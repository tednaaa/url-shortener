package redirect

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"
)

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name=UrlGetter
type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, response.Error("invalid request"))

			return
		}

		resURL, err := urlGetter.GetUrl(alias)
		if errors.Is(err, storage.ErrorUrlNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, response.Error("not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			render.JSON(w, r, response.Error("internal error"))

			return
		}

		log.Info("got url", slog.String("url", resURL))

		http.Redirect(w, r, resURL, http.StatusFound)
	}
}