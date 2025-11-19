package protos

type EchelonSaveRequest struct {
	RequestPacket
	EchelonDB      EchelonDB           `json:",omitempty,omitzero"`
	AssistUseInfos []ClanAssistUseInfo `json:",omitempty,omitzero"`
	IsPractice     bool                `json:",omitempty,omitzero"`
}
