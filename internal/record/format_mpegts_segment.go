package record

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bluenviron/mediamtx/internal/logger"
)

type formatMPEGTSSegment struct {
	f         *formatMPEGTS
	startDTS  time.Duration
	lastFlush time.Duration

	created time.Time
	path    string
	fi      *os.File
}

func newFormatMPEGTSSegment(f *formatMPEGTS, startDTS time.Duration) *formatMPEGTSSegment {
	s := &formatMPEGTSSegment{
		f:         f,
		startDTS:  startDTS,
		lastFlush: startDTS,
		created:   timeNow(),
	}

	f.dw.setTarget(s)

	return s
}

func (s *formatMPEGTSSegment) close() error {
	err := s.f.bw.Flush()

	if s.fi != nil {
		s.f.a.agent.Log(logger.Debug, "closing segment %s", s.path)
		err2 := s.fi.Close()
		if err == nil {
			err = err2
		}

		if err2 == nil {
			s.f.a.agent.OnSegmentComplete(s.path)

			if s.f.a.stor.Use {
				stat, err3 := os.Stat(s.path)
				if err3 == nil {
					paths := strings.Split(s.path, "/")
					err4 := s.f.a.stor.Req.ExecQuery(
						fmt.Sprintf(
							s.f.a.stor.Sql.UpdateSize,
							fmt.Sprint(stat.Size()),
							time.Now().Format("2006-01-02 15:04:05"),
							paths[len(paths)-1],
						))
					if err4 != nil {
						return err4
					}
					return err
				}
				err = err3
			}
		}
	}

	return err
}

func (s *formatMPEGTSSegment) Write(p []byte) (int, error) {
	if s.fi == nil {

		var err error

		if s.f.a.stor.DbDrives {

			data, err := s.f.a.stor.Req.SelectData(s.f.a.stor.Sql.GetDrives)

			if err != nil {
				return 0, err
			}
			for _, line := range data {
				drives = append(drives, line[0].(string))
			}
			s.f.a.free = getMostFreeDisk(drives)
			if s.f.a.stor.DbUseCodeMP && s.f.a.stor.UseDbPathStream {
				s.f.a.codeMp, err = s.f.a.stor.Req.SelectPathStream(fmt.Sprintf(s.f.a.stor.Sql.GetCodeMP, s.f.a.agent.StreamName))
				if err != nil {
					return 0, err
				}
				s.f.a.pathStream, err = s.f.a.stor.Req.SelectPathStream(fmt.Sprintf(s.f.a.stor.Sql.GetPathStream, s.f.a.agent.StreamName))
				if err != nil {
					return 0, err
				}
				s.path = fmt.Sprintf(s.f.a.free+path(s.created).encode(s.f.a.pathFormat), s.f.a.codeMp, s.f.a.pathStream)
			}

			if s.f.a.stor.DbUseCodeMP {
				s.f.a.codeMp, err = s.f.a.stor.Req.SelectPathStream(fmt.Sprintf(s.f.a.stor.Sql.GetCodeMP, s.f.a.agent.StreamName))
				if err != nil {
					return 0, err
				}
				s.path = fmt.Sprintf(s.f.a.free+path(s.created).encode(s.f.a.pathFormat), s.f.a.codeMp)
			}
			if s.f.a.stor.UseDbPathStream {
				s.f.a.pathStream, err = s.f.a.stor.Req.SelectPathStream(fmt.Sprintf(s.f.a.stor.Sql.GetPathStream, s.f.a.agent.StreamName))
				if err != nil {
					return 0, err
				}
				s.path = fmt.Sprintf(s.f.a.free+path(s.created).encode(s.f.a.pathFormat), s.f.a.pathStream)
			}
			if !s.f.a.stor.DbUseCodeMP && !s.f.a.stor.UseDbPathStream {
				s.path = fmt.Sprintf(s.f.a.free + path(s.created).encode(s.f.a.pathFormat))
			}
		} else {
			s.path = path(s.created).encode(s.f.a.pathFormat)
		}

		s.f.a.agent.Log(logger.Debug, "creating segment %s", s.path)

		err = os.MkdirAll(filepath.Dir(s.path), 0o755)
		if err != nil {
			return 0, err
		}

		fi, err := os.Create(s.path)
		if err != nil {
			return 0, err
		}

		if s.f.a.stor.Use {
			paths := strings.Split(s.path, "/")
			pathRec := strings.Join(paths[:len(paths)-1], "/")
			if s.f.a.stor.UseDbPathStream {
				err := s.f.a.stor.Req.ExecQuery(
					fmt.Sprintf(
						s.f.a.stor.Sql.InsertPath,
						"pathStream",
						pathRec+"/",
						paths[len(paths)-1],
						time.Now().Format("2006-01-02 15:04:05"),
						s.f.a.pathStream,
					),
				)
				if err != nil {
					os.Remove(s.path)
					return 0, err
				}
			} else {
				err := s.f.a.stor.Req.ExecQuery(
					fmt.Sprintf(
						s.f.a.stor.Sql.InsertPath,
						"stream",
						pathRec+"/",
						paths[len(paths)-1],
						time.Now().Format("2006-01-02 15:04:05"),
						s.f.a.agent.PathName,
					),
				)
				if err != nil {
					os.Remove(s.path)
					return 0, err
				}
			}

		}

		s.f.a.agent.OnSegmentCreate(s.path)

		s.fi = fi
	}

	return s.fi.Write(p)
}
