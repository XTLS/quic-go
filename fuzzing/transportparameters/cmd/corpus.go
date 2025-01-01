package main

import (
	"log"
	"math"
	"net/netip"
	"time"

	"golang.org/x/exp/rand"

	"github.com/xtls/quic-go/fuzzing/internal/helper"
	"github.com/xtls/quic-go/fuzzing/transportparameters"
	"github.com/xtls/quic-go/internal/protocol"
	"github.com/xtls/quic-go/internal/wire"
)

func getRandomData(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

func getRandomValue() uint64 {
	maxVals := []int64{math.MaxUint8 / 4, math.MaxUint16 / 4, math.MaxUint32 / 4, math.MaxUint64 / 4}
	return uint64(rand.Int63n(maxVals[int(rand.Int31n(4))]))
}

func main() {
	for i := 0; i < 30; i++ {
		tp := &wire.TransportParameters{
			InitialMaxStreamDataBidiLocal:  protocol.ByteCount(getRandomValue()),
			InitialMaxStreamDataBidiRemote: protocol.ByteCount(getRandomValue()),
			InitialMaxStreamDataUni:        protocol.ByteCount(getRandomValue()),
			InitialMaxData:                 protocol.ByteCount(getRandomValue()),
			MaxAckDelay:                    time.Duration(getRandomValue()),
			AckDelayExponent:               uint8(getRandomValue()),
			DisableActiveMigration:         getRandomValue()%2 == 0,
			MaxUDPPayloadSize:              protocol.ByteCount(getRandomValue()),
			MaxUniStreamNum:                protocol.StreamNum(getRandomValue()),
			MaxBidiStreamNum:               protocol.StreamNum(getRandomValue()),
			MaxIdleTimeout:                 time.Duration(getRandomValue()),
			ActiveConnectionIDLimit:        getRandomValue() + 2,
		}
		if rand.Int()%2 == 0 {
			tp.OriginalDestinationConnectionID = protocol.ParseConnectionID(getRandomData(rand.Intn(21)))
		}
		if rand.Int()%2 == 0 {
			tp.InitialSourceConnectionID = protocol.ParseConnectionID(getRandomData(rand.Intn(21)))
		}
		if rand.Int()%2 == 0 {
			connID := protocol.ParseConnectionID(getRandomData(rand.Intn(21)))
			tp.RetrySourceConnectionID = &connID
		}
		if rand.Int()%2 == 0 {
			var token protocol.StatelessResetToken
			rand.Read(token[:])
			tp.StatelessResetToken = &token
		}
		if rand.Int()%2 == 0 {
			var token protocol.StatelessResetToken
			rand.Read(token[:])
			var ip4 [4]byte
			rand.Read(ip4[:])
			var ip6 [16]byte
			rand.Read(ip6[:])
			tp.PreferredAddress = &wire.PreferredAddress{
				IPv4:                netip.AddrPortFrom(netip.AddrFrom4(ip4), uint16(rand.Int())),
				IPv6:                netip.AddrPortFrom(netip.AddrFrom16(ip6), uint16(rand.Int())),
				ConnectionID:        protocol.ParseConnectionID(getRandomData(rand.Intn(21))),
				StatelessResetToken: token,
			}
		}

		var data []byte
		if rand.Int()%2 == 0 {
			pers := protocol.PerspectiveServer
			if rand.Int()%2 == 0 {
				pers = protocol.PerspectiveClient
			}
			data = tp.Marshal(pers)
		} else {
			data = tp.MarshalForSessionTicket(nil)
		}
		if err := helper.WriteCorpusFileWithPrefix("corpus", data, transportparameters.PrefixLen); err != nil {
			log.Fatal(err)
		}
	}
}
