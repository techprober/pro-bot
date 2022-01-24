package main

import "os"

const StartCommand string = "/start"
const BotTag string = "@TechProberBot"

const TelegramApiBaseUrl string = "https://api.telegram.org/bot"
const TelegramApiSendMessagePath string = "/sendMessage"

var TelegramApi string = TelegramApiBaseUrl + os.Getenv("TELEGRAM_TOKEN") + TelegramApiSendMessagePath
