package main

import(
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-3676004078884-3674004270018-F95OYMKQXyk8VEXabM9xpTsv")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A03KTG6V7QT-3686283486785-13ca99c049dcc69c0a99e4d433b0757b8cec4f0432cc2a0285ce4995550f7af5")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example: "my yob is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				println(err)
			}
			age := 2022-yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}

}