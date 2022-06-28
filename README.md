<!-- PROJECT LOGO -->
<br />
<p align="center">

<h3 align="center">Final Project Golang Training by Hacktiv8</h3>

  <p align="center">
    a MyGram App Golang REST API
    <br />
    <br />
    <a href="https://github.com/HafidhIrsyad/final-project-golang-hacktiv8/issues">Report Bug</a>
    ·
    <a href="https://github.com/HafidhIrsyad/final-project-golang-hacktiv8/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
    * [Built With](#built-with)
* [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
* [Contributing](#contributing)
* [Contact](#contact)


<!-- ABOUT THE PROJECT -->
## About The Project

### Standard Naming Convention
* FOLDER = camelCase
* FILE = snake_case
* FUNCTION = PascalCase

### Feature

* Register & Login
* CRUD User
* CRUD Social Media
* CRUD Comment
* CRUD Photo

### Built With

* [Go as Programming Language](https://golang.org/)
* [JWT Auth as Authentication](https://github.com/dgrijalva/jwt-go)
* [PostgreSQL as Database](https://www.postgresql.org/)
* [Gorilla Mux as HTTP Router](https://github.com/gorilla/mux)

### Usage
* [Postman Collections](https://www.getpostman.com/collections/45ad78085ee86fbcdf50)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running, follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Install Golang, PostgreSQL, and Postman for testing
* create an `.env` file

```bash
export SECRET_KEY="S3CR3T"
export DB_PORT="5432"
export DB_NAME=""
export DB_ADDRESS="localhost"
export DB_USERNAME=""
export DB_PASSWORD=""
```

### Installation

1. Clone the repo (in Folder htdocs)
```sh
git clone git@github.com:HafidhIrsyad/final-project-golang-hacktiv8.git
```
2. Install module with get
```sh
go get
go mod tidy
```
3. Run
```sh
source .env
go run main.go
```
4. Access via url
```JS
localhost:port
```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- CONTACT -->
## Contact
[![](https://img.shields.io/badge/LinkedIn_Hafidh-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/hafidhirsyad/)

