package wecom

type ExternalContact struct {
	ExternalUserID string `json:"external_userid"`
	Name           string `json:"name,omitempty"`
	Position       string `json:"position,omitempty"`
}

type externalListResponse struct {
	APIError
	ExternalUserID []string `json:"external_userid"`
}

type externalGetResponse struct {
	APIError
	ExternalContact ExternalContact `json:"external_contact"`
}
