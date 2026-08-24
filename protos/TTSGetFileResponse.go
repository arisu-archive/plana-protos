package protos

type TTSGetFileResponse struct {
	ResponsePacket
	IsFileReady  bool `json:",omitempty,omitzero"`
	TTSFileS3Uri string
}
