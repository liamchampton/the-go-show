package main

////// UNCOMMENT TO USE THE STANDARD LIBRARY //////
// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprintf(w, "Hello, World!")
// }

////////////////////////////////////////////////////////////////

////// UNCOMMENT TO USE GIN //////
// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", handler)
// 	r.Run(":8080")
// }

// func handler(c *gin.Context) {
// 	c.String(200, "Hello, World!")
// }

////////////////////////////////////////////////////////////////

////// UNCOMMENT TO USE FIBER //////
// import (
// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!")
// })

// 	app.Listen(":8080")
// }

////////////////////////////////////////////////////////////////

////// UNCOMMENT TO USE THE STANDARD LIBRARY WITH HTML (index.html) //////
// import (
// 	"fmt"
// 	"os"

// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// Read the contents of the HTML file
// 	html, err := os.ReadFile("index.html")
// 	if err != nil {
// 		http.Error(w, "Unable to load HTML file", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the content type header to HTML
// 	w.Header().Set("Content-Type", "text/html")

// 	// Write the HTML file to the response
// 	fmt.Fprintf(w, "%s", html)
// }
