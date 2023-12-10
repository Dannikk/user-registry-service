# User registry service

This service application was developed to demonstrate skills in working with Golang, PostgreSQL and Redis.
The application is written according to __clean architecture__ and __SOLID__ principles.

# Configuration

Сreate `.env` file and fill it according to the [example](./.env_example).

You need to configure the following values:

- POSTGRES_USER
- POSTGRES_PW
- POSTGRES_DB (can be default value)
- PGADMIN_MAIL
- PGADMIN_PW

# Deploy with docker compose

``` shell
$ docker compose up
```

Stop the containers with
``` shell
$ docker compose down
# To delete all data run:
$ docker compose down -v
```

# Use the pgAdmin 
By default the pgAdmin web interface will be available at port 5050 (http://localhost:5050). But you can change the pgAdmin port in [docker-compose.yml](./docker-compose.yml)

After logging in with your credentials of the .env file, you can add your database to pgAdmin. 
1. Right-click "Servers" in the top-left corner and select "Create" -> "Server..."
2. Name your connection
3. Change to the "Connection" tab and add the connection details:
- Hostname: "postgres" (this would normally be your IP address of the postgres database - however, docker can resolve this container ip by its name)
- Port: "5432"
- Maintenance Database: $POSTGRES_DB (see .env)
- Username: $POSTGRES_USER (see .env)
- Password: $POSTGRES_PW (see .env)
