package handlers

import (
	"context"
	"cookbook/db"
	"cookbook/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	database   = "cookbook"
	collection = "recipes"
)

func HandleRecipes(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		GetRecipes(ctx)
	case http.MethodPost:
		PostRecipe(ctx)
	default:
		ctx.String(http.StatusNotFound, "404 page not found")
		return
	}
}

func GetRecipes(ctx *gin.Context) {
	coll := db.GetCollection(database, collection)

	cursor, err := coll.Find(context.Background(), gin.H{})
	if err != nil {
		ctx.Error(NewAPIError(http.StatusInternalServerError, "failed to query db", err.Error()))
		return
	}

	var res []models.Recipe
	err = cursor.All(context.Background(), &res)
	if err != nil {
		ctx.Error(NewAPIError(http.StatusInternalServerError, "failed to parse db return", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func PostRecipe(ctx *gin.Context) {
	coll := db.GetCollection(database, collection)

	var recipe models.Recipe
	err := json.NewDecoder(ctx.Request.Body).Decode(&recipe)
	if err != nil {
		ctx.Error(NewAPIError(http.StatusBadRequest, "failed to read request body", err.Error()))
		return
	}

	res, err := coll.InsertOne(context.Background(), &recipe)
	if err != nil {
		ctx.Error(NewAPIError(http.StatusInternalServerError, "failed to insert recipe", err.Error()))
		return
	}

	logrus.Infof("RES: %+v", res)
	recipe.ID = models.ObjectID(res.InsertedID.(primitive.ObjectID).Hex())

	ctx.JSON(http.StatusCreated, recipe)

}

func DeleteRecipe(ctx *gin.Context) {
	coll := db.GetCollection(database, collection)

	id := ctx.Param("id")

	//TODO fix the below filter so it actually deletes by ID value.
	res, err := coll.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		ctx.Error(NewAPIError(http.StatusBadRequest, "failed to delete recipe", err.Error()))
		return
	}

	logrus.Infof("Res: %+v", res)

	ctx.Status(http.StatusOK)
}
