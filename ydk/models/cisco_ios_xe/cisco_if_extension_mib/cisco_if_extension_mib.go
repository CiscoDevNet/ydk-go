// A MIB module for extending the IF-MIB (RFC2863)
// to add objects which provide additional information
// about interfaces not available in other MIBS.
// This MIB replaces the OLD-CISCO-INTERFACES-MIB.
// 
// GLOSSARY :
// 
// Virtual Switch - A physical switch partitioned into 
//         multiple logical switches.
// 
// Interface Sharing - An interface can be shared among 
//         multiple virtual switches.
// 
// Speed Group - An interface is capable of operating in any one of
// the speed range depending on the capability of the hardware.
// 
// Virtual Link (VL) - Virtual Link is a logical connectivity 
//     between two end points. A physical interface can 
//     have multiple Virtual Links.
// 
// No Drop Virtual Link - According to 802.3 standard, 
//     No drop specifies lossless service on a virtual link.
// 
// Drop Virtual Link - According to 802.3 standard,
//     Traffic drop may occur on this virtual Link.
package cisco_if_extension_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_if_extension_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IF-EXTENSION-MIB CISCO-IF-EXTENSION-MIB}", reflect.TypeOf(CISCOIFEXTENSIONMIB{}))
    ydk.RegisterEntity("CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB", reflect.TypeOf(CISCOIFEXTENSIONMIB{}))
}

// IfIndexPersistenceState represents take all the three values enable/disable/global.
type IfIndexPersistenceState string

const (
    IfIndexPersistenceState_disable IfIndexPersistenceState = "disable"

    IfIndexPersistenceState_enable IfIndexPersistenceState = "enable"

    IfIndexPersistenceState_global IfIndexPersistenceState = "global"
)

