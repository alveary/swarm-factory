package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

// User is a new member
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
	var port int
	flag.IntVar(&port, "p", 9001, "the port number")
	flag.Parse()

	m := AppEngine()
	m.RunOnAddr(":" + strconv.Itoa(port))
}
