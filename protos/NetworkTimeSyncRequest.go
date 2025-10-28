package protos

type NetworkTimeSyncRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
