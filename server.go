package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	*http.Server
}

// NewServer creates and configures a server serving all application routes.
//
// The server implements a graceful shutdown and utilizes zap.Logger for logging purposes.
// chi.Mux is used for registering some convenient middlewares and easy configuration of
// routes using different http verbs.
func NewServer(listenAddr string) (*Server, error) {

	api := newAPI()

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: api,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	return &Server{srv}, nil

}

func newAPI() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Proudly served with Go and HTTPS!\n")
	})
	mux.HandleFunc("/secret/", func(w http.ResponseWriter, req *http.Request) {
		user, pass, ok := req.BasicAuth()
		if ok && verifyUserPass(user, pass) {
			fmt.Fprintf(w, "You get to see the secret\n")
		} else {
			// i should redirect to login page
			w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})

	// register more routes over here...
	return mux
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *Server) Start() {
	fmt.Println("Starting server...")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Could not listen on %s\n", srv.Addr)
			log.Printf("%+v", err)
		}
	}()
	fmt.Println("Server is ready to handle requests")
	srv.gracefulShutdown()
}

// Start runs ListenAndServeTLS on the http.Server with graceful shutdown
func (srv *Server) StartTLS(certFile, keyFile string) {
	fmt.Println("Starting server...")

	go func() {
		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Could not listen on %s\n", srv.Addr)
			log.Printf("%+v", err)
			os.Exit(-1)
		}
	}()
	fmt.Println("Server is ready to handle requests")
	srv.gracefulShutdown()
}
func (srv *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Printf("Server is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Could not gracefuly shutdown the server", err)
	}
	fmt.Println("Server stopped")
}
