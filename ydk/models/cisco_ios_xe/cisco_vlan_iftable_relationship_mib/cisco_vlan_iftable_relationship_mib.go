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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The cviVlanInterfaceIndexTable provides a way to translate a VLAN-id in to
    // an ifIndex, so that  the routed VLAN interface's routing configuration  can
    // be obtained from interface entry in ipRouteTable.  Note that some routers
    // can have interfaces to multiple VLAN management domains, and therefore can
    // have multiple  routed VLAN interfaces which connect to different VLANs 
    // having the same VLAN-id.  Thus, it is possible to have  multiple rows in
    // this table for the same VLAN-id.  The cviVlanInterfaceIndexTable also
    // provides a way to find the VLAN-id from an ifTable VLAN's ifIndex.
    CviVlanInterfaceIndexTable CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable
}

func (cISCOVLANIFTABLERELATIONSHIPMIB *CISCOVLANIFTABLERELATIONSHIPMIB) GetEntityData() *types.CommonEntityData {
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.YFilter = cISCOVLANIFTABLERELATIONSHIPMIB.YFilter
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.YangName = "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.ParentYangName = "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.SegmentPath = "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB:CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Children.Append("cviVlanInterfaceIndexTable", types.YChild{"CviVlanInterfaceIndexTable", &cISCOVLANIFTABLERELATIONSHIPMIB.CviVlanInterfaceIndexTable})
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.YListKeys = []string {}

    return &(cISCOVLANIFTABLERELATIONSHIPMIB.EntityData)
}

// CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable
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
type CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a routed VLAN interface, its corresponding physical
    // port if any, and the ifTable entry for the routed VLAN interface.  Entries
    // are created by the agent when the routed VLAN interface is created. 
    // Operational status of routing does not affect the entries listed here.  For
    // routing configuration please refer to ipRouteTable.  Entries are deleted by
    // the agent when the routed VLAN interface is removed from the system
    // configuration. The type is slice of
    // CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable_CviVlanInterfaceIndexEntry.
    CviVlanInterfaceIndexEntry []*CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable_CviVlanInterfaceIndexEntry
}

func (cviVlanInterfaceIndexTable *CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable) GetEntityData() *types.CommonEntityData {
    cviVlanInterfaceIndexTable.EntityData.YFilter = cviVlanInterfaceIndexTable.YFilter
    cviVlanInterfaceIndexTable.EntityData.YangName = "cviVlanInterfaceIndexTable"
    cviVlanInterfaceIndexTable.EntityData.BundleName = "cisco_ios_xe"
    cviVlanInterfaceIndexTable.EntityData.ParentYangName = "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
    cviVlanInterfaceIndexTable.EntityData.SegmentPath = "cviVlanInterfaceIndexTable"
    cviVlanInterfaceIndexTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cviVlanInterfaceIndexTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cviVlanInterfaceIndexTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cviVlanInterfaceIndexTable.EntityData.Children = types.NewOrderedMap()
    cviVlanInterfaceIndexTable.EntityData.Children.Append("cviVlanInterfaceIndexEntry", types.YChild{"CviVlanInterfaceIndexEntry", nil})
    for i := range cviVlanInterfaceIndexTable.CviVlanInterfaceIndexEntry {
        cviVlanInterfaceIndexTable.EntityData.Children.Append(types.GetSegmentPath(cviVlanInterfaceIndexTable.CviVlanInterfaceIndexEntry[i]), types.YChild{"CviVlanInterfaceIndexEntry", cviVlanInterfaceIndexTable.CviVlanInterfaceIndexEntry[i]})
    }
    cviVlanInterfaceIndexTable.EntityData.Leafs = types.NewOrderedMap()

    cviVlanInterfaceIndexTable.EntityData.YListKeys = []string {}

    return &(cviVlanInterfaceIndexTable.EntityData)
}

// CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable_CviVlanInterfaceIndexEntry
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
type CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable_CviVlanInterfaceIndexEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN-id number of the routed VLAN interface.
    // The type is interface{} with range: 0..4095.
    CviVlanId interface{}

    // This attribute is a key. For subinterfaces, this object is the ifIndex of
    // the physical interface for the subinterface.  For Switch Virtual Interfaces
    // (SVIs), this object is zero. The type is interface{} with range:
    // 0..2147483647.
    CviPhysicalIfIndex interface{}

    // The index for the ifTable entry associated with this routed VLAN interface.
    // The type is interface{} with range: 1..2147483647.
    CviRoutedVlanIfIndex interface{}
}

func (cviVlanInterfaceIndexEntry *CISCOVLANIFTABLERELATIONSHIPMIB_CviVlanInterfaceIndexTable_CviVlanInterfaceIndexEntry) GetEntityData() *types.CommonEntityData {
    cviVlanInterfaceIndexEntry.EntityData.YFilter = cviVlanInterfaceIndexEntry.YFilter
    cviVlanInterfaceIndexEntry.EntityData.YangName = "cviVlanInterfaceIndexEntry"
    cviVlanInterfaceIndexEntry.EntityData.BundleName = "cisco_ios_xe"
    cviVlanInterfaceIndexEntry.EntityData.ParentYangName = "cviVlanInterfaceIndexTable"
    cviVlanInterfaceIndexEntry.EntityData.SegmentPath = "cviVlanInterfaceIndexEntry" + types.AddKeyToken(cviVlanInterfaceIndexEntry.CviVlanId, "cviVlanId") + types.AddKeyToken(cviVlanInterfaceIndexEntry.CviPhysicalIfIndex, "cviPhysicalIfIndex")
    cviVlanInterfaceIndexEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cviVlanInterfaceIndexEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cviVlanInterfaceIndexEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cviVlanInterfaceIndexEntry.EntityData.Children = types.NewOrderedMap()
    cviVlanInterfaceIndexEntry.EntityData.Leafs = types.NewOrderedMap()
    cviVlanInterfaceIndexEntry.EntityData.Leafs.Append("cviVlanId", types.YLeaf{"CviVlanId", cviVlanInterfaceIndexEntry.CviVlanId})
    cviVlanInterfaceIndexEntry.EntityData.Leafs.Append("cviPhysicalIfIndex", types.YLeaf{"CviPhysicalIfIndex", cviVlanInterfaceIndexEntry.CviPhysicalIfIndex})
    cviVlanInterfaceIndexEntry.EntityData.Leafs.Append("cviRoutedVlanIfIndex", types.YLeaf{"CviRoutedVlanIfIndex", cviVlanInterfaceIndexEntry.CviRoutedVlanIfIndex})

    cviVlanInterfaceIndexEntry.EntityData.YListKeys = []string {"CviVlanId", "CviPhysicalIfIndex"}

    return &(cviVlanInterfaceIndexEntry.EntityData)
}

