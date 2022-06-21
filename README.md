# bookrestapi

## How to run the application

1.  Clone the application with `https://github.com/emanuelquerty/bookrestapi.git`

2.  Use the postgres dump `books_db.sql` to create the database, create the tables and insert some dump data.

3.  Once the application is cloned and database is created, change the Connection String as per your postgres username, password and database name on line No. 12 in /pkg/db/database.go

4.  Before running the application, you need to import the pq (postgres) driver that the database/sql package requires. To do that, just run the following command in your terminal:

    ```shell
        go get "github.com/lib/pq"
    ```

5.  To run the application, please use the following command:

        ```shell
            go run main.go
        ```

    > Note: By default the server is listening on port **3000**.

## Endpoints Description

### Get All Books

```JSON
    URL - *http://localhost:3000/api/v1/books*
    Method - GET
```

### Get Book By ID

```JSON
    URL - *http://localhost:3000/api/v1/books/2*
    Method - GET
```

### Create a new Book Entry

```JSON
   URL - *http://localhost:3000/api/v1/books*
    Method - POST
    Body - (content-type = application/json)
    {
        "title": "Rennaissance",
        "published_year": 2022,
        "author_id": 2
    }
```

### Update Entry

```JSON
   URL - *http://localhost:3000/api/v1/books/2*
    Method - PUT
    Body - (content-type = application/json)
   {
        "title": "This is a new updated title",
        "published_year": 1840,
        "author_id": 2
    }
```

### Delete Entry

```JSON
     URL - *http://localhost:3000/api/v1/books/2*
    Method - DELETE
```

## Hope everything works. Thank you.
