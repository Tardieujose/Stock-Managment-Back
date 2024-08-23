package product

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/models"
)

// Handler para obtener todos los productos
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
    pageQuery := r.URL.Query().Get("page")
    limitQuery := r.URL.Query().Get("limit")

    // Si no se proporcionan parámetros, devuelve todos los productos
    if pageQuery == "" && limitQuery == "" {
        products, err := GetAllProducts()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        response := map[string]interface{}{
            "products": products,
            "total":    len(products),
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        return
    }

    page, err := strconv.Atoi(pageQuery)
    if err != nil || page < 1 {
        page = 1
    }

    limit, err := strconv.Atoi(limitQuery)
    if err != nil || limit < 1 {
        limit = 10
    }

    products, total, err := GetAllProductsWithPag(page, limit)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]interface{}{
        "products": products,
        "total":    total,
        "page":     page,
        "limit":    limit,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Handler para crear un nuevo producto
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    json.NewDecoder(r.Body).Decode(&product)
    err := CreateProduct(product)  // Llamada directa a la función CreateProduct
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}



