package protos

type MiniGameTableBoardEncounterInputRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	ObjectServerId int64 `json:",omitempty,omitzero"`
	EncounterStage int32 `json:",omitempty,omitzero"`
	SelectedIndex  int32 `json:",omitempty,omitzero"`
}
