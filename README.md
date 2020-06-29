# TSA Coding Challenge API

![build status](https://github.com/graciasrochelle/TSACodingChallengeAPI/workflows/Go/badge.svg?branch=master)

## Aim of the project

A simple GoLang RESTful API that accepts and persists contact information using in-memory or AWS RDS SQL database storage. The api has two functionalities - Get list of contacts and Create a new contact.

## Usage

### Get list of Contacts

#### Request

`GET /contacts`

```
Responses:
    200: ContactsResponse - array of contacts
	400: BadRequestError
	500: InteralServiceError
	502: BadGatewayError
```

### Create a new Contact

#### Request

`POST /contact`

```
Request Body:
    {
        "email": "fredrik_idestam@test.com",
        "fullName": "fredrik IDESTAM",
        "phoneNumbers": [
            "+6139888998"
        ]
    }
```

```
Responses:
    201: ContactResponse - empty response
    400: ValidationError
    500: InternalServerError
    502: BadGateway
```

## Components and Design

![Design](/images/diagram.png)

The api is designed using factory pattern in Go-Lang. It also uses SQL relational database to store and retrieve contact details. A [SQL Driver](github.com/denisenkom/go-mssqldb) was used.

A fall back mechanism is added if this database is not available to store and read contacts from an file that is read at run time.

This application is deployed on AWS EC2 instance.

### Possible extension

1. _GitHub actions_ is integrated to run tests. The deployment of this project to AWS EC2 instance is manual and can be automated using github actions.

2. _OAuth_ can be added to provide Server-Client Authentication

## Installation and Usage

### Configuration

Configuration is done using Json file. Add in the connection string save data into SQL relational database. Click to view [Database Schema](sql-schema.sql)

### Run API via Docker

`docker build -t yourusername/tsacodingchallengeapi .`

`docker run -d -p 10010:10010 yourusername/tsacodingchallengeapi`

### Deploy API to AWS EC2 instance

#### Clone Project

1. Setup EC2 instance
2. Connect to AWS EC2 via SSH
3. Install Golang on EC2 instance
4. Set Go environment variables
5. Clone the Go Application

#### Run application via Docker

1. Install docker on EC2 instance

2. Build Project Dockerfile

   `docker build -t tsacodingchallenge .`

3. Run Project Dockerfile

   `docker run --publish 10010:10010 --detach --name tda tsacodingchallenge:latest`

## References

1. [Deploy to AWS EC2 instance](https://hackernoon.com/deploying-a-go-application-on-aws-ec2-76390c09c2c5)

2. [Golang SQL Databases](https://flaviocopes.com/golang-sql-database/)

3. [Stackover](https://stackoverflow.com/)

4. Open Source Community
