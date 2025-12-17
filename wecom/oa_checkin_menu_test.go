package wecom

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"
)

func TestOAAndMenuAndRaw(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/cgi-bin/gettoken", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "access_token": "T", "expires_in": 7200})
    })
    mux.HandleFunc("/cgi-bin/oa/getapprovaldetail", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "approval_info": map[string]any{"sp_no": "sp1"}})
    })
    mux.HandleFunc("/cgi-bin/oa/gettemplatedetail", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "template": map[string]any{"template_id": "tpl1"}})
    })
    mux.HandleFunc("/cgi-bin/checkin/getcheckindata", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "checkindata": []map[string]any{{"userid": "u1"}}})
    })
    mux.HandleFunc("/cgi-bin/checkin/getdaydata", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "daydata": []map[string]any{{"userid": "u1"}}})
    })
    mux.HandleFunc("/cgi-bin/oa/calendar/add", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "cal_id": "cal1"})
    })
    mux.HandleFunc("/cgi-bin/oa/calendar/get", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "calendar": map[string]any{"cal_id": "cal1"}})
    })
    mux.HandleFunc("/cgi-bin/oa/calendar/update", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/oa/calendar/del", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/oa/schedule/add", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "schedule_id": "sch1"})
    })
    mux.HandleFunc("/cgi-bin/oa/schedule/get", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "schedule": map[string]any{"schedule_id": "sch1"}})
    })
    mux.HandleFunc("/cgi-bin/oa/schedule/update", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/oa/schedule/del", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/menu/create", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/menu/delete", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok"})
    })
    mux.HandleFunc("/cgi-bin/menu/get", func(w http.ResponseWriter, r *http.Request) {
        _ = json.NewEncoder(w).Encode(map[string]any{"errcode": 0, "errmsg": "ok", "menu": map[string]any{"button": []map[string]any{{"name": "A"}}}})
    })
    mux.HandleFunc("/cgi-bin/media/get", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/octet-stream")
        _, _ = w.Write([]byte("BIN"))
    })
    srv := httptest.NewServer(mux)
    defer srv.Close()
    c, _ := NewClient(Config{CorpID: "id", CorpSecret: "secret", BaseURL: srv.URL})
    ctx := context.Background()
    var app oaApprovalGetDetailResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/getapprovaldetail", oaApprovalGetDetailReq{SpNo: "sp1"}, &app)
    var tpl oaApprovalGetTemplateDetailResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/gettemplatedetail", oaApprovalGetTemplateDetailReq{TemplateID: "tpl1"}, &tpl)
    var cd checkinGetDataResp
    _ = c.PostJSON(ctx, "/cgi-bin/checkin/getcheckindata", checkinGetDataReq{StartTime: 1, EndTime: 2, UserIDList: []string{"u1"}}, &cd)
    var dd checkinGetDayDataResp
    _ = c.PostJSON(ctx, "/cgi-bin/checkin/getdaydata", checkinGetDayDataReq{Date: 20240101, UserIDList: []string{"u1"}}, &dd)
    var calAdd oaCalendarAddResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/calendar/add", oaCalendarAddReq{Organizer: "u1", Summary: "S"}, &calAdd)
    var calGet oaCalendarGetResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/calendar/get", oaCalendarGetReq{CalID: calAdd.CalID}, &calGet)
    var calUpd oaCalendarUpdateResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/calendar/update", oaCalendarUpdateReq{CalID: calAdd.CalID, Summary: "S2"}, &calUpd)
    var calDel oaCalendarDelResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/calendar/del", oaCalendarDelReq{CalID: calAdd.CalID}, &calDel)
    var schAdd oaScheduleAddResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/schedule/add", oaScheduleAddReq{CalID: calAdd.CalID, Organizer: "u1", Summary: "SS", StartTime: 1, EndTime: 2}, &schAdd)
    var schGet oaScheduleGetResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/schedule/get", oaScheduleGetReq{ScheduleID: schAdd.ScheduleID}, &schGet)
    var schUpd oaScheduleUpdateResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/schedule/update", oaScheduleUpdateReq{ScheduleID: schAdd.ScheduleID, Summary: "SS2"}, &schUpd)
    var schDel oaScheduleDelResp
    _ = c.PostJSON(ctx, "/cgi-bin/oa/schedule/del", oaScheduleDelReq{ScheduleID: schAdd.ScheduleID}, &schDel)
    var mCreate menuCreateResp
    _ = c.PostJSON(ctx, "/cgi-bin/menu/create", menuCreateReq{AgentID: 1000001, Button: []menuButton{{Name: "A", Type: "click", Key: "K"}}}, &mCreate)
    var mDel menuDeleteResp
    _ = c.GetJSON(ctx, "/cgi-bin/menu/delete", url.Values{"agentid": []string{"1000001"}}, &mDel)
    var mGet menuGetResp
    _ = c.GetJSON(ctx, "/cgi-bin/menu/get", url.Values{"agentid": []string{"1000001"}}, &mGet)
    b, e, err := c.GetRaw(ctx, "/cgi-bin/media/get", nil)
    if err != nil { t.Fatal(err) }
    if e != nil { t.Fatal(e) }
    if string(b) != "BIN" { t.Fatalf("unexpected raw: %s", string(b)) }
}

