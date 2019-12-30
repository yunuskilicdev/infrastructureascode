.PHONY: clean build


clean:
	rm -rf ./bin/cacheClient/cacheClient


build:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o bin/cacheClient/cacheClient ./CacheClient
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o bin/dbClient/dbClient ./DbClient
