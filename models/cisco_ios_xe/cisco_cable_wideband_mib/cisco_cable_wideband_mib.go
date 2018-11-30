// This is the MIB module for the support of Channel Bonding
// Protocol for the Cable Modem Termination System (CMTS).
// 
// Wideband DOCSIS is a method of increasing downstream
// bandwidth by simultaneously transmitting DOCSIS data
// over multiple RF channels. This DOCSIS data is
// organized as a sequence of 188-byte MPEG-TS data packets.
// 
// A Wideband CMTS (WCMTS) is a CMTS that also transmits
// Wideband MPEG-TS packets.
// 
// An Edge QAM (Quadrature Amplitude Modulation) device is one
// which provides the QAM and used to couple the Wideband MPEG-TS
// packet onto the HFC plant.
// 
// A Wideband Cable Modem (WCM) is a CableModem (CM) that
// is able to receive Wideband MPEG-TS packets.
// 
// A wideband channel or Bonded Group (BG) is a logical
// grouping of one or more physical RF channels over which
// MPEG-TS packets are carried. Wideband channel carries DOCSIS
// bonded packets encapsulated in MPEG-TS packets from a
// WCMTS to one or more WCMs.
// 
// Packets outgoing from WCMTS to the WCM are formatted with
// the DOCSIS Header. The DOCSIS packets are then formatted into
// MPEG-TS data packets. These are 188 byte MPEG packets
// containing the DOCSIS information.
// Within DOCSIS Header of the WB Channel there is an extended
// header called, DOCSIS Bonding Extended Header, the format of
// which is shown below:
// 
//  --------------------------------------------------------
// |  ----------------------------------------------------  |
// | |  TYPE | LEN |       BSID           |      SEQ        |
// |  ----------------------------------------------------  |
// | |    byte 0   |     byte 1-2         |    byte 3-4     |
// |--------------------------------------------------------|
//               DOCSIS Bonding Extended Header
// 
// BSID is the Bonding Service IDentifier, it defines a sequence
// number for a Wideband channel. It is used by the WCM to
// re-sequence packets received over downstream channels to
// maintain packet order. SEQ is per service flow sequence number.
// Whereas TYPE is the type of the Bonding Extended Header and LEN
// specifies its length.
// 
// A Narrowband Channel is a standard DOCSIS downstream channel
// which contains exactly one RF channel.
// The wideband protocol utilizes the existing narrowband
// downstream channel for carrying the MAC management and
// signaling messages and the associated narrowband upstream
// for return data traffic and signaling.
// 
// The channel bonding protocol supports multiple wideband
// bonded groups. This will allow the WCM to listen to multiple
// bonded groups at the same time. This would support (for example)
// multicast video being sent to a CPE device on the LAN side of
// the WCM in addition to standard DOCSIS data.
// 
// Channel Bonding allows two types of Bonding Group (BG) interfaces.
// These are Secondary BG interface and non-Secondary BG
// interface. The Secondary BG interfaces will carry the multicast
// traffic, whereas, the non-Secondary BG interface will carry the
// non-multicast traffic.
// This MIB also allows for configuration of the RF channels
// on the WCMTS, as well as the association between the RF and
// narrowband downstream channels with the BG channel.
// 
// Fiber Node is an optical node which terminates the fiber based
// downstream signal as an electrical signal onto a coaxial RF cable.
// 
// DEPI:  Downstream external physical interface. 
// TSID:  MPEG2 Transport Stream ID.
// SFP:   Small Form-Factor Pluggable.
package cisco_cable_wideband_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cable_wideband_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CABLE-WIDEBAND-MIB CISCO-CABLE-WIDEBAND-MIB}", reflect.TypeOf(CISCOCABLEWIDEBANDMIB{}))
    ydk.RegisterEntity("CISCO-CABLE-WIDEBAND-MIB:CISCO-CABLE-WIDEBAND-MIB", reflect.TypeOf(CISCOCABLEWIDEBANDMIB{}))
}

// CISCOCABLEWIDEBANDMIB
type CISCOCABLEWIDEBANDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoCableWidebandMIBObjects CISCOCABLEWIDEBANDMIB_CiscoCableWidebandMIBObjects

    // This table contains attributes of the physical RF channels. MPEG-TS packets
    // are sent across RF channels within a wideband channel.  These physical RF
    // channels might be present on a different system but the WCMTS entity
    // requires the knowledge of that system for its operation.
    CcwbRFChannelTable CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable

    // A wideband channel is a logical grouping of one or more physical RF
    // channels over which Wideband MPEG-TS packets are carried.  This table
    // contains association information of the wideband channels to the RF
    // channels that are available for the WCMTS.
    CcwbWBtoRFMappingTable CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable

    // This table contains information of the mapping of the wideband channels to
    // the Narrowband channels that are available on the WCMTS.  The wideband
    // protocol utilizes the existing narrowband downstream channel for carrying
    // the MAC management and signaling messages and the associated narrowband
    // upstream for return data traffic and signaling.
    CcwbWBtoNBMappingTable CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable

    // This table provides information about a Wideband BG interface, whether its
    // configured to carry multicast or non-multicast traffic. For multicast the
    // BG interface type is Secondary and for non-multicast its non-Secondary.
    // Other objects could be added to this later in the future.
    CcwbWBBondGrpTable CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable

    // This table contains Wideband Cable Modem(WCM) connectivity state. A WCM
    // connectivity state can be associated with multiple Wideband interfaces.
    CcwbWBCmStatusTable CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable

    // An entry in this table exists for each Wideband Cable Modem which links to
    // one or more WB interface.
    CcwbWBCmStatusExtTable CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable

    // This table contains the description of a Fiber Node on a CMTS. An entry in
    // this table exists for each FiberNode ID.
    CcwbFiberNodeDescrTable CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable

    // This table provides configuration information of each Fiber node. It will
    // provide topology information of each Fiber node, which includes information
    // such as, Narrowband and Wideband QAMs.
    CcwbFiberNodeTable CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable
}

