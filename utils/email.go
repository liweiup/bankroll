package utils

import (
	"bankroll/global"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)
func SendEmail(subject,content string,) {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = global.Config.Email.From
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	//em.To = []string{global.Config.Email.To,"365625376@qq.com"}
	em.To = []string{global.Config.Email.To}
	// 设置主题
	em.Subject = subject
	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(content)
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", global.Config.Email.From, global.Config.Email.Secret, "smtp.qq.com"))
	if err != nil {
		log.Println("send error ... " + err.Error())
	}
	log.Println("send successfully ... ")
}
