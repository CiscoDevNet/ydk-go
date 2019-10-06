// This is the MIB Module for DOCSIS 2.0 compliant Radio
// Frequency (RF) interfaces in Cable Modems (CM) and
// Cable Modem Termination Systems (CMTS).
package docs_if_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_if_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-IF-MIB DOCS-IF-MIB}", reflect.TypeOf(DOCSIFMIB{}))
    ydk.RegisterEntity("DOCS-IF-MIB:DOCS-IF-MIB", reflect.TypeOf(DOCSIFMIB{}))
}

// DocsisUpstreamTypeStatus represents this type is used to specifically identify PHY mode.
type DocsisUpstreamTypeStatus string

const (
    DocsisUpstreamTypeStatus_unknown DocsisUpstreamTypeStatus = "unknown"

    DocsisUpstreamTypeStatus_tdma DocsisUpstreamTypeStatus = "tdma"

    DocsisUpstreamTypeStatus_atdma DocsisUpstreamTypeStatus = "atdma"

    DocsisUpstreamTypeStatus_scdma DocsisUpstreamTypeStatus = "scdma"
)

// DocsisQosVersion represents Indicates the quality of service level.
type DocsisQosVersion string

const (
    DocsisQosVersion_docsis10 DocsisQosVersion = "docsis10"

    DocsisQosVersion_docsis11 DocsisQosVersion = "docsis11"
)

// DocsisUpstreamType represents Indicates the DOCSIS Upstream Channel Type.
type DocsisUpstreamType string

const (
    DocsisUpstreamType_unknown DocsisUpstreamType = "unknown"

    DocsisUpstreamType_tdma DocsisUpstreamType = "tdma"

    DocsisUpstreamType_atdma DocsisUpstreamType = "atdma"

    DocsisUpstreamType_scdma DocsisUpstreamType = "scdma"

    DocsisUpstreamType_tdmaAndAtdma DocsisUpstreamType = "tdmaAndAtdma"
)

// DocsisVersion represents 'docsis30' indicates DOCSIS 3.0.
type DocsisVersion string

const (
    DocsisVersion_docsis10 DocsisVersion = "docsis10"

    DocsisVersion_docsis11 DocsisVersion = "docsis11"

    DocsisVersion_docsis20 DocsisVersion = "docsis20"

    DocsisVersion_docsis30 DocsisVersion = "docsis30"
)

// DOCSIFMIB
type DOCSIFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DocsIfBaseObjects DOCSIFMIB_DocsIfBaseObjects

    
    DocsIfCmtsObjects DOCSIFMIB_DocsIfCmtsObjects

    // This table describes the attributes of downstream channels (frequency
    // bands).
    DocsIfDownstreamChannelTable DOCSIFMIB_DocsIfDownstreamChannelTable

    // This table describes the attributes of attached upstream channels.
    DocsIfUpstreamChannelTable DOCSIFMIB_DocsIfUpstreamChannelTable

    // Describes the attributes for each class of service.
    DocsIfQosProfileTable DOCSIFMIB_DocsIfQosProfileTable

    // At the CM, describes the PHY signal quality of downstream channels. At the
    // CMTS, describes the PHY signal quality of upstream channels. At the CMTS,
    // this table may exclude contention intervals.
    DocsIfSignalQualityTable DOCSIFMIB_DocsIfSignalQualityTable

    // Describes the attributes of each CM MAC interface, extending the
    // information available from ifEntry.
    DocsIfCmMacTable DOCSIFMIB_DocsIfCmMacTable

    // This table maintains a number of status objects and counters for Cable
    // Modems.
    DocsIfCmStatusTable DOCSIFMIB_DocsIfCmStatusTable

    // Describes the attributes of each upstream service queue on a CM.
    DocsIfCmServiceTable DOCSIFMIB_DocsIfCmServiceTable

    // Describes the attributes of each CMTS MAC interface, extending the
    // information available from ifEntry. Mandatory for all CMTS devices.
    DocsIfCmtsMacTable DOCSIFMIB_DocsIfCmtsMacTable

    // For the MAC layer, this group maintains a number of status objects and
    // counters.
    DocsIfCmtsStatusTable DOCSIFMIB_DocsIfCmtsStatusTable

    // A set of objects in the CMTS, maintained for each Cable Modem connected to
    // this CMTS.
    DocsIfCmtsCmStatusTable DOCSIFMIB_DocsIfCmtsCmStatusTable

    // Describes the attributes of upstream service queues in a Cable Modem
    // Termination System.
    DocsIfCmtsServiceTable DOCSIFMIB_DocsIfCmtsServiceTable

    // Describes a modulation profile associated with one or more upstream
    // channels.
    DocsIfCmtsModulationTable DOCSIFMIB_DocsIfCmtsModulationTable

    // This is a table to provide a quick access index into the
    // docsIfCmtsCmStatusTable. There is exactly one row in this table for each
    // row in the docsIfCmtsCmStatusTable. In general, the management station
    // should use this table only to get a pointer into the
    // docsIfCmtsCmStatusTable (which corresponds to the CM's RF interface MAC
    // address), and should not iterate (e.g. GetNext through) this table.
    DocsIfCmtsMacToCmTable DOCSIFMIB_DocsIfCmtsMacToCmTable

    // Reports utilization statistics for attached upstream and downstream
    // physical channels.
    DocsIfCmtsChannelUtilizationTable DOCSIFMIB_DocsIfCmtsChannelUtilizationTable

    // This table is implemented at the CMTS to collect downstream channel
    // statistics for utilization calculations.
    DocsIfCmtsDownChannelCounterTable DOCSIFMIB_DocsIfCmtsDownChannelCounterTable

    // This table is implemented at the CMTS to provide upstream channel
    // statistics appropriate for channel utilization calculations.
    DocsIfCmtsUpChannelCounterTable DOCSIFMIB_DocsIfCmtsUpChannelCounterTable
}

func (dOCSIFMIB *DOCSIFMIB) GetEntityData() *types.CommonEntityData {
    dOCSIFMIB.EntityData.YFilter = dOCSIFMIB.YFilter
    dOCSIFMIB.EntityData.YangName = "DOCS-IF-MIB"
    dOCSIFMIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSIFMIB.EntityData.ParentYangName = "DOCS-IF-MIB"
    dOCSIFMIB.EntityData.SegmentPath = "DOCS-IF-MIB:DOCS-IF-MIB"
    dOCSIFMIB.EntityData.AbsolutePath = dOCSIFMIB.EntityData.SegmentPath
    dOCSIFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSIFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSIFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSIFMIB.EntityData.Children = types.NewOrderedMap()
    dOCSIFMIB.EntityData.Children.Append("docsIfBaseObjects", types.YChild{"DocsIfBaseObjects", &dOCSIFMIB.DocsIfBaseObjects})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsObjects", types.YChild{"DocsIfCmtsObjects", &dOCSIFMIB.DocsIfCmtsObjects})
    dOCSIFMIB.EntityData.Children.Append("docsIfDownstreamChannelTable", types.YChild{"DocsIfDownstreamChannelTable", &dOCSIFMIB.DocsIfDownstreamChannelTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfUpstreamChannelTable", types.YChild{"DocsIfUpstreamChannelTable", &dOCSIFMIB.DocsIfUpstreamChannelTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfQosProfileTable", types.YChild{"DocsIfQosProfileTable", &dOCSIFMIB.DocsIfQosProfileTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfSignalQualityTable", types.YChild{"DocsIfSignalQualityTable", &dOCSIFMIB.DocsIfSignalQualityTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmMacTable", types.YChild{"DocsIfCmMacTable", &dOCSIFMIB.DocsIfCmMacTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmStatusTable", types.YChild{"DocsIfCmStatusTable", &dOCSIFMIB.DocsIfCmStatusTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmServiceTable", types.YChild{"DocsIfCmServiceTable", &dOCSIFMIB.DocsIfCmServiceTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsMacTable", types.YChild{"DocsIfCmtsMacTable", &dOCSIFMIB.DocsIfCmtsMacTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsStatusTable", types.YChild{"DocsIfCmtsStatusTable", &dOCSIFMIB.DocsIfCmtsStatusTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsCmStatusTable", types.YChild{"DocsIfCmtsCmStatusTable", &dOCSIFMIB.DocsIfCmtsCmStatusTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsServiceTable", types.YChild{"DocsIfCmtsServiceTable", &dOCSIFMIB.DocsIfCmtsServiceTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsModulationTable", types.YChild{"DocsIfCmtsModulationTable", &dOCSIFMIB.DocsIfCmtsModulationTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsMacToCmTable", types.YChild{"DocsIfCmtsMacToCmTable", &dOCSIFMIB.DocsIfCmtsMacToCmTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsChannelUtilizationTable", types.YChild{"DocsIfCmtsChannelUtilizationTable", &dOCSIFMIB.DocsIfCmtsChannelUtilizationTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsDownChannelCounterTable", types.YChild{"DocsIfCmtsDownChannelCounterTable", &dOCSIFMIB.DocsIfCmtsDownChannelCounterTable})
    dOCSIFMIB.EntityData.Children.Append("docsIfCmtsUpChannelCounterTable", types.YChild{"DocsIfCmtsUpChannelCounterTable", &dOCSIFMIB.DocsIfCmtsUpChannelCounterTable})
    dOCSIFMIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSIFMIB.EntityData.YListKeys = []string {}

    return &(dOCSIFMIB.EntityData)
}

// DOCSIFMIB_DocsIfBaseObjects
type DOCSIFMIB_DocsIfBaseObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indication of the DOCSIS capability of the device. The type is
    // DocsisVersion.
    DocsIfDocsisBaseCapability interface{}
}

func (docsIfBaseObjects *DOCSIFMIB_DocsIfBaseObjects) GetEntityData() *types.CommonEntityData {
    docsIfBaseObjects.EntityData.YFilter = docsIfBaseObjects.YFilter
    docsIfBaseObjects.EntityData.YangName = "docsIfBaseObjects"
    docsIfBaseObjects.EntityData.BundleName = "cisco_ios_xe"
    docsIfBaseObjects.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfBaseObjects.EntityData.SegmentPath = "docsIfBaseObjects"
    docsIfBaseObjects.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfBaseObjects.EntityData.SegmentPath
    docsIfBaseObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfBaseObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfBaseObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfBaseObjects.EntityData.Children = types.NewOrderedMap()
    docsIfBaseObjects.EntityData.Leafs = types.NewOrderedMap()
    docsIfBaseObjects.EntityData.Leafs.Append("docsIfDocsisBaseCapability", types.YLeaf{"DocsIfDocsisBaseCapability", docsIfBaseObjects.DocsIfDocsisBaseCapability})

    docsIfBaseObjects.EntityData.YListKeys = []string {}

    return &(docsIfBaseObjects.EntityData)
}

// DOCSIFMIB_DocsIfCmtsObjects
type DOCSIFMIB_DocsIfCmtsObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object specifies permitted methods of creating entries in
    // docsIfQosProfileTable. CreateByManagement(0) is set if entries can be
    // created using SNMP. UpdateByManagement(1) is set if updating entries using
    // SNMP is permitted. CreateByModems(2) is set if entries can be created based
    // on information in REG-REQ MAC messages received from Cable Modems.
    // Information in this object is only applicable if docsIfQosProfileTable is
    // implemented as read-create. Otherwise, this object is implemented as
    // read-only and returns CreateByModems(2). Either CreateByManagement(0) or
    // CreateByModems(1) must be set when writing to this object. Note that BITS
    // objects are encoded most significant bit first. For example, if bit 2 is
    // set, the value of this object is the octet string '20'H. The type is
    // map[string]bool.
    DocsIfCmtsQosProfilePermissions interface{}

    // The time interval in seconds over which the channel utilization index is
    // calculated. All upstream/downstream channels use the same
    // docsIfCmtsChannelUtilizationInterval.  Setting a value of zero disables
    // utilization reporting. A channel utilization index is calculated over a
    // fixed window applying to the most recent
    // docsIfCmtsChannelUtilizationInterval. It would therefore be prudent to use
    // a relatively short docsIfCmtsChannelUtilizationInterval. The type is
    // interface{} with range: 0..86400. Units are seconds.
    DocsIfCmtsChannelUtilizationInterval interface{}
}

func (docsIfCmtsObjects *DOCSIFMIB_DocsIfCmtsObjects) GetEntityData() *types.CommonEntityData {
    docsIfCmtsObjects.EntityData.YFilter = docsIfCmtsObjects.YFilter
    docsIfCmtsObjects.EntityData.YangName = "docsIfCmtsObjects"
    docsIfCmtsObjects.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsObjects.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsObjects.EntityData.SegmentPath = "docsIfCmtsObjects"
    docsIfCmtsObjects.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsObjects.EntityData.SegmentPath
    docsIfCmtsObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsObjects.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsObjects.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsObjects.EntityData.Leafs.Append("docsIfCmtsQosProfilePermissions", types.YLeaf{"DocsIfCmtsQosProfilePermissions", docsIfCmtsObjects.DocsIfCmtsQosProfilePermissions})
    docsIfCmtsObjects.EntityData.Leafs.Append("docsIfCmtsChannelUtilizationInterval", types.YLeaf{"DocsIfCmtsChannelUtilizationInterval", docsIfCmtsObjects.DocsIfCmtsChannelUtilizationInterval})

    docsIfCmtsObjects.EntityData.YListKeys = []string {}

    return &(docsIfCmtsObjects.EntityData)
}

// DOCSIFMIB_DocsIfDownstreamChannelTable
// This table describes the attributes of downstream
// channels (frequency bands).
type DOCSIFMIB_DocsIfDownstreamChannelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides a list of attributes for a single Downstream channel. An
    // entry in this table exists for each ifEntry with an ifType of
    // docsCableDownstream(128). The type is slice of
    // DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry.
    DocsIfDownstreamChannelEntry []*DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry
}

func (docsIfDownstreamChannelTable *DOCSIFMIB_DocsIfDownstreamChannelTable) GetEntityData() *types.CommonEntityData {
    docsIfDownstreamChannelTable.EntityData.YFilter = docsIfDownstreamChannelTable.YFilter
    docsIfDownstreamChannelTable.EntityData.YangName = "docsIfDownstreamChannelTable"
    docsIfDownstreamChannelTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfDownstreamChannelTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfDownstreamChannelTable.EntityData.SegmentPath = "docsIfDownstreamChannelTable"
    docsIfDownstreamChannelTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfDownstreamChannelTable.EntityData.SegmentPath
    docsIfDownstreamChannelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfDownstreamChannelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfDownstreamChannelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfDownstreamChannelTable.EntityData.Children = types.NewOrderedMap()
    docsIfDownstreamChannelTable.EntityData.Children.Append("docsIfDownstreamChannelEntry", types.YChild{"DocsIfDownstreamChannelEntry", nil})
    for i := range docsIfDownstreamChannelTable.DocsIfDownstreamChannelEntry {
        docsIfDownstreamChannelTable.EntityData.Children.Append(types.GetSegmentPath(docsIfDownstreamChannelTable.DocsIfDownstreamChannelEntry[i]), types.YChild{"DocsIfDownstreamChannelEntry", docsIfDownstreamChannelTable.DocsIfDownstreamChannelEntry[i]})
    }
    docsIfDownstreamChannelTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfDownstreamChannelTable.EntityData.YListKeys = []string {}

    return &(docsIfDownstreamChannelTable.EntityData)
}

// DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry
// An entry provides a list of attributes for a single
// Downstream channel.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableDownstream(128).
type DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The Cable Modem Termination System (CMTS) identification of the downstream
    // channel within this particular MAC interface. If the interface is down, the
    // object returns the most current value. If the downstream channel ID is
    // unknown, this object returns a value of 0. The type is interface{} with
    // range: 0..255.
    DocsIfDownChannelId interface{}

    // The center of the downstream frequency associated with this channel. This
    // object will return the current tuner frequency. If a CMTS provides IF
    // output, this object will return 0, unless this CMTS is in control of the
    // final downstream RF frequency.  See the associated compliance object for a
    // description of valid frequencies that may be written to this object. The
    // type is interface{} with range: 0..1000000000. Units are hertz.
    DocsIfDownChannelFrequency interface{}

    // The bandwidth of this downstream channel. Most implementations are expected
    // to support a channel width of 6 MHz (North America) and/or 8 MHz (Europe). 
    // See the associated compliance object for a description of the valid channel
    // widths for this object. The type is interface{} with range: 0..16000000.
    // Units are hertz.
    DocsIfDownChannelWidth interface{}

    // The modulation type associated with this downstream channel. If the
    // interface is down, this object either returns the configured value (CMTS),
    // the most current value (CM), or the value of unknown(1).  See the
    // associated conformance object for write conditions and limitations. See the
    // reference for specifics on the modulation profiles implied by qam64 and
    // qam256. The type is DocsIfDownChannelModulation.
    DocsIfDownChannelModulation interface{}

    // The Forward Error Correction (FEC) interleaving used for this downstream
    // channel. Values are defined as follows: taps8Increment16(3):   protection
    // 5.9/4.1 usec,                        latency .22/.15 msec
    // taps16Increment8(4):   protection 12/8.2 usec,                       
    // latency .48/.33 msec taps32Increment4(5):   protection 24/16 usec,         
    // latency .98/.68 msec taps64Increment2(6):   protection 47/33 usec,         
    // latency 2/1.4 msec taps128Increment1(7):  protection 95/66 usec,           
    // latency 4/2.8 msec taps12increment17(8):  protection 18/14 usec,           
    // latency 0.43/0.32 msec                        taps12increment17 is
    // implemented in                         conformance with EuroDOCSIS document
    // 'Adapted MIB-definitions - and a                         clarification for
    // MPEG-related issues - for                         EuroDOCSIS cable modem
    // systems' by tComLabs                        and should only be used for a
    // EuroDOCSIS MAC                         interface.     If the interface is
    // down, this object either returns the configured value (CMTS), the most
    // current value (CM), or the value of unknown(1). The value of other(2) is
    // returned if the interleave is known but not defined in the above list. See
    // the associated conformance object for write conditions and limitations. See
    // the reference for the FEC configuration described by the setting of this
    // object. The type is DocsIfDownChannelInterleave.
    DocsIfDownChannelInterleave interface{}

    // At the CMTS, the operational transmit power. At the CM, the received power
    // level. May be set to zero at the CM if power level measurement is not
    // supported. If the interface is down, this object either returns the
    // configured value (CMTS), the most current value (CM) or the value of 0. See
    // the associated conformance object for write conditions and limitations. See
    // the reference for recommended and required power levels. The type is
    // interface{} with range: -2147483648..2147483647. Units are dBmV.
    DocsIfDownChannelPower interface{}

    // The value of this object indicates the conformance of the implementation to
    // important regional cable standards. annexA : Annex A from ITU-J83 is used.
    // annexB : Annex B from ITU-J83 is used. annexC : Annex C from ITU-J83 is
    // used.  AnnexB is used for DOCSIS implementations. The type is
    // DocsIfDownChannelAnnex.
    DocsIfDownChannelAnnex interface{}

    // The storage type for this conceptual row. Entries with this object set to
    // permanent(4) do not require write operations for read-write objects. The
    // type is StorageType.
    DocsIfDownChannelStorageType interface{}
}

func (docsIfDownstreamChannelEntry *DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry) GetEntityData() *types.CommonEntityData {
    docsIfDownstreamChannelEntry.EntityData.YFilter = docsIfDownstreamChannelEntry.YFilter
    docsIfDownstreamChannelEntry.EntityData.YangName = "docsIfDownstreamChannelEntry"
    docsIfDownstreamChannelEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfDownstreamChannelEntry.EntityData.ParentYangName = "docsIfDownstreamChannelTable"
    docsIfDownstreamChannelEntry.EntityData.SegmentPath = "docsIfDownstreamChannelEntry" + types.AddKeyToken(docsIfDownstreamChannelEntry.IfIndex, "ifIndex")
    docsIfDownstreamChannelEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfDownstreamChannelTable/" + docsIfDownstreamChannelEntry.EntityData.SegmentPath
    docsIfDownstreamChannelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfDownstreamChannelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfDownstreamChannelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfDownstreamChannelEntry.EntityData.Children = types.NewOrderedMap()
    docsIfDownstreamChannelEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfDownstreamChannelEntry.IfIndex})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelId", types.YLeaf{"DocsIfDownChannelId", docsIfDownstreamChannelEntry.DocsIfDownChannelId})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelFrequency", types.YLeaf{"DocsIfDownChannelFrequency", docsIfDownstreamChannelEntry.DocsIfDownChannelFrequency})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelWidth", types.YLeaf{"DocsIfDownChannelWidth", docsIfDownstreamChannelEntry.DocsIfDownChannelWidth})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelModulation", types.YLeaf{"DocsIfDownChannelModulation", docsIfDownstreamChannelEntry.DocsIfDownChannelModulation})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelInterleave", types.YLeaf{"DocsIfDownChannelInterleave", docsIfDownstreamChannelEntry.DocsIfDownChannelInterleave})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelPower", types.YLeaf{"DocsIfDownChannelPower", docsIfDownstreamChannelEntry.DocsIfDownChannelPower})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelAnnex", types.YLeaf{"DocsIfDownChannelAnnex", docsIfDownstreamChannelEntry.DocsIfDownChannelAnnex})
    docsIfDownstreamChannelEntry.EntityData.Leafs.Append("docsIfDownChannelStorageType", types.YLeaf{"DocsIfDownChannelStorageType", docsIfDownstreamChannelEntry.DocsIfDownChannelStorageType})

    docsIfDownstreamChannelEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfDownstreamChannelEntry.EntityData)
}

// DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex represents AnnexB is used for DOCSIS implementations
type DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex string

const (
    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex_unknown DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex = "unknown"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex_other DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex = "other"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex_annexA DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex = "annexA"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex_annexB DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex = "annexB"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex_annexC DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelAnnex = "annexC"
)

// DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave represents configuration described by the setting of this object.
type DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave string

const (
    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_unknown DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "unknown"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_other DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "other"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps8Increment16 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps8Increment16"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps16Increment8 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps16Increment8"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps32Increment4 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps32Increment4"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps64Increment2 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps64Increment2"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps128Increment1 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps128Increment1"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave_taps12increment17 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelInterleave = "taps12increment17"
)

// DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation represents modulation profiles implied by qam64 and qam256.
type DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation string

const (
    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation_unknown DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation = "unknown"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation_other DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation = "other"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation_qam64 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation = "qam64"

    DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation_qam256 DOCSIFMIB_DocsIfDownstreamChannelTable_DocsIfDownstreamChannelEntry_DocsIfDownChannelModulation = "qam256"
)

// DOCSIFMIB_DocsIfUpstreamChannelTable
// This table describes the attributes of attached upstream
// channels.
type DOCSIFMIB_DocsIfUpstreamChannelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of attributes for a single upstream channel. For Docsis 2.0 CMTSs, an
    // entry in this table exists for  each ifEntry with an ifType of
    // docsCableUpstreamChannel (205). For Docsis 1.x CM/CMTSs and Docsis 2.0 CMs,
    // an entry in this table exists  for each ifEntry with an ifType of
    // docsCableUpstreamInterface (129). The type is slice of
    // DOCSIFMIB_DocsIfUpstreamChannelTable_DocsIfUpstreamChannelEntry.
    DocsIfUpstreamChannelEntry []*DOCSIFMIB_DocsIfUpstreamChannelTable_DocsIfUpstreamChannelEntry
}

func (docsIfUpstreamChannelTable *DOCSIFMIB_DocsIfUpstreamChannelTable) GetEntityData() *types.CommonEntityData {
    docsIfUpstreamChannelTable.EntityData.YFilter = docsIfUpstreamChannelTable.YFilter
    docsIfUpstreamChannelTable.EntityData.YangName = "docsIfUpstreamChannelTable"
    docsIfUpstreamChannelTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfUpstreamChannelTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfUpstreamChannelTable.EntityData.SegmentPath = "docsIfUpstreamChannelTable"
    docsIfUpstreamChannelTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfUpstreamChannelTable.EntityData.SegmentPath
    docsIfUpstreamChannelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfUpstreamChannelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfUpstreamChannelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfUpstreamChannelTable.EntityData.Children = types.NewOrderedMap()
    docsIfUpstreamChannelTable.EntityData.Children.Append("docsIfUpstreamChannelEntry", types.YChild{"DocsIfUpstreamChannelEntry", nil})
    for i := range docsIfUpstreamChannelTable.DocsIfUpstreamChannelEntry {
        docsIfUpstreamChannelTable.EntityData.Children.Append(types.GetSegmentPath(docsIfUpstreamChannelTable.DocsIfUpstreamChannelEntry[i]), types.YChild{"DocsIfUpstreamChannelEntry", docsIfUpstreamChannelTable.DocsIfUpstreamChannelEntry[i]})
    }
    docsIfUpstreamChannelTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfUpstreamChannelTable.EntityData.YListKeys = []string {}

    return &(docsIfUpstreamChannelTable.EntityData)
}

// DOCSIFMIB_DocsIfUpstreamChannelTable_DocsIfUpstreamChannelEntry
// List of attributes for a single upstream channel. For
// Docsis 2.0 CMTSs, an entry in this table exists for 
// each ifEntry with an ifType of docsCableUpstreamChannel (205).
// For Docsis 1.x CM/CMTSs and Docsis 2.0 CMs, an entry in this table exists 
// for each ifEntry with an ifType of docsCableUpstreamInterface (129).
type DOCSIFMIB_DocsIfUpstreamChannelTable_DocsIfUpstreamChannelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The CMTS identification of the upstream channel. The type is interface{}
    // with range: 0..255.
    DocsIfUpChannelId interface{}

    // The center of the frequency band associated with this upstream interface.
    // This object returns 0 if the frequency is undefined or unknown. Minimum
    // permitted upstream frequency is 5,000,000 Hz for current technology.  See
    // the associated conformance object for write conditions and limitations. The
    // type is interface{} with range: 0..1000000000. Units are hertz.
    DocsIfUpChannelFrequency interface{}

    // The bandwidth of this upstream interface. This object returns 0 if the
    // interface width is undefined or unknown. Minimum permitted interface width
    // is 200,000 Hz currently. See the associated conformance object for write
    // conditions and limitations. The type is interface{} with range:
    // 0..64000000. Units are hertz.
    DocsIfUpChannelWidth interface{}

    // An entry identical to the docsIfModIndex in the docsIfCmtsModulationTable
    // that describes this channel. This channel is further instantiated there by
    // a grouping of interval usage codes which together fully describe the
    // channel modulation. This object returns 0 if the docsIfCmtsModulationTable
    // entry does not exist or docsIfCmtsModulationTable is empty. See the
    // associated conformance object for write conditions and limitations. The
    // type is interface{} with range: 0..4294967295.
    DocsIfUpChannelModulationProfile interface{}

    // Applicable to TDMA and ATDMA channel types only. The number of 6.25
    // microsecond ticks in each upstream mini- slot. Returns zero if the value is
    // undefined, unknown or in case of an SCDMA channel. See the associated
    // conformance object for write conditions and limitations. . The type is
    // interface{} with range: 0..4294967295.
    DocsIfUpChannelSlotSize interface{}

    // At the CM, a measure of the current round trip time obtained from the
    // ranging offset (initial ranging offset + ranging offset adjustments). At
    // the CMTS, the maximum of timing offset, among all the CMs that  are/were
    // present on the channel, taking into account all ( initial +  periodic
    // )timing offset corrections that were sent for each of the CMs. Generally,
    // these measurements are positive, but if the  measurements are negative, the
    // value of this object is zero. Used for  timing of CM upstream transmissions
    // to ensure synchronized arrivals at  the CMTS. Units are in terms of (6.25
    // microseconds/64). The type is interface{} with range: 0..4294967295.
    DocsIfUpChannelTxTimingOffset interface{}

    // The initial random backoff window to use when retrying Ranging Requests.
    // Expressed as a power of 2. A value of 16 at the CMTS indicates that a
    // proprietary adaptive retry mechanism is to be used. See the associated
    // conformance object for write conditions and limitations. The type is
    // interface{} with range: 0..16.
    DocsIfUpChannelRangingBackoffStart interface{}

    // The final random backoff window to use when retrying Ranging Requests.
    // Expressed as a power of 2. A value of 16 at the CMTS indicates that a
    // proprietary adaptive retry mechanism is to be used. See the associated
    // conformance object for write conditions and limitations. The type is
    // interface{} with range: 0..16.
    DocsIfUpChannelRangingBackoffEnd interface{}

    // The initial random backoff window to use when retrying transmissions.
    // Expressed as a power of 2. A value of 16 at the CMTS indicates that a
    // proprietary adaptive retry mechanism is to be used. See the associated
    // conformance object for write conditions and limitations. The type is
    // interface{} with range: 0..16.
    DocsIfUpChannelTxBackoffStart interface{}

    // The final random backoff window to use when retrying transmissions.
    // Expressed as a power of 2. A value of 16 at the CMTS indicates that a
    // proprietary adaptive retry mechanism is to be used. See the associated
    // conformance object for write conditions and limitations. The type is
    // interface{} with range: 0..16.
    DocsIfUpChannelTxBackoffEnd interface{}

    // Applicable for SCDMA channel types only. Number of active codes. Returns
    // zero for Non-SCDMA channel types. Note that legal  values from 64..128 MUST
    // be non-prime. The type is interface{} with range: 0..0 | 64..128.
    DocsIfUpChannelScdmaActiveCodes interface{}

    // Applicable for SCDMA channel types only. The number of SCDMA codes per
    // mini-slot. Returns zero if the value is undefined, unknown or in case of a
    // TDMA or ATDMA channel. The type is interface{} with range: 0..0 | 2..32.
    DocsIfUpChannelScdmaCodesPerSlot interface{}

    // Applicable for SCDMA channel types only. SCDMA Frame size in units of
    // spreading intervals.  This value returns zero for non SCDMA Profiles. The
    // type is interface{} with range: 0..32.
    DocsIfUpChannelScdmaFrameSize interface{}

    // Applicable for SCDMA channel types only. 15 bit seed used for code hopping
    // sequence initialization. Returns zero for non-SCDMA channel types. The type
    // is interface{} with range: 0..32767.
    DocsIfUpChannelScdmaHoppingSeed interface{}

    // Defines the Upstream channel type. Given the channel type, other channel
    // attributes can be checked for value validity at the time of entry creation
    // and update. The type is DocsisUpstreamType.
    DocsIfUpChannelType interface{}

    // Intended for use when a temporary inactive upstream table row is created
    // for the purpose of manipulating SCDMA parameters for an active row. Refer
    // to the descriptions of docsIfUpChannelStatus  and docsIfUpChannelUpdate for
    // details of this procedure. This object contains the ifIndex value of the
    // active upstream row whose SCDMA parameters are to be adjusted. Although
    // this object was created to facilitate SCDMA parameter adjustment, it may
    // also be used at the vendor's discretion for non-SCDMA parameter adjustment.
    // This object must contain a value of zero for active upstream rows. The type
    // is interface{} with range: 0..2147483647.
    DocsIfUpChannelCloneFrom interface{}

    // Used to perform the transfer of adjusted SCDMA parameters from the
    // temporary upstream row to the active upstream row indicated by the
    // docsIfUpChannelCloneFrom object. The transfer is initiated through  an SNMP
    // SET of TRUE to this object. The SNMP SET will fail with a GEN_ERROR
    // (snmpv1) or COMMIT_FAILED_ERROR (snmpv2c/v3) if the adjusted SCDMA
    // parameter values are not compatible with each other. Although this object
    // was created to facilitate SCDMA parameter adjustment, it may also be used
    // at the vendor's discretion for non-SCDMA parameter adjustment. An SNMP GET
    // of this object always returns FALSE. The type is bool.
    DocsIfUpChannelUpdate interface{}

    // This object is generally intended to be used for the creation of a
    // temporary inactive upstream row for the purpose of adjusting the SCDMA
    // channel parameters of an active upstream row.  The recommended procedure
    // is: 1) Create an inactive row through an SNMP SET using createAndWait(5).  
    // Use an ifIndex value outside the operational range of the system. 2) Set
    // the docsIfUpChannelCloneFrom field to the ifIndex value of    the active
    // row whose SCDMA parameters require adjustment. 3) Adjust the SCDMA
    // parameter values using the new temporary inactive    row. 4) Update the
    // active row by setting object docsIfUpChannelUpdate to    TRUE. This SET
    // will fail if the adjusted SCDMA parameters are not    compatible with each
    // other. 5) Delete the temporary row through an SNMP SET using DELETE.  The
    // following restrictions apply to this object: 1) This object must contain a
    // value of active(1) for active rows. 2) Temporary inactive rows must be
    // created using createAndWait(5). 3) The only possible status change of a row
    // created using     createAndWait(5) (ie notInService(2)) is to destroy(6).  
    // These temporary rows must never become active. 4) A status transition from
    // active (1) to destroy (6) is not    permitted. Entries with
    // docsIfUpChannelStatus set to active(1)     are logically linked to a
    // physical interface, not temporarily     created to clone parameters. The
    // Interface MIB (RFC 2863)    ifAdminStatus should be used to take an
    // Upstream Channel offline.  Although this object was created to facilitate
    // SCDMA parameter adjustment, it may also be used at the vendor's discretion
    // for non-SCDMA parameter adjustment. The type is RowStatus.
    DocsIfUpChannelStatus interface{}

    // At the CMTS, used to enable or disable pre-equalization on the upstream
    // channel represented by this table instance.  At the CM, this object is
    // read-only and reflects the status of  pre-equalization as represented in
    // the RNG-RSP. The type is bool.
    DocsIfUpChannelPreEqEnable interface{}
}

func (docsIfUpstreamChannelEntry *DOCSIFMIB_DocsIfUpstreamChannelTable_DocsIfUpstreamChannelEntry) GetEntityData() *types.CommonEntityData {
    docsIfUpstreamChannelEntry.EntityData.YFilter = docsIfUpstreamChannelEntry.YFilter
    docsIfUpstreamChannelEntry.EntityData.YangName = "docsIfUpstreamChannelEntry"
    docsIfUpstreamChannelEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfUpstreamChannelEntry.EntityData.ParentYangName = "docsIfUpstreamChannelTable"
    docsIfUpstreamChannelEntry.EntityData.SegmentPath = "docsIfUpstreamChannelEntry" + types.AddKeyToken(docsIfUpstreamChannelEntry.IfIndex, "ifIndex")
    docsIfUpstreamChannelEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfUpstreamChannelTable/" + docsIfUpstreamChannelEntry.EntityData.SegmentPath
    docsIfUpstreamChannelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfUpstreamChannelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfUpstreamChannelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfUpstreamChannelEntry.EntityData.Children = types.NewOrderedMap()
    docsIfUpstreamChannelEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfUpstreamChannelEntry.IfIndex})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelId", types.YLeaf{"DocsIfUpChannelId", docsIfUpstreamChannelEntry.DocsIfUpChannelId})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelFrequency", types.YLeaf{"DocsIfUpChannelFrequency", docsIfUpstreamChannelEntry.DocsIfUpChannelFrequency})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelWidth", types.YLeaf{"DocsIfUpChannelWidth", docsIfUpstreamChannelEntry.DocsIfUpChannelWidth})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelModulationProfile", types.YLeaf{"DocsIfUpChannelModulationProfile", docsIfUpstreamChannelEntry.DocsIfUpChannelModulationProfile})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelSlotSize", types.YLeaf{"DocsIfUpChannelSlotSize", docsIfUpstreamChannelEntry.DocsIfUpChannelSlotSize})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelTxTimingOffset", types.YLeaf{"DocsIfUpChannelTxTimingOffset", docsIfUpstreamChannelEntry.DocsIfUpChannelTxTimingOffset})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelRangingBackoffStart", types.YLeaf{"DocsIfUpChannelRangingBackoffStart", docsIfUpstreamChannelEntry.DocsIfUpChannelRangingBackoffStart})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelRangingBackoffEnd", types.YLeaf{"DocsIfUpChannelRangingBackoffEnd", docsIfUpstreamChannelEntry.DocsIfUpChannelRangingBackoffEnd})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelTxBackoffStart", types.YLeaf{"DocsIfUpChannelTxBackoffStart", docsIfUpstreamChannelEntry.DocsIfUpChannelTxBackoffStart})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelTxBackoffEnd", types.YLeaf{"DocsIfUpChannelTxBackoffEnd", docsIfUpstreamChannelEntry.DocsIfUpChannelTxBackoffEnd})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelScdmaActiveCodes", types.YLeaf{"DocsIfUpChannelScdmaActiveCodes", docsIfUpstreamChannelEntry.DocsIfUpChannelScdmaActiveCodes})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelScdmaCodesPerSlot", types.YLeaf{"DocsIfUpChannelScdmaCodesPerSlot", docsIfUpstreamChannelEntry.DocsIfUpChannelScdmaCodesPerSlot})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelScdmaFrameSize", types.YLeaf{"DocsIfUpChannelScdmaFrameSize", docsIfUpstreamChannelEntry.DocsIfUpChannelScdmaFrameSize})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelScdmaHoppingSeed", types.YLeaf{"DocsIfUpChannelScdmaHoppingSeed", docsIfUpstreamChannelEntry.DocsIfUpChannelScdmaHoppingSeed})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelType", types.YLeaf{"DocsIfUpChannelType", docsIfUpstreamChannelEntry.DocsIfUpChannelType})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelCloneFrom", types.YLeaf{"DocsIfUpChannelCloneFrom", docsIfUpstreamChannelEntry.DocsIfUpChannelCloneFrom})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelUpdate", types.YLeaf{"DocsIfUpChannelUpdate", docsIfUpstreamChannelEntry.DocsIfUpChannelUpdate})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelStatus", types.YLeaf{"DocsIfUpChannelStatus", docsIfUpstreamChannelEntry.DocsIfUpChannelStatus})
    docsIfUpstreamChannelEntry.EntityData.Leafs.Append("docsIfUpChannelPreEqEnable", types.YLeaf{"DocsIfUpChannelPreEqEnable", docsIfUpstreamChannelEntry.DocsIfUpChannelPreEqEnable})

    docsIfUpstreamChannelEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfUpstreamChannelEntry.EntityData)
}

