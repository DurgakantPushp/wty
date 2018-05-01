package api

import (
	"encoding/json"
	"net/http"

	"github.com/wty/auth"
	"github.com/wty/models"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
}

// UserSignup -
func (api *API) UserSignup(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil ||
		jsondata.UserName == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(jsondata.UserName) {
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(jsondata.UserName, jsondata.Password, jsondata.Role)

	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))
}

// UserLogin -
func (api *API) UserLogin(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.UserName == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	user := api.users.FindUser(jsondata.UserName)
	if user.Username == "" {
		http.Error(w, "username not found", http.StatusBadRequest)
		return
	}

	if !api.users.CheckPassword(user.Password, jsondata.Password) {
		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsontoken))

}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {

	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
