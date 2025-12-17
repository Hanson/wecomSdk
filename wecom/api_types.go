package wecom

type UserUpdateReq struct {
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

type UserUpdateResp struct {
	APIError
}

type UserCreateReq struct {
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

type UserCreateResp struct {
	APIError
}

type UserGetResp struct {
	APIError
	User
}

type UserSimpleListReq struct {
	DepartmentID int `json:"department_id"`
	FetchChild   int `json:"fetch_child"`
}

type UserSimpleListResp struct {
	APIError
	UserList []struct {
		UserID string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
}

type UserListReq struct {
	DepartmentID int `json:"department_id"`
	FetchChild   int `json:"fetch_child"`
}

type UserListResp struct {
	APIError
	UserList []User `json:"userlist"`
}

type UserDeleteResp struct {
	APIError
}

type messageSendResp = MessageSendResponse

type MediaUploadReq struct {
	Type     string
	Filename string
	Data     []byte
}

type MediaUploadResp = MediaUploadResponse

type AgentGetReq struct {
	AgentID int `json:"agentid"`
}

type AgentGetResp struct {
	APIError
	AgentID       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url,omitempty"`
	Description   string `json:"description,omitempty"`
}

type ExternalContactListResp struct {
	APIError
	ExternalUserID []string `json:"external_userid"`
}

type ExternalContactGetReq struct {
	ExternalUserID string `json:"external_userid"`
}

type ExternalContactGetResp struct {
	APIError
	ExternalContact ExternalContact `json:"external_contact"`
}

type DepartmentCreateReq struct {
	Name     string `json:"name"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

type DepartmentCreateResp struct {
	APIError
	ID int `json:"id"`
}

type DepartmentUpdateReq struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	ParentID int    `json:"parentid,omitempty"`
	Order    int    `json:"order,omitempty"`
}

type DepartmentUpdateResp struct {
	APIError
}

type DepartmentDeleteResp struct {
	APIError
}

type DepartmentListResp struct {
	APIError
	Department []Department `json:"department"`
}

type TagCreateReq struct {
	TagName string `json:"tagname"`
	TagID   int    `json:"tagid,omitempty"`
}

type TagCreateResp struct {
	APIError
	TagID int `json:"tagid"`
}

type TagUpdateReq struct {
	TagID   int    `json:"tagid"`
	TagName string `json:"tagname"`
}

type TagUpdateResp struct {
	APIError
}

type TagDeleteResp struct {
	APIError
}

type TagListResp struct {
	APIError
	TagList []Tag `json:"taglist"`
}

type TagGetReq struct {
	TagID int `json:"tagid"`
}

type TagGetResp struct {
	APIError
	UserList []struct {
		UserID string `json:"userid"`
		Name   string `json:"name,omitempty"`
	} `json:"userlist"`
}

type TagAddTagUsersReq struct {
	TagID   int      `json:"tagid"`
	UserIDs []string `json:"userlist"`
}

type TagAddTagUsersResp struct {
	APIError
}

type TagDelTagUsersReq struct {
	TagID   int      `json:"tagid"`
	UserIDs []string `json:"userlist"`
}
type TagDelTagUsersResp struct{ APIError }

type TagAddTagUsersWithDeptReq struct {
	TagID    int      `json:"tagid"`
	UserIDs  []string `json:"userlist,omitempty"`
	PartyIDs []int    `json:"partylist,omitempty"`
}

type TagAddTagUsersWithDeptResp struct{ APIError }

type ExternalAddContactWayReq struct {
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

type ExternalAddContactWayResp struct {
	APIError
	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code,omitempty"`
}

type ExternalGetContactWayReq struct {
	ConfigID string `json:"config_id"`
}
type ExternalGetContactWayResp struct {
	APIError
	ContactWay any `json:"contact_way"`
}

type ExternalListContactWayReq struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor,omitempty"`
}
type ExternalListContactWayResp struct {
	APIError
	ContactWay []any  `json:"contact_way"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type ExternalUpdateContactWayReq struct {
	ConfigID string `json:"config_id"`
	Remark   string `json:"remark,omitempty"`
}
type ExternalUpdateContactWayResp struct{ APIError }

type ExternalDelContactWayReq struct {
	ConfigID string `json:"config_id"`
}
type ExternalDelContactWayResp struct{ APIError }

type ExternalGroupChatListReq struct {
	StatusFilter int `json:"status_filter,omitempty"`
	OwnerFilter  struct {
		UserIDList []string `json:"userid_list"`
	} `json:"owner_filter,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
type ExternalGroupChatListResp struct {
	APIError
	GroupChatList []struct {
		ChatID string `json:"chat_id"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type ExternalGroupChatGetReq struct {
	ChatID string `json:"chat_id"`
}
type ExternalGroupChatGetResp struct {
	APIError
	GroupChat any `json:"group_chat"`
}

type KfAccountListResp struct {
	APIError
	AccountList []struct {
		OpenKfID string `json:"open_kfid"`
		Name     string `json:"name,omitempty"`
	} `json:"account_list"`
}

type KfSendMsgReq struct {
	OpenKfID string `json:"open_kfid"`
	ToUser   string `json:"touser"`
	MsgType  string `json:"msgtype"`
	Text     struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
}

type KfSendMsgResp struct {
	APIError
	MsgID string `json:"msgid,omitempty"`
}

type KfSyncMsgReq struct {
	Cursor   string `json:"cursor,omitempty"`
	Token    string `json:"token,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	OpenKfID string `json:"open_kfid,omitempty"`
}

type KfSyncMsgResp struct {
	APIError
	MsgList    []any  `json:"msg_list"`
	HasMore    bool   `json:"has_more"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type KfServiceStateGetReq struct {
	OpenKfID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
}

type KfServiceStateGetResp struct {
	APIError
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid,omitempty"`
}

type KfServiceStateTransReq struct {
	OpenKfID       string `json:"open_kfid"`
	ExternalUserID string `json:"external_userid"`
	ServiceState   int    `json:"service_state"`
	ServicerUserID string `json:"servicer_userid,omitempty"`
}

type KfServiceStateTransResp struct{ APIError }

type KfAddContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
	Scene    int    `json:"scene,omitempty"`
	Style    int    `json:"style,omitempty"`
}
type KfAddContactWayResp struct {
	APIError
	URL    string `json:"url,omitempty"`
	QRCode string `json:"qr_code,omitempty"`
}
type KfGetContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
}
type KfGetContactWayResp struct {
	APIError
	URL    string `json:"url,omitempty"`
	QRCode string `json:"qr_code,omitempty"`
}
type KfDelContactWayReq struct {
	OpenKfID string `json:"open_kfid"`
}
type KfDelContactWayResp struct{ APIError }

type ExternalRemarkReq struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	Remark         string   `json:"remark,omitempty"`
	Description    string   `json:"description,omitempty"`
	RemarkCompany  string   `json:"remark_company,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	RemarkEmails   []string `json:"remark_emails,omitempty"`
}

type ExternalRemarkResp struct{ APIError }

type ExternalFollowUserListResp struct {
	APIError
	FollowUser []string `json:"follow_user"`
}

type ExternalCorpTag struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	Tag       []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CreateTime int64  `json:"create_time"`
	} `json:"tag"`
}

type ExternalAddCorpTagReq struct {
	GroupID   string `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	Order     int    `json:"order,omitempty"`
	Tag       []struct {
		Name  string `json:"name"`
		Order int    `json:"order,omitempty"`
	} `json:"tag,omitempty"`
}
type ExternalAddCorpTagResp struct {
	APIError
	TagGroup ExternalCorpTag `json:"tag_group"`
}

type ExternalEditCorpTagReq struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Order int    `json:"order,omitempty"`
}
type ExternalEditCorpTagResp struct{ APIError }

