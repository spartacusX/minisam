package models

type SoftWareCounter struct {
	LCounterId         int
	Name               string
	Status             string
	IsInternal         bool
	LicUseRights       int
	EntCount           int
	SoftInstallCount   int
	UnusedInstall      int
	LBrandId           int
	LContractId        int
	LQueryId           int
	LDescId            int
	LEntQueryId        int
	LLicQueryId        int
	LLicTypeId         int
	LCntrModelId       int
	LPreviousId        int
	LInstQueryId       int
	LSupervId          int
	LTenantId          int
	LUnusedSoftQueryId int
	LUpgFromCntrId     int
	LUpgToCntrId       int
	LAppDataVerId      int

	BAllLicenses       bool
	BAllSoftware       bool
	BAutomated         bool
	BCountDisappeared  bool
	BCountEnt          bool
	BCountInst         bool
	BCountLic          bool
	BCountSuiteCompo   bool
	BFamily            bool
	BHasPreviousCntr   bool
	BInternal          bool
	BLicRightsTransfer bool
	BLicUpgrade        bool
	BOldMode           bool
	BType              bool
	BUseCountFactor    bool
	BUseMetering       bool

	Code
	Context
	dCompliancy
	dCompliancyUpg
	dEntCount
	dEntitled
	dLicUseRights
	dLicUseRightsUpg
	dLicUseRightsUpgMx
	dRightsUpgTrans
	dSoftInstallCount
	dtCalc
	dUnusedInstall
	EntContext
	EntCountFormula
	EntCountSQL
	EntGroupBy
	InstContext
	InstCountFormula
	InstCountSQL
	InstGroupBy
	LicContext
	LicCountFormula
	LicCountSQL
	LicGroupBy
	LResultId
	Name
	seEntCountMode
	seInstallCountMode
	seInstallCountType
	seLicCountMode
	tsUnusedDuration
	Type
}

const (
	NORMAL = iota
	WARNING
	ERROR
)
