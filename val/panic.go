package val

import (
	"github.com/nanoeru/copal/val/vpanic"
)

var boolPanic = vpanic.NewValPanic("bool")
var chanPanic = vpanic.NewValPanic("chan")
var intPanic = vpanic.NewValPanic("int")
var floatPanic = vpanic.NewValPanic("float")
var stringPanic = vpanic.NewValPanic("string")
var funcPanic = vpanic.NewValPanic("func")
var mapPanic = vpanic.NewValPanic("map")
var nilPanic = vpanic.NewValPanic("nil")
var slicePanic = vpanic.NewValPanic("slice")
