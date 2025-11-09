package protos

type RequestPacket struct {
	BasePacket
	Resendable                    bool    `json:",omitempty,omitzero"`
	Hash                          int64   `json:",omitempty,omitzero"`
	IsTest                        bool    `json:",omitempty,omitzero"`
	ModifiedServerTime__DebugOnly *MxTime `json:",omitempty,omitzero"`
}
