# XM-Golang-Exercise

## Description
A RESTful microservice that does CRUD operations for companies through protected endpoints that requires a JWT token. After any mutation operation the service produces a Kafka event to notify other microservices that uses this information.

## How To Use

1. Make sure you have Docker installed on your local machine.
2. Run the following command to start the database:
```
docker-compose -f ./external/database/docker-compose.yml up -d
``` 
3. Build the Dockerfile for the service using the following command:
```
docker build . -t xmgolangexercise:latest
```
4. Run the service by running the following command:
```
docker-compose up -d
```
5. You can access the service from the following url:
```
http://localhost:8000
```

## Documentation
You can find the documentation of the api in the *openapi.yaml* file which follows OpenAPI 3.0.3 spec and you can inspect it here: <https://editor.swagger.io>

## Database
I decided to use PostgreSQl database and run it locally through Docker.

## Kafka
I decided to use Kafka cloud because I ran into some issues trying to run it locally.

## Secrets
I keep the secrets in the repo just for the sake of the task but if it was an actual deployment I would store the secrets more safely.