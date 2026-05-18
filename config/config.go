package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token   string
	BaseUrl string
)

func init() {
	_ = godotenv.Load()

	Token = os.Getenv("BOT_TOKEN")
	if Token == "" {
		panic("BOT_TOKEN environment variable is required")
	}
	BaseUrl = os.Getenv("BASE_URL")
	if BaseUrl == "" {
		BaseUrl = "https://www.kookapp.cn/api"
	}
	fmt.Printf("Config loaded: BaseUrl=%s, Token=%s...%s\n", BaseUrl, Token[:4], Token[len(Token)-4:])
}