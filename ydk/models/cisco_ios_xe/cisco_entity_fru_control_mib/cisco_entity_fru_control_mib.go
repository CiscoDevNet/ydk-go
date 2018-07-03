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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CefcFRUPower CISCOENTITYFRUCONTROLMIB_CefcFRUPower

    
    CefcMIBNotificationEnables CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables

    // This table lists the redundancy mode and the operational status of the
    // power supply groups in the system.
    CefcFRUPowerSupplyGroupTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable

    // This table lists the power-related administrative status and operational
    // status of the manageable components in the system.
    CefcFRUPowerStatusTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable

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
    CefcFRUPowerSupplyValueTable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable

    // A cefcModuleTable entry lists the operational and administrative status
    // information for ENTITY-MIB entPhysicalTable entries for manageable
    // components of type PhysicalClass module(9).
    CefcModuleTable CISCOENTITYFRUCONTROLMIB_CefcModuleTable

    // This table sparsely augments the cefcModuleTable (i.e., every row in this
    // table corresponds to a row in the cefcModuleTable but not necessarily
    // vice-versa).  A cefcIntelliModuleTable entry lists the information specific
    // to intelligent modules which cannot be provided by the cefcModuleTable.
    CefcIntelliModuleTable CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable

    // This table sparsely augments the cefcModuleTable (i.e., every row in this
    // table corresponds to a row in the cefcModuleTable but not necessarily
    // vice-versa).  A cefcModuleLocalSwitchingTable entry lists the information
    // specific to local switching capable modules which cannot be provided by the
    // cefcModuleTable.
    CefcModuleLocalSwitchingTable CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable

    // This table contains the operational status information for all ENTITY-MIB
    // entPhysicalTable entries which have  an entPhysicalClass of 'fan';
    // specifically, all   entPhysicalTable entries which represent either: one 
    // physical fan, or a single physical 'fan tray' which is a manufactured
    // (inseparable in the field) combination of  multiple fans.
    CefcFanTrayStatusTable CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable

    // This table contains one row per physical entity.
    CefcPhysicalTable CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable

    // This table contains the power input information for all the power supplies
    // that have entPhysicalTable entries with 'powerSupply' in the
    // entPhysicalClass.   The entries are created by the agent at the system
    // power-up or power supply insertion.  Entries are deleted by the agent upon
    // power supply removal.  The number of entries is determined by the number of
    // power supplies and number of power inputs on the power  supply.
    CefcPowerSupplyInputTable CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable

    // This table contains a list of possible output mode for the power supplies,
    // whose ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of
    // 'powerSupply'. It also indicate which mode is the operational mode within
    // the system.
    CefcPowerSupplyOutputTable CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable

    // This table contains the cooling capacity information of the chassis whose
    // ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of 'chassis'.
    CefcChassisCoolingTable CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable

    // This table contains the cooling capacity information of the fans whose
    // ENTITY-MIB entPhysicalTable entries have an entPhysicalClass of 'fan'.
    CefcFanCoolingTable CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable

    // This table contains the cooling requirement for all the manageable
    // components of type entPhysicalClass 'module'.
    CefcModuleCoolingTable CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable

    // This table contains a list of the possible cooling capacity modes and
    // properties of the fans, whose  ENTITY-MIB entPhysicalTable entries have an 
    // entPhysicalClass of 'fan'.
    CefcFanCoolingCapTable CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable

    // This table contains the connector power ratings of FRUs.   Only components
    // with power connector rating  management are listed in this table.
    CefcConnectorRatingTable CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable

    // This table contains the total power consumption information for modules
    // whose ENTITY-MIB  entPhysicalTable entries have an entPhysicalClass  of
    // 'module'.
    CefcModulePowerConsumptionTable CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable
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

    cISCOENTITYFRUCONTROLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPower", types.YChild{"CefcFRUPower", &cISCOENTITYFRUCONTROLMIB.CefcFRUPower})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcMIBNotificationEnables", types.YChild{"CefcMIBNotificationEnables", &cISCOENTITYFRUCONTROLMIB.CefcMIBNotificationEnables})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerSupplyGroupTable", types.YChild{"CefcFRUPowerSupplyGroupTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerSupplyGroupTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerStatusTable", types.YChild{"CefcFRUPowerStatusTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerStatusTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFRUPowerSupplyValueTable", types.YChild{"CefcFRUPowerSupplyValueTable", &cISCOENTITYFRUCONTROLMIB.CefcFRUPowerSupplyValueTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcModuleTable", types.YChild{"CefcModuleTable", &cISCOENTITYFRUCONTROLMIB.CefcModuleTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcIntelliModuleTable", types.YChild{"CefcIntelliModuleTable", &cISCOENTITYFRUCONTROLMIB.CefcIntelliModuleTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcModuleLocalSwitchingTable", types.YChild{"CefcModuleLocalSwitchingTable", &cISCOENTITYFRUCONTROLMIB.CefcModuleLocalSwitchingTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFanTrayStatusTable", types.YChild{"CefcFanTrayStatusTable", &cISCOENTITYFRUCONTROLMIB.CefcFanTrayStatusTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcPhysicalTable", types.YChild{"CefcPhysicalTable", &cISCOENTITYFRUCONTROLMIB.CefcPhysicalTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcPowerSupplyInputTable", types.YChild{"CefcPowerSupplyInputTable", &cISCOENTITYFRUCONTROLMIB.CefcPowerSupplyInputTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcPowerSupplyOutputTable", types.YChild{"CefcPowerSupplyOutputTable", &cISCOENTITYFRUCONTROLMIB.CefcPowerSupplyOutputTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcChassisCoolingTable", types.YChild{"CefcChassisCoolingTable", &cISCOENTITYFRUCONTROLMIB.CefcChassisCoolingTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFanCoolingTable", types.YChild{"CefcFanCoolingTable", &cISCOENTITYFRUCONTROLMIB.CefcFanCoolingTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcModuleCoolingTable", types.YChild{"CefcModuleCoolingTable", &cISCOENTITYFRUCONTROLMIB.CefcModuleCoolingTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcFanCoolingCapTable", types.YChild{"CefcFanCoolingCapTable", &cISCOENTITYFRUCONTROLMIB.CefcFanCoolingCapTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcConnectorRatingTable", types.YChild{"CefcConnectorRatingTable", &cISCOENTITYFRUCONTROLMIB.CefcConnectorRatingTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Children.Append("cefcModulePowerConsumptionTable", types.YChild{"CefcModulePowerConsumptionTable", &cISCOENTITYFRUCONTROLMIB.CefcModulePowerConsumptionTable})
    cISCOENTITYFRUCONTROLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYFRUCONTROLMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYFRUCONTROLMIB.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPower
type CISCOENTITYFRUCONTROLMIB_CefcFRUPower struct {
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
    CefcMaxDefaultInLinePower interface{}

    // The system will provide power to the device connecting to the FRU if the
    // device needs power, like an IP Phone. We call the providing power inline
    // power.  This MIB object controls the maximum default inline power for the
    // device connecting to the FRU in the system. The type is interface{} with
    // range: 0..4294967295. Units are miliwatts.
    CefcMaxDefaultHighInLinePower interface{}
}

func (cefcFRUPower *CISCOENTITYFRUCONTROLMIB_CefcFRUPower) GetEntityData() *types.CommonEntityData {
    cefcFRUPower.EntityData.YFilter = cefcFRUPower.YFilter
    cefcFRUPower.EntityData.YangName = "cefcFRUPower"
    cefcFRUPower.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPower.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPower.EntityData.SegmentPath = "cefcFRUPower"
    cefcFRUPower.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPower.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPower.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPower.EntityData.Children = types.NewOrderedMap()
    cefcFRUPower.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPower.EntityData.Leafs.Append("cefcMaxDefaultInLinePower", types.YLeaf{"CefcMaxDefaultInLinePower", cefcFRUPower.CefcMaxDefaultInLinePower})
    cefcFRUPower.EntityData.Leafs.Append("cefcMaxDefaultHighInLinePower", types.YLeaf{"CefcMaxDefaultHighInLinePower", cefcFRUPower.CefcMaxDefaultHighInLinePower})

    cefcFRUPower.EntityData.YListKeys = []string {}

    return &(cefcFRUPower.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables
type CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the following
    // notifications: cefcModuleStatusChange, cefcPowerStatusChange, 
    // cefcFRUInserted, cefcFRURemoved,  cefcUnrecognizedFRU and
    // cefcFanTrayStatusChange.  A false value will prevent these notifications
    // from being generated. The type is bool.
    CefcMIBEnableStatusNotification interface{}

    // This variable indicates whether the system produces the
    // cefcPowerSupplyOutputChange  notifications when the output capacity of  a
    // power supply has changed. A false value  will prevent this notification to
    // generated. The type is bool.
    CefcEnablePSOutputChangeNotif interface{}
}

func (cefcMIBNotificationEnables *CISCOENTITYFRUCONTROLMIB_CefcMIBNotificationEnables) GetEntityData() *types.CommonEntityData {
    cefcMIBNotificationEnables.EntityData.YFilter = cefcMIBNotificationEnables.YFilter
    cefcMIBNotificationEnables.EntityData.YangName = "cefcMIBNotificationEnables"
    cefcMIBNotificationEnables.EntityData.BundleName = "cisco_ios_xe"
    cefcMIBNotificationEnables.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcMIBNotificationEnables.EntityData.SegmentPath = "cefcMIBNotificationEnables"
    cefcMIBNotificationEnables.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcMIBNotificationEnables.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcMIBNotificationEnables.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcMIBNotificationEnables.EntityData.Children = types.NewOrderedMap()
    cefcMIBNotificationEnables.EntityData.Leafs = types.NewOrderedMap()
    cefcMIBNotificationEnables.EntityData.Leafs.Append("cefcMIBEnableStatusNotification", types.YLeaf{"CefcMIBEnableStatusNotification", cefcMIBNotificationEnables.CefcMIBEnableStatusNotification})
    cefcMIBNotificationEnables.EntityData.Leafs.Append("cefcEnablePSOutputChangeNotif", types.YLeaf{"CefcEnablePSOutputChangeNotif", cefcMIBNotificationEnables.CefcEnablePSOutputChangeNotif})

    cefcMIBNotificationEnables.EntityData.YListKeys = []string {}

    return &(cefcMIBNotificationEnables.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable
// This table lists the redundancy mode and the
// operational status of the power supply groups
// in the system.
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyGroupTable entry lists the desired redundancy mode,
    // the units of the power outputs and the  available and drawn current for the
    // power supply group.  Entries are created by the agent when a power supply
    // group is added to the entPhysicalTable. Entries are deleted by  the agent
    // at power supply group removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry.
    CefcFRUPowerSupplyGroupEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry
}

func (cefcFRUPowerSupplyGroupTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyGroupTable.EntityData.YFilter = cefcFRUPowerSupplyGroupTable.YFilter
    cefcFRUPowerSupplyGroupTable.EntityData.YangName = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerSupplyGroupTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerSupplyGroupTable.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerSupplyGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerSupplyGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerSupplyGroupTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupTable.EntityData.Children.Append("cefcFRUPowerSupplyGroupEntry", types.YChild{"CefcFRUPowerSupplyGroupEntry", nil})
    for i := range cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry {
        cefcFRUPowerSupplyGroupTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry[i]), types.YChild{"CefcFRUPowerSupplyGroupEntry", cefcFRUPowerSupplyGroupTable.CefcFRUPowerSupplyGroupEntry[i]})
    }
    cefcFRUPowerSupplyGroupTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerSupplyGroupTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerSupplyGroupTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry
// An cefcFRUPowerSupplyGroupTable entry lists the desired
// redundancy mode, the units of the power outputs and the 
// available and drawn current for the power supply group.
// 
// Entries are created by the agent when a power supply group
// is added to the entPhysicalTable. Entries are deleted by 
// the agent at power supply group removal.
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The administratively desired power supply redundancy mode. The type is
    // PowerRedundancyType.
    CefcPowerRedundancyMode interface{}

    // The units of primary supply to interpret cefcTotalAvailableCurrent and
    // cefcTotalDrawnCurrent as power.  For example, one 1000-watt power supply
    // could  deliver 100 amperes at 10 volts DC.  So the value of cefcPowerUnits
    // would be 'at 10 volts DC'.  cefcPowerUnits is for display purposes only.
    // The type is string.
    CefcPowerUnits interface{}

    // Total current available for FRU usage. The type is interface{} with range:
    // -1000000000..1000000000.
    CefcTotalAvailableCurrent interface{}

    // Total current drawn by powered-on FRUs. The type is interface{} with range:
    // -1000000000..1000000000.
    CefcTotalDrawnCurrent interface{}

    // The power supply redundancy operational mode. The type is
    // PowerRedundancyType.
    CefcPowerRedundancyOperMode interface{}

    // This object has the value of notApplicable(1) when
    // cefcPowerRedundancyOperMode of the instance does not have the value of
    // nonRedundant(4).  The other values explain the reason why 
    // cefcPowerRedundancyOperMode is nonRedundant(4), e.g.  unknown(2)           
    // the reason is not identified.  singleSupply(3)        There is only one
    // power supply                        in the group.  mismatchedSupplies(4) 
    // There are more than one power                        supplies in the
    // groups. However                        they are mismatched and can not     
    // work redundantly.  supplyError(5)         Some power supply or supplies    
    // does or do not working properly. The type is CefcPowerNonRedundantReason.
    CefcPowerNonRedundantReason interface{}

    // Total inline current drawn for inline operation. The type is interface{}
    // with range: -1000000000..1000000000.
    CefcTotalDrawnInlineCurrent interface{}
}

func (cefcFRUPowerSupplyGroupEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyGroupEntry.EntityData.YFilter = cefcFRUPowerSupplyGroupEntry.YFilter
    cefcFRUPowerSupplyGroupEntry.EntityData.YangName = "cefcFRUPowerSupplyGroupEntry"
    cefcFRUPowerSupplyGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerSupplyGroupEntry.EntityData.ParentYangName = "cefcFRUPowerSupplyGroupTable"
    cefcFRUPowerSupplyGroupEntry.EntityData.SegmentPath = "cefcFRUPowerSupplyGroupEntry" + types.AddKeyToken(cefcFRUPowerSupplyGroupEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerSupplyGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerSupplyGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerSupplyGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerSupplyGroupEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerSupplyGroupEntry.EntPhysicalIndex})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerRedundancyMode", types.YLeaf{"CefcPowerRedundancyMode", cefcFRUPowerSupplyGroupEntry.CefcPowerRedundancyMode})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerUnits", types.YLeaf{"CefcPowerUnits", cefcFRUPowerSupplyGroupEntry.CefcPowerUnits})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcTotalAvailableCurrent", types.YLeaf{"CefcTotalAvailableCurrent", cefcFRUPowerSupplyGroupEntry.CefcTotalAvailableCurrent})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcTotalDrawnCurrent", types.YLeaf{"CefcTotalDrawnCurrent", cefcFRUPowerSupplyGroupEntry.CefcTotalDrawnCurrent})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerRedundancyOperMode", types.YLeaf{"CefcPowerRedundancyOperMode", cefcFRUPowerSupplyGroupEntry.CefcPowerRedundancyOperMode})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcPowerNonRedundantReason", types.YLeaf{"CefcPowerNonRedundantReason", cefcFRUPowerSupplyGroupEntry.CefcPowerNonRedundantReason})
    cefcFRUPowerSupplyGroupEntry.EntityData.Leafs.Append("cefcTotalDrawnInlineCurrent", types.YLeaf{"CefcTotalDrawnInlineCurrent", cefcFRUPowerSupplyGroupEntry.CefcTotalDrawnInlineCurrent})

    cefcFRUPowerSupplyGroupEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerSupplyGroupEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason represents                        does or do not working properly.
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason string

const (
    CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason_notApplicable CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason = "notApplicable"

    CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason_unknown CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason = "unknown"

    CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason_singleSupply CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason = "singleSupply"

    CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason_mismatchedSupplies CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason = "mismatchedSupplies"

    CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason_supplyError CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyGroupTable_CefcFRUPowerSupplyGroupEntry_CefcPowerNonRedundantReason = "supplyError"
)

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable
// This table lists the power-related administrative status
// and operational status of the manageable components
// in the system.
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An cefcFRUPowerStatusTable entry lists the desired administrative status,
    // the operational status of the  power manageable component, and the current
    // required by  the component for operation.  Entries are created by the agent
    // at system power-up or  the insertion of the component.  Entries are deleted
    // by the agent at the removal of the component.  Only components with power
    // control are listed in the  table. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry.
    CefcFRUPowerStatusEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry
}

func (cefcFRUPowerStatusTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerStatusTable.EntityData.YFilter = cefcFRUPowerStatusTable.YFilter
    cefcFRUPowerStatusTable.EntityData.YangName = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerStatusTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerStatusTable.EntityData.SegmentPath = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerStatusTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerStatusTable.EntityData.Children.Append("cefcFRUPowerStatusEntry", types.YChild{"CefcFRUPowerStatusEntry", nil})
    for i := range cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry {
        cefcFRUPowerStatusTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry[i]), types.YChild{"CefcFRUPowerStatusEntry", cefcFRUPowerStatusTable.CefcFRUPowerStatusEntry[i]})
    }
    cefcFRUPowerStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerStatusTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerStatusTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // Administratively desired FRU power state. The type is PowerAdminType.
    CefcFRUPowerAdminStatus interface{}

    // Operational FRU power state. The type is PowerOperType.
    CefcFRUPowerOperStatus interface{}

    // Current supplied by the FRU (positive values) or current required to
    // operate the FRU (negative values). The type is interface{} with range:
    // -1000000000..1000000000.
    CefcFRUCurrent interface{}

    // This object indicates the set of supported power capabilities of the FRU. 
    // realTimeCurrent(0) -     cefcFRURealTimeCurrent is supported by the FRU.
    // The type is map[string]bool.
    CefcFRUPowerCapability interface{}

    // This object indicates the realtime value of current supplied by the FRU
    // (positive values) or the realtime value of current drawn by the FRU
    // (negative values). The type is interface{} with range:
    // -1000000000..1000000000.
    CefcFRURealTimeCurrent interface{}
}

func (cefcFRUPowerStatusEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerStatusTable_CefcFRUPowerStatusEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerStatusEntry.EntityData.YFilter = cefcFRUPowerStatusEntry.YFilter
    cefcFRUPowerStatusEntry.EntityData.YangName = "cefcFRUPowerStatusEntry"
    cefcFRUPowerStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerStatusEntry.EntityData.ParentYangName = "cefcFRUPowerStatusTable"
    cefcFRUPowerStatusEntry.EntityData.SegmentPath = "cefcFRUPowerStatusEntry" + types.AddKeyToken(cefcFRUPowerStatusEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerStatusEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerStatusEntry.EntPhysicalIndex})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUPowerAdminStatus", types.YLeaf{"CefcFRUPowerAdminStatus", cefcFRUPowerStatusEntry.CefcFRUPowerAdminStatus})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUPowerOperStatus", types.YLeaf{"CefcFRUPowerOperStatus", cefcFRUPowerStatusEntry.CefcFRUPowerOperStatus})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUCurrent", types.YLeaf{"CefcFRUCurrent", cefcFRUPowerStatusEntry.CefcFRUCurrent})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRUPowerCapability", types.YLeaf{"CefcFRUPowerCapability", cefcFRUPowerStatusEntry.CefcFRUPowerCapability})
    cefcFRUPowerStatusEntry.EntityData.Leafs.Append("cefcFRURealTimeCurrent", types.YLeaf{"CefcFRURealTimeCurrent", cefcFRUPowerStatusEntry.CefcFRURealTimeCurrent})

    cefcFRUPowerStatusEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerStatusEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable
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
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An cefcFRUPowerSupplyValueTable entry lists the current provided by the FRU
    // for operation.  Entries are created by the agent at system power-up or  FRU
    // insertion.  Entries are deleted by the agent at FRU removal.  Only power
    // supply FRUs are listed in the table. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry.
    CefcFRUPowerSupplyValueEntry []*CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry
}

