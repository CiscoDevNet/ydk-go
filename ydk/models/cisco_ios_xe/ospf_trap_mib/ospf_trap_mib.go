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

    
    OspfTrapControl OSPFTRAPMIB_OspfTrapControl
}

func (oSPFTRAPMIB *OSPFTRAPMIB) GetEntityData() *types.CommonEntityData {
    oSPFTRAPMIB.EntityData.YFilter = oSPFTRAPMIB.YFilter
    oSPFTRAPMIB.EntityData.YangName = "OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.BundleName = "cisco_ios_xe"
    oSPFTRAPMIB.EntityData.ParentYangName = "OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.SegmentPath = "OSPF-TRAP-MIB:OSPF-TRAP-MIB"
    oSPFTRAPMIB.EntityData.AbsolutePath = oSPFTRAPMIB.EntityData.SegmentPath
    oSPFTRAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    oSPFTRAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    oSPFTRAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    oSPFTRAPMIB.EntityData.Children = types.NewOrderedMap()
    oSPFTRAPMIB.EntityData.Children.Append("ospfTrapControl", types.YChild{"OspfTrapControl", &oSPFTRAPMIB.OspfTrapControl})
    oSPFTRAPMIB.EntityData.Leafs = types.NewOrderedMap()

    oSPFTRAPMIB.EntityData.YListKeys = []string {}

    return &(oSPFTRAPMIB.EntityData)
}

// OSPFTRAPMIB_OspfTrapControl
type OSPFTRAPMIB_OspfTrapControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A 4-octet string serving as a bit map for the trap events defined by the
    // OSPF traps.  This object is used to enable and disable specific OSPF traps
    // where a 1 in the bit field represents enabled.  The right-most bit (least
    // significant) represents trap 0.  This object is persistent and when written
    // the entity SHOULD save the change to non-volatile storage. The type is
    // string with length: 4.
    OspfSetTrap interface{}

    // Potential types of configuration conflicts. Used by the ospfConfigError and
    // ospfConfigVirtError traps.  When the last value of a trap using this object
    // is needed, but no traps of that type have been sent, this value pertaining
    // to this object should be returned as noError. The type is
    // OspfConfigErrorType.
    OspfConfigErrorType interface{}

    // OSPF packet types.  When the last value of a trap using this object is
    // needed, but no traps of that type have been sent, this value pertaining to
    // this object should be returned as nullPacket. The type is OspfPacketType.
    OspfPacketType interface{}

    // The IP address of an inbound packet that cannot be identified by a neighbor
    // instance.  When the last value of a trap using this object is needed, but
    // no traps of that type have been sent, this value pertaining to this object
    // should be returned as 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OspfPacketSrc interface{}
}

func (ospfTrapControl *OSPFTRAPMIB_OspfTrapControl) GetEntityData() *types.CommonEntityData {
    ospfTrapControl.EntityData.YFilter = ospfTrapControl.YFilter
    ospfTrapControl.EntityData.YangName = "ospfTrapControl"
    ospfTrapControl.EntityData.BundleName = "cisco_ios_xe"
    ospfTrapControl.EntityData.ParentYangName = "OSPF-TRAP-MIB"
    ospfTrapControl.EntityData.SegmentPath = "ospfTrapControl"
    ospfTrapControl.EntityData.AbsolutePath = "OSPF-TRAP-MIB:OSPF-TRAP-MIB/" + ospfTrapControl.EntityData.SegmentPath
    ospfTrapControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospfTrapControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospfTrapControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospfTrapControl.EntityData.Children = types.NewOrderedMap()
    ospfTrapControl.EntityData.Leafs = types.NewOrderedMap()
    ospfTrapControl.EntityData.Leafs.Append("ospfSetTrap", types.YLeaf{"OspfSetTrap", ospfTrapControl.OspfSetTrap})
    ospfTrapControl.EntityData.Leafs.Append("ospfConfigErrorType", types.YLeaf{"OspfConfigErrorType", ospfTrapControl.OspfConfigErrorType})
    ospfTrapControl.EntityData.Leafs.Append("ospfPacketType", types.YLeaf{"OspfPacketType", ospfTrapControl.OspfPacketType})
    ospfTrapControl.EntityData.Leafs.Append("ospfPacketSrc", types.YLeaf{"OspfPacketSrc", ospfTrapControl.OspfPacketSrc})

    ospfTrapControl.EntityData.YListKeys = []string {}

    return &(ospfTrapControl.EntityData)
}

// OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType represents noError.
type OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType string

const (
    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_badVersion OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "badVersion"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_areaMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "areaMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_unknownNbmaNbr OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "unknownNbmaNbr"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_unknownVirtualNbr OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "unknownVirtualNbr"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_authTypeMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "authTypeMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_authFailure OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "authFailure"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_netMaskMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "netMaskMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_helloIntervalMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "helloIntervalMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_deadIntervalMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "deadIntervalMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_optionMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "optionMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_mtuMismatch OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "mtuMismatch"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_duplicateRouterId OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "duplicateRouterId"

    OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType_noError OSPFTRAPMIB_OspfTrapControl_OspfConfigErrorType = "noError"
)

// OSPFTRAPMIB_OspfTrapControl_OspfPacketType represents to this object should be returned as nullPacket.
type OSPFTRAPMIB_OspfTrapControl_OspfPacketType string

const (
    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_hello OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "hello"

    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_dbDescript OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "dbDescript"

    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_lsReq OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "lsReq"

    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_lsUpdate OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "lsUpdate"

    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_lsAck OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "lsAck"

    OSPFTRAPMIB_OspfTrapControl_OspfPacketType_nullPacket OSPFTRAPMIB_OspfTrapControl_OspfPacketType = "nullPacket"
)

