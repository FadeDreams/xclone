## XClone Application

XClone is a Go-based application built with a modular architecture, integrating Chi as a router, PostgreSQL as the database, and JWT for secure authentication and authorization. The codebase implements a Hexagonal Architecture, separating concerns into domain, service, and adapter layers. Database migrations are handled using the PostgreSQL package, and the server is configured to run on port 8080.

## Overview

- **Modular Architecture:** Built with a modular structure, incorporating Chi as a router, PostgreSQL for the database, and JWT for secure authentication and authorization.

- **Hexagonal Architecture:** The codebase follows a Hexagonal Architecture, distinguishing domain, service, and adapter layers to maintain separation of concerns.

- **Database Migrations:** PostgreSQL package is used for seamless handling of database migrations.

- **GraphQL API:** Powered by gqlgen, the application provides a GraphQL playground at the root endpoint and a query endpoint at "/query".

- **Testing:** The application includes comprehensive test coverage to ensure reliability and robustness.

- **Token Service:** Innovative token service is employed for user authentication.

## Project Structure
```plaintext
/xclone
|-- auth.go
|-- auth_test.go
|-- config
|-- domain
|-- go.sum
|-- Makefile
|-- README.md
|-- tweet.go
|-- context.go
|-- faker
|-- gqlgen.yml
|-- init-scripts
|-- mocks
|-- refresh_token.go
|-- user.go
|-- cmd
|-- docker-compose.yml
|-- go.mod
|-- graph
|-- jwt
|-- postgres
|-- server_go
|-- uuid
```

## Getting Started
To test the app, follow these steps:
1. Install mockery: `brew install mockery`
2. Run tests: `go test ./domain -v`
3. Make migrations: `make migrations`

For GraphQL functionality, refer to the [How To GraphQL guide](https://www.howtographql.com/graphql-go/1-getting-started/).

### Example GraphQL Queries

#### Register
```graphql
mutation {
  register(
    input: {
      email: "x@x.com"
      username: "xxxxxxxx"
      password: "xxxxxxxx"
      confirmPassword: "xxxxxxxx"
    }
  ) {
    accessToken
    user {
      id
      email
      username
      createdAt
    }
  }
}
```

#### Login
```graphql
mutation {
  login(input: {
    email: "x@x.com"
    password: "xxxxxxxx"
  }) {
    accessToken
    user {
      id
      email
    }
  }
}
```

Feel free to contribute, report issues, or provide feedback. Let's collaborate to enhance and optimize the XClone application!
