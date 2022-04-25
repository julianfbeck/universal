package main

import "github.com/julianfbeck/universal/backend/services/stock/route"

func main() {
	engine := route.Create()
	engine.Listen(":8080")
}
