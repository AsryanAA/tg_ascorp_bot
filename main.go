package main

import (
	"flag"
	"log"

	"ascorp/clients/tg"
	"ascorp/consts"
)

func main() {
	tgClient := tg.New("api.telegram.org", mustToken())

	lastUpdateID, limit := 0, 1
	for {
		updates, _ := tgClient.Updates(lastUpdateID, limit)

		if len(updates) == 0 {
			continue
		}

		for _, update := range updates {
			if update.Message.IsCommand() {
				switch update.Message.Text {
				case consts.Start:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgWelcome)
				case consts.Subscribe:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgOK)
				case consts.AboutMe:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgAboutMe)
				case consts.CanNot:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgCanNot)
				case consts.CMDs:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgCMDs)
				default:
					tgClient.SendMessage(update.Message.Chat.ID, consts.MsgUnknownCommand)
				}
			}
		}

		lastUpdateID = updates[len(updates)-1].ID + 1
	}
}

func mustToken() string {
	token := flag.String("token", "", "token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}

	return *token
}