func (cefcFRUPowerSupplyValueTable *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyValueTable.EntityData.YFilter = cefcFRUPowerSupplyValueTable.YFilter
    cefcFRUPowerSupplyValueTable.EntityData.YangName = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerSupplyValueTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFRUPowerSupplyValueTable.EntityData.SegmentPath = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerSupplyValueTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerSupplyValueTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerSupplyValueTable.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyValueTable.EntityData.Children.Append("cefcFRUPowerSupplyValueEntry", types.YChild{"CefcFRUPowerSupplyValueEntry", nil})
    for i := range cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry {
        cefcFRUPowerSupplyValueTable.EntityData.Children.Append(types.GetSegmentPath(cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry[i]), types.YChild{"CefcFRUPowerSupplyValueEntry", cefcFRUPowerSupplyValueTable.CefcFRUPowerSupplyValueEntry[i]})
    }
    cefcFRUPowerSupplyValueTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFRUPowerSupplyValueTable.EntityData.YListKeys = []string {}

    return &(cefcFRUPowerSupplyValueTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry
// An cefcFRUPowerSupplyValueTable entry lists the current
// provided by the FRU for operation.
// 
// Entries are created by the agent at system power-up or 
// FRU insertion.  Entries are deleted by the agent at FRU
// removal.
// 
// Only power supply FRUs are listed in the table.
type CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // Total current that could be supplied by the FRU (positive values) for
    // system operations. The type is interface{} with range:
    // -1000000000..1000000000.
    CefcFRUTotalSystemCurrent interface{}

    // Amount of current drawn by the FRU's in the system towards system
    // operations from this FRU. The type is interface{} with range:
    // -1000000000..1000000000.
    CefcFRUDrawnSystemCurrent interface{}

    // Total current supplied by the FRU (positive values) for inline operations.
    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUTotalInlineCurrent interface{}

    // Amount of current that is being drawn from this FRU for inline operation.
    // The type is interface{} with range: -1000000000..1000000000.
    CefcFRUDrawnInlineCurrent interface{}
}

func (cefcFRUPowerSupplyValueEntry *CISCOENTITYFRUCONTROLMIB_CefcFRUPowerSupplyValueTable_CefcFRUPowerSupplyValueEntry) GetEntityData() *types.CommonEntityData {
    cefcFRUPowerSupplyValueEntry.EntityData.YFilter = cefcFRUPowerSupplyValueEntry.YFilter
    cefcFRUPowerSupplyValueEntry.EntityData.YangName = "cefcFRUPowerSupplyValueEntry"
    cefcFRUPowerSupplyValueEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFRUPowerSupplyValueEntry.EntityData.ParentYangName = "cefcFRUPowerSupplyValueTable"
    cefcFRUPowerSupplyValueEntry.EntityData.SegmentPath = "cefcFRUPowerSupplyValueEntry" + types.AddKeyToken(cefcFRUPowerSupplyValueEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFRUPowerSupplyValueEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFRUPowerSupplyValueEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFRUPowerSupplyValueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFRUPowerSupplyValueEntry.EntityData.Children = types.NewOrderedMap()
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFRUPowerSupplyValueEntry.EntPhysicalIndex})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUTotalSystemCurrent", types.YLeaf{"CefcFRUTotalSystemCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUTotalSystemCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUDrawnSystemCurrent", types.YLeaf{"CefcFRUDrawnSystemCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUDrawnSystemCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUTotalInlineCurrent", types.YLeaf{"CefcFRUTotalInlineCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUTotalInlineCurrent})
    cefcFRUPowerSupplyValueEntry.EntityData.Leafs.Append("cefcFRUDrawnInlineCurrent", types.YLeaf{"CefcFRUDrawnInlineCurrent", cefcFRUPowerSupplyValueEntry.CefcFRUDrawnInlineCurrent})

    cefcFRUPowerSupplyValueEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFRUPowerSupplyValueEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleTable
// A cefcModuleTable entry lists the operational and
// administrative status information for ENTITY-MIB
// entPhysicalTable entries for manageable components
// of type PhysicalClass module(9).
type CISCOENTITYFRUCONTROLMIB_CefcModuleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModuleStatusTable entry lists the operational and administrative
    // status information for ENTITY-MIB entPhysicalTable entries for manageable
    // components  of type PhysicalClass module(9).  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry.
    CefcModuleEntry []*CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry
}

func (cefcModuleTable *CISCOENTITYFRUCONTROLMIB_CefcModuleTable) GetEntityData() *types.CommonEntityData {
    cefcModuleTable.EntityData.YFilter = cefcModuleTable.YFilter
    cefcModuleTable.EntityData.YangName = "cefcModuleTable"
    cefcModuleTable.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcModuleTable.EntityData.SegmentPath = "cefcModuleTable"
    cefcModuleTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleTable.EntityData.Children = types.NewOrderedMap()
    cefcModuleTable.EntityData.Children.Append("cefcModuleEntry", types.YChild{"CefcModuleEntry", nil})
    for i := range cefcModuleTable.CefcModuleEntry {
        cefcModuleTable.EntityData.Children.Append(types.GetSegmentPath(cefcModuleTable.CefcModuleEntry[i]), types.YChild{"CefcModuleEntry", cefcModuleTable.CefcModuleEntry[i]})
    }
    cefcModuleTable.EntityData.Leafs = types.NewOrderedMap()

    cefcModuleTable.EntityData.YListKeys = []string {}

    return &(cefcModuleTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry
// A cefcModuleStatusTable entry lists the operational and
// administrative status information for ENTITY-MIB
// entPhysicalTable entries for manageable components 
// of type PhysicalClass module(9).
// 
// Entries are created by the agent at the system power-up or
// module insertion.
// 
// Entries are deleted by the agent upon module removal.
type CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This object provides administrative control of the module. The type is
    // ModuleAdminType.
    CefcModuleAdminStatus interface{}

    // This object shows the module's operational state. The type is
    // ModuleOperType.
    CefcModuleOperStatus interface{}

    // This object identifies the reason for the last reset performed on the
    // module. The type is ModuleResetReasonType.
    CefcModuleResetReason interface{}

    // The value of sysUpTime at the time the cefcModuleOperStatus is changed. The
    // type is interface{} with range: 0..4294967295.
    CefcModuleStatusLastChangeTime interface{}

    // The value of sysUpTime when the configuration was most recently cleared.
    // The type is interface{} with range: 0..4294967295.
    CefcModuleLastClearConfigTime interface{}

    // A description qualifying the module reset reason specified in
    // cefcModuleResetReason.   Examples:   command xyz                 missing
    // task   switch over   watchdog timeout       etc. 
    // cefcModuleResetReasonDescription is for display purposes only. NMS
    // applications must not parse. The type is string.
    CefcModuleResetReasonDescription interface{}

    // This object displays human-readable textual string which describes the
    // cause of the last state change of the module. This object contains zero
    // length string if no meaningful reason could be provided.  Examples:
    // 'Invalid software version' 'Software download failed' 'Software version
    // mismatch' 'Module is in standby state' etc.  This object is for display
    // purposes only. NMS applications must not parse this object and take any
    // decision based on its value. The type is string.
    CefcModuleStateChangeReasonDescr interface{}

    // This object provides the up time for the module since it was last
    // re-initialized.  This object is not persistent; if a module reset, restart,
    // power off, the up time starts from zero. The type is interface{} with
    // range: 0..4294967295.
    CefcModuleUpTime interface{}
}

func (cefcModuleEntry *CISCOENTITYFRUCONTROLMIB_CefcModuleTable_CefcModuleEntry) GetEntityData() *types.CommonEntityData {
    cefcModuleEntry.EntityData.YFilter = cefcModuleEntry.YFilter
    cefcModuleEntry.EntityData.YangName = "cefcModuleEntry"
    cefcModuleEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleEntry.EntityData.ParentYangName = "cefcModuleTable"
    cefcModuleEntry.EntityData.SegmentPath = "cefcModuleEntry" + types.AddKeyToken(cefcModuleEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcModuleEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleEntry.EntityData.Children = types.NewOrderedMap()
    cefcModuleEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcModuleEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcModuleEntry.EntPhysicalIndex})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleAdminStatus", types.YLeaf{"CefcModuleAdminStatus", cefcModuleEntry.CefcModuleAdminStatus})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleOperStatus", types.YLeaf{"CefcModuleOperStatus", cefcModuleEntry.CefcModuleOperStatus})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleResetReason", types.YLeaf{"CefcModuleResetReason", cefcModuleEntry.CefcModuleResetReason})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleStatusLastChangeTime", types.YLeaf{"CefcModuleStatusLastChangeTime", cefcModuleEntry.CefcModuleStatusLastChangeTime})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleLastClearConfigTime", types.YLeaf{"CefcModuleLastClearConfigTime", cefcModuleEntry.CefcModuleLastClearConfigTime})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleResetReasonDescription", types.YLeaf{"CefcModuleResetReasonDescription", cefcModuleEntry.CefcModuleResetReasonDescription})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleStateChangeReasonDescr", types.YLeaf{"CefcModuleStateChangeReasonDescr", cefcModuleEntry.CefcModuleStateChangeReasonDescr})
    cefcModuleEntry.EntityData.Leafs.Append("cefcModuleUpTime", types.YLeaf{"CefcModuleUpTime", cefcModuleEntry.CefcModuleUpTime})

    cefcModuleEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcModuleEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable
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
type CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable struct {
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
    // CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry.
    CefcIntelliModuleEntry []*CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry
}

