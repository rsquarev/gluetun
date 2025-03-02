package server

import (
	"context"
	"fmt"

	"github.com/qdm12/gluetun/internal/httpserver"
	"github.com/qdm12/gluetun/internal/models"
)

func New(ctx context.Context, address string, logEnabled bool, logger Logger,
	buildInfo models.BuildInformation, openvpnLooper VPNLooper,
	pfGetter PortForwardedGetter, unboundLooper DNSLoop,
	updaterLooper UpdaterLooper, publicIPLooper PublicIPLoop, storage Storage) (
	server *httpserver.Server, err error) {
	handler := newHandler(ctx, logger, logEnabled, buildInfo,
		openvpnLooper, pfGetter, unboundLooper, updaterLooper, publicIPLooper, storage)

	httpServerSettings := httpserver.Settings{
		Address: address,
		Handler: handler,
		Logger:  logger,
	}

	server, err = httpserver.New(httpServerSettings)
	if err != nil {
		return nil, fmt.Errorf("cannot create server: %w", err)
	}

	return server, nil
}
