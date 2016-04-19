package main

import (
	"crypto/rand"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

const (
	numbers    string = `23456789`
	characters string = `ABCDEFGHJKLMNPRSTUVWXYZabcdefghikmnopqrstuvwxyz`
	dictionary string = numbers + characters
	dicLen     int    = len(dictionary)
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOTTOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	for err == nil {
		err = botman(bot, update)
	}
	log.Fatal(err)
}

func botman(b *tgbotapi.BotAPI, u tgbotapi.UpdateConfig) error {

	updates, err := b.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {

		chatID := update.Message.Chat.ID
		text := update.Message.Text
		if b.Debug == true {
			log.Printf("[%s] %s", update.Message.From.UserName, text)
		}

		switch text {
		case "/easy":
			{
				msg := tgbotapi.NewMessage(chatID, genPass(10))
				b.Send(msg)
			}
		case "/hard":
			{
				msg := tgbotapi.NewMessage(chatID, genPass(20))
				b.Send(msg)
			}
		default:
			{
				msg := tgbotapi.NewMessage(chatID, "WAT")
				b.Send(msg)
			}
		}
	}
	return nil
}

func genPass(level byte) string {

	bytes := make([]byte, level)
	rand.Read(bytes)

	result := ""
	for strings.ContainsAny(result, numbers) == false {
		for i, b := range bytes {
			bytes[i] = dictionary[b%byte(dicLen)]
		}
		result = string(bytes)
	}
	return result
}
