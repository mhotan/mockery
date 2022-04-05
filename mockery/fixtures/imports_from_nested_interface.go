package test

import (
	"github.com/EngHabu/mockery/mockery/fixtures/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
