package main

import (
	"github.com/labstack/echo/v4"
	"go-practice/web/scrapper"
	"os"
	"strings"
)

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//scrapper.Scrape("python")
}

func handleHome(c echo.Context) error {
	return c.File("web/home.html")
}

const fileName string = "web/jobs.csv"

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, "job.csv")
}
