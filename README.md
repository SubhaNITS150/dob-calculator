
# DOB-Calculator

A backend application which fetches user DOB and calculates age in years via api call using Go Programming Language. 


## Installation

Clone the git repository using the following steps:- 

```bash
  git clone https://github.com/SubhaNITS150/dob-calculator.git
  cd dob-calculator
  go mod tidy

```
Note:- Usually we do not push the .env file to the github. Here I have uploaded it on github to run the project. During deployment, we upload it on environment variables section.
## Starting the server

After cloning the project, here are the next steps:- 


Start the server

```bash
  go run cmd/server/main.go
```


## Frontend

To run the frontend, open the index.html file after running the server. It is not required to run the index.go file.

## Run Locally 

Run the terminal client on a separate terminal

```bash
  go run index.go
```


## Server Running
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765973944/eadea5b8-3358-44e7-8363-bc54d2225dba_ms85ie.jpg)

## Running index.go
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765974182/0eef41e3-508c-41a7-9adf-f3792d132a94_rovxzp.jpg)
## Running tests
1. View all Users
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765974402/4ba0f980-ce1f-4c5a-be2e-034494a92183_qxhtyq.jpg)

2. Create User
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765974599/7f0b4663-6341-4870-a970-82f6012b9714_s3kuye.jpg)

3. Update User
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765974719/78c49ae8-0396-4667-a0e7-3493d749cff5_avl6nh.jpg)

4. Delete User
![](https://res.cloudinary.com/dludtk5vz/image/upload/v1765974781/862f3cb9-2ee5-4476-a52c-65a24dea9c2c_xbwums.jpg)
## API Reference

#### Checking health of API

```http
  GET /health
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `_` | `_` | Checks if the server is running |

#### Get a user with id
```http
  GET /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `_`      | `_` | Fetches the user with calculated age |

#### Get all items

```http
  GET /users
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `_`      | `_` | Fetches all users with calculated age |

#### Creates a new user and returns the user details.

```http
  POST /users
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**:- Name of the user |
| `dob` | `string` | **Required**:- Date of Birth `(YYYY-MM-DD)`

#### Updates an existing user’s details and recalculates age.

```http
  PUT /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**:- UUID of the user |
| `name`      | `string` | **Required**:- Updated name |
| `dob` | `string` | **Required**:- Updated Date of Birth `(YYYY-MM-DD)`

#### Deletes the specified user.

```http
    DELETE /users/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**:- UUID of the user |






## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DATABASE_URL = postgresql://postgres.xtmmwtimsjupkzunfjrx:Subha%402910@aws-1-ap-southeast-2.pooler.supabase.com:5432/postgres?sslmode=require`

`PORT = 8090`


## Features

- Frontend Included (using HTML, CSS, JS)
- Live previews
- User profile 


## Authors

- [@SubhaNITS150](https://www.github.com/SubhaNITS150)

