// The MIB module to describe generic objects for
// ethernet-like network interfaces. 
// 
// This MIB provides ethernet-like network interfaces 
// information that are either excluded by EtherLike-MIB 
// or specific to Cisco products.
package cisco_etherlike_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_etherlike_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ETHERLIKE-EXT-MIB CISCO-ETHERLIKE-EXT-MIB}", reflect.TypeOf(CISCOETHERLIKEEXTMIB{}))
    ydk.RegisterEntity("CISCO-ETHERLIKE-EXT-MIB:CISCO-ETHERLIKE-EXT-MIB", reflect.TypeOf(CISCOETHERLIKEEXTMIB{}))
}

// CISCOETHERLIKEEXTMIB
type CISCOETHERLIKEEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of additional descriptive and status information about the MAC
    // Control PAUSE  function on the ethernet-like interfaces  attached to a
    // particular system, in extension to dot3PauseTable in EtherLike-MIB. There
    // will be  one row in this table for each ethernet-like  interface in the
    // system which supports the MAC  Control PAUSE function.
    Ceedot3Pauseexttable CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable

    // This table provides the subinterface related information associated to the
    // Ethernet-like interfaces.  The subinterface is a division of one physical
    // interface into multiple logical interfaces. As an example of what a typical
    // subinterface setup might look like on a device, a single Ethernet port such
    // as GigabitEthernet0/0 would be subdivided into Gi0/0.1, Gi0/0.2, Gi0/0.3
    // and so on, each one performing as if it were a separate interface.
    Ceesubinterfacetable CISCOETHERLIKEEXTMIB_Ceesubinterfacetable
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCOETHERLIKEEXTMIB.EntityData.YFilter = cISCOETHERLIKEEXTMIB.YFilter
    cISCOETHERLIKEEXTMIB.EntityData.YangName = "CISCO-ETHERLIKE-EXT-MIB"
    cISCOETHERLIKEEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOETHERLIKEEXTMIB.EntityData.ParentYangName = "CISCO-ETHERLIKE-EXT-MIB"
    cISCOETHERLIKEEXTMIB.EntityData.SegmentPath = "CISCO-ETHERLIKE-EXT-MIB:CISCO-ETHERLIKE-EXT-MIB"
    cISCOETHERLIKEEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOETHERLIKEEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOETHERLIKEEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOETHERLIKEEXTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOETHERLIKEEXTMIB.EntityData.Children["ceeDot3PauseExtTable"] = types.YChild{"Ceedot3Pauseexttable", &cISCOETHERLIKEEXTMIB.Ceedot3Pauseexttable}
    cISCOETHERLIKEEXTMIB.EntityData.Children["ceeSubInterfaceTable"] = types.YChild{"Ceesubinterfacetable", &cISCOETHERLIKEEXTMIB.Ceesubinterfacetable}
    cISCOETHERLIKEEXTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOETHERLIKEEXTMIB.EntityData)
}

// CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable
// A list of additional descriptive and status
// information about the MAC Control PAUSE 
// function on the ethernet-like interfaces 
// attached to a particular system, in extension to
// dot3PauseTable in EtherLike-MIB. There will be 
// one row in this table for each ethernet-like 
// interface in the system which supports the MAC 
// Control PAUSE function.
type CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing additional information about the MAC
    // Control PAUSE function  on a single ethernet-like interface, in extension 
    // to dot3PauseEntry in Etherlike-MIB. The type is slice of
    // CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry.
    Ceedot3Pauseextentry []CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetEntityData() *types.CommonEntityData {
    ceedot3Pauseexttable.EntityData.YFilter = ceedot3Pauseexttable.YFilter
    ceedot3Pauseexttable.EntityData.YangName = "ceeDot3PauseExtTable"
    ceedot3Pauseexttable.EntityData.BundleName = "cisco_ios_xe"
    ceedot3Pauseexttable.EntityData.ParentYangName = "CISCO-ETHERLIKE-EXT-MIB"
    ceedot3Pauseexttable.EntityData.SegmentPath = "ceeDot3PauseExtTable"
    ceedot3Pauseexttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceedot3Pauseexttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceedot3Pauseexttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceedot3Pauseexttable.EntityData.Children = make(map[string]types.YChild)
    ceedot3Pauseexttable.EntityData.Children["ceeDot3PauseExtEntry"] = types.YChild{"Ceedot3Pauseextentry", nil}
    for i := range ceedot3Pauseexttable.Ceedot3Pauseextentry {
        ceedot3Pauseexttable.EntityData.Children[types.GetSegmentPath(&ceedot3Pauseexttable.Ceedot3Pauseextentry[i])] = types.YChild{"Ceedot3Pauseextentry", &ceedot3Pauseexttable.Ceedot3Pauseextentry[i]}
    }
    ceedot3Pauseexttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceedot3Pauseexttable.EntityData)
}

// CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry
// An entry in the table, containing additional
// information about the MAC Control PAUSE function 
// on a single ethernet-like interface, in extension 
// to dot3PauseEntry in Etherlike-MIB.
type CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3Statstable_Dot3Statsentry_Dot3Statsindex
    Dot3Statsindex interface{}

    // Indicates preference to send or process pause frames on this interface.
    // txDesired(0)  -  indicates preference to send pause                  
    // frames, but autonegotiates flow                   control. This bit can
    // only be                   turned on when the corresponding                 
    // instance of dot3PauseAdminMode                   has the value of
    // 'enabledXmit' or                   'enabledXmitAndRcv'. rxDesired(1)  - 
    // indicates preference to process                   pause frames, but
    // autonegotiates                   flow control. This bit can only be        
    // turned on when the corresponding                   instance of
    // dot3PauseAdminMode                   has the value of 'enabledRcv' or      
    // 'enabledXmitAndRcv'. The type is map[string]bool.
    Ceedot3Pauseextadminmode interface{}

    // Provides additional information about the flow control operational status
    // on this interface. txDisagree(0) - the transmit pause function on          
    // this interface is disabled due to                  disagreement from the
    // far end on                  negotiation. rxDisagree(1) - the receive pause
    // function on                   this interface is disabled due to            
    // disagreement from the far end on                  negotiation. txDesired(2)
    // - the transmit pause function on                  this interface is
    // desired. rxDesired(3)  - the receive pause function on                  
    // this interface is desired. The type is map[string]bool.
    Ceedot3Pauseextopermode interface{}
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetEntityData() *types.CommonEntityData {
    ceedot3Pauseextentry.EntityData.YFilter = ceedot3Pauseextentry.YFilter
    ceedot3Pauseextentry.EntityData.YangName = "ceeDot3PauseExtEntry"
    ceedot3Pauseextentry.EntityData.BundleName = "cisco_ios_xe"
    ceedot3Pauseextentry.EntityData.ParentYangName = "ceeDot3PauseExtTable"
    ceedot3Pauseextentry.EntityData.SegmentPath = "ceeDot3PauseExtEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", ceedot3Pauseextentry.Dot3Statsindex) + "']"
    ceedot3Pauseextentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceedot3Pauseextentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceedot3Pauseextentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceedot3Pauseextentry.EntityData.Children = make(map[string]types.YChild)
    ceedot3Pauseextentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceedot3Pauseextentry.EntityData.Leafs["dot3StatsIndex"] = types.YLeaf{"Dot3Statsindex", ceedot3Pauseextentry.Dot3Statsindex}
    ceedot3Pauseextentry.EntityData.Leafs["ceeDot3PauseExtAdminMode"] = types.YLeaf{"Ceedot3Pauseextadminmode", ceedot3Pauseextentry.Ceedot3Pauseextadminmode}
    ceedot3Pauseextentry.EntityData.Leafs["ceeDot3PauseExtOperMode"] = types.YLeaf{"Ceedot3Pauseextopermode", ceedot3Pauseextentry.Ceedot3Pauseextopermode}
    return &(ceedot3Pauseextentry.EntityData)
}

