package protos

type CraftShiftingCompleteProcessAllResponse struct {
	ResponsePacket
	CraftInfoDBs   []ShiftingCraftInfoDB
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
}
