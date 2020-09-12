package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jatindumka/bookstore_users-api/domain/users"
	"github.com/jatindumka/bookstore_users-api/services"
	"github.com/jatindumka/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		nil
	}
	c.String(http.StatusOK, user)

}

// func FindUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")

// }
