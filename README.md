# LRU-Cache web Application

This project is a web application that implements a Least Recently Used (LRU) cache system. It consists of a backend server built with Go (Golang) and a frontend client built with React.

## Features 
- Implements an LRU cache with a fixed size.
- Supports setting and getting key-pairs with expiration times.
- Periodically evicts expired item from the cache.
 
## Getting Started
 
Follow these steps to run the application locally:
 
### Prerequisites

- GO
- React JS
- Docker
- Docker Compose
 
### Instructions
 
1. Clone the repository: git clone https://github.com/pratibh13/LRU-Cache.git
 
2) Navigate to the project directory: cd LRU-Cache
3) Build and run the application using Docker Compose: docker-compose up --build
Once the containers are up and running, you can access the frontend app at http://localhost:3000 and the backend API at http://localhost:8080.
   - OR ( Mostly Use this one) 
   1) Split the terminal
   2) Navigate to backend and run go run main.go // these will start your backend server
   3) On second terminal Navigate to: frontend/lru-cache-app  then run npm start // these will start your client server 
 
### USAGE
Use the frontend application to interact with the LRU cache system:
Enter a key-value pair and click "Set Key" to add it to the cache.
Enter a key and click "Get Key" to retrieve its corresponding value from the cache.

## API Endpoints
- GET /cache?key=<key>: Retrieve the value associated with the specified key from the cache.
- POST /cache: Set a new key-value pair in the cache. Example JSON request body: {"key": "mykey", "value": "myvalue"}

## Dependencies
This project uses the following dependencies:
 
- github.com/gorilla/mux for routing HTTP requests in the server.
- github.com/rs/cors for handling Cross-Origin Resource Sharing (CORS) in the server.
- You can install these dependencies using the go get command: go get github.com/gorilla/mux github.com/rs/cors
### CONTRIBUTING
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or create a pull request.
