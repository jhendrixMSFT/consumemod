package main

import (
	"fmt"

	"github.com/jhendrixMSFT/go-fake-sdk/services/bar/2018-01-01/bar"
	oldfoo "github.com/jhendrixMSFT/go-fake-sdk/services/foo/2018-01-01/foo"
	newfoo "github.com/jhendrixMSFT/go-fake-sdk/services/foo/2018-02-01/foo"
)

func main() {
	fmt.Println(bar.Report())
	fmt.Println(oldfoo.Report())
	fmt.Println(newfoo.Report())
}
