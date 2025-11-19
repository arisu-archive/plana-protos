package protos

type ClanAllAssistListResponse struct {
	ResponsePacket
	AssistCharacterDBs            []AssistCharacterDB       `json:",omitempty,omitzero"`
	AssistCharacterRentHistoryDBs []ClanAssistRentHistoryDB `json:",omitempty,omitzero"`
	ClanDBId                      int64                     `json:",omitempty,omitzero"`
}