// CISCOIFEXTENSIONMIB
type CISCOIFEXTENSIONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscoifextsystemconfig CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig

    // This  table contains interface packet statistics which are not available in
    // IF-MIB(RFC2863).   As an example, some interfaces to which objects in this
    // table are applicable are as follows :          o Ethernet         o
    // FastEthernet         o ATM         o BRI         o Sonet         o
    // GigabitEthernet  Some objects defined in this table may be  applicable to
    // physical interfaces only. As a result, this table may be sparse for some
    // logical interfaces.
    Cieifpacketstatstable CISCOIFEXTENSIONMIB_Cieifpacketstatstable

    // This  table contains objects which provide more information about interface
    // properties not available in IF-MIB (RFC 2863).  Some objects defined in
    // this table may be applicable to physical interfaces only. As a result, this
    // table may be sparse for logical interfaces.
    Cieifinterfacetable CISCOIFEXTENSIONMIB_Cieifinterfacetable

    // This table contains objects for providing the 'ifIndex', interface
    // operational mode and  interface operational cause for all the  interfaces
    // in the modules.  This table contains one entry for each  64 interfaces in
    // an module.  This table provides efficient way of encoding  'ifIndex',
    // interface operational mode and interface operational cause, from the point 
    // of retrieval, by combining the values a set  of 64 interfaces in a single
    // MIB object.
    Cieifstatuslisttable CISCOIFEXTENSIONMIB_Cieifstatuslisttable

    // This table contains VL (Virtual Link) statistics for a capable interface. 
    // Objects defined in this table may be  applicable to physical interfaces
    // only.
    Cieifvlstatstable CISCOIFEXTENSIONMIB_Cieifvlstatstable

    // This table lists configuration data relating to ifIndex persistence.  This
    // table has a sparse dependent relationship on the ifTable, containing a row
    // for each ifEntry corresponding to an interface for which ifIndex
    // persistence is supported.
    Cieifindexpersistencetable CISCOIFEXTENSIONMIB_Cieifindexpersistencetable

    // A list of the interfaces that support the 802.1q custom Ethertype feature.
    Cieifdot1Qcustomethertypetable CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable

    // This table contains the interface utilization rates for inbound and
    // outbound traffic on an interface.
    Cieifutiltable CISCOIFEXTENSIONMIB_Cieifutiltable

    // This table contains the mappings of the ifIndex of an interface to its
    // corresponding dot1dBasePort value.
    Cieifdot1Dbasemappingtable CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable

    // This table contains objects for providing the 'ifName' to 'ifIndex'
    // mapping. This table contains one entry for each valid 'ifName' available in
    // the system. Upon the first request, the implementation of this table will
    // get all the available ifNames, and it will populate the entries in this
    // table, it maintains this ifNames in a cache for ~30 seconds.
    Cieifnamemappingtable CISCOIFEXTENSIONMIB_Cieifnamemappingtable
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetFilter() yfilter.YFilter { return cISCOIFEXTENSIONMIB.YFilter }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) SetFilter(yf yfilter.YFilter) { cISCOIFEXTENSIONMIB.YFilter = yf }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetGoName(yname string) string {
    if yname == "ciscoIfExtSystemConfig" { return "Ciscoifextsystemconfig" }
    if yname == "cieIfPacketStatsTable" { return "Cieifpacketstatstable" }
    if yname == "cieIfInterfaceTable" { return "Cieifinterfacetable" }
    if yname == "cieIfStatusListTable" { return "Cieifstatuslisttable" }
    if yname == "cieIfVlStatsTable" { return "Cieifvlstatstable" }
    if yname == "cieIfIndexPersistenceTable" { return "Cieifindexpersistencetable" }
    if yname == "cieIfDot1qCustomEtherTypeTable" { return "Cieifdot1Qcustomethertypetable" }
    if yname == "cieIfUtilTable" { return "Cieifutiltable" }
    if yname == "cieIfDot1dBaseMappingTable" { return "Cieifdot1Dbasemappingtable" }
    if yname == "cieIfNameMappingTable" { return "Cieifnamemappingtable" }
    return ""
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetSegmentPath() string {
    return "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB"
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoIfExtSystemConfig" {
        return &cISCOIFEXTENSIONMIB.Ciscoifextsystemconfig
    }
    if childYangName == "cieIfPacketStatsTable" {
        return &cISCOIFEXTENSIONMIB.Cieifpacketstatstable
    }
    if childYangName == "cieIfInterfaceTable" {
        return &cISCOIFEXTENSIONMIB.Cieifinterfacetable
    }
    if childYangName == "cieIfStatusListTable" {
        return &cISCOIFEXTENSIONMIB.Cieifstatuslisttable
    }
    if childYangName == "cieIfVlStatsTable" {
        return &cISCOIFEXTENSIONMIB.Cieifvlstatstable
    }
    if childYangName == "cieIfIndexPersistenceTable" {
        return &cISCOIFEXTENSIONMIB.Cieifindexpersistencetable
    }
    if childYangName == "cieIfDot1qCustomEtherTypeTable" {
        return &cISCOIFEXTENSIONMIB.Cieifdot1Qcustomethertypetable
    }
    if childYangName == "cieIfUtilTable" {
        return &cISCOIFEXTENSIONMIB.Cieifutiltable
    }
    if childYangName == "cieIfDot1dBaseMappingTable" {
        return &cISCOIFEXTENSIONMIB.Cieifdot1Dbasemappingtable
    }
    if childYangName == "cieIfNameMappingTable" {
        return &cISCOIFEXTENSIONMIB.Cieifnamemappingtable
    }
    return nil
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoIfExtSystemConfig"] = &cISCOIFEXTENSIONMIB.Ciscoifextsystemconfig
    children["cieIfPacketStatsTable"] = &cISCOIFEXTENSIONMIB.Cieifpacketstatstable
    children["cieIfInterfaceTable"] = &cISCOIFEXTENSIONMIB.Cieifinterfacetable
    children["cieIfStatusListTable"] = &cISCOIFEXTENSIONMIB.Cieifstatuslisttable
    children["cieIfVlStatsTable"] = &cISCOIFEXTENSIONMIB.Cieifvlstatstable
    children["cieIfIndexPersistenceTable"] = &cISCOIFEXTENSIONMIB.Cieifindexpersistencetable
    children["cieIfDot1qCustomEtherTypeTable"] = &cISCOIFEXTENSIONMIB.Cieifdot1Qcustomethertypetable
    children["cieIfUtilTable"] = &cISCOIFEXTENSIONMIB.Cieifutiltable
    children["cieIfDot1dBaseMappingTable"] = &cISCOIFEXTENSIONMIB.Cieifdot1Dbasemappingtable
    children["cieIfNameMappingTable"] = &cISCOIFEXTENSIONMIB.Cieifnamemappingtable
    return children
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetYangName() string { return "CISCO-IF-EXTENSION-MIB" }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) SetParent(parent types.Entity) { cISCOIFEXTENSIONMIB.parent = parent }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetParent() types.Entity { return cISCOIFEXTENSIONMIB.parent }

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig
type CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global system MTU in octets. This object specifies the MTU on all
    // interfaces. However, the value specified by cieIfMtu takes precedence for
    // an interface, which means that the interface's MTU uses the value specified
    // by cieIfMtu, if it is configured. The type is interface{} with range:
    // -2147483648..2147483647.
    Ciesystemmtu interface{}

    // Indicates whether cieLinkUp/cieLinkDown or standard mib-II defined
    // linkUp/Down or both, notifications should be generated for the interfaces
    // in the system.  'standard'  - only generate standard defined              
    // mib-II linkUp/linkDown notification               if
    // 'ifLinkUpDownTrapEnable' for                the interface is 'enabled'.
    // 'cisco'     - only generate cieLinkUp/cieLinkDown              
    // notifications for an interface if               the
    // 'ifLinkUpDownTrapEnable' for the               interface is 'enabled'.  If
    // both bits are selected then linkUp/linkDown and cieLinkUp/cieLinkDown are
    // both generated for an  interface if the 'ifLinkUpDownTrapEnable' for the
    // interface is 'enabled'. The type is map[string]bool.
    Cielinkupdownenable interface{}

    // Indicates whether to send the extra varbinds in addition to the varbinds
    // defined  in linkUp/linkDown notifications.  'standard'   - only send the
    // varbinds defined in                the standard linkUp/linkDown            
    // notification.   'additional' - send the extra varbinds in addition         
    // to the defined ones. 'other'      - any other config not covered by the
    // above.                This value is read-only. The type is
    // Ciestandardlinkupdownvarbinds.
    Ciestandardlinkupdownvarbinds interface{}

    // This object specifies whether ifIndex values persist across
    // reinitialization of the device.  ifIndex persistence means that the mapping
    // between the ifDescr object values and the ifIndex object values will be
    // retained across reboots.  Applications such as device inventory, billing,
    // and fault detection depend on the maintenance of the correspondence between
    // particular ifIndex values and their interfaces. During reboot or insertion
    // of a new card, the data to correlate the interfaces to the ifIndex may
    // become invalid in absence of ifIndex persistence feature.  ifIndex
    // persistence for an interface ensures ifIndex value for the interface will
    // remain the same after a system reboot. Hence, this feature allows users to
    // avoid the workarounds required for consistent interface identification
    // across reinitialization.  Due to change in syntax, this object is
    // deprecated by cieIfIndexGlobalPersistence. The type is bool.
    Cieifindexpersistence interface{}

    // This object specifies whether the system generates a
    // cieDelayedLinkUpDownNotif notification. The type is bool.
    Ciedelayedlinkupdownnotifenable interface{}

    // This object specifies the interval of time an interface's operational
    // status must remain stable following a transition before the system will
    // generate a cieDelayedLinkUpDownNotif. The type is interface{} with range:
    // 1..60. Units are minutes.
    Ciedelayedlinkupdownnotifdelay interface{}

    // This object specifies whether ifIndex values persist across
    // reinitialization of the device.  ifIndex persistence means that the mapping
    // between the ifDescr object values and the ifIndex object values will be
    // retained across reboots.  Applications such as device inventory, billing,
    // and fault detection depend on the maintenance of the correspondence between
    // particular ifIndex values and their interfaces. During reboot or insertion
    // of a new card, the data to correlate the interfaces to the ifIndex may
    // become invalid in absence of ifIndex persistence feature.  ifIndex
    // persistence for an interface ensures ifIndex value for the interface will
    // remain the same after a system reboot. Hence, this feature allows users to
    // avoid the workarounds required for consistent interface identification
    // across reinitialization.  The allowed values for this object are either
    // enable or disable. global value is not allowed. The type is
    // IfIndexPersistenceState.
    Cieifindexglobalpersistence interface{}

    // This object specifies whether standard mib-II defined linkUp/ linkDown,
    // extended linkUp/linkDown (with extra varbinds in addition to the varbinds
    // defined in linkUp/linkDown) or cieLinkUp/cieLinkDown notifications should
    // be generated for the interfaces in the system.  'standardLinkUp'     -
    // generate standard defined mib-II                         linkUp
    // notification if                         'ifLinkUpDownTrapEnable' for the   
    // interface is 'enabled'. 'standardLinkDown'   - generate standard defined
    // mib-II                         linkDown notification if                    
    // 'ifLinkUpDownTrapEnable' for the                         interface is
    // 'enabled'.   'additionalLinkUp'   - generate linkUp notification with      
    // additional varbinds if                         'ifLinkUpDownTrapEnable' for
    // the                         interface is 'enabled'.   'additionalLinkDown'
    // - generate linkDown notification with                        additional
    // varbinds if                         'ifLinkUpDownTrapEnable' for the       
    // interface is 'enabled'. 'ciscoLinkUp'        - generate cieLinkUp
    // notification                        if the 'ifLinkUpDownTrapEnable' for the
    // interface is 'enabled'. 'ciscoLinkDown'      - generate cieLinkDown
    // notification                        if the 'ifLinkUpDownTrapEnable' for the
    // interface is 'enabled'.  If multiple bits are set then multiple
    // notifications will be generated for an interface if the
    // 'ifLinkUpDownTrapEnable'  for the interface is 'enabled'. The type is
    // map[string]bool.
    Cielinkupdownconfig interface{}
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetFilter() yfilter.YFilter { return ciscoifextsystemconfig.YFilter }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) SetFilter(yf yfilter.YFilter) { ciscoifextsystemconfig.YFilter = yf }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetGoName(yname string) string {
    if yname == "cieSystemMtu" { return "Ciesystemmtu" }
    if yname == "cieLinkUpDownEnable" { return "Cielinkupdownenable" }
    if yname == "cieStandardLinkUpDownVarbinds" { return "Ciestandardlinkupdownvarbinds" }
    if yname == "cieIfIndexPersistence" { return "Cieifindexpersistence" }
    if yname == "cieDelayedLinkUpDownNotifEnable" { return "Ciedelayedlinkupdownnotifenable" }
    if yname == "cieDelayedLinkUpDownNotifDelay" { return "Ciedelayedlinkupdownnotifdelay" }
    if yname == "cieIfIndexGlobalPersistence" { return "Cieifindexglobalpersistence" }
    if yname == "cieLinkUpDownConfig" { return "Cielinkupdownconfig" }
    return ""
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetSegmentPath() string {
    return "ciscoIfExtSystemConfig"
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cieSystemMtu"] = ciscoifextsystemconfig.Ciesystemmtu
    leafs["cieLinkUpDownEnable"] = ciscoifextsystemconfig.Cielinkupdownenable
    leafs["cieStandardLinkUpDownVarbinds"] = ciscoifextsystemconfig.Ciestandardlinkupdownvarbinds
    leafs["cieIfIndexPersistence"] = ciscoifextsystemconfig.Cieifindexpersistence
    leafs["cieDelayedLinkUpDownNotifEnable"] = ciscoifextsystemconfig.Ciedelayedlinkupdownnotifenable
    leafs["cieDelayedLinkUpDownNotifDelay"] = ciscoifextsystemconfig.Ciedelayedlinkupdownnotifdelay
    leafs["cieIfIndexGlobalPersistence"] = ciscoifextsystemconfig.Cieifindexglobalpersistence
    leafs["cieLinkUpDownConfig"] = ciscoifextsystemconfig.Cielinkupdownconfig
    return leafs
}

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetYangName() string { return "ciscoIfExtSystemConfig" }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) SetParent(parent types.Entity) { ciscoifextsystemconfig.parent = parent }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetParent() types.Entity { return ciscoifextsystemconfig.parent }

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds represents                This value is read-only.
type CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds string

const (
    CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds_standard CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds = "standard"

    CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds_additional CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds = "additional"

    CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds_other CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig_Ciestandardlinkupdownvarbinds = "other"
)

