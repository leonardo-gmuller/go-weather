
# Go Weather

**Go Weather** is a Go-powered API that retrieves the current weather of a city based on a Brazilian postal code (CEP). It integrates with external APIs to fetch address and weather data, and returns a clean JSON response with temperature and location info.

## ✨ Technologies Used

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Google Cloud Run](https://cloud.google.com/run)
- External APIs:
  - [ViaCEP](https://viacep.com.br/)
  - [WeatherAPI](https://www.weatherapi.com/)

## 🏗️ Architecture

The project follows a clean architecture pattern with the following structure:

```
app/
├── cmd/
│   └── server/         # Application entry point (main package)
├── config/             # Configuration loading (env, flags, etc.)
├── domain/             # Core business logic and entities
│   ├── dto/            # Data Transfer Objects
│   └── usecase/        # Business use cases and interfaces
├── gateway/            # External service integrations
│   ├── api/            # Handlers for this API's endpoints
│   └── client/         # External API clients and adapters (e.g.,      WeatherAPI, ViaCEP)
tests/
└── mocks/              # Test mocks and interfaces
```

The application uses dependency injection to keep the components loosely coupled and testable.

## 🚀 Running Locally

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/go-weather.git
    cd go-weather
    ```

2. Set required environment variables:

    - Copy `.env_template` to `.env` and fill in your values:
      ```bash
      cp .env_template .env
      ```
      Edit `.env` to set your `WEATHER_API_KEY` and any other required variables.

    **OR** export them directly in your shell:
    ```bash
    export WEATHER_API_KEY=your_weatherapi_key
    export PORT=8080
    # Add any other required variables here
    ```

3. Run the application locally with Go:
    ```bash
    go run ./app/cmd/server
    ```

    **OR** with [Air](https://github.com/cosmtrek/air) for live reloading:
    ```bash
    air
    ```

**OR** using Docker:

```bash
docker build -t go-weather .
docker run -p 8080:8080 --env-file .env go-weather
```

Once running, the app will be available at:  
`http://localhost:8080/api/v1/weather/{cep}`

Replace `{cep}` with the desired Brazilian postal code.

## 🧪 Running Tests

Automated tests are located in their respective directories, such as `handler_test.go` inside the `handlers` package.

To run tests for a specific package:

```bash
go test ./app/gateway/api/handler
```

To run all tests in the project:

```bash
go test ./...
```

## 📦 Example Request (DEMO)

Example using `curl`:

```bash
curl -X GET "https://go-weather-w7igmsgiwa-uc.a.run.app/api/v1/weather/01001000"
```

Expected response:

```json
{
  "temp_C": 22.5,
  "temp_F": 72.5,
  "temp_K": 295.65
}
```

## ⚠️ Error Handling

The system responds appropriately under the following scenarios:

### ✅ Success
- **HTTP Status**: `200 OK`
- **Response Body**:
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.65
  }
  ```

---

### ❌ Invalid ZIP Code (correct format, but invalid content)

- **HTTP Status**: `422 Unprocessable Entity`
- **Response Body**:
  ```json
  {
    "error": "invalid zipcode"
  }
  ```

---

### ❌ ZIP Code Not Found

- **HTTP Status**: `404 Not Found`
- **Response Body**:
  ```json
  {
    "error": "can not find zipcode"
  }
  ```

---

## ✅ Project Status

This project was built for educational purposes to practice backend development using Go, consuming external APIs, and deploying to a serverless environment on Google Cloud.

## 📄 License

No license has been defined for this project.
