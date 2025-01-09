package controllers

import (
    "encoding/json"
    "net/http"
    "linkupapi/models"
    "linkupapi/services"
)

var proposalService = services.ProposalService{}

func CreateProposal(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var proposal models.Proposal
    if err := json.NewDecoder(r.Body).Decode(&proposal); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err := proposalService.InitiateProposal(proposal)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"status": "Proposal initiated"})
}
