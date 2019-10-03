// A single SNMP agent sometimes needs to support multiple
// instances of the same MIB module, and does so through the
// use of multiple SNMP contexts.  This typically occurs because
// the technology has evolved to have extra dimension(s), i.e.,
// one or more extra data and/or identifier values which are
// different in the different contexts, but were not defined in
// INDEX clause(s) of the original MIB module.  In such cases,
// network management applications need to know the specific
// data/identifier values in each context, and this MIB module
// provides mapping tables which contain that information.
// 
// Within a network there can be multiple Virtual Private
// Networks (VPNs) configured using Virtual Routing and
// Forwarding Instances (VRFs). Within a VPN there can be
// multiple topologies when Multi-topology Routing (MTR) is
// used. Also, Interior Gateway Protocols (IGPs) can have
// multiple protocol instances running on the device.
// A network can have multiple broadcast domains configured
// using Bridge Domain Identifiers.
// 
// With MTR routing, VRFs, and Bridge domains, a router now 
// needs to support multiple instances of several existing 
// MIB modules, and this can be achieved if the router's SNMP 
// agent provides access to each instance of the same MIB module 
// via a different SNMP context (see Section 3.1.1 of RFC 3411).
// For MTR routing, VRFs, and Bridge domains, a different SNMP 
// context is needed depending on one or more of the following: 
// the VRF, the topology-identifier, the routing protocol instance,
// and the bridge domain identifier.
// In other words, the router's management information can be
// accessed through multiple SNMP contexts where each such
// context represents a specific VRF, a specific
// topology-identifier, a specific routing protocol instance 
// and/or a bridge domain identifier. This MIB module provides 
// a mapping of each such SNMP context to the corresponding VRF,
// the corresponding topology, the corresponding routing protocol 
// instance, and the corresponding bridge domain identifier.
// Some SNMP contexts are independent of VRFs, independent of
// a topology, independent of a routing protocol instance, or 
// independent of a bridge domain and in such a case, the mapping
// is to the zero length string.
// 
// With the Cisco package licensing strategy, the features 
// available in the image are grouped into multiple packages 
// and each packages can be managed to operate at different 
// feature levels based on the available license. This MIB 
// module provides option to associate an SNMP context to a 
// feature package group. This will allow manageability of 
// license MIB objects specific to a feature package group.
// 
// As technology evolves more we may need additional
// identifiers to identify the context. Then we would need
// to add those additional identifiers into the mapping.
package cisco_context_mapping_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_context_mapping_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CONTEXT-MAPPING-MIB CISCO-CONTEXT-MAPPING-MIB}", reflect.TypeOf(CISCOCONTEXTMAPPINGMIB{}))
    ydk.RegisterEntity("CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB", reflect.TypeOf(CISCOCONTEXTMAPPINGMIB{}))
}

