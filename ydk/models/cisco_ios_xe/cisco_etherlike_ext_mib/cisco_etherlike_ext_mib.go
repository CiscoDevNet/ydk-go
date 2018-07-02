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
    CeeDot3PauseExtTable CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable

    // This table provides the subinterface related information associated to the
    // Ethernet-like interfaces.  The subinterface is a division of one physical
    // interface into multiple logical interfaces. As an example of what a typical
    // subinterface setup might look like on a device, a single Ethernet port such
    // as GigabitEthernet0/0 would be subdivided into Gi0/0.1, Gi0/0.2, Gi0/0.3
    // and so on, each one performing as if it were a separate interface.
    CeeSubInterfaceTable CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable
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

    cISCOETHERLIKEEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOETHERLIKEEXTMIB.EntityData.Children.Append("ceeDot3PauseExtTable", types.YChild{"CeeDot3PauseExtTable", &cISCOETHERLIKEEXTMIB.CeeDot3PauseExtTable})
    cISCOETHERLIKEEXTMIB.EntityData.Children.Append("ceeSubInterfaceTable", types.YChild{"CeeSubInterfaceTable", &cISCOETHERLIKEEXTMIB.CeeSubInterfaceTable})
    cISCOETHERLIKEEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOETHERLIKEEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCOETHERLIKEEXTMIB.EntityData)
}

// CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable
// A list of additional descriptive and status
// information about the MAC Control PAUSE 
// function on the ethernet-like interfaces 
// attached to a particular system, in extension to
// dot3PauseTable in EtherLike-MIB. There will be 
// one row in this table for each ethernet-like 
// interface in the system which supports the MAC 
// Control PAUSE function.
type CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table, containing additional information about the MAC
    // Control PAUSE function  on a single ethernet-like interface, in extension 
    // to dot3PauseEntry in Etherlike-MIB. The type is slice of
    // CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable_CeeDot3PauseExtEntry.
    CeeDot3PauseExtEntry []*CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable_CeeDot3PauseExtEntry
}

func (ceeDot3PauseExtTable *CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable) GetEntityData() *types.CommonEntityData {
    ceeDot3PauseExtTable.EntityData.YFilter = ceeDot3PauseExtTable.YFilter
    ceeDot3PauseExtTable.EntityData.YangName = "ceeDot3PauseExtTable"
    ceeDot3PauseExtTable.EntityData.BundleName = "cisco_ios_xe"
    ceeDot3PauseExtTable.EntityData.ParentYangName = "CISCO-ETHERLIKE-EXT-MIB"
    ceeDot3PauseExtTable.EntityData.SegmentPath = "ceeDot3PauseExtTable"
    ceeDot3PauseExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceeDot3PauseExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceeDot3PauseExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceeDot3PauseExtTable.EntityData.Children = types.NewOrderedMap()
    ceeDot3PauseExtTable.EntityData.Children.Append("ceeDot3PauseExtEntry", types.YChild{"CeeDot3PauseExtEntry", nil})
    for i := range ceeDot3PauseExtTable.CeeDot3PauseExtEntry {
        ceeDot3PauseExtTable.EntityData.Children.Append(types.GetSegmentPath(ceeDot3PauseExtTable.CeeDot3PauseExtEntry[i]), types.YChild{"CeeDot3PauseExtEntry", ceeDot3PauseExtTable.CeeDot3PauseExtEntry[i]})
    }
    ceeDot3PauseExtTable.EntityData.Leafs = types.NewOrderedMap()

    ceeDot3PauseExtTable.EntityData.YListKeys = []string {}

    return &(ceeDot3PauseExtTable.EntityData)
}

// CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable_CeeDot3PauseExtEntry
// An entry in the table, containing additional
// information about the MAC Control PAUSE function 
// on a single ethernet-like interface, in extension 
// to dot3PauseEntry in Etherlike-MIB.
type CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable_CeeDot3PauseExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // etherlike_mib.EtherLikeMIB_Dot3StatsTable_Dot3StatsEntry_Dot3StatsIndex
    Dot3StatsIndex interface{}

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
    CeeDot3PauseExtAdminMode interface{}

    // Provides additional information about the flow control operational status
    // on this interface. txDisagree(0) - the transmit pause function on          
    // this interface is disabled due to                  disagreement from the
    // far end on                  negotiation. rxDisagree(1) - the receive pause
    // function on                   this interface is disabled due to            
    // disagreement from the far end on                  negotiation. txDesired(2)
    // - the transmit pause function on                  this interface is
    // desired. rxDesired(3)  - the receive pause function on                  
    // this interface is desired. The type is map[string]bool.
    CeeDot3PauseExtOperMode interface{}
}

