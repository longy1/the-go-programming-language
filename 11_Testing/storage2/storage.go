package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 {
	return 9.8e8
}

const sender = "notification@example.com"
const password = "password"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

func checkQuota(username string) {
	used := bytesInUse(username)
	const quota = 1e9 // 1 GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
