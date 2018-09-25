package main

import (
	"fmt"

	//"github.com/jhendrixMSFT/moremods/bar/2018-01-01/bar"
	//"github.com/jhendrixMSFT/modtest/foo"
	//oldfoo "github.com/jhendrixMSFT/modtest/foo/2018-01-01/foo"
	//newfoo "github.com/jhendrixMSFT/modtest/foo/2018-02-01/foo"
	"github.com/jhendrixMSFT/moremods/profiles/latest/bar"
	"github.com/jhendrixMSFT/moremods/profiles/latest/foo"
)

func main() {
	//fmt.Println(bar.Report())
	//fmt.Println(oldfoo.Report())
	//fmt.Println(newfoo.Report())
	fmt.Println(bar.Report())
	fmt.Println(foo.Report())
}
