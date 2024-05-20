package record

import (
	"time"

	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/logger"
	"github.com/bluenviron/mediamtx/internal/storage"
	"github.com/bluenviron/mediamtx/internal/stream"
	"github.com/bluenviron/mediamtx/internal/errorSQL"
)

// Agent writes recordings to disk.
type Agent struct {
	WriteQueueSize    int
	PathFormat        string
	PathFormats       []string
	Format            conf.RecordFormat
	PartDuration      time.Duration
	SegmentDuration   time.Duration
	PathName          string
	StreamName        string
	Stream            *stream.Stream
	OnSegmentCreate   OnSegmentFunc
	OnSegmentComplete OnSegmentFunc
	Parent            logger.Writer

	restartPause time.Duration

	currentInstance *agentInstance

	terminate chan struct{}
	done      chan struct{}

	Stor        storage.Storage
	RecordAudio bool

	PathStream string
	CodeMp     string

	Filesqlerror *errorsql.Filesqlerror
}

// Initialize initializes Agent.
func (w *Agent) Initialize() {
	if w.OnSegmentCreate == nil {
		w.OnSegmentCreate = func(string) {
		}
	}
	if w.OnSegmentComplete == nil {
		w.OnSegmentComplete = func(string) {
		}
	}
	if w.restartPause == 0 {
		w.restartPause = 2 * time.Second
	}

	w.terminate = make(chan struct{})
	w.done = make(chan struct{})

	w.currentInstance = &agentInstance{
		agent:       w,
		stor:        w.Stor,
		recordAudio: w.RecordAudio,
	}
	w.currentInstance.initialize()

	go w.run()
}

// Log implements logger.Writer.
func (w *Agent) Log(level logger.Level, format string, args ...interface{}) {
	w.Parent.Log(level, "[record] "+format, args...)
}

// Close closes the agent.
func (w *Agent) Close() {
	w.Log(logger.Info, "recording stopped")
	close(w.terminate)
	<-w.done
}

func (w *Agent) run() {
	defer close(w.done)

	for {
		select {
		case <-w.currentInstance.done:
			w.currentInstance.close()
		case <-w.terminate:
			w.currentInstance.close()
			return
		}

		select {
		case <-time.After(w.restartPause):
		case <-w.terminate:
			return
		}

		w.currentInstance = &agentInstance{
			agent:       w,
			stor:        w.Stor,
			recordAudio: w.RecordAudio,
		}
		w.currentInstance.initialize()
	}
}
