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

    
    CiscoIfExtSystemConfig CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig

    // This  table contains interface packet statistics which are not available in
    // IF-MIB(RFC2863).   As an example, some interfaces to which objects in this
    // table are applicable are as follows :          o Ethernet         o
    // FastEthernet         o ATM         o BRI         o Sonet         o
    // GigabitEthernet  Some objects defined in this table may be  applicable to
    // physical interfaces only. As a result, this table may be sparse for some
    // logical interfaces.
    CieIfPacketStatsTable CISCOIFEXTENSIONMIB_CieIfPacketStatsTable

    // This  table contains objects which provide more information about interface
    // properties not available in IF-MIB (RFC 2863).  Some objects defined in
    // this table may be applicable to physical interfaces only. As a result, this
    // table may be sparse for logical interfaces.
    CieIfInterfaceTable CISCOIFEXTENSIONMIB_CieIfInterfaceTable

    // This table contains objects for providing the 'ifIndex', interface
    // operational mode and  interface operational cause for all the  interfaces
    // in the modules.  This table contains one entry for each  64 interfaces in
    // an module.  This table provides efficient way of encoding  'ifIndex',
    // interface operational mode and interface operational cause, from the point 
    // of retrieval, by combining the values a set  of 64 interfaces in a single
    // MIB object.
    CieIfStatusListTable CISCOIFEXTENSIONMIB_CieIfStatusListTable

    // This table contains VL (Virtual Link) statistics for a capable interface. 
    // Objects defined in this table may be  applicable to physical interfaces
    // only.
    CieIfVlStatsTable CISCOIFEXTENSIONMIB_CieIfVlStatsTable

    // This table lists configuration data relating to ifIndex persistence.  This
    // table has a sparse dependent relationship on the ifTable, containing a row
    // for each ifEntry corresponding to an interface for which ifIndex
    // persistence is supported.
    CieIfIndexPersistenceTable CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable

    // A list of the interfaces that support the 802.1q custom Ethertype feature.
    CieIfDot1qCustomEtherTypeTable CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable

    // This table contains the interface utilization rates for inbound and
    // outbound traffic on an interface.
    CieIfUtilTable CISCOIFEXTENSIONMIB_CieIfUtilTable

    // This table contains the mappings of the ifIndex of an interface to its
    // corresponding dot1dBasePort value.
    CieIfDot1dBaseMappingTable CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable

    // This table contains objects for providing the 'ifName' to 'ifIndex'
    // mapping. This table contains one entry for each valid 'ifName' available in
    // the system. Upon the first request, the implementation of this table will
    // get all the available ifNames, and it will populate the entries in this
    // table, it maintains this ifNames in a cache for ~30 seconds.
    CieIfNameMappingTable CISCOIFEXTENSIONMIB_CieIfNameMappingTable
}

func (cISCOIFEXTENSIONMIB *CISCOIFEXTENSIONMIB) GetEntityData() *types.CommonEntityData {
    cISCOIFEXTENSIONMIB.EntityData.YFilter = cISCOIFEXTENSIONMIB.YFilter
    cISCOIFEXTENSIONMIB.EntityData.YangName = "CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIFEXTENSIONMIB.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.SegmentPath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB"
    cISCOIFEXTENSIONMIB.EntityData.AbsolutePath = cISCOIFEXTENSIONMIB.EntityData.SegmentPath
    cISCOIFEXTENSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIFEXTENSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIFEXTENSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIFEXTENSIONMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("ciscoIfExtSystemConfig", types.YChild{"CiscoIfExtSystemConfig", &cISCOIFEXTENSIONMIB.CiscoIfExtSystemConfig})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfPacketStatsTable", types.YChild{"CieIfPacketStatsTable", &cISCOIFEXTENSIONMIB.CieIfPacketStatsTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfInterfaceTable", types.YChild{"CieIfInterfaceTable", &cISCOIFEXTENSIONMIB.CieIfInterfaceTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfStatusListTable", types.YChild{"CieIfStatusListTable", &cISCOIFEXTENSIONMIB.CieIfStatusListTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfVlStatsTable", types.YChild{"CieIfVlStatsTable", &cISCOIFEXTENSIONMIB.CieIfVlStatsTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfIndexPersistenceTable", types.YChild{"CieIfIndexPersistenceTable", &cISCOIFEXTENSIONMIB.CieIfIndexPersistenceTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfDot1qCustomEtherTypeTable", types.YChild{"CieIfDot1qCustomEtherTypeTable", &cISCOIFEXTENSIONMIB.CieIfDot1qCustomEtherTypeTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfUtilTable", types.YChild{"CieIfUtilTable", &cISCOIFEXTENSIONMIB.CieIfUtilTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfDot1dBaseMappingTable", types.YChild{"CieIfDot1dBaseMappingTable", &cISCOIFEXTENSIONMIB.CieIfDot1dBaseMappingTable})
    cISCOIFEXTENSIONMIB.EntityData.Children.Append("cieIfNameMappingTable", types.YChild{"CieIfNameMappingTable", &cISCOIFEXTENSIONMIB.CieIfNameMappingTable})
    cISCOIFEXTENSIONMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIFEXTENSIONMIB.EntityData.YListKeys = []string {}

    return &(cISCOIFEXTENSIONMIB.EntityData)
}

// CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig
type CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global system MTU in octets. This object specifies the MTU on all
    // interfaces. However, the value specified by cieIfMtu takes precedence for
    // an interface, which means that the interface's MTU uses the value specified
    // by cieIfMtu, if it is configured. The type is interface{} with range:
    // -2147483648..2147483647.
    CieSystemMtu interface{}

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
    CieLinkUpDownEnable interface{}

    // Indicates whether to send the extra varbinds in addition to the varbinds
    // defined  in linkUp/linkDown notifications.  'standard'   - only send the
    // varbinds defined in                the standard linkUp/linkDown            
    // notification.   'additional' - send the extra varbinds in addition         
    // to the defined ones. 'other'      - any other config not covered by the
    // above.                This value is read-only. The type is
    // CieStandardLinkUpDownVarbinds.
    CieStandardLinkUpDownVarbinds interface{}

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
    CieIfIndexPersistence interface{}

    // This object specifies whether the system generates a
    // cieDelayedLinkUpDownNotif notification. The type is bool.
    CieDelayedLinkUpDownNotifEnable interface{}

    // This object specifies the interval of time an interface's operational
    // status must remain stable following a transition before the system will
    // generate a cieDelayedLinkUpDownNotif. The type is interface{} with range:
    // 1..60. Units are minutes.
    CieDelayedLinkUpDownNotifDelay interface{}

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
    CieIfIndexGlobalPersistence interface{}

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
    CieLinkUpDownConfig interface{}
}

func (ciscoIfExtSystemConfig *CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig) GetEntityData() *types.CommonEntityData {
    ciscoIfExtSystemConfig.EntityData.YFilter = ciscoIfExtSystemConfig.YFilter
    ciscoIfExtSystemConfig.EntityData.YangName = "ciscoIfExtSystemConfig"
    ciscoIfExtSystemConfig.EntityData.BundleName = "cisco_ios_xe"
    ciscoIfExtSystemConfig.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    ciscoIfExtSystemConfig.EntityData.SegmentPath = "ciscoIfExtSystemConfig"
    ciscoIfExtSystemConfig.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + ciscoIfExtSystemConfig.EntityData.SegmentPath
    ciscoIfExtSystemConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoIfExtSystemConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoIfExtSystemConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoIfExtSystemConfig.EntityData.Children = types.NewOrderedMap()
    ciscoIfExtSystemConfig.EntityData.Leafs = types.NewOrderedMap()
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieSystemMtu", types.YLeaf{"CieSystemMtu", ciscoIfExtSystemConfig.CieSystemMtu})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieLinkUpDownEnable", types.YLeaf{"CieLinkUpDownEnable", ciscoIfExtSystemConfig.CieLinkUpDownEnable})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieStandardLinkUpDownVarbinds", types.YLeaf{"CieStandardLinkUpDownVarbinds", ciscoIfExtSystemConfig.CieStandardLinkUpDownVarbinds})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieIfIndexPersistence", types.YLeaf{"CieIfIndexPersistence", ciscoIfExtSystemConfig.CieIfIndexPersistence})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieDelayedLinkUpDownNotifEnable", types.YLeaf{"CieDelayedLinkUpDownNotifEnable", ciscoIfExtSystemConfig.CieDelayedLinkUpDownNotifEnable})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieDelayedLinkUpDownNotifDelay", types.YLeaf{"CieDelayedLinkUpDownNotifDelay", ciscoIfExtSystemConfig.CieDelayedLinkUpDownNotifDelay})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieIfIndexGlobalPersistence", types.YLeaf{"CieIfIndexGlobalPersistence", ciscoIfExtSystemConfig.CieIfIndexGlobalPersistence})
    ciscoIfExtSystemConfig.EntityData.Leafs.Append("cieLinkUpDownConfig", types.YLeaf{"CieLinkUpDownConfig", ciscoIfExtSystemConfig.CieLinkUpDownConfig})

    ciscoIfExtSystemConfig.EntityData.YListKeys = []string {}

    return &(ciscoIfExtSystemConfig.EntityData)
}

// CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds represents                This value is read-only.
type CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds string

const (
    CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds_standard CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds = "standard"

    CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds_additional CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds = "additional"

    CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds_other CISCOIFEXTENSIONMIB_CiscoIfExtSystemConfig_CieStandardLinkUpDownVarbinds = "other"
)

// CISCOIFEXTENSIONMIB_CieIfPacketStatsTable
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
type CISCOIFEXTENSIONMIB_CieIfPacketStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfPacketStatsTable. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfPacketStatsTable_CieIfPacketStatsEntry.
    CieIfPacketStatsEntry []*CISCOIFEXTENSIONMIB_CieIfPacketStatsTable_CieIfPacketStatsEntry
}

func (cieIfPacketStatsTable *CISCOIFEXTENSIONMIB_CieIfPacketStatsTable) GetEntityData() *types.CommonEntityData {
    cieIfPacketStatsTable.EntityData.YFilter = cieIfPacketStatsTable.YFilter
    cieIfPacketStatsTable.EntityData.YangName = "cieIfPacketStatsTable"
    cieIfPacketStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfPacketStatsTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfPacketStatsTable.EntityData.SegmentPath = "cieIfPacketStatsTable"
    cieIfPacketStatsTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfPacketStatsTable.EntityData.SegmentPath
    cieIfPacketStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfPacketStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfPacketStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfPacketStatsTable.EntityData.Children = types.NewOrderedMap()
    cieIfPacketStatsTable.EntityData.Children.Append("cieIfPacketStatsEntry", types.YChild{"CieIfPacketStatsEntry", nil})
    for i := range cieIfPacketStatsTable.CieIfPacketStatsEntry {
        cieIfPacketStatsTable.EntityData.Children.Append(types.GetSegmentPath(cieIfPacketStatsTable.CieIfPacketStatsEntry[i]), types.YChild{"CieIfPacketStatsEntry", cieIfPacketStatsTable.CieIfPacketStatsEntry[i]})
    }
    cieIfPacketStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfPacketStatsTable.EntityData.YListKeys = []string {}

    return &(cieIfPacketStatsTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfPacketStatsTable_CieIfPacketStatsEntry
// An entry into the cieIfPacketStatsTable.
type CISCOIFEXTENSIONMIB_CieIfPacketStatsTable_CieIfPacketStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // input  packet was received.  Discontinuities in the value of this variable
    // can occur at re-initialization of the management system, and at other times
    // as  indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CieIfLastInTime interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // output  packet was transmitted.  Discontinuities in the value of this
    // variable can occur at re-initialization of the management system, and at
    // other times as  indicated by the values of  cieIfPacketDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    CieIfLastOutTime interface{}

    // This object represents the elapsed time in milliseconds since last protocol
    // output packet could not be successfully transmitted.  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CieIfLastOutHangTime interface{}

    // The number of packets input on a particular physical interface which were
    // dropped as they were smaller than the minimum allowable  physical media
    // limit.  Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    CieIfInRuntsErrs interface{}

    // The number of input packets on a particular physical interface which were
    // dropped as  they were larger than the ifMtu (largest  permitted  size of a
    // packet which can be  sent/received on an interface).  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    CieIfInGiantsErrs interface{}

    // The number of input packets on a physical interface which were misaligned
    // or had framing errors. This happens when the  format of the incoming packet
    // on a physical interface is incorrect.  Discontinuities in the value of this
    // variable can occur at re-initialization of the management system, and at
    // other times as  indicated by the values of  cieIfPacketDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    CieIfInFramingErrs interface{}

    // The number of input packets which arrived on a particular physical
    // interface which  were too quick for the hardware to receive and hence the
    // receiver ran out of buffers.  Discontinuities in the value of this variable
    // can occur at re-initialization of the management system, and at other times
    // as  indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    CieIfInOverrunErrs interface{}

    // The number of input packets which were simply ignored by this physical
    // interface due to  insufficient resources to handle the incoming packets. 
    // For example, this could indicate that the input receive buffers are not
    // available or that the receiver lost a packet.  Discontinuities in the value
    // of this variable can occur at re-initialization of the management system,
    // and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    CieIfInIgnored interface{}

    // Number of input packets which were dropped because the receiver aborted. 
    // Examples of this could be when an abort sequence aborted the input frame or
    // when there is a collision in an ethernet segment.  Discontinuities in the
    // value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    CieIfInAbortErrs interface{}

    // The number of input packets which were dropped.  Some reasons why this
    // object could be  incremented are:  o  Input queue is full. o  Errors at the
    // receiver hardware     while receiving the packet.  Discontinuities in the
    // value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfPacketDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    CieIfInputQueueDrops interface{}

    // This object indicates the  number of output packets dropped by the
    // interface even though no error had been detected to prevent them being
    // transmitted.   The packet could be dropped for many reasons, which could
    // range from the interface being down to errors in the format of the packet. 
    // Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfPacketDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    CieIfOutputQueueDrops interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters suffered a discontinuity.   If no such
    // discontinuities have occurred  since the last re-initialization of the 
    // local management subsystem, then this  object contains a value of zero. The
    // type is interface{} with range: 0..4294967295.
    CieIfPacketDiscontinuityTime interface{}
}

func (cieIfPacketStatsEntry *CISCOIFEXTENSIONMIB_CieIfPacketStatsTable_CieIfPacketStatsEntry) GetEntityData() *types.CommonEntityData {
    cieIfPacketStatsEntry.EntityData.YFilter = cieIfPacketStatsEntry.YFilter
    cieIfPacketStatsEntry.EntityData.YangName = "cieIfPacketStatsEntry"
    cieIfPacketStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfPacketStatsEntry.EntityData.ParentYangName = "cieIfPacketStatsTable"
    cieIfPacketStatsEntry.EntityData.SegmentPath = "cieIfPacketStatsEntry" + types.AddKeyToken(cieIfPacketStatsEntry.IfIndex, "ifIndex")
    cieIfPacketStatsEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfPacketStatsTable/" + cieIfPacketStatsEntry.EntityData.SegmentPath
    cieIfPacketStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfPacketStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfPacketStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfPacketStatsEntry.EntityData.Children = types.NewOrderedMap()
    cieIfPacketStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfPacketStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfPacketStatsEntry.IfIndex})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfLastInTime", types.YLeaf{"CieIfLastInTime", cieIfPacketStatsEntry.CieIfLastInTime})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfLastOutTime", types.YLeaf{"CieIfLastOutTime", cieIfPacketStatsEntry.CieIfLastOutTime})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfLastOutHangTime", types.YLeaf{"CieIfLastOutHangTime", cieIfPacketStatsEntry.CieIfLastOutHangTime})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInRuntsErrs", types.YLeaf{"CieIfInRuntsErrs", cieIfPacketStatsEntry.CieIfInRuntsErrs})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInGiantsErrs", types.YLeaf{"CieIfInGiantsErrs", cieIfPacketStatsEntry.CieIfInGiantsErrs})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInFramingErrs", types.YLeaf{"CieIfInFramingErrs", cieIfPacketStatsEntry.CieIfInFramingErrs})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInOverrunErrs", types.YLeaf{"CieIfInOverrunErrs", cieIfPacketStatsEntry.CieIfInOverrunErrs})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInIgnored", types.YLeaf{"CieIfInIgnored", cieIfPacketStatsEntry.CieIfInIgnored})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInAbortErrs", types.YLeaf{"CieIfInAbortErrs", cieIfPacketStatsEntry.CieIfInAbortErrs})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfInputQueueDrops", types.YLeaf{"CieIfInputQueueDrops", cieIfPacketStatsEntry.CieIfInputQueueDrops})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfOutputQueueDrops", types.YLeaf{"CieIfOutputQueueDrops", cieIfPacketStatsEntry.CieIfOutputQueueDrops})
    cieIfPacketStatsEntry.EntityData.Leafs.Append("cieIfPacketDiscontinuityTime", types.YLeaf{"CieIfPacketDiscontinuityTime", cieIfPacketStatsEntry.CieIfPacketDiscontinuityTime})

    cieIfPacketStatsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfPacketStatsEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable
