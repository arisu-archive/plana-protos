package protos

type CafeDeployFurnitureResponse struct {
	ResponsePacket
	CafeDB               CafeDB
	NewFurnitureServerId int64 `json:",omitempty,omitzero"`
	ChangedFurnitureDBs  []FurnitureDB
}