func (cISCOCABLEWIDEBANDMIB *CISCOCABLEWIDEBANDMIB) GetEntityData() *types.CommonEntityData {
    cISCOCABLEWIDEBANDMIB.EntityData.YFilter = cISCOCABLEWIDEBANDMIB.YFilter
    cISCOCABLEWIDEBANDMIB.EntityData.YangName = "CISCO-CABLE-WIDEBAND-MIB"
    cISCOCABLEWIDEBANDMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCABLEWIDEBANDMIB.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    cISCOCABLEWIDEBANDMIB.EntityData.SegmentPath = "CISCO-CABLE-WIDEBAND-MIB:CISCO-CABLE-WIDEBAND-MIB"
    cISCOCABLEWIDEBANDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCABLEWIDEBANDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCABLEWIDEBANDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCABLEWIDEBANDMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ciscoCableWidebandMIBObjects", types.YChild{"CiscoCableWidebandMIBObjects", &cISCOCABLEWIDEBANDMIB.CiscoCableWidebandMIBObjects})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbRFChannelTable", types.YChild{"CcwbRFChannelTable", &cISCOCABLEWIDEBANDMIB.CcwbRFChannelTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbWBtoRFMappingTable", types.YChild{"CcwbWBtoRFMappingTable", &cISCOCABLEWIDEBANDMIB.CcwbWBtoRFMappingTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbWBtoNBMappingTable", types.YChild{"CcwbWBtoNBMappingTable", &cISCOCABLEWIDEBANDMIB.CcwbWBtoNBMappingTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbWBBondGrpTable", types.YChild{"CcwbWBBondGrpTable", &cISCOCABLEWIDEBANDMIB.CcwbWBBondGrpTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbWBCmStatusTable", types.YChild{"CcwbWBCmStatusTable", &cISCOCABLEWIDEBANDMIB.CcwbWBCmStatusTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbWBCmStatusExtTable", types.YChild{"CcwbWBCmStatusExtTable", &cISCOCABLEWIDEBANDMIB.CcwbWBCmStatusExtTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbFiberNodeDescrTable", types.YChild{"CcwbFiberNodeDescrTable", &cISCOCABLEWIDEBANDMIB.CcwbFiberNodeDescrTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Children.Append("ccwbFiberNodeTable", types.YChild{"CcwbFiberNodeTable", &cISCOCABLEWIDEBANDMIB.CcwbFiberNodeTable})
    cISCOCABLEWIDEBANDMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCABLEWIDEBANDMIB.EntityData.YListKeys = []string {}

    return &(cISCOCABLEWIDEBANDMIB.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CiscoCableWidebandMIBObjects
type CISCOCABLEWIDEBANDMIB_CiscoCableWidebandMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time interval in seconds over which the RF channels utilization index
    // is calculated. All RF channels use the same interval.  Setting a value of
    // zero disables utilization reporting.  This value should be persisted
    // accross CMTS reinitializations. The type is interface{} with range:
    // 0..86400. Units are seconds.
    CcwbRFChanUtilInterval interface{}

    // This object specifies whether ccwbSFPLinkDownNotification and
    // ccwbSFPLinkUpNotification are generated. The type is interface{} with
    // range: 0..1.
    CcwbSFPLinkTrapEnable interface{}
}

func (ciscoCableWidebandMIBObjects *CISCOCABLEWIDEBANDMIB_CiscoCableWidebandMIBObjects) GetEntityData() *types.CommonEntityData {
    ciscoCableWidebandMIBObjects.EntityData.YFilter = ciscoCableWidebandMIBObjects.YFilter
    ciscoCableWidebandMIBObjects.EntityData.YangName = "ciscoCableWidebandMIBObjects"
    ciscoCableWidebandMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoCableWidebandMIBObjects.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ciscoCableWidebandMIBObjects.EntityData.SegmentPath = "ciscoCableWidebandMIBObjects"
    ciscoCableWidebandMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoCableWidebandMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoCableWidebandMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoCableWidebandMIBObjects.EntityData.Children = types.NewOrderedMap()
    ciscoCableWidebandMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoCableWidebandMIBObjects.EntityData.Leafs.Append("ccwbRFChanUtilInterval", types.YLeaf{"CcwbRFChanUtilInterval", ciscoCableWidebandMIBObjects.CcwbRFChanUtilInterval})
    ciscoCableWidebandMIBObjects.EntityData.Leafs.Append("ccwbSFPLinkTrapEnable", types.YLeaf{"CcwbSFPLinkTrapEnable", ciscoCableWidebandMIBObjects.CcwbSFPLinkTrapEnable})

    ciscoCableWidebandMIBObjects.EntityData.YListKeys = []string {}

    return &(ciscoCableWidebandMIBObjects.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable
// This table contains attributes of the physical
// RF channels. MPEG-TS packets are sent across RF
// channels within a wideband channel.
// 
// These physical RF channels might be present on a
// different system but the WCMTS entity requires
// the knowledge of that system for its operation.
type CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides a list of attributes for a single downstream RF channel
    // per WCMTS entity.  An entry in this table exists for each configured RF
    // channel on the WCMTS entity that provides the wideband DOCSIS
    // functionality. The index, entPhysicalIndex, used here is the physical index
    // of the wideband controller card. Since RF channels are considered part of
    // the Wideband controller card, hence entPhysicalIndex is used for
    // associating RF channels. The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry.
    CcwbRFChannelEntry []*CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry
}

func (ccwbRFChannelTable *CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable) GetEntityData() *types.CommonEntityData {
    ccwbRFChannelTable.EntityData.YFilter = ccwbRFChannelTable.YFilter
    ccwbRFChannelTable.EntityData.YangName = "ccwbRFChannelTable"
    ccwbRFChannelTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbRFChannelTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbRFChannelTable.EntityData.SegmentPath = "ccwbRFChannelTable"
    ccwbRFChannelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbRFChannelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbRFChannelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbRFChannelTable.EntityData.Children = types.NewOrderedMap()
    ccwbRFChannelTable.EntityData.Children.Append("ccwbRFChannelEntry", types.YChild{"CcwbRFChannelEntry", nil})
    for i := range ccwbRFChannelTable.CcwbRFChannelEntry {
        ccwbRFChannelTable.EntityData.Children.Append(types.GetSegmentPath(ccwbRFChannelTable.CcwbRFChannelEntry[i]), types.YChild{"CcwbRFChannelEntry", ccwbRFChannelTable.CcwbRFChannelEntry[i]})
    }
    ccwbRFChannelTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbRFChannelTable.EntityData.YListKeys = []string {}

    return &(ccwbRFChannelTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry
// An entry provides a list of attributes for a single
// downstream RF channel per WCMTS entity.
// 
// An entry in this table exists for each configured
// RF channel on the WCMTS entity that provides the
// wideband DOCSIS functionality. The index, entPhysicalIndex,
// used here is the physical index of the wideband controller
// card. Since RF channels are considered part of the Wideband
// controller card, hence entPhysicalIndex is used for
// associating RF channels.
type CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The WCMTS identification of the RF channel. The
    // range of this object is limited to 0-18 in the case of annex A/256qam, and
    // 0-23 for Annex B and C. The type is interface{} with range: 0..255.
    CcwbRFChannelNum interface{}

    // The center of the downstream frequency associated with this RF channel. The
    // final downstream RF frequency may be provided by an edge QAM device or the
    // CMTS itself. See the associated compliance object for a description of
    // valid frequencies that may be written to this object.  If the downstream
    // frequency associated with this RF channel is greater than the maximum value
    // reportable by this object then this object should report its maximum value
    // (1,000,000,000) and ccwbRFChannelFrequencyRev1 must be used to report the
    // downstream frequency.  This object is deprecated and replaced by
    // ccwbRFChannelFrequencyRev1. The type is interface{} with range:
    // 0..1000000000. Units are hertz.
    CcwbRFChannelFrequency interface{}

    // The bandwidth of this downstream RF channel. Most implementations are
    // expected to support a channel width of 6 MHz (North America) and/or 8 MHz
    // (Europe). See the associated compliance object for a description of the
    // valid channel widths for this object. The type is interface{} with range:
    // 0..16000000. Units are hertz.
    CcwbRFChannelWidth interface{}

    // The modulation type associated with this downstream RF channel. See the
    // associated conformance object for write conditions an limitations. See the
    // DOCSIS specification for specifics on the modulation profiles implied by
    // qam64 qam256 and qam1024. qam64, qam256 and qam1024 are various modulation
    // schemes often used in digital cable and cable modem applications. The
    // numbers 64/256/1024 in qam represent constellation points, which is the
    // measurement of qam transmission capability, the higher the number, higher
    // the bits that can be transmitted. The type is CcwbRFChannelModulation.
    CcwbRFChannelModulation interface{}

    // The value of this object indicates the conformance of the implementation to
    // important regional cable standards. annexA : Annex A from ITU-J83 is used.
    // annexB : Annex B from ITU-J83 is used. annexC : Annex C from ITU-J83 is
    // used. The type is CcwbRFChannelAnnex.
    CcwbRFChannelAnnex interface{}

    // The number of MPEG packets transmitted on this RF channel. The type is
    // interface{} with range: 0..18446744073709551615. Units are Packets.
    CcwbRFChannelMpegPkts interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcwbRFChannelStorageType interface{}

    // Controls and reflects the status of rows in this table. It can be used for
    // creating and deleting entries in this table.  The ccwbRFChannelAnnex and
    // ccwbRFChannelModulation must be valid for a row to be created. The type is
    // RowStatus.
    CcwbRFChannelRowStatus interface{}

    // The calculated and truncated utilization for this RF channel over the
    // previous complete measuring interval. The configured duration of the
    // measurement intervals is defined in the ccwbRFChanUtilInterval object.  The
    // utilization index is a percentage expressing the ratio between bytes used
    // to transmit data versus the total number of bytes transmitted in the raw
    // bandwidth of the RF channel. The type is interface{} with range: 0..100.
    // Units are percent.
    CcwbRFChannelUtilization interface{}

    // The center of the downstream frequency associated with this RF channel. The
    // final downstream RF frequency may be provided by an edge QAM device or the
    // CMTS itself. See the associated compliance object for a description of
    // valid frequencies that may be written to this object. The type is
    // interface{} with range: 55000000..1050000000. Units are hertz.
    CcwbRFChannelFrequencyRev1 interface{}

    // The type of internet address. This object identifies the internet address
    // type specified by ccwbRFChanQamIPAddress object. The type is
    // InetAddressType.
    CcwbRFChanQamIPAddressType interface{}

    // The IP address of the edge QAM device or the CMTS cable interface which
    // provides the physical RF channel. The IP address will be of the type
    // represented by object ccwbRFChanQamIPAddressType. The type is string with
    // length: 0..255.
    CcwbRFChanQamIPAddress interface{}

    // The MAC address of the edge QAM device or next hop router which might be
    // present between the WCMTS and the edge QAM. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CcwbRFChanQamMacAddress interface{}

    // The port number on the edge QAM associated with this RF channel. The type
    // is interface{} with range: 0..65535.
    CcwbRFChanQamUdpPort interface{}

    // The value of the TOS field in the IP header for all Ethernet frames
    // destined for the given RF channel. The type is interface{} with range:
    // 0..255.
    CcwbRFChanQamTos interface{}

    // The VLAN ID to be inserted in the Ethernet frames when using 802.1q frames
    // instead of normal 802.1 frames for the given RF channel. The value of 0
    // indicates that 802.1 frames are being used and is not supported in setting
    // this object. The type is interface{} with range: 0..4095.
    CcwbRFChanQamVlanId interface{}

    // The priority bits used when inserting Ethernet 802.1q VLAN tags into the
    // Ethernet frames destined for a given RF channel. Priority Bits (0=Best
    // effort, 1=background, 2=spare, 3=excellent effort, 4=controlled load,
    // 5=video, 6=voice, 7=network control). The type is interface{} with range:
    // 0..7.
    CcwbRFChanQamPriorityBits interface{}

    // The DEPI remote ID on edge QAM associated with this RF channel. The type is
    // interface{} with range: 0..4294967295.
    CcwbRFChanQamDepiRemoteId interface{}

    // This object specifies the name of the DEPI tunnel which determines the DEPI
    // data session configuration associated with this RF channel. The type is
    // string with length: 0..255.
    CcwbRFChanQamDepiTunnel interface{}

    // This object specifies the MPEG2 transport stream ID which is associated
    // with this RF channel. The type is interface{} with range: 0..65535.
    CcwbRFChanQamTsid interface{}
}

func (ccwbRFChannelEntry *CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry) GetEntityData() *types.CommonEntityData {
    ccwbRFChannelEntry.EntityData.YFilter = ccwbRFChannelEntry.YFilter
    ccwbRFChannelEntry.EntityData.YangName = "ccwbRFChannelEntry"
    ccwbRFChannelEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbRFChannelEntry.EntityData.ParentYangName = "ccwbRFChannelTable"
    ccwbRFChannelEntry.EntityData.SegmentPath = "ccwbRFChannelEntry" + types.AddKeyToken(ccwbRFChannelEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(ccwbRFChannelEntry.CcwbRFChannelNum, "ccwbRFChannelNum")
    ccwbRFChannelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbRFChannelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbRFChannelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbRFChannelEntry.EntityData.Children = types.NewOrderedMap()
    ccwbRFChannelEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbRFChannelEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ccwbRFChannelEntry.EntPhysicalIndex})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelNum", types.YLeaf{"CcwbRFChannelNum", ccwbRFChannelEntry.CcwbRFChannelNum})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelFrequency", types.YLeaf{"CcwbRFChannelFrequency", ccwbRFChannelEntry.CcwbRFChannelFrequency})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelWidth", types.YLeaf{"CcwbRFChannelWidth", ccwbRFChannelEntry.CcwbRFChannelWidth})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelModulation", types.YLeaf{"CcwbRFChannelModulation", ccwbRFChannelEntry.CcwbRFChannelModulation})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelAnnex", types.YLeaf{"CcwbRFChannelAnnex", ccwbRFChannelEntry.CcwbRFChannelAnnex})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelMpegPkts", types.YLeaf{"CcwbRFChannelMpegPkts", ccwbRFChannelEntry.CcwbRFChannelMpegPkts})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelStorageType", types.YLeaf{"CcwbRFChannelStorageType", ccwbRFChannelEntry.CcwbRFChannelStorageType})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelRowStatus", types.YLeaf{"CcwbRFChannelRowStatus", ccwbRFChannelEntry.CcwbRFChannelRowStatus})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelUtilization", types.YLeaf{"CcwbRFChannelUtilization", ccwbRFChannelEntry.CcwbRFChannelUtilization})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChannelFrequencyRev1", types.YLeaf{"CcwbRFChannelFrequencyRev1", ccwbRFChannelEntry.CcwbRFChannelFrequencyRev1})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamIPAddressType", types.YLeaf{"CcwbRFChanQamIPAddressType", ccwbRFChannelEntry.CcwbRFChanQamIPAddressType})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamIPAddress", types.YLeaf{"CcwbRFChanQamIPAddress", ccwbRFChannelEntry.CcwbRFChanQamIPAddress})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamMacAddress", types.YLeaf{"CcwbRFChanQamMacAddress", ccwbRFChannelEntry.CcwbRFChanQamMacAddress})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamUdpPort", types.YLeaf{"CcwbRFChanQamUdpPort", ccwbRFChannelEntry.CcwbRFChanQamUdpPort})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamTos", types.YLeaf{"CcwbRFChanQamTos", ccwbRFChannelEntry.CcwbRFChanQamTos})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamVlanId", types.YLeaf{"CcwbRFChanQamVlanId", ccwbRFChannelEntry.CcwbRFChanQamVlanId})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamPriorityBits", types.YLeaf{"CcwbRFChanQamPriorityBits", ccwbRFChannelEntry.CcwbRFChanQamPriorityBits})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamDepiRemoteId", types.YLeaf{"CcwbRFChanQamDepiRemoteId", ccwbRFChannelEntry.CcwbRFChanQamDepiRemoteId})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamDepiTunnel", types.YLeaf{"CcwbRFChanQamDepiTunnel", ccwbRFChannelEntry.CcwbRFChanQamDepiTunnel})
    ccwbRFChannelEntry.EntityData.Leafs.Append("ccwbRFChanQamTsid", types.YLeaf{"CcwbRFChanQamTsid", ccwbRFChannelEntry.CcwbRFChanQamTsid})

    ccwbRFChannelEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CcwbRFChannelNum"}

    return &(ccwbRFChannelEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex represents annexC : Annex C from ITU-J83 is used.
type CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex string

const (
    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex_annexA CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex = "annexA"

    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex_annexB CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex = "annexB"

    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex_annexC CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelAnnex = "annexC"
)

// CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation represents can be transmitted.
type CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation string

const (
    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation_qam64 CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation = "qam64"

    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation_qam256 CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation = "qam256"

    CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation_qam1024 CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelModulation = "qam1024"
)

// CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable
// A wideband channel is a logical grouping of one or more
// physical RF channels over which Wideband MPEG-TS packets
// are carried.
// 
// This table contains association information of the wideband
// channels to the RF channels that are available for the WCMTS.
type CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides a list of attributes for a single association between a
    // wideband channel and a RF channel.  An entry in this table exists for each
    // association between a wideband channel and RF channel on the WCMTS. It is
    // indexed by the ifIndex of the wideband channel and entPhysicalIndex and
    // ccwbRFChannelNum which represents the RF channel.  This object may be
    // modified or deleted once they are already created. The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable_CcwbWBtoRFMappingEntry.
    CcwbWBtoRFMappingEntry []*CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable_CcwbWBtoRFMappingEntry
}

