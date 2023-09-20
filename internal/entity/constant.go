package entity

type SortAttribute int

const (
	SortByName SortAttribute = iota
	SortByCreateTime
)

type SortDirection int

const (
	Asc SortDirection = iota
	Desc
)
