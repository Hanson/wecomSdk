package wecom

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strconv"
    "testing"
)

func mockServerAll() *httptest.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200}) })
    mux.HandleFunc("/cgi-bin/message/send", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "msgid": "m1"}) })
    mux.HandleFunc("/cgi-bin/media/upload", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "type": "image", "media_id": "mid"}) })
    mux.HandleFunc("/cgi-bin/user/create", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/user/update", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/user/delete", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/user/get", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "userid": "u1", "name": "n1"}) })
    mux.HandleFunc("/cgi-bin/department/create", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "id": 2}) })
    mux.HandleFunc("/cgi-bin/department/update", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/department/delete", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/department/list", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "department": []map[string]any{{"id": 1, "name": "root"}}}) })
    mux.HandleFunc("/cgi-bin/tag/create", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "tagid": 1}) })
    mux.HandleFunc("/cgi-bin/tag/update", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/tag/delete", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/tag/list", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "taglist": []map[string]any{{"tagid": 1, "tagname": "t"}}}) })
    mux.HandleFunc("/cgi-bin/tag/get", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "userlist": []map[string]any{{"userid": "u1"}}}) })
    mux.HandleFunc("/cgi-bin/tag/addtagusers", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/tag/deltagusers", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/agent/get", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "agentid": 1000001, "name": "app"}) })
    mux.HandleFunc("/cgi-bin/externalcontact/list", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "external_userid": []string{"ext1", "ext2"}}) })
    mux.HandleFunc("/cgi-bin/externalcontact/get", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "external_contact": map[string]any{"external_userid": "ext1", "name": "alice"}}) })
    mux.HandleFunc("/cgi-bin/externalcontact/remark", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/externalcontact/get_follow_user_list", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "follow_user": []string{"zhangsan", "lisi"}}) })
    mux.HandleFunc("/cgi-bin/externalcontact/add_corp_tag", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "tag_group": map[string]any{"group_id": "g1", "group_name": "G", "tag": []map[string]any{{"id": "t1", "name": "VIP"}}}}) })
    mux.HandleFunc("/cgi-bin/externalcontact/edit_corp_tag", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/externalcontact/del_corp_tag", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/externalcontact/get_corp_tag", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "tag_group": []map[string]any{{"group_id": "g1", "group_name": "G", "tag": []map[string]any{{"id": "t1", "name": "VIP"}}}}}) })
    mux.HandleFunc("/cgi-bin/user/batchdelete", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"}) })
    mux.HandleFunc("/cgi-bin/user/convert_to_openid", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "openid": "o123"}) })
    mux.HandleFunc("/cgi-bin/user/convert_to_userid", func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "userid": "u123"}) })
    return httptest.NewServer(mux)
}

