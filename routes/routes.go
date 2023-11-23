// handlers.go

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ferna/ApiGo/db"
	"github.com/ferna/ApiGo/models"


)

// Crear
func PostIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	err := json.NewDecoder(r.Body).Decode(&ice)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	createdIce := db.DB.Create(&ice)
	if createdIce.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(createdIce.Error.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ice)
}

// Leer todo
func GetAll(w http.ResponseWriter, r *http.Request) {
	var ices []models.IceCream
	db.DB.Find(&ices)

	json.NewEncoder(w).Encode(&ices)
}

// Leer
func GetIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&ice)
}

// Actualizar
func UpdateIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	var updatedIce models.IceCream

	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedIce)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	db.DB.Model(&ice).Updates(updatedIce)

	json.NewEncoder(w).Encode(&ice)
}

// Borrar
func DeleteIceCream(w http.ResponseWriter, r *http.Request) {
	var ice models.IceCream
	params := mux.Vars(r)
	db.DB.First(&ice, params["id"])

	if ice.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sabor de helado no encontrado"))
		return
	}

	db.DB.Unscoped().Delete(&ice)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ice)
}
