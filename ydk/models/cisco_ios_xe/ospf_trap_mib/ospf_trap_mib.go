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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ospftrapcontrol OSPFTRAPMIB_Ospftrapcontrol
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetEntityData() *types.CommonEntityData {
    oSPFTRAPMIB.EntityData.YFilter = oSPFTRAPMIB.YFilter
    oSPFTRAPMIB.EntityData.YangName = "OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.BundleName = "cisco_ios_xe"
    oSPFTRAPMIB.EntityData.ParentYangName = "OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.SegmentPath = "OSPF-TRAP-MIB:OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    oSPFTRAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    oSPFTRAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    oSPFTRAPMIB.EntityData.Children = make(map[string]types.YChild)
    oSPFTRAPMIB.EntityData.Children["ospfTrapControl"] = types.YChild{"Ospftrapcontrol", &oSPFTRAPMIB.Ospftrapcontrol}
    oSPFTRAPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oSPFTRAPMIB.EntityData)
}

// OSPFTRAPMIB_Ospftrapcontrol
type OSPFTRAPMIB_Ospftrapcontrol struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ospfpacketsrc interface{}
}

func (ospftrapcontrol *OSPFTRAPMIB_Ospftrapcontrol) GetEntityData() *types.CommonEntityData {
    ospftrapcontrol.EntityData.YFilter = ospftrapcontrol.YFilter
    ospftrapcontrol.EntityData.YangName = "ospfTrapControl"
    ospftrapcontrol.EntityData.BundleName = "cisco_ios_xe"
    ospftrapcontrol.EntityData.ParentYangName = "OSPF-TRAP-MIB"
    ospftrapcontrol.EntityData.SegmentPath = "ospfTrapControl"
    ospftrapcontrol.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospftrapcontrol.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospftrapcontrol.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospftrapcontrol.EntityData.Children = make(map[string]types.YChild)
    ospftrapcontrol.EntityData.Leafs = make(map[string]types.YLeaf)
    ospftrapcontrol.EntityData.Leafs["ospfSetTrap"] = types.YLeaf{"Ospfsettrap", ospftrapcontrol.Ospfsettrap}
    ospftrapcontrol.EntityData.Leafs["ospfConfigErrorType"] = types.YLeaf{"Ospfconfigerrortype", ospftrapcontrol.Ospfconfigerrortype}
    ospftrapcontrol.EntityData.Leafs["ospfPacketType"] = types.YLeaf{"Ospfpackettype", ospftrapcontrol.Ospfpackettype}
    ospftrapcontrol.EntityData.Leafs["ospfPacketSrc"] = types.YLeaf{"Ospfpacketsrc", ospftrapcontrol.Ospfpacketsrc}
    return &(ospftrapcontrol.EntityData)
}

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

