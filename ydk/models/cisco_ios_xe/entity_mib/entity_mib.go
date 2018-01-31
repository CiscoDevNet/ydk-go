// The MIB module for representing multiple logical
// entities supported by a single SNMP agent.
// 
// Copyright (C) The Internet Society (2005).  This
// version of this MIB module is part of RFC 4133; see
// the RFC itself for full legal notices.
package entity_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package entity_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ENTITY-MIB ENTITY-MIB}", reflect.TypeOf(ENTITYMIB{}))
    ydk.RegisterEntity("ENTITY-MIB:ENTITY-MIB", reflect.TypeOf(ENTITYMIB{}))
}

// PhysicalClass represents class is some sort of central processing unit.
type PhysicalClass string

const (
    PhysicalClass_other PhysicalClass = "other"

    PhysicalClass_unknown PhysicalClass = "unknown"

    PhysicalClass_chassis PhysicalClass = "chassis"

    PhysicalClass_backplane PhysicalClass = "backplane"

    PhysicalClass_container PhysicalClass = "container"

    PhysicalClass_powerSupply PhysicalClass = "powerSupply"

    PhysicalClass_fan PhysicalClass = "fan"

    PhysicalClass_sensor PhysicalClass = "sensor"

    PhysicalClass_module PhysicalClass = "module"

    PhysicalClass_port PhysicalClass = "port"

    PhysicalClass_stack PhysicalClass = "stack"

    PhysicalClass_cpu PhysicalClass = "cpu"
)

// ENTITYMIB
type ENTITYMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Entitygeneral ENTITYMIB_Entitygeneral

    // This table contains one row per physical entity.  There is always at least
    // one row for an 'overall' physical entity.
    Entphysicaltable ENTITYMIB_Entphysicaltable

    // This table contains one row per logical entity.  For agents that implement
    // more than one naming scope, at least one entry must exist.  Agents which
    // instantiate all MIB objects within a single naming scope are not required
    // to implement this table.
    Entlogicaltable ENTITYMIB_Entlogicaltable

    // This table contains zero or more rows of logical entity to physical
    // equipment associations.  For each logical entity known by this agent, there
    // are zero or more mappings to the physical resources, which are used to
    // realize that logical entity.  An agent should limit the number and nature
    // of entries in this table such that only meaningful and non-redundant
    // information is returned.  For example, in a system that contains a single
    // power supply, mappings between logical entities and the power supply are
    // not useful and should not be included.  Also, only the most appropriate
    // physical component, which is closest to the root of a particular
    // containment tree, should be identified in an entLPMapping entry.  For
    // example, suppose a bridge is realized on a particular module, and all ports
    // on that module are ports on this bridge.  A mapping between the bridge and
    // the module would be useful, but additional mappings between the bridge and
    // each of the ports on that module would be redundant (because the
    // entPhysicalContainedIn hierarchy can provide the same information).  On the
    // other hand, if more than one bridge were utilizing ports on this module,
    // then mappings between each bridge and the ports it used would be
    // appropriate.  Also, in the case of a single backplane repeater, a mapping
    // for the backplane to the single repeater entity is not necessary.
    Entlpmappingtable ENTITYMIB_Entlpmappingtable

    // This table contains zero or more rows, representing mappings of logical
    // entity and physical component to external MIB identifiers.  Each physical
    // port in the system may be associated with a mapping to an external
    // identifier, which itself is associated with a particular logical entity's
    // naming scope.  A 'wildcard' mechanism is provided to indicate that an
    // identifier is associated with more than one logical entity.
    Entaliasmappingtable ENTITYMIB_Entaliasmappingtable

    // A table that exposes the container/'containee' relationships between
    // physical entities.  This table provides all the information found by
    // constructing the virtual containment tree for a given entPhysicalTable, but
    // in a more direct format.  In the event a physical entity is contained by
    // more than one other physical entity (e.g., double-wide modules), this table
    // should include these additional mappings, which cannot be represented in
    // the entPhysicalTable virtual containment tree.
    Entphysicalcontainstable ENTITYMIB_Entphysicalcontainstable
}

func (eNTITYMIB *ENTITYMIB) GetFilter() yfilter.YFilter { return eNTITYMIB.YFilter }

func (eNTITYMIB *ENTITYMIB) SetFilter(yf yfilter.YFilter) { eNTITYMIB.YFilter = yf }

func (eNTITYMIB *ENTITYMIB) GetGoName(yname string) string {
    if yname == "entityGeneral" { return "Entitygeneral" }
    if yname == "entPhysicalTable" { return "Entphysicaltable" }
    if yname == "entLogicalTable" { return "Entlogicaltable" }
    if yname == "entLPMappingTable" { return "Entlpmappingtable" }
    if yname == "entAliasMappingTable" { return "Entaliasmappingtable" }
    if yname == "entPhysicalContainsTable" { return "Entphysicalcontainstable" }
    return ""
}

func (eNTITYMIB *ENTITYMIB) GetSegmentPath() string {
    return "ENTITY-MIB:ENTITY-MIB"
}

func (eNTITYMIB *ENTITYMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entityGeneral" {
        return &eNTITYMIB.Entitygeneral
    }
    if childYangName == "entPhysicalTable" {
        return &eNTITYMIB.Entphysicaltable
    }
    if childYangName == "entLogicalTable" {
        return &eNTITYMIB.Entlogicaltable
    }
    if childYangName == "entLPMappingTable" {
        return &eNTITYMIB.Entlpmappingtable
    }
    if childYangName == "entAliasMappingTable" {
        return &eNTITYMIB.Entaliasmappingtable
    }
    if childYangName == "entPhysicalContainsTable" {
        return &eNTITYMIB.Entphysicalcontainstable
    }
    return nil
}

func (eNTITYMIB *ENTITYMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entityGeneral"] = &eNTITYMIB.Entitygeneral
    children["entPhysicalTable"] = &eNTITYMIB.Entphysicaltable
    children["entLogicalTable"] = &eNTITYMIB.Entlogicaltable
    children["entLPMappingTable"] = &eNTITYMIB.Entlpmappingtable
    children["entAliasMappingTable"] = &eNTITYMIB.Entaliasmappingtable
    children["entPhysicalContainsTable"] = &eNTITYMIB.Entphysicalcontainstable
    return children
}

