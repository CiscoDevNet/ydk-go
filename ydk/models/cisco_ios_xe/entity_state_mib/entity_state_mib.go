// This MIB defines a state extension to the Entity MIB.
// 
// Copyright (C) The Internet Society 2005.  This version
// of this MIB module is part of RFC 4268; see the RFC
// itself for full legal notices.
package entity_state_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package entity_state_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ENTITY-STATE-MIB ENTITY-STATE-MIB}", reflect.TypeOf(ENTITYSTATEMIB{}))
    ydk.RegisterEntity("ENTITY-STATE-MIB:ENTITY-STATE-MIB", reflect.TypeOf(ENTITYSTATEMIB{}))
}

// ENTITYSTATEMIB
type ENTITYSTATEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table of information about state/status of entities. This is a sparse
    // augment of the entPhysicalTable.  Entries appear in this table for values
    // of entPhysicalClass [RFC4133] that in this implementation are able to
    // report any of the state or status stored in this table.
    EntStateTable ENTITYSTATEMIB_EntStateTable
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetEntityData() *types.CommonEntityData {
    eNTITYSTATEMIB.EntityData.YFilter = eNTITYSTATEMIB.YFilter
    eNTITYSTATEMIB.EntityData.YangName = "ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.BundleName = "cisco_ios_xe"
    eNTITYSTATEMIB.EntityData.ParentYangName = "ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.SegmentPath = "ENTITY-STATE-MIB:ENTITY-STATE-MIB"
    eNTITYSTATEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eNTITYSTATEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eNTITYSTATEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eNTITYSTATEMIB.EntityData.Children = types.NewOrderedMap()
    eNTITYSTATEMIB.EntityData.Children.Append("entStateTable", types.YChild{"EntStateTable", &eNTITYSTATEMIB.EntStateTable})
    eNTITYSTATEMIB.EntityData.Leafs = types.NewOrderedMap()

    eNTITYSTATEMIB.EntityData.YListKeys = []string {}

    return &(eNTITYSTATEMIB.EntityData)
}

// ENTITYSTATEMIB_EntStateTable
// A table of information about state/status of entities.
// This is a sparse augment of the entPhysicalTable.  Entries
// appear in this table for values of
// entPhysicalClass [RFC4133] that in this implementation
// are able to report any of the state or status stored in
// this table.
type ENTITYSTATEMIB_EntStateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State information about this physical entity. The type is slice of
    // ENTITYSTATEMIB_EntStateTable_EntStateEntry.
    EntStateEntry []*ENTITYSTATEMIB_EntStateTable_EntStateEntry
}

func (entStateTable *ENTITYSTATEMIB_EntStateTable) GetEntityData() *types.CommonEntityData {
    entStateTable.EntityData.YFilter = entStateTable.YFilter
    entStateTable.EntityData.YangName = "entStateTable"
    entStateTable.EntityData.BundleName = "cisco_ios_xe"
    entStateTable.EntityData.ParentYangName = "ENTITY-STATE-MIB"
    entStateTable.EntityData.SegmentPath = "entStateTable"
    entStateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entStateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entStateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entStateTable.EntityData.Children = types.NewOrderedMap()
    entStateTable.EntityData.Children.Append("entStateEntry", types.YChild{"EntStateEntry", nil})
    for i := range entStateTable.EntStateEntry {
        entStateTable.EntityData.Children.Append(types.GetSegmentPath(entStateTable.EntStateEntry[i]), types.YChild{"EntStateEntry", entStateTable.EntStateEntry[i]})
    }
    entStateTable.EntityData.Leafs = types.NewOrderedMap()

    entStateTable.EntityData.YListKeys = []string {}

    return &(entStateTable.EntityData)
}

