package contacts

import (
	"github.com/bitcoin-sv/spv-wallet/engine"
	"github.com/bitcoin-sv/spv-wallet/models"
)

// CreateContact is the model for creating a contact
type CreateContact struct {
	Metadata engine.Metadata `json:"metadata"`
}

// UpdateContact is the model for updating a contact
type UpdateContact struct {
	XPubID   string               `json:"xpub_id"`
	FullName string               `json:"full_name"`
	Paymail  string               `json:"paymail"`
	PubKey   string               `json:"pubKey"`
	Status   models.ContactStatus `json:"status"`
	Metadata engine.Metadata      `json:"metadata"`
}