// DOCSIFMIB_DocsIfQosProfileTable
// Describes the attributes for each class of service.
type DOCSIFMIB_DocsIfQosProfileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes the attributes for a single class of service.  If implemented as
    // read-create in the Cable Modem Termination System, creation of entries in
    // this table is controlled by the value of docsIfCmtsQosProfilePermissions. 
    // If implemented as read-only, entries are created based on information in
    // REG-REQ MAC messages received from Cable Modems (Cable Modem Termination
    // System implementation), or based on information extracted from the TFTP
    // option file (Cable Modem implementation). In the Cable Modem Termination
    // system, read-only entries are removed if no longer referenced by
    // docsIfCmtsServiceTable.  An entry in this table must not be removed while
    // it is referenced by an entry in docsIfCmServiceTable (Cable Modem) or
    // docsIfCmtsServiceTable (Cable Modem Termination System).  An entry in this
    // table should not be changeable while it is referenced by an entry in
    // docsIfCmtsServiceTable.  If this table is created automatically, there
    // should only be a single entry for each Class of Service. Multiple entries
    // with the same Class of Service parameters are not recommended. The type is
    // slice of DOCSIFMIB_DocsIfQosProfileTable_DocsIfQosProfileEntry.
    DocsIfQosProfileEntry []*DOCSIFMIB_DocsIfQosProfileTable_DocsIfQosProfileEntry
}

func (docsIfQosProfileTable *DOCSIFMIB_DocsIfQosProfileTable) GetEntityData() *types.CommonEntityData {
    docsIfQosProfileTable.EntityData.YFilter = docsIfQosProfileTable.YFilter
    docsIfQosProfileTable.EntityData.YangName = "docsIfQosProfileTable"
    docsIfQosProfileTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfQosProfileTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfQosProfileTable.EntityData.SegmentPath = "docsIfQosProfileTable"
    docsIfQosProfileTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfQosProfileTable.EntityData.SegmentPath
    docsIfQosProfileTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfQosProfileTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfQosProfileTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfQosProfileTable.EntityData.Children = types.NewOrderedMap()
    docsIfQosProfileTable.EntityData.Children.Append("docsIfQosProfileEntry", types.YChild{"DocsIfQosProfileEntry", nil})
    for i := range docsIfQosProfileTable.DocsIfQosProfileEntry {
        docsIfQosProfileTable.EntityData.Children.Append(types.GetSegmentPath(docsIfQosProfileTable.DocsIfQosProfileEntry[i]), types.YChild{"DocsIfQosProfileEntry", docsIfQosProfileTable.DocsIfQosProfileEntry[i]})
    }
    docsIfQosProfileTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfQosProfileTable.EntityData.YListKeys = []string {}

    return &(docsIfQosProfileTable.EntityData)
}

// DOCSIFMIB_DocsIfQosProfileTable_DocsIfQosProfileEntry
// Describes the attributes for a single class of service.
// 
// If implemented as read-create in the Cable Modem
// Termination System, creation of entries in this table is
// controlled by the value of docsIfCmtsQosProfilePermissions.
// 
// If implemented as read-only, entries are created based
// on information in REG-REQ MAC messages received from
// Cable Modems (Cable Modem Termination System
// implementation), or based on information extracted from
// the TFTP option file (Cable Modem implementation).
// In the Cable Modem Termination system, read-only entries
// are removed if no longer referenced by
// docsIfCmtsServiceTable.
// 
// An entry in this table must not be removed while it is
// referenced by an entry in docsIfCmServiceTable (Cable Modem)
// or docsIfCmtsServiceTable (Cable Modem Termination System).
// 
// An entry in this table should not be changeable while
// it is referenced by an entry in docsIfCmtsServiceTable.
// 
// If this table is created automatically, there should only
// be a single entry for each Class of Service. Multiple
// entries with the same Class of Service parameters are not
// recommended.
type DOCSIFMIB_DocsIfQosProfileTable_DocsIfQosProfileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value that uniquely identifies an entry
    // in the docsIfQosProfileTable. The type is interface{} with range: 1..16383.
    DocsIfQosProfIndex interface{}

    // A relative priority assigned to this service when allocating bandwidth.
    // Zero indicates lowest priority and seven indicates highest priority.
    // Interpretation of priority is device-specific. MUST NOT be changed while
    // this row is active. The type is interface{} with range: 0..7.
    DocsIfQosProfPriority interface{}

    // The maximum upstream bandwidth, in bits per second, allowed for a service
    // with this service class. Zero if there is no restriction of upstream
    // bandwidth. MUST NOT be changed while this row is active. The type is
    // interface{} with range: 0..100000000.
    DocsIfQosProfMaxUpBandwidth interface{}

    // Minimum guaranteed upstream bandwidth, in bits per second, allowed for a
    // service with this service class. MUST NOT be changed while this row is
    // active. The type is interface{} with range: 0..100000000.
    DocsIfQosProfGuarUpBandwidth interface{}

    // The maximum downstream bandwidth, in bits per second, allowed for a service
    // with this service class. Zero if there is no restriction of downstream
    // bandwidth. MUST NOT be changed while this row is active. The type is
    // interface{} with range: 0..100000000.
    DocsIfQosProfMaxDownBandwidth interface{}

    // The maximum number of mini-slots that may be requested for a single
    // upstream transmission. A value of zero means there is no limit. MUST NOT be
    // changed while this row is active. This object has been deprecated and
    // replaced by  docsIfQosProfMaxTransmitBurst, to fix a mismatch of the units
    // and value range with respect to the DOCSIS Maximum Upstream Channel
    // Transmit Burst Configuration Setting. The type is interface{} with range:
    // 0..255.
    DocsIfQosProfMaxTxBurst interface{}

    // Indicates whether Baseline Privacy is enabled for this service class. MUST
    // NOT be changed while this row is active. The type is bool.
    DocsIfQosProfBaselinePrivacy interface{}

    // This is object is to used to create or delete rows in this table.  This
    // object MUST NOT be changed from active while the row is referenced by the
    // any entry in either docsIfCmServiceTable (on the CM), or the
    // docsIfCmtsServiceTable (on the CMTS). The type is RowStatus.
    DocsIfQosProfStatus interface{}

    // The maximum number of bytes that may be requested for a  single upstream
    // transmission. A value of zero means there  is no limit. Note: This value
    // does not include any  physical layer overhead. MUST NOT be changed while
    // this row is active. The type is interface{} with range: 0..65535.
    DocsIfQosProfMaxTransmitBurst interface{}
}

func (docsIfQosProfileEntry *DOCSIFMIB_DocsIfQosProfileTable_DocsIfQosProfileEntry) GetEntityData() *types.CommonEntityData {
    docsIfQosProfileEntry.EntityData.YFilter = docsIfQosProfileEntry.YFilter
    docsIfQosProfileEntry.EntityData.YangName = "docsIfQosProfileEntry"
    docsIfQosProfileEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfQosProfileEntry.EntityData.ParentYangName = "docsIfQosProfileTable"
    docsIfQosProfileEntry.EntityData.SegmentPath = "docsIfQosProfileEntry" + types.AddKeyToken(docsIfQosProfileEntry.DocsIfQosProfIndex, "docsIfQosProfIndex")
    docsIfQosProfileEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfQosProfileTable/" + docsIfQosProfileEntry.EntityData.SegmentPath
    docsIfQosProfileEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfQosProfileEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfQosProfileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfQosProfileEntry.EntityData.Children = types.NewOrderedMap()
    docsIfQosProfileEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfIndex", types.YLeaf{"DocsIfQosProfIndex", docsIfQosProfileEntry.DocsIfQosProfIndex})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfPriority", types.YLeaf{"DocsIfQosProfPriority", docsIfQosProfileEntry.DocsIfQosProfPriority})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfMaxUpBandwidth", types.YLeaf{"DocsIfQosProfMaxUpBandwidth", docsIfQosProfileEntry.DocsIfQosProfMaxUpBandwidth})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfGuarUpBandwidth", types.YLeaf{"DocsIfQosProfGuarUpBandwidth", docsIfQosProfileEntry.DocsIfQosProfGuarUpBandwidth})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfMaxDownBandwidth", types.YLeaf{"DocsIfQosProfMaxDownBandwidth", docsIfQosProfileEntry.DocsIfQosProfMaxDownBandwidth})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfMaxTxBurst", types.YLeaf{"DocsIfQosProfMaxTxBurst", docsIfQosProfileEntry.DocsIfQosProfMaxTxBurst})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfBaselinePrivacy", types.YLeaf{"DocsIfQosProfBaselinePrivacy", docsIfQosProfileEntry.DocsIfQosProfBaselinePrivacy})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfStatus", types.YLeaf{"DocsIfQosProfStatus", docsIfQosProfileEntry.DocsIfQosProfStatus})
    docsIfQosProfileEntry.EntityData.Leafs.Append("docsIfQosProfMaxTransmitBurst", types.YLeaf{"DocsIfQosProfMaxTransmitBurst", docsIfQosProfileEntry.DocsIfQosProfMaxTransmitBurst})

    docsIfQosProfileEntry.EntityData.YListKeys = []string {"DocsIfQosProfIndex"}

    return &(docsIfQosProfileEntry.EntityData)
}

// DOCSIFMIB_DocsIfSignalQualityTable
// At the CM, describes the PHY signal quality of downstream
// channels. At the CMTS, describes the PHY signal quality of
// upstream channels. At the CMTS, this table may exclude
// contention intervals.
type DOCSIFMIB_DocsIfSignalQualityTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // At the CM, describes the PHY characteristics of a downstream channel. At
    // the CMTS, describes the PHY signal quality of an upstream channel. An entry
    // in this table exists for each ifEntry with an ifType of
    // docsCableUpstreamChannel(205) for Cable Modem Termination Systems and
    // docsCableDownstream(128) for Cable Modems. The type is slice of
    // DOCSIFMIB_DocsIfSignalQualityTable_DocsIfSignalQualityEntry.
    DocsIfSignalQualityEntry []*DOCSIFMIB_DocsIfSignalQualityTable_DocsIfSignalQualityEntry
}

func (docsIfSignalQualityTable *DOCSIFMIB_DocsIfSignalQualityTable) GetEntityData() *types.CommonEntityData {
    docsIfSignalQualityTable.EntityData.YFilter = docsIfSignalQualityTable.YFilter
    docsIfSignalQualityTable.EntityData.YangName = "docsIfSignalQualityTable"
    docsIfSignalQualityTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfSignalQualityTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfSignalQualityTable.EntityData.SegmentPath = "docsIfSignalQualityTable"
    docsIfSignalQualityTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfSignalQualityTable.EntityData.SegmentPath
    docsIfSignalQualityTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfSignalQualityTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfSignalQualityTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfSignalQualityTable.EntityData.Children = types.NewOrderedMap()
    docsIfSignalQualityTable.EntityData.Children.Append("docsIfSignalQualityEntry", types.YChild{"DocsIfSignalQualityEntry", nil})
    for i := range docsIfSignalQualityTable.DocsIfSignalQualityEntry {
        docsIfSignalQualityTable.EntityData.Children.Append(types.GetSegmentPath(docsIfSignalQualityTable.DocsIfSignalQualityEntry[i]), types.YChild{"DocsIfSignalQualityEntry", docsIfSignalQualityTable.DocsIfSignalQualityEntry[i]})
    }
    docsIfSignalQualityTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfSignalQualityTable.EntityData.YListKeys = []string {}

    return &(docsIfSignalQualityTable.EntityData)
}

// DOCSIFMIB_DocsIfSignalQualityTable_DocsIfSignalQualityEntry
// At the CM, describes the PHY characteristics of a
// downstream channel. At the CMTS, describes the PHY signal
// quality of an upstream channel.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableUpstreamChannel(205) for Cable Modem Termination
// Systems and docsCableDownstream(128) for Cable Modems.
type DOCSIFMIB_DocsIfSignalQualityTable_DocsIfSignalQualityEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // true(1) if this CMTS includes contention intervals in the counters in this
    // table. Always false(2) for CMs. The type is bool.
    DocsIfSigQIncludesContention interface{}

    // Codewords received on this channel without error. This includes all
    // codewords, whether or not they were part of frames destined for this
    // device. The type is interface{} with range: 0..4294967295.
    DocsIfSigQUnerroreds interface{}

    // Codewords received on this channel with correctable errors. This includes
    // all codewords, whether or not they were part of frames destined for this
    // device. The type is interface{} with range: 0..4294967295.
    DocsIfSigQCorrecteds interface{}

    // Codewords received on this channel with uncorrectable errors. This includes
    // all codewords, whether or not they were part of frames destined for this
    // device. The type is interface{} with range: 0..4294967295.
    DocsIfSigQUncorrectables interface{}

    // Signal/Noise ratio as perceived for this channel. At the CM, describes the
    // Signal/Noise of the downstream channel.  At the CMTS, describes the average
    // Signal/Noise of the upstream channel. The type is interface{} with range:
    // -2147483648..2147483647. Units are dB.
    DocsIfSigQSignalNoise interface{}

    // Total microreflections including in-channel response as perceived on this
    // interface, measured in dBc below the signal level. This object is not
    // assumed to return an absolutely accurate value, but should give a rough
    // indication of microreflections received on this interface. It is up to the
    // implementer to provide information as accurate as possible. The type is
    // interface{} with range: 0..255. Units are dBc.
    DocsIfSigQMicroreflections interface{}

    // At the CM, returns the equalization data for the downstream channel. At the
    // CMTS, returns the average equalization data for the upstream channel.
    // Returns an empty string if the value is unknown or if there is no
    // equalization data available or defined. The type is string.
    DocsIfSigQEqualizationData interface{}

    // Codewords received on this channel without error. This includes all
    // codewords, whether or not they were part of frames destined for this
    // device. This is the 64 bit version of docsIfSigQUnerroreds. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfSigQExtUnerroreds interface{}

    // Codewords received on this channel with correctable errors. This includes
    // all codewords, whether or not they were part of frames destined for this
    // device. This is the 64 bit version of docsIfSigQCorrecteds. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfSigQExtCorrecteds interface{}

    // Codewords received on this channel with uncorrectable errors. This includes
    // all codewords, whether or not they were part of frames destined for this
    // device. This is the 64 bit version of docsIfSigQUncorrectables. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfSigQExtUncorrectables interface{}
}

func (docsIfSignalQualityEntry *DOCSIFMIB_DocsIfSignalQualityTable_DocsIfSignalQualityEntry) GetEntityData() *types.CommonEntityData {
    docsIfSignalQualityEntry.EntityData.YFilter = docsIfSignalQualityEntry.YFilter
    docsIfSignalQualityEntry.EntityData.YangName = "docsIfSignalQualityEntry"
    docsIfSignalQualityEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfSignalQualityEntry.EntityData.ParentYangName = "docsIfSignalQualityTable"
    docsIfSignalQualityEntry.EntityData.SegmentPath = "docsIfSignalQualityEntry" + types.AddKeyToken(docsIfSignalQualityEntry.IfIndex, "ifIndex")
    docsIfSignalQualityEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfSignalQualityTable/" + docsIfSignalQualityEntry.EntityData.SegmentPath
    docsIfSignalQualityEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfSignalQualityEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfSignalQualityEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfSignalQualityEntry.EntityData.Children = types.NewOrderedMap()
    docsIfSignalQualityEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfSignalQualityEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfSignalQualityEntry.IfIndex})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQIncludesContention", types.YLeaf{"DocsIfSigQIncludesContention", docsIfSignalQualityEntry.DocsIfSigQIncludesContention})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQUnerroreds", types.YLeaf{"DocsIfSigQUnerroreds", docsIfSignalQualityEntry.DocsIfSigQUnerroreds})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQCorrecteds", types.YLeaf{"DocsIfSigQCorrecteds", docsIfSignalQualityEntry.DocsIfSigQCorrecteds})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQUncorrectables", types.YLeaf{"DocsIfSigQUncorrectables", docsIfSignalQualityEntry.DocsIfSigQUncorrectables})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQSignalNoise", types.YLeaf{"DocsIfSigQSignalNoise", docsIfSignalQualityEntry.DocsIfSigQSignalNoise})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQMicroreflections", types.YLeaf{"DocsIfSigQMicroreflections", docsIfSignalQualityEntry.DocsIfSigQMicroreflections})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQEqualizationData", types.YLeaf{"DocsIfSigQEqualizationData", docsIfSignalQualityEntry.DocsIfSigQEqualizationData})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQExtUnerroreds", types.YLeaf{"DocsIfSigQExtUnerroreds", docsIfSignalQualityEntry.DocsIfSigQExtUnerroreds})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQExtCorrecteds", types.YLeaf{"DocsIfSigQExtCorrecteds", docsIfSignalQualityEntry.DocsIfSigQExtCorrecteds})
    docsIfSignalQualityEntry.EntityData.Leafs.Append("docsIfSigQExtUncorrectables", types.YLeaf{"DocsIfSigQExtUncorrectables", docsIfSignalQualityEntry.DocsIfSigQExtUncorrectables})

    docsIfSignalQualityEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfSignalQualityEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmMacTable
// Describes the attributes of each CM MAC interface,
// extending the information available from ifEntry.
type DOCSIFMIB_DocsIfCmMacTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing objects describing attributes of each MAC entry,
    // extending the information in ifEntry. An entry in this table exists for
    // each ifEntry with an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSIFMIB_DocsIfCmMacTable_DocsIfCmMacEntry.
    DocsIfCmMacEntry []*DOCSIFMIB_DocsIfCmMacTable_DocsIfCmMacEntry
}

func (docsIfCmMacTable *DOCSIFMIB_DocsIfCmMacTable) GetEntityData() *types.CommonEntityData {
    docsIfCmMacTable.EntityData.YFilter = docsIfCmMacTable.YFilter
    docsIfCmMacTable.EntityData.YangName = "docsIfCmMacTable"
    docsIfCmMacTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmMacTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmMacTable.EntityData.SegmentPath = "docsIfCmMacTable"
    docsIfCmMacTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmMacTable.EntityData.SegmentPath
    docsIfCmMacTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmMacTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmMacTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmMacTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmMacTable.EntityData.Children.Append("docsIfCmMacEntry", types.YChild{"DocsIfCmMacEntry", nil})
    for i := range docsIfCmMacTable.DocsIfCmMacEntry {
        docsIfCmMacTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmMacTable.DocsIfCmMacEntry[i]), types.YChild{"DocsIfCmMacEntry", docsIfCmMacTable.DocsIfCmMacEntry[i]})
    }
    docsIfCmMacTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmMacTable.EntityData.YListKeys = []string {}

    return &(docsIfCmMacTable.EntityData)
}

// DOCSIFMIB_DocsIfCmMacTable_DocsIfCmMacEntry
// An entry containing objects describing attributes of
// each MAC entry, extending the information in ifEntry.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableMaclayer(127).
type DOCSIFMIB_DocsIfCmMacTable_DocsIfCmMacEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Identifies the CMTS that is believed to control this MAC domain. At the CM,
    // this will be the source address from SYNC, MAP, and other MAC-layer
    // messages. If the CMTS is unknown, returns 00-00-00-00-00-00. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIfCmCmtsAddress interface{}

    // Identifies the capabilities of the MAC implementation at this interface.
    // Note that packet transmission is always supported. Therefore, there is no
    // specific bit required to explicitly indicate this capability. Note that
    // BITS objects are encoded most significant bit first. For example, if bit 1
    // is set, the value of this object is the octet string '40'H. The type is
    // map[string]bool.
    DocsIfCmCapabilities interface{}

    // Waiting time for a Ranging Response packet. The type is interface{} with
    // range: 0..4294967295.
    DocsIfCmRangingRespTimeout interface{}

    // Waiting time for a Ranging Response packet. The type is interface{} with
    // range: 0..2147483647.
    DocsIfCmRangingTimeout interface{}
}

