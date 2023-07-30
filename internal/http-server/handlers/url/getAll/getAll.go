package getAll

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"
)

type Urls struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type Response struct {
	response.Response
	Urls []Urls `json:"urls"`
}

type UrlsGetter interface {
	GetUrls() ([]Urls, error)
}

func New(log *slog.Logger, urlsGetter UrlsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const operation = "handlers.url.get-all.New"

		log = log.With(
			slog.String("operation", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		urls, err := urlsGetter.GetUrls()

		if errors.Is(err, storage.ErrorEmptyUrls) {
			log.Info("no one existing alias")

			render.JSON(w, r, response.Error("no one existing alias"))

			return
		}

		render.JSON(w, r, Response{
			Response: response.OK(),
			Urls:     urls,
		})
	}
}