func (cefcIntelliModuleTable *CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable) GetEntityData() *types.CommonEntityData {
    cefcIntelliModuleTable.EntityData.YFilter = cefcIntelliModuleTable.YFilter
    cefcIntelliModuleTable.EntityData.YangName = "cefcIntelliModuleTable"
    cefcIntelliModuleTable.EntityData.BundleName = "cisco_ios_xe"
    cefcIntelliModuleTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcIntelliModuleTable.EntityData.SegmentPath = "cefcIntelliModuleTable"
    cefcIntelliModuleTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcIntelliModuleTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcIntelliModuleTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcIntelliModuleTable.EntityData.Children = types.NewOrderedMap()
    cefcIntelliModuleTable.EntityData.Children.Append("cefcIntelliModuleEntry", types.YChild{"CefcIntelliModuleEntry", nil})
    for i := range cefcIntelliModuleTable.CefcIntelliModuleEntry {
        cefcIntelliModuleTable.EntityData.Children.Append(types.GetSegmentPath(cefcIntelliModuleTable.CefcIntelliModuleEntry[i]), types.YChild{"CefcIntelliModuleEntry", cefcIntelliModuleTable.CefcIntelliModuleEntry[i]})
    }
    cefcIntelliModuleTable.EntityData.Leafs = types.NewOrderedMap()

    cefcIntelliModuleTable.EntityData.YListKeys = []string {}

    return &(cefcIntelliModuleTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The type of Internet address by which the intelligent module is reachable.
    // The type is InetAddressType.
    CefcIntelliModuleIPAddrType interface{}

    // The Internet address configured for the intelligent module. The type of
    // this address is  determined by the value of the object 
    // cefcIntelliModuleIPAddrType. The type is string with length: 0..255.
    CefcIntelliModuleIPAddr interface{}
}

func (cefcIntelliModuleEntry *CISCOENTITYFRUCONTROLMIB_CefcIntelliModuleTable_CefcIntelliModuleEntry) GetEntityData() *types.CommonEntityData {
    cefcIntelliModuleEntry.EntityData.YFilter = cefcIntelliModuleEntry.YFilter
    cefcIntelliModuleEntry.EntityData.YangName = "cefcIntelliModuleEntry"
    cefcIntelliModuleEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcIntelliModuleEntry.EntityData.ParentYangName = "cefcIntelliModuleTable"
    cefcIntelliModuleEntry.EntityData.SegmentPath = "cefcIntelliModuleEntry" + types.AddKeyToken(cefcIntelliModuleEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcIntelliModuleEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcIntelliModuleEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcIntelliModuleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcIntelliModuleEntry.EntityData.Children = types.NewOrderedMap()
    cefcIntelliModuleEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcIntelliModuleEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcIntelliModuleEntry.EntPhysicalIndex})
    cefcIntelliModuleEntry.EntityData.Leafs.Append("cefcIntelliModuleIPAddrType", types.YLeaf{"CefcIntelliModuleIPAddrType", cefcIntelliModuleEntry.CefcIntelliModuleIPAddrType})
    cefcIntelliModuleEntry.EntityData.Leafs.Append("cefcIntelliModuleIPAddr", types.YLeaf{"CefcIntelliModuleIPAddr", cefcIntelliModuleEntry.CefcIntelliModuleIPAddr})

    cefcIntelliModuleEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcIntelliModuleEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable
// This table sparsely augments the cefcModuleTable
// (i.e., every row in this table corresponds to a row in
// the cefcModuleTable but not necessarily vice-versa).
// 
// A cefcModuleLocalSwitchingTable entry lists the
// information specific to local switching capable
// modules which cannot be provided by the
// cefcModuleTable.
type CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModuleLocalSwitchingTable entry lists the information specific to a
    // local switching capable module which cannot be provided by this module's
    // corresponding instance in the cefcModuleTable. Only a module which is
    // capable of local switching has its entry here.  An entry of this table is
    // created if a module which is capable of local switching is detected by the
    // managed system.  An entry of this table is deleted if the removal of this
    // module. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry.
    CefcModuleLocalSwitchingEntry []*CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry
}

