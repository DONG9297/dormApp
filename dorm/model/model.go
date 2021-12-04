package model

// User 用户格式
type User struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Phone     string `json:"phone"`
	StudentID string `json:"stu_id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
}
type UserDorm struct {
	UserID int
	DormID int
}

type Building struct {
	ID   int
	Name string
}
type Unit struct {
	ID         int
	Name       string
	BuildingID int
}
type Dorm struct {
	ID            int
	Name          string
	Gender        string
	TotalBeds     int
	AvailableBeds int
	UnitID        int
}

type DormInfo struct {
	BuildingName  string `json:"building_name"`
	UnitName      string `json:"unit_name"`
	DormName      string `json:"dorm_name"`
	DormID        int    `json:"dorm_id"`
	AvailableBeds int    `json:"available_beds"`
}

type DormListItem struct {
	BuildingName  string `json:"building_name"`
	AvailableBeds int    `json:"available_beds"`
	Count         int    `json:"count"`
}

type Order struct {
	ID         string
	UserID     int
	Count      int
	BuildingID int
	Gender     string
	CreateTime string
	State      int // 0 未处理， 1 成功， 2 失败
}
type OrderItem struct {
	ID      int
	OrderID string
	UserID  int
}
