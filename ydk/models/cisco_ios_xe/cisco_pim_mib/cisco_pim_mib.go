// This MIB module defines the cisco specific variables
// for Protocol Independent Multicast (PIM) management. 
// These definitions are an extension of those defined in
// the IETF PIM MIB (RFC 2934).
package cisco_pim_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_pim_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-PIM-MIB CISCO-PIM-MIB}", reflect.TypeOf(CISCOPIMMIB{}))
    ydk.RegisterEntity("CISCO-PIM-MIB:CISCO-PIM-MIB", reflect.TypeOf(CISCOPIMMIB{}))
}

// CISCOPIMMIB
type CISCOPIMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cpim CISCOPIMMIB_Cpim

    
    CiscoPimMIBNotificationObjects CISCOPIMMIB_CiscoPimMIBNotificationObjects
}

func (cISCOPIMMIB *CISCOPIMMIB) GetEntityData() *types.CommonEntityData {
    cISCOPIMMIB.EntityData.YFilter = cISCOPIMMIB.YFilter
    cISCOPIMMIB.EntityData.YangName = "CISCO-PIM-MIB"
    cISCOPIMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPIMMIB.EntityData.ParentYangName = "CISCO-PIM-MIB"
    cISCOPIMMIB.EntityData.SegmentPath = "CISCO-PIM-MIB:CISCO-PIM-MIB"
    cISCOPIMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPIMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPIMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPIMMIB.EntityData.Children = types.NewOrderedMap()
    cISCOPIMMIB.EntityData.Children.Append("cpim", types.YChild{"Cpim", &cISCOPIMMIB.Cpim})
    cISCOPIMMIB.EntityData.Children.Append("ciscoPimMIBNotificationObjects", types.YChild{"CiscoPimMIBNotificationObjects", &cISCOPIMMIB.CiscoPimMIBNotificationObjects})
    cISCOPIMMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOPIMMIB.EntityData.YListKeys = []string {}

    return &(cISCOPIMMIB.EntityData)
}

// CISCOPIMMIB_Cpim
type CISCOPIMMIB_Cpim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A count of the number of invalid PIM Register messages received by this
    // device. A PIM Register message is termed invalid if  o the encapsulated IP
    // header is malformed, o the destination of the PIM Register message is not
    // the   RP (Rendezvous Point) for the group in question, o the source/DR
    // (Designated Router) address is not a valid   unicast address. The type is
    // interface{} with range: 0..4294967295.
    CpimInvalidRegisterMsgsRcvd interface{}

    // A count of the number of invalid PIM Join/Prune messages received by this
    // device. A PIM Join/Prune message is termed invalid if  o the RP specified
    // in the packet is not the RP for   the group in question. The type is
    // interface{} with range: 0..4294967295.
    CpimInvalidJoinPruneMsgsRcvd interface{}

    // The type of the last invalid message that was received by this device. The
    // type is CpimLastErrorType.
    CpimLastErrorType interface{}

    // Represents the type of address stored in cpimLastErrorOrigin. The value of
    // this object is irrelevant if the value of cpimLastErrorType is none(1). The
    // type is InetAddressType.
    CpimLastErrorOriginType interface{}

    // This object represents the Network Layer Address of the source that
    // originated the last invalid packet. The type of address stored depends on
    // the value in cpimLastErrorOriginType.   The value of this object represents
    // the Network Layer Address of the Designated Router (DR) whenever the value
    // of cpimLastErrorGroup is a zero-length address,  for eg. when encapsulated
    // IP header is malformed.  The value of this object is irrelevant if the
    // value of cpimLastErrorType is none(1). The type is string with length:
    // 0..255.
    CpimLastErrorOrigin interface{}

    // Represents the type of address stored in cpimLastErrorGroup. The value of
    // this object is unknown(0) if there is a problem in the packet received from
    // the DR.  The value of this object is irrelevant if the value of
    // cpimLastErrorType is none(1). The type is InetAddressType.
    CpimLastErrorGroupType interface{}

    // The IP multicast group address to which the last invalid packet was
    // addressed. The type of address stored depends on the value in
    // cpimLastErrorGroupType.   The value of this object is a zero-length
    // InetAddress if there is a problem in the packet received from the DR, for
    // eg. a malformed encapsulated IP header.   The value of this object is
    // irrelevant if the value of cpimLastErrorType is none(1). The type is string
    // with length: 0..255.
    CpimLastErrorGroup interface{}

    // Represents the type of address stored in cpimLastErrorRP. The value of this
    // object is irrelevant if the value of cpimLastErrorType is none(1). The type
    // is InetAddressType.
    CpimLastErrorRPType interface{}

    // The address of the RP, as per the last invalid packet. The type of address
    // stored depends on the value in cpimLastErrorRPType.   The value of this
    // object is irrelevant if the value of cpimLastErrorType is none(1). The type
    // is string with length: 0..255.
    CpimLastErrorRP interface{}
}

