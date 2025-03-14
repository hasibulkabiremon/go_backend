# Simple Go Web Server

A basic web server implemented in Go that demonstrates handling static files, form submissions, and basic routing.

## Features

- Static file serving
- Form handling with POST requests
- Basic routing with multiple endpoints
- Error handling for invalid routes and methods

## Project Structure 

## Setup and Running

1. Make sure you have Go installed on your system
2. Clone this repository
3. Navigate to the web_server directory
4. Run the server:
```bash
go run main.go
```
The server will start at `http://localhost:8080`

## Available Endpoints

### 1. Static File Server (`/`)
- Serves static files from the `static` directory
- Example: `http://localhost:8080/form.html`

### 2. Hello Endpoint (`/hello`)
- Method: GET
- Returns a welcome message
- Example Response: "Welcome to my website!"
- Error handling for incorrect methods or paths

### 3. Form Handler (`/form`)
- Method: POST
- Handles form submissions with name and address fields
- Example form submission response: 

## Example Usage

### Accessing the Form
1. Open your browser and go to `http://localhost:8080/form.html`
2. Fill out the name and address fields
3. Submit the form to see the processed results

### Testing the Hello Endpoint
- Visit `http://localhost:8080/hello` in your browser
- You should see the welcome message

## Error Cases

The server handles various error cases:
- 404 Not Found for invalid paths
- Method not supported for incorrect HTTP methods
- Form parsing errors

## API Examples

### Using cURL

1. Test the hello endpoint:
```bash
curl http://localhost:8080/hello
```

2. Submit a form:
```bash
curl -X POST -d "name=John&address=123%20Main%20St" http://localhost:8080/form
```

## Notes

- The server runs on port 8080 by default
- Static files must be placed in the `static` directory to be served
- Form submissions are processed but not stored (for demonstration purposes only)