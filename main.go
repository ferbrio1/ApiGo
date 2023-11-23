// main.go
package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ferna/ApiGo/db"
	"github.com/ferna/ApiGo/models"
	"github.com/ferna/ApiGo/routes"

)


func main() {
	// Establece la conexión a la base de datos
	db.DBConnection()

	// Realiza las migraciones de la base de datos
	db.DB.AutoMigrate(&models.IceCream{})

	// Configura las rutas
	r := mux.NewRouter()
	r.HandleFunc("/icecream", routes.PostIceCream).Methods("POST")
	r.HandleFunc("/icecream", routes.GetAll).Methods("GET")
	r.HandleFunc("/icecream/{id}", routes.GetIceCream).Methods("GET")
	r.HandleFunc("/icecream/{id}", routes.UpdateIceCream).Methods("PUT")
	r.HandleFunc("/icecream/{id}", routes.DeleteIceCream).Methods("DELETE")

	// Inicia el servidor HTTP en el puerto 3000
	http.ListenAndServe(":3000", r)
}
