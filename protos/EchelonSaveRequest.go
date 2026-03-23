package protos

type EchelonSaveRequest struct {
	RequestPacket
	EchelonDB      *EchelonDB `json:",omitempty,omitzero"`
	AssistUseInfos []*ClanAssistUseInfo
	IsPractice     bool `json:",omitempty,omitzero"`
}
