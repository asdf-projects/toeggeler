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

Change variables in _.env_ if you wish to re-name the database file or use another port for the server.
Defaults (both are strings): 
* Port: **:8080**
* Database file: **./toeggeler.sqlite**


# Using the API

<details>

<summary>User Management</summary>

<br />

**GET /api/users** to get all users

```
// Response 
[
    {
        "id": 1,
        "username": "Franz",
        "mail": "franz@net.com",   
    }
]
```
* **GET /api/users/{name}** to get a single user
```
// Response 
{
    "id": 1,
    "username": "Franz",
    "mail": "franz@net.com",   
}
```
* **POST /api/users** to create a new user
```
// Request
{
    "username": "Franz",
    "mail": "franz@net.com"
    "password": "insecure"
}

// Response 
{
    "id": 1,
    "username": "Franz",
    "mail": "franz@net.com",   
}
```
* **PUT /api/users/{name}** to update an existing user
```
// Request
{
    // only mail updatable for now
    "mail": "franz2@net.com"
}

// Response 
{
    "id": 1,
    "username": "Franz",
    "mail": "franz2@net.com",   
}
```
* **DELETE /api/users/{name}** to delete an existing user
</details>

<details>
<summary>Game Management</summary>

<br/>

* **POST /api/games** to submit a completed game
```
// Request
[
    // "GAME_START" requires "team1" and "team2" properties 
    {
        "timestamp": 1000000, // unix timestamp
        "event": "GAME_START",
        "team1": {
            "offense": 1,
            "defense": 2
        },
        "team2: {
            "offense": 3,
            "defense": 4
        }
    },
    // any type of goal event requires the player id 
    {
        "timestamp": 100000,
        "event": "GOAL" | "OWN_GOAL" | "FOETELI",
        "player": 1
    }
    {
        "timestamp": 130000,
        "event": "GAME_END"
    }
]

// Response
{
    "id": "IdIdIdIdId"
}
```
</details>