package main

import (
	"net/http"

	gomob "github.com/iranika/gomob"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

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
	var products []gomob.ProductInfo

	for _, url := range u.Urls {
		p := gomob.GetProductInfo(url)
		gomob.SetProductInfo(p)
		products = append(products, p)
	}

	return c.JSON(http.StatusOK, products)
}
