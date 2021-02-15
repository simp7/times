package clock

import (
	"github.com/simp7/times"
	"github.com/simp7/times/formatter"
)

//Standard returns Clock that implements time.Standard and formatter.Standard adopting minimum unit as second. Its notation would be 12 so AM/PM notation would be added.
var (
	Standard = New(times.Sec, formatter.Clock(true))
)
