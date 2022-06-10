package provider

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/constants/providers"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/common"
	"github.com/qdm12/gluetun/internal/provider/custom"
	"github.com/qdm12/gluetun/internal/provider/cyberghost"
	"github.com/qdm12/gluetun/internal/provider/expressvpn"
	"github.com/qdm12/gluetun/internal/provider/fastestvpn"
	"github.com/qdm12/gluetun/internal/provider/hidemyass"
	"github.com/qdm12/gluetun/internal/provider/ipvanish"
	"github.com/qdm12/gluetun/internal/provider/ivpn"
	"github.com/qdm12/gluetun/internal/provider/mullvad"
	"github.com/qdm12/gluetun/internal/provider/nordvpn"
	"github.com/qdm12/gluetun/internal/provider/perfectprivacy"
	"github.com/qdm12/gluetun/internal/provider/privado"
	"github.com/qdm12/gluetun/internal/provider/privateinternetaccess"
	"github.com/qdm12/gluetun/internal/provider/privatevpn"
	"github.com/qdm12/gluetun/internal/provider/protonvpn"
	"github.com/qdm12/gluetun/internal/provider/purevpn"
	"github.com/qdm12/gluetun/internal/provider/surfshark"
	"github.com/qdm12/gluetun/internal/provider/torguard"
	"github.com/qdm12/gluetun/internal/provider/vpnunlimited"
	"github.com/qdm12/gluetun/internal/provider/vyprvpn"
	"github.com/qdm12/gluetun/internal/provider/wevpn"
	"github.com/qdm12/gluetun/internal/provider/windscribe"
)

type Providers struct {
	providerNameToProvider map[string]Provider
}

type Storage interface {
	FilterServers(provider string, selection settings.ServerSelection) (
		servers []models.Server, err error)
	GetServerByName(provider, name string) (server models.Server, ok bool)
}

func NewProviders(storage Storage, timeNow func() time.Time,
	updaterWarner common.Warner, client *http.Client, unzipper common.Unzipper) *Providers {
	randSource := rand.NewSource(timeNow().UnixNano())

	targetLength := len(providers.AllWithCustom())
	providerNameToProvider := make(map[string]Provider, targetLength)
	providerNameToProvider[providers.Custom] = custom.New()
	providerNameToProvider[providers.Cyberghost] = cyberghost.New(storage, randSource)
	providerNameToProvider[providers.Expressvpn] = expressvpn.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Fastestvpn] = fastestvpn.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.HideMyAss] = hidemyass.New(storage, randSource, client, updaterWarner)
	providerNameToProvider[providers.Ipvanish] = ipvanish.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Ivpn] = ivpn.New(storage, randSource, client, updaterWarner)
	providerNameToProvider[providers.Mullvad] = mullvad.New(storage, randSource, client)
	providerNameToProvider[providers.Nordvpn] = nordvpn.New(storage, randSource, client, updaterWarner)
	providerNameToProvider[providers.Perfectprivacy] = perfectprivacy.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Privado] = privado.New(storage, randSource, client, unzipper, updaterWarner)
	providerNameToProvider[providers.PrivateInternetAccess] = privateinternetaccess.New(storage, randSource, timeNow, client) //nolint:lll
	providerNameToProvider[providers.Privatevpn] = privatevpn.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Protonvpn] = protonvpn.New(storage, randSource, client, updaterWarner)
	providerNameToProvider[providers.Purevpn] = purevpn.New(storage, randSource, client, unzipper, updaterWarner)
	providerNameToProvider[providers.Surfshark] = surfshark.New(storage, randSource, client, unzipper, updaterWarner)
	providerNameToProvider[providers.Torguard] = torguard.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.VPNUnlimited] = vpnunlimited.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Vyprvpn] = vyprvpn.New(storage, randSource, unzipper, updaterWarner)
	providerNameToProvider[providers.Wevpn] = wevpn.New(storage, randSource, updaterWarner)
	providerNameToProvider[providers.Windscribe] = windscribe.New(storage, randSource, client, updaterWarner)

	if len(providerNameToProvider) != targetLength {
		// Programming sanity check
		panic(fmt.Sprintf("invalid number of providers, expected %d but got %d",
			targetLength, len(providerNameToProvider)))
	}

	return &Providers{
		providerNameToProvider: providerNameToProvider,
	}
}

func (p *Providers) Get(providerName string) (provider Provider) {
	provider, ok := p.providerNameToProvider[providerName]
	if !ok {
		panic(fmt.Sprintf("provider %q not found", providerName))
	}
	return provider
}
