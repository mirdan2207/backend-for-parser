package samples

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	h := NewSampleHandler(NewSampleService(db))

	r.HandleFunc("/samples/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		h.GetSamplesAll(w, r)
	}).Methods("GET")
	r.HandleFunc("/samples/{user_id}/{sample_id}", func(w http.ResponseWriter, r *http.Request) {
		h.GetSampleByID(w, r)
	}).Methods("GET")
	r.HandleFunc("/samples/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		h.CreateSample(w, r)
	}).Methods("POST")
	r.HandleFunc("/samples/{user_id}/{sample_id}", func(w http.ResponseWriter, r *http.Request) {
		h.UpdateSample(w, r)
	}).Methods("PUT")
	r.HandleFunc("/samples/{user_id}/{sample_id}", func(w http.ResponseWriter, r *http.Request) {
		h.DeleteSample(w, r)
	}).Methods("DELETE")
}

type SampleHandler struct {
	service *SampleService
}

func NewSampleHandler(service *SampleService) *SampleHandler {
    return &SampleHandler{service: service}
}

func (h *SampleHandler) GetSamplesAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err!= nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

	samples, err := h.service.GetSamplesAllSamples(userID)
	if err != nil {
        http.Error(w, "Failed to get samples", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(samples)
}

func (h *SampleHandler) GetSampleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["user_id"])
    if err!= nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    sampleID, err := strconv.Atoi(vars["sample_id"])
    if err!= nil {
        http.Error(w, "Invalid sample ID", http.StatusBadRequest)
        return
    }

    sample, err := h.service.GetSampleByID(userID, sampleID)
    if err!= nil {
        http.Error(w, "Failed to get sample", http.StatusInternalServerError)
        return
    }

	if sample == nil {
		http.Error(w, "No sample found", http.StatusNotFound)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(sample)
}

func (h *SampleHandler) CreateSample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err!= nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

	var sample Sample
	if err := json.NewDecoder(r.Body).Decode(&sample); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateSample(userID, &sample); err != nil {
		http.Error(w, "Failed to create Sample", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sample)
}

func (h *SampleHandler) UpdateSample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["user_id"])
    if err!= nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    sampleID, err := strconv.Atoi(vars["sample_id"])
    if err!= nil {
        http.Error(w, "Invalid sample ID", http.StatusBadRequest)
        return
    }

    var sample Sample
    if err := json.NewDecoder(r.Body).Decode(&sample); err!= nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.service.UpdateSample(userID, sampleID, &sample); err!= nil {
        http.Error(w, "Failed to update sample", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sample)
}

func (h *SampleHandler) DeleteSample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["user_id"])
    if err!= nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    sampleID, err := strconv.Atoi(vars["sample_id"])
    if err!= nil {
        http.Error(w, "Invalid sample ID", http.StatusBadRequest)
        return
    }

    if err := h.service.DeleteSample(userID, sampleID); err!= nil {
        http.Error(w, "Failed to delete sample", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
    json.NewEncoder(w).Encode(map[string]string{"message": "Sample deleted"})
}