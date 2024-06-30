package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	NRA      	string `json:"nra"`
	Nama		string `json:"nama"`
	Email    	string `json:"email"`
	No_Wa    	string `json:"no_wa"`
	Password 	string `json:"password" validate:"required,min=8"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var db *sql.DB

func generateRandomKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

var validate *validator.Validate

func GeneratePasswordHash(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}


func main() {
	var err error

	validate = validator.New()

	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_anggota")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/signup", SignupHandler).Methods("POST")

	r.HandleFunc("/signin", SigninHandler).Methods("POST")

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	fmt.Println("Server is running on port 4040")
	log.Fatal(http.ListenAndServe(":4040", corsMiddleware(r)))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	// Validate email format
	if !isValidEmail(user.Email) {
		http.Error(w, `{"error": "Invalid email format"}`, http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		log.Println("Validation error:", err)
		http.Error(w, `{"error": "Invalid input data"}`, http.StatusBadRequest)
		return
	}

	fmt.Printf("User baru: %+v\n", user)

	// Hash the user's password
	hashedPassword, err := GeneratePasswordHash(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	// Insert user into database
	_, err = db.Exec("INSERT INTO signup (NRA, NamaLengkap, Email, NoWa, Password) VALUES (?, ?, ?,?,?)", user.NRA, user.Nama, user.Email, user.No_Wa, hashedPassword)
	if err != nil {
		log.Println("Error inserting user into database:", err)
		http.Error(w, `{"error": "Failed to create user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "User created successfully"}`))
}

func isValidEmail(email string) bool {
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailRegex)

	// Validate the email against the regex pattern
	return re.MatchString(email)
}


func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	err = validate.StructPartial(user, "Password")
	if err != nil {
		log.Println("Validation error:", err)
		http.Error(w, `{"error": "Invalid input data"}`, http.StatusBadRequest)
		return
	}

	// Menentukan inputUsername berdasarkan input pengguna
	var inputUsername string
	if user.Email != "" {
		inputUsername = user.Email
	} else if user.NRA != "" {
		inputUsername = user.NRA
	} else {
		http.Error(w, `{"error": "NRA or Email required"}`, http.StatusBadRequest)
		return
	}

	var storedPasswordHash string
	err = db.QueryRow("SELECT password FROM signup WHERE nra = ? OR email = ?", inputUsername, inputUsername).Scan(&storedPasswordHash)
	if err != nil {
		log.Println("Error querying database:", err)
		if err == sql.ErrNoRows {
			http.Error(w, `{"error": "Invalid NRA or password"}`, http.StatusUnauthorized)
		} else {
			http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		}
		return
	}

	// Memverifikasi password dengan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(user.Password))
	if err != nil {
		log.Println("Password does not match stored hash:", err)
		http.Error(w, `{"error": "Invalid NRA or password"}`, http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: inputUsername, // Menggunakan inputUsername yang ditentukan
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey, err := generateRandomKey(32)
	if err != nil {
		log.Println("Error generating JWT key:", err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error signing JWT token:", err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", tokenString)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + tokenString + `", "redirect": "dashboard"}`))
}


