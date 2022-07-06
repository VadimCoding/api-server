# Api-Server
A small api-server in Go

## What it is ?
Simply 2 API endpoints: 
- /ready
- /{param}

### /ready
**Method**: GET

**Returns**: {"status": "ok"}

### /{param}
**Method**: GET

**Description**: The result of `param` pass in argument the funciton word2Number, which transform a word/sentence in its number equivalent ("abc" -> 123, "aba" -> 121, "Dad" -> 414)

Returns: {"result": STRING}

## How to use it ?
### Source code
1. Pull the code
2. Go in the /app folder `cd app`
3. Launch the binary `go run main.go` or build the binary first `go build -o api-server` and then run it `./api-server`
4. Once you'll see `Listening on port 4200...` you can test the endpoint with your local browser with `http://localhost:4200/ready`

### From docker image
You can choose to directly use the docker image and run it in a container.
1. Get the image from `docker pull vadimdocker/api-server2`
2. Run it with `docker run vadimdocker/api-server2` 
3. Depending on your hosting configuration for docker, you might need to expose ports.

## CI/CD
The CI is builind and pushing a new image when a new version of the code is produce on main with a tag. The git tag will be use for the image version on dockerhub.
