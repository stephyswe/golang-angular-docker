Angular and Golang: A Rapid Guide - Advanced

Run on WSL Ubuntu - ang-go

```erlang
# VSC - Go
# Extention - Go

Fix 
go get -u golang.org/x/tools/gopls
```
1. Introduction
```
...
```

2. Setup
    
    ```erlang
    # Windows Install 
    https://go.dev/dl/ - Microsoft Windows
    
    # WSL - Ubuntu Install
    https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374
    
    wget https://dl.google.com/go/go1.16.1.linux-amd64.tar.gz
    sudo tar -xvf go1.16.1.linux-amd64.tar.gz
    sudo mv go /usr/local
    
    cd ~
    explorer.exe .
    
    # Edit .bashrc
    """
    ...
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
    """
    
    go version
    
    # Go
    mkdir ang-go && cd ang-go
    mkdir go-ambassador && cd go-ambassador
    go mod init ambassador # create go.mod
    go get github.com/gofiber/fiber/v2
    go run main.go
    
    # http://localhost:3000/
    ```
        
    
3. Docker
    
    ```erlang
    sudo docker-compose up --build
    # http://localhost:8000
    ```
        
4. Database
    
    ```erlang
    sudo docker-compose up --build # twice
    # http://localhost:8000
    ```
        
5. Migrations
    
    ```erlang
    # HeidiSQL
    https://www.heidisql.com/download.php
    
    # Open to MYSQL
    """
    HOST: localhost
    USER/PW: root
    PORT: 33066
    DATABASE: ambassador
    """
    
    # Add table
    
    > Files
    
    sudo docker-compose up --build
    ```
        
6. Live Reloading
    
    ```erlang
    # Live Reload - https://thedevelopercafe.com/articles/live-reload-in-go-with-air-4eff64b7a642
    # Github - https://github.com/cosmtrek/air
    
    sudo docker-compose up --build
    ```

7. Admin Authentication Endpoints
    ```erlang
    # Endpoint - api/admin/

    # POST - register
    # POST - login
    # GET- user
    # POST - logout
    # PUT - /users/info
    # PUT - /users/password
    ```

8. Routes
    
    ```erlang
    # Postman
    
    POST - http://localhost:8000/api/admin/register
    """
    {
        "message": "hello"
    }
    """
    ```
        
9. Register
    
    ```erlang
    # GO Package - Bcrypt
    https://pkg.go.dev/
    go get golang.org/x/crypto/bcrypt@v0.0.0-20210513164829-c07d793c2f9a
    
    # Validate - Postman
    POST - http://localhost:8000/api/admin/register
    {
        "first_name": "a",
        "last_name": "a",
        "email": "a@a.com",
        "password": "a",
        "password_confirm": "a2"
    }
    ...
    {
        "message": "passwords do not match"
    }
    
    {
        "first_name": "a",
        "last_name": "a",
        "email": "a@a.com",
        "password": "a",
        "password_confirm": "a"
    }
    ...
    {
        "Id": 1,
        "FirstName": "a",
        "LastName": "a",
        "Email": "a@a.com",
        "Password": "JDJhJDEyJHNpLm5ncTFEMkRlcnEuWlRBN1dTWk9BeXhYeHpodzFqc3dVWTVTVmYyM3FoS0NQcWlySkF1",
        "IsAmbassador": false
    }
    ```
        
10. Login
    ```erlang
    # Validate - Postman
    POST - http://localhost:8000/api/admin/login
    
    # Wrong password
    {
        "email": "a@a.com",
        "password": "b",
    }
    ...
    {
        "message": "Invalid credentials"
    }
    
    # Wrong email
    {
        "email": "a@a1.com",
        "password": "a",
    }
    ...
    {
        "message": "Invalid credentials"
    }
    
    # Right
    {
        "email": "a@a.com",
        "password": "a",
    }
    ...
    {
        "Id": 1,
        "FirstName": "a",
        "LastName": "a",
        "Email": "a@a.com",
        "Password": "JDJhJDEyJHNpLm5ncTFEMkRlcnEuWlRBN1dTWk9BeXhYeHpodzFqc3dVWTVTVmYyM3FoS0NQcWlySkF1",
        "IsAmbassador": false
    }
    ```
    
11. Jwt
    
    ```erlang
    POST - http://localhost:8000/api/admin/login
    
    # Token returned
    {
        "email": "a@a.com",
        "password": "a"
    }
    ...
    {  "message": "success" } - COOKIE
    ```
    
12. Methods (SetPassword, ComparePassword)
    
    ```erlang
    # Validate - Postman
    POST - http://localhost:8000/api/admin/register
    
    # Right
    {
        "first_name": "a2",
        "last_name": "a2",
        "email": "a@a2.com",
        "password": "a",
        "password_confirm": "a"
    }
    ...
    { "user": "userInfo..." }
    ```
        
13. Authenticated User
    
    ```erlang
    GET - http://localhost:8000/api/admin/user
    
    # Token returned
    {
        "email": "a@a.com",
        "password": "a"
    }
    ...
    {  "message": "success" } - COOKIE
    ```
        
14. Formatting Output - (user.go)
    
    ```erlang
    # GET - http://localhost:8000/api/admin/user
    
    {
        "id": 4,
        "first_name": "sa",
        "last_name": "a",
        "email": "a@a.com"
    }
    ```
15. Logout
    
    ```erlang
    POST - http://localhost:8000/api/admin/logout
    {  "message": "success" } - COOKIE
    ```
    
16. Middlewares
    
    ```erlang
    # All Routes - Recheck
    
    > authController
    > auth
    > routes
    ```
    
