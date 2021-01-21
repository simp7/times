package formatter

import (
	"times/model/tobject"
)

type TimeFormatter interface {
	Format(t tobject.Time) string
}