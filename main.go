package main

import (
	"fmt"
	_ "gin-swagger-example/docs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin CRUD Swagger Example API
// @version 1.0
// @description This is a sample server for a Gin-Swagger example with CRUD operations.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// mock database
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

func main() {
	r := gin.Default()

	// Routes
	r.GET("/api/v1/users", getUsers)
	r.POST("/api/v1/users", createUser)
	r.GET("/api/v1/users/:id", getUser)
	r.PUT("/api/v1/users/:id", updateUser)
	r.DELETE("/api/v1/users/:id", deleteUser)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// getUsers godoc
// @Summary Get all users
// @Description Get all users from the system
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// createUser godoc
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User data"
// @Success 201 {object} User
// @Router /users [post]
func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = users[len(users)-1].ID + 1
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// getUser godoc
// @Summary Get a user by ID
// @Description Get a single user from the system by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// updateUser godoc
// @Summary Update a user by ID
// @Description Update a single user in the system by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body User true "User data"
// @Success 200 {object} User
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			if err := c.BindJSON(&users[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// deleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a single user from the system by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"message": "user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}
