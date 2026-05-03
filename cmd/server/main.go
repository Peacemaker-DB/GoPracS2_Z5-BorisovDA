package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"example.com/gopracs2-z5-borisovda/internal/config"
	"example.com/gopracs2-z5-borisovda/internal/httpapi"
	"example.com/gopracs2-z5-borisovda/internal/student"
)

func main() {
	cfg := config.New()

	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo, err := student.NewRepo(db)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	handler := httpapi.NewHandler(repo)

	mux := http.NewServeMux()
	httpapi.RegisterRoutes(mux, handler)

	startRedirectServer(cfg.HTTPAddr, cfg.HTTPSHost)

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
		},
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("HTTPS server started on https://" + cfg.HTTPSHost)

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func startRedirectServer(addr string, httpsHost string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := "https://" + httpsHost + r.URL.RequestURI()
		http.Redirect(w, r, target, http.StatusMovedPermanently)
	})

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("HTTP redirect server started on http://localhost" + addr)

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}