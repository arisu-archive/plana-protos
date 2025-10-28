package protos

type AccountGetTutorialResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TutorialIds []int64 `json:",omitempty,omitzero"`
}
