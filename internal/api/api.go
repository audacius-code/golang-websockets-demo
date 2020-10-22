/*
Copyright © 2020 Fernando Silva <audacius@pm.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package api

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"
)

// For demo purpose only. In a real scenario, it would be part of the `auth`
// domain: router, routes, logic, types, interfaces, and etc.
var jwtSecret = viper.GetString("JWT_SECRET")

// For demo purpose only. In a real scenario, it would be part of the `user`
// domain: router, routes, logic, types, interfaces, and etc.
type userLoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
 * Handlers
 */

func login(c echo.Context) error {
	// For demo purpose only
	username := viper.GetString("AUTH_USERNAME")
	password := viper.GetString("AUTH_PWD")

	u := new(userLoginInfo)
	err := c.Bind(&u)
	if err != nil {
		return err
	}

	// In a real scenario, it would be checked first checked against a hot storage,
	// for example Redis, if it's not there, check against the cold storage - for
	// example MySQL.
	//
	// Note: Password would be compared hashed not plain
	if u.Username != username || u.Password != password {
		return echo.ErrUnauthorized
	}

	t, err := GenerateToken()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Send JWT token information, specifically the access level
func getLevel(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	level := claims["level"].(string)

	return c.JSONBlob(http.StatusOK, []byte(`{"level": `+level+`}`))
}

// Establish the websocket connection.
// For demo purposes: it's a very simple implementation of the classic `echo`
func connect(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		for {
			// Read
			var msg string
			err := websocket.Message.Receive(ws, &msg)

			if err == nil {
				fmt.Printf("Received message from \"%s\": %s\n", ws.RemoteAddr(), msg)
			}

			// Write
			websocket.Message.Send(ws, "Echo: "+msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

// Run the API server
func Run(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Website
	e.Static("/", "website")

	// Entrypoint
	e.POST("/login", login)

	// API
	r := e.Group("/api/v1")
	r.Use(middleware.JWT([]byte(jwtSecret)))

	// API Routes
	r.GET("/token/level", getLevel)

	// Websocket demo
	e.GET("/ws", connect)

	e.Logger.Fatal(e.Start(":" + port))
}
