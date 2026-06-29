package utils

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	//"github.com/saisnehith876/StreamMovies/Server/StreamMovie_Server/models"
	//"github.com/golang-jwt/jwt/v5"
	jwt "github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Role      string
	UserId    string
	jwt.RegisteredClaims
}

// SECRET_KEY/SECRET_REFRESH_KEY used to be package-level vars initialized straight from
// os.Getenv. That initialization runs before main() ever gets a chance to call
// godotenv.Load(), so they were always being set to "" unless the secrets were exported
// in the real shell environment. Resolving them lazily fixes that.
func getSecretKey() string {
	if key := os.Getenv("SECRET_KEY"); key != "" {
		return key
	}
	_ = godotenv.Load(".env")
	return os.Getenv("SECRET_KEY")
}

func getRefreshSecretKey() string {
	if key := os.Getenv("SECRET_REFRESH_KEY"); key != "" {
		return key
	}
	_ = godotenv.Load(".env")
	return os.Getenv("SECRET_REFRESH_KEY")
}

func GenerateAllTokens(email, firstName, lastName, role, userId string) (string, string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		UserId:    userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MagicStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(getSecretKey()))

	if err != nil {
		return "", "", err
	}

	refreshClaims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		UserId:    userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MagicStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(getRefreshSecretKey()))

	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil

}

func UpdateAllTokens(userId, token, refreshToken string, client *mongo.Client) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateData := bson.M{
		"$set": bson.M{
			"token":         token,
			"refresh_token": refreshToken,
			"update_at":     updateAt,
		},
	}

	var userCollection *mongo.Collection = database.OpenCollection("users", client)

	_, err = userCollection.UpdateOne(ctx, bson.M{"user_id": userId}, updateData)

	if err != nil {
		return err
	}
	return nil
}

func GetAccessToken(c *gin.Context) (string, error) {
	// authHeader := c.Request.Header.Get("Authorization")
	// if authHeader == "" {
	// 	return "", errors.New("Authorization header is required")
	// }
	// tokenString := authHeader[len("Bearer "):]

	// if tokenString == "" {
	// 	return "", errors.New("Bearer token is required")
	// }
	tokenString, err := c.Cookie("access_token")
	if err != nil {

		return "", err
	}

	return tokenString, nil

}

func ValidateToken(tokenString string) (*SignedDetails, error) {
	claims := &SignedDetails{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil

}

func GetUserIdFromContext(c *gin.Context) (string, error) {
	userId, exists := c.Get("userId")

	if !exists {
		return "", errors.New("userId does not exists in this context")
	}

	id, ok := userId.(string)

	if !ok {
		return "", errors.New("unable to retrieve userId")
	}

	return id, nil

}

func GetRoleFromContext(c *gin.Context) (string, error) {
	role, exists := c.Get("role")

	if !exists {
		return "", errors.New("role does not exists in this context")
	}

	memberRole, ok := role.(string)

	if !ok {
		return "", errors.New("unable to retrieve userId")
	}

	return memberRole, nil

}

func ValidateRefreshToken(tokenString string) (*SignedDetails, error) {
	claims := &SignedDetails{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(getRefreshSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("refresh token has expired")
	}

	return claims, nil
}
