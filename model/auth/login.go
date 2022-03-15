package auth

type Data struct {
	Name        string `json:"name"`         //用户名
	Phone       string `json:"phone"`        //用户手机号
	Email       string `json:"email"`        //用户邮箱
	Password    string `json:"password"`     //用户密码
	NewPassword string `json:"new_password"` //更替的新密码
	Code        string `json:"code"`         //验证码
}
