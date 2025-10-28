package protos

type MultiFloorRaidSyncResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MultiFloorRaidDBs []MultiFloorRaidDB `json:",omitempty,omitzero"`
}
