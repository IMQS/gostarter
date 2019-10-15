# META META META BEGIN
This is a "starter kit" Go project for a new service.

The idea is that you fork this repo into a new git repo, and make it your own. To find all the places where you need to change things, search throughout the project for the word "starter", and replace it appropriately.

This project has the following things setup for you, which you can change and extend:
* Load config from service, with HTTP port and DB connection info
* Documentation skeleton
* Dockerfile
* Migration system
* ORM models
* Unit tests that
    * Mock authentication
    * Wipe your Postgres database before every test
    * Connect to the docker Postgres image imqs/postgres:unittest
    * Make it easy to test JSON in/out

Delete this META section
# META META META END

# Go Starter
Describe what this service is, and why it exists

For documentation, see http://docs.imqs.co.za/#Services-Starter

# Building
```shell
go build github.com/IMQS/gostarter/starter
```

# Testing
```shell
docker run -p 5432:5432 imqs/postgres:unittest-10.5
go test github.com/IMQS/gostarter/starter
```