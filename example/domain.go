package example

import (
	"time"
)

type Gender int



const (
	Man Gender = iota
	Women
	Unknown
)

type TimeStamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// +jam /clefs/slicer/struct.go
type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Gender    Gender
	TimeStamps
}

// +jam /clefs/slicer/struct.go
type City struct {
	PrefectureName string
	Name           string
	Population     int
	IsCapital      bool
}
