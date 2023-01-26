package handlers

import (
	"context"
	"url-shortner/database"
	"url-shortner/database/entity"
	u "url-shortner/pkg/uuid"
	"url-shortner/utils"

	log "github.com/drj-sharma/glogger"
)

var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

type UrlPayload struct {
	Url  string
	Code string
}

type UrlHandler struct {
	Payload UrlPayload
}

// get mapped url by using code of shortened url
func (urlHandler *UrlHandler) GetOriginalUrlByCode(ctx context.Context) (entity.Url, error) {
	code := urlHandler.Payload.Code
	database := database.Database{}
	connection, err := database.Connect()
	url := entity.Url{
		ShortUrlCode: code,
	}
	if err != nil {
		logger.Debug("Not able to fetch url with given short code", err)
		return url, err
	}
	var result entity.Url
	connection.Table("urls").Select("original_url, short_url_code").Where("short_url_code = ?", url.ShortUrlCode).Scan(&result)
	return result, nil
}

// save original url as well as new shortened url code in database
func (urlHandler *UrlHandler) SaveUrlEntry(ctx context.Context) (entity.Url, error) {
	uuid := u.UUID{}
	code := uuid.GetCode()
	database := database.Database{}
	connection, err := database.Connect()
	utils.IsValidUrl(urlHandler.Payload.Url)
	url := entity.Url{
		OriginalUrl:  urlHandler.Payload.Url,
		ShortUrlCode: code,
	}
	if err != nil {
		logger.Debug("Not able to connect with database", err)
		return url, err
	}
	result := connection.Create(&url)
	if result.Error != nil {
		logger.Debug("not able to create short url entry", result.Error)
		return url, result.Error
	}
	return url, nil
}
