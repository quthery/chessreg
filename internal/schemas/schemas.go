package schemas

type UserInsert struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}
