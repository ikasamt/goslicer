package goslicer

import "time"

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

type User struct {
	ID string
	Age int
	FirstName string
	LastName string
	Gender Gender
	TimeStamps
}

type Item struct {
	ID string
	Name string
	TimeStamps
}

type Product struct {
	ID string
	ShortName string
	Name string
	ItemID int
	TimeStamps
}