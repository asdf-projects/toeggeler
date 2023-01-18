package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steinm91/toeggeler/toeggeler-server/models"
)

type UserController struct {
	UserService *models.UserService
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Mail     string `json:"mail" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Mail string `json:"mail" binding:"required,email"`
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  create a new user
// @Tags		 users
// @Accept       json
// @Produce      json
// @Param        user body CreateUserRequest true "Create user"
// @Success      200  {object}  models.User
// @Router       /users [post]
func (userCtrl UserController) CreateUser(c *gin.Context) {
	var userRequest CreateUserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid object provided")
		return
	}

	user, err := userCtrl.UserService.Create(
		userRequest.Username,
		userRequest.Mail,
		userRequest.Password,
	)
	if err != nil {
		if errors.Is(err, models.ErrUserExists) {
			c.String(http.StatusBadRequest, "User/E-Mail already exists")
		} else {
			c.String(http.StatusInternalServerError, "Could not create new user")
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary      Get a list of all available users
// @Description  Get a list of all available users
// @Tags		 users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.User
// @Router       /users [get]
func (userCtrl UserController) GetUsers(c *gin.Context) {
	users, err := userCtrl.UserService.GetUsers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Get a user by Id
// @Description  Get a user by Id
// @Tags		 users
// @Accept       json
// @Produce      json
// @Param		 id path int true "User ID"
// @Success      200  {object}  []models.User
// @Failure      404
// @Router       /users/{id} [get]
func (userCtrl UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := userCtrl.UserService.GetUser(id)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.String(http.StatusNotFound, "User not found")
		} else {
			c.String(http.StatusInternalServerError, "Could not get user")
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  update an existing user
// @Tags		 users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        user body UpdateUserRequest true "User user"
// @Success      200  {object}  models.User
// @Failure      404
// @Router       /users/{id} [put]
func (userCtrl UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateRequest UpdateUserRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.String(http.StatusBadRequest, "Invalid payload")
	}

	user, err := userCtrl.UserService.UpdateUser(id, updateRequest.Mail)
	if err != nil {
		if errors.Is(err, models.ErrUserExists) {
			c.String(http.StatusBadRequest, "Mail already exists")
		} else {
			c.String(http.StatusInternalServerError, "Could not update user")
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary      Delete an existing user
// @Description  delete an existing user
// @Tags		 users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200
// @Failure      404
// @Router       /users/{id} [delete]
func (userCtrl UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := userCtrl.UserService.DeleteUser(id)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.String(http.StatusBadRequest, "User not found")
		} else {
			c.String(http.StatusInternalServerError, "Could not delete user")
		}
		return
	}

	c.Status(http.StatusOK)
}
