package cgo

//go:generate go run github.com/igadmg/colonization/cmd/gog "-ecs=$GOFILE"

// #cgo CXXFLAGS: -std=c++17
// #include "cgo.h"
import "C"
