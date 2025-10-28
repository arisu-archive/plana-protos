package protos

type AccountCheckYostarRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UID int64 `json:",omitempty,omitzero"`
	EnterTicket string `json:",omitempty,omitzero"`
	PassCookieResult bool `json:",omitempty,omitzero"`
	Cookie string `json:",omitempty,omitzero"`
	ClientGeneratedKey string `json:",omitempty,omitzero"`
	ClientGeneratedIV string `json:",omitempty,omitzero"`
}
