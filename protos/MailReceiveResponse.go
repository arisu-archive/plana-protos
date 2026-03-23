package protos

type MailReceiveResponse struct {
	ResponsePacket
	MailServerIds     []int64
	ParcelResultDB    *ParcelResultDB `json:",omitempty,omitzero"`
	BattlePassInfoDBs []*BattlePassInfoDB
}
