package main

import (
	"fmt"
	"net/http"
	"url-shortner/database"
	"url-shortner/handlers"

	log "github.com/drj-sharma/glogger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

const Host = "localhost:8080/"

var router *gin.Engine

func main() {
	u := uuid.New()
	fmt.Println(u[0:8])
	router := gin.Default()
	database.Migration()
	router.GET("/:code", func(ctx *gin.Context) {
		code := ctx.Param("code")
		logger.Info("code: ", code)
		payload := handlers.UrlPayload{
			Code: code,
		}
		urlHandler := handlers.UrlHandler{
			Payload: payload,
		}
		url, err := urlHandler.GetOriginalUrlByCode(ctx)
		logger.Info("url: ", url)
		logger.Info("err: ", err)
		if err != nil {
			ctx.HTML(
				http.StatusNotFound, "text/html; charset=utf-8", "Not Found",
			)
			return
		}
		// Redirect to the original mapped url
		ctx.Redirect(http.StatusMovedPermanently, url.OriginalUrl)
	})

	router.GET("/get-short-url", func(ctx *gin.Context) {
		logger.Info("params -> ", ctx.Query("url"))
		url := ctx.Query("url")
		payload := handlers.UrlPayload{
			Url: url,
		}
		urlHandler := handlers.UrlHandler{
			Payload: payload,
		}
		result, err := urlHandler.SaveUrlEntry(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "/")
		}
		logger.Info("url: ", url)
		// return shortenUrl
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"shortenUrl": Host + result.ShortUrlCode,
			},
		)
	})
	// Start serving the application
	router.Run()
}