func (eNTITYMIB *ENTITYMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eNTITYMIB *ENTITYMIB) GetBundleName() string { return "cisco_ios_xe" }

func (eNTITYMIB *ENTITYMIB) GetYangName() string { return "ENTITY-MIB" }

func (eNTITYMIB *ENTITYMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eNTITYMIB *ENTITYMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eNTITYMIB *ENTITYMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eNTITYMIB *ENTITYMIB) SetParent(parent types.Entity) { eNTITYMIB.parent = parent }

func (eNTITYMIB *ENTITYMIB) GetParent() types.Entity { return eNTITYMIB.parent }

func (eNTITYMIB *ENTITYMIB) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entitygeneral
type ENTITYMIB_Entitygeneral struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time a conceptual row is created, modified,
    // or deleted in any of these tables:         - entPhysicalTable         -
    // entLogicalTable         - entLPMappingTable         - entAliasMappingTable 
    // - entPhysicalContainsTable. The type is interface{} with range:
    // 0..4294967295.
    Entlastchangetime interface{}
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetFilter() yfilter.YFilter { return entitygeneral.YFilter }

func (entitygeneral *ENTITYMIB_Entitygeneral) SetFilter(yf yfilter.YFilter) { entitygeneral.YFilter = yf }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetGoName(yname string) string {
    if yname == "entLastChangeTime" { return "Entlastchangetime" }
    return ""
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetSegmentPath() string {
    return "entityGeneral"
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entLastChangeTime"] = entitygeneral.Entlastchangetime
    return leafs
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetBundleName() string { return "cisco_ios_xe" }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetYangName() string { return "entityGeneral" }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entitygeneral *ENTITYMIB_Entitygeneral) SetParent(parent types.Entity) { entitygeneral.parent = parent }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetParent() types.Entity { return entitygeneral.parent }

func (entitygeneral *ENTITYMIB_Entitygeneral) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entphysicaltable
// This table contains one row per physical entity.  There is
// always at least one row for an 'overall' physical entity.
type ENTITYMIB_Entphysicaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular physical entity.  Each entry provides
    // objects (entPhysicalDescr, entPhysicalVendorType, and entPhysicalClass) to
    // help an NMS identify and characterize the entry, and objects
    // (entPhysicalContainedIn and entPhysicalParentRelPos) to help an NMS relate
    // the particular entry to other entries in this table. The type is slice of
    // ENTITYMIB_Entphysicaltable_Entphysicalentry.
    Entphysicalentry []ENTITYMIB_Entphysicaltable_Entphysicalentry
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetFilter() yfilter.YFilter { return entphysicaltable.YFilter }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) SetFilter(yf yfilter.YFilter) { entphysicaltable.YFilter = yf }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetGoName(yname string) string {
    if yname == "entPhysicalEntry" { return "Entphysicalentry" }
    return ""
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetSegmentPath() string {
    return "entPhysicalTable"
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entPhysicalEntry" {
        for _, c := range entphysicaltable.Entphysicalentry {
            if entphysicaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYMIB_Entphysicaltable_Entphysicalentry{}
        entphysicaltable.Entphysicalentry = append(entphysicaltable.Entphysicalentry, child)
        return &entphysicaltable.Entphysicalentry[len(entphysicaltable.Entphysicalentry)-1]
    }
    return nil
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entphysicaltable.Entphysicalentry {
        children[entphysicaltable.Entphysicalentry[i].GetSegmentPath()] = &entphysicaltable.Entphysicalentry[i]
    }
    return children
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetBundleName() string { return "cisco_ios_xe" }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetYangName() string { return "entPhysicalTable" }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) SetParent(parent types.Entity) { entphysicaltable.parent = parent }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetParent() types.Entity { return entphysicaltable.parent }

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entphysicaltable_Entphysicalentry
// Information about a particular physical entity.
// 
// Each entry provides objects (entPhysicalDescr,
// entPhysicalVendorType, and entPhysicalClass) to help an NMS
// identify and characterize the entry, and objects
// (entPhysicalContainedIn and entPhysicalParentRelPos) to help
// an NMS relate the particular entry to other entries in this
// table.
type ENTITYMIB_Entphysicaltable_Entphysicalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index for this entry. The type is interface{}
    // with range: 1..2147483647.
    Entphysicalindex interface{}

    // A textual description of physical entity.  This object should contain a
    // string that identifies the manufacturer's name for the physical entity, and
    // should be set to a distinct value for each version or model of the physical
    // entity. The type is string.
    Entphysicaldescr interface{}

    // An indication of the vendor-specific hardware type of the physical entity. 
    // Note that this is different from the definition of MIB-II's sysObjectID. 
    // An agent should set this object to an enterprise-specific registration
    // identifier value indicating the specific equipment type in detail.  The
    // associated instance of entPhysicalClass is used to indicate the general
    // type of hardware device.  If no vendor-specific registration identifier
    // exists for this physical entity, or the value is unknown by this agent,
    // then the value { 0 0 } is returned. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Entphysicalvendortype interface{}

    // The value of entPhysicalIndex for the physical entity which 'contains' this
    // physical entity.  A value of zero indicates this physical entity is not
    // contained in any other physical entity.  Note that the set of 'containment'
    // relationships define a strict hierarchy; that is, recursion is not allowed.
    // In the event that a physical entity is contained by more than one physical
    // entity (e.g., double-wide modules), this object should identify the
    // containing entity with the lowest value of entPhysicalIndex. The type is
    // interface{} with range: 0..2147483647.
    Entphysicalcontainedin interface{}

    // An indication of the general hardware type of the physical entity.  An
    // agent should set this object to the standard enumeration value that most
    // accurately indicates the general class of the physical entity, or the
    // primary class if there is more than one entity.  If no appropriate standard
    // registration identifier exists for this physical entity, then the value
    // 'other(1)' is returned.  If the value is unknown by this agent, then the
    // value 'unknown(2)' is returned. The type is PhysicalClass.
    Entphysicalclass interface{}

    // An indication of the relative position of this 'child' component among all
    // its 'sibling' components.  Sibling components are defined as
    // entPhysicalEntries that share the same instance values of each of the
    // entPhysicalContainedIn and entPhysicalClass objects.  An NMS can use this
    // object to identify the relative ordering for all sibling components of a
    // particular parent (identified by the entPhysicalContainedIn instance in
    // each sibling entry).  If possible, this value should match any external
    // labeling of the physical component.  For example, for a container (e.g.,
    // card slot) labeled as 'slot #3', entPhysicalParentRelPos should have the
    // value '3'.  Note that the entPhysicalEntry for the module plugged in slot 3
    // should have an entPhysicalParentRelPos value of '1'.  If the physical
    // position of this component does not match any external numbering or clearly
    // visible ordering, then user documentation or other external reference
    // material should be used to determine the parent-relative position. If this
    // is not possible, then the agent should assign a consistent (but possibly
    // arbitrary) ordering to a given set of 'sibling' components, perhaps based
    // on internal representation of the components.   If the agent cannot
    // determine the parent-relative position for some reason, or if the
    // associated value of entPhysicalContainedIn is '0', then the value '-1' is
    // returned.  Otherwise, a non-negative integer is returned, indicating the
    // parent-relative position of this physical entity.  Parent-relative ordering
    // normally starts from '1' and continues to 'N', where 'N' represents the
    // highest positioned child entity.  However, if the physical entities (e.g.,
    // slots) are labeled from a starting position of zero, then the first sibling
    // should be associated with an entPhysicalParentRelPos value of '0'.  Note
    // that this ordering may be sparse or dense, depending on agent
    // implementation.  The actual values returned are not globally meaningful, as
    // each 'parent' component may use different numbering algorithms.  The
    // ordering is only meaningful among siblings of the same parent component. 
    // The agent should retain parent-relative position values across reboots,
    // either through algorithmic assignment or use of non-volatile storage. The
    // type is interface{} with range: -1..2147483647.
    Entphysicalparentrelpos interface{}

    // The textual name of the physical entity.  The value of this object should
    // be the name of the component as assigned by the local device and should be
    // suitable for use in commands entered at the device's `console'.  This might
    // be a text name (e.g., `console') or a simple component number (e.g., port
    // or module number, such as `1'), depending on the physical component naming
    // syntax of the device.  If there is no local name, or if this object is
    // otherwise not applicable, then this object contains a zero-length string. 
    // Note that the value of entPhysicalName for two physical entities will be
    // the same in the event that the console interface does not distinguish
    // between them, e.g., slot-1 and the card in slot-1. The type is string.
    Entphysicalname interface{}

    // The vendor-specific hardware revision string for the physical entity.  The
    // preferred value is the hardware revision identifier actually printed on the
    // component itself (if present).  Note that if revision information is stored
    // internally in a non-printable (e.g., binary) format, then the agent must
    // convert such information to a printable format, in an
    // implementation-specific manner.  If no specific hardware revision string is
    // associated with the physical component, or if this information is unknown
    // to the agent, then this object will contain a zero-length string. The type
    // is string.
    Entphysicalhardwarerev interface{}

    // The vendor-specific firmware revision string for the physical entity.  Note
    // that if revision information is stored internally in a non-printable (e.g.,
    // binary) format, then the agent must convert such information to a printable
    // format, in an implementation-specific manner.  If no specific firmware
    // programs are associated with the physical component, or if this information
    // is unknown to the agent, then this object will contain a zero-length
    // string. The type is string.
    Entphysicalfirmwarerev interface{}

    // The vendor-specific software revision string for the physical entity.  Note
    // that if revision information is stored internally in a   non-printable
    // (e.g., binary) format, then the agent must convert such information to a
    // printable format, in an implementation-specific manner.  If no specific
    // software programs are associated with the physical component, or if this
    // information is unknown to the agent, then this object will contain a
    // zero-length string. The type is string.
    Entphysicalsoftwarerev interface{}

    // The vendor-specific serial number string for the physical entity.  The
    // preferred value is the serial number string actually printed on the
    // component itself (if present).  On the first instantiation of an physical
    // entity, the value of entPhysicalSerialNum associated with that entity is
    // set to the correct vendor-assigned serial number, if this information is
    // available to the agent.  If a serial number is unknown or non-existent, the
    // entPhysicalSerialNum will be set to a zero-length string instead.  Note
    // that implementations that can correctly identify the serial numbers of all
    // installed physical entities do not need to provide write access to the
    // entPhysicalSerialNum object.  Agents which cannot provide non-volatile
    // storage for the entPhysicalSerialNum strings are not required to implement
    // write access for this object.  Not every physical component will have a
    // serial number, or even need one.  Physical entities for which the
    // associated value of the entPhysicalIsFRU object is equal to 'false(2)'
    // (e.g., the repeater ports within a repeater module), do not need their own
    // unique serial number.  An agent does not have to provide write access for
    // such entities, and may return a zero-length string.  If write access is
    // implemented for an instance of entPhysicalSerialNum, and a value is written
    // into the instance, the agent must retain the supplied value in the
    // entPhysicalSerialNum instance (associated with the same physical entity)
    // for as long as that entity remains instantiated.  This includes
    // instantiations across all re-initializations/reboots of the network
    // management system, including those resulting in a change of the physical  
    // entity's entPhysicalIndex value. The type is string with length: 0..32.
    Entphysicalserialnum interface{}

    // The name of the manufacturer of this physical component. The preferred
    // value is the manufacturer name string actually printed on the component
    // itself (if present).  Note that comparisons between instances of the
    // entPhysicalModelName, entPhysicalFirmwareRev, entPhysicalSoftwareRev, and
    // the entPhysicalSerialNum objects, are only meaningful amongst
    // entPhysicalEntries with the same value of entPhysicalMfgName.  If the
    // manufacturer name string associated with the physical component is unknown
    // to the agent, then this object will contain a zero-length string. The type
    // is string.
    Entphysicalmfgname interface{}

    // The vendor-specific model name identifier string associated with this
    // physical component.  The preferred value is the customer-visible part
    // number, which may be printed on the component itself.  If the model name
    // string associated with the physical component is unknown to the agent, then
    // this object will contain a zero-length string. The type is string.
    Entphysicalmodelname interface{}

    // This object is an 'alias' name for the physical entity, as specified by a
    // network manager, and provides a non-volatile 'handle' for the physical
    // entity.  On the first instantiation of a physical entity, the value   of
    // entPhysicalAlias associated with that entity is set to the zero-length
    // string.  However, the agent may set the value to a locally unique default
    // value, instead of a zero-length string.  If write access is implemented for
    // an instance of entPhysicalAlias, and a value is written into the instance,
    // the agent must retain the supplied value in the entPhysicalAlias instance
    // (associated with the same physical entity) for as long as that entity
    // remains instantiated. This includes instantiations across all
    // re-initializations/reboots of the network management system, including
    // those resulting in a change of the physical entity's entPhysicalIndex
    // value. The type is string with length: 0..32.
    Entphysicalalias interface{}

    // This object is a user-assigned asset tracking identifier (as specified by a
    // network manager) for the physical entity, and provides non-volatile storage
    // of this information.  On the first instantiation of a physical entity, the
    // value of entPhysicalAssetID associated with that entity is set to the
    // zero-length string.  Not every physical component will have an asset
    // tracking identifier, or even need one.  Physical entities for which the
    // associated value of the entPhysicalIsFRU object is equal to 'false(2)'
    // (e.g., the repeater ports within a repeater module), do not need their own
    // unique asset tracking identifier.  An agent does not have to provide write
    // access for such entities, and may instead return a zero-length string.  If
    // write access is implemented for an instance of entPhysicalAssetID, and a
    // value is written into the instance, the agent must retain the supplied
    // value in the entPhysicalAssetID instance (associated with the same physical
    // entity) for as long as that entity remains instantiated.  This includes
    // instantiations across all re-initializations/reboots of the network
    // management system, including those resulting in a change of the physical
    // entity's entPhysicalIndex value.   If no asset tracking information is
    // associated with the physical component, then this object will contain a
    // zero-length string. The type is string with length: 0..32.
    Entphysicalassetid interface{}

    // This object indicates whether or not this physical entity is considered a
    // 'field replaceable unit' by the vendor.  If this object contains the value
    // 'true(1)' then this entPhysicalEntry identifies a field replaceable unit. 
    // For all entPhysicalEntries that represent components permanently contained
    // within a field replaceable unit, the value 'false(2)' should be returned
    // for this object. The type is bool.
    Entphysicalisfru interface{}

    // This object contains the date of manufacturing of the managed entity.  If
    // the manufacturing date is unknown or not supported, the object is not
    // instantiated.  The special value '0000000000000000'H may also be returned
    // in this case. The type is string.
    Entphysicalmfgdate interface{}

    // This object contains additional identification information about the
    // physical entity.  The object contains URIs and, therefore, the syntax of
    // this object must conform to RFC 3986, section 2.  Multiple URIs may be
    // present and are separated by white space characters.  Leading and trailing
    // white space characters are ignored.  If no additional identification
    // information is known about the physical entity or supported, the object is
    // not instantiated.  A zero length octet string may also be   returned in
    // this case. The type is string.
    Entphysicaluris interface{}

    // This object represents the vendor-specific second serial number string for
    // the physical entity.  The first serial number string of the physical entity
    // is represented in the value of corresponding  instance of the
    // 'entPhysicalSerialNum' object.  On the first instantiation of an physical
    // entity, the value of this object is the correct vendor-assigned second 
    // serial number, if this information is available to the agent.   If the
    // second serial number is unknown or non-existent, then  the value of this
    // object will be a zero-length string instead.  Note that implementations
    // which can correctly identify the second serial numbers of all installed
    // physical entities do  not need to provide write access to this object. 
    // Agents which cannot provide non-volatile storage for the  second serial
    // number strings are not required to implement  write access for this object.
    // Not every physical component will have a serial number, or even need one. 
    // Physical entities for which the associated value of the entPhysicalIsFRU
    // object is equal to 'false(2)' (e.g., the repeater ports within a repeater
    // module), do not need their own unique serial number. An agent does not have
    // to provide write access for such entities, and may return a zero-length
    // string.  If write access is implemented for an instance of
    // 'ceEntPhysicalSecondSerialNum', and a value is written into  the instance,
    // the agent must retain the supplied value in the
    // 'ceEntPhysicalSecondSerialNum' instance associated with the  same physical
    // entity for as long as that entity remains instantiated. This includes
    // instantiations across all re- initializations/reboots of the network
    // management system, including those which result in a change of the physical
    // entity's entPhysicalIndex value. The type is string with length: 0..32.
    Ceentphysicalsecondserialnum interface{}
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetFilter() yfilter.YFilter { return entphysicalentry.YFilter }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) SetFilter(yf yfilter.YFilter) { entphysicalentry.YFilter = yf }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entPhysicalDescr" { return "Entphysicaldescr" }
    if yname == "entPhysicalVendorType" { return "Entphysicalvendortype" }
    if yname == "entPhysicalContainedIn" { return "Entphysicalcontainedin" }
    if yname == "entPhysicalClass" { return "Entphysicalclass" }
    if yname == "entPhysicalParentRelPos" { return "Entphysicalparentrelpos" }
    if yname == "entPhysicalName" { return "Entphysicalname" }
    if yname == "entPhysicalHardwareRev" { return "Entphysicalhardwarerev" }
    if yname == "entPhysicalFirmwareRev" { return "Entphysicalfirmwarerev" }
    if yname == "entPhysicalSoftwareRev" { return "Entphysicalsoftwarerev" }
    if yname == "entPhysicalSerialNum" { return "Entphysicalserialnum" }
    if yname == "entPhysicalMfgName" { return "Entphysicalmfgname" }
    if yname == "entPhysicalModelName" { return "Entphysicalmodelname" }
    if yname == "entPhysicalAlias" { return "Entphysicalalias" }
    if yname == "entPhysicalAssetID" { return "Entphysicalassetid" }
    if yname == "entPhysicalIsFRU" { return "Entphysicalisfru" }
    if yname == "entPhysicalMfgDate" { return "Entphysicalmfgdate" }
    if yname == "entPhysicalUris" { return "Entphysicaluris" }
    if yname == "ceEntPhysicalSecondSerialNum" { return "Ceentphysicalsecondserialnum" }
    return ""
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetSegmentPath() string {
    return "entPhysicalEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalentry.Entphysicalindex) + "']"
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entphysicalentry.Entphysicalindex
    leafs["entPhysicalDescr"] = entphysicalentry.Entphysicaldescr
    leafs["entPhysicalVendorType"] = entphysicalentry.Entphysicalvendortype
    leafs["entPhysicalContainedIn"] = entphysicalentry.Entphysicalcontainedin
    leafs["entPhysicalClass"] = entphysicalentry.Entphysicalclass
    leafs["entPhysicalParentRelPos"] = entphysicalentry.Entphysicalparentrelpos
    leafs["entPhysicalName"] = entphysicalentry.Entphysicalname
    leafs["entPhysicalHardwareRev"] = entphysicalentry.Entphysicalhardwarerev
    leafs["entPhysicalFirmwareRev"] = entphysicalentry.Entphysicalfirmwarerev
    leafs["entPhysicalSoftwareRev"] = entphysicalentry.Entphysicalsoftwarerev
    leafs["entPhysicalSerialNum"] = entphysicalentry.Entphysicalserialnum
    leafs["entPhysicalMfgName"] = entphysicalentry.Entphysicalmfgname
    leafs["entPhysicalModelName"] = entphysicalentry.Entphysicalmodelname
    leafs["entPhysicalAlias"] = entphysicalentry.Entphysicalalias
    leafs["entPhysicalAssetID"] = entphysicalentry.Entphysicalassetid
    leafs["entPhysicalIsFRU"] = entphysicalentry.Entphysicalisfru
    leafs["entPhysicalMfgDate"] = entphysicalentry.Entphysicalmfgdate
    leafs["entPhysicalUris"] = entphysicalentry.Entphysicaluris
    leafs["ceEntPhysicalSecondSerialNum"] = entphysicalentry.Ceentphysicalsecondserialnum
    return leafs
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetBundleName() string { return "cisco_ios_xe" }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetYangName() string { return "entPhysicalEntry" }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) SetParent(parent types.Entity) { entphysicalentry.parent = parent }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetParent() types.Entity { return entphysicalentry.parent }

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetParentYangName() string { return "entPhysicalTable" }

