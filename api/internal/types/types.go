// Code generated by goctl. DO NOT EDIT.
package types

type Req struct {
	Id int64 `path:"id"`
}

type Reply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