17. Profile
    
    ```erlang
    
    # 1. Update UserInfo
    
    # PUT - /admin/users/info
    """
    {
        "first_name": "b",
        "last_name": "b",
        "email": "b@b.com"
    }
    ...
    {
        "id": 4,
        "first_name": "bab",
        "last_name": "ba",
        "email": "baba@b.com"
    }
    """
    
    # GET - /admin/user
    """
    {
        "id": 4,
        "first_name": "bab",
        "last_name": "ba",
        "email": "baba@b.com"
    }
    """
    
    # 2. Update Password
    # PUT - /admin/users/password
    """
    {
        "password": "b",
        "password_confirm": "b"
    }
    ...
    {
        "id": 4,
        "first_name": "",
        "last_name": "",
        "email": ""
    }
    """
    
    # POST - /admin/login
    """
    {
        "password": "b",
        "password_confirm": "b"
    }
    ...
    { "message": "success" }
    
    ```

18. Admin Endpoints
    
    ```erlang
    # Endpoint - api/admin/
    
    # GET/POST - products
    # GET/PUT/DELETE - products/:id
    # GET - users/:id/links
    # GET - orders
    # GET - ambassadors
    ```

19. Ambassadors
    
    ```erlang
    go get github.com/bxcodec/faker/v3@v3.6.0
    # Error - Run Command PopulateUser
    go run src/commands/populateUsers.go - # Fails DB error
    ...
    # Solution - Run PopulateUser inside Docker
    docker-compose exec backend sh
    # go run src/commands/populateUsers.go
    
    # /admin/ambassadors - #Shows 30 entities
    
    ```
    
20. Products
    
    ```erlang
    # Endpoints
    # Get Products - GET - /admin/products 
    # Get Product - GET - /admin/products/1
    # Create Products - POST - /admin/products
    # Update Product - PUT - admin/products/1
    # Delete Product - DEL - admin/products/1
    
    # Get Products - # []
    
    # Create Products
    """
    {
        "title": "",
        "description": "b",
        "image": "img",
        "price": 10
    }
    ...
    {
        "id": 1,
        "title": "title",
        "description": "desc",
        "image": "img",
        "price": 10
    }
    """
    # Get Product
    """
    {
    		"id": 1,
        "title": "",
        "description": "b",
        "image": "img",
        "price": 10
    }
    """
    
    # Update Product
    """
    {
        "title": "title",
        "description": "desc 2",
        "image": "img 2",
        "price": 20
    }
    ...
    {
        "id": 1,
        "title": "title3",
        "description": "",
        "image": "",
        "price": 0
    }
    """
    
    # Delete Product - # []
    
     
    
    ```
    
21. Embedded Structs (model ID, runProducts)
    
    ```erlang
    #  Run PopulateProducts inside Docker
    docker-compose exec backend sh
    # go run src/commands/populateProducts.go
    
    # /admin/products - #Shows 30 entities
    ```
    
22. Links
    
    ```erlang
    # /admin/users/1/links - # []
    ```
    
23. Orders
    
    ```erlang
    #  Run populateOrders inside Docker
    docker-compose exec backend sh
    # go run src/commands/populateOrders.go
    
    # GET - /orders - #Shows 30 entities - order_item is null
    """
    {
            "id": 1,
            "transaction_id": "",
            "user_id": 17,
            "code": "fsQpDHO",
            "ambassador_email": "HjDyGrm@GomdfJF.com",
            "name": "",
            "email": "oLUVnCK@mlWSsuk.net",
            "address": "",
            "city": "",
            "country": "",
            "zip": "",
            "order_item": null
        },
    """
    ```
    
24. Preloading
    
    ```erlang
    # GET - /orders - #Shows 30 entities - order_item too
    """
    {
            "id": 1,
            "transaction_id": "",
            "user_id": 17,
            "code": "fsQpDHO",
            "ambassador_email": "HjDyGrm@GomdfJF.com",
            "name": "",
            "email": "oLUVnCK@mlWSsuk.net",
            "address": "",
            "city": "",
            "country": "",
            "zip": "",
            "order_items": [
                {
                    "id": 1,
                    "order_id": 1,
                    "product_title": "dolorum",
                    "price": 60,
                    "quantity": 2,
                    "admin_revenue": 108,
                    "ambassador_revenue": 12
                },
                {
                    "id": 2,
                    "order_id": 1,
                    "product_title": "debitis",
                    "price": 32,
                    "quantity": 0,
                    "admin_revenue": 0,
                    "ambassador_revenue": 0
                },
                {
                    "id": 3,
                    "order_id": 1,
                    "product_title": "aliquam",
                    "price": 92,
                    "quantity": 3,
                    "admin_revenue": 248.39999999999998,
                    "ambassador_revenue": 27.6
                }
            ]
        },
    """
    
    # Methods FulLname & GetTotal
    # GET - /orders
    """
    {
            "id": 1,
            ...
            "name": "Murphy Upton",
    				...
            "total": 396,
            "order_items": [...]
        },
    """
    
    # Method Orders in Links
    # GET - /users/10/links - # []
    ```
    
25. Ambassador Authentication Endpoints
    
    ```erlang
    # Endpoint - api/ambassadors/
    
    # POST - register
    # POST - login
    # GET- user
    # POST - logout
    # PUT - /users/info
    # PUT - /users/password
    ```
    
26. Multiple Routes
    
    ```erlang
    # POST - register
    """
    {
        "first_name": "ab",
        "last_name": "ab",
        "email": "d@d.com",
        "password": "a",
        "password_confirm": "a"
    }
    ...
    {
        "id": 42,
        "first_name": "ab",
        "last_name": "ab",
        "email": "d@d.com"
    }
    """
    
    ```