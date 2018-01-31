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
    parent types.Entity
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

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetFilter() yfilter.YFilter { return cISCOCONTEXTMAPPINGMIB.YFilter }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) SetFilter(yf yfilter.YFilter) { cISCOCONTEXTMAPPINGMIB.YFilter = yf }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetGoName(yname string) string {
    if yname == "cContextMappingTable" { return "Ccontextmappingtable" }
    if yname == "cContextMappingBridgeDomainTable" { return "Ccontextmappingbridgedomaintable" }
    if yname == "cContextMappingBridgeInstanceTable" { return "Ccontextmappingbridgeinstancetable" }
    if yname == "cContextMappingLicenseGroupTable" { return "Ccontextmappinglicensegrouptable" }
    return ""
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetSegmentPath() string {
    return "CISCO-CONTEXT-MAPPING-MIB:CISCO-CONTEXT-MAPPING-MIB"
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cContextMappingTable" {
        return &cISCOCONTEXTMAPPINGMIB.Ccontextmappingtable
    }
    if childYangName == "cContextMappingBridgeDomainTable" {
        return &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgedomaintable
    }
    if childYangName == "cContextMappingBridgeInstanceTable" {
        return &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgeinstancetable
    }
    if childYangName == "cContextMappingLicenseGroupTable" {
        return &cISCOCONTEXTMAPPINGMIB.Ccontextmappinglicensegrouptable
    }
    return nil
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cContextMappingTable"] = &cISCOCONTEXTMAPPINGMIB.Ccontextmappingtable
    children["cContextMappingBridgeDomainTable"] = &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgedomaintable
    children["cContextMappingBridgeInstanceTable"] = &cISCOCONTEXTMAPPINGMIB.Ccontextmappingbridgeinstancetable
    children["cContextMappingLicenseGroupTable"] = &cISCOCONTEXTMAPPINGMIB.Ccontextmappinglicensegrouptable
    return children
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) SetParent(parent types.Entity) { cISCOCONTEXTMAPPINGMIB.parent = parent }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetParent() types.Entity { return cISCOCONTEXTMAPPINGMIB.parent }

func (cISCOCONTEXTMAPPINGMIB *CISCOCONTEXTMAPPINGMIB) GetParentYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information relating to a single mapping of cContextMappingVacmContextName
    // to the corresponding VRF, the corresponding topology, and the corresponding
    // routing protocol instance. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry.
    Ccontextmappingentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetFilter() yfilter.YFilter { return ccontextmappingtable.YFilter }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) SetFilter(yf yfilter.YFilter) { ccontextmappingtable.YFilter = yf }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetGoName(yname string) string {
    if yname == "cContextMappingEntry" { return "Ccontextmappingentry" }
    return ""
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetSegmentPath() string {
    return "cContextMappingTable"
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cContextMappingEntry" {
        for _, c := range ccontextmappingtable.Ccontextmappingentry {
            if ccontextmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry{}
        ccontextmappingtable.Ccontextmappingentry = append(ccontextmappingtable.Ccontextmappingentry, child)
        return &ccontextmappingtable.Ccontextmappingentry[len(ccontextmappingtable.Ccontextmappingentry)-1]
    }
    return nil
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccontextmappingtable.Ccontextmappingentry {
        children[ccontextmappingtable.Ccontextmappingentry[i].GetSegmentPath()] = &ccontextmappingtable.Ccontextmappingentry[i]
    }
    return children
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetYangName() string { return "cContextMappingTable" }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) SetParent(parent types.Entity) { ccontextmappingtable.parent = parent }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetParent() types.Entity { return ccontextmappingtable.parent }

func (ccontextmappingtable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable) GetParentYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

// CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry
// Information relating to a single mapping of
// cContextMappingVacmContextName to the corresponding VRF,
// the corresponding topology, and the corresponding routing
// protocol instance.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry struct {
    parent types.Entity
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

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetFilter() yfilter.YFilter { return ccontextmappingentry.YFilter }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) SetFilter(yf yfilter.YFilter) { ccontextmappingentry.YFilter = yf }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetGoName(yname string) string {
    if yname == "cContextMappingVacmContextName" { return "Ccontextmappingvacmcontextname" }
    if yname == "cContextMappingVrfName" { return "Ccontextmappingvrfname" }
    if yname == "cContextMappingTopologyName" { return "Ccontextmappingtopologyname" }
    if yname == "cContextMappingProtoInstName" { return "Ccontextmappingprotoinstname" }
    if yname == "cContextMappingStorageType" { return "Ccontextmappingstoragetype" }
    if yname == "cContextMappingRowStatus" { return "Ccontextmappingrowstatus" }
    return ""
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetSegmentPath() string {
    return "cContextMappingEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingentry.Ccontextmappingvacmcontextname) + "']"
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cContextMappingVacmContextName"] = ccontextmappingentry.Ccontextmappingvacmcontextname
    leafs["cContextMappingVrfName"] = ccontextmappingentry.Ccontextmappingvrfname
    leafs["cContextMappingTopologyName"] = ccontextmappingentry.Ccontextmappingtopologyname
    leafs["cContextMappingProtoInstName"] = ccontextmappingentry.Ccontextmappingprotoinstname
    leafs["cContextMappingStorageType"] = ccontextmappingentry.Ccontextmappingstoragetype
    leafs["cContextMappingRowStatus"] = ccontextmappingentry.Ccontextmappingrowstatus
    return leafs
}

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetYangName() string { return "cContextMappingEntry" }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) SetParent(parent types.Entity) { ccontextmappingentry.parent = parent }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetParent() types.Entity { return ccontextmappingentry.parent }

