package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/server/dto"
	"github.com/kanumone/avito_test/internal/storage/entities"
)

type Getter interface {
	UserSlugs(userID int64) ([]entities.Slug, error)
}

func UserSlugs(g Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Get"
		_, err := w.Write([]byte("User slugs"))
		if err != nil {
			helpers.LogErr(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

type Updator interface {
	SlugToUser(data dto.UserSlug) error
}

func UpdateUser(u Updator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.slugs.Put"
		data := dto.UserSlug{}
		err := helpers.ParseJson(r.Body, &data)
		if err != nil {
			helpers.LogErr(op, err)
			w.WriteHeader(http.StatusBadRequest)
			response.Error("invalid json")
			return
		}
		log.Printf("data: %v\n", data)
		fmt.Fprint(w, "OK")
		// ok := u.SlugToUser(data)
		// if ok {
		// 	w.Write(response.OK("updated successfully"))
		// } else {
		// 	w.Write(response.Error("something went wrong"))
		// }
	}
}
