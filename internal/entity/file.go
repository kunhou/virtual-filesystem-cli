package entity

import "time"

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

type ListFileOption struct {
	Sort SortOption
}

type CreateFileParam struct {
	Username   string
	FolderName string

	Name        string
	Description string
}
