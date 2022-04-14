# Movies API POC

This project should be used to serve as a POC for the Go Masterclass internal project developed at Wizeline

Based on the efforts made on this project, improvements and enhancements can be added to provide a better experience for the potential students of the Mastrerclass

## Requirements

The following software is required to be able to run this project, make sure to have all of this installed before proceeding

- [Git](https://git-scm.com/downloads)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- [Go](https://go.dev/dl/)
- [Air](https://github.com/cosmtrek/air)

## Setting up the environment

1. Clone the repo to your machine

2. Run `docker-compose up` on your terminal to create the localstack container, this will setup a local dynamodb instance that will be accessible on `localhost:4566`. This can take a few minutes if it's your first time running it, since docker needs to pull the localstack image to your machine
    - Take a look at the `docker-compose.yml` to learn more about it 

3. Create the movies table by running the following command un your terminal:
```
aws dynamodb --endpoint-url=http://localhost:4566 create-table \
    --table-name movies \
    --region us-east-1 \
    --attribute-definitions \
        AttributeName=id,AttributeType=S \
    --key-schema \
        AttributeName=id,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5
```

3. Run `go mod tidy` on your terminal to download and install the project dependencies

4. Duplicate the `config_dist.json` into a new file called `config.json` and fill it with the appropriate values

5. Run `air` on the root of the project on your terminal to start the API. If everything goes well you should see the following log `starting API server on port X`
    - If you modify source code files while air is running the API will restart automatically

You're done! After this you should be able to make requests to the API

If you want to **stop** the localstack container just run `docker-compose stop`, to bring it back up run `docker-compose restart`

If you want to **remove** the container run `docker-compose down`, be careful since this will delete the instance, the table and the data that could be stored already