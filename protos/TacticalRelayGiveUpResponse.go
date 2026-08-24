package protos

type TacticalRelayGiveUpResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
