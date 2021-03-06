check-swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger swagger.yaml

generate-client: swagger
	swagger generate client -f swagger.yaml