func (ccwbWBtoRFMappingTable *CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable) GetEntityData() *types.CommonEntityData {
    ccwbWBtoRFMappingTable.EntityData.YFilter = ccwbWBtoRFMappingTable.YFilter
    ccwbWBtoRFMappingTable.EntityData.YangName = "ccwbWBtoRFMappingTable"
    ccwbWBtoRFMappingTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBtoRFMappingTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbWBtoRFMappingTable.EntityData.SegmentPath = "ccwbWBtoRFMappingTable"
    ccwbWBtoRFMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBtoRFMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBtoRFMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBtoRFMappingTable.EntityData.Children = types.NewOrderedMap()
    ccwbWBtoRFMappingTable.EntityData.Children.Append("ccwbWBtoRFMappingEntry", types.YChild{"CcwbWBtoRFMappingEntry", nil})
    for i := range ccwbWBtoRFMappingTable.CcwbWBtoRFMappingEntry {
        ccwbWBtoRFMappingTable.EntityData.Children.Append(types.GetSegmentPath(ccwbWBtoRFMappingTable.CcwbWBtoRFMappingEntry[i]), types.YChild{"CcwbWBtoRFMappingEntry", ccwbWBtoRFMappingTable.CcwbWBtoRFMappingEntry[i]})
    }
    ccwbWBtoRFMappingTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbWBtoRFMappingTable.EntityData.YListKeys = []string {}

    return &(ccwbWBtoRFMappingTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable_CcwbWBtoRFMappingEntry
// An entry provides a list of attributes for a single
// association between a wideband channel and a RF channel.
// 
// An entry in this table exists for each association
// between a wideband channel and RF channel on the WCMTS.
// It is indexed by the ifIndex of the wideband channel
// and entPhysicalIndex and ccwbRFChannelNum which
// represents the RF channel.
// 
// This object may be modified or deleted once they are
// already created.
type CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable_CcwbWBtoRFMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_cable_wideband_mib.CISCOCABLEWIDEBANDMIB_CcwbRFChannelTable_CcwbRFChannelEntry_CcwbRFChannelNum
    CcwbRFChannelNum interface{}

    // The percentage of the RF channel bandwidth allocated for this wideband
    // channel. The type is interface{} with range: 1..100. Units are percent.
    CcwbWBtoRFBandwidth interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcwbWBtoRFStorageType interface{}

    // Controls and reflects the status of rows in this table. It can be used for
    // creating and deleting entries in this table.  The ccwbWBtoRFBandwidth must
    // be valid for a row to be created. When ccwbWBtoRFRowStatus is 'active', the
    // object ccwbWBtoRFBandwidth can be modified. The type is RowStatus.
    CcwbWBtoRFRowStatus interface{}
}

func (ccwbWBtoRFMappingEntry *CISCOCABLEWIDEBANDMIB_CcwbWBtoRFMappingTable_CcwbWBtoRFMappingEntry) GetEntityData() *types.CommonEntityData {
    ccwbWBtoRFMappingEntry.EntityData.YFilter = ccwbWBtoRFMappingEntry.YFilter
    ccwbWBtoRFMappingEntry.EntityData.YangName = "ccwbWBtoRFMappingEntry"
    ccwbWBtoRFMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBtoRFMappingEntry.EntityData.ParentYangName = "ccwbWBtoRFMappingTable"
    ccwbWBtoRFMappingEntry.EntityData.SegmentPath = "ccwbWBtoRFMappingEntry" + types.AddKeyToken(ccwbWBtoRFMappingEntry.IfIndex, "ifIndex") + types.AddKeyToken(ccwbWBtoRFMappingEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(ccwbWBtoRFMappingEntry.CcwbRFChannelNum, "ccwbRFChannelNum")
    ccwbWBtoRFMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBtoRFMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBtoRFMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBtoRFMappingEntry.EntityData.Children = types.NewOrderedMap()
    ccwbWBtoRFMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ccwbWBtoRFMappingEntry.IfIndex})
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ccwbWBtoRFMappingEntry.EntPhysicalIndex})
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("ccwbRFChannelNum", types.YLeaf{"CcwbRFChannelNum", ccwbWBtoRFMappingEntry.CcwbRFChannelNum})
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("ccwbWBtoRFBandwidth", types.YLeaf{"CcwbWBtoRFBandwidth", ccwbWBtoRFMappingEntry.CcwbWBtoRFBandwidth})
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("ccwbWBtoRFStorageType", types.YLeaf{"CcwbWBtoRFStorageType", ccwbWBtoRFMappingEntry.CcwbWBtoRFStorageType})
    ccwbWBtoRFMappingEntry.EntityData.Leafs.Append("ccwbWBtoRFRowStatus", types.YLeaf{"CcwbWBtoRFRowStatus", ccwbWBtoRFMappingEntry.CcwbWBtoRFRowStatus})

    ccwbWBtoRFMappingEntry.EntityData.YListKeys = []string {"IfIndex", "EntPhysicalIndex", "CcwbRFChannelNum"}

    return &(ccwbWBtoRFMappingEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable
// This table contains information of the mapping
// of the wideband channels to the Narrowband channels
// that are available on the WCMTS.
// 
// The wideband protocol utilizes the existing narrowband
// downstream channel for carrying the MAC management and
// signaling messages and the associated narrowband upstream
// for return data traffic and signaling.
type CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides a list of attributes for a association between a wideband
    // channel and a narrowband channel.  An entry in this table exists for each
    // association between a wideband channel and narrowband channel on the WCMTS.
    // The valid ifType for the ifIndex used here is,
    // ciscoDocsCableWBDownstream(224). The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable_CcwbWBtoNBMappingEntry.
    CcwbWBtoNBMappingEntry []*CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable_CcwbWBtoNBMappingEntry
}

