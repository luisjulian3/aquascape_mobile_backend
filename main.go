package main

import service "github.com/luisjulian3/aquascape_mobile_backend/services"

func main() {
	// expose echo http server
	service.EchoHTTPService()
}
