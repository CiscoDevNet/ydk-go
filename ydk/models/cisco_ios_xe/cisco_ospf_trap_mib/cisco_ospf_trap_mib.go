// This MIB module describes new/modified notification
// objects/events, which are defined in the latest
// version for OSPF MIB IETF draft
// draftietf-ospf-mib-update-05.txt. Support for OSPF 
// Sham link is also added
package cisco_ospf_trap_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ospf_trap_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-OSPF-TRAP-MIB CISCO-OSPF-TRAP-MIB}", reflect.TypeOf(CISCOOSPFTRAPMIB{}))
    ydk.RegisterEntity("CISCO-OSPF-TRAP-MIB:CISCO-OSPF-TRAP-MIB", reflect.TypeOf(CISCOOSPFTRAPMIB{}))
}

// CISCOOSPFTRAPMIB
type CISCOOSPFTRAPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cospftrapcontrol CISCOOSPFTRAPMIB_Cospftrapcontrol
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetFilter() yfilter.YFilter { return cISCOOSPFTRAPMIB.YFilter }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) SetFilter(yf yfilter.YFilter) { cISCOOSPFTRAPMIB.YFilter = yf }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetGoName(yname string) string {
    if yname == "cospfTrapControl" { return "Cospftrapcontrol" }
    return ""
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetSegmentPath() string {
    return "CISCO-OSPF-TRAP-MIB:CISCO-OSPF-TRAP-MIB"
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfTrapControl" {
        return &cISCOOSPFTRAPMIB.Cospftrapcontrol
    }
    return nil
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cospfTrapControl"] = &cISCOOSPFTRAPMIB.Cospftrapcontrol
    return children
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetYangName() string { return "CISCO-OSPF-TRAP-MIB" }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) SetParent(parent types.Entity) { cISCOOSPFTRAPMIB.parent = parent }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetParent() types.Entity { return cISCOOSPFTRAPMIB.parent }

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetParentYangName() string { return "CISCO-OSPF-TRAP-MIB" }

// CISCOOSPFTRAPMIB_Cospftrapcontrol
type CISCOOSPFTRAPMIB_Cospftrapcontrol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An octet string serving as a bit  map  for the trap events defined by the
    // OSPF traps in  this MIB. This object is used to enable and   disable 
    // specific OSPF   traps   where  a  1   in  the  corresponding bit  field
    // represents  enabled. The type is map[string]bool.
    Cospfsettrap interface{}

    // Potential types of configuration conflicts. Used  by the cospfConfigError
    // and cospfConfigVirtError traps. When the last value of a trap using this
    // object is needed, but no traps of that type have been sent, this value
    // pertaining to this object should be returned as noError. The type is
    // Cospfconfigerrortype.
    Cospfconfigerrortype interface{}

    // OSPF packet types. When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as nullPacket. The type is Cospfpackettype.
    Cospfpackettype interface{}

    // The IP address of an inbound packet that can- not be identified by a
    // neighbor instance. When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfpacketsrc interface{}
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetFilter() yfilter.YFilter { return cospftrapcontrol.YFilter }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) SetFilter(yf yfilter.YFilter) { cospftrapcontrol.YFilter = yf }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetGoName(yname string) string {
    if yname == "cospfSetTrap" { return "Cospfsettrap" }
    if yname == "cospfConfigErrorType" { return "Cospfconfigerrortype" }
    if yname == "cospfPacketType" { return "Cospfpackettype" }
    if yname == "cospfPacketSrc" { return "Cospfpacketsrc" }
    return ""
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetSegmentPath() string {
    return "cospfTrapControl"
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfSetTrap"] = cospftrapcontrol.Cospfsettrap
    leafs["cospfConfigErrorType"] = cospftrapcontrol.Cospfconfigerrortype
    leafs["cospfPacketType"] = cospftrapcontrol.Cospfpackettype
    leafs["cospfPacketSrc"] = cospftrapcontrol.Cospfpacketsrc
    return leafs
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetBundleName() string { return "cisco_ios_xe" }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetYangName() string { return "cospfTrapControl" }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) SetParent(parent types.Entity) { cospftrapcontrol.parent = parent }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetParent() types.Entity { return cospftrapcontrol.parent }

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetParentYangName() string { return "CISCO-OSPF-TRAP-MIB" }

// CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype represents to this object should be returned as noError.
type CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype string

const (
    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_badVersion CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "badVersion"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_areaMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "areaMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_unknownNbmaNbr CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "unknownNbmaNbr"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_unknownVirtualNbr CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "unknownVirtualNbr"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_authTypeMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "authTypeMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_authFailure CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "authFailure"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_netMaskMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "netMaskMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_helloIntervalMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "helloIntervalMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_deadIntervalMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "deadIntervalMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_optionMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "optionMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_mtuMismatch CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "mtuMismatch"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_noError CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "noError"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype_unknownShamLinkNbr CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfconfigerrortype = "unknownShamLinkNbr"
)

// CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype represents to this object should be returned as nullPacket.
type CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype string

const (
    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_hello CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "hello"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_dbDescript CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "dbDescript"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_lsReq CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "lsReq"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_lsUpdate CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "lsUpdate"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_lsAck CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "lsAck"

    CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype_nullPacket CISCOOSPFTRAPMIB_Cospftrapcontrol_Cospfpackettype = "nullPacket"
)

