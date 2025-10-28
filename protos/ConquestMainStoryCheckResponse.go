package protos

type ConquestMainStoryCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConquestMainStorySummary ConquestMainStorySummary `json:",omitempty,omitzero"`
}
