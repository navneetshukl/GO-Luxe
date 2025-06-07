package main

import "github.com/navneetshukl/luxe"

func main() {
	luxxer := luxe.New()

	luxxer.GET("/a",GET)

	luxxer.Run()
}

func GET(l *luxe.LTX){
	jsonResp:=`{
		"name":"Navneet",
	}`
	l.SendJSON(200,jsonResp)
}
