
# Go URL Shortener API with Redis

This is a simple and efficient URL shortening backend service written in Go, using Redis for storage. It allows users to submit long URLs and receive short redirectable codes in return.


## Acknowledgements

 - [Awesome Readme Templates](https://awesomeopensource.com/project/elangosundar/awesome-README-templates)
 - [Awesome README](https://github.com/matiassingers/awesome-readme)
 - [How to write a Good readme](https://bulldogjob.com/news/449-how-to-write-a-good-readme-for-your-github-project)


## Features

- Generate unique short codes for any URL
- Store mappings in Redis for fast access
- Handle redirection via short URLs
- Easily test with Postman or Curl
- Lightweight and easy to deploy


## Tech Stack

**Language**: Go (Golang)
- **Framework**: Gorilla Mux
- **Database**: Redis
- **HTTP**: net/http package
- **Others**: JSON handling, RESTful API



## Installation

1. Clone the repo:
   ```bash
   git clone https://github.com/BarakaCaleb/url-shortener-go.git
   cd url-shortener-go
    
    go mod tidy


    docker run -d -p 6379:6379 redis


    go run main.go





# 📡 API Endpoints
#### POST /shorten

Request:

```
{
  "long_url": "https://example.com/some/long/path"
}
```

Response:

```
{
  "short_code": "abc123",
  "short_url": "http://localhost:8080/abc123"
}
```

#### GET /{code}

Redirects the user to the original long URL associated with the short code.

Example:

```
GET http://localhost:8080/abc123
```


## Run Locally
```
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"long_url": "https://example.com"}'

```


## License

[MIT](https://choosealicense.com/licenses/mit/)


