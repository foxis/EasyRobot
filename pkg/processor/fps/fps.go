package fps

import (
	"time"

	. "github.com/foxis/EasyRobot/pkg/logger"
	"github.com/foxis/EasyRobot/pkg/pipeline"
	"github.com/foxis/EasyRobot/pkg/pipeline/steps"
	"github.com/foxis/EasyRobot/pkg/plugin"
	"github.com/foxis/EasyRobot/pkg/store"
)

func init() {
	pipeline.Register(NAME, New)
}

func New(opts ...plugin.Option) (pipeline.Step, error) {
	o := Options{
		numFrames: 30,
	}
	plugin.ApplyOptions(&o, opts...)
	newOpts := append([]plugin.Option{plugin.WithName(NAME)}, opts...)
	newOpts = append(newOpts, WithFPSCounter(o.numFrames, o.suffix))
	return steps.NewProcessor(newOpts...)
}

func WithFPSCounter(numFrames int, suffix store.FQDNType) plugin.Option {
	fps := store.FPS
	drop := store.DROPPED_FRAMES
	if suffix != store.NONE {
		fps += suffix
		drop += suffix
	}
	return steps.WithProcessor(&fpsCounter{
		nameFPS:     fps,
		nameDropped: drop,
		numFrames:   numFrames,
	})
}

type fpsCounter struct {
	numFrames     int
	nameFPS       store.FQDNType
	nameDropped   store.FQDNType
	lastTimestamp int64
	lastIndex     int64
	dropped       int64
	fps           float32
}

func (s *fpsCounter) Process(src, dst store.Store) error {
	index, err := src.Index()
	if err != nil {
		Log.Error().Err(err).Msg(store.INDEX.String())
		return nil
	}

	timestamp, err := src.Timestamp()
	if err != nil {
		Log.Error().Err(err).Msg(store.TIMESTAMP.String())
		return nil
	}

	if index%int64(s.numFrames) == 0 {
		duration := time.Duration(timestamp - s.lastTimestamp)
		s.lastTimestamp = timestamp
		s.lastIndex = index
		if s.lastIndex >= 0 {
			s.fps = float32(s.numFrames) / float32(duration.Seconds())
		}
	}

	dst.SetFPS(float32(s.fps))
	dst.SetDropCount(int64(s.dropped))
	// dst.SetName(src.Name())

	return nil
}

func (s *fpsCounter) Init() error {
	s.lastTimestamp = time.Now().UnixNano()
	s.lastIndex = -1
	return nil
}

func (s *fpsCounter) Reset() {
	s.dropped = 0
}

func (s *fpsCounter) Close() {
	s.dropped = 0
}