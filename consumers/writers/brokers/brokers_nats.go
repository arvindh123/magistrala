// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

//go:build !rabbitmq
// +build !rabbitmq

package brokers

import (
	"context"
	"log/slog"
	"time"

	"github.com/absmach/supermq/pkg/messaging"
	broker "github.com/absmach/supermq/pkg/messaging/nats"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	AllTopic = "writers.>"

	prefix = "writers"
)

var cfg = jetstream.StreamConfig{
	Name:              "writers",
	Description:       "SuperMQ Rules Engine stream for handling internal messages",
	Subjects:          []string{"writers.>"},
	Retention:         jetstream.LimitsPolicy,
	MaxMsgsPerSubject: 1e6,
	MaxAge:            time.Hour * 24,
	MaxMsgSize:        1024 * 1024,
	Discard:           jetstream.DiscardOld,
	Storage:           jetstream.FileStorage,
}

func NewPubSub(ctx context.Context, url string, logger *slog.Logger) (messaging.PubSub, error) {
	pb, err := broker.NewPubSub(ctx, url, logger, broker.Prefix(prefix), broker.JSStreamConfig(cfg))
	if err != nil {
		return nil, err
	}

	return pb, nil
}

func NewPublisher(ctx context.Context, url string) (messaging.Publisher, error) {
	pb, err := broker.NewPublisher(ctx, url, broker.Prefix(prefix), broker.JSStreamConfig(cfg))
	if err != nil {
		return nil, err
	}

	return pb, nil
}
