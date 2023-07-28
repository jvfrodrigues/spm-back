This is the backend for my simple password manager application, it provides a REST API with routes for acessing, creating, deleting and updating password entries.

## Getting Started

### Stack:

The application runs on [Go](https://go.dev/) 1.19 for creating my API I use a package that I've worked with before called [Gin](github.com/gin-gonic/gin).
I've also created a Dockerfile to run the project.

### Running:

Locally:

```bash
go run ./main.go
```

Docker:

```bash
docker run -p 8080:8080 spm-back
```

The api will run on [http://localhost:8080](http://localhost:8080) you can call it using from cmd or using an API development platform like [Insomnia](https://insomnia.rest/) I provide a collection with the available routes in the project.

## Approach

I wanted to do something very simple but still follow the patterns I use in my day to day work.
Normally I use SOLID principles in my projects so I made sure to follow them here.
For instance my password_controller calls the repository without knowing how the repository does the actual actions of creating, deleting and finding the passwords that I want to create. Today we are using an in memory variable that keeps all are passwords while the server runs but if in the we were to add use a database to store our passwords we would be able to only change our repository without affecting the rest of our application.