// CISCOIFEXTENSIONMIB_Cieifpacketstatstable
// This  table contains interface packet
// statistics which are not available in 
// IF-MIB(RFC2863). 
// 
// As an example, some interfaces to which
// objects in this table are applicable are
// as follows :
// 
//         o Ethernet
//         o FastEthernet
//         o ATM
//         o BRI
//         o Sonet
//         o GigabitEthernet
// 
// Some objects defined in this table may be 
// applicable to physical interfaces only.
// As a result, this table may be sparse for
// some logical interfaces.
type CISCOIFEXTENSIONMIB_Cieifpacketstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry into the cieIfPacketStatsTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry.
    Cieifpacketstatsentry []CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetFilter() yfilter.YFilter { return cieifpacketstatstable.YFilter }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) SetFilter(yf yfilter.YFilter) { cieifpacketstatstable.YFilter = yf }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetGoName(yname string) string {
    if yname == "cieIfPacketStatsEntry" { return "Cieifpacketstatsentry" }
    return ""
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetSegmentPath() string {
    return "cieIfPacketStatsTable"
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfPacketStatsEntry" {
        for _, c := range cieifpacketstatstable.Cieifpacketstatsentry {
            if cieifpacketstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry{}
        cieifpacketstatstable.Cieifpacketstatsentry = append(cieifpacketstatstable.Cieifpacketstatsentry, child)
        return &cieifpacketstatstable.Cieifpacketstatsentry[len(cieifpacketstatstable.Cieifpacketstatsentry)-1]
    }
    return nil
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifpacketstatstable.Cieifpacketstatsentry {
        children[cieifpacketstatstable.Cieifpacketstatsentry[i].GetSegmentPath()] = &cieifpacketstatstable.Cieifpacketstatsentry[i]
    }
    return children
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetYangName() string { return "cieIfPacketStatsTable" }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) SetParent(parent types.Entity) { cieifpacketstatstable.parent = parent }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetParent() types.Entity { return cieifpacketstatstable.parent }

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry
// An entry into the cieIfPacketStatsTable.
type CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // input  packet was received.  Discontinuities in the value of this variable
    // can occur at re-initialization of the management system, and at other times
    // as  indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Cieiflastintime interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // output  packet was transmitted.  Discontinuities in the value of this
    // variable can occur at re-initialization of the management system, and at
    // other times as  indicated by the values of  cieIfPacketDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    Cieiflastouttime interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // output packet could not be successfully transmitted.  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Cieiflastouthangtime interface{}

    // The number of packets input on a particular physical interface which were
    // dropped as they were smaller than the minimum allowable  physical media
    // limit.  Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Cieifinruntserrs interface{}

    // The number of input packets on a particular physical interface which were
    // dropped as  they were larger than the ifMtu (largest  permitted  size of a
    // packet which can be  sent/received on an interface).  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Cieifingiantserrs interface{}

    // The number of input packets on a physical interface which were misaligned
    // or had framing errors. This happens when the  format of the incoming packet
    // on a physical interface is incorrect.  Discontinuities in the value of this
    // variable can occur at re-initialization of the management system, and at
    // other times as  indicated by the values of  cieIfPacketDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Cieifinframingerrs interface{}

    // The number of input packets which arrived on a particular physical
    // interface which  were too quick for the hardware to receive and hence the
    // receiver ran out of buffers.  Discontinuities in the value of this variable
    // can occur at re-initialization of the management system, and at other times
    // as  indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Cieifinoverrunerrs interface{}

    // The number of input packets which were simply ignored by this physical
    // interface due to  insufficient resources to handle the incoming packets. 
    // For example, this could indicate that the input receive buffers are not
    // available or that the receiver lost a packet.  Discontinuities in the value
    // of this variable can occur at re-initialization of the management system,
    // and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Cieifinignored interface{}

    // Number of input packets which were dropped because the receiver aborted. 
    // Examples of this could be when an abort sequence aborted the input frame or
    // when there is a collision in an ethernet segment.  Discontinuities in the
    // value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Cieifinaborterrs interface{}

    // The number of input packets which were dropped.  Some reasons why this
    // object could be  incremented are:  o  Input queue is full. o  Errors at the
    // receiver hardware     while receiving the packet.  Discontinuities in the
    // value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Cieifinputqueuedrops interface{}

    // This object indicates the  number of output packets dropped by the
    // interface even though no error had been detected to prevent them being
    // transmitted.   The packet could be dropped for many reasons, which could
    // range from the interface being down to errors in the format of the packet. 
    // Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Cieifoutputqueuedrops interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters suffered a discontinuity.   If no such
    // discontinuities have occurred  since the last re-initialization of the 
    // local management subsystem, then this  object contains a value of zero. The
    // type is interface{} with range: 0..4294967295.
    Cieifpacketdiscontinuitytime interface{}
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetFilter() yfilter.YFilter { return cieifpacketstatsentry.YFilter }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) SetFilter(yf yfilter.YFilter) { cieifpacketstatsentry.YFilter = yf }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfLastInTime" { return "Cieiflastintime" }
    if yname == "cieIfLastOutTime" { return "Cieiflastouttime" }
    if yname == "cieIfLastOutHangTime" { return "Cieiflastouthangtime" }
    if yname == "cieIfInRuntsErrs" { return "Cieifinruntserrs" }
    if yname == "cieIfInGiantsErrs" { return "Cieifingiantserrs" }
    if yname == "cieIfInFramingErrs" { return "Cieifinframingerrs" }
    if yname == "cieIfInOverrunErrs" { return "Cieifinoverrunerrs" }
    if yname == "cieIfInIgnored" { return "Cieifinignored" }
    if yname == "cieIfInAbortErrs" { return "Cieifinaborterrs" }
    if yname == "cieIfInputQueueDrops" { return "Cieifinputqueuedrops" }
    if yname == "cieIfOutputQueueDrops" { return "Cieifoutputqueuedrops" }
    if yname == "cieIfPacketDiscontinuityTime" { return "Cieifpacketdiscontinuitytime" }
    return ""
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetSegmentPath() string {
    return "cieIfPacketStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifpacketstatsentry.Ifindex) + "']"
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifpacketstatsentry.Ifindex
    leafs["cieIfLastInTime"] = cieifpacketstatsentry.Cieiflastintime
    leafs["cieIfLastOutTime"] = cieifpacketstatsentry.Cieiflastouttime
    leafs["cieIfLastOutHangTime"] = cieifpacketstatsentry.Cieiflastouthangtime
    leafs["cieIfInRuntsErrs"] = cieifpacketstatsentry.Cieifinruntserrs
    leafs["cieIfInGiantsErrs"] = cieifpacketstatsentry.Cieifingiantserrs
    leafs["cieIfInFramingErrs"] = cieifpacketstatsentry.Cieifinframingerrs
    leafs["cieIfInOverrunErrs"] = cieifpacketstatsentry.Cieifinoverrunerrs
    leafs["cieIfInIgnored"] = cieifpacketstatsentry.Cieifinignored
    leafs["cieIfInAbortErrs"] = cieifpacketstatsentry.Cieifinaborterrs
    leafs["cieIfInputQueueDrops"] = cieifpacketstatsentry.Cieifinputqueuedrops
    leafs["cieIfOutputQueueDrops"] = cieifpacketstatsentry.Cieifoutputqueuedrops
    leafs["cieIfPacketDiscontinuityTime"] = cieifpacketstatsentry.Cieifpacketdiscontinuitytime
    return leafs
}

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetYangName() string { return "cieIfPacketStatsEntry" }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) SetParent(parent types.Entity) { cieifpacketstatsentry.parent = parent }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetParent() types.Entity { return cieifpacketstatsentry.parent }

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetParentYangName() string { return "cieIfPacketStatsTable" }

