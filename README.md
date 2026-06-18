# Task API

A simple RESTful Task Management API built with Go using the standard library `net/http` and `ServeMux`.

## Features

* Create Tasks
* Get All Tasks
* Get Task By ID
* Update Tasks
* Delete Tasks
* JSON Request and Response Handling
* Route Parameters using ServeMux
* Input Validation
* In-Memory Storage

## Project Structure

```text
task-api/
│
├── main.go
├── go.mod
│
├── handlers/
│   └── handlers.go
│
├── models/
│   └── task.go
│
└── store/
    └── store.go
```

## Tech Stack

* Go
* net/http
* ServeMux
* JSON Encoding/Decoding

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd task-api
```

Install dependencies:

```bash
go mod tidy
```

## Run the Application

```bash
go run .
```

Server starts on:

```text
http://localhost:8080
```

## API Endpoints

### Create Task

```http
POST /tasks
```

Request:

```json
{
  "title": "Learn Go",
  "status": "pending"
}
```

Response:

```json
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending"
}
```

---

### Get All Tasks

```http
GET /tasks
```

Response:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "status": "pending"
  }
]
```

---

### Get Task By ID

```http
GET /tasks/{id}
```

Example:

```http
GET /tasks/1
```

---

### Update Task

```http
PUT /tasks/{id}
```

Request:

```json
{
  "title": "Master Go",
  "status": "done"
}
```

---

### Delete Task

```http
DELETE /tasks/{id}
```

Response:

```json
{
  "message": "Task deleted successfully"
}
```

## Task Model

```json
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending",
  "created_at": "2025-09-01T10:00:00Z"
}
```

## Author

Divakaran
