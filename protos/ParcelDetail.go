package protos

type ParcelDetail struct {
	OriginParcel         ParcelInfo
	MailSendParcel       ParcelInfo
	ConvertedParcelInfos []ParcelInfo
	ParcelChangeType     ParcelChangeType `json:",omitempty,omitzero"`
}
