package models

type SoftWareCounter struct {
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
	Code               string
	Context            string
	EntContext         string
	EntCountFormula    string
	EntCountSQL        string
	EntGroupBy         string
	InstContext        string
	InstCountFormula   string
	InstCountSQL       string
	InstGroupBy        string
	IsInternal         bool
	LAppDataVerId      int
	LBrandId           int
	LCntrModelId       int
	LCompliancy        int
	LCompliancyUpg     int
	LContractId        int
	LCounterId         int
	LDescId            int
	LEntCount          int
	LEntitled          int
	LEntQueryId        int
	LicContext         string
	LicCountFormula    string
	LicCountSQL        string
	LicGroupBy         string
	LInstQueryId       int
	LLicQueryId        int
	LLicTypeId         int
	LLicUseRights      int
	LLicUseRightsUpg   int
	LLicUseRightsUpgMx int
	LPreviousId        int
	LQueryId           int
	LResultId          int
	LRightsUpgTrans    int
	LSoftInstallCount  int
	LSupervId          int
	LtCalc             int
	LTenantId          int
	LUnusedInstall     int
	LUnusedSoftQueryId int
	LUpgFromCntrId     int
	LUpgToCntrId       int
	Name               string
	SeEntCountMode     int
	SeInstallCountMode int
	SeInstallCountType int
	SeLicCountMode     int
	Status             string
	TsUnusedDuration   int
	Type               string
	UnusedInstall      int
}

func init() {

}

const (
	NORMAL = iota
	WARNING
	ERROR
)
