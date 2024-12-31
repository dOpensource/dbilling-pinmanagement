package dbilling_pinmanagement

type Pin struct {
	Id              uint32 `json:"pin_id" form:"route_id"`
	Pin             string `json:"pin" form:"redirect_did" binding:"required"`
	MinutesPerMonth string `json:"minutespermonth" form:"trunk" binding:"required"`
	UserId          uint32 `json:"user_id" form:"trunk" binding:"required"`
	Status          string `json:"status" form:"trunk"`
}
