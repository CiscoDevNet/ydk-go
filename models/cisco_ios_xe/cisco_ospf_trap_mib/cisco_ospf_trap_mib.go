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

    
    CospfTrapControl CISCOOSPFTRAPMIB_CospfTrapControl
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

    cISCOOSPFTRAPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOOSPFTRAPMIB.EntityData.Children.Append("cospfTrapControl", types.YChild{"CospfTrapControl", &cISCOOSPFTRAPMIB.CospfTrapControl})
    cISCOOSPFTRAPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOOSPFTRAPMIB.EntityData.YListKeys = []string {}

    return &(cISCOOSPFTRAPMIB.EntityData)
}

// CISCOOSPFTRAPMIB_CospfTrapControl
type CISCOOSPFTRAPMIB_CospfTrapControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An octet string serving as a bit  map  for the trap events defined by the
    // OSPF traps in  this MIB. This object is used to enable and   disable 
    // specific OSPF   traps   where  a  1   in  the  corresponding bit  field
    // represents  enabled. The type is map[string]bool.
    CospfSetTrap interface{}

    // Potential types of configuration conflicts. Used  by the cospfConfigError
    // and cospfConfigVirtError traps. When the last value of a trap using this
    // object is needed, but no traps of that type have been sent, this value
    // pertaining to this object should be returned as noError. The type is
    // CospfConfigErrorType.
    CospfConfigErrorType interface{}

    // OSPF packet types. When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as nullPacket. The type is CospfPacketType.
    CospfPacketType interface{}

    // The IP address of an inbound packet that can- not be identified by a
    // neighbor instance. When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CospfPacketSrc interface{}
}

func (cospfTrapControl *CISCOOSPFTRAPMIB_CospfTrapControl) GetEntityData() *types.CommonEntityData {
    cospfTrapControl.EntityData.YFilter = cospfTrapControl.YFilter
    cospfTrapControl.EntityData.YangName = "cospfTrapControl"
    cospfTrapControl.EntityData.BundleName = "cisco_ios_xe"
    cospfTrapControl.EntityData.ParentYangName = "CISCO-OSPF-TRAP-MIB"
    cospfTrapControl.EntityData.SegmentPath = "cospfTrapControl"
    cospfTrapControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfTrapControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfTrapControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfTrapControl.EntityData.Children = types.NewOrderedMap()
    cospfTrapControl.EntityData.Leafs = types.NewOrderedMap()
    cospfTrapControl.EntityData.Leafs.Append("cospfSetTrap", types.YLeaf{"CospfSetTrap", cospfTrapControl.CospfSetTrap})
    cospfTrapControl.EntityData.Leafs.Append("cospfConfigErrorType", types.YLeaf{"CospfConfigErrorType", cospfTrapControl.CospfConfigErrorType})
    cospfTrapControl.EntityData.Leafs.Append("cospfPacketType", types.YLeaf{"CospfPacketType", cospfTrapControl.CospfPacketType})
    cospfTrapControl.EntityData.Leafs.Append("cospfPacketSrc", types.YLeaf{"CospfPacketSrc", cospfTrapControl.CospfPacketSrc})

    cospfTrapControl.EntityData.YListKeys = []string {}

    return &(cospfTrapControl.EntityData)
}

// CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType represents to this object should be returned as noError.
type CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType string

const (
    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_badVersion CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "badVersion"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_areaMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "areaMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_unknownNbmaNbr CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "unknownNbmaNbr"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_unknownVirtualNbr CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "unknownVirtualNbr"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_authTypeMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "authTypeMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_authFailure CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "authFailure"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_netMaskMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "netMaskMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_helloIntervalMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "helloIntervalMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_deadIntervalMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "deadIntervalMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_optionMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "optionMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_mtuMismatch CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "mtuMismatch"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_noError CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "noError"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType_unknownShamLinkNbr CISCOOSPFTRAPMIB_CospfTrapControl_CospfConfigErrorType = "unknownShamLinkNbr"
)

// CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType represents to this object should be returned as nullPacket.
type CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType string

const (
    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_hello CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "hello"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_dbDescript CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "dbDescript"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_lsReq CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "lsReq"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_lsUpdate CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "lsUpdate"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_lsAck CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "lsAck"

    CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType_nullPacket CISCOOSPFTRAPMIB_CospfTrapControl_CospfPacketType = "nullPacket"
)

