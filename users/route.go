package users

import (
	"net/http"
)

//InitMux to share mux server with main process
func InitMux(muxServer *http.ServeMux) {
	handlers := map[string]func(w http.ResponseWriter, r *http.Request){
		//Validation
		"/validation/isHexString":     isHexString,
		"/validation/isValidPassword": isValidPassword,
		"/validation/isValidUsername": isValidUsername,
		"/validation/isValidEmail":    isValidEmail,
		"/validation/isValidValue":    isValidValue,
		//User APIs
		"/user/register":       register,
		"/user/login":          login,
		"/user/logout":         logout,
		"/user/getUser":        getUser,
		"/user/updateProfile":  updateProfile,
		"/user/updatePassword": updatePassword,
		"/user/getProfile":     getProfile}

	//Append to mux server
	for key, value := range handlers {
		muxServer.HandleFunc(key, value)
	}
}
