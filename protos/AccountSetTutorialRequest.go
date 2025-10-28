package protos

type AccountSetTutorialRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TutorialIds []int64 `json:",omitempty,omitzero"`
}