// CISCOIFEXTENSIONMIB_Cieifinterfacetable
// This  table contains objects which provide
// more information about interface  
// properties not available in IF-MIB
// (RFC 2863).
// 
// Some objects defined in this table may be
// applicable to physical interfaces only.
// As a result, this table may be sparse for
// logical interfaces.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry into the cieIfInterfaceTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry.
    Cieifinterfaceentry []CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetFilter() yfilter.YFilter { return cieifinterfacetable.YFilter }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) SetFilter(yf yfilter.YFilter) { cieifinterfacetable.YFilter = yf }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetGoName(yname string) string {
    if yname == "cieIfInterfaceEntry" { return "Cieifinterfaceentry" }
    return ""
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetSegmentPath() string {
    return "cieIfInterfaceTable"
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfInterfaceEntry" {
        for _, c := range cieifinterfacetable.Cieifinterfaceentry {
            if cieifinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry{}
        cieifinterfacetable.Cieifinterfaceentry = append(cieifinterfacetable.Cieifinterfaceentry, child)
        return &cieifinterfacetable.Cieifinterfaceentry[len(cieifinterfacetable.Cieifinterfaceentry)-1]
    }
    return nil
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifinterfacetable.Cieifinterfaceentry {
        children[cieifinterfacetable.Cieifinterfaceentry[i].GetSegmentPath()] = &cieifinterfacetable.Cieifinterfaceentry[i]
    }
    return children
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetYangName() string { return "cieIfInterfaceTable" }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) SetParent(parent types.Entity) { cieifinterfacetable.parent = parent }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetParent() types.Entity { return cieifinterfacetable.parent }

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry
// An entry into the cieIfInterfaceTable.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of times the interface was internally reset and brought up. 
    // Some of the actions which can cause this counter to increment are :  o 
    // Bringing an interface up using the     interface CLI command.  o  Clearing
    // the interface with the exec    CLI command.  o  Bringing the interface up
    // via SNMP.  Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfInterfaceDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Cieifresetcount interface{}

    // A keepalive is a small, layer-2 message that is transmitted by a network
    // device  to let directly-connected network devices know of its presence. 
    // This object returns 'true' if keepalives are enabled on this interface. If
    // keepalives are not enabled, 'false' is returned.  Setting this object to
    // TRUE or FALSE enables or disables (respectively) keepalive on this 
    // interface. The type is bool.
    Cieifkeepaliveenabled interface{}

    // This object displays a human-readable textual string which describes the 
    // cause of the last state change of the  interface.  Examples of the values
    // this object can take are:  o  'Lost Carrier' o  'administratively down' o 
    // 'up' o  'down'. The type is string.
    Cieifstatechangereason interface{}

    // Number of times interface saw the carrier signal transition.  For example,
    // if a T1 line is unplugged,  then framer will detect the loss of signal 
    // (LOS) on the line  and will count it as a transition.  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfInterfaceDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Cieifcarriertransitioncount interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters  suffered  a discontinuity.   If no such
    // discontinuities have occurred  since the last re-initialization of the 
    // local management subsystem, then this  object contains a value of zero. The
    // type is interface{} with range: 0..4294967295.
    Cieifinterfacediscontinuitytime interface{}

    // The DHCP mode configured by the administrator. If 'true' the DHCP is
    // enabled. In which case an IP address is requested in DHCP. This is in
    // addition to any that are  configured by the administrator in
    // 'ciiIPAddressTable' or 'ciiIPIfAddressTable' in CISCO-IP-IF-MIB. If 'false'
    // the DHCP is disabled. In which case all IP addresses are configured by the
    // administrator in 'ciiIPAddressTable' or  'ciiIPIfAddressTable'. For
    // interfaces, for which DHCP cannot be or is not supported, then this object
    // has the value 'false'. The type is bool.
    Cieifdhcpmode interface{}

    // The MTU configured by the administrator. This object is exactly same as
    // 'ifMtu' in  ifTable from IF-MIB for the same ifIndex value , except that it
    // is configurable by the administrator. For more description of this object
    // refer to 'ifMtu' in IF-MIB. The type is interface{} with range:
    // 40..2147483647.
    Cieifmtu interface{}

    // The ContextName denotes the interface 'context' and is used to logically
    // separate the MIB management. RFC 2571 and RFC 2737 describe this approach.
    // When the agent supports a different SNMP  context, as detailed in RFC 2571
    // and  RFC 2737, for different interfaces, then the value of this object
    // specifies the context name used for this interface. The type is string with
    // length: 0..32.
    Cieifcontextname interface{}

    // This object represents the detailed operational cause reason for the
    // current  operational state of the interface.  The current operational state
    // of the interface  is given by the 'ifOperStatus' defined  in IF-MIB.   The
    // corresponding instance of  'cieIfOperStatusCauseDescr' must be used to  get
    // the information about the operational  cause value mentioned in this
    // object.  For interfaces whose 'ifOperStatus' is 'down'  the objects
    // 'cieIfOperStatusCause' and  'cieIfOperStatusCauseDescr' together provides 
    // the information about the operational cause  reason and the description of
    // the cause.   The value of this object will be 'none' for all the
    // 'ifOperStatus' values except for  'down'. Its value will be one status
    // cause  defined in the 'IfOperStatusReason' textual  convention if
    // 'ifOperStatus' is 'down'.   The value of this object will be 'other'  if
    // the operational status cause is not one  defined in 'IfOperStatusReason'.
    // The type is IfOperStatusReason.
    Cieifoperstatuscause interface{}

    // The description for the cause of current operational state of the
    // interface, given  by the object 'cieIfOperStatusCause'.  For an interface
    // whose 'ifOperStatus' is not 'down' the value of this object will be 
    // 'none'. The type is string.
    Cieifoperstatuscausedescr interface{}

    // An estimate of the interface's current receive bandwidth in bits per
    // second.  This object is provided for interface with asymmetric interface
    // speeds like ADSL and should be used in conjunction with ifSpeed object. 
    // For interfaces which do not vary in bandwidth or for those where no
    // accurate estimation can be made, this object should contain the nominal
    // bandwidth. If the bandwidth of the interface is greater than the maximum
    // value reportable by this object then this object should report its maximum
    // value (4,294,967,295) and ifHighSpeed must be used to report the interace's
    // speed.  For a sub-layer which has no concept of bandwidth, this object
    // should be zero. The type is interface{} with range: 0..4294967295.
    Cieifspeedreceive interface{}

    // An estimate of the interface's current receive bandwidth in units of
    // 1,000,000 bits per second.  If this object reports a value of `n' then the
    // speed of the interface is somewhere in the range of `n-500,000' to
    // `n+499,999'.  For interfaces which do not vary in bandwidth or for those
    // where no accurate estimation can be made, this object should contain the
    // nominal bandwidth.  For a sub-layer which has no concept of bandwidth, this
    // object should be zero. The type is interface{} with range: 0..4294967295.
    Cieifhighspeedreceive interface{}

    // This data type is used to model an administratively assigned name of the
    // current owner of the interface resource. This  information is taken from
    // the NVT ASCII character set.  It is  suggested that this name contain one
    // or more of the following:  SnmpEngineID, IP address, management station
    // name, network  manager's name, location, or phone number. SNMP access
    // control is articulated entirely in terms of the  contents of MIB views;
    // access to a particular SNMP object  instance depends only upon its presence
    // or absence in a  particular MIB view and never upon its value or the value
    // of  related object instances. Thus, this object affords resolution of
    // resource contention  only among cooperating managers; this object realizes
    // no access control function with respect to uncooperative parties. The type
    // is string with length: 0..80.
    Cieifowner interface{}

    // This object indicates the current configuration of interface sharing on the
    // given interface.  'notApplicable' - the interface sharing configuration on 
    // this interface is not applicable.  'ownerDedicated' - the interface is in
    // the dedicated mode             to the binding physical interface.
    // 'ownerShared' - the interface is shared amongst virtual switches         
    // and this interface physically belongs to a its           virtual switch.  
    // 'sharedOnly' - the interface is in purely shared mode. The type is
    // Cieifsharedconfig.
    Cieifsharedconfig interface{}

    // This object specifies the current speed group configuration on the given
    // interface.  'notApplicable' - the interface speed group configuration on   
    // this interface is not applicable. It is a              read-only value.
    // '10G' - the interface speed group configuration on             this
    // interface as 10G. '1G-2G-4G-8G' - the interface speed group configuration  
    // on this interface as 1G-2G-4G-8G. '2G-4G-8G-16G' - the interface speed
    // group configuration              on this interface as 2G-4G-8G-16G. The
    // type is Cieifspeedgroupconfig.
    Cieifspeedgroupconfig interface{}

    // This object specifies the current transceiver frequency configuration on
    // the given interface.  'notApplicable' - the interface transceiver frequency
    // 				  configuration on this interface  				  is not applicable. It is a
    // read-only value. 'FibreChannel' - the interface transceiver frequency 				
    // configuration on this interface as                   Fibre Channel.
    // 'Ethernet'	  -  the interface transceiver frequency on 				 this interface
    // as Ethernet. The type is Cieiftransceiverfrequencyconfig.
    Cieiftransceiverfrequencyconfig interface{}

    // This object specifies the current switchport fill pattern configuration on
    // the given interface.  'arbff8G' - the inter frame gap fill pattern is
    // 			ARBFF for 8G speed. 'idle8G' - the inter frame gap fill pattern is 		  
    // IDLE for 8G speed. The type is Cieiffillpatternconfig.
    Cieiffillpatternconfig interface{}

    // This object specifies the current switchport biterrors configuration on the
    // given interface.  If 'true(1)' the ignore bit errors is enabled.In which
    // case the interface ignores bit errors. If 'false(2)' the ignore bit errors
    // is disabled. In which  case the interface acts on the bit errors.  For
    // interfaces, for which bit errors  is not supported, then this object has
    // the value 'true(1)'. The type is bool.
    Cieifignorebiterrorsconfig interface{}

    // This object specifies the current interrupt threshold configuration on the
    // given interface.  'If 'true(1)' the ignore interrupt thresholds is enabled.
    // In which case the interface ignores interrupt thresholds. If 'false(2)' the
    // ignore interrupt thresholds is disabled. In which case the interface acts
    // on the interrupt  thresholds.  For interfaces, for which interrupt
    // thresholds  is not supported, then this object has the  value 'true(1)'.
    // The type is bool.
    Cieifignoreinterruptthresholdconfig interface{}
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetFilter() yfilter.YFilter { return cieifinterfaceentry.YFilter }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) SetFilter(yf yfilter.YFilter) { cieifinterfaceentry.YFilter = yf }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfResetCount" { return "Cieifresetcount" }
    if yname == "cieIfKeepAliveEnabled" { return "Cieifkeepaliveenabled" }
    if yname == "cieIfStateChangeReason" { return "Cieifstatechangereason" }
    if yname == "cieIfCarrierTransitionCount" { return "Cieifcarriertransitioncount" }
    if yname == "cieIfInterfaceDiscontinuityTime" { return "Cieifinterfacediscontinuitytime" }
    if yname == "cieIfDhcpMode" { return "Cieifdhcpmode" }
    if yname == "cieIfMtu" { return "Cieifmtu" }
    if yname == "cieIfContextName" { return "Cieifcontextname" }
    if yname == "cieIfOperStatusCause" { return "Cieifoperstatuscause" }
    if yname == "cieIfOperStatusCauseDescr" { return "Cieifoperstatuscausedescr" }
    if yname == "cieIfSpeedReceive" { return "Cieifspeedreceive" }
    if yname == "cieIfHighSpeedReceive" { return "Cieifhighspeedreceive" }
    if yname == "cieIfOwner" { return "Cieifowner" }
    if yname == "cieIfSharedConfig" { return "Cieifsharedconfig" }
    if yname == "cieIfSpeedGroupConfig" { return "Cieifspeedgroupconfig" }
    if yname == "cieIfTransceiverFrequencyConfig" { return "Cieiftransceiverfrequencyconfig" }
    if yname == "cieIfFillPatternConfig" { return "Cieiffillpatternconfig" }
    if yname == "cieIfIgnoreBitErrorsConfig" { return "Cieifignorebiterrorsconfig" }
    if yname == "cieIfIgnoreInterruptThresholdConfig" { return "Cieifignoreinterruptthresholdconfig" }
    return ""
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetSegmentPath() string {
    return "cieIfInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifinterfaceentry.Ifindex) + "']"
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifinterfaceentry.Ifindex
    leafs["cieIfResetCount"] = cieifinterfaceentry.Cieifresetcount
    leafs["cieIfKeepAliveEnabled"] = cieifinterfaceentry.Cieifkeepaliveenabled
    leafs["cieIfStateChangeReason"] = cieifinterfaceentry.Cieifstatechangereason
    leafs["cieIfCarrierTransitionCount"] = cieifinterfaceentry.Cieifcarriertransitioncount
    leafs["cieIfInterfaceDiscontinuityTime"] = cieifinterfaceentry.Cieifinterfacediscontinuitytime
    leafs["cieIfDhcpMode"] = cieifinterfaceentry.Cieifdhcpmode
    leafs["cieIfMtu"] = cieifinterfaceentry.Cieifmtu
    leafs["cieIfContextName"] = cieifinterfaceentry.Cieifcontextname
    leafs["cieIfOperStatusCause"] = cieifinterfaceentry.Cieifoperstatuscause
    leafs["cieIfOperStatusCauseDescr"] = cieifinterfaceentry.Cieifoperstatuscausedescr
    leafs["cieIfSpeedReceive"] = cieifinterfaceentry.Cieifspeedreceive
    leafs["cieIfHighSpeedReceive"] = cieifinterfaceentry.Cieifhighspeedreceive
    leafs["cieIfOwner"] = cieifinterfaceentry.Cieifowner
    leafs["cieIfSharedConfig"] = cieifinterfaceentry.Cieifsharedconfig
    leafs["cieIfSpeedGroupConfig"] = cieifinterfaceentry.Cieifspeedgroupconfig
    leafs["cieIfTransceiverFrequencyConfig"] = cieifinterfaceentry.Cieiftransceiverfrequencyconfig
    leafs["cieIfFillPatternConfig"] = cieifinterfaceentry.Cieiffillpatternconfig
    leafs["cieIfIgnoreBitErrorsConfig"] = cieifinterfaceentry.Cieifignorebiterrorsconfig
    leafs["cieIfIgnoreInterruptThresholdConfig"] = cieifinterfaceentry.Cieifignoreinterruptthresholdconfig
    return leafs
}

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetYangName() string { return "cieIfInterfaceEntry" }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) SetParent(parent types.Entity) { cieifinterfaceentry.parent = parent }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetParent() types.Entity { return cieifinterfaceentry.parent }

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetParentYangName() string { return "cieIfInterfaceTable" }

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig represents 		   IDLE for 8G speed.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig string

