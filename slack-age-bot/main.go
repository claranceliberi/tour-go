package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-549525876311-3424042919110-Avb5tUC0kIvFNNOP7wqaiM1g")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03C82BU4LF-3430695231234-177b8227bbe7beec8dea38ad6f5050560668768cde630d1bb6e4f665e30c3234")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2000",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")

			yob, err := strconv.Atoi(year)

			if err != nil {
				log.Println("Error occured..........")
				log.Println(err)
			}

			age := 2022 - yob

			r := fmt.Sprintf("age is %d", age)

			log.Println("---------------------------------------", r)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
