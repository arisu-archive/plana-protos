package protos

type EchelonSaveRequest struct {
	RequestPacket
	EchelonDB      EchelonDB
	AssistUseInfos []ClanAssistUseInfo
	IsPractice     bool `json:",omitempty,omitzero"`
}
