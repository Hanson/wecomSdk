# WeCom SDK for Go（企微自建应用 SDK）

面向企业微信开放平台自建应用的 Go SDK，覆盖令牌管理、消息发送、媒体上传、通讯录（用户/部门/标签）、应用信息查询、客户联系（外部联系人），并附带可运行的单元与真实集成测试。

## 特性
- 令牌缓存与并发安全，过期前自动刷新
- 统一 HTTPS/JSON 请求封装，自动注入 `access_token`
- 消息类型支持：文本、Markdown、图片、文件、图文（news）
- 媒体上传（multipart/form-data）
- 通讯录：用户/部门/标签增改删查
- 应用信息查询（`agent/get`）
- 客户联系：外部联系人列表与详情查询
- 自包含本地模拟测试与真实接口测试（读取环境变量）

## 交流群

<img width="434" height="514" alt="image" src="https://github.com/user-attachments/assets/6d68a770-f0b1-44ff-a52c-0aff79c065c5" />

## 安装
本仓库为本地模块，直接在你的代码中按模块名引用：

```go
import "github.com/Hanson/wecomSdk"
```

Go 版本：`go 1.21`

## 快速开始
```go
package main

import (
    "context"
    "fmt"
    "github.com/Hanson/wecomSdk"
)

func main() {
    c, _ := wecom.NewClient(wecom.Config{
        CorpID:     "你的CorpID",
        CorpSecret: "你的自建应用Secret",
    })
    msg := wecom.TextMessage{ToUser: "zhangsan", AgentID: 1000001}
    msg.Text.Content = "hello from wecom SDK"
    r, err := c.SendTextMessage(context.Background(), msg)
    if err != nil { panic(err) }
    fmt.Println("msgid:", r.MsgID)
}
```

## 使用示例
### 统一调用约定与类型命名
- 显式指定接口路径进行调用：`GetJSON/PostJSON/CallJSON/GetJSONWithReq`
- 请求与返回全部结构化，并按接口路径命名：
  - `/cgi-bin/user/update` → `userUpdateReq` / `userUpdateResp`
  - `/cgi-bin/user/simplelist` → `userSimpleListReq` / `userSimpleListResp`
  - `/cgi-bin/kf/send_msg` → `kfSendMsgReq` / `kfSendMsgResp`
- 错误统一：当返回体内 `errcode!=0` 时，返回 `wecom.Error{Code, Message, Raw}`；`out` 仍会填充原生返回内容。

### 回调服务（URL 验证与消息接收）
```go
cc, _ := wecom.NewCallbackCrypto("你的Token", "你的EncodingAESKey(43位)", "你的ReceiveID")
http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodGet {
    wecom.ValidateURL(w, r, cc)
    return
  }
  wecom.ReceiveMessage(w, r, cc, func(msg []byte) []byte { return nil })
})
```
提示：GET 验证在 1 秒内返回明文；POST 校验 `msg_signature/timestamp/nonce` 并解密 `Encrypt`，需比对 ReceiveID。

### 更新成员（结构化 req/resp + 显式路径）
```go
ctx := context.Background()
var out wecom.userUpdateResp
payload := wecom.userUpdateReq{
  UserID:     "zhangsan",
  Name:       "李四",
  Department: []int{1},
}
err := c.PostJSON(ctx, "/cgi-bin/user/update", payload, &out)
```

### 获取成员简单列表（结构化 GET 请求）
```go
var out wecom.userSimpleListResp
err := c.GetJSONWithReq(ctx, "/cgi-bin/user/simplelist",
  wecom.userSimpleListReq{DepartmentID: 1, FetchChild: 1},
  &out,
)
```

### 微信客服：发送文本消息
```go
req := wecom.kfSendMsgReq{OpenKfID: "kf_123", ToUser: "external_userid", MsgType: "text"}
req.Text.Content = "hello"
var resp wecom.kfSendMsgResp
err := c.PostJSON(ctx, "/cgi-bin/kf/send_msg", req, &resp)
```

- 获取令牌
```go
ctx := context.Background()
tok, err := c.GetAccessToken(ctx)
```

- 发送文本消息
```go
msg := wecom.TextMessage{ToUser: "zhangsan", AgentID: 1000001}
msg.Text.Content = "hello"
r, err := c.SendTextMessage(ctx, msg)
```

- 发送 Markdown
```go
m := wecom.MarkdownMessage{ToUser: "zhangsan", AgentID: 1000001}
m.Markdown.Content = "**Hello**"
_, err := c.SendMarkdownMessage(ctx, m)
```

- 上传图片并发送图片消息
```go
up, err := c.UploadMedia(ctx, "image", "a.png", imgBytes)
img := wecom.ImageMessage{ToUser: "zhangsan", AgentID: 1000001}
img.Image.MediaID = up.MediaID
_, err = c.SendImageMessage(ctx, img)
```

