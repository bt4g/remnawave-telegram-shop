package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
)

type config struct {
	telegramToken          string
	price1                 int
	price3                 int
	price6                 int
	remnawaveUrl           string
	remnawaveToken         string
	remnawaveMode          string
	databaseURL            string
	cryptoPayURL           string
	cryptoPayToken         string
	botURL                 string
	yookasaURL             string
	yookasaShopId          string
	yookasaSecretKey       string
	yookasaEmail           string
	countries              map[string]string
	trafficLimit           int64
	feedbackURL            string
	channelURL             string
	serverStatusURL        string
	supportURL             string
	isYookasaEnabled       bool
	isCryptoEnabled        bool
	isTelegramStarsEnabled bool
	adminTelegramId        int64
	trialDays              int
	trialTrafficLimit      int64
}

var conf config

func TrialTrafficLimit() int64 {
	return conf.trialTrafficLimit * bytesInGigabyte
}

func TrialDays() int {
	return conf.trialDays
}
func FeedbackURL() string {
	return conf.feedbackURL
}

func ChannelURL() string {
	return conf.channelURL
}

func ServerStatusURL() string {
	return conf.serverStatusURL
}

func SupportURL() string {
	return conf.supportURL
}

func YookasaEmail() string {
	return conf.yookasaEmail
}
func Price1() int {
	return conf.price1
}
func Price3() int {
	return conf.price3
}
func Price6() int {
	return conf.price6
}
func TelegramToken() string {
	return conf.telegramToken
}
func RemnawaveUrl() string {
	return conf.remnawaveUrl
}
func DadaBaseUrl() string {
	return conf.databaseURL
}
func RemnawaveToken() string {
	return conf.remnawaveToken
}
func RemnawaveMode() string {
	return conf.remnawaveMode
}
func CryptoPayUrl() string {
	return conf.cryptoPayURL
}
func CryptoPayToken() string {
	return conf.cryptoPayToken
}
func BotURL() string {
	return conf.botURL
}
func SetBotURL(botURL string) {
	conf.botURL = botURL
}
func YookasaUrl() string {
	return conf.yookasaURL
}
func YookasaShopId() string {
	return conf.yookasaShopId
}
func YookasaSecretKey() string {
	return conf.yookasaSecretKey
}
func SetCountries(countries map[string]string) {
	conf.countries = countries
}
func Countries() map[string]string {
	return conf.countries
}
func TrafficLimit() int64 {
	return conf.trafficLimit * bytesInGigabyte
}

func IsCryptoPayEnabled() bool {
	return conf.isCryptoEnabled
}

func IsYookasaEnabled() bool {
	return conf.isYookasaEnabled
}

func IsTelegramStarsEnabled() bool {
	return conf.isTelegramStarsEnabled
}

func GetAdminTelegramId() int64 {
	return conf.adminTelegramId
}

const bytesInGigabyte = 1073741824

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("Env file not found")
	}

	conf.adminTelegramId, err = strconv.ParseInt(os.Getenv("ADMIN_TELEGRAM_ID"), 10, 64)
	if err != nil {
		panic("ADMIN_TELEGRAM_ID .env variable not set")
	}

	conf.telegramToken = os.Getenv("TELEGRAM_TOKEN")
	if conf.telegramToken == "" {
		panic("TELEGRAM_TOKEN .env variable not set")
	}

	conf.trialTrafficLimit, err = strconv.ParseInt(os.Getenv("TRIAL_TRAFFIC_LIMIT"), 10, 64)
	if err != nil {
		panic("TRIAL_TRAFFIC_LIMIT .env variable not set")
	}

	conf.trialDays, err = strconv.Atoi(os.Getenv("TRIAL_DAYS"))
	if err != nil {
		panic("TRIAL_DAYS .env variable not set")
	}

	strPrice := os.Getenv("PRICE_1")
	if strPrice == "" {
		panic("PRICE_1 .env variable not set")
	}
	price, err := strconv.Atoi(strPrice)
	if err != nil {
		panic(err)
	}
	conf.price1 = price

	strPrice3 := os.Getenv("PRICE_3")
	if strPrice3 == "" {
		panic("PRICE_3 .env variable not set")
	}
	price3, err := strconv.Atoi(strPrice3)
	if err != nil {
		panic(err)
	}
	conf.price3 = price3

	strPrice6 := os.Getenv("PRICE_6")
	if strPrice6 == "" {
		panic("PRICE_6 .env variable not set")
	}
	price6, err := strconv.Atoi(strPrice6)
	if err != nil {
		panic(err)
	}
	conf.price6 = price6

	conf.remnawaveUrl = os.Getenv("REMNAWAVE_URL")
	if conf.remnawaveUrl == "" {
		panic("REMNAWAVE_URL .env variable not set")
	}

	conf.remnawaveMode = os.Getenv("REMNAWAVE_MODE")
	if conf.remnawaveMode == "" {
		conf.remnawaveMode = "remote"
	} else if conf.remnawaveMode != "remote" && conf.remnawaveMode != "local" {
		panic("REMNAWAVE_MODE .env variable must be either 'remote' or 'local'")
	}

	conf.remnawaveToken = os.Getenv("REMNAWAVE_TOKEN")
	if conf.remnawaveToken == "" {
		panic("REMNAWAVE_TOKEN .env variable not set")
	}

	conf.databaseURL = os.Getenv("DATABASE_URL")
	if conf.databaseURL == "" {
		panic("DADA_BASE_URL .env variable not set")
	}

	conf.isTelegramStarsEnabled = os.Getenv("TELEGRAM_STARS_ENABLED") == "true"

	conf.isCryptoEnabled = os.Getenv("CRYPTO_PAY_ENABLED") == "true"
	if conf.isCryptoEnabled {
		conf.cryptoPayURL = os.Getenv("CRYPTO_PAY_URL")
		if conf.cryptoPayURL == "" {
			panic("CRYPTO_PAY_URL .env variable not set")
		}
		conf.cryptoPayToken = os.Getenv("CRYPTO_PAY_TOKEN")
		if conf.cryptoPayToken == "" {
			panic("CRYPTO_PAY_TOKEN .env variable not set")
		}
	}

	conf.isYookasaEnabled = os.Getenv("YOOKASA_ENABLED") == "true"
	if conf.isYookasaEnabled {
		conf.yookasaURL = os.Getenv("YOOKASA_URL")
		conf.yookasaShopId = os.Getenv("YOOKASA_SHOP_ID")
		conf.yookasaSecretKey = os.Getenv("YOOKASA_SECRET_KEY")

		if conf.yookasaURL == "" || conf.yookasaShopId == "" || conf.yookasaSecretKey == "" {
			panic("YOOKASA_URL, YOOKASA_SHOP_ID, YOOKASA_SECRET_KEY .env variables not set")
		}

		conf.yookasaEmail = os.Getenv("YOOKASA_EMAIL")
		if conf.yookasaEmail == "" {
			panic("YOOKASA_EMAIL .env variable not set")
		}
	}

	strLimit := os.Getenv("TRAFFIC_LIMIT")
	if strLimit == "" {
		panic("TRAFFIC_LIMIT .env variable not set")
	}
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		panic(err)
	}
	conf.trafficLimit = int64(limit)

	conf.serverStatusURL = os.Getenv("SERVER_STATUS_URL")
	conf.supportURL = os.Getenv("SUPPORT_URL")
	conf.feedbackURL = os.Getenv("FEEDBACK_URL")
	conf.channelURL = os.Getenv("CHANNEL_URL")

}
