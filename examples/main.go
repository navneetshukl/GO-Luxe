package main

import "github.com/navneetshukl/luxe"

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
	l.SendJSON(200, luxe.H{
		"name": "Navneet",
	})
}

