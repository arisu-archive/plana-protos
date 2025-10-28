package protos

type AccountAuth2Response struct {
	AccountAuthResponse
	Protocol Protocol `json:",omitempty,omitzero"`
}
