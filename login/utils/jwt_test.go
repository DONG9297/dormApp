package utils

import (
	"login/model"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := &model.User{
		ID:        1,
		UID:       "12345678",
		Phone:     "18312345678",
		StudentID: "1234567890",
		Name:      "张三",
		Gender:    "男",
	}
	token, err := GenerateToken(user)
	if err != nil {
		println(err)
	}
	println(token)
}

func TestParseToken(t *testing.T) {
	user, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzg2MjA4NzIsImdlbmRlciI6IueUtyIsImlhdCI6MTYzODUzNDQ3MiwiaWQiOjEsIm5hbWUiOiLlvKDkuIkiLCJwaG9uZSI6IjE4MzEyMzQ1Njc4Iiwic3R1X2lkIjoiMTIzNDU2Nzg5MCIsInVpZCI6IjEyMzQ1Njc4In0.P32qsEw__3JrcXVUsfXjmBOUoWFNT8v-4eEjxyAchDc")
	if err != nil {
		println(err)
	}
	println(user.ID)
	println(user.UID)
	println(user.Phone)
	println(user.StudentID)
	println(user.Name)
	println(user.Gender)
}
