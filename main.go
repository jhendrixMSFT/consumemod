package main

import (
	"github.com/jhendrixMSFT/azure-sdk-for-go2/services/compute/mgmt/2016-03-30/compute/v2"
	"github.com/jhendrixMSFT/go-autorest2/v11/autorest/azure/auth"
)

func main() {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	client := compute.NewVirtualMachinesClient("subID")
	client.Authorizer = a
	client.FakeV2Method()
}
