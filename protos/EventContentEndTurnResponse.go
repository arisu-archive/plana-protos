package protos

type EventContentEndTurnResponse struct {
	ResponsePacket
	SaveDataDB        EventContentMainStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
}
