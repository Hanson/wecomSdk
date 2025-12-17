package wecom

import (
    wecomsub "github.com/Hanson/wecomSdk/wecom"
)

type Config = wecomsub.Config
type Client = wecomsub.Client
type APIError = wecomsub.APIError
type Error = wecomsub.Error
type TokenProvider = wecomsub.TokenProvider

// messages
type TextMessage = wecomsub.TextMessage
type MarkdownMessage = wecomsub.MarkdownMessage
type ImageMessage = wecomsub.ImageMessage
type FileMessage = wecomsub.FileMessage
type NewsMessage = wecomsub.NewsMessage
type MessageSendResponse = wecomsub.MessageSendResponse

// media
type MediaUploadResponse = wecomsub.MediaUploadResponse

// generic request helpers
var (
    // methods on Client are available via type alias; expose constructor here
    _ = wecomsub.NewClient
)

func NewClient(cfg Config) (*Client, error) { return wecomsub.NewClient(cfg) }

// api typed req/resp
type (
    userUpdateReq = wecomsub.UserUpdateReq
    userUpdateResp = wecomsub.UserUpdateResp
    userCreateReq = wecomsub.UserCreateReq
    userCreateResp = wecomsub.UserCreateResp
    userGetResp = wecomsub.UserGetResp
    userDeleteResp = wecomsub.UserDeleteResp
    userSimpleListReq = wecomsub.UserSimpleListReq
    userSimpleListResp = wecomsub.UserSimpleListResp
    userListReq = wecomsub.UserListReq
    userListResp = wecomsub.UserListResp

    departmentCreateReq = wecomsub.DepartmentCreateReq
    departmentCreateResp = wecomsub.DepartmentCreateResp
    departmentUpdateReq = wecomsub.DepartmentUpdateReq
    departmentUpdateResp = wecomsub.DepartmentUpdateResp
    departmentDeleteResp = wecomsub.DepartmentDeleteResp
    departmentListResp = wecomsub.DepartmentListResp

    tagCreateReq = wecomsub.TagCreateReq
    tagCreateResp = wecomsub.TagCreateResp
    tagUpdateReq = wecomsub.TagUpdateReq
    tagUpdateResp = wecomsub.TagUpdateResp
    tagDeleteResp = wecomsub.TagDeleteResp
    tagListResp = wecomsub.TagListResp
    tagGetReq = wecomsub.TagGetReq
    tagGetResp = wecomsub.TagGetResp
    tagAddTagUsersReq = wecomsub.TagAddTagUsersReq
    tagAddTagUsersResp = wecomsub.TagAddTagUsersResp
    tagDelTagUsersReq = wecomsub.TagDelTagUsersReq
    tagDelTagUsersResp = wecomsub.TagDelTagUsersResp
    tagAddTagUsersWithDeptReq = wecomsub.TagAddTagUsersWithDeptReq
    tagAddTagUsersWithDeptResp = wecomsub.TagAddTagUsersWithDeptResp

    externalContactListResp = wecomsub.ExternalContactListResp
    externalContactGetReq = wecomsub.ExternalContactGetReq
    externalContactGetResp = wecomsub.ExternalContactGetResp
    externalAddContactWayReq = wecomsub.ExternalAddContactWayReq
    externalAddContactWayResp = wecomsub.ExternalAddContactWayResp
    externalGetContactWayReq = wecomsub.ExternalGetContactWayReq
    externalGetContactWayResp = wecomsub.ExternalGetContactWayResp
    externalListContactWayReq = wecomsub.ExternalListContactWayReq
    externalListContactWayResp = wecomsub.ExternalListContactWayResp
    externalUpdateContactWayReq = wecomsub.ExternalUpdateContactWayReq
    externalUpdateContactWayResp = wecomsub.ExternalUpdateContactWayResp
    externalDelContactWayReq = wecomsub.ExternalDelContactWayReq
    externalDelContactWayResp = wecomsub.ExternalDelContactWayResp
    externalGroupChatListReq = wecomsub.ExternalGroupChatListReq
    externalGroupChatListResp = wecomsub.ExternalGroupChatListResp
    externalGroupChatGetReq = wecomsub.ExternalGroupChatGetReq
    externalGroupChatGetResp = wecomsub.ExternalGroupChatGetResp
    externalRemarkReq = wecomsub.ExternalRemarkReq
    externalRemarkResp = wecomsub.ExternalRemarkResp
    externalFollowUserListResp = wecomsub.ExternalFollowUserListResp
    externalAddCorpTagReq = wecomsub.ExternalAddCorpTagReq
    externalAddCorpTagResp = wecomsub.ExternalAddCorpTagResp
    externalEditCorpTagReq = wecomsub.ExternalEditCorpTagReq
    externalEditCorpTagResp = wecomsub.ExternalEditCorpTagResp
    externalDelCorpTagReq = wecomsub.ExternalDelCorpTagReq
    externalDelCorpTagResp = wecomsub.ExternalDelCorpTagResp
    externalGetCorpTagReq = wecomsub.ExternalGetCorpTagReq
    externalGetCorpTagResp = wecomsub.ExternalGetCorpTagResp

    agentGetReq = wecomsub.AgentGetReq
    agentGetResp = wecomsub.AgentGetResp

    userBatchDeleteReq = wecomsub.UserBatchDeleteReq
    userBatchDeleteResp = wecomsub.UserBatchDeleteResp
    userConvertToOpenIDReq = wecomsub.UserConvertToOpenIDReq
    userConvertToOpenIDResp = wecomsub.UserConvertToOpenIDResp
    userConvertToUserIDReq = wecomsub.UserConvertToUserIDReq
    userConvertToUserIDResp = wecomsub.UserConvertToUserIDResp

    kfAccountListResp = wecomsub.KfAccountListResp
    kfSendMsgReq = wecomsub.KfSendMsgReq
    kfSendMsgResp = wecomsub.KfSendMsgResp
    kfSyncMsgReq = wecomsub.KfSyncMsgReq
    kfSyncMsgResp = wecomsub.KfSyncMsgResp
    kfServiceStateGetReq = wecomsub.KfServiceStateGetReq
    kfServiceStateGetResp = wecomsub.KfServiceStateGetResp
    kfServiceStateTransReq = wecomsub.KfServiceStateTransReq
    kfServiceStateTransResp = wecomsub.KfServiceStateTransResp
    kfAddContactWayReq = wecomsub.KfAddContactWayReq
    kfAddContactWayResp = wecomsub.KfAddContactWayResp
    kfGetContactWayReq = wecomsub.KfGetContactWayReq
    kfGetContactWayResp = wecomsub.KfGetContactWayResp
    kfDelContactWayReq = wecomsub.KfDelContactWayReq
    kfDelContactWayResp = wecomsub.KfDelContactWayResp

    oaApprovalGetDetailReq = wecomsub.OaApprovalGetDetailReq
    oaApprovalGetDetailResp = wecomsub.OaApprovalGetDetailResp
    oaApprovalGetTemplateDetailReq = wecomsub.OaApprovalGetTemplateDetailReq
    oaApprovalGetTemplateDetailResp = wecomsub.OaApprovalGetTemplateDetailResp

    checkinGetDataReq = wecomsub.CheckinGetDataReq
    checkinGetDataResp = wecomsub.CheckinGetDataResp
    checkinGetDayDataReq = wecomsub.CheckinGetDayDataReq
    checkinGetDayDataResp = wecomsub.CheckinGetDayDataResp

    oaCalendarAddReq = wecomsub.OaCalendarAddReq
    oaCalendarAddResp = wecomsub.OaCalendarAddResp
    oaCalendarUpdateReq = wecomsub.OaCalendarUpdateReq
    oaCalendarUpdateResp = wecomsub.OaCalendarUpdateResp
    oaCalendarGetReq = wecomsub.OaCalendarGetReq
    oaCalendarGetResp = wecomsub.OaCalendarGetResp
    oaCalendarDelReq = wecomsub.OaCalendarDelReq
    oaCalendarDelResp = wecomsub.OaCalendarDelResp

    oaScheduleAddReq = wecomsub.OaScheduleAddReq
    oaScheduleAddResp = wecomsub.OaScheduleAddResp
    oaScheduleUpdateReq = wecomsub.OaScheduleUpdateReq
    oaScheduleUpdateResp = wecomsub.OaScheduleUpdateResp
    oaScheduleGetReq = wecomsub.OaScheduleGetReq
    oaScheduleGetResp = wecomsub.OaScheduleGetResp
    oaScheduleDelReq = wecomsub.OaScheduleDelReq
    oaScheduleDelResp = wecomsub.OaScheduleDelResp

    menuButton = wecomsub.MenuButton
    menuCreateReq = wecomsub.MenuCreateReq
    menuCreateResp = wecomsub.MenuCreateResp
    menuDeleteResp = wecomsub.MenuDeleteResp
    menuGetReq = wecomsub.MenuGetReq
    menuGetResp = wecomsub.MenuGetResp
)
