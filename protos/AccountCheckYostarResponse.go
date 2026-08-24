package protos

type AccountCheckYostarResponse struct {
	ResponsePacket
	ResultState  int32 `json:",omitempty,omitzero"`
	ResultMessag string
	Birth        string
	EncryptedKey string
	SignedKey    string
	EncryptedIV  string
	SignedIV     string
}
