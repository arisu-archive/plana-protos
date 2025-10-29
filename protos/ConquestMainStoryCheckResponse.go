package protos

type ConquestMainStoryCheckResponse struct {
	ResponsePacket
	ConquestMainStorySummary ConquestMainStorySummary `json:",omitempty,omitzero"`
}
