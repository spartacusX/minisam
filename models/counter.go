package models

type SoftWareCounter struct {
	Name             string
	Status           string
	IsInternal       bool
	LicUseRights     int
	EntCount         int
	SoftInstallCount int
	UnusedInstall    int
}

const (
	NORMAL = iota
	WARNING
	ERROR
)
