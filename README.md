# url-shortener

This project is an educational application that implements the functionality of shortening long URLs. The project works as an HTTP server and a gRPC server, handling URL shortening requests and providing an API for managing the data.

### Requirements

- Golang
- Postgres

## Usage

### HTTP Server

The HTTP server handles URL shortening requests and provides an API for managing the data. The following are the available HTTP server handlers:

- GET /{id} - Get the full URL by its identifier (id)
- POST / - Create a new shortened URL
- POST /api/shorten - Create a new shortened URL via the API
- GET /api/user/urls - Get a list of all shortened URLs for the user
- POST /api/shorten/batch - Create multiple shortened URLs at once
- DELETE /api/user/urls - Delete all shortened URLs for the user

### gRPC Server

The gRPC server is an alternative way to interact with the application. It provides support for remote procedure calls and supports the same functions as the HTTP server. For more information on using gRPC, refer to the gRPC documentation.
