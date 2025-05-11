package valkey

import (
	"time"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/valkey-io/valkey-go"
)

var log = clog.NewWithPlugin("valkey")

type Valkey struct {
	Next           plugin.Handler
	valkeyClient   *valkey.Client
	valkeyAddress  string
	valkeyPassword string
	connectTimeout int
	readTimeout    int
	keyPrefix      string
	keySuffix      string
	Ttl            uint32
	Zones          []string
	LastZoneUpdate time.Time
}

func (v *Valkey) Connect() {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{v.valkeyAddress}, Password: v.valkeyPassword})
	if err != nil {
		log.Error("Could not create Valkey client")
	}
	v.valkeyClient = &client
}

const (
	defaultTtl     = 300
	hostmaster     = "hostmaster"
	zoneUpdateTime = 10 * time.Minute
	transferLength = 1000
)
