package network

import (
	"encoding/json"
	"net/http"
)

// 公网IP先关

// 查询结果结构体
// 采用接口： 进行查询
type Query struct {
	Success string  `json:"success"`
	Result  *Result `json:"result"`
}
type Result struct {
	IP        string `json:"ip"`
	Proxy     string `json:"proxy"`
	Att       string `json:"att"`
	Operators string `json:"operators"`
}

// 查询公网IP
func PublicIPInfo() (*Query, error) {
	rsp, err := http.Get("http://api.k780.com/?app=ip.local&format=json")
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var query Query
	if err := json.NewDecoder(rsp.Body).Decode(&query); err != nil {
		return nil, err
	}
	return &query, nil
}
