# Go Fiber GORM Microservice Boilerplate

Under Construction 

Tech Stacks
```
1. Go Programming Language
2. MySQL
3. Go Fiber Framework
4. Gorm ORM
5. JWT (Json Web Token)
6. Docker
7. Swagger
```

Login API
```
curl --location --request POST 'http://localhost:8080/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"yourmail@gmail.com",
    "password":"yourpassword"
}'
```

INSERT USER
```
INSERT INTO `users` VALUES (1, 'yourusername', 'youremail@mail.com', 'firstname', 'lastname', 1, '$2a$10$UKIkKmrZZ7H1gnzwyUkJ2.XK2u2IpqplinOXwbKqLia5xl9dcBVRW', '2023-06-11 05:12:42', '2023-06-11 05:12:42');
```
```
$2a$10$UKIkKmrZZ7H1gnzwyUkJ2.XK2u2IpqplinOXwbKqLia5xl9dcBVRW = Password
```

Refresh Token
```
curl --location --request POST 'http://localhost:8080/v1/auth/access-token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "refreshToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTY4ODgzMDYzM30.UL-y7LzpuLq3mQMORkYEACZnTXG5qTCz_hP8UuQlK2M"
}'
````

Medicine 
GET ALL with Pagination 
```
curl --location --request GET 'http://localhost:8080/v1/medicine?limit=10&page=1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidHlwZSI6ImFjY2VzcyIsImV4cCI6MTY4ODg1ODU3N30.j_YDKmaQGxBae1OilFSyr65MlFfgP3Zq-rjEBA4W8nk' \
--header 'Content-Type: application/json' \
```

Running Docker
1. Open Aplication Root Folder in the terminal.
2. Run Docker Compose Build
```
docker-compose build
```
3. Run Docker Compose Create
```
docker-compose create
```
4. Run Docker Compose Start
```
docker-compose start
```

Delete Container
```
docker-compose down
```

SWAGGER
```
Coming Soon
```

