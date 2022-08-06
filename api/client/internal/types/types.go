// Code generated by goctl. DO NOT EDIT.
package types

type InsertClientReq struct {
	WorkshopId  int64  `json:"workshop_id,optional"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number,optional"`
	Email       string `json:"email,optional"`
	Address     string `json:"address,optional"`
	Remark      string `json:"remark,optional"`
}

type InsertClientReply struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type DeleteCientReq struct {
	Id int64 `json:"id"`
}

type DeleteCientReply struct {
	Message string `json:"message"`
}

type UpdateClientReq struct {
	Id          int64  `json:"id"`
	WorkshopId  int64  `json:"workshop_id,optional"`
	Name        string `json:"name,optional"`
	PhoneNumber string `json:"phone_number,optional"`
	Email       string `json:"email,optional"`
	Address     string `json:"address,optional"`
	Remark      string `json:"remark,optional"`
}

type UpdateClientReply struct {
	Message string `json:"message"`
}

type GetClientInfoByIdReq struct {
	Id int64 `json:"id"`
}

type ClientInfo struct {
	Id          int64  `json:"id"`
	WorkshopId  int64  `json:"workshop_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Remark      string `json:"remark"`
	Version     int64  `json:"version"`
}

type GetClientListByWorkshopReq struct {
	WorkshopId int64 `json:"workshop_id"`
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetClientListByWorkshopReply struct {
	Count   int64        `json:"count"`
	List    []ClientInfo `json:"list"`
	Message string       `json:"message"`
}

type GetClientListByEnterpriseReq struct {
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetClientListByEnterpriseReply struct {
	Count   int64        `json:"count"`
	List    []ClientInfo `json:"list"`
	Message string       `json:"message"`
}

type SetUnitPriceByClientReq struct {
	ClientId     int64           `json:"client_id"`
	UnitPriceMap map[int64]int64 `json:"unit_price_map"`
}

type SetUnitPriceByClientReply struct {
	Message string `json:"message"`
}

type GetUnitPriceByClientReq struct {
	ClientId int64 `json:"client_id"`
}

type GetUnitPriceByClientReply struct {
	Message      string          `json:"message"`
	UnitPriceMap map[int64]int64 `json:"unit_price_map"`
}
