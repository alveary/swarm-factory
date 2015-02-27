package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-martini/martini"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//
type FactoryTestSuite struct {
	suite.Suite
	app *martini.ClassicMartini
}

//
func (suite *FactoryTestSuite) SetupTest() {
	suite.app = AppEngine()
}

func (suite *FactoryTestSuite) TestFactoryWithInvalidMailAndPassword() {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(`{email: "falk.hoppe@innoq.com", password: "1234"}`))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	w := httptest.NewRecorder()

	suite.app.ServeHTTP(w, req)
	assert.Equal(suite.T(), w.Code, 201, "User should be created")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(FactoryTestSuite))
}
