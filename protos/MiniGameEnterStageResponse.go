package protos

type MiniGameEnterStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