const (
    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig_arbff8G CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig = "arbff8G"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig_idle8G CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiffillpatternconfig = "idle8G"
)

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig represents 'sharedOnly' - the interface is in purely shared mode.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig string

const (
    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig_notApplicable CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig = "notApplicable"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig_ownerDedicated CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig = "ownerDedicated"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig_ownerShared CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig = "ownerShared"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig_sharedOnly CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifsharedconfig = "sharedOnly"
)

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig represents             on this interface as 2G-4G-8G-16G.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig string

const (
    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig_notApplicable CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig = "notApplicable"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig_tenG CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig = "tenG"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig_oneTwoFourEightG CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig = "oneTwoFourEightG"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig_twoFourEightSixteenG CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieifspeedgroupconfig = "twoFourEightSixteenG"
)

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig represents 				 this interface as Ethernet.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig string

const (
    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig_notApplicable CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig = "notApplicable"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig_fibreChannel CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig = "fibreChannel"

    CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig_ethernet CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry_Cieiftransceiverfrequencyconfig = "ethernet"
)

// CISCOIFEXTENSIONMIB_Cieifstatuslisttable
// This table contains objects for providing
// the 'ifIndex', interface operational mode and 
// interface operational cause for all the 
// interfaces in the modules.
// 
// This table contains one entry for each 
// 64 interfaces in an module.
// 
// This table provides efficient way of encoding 
// 'ifIndex', interface operational mode and
// interface operational cause, from the point 
// of retrieval, by combining the values a set 
// of 64 interfaces in a single MIB object.
type CISCOIFEXTENSIONMIB_Cieifstatuslisttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents the 'ifIndex', interface operational mode and
    // interface  operational cause for a set of 64 interfaces  in a module. The
    // type is slice of
    // CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry.
    Cieifstatuslistentry []CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetFilter() yfilter.YFilter { return cieifstatuslisttable.YFilter }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) SetFilter(yf yfilter.YFilter) { cieifstatuslisttable.YFilter = yf }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetGoName(yname string) string {
    if yname == "cieIfStatusListEntry" { return "Cieifstatuslistentry" }
    return ""
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetSegmentPath() string {
    return "cieIfStatusListTable"
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfStatusListEntry" {
        for _, c := range cieifstatuslisttable.Cieifstatuslistentry {
            if cieifstatuslisttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry{}
        cieifstatuslisttable.Cieifstatuslistentry = append(cieifstatuslisttable.Cieifstatuslistentry, child)
        return &cieifstatuslisttable.Cieifstatuslistentry[len(cieifstatuslisttable.Cieifstatuslistentry)-1]
    }
    return nil
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifstatuslisttable.Cieifstatuslistentry {
        children[cieifstatuslisttable.Cieifstatuslistentry[i].GetSegmentPath()] = &cieifstatuslisttable.Cieifstatuslistentry[i]
    }
    return children
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetYangName() string { return "cieIfStatusListTable" }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) SetParent(parent types.Entity) { cieifstatuslisttable.parent = parent }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetParent() types.Entity { return cieifstatuslisttable.parent }

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry
// Each entry represents the 'ifIndex',
// interface operational mode and interface 
// operational cause for a set of 64 interfaces 
// in a module.
type CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. An arbitrary integer value, greater than zero,
    // which identifies a list of 64 interfaces within a module. The type is
    // interface{} with range: 1..33554432.
    Cieifstatuslistindex interface{}

    // This object represents the 'ifIndex' for a set of 64 interfaces in the
    // module. The type is string with length: 0..256.
    Cieinterfacesindex interface{}

    // This object represents the operational mode for a set of 64 interfaces in
    // the module. The type is string with length: 0..64.
    Cieinterfacesopermode interface{}

    // This object represents the operational status cause for a set of 64
    // interfaces in the  module. The type is string with length: 0..128.
    Cieinterfacesopercause interface{}

    // This object indicates the status for a set of 64 interfaces in a module
    // regarding whether or not each interface is  administratively assigned a
    // name of the current owner of the  interface resource as per cieIfOwner. The
    // type is string with length: 0..8.
    Cieinterfaceownershipbitmap interface{}
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetFilter() yfilter.YFilter { return cieifstatuslistentry.YFilter }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) SetFilter(yf yfilter.YFilter) { cieifstatuslistentry.YFilter = yf }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cieIfStatusListIndex" { return "Cieifstatuslistindex" }
    if yname == "cieInterfacesIndex" { return "Cieinterfacesindex" }
    if yname == "cieInterfacesOperMode" { return "Cieinterfacesopermode" }
    if yname == "cieInterfacesOperCause" { return "Cieinterfacesopercause" }
    if yname == "cieInterfaceOwnershipBitmap" { return "Cieinterfaceownershipbitmap" }
    return ""
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetSegmentPath() string {
    return "cieIfStatusListEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cieifstatuslistentry.Entphysicalindex) + "']" + "[cieIfStatusListIndex='" + fmt.Sprintf("%v", cieifstatuslistentry.Cieifstatuslistindex) + "']"
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cieifstatuslistentry.Entphysicalindex
    leafs["cieIfStatusListIndex"] = cieifstatuslistentry.Cieifstatuslistindex
    leafs["cieInterfacesIndex"] = cieifstatuslistentry.Cieinterfacesindex
    leafs["cieInterfacesOperMode"] = cieifstatuslistentry.Cieinterfacesopermode
    leafs["cieInterfacesOperCause"] = cieifstatuslistentry.Cieinterfacesopercause
    leafs["cieInterfaceOwnershipBitmap"] = cieifstatuslistentry.Cieinterfaceownershipbitmap
    return leafs
}

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetYangName() string { return "cieIfStatusListEntry" }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) SetParent(parent types.Entity) { cieifstatuslistentry.parent = parent }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetParent() types.Entity { return cieifstatuslistentry.parent }

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetParentYangName() string { return "cieIfStatusListTable" }

