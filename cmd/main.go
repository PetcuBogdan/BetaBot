package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
)

func initializeRouter(){
mux := chi.NewRouter()

mux.Use(middleware.Recoverer)
mux.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://*", "http://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"Accepted", "Content-Type", "Authorization"},
    ExposedHeaders: []string{"Link"},
    AllowCredentials: true,
    MaxAge: 300,
}))

mux.Get("/cardId", CardId)
mux.Post("/cardId", CardId)
mux.Get("/cards", Cards)
mux.Post("/cards", Cards)
mux.Get("/cardAdd", AddCard)
mux.Post("/cardAdd", AddCard)
mux.Get("/cardDelete", DeleteCard)
mux.Post("/cardDelete", DeleteCard)
mux.Get("/cardEdit", EditCard)
mux.Post("/cardEdit", EditCard)

mux.Get("/profileId", ProfileId)
mux.Post("/profileId", ProfileId)
mux.Get("/profiles", Profiles)
mux.Post("/profiles", Profiles)
mux.Get("/profileAdd", AddProfile)
mux.Post("/profileAdd", AddProfile)
mux.Get("/profileDelete", DeleteProfile)
mux.Post("/profileDelete", DeleteProfile)
mux.Get("/profileEdit", EditProfile)
mux.Post("/profileEdit", EditProfile)


mux.Get("/taskId", TaskId)
mux.Post("/taskId", TaskId)
mux.Get("/tasks", Tasks)
mux.Post("/tasks", Tasks)
mux.Get("/taskAdd", AddTask)
mux.Post("/taskAdd", AddTask)
mux.Get("/taskDelete", DeleteTask)
mux.Post("/taskDelete", DeleteTask)
mux.Get("/taskStart", StartTask)
mux.Post("/taskStart", StartTask)
mux.Get("/taskStop", StopTask)
mux.Post("/taskStop", StopTask)

mux.Get("/getPalaceProducts",getPalaceProductsList)
mux.Post("/getPalaceProducts",getPalaceProductsList)
mux.Get("/getSupremeProducts",getSupremeProductsList)
mux.Post("/getSupremeProducts",getSupremeProductsList)

log.Fatal(http.ListenAndServe(":9000", mux))
}

func main() {
	category := []string{"CK1 Palace", "Jackets", "Shirting", "Tops", "Bottoms", "Tracksuits", "Hoods", "Sweatshirts", "T-Shirts", "Hats", "Footwear", "Bags", "Accessories", "Hardware"}
	for _,cat := range category{
	 	getProductsByCategory(cat)
	}
	initializeRouter()
}