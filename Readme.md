

GoLang Task with MoP - EB
======================

REST API that will allow users to register/login, after which it will be possible to use all other endpoints.

It provides Postman config file that you can load into your Postman and test all possible endpoints.

Frontend for this backend is a simple website with a small shop feature that will allow users, after login, to purchase simple items.
Shipment is handled in another part of the system and is not part of this API.

This api mostly handles purchases, with basic functionality.

## Table of content

- [Prerequirements](#prerequirements)
    - [GO & MySQL](#go) or
    - [Docker](#docker)
- [Running the app](#running)
- [Testing](#testing)

## Prerequirements

- GO & MySql
- Docker & Docker Compose

## Data Migration

This app uses [jinzhu/gorm](http://github.com/jinzhu/gorm) as ORM library.

All data migrations and sample data seedings are handled through [api/seed/seeder.go](api/seed/seeder.go) and referenced at [api/server.go](api/server.go)


## Running the app

In project directory:
```
go run main.go
```

Running with Docker, in project directory:
```
docker-compose build
docker-compose up
```

## Testing

Currently, we don't have any unit tests.

To test manually with Postman ( TODO: Postman test runner ),
open Postman, an "Import" the file [MOP - EB.postman_collection.json](MOP%20-%20EB.postman_collection.json) from the root of the app folder.

Then this is the recommended testing order:

 - Get Users
 - Get Products
 - Get User ( Sample user show )
 - Register
 - Login ( We use registered user from above, not from seeder. Copy the token in your clipboard)
 - Get All Purchases ( Paste the copied token as Bearer Token. Do the same for all next requests. )
 - Get Own Purchases ( Should be empty for the new user )
 - Create Purchase ( Purchase for a product and currently logged in user is made )
 - Get Own Purchases ( Should be 1 record the new user )
 - Dashboard - Public ( displays witgets with no userID asigned and marked as public )
 - Dashboard - Private

