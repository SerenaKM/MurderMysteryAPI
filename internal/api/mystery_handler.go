package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/serenakm/MurderMysteryAPI/internal/store"
	"github.com/serenakm/MurderMysteryAPI/internal/utils"
)

// handling data
type MysteryHandler struct {
	mysteryStore store.MysteryStore
	logger *log.Logger
}

// constructor
func NewMysteryHandler(mysteryStore store.MysteryStore, logger *log.Logger) *MysteryHandler {
	return &MysteryHandler{
		mysteryStore: mysteryStore,
		logger: logger,
	}
}

func (mh *MysteryHandler) HandleGetMysteryByID(w http.ResponseWriter, r *http.Request) {
	mysteryID, err := utils.ReadIDParam(r)
	if err != nil {
		mh.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"Error":"Invalid case id"})
		return
	}

	mysteryCase, err := mh.mysteryStore.GetCaseByID(mysteryID)
	if err != nil {
		mh.logger.Printf("ERROR: getCaseByID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error":"Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"Case": mysteryCase})
}

func (mh *MysteryHandler) HandleCreateMystery(w http.ResponseWriter, r *http.Request) {
	var mysteryCase store.Case
	err := json.NewDecoder(r.Body).Decode(&mysteryCase) // parse the JSON from the request body into a struct
	if err != nil {
		mh.logger.Printf("ERROR: decodingCreateMystery: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"Error":"Invalid request sent"})
		return
	}

	createdCase, err := mh.mysteryStore.CreateCase(&mysteryCase)
	if err != nil {
		mh.logger.Printf("ERROR: createCase: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error":"Failed to create case"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"Case": createdCase})
}

func (mh *MysteryHandler) HandleDeleteCase (w http.ResponseWriter, r *http.Request) {
	mysteryID, err := utils.ReadIDParam(r)
	if err != nil {
		mh.logger.Printf("ERROR: readIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"Error":"Invalid case delete id"})
		return
	}

	err = mh.mysteryStore.DeleteCase(mysteryID)
	if err == sql.ErrNoRows {
		mh.logger.Printf("ERROR: deletingCase: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"Error":"Case not found"})
		return
	}

	if err != nil {
		mh.logger.Printf("ERROR: deletingCase: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"Error":"Error deleting case"})
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, utils.Envelope{"Case":"Deleted"})
}