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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cpim CISCOPIMMIB_Cpim

    
    Ciscopimmibnotificationobjects CISCOPIMMIB_Ciscopimmibnotificationobjects
}

func (cISCOPIMMIB *CISCOPIMMIB) GetFilter() yfilter.YFilter { return cISCOPIMMIB.YFilter }

func (cISCOPIMMIB *CISCOPIMMIB) SetFilter(yf yfilter.YFilter) { cISCOPIMMIB.YFilter = yf }

func (cISCOPIMMIB *CISCOPIMMIB) GetGoName(yname string) string {
    if yname == "cpim" { return "Cpim" }
    if yname == "ciscoPimMIBNotificationObjects" { return "Ciscopimmibnotificationobjects" }
    return ""
}

func (cISCOPIMMIB *CISCOPIMMIB) GetSegmentPath() string {
    return "CISCO-PIM-MIB:CISCO-PIM-MIB"
}

func (cISCOPIMMIB *CISCOPIMMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpim" {
        return &cISCOPIMMIB.Cpim
    }
    if childYangName == "ciscoPimMIBNotificationObjects" {
        return &cISCOPIMMIB.Ciscopimmibnotificationobjects
    }
    return nil
}

func (cISCOPIMMIB *CISCOPIMMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpim"] = &cISCOPIMMIB.Cpim
    children["ciscoPimMIBNotificationObjects"] = &cISCOPIMMIB.Ciscopimmibnotificationobjects
    return children
}

func (cISCOPIMMIB *CISCOPIMMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOPIMMIB *CISCOPIMMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOPIMMIB *CISCOPIMMIB) GetYangName() string { return "CISCO-PIM-MIB" }

func (cISCOPIMMIB *CISCOPIMMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOPIMMIB *CISCOPIMMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOPIMMIB *CISCOPIMMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOPIMMIB *CISCOPIMMIB) SetParent(parent types.Entity) { cISCOPIMMIB.parent = parent }

func (cISCOPIMMIB *CISCOPIMMIB) GetParent() types.Entity { return cISCOPIMMIB.parent }

func (cISCOPIMMIB *CISCOPIMMIB) GetParentYangName() string { return "CISCO-PIM-MIB" }

// CISCOPIMMIB_Cpim
type CISCOPIMMIB_Cpim struct {
    parent types.Entity
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

func (cpim *CISCOPIMMIB_Cpim) GetFilter() yfilter.YFilter { return cpim.YFilter }

func (cpim *CISCOPIMMIB_Cpim) SetFilter(yf yfilter.YFilter) { cpim.YFilter = yf }

func (cpim *CISCOPIMMIB_Cpim) GetGoName(yname string) string {
    if yname == "cpimInvalidRegisterMsgsRcvd" { return "Cpiminvalidregistermsgsrcvd" }
    if yname == "cpimInvalidJoinPruneMsgsRcvd" { return "Cpiminvalidjoinprunemsgsrcvd" }
    if yname == "cpimLastErrorType" { return "Cpimlasterrortype" }
    if yname == "cpimLastErrorOriginType" { return "Cpimlasterrororigintype" }
    if yname == "cpimLastErrorOrigin" { return "Cpimlasterrororigin" }
    if yname == "cpimLastErrorGroupType" { return "Cpimlasterrorgrouptype" }
    if yname == "cpimLastErrorGroup" { return "Cpimlasterrorgroup" }
    if yname == "cpimLastErrorRPType" { return "Cpimlasterrorrptype" }
    if yname == "cpimLastErrorRP" { return "Cpimlasterrorrp" }
    return ""
}

func (cpim *CISCOPIMMIB_Cpim) GetSegmentPath() string {
    return "cpim"
}

func (cpim *CISCOPIMMIB_Cpim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpim *CISCOPIMMIB_Cpim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpim *CISCOPIMMIB_Cpim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpimInvalidRegisterMsgsRcvd"] = cpim.Cpiminvalidregistermsgsrcvd
    leafs["cpimInvalidJoinPruneMsgsRcvd"] = cpim.Cpiminvalidjoinprunemsgsrcvd
    leafs["cpimLastErrorType"] = cpim.Cpimlasterrortype
    leafs["cpimLastErrorOriginType"] = cpim.Cpimlasterrororigintype
    leafs["cpimLastErrorOrigin"] = cpim.Cpimlasterrororigin
    leafs["cpimLastErrorGroupType"] = cpim.Cpimlasterrorgrouptype
    leafs["cpimLastErrorGroup"] = cpim.Cpimlasterrorgroup
    leafs["cpimLastErrorRPType"] = cpim.Cpimlasterrorrptype
    leafs["cpimLastErrorRP"] = cpim.Cpimlasterrorrp
    return leafs
}

func (cpim *CISCOPIMMIB_Cpim) GetBundleName() string { return "cisco_ios_xe" }

func (cpim *CISCOPIMMIB_Cpim) GetYangName() string { return "cpim" }

func (cpim *CISCOPIMMIB_Cpim) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpim *CISCOPIMMIB_Cpim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpim *CISCOPIMMIB_Cpim) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpim *CISCOPIMMIB_Cpim) SetParent(parent types.Entity) { cpim.parent = parent }

func (cpim *CISCOPIMMIB_Cpim) GetParent() types.Entity { return cpim.parent }

func (cpim *CISCOPIMMIB_Cpim) GetParentYangName() string { return "CISCO-PIM-MIB" }

// CISCOPIMMIB_Cpim_Cpimlasterrortype represents this device.
type CISCOPIMMIB_Cpim_Cpimlasterrortype string

const (
    CISCOPIMMIB_Cpim_Cpimlasterrortype_none CISCOPIMMIB_Cpim_Cpimlasterrortype = "none"

    CISCOPIMMIB_Cpim_Cpimlasterrortype_invalidRegister CISCOPIMMIB_Cpim_Cpimlasterrortype = "invalidRegister"

    CISCOPIMMIB_Cpim_Cpimlasterrortype_invalidJoinPrune CISCOPIMMIB_Cpim_Cpimlasterrortype = "invalidJoinPrune"
)

// CISCOPIMMIB_Ciscopimmibnotificationobjects
type CISCOPIMMIB_Ciscopimmibnotificationobjects struct {
    parent types.Entity
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

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetFilter() yfilter.YFilter { return ciscopimmibnotificationobjects.YFilter }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) SetFilter(yf yfilter.YFilter) { ciscopimmibnotificationobjects.YFilter = yf }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetGoName(yname string) string {
    if yname == "cpimRPMappingChangeType" { return "Cpimrpmappingchangetype" }
    return ""
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetSegmentPath() string {
    return "ciscoPimMIBNotificationObjects"
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpimRPMappingChangeType"] = ciscopimmibnotificationobjects.Cpimrpmappingchangetype
    return leafs
}

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetYangName() string { return "ciscoPimMIBNotificationObjects" }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) SetParent(parent types.Entity) { ciscopimmibnotificationobjects.parent = parent }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetParent() types.Entity { return ciscopimmibnotificationobjects.parent }

func (ciscopimmibnotificationobjects *CISCOPIMMIB_Ciscopimmibnotificationobjects) GetParentYangName() string { return "CISCO-PIM-MIB" }

// CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype represents   that generated after modification.
type CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype string

const (
    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_newMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "newMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_deletedMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "deletedMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_modifiedOldMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "modifiedOldMapping"

    CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype_modifiedNewMapping CISCOPIMMIB_Ciscopimmibnotificationobjects_Cpimrpmappingchangetype = "modifiedNewMapping"
)

