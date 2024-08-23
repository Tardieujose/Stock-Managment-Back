package product


import (
    // "encoding/json"
    // "net/http"
    // "strconv"
    "fmt"
    "context"
    "log"
    "time"
    // "github.com/gorilla/mux"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllProducts obtiene todos los productos de la base de datos
func GetAllProducts() ([]models.Product, error) {
    var products []models.Product

    collection := database.GetCollection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var product models.Product
        if err := cursor.Decode(&product); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

func GetAllProductsWithPag(page, limit int) ([]models.Product, int64, error) {
    var products []models.Product

    collection := database.GetCollection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Cálculo del número de documentos a omitir
    skip := (page - 1) * limit

    // Opción para establecer límite y omisión
    findOptions := options.Find()
    findOptions.SetLimit(int64(limit))
    findOptions.SetSkip(int64(skip))

    cursor, err := collection.Find(ctx, bson.M{}, findOptions)
    if err != nil {
        return nil, 0, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var product models.Product
        if err := cursor.Decode(&product); err != nil {
            return nil, 0, err
        }
        products = append(products, product)
    }

    if err := cursor.Err(); err != nil {
        return nil, 0, err
    }

    // Obtener el total de productos
    total, err := collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return nil, 0, err
    }

    return products, total, nil
}

// GetProductByID obtiene un producto por su ID
func GetProductByID(productID string) (models.Product, error) {
    var product models.Product

    collection := database.GetCollection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Buscar el producto por ID
    filter := bson.M{"product_id": productID}
    err := collection.FindOne(ctx, filter).Decode(&product)
    if err != nil {
        return product, err
    }

    return product, nil
}

// UpdateProduct actualiza el producto en la base de datos
func UpdateProduct(product models.Product) error {
    collection := database.GetCollection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Filtrar por ProductID
    filter := bson.M{"product_id": product.ProductID}
    update := bson.M{
        "$set": bson.M{
            "name":        product.Name,
            "price":       product.Price,
            "quantity":    product.Quantity,
            "description": product.Description,
        },
    }

    _, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    log.Println("Producto actualizado exitosamente:", product.ProductID)
    return nil
}

const counterCollection = "counters"

// GetNextSequenceValue obtiene el siguiente valor de secuencia para el contador
func GetNextSequenceValue(sequenceName string) (int, error) {
    collection := database.GetCollection(counterCollection)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"_id": sequenceName}
    update := bson.M{"$inc": bson.M{"sequence_value": 1}}
    options := options.FindOneAndUpdate().SetReturnDocument(options.After)

    var result struct {
        SequenceValue int `bson:"sequence_value"`
    }
    err := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            // Si el contador no existe, crear uno nuevo con el valor inicial
            _, err = collection.InsertOne(ctx, bson.M{
                "_id":            sequenceName,
                "sequence_value": 1,
            })
            if err != nil {
                return 0, err
            }
            return 1, nil
        }
        return 0, err
    }

    return result.SequenceValue, nil
}

// CreateProduct crea un nuevo producto con un ID único
func CreateProduct(product models.Product) error {
    productID, err := GetNextSequenceValue("productid")
    if err != nil {
        return err
    }

    // Asigna el nuevo productID al producto
    product.ProductID = fmt.Sprintf("%d", productID)

    collection := database.GetCollection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, product)
    if err != nil {
        return err
    }

    log.Println("Producto creado exitosamente:", product.ProductID)
    return nil
}

