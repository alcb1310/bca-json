# Budget Control Applicaton

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Authors](#authors)
- [FAQ](#faq)

### Configuration

To confiure the application create a `.env` file in the root directory of the application with the following content:

```.env
PORT=3000 # Port to listen on

PGDATABASE=<database name> # Database name
PGHOST=<database host> # Database host
PGUSER=<database user> # Database user
PGPASSWORD=<database password> # Database password
PGPORT=<database port> # Database port
```

### Build

To install this application run the following command:

```bash
make build
./bin/app
```

that will build and run the application.

## Overview

This is a simple budget control application that allows construction companys to add expenses and track their budget.

## Features

### Registration

To create an account you will need to register, for that you will need to make a post to the `/register` endpoint with the following content:

```json
{
    "ruc": "123456789",
    "name": "Company Name",
    "employees": 10,
    "email": "example@example.com",
    "password": "password123",
    "username": "Test User"
}
```

As a result of that you will get an authorized email and password for you to manage the budget information of the company which you registered with.

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
[Chi](https://go-chi.io/#/)
[Testify](https://github.com/stretchr/testify)

### Editor

![Neovim](https://img.shields.io/badge/NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)

### Version Control

![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

### CI/CD

![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

### Code information

|  | Code | Test |
| :--- | ---: | ---: |
| Lines of code | 402 | 463 |
| Number of files | 12 | 5 |
| Average | 33 | 92 |

## Authors

- Andrés Court [![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/alcb1310) [![X](https://img.shields.io/badge/X-%23000000.svg?style=for-the-badge&logo=X&logoColor=white)](https://x.com/alcb1310) [![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/andres-court/)

## FAQ

1. Can I have one account for multiple companies? No, each email is associated to only one company.
2. Can I have multiple accounts for one company? Yes, you can have multiple accounts for one company with different credentials.

