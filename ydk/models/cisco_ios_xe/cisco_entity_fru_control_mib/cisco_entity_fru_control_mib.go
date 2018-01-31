// The CISCO-ENTITY-FRU-CONTROL-MIB is used to monitor
// and configure operational status of 
// Field Replaceable Units (FRUs) and other managable 
// physical entities of the system listed in the 
// Entity-MIB (RFC 2737) entPhysicalTable. 
// 
// FRUs include assemblies such as power supplies, fans, 
// processor modules, interface modules, etc.
package cisco_entity_fru_control_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_entity_fru_control_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENTITY-FRU-CONTROL-MIB CISCO-ENTITY-FRU-CONTROL-MIB}", reflect.TypeOf(CISCOENTITYFRUCONTROLMIB{}))
    ydk.RegisterEntity("CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB", reflect.TypeOf(CISCOENTITYFRUCONTROLMIB{}))
}

// ModuleAdminType represents                        service, set by CLI.
type ModuleAdminType string

const (
    ModuleAdminType_enabled ModuleAdminType = "enabled"

    ModuleAdminType_disabled ModuleAdminType = "disabled"

    ModuleAdminType_reset ModuleAdminType = "reset"

    ModuleAdminType_outOfServiceAdmin ModuleAdminType = "outOfServiceAdmin"
)

// FRUCoolingUnit represents watts(2)  Watts
type FRUCoolingUnit string

const (
    FRUCoolingUnit_cfm FRUCoolingUnit = "cfm"

    FRUCoolingUnit_watts FRUCoolingUnit = "watts"
)

// PowerAdminType represents in response to a management protocol retrieval operation.
type PowerAdminType string

const (
    PowerAdminType_on PowerAdminType = "on"

    PowerAdminType_off PowerAdminType = "off"

    PowerAdminType_inlineAuto PowerAdminType = "inlineAuto"

    PowerAdminType_inlineOn PowerAdminType = "inlineOn"

    PowerAdminType_powerCycle PowerAdminType = "powerCycle"
)

// ModuleOperType represents fwDownloadFailure(27) Module firmware download failed.
type ModuleOperType string

const (
    ModuleOperType_unknown ModuleOperType = "unknown"

    ModuleOperType_ok ModuleOperType = "ok"

    ModuleOperType_disabled ModuleOperType = "disabled"

    ModuleOperType_okButDiagFailed ModuleOperType = "okButDiagFailed"

    ModuleOperType_boot ModuleOperType = "boot"

    ModuleOperType_selfTest ModuleOperType = "selfTest"

    ModuleOperType_failed ModuleOperType = "failed"

    ModuleOperType_missing ModuleOperType = "missing"

    ModuleOperType_mismatchWithParent ModuleOperType = "mismatchWithParent"

    ModuleOperType_mismatchConfig ModuleOperType = "mismatchConfig"

    ModuleOperType_diagFailed ModuleOperType = "diagFailed"

    ModuleOperType_dormant ModuleOperType = "dormant"

    ModuleOperType_outOfServiceAdmin ModuleOperType = "outOfServiceAdmin"

    ModuleOperType_outOfServiceEnvTemp ModuleOperType = "outOfServiceEnvTemp"

    ModuleOperType_poweredDown ModuleOperType = "poweredDown"

    ModuleOperType_poweredUp ModuleOperType = "poweredUp"

    ModuleOperType_powerDenied ModuleOperType = "powerDenied"

    ModuleOperType_powerCycled ModuleOperType = "powerCycled"

    ModuleOperType_okButPowerOverWarning ModuleOperType = "okButPowerOverWarning"

    ModuleOperType_okButPowerOverCritical ModuleOperType = "okButPowerOverCritical"

    ModuleOperType_syncInProgress ModuleOperType = "syncInProgress"

    ModuleOperType_upgrading ModuleOperType = "upgrading"

    ModuleOperType_okButAuthFailed ModuleOperType = "okButAuthFailed"

    ModuleOperType_mdr ModuleOperType = "mdr"

    ModuleOperType_fwMismatchFound ModuleOperType = "fwMismatchFound"

    ModuleOperType_fwDownloadSuccess ModuleOperType = "fwDownloadSuccess"

    ModuleOperType_fwDownloadFailure ModuleOperType = "fwDownloadFailure"
)

// ModuleResetReasonType represents                                 violation.
type ModuleResetReasonType string

const (
    ModuleResetReasonType_unknown ModuleResetReasonType = "unknown"

    ModuleResetReasonType_powerUp ModuleResetReasonType = "powerUp"

    ModuleResetReasonType_parityError ModuleResetReasonType = "parityError"

    ModuleResetReasonType_clearConfigReset ModuleResetReasonType = "clearConfigReset"

    ModuleResetReasonType_manualReset ModuleResetReasonType = "manualReset"

    ModuleResetReasonType_watchDogTimeoutReset ModuleResetReasonType = "watchDogTimeoutReset"

    ModuleResetReasonType_resourceOverflowReset ModuleResetReasonType = "resourceOverflowReset"

    ModuleResetReasonType_missingTaskReset ModuleResetReasonType = "missingTaskReset"

    ModuleResetReasonType_lowVoltageReset ModuleResetReasonType = "lowVoltageReset"

    ModuleResetReasonType_controllerReset ModuleResetReasonType = "controllerReset"

    ModuleResetReasonType_systemReset ModuleResetReasonType = "systemReset"

    ModuleResetReasonType_switchoverReset ModuleResetReasonType = "switchoverReset"

    ModuleResetReasonType_upgradeReset ModuleResetReasonType = "upgradeReset"

    ModuleResetReasonType_downgradeReset ModuleResetReasonType = "downgradeReset"

    ModuleResetReasonType_cacheErrorReset ModuleResetReasonType = "cacheErrorReset"

    ModuleResetReasonType_deviceDriverReset ModuleResetReasonType = "deviceDriverReset"

    ModuleResetReasonType_softwareExceptionReset ModuleResetReasonType = "softwareExceptionReset"

    ModuleResetReasonType_restoreConfigReset ModuleResetReasonType = "restoreConfigReset"

    ModuleResetReasonType_abortRevReset ModuleResetReasonType = "abortRevReset"

    ModuleResetReasonType_burnBootReset ModuleResetReasonType = "burnBootReset"

    ModuleResetReasonType_standbyCdHealthierReset ModuleResetReasonType = "standbyCdHealthierReset"

    ModuleResetReasonType_nonNativeConfigClearReset ModuleResetReasonType = "nonNativeConfigClearReset"

    ModuleResetReasonType_memoryProtectionErrorReset ModuleResetReasonType = "memoryProtectionErrorReset"
)

// PowerRedundancyType represents     support power redundancy with single input.
type PowerRedundancyType string

const (
    PowerRedundancyType_notsupported PowerRedundancyType = "notsupported"

    PowerRedundancyType_redundant PowerRedundancyType = "redundant"

    PowerRedundancyType_combined PowerRedundancyType = "combined"

    PowerRedundancyType_nonRedundant PowerRedundancyType = "nonRedundant"

    PowerRedundancyType_psRedundant PowerRedundancyType = "psRedundant"

    PowerRedundancyType_inPwrSrcRedundant PowerRedundancyType = "inPwrSrcRedundant"

    PowerRedundancyType_psRedundantSingleInput PowerRedundancyType = "psRedundantSingleInput"
)

// PowerOperType represents                           FRU has failed.
type PowerOperType string

const (
    PowerOperType_offEnvOther PowerOperType = "offEnvOther"

    PowerOperType_on PowerOperType = "on"

    PowerOperType_offAdmin PowerOperType = "offAdmin"

    PowerOperType_offDenied PowerOperType = "offDenied"

    PowerOperType_offEnvPower PowerOperType = "offEnvPower"

    PowerOperType_offEnvTemp PowerOperType = "offEnvTemp"

    PowerOperType_offEnvFan PowerOperType = "offEnvFan"

    PowerOperType_failed PowerOperType = "failed"

    PowerOperType_onButFanFail PowerOperType = "onButFanFail"

    PowerOperType_offCooling PowerOperType = "offCooling"

    PowerOperType_offConnectorRating PowerOperType = "offConnectorRating"

    PowerOperType_onButInlinePowerFail PowerOperType = "onButInlinePowerFail"
)

