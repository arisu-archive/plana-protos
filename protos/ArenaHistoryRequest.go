package protos

type ArenaHistoryRequest struct {
	RequestPacket
	SearchStartDate MxTime `json:",omitempty,omitzero"`
	Count           int32  `json:",omitempty,omitzero"`
}
