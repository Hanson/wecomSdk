package wecom

import "strconv"

type Department struct {
    ID       int    `json:"id,omitempty"`
    Name     string `json:"name"`
    ParentID int    `json:"parentid,omitempty"`
    Order    int    `json:"order,omitempty"`
}

type deptListResponse struct { APIError; Department []Department `json:"department"` }

func intToString(n int) string { return strconv.Itoa(n) }
