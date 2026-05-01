package protos

type AccountCreateRequest struct {
	RequestPacket
	DevId           string `json:",omitempty,omitzero"`
	Version         int64  `json:",omitempty,omitzero"`
	IMEI            int64  `json:",omitempty,omitzero"`
	AccessIP        string `json:",omitempty,omitzero"`
	MarketId        string `json:",omitempty,omitzero"`
	UserType        string `json:",omitempty,omitzero"`
	AdvertisementId string `json:",omitempty,omitzero"`
	OSType          string `json:",omitempty,omitzero"`
	OSVersion       string `json:",omitempty,omitzero"`
}
