package mappings

import (
	"github.com/bitcoin-sv/spv-wallet/engine"
	"github.com/bitcoin-sv/spv-wallet/mappings/common"
	spvwalletmodels "github.com/bitcoin-sv/spv-wallet/models"
)

// MapToTransactionContract will map the model from spv-wallet to the spv-wallet-models contract
func MapToTransactionContract(t *engine.Transaction) *spvwalletmodels.Transaction {
	if t == nil {
		return nil
	}

	model := spvwalletmodels.Transaction{
		Model:                *common.MapToContract(&t.Model),
		ID:                   t.ID,
		Hex:                  t.Hex,
		XpubInIDs:            t.XpubInIDs,
		XpubOutIDs:           t.XpubOutIDs,
		BlockHash:            t.BlockHash,
		BlockHeight:          t.BlockHeight,
		Fee:                  t.Fee,
		NumberOfInputs:       t.NumberOfInputs,
		NumberOfOutputs:      t.NumberOfOutputs,
		DraftID:              t.DraftID,
		TotalValue:           t.TotalValue,
		Status:               string(t.Status),
		TransactionDirection: string(t.Direction),
	}

	processMetadata(t, t.XPubID, &model)
	processOutputValue(t, t.XPubID, &model)

	return &model
}

// MapToTransactionContractForAdmin will map the model from spv-wallet to the spv-wallet-models contract for admin
func MapToTransactionContractForAdmin(t *engine.Transaction) *spvwalletmodels.Transaction {
	if t == nil {
		return nil
	}

	model := spvwalletmodels.Transaction{
		Model:           *common.MapToContract(&t.Model),
		ID:              t.ID,
		Hex:             t.Hex,
		XpubInIDs:       t.XpubInIDs,
		XpubOutIDs:      t.XpubOutIDs,
		BlockHash:       t.BlockHash,
		BlockHeight:     t.BlockHeight,
		Fee:             t.Fee,
		NumberOfInputs:  t.NumberOfInputs,
		NumberOfOutputs: t.NumberOfOutputs,
		DraftID:         t.DraftID,
		TotalValue:      t.TotalValue,
		Status:          string(t.Status),
		Outputs:         t.XpubOutputValue,
	}

	processMetadata(t, t.XPubID, &model)

	return &model
}

func processMetadata(t *engine.Transaction, xpubID string, model *spvwalletmodels.Transaction) {
	if len(t.XpubMetadata) > 0 && len(t.XpubMetadata[xpubID]) > 0 {
		if t.Model.Metadata == nil {
			model.Model.Metadata = make(spvwalletmodels.Metadata)
		}
		for key, value := range t.XpubMetadata[xpubID] {
			model.Model.Metadata[key] = value
		}
	}
}

func processOutputValue(t *engine.Transaction, xpubID string, model *spvwalletmodels.Transaction) {
	model.OutputValue = int64(0)
	if len(t.XpubOutputValue) > 0 && t.XpubOutputValue[xpubID] != 0 {
		model.OutputValue = t.XpubOutputValue[xpubID]
	}

	if model.OutputValue > 0 {
		model.TransactionDirection = "incoming"
	} else {
		model.TransactionDirection = "outgoing"
	}
}

// MapToTransactionSPV will map the model from spv-wallet-models to the spv-wallet contract
func MapToTransactionSPV(t *spvwalletmodels.Transaction) *engine.Transaction {
	if t == nil {
		return nil
	}

	return &engine.Transaction{
		Model:           *common.MapToModel(&t.Model),
		TransactionBase: engine.TransactionBase{ID: t.ID, Hex: t.Hex},
		XpubInIDs:       t.XpubInIDs,
		XpubOutIDs:      t.XpubOutIDs,
		BlockHash:       t.BlockHash,
		BlockHeight:     t.BlockHeight,
		Fee:             t.Fee,
		NumberOfInputs:  t.NumberOfInputs,
		NumberOfOutputs: t.NumberOfOutputs,
		DraftID:         t.DraftID,
		TotalValue:      t.TotalValue,
		OutputValue:     t.OutputValue,
		Status:          engine.SyncStatus(t.Status),
		Direction:       engine.TransactionDirection(t.TransactionDirection),
	}
}

