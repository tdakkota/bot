package main

import (
	"net/http"
	"os"
	"strconv"

	"golang.org/x/xerrors"

	"github.com/gotd/bot/internal/gentext"
)

func setupGPT2(client *http.Client) (gentext.Net, error) {
	g := gentext.NewGPT2().WithClient(client)
	if v := os.Getenv("GPT_MIN_SIZE"); v != "" {
		min, err := strconv.Atoi(v)
		if err != nil {
			return nil, xerrors.Errorf("GPT_MIN_SIZE %q is invalid: %w", v, err)
		}
		g = g.WithMinLength(min)
	}
	if v := os.Getenv("GPT_MAX_SIZE"); v != "" {
		max, err := strconv.Atoi(v)
		if err != nil {
			return nil, xerrors.Errorf("GPT_MAX_SIZE %q is invalid: %w", v, err)
		}
		g = g.WithMaxLength(max)
	}

	return g, nil
}

func setupGPT3(client *http.Client) (gentext.Net, error) {
	g := gentext.NewGPT3().WithClient(client)
	return g, nil
}
