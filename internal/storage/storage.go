package storage

import "errors"

var (
	ErrorUrlExists = errors.New("url exists")
	ErrorEmptyUrls = errors.New("no one existing alias")
)