type ExternalDelCorpTagReq struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}
type ExternalDelCorpTagResp struct{ APIError }

type ExternalGetCorpTagReq struct {
	TagID   []string `json:"tag_id,omitempty"`
	GroupID []string `json:"group_id,omitempty"`
}
type ExternalGetCorpTagResp struct {
	APIError
	TagGroup []ExternalCorpTag `json:"tag_group"`
}

type UserBatchDeleteReq struct {
	UserIDList []string `json:"useridlist"`
}
type UserBatchDeleteResp struct{ APIError }

type UserConvertToOpenIDReq struct {
	UserID string `json:"userid"`
}
type UserConvertToOpenIDResp struct {
	APIError
	OpenID string `json:"openid"`
}

type UserConvertToUserIDReq struct {
	OpenID string `json:"openid"`
}
type UserConvertToUserIDResp struct {
	APIError
	UserID string `json:"userid"`
}

// OA 审批
type OaApprovalGetDetailReq struct {
	SpNo string `json:"sp_no"`
}
type OaApprovalGetDetailResp struct {
	APIError
	ApprovalInfo any `json:"approval_info"`
}
type OaApprovalGetTemplateDetailReq struct {
	TemplateID string `json:"template_id"`
}
type OaApprovalGetTemplateDetailResp struct {
	APIError
	Template any `json:"template"`
}