// CISCOIFEXTENSIONMIB_Cieifvlstatstable
// This table contains VL (Virtual Link) statistics
// for a capable interface.
// 
// Objects defined in this table may be 
// applicable to physical interfaces only.
type CISCOIFEXTENSIONMIB_Cieifvlstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each row contains managed objects for Virtual Link statistics on interface
    // capable of  providing this information. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry.
    Cieifvlstatsentry []CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetFilter() yfilter.YFilter { return cieifvlstatstable.YFilter }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) SetFilter(yf yfilter.YFilter) { cieifvlstatstable.YFilter = yf }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetGoName(yname string) string {
    if yname == "cieIfVlStatsEntry" { return "Cieifvlstatsentry" }
    return ""
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetSegmentPath() string {
    return "cieIfVlStatsTable"
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfVlStatsEntry" {
        for _, c := range cieifvlstatstable.Cieifvlstatsentry {
            if cieifvlstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry{}
        cieifvlstatstable.Cieifvlstatsentry = append(cieifvlstatstable.Cieifvlstatsentry, child)
        return &cieifvlstatstable.Cieifvlstatsentry[len(cieifvlstatstable.Cieifvlstatsentry)-1]
    }
    return nil
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifvlstatstable.Cieifvlstatsentry {
        children[cieifvlstatstable.Cieifvlstatsentry[i].GetSegmentPath()] = &cieifvlstatstable.Cieifvlstatsentry[i]
    }
    return children
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetYangName() string { return "cieIfVlStatsTable" }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) SetParent(parent types.Entity) { cieifvlstatstable.parent = parent }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetParent() types.Entity { return cieifvlstatstable.parent }

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry
// Each row contains managed objects for
// Virtual Link statistics on interface capable of 
// providing this information.
type CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object indicates the number of input packets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifnodropvlinpkts interface{}

    // This object indicates the number of input octets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifnodropvlinoctets interface{}

    // This object indicates the number of output packets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifnodropvloutpkts interface{}

    // This object indicates the number of output octets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifnodropvloutoctets interface{}

    // This object indicates the number of input packets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifdropvlinpkts interface{}

    // This object indicates the number of input octets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifdropvlinoctets interface{}

    // This object indicates the number of output packets on all Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifdropvloutpkts interface{}

    // This object indicates the number of output octets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    Cieifdropvloutoctets interface{}
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetFilter() yfilter.YFilter { return cieifvlstatsentry.YFilter }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) SetFilter(yf yfilter.YFilter) { cieifvlstatsentry.YFilter = yf }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfNoDropVlInPkts" { return "Cieifnodropvlinpkts" }
    if yname == "cieIfNoDropVlInOctets" { return "Cieifnodropvlinoctets" }
    if yname == "cieIfNoDropVlOutPkts" { return "Cieifnodropvloutpkts" }
    if yname == "cieIfNoDropVlOutOctets" { return "Cieifnodropvloutoctets" }
    if yname == "cieIfDropVlInPkts" { return "Cieifdropvlinpkts" }
    if yname == "cieIfDropVlInOctets" { return "Cieifdropvlinoctets" }
    if yname == "cieIfDropVlOutPkts" { return "Cieifdropvloutpkts" }
    if yname == "cieIfDropVlOutOctets" { return "Cieifdropvloutoctets" }
    return ""
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetSegmentPath() string {
    return "cieIfVlStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifvlstatsentry.Ifindex) + "']"
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifvlstatsentry.Ifindex
    leafs["cieIfNoDropVlInPkts"] = cieifvlstatsentry.Cieifnodropvlinpkts
    leafs["cieIfNoDropVlInOctets"] = cieifvlstatsentry.Cieifnodropvlinoctets
    leafs["cieIfNoDropVlOutPkts"] = cieifvlstatsentry.Cieifnodropvloutpkts
    leafs["cieIfNoDropVlOutOctets"] = cieifvlstatsentry.Cieifnodropvloutoctets
    leafs["cieIfDropVlInPkts"] = cieifvlstatsentry.Cieifdropvlinpkts
    leafs["cieIfDropVlInOctets"] = cieifvlstatsentry.Cieifdropvlinoctets
    leafs["cieIfDropVlOutPkts"] = cieifvlstatsentry.Cieifdropvloutpkts
    leafs["cieIfDropVlOutOctets"] = cieifvlstatsentry.Cieifdropvloutoctets
    return leafs
}

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetYangName() string { return "cieIfVlStatsEntry" }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) SetParent(parent types.Entity) { cieifvlstatsentry.parent = parent }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetParent() types.Entity { return cieifvlstatsentry.parent }

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetParentYangName() string { return "cieIfVlStatsTable" }

