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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cospftrapcontrol CISCOOSPFTRAPMIB_Cospftrapcontrol
}

func (cISCOOSPFTRAPMIB *CISCOOSPFTRAPMIB) GetEntityData() *types.CommonEntityData {
    cISCOOSPFTRAPMIB.EntityData.YFilter = cISCOOSPFTRAPMIB.YFilter
    cISCOOSPFTRAPMIB.EntityData.YangName = "CISCO-OSPF-TRAP-MIB"
    cISCOOSPFTRAPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOOSPFTRAPMIB.EntityData.ParentYangName = "CISCO-OSPF-TRAP-MIB"
    cISCOOSPFTRAPMIB.EntityData.SegmentPath = "CISCO-OSPF-TRAP-MIB:CISCO-OSPF-TRAP-MIB"
    cISCOOSPFTRAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOOSPFTRAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOOSPFTRAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOOSPFTRAPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOOSPFTRAPMIB.EntityData.Children["cospfTrapControl"] = types.YChild{"Cospftrapcontrol", &cISCOOSPFTRAPMIB.Cospftrapcontrol}
    cISCOOSPFTRAPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOOSPFTRAPMIB.EntityData)
}

// CISCOOSPFTRAPMIB_Cospftrapcontrol
type CISCOOSPFTRAPMIB_Cospftrapcontrol struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfpacketsrc interface{}
}

func (cospftrapcontrol *CISCOOSPFTRAPMIB_Cospftrapcontrol) GetEntityData() *types.CommonEntityData {
    cospftrapcontrol.EntityData.YFilter = cospftrapcontrol.YFilter
    cospftrapcontrol.EntityData.YangName = "cospfTrapControl"
    cospftrapcontrol.EntityData.BundleName = "cisco_ios_xe"
    cospftrapcontrol.EntityData.ParentYangName = "CISCO-OSPF-TRAP-MIB"
    cospftrapcontrol.EntityData.SegmentPath = "cospfTrapControl"
    cospftrapcontrol.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospftrapcontrol.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospftrapcontrol.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospftrapcontrol.EntityData.Children = make(map[string]types.YChild)
    cospftrapcontrol.EntityData.Leafs = make(map[string]types.YLeaf)
    cospftrapcontrol.EntityData.Leafs["cospfSetTrap"] = types.YLeaf{"Cospfsettrap", cospftrapcontrol.Cospfsettrap}
    cospftrapcontrol.EntityData.Leafs["cospfConfigErrorType"] = types.YLeaf{"Cospfconfigerrortype", cospftrapcontrol.Cospfconfigerrortype}
    cospftrapcontrol.EntityData.Leafs["cospfPacketType"] = types.YLeaf{"Cospfpackettype", cospftrapcontrol.Cospfpackettype}
    cospftrapcontrol.EntityData.Leafs["cospfPacketSrc"] = types.YLeaf{"Cospfpacketsrc", cospftrapcontrol.Cospfpacketsrc}
    return &(cospftrapcontrol.EntityData)
}

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