// ENTITYMIB_Entlogicaltable
// This table contains one row per logical entity.  For agents
// that implement more than one naming scope, at least one
// entry must exist.  Agents which instantiate all MIB objects
// within a single naming scope are not required to implement
// this table.
type ENTITYMIB_Entlogicaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular logical entity.  Entities may be managed by
    // this agent or other SNMP agents (possibly) in the same chassis. The type is
    // slice of ENTITYMIB_Entlogicaltable_Entlogicalentry.
    Entlogicalentry []ENTITYMIB_Entlogicaltable_Entlogicalentry
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetFilter() yfilter.YFilter { return entlogicaltable.YFilter }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) SetFilter(yf yfilter.YFilter) { entlogicaltable.YFilter = yf }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetGoName(yname string) string {
    if yname == "entLogicalEntry" { return "Entlogicalentry" }
    return ""
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetSegmentPath() string {
    return "entLogicalTable"
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entLogicalEntry" {
        for _, c := range entlogicaltable.Entlogicalentry {
            if entlogicaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYMIB_Entlogicaltable_Entlogicalentry{}
        entlogicaltable.Entlogicalentry = append(entlogicaltable.Entlogicalentry, child)
        return &entlogicaltable.Entlogicalentry[len(entlogicaltable.Entlogicalentry)-1]
    }
    return nil
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entlogicaltable.Entlogicalentry {
        children[entlogicaltable.Entlogicalentry[i].GetSegmentPath()] = &entlogicaltable.Entlogicalentry[i]
    }
    return children
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetBundleName() string { return "cisco_ios_xe" }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetYangName() string { return "entLogicalTable" }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) SetParent(parent types.Entity) { entlogicaltable.parent = parent }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetParent() types.Entity { return entlogicaltable.parent }

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entlogicaltable_Entlogicalentry
// Information about a particular logical entity.  Entities
// may be managed by this agent or other SNMP agents (possibly)
// in the same chassis.
type ENTITYMIB_Entlogicaltable_Entlogicalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of this object uniquely identifies the
    // logical entity.  The value should be a small positive integer; index values
    // for different logical entities are not necessarily contiguous. The type is
    // interface{} with range: 1..2147483647.
    Entlogicalindex interface{}

    // A textual description of the logical entity.  This object should contain a
    // string that identifies the manufacturer's name for the logical entity, and
    // should be set to a distinct value for each version of the logical entity.
    // The type is string.
    Entlogicaldescr interface{}

    // An indication of the type of logical entity.  This will typically be the
    // OBJECT IDENTIFIER name of the node in the SMI's naming hierarchy which
    // represents the major MIB module, or the majority of the MIB modules,
    // supported by the logical entity.  For example:    a logical entity of a
    // regular host/router -> mib-2    a logical entity of a 802.1d bridge ->
    // dot1dBridge    a logical entity of a 802.3 repeater -> snmpDot3RptrMgmt If
    // an appropriate node in the SMI's naming hierarchy cannot be identified, the
    // value 'mib-2' should be used. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Entlogicaltype interface{}

    // An SNMPv1 or SNMPv2C community-string, which can be used to access detailed
    // management information for this logical entity.  The agent should allow
    // read access with this community string (to an appropriate subset of all
    // managed objects) and may also return a community string based on the
    // privileges of the request used to read this object.  Note that an agent may
    // return a community string with read-only privileges, even if this object is
    // accessed with a read-write community string.  However, the agent must take 
    // care not to return a community string that allows more privileges than the
    // community string used to access this object.  A compliant SNMP agent may
    // wish to conserve naming scopes by representing multiple logical entities in
    // a single 'default' naming scope.  This is possible when the logical
    // entities, represented by the same value of entLogicalCommunity, have no
    // object instances in common.  For example, 'bridge1' and 'repeater1' may be
    // part of the main naming scope, but at least one additional community string
    // is needed to represent 'bridge2' and 'repeater2'.  Logical entities
    // 'bridge1' and 'repeater1' would be represented by sysOREntries associated
    // with the 'default' naming scope.  For agents not accessible via SNMPv1 or
    // SNMPv2C, the value of this object is the empty string.  This object may
    // also contain an empty string if a community string has not yet been
    // assigned by the agent, or if no community string with suitable access
    // rights can be returned for a particular SNMP request.  Note that this
    // object is deprecated.  Agents which implement SNMPv3 access should use the
    // entLogicalContextEngineID and entLogicalContextName objects to identify the
    // context associated with each logical entity.  SNMPv3 agents may return a
    // zero-length string for this object, or may continue to return a community
    // string (e.g., tri-lingual agent support). The type is string with length:
    // 0..255.
    Entlogicalcommunity interface{}

    // The transport service address by which the logical entity receives network
    // management traffic, formatted according to the corresponding value of
    // entLogicalTDomain.  For snmpUDPDomain, a TAddress is 6 octets long: the
    // initial 4 octets contain the IP-address in network-byte order and the last
    // 2 contain the UDP port in network-byte order. Consult 'Transport Mappings
    // for the Simple Network Management Protocol' (STD 62, RFC 3417 [RFC3417])
    // for further information on snmpUDPDomain. The type is string with length:
    // 1..255.
    Entlogicaltaddress interface{}

    // Indicates the kind of transport service by which the logical entity
    // receives network management traffic. Possible values for this object are
    // presently found in the Transport Mappings for Simple Network Management
    // Protocol' (STD 62, RFC 3417 [RFC3417]). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Entlogicaltdomain interface{}

    // The authoritative contextEngineID that can be used to send an SNMP message
    // concerning information held by this logical entity, to the address
    // specified by the associated 'entLogicalTAddress/entLogicalTDomain' pair. 
    // This object, together with the associated entLogicalContextName object,
    // defines the context associated with a particular logical entity, and allows
    // access to SNMP engines identified by a contextEngineId and contextName
    // pair.  If no value has been configured by the agent, a zero-length string
    // is returned, or the agent may choose not to instantiate this object at all.
    // The type is string with length: 0..32.
    Entlogicalcontextengineid interface{}

    // The contextName that can be used to send an SNMP message concerning
    // information held by this logical entity, to the address specified by the
    // associated 'entLogicalTAddress/entLogicalTDomain' pair.  This object,
    // together with the associated entLogicalContextEngineID object, defines the
    // context associated with a particular logical entity, and allows   access to
    // SNMP engines identified by a contextEngineId and contextName pair.  If no
    // value has been configured by the agent, a zero-length string is returned,
    // or the agent may choose not to instantiate this object at all. The type is
    // string.
    Entlogicalcontextname interface{}
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetFilter() yfilter.YFilter { return entlogicalentry.YFilter }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) SetFilter(yf yfilter.YFilter) { entlogicalentry.YFilter = yf }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetGoName(yname string) string {
    if yname == "entLogicalIndex" { return "Entlogicalindex" }
    if yname == "entLogicalDescr" { return "Entlogicaldescr" }
    if yname == "entLogicalType" { return "Entlogicaltype" }
    if yname == "entLogicalCommunity" { return "Entlogicalcommunity" }
    if yname == "entLogicalTAddress" { return "Entlogicaltaddress" }
    if yname == "entLogicalTDomain" { return "Entlogicaltdomain" }
    if yname == "entLogicalContextEngineID" { return "Entlogicalcontextengineid" }
    if yname == "entLogicalContextName" { return "Entlogicalcontextname" }
    return ""
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetSegmentPath() string {
    return "entLogicalEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlogicalentry.Entlogicalindex) + "']"
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entLogicalIndex"] = entlogicalentry.Entlogicalindex
    leafs["entLogicalDescr"] = entlogicalentry.Entlogicaldescr
    leafs["entLogicalType"] = entlogicalentry.Entlogicaltype
    leafs["entLogicalCommunity"] = entlogicalentry.Entlogicalcommunity
    leafs["entLogicalTAddress"] = entlogicalentry.Entlogicaltaddress
    leafs["entLogicalTDomain"] = entlogicalentry.Entlogicaltdomain
    leafs["entLogicalContextEngineID"] = entlogicalentry.Entlogicalcontextengineid
    leafs["entLogicalContextName"] = entlogicalentry.Entlogicalcontextname
    return leafs
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetBundleName() string { return "cisco_ios_xe" }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetYangName() string { return "entLogicalEntry" }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) SetParent(parent types.Entity) { entlogicalentry.parent = parent }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetParent() types.Entity { return entlogicalentry.parent }

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetParentYangName() string { return "entLogicalTable" }

// ENTITYMIB_Entlpmappingtable
// This table contains zero or more rows of logical entity to
// physical equipment associations.  For each logical entity
// known by this agent, there are zero or more mappings to the
// physical resources, which are used to realize that logical
// entity.
// 
// An agent should limit the number and nature of entries in
// this table such that only meaningful and non-redundant
// information is returned.  For example, in a system that
// contains a single power supply, mappings between logical
// entities and the power supply are not useful and should not
// be included.
// 
// Also, only the most appropriate physical component, which is
// closest to the root of a particular containment tree, should
// be identified in an entLPMapping entry.
// 
// For example, suppose a bridge is realized on a particular
// module, and all ports on that module are ports on this
// bridge.  A mapping between the bridge and the module would
// be useful, but additional mappings between the bridge and
// each of the ports on that module would be redundant (because
// the entPhysicalContainedIn hierarchy can provide the same
// information).  On the other hand, if more than one bridge
// were utilizing ports on this module, then mappings between
// each bridge and the ports it used would be appropriate.
// 
// Also, in the case of a single backplane repeater, a mapping
// for the backplane to the single repeater entity is not
// necessary.
type ENTITYMIB_Entlpmappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular logical entity to physical equipment
    // association.  Note that the nature of the association is not specifically
    // identified in this entry. It is expected that sufficient information exists
    // in the MIBs used to manage a particular logical entity to infer how
    // physical component information is utilized. The type is slice of
    // ENTITYMIB_Entlpmappingtable_Entlpmappingentry.
    Entlpmappingentry []ENTITYMIB_Entlpmappingtable_Entlpmappingentry
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetFilter() yfilter.YFilter { return entlpmappingtable.YFilter }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) SetFilter(yf yfilter.YFilter) { entlpmappingtable.YFilter = yf }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetGoName(yname string) string {
    if yname == "entLPMappingEntry" { return "Entlpmappingentry" }
    return ""
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetSegmentPath() string {
    return "entLPMappingTable"
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entLPMappingEntry" {
        for _, c := range entlpmappingtable.Entlpmappingentry {
            if entlpmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYMIB_Entlpmappingtable_Entlpmappingentry{}
        entlpmappingtable.Entlpmappingentry = append(entlpmappingtable.Entlpmappingentry, child)
        return &entlpmappingtable.Entlpmappingentry[len(entlpmappingtable.Entlpmappingentry)-1]
    }
    return nil
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entlpmappingtable.Entlpmappingentry {
        children[entlpmappingtable.Entlpmappingentry[i].GetSegmentPath()] = &entlpmappingtable.Entlpmappingentry[i]
    }
    return children
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetYangName() string { return "entLPMappingTable" }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) SetParent(parent types.Entity) { entlpmappingtable.parent = parent }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetParent() types.Entity { return entlpmappingtable.parent }

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entlpmappingtable_Entlpmappingentry
// Information about a particular logical entity to physical
// equipment association.  Note that the nature of the
// association is not specifically identified in this entry.
// It is expected that sufficient information exists in the
// MIBs used to manage a particular logical entity to infer how
// physical component information is utilized.
type ENTITYMIB_Entlpmappingtable_Entlpmappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entlogicaltable_Entlogicalentry_Entlogicalindex
    Entlogicalindex interface{}

    // This attribute is a key. The value of this object identifies the index
    // value of a particular entPhysicalEntry associated with the indicated
    // entLogicalEntity. The type is interface{} with range: 1..2147483647.
    Entlpphysicalindex interface{}
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetFilter() yfilter.YFilter { return entlpmappingentry.YFilter }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) SetFilter(yf yfilter.YFilter) { entlpmappingentry.YFilter = yf }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetGoName(yname string) string {
    if yname == "entLogicalIndex" { return "Entlogicalindex" }
    if yname == "entLPPhysicalIndex" { return "Entlpphysicalindex" }
    return ""
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetSegmentPath() string {
    return "entLPMappingEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlogicalindex) + "']" + "[entLPPhysicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlpphysicalindex) + "']"
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entLogicalIndex"] = entlpmappingentry.Entlogicalindex
    leafs["entLPPhysicalIndex"] = entlpmappingentry.Entlpphysicalindex
    return leafs
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetYangName() string { return "entLPMappingEntry" }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) SetParent(parent types.Entity) { entlpmappingentry.parent = parent }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetParent() types.Entity { return entlpmappingentry.parent }

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetParentYangName() string { return "entLPMappingTable" }

