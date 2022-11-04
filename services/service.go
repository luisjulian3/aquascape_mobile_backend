package service

import (
	"aquascape_backend/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	http "net/http"
	"strings"
)

func EchoHTTPService() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// routes
	e.GET("/PHScale", GetDataPHScale())
	e.GET("/Temp", GetDataTemp())
	e.GET("/Fan", GetDataFan())
	e.GET("/Lamp", GetDataLamp())
	e.PUT("/Fan/UpdateFanTrue", UpdateDataFanTrue())
	e.PUT("/Fan/UpdateFanFalse", UpdateDataFanFalse())
	e.PUT("/Lamp/UpdateLampTrue", UpdateDataLampTrue())
	e.PUT("/Lamp/UpdateLampFalse", UpdateDataLampFalse())
	// e.POST("url", func(s))
	// e.DELETE("url", func(s))

	// run actual server
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
