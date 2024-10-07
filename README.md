# Fiber CRUD REST API Demo

## cURL requests

### Books

- Get all Books

  ```bash
  curl --location 'localhost:8080/api/books'
  ```

- Get a book by ID

  ```bash
  curl --location 'localhost:8080/api/books/1'
  ```

- Create a book

  ```bash
  curl --location 'localhost:8080/api/books' \
  --header 'Content-Type: application/json' \
  --data '{
      "title": "Sample Book",
      "author": "John Doe",
      "description": "A sample book"
    }'
  ```

- Update a book

  ```bash
  curl --location --request PUT 'localhost:8080/api/books' \
  --header 'Content-Type: application/json' \
  --data '{
      "id": 1
      "title": "Sample Book",
      "author": "John Doe",
      "description": "A sample book"
    }'
  ```
