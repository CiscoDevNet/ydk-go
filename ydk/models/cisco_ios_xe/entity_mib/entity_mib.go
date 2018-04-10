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
    EntityData types.CommonEntityData
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

func (eNTITYMIB *ENTITYMIB) GetEntityData() *types.CommonEntityData {
    eNTITYMIB.EntityData.YFilter = eNTITYMIB.YFilter
    eNTITYMIB.EntityData.YangName = "ENTITY-MIB"
    eNTITYMIB.EntityData.BundleName = "cisco_ios_xe"
    eNTITYMIB.EntityData.ParentYangName = "ENTITY-MIB"
    eNTITYMIB.EntityData.SegmentPath = "ENTITY-MIB:ENTITY-MIB"
    eNTITYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eNTITYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eNTITYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eNTITYMIB.EntityData.Children = make(map[string]types.YChild)
    eNTITYMIB.EntityData.Children["entityGeneral"] = types.YChild{"Entitygeneral", &eNTITYMIB.Entitygeneral}
    eNTITYMIB.EntityData.Children["entPhysicalTable"] = types.YChild{"Entphysicaltable", &eNTITYMIB.Entphysicaltable}
    eNTITYMIB.EntityData.Children["entLogicalTable"] = types.YChild{"Entlogicaltable", &eNTITYMIB.Entlogicaltable}
    eNTITYMIB.EntityData.Children["entLPMappingTable"] = types.YChild{"Entlpmappingtable", &eNTITYMIB.Entlpmappingtable}
    eNTITYMIB.EntityData.Children["entAliasMappingTable"] = types.YChild{"Entaliasmappingtable", &eNTITYMIB.Entaliasmappingtable}
    eNTITYMIB.EntityData.Children["entPhysicalContainsTable"] = types.YChild{"Entphysicalcontainstable", &eNTITYMIB.Entphysicalcontainstable}
    eNTITYMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eNTITYMIB.EntityData)
}

// ENTITYMIB_Entitygeneral
type ENTITYMIB_Entitygeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time a conceptual row is created, modified,
    // or deleted in any of these tables:         - entPhysicalTable         -
    // entLogicalTable         - entLPMappingTable         - entAliasMappingTable 
    // - entPhysicalContainsTable. The type is interface{} with range:
    // 0..4294967295.
    Entlastchangetime interface{}
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetEntityData() *types.CommonEntityData {
    entitygeneral.EntityData.YFilter = entitygeneral.YFilter
    entitygeneral.EntityData.YangName = "entityGeneral"
    entitygeneral.EntityData.BundleName = "cisco_ios_xe"
    entitygeneral.EntityData.ParentYangName = "ENTITY-MIB"
    entitygeneral.EntityData.SegmentPath = "entityGeneral"
    entitygeneral.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entitygeneral.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entitygeneral.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entitygeneral.EntityData.Children = make(map[string]types.YChild)
    entitygeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    entitygeneral.EntityData.Leafs["entLastChangeTime"] = types.YLeaf{"Entlastchangetime", entitygeneral.Entlastchangetime}
    return &(entitygeneral.EntityData)
}