// This  table contains objects which provide
// more information about interface  
// properties not available in IF-MIB
// (RFC 2863).
// 
// Some objects defined in this table may be
// applicable to physical interfaces only.
// As a result, this table may be sparse for
// logical interfaces.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfInterfaceTable. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry.
    CieIfInterfaceEntry []*CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry
}

func (cieIfInterfaceTable *CISCOIFEXTENSIONMIB_CieIfInterfaceTable) GetEntityData() *types.CommonEntityData {
    cieIfInterfaceTable.EntityData.YFilter = cieIfInterfaceTable.YFilter
    cieIfInterfaceTable.EntityData.YangName = "cieIfInterfaceTable"
    cieIfInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfInterfaceTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfInterfaceTable.EntityData.SegmentPath = "cieIfInterfaceTable"
    cieIfInterfaceTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfInterfaceTable.EntityData.SegmentPath
    cieIfInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfInterfaceTable.EntityData.Children = types.NewOrderedMap()
    cieIfInterfaceTable.EntityData.Children.Append("cieIfInterfaceEntry", types.YChild{"CieIfInterfaceEntry", nil})
    for i := range cieIfInterfaceTable.CieIfInterfaceEntry {
        cieIfInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(cieIfInterfaceTable.CieIfInterfaceEntry[i]), types.YChild{"CieIfInterfaceEntry", cieIfInterfaceTable.CieIfInterfaceEntry[i]})
    }
    cieIfInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfInterfaceTable.EntityData.YListKeys = []string {}

    return &(cieIfInterfaceTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry
// An entry into the cieIfInterfaceTable.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of times the interface was internally reset and brought up. 
    // Some of the actions which can cause this counter to increment are :  o 
    // Bringing an interface up using the     interface CLI command.  o  Clearing
    // the interface with the exec    CLI command.  o  Bringing the interface up
    // via SNMP.  Discontinuities in the value of this variable can occur at
    // re-initialization of the management system, and at other times as 
    // indicated by the values of  cieIfInterfaceDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    CieIfResetCount interface{}

    // A keepalive is a small, layer-2 message that is transmitted by a network
    // device  to let directly-connected network devices know of its presence. 
    // This object returns 'true' if keepalives are enabled on this interface. If
    // keepalives are not enabled, 'false' is returned.  Setting this object to
    // TRUE or FALSE enables or disables (respectively) keepalive on this 
    // interface. The type is bool.
    CieIfKeepAliveEnabled interface{}

    // This object displays a human-readable textual string which describes the 
    // cause of the last state change of the  interface.  Examples of the values
    // this object can take are:  o  'Lost Carrier' o  'administratively down' o 
    // 'up' o  'down'. The type is string.
    CieIfStateChangeReason interface{}

    // Number of times interface saw the carrier signal transition.  For example,
    // if a T1 line is unplugged,  then framer will detect the loss of signal 
    // (LOS) on the line  and will count it as a transition.  Discontinuities in
    // the value of this variable can occur at re-initialization of the management
    // system, and at other times as  indicated by the values of 
    // cieIfInterfaceDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    CieIfCarrierTransitionCount interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters  suffered  a discontinuity.   If no such
    // discontinuities have occurred  since the last re-initialization of the 
    // local management subsystem, then this  object contains a value of zero. The
    // type is interface{} with range: 0..4294967295.
    CieIfInterfaceDiscontinuityTime interface{}

    // The DHCP mode configured by the administrator. If 'true' the DHCP is
    // enabled. In which case an IP address is requested in DHCP. This is in
    // addition to any that are  configured by the administrator in
    // 'ciiIPAddressTable' or 'ciiIPIfAddressTable' in CISCO-IP-IF-MIB. If 'false'
    // the DHCP is disabled. In which case all IP addresses are configured by the
    // administrator in 'ciiIPAddressTable' or  'ciiIPIfAddressTable'. For
    // interfaces, for which DHCP cannot be or is not supported, then this object
    // has the value 'false'. The type is bool.
    CieIfDhcpMode interface{}

    // The MTU configured by the administrator. This object is exactly same as
    // 'ifMtu' in  ifTable from IF-MIB for the same ifIndex value , except that it
    // is configurable by the administrator. For more description of this object
    // refer to 'ifMtu' in IF-MIB. The type is interface{} with range:
    // 40..2147483647.
    CieIfMtu interface{}

    // The ContextName denotes the interface 'context' and is used to logically
    // separate the MIB management. RFC 2571 and RFC 2737 describe this approach.
    // When the agent supports a different SNMP  context, as detailed in RFC 2571
    // and  RFC 2737, for different interfaces, then the value of this object
    // specifies the context name used for this interface. The type is string with
    // length: 0..32.
    CieIfContextName interface{}

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
    CieIfOperStatusCause interface{}

    // The description for the cause of current operational state of the
    // interface, given  by the object 'cieIfOperStatusCause'.  For an interface
    // whose 'ifOperStatus' is not 'down' the value of this object will be 
    // 'none'. The type is string.
    CieIfOperStatusCauseDescr interface{}

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
    CieIfSpeedReceive interface{}

    // An estimate of the interface's current receive bandwidth in units of
    // 1,000,000 bits per second.  If this object reports a value of `n' then the
    // speed of the interface is somewhere in the range of `n-500,000' to
    // `n+499,999'.  For interfaces which do not vary in bandwidth or for those
    // where no accurate estimation can be made, this object should contain the
    // nominal bandwidth.  For a sub-layer which has no concept of bandwidth, this
    // object should be zero. The type is interface{} with range: 0..4294967295.
    CieIfHighSpeedReceive interface{}

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
    CieIfOwner interface{}

    // This object indicates the current configuration of interface sharing on the
    // given interface.  'notApplicable' - the interface sharing configuration on 
    // this interface is not applicable.  'ownerDedicated' - the interface is in
    // the dedicated mode             to the binding physical interface.
    // 'ownerShared' - the interface is shared amongst virtual switches         
    // and this interface physically belongs to a its           virtual switch.  
    // 'sharedOnly' - the interface is in purely shared mode. The type is
    // CieIfSharedConfig.
    CieIfSharedConfig interface{}

    // This object specifies the current speed group configuration on the given
    // interface.  'notApplicable' - the interface speed group configuration on   
    // this interface is not applicable. It is a              read-only value.
    // '10G' - the interface speed group configuration on             this
    // interface as 10G. '1G-2G-4G-8G' - the interface speed group configuration  
    // on this interface as 1G-2G-4G-8G. '2G-4G-8G-16G' - the interface speed
    // group configuration              on this interface as 2G-4G-8G-16G. The
    // type is CieIfSpeedGroupConfig.
    CieIfSpeedGroupConfig interface{}

    // This object specifies the current transceiver frequency configuration on
    // the given interface.  'notApplicable' - the interface transceiver frequency
    // 				  configuration on this interface  				  is not applicable. It is a
    // read-only value. 'FibreChannel' - the interface transceiver frequency 				
    // configuration on this interface as                   Fibre Channel.
    // 'Ethernet'	  -  the interface transceiver frequency on 				 this interface
    // as Ethernet. The type is CieIfTransceiverFrequencyConfig.
    CieIfTransceiverFrequencyConfig interface{}

    // This object specifies the current switchport fill pattern configuration on
    // the given interface.  'arbff8G' - the inter frame gap fill pattern is
    // 			ARBFF for 8G speed. 'idle8G' - the inter frame gap fill pattern is 		  
    // IDLE for 8G speed. The type is CieIfFillPatternConfig.
    CieIfFillPatternConfig interface{}

    // This object specifies the current switchport biterrors configuration on the
    // given interface.  If 'true(1)' the ignore bit errors is enabled.In which
    // case the interface ignores bit errors. If 'false(2)' the ignore bit errors
    // is disabled. In which  case the interface acts on the bit errors.  For
    // interfaces, for which bit errors  is not supported, then this object has
    // the value 'true(1)'. The type is bool.
    CieIfIgnoreBitErrorsConfig interface{}

    // This object specifies the current interrupt threshold configuration on the
    // given interface.  'If 'true(1)' the ignore interrupt thresholds is enabled.
    // In which case the interface ignores interrupt thresholds. If 'false(2)' the
    // ignore interrupt thresholds is disabled. In which case the interface acts
    // on the interrupt  thresholds.  For interfaces, for which interrupt
    // thresholds  is not supported, then this object has the  value 'true(1)'.
    // The type is bool.
    CieIfIgnoreInterruptThresholdConfig interface{}
}

