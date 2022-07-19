# Simple Terraform provider

Run services:

```sh
cd services
go run services
```

Build and install provider:

```sh
cd simple-provider
go build -o terraform-provider-simpleprovider
mv terraform-provider-simpleprovider ~/.terraform.d/plugins/hashicorp.com/edu/simpleprovider/0.1/linux_amd64
```

Try example:

```sh
cd examples
terraform init
terraform apply
```