func (ccwbWBtoNBMappingTable *CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable) GetEntityData() *types.CommonEntityData {
    ccwbWBtoNBMappingTable.EntityData.YFilter = ccwbWBtoNBMappingTable.YFilter
    ccwbWBtoNBMappingTable.EntityData.YangName = "ccwbWBtoNBMappingTable"
    ccwbWBtoNBMappingTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBtoNBMappingTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbWBtoNBMappingTable.EntityData.SegmentPath = "ccwbWBtoNBMappingTable"
    ccwbWBtoNBMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBtoNBMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBtoNBMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBtoNBMappingTable.EntityData.Children = types.NewOrderedMap()
    ccwbWBtoNBMappingTable.EntityData.Children.Append("ccwbWBtoNBMappingEntry", types.YChild{"CcwbWBtoNBMappingEntry", nil})
    for i := range ccwbWBtoNBMappingTable.CcwbWBtoNBMappingEntry {
        ccwbWBtoNBMappingTable.EntityData.Children.Append(types.GetSegmentPath(ccwbWBtoNBMappingTable.CcwbWBtoNBMappingEntry[i]), types.YChild{"CcwbWBtoNBMappingEntry", ccwbWBtoNBMappingTable.CcwbWBtoNBMappingEntry[i]})
    }
    ccwbWBtoNBMappingTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbWBtoNBMappingTable.EntityData.YListKeys = []string {}

    return &(ccwbWBtoNBMappingTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable_CcwbWBtoNBMappingEntry
// An entry provides a list of attributes for a
// association between a wideband channel and a narrowband
// channel.
// 
// An entry in this table exists for each association
// between a wideband channel and narrowband channel on the
// WCMTS. The valid ifType for the ifIndex used here is,
// ciscoDocsCableWBDownstream(224).
type CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable_CcwbWBtoNBMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The ifIndex of the narrowband cable interface
    // associated with this wideband channel. The type is interface{} with range:
    // 1..2147483647.
    CcwbWBtoNBifIndexForNB interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcwbWBtoNBStorageType interface{}

    // Controls and reflects the status of rows in this table. It can be used for
    // creating and deleting entries in this table. The object
    // ccwbWBtoNBifIndexForNB must be valid in order for row to be created. The
    // type is RowStatus.
    CcwbWBtoNBRowStatus interface{}
}

func (ccwbWBtoNBMappingEntry *CISCOCABLEWIDEBANDMIB_CcwbWBtoNBMappingTable_CcwbWBtoNBMappingEntry) GetEntityData() *types.CommonEntityData {
    ccwbWBtoNBMappingEntry.EntityData.YFilter = ccwbWBtoNBMappingEntry.YFilter
    ccwbWBtoNBMappingEntry.EntityData.YangName = "ccwbWBtoNBMappingEntry"
    ccwbWBtoNBMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBtoNBMappingEntry.EntityData.ParentYangName = "ccwbWBtoNBMappingTable"
    ccwbWBtoNBMappingEntry.EntityData.SegmentPath = "ccwbWBtoNBMappingEntry" + types.AddKeyToken(ccwbWBtoNBMappingEntry.IfIndex, "ifIndex") + types.AddKeyToken(ccwbWBtoNBMappingEntry.CcwbWBtoNBifIndexForNB, "ccwbWBtoNBifIndexForNB")
    ccwbWBtoNBMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBtoNBMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBtoNBMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBtoNBMappingEntry.EntityData.Children = types.NewOrderedMap()
    ccwbWBtoNBMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbWBtoNBMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ccwbWBtoNBMappingEntry.IfIndex})
    ccwbWBtoNBMappingEntry.EntityData.Leafs.Append("ccwbWBtoNBifIndexForNB", types.YLeaf{"CcwbWBtoNBifIndexForNB", ccwbWBtoNBMappingEntry.CcwbWBtoNBifIndexForNB})
    ccwbWBtoNBMappingEntry.EntityData.Leafs.Append("ccwbWBtoNBStorageType", types.YLeaf{"CcwbWBtoNBStorageType", ccwbWBtoNBMappingEntry.CcwbWBtoNBStorageType})
    ccwbWBtoNBMappingEntry.EntityData.Leafs.Append("ccwbWBtoNBRowStatus", types.YLeaf{"CcwbWBtoNBRowStatus", ccwbWBtoNBMappingEntry.CcwbWBtoNBRowStatus})

    ccwbWBtoNBMappingEntry.EntityData.YListKeys = []string {"IfIndex", "CcwbWBtoNBifIndexForNB"}

    return &(ccwbWBtoNBMappingEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable
// This table provides information about a
// Wideband BG interface, whether its configured
// to carry multicast or non-multicast traffic.
// For multicast the BG interface type is
// Secondary and for non-multicast its non-Secondary.
// Other objects could be added to this later in
// the future.
type CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table provides information about each Wideband BG
    // interface whose ifType is ciscoDocsCableWBDownstream(224). The type is
    // slice of CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable_CcwbWBBondGrpEntry.
    CcwbWBBondGrpEntry []*CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable_CcwbWBBondGrpEntry
}

