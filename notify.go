package notify

import (
	"errors"
	"github.com/AlphaAILabs/go-notify/fwalert"
	"strconv"
	"strings"

	"github.com/AlphaAILabs/go-notify/dingtalk"
	"github.com/AlphaAILabs/go-notify/discord"
	"github.com/AlphaAILabs/go-notify/email"
	"github.com/AlphaAILabs/go-notify/lark"
	"github.com/AlphaAILabs/go-notify/pagerduty"
	"github.com/AlphaAILabs/go-notify/pushover"
	"github.com/AlphaAILabs/go-notify/slack"
	"github.com/AlphaAILabs/go-notify/telegram"
)

type Platform string

const (
	PlatformSlack     Platform = "Slack"
	PlatformPushover           = "Pushover"
	PlatformPagerduty          = "Pagerduty"
	PlatformDiscord            = "Discord"
	PlatformTelegram           = "Telegram"
	PlatformDingTalk           = "DingTalk"
	PlatformEmail              = "Email"
	PlatformLark               = "Lark"
	PlatformFwAlert            = "FwAlert"
)

type Notify struct {
	config *Config
}

type Config struct {
	Platform Platform `mapstructure:"platform"`
	Token    string   `mapstructure:"token"`
	Channel  string   `mapstructure:"channel"`
	Source   string   `mapstructure:"source"`
	Severity string   `mapstructure:"severity"`
	User     string   `mapstructure:"user"`
	Password string   `mapstructure:"password"`
	Host     string   `mapstructure:"host"`
}

func NewNotify(config *Config) *Notify {
	return &Notify{
		config: config,
	}
}

func (n *Notify) Send(msg string) error {
	switch n.config.Platform {
	case PlatformPushover:
		return n.sendPushOverNotify(msg)
	case PlatformSlack:
		return n.sendSlackNotify(msg)
	case PlatformPagerduty:
		return n.sendPagerdutyNotify(msg)
	case PlatformDiscord:
		return n.sendDiscordNotify(msg)
	case PlatformTelegram:
		return n.sendTelegramNotify(msg)
	case PlatformDingTalk:
		return n.sendDingTalkNotify(msg)
	case PlatformEmail:
		return n.sendEmailNotify(msg)
	case PlatformLark:
		return n.sendLarkNotify(msg)
	case PlatformFwAlert:
		return n.sendFwAlertNotify(msg)
	default:
		return errors.New("not supported notify platform")
	}
}

func (n *Notify) sendPushOverNotify(msg string) error {
	app := pushover.New(pushover.Options{
		Token: n.config.Token,
		User:  n.config.Channel,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendSlackNotify(msg string) error {
	app := slack.New(slack.Options{
		Token:   n.config.Token,
		Channel: n.config.Channel,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendPagerdutyNotify(msg string) error {
	app := pagerduty.New(pagerduty.Options{
		Token:    n.config.Token,
		Source:   n.config.Source,
		Severity: n.config.Severity,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendDiscordNotify(msg string) error {
	app := discord.New(discord.Options{
		Token:   n.config.Token,
		Channel: n.config.Channel,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendTelegramNotify(msg string) error {
	var _channel int64
	var _chatName string
	if strings.Contains(n.config.Channel, "@") {
		_channel = 0
		_chatName = n.config.Channel
	} else {
		_channel, _ = strconv.ParseInt(n.config.Channel, 10, 64)
		_chatName = ""
	}

	app := telegram.New(telegram.Options{
		Token:    n.config.Token,
		Channel:  _channel,
		ChatName: _chatName,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendDingTalkNotify(msg string) error {
	app := dingtalk.New(dingtalk.Options{
		WebhookUrl: n.config.Channel,
		Secret:     n.config.Token,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendEmailNotify(msg string) error {
	app := email.New(email.Options{
		ToEmail:  n.config.Token,
		User:     n.config.User,
		Password: n.config.Password,
		Host:     n.config.Host,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendLarkNotify(msg string) error {
	app := lark.New(lark.Options{
		Token: n.config.Token,
	})
	err := app.Send(msg)
	return err
}

func (n *Notify) sendFwAlertNotify(msg string) error {
	app := fwalert.New(fwalert.Options{
		Token: n.config.Token,
	})
	err := app.Send(msg)
	return err
}