func (cieIfInterfaceEntry *CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry) GetEntityData() *types.CommonEntityData {
    cieIfInterfaceEntry.EntityData.YFilter = cieIfInterfaceEntry.YFilter
    cieIfInterfaceEntry.EntityData.YangName = "cieIfInterfaceEntry"
    cieIfInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfInterfaceEntry.EntityData.ParentYangName = "cieIfInterfaceTable"
    cieIfInterfaceEntry.EntityData.SegmentPath = "cieIfInterfaceEntry" + types.AddKeyToken(cieIfInterfaceEntry.IfIndex, "ifIndex")
    cieIfInterfaceEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfInterfaceTable/" + cieIfInterfaceEntry.EntityData.SegmentPath
    cieIfInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    cieIfInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfInterfaceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfInterfaceEntry.IfIndex})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfResetCount", types.YLeaf{"CieIfResetCount", cieIfInterfaceEntry.CieIfResetCount})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfKeepAliveEnabled", types.YLeaf{"CieIfKeepAliveEnabled", cieIfInterfaceEntry.CieIfKeepAliveEnabled})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfStateChangeReason", types.YLeaf{"CieIfStateChangeReason", cieIfInterfaceEntry.CieIfStateChangeReason})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfCarrierTransitionCount", types.YLeaf{"CieIfCarrierTransitionCount", cieIfInterfaceEntry.CieIfCarrierTransitionCount})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfInterfaceDiscontinuityTime", types.YLeaf{"CieIfInterfaceDiscontinuityTime", cieIfInterfaceEntry.CieIfInterfaceDiscontinuityTime})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfDhcpMode", types.YLeaf{"CieIfDhcpMode", cieIfInterfaceEntry.CieIfDhcpMode})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfMtu", types.YLeaf{"CieIfMtu", cieIfInterfaceEntry.CieIfMtu})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfContextName", types.YLeaf{"CieIfContextName", cieIfInterfaceEntry.CieIfContextName})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfOperStatusCause", types.YLeaf{"CieIfOperStatusCause", cieIfInterfaceEntry.CieIfOperStatusCause})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfOperStatusCauseDescr", types.YLeaf{"CieIfOperStatusCauseDescr", cieIfInterfaceEntry.CieIfOperStatusCauseDescr})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfSpeedReceive", types.YLeaf{"CieIfSpeedReceive", cieIfInterfaceEntry.CieIfSpeedReceive})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfHighSpeedReceive", types.YLeaf{"CieIfHighSpeedReceive", cieIfInterfaceEntry.CieIfHighSpeedReceive})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfOwner", types.YLeaf{"CieIfOwner", cieIfInterfaceEntry.CieIfOwner})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfSharedConfig", types.YLeaf{"CieIfSharedConfig", cieIfInterfaceEntry.CieIfSharedConfig})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfSpeedGroupConfig", types.YLeaf{"CieIfSpeedGroupConfig", cieIfInterfaceEntry.CieIfSpeedGroupConfig})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfTransceiverFrequencyConfig", types.YLeaf{"CieIfTransceiverFrequencyConfig", cieIfInterfaceEntry.CieIfTransceiverFrequencyConfig})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfFillPatternConfig", types.YLeaf{"CieIfFillPatternConfig", cieIfInterfaceEntry.CieIfFillPatternConfig})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfIgnoreBitErrorsConfig", types.YLeaf{"CieIfIgnoreBitErrorsConfig", cieIfInterfaceEntry.CieIfIgnoreBitErrorsConfig})
    cieIfInterfaceEntry.EntityData.Leafs.Append("cieIfIgnoreInterruptThresholdConfig", types.YLeaf{"CieIfIgnoreInterruptThresholdConfig", cieIfInterfaceEntry.CieIfIgnoreInterruptThresholdConfig})

    cieIfInterfaceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfInterfaceEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig represents 		   IDLE for 8G speed.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig string

const (
    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig_arbff8G CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig = "arbff8G"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig_idle8G CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfFillPatternConfig = "idle8G"
)

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig represents 'sharedOnly' - the interface is in purely shared mode.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig string

const (
    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig_notApplicable CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig = "notApplicable"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig_ownerDedicated CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig = "ownerDedicated"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig_ownerShared CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig = "ownerShared"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig_sharedOnly CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSharedConfig = "sharedOnly"
)

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig represents             on this interface as 2G-4G-8G-16G.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig string

const (
    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig_notApplicable CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig = "notApplicable"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig_tenG CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig = "tenG"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig_oneTwoFourEightG CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig = "oneTwoFourEightG"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig_twoFourEightSixteenG CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfSpeedGroupConfig = "twoFourEightSixteenG"
)

// CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig represents 				 this interface as Ethernet.
type CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig string

const (
    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig_notApplicable CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig = "notApplicable"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig_fibreChannel CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig = "fibreChannel"

    CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig_ethernet CISCOIFEXTENSIONMIB_CieIfInterfaceTable_CieIfInterfaceEntry_CieIfTransceiverFrequencyConfig = "ethernet"
)

// CISCOIFEXTENSIONMIB_CieIfStatusListTable
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
type CISCOIFEXTENSIONMIB_CieIfStatusListTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents the 'ifIndex', interface operational mode and
    // interface  operational cause for a set of 64 interfaces  in a module. The
    // type is slice of
    // CISCOIFEXTENSIONMIB_CieIfStatusListTable_CieIfStatusListEntry.
    CieIfStatusListEntry []*CISCOIFEXTENSIONMIB_CieIfStatusListTable_CieIfStatusListEntry
}

