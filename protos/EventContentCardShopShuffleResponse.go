package protos

type EventContentCardShopShuffleResponse struct {
	ResponsePacket
	CardShopElementDBs []CardShopElementDB `json:",omitempty,omitzero"`
}
