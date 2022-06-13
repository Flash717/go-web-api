# Web Service + Database in Golang

## Prerequisites
Golang 1.18+

### Environment setup
This project uses `.env` files to specify values for your application
* DB_USER: Database user name
* DB_PASS: Database user password
* DB_ADDR: Database address
* DB_NAME: Database name
* API_PORT: port on which the API shall launch

### Database
This project is using MySQL as database. 

Table `album`:
| Column | Type               |
| ------ | ------------------ |
| id     | int auto_increment |
| artist | varchar            |
| title  | varchar            |
| price  | decimal            |