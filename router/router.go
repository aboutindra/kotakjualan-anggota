package router

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kotakjualan-anggota/controller"
	"kotakjualan-anggota/db"
	"time"
)

type Router struct {
}

const dbName = "AnggotaDB"
const collectionName = "User"

func (r Router) GetById(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	obj, _ := primitive.ObjectIDFromHex(c.Params("id"))

	var filter controller.Id

	filter.Id = obj

	fmt.Println(obj)

	c.Body()

	//var filter bson.M = bson.M{}
	//
	//filter = bson.M{"_id": obj}

	//if c.Params("id") != "" {
	//	id := c.Params("id")
	//	objID, _ := primitive.ObjectIDFromHex(id)
	//	filter = bson.M{"_id": objID}
	//}

	//var results []bson.M

	var results controller.Anggota

	var arr []interface{}

	cur, err := collection.Find(context.Background(), filter)
	//defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	fmt.Println(cur)

	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)

	for cur.Next(ct) {
		cur.Decode(&results)
		fmt.Println("Hello")
		fmt.Println(results)
		arr = append(arr, results)
	}

	//cur.All(context.Background(), &results)

	//if results == nil {
	//	c.SendStatus(404)
	//	return
	//}

	json, _ := json.Marshal(arr)
	c.Send(json)
}

func (r Router) CreateAnggota(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var anggota controller.Anggota
	json.Unmarshal([]byte(c.Body()), &anggota)

	res, err := collection.InsertOne(context.Background(), anggota)

	if err != nil {
		c.Status(500).Send(res)
	}

	response, _ := json.Marshal(res)
	c.Send(response)

}

func (r Router) UpdateAnggota(c *fiber.Ctx) {
	collection, err := db.GetMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var anggota controller.Anggota
	json.Unmarshal([]byte(c.Body()), &anggota)

	update := bson.M{
		"$set": anggota,
	}

	_id, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": _id}, update)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)

}
