# Go-React-Chat

This is a Websocket chat room implementation with Golang backend, React Typescript frontend and PostgresQL
User could sign up and login to the public chat room, could see all the history chat records and realtime incoming messages, user could also vote up or down to others messages.
user login is token based, after sign in with valid credential, server will sign a jwt token with secret key and let client keep it. Client need to send request with token to verify their identity.
Sending and receiving chats are through webSocket, after login, user will ask to connect server through websocket. whenever server receive a new message or upvote/downvote, server will handle the message,
create new record/ update exist record in database and broadcast to every connected clients.

## Package
- gin
- gorm
- jwt
- gorilla/websocket
- react-use-websocket

## Docker
- build docker images `docker-compose build`
- run db `docker-compose up db -d`
- run backend `docker-compose up api -d`, backend will start listen on port 8080 
- run frontend `docker-compose up frontend -d`, react will be served on `http://localhost:3000/`

## Build Locally
- run db `docker-compose up db -d`
- switch to backend directory, build go with `go build ./main.go`, then run `./main`
- switch to frontend directory, install dependencies with `npm i`, then run `npm start`
- backend will be listening on port 8080, and react could be access on `http://localhost:3000/`

