package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nantaphop/go-clean-arch/app/models"
	"github.com/nantaphop/go-clean-arch/app/user"
)

type restUserDelivery struct {
	userUc user.Usecase
}

// NewUserDelivery will return rest controller of users
func NewUserDelivery(r *gin.Engine, userUc user.Usecase) {
	userDelivery := restUserDelivery{userUc}
	r.POST("/users", userDelivery.createUser)
}

func (d restUserDelivery) createUser(c *gin.Context) {
	var u models.User
	c.BindJSON(&u)
	result, err := d.userUc.CreateUser(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
}
