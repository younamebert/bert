package tools

import "time"

type Tabname struct {
	Query string `json:"query"`
}

func Rday() int64 {
	r := time.Now().Unix()
	return r
}
