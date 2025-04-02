# MIDDLEWARE WITH GOFIBER

Simple Rest API to implement middleware using Go-Fiber framework

## Features

### User

- Login
- Register

### Task

- Create new Task
- Get All Public Tasks
- Get All User's Tasks
- Get Task By Id

## Quick Start

Prepare your `.env` file and place it in root project. the example is provided on `.env.example`.
Then run this :

```cmd
> make all-dev
```

Base endpoint is:

```text
Host: localhost
Port: 3000
Prefix: /api/v1
```

## Building

### Prerequisite

- [Golang v1.24](https://go.dev/dl/)
- Postgresql

Environment Variables:

**NOTE:** Make sure you have Postgres database running and the details is set on `.env`

### Local

```cmd
> make clean
> make install
> make run
```

## Response

Responses are always the same between success and error response

### Success Response

| Key     | Type                |
|---------|---------------------|
| status  | `string`            |
| code    | `int`               |
| message | `string`            |
| data    | `object` `OPTIONAL` |

For Example:

```json
{
    "status": "success",
    "message": "login success",
    "code": 200,
    "data": {
        "token_type": "Bearer",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNzQzNjA5NTg5LCJpZCI6ImFjMmI2MTRjLTU3NDQtNDhmZS04ZmM0LWM4ZTZiNDQ1YWE5OCJ9.-2htsL5I_Pd6DRl0dN7DeiR1_G_HdHf5w4imvnfcdH8"
    }
}
```

### Error Response

| Key     | Type                |
|---------|---------------------|
| status  | `string`            |
| code    | `int`               |
| message | `string`            |
| data    | `object` `OPTIONAL` |

For Example:

```json
{
    "status": "error",
    "message": "BAD_REQUEST",
    "code": 400,
    "data": null
}
```

## Usage

All endpoint documentations can be viewed at [localhost:3000/api/v1/swagger/index.html](localhost:3000/api/v1/swagger/index.html) (Change the ip and port depending your setup on `.env`)

**NOTE:** Make sure your server is up and running.

### Users

Prefix: `/users`

Example: `localhost:3000/api/v1/users`

### Tasks

Prefix: `/tasks`

Example: `localhost:3000/api/v1/users`