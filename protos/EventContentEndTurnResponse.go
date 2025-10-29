package protos

type EventContentEndTurnResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
}
