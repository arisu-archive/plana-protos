package protos

type EchelonSaveResponse struct {
	ResponsePacket
	EchelonDB EchelonDB `json:",omitempty,omitzero"`
}
