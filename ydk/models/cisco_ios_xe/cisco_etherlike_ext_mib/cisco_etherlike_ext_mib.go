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
    parent types.Entity
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

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetFilter() yfilter.YFilter { return cISCOETHERLIKEEXTMIB.YFilter }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) SetFilter(yf yfilter.YFilter) { cISCOETHERLIKEEXTMIB.YFilter = yf }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetGoName(yname string) string {
    if yname == "ceeDot3PauseExtTable" { return "Ceedot3Pauseexttable" }
    if yname == "ceeSubInterfaceTable" { return "Ceesubinterfacetable" }
    return ""
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetSegmentPath() string {
    return "CISCO-ETHERLIKE-EXT-MIB:CISCO-ETHERLIKE-EXT-MIB"
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceeDot3PauseExtTable" {
        return &cISCOETHERLIKEEXTMIB.Ceedot3Pauseexttable
    }
    if childYangName == "ceeSubInterfaceTable" {
        return &cISCOETHERLIKEEXTMIB.Ceesubinterfacetable
    }
    return nil
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ceeDot3PauseExtTable"] = &cISCOETHERLIKEEXTMIB.Ceedot3Pauseexttable
    children["ceeSubInterfaceTable"] = &cISCOETHERLIKEEXTMIB.Ceesubinterfacetable
    return children
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetYangName() string { return "CISCO-ETHERLIKE-EXT-MIB" }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) SetParent(parent types.Entity) { cISCOETHERLIKEEXTMIB.parent = parent }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetParent() types.Entity { return cISCOETHERLIKEEXTMIB.parent }

