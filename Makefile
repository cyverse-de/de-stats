install-swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: install-swagger
	GO111MODULE=on go mod vendor && GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
