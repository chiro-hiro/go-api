package users

import (
	"errors"
	"net/http"

	"github.com/chiro-hiro/go-api/exports"
)

func isHexString(w http.ResponseWriter, r *http.Request) {
	w.Write(exports.ExportAPI(IsHexString(r.PostFormValue("value"))))
}

func isValidPassword(w http.ResponseWriter, r *http.Request) {
	w.Write(exports.ExportAPI(IsValidPassword(r.PostFormValue("value"))))
}

func isValidUsername(w http.ResponseWriter, r *http.Request) {
	w.Write(exports.ExportAPI(IsValidUsername(r.PostFormValue("value"))))
}

func isValidEmail(w http.ResponseWriter, r *http.Request) {
	w.Write(exports.ExportAPI(IsValidEmail(r.PostFormValue("value"))))
}

func isValidValue(w http.ResponseWriter, r *http.Request) {
	w.Write(exports.ExportAPI(IsValidValue(r.PostFormValue("value"))))
}

func getSession(r *http.Request) (session string) {
	if sessions, ok := r.Header["X-Session-Id"]; ok {
		if len(sessions) > 0 {
			session = sessions[0]
		}
	}
	return
}

func register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	//Validate username
	if valid, _ := IsValidUsername(username); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid username")))
		return
	}
	//Validate email
	if valid, _ := IsValidEmail(email); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid email")))
		return
	}
	//Validate password
	if valid, _ := IsValidPassword(password); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid password")))
		return
	}
	w.Write(exports.ExportAPI(Register(username, email, password)))
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//Validate password
	if valid, _ := IsValidPassword(password); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid password")))
		return
	}
	w.Write(exports.ExportAPI(Login(username, password)))
}

func logout(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	//Validate session
	if valid, _ := IsHexString(session); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid session")))
		return
	}
	w.Write(exports.ExportAPI(Logout(session)))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	//Validate session
	if valid, _ := IsHexString(session); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid session")))
		return
	}
	w.Write(exports.ExportAPI(GetUserBySession(session)))
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	field := r.PostFormValue("field")
	value := r.PostFormValue("value")
	//Validate session
	if valid, _ := IsHexString(session); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid session")))
		return
	}
	//Validate field name
	if valid, _ := IsValidField(field); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid field")))
		return
	}
	//Validate field value
	if valid, _ := IsValidValue(value); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid value")))
		return
	}
	w.Write(exports.ExportAPI(UpdateProfileFieldBySession(session, field, value)))
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	currentPassword := r.PostFormValue("currentPassword")
	newPassword := r.PostFormValue("newPassword")
	session := getSession(r)
	//Validate current password
	if valid, _ := IsValidPassword(currentPassword); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid current password")))
		return
	}
	//Validate new password
	if valid, _ := IsValidPassword(newPassword); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid new password")))
		return
	}
	w.Write(exports.ExportAPI(UpdatePassword(session, currentPassword, newPassword)))
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	session := getSession(r)
	//Validate session
	if valid, _ := IsHexString(session); !valid {
		w.Write(exports.ExportAPI(valid, errors.New("Invalid session")))
		return
	}
	w.Write(exports.ExportAPI(GetUserProfileBySession(session)))
}
