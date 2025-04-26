package valkey

import (
	"time"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var log = clog.NewWithPlugin("valkey")

type Valkey struct {
	Next           plugin.Handler
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

func (valkey *Valkey) LoadZones() {
	var (
		replay interface{}
		err    error
		zones  []string
	)

	// conn := valkey
}
