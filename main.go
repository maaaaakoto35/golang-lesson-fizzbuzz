package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!\n") })
	e.GET("/fizzbuzz/:num", func(c echo.Context) error {
		num, err := strconv.Atoi(c.Param("num"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		res, err := fizzbuzz(num)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":3333"))
}

func fizzbuzz(num int) (result string, err error) {
	// exception
	if num < 0 {
		err = newError("this is negative number")
		return
	}
	for count := 1; count <= num; count++ {
		if count%15 == 0 {
			result += "FIZZ BUZZ\n"
		} else if count%5 == 0 {
			result += "BUZZ\n"
		} else if count%3 == 0 {
			result += "FIZZ\n"
		} else {
			str := strconv.Itoa(count)
			result += str + "\n"
		}
	}
	return
}

func newError(str string) error {
	return errors.New(str)
}