// ENTITYMIB_Entaliasmappingtable
// This table contains zero or more rows, representing
// mappings of logical entity and physical component to
// external MIB identifiers.  Each physical port in the system
// may be associated with a mapping to an external identifier,
// which itself is associated with a particular logical
// entity's naming scope.  A 'wildcard' mechanism is provided
// to indicate that an identifier is associated with more than
// one logical entity.
type ENTITYMIB_Entaliasmappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular physical equipment, logical   entity to
    // external identifier binding.  Each logical entity/physical component pair
    // may be associated with one alias mapping.  The logical entity index may
    // also be used as a 'wildcard' (refer to the entAliasLogicalIndexOrZero
    // object DESCRIPTION clause for details.)  Note that only entPhysicalIndex
    // values that represent physical ports (i.e., associated entPhysicalClass
    // value is 'port(10)') are permitted to exist in this table. The type is
    // slice of ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry.
    Entaliasmappingentry []ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetFilter() yfilter.YFilter { return entaliasmappingtable.YFilter }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) SetFilter(yf yfilter.YFilter) { entaliasmappingtable.YFilter = yf }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetGoName(yname string) string {
    if yname == "entAliasMappingEntry" { return "Entaliasmappingentry" }
    return ""
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetSegmentPath() string {
    return "entAliasMappingTable"
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entAliasMappingEntry" {
        for _, c := range entaliasmappingtable.Entaliasmappingentry {
            if entaliasmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry{}
        entaliasmappingtable.Entaliasmappingentry = append(entaliasmappingtable.Entaliasmappingentry, child)
        return &entaliasmappingtable.Entaliasmappingentry[len(entaliasmappingtable.Entaliasmappingentry)-1]
    }
    return nil
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entaliasmappingtable.Entaliasmappingentry {
        children[entaliasmappingtable.Entaliasmappingentry[i].GetSegmentPath()] = &entaliasmappingtable.Entaliasmappingentry[i]
    }
    return children
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetYangName() string { return "entAliasMappingTable" }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) SetParent(parent types.Entity) { entaliasmappingtable.parent = parent }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetParent() types.Entity { return entaliasmappingtable.parent }

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry
// Information about a particular physical equipment, logical
// 
// 
// entity to external identifier binding.  Each logical
// entity/physical component pair may be associated with one
// alias mapping.  The logical entity index may also be used as
// a 'wildcard' (refer to the entAliasLogicalIndexOrZero object
// DESCRIPTION clause for details.)
// 
// Note that only entPhysicalIndex values that represent
// physical ports (i.e., associated entPhysicalClass value is
// 'port(10)') are permitted to exist in this table.
type ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The value of this object identifies the logical
    // entity that defines the naming scope for the associated instance of the
    // 'entAliasMappingIdentifier' object.  If this object has a non-zero value,
    // then it identifies the logical entity named by the same value of
    // entLogicalIndex.  If this object has a value of zero, then the mapping
    // between the physical component and the alias identifier for this
    // entAliasMapping entry is associated with all unspecified logical entities. 
    // That is, a value of zero (the default mapping) identifies any logical
    // entity that does not have an explicit entry in this table for a particular
    // entPhysicalIndex/entAliasMappingIdentifier pair.  For example, to indicate
    // that a particular interface (e.g., physical component 33) is identified by
    // the same value of ifIndex for all logical entities, the following instance
    // might exist:          entAliasMappingIdentifier.33.0 = ifIndex.5  In the
    // event an entPhysicalEntry is associated differently for some logical
    // entities, additional entAliasMapping entries may exist, e.g.:          
    // entAliasMappingIdentifier.33.0 = ifIndex.6        
    // entAliasMappingIdentifier.33.4 =  ifIndex.1        
    // entAliasMappingIdentifier.33.5 =  ifIndex.1        
    // entAliasMappingIdentifier.33.10 = ifIndex.12  Note that entries with
    // non-zero entAliasLogicalIndexOrZero index values have precedence over
    // zero-indexed entries.  In this example, all logical entities except 4, 5,
    // and 10, associate physical entity 33 with ifIndex.6. The type is
    // interface{} with range: 0..2147483647.
    Entaliaslogicalindexorzero interface{}

    // The value of this object identifies a particular conceptual row associated
    // with the indicated entPhysicalIndex and entLogicalIndex pair.  Because only
    // physical ports are modeled in this table, only entries that represent
    // interfaces or ports are allowed.  If an ifEntry exists on behalf of a
    // particular physical port, then this object should identify the associated
    // 'ifEntry'. For repeater ports, the appropriate row in the
    // 'rptrPortGroupTable' should be identified instead.  For example, suppose a
    // physical port was represented by entPhysicalEntry.3, entLogicalEntry.15
    // existed for a repeater, and entLogicalEntry.22 existed for a bridge.  Then
    // there might be two related instances of entAliasMappingIdentifier:   
    // entAliasMappingIdentifier.3.15 == rptrPortGroupIndex.5.2   
    // entAliasMappingIdentifier.3.22 == ifIndex.17 It is possible that other
    // mappings (besides interfaces and repeater ports) may be defined in the
    // future, as required.  Bridge ports are identified by examining the Bridge
    // MIB and appropriate ifEntries associated with each 'dot1dBasePort', and are
    // thus not represented in this table. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Entaliasmappingidentifier interface{}
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetFilter() yfilter.YFilter { return entaliasmappingentry.YFilter }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) SetFilter(yf yfilter.YFilter) { entaliasmappingentry.YFilter = yf }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entAliasLogicalIndexOrZero" { return "Entaliaslogicalindexorzero" }
    if yname == "entAliasMappingIdentifier" { return "Entaliasmappingidentifier" }
    return ""
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetSegmentPath() string {
    return "entAliasMappingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entaliasmappingentry.Entphysicalindex) + "']" + "[entAliasLogicalIndexOrZero='" + fmt.Sprintf("%v", entaliasmappingentry.Entaliaslogicalindexorzero) + "']"
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entaliasmappingentry.Entphysicalindex
    leafs["entAliasLogicalIndexOrZero"] = entaliasmappingentry.Entaliaslogicalindexorzero
    leafs["entAliasMappingIdentifier"] = entaliasmappingentry.Entaliasmappingidentifier
    return leafs
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetYangName() string { return "entAliasMappingEntry" }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) SetParent(parent types.Entity) { entaliasmappingentry.parent = parent }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetParent() types.Entity { return entaliasmappingentry.parent }

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetParentYangName() string { return "entAliasMappingTable" }

