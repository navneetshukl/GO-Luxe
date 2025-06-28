package main

import (
	"fmt"

	luxe "github.com/navneetshukl/GO-Luxe"
)

func main() {
	luxxer := luxe.New()

	luxxer.GET("/a", GETPath)
	luxxer.GET("/b", func(l *luxe.LTX) {
		l.SendJSON(200, luxe.H{
			"message": "This is get method",
		})
	})

	luxxer.Run()
}

func GETPath(l *luxe.LTX) {
	fmt.Println("Path is ",l.Request.Query)
	fmt.Println("AllQuery is ",l.Request.Params)
	fmt.Println("Title is ",l.GetParam("title"))
	l.SendJSON(200, luxe.H{
		"name": "Navneet",
	})
}
