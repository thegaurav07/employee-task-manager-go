# Employee Task Manager (Go + PostgreSQL)
 
## Project Overview
 
Employee Task Manager is a full-stack web application built using Golang, Gin Framework, PostgreSQL, GORM, HTML, CSS, and JavaScript.
 
This project allows users to:
 
- Create tasks
- View tasks
- Store task data in PostgreSQL database
- Connect frontend with backend APIs
 
---
 
# Tech Stack
 
## Backend
 
- Golang
- Gin Framework
- GORM
- PostgreSQL
 
## Frontend
 
- HTML
- CSS
- JavaScript
 
## Tools
 
- VS Code
- Git
- GitHub
- Postman
 
---
 
# Project Structure
 
```bash
employee-task-manager/
│
├── controllers/
│   └── taskcontroller.go
│
├── database/
│   └── db.go
│
├── models/
│   └── task.go
│
├── routes/
│   └── routes.go
│
├── static/
│   ├── style.css
│   └── script.js
│
├── templates/
│   └── index.html
│
├── main.go
├── go.mod
└── go.sum
```
 
---
 
# Features
 
- Add new tasks
- View all tasks
- REST API integration
- PostgreSQL database connection
- Frontend and backend integration
- Responsive UI
 
---
 
# Installation Steps
 
## 1. Clone Repository
 
```bash
git clone https://github.com/thegaurav67/employee-task-manager-go.git
```
 
---
 
## 2. Open Project
 
```bash
cd employee-task-manager-go
```
 
---
 
## 3. Install Dependencies
 
```bash
go mod tidy
```
 
---
 
## 4. Configure PostgreSQL
 
Update database connection inside:
 
```bash
database/db.go
```
 
Example:
 
```go
dsn := "host=localhost user=postgres password=12345 dbname=taskdb port=5432 sslmode=disable"
```
 
---
 
## 5. Run Project
 
```bash
go run main.go
```
 
---
 
# API Endpoints
 
## Create Task
 
```http
POST /tasks
```
 
## Get Tasks
 
```http
GET /tasks
```
 
---
 
# Frontend
 
Frontend files are available inside:
 
```bash
static/
```
 
and
 
```bash
templates/
```
 
---
 
# GitHub Commands Used
 
## Initialize Git
 
```bash
git init
```
 
## Add Files
 
```bash
git add .
```
 
## Commit Changes
 
```bash
git commit -m "Initial Commit"
```
 
## Connect GitHub Repository
 
```bash
git remote add origin <repository-link>
```
 
## Push Code
 
```bash
git push -u origin main
```
 
---
 
# Future Improvements
 
- Update task feature
- Delete task feature
- User authentication using JWT
- Docker support
- Deployment on Render/Railway
 
---
 
# Author
 
Kumar Gaurav
 
GitHub:
 
https://github.com/thegaurav67
