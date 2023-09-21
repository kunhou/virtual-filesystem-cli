package entity

type SortAttribute int

const (
	SortByName SortAttribute = iota
	SortByCreateTime

	ListResourceTimeFormat = "2006-01-02 15:04:05"
)

type SortDirection int

const (
	Asc SortDirection = iota
	Desc
)