// CISCOENTITYFRUCONTROLMIB
type CISCOENTITYFRUCONTROLMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cefcfrupower CISCOENTITYFRUCONTROLMIB_Cefcfrupower

    
    Cefcmibnotificationenables CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables

    // This table lists the redundancy mode and the operational status of the
    // power supply groups in the system.
    Cefcfrupowersupplygrouptable CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable

    // This table lists the power-related administrative status and operational
    // status of the manageable components in the system.
    Cefcfrupowerstatustable CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable

    // This table lists the power capacity of a power FRU in the system if it
    // provides variable power. Power supplies usually provide either system or
    // inline power. They cannot be  controlled by software to dictate how they
    // distribute power. We can also have what are known as variable power
    // supplies. They can provide both system and inline power and can be  varied
    // within hardware defined ranges for system and inline limited by a total
    // maximum combined output. They could be configured by the user via CLI or
    // SNMP or be controlled by software internally. This table supplements the
    // information in the cefcFRUPowerStatusTable for power supply FRUs. The 
    // cefcFRUCurrent attribute in that table provides the overall current the
    // power supply FRU can provide while this table  gives us the individual
    // contribution towards system and  inline power.
    Cefcfrupowersupplyvaluetable CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable

    // A cefcModuleTable entry lists the operational and administrative status
    // information for ENTITY-MIB entPhysicalTable entries for manageable
    // components of type PhysicalClass module(9).
    Cefcmoduletable CISCOENTITYFRUCONTROLMIB_Cefcmoduletable

    // This table sparsely augments the cefcModuleTable (i.e., every row in this
    // table corresponds to a row in the cefcModuleTable but not necessarily
    // vice-versa).  A cefcIntelliModuleTable entry lists the information specific
    // to intelligent modules which cannot be provided by the cefcModuleTable.
    Cefcintellimoduletable CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable

    // This table sparsely augments the cefcModuleTable (i.e., every row in this
    // table corresponds to a row in the cefcModuleTable but not necessarily
    // vice-versa).  A cefcModuleLocalSwitchingTable entry lists the information
    // specific to local switching capable modules which cannot be provided by the
    // cefcModuleTable.
    Cefcmodulelocalswitchingtable CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable

    // This table contains the operational status information for all ENTITY-MIB
    // entPhysicalTable entries which have  an entPhysicalClass of 'fan';
    // specifically, all   entPhysicalTable entries which represent either: one 
    // physical fan, or a single physical 'fan tray' which is a manufactured
    // (inseparable in the field) combination of  multiple fans.
    Cefcfantraystatustable CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable

    // This table contains one row per physical entity.
    Cefcphysicaltable CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable

    // This table contains the power input information for all the power supplies
    // that have entPhysicalTable entries with 'powerSupply' in the
    // entPhysicalClass.   The entries are created by the agent at the system
    // power-up or power supply insertion.  Entries are deleted by the agent upon
    // power supply removal.  The number of entries is determined by the number of
    // power supplies and number of power inputs on the power  supply.
    Cefcpowersupplyinputtable CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable

    // This table contains a list of possible output mode for the power supplies,
    // whose ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of
    // 'powerSupply'. It also indicate which mode is the operational mode within
    // the system.
    Cefcpowersupplyoutputtable CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable

    // This table contains the cooling capacity information of the chassis whose
    // ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of 'chassis'.
    Cefcchassiscoolingtable CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable

    // This table contains the cooling capacity information of the fans whose
    // ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of 'fan'.
    Cefcfancoolingtable CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable

    // This table contains the cooling requirement for all the manageable
    // components of type entPhysicalClass 'module'.
    Cefcmodulecoolingtable CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable

    // This table contains a list of the possible cooling capacity modes and
    // properties of the fans, whose  ENTITY-MIB entPhysicalTable entries have an 
    // entPhysicalClass of 'fan'.
    Cefcfancoolingcaptable CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable

    // This table contains the connector power ratings of FRUs.   Only components
    // with power connector rating  management are listed in this table.
    Cefcconnectorratingtable CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable

    // This table contains the total power consumption information for modules
    // whose ENTITY-MIB  entPhysicalTable entries have an entPhysicalClass  of
    // 'module'.
    Cefcmodulepowerconsumptiontable CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetFilter() yfilter.YFilter { return cISCOENTITYFRUCONTROLMIB.YFilter }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) SetFilter(yf yfilter.YFilter) { cISCOENTITYFRUCONTROLMIB.YFilter = yf }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetGoName(yname string) string {
    if yname == "cefcFRUPower" { return "Cefcfrupower" }
    if yname == "cefcMIBNotificationEnables" { return "Cefcmibnotificationenables" }
    if yname == "cefcFRUPowerSupplyGroupTable" { return "Cefcfrupowersupplygrouptable" }
    if yname == "cefcFRUPowerStatusTable" { return "Cefcfrupowerstatustable" }
    if yname == "cefcFRUPowerSupplyValueTable" { return "Cefcfrupowersupplyvaluetable" }
    if yname == "cefcModuleTable" { return "Cefcmoduletable" }
    if yname == "cefcIntelliModuleTable" { return "Cefcintellimoduletable" }
    if yname == "cefcModuleLocalSwitchingTable" { return "Cefcmodulelocalswitchingtable" }
    if yname == "cefcFanTrayStatusTable" { return "Cefcfantraystatustable" }
    if yname == "cefcPhysicalTable" { return "Cefcphysicaltable" }
    if yname == "cefcPowerSupplyInputTable" { return "Cefcpowersupplyinputtable" }
    if yname == "cefcPowerSupplyOutputTable" { return "Cefcpowersupplyoutputtable" }
    if yname == "cefcChassisCoolingTable" { return "Cefcchassiscoolingtable" }
    if yname == "cefcFanCoolingTable" { return "Cefcfancoolingtable" }
    if yname == "cefcModuleCoolingTable" { return "Cefcmodulecoolingtable" }
    if yname == "cefcFanCoolingCapTable" { return "Cefcfancoolingcaptable" }
    if yname == "cefcConnectorRatingTable" { return "Cefcconnectorratingtable" }
    if yname == "cefcModulePowerConsumptionTable" { return "Cefcmodulepowerconsumptiontable" }
    return ""
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetSegmentPath() string {
    return "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB"
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFRUPower" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfrupower
    }
    if childYangName == "cefcMIBNotificationEnables" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcmibnotificationenables
    }
    if childYangName == "cefcFRUPowerSupplyGroupTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplygrouptable
    }
    if childYangName == "cefcFRUPowerStatusTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfrupowerstatustable
    }
    if childYangName == "cefcFRUPowerSupplyValueTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplyvaluetable
    }
    if childYangName == "cefcModuleTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcmoduletable
    }
    if childYangName == "cefcIntelliModuleTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcintellimoduletable
    }
    if childYangName == "cefcModuleLocalSwitchingTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcmodulelocalswitchingtable
    }
    if childYangName == "cefcFanTrayStatusTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfantraystatustable
    }
    if childYangName == "cefcPhysicalTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcphysicaltable
    }
    if childYangName == "cefcPowerSupplyInputTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyinputtable
    }
    if childYangName == "cefcPowerSupplyOutputTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyoutputtable
    }
    if childYangName == "cefcChassisCoolingTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcchassiscoolingtable
    }
    if childYangName == "cefcFanCoolingTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingtable
    }
    if childYangName == "cefcModuleCoolingTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcmodulecoolingtable
    }
    if childYangName == "cefcFanCoolingCapTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingcaptable
    }
    if childYangName == "cefcConnectorRatingTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcconnectorratingtable
    }
    if childYangName == "cefcModulePowerConsumptionTable" {
        return &cISCOENTITYFRUCONTROLMIB.Cefcmodulepowerconsumptiontable
    }
    return nil
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cefcFRUPower"] = &cISCOENTITYFRUCONTROLMIB.Cefcfrupower
    children["cefcMIBNotificationEnables"] = &cISCOENTITYFRUCONTROLMIB.Cefcmibnotificationenables
    children["cefcFRUPowerSupplyGroupTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplygrouptable
    children["cefcFRUPowerStatusTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfrupowerstatustable
    children["cefcFRUPowerSupplyValueTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplyvaluetable
    children["cefcModuleTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcmoduletable
    children["cefcIntelliModuleTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcintellimoduletable
    children["cefcModuleLocalSwitchingTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcmodulelocalswitchingtable
    children["cefcFanTrayStatusTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfantraystatustable
    children["cefcPhysicalTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcphysicaltable
    children["cefcPowerSupplyInputTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyinputtable
    children["cefcPowerSupplyOutputTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyoutputtable
    children["cefcChassisCoolingTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcchassiscoolingtable
    children["cefcFanCoolingTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingtable
    children["cefcModuleCoolingTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcmodulecoolingtable
    children["cefcFanCoolingCapTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingcaptable
    children["cefcConnectorRatingTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcconnectorratingtable
    children["cefcModulePowerConsumptionTable"] = &cISCOENTITYFRUCONTROLMIB.Cefcmodulepowerconsumptiontable
    return children
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) SetParent(parent types.Entity) { cISCOENTITYFRUCONTROLMIB.parent = parent }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetParent() types.Entity { return cISCOENTITYFRUCONTROLMIB.parent }

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupower
type CISCOENTITYFRUCONTROLMIB_Cefcfrupower struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The system will provide power to the device connecting to the FRU if the
    // device needs power, like an IP Phone. We call the providing power inline
    // power.  This MIB object controls the maximum default inline power for the
    // device connecting to the FRU in the system. If the maximum default inline
    // power of the device is greater than the maximum value reportable by this
    // object, then this object should report its maximum reportable value (12500)
    // and cefcMaxDefaultHighInLinePower must be used to report the actual maximum
    // default inline power. The type is interface{} with range: 0..12500. Units
    // are miliwatts.
    Cefcmaxdefaultinlinepower interface{}

    // The system will provide power to the device connecting to the FRU if the
    // device needs power, like an IP Phone. We call the providing power inline
    // power.  This MIB object controls the maximum default inline power for the
    // device connecting to the FRU in the system. The type is interface{} with
    // range: 0..4294967295. Units are miliwatts.
    Cefcmaxdefaulthighinlinepower interface{}
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetFilter() yfilter.YFilter { return cefcfrupower.YFilter }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) SetFilter(yf yfilter.YFilter) { cefcfrupower.YFilter = yf }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetGoName(yname string) string {
    if yname == "cefcMaxDefaultInLinePower" { return "Cefcmaxdefaultinlinepower" }
    if yname == "cefcMaxDefaultHighInLinePower" { return "Cefcmaxdefaulthighinlinepower" }
    return ""
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetSegmentPath() string {
    return "cefcFRUPower"
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefcMaxDefaultInLinePower"] = cefcfrupower.Cefcmaxdefaultinlinepower
    leafs["cefcMaxDefaultHighInLinePower"] = cefcfrupower.Cefcmaxdefaulthighinlinepower
    return leafs
}

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetYangName() string { return "cefcFRUPower" }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) SetParent(parent types.Entity) { cefcfrupower.parent = parent }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetParent() types.Entity { return cefcfrupower.parent }

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables
type CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the following
    // notifications: cefcModuleStatusChange, cefcPowerStatusChange, 
    // cefcFRUInserted, cefcFRURemoved,  cefcUnrecognizedFRU and
    // cefcFanTrayStatusChange.  A false value will prevent these notifications
    // from being generated. The type is bool.
    Cefcmibenablestatusnotification interface{}

    // This variable indicates whether the system produces the
    // cefcPowerSupplyOutputChange  notifications when the output capacity of  a
    // power supply has changed. A false value  will prevent this notification to
    // generated. The type is bool.
    Cefcenablepsoutputchangenotif interface{}
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetFilter() yfilter.YFilter { return cefcmibnotificationenables.YFilter }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) SetFilter(yf yfilter.YFilter) { cefcmibnotificationenables.YFilter = yf }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetGoName(yname string) string {
    if yname == "cefcMIBEnableStatusNotification" { return "Cefcmibenablestatusnotification" }
    if yname == "cefcEnablePSOutputChangeNotif" { return "Cefcenablepsoutputchangenotif" }
    return ""
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetSegmentPath() string {
    return "cefcMIBNotificationEnables"
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cefcMIBEnableStatusNotification"] = cefcmibnotificationenables.Cefcmibenablestatusnotification
    leafs["cefcEnablePSOutputChangeNotif"] = cefcmibnotificationenables.Cefcenablepsoutputchangenotif
    return leafs
}

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetYangName() string { return "cefcMIBNotificationEnables" }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) SetParent(parent types.Entity) { cefcmibnotificationenables.parent = parent }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetParent() types.Entity { return cefcmibnotificationenables.parent }

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable
// This table lists the redundancy mode and the
// operational status of the power supply groups
// in the system.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyGroupTable entry lists the desired redundancy mode,
    // the units of the power outputs and the  available and drawn current for the
    // power supply group.  Entries are created by the agent when a power supply
    // group is added to the entPhysicalTable. Entries are deleted by  the agent
    // at power supply group removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry.
    Cefcfrupowersupplygroupentry []CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetFilter() yfilter.YFilter { return cefcfrupowersupplygrouptable.YFilter }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) SetFilter(yf yfilter.YFilter) { cefcfrupowersupplygrouptable.YFilter = yf }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetGoName(yname string) string {
    if yname == "cefcFRUPowerSupplyGroupEntry" { return "Cefcfrupowersupplygroupentry" }
    return ""
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetSegmentPath() string {
    return "cefcFRUPowerSupplyGroupTable"
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFRUPowerSupplyGroupEntry" {
        for _, c := range cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry {
            if cefcfrupowersupplygrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry{}
        cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry = append(cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry, child)
        return &cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry[len(cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry)-1]
    }
    return nil
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry {
        children[cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry[i].GetSegmentPath()] = &cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry[i]
    }
    return children
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetYangName() string { return "cefcFRUPowerSupplyGroupTable" }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) SetParent(parent types.Entity) { cefcfrupowersupplygrouptable.parent = parent }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetParent() types.Entity { return cefcfrupowersupplygrouptable.parent }

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry
// An cefcFRUPowerSupplyGroupTable entry lists the desired
// redundancy mode, the units of the power outputs and the 
// available and drawn current for the power supply group.
// 
// Entries are created by the agent when a power supply group
// is added to the entPhysicalTable. Entries are deleted by 
// the agent at power supply group removal.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The administratively desired power supply redundancy mode. The type is
    // PowerRedundancyType.
    Cefcpowerredundancymode interface{}

    // The units of primary supply to interpret cefcTotalAvailableCurrent and
    // cefcTotalDrawnCurrent as power.  For example, one 1000-watt power supply
    // could  deliver 100 amperes at 10 volts DC.  So the value of cefcPowerUnits
    // would be 'at 10 volts DC'.  cefcPowerUnits is for display purposes only.
    // The type is string.
    Cefcpowerunits interface{}

    // Total current available for FRU usage. The type is interface{} with range:
    // -1000000000..1000000000.
    Cefctotalavailablecurrent interface{}

    // Total current drawn by powered-on FRUs. The type is interface{} with range:
    // -1000000000..1000000000.
    Cefctotaldrawncurrent interface{}

    // The power supply redundancy operational mode. The type is
    // PowerRedundancyType.
    Cefcpowerredundancyopermode interface{}

    // This object has the value of notApplicable(1) when
    // cefcPowerRedundancyOperMode of the instance does not have the value of
    // nonRedundant(4).  The other values explain the reason why 
    // cefcPowerRedundancyOperMode is nonRedundant(4), e.g.  unknown(2)           
    // the reason is not identified.  singleSupply(3)        There is only one
    // power supply                        in the group.  mismatchedSupplies(4) 
    // There are more than one power                        supplies in the
    // groups. However                        they are mismatched and can not     
    // work redundantly.  supplyError(5)         Some power supply or supplies    
    // does or do not working properly. The type is Cefcpowernonredundantreason.
    Cefcpowernonredundantreason interface{}

    // Total inline current drawn for inline operation. The type is interface{}
    // with range: -1000000000..1000000000.
    Cefctotaldrawninlinecurrent interface{}
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetFilter() yfilter.YFilter { return cefcfrupowersupplygroupentry.YFilter }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) SetFilter(yf yfilter.YFilter) { cefcfrupowersupplygroupentry.YFilter = yf }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcPowerRedundancyMode" { return "Cefcpowerredundancymode" }
    if yname == "cefcPowerUnits" { return "Cefcpowerunits" }
    if yname == "cefcTotalAvailableCurrent" { return "Cefctotalavailablecurrent" }
    if yname == "cefcTotalDrawnCurrent" { return "Cefctotaldrawncurrent" }
    if yname == "cefcPowerRedundancyOperMode" { return "Cefcpowerredundancyopermode" }
    if yname == "cefcPowerNonRedundantReason" { return "Cefcpowernonredundantreason" }
    if yname == "cefcTotalDrawnInlineCurrent" { return "Cefctotaldrawninlinecurrent" }
    return ""
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetSegmentPath() string {
    return "cefcFRUPowerSupplyGroupEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowersupplygroupentry.Entphysicalindex) + "']"
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfrupowersupplygroupentry.Entphysicalindex
    leafs["cefcPowerRedundancyMode"] = cefcfrupowersupplygroupentry.Cefcpowerredundancymode
    leafs["cefcPowerUnits"] = cefcfrupowersupplygroupentry.Cefcpowerunits
    leafs["cefcTotalAvailableCurrent"] = cefcfrupowersupplygroupentry.Cefctotalavailablecurrent
    leafs["cefcTotalDrawnCurrent"] = cefcfrupowersupplygroupentry.Cefctotaldrawncurrent
    leafs["cefcPowerRedundancyOperMode"] = cefcfrupowersupplygroupentry.Cefcpowerredundancyopermode
    leafs["cefcPowerNonRedundantReason"] = cefcfrupowersupplygroupentry.Cefcpowernonredundantreason
    leafs["cefcTotalDrawnInlineCurrent"] = cefcfrupowersupplygroupentry.Cefctotaldrawninlinecurrent
    return leafs
}

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetYangName() string { return "cefcFRUPowerSupplyGroupEntry" }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) SetParent(parent types.Entity) { cefcfrupowersupplygroupentry.parent = parent }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetParent() types.Entity { return cefcfrupowersupplygroupentry.parent }

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetParentYangName() string { return "cefcFRUPowerSupplyGroupTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason represents                        does or do not working properly.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason string

const (
    CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason_notApplicable CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason = "notApplicable"

    CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason_unknown CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason = "unknown"

    CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason_singleSupply CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason = "singleSupply"

    CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason_mismatchedSupplies CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason = "mismatchedSupplies"

    CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason_supplyError CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry_Cefcpowernonredundantreason = "supplyError"
)

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable
// This table lists the power-related administrative status
// and operational status of the manageable components
// in the system.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An cefcFRUPowerStatusTable entry lists the desired administrative status,
    // the operational status of the  power manageable component, and the current
    // required by  the component for operation.  Entries are created by the agent
    // at system power-up or  the insertion of the component.  Entries are deleted
    // by the agent at the removal of the component.  Only components with power
    // control are listed in the  table. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry.
    Cefcfrupowerstatusentry []CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetFilter() yfilter.YFilter { return cefcfrupowerstatustable.YFilter }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) SetFilter(yf yfilter.YFilter) { cefcfrupowerstatustable.YFilter = yf }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetGoName(yname string) string {
    if yname == "cefcFRUPowerStatusEntry" { return "Cefcfrupowerstatusentry" }
    return ""
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetSegmentPath() string {
    return "cefcFRUPowerStatusTable"
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFRUPowerStatusEntry" {
        for _, c := range cefcfrupowerstatustable.Cefcfrupowerstatusentry {
            if cefcfrupowerstatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry{}
        cefcfrupowerstatustable.Cefcfrupowerstatusentry = append(cefcfrupowerstatustable.Cefcfrupowerstatusentry, child)
        return &cefcfrupowerstatustable.Cefcfrupowerstatusentry[len(cefcfrupowerstatustable.Cefcfrupowerstatusentry)-1]
    }
    return nil
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfrupowerstatustable.Cefcfrupowerstatusentry {
        children[cefcfrupowerstatustable.Cefcfrupowerstatusentry[i].GetSegmentPath()] = &cefcfrupowerstatustable.Cefcfrupowerstatusentry[i]
    }
    return children
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetYangName() string { return "cefcFRUPowerStatusTable" }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) SetParent(parent types.Entity) { cefcfrupowerstatustable.parent = parent }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetParent() types.Entity { return cefcfrupowerstatustable.parent }

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry
// An cefcFRUPowerStatusTable entry lists the desired
// administrative status, the operational status of the 
// power manageable component, and the current required by 
// the component for operation.
// 
// Entries are created by the agent at system power-up or 
// the insertion of the component.  Entries are deleted by
// the agent at the removal of the component.
// 
// Only components with power control are listed in the 
// table.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // Administratively desired FRU power state. The type is PowerAdminType.
    Cefcfrupoweradminstatus interface{}

    // Operational FRU power state. The type is PowerOperType.
    Cefcfrupoweroperstatus interface{}

    // Current supplied by the FRU (positive values) or current required to
    // operate the FRU (negative values). The type is interface{} with range:
    // -1000000000..1000000000.
    Cefcfrucurrent interface{}

    // This object indicates the set of supported power capabilities of the FRU. 
    // realTimeCurrent(0) -     cefcFRURealTimeCurrent is supported by the FRU.
    // The type is map[string]bool.
    Cefcfrupowercapability interface{}

    // This object indicates the realtime value of current supplied by the FRU
    // (positive values) or the realtime value of current drawn by the FRU
    // (negative values). The type is interface{} with range:
    // -1000000000..1000000000.
    Cefcfrurealtimecurrent interface{}
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetFilter() yfilter.YFilter { return cefcfrupowerstatusentry.YFilter }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) SetFilter(yf yfilter.YFilter) { cefcfrupowerstatusentry.YFilter = yf }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcFRUPowerAdminStatus" { return "Cefcfrupoweradminstatus" }
    if yname == "cefcFRUPowerOperStatus" { return "Cefcfrupoweroperstatus" }
    if yname == "cefcFRUCurrent" { return "Cefcfrucurrent" }
    if yname == "cefcFRUPowerCapability" { return "Cefcfrupowercapability" }
    if yname == "cefcFRURealTimeCurrent" { return "Cefcfrurealtimecurrent" }
    return ""
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetSegmentPath() string {
    return "cefcFRUPowerStatusEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowerstatusentry.Entphysicalindex) + "']"
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfrupowerstatusentry.Entphysicalindex
    leafs["cefcFRUPowerAdminStatus"] = cefcfrupowerstatusentry.Cefcfrupoweradminstatus
    leafs["cefcFRUPowerOperStatus"] = cefcfrupowerstatusentry.Cefcfrupoweroperstatus
    leafs["cefcFRUCurrent"] = cefcfrupowerstatusentry.Cefcfrucurrent
    leafs["cefcFRUPowerCapability"] = cefcfrupowerstatusentry.Cefcfrupowercapability
    leafs["cefcFRURealTimeCurrent"] = cefcfrupowerstatusentry.Cefcfrurealtimecurrent
    return leafs
}

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetYangName() string { return "cefcFRUPowerStatusEntry" }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) SetParent(parent types.Entity) { cefcfrupowerstatusentry.parent = parent }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetParent() types.Entity { return cefcfrupowerstatusentry.parent }

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetParentYangName() string { return "cefcFRUPowerStatusTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable
// This table lists the power capacity of a power FRU in the
// system if it provides variable power. Power supplies usually
// provide either system or inline power. They cannot be 
// controlled by software to dictate how they distribute power.
// We can also have what are known as variable power supplies.
// They can provide both system and inline power and can be 
// varied within hardware defined ranges for system and inline
// limited by a total maximum combined output. They could be
// configured by the user via CLI or SNMP or be controlled by
// software internally.
// This table supplements the information in the
// cefcFRUPowerStatusTable for power supply FRUs. The 
// cefcFRUCurrent attribute in that table provides the overall
// current the power supply FRU can provide while this table 
// gives us the individual contribution towards system and 
// inline power.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyValueTable entry lists the current provided by the FRU
    // for operation.  Entries are created by the agent at system power-up or  FRU
    // insertion.  Entries are deleted by the agent at FRU removal.  Only power
    // supply FRUs are listed in the table. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry.
    Cefcfrupowersupplyvalueentry []CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetFilter() yfilter.YFilter { return cefcfrupowersupplyvaluetable.YFilter }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) SetFilter(yf yfilter.YFilter) { cefcfrupowersupplyvaluetable.YFilter = yf }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetGoName(yname string) string {
    if yname == "cefcFRUPowerSupplyValueEntry" { return "Cefcfrupowersupplyvalueentry" }
    return ""
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetSegmentPath() string {
    return "cefcFRUPowerSupplyValueTable"
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFRUPowerSupplyValueEntry" {
        for _, c := range cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry {
            if cefcfrupowersupplyvaluetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry{}
        cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry = append(cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry, child)
        return &cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry[len(cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry)-1]
    }
    return nil
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry {
        children[cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry[i].GetSegmentPath()] = &cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry[i]
    }
    return children
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetYangName() string { return "cefcFRUPowerSupplyValueTable" }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) SetParent(parent types.Entity) { cefcfrupowersupplyvaluetable.parent = parent }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetParent() types.Entity { return cefcfrupowersupplyvaluetable.parent }

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry
// An cefcFRUPowerSupplyValueTable entry lists the current
// provided by the FRU for operation.
// 
// Entries are created by the agent at system power-up or 
// FRU insertion.  Entries are deleted by the agent at FRU
// removal.
// 
// Only power supply FRUs are listed in the table.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // Total current that could be supplied by the FRU (positive values) for
    // system operations. The type is interface{} with range:
    // -1000000000..1000000000.
    Cefcfrutotalsystemcurrent interface{}

    // Amount of current drawn by the FRU's in the system towards system
    // operations from this FRU. The type is interface{} with range:
    // -1000000000..1000000000.
    Cefcfrudrawnsystemcurrent interface{}

    // Total current supplied by the FRU (positive values) for inline operations.
    // The type is interface{} with range: -1000000000..1000000000.
    Cefcfrutotalinlinecurrent interface{}

    // Amount of current that is being drawn from this FRU for inline operation.
    // The type is interface{} with range: -1000000000..1000000000.
    Cefcfrudrawninlinecurrent interface{}
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetFilter() yfilter.YFilter { return cefcfrupowersupplyvalueentry.YFilter }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) SetFilter(yf yfilter.YFilter) { cefcfrupowersupplyvalueentry.YFilter = yf }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcFRUTotalSystemCurrent" { return "Cefcfrutotalsystemcurrent" }
    if yname == "cefcFRUDrawnSystemCurrent" { return "Cefcfrudrawnsystemcurrent" }
    if yname == "cefcFRUTotalInlineCurrent" { return "Cefcfrutotalinlinecurrent" }
    if yname == "cefcFRUDrawnInlineCurrent" { return "Cefcfrudrawninlinecurrent" }
    return ""
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetSegmentPath() string {
    return "cefcFRUPowerSupplyValueEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowersupplyvalueentry.Entphysicalindex) + "']"
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfrupowersupplyvalueentry.Entphysicalindex
    leafs["cefcFRUTotalSystemCurrent"] = cefcfrupowersupplyvalueentry.Cefcfrutotalsystemcurrent
    leafs["cefcFRUDrawnSystemCurrent"] = cefcfrupowersupplyvalueentry.Cefcfrudrawnsystemcurrent
    leafs["cefcFRUTotalInlineCurrent"] = cefcfrupowersupplyvalueentry.Cefcfrutotalinlinecurrent
    leafs["cefcFRUDrawnInlineCurrent"] = cefcfrupowersupplyvalueentry.Cefcfrudrawninlinecurrent
    return leafs
}

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetYangName() string { return "cefcFRUPowerSupplyValueEntry" }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) SetParent(parent types.Entity) { cefcfrupowersupplyvalueentry.parent = parent }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetParent() types.Entity { return cefcfrupowersupplyvalueentry.parent }

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetParentYangName() string { return "cefcFRUPowerSupplyValueTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcmoduletable
// A cefcModuleTable entry lists the operational and
// administrative status information for ENTITY-MIB
// entPhysicalTable entries for manageable components
// of type PhysicalClass module(9).
type CISCOENTITYFRUCONTROLMIB_Cefcmoduletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcModuleStatusTable entry lists the operational and administrative
    // status information for ENTITY-MIB entPhysicalTable entries for manageable
    // components  of type PhysicalClass module(9).  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry.
    Cefcmoduleentry []CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetFilter() yfilter.YFilter { return cefcmoduletable.YFilter }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) SetFilter(yf yfilter.YFilter) { cefcmoduletable.YFilter = yf }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetGoName(yname string) string {
    if yname == "cefcModuleEntry" { return "Cefcmoduleentry" }
    return ""
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetSegmentPath() string {
    return "cefcModuleTable"
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcModuleEntry" {
        for _, c := range cefcmoduletable.Cefcmoduleentry {
            if cefcmoduletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry{}
        cefcmoduletable.Cefcmoduleentry = append(cefcmoduletable.Cefcmoduleentry, child)
        return &cefcmoduletable.Cefcmoduleentry[len(cefcmoduletable.Cefcmoduleentry)-1]
    }
    return nil
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcmoduletable.Cefcmoduleentry {
        children[cefcmoduletable.Cefcmoduleentry[i].GetSegmentPath()] = &cefcmoduletable.Cefcmoduleentry[i]
    }
    return children
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetYangName() string { return "cefcModuleTable" }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) SetParent(parent types.Entity) { cefcmoduletable.parent = parent }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetParent() types.Entity { return cefcmoduletable.parent }

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry
// A cefcModuleStatusTable entry lists the operational and
// administrative status information for ENTITY-MIB
// entPhysicalTable entries for manageable components 
// of type PhysicalClass module(9).
// 
// Entries are created by the agent at the system power-up or
// module insertion.
// 
// Entries are deleted by the agent upon module removal.
type CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This object provides administrative control of the module. The type is
    // ModuleAdminType.
    Cefcmoduleadminstatus interface{}

    // This object shows the module's operational state. The type is
    // ModuleOperType.
    Cefcmoduleoperstatus interface{}

    // This object identifies the reason for the last reset performed on the
    // module. The type is ModuleResetReasonType.
    Cefcmoduleresetreason interface{}

    // The value of sysUpTime at the time the cefcModuleOperStatus is changed. The
    // type is interface{} with range: 0..4294967295.
    Cefcmodulestatuslastchangetime interface{}

    // The value of sysUpTime when the configuration was most recently cleared.
    // The type is interface{} with range: 0..4294967295.
    Cefcmodulelastclearconfigtime interface{}

    // A description qualifying the module reset reason specified in
    // cefcModuleResetReason.   Examples:   command xyz                 missing
    // task   switch over   watchdog timeout       etc. 
    // cefcModuleResetReasonDescription is for display purposes only. NMS
    // applications must not parse. The type is string.
    Cefcmoduleresetreasondescription interface{}

    // This object displays human-readable textual string which describes the
    // cause of the last state change of the module. This object contains zero
    // length string if no meaningful reason could be provided.  Examples:
    // 'Invalid software version' 'Software download failed' 'Software version
    // mismatch' 'Module is in standby state' etc.  This object is for display
    // purposes only. NMS applications must not parse this object and take any
    // decision based on its value. The type is string.
    Cefcmodulestatechangereasondescr interface{}

    // This object provides the up time for the module since it was last
    // re-initialized.  This object is not persistent; if a module reset, restart,
    // power off, the up time starts from zero. The type is interface{} with
    // range: 0..4294967295.
    Cefcmoduleuptime interface{}
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetFilter() yfilter.YFilter { return cefcmoduleentry.YFilter }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) SetFilter(yf yfilter.YFilter) { cefcmoduleentry.YFilter = yf }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcModuleAdminStatus" { return "Cefcmoduleadminstatus" }
    if yname == "cefcModuleOperStatus" { return "Cefcmoduleoperstatus" }
    if yname == "cefcModuleResetReason" { return "Cefcmoduleresetreason" }
    if yname == "cefcModuleStatusLastChangeTime" { return "Cefcmodulestatuslastchangetime" }
    if yname == "cefcModuleLastClearConfigTime" { return "Cefcmodulelastclearconfigtime" }
    if yname == "cefcModuleResetReasonDescription" { return "Cefcmoduleresetreasondescription" }
    if yname == "cefcModuleStateChangeReasonDescr" { return "Cefcmodulestatechangereasondescr" }
    if yname == "cefcModuleUpTime" { return "Cefcmoduleuptime" }
    return ""
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetSegmentPath() string {
    return "cefcModuleEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmoduleentry.Entphysicalindex) + "']"
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcmoduleentry.Entphysicalindex
    leafs["cefcModuleAdminStatus"] = cefcmoduleentry.Cefcmoduleadminstatus
    leafs["cefcModuleOperStatus"] = cefcmoduleentry.Cefcmoduleoperstatus
    leafs["cefcModuleResetReason"] = cefcmoduleentry.Cefcmoduleresetreason
    leafs["cefcModuleStatusLastChangeTime"] = cefcmoduleentry.Cefcmodulestatuslastchangetime
    leafs["cefcModuleLastClearConfigTime"] = cefcmoduleentry.Cefcmodulelastclearconfigtime
    leafs["cefcModuleResetReasonDescription"] = cefcmoduleentry.Cefcmoduleresetreasondescription
    leafs["cefcModuleStateChangeReasonDescr"] = cefcmoduleentry.Cefcmodulestatechangereasondescr
    leafs["cefcModuleUpTime"] = cefcmoduleentry.Cefcmoduleuptime
    return leafs
}

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetYangName() string { return "cefcModuleEntry" }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) SetParent(parent types.Entity) { cefcmoduleentry.parent = parent }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetParent() types.Entity { return cefcmoduleentry.parent }

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetParentYangName() string { return "cefcModuleTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable
// This table sparsely augments the
// cefcModuleTable (i.e., every row in
// this table corresponds to a row in
// the cefcModuleTable but not necessarily
// vice-versa).
// 
// A cefcIntelliModuleTable entry lists the
// information specific to intelligent
// modules which cannot be provided by the
// cefcModuleTable.
type CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcIntelliModuleTable entry lists the information specific to an
    // intelligent module which cannot be provided by this module's corresponding
    // instance in the cefcModuleTable. Only an intelligent module with Internet
    // address configured has its entry here.  An entry of this table is created
    // if an  intelligent module is detected by the  managed system and its
    // management Internet address is configured on the intelligent  module.  An
    // entry of this table is deleted if the  removal of Internet address
    // configuration of  this module or the module itself. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry.
    Cefcintellimoduleentry []CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetFilter() yfilter.YFilter { return cefcintellimoduletable.YFilter }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) SetFilter(yf yfilter.YFilter) { cefcintellimoduletable.YFilter = yf }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetGoName(yname string) string {
    if yname == "cefcIntelliModuleEntry" { return "Cefcintellimoduleentry" }
    return ""
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetSegmentPath() string {
    return "cefcIntelliModuleTable"
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcIntelliModuleEntry" {
        for _, c := range cefcintellimoduletable.Cefcintellimoduleentry {
            if cefcintellimoduletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry{}
        cefcintellimoduletable.Cefcintellimoduleentry = append(cefcintellimoduletable.Cefcintellimoduleentry, child)
        return &cefcintellimoduletable.Cefcintellimoduleentry[len(cefcintellimoduletable.Cefcintellimoduleentry)-1]
    }
    return nil
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcintellimoduletable.Cefcintellimoduleentry {
        children[cefcintellimoduletable.Cefcintellimoduleentry[i].GetSegmentPath()] = &cefcintellimoduletable.Cefcintellimoduleentry[i]
    }
    return children
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetYangName() string { return "cefcIntelliModuleTable" }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) SetParent(parent types.Entity) { cefcintellimoduletable.parent = parent }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetParent() types.Entity { return cefcintellimoduletable.parent }

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry
// A cefcIntelliModuleTable entry lists the
// information specific to an intelligent
// module which cannot be provided by
// this module's corresponding instance in
// the cefcModuleTable. Only an intelligent
// module with Internet address configured has
// its entry here.
// 
// An entry of this table is created if an 
// intelligent module is detected by the 
// managed system and its management Internet
// address is configured on the intelligent 
// module.
// 
// An entry of this table is deleted if the 
// removal of Internet address configuration of 
// this module or the module itself.
type CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The type of Internet address by which the intelligent module is reachable.
    // The type is InetAddressType.
    Cefcintellimoduleipaddrtype interface{}

    // The Internet address configured for the intelligent module. The type of
    // this address is  determined by the value of the object 
    // cefcIntelliModuleIPAddrType. The type is string with length: 0..255.
    Cefcintellimoduleipaddr interface{}
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetFilter() yfilter.YFilter { return cefcintellimoduleentry.YFilter }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) SetFilter(yf yfilter.YFilter) { cefcintellimoduleentry.YFilter = yf }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcIntelliModuleIPAddrType" { return "Cefcintellimoduleipaddrtype" }
    if yname == "cefcIntelliModuleIPAddr" { return "Cefcintellimoduleipaddr" }
    return ""
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetSegmentPath() string {
    return "cefcIntelliModuleEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcintellimoduleentry.Entphysicalindex) + "']"
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcintellimoduleentry.Entphysicalindex
    leafs["cefcIntelliModuleIPAddrType"] = cefcintellimoduleentry.Cefcintellimoduleipaddrtype
    leafs["cefcIntelliModuleIPAddr"] = cefcintellimoduleentry.Cefcintellimoduleipaddr
    return leafs
}

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetYangName() string { return "cefcIntelliModuleEntry" }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) SetParent(parent types.Entity) { cefcintellimoduleentry.parent = parent }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetParent() types.Entity { return cefcintellimoduleentry.parent }

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetParentYangName() string { return "cefcIntelliModuleTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable
// This table sparsely augments the cefcModuleTable
// (i.e., every row in this table corresponds to a row in
// the cefcModuleTable but not necessarily vice-versa).
// 
// A cefcModuleLocalSwitchingTable entry lists the
// information specific to local switching capable
// modules which cannot be provided by the
// cefcModuleTable.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcModuleLocalSwitchingTable entry lists the information specific to a
    // local switching capable module which cannot be provided by this module's
    // corresponding instance in the cefcModuleTable. Only a module which is
    // capable of local switching has its entry here.  An entry of this table is
    // created if a module which is capable of local switching is detected by the
    // managed system.  An entry of this table is deleted if the removal of this
    // module. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry.
    Cefcmodulelocalswitchingentry []CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetFilter() yfilter.YFilter { return cefcmodulelocalswitchingtable.YFilter }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) SetFilter(yf yfilter.YFilter) { cefcmodulelocalswitchingtable.YFilter = yf }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetGoName(yname string) string {
    if yname == "cefcModuleLocalSwitchingEntry" { return "Cefcmodulelocalswitchingentry" }
    return ""
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetSegmentPath() string {
    return "cefcModuleLocalSwitchingTable"
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcModuleLocalSwitchingEntry" {
        for _, c := range cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry {
            if cefcmodulelocalswitchingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry{}
        cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry = append(cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry, child)
        return &cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry[len(cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry)-1]
    }
    return nil
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry {
        children[cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry[i].GetSegmentPath()] = &cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry[i]
    }
    return children
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetYangName() string { return "cefcModuleLocalSwitchingTable" }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) SetParent(parent types.Entity) { cefcmodulelocalswitchingtable.parent = parent }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetParent() types.Entity { return cefcmodulelocalswitchingtable.parent }

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry
// A cefcModuleLocalSwitchingTable entry lists the
// information specific to a local switching capable
// module which cannot be provided by this module's
// corresponding instance in the cefcModuleTable.
// Only a module which is capable of local switching
// has its entry here.
// 
// An entry of this table is created if a module which
// is capable of local switching is detected by the
// managed system.
// 
// An entry of this table is deleted if the
// removal of this module.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This object specifies the mode of local switching.  enabled(1)  - local
    // switching is enabled. disabled(2) - local switching is disabled. The type
    // is Cefcmodulelocalswitchingmode.
    Cefcmodulelocalswitchingmode interface{}
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetFilter() yfilter.YFilter { return cefcmodulelocalswitchingentry.YFilter }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) SetFilter(yf yfilter.YFilter) { cefcmodulelocalswitchingentry.YFilter = yf }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcModuleLocalSwitchingMode" { return "Cefcmodulelocalswitchingmode" }
    return ""
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetSegmentPath() string {
    return "cefcModuleLocalSwitchingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulelocalswitchingentry.Entphysicalindex) + "']"
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcmodulelocalswitchingentry.Entphysicalindex
    leafs["cefcModuleLocalSwitchingMode"] = cefcmodulelocalswitchingentry.Cefcmodulelocalswitchingmode
    return leafs
}

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetYangName() string { return "cefcModuleLocalSwitchingEntry" }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) SetParent(parent types.Entity) { cefcmodulelocalswitchingentry.parent = parent }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetParent() types.Entity { return cefcmodulelocalswitchingentry.parent }

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetParentYangName() string { return "cefcModuleLocalSwitchingTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode represents disabled(2) - local switching is disabled.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode string

const (
    CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode_enabled CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode = "enabled"

    CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode_disabled CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry_Cefcmodulelocalswitchingmode = "disabled"
)

// CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable
// This table contains the operational status information
// for all ENTITY-MIB entPhysicalTable entries which have 
// an entPhysicalClass of 'fan'; specifically, all  
// entPhysicalTable entries which represent either: one 
// physical fan, or a single physical 'fan tray' which is a
// manufactured (inseparable in the field) combination of 
// multiple fans.
type CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An cefcFanTrayStatusTable entry lists the operational status information
    // for the ENTITY-MIB entPhysicalTable  entry which is identified by the value
    // of entPhysicalIndex. The value of entPhysicalClass for the identified entry
    // will be 'fan', and the represented physical entity will be  either: one
    // physical fan, or a single physical 'fan tray'  which is a manufactured
    // (inseparable in the field)  combination of multiple fans.  Entries are
    // created by the agent at system power-up or  fan or fan tray insertion. 
    // Entries are deleted  by the agent at the fan or fan tray removal. The type
    // is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry.
    Cefcfantraystatusentry []CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetFilter() yfilter.YFilter { return cefcfantraystatustable.YFilter }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) SetFilter(yf yfilter.YFilter) { cefcfantraystatustable.YFilter = yf }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetGoName(yname string) string {
    if yname == "cefcFanTrayStatusEntry" { return "Cefcfantraystatusentry" }
    return ""
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetSegmentPath() string {
    return "cefcFanTrayStatusTable"
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFanTrayStatusEntry" {
        for _, c := range cefcfantraystatustable.Cefcfantraystatusentry {
            if cefcfantraystatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry{}
        cefcfantraystatustable.Cefcfantraystatusentry = append(cefcfantraystatustable.Cefcfantraystatusentry, child)
        return &cefcfantraystatustable.Cefcfantraystatusentry[len(cefcfantraystatustable.Cefcfantraystatusentry)-1]
    }
    return nil
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfantraystatustable.Cefcfantraystatusentry {
        children[cefcfantraystatustable.Cefcfantraystatusentry[i].GetSegmentPath()] = &cefcfantraystatustable.Cefcfantraystatusentry[i]
    }
    return children
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetYangName() string { return "cefcFanTrayStatusTable" }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) SetParent(parent types.Entity) { cefcfantraystatustable.parent = parent }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetParent() types.Entity { return cefcfantraystatustable.parent }

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry
// An cefcFanTrayStatusTable entry lists the operational
// status information for the ENTITY-MIB entPhysicalTable 
// entry which is identified by the value of entPhysicalIndex.
// The value of entPhysicalClass for the identified entry will
// be 'fan', and the represented physical entity will be 
// either: one physical fan, or a single physical 'fan tray' 
// which is a manufactured (inseparable in the field) 
// combination of multiple fans.
// 
// Entries are created by the agent at system power-up or 
// fan or fan tray insertion.  Entries are deleted 
// by the agent at the fan or fan tray removal.
type CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The operational state of the fan or fan tray. unknown(1) - unknown. up(2) -
    // powered on. down(3) - powered down. warning(4) - partial failure, needs
    // replacement               as soon as possible. The type is
    // Cefcfantrayoperstatus.
    Cefcfantrayoperstatus interface{}
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetFilter() yfilter.YFilter { return cefcfantraystatusentry.YFilter }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) SetFilter(yf yfilter.YFilter) { cefcfantraystatusentry.YFilter = yf }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcFanTrayOperStatus" { return "Cefcfantrayoperstatus" }
    return ""
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetSegmentPath() string {
    return "cefcFanTrayStatusEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfantraystatusentry.Entphysicalindex) + "']"
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfantraystatusentry.Entphysicalindex
    leafs["cefcFanTrayOperStatus"] = cefcfantraystatusentry.Cefcfantrayoperstatus
    return leafs
}

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetYangName() string { return "cefcFanTrayStatusEntry" }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) SetParent(parent types.Entity) { cefcfantraystatusentry.parent = parent }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetParent() types.Entity { return cefcfantraystatusentry.parent }

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetParentYangName() string { return "cefcFanTrayStatusTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus represents              as soon as possible.
type CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus string

const (
    CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus_unknown CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus = "unknown"

    CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus_up CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus = "up"

    CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus_down CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus = "down"

    CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus_warning CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry_Cefcfantrayoperstatus = "warning"
)

// CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable
// This table contains one row per physical entity.
type CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular physical entity. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry.
    Cefcphysicalentry []CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetFilter() yfilter.YFilter { return cefcphysicaltable.YFilter }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) SetFilter(yf yfilter.YFilter) { cefcphysicaltable.YFilter = yf }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetGoName(yname string) string {
    if yname == "cefcPhysicalEntry" { return "Cefcphysicalentry" }
    return ""
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetSegmentPath() string {
    return "cefcPhysicalTable"
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcPhysicalEntry" {
        for _, c := range cefcphysicaltable.Cefcphysicalentry {
            if cefcphysicaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry{}
        cefcphysicaltable.Cefcphysicalentry = append(cefcphysicaltable.Cefcphysicalentry, child)
        return &cefcphysicaltable.Cefcphysicalentry[len(cefcphysicaltable.Cefcphysicalentry)-1]
    }
    return nil
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcphysicaltable.Cefcphysicalentry {
        children[cefcphysicaltable.Cefcphysicalentry[i].GetSegmentPath()] = &cefcphysicaltable.Cefcphysicalentry[i]
    }
    return children
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetYangName() string { return "cefcPhysicalTable" }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) SetParent(parent types.Entity) { cefcphysicaltable.parent = parent }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetParent() types.Entity { return cefcphysicaltable.parent }

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry
// Information about a particular physical entity.
type CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The status of this physical entity. other(1) - the status is not any of the
    // listed below. supported(2) - this entity is supported. unsupported(3) -
    // this entity is unsupported. incompatible(4) - this entity is incompatible.
    // It would be unsupported(3), if the ID read from Serial EPROM is not
    // supported. It would be incompatible(4), if in the present configuration
    // this FRU is not supported. The type is Cefcphysicalstatus.
    Cefcphysicalstatus interface{}
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetFilter() yfilter.YFilter { return cefcphysicalentry.YFilter }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) SetFilter(yf yfilter.YFilter) { cefcphysicalentry.YFilter = yf }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcPhysicalStatus" { return "Cefcphysicalstatus" }
    return ""
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetSegmentPath() string {
    return "cefcPhysicalEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcphysicalentry.Entphysicalindex) + "']"
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcphysicalentry.Entphysicalindex
    leafs["cefcPhysicalStatus"] = cefcphysicalentry.Cefcphysicalstatus
    return leafs
}

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetYangName() string { return "cefcPhysicalEntry" }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) SetParent(parent types.Entity) { cefcphysicalentry.parent = parent }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetParent() types.Entity { return cefcphysicalentry.parent }

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetParentYangName() string { return "cefcPhysicalTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus represents in the present configuration this FRU is not supported.
type CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus string

const (
    CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus_other CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus = "other"

    CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus_supported CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus = "supported"

    CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus_unsupported CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus = "unsupported"

    CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus_incompatible CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry_Cefcphysicalstatus = "incompatible"
)

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable
// This table contains the power input information
// for all the power supplies that have entPhysicalTable
// entries with 'powerSupply' in the entPhysicalClass. 
// 
// The entries are created by the agent at the system
// power-up or power supply insertion.
// 
// Entries are deleted by the agent upon power supply
// removal.
// 
// The number of entries is determined by the number of
// power supplies and number of power inputs on the power 
// supply.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing power input management information applicable to a
    // particular power supply and input. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry.
    Cefcpowersupplyinputentry []CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetFilter() yfilter.YFilter { return cefcpowersupplyinputtable.YFilter }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) SetFilter(yf yfilter.YFilter) { cefcpowersupplyinputtable.YFilter = yf }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetGoName(yname string) string {
    if yname == "cefcPowerSupplyInputEntry" { return "Cefcpowersupplyinputentry" }
    return ""
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetSegmentPath() string {
    return "cefcPowerSupplyInputTable"
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcPowerSupplyInputEntry" {
        for _, c := range cefcpowersupplyinputtable.Cefcpowersupplyinputentry {
            if cefcpowersupplyinputtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry{}
        cefcpowersupplyinputtable.Cefcpowersupplyinputentry = append(cefcpowersupplyinputtable.Cefcpowersupplyinputentry, child)
        return &cefcpowersupplyinputtable.Cefcpowersupplyinputentry[len(cefcpowersupplyinputtable.Cefcpowersupplyinputentry)-1]
    }
    return nil
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcpowersupplyinputtable.Cefcpowersupplyinputentry {
        children[cefcpowersupplyinputtable.Cefcpowersupplyinputentry[i].GetSegmentPath()] = &cefcpowersupplyinputtable.Cefcpowersupplyinputentry[i]
    }
    return children
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetYangName() string { return "cefcPowerSupplyInputTable" }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) SetParent(parent types.Entity) { cefcpowersupplyinputtable.parent = parent }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetParent() types.Entity { return cefcpowersupplyinputtable.parent }

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry
// An entry containing power input management information
// applicable to a particular power supply and input.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. A unique value, greater than zero, for each input
    // on a power supply. The type is interface{} with range: 0..4294967295.
    Cefcpowersupplyinputindex interface{}

    // The type of an input power detected on the power supply.  unknown(1): No
    // input power is detected.  acLow(2): Lower rating AC input power is
    // detected.  acHigh(3): Higher rating AC input power is detected.  dcLow(4):
    // Lower rating DC input power is detected.  dcHigh(5): Higher rating DC input
    // power is detected. The type is Cefcpowersupplyinputtype.
    Cefcpowersupplyinputtype interface{}
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetFilter() yfilter.YFilter { return cefcpowersupplyinputentry.YFilter }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) SetFilter(yf yfilter.YFilter) { cefcpowersupplyinputentry.YFilter = yf }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcPowerSupplyInputIndex" { return "Cefcpowersupplyinputindex" }
    if yname == "cefcPowerSupplyInputType" { return "Cefcpowersupplyinputtype" }
    return ""
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetSegmentPath() string {
    return "cefcPowerSupplyInputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcpowersupplyinputentry.Entphysicalindex) + "']" + "[cefcPowerSupplyInputIndex='" + fmt.Sprintf("%v", cefcpowersupplyinputentry.Cefcpowersupplyinputindex) + "']"
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcpowersupplyinputentry.Entphysicalindex
    leafs["cefcPowerSupplyInputIndex"] = cefcpowersupplyinputentry.Cefcpowersupplyinputindex
    leafs["cefcPowerSupplyInputType"] = cefcpowersupplyinputentry.Cefcpowersupplyinputtype
    return leafs
}

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetYangName() string { return "cefcPowerSupplyInputEntry" }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) SetParent(parent types.Entity) { cefcpowersupplyinputentry.parent = parent }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetParent() types.Entity { return cefcpowersupplyinputentry.parent }

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetParentYangName() string { return "cefcPowerSupplyInputTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype represents dcHigh(5): Higher rating DC input power is detected.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype string

const (
    CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype_unknown CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype = "unknown"

    CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype_acLow CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype = "acLow"

    CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype_acHigh CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype = "acHigh"

    CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype_dcLow CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype = "dcLow"

    CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype_dcHigh CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry_Cefcpowersupplyinputtype = "dcHigh"
)

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable
// This table contains a list of possible output
// mode for the power supplies, whose ENTITY-MIB
// entPhysicalTable entries have an entPhysicalClass
// of 'powerSupply'. It also indicate which mode
// is the operational mode within the system.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcPowerSupplyOutputTable entry lists the power output capacity and its
    // operational status for manageable components of type PhysicalClass
    // 'powerSupply'.  Entries are created by the agent at the system power-up or
    // power supply insertion.  Entries are deleted by the agent upon power supply
    // removal.  The number of entries of a power supply is determined by the
    // power supply. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry.
    Cefcpowersupplyoutputentry []CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetFilter() yfilter.YFilter { return cefcpowersupplyoutputtable.YFilter }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) SetFilter(yf yfilter.YFilter) { cefcpowersupplyoutputtable.YFilter = yf }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetGoName(yname string) string {
    if yname == "cefcPowerSupplyOutputEntry" { return "Cefcpowersupplyoutputentry" }
    return ""
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetSegmentPath() string {
    return "cefcPowerSupplyOutputTable"
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcPowerSupplyOutputEntry" {
        for _, c := range cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry {
            if cefcpowersupplyoutputtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry{}
        cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry = append(cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry, child)
        return &cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry[len(cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry)-1]
    }
    return nil
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry {
        children[cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry[i].GetSegmentPath()] = &cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry[i]
    }
    return children
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetYangName() string { return "cefcPowerSupplyOutputTable" }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) SetParent(parent types.Entity) { cefcpowersupplyoutputtable.parent = parent }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetParent() types.Entity { return cefcpowersupplyoutputtable.parent }

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry
// A cefcPowerSupplyOutputTable entry lists the
// power output capacity and its operational status
// for manageable components of type PhysicalClass
// 'powerSupply'.
// 
// Entries are created by the agent at the system
// power-up or power supply insertion.
// 
// Entries are deleted by the agent upon power supply
// removal.
// 
// The number of entries of a power supply is determined
// by the power supply.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. A unique value, greater than zero, for each
    // possible output mode on a power supply. The type is interface{} with range:
    // 0..4294967295.
    Cefcpsoutputmodeindex interface{}

    // The output capacity of the power supply. The type is interface{} with
    // range: -1000000000..1000000000.
    Cefcpsoutputmodecurrent interface{}

    // A value of 'true' indicates that this mode is the operational mode of the
    // power supply output capacity.  A value of 'false' indicates that this mode
    // is not the operational mode of the power supply output capacity.  For a
    // given power supply's entPhysicalIndex,  at most one instance of this object
    // can have the value of true(1). The type is bool.
    Cefcpsoutputmodeinoperation interface{}
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetFilter() yfilter.YFilter { return cefcpowersupplyoutputentry.YFilter }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) SetFilter(yf yfilter.YFilter) { cefcpowersupplyoutputentry.YFilter = yf }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcPSOutputModeIndex" { return "Cefcpsoutputmodeindex" }
    if yname == "cefcPSOutputModeCurrent" { return "Cefcpsoutputmodecurrent" }
    if yname == "cefcPSOutputModeInOperation" { return "Cefcpsoutputmodeinoperation" }
    return ""
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetSegmentPath() string {
    return "cefcPowerSupplyOutputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcpowersupplyoutputentry.Entphysicalindex) + "']" + "[cefcPSOutputModeIndex='" + fmt.Sprintf("%v", cefcpowersupplyoutputentry.Cefcpsoutputmodeindex) + "']"
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcpowersupplyoutputentry.Entphysicalindex
    leafs["cefcPSOutputModeIndex"] = cefcpowersupplyoutputentry.Cefcpsoutputmodeindex
    leafs["cefcPSOutputModeCurrent"] = cefcpowersupplyoutputentry.Cefcpsoutputmodecurrent
    leafs["cefcPSOutputModeInOperation"] = cefcpowersupplyoutputentry.Cefcpsoutputmodeinoperation
    return leafs
}

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetYangName() string { return "cefcPowerSupplyOutputEntry" }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) SetParent(parent types.Entity) { cefcpowersupplyoutputentry.parent = parent }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetParent() types.Entity { return cefcpowersupplyoutputentry.parent }

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetParentYangName() string { return "cefcPowerSupplyOutputTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable
// This table contains the cooling capacity
// information of the chassis whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'chassis'.
type CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcChassisCoolingEntry lists the maximum cooling capacity that could be
    // provided  for one slot on the manageable components of type PhysicalClass
    // 'chassis'.  Entries are created by the agent if the corresponding entry is
    // created in ENTITY-MIB entPhysicalTable.  Entries are deleted by the agent
    // if the corresponding entry is deleted in ENTITY-MIB entPhysicalTable. The
    // type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry.
    Cefcchassiscoolingentry []CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetFilter() yfilter.YFilter { return cefcchassiscoolingtable.YFilter }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) SetFilter(yf yfilter.YFilter) { cefcchassiscoolingtable.YFilter = yf }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetGoName(yname string) string {
    if yname == "cefcChassisCoolingEntry" { return "Cefcchassiscoolingentry" }
    return ""
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetSegmentPath() string {
    return "cefcChassisCoolingTable"
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcChassisCoolingEntry" {
        for _, c := range cefcchassiscoolingtable.Cefcchassiscoolingentry {
            if cefcchassiscoolingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry{}
        cefcchassiscoolingtable.Cefcchassiscoolingentry = append(cefcchassiscoolingtable.Cefcchassiscoolingentry, child)
        return &cefcchassiscoolingtable.Cefcchassiscoolingentry[len(cefcchassiscoolingtable.Cefcchassiscoolingentry)-1]
    }
    return nil
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcchassiscoolingtable.Cefcchassiscoolingentry {
        children[cefcchassiscoolingtable.Cefcchassiscoolingentry[i].GetSegmentPath()] = &cefcchassiscoolingtable.Cefcchassiscoolingentry[i]
    }
    return children
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetYangName() string { return "cefcChassisCoolingTable" }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) SetParent(parent types.Entity) { cefcchassiscoolingtable.parent = parent }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetParent() types.Entity { return cefcchassiscoolingtable.parent }

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry
// A cefcChassisCoolingEntry lists the maximum
// cooling capacity that could be provided 
// for one slot on the manageable components of type
// PhysicalClass 'chassis'.
// 
// Entries are created by the agent if the corresponding
// entry is created in ENTITY-MIB entPhysicalTable.
// 
// Entries are deleted by the agent if the corresponding
// entry is deleted in ENTITY-MIB entPhysicalTable.
type CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The maximum cooling capacity that could be provided for any slot in this
    // chassis.  The default unit of the cooling capacity is 'cfm', if
    // cefcChassisPerSlotCoolingUnit is not supported. The type is interface{}
    // with range: 0..4294967295.
    Cefcchassisperslotcoolingcap interface{}

    // The unit of the maximum cooling capacity for any slot in this chassis. The
    // type is FRUCoolingUnit.
    Cefcchassisperslotcoolingunit interface{}
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetFilter() yfilter.YFilter { return cefcchassiscoolingentry.YFilter }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) SetFilter(yf yfilter.YFilter) { cefcchassiscoolingentry.YFilter = yf }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcChassisPerSlotCoolingCap" { return "Cefcchassisperslotcoolingcap" }
    if yname == "cefcChassisPerSlotCoolingUnit" { return "Cefcchassisperslotcoolingunit" }
    return ""
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetSegmentPath() string {
    return "cefcChassisCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcchassiscoolingentry.Entphysicalindex) + "']"
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcchassiscoolingentry.Entphysicalindex
    leafs["cefcChassisPerSlotCoolingCap"] = cefcchassiscoolingentry.Cefcchassisperslotcoolingcap
    leafs["cefcChassisPerSlotCoolingUnit"] = cefcchassiscoolingentry.Cefcchassisperslotcoolingunit
    return leafs
}

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetYangName() string { return "cefcChassisCoolingEntry" }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) SetParent(parent types.Entity) { cefcchassiscoolingentry.parent = parent }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetParent() types.Entity { return cefcchassiscoolingentry.parent }

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetParentYangName() string { return "cefcChassisCoolingTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable
// This table contains the cooling capacity
// information of the fans whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcFanCoolingEntry lists the cooling capacity that is provided by the 
    // manageable components of type PhysicalClass  'fan'.  Entries are created by
    // the agent if the corresponding entry is created in ENTITY-MIB
    // entPhysicalTable.  Entries are deleted by the agent if the corresponding
    // entry is deleted in ENTITY-MIB entPhysicalTable. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry.
    Cefcfancoolingentry []CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetFilter() yfilter.YFilter { return cefcfancoolingtable.YFilter }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) SetFilter(yf yfilter.YFilter) { cefcfancoolingtable.YFilter = yf }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetGoName(yname string) string {
    if yname == "cefcFanCoolingEntry" { return "Cefcfancoolingentry" }
    return ""
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetSegmentPath() string {
    return "cefcFanCoolingTable"
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFanCoolingEntry" {
        for _, c := range cefcfancoolingtable.Cefcfancoolingentry {
            if cefcfancoolingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry{}
        cefcfancoolingtable.Cefcfancoolingentry = append(cefcfancoolingtable.Cefcfancoolingentry, child)
        return &cefcfancoolingtable.Cefcfancoolingentry[len(cefcfancoolingtable.Cefcfancoolingentry)-1]
    }
    return nil
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfancoolingtable.Cefcfancoolingentry {
        children[cefcfancoolingtable.Cefcfancoolingentry[i].GetSegmentPath()] = &cefcfancoolingtable.Cefcfancoolingentry[i]
    }
    return children
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetYangName() string { return "cefcFanCoolingTable" }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) SetParent(parent types.Entity) { cefcfancoolingtable.parent = parent }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetParent() types.Entity { return cefcfancoolingtable.parent }

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry
// A cefcFanCoolingEntry lists the cooling
// capacity that is provided by the 
// manageable components of type PhysicalClass 
// 'fan'.
// 
// Entries are created by the agent if the corresponding
// entry is created in ENTITY-MIB entPhysicalTable.
// 
// Entries are deleted by the agent if the corresponding
// entry is deleted in ENTITY-MIB entPhysicalTable.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The cooling capacity that is provided by this fan.  The default unit of the
    // fan cooling capacity is 'cfm', if cefcFanCoolingCapacityUnit is not
    // supported. The type is interface{} with range: 0..4294967295.
    Cefcfancoolingcapacity interface{}

    // The unit of the fan cooling capacity. The type is FRUCoolingUnit.
    Cefcfancoolingcapacityunit interface{}
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetFilter() yfilter.YFilter { return cefcfancoolingentry.YFilter }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) SetFilter(yf yfilter.YFilter) { cefcfancoolingentry.YFilter = yf }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcFanCoolingCapacity" { return "Cefcfancoolingcapacity" }
    if yname == "cefcFanCoolingCapacityUnit" { return "Cefcfancoolingcapacityunit" }
    return ""
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetSegmentPath() string {
    return "cefcFanCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfancoolingentry.Entphysicalindex) + "']"
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfancoolingentry.Entphysicalindex
    leafs["cefcFanCoolingCapacity"] = cefcfancoolingentry.Cefcfancoolingcapacity
    leafs["cefcFanCoolingCapacityUnit"] = cefcfancoolingentry.Cefcfancoolingcapacityunit
    return leafs
}

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetYangName() string { return "cefcFanCoolingEntry" }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) SetParent(parent types.Entity) { cefcfancoolingentry.parent = parent }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetParent() types.Entity { return cefcfancoolingentry.parent }

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetParentYangName() string { return "cefcFanCoolingTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable
// This table contains the cooling requirement for
// all the manageable components of type entPhysicalClass
// 'module'.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcModuleCoolingEntry lists the cooling requirement for a manageable
    // components of type entPhysicalClass 'module'.  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry.
    Cefcmodulecoolingentry []CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetFilter() yfilter.YFilter { return cefcmodulecoolingtable.YFilter }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) SetFilter(yf yfilter.YFilter) { cefcmodulecoolingtable.YFilter = yf }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetGoName(yname string) string {
    if yname == "cefcModuleCoolingEntry" { return "Cefcmodulecoolingentry" }
    return ""
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetSegmentPath() string {
    return "cefcModuleCoolingTable"
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcModuleCoolingEntry" {
        for _, c := range cefcmodulecoolingtable.Cefcmodulecoolingentry {
            if cefcmodulecoolingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry{}
        cefcmodulecoolingtable.Cefcmodulecoolingentry = append(cefcmodulecoolingtable.Cefcmodulecoolingentry, child)
        return &cefcmodulecoolingtable.Cefcmodulecoolingentry[len(cefcmodulecoolingtable.Cefcmodulecoolingentry)-1]
    }
    return nil
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcmodulecoolingtable.Cefcmodulecoolingentry {
        children[cefcmodulecoolingtable.Cefcmodulecoolingentry[i].GetSegmentPath()] = &cefcmodulecoolingtable.Cefcmodulecoolingentry[i]
    }
    return children
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetYangName() string { return "cefcModuleCoolingTable" }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) SetParent(parent types.Entity) { cefcmodulecoolingtable.parent = parent }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetParent() types.Entity { return cefcmodulecoolingtable.parent }

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry
// A cefcModuleCoolingEntry lists the cooling
// requirement for a manageable components of type
// entPhysicalClass 'module'.
// 
// Entries are created by the agent at the system
// power-up or module insertion.
// 
// Entries are deleted by the agent upon module
// removal.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The cooling requirement of the module and its daughter cards.  The default
    // unit of the module cooling requirement is 'cfm', if cefcModuleCoolingUnit
    // is not supported. The type is interface{} with range: 0..4294967295.
    Cefcmodulecooling interface{}

    // The unit of the cooling requirement of the module and its daughter cards.
    // The type is FRUCoolingUnit.
    Cefcmodulecoolingunit interface{}
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetFilter() yfilter.YFilter { return cefcmodulecoolingentry.YFilter }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) SetFilter(yf yfilter.YFilter) { cefcmodulecoolingentry.YFilter = yf }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcModuleCooling" { return "Cefcmodulecooling" }
    if yname == "cefcModuleCoolingUnit" { return "Cefcmodulecoolingunit" }
    return ""
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetSegmentPath() string {
    return "cefcModuleCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulecoolingentry.Entphysicalindex) + "']"
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcmodulecoolingentry.Entphysicalindex
    leafs["cefcModuleCooling"] = cefcmodulecoolingentry.Cefcmodulecooling
    leafs["cefcModuleCoolingUnit"] = cefcmodulecoolingentry.Cefcmodulecoolingunit
    return leafs
}

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetYangName() string { return "cefcModuleCoolingEntry" }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) SetParent(parent types.Entity) { cefcmodulecoolingentry.parent = parent }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetParent() types.Entity { return cefcmodulecoolingentry.parent }

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetParentYangName() string { return "cefcModuleCoolingTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable
// This table contains a list of the possible cooling
// capacity modes and properties of the fans, whose 
// ENTITY-MIB entPhysicalTable entries have an 
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcFanCoolingCapacityEntry lists the cooling capacity mode of a
    // manageable components of type entPhysicalClass 'fan'. It also lists the
    // corresponding cooling capacity provided and the power consumed by the fan
    // on this mode.   Entries are created by the agent if the corresponding entry
    // is created in ENTITY-MIB entPhysicalTable.  Entries are deleted by the
    // agent if the corresponding entry is deleted in ENTITY-MIB entPhysicalTable.
    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry.
    Cefcfancoolingcapentry []CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetFilter() yfilter.YFilter { return cefcfancoolingcaptable.YFilter }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) SetFilter(yf yfilter.YFilter) { cefcfancoolingcaptable.YFilter = yf }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetGoName(yname string) string {
    if yname == "cefcFanCoolingCapEntry" { return "Cefcfancoolingcapentry" }
    return ""
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetSegmentPath() string {
    return "cefcFanCoolingCapTable"
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcFanCoolingCapEntry" {
        for _, c := range cefcfancoolingcaptable.Cefcfancoolingcapentry {
            if cefcfancoolingcaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry{}
        cefcfancoolingcaptable.Cefcfancoolingcapentry = append(cefcfancoolingcaptable.Cefcfancoolingcapentry, child)
        return &cefcfancoolingcaptable.Cefcfancoolingcapentry[len(cefcfancoolingcaptable.Cefcfancoolingcapentry)-1]
    }
    return nil
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcfancoolingcaptable.Cefcfancoolingcapentry {
        children[cefcfancoolingcaptable.Cefcfancoolingcapentry[i].GetSegmentPath()] = &cefcfancoolingcaptable.Cefcfancoolingcapentry[i]
    }
    return children
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetYangName() string { return "cefcFanCoolingCapTable" }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) SetParent(parent types.Entity) { cefcfancoolingcaptable.parent = parent }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetParent() types.Entity { return cefcfancoolingcaptable.parent }

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry
// A cefcFanCoolingCapacityEntry lists the cooling
// capacity mode of a manageable components of type
// entPhysicalClass 'fan'. It also lists the corresponding
// cooling capacity provided and the power consumed
// by the fan on this mode.
// 
// 
// Entries are created by the agent if the corresponding
// entry is created in ENTITY-MIB entPhysicalTable.
// 
// Entries are deleted by the agent if the corresponding
// entry is deleted in ENTITY-MIB entPhysicalTable.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. An arbitrary value that uniquely identifies a
    // cooling capacity mode for a fan. The type is interface{} with range:
    // 1..4095.
    Cefcfancoolingcapindex interface{}

    // A textual description of the cooling capacity mode of the fan. The type is
    // string.
    Cefcfancoolingcapmodedescr interface{}

    // The cooling capacity that could be provided when the fan is operating in
    // this mode.  The default unit of the cooling capacity is 'cfm', if
    // cefcFanCoolingCapCapacityUnit is not supported. The type is interface{}
    // with range: 0..4294967295.
    Cefcfancoolingcapcapacity interface{}

    // The power consumption of the fan when operating in in this mode. The type
    // is interface{} with range: -1000000000..1000000000.
    Cefcfancoolingcapcurrent interface{}

    // The unit of the fan cooling capacity when operating in this mode. The type
    // is FRUCoolingUnit.
    Cefcfancoolingcapcapacityunit interface{}
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetFilter() yfilter.YFilter { return cefcfancoolingcapentry.YFilter }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) SetFilter(yf yfilter.YFilter) { cefcfancoolingcapentry.YFilter = yf }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcFanCoolingCapIndex" { return "Cefcfancoolingcapindex" }
    if yname == "cefcFanCoolingCapModeDescr" { return "Cefcfancoolingcapmodedescr" }
    if yname == "cefcFanCoolingCapCapacity" { return "Cefcfancoolingcapcapacity" }
    if yname == "cefcFanCoolingCapCurrent" { return "Cefcfancoolingcapcurrent" }
    if yname == "cefcFanCoolingCapCapacityUnit" { return "Cefcfancoolingcapcapacityunit" }
    return ""
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetSegmentPath() string {
    return "cefcFanCoolingCapEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfancoolingcapentry.Entphysicalindex) + "']" + "[cefcFanCoolingCapIndex='" + fmt.Sprintf("%v", cefcfancoolingcapentry.Cefcfancoolingcapindex) + "']"
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcfancoolingcapentry.Entphysicalindex
    leafs["cefcFanCoolingCapIndex"] = cefcfancoolingcapentry.Cefcfancoolingcapindex
    leafs["cefcFanCoolingCapModeDescr"] = cefcfancoolingcapentry.Cefcfancoolingcapmodedescr
    leafs["cefcFanCoolingCapCapacity"] = cefcfancoolingcapentry.Cefcfancoolingcapcapacity
    leafs["cefcFanCoolingCapCurrent"] = cefcfancoolingcapentry.Cefcfancoolingcapcurrent
    leafs["cefcFanCoolingCapCapacityUnit"] = cefcfancoolingcapentry.Cefcfancoolingcapcapacityunit
    return leafs
}

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetYangName() string { return "cefcFanCoolingCapEntry" }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) SetParent(parent types.Entity) { cefcfancoolingcapentry.parent = parent }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetParent() types.Entity { return cefcfancoolingcapentry.parent }

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetParentYangName() string { return "cefcFanCoolingCapTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable
// This table contains the connector power
// ratings of FRUs. 
// 
// Only components with power connector rating 
// management are listed in this table.
type CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcConnectorRatingEntry lists the power connector rating information of
    // a  component in the system.  An entry or entries are created by the agent
    // when an physical entity with connector rating  management is added to the
    // ENTITY-MIB  entPhysicalTable. An entry is deleted  by the agent at the
    // entity removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry.
    Cefcconnectorratingentry []CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetFilter() yfilter.YFilter { return cefcconnectorratingtable.YFilter }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) SetFilter(yf yfilter.YFilter) { cefcconnectorratingtable.YFilter = yf }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetGoName(yname string) string {
    if yname == "cefcConnectorRatingEntry" { return "Cefcconnectorratingentry" }
    return ""
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetSegmentPath() string {
    return "cefcConnectorRatingTable"
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcConnectorRatingEntry" {
        for _, c := range cefcconnectorratingtable.Cefcconnectorratingentry {
            if cefcconnectorratingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry{}
        cefcconnectorratingtable.Cefcconnectorratingentry = append(cefcconnectorratingtable.Cefcconnectorratingentry, child)
        return &cefcconnectorratingtable.Cefcconnectorratingentry[len(cefcconnectorratingtable.Cefcconnectorratingentry)-1]
    }
    return nil
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcconnectorratingtable.Cefcconnectorratingentry {
        children[cefcconnectorratingtable.Cefcconnectorratingentry[i].GetSegmentPath()] = &cefcconnectorratingtable.Cefcconnectorratingentry[i]
    }
    return children
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetYangName() string { return "cefcConnectorRatingTable" }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) SetParent(parent types.Entity) { cefcconnectorratingtable.parent = parent }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetParent() types.Entity { return cefcconnectorratingtable.parent }

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry
// A cefcConnectorRatingEntry lists the
// power connector rating information of a 
// component in the system.
// 
// An entry or entries are created by the agent
// when an physical entity with connector rating 
// management is added to the ENTITY-MIB 
// entPhysicalTable. An entry is deleted 
// by the agent at the entity removal.
type CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The maximum power that the component's connector can withdraw. The type is
    // interface{} with range: -1000000000..1000000000.
    Cefcconnectorrating interface{}
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetFilter() yfilter.YFilter { return cefcconnectorratingentry.YFilter }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) SetFilter(yf yfilter.YFilter) { cefcconnectorratingentry.YFilter = yf }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcConnectorRating" { return "Cefcconnectorrating" }
    return ""
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetSegmentPath() string {
    return "cefcConnectorRatingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcconnectorratingentry.Entphysicalindex) + "']"
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcconnectorratingentry.Entphysicalindex
    leafs["cefcConnectorRating"] = cefcconnectorratingentry.Cefcconnectorrating
    return leafs
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetYangName() string { return "cefcConnectorRatingEntry" }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) SetParent(parent types.Entity) { cefcconnectorratingentry.parent = parent }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetParent() types.Entity { return cefcconnectorratingentry.parent }

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetParentYangName() string { return "cefcConnectorRatingTable" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable
// This table contains the total power consumption
// information for modules whose ENTITY-MIB 
// entPhysicalTable entries have an entPhysicalClass 
// of 'module'.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A cefcModulePowerConsumptionEntry lists the total power consumption of a
    // manageable components of type entPhysicalClass 'module'.  Entries are
    // created by the agent at the system power-up or module insertion.  Entries
    // are deleted by the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry.
    Cefcmodulepowerconsumptionentry []CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetFilter() yfilter.YFilter { return cefcmodulepowerconsumptiontable.YFilter }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) SetFilter(yf yfilter.YFilter) { cefcmodulepowerconsumptiontable.YFilter = yf }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetGoName(yname string) string {
    if yname == "cefcModulePowerConsumptionEntry" { return "Cefcmodulepowerconsumptionentry" }
    return ""
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetSegmentPath() string {
    return "cefcModulePowerConsumptionTable"
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cefcModulePowerConsumptionEntry" {
        for _, c := range cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry {
            if cefcmodulepowerconsumptiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry{}
        cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry = append(cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry, child)
        return &cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry[len(cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry)-1]
    }
    return nil
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry {
        children[cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry[i].GetSegmentPath()] = &cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry[i]
    }
    return children
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetYangName() string { return "cefcModulePowerConsumptionTable" }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) SetParent(parent types.Entity) { cefcmodulepowerconsumptiontable.parent = parent }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetParent() types.Entity { return cefcmodulepowerconsumptiontable.parent }

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetParentYangName() string { return "CISCO-ENTITY-FRU-CONTROL-MIB" }

// CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry
// A cefcModulePowerConsumptionEntry lists the total
// power consumption of a manageable components of type
// entPhysicalClass 'module'.
// 
// Entries are created by the agent at the system
// power-up or module insertion.
// 
// Entries are deleted by the agent upon module
// removal.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The combined power consumption to operate the module and its submodule(s)
    // and inline-power device(s). The type is interface{} with range:
    // -1000000000..1000000000.
    Cefcmodulepowerconsumption interface{}
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetFilter() yfilter.YFilter { return cefcmodulepowerconsumptionentry.YFilter }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) SetFilter(yf yfilter.YFilter) { cefcmodulepowerconsumptionentry.YFilter = yf }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cefcModulePowerConsumption" { return "Cefcmodulepowerconsumption" }
    return ""
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetSegmentPath() string {
    return "cefcModulePowerConsumptionEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulepowerconsumptionentry.Entphysicalindex) + "']"
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cefcmodulepowerconsumptionentry.Entphysicalindex
    leafs["cefcModulePowerConsumption"] = cefcmodulepowerconsumptionentry.Cefcmodulepowerconsumption
    return leafs
}

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetBundleName() string { return "cisco_ios_xe" }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetYangName() string { return "cefcModulePowerConsumptionEntry" }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) SetParent(parent types.Entity) { cefcmodulepowerconsumptionentry.parent = parent }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetParent() types.Entity { return cefcmodulepowerconsumptionentry.parent }

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetParentYangName() string { return "cefcModulePowerConsumptionTable" }