func (docsIfCmMacEntry *DOCSIFMIB_DocsIfCmMacTable_DocsIfCmMacEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmMacEntry.EntityData.YFilter = docsIfCmMacEntry.YFilter
    docsIfCmMacEntry.EntityData.YangName = "docsIfCmMacEntry"
    docsIfCmMacEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmMacEntry.EntityData.ParentYangName = "docsIfCmMacTable"
    docsIfCmMacEntry.EntityData.SegmentPath = "docsIfCmMacEntry" + types.AddKeyToken(docsIfCmMacEntry.IfIndex, "ifIndex")
    docsIfCmMacEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmMacTable/" + docsIfCmMacEntry.EntityData.SegmentPath
    docsIfCmMacEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmMacEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmMacEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmMacEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmMacEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmMacEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmMacEntry.IfIndex})
    docsIfCmMacEntry.EntityData.Leafs.Append("docsIfCmCmtsAddress", types.YLeaf{"DocsIfCmCmtsAddress", docsIfCmMacEntry.DocsIfCmCmtsAddress})
    docsIfCmMacEntry.EntityData.Leafs.Append("docsIfCmCapabilities", types.YLeaf{"DocsIfCmCapabilities", docsIfCmMacEntry.DocsIfCmCapabilities})
    docsIfCmMacEntry.EntityData.Leafs.Append("docsIfCmRangingRespTimeout", types.YLeaf{"DocsIfCmRangingRespTimeout", docsIfCmMacEntry.DocsIfCmRangingRespTimeout})
    docsIfCmMacEntry.EntityData.Leafs.Append("docsIfCmRangingTimeout", types.YLeaf{"DocsIfCmRangingTimeout", docsIfCmMacEntry.DocsIfCmRangingTimeout})

    docsIfCmMacEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmMacEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmStatusTable
// This table maintains a number of status objects
// and counters for Cable Modems.
type DOCSIFMIB_DocsIfCmStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of status objects and counters for a single MAC layer instance in a
    // Cable Modem. An entry in this table exists for each ifEntry with an ifType
    // of docsCableMaclayer(127). The type is slice of
    // DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry.
    DocsIfCmStatusEntry []*DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry
}

func (docsIfCmStatusTable *DOCSIFMIB_DocsIfCmStatusTable) GetEntityData() *types.CommonEntityData {
    docsIfCmStatusTable.EntityData.YFilter = docsIfCmStatusTable.YFilter
    docsIfCmStatusTable.EntityData.YangName = "docsIfCmStatusTable"
    docsIfCmStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmStatusTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmStatusTable.EntityData.SegmentPath = "docsIfCmStatusTable"
    docsIfCmStatusTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmStatusTable.EntityData.SegmentPath
    docsIfCmStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmStatusTable.EntityData.Children.Append("docsIfCmStatusEntry", types.YChild{"DocsIfCmStatusEntry", nil})
    for i := range docsIfCmStatusTable.DocsIfCmStatusEntry {
        docsIfCmStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmStatusTable.DocsIfCmStatusEntry[i]), types.YChild{"DocsIfCmStatusEntry", docsIfCmStatusTable.DocsIfCmStatusEntry[i]})
    }
    docsIfCmStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmStatusTable.EntityData.YListKeys = []string {}

    return &(docsIfCmStatusTable.EntityData)
}

// DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry
// A set of status objects and counters for a single MAC
// layer instance in a Cable Modem.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableMaclayer(127).
type DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Current Cable Modem connectivity state, as specified in the RF Interface
    // Specification. Interpretations for state values 1-12 are clearly outlined
    // in the Document [25] reference given below. As stated in the description
    // for object docsIfCmtsCmStatusValue, accessDenied(13)indicates the CMTS has
    // sent a Registration Aborted message to the CM. The type is
    // DocsIfCmStatusValue.
    DocsIfCmStatusValue interface{}

    // Status code for this Cable Modem as defined in the RF Interface
    // Specification. The status code consists of a single character indicating
    // error groups, followed by a two- or three-digit number indicating the
    // status condition. The type is string.
    DocsIfCmStatusCode interface{}

    // The operational transmit power for the attached upstream channel. The type
    // is interface{} with range: -2147483648..2147483647. Units are dBmV.
    DocsIfCmStatusTxPower interface{}

    // Number of times the CM reset or initialized this interface. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmStatusResets interface{}

    // Number of times the CM lost synchronization with the downstream channel.
    // The type is interface{} with range: 0..4294967295.
    DocsIfCmStatusLostSyncs interface{}

    // Number of times the CM received invalid MAP messages. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmStatusInvalidMaps interface{}

    // Number of times the CM received invalid UCD messages. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmStatusInvalidUcds interface{}

    // Number of times the CM received invalid ranging response messages. The type
    // is interface{} with range: 0..4294967295.
    DocsIfCmStatusInvalidRangingResponses interface{}

    // Number of times the CM received invalid registration response messages. The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmStatusInvalidRegistrationResponses interface{}

    // Number of times counter T1 expired in the CM. The type is interface{} with
    // range: 0..4294967295.
    DocsIfCmStatusT1Timeouts interface{}

    // Number of times counter T2 expired in the CM. The type is interface{} with
    // range: 0..4294967295.
    DocsIfCmStatusT2Timeouts interface{}

    // Number of times counter T3 expired in the CM. The type is interface{} with
    // range: 0..4294967295.
    DocsIfCmStatusT3Timeouts interface{}

    // Number of times counter T4 expired in the CM. The type is interface{} with
    // range: 0..4294967295.
    DocsIfCmStatusT4Timeouts interface{}

    // Number of times the ranging process was aborted by the CMTS. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmStatusRangingAborteds interface{}

    // Indication whether the device has registered using 1.0 Class of Service or
    // 1.1 Quality of Service. An unregistered CM should indicate 1.1 QOS for a 
    // docsIfDocsisBaseCapability value of Docsis 1.1/2.0. An unregistered 	  CM
    // should indicate 1.0 COS for a docsIfDocsisBaseCapability value  of Docsis
    // 1.0. This object mirrors docsIfCmDocsisOperMode from the docsIfExt mib. The
    // type is DocsisQosVersion.
    DocsIfCmStatusDocsisOperMode interface{}

    // Indicates modulation type status currently used by the CM. Since this
    // object specifically identifies PHY mode, the shared upstream channel type
    // is not permitted. The type is DocsisUpstreamTypeStatus.
    DocsIfCmStatusModulationType interface{}

    // Pre-equalization data for this CM after convolution with  data indicated in
    // the RNG-RSP. Returns an empty string if the value is unknown or if there 
    // is no equalization data available or defined. The value should  be
    // formatted as defined in the following REFERENCE. The type is string.
    DocsIfCmStatusEqualizationData interface{}
}

func (docsIfCmStatusEntry *DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmStatusEntry.EntityData.YFilter = docsIfCmStatusEntry.YFilter
    docsIfCmStatusEntry.EntityData.YangName = "docsIfCmStatusEntry"
    docsIfCmStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmStatusEntry.EntityData.ParentYangName = "docsIfCmStatusTable"
    docsIfCmStatusEntry.EntityData.SegmentPath = "docsIfCmStatusEntry" + types.AddKeyToken(docsIfCmStatusEntry.IfIndex, "ifIndex")
    docsIfCmStatusEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmStatusTable/" + docsIfCmStatusEntry.EntityData.SegmentPath
    docsIfCmStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmStatusEntry.IfIndex})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusValue", types.YLeaf{"DocsIfCmStatusValue", docsIfCmStatusEntry.DocsIfCmStatusValue})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusCode", types.YLeaf{"DocsIfCmStatusCode", docsIfCmStatusEntry.DocsIfCmStatusCode})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusTxPower", types.YLeaf{"DocsIfCmStatusTxPower", docsIfCmStatusEntry.DocsIfCmStatusTxPower})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusResets", types.YLeaf{"DocsIfCmStatusResets", docsIfCmStatusEntry.DocsIfCmStatusResets})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusLostSyncs", types.YLeaf{"DocsIfCmStatusLostSyncs", docsIfCmStatusEntry.DocsIfCmStatusLostSyncs})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusInvalidMaps", types.YLeaf{"DocsIfCmStatusInvalidMaps", docsIfCmStatusEntry.DocsIfCmStatusInvalidMaps})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusInvalidUcds", types.YLeaf{"DocsIfCmStatusInvalidUcds", docsIfCmStatusEntry.DocsIfCmStatusInvalidUcds})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusInvalidRangingResponses", types.YLeaf{"DocsIfCmStatusInvalidRangingResponses", docsIfCmStatusEntry.DocsIfCmStatusInvalidRangingResponses})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusInvalidRegistrationResponses", types.YLeaf{"DocsIfCmStatusInvalidRegistrationResponses", docsIfCmStatusEntry.DocsIfCmStatusInvalidRegistrationResponses})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusT1Timeouts", types.YLeaf{"DocsIfCmStatusT1Timeouts", docsIfCmStatusEntry.DocsIfCmStatusT1Timeouts})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusT2Timeouts", types.YLeaf{"DocsIfCmStatusT2Timeouts", docsIfCmStatusEntry.DocsIfCmStatusT2Timeouts})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusT3Timeouts", types.YLeaf{"DocsIfCmStatusT3Timeouts", docsIfCmStatusEntry.DocsIfCmStatusT3Timeouts})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusT4Timeouts", types.YLeaf{"DocsIfCmStatusT4Timeouts", docsIfCmStatusEntry.DocsIfCmStatusT4Timeouts})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusRangingAborteds", types.YLeaf{"DocsIfCmStatusRangingAborteds", docsIfCmStatusEntry.DocsIfCmStatusRangingAborteds})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusDocsisOperMode", types.YLeaf{"DocsIfCmStatusDocsisOperMode", docsIfCmStatusEntry.DocsIfCmStatusDocsisOperMode})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusModulationType", types.YLeaf{"DocsIfCmStatusModulationType", docsIfCmStatusEntry.DocsIfCmStatusModulationType})
    docsIfCmStatusEntry.EntityData.Leafs.Append("docsIfCmStatusEqualizationData", types.YLeaf{"DocsIfCmStatusEqualizationData", docsIfCmStatusEntry.DocsIfCmStatusEqualizationData})

    docsIfCmStatusEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmStatusEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue represents Aborted message to the CM.
type DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue string

const (
    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_other DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "other"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_notReady DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "notReady"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_notSynchronized DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "notSynchronized"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_phySynchronized DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "phySynchronized"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_usParametersAcquired DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "usParametersAcquired"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_rangingComplete DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "rangingComplete"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_ipComplete DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "ipComplete"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_todEstablished DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "todEstablished"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_securityEstablished DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "securityEstablished"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_paramTransferComplete DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "paramTransferComplete"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_registrationComplete DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "registrationComplete"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_operational DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "operational"

    DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue_accessDenied DOCSIFMIB_DocsIfCmStatusTable_DocsIfCmStatusEntry_DocsIfCmStatusValue = "accessDenied"
)

// DOCSIFMIB_DocsIfCmServiceTable
// Describes the attributes of each upstream service queue
// on a CM.
type DOCSIFMIB_DocsIfCmServiceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes the attributes of an upstream bandwidth service queue. An entry
    // in this table exists for each Service ID. The primary index is an ifIndex
    // with an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSIFMIB_DocsIfCmServiceTable_DocsIfCmServiceEntry.
    DocsIfCmServiceEntry []*DOCSIFMIB_DocsIfCmServiceTable_DocsIfCmServiceEntry
}

func (docsIfCmServiceTable *DOCSIFMIB_DocsIfCmServiceTable) GetEntityData() *types.CommonEntityData {
    docsIfCmServiceTable.EntityData.YFilter = docsIfCmServiceTable.YFilter
    docsIfCmServiceTable.EntityData.YangName = "docsIfCmServiceTable"
    docsIfCmServiceTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmServiceTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmServiceTable.EntityData.SegmentPath = "docsIfCmServiceTable"
    docsIfCmServiceTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmServiceTable.EntityData.SegmentPath
    docsIfCmServiceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmServiceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmServiceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmServiceTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmServiceTable.EntityData.Children.Append("docsIfCmServiceEntry", types.YChild{"DocsIfCmServiceEntry", nil})
    for i := range docsIfCmServiceTable.DocsIfCmServiceEntry {
        docsIfCmServiceTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmServiceTable.DocsIfCmServiceEntry[i]), types.YChild{"DocsIfCmServiceEntry", docsIfCmServiceTable.DocsIfCmServiceEntry[i]})
    }
    docsIfCmServiceTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmServiceTable.EntityData.YListKeys = []string {}

    return &(docsIfCmServiceTable.EntityData)
}

// DOCSIFMIB_DocsIfCmServiceTable_DocsIfCmServiceEntry
// Describes the attributes of an upstream bandwidth service
// queue.
// An entry in this table exists for each Service ID.
// The primary index is an ifIndex with an ifType of
// docsCableMaclayer(127).
type DOCSIFMIB_DocsIfCmServiceTable_DocsIfCmServiceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Identifies a service queue for upstream bandwidth.
    // The attributes of this service queue are shared between the CM and the
    // CMTS. The CMTS allocates upstream bandwidth to this service queue based on
    // requests from the CM and on the class of service associated with this
    // queue. The type is interface{} with range: 1..16383.
    DocsIfCmServiceId interface{}

    // The index in docsIfQosProfileTable describing the quality of service
    // attributes associated with this particular service. If no associated entry
    // in docsIfQosProfileTable exists, this object returns a value of zero. The
    // type is interface{} with range: 0..16383.
    DocsIfCmServiceQosProfile interface{}

    // The number of upstream mini-slots which have been used to transmit data
    // PDUs in immediate (contention) mode. This includes only those PDUs that are
    // presumed to have arrived at the headend (i.e., those which were explicitly
    // acknowledged.) It does not include retransmission attempts or mini-slots
    // used by Requests. The type is interface{} with range: 0..4294967295.
    DocsIfCmServiceTxSlotsImmed interface{}

    // The number of upstream mini-slots which have been used to transmit data
    // PDUs in dedicated mode (i.e., as a result of a unicast Data Grant). The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmServiceTxSlotsDed interface{}

    // The number of attempts to transmit data PDUs containing requests for
    // acknowledgment that did not result in acknowledgment. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmServiceTxRetries interface{}

    // The number of data PDUs transmission failures due to excessive retries
    // without acknowledgment. The type is interface{} with range: 0..4294967295.
    DocsIfCmServiceTxExceededs interface{}

    // The number of attempts to transmit bandwidth requests which did not result
    // in acknowledgment. The type is interface{} with range: 0..4294967295.
    DocsIfCmServiceRqRetries interface{}

    // The number of requests for bandwidth which failed due to excessive retries
    // without acknowledgment. The type is interface{} with range: 0..4294967295.
    DocsIfCmServiceRqExceededs interface{}

    // The number of upstream mini-slots which have been used to transmit data
    // PDUs in immediate (contention) mode. This includes only those PDUs that are
    // presumed to have arrived at the headend (i.e., those which were explicitly
    // acknowledged.) It does not include retransmission attempts or mini-slots
    // used by Requests. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmServiceExtTxSlotsImmed interface{}

    // The number of upstream mini-slots which have been used to transmit data
    // PDUs in dedicated mode (i.e., as a result of a unicast Data Grant). The
    // type is interface{} with range: 0..18446744073709551615.
    DocsIfCmServiceExtTxSlotsDed interface{}
}

func (docsIfCmServiceEntry *DOCSIFMIB_DocsIfCmServiceTable_DocsIfCmServiceEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmServiceEntry.EntityData.YFilter = docsIfCmServiceEntry.YFilter
    docsIfCmServiceEntry.EntityData.YangName = "docsIfCmServiceEntry"
    docsIfCmServiceEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmServiceEntry.EntityData.ParentYangName = "docsIfCmServiceTable"
    docsIfCmServiceEntry.EntityData.SegmentPath = "docsIfCmServiceEntry" + types.AddKeyToken(docsIfCmServiceEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIfCmServiceEntry.DocsIfCmServiceId, "docsIfCmServiceId")
    docsIfCmServiceEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmServiceTable/" + docsIfCmServiceEntry.EntityData.SegmentPath
    docsIfCmServiceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmServiceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmServiceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmServiceEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmServiceEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmServiceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmServiceEntry.IfIndex})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceId", types.YLeaf{"DocsIfCmServiceId", docsIfCmServiceEntry.DocsIfCmServiceId})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceQosProfile", types.YLeaf{"DocsIfCmServiceQosProfile", docsIfCmServiceEntry.DocsIfCmServiceQosProfile})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceTxSlotsImmed", types.YLeaf{"DocsIfCmServiceTxSlotsImmed", docsIfCmServiceEntry.DocsIfCmServiceTxSlotsImmed})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceTxSlotsDed", types.YLeaf{"DocsIfCmServiceTxSlotsDed", docsIfCmServiceEntry.DocsIfCmServiceTxSlotsDed})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceTxRetries", types.YLeaf{"DocsIfCmServiceTxRetries", docsIfCmServiceEntry.DocsIfCmServiceTxRetries})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceTxExceededs", types.YLeaf{"DocsIfCmServiceTxExceededs", docsIfCmServiceEntry.DocsIfCmServiceTxExceededs})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceRqRetries", types.YLeaf{"DocsIfCmServiceRqRetries", docsIfCmServiceEntry.DocsIfCmServiceRqRetries})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceRqExceededs", types.YLeaf{"DocsIfCmServiceRqExceededs", docsIfCmServiceEntry.DocsIfCmServiceRqExceededs})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceExtTxSlotsImmed", types.YLeaf{"DocsIfCmServiceExtTxSlotsImmed", docsIfCmServiceEntry.DocsIfCmServiceExtTxSlotsImmed})
    docsIfCmServiceEntry.EntityData.Leafs.Append("docsIfCmServiceExtTxSlotsDed", types.YLeaf{"DocsIfCmServiceExtTxSlotsDed", docsIfCmServiceEntry.DocsIfCmServiceExtTxSlotsDed})

    docsIfCmServiceEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIfCmServiceId"}

    return &(docsIfCmServiceEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsMacTable
// Describes the attributes of each CMTS MAC interface,
// extending the information available from ifEntry.
// Mandatory for all CMTS devices.
type DOCSIFMIB_DocsIfCmtsMacTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing objects describing attributes of each MAC entry,
    // extending the information in ifEntry. An entry in this table exists for
    // each ifEntry with an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSIFMIB_DocsIfCmtsMacTable_DocsIfCmtsMacEntry.
    DocsIfCmtsMacEntry []*DOCSIFMIB_DocsIfCmtsMacTable_DocsIfCmtsMacEntry
}

func (docsIfCmtsMacTable *DOCSIFMIB_DocsIfCmtsMacTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsMacTable.EntityData.YFilter = docsIfCmtsMacTable.YFilter
    docsIfCmtsMacTable.EntityData.YangName = "docsIfCmtsMacTable"
    docsIfCmtsMacTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsMacTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsMacTable.EntityData.SegmentPath = "docsIfCmtsMacTable"
    docsIfCmtsMacTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsMacTable.EntityData.SegmentPath
    docsIfCmtsMacTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsMacTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsMacTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsMacTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsMacTable.EntityData.Children.Append("docsIfCmtsMacEntry", types.YChild{"DocsIfCmtsMacEntry", nil})
    for i := range docsIfCmtsMacTable.DocsIfCmtsMacEntry {
        docsIfCmtsMacTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsMacTable.DocsIfCmtsMacEntry[i]), types.YChild{"DocsIfCmtsMacEntry", docsIfCmtsMacTable.DocsIfCmtsMacEntry[i]})
    }
    docsIfCmtsMacTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsMacTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsMacTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsMacTable_DocsIfCmtsMacEntry
