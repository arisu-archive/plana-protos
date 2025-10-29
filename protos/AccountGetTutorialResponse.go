package protos

type AccountGetTutorialResponse struct {
	ResponsePacket
	TutorialIds []int64 `json:",omitempty,omitzero"`
}
