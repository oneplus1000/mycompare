package main

import (
	"log"

	"github.com/oneplus1000/mycompare/libmycompare"
)

func main() {
	var my = libmycompare.MyCompare{
		SrcConfig: libmycompare.ConnConfig{
			ConnString: "",
		},
		DestConfig: libmycompare.ConnConfig{
			ConnString: "",
		},
	}
	err := my.Run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