// ENTITYMIB_Entphysicaltable
// This table contains one row per physical entity.  There is
// always at least one row for an 'overall' physical entity.
type ENTITYMIB_Entphysicaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular physical entity.  Each entry provides
    // objects (entPhysicalDescr, entPhysicalVendorType, and entPhysicalClass) to
    // help an NMS identify and characterize the entry, and objects
    // (entPhysicalContainedIn and entPhysicalParentRelPos) to help an NMS relate
    // the particular entry to other entries in this table. The type is slice of
    // ENTITYMIB_Entphysicaltable_Entphysicalentry.
    Entphysicalentry []ENTITYMIB_Entphysicaltable_Entphysicalentry
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetEntityData() *types.CommonEntityData {
    entphysicaltable.EntityData.YFilter = entphysicaltable.YFilter
    entphysicaltable.EntityData.YangName = "entPhysicalTable"
    entphysicaltable.EntityData.BundleName = "cisco_ios_xe"
    entphysicaltable.EntityData.ParentYangName = "ENTITY-MIB"
    entphysicaltable.EntityData.SegmentPath = "entPhysicalTable"
    entphysicaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysicaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysicaltable.EntityData.Children = make(map[string]types.YChild)
    entphysicaltable.EntityData.Children["entPhysicalEntry"] = types.YChild{"Entphysicalentry", nil}
    for i := range entphysicaltable.Entphysicalentry {
        entphysicaltable.EntityData.Children[types.GetSegmentPath(&entphysicaltable.Entphysicalentry[i])] = types.YChild{"Entphysicalentry", &entphysicaltable.Entphysicalentry[i]}
    }
    entphysicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entphysicaltable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetEntityData() *types.CommonEntityData {
    entphysicalentry.EntityData.YFilter = entphysicalentry.YFilter
    entphysicalentry.EntityData.YangName = "entPhysicalEntry"
    entphysicalentry.EntityData.BundleName = "cisco_ios_xe"
    entphysicalentry.EntityData.ParentYangName = "entPhysicalTable"
    entphysicalentry.EntityData.SegmentPath = "entPhysicalEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalentry.Entphysicalindex) + "']"
    entphysicalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysicalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysicalentry.EntityData.Children = make(map[string]types.YChild)
    entphysicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entphysicalentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entphysicalentry.Entphysicalindex}
    entphysicalentry.EntityData.Leafs["entPhysicalDescr"] = types.YLeaf{"Entphysicaldescr", entphysicalentry.Entphysicaldescr}
    entphysicalentry.EntityData.Leafs["entPhysicalVendorType"] = types.YLeaf{"Entphysicalvendortype", entphysicalentry.Entphysicalvendortype}
    entphysicalentry.EntityData.Leafs["entPhysicalContainedIn"] = types.YLeaf{"Entphysicalcontainedin", entphysicalentry.Entphysicalcontainedin}
    entphysicalentry.EntityData.Leafs["entPhysicalClass"] = types.YLeaf{"Entphysicalclass", entphysicalentry.Entphysicalclass}
    entphysicalentry.EntityData.Leafs["entPhysicalParentRelPos"] = types.YLeaf{"Entphysicalparentrelpos", entphysicalentry.Entphysicalparentrelpos}
    entphysicalentry.EntityData.Leafs["entPhysicalName"] = types.YLeaf{"Entphysicalname", entphysicalentry.Entphysicalname}
    entphysicalentry.EntityData.Leafs["entPhysicalHardwareRev"] = types.YLeaf{"Entphysicalhardwarerev", entphysicalentry.Entphysicalhardwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalFirmwareRev"] = types.YLeaf{"Entphysicalfirmwarerev", entphysicalentry.Entphysicalfirmwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalSoftwareRev"] = types.YLeaf{"Entphysicalsoftwarerev", entphysicalentry.Entphysicalsoftwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalSerialNum"] = types.YLeaf{"Entphysicalserialnum", entphysicalentry.Entphysicalserialnum}
    entphysicalentry.EntityData.Leafs["entPhysicalMfgName"] = types.YLeaf{"Entphysicalmfgname", entphysicalentry.Entphysicalmfgname}
    entphysicalentry.EntityData.Leafs["entPhysicalModelName"] = types.YLeaf{"Entphysicalmodelname", entphysicalentry.Entphysicalmodelname}
    entphysicalentry.EntityData.Leafs["entPhysicalAlias"] = types.YLeaf{"Entphysicalalias", entphysicalentry.Entphysicalalias}
    entphysicalentry.EntityData.Leafs["entPhysicalAssetID"] = types.YLeaf{"Entphysicalassetid", entphysicalentry.Entphysicalassetid}
    entphysicalentry.EntityData.Leafs["entPhysicalIsFRU"] = types.YLeaf{"Entphysicalisfru", entphysicalentry.Entphysicalisfru}
    entphysicalentry.EntityData.Leafs["entPhysicalMfgDate"] = types.YLeaf{"Entphysicalmfgdate", entphysicalentry.Entphysicalmfgdate}
    entphysicalentry.EntityData.Leafs["entPhysicalUris"] = types.YLeaf{"Entphysicaluris", entphysicalentry.Entphysicaluris}
    entphysicalentry.EntityData.Leafs["ceEntPhysicalSecondSerialNum"] = types.YLeaf{"Ceentphysicalsecondserialnum", entphysicalentry.Ceentphysicalsecondserialnum}
    return &(entphysicalentry.EntityData)
}

// ENTITYMIB_Entlogicaltable
// This table contains one row per logical entity.  For agents
// that implement more than one naming scope, at least one
// entry must exist.  Agents which instantiate all MIB objects
// within a single naming scope are not required to implement
// this table.
type ENTITYMIB_Entlogicaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular logical entity.  Entities may be managed by
    // this agent or other SNMP agents (possibly) in the same chassis. The type is
    // slice of ENTITYMIB_Entlogicaltable_Entlogicalentry.
    Entlogicalentry []ENTITYMIB_Entlogicaltable_Entlogicalentry
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetEntityData() *types.CommonEntityData {
    entlogicaltable.EntityData.YFilter = entlogicaltable.YFilter
    entlogicaltable.EntityData.YangName = "entLogicalTable"
    entlogicaltable.EntityData.BundleName = "cisco_ios_xe"
    entlogicaltable.EntityData.ParentYangName = "ENTITY-MIB"
    entlogicaltable.EntityData.SegmentPath = "entLogicalTable"
    entlogicaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entlogicaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entlogicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entlogicaltable.EntityData.Children = make(map[string]types.YChild)
    entlogicaltable.EntityData.Children["entLogicalEntry"] = types.YChild{"Entlogicalentry", nil}
    for i := range entlogicaltable.Entlogicalentry {
        entlogicaltable.EntityData.Children[types.GetSegmentPath(&entlogicaltable.Entlogicalentry[i])] = types.YChild{"Entlogicalentry", &entlogicaltable.Entlogicalentry[i]}
    }
    entlogicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entlogicaltable.EntityData)
}

