package request

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/p2034/HACK-2022-DREAM-DAIRY/internal/database"
)

type logout_req struct {
	Token  string `json:"token"`
	Userid int    `json:"userid"`
}

func LogoutRequest(w http.ResponseWriter, r *http.Request) {
	// check method
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	// get params
	var body logout_req
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// delete token
	db := database.OpenDB()
	defer db.Close()
	database.Token_delete(db, body.Token, body.Userid)

	w.WriteHeader(200)
}
