default: build

build:
	go build -o terraform-provider-example
	mv terraform-provider-example ~/.terraform.d/plugins/hashicorp.com/edu/example/0.1/darwin_amd64/

api:
	go run api/main.go