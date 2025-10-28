package protos

type AccountAuth2Request struct {
	AccountAuthRequest
	Protocol Protocol `json:",omitempty,omitzero"`
}
