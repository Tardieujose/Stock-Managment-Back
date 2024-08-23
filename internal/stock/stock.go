package stock

import (
    "context"
    "log"
    "time"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/models"
    "go.mongodb.org/mongo-driver/bson"
    // "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// GetStock obtiene la lista de stock de todos los productos
func GetStock() ([]models.Stock, error) {
    var stockList []models.Stock

    collection := database.GetCollection("stock")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var stock models.Stock
        if err := cursor.Decode(&stock); err != nil {
            return nil, err
        }
        stockList = append(stockList, stock)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return stockList, nil
}

// UpdateStock actualiza el stock de un producto en la base de datos
func UpdateStock(stock models.Stock) error {
    collection := database.GetCollection("stock")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"product_id": stock.ProductID}
    update := bson.M{"$set": bson.M{"quantity": stock.Quantity}}

    _, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
    if err != nil {
        return err
    }

    log.Println("Stock actualizado para el producto ID:", stock.ProductID)
    return nil
}