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
    Cvivlaninterfaceindextable CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable
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

    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Children["cviVlanInterfaceIndexTable"] = types.YChild{"Cvivlaninterfaceindextable", &cISCOVLANIFTABLERELATIONSHIPMIB.Cvivlaninterfaceindextable}
    cISCOVLANIFTABLERELATIONSHIPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVLANIFTABLERELATIONSHIPMIB.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cvivlaninterfaceindextable *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable) GetEntityData() *types.CommonEntityData {
    cvivlaninterfaceindextable.EntityData.YFilter = cvivlaninterfaceindextable.YFilter
    cvivlaninterfaceindextable.EntityData.YangName = "cviVlanInterfaceIndexTable"
    cvivlaninterfaceindextable.EntityData.BundleName = "cisco_ios_xe"
    cvivlaninterfaceindextable.EntityData.ParentYangName = "CISCO-VLAN-IFTABLE-RELATIONSHIP-MIB"
    cvivlaninterfaceindextable.EntityData.SegmentPath = "cviVlanInterfaceIndexTable"
    cvivlaninterfaceindextable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvivlaninterfaceindextable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvivlaninterfaceindextable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvivlaninterfaceindextable.EntityData.Children = make(map[string]types.YChild)
    cvivlaninterfaceindextable.EntityData.Children["cviVlanInterfaceIndexEntry"] = types.YChild{"Cvivlaninterfaceindexentry", nil}
    for i := range cvivlaninterfaceindextable.Cvivlaninterfaceindexentry {
        cvivlaninterfaceindextable.EntityData.Children[types.GetSegmentPath(&cvivlaninterfaceindextable.Cvivlaninterfaceindexentry[i])] = types.YChild{"Cvivlaninterfaceindexentry", &cvivlaninterfaceindextable.Cvivlaninterfaceindexentry[i]}
    }
    cvivlaninterfaceindextable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cvivlaninterfaceindextable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cvivlaninterfaceindexentry *CISCOVLANIFTABLERELATIONSHIPMIB_Cvivlaninterfaceindextable_Cvivlaninterfaceindexentry) GetEntityData() *types.CommonEntityData {
    cvivlaninterfaceindexentry.EntityData.YFilter = cvivlaninterfaceindexentry.YFilter
    cvivlaninterfaceindexentry.EntityData.YangName = "cviVlanInterfaceIndexEntry"
    cvivlaninterfaceindexentry.EntityData.BundleName = "cisco_ios_xe"
    cvivlaninterfaceindexentry.EntityData.ParentYangName = "cviVlanInterfaceIndexTable"
    cvivlaninterfaceindexentry.EntityData.SegmentPath = "cviVlanInterfaceIndexEntry" + "[cviVlanId='" + fmt.Sprintf("%v", cvivlaninterfaceindexentry.Cvivlanid) + "']" + "[cviPhysicalIfIndex='" + fmt.Sprintf("%v", cvivlaninterfaceindexentry.Cviphysicalifindex) + "']"
    cvivlaninterfaceindexentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cvivlaninterfaceindexentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cvivlaninterfaceindexentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cvivlaninterfaceindexentry.EntityData.Children = make(map[string]types.YChild)
    cvivlaninterfaceindexentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cvivlaninterfaceindexentry.EntityData.Leafs["cviVlanId"] = types.YLeaf{"Cvivlanid", cvivlaninterfaceindexentry.Cvivlanid}
    cvivlaninterfaceindexentry.EntityData.Leafs["cviPhysicalIfIndex"] = types.YLeaf{"Cviphysicalifindex", cvivlaninterfaceindexentry.Cviphysicalifindex}
    cvivlaninterfaceindexentry.EntityData.Leafs["cviRoutedVlanIfIndex"] = types.YLeaf{"Cviroutedvlanifindex", cvivlaninterfaceindexentry.Cviroutedvlanifindex}
    return &(cvivlaninterfaceindexentry.EntityData)
}

