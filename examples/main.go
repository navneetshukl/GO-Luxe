package main

import "github.com/navneetshukl/luxe"

func main() {
	luxxer := luxe.New()

	luxxer.GET("/a",GETPath)

	luxxer.Run()
}

func GETPath(l *luxe.LTX){
	jsonResp:=`{
		"name":"Navneet",
	}`
	l.SendJSON(200,jsonResp)
}
