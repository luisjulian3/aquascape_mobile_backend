package service

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/luisjulian3/aquascape_mobile_backend/models"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
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

	//Sensor Post Data using params
	e.GET("/sensor/postsensor", PostSensor())

	//Fan - > FireStore hit
	e.GET("/fan", GetFan())
	e.GET("/fan/true", PostFanTrue())
	e.GET("/fan/false", PostFanFalse())

	//Lamp - > FireStore hit
	e.GET("/lamp", GetLamp())
	e.GET("/lamp/true", PostLampTrue())
	e.GET("/lamp/false", PostLampFalse())

	//Temp - > FireStore
	e.GET("/temp/data", GetTempData()) // 1 data
	e.GET("/temp/real", GetTempReal()) //RealTimeData

	//PHScale - > Firestore
	e.GET("/phscale/data", GetPHScaleData()) //1 data
	e.GET("/phscale/real", GetPHScaleReal()) //RealTimeData

	e.Logger.Fatal(e.Start(":8080"))
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

func PostLampTrue() echo.HandlerFunc {
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

		postdatalamp, err := client.Collection("device").Doc("devices").Set(ctx, map[string]interface{}{
			"lamp": true,
		}, firestore.MergeAll)

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		fmt.Printf("Document data: %#v\n", postdatalamp)

		return c.JSON(http.StatusOK, postdatalamp)
	}
}

func PostLampFalse() echo.HandlerFunc {
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

		postdatalamp, err := client.Collection("device").Doc("devices").Set(ctx, map[string]interface{}{
			"lamp": false,
		}, firestore.MergeAll)

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		fmt.Printf("Document data: %#v\n", postdatalamp)

		return c.JSON(http.StatusOK, postdatalamp)
	}
}

func PostFanFalse() echo.HandlerFunc {
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

		postdatalamp, err := client.Collection("device").Doc("devices").Set(ctx, map[string]interface{}{
			"fan": false,
		}, firestore.MergeAll)

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		fmt.Printf("Document data: %#v\n", postdatalamp)

		return c.JSON(http.StatusOK, postdatalamp)
	}
}

func PostFanTrue() echo.HandlerFunc {
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

		postdatalamp, err := client.Collection("device").Doc("devices").Set(ctx, map[string]interface{}{
			"fan": true,
		}, firestore.MergeAll)

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

		fmt.Printf("Document data: %#v\n", postdatalamp)

		return c.JSON(http.StatusOK, postdatalamp)
	}
}

func GetTempData() echo.HandlerFunc {
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

		iter := client.Collection("sensor").Documents(ctx)

		var results []models.NewResultDataPH
		for {
			var result models.NewResultDataPH
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			data := doc.Data()["temp"]
			time := doc.Data()["time"]

			result.Time = fmt.Sprintf("%v", time)
			result.Value = fmt.Sprintf("%v", data)
			results = append(results, result)
		}
		return c.JSON(http.StatusOK, results)
	}
}

func GetPHScaleData() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		sa := option.WithCredentialsFile("keyF.json")
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		client, err := app.Firestore(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer client.Close()

		iter := client.Collection("sensor").Documents(ctx)

		var results []models.NewResultDataPH
		for {
			var result models.NewResultDataPH
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			data := doc.Data()["phscale"]
			time := doc.Data()["time"]

			result.Time = fmt.Sprintf("%v", time)
			result.Value = fmt.Sprintf("%v", data)
			results = append(results, result)
		}
		return c.JSON(http.StatusOK, results)
	}
}

func GetTempReal() echo.HandlerFunc {
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

		iter := client.Collection("sensor").OrderBy("temp", firestore.Asc).Documents(ctx)

		var test []string
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			data := doc.Data()["temp"]
			test = append(test, fmt.Sprintf("%v", data))
		}

		return c.JSON(http.StatusOK, test[len(test)-1])
	}
}

func GetPHScaleReal() echo.HandlerFunc {
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

		iter := client.Collection("sensor").OrderBy("phscale", firestore.Asc).Documents(ctx)

		var test []string
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			data := doc.Data()["phscale"]
			test = append(test, fmt.Sprintf("%v", data))

		}
		//test = json.Unmarshal([]byte(dataPhscale{}), test)
		return c.JSON(http.StatusOK, test[len(test)-1])
	}
}

func GetFan() echo.HandlerFunc {
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

		dsnap, err := client.Collection("device").Doc("devices").Get(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		m := dsnap.Data()["fan"]
		fmt.Printf("Document data: %#v\n", m)

		return c.JSON(http.StatusOK, m)
	}
}

func GetLamp() echo.HandlerFunc {
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

		dsnap, err := client.Collection("device").Doc("devices").Get(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		m := dsnap.Data()["lamp"]
		fmt.Printf("Document data: %#v\n", m)

		return c.JSON(http.StatusOK, m)
	}
}
