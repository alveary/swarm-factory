package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

// User is a new swarm family member
type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()

	m.Post("/", binding.Json(User{}), func(user User, resp http.ResponseWriter) {
		time.Sleep(10 * time.Second)
		fmt.Println(fmt.Sprintf("%s : %s", user.Email, user.Password))

		resp.WriteHeader(http.StatusCreated)
	})

	return m
}

func main() {
	m := AppEngine()
	m.RunOnAddr(":9001")
}