// 考勤打卡
type CheckinGetDataReq struct {
	StartTime  int      `json:"starttime"`
	EndTime    int      `json:"endtime"`
	UserIDList []string `json:"useridlist"`
}
type CheckinGetDataResp struct {
	APIError
	CheckinData []any `json:"checkindata"`
}
type CheckinGetDayDataReq struct {
	Date       int      `json:"date"`
	UserIDList []string `json:"useridlist"`
}
type CheckinGetDayDataResp struct {
	APIError
	DayData []any `json:"daydata"`
}

// 日历
type OaCalendarAddReq struct {
	Organizer string `json:"organizer"`
	Summary   string `json:"summary"`
}
type OaCalendarAddResp struct {
	APIError
	CalID string `json:"cal_id"`
}
type OaCalendarUpdateReq struct {
	CalID   string `json:"cal_id"`
	Summary string `json:"summary,omitempty"`
}
type OaCalendarUpdateResp struct{ APIError }
type OaCalendarGetReq struct {
	CalID string `json:"cal_id"`
}
type OaCalendarGetResp struct {
	APIError
	Calendar any `json:"calendar"`
}
type OaCalendarDelReq struct {
	CalID string `json:"cal_id"`
}
type OaCalendarDelResp struct{ APIError }

// 日程
type OaScheduleAddReq struct {
	CalID     string `json:"cal_id"`
	Organizer string `json:"organizer"`
	Summary   string `json:"summary"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
}
type OaScheduleAddResp struct {
	APIError
	ScheduleID string `json:"schedule_id"`
}
type OaScheduleUpdateReq struct {
	ScheduleID string `json:"schedule_id"`
	Summary    string `json:"summary,omitempty"`
}
type OaScheduleUpdateResp struct{ APIError }
type OaScheduleGetReq struct {
	ScheduleID string `json:"schedule_id"`
}
type OaScheduleGetResp struct {
	APIError
	Schedule any `json:"schedule"`
}
type OaScheduleDelReq struct {
	ScheduleID string `json:"schedule_id"`
}
type OaScheduleDelResp struct{ APIError }

// 应用菜单
type MenuButton struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name"`
	Key       string       `json:"key,omitempty"`
	URL       string       `json:"url,omitempty"`
	SubButton []MenuButton `json:"sub_button,omitempty"`
}
type MenuCreateReq struct {
	AgentID int          `json:"agentid"`
	Button  []MenuButton `json:"button"`
}
type MenuCreateResp struct{ APIError }
type MenuDeleteResp struct{ APIError }
type MenuGetReq struct {
	AgentID int `json:"agentid"`
}
type MenuGetResp struct {
	APIError
	Menu any `json:"menu"`
}
