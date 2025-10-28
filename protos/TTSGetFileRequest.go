package protos

type TTSGetFileRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
