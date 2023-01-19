package main

//go:generate mkdir -p api
//go:generate oapi-codegen -package=api -generate=server,spec,types,strict-server -import-mapping=models.yaml:oapi-codegen-subtype/models -o=api/api.go doc/api.yaml

//go:generate mkdir -p models
//go:generate oapi-codegen -package=models -generate=spec,types,skip-prune -o=models/models.go doc/models.yaml
