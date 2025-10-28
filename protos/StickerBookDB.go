package protos

type StickerBookDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	UnusedStickerDBs []StickerDB `json:",omitempty,omitzero"`
	UsedStickerDBs []StickerDB `json:",omitempty,omitzero"`
}
