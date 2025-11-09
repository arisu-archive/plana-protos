package protos

type SystemVersionResponse struct {
	ResponsePacket
	CurrentVersion int64 `json:",omitempty,omitzero"`
	MinimumVersion int64 `json:",omitempty,omitzero"`
	IsDevelopment  bool  `json:",omitempty,omitzero"`
}