func TestAllAPIs(t *testing.T) {
    srv := mockServerAll(); defer srv.Close()
    c, err := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL}); if err != nil { t.Fatal(err) }
    ctx := context.Background()
    _, err = c.SendTextMessage(ctx, TextMessage{AgentID: 1}); if err != nil { t.Fatal(err) }
    _, err = c.SendMarkdownMessage(ctx, MarkdownMessage{AgentID: 1}); if err != nil { t.Fatal(err) }
    _, err = c.SendImageMessage(ctx, ImageMessage{AgentID: 1}); if err != nil { t.Fatal(err) }
    _, err = c.SendNewsMessage(ctx, NewsMessage{AgentID: 1}); if err != nil { t.Fatal(err) }
    _, err = c.UploadMedia(ctx, "image", "a.png", []byte("data")); if err != nil { t.Fatal(err) }
    var userCreateOut UserCreateResp
    err = c.PostJSON(ctx, "/cgi-bin/user/create", UserCreateReq{UserID: "u1", Name: "n1", Department: []int{1}}, &userCreateOut); if err != nil { t.Fatal(err) }
    var userUpdateOut UserUpdateResp
    err = c.PostJSON(ctx, "/cgi-bin/user/update", UserUpdateReq{UserID: "u1", Name: "n2"}, &userUpdateOut); if err != nil { t.Fatal(err) }
    var userDeleteOut UserDeleteResp
    err = c.GetJSON(ctx, "/cgi-bin/user/delete", url.Values{"userid": []string{"u1"}}, &userDeleteOut); if err != nil { t.Fatal(err) }
    var userGet UserGetResp
    err = c.GetJSON(ctx, "/cgi-bin/user/get", url.Values{"userid": []string{"u1"}}, &userGet); if err != nil || userGet.User.UserID != "u1" { t.Fatal(err) }
    var deptCreateOut DepartmentCreateResp
    err = c.PostJSON(ctx, "/cgi-bin/department/create", DepartmentCreateReq{Name: "d"}, &deptCreateOut); if err != nil || deptCreateOut.ID != 2 { t.Fatal(err) }
    var deptUpdateOut DepartmentUpdateResp
    err = c.PostJSON(ctx, "/cgi-bin/department/update", DepartmentUpdateReq{ID: deptCreateOut.ID, Name: "d2"}, &deptUpdateOut); if err != nil { t.Fatal(err) }
    var deptDeleteOut DepartmentDeleteResp
    err = c.GetJSON(ctx, "/cgi-bin/department/delete", url.Values{"id": []string{strconv.Itoa(deptCreateOut.ID)}}, &deptDeleteOut); if err != nil { t.Fatal(err) }
    var deptListOut DepartmentListResp
    err = c.GetJSON(ctx, "/cgi-bin/department/list", nil, &deptListOut); if err != nil || len(deptListOut.Department) == 0 { t.Fatal(err) }
    var tagCreateOut TagCreateResp
    err = c.PostJSON(ctx, "/cgi-bin/tag/create", TagCreateReq{TagName: "t"}, &tagCreateOut); if err != nil || tagCreateOut.TagID != 1 { t.Fatal(err) }
    var tagUpdateOut TagUpdateResp
    err = c.PostJSON(ctx, "/cgi-bin/tag/update", TagUpdateReq{TagID: tagCreateOut.TagID, TagName: "t2"}, &tagUpdateOut); if err != nil { t.Fatal(err) }
    var tagDeleteOut TagDeleteResp
    err = c.GetJSON(ctx, "/cgi-bin/tag/delete", url.Values{"tagid": []string{strconv.Itoa(tagCreateOut.TagID)}}, &tagDeleteOut); if err != nil { t.Fatal(err) }
    var tagListOut TagListResp
    err = c.GetJSON(ctx, "/cgi-bin/tag/list", nil, &tagListOut); if err != nil { t.Fatal(err) }
    var tagGetOut TagGetResp
    err = c.GetJSON(ctx, "/cgi-bin/tag/get", url.Values{"tagid": []string{strconv.Itoa(tagCreateOut.TagID)}}, &tagGetOut); if err != nil { t.Fatal(err) }
    var tagAddUsersOut TagAddTagUsersResp
    err = c.PostJSON(ctx, "/cgi-bin/tag/addtagusers", TagAddTagUsersReq{TagID: tagCreateOut.TagID, UserIDs: []string{"u1"}}, &tagAddUsersOut); if err != nil { t.Fatal(err) }
    var tagDelUsersOut TagDelTagUsersResp
    err = c.PostJSON(ctx, "/cgi-bin/tag/deltagusers", TagDelTagUsersReq{TagID: tagCreateOut.TagID, UserIDs: []string{"u1"}}, &tagDelUsersOut); if err != nil { t.Fatal(err) }
    var agentOut AgentGetResp
    err = c.GetJSON(ctx, "/cgi-bin/agent/get", url.Values{"agentid": []string{"1000001"}}, &agentOut); if err != nil || agentOut.AgentID != 1000001 { t.Fatal(err) }
    var extList ExternalContactListResp
    err = c.GetJSON(ctx, "/cgi-bin/externalcontact/list", url.Values{"userid": []string{"u1"}}, &extList); if err != nil { t.Fatal(err) }
    var extGet ExternalContactGetResp
    err = c.GetJSON(ctx, "/cgi-bin/externalcontact/get", url.Values{"external_userid": []string{"ext1"}}, &extGet); if err != nil || extGet.ExternalContact.ExternalUserID != "ext1" { t.Fatal(err) }
    var extRemark ExternalRemarkResp
    err = c.PostJSON(ctx, "/cgi-bin/externalcontact/remark", ExternalRemarkReq{UserID: "u1", ExternalUserID: "ext1", Remark: "VIP"}, &extRemark); if err != nil { t.Fatal(err) }
    var extFollow ExternalFollowUserListResp
    err = c.GetJSON(ctx, "/cgi-bin/externalcontact/get_follow_user_list", nil, &extFollow); if err != nil || len(extFollow.FollowUser) == 0 { t.Fatal(err) }
    var addTag ExternalAddCorpTagResp
    err = c.PostJSON(ctx, "/cgi-bin/externalcontact/add_corp_tag", ExternalAddCorpTagReq{GroupName: "G", Tag: []struct{ Name string `json:"name"`; Order int `json:"order,omitempty"` }{{Name: "VIP"}}}, &addTag); if err != nil || addTag.TagGroup.GroupID == "" { t.Fatal(err) }
    var editTag ExternalEditCorpTagResp
    err = c.PostJSON(ctx, "/cgi-bin/externalcontact/edit_corp_tag", ExternalEditCorpTagReq{ID: "t1", Name: "SVIP"}, &editTag); if err != nil { t.Fatal(err) }
    var delTag ExternalDelCorpTagResp
    err = c.PostJSON(ctx, "/cgi-bin/externalcontact/del_corp_tag", ExternalDelCorpTagReq{TagID: []string{"t1"}}, &delTag); if err != nil { t.Fatal(err) }
    var getTag ExternalGetCorpTagResp
    err = c.PostJSON(ctx, "/cgi-bin/externalcontact/get_corp_tag", ExternalGetCorpTagReq{}, &getTag); if err != nil { t.Fatal(err) }
    var batchDel UserBatchDeleteResp
    err = c.PostJSON(ctx, "/cgi-bin/user/batchdelete", UserBatchDeleteReq{UserIDList: []string{"u1","u2"}}, &batchDel); if err != nil { t.Fatal(err) }
    var toOpen UserConvertToOpenIDResp
    err = c.PostJSON(ctx, "/cgi-bin/user/convert_to_openid", UserConvertToOpenIDReq{UserID: "u1"}, &toOpen); if err != nil || toOpen.OpenID == "" { t.Fatal(err) }
    var toUser UserConvertToUserIDResp
    err = c.PostJSON(ctx, "/cgi-bin/user/convert_to_userid", UserConvertToUserIDReq{OpenID: "o123"}, &toUser); if err != nil || toUser.UserID == "" { t.Fatal(err) }
}

func BenchmarkSendText(b *testing.B) {
    srv := mockServerAll(); defer srv.Close()
    c, _ := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
    ctx := context.Background()
    for i := 0; i < b.N; i++ { _, _ = c.SendTextMessage(ctx, TextMessage{AgentID: 1}) }
}
