package stock

import (
    "encoding/json"
    "net/http"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/models"
)

func GetStockHandler(w http.ResponseWriter, r *http.Request) {
    stock, err := GetStock()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(stock)
}

func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
    var stock models.Stock
    json.NewDecoder(r.Body).Decode(&stock)
    UpdateStock(stock)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(stock)
}