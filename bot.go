package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

type Bot struct {
	state  *State
	client *telegram.Client
	rpc    *tg.Client
	logger *zap.Logger

	m Metrics
}

func NewBot(state *State, client *telegram.Client) *Bot {
	return &Bot{
		state:  state,
		client: client,
		rpc:    tg.NewClient(client),
		logger: zap.NewNop(),
	}
}

// WithLogger sets logger.
func (b *Bot) WithLogger(logger *zap.Logger) *Bot {
	b.logger = logger
	return b
}

func (b *Bot) WithStart(t time.Time) *Bot {
	b.m.Start = t
	return b
}

// sendMessage calls SendMessage and handles IsUserBlocked error.
func (b *Bot) sendMessage(ctx tg.UpdateContext, m *tg.MessagesSendMessageRequest) error {
	if err := b.client.SendMessage(ctx, m); err != nil {
		if tg.IsUserBlocked(err) {
			b.logger.Debug("Bot is blocked by user")
			return nil
		}
		return xerrors.Errorf("send message: %w", err)
	}

	// Increasing total response count metric.
	b.m.Responses.Inc()

	return nil
}

// sendMedia calls SendMedia and handles IsUserBlocked error.
func (b *Bot) sendMedia(ctx tg.UpdateContext, m *tg.MessagesSendMediaRequest) error {
	if m.RandomID == 0 {
		id, err := b.client.RandInt64()
		if err != nil {
			return xerrors.Errorf("gen id: %w", err)
		}
		m.RandomID = id
	}

	_, err := b.rpc.MessagesSendMedia(ctx, m)
	if err != nil {
		return xerrors.Errorf("send: %w", err)
	}
	// Increasing total response count metric.
	b.m.Responses.Inc()

	return nil
}

func (b *Bot) handleUser(ctx tg.UpdateContext, user *tg.User, m *tg.Message) error {
	b.logger.Info("Got message",
		zap.String("text", m.Message),
		zap.Int("user_id", user.ID),
		zap.String("user_first_name", user.FirstName),
		zap.String("username", user.Username),
	)

	reply := &tg.MessagesSendMessageRequest{
		Message:      fmt.Sprintf("No u %s, @%s", m.Message, user.Username),
		ReplyToMsgID: m.ID,
		Peer: &tg.InputPeerUser{
			UserID:     user.ID,
			AccessHash: user.AccessHash,
		},
	}

	if err := b.sendMessage(ctx, reply); err != nil {
		return xerrors.Errorf("send: %w", err)
	}

	return nil
}

func (b *Bot) handleChat(ctx tg.UpdateContext, peer *tg.Chat, m *tg.Message) error {
	b.logger.Info("Got message from chat",
		zap.String("text", m.Message),
		zap.Int("channel_id", peer.ID),
	)

	return b.answer(ctx, m, &tg.InputPeerChat{
		ChatID: peer.ID,
	}, m.ID)
}

func (b *Bot) handleChannel(ctx tg.UpdateContext, peer *tg.Channel, m *tg.Message) error {
	b.logger.Info("Got message from channel",
		zap.String("text", m.Message),
		zap.Int("channel_id", peer.ID),
	)

	return b.answer(ctx, m, &tg.InputPeerChannel{
		ChannelID:  peer.ID,
		AccessHash: peer.AccessHash,
	}, m.ID)
}

func (b *Bot) Handle(ctx tg.UpdateContext, msg tg.MessageClass) error {
	switch m := msg.(type) {
	case *tg.Message:
		if m.Out {
			return nil
		}

		b.m.Messages.Inc()

		switch peer := m.PeerID.(type) {
		case *tg.PeerUser:
			user, ok := ctx.Users[peer.UserID]
			if !ok {
				return xerrors.Errorf("unknown user ID %d", peer.UserID)
			}
			return b.handleUser(ctx, user, m)
		case *tg.PeerChat:
			chat, ok := ctx.Chats[peer.ChatID]
			if !ok {
				return xerrors.Errorf("unknown chat ID %d", peer.ChatID)
			}
			return b.handleChat(ctx, chat, m)
		case *tg.PeerChannel:
			channel, ok := ctx.Channels[peer.ChannelID]
			if !ok {
				return xerrors.Errorf("unknown channel ID %d", peer.ChannelID)
			}
			return b.handleChannel(ctx, channel, m)
		}
	}

	return nil
}

func (b *Bot) OnNewMessage(ctx tg.UpdateContext, u *tg.UpdateNewMessage) error {
	if err := b.Handle(ctx, u.Message); err != nil {
		return xerrors.Errorf("handle: %w", err)
	}

	if err := b.state.Commit(u.Pts); err != nil {
		return xerrors.Errorf("commit: %w", err)
	}
	return nil
}

func (b *Bot) OnNewChannelMessage(ctx tg.UpdateContext, u *tg.UpdateNewChannelMessage) error {
	if err := b.Handle(ctx, u.Message); err != nil {
		return xerrors.Errorf("handle: %w", err)
	}

	if err := b.state.Commit(u.Pts); err != nil {
		return xerrors.Errorf("commit: %w", err)
	}
	return nil
}
