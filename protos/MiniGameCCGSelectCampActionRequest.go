package protos

type MiniGameCCGSelectCampActionRequest struct {
	RequestPacket
	EventContentId  int64                 `json:",omitempty,omitzero"`
	SelectedOption  MiniGameCCGCampOption `json:",omitempty,omitzero"`
	RemoveCardDBIds []int32               `json:",omitempty,omitzero"`
}
