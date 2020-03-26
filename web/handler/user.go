package handler

import (
	"encoding/json"
	"github.com/devgek/webskeleton/config"
	"github.com/devgek/webskeleton/web"
	"github.com/devgek/webskeleton/web/viewmodel"
	"log"
	"net/http"
	"strconv"
)

//HandleUsers ...
func HandleUsers(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//show user list
		contextData := web.NewContextData()
		ctx := web.ToContext(r.Context(), contextData)

		users, err := env.Services.GetAllUsers()
		viewData := web.NewViewDataWithContextData(web.FromContext(r.Context()))
		viewData["Users"] = users
		viewData["EditEntityType"] = env.MessageLocator.GetString("entity.user")
		if err != nil {
			viewData["ErrorMessage"] = err.Error()
		}
		web.RenderTemplate(w, r.WithContext(ctx), "users.html", viewData)
		return
	})

}

//HandleUserEdit ...
func HandleUserEdit(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oName := r.FormValue("gkvName")
		oEmail := r.FormValue("gkvEmail")
		oAdmin := r.FormValue("gkvAdmin")
		log.Println(oAdmin)

		contextData := web.NewContextData()
		web.ToContext(r.Context(), contextData)

		u, err := env.Services.UpdateUser(oName, oEmail, oAdmin == "true")

		vd := web.NewViewData()
		userEditResponse := viewmodel.NewUserEditResponse()
		if err != nil {
			userEditResponse.IsError = true
			userEditResponse.Message = env.MessageLocator.GetString("msg.error.user.edit")
			userEditResponse.Name = oName
			userEditResponse.Email = oEmail
			userEditResponse.Admin = (oAdmin == "true")
		} else {
			userEditResponse.Message = env.MessageLocator.GetString("msg.success.user.edit")
			userEditResponse.Name = u.Name
			userEditResponse.Email = u.Email
			userEditResponse.Admin = u.Admin
		}

		vd["Response"] = userEditResponse
		json.NewEncoder(w).Encode(vd)
		return
	})

}

//HandleUserNew ...
func HandleUserNew(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oName := r.FormValue("gkvName")
		oPass := r.FormValue("gkvPass")
		oEmail := r.FormValue("gkvEmail")
		oAdmin := r.FormValue("gkvAdmin")
		log.Println(oAdmin)

		contextData := web.NewContextData()
		web.ToContext(r.Context(), contextData)

		u, err := env.Services.CreateUser(oName, oPass, oEmail, oAdmin == "true")

		vd := web.NewViewData()
		userEditResponse := viewmodel.NewUserEditResponse()
		if err != nil {
			userEditResponse.IsError = true
			userEditResponse.Message = env.MessageLocator.GetString("msg.error.user.create")
			userEditResponse.Name = oName
			userEditResponse.Pass = oPass
			userEditResponse.Email = oEmail
			userEditResponse.Admin = (oAdmin == "true")
		} else {
			userEditResponse.Message = env.MessageLocator.GetString("msg.success.user.create")
			userEditResponse.Name = u.Name
			userEditResponse.Pass = string(u.Pass)
			userEditResponse.Email = u.Email
			userEditResponse.Admin = u.Admin
		}

		vd["Response"] = userEditResponse
		json.NewEncoder(w).Encode(vd)
		return
	})

}

//HandleUserDelete ...
func HandleUserDelete(env *config.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oID := r.FormValue("gkvObjId")
		ioID, _ := strconv.Atoi(oID)

		contextData := web.NewContextData()
		web.ToContext(r.Context(), contextData)

		err := env.Services.DeleteUser(uint(ioID))

		vd := web.NewViewData()
		baseResponse := &viewmodel.BaseResponse{}
		if err != nil {
			baseResponse.IsError = true
			baseResponse.Message = env.MessageLocator.GetString("msg.error.user.delete")
		} else {
			baseResponse.Message = env.MessageLocator.GetString("msg.success.user.delete")
		}

		vd["Response"] = baseResponse
		json.NewEncoder(w).Encode(vd)
		return
	})

}
