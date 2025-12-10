package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wonli/aqi"
	"github.com/wonli/aqi/ws"
)

func main() {
	app := aqi.Init(
		aqi.ConfigFile("config.yaml"),
		aqi.HttpServer("Aqi", "port"),
	)

	// launchWS(app)
	launchHTTP(app)

	app.Start()
}

func launchHTTP(app *aqi.AppConfig) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fmt.Println(r.Header)
		fmt.Println(r.URL)
		fmt.Println(r.Method)
		fmt.Println(r.Host)
		fmt.Println(r.Proto)
		// fmt.Println(r.FormValue("address"))
		// fmt.Println(r.URL.Query()["abc"])
		w.Write([]byte("server running"))
	})

	app.WithHttpServer(h)

	app.Start()
}

func launchWS(app *aqi.AppConfig) {
	engine := gin.Default()
	// Handler
	engine.GET("/ws", func(c *gin.Context) {
		ws.HttpHandler(c.Writer, c.Request)
	})

	// Router
	wsr := ws.NewRouter()
	wsr.Add("hi", func(a *ws.Context) {
		a.Send(ws.H{
			"hi": time.Now(),
		})
	})

	app.WithHttpServer(engine)
}
