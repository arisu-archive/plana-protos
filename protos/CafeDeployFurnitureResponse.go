package protos

type CafeDeployFurnitureResponse struct {
	ResponsePacket
	CafeDB               CafeDB        `json:",omitempty,omitzero"`
	NewFurnitureServerId int64         `json:",omitempty,omitzero"`
	ChangedFurnitureDBs  []FurnitureDB `json:",omitempty,omitzero"`
}
