package main

import (
	"catfactsbot/internal/app/botservice"
	"catfactsbot/internal/utils"
	"catfactsbot/providers/catfacts"
	"catfactsbot/providers/telegram"
)

func main() {

	cfProv := catfacts.New()
	tgProv := telegram.New("https://api.telegram.org/", utils.ReadEnvVar("BOT_API"), 0)
	bot := botservice.New()

	bot.Serve(cfProv, tgProv)
}
