package protos

type AccountCheckYostarResponse struct {
	ResponsePacket
	ResultState  int32  `json:",omitempty,omitzero"`
	ResultMessag string `json:",omitempty,omitzero"`
	Birth        string `json:",omitempty,omitzero"`
	EncryptedKey string `json:",omitempty,omitzero"`
	SignedKey    string `json:",omitempty,omitzero"`
	EncryptedIV  string `json:",omitempty,omitzero"`
	SignedIV     string `json:",omitempty,omitzero"`
}
