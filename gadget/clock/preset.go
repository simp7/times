package clock

import (
	"github.com/simp7/times/format"
	"time"
)

//Standard returns gadget.Clock that implements object.Standard and format.Standard adopting minimum unit as second. Its notation would be 12 so AM/PM notation would be added.
var (
	Standard = New(time.Second, format.Clock.Notation12)
)
