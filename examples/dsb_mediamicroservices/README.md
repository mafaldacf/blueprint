# DeathStarBench MediaMicroservices

## Getting started

Prerequisites for this tutorial:
* [thrift compiler](https://thrift.apache.org/download) is installed
* docker is installed

## Compiling the application

To compile the application, we execute `wiring/main.go` and specify which wiring spec to compile. To view options and list wiring specs, run:

```zsh
go run wiring/main.go -h
```

If you encounter errors like because of missing modules that are suposed to be replaced by local ones, do:

```zsh
export GOFLAGS=-mod=mod
export GOWORK=off
```
OR
```zsh
cd wiring
go clean -cache -modcache
export GOFLAGS=-mod=mod
export GOWORK=off
go mod tidy
cd ..
```

The following will compile the `docker` wiring spec to the directory `build`. This will fail if the pre-requisite thrift compiler is not installed.

```zsh
rm -rf build
go run wiring/main.go -w docker -o build
```

## Running the application

To run the application, navigate to `build/docker` and run `docker compose up`. Use flag `--build` to build images if code is changed.

```zsh
# remove dangling images (mostly untagged and used by others)
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

docker-compose --env-file build/.env -f build/docker/docker-compose.yml up --build
```  

If you see Docker complain about missing environment variables, edit the `.env` file in `build/docker` and put the following:

```zsh
API_SERVICE_HTTP_BIND_ADDR=12345
MOVIEID_DB_BIND_ADDR=12346
MOVIEID_DB_DIAL_ADDR=12346
MOVIEID_SERVICE_THRIFT_BIND_ADDR=12347
MOVIEID_SERVICE_THRIFT_DIAL_ADDR=12347
MOVIEINFO_DB_BIND_ADDR=12348
MOVIEINFO_DB_DIAL_ADDR=12348
MOVIEINFO_SERVICE_THRIFT_BIND_ADDR=12349
MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR=12349
UNIQUEID_SERVICE_THRIFT_BIND_ADDR=12350
UNIQUEID_SERVICE_THRIFT_DIAL_ADDR=12350
```

## Testing API Requests

```go
type APIService interface {
	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
	ComposeUniqueId(ctx context.Context, reqID int64) (int64, error)
}
```

```zsh
# MovieIdService -> RegisterMovieId()
curl http://localhost:12345/RegisterMovieId?reqID=0\&movieID=mymovieid0\&title=mymovietitle0
curl http://localhost:12345/RegisterMovieId?reqID=1\&movieID=mymovieid1\&title=mymovietitle1
```

```zsh
# MovieInfoService -> WriteMovieInfo()
curl http://localhost:12345/WriteMovieInfo?reqID=0\&movieID=mymovieid0\&title=mymovietitle0\&casts=mycastslist0
curl http://localhost:12345/WriteMovieInfo?reqID=1\&movieID=mymovieid1\&title=mymovietitle1\&casts=mycastslist1
```

```zsh
# UniqueIdService -> ComposeUniqueId()
curl http://localhost:12345/ComposeUniqueId?reqID=0
curl http://localhost:12345/ComposeUniqueId?reqID=1
```
