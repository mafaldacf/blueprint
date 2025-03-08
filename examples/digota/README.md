# Digota

## Getting started

Prerequisites for this tutorial:
* [thrift compiler](https://thrift.apache.org/download) is installed
* docker is installed

## Compiling the application

To compile the application, we execute `wiring/main.go` and specify which wiring spec to compile. To view options and list wiring specs, run:

```
go run wiring/main.go -h
```

If you encounter errors like because of missing modules that are suposed to be replaced by local ones, do:

```
export GOFLAGS=-mod=mod
export GOWORK=off
```
OR
```
cd wiring
go clean -cache -modcache
export GOFLAGS=-mod=mod
export GOWORK=off
go mod tidy
cd ..
```

The following will compile the `docker` wiring spec to the directory `build`. This will fail if the pre-requisite thrift compiler is not installed.

```
rm -rf build
go run wiring/main.go -w docker -o build
```

## Running the application

To run the application, navigate to `build/docker` and run `docker compose up`. Use flag `--build` to build images if code is changed.

```
# remove dangling images (mostly untagged and used by others)
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

docker-compose --env-file build/.env -f build/docker/docker-compose.yml up --build
```  

If you see Docker complain about missing environment variables, edit the `.env` file in `build/docker` and put the following:

```
FRONTEND_HTTP_BIND_ADDR=12345
FRONTEND_HTTP_DIAL_ADDR=12345
NOTIFICATION_QUEUE_BIND_ADDR=12346
NOTIFICATION_QUEUE_DIAL_ADDR=12346
NOTIFY_SERVICE_SERVICE_THRIFT_BIND_ADDR=12347
NOTIFY_SERVICE_SERVICE_THRIFT_DIAL_ADDR=12347
POST_DB_BIND_ADDR=12348
POST_DB_DIAL_ADDR=12348
STORAGE_SERVICE_SERVICE_THRIFT_BIND_ADDR=12349
STORAGE_SERVICE_SERVICE_THRIFT_DIAL_ADDR=storage_service_service:12349
```

## TEST FRONTEND

```
# UploadPost (text) -> postId
curl http://localhost:12348/UploadPost?text=helloworld\&username=mafalda
curl http://localhost:12348/FetchPost?postID=123
```