// CISCOIFEXTENSIONMIB_Cieifindexpersistencetable
// This table lists configuration data relating to ifIndex
// persistence.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each ifEntry corresponding to an interface
// for which ifIndex persistence is supported.
type CISCOIFEXTENSIONMIB_Cieifindexpersistencetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents ifindex persistence configuration for an interface
    // specified by ifIndex. Whenever an interface which supports ifindex
    // persistence is created/destroyed in the ifTable, the corresponding ifindex
    // persistence entry is created/destroyed respectively. Some of the interfaces
    // may not support ifindex persistence, for example, a dynamic interface, such
    // as a PPP connection or a IP subscriber interface. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry.
    Cieifindexpersistenceentry []CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetFilter() yfilter.YFilter { return cieifindexpersistencetable.YFilter }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) SetFilter(yf yfilter.YFilter) { cieifindexpersistencetable.YFilter = yf }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetGoName(yname string) string {
    if yname == "cieIfIndexPersistenceEntry" { return "Cieifindexpersistenceentry" }
    return ""
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetSegmentPath() string {
    return "cieIfIndexPersistenceTable"
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfIndexPersistenceEntry" {
        for _, c := range cieifindexpersistencetable.Cieifindexpersistenceentry {
            if cieifindexpersistencetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry{}
        cieifindexpersistencetable.Cieifindexpersistenceentry = append(cieifindexpersistencetable.Cieifindexpersistenceentry, child)
        return &cieifindexpersistencetable.Cieifindexpersistenceentry[len(cieifindexpersistencetable.Cieifindexpersistenceentry)-1]
    }
    return nil
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifindexpersistencetable.Cieifindexpersistenceentry {
        children[cieifindexpersistencetable.Cieifindexpersistenceentry[i].GetSegmentPath()] = &cieifindexpersistencetable.Cieifindexpersistenceentry[i]
    }
    return children
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetYangName() string { return "cieIfIndexPersistenceTable" }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) SetParent(parent types.Entity) { cieifindexpersistencetable.parent = parent }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetParent() types.Entity { return cieifindexpersistencetable.parent }

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry
// Each entry represents ifindex persistence configuration for an
// interface specified by ifIndex. Whenever an interface which
// supports ifindex persistence is created/destroyed in the
// ifTable, the corresponding ifindex persistence entry is
// created/destroyed respectively. Some of the interfaces may not
// support ifindex persistence, for example, a dynamic interface,
// such as a PPP connection or a IP subscriber interface.
type CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object specifies whether the interface's ifIndex value persist across
    // reinitialization.  Due to change in syntax, this object is deprecated by
    // cieIfIndexPersistenceControl. The type is bool.
    Cieifindexpersistenceenabled interface{}

    // This object specifies whether the interface's ifIndex value persist across
    // reinitialization. In global state, the interface uses the global setting
    // data for persistence i.e. cieIfIndexGlobalPersistence. The type is
    // IfIndexPersistenceState.
    Cieifindexpersistencecontrol interface{}
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetFilter() yfilter.YFilter { return cieifindexpersistenceentry.YFilter }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) SetFilter(yf yfilter.YFilter) { cieifindexpersistenceentry.YFilter = yf }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfIndexPersistenceEnabled" { return "Cieifindexpersistenceenabled" }
    if yname == "cieIfIndexPersistenceControl" { return "Cieifindexpersistencecontrol" }
    return ""
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetSegmentPath() string {
    return "cieIfIndexPersistenceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifindexpersistenceentry.Ifindex) + "']"
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifindexpersistenceentry.Ifindex
    leafs["cieIfIndexPersistenceEnabled"] = cieifindexpersistenceentry.Cieifindexpersistenceenabled
    leafs["cieIfIndexPersistenceControl"] = cieifindexpersistenceentry.Cieifindexpersistencecontrol
    return leafs
}

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetYangName() string { return "cieIfIndexPersistenceEntry" }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) SetParent(parent types.Entity) { cieifindexpersistenceentry.parent = parent }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetParent() types.Entity { return cieifindexpersistenceentry.parent }

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetParentYangName() string { return "cieIfIndexPersistenceTable" }

// CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable
// A list of the interfaces that support
// the 802.1q custom Ethertype feature.
type CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the custom EtherType information for the interface. 
    // Only interfaces with custom 802.1q ethertype control are listed in the 
    // table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry.
    Cieifdot1Qcustomethertypeentry []CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetFilter() yfilter.YFilter { return cieifdot1Qcustomethertypetable.YFilter }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) SetFilter(yf yfilter.YFilter) { cieifdot1Qcustomethertypetable.YFilter = yf }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetGoName(yname string) string {
    if yname == "cieIfDot1qCustomEtherTypeEntry" { return "Cieifdot1Qcustomethertypeentry" }
    return ""
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetSegmentPath() string {
    return "cieIfDot1qCustomEtherTypeTable"
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfDot1qCustomEtherTypeEntry" {
        for _, c := range cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry {
            if cieifdot1Qcustomethertypetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry{}
        cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry = append(cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry, child)
        return &cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry[len(cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry)-1]
    }
    return nil
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry {
        children[cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry[i].GetSegmentPath()] = &cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry[i]
    }
    return children
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetYangName() string { return "cieIfDot1qCustomEtherTypeTable" }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) SetParent(parent types.Entity) { cieifdot1Qcustomethertypetable.parent = parent }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetParent() types.Entity { return cieifdot1Qcustomethertypetable.parent }

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry
// An entry containing the custom EtherType
// information for the interface.
// 
// Only interfaces with custom 802.1q
// ethertype control are listed in the 
// table.
type CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The Dot1qEtherType allow administrator to select a non-standard (other than
    // 0x8100) 2-byte ethertype for the interface to  interoperate with third
    // party vendor's system that do not use the standard 0x8100 ethertype to
    // identify 802.1q-tagged frames.  The current administrative value of the 
    // 802.1q ethertype for the interface.  The administrative 802.1q ethertype
    // value may  differ from the operational 802.1q ethertype value.  On some
    // platforms, 802.1q ethertype may be assigned per group rather than per port.
    // If multiple ports belong to a port group, the 802.1q ethertype assigned to
    // any of the ports in such group will apply to all ports in the same group. 
    // To configure non-standard dot1q ethertype is only recommended when the
    // Cisco device is connected to any third party vendor device. Also be advised
    // that the custom ethertype value needs to be changed in the whole cloud of 
    // Cisco device with the same custom ethertype  value if the third party
    // device are separated  by number of Cisco device in the middle. The type is
    // interface{} with range: 0..65535.
    Cieifdot1Qcustomadminethertype interface{}

    // The current operational value of the 802.1q ethertype for the interface.
    // The type is interface{} with range: 0..65535.
    Cieifdot1Qcustomoperethertype interface{}
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetFilter() yfilter.YFilter { return cieifdot1Qcustomethertypeentry.YFilter }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) SetFilter(yf yfilter.YFilter) { cieifdot1Qcustomethertypeentry.YFilter = yf }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfDot1qCustomAdminEtherType" { return "Cieifdot1Qcustomadminethertype" }
    if yname == "cieIfDot1qCustomOperEtherType" { return "Cieifdot1Qcustomoperethertype" }
    return ""
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetSegmentPath() string {
    return "cieIfDot1qCustomEtherTypeEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifdot1Qcustomethertypeentry.Ifindex) + "']"
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifdot1Qcustomethertypeentry.Ifindex
    leafs["cieIfDot1qCustomAdminEtherType"] = cieifdot1Qcustomethertypeentry.Cieifdot1Qcustomadminethertype
    leafs["cieIfDot1qCustomOperEtherType"] = cieifdot1Qcustomethertypeentry.Cieifdot1Qcustomoperethertype
    return leafs
}

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetYangName() string { return "cieIfDot1qCustomEtherTypeEntry" }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) SetParent(parent types.Entity) { cieifdot1Qcustomethertypeentry.parent = parent }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetParent() types.Entity { return cieifdot1Qcustomethertypeentry.parent }

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetParentYangName() string { return "cieIfDot1qCustomEtherTypeTable" }

