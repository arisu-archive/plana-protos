package protos

type AccountSetTutorialRequest struct {
	RequestPacket
	TutorialIds []int64 `json:",omitempty,omitzero"`
}
