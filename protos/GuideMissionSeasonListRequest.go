package protos

type GuideMissionSeasonListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