// CISCOETHERLIKEEXTMIB_Ceesubinterfacetable
// This table provides the subinterface related information
// associated to the Ethernet-like interfaces.
// 
// The subinterface is a division of one physical interface into
// multiple logical interfaces. As an example of what a typical
// subinterface setup might look like on a device, a single
// Ethernet port such as GigabitEthernet0/0 would be subdivided
// into Gi0/0.1, Gi0/0.2, Gi0/0.3 and so on, each one performing as
// if it were a separate interface.
type CISCOETHERLIKEEXTMIB_Ceesubinterfacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains a row for each Ethernet-like interface by it's ifTable
    // ifIndex in the system, which supports the sub-interface.  An entry is
    // created by an agent, when it detects a Ethernet-like interface is created
    // in ifTable and it  can support sub-interface.  An entry is deleted by an
    // agent, when the ifTable entry associated to the Ethernet-like interface is
    // deleted. Typically, when the card is removed from the device. The type is
    // slice of CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry.
    Ceesubinterfaceentry []CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetEntityData() *types.CommonEntityData {
    ceesubinterfacetable.EntityData.YFilter = ceesubinterfacetable.YFilter
    ceesubinterfacetable.EntityData.YangName = "ceeSubInterfaceTable"
    ceesubinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    ceesubinterfacetable.EntityData.ParentYangName = "CISCO-ETHERLIKE-EXT-MIB"
    ceesubinterfacetable.EntityData.SegmentPath = "ceeSubInterfaceTable"
    ceesubinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceesubinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceesubinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceesubinterfacetable.EntityData.Children = make(map[string]types.YChild)
    ceesubinterfacetable.EntityData.Children["ceeSubInterfaceEntry"] = types.YChild{"Ceesubinterfaceentry", nil}
    for i := range ceesubinterfacetable.Ceesubinterfaceentry {
        ceesubinterfacetable.EntityData.Children[types.GetSegmentPath(&ceesubinterfacetable.Ceesubinterfaceentry[i])] = types.YChild{"Ceesubinterfaceentry", &ceesubinterfacetable.Ceesubinterfaceentry[i]}
    }
    ceesubinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceesubinterfacetable.EntityData)
}

// CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry
// This table contains a row for each Ethernet-like interface
// by it's ifTable ifIndex in the system, which supports the
// sub-interface.
// 
// An entry is created by an agent, when it detects a
// Ethernet-like interface is created in ifTable and it 
// can support sub-interface.
// 
// An entry is deleted by an agent, when the ifTable entry
// associated to the Ethernet-like interface is deleted.
// Typically, when the card is removed from the device.
type CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object represents the number of subinterfaces created on a
    // Ethernet-like interface. The type is interface{} with range: 0..4294967295.
    // Units are subifs.
    Ceesubinterfacecount interface{}
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetEntityData() *types.CommonEntityData {
    ceesubinterfaceentry.EntityData.YFilter = ceesubinterfaceentry.YFilter
    ceesubinterfaceentry.EntityData.YangName = "ceeSubInterfaceEntry"
    ceesubinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    ceesubinterfaceentry.EntityData.ParentYangName = "ceeSubInterfaceTable"
    ceesubinterfaceentry.EntityData.SegmentPath = "ceeSubInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", ceesubinterfaceentry.Ifindex) + "']"
    ceesubinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceesubinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceesubinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceesubinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    ceesubinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceesubinterfaceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", ceesubinterfaceentry.Ifindex}
    ceesubinterfaceentry.EntityData.Leafs["ceeSubInterfaceCount"] = types.YLeaf{"Ceesubinterfacecount", ceesubinterfaceentry.Ceesubinterfacecount}
    return &(ceesubinterfaceentry.EntityData)
}

