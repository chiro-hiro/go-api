package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"

	"github.com/chiro-hiro/go-api/users"
	_ "github.com/go-sql-driver/mysql"
)

//HTTP redirect to HTTPS
func httpRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Location", "https://localhost")
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte("HTTP isn't allowed."))
}

//Accept cross site API call
func acceptCrossSiteAPI(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptedOrigin := map[string]bool{"http://localhost:8080": true}
		origin := ""
		if origins, ok := r.Header["Origin"]; ok {
			if len(origins) > 0 {
				origin = origins[0]
			}
		}
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "x-session-id")
			w.Header().Set("Access-Control-Allow-Method", "POST")
			if origin != "" && acceptedOrigin[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.Write([]byte(""))
			return
		}

		if r.Method == "POST" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "x-session-id")
			if origin != "" && acceptedOrigin[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			handler.ServeHTTP(w, r)
		}

		log.Println("Handler finished processing request")
	})
}

func main() {

	//Open database connect
	db, err := sql.Open("mysql", "[username]:[password]/[database]")
	if err != nil {
		panic(err)
	}

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	//Create new mux server
	muxServer := http.NewServeMux()

	//Create new mux server
	redirectMux := http.NewServeMux()
	redirectMux.HandleFunc("/", httpRedirect)

	//Register users package to use database
	users.InitDB(db)

	//Mux server hook
	users.InitMux(muxServer)

	//Wrapp cross site API
	wrappedMux := acceptCrossSiteAPI(muxServer)

	srv := &http.Server{
		Addr:         "0.0.0.0:443",
		Handler:      wrappedMux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	//Start redirect server
	go func() {
		if err := http.ListenAndServe("0.0.0.0:80", redirectMux); err != nil {
			panic(err)
		} else {
			log.Println("Start redirect server")
		}
	}()

	//Start TLS server
	if err := srv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		panic(err)
	} else {
		log.Println("Start HTTPS server")
	}

}
