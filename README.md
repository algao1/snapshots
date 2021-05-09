# Image Repo

This project was created for the `Fall 2021 - Shopify
Developer Intern Challenge Question`.

## Usage

### Requirements

Docker or go1.16+ is required.

Additionally, a `.env` file is required containing all the API keys. It should be in the format

```
# MongoDB
MONGO_URI = 
MONGO_DB = 
MONGO_ACCS = 
MONGO_IMGS =

# DigitalOcean Spaces
SPACES_KEY = 
SPACES_SECRET = 
SPACES_ENDPOINT = 
SPACES_REGION = 
SPACES_BUCKET = 

# Redis
CACHE_URL=localhost
CACHE_PORT=6379
CACHE_PASS=
CACHE_DB=0
```

### With Docker and Docker Compose

The image repository server can be started with

```console
docker-compose up
```

and shut down with

```console
docker-compose down
```

alternatively, the Redis cache and server can be started individually

```console
docker run -dp 6379:6379 redis

docker build -t shopify-server -f Dockerfile.server .

docker run --rm -it -p 10000:10000 --env CACHE_URL=host.docker.internal shopify-server -server_addr :10000
```

The image repository client can be built and started with

```console
docker build -t shopify-client -f Dockerfile.client .

docker run -it --rm --network="host" -v ${PWD}:/root/mount shopify-client
```

### With Go 1.16+

A Redis cache will still be required on localhost:6379, but the server and client can be started with

```console
go run cmd/server/server.go
```

and

```console
go run cmd/client/client.go
```

### Using the Client

There are currently 5 commands

```
reg [username] [password] - registers username and password

login [username] [password] - logs in using username and password

up [0|1] [regex] [directories...] - uploads file with public (0) or private (1) access in listed directories matching regex

down [id] [directory] - downloads the file with id to specified directory

ls [-n] - lists all viewable images, 'ls -n' will view the next page
```

Note: when using the client with Docker, all directories must be prefixed by mount/ .