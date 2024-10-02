package handlers

import (
	"ApiDoc/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 				Creates a new user
// @Description 		This endpoint let the user register themself by providing user information
// @Accepts 				json
// @Produce 				json
// @Param 					user 	body	 	models.Users 	true "user object"
// @Success 				201 	{object} 	models.Users 	"user created successfully"
// @Failure 				400 	{object} 	string 				"Bad request"
// @Router 					/adduser [post]
func AddUserHandler(u *[]models.Users, c *gin.Context) {
	var userinfo models.Users

	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	userinfo.UserId = len(*u)
	*u = append(*u, userinfo)

	c.JSON(http.StatusOK, gin.H{"Message": "User added succesfully"})

}

// @Summery 				Get existing users info
// @Description 		This end-point takes the user id and return the useinfo with that id
// @Accepts 				json
// @Produce 				json
// @Param 					id   path 	int 	true 	"User ID"
// @Success					200 	{Object} 	model.Users 	"Info about the user"
// @Fail 						404 	{object} 	string 				"No such user found"
// @Router 					/getuser/{id} 	[get]
func GetUserInfo(u *[]models.Users, c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user id given"})
		return
	}

	if id < 0 || id > len(*u) {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such user found"})
		return
	}

	user := (*u)[id]
	c.JSONP(http.StatusOK, user)

}
