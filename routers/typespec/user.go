package typespec

type User struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	NicName string `form:"nickName" json:"nickName"`
	Sex     string `form:"sex" json:"sex"`
	Age     int    `form:"age" json:"age"`
}

//	Id      int64  `form:"id" json:"id"`
type AddUserRequest struct {
	User
}

type AddUserResponse struct {
	Id int64 `json:"id"`
}

type GetUserRequest struct {
	Id int64 `form:"id" json:"id"`
}

type GetUserResponse struct {
	Id int64 `json:"id"`
	User
	CreateOn string `json:"createOn"`
	ModifyOn string `json:"modifyOn"`
}

type GetUserListRequest struct {
	Offset int `form:"offset,default=0"`  // 分页偏移
	Length int `form:"length,default=10"` // 分页每页显示条数
	User
}

// GetVersionListResponse -
type GetUserListResponse struct {
	Offset int               `json:"offset"` // 分页偏移
	Length int               `json:"length"` // 分页每页显示条数
	Total  int64             `json:"total"`  // 总条数
	List   []GetUserResponse `json:"list"`
}
