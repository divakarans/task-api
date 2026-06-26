# Task Management Application

A full-stack Task Management application built with **React**, **Go (Gin)**, and **SQLite**. The application allows users to create, view, update, and delete tasks through a clean and responsive user interface.

---

## 🚀 Features

* Create new tasks
* View all tasks in a modal
* Edit existing tasks
* Delete tasks
* Persistent data storage using SQLite
* RESTful API built with Gin
* Responsive React frontend
* Modal-based task management interface

---

## 🛠️ Tech Stack

### Frontend

* React
* Vite
* JavaScript
* CSS

### Backend

* Go
* Gin Web Framework
* SQLite

---

## 📂 Project Structure

```text
task-api/
│
├── backend/
│   ├── db/
│   ├── handlers/
│   ├── models/
│   ├── store/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── tasks.db
│
├── frontend/
│   ├── public/
│   ├── src/
│   ├── package.json
│   ├── vite.config.js
│   └── ...
│
└── README.md
```

---

## ⚙️ Backend Setup

Navigate to the backend directory:

```bash
cd backend
```

Install dependencies:

```bash
go mod tidy
```

Run the server:

```bash
go run .
```

The backend runs on:

```
http://localhost:8080
```

---

## ⚙️ Frontend Setup

Navigate to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm run dev
```

The frontend runs on:

```
http://localhost:5173
```

---

## 📡 API Endpoints

| Method | Endpoint     | Description             |
| ------ | ------------ | ----------------------- |
| GET    | `/tasks`     | Retrieve all tasks      |
| GET    | `/tasks/:id` | Retrieve a task by ID   |
| POST   | `/tasks`     | Create a new task       |
| PUT    | `/tasks/:id` | Update an existing task |
| DELETE | `/tasks/:id` | Delete a task           |

---

# 📸 Application Screenshots

## Home Page

<img width="1918" height="867" alt="Home" src="https://github.com/user-attachments/assets/834ccb90-cddf-4b96-9cac-4869d2849fff" />

---

## Task List

<img width="1918" height="868" alt="Task List" src="https://github.com/user-attachments/assets/a317f06f-2b02-4725-a851-66dc66fce13d" />

---

## Edit Task

<img width="1918" height="870" alt="Edit Task" src="https://github.com/user-attachments/assets/a999f726-c90f-4a00-97f7-e3e563e2d81c" />

---

## 📖 Learning Outcomes

This project helped in understanding:

* React Components
* Props
* State Management using `useState`
* Side Effects using `useEffect`
* Controlled Forms
* Conditional Rendering
* Parent-Child Communication
* CRUD Operations
* REST API Integration
* Fetch API
* Go Gin Framework
* SQLite Database
* Backend Architecture
* Component-based UI Design

---

## 🔮 Future Improvements

* Authentication & Authorization
* Task Search
* Task Filtering
* Task Sorting
* Pagination
* Toast Notifications
* Confirmation Modal for Delete
* React Router
* Axios Integration
* TanStack Query
* React Hook Form
* Dark Mode
* Drag & Drop Task Management

---

## 👨‍💻 Author

**Divakaran S**

GitHub: https://github.com/divakarans
