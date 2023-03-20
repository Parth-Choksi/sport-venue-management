package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sport-venue-management/configs"
	"sport-venue-management/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userDetailCollection *mongo.Collection = configs.GetCollection(configs.DB, "userDetails")
var validate = validator.New()
var SECRET_KEY = []byte("gosecretkey")

func RegisterUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var user models.UserDetails
	json.NewDecoder(req.Body).Decode(&user)
	user.User_Id = uuid.NewString()
	user.Password = getHash([]byte(user.Password))
	// collection := configs.GetCollection("sports-venue-management",'')
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := userDetailCollection.InsertOne(ctx, user)
	json.NewEncoder(res).Encode(result)
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func LoginUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user models.UserDetails
	var dbUser models.UserDetails

	json.NewDecoder(req.Body).Decode(&user)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := userDetailCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		res.Write([]byte(`{"response":"Wrong Password"}`))
	}

	jwtToken, err := GenerateJWT()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	res.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT Token")
		return "", err
	}
	return tokenString, nil
}
