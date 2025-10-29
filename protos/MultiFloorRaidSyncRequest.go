package protos

type MultiFloorRaidSyncRequest struct {
	RequestPacket
	SeasonId *int64 `json:",omitempty,omitzero"`
}