// MapToTransactionConfigSPV will map the transaction-config model from spv-wallet to the spv-wallet-models contract
func MapToTransactionConfigSPV(tx *spvwalletmodels.TransactionConfig) *engine.TransactionConfig {
	if tx == nil {
		return nil
	}

	return &engine.TransactionConfig{
		ChangeDestinations:         mapToSPVDestinations(tx),
		ChangeDestinationsStrategy: engine.ChangeStrategy(tx.ChangeStrategy),
		ChangeMinimumSatoshis:      tx.ChangeMinimumSatoshis,
		ChangeNumberOfDestinations: tx.ChangeNumberOfDestinations,
		ChangeSatoshis:             tx.ChangeSatoshis,
		ExpiresIn:                  tx.ExpiresIn,
		Fee:                        tx.Fee,
		FeeUnit:                    MapToFeeUnitSPV(tx.FeeUnit),
		FromUtxos:                  mapToSPVFromUtxos(tx),
		IncludeUtxos:               mapToSPVIncludeUtxos(tx),
		Inputs:                     mapToSPVInputs(tx),
		Outputs:                    mapToSPVOutputs(tx),
		SendAllTo:                  MapToTransactionOutputSPV(tx.SendAllTo),
		Sync:                       MapToSyncConfigSPV(tx.Sync),
	}
}

func mapToSPVOutputs(tx *spvwalletmodels.TransactionConfig) []*engine.TransactionOutput {
	if tx.Outputs == nil {
		return nil
	}

	outputs := make([]*engine.TransactionOutput, 0)
	for _, output := range tx.Outputs {
		outputs = append(outputs, MapToTransactionOutputSPV(output))
	}
	return outputs
}

func mapToSPVInputs(tx *spvwalletmodels.TransactionConfig) []*engine.TransactionInput {
	if tx.Inputs == nil {
		return nil
	}

	inputs := make([]*engine.TransactionInput, 0)
	for _, input := range tx.Inputs {
		inputs = append(inputs, MapToTransactionInputSPV(input))
	}
	return inputs
}

func mapToSPVIncludeUtxos(tx *spvwalletmodels.TransactionConfig) []*engine.UtxoPointer {
	if tx.IncludeUtxos == nil {
		return nil
	}

	includeUtxos := make([]*engine.UtxoPointer, 0)
	for _, utxo := range tx.IncludeUtxos {
		includeUtxos = append(includeUtxos, MapToUtxoPointerSPV(utxo))
	}
	return includeUtxos
}

func mapToSPVFromUtxos(tx *spvwalletmodels.TransactionConfig) []*engine.UtxoPointer {
	if tx.FromUtxos == nil {
		return nil
	}

	fromUtxos := make([]*engine.UtxoPointer, 0)
	for _, utxo := range tx.FromUtxos {
		fromUtxos = append(fromUtxos, MapToUtxoPointerSPV(utxo))
	}
	return fromUtxos
}

func mapToSPVDestinations(tx *spvwalletmodels.TransactionConfig) []*engine.Destination {
	if tx.ChangeDestinations == nil {
		return nil
	}

	destinations := make([]*engine.Destination, 0)
	for _, destination := range tx.ChangeDestinations {
		destinations = append(destinations, MapToDestinationSPV(destination))
	}
	return destinations
}

