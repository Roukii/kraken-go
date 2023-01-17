# kraken-go
A go version of the kraken api generated from OpenApi doc

# Warning
Use this at your own risk, kraken's openapi file is not generated perfectly, there is quite a few modification to make to the file at each generation.
I know there is nothing really optimized yet, it's work in progress.

# OpenApi
To update the openapi configuration first download the openapi code generator at : 
https://github.com/deepmap/oapi-codegen

To get the latest openapi file for kraken go to : 
https://docs.kraken.com/rest/


To generate the kraken.gen.go use the command:
`oapi-codegen -response-type-suffix Response -package openapi openapi/openapi.json > openapi/kraken.gen.go`