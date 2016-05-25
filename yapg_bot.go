package main

import (
	"crypto/rand"
	"log"
	"os"
	"strings"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	numbers     string = `23456789`
	uCaseLetter string = `ABCDEFGHJKLMNPRSTUVWXYZ`
	lCaseLetter string = `abcdefghikmnopqrstuvwxyz`
	dictionary  string = numbers + uCaseLetter + lCaseLetter
	dicLen      int    = len(dictionary)
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

	err = botman(bot, update)
	if err != nil {
		log.Fatal(err)
	}
}

func botman(b *tgbotapi.BotAPI, u tgbotapi.UpdateConfig) error {
	const defaultMessage = `Choose your destiny!`

	keyboard := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(`/easy`),
		tgbotapi.NewKeyboardButton(`/hard`))

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
		case "/start":
			{
				msg := tgbotapi.NewMessage(chatID, defaultMessage)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(keyboard)
				b.Send(msg)
			}
		case "/stop":
			{
				msg := tgbotapi.NewMessage(chatID, `SeeYa!`)
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{HideKeyboard: true}
				b.Send(msg)
			}
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
				msg := tgbotapi.NewMessage(chatID, defaultMessage)
				b.Send(msg)
			}
		}
	}
	return nil
}

func genPass(level byte) string {

	bytes := make([]byte, level)
	rand.Read(bytes)

	var result string
	for strings.ContainsAny(result, numbers) == false {
		for i, b := range bytes {
			bytes[i] = dictionary[b%byte(dicLen)]
		}
		result = string(bytes)
	}
	return result
}
