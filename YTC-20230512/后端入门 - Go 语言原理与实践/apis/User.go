package apis

// 实现user结构体
type User struct {
	Name     string
	Password string
}

func (u User) CheckPassword(password string) bool {
	return u.Password == password
}

func (u *User) ResetPassword(password string) {
	u.Password = password
}
