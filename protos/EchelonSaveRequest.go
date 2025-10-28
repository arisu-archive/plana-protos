package protos

type EchelonSaveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonDB EchelonDB `json:",omitempty,omitzero"`
	AssistUseInfos []ClanAssistUseInfo `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
}
