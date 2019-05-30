package pprint

import (
	"testing"
	"time"
)

type Policy struct {
	Allow string
	Deny  string
}

type Host struct {
	ip     string
	alias  []string
	policy Policy
}
type Redis struct {
	name string // lower-case
	port uint   // lower-case
	host Host
}

type Web struct {
	Host    string
	port    int32 // lower-case
	Timeout time.Duration
	Rate    float32
	Score   []float32
	IP      []string
	MySQL *struct {
		Name string
		port int64 // lower-case
	}
	redis *Redis
}

func TestFormat(t *testing.T) {
	w := &Web{
		Host:    "web host",
		port:    1234,
		Timeout: 5 * time.Second,
		Rate:    0.32,
		Score:   []float32{1.1, 2.2, 3.3},
		IP:      []string{"192.168.1.1", "127.0.0.1", "localhost"},
/*		MySQL: &struct {
			Name string
			port int64
		}{Name: "mysqldb", port: 3306},*/
		redis: &Redis{"rdb", 6379, Host{"adf", []string{"alias1", "alias2"}, Policy{"allow policy", "deny policy"}}},
	}
	// 调用pprint 格式化输出
	Format(w)
}
