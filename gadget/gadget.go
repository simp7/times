package gadget

import "github.com/simp7/times/model/tobject"

type Gadget interface {
	Add(action func(current string))
	AddAlarm(action func(current string), when tobject.Time)
	Start()
	End() string
}
