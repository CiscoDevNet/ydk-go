// This module defines textual conventions used in
// Storage Area Network technology specific mibs.
package cisco_st_tc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_st_tc"))
}

// FcPortTypes represents               to N-port Virtualizer (NPV).
type FcPortTypes string

const (
    FcPortTypes_auto FcPortTypes = "auto"

    FcPortTypes_fPort FcPortTypes = "fPort"

    FcPortTypes_flPort FcPortTypes = "flPort"

    FcPortTypes_ePort FcPortTypes = "ePort"

    FcPortTypes_bPort FcPortTypes = "bPort"

    FcPortTypes_fxPort FcPortTypes = "fxPort"

    FcPortTypes_sdPort FcPortTypes = "sdPort"

    FcPortTypes_tlPort FcPortTypes = "tlPort"

    FcPortTypes_nPort FcPortTypes = "nPort"

    FcPortTypes_nlPort FcPortTypes = "nlPort"

    FcPortTypes_nxPort FcPortTypes = "nxPort"

    FcPortTypes_tePort FcPortTypes = "tePort"

    FcPortTypes_fvPort FcPortTypes = "fvPort"

    FcPortTypes_portOperDown FcPortTypes = "portOperDown"

    FcPortTypes_stPort FcPortTypes = "stPort"

    FcPortTypes_npPort FcPortTypes = "npPort"

    FcPortTypes_tfPort FcPortTypes = "tfPort"

    FcPortTypes_tnpPort FcPortTypes = "tnpPort"
)

// FcPortTxTypes represents .
type FcPortTxTypes string

const (
    FcPortTxTypes_unknown FcPortTxTypes = "unknown"

    FcPortTxTypes_longWaveLaser FcPortTxTypes = "longWaveLaser"

    FcPortTxTypes_shortWaveLaser FcPortTxTypes = "shortWaveLaser"

    FcPortTxTypes_longWaveLaserCostReduced FcPortTxTypes = "longWaveLaserCostReduced"

    FcPortTxTypes_electrical FcPortTxTypes = "electrical"

    FcPortTxTypes_tenGigBaseSr FcPortTxTypes = "tenGigBaseSr"

    FcPortTxTypes_tenGigBaseLr FcPortTxTypes = "tenGigBaseLr"

    FcPortTxTypes_tenGigBaseEr FcPortTxTypes = "tenGigBaseEr"

    FcPortTxTypes_tenGigBaseLx4 FcPortTxTypes = "tenGigBaseLx4"

    FcPortTxTypes_tenGigBaseSw FcPortTxTypes = "tenGigBaseSw"

    FcPortTxTypes_tenGigBaseLw FcPortTxTypes = "tenGigBaseLw"

    FcPortTxTypes_tenGigBaseEw FcPortTxTypes = "tenGigBaseEw"
)

// FcPortModuleTypes represents .
type FcPortModuleTypes string

const (
    FcPortModuleTypes_unknown FcPortModuleTypes = "unknown"

    FcPortModuleTypes_other FcPortModuleTypes = "other"

    FcPortModuleTypes_gbic FcPortModuleTypes = "gbic"

    FcPortModuleTypes_embedded FcPortModuleTypes = "embedded"

    FcPortModuleTypes_glm FcPortModuleTypes = "glm"

    FcPortModuleTypes_gbicWithSerialID FcPortModuleTypes = "gbicWithSerialID"

    FcPortModuleTypes_gbicWithoutSerialID FcPortModuleTypes = "gbicWithoutSerialID"

    FcPortModuleTypes_sfpWithSerialID FcPortModuleTypes = "sfpWithSerialID"

    FcPortModuleTypes_sfpWithoutSerialID FcPortModuleTypes = "sfpWithoutSerialID"

    FcPortModuleTypes_xfp FcPortModuleTypes = "xfp"

    FcPortModuleTypes_x2Short FcPortModuleTypes = "x2Short"

    FcPortModuleTypes_x2Medium FcPortModuleTypes = "x2Medium"

    FcPortModuleTypes_x2Tall FcPortModuleTypes = "x2Tall"

    FcPortModuleTypes_xpakShort FcPortModuleTypes = "xpakShort"

    FcPortModuleTypes_xpakMedium FcPortModuleTypes = "xpakMedium"

    FcPortModuleTypes_xpakTall FcPortModuleTypes = "xpakTall"

    FcPortModuleTypes_xenpak FcPortModuleTypes = "xenpak"

    FcPortModuleTypes_sfpDwdm FcPortModuleTypes = "sfpDwdm"

    FcPortModuleTypes_qsfp FcPortModuleTypes = "qsfp"

    FcPortModuleTypes_x2Dwdm FcPortModuleTypes = "x2Dwdm"
)

