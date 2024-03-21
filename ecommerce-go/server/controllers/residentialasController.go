package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	middlewares "github.com/umangraval/Go-Mongodb-REST-boilerplate/handlers"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/models"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/validators"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var GetResidentialas = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var residentialas []*models.Residential
	params := mux.Vars(request)
	category := params["category"]
	collection := client.Database("resident").Collection("residentials")
	filter := bson.M{}
	if category != "Все" {
		filter["category"] = category
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var residential models.Residential
		err := cursor.Decode(&residential)
		if err != nil {
			log.Fatal(err)
		}

		residentialas = append(residentialas, &residential)
	}
	if err := cursor.Err(); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessArrResidentialasRespond(residentialas, response)
})

var CreateResidential = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var residential models.Residential
	err := json.NewDecoder(request.Body).Decode(&residential)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	if ok, errors := validators.ValidateInputs(residential); !ok {
		middlewares.ValidationResponse(errors, response)
		return
	}
	collection := client.Database("resident").Collection("residentials")
	result, err := collection.InsertOne(context.TODO(), residential)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)

})

var GetResidentEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("id:", id)
	var residential models.Residential

	collection := client.Database("resident").Collection("residentials")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&residential)
	if err != nil {
		middlewares.ErrorResponse("Person does not exist", response)
		return
	}
	middlewares.SuccessResidentialasRespond(&residential, response)
})

var DeleteResidentEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("id:", id)
	var residential models.Residential

	collection := client.Database("resident").Collection("residentials")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&residential)
	if err != nil {
		middlewares.ErrorResponse("Person does not exist", response)
		return
	}
	_, derr := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if derr != nil {
		middlewares.ServerErrResponse(derr.Error(), response)
		return
	}
	middlewares.SuccessResponse("Deleted", response)
})
