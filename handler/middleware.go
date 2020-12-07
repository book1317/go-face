package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bodyDumpConfig() middleware.BodyDumpConfig {
	handler := func(c echo.Context, req, res []byte) {
		log.Println("headers:", c.Request().Header)
		log.Println("request:", string(req))
		log.Println("response:", string(res))
	}
	return middleware.BodyDumpConfig{Handler: handler}
}

func notFoundOnProduction(state string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if state != "dev" {
				return c.NoContent(http.StatusNotFound)
			}
			return next(c)
		}
	}
}

func enable(b bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !b {
				return c.NoContent(http.StatusBadRequest)
			}
			return next(c)
		}
	}
}

// func verifyPassword(cv *app.Config) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			var bodyBytes []byte
// 			if c.Request().Body != nil {
// 				bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
// 			}

// 			var merchant model.Merchant
// 			if err := json.Unmarshal(bodyBytes, &merchant); err != nil {
// 				// panic(err)
// 				return c.JSON(http.StatusBadRequest, err)
// 			}

// 			// Restore the io.ReadCloser to its original state
// 			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

// 			var tempMerchant model.Merchant
// 			coll := cv.MongoDB.Client.Database("quiz").Collection("merchant")
// 			coll.FindOne(c.Request().Context(), bson.M{"username": merchant.Username}).Decode(&tempMerchant)
// 			if (tempMerchant == model.Merchant{}) {
// 				return c.JSON(http.StatusBadRequest, "no username")
// 			}

// 			if tempMerchant.Password != merchant.Password {
// 				// return errors.New("wrong password")
// 				return c.JSON(http.StatusBadRequest, "wrong password")
// 			}

// 			return next(c)
// 		}
// 	}
// }
