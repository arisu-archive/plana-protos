package protos

type GuideMissionSeasonListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GuideMissionSeasonDBs []GuideMissionSeasonDB `json:",omitempty,omitzero"`
}
