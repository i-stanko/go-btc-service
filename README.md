# go-btc-service
**go-btc-service** is a simple API service that allows you to retrieve the current exchange rate of Bitcoin (BTC) to Ukrainian Hryvnia (UAH), subscribe for rate change notifications, and send rate updates to subscribed email addresses. It is implemented using the Go programming language.

[![asciicast](https://asciinema.org/a/587013.svg)](https://asciinema.org/a/587013)

## Requirements
- Go 
- Docker

## Installation and Usage
- Clone the repository
```
git clone https://github.com/i-stanko/go-btc-service.git
```

- Navigate to the project directory
```
cd go-btc-service
```

- Start the server
```
go run main.go data.go 
```

The server will be available at `http://localhost:8080`. Make requests to the API using an API client such as curl or Postman, or make requests from your own program.

## API

- **Get Current BTC to UAH Exchange Rate.** This endpoint returns the current exchange rate of Bitcoin to Ukrainian Hryvnia.
```
curl http://localhost:8080/api/rate
```

- **Subscribe Email for Rate Updates.** This endpoint adds an email address to the subscription list for receiving rate updates.
```
curl -X POST \
    -d "email=mail@example.com" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    http://localhost:8080/api/subscribe
```

- **Send Rate Updates to Subscribed Emails.** This endpoint sends rate updates to all subscribed email addresses.
```
curl -X POST \
    -H "Content-Type: application/json" \
    http://localhost:8080/api/sendEmails
```

## Docker
To deploy this service in a Docker container, follow these steps:

- Build the Docker image
```
docker build -t go-btc-service .
```

- Run the Docker container
```
docker run -p 8080:8080 go-btc-service
```
The server will be available at `http://localhost:8080`.

## Additional Information

- Data is stored and managed using the file system.
- Swagger documentation is available at [gses2swagger.yaml](https://github.com/AndriiPopovych/gses/blob/main/gses2swagger.yaml).
- If you have any further questions or need assistance, please let me know.
