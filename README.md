# Data Integration API
> Fetch data from a Legacy API and serve a CRUD



## Instructions
- Install [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/);
- In your terminal, run ```docker-compose up```;
- Wait until all containers are ready;

You will know that the API will be ready to receive requests when you see something like this:

```
data-integration-container | 2021/03/17 22:15:28 Connected!
data-integration-container | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
data-integration-container | 
data-integration-container | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
data-integration-container |  - using env:	export GIN_MODE=release
data-integration-container |  - using code:	gin.SetMode(gin.ReleaseMode)
data-integration-container | 
data-integration-container | [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
data-integration-container | [GIN-debug] POST   /api/v1/auth/login        --> github.com/klasrak/data-integration/api/handlers.AuthHandler.Login-fm (3 handlers)
data-integration-container | [GIN-debug] POST   /api/v1/auth/logout       --> github.com/klasrak/data-integration/api/handlers.AuthHandler.Logout-fm (3 handlers)
data-integration-container | [GIN-debug] POST   /api/v1/auth/refresh-token --> github.com/klasrak/data-integration/api/handlers.AuthHandler.Refresh-fm (3 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/users/get         --> github.com/klasrak/data-integration/api/handlers.UserHandler.Find-fm (4 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/users/get/:email  --> github.com/klasrak/data-integration/api/handlers.UserHandler.FindByEmail-fm (4 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/negativations/fetch --> github.com/klasrak/data-integration/api/handlers.Handler.Fetch-fm (4 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/negativations/get --> github.com/klasrak/data-integration/api/handlers.Handler.GetAll-fm (4 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/negativations/get/:customerDocument --> github.com/klasrak/data-integration/api/handlers.Handler.Get-fm (4 handlers)
data-integration-container | [GIN-debug] GET    /api/v1/negativations/get-id/:id --> github.com/klasrak/data-integration/api/handlers.Handler.GetByID-fm (4 handlers)
data-integration-container | [GIN-debug] POST   /api/v1/negativations/create --> github.com/klasrak/data-integration/api/handlers.Handler.Create-fm (4 handlers)
data-integration-container | [GIN-debug] PUT    /api/v1/negativations/update/:id --> github.com/klasrak/data-integration/api/handlers.Handler.Update-fm (4 handlers)
data-integration-container | [GIN-debug] DELETE /api/v1/negativations/delete/:id --> github.com/klasrak/data-integration/api/handlers.Handler.Delete-fm (4 handlers)
data-integration-container | [GIN-debug] Listening and serving HTTP on :8080
```

## Docs
With the API running, you can access http://localhost:8080/swagger/index.html to see the documentation.

