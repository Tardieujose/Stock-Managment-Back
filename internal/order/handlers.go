package order

import (
	"encoding/json"
	"net/http"

	"github.com/Tardieujose/Stock-Managment-Back/pkg/models"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
    orders, err := GetAllOrders()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(orders)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)
    CreateOrder(order)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(order)
}