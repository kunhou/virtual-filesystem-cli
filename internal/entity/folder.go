package entity

import "time"

type Folder struct {
	Name        string
	Description string
	CreatedAt   time.Time

	Files []*File
}

type ListFolderOption struct {
	Sort FolderSort
}

type FolderSort struct {
	Attribute SortAttribute
	Direction SortDirection
}
