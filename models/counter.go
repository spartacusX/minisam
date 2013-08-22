package models

type SoftWareCounter struct {
	Name             string
	IsInternal       bool
	LicUseRights     int
	EntCount         int
	SoftInstallCount int
	UnusedInstall    int
}
