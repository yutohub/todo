# TODO
Create an API to control a simple TODO list using Go's standard library.

## HTTP Method
Follows REST architecture style constraints.  
| HTTP Method | Query String | Request Body | Path | 
| --- | --- | --- | --- |
| GET | ◎ |   |   |
| POST |   | ◎ |   |
| PUT |   | ◎ |   |
| DELETE |   |   | ◎ |

## Example
You can check using Postman.  
> Postman is an API platform for developers to design, build, test and iterate their APIs.  
- Start the server
    ```
    $ go run main.go
    ```
- GET
    ```
    GET /todo/ HTTP/1.1
    Host: localhost:8080
    ```
- POST
    ```
    POST /todo/ HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Content-Length: 36

    {
        "task": "お洗濯",
        "done": 0
    }
    ```
- PUT
    ```
    PUT /todo/ HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Content-Length: 49

    {
        "id": 1,
        "task": "お掃除",
        "done": 1
    }
    ```
- DELETE
    ```
    DELETE /todo/1 HTTP/1.1
    Host: localhost:8080
    ```
## Specification

- REST API
- Only standard library
- CRUD Operation
- Connection to Database
- MVC Model Based
- Path routing
- Exchange in JSON

## Other
I created this for my own Golang study.
If anyone has any advice, feel free to share it with me.