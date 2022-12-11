package bot
import(
	"fmt"
	"github.com/MattLumm/GoLangDiscordPingBot/config" //NOTE: Token Removed for protection on the files uploaded to GitHub
	"github.com/bwmarrin/discordgo" //Go pakcage that provides low lvl bindings to the Discord chat client API
)

var BotID string
var goBot *discordgo.Session

func Start(){
	goBot, err := discordgo.New("Bot " + config.Token) //must have space after Bot

	if err !=nil {
		fmt.Println(err.Error())
		return
	}

	u, err:= goBot.User("@me")

	if err !=nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID
	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err !=nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){ 

	if m.Author.ID == BotID{ //varifies that message is not from the Bot so it doesn't infitely loop on it's own msgs
		return
	}

	if m.Content == "ping"{
		_,_ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}