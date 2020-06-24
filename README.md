# TSA Coding Challenge API 
![Go](https://github.com/graciasrochelle/TSACodingChallengeAPI/workflows/Go/badge.svg?branch=master)

TSACodingChallengeAPI is an GoLang RESTful API that accepts and persists data for a "Contact" into a SQL relational database.

# TSACodingChallengeAPI API

## Get list of Contacts

### Request

`GET /contacts`

## Create a new Contact

### Request

`POST /contact`

# Run API via Docker

`docker build -t yourusername/tsacodingchallengeapi .`

`docker run -d -p 10010:10010 yourusername/tsacodingchallengeapi`