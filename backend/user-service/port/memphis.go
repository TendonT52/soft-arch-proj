package port

import "github.com/memphisdev/memphis.go"

type MemphisPort interface {
	CreateProducer(stationName string, name string, opts ...memphis.ProducerOpt) (*memphis.Producer, error)
}
