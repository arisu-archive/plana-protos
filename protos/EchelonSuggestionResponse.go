package protos

type EchelonSuggestionResponse struct {
	ResponsePacket
	EchelonDB *EchelonDB `json:",omitempty,omitzero"`
	GroundId  int64      `json:",omitempty,omitzero"`
}
