# Weather API Documentation

https://roadmap.sh/projects/weather-api-wrapper-service

## Overview

The Weather API is a simple service that retrieves weather data for a specified city. It supports caching, leverages OpenWeatherMap API, and requires an API key for access.

---

## Requirements

- **Go**: Version 1.23.4 or higher.
- **OpenWeatherMap API Key**: Required for fetching weather data.
- **Environment Variables**: Define the following in a `.env` file or system variables:

  ```plaintext
  PORT=8080
  OPENWEATHER_API_KEY=<your_openweathermap_api_key>
  ```

- **Dependencies**:
  Install dependencies with:
  ```bash
  go get github.com/joho/godotenv
  go get github.com/rs/zerolog
  ```

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd Weather-API
   ```

2. Build the application:
   ```bash
   go build -o weather-api.exe
   ```

3. Create a `.env` file in the root directory and add your configuration:
   ```plaintext
   PORT=8080
   OPENWEATHER_API_KEY=<your_openweathermap_api_key>
   ```

4. Run the application:
   ```bash
   .\weather-api.exe
   ```

---

## API Usage

### Base URL

```
http://localhost:<PORT>
```

### Endpoints

#### `GET /`

- **Description**: Fetches weather data for a specified city.
- **Headers**:
    - `X-API-Key`: Required API key for authentication.
- **Query Parameters**:
    - `city`: Name of the city (default: Moscow).

#### Example Request (PowerShell)

```powershell
$Headers = @{ "X-API-Key" = "<your_api_key>" }
$Url = "http://localhost:8080/?city=Moscow"
$response = Invoke-WebRequest -Uri $Url -Headers $Headers -Method GET
$response.Content
```

#### Example Response

```json
{
  "city": "Moscow",
  "temperature": 2.24,
  "units": "Celsius"
}
```

#### Error Response

If the API key is missing or invalid:

```json
{
  "error": "API key is missing or invalid"
}
```

---

## Development Notes

- **Logging**:
    - Uses `zerolog` for structured logging.
    - Logs are displayed in a human-readable format on the console.
- **Environment Management**:
    - `.env` file is used to load environment variables via `godotenv`.
- **Caching**:
    - In-memory caching is implemented using a mock cache (`MockCache`) for simplicity.

---

## Future Improvements

1. Integrate a persistent caching layer (e.g., Redis).
2. Add support for more query parameters (e.g., units: metric/imperial).
3. Implement rate limiting for API requests.
4. Add Docker support for containerized deployment.

---

## Support

For any questions or issues, please contact the project maintainer at: cooler19860212@yahoo.com.

