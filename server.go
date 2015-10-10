package powerplug

import (
	"fmt"
	"net/http"

	"go-zoo/bone"
)

func init() {
	http.HandleFunc("/", handler)
}

func init() {
	mux := bone.New()

	// mux.Get, Post, etc ... takes http.Handler
	mux.Get("/home/:id", http.HandlerFunc(HomeHandler))
	mux.Get("/profil/:id/:var", http.HandlerFunc(ProfilHandler))
	mux.Post("/data", http.HandlerFunc(DataHandler))

	// Support REGEX Route params
	mux.Get("/index/#id^[0-9]$", http.HandleFunc(IndexHandler))

	// Handle take http.Handler
	mux.Handle("/", http.HandlerFunc(RootHandler))

	// GetFunc, PostFunc etc ... takes http.HandlerFunc
	mux.GetFunc("/test", Handler)

	http.ListenAndServe(":8080", mux)
}

func Handler(rw http.ResponseWriter, req *http.Request) {
	// Get the value of the "id" parameters.
	val := bone.GetValue(req, "id")

	rw.Write([]byte(val))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