func (ccwbWBBondGrpTable *CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable) GetEntityData() *types.CommonEntityData {
    ccwbWBBondGrpTable.EntityData.YFilter = ccwbWBBondGrpTable.YFilter
    ccwbWBBondGrpTable.EntityData.YangName = "ccwbWBBondGrpTable"
    ccwbWBBondGrpTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBBondGrpTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbWBBondGrpTable.EntityData.SegmentPath = "ccwbWBBondGrpTable"
    ccwbWBBondGrpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBBondGrpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBBondGrpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBBondGrpTable.EntityData.Children = types.NewOrderedMap()
    ccwbWBBondGrpTable.EntityData.Children.Append("ccwbWBBondGrpEntry", types.YChild{"CcwbWBBondGrpEntry", nil})
    for i := range ccwbWBBondGrpTable.CcwbWBBondGrpEntry {
        ccwbWBBondGrpTable.EntityData.Children.Append(types.GetSegmentPath(ccwbWBBondGrpTable.CcwbWBBondGrpEntry[i]), types.YChild{"CcwbWBBondGrpEntry", ccwbWBBondGrpTable.CcwbWBBondGrpEntry[i]})
    }
    ccwbWBBondGrpTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbWBBondGrpTable.EntityData.YListKeys = []string {}

    return &(ccwbWBBondGrpTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable_CcwbWBBondGrpEntry
// An entry in this table provides information
// about each Wideband BG interface whose ifType is
// ciscoDocsCableWBDownstream(224).
type CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable_CcwbWBBondGrpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object has the value 'true(1)' if the WB interface(BG) is Seconday and
    // the value 'false(2)' for non-Secondary. The type is bool.
    CcwbWBBondGrpSecondary interface{}
}

func (ccwbWBBondGrpEntry *CISCOCABLEWIDEBANDMIB_CcwbWBBondGrpTable_CcwbWBBondGrpEntry) GetEntityData() *types.CommonEntityData {
    ccwbWBBondGrpEntry.EntityData.YFilter = ccwbWBBondGrpEntry.YFilter
    ccwbWBBondGrpEntry.EntityData.YangName = "ccwbWBBondGrpEntry"
    ccwbWBBondGrpEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBBondGrpEntry.EntityData.ParentYangName = "ccwbWBBondGrpTable"
    ccwbWBBondGrpEntry.EntityData.SegmentPath = "ccwbWBBondGrpEntry" + types.AddKeyToken(ccwbWBBondGrpEntry.IfIndex, "ifIndex")
    ccwbWBBondGrpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBBondGrpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBBondGrpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBBondGrpEntry.EntityData.Children = types.NewOrderedMap()
    ccwbWBBondGrpEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbWBBondGrpEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ccwbWBBondGrpEntry.IfIndex})
    ccwbWBBondGrpEntry.EntityData.Leafs.Append("ccwbWBBondGrpSecondary", types.YLeaf{"CcwbWBBondGrpSecondary", ccwbWBBondGrpEntry.CcwbWBBondGrpSecondary})

    ccwbWBBondGrpEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(ccwbWBBondGrpEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable
// This table contains Wideband Cable Modem(WCM) connectivity state.
// A WCM connectivity state can be associated with multiple
// Wideband interfaces.
type CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status information for a single Wideband Cable Modem. An entry in this
    // table exists for each Wideband Cable Modem that is connected to the WCMTS.
    // The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry.
    CcwbWBCmStatusEntry []*CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry
}

