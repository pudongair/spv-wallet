package mappings

import (
	"github.com/bitcoin-sv/spv-wallet/engine"
	"github.com/bitcoin-sv/spv-wallet/mappings/common"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

// MapToXpubContract will map the xpub model from spv-wallet to the spv-wallet-models contract
func MapToXpubContract(xpub *engine.Xpub) *response.Xpub {
	if xpub == nil {
		return nil
	}

	return &response.Xpub{
		Model:           *common.MapToContract(&xpub.Model),
		ID:              xpub.ID,
		CurrentBalance:  xpub.CurrentBalance,
		NextInternalNum: xpub.NextInternalNum,
		NextExternalNum: xpub.NextExternalNum,
	}
}
