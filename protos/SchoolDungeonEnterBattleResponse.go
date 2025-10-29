package protos

type SchoolDungeonEnterBattleResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