// An entry containing objects describing attributes of each
// MAC entry, extending the information in ifEntry.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableMaclayer(127).
type DOCSIFMIB_DocsIfCmtsMacTable_DocsIfCmtsMacEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Identifies the capabilities of the CMTS MAC implementation at this
    // interface. Note that packet transmission is always supported. Therefore,
    // there is no specific bit required to explicitly indicate this capability.
    // Note that BITS objects are encoded most significant bit first. For example,
    // if bit 1 is set, the value of this object is the octet string '40'H. The
    // type is map[string]bool.
    DocsIfCmtsCapabilities interface{}

    // The interval between CMTS transmission of successive SYNC messages at this
    // interface. The type is interface{} with range: 1..200. Units are
    // Milliseconds.
    DocsIfCmtsSyncInterval interface{}

    // The interval between CMTS transmission of successive Upstream Channel
    // Descriptor messages for each upstream channel at this interface. The type
    // is interface{} with range: 1..2000. Units are Milliseconds.
    DocsIfCmtsUcdInterval interface{}

    // The maximum number of service IDs that may be simultaneously active. The
    // type is interface{} with range: 1..16383.
    DocsIfCmtsMaxServiceIds interface{}

    // The amount of time to elapse between each broadcast station maintenance
    // grant. Broadcast station maintenance grants are used to allow new cable
    // modems to join the network. Zero indicates that a vendor-specific algorithm
    // is used instead of a fixed time. Maximum amount of time permitted by the
    // specification is 2 seconds. The type is interface{} with range:
    // 0..4294967295.
    DocsIfCmtsInsertionInterval interface{}

    // The maximum number of attempts to make on invitations for ranging requests.
    // A value of zero means the system should attempt to range forever. The type
    // is interface{} with range: 0..1024.
    DocsIfCmtsInvitedRangingAttempts interface{}

    // The amount of time to elapse between each broadcast station maintenance
    // grant. Broadcast station maintenance grants are used to allow new cable
    // modems to join the network. Zero indicates that a vendor-specific algorithm
    // is used instead of a fixed time. Maximum amount of time permitted by the
    // specification is 2 seconds. The type is interface{} with range:
    // 0..2147483647.
    DocsIfCmtsInsertInterval interface{}

    // The storage type for this conceptual row. Entries with this object set to
    // permanent(4) do not require write operations for read-write objects. The
    // type is StorageType.
    DocsIfCmtsMacStorageType interface{}
}

func (docsIfCmtsMacEntry *DOCSIFMIB_DocsIfCmtsMacTable_DocsIfCmtsMacEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsMacEntry.EntityData.YFilter = docsIfCmtsMacEntry.YFilter
    docsIfCmtsMacEntry.EntityData.YangName = "docsIfCmtsMacEntry"
    docsIfCmtsMacEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsMacEntry.EntityData.ParentYangName = "docsIfCmtsMacTable"
    docsIfCmtsMacEntry.EntityData.SegmentPath = "docsIfCmtsMacEntry" + types.AddKeyToken(docsIfCmtsMacEntry.IfIndex, "ifIndex")
    docsIfCmtsMacEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsMacTable/" + docsIfCmtsMacEntry.EntityData.SegmentPath
    docsIfCmtsMacEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsMacEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsMacEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsMacEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsMacEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsMacEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsMacEntry.IfIndex})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsCapabilities", types.YLeaf{"DocsIfCmtsCapabilities", docsIfCmtsMacEntry.DocsIfCmtsCapabilities})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsSyncInterval", types.YLeaf{"DocsIfCmtsSyncInterval", docsIfCmtsMacEntry.DocsIfCmtsSyncInterval})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsUcdInterval", types.YLeaf{"DocsIfCmtsUcdInterval", docsIfCmtsMacEntry.DocsIfCmtsUcdInterval})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsMaxServiceIds", types.YLeaf{"DocsIfCmtsMaxServiceIds", docsIfCmtsMacEntry.DocsIfCmtsMaxServiceIds})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsInsertionInterval", types.YLeaf{"DocsIfCmtsInsertionInterval", docsIfCmtsMacEntry.DocsIfCmtsInsertionInterval})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsInvitedRangingAttempts", types.YLeaf{"DocsIfCmtsInvitedRangingAttempts", docsIfCmtsMacEntry.DocsIfCmtsInvitedRangingAttempts})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsInsertInterval", types.YLeaf{"DocsIfCmtsInsertInterval", docsIfCmtsMacEntry.DocsIfCmtsInsertInterval})
    docsIfCmtsMacEntry.EntityData.Leafs.Append("docsIfCmtsMacStorageType", types.YLeaf{"DocsIfCmtsMacStorageType", docsIfCmtsMacEntry.DocsIfCmtsMacStorageType})

    docsIfCmtsMacEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmtsMacEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsStatusTable
// For the MAC layer, this group maintains a number of
// status objects and counters.
type DOCSIFMIB_DocsIfCmtsStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status entry for a single MAC layer. An entry in this table exists for each
    // ifEntry with an ifType of docsCableMaclayer(127). The type is slice of
    // DOCSIFMIB_DocsIfCmtsStatusTable_DocsIfCmtsStatusEntry.
    DocsIfCmtsStatusEntry []*DOCSIFMIB_DocsIfCmtsStatusTable_DocsIfCmtsStatusEntry
}

func (docsIfCmtsStatusTable *DOCSIFMIB_DocsIfCmtsStatusTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsStatusTable.EntityData.YFilter = docsIfCmtsStatusTable.YFilter
    docsIfCmtsStatusTable.EntityData.YangName = "docsIfCmtsStatusTable"
    docsIfCmtsStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsStatusTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsStatusTable.EntityData.SegmentPath = "docsIfCmtsStatusTable"
    docsIfCmtsStatusTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsStatusTable.EntityData.SegmentPath
    docsIfCmtsStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsStatusTable.EntityData.Children.Append("docsIfCmtsStatusEntry", types.YChild{"DocsIfCmtsStatusEntry", nil})
    for i := range docsIfCmtsStatusTable.DocsIfCmtsStatusEntry {
        docsIfCmtsStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsStatusTable.DocsIfCmtsStatusEntry[i]), types.YChild{"DocsIfCmtsStatusEntry", docsIfCmtsStatusTable.DocsIfCmtsStatusEntry[i]})
    }
    docsIfCmtsStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsStatusTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsStatusTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsStatusTable_DocsIfCmtsStatusEntry
// Status entry for a single MAC layer.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableMaclayer(127).
type DOCSIFMIB_DocsIfCmtsStatusTable_DocsIfCmtsStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object counts invalid RNG-REQ messages received on this interface. The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusInvalidRangeReqs interface{}

    // This object counts ranging attempts that were explicitly aborted by the
    // CMTS. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusRangingAborteds interface{}

    // This object counts invalid REG-REQ messages received on this interface.
    // That is, syntax, out of range parameters, or erroneous requests. The type
    // is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusInvalidRegReqs interface{}

    // This object counts failed registration attempts. Included are
    // docsIfCmtsStatusInvalidRegReqs, authentication and class of  service
    // failures. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusFailedRegReqs interface{}

    // This object counts invalid data request messages received on this
    // interface. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusInvalidDataReqs interface{}

    // This object counts the number of times counter T5 expired on this
    // interface. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsStatusT5Timeouts interface{}
}

func (docsIfCmtsStatusEntry *DOCSIFMIB_DocsIfCmtsStatusTable_DocsIfCmtsStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsStatusEntry.EntityData.YFilter = docsIfCmtsStatusEntry.YFilter
    docsIfCmtsStatusEntry.EntityData.YangName = "docsIfCmtsStatusEntry"
    docsIfCmtsStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsStatusEntry.EntityData.ParentYangName = "docsIfCmtsStatusTable"
    docsIfCmtsStatusEntry.EntityData.SegmentPath = "docsIfCmtsStatusEntry" + types.AddKeyToken(docsIfCmtsStatusEntry.IfIndex, "ifIndex")
    docsIfCmtsStatusEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsStatusTable/" + docsIfCmtsStatusEntry.EntityData.SegmentPath
    docsIfCmtsStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsStatusEntry.IfIndex})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusInvalidRangeReqs", types.YLeaf{"DocsIfCmtsStatusInvalidRangeReqs", docsIfCmtsStatusEntry.DocsIfCmtsStatusInvalidRangeReqs})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusRangingAborteds", types.YLeaf{"DocsIfCmtsStatusRangingAborteds", docsIfCmtsStatusEntry.DocsIfCmtsStatusRangingAborteds})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusInvalidRegReqs", types.YLeaf{"DocsIfCmtsStatusInvalidRegReqs", docsIfCmtsStatusEntry.DocsIfCmtsStatusInvalidRegReqs})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusFailedRegReqs", types.YLeaf{"DocsIfCmtsStatusFailedRegReqs", docsIfCmtsStatusEntry.DocsIfCmtsStatusFailedRegReqs})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusInvalidDataReqs", types.YLeaf{"DocsIfCmtsStatusInvalidDataReqs", docsIfCmtsStatusEntry.DocsIfCmtsStatusInvalidDataReqs})
    docsIfCmtsStatusEntry.EntityData.Leafs.Append("docsIfCmtsStatusT5Timeouts", types.YLeaf{"DocsIfCmtsStatusT5Timeouts", docsIfCmtsStatusEntry.DocsIfCmtsStatusT5Timeouts})

    docsIfCmtsStatusEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmtsStatusEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsCmStatusTable
// A set of objects in the CMTS, maintained for each
// Cable Modem connected to this CMTS.
type DOCSIFMIB_DocsIfCmtsCmStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status information for a single Cable Modem. An entry in this table exists
    // for each Cable Modem that is connected to the CMTS implementing this table.
    // The type is slice of
    // DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry.
    DocsIfCmtsCmStatusEntry []*DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry
}

func (docsIfCmtsCmStatusTable *DOCSIFMIB_DocsIfCmtsCmStatusTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsCmStatusTable.EntityData.YFilter = docsIfCmtsCmStatusTable.YFilter
    docsIfCmtsCmStatusTable.EntityData.YangName = "docsIfCmtsCmStatusTable"
    docsIfCmtsCmStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsCmStatusTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsCmStatusTable.EntityData.SegmentPath = "docsIfCmtsCmStatusTable"
    docsIfCmtsCmStatusTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsCmStatusTable.EntityData.SegmentPath
    docsIfCmtsCmStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsCmStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsCmStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsCmStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsCmStatusTable.EntityData.Children.Append("docsIfCmtsCmStatusEntry", types.YChild{"DocsIfCmtsCmStatusEntry", nil})
    for i := range docsIfCmtsCmStatusTable.DocsIfCmtsCmStatusEntry {
        docsIfCmtsCmStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsCmStatusTable.DocsIfCmtsCmStatusEntry[i]), types.YChild{"DocsIfCmtsCmStatusEntry", docsIfCmtsCmStatusTable.DocsIfCmtsCmStatusEntry[i]})
    }
    docsIfCmtsCmStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsCmStatusTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsCmStatusTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry
// Status information for a single Cable Modem.
// An entry in this table exists for each Cable Modem
// that is connected to the CMTS implementing this table.
type DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index value to uniquely identify an entry in this
    // table. For an individual Cable Modem, this index value should not change
    // during CMTS uptime. The type is interface{} with range: 1..2147483647.
    DocsIfCmtsCmStatusIndex interface{}

    // MAC address of this Cable Modem. If the Cable Modem has multiple MAC
    // addresses, this is the MAC address associated with the Cable interface. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIfCmtsCmStatusMacAddress interface{}

    // IP address of this Cable Modem. If the Cable Modem has no IP address
    // assigned, or the IP address is unknown, this object returns a value of
    // 0.0.0.0. If the Cable Modem has multiple IP addresses, this object returns
    // the IP address associated with the Cable interface. This object has been
    // deprecated and replaced by docsIfCmtsCmStatusInetAddressType and
    // docsIfCmtsCmStatusInetAddress, to enable IPv6 addressing in the future. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DocsIfCmtsCmStatusIpAddress interface{}

    // IfIndex of the downstream channel this CM is connected to. If the
    // downstream channel is unknown, this object returns a value of zero. The
    // type is interface{} with range: 0..2147483647.
    DocsIfCmtsCmStatusDownChannelIfIndex interface{}

    // IfIndex of the upstream channel this CM is connected to. If the upstream
    // channel is unknown, this object returns a value of zero. The type is
    // interface{} with range: 0..2147483647.
    DocsIfCmtsCmStatusUpChannelIfIndex interface{}

    // The receive power as perceived for upstream data from this Cable Modem. If
    // the receive power is unknown, this object returns a value of zero. The type
    // is interface{} with range: -2147483648..2147483647. Units are dBmV.
    DocsIfCmtsCmStatusRxPower interface{}

    // A measure of the current round trip time for this CM. Used for timing of CM
    // upstream transmissions to ensure synchronized arrivals at the CMTS. Units
    // are in terms of 6.25 microseconds/(64*256). Returns zero if the value is
    // unknown. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsCmStatusTimingOffset interface{}

    // Equalization data for this CM. Returns an empty string if the value is
    // unknown or if there is no equalization data available or defined. The type
    // is string.
    DocsIfCmtsCmStatusEqualizationData interface{}

    // Current Cable Modem connectivity state, as specified in the RF Interface
    // Specification. Returned status information is the CM status as assumed by
    // the CMTS, and indicates the following events: other(1)    Any state other
    // than below. ranging(2)    The CMTS has received an Initial Ranging Request 
    // message from the CM, and the ranging process is not  yet     complete.
    // rangingAborted(3)    The CMTS has sent a Ranging Abort message to the CM.
    // rangingComplete(4)    The CMTS has sent a Ranging Complete message to the
    // CM. ipComplete(5)    The CMTS has received a DHCP reply message and
    // forwarded    it to the CM. registrationComplete(6)    The CMTS has sent a
    // Registration Response message to the CM. accessDenied(7)    The CMTS has
    // sent a Registration Aborted message    to the CM. --           deprecated
    // value --           operational(8) --              If Baseline Privacy is
    // enabled for the CM, the CMTS --              has completed Baseline Privacy
    // initialization. If Baseline --              Privacy is not enabled,
    // equivalent to registrationComplete. registeredBPIInitializing(9)   
    // Baseline Privacy is enabled, CMTS is in the process of     completing the
    // Baseline Privacy initialization. This state     can last for a significant
    // time in the case of failures     during The process. After Baseline Privacy
    // initialization     Complete, the CMTS will report back the value   
    // registrationComplete(6).     The CMTS only needs to report states it is
    // able to detect. The type is DocsIfCmtsCmStatusValue.
    DocsIfCmtsCmStatusValue interface{}

    // Codewords received without error from this Cable Modem. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsCmStatusUnerroreds interface{}

    // Codewords received with correctable errors from this Cable Modem. The type
    // is interface{} with range: 0..4294967295.
    DocsIfCmtsCmStatusCorrecteds interface{}

    // Codewords received with uncorrectable errors from this Cable Modem. The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmtsCmStatusUncorrectables interface{}

    // Signal/Noise ratio as perceived for upstream data from this Cable Modem. If
    // the Signal/Noise is unknown, this object returns a value of zero. The type
    // is interface{} with range: -2147483648..2147483647. Units are dB.
    DocsIfCmtsCmStatusSignalNoise interface{}

    // Total microreflections including in-channel response as perceived on this
    // interface, measured in dBc below the signal level. This object is not
    // assumed to return an absolutely accurate value, but should give a rough
    // indication of microreflections received on this interface. It is up to the
    // implementer to provide information as accurate as possible. The type is
    // interface{} with range: 0..255. Units are dBc.
    DocsIfCmtsCmStatusMicroreflections interface{}

    // Codewords received without error from this Cable Modem. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfCmtsCmStatusExtUnerroreds interface{}

    // Codewords received with correctable errors from this Cable Modem. The type
    // is interface{} with range: 0..18446744073709551615.
    DocsIfCmtsCmStatusExtCorrecteds interface{}

    // Codewords received with uncorrectable errors from this Cable Modem. The
    // type is interface{} with range: 0..18446744073709551615.
    DocsIfCmtsCmStatusExtUncorrectables interface{}

    // Indication whether the CM has registered using 1.0 Class of Service or 1.1
    // Quality of Service. This object mirrors docsIfCmtsCmStatusDocsisMode from
    // the  docsIfExt mib. The type is DocsisQosVersion.
    DocsIfCmtsCmStatusDocsisRegMode interface{}

    // Indicates modulation type currently used by the CM. Since this object
    // specifically identifies PHY mode, the shared type is not permitted. If the
    // upstream channel is unknown,  this object returns a value of zero. The type
    // is DocsisUpstreamTypeStatus.
    DocsIfCmtsCmStatusModulationType interface{}

    // The type of internet address of docsIfCmtsCmStatusInetAddress. If the cable
    // modem Internet address is unassigned or unknown, then the value of this
    // object is unknown(0). The type is InetAddressType.
    DocsIfCmtsCmStatusInetAddressType interface{}

    // Internet address of this Cable Modem. If the Cable Modem has no Internet
    // address assigned, or the Internet address is unknown, the value of this
    // object is the empty string. If the Cable Modem has multiple Internet
    // addresses, this object returns the Internet address associated with the
    // Cable (i.e. RF MAC) interface. The type is string with length: 0..255.
    DocsIfCmtsCmStatusInetAddress interface{}

    // The value of sysUpTime when docsIfCmtsCmStatusValue was last updated. The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmtsCmStatusValueLastUpdate interface{}
}

