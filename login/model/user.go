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