// MapToTransactionConfigContract will map the transaction-config model from spv-wallet-models to the spv-wallet contract
func MapToTransactionConfigContract(tx *engine.TransactionConfig) *spvwalletmodels.TransactionConfig {
	if tx == nil {
		return nil
	}

	return &spvwalletmodels.TransactionConfig{
		ChangeDestinations:         mapToContractDestinations(tx),
		ChangeStrategy:             string(tx.ChangeDestinationsStrategy),
		ChangeMinimumSatoshis:      tx.ChangeMinimumSatoshis,
		ChangeNumberOfDestinations: tx.ChangeNumberOfDestinations,
		ChangeSatoshis:             tx.ChangeSatoshis,
		ExpiresIn:                  tx.ExpiresIn,
		FeeUnit:                    MapToFeeUnitContract(tx.FeeUnit),
		FromUtxos:                  mapToContractFromUtxos(tx),
		IncludeUtxos:               mapToContractIncludeUtxos(tx),
		Inputs:                     mapToContractInputs(tx),
		Outputs:                    mapToContractOutputs(tx),
		SendAllTo:                  MapToTransactionOutputContract(tx.SendAllTo),
		Sync:                       MapToSyncConfigContract(tx.Sync),
	}
}

func mapToContractOutputs(tx *engine.TransactionConfig) []*spvwalletmodels.TransactionOutput {
	if tx.Outputs == nil {
		return nil
	}

	outputs := make([]*spvwalletmodels.TransactionOutput, 0)
	for _, output := range tx.Outputs {
		outputs = append(outputs, MapToTransactionOutputContract(output))
	}
	return outputs
}

func mapToContractInputs(tx *engine.TransactionConfig) []*spvwalletmodels.TransactionInput {
	if tx.Inputs == nil {
		return nil
	}

	inputs := make([]*spvwalletmodels.TransactionInput, 0)
	for _, input := range tx.Inputs {
		inputs = append(inputs, MapToTransactionInputContract(input))
	}
	return inputs
}

func mapToContractIncludeUtxos(tx *engine.TransactionConfig) []*spvwalletmodels.UtxoPointer {
	if tx.IncludeUtxos == nil {
		return nil
	}

	includeUtxos := make([]*spvwalletmodels.UtxoPointer, 0)
	for _, utxo := range tx.IncludeUtxos {
		includeUtxos = append(includeUtxos, MapToUtxoPointer(utxo))
	}
	return includeUtxos
}

func mapToContractFromUtxos(tx *engine.TransactionConfig) []*spvwalletmodels.UtxoPointer {
	if tx.FromUtxos == nil {
		return nil
	}

	fromUtxos := make([]*spvwalletmodels.UtxoPointer, 0)
	for _, utxo := range tx.FromUtxos {
		fromUtxos = append(fromUtxos, MapToUtxoPointer(utxo))
	}
	return fromUtxos
}

func mapToContractDestinations(tx *engine.TransactionConfig) []*spvwalletmodels.Destination {
	if tx.ChangeDestinations == nil {
		return nil
	}

	destinations := make([]*spvwalletmodels.Destination, 0)
	for _, destination := range tx.ChangeDestinations {
		destinations = append(destinations, MapToDestinationContract(destination))
	}
	return destinations
}

// MapToDraftTransactionContract will map the transaction-output model from spv-wallet to the spv-wallet-models contract
func MapToDraftTransactionContract(tx *engine.DraftTransaction) *spvwalletmodels.DraftTransaction {
	if tx == nil {
		return nil
	}

	return &spvwalletmodels.DraftTransaction{
		Model:         *common.MapToContract(&tx.Model),
		ID:            tx.ID,
		Hex:           tx.Hex,
		XpubID:        tx.XpubID,
		ExpiresAt:     tx.ExpiresAt,
		Configuration: *MapToTransactionConfigContract(&tx.Configuration),
	}
}

// MapToTransactionInputContract will map the transaction-output model from spv-wallet-models to the spv-wallet contract
func MapToTransactionInputContract(inp *engine.TransactionInput) *spvwalletmodels.TransactionInput {
	if inp == nil {
		return nil
	}

	return &spvwalletmodels.TransactionInput{
		Utxo:        *MapToUtxoContract(&inp.Utxo),
		Destination: *MapToDestinationContract(&inp.Destination),
	}
}