// ENTITYSTATEMIB_EntStateTable_EntStateEntry
// State information about this physical entity.
type ENTITYSTATEMIB_EntStateTable_EntStateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The value of this object is the date and time when the value of any of
    // entStateAdmin, entStateOper, entStateUsage, entStateAlarm, or
    // entStateStandby changed for this entity.  If there has been no change since
    // the last re-initialization of the local system, this object contains the
    // date and time of local system initialization.  If there has been no change
    // since the entity was added to the local system, this object contains the
    // date and time of the insertion. The type is string.
    EntStateLastChanged interface{}

    // The administrative state for this entity.  This object refers to an
    // entities administrative permission to service both other entities within
    // its containment hierarchy as well other users of its services defined by
    // means outside the scope of this MIB.  Setting this object to 'notSupported'
    // will result in an 'inconsistentValue' error.  For entities that do not
    // support administrative state, all set operations will result in an
    // 'inconsistentValue' error.  Some physical entities exhibit only a subset of
    // the remaining administrative state values.  Some entities cannot be locked,
    // and hence this object exhibits only the 'unlocked' state.  Other entities
    // cannot be shutdown gracefully, and hence this object does not exhibit the
    // 'shuttingDown' state.  A value of 'inconsistentValue' will be returned if
    // attempts are made to set this object to values not supported by its
    // administrative model. The type is EntityAdminState.
    EntStateAdmin interface{}

    // The operational state for this entity.  Note that unlike the state model
    // used within the Interfaces MIB [RFC2863], this object does not follow the
    // administrative state.  An administrative state of down does not predict an
    // operational state of disabled.  A value of 'testing' means that entity
    // currently being tested and cannot therefore report whether it is
    // operational or not.  A value of 'disabled' means that an entity is totally
    // inoperable and unable to provide service both to entities within its
    // containment hierarchy, or to other receivers of its service as defined in
    // ways outside the scope of this MIB.  A value of 'enabled' means that an
    // entity is fully or partially operable and able to provide service both to 
    // entities within its containment hierarchy, or to other receivers of its
    // service as defined in ways outside the scope of this MIB.  Note that some
    // implementations may not be able to accurately report entStateOper while the
    // entStateAdmin object has a value other than 'unlocked'. In these cases,
    // this object MUST have a value of 'unknown'. The type is EntityOperState.
    EntStateOper interface{}

    // The usage state for this entity.  This object refers to an entity's ability
    // to service more physical entities in a containment hierarchy.  A value of
    // 'idle' means this entity is able to contain other entities but that no
    // other entity is currently contained within this entity.  A value of
    // 'active' means that at least one entity is contained within this entity,
    // but that it could handle more.  A value of 'busy' means that the entity is
    // unable to handle any additional entities being contained in it.  Some
    // entities will exhibit only a subset of the usage state values.  Entities
    // that are unable to ever service any entities within a containment hierarchy
    // will always have a usage state of 'busy'.  Some entities will only ever be
    // able to support one entity within its containment hierarchy and will
    // therefore only exhibit values of 'idle' and 'busy'. The type is
    // EntityUsageState.
    EntStateUsage interface{}

    // The alarm status for this entity.  It does not include the alarms raised on
    // child components within its containment hierarchy.  A value of 'unknown'
    // means that this entity is  unable to report alarm state.  Note that this
    // differs from 'indeterminate', which means that alarm state is supported and
    // there are alarms against this entity, but the severity of some of the
    // alarms is not known.  If no bits are set, then this entity supports
    // reporting of alarms, but there are currently no active alarms against this
    // entity. The type is map[string]bool.
    EntStateAlarm interface{}

    // The standby status for this entity.  Some entities will exhibit only a
    // subset of the remaining standby state values.  If this entity cannot
    // operate in a standby role, the value of this object will always be
    // 'providingService'. The type is EntityStandbyStatus.
    EntStateStandby interface{}
}

func (entStateEntry *ENTITYSTATEMIB_EntStateTable_EntStateEntry) GetEntityData() *types.CommonEntityData {
    entStateEntry.EntityData.YFilter = entStateEntry.YFilter
    entStateEntry.EntityData.YangName = "entStateEntry"
    entStateEntry.EntityData.BundleName = "cisco_ios_xe"
    entStateEntry.EntityData.ParentYangName = "entStateTable"
    entStateEntry.EntityData.SegmentPath = "entStateEntry" + types.AddKeyToken(entStateEntry.EntPhysicalIndex, "entPhysicalIndex")
    entStateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    entStateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    entStateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    entStateEntry.EntityData.Children = types.NewOrderedMap()
    entStateEntry.EntityData.Leafs = types.NewOrderedMap()
    entStateEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", entStateEntry.EntPhysicalIndex})
    entStateEntry.EntityData.Leafs.Append("entStateLastChanged", types.YLeaf{"EntStateLastChanged", entStateEntry.EntStateLastChanged})
    entStateEntry.EntityData.Leafs.Append("entStateAdmin", types.YLeaf{"EntStateAdmin", entStateEntry.EntStateAdmin})
    entStateEntry.EntityData.Leafs.Append("entStateOper", types.YLeaf{"EntStateOper", entStateEntry.EntStateOper})
    entStateEntry.EntityData.Leafs.Append("entStateUsage", types.YLeaf{"EntStateUsage", entStateEntry.EntStateUsage})
    entStateEntry.EntityData.Leafs.Append("entStateAlarm", types.YLeaf{"EntStateAlarm", entStateEntry.EntStateAlarm})
    entStateEntry.EntityData.Leafs.Append("entStateStandby", types.YLeaf{"EntStateStandby", entStateEntry.EntStateStandby})

    entStateEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(entStateEntry.EntityData)
}

