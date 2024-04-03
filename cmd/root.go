package cmd

import (
	"fmt"
	"strings"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/spf13/cobra"
)

type SlackNotification struct {
	WebhookURL    string //  Common
	Channel       string // Common
	LineMessage   string // Common
	GitRepository string
	GitBranch     string
	GitCommidId   string
	GitCommitter  string
	GitMessage    string
	Env           string
	Service       string

	// url
	UrlToCodePipeline string
	UrlToGitCommit    string
}

var (
	sn      SlackNotification
	rootCmd = &cobra.Command{
		Use:   "Slack-Noti-Job",
		Short: "Slack-Noti-Job",
		Long:  "Slack-Noti-Job",
		Run: func(cmd *cobra.Command, args []string) {

			// Not Exist Slack Webhook URL
			if strings.Trim(sn.WebhookURL, "") == "" {
				return
			}

			// Connect WebHook URL

			if !strings.HasPrefix(sn.Channel, "#") {
				sn.Channel = fmt.Sprintf("#%s", sn.Channel)
			}

			payload := slack.Payload{
				Username:    "leedonggyu",
				Channel:     sn.Channel,
				IconEmoji:   ":fire:",
				Attachments: []slack.Attachment{getSlackAttachment()},
			}

			err := slack.Send(sn.WebhookURL, "", payload)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func getSlackAttachment() slack.Attachment {
	attach := slack.Attachment{
		Title: &sn.LineMessage,
	}

	attach.AddField(slack.Field{Title: "Repository", Value: sn.GitRepository, Short: true})
	attach.AddField(slack.Field{Title: "Message", Value: sn.GitMessage, Short: true})

	attach.AddField(slack.Field{Title: "Branch", Value: sn.GitBranch, Short: true})
	attach.AddField(slack.Field{Title: "CommitId", Value: sn.GitCommidId, Short: true})
	attach.AddField(slack.Field{Title: "Committer", Value: sn.GitCommitter, Short: true})

	attach.AddField(slack.Field{Title: "Env", Value: sn.Env, Short: true})
	attach.AddField(slack.Field{Title: "Services", Value: sn.Service, Short: true})

	attach.AddAction(slack.Action{
		Text:  "Go To CodePipeline",
		Style: "primary",
		Type:  "button",
		Url:   sn.UrlToCodePipeline,
	})

	attach.AddAction(slack.Action{
		Text:  "Go To GithubCommit",
		Style: "primary",
		Type:  "button",
		Url:   sn.UrlToGitCommit,
	})

	return attach
}

func initial() {

	// intitial
	rootCmd.PersistentFlags().StringVarP(&sn.WebhookURL, "webhook-url", "w", "", "[Required] Slack Web Hook URL")
	rootCmd.PersistentFlags().StringVarP(&sn.Channel, "channel", "c", "", "[Required] Channel")
	rootCmd.PersistentFlags().StringVarP(&sn.LineMessage, "lineMessage", "l", "", "[Required] Line Message")

	rootCmd.PersistentFlags().StringVarP(&sn.GitRepository, "git-repository", "r", "", "[Reqruied] Git Repository")
	rootCmd.PersistentFlags().StringVarP(&sn.GitBranch, "git-branch", "b", "", "[Required] Git branch")
	rootCmd.PersistentFlags().StringVarP(&sn.GitCommidId, "git-commitId ", "i", "", "[Required] Git CommitId")
	rootCmd.PersistentFlags().StringVarP(&sn.GitCommitter, "git-committer", "t", "", "[Required] Git Committer")
	rootCmd.PersistentFlags().StringVarP(&sn.GitMessage, "git-message", "m", "", "[Required] Git Message")
	rootCmd.PersistentFlags().StringVarP(&sn.Env, "env", "e", "", "[Optional] Environments")
	rootCmd.PersistentFlags().StringVarP(&sn.Service, "service", "s", "", "[Optional] Services")

	rootCmd.PersistentFlags().StringVarP(&sn.UrlToCodePipeline, "codepipelineUrl", "p", "", "[Optional] CodePipeline URL")
	rootCmd.PersistentFlags().StringVarP(&sn.UrlToGitCommit, "githubCommitUrl", "o", "", "[Optional] CommitID")
}

func Execute() {
	initial()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
