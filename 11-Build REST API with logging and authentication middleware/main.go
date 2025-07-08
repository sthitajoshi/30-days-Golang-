package main

import (
	"log/slog"
	"net/http"
	"os"
)

// // THESE ALL ARE FOR LEARING PROPOSE

// //The standard pattern
// func messageHandler(message string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(message))
// 	})
// }
// func main() {
// 	mux := http.NewServeMux()

// 	mux.Handle("GET/", messageHandler("hello world"))
// 	log.Print("listening on :3000...")
// 	err := http.ListenAndServe(":3000", mux)
// 	log.Fatal(err)
// }

// //Using middleware on specific routes
// func middlewareOne(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareOne")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareOne again")
// 	})
// }
// func fooHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.Path, "executing fooHandler")
// 	w.Write([]byte("ok"))
// }
// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.Path, "executing BraHandler")
// 	w.Write([]byte("bAR ok"))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/foo", fooHandler)
// 	mux.Handle("/bar", middlewareOne(http.HandlerFunc(barHandler)))

// 	log.Print("listening on :3000...")
// 	err := http.ListenAndServe(":3000", mux)
// 	log.Fatal(err)
// }

// // Chaining middleware
// func middlewareOne(next http.Handler) http.Handler { //here handler have a libary where there is define w http.ResponseWriter, r *http.Request
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareOne")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareOne again")
// 	})
// }

// func middlewareTwo(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareTwo")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareTwo again")
// 	})
// }

// func middlewareThree(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareThree")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareThree again")
// 	})
// }

// func middlewareFour(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareFour")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareFour again")
// 	})
// }

// func middlewareFive(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.URL.Path, "executing middlewareFive")
// 		next.ServeHTTP(w, r)
// 		log.Println(r.URL.Path, "executing middlewareFive again")
// 	})
// }

// func fooHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.Path, "executing fooHandler")
// 	w.Write([]byte("OK"))
// }

// func barHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.Path, "executing barHandler")
// 	w.Write([]byte("OK"))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// Apply middlewareThree and middlewareFour to GET /foo
// 	mux.Handle("GET /foo", middlewareThree(middlewareFour(http.HandlerFunc(fooHandler))))

// 	// Apply middlewareFour and middlewareFive to GET /bar
// 	mux.Handle("GET /bar", middlewareFour(middlewareFive(http.HandlerFunc(barHandler))))

// 	log.Println("listening on :3000...")
// 	// Apply middlewareOne and middlewareTwo to the entire http.ServeMux
// 	err := http.ListenAndServe(":3000", middlewareOne(middlewareTwo(mux)))
// 	log.Fatal(err)
// }



// Build REST API with logging and authentication middleware

func serverHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Go")
		next.ServeHTTP(w, r)
	})
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := slog.Group("user", "ip", ip)
		requestAttrs := slog.Group("request", "method", method, "url", url, "proto", proto)

		slog.Info("request received", userAttrs, requestAttrs)
		next.ServeHTTP(w, r)
	})
}

func requireBasicAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validUsername := "admin"
		validPassword := "secret"

		username, password, ok := r.BasicAuth()
		if !ok || username != validUsername || password != validPassword {
			w.Header().Set("WWW-Authenticate", `Basic realm="protected"`)
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func admin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin dashboard - you are authenticated!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	// Use the requireBasicAuthentication middleware on the GET /admin route only.
	mux.Handle("GET /admin", requireBasicAuthentication(http.HandlerFunc(admin)))

	slog.Info("listening on :3000...")
	// Use the serverHeader and logRequest middleware on all routes.
	err := http.ListenAndServe(":3000", serverHeader(logRequest(mux)))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
