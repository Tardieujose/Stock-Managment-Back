package order

import (
    "context"
    "log"
    "time"
    "fmt"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/models"
    "github.com/Tardieujose/Stock-Managment-Back/internal/product"
    "go.mongodb.org/mongo-driver/bson"
)

// GetAllOrders obtiene todas las órdenes de la base de datos
func GetAllOrders() ([]models.Order, error) {
    var orders []models.Order

    collection := database.GetCollection("orders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var order models.Order
        if err := cursor.Decode(&order); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return orders, nil
}

func CreateOrderWithProductUpdates(order models.Order) error {
    // Obtener el producto
    prod, err := product.GetProductByID(order.ProductID)
    if err != nil {
        return err
    }

    // Verificar si hay suficiente stock
    if prod.Quantity < float64(order.Quantity) {
        return fmt.Errorf("not enough stock for product %s", prod.ProductID)
    }

    // Calcular el total de la orden
    order.Total = float64(order.Quantity) * prod.Price

    // Iniciar sesión de base de datos para realizar las operaciones como una transacción
    collection := database.GetCollection("orders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Crear la orden en la base de datos
    _, err = collection.InsertOne(ctx, order)
    if err != nil {
        return err
    }

    // Actualizar el stock del producto
    prod.Quantity -= float64(order.Quantity)
    err = product.UpdateProduct(prod)
    if err != nil {
        return err
    }

    log.Println("Orden creada y stock actualizado exitosamente:", order.ID)
    return nil
}

// CreateOrder crea una nueva orden en la base de datos
// (opcional: eliminar si no se necesita)
func CreateOrder(order models.Order) error {
    collection := database.GetCollection("orders")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, order)
    if err != nil {
        return err
    }

    log.Println("Orden creada exitosamente:", order.ID)
    return nil
}
