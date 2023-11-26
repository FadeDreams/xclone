### to test the app
- brew install mockery
- go test ./domain -v
- make migrations

- https://www.howtographql.com/graphql-go/1-getting-started/
- go run github.com/99designs/gqlgen init
- 
- register
`
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

`
-login
`
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

`


-
