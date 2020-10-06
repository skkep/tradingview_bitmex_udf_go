package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var base_uri = "https://www.bitmex.com/api/udf/"

func main() {
	//에코 인스턴스 생성
	e := echo.New()
	//미들웨어 선언
	e.Use(middleware.Logger())  //http 요청 기록
	e.Use(middleware.Recover()) //패닉 복구
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/time", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.FormatInt(time.Now().UTC().Unix(), 10))
	})

	e.GET("/config", GetConfig)
	e.GET("/symbols", Symbols)
	e.GET("/search", Search)
	e.GET("/history", History)
	e.Logger.Fatal(e.Start(":80"))

}

func SendReq(path string) string {
	resp1, _ := http.Get(base_uri + path)
	bytes, _ := ioutil.ReadAll(resp1.Body)
	return string(bytes)
}

func History(context echo.Context) error {
	return context.String(200, SendReq("/history?"+context.QueryString()))
}

func Search(context echo.Context) error {
	return context.String(200, SendReq("/search?"+context.QueryString()))

}

func Symbols(context echo.Context) error {
	return context.String(200, SendReq("/symbols?"+context.QueryString()))

}

func GetConfig(context echo.Context) error {
	return context.String(200, SendReq("/config?"+context.QueryString()))

}
