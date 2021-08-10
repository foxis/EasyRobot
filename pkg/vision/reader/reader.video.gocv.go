package reader

import (
	"time"

	"github.com/foxis/EasyRobot/pkg/backend"
	. "github.com/foxis/EasyRobot/pkg/logger"
	"github.com/foxis/EasyRobot/pkg/pipeline"
	"github.com/foxis/EasyRobot/pkg/pipeline/steps"
	"github.com/foxis/EasyRobot/pkg/plugin"

	"gocv.io/x/gocv"
)

type readerVideoGoCV struct {
	readerOpts
	dev *gocv.VideoCapture
}

const VIDEO_GOCV_NAME = "rvid"

func init() {
	pipeline.Register(VIDEO_GOCV_NAME, NewVideoFileGoCV)
}

func NewVideoFileGoCV(opts ...plugin.Option) (pipeline.Step, error) {
	o := readerOpts{}
	plugin.ApplyOptions(&o, opts...)
	newOpts := opts
	newOpts = append(newOpts, WithVideoReaderGoCV(o.fname))
	return steps.NewReader(newOpts...)
}

func WithVideoReaderGoCV(file string) plugin.Option {
	return steps.WithSourceReader(&readerVideoGoCV{readerOpts: readerOpts{fname: file}})
}

func (s *readerVideoGoCV) Read(o steps.SourceOptions) (img interface{}, path string, index int, timestamp int64, err error) {
	mat := backend.NewGoCVMat()
	s.dev.Read(mat)
	if !mat.Empty() {
		img = mat
	}

	timestamp = time.Now().UnixNano()
	index = s.index
	s.index++
	return
}

func (s *readerVideoGoCV) Reset() {
	s.index = 0
}

func (s *readerVideoGoCV) Open() error {
	Log.Debug().Str("file", s.fname).Msg("Open")
	dev, err := gocv.VideoCaptureFile(s.fname)
	if err != nil {
		Log.Error().Err(err).Str("file", s.fname)
		return err
	}
	s.index = 0
	s.dev = dev
	return nil
}

func (s *readerVideoGoCV) Close() {
	if s.dev != nil {
		s.dev.Close()
	}
}

func (s *readerVideoGoCV) Name() string {
	return VIDEO_GOCV_NAME
}