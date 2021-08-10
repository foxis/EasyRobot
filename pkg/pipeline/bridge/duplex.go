package bridge

import (
	"context"

	"github.com/foxis/EasyRobot/pkg/pipeline"
	"github.com/foxis/EasyRobot/pkg/plugin"
)

const BRIDGE_DUPLEX_NAME = "rw"

type duplex struct {
	options Options
	recv    pipeline.Step
	send    pipeline.Step
}

func init() {
	pipeline.Register(BRIDGE_DUPLEX_NAME, NewDuplexBridge)
}

func NewDuplexBridge(opts ...plugin.Option) (pipeline.Step, error) {
	recv, err := NewBridgeReceiver(opts...)
	if err != nil {
		return nil, err
	}
	send, err := NewBridgeSender(opts...)
	if err != nil {
		return nil, err
	}

	step := &duplex{
		options: Options{
			base: plugin.DefaultOptions(),
		},
		recv: recv,
		send: send,
	}
	plugin.ApplyOptions(&step.options, opts...)
	plugin.ApplyOptions(&step.options.base, opts...)

	step.Reset()
	return step, nil
}

func (s *duplex) In(ch <-chan pipeline.Data) {
	s.send.In(ch)
}

func (s *duplex) Out() <-chan pipeline.Data {
	return s.recv.Out()
}

func (s *duplex) Run(ctx context.Context) {
	go s.recv.Run(ctx)
	s.send.Run(ctx)
}

func (s *duplex) Reset() {
	s.send.Reset()
	s.recv.Reset()
}