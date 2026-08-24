package protos

type AccountCheckYostarRequest struct {
	RequestPacket
	UID                int64 `json:",omitempty,omitzero"`
	EnterTicket        string
	PassCookieResult   bool `json:",omitempty,omitzero"`
	Cookie             string
	ClientGeneratedKey string
	ClientGeneratedIV  string
}
