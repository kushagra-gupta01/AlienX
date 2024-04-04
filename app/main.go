package main

import (
	"github.com/kushagra-gupta01/AlienX"
	"github.com/kushagra-gupta01/AlienX/app/view/profile"
	"github.com/kushagra-gupta01/AlienX/app/view/dashboard"
)

func main() {
	app := AlienX.New()
	app.Get("/profile", HandleProfileIndex)
	app.Get("/dashboard", HandleDashboardIndex)
	app.Start(":3000")
}

func HandleProfileIndex(c *AlienX.Context) error {
	user := profile.User{
		FirstName: "kk",
		LastName:  "g",
		Email:     "kk@kk.in",
	}
	return c.Render(profile.Index(user))
}

func HandleDashboardIndex(c *AlienX.Context) error {
	return c.Render(dashboard.Index())
}