func (cISCOETHERLIKEEXTMIB *CISCOETHERLIKEEXTMIB) GetParentYangName() string { return "CISCO-ETHERLIKE-EXT-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table, containing additional information about the MAC
    // Control PAUSE function  on a single ethernet-like interface, in extension 
    // to dot3PauseEntry in Etherlike-MIB. The type is slice of
    // CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry.
    Ceedot3Pauseextentry []CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetFilter() yfilter.YFilter { return ceedot3Pauseexttable.YFilter }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) SetFilter(yf yfilter.YFilter) { ceedot3Pauseexttable.YFilter = yf }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetGoName(yname string) string {
    if yname == "ceeDot3PauseExtEntry" { return "Ceedot3Pauseextentry" }
    return ""
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetSegmentPath() string {
    return "ceeDot3PauseExtTable"
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceeDot3PauseExtEntry" {
        for _, c := range ceedot3Pauseexttable.Ceedot3Pauseextentry {
            if ceedot3Pauseexttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry{}
        ceedot3Pauseexttable.Ceedot3Pauseextentry = append(ceedot3Pauseexttable.Ceedot3Pauseextentry, child)
        return &ceedot3Pauseexttable.Ceedot3Pauseextentry[len(ceedot3Pauseexttable.Ceedot3Pauseextentry)-1]
    }
    return nil
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceedot3Pauseexttable.Ceedot3Pauseextentry {
        children[ceedot3Pauseexttable.Ceedot3Pauseextentry[i].GetSegmentPath()] = &ceedot3Pauseexttable.Ceedot3Pauseextentry[i]
    }
    return children
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetBundleName() string { return "cisco_ios_xe" }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetYangName() string { return "ceeDot3PauseExtTable" }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) SetParent(parent types.Entity) { ceedot3Pauseexttable.parent = parent }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetParent() types.Entity { return ceedot3Pauseexttable.parent }

func (ceedot3Pauseexttable *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable) GetParentYangName() string { return "CISCO-ETHERLIKE-EXT-MIB" }

// CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry
// An entry in the table, containing additional
// information about the MAC Control PAUSE function 
// on a single ethernet-like interface, in extension 
// to dot3PauseEntry in Etherlike-MIB.
type CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry struct {
    parent types.Entity
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

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetFilter() yfilter.YFilter { return ceedot3Pauseextentry.YFilter }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) SetFilter(yf yfilter.YFilter) { ceedot3Pauseextentry.YFilter = yf }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetGoName(yname string) string {
    if yname == "dot3StatsIndex" { return "Dot3Statsindex" }
    if yname == "ceeDot3PauseExtAdminMode" { return "Ceedot3Pauseextadminmode" }
    if yname == "ceeDot3PauseExtOperMode" { return "Ceedot3Pauseextopermode" }
    return ""
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetSegmentPath() string {
    return "ceeDot3PauseExtEntry" + "[dot3StatsIndex='" + fmt.Sprintf("%v", ceedot3Pauseextentry.Dot3Statsindex) + "']"
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dot3StatsIndex"] = ceedot3Pauseextentry.Dot3Statsindex
    leafs["ceeDot3PauseExtAdminMode"] = ceedot3Pauseextentry.Ceedot3Pauseextadminmode
    leafs["ceeDot3PauseExtOperMode"] = ceedot3Pauseextentry.Ceedot3Pauseextopermode
    return leafs
}

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetYangName() string { return "ceeDot3PauseExtEntry" }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) SetParent(parent types.Entity) { ceedot3Pauseextentry.parent = parent }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetParent() types.Entity { return ceedot3Pauseextentry.parent }

func (ceedot3Pauseextentry *CISCOETHERLIKEEXTMIB_Ceedot3Pauseexttable_Ceedot3Pauseextentry) GetParentYangName() string { return "ceeDot3PauseExtTable" }

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
    parent types.Entity
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

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetFilter() yfilter.YFilter { return ceesubinterfacetable.YFilter }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) SetFilter(yf yfilter.YFilter) { ceesubinterfacetable.YFilter = yf }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetGoName(yname string) string {
    if yname == "ceeSubInterfaceEntry" { return "Ceesubinterfaceentry" }
    return ""
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetSegmentPath() string {
    return "ceeSubInterfaceTable"
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceeSubInterfaceEntry" {
        for _, c := range ceesubinterfacetable.Ceesubinterfaceentry {
            if ceesubinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry{}
        ceesubinterfacetable.Ceesubinterfaceentry = append(ceesubinterfacetable.Ceesubinterfaceentry, child)
        return &ceesubinterfacetable.Ceesubinterfaceentry[len(ceesubinterfacetable.Ceesubinterfaceentry)-1]
    }
    return nil
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceesubinterfacetable.Ceesubinterfaceentry {
        children[ceesubinterfacetable.Ceesubinterfaceentry[i].GetSegmentPath()] = &ceesubinterfacetable.Ceesubinterfaceentry[i]
    }
    return children
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetYangName() string { return "ceeSubInterfaceTable" }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) SetParent(parent types.Entity) { ceesubinterfacetable.parent = parent }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetParent() types.Entity { return ceesubinterfacetable.parent }

func (ceesubinterfacetable *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable) GetParentYangName() string { return "CISCO-ETHERLIKE-EXT-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object represents the number of subinterfaces created on a
    // Ethernet-like interface. The type is interface{} with range: 0..4294967295.
    // Units are subifs.
    Ceesubinterfacecount interface{}
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetFilter() yfilter.YFilter { return ceesubinterfaceentry.YFilter }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) SetFilter(yf yfilter.YFilter) { ceesubinterfaceentry.YFilter = yf }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "ceeSubInterfaceCount" { return "Ceesubinterfacecount" }
    return ""
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetSegmentPath() string {
    return "ceeSubInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", ceesubinterfaceentry.Ifindex) + "']"
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = ceesubinterfaceentry.Ifindex
    leafs["ceeSubInterfaceCount"] = ceesubinterfaceentry.Ceesubinterfacecount
    return leafs
}

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetYangName() string { return "ceeSubInterfaceEntry" }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) SetParent(parent types.Entity) { ceesubinterfaceentry.parent = parent }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetParent() types.Entity { return ceesubinterfaceentry.parent }

func (ceesubinterfaceentry *CISCOETHERLIKEEXTMIB_Ceesubinterfacetable_Ceesubinterfaceentry) GetParentYangName() string { return "ceeSubInterfaceTable" }

