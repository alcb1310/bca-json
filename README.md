# Budget Control Application

## Tech Stack

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)

## Editor

[![Neovim](https://img.shields.io/badge/NeoVim-%2357A143.svg?&style=for-the-badge&logo=neovim&logoColor=white)](https://neovim.io/)

## Socials

[![X](https://img.shields.io/badge/X-%23000000.svg?style=for-the-badge&logo=X&logoColor=white)](https://x.com/alcb1310)
[![GitHub](https://img.shields.io/badge/GitHub-%23121011.svg?style=for-the-badge&logo=GitHub&logoColor=white)](https://github.com/alcb1310)
[![built with Codeium](https://codeium.com/badges/main)](https://codeium.com/profile/alcb1310)

## Installation

## Environment Variables

| ***Variable*** | ***Type*** | ***Description*** |
| :--- | :--- | :--- |
| **PORT** | String | Port the application listens on |
| **DB_HOST** | String | Database host address |
| **DB_PORT** | String | Database port |
| **DB_USER** | String | Database user name |
| **DB_PASSWORD** | String | Database user password |
| **DB_NAME** | String | Database name |
| **JWT_SECRET** | String | JWT secret key |

## Routes

### Unauthenticated Routes

| ***Route*** | ***Method*** | ***Description*** |
| :--- | :--- | :--- |
| **/api/v2/companies* | POST | Register a new company |
| **/api/v2/login** | POST | Login a user |

### Authenticated Routes

| ***Route*** | ***Method*** | ***Description*** |
| :--- | :--- | :--- |
| **/api/vw2/bca/users** | GET | Get all users |
| **/api/vw2/bca/users** | POST | Create a new user |
| **/api/vw2/bca/users/me** | GET | Get the current user |
| **/api/vw2/bca/users/{id}** | GET | Get a user |
| **/api/vw2/bca/users/{id}** | PUT | Update a user |
| **/api/vw2/bca/users/{id}** | DELETE | Delete a user |
