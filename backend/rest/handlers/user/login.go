package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if(err != nil) {
		util.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password)
	if(err != nil) {
		util.SendError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateJWT(h.cnf.JwtSecretKey, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
	})
	if (err != nil) {
		util.SendError(w, http.StatusInternalServerError, "Error creating access token")
		return
	}

	util.SendData(w,http.StatusOK, accessToken)
}