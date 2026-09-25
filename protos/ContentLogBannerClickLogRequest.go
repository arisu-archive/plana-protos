package protos

type ContentLogBannerClickLogRequest struct {
	RequestPacket
	Platform   string `json:",omitempty,omitzero"`
	Device     string `json:",omitempty,omitzero"`
	BannerLogs []*BannerLog
}
