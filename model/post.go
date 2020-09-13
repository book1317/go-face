package model

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	ObjectID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content  string             `json:"content" bson:"content"`
	Like     int                `json:"like" bson:"like"`
	IsLike   bool               `json:"isLike" bson:"isLike"`
	Owner    Profile            `json:"owner" bson:"owner"`
}

type Post struct {
	ObjectID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content  string             `json:"content" bson:"content"`
	Like     int                `json:"like" bson:"like"`
	Comments []Comment          `json:"comments" bson:"comments"`
	IsLike   bool               `json:"isLike" bson:"isLike"`
	Owner    Profile            `json:"owner" bson:"owner"`
}

func (db Database) CreatePost(c echo.Context) error {
	fmt.Printf("CreatePost=======")
	post := new(Post)

	if err := c.Bind(post); err != nil {
		fmt.Printf("post ==== > %+v", post)
		fmt.Println("error Bind Post")
		return err
	}

	postID, err := db.insertPostDB(*post)
	if err != nil {
		return err
	}

	postWithID := post
	postWithID.ObjectID = postID

	return c.JSON(http.StatusCreated, postWithID)
}

func (db Database) GetPosts(c echo.Context) error {
	Posts, err := getPostDB(db.Client)
	if err != nil {
		fmt.Println("error ====> document not found")
		return err
	}
	fmt.Printf("Posts =====> %+v", Posts)
	return c.JSON(http.StatusOK, Posts)
}

func (db Database) insertPostDB(post Post) (primitive.ObjectID, error) {
	col := db.Client.Database(db_facebook).Collection(co_post)
	result, err := col.InsertOne(context.TODO(), post)
	fmt.Printf("result =====> %+v", result)
	postID, _ := result.InsertedID.(primitive.ObjectID)
	return postID, err
}

func getPostDB(client *mongo.Client) ([]Post, error) {
	fmt.Println("====> getPostDB")
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