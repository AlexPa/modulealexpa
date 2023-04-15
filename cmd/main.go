package main

import (
	"fmt"
	"golang.org/x/text/language"
	"rsc.io/sampler"
)

func main() {

	fmt.Println(sampler.Hello(language.Make("cn-CN")))
}
