package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

var cyberchefFS embed.FS

func main() {
	cyberchefSubFS, err := fs.Sub(cyberchefFS, "cyberchef")
	if err != nil {
		log.Fatal(err)
	}

	setupRoutes(cyberchefSubFS)

	fmt.Println("🚀 Master server running on https://localhost:8443")
	fmt.Println("🔧 CyberChef at: https://localhost:8443/#/cyberchef")
	fmt.Println("📋 Status at: https://localhost:8443/#/status")

	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}

func setupRoutes(cyberchefFS fs.FS) {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/cyberchef/", http.StripPrefix("/cyberchef/", http.FileServer(http.FS(cyberchefFS))))
}
