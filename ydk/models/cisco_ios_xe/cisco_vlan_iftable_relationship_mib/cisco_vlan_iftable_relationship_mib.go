// Cisco VLAN ifTable Relationship MIB lists VLAN-id and ifIndex
// information for routed VLAN interfaces.  
// 
// A routed VLAN interface is the router interface or sub-interface 
// to which the router's IP address on the VLAN is attached.  
// For example, an ISL, SDE, or 802.1Q encapsulated
// subinterface, or Switched Virtual Interface (SVI).
package cisco_vlan_iftable_relationship_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_vlan_iftable_relationship_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB}", reflect.TypeOf(CISCOVLANIFTABLERELATIONSHIPMIB{}))
    ydk.RegisterEntity("CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB:CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB", reflect.TypeOf(CISCOVLANIFTABLERELATIONSHIPMIB{}))
}

// CISCOVLANIFTABLERELATIONSHIPMIB
type CISCOVLANIFTABLERELATIONSHIPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The cviVlanInterfaceIndexTable provides a way to translate a VLAN-id in to
    // an ifIndex, so that  the routed VLAN interface's routing configuration  can
    // be obtained from interface entry in ipRouteTable.  Note that some routers
    // can have interfaces to multiple VLAN management domains, and therefore can
    // have multiple  routed VLAN interfaces which connect to different VLANs 
    // having the same VLAN-id.  Thus, it is possible to have  multiple rows in
    // this table for the same VLAN-id.  The cviVlanInterfaceIndexTable also
    // provides a way to find the VLAN-id from an ifTable VLAN's ifIndex.
    Cvivlaninterfaceindextable CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetFilter() yfilter.YFilter { return cISCOVLANIFTABLERELATIONSHIPMIB.YFilter }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) SetFilter(yf yfilter.YFilter) { cISCOVLANIFTABLERELATIONSHIPMIB.YFilter = yf }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetGoName(yname string) string {
    if yname == "cviVlanInterfaceIndexTable" { return "Cvivlaninterfaceindextable" }
    return ""
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetSegmentPath() string {
    return "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB:CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cviVlanInterfaceIndexTable" {
        return &cISCOVLANIFTABLERELATIONSHIPMIB.Cvivlaninterfaceindextable
    }
    return nil
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cviVlanInterfaceIndexTable"] = &cISCOVLANIFTABLERELATIONSHIPMIB.Cvivlaninterfaceindextable
    return children
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetYangName() string { return "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB" }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) SetParent(parent types.Entity) { cISCOVLANIFTABLERELATIONSHIPMIB.parent = parent }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetParent() types.Entity { return cISCOVLANIFTABLERELATIONSHIPMIB.parent }

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetParentYangName() string { return "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB" }

// CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable
// The cviVlanInterfaceIndexTable provides a way to
// translate a VLAN-id in to an ifIndex, so that 
// the routed VLAN interface's routing configuration 
// can be obtained from interface entry in ipRouteTable.
// 
// Note that some routers can have interfaces to multiple
// VLAN management domains, and therefore can have multiple 
// routed VLAN interfaces which connect to different VLANs 
// having the same VLAN-id.  Thus, it is possible to have 
// multiple rows in this table for the same VLAN-id.
// 
// The cviVlanInterfaceIndexTable also provides a way
// to find the VLAN-id from an ifTable VLAN's ifIndex.
type CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a routed VLAN interface, its corresponding physical
    // port if any, and the ifTable entry for the routed VLAN interface.  Entries
    // are created by the agent when the routed VLAN interface is created. 
    // Operational status of routing does not affect the entries listed here.  For
    // routing configuration please refer to ipRouteTable.  Entries are deleted by
    // the agent when the routed VLAN interface is removed from the system
    // configuration. The type is slice of
    // CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry.
    Cvivlaninterfaceindexentry []CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetFilter() yfilter.YFilter { return cvivlaninterfaceindextable.YFilter }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) SetFilter(yf yfilter.YFilter) { cvivlaninterfaceindextable.YFilter = yf }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetGoName(yname string) string {
    if yname == "cviVlanInterfaceIndexEntry" { return "Cvivlaninterfaceindexentry" }
    return ""
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetSegmentPath() string {
    return "cviVlanInterfaceIndexTable"
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cviVlanInterfaceIndexEntry" {
        for _, c := range cvivlaninterfaceindextable.Cvivlaninterfaceindexentry {
            if cvivlaninterfaceindextable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry{}
        cvivlaninterfaceindextable.Cvivlaninterfaceindexentry = append(cvivlaninterfaceindextable.Cvivlaninterfaceindexentry, child)
        return &cvivlaninterfaceindextable.Cvivlaninterfaceindexentry[len(cvivlaninterfaceindextable.Cvivlaninterfaceindexentry)-1]
    }
    return nil
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvivlaninterfaceindextable.Cvivlaninterfaceindexentry {
        children[cvivlaninterfaceindextable.Cvivlaninterfaceindexentry[i].GetSegmentPath()] = &cvivlaninterfaceindextable.Cvivlaninterfaceindexentry[i]
    }
    return children
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetBundleName() string { return "cisco_ios_xe" }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetYangName() string { return "cviVlanInterfaceIndexTable" }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) SetParent(parent types.Entity) { cvivlaninterfaceindextable.parent = parent }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetParent() types.Entity { return cvivlaninterfaceindextable.parent }

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetParentYangName() string { return "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB" }

// CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry
// Each entry represents a routed VLAN interface, its
// corresponding physical port if any, and the ifTable entry
// for the routed VLAN interface.
// 
// Entries are created by the agent when the routed VLAN interface
// is created.  Operational status of routing does not affect
// the entries listed here.  For routing configuration please refer
// to ipRouteTable.
// 
// Entries are deleted by the agent when the routed VLAN interface
// is removed from the system configuration.
type CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN-id number of the routed VLAN interface.
    // The type is interface{} with range: 0..4095.
    Cvivlanid interface{}

    // This attribute is a key. For subinterfaces, this object is the ifIndex of
    // the physical interface for the subinterface.  For Switch Virtual Interfaces
    // (SVIs), this object is zero. The type is interface{} with range:
    // 0..2147483647.
    Cviphysicalifindex interface{}

    // The index for the ifTable entry associated with this routed VLAN interface.
    // The type is interface{} with range: 1..2147483647.
    Cviroutedvlanifindex interface{}
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetFilter() yfilter.YFilter { return cvivlaninterfaceindexentry.YFilter }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) SetFilter(yf yfilter.YFilter) { cvivlaninterfaceindexentry.YFilter = yf }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetGoName(yname string) string {
    if yname == "cviVlanId" { return "Cvivlanid" }
    if yname == "cviPhysicalIfIndex" { return "Cviphysicalifindex" }
    if yname == "cviRoutedVlanIfIndex" { return "Cviroutedvlanifindex" }
    return ""
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetSegmentPath() string {
    return "cviVlanInterfaceIndexEntry" + "[cviVlanId='" + fmt.Sprintf("%v", cvivlaninterfaceindexentry.Cvivlanid) + "']" + "[cviPhysicalIfIndex='" + fmt.Sprintf("%v", cvivlaninterfaceindexentry.Cviphysicalifindex) + "']"
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cviVlanId"] = cvivlaninterfaceindexentry.Cvivlanid
    leafs["cviPhysicalIfIndex"] = cvivlaninterfaceindexentry.Cviphysicalifindex
    leafs["cviRoutedVlanIfIndex"] = cvivlaninterfaceindexentry.Cviroutedvlanifindex
    return leafs
}

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetYangName() string { return "cviVlanInterfaceIndexEntry" }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) SetParent(parent types.Entity) { cvivlaninterfaceindexentry.parent = parent }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetParent() types.Entity { return cvivlaninterfaceindexentry.parent }

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetParentYangName() string { return "cviVlanInterfaceIndexTable" }

