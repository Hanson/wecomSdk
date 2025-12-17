package wecom

type APIError struct { ErrCode int `json:"errcode"`; ErrMsg string `json:"errmsg"` }
type MessageSendResponse struct { APIError; InvalidUser string `json:"invaliduser,omitempty"`; InvalidParty string `json:"invalidparty,omitempty"`; InvalidTag string `json:"invalidtag,omitempty"`; MsgID string `json:"msgid,omitempty"` }
