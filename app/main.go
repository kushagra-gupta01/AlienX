package main

import (
	"github.com/kushagra-gupta01/alienx"
)

func main() {
	app := AlienX.New()
	app.Get("/profile", HandleUserProfile)
	app.Start(":3000")
}

func HandleUserProfile(c *AlienX.Context) error {
	return c.Render(profile.Index())
}
