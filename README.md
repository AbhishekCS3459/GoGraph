# GoGraph API

GoGraph is a simple API built with Go that finds all possible paths between two nodes in a directed graph using Depth-First Search (DFS).

## Features
- Health check endpoint
- Find all paths between two nodes in a directed graph
- Built using `gorilla/mux` for routing

## Installation & Setup

### Prerequisites
- Go 1.18+

### Clone the Repository
```sh
git clone https://github.com/AbhishekCS3459/goGraph.git
cd goGraph
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Server
```sh
go run cmd/main.go
```
The server will start on **http://localhost:8080**.

---

## API Endpoints

### 1. Health Check
- **Endpoint:** `/`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "message": "I am alive!"
  }
  ```

### 2. Find Paths in a Graph
- **Endpoint:** `/find-path`
- **Method:** `POST`
- **Headers:**
  - `Content-Type: application/json`
- **Request Body:**
  ```json
  {
    "graph": {
      "edges": {
        "0": [1, 2],
        "1": [2, 3],
        "2": [3],
        "3": [4],
        "4": []
      }
    },
    "startNode": 0,
    "endNode": 4
  }
  ```
- **Response (Example):**
  ```json
  {
    "paths": [
      [0, 1, 3, 4],
      [0, 2, 3, 4],
      [0, 1, 2, 3, 4]
    ]
  }
  ```

### Possible Errors
- If the graph is empty:
  ```json
  {
    "error": "Empty graph"
  }
  ```
- If the start node has no outgoing edges:
  ```json
  {
    "error": "Start node has no outgoing edges"
  }
  ```
- If no path exists:
  ```json
  {
    "error": "No path exists between start and end nodes"
  }
  ```

## Testing with Postman
1. Open **Postman**.
2. Create a new request.
3. Set method to `POST`.
4. Enter URL: `http://localhost:8080/find-path`.
5. Go to **Body > raw**, select `JSON`, and paste the request body.
6. Click **Send** to see the response.

## License
This project is open-source under the MIT License.

---

### Author
Developed by [Abhishek Kumar Verma](https://github.com/AbhishekCS3459).

