package main

import (
	"fmt"
	"github.com/QuanLab/gotaho-service/config"
	"github.com/QuanLab/gotaho-service/model"
	"github.com/QuanLab/gotaho-service/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

func main() {
	var conf = config.Get().Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET(conf.BaseContextPath+ "jobs", func(c echo.Context) error {
		// api/v1/jobs?limit=20&offset=0
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err!= nil {
			limit = 20
		}
		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err!= nil {
			offset = 0
		}
		return c.JSON(http.StatusOK, model.GetJobList(limit, offset))
	})


	e.GET(conf.BaseContextPath+ "search/jobs", func(c echo.Context) error {
		// api/v1/jobs?limit=20&offset=0
		name := c.QueryParam("q")
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err!= nil {
			limit = 20
		}
		return c.JSON(http.StatusOK, service.SearchJobsByName(name, limit))
	})

	e.POST(conf.BaseContextPath+ "startJob", func(c echo.Context) error {
		var job model.Job
		err := c.Bind(&job)
		if err != nil {
			return c.JSON(http.StatusForbidden, "")
		}

		err = service.StartJob(job)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		return c.JSON(http.StatusOK, "")
	})

	e.POST(conf.BaseContextPath+ "stopJob", func(c echo.Context) error {
		var job model.Job
		err := c.Bind(&job)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "")
		}
		err = service.StopJob(job)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		return c.JSON(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
