package main

import "github.com/akinozgen/go-hello/models"

func main() {
	var akin models.Insan = models.NewInsan(3, "atkafasi", 22)

	akin.SetName("Akın Özgen").SetAge(24).SetId(1)

	println(akin.GetName())
}
