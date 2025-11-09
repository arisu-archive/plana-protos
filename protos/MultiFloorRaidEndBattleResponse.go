package protos

type MultiFloorRaidEndBattleResponse struct {
	ResponsePacket
	MultiFloorRaidDB MultiFloorRaidDB `json:",omitempty,omitzero"`
	ParcelResultDB   ParcelResultDB   `json:",omitempty,omitzero"`
}