// ENTITYMIB_Entlogicaltable_Entlogicalentry
// Information about a particular logical entity.  Entities
// may be managed by this agent or other SNMP agents (possibly)
// in the same chassis.
type ENTITYMIB_Entlogicaltable_Entlogicalentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetEntityData() *types.CommonEntityData {
    entlogicalentry.EntityData.YFilter = entlogicalentry.YFilter
    entlogicalentry.EntityData.YangName = "entLogicalEntry"
    entlogicalentry.EntityData.BundleName = "cisco_ios_xe"
    entlogicalentry.EntityData.ParentYangName = "entLogicalTable"
    entlogicalentry.EntityData.SegmentPath = "entLogicalEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlogicalentry.Entlogicalindex) + "']"
    entlogicalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entlogicalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entlogicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entlogicalentry.EntityData.Children = make(map[string]types.YChild)
    entlogicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entlogicalentry.EntityData.Leafs["entLogicalIndex"] = types.YLeaf{"Entlogicalindex", entlogicalentry.Entlogicalindex}
    entlogicalentry.EntityData.Leafs["entLogicalDescr"] = types.YLeaf{"Entlogicaldescr", entlogicalentry.Entlogicaldescr}
    entlogicalentry.EntityData.Leafs["entLogicalType"] = types.YLeaf{"Entlogicaltype", entlogicalentry.Entlogicaltype}
    entlogicalentry.EntityData.Leafs["entLogicalCommunity"] = types.YLeaf{"Entlogicalcommunity", entlogicalentry.Entlogicalcommunity}
    entlogicalentry.EntityData.Leafs["entLogicalTAddress"] = types.YLeaf{"Entlogicaltaddress", entlogicalentry.Entlogicaltaddress}
    entlogicalentry.EntityData.Leafs["entLogicalTDomain"] = types.YLeaf{"Entlogicaltdomain", entlogicalentry.Entlogicaltdomain}
    entlogicalentry.EntityData.Leafs["entLogicalContextEngineID"] = types.YLeaf{"Entlogicalcontextengineid", entlogicalentry.Entlogicalcontextengineid}
    entlogicalentry.EntityData.Leafs["entLogicalContextName"] = types.YLeaf{"Entlogicalcontextname", entlogicalentry.Entlogicalcontextname}
    return &(entlogicalentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular logical entity to physical equipment
    // association.  Note that the nature of the association is not specifically
    // identified in this entry. It is expected that sufficient information exists
    // in the MIBs used to manage a particular logical entity to infer how
    // physical component information is utilized. The type is slice of
    // ENTITYMIB_Entlpmappingtable_Entlpmappingentry.
    Entlpmappingentry []ENTITYMIB_Entlpmappingtable_Entlpmappingentry
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetEntityData() *types.CommonEntityData {
    entlpmappingtable.EntityData.YFilter = entlpmappingtable.YFilter
    entlpmappingtable.EntityData.YangName = "entLPMappingTable"
    entlpmappingtable.EntityData.BundleName = "cisco_ios_xe"
    entlpmappingtable.EntityData.ParentYangName = "ENTITY-MIB"
    entlpmappingtable.EntityData.SegmentPath = "entLPMappingTable"
    entlpmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entlpmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entlpmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entlpmappingtable.EntityData.Children = make(map[string]types.YChild)
    entlpmappingtable.EntityData.Children["entLPMappingEntry"] = types.YChild{"Entlpmappingentry", nil}
    for i := range entlpmappingtable.Entlpmappingentry {
        entlpmappingtable.EntityData.Children[types.GetSegmentPath(&entlpmappingtable.Entlpmappingentry[i])] = types.YChild{"Entlpmappingentry", &entlpmappingtable.Entlpmappingentry[i]}
    }
    entlpmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entlpmappingtable.EntityData)
}

// ENTITYMIB_Entlpmappingtable_Entlpmappingentry
// Information about a particular logical entity to physical
// equipment association.  Note that the nature of the
// association is not specifically identified in this entry.
// It is expected that sufficient information exists in the
// MIBs used to manage a particular logical entity to infer how
// physical component information is utilized.
type ENTITYMIB_Entlpmappingtable_Entlpmappingentry struct {
    EntityData types.CommonEntityData
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

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetEntityData() *types.CommonEntityData {
    entlpmappingentry.EntityData.YFilter = entlpmappingentry.YFilter
    entlpmappingentry.EntityData.YangName = "entLPMappingEntry"
    entlpmappingentry.EntityData.BundleName = "cisco_ios_xe"
    entlpmappingentry.EntityData.ParentYangName = "entLPMappingTable"
    entlpmappingentry.EntityData.SegmentPath = "entLPMappingEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlogicalindex) + "']" + "[entLPPhysicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlpphysicalindex) + "']"
    entlpmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entlpmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entlpmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entlpmappingentry.EntityData.Children = make(map[string]types.YChild)
    entlpmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entlpmappingentry.EntityData.Leafs["entLogicalIndex"] = types.YLeaf{"Entlogicalindex", entlpmappingentry.Entlogicalindex}
    entlpmappingentry.EntityData.Leafs["entLPPhysicalIndex"] = types.YLeaf{"Entlpphysicalindex", entlpmappingentry.Entlpphysicalindex}
    return &(entlpmappingentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetEntityData() *types.CommonEntityData {
    entaliasmappingtable.EntityData.YFilter = entaliasmappingtable.YFilter
    entaliasmappingtable.EntityData.YangName = "entAliasMappingTable"
    entaliasmappingtable.EntityData.BundleName = "cisco_ios_xe"
    entaliasmappingtable.EntityData.ParentYangName = "ENTITY-MIB"
    entaliasmappingtable.EntityData.SegmentPath = "entAliasMappingTable"
    entaliasmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entaliasmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entaliasmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entaliasmappingtable.EntityData.Children = make(map[string]types.YChild)
    entaliasmappingtable.EntityData.Children["entAliasMappingEntry"] = types.YChild{"Entaliasmappingentry", nil}
    for i := range entaliasmappingtable.Entaliasmappingentry {
        entaliasmappingtable.EntityData.Children[types.GetSegmentPath(&entaliasmappingtable.Entaliasmappingentry[i])] = types.YChild{"Entaliasmappingentry", &entaliasmappingtable.Entaliasmappingentry[i]}
    }
    entaliasmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entaliasmappingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Entaliasmappingidentifier interface{}
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetEntityData() *types.CommonEntityData {
    entaliasmappingentry.EntityData.YFilter = entaliasmappingentry.YFilter
    entaliasmappingentry.EntityData.YangName = "entAliasMappingEntry"
    entaliasmappingentry.EntityData.BundleName = "cisco_ios_xe"
    entaliasmappingentry.EntityData.ParentYangName = "entAliasMappingTable"
    entaliasmappingentry.EntityData.SegmentPath = "entAliasMappingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entaliasmappingentry.Entphysicalindex) + "']" + "[entAliasLogicalIndexOrZero='" + fmt.Sprintf("%v", entaliasmappingentry.Entaliaslogicalindexorzero) + "']"
    entaliasmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entaliasmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entaliasmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entaliasmappingentry.EntityData.Children = make(map[string]types.YChild)
    entaliasmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entaliasmappingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entaliasmappingentry.Entphysicalindex}
    entaliasmappingentry.EntityData.Leafs["entAliasLogicalIndexOrZero"] = types.YLeaf{"Entaliaslogicalindexorzero", entaliasmappingentry.Entaliaslogicalindexorzero}
    entaliasmappingentry.EntityData.Leafs["entAliasMappingIdentifier"] = types.YLeaf{"Entaliasmappingidentifier", entaliasmappingentry.Entaliasmappingidentifier}
    return &(entaliasmappingentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single container/'containee' relationship. The type is slice of
    // ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry.
    Entphysicalcontainsentry []ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetEntityData() *types.CommonEntityData {
    entphysicalcontainstable.EntityData.YFilter = entphysicalcontainstable.YFilter
    entphysicalcontainstable.EntityData.YangName = "entPhysicalContainsTable"
    entphysicalcontainstable.EntityData.BundleName = "cisco_ios_xe"
    entphysicalcontainstable.EntityData.ParentYangName = "ENTITY-MIB"
    entphysicalcontainstable.EntityData.SegmentPath = "entPhysicalContainsTable"
    entphysicalcontainstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysicalcontainstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysicalcontainstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysicalcontainstable.EntityData.Children = make(map[string]types.YChild)
    entphysicalcontainstable.EntityData.Children["entPhysicalContainsEntry"] = types.YChild{"Entphysicalcontainsentry", nil}
    for i := range entphysicalcontainstable.Entphysicalcontainsentry {
        entphysicalcontainstable.EntityData.Children[types.GetSegmentPath(&entphysicalcontainstable.Entphysicalcontainsentry[i])] = types.YChild{"Entphysicalcontainsentry", &entphysicalcontainstable.Entphysicalcontainsentry[i]}
    }
    entphysicalcontainstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entphysicalcontainstable.EntityData)
}

// ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
// A single container/'containee' relationship.
type ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The value of entPhysicalIndex for the contained
    // physical entity. The type is interface{} with range: 1..2147483647.
    Entphysicalchildindex interface{}
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetEntityData() *types.CommonEntityData {
    entphysicalcontainsentry.EntityData.YFilter = entphysicalcontainsentry.YFilter
    entphysicalcontainsentry.EntityData.YangName = "entPhysicalContainsEntry"
    entphysicalcontainsentry.EntityData.BundleName = "cisco_ios_xe"
    entphysicalcontainsentry.EntityData.ParentYangName = "entPhysicalContainsTable"
    entphysicalcontainsentry.EntityData.SegmentPath = "entPhysicalContainsEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalindex) + "']" + "[entPhysicalChildIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalchildindex) + "']"
    entphysicalcontainsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entphysicalcontainsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entphysicalcontainsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entphysicalcontainsentry.EntityData.Children = make(map[string]types.YChild)
    entphysicalcontainsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entphysicalcontainsentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entphysicalcontainsentry.Entphysicalindex}
    entphysicalcontainsentry.EntityData.Leafs["entPhysicalChildIndex"] = types.YLeaf{"Entphysicalchildindex", entphysicalcontainsentry.Entphysicalchildindex}
    return &(entphysicalcontainsentry.EntityData)
}

