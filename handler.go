package valkey

// func (v *Valkey) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
// 	state := request.Request{W: w, Req: r}

// 	qname := state.Name()
// 	qtype := state.Type()

// 	if time.Since(v.LastZoneUpdate) > zoneUpdateTime {
// 		v.LoadZones()
// 	}

// 	zone := plugin.Zones(v.Zones).Matches(qname)
// }

func (v *Valkey) Name() string { return "valkey" }
