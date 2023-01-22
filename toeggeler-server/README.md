# Running the server

## Pre-requisites
Go to https://go.dev/doc/install to learn how to setup Go

## Running
Go to the folder containing the _main.go_ file and run the following command:

`go run .`

or

`go build . && ./toeggeler-server` 

The modules defined in _go.mod_ should automatically be loaded.

## Customizing

Change variables in _config.toml_ if you wish to re-name the database file or use another port for the server.
Defaults:

* **common.dev:** true
<br>_Enables the development mode, which basically means more logs_

* **server.port:** ":8080"
<br>_Set the server port_

* **server.enableJwt:** false
<br>_Enables mandatory JWT Token validation for authorized routes_

* **database.file:** "./toeggeler.sqlite"
<br>_Filename for the SQLite database_


# Using the API
Open http://localhost:8080/api/swagger/index.html to get an overview of the available API endpoints and try them out.

When setting **server.enableJwt** to true, some APIs require a valid JWT token. This is indicated by a response with http code 401.

Use _POST /api/authenticate_ to get a JWT Token. Use the button **Authorize** in the Swagger UI (top right) and enter the token received from the authentication response.
Currently the token is hard-coded to be valid for half an hour.