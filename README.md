# REST API for TASK-MANAGEMENT

## In this project:
- Developed web-application following REST API design.
- Framework <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Clean architecture in building structure of application. Built dependencies.
- Work with database MySQL. Launch from Docker. Migrations.
- Configuration with help of <a href="https://github.com/spf13/viper">spf13/viper</a>. Work with environment variables.
- Work with MySQL with help of library <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Registration and authentication. Work with JWT. Middleware.
- SQL queries.
- Graceful Shutdown.

### Endpoints:

### - POST /auth/sign-up

Creates new user

##### Example Input:
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input:
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```

##### Example Response:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /api/tasks

Creates new task

##### Example Input:
```
{
    "title": "https://github.com/5t4lk/task-management",
	"description": "REST API project",
	"status": "in progress",
	"end_date": "31.06.2023"    
} 
```

### GET /api/tasks

Returns all user tasks

##### Example Response:
```
{
	"tasks": [
            {
                "id": "1",
                "title": "https://github.com/5t4lk/task-management",
                "description": "REST API project"
            }
    ]
} 
```

### DELETE /api/tasks

Deletes task by ID:

##### Example Input:
```
{
	"id": "1"
} 
```


## Requirements
- go 1.20
- docker & docker-compose


### To launch:

```
make build
make run
```

If you are launching application first time, you need to apply migrations:

```
make migrate
```