### 用户管理（增/改/删/查）
```go
// 创建
var createOut wecom.userCreateResp
_ = c.PostJSON(ctx, "/cgi-bin/user/create", wecom.userCreateReq{UserID: "u1", Name: "n1", Department: []int{1}}, &createOut)
// 更新
var updateOut wecom.userUpdateResp
_ = c.PostJSON(ctx, "/cgi-bin/user/update", wecom.userUpdateReq{UserID: "u1", Name: "n2"}, &updateOut)
// 删除
var delOut wecom.userDeleteResp
_ = c.GetJSON(ctx, "/cgi-bin/user/delete", url.Values{"userid": []string{"u1"}}, &delOut)
// 查询
var getOut wecom.userGetResp
_ = c.GetJSON(ctx, "/cgi-bin/user/get", url.Values{"userid": []string{"u1"}}, &getOut)
```

### 部门管理（增/改/删/查）
```go
// 创建
var deptCreate wecom.departmentCreateResp
_ = c.PostJSON(ctx, "/cgi-bin/department/create", wecom.departmentCreateReq{Name: "d"}, &deptCreate)
// 更新
var deptUpdate wecom.departmentUpdateResp
_ = c.PostJSON(ctx, "/cgi-bin/department/update", wecom.departmentUpdateReq{ID: deptCreate.ID, Name: "d2"}, &deptUpdate)
// 删除
var deptDelete wecom.departmentDeleteResp
_ = c.GetJSON(ctx, "/cgi-bin/department/delete", url.Values{"id": []string{strconv.Itoa(deptCreate.ID)}}, &deptDelete)
// 查询
var deptList wecom.departmentListResp
_ = c.GetJSON(ctx, "/cgi-bin/department/list", nil, &deptList)
```

### 标签管理与成员维护
```go
var tagCreate wecom.tagCreateResp
_ = c.PostJSON(ctx, "/cgi-bin/tag/create", wecom.tagCreateReq{TagName: "t"}, &tagCreate)
var tagUpdate wecom.tagUpdateResp
_ = c.PostJSON(ctx, "/cgi-bin/tag/update", wecom.tagUpdateReq{TagID: tagCreate.TagID, TagName: "t2"}, &tagUpdate)
var tagList wecom.tagListResp
_ = c.GetJSON(ctx, "/cgi-bin/tag/list", nil, &tagList)
var tagGet wecom.tagGetResp
_ = c.GetJSON(ctx, "/cgi-bin/tag/get", url.Values{"tagid": []string{strconv.Itoa(tagCreate.TagID)}}, &tagGet)
var tagAdd wecom.tagAddTagUsersResp
_ = c.PostJSON(ctx, "/cgi-bin/tag/addtagusers", wecom.tagAddTagUsersReq{TagID: tagCreate.TagID, UserIDs: []string{"u1"}}, &tagAdd)
var tagDel wecom.tagDelTagUsersResp
_ = c.PostJSON(ctx, "/cgi-bin/tag/deltagusers", wecom.tagDelTagUsersReq{TagID: tagCreate.TagID, UserIDs: []string{"u1"}}, &tagDel)
var tagDelete wecom.tagDeleteResp
_ = c.GetJSON(ctx, "/cgi-bin/tag/delete", url.Values{"tagid": []string{strconv.Itoa(tagCreate.TagID)}}, &tagDelete)
```

- 应用信息查询
```go
a, _ := c.GetAgent(ctx, 1000001)
```

### 客户联系（外部联系人）
```go
var extList wecom.externalContactListResp
_ = c.GetJSON(ctx, "/cgi-bin/externalcontact/list", url.Values{"userid": []string{"zhangsan"}}, &extList)
var extGet wecom.externalContactGetResp
_ = c.GetJSON(ctx, "/cgi-bin/externalcontact/get", url.Values{"external_userid": []string{extList.ExternalUserID[0]}}, &extGet)
// 备注
var extRemark wecom.externalRemarkResp
_ = c.PostJSON(ctx, "/cgi-bin/externalcontact/remark", wecom.externalRemarkReq{UserID: "zhangsan", ExternalUserID: extList.ExternalUserID[0], Remark: "VIP"}, &extRemark)
// 跟进人列表
var extFollow wecom.externalFollowUserListResp
_ = c.GetJSON(ctx, "/cgi-bin/externalcontact/get_follow_user_list", nil, &extFollow)
// 企业客户标签
var addTag wecom.externalAddCorpTagResp
_ = c.PostJSON(ctx, "/cgi-bin/externalcontact/add_corp_tag", wecom.externalAddCorpTagReq{GroupName: "客户标签", Tag: []struct{ Name string `json:"name"`; Order int `json:"order,omitempty"` }{{Name: "VIP"}}}, &addTag)
var getTag wecom.externalGetCorpTagResp
_ = c.PostJSON(ctx, "/cgi-bin/externalcontact/get_corp_tag", wecom.externalGetCorpTagReq{}, &getTag)
```

## 真实集成测试
为方便在真实企微环境验证，仓库内提供读取环境变量的测试用例：

- 环境变量
```
WECOM_CORP_ID     企业CorpID
WECOM_APP_SECRET  自建应用Secret
WECOM_AGENT_ID    自建应用AgentID
WECOM_TO_USER     用于接收消息的成员userid（示例：demo_user）
WECOM_USER_ID     内部成员userid，用于外部联系人列表（可与 WECOM_TO_USER 同步）
```

