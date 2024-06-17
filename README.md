# Chat Room Application

This is a simple chat room application built in Go. It uses websockets to provide real-time communication between the server and the clients.

## How It Works

The application is structured around the concept of a Room, which is a chat room where clients can join, leave, and send messages. Each Client represents a user connected to the room. Messages sent by a client are forwarded to all other clients in the room.

The application uses the Gorilla Websocket library for handling websocket connections.

## Required Libraries

This application requires the following Go libraries:
- Gorilla Websocket v1.5.3

## Installation

1. Ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. Clone this repository to your local machine:
   ```sh
   git clone <repository-url>

3. Navigate to the project directory:
    ```sh
    cd <project-directory>
    ```
4. Download the required Go dependencies:
    ```sh
        Copy code
        go mod download
    ```
5. Build the application:
    ```sh
    Copy code
    go build
    ```
6. Run the application:
    ```sh
    ./<application-name>
    ```
By default, the application will start a server on localhost:8080. You can connect to the chat room by opening a web browser and navigating to http://localhost:8080.

## Files
- main.go: This is the entry point of the application. It sets up the HTTP server and the chat room.
- room.go: This file defines the Room type and its methods. A Room represents a chat room where clients can join, leave, and send messages.
- client.go: This file defines the Client type and its methods. A Client represents a user connected to the room.
- templates/chat.html: This is the HTML template for the chat room interface.
- templates/js/chat.js: This is the JavaScript file that handles the client-side logic of the chat room.