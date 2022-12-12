package middlewares

import (
	"context"
	"github.com/labstack/gommon/log"
	"strings"

	"github.com/labstack/echo/v4"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("serviceAccountkey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}
		log.Fatalf("token fail", token)

		c.Set("token", token)
		return next(c)
	}
}
