package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kanumone/avito_test/internal/lib/api/response"
	"github.com/kanumone/avito_test/internal/lib/helpers"
	"github.com/kanumone/avito_test/internal/storage/entities"
	"github.com/kanumone/avito_test/internal/transport/dto"
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
			response.SendError(w, response.InvalidJson)
			return
		}

		err = validate.Struct(user)
		if err != nil {
			response.ValidationError(w, err.(validator.ValidationErrors))
			return
		}

		if res, err := g.UserSlugs(user.ID); err != nil {
			response.SendError(w, err)
		} else {
			w.Write(response.OK(dto.SlugSliceToDTO(res)))
		}
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
			response.SendError(w, response.InvalidJson)
			return
		}

		if res, err := u.SlugToUser(data); err != nil {
			response.SendError(w, err)
			return
		} else {
			response.Send(w, dto.UserSlugRes{
				Added:   res.Added,
				Deleted: res.Deleted,
			})
		}
	}
}