func (cieIfStatusListTable *CISCOIFEXTENSIONMIB_CieIfStatusListTable) GetEntityData() *types.CommonEntityData {
    cieIfStatusListTable.EntityData.YFilter = cieIfStatusListTable.YFilter
    cieIfStatusListTable.EntityData.YangName = "cieIfStatusListTable"
    cieIfStatusListTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfStatusListTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfStatusListTable.EntityData.SegmentPath = "cieIfStatusListTable"
    cieIfStatusListTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfStatusListTable.EntityData.SegmentPath
    cieIfStatusListTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfStatusListTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfStatusListTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfStatusListTable.EntityData.Children = types.NewOrderedMap()
    cieIfStatusListTable.EntityData.Children.Append("cieIfStatusListEntry", types.YChild{"CieIfStatusListEntry", nil})
    for i := range cieIfStatusListTable.CieIfStatusListEntry {
        cieIfStatusListTable.EntityData.Children.Append(types.GetSegmentPath(cieIfStatusListTable.CieIfStatusListEntry[i]), types.YChild{"CieIfStatusListEntry", cieIfStatusListTable.CieIfStatusListEntry[i]})
    }
    cieIfStatusListTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfStatusListTable.EntityData.YListKeys = []string {}

    return &(cieIfStatusListTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfStatusListTable_CieIfStatusListEntry
// Each entry represents the 'ifIndex',
// interface operational mode and interface 
// operational cause for a set of 64 interfaces 
// in a module.
type CISCOIFEXTENSIONMIB_CieIfStatusListTable_CieIfStatusListEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. An arbitrary integer value, greater than zero,
    // which identifies a list of 64 interfaces within a module. The type is
    // interface{} with range: 1..33554432.
    CieIfStatusListIndex interface{}

    // This object represents the 'ifIndex' for a set of 64 interfaces in the
    // module. The type is string with length: 0..256.
    CieInterfacesIndex interface{}

    // This object represents the operational mode for a set of 64 interfaces in
    // the module. The type is string with length: 0..64.
    CieInterfacesOperMode interface{}

    // This object represents the operational status cause for a set of 64
    // interfaces in the  module. The type is string with length: 0..128.
    CieInterfacesOperCause interface{}

    // This object indicates the status for a set of 64 interfaces in a module
    // regarding whether or not each interface is  administratively assigned a
    // name of the current owner of the  interface resource as per cieIfOwner. The
    // type is string with length: 0..8.
    CieInterfaceOwnershipBitmap interface{}
}

func (cieIfStatusListEntry *CISCOIFEXTENSIONMIB_CieIfStatusListTable_CieIfStatusListEntry) GetEntityData() *types.CommonEntityData {
    cieIfStatusListEntry.EntityData.YFilter = cieIfStatusListEntry.YFilter
    cieIfStatusListEntry.EntityData.YangName = "cieIfStatusListEntry"
    cieIfStatusListEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfStatusListEntry.EntityData.ParentYangName = "cieIfStatusListTable"
    cieIfStatusListEntry.EntityData.SegmentPath = "cieIfStatusListEntry" + types.AddKeyToken(cieIfStatusListEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cieIfStatusListEntry.CieIfStatusListIndex, "cieIfStatusListIndex")
    cieIfStatusListEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfStatusListTable/" + cieIfStatusListEntry.EntityData.SegmentPath
    cieIfStatusListEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfStatusListEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfStatusListEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfStatusListEntry.EntityData.Children = types.NewOrderedMap()
    cieIfStatusListEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfStatusListEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cieIfStatusListEntry.EntPhysicalIndex})
    cieIfStatusListEntry.EntityData.Leafs.Append("cieIfStatusListIndex", types.YLeaf{"CieIfStatusListIndex", cieIfStatusListEntry.CieIfStatusListIndex})
    cieIfStatusListEntry.EntityData.Leafs.Append("cieInterfacesIndex", types.YLeaf{"CieInterfacesIndex", cieIfStatusListEntry.CieInterfacesIndex})
    cieIfStatusListEntry.EntityData.Leafs.Append("cieInterfacesOperMode", types.YLeaf{"CieInterfacesOperMode", cieIfStatusListEntry.CieInterfacesOperMode})
    cieIfStatusListEntry.EntityData.Leafs.Append("cieInterfacesOperCause", types.YLeaf{"CieInterfacesOperCause", cieIfStatusListEntry.CieInterfacesOperCause})
    cieIfStatusListEntry.EntityData.Leafs.Append("cieInterfaceOwnershipBitmap", types.YLeaf{"CieInterfaceOwnershipBitmap", cieIfStatusListEntry.CieInterfaceOwnershipBitmap})

    cieIfStatusListEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CieIfStatusListIndex"}

    return &(cieIfStatusListEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfVlStatsTable
// This table contains VL (Virtual Link) statistics
// for a capable interface.
// 
// Objects defined in this table may be 
// applicable to physical interfaces only.
type CISCOIFEXTENSIONMIB_CieIfVlStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row contains managed objects for Virtual Link statistics on interface
    // capable of  providing this information. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfVlStatsTable_CieIfVlStatsEntry.
    CieIfVlStatsEntry []*CISCOIFEXTENSIONMIB_CieIfVlStatsTable_CieIfVlStatsEntry
}

func (cieIfVlStatsTable *CISCOIFEXTENSIONMIB_CieIfVlStatsTable) GetEntityData() *types.CommonEntityData {
    cieIfVlStatsTable.EntityData.YFilter = cieIfVlStatsTable.YFilter
    cieIfVlStatsTable.EntityData.YangName = "cieIfVlStatsTable"
    cieIfVlStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfVlStatsTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfVlStatsTable.EntityData.SegmentPath = "cieIfVlStatsTable"
    cieIfVlStatsTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfVlStatsTable.EntityData.SegmentPath
    cieIfVlStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfVlStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfVlStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfVlStatsTable.EntityData.Children = types.NewOrderedMap()
    cieIfVlStatsTable.EntityData.Children.Append("cieIfVlStatsEntry", types.YChild{"CieIfVlStatsEntry", nil})
    for i := range cieIfVlStatsTable.CieIfVlStatsEntry {
        cieIfVlStatsTable.EntityData.Children.Append(types.GetSegmentPath(cieIfVlStatsTable.CieIfVlStatsEntry[i]), types.YChild{"CieIfVlStatsEntry", cieIfVlStatsTable.CieIfVlStatsEntry[i]})
    }
    cieIfVlStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfVlStatsTable.EntityData.YListKeys = []string {}

    return &(cieIfVlStatsTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfVlStatsTable_CieIfVlStatsEntry
// Each row contains managed objects for
// Virtual Link statistics on interface capable of 
// providing this information.
type CISCOIFEXTENSIONMIB_CieIfVlStatsTable_CieIfVlStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object indicates the number of input packets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfNoDropVlInPkts interface{}

    // This object indicates the number of input octets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfNoDropVlInOctets interface{}

    // This object indicates the number of output packets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfNoDropVlOutPkts interface{}

    // This object indicates the number of output octets on all No-Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfNoDropVlOutOctets interface{}

    // This object indicates the number of input packets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfDropVlInPkts interface{}

    // This object indicates the number of input octets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfDropVlInOctets interface{}

    // This object indicates the number of output packets on all Drop Virtual
    // Links belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfDropVlOutPkts interface{}

    // This object indicates the number of output octets on all Drop Virtual Links
    // belonged  to this interface. The type is interface{} with range:
    // 0..18446744073709551615.
    CieIfDropVlOutOctets interface{}
}

func (cieIfVlStatsEntry *CISCOIFEXTENSIONMIB_CieIfVlStatsTable_CieIfVlStatsEntry) GetEntityData() *types.CommonEntityData {
    cieIfVlStatsEntry.EntityData.YFilter = cieIfVlStatsEntry.YFilter
    cieIfVlStatsEntry.EntityData.YangName = "cieIfVlStatsEntry"
    cieIfVlStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfVlStatsEntry.EntityData.ParentYangName = "cieIfVlStatsTable"
    cieIfVlStatsEntry.EntityData.SegmentPath = "cieIfVlStatsEntry" + types.AddKeyToken(cieIfVlStatsEntry.IfIndex, "ifIndex")
    cieIfVlStatsEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfVlStatsTable/" + cieIfVlStatsEntry.EntityData.SegmentPath
    cieIfVlStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfVlStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfVlStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfVlStatsEntry.EntityData.Children = types.NewOrderedMap()
    cieIfVlStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfVlStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfVlStatsEntry.IfIndex})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfNoDropVlInPkts", types.YLeaf{"CieIfNoDropVlInPkts", cieIfVlStatsEntry.CieIfNoDropVlInPkts})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfNoDropVlInOctets", types.YLeaf{"CieIfNoDropVlInOctets", cieIfVlStatsEntry.CieIfNoDropVlInOctets})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfNoDropVlOutPkts", types.YLeaf{"CieIfNoDropVlOutPkts", cieIfVlStatsEntry.CieIfNoDropVlOutPkts})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfNoDropVlOutOctets", types.YLeaf{"CieIfNoDropVlOutOctets", cieIfVlStatsEntry.CieIfNoDropVlOutOctets})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfDropVlInPkts", types.YLeaf{"CieIfDropVlInPkts", cieIfVlStatsEntry.CieIfDropVlInPkts})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfDropVlInOctets", types.YLeaf{"CieIfDropVlInOctets", cieIfVlStatsEntry.CieIfDropVlInOctets})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfDropVlOutPkts", types.YLeaf{"CieIfDropVlOutPkts", cieIfVlStatsEntry.CieIfDropVlOutPkts})
    cieIfVlStatsEntry.EntityData.Leafs.Append("cieIfDropVlOutOctets", types.YLeaf{"CieIfDropVlOutOctets", cieIfVlStatsEntry.CieIfDropVlOutOctets})

    cieIfVlStatsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfVlStatsEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable
// This table lists configuration data relating to ifIndex
// persistence.
// 
// This table has a sparse dependent relationship on the ifTable,
// containing a row for each ifEntry corresponding to an interface
// for which ifIndex persistence is supported.
type CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents ifindex persistence configuration for an interface
    // specified by ifIndex. Whenever an interface which supports ifindex
    // persistence is created/destroyed in the ifTable, the corresponding ifindex
    // persistence entry is created/destroyed respectively. Some of the interfaces
    // may not support ifindex persistence, for example, a dynamic interface, such
    // as a PPP connection or a IP subscriber interface. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable_CieIfIndexPersistenceEntry.
    CieIfIndexPersistenceEntry []*CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable_CieIfIndexPersistenceEntry
}

func (cieIfIndexPersistenceTable *CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable) GetEntityData() *types.CommonEntityData {
    cieIfIndexPersistenceTable.EntityData.YFilter = cieIfIndexPersistenceTable.YFilter
    cieIfIndexPersistenceTable.EntityData.YangName = "cieIfIndexPersistenceTable"
    cieIfIndexPersistenceTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfIndexPersistenceTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfIndexPersistenceTable.EntityData.SegmentPath = "cieIfIndexPersistenceTable"
    cieIfIndexPersistenceTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfIndexPersistenceTable.EntityData.SegmentPath
    cieIfIndexPersistenceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfIndexPersistenceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfIndexPersistenceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfIndexPersistenceTable.EntityData.Children = types.NewOrderedMap()
    cieIfIndexPersistenceTable.EntityData.Children.Append("cieIfIndexPersistenceEntry", types.YChild{"CieIfIndexPersistenceEntry", nil})
    for i := range cieIfIndexPersistenceTable.CieIfIndexPersistenceEntry {
        cieIfIndexPersistenceTable.EntityData.Children.Append(types.GetSegmentPath(cieIfIndexPersistenceTable.CieIfIndexPersistenceEntry[i]), types.YChild{"CieIfIndexPersistenceEntry", cieIfIndexPersistenceTable.CieIfIndexPersistenceEntry[i]})
    }
    cieIfIndexPersistenceTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfIndexPersistenceTable.EntityData.YListKeys = []string {}

    return &(cieIfIndexPersistenceTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable_CieIfIndexPersistenceEntry
// Each entry represents ifindex persistence configuration for an
// interface specified by ifIndex. Whenever an interface which
// supports ifindex persistence is created/destroyed in the
// ifTable, the corresponding ifindex persistence entry is
// created/destroyed respectively. Some of the interfaces may not
// support ifindex persistence, for example, a dynamic interface,
// such as a PPP connection or a IP subscriber interface.
type CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable_CieIfIndexPersistenceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object specifies whether the interface's ifIndex value persist across
    // reinitialization.  Due to change in syntax, this object is deprecated by
    // cieIfIndexPersistenceControl. The type is bool.
    CieIfIndexPersistenceEnabled interface{}

    // This object specifies whether the interface's ifIndex value persist across
    // reinitialization. In global state, the interface uses the global setting
    // data for persistence i.e. cieIfIndexGlobalPersistence. The type is
    // IfIndexPersistenceState.
    CieIfIndexPersistenceControl interface{}
}

func (cieIfIndexPersistenceEntry *CISCOIFEXTENSIONMIB_CieIfIndexPersistenceTable_CieIfIndexPersistenceEntry) GetEntityData() *types.CommonEntityData {
    cieIfIndexPersistenceEntry.EntityData.YFilter = cieIfIndexPersistenceEntry.YFilter
    cieIfIndexPersistenceEntry.EntityData.YangName = "cieIfIndexPersistenceEntry"
    cieIfIndexPersistenceEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfIndexPersistenceEntry.EntityData.ParentYangName = "cieIfIndexPersistenceTable"
    cieIfIndexPersistenceEntry.EntityData.SegmentPath = "cieIfIndexPersistenceEntry" + types.AddKeyToken(cieIfIndexPersistenceEntry.IfIndex, "ifIndex")
    cieIfIndexPersistenceEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfIndexPersistenceTable/" + cieIfIndexPersistenceEntry.EntityData.SegmentPath
    cieIfIndexPersistenceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfIndexPersistenceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfIndexPersistenceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfIndexPersistenceEntry.EntityData.Children = types.NewOrderedMap()
    cieIfIndexPersistenceEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfIndexPersistenceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfIndexPersistenceEntry.IfIndex})
    cieIfIndexPersistenceEntry.EntityData.Leafs.Append("cieIfIndexPersistenceEnabled", types.YLeaf{"CieIfIndexPersistenceEnabled", cieIfIndexPersistenceEntry.CieIfIndexPersistenceEnabled})
    cieIfIndexPersistenceEntry.EntityData.Leafs.Append("cieIfIndexPersistenceControl", types.YLeaf{"CieIfIndexPersistenceControl", cieIfIndexPersistenceEntry.CieIfIndexPersistenceControl})

    cieIfIndexPersistenceEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfIndexPersistenceEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable
// A list of the interfaces that support
// the 802.1q custom Ethertype feature.
type CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the custom EtherType information for the interface. 
    // Only interfaces with custom 802.1q ethertype control are listed in the 
    // table. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable_CieIfDot1qCustomEtherTypeEntry.
    CieIfDot1qCustomEtherTypeEntry []*CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable_CieIfDot1qCustomEtherTypeEntry
}