func (docsIfCmtsCmStatusEntry *DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsCmStatusEntry.EntityData.YFilter = docsIfCmtsCmStatusEntry.YFilter
    docsIfCmtsCmStatusEntry.EntityData.YangName = "docsIfCmtsCmStatusEntry"
    docsIfCmtsCmStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsCmStatusEntry.EntityData.ParentYangName = "docsIfCmtsCmStatusTable"
    docsIfCmtsCmStatusEntry.EntityData.SegmentPath = "docsIfCmtsCmStatusEntry" + types.AddKeyToken(docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusIndex, "docsIfCmtsCmStatusIndex")
    docsIfCmtsCmStatusEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsCmStatusTable/" + docsIfCmtsCmStatusEntry.EntityData.SegmentPath
    docsIfCmtsCmStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsCmStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsCmStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsCmStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsCmStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIndex", types.YLeaf{"DocsIfCmtsCmStatusIndex", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusIndex})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusMacAddress", types.YLeaf{"DocsIfCmtsCmStatusMacAddress", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusMacAddress})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusIpAddress", types.YLeaf{"DocsIfCmtsCmStatusIpAddress", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusIpAddress})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusDownChannelIfIndex", types.YLeaf{"DocsIfCmtsCmStatusDownChannelIfIndex", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusDownChannelIfIndex})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusUpChannelIfIndex", types.YLeaf{"DocsIfCmtsCmStatusUpChannelIfIndex", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusUpChannelIfIndex})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusRxPower", types.YLeaf{"DocsIfCmtsCmStatusRxPower", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusRxPower})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusTimingOffset", types.YLeaf{"DocsIfCmtsCmStatusTimingOffset", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusTimingOffset})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusEqualizationData", types.YLeaf{"DocsIfCmtsCmStatusEqualizationData", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusEqualizationData})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusValue", types.YLeaf{"DocsIfCmtsCmStatusValue", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusValue})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusUnerroreds", types.YLeaf{"DocsIfCmtsCmStatusUnerroreds", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusUnerroreds})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusCorrecteds", types.YLeaf{"DocsIfCmtsCmStatusCorrecteds", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusCorrecteds})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusUncorrectables", types.YLeaf{"DocsIfCmtsCmStatusUncorrectables", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusUncorrectables})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusSignalNoise", types.YLeaf{"DocsIfCmtsCmStatusSignalNoise", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusSignalNoise})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusMicroreflections", types.YLeaf{"DocsIfCmtsCmStatusMicroreflections", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusMicroreflections})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusExtUnerroreds", types.YLeaf{"DocsIfCmtsCmStatusExtUnerroreds", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusExtUnerroreds})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusExtCorrecteds", types.YLeaf{"DocsIfCmtsCmStatusExtCorrecteds", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusExtCorrecteds})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusExtUncorrectables", types.YLeaf{"DocsIfCmtsCmStatusExtUncorrectables", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusExtUncorrectables})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusDocsisRegMode", types.YLeaf{"DocsIfCmtsCmStatusDocsisRegMode", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusDocsisRegMode})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusModulationType", types.YLeaf{"DocsIfCmtsCmStatusModulationType", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusModulationType})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusInetAddressType", types.YLeaf{"DocsIfCmtsCmStatusInetAddressType", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusInetAddressType})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusInetAddress", types.YLeaf{"DocsIfCmtsCmStatusInetAddress", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusInetAddress})
    docsIfCmtsCmStatusEntry.EntityData.Leafs.Append("docsIfCmtsCmStatusValueLastUpdate", types.YLeaf{"DocsIfCmtsCmStatusValueLastUpdate", docsIfCmtsCmStatusEntry.DocsIfCmtsCmStatusValueLastUpdate})

    docsIfCmtsCmStatusEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmStatusIndex"}

    return &(docsIfCmtsCmStatusEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue represents    The CMTS only needs to report states it is able to detect.
type DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue string

const (
    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_other DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "other"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_ranging DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "ranging"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_rangingAborted DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "rangingAborted"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_rangingComplete DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "rangingComplete"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_ipComplete DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "ipComplete"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_registrationComplete DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "registrationComplete"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_accessDenied DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "accessDenied"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_operational DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "operational"

    DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue_registeredBPIInitializing DOCSIFMIB_DocsIfCmtsCmStatusTable_DocsIfCmtsCmStatusEntry_DocsIfCmtsCmStatusValue = "registeredBPIInitializing"
)

// DOCSIFMIB_DocsIfCmtsServiceTable
// Describes the attributes of upstream service queues
// in a Cable Modem Termination System.
type DOCSIFMIB_DocsIfCmtsServiceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes the attributes of a single upstream bandwidth service queue.
    // Entries in this table exist for each ifEntry with an ifType of
    // docsCableMaclayer(127), and for each service queue (Service ID) within this
    // MAC layer. Entries in this table are created with the creation of
    // individual Service IDs by the MAC layer and removed when a Service ID is
    // removed. The type is slice of
    // DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry.
    DocsIfCmtsServiceEntry []*DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry
}

func (docsIfCmtsServiceTable *DOCSIFMIB_DocsIfCmtsServiceTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsServiceTable.EntityData.YFilter = docsIfCmtsServiceTable.YFilter
    docsIfCmtsServiceTable.EntityData.YangName = "docsIfCmtsServiceTable"
    docsIfCmtsServiceTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsServiceTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsServiceTable.EntityData.SegmentPath = "docsIfCmtsServiceTable"
    docsIfCmtsServiceTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsServiceTable.EntityData.SegmentPath
    docsIfCmtsServiceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsServiceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsServiceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsServiceTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsServiceTable.EntityData.Children.Append("docsIfCmtsServiceEntry", types.YChild{"DocsIfCmtsServiceEntry", nil})
    for i := range docsIfCmtsServiceTable.DocsIfCmtsServiceEntry {
        docsIfCmtsServiceTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsServiceTable.DocsIfCmtsServiceEntry[i]), types.YChild{"DocsIfCmtsServiceEntry", docsIfCmtsServiceTable.DocsIfCmtsServiceEntry[i]})
    }
    docsIfCmtsServiceTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsServiceTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsServiceTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry
// Describes the attributes of a single upstream bandwidth
// service queue.
// Entries in this table exist for each ifEntry with an
// ifType of docsCableMaclayer(127), and for each service
// queue (Service ID) within this MAC layer.
// Entries in this table are created with the creation of
// individual Service IDs by the MAC layer and removed
// when a Service ID is removed.
type DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Identifies a service queue for upstream bandwidth.
    // The attributes of this service queue are shared between the Cable Modem and
    // the Cable Modem Termination System. The CMTS allocates upstream bandwidth
    // to this service queue based on requests from the CM and on the class of
    // service associated with this queue. The type is interface{} with range:
    // 1..16383.
    DocsIfCmtsServiceId interface{}

    // Pointer to an entry in docsIfCmtsCmStatusTable identifying the Cable Modem
    // using this Service Queue. If multiple Cable Modems are using this Service
    // Queue, the value of this object is zero. This object has been deprecated
    // and replaced by docsIfCmtsServiceNewCmStatusIndex, to fix a mismatch of the
    // value range with respect to docsIfCmtsCmStatusIndex (1..2147483647). The
    // type is interface{} with range: 0..2147483647.
    DocsIfCmtsServiceCmStatusIndex interface{}

    // Allows a service class for a particular modem to be suppressed,
    // (re-)enabled, or deleted altogether. The type is
    // DocsIfCmtsServiceAdminStatus.
    DocsIfCmtsServiceAdminStatus interface{}

    // The index in docsIfQosProfileTable describing the quality of service
    // attributes associated with this particular service. If no associated
    // docsIfQosProfileTable entry exists, this object returns a value of zero.
    // The type is interface{} with range: 0..16383.
    DocsIfCmtsServiceQosProfile interface{}

    // The value of sysUpTime when this entry was created. The type is interface{}
    // with range: 0..4294967295.
    DocsIfCmtsServiceCreateTime interface{}

    // The cumulative number of Packet Data octets received on this Service ID.
    // The count does not include the size of the Cable MAC header. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsServiceInOctets interface{}

    // The cumulative number of Packet Data packets received on this Service ID.
    // The type is interface{} with range: 0..4294967295.
    DocsIfCmtsServiceInPackets interface{}

    // Pointer (via docsIfCmtsCmStatusIndex) to an entry in
    // docsIfCmtsCmStatusTable identifying the Cable Modem using this Service
    // Queue. If multiple Cable Modems are using this Service Queue, the value of
    // this object is zero. The type is interface{} with range: 0..2147483647.
    DocsIfCmtsServiceNewCmStatusIndex interface{}
}

func (docsIfCmtsServiceEntry *DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsServiceEntry.EntityData.YFilter = docsIfCmtsServiceEntry.YFilter
    docsIfCmtsServiceEntry.EntityData.YangName = "docsIfCmtsServiceEntry"
    docsIfCmtsServiceEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsServiceEntry.EntityData.ParentYangName = "docsIfCmtsServiceTable"
    docsIfCmtsServiceEntry.EntityData.SegmentPath = "docsIfCmtsServiceEntry" + types.AddKeyToken(docsIfCmtsServiceEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIfCmtsServiceEntry.DocsIfCmtsServiceId, "docsIfCmtsServiceId")
    docsIfCmtsServiceEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsServiceTable/" + docsIfCmtsServiceEntry.EntityData.SegmentPath
    docsIfCmtsServiceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsServiceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsServiceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsServiceEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsServiceEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsServiceEntry.IfIndex})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceId", types.YLeaf{"DocsIfCmtsServiceId", docsIfCmtsServiceEntry.DocsIfCmtsServiceId})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceCmStatusIndex", types.YLeaf{"DocsIfCmtsServiceCmStatusIndex", docsIfCmtsServiceEntry.DocsIfCmtsServiceCmStatusIndex})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceAdminStatus", types.YLeaf{"DocsIfCmtsServiceAdminStatus", docsIfCmtsServiceEntry.DocsIfCmtsServiceAdminStatus})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceQosProfile", types.YLeaf{"DocsIfCmtsServiceQosProfile", docsIfCmtsServiceEntry.DocsIfCmtsServiceQosProfile})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceCreateTime", types.YLeaf{"DocsIfCmtsServiceCreateTime", docsIfCmtsServiceEntry.DocsIfCmtsServiceCreateTime})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceInOctets", types.YLeaf{"DocsIfCmtsServiceInOctets", docsIfCmtsServiceEntry.DocsIfCmtsServiceInOctets})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceInPackets", types.YLeaf{"DocsIfCmtsServiceInPackets", docsIfCmtsServiceEntry.DocsIfCmtsServiceInPackets})
    docsIfCmtsServiceEntry.EntityData.Leafs.Append("docsIfCmtsServiceNewCmStatusIndex", types.YLeaf{"DocsIfCmtsServiceNewCmStatusIndex", docsIfCmtsServiceEntry.DocsIfCmtsServiceNewCmStatusIndex})

    docsIfCmtsServiceEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIfCmtsServiceId"}

    return &(docsIfCmtsServiceEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus represents suppressed, (re-)enabled, or deleted altogether.
type DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus string

const (
    DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus_enabled DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus = "enabled"

    DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus_disabled DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus = "disabled"

    DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus_destroyed DOCSIFMIB_DocsIfCmtsServiceTable_DocsIfCmtsServiceEntry_DocsIfCmtsServiceAdminStatus = "destroyed"
)

// DOCSIFMIB_DocsIfCmtsModulationTable
// Describes a modulation profile associated with one or more
// upstream channels.
type DOCSIFMIB_DocsIfCmtsModulationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describes a modulation profile for an Interval Usage Code for one or more
    // upstream channels. Entries in this table are created by the operator.
    // Initial default entries may be created at system initialization time. No
    // individual objects have to be specified in order to create an entry in this
    // table. Note that some objects do not have DEFVALs, but do have calculated
    // defaults and need not be specified during row creation. There is no
    // restriction on the changing of values in this table while their associated
    // rows are active. The type is slice of
    // DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry.
    DocsIfCmtsModulationEntry []*DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry
}

func (docsIfCmtsModulationTable *DOCSIFMIB_DocsIfCmtsModulationTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsModulationTable.EntityData.YFilter = docsIfCmtsModulationTable.YFilter
    docsIfCmtsModulationTable.EntityData.YangName = "docsIfCmtsModulationTable"
    docsIfCmtsModulationTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsModulationTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsModulationTable.EntityData.SegmentPath = "docsIfCmtsModulationTable"
    docsIfCmtsModulationTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsModulationTable.EntityData.SegmentPath
    docsIfCmtsModulationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsModulationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsModulationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsModulationTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsModulationTable.EntityData.Children.Append("docsIfCmtsModulationEntry", types.YChild{"DocsIfCmtsModulationEntry", nil})
    for i := range docsIfCmtsModulationTable.DocsIfCmtsModulationEntry {
        docsIfCmtsModulationTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsModulationTable.DocsIfCmtsModulationEntry[i]), types.YChild{"DocsIfCmtsModulationEntry", docsIfCmtsModulationTable.DocsIfCmtsModulationEntry[i]})
    }
    docsIfCmtsModulationTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsModulationTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsModulationTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry
// Describes a modulation profile for an Interval Usage Code
// for one or more upstream channels.
// Entries in this table are created by the operator. Initial
// default entries may be created at system initialization
// time. No individual objects have to be specified in order
// to create an entry in this table.
// Note that some objects do not have DEFVALs, but do have
// calculated defaults and need not be specified during row
// creation.
// There is no restriction on the changing of values in this
// table while their associated rows are active.
type DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index into the Channel Modulation table
    // representing a group of Interval Usage Codes, all associated with the same
    // channel. The type is interface{} with range: 1..2147483647.
    DocsIfCmtsModIndex interface{}

    // This attribute is a key. An index into the Channel Modulation table which,
    // when grouped with other Interval Usage Codes, fully instantiate all
    // modulation sets for a given upstream channel. The type is
    // DocsIfCmtsModIntervalUsageCode.
    DocsIfCmtsModIntervalUsageCode interface{}

    // Controls and reflects the status of rows in this table. The type is
    // RowStatus.
    DocsIfCmtsModControl interface{}

    // The modulation type used on this channel. Returns other(1) if the
    // modulation type is neither  qpsk, qam16, qam8, qam32, qam64 or qam128. Type
    // qam128 is used for SCDMA channels only. See the reference for the
    // modulation profiles implied by different modulation types. The type is
    // DocsIfCmtsModType.
    DocsIfCmtsModType interface{}

    // The preamble length for this modulation profile in bits. Default value is
    // the minimum needed by the implementation at the CMTS for the given
    // modulation profile. The type is interface{} with range: 0..1536.
    DocsIfCmtsModPreambleLen interface{}

    // Specifies whether or not differential encoding is used on this channel. The
    // type is bool.
    DocsIfCmtsModDifferentialEncoding interface{}

    // The number of correctable errored bytes (t) used in forward error
    // correction code. The value of 0 indicates no correction is employed. The
    // number of check bytes appended will be twice this value. The type is
    // interface{} with range: 0..16.
    DocsIfCmtsModFECErrorCorrection interface{}

    // The number of data bytes (k) in the forward error correction codeword. This
    // object is not used if docsIfCmtsModFECErrorCorrection is zero. The type is
    // interface{} with range: 1..255.
    DocsIfCmtsModFECCodewordLength interface{}

    // The 15 bit seed value for the scrambler polynomial. The type is interface{}
    // with range: 0..32767.
    DocsIfCmtsModScramblerSeed interface{}

    // The maximum number of mini-slots that can be transmitted during this
    // channel's burst time. Returns zero if the burst length is bounded by the
    // allocation MAP rather than this profile. Default value is 0 except for
    // shortData, where it is 8. The type is interface{} with range: 0..255.
    DocsIfCmtsModMaxBurstSize interface{}

    // The number of symbol-times which must follow the end of this channel's
    // burst. Default value is the minimum time needed by the implementation for
    // this modulation profile. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsModGuardTimeSize interface{}

    // Indicates if the last FEC codeword is truncated. The type is bool.
    DocsIfCmtsModLastCodewordShortened interface{}

    // Indicates if the scrambler is employed. The type is bool.
    DocsIfCmtsModScrambler interface{}

    // ATDMA Byte Interleaver Depth (Ir). This object returns 1 for               
    // non ATDMA profiles. . The type is interface{} with range: 0..4294967295.
    DocsIfCmtsModByteInterleaverDepth interface{}

    // ATDMA Byte Interleaver Block size (Br). This object returns  zero for non
    // ATDMA profiles . The type is interface{} with range: 0..4294967295.
    DocsIfCmtsModByteInterleaverBlockSize interface{}

    // Preamble type for DOCSIS 2.0 bursts. The value 'unknown(0)'  represents a
    // row entry consisting only of DOCSIS 1.x bursts. The type is
    // DocsIfCmtsModPreambleType.
    DocsIfCmtsModPreambleType interface{}

    // Trellis Code Modulation (TCM) On/Off. This value returns false for  non
    // S-CDMA profiles. The type is bool.
    DocsIfCmtsModTcmErrorCorrectionOn interface{}

    // S-CDMA Interleaver step size. This value returns zero  for non S-CDMA
    // profiles. The type is interface{} with range: 0..32.
    DocsIfCmtsModScdmaInterleaverStepSize interface{}

    // S-CDMA spreader. This value returns false for non S-CDMA profiles. Default
    // value for IUC 3 and 4 is OFF, for  all other IUCs it is ON. The type is
    // bool.
    DocsIfCmtsModScdmaSpreaderEnable interface{}

    // S-CDMA sub-frame size. This value returns zero  for non S-CDMA profiles.
    // The type is interface{} with range: 0..128.
    DocsIfCmtsModScdmaSubframeCodes interface{}

    // Describes the modulation channel type for this modulation entry. The type
    // is DocsisUpstreamType.
    DocsIfCmtsModChannelType interface{}
}

func (docsIfCmtsModulationEntry *DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsModulationEntry.EntityData.YFilter = docsIfCmtsModulationEntry.YFilter
    docsIfCmtsModulationEntry.EntityData.YangName = "docsIfCmtsModulationEntry"
    docsIfCmtsModulationEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsModulationEntry.EntityData.ParentYangName = "docsIfCmtsModulationTable"
    docsIfCmtsModulationEntry.EntityData.SegmentPath = "docsIfCmtsModulationEntry" + types.AddKeyToken(docsIfCmtsModulationEntry.DocsIfCmtsModIndex, "docsIfCmtsModIndex") + types.AddKeyToken(docsIfCmtsModulationEntry.DocsIfCmtsModIntervalUsageCode, "docsIfCmtsModIntervalUsageCode")
    docsIfCmtsModulationEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsModulationTable/" + docsIfCmtsModulationEntry.EntityData.SegmentPath
    docsIfCmtsModulationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsModulationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsModulationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsModulationEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsModulationEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModIndex", types.YLeaf{"DocsIfCmtsModIndex", docsIfCmtsModulationEntry.DocsIfCmtsModIndex})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModIntervalUsageCode", types.YLeaf{"DocsIfCmtsModIntervalUsageCode", docsIfCmtsModulationEntry.DocsIfCmtsModIntervalUsageCode})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModControl", types.YLeaf{"DocsIfCmtsModControl", docsIfCmtsModulationEntry.DocsIfCmtsModControl})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModType", types.YLeaf{"DocsIfCmtsModType", docsIfCmtsModulationEntry.DocsIfCmtsModType})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModPreambleLen", types.YLeaf{"DocsIfCmtsModPreambleLen", docsIfCmtsModulationEntry.DocsIfCmtsModPreambleLen})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModDifferentialEncoding", types.YLeaf{"DocsIfCmtsModDifferentialEncoding", docsIfCmtsModulationEntry.DocsIfCmtsModDifferentialEncoding})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModFECErrorCorrection", types.YLeaf{"DocsIfCmtsModFECErrorCorrection", docsIfCmtsModulationEntry.DocsIfCmtsModFECErrorCorrection})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModFECCodewordLength", types.YLeaf{"DocsIfCmtsModFECCodewordLength", docsIfCmtsModulationEntry.DocsIfCmtsModFECCodewordLength})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModScramblerSeed", types.YLeaf{"DocsIfCmtsModScramblerSeed", docsIfCmtsModulationEntry.DocsIfCmtsModScramblerSeed})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModMaxBurstSize", types.YLeaf{"DocsIfCmtsModMaxBurstSize", docsIfCmtsModulationEntry.DocsIfCmtsModMaxBurstSize})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModGuardTimeSize", types.YLeaf{"DocsIfCmtsModGuardTimeSize", docsIfCmtsModulationEntry.DocsIfCmtsModGuardTimeSize})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModLastCodewordShortened", types.YLeaf{"DocsIfCmtsModLastCodewordShortened", docsIfCmtsModulationEntry.DocsIfCmtsModLastCodewordShortened})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModScrambler", types.YLeaf{"DocsIfCmtsModScrambler", docsIfCmtsModulationEntry.DocsIfCmtsModScrambler})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModByteInterleaverDepth", types.YLeaf{"DocsIfCmtsModByteInterleaverDepth", docsIfCmtsModulationEntry.DocsIfCmtsModByteInterleaverDepth})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModByteInterleaverBlockSize", types.YLeaf{"DocsIfCmtsModByteInterleaverBlockSize", docsIfCmtsModulationEntry.DocsIfCmtsModByteInterleaverBlockSize})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModPreambleType", types.YLeaf{"DocsIfCmtsModPreambleType", docsIfCmtsModulationEntry.DocsIfCmtsModPreambleType})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModTcmErrorCorrectionOn", types.YLeaf{"DocsIfCmtsModTcmErrorCorrectionOn", docsIfCmtsModulationEntry.DocsIfCmtsModTcmErrorCorrectionOn})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModScdmaInterleaverStepSize", types.YLeaf{"DocsIfCmtsModScdmaInterleaverStepSize", docsIfCmtsModulationEntry.DocsIfCmtsModScdmaInterleaverStepSize})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModScdmaSpreaderEnable", types.YLeaf{"DocsIfCmtsModScdmaSpreaderEnable", docsIfCmtsModulationEntry.DocsIfCmtsModScdmaSpreaderEnable})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModScdmaSubframeCodes", types.YLeaf{"DocsIfCmtsModScdmaSubframeCodes", docsIfCmtsModulationEntry.DocsIfCmtsModScdmaSubframeCodes})
    docsIfCmtsModulationEntry.EntityData.Leafs.Append("docsIfCmtsModChannelType", types.YLeaf{"DocsIfCmtsModChannelType", docsIfCmtsModulationEntry.DocsIfCmtsModChannelType})

    docsIfCmtsModulationEntry.EntityData.YListKeys = []string {"DocsIfCmtsModIndex", "DocsIfCmtsModIntervalUsageCode"}

    return &(docsIfCmtsModulationEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode represents channel.
type DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode string

const (
    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_request DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "request"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_requestData DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "requestData"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_initialRanging DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "initialRanging"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_periodicRanging DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "periodicRanging"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_shortData DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "shortData"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_longData DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "longData"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_advPhyShortData DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "advPhyShortData"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_advPhyLongData DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "advPhyLongData"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode_ugs DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModIntervalUsageCode = "ugs"
)

// DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType represents represents a row entry consisting only of DOCSIS 1.x bursts.
type DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType string

const (
    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType_unknown DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType = "unknown"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType_qpsk0 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType = "qpsk0"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType_qpsk1 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModPreambleType = "qpsk1"
)

// DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType represents implied by different modulation types.
type DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType string

const (
    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_other DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "other"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qpsk DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qpsk"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qam16 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qam16"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qam8 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qam8"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qam32 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qam32"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qam64 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qam64"

    DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType_qam128 DOCSIFMIB_DocsIfCmtsModulationTable_DocsIfCmtsModulationEntry_DocsIfCmtsModType = "qam128"
)

// DOCSIFMIB_DocsIfCmtsMacToCmTable
// This is a table to provide a quick access index into the
// docsIfCmtsCmStatusTable. There is exactly one row in this
// table for each row in the docsIfCmtsCmStatusTable. In
// general, the management station should use this table only
// to get a pointer into the docsIfCmtsCmStatusTable (which
// corresponds to the CM's RF interface MAC address), and
// should not iterate (e.g. GetNext through) this table.
type DOCSIFMIB_DocsIfCmtsMacToCmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the docsIfCmtsMacToCmTable. An entry in this table exists for each
    // Cable Modem that is connected to the CMTS implementing this table. The type
    // is slice of DOCSIFMIB_DocsIfCmtsMacToCmTable_DocsIfCmtsMacToCmEntry.
    DocsIfCmtsMacToCmEntry []*DOCSIFMIB_DocsIfCmtsMacToCmTable_DocsIfCmtsMacToCmEntry
}

