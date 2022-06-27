# api-server
A small api-server in Go

## What is it ?
I'm exposing 2 endpoints: 
- /ready
- /{param}

### /ready
Method: GET
Returns: {"status": "ok"}

### /{param}
Method: GET
Description: The transformation of the "param" pass in argument of the request from the funciton word2Number, which transform a word/sentence in its number equivalent ("abc" -> 123, "aba" -> 121, "Dad" -> 414)
Returns: {"result": STRING}

## How to use it ?
1. Pull the code
2. Go in the /app folder `cd app`
3. Launch the binary `go run main.go` or build the binary first `go build -o api-server` and then run it `./api-server`
4. Once you'll see `Listening on port 4200...` you can test the endpoint with your local browser with `http://localhost:4200/ready`
