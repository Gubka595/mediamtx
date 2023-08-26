package message

import (
	"bytes"
	"testing"
	"time"

	"github.com/notedit/rtmp/format/flv/flvio"
	"github.com/stretchr/testify/require"

	"github.com/bluenviron/mediamtx/internal/rtmp/bytecounter"
)

var readWriterCases = []struct {
	name string
	dec  Message
	enc  []byte
}{
	{
		"acknowledge",
		&Acknowledge{
			Value: 45953968,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x3,
			0x0, 0x0, 0x0, 0x0, 0x2, 0xbd, 0x33, 0xb0,
		},
	},
	{
		"audio mpeg1",
		&Audio{
			ChunkStreamID:   7,
			DTS:             6013806 * time.Millisecond,
			MessageStreamID: 4534543,
			Codec:           CodecMPEG1Audio,
			Rate:            flvio.SOUND_44Khz,
			Depth:           flvio.SOUND_16BIT,
			Channels:        flvio.SOUND_STEREO,
			Payload:         []byte{0x01, 0x02, 0x03, 0x04},
		},
		[]byte{
			0x7, 0x5b, 0xc3, 0x6e, 0x0, 0x0, 0x5, 0x8, 0x0, 0x45, 0x31, 0xf, 0x2f,
			0x01, 0x02, 0x03, 0x04,
		},
	},
	{
		"audio mpeg4",
		&Audio{
			ChunkStreamID:   7,
			DTS:             6013806 * time.Millisecond,
			MessageStreamID: 4534543,
			Codec:           CodecMPEG4Audio,
			Rate:            flvio.SOUND_44Khz,
			Depth:           flvio.SOUND_16BIT,
			Channels:        flvio.SOUND_STEREO,
			AACType:         AudioAACTypeAU,
			Payload:         []byte{0x5A, 0xC0, 0x77, 0x40},
		},
		[]byte{
			0x7, 0x5b, 0xc3, 0x6e, 0x0, 0x0, 0x6, 0x8,
			0x0, 0x45, 0x31, 0xf, 0xaf, 0x1, 0x5a, 0xc0,
			0x77, 0x40,
		},
	},
	{
		"command amf0",
		&CommandAMF0{
			ChunkStreamID:   3,
			MessageStreamID: 345243,
			Name:            "i8yythrergre",
			CommandID:       56456,
			Arguments: []interface{}{
				flvio.AMFMap{
					{K: "k1", V: "v1"},
					{K: "k2", V: "v2"},
				},
				nil,
			},
		},
		[]byte{
			0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2f, 0x14,
			0x0, 0x5, 0x44, 0x9b, 0x2, 0x0, 0xc, 0x69,
			0x38, 0x79, 0x79, 0x74, 0x68, 0x72, 0x65, 0x72,
			0x67, 0x72, 0x65, 0x0, 0x40, 0xeb, 0x91, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x2, 0x6b,
			0x31, 0x2, 0x0, 0x2, 0x76, 0x31, 0x0, 0x2,
			0x6b, 0x32, 0x2, 0x0, 0x2, 0x76, 0x32, 0x0,
			0x0, 0x9, 0x5,
		},
	},
	{
		"data amf0",
		&DataAMF0{
			ChunkStreamID:   3,
			MessageStreamID: 345243,
			Payload: []interface{}{
				float64(234),
				"string",
				nil,
			},
		},
		[]byte{
			0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x13, 0x12,
			0x0, 0x5, 0x44, 0x9b, 0x0, 0x40, 0x6d, 0x40,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x6,
			0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x05,
		},
	},
	{
		"set chunk size",
		&SetChunkSize{
			Value: 10000,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
		},
	},
	{
		"set peer bandwidth",
		&SetChunkSize{
			Value: 10000,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
		},
	},
	{
		"set window ack size",
		&SetChunkSize{
			Value: 10000,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x1,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x27, 0x10,
		},
	},
	{
		"user control ping request",
		&UserControlPingRequest{
			ServerTime: 569834435,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x21, 0xf6,
			0xfb, 0xc3,
		},
	},
	{
		"user control ping response",
		&UserControlPingResponse{
			ServerTime: 569834435,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x21, 0xf6,
			0xfb, 0xc3,
		},
	},
	{
		"user control set buffer length",
		&UserControlSetBufferLength{
			StreamID:     35534,
			BufferLength: 235345,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0,
			0x8a, 0xce, 0x0, 0x3, 0x97, 0x51,
		},
	},
	{
		"user control stream begin",
		&UserControlStreamBegin{
			StreamID: 35534,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x8a, 0xce,
		},
	},
	{
		"user control stream dry",
		&UserControlStreamDry{
			StreamID: 35534,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0,
			0x8a, 0xce,
		},
	},
	{
		"user control stream eof",
		&UserControlStreamEOF{
			StreamID: 35534,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0,
			0x8a, 0xce,
		},
	},
	{
		"user control stream is recorded",
		&UserControlStreamIsRecorded{
			StreamID: 35534,
		},
		[]byte{
			0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x4,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0,
			0x8a, 0xce,
		},
	},
	{
		"video",
		&Video{
			ChunkStreamID:   6,
			DTS:             2543534 * time.Millisecond,
			MessageStreamID: 0x1000000,
			Codec:           CodecH264,
			IsKeyFrame:      true,
			Type:            VideoTypeConfig,
			PTSDelta:        10 * time.Millisecond,
			Payload:         []byte{0x01, 0x02, 0x03},
		},
		[]byte{
			0x06, 0x26, 0xcf, 0xae, 0x00, 0x00, 0x08, 0x09,
			0x01, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00,
			0x0a, 0x01, 0x02, 0x03,
		},
	},
	{
		"extended sequence start",
		&ExtendedSequenceStart{
			ChunkStreamID:   4,
			MessageStreamID: 0x1000000,
			FourCC:          FourCCHEVC,
			Config:          []byte{0x01, 0x02, 0x03},
		},
		[]byte{
			0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x09,
			0x01, 0x00, 0x00, 0x00, 0x80, 0x68, 0x76, 0x63,
			0x31, 0x01, 0x02, 0x03,
		},
	},
	{
		"extended coded frames",
		&ExtendedCodedFrames{
			ChunkStreamID:   4,
			DTS:             15100 * time.Millisecond,
			MessageStreamID: 0x1000000,
			FourCC:          FourCCHEVC,
			PTSDelta:        30 * time.Millisecond,
			Payload:         []byte{0x01, 0x02, 0x03},
		},
		[]byte{
			0x04, 0x00, 0x3a, 0xfc, 0x00, 0x00, 0x0b, 0x09,
			0x01, 0x00, 0x00, 0x00, 0x81, 0x68, 0x76, 0x63,
			0x31, 0x00, 0x00, 0x1e, 0x01, 0x02, 0x03,
		},
	},
	{
		"extended frames x",
		&ExtendedFramesX{
			ChunkStreamID:   4,
			DTS:             15100 * time.Millisecond,
			MessageStreamID: 0x1000000,
			FourCC:          FourCCHEVC,
			Payload:         []byte{0x01, 0x02, 0x03},
		},
		[]byte{
			0x04, 0x00, 0x3a, 0xfc, 0x00, 0x00, 0x08, 0x09,
			0x01, 0x00, 0x00, 0x00, 0x83, 0x68, 0x76, 0x63,
			0x31, 0x01, 0x02, 0x03,
		},
	},
}

func TestReader(t *testing.T) {
	for _, ca := range readWriterCases {
		t.Run(ca.name, func(t *testing.T) {
			bc := bytecounter.NewReader(bytes.NewReader(ca.enc))
			r := NewReader(bc, bc, nil)
			dec, err := r.Read()
			require.NoError(t, err)
			require.Equal(t, ca.dec, dec)
		})
	}
}

func FuzzReader(f *testing.F) {
	f.Add([]byte{
		0x04, 0x00, 0x3a, 0xfc, 0x00, 0x00, 0x08, 0x09,
		0x01, 0x00, 0x00, 0x00, 0x88, 0x68, 0x76, 0x63,
		0x31, 0x01, 0x02, 0x03,
	})
	f.Fuzz(func(t *testing.T, b []byte) {
		bc := bytecounter.NewReader(bytes.NewReader(b))
		r := NewReader(bc, bc, nil)
		r.Read() //nolint:errcheck
	})
}
