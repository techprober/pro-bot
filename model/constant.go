package model

import "os"

const StartCommand string = "/start"
const BotTag string = "@TechProberBot"

const TelegramApiBaseUrl string = "https://api.telegram.org/bot"
const TelegramApiSendMessagePath string = "/sendMessage"
const TelegramApiSendStickerPath string = "/sendSticker"

const GoodMorningStickerRef = "CAACAgUAAxkBAAMoYe4q6wt_DVzPJjWKuuQRHm9zlxAAAtQDAAJnwDBWd7ZQ91obYOEjBA"

var TelegramApi string = TelegramApiBaseUrl + os.Getenv("TELEGRAM_TOKEN")
