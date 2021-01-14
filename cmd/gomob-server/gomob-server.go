package main

import (
	"net/http"

	"github.com/iranika/gomob"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/dlsitesq", dlsitesq)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type Product struct {
	Urls []string `json: "urls"`
}

// Handler
func dlsitesq(c echo.Context) error {
	u := new(Product)
	if err := c.Bind(u); err != nil {
		return err
	}
	var product []gomob.ProductInfo

	for i, url := range u.Urls {
		product += gomob.getProductInfo(url)
	}
	return c.JSON(http.StatusOK, u)
}
