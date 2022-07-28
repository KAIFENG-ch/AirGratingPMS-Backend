// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	EnterpriseId int64  `form:"enterprise_id"`
	Username     string `form:"username"`
	Password     string `form:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type Staffer struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	WorkshopId  int64  `json:"workshop_id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Remark      string `json:"remark"`
}

type CreateStafferReq struct {
	Username    string `json:"username"`
	WorkshopId  int64  `json:"workshop_id"`
	Name        string `json:"name"`
	Role        string `json:"role,options=employee|manager"`
	Gender      string `json:"gender,options=male|female"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Remark      string `json:"remark"`
}

type ChangeStafferInfoReq struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	WorkshopId  int64  `json:"workshop_id"`
	Name        string `json:"name"`
	Role        string `json:"role,options=employee|manager"`
	Gender      string `json:"gender,optional,options=male|female"`
	PhoneNumber string `json:"phone_number,optional"`
	Email       string `json:"email,optional"`
	Address     string `json:"address,optional"`
	Remark      string `json:"remark,optional"`
}

type PartialChangeStafferInfoReq struct {
	Id          int64  `json:"id"`
	Username    string `json:"username,optional"`
	WorkshopId  int64  `json:"workshop_id,optional"`
	Name        string `json:"name"`
	Role        string `json:"role,optional,options=employee|manager"`
	Gender      string `json:"gender,optional,options=male|female"`
	PhoneNumber string `json:"phone_number,optional"`
	Email       string `json:"email,optional"`
	Address     string `json:"address,optional"`
	Remark      string `json:"remark,optional"`
}

type CreateReply struct {
	Message  string `json:"message"`
	Password string `json:"password"`
}

type ChangeStafferInfoReply struct {
	Message string `json:"message"`
}

type PartialChangeStafferInfoReply struct {
	Message       string   `json:"message"`
	ChangedFields []string `json:"changed_fields"`
}

type GetStafferListReq struct {
	PageNumber int32 `form:"page_number,default=1"`
	PageSize   int32 `form:"page_size,default=20"`
}

type GetStafferListReply struct {
	Message     string    `json:"message"`
	StafferList []Staffer `json:"staffer_list"`
}

type GetStafferListByWorkshopReq struct {
	WorkshopId int64 `form:"workshop_id"`
	PageNumber int32 `form:"page_number,default=1"`
	PageSize   int32 `form:"page_size,default=20"`
}

type GetStafferListByWorkshopReply struct {
	Message     string    `json:"message"`
	StafferList []Staffer `json:"staffer_list"`
}

type ChangeStafferPasswordReq struct {
	OldPassword string `form:"old_password"`
	NewPassword string `form:"new_password"`
}

type ChangeStafferPasswordReply struct {
	Message string `json:"message"`
}

type ResetStafferPasswordReq struct {
	Id int64 `form:"id"`
}

type ResetStafferPasswordReply struct {
	Message  string `json:"message"`
	Password string `json:"password"`
}
