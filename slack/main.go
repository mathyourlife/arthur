package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func reactions(api *slack.Client) {
	var (
		postAsUserName  string
		postAsUserID    string
		postToUserName  string
		postToUserID    string
		postToChannelID string
	)

	// Find the user to post as.
	authTest, err := api.AuthTest()
	if err != nil {
		fmt.Printf("Error getting channels: %s\n", err)
		return
	}

	// Post as the authenticated user.
	postAsUserName = authTest.User
	postAsUserID = authTest.UserID

	// Posting to DM with self causes a conversation with slackbot.
	postToUserName = authTest.User
	postToUserID = authTest.UserID

	// Find the channel.
	_, _, chanID, err := api.OpenIMChannel(postToUserID)
	if err != nil {
		fmt.Printf("Error opening IM: %s\n", err)
		return
	}
	postToChannelID = chanID

	fmt.Printf("Posting as %s (%s) in DM with %s (%s), channel %s\n", postAsUserName, postAsUserID, postToUserName, postToUserID, postToChannelID)

	// Post a message.
	channelID, timestamp, err := api.PostMessage(postToChannelID, slack.MsgOptionText("Is this any good?", false))
	if err != nil {
		fmt.Printf("Error posting message: %s\n", err)
		return
	}

	// Grab a reference to the message.
	msgRef := slack.NewRefToMessage(channelID, timestamp)

	// React with :+1:
	if err = api.AddReaction("+1", msgRef); err != nil {
		fmt.Printf("Error adding reaction: %s\n", err)
		return
	}

	// React with :-1:
	if err = api.AddReaction("cry", msgRef); err != nil {
		fmt.Printf("Error adding reaction: %s\n", err)
		return
	}

	// Get all reactions on the message.
	msgReactions, err := api.GetReactions(msgRef, slack.NewGetReactionsParameters())
	if err != nil {
		fmt.Printf("Error getting reactions: %s\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("%d reactions to message...\n", len(msgReactions))
	for _, r := range msgReactions {
		fmt.Printf("  %d users say %s\n", r.Count, r.Name)
	}

	// List all of the users reactions.
	listReactions, _, err := api.ListReactions(slack.NewListReactionsParameters())
	if err != nil {
		fmt.Printf("Error listing reactions: %s\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("All reactions by %s...\n", authTest.User)
	for _, item := range listReactions {
		fmt.Printf("%d on a %s...\n", len(item.Reactions), item.Type)
		for _, r := range item.Reactions {
			fmt.Printf("  %s (along with %d others)\n", r.Name, r.Count-1)
		}
	}

	// Remove the :cry: reaction.
	err = api.RemoveReaction("cry", msgRef)
	if err != nil {
		fmt.Printf("Error remove reaction: %s\n", err)
		return
	}

	// Get all reactions on the message.
	msgReactions, err = api.GetReactions(msgRef, slack.NewGetReactionsParameters())
	if err != nil {
		fmt.Printf("Error getting reactions: %s\n", err)
		return
	}
	fmt.Printf("\n")
	fmt.Printf("%d reactions to message after removing cry...\n", len(msgReactions))
	for _, r := range msgReactions {
		fmt.Printf("  %d users say %s\n", r.Count, r.Name)
	}
}

func writeMessage(api *slack.Client) {
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		// Uncomment the following part to send a field too
		// Fields: []slack.AttachmentField{
		// 	slack.AttachmentField{
		// 		Title: "a",
		// 		Value: "no",
		// 	},
		// },
	}

	channelID, timestamp, err := api.PostMessage(
		"C019QSTFU02",
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}

func rtm(api *slack.Client) {
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		log.Println(msg)
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "C2147483705"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		case *slack.ConnectionErrorEvent:
			log.Printf("&slack.ConnectionErrorEvent %#v", ev.ErrorObj)

		default:
			log.Printf("default %#v", ev)
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func users() {
	user, err := api.GetUserInfo("U023BECGF")
if err != nil {
	fmt.Printf("%s\n", err)
	return
}
fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))

	reactions(api)
}
