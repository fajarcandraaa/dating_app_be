# dating_app_be

<!-- ABOUT THE PROJECT -->
## About The Project

This is a simple dating app exercise. 
For the exercise, I have to build backend system for Dating Mobile App.
The basic functions I have reached are :
1. Sign Up
2. Login
3. View other dating profiles with several condition like :
 - If users is not in premium package, they only have 10 quota in  a day
 - If users is premium package, they have take 1 of 2 choices premium features like : No swipe quota or verified label

## Code Structure 

1. `config`, the directory to collecting all main configuration. From database connection to API routers
2. `handler`, this is layer for handler function, the function that will be called in router
3. `helpers`, is for reusable function
4. `internal/dto`, or data transfer object, used to collect all function that used for convert data
5. `internal/entity`, this is layer that represent our database
6. `internal/presentation`, contains about the struct that will be present like request and response format
7. `internal/reposiroty`, is layer to communication into database
8. `internal/usecase`, is business logic
9. `middleware`, contain authentication config
10. `mocking`, contain mock function and seeder data
11. `pkg`, the directory that will be used to setup all third pary library/packages


### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)
* [PostgreSQL](https://www.postgresql.org/)
* [Redis](https://redis.io)
* [Linter](https://github.com/golangci/golangci-lint)
* [Medis](https://getmedis.com/)

<!-- GETTING STARTED -->
## Getting Started
Before we get started, it's important to know that  before you run this code you have to make sure that `Redis` is already exist and ready to run on your device. Than this code use a custom command to execute it with makefile to make more simple command like :
1. make update
2. make tidy
3. make start

So, let start it.
1. After clone this repository, just run `make update`.
2. Setup your `.env` file such as database connection and redis connection based on default setting on you device.
3. To make sure that all dependency is run well, than run `make tidy`.
4. Finally, you can run your project with command: `make start`.
5. Go to postman and set url like `http://localhost:8080/`, for information that port to run this project depend on configuratin on `.env`

And for additional information, i'm alredy setup unit-testing, just run `make test-service`.

## Afterword
Hopefully, it can be easily understood and useful. Thank you~