func (ccwbWBCmStatusTable *CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable) GetEntityData() *types.CommonEntityData {
    ccwbWBCmStatusTable.EntityData.YFilter = ccwbWBCmStatusTable.YFilter
    ccwbWBCmStatusTable.EntityData.YangName = "ccwbWBCmStatusTable"
    ccwbWBCmStatusTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBCmStatusTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbWBCmStatusTable.EntityData.SegmentPath = "ccwbWBCmStatusTable"
    ccwbWBCmStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBCmStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBCmStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBCmStatusTable.EntityData.Children = types.NewOrderedMap()
    ccwbWBCmStatusTable.EntityData.Children.Append("ccwbWBCmStatusEntry", types.YChild{"CcwbWBCmStatusEntry", nil})
    for i := range ccwbWBCmStatusTable.CcwbWBCmStatusEntry {
        ccwbWBCmStatusTable.EntityData.Children.Append(types.GetSegmentPath(ccwbWBCmStatusTable.CcwbWBCmStatusEntry[i]), types.YChild{"CcwbWBCmStatusEntry", ccwbWBCmStatusTable.CcwbWBCmStatusEntry[i]})
    }
    ccwbWBCmStatusTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbWBCmStatusTable.EntityData.YListKeys = []string {}

    return &(ccwbWBCmStatusTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry
// Status information for a single Wideband Cable Modem.
// An entry in this table exists for each Wideband Cable Modem
// that is connected to the WCMTS.
type CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // Current WB Cable Modem connectivity state, as specified in the RF Interface
    // Specification. Returned status information is the WCM status as assumed by
    // the WCMTS, and indicates the following events:  The enumerations are:
    // offline(1)                : modem considered offline. others(2)            
    // : states is in                             ccwbWBCmStatusValue.
    // initRangingRcvd(3)        : modem sent initial ranging. initDhcpReqRcvd(4) 
    // : dhcp request received. onlineNetAccessDisabled(5): modem registered, but
    // network                             access for the WCM is disabled.
    // onlineKekAssigned(6)      : modem registered, BPI enabled                  
    // and KEK assigned. onlineTekAssigned(7)      : modem registered, BPI enabled
    // and TEK assigned. rejectBadMic(8)           : modem did attempt to register
    // but registration was refused                             due to bad mic.
    // rejectBadCos(9)           : modem did attempt to register                  
    // but registration was refused                             due to bad COS.
    // kekRejected(10)           : KEK modem key assignment                       
    // rejected. tekRejected(11)           : TEK modem key assignment             
    // rejected. online(12)                : modem registered, enabled for        
    // data. initTftpPacketRcvd(13)    : tftp packet received and option          
    // file transfer started. initTodRquestRcvd(14)     : Time of the Day (TOD)
    // request                             received. reset(15)                 :
    // modem is resetting. rangingInProgress(16)     : initial ranging is in
    // progress. dhcpGotIpAddr(17)         : modem has got an IP address
    // rejStaleConfig(18)        : modem did attempt to register                  
    // but registration was refused                             due to stale
    // Config. rejIpSpoof(19)            : modem did attempt to register          
    // but registration was refused                             due to IP Spoof.
    // rejClassFail(20)          : modem did attempt to register                  
    // but registration was refused                             due to Class
    // failure. rejRegNack(21)            : modem did attempt to register         
    // but no acknowledgement was                             received.
    // bpiKekExpired(22)         : KEK modem key assignment                       
    // expired. bpiTekExpired(23)         : TEK modem key assignment              
    // expired. shutdown(24)              : modem is in shutdown state.
    // channelChgInitRangingRcvd(25)  : modem sent initial ranging                
    // during channel change. channelChgRangingInProgress(26) : initial ranging is
    // in                                   progress during channel               
    // change. wbOnline(27)               : Wideband modem is online.
    // wbOnlinePrivacy(28)        : Wideband modem is online Privacy              
    // enabled. wbOnlineKekAssigned(29)    : Wideband modem is online             
    // and KEK assigned. wbOnlineTekAssigned(30)    : Wideband modem is online    
    // and TEK assigned. wbOnlineNetAccessDis(31)   : Wideband modem registered
    // but                              Network access disabled. wbKekReject(32)  
    // : KEK wideband modem key assignment                              rejected.
    // wbTekReject(33)            : TEK wideband modem key assignment             
    // rejected. wbNetAccessDisReject(34)   : wideband modem rejected -           
    // Net access disabled. wbPrivacyEnabReject(35)    : wideband modem rejected  
    // Privacy enabled. wbKekExpire(36)            : KEK Wideband modem key
    // assignment                              expired. wbTekExpire(37)           
    // : TEK wideband modem key assignment                              rejected.
    // wbNetAccessDisExpire(38)   : wideband modem expire - Network               
    // access disabled. wbPrivacyEnabExpire(39)    : wideband modem expire -
    // Privacy                              enabled.   This ccwbWBCmStatusValue
    // could return initRangingRcvd(3) or rangingInProgress(16) when the
    // ccwbWBCmStatusValue is ranging(2).  This ccwbWBCmStatusValue will return
    // others(2) when the ccwbWBCmStatusValue states is either rangingAborted(3),
    // rangingComplete(4), and ipComplete(5).  This ccwbWBCmStatusValue could
    // return wbonline(27), or onlineNetAccessDisabled(5) for WCM with BPI
    // disabled, or onlineNetAccessDisabled(5) or onlineTekAssigned(7) for WCM
    // with BPI enabled, when the ccwbWBCmStatusValue is registrationComplete(6). 
    // This ccwbWBCmStatusValue could return either rejectBadMic(8),
    // rejectBadCos(9) rejStaleConfig(19) or rejRegNack(22) when the
    // ccwbWBCmStatusValue is accessDenied(7) for possible reasons of cable modem
    // registration abort.  This ccwbWBCmStatusValue could return either
    // onlineKekAssigned(6), kekRejected(10), tekRejected(11), or online(12) for
    // the WCM with BPI enabled when the ccwbWBCmStatusValue is
    // registeredBPIInitializing(9).  The state rejectBadCos(9) is not applicable
    // for DOCSIS1.1 modems. The WCMTS only needs to report states it is able to
    // detect. The type is CcwbWBCmStatusValue.
    CcwbWBCmStatusValue interface{}
}

func (ccwbWBCmStatusEntry *CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry) GetEntityData() *types.CommonEntityData {
    ccwbWBCmStatusEntry.EntityData.YFilter = ccwbWBCmStatusEntry.YFilter
    ccwbWBCmStatusEntry.EntityData.YangName = "ccwbWBCmStatusEntry"
    ccwbWBCmStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBCmStatusEntry.EntityData.ParentYangName = "ccwbWBCmStatusTable"
    ccwbWBCmStatusEntry.EntityData.SegmentPath = "ccwbWBCmStatusEntry" + types.AddKeyToken(ccwbWBCmStatusEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    ccwbWBCmStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBCmStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBCmStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBCmStatusEntry.EntityData.Children = types.NewOrderedMap()
    ccwbWBCmStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbWBCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", ccwbWBCmStatusEntry.DocsIfCmtsCmStatusIndex})
    ccwbWBCmStatusEntry.EntityData.Leafs.Append("ccwbWBCmStatusValue", types.YLeaf{"CcwbWBCmStatusValue", ccwbWBCmStatusEntry.CcwbWBCmStatusValue})

    ccwbWBCmStatusEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(ccwbWBCmStatusEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue represents The WCMTS only needs to report states it is able to detect.
type CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue string

const (
    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_offline CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "offline"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_others CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "others"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_initRangingRcvd CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "initRangingRcvd"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_initDhcpReqRcvd CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "initDhcpReqRcvd"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_onlineNetAccessDisabled CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "onlineNetAccessDisabled"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_onlineKekAssigned CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "onlineKekAssigned"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_onlineTekAssigned CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "onlineTekAssigned"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejectBadMic CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejectBadMic"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejectBadCos CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejectBadCos"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_kekRejected CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "kekRejected"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_tekRejected CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "tekRejected"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_online CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "online"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_initTftpPacketRcvd CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "initTftpPacketRcvd"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_initTodRequestRcvd CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "initTodRequestRcvd"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_reset CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "reset"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rangingInProgress CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rangingInProgress"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_dhcpGotIpAddr CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "dhcpGotIpAddr"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejStaleConfig CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejStaleConfig"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejIpSpoof CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejIpSpoof"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejClassFail CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejClassFail"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_rejRegNack CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "rejRegNack"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_bpiKekExpired CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "bpiKekExpired"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_bpiTekExpired CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "bpiTekExpired"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_shutdown CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "shutdown"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_channelChgInitRangingRcvd CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "channelChgInitRangingRcvd"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_channelChgRangingInProgress CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "channelChgRangingInProgress"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbOnline CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbOnline"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbOnlinePrivacy CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbOnlinePrivacy"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbOnlineKekAssigned CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbOnlineKekAssigned"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbOnlineTekAssigned CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbOnlineTekAssigned"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbOnlineNetAccessDis CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbOnlineNetAccessDis"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbKekReject CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbKekReject"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbTekReject CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbTekReject"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbNetAccessDisReject CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbNetAccessDisReject"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbPrivacyEnabReject CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbPrivacyEnabReject"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbKekExpire CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbKekExpire"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbTekExpire CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbTekExpire"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbNetAccessDisExpire CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbNetAccessDisExpire"

    CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue_wbPrivacyEnabExpire CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusTable_CcwbWBCmStatusEntry_CcwbWBCmStatusValue = "wbPrivacyEnabExpire"
)

// CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable
// An entry in this table exists for each Wideband
// Cable Modem which links to one or more WB interface.
type CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry exists for each Wideband Cable Modem(WCM) which links to one or
    // more WB interface. The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable_CcwbWBCmStatusExtEntry.
    CcwbWBCmStatusExtEntry []*CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable_CcwbWBCmStatusExtEntry
}

func (ccwbWBCmStatusExtTable *CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable) GetEntityData() *types.CommonEntityData {
    ccwbWBCmStatusExtTable.EntityData.YFilter = ccwbWBCmStatusExtTable.YFilter
    ccwbWBCmStatusExtTable.EntityData.YangName = "ccwbWBCmStatusExtTable"
    ccwbWBCmStatusExtTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBCmStatusExtTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbWBCmStatusExtTable.EntityData.SegmentPath = "ccwbWBCmStatusExtTable"
    ccwbWBCmStatusExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBCmStatusExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBCmStatusExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBCmStatusExtTable.EntityData.Children = types.NewOrderedMap()
    ccwbWBCmStatusExtTable.EntityData.Children.Append("ccwbWBCmStatusExtEntry", types.YChild{"CcwbWBCmStatusExtEntry", nil})
    for i := range ccwbWBCmStatusExtTable.CcwbWBCmStatusExtEntry {
        ccwbWBCmStatusExtTable.EntityData.Children.Append(types.GetSegmentPath(ccwbWBCmStatusExtTable.CcwbWBCmStatusExtEntry[i]), types.YChild{"CcwbWBCmStatusExtEntry", ccwbWBCmStatusExtTable.CcwbWBCmStatusExtEntry[i]})
    }
    ccwbWBCmStatusExtTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbWBCmStatusExtTable.EntityData.YListKeys = []string {}

    return &(ccwbWBCmStatusExtTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable_CcwbWBCmStatusExtEntry
// This entry exists for each Wideband Cable Modem(WCM)
// which links to one or more WB interface.
type CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable_CcwbWBCmStatusExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // docs_if_mib.DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusIndex
    DocsIfCmtsCmStatusIndex interface{}

    // This attribute is a key. The value of this object uniquely identifies an
    // association between a WCM and a BG. The type is interface{} with range:
    // 1..2147483647.
    CcwbWBCmStatusExtIndex interface{}

    // ifIndex of the wideband channel associated with the WCM. The type is
    // interface{} with range: 1..2147483647.
    CcwbWBCmWBIfindex interface{}
}

func (ccwbWBCmStatusExtEntry *CISCOCABLEWIDEBANDMIB_CcwbWBCmStatusExtTable_CcwbWBCmStatusExtEntry) GetEntityData() *types.CommonEntityData {
    ccwbWBCmStatusExtEntry.EntityData.YFilter = ccwbWBCmStatusExtEntry.YFilter
    ccwbWBCmStatusExtEntry.EntityData.YangName = "ccwbWBCmStatusExtEntry"
    ccwbWBCmStatusExtEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbWBCmStatusExtEntry.EntityData.ParentYangName = "ccwbWBCmStatusExtTable"
    ccwbWBCmStatusExtEntry.EntityData.SegmentPath = "ccwbWBCmStatusExtEntry" + types.AddKeyToken(ccwbWBCmStatusExtEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex") + types.AddKeyToken(ccwbWBCmStatusExtEntry.CcwbWBCmStatusExtIndex, "ccwbWBCmStatusExtIndex")
    ccwbWBCmStatusExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbWBCmStatusExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbWBCmStatusExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbWBCmStatusExtEntry.EntityData.Children = types.NewOrderedMap()
    ccwbWBCmStatusExtEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbWBCmStatusExtEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", ccwbWBCmStatusExtEntry.DocsIfCmtsCmStatusIndex})
    ccwbWBCmStatusExtEntry.EntityData.Leafs.Append("ccwbWBCmStatusExtIndex", types.YLeaf{"CcwbWBCmStatusExtIndex", ccwbWBCmStatusExtEntry.CcwbWBCmStatusExtIndex})
    ccwbWBCmStatusExtEntry.EntityData.Leafs.Append("ccwbWBCmWBIfindex", types.YLeaf{"CcwbWBCmWBIfindex", ccwbWBCmStatusExtEntry.CcwbWBCmWBIfindex})

    ccwbWBCmStatusExtEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex", "CcwbWBCmStatusExtIndex"}

    return &(ccwbWBCmStatusExtEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable
// This table contains the description of a Fiber Node
// on a CMTS. An entry in this table exists for each
// FiberNode ID.
type CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry provides the description of each fiber node on the CMTS and it
    // is part of the Fiber node configuration. The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable_CcwbFiberNodeDescrEntry.
    CcwbFiberNodeDescrEntry []*CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable_CcwbFiberNodeDescrEntry
}

func (ccwbFiberNodeDescrTable *CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable) GetEntityData() *types.CommonEntityData {
    ccwbFiberNodeDescrTable.EntityData.YFilter = ccwbFiberNodeDescrTable.YFilter
    ccwbFiberNodeDescrTable.EntityData.YangName = "ccwbFiberNodeDescrTable"
    ccwbFiberNodeDescrTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbFiberNodeDescrTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbFiberNodeDescrTable.EntityData.SegmentPath = "ccwbFiberNodeDescrTable"
    ccwbFiberNodeDescrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbFiberNodeDescrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbFiberNodeDescrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbFiberNodeDescrTable.EntityData.Children = types.NewOrderedMap()
    ccwbFiberNodeDescrTable.EntityData.Children.Append("ccwbFiberNodeDescrEntry", types.YChild{"CcwbFiberNodeDescrEntry", nil})
    for i := range ccwbFiberNodeDescrTable.CcwbFiberNodeDescrEntry {
        ccwbFiberNodeDescrTable.EntityData.Children.Append(types.GetSegmentPath(ccwbFiberNodeDescrTable.CcwbFiberNodeDescrEntry[i]), types.YChild{"CcwbFiberNodeDescrEntry", ccwbFiberNodeDescrTable.CcwbFiberNodeDescrEntry[i]})
    }
    ccwbFiberNodeDescrTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbFiberNodeDescrTable.EntityData.YListKeys = []string {}

    return &(ccwbFiberNodeDescrTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable_CcwbFiberNodeDescrEntry
// This entry provides the description of each fiber node
// on the CMTS and it is part of the Fiber node configuration.
type CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable_CcwbFiberNodeDescrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..1000. Refers to
    // cisco_cable_wideband_mib.CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry_CcwbFiberNodeID
    CcwbFiberNodeID interface{}

    // This object contains the user configured description string of the fiber
    // node. The type is string with length: 0..80.
    CcwbFiberNodeDescription interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcwbFiberNodeDescrStorageType interface{}

    // Controls and reflects the status of rows in this table. It can be used for
    // creating and deleting entries in this table. ccwbFiberNodeDescription must
    // not be null in order for row to be created. The type is RowStatus.
    CcwbFiberNodeDescrRowStatus interface{}
}

func (ccwbFiberNodeDescrEntry *CISCOCABLEWIDEBANDMIB_CcwbFiberNodeDescrTable_CcwbFiberNodeDescrEntry) GetEntityData() *types.CommonEntityData {
    ccwbFiberNodeDescrEntry.EntityData.YFilter = ccwbFiberNodeDescrEntry.YFilter
    ccwbFiberNodeDescrEntry.EntityData.YangName = "ccwbFiberNodeDescrEntry"
    ccwbFiberNodeDescrEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbFiberNodeDescrEntry.EntityData.ParentYangName = "ccwbFiberNodeDescrTable"
    ccwbFiberNodeDescrEntry.EntityData.SegmentPath = "ccwbFiberNodeDescrEntry" + types.AddKeyToken(ccwbFiberNodeDescrEntry.CcwbFiberNodeID, "ccwbFiberNodeID")
    ccwbFiberNodeDescrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbFiberNodeDescrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbFiberNodeDescrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbFiberNodeDescrEntry.EntityData.Children = types.NewOrderedMap()
    ccwbFiberNodeDescrEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbFiberNodeDescrEntry.EntityData.Leafs.Append("ccwbFiberNodeID", types.YLeaf{"CcwbFiberNodeID", ccwbFiberNodeDescrEntry.CcwbFiberNodeID})
    ccwbFiberNodeDescrEntry.EntityData.Leafs.Append("ccwbFiberNodeDescription", types.YLeaf{"CcwbFiberNodeDescription", ccwbFiberNodeDescrEntry.CcwbFiberNodeDescription})
    ccwbFiberNodeDescrEntry.EntityData.Leafs.Append("ccwbFiberNodeDescrStorageType", types.YLeaf{"CcwbFiberNodeDescrStorageType", ccwbFiberNodeDescrEntry.CcwbFiberNodeDescrStorageType})
    ccwbFiberNodeDescrEntry.EntityData.Leafs.Append("ccwbFiberNodeDescrRowStatus", types.YLeaf{"CcwbFiberNodeDescrRowStatus", ccwbFiberNodeDescrEntry.CcwbFiberNodeDescrRowStatus})

    ccwbFiberNodeDescrEntry.EntityData.YListKeys = []string {"CcwbFiberNodeID"}

    return &(ccwbFiberNodeDescrEntry.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable
// This table provides configuration information of each Fiber node.
// It will provide topology information of each Fiber node, which
// includes information such as, Narrowband and Wideband QAMs.
type CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table exists for each FiberNode ID that is in use. It uses
    // two indices, i.e. ccwbFiberNodeID which is the Fiber node ID, and
    // ccwbFiberNodeGlobRFID, which is the combined bit mask of Narrowband RF
    // channels and Wideband rf-ports(rf-channels). The type is slice of
    // CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry.
    CcwbFiberNodeEntry []*CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry
}

func (ccwbFiberNodeTable *CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable) GetEntityData() *types.CommonEntityData {
    ccwbFiberNodeTable.EntityData.YFilter = ccwbFiberNodeTable.YFilter
    ccwbFiberNodeTable.EntityData.YangName = "ccwbFiberNodeTable"
    ccwbFiberNodeTable.EntityData.BundleName = "cisco_ios_xe"
    ccwbFiberNodeTable.EntityData.ParentYangName = "CISCO-CABLE-WIDEBAND-MIB"
    ccwbFiberNodeTable.EntityData.SegmentPath = "ccwbFiberNodeTable"
    ccwbFiberNodeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbFiberNodeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbFiberNodeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbFiberNodeTable.EntityData.Children = types.NewOrderedMap()
    ccwbFiberNodeTable.EntityData.Children.Append("ccwbFiberNodeEntry", types.YChild{"CcwbFiberNodeEntry", nil})
    for i := range ccwbFiberNodeTable.CcwbFiberNodeEntry {
        ccwbFiberNodeTable.EntityData.Children.Append(types.GetSegmentPath(ccwbFiberNodeTable.CcwbFiberNodeEntry[i]), types.YChild{"CcwbFiberNodeEntry", ccwbFiberNodeTable.CcwbFiberNodeEntry[i]})
    }
    ccwbFiberNodeTable.EntityData.Leafs = types.NewOrderedMap()

    ccwbFiberNodeTable.EntityData.YListKeys = []string {}

    return &(ccwbFiberNodeTable.EntityData)
}

// CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry
// An entry in this table exists for each FiberNode ID that is in use.
// It uses two indices, i.e. ccwbFiberNodeID which is the
// Fiber node ID, and ccwbFiberNodeGlobRFID, which is the combined
// bit mask of Narrowband RF channels and Wideband
// rf-ports(rf-channels).
type CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object represents the Fiber node ID. Each
    // Fiber node configuration on the CMTS is assigned a unique Fiber node ID.
    // The type is interface{} with range: 1..1000.
    CcwbFiberNodeID interface{}

    // This attribute is a key. This is the RF ID of both Narrowband and Wideband
    // QAMs(rf-channels) combined. The type is interface{} with range:
    // 0..4294967295.
    CcwbFiberNodeGlobRFID interface{}

    // This object represents the Narrowband Ifindex of the  RF downstream channel
    // which is part of the Fiber node configuation. The type is interface{} with
    // range: 0..2147483647.
    CcwbFiberNodeNBIfIndx interface{}

    // This object represents the entity physical index of Wideband controller
    // card. This card contains the RF port which is part of the
    // ccwbFiberNodeGlobRFID bit mask. A value of zero means the index is invalid.
    // ccwbFiberNodeWBRFPort and ccwbFiberNodeWBContlrPhyIndx are mutually
    // inclusive. The type is interface{} with range: 0..2147483647.
    CcwbFiberNodeWBContlrPhyIndx interface{}

    // This object represents the RF downstream channel IDs (rf-ports) of the
    // wideband controller card. Each Wideband controller can have 24 RF channels.
    // ccwbFiberNodeWBRFPort and ccwbFiberNodeWBContlrPhyIndx are mutually
    // inclusive. The type is interface{} with range: 0..1024.
    CcwbFiberNodeWBRFPort interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcwbFiberNodeStorageType interface{}

    // Controls and reflects the status of rows in this table. It can be used for
    // creating and deleting entries in this table. The type is RowStatus.
    CcwbFiberNodeRowStatus interface{}
}

func (ccwbFiberNodeEntry *CISCOCABLEWIDEBANDMIB_CcwbFiberNodeTable_CcwbFiberNodeEntry) GetEntityData() *types.CommonEntityData {
    ccwbFiberNodeEntry.EntityData.YFilter = ccwbFiberNodeEntry.YFilter
    ccwbFiberNodeEntry.EntityData.YangName = "ccwbFiberNodeEntry"
    ccwbFiberNodeEntry.EntityData.BundleName = "cisco_ios_xe"
    ccwbFiberNodeEntry.EntityData.ParentYangName = "ccwbFiberNodeTable"
    ccwbFiberNodeEntry.EntityData.SegmentPath = "ccwbFiberNodeEntry" + types.AddKeyToken(ccwbFiberNodeEntry.CcwbFiberNodeID, "ccwbFiberNodeID") + types.AddKeyToken(ccwbFiberNodeEntry.CcwbFiberNodeGlobRFID, "ccwbFiberNodeGlobRFID")
    ccwbFiberNodeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccwbFiberNodeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccwbFiberNodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccwbFiberNodeEntry.EntityData.Children = types.NewOrderedMap()
    ccwbFiberNodeEntry.EntityData.Leafs = types.NewOrderedMap()
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeID", types.YLeaf{"CcwbFiberNodeID", ccwbFiberNodeEntry.CcwbFiberNodeID})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeGlobRFID", types.YLeaf{"CcwbFiberNodeGlobRFID", ccwbFiberNodeEntry.CcwbFiberNodeGlobRFID})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeNBIfIndx", types.YLeaf{"CcwbFiberNodeNBIfIndx", ccwbFiberNodeEntry.CcwbFiberNodeNBIfIndx})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeWBContlrPhyIndx", types.YLeaf{"CcwbFiberNodeWBContlrPhyIndx", ccwbFiberNodeEntry.CcwbFiberNodeWBContlrPhyIndx})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeWBRFPort", types.YLeaf{"CcwbFiberNodeWBRFPort", ccwbFiberNodeEntry.CcwbFiberNodeWBRFPort})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeStorageType", types.YLeaf{"CcwbFiberNodeStorageType", ccwbFiberNodeEntry.CcwbFiberNodeStorageType})
    ccwbFiberNodeEntry.EntityData.Leafs.Append("ccwbFiberNodeRowStatus", types.YLeaf{"CcwbFiberNodeRowStatus", ccwbFiberNodeEntry.CcwbFiberNodeRowStatus})

    ccwbFiberNodeEntry.EntityData.YListKeys = []string {"CcwbFiberNodeID", "CcwbFiberNodeGlobRFID"}

    return &(ccwbFiberNodeEntry.EntityData)
}