func (cpim *CISCOPIMMIB_Cpim) GetEntityData() *types.CommonEntityData {
    cpim.EntityData.YFilter = cpim.YFilter
    cpim.EntityData.YangName = "cpim"
    cpim.EntityData.BundleName = "cisco_ios_xe"
    cpim.EntityData.ParentYangName = "CISCO-PIM-MIB"
    cpim.EntityData.SegmentPath = "cpim"
    cpim.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpim.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpim.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpim.EntityData.Children = types.NewOrderedMap()
    cpim.EntityData.Leafs = types.NewOrderedMap()
    cpim.EntityData.Leafs.Append("cpimInvalidRegisterMsgsRcvd", types.YLeaf{"CpimInvalidRegisterMsgsRcvd", cpim.CpimInvalidRegisterMsgsRcvd})
    cpim.EntityData.Leafs.Append("cpimInvalidJoinPruneMsgsRcvd", types.YLeaf{"CpimInvalidJoinPruneMsgsRcvd", cpim.CpimInvalidJoinPruneMsgsRcvd})
    cpim.EntityData.Leafs.Append("cpimLastErrorType", types.YLeaf{"CpimLastErrorType", cpim.CpimLastErrorType})
    cpim.EntityData.Leafs.Append("cpimLastErrorOriginType", types.YLeaf{"CpimLastErrorOriginType", cpim.CpimLastErrorOriginType})
    cpim.EntityData.Leafs.Append("cpimLastErrorOrigin", types.YLeaf{"CpimLastErrorOrigin", cpim.CpimLastErrorOrigin})
    cpim.EntityData.Leafs.Append("cpimLastErrorGroupType", types.YLeaf{"CpimLastErrorGroupType", cpim.CpimLastErrorGroupType})
    cpim.EntityData.Leafs.Append("cpimLastErrorGroup", types.YLeaf{"CpimLastErrorGroup", cpim.CpimLastErrorGroup})
    cpim.EntityData.Leafs.Append("cpimLastErrorRPType", types.YLeaf{"CpimLastErrorRPType", cpim.CpimLastErrorRPType})
    cpim.EntityData.Leafs.Append("cpimLastErrorRP", types.YLeaf{"CpimLastErrorRP", cpim.CpimLastErrorRP})

    cpim.EntityData.YListKeys = []string {}

    return &(cpim.EntityData)
}

// CISCOPIMMIB_Cpim_CpimLastErrorType represents this device.
type CISCOPIMMIB_Cpim_CpimLastErrorType string

const (
    CISCOPIMMIB_Cpim_CpimLastErrorType_none CISCOPIMMIB_Cpim_CpimLastErrorType = "none"

    CISCOPIMMIB_Cpim_CpimLastErrorType_invalidRegister CISCOPIMMIB_Cpim_CpimLastErrorType = "invalidRegister"

    CISCOPIMMIB_Cpim_CpimLastErrorType_invalidJoinPrune CISCOPIMMIB_Cpim_CpimLastErrorType = "invalidJoinPrune"
)

// CISCOPIMMIB_CiscoPimMIBNotificationObjects
type CISCOPIMMIB_CiscoPimMIBNotificationObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes the operation that resulted in generation of cpimRPMappingChange
    // notification.  o newMapping, as the name suggests indicates that a new  
    // mapping has been added into the pimRPSetTable,  o deletedMapping indicates
    // that a mapping has been    deleted from the pimRPSetTable, and, o
    // modifiedXXXMapping indicates that an RP mapping (which   already existed in
    // the table) has been modified.   The two modifications types i.e.
    // modifiedOldMapping   and modifiedNewMapping, are defined to differentiate  
    // the notification generated before modification from   that generated after
    // modification. The type is CpimRPMappingChangeType.
    CpimRPMappingChangeType interface{}
}

func (ciscoPimMIBNotificationObjects *CISCOPIMMIB_CiscoPimMIBNotificationObjects) GetEntityData() *types.CommonEntityData {
    ciscoPimMIBNotificationObjects.EntityData.YFilter = ciscoPimMIBNotificationObjects.YFilter
    ciscoPimMIBNotificationObjects.EntityData.YangName = "ciscoPimMIBNotificationObjects"
    ciscoPimMIBNotificationObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoPimMIBNotificationObjects.EntityData.ParentYangName = "CISCO-PIM-MIB"
    ciscoPimMIBNotificationObjects.EntityData.SegmentPath = "ciscoPimMIBNotificationObjects"
    ciscoPimMIBNotificationObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoPimMIBNotificationObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoPimMIBNotificationObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoPimMIBNotificationObjects.EntityData.Children = types.NewOrderedMap()
    ciscoPimMIBNotificationObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoPimMIBNotificationObjects.EntityData.Leafs.Append("cpimRPMappingChangeType", types.YLeaf{"CpimRPMappingChangeType", ciscoPimMIBNotificationObjects.CpimRPMappingChangeType})

    ciscoPimMIBNotificationObjects.EntityData.YListKeys = []string {}

    return &(ciscoPimMIBNotificationObjects.EntityData)
}

// CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType represents   that generated after modification.
type CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType string

const (
    CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType_newMapping CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType = "newMapping"

    CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType_deletedMapping CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType = "deletedMapping"

    CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType_modifiedOldMapping CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType = "modifiedOldMapping"

    CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType_modifiedNewMapping CISCOPIMMIB_CiscoPimMIBNotificationObjects_CpimRPMappingChangeType = "modifiedNewMapping"
)

