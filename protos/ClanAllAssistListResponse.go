package protos

type ClanAllAssistListResponse struct {
	ResponsePacket
	AssistCharacterDBs            []*AssistCharacterDB
	AssistCharacterRentHistoryDBs []*ClanAssistRentHistoryDB
	ClanDBId                      int64 `json:",omitempty,omitzero"`
}
