package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func pad(hexString string) string {
	if len(hexString)%2 == 1 {
		return "0" + hexString
	}
	return hexString
}

func main() {
	a := map[string]bool{"hello": true}
	fmt.Println(a["hello"])
}

/*

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/api/", test)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

func main() {
	/*
			fmt.Println(isVailidUsername("chi"))
			fmt.Println(isVailidUsername("chiro"))
			fmt.Println(isVailidUsername("8chiro"))
			fmt.Println(isVailidUsername("8chiro_"))
			fmt.Println(isHexString("aff123"))
			fmt.Println(isHexString("affF123"))

			u := User{ID: 123, Username: "Chiro"}

			fmt.Println(u.ChangePassword("dsadsadsadsa23232"))
			fmt.Println(u.password)
			fmt.Printf("%x", u.salt)

		v := make([]byte, 20)
		fmt.Println(v)

		u := User{ID: []byte("Hello"), Username: "chiro"}

		fmt.Println(u.ChangePassword("dsadsadsadsa23232"))
		fmt.Printf("%x\n", u.password)
		fmt.Printf("%x\n", u.salt)

		u.NewSession()
		u.NewSession()
		u.NewSession()
		u.NewSession()
		fmt.Println("Session:", u.sessions)
		fmt.Println(u.Login("dsadsadsadsa23232"))
		fmt.Println("Session:", u.sessions)

	db, _ := sql.Open("mysql", "root:!thepirate@/test")
	type data struct {
		id       []byte
		username string
		password []byte
		salt     []byte
		sessions []byte
	}

	a := data{}
	db.QueryRow("SELECT * FROM `users` WHERE `username`=?;", "chiro").Scan(&a.id, &a.username, &a.password, &a.salt, &a.sessions)

	fmt.Println(a)
	db.Close()

}	*/
