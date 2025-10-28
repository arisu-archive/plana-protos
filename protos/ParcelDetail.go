package protos

type ParcelDetail struct {
	OriginParcel ParcelInfo `json:",omitempty,omitzero"`
	MailSendParcel ParcelInfo `json:",omitempty,omitzero"`
	ConvertedParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
	ParcelChangeType ParcelChangeType `json:",omitempty,omitzero"`
}
