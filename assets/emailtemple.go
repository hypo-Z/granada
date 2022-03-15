package assets

func NewEmailBody(code string, codeExp string, templateId string) string {
	template, ok := Templates[templateId]
	if !ok {
		return ""
	}
	return template[0] + codeExp + template[1] + code + template[2]
}

var Templates = map[string][]string{
	"1": {
		`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>用户激活邮件</title>
	</head>
	<body>
	<div class="box" style="width: 658px;padding-top: 10px;;margin: 0 auto;background: #f7f8fa;border-radius: 10px;">

        <div class="content" style="width: 90%;margin: 0 auto;padding: 20px 20px 0 20px;background: white;border-radius: 10px;">

            <div class="content-logo" style=" width: 50px;height: 50px;margin: 15px auto;">

                <img src="https://preactjs.com/favicon.ico" alt="" style="width: 100%;height: 100%;">

            </div>

            <div class="content-text" style=" width: 100%; padding: 10px 0;border-top: 2px solid skyblue;">

                <h2 style=" margin: 5px 0 10px 0; text-align: center;">电子邮箱验证码</h2>

                <p style="text-align: center;">请在 <span style="color: black;font-weight: 600;"> `, `分钟内

                    </span>输入您的电子邮箱验证码。</p>

                <p style=" margin: 5px 0;height: 50px;line-height: 50px;text-align: center;"><span

                        style="color: #673ab8; font-size: 32px;"> `, ` </span></p>

            </div>

        </div>

        <div style="text-align: center;padding: 20px 0 ;">此为系统邮件，请勿回复。</div>

    </div>
	</body>
	</html>`,
	},
}
