package handlers

import (
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
		const op = "server.handlers.user.UserSlugs"
		user := dto.User{}
		err := helpers.ParseJson(r.Body, &user)
		if err != nil {
			helpers.LogErr(op, err)
			response.Error("invalid json")
		}
		res, err := g.UserSlugs(user.ID)
		if err != nil {
			helpers.LogErr(op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response.OK(dto.SlugSliceToDTO(res)))
	}
}

type Updator interface {
	SlugToUser(data dto.UserSlugReq) (entities.AddedDeleted, error)
}

func UpdateUser(u Updator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "server.handlers.user.UpdateUser"
		data := dto.UserSlugReq{}
		err := helpers.ParseJson(r.Body, &data)
		if err != nil {
			helpers.LogErr(op, err)
			response.SendError(w, err)
			return
		}
		res, err := u.SlugToUser(data)
		if err != nil {
			response.SendError(w, err)
			return
		}
		response.Send(w, dto.UserSlugRes{
			Added:   res.Added,
			Deleted: res.Deleted,
		})
	}
}
