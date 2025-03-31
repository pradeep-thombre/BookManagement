Here’s a `README.md` file for your Book Management Golang application, following the format you provided:

```markdown
# Book Management Golang Application

## Overview

This is a Book Management application developed in Go. It provides a set of APIs for managing books in a MongoDB database. The application allows for CRUD (Create, Read, Update, Delete) operations on books and supports interacting with a MongoDB database locally.

## Run Locally

### Clone the project

```bash
git clone https://github.com/pradeep-thombre/BookManagement.git
```

### Go to the project directory

```bash
cd book-management
```

### Import dependencies

```bash
go mod tidy
```

### Start the server

```bash
go run .
```

The application will start running locally on port `3000`.

## API Reference

### Get all Books

```http
GET /books
```

Gets a list of all books and the total count of books present.

### Get Book by Id

```http
GET /books/${id}
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. ID of the book to fetch  |

Gets the book by the provided ID.

### Create a new Book

```http
POST /books
```

Payload:
```json
{
    "title": "string",        // required
    "author": "string",       // required
    "year": 2022              // required
}
```

Creates a new book with the provided payload and returns the ID of the created book.

### Update Book by Id

```http
PATCH /books/${id}
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. ID of the book to update |

Payload:
```json
{
    "title": "string",        // required
    "author": "string",       // required
    "year": 2022              // required
}
```

Updates the book details by the provided ID and payload.

### Delete Book by Id

```http
DELETE /books/${id}
```

| Parameter | Type     | Description                            |
| :-------- | :------- | :---------------------------------     |
| `id`      | `string` | **Required**. ID of the book to delete |

Deletes the book by the provided ID.

## Testing

You can test the application using the following command:

```bash
go test -v ./...
```

## Authors

- [Pradeep Thombre](https://www.github.com/Pradeep-Thombre)

## 🛠 Tech Stacks

- Golang
- MongoDB
- Gin Framework
- Ginkgo
- Gomega


## Support

For support, email us at [pradeepbthombre@gmail.com](mailto:pradeepbthombre@gmail.com)
```