package wecom

type userUpdateReq struct {
	UserID       string `json:"userid"`
	Name         string `json:"name,omitempty"`
	Department   []int  `json:"department,omitempty"`
	Order        []int  `json:"order,omitempty"`
	Position     string `json:"position,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Email        string `json:"email,omitempty"`
	BizMail      string `json:"biz_mail,omitempty"`
	BizMailAlias struct {
		Item []string `json:"item"`
	} `json:"biz_mail_alias,omitempty"`
	IsLeaderInDept   []int    `json:"is_leader_in_dept,omitempty"`
	DirectLeader     []string `json:"direct_leader,omitempty"`
	Enable           int      `json:"enable,omitempty"`
	AvatarMediaID    string   `json:"avatar_mediaid,omitempty"`
	Telephone        string   `json:"telephone,omitempty"`
	Alias            string   `json:"alias,omitempty"`
	Address          string   `json:"address,omitempty"`
	MainDepartment   int      `json:"main_department,omitempty"`
	ExtAttr          any      `json:"extattr,omitempty"`
	ExternalPosition string   `json:"external_position,omitempty"`
	ExternalProfile  any      `json:"external_profile,omitempty"`
}

type userUpdateResp struct {
	APIError
}

type userCreateReq struct {
	UserID         string `json:"userid"`
	Name           string `json:"name"`
	Department     []int  `json:"department"`
	Position       string `json:"position,omitempty"`
	Mobile         string `json:"mobile,omitempty"`
	Gender         string `json:"gender,omitempty"`
	Email          string `json:"email,omitempty"`
	Alias          string `json:"alias,omitempty"`
	Enable         int    `json:"enable,omitempty"`
	MainDepartment int    `json:"main_department,omitempty"`
}

type userCreateResp struct {
	APIError
}

type userGetResp struct {
	APIError
	User
}

type userSimpleListReq struct {
	DepartmentID int `json:"department_id"`
	FetchChild   int `json:"fetch_child"`
}

type userSimpleListResp struct {
	APIError
	UserList []struct {
		UserID string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
}

type userListReq struct {
	DepartmentID int `json:"department_id"`
	FetchChild   int `json:"fetch_child"`
}

type userListResp struct {
	APIError
	UserList []User `json:"userlist"`
}

type userDeleteResp struct {
	APIError
}

type messageSendResp = MessageSendResponse

type mediaUploadReq struct {
	Type     string
	Filename string
	Data     []byte
}

type mediaUploadResp = MediaUploadResponse

type agentGetReq struct {
	AgentID int `json:"agentid"`
}

type agentGetResp struct {
	APIError
	AgentID       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	Description   string `json:"description,omitempty"`
}

type externalContactListResp struct {
	APIError
	ExternalUserID []string `json:"external_userid"`
}

type externalContactGetReq struct {
	ExternalUserID string `json:"external_userid"`
}

type externalContactGetResp struct {
	APIError
	ExternalContact ExternalContact `json:"external_contact"`
}

type departmentCreateReq struct {
	Name     string `json:"name"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

type departmentCreateResp struct {
	APIError
	ID int `json:"id"`
}

type departmentUpdateReq struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

type departmentUpdateResp struct {
	APIError
}

type departmentDeleteResp struct {
	APIError
}

type departmentListResp struct {
	APIError
	Department []Department `json:"department"`
}

type tagCreateReq struct {
	TagName string `json:"tagname"`
	TagID   int    `json:"tagid,omitempty"`
}

type tagCreateResp struct {
	APIError
	TagID int `json:"tagid"`
}

type tagUpdateReq struct {
	TagID   int    `json:"tagid"`
	TagName string `json:"tagname"`
}

type tagUpdateResp struct {
	APIError
}

type tagDeleteResp struct {
	APIError
}

type tagListResp struct {
	APIError
	TagList []Tag `json:"taglist"`
}

type tagGetReq struct {
	TagID int `json:"tagid"`
}

type tagGetResp struct {
	APIError
	UserList []struct {
		UserID string `json:"userid"`
		Name   string `json:"name,omitempty"`
	} `json:"userlist"`
}

type tagAddTagUsersReq struct {
	TagID   int      `json:"tagid"`
	UserIDs []string `json:"userlist"`
}

type tagAddTagUsersResp struct {
	APIError
}

type tagDelTagUsersReq struct {
	TagID   int      `json:"tagid"`
	UserIDs []string `json:"userlist"`
}

type tagAddTagUsersWithDeptReq struct {
	TagID    int      `json:"tagid"`
	UserIDs  []string `json:"userlist,omitempty"`
	PartyIDs []int    `json:"partylist,omitempty"`
}

type tagAddTagUsersWithDeptResp struct{ APIError }

type externalAddContactWayReq struct {
	Type       int      `json:"type"`
	Scene      int      `json:"scene"`
	Style      int      `json:"style,omitempty"`
	Remark     string   `json:"remark,omitempty"`
	SkipVerify bool     `json:"skip_verify,omitempty"`
	State      string   `json:"state,omitempty"`
	User       []string `json:"user,omitempty"`
	Party      []int    `json:"party,omitempty"`
	IsTmp      bool     `json:"is_tmp,omitempty"`
	ExpiresIn  int      `json:"expires_in,omitempty"`
}

type externalAddContactWayResp struct {
	APIError
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code,omitempty"`
}

type externalGetContactWayReq struct {
	ConfigID string `json:"config_id"`
}
type externalGetContactWayResp struct {
	APIError
	ContactWay any `json:"contact_way"`
}

type externalListContactWayReq struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor,omitempty"`
}
type externalListContactWayResp struct {
	APIError
	ContactWay []any  `json:"contact_way"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type externalUpdateContactWayReq struct {
	ConfigID string `json:"config_id"`
	Remark   string `json:"remark,omitempty"`
}
type externalUpdateContactWayResp struct{ APIError }

type externalDelContactWayReq struct {
	ConfigID string `json:"config_id"`
}
type externalDelContactWayResp struct{ APIError }

type externalGroupChatListReq struct {
	StatusFilter int `json:"status_filter,omitempty"`
	OwnerFilter  struct {
		UserIDList []string `json:"userid_list"`
	} `json:"owner_filter,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type externalGroupChatListResp struct {
	APIError
	GroupChatList []struct {
		ChatID string `json:"chat_id"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type externalGroupChatGetReq struct {
	ChatID string `json:"chat_id"`
}
type externalGroupChatGetResp struct {
	APIError
	GroupChat any `json:"group_chat"`
}

type kfAccountListResp struct {
	APIError
	AccountList []struct {
		OpenKfID string `json:"open_kfid"`
		Name     string `json:"name,omitempty"`
	} `json:"account_list"`
}

type kfSendMsgReq struct {
	OpenKfID string `json:"open_kfid"`
	ToUser   string `json:"touser"`
	MsgType  string `json:"msgtype"`
	Text     struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
}

type kfSendMsgResp struct {
	APIError
	MsgID string `json:"msgid,omitempty"`
}

type kfSyncMsgReq struct {
	Cursor   string `json:"cursor,omitempty"`
	Token    string `json:"token,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	OpenKfID string `json:"open_kfid,omitempty"`
}

type kfSyncMsgResp struct {
	APIError
	MsgList    []any  `json:"msg_list"`
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type kfServiceStateGetReq struct {
	OpenKfID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

type kfServiceStateGetResp struct {
	APIError
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid,omitempty"`
}

type kfServiceStateTransReq struct {
	OpenKfID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid,omitempty"`
}

type kfServiceStateTransResp struct{ APIError }

type kfAddContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
	Scene    int    `json:"scene,omitempty"`
	Style    int    `json:"style,omitempty"`
}
type kfAddContactWayResp struct {
	APIError
	URL    string `json:"url,omitempty"`
	QRCode string `json:"qr_code,omitempty"`
}
type kfGetContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
}
type kfGetContactWayResp struct {
	APIError
	URL    string `json:"url,omitempty"`
	QRCode string `json:"qr_code,omitempty"`
}
type kfDelContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
}
type kfDelContactWayResp struct{ APIError }

type externalRemarkReq struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	Remark         string   `json:"remark,omitempty"`
	Description    string   `json:"description,omitempty"`
	RemarkCompany  string   `json:"remark_company,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	RemarkEmails   []string `json:"remark_emails,omitempty"`
}

type externalRemarkResp struct{ APIError }

type externalFollowUserListResp struct {
	APIError
	FollowUser []string `json:"follow_user"`
}

type externalCorpTag struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Tag       []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CreateTime int64  `json:"create_time"`
	} `json:"tag"`
}

type externalAddCorpTagReq struct {
	GroupID   string `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Order     int    `json:"order,omitempty"`
	Tag       []struct {
		Name  string `json:"name"`
		Order int    `json:"order,omitempty"`
	} `json:"tag,omitempty"`
}
type externalAddCorpTagResp struct {
	APIError
	TagGroup externalCorpTag `json:"tag_group"`
}

type externalEditCorpTagReq struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order int    `json:"order,omitempty"`
}
type externalEditCorpTagResp struct{ APIError }

type externalDelCorpTagReq struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}
type externalDelCorpTagResp struct{ APIError }

type externalGetCorpTagReq struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}
type externalGetCorpTagResp struct {
	APIError
	TagGroup []externalCorpTag `json:"tag_group"`
}

