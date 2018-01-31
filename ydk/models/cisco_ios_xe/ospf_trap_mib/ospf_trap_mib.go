// The MIB module to describe traps for the OSPF
// Version 2 Protocol.
// 
// Copyright (C) The IETF Trust (2006).
// This version of this MIB module is part of
// RFC 4750;  see the RFC itself for full legal
// notices.
package ospf_trap_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ospf_trap_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:OSPF-TRAP-MIB OSPF-TRAP-MIB}", reflect.TypeOf(OSPFTRAPMIB{}))
    ydk.RegisterEntity("OSPF-TRAP-MIB:OSPF-TRAP-MIB", reflect.TypeOf(OSPFTRAPMIB{}))
}

// OSPFTRAPMIB
type OSPFTRAPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ospftrapcontrol OSPFTRAPMIB_Ospftrapcontrol
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetFilter() yfilter.YFilter { return oSPFTRAPMIB.YFilter }

func (oSPFTRAPMIB *OSPFTRAPMIB) SetFilter(yf yfilter.YFilter) { oSPFTRAPMIB.YFilter = yf }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetGoName(yname string) string {
    if yname == "ospfTrapControl" { return "Ospftrapcontrol" }
    return ""
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetSegmentPath() string {
    return "OSPF-TRAP-MIB:OSPF-TRAP-MIB"
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfTrapControl" {
        return &oSPFTRAPMIB.Ospftrapcontrol
    }
    return nil
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfTrapControl"] = &oSPFTRAPMIB.Ospftrapcontrol
    return children
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetYangName() string { return "OSPF-TRAP-MIB" }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (oSPFTRAPMIB *OSPFTRAPMIB) SetParent(parent types.Entity) { oSPFTRAPMIB.parent = parent }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetParent() types.Entity { return oSPFTRAPMIB.parent }

func (oSPFTRAPMIB *OSPFTRAPMIB) GetParentYangName() string { return "OSPF-TRAP-MIB" }

// OSPFTRAPMIB_Ospftrapcontrol
type OSPFTRAPMIB_Ospftrapcontrol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A 4-octet string serving as a bit map for the trap events defined by the
    // OSPF traps.  This object is used to enable and disable specific OSPF traps
    // where a 1 in the bit field represents enabled.  The right-most bit (least
    // significant) represents trap 0.  This object is persistent and when written
    // the entity SHOULD save the change to non-volatile storage. The type is
    // string with length: 4.
    Ospfsettrap interface{}

    // Potential types of configuration conflicts. Used by the ospfConfigError and
    // ospfConfigVirtError traps.  When the last value of a trap using this object
    // is needed, but no traps of that type have been sent, this value pertaining
    // to this object should be returned as noError. The type is
    // Ospfconfigerrortype.
    Ospfconfigerrortype interface{}

    // OSPF packet types.  When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as nullPacket. The type is Ospfpackettype.
    Ospfpackettype interface{}

    // The IP address of an inbound packet that cannot be identified by a neighbor
    // instance.  When the last value of a trap using this object is needed, but
    // no traps of that type have been sent, this value pertaining to this object
    // should be returned as 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ospfpacketsrc interface{}
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetFilter() yfilter.YFilter { return ospftrapcontrol.YFilter }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) SetFilter(yf yfilter.YFilter) { ospftrapcontrol.YFilter = yf }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetGoName(yname string) string {
    if yname == "ospfSetTrap" { return "Ospfsettrap" }
    if yname == "ospfConfigErrorType" { return "Ospfconfigerrortype" }
    if yname == "ospfPacketType" { return "Ospfpackettype" }
    if yname == "ospfPacketSrc" { return "Ospfpacketsrc" }
    return ""
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetSegmentPath() string {
    return "ospfTrapControl"
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfSetTrap"] = ospftrapcontrol.Ospfsettrap
    leafs["ospfConfigErrorType"] = ospftrapcontrol.Ospfconfigerrortype
    leafs["ospfPacketType"] = ospftrapcontrol.Ospfpackettype
    leafs["ospfPacketSrc"] = ospftrapcontrol.Ospfpacketsrc
    return leafs
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetBundleName() string { return "cisco_ios_xe" }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetYangName() string { return "ospfTrapControl" }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) SetParent(parent types.Entity) { ospftrapcontrol.parent = parent }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetParent() types.Entity { return ospftrapcontrol.parent }

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetParentYangName() string { return "OSPF-TRAP-MIB" }

// OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype represents noError.
type OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype string

const (
    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_badVersion OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "badVersion"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_areaMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "areaMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_unknownNbmaNbr OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "unknownNbmaNbr"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_unknownVirtualNbr OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "unknownVirtualNbr"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_authTypeMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "authTypeMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_authFailure OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "authFailure"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_netMaskMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "netMaskMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_helloIntervalMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "helloIntervalMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_deadIntervalMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "deadIntervalMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_optionMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "optionMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_mtuMismatch OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "mtuMismatch"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_duplicateRouterId OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "duplicateRouterId"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype_noError OSPFTRAPMIB_Ospftrapcontrol_Ospfconfigerrortype = "noError"
)

// OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype represents to this object should be returned as nullPacket.
type OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype string

const (
    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_hello OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "hello"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_dbDescript OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "dbDescript"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_lsReq OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "lsReq"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_lsUpdate OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "lsUpdate"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_lsAck OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "lsAck"

    OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype_nullPacket OSPFTRAPMIB_Ospftrapcontrol_Ospfpackettype = "nullPacket"
)

