# GO URL Shortener
![Shiva's URL Shortener](https://github.com/Shiva-2103/GO_URL_SHORTENER/blob/main/GO_URL.jpg)



A simple and efficient URL Shortener built using Golang, Redis, Docker Compose, and the Fiber framework. This project allows users to shorten long URLs and easily redirect to the original URLs using the generated short links.

## Features

- **Shorten URLs**: Generate compact and unique short URLs.
- **Fast and Lightweight**: Built using Golang and the Fiber framework for high performance.
- **Persistent Storage**: Uses Redis for storing URL mappings.
- **Dockerized**: Fully containerized using Docker and Docker Compose for seamless deployment.
- **Scalable**: Designed to handle high traffic with ease.

## Tech Stack

- **Backend**: Golang
- **Framework**: Fiber
- **Database**: Redis
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Docker and Docker Compose installed on your system.
- Go installed (for local development without Docker).

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/Shiva-2103/GO_URL_SHORTENER.git
cd GO_URL_SHORTENER
```
## Running with Docker Compose
1. Build and start the containers:
```bash
docker-compose up --build
```
2. The application will be accessible at ``` http://localhost:3000. ```

## Running Locally without Docker
1. Install dependencies:
```bash
go mod tidy
```
2. Start the Redis server (if not already running):
```bash
redis-server
```
3. Run the application:
```bash
go run main.go
```
4. The application will be accessible at ``` http://localhost:3000. ```

## API Endpoints
## Shorten URL
POST  ``` /shorten ```
Body:
``` json
{
  "url": "https://example.com",
  "expiry": 3600
}
```
Response:
```json
{
  "short_url": "http://localhost:3000/<short_code>"
}
```
## Redirect to Original URL
GET /<short_code>
Redirects the user to the original URL.

## Project Structure
```
GO_URL_SHORTENER/
│
├── main.go          # Entry point of the application
├── handler/         # Handlers for API endpoints
├── router/          # Routes and middleware
├── storage/         # Logic for Redis interactions
├── Dockerfile       # Dockerfile for containerizing the app
├── docker-compose.yml # Docker Compose configuration
├── go.mod           # Module dependencies
└── README.md        # Project documentation
```

## Future Enhancements
 - Add analytics for short URL usage (e.g., click counts, last accessed time).
 - Implement user authentication for personalized URL shortening.
 - Support for custom short URL aliases.
 - Enhanced error handling and logging.

## Contributing
Contributions are welcome! Feel free to fork this repository, create a new branch, and submit a pull request.
