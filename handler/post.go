package handler

import (
	"fmt"
	"go-face/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func (db Database) CreatePost(c echo.Context) error {
	fmt.Println("CreatePost =======")
	var post model.Post
	if err := c.Bind(&post); err != nil {
		return err
	}

	postID, err := model.InsertPostDB(db.Client, post)
	if err != nil {
		return err
	}

	postWithID := post
	postWithID.ObjectID = postID

	return c.JSON(http.StatusCreated, postWithID)
}

func (db Database) InserComment(c echo.Context) error {
	fmt.Println("InserComment")
	comment := new(model.Comment)
	postId := c.Param("id")

	if err := c.Bind(comment); err != nil {
		fmt.Println("error Bind Post")
		return err
	}

	err := model.InserCommentToPostByIdDB(db.Client, postId, comment)
	if err != nil {
		fmt.Println("error ====> cannot inser comment")
		return err
	}
	post, err := model.GetPostByIdDB(db.Client, postId)
	if err != nil {
		fmt.Println("error ====> cannot get post")
		return err
	}

	return c.JSON(http.StatusOK, gin.H{"data": post})
}

func (db Database) GetPosts(c echo.Context) error {
	posts, err := model.GetPostsDB(db.Client)
	if err != nil {
		fmt.Println("error ====> document not found ", err)
		return err
	}
	// fmt.Printf("Posts =====> %+v", Posts)
	return c.JSON(http.StatusOK, gin.H{"data": posts})
}