func (cefcModuleLocalSwitchingTable *CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable) GetEntityData() *types.CommonEntityData {
    cefcModuleLocalSwitchingTable.EntityData.YFilter = cefcModuleLocalSwitchingTable.YFilter
    cefcModuleLocalSwitchingTable.EntityData.YangName = "cefcModuleLocalSwitchingTable"
    cefcModuleLocalSwitchingTable.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleLocalSwitchingTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcModuleLocalSwitchingTable.EntityData.SegmentPath = "cefcModuleLocalSwitchingTable"
    cefcModuleLocalSwitchingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleLocalSwitchingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleLocalSwitchingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleLocalSwitchingTable.EntityData.Children = types.NewOrderedMap()
    cefcModuleLocalSwitchingTable.EntityData.Children.Append("cefcModuleLocalSwitchingEntry", types.YChild{"CefcModuleLocalSwitchingEntry", nil})
    for i := range cefcModuleLocalSwitchingTable.CefcModuleLocalSwitchingEntry {
        cefcModuleLocalSwitchingTable.EntityData.Children.Append(types.GetSegmentPath(cefcModuleLocalSwitchingTable.CefcModuleLocalSwitchingEntry[i]), types.YChild{"CefcModuleLocalSwitchingEntry", cefcModuleLocalSwitchingTable.CefcModuleLocalSwitchingEntry[i]})
    }
    cefcModuleLocalSwitchingTable.EntityData.Leafs = types.NewOrderedMap()

    cefcModuleLocalSwitchingTable.EntityData.YListKeys = []string {}

    return &(cefcModuleLocalSwitchingTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This object specifies the mode of local switching.  enabled(1)  - local
    // switching is enabled. disabled(2) - local switching is disabled. The type
    // is CefcModuleLocalSwitchingMode.
    CefcModuleLocalSwitchingMode interface{}
}

func (cefcModuleLocalSwitchingEntry *CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry) GetEntityData() *types.CommonEntityData {
    cefcModuleLocalSwitchingEntry.EntityData.YFilter = cefcModuleLocalSwitchingEntry.YFilter
    cefcModuleLocalSwitchingEntry.EntityData.YangName = "cefcModuleLocalSwitchingEntry"
    cefcModuleLocalSwitchingEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleLocalSwitchingEntry.EntityData.ParentYangName = "cefcModuleLocalSwitchingTable"
    cefcModuleLocalSwitchingEntry.EntityData.SegmentPath = "cefcModuleLocalSwitchingEntry" + types.AddKeyToken(cefcModuleLocalSwitchingEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcModuleLocalSwitchingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleLocalSwitchingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleLocalSwitchingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleLocalSwitchingEntry.EntityData.Children = types.NewOrderedMap()
    cefcModuleLocalSwitchingEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcModuleLocalSwitchingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcModuleLocalSwitchingEntry.EntPhysicalIndex})
    cefcModuleLocalSwitchingEntry.EntityData.Leafs.Append("cefcModuleLocalSwitchingMode", types.YLeaf{"CefcModuleLocalSwitchingMode", cefcModuleLocalSwitchingEntry.CefcModuleLocalSwitchingMode})

    cefcModuleLocalSwitchingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcModuleLocalSwitchingEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode represents disabled(2) - local switching is disabled.
type CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode string

const (
    CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode_enabled CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode = "enabled"

    CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode_disabled CISCOENTITYFRUCONTROLMIB_CefcModuleLocalSwitchingTable_CefcModuleLocalSwitchingEntry_CefcModuleLocalSwitchingMode = "disabled"
)

// CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable
// This table contains the operational status information
// for all ENTITY-MIB entPhysicalTable entries which have 
// an entPhysicalClass of 'fan'; specifically, all  
// entPhysicalTable entries which represent either: one 
// physical fan, or a single physical 'fan tray' which is a
// manufactured (inseparable in the field) combination of 
// multiple fans.
type CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable struct {
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
    // CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry.
    CefcFanTrayStatusEntry []*CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry
}

func (cefcFanTrayStatusTable *CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable) GetEntityData() *types.CommonEntityData {
    cefcFanTrayStatusTable.EntityData.YFilter = cefcFanTrayStatusTable.YFilter
    cefcFanTrayStatusTable.EntityData.YangName = "cefcFanTrayStatusTable"
    cefcFanTrayStatusTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFanTrayStatusTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFanTrayStatusTable.EntityData.SegmentPath = "cefcFanTrayStatusTable"
    cefcFanTrayStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanTrayStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanTrayStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanTrayStatusTable.EntityData.Children = types.NewOrderedMap()
    cefcFanTrayStatusTable.EntityData.Children.Append("cefcFanTrayStatusEntry", types.YChild{"CefcFanTrayStatusEntry", nil})
    for i := range cefcFanTrayStatusTable.CefcFanTrayStatusEntry {
        cefcFanTrayStatusTable.EntityData.Children.Append(types.GetSegmentPath(cefcFanTrayStatusTable.CefcFanTrayStatusEntry[i]), types.YChild{"CefcFanTrayStatusEntry", cefcFanTrayStatusTable.CefcFanTrayStatusEntry[i]})
    }
    cefcFanTrayStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFanTrayStatusTable.EntityData.YListKeys = []string {}

    return &(cefcFanTrayStatusTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The operational state of the fan or fan tray. unknown(1) - unknown. up(2) -
    // powered on. down(3) - powered down. warning(4) - partial failure, needs
    // replacement               as soon as possible. The type is
    // CefcFanTrayOperStatus.
    CefcFanTrayOperStatus interface{}
}

func (cefcFanTrayStatusEntry *CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry) GetEntityData() *types.CommonEntityData {
    cefcFanTrayStatusEntry.EntityData.YFilter = cefcFanTrayStatusEntry.YFilter
    cefcFanTrayStatusEntry.EntityData.YangName = "cefcFanTrayStatusEntry"
    cefcFanTrayStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFanTrayStatusEntry.EntityData.ParentYangName = "cefcFanTrayStatusTable"
    cefcFanTrayStatusEntry.EntityData.SegmentPath = "cefcFanTrayStatusEntry" + types.AddKeyToken(cefcFanTrayStatusEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFanTrayStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanTrayStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanTrayStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanTrayStatusEntry.EntityData.Children = types.NewOrderedMap()
    cefcFanTrayStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFanTrayStatusEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFanTrayStatusEntry.EntPhysicalIndex})
    cefcFanTrayStatusEntry.EntityData.Leafs.Append("cefcFanTrayOperStatus", types.YLeaf{"CefcFanTrayOperStatus", cefcFanTrayStatusEntry.CefcFanTrayOperStatus})

    cefcFanTrayStatusEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFanTrayStatusEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus represents              as soon as possible.
type CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus string

const (
    CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus_unknown CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus = "unknown"

    CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus_up CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus = "up"

    CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus_down CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus = "down"

    CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus_warning CISCOENTITYFRUCONTROLMIB_CefcFanTrayStatusTable_CefcFanTrayStatusEntry_CefcFanTrayOperStatus = "warning"
)

// CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable
// This table contains one row per physical entity.
type CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical entity. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry.
    CefcPhysicalEntry []*CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry
}

func (cefcPhysicalTable *CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable) GetEntityData() *types.CommonEntityData {
    cefcPhysicalTable.EntityData.YFilter = cefcPhysicalTable.YFilter
    cefcPhysicalTable.EntityData.YangName = "cefcPhysicalTable"
    cefcPhysicalTable.EntityData.BundleName = "cisco_ios_xe"
    cefcPhysicalTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcPhysicalTable.EntityData.SegmentPath = "cefcPhysicalTable"
    cefcPhysicalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPhysicalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPhysicalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPhysicalTable.EntityData.Children = types.NewOrderedMap()
    cefcPhysicalTable.EntityData.Children.Append("cefcPhysicalEntry", types.YChild{"CefcPhysicalEntry", nil})
    for i := range cefcPhysicalTable.CefcPhysicalEntry {
        cefcPhysicalTable.EntityData.Children.Append(types.GetSegmentPath(cefcPhysicalTable.CefcPhysicalEntry[i]), types.YChild{"CefcPhysicalEntry", cefcPhysicalTable.CefcPhysicalEntry[i]})
    }
    cefcPhysicalTable.EntityData.Leafs = types.NewOrderedMap()

    cefcPhysicalTable.EntityData.YListKeys = []string {}

    return &(cefcPhysicalTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry
// Information about a particular physical entity.
type CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The status of this physical entity. other(1) - the status is not any of the
    // listed below. supported(2) - this entity is supported. unsupported(3) -
    // this entity is unsupported. incompatible(4) - this entity is incompatible.
    // It would be unsupported(3), if the ID read from Serial EPROM is not
    // supported. It would be incompatible(4), if in the present configuration
    // this FRU is not supported. The type is CefcPhysicalStatus.
    CefcPhysicalStatus interface{}
}

func (cefcPhysicalEntry *CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry) GetEntityData() *types.CommonEntityData {
    cefcPhysicalEntry.EntityData.YFilter = cefcPhysicalEntry.YFilter
    cefcPhysicalEntry.EntityData.YangName = "cefcPhysicalEntry"
    cefcPhysicalEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcPhysicalEntry.EntityData.ParentYangName = "cefcPhysicalTable"
    cefcPhysicalEntry.EntityData.SegmentPath = "cefcPhysicalEntry" + types.AddKeyToken(cefcPhysicalEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcPhysicalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPhysicalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPhysicalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPhysicalEntry.EntityData.Children = types.NewOrderedMap()
    cefcPhysicalEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcPhysicalEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcPhysicalEntry.EntPhysicalIndex})
    cefcPhysicalEntry.EntityData.Leafs.Append("cefcPhysicalStatus", types.YLeaf{"CefcPhysicalStatus", cefcPhysicalEntry.CefcPhysicalStatus})

    cefcPhysicalEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcPhysicalEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus represents in the present configuration this FRU is not supported.
type CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus string

const (
    CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus_other CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus = "other"

    CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus_supported CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus = "supported"

    CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus_unsupported CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus = "unsupported"

    CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus_incompatible CISCOENTITYFRUCONTROLMIB_CefcPhysicalTable_CefcPhysicalEntry_CefcPhysicalStatus = "incompatible"
)

// CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable
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
type CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing power input management information applicable to a
    // particular power supply and input. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry.
    CefcPowerSupplyInputEntry []*CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry
}

func (cefcPowerSupplyInputTable *CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable) GetEntityData() *types.CommonEntityData {
    cefcPowerSupplyInputTable.EntityData.YFilter = cefcPowerSupplyInputTable.YFilter
    cefcPowerSupplyInputTable.EntityData.YangName = "cefcPowerSupplyInputTable"
    cefcPowerSupplyInputTable.EntityData.BundleName = "cisco_ios_xe"
    cefcPowerSupplyInputTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcPowerSupplyInputTable.EntityData.SegmentPath = "cefcPowerSupplyInputTable"
    cefcPowerSupplyInputTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPowerSupplyInputTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPowerSupplyInputTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPowerSupplyInputTable.EntityData.Children = types.NewOrderedMap()
    cefcPowerSupplyInputTable.EntityData.Children.Append("cefcPowerSupplyInputEntry", types.YChild{"CefcPowerSupplyInputEntry", nil})
    for i := range cefcPowerSupplyInputTable.CefcPowerSupplyInputEntry {
        cefcPowerSupplyInputTable.EntityData.Children.Append(types.GetSegmentPath(cefcPowerSupplyInputTable.CefcPowerSupplyInputEntry[i]), types.YChild{"CefcPowerSupplyInputEntry", cefcPowerSupplyInputTable.CefcPowerSupplyInputEntry[i]})
    }
    cefcPowerSupplyInputTable.EntityData.Leafs = types.NewOrderedMap()

    cefcPowerSupplyInputTable.EntityData.YListKeys = []string {}

    return &(cefcPowerSupplyInputTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry
// An entry containing power input management information
// applicable to a particular power supply and input.
type CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. A unique value, greater than zero, for each input
    // on a power supply. The type is interface{} with range: 0..4294967295.
    CefcPowerSupplyInputIndex interface{}

    // The type of an input power detected on the power supply.  unknown(1): No
    // input power is detected.  acLow(2): Lower rating AC input power is
    // detected.  acHigh(3): Higher rating AC input power is detected.  dcLow(4):
    // Lower rating DC input power is detected.  dcHigh(5): Higher rating DC input
    // power is detected. The type is CefcPowerSupplyInputType.
    CefcPowerSupplyInputType interface{}
}

func (cefcPowerSupplyInputEntry *CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry) GetEntityData() *types.CommonEntityData {
    cefcPowerSupplyInputEntry.EntityData.YFilter = cefcPowerSupplyInputEntry.YFilter
    cefcPowerSupplyInputEntry.EntityData.YangName = "cefcPowerSupplyInputEntry"
    cefcPowerSupplyInputEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcPowerSupplyInputEntry.EntityData.ParentYangName = "cefcPowerSupplyInputTable"
    cefcPowerSupplyInputEntry.EntityData.SegmentPath = "cefcPowerSupplyInputEntry" + types.AddKeyToken(cefcPowerSupplyInputEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefcPowerSupplyInputEntry.CefcPowerSupplyInputIndex, "cefcPowerSupplyInputIndex")
    cefcPowerSupplyInputEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPowerSupplyInputEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPowerSupplyInputEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPowerSupplyInputEntry.EntityData.Children = types.NewOrderedMap()
    cefcPowerSupplyInputEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcPowerSupplyInputEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcPowerSupplyInputEntry.EntPhysicalIndex})
    cefcPowerSupplyInputEntry.EntityData.Leafs.Append("cefcPowerSupplyInputIndex", types.YLeaf{"CefcPowerSupplyInputIndex", cefcPowerSupplyInputEntry.CefcPowerSupplyInputIndex})
    cefcPowerSupplyInputEntry.EntityData.Leafs.Append("cefcPowerSupplyInputType", types.YLeaf{"CefcPowerSupplyInputType", cefcPowerSupplyInputEntry.CefcPowerSupplyInputType})

    cefcPowerSupplyInputEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefcPowerSupplyInputIndex"}

    return &(cefcPowerSupplyInputEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType represents dcHigh(5): Higher rating DC input power is detected.
type CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType string

const (
    CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType_unknown CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType = "unknown"

    CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType_acLow CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType = "acLow"

    CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType_acHigh CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType = "acHigh"

    CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType_dcLow CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType = "dcLow"

    CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType_dcHigh CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyInputTable_CefcPowerSupplyInputEntry_CefcPowerSupplyInputType = "dcHigh"
)

// CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable
// This table contains a list of possible output
// mode for the power supplies, whose ENTITY-MIB
// entPhysicalTable entries have an entPhysicalClass
// of 'powerSupply'. It also indicate which mode
// is the operational mode within the system.
type CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcPowerSupplyOutputTable entry lists the power output capacity and its
    // operational status for manageable components of type PhysicalClass
    // 'powerSupply'.  Entries are created by the agent at the system power-up or
    // power supply insertion.  Entries are deleted by the agent upon power supply
    // removal.  The number of entries of a power supply is determined by the
    // power supply. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable_CefcPowerSupplyOutputEntry.
    CefcPowerSupplyOutputEntry []*CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable_CefcPowerSupplyOutputEntry
}

func (cefcPowerSupplyOutputTable *CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable) GetEntityData() *types.CommonEntityData {
    cefcPowerSupplyOutputTable.EntityData.YFilter = cefcPowerSupplyOutputTable.YFilter
    cefcPowerSupplyOutputTable.EntityData.YangName = "cefcPowerSupplyOutputTable"
    cefcPowerSupplyOutputTable.EntityData.BundleName = "cisco_ios_xe"
    cefcPowerSupplyOutputTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcPowerSupplyOutputTable.EntityData.SegmentPath = "cefcPowerSupplyOutputTable"
    cefcPowerSupplyOutputTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPowerSupplyOutputTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPowerSupplyOutputTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPowerSupplyOutputTable.EntityData.Children = types.NewOrderedMap()
    cefcPowerSupplyOutputTable.EntityData.Children.Append("cefcPowerSupplyOutputEntry", types.YChild{"CefcPowerSupplyOutputEntry", nil})
    for i := range cefcPowerSupplyOutputTable.CefcPowerSupplyOutputEntry {
        cefcPowerSupplyOutputTable.EntityData.Children.Append(types.GetSegmentPath(cefcPowerSupplyOutputTable.CefcPowerSupplyOutputEntry[i]), types.YChild{"CefcPowerSupplyOutputEntry", cefcPowerSupplyOutputTable.CefcPowerSupplyOutputEntry[i]})
    }
    cefcPowerSupplyOutputTable.EntityData.Leafs = types.NewOrderedMap()

    cefcPowerSupplyOutputTable.EntityData.YListKeys = []string {}

    return &(cefcPowerSupplyOutputTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable_CefcPowerSupplyOutputEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable_CefcPowerSupplyOutputEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. A unique value, greater than zero, for each
    // possible output mode on a power supply. The type is interface{} with range:
    // 0..4294967295.
    CefcPSOutputModeIndex interface{}

    // The output capacity of the power supply. The type is interface{} with
    // range: -1000000000..1000000000.
    CefcPSOutputModeCurrent interface{}

    // A value of 'true' indicates that this mode is the operational mode of the
    // power supply output capacity.  A value of 'false' indicates that this mode
    // is not the operational mode of the power supply output capacity.  For a
    // given power supply's entPhysicalIndex,  at most one instance of this object
    // can have the value of true(1). The type is bool.
    CefcPSOutputModeInOperation interface{}
}

func (cefcPowerSupplyOutputEntry *CISCOENTITYFRUCONTROLMIB_CefcPowerSupplyOutputTable_CefcPowerSupplyOutputEntry) GetEntityData() *types.CommonEntityData {
    cefcPowerSupplyOutputEntry.EntityData.YFilter = cefcPowerSupplyOutputEntry.YFilter
    cefcPowerSupplyOutputEntry.EntityData.YangName = "cefcPowerSupplyOutputEntry"
    cefcPowerSupplyOutputEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcPowerSupplyOutputEntry.EntityData.ParentYangName = "cefcPowerSupplyOutputTable"
    cefcPowerSupplyOutputEntry.EntityData.SegmentPath = "cefcPowerSupplyOutputEntry" + types.AddKeyToken(cefcPowerSupplyOutputEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefcPowerSupplyOutputEntry.CefcPSOutputModeIndex, "cefcPSOutputModeIndex")
    cefcPowerSupplyOutputEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcPowerSupplyOutputEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcPowerSupplyOutputEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcPowerSupplyOutputEntry.EntityData.Children = types.NewOrderedMap()
    cefcPowerSupplyOutputEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcPowerSupplyOutputEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcPowerSupplyOutputEntry.EntPhysicalIndex})
    cefcPowerSupplyOutputEntry.EntityData.Leafs.Append("cefcPSOutputModeIndex", types.YLeaf{"CefcPSOutputModeIndex", cefcPowerSupplyOutputEntry.CefcPSOutputModeIndex})
    cefcPowerSupplyOutputEntry.EntityData.Leafs.Append("cefcPSOutputModeCurrent", types.YLeaf{"CefcPSOutputModeCurrent", cefcPowerSupplyOutputEntry.CefcPSOutputModeCurrent})
    cefcPowerSupplyOutputEntry.EntityData.Leafs.Append("cefcPSOutputModeInOperation", types.YLeaf{"CefcPSOutputModeInOperation", cefcPowerSupplyOutputEntry.CefcPSOutputModeInOperation})

    cefcPowerSupplyOutputEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefcPSOutputModeIndex"}

    return &(cefcPowerSupplyOutputEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable
// This table contains the cooling capacity
// information of the chassis whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'chassis'.
type CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcChassisCoolingEntry lists the maximum cooling capacity that could be
    // provided  for one slot on the manageable components of type PhysicalClass
    // 'chassis'.  Entries are created by the agent if the corresponding entry is
    // created in ENTITY-MIB entPhysicalTable.  Entries are deleted by the agent
    // if the corresponding entry is deleted in ENTITY-MIB entPhysicalTable. The
    // type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable_CefcChassisCoolingEntry.
    CefcChassisCoolingEntry []*CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable_CefcChassisCoolingEntry
}

func (cefcChassisCoolingTable *CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable) GetEntityData() *types.CommonEntityData {
    cefcChassisCoolingTable.EntityData.YFilter = cefcChassisCoolingTable.YFilter
    cefcChassisCoolingTable.EntityData.YangName = "cefcChassisCoolingTable"
    cefcChassisCoolingTable.EntityData.BundleName = "cisco_ios_xe"
    cefcChassisCoolingTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcChassisCoolingTable.EntityData.SegmentPath = "cefcChassisCoolingTable"
    cefcChassisCoolingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcChassisCoolingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcChassisCoolingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcChassisCoolingTable.EntityData.Children = types.NewOrderedMap()
    cefcChassisCoolingTable.EntityData.Children.Append("cefcChassisCoolingEntry", types.YChild{"CefcChassisCoolingEntry", nil})
    for i := range cefcChassisCoolingTable.CefcChassisCoolingEntry {
        cefcChassisCoolingTable.EntityData.Children.Append(types.GetSegmentPath(cefcChassisCoolingTable.CefcChassisCoolingEntry[i]), types.YChild{"CefcChassisCoolingEntry", cefcChassisCoolingTable.CefcChassisCoolingEntry[i]})
    }
    cefcChassisCoolingTable.EntityData.Leafs = types.NewOrderedMap()

    cefcChassisCoolingTable.EntityData.YListKeys = []string {}

    return &(cefcChassisCoolingTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable_CefcChassisCoolingEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable_CefcChassisCoolingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The maximum cooling capacity that could be provided for any slot in this
    // chassis.  The default unit of the cooling capacity is 'cfm', if
    // cefcChassisPerSlotCoolingUnit is not supported. The type is interface{}
    // with range: 0..4294967295.
    CefcChassisPerSlotCoolingCap interface{}

    // The unit of the maximum cooling capacity for any slot in this chassis. The
    // type is FRUCoolingUnit.
    CefcChassisPerSlotCoolingUnit interface{}
}

func (cefcChassisCoolingEntry *CISCOENTITYFRUCONTROLMIB_CefcChassisCoolingTable_CefcChassisCoolingEntry) GetEntityData() *types.CommonEntityData {
    cefcChassisCoolingEntry.EntityData.YFilter = cefcChassisCoolingEntry.YFilter
    cefcChassisCoolingEntry.EntityData.YangName = "cefcChassisCoolingEntry"
    cefcChassisCoolingEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcChassisCoolingEntry.EntityData.ParentYangName = "cefcChassisCoolingTable"
    cefcChassisCoolingEntry.EntityData.SegmentPath = "cefcChassisCoolingEntry" + types.AddKeyToken(cefcChassisCoolingEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcChassisCoolingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcChassisCoolingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcChassisCoolingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcChassisCoolingEntry.EntityData.Children = types.NewOrderedMap()
    cefcChassisCoolingEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcChassisCoolingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcChassisCoolingEntry.EntPhysicalIndex})
    cefcChassisCoolingEntry.EntityData.Leafs.Append("cefcChassisPerSlotCoolingCap", types.YLeaf{"CefcChassisPerSlotCoolingCap", cefcChassisCoolingEntry.CefcChassisPerSlotCoolingCap})
    cefcChassisCoolingEntry.EntityData.Leafs.Append("cefcChassisPerSlotCoolingUnit", types.YLeaf{"CefcChassisPerSlotCoolingUnit", cefcChassisCoolingEntry.CefcChassisPerSlotCoolingUnit})

    cefcChassisCoolingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcChassisCoolingEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable
// This table contains the cooling capacity
// information of the fans whose ENTITY-MIB
// entPhysicalTable entries have an
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcFanCoolingEntry lists the cooling capacity that is provided by the 
    // manageable components of type PhysicalClass  'fan'.  Entries are created by
    // the agent if the corresponding entry is created in ENTITY-MIB
    // entPhysicalTable.  Entries are deleted by the agent if the corresponding
    // entry is deleted in ENTITY-MIB entPhysicalTable. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable_CefcFanCoolingEntry.
    CefcFanCoolingEntry []*CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable_CefcFanCoolingEntry
}

func (cefcFanCoolingTable *CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable) GetEntityData() *types.CommonEntityData {
    cefcFanCoolingTable.EntityData.YFilter = cefcFanCoolingTable.YFilter
    cefcFanCoolingTable.EntityData.YangName = "cefcFanCoolingTable"
    cefcFanCoolingTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFanCoolingTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFanCoolingTable.EntityData.SegmentPath = "cefcFanCoolingTable"
    cefcFanCoolingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanCoolingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanCoolingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanCoolingTable.EntityData.Children = types.NewOrderedMap()
    cefcFanCoolingTable.EntityData.Children.Append("cefcFanCoolingEntry", types.YChild{"CefcFanCoolingEntry", nil})
    for i := range cefcFanCoolingTable.CefcFanCoolingEntry {
        cefcFanCoolingTable.EntityData.Children.Append(types.GetSegmentPath(cefcFanCoolingTable.CefcFanCoolingEntry[i]), types.YChild{"CefcFanCoolingEntry", cefcFanCoolingTable.CefcFanCoolingEntry[i]})
    }
    cefcFanCoolingTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFanCoolingTable.EntityData.YListKeys = []string {}

    return &(cefcFanCoolingTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable_CefcFanCoolingEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable_CefcFanCoolingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The cooling capacity that is provided by this fan.  The default unit of the
    // fan cooling capacity is 'cfm', if cefcFanCoolingCapacityUnit is not
    // supported. The type is interface{} with range: 0..4294967295.
    CefcFanCoolingCapacity interface{}

    // The unit of the fan cooling capacity. The type is FRUCoolingUnit.
    CefcFanCoolingCapacityUnit interface{}
}

func (cefcFanCoolingEntry *CISCOENTITYFRUCONTROLMIB_CefcFanCoolingTable_CefcFanCoolingEntry) GetEntityData() *types.CommonEntityData {
    cefcFanCoolingEntry.EntityData.YFilter = cefcFanCoolingEntry.YFilter
    cefcFanCoolingEntry.EntityData.YangName = "cefcFanCoolingEntry"
    cefcFanCoolingEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFanCoolingEntry.EntityData.ParentYangName = "cefcFanCoolingTable"
    cefcFanCoolingEntry.EntityData.SegmentPath = "cefcFanCoolingEntry" + types.AddKeyToken(cefcFanCoolingEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcFanCoolingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanCoolingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanCoolingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanCoolingEntry.EntityData.Children = types.NewOrderedMap()
    cefcFanCoolingEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFanCoolingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFanCoolingEntry.EntPhysicalIndex})
    cefcFanCoolingEntry.EntityData.Leafs.Append("cefcFanCoolingCapacity", types.YLeaf{"CefcFanCoolingCapacity", cefcFanCoolingEntry.CefcFanCoolingCapacity})
    cefcFanCoolingEntry.EntityData.Leafs.Append("cefcFanCoolingCapacityUnit", types.YLeaf{"CefcFanCoolingCapacityUnit", cefcFanCoolingEntry.CefcFanCoolingCapacityUnit})

    cefcFanCoolingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcFanCoolingEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable
// This table contains the cooling requirement for
// all the manageable components of type entPhysicalClass
// 'module'.
type CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModuleCoolingEntry lists the cooling requirement for a manageable
    // components of type entPhysicalClass 'module'.  Entries are created by the
    // agent at the system power-up or module insertion.  Entries are deleted by
    // the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable_CefcModuleCoolingEntry.
    CefcModuleCoolingEntry []*CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable_CefcModuleCoolingEntry
}

func (cefcModuleCoolingTable *CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable) GetEntityData() *types.CommonEntityData {
    cefcModuleCoolingTable.EntityData.YFilter = cefcModuleCoolingTable.YFilter
    cefcModuleCoolingTable.EntityData.YangName = "cefcModuleCoolingTable"
    cefcModuleCoolingTable.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleCoolingTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcModuleCoolingTable.EntityData.SegmentPath = "cefcModuleCoolingTable"
    cefcModuleCoolingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleCoolingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleCoolingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleCoolingTable.EntityData.Children = types.NewOrderedMap()
    cefcModuleCoolingTable.EntityData.Children.Append("cefcModuleCoolingEntry", types.YChild{"CefcModuleCoolingEntry", nil})
    for i := range cefcModuleCoolingTable.CefcModuleCoolingEntry {
        cefcModuleCoolingTable.EntityData.Children.Append(types.GetSegmentPath(cefcModuleCoolingTable.CefcModuleCoolingEntry[i]), types.YChild{"CefcModuleCoolingEntry", cefcModuleCoolingTable.CefcModuleCoolingEntry[i]})
    }
    cefcModuleCoolingTable.EntityData.Leafs = types.NewOrderedMap()

    cefcModuleCoolingTable.EntityData.YListKeys = []string {}

    return &(cefcModuleCoolingTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable_CefcModuleCoolingEntry
// A cefcModuleCoolingEntry lists the cooling
// requirement for a manageable components of type
// entPhysicalClass 'module'.
// 
// Entries are created by the agent at the system
// power-up or module insertion.
// 
// Entries are deleted by the agent upon module
// removal.
type CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable_CefcModuleCoolingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The cooling requirement of the module and its daughter cards.  The default
    // unit of the module cooling requirement is 'cfm', if cefcModuleCoolingUnit
    // is not supported. The type is interface{} with range: 0..4294967295.
    CefcModuleCooling interface{}

    // The unit of the cooling requirement of the module and its daughter cards.
    // The type is FRUCoolingUnit.
    CefcModuleCoolingUnit interface{}
}

func (cefcModuleCoolingEntry *CISCOENTITYFRUCONTROLMIB_CefcModuleCoolingTable_CefcModuleCoolingEntry) GetEntityData() *types.CommonEntityData {
    cefcModuleCoolingEntry.EntityData.YFilter = cefcModuleCoolingEntry.YFilter
    cefcModuleCoolingEntry.EntityData.YangName = "cefcModuleCoolingEntry"
    cefcModuleCoolingEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcModuleCoolingEntry.EntityData.ParentYangName = "cefcModuleCoolingTable"
    cefcModuleCoolingEntry.EntityData.SegmentPath = "cefcModuleCoolingEntry" + types.AddKeyToken(cefcModuleCoolingEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcModuleCoolingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModuleCoolingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModuleCoolingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModuleCoolingEntry.EntityData.Children = types.NewOrderedMap()
    cefcModuleCoolingEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcModuleCoolingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcModuleCoolingEntry.EntPhysicalIndex})
    cefcModuleCoolingEntry.EntityData.Leafs.Append("cefcModuleCooling", types.YLeaf{"CefcModuleCooling", cefcModuleCoolingEntry.CefcModuleCooling})
    cefcModuleCoolingEntry.EntityData.Leafs.Append("cefcModuleCoolingUnit", types.YLeaf{"CefcModuleCoolingUnit", cefcModuleCoolingEntry.CefcModuleCoolingUnit})

    cefcModuleCoolingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcModuleCoolingEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable
// This table contains a list of the possible cooling
// capacity modes and properties of the fans, whose 
// ENTITY-MIB entPhysicalTable entries have an 
// entPhysicalClass of 'fan'.
type CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcFanCoolingCapacityEntry lists the cooling capacity mode of a
    // manageable components of type entPhysicalClass 'fan'. It also lists the
    // corresponding cooling capacity provided and the power consumed by the fan
    // on this mode.   Entries are created by the agent if the corresponding entry
    // is created in ENTITY-MIB entPhysicalTable.  Entries are deleted by the
    // agent if the corresponding entry is deleted in ENTITY-MIB entPhysicalTable.
    // The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable_CefcFanCoolingCapEntry.
    CefcFanCoolingCapEntry []*CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable_CefcFanCoolingCapEntry
}

func (cefcFanCoolingCapTable *CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable) GetEntityData() *types.CommonEntityData {
    cefcFanCoolingCapTable.EntityData.YFilter = cefcFanCoolingCapTable.YFilter
    cefcFanCoolingCapTable.EntityData.YangName = "cefcFanCoolingCapTable"
    cefcFanCoolingCapTable.EntityData.BundleName = "cisco_ios_xe"
    cefcFanCoolingCapTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcFanCoolingCapTable.EntityData.SegmentPath = "cefcFanCoolingCapTable"
    cefcFanCoolingCapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanCoolingCapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanCoolingCapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanCoolingCapTable.EntityData.Children = types.NewOrderedMap()
    cefcFanCoolingCapTable.EntityData.Children.Append("cefcFanCoolingCapEntry", types.YChild{"CefcFanCoolingCapEntry", nil})
    for i := range cefcFanCoolingCapTable.CefcFanCoolingCapEntry {
        cefcFanCoolingCapTable.EntityData.Children.Append(types.GetSegmentPath(cefcFanCoolingCapTable.CefcFanCoolingCapEntry[i]), types.YChild{"CefcFanCoolingCapEntry", cefcFanCoolingCapTable.CefcFanCoolingCapEntry[i]})
    }
    cefcFanCoolingCapTable.EntityData.Leafs = types.NewOrderedMap()

    cefcFanCoolingCapTable.EntityData.YListKeys = []string {}

    return &(cefcFanCoolingCapTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable_CefcFanCoolingCapEntry
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
type CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable_CefcFanCoolingCapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. An arbitrary value that uniquely identifies a
    // cooling capacity mode for a fan. The type is interface{} with range:
    // 1..4095.
    CefcFanCoolingCapIndex interface{}

    // A textual description of the cooling capacity mode of the fan. The type is
    // string.
    CefcFanCoolingCapModeDescr interface{}

    // The cooling capacity that could be provided when the fan is operating in
    // this mode.  The default unit of the cooling capacity is 'cfm', if
    // cefcFanCoolingCapCapacityUnit is not supported. The type is interface{}
    // with range: 0..4294967295.
    CefcFanCoolingCapCapacity interface{}

    // The power consumption of the fan when operating in in this mode. The type
    // is interface{} with range: -1000000000..1000000000.
    CefcFanCoolingCapCurrent interface{}

    // The unit of the fan cooling capacity when operating in this mode. The type
    // is FRUCoolingUnit.
    CefcFanCoolingCapCapacityUnit interface{}
}

func (cefcFanCoolingCapEntry *CISCOENTITYFRUCONTROLMIB_CefcFanCoolingCapTable_CefcFanCoolingCapEntry) GetEntityData() *types.CommonEntityData {
    cefcFanCoolingCapEntry.EntityData.YFilter = cefcFanCoolingCapEntry.YFilter
    cefcFanCoolingCapEntry.EntityData.YangName = "cefcFanCoolingCapEntry"
    cefcFanCoolingCapEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcFanCoolingCapEntry.EntityData.ParentYangName = "cefcFanCoolingCapTable"
    cefcFanCoolingCapEntry.EntityData.SegmentPath = "cefcFanCoolingCapEntry" + types.AddKeyToken(cefcFanCoolingCapEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cefcFanCoolingCapEntry.CefcFanCoolingCapIndex, "cefcFanCoolingCapIndex")
    cefcFanCoolingCapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcFanCoolingCapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcFanCoolingCapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcFanCoolingCapEntry.EntityData.Children = types.NewOrderedMap()
    cefcFanCoolingCapEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcFanCoolingCapEntry.EntPhysicalIndex})
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("cefcFanCoolingCapIndex", types.YLeaf{"CefcFanCoolingCapIndex", cefcFanCoolingCapEntry.CefcFanCoolingCapIndex})
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("cefcFanCoolingCapModeDescr", types.YLeaf{"CefcFanCoolingCapModeDescr", cefcFanCoolingCapEntry.CefcFanCoolingCapModeDescr})
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("cefcFanCoolingCapCapacity", types.YLeaf{"CefcFanCoolingCapCapacity", cefcFanCoolingCapEntry.CefcFanCoolingCapCapacity})
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("cefcFanCoolingCapCurrent", types.YLeaf{"CefcFanCoolingCapCurrent", cefcFanCoolingCapEntry.CefcFanCoolingCapCurrent})
    cefcFanCoolingCapEntry.EntityData.Leafs.Append("cefcFanCoolingCapCapacityUnit", types.YLeaf{"CefcFanCoolingCapCapacityUnit", cefcFanCoolingCapEntry.CefcFanCoolingCapCapacityUnit})

    cefcFanCoolingCapEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CefcFanCoolingCapIndex"}

    return &(cefcFanCoolingCapEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable
// This table contains the connector power
// ratings of FRUs. 
// 
// Only components with power connector rating 
// management are listed in this table.
type CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcConnectorRatingEntry lists the power connector rating information of
    // a  component in the system.  An entry or entries are created by the agent
    // when an physical entity with connector rating  management is added to the
    // ENTITY-MIB  entPhysicalTable. An entry is deleted  by the agent at the
    // entity removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable_CefcConnectorRatingEntry.
    CefcConnectorRatingEntry []*CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable_CefcConnectorRatingEntry
}

func (cefcConnectorRatingTable *CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable) GetEntityData() *types.CommonEntityData {
    cefcConnectorRatingTable.EntityData.YFilter = cefcConnectorRatingTable.YFilter
    cefcConnectorRatingTable.EntityData.YangName = "cefcConnectorRatingTable"
    cefcConnectorRatingTable.EntityData.BundleName = "cisco_ios_xe"
    cefcConnectorRatingTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcConnectorRatingTable.EntityData.SegmentPath = "cefcConnectorRatingTable"
    cefcConnectorRatingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcConnectorRatingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcConnectorRatingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcConnectorRatingTable.EntityData.Children = types.NewOrderedMap()
    cefcConnectorRatingTable.EntityData.Children.Append("cefcConnectorRatingEntry", types.YChild{"CefcConnectorRatingEntry", nil})
    for i := range cefcConnectorRatingTable.CefcConnectorRatingEntry {
        cefcConnectorRatingTable.EntityData.Children.Append(types.GetSegmentPath(cefcConnectorRatingTable.CefcConnectorRatingEntry[i]), types.YChild{"CefcConnectorRatingEntry", cefcConnectorRatingTable.CefcConnectorRatingEntry[i]})
    }
    cefcConnectorRatingTable.EntityData.Leafs = types.NewOrderedMap()

    cefcConnectorRatingTable.EntityData.YListKeys = []string {}

    return &(cefcConnectorRatingTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable_CefcConnectorRatingEntry
// A cefcConnectorRatingEntry lists the
// power connector rating information of a 
// component in the system.
// 
// An entry or entries are created by the agent
// when an physical entity with connector rating 
// management is added to the ENTITY-MIB 
// entPhysicalTable. An entry is deleted 
// by the agent at the entity removal.
type CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable_CefcConnectorRatingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The maximum power that the component's connector can withdraw. The type is
    // interface{} with range: -1000000000..1000000000.
    CefcConnectorRating interface{}
}

func (cefcConnectorRatingEntry *CISCOENTITYFRUCONTROLMIB_CefcConnectorRatingTable_CefcConnectorRatingEntry) GetEntityData() *types.CommonEntityData {
    cefcConnectorRatingEntry.EntityData.YFilter = cefcConnectorRatingEntry.YFilter
    cefcConnectorRatingEntry.EntityData.YangName = "cefcConnectorRatingEntry"
    cefcConnectorRatingEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcConnectorRatingEntry.EntityData.ParentYangName = "cefcConnectorRatingTable"
    cefcConnectorRatingEntry.EntityData.SegmentPath = "cefcConnectorRatingEntry" + types.AddKeyToken(cefcConnectorRatingEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcConnectorRatingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcConnectorRatingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcConnectorRatingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcConnectorRatingEntry.EntityData.Children = types.NewOrderedMap()
    cefcConnectorRatingEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcConnectorRatingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcConnectorRatingEntry.EntPhysicalIndex})
    cefcConnectorRatingEntry.EntityData.Leafs.Append("cefcConnectorRating", types.YLeaf{"CefcConnectorRating", cefcConnectorRatingEntry.CefcConnectorRating})

    cefcConnectorRatingEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcConnectorRatingEntry.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable
// This table contains the total power consumption
// information for modules whose ENTITY-MIB 
// entPhysicalTable entries have an entPhysicalClass 
// of 'module'.
type CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cefcModulePowerConsumptionEntry lists the total power consumption of a
    // manageable components of type entPhysicalClass 'module'.  Entries are
    // created by the agent at the system power-up or module insertion.  Entries
    // are deleted by the agent upon module removal. The type is slice of
    // CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable_CefcModulePowerConsumptionEntry.
    CefcModulePowerConsumptionEntry []*CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable_CefcModulePowerConsumptionEntry
}

func (cefcModulePowerConsumptionTable *CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable) GetEntityData() *types.CommonEntityData {
    cefcModulePowerConsumptionTable.EntityData.YFilter = cefcModulePowerConsumptionTable.YFilter
    cefcModulePowerConsumptionTable.EntityData.YangName = "cefcModulePowerConsumptionTable"
    cefcModulePowerConsumptionTable.EntityData.BundleName = "cisco_ios_xe"
    cefcModulePowerConsumptionTable.EntityData.ParentYangName = "CISCO-ENTITY-FRU-CONTROL-MIB"
    cefcModulePowerConsumptionTable.EntityData.SegmentPath = "cefcModulePowerConsumptionTable"
    cefcModulePowerConsumptionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModulePowerConsumptionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModulePowerConsumptionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModulePowerConsumptionTable.EntityData.Children = types.NewOrderedMap()
    cefcModulePowerConsumptionTable.EntityData.Children.Append("cefcModulePowerConsumptionEntry", types.YChild{"CefcModulePowerConsumptionEntry", nil})
    for i := range cefcModulePowerConsumptionTable.CefcModulePowerConsumptionEntry {
        cefcModulePowerConsumptionTable.EntityData.Children.Append(types.GetSegmentPath(cefcModulePowerConsumptionTable.CefcModulePowerConsumptionEntry[i]), types.YChild{"CefcModulePowerConsumptionEntry", cefcModulePowerConsumptionTable.CefcModulePowerConsumptionEntry[i]})
    }
    cefcModulePowerConsumptionTable.EntityData.Leafs = types.NewOrderedMap()

    cefcModulePowerConsumptionTable.EntityData.YListKeys = []string {}

    return &(cefcModulePowerConsumptionTable.EntityData)
}

// CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable_CefcModulePowerConsumptionEntry
// A cefcModulePowerConsumptionEntry lists the total
// power consumption of a manageable components of type
// entPhysicalClass 'module'.
// 
// Entries are created by the agent at the system
// power-up or module insertion.
// 
// Entries are deleted by the agent upon module
// removal.
type CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable_CefcModulePowerConsumptionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The combined power consumption to operate the module and its submodule(s)
    // and inline-power device(s). The type is interface{} with range:
    // -1000000000..1000000000.
    CefcModulePowerConsumption interface{}
}

