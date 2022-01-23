package main

import "os"

const PunchCommand string = "/punch"
const StartCommand string = "/start"
const BotTag string = "@TechProberBot"

const TelegramApiBaseUrl string = "https://api.telegram.org/bot"
const TelegramApiSendMessage string = "/sendMessage"
const TelegramToken string = "TELEGRAM_BOT_TOKEN"

var TelegramApi string = TelegramApiBaseUrl + os.Getenv("TELEGRAM_TOKEN") + TelegramApiSendMessage
