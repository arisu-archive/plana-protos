package protos

type TTSGetFileResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsFileReady bool `json:",omitempty,omitzero"`
	TTSFileS3Uri string `json:",omitempty,omitzero"`
}
