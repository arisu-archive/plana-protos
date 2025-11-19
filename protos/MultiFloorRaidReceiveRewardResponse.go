package protos

type MultiFloorRaidReceiveRewardResponse struct {
	ResponsePacket
	MultiFloorRaidDB MultiFloorRaidDB `json:",omitempty,omitzero"`
	ParcelResultDB   ParcelResultDB   `json:",omitempty,omitzero"`
}
