package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kushagra-gupta01/AlienX"
	"github.com/kushagra-gupta01/AlienX/app/view/dashboard"
	"github.com/kushagra-gupta01/AlienX/app/view/profile"
)

func main() {
	app := AlienX.New()
	app.Plug(WithRequestId, WithAuth)
	ph := NewProfileHandler(&NOOPSB{})
	app.Get("/profile", ph.HandleProfileIndex)
	app.Get("/dashboard", HandleDashboardIndex)
	app.Start(":3000")
}

func WithAuth(h AlienX.Handler) AlienX.Handler {
	return func(c *AlienX.Context) error {
		fmt.Println("auth")
		c.Set("email", "kk@king.com")
		return h(c)
	}
}

func WithRequestId(h AlienX.Handler) AlienX.Handler {
	return func(c *AlienX.Context) error {
fmt.Println("request")
		c.Set("requestID", uuid.New())
		return h(c)
	}
}

type SupabaseClient interface{
	Auth(foo string) error
	//...some methods
}

type NOOPSB struct{}

func (NOOPSB) Auth(foo string) error{return nil}

type ProfileHandler struct{
	sbClient SupabaseClient
	//...
}

func NewProfileHandler(sb SupabaseClient) *ProfileHandler{
	return &ProfileHandler{
		sbClient: sb,
	}
}

func (h *ProfileHandler) HandleProfileIndex(c *AlienX.Context) error {
	user := profile.User{
		FirstName: "kk",
		LastName:  "g",
		Email:     "kk@kk.in",
	}
	return c.Render(profile.Index(user))
}

func HandleDashboardIndex(c *AlienX.Context) error {
	fmt.Println(c.Get("requestID"))
	return c.Render(dashboard.Index())
}