// FcIfSpeed represents                      maximum of 16Gbit.
type FcIfSpeed string

const (
    FcIfSpeed_auto FcIfSpeed = "auto"

    FcIfSpeed_oneG FcIfSpeed = "oneG"

    FcIfSpeed_twoG FcIfSpeed = "twoG"

    FcIfSpeed_fourG FcIfSpeed = "fourG"

    FcIfSpeed_autoMaxTwoG FcIfSpeed = "autoMaxTwoG"

    FcIfSpeed_eightG FcIfSpeed = "eightG"

    FcIfSpeed_autoMaxFourG FcIfSpeed = "autoMaxFourG"

    FcIfSpeed_tenG FcIfSpeed = "tenG"

    FcIfSpeed_autoMaxEightG FcIfSpeed = "autoMaxEightG"

    FcIfSpeed_sixteenG FcIfSpeed = "sixteenG"

    FcIfSpeed_autoMaxSixteenG FcIfSpeed = "autoMaxSixteenG"
)

// FcAddressType represents fcid(2) - address is FCID.
type FcAddressType string

const (
    FcAddressType_wwn FcAddressType = "wwn"

    FcAddressType_fcid FcAddressType = "fcid"
)

// InterfaceOperMode represents                       to N-port Virtualizer (NPV).
type InterfaceOperMode string

const (
    InterfaceOperMode_auto InterfaceOperMode = "auto"

    InterfaceOperMode_fPort InterfaceOperMode = "fPort"

    InterfaceOperMode_flPort InterfaceOperMode = "flPort"

    InterfaceOperMode_ePort InterfaceOperMode = "ePort"

    InterfaceOperMode_bPort InterfaceOperMode = "bPort"

    InterfaceOperMode_fxPort InterfaceOperMode = "fxPort"

    InterfaceOperMode_sdPort InterfaceOperMode = "sdPort"

    InterfaceOperMode_tlPort InterfaceOperMode = "tlPort"

    InterfaceOperMode_nPort InterfaceOperMode = "nPort"

    InterfaceOperMode_nlPort InterfaceOperMode = "nlPort"

    InterfaceOperMode_nxPort InterfaceOperMode = "nxPort"

    InterfaceOperMode_tePort InterfaceOperMode = "tePort"

    InterfaceOperMode_fvPort InterfaceOperMode = "fvPort"

    InterfaceOperMode_portOperDown InterfaceOperMode = "portOperDown"

    InterfaceOperMode_stPort InterfaceOperMode = "stPort"

    InterfaceOperMode_mgmtPort InterfaceOperMode = "mgmtPort"

    InterfaceOperMode_ipsPort InterfaceOperMode = "ipsPort"

    InterfaceOperMode_evPort InterfaceOperMode = "evPort"

    InterfaceOperMode_npPort InterfaceOperMode = "npPort"

    InterfaceOperMode_tfPort InterfaceOperMode = "tfPort"

    InterfaceOperMode_tnpPort InterfaceOperMode = "tnpPort"
)

// FcIfServiceStateType represents               operational.
type FcIfServiceStateType string

const (
    FcIfServiceStateType_inService FcIfServiceStateType = "inService"

    FcIfServiceStateType_outOfService FcIfServiceStateType = "outOfService"
)

// FcIfSfpDiagLevelType represents transmit and receive power.
type FcIfSfpDiagLevelType string

const (
    FcIfSfpDiagLevelType_unknown FcIfSfpDiagLevelType = "unknown"

    FcIfSfpDiagLevelType_normal FcIfSfpDiagLevelType = "normal"

    FcIfSfpDiagLevelType_lowWarning FcIfSfpDiagLevelType = "lowWarning"

    FcIfSfpDiagLevelType_lowAlarm FcIfSfpDiagLevelType = "lowAlarm"

    FcIfSfpDiagLevelType_highWarning FcIfSfpDiagLevelType = "highWarning"

    FcIfSfpDiagLevelType_highAlarm FcIfSfpDiagLevelType = "highAlarm"
)

