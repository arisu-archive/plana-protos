package protos

type CafePresetDetailResponse struct {
	ResponsePacket
	DeployCountByFurnitureId map[int64]int64 `json:",omitempty,omitzero"`
}
