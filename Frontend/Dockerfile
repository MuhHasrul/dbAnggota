#Use the official node.js image as a base
FROM node:20-alpine

#Set the working directory in the container

WORKDIR /app

# Copy package.json dan package-lock.json file

COPY pakcage*.json.

RUN npm install

COPY . .

EXPOSE 5173

CMD ["npm","run","dev"]
