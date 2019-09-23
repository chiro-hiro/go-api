package users

import (
	"database/sql"
	"encoding/hex"
)

var db *sql.DB

//InitDB database
func InitDB(_db *sql.DB) {
	db = _db
}

//User structure for basic user account
type User struct {
	ID       string
	Username string
	Email    string
}

//Register user register
func Register(username string, email string, password string) (result bool, err error) {
	err = db.QueryRow("SELECT createNewUser(?, ?, UNHEX(?))", username, email, password).Scan(&result)
	return
}

//Login user login
func Login(username string, password string) (result string, err error) {
	var sessionID []byte
	err = db.QueryRow("SELECT loginUser(?, UNHEX(?))", username, password).Scan(&sessionID)
	result = hex.EncodeToString(sessionID)
	return
}

//Logout user logout
func Logout(sessionID string) (result bool, err error) {
	err = db.QueryRow("SELECT logoutUser(UNHEX(?))", sessionID).Scan(&result)
	return
}

//GetUserBySession get user basic profile by session ID
func GetUserBySession(sessionID string) (result User, err error) {
	var id []byte
	err = db.QueryRow("CALL getUserBySession(UNHEX(?))", sessionID).Scan(&id, &result.Username, &result.Email)
	if err == nil {
		result.ID = hex.EncodeToString(id)
	}
	return
}

//UpdateProfile update an user profile field
func UpdateProfile(userID string, field string, value string) (result bool, err error) {
	err = db.QueryRow("SELECT updateProfileField(UNHEX(?), ?, ?)", userID, field, value).Scan(&result)
	return
}

//UpdatePassword update an user profile field
func UpdatePassword(sessionID string, currentPassword string, newPassword string) (result bool, err error) {
	err = db.QueryRow("SELECT updatePassword(UNHEX(?), UNHEX(?), UNHEX(?))", sessionID, currentPassword, newPassword).Scan(&result)
	return
}

//UpdateProfileFieldBySession update an user profile field
func UpdateProfileFieldBySession(sessionID string, field string, value string) (result bool, err error) {
	err = db.QueryRow("SELECT updateProfileFieldBySession(UNHEX(?), ?, ?)", sessionID, field, value).Scan(&result)
	return
}

//GetUserProfileBySession get user profile info by given session ID
func GetUserProfileBySession(sessionID string) (result map[string]string, err error) {
	rows, err := db.Query("CALL getProfileBySession(UNHEX(?))", sessionID)
	if err == nil {
		tmp := make(map[string]string)
		for rows.Next() == true {
			var id []byte
			var field string
			var value string
			rows.Scan(&id, &field, &value)
			if len(tmp["userID"]) == 0 {
				tmp["userID"] = hex.EncodeToString(id)
			}
			tmp[field] = value
		}
		result = tmp
	}
	err = rows.Close()
	return
}
