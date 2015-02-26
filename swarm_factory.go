package main

import "github.com/go-martini/martini"

// AppEngine for web engine setup
func AppEngine() *martini.ClassicMartini {
	m := martini.Classic()
	return m
}

func main() {
	m := AppEngine()
	m.Run()
}
