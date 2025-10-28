package protos

type MissionSyncResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
