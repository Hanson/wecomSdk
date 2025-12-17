package wecom

type Tag struct { TagID int `json:"tagid,omitempty"`; TagName string `json:"tagname"` }
type tagListResponse struct { APIError; TagList []Tag `json:"taglist"` }
type tagUsersResponse struct { APIError; UserList []struct { UserID string `json:"userid"`; Name string `json:"name,omitempty"` } `json:"userlist"` }
