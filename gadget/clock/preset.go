package clock

import (
	"github.com/simp7/times/formatter"
	"time"
)

//Standard returns gadget.Clock that implements timeobject.Standard and formatter.Standard adopting minimum unit as second. Its notation would be 12 so AM/PM notation would be added.
var (
	Standard = New(time.Second, formatter.Clock(true))
)
