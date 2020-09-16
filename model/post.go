package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	ObjectID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content  string             `json:"content" bson:"content"`
	Like     int                `json:"like" bson:"like"`
	OwnerId  primitive.ObjectID `json:"owner_id" bson:"owner_id"`
}

type Post struct {
	ObjectID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content  string             `json:"content" bson:"content"`
	Like     int                `json:"like" bson:"like"`
	Comments []Comment          `json:"comments" bson:"comments"`
	OwnerId  primitive.ObjectID `json:"owner_id" bson:"owner_id"`
}

func InsertPostDB(client *mongo.Client, post *Post) (primitive.ObjectID, error) {
	col := client.Database(db_facebook).Collection(co_post)
	result, err := col.InsertOne(context.TODO(), post)
	postID, _ := result.InsertedID.(primitive.ObjectID)
	return postID, err
}

func GetPostsDB(client *mongo.Client) ([]Post, error) {
	fmt.Println("====> getPostsDB")
	var result []Post
	col := client.Database(db_facebook).Collection(co_post)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(context.TODO())
	} else {
		for cursor.Next(context.TODO()) {
			var res bson.M
			err := cursor.Decode(&res)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
			} else {
				var Post Post
				bsonBytes, _ := bson.Marshal(res)
				bson.Unmarshal(bsonBytes, &Post)
				result = append(result, Post)
			}
		}
	}
	return result, err
}
func GetPostByIdDB(client *mongo.Client, postID string) (Post, error) {
	var post Post
	col := client.Database(db_facebook).Collection(co_post)
	id, _ := primitive.ObjectIDFromHex(postID)
	err := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&post)
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
	}
	return post, err
}

func InserCommentToPostByIdDB(client *mongo.Client, postID string, comment *Comment) error {
	id, _ := primitive.ObjectIDFromHex(postID)
	col := client.Database(db_facebook).Collection(co_post)
	result, err := col.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$push", bson.D{{"comments", comment}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}
