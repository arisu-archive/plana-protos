package protos

type MultiFloorRaidSyncResponse struct {
	ResponsePacket
	MultiFloorRaidDBs []MultiFloorRaidDB `json:",omitempty,omitzero"`
}
