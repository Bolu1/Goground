package main

import(
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3676004078884-3674004270018-F95OYMKQXyk8VEXabM9xpTsv")
	os.Setenv("CHANNEL_ID", "C03LHAWCMCG")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"text.ts"}

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err !=nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s has been uploaded to URL:%s\n", file.Name, file.URL)
	}
}