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
    EntityData types.CommonEntityData
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

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOIFEXTENSIONMIB.EntityData.YFilter = cISCOIFEXTENSIONMIB.YFilter
    cISCOIFEXTENSIONMIB.EntityData.YangName = "CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIFEXTENSIONMIB.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.SegmentPath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIFEXTENSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIFEXTENSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIFEXTENSIONMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIFEXTENSIONMIB.EntityData.Children["ciscoIfExtSystemConfig"] = types.YChild{"Ciscoifextsystemconfig", &cISCOIFEXTENSIONMIB.Ciscoifextsystemconfig}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfPacketStatsTable"] = types.YChild{"Cieifpacketstatstable", &cISCOIFEXTENSIONMIB.Cieifpacketstatstable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfInterfaceTable"] = types.YChild{"Cieifinterfacetable", &cISCOIFEXTENSIONMIB.Cieifinterfacetable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfStatusListTable"] = types.YChild{"Cieifstatuslisttable", &cISCOIFEXTENSIONMIB.Cieifstatuslisttable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfVlStatsTable"] = types.YChild{"Cieifvlstatstable", &cISCOIFEXTENSIONMIB.Cieifvlstatstable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfIndexPersistenceTable"] = types.YChild{"Cieifindexpersistencetable", &cISCOIFEXTENSIONMIB.Cieifindexpersistencetable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfDot1qCustomEtherTypeTable"] = types.YChild{"Cieifdot1Qcustomethertypetable", &cISCOIFEXTENSIONMIB.Cieifdot1Qcustomethertypetable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfUtilTable"] = types.YChild{"Cieifutiltable", &cISCOIFEXTENSIONMIB.Cieifutiltable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfDot1dBaseMappingTable"] = types.YChild{"Cieifdot1Dbasemappingtable", &cISCOIFEXTENSIONMIB.Cieifdot1Dbasemappingtable}
    cISCOIFEXTENSIONMIB.EntityData.Children["cieIfNameMappingTable"] = types.YChild{"Cieifnamemappingtable", &cISCOIFEXTENSIONMIB.Cieifnamemappingtable}
    cISCOIFEXTENSIONMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIFEXTENSIONMIB.EntityData)
}

// CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig
type CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig struct {
    EntityData types.CommonEntityData
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

func (ciscoifextsystemconfig *CISCOIFEXTENSIONMIB_Ciscoifextsystemconfig) GetEntityData() *types.CommonEntityData {
    ciscoifextsystemconfig.EntityData.YFilter = ciscoifextsystemconfig.YFilter
    ciscoifextsystemconfig.EntityData.YangName = "ciscoIfExtSystemConfig"
    ciscoifextsystemconfig.EntityData.BundleName = "cisco_ios_xe"
    ciscoifextsystemconfig.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    ciscoifextsystemconfig.EntityData.SegmentPath = "ciscoIfExtSystemConfig"
    ciscoifextsystemconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoifextsystemconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoifextsystemconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoifextsystemconfig.EntityData.Children = make(map[string]types.YChild)
    ciscoifextsystemconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoifextsystemconfig.EntityData.Leafs["cieSystemMtu"] = types.YLeaf{"Ciesystemmtu", ciscoifextsystemconfig.Ciesystemmtu}
    ciscoifextsystemconfig.EntityData.Leafs["cieLinkUpDownEnable"] = types.YLeaf{"Cielinkupdownenable", ciscoifextsystemconfig.Cielinkupdownenable}
    ciscoifextsystemconfig.EntityData.Leafs["cieStandardLinkUpDownVarbinds"] = types.YLeaf{"Ciestandardlinkupdownvarbinds", ciscoifextsystemconfig.Ciestandardlinkupdownvarbinds}
    ciscoifextsystemconfig.EntityData.Leafs["cieIfIndexPersistence"] = types.YLeaf{"Cieifindexpersistence", ciscoifextsystemconfig.Cieifindexpersistence}
    ciscoifextsystemconfig.EntityData.Leafs["cieDelayedLinkUpDownNotifEnable"] = types.YLeaf{"Ciedelayedlinkupdownnotifenable", ciscoifextsystemconfig.Ciedelayedlinkupdownnotifenable}
    ciscoifextsystemconfig.EntityData.Leafs["cieDelayedLinkUpDownNotifDelay"] = types.YLeaf{"Ciedelayedlinkupdownnotifdelay", ciscoifextsystemconfig.Ciedelayedlinkupdownnotifdelay}
    ciscoifextsystemconfig.EntityData.Leafs["cieIfIndexGlobalPersistence"] = types.YLeaf{"Cieifindexglobalpersistence", ciscoifextsystemconfig.Cieifindexglobalpersistence}
    ciscoifextsystemconfig.EntityData.Leafs["cieLinkUpDownConfig"] = types.YLeaf{"Cielinkupdownconfig", ciscoifextsystemconfig.Cielinkupdownconfig}
    return &(ciscoifextsystemconfig.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfPacketStatsTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry.
    Cieifpacketstatsentry []CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry
}

func (cieifpacketstatstable *CISCOIFEXTENSIONMIB_Cieifpacketstatstable) GetEntityData() *types.CommonEntityData {
    cieifpacketstatstable.EntityData.YFilter = cieifpacketstatstable.YFilter
    cieifpacketstatstable.EntityData.YangName = "cieIfPacketStatsTable"
    cieifpacketstatstable.EntityData.BundleName = "cisco_ios_xe"
    cieifpacketstatstable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifpacketstatstable.EntityData.SegmentPath = "cieIfPacketStatsTable"
    cieifpacketstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifpacketstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifpacketstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifpacketstatstable.EntityData.Children = make(map[string]types.YChild)
    cieifpacketstatstable.EntityData.Children["cieIfPacketStatsEntry"] = types.YChild{"Cieifpacketstatsentry", nil}
    for i := range cieifpacketstatstable.Cieifpacketstatsentry {
        cieifpacketstatstable.EntityData.Children[types.GetSegmentPath(&cieifpacketstatstable.Cieifpacketstatsentry[i])] = types.YChild{"Cieifpacketstatsentry", &cieifpacketstatstable.Cieifpacketstatsentry[i]}
    }
    cieifpacketstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifpacketstatstable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry
// An entry into the cieIfPacketStatsTable.
type CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry struct {
    EntityData types.CommonEntityData
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

func (cieifpacketstatsentry *CISCOIFEXTENSIONMIB_Cieifpacketstatstable_Cieifpacketstatsentry) GetEntityData() *types.CommonEntityData {
    cieifpacketstatsentry.EntityData.YFilter = cieifpacketstatsentry.YFilter
    cieifpacketstatsentry.EntityData.YangName = "cieIfPacketStatsEntry"
    cieifpacketstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cieifpacketstatsentry.EntityData.ParentYangName = "cieIfPacketStatsTable"
    cieifpacketstatsentry.EntityData.SegmentPath = "cieIfPacketStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifpacketstatsentry.Ifindex) + "']"
    cieifpacketstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifpacketstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifpacketstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifpacketstatsentry.EntityData.Children = make(map[string]types.YChild)
    cieifpacketstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifpacketstatsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifpacketstatsentry.Ifindex}
    cieifpacketstatsentry.EntityData.Leafs["cieIfLastInTime"] = types.YLeaf{"Cieiflastintime", cieifpacketstatsentry.Cieiflastintime}
    cieifpacketstatsentry.EntityData.Leafs["cieIfLastOutTime"] = types.YLeaf{"Cieiflastouttime", cieifpacketstatsentry.Cieiflastouttime}
    cieifpacketstatsentry.EntityData.Leafs["cieIfLastOutHangTime"] = types.YLeaf{"Cieiflastouthangtime", cieifpacketstatsentry.Cieiflastouthangtime}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInRuntsErrs"] = types.YLeaf{"Cieifinruntserrs", cieifpacketstatsentry.Cieifinruntserrs}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInGiantsErrs"] = types.YLeaf{"Cieifingiantserrs", cieifpacketstatsentry.Cieifingiantserrs}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInFramingErrs"] = types.YLeaf{"Cieifinframingerrs", cieifpacketstatsentry.Cieifinframingerrs}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInOverrunErrs"] = types.YLeaf{"Cieifinoverrunerrs", cieifpacketstatsentry.Cieifinoverrunerrs}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInIgnored"] = types.YLeaf{"Cieifinignored", cieifpacketstatsentry.Cieifinignored}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInAbortErrs"] = types.YLeaf{"Cieifinaborterrs", cieifpacketstatsentry.Cieifinaborterrs}
    cieifpacketstatsentry.EntityData.Leafs["cieIfInputQueueDrops"] = types.YLeaf{"Cieifinputqueuedrops", cieifpacketstatsentry.Cieifinputqueuedrops}
    cieifpacketstatsentry.EntityData.Leafs["cieIfOutputQueueDrops"] = types.YLeaf{"Cieifoutputqueuedrops", cieifpacketstatsentry.Cieifoutputqueuedrops}
    cieifpacketstatsentry.EntityData.Leafs["cieIfPacketDiscontinuityTime"] = types.YLeaf{"Cieifpacketdiscontinuitytime", cieifpacketstatsentry.Cieifpacketdiscontinuitytime}
    return &(cieifpacketstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfInterfaceTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry.
    Cieifinterfaceentry []CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry
}

func (cieifinterfacetable *CISCOIFEXTENSIONMIB_Cieifinterfacetable) GetEntityData() *types.CommonEntityData {
    cieifinterfacetable.EntityData.YFilter = cieifinterfacetable.YFilter
    cieifinterfacetable.EntityData.YangName = "cieIfInterfaceTable"
    cieifinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    cieifinterfacetable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifinterfacetable.EntityData.SegmentPath = "cieIfInterfaceTable"
    cieifinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifinterfacetable.EntityData.Children = make(map[string]types.YChild)
    cieifinterfacetable.EntityData.Children["cieIfInterfaceEntry"] = types.YChild{"Cieifinterfaceentry", nil}
    for i := range cieifinterfacetable.Cieifinterfaceentry {
        cieifinterfacetable.EntityData.Children[types.GetSegmentPath(&cieifinterfacetable.Cieifinterfaceentry[i])] = types.YChild{"Cieifinterfaceentry", &cieifinterfacetable.Cieifinterfaceentry[i]}
    }
    cieifinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifinterfacetable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry
// An entry into the cieIfInterfaceTable.
type CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry struct {
    EntityData types.CommonEntityData
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

func (cieifinterfaceentry *CISCOIFEXTENSIONMIB_Cieifinterfacetable_Cieifinterfaceentry) GetEntityData() *types.CommonEntityData {
    cieifinterfaceentry.EntityData.YFilter = cieifinterfaceentry.YFilter
    cieifinterfaceentry.EntityData.YangName = "cieIfInterfaceEntry"
    cieifinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    cieifinterfaceentry.EntityData.ParentYangName = "cieIfInterfaceTable"
    cieifinterfaceentry.EntityData.SegmentPath = "cieIfInterfaceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifinterfaceentry.Ifindex) + "']"
    cieifinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    cieifinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifinterfaceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifinterfaceentry.Ifindex}
    cieifinterfaceentry.EntityData.Leafs["cieIfResetCount"] = types.YLeaf{"Cieifresetcount", cieifinterfaceentry.Cieifresetcount}
    cieifinterfaceentry.EntityData.Leafs["cieIfKeepAliveEnabled"] = types.YLeaf{"Cieifkeepaliveenabled", cieifinterfaceentry.Cieifkeepaliveenabled}
    cieifinterfaceentry.EntityData.Leafs["cieIfStateChangeReason"] = types.YLeaf{"Cieifstatechangereason", cieifinterfaceentry.Cieifstatechangereason}
    cieifinterfaceentry.EntityData.Leafs["cieIfCarrierTransitionCount"] = types.YLeaf{"Cieifcarriertransitioncount", cieifinterfaceentry.Cieifcarriertransitioncount}
    cieifinterfaceentry.EntityData.Leafs["cieIfInterfaceDiscontinuityTime"] = types.YLeaf{"Cieifinterfacediscontinuitytime", cieifinterfaceentry.Cieifinterfacediscontinuitytime}
    cieifinterfaceentry.EntityData.Leafs["cieIfDhcpMode"] = types.YLeaf{"Cieifdhcpmode", cieifinterfaceentry.Cieifdhcpmode}
    cieifinterfaceentry.EntityData.Leafs["cieIfMtu"] = types.YLeaf{"Cieifmtu", cieifinterfaceentry.Cieifmtu}
    cieifinterfaceentry.EntityData.Leafs["cieIfContextName"] = types.YLeaf{"Cieifcontextname", cieifinterfaceentry.Cieifcontextname}
    cieifinterfaceentry.EntityData.Leafs["cieIfOperStatusCause"] = types.YLeaf{"Cieifoperstatuscause", cieifinterfaceentry.Cieifoperstatuscause}
    cieifinterfaceentry.EntityData.Leafs["cieIfOperStatusCauseDescr"] = types.YLeaf{"Cieifoperstatuscausedescr", cieifinterfaceentry.Cieifoperstatuscausedescr}
    cieifinterfaceentry.EntityData.Leafs["cieIfSpeedReceive"] = types.YLeaf{"Cieifspeedreceive", cieifinterfaceentry.Cieifspeedreceive}
    cieifinterfaceentry.EntityData.Leafs["cieIfHighSpeedReceive"] = types.YLeaf{"Cieifhighspeedreceive", cieifinterfaceentry.Cieifhighspeedreceive}
    cieifinterfaceentry.EntityData.Leafs["cieIfOwner"] = types.YLeaf{"Cieifowner", cieifinterfaceentry.Cieifowner}
    cieifinterfaceentry.EntityData.Leafs["cieIfSharedConfig"] = types.YLeaf{"Cieifsharedconfig", cieifinterfaceentry.Cieifsharedconfig}
    cieifinterfaceentry.EntityData.Leafs["cieIfSpeedGroupConfig"] = types.YLeaf{"Cieifspeedgroupconfig", cieifinterfaceentry.Cieifspeedgroupconfig}
    cieifinterfaceentry.EntityData.Leafs["cieIfTransceiverFrequencyConfig"] = types.YLeaf{"Cieiftransceiverfrequencyconfig", cieifinterfaceentry.Cieiftransceiverfrequencyconfig}
    cieifinterfaceentry.EntityData.Leafs["cieIfFillPatternConfig"] = types.YLeaf{"Cieiffillpatternconfig", cieifinterfaceentry.Cieiffillpatternconfig}
    cieifinterfaceentry.EntityData.Leafs["cieIfIgnoreBitErrorsConfig"] = types.YLeaf{"Cieifignorebiterrorsconfig", cieifinterfaceentry.Cieifignorebiterrorsconfig}
    cieifinterfaceentry.EntityData.Leafs["cieIfIgnoreInterruptThresholdConfig"] = types.YLeaf{"Cieifignoreinterruptthresholdconfig", cieifinterfaceentry.Cieifignoreinterruptthresholdconfig}
    return &(cieifinterfaceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents the 'ifIndex', interface operational mode and
    // interface  operational cause for a set of 64 interfaces  in a module. The
    // type is slice of
    // CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry.
    Cieifstatuslistentry []CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry
}

func (cieifstatuslisttable *CISCOIFEXTENSIONMIB_Cieifstatuslisttable) GetEntityData() *types.CommonEntityData {
    cieifstatuslisttable.EntityData.YFilter = cieifstatuslisttable.YFilter
    cieifstatuslisttable.EntityData.YangName = "cieIfStatusListTable"
    cieifstatuslisttable.EntityData.BundleName = "cisco_ios_xe"
    cieifstatuslisttable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifstatuslisttable.EntityData.SegmentPath = "cieIfStatusListTable"
    cieifstatuslisttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifstatuslisttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifstatuslisttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifstatuslisttable.EntityData.Children = make(map[string]types.YChild)
    cieifstatuslisttable.EntityData.Children["cieIfStatusListEntry"] = types.YChild{"Cieifstatuslistentry", nil}
    for i := range cieifstatuslisttable.Cieifstatuslistentry {
        cieifstatuslisttable.EntityData.Children[types.GetSegmentPath(&cieifstatuslisttable.Cieifstatuslistentry[i])] = types.YChild{"Cieifstatuslistentry", &cieifstatuslisttable.Cieifstatuslistentry[i]}
    }
    cieifstatuslisttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifstatuslisttable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry
// Each entry represents the 'ifIndex',
// interface operational mode and interface 
// operational cause for a set of 64 interfaces 
// in a module.
type CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry struct {
    EntityData types.CommonEntityData
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

func (cieifstatuslistentry *CISCOIFEXTENSIONMIB_Cieifstatuslisttable_Cieifstatuslistentry) GetEntityData() *types.CommonEntityData {
    cieifstatuslistentry.EntityData.YFilter = cieifstatuslistentry.YFilter
    cieifstatuslistentry.EntityData.YangName = "cieIfStatusListEntry"
    cieifstatuslistentry.EntityData.BundleName = "cisco_ios_xe"
    cieifstatuslistentry.EntityData.ParentYangName = "cieIfStatusListTable"
    cieifstatuslistentry.EntityData.SegmentPath = "cieIfStatusListEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cieifstatuslistentry.Entphysicalindex) + "']" + "[cieIfStatusListIndex='" + fmt.Sprintf("%v", cieifstatuslistentry.Cieifstatuslistindex) + "']"
    cieifstatuslistentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifstatuslistentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifstatuslistentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifstatuslistentry.EntityData.Children = make(map[string]types.YChild)
    cieifstatuslistentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifstatuslistentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cieifstatuslistentry.Entphysicalindex}
    cieifstatuslistentry.EntityData.Leafs["cieIfStatusListIndex"] = types.YLeaf{"Cieifstatuslistindex", cieifstatuslistentry.Cieifstatuslistindex}
    cieifstatuslistentry.EntityData.Leafs["cieInterfacesIndex"] = types.YLeaf{"Cieinterfacesindex", cieifstatuslistentry.Cieinterfacesindex}
    cieifstatuslistentry.EntityData.Leafs["cieInterfacesOperMode"] = types.YLeaf{"Cieinterfacesopermode", cieifstatuslistentry.Cieinterfacesopermode}
    cieifstatuslistentry.EntityData.Leafs["cieInterfacesOperCause"] = types.YLeaf{"Cieinterfacesopercause", cieifstatuslistentry.Cieinterfacesopercause}
    cieifstatuslistentry.EntityData.Leafs["cieInterfaceOwnershipBitmap"] = types.YLeaf{"Cieinterfaceownershipbitmap", cieifstatuslistentry.Cieinterfaceownershipbitmap}
    return &(cieifstatuslistentry.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifvlstatstable
// This table contains VL (Virtual Link) statistics
// for a capable interface.
// 
// Objects defined in this table may be 
// applicable to physical interfaces only.
type CISCOIFEXTENSIONMIB_Cieifvlstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains managed objects for Virtual Link statistics on interface
    // capable of  providing this information. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry.
    Cieifvlstatsentry []CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry
}

func (cieifvlstatstable *CISCOIFEXTENSIONMIB_Cieifvlstatstable) GetEntityData() *types.CommonEntityData {
    cieifvlstatstable.EntityData.YFilter = cieifvlstatstable.YFilter
    cieifvlstatstable.EntityData.YangName = "cieIfVlStatsTable"
    cieifvlstatstable.EntityData.BundleName = "cisco_ios_xe"
    cieifvlstatstable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifvlstatstable.EntityData.SegmentPath = "cieIfVlStatsTable"
    cieifvlstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifvlstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifvlstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifvlstatstable.EntityData.Children = make(map[string]types.YChild)
    cieifvlstatstable.EntityData.Children["cieIfVlStatsEntry"] = types.YChild{"Cieifvlstatsentry", nil}
    for i := range cieifvlstatstable.Cieifvlstatsentry {
        cieifvlstatstable.EntityData.Children[types.GetSegmentPath(&cieifvlstatstable.Cieifvlstatsentry[i])] = types.YChild{"Cieifvlstatsentry", &cieifvlstatstable.Cieifvlstatsentry[i]}
    }
    cieifvlstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifvlstatstable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry
// Each row contains managed objects for
// Virtual Link statistics on interface capable of 
// providing this information.
type CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry struct {
    EntityData types.CommonEntityData
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

func (cieifvlstatsentry *CISCOIFEXTENSIONMIB_Cieifvlstatstable_Cieifvlstatsentry) GetEntityData() *types.CommonEntityData {
    cieifvlstatsentry.EntityData.YFilter = cieifvlstatsentry.YFilter
    cieifvlstatsentry.EntityData.YangName = "cieIfVlStatsEntry"
    cieifvlstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cieifvlstatsentry.EntityData.ParentYangName = "cieIfVlStatsTable"
    cieifvlstatsentry.EntityData.SegmentPath = "cieIfVlStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifvlstatsentry.Ifindex) + "']"
    cieifvlstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifvlstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifvlstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifvlstatsentry.EntityData.Children = make(map[string]types.YChild)
    cieifvlstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifvlstatsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifvlstatsentry.Ifindex}
    cieifvlstatsentry.EntityData.Leafs["cieIfNoDropVlInPkts"] = types.YLeaf{"Cieifnodropvlinpkts", cieifvlstatsentry.Cieifnodropvlinpkts}
    cieifvlstatsentry.EntityData.Leafs["cieIfNoDropVlInOctets"] = types.YLeaf{"Cieifnodropvlinoctets", cieifvlstatsentry.Cieifnodropvlinoctets}
    cieifvlstatsentry.EntityData.Leafs["cieIfNoDropVlOutPkts"] = types.YLeaf{"Cieifnodropvloutpkts", cieifvlstatsentry.Cieifnodropvloutpkts}
    cieifvlstatsentry.EntityData.Leafs["cieIfNoDropVlOutOctets"] = types.YLeaf{"Cieifnodropvloutoctets", cieifvlstatsentry.Cieifnodropvloutoctets}
    cieifvlstatsentry.EntityData.Leafs["cieIfDropVlInPkts"] = types.YLeaf{"Cieifdropvlinpkts", cieifvlstatsentry.Cieifdropvlinpkts}
    cieifvlstatsentry.EntityData.Leafs["cieIfDropVlInOctets"] = types.YLeaf{"Cieifdropvlinoctets", cieifvlstatsentry.Cieifdropvlinoctets}
    cieifvlstatsentry.EntityData.Leafs["cieIfDropVlOutPkts"] = types.YLeaf{"Cieifdropvloutpkts", cieifvlstatsentry.Cieifdropvloutpkts}
    cieifvlstatsentry.EntityData.Leafs["cieIfDropVlOutOctets"] = types.YLeaf{"Cieifdropvloutoctets", cieifvlstatsentry.Cieifdropvloutoctets}
    return &(cieifvlstatsentry.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifindexpersistencetable
// This table lists configuration data relating to ifIndex
// persistence.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each ifEntry corresponding to an interface
// for which ifIndex persistence is supported.
type CISCOIFEXTENSIONMIB_Cieifindexpersistencetable struct {
    EntityData types.CommonEntityData
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

func (cieifindexpersistencetable *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable) GetEntityData() *types.CommonEntityData {
    cieifindexpersistencetable.EntityData.YFilter = cieifindexpersistencetable.YFilter
    cieifindexpersistencetable.EntityData.YangName = "cieIfIndexPersistenceTable"
    cieifindexpersistencetable.EntityData.BundleName = "cisco_ios_xe"
    cieifindexpersistencetable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifindexpersistencetable.EntityData.SegmentPath = "cieIfIndexPersistenceTable"
    cieifindexpersistencetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifindexpersistencetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifindexpersistencetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifindexpersistencetable.EntityData.Children = make(map[string]types.YChild)
    cieifindexpersistencetable.EntityData.Children["cieIfIndexPersistenceEntry"] = types.YChild{"Cieifindexpersistenceentry", nil}
    for i := range cieifindexpersistencetable.Cieifindexpersistenceentry {
        cieifindexpersistencetable.EntityData.Children[types.GetSegmentPath(&cieifindexpersistencetable.Cieifindexpersistenceentry[i])] = types.YChild{"Cieifindexpersistenceentry", &cieifindexpersistencetable.Cieifindexpersistenceentry[i]}
    }
    cieifindexpersistencetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifindexpersistencetable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry
// Each entry represents ifindex persistence configuration for an
// interface specified by ifIndex. Whenever an interface which
// supports ifindex persistence is created/destroyed in the
// ifTable, the corresponding ifindex persistence entry is
// created/destroyed respectively. Some of the interfaces may not
// support ifindex persistence, for example, a dynamic interface,
// such as a PPP connection or a IP subscriber interface.
type CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry struct {
    EntityData types.CommonEntityData
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

func (cieifindexpersistenceentry *CISCOIFEXTENSIONMIB_Cieifindexpersistencetable_Cieifindexpersistenceentry) GetEntityData() *types.CommonEntityData {
    cieifindexpersistenceentry.EntityData.YFilter = cieifindexpersistenceentry.YFilter
    cieifindexpersistenceentry.EntityData.YangName = "cieIfIndexPersistenceEntry"
    cieifindexpersistenceentry.EntityData.BundleName = "cisco_ios_xe"
    cieifindexpersistenceentry.EntityData.ParentYangName = "cieIfIndexPersistenceTable"
    cieifindexpersistenceentry.EntityData.SegmentPath = "cieIfIndexPersistenceEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifindexpersistenceentry.Ifindex) + "']"
    cieifindexpersistenceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifindexpersistenceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifindexpersistenceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifindexpersistenceentry.EntityData.Children = make(map[string]types.YChild)
    cieifindexpersistenceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifindexpersistenceentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifindexpersistenceentry.Ifindex}
    cieifindexpersistenceentry.EntityData.Leafs["cieIfIndexPersistenceEnabled"] = types.YLeaf{"Cieifindexpersistenceenabled", cieifindexpersistenceentry.Cieifindexpersistenceenabled}
    cieifindexpersistenceentry.EntityData.Leafs["cieIfIndexPersistenceControl"] = types.YLeaf{"Cieifindexpersistencecontrol", cieifindexpersistenceentry.Cieifindexpersistencecontrol}
    return &(cieifindexpersistenceentry.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable
// A list of the interfaces that support
// the 802.1q custom Ethertype feature.
type CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the custom EtherType information for the interface. 
    // Only interfaces with custom 802.1q ethertype control are listed in the 
    // table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry.
    Cieifdot1Qcustomethertypeentry []CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry
}

func (cieifdot1Qcustomethertypetable *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable) GetEntityData() *types.CommonEntityData {
    cieifdot1Qcustomethertypetable.EntityData.YFilter = cieifdot1Qcustomethertypetable.YFilter
    cieifdot1Qcustomethertypetable.EntityData.YangName = "cieIfDot1qCustomEtherTypeTable"
    cieifdot1Qcustomethertypetable.EntityData.BundleName = "cisco_ios_xe"
    cieifdot1Qcustomethertypetable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifdot1Qcustomethertypetable.EntityData.SegmentPath = "cieIfDot1qCustomEtherTypeTable"
    cieifdot1Qcustomethertypetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifdot1Qcustomethertypetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifdot1Qcustomethertypetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifdot1Qcustomethertypetable.EntityData.Children = make(map[string]types.YChild)
    cieifdot1Qcustomethertypetable.EntityData.Children["cieIfDot1qCustomEtherTypeEntry"] = types.YChild{"Cieifdot1Qcustomethertypeentry", nil}
    for i := range cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry {
        cieifdot1Qcustomethertypetable.EntityData.Children[types.GetSegmentPath(&cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry[i])] = types.YChild{"Cieifdot1Qcustomethertypeentry", &cieifdot1Qcustomethertypetable.Cieifdot1Qcustomethertypeentry[i]}
    }
    cieifdot1Qcustomethertypetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifdot1Qcustomethertypetable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry
// An entry containing the custom EtherType
// information for the interface.
// 
// Only interfaces with custom 802.1q
// ethertype control are listed in the 
// table.
type CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry struct {
    EntityData types.CommonEntityData
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

func (cieifdot1Qcustomethertypeentry *CISCOIFEXTENSIONMIB_Cieifdot1Qcustomethertypetable_Cieifdot1Qcustomethertypeentry) GetEntityData() *types.CommonEntityData {
    cieifdot1Qcustomethertypeentry.EntityData.YFilter = cieifdot1Qcustomethertypeentry.YFilter
    cieifdot1Qcustomethertypeentry.EntityData.YangName = "cieIfDot1qCustomEtherTypeEntry"
    cieifdot1Qcustomethertypeentry.EntityData.BundleName = "cisco_ios_xe"
    cieifdot1Qcustomethertypeentry.EntityData.ParentYangName = "cieIfDot1qCustomEtherTypeTable"
    cieifdot1Qcustomethertypeentry.EntityData.SegmentPath = "cieIfDot1qCustomEtherTypeEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifdot1Qcustomethertypeentry.Ifindex) + "']"
    cieifdot1Qcustomethertypeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifdot1Qcustomethertypeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifdot1Qcustomethertypeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifdot1Qcustomethertypeentry.EntityData.Children = make(map[string]types.YChild)
    cieifdot1Qcustomethertypeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifdot1Qcustomethertypeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifdot1Qcustomethertypeentry.Ifindex}
    cieifdot1Qcustomethertypeentry.EntityData.Leafs["cieIfDot1qCustomAdminEtherType"] = types.YLeaf{"Cieifdot1Qcustomadminethertype", cieifdot1Qcustomethertypeentry.Cieifdot1Qcustomadminethertype}
    cieifdot1Qcustomethertypeentry.EntityData.Leafs["cieIfDot1qCustomOperEtherType"] = types.YLeaf{"Cieifdot1Qcustomoperethertype", cieifdot1Qcustomethertypeentry.Cieifdot1Qcustomoperethertype}
    return &(cieifdot1Qcustomethertypeentry.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifutiltable
// This table contains the interface utilization
// rates for inbound and outbound traffic on an
// interface.
type CISCOIFEXTENSIONMIB_Cieifutiltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing utilization rates for the interface.  Every interface
    // for which the  inbound and  outbound traffic information is available has a
    // corresponding entry in this table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry.
    Cieifutilentry []CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry
}

func (cieifutiltable *CISCOIFEXTENSIONMIB_Cieifutiltable) GetEntityData() *types.CommonEntityData {
    cieifutiltable.EntityData.YFilter = cieifutiltable.YFilter
    cieifutiltable.EntityData.YangName = "cieIfUtilTable"
    cieifutiltable.EntityData.BundleName = "cisco_ios_xe"
    cieifutiltable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifutiltable.EntityData.SegmentPath = "cieIfUtilTable"
    cieifutiltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifutiltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifutiltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifutiltable.EntityData.Children = make(map[string]types.YChild)
    cieifutiltable.EntityData.Children["cieIfUtilEntry"] = types.YChild{"Cieifutilentry", nil}
    for i := range cieifutiltable.Cieifutilentry {
        cieifutiltable.EntityData.Children[types.GetSegmentPath(&cieifutiltable.Cieifutilentry[i])] = types.YChild{"Cieifutilentry", &cieifutiltable.Cieifutilentry[i]}
    }
    cieifutiltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifutiltable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry
// An entry containing utilization rates for the
// interface.
// 
// Every interface for which the  inbound and 
// outbound traffic information is available
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry struct {
    EntityData types.CommonEntityData
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

func (cieifutilentry *CISCOIFEXTENSIONMIB_Cieifutiltable_Cieifutilentry) GetEntityData() *types.CommonEntityData {
    cieifutilentry.EntityData.YFilter = cieifutilentry.YFilter
    cieifutilentry.EntityData.YangName = "cieIfUtilEntry"
    cieifutilentry.EntityData.BundleName = "cisco_ios_xe"
    cieifutilentry.EntityData.ParentYangName = "cieIfUtilTable"
    cieifutilentry.EntityData.SegmentPath = "cieIfUtilEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifutilentry.Ifindex) + "']"
    cieifutilentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifutilentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifutilentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifutilentry.EntityData.Children = make(map[string]types.YChild)
    cieifutilentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifutilentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifutilentry.Ifindex}
    cieifutilentry.EntityData.Leafs["cieIfInPktRate"] = types.YLeaf{"Cieifinpktrate", cieifutilentry.Cieifinpktrate}
    cieifutilentry.EntityData.Leafs["cieIfInOctetRate"] = types.YLeaf{"Cieifinoctetrate", cieifutilentry.Cieifinoctetrate}
    cieifutilentry.EntityData.Leafs["cieIfOutPktRate"] = types.YLeaf{"Cieifoutpktrate", cieifutilentry.Cieifoutpktrate}
    cieifutilentry.EntityData.Leafs["cieIfOutOctetRate"] = types.YLeaf{"Cieifoutoctetrate", cieifutilentry.Cieifoutoctetrate}
    cieifutilentry.EntityData.Leafs["cieIfInterval"] = types.YLeaf{"Cieifinterval", cieifutilentry.Cieifinterval}
    return &(cieifutilentry.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable
// This table contains the mappings of the
// ifIndex of an interface to its
// corresponding dot1dBasePort value.
type CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the mapping between the ifIndex value of an interface
    // and its corresponding dot1dBasePort value.  Every interface which has been
    // assigned a dot1dBasePort value by the system has a corresponding entry in
    // this table. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry.
    Cieifdot1Dbasemappingentry []CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry
}

func (cieifdot1Dbasemappingtable *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable) GetEntityData() *types.CommonEntityData {
    cieifdot1Dbasemappingtable.EntityData.YFilter = cieifdot1Dbasemappingtable.YFilter
    cieifdot1Dbasemappingtable.EntityData.YangName = "cieIfDot1dBaseMappingTable"
    cieifdot1Dbasemappingtable.EntityData.BundleName = "cisco_ios_xe"
    cieifdot1Dbasemappingtable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifdot1Dbasemappingtable.EntityData.SegmentPath = "cieIfDot1dBaseMappingTable"
    cieifdot1Dbasemappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifdot1Dbasemappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifdot1Dbasemappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifdot1Dbasemappingtable.EntityData.Children = make(map[string]types.YChild)
    cieifdot1Dbasemappingtable.EntityData.Children["cieIfDot1dBaseMappingEntry"] = types.YChild{"Cieifdot1Dbasemappingentry", nil}
    for i := range cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry {
        cieifdot1Dbasemappingtable.EntityData.Children[types.GetSegmentPath(&cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry[i])] = types.YChild{"Cieifdot1Dbasemappingentry", &cieifdot1Dbasemappingtable.Cieifdot1Dbasemappingentry[i]}
    }
    cieifdot1Dbasemappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifdot1Dbasemappingtable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry
// An entry containing the mapping between
// the ifIndex value of an interface and its
// corresponding dot1dBasePort value.
// 
// Every interface which has been assigned
// a dot1dBasePort value by the system
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The dot1dBasePort value for this interface. The type is interface{} with
    // range: 1..65535.
    Cieifdot1Dbasemappingport interface{}
}

func (cieifdot1Dbasemappingentry *CISCOIFEXTENSIONMIB_Cieifdot1Dbasemappingtable_Cieifdot1Dbasemappingentry) GetEntityData() *types.CommonEntityData {
    cieifdot1Dbasemappingentry.EntityData.YFilter = cieifdot1Dbasemappingentry.YFilter
    cieifdot1Dbasemappingentry.EntityData.YangName = "cieIfDot1dBaseMappingEntry"
    cieifdot1Dbasemappingentry.EntityData.BundleName = "cisco_ios_xe"
    cieifdot1Dbasemappingentry.EntityData.ParentYangName = "cieIfDot1dBaseMappingTable"
    cieifdot1Dbasemappingentry.EntityData.SegmentPath = "cieIfDot1dBaseMappingEntry" + "[ifIndex='" + fmt.Sprintf("%v", cieifdot1Dbasemappingentry.Ifindex) + "']"
    cieifdot1Dbasemappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifdot1Dbasemappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifdot1Dbasemappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifdot1Dbasemappingentry.EntityData.Children = make(map[string]types.YChild)
    cieifdot1Dbasemappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifdot1Dbasemappingentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cieifdot1Dbasemappingentry.Ifindex}
    cieifdot1Dbasemappingentry.EntityData.Leafs["cieIfDot1dBaseMappingPort"] = types.YLeaf{"Cieifdot1Dbasemappingport", cieifdot1Dbasemappingentry.Cieifdot1Dbasemappingport}
    return &(cieifdot1Dbasemappingentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfNameMappingTable. The type is slice of
    // CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry.
    Cieifnamemappingentry []CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry
}

func (cieifnamemappingtable *CISCOIFEXTENSIONMIB_Cieifnamemappingtable) GetEntityData() *types.CommonEntityData {
    cieifnamemappingtable.EntityData.YFilter = cieifnamemappingtable.YFilter
    cieifnamemappingtable.EntityData.YangName = "cieIfNameMappingTable"
    cieifnamemappingtable.EntityData.BundleName = "cisco_ios_xe"
    cieifnamemappingtable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieifnamemappingtable.EntityData.SegmentPath = "cieIfNameMappingTable"
    cieifnamemappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifnamemappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifnamemappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifnamemappingtable.EntityData.Children = make(map[string]types.YChild)
    cieifnamemappingtable.EntityData.Children["cieIfNameMappingEntry"] = types.YChild{"Cieifnamemappingentry", nil}
    for i := range cieifnamemappingtable.Cieifnamemappingentry {
        cieifnamemappingtable.EntityData.Children[types.GetSegmentPath(&cieifnamemappingtable.Cieifnamemappingentry[i])] = types.YChild{"Cieifnamemappingentry", &cieifnamemappingtable.Cieifnamemappingentry[i]}
    }
    cieifnamemappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cieifnamemappingtable.EntityData)
}

// CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry
// An entry into the cieIfNameMappingTable.
type CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry struct {
    EntityData types.CommonEntityData
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

func (cieifnamemappingentry *CISCOIFEXTENSIONMIB_Cieifnamemappingtable_Cieifnamemappingentry) GetEntityData() *types.CommonEntityData {
    cieifnamemappingentry.EntityData.YFilter = cieifnamemappingentry.YFilter
    cieifnamemappingentry.EntityData.YangName = "cieIfNameMappingEntry"
    cieifnamemappingentry.EntityData.BundleName = "cisco_ios_xe"
    cieifnamemappingentry.EntityData.ParentYangName = "cieIfNameMappingTable"
    cieifnamemappingentry.EntityData.SegmentPath = "cieIfNameMappingEntry" + "[cieIfName='" + fmt.Sprintf("%v", cieifnamemappingentry.Cieifname) + "']"
    cieifnamemappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieifnamemappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieifnamemappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieifnamemappingentry.EntityData.Children = make(map[string]types.YChild)
    cieifnamemappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cieifnamemappingentry.EntityData.Leafs["cieIfName"] = types.YLeaf{"Cieifname", cieifnamemappingentry.Cieifname}
    cieifnamemappingentry.EntityData.Leafs["cieIfIndex"] = types.YLeaf{"Cieifindex", cieifnamemappingentry.Cieifindex}
    return &(cieifnamemappingentry.EntityData)
}