// ENTITYMIB_Entphysicalcontainstable
// A table that exposes the container/'containee'
// relationships between physical entities.  This table
// provides all the information found by constructing the
// virtual containment tree for a given entPhysicalTable, but
// in a more direct format.
// 
// In the event a physical entity is contained by more than one
// other physical entity (e.g., double-wide modules), this
// table should include these additional mappings, which cannot
// be represented in the entPhysicalTable virtual containment
// tree.
type ENTITYMIB_Entphysicalcontainstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single container/'containee' relationship. The type is slice of
    // ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry.
    Entphysicalcontainsentry []ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetFilter() yfilter.YFilter { return entphysicalcontainstable.YFilter }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) SetFilter(yf yfilter.YFilter) { entphysicalcontainstable.YFilter = yf }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetGoName(yname string) string {
    if yname == "entPhysicalContainsEntry" { return "Entphysicalcontainsentry" }
    return ""
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetSegmentPath() string {
    return "entPhysicalContainsTable"
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entPhysicalContainsEntry" {
        for _, c := range entphysicalcontainstable.Entphysicalcontainsentry {
            if entphysicalcontainstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry{}
        entphysicalcontainstable.Entphysicalcontainsentry = append(entphysicalcontainstable.Entphysicalcontainsentry, child)
        return &entphysicalcontainstable.Entphysicalcontainsentry[len(entphysicalcontainstable.Entphysicalcontainsentry)-1]
    }
    return nil
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entphysicalcontainstable.Entphysicalcontainsentry {
        children[entphysicalcontainstable.Entphysicalcontainsentry[i].GetSegmentPath()] = &entphysicalcontainstable.Entphysicalcontainsentry[i]
    }
    return children
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetBundleName() string { return "cisco_ios_xe" }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetYangName() string { return "entPhysicalContainsTable" }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) SetParent(parent types.Entity) { entphysicalcontainstable.parent = parent }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetParent() types.Entity { return entphysicalcontainstable.parent }

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetParentYangName() string { return "ENTITY-MIB" }

// ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
// A single container/'containee' relationship.
type ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The value of entPhysicalIndex for the contained
    // physical entity. The type is interface{} with range: 1..2147483647.
    Entphysicalchildindex interface{}
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetFilter() yfilter.YFilter { return entphysicalcontainsentry.YFilter }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) SetFilter(yf yfilter.YFilter) { entphysicalcontainsentry.YFilter = yf }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entPhysicalChildIndex" { return "Entphysicalchildindex" }
    return ""
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetSegmentPath() string {
    return "entPhysicalContainsEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalindex) + "']" + "[entPhysicalChildIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalchildindex) + "']"
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entphysicalcontainsentry.Entphysicalindex
    leafs["entPhysicalChildIndex"] = entphysicalcontainsentry.Entphysicalchildindex
    return leafs
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetBundleName() string { return "cisco_ios_xe" }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetYangName() string { return "entPhysicalContainsEntry" }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) SetParent(parent types.Entity) { entphysicalcontainsentry.parent = parent }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetParent() types.Entity { return entphysicalcontainsentry.parent }

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetParentYangName() string { return "entPhysicalContainsTable" }