- 可信 IP
  - 必须在管理端为自建应用配置“可信IP”（新创建应用强制要求），否则接口返回错误码 `60020`
  - 可在 PowerShell 查询当前出口公网 IP：
    - `Invoke-RestMethod https://api.ipify.org`

- 运行真实测试（示例，PowerShell）
```powershell
$env:WECOM_CORP_ID="你的CorpID"
$env:WECOM_AGENT_ID="你的AgentID"
$env:WECOM_APP_SECRET="你的Secret"
$env:WECOM_TO_USER="demo_user"
# 实时调用，禁用缓存：
go test -v -run TestRealTokenAndAgent -count=1 github.com/Hanson/wecomSdk
# 发送文本消息到指定成员：
go test -v -run TestRealSendTextMessage -count=1 github.com/Hanson/wecomSdk
# 客户联系（外部联系人）列表与详情，并打印返回：
$env:WECOM_USER_ID="demo_user"
go test -v -run TestRealExternalContact -count=1 github.com/Hanson/wecomSdk
```

## 本地模拟测试
无需外网与真实配置，使用模拟服务验证主要流程：

```bash
go test -v ./...
```
- 集成用例：`integration_test.go`
- 令牌与消息：`client_test.go`

## 注意事项
- 所有接口使用 HTTPS、JSON、UTF-8
- 令牌需缓存并在过期前刷新，避免频繁 `gettoken` 触发频率限制
- 某些接口需要特定令牌类型或权限范围（如通讯录同步助手与自建应用令牌不同）
- 不要在代码仓库提交真实密钥与令牌
 - 微信客服与客户联系模块需要额外权限；若返回 `48002` 表示无接口权限，`60020` 表示可信 IP 未配置

## 版本与更新日志
- 当前版本：`v0.0.1`
- 每日更新详见 `CHANGELOG.md`

## 参考
- 企业微信开放平台：开发前必读、接口调用流程与可信 IP 要求
- 文档入口：`https://developer.work.weixin.qq.com/document/path/90664`
### 其他常用接口
```go
// 批量删除成员
var batchDel wecom.userBatchDeleteResp
_ = c.PostJSON(ctx, "/cgi-bin/user/batchdelete", wecom.userBatchDeleteReq{UserIDList: []string{"u1","u2"}}, &batchDel)
// ID 转换
var toOpen wecom.userConvertToOpenIDResp
_ = c.PostJSON(ctx, "/cgi-bin/user/convert_to_openid", wecom.userConvertToOpenIDReq{UserID: "u1"}, &toOpen)
var toUser wecom.userConvertToUserIDResp
_ = c.PostJSON(ctx, "/cgi-bin/user/convert_to_userid", wecom.userConvertToUserIDReq{OpenID: toOpen.OpenID}, &toUser)
```

### OA 审批/模板
```go
var app wecom.oaApprovalGetDetailResp
_ = c.PostJSON(ctx, "/cgi-bin/oa/getapprovaldetail", wecom.oaApprovalGetDetailReq{SpNo: "SP2025"}, &app)
var tpl wecom.oaApprovalGetTemplateDetailResp
_ = c.PostJSON(ctx, "/cgi-bin/oa/gettemplatedetail", wecom.oaApprovalGetTemplateDetailReq{TemplateID: "TPL2025"}, &tpl)
```

### 考勤打卡
```go
var cd wecom.checkinGetDataResp
_ = c.PostJSON(ctx, "/cgi-bin/checkin/getcheckindata", wecom.checkinGetDataReq{StartTime: 1734393600, EndTime: 1734480000, UserIDList: []string{"u1"}}, &cd)
var dd wecom.checkinGetDayDataResp
_ = c.PostJSON(ctx, "/cgi-bin/checkin/getdaydata", wecom.checkinGetDayDataReq{Date: 20250101, UserIDList: []string{"u1"}}, &dd)
```

### 日历与日程
```go
var calAdd wecom.oaCalendarAddResp
_ = c.PostJSON(ctx, "/cgi-bin/oa/calendar/add", wecom.oaCalendarAddReq{Organizer: "u1", Summary: "项目日历"}, &calAdd)
var schAdd wecom.oaScheduleAddResp
_ = c.PostJSON(ctx, "/cgi-bin/oa/schedule/add", wecom.oaScheduleAddReq{CalID: calAdd.CalID, Organizer: "u1", Summary: "评审", StartTime: 1734400000, EndTime: 1734403600}, &schAdd)
```

### 应用菜单
```go
var mCreate wecom.menuCreateResp
_ = c.PostJSON(ctx, "/cgi-bin/menu/create", wecom.menuCreateReq{AgentID: 1000001, Button: []wecom.menuButton{{Name: "功能", Type: "click", Key: "FUNC"}}}, &mCreate)
var mGet wecom.menuGetResp
_ = c.GetJSON(ctx, "/cgi-bin/menu/get", url.Values{"agentid": []string{"1000001"}}, &mGet)
```

### 素材（二进制获取）
```go
data, apiErr, err := c.GetRaw(ctx, "/cgi-bin/media/get", url.Values{"media_id": []string{"MID"}})
```
