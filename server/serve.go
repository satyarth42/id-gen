package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/satyarth42/id-gen/config"
	"github.com/satyarth42/id-gen/logic"
)

func HandleGenerateID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	id, err := logic.GenerateID()

	if err != nil {
		log.Printf("error returned in generating ID, err: %+v", err)

		http.Error(w, fmt.Sprintf("{\"error\":%s}", err.Error()), http.StatusTooManyRequests)
		return
	}

	resp := struct {
		ID int64 `json:"id"`
	}{
		ID: id,
	}

	respBytes, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func Serve() {
	config.LoadConfig()

	router := mux.NewRouter()

	router.HandleFunc("/id", HandleGenerateID).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.GetConfig().Port),
		Handler: router,

		WriteTimeout: time.Second,
		ReadTimeout:  100 * time.Millisecond,
	}

	log.Printf("server starting at %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
