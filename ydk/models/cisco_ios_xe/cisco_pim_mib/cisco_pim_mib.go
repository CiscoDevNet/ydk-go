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

    
    Ciscopimmibnotificationobjects CISCOPIMMIB_Ciscopimmibnotificationobjects
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

    cISCOPIMMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOPIMMIB.EntityData.Children["cpim"] = types.YChild{"Cpim", &cISCOPIMMIB.Cpim}
    cISCOPIMMIB.EntityData.Children["ciscoPimMIBNotificationObjects"] = types.YChild{"Ciscopimmibnotificationobjects", &cISCOPIMMIB.Ciscopimmibnotificationobjects}
    cISCOPIMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Cpiminvalidregistermsgsrcvd interface{}

    // A count of the number of invalid PIM Join/Prune messages received by this
    // device. A PIM Join/Prune message is termed invalid if  o the RP specified
    // in the packet is not the RP for   the group in question. The type is
    // interface{} with range: 0..4294967295.
    Cpiminvalidjoinprunemsgsrcvd interface{}

    // The type of the last invalid message that was received by this device. The
    // type is Cpimlasterrortype.
    Cpimlasterrortype interface{}

    // Represents the type of address stored in cpimLastErrorOrigin. The value of
    // this object is irrelevant if the value of cpimLastErrorType is none(1). The
    // type is InetAddressType.
    Cpimlasterrororigintype interface{}

    // This object represents the Network Layer Address of the source that
    // originated the last invalid packet. The type of address stored depends on
    // the value in cpimLastErrorOriginType.   The value of this object represents
    // the Network Layer Address of the Designated Router (DR) whenever the value
    // of cpimLastErrorGroup is a zero-length address,  for eg. when encapsulated
    // IP header is malformed.  The value of this object is irrelevant if the
    // value of cpimLastErrorType is none(1). The type is string with length:
    // 0..255.
    Cpimlasterrororigin interface{}

    // Represents the type of address stored in cpimLastErrorGroup. The value of
    // this object is unknown(0) if there is a problem in the packet received from
    // the DR.  The value of this object is irrelevant if the value of
    // cpimLastErrorType is none(1). The type is InetAddressType.
    Cpimlasterrorgrouptype interface{}

    // The IP multicast group address to which the last invalid packet was
    // addressed. The type of address stored depends on the value in
    // cpimLastErrorGroupType.   The value of this object is a zero-length
    // InetAddress if there is a problem in the packet received from the DR, for
    // eg. a malformed encapsulated IP header.   The value of this object is
    // irrelevant if the value of cpimLastErrorType is none(1). The type is string
    // with length: 0..255.
    Cpimlasterrorgroup interface{}

    // Represents the type of address stored in cpimLastErrorRP. The value of this
    // object is irrelevant if the value of cpimLastErrorType is none(1). The type
    // is InetAddressType.
    Cpimlasterrorrptype interface{}

    // The address of the RP, as per the last invalid packet. The type of address
    // stored depends on the value in cpimLastErrorRPType.   The value of this
    // object is irrelevant if the value of cpimLastErrorType is none(1). The type
    // is string with length: 0..255.
    Cpimlasterrorrp interface{}
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

    cpim.EntityData.Children = make(map[string]types.YChild)
    cpim.EntityData.Leafs = make(map[string]types.YLeaf)
    cpim.EntityData.Leafs["cpimInvalidRegisterMsgsRcvd"] = types.YLeaf{"Cpiminvalidregistermsgsrcvd", cpim.Cpiminvalidregistermsgsrcvd}
    cpim.EntityData.Leafs["cpimInvalidJoinPruneMsgsRcvd"] = types.YLeaf{"Cpiminvalidjoinprunemsgsrcvd", cpim.Cpiminvalidjoinprunemsgsrcvd}
    cpim.EntityData.Leafs["cpimLastErrorType"] = types.YLeaf{"Cpimlasterrortype", cpim.Cpimlasterrortype}
    cpim.EntityData.Leafs["cpimLastErrorOriginType"] = types.YLeaf{"Cpimlasterrororigintype", cpim.Cpimlasterrororigintype}
    cpim.EntityData.Leafs["cpimLastErrorOrigin"] = types.YLeaf{"Cpimlasterrororigin", cpim.Cpimlasterrororigin}
    cpim.EntityData.Leafs["cpimLastErrorGroupType"] = types.YLeaf{"Cpimlasterrorgrouptype", cpim.Cpimlasterrorgrouptype}
    cpim.EntityData.Leafs["cpimLastErrorGroup"] = types.YLeaf{"Cpimlasterrorgroup", cpim.Cpimlasterrorgroup}
    cpim.EntityData.Leafs["cpimLastErrorRPType"] = types.YLeaf{"Cpimlasterrorrptype", cpim.Cpimlasterrorrptype}
    cpim.EntityData.Leafs["cpimLastErrorRP"] = types.YLeaf{"Cpimlasterrorrp", cpim.Cpimlasterrorrp}
    return &(cpim.EntityData)
}

// CISCOPIMMIB_Cpim_Cpimlasterrortype represents this device.
type CISCOPIMMIB_Cpim_Cpimlasterrortype string

const (
    CISCOPIMMIB_Cpim_Cpimlasterrortype_none CISCOPIMMIB_Cpim_Cpimlasterrortype = "none"

    CISCOPIMMIB_Cpim_Cpimlasterrortype_invalidRegister CISCOPIMMIB_Cpim_Cpimlasterrortype = "invalidRegister"

    CISCOPIMMIB_Cpim_Cpimlasterrortype_invalidJoinPrune CISCOPIMMIB_Cpim_Cpimlasterrortype = "invalidJoinPrune"
)

// CISCOPIMMIB_Ciscopimmibnotificationobjects
type CISCOPIMMIB_Ciscopimmibnotificationobjects struct {
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
    // modification. The type is Cpimrpmappingchangetype.
    Cpimrpmappingchangetype interface{}
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetEntityData() *types.CommonEntityData {
    ciscopimmibnotificationobjects.EntityData.YFilter = ciscopimmibnotificationobjects.YFilter
    ciscopimmibnotificationobjects.EntityData.YangName = "ciscoPimMIBNotificationObjects"
    ciscopimmibnotificationobjects.EntityData.BundleName = "cisco_ios_xe"
    ciscopimmibnotificationobjects.EntityData.ParentYangName = "CISCO-PIM-MIB"
    ciscopimmibnotificationobjects.EntityData.SegmentPath = "ciscoPimMIBNotificationObjects"
    ciscopimmibnotificationobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscopimmibnotificationobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscopimmibnotificationobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscopimmibnotificationobjects.EntityData.Children = make(map[string]types.YChild)
    ciscopimmibnotificationobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscopimmibnotificationobjects.EntityData.Leafs["cpimRPMappingChangeType"] = types.YLeaf{"Cpimrpmappingchangetype", ciscopimmibnotificationobjects.Cpimrpmappingchangetype}
    return &(ciscopimmibnotificationobjects.EntityData)
}

// CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype represents   that generated after modification.
type CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype string

const (
    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_newMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "newMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_deletedMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "deletedMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_modifiedOldMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "modifiedOldMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_modifiedNewMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "modifiedNewMapping"
)

