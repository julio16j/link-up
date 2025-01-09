package services

import (
    "encoding/json"
    "linkupapi/models"
    "linkupapi/repositories"
)

type ProposalService struct {
    Repo repositories.ProposalRepository
}

func (s *ProposalService) InitiateProposal(proposal models.Proposal) error {
    message, err := json.Marshal(proposal)
    if err != nil {
        return err
    }
    return s.Repo.PublishProposal(message)
}
