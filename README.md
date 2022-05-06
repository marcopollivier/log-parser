# Log Parser

Log parser tool

## Environment 

### Running in a environment with Go
> Go 1.18 is needed

## Configurations and Properties

### Logging 

To enable logging, change the configuration property `LoggingEnable` on 
[JSON file `config/config.json`](config/config.json)

Example: 

```json
{
  "LoggingEnable": true
}
```

## Run application

#### Executing via code
```bash
$ go run cmd/server/main.go
``` 

#### On Linux 
```bash
$ GOOS=linux GOARCH=amd64 go build -o log-parser cmd/server/main.go
$ ./log-parser
```

#### On MacOS 
```bash
$ GOOS=darwin GOARCH=amd64 go build -o log-parser cmd/server/main.go
$ ./log-parser
```

### Running in a Docker environment
- Build a Docker image
```bash
$ docker build -t log-parser .
```

- Create and Run via Docker Container
```bash
$ docker run -it log-parser
```

- Create and Run via Docker Compose
```bash
$ docker-compose run log-parser
```