type userBatchDeleteReq struct {
	UserIDList []string `json:"useridlist"`
}
type userBatchDeleteResp struct{ APIError }

type userConvertToOpenIDReq struct {
	UserID string `json:"userid"`
}
type userConvertToOpenIDResp struct {
	APIError
	OpenID string `json:"openid"`
}

type userConvertToUserIDReq struct {
	OpenID string `json:"openid"`
}
type userConvertToUserIDResp struct {
	APIError
	UserID string `json:"userid"`
}

// OA 审批
type oaApprovalGetDetailReq struct {
	SpNo string `json:"sp_no"`
}
type oaApprovalGetDetailResp struct {
	APIError
	ApprovalInfo any `json:"approval_info"`
}
type oaApprovalGetTemplateDetailReq struct {
	TemplateID string `json:"template_id"`
}
type oaApprovalGetTemplateDetailResp struct {
	APIError
	Template any `json:"template"`
}

// 考勤打卡
type checkinGetDataReq struct {
	StartTime  int      `json:"starttime"`
	EndTime    int      `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}
type checkinGetDataResp struct {
	APIError
	CheckinData []any `json:"checkindata"`
}
type checkinGetDayDataReq struct {
	Date       int      `json:"date"`
	UserIDList []string `json:"useridlist"`
}
type checkinGetDayDataResp struct {
	APIError
	DayData []any `json:"daydata"`
}

// 日历
type oaCalendarAddReq struct {
	Organizer string `json:"organizer"`
	Summary   string `json:"summary"`
}
type oaCalendarAddResp struct {
	APIError
	CalID string `json:"cal_id"`
}
type oaCalendarUpdateReq struct {
	CalID   string `json:"cal_id"`
	Summary string `json:"summary,omitempty"`
}
type oaCalendarUpdateResp struct{ APIError }
type oaCalendarGetReq struct {
	CalID string `json:"cal_id"`
}
type oaCalendarGetResp struct {
	APIError
	Calendar any `json:"calendar"`
}
type oaCalendarDelReq struct {
	CalID string `json:"cal_id"`
}
type oaCalendarDelResp struct{ APIError }

// 日程
type oaScheduleAddReq struct {
	CalID     string `json:"cal_id"`
	Organizer string `json:"organizer"`
	Summary   string `json:"summary"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}
type oaScheduleAddResp struct {
	APIError
	ScheduleID string `json:"schedule_id"`
}
type oaScheduleUpdateReq struct {
	ScheduleID string `json:"schedule_id"`
	Summary    string `json:"summary,omitempty"`
}
type oaScheduleUpdateResp struct{ APIError }
type oaScheduleGetReq struct {
	ScheduleID string `json:"schedule_id"`
}
type oaScheduleGetResp struct {
	APIError
	Schedule any `json:"schedule"`
}
type oaScheduleDelReq struct {
	ScheduleID string `json:"schedule_id"`
}
type oaScheduleDelResp struct{ APIError }

// 应用菜单
type menuButton struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name"`
	Key       string       `json:"key,omitempty"`
	URL       string       `json:"url,omitempty"`
	SubButton []menuButton `json:"sub_button,omitempty"`
}
type menuCreateReq struct {
	AgentID int          `json:"agentid"`
	Button  []menuButton `json:"button"`
}
type menuCreateResp struct{ APIError }
type menuDeleteResp struct{ APIError }
type menuGetReq struct {
	AgentID int `json:"agentid"`
}
type menuGetResp struct {
	APIError
	Menu any `json:"menu"`
}

type tagDelTagUsersResp struct {
	APIError
}
