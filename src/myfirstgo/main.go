package main

import (
	"myfirstgo/modules/myhttpserver"
	"myfirstgo/modules/mylibs"
)

func main() {
	println(mylibs.Hello())
	println("main run")
	myhttpserver.CreateHttpServer("8090")
}

//fmt.Println(fmt.Sprint("Hello ", 23))
