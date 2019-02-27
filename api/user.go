package api

import (
	"github.com/Fallensouls/Pandora/cache"
	"github.com/Fallensouls/Pandora/models"
	"github.com/Fallensouls/Pandora/util/jsonutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var (
		user models.User
		err  error
	)
	defer func() { c.Set("error", err) }()

	if c.BindJSON(&user) != nil {
		return
	}

	if err = user.AddUser(); err != nil {
		return
	}

	// TODO: send new user an email to activate his account.
	//token, _ := jsonutil.GenerateAccessJWT(user.Id)
	//url := "http://pandora.com/auth/activate?token=" + token
	//log.Println(url)
	//send email...
	c.Status(http.StatusOK)
}

func ActivateUser(c *gin.Context) {
	token := c.Query("token")
	id, _, err := jsonutil.ValidateAccessJWT(token)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var user models.User
	user.Id = id

	if err := user.ActivateUser(); err != nil {
		c.Set("error", err)
		return
	}

	c.Status(http.StatusOK)
}

func RestrictUser(c *gin.Context) {
	id := c.GetInt64("id")
	var user models.User
	user.Id = id
	if err := user.RestrictUser(); err != nil {
		c.Set("error", err)
		return
	}

	if err := cache.SetJWTDeadline(id); err != nil {
		c.Set("error", err)
		return
	}

	c.Status(http.StatusOK)
}

func BanUser(c *gin.Context) {
	id := c.GetInt64("id")
	var user models.User
	user.Id = id
	if err := user.BanUser(); err != nil {
		c.Set("error", err)
		return
	}
	if err := cache.SetJWTDeadline(id); err != nil {
		c.Set("error", err)
		return
	}

	c.Status(http.StatusOK)
}

func UpdateProfile(c *gin.Context) {
	var user models.User
	if c.BindJSON(&user) != nil {
		return
	}

	id := c.GetInt64("id")
	if err := user.UpdateUserProfile(id); err != nil {
		c.Set("error", err)
		return
	}

	c.Status(http.StatusOK)
}

func GetProfile(c *gin.Context) {
	id := c.GetInt64("id")
	var user models.User

	if err := user.GetUser(id); err != nil {
		c.Set("error", err)
		return
	}

	c.JSON(http.StatusOK, Response{Data: user})
}
