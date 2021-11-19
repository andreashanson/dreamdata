package config

import "os"

type Config struct {
	SMTPConfig    *SMTPConfig
	MailjetConfig *MailjetConfig
}

type SMTPConfig struct {
	SMTPHost     string `env:"SMTP_HOST"`
	SMTPUser     string `env:"SMTP_USER"`
	SMTPPassword string `env:"SMTP_PASSWORD"`
}

type MailjetConfig struct {
	Key1 string `env:"MAILJET_KEY1"`
	Key2 string `env:"MAILJET_KEY2"`
}

func NewConfig() *Config {
	return &Config{
		SMTPConfig: &SMTPConfig{
			SMTPHost:     os.Getenv("SMTP_HOST"),
			SMTPUser:     os.Getenv("SMTP_USER"),
			SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		},
		MailjetConfig: &MailjetConfig{
			Key1: os.Getenv("MAILJET_KEY1"),
			Key2: os.Getenv("MAILJET_KEY2"),
		},
	}
}