// CISCOCONTEXTMAPPINGMIB
type CISCOCONTEXTMAPPINGMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains information on which cContextMappingVacmContextName is
    // mapped to which VRF, topology, and routing protocol instance.  This table
    // is indexed by SNMP VACM context.  Configuring a row in this table for an
    // SNMP context does not require that the context be already defined, i.e., a
    // row can be created in this table for a context before the corresponding row
    // is created in RFC 3415's vacmContextTable.  To create a row in this table,
    // a manager must set cContextMappingRowStatus to either 'createAndGo' or
    // 'createAndWait'.  To delete a row in this table, a manager must set
    // cContextMappingRowStatus to 'destroy'.
    CContextMappingTable CISCOCONTEXTMAPPINGMIB_CContextMappingTable

    // This table contains information on which cContextMappingVacmContextName is
    // mapped to which bridge domain.  A Bridge Domain is one of the means by
    // which it is possible  to define an Ethernet broadcast domain on a bridging
    // device.  A network can have multiple broadcast domains configured. This
    // table helps the network management personnel to find  out the  details of
    // various broadcast domains configured  in the network.  An entry need to
    // exist in cContextMappingTable, to create  an entry in this table.
    CContextMappingBridgeDomainTable CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable

    // This table contains information on mapping between
    // cContextMappingVacmContextName and bridge instance.  Bridge instance is an
    // instance of a physical or logical  bridge which has unique bridge-id.  If
    // an entry is deleted from cContextMappingTable, the corresponding entry in
    // this table will also get deleted.  If an entry needs to be created in this
    // table, the corresponding entry must exist in cContextMappingTable.
    CContextMappingBridgeInstanceTable CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable

    // This table contains information on which cContextMappingVacmContextName is
    // mapped to which License Group. Group level licensing is used where each
    // Technology Package is enabled via a License.
    CContextMappingLicenseGroupTable CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetEntityData() *types.CommonEntityData {
    cISCOCONTEXTMAPPINGMIB.EntityData.YFilter = cISCOCONTEXTMAPPINGMIB.YFilter
    cISCOCONTEXTMAPPINGMIB.EntityData.YangName = "CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCONTEXTMAPPINGMIB.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.SegmentPath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.AbsolutePath = cISCOCONTEXTMAPPINGMIB.EntityData.SegmentPath
    cISCOCONTEXTMAPPINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCONTEXTMAPPINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCONTEXTMAPPINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCONTEXTMAPPINGMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCONTEXTMAPPINGMIB.EntityData.Children.Append("cContextMappingTable", types.YChild{"CContextMappingTable", &cISCOCONTEXTMAPPINGMIB.CContextMappingTable})
    cISCOCONTEXTMAPPINGMIB.EntityData.Children.Append("cContextMappingBridgeDomainTable", types.YChild{"CContextMappingBridgeDomainTable", &cISCOCONTEXTMAPPINGMIB.CContextMappingBridgeDomainTable})
    cISCOCONTEXTMAPPINGMIB.EntityData.Children.Append("cContextMappingBridgeInstanceTable", types.YChild{"CContextMappingBridgeInstanceTable", &cISCOCONTEXTMAPPINGMIB.CContextMappingBridgeInstanceTable})
    cISCOCONTEXTMAPPINGMIB.EntityData.Children.Append("cContextMappingLicenseGroupTable", types.YChild{"CContextMappingLicenseGroupTable", &cISCOCONTEXTMAPPINGMIB.CContextMappingLicenseGroupTable})
    cISCOCONTEXTMAPPINGMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCONTEXTMAPPINGMIB.EntityData.YListKeys = []string {}

    return &(cISCOCONTEXTMAPPINGMIB.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingTable
// This table contains information on which
// cContextMappingVacmContextName is mapped to
// which VRF, topology, and routing protocol instance.
// 
// This table is indexed by SNMP VACM context.
// 
// Configuring a row in this table for an SNMP context
// does not require that the context be already defined,
// i.e., a row can be created in this table for a context
// before the corresponding row is created in RFC 3415's
// vacmContextTable.
// 
// To create a row in this table, a manager must set
// cContextMappingRowStatus to either 'createAndGo' or
// 'createAndWait'.
// 
// To delete a row in this table, a manager must set
// cContextMappingRowStatus to 'destroy'.
type CISCOCONTEXTMAPPINGMIB_CContextMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the corresponding VRF, the corresponding topology, and the corresponding
    // routing protocol instance. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry.
    CContextMappingEntry []*CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry
}

func (cContextMappingTable *CISCOCONTEXTMAPPINGMIB_CContextMappingTable) GetEntityData() *types.CommonEntityData {
    cContextMappingTable.EntityData.YFilter = cContextMappingTable.YFilter
    cContextMappingTable.EntityData.YangName = "cContextMappingTable"
    cContextMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingTable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cContextMappingTable.EntityData.SegmentPath = "cContextMappingTable"
    cContextMappingTable.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/" + cContextMappingTable.EntityData.SegmentPath
    cContextMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingTable.EntityData.Children = types.NewOrderedMap()
    cContextMappingTable.EntityData.Children.Append("cContextMappingEntry", types.YChild{"CContextMappingEntry", nil})
    for i := range cContextMappingTable.CContextMappingEntry {
        cContextMappingTable.EntityData.Children.Append(types.GetSegmentPath(cContextMappingTable.CContextMappingEntry[i]), types.YChild{"CContextMappingEntry", cContextMappingTable.CContextMappingEntry[i]})
    }
    cContextMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cContextMappingTable.EntityData.YListKeys = []string {}

    return &(cContextMappingTable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry
// Information relating to a single mapping of
// cContextMappingVacmContextName to the corresponding VRF,
// the corresponding topology, and the corresponding routing
// protocol instance.
type CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The vacmContextName given to the SNMP context. 
    // This is a human readable name identifying a particular SNMP VACM context at
    // a particular SNMP entity. The empty contextName (zero length) represents
    // the default context. The type is string with length: 0..32.
    CContextMappingVacmContextName interface{}

    // The value of an instance of this object identifies the name given to the
    // VRF to which the SNMP context is mapped to.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this VRF.  When the value of this object is
    // the zero length string it indicates that the SNMP context is independent of
    // any VRF. The type is string with length: 0..32.
    CContextMappingVrfName interface{}

    // The value of an instance of this object identifies the name given to the
    // topology to which the SNMP context is mapped to.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this topology.  When the value of this object
    // is the zero length string it indicates that the SNMP context is independent
    // of any topology. The type is string with length: 0..32.
    CContextMappingTopologyName interface{}

    // The value of an instance of this object identifies the name given to the
    // protocol instance to which the SNMP context is mapped to.  This is
    // typically a human-readable string. This is the same ASCII string used in
    // the router's console interface to refer to this protocol instance.  When
    // the value of this object is the zero length string it indicates that the
    // SNMP context is independent of any protocol instance. The type is string
    // with length: 0..32.
    CContextMappingProtoInstName interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    CContextMappingStorageType interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    CContextMappingRowStatus interface{}
}

func (cContextMappingEntry *CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry) GetEntityData() *types.CommonEntityData {
    cContextMappingEntry.EntityData.YFilter = cContextMappingEntry.YFilter
    cContextMappingEntry.EntityData.YangName = "cContextMappingEntry"
    cContextMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingEntry.EntityData.ParentYangName = "cContextMappingTable"
    cContextMappingEntry.EntityData.SegmentPath = "cContextMappingEntry" + types.AddKeyToken(cContextMappingEntry.CContextMappingVacmContextName, "cContextMappingVacmContextName")
    cContextMappingEntry.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/cContextMappingTable/" + cContextMappingEntry.EntityData.SegmentPath
    cContextMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingEntry.EntityData.Children = types.NewOrderedMap()
    cContextMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingVacmContextName", types.YLeaf{"CContextMappingVacmContextName", cContextMappingEntry.CContextMappingVacmContextName})
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingVrfName", types.YLeaf{"CContextMappingVrfName", cContextMappingEntry.CContextMappingVrfName})
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingTopologyName", types.YLeaf{"CContextMappingTopologyName", cContextMappingEntry.CContextMappingTopologyName})
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingProtoInstName", types.YLeaf{"CContextMappingProtoInstName", cContextMappingEntry.CContextMappingProtoInstName})
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingStorageType", types.YLeaf{"CContextMappingStorageType", cContextMappingEntry.CContextMappingStorageType})
    cContextMappingEntry.EntityData.Leafs.Append("cContextMappingRowStatus", types.YLeaf{"CContextMappingRowStatus", cContextMappingEntry.CContextMappingRowStatus})

    cContextMappingEntry.EntityData.YListKeys = []string {"CContextMappingVacmContextName"}

    return &(cContextMappingEntry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable
// This table contains information on which
// cContextMappingVacmContextName is mapped to
// which bridge domain.
// 
// A Bridge Domain is one of the means by which it is possible 
// to define an Ethernet broadcast domain on a bridging device. 
// A network can have multiple broadcast domains configured.
// This table helps the network management personnel to find 
// out the  details of various broadcast domains configured 
// in the network.
// 
// An entry need to exist in cContextMappingTable, to create 
// an entry in this table.
type CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the  corresponding bridge domain.  To create a row in this table, a
    // manager must set cContextMappingBridgeDomainRowStatus to either 
    // 'createAndGo' or 'createAndWait'.  To delete a row in this table, a manager
    // must set cContextMappingBridgeDomainRowStatus to 'destroy'. The type is
    // slice of
    // CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable_CContextMappingBridgeDomainEntry.
    CContextMappingBridgeDomainEntry []*CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable_CContextMappingBridgeDomainEntry
}

func (cContextMappingBridgeDomainTable *CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable) GetEntityData() *types.CommonEntityData {
    cContextMappingBridgeDomainTable.EntityData.YFilter = cContextMappingBridgeDomainTable.YFilter
    cContextMappingBridgeDomainTable.EntityData.YangName = "cContextMappingBridgeDomainTable"
    cContextMappingBridgeDomainTable.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingBridgeDomainTable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cContextMappingBridgeDomainTable.EntityData.SegmentPath = "cContextMappingBridgeDomainTable"
    cContextMappingBridgeDomainTable.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/" + cContextMappingBridgeDomainTable.EntityData.SegmentPath
    cContextMappingBridgeDomainTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingBridgeDomainTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingBridgeDomainTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingBridgeDomainTable.EntityData.Children = types.NewOrderedMap()
    cContextMappingBridgeDomainTable.EntityData.Children.Append("cContextMappingBridgeDomainEntry", types.YChild{"CContextMappingBridgeDomainEntry", nil})
    for i := range cContextMappingBridgeDomainTable.CContextMappingBridgeDomainEntry {
        cContextMappingBridgeDomainTable.EntityData.Children.Append(types.GetSegmentPath(cContextMappingBridgeDomainTable.CContextMappingBridgeDomainEntry[i]), types.YChild{"CContextMappingBridgeDomainEntry", cContextMappingBridgeDomainTable.CContextMappingBridgeDomainEntry[i]})
    }
    cContextMappingBridgeDomainTable.EntityData.Leafs = types.NewOrderedMap()

    cContextMappingBridgeDomainTable.EntityData.YListKeys = []string {}

    return &(cContextMappingBridgeDomainTable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable_CContextMappingBridgeDomainEntry
// Information relating to a single mapping of
// cContextMappingVacmContextName to the 
// corresponding bridge domain.
// 
// To create a row in this table, a manager must set
// cContextMappingBridgeDomainRowStatus to either 
// 'createAndGo' or 'createAndWait'.
// 
// To delete a row in this table, a manager must set
// cContextMappingBridgeDomainRowStatus to 'destroy'.
type CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable_CContextMappingBridgeDomainEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry_CContextMappingVacmContextName
    CContextMappingVacmContextName interface{}

    // The value of an instance of this object identifies the bridge domain to
    // which the SNMP context is  mapped to. The type is interface{} with range:
    // 1..65535.
    CContextMappingBridgeDomainIdentifier interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    CContextMappingBridgeDomainStorageType interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    CContextMappingBridgeDomainRowStatus interface{}
}

func (cContextMappingBridgeDomainEntry *CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeDomainTable_CContextMappingBridgeDomainEntry) GetEntityData() *types.CommonEntityData {
    cContextMappingBridgeDomainEntry.EntityData.YFilter = cContextMappingBridgeDomainEntry.YFilter
    cContextMappingBridgeDomainEntry.EntityData.YangName = "cContextMappingBridgeDomainEntry"
    cContextMappingBridgeDomainEntry.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingBridgeDomainEntry.EntityData.ParentYangName = "cContextMappingBridgeDomainTable"
    cContextMappingBridgeDomainEntry.EntityData.SegmentPath = "cContextMappingBridgeDomainEntry" + types.AddKeyToken(cContextMappingBridgeDomainEntry.CContextMappingVacmContextName, "cContextMappingVacmContextName")
    cContextMappingBridgeDomainEntry.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/cContextMappingBridgeDomainTable/" + cContextMappingBridgeDomainEntry.EntityData.SegmentPath
    cContextMappingBridgeDomainEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingBridgeDomainEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingBridgeDomainEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingBridgeDomainEntry.EntityData.Children = types.NewOrderedMap()
    cContextMappingBridgeDomainEntry.EntityData.Leafs = types.NewOrderedMap()
    cContextMappingBridgeDomainEntry.EntityData.Leafs.Append("cContextMappingVacmContextName", types.YLeaf{"CContextMappingVacmContextName", cContextMappingBridgeDomainEntry.CContextMappingVacmContextName})
    cContextMappingBridgeDomainEntry.EntityData.Leafs.Append("cContextMappingBridgeDomainIdentifier", types.YLeaf{"CContextMappingBridgeDomainIdentifier", cContextMappingBridgeDomainEntry.CContextMappingBridgeDomainIdentifier})
    cContextMappingBridgeDomainEntry.EntityData.Leafs.Append("cContextMappingBridgeDomainStorageType", types.YLeaf{"CContextMappingBridgeDomainStorageType", cContextMappingBridgeDomainEntry.CContextMappingBridgeDomainStorageType})
    cContextMappingBridgeDomainEntry.EntityData.Leafs.Append("cContextMappingBridgeDomainRowStatus", types.YLeaf{"CContextMappingBridgeDomainRowStatus", cContextMappingBridgeDomainEntry.CContextMappingBridgeDomainRowStatus})

    cContextMappingBridgeDomainEntry.EntityData.YListKeys = []string {"CContextMappingVacmContextName"}

    return &(cContextMappingBridgeDomainEntry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable
// This table contains information on mapping between
// cContextMappingVacmContextName and bridge instance.
// 
// Bridge instance is an instance of a physical or logical 
// bridge which has unique bridge-id.
// 
// If an entry is deleted from cContextMappingTable, the
// corresponding entry in this table will also get deleted.
// 
// If an entry needs to be created in this table, the
// corresponding entry must exist in cContextMappingTable.
type CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the  corresponding bridge instance.  To create a row in this table, a
    // manager must set cContextMappingBridgeInstRowStatus to either 
    // 'createAndGo' or 'createAndWait'.  To delete a row in this table, a manager
    // must set cContextMappingBridgeInstRowStatus to 'destroy'. The type is slice
    // of
    // CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable_CContextMappingBridgeInstanceEntry.
    CContextMappingBridgeInstanceEntry []*CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable_CContextMappingBridgeInstanceEntry
}

func (cContextMappingBridgeInstanceTable *CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable) GetEntityData() *types.CommonEntityData {
    cContextMappingBridgeInstanceTable.EntityData.YFilter = cContextMappingBridgeInstanceTable.YFilter
    cContextMappingBridgeInstanceTable.EntityData.YangName = "cContextMappingBridgeInstanceTable"
    cContextMappingBridgeInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingBridgeInstanceTable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cContextMappingBridgeInstanceTable.EntityData.SegmentPath = "cContextMappingBridgeInstanceTable"
    cContextMappingBridgeInstanceTable.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/" + cContextMappingBridgeInstanceTable.EntityData.SegmentPath
    cContextMappingBridgeInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingBridgeInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingBridgeInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingBridgeInstanceTable.EntityData.Children = types.NewOrderedMap()
    cContextMappingBridgeInstanceTable.EntityData.Children.Append("cContextMappingBridgeInstanceEntry", types.YChild{"CContextMappingBridgeInstanceEntry", nil})
    for i := range cContextMappingBridgeInstanceTable.CContextMappingBridgeInstanceEntry {
        cContextMappingBridgeInstanceTable.EntityData.Children.Append(types.GetSegmentPath(cContextMappingBridgeInstanceTable.CContextMappingBridgeInstanceEntry[i]), types.YChild{"CContextMappingBridgeInstanceEntry", cContextMappingBridgeInstanceTable.CContextMappingBridgeInstanceEntry[i]})
    }
    cContextMappingBridgeInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    cContextMappingBridgeInstanceTable.EntityData.YListKeys = []string {}

    return &(cContextMappingBridgeInstanceTable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable_CContextMappingBridgeInstanceEntry
// Information relating to a single mapping of
// cContextMappingVacmContextName to the 
// corresponding bridge instance.
// 
// To create a row in this table, a manager must set
// cContextMappingBridgeInstRowStatus to either 
// 'createAndGo' or 'createAndWait'.
// 
// To delete a row in this table, a manager must set
// cContextMappingBridgeInstRowStatus to 'destroy'.
type CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable_CContextMappingBridgeInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry_CContextMappingVacmContextName
    CContextMappingVacmContextName interface{}

    // The object identifies the name given to bridge instance to which the SNMP
    // context is mapped to.  Value of this object cannot be changed when the 
    // RowStatus object in the same row is 'active'.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this bridge instance.  When the value of this
    // object is a zero length string, it indicates that the SNMP context is
    // independent of any bridge instances. The type is string.
    CContextMappingBridgeInstName interface{}

    // The storage type for this conceptual row.  Value of this object cannot be
    // changed when the  RowStatus object in the same row is 'active'.  Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    CContextMappingBridgeInstStorageType interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    CContextMappingBridgeInstRowStatus interface{}
}

func (cContextMappingBridgeInstanceEntry *CISCOCONTEXTMAPPINGMIB_CContextMappingBridgeInstanceTable_CContextMappingBridgeInstanceEntry) GetEntityData() *types.CommonEntityData {
    cContextMappingBridgeInstanceEntry.EntityData.YFilter = cContextMappingBridgeInstanceEntry.YFilter
    cContextMappingBridgeInstanceEntry.EntityData.YangName = "cContextMappingBridgeInstanceEntry"
    cContextMappingBridgeInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingBridgeInstanceEntry.EntityData.ParentYangName = "cContextMappingBridgeInstanceTable"
    cContextMappingBridgeInstanceEntry.EntityData.SegmentPath = "cContextMappingBridgeInstanceEntry" + types.AddKeyToken(cContextMappingBridgeInstanceEntry.CContextMappingVacmContextName, "cContextMappingVacmContextName")
    cContextMappingBridgeInstanceEntry.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/cContextMappingBridgeInstanceTable/" + cContextMappingBridgeInstanceEntry.EntityData.SegmentPath
    cContextMappingBridgeInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingBridgeInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingBridgeInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingBridgeInstanceEntry.EntityData.Children = types.NewOrderedMap()
    cContextMappingBridgeInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    cContextMappingBridgeInstanceEntry.EntityData.Leafs.Append("cContextMappingVacmContextName", types.YLeaf{"CContextMappingVacmContextName", cContextMappingBridgeInstanceEntry.CContextMappingVacmContextName})
    cContextMappingBridgeInstanceEntry.EntityData.Leafs.Append("cContextMappingBridgeInstName", types.YLeaf{"CContextMappingBridgeInstName", cContextMappingBridgeInstanceEntry.CContextMappingBridgeInstName})
    cContextMappingBridgeInstanceEntry.EntityData.Leafs.Append("cContextMappingBridgeInstStorageType", types.YLeaf{"CContextMappingBridgeInstStorageType", cContextMappingBridgeInstanceEntry.CContextMappingBridgeInstStorageType})
    cContextMappingBridgeInstanceEntry.EntityData.Leafs.Append("cContextMappingBridgeInstRowStatus", types.YLeaf{"CContextMappingBridgeInstRowStatus", cContextMappingBridgeInstanceEntry.CContextMappingBridgeInstRowStatus})

    cContextMappingBridgeInstanceEntry.EntityData.YListKeys = []string {"CContextMappingVacmContextName"}

    return &(cContextMappingBridgeInstanceEntry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable
// This table contains information on which
// cContextMappingVacmContextName is mapped to
// which License Group.
// Group level licensing is used where each
// Technology Package is enabled via a License.
type CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of CContextMappingVacmContextName
    // to the corresponding License Group. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable_CContextMappingLicenseGroupEntry.
    CContextMappingLicenseGroupEntry []*CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable_CContextMappingLicenseGroupEntry
}

func (cContextMappingLicenseGroupTable *CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable) GetEntityData() *types.CommonEntityData {
    cContextMappingLicenseGroupTable.EntityData.YFilter = cContextMappingLicenseGroupTable.YFilter
    cContextMappingLicenseGroupTable.EntityData.YangName = "cContextMappingLicenseGroupTable"
    cContextMappingLicenseGroupTable.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingLicenseGroupTable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cContextMappingLicenseGroupTable.EntityData.SegmentPath = "cContextMappingLicenseGroupTable"
    cContextMappingLicenseGroupTable.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/" + cContextMappingLicenseGroupTable.EntityData.SegmentPath
    cContextMappingLicenseGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingLicenseGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingLicenseGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingLicenseGroupTable.EntityData.Children = types.NewOrderedMap()
    cContextMappingLicenseGroupTable.EntityData.Children.Append("cContextMappingLicenseGroupEntry", types.YChild{"CContextMappingLicenseGroupEntry", nil})
    for i := range cContextMappingLicenseGroupTable.CContextMappingLicenseGroupEntry {
        cContextMappingLicenseGroupTable.EntityData.Children.Append(types.GetSegmentPath(cContextMappingLicenseGroupTable.CContextMappingLicenseGroupEntry[i]), types.YChild{"CContextMappingLicenseGroupEntry", cContextMappingLicenseGroupTable.CContextMappingLicenseGroupEntry[i]})
    }
    cContextMappingLicenseGroupTable.EntityData.Leafs = types.NewOrderedMap()

    cContextMappingLicenseGroupTable.EntityData.YListKeys = []string {}

    return &(cContextMappingLicenseGroupTable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable_CContextMappingLicenseGroupEntry
// Information relating to a single mapping of
// CContextMappingVacmContextName to the
// corresponding License Group.
type CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable_CContextMappingLicenseGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_CContextMappingTable_CContextMappingEntry_CContextMappingVacmContextName
    CContextMappingVacmContextName interface{}

    // The value of an instance of this object identifies the name given to the
    // Group to which the SNMP context is mapped.  Feature sets from all groups
    // will be combined to form  universal image. User can configure multiple
    // groups as needed.  For example: In Next generation ISRs will use the
    // universal image package level licensing model for its licensing need. Each
    // group has the feature set needed for that specific technology. Feature sets
    // from different groups are combined to  form universal image and each
    // feature set for a group  can be enabled using a valid license key. There
    // will  be a base level ipbase package in which the router  boots with out
    // any license key.  The following are the different Technology Groups.
    // 1.crypto 2.data 3.ip 4.legacy 5.novpn-security 6.security 7.uc. The type is
    // string with length: 0..32.
    CContextMappingLicenseGroupName interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    CContextMappingLicenseGroupStorageType interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    CContextMappingLicenseGroupRowStatus interface{}
}

func (cContextMappingLicenseGroupEntry *CISCOCONTEXTMAPPINGMIB_CContextMappingLicenseGroupTable_CContextMappingLicenseGroupEntry) GetEntityData() *types.CommonEntityData {
    cContextMappingLicenseGroupEntry.EntityData.YFilter = cContextMappingLicenseGroupEntry.YFilter
    cContextMappingLicenseGroupEntry.EntityData.YangName = "cContextMappingLicenseGroupEntry"
    cContextMappingLicenseGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    cContextMappingLicenseGroupEntry.EntityData.ParentYangName = "cContextMappingLicenseGroupTable"
    cContextMappingLicenseGroupEntry.EntityData.SegmentPath = "cContextMappingLicenseGroupEntry" + types.AddKeyToken(cContextMappingLicenseGroupEntry.CContextMappingVacmContextName, "cContextMappingVacmContextName")
    cContextMappingLicenseGroupEntry.EntityData.AbsolutePath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB/cContextMappingLicenseGroupTable/" + cContextMappingLicenseGroupEntry.EntityData.SegmentPath
    cContextMappingLicenseGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cContextMappingLicenseGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cContextMappingLicenseGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cContextMappingLicenseGroupEntry.EntityData.Children = types.NewOrderedMap()
    cContextMappingLicenseGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    cContextMappingLicenseGroupEntry.EntityData.Leafs.Append("cContextMappingVacmContextName", types.YLeaf{"CContextMappingVacmContextName", cContextMappingLicenseGroupEntry.CContextMappingVacmContextName})
    cContextMappingLicenseGroupEntry.EntityData.Leafs.Append("cContextMappingLicenseGroupName", types.YLeaf{"CContextMappingLicenseGroupName", cContextMappingLicenseGroupEntry.CContextMappingLicenseGroupName})
    cContextMappingLicenseGroupEntry.EntityData.Leafs.Append("cContextMappingLicenseGroupStorageType", types.YLeaf{"CContextMappingLicenseGroupStorageType", cContextMappingLicenseGroupEntry.CContextMappingLicenseGroupStorageType})
    cContextMappingLicenseGroupEntry.EntityData.Leafs.Append("cContextMappingLicenseGroupRowStatus", types.YLeaf{"CContextMappingLicenseGroupRowStatus", cContextMappingLicenseGroupEntry.CContextMappingLicenseGroupRowStatus})

    cContextMappingLicenseGroupEntry.EntityData.YListKeys = []string {"CContextMappingVacmContextName"}

    return &(cContextMappingLicenseGroupEntry.EntityData)
}

