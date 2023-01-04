package api

import (
	"encoding/json"

	"github.com/anthdm/weavebox"
)

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationHandler struct {
	// userStore
}

func (h *AuthenticationHandler) AuthenticateUser(ctx *weavebox.Context) error {
	authReq := &AuthenticationRequest{}
	if err := json.NewDecoder(ctx.Request().Body).Decode(authReq); err != nil {
		return err
	}
	return nil
}
