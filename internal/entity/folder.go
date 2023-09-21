package entity

import "time"

type Folder struct {
	Name        string
	Description string
	CreatedAt   time.Time

	Files []*File
}

type ListFolderOption struct {
	Sort SortOption
}

type SortOption struct {
	Attribute SortAttribute
	Direction SortDirection
}

type CreateFolderParam struct {
	Username    string
	Name        string
	Description string
}