func (cieIfDot1qCustomEtherTypeTable *CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable) GetEntityData() *types.CommonEntityData {
    cieIfDot1qCustomEtherTypeTable.EntityData.YFilter = cieIfDot1qCustomEtherTypeTable.YFilter
    cieIfDot1qCustomEtherTypeTable.EntityData.YangName = "cieIfDot1qCustomEtherTypeTable"
    cieIfDot1qCustomEtherTypeTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfDot1qCustomEtherTypeTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfDot1qCustomEtherTypeTable.EntityData.SegmentPath = "cieIfDot1qCustomEtherTypeTable"
    cieIfDot1qCustomEtherTypeTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfDot1qCustomEtherTypeTable.EntityData.SegmentPath
    cieIfDot1qCustomEtherTypeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfDot1qCustomEtherTypeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfDot1qCustomEtherTypeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfDot1qCustomEtherTypeTable.EntityData.Children = types.NewOrderedMap()
    cieIfDot1qCustomEtherTypeTable.EntityData.Children.Append("cieIfDot1qCustomEtherTypeEntry", types.YChild{"CieIfDot1qCustomEtherTypeEntry", nil})
    for i := range cieIfDot1qCustomEtherTypeTable.CieIfDot1qCustomEtherTypeEntry {
        cieIfDot1qCustomEtherTypeTable.EntityData.Children.Append(types.GetSegmentPath(cieIfDot1qCustomEtherTypeTable.CieIfDot1qCustomEtherTypeEntry[i]), types.YChild{"CieIfDot1qCustomEtherTypeEntry", cieIfDot1qCustomEtherTypeTable.CieIfDot1qCustomEtherTypeEntry[i]})
    }
    cieIfDot1qCustomEtherTypeTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfDot1qCustomEtherTypeTable.EntityData.YListKeys = []string {}

    return &(cieIfDot1qCustomEtherTypeTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable_CieIfDot1qCustomEtherTypeEntry
// An entry containing the custom EtherType
// information for the interface.
// 
// Only interfaces with custom 802.1q
// ethertype control are listed in the 
// table.
type CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable_CieIfDot1qCustomEtherTypeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

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
    CieIfDot1qCustomAdminEtherType interface{}

    // The current operational value of the 802.1q ethertype for the interface.
    // The type is interface{} with range: 0..65535.
    CieIfDot1qCustomOperEtherType interface{}
}

func (cieIfDot1qCustomEtherTypeEntry *CISCOIFEXTENSIONMIB_CieIfDot1qCustomEtherTypeTable_CieIfDot1qCustomEtherTypeEntry) GetEntityData() *types.CommonEntityData {
    cieIfDot1qCustomEtherTypeEntry.EntityData.YFilter = cieIfDot1qCustomEtherTypeEntry.YFilter
    cieIfDot1qCustomEtherTypeEntry.EntityData.YangName = "cieIfDot1qCustomEtherTypeEntry"
    cieIfDot1qCustomEtherTypeEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfDot1qCustomEtherTypeEntry.EntityData.ParentYangName = "cieIfDot1qCustomEtherTypeTable"
    cieIfDot1qCustomEtherTypeEntry.EntityData.SegmentPath = "cieIfDot1qCustomEtherTypeEntry" + types.AddKeyToken(cieIfDot1qCustomEtherTypeEntry.IfIndex, "ifIndex")
    cieIfDot1qCustomEtherTypeEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfDot1qCustomEtherTypeTable/" + cieIfDot1qCustomEtherTypeEntry.EntityData.SegmentPath
    cieIfDot1qCustomEtherTypeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfDot1qCustomEtherTypeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfDot1qCustomEtherTypeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfDot1qCustomEtherTypeEntry.EntityData.Children = types.NewOrderedMap()
    cieIfDot1qCustomEtherTypeEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfDot1qCustomEtherTypeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfDot1qCustomEtherTypeEntry.IfIndex})
    cieIfDot1qCustomEtherTypeEntry.EntityData.Leafs.Append("cieIfDot1qCustomAdminEtherType", types.YLeaf{"CieIfDot1qCustomAdminEtherType", cieIfDot1qCustomEtherTypeEntry.CieIfDot1qCustomAdminEtherType})
    cieIfDot1qCustomEtherTypeEntry.EntityData.Leafs.Append("cieIfDot1qCustomOperEtherType", types.YLeaf{"CieIfDot1qCustomOperEtherType", cieIfDot1qCustomEtherTypeEntry.CieIfDot1qCustomOperEtherType})

    cieIfDot1qCustomEtherTypeEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfDot1qCustomEtherTypeEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfUtilTable
// This table contains the interface utilization
// rates for inbound and outbound traffic on an
// interface.
type CISCOIFEXTENSIONMIB_CieIfUtilTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing utilization rates for the interface.  Every interface
    // for which the  inbound and  outbound traffic information is available has a
    // corresponding entry in this table. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfUtilTable_CieIfUtilEntry.
    CieIfUtilEntry []*CISCOIFEXTENSIONMIB_CieIfUtilTable_CieIfUtilEntry
}

func (cieIfUtilTable *CISCOIFEXTENSIONMIB_CieIfUtilTable) GetEntityData() *types.CommonEntityData {
    cieIfUtilTable.EntityData.YFilter = cieIfUtilTable.YFilter
    cieIfUtilTable.EntityData.YangName = "cieIfUtilTable"
    cieIfUtilTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfUtilTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfUtilTable.EntityData.SegmentPath = "cieIfUtilTable"
    cieIfUtilTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfUtilTable.EntityData.SegmentPath
    cieIfUtilTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfUtilTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfUtilTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfUtilTable.EntityData.Children = types.NewOrderedMap()
    cieIfUtilTable.EntityData.Children.Append("cieIfUtilEntry", types.YChild{"CieIfUtilEntry", nil})
    for i := range cieIfUtilTable.CieIfUtilEntry {
        cieIfUtilTable.EntityData.Children.Append(types.GetSegmentPath(cieIfUtilTable.CieIfUtilEntry[i]), types.YChild{"CieIfUtilEntry", cieIfUtilTable.CieIfUtilEntry[i]})
    }
    cieIfUtilTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfUtilTable.EntityData.YListKeys = []string {}

    return &(cieIfUtilTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfUtilTable_CieIfUtilEntry
// An entry containing utilization rates for the
// interface.
// 
// Every interface for which the  inbound and 
// outbound traffic information is available
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_CieIfUtilTable_CieIfUtilEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the inbound packet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfInPktRate is the
    // exponentially-decayed moving average of inbound packet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    CieIfInPktRate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the inbound octet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfInOctetRate is the
    // exponentially-decayed moving average of inbound octet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are octets per second.
    CieIfInOctetRate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the outbound packet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfOutPktRate is the
    // exponentially-decayed moving average of outbound packet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    CieIfOutPktRate interface{}

    // By default, this is the five minute exponentially-decayed moving average of
    // the outbound octet rate for this interface. However, if the corresponding
    // instance of cieIfInterval is instantiated with a value which specifies an
    // interval different from 5-minutes, then cieIfOutOctetRate is the
    // exponentially-decayed moving average of outbound octet rate over this
    // different time interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are octets per second.
    CieIfOutOctetRate interface{}

    // This object specifies the time interval over which the inbound and outbound
    // traffic rates are calculated for this interface. The type is interface{}
    // with range: 1..4294967295. Units are seconds.
    CieIfInterval interface{}
}

func (cieIfUtilEntry *CISCOIFEXTENSIONMIB_CieIfUtilTable_CieIfUtilEntry) GetEntityData() *types.CommonEntityData {
    cieIfUtilEntry.EntityData.YFilter = cieIfUtilEntry.YFilter
    cieIfUtilEntry.EntityData.YangName = "cieIfUtilEntry"
    cieIfUtilEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfUtilEntry.EntityData.ParentYangName = "cieIfUtilTable"
    cieIfUtilEntry.EntityData.SegmentPath = "cieIfUtilEntry" + types.AddKeyToken(cieIfUtilEntry.IfIndex, "ifIndex")
    cieIfUtilEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfUtilTable/" + cieIfUtilEntry.EntityData.SegmentPath
    cieIfUtilEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfUtilEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfUtilEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfUtilEntry.EntityData.Children = types.NewOrderedMap()
    cieIfUtilEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfUtilEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfUtilEntry.IfIndex})
    cieIfUtilEntry.EntityData.Leafs.Append("cieIfInPktRate", types.YLeaf{"CieIfInPktRate", cieIfUtilEntry.CieIfInPktRate})
    cieIfUtilEntry.EntityData.Leafs.Append("cieIfInOctetRate", types.YLeaf{"CieIfInOctetRate", cieIfUtilEntry.CieIfInOctetRate})
    cieIfUtilEntry.EntityData.Leafs.Append("cieIfOutPktRate", types.YLeaf{"CieIfOutPktRate", cieIfUtilEntry.CieIfOutPktRate})
    cieIfUtilEntry.EntityData.Leafs.Append("cieIfOutOctetRate", types.YLeaf{"CieIfOutOctetRate", cieIfUtilEntry.CieIfOutOctetRate})
    cieIfUtilEntry.EntityData.Leafs.Append("cieIfInterval", types.YLeaf{"CieIfInterval", cieIfUtilEntry.CieIfInterval})

    cieIfUtilEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfUtilEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable
// This table contains the mappings of the
// ifIndex of an interface to its
// corresponding dot1dBasePort value.
type CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the mapping between the ifIndex value of an interface
    // and its corresponding dot1dBasePort value.  Every interface which has been
    // assigned a dot1dBasePort value by the system has a corresponding entry in
    // this table. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable_CieIfDot1dBaseMappingEntry.
    CieIfDot1dBaseMappingEntry []*CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable_CieIfDot1dBaseMappingEntry
}

