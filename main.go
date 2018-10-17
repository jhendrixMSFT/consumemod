package main

import (
	"context"
	"fmt"

	"github.com/jhendrixMSFT/azure-sdk-for-go2/profiles/2017-03-09/resources/mgmt/resources"
	"github.com/jhendrixMSFT/go-autorest2/autorest/azure/auth"
)

func main() {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	client := resources.NewClient("subID")
	client.Authorizer = a
	for page, err := client.List(context.Background(), "", "", nil); page.NotDone(); err = page.Next() {
		if err != nil {
			panic(err)
		}
		for _, v := range page.Values() {
			fmt.Println(*v.ID)
		}
	}
}
