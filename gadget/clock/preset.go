package clock

import (
	"github.com/simp7/times"
	"github.com/simp7/times/formatter"
)

//Standard returns gadget.Clock that implements ttime.Standard and formatter.Standard adopting minimum unit as second. Its notation would be 12 so AM/PM notation would be added.
var (
	Standard = New(times.Sec, formatter.Clock(true))
)
