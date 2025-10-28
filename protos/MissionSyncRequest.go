package protos

type MissionSyncRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
