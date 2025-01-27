package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/keratin/authn-server/app"
	"github.com/keratin/authn-server/app/services"
)

func GetAccount(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paramID *int
		var paramUsername *string

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err == nil {
			paramID = &id
		}

		username := mux.Vars(r)["username"]
		if username != "" {
			paramUsername = &username
		}

		if paramID == nil && paramUsername == nil {
			WriteNotFound(w, "account")
			return
		}

		account, err := services.AccountGetter(app.AccountStore, services.AccountGetterParams{
			AccountID: paramID,
			Username:  paramUsername,
		})
		if err != nil {
			if _, ok := err.(services.FieldErrors); ok {
				WriteNotFound(w, "account")
				return
			}

			panic(err)
		}

		WriteData(w, http.StatusOK, account)
	}
}
