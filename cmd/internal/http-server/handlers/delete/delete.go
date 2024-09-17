package delete

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"short_link/cmd/internal/lib/api/responce"
	sl "short_link/cmd/internal/lib/logger/slog"
	"short_link/cmd/internal/storage"
)

type Request struct {
	Alias string `json:"alias"`
}

type UrlGetter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if len(alias) == 0 {
			log.Error("alias is empty")

			render.JSON(w, r, responce.BadRequestError("alias is empty"))

			return
		}

		err := urlGetter.DeleteURL(alias)
		if errors.Is(err, storage.ErrNotFound) {
			log.Error("url not found", sl.Err(err))

			render.JSON(w, r, responce.BadRequestError(fmt.Sprintf("url with alias %s not found", alias)))
			return
		}

		if err != nil {
			log.Error("field to get URL", sl.Err(err))

			render.JSON(w, r, responce.ServerError("Internal Server Error"))

			return
		}

		log.Info("successfully deleted")

		render.JSON(w, r, http.StatusAccepted)

	}
}