// MapToTransactionInputSPV will map the transaction-output model from spv-wallet to the spv-wallet-models contract
func MapToTransactionInputSPV(inp *spvwalletmodels.TransactionInput) *engine.TransactionInput {
	if inp == nil {
		return nil
	}

	return &engine.TransactionInput{
		Utxo:        *MapToUtxoSPV(&inp.Utxo),
		Destination: *MapToDestinationSPV(&inp.Destination),
	}
}

// MapToTransactionOutputContract will map the transaction-output model from spv-wallet to the spv-wallet-models contract
func MapToTransactionOutputContract(out *engine.TransactionOutput) *spvwalletmodels.TransactionOutput {
	if out == nil {
		return nil
	}

	scriptOutputs := make([]*spvwalletmodels.ScriptOutput, 0)
	for _, scriptOutput := range out.Scripts {
		scriptOutputs = append(scriptOutputs, MapToScriptOutputContract(scriptOutput))
	}

	return &spvwalletmodels.TransactionOutput{
		OpReturn:     MapToOpReturnContract(out.OpReturn),
		PaymailP4:    MapToPaymailP4Contract(out.PaymailP4),
		Satoshis:     out.Satoshis,
		Script:       out.Script,
		Scripts:      scriptOutputs,
		To:           out.To,
		UseForChange: out.UseForChange,
	}
}

// MapToTransactionOutputSPV will map the transaction-output model from spv-wallet-models to the spv-wallet contract
func MapToTransactionOutputSPV(out *spvwalletmodels.TransactionOutput) *engine.TransactionOutput {
	if out == nil {
		return nil
	}

	scriptOutputs := make([]*engine.ScriptOutput, 0)
	for _, scriptOutput := range out.Scripts {
		scriptOutputs = append(scriptOutputs, MapToScriptOutputSPV(scriptOutput))
	}

	return &engine.TransactionOutput{
		OpReturn:     MapToOpReturnSPV(out.OpReturn),
		PaymailP4:    MapToPaymailP4SPV(out.PaymailP4),
		Satoshis:     out.Satoshis,
		Script:       out.Script,
		Scripts:      scriptOutputs,
		To:           out.To,
		UseForChange: out.UseForChange,
	}
}

// MapToMapProtocolContract will map the transaction-output model from spv-wallet to the spv-wallet-models contract
func MapToMapProtocolContract(mp *engine.MapProtocol) *spvwalletmodels.MapProtocol {
	if mp == nil {
		return nil
	}

	return &spvwalletmodels.MapProtocol{
		App:  mp.App,
		Keys: mp.Keys,
		Type: mp.Type,
	}
}

// MapToMapProtocolSPV will map the transaction-output model from spv-wallet-models to the spv-wallet contract
func MapToMapProtocolSPV(mp *spvwalletmodels.MapProtocol) *engine.MapProtocol {
	if mp == nil {
		return nil
	}

	return &engine.MapProtocol{
		App:  mp.App,
		Keys: mp.Keys,
		Type: mp.Type,
	}
}

// MapToOpReturnContract will map the transaction-output model from spv-wallet to the spv-wallet-models contract
func MapToOpReturnContract(op *engine.OpReturn) *spvwalletmodels.OpReturn {
	if op == nil {
		return nil
	}

	return &spvwalletmodels.OpReturn{
		Hex:         op.Hex,
		HexParts:    op.HexParts,
		Map:         MapToMapProtocolContract(op.Map),
		StringParts: op.StringParts,
	}
}

// MapToOpReturnSPV will map the op-return model from spv-wallet-models to the spv-wallet contract
func MapToOpReturnSPV(op *spvwalletmodels.OpReturn) *engine.OpReturn {
	if op == nil {
		return nil
	}

	return &engine.OpReturn{
		Hex:         op.Hex,
		HexParts:    op.HexParts,
		Map:         MapToMapProtocolSPV(op.Map),
		StringParts: op.StringParts,
	}
}