func (docsIfCmtsMacToCmTable *DOCSIFMIB_DocsIfCmtsMacToCmTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsMacToCmTable.EntityData.YFilter = docsIfCmtsMacToCmTable.YFilter
    docsIfCmtsMacToCmTable.EntityData.YangName = "docsIfCmtsMacToCmTable"
    docsIfCmtsMacToCmTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsMacToCmTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsMacToCmTable.EntityData.SegmentPath = "docsIfCmtsMacToCmTable"
    docsIfCmtsMacToCmTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsMacToCmTable.EntityData.SegmentPath
    docsIfCmtsMacToCmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsMacToCmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsMacToCmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsMacToCmTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsMacToCmTable.EntityData.Children.Append("docsIfCmtsMacToCmEntry", types.YChild{"DocsIfCmtsMacToCmEntry", nil})
    for i := range docsIfCmtsMacToCmTable.DocsIfCmtsMacToCmEntry {
        docsIfCmtsMacToCmTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsMacToCmTable.DocsIfCmtsMacToCmEntry[i]), types.YChild{"DocsIfCmtsMacToCmEntry", docsIfCmtsMacToCmTable.DocsIfCmtsMacToCmEntry[i]})
    }
    docsIfCmtsMacToCmTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsMacToCmTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsMacToCmTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsMacToCmTable_DocsIfCmtsMacToCmEntry
// A row in the docsIfCmtsMacToCmTable.
// An entry in this table exists for each Cable Modem
// that is connected to the CMTS implementing this table.
type DOCSIFMIB_DocsIfCmtsMacToCmTable_DocsIfCmtsMacToCmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The RF side MAC address for the referenced CM.
    // (E.g. the interface on the CM that has docsCableMacLayer(127) as its
    // ifType. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIfCmtsCmMac interface{}

    // An row index into docsIfCmtsCmStatusTable. When queried with the correct
    // instance value (e.g. a CM's MAC address), returns the index in
    // docsIfCmtsCmStatusTable which represents that CM. The type is interface{}
    // with range: 1..2147483647.
    DocsIfCmtsCmPtr interface{}
}

func (docsIfCmtsMacToCmEntry *DOCSIFMIB_DocsIfCmtsMacToCmTable_DocsIfCmtsMacToCmEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsMacToCmEntry.EntityData.YFilter = docsIfCmtsMacToCmEntry.YFilter
    docsIfCmtsMacToCmEntry.EntityData.YangName = "docsIfCmtsMacToCmEntry"
    docsIfCmtsMacToCmEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsMacToCmEntry.EntityData.ParentYangName = "docsIfCmtsMacToCmTable"
    docsIfCmtsMacToCmEntry.EntityData.SegmentPath = "docsIfCmtsMacToCmEntry" + types.AddKeyToken(docsIfCmtsMacToCmEntry.DocsIfCmtsCmMac, "docsIfCmtsCmMac")
    docsIfCmtsMacToCmEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsMacToCmTable/" + docsIfCmtsMacToCmEntry.EntityData.SegmentPath
    docsIfCmtsMacToCmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsMacToCmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsMacToCmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsMacToCmEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsMacToCmEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsMacToCmEntry.EntityData.Leafs.Append("docsIfCmtsCmMac", types.YLeaf{"DocsIfCmtsCmMac", docsIfCmtsMacToCmEntry.DocsIfCmtsCmMac})
    docsIfCmtsMacToCmEntry.EntityData.Leafs.Append("docsIfCmtsCmPtr", types.YLeaf{"DocsIfCmtsCmPtr", docsIfCmtsMacToCmEntry.DocsIfCmtsCmPtr})

    docsIfCmtsMacToCmEntry.EntityData.YListKeys = []string {"DocsIfCmtsCmMac"}

    return &(docsIfCmtsMacToCmEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsChannelUtilizationTable
// Reports utilization statistics for attached upstream and
// downstream physical channels.
type DOCSIFMIB_DocsIfCmtsChannelUtilizationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Utilization statistics for a single upstream or downstream physical
    // channel. An entry exists in this table for each ifEntry with an ifType
    // equal to docsCableDownstreamInterface (128) or docsCableUpstreamInterface
    // (129). The type is slice of
    // DOCSIFMIB_DocsIfCmtsChannelUtilizationTable_DocsIfCmtsChannelUtilizationEntry.
    DocsIfCmtsChannelUtilizationEntry []*DOCSIFMIB_DocsIfCmtsChannelUtilizationTable_DocsIfCmtsChannelUtilizationEntry
}

func (docsIfCmtsChannelUtilizationTable *DOCSIFMIB_DocsIfCmtsChannelUtilizationTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsChannelUtilizationTable.EntityData.YFilter = docsIfCmtsChannelUtilizationTable.YFilter
    docsIfCmtsChannelUtilizationTable.EntityData.YangName = "docsIfCmtsChannelUtilizationTable"
    docsIfCmtsChannelUtilizationTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsChannelUtilizationTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsChannelUtilizationTable.EntityData.SegmentPath = "docsIfCmtsChannelUtilizationTable"
    docsIfCmtsChannelUtilizationTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsChannelUtilizationTable.EntityData.SegmentPath
    docsIfCmtsChannelUtilizationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsChannelUtilizationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsChannelUtilizationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsChannelUtilizationTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsChannelUtilizationTable.EntityData.Children.Append("docsIfCmtsChannelUtilizationEntry", types.YChild{"DocsIfCmtsChannelUtilizationEntry", nil})
    for i := range docsIfCmtsChannelUtilizationTable.DocsIfCmtsChannelUtilizationEntry {
        docsIfCmtsChannelUtilizationTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsChannelUtilizationTable.DocsIfCmtsChannelUtilizationEntry[i]), types.YChild{"DocsIfCmtsChannelUtilizationEntry", docsIfCmtsChannelUtilizationTable.DocsIfCmtsChannelUtilizationEntry[i]})
    }
    docsIfCmtsChannelUtilizationTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsChannelUtilizationTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsChannelUtilizationTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsChannelUtilizationTable_DocsIfCmtsChannelUtilizationEntry
// Utilization statistics for a single upstream or downstream
// physical channel. An entry exists in this table for each
// ifEntry with an ifType equal to docsCableDownstreamInterface
// (128) or docsCableUpstreamInterface (129).
type DOCSIFMIB_DocsIfCmtsChannelUtilizationTable_DocsIfCmtsChannelUtilizationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The secondary index into this table. Indicates the
    // IANA interface type associated with this physical channel. Only
    // docsCableDownstreamInterface (128) and docsCableUpstreamInterface (129) are
    // valid. The type is IANAifType.
    DocsIfCmtsChannelUtIfType interface{}

    // This attribute is a key. The tertiary index into this table. Indicates the
    // CMTS identifier for this physical channel. The type is interface{} with
    // range: 0..255.
    DocsIfCmtsChannelUtId interface{}

    // The calculated and truncated utilization index for this physical upstream
    // or downstream channel, accurate as of  the most recent
    // docsIfCmtsChannelUtilizationInterval.  Upstream Channel Utilization Index:
    // The upstream channel utilization index is expressed as a  percentage of
    // minislots utilized on the physical channel, regardless  of burst type. For
    // an Initial Maintenance region, the minislots  for the complete region are
    // considered utilized if the CMTS  received an upstream burst within the
    // region from any CM on the  physical channel.  For contention REQ and
    // REQ/DATA regions, the     minislots for a transmission opportunity within
    // the region are  considered utilized if the CMTS received an upstream burst
    // within  the opportunity from any CM on the physical channel. For all other 
    // regions, utilized minislots are those in which the CMTS granted bandwidth
    // to any unicast SID on the physical channel.  For an upstream interface that
    // has multiple logical upstream  channels enabled, the utilization index is a
    // weighted sum of  utilization indices for the logical channels. The weight
    // for  each utilization index is the percentage of upstream minislots 
    // allocated for the corresponding logical channel. Example:  If 75% of
    // bandwidth is allocated to the first logical channel  and 25% to the second,
    // and the utilization indices for each are  60 and 40 respectively, the
    // utilization index for the upstream  physical channel is (60 * 0.75) + (40 *
    // 0.25) = 55. This figure  applies to the most recent utilization interval.  
    // Downstream Channel Utilization Index: The downstream channel utilization
    // index is a percentage expressing  the ratio between bytes used to transmit
    // data versus the total number  of bytes transmitted in the raw bandwidth of
    // the MPEG channel. As with the upstream utilization index, the calculated
    // value represents  the most recent utilization interval. Formula: Downstream
    // utilization index =  (100 * (data bytes / raw bytes)) =  (100 * ((raw bytes
    // - stuffed bytes) / raw bytes))  Definitions: Data bytes: Number of bytes
    // transmitted as data in the                                        
    // docsIfCmtsChannelUtilizationInterval.  Stuffed bytes: Number of filler
    // bytes transmitted as non-data in the                
    // DocsIfCmtsChannelUtilizationInterval. Raw bandwidth: Total number of bytes
    // available for transmitting data,                not including bytes used
    // for headers and other overhead. Raw bytes: (raw bandwidth *
    // docsIfCmtsChannelUtilizationInterval). The type is interface{} with range:
    // 0..100. Units are percent.
    DocsIfCmtsChannelUtUtilization interface{}
}

func (docsIfCmtsChannelUtilizationEntry *DOCSIFMIB_DocsIfCmtsChannelUtilizationTable_DocsIfCmtsChannelUtilizationEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsChannelUtilizationEntry.EntityData.YFilter = docsIfCmtsChannelUtilizationEntry.YFilter
    docsIfCmtsChannelUtilizationEntry.EntityData.YangName = "docsIfCmtsChannelUtilizationEntry"
    docsIfCmtsChannelUtilizationEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsChannelUtilizationEntry.EntityData.ParentYangName = "docsIfCmtsChannelUtilizationTable"
    docsIfCmtsChannelUtilizationEntry.EntityData.SegmentPath = "docsIfCmtsChannelUtilizationEntry" + types.AddKeyToken(docsIfCmtsChannelUtilizationEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIfCmtsChannelUtilizationEntry.DocsIfCmtsChannelUtIfType, "docsIfCmtsChannelUtIfType") + types.AddKeyToken(docsIfCmtsChannelUtilizationEntry.DocsIfCmtsChannelUtId, "docsIfCmtsChannelUtId")
    docsIfCmtsChannelUtilizationEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsChannelUtilizationTable/" + docsIfCmtsChannelUtilizationEntry.EntityData.SegmentPath
    docsIfCmtsChannelUtilizationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsChannelUtilizationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsChannelUtilizationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsChannelUtilizationEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsChannelUtilizationEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsChannelUtilizationEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsChannelUtilizationEntry.IfIndex})
    docsIfCmtsChannelUtilizationEntry.EntityData.Leafs.Append("docsIfCmtsChannelUtIfType", types.YLeaf{"DocsIfCmtsChannelUtIfType", docsIfCmtsChannelUtilizationEntry.DocsIfCmtsChannelUtIfType})
    docsIfCmtsChannelUtilizationEntry.EntityData.Leafs.Append("docsIfCmtsChannelUtId", types.YLeaf{"DocsIfCmtsChannelUtId", docsIfCmtsChannelUtilizationEntry.DocsIfCmtsChannelUtId})
    docsIfCmtsChannelUtilizationEntry.EntityData.Leafs.Append("docsIfCmtsChannelUtUtilization", types.YLeaf{"DocsIfCmtsChannelUtUtilization", docsIfCmtsChannelUtilizationEntry.DocsIfCmtsChannelUtUtilization})

    docsIfCmtsChannelUtilizationEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIfCmtsChannelUtIfType", "DocsIfCmtsChannelUtId"}

    return &(docsIfCmtsChannelUtilizationEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsDownChannelCounterTable
// This table is implemented at the CMTS to collect downstream
// channel statistics for utilization calculations.
type DOCSIFMIB_DocsIfCmtsDownChannelCounterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry provides a list of traffic counters for a single downstream
    // channel. An entry in this table exists for each ifEntry with an ifType of
    // docsCableDownstream(128). The type is slice of
    // DOCSIFMIB_DocsIfCmtsDownChannelCounterTable_DocsIfCmtsDownChannelCounterEntry.
    DocsIfCmtsDownChannelCounterEntry []*DOCSIFMIB_DocsIfCmtsDownChannelCounterTable_DocsIfCmtsDownChannelCounterEntry
}

func (docsIfCmtsDownChannelCounterTable *DOCSIFMIB_DocsIfCmtsDownChannelCounterTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsDownChannelCounterTable.EntityData.YFilter = docsIfCmtsDownChannelCounterTable.YFilter
    docsIfCmtsDownChannelCounterTable.EntityData.YangName = "docsIfCmtsDownChannelCounterTable"
    docsIfCmtsDownChannelCounterTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsDownChannelCounterTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsDownChannelCounterTable.EntityData.SegmentPath = "docsIfCmtsDownChannelCounterTable"
    docsIfCmtsDownChannelCounterTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsDownChannelCounterTable.EntityData.SegmentPath
    docsIfCmtsDownChannelCounterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsDownChannelCounterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsDownChannelCounterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsDownChannelCounterTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsDownChannelCounterTable.EntityData.Children.Append("docsIfCmtsDownChannelCounterEntry", types.YChild{"DocsIfCmtsDownChannelCounterEntry", nil})
    for i := range docsIfCmtsDownChannelCounterTable.DocsIfCmtsDownChannelCounterEntry {
        docsIfCmtsDownChannelCounterTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsDownChannelCounterTable.DocsIfCmtsDownChannelCounterEntry[i]), types.YChild{"DocsIfCmtsDownChannelCounterEntry", docsIfCmtsDownChannelCounterTable.DocsIfCmtsDownChannelCounterEntry[i]})
    }
    docsIfCmtsDownChannelCounterTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsDownChannelCounterTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsDownChannelCounterTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsDownChannelCounterTable_DocsIfCmtsDownChannelCounterEntry
// An entry provides a list of traffic counters for a single
// downstream channel.
// An entry in this table exists for each ifEntry with an
// ifType of docsCableDownstream(128).
type DOCSIFMIB_DocsIfCmtsDownChannelCounterTable_DocsIfCmtsDownChannelCounterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The Cable Modem Termination System (CMTS) identification of the downstream
    // channel within this particular MAC interface. If the interface is down, the
    // object returns the most current value. If the downstream channel ID is
    // unknown, this object returns a value of 0. The type is interface{} with
    // range: 0..255.
    DocsIfCmtsDownChnlCtrId interface{}

    // At the CMTS, the total number of bytes in the Payload portion of MPEG
    // Packets (ie. not including MPEG header or pointer_field) transported by
    // this downstream channel since CMTS initialization. This is the 32 bit
    // version of docsIfCmtsDownChnlCtrExtTotalBytes, included to provide back
    // compatibility with SNMPv1 managers. The type is interface{} with range:
    // 0..4294967295.
    DocsIfCmtsDownChnlCtrTotalBytes interface{}

    // At the CMTS, the total number of DOCSIS data bytes transported by this
    // downstream channel since CMTS initialization. The number of data bytes is
    // defined as the total number of bytes transported in DOCSIS payloads minus
    // the number of stuff bytes transported in DOCSIS payloads. This is the 32
    // bit version of docsIfCmtsDownChnlCtrExtUsedBytes, included to provide back
    // compatibility with SNMPv1 managers. The type is interface{} with range:
    // 0..4294967295.
    DocsIfCmtsDownChnlCtrUsedBytes interface{}

    // At the CMTS, the total number of bytes in the Payload portion of MPEG
    // Packets (ie. not including MPEG header or pointer_field) transported by
    // this downstream channel since CMTS initialization. This is the 64 bit
    // version of docsIfCmtsDownChnlCtrTotalBytes, and will not be accessible to
    // SNMPv1 managers. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsDownChnlCtrExtTotalBytes interface{}

    // At the CMTS, the total number of DOCSIS data bytes transported by this
    // downstream channel since CMTS initialization. The number of data bytes is
    // defined as the total number of bytes transported in DOCSIS payloads minus
    // the number of stuff bytes transported in DOCSIS payloads. This is the 64
    // bit version of docsIfCmtsDownChnlCtrUsedBytes, and will not be accessible
    // to SNMPv1 managers. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsDownChnlCtrExtUsedBytes interface{}
}

