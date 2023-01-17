# kraken-go
A go version of the kraken api generated from OpenApi doc


# OpenApi
To update the openapi configuration first download the openapi code generator at : 
https://github.com/deepmap/oapi-codegen

To get the latest openapi file for kraken go to : 
https://docs.kraken.com/rest/


To generate the kraken.gen.go use the command:
`oapi-codegen -response-type-suffix Response -package openapi openapi/openapi.json > openapi/kraken.gen.go`