package format

import (
	"github.com/grupopulpo/vdk/av/avutil"
	"github.com/grupopulpo/vdk/format/aac"
	"github.com/grupopulpo/vdk/format/flv"
	"github.com/grupopulpo/vdk/format/mp4"
	"github.com/grupopulpo/vdk/format/rtmp"
	"github.com/grupopulpo/vdk/format/rtsp"
	"github.com/grupopulpo/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
