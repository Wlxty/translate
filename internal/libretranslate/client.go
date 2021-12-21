package libretranslate

import "go.uber.org/zap"

type Client struct {
	Logger *zap.SugaredLogger
	Host   string
}

func NewClient(Logger *zap.SugaredLogger, Host string) *Client {
	client := Client{Logger, Host}
	return &client
}
