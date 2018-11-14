package main

import (
	"github.com/jhendrixMSFT/azure-sdk-for-go3/services/compute/mgmt/2016-03-30/compute"
	"github.com/jhendrixMSFT/go-autorest3/autorest/azure/auth"
)

func main() {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	client := compute.NewVirtualMachinesClient("subID")
	client.Authorizer = a
}
