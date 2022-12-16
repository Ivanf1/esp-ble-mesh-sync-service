package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Ivanf1/esp-ble-mesh-sync-service/pkg/db"
	"github.com/jmoiron/sqlx/types"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc(os.Getenv("API_MESH_CONFIGURATION_BASE"), s.configurationRouter)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) configurationRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetConfiguration(w, r)
	case "POST":
		s.handleUpdateConfiguration(w, r)
	}
}

func (s *Server) handleGetConfiguration(w http.ResponseWriter, r *http.Request) {
	log.Println("GET configuration")

	httpData := db.BleMeshConfigurationHTTP{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &httpData)

	result := db.BleMeshConfigurationDB{}

	db.GetConfiguration(httpData.ID, &result)

	j, err := json.Marshal(result)
	if err != nil {
		log.Fatal("json marshal error: ", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (s *Server) handleUpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	log.Println("POST configuration")

	httpData := db.BleMeshConfigurationHTTP{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &httpData)

	// remove space characters
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, httpData.Config)
	if err != nil {
		log.Fatal("json compact error: ", err)
	}

	updatedConfiguration := db.BleMeshConfigurationDB{}
	updatedConfiguration.ID = httpData.ID
	updatedConfiguration.Config = types.JSONText(buffer.String())

	db.UpdateConfiguration(&updatedConfiguration)
}