// CISCOIFEXTENSIONMIB_Cieifutiltable
// This table contains the interface utilization
// rates for inbound and outbound traffic on an
// interface.
type CISCOIFEXTENSIONMIB_Cieifutiltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing utilization rates for the interface.  Every interface
    // for which the  inbound and  outbound traffic information is available has a
    // corresponding entry in this table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry.
    Cieifutilentry []CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetFilter() yfilter.YFilter { return cieifutiltable.YFilter }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) SetFilter(yf yfilter.YFilter) { cieifutiltable.YFilter = yf }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetGoName(yname string) string {
    if yname == "cieIfUtilEntry" { return "Cieifutilentry" }
    return ""
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetSegmentPath() string {
    return "cieIfUtilTable"
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfUtilEntry" {
        for _, c := range cieifutiltable.Cieifutilentry {
            if cieifutiltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry{}
        cieifutiltable.Cieifutilentry = append(cieifutiltable.Cieifutilentry, child)
        return &cieifutiltable.Cieifutilentry[len(cieifutiltable.Cieifutilentry)-1]
    }
    return nil
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifutiltable.Cieifutilentry {
        children[cieifutiltable.Cieifutilentry[i].GetSegmentPath()] = &cieifutiltable.Cieifutilentry[i]
    }
    return children
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetYangName() string { return "cieIfUtilTable" }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) SetParent(parent types.Entity) { cieifutiltable.parent = parent }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetParent() types.Entity { return cieifutiltable.parent }

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry
// An entry containing utilization rates for the
// interface.
// 
// Every interface for which the  inbound and 
// outbound traffic information is available
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the inbound packet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfInPktRate is the
    // exponentially-decayed moving average of inbound packet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    Cieifinpktrate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the inbound octet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfInOctetRate is the
    // exponentially-decayed moving average of inbound octet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are octets per second.
    Cieifinoctetrate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the outbound packet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfOutPktRate is the
    // exponentially-decayed moving average of outbound packet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    Cieifoutpktrate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the outbound octet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfOutOctetRate is the
    // exponentially-decayed moving average of outbound octet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are octets per second.
    Cieifoutoctetrate interface{}

    // This object specifies the time interval over which the inbound and outbound
    // traffic rates are calculated for this interface. The type is interface{}
    // with range: 1..4294967295. Units are seconds.
    Cieifinterval interface{}
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetFilter() yfilter.YFilter { return cieifutilentry.YFilter }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) SetFilter(yf yfilter.YFilter) { cieifutilentry.YFilter = yf }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfInPktRate" { return "Cieifinpktrate" }
    if yname == "cieIfInOctetRate" { return "Cieifinoctetrate" }
    if yname == "cieIfOutPktRate" { return "Cieifoutpktrate" }
    if yname == "cieIfOutOctetRate" { return "Cieifoutoctetrate" }
    if yname == "cieIfInterval" { return "Cieifinterval" }
    return ""
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetSegmentPath() string {
    return "cieIfUtilEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifutilentry.Ifindex) + "']"
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifutilentry.Ifindex
    leafs["cieIfInPktRate"] = cieifutilentry.Cieifinpktrate
    leafs["cieIfInOctetRate"] = cieifutilentry.Cieifinoctetrate
    leafs["cieIfOutPktRate"] = cieifutilentry.Cieifoutpktrate
    leafs["cieIfOutOctetRate"] = cieifutilentry.Cieifoutoctetrate
    leafs["cieIfInterval"] = cieifutilentry.Cieifinterval
    return leafs
}

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetYangName() string { return "cieIfUtilEntry" }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) SetParent(parent types.Entity) { cieifutilentry.parent = parent }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetParent() types.Entity { return cieifutilentry.parent }

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetParentYangName() string { return "cieIfUtilTable" }

// CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable
// This table contains the mappings of the
// ifIndex of an interface to its
// corresponding dot1dBasePort value.
type CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the mapping between the ifIndex value of an interface
    // and its corresponding dot1dBasePort value.  Every interface which has been
    // assigned a dot1dBasePort value by the system has a corresponding entry in
    // this table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry.
    Cieifdot1Dbasemappingentry []CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetFilter() yfilter.YFilter { return cieifdot1Dbasemappingtable.YFilter }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) SetFilter(yf yfilter.YFilter) { cieifdot1Dbasemappingtable.YFilter = yf }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetGoName(yname string) string {
    if yname == "cieIfDot1dBaseMappingEntry" { return "Cieifdot1Dbasemappingentry" }
    return ""
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetSegmentPath() string {
    return "cieIfDot1dBaseMappingTable"
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfDot1dBaseMappingEntry" {
        for _, c := range cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry {
            if cieifdot1Dbasemappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry{}
        cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry = append(cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry, child)
        return &cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry[len(cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry)-1]
    }
    return nil
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry {
        children[cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry[i].GetSegmentPath()] = &cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry[i]
    }
    return children
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetYangName() string { return "cieIfDot1dBaseMappingTable" }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) SetParent(parent types.Entity) { cieifdot1Dbasemappingtable.parent = parent }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetParent() types.Entity { return cieifdot1Dbasemappingtable.parent }

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry
// An entry containing the mapping between
// the ifIndex value of an interface and its
// corresponding dot1dBasePort value.
// 
// Every interface which has been assigned
// a dot1dBasePort value by the system
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The dot1dBasePort value for this interface. The type is interface{} with
    // range: 1..65535.
    Cieifdot1Dbasemappingport interface{}
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetFilter() yfilter.YFilter { return cieifdot1Dbasemappingentry.YFilter }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) SetFilter(yf yfilter.YFilter) { cieifdot1Dbasemappingentry.YFilter = yf }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cieIfDot1dBaseMappingPort" { return "Cieifdot1Dbasemappingport" }
    return ""
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetSegmentPath() string {
    return "cieIfDot1dBaseMappingEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifdot1Dbasemappingentry.Ifindex) + "']"
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cieifdot1Dbasemappingentry.Ifindex
    leafs["cieIfDot1dBaseMappingPort"] = cieifdot1Dbasemappingentry.Cieifdot1Dbasemappingport
    return leafs
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetYangName() string { return "cieIfDot1dBaseMappingEntry" }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) SetParent(parent types.Entity) { cieifdot1Dbasemappingentry.parent = parent }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetParent() types.Entity { return cieifdot1Dbasemappingentry.parent }

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetParentYangName() string { return "cieIfDot1dBaseMappingTable" }

// CISCOIFEXTENSIONMIB_Cieifnamemappingtable
// This table contains objects for providing
// the 'ifName' to 'ifIndex' mapping.
// This table contains one entry for each
// valid 'ifName' available in the system.
// Upon the first request, the implementation
// of this table will get all the available
// ifNames, and it will populate the entries
// in this table, it maintains this ifNames
// in a cache for ~30 seconds.
type CISCOIFEXTENSIONMIB_Cieifnamemappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry into the cieIfNameMappingTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry.
    Cieifnamemappingentry []CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetFilter() yfilter.YFilter { return cieifnamemappingtable.YFilter }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) SetFilter(yf yfilter.YFilter) { cieifnamemappingtable.YFilter = yf }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetGoName(yname string) string {
    if yname == "cieIfNameMappingEntry" { return "Cieifnamemappingentry" }
    return ""
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetSegmentPath() string {
    return "cieIfNameMappingTable"
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cieIfNameMappingEntry" {
        for _, c := range cieifnamemappingtable.Cieifnamemappingentry {
            if cieifnamemappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry{}
        cieifnamemappingtable.Cieifnamemappingentry = append(cieifnamemappingtable.Cieifnamemappingentry, child)
        return &cieifnamemappingtable.Cieifnamemappingentry[len(cieifnamemappingtable.Cieifnamemappingentry)-1]
    }
    return nil
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cieifnamemappingtable.Cieifnamemappingentry {
        children[cieifnamemappingtable.Cieifnamemappingentry[i].GetSegmentPath()] = &cieifnamemappingtable.Cieifnamemappingentry[i]
    }
    return children
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetYangName() string { return "cieIfNameMappingTable" }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) SetParent(parent types.Entity) { cieifnamemappingtable.parent = parent }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetParent() types.Entity { return cieifnamemappingtable.parent }

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetParentYangName() string { return "CISCO-IF-EXTENSION-MIB" }

// CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry
// An entry into the cieIfNameMappingTable.
type CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Represents an interface name mentioned in the
    // 'ifName' object of this system. The type is string with length: 1..112.
    Cieifname interface{}

    // This object represents the 'ifIndex' corresponding to the interface name
    // mentioned in the 'cieIfName' object of this instance. If the 'ifName'
    // mentioned in the 'cieIfName'  object of this instance corresponds to
    // multiple 'ifIndex' values, then the value of this object is the numerically
    // smallest of those multiple  'ifIndex' values. The type is interface{} with
    // range: 0..2147483647.
    Cieifindex interface{}
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetFilter() yfilter.YFilter { return cieifnamemappingentry.YFilter }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) SetFilter(yf yfilter.YFilter) { cieifnamemappingentry.YFilter = yf }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetGoName(yname string) string {
    if yname == "cieIfName" { return "Cieifname" }
    if yname == "cieIfIndex" { return "Cieifindex" }
    return ""
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetSegmentPath() string {
    return "cieIfNameMappingEntry" + "[cieIfName='" + fmt.Sprintf("%v", cieifnamemappingentry.Cieifname) + "']"
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cieIfName"] = cieifnamemappingentry.Cieifname
    leafs["cieIfIndex"] = cieifnamemappingentry.Cieifindex
    return leafs
}

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetYangName() string { return "cieIfNameMappingEntry" }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) SetParent(parent types.Entity) { cieifnamemappingentry.parent = parent }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetParent() types.Entity { return cieifnamemappingentry.parent }

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetParentYangName() string { return "cieIfNameMappingTable" }

