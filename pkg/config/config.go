package config

import "os"

type Config struct {
	SMTPConfig    SMTPConfig
	MailjetConfig MailjetConfig
}

type SMTPConfig struct {
	SMTPHost     string
	SMTPUser     string
	SMTPPassword string
}

type MailjetConfig struct {
	Key1 string
	Key2 string
}

func NewConfig() Config {
	return Config{
		SMTPConfig: SMTPConfig{
			SMTPHost:     os.Getenv("SMPT_HOST"),
			SMTPUser:     os.Getenv("SMTP_USER"),
			SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		},
		MailjetConfig: MailjetConfig{
			Key1: os.Getenv("MAILJET_KEY1"),
			Key2: os.Getenv("MAILJET_KEY2"),
		},
	}
}
