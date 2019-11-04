package util

//Using code of https://github.com/ashwanthkumar/slack-go-webhook

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"os"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	Style string `json:"style"`
}

type Attachment struct {
	Fallback     *string   `json:"fallback"`
	Color        *string   `json:"color"`
	PreText      *string   `json:"pretext"`
	AuthorName   *string   `json:"author_name"`
	AuthorLink   *string   `json:"author_link"`
	AuthorIcon   *string   `json:"author_icon"`
	Title        *string   `json:"title"`
	TitleLink    *string   `json:"title_link"`
	Text         *string   `json:"text"`
	ImageUrl     *string   `json:"image_url"`
	Fields       []*Field  `json:"fields"`
	Footer       *string   `json:"footer"`
	FooterIcon   *string   `json:"footer_icon"`
	Timestamp    *int64    `json:"ts"`
	MarkdownIn   *[]string `json:"mrkdwn_in"`
	Actions      []*Action `json:"actions"`
	CallbackID   *string   `json:"callback_id"`
	ThumbnailUrl *string   `json:"thumb_url"`
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func (attachment *Attachment) AddAction(action Action) *Attachment {
	attachment.Actions = append(attachment.Actions, &action)
	return attachment
}
func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	_, _ = req, via
	return fmt.Errorf("Incorrect token (redirection)")
}

func send(webhookUrl string, proxy string, payload Payload) []error {
	request := gorequest.New().Proxy(proxy)
	resp, _, err := request.
		Post(webhookUrl).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("Error sending slack msg. Status: %v", resp.Status)}
	}
	return nil
}

func SendLogToSlack(site string, backupType string, link string, status bool, extra string) {
	url := os.Getenv("Slack")
	_ = url
	attachment := Attachment{}
	var result string
	if status {
		result = "Success"
	} else {
		result = "Failed"
	}
	attachment.AddField(Field{Title: "Site", Value: site}).
		AddField(Field{Title: "Backup type", Value: backupType}).
		AddField(Field{Title: "Status", Value: result}).
		AddField(Field{Title: "Additional info", Value: extra})
	attachment.AddAction(Action{Type: "button", Text: "More info", Url: link})
	payload := Payload{
		Text:        "Failed on site backup:",
		Channel:     "backup_logs",
		Attachments: []Attachment{attachment},
	}
	err := send(url, "", payload)
	if len(err) != 0 {
		fmt.Print(err)
	}
	return
}
