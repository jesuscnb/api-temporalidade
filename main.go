package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/jesuscnb/api-temporalidade/config"
	. "github.com/jesuscnb/api-temporalidade/dao"
	. "github.com/jesuscnb/api-temporalidade/models"
)

var config = Config{}
var dao = TemporalidadesDAO{}

// GET list of Temporalidades
func AllTemporalidadesEndPoint(w http.ResponseWriter, r *http.Request) {
	temporalidades, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, temporalidades)
}

// GET a Temporalidade by its ID
func FindTemporalidadeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	temporalidade, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Temporalidade ID")
		return
	}
	respondWithJson(w, http.StatusOK, temporalidade)
}

// POST a new Temporalidade
func CreateTemporalidadeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var temporalidade Temporalidade
	if err := json.NewDecoder(r.Body).Decode(&temporalidade); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	temporalidade.ID = bson.NewObjectId()
	if err := dao.Insert(temporalidade); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, temporalidade)
}

// PUT update an existing Temporalidade
func UpdateTemporalidadeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var temporalidade Temporalidade
	if err := json.NewDecoder(r.Body).Decode(&temporalidade); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(temporalidade); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing Temporalidade
func DeleteTemporalidadeEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var temporalidade Temporalidade
	if err := json.NewDecoder(r.Body).Decode(&temporalidade); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(temporalidade); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/temporalidades", AllTemporalidadesEndPoint).Methods("GET")
	r.HandleFunc("/temporalidades", CreateTemporalidadeEndPoint).Methods("POST")
	r.HandleFunc("/temporalidades", UpdateTemporalidadeEndPoint).Methods("PUT")
	r.HandleFunc("/temporalidades/{id}", DeleteTemporalidadeEndPoint).Methods("DELETE")
	r.HandleFunc("/temporalidades/{id}", FindTemporalidadeEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
