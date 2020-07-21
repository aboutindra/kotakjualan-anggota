package main

import (
	"github.com/gofiber/fiber"
	"kotakjualan-anggota/router"
)

const port = 8000

var r router.Router

func init() {
	r = router.Router{}
}

func main() {
	app := fiber.New()
	app.Get("v1/api/anggota/:id?", r.GetById)
	app.Post("v1/api/anggota/create", r.CreateAnggota)
	app.Put("v1/api/anggota/update/:id", r.UpdateAnggota)

	app.Listen(port)
}
