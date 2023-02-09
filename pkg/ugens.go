package ugens

import (
	"github.com/CameronGorrie/sc"
	"github.com/CameronGorrie/ugens/pkg/mpe"
	"github.com/CameronGorrie/ugens/pkg/samples"
	"github.com/CameronGorrie/ugens/pkg/simple"
	"github.com/CameronGorrie/ugens/pkg/soundin"
)

var Lib = map[string]sc.UgenFunc{
	"bpf":   mpe.BPF,
	"buf":   samples.Buf,
	"mic":   soundin.Mic,
	"delay": simple.Delay,
	"saw":   simple.SimpleSaw,
	"sine":  simple.SimpleSine,
}