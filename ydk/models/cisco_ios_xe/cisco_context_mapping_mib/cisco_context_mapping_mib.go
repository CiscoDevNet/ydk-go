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
    Ccontextmappingtable CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable

    // This table contains information on which cContextMappingVacmContextName is
    // mapped to which bridge domain.  A Bridge Domain is one of the means by
    // which it is possible  to define an Ethernet broadcast domain on a bridging
    // device.  A network can have multiple broadcast domains configured. This
    // table helps the network management personnel to find  out the  details of
    // various broadcast domains configured  in the network.  An entry need to
    // exist in cContextMappingTable, to create  an entry in this table.
    Ccontextmappingbridgedomaintable CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable

    // This table contains information on mapping between
    // cContextMappingVacmContextName and bridge instance.  Bridge instance is an
    // instance of a physical or logical  bridge which has unique bridge-id.  If
    // an entry is deleted from cContextMappingTable, the corresponding entry in
    // this table will also get deleted.  If an entry needs to be created in this
    // table, the corresponding entry must exist in cContextMappingTable.
    Ccontextmappingbridgeinstancetable CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable

    // This table contains information on which cContextMappingVacmContextName is
    // mapped to which License Group. Group level licensing is used where each
    // Technology Package is enabled via a License.
    Ccontextmappinglicensegrouptable CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetEntityData() *types.CommonEntityData {
    cISCOCONTEXTMAPPINGMIB.EntityData.YFilter = cISCOCONTEXTMAPPINGMIB.YFilter
    cISCOCONTEXTMAPPINGMIB.EntityData.YangName = "CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCONTEXTMAPPINGMIB.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.SegmentPath = "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB"
    cISCOCONTEXTMAPPINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCONTEXTMAPPINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCONTEXTMAPPINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCONTEXTMAPPINGMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOCONTEXTMAPPINGMIB.EntityData.Children["cContextMappingTable"] = types.YChild{"Ccontextmappingtable", &cISCOCONTEXTMAPPINGMIB.Ccontextmappingtable}
    cISCOCONTEXTMAPPINGMIB.EntityData.Children["cContextMappingBridgeDomainTable"] = types.YChild{"Ccontextmappingbridgedomaintable", &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgedomaintable}
    cISCOCONTEXTMAPPINGMIB.EntityData.Children["cContextMappingBridgeInstanceTable"] = types.YChild{"Ccontextmappingbridgeinstancetable", &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgeinstancetable}
    cISCOCONTEXTMAPPINGMIB.EntityData.Children["cContextMappingLicenseGroupTable"] = types.YChild{"Ccontextmappinglicensegrouptable", &cISCOCONTEXTMAPPINGMIB.Ccontextmappinglicensegrouptable}
    cISCOCONTEXTMAPPINGMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOCONTEXTMAPPINGMIB.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable
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
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the corresponding VRF, the corresponding topology, and the corresponding
    // routing protocol instance. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry.
    Ccontextmappingentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetEntityData() *types.CommonEntityData {
    ccontextmappingtable.EntityData.YFilter = ccontextmappingtable.YFilter
    ccontextmappingtable.EntityData.YangName = "cContextMappingTable"
    ccontextmappingtable.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingtable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    ccontextmappingtable.EntityData.SegmentPath = "cContextMappingTable"
    ccontextmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingtable.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingtable.EntityData.Children["cContextMappingEntry"] = types.YChild{"Ccontextmappingentry", nil}
    for i := range ccontextmappingtable.Ccontextmappingentry {
        ccontextmappingtable.EntityData.Children[types.GetSegmentPath(&ccontextmappingtable.Ccontextmappingentry[i])] = types.YChild{"Ccontextmappingentry", &ccontextmappingtable.Ccontextmappingentry[i]}
    }
    ccontextmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccontextmappingtable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry
// Information relating to a single mapping of
// cContextMappingVacmContextName to the corresponding VRF,
// the corresponding topology, and the corresponding routing
// protocol instance.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The vacmContextName given to the SNMP context. 
    // This is a human readable name identifying a particular SNMP VACM context at
    // a particular SNMP entity. The empty contextName (zero length) represents
    // the default context. The type is string with length: 0..32.
    Ccontextmappingvacmcontextname interface{}

    // The value of an instance of this object identifies the name given to the
    // VRF to which the SNMP context is mapped to.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this VRF.  When the value of this object is
    // the zero length string it indicates that the SNMP context is independent of
    // any VRF. The type is string with length: 0..32.
    Ccontextmappingvrfname interface{}

    // The value of an instance of this object identifies the name given to the
    // topology to which the SNMP context is mapped to.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this topology.  When the value of this object
    // is the zero length string it indicates that the SNMP context is independent
    // of any topology. The type is string with length: 0..32.
    Ccontextmappingtopologyname interface{}

    // The value of an instance of this object identifies the name given to the
    // protocol instance to which the SNMP context is mapped to.  This is
    // typically a human-readable string. This is the same ASCII string used in
    // the router's console interface to refer to this protocol instance.  When
    // the value of this object is the zero length string it indicates that the
    // SNMP context is independent of any protocol instance. The type is string
    // with length: 0..32.
    Ccontextmappingprotoinstname interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Ccontextmappingstoragetype interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    Ccontextmappingrowstatus interface{}
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetEntityData() *types.CommonEntityData {
    ccontextmappingentry.EntityData.YFilter = ccontextmappingentry.YFilter
    ccontextmappingentry.EntityData.YangName = "cContextMappingEntry"
    ccontextmappingentry.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingentry.EntityData.ParentYangName = "cContextMappingTable"
    ccontextmappingentry.EntityData.SegmentPath = "cContextMappingEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingentry.Ccontextmappingvacmcontextname) + "']"
    ccontextmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingentry.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccontextmappingentry.EntityData.Leafs["cContextMappingVacmContextName"] = types.YLeaf{"Ccontextmappingvacmcontextname", ccontextmappingentry.Ccontextmappingvacmcontextname}
    ccontextmappingentry.EntityData.Leafs["cContextMappingVrfName"] = types.YLeaf{"Ccontextmappingvrfname", ccontextmappingentry.Ccontextmappingvrfname}
    ccontextmappingentry.EntityData.Leafs["cContextMappingTopologyName"] = types.YLeaf{"Ccontextmappingtopologyname", ccontextmappingentry.Ccontextmappingtopologyname}
    ccontextmappingentry.EntityData.Leafs["cContextMappingProtoInstName"] = types.YLeaf{"Ccontextmappingprotoinstname", ccontextmappingentry.Ccontextmappingprotoinstname}
    ccontextmappingentry.EntityData.Leafs["cContextMappingStorageType"] = types.YLeaf{"Ccontextmappingstoragetype", ccontextmappingentry.Ccontextmappingstoragetype}
    ccontextmappingentry.EntityData.Leafs["cContextMappingRowStatus"] = types.YLeaf{"Ccontextmappingrowstatus", ccontextmappingentry.Ccontextmappingrowstatus}
    return &(ccontextmappingentry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable
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
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the  corresponding bridge domain.  To create a row in this table, a
    // manager must set cContextMappingBridgeDomainRowStatus to either 
    // 'createAndGo' or 'createAndWait'.  To delete a row in this table, a manager
    // must set cContextMappingBridgeDomainRowStatus to 'destroy'. The type is
    // slice of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry.
    Ccontextmappingbridgedomainentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetEntityData() *types.CommonEntityData {
    ccontextmappingbridgedomaintable.EntityData.YFilter = ccontextmappingbridgedomaintable.YFilter
    ccontextmappingbridgedomaintable.EntityData.YangName = "cContextMappingBridgeDomainTable"
    ccontextmappingbridgedomaintable.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingbridgedomaintable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    ccontextmappingbridgedomaintable.EntityData.SegmentPath = "cContextMappingBridgeDomainTable"
    ccontextmappingbridgedomaintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingbridgedomaintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingbridgedomaintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingbridgedomaintable.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingbridgedomaintable.EntityData.Children["cContextMappingBridgeDomainEntry"] = types.YChild{"Ccontextmappingbridgedomainentry", nil}
    for i := range ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry {
        ccontextmappingbridgedomaintable.EntityData.Children[types.GetSegmentPath(&ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry[i])] = types.YChild{"Ccontextmappingbridgedomainentry", &ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry[i]}
    }
    ccontextmappingbridgedomaintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccontextmappingbridgedomaintable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry
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
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry_Ccontextmappingvacmcontextname
    Ccontextmappingvacmcontextname interface{}

    // The value of an instance of this object identifies the bridge domain to
    // which the SNMP context is  mapped to. The type is interface{} with range:
    // 1..65535.
    Ccontextmappingbridgedomainidentifier interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Ccontextmappingbridgedomainstoragetype interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    Ccontextmappingbridgedomainrowstatus interface{}
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetEntityData() *types.CommonEntityData {
    ccontextmappingbridgedomainentry.EntityData.YFilter = ccontextmappingbridgedomainentry.YFilter
    ccontextmappingbridgedomainentry.EntityData.YangName = "cContextMappingBridgeDomainEntry"
    ccontextmappingbridgedomainentry.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingbridgedomainentry.EntityData.ParentYangName = "cContextMappingBridgeDomainTable"
    ccontextmappingbridgedomainentry.EntityData.SegmentPath = "cContextMappingBridgeDomainEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingbridgedomainentry.Ccontextmappingvacmcontextname) + "']"
    ccontextmappingbridgedomainentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingbridgedomainentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingbridgedomainentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingbridgedomainentry.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingbridgedomainentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccontextmappingbridgedomainentry.EntityData.Leafs["cContextMappingVacmContextName"] = types.YLeaf{"Ccontextmappingvacmcontextname", ccontextmappingbridgedomainentry.Ccontextmappingvacmcontextname}
    ccontextmappingbridgedomainentry.EntityData.Leafs["cContextMappingBridgeDomainIdentifier"] = types.YLeaf{"Ccontextmappingbridgedomainidentifier", ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainidentifier}
    ccontextmappingbridgedomainentry.EntityData.Leafs["cContextMappingBridgeDomainStorageType"] = types.YLeaf{"Ccontextmappingbridgedomainstoragetype", ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainstoragetype}
    ccontextmappingbridgedomainentry.EntityData.Leafs["cContextMappingBridgeDomainRowStatus"] = types.YLeaf{"Ccontextmappingbridgedomainrowstatus", ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainrowstatus}
    return &(ccontextmappingbridgedomainentry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable
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
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the  corresponding bridge instance.  To create a row in this table, a
    // manager must set cContextMappingBridgeInstRowStatus to either 
    // 'createAndGo' or 'createAndWait'.  To delete a row in this table, a manager
    // must set cContextMappingBridgeInstRowStatus to 'destroy'. The type is slice
    // of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry.
    Ccontextmappingbridgeinstanceentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetEntityData() *types.CommonEntityData {
    ccontextmappingbridgeinstancetable.EntityData.YFilter = ccontextmappingbridgeinstancetable.YFilter
    ccontextmappingbridgeinstancetable.EntityData.YangName = "cContextMappingBridgeInstanceTable"
    ccontextmappingbridgeinstancetable.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingbridgeinstancetable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    ccontextmappingbridgeinstancetable.EntityData.SegmentPath = "cContextMappingBridgeInstanceTable"
    ccontextmappingbridgeinstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingbridgeinstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingbridgeinstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingbridgeinstancetable.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingbridgeinstancetable.EntityData.Children["cContextMappingBridgeInstanceEntry"] = types.YChild{"Ccontextmappingbridgeinstanceentry", nil}
    for i := range ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry {
        ccontextmappingbridgeinstancetable.EntityData.Children[types.GetSegmentPath(&ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry[i])] = types.YChild{"Ccontextmappingbridgeinstanceentry", &ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry[i]}
    }
    ccontextmappingbridgeinstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccontextmappingbridgeinstancetable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry
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
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry_Ccontextmappingvacmcontextname
    Ccontextmappingvacmcontextname interface{}

    // The object identifies the name given to bridge instance to which the SNMP
    // context is mapped to.  Value of this object cannot be changed when the 
    // RowStatus object in the same row is 'active'.  This is typically a
    // human-readable string. This is the same ASCII string used in the router's
    // console interface to refer to this bridge instance.  When the value of this
    // object is a zero length string, it indicates that the SNMP context is
    // independent of any bridge instances. The type is string.
    Ccontextmappingbridgeinstname interface{}

    // The storage type for this conceptual row.  Value of this object cannot be
    // changed when the  RowStatus object in the same row is 'active'.  Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    Ccontextmappingbridgeinststoragetype interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    Ccontextmappingbridgeinstrowstatus interface{}
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetEntityData() *types.CommonEntityData {
    ccontextmappingbridgeinstanceentry.EntityData.YFilter = ccontextmappingbridgeinstanceentry.YFilter
    ccontextmappingbridgeinstanceentry.EntityData.YangName = "cContextMappingBridgeInstanceEntry"
    ccontextmappingbridgeinstanceentry.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappingbridgeinstanceentry.EntityData.ParentYangName = "cContextMappingBridgeInstanceTable"
    ccontextmappingbridgeinstanceentry.EntityData.SegmentPath = "cContextMappingBridgeInstanceEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingbridgeinstanceentry.Ccontextmappingvacmcontextname) + "']"
    ccontextmappingbridgeinstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappingbridgeinstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappingbridgeinstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappingbridgeinstanceentry.EntityData.Children = make(map[string]types.YChild)
    ccontextmappingbridgeinstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccontextmappingbridgeinstanceentry.EntityData.Leafs["cContextMappingVacmContextName"] = types.YLeaf{"Ccontextmappingvacmcontextname", ccontextmappingbridgeinstanceentry.Ccontextmappingvacmcontextname}
    ccontextmappingbridgeinstanceentry.EntityData.Leafs["cContextMappingBridgeInstName"] = types.YLeaf{"Ccontextmappingbridgeinstname", ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinstname}
    ccontextmappingbridgeinstanceentry.EntityData.Leafs["cContextMappingBridgeInstStorageType"] = types.YLeaf{"Ccontextmappingbridgeinststoragetype", ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinststoragetype}
    ccontextmappingbridgeinstanceentry.EntityData.Leafs["cContextMappingBridgeInstRowStatus"] = types.YLeaf{"Ccontextmappingbridgeinstrowstatus", ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinstrowstatus}
    return &(ccontextmappingbridgeinstanceentry.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable
// This table contains information on which
// cContextMappingVacmContextName is mapped to
// which License Group.
// Group level licensing is used where each
// Technology Package is enabled via a License.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information relating to a single mapping of CContextMappingVacmContextName
    // to the corresponding License Group. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry.
    Ccontextmappinglicensegroupentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetEntityData() *types.CommonEntityData {
    ccontextmappinglicensegrouptable.EntityData.YFilter = ccontextmappinglicensegrouptable.YFilter
    ccontextmappinglicensegrouptable.EntityData.YangName = "cContextMappingLicenseGroupTable"
    ccontextmappinglicensegrouptable.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappinglicensegrouptable.EntityData.ParentYangName = "CISCO-CONTEXT-MAPPING-MIB"
    ccontextmappinglicensegrouptable.EntityData.SegmentPath = "cContextMappingLicenseGroupTable"
    ccontextmappinglicensegrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappinglicensegrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappinglicensegrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappinglicensegrouptable.EntityData.Children = make(map[string]types.YChild)
    ccontextmappinglicensegrouptable.EntityData.Children["cContextMappingLicenseGroupEntry"] = types.YChild{"Ccontextmappinglicensegroupentry", nil}
    for i := range ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry {
        ccontextmappinglicensegrouptable.EntityData.Children[types.GetSegmentPath(&ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry[i])] = types.YChild{"Ccontextmappinglicensegroupentry", &ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry[i]}
    }
    ccontextmappinglicensegrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccontextmappinglicensegrouptable.EntityData)
}

// CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry
// Information relating to a single mapping of
// CContextMappingVacmContextName to the
// corresponding License Group.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_context_mapping_mib.CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry_Ccontextmappingvacmcontextname
    Ccontextmappingvacmcontextname interface{}

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
    Ccontextmappinglicensegroupname interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Ccontextmappinglicensegroupstoragetype interface{}

    // This object facilitates the creation, modification, or deletion of a
    // conceptual row in this table. The type is RowStatus.
    Ccontextmappinglicensegrouprowstatus interface{}
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetEntityData() *types.CommonEntityData {
    ccontextmappinglicensegroupentry.EntityData.YFilter = ccontextmappinglicensegroupentry.YFilter
    ccontextmappinglicensegroupentry.EntityData.YangName = "cContextMappingLicenseGroupEntry"
    ccontextmappinglicensegroupentry.EntityData.BundleName = "cisco_ios_xe"
    ccontextmappinglicensegroupentry.EntityData.ParentYangName = "cContextMappingLicenseGroupTable"
    ccontextmappinglicensegroupentry.EntityData.SegmentPath = "cContextMappingLicenseGroupEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappinglicensegroupentry.Ccontextmappingvacmcontextname) + "']"
    ccontextmappinglicensegroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccontextmappinglicensegroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccontextmappinglicensegroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccontextmappinglicensegroupentry.EntityData.Children = make(map[string]types.YChild)
    ccontextmappinglicensegroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccontextmappinglicensegroupentry.EntityData.Leafs["cContextMappingVacmContextName"] = types.YLeaf{"Ccontextmappingvacmcontextname", ccontextmappinglicensegroupentry.Ccontextmappingvacmcontextname}
    ccontextmappinglicensegroupentry.EntityData.Leafs["cContextMappingLicenseGroupName"] = types.YLeaf{"Ccontextmappinglicensegroupname", ccontextmappinglicensegroupentry.Ccontextmappinglicensegroupname}
    ccontextmappinglicensegroupentry.EntityData.Leafs["cContextMappingLicenseGroupStorageType"] = types.YLeaf{"Ccontextmappinglicensegroupstoragetype", ccontextmappinglicensegroupentry.Ccontextmappinglicensegroupstoragetype}
    ccontextmappinglicensegroupentry.EntityData.Leafs["cContextMappingLicenseGroupRowStatus"] = types.YLeaf{"Ccontextmappinglicensegrouprowstatus", ccontextmappinglicensegroupentry.Ccontextmappinglicensegrouprowstatus}
    return &(ccontextmappinglicensegroupentry.EntityData)
}

