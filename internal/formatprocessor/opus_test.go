package formatprocessor

import (
	"testing"

	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/mediamtx/internal/unit"
	"github.com/pion/rtp"
	"github.com/stretchr/testify/require"
)

func TestOpusEncode(t *testing.T) {
	forma := &format.Opus{
		PayloadTyp:   96,
		ChannelCount: 2,
	}

	p, err := New(1472, forma, true)
	require.NoError(t, err)

	unit := &unit.Opus{
		Packets: [][]byte{
			{
				0xfc, 0x1e, 0x61, 0x96, 0xfc, 0xf7, 0x9b, 0x23,
				0x5b, 0xc9, 0x56, 0xad, 0x05, 0x12, 0x2f, 0x6c,
				0xc0, 0x0c, 0x2c, 0x17, 0x7b, 0x1a, 0xde, 0x1b,
				0x37, 0x89, 0xc5, 0xbb, 0x34, 0xbb, 0x1c, 0x74,
				0x7c, 0x18, 0x0a, 0xde, 0xa1, 0x2b, 0x86, 0x1d,
				0x60, 0xa2, 0xb6, 0xce, 0xe7, 0x0e, 0x17, 0x1b,
				0xc7, 0xd4, 0xd1, 0x2a, 0x68, 0x1f, 0x05, 0x2b,
				0x22, 0x80, 0x68, 0x12, 0x0c, 0x45, 0xbc, 0x3a,
				0xd2, 0x1b, 0xf2, 0x8a, 0x77, 0x5f, 0x2b, 0x34,
				0x97, 0x34, 0x09, 0x6d, 0x05, 0x5f, 0x48, 0x0c,
				0x45, 0xb5, 0xae, 0x2a, 0x90, 0x21, 0xda, 0xfb,
				0x5b, 0x10,
			},
			{
				0xfc, 0x1d, 0x61, 0x96, 0xfa, 0x7a, 0x90, 0x59,
				0xb7, 0x10, 0xd7, 0x03, 0x84, 0x27, 0x3f, 0x52,
				0x9f, 0xd7, 0x38, 0x9f, 0xbc, 0xff, 0x7d, 0x62,
				0xe8, 0x60, 0x64, 0x54, 0x9d, 0x1a, 0xd1, 0x7c,
				0x1a, 0x0a, 0x76, 0xd2, 0x30, 0x9b, 0xbe, 0xc7,
				0x5d, 0x37, 0x42, 0xe7, 0xdd, 0xfc, 0xc7, 0x03,
				0xab, 0x90, 0xd9, 0x9b, 0xad, 0xf4, 0x88, 0xd3,
				0x81, 0xbf, 0xd0, 0x68, 0x10, 0x0c, 0x46, 0xa4,
				0xe8, 0x83, 0xd9, 0x6b, 0x7a, 0x25, 0xed, 0x81,
				0xf6, 0x92, 0x14, 0x70, 0x6c, 0x48, 0x0c, 0x45,
				0xb4, 0xf2, 0x61, 0x9b, 0xd7, 0x62, 0x58, 0x87,
			},
			{
				0xfc, 0x1e, 0x61, 0x96, 0xfa, 0x7f, 0x61, 0x51,
				0x2f, 0x10, 0x81, 0x22, 0x01, 0x81, 0x46, 0x5e,
				0xbd, 0xb0, 0x87, 0xcb, 0x0a, 0xa6, 0xd6, 0xd3,
				0xec, 0x7c, 0x9e, 0xf5, 0x07, 0x5a, 0x07, 0x1b,
				0x7c, 0x19, 0x0a, 0xde, 0xa1, 0x38, 0xe4, 0x51,
				0x7f, 0x54, 0x6f, 0x91, 0x9f, 0xda, 0x2b, 0x40,
				0x80, 0x36, 0xeb, 0xe3, 0xc2, 0x58, 0x12, 0x55,
				0x80, 0x65, 0x14, 0x68, 0x12, 0x0c, 0x46, 0xa4,
				0xdc, 0x6e, 0x62, 0x79, 0xf9, 0x09, 0x28, 0x11,
				0xab, 0xab, 0xd3, 0x84, 0x7d, 0x93, 0x34, 0x48,
				0x0c, 0x46, 0x9a, 0x0c, 0xb0, 0xfe, 0x7a, 0x5b,
				0x22, 0x85, 0x4a, 0x73, 0x0d, 0xd0,
			},
		},
	}

	err = p.ProcessUnit(unit)
	require.NoError(t, err)
	require.Equal(t, []*rtp.Packet{
		{
			Header: rtp.Header{
				Version:        2,
				PayloadType:    96,
				SequenceNumber: unit.RTPPackets[0].SequenceNumber,
				Timestamp:      unit.RTPPackets[0].Timestamp,
				SSRC:           unit.RTPPackets[0].SSRC,
			},
			Payload: []byte{
				0xfc, 0x1e, 0x61, 0x96, 0xfc, 0xf7, 0x9b, 0x23,
				0x5b, 0xc9, 0x56, 0xad, 0x05, 0x12, 0x2f, 0x6c,
				0xc0, 0x0c, 0x2c, 0x17, 0x7b, 0x1a, 0xde, 0x1b,
				0x37, 0x89, 0xc5, 0xbb, 0x34, 0xbb, 0x1c, 0x74,
				0x7c, 0x18, 0x0a, 0xde, 0xa1, 0x2b, 0x86, 0x1d,
				0x60, 0xa2, 0xb6, 0xce, 0xe7, 0x0e, 0x17, 0x1b,
				0xc7, 0xd4, 0xd1, 0x2a, 0x68, 0x1f, 0x05, 0x2b,
				0x22, 0x80, 0x68, 0x12, 0x0c, 0x45, 0xbc, 0x3a,
				0xd2, 0x1b, 0xf2, 0x8a, 0x77, 0x5f, 0x2b, 0x34,
				0x97, 0x34, 0x09, 0x6d, 0x05, 0x5f, 0x48, 0x0c,
				0x45, 0xb5, 0xae, 0x2a, 0x90, 0x21, 0xda, 0xfb,
				0x5b, 0x10,
			},
		},
		{
			Header: rtp.Header{
				Version:        2,
				PayloadType:    96,
				SequenceNumber: unit.RTPPackets[0].SequenceNumber + 1,
				Timestamp:      unit.RTPPackets[0].Timestamp + 960,
				SSRC:           unit.RTPPackets[0].SSRC,
			},
			Payload: []byte{
				0xfc, 0x1d, 0x61, 0x96, 0xfa, 0x7a, 0x90, 0x59,
				0xb7, 0x10, 0xd7, 0x03, 0x84, 0x27, 0x3f, 0x52,
				0x9f, 0xd7, 0x38, 0x9f, 0xbc, 0xff, 0x7d, 0x62,
				0xe8, 0x60, 0x64, 0x54, 0x9d, 0x1a, 0xd1, 0x7c,
				0x1a, 0x0a, 0x76, 0xd2, 0x30, 0x9b, 0xbe, 0xc7,
				0x5d, 0x37, 0x42, 0xe7, 0xdd, 0xfc, 0xc7, 0x03,
				0xab, 0x90, 0xd9, 0x9b, 0xad, 0xf4, 0x88, 0xd3,
				0x81, 0xbf, 0xd0, 0x68, 0x10, 0x0c, 0x46, 0xa4,
				0xe8, 0x83, 0xd9, 0x6b, 0x7a, 0x25, 0xed, 0x81,
				0xf6, 0x92, 0x14, 0x70, 0x6c, 0x48, 0x0c, 0x45,
				0xb4, 0xf2, 0x61, 0x9b, 0xd7, 0x62, 0x58, 0x87,
			},
		},
		{
			Header: rtp.Header{
				Version:        2,
				PayloadType:    96,
				SequenceNumber: unit.RTPPackets[0].SequenceNumber + 2,
				Timestamp:      unit.RTPPackets[0].Timestamp + 960*2,
				SSRC:           unit.RTPPackets[0].SSRC,
			},
			Payload: []byte{
				0xfc, 0x1e, 0x61, 0x96, 0xfa, 0x7f, 0x61, 0x51,
				0x2f, 0x10, 0x81, 0x22, 0x01, 0x81, 0x46, 0x5e,
				0xbd, 0xb0, 0x87, 0xcb, 0x0a, 0xa6, 0xd6, 0xd3,
				0xec, 0x7c, 0x9e, 0xf5, 0x07, 0x5a, 0x07, 0x1b,
				0x7c, 0x19, 0x0a, 0xde, 0xa1, 0x38, 0xe4, 0x51,
				0x7f, 0x54, 0x6f, 0x91, 0x9f, 0xda, 0x2b, 0x40,
				0x80, 0x36, 0xeb, 0xe3, 0xc2, 0x58, 0x12, 0x55,
				0x80, 0x65, 0x14, 0x68, 0x12, 0x0c, 0x46, 0xa4,
				0xdc, 0x6e, 0x62, 0x79, 0xf9, 0x09, 0x28, 0x11,
				0xab, 0xab, 0xd3, 0x84, 0x7d, 0x93, 0x34, 0x48,
				0x0c, 0x46, 0x9a, 0x0c, 0xb0, 0xfe, 0x7a, 0x5b,
				0x22, 0x85, 0x4a, 0x73, 0x0d, 0xd0,
			},
		},
	}, unit.RTPPackets)
}
