package user

import (
    "context"
    "time"
    "github.com/Tardieujose/Stock-Managment-Back/pkg/database"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Password string             `bson:"password"`
    Role     string             `bson:"role"`
}

func CreateUser(user User) (*mongo.InsertOneResult, error) {
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    return collection.InsertOne(ctx, user)
}

func FindUserByUsername(username string) (*User, error) {
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user User
    err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}