func (ccontextmappingentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingtable_Ccontextmappingentry) GetParentYangName() string { return "cContextMappingTable" }

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
    parent types.Entity
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

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetFilter() yfilter.YFilter { return ccontextmappingbridgedomaintable.YFilter }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) SetFilter(yf yfilter.YFilter) { ccontextmappingbridgedomaintable.YFilter = yf }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetGoName(yname string) string {
    if yname == "cContextMappingBridgeDomainEntry" { return "Ccontextmappingbridgedomainentry" }
    return ""
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetSegmentPath() string {
    return "cContextMappingBridgeDomainTable"
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cContextMappingBridgeDomainEntry" {
        for _, c := range ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry {
            if ccontextmappingbridgedomaintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry{}
        ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry = append(ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry, child)
        return &ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry[len(ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry)-1]
    }
    return nil
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry {
        children[ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry[i].GetSegmentPath()] = &ccontextmappingbridgedomaintable.Ccontextmappingbridgedomainentry[i]
    }
    return children
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetYangName() string { return "cContextMappingBridgeDomainTable" }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) SetParent(parent types.Entity) { ccontextmappingbridgedomaintable.parent = parent }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetParent() types.Entity { return ccontextmappingbridgedomaintable.parent }

func (ccontextmappingbridgedomaintable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable) GetParentYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

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
    parent types.Entity
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

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetFilter() yfilter.YFilter { return ccontextmappingbridgedomainentry.YFilter }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) SetFilter(yf yfilter.YFilter) { ccontextmappingbridgedomainentry.YFilter = yf }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetGoName(yname string) string {
    if yname == "cContextMappingVacmContextName" { return "Ccontextmappingvacmcontextname" }
    if yname == "cContextMappingBridgeDomainIdentifier" { return "Ccontextmappingbridgedomainidentifier" }
    if yname == "cContextMappingBridgeDomainStorageType" { return "Ccontextmappingbridgedomainstoragetype" }
    if yname == "cContextMappingBridgeDomainRowStatus" { return "Ccontextmappingbridgedomainrowstatus" }
    return ""
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetSegmentPath() string {
    return "cContextMappingBridgeDomainEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingbridgedomainentry.Ccontextmappingvacmcontextname) + "']"
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cContextMappingVacmContextName"] = ccontextmappingbridgedomainentry.Ccontextmappingvacmcontextname
    leafs["cContextMappingBridgeDomainIdentifier"] = ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainidentifier
    leafs["cContextMappingBridgeDomainStorageType"] = ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainstoragetype
    leafs["cContextMappingBridgeDomainRowStatus"] = ccontextmappingbridgedomainentry.Ccontextmappingbridgedomainrowstatus
    return leafs
}

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetYangName() string { return "cContextMappingBridgeDomainEntry" }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) SetParent(parent types.Entity) { ccontextmappingbridgedomainentry.parent = parent }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetParent() types.Entity { return ccontextmappingbridgedomainentry.parent }

func (ccontextmappingbridgedomainentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgedomaintable_Ccontextmappingbridgedomainentry) GetParentYangName() string { return "cContextMappingBridgeDomainTable" }

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
    parent types.Entity
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

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetFilter() yfilter.YFilter { return ccontextmappingbridgeinstancetable.YFilter }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) SetFilter(yf yfilter.YFilter) { ccontextmappingbridgeinstancetable.YFilter = yf }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetGoName(yname string) string {
    if yname == "cContextMappingBridgeInstanceEntry" { return "Ccontextmappingbridgeinstanceentry" }
    return ""
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetSegmentPath() string {
    return "cContextMappingBridgeInstanceTable"
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cContextMappingBridgeInstanceEntry" {
        for _, c := range ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry {
            if ccontextmappingbridgeinstancetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry{}
        ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry = append(ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry, child)
        return &ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry[len(ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry)-1]
    }
    return nil
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry {
        children[ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry[i].GetSegmentPath()] = &ccontextmappingbridgeinstancetable.Ccontextmappingbridgeinstanceentry[i]
    }
    return children
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetYangName() string { return "cContextMappingBridgeInstanceTable" }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) SetParent(parent types.Entity) { ccontextmappingbridgeinstancetable.parent = parent }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetParent() types.Entity { return ccontextmappingbridgeinstancetable.parent }

func (ccontextmappingbridgeinstancetable *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable) GetParentYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

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
    parent types.Entity
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

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetFilter() yfilter.YFilter { return ccontextmappingbridgeinstanceentry.YFilter }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) SetFilter(yf yfilter.YFilter) { ccontextmappingbridgeinstanceentry.YFilter = yf }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetGoName(yname string) string {
    if yname == "cContextMappingVacmContextName" { return "Ccontextmappingvacmcontextname" }
    if yname == "cContextMappingBridgeInstName" { return "Ccontextmappingbridgeinstname" }
    if yname == "cContextMappingBridgeInstStorageType" { return "Ccontextmappingbridgeinststoragetype" }
    if yname == "cContextMappingBridgeInstRowStatus" { return "Ccontextmappingbridgeinstrowstatus" }
    return ""
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetSegmentPath() string {
    return "cContextMappingBridgeInstanceEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappingbridgeinstanceentry.Ccontextmappingvacmcontextname) + "']"
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cContextMappingVacmContextName"] = ccontextmappingbridgeinstanceentry.Ccontextmappingvacmcontextname
    leafs["cContextMappingBridgeInstName"] = ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinstname
    leafs["cContextMappingBridgeInstStorageType"] = ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinststoragetype
    leafs["cContextMappingBridgeInstRowStatus"] = ccontextmappingbridgeinstanceentry.Ccontextmappingbridgeinstrowstatus
    return leafs
}

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetYangName() string { return "cContextMappingBridgeInstanceEntry" }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) SetParent(parent types.Entity) { ccontextmappingbridgeinstanceentry.parent = parent }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetParent() types.Entity { return ccontextmappingbridgeinstanceentry.parent }

func (ccontextmappingbridgeinstanceentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappingbridgeinstancetable_Ccontextmappingbridgeinstanceentry) GetParentYangName() string { return "cContextMappingBridgeInstanceTable" }

// CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable
// This table contains information on which
// cContextMappingVacmContextName is mapped to
// which License Group.
// Group level licensing is used where each
// Technology Package is enabled via a License.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information relating to a single mapping of CContextMappingVacmContextName
    // to the corresponding License Group. The type is slice of
    // CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry.
    Ccontextmappinglicensegroupentry []CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetFilter() yfilter.YFilter { return ccontextmappinglicensegrouptable.YFilter }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) SetFilter(yf yfilter.YFilter) { ccontextmappinglicensegrouptable.YFilter = yf }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetGoName(yname string) string {
    if yname == "cContextMappingLicenseGroupEntry" { return "Ccontextmappinglicensegroupentry" }
    return ""
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetSegmentPath() string {
    return "cContextMappingLicenseGroupTable"
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cContextMappingLicenseGroupEntry" {
        for _, c := range ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry {
            if ccontextmappinglicensegrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry{}
        ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry = append(ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry, child)
        return &ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry[len(ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry)-1]
    }
    return nil
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry {
        children[ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry[i].GetSegmentPath()] = &ccontextmappinglicensegrouptable.Ccontextmappinglicensegroupentry[i]
    }
    return children
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetYangName() string { return "cContextMappingLicenseGroupTable" }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) SetParent(parent types.Entity) { ccontextmappinglicensegrouptable.parent = parent }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetParent() types.Entity { return ccontextmappinglicensegrouptable.parent }

func (ccontextmappinglicensegrouptable *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable) GetParentYangName() string { return "CISCO-CONTEXT-MAPPING-MIB" }

// CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry
// Information relating to a single mapping of
// CContextMappingVacmContextName to the
// corresponding License Group.
type CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry struct {
    parent types.Entity
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

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetFilter() yfilter.YFilter { return ccontextmappinglicensegroupentry.YFilter }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) SetFilter(yf yfilter.YFilter) { ccontextmappinglicensegroupentry.YFilter = yf }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetGoName(yname string) string {
    if yname == "cContextMappingVacmContextName" { return "Ccontextmappingvacmcontextname" }
    if yname == "cContextMappingLicenseGroupName" { return "Ccontextmappinglicensegroupname" }
    if yname == "cContextMappingLicenseGroupStorageType" { return "Ccontextmappinglicensegroupstoragetype" }
    if yname == "cContextMappingLicenseGroupRowStatus" { return "Ccontextmappinglicensegrouprowstatus" }
    return ""
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetSegmentPath() string {
    return "cContextMappingLicenseGroupEntry" + "[cContextMappingVacmContextName='" + fmt.Sprintf("%v", ccontextmappinglicensegroupentry.Ccontextmappingvacmcontextname) + "']"
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cContextMappingVacmContextName"] = ccontextmappinglicensegroupentry.Ccontextmappingvacmcontextname
    leafs["cContextMappingLicenseGroupName"] = ccontextmappinglicensegroupentry.Ccontextmappinglicensegroupname
    leafs["cContextMappingLicenseGroupStorageType"] = ccontextmappinglicensegroupentry.Ccontextmappinglicensegroupstoragetype
    leafs["cContextMappingLicenseGroupRowStatus"] = ccontextmappinglicensegroupentry.Ccontextmappinglicensegrouprowstatus
    return leafs
}

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetYangName() string { return "cContextMappingLicenseGroupEntry" }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) SetParent(parent types.Entity) { ccontextmappinglicensegroupentry.parent = parent }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetParent() types.Entity { return ccontextmappinglicensegroupentry.parent }

func (ccontextmappinglicensegroupentry *CISCOCONTEXTMAPPINGMIB_Ccontextmappinglicensegrouptable_Ccontextmappinglicensegroupentry) GetParentYangName() string { return "cContextMappingLicenseGroupTable" }

