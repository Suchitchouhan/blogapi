// controllers/blog_controller.go

package controllers

import (
	"blogapi/models"
	"blogapi/payloads"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlogPosts(c *gin.Context) {
	var blogs []models.BlogPost
	result := models.DB.Find(&blogs)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": blogs})
	}
	// Implement code to retrieve all blog posts from the database
}

func GetBlogPost(c *gin.Context) {
	var blog models.BlogPost
	id := c.Param("id")
	if result := models.DB.Where("id = ?", id).First(&blog).RowsAffected; result == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": blog})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Found"})
	}
	// Implement code to retrieve a specific blog post by ID from the database
}

func CreateBlogPost(c *gin.Context) {
	var payloads payloads.BlogPayload
	var blog models.BlogPost
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := models.DB.Where("Title = ?", payloads.Title).First(&blog).RowsAffected; result == 0 {
		blog := models.BlogPost{Title: payloads.Title, Content: payloads.Content}
		models.DB.Create(&blog)
		c.JSON(http.StatusOK, gin.H{"message": "Blog has been successfully created."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "blog is already Exists."})
	}

	// Implement code to create a new blog post and save it to the database
}

func UpdateBlogPost(c *gin.Context) {
	var payloads payloads.BlogPayload
	var blog models.BlogPost
	id := c.Param("id")
	if err := c.ShouldBindJSON(&payloads); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := models.DB.Where("id = ?", id).First(&blog)
	if result.RowsAffected == 1 {
		updatedBlog := models.BlogPost{Title: payloads.Title, Content: payloads.Content}
		result.Updates(&updatedBlog)
		c.JSON(http.StatusOK, gin.H{"message": "Blog has been successfully updatw."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "blog is not Exists."})
	}

	// Implement code to update an existing blog post in the database
}

func DeleteBlogPost(c *gin.Context) {
	var blog models.BlogPost
	id := c.Param("id")
	result := models.DB.Where("id = ?", id).Delete(&blog)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"message": "blog has been successfully deleted"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog is not exists",
			"error":   result.Error,
		})
	}
	// Implement code to delete a blog post from the database
}
