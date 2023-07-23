package save

import (
	"errors"
	"net/http"
	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/lib/random"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

const aliasLength = 6

//go:generate go run github.com/vektra/mockery/v2@v2.32.0 --name UrlSaver
type UrlSaver interface {
	SaveUrl(urlToSave string, alias string) (int64, error)
}

func New(logger *slog.Logger, urlSaver UrlSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const operation = "handlers.url.save.New"

		logger = logger.With(
			slog.String("operation", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var request Request

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			logger.Error("Failed to decode request body", sl.Err(err))

			render.JSON(w, r, response.Error("Failed to decode request"))

			return
		}

		logger.Info("request body decoded", slog.Any("request", request))

		if err := validator.New().Struct(request); err != nil {
			validateErr := err.(validator.ValidationErrors)
			logger.Error("Invalid request", sl.Err(err))

			render.JSON(w, r, response.Error("Invalid request"))
			render.JSON(w, r, response.ValidationError(validateErr))

			return
		}

		alias := request.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

		id, err := urlSaver.SaveUrl(request.URL, alias)
		if errors.Is(err, storage.ErrorUrlExists) {
			logger.Info("url already exists", slog.String("url", request.URL))

			render.JSON(w, r, response.Error("url already exists"))

			return
		}
		if err != nil {
			logger.Error("failed to add url", sl.Err(err))

			render.JSON(w, r, response.Error("failed to add url"))

			return
		}

		logger.Info("url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
