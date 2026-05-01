package protos

type SessionKey struct {
	AccountServerId int64  `json:",omitempty,omitzero"`
	MxToken         string `json:",omitempty,omitzero"`
}
