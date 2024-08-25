package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

// GenerateJWT genera un token JWT para un usuario
func GenerateJWT(username, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        Role:     role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, err
    }

    return claims, nil
}
