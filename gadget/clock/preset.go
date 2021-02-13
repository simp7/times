package clock

import (
	"github.com/simp7/times/model/formatter"
	"github.com/simp7/times/model/tobject"
)

//Standard returns Clock that implements tobject.Standard and formatter.Standard adopting minimum unit as second. Its notation would be 12 so AM/PM notation would be added.
var (
	Standard = New(tobject.Sec, formatter.Clock(true))
)