func (cefcModulePowerConsumptionEntry *CISCOENTITYFRUCONTROLMIB_CefcModulePowerConsumptionTable_CefcModulePowerConsumptionEntry) GetEntityData() *types.CommonEntityData {
    cefcModulePowerConsumptionEntry.EntityData.YFilter = cefcModulePowerConsumptionEntry.YFilter
    cefcModulePowerConsumptionEntry.EntityData.YangName = "cefcModulePowerConsumptionEntry"
    cefcModulePowerConsumptionEntry.EntityData.BundleName = "cisco_ios_xe"
    cefcModulePowerConsumptionEntry.EntityData.ParentYangName = "cefcModulePowerConsumptionTable"
    cefcModulePowerConsumptionEntry.EntityData.SegmentPath = "cefcModulePowerConsumptionEntry" + types.AddKeyToken(cefcModulePowerConsumptionEntry.EntPhysicalIndex, "entPhysicalIndex")
    cefcModulePowerConsumptionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cefcModulePowerConsumptionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cefcModulePowerConsumptionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cefcModulePowerConsumptionEntry.EntityData.Children = types.NewOrderedMap()
    cefcModulePowerConsumptionEntry.EntityData.Leafs = types.NewOrderedMap()
    cefcModulePowerConsumptionEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cefcModulePowerConsumptionEntry.EntPhysicalIndex})
    cefcModulePowerConsumptionEntry.EntityData.Leafs.Append("cefcModulePowerConsumption", types.YLeaf{"CefcModulePowerConsumption", cefcModulePowerConsumptionEntry.CefcModulePowerConsumption})

    cefcModulePowerConsumptionEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(cefcModulePowerConsumptionEntry.EntityData)
}

