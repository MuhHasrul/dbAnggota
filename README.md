# dbAnggota

Projek ini adalah aplikasi berbasis web untuk manajemen data anggota, status anggota, dan badan pengurus harian. Backend aplikasi ini menggunakan MySQL, sedangkan frontend menggunakan React. Projek ini dikontainerisasi menggunakan Docker untuk memudahkan deployment.

## Tujuan

Projek ini dibuat oleh MarvelStudioXII dengan tujuan untuk:
- Melihat total anggota
- Memantau status anggota
- Mengelola badan pengurus harian

## Struktur Projek

	```bash
	dbAnggota/
	├── backend/
	│ ├── Dockerfile
	│ ├── src/
	│ ├── .env
	│ ├── db.sql
	│ └── ...
	├── frontend/
	│ ├── Dockerfile
	│ ├── src/
	│ ├── public/
	│ ├── package.json
	│ └── ...
	├── docker-compose.yml
	└── README.md



## Prasyarat

Pastikan Anda telah menginstal Docker dan Docker Compose di mesin Anda.

## Setup dan Menjalankan Projek

1. ## Clone repositori ini 

   ```bash
   git clone https://github.com/MarvelStudioXII/dbAnggota.git
   cd dbAnggota

2. ## Setup backend

	- Masuk ke direktoru 'backend'
	cd backend
	- Buat file .env dan tambahkan konfigurasi database Anda
	```bash
	DB_HOST=db
	DB_USER=root
	DB_PASSWORD=myp4ssword
	DB_NAME=dbAnggota

3. ## Setup Frontend

	- Masuk ke direktori frontend
	cd frontend
	- Install dependencies
	npm install
4. ## Jalankan dengan Docker Compose

	- Kembali ke direktori root projek
	cd ..
	- Jalankan Docker Compose
	```bash
	docker-compose up --build
5. ## Akses Aplikasi
	- Frontend dapat diakses di http://localhost:3000
	- Backend API dapat diakses di http://localhost:5000

***Struktur Docker***

## Backend
Dockerfile untuk backend terletak di backend/Dockerfile:
# Gunakan image MySQL sebagai dasar
FROM mysql:5.7

## Salin skrip SQL ke dalam container
COPY db.sql /docker-entrypoint-initdb.d/

## Set environment variables
ENV MYSQL_ROOT_PASSWORD=yourpassword
ENV MYSQL_DATABASE=dbAnggota

## Frontend
Dockerfile untuk frontend terletak di frontend/Dockerfile:
## Gunakan image Node.js sebagai dasar
FROM node:14

***Set working directory***
WORKDIR /app

***Salin package.json dan install dependencies***
COPY package.json ./
RUN npm install

***Salin seluruh proyek***
COPY . .

***Build aplikasi React***
RUN npm run build

***Expose port***
EXPOSE 3000

***Jalankan aplikasi***
CMD ["npm", "start"]

***Docker-compose***
`docker-compose.yml` menghubungkan frontend dan backend:
version: '3.8'

