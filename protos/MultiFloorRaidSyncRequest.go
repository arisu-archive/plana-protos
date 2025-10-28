package protos

type MultiFloorRaidSyncRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SeasonId *int64 `json:",omitempty,omitzero"`
}
