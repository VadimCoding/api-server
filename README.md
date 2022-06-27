# api-server
A small api-server in Go

I'm exposing 2 endpoints: 
- /ready
- /{param}

## /ready
Method: GET
Returns: {"status": "ok"}

## /{param}
Method: GET
Description: The transformation of the "param" pass in argument of the request from the funciton word2Number, which transform a word/sentence in its number equivalent ("abc" -> 123, "aba" -> 121, "Dad" -> 414)
Returns: {"result": STRING}
