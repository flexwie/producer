package api

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"felixwie.com/producer/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/mux"
// )

// func createInventory(c *gin.Context) {
// 	db := models.GetDB()

// 	var data models.Inventory
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(rw, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	result := db.Create(&data)
// 	if result.Error != nil {
// 		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(rw, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.Write(json)

// }

// func getAllInventory(c *gin.Context) {
// 	db := models.GetDB()

// 	var data []models.Inventory
// 	result := db.Joins("LEFT JOIN refill_histories ON refill_histories.inventory_id = inventories.id").Find(&data)

// 	if result.Error != nil {
// 		http.Error(rw, result.Error.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	json, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(rw, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.Write(json)
// }

// func getOneInventory(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	db := models.GetDB()

// 	log.Println(vars)

// 	var data models.Inventory
// 	result := db.Joins("RefillHistory").Joins("Produce").Where("id = ?", vars["id"]).Find(&data)

// 	if result.Error != nil {
// 		http.Error(rw, result.Error.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	json, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(rw, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.Write(json)
// }

// func deleteInventory(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	log.Printf("vars: %v", id)

// 	db := models.GetDB()

// 	result := db.Where("id = ?", id).Delete(&models.Inventory{})

// 	if result.Error != nil {
// 		http.Error(rw, result.Error.Error(), http.StatusBadRequest)
// 		return
// 	}
// }
