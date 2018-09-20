package pages

import (
	"github.com/abaft/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pe5er/radiotutor/user"
)

func LoginGET(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get("loggedIn")

	if v == nil {
		c.HTML(200, "login.html", nil)
	} else {
		c.HTML(200, "login-successful.html", gin.H{"User": v.(user.User)})
	}
}

func LoginPOST(c *gin.Context) {
	session := sessions.Default(c)

	rawUsername, ok := c.GetPostForm("username")
	if !ok || rawUsername == "" {
		LoginGET(c)
		return
	}
	rawPassword, ok := c.GetPostForm("password")
	if !ok || rawPassword == "" {
		LoginGET(c)
		return
	}

	u, err := user.AuthAttempt(rawUsername, rawPassword)
	if err != nil {
		c.HTML(200, "login.html", gin.H{
			"ErrorTitle":   "ERROR",
			"ErrorMessage": err.Error(),
		})
		return
	}
	session.Set("loggedIn", u)
	session.Save()
	LoginGET(c)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("loggedIn")
	session.Save()
	Home(c)
}

func RemoveUser(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("loggedIn")
	if v != nil {
		if err := user.DeleteUser(v.(user.User)); err != nil {
			c.String(200, "ERROR: "+err.Error())
		}
	}
	session.Delete("loggedIn")
	session.Save()
	Home(c)
}

func AccountGET(c *gin.Context) {
	session := sessions.Default(c)

	v := session.Get("loggedIn")

	if v == nil {
		c.HTML(200, "login.html", nil)
	} else {
		c.HTML(200, "account.html", gin.H{"User": v.(user.User)})
	}
}