func (cieIfDot1dBaseMappingTable *CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable) GetEntityData() *types.CommonEntityData {
    cieIfDot1dBaseMappingTable.EntityData.YFilter = cieIfDot1dBaseMappingTable.YFilter
    cieIfDot1dBaseMappingTable.EntityData.YangName = "cieIfDot1dBaseMappingTable"
    cieIfDot1dBaseMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfDot1dBaseMappingTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfDot1dBaseMappingTable.EntityData.SegmentPath = "cieIfDot1dBaseMappingTable"
    cieIfDot1dBaseMappingTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfDot1dBaseMappingTable.EntityData.SegmentPath
    cieIfDot1dBaseMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfDot1dBaseMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfDot1dBaseMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfDot1dBaseMappingTable.EntityData.Children = types.NewOrderedMap()
    cieIfDot1dBaseMappingTable.EntityData.Children.Append("cieIfDot1dBaseMappingEntry", types.YChild{"CieIfDot1dBaseMappingEntry", nil})
    for i := range cieIfDot1dBaseMappingTable.CieIfDot1dBaseMappingEntry {
        cieIfDot1dBaseMappingTable.EntityData.Children.Append(types.GetSegmentPath(cieIfDot1dBaseMappingTable.CieIfDot1dBaseMappingEntry[i]), types.YChild{"CieIfDot1dBaseMappingEntry", cieIfDot1dBaseMappingTable.CieIfDot1dBaseMappingEntry[i]})
    }
    cieIfDot1dBaseMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfDot1dBaseMappingTable.EntityData.YListKeys = []string {}

    return &(cieIfDot1dBaseMappingTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable_CieIfDot1dBaseMappingEntry
// An entry containing the mapping between
// the ifIndex value of an interface and its
// corresponding dot1dBasePort value.
// 
// Every interface which has been assigned
// a dot1dBasePort value by the system
// has a corresponding entry in this table.
type CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable_CieIfDot1dBaseMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The dot1dBasePort value for this interface. The type is interface{} with
    // range: 1..65535.
    CieIfDot1dBaseMappingPort interface{}
}

func (cieIfDot1dBaseMappingEntry *CISCOIFEXTENSIONMIB_CieIfDot1dBaseMappingTable_CieIfDot1dBaseMappingEntry) GetEntityData() *types.CommonEntityData {
    cieIfDot1dBaseMappingEntry.EntityData.YFilter = cieIfDot1dBaseMappingEntry.YFilter
    cieIfDot1dBaseMappingEntry.EntityData.YangName = "cieIfDot1dBaseMappingEntry"
    cieIfDot1dBaseMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfDot1dBaseMappingEntry.EntityData.ParentYangName = "cieIfDot1dBaseMappingTable"
    cieIfDot1dBaseMappingEntry.EntityData.SegmentPath = "cieIfDot1dBaseMappingEntry" + types.AddKeyToken(cieIfDot1dBaseMappingEntry.IfIndex, "ifIndex")
    cieIfDot1dBaseMappingEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfDot1dBaseMappingTable/" + cieIfDot1dBaseMappingEntry.EntityData.SegmentPath
    cieIfDot1dBaseMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfDot1dBaseMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfDot1dBaseMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfDot1dBaseMappingEntry.EntityData.Children = types.NewOrderedMap()
    cieIfDot1dBaseMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfDot1dBaseMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cieIfDot1dBaseMappingEntry.IfIndex})
    cieIfDot1dBaseMappingEntry.EntityData.Leafs.Append("cieIfDot1dBaseMappingPort", types.YLeaf{"CieIfDot1dBaseMappingPort", cieIfDot1dBaseMappingEntry.CieIfDot1dBaseMappingPort})

    cieIfDot1dBaseMappingEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cieIfDot1dBaseMappingEntry.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfNameMappingTable
// This table contains objects for providing
// the 'ifName' to 'ifIndex' mapping.
// This table contains one entry for each
// valid 'ifName' available in the system.
// Upon the first request, the implementation
// of this table will get all the available
// ifNames, and it will populate the entries
// in this table, it maintains this ifNames
// in a cache for ~30 seconds.
type CISCOIFEXTENSIONMIB_CieIfNameMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry into the cieIfNameMappingTable. The type is slice of
    // CISCOIFEXTENSIONMIB_CieIfNameMappingTable_CieIfNameMappingEntry.
    CieIfNameMappingEntry []*CISCOIFEXTENSIONMIB_CieIfNameMappingTable_CieIfNameMappingEntry
}

func (cieIfNameMappingTable *CISCOIFEXTENSIONMIB_CieIfNameMappingTable) GetEntityData() *types.CommonEntityData {
    cieIfNameMappingTable.EntityData.YFilter = cieIfNameMappingTable.YFilter
    cieIfNameMappingTable.EntityData.YangName = "cieIfNameMappingTable"
    cieIfNameMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cieIfNameMappingTable.EntityData.ParentYangName = "CISCO-IF-EXTENSION-MIB"
    cieIfNameMappingTable.EntityData.SegmentPath = "cieIfNameMappingTable"
    cieIfNameMappingTable.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/" + cieIfNameMappingTable.EntityData.SegmentPath
    cieIfNameMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfNameMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfNameMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfNameMappingTable.EntityData.Children = types.NewOrderedMap()
    cieIfNameMappingTable.EntityData.Children.Append("cieIfNameMappingEntry", types.YChild{"CieIfNameMappingEntry", nil})
    for i := range cieIfNameMappingTable.CieIfNameMappingEntry {
        cieIfNameMappingTable.EntityData.Children.Append(types.GetSegmentPath(cieIfNameMappingTable.CieIfNameMappingEntry[i]), types.YChild{"CieIfNameMappingEntry", cieIfNameMappingTable.CieIfNameMappingEntry[i]})
    }
    cieIfNameMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cieIfNameMappingTable.EntityData.YListKeys = []string {}

    return &(cieIfNameMappingTable.EntityData)
}

// CISCOIFEXTENSIONMIB_CieIfNameMappingTable_CieIfNameMappingEntry
// An entry into the cieIfNameMappingTable.
type CISCOIFEXTENSIONMIB_CieIfNameMappingTable_CieIfNameMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Represents an interface name mentioned in the
    // 'ifName' object of this system. The type is string with length: 1..112.
    CieIfName interface{}

    // This object represents the 'ifIndex' corresponding to the interface name
    // mentioned in the 'cieIfName' object of this instance. If the 'ifName'
    // mentioned in the 'cieIfName'  object of this instance corresponds to
    // multiple 'ifIndex' values, then the value of this object is the numerically
    // smallest of those multiple  'ifIndex' values. The type is interface{} with
    // range: 0..2147483647.
    CieIfIndex interface{}
}

func (cieIfNameMappingEntry *CISCOIFEXTENSIONMIB_CieIfNameMappingTable_CieIfNameMappingEntry) GetEntityData() *types.CommonEntityData {
    cieIfNameMappingEntry.EntityData.YFilter = cieIfNameMappingEntry.YFilter
    cieIfNameMappingEntry.EntityData.YangName = "cieIfNameMappingEntry"
    cieIfNameMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cieIfNameMappingEntry.EntityData.ParentYangName = "cieIfNameMappingTable"
    cieIfNameMappingEntry.EntityData.SegmentPath = "cieIfNameMappingEntry" + types.AddKeyToken(cieIfNameMappingEntry.CieIfName, "cieIfName")
    cieIfNameMappingEntry.EntityData.AbsolutePath = "CISCO-IF-EXTENSION-MIB:CISCO-IF-EXTENSION-MIB/cieIfNameMappingTable/" + cieIfNameMappingEntry.EntityData.SegmentPath
    cieIfNameMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cieIfNameMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cieIfNameMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cieIfNameMappingEntry.EntityData.Children = types.NewOrderedMap()
    cieIfNameMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cieIfNameMappingEntry.EntityData.Leafs.Append("cieIfName", types.YLeaf{"CieIfName", cieIfNameMappingEntry.CieIfName})
    cieIfNameMappingEntry.EntityData.Leafs.Append("cieIfIndex", types.YLeaf{"CieIfIndex", cieIfNameMappingEntry.CieIfIndex})

    cieIfNameMappingEntry.EntityData.YListKeys = []string {"CieIfName"}

    return &(cieIfNameMappingEntry.EntityData)
}

