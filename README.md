# Go postgres docker example

See also python version: https://github.com/KarmaComputing/db-linking-postgres-container-app-python

- docker --link to container hostname
- type cast of port number to int

# Build database
```
docker build -t db --file DockerfilePostgres .
```

# Start database
```
docker run --rm --name db db
```
# Build app
```
docker build -t app .
```
# Run app
```
cp .env.example .env
docker run --link db --env-file .env app
```

# Observe:

```
Connected OK
```

