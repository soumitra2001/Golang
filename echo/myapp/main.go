package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (v ProductValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	e := echo.New()
	v := validator.New()

	products := []map[int]string{{1: "mobile"}, {2: "laptop"}, {3: "tv"}}

	e.GET("/products", func(r echo.Context) error {
		return r.JSON(http.StatusOK, products)
	})

	e.GET("/product/:id", func(r echo.Context) error {
		var product map[int]string
		id, err := strconv.Atoi(r.Param("id"))

		if err != nil {
			log.Fatal(err)
		}

		for _, p := range products {
			for key, _ := range p {
				if key == id {
					product = p
				}
			}
		}

		if product == nil {
			r.String(http.StatusBadRequest, "Invalid product id")
			return nil
		}

		return r.JSON(http.StatusOK, product)

	})

	e.POST("/product", func(c echo.Context) error {
		type Product struct {
			Name    string `json:"product_name" validate:"required,min=2"`
			Vendor  string `json:"vendor" validate:"min=4,max=10"`
			Email   string `json:"email" validate:"required_with=Vendor,email"`
			Website string `json:"website" validate:"url"`
			Country string `json:"country" validate:"len=2"`
		}

		var reqBody Product
		e.Validator = &ProductValidator{validator: v}

		if err := c.Bind(&reqBody); err != nil {
			return err
		}

		// if err := v.Struct(reqBody); err != nil {
		// 	return err
		// }
		if err := c.Validate(reqBody); err != nil {
			return err
		}

		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}

		products = append(products, product)

		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
