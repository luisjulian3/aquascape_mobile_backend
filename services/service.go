package service

import (
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/luisjulian3/aquascape_mobile_backend/models"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"io/ioutil"
	http "net/http"
	"strings"
	"time"
)

func EchoHTTPService() {

	e := echo.New()

	// Middleware
	//e.Use(middlewares.Auth())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//e.GET("/api/private", privateAPI, middleware.Auth())

	// routes

	e.GET("/PHScale", GetDataPHScale())
	e.GET("/Temp", GetDataTemp())
	e.GET("/Fan", GetDataFan())
	e.GET("/Lamp", GetDataLamp())
	e.PUT("/Fan/UpdateFanTrue", UpdateDataFanTrue())
	e.PUT("/Fan/UpdateFanFalse", UpdateDataFanFalse())
	e.PUT("/Lamp/UpdateLampTrue", UpdateDataLampTrue())
	e.PUT("/Lamp/UpdateLampFalse", UpdateDataLampFalse())
	e.GET("/sensor/postsensor", PostSensor())
	// e.POST("url", func(s))
	// e.DELETE("url", func(s))

	// run actual server
	e.Logger.Fatal(e.Start(":8080"))
}

/*func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	log.Printf("test123")
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}*/

func getUser(ctx context.Context, app *firebase.App) *auth.UserRecord {
	uid := "some_string_uid"

	// [START get_user]
	// Get an auth client from the firebase.App
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user]
	return u
}

func GetDataPHScale() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/PHScale.json"
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultPH

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func GetDataTemp() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Temperature.json"
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultTemp

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func GetDataFan() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultFan

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func GetDataLamp() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Lamp.json"
		resp, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultLamp

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateDataFanTrue() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": true}"

		payload := strings.NewReader(body)

		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultFan

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateDataFanFalse() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": false}"

		payload := strings.NewReader(body)

		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Fan.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultFan

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateDataLampTrue() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": true}"

		payload := strings.NewReader(body)

		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Lamp.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultLamp

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func UpdateDataLampFalse() echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &http.Client{}

		body := "{\"status\": false}"

		payload := strings.NewReader(body)

		url := "https://aquascape-mobile-default-rtdb.asia-southeast1.firebasedatabase.app/ESP8266_Aqua/Lamp.json"

		req, err := http.NewRequest(http.MethodPut, url, payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var results models.ResultLamp

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = json.Unmarshal(response, &results)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, results)
	}
}

func PostSensor() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		sa := option.WithCredentialsFile("keyF.json")
		//conf := &config.Config{ProjectID: "aquascape-mobile"}
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		client, err := app.Firestore(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer client.Close()

		phscale := c.QueryParam("phscale")
		temp := c.QueryParam("temp")

		postdatasensor, _, err := client.Collection("sensor").Add(ctx, map[string]interface{}{
			"phscale": phscale,
			"temp":    temp,
			"time":    time.Now(),
		})

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		fmt.Printf("Document data: %#v\n", postdatasensor)

		return c.JSON(http.StatusOK, postdatasensor)
	}
}
