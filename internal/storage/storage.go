package storage

import "errors"

var (
	ErrorUrlNotFound = errors.New("url not found")
	ErrorUrlExists   = errors.New("url exists")
)
