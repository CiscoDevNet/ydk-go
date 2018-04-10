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

// PowerAdminType represents in response to a management protocol retrieval operation.
type PowerAdminType string

const (
    PowerAdminType_on PowerAdminType = "on"

    PowerAdminType_off PowerAdminType = "off"

    PowerAdminType_inlineAuto PowerAdminType = "inlineAuto"

    PowerAdminType_inlineOn PowerAdminType = "inlineOn"

    PowerAdminType_powerCycle PowerAdminType = "powerCycle"
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

// ModuleAdminType represents                        service, set by CLI.
type ModuleAdminType string

const (
    ModuleAdminType_enabled ModuleAdminType = "enabled"

    ModuleAdminType_disabled ModuleAdminType = "disabled"

    ModuleAdminType_reset ModuleAdminType = "reset"

    ModuleAdminType_outOfServiceAdmin ModuleAdminType = "outOfServiceAdmin"
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

// FRUCoolingUnit represents watts(2)  Watts
type FRUCoolingUnit string

const (
    FRUCoolingUnit_cfm FRUCoolingUnit = "cfm"

    FRUCoolingUnit_watts FRUCoolingUnit = "watts"
)

// CISCOENTITYFRUCONTROLMIB
type CISCOENTITYFRUCONTROLMIB struct {
    EntityData types.CommonEntityData
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

func (cISCOENTITYFRUCONTROLMIB *CISCOENTITYFRUCONTROLMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYFRUCONTROLMIB.EntityData.YFilter = cISCOENTITYFRUCONTROLMIB.YFilter
    cISCOENTITYFRUCONTROLMIB.EntityData.YangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENTITYFRUCONTROLMIB.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.SegmentPath = "CISCO-ENTITY-FRU-CONTROL-MIB:CISCO-ENTITY-FRU-CONTROL-MIB"
    cISCOENTITYFRUCONTROLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENTITYFRUCONTROLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENTITYFRUCONTROLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENTITYFRUCONTROLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFRUPower"] = types.YChild{"Cefcfrupower", &cISCOENTITYFRUCONTROLMIB.Cefcfrupower}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcMIBNotificationEnables"] = types.YChild{"Cefcmibnotificationenables", &cISCOENTITYFRUCONTROLMIB.Cefcmibnotificationenables}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFRUPowerSupplyGroupTable"] = types.YChild{"Cefcfrupowersupplygrouptable", &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplygrouptable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFRUPowerStatusTable"] = types.YChild{"Cefcfrupowerstatustable", &cISCOENTITYFRUCONTROLMIB.Cefcfrupowerstatustable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFRUPowerSupplyValueTable"] = types.YChild{"Cefcfrupowersupplyvaluetable", &cISCOENTITYFRUCONTROLMIB.Cefcfrupowersupplyvaluetable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcModuleTable"] = types.YChild{"Cefcmoduletable", &cISCOENTITYFRUCONTROLMIB.Cefcmoduletable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcIntelliModuleTable"] = types.YChild{"Cefcintellimoduletable", &cISCOENTITYFRUCONTROLMIB.Cefcintellimoduletable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcModuleLocalSwitchingTable"] = types.YChild{"Cefcmodulelocalswitchingtable", &cISCOENTITYFRUCONTROLMIB.Cefcmodulelocalswitchingtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFanTrayStatusTable"] = types.YChild{"Cefcfantraystatustable", &cISCOENTITYFRUCONTROLMIB.Cefcfantraystatustable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcPhysicalTable"] = types.YChild{"Cefcphysicaltable", &cISCOENTITYFRUCONTROLMIB.Cefcphysicaltable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcPowerSupplyInputTable"] = types.YChild{"Cefcpowersupplyinputtable", &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyinputtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcPowerSupplyOutputTable"] = types.YChild{"Cefcpowersupplyoutputtable", &cISCOENTITYFRUCONTROLMIB.Cefcpowersupplyoutputtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcChassisCoolingTable"] = types.YChild{"Cefcchassiscoolingtable", &cISCOENTITYFRUCONTROLMIB.Cefcchassiscoolingtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFanCoolingTable"] = types.YChild{"Cefcfancoolingtable", &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcModuleCoolingTable"] = types.YChild{"Cefcmodulecoolingtable", &cISCOENTITYFRUCONTROLMIB.Cefcmodulecoolingtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcFanCoolingCapTable"] = types.YChild{"Cefcfancoolingcaptable", &cISCOENTITYFRUCONTROLMIB.Cefcfancoolingcaptable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcConnectorRatingTable"] = types.YChild{"Cefcconnectorratingtable", &cISCOENTITYFRUCONTROLMIB.Cefcconnectorratingtable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Children["cefcModulePowerConsumptionTable"] = types.YChild{"Cefcmodulepowerconsumptiontable", &cISCOENTITYFRUCONTROLMIB.Cefcmodulepowerconsumptiontable}
    cISCOENTITYFRUCONTROLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENTITYFRUCONTROLMIB.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcfrupower
type CISCOENTITYFRUCONTROLMIB_Cefcfrupower struct {
    EntityData types.CommonEntityData
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

func (cefcfrupower *CISCOENTITYFRUCONTROLMIB_Cefcfrupower) GetEntityData() *types.CommonEntityData {
    cefcfrupower.EntityData.YFilter = cefcfrupower.YFilter
    cefcfrupower.EntityData.YangName = "cefcFRUPower"
    cefcfrupower.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupower.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfrupower.EntityData.SegmentPath = "cefcFRUPower"
    cefcfrupower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupower.EntityData.Children = make(map[string]types.YChild)
    cefcfrupower.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfrupower.EntityData.Leafs["cefcMaxDefaultInLinePower"] = types.YLeaf{"Cefcmaxdefaultinlinepower", cefcfrupower.Cefcmaxdefaultinlinepower}
    cefcfrupower.EntityData.Leafs["cefcMaxDefaultHighInLinePower"] = types.YLeaf{"Cefcmaxdefaulthighinlinepower", cefcfrupower.Cefcmaxdefaulthighinlinepower}
    return &(cefcfrupower.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables
type CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables struct {
    EntityData types.CommonEntityData
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

func (cefcmibnotificationenables *CISCOENTITYFRUCONTROLMIB_Cefcmibnotificationenables) GetEntityData() *types.CommonEntityData {
    cefcmibnotificationenables.EntityData.YFilter = cefcmibnotificationenables.YFilter
    cefcmibnotificationenables.EntityData.YangName = "cefcMIBNotificationEnables"
    cefcmibnotificationenables.EntityData.BundleName = "cisco_ios_xe"
    cefcmibnotificationenables.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcmibnotificationenables.EntityData.SegmentPath = "cefcMIBNotificationEnables"
    cefcmibnotificationenables.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmibnotificationenables.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmibnotificationenables.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmibnotificationenables.EntityData.Children = make(map[string]types.YChild)
    cefcmibnotificationenables.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcmibnotificationenables.EntityData.Leafs["cefcMIBEnableStatusNotification"] = types.YLeaf{"Cefcmibenablestatusnotification", cefcmibnotificationenables.Cefcmibenablestatusnotification}
    cefcmibnotificationenables.EntityData.Leafs["cefcEnablePSOutputChangeNotif"] = types.YLeaf{"Cefcenablepsoutputchangenotif", cefcmibnotificationenables.Cefcenablepsoutputchangenotif}
    return &(cefcmibnotificationenables.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable
// This table lists the redundancy mode and the
// operational status of the power supply groups
// in the system.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyGroupTable entry lists the desired redundancy mode,
    // the units of the power outputs and the  available and drawn current for the
    // power supply group.  Entries are created by the agent when a power supply
    // group is added to the entPhysicalTable. Entries are deleted by  the agent
    // at power supply group removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry.
    Cefcfrupowersupplygroupentry []CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry
}

func (cefcfrupowersupplygrouptable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable) GetEntityData() *types.CommonEntityData {
    cefcfrupowersupplygrouptable.EntityData.YFilter = cefcfrupowersupplygrouptable.YFilter
    cefcfrupowersupplygrouptable.EntityData.YangName = "cefcFRUPowerSupplyGroupTable"
    cefcfrupowersupplygrouptable.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowersupplygrouptable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfrupowersupplygrouptable.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupTable"
    cefcfrupowersupplygrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowersupplygrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowersupplygrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowersupplygrouptable.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowersupplygrouptable.EntityData.Children["cefcFRUPowerSupplyGroupEntry"] = types.YChild{"Cefcfrupowersupplygroupentry", nil}
    for i := range cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry {
        cefcfrupowersupplygrouptable.EntityData.Children[types.GetSegmentPath(&cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry[i])] = types.YChild{"Cefcfrupowersupplygroupentry", &cefcfrupowersupplygrouptable.Cefcfrupowersupplygroupentry[i]}
    }
    cefcfrupowersupplygrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfrupowersupplygrouptable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry
// An cefcFRUPowerSupplyGroupTable entry lists the desired
// redundancy mode, the units of the power outputs and the 
// available and drawn current for the power supply group.
// 
// Entries are created by the agent when a power supply group
// is added to the entPhysicalTable. Entries are deleted by 
// the agent at power supply group removal.
type CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry struct {
    EntityData types.CommonEntityData
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

func (cefcfrupowersupplygroupentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplygrouptable_Cefcfrupowersupplygroupentry) GetEntityData() *types.CommonEntityData {
    cefcfrupowersupplygroupentry.EntityData.YFilter = cefcfrupowersupplygroupentry.YFilter
    cefcfrupowersupplygroupentry.EntityData.YangName = "cefcFRUPowerSupplyGroupEntry"
    cefcfrupowersupplygroupentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowersupplygroupentry.EntityData.ParentYangName = "cefcFRUPowerSupplyGroupTable"
    cefcfrupowersupplygroupentry.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowersupplygroupentry.Entphysicalindex) + "']"
    cefcfrupowersupplygroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowersupplygroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowersupplygroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowersupplygroupentry.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowersupplygroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfrupowersupplygroupentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfrupowersupplygroupentry.Entphysicalindex}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcPowerRedundancyMode"] = types.YLeaf{"Cefcpowerredundancymode", cefcfrupowersupplygroupentry.Cefcpowerredundancymode}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcPowerUnits"] = types.YLeaf{"Cefcpowerunits", cefcfrupowersupplygroupentry.Cefcpowerunits}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcTotalAvailableCurrent"] = types.YLeaf{"Cefctotalavailablecurrent", cefcfrupowersupplygroupentry.Cefctotalavailablecurrent}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcTotalDrawnCurrent"] = types.YLeaf{"Cefctotaldrawncurrent", cefcfrupowersupplygroupentry.Cefctotaldrawncurrent}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcPowerRedundancyOperMode"] = types.YLeaf{"Cefcpowerredundancyopermode", cefcfrupowersupplygroupentry.Cefcpowerredundancyopermode}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcPowerNonRedundantReason"] = types.YLeaf{"Cefcpowernonredundantreason", cefcfrupowersupplygroupentry.Cefcpowernonredundantreason}
    cefcfrupowersupplygroupentry.EntityData.Leafs["cefcTotalDrawnInlineCurrent"] = types.YLeaf{"Cefctotaldrawninlinecurrent", cefcfrupowersupplygroupentry.Cefctotaldrawninlinecurrent}
    return &(cefcfrupowersupplygroupentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfrupowerstatustable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable) GetEntityData() *types.CommonEntityData {
    cefcfrupowerstatustable.EntityData.YFilter = cefcfrupowerstatustable.YFilter
    cefcfrupowerstatustable.EntityData.YangName = "cefcFRUPowerStatusTable"
    cefcfrupowerstatustable.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowerstatustable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfrupowerstatustable.EntityData.SegmentPath = "cefcFRUPowerStatusTable"
    cefcfrupowerstatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowerstatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowerstatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowerstatustable.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowerstatustable.EntityData.Children["cefcFRUPowerStatusEntry"] = types.YChild{"Cefcfrupowerstatusentry", nil}
    for i := range cefcfrupowerstatustable.Cefcfrupowerstatusentry {
        cefcfrupowerstatustable.EntityData.Children[types.GetSegmentPath(&cefcfrupowerstatustable.Cefcfrupowerstatusentry[i])] = types.YChild{"Cefcfrupowerstatusentry", &cefcfrupowerstatustable.Cefcfrupowerstatusentry[i]}
    }
    cefcfrupowerstatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfrupowerstatustable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfrupowerstatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowerstatustable_Cefcfrupowerstatusentry) GetEntityData() *types.CommonEntityData {
    cefcfrupowerstatusentry.EntityData.YFilter = cefcfrupowerstatusentry.YFilter
    cefcfrupowerstatusentry.EntityData.YangName = "cefcFRUPowerStatusEntry"
    cefcfrupowerstatusentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowerstatusentry.EntityData.ParentYangName = "cefcFRUPowerStatusTable"
    cefcfrupowerstatusentry.EntityData.SegmentPath = "cefcFRUPowerStatusEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowerstatusentry.Entphysicalindex) + "']"
    cefcfrupowerstatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowerstatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowerstatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowerstatusentry.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowerstatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfrupowerstatusentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfrupowerstatusentry.Entphysicalindex}
    cefcfrupowerstatusentry.EntityData.Leafs["cefcFRUPowerAdminStatus"] = types.YLeaf{"Cefcfrupoweradminstatus", cefcfrupowerstatusentry.Cefcfrupoweradminstatus}
    cefcfrupowerstatusentry.EntityData.Leafs["cefcFRUPowerOperStatus"] = types.YLeaf{"Cefcfrupoweroperstatus", cefcfrupowerstatusentry.Cefcfrupoweroperstatus}
    cefcfrupowerstatusentry.EntityData.Leafs["cefcFRUCurrent"] = types.YLeaf{"Cefcfrucurrent", cefcfrupowerstatusentry.Cefcfrucurrent}
    cefcfrupowerstatusentry.EntityData.Leafs["cefcFRUPowerCapability"] = types.YLeaf{"Cefcfrupowercapability", cefcfrupowerstatusentry.Cefcfrupowercapability}
    cefcfrupowerstatusentry.EntityData.Leafs["cefcFRURealTimeCurrent"] = types.YLeaf{"Cefcfrurealtimecurrent", cefcfrupowerstatusentry.Cefcfrurealtimecurrent}
    return &(cefcfrupowerstatusentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyValueTable entry lists the current provided by the FRU
    // for operation.  Entries are created by the agent at system power-up or  FRU
    // insertion.  Entries are deleted by the agent at FRU removal.  Only power
    // supply FRUs are listed in the table. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry.
    Cefcfrupowersupplyvalueentry []CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry
}

func (cefcfrupowersupplyvaluetable *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable) GetEntityData() *types.CommonEntityData {
    cefcfrupowersupplyvaluetable.EntityData.YFilter = cefcfrupowersupplyvaluetable.YFilter
    cefcfrupowersupplyvaluetable.EntityData.YangName = "cefcFRUPowerSupplyValueTable"
    cefcfrupowersupplyvaluetable.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowersupplyvaluetable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfrupowersupplyvaluetable.EntityData.SegmentPath = "cefcFRUPowerSupplyValueTable"
    cefcfrupowersupplyvaluetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowersupplyvaluetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowersupplyvaluetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowersupplyvaluetable.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowersupplyvaluetable.EntityData.Children["cefcFRUPowerSupplyValueEntry"] = types.YChild{"Cefcfrupowersupplyvalueentry", nil}
    for i := range cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry {
        cefcfrupowersupplyvaluetable.EntityData.Children[types.GetSegmentPath(&cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry[i])] = types.YChild{"Cefcfrupowersupplyvalueentry", &cefcfrupowersupplyvaluetable.Cefcfrupowersupplyvalueentry[i]}
    }
    cefcfrupowersupplyvaluetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfrupowersupplyvaluetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfrupowersupplyvalueentry *CISCOENTITYFRUCONTROLMIB_Cefcfrupowersupplyvaluetable_Cefcfrupowersupplyvalueentry) GetEntityData() *types.CommonEntityData {
    cefcfrupowersupplyvalueentry.EntityData.YFilter = cefcfrupowersupplyvalueentry.YFilter
    cefcfrupowersupplyvalueentry.EntityData.YangName = "cefcFRUPowerSupplyValueEntry"
    cefcfrupowersupplyvalueentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfrupowersupplyvalueentry.EntityData.ParentYangName = "cefcFRUPowerSupplyValueTable"
    cefcfrupowersupplyvalueentry.EntityData.SegmentPath = "cefcFRUPowerSupplyValueEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfrupowersupplyvalueentry.Entphysicalindex) + "']"
    cefcfrupowersupplyvalueentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfrupowersupplyvalueentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfrupowersupplyvalueentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfrupowersupplyvalueentry.EntityData.Children = make(map[string]types.YChild)
    cefcfrupowersupplyvalueentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfrupowersupplyvalueentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfrupowersupplyvalueentry.Entphysicalindex}
    cefcfrupowersupplyvalueentry.EntityData.Leafs["cefcFRUTotalSystemCurrent"] = types.YLeaf{"Cefcfrutotalsystemcurrent", cefcfrupowersupplyvalueentry.Cefcfrutotalsystemcurrent}
    cefcfrupowersupplyvalueentry.EntityData.Leafs["cefcFRUDrawnSystemCurrent"] = types.YLeaf{"Cefcfrudrawnsystemcurrent", cefcfrupowersupplyvalueentry.Cefcfrudrawnsystemcurrent}
    cefcfrupowersupplyvalueentry.EntityData.Leafs["cefcFRUTotalInlineCurrent"] = types.YLeaf{"Cefcfrutotalinlinecurrent", cefcfrupowersupplyvalueentry.Cefcfrutotalinlinecurrent}
    cefcfrupowersupplyvalueentry.EntityData.Leafs["cefcFRUDrawnInlineCurrent"] = types.YLeaf{"Cefcfrudrawninlinecurrent", cefcfrupowersupplyvalueentry.Cefcfrudrawninlinecurrent}
    return &(cefcfrupowersupplyvalueentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcmoduletable
// A cefcModuleTable entry lists the operational and
// administrative status information for ENTITY-MIB
// entPhysicalTable entries for manageable components
// of type PhysicalClass module(9).
type CISCOENTITYFRUCONTROLMIB_Cefcmoduletable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModuleStatusTable entry lists the operational and administrative
    // status information for ENTITY-MIB entPhysicalTable entries for manageable
    // components  of type PhysicalClass module(9).  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry.
    Cefcmoduleentry []CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry
}

func (cefcmoduletable *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable) GetEntityData() *types.CommonEntityData {
    cefcmoduletable.EntityData.YFilter = cefcmoduletable.YFilter
    cefcmoduletable.EntityData.YangName = "cefcModuleTable"
    cefcmoduletable.EntityData.BundleName = "cisco_ios_xe"
    cefcmoduletable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcmoduletable.EntityData.SegmentPath = "cefcModuleTable"
    cefcmoduletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmoduletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmoduletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmoduletable.EntityData.Children = make(map[string]types.YChild)
    cefcmoduletable.EntityData.Children["cefcModuleEntry"] = types.YChild{"Cefcmoduleentry", nil}
    for i := range cefcmoduletable.Cefcmoduleentry {
        cefcmoduletable.EntityData.Children[types.GetSegmentPath(&cefcmoduletable.Cefcmoduleentry[i])] = types.YChild{"Cefcmoduleentry", &cefcmoduletable.Cefcmoduleentry[i]}
    }
    cefcmoduletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcmoduletable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcmoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcmoduletable_Cefcmoduleentry) GetEntityData() *types.CommonEntityData {
    cefcmoduleentry.EntityData.YFilter = cefcmoduleentry.YFilter
    cefcmoduleentry.EntityData.YangName = "cefcModuleEntry"
    cefcmoduleentry.EntityData.BundleName = "cisco_ios_xe"
    cefcmoduleentry.EntityData.ParentYangName = "cefcModuleTable"
    cefcmoduleentry.EntityData.SegmentPath = "cefcModuleEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmoduleentry.Entphysicalindex) + "']"
    cefcmoduleentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmoduleentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmoduleentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmoduleentry.EntityData.Children = make(map[string]types.YChild)
    cefcmoduleentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcmoduleentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcmoduleentry.Entphysicalindex}
    cefcmoduleentry.EntityData.Leafs["cefcModuleAdminStatus"] = types.YLeaf{"Cefcmoduleadminstatus", cefcmoduleentry.Cefcmoduleadminstatus}
    cefcmoduleentry.EntityData.Leafs["cefcModuleOperStatus"] = types.YLeaf{"Cefcmoduleoperstatus", cefcmoduleentry.Cefcmoduleoperstatus}
    cefcmoduleentry.EntityData.Leafs["cefcModuleResetReason"] = types.YLeaf{"Cefcmoduleresetreason", cefcmoduleentry.Cefcmoduleresetreason}
    cefcmoduleentry.EntityData.Leafs["cefcModuleStatusLastChangeTime"] = types.YLeaf{"Cefcmodulestatuslastchangetime", cefcmoduleentry.Cefcmodulestatuslastchangetime}
    cefcmoduleentry.EntityData.Leafs["cefcModuleLastClearConfigTime"] = types.YLeaf{"Cefcmodulelastclearconfigtime", cefcmoduleentry.Cefcmodulelastclearconfigtime}
    cefcmoduleentry.EntityData.Leafs["cefcModuleResetReasonDescription"] = types.YLeaf{"Cefcmoduleresetreasondescription", cefcmoduleentry.Cefcmoduleresetreasondescription}
    cefcmoduleentry.EntityData.Leafs["cefcModuleStateChangeReasonDescr"] = types.YLeaf{"Cefcmodulestatechangereasondescr", cefcmoduleentry.Cefcmodulestatechangereasondescr}
    cefcmoduleentry.EntityData.Leafs["cefcModuleUpTime"] = types.YLeaf{"Cefcmoduleuptime", cefcmoduleentry.Cefcmoduleuptime}
    return &(cefcmoduleentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcintellimoduletable *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable) GetEntityData() *types.CommonEntityData {
    cefcintellimoduletable.EntityData.YFilter = cefcintellimoduletable.YFilter
    cefcintellimoduletable.EntityData.YangName = "cefcIntelliModuleTable"
    cefcintellimoduletable.EntityData.BundleName = "cisco_ios_xe"
    cefcintellimoduletable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcintellimoduletable.EntityData.SegmentPath = "cefcIntelliModuleTable"
    cefcintellimoduletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcintellimoduletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcintellimoduletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcintellimoduletable.EntityData.Children = make(map[string]types.YChild)
    cefcintellimoduletable.EntityData.Children["cefcIntelliModuleEntry"] = types.YChild{"Cefcintellimoduleentry", nil}
    for i := range cefcintellimoduletable.Cefcintellimoduleentry {
        cefcintellimoduletable.EntityData.Children[types.GetSegmentPath(&cefcintellimoduletable.Cefcintellimoduleentry[i])] = types.YChild{"Cefcintellimoduleentry", &cefcintellimoduletable.Cefcintellimoduleentry[i]}
    }
    cefcintellimoduletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcintellimoduletable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcintellimoduleentry *CISCOENTITYFRUCONTROLMIB_Cefcintellimoduletable_Cefcintellimoduleentry) GetEntityData() *types.CommonEntityData {
    cefcintellimoduleentry.EntityData.YFilter = cefcintellimoduleentry.YFilter
    cefcintellimoduleentry.EntityData.YangName = "cefcIntelliModuleEntry"
    cefcintellimoduleentry.EntityData.BundleName = "cisco_ios_xe"
    cefcintellimoduleentry.EntityData.ParentYangName = "cefcIntelliModuleTable"
    cefcintellimoduleentry.EntityData.SegmentPath = "cefcIntelliModuleEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcintellimoduleentry.Entphysicalindex) + "']"
    cefcintellimoduleentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcintellimoduleentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcintellimoduleentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcintellimoduleentry.EntityData.Children = make(map[string]types.YChild)
    cefcintellimoduleentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcintellimoduleentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcintellimoduleentry.Entphysicalindex}
    cefcintellimoduleentry.EntityData.Leafs["cefcIntelliModuleIPAddrType"] = types.YLeaf{"Cefcintellimoduleipaddrtype", cefcintellimoduleentry.Cefcintellimoduleipaddrtype}
    cefcintellimoduleentry.EntityData.Leafs["cefcIntelliModuleIPAddr"] = types.YLeaf{"Cefcintellimoduleipaddr", cefcintellimoduleentry.Cefcintellimoduleipaddr}
    return &(cefcintellimoduleentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcmodulelocalswitchingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable) GetEntityData() *types.CommonEntityData {
    cefcmodulelocalswitchingtable.EntityData.YFilter = cefcmodulelocalswitchingtable.YFilter
    cefcmodulelocalswitchingtable.EntityData.YangName = "cefcModuleLocalSwitchingTable"
    cefcmodulelocalswitchingtable.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulelocalswitchingtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcmodulelocalswitchingtable.EntityData.SegmentPath = "cefcModuleLocalSwitchingTable"
    cefcmodulelocalswitchingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulelocalswitchingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulelocalswitchingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulelocalswitchingtable.EntityData.Children = make(map[string]types.YChild)
    cefcmodulelocalswitchingtable.EntityData.Children["cefcModuleLocalSwitchingEntry"] = types.YChild{"Cefcmodulelocalswitchingentry", nil}
    for i := range cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry {
        cefcmodulelocalswitchingtable.EntityData.Children[types.GetSegmentPath(&cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry[i])] = types.YChild{"Cefcmodulelocalswitchingentry", &cefcmodulelocalswitchingtable.Cefcmodulelocalswitchingentry[i]}
    }
    cefcmodulelocalswitchingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcmodulelocalswitchingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcmodulelocalswitchingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulelocalswitchingtable_Cefcmodulelocalswitchingentry) GetEntityData() *types.CommonEntityData {
    cefcmodulelocalswitchingentry.EntityData.YFilter = cefcmodulelocalswitchingentry.YFilter
    cefcmodulelocalswitchingentry.EntityData.YangName = "cefcModuleLocalSwitchingEntry"
    cefcmodulelocalswitchingentry.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulelocalswitchingentry.EntityData.ParentYangName = "cefcModuleLocalSwitchingTable"
    cefcmodulelocalswitchingentry.EntityData.SegmentPath = "cefcModuleLocalSwitchingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulelocalswitchingentry.Entphysicalindex) + "']"
    cefcmodulelocalswitchingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulelocalswitchingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulelocalswitchingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulelocalswitchingentry.EntityData.Children = make(map[string]types.YChild)
    cefcmodulelocalswitchingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcmodulelocalswitchingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcmodulelocalswitchingentry.Entphysicalindex}
    cefcmodulelocalswitchingentry.EntityData.Leafs["cefcModuleLocalSwitchingMode"] = types.YLeaf{"Cefcmodulelocalswitchingmode", cefcmodulelocalswitchingentry.Cefcmodulelocalswitchingmode}
    return &(cefcmodulelocalswitchingentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfantraystatustable *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable) GetEntityData() *types.CommonEntityData {
    cefcfantraystatustable.EntityData.YFilter = cefcfantraystatustable.YFilter
    cefcfantraystatustable.EntityData.YangName = "cefcFanTrayStatusTable"
    cefcfantraystatustable.EntityData.BundleName = "cisco_ios_xe"
    cefcfantraystatustable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfantraystatustable.EntityData.SegmentPath = "cefcFanTrayStatusTable"
    cefcfantraystatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfantraystatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfantraystatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfantraystatustable.EntityData.Children = make(map[string]types.YChild)
    cefcfantraystatustable.EntityData.Children["cefcFanTrayStatusEntry"] = types.YChild{"Cefcfantraystatusentry", nil}
    for i := range cefcfantraystatustable.Cefcfantraystatusentry {
        cefcfantraystatustable.EntityData.Children[types.GetSegmentPath(&cefcfantraystatustable.Cefcfantraystatusentry[i])] = types.YChild{"Cefcfantraystatusentry", &cefcfantraystatustable.Cefcfantraystatusentry[i]}
    }
    cefcfantraystatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfantraystatustable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfantraystatusentry *CISCOENTITYFRUCONTROLMIB_Cefcfantraystatustable_Cefcfantraystatusentry) GetEntityData() *types.CommonEntityData {
    cefcfantraystatusentry.EntityData.YFilter = cefcfantraystatusentry.YFilter
    cefcfantraystatusentry.EntityData.YangName = "cefcFanTrayStatusEntry"
    cefcfantraystatusentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfantraystatusentry.EntityData.ParentYangName = "cefcFanTrayStatusTable"
    cefcfantraystatusentry.EntityData.SegmentPath = "cefcFanTrayStatusEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfantraystatusentry.Entphysicalindex) + "']"
    cefcfantraystatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfantraystatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfantraystatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfantraystatusentry.EntityData.Children = make(map[string]types.YChild)
    cefcfantraystatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfantraystatusentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfantraystatusentry.Entphysicalindex}
    cefcfantraystatusentry.EntityData.Leafs["cefcFanTrayOperStatus"] = types.YLeaf{"Cefcfantrayoperstatus", cefcfantraystatusentry.Cefcfantrayoperstatus}
    return &(cefcfantraystatusentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical entity. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry.
    Cefcphysicalentry []CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry
}

func (cefcphysicaltable *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable) GetEntityData() *types.CommonEntityData {
    cefcphysicaltable.EntityData.YFilter = cefcphysicaltable.YFilter
    cefcphysicaltable.EntityData.YangName = "cefcPhysicalTable"
    cefcphysicaltable.EntityData.BundleName = "cisco_ios_xe"
    cefcphysicaltable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcphysicaltable.EntityData.SegmentPath = "cefcPhysicalTable"
    cefcphysicaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcphysicaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcphysicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcphysicaltable.EntityData.Children = make(map[string]types.YChild)
    cefcphysicaltable.EntityData.Children["cefcPhysicalEntry"] = types.YChild{"Cefcphysicalentry", nil}
    for i := range cefcphysicaltable.Cefcphysicalentry {
        cefcphysicaltable.EntityData.Children[types.GetSegmentPath(&cefcphysicaltable.Cefcphysicalentry[i])] = types.YChild{"Cefcphysicalentry", &cefcphysicaltable.Cefcphysicalentry[i]}
    }
    cefcphysicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcphysicaltable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry
// Information about a particular physical entity.
type CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry struct {
    EntityData types.CommonEntityData
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

func (cefcphysicalentry *CISCOENTITYFRUCONTROLMIB_Cefcphysicaltable_Cefcphysicalentry) GetEntityData() *types.CommonEntityData {
    cefcphysicalentry.EntityData.YFilter = cefcphysicalentry.YFilter
    cefcphysicalentry.EntityData.YangName = "cefcPhysicalEntry"
    cefcphysicalentry.EntityData.BundleName = "cisco_ios_xe"
    cefcphysicalentry.EntityData.ParentYangName = "cefcPhysicalTable"
    cefcphysicalentry.EntityData.SegmentPath = "cefcPhysicalEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcphysicalentry.Entphysicalindex) + "']"
    cefcphysicalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcphysicalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcphysicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcphysicalentry.EntityData.Children = make(map[string]types.YChild)
    cefcphysicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcphysicalentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcphysicalentry.Entphysicalindex}
    cefcphysicalentry.EntityData.Leafs["cefcPhysicalStatus"] = types.YLeaf{"Cefcphysicalstatus", cefcphysicalentry.Cefcphysicalstatus}
    return &(cefcphysicalentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing power input management information applicable to a
    // particular power supply and input. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry.
    Cefcpowersupplyinputentry []CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry
}

func (cefcpowersupplyinputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable) GetEntityData() *types.CommonEntityData {
    cefcpowersupplyinputtable.EntityData.YFilter = cefcpowersupplyinputtable.YFilter
    cefcpowersupplyinputtable.EntityData.YangName = "cefcPowerSupplyInputTable"
    cefcpowersupplyinputtable.EntityData.BundleName = "cisco_ios_xe"
    cefcpowersupplyinputtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcpowersupplyinputtable.EntityData.SegmentPath = "cefcPowerSupplyInputTable"
    cefcpowersupplyinputtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcpowersupplyinputtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcpowersupplyinputtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcpowersupplyinputtable.EntityData.Children = make(map[string]types.YChild)
    cefcpowersupplyinputtable.EntityData.Children["cefcPowerSupplyInputEntry"] = types.YChild{"Cefcpowersupplyinputentry", nil}
    for i := range cefcpowersupplyinputtable.Cefcpowersupplyinputentry {
        cefcpowersupplyinputtable.EntityData.Children[types.GetSegmentPath(&cefcpowersupplyinputtable.Cefcpowersupplyinputentry[i])] = types.YChild{"Cefcpowersupplyinputentry", &cefcpowersupplyinputtable.Cefcpowersupplyinputentry[i]}
    }
    cefcpowersupplyinputtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcpowersupplyinputtable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry
// An entry containing power input management information
// applicable to a particular power supply and input.
type CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry struct {
    EntityData types.CommonEntityData
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

func (cefcpowersupplyinputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyinputtable_Cefcpowersupplyinputentry) GetEntityData() *types.CommonEntityData {
    cefcpowersupplyinputentry.EntityData.YFilter = cefcpowersupplyinputentry.YFilter
    cefcpowersupplyinputentry.EntityData.YangName = "cefcPowerSupplyInputEntry"
    cefcpowersupplyinputentry.EntityData.BundleName = "cisco_ios_xe"
    cefcpowersupplyinputentry.EntityData.ParentYangName = "cefcPowerSupplyInputTable"
    cefcpowersupplyinputentry.EntityData.SegmentPath = "cefcPowerSupplyInputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcpowersupplyinputentry.Entphysicalindex) + "']" + "[cefcPowerSupplyInputIndex='" + fmt.Sprintf("%v", cefcpowersupplyinputentry.Cefcpowersupplyinputindex) + "']"
    cefcpowersupplyinputentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcpowersupplyinputentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcpowersupplyinputentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcpowersupplyinputentry.EntityData.Children = make(map[string]types.YChild)
    cefcpowersupplyinputentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcpowersupplyinputentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcpowersupplyinputentry.Entphysicalindex}
    cefcpowersupplyinputentry.EntityData.Leafs["cefcPowerSupplyInputIndex"] = types.YLeaf{"Cefcpowersupplyinputindex", cefcpowersupplyinputentry.Cefcpowersupplyinputindex}
    cefcpowersupplyinputentry.EntityData.Leafs["cefcPowerSupplyInputType"] = types.YLeaf{"Cefcpowersupplyinputtype", cefcpowersupplyinputentry.Cefcpowersupplyinputtype}
    return &(cefcpowersupplyinputentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcpowersupplyoutputtable *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable) GetEntityData() *types.CommonEntityData {
    cefcpowersupplyoutputtable.EntityData.YFilter = cefcpowersupplyoutputtable.YFilter
    cefcpowersupplyoutputtable.EntityData.YangName = "cefcPowerSupplyOutputTable"
    cefcpowersupplyoutputtable.EntityData.BundleName = "cisco_ios_xe"
    cefcpowersupplyoutputtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcpowersupplyoutputtable.EntityData.SegmentPath = "cefcPowerSupplyOutputTable"
    cefcpowersupplyoutputtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcpowersupplyoutputtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcpowersupplyoutputtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcpowersupplyoutputtable.EntityData.Children = make(map[string]types.YChild)
    cefcpowersupplyoutputtable.EntityData.Children["cefcPowerSupplyOutputEntry"] = types.YChild{"Cefcpowersupplyoutputentry", nil}
    for i := range cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry {
        cefcpowersupplyoutputtable.EntityData.Children[types.GetSegmentPath(&cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry[i])] = types.YChild{"Cefcpowersupplyoutputentry", &cefcpowersupplyoutputtable.Cefcpowersupplyoutputentry[i]}
    }
    cefcpowersupplyoutputtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcpowersupplyoutputtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcpowersupplyoutputentry *CISCOENTITYFRUCONTROLMIB_Cefcpowersupplyoutputtable_Cefcpowersupplyoutputentry) GetEntityData() *types.CommonEntityData {
    cefcpowersupplyoutputentry.EntityData.YFilter = cefcpowersupplyoutputentry.YFilter
    cefcpowersupplyoutputentry.EntityData.YangName = "cefcPowerSupplyOutputEntry"
    cefcpowersupplyoutputentry.EntityData.BundleName = "cisco_ios_xe"
    cefcpowersupplyoutputentry.EntityData.ParentYangName = "cefcPowerSupplyOutputTable"
    cefcpowersupplyoutputentry.EntityData.SegmentPath = "cefcPowerSupplyOutputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcpowersupplyoutputentry.Entphysicalindex) + "']" + "[cefcPSOutputModeIndex='" + fmt.Sprintf("%v", cefcpowersupplyoutputentry.Cefcpsoutputmodeindex) + "']"
    cefcpowersupplyoutputentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcpowersupplyoutputentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcpowersupplyoutputentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcpowersupplyoutputentry.EntityData.Children = make(map[string]types.YChild)
    cefcpowersupplyoutputentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcpowersupplyoutputentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcpowersupplyoutputentry.Entphysicalindex}
    cefcpowersupplyoutputentry.EntityData.Leafs["cefcPSOutputModeIndex"] = types.YLeaf{"Cefcpsoutputmodeindex", cefcpowersupplyoutputentry.Cefcpsoutputmodeindex}
    cefcpowersupplyoutputentry.EntityData.Leafs["cefcPSOutputModeCurrent"] = types.YLeaf{"Cefcpsoutputmodecurrent", cefcpowersupplyoutputentry.Cefcpsoutputmodecurrent}
    cefcpowersupplyoutputentry.EntityData.Leafs["cefcPSOutputModeInOperation"] = types.YLeaf{"Cefcpsoutputmodeinoperation", cefcpowersupplyoutputentry.Cefcpsoutputmodeinoperation}
    return &(cefcpowersupplyoutputentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable
// This table contains the cooling capacity
// information of the chassis whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'chassis'.
type CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable struct {
    EntityData types.CommonEntityData
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

func (cefcchassiscoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable) GetEntityData() *types.CommonEntityData {
    cefcchassiscoolingtable.EntityData.YFilter = cefcchassiscoolingtable.YFilter
    cefcchassiscoolingtable.EntityData.YangName = "cefcChassisCoolingTable"
    cefcchassiscoolingtable.EntityData.BundleName = "cisco_ios_xe"
    cefcchassiscoolingtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcchassiscoolingtable.EntityData.SegmentPath = "cefcChassisCoolingTable"
    cefcchassiscoolingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcchassiscoolingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcchassiscoolingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcchassiscoolingtable.EntityData.Children = make(map[string]types.YChild)
    cefcchassiscoolingtable.EntityData.Children["cefcChassisCoolingEntry"] = types.YChild{"Cefcchassiscoolingentry", nil}
    for i := range cefcchassiscoolingtable.Cefcchassiscoolingentry {
        cefcchassiscoolingtable.EntityData.Children[types.GetSegmentPath(&cefcchassiscoolingtable.Cefcchassiscoolingentry[i])] = types.YChild{"Cefcchassiscoolingentry", &cefcchassiscoolingtable.Cefcchassiscoolingentry[i]}
    }
    cefcchassiscoolingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcchassiscoolingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcchassiscoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcchassiscoolingtable_Cefcchassiscoolingentry) GetEntityData() *types.CommonEntityData {
    cefcchassiscoolingentry.EntityData.YFilter = cefcchassiscoolingentry.YFilter
    cefcchassiscoolingentry.EntityData.YangName = "cefcChassisCoolingEntry"
    cefcchassiscoolingentry.EntityData.BundleName = "cisco_ios_xe"
    cefcchassiscoolingentry.EntityData.ParentYangName = "cefcChassisCoolingTable"
    cefcchassiscoolingentry.EntityData.SegmentPath = "cefcChassisCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcchassiscoolingentry.Entphysicalindex) + "']"
    cefcchassiscoolingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcchassiscoolingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcchassiscoolingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcchassiscoolingentry.EntityData.Children = make(map[string]types.YChild)
    cefcchassiscoolingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcchassiscoolingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcchassiscoolingentry.Entphysicalindex}
    cefcchassiscoolingentry.EntityData.Leafs["cefcChassisPerSlotCoolingCap"] = types.YLeaf{"Cefcchassisperslotcoolingcap", cefcchassiscoolingentry.Cefcchassisperslotcoolingcap}
    cefcchassiscoolingentry.EntityData.Leafs["cefcChassisPerSlotCoolingUnit"] = types.YLeaf{"Cefcchassisperslotcoolingunit", cefcchassiscoolingentry.Cefcchassisperslotcoolingunit}
    return &(cefcchassiscoolingentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable
// This table contains the cooling capacity
// information of the fans whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcFanCoolingEntry lists the cooling capacity that is provided by the 
    // manageable components of type PhysicalClass  'fan'.  Entries are created by
    // the agent if the corresponding entry is created in ENTITY-MIB
    // entPhysicalTable.  Entries are deleted by the agent if the corresponding
    // entry is deleted in ENTITY-MIB entPhysicalTable. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry.
    Cefcfancoolingentry []CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry
}

func (cefcfancoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable) GetEntityData() *types.CommonEntityData {
    cefcfancoolingtable.EntityData.YFilter = cefcfancoolingtable.YFilter
    cefcfancoolingtable.EntityData.YangName = "cefcFanCoolingTable"
    cefcfancoolingtable.EntityData.BundleName = "cisco_ios_xe"
    cefcfancoolingtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfancoolingtable.EntityData.SegmentPath = "cefcFanCoolingTable"
    cefcfancoolingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfancoolingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfancoolingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfancoolingtable.EntityData.Children = make(map[string]types.YChild)
    cefcfancoolingtable.EntityData.Children["cefcFanCoolingEntry"] = types.YChild{"Cefcfancoolingentry", nil}
    for i := range cefcfancoolingtable.Cefcfancoolingentry {
        cefcfancoolingtable.EntityData.Children[types.GetSegmentPath(&cefcfancoolingtable.Cefcfancoolingentry[i])] = types.YChild{"Cefcfancoolingentry", &cefcfancoolingtable.Cefcfancoolingentry[i]}
    }
    cefcfancoolingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfancoolingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfancoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingtable_Cefcfancoolingentry) GetEntityData() *types.CommonEntityData {
    cefcfancoolingentry.EntityData.YFilter = cefcfancoolingentry.YFilter
    cefcfancoolingentry.EntityData.YangName = "cefcFanCoolingEntry"
    cefcfancoolingentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfancoolingentry.EntityData.ParentYangName = "cefcFanCoolingTable"
    cefcfancoolingentry.EntityData.SegmentPath = "cefcFanCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfancoolingentry.Entphysicalindex) + "']"
    cefcfancoolingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfancoolingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfancoolingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfancoolingentry.EntityData.Children = make(map[string]types.YChild)
    cefcfancoolingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfancoolingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfancoolingentry.Entphysicalindex}
    cefcfancoolingentry.EntityData.Leafs["cefcFanCoolingCapacity"] = types.YLeaf{"Cefcfancoolingcapacity", cefcfancoolingentry.Cefcfancoolingcapacity}
    cefcfancoolingentry.EntityData.Leafs["cefcFanCoolingCapacityUnit"] = types.YLeaf{"Cefcfancoolingcapacityunit", cefcfancoolingentry.Cefcfancoolingcapacityunit}
    return &(cefcfancoolingentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable
// This table contains the cooling requirement for
// all the manageable components of type entPhysicalClass
// 'module'.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModuleCoolingEntry lists the cooling requirement for a manageable
    // components of type entPhysicalClass 'module'.  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry.
    Cefcmodulecoolingentry []CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry
}

func (cefcmodulecoolingtable *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable) GetEntityData() *types.CommonEntityData {
    cefcmodulecoolingtable.EntityData.YFilter = cefcmodulecoolingtable.YFilter
    cefcmodulecoolingtable.EntityData.YangName = "cefcModuleCoolingTable"
    cefcmodulecoolingtable.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulecoolingtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcmodulecoolingtable.EntityData.SegmentPath = "cefcModuleCoolingTable"
    cefcmodulecoolingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulecoolingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulecoolingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulecoolingtable.EntityData.Children = make(map[string]types.YChild)
    cefcmodulecoolingtable.EntityData.Children["cefcModuleCoolingEntry"] = types.YChild{"Cefcmodulecoolingentry", nil}
    for i := range cefcmodulecoolingtable.Cefcmodulecoolingentry {
        cefcmodulecoolingtable.EntityData.Children[types.GetSegmentPath(&cefcmodulecoolingtable.Cefcmodulecoolingentry[i])] = types.YChild{"Cefcmodulecoolingentry", &cefcmodulecoolingtable.Cefcmodulecoolingentry[i]}
    }
    cefcmodulecoolingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcmodulecoolingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcmodulecoolingentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulecoolingtable_Cefcmodulecoolingentry) GetEntityData() *types.CommonEntityData {
    cefcmodulecoolingentry.EntityData.YFilter = cefcmodulecoolingentry.YFilter
    cefcmodulecoolingentry.EntityData.YangName = "cefcModuleCoolingEntry"
    cefcmodulecoolingentry.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulecoolingentry.EntityData.ParentYangName = "cefcModuleCoolingTable"
    cefcmodulecoolingentry.EntityData.SegmentPath = "cefcModuleCoolingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulecoolingentry.Entphysicalindex) + "']"
    cefcmodulecoolingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulecoolingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulecoolingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulecoolingentry.EntityData.Children = make(map[string]types.YChild)
    cefcmodulecoolingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcmodulecoolingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcmodulecoolingentry.Entphysicalindex}
    cefcmodulecoolingentry.EntityData.Leafs["cefcModuleCooling"] = types.YLeaf{"Cefcmodulecooling", cefcmodulecoolingentry.Cefcmodulecooling}
    cefcmodulecoolingentry.EntityData.Leafs["cefcModuleCoolingUnit"] = types.YLeaf{"Cefcmodulecoolingunit", cefcmodulecoolingentry.Cefcmodulecoolingunit}
    return &(cefcmodulecoolingentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable
// This table contains a list of the possible cooling
// capacity modes and properties of the fans, whose 
// ENTITY-MIB entPhysicalTable entries have an 
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable struct {
    EntityData types.CommonEntityData
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

func (cefcfancoolingcaptable *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable) GetEntityData() *types.CommonEntityData {
    cefcfancoolingcaptable.EntityData.YFilter = cefcfancoolingcaptable.YFilter
    cefcfancoolingcaptable.EntityData.YangName = "cefcFanCoolingCapTable"
    cefcfancoolingcaptable.EntityData.BundleName = "cisco_ios_xe"
    cefcfancoolingcaptable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcfancoolingcaptable.EntityData.SegmentPath = "cefcFanCoolingCapTable"
    cefcfancoolingcaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfancoolingcaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfancoolingcaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfancoolingcaptable.EntityData.Children = make(map[string]types.YChild)
    cefcfancoolingcaptable.EntityData.Children["cefcFanCoolingCapEntry"] = types.YChild{"Cefcfancoolingcapentry", nil}
    for i := range cefcfancoolingcaptable.Cefcfancoolingcapentry {
        cefcfancoolingcaptable.EntityData.Children[types.GetSegmentPath(&cefcfancoolingcaptable.Cefcfancoolingcapentry[i])] = types.YChild{"Cefcfancoolingcapentry", &cefcfancoolingcaptable.Cefcfancoolingcapentry[i]}
    }
    cefcfancoolingcaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcfancoolingcaptable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcfancoolingcapentry *CISCOENTITYFRUCONTROLMIB_Cefcfancoolingcaptable_Cefcfancoolingcapentry) GetEntityData() *types.CommonEntityData {
    cefcfancoolingcapentry.EntityData.YFilter = cefcfancoolingcapentry.YFilter
    cefcfancoolingcapentry.EntityData.YangName = "cefcFanCoolingCapEntry"
    cefcfancoolingcapentry.EntityData.BundleName = "cisco_ios_xe"
    cefcfancoolingcapentry.EntityData.ParentYangName = "cefcFanCoolingCapTable"
    cefcfancoolingcapentry.EntityData.SegmentPath = "cefcFanCoolingCapEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcfancoolingcapentry.Entphysicalindex) + "']" + "[cefcFanCoolingCapIndex='" + fmt.Sprintf("%v", cefcfancoolingcapentry.Cefcfancoolingcapindex) + "']"
    cefcfancoolingcapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcfancoolingcapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcfancoolingcapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcfancoolingcapentry.EntityData.Children = make(map[string]types.YChild)
    cefcfancoolingcapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcfancoolingcapentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcfancoolingcapentry.Entphysicalindex}
    cefcfancoolingcapentry.EntityData.Leafs["cefcFanCoolingCapIndex"] = types.YLeaf{"Cefcfancoolingcapindex", cefcfancoolingcapentry.Cefcfancoolingcapindex}
    cefcfancoolingcapentry.EntityData.Leafs["cefcFanCoolingCapModeDescr"] = types.YLeaf{"Cefcfancoolingcapmodedescr", cefcfancoolingcapentry.Cefcfancoolingcapmodedescr}
    cefcfancoolingcapentry.EntityData.Leafs["cefcFanCoolingCapCapacity"] = types.YLeaf{"Cefcfancoolingcapcapacity", cefcfancoolingcapentry.Cefcfancoolingcapcapacity}
    cefcfancoolingcapentry.EntityData.Leafs["cefcFanCoolingCapCurrent"] = types.YLeaf{"Cefcfancoolingcapcurrent", cefcfancoolingcapentry.Cefcfancoolingcapcurrent}
    cefcfancoolingcapentry.EntityData.Leafs["cefcFanCoolingCapCapacityUnit"] = types.YLeaf{"Cefcfancoolingcapcapacityunit", cefcfancoolingcapentry.Cefcfancoolingcapcapacityunit}
    return &(cefcfancoolingcapentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable
// This table contains the connector power
// ratings of FRUs. 
// 
// Only components with power connector rating 
// management are listed in this table.
type CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcConnectorRatingEntry lists the power connector rating information of
    // a  component in the system.  An entry or entries are created by the agent
    // when an physical entity with connector rating  management is added to the
    // ENTITY-MIB  entPhysicalTable. An entry is deleted  by the agent at the
    // entity removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry.
    Cefcconnectorratingentry []CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry
}

func (cefcconnectorratingtable *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable) GetEntityData() *types.CommonEntityData {
    cefcconnectorratingtable.EntityData.YFilter = cefcconnectorratingtable.YFilter
    cefcconnectorratingtable.EntityData.YangName = "cefcConnectorRatingTable"
    cefcconnectorratingtable.EntityData.BundleName = "cisco_ios_xe"
    cefcconnectorratingtable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcconnectorratingtable.EntityData.SegmentPath = "cefcConnectorRatingTable"
    cefcconnectorratingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcconnectorratingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcconnectorratingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcconnectorratingtable.EntityData.Children = make(map[string]types.YChild)
    cefcconnectorratingtable.EntityData.Children["cefcConnectorRatingEntry"] = types.YChild{"Cefcconnectorratingentry", nil}
    for i := range cefcconnectorratingtable.Cefcconnectorratingentry {
        cefcconnectorratingtable.EntityData.Children[types.GetSegmentPath(&cefcconnectorratingtable.Cefcconnectorratingentry[i])] = types.YChild{"Cefcconnectorratingentry", &cefcconnectorratingtable.Cefcconnectorratingentry[i]}
    }
    cefcconnectorratingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcconnectorratingtable.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The maximum power that the component's connector can withdraw. The type is
    // interface{} with range: -1000000000..1000000000.
    Cefcconnectorrating interface{}
}

func (cefcconnectorratingentry *CISCOENTITYFRUCONTROLMIB_Cefcconnectorratingtable_Cefcconnectorratingentry) GetEntityData() *types.CommonEntityData {
    cefcconnectorratingentry.EntityData.YFilter = cefcconnectorratingentry.YFilter
    cefcconnectorratingentry.EntityData.YangName = "cefcConnectorRatingEntry"
    cefcconnectorratingentry.EntityData.BundleName = "cisco_ios_xe"
    cefcconnectorratingentry.EntityData.ParentYangName = "cefcConnectorRatingTable"
    cefcconnectorratingentry.EntityData.SegmentPath = "cefcConnectorRatingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcconnectorratingentry.Entphysicalindex) + "']"
    cefcconnectorratingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcconnectorratingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcconnectorratingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcconnectorratingentry.EntityData.Children = make(map[string]types.YChild)
    cefcconnectorratingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcconnectorratingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcconnectorratingentry.Entphysicalindex}
    cefcconnectorratingentry.EntityData.Leafs["cefcConnectorRating"] = types.YLeaf{"Cefcconnectorrating", cefcconnectorratingentry.Cefcconnectorrating}
    return &(cefcconnectorratingentry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable
// This table contains the total power consumption
// information for modules whose ENTITY-MIB 
// entPhysicalTable entries have an entPhysicalClass 
// of 'module'.
type CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModulePowerConsumptionEntry lists the total power consumption of a
    // manageable components of type entPhysicalClass 'module'.  Entries are
    // created by the agent at the system power-up or module insertion.  Entries
    // are deleted by the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry.
    Cefcmodulepowerconsumptionentry []CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry
}

func (cefcmodulepowerconsumptiontable *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable) GetEntityData() *types.CommonEntityData {
    cefcmodulepowerconsumptiontable.EntityData.YFilter = cefcmodulepowerconsumptiontable.YFilter
    cefcmodulepowerconsumptiontable.EntityData.YangName = "cefcModulePowerConsumptionTable"
    cefcmodulepowerconsumptiontable.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulepowerconsumptiontable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcmodulepowerconsumptiontable.EntityData.SegmentPath = "cefcModulePowerConsumptionTable"
    cefcmodulepowerconsumptiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulepowerconsumptiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulepowerconsumptiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulepowerconsumptiontable.EntityData.Children = make(map[string]types.YChild)
    cefcmodulepowerconsumptiontable.EntityData.Children["cefcModulePowerConsumptionEntry"] = types.YChild{"Cefcmodulepowerconsumptionentry", nil}
    for i := range cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry {
        cefcmodulepowerconsumptiontable.EntityData.Children[types.GetSegmentPath(&cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry[i])] = types.YChild{"Cefcmodulepowerconsumptionentry", &cefcmodulepowerconsumptiontable.Cefcmodulepowerconsumptionentry[i]}
    }
    cefcmodulepowerconsumptiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cefcmodulepowerconsumptiontable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cefcmodulepowerconsumptionentry *CISCOENTITYFRUCONTROLMIB_Cefcmodulepowerconsumptiontable_Cefcmodulepowerconsumptionentry) GetEntityData() *types.CommonEntityData {
    cefcmodulepowerconsumptionentry.EntityData.YFilter = cefcmodulepowerconsumptionentry.YFilter
    cefcmodulepowerconsumptionentry.EntityData.YangName = "cefcModulePowerConsumptionEntry"
    cefcmodulepowerconsumptionentry.EntityData.BundleName = "cisco_ios_xe"
    cefcmodulepowerconsumptionentry.EntityData.ParentYangName = "cefcModulePowerConsumptionTable"
    cefcmodulepowerconsumptionentry.EntityData.SegmentPath = "cefcModulePowerConsumptionEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cefcmodulepowerconsumptionentry.Entphysicalindex) + "']"
    cefcmodulepowerconsumptionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcmodulepowerconsumptionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcmodulepowerconsumptionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcmodulepowerconsumptionentry.EntityData.Children = make(map[string]types.YChild)
    cefcmodulepowerconsumptionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cefcmodulepowerconsumptionentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cefcmodulepowerconsumptionentry.Entphysicalindex}
    cefcmodulepowerconsumptionentry.EntityData.Leafs["cefcModulePowerConsumption"] = types.YLeaf{"Cefcmodulepowerconsumption", cefcmodulepowerconsumptionentry.Cefcmodulepowerconsumption}
    return &(cefcmodulepowerconsumptionentry.EntityData)
}

