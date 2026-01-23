package main

import (
	"errors"
	"os"

	"github.com/agentio/chatter/gen/xrpc"
	client "github.com/agentio/slink/pkg/client"

	"github.com/agentio/slink/pkg/common"
	"github.com/agentio/slink/pkg/resolve"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func main() {
	if err := cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func cmd() *cobra.Command {
	var loglevel string
	var text string
	var from string
	var to string
	var password string

	cmd := &cobra.Command{
		Use:   "chatter",
		Short: "Sends messages over Bluesky chat",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.SetLogLevel(loglevel); err != nil {
				return err
			}
			if text == "" || from == "" || to == "" || password == "" {
				return errors.New("all of --text, --from, --to, and --password must be set")
			}
			fromDid, err := resolve.Handle(cmd.Context(), from)
			if err != nil {
				return err
			}
			fromDidDocument, err := resolve.Did(cmd.Context(), fromDid)
			if err != nil {
				return err
			}
			c := client.NewClientWithOptions(client.ClientOptions{
				Host: fromDidDocument.Service[0].ServiceEndpoint,
			})
			result, err := xrpc.ServerCreateSession(cmd.Context(), c, &xrpc.ServerCreateSession_Input{
				Identifier: fromDid,
				Password:   password,
			})
			if err != nil {
				return err
			}
			log.Infof("%s", result.AccessJwt)
			toDid, err := resolve.Handle(cmd.Context(), to)
			if err != nil {
				return err
			}
			log.Infof("%s %s %s", fromDid, toDid, fromDidDocument.Service[0].ServiceEndpoint)
			c = client.NewClientWithOptions(client.ClientOptions{
				Host:          fromDidDocument.Service[0].ServiceEndpoint,
				Authorization: "Bearer " + result.AccessJwt,
				ATProtoProxy:  "did:web:api.bsky.chat#bsky_chat",
			})
			result2, err := xrpc.ChatBskyConvoGetConvoForMembers(cmd.Context(), c, []string{fromDid, toDid})
			if err != nil {
				return err
			}
			convoId := result2.Convo.Id
			log.Infof("convo id %s", convoId)
			result3, err := xrpc.ChatBskyConvoSendMessage(cmd.Context(), c, &xrpc.ChatBskyConvoSendMessage_Input{
				ConvoId: convoId,
				Message: &xrpc.ChatBskyConvoDefs_MessageInput{
					Text: text,
				},
			})
			log.Infof("%+v", result3)
			return nil
		},
	}
	cmd.Flags().StringVar(&text, "text", "", "message text (required)")
	cmd.Flags().StringVar(&from, "from", "", "message sender (required)")
	cmd.Flags().StringVar(&to, "to", "", "messsage recipient (required)")
	cmd.Flags().StringVar(&password, "password", "", "the sender's password (required)")
	cmd.Flags().StringVarP(&loglevel, "log-level", "l", "warn", "log level (debug, info, warn, error, fatal)")
	return cmd
}
