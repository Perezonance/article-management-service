package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Perezonance/article-management-service/internal/controllers"
	"github.com/Perezonance/article-management-service/internal/server"
	"github.com/Perezonance/article-management-service/internal/storage"
	l "github.com/Perezonance/article-management-service/internal/util/logger"
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	r := mux.NewRouter()

	db := storage.NewMockDynamo()

	s := server.NewServer(db)

	c := controllers.NewController(s)

	r.HandleFunc("/articles", c.GetArticlesHandler).Methods(http.MethodGet)
	r.HandleFunc("/articles", c.PostArticleHandler).Methods(http.MethodPost)

	r.HandleFunc("/articles/{articleID}", c.GetArticleByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/articles/{articleID}", c.UpdateArticleByIDHandler).Methods(http.MethodPut)
	r.HandleFunc("/articles/{articleID}", c.DeleteArticleByIDHandler).Methods(http.MethodDelete)

	r.HandleFunc("/articles/{userID}", c.GetArticleByUserIDHandler).Methods(http.MethodGet)

	//GET /articles?id=id1, id2, idn...
	// r.HandleFunc("/articles", c.GetMultArticleByIDHandler).Methods(http.MethodGet).Queries("ids")

	l.InfoLog("Server Initialized")
	//Graceful shut down procedure...
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			l.ErrorLog("Server encountered error while serving", err)
		}
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	//Block until signal emitted
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	os.Exit(0)
}
