package formatter

import (
	"github.com/simp7/times/model/tobject"
)

type TimeFormatter interface {
	Format(t tobject.Time) string
}