func (ceeDot3PauseExtEntry *CISCOETHERLIKEEXTMIB_CeeDot3PauseExtTable_CeeDot3PauseExtEntry) GetEntityData() *types.CommonEntityData {
    ceeDot3PauseExtEntry.EntityData.YFilter = ceeDot3PauseExtEntry.YFilter
    ceeDot3PauseExtEntry.EntityData.YangName = "ceeDot3PauseExtEntry"
    ceeDot3PauseExtEntry.EntityData.BundleName = "cisco_ios_xe"
    ceeDot3PauseExtEntry.EntityData.ParentYangName = "ceeDot3PauseExtTable"
    ceeDot3PauseExtEntry.EntityData.SegmentPath = "ceeDot3PauseExtEntry" + types.AddKeyToken(ceeDot3PauseExtEntry.Dot3StatsIndex, "dot3StatsIndex")
    ceeDot3PauseExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceeDot3PauseExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceeDot3PauseExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceeDot3PauseExtEntry.EntityData.Children = types.NewOrderedMap()
    ceeDot3PauseExtEntry.EntityData.Leafs = types.NewOrderedMap()
    ceeDot3PauseExtEntry.EntityData.Leafs.Append("dot3StatsIndex", types.YLeaf{"Dot3StatsIndex", ceeDot3PauseExtEntry.Dot3StatsIndex})
    ceeDot3PauseExtEntry.EntityData.Leafs.Append("ceeDot3PauseExtAdminMode", types.YLeaf{"CeeDot3PauseExtAdminMode", ceeDot3PauseExtEntry.CeeDot3PauseExtAdminMode})
    ceeDot3PauseExtEntry.EntityData.Leafs.Append("ceeDot3PauseExtOperMode", types.YLeaf{"CeeDot3PauseExtOperMode", ceeDot3PauseExtEntry.CeeDot3PauseExtOperMode})

    ceeDot3PauseExtEntry.EntityData.YListKeys = []string {"Dot3StatsIndex"}

    return &(ceeDot3PauseExtEntry.EntityData)
}

// CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable
// This table provides the subinterface related information
// associated to the Ethernet-like interfaces.
// 
// The subinterface is a division of one physical interface into
// multiple logical interfaces. As an example of what a typical
// subinterface setup might look like on a device, a single
// Ethernet port such as GigabitEthernet0/0 would be subdivided
// into Gi0/0.1, Gi0/0.2, Gi0/0.3 and so on, each one performing as
// if it were a separate interface.
type CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains a row for each Ethernet-like interface by it's ifTable
    // ifIndex in the system, which supports the sub-interface.  An entry is
    // created by an agent, when it detects a Ethernet-like interface is created
    // in ifTable and it  can support sub-interface.  An entry is deleted by an
    // agent, when the ifTable entry associated to the Ethernet-like interface is
    // deleted. Typically, when the card is removed from the device. The type is
    // slice of CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable_CeeSubInterfaceEntry.
    CeeSubInterfaceEntry []*CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable_CeeSubInterfaceEntry
}

func (ceeSubInterfaceTable *CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable) GetEntityData() *types.CommonEntityData {
    ceeSubInterfaceTable.EntityData.YFilter = ceeSubInterfaceTable.YFilter
    ceeSubInterfaceTable.EntityData.YangName = "ceeSubInterfaceTable"
    ceeSubInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    ceeSubInterfaceTable.EntityData.ParentYangName = "CISCO-ETHERLIKE-EXT-MIB"
    ceeSubInterfaceTable.EntityData.SegmentPath = "ceeSubInterfaceTable"
    ceeSubInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceeSubInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceeSubInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceeSubInterfaceTable.EntityData.Children = types.NewOrderedMap()
    ceeSubInterfaceTable.EntityData.Children.Append("ceeSubInterfaceEntry", types.YChild{"CeeSubInterfaceEntry", nil})
    for i := range ceeSubInterfaceTable.CeeSubInterfaceEntry {
        ceeSubInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(ceeSubInterfaceTable.CeeSubInterfaceEntry[i]), types.YChild{"CeeSubInterfaceEntry", ceeSubInterfaceTable.CeeSubInterfaceEntry[i]})
    }
    ceeSubInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    ceeSubInterfaceTable.EntityData.YListKeys = []string {}

    return &(ceeSubInterfaceTable.EntityData)
}

// CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable_CeeSubInterfaceEntry
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
type CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable_CeeSubInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object represents the number of subinterfaces created on a
    // Ethernet-like interface. The type is interface{} with range: 0..4294967295.
    // Units are subifs.
    CeeSubInterfaceCount interface{}
}

func (ceeSubInterfaceEntry *CISCOETHERLIKEEXTMIB_CeeSubInterfaceTable_CeeSubInterfaceEntry) GetEntityData() *types.CommonEntityData {
    ceeSubInterfaceEntry.EntityData.YFilter = ceeSubInterfaceEntry.YFilter
    ceeSubInterfaceEntry.EntityData.YangName = "ceeSubInterfaceEntry"
    ceeSubInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    ceeSubInterfaceEntry.EntityData.ParentYangName = "ceeSubInterfaceTable"
    ceeSubInterfaceEntry.EntityData.SegmentPath = "ceeSubInterfaceEntry" + types.AddKeyToken(ceeSubInterfaceEntry.IfIndex, "ifIndex")
    ceeSubInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceeSubInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceeSubInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceeSubInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    ceeSubInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    ceeSubInterfaceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ceeSubInterfaceEntry.IfIndex})
    ceeSubInterfaceEntry.EntityData.Leafs.Append("ceeSubInterfaceCount", types.YLeaf{"CeeSubInterfaceCount", ceeSubInterfaceEntry.CeeSubInterfaceCount})

    ceeSubInterfaceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(ceeSubInterfaceEntry.EntityData)
}

