package protos

type StickerBookDB struct {
	AccountId        int64 `json:",omitempty,omitzero"`
	UnusedStickerDBs []*StickerDB
	UsedStickerDBs   []*StickerDB
}
