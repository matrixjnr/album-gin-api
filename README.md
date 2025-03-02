
# Album Gin API

This project is a RESTful API for managing a collection of music albums, built using the Gin framework in Go.

## Setup Instructions

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/album-gin-api.git
    cd album-gin-api
    ```

2. **Install dependencies:**
    ```sh
    go mod download
    ```

3. **Run the application:**
    ```sh
    go run main.go
    ```

4. **Access the API:**
    Open your browser or API client and navigate to `http://localhost:8080/albums`

## Endpoints

- `GET /albums` - Retrieve a list of albums
- `POST /albums` - Add a new album
- `GET /albums/:id` - Retrieve a specific album by ID
- `PUT /albums/:id` - Update an album by ID
- `DELETE /albums/:id` - Delete an album by ID

## License

This project is licensed under the MIT License.
