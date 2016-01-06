package controllers

import (
	"github.com/revel/revel"
	"revel_cms/app/models"
)

func (c App) Login(email, password string) revel.Result {
	person := c.getUser(email)
	if person != nil {
		err := compare([]byte(person.Password), []byte(password))
		if err == nil {
			c.Session["user"] = person.Email
			c.Session.SetDefaultExpiration()
			c.Flash.Success("Bem Vindo ," + person.Name)
			return c.Redirect(App.Index)
		}
		c.Flash.Error("Login Failed try again")
		return c.Render()
	}
	return c.Render()
}