![Swagger](https://i.ibb.co/1Z2xgxm/swagger.png)

**If you use the [Insomnia REST Client](https://insomnia.rest/download), you can import the schema from [here](https://github.com/klasrak/data-integration/blob/master/config/Insomnia_2021-03-17).** 


## Getting started
All routes, with the exception of login, require the user to be authenticated, so first log in with the pre-registered admin user.

**Email**: admin@mail.com <br/>
**Password**: abcd1234

```sh
$ curl --location --request POST 'http://localhost:8080/api/v1/auth/login' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "email": "admin@mail.com",
    "password": "abcd1234"
  }'
```

Once logged in, the answer will be an access token valid for 1 hour and a refresh token.

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NVVUlEIjoiNjdjYjU3MjEtYTE5MC00NzVkLWEwNjgtMDEyMzhkMWE4OWI2IiwiZW1haWwiOiJhZG1pbkBtYWlsLmNvbSIsImV4cCI6MTYxNjAyNDM3OSwidXNlcklEIjoiNjA1MjZiOWMwZTEzNzIzMjEyMzEwYjlkIn0.8cL6Lk20l8Iw8CNsIHNirKleO8jUMr4JAHx7O9pa90s",
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQG1haWwuY29tIiwiZXhwIjoxNjE2NjI1NTc5LCJyZWZyZXNoVVVJRCI6IjY3Y2I1NzIxLWExOTAtNDc1ZC1hMDY4LTAxMjM4ZDFhODliNisrNjA1MjZiOWMwZTEzNzIzMjEyMzEwYjlkIiwidXNlcklEIjoiNjA1MjZiOWMwZTEzNzIzMjEyMzEwYjlkIn0.esm9rdSdQ1Q-b50VCmSekLdeGPHmEMggaVKsHp4EyPM"
}
```

To integrate the data from the Legacy API, you must call the endpoint `/negativations/fetch`. **Remember to send the Authorization header with the Bearer Token**

```sh
$ curl --location --request GET 'http://localhost:8080/api/v1/negativations/fetch' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NVVUlEIjoiMGNjOTc3NGUtNGYwZi00ODFmLThlZGItYTgzM2E0MWVlOWI3IiwiZW1haWwiOiJhZG1pbkBtYWlsLmNvbSIsImV4cCI6MTYxNjAyNjEzOCwidXNlcklEIjoiNjA1MjZiOWMwZTEzNzIzMjEyMzEwYjlkIn0.plv7yZSNgmt5ofvnAyuBo3j8BMgjlABWAYc6AwxmxcY'
```

The legacy API will be called, and the data returned will be saved to a MongoDB instance only the first time you call. 

```json
[
    {
        "companyDocument": "59291534000167",
        "companyName": "ABC S.A.",
        "customerDocument": "51537476467",
        "value": 1235.23,
        "contract": "bc063153-fb9e-4334-9a6c-0d069a42065b",
        "debtDate": "2015-11-13T20:32:51-03:00",
        "inclusionDate": "2020-11-13T20:32:51-03:00"
    },
    {
        "companyDocument": "77723018000146",
        "companyName": "123 S.A.",
        "customerDocument": "51537476467",
        "value": 400,
        "contract": "5f206825-3cfe-412f-8302-cc1b24a179b0",
        "debtDate": "2015-10-12T20:32:51-03:00",
        "inclusionDate": "2020-10-12T20:32:51-03:00"
    },
    {
        "companyDocument": "04843574000182",
        "companyName": "DBZ S.A.",
        "customerDocument": "26658236674",
        "value": 59.99,
        "contract": "3132f136-3889-4efb-bf92-e1efbb3fe15e",
        "debtDate": "2015-09-11T20:32:51-03:00",
        "inclusionDate": "2020-09-11T20:32:51-03:00"
    },
    {
        "companyDocument": "23993551000107",
        "companyName": "XPTO S.A.",
        "customerDocument": "62824334010",
        "value": 230.5,
        "contract": "8b441dbb-3bb4-4fc9-9b46-bdaad00a7a98",
        "debtDate": "2015-08-10T20:32:51-03:00",
        "inclusionDate": "2020-08-10T20:32:51-03:00"
    },
    {
        "companyDocument": "70170935000100",
        "companyName": "ASD S.A.",
        "customerDocument": "25124543043",
        "value": 10340.67,
        "contract": "d6628a0e-d4dd-4f14-8591-2ddc7f1bbeff",
        "debtDate": "2015-07-09T20:32:51-03:00",
        "inclusionDate": "2020-07-09T20:32:51-03:00"
    }
]
```

If you call the endpoint again and no new data is returned to be inserted, an error will be returned.

```json
{
    "statusCode": 400,
    "message": "legacy API data already fetched and there is no new data to save"
}
```

From here, you can perform any CRUD operation. Take a look at the [documentation](http://localhost:8080/swagger/index.html) for more information.

## Test

There are mocks for the [repository layer](https://github.com/klasrak/data-integration/tree/master/mocks), and tests implemented on the internal and handler layer.

To run the tests, simply run the command `go test ./...` at the root of the project.

## TODO list
- **WRITE MORE TESTS!!!!!**
- Add a cache layer
- Revoge active tokens

## Purpose of the project

This project was (and will continue to be) developed for the purpose of studies. Any feedback or contribution, positive or not, will be very welcome and will be taken into consideration.

## License

This project follows [MIT License](https://github.com/klasrak/data-integration/blob/master/LICENSE).
