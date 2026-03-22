package protos

type EventContentEnterMainStageResponse struct {
	ResponsePacket
	SaveDataDB   EventContentMainStageSaveDB
	IsOnSubEvent bool `json:",omitempty,omitzero"`
}