func (docsIfCmtsDownChannelCounterEntry *DOCSIFMIB_DocsIfCmtsDownChannelCounterTable_DocsIfCmtsDownChannelCounterEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsDownChannelCounterEntry.EntityData.YFilter = docsIfCmtsDownChannelCounterEntry.YFilter
    docsIfCmtsDownChannelCounterEntry.EntityData.YangName = "docsIfCmtsDownChannelCounterEntry"
    docsIfCmtsDownChannelCounterEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsDownChannelCounterEntry.EntityData.ParentYangName = "docsIfCmtsDownChannelCounterTable"
    docsIfCmtsDownChannelCounterEntry.EntityData.SegmentPath = "docsIfCmtsDownChannelCounterEntry" + types.AddKeyToken(docsIfCmtsDownChannelCounterEntry.IfIndex, "ifIndex")
    docsIfCmtsDownChannelCounterEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsDownChannelCounterTable/" + docsIfCmtsDownChannelCounterEntry.EntityData.SegmentPath
    docsIfCmtsDownChannelCounterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsDownChannelCounterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsDownChannelCounterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsDownChannelCounterEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsDownChannelCounterEntry.IfIndex})
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsDownChnlCtrId", types.YLeaf{"DocsIfCmtsDownChnlCtrId", docsIfCmtsDownChannelCounterEntry.DocsIfCmtsDownChnlCtrId})
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsDownChnlCtrTotalBytes", types.YLeaf{"DocsIfCmtsDownChnlCtrTotalBytes", docsIfCmtsDownChannelCounterEntry.DocsIfCmtsDownChnlCtrTotalBytes})
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsDownChnlCtrUsedBytes", types.YLeaf{"DocsIfCmtsDownChnlCtrUsedBytes", docsIfCmtsDownChannelCounterEntry.DocsIfCmtsDownChnlCtrUsedBytes})
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsDownChnlCtrExtTotalBytes", types.YLeaf{"DocsIfCmtsDownChnlCtrExtTotalBytes", docsIfCmtsDownChannelCounterEntry.DocsIfCmtsDownChnlCtrExtTotalBytes})
    docsIfCmtsDownChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsDownChnlCtrExtUsedBytes", types.YLeaf{"DocsIfCmtsDownChnlCtrExtUsedBytes", docsIfCmtsDownChannelCounterEntry.DocsIfCmtsDownChnlCtrExtUsedBytes})

    docsIfCmtsDownChannelCounterEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmtsDownChannelCounterEntry.EntityData)
}

// DOCSIFMIB_DocsIfCmtsUpChannelCounterTable
// This table is implemented at the CMTS to provide upstream
// channel statistics appropriate for channel utilization
// calculations.
type DOCSIFMIB_DocsIfCmtsUpChannelCounterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of traffic statistics for a single upstream channel. For Docsis 2.0
    // CMTSs, an entry in this table exists for  each ifEntry with an ifType of
    // docsCableUpstreamChannel (205). For Docsis 1.x CMTSs, an entry in this
    // table exists for each ifEntry with an ifType of docsCableUpstreamInterface
    // (129). The type is slice of
    // DOCSIFMIB_DocsIfCmtsUpChannelCounterTable_DocsIfCmtsUpChannelCounterEntry.
    DocsIfCmtsUpChannelCounterEntry []*DOCSIFMIB_DocsIfCmtsUpChannelCounterTable_DocsIfCmtsUpChannelCounterEntry
}

func (docsIfCmtsUpChannelCounterTable *DOCSIFMIB_DocsIfCmtsUpChannelCounterTable) GetEntityData() *types.CommonEntityData {
    docsIfCmtsUpChannelCounterTable.EntityData.YFilter = docsIfCmtsUpChannelCounterTable.YFilter
    docsIfCmtsUpChannelCounterTable.EntityData.YangName = "docsIfCmtsUpChannelCounterTable"
    docsIfCmtsUpChannelCounterTable.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsUpChannelCounterTable.EntityData.ParentYangName = "DOCS-IF-MIB"
    docsIfCmtsUpChannelCounterTable.EntityData.SegmentPath = "docsIfCmtsUpChannelCounterTable"
    docsIfCmtsUpChannelCounterTable.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/" + docsIfCmtsUpChannelCounterTable.EntityData.SegmentPath
    docsIfCmtsUpChannelCounterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsUpChannelCounterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsUpChannelCounterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsUpChannelCounterTable.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsUpChannelCounterTable.EntityData.Children.Append("docsIfCmtsUpChannelCounterEntry", types.YChild{"DocsIfCmtsUpChannelCounterEntry", nil})
    for i := range docsIfCmtsUpChannelCounterTable.DocsIfCmtsUpChannelCounterEntry {
        docsIfCmtsUpChannelCounterTable.EntityData.Children.Append(types.GetSegmentPath(docsIfCmtsUpChannelCounterTable.DocsIfCmtsUpChannelCounterEntry[i]), types.YChild{"DocsIfCmtsUpChannelCounterEntry", docsIfCmtsUpChannelCounterTable.DocsIfCmtsUpChannelCounterEntry[i]})
    }
    docsIfCmtsUpChannelCounterTable.EntityData.Leafs = types.NewOrderedMap()

    docsIfCmtsUpChannelCounterTable.EntityData.YListKeys = []string {}

    return &(docsIfCmtsUpChannelCounterTable.EntityData)
}

// DOCSIFMIB_DocsIfCmtsUpChannelCounterTable_DocsIfCmtsUpChannelCounterEntry
// List of traffic statistics for a single upstream channel.
// For Docsis 2.0 CMTSs, an entry in this table exists for 
// each ifEntry with an ifType of docsCableUpstreamChannel (205).
// For Docsis 1.x CMTSs, an entry in this table exists for each
// ifEntry with an ifType of docsCableUpstreamInterface (129).
type DOCSIFMIB_DocsIfCmtsUpChannelCounterTable_DocsIfCmtsUpChannelCounterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The CMTS identification of the upstream channel. The type is interface{}
    // with range: 0..255.
    DocsIfCmtsUpChnlCtrId interface{}

    // Current count, from CMTS initialization, of all minislots defined for this
    // upstream logical channel. This count includes all IUCs and SIDs, even those
    // allocated to the NULL SID for a 2.0 logical channel which is inactive. This
    // is the 32 bit version of docsIfCmtsUpChnlCtrExtTotalMslots, and is included
    // for back compatibility with SNMPv1 managers. Support for this object is
    // mandatory. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrTotalMslots interface{}

    // Current count, from CMTS initialization, of unicast granted minislots on
    // the upstream logical channel, regardless of burst type. Unicast granted
    // minislots are those in which the CMTS assigned bandwidth to any unicast SID
    // on the logical channel. This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtUcastGrantedMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is mandatory.
    // The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrUcastGrantedMslots interface{}

    // Current count, from CMTS initialization, of contention minislots defined
    // for this upstream logical channel. This count includes all minislots
    // assigned to a broadcast or multicast SID on the logical channel.  This is
    // the 32 bit version of docsIfCmtsUpChnlCtrExtTotalCntnMslots, and is
    // included for back compatibility with SNMPv1 managers. Support for this
    // object is mandatory. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrTotalCntnMslots interface{}

    // Current count, from CMTS initialization, of contention minislots utilized
    // on the upstream logical channel. For contention regions, utilized minislots
    // are those in which the CMTS correctly received an upstream burst from any
    // CM on the upstream logical channel. This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtUsedCntnMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is mandatory.
    // The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrUsedCntnMslots interface{}

    // Current count, from CMTS initialization, of all minislots defined for this
    // upstream logical channel. This count includes all IUCs and SIDs, even those
    // allocated to the NULL SID for a 2.0 logical channel which is inactive. This
    // is the 64 bit version of docsIfCmtsUpChnlCtrTotalMslots, and will not be
    // accessible to SNMPv1 managers. Support for this object is mandatory. The
    // type is interface{} with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtTotalMslots interface{}

    // Current count, from CMTS initialization, of unicast granted minislots on
    // the upstream logical channel, regardless of burst type. Unicast granted
    // minislots are those in which the CMTS assigned bandwidth to any unicast SID
    // on the logical channel. This is the 64 bit version of
    // docsIfCmtsUpChnlCtrUcastGrantedMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is mandatory. The type is interface{}
    // with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtUcastGrantedMslots interface{}

    // Current count, from CMTS initialization, of contention minislots defined
    // for this upstream logical channel. This count includes all minislots
    // assigned to a broadcast or multicast SID on the logical channel.  This is
    // the 64 bit version of docsIfCmtsUpChnlCtrTotalCntnMslots, and will not be
    // accessible to SNMPv1 managers. Support for this object is mandatory. The
    // type is interface{} with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtTotalCntnMslots interface{}

    // Current count, from CMTS initialization, of contention minislots utilized
    // on the upstream logical channel. For contention regions, utilized minislots
    // are those in which the CMTS correctly received an upstream burst from any
    // CM on the upstream logical channel. This is the 64 bit version of
    // docsIfCmtsUpChnlCtrUsedCntnMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is mandatory. The type is interface{}
    // with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtUsedCntnMslots interface{}

    // Current count, from CMTS initialization, of contention minislots  subjected
    // to collisions on the upstream logical channel. For  contention regions,
    // these are the minislots applicable to bursts  that the CMTS detected, but
    // could not correctly receive.   This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtCollCntnMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, a value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrCollCntnMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // defined for this upstream logical channel. This count  includes all
    // minislots for IUC1 assigned to a broadcast or multicast  SID on the logical
    // channel.  This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtTotalCntnReqMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrTotalCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // utilized on this upstream logical channel. This count  includes all
    // contention minislots for IUC1 applicable to bursts that the CMTS correctly
    // received.              This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtUsedCntnReqMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrUsedCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // subjected to collisions on this upstream logical channel.   This includes
    // all contention minislots for IUC1 applicable to bursts that the CMTS
    // detected, but could not correctly receive.              This is the 32 bit
    // version of docsIfCmtsUpChnlCtrExtCollCntnReqMslots, and is included for
    // back compatibility with SNMPv1 managers. Support for this object is
    // optional. If the object is not supported, A value of zero is returned. The
    // type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrCollCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots defined for this upstream logical channel. This count  includes
    // all minislots for IUC2 assigned to a broadcast or multicast  SID on the
    // logical channel.  This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtTotalCntnReqDataMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrTotalCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots utilized on this upstream logical channel. This   includes all
    // contention minislots for IUC2 applicable to bursts that the CMTS correctly
    // received.              This is the 32 bit version of 
    // docsIfCmtsUpChnlCtrExtUsedCntnReqDataMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrUsedCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots subjected to collisions on this upstream logical  channel. This
    // includes all contention minislots for IUC2 applicable  to bursts that the
    // CMTS detected, but could not correctly receive.              This is the 32
    // bit version of  docsIfCmtsUpChnlCtrExtCollCntnReqDataMslots, and is
    // included for back compatibility with SNMPv1 managers. Support for this
    // object is optional. If the object is not supported, A value of zero is
    // returned. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrCollCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of contention initial maintenance
    // minislots defined for this upstream logical channel.  This includes all
    // minislots for IUC3 assigned to a broadcast or  multicast SID on the logical
    // channel.  This is the 32 bit version of
    // docsIfCmtsUpChnlCtrExtTotalCntnInitMaintMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrTotalCntnInitMaintMslots interface{}

    // Current count, from CMTS initialization, of contention initial maintenance
    // minislots utilized on this upstream logical channel.    This includes all
    // contention minislots for IUC3 applicable to bursts that the CMTS correctly
    // received.              This is the 32 bit version of 
    // docsIfCmtsUpChnlCtrExtUsedCntnInitMaintMslots, and is included for back
    // compatibility with SNMPv1 managers. Support for this object is optional. If
    // the object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrUsedCntnInitMaintMslots interface{}

    // Current count, from CMTS initialization, of contention initial maintenance
    // minislots subjected to collisions on this upstream  logical channel. This
    // includes all contention minislots for IUC3 applicable to bursts that the
    // CMTS detected, but could not correctly receive.              This is the 32
    // bit version of  docsIfCmtsUpChnlCtrExtCollCntnInitMaintMslots, and is
    // included for back compatibility with SNMPv1 managers. Support for this
    // object is optional. If the object is not supported, A value of zero is
    // returned. The type is interface{} with range: 0..4294967295.
    DocsIfCmtsUpChnlCtrCollCntnInitMaintMslots interface{}

    // Current count, from CMTS initialization, of collision contention  minislots
    // on the upstream logical channel. For contention regions, these are the
    // minislots applicable to bursts that the CMTS detected,   but could not
    // correctly receive. This is the 64 bit version of
    // docsIfCmtsUpChnlCtrCollCntnMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is optional. If the object is not
    // supported, a value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtCollCntnMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // defined for this upstream logical channel. This count  includes all
    // minislots for IUC1 assigned to a broadcast or multicast  SID on the logical
    // channel.  This is the 64 bit version of
    // docsIfCmtsUpChnlCtrTotalCntnReqMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtTotalCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // utilized on this upstream logical channel. This count  includes all
    // contention minislots for IUC1 applicable to bursts that the CMTS correctly
    // received.              This is the 64 bit version of
    // docsIfCmtsUpChnlCtrUsedCntnReqMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtUsedCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request minislots
    // subjected to collisions on this upstream logical channel. This includes all
    // contention minislots for IUC1 applicable to bursts that the CMTS detected,
    // but could not correctly receive.              This is the 64 bit version of
    // docsIfCmtsUpChnlCtrCollCntnReqMslots, and will not be accessible to SNMPv1
    // managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtCollCntnReqMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots defined for this upstream logical channel. This count  includes
    // all minislots for IUC2 assigned to a broadcast or multicast  SID on the
    // logical channel.  This is the 64 bit version of
    // docsIfCmtsUpChnlCtrTotalCntnReqDataMslots, and will not be accessible to
    // SNMPv1 managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtTotalCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots utilized on this upstream logical channel. This   includes all
    // contention minislots for IUC2 applicable to bursts that the CMTS correctly
    // received.              This is the 64 bit version of
    // docsIfCmtsUpChnlCtrUsedCntnReqDataMslots, and will not be accessible to
    // SNMPv1 managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtUsedCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of contention request data
    // minislots subjected to collisions on this upstream logical  channel. This
    // includes all contention minislots for IUC2 applicable  to bursts that the
    // CMTS detected, but could not correctly receive.              This is the 64
    // bit version of docsIfCmtsUpChnlCtrCollCntnReqDataMslots, and will not be
    // accessible to SNMPv1 managers. Support for this object is optional. If the
    // object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtCollCntnReqDataMslots interface{}

    // Current count, from CMTS initialization, of initial maintenance minislots
    // defined for this upstream logical channel. This count  includes all
    // minislots for IUC3 assigned to a broadcast or multicast  SID on the logical
    // channel.  This is the 64 bit version of 
    // docsIfCmtsUpChnlCtrTotalCntnInitMaintMslots, and will not be accessible to
    // SNMPv1 managers. Support for this object is optional. If the object is not
    // supported, A value of zero is returned. The type is interface{} with range:
    // 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtTotalCntnInitMaintMslots interface{}

    // Current count, from CMTS initialization, of initial maintenance minislots
    // utilized on this upstream logical channel. This   includes all contention
    // minislots for IUC3 applicable to bursts that the CMTS correctly received.  
    // This is the 64 bit version of docsIfCmtsUpChnlCtrUsedCntnInitMaintMslots,
    // and will not be accessible to SNMPv1 managers. Support for this object is
    // optional. If the object is not supported, A value of zero is returned. The
    // type is interface{} with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtUsedCntnInitMaintMslots interface{}

    // Current count, from CMTS initialization, of contention initial maintenance
    // minislots subjected to collisions on this upstream  logical channel. This
    // includes all contention minislots for IUC3 applicable to bursts that the
    // CMTS detected, but could not correctly receive.              This is the 64
    // bit version of docsIfCmtsUpChnlCtrCollCntnInitMaintMslots, and will not be
    // accessible to SNMPv1 managers. Support for this object is optional. If the
    // object is not supported, A value of zero is returned. The type is
    // interface{} with range: 0..18446744073709551615.
    DocsIfCmtsUpChnlCtrExtCollCntnInitMaintMslots interface{}
}

func (docsIfCmtsUpChannelCounterEntry *DOCSIFMIB_DocsIfCmtsUpChannelCounterTable_DocsIfCmtsUpChannelCounterEntry) GetEntityData() *types.CommonEntityData {
    docsIfCmtsUpChannelCounterEntry.EntityData.YFilter = docsIfCmtsUpChannelCounterEntry.YFilter
    docsIfCmtsUpChannelCounterEntry.EntityData.YangName = "docsIfCmtsUpChannelCounterEntry"
    docsIfCmtsUpChannelCounterEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIfCmtsUpChannelCounterEntry.EntityData.ParentYangName = "docsIfCmtsUpChannelCounterTable"
    docsIfCmtsUpChannelCounterEntry.EntityData.SegmentPath = "docsIfCmtsUpChannelCounterEntry" + types.AddKeyToken(docsIfCmtsUpChannelCounterEntry.IfIndex, "ifIndex")
    docsIfCmtsUpChannelCounterEntry.EntityData.AbsolutePath = "DOCS-IF-MIB:DOCS-IF-MIB/docsIfCmtsUpChannelCounterTable/" + docsIfCmtsUpChannelCounterEntry.EntityData.SegmentPath
    docsIfCmtsUpChannelCounterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIfCmtsUpChannelCounterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIfCmtsUpChannelCounterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIfCmtsUpChannelCounterEntry.EntityData.Children = types.NewOrderedMap()
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIfCmtsUpChannelCounterEntry.IfIndex})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrId", types.YLeaf{"DocsIfCmtsUpChnlCtrId", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrId})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrTotalMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrTotalMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrTotalMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrUcastGrantedMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrUcastGrantedMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrUcastGrantedMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrTotalCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrTotalCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrTotalCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrUsedCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrUsedCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrUsedCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtTotalMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtTotalMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtTotalMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtUcastGrantedMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtUcastGrantedMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtUcastGrantedMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtTotalCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtTotalCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtTotalCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtUsedCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtUsedCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtUsedCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrCollCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrCollCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrCollCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrTotalCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrTotalCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrTotalCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrUsedCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrUsedCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrUsedCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrCollCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrCollCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrCollCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrTotalCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrTotalCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrTotalCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrUsedCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrUsedCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrUsedCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrCollCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrCollCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrCollCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrTotalCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrTotalCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrTotalCntnInitMaintMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrUsedCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrUsedCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrUsedCntnInitMaintMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrCollCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrCollCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrCollCntnInitMaintMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtCollCntnMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtCollCntnMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtCollCntnMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtTotalCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtTotalCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtTotalCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtUsedCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtUsedCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtUsedCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtCollCntnReqMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtCollCntnReqMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtCollCntnReqMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtTotalCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtTotalCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtTotalCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtUsedCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtUsedCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtUsedCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtCollCntnReqDataMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtCollCntnReqDataMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtCollCntnReqDataMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtTotalCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtTotalCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtTotalCntnInitMaintMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtUsedCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtUsedCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtUsedCntnInitMaintMslots})
    docsIfCmtsUpChannelCounterEntry.EntityData.Leafs.Append("docsIfCmtsUpChnlCtrExtCollCntnInitMaintMslots", types.YLeaf{"DocsIfCmtsUpChnlCtrExtCollCntnInitMaintMslots", docsIfCmtsUpChannelCounterEntry.DocsIfCmtsUpChnlCtrExtCollCntnInitMaintMslots})

    docsIfCmtsUpChannelCounterEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIfCmtsUpChannelCounterEntry.EntityData)
}

