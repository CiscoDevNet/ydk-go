// This is the MIB Module for Cable Spectrum Management for
// DOCSIS-compliant Cable Modem Termination Systems (CMTS).
// 
// Spectrum management is a software/hardware feature provided
// in the CMTS so that the CMTS may sense both downstream and
// upstream plant impairments, report them to a management
// entity, and automatically mitigate them where possible.
// 
// The CMTS directly senses upstream transmission errors.It
// may also indirectly monitor the condition of the plant by
// keeping a record of modem state changes.  It is desireable
// to perform these functions without reducing throughput or
// latency and without creating additional packet overhead on
// the RF plant.
// 
// The purpose of cable Spectrum Management is to prevent long
// term service interruptions caused by upstream noise events
// in the cable plant.  It is also used for fault management
// and trouble shooting the cable network.  When modems are
// detected to go on-line and off-line by flap detectors, the
// cable operators can look at the flap list and spectrum
// tables to determine the possible causes.
package cisco_cable_spectrum_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cable_spectrum_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CABLE-SPECTRUM-MIB CISCO-CABLE-SPECTRUM-MIB}", reflect.TypeOf(CISCOCABLESPECTRUMMIB{}))
    ydk.RegisterEntity("CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB", reflect.TypeOf(CISCOCABLESPECTRUMMIB{}))
}

// CCSRequestOperation represents - 'abort', which is used to abort the test;
type CCSRequestOperation string

const (
    CCSRequestOperation_none CCSRequestOperation = "none"

    CCSRequestOperation_start CCSRequestOperation = "start"

    CCSRequestOperation_abort CCSRequestOperation = "abort"
)

// CCSRequestOperState represents - 'others', other errors;
type CCSRequestOperState string

const (
    CCSRequestOperState_idle CCSRequestOperState = "idle"

    CCSRequestOperState_pending CCSRequestOperState = "pending"

    CCSRequestOperState_running CCSRequestOperState = "running"

    CCSRequestOperState_noError CCSRequestOperState = "noError"

    CCSRequestOperState_aborted CCSRequestOperState = "aborted"

    CCSRequestOperState_notOnLine CCSRequestOperState = "notOnLine"

    CCSRequestOperState_invalidMac CCSRequestOperState = "invalidMac"

    CCSRequestOperState_timeOut CCSRequestOperState = "timeOut"

    CCSRequestOperState_fftBusy CCSRequestOperState = "fftBusy"

    CCSRequestOperState_fftFailed CCSRequestOperState = "fftFailed"

    CCSRequestOperState_others CCSRequestOperState = "others"
)

// CISCOCABLESPECTRUMMIB
type CISCOCABLESPECTRUMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CcsFlapObjects CISCOCABLESPECTRUMMIB_CcsFlapObjects

    // This table keeps the records of modem state changes. It can be used to
    // identify the problematic cable modems. An entry can be deleted from the
    // table but can not be added to the table.
    CcsFlapTable CISCOCABLESPECTRUMMIB_CcsFlapTable

    // This table keeps the records of modem state changes, so it can be used to
    // identify the problematic cable  modems. An entry per modem is added to the
    // table automatically  by the system when it detects any state changes to the
    // modem.  Therefore, it can be deleted but  can not be added by the
    // management system.
    CcsCmFlapTable CISCOCABLESPECTRUMMIB_CcsCmFlapTable

    // This table contains the spectrum data requests.  There are two types of
    // request: background noise and SNR. Refer to ccsSpectrumRequestIfIndex and
    // ccsSpectrumRequestMacAddr DESCRIPTIONS on how the type of request is
    // determined.
    CcsSpectrumRequestTable CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable

    // This table contains the receiving power or background noise measurement
    // based on the criteria that is set in the ccsSpectrumRequestEntry.
    CcsSpectrumDataTable CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable

    // A table of CNR requests.
    CcsSNRRequestTable CISCOCABLESPECTRUMMIB_CcsSNRRequestTable

    // This table contains the cable upstream interfaces assigned to a spectrum
    // group. A spectrum group contains one or more fixed frequencies or frequency
    // bands which can be assigned to cable upstream interfaces in the spectrum
    // group.
    CcsUpInSpecGroupTable CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable

    // This table contains the cable upstream interfaces in a fiber-node.Each
    // fiber-node uniquely represents an RF domain.Cable upstream interfaces in
    // the same fiber-node are physically combined together into the same RF
    // domain.
    CcsUpInFiberNodeTable CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable

    // This table contains the attributes of the cable upstream interfaces, ifType
    // of docsCableUpstream(129), to be used for improving performance and
    // proactive hopping.  Proactive hopping is achieved by setting the SNR 
    // polling period over the in-use band without CM signals.
    CcsUpSpecMgmtTable CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable

    // This table contains the frequency and band configuration of the spectrum
    // group.
    CcsSpecGroupFreqTable CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable
}

func (cISCOCABLESPECTRUMMIB *CISCOCABLESPECTRUMMIB) GetEntityData() *types.CommonEntityData {
    cISCOCABLESPECTRUMMIB.EntityData.YFilter = cISCOCABLESPECTRUMMIB.YFilter
    cISCOCABLESPECTRUMMIB.EntityData.YangName = "CISCO-CABLE-SPECTRUM-MIB"
    cISCOCABLESPECTRUMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCABLESPECTRUMMIB.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    cISCOCABLESPECTRUMMIB.EntityData.SegmentPath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB"
    cISCOCABLESPECTRUMMIB.EntityData.AbsolutePath = cISCOCABLESPECTRUMMIB.EntityData.SegmentPath
    cISCOCABLESPECTRUMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCABLESPECTRUMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCABLESPECTRUMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCABLESPECTRUMMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsFlapObjects", types.YChild{"CcsFlapObjects", &cISCOCABLESPECTRUMMIB.CcsFlapObjects})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsFlapTable", types.YChild{"CcsFlapTable", &cISCOCABLESPECTRUMMIB.CcsFlapTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsCmFlapTable", types.YChild{"CcsCmFlapTable", &cISCOCABLESPECTRUMMIB.CcsCmFlapTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsSpectrumRequestTable", types.YChild{"CcsSpectrumRequestTable", &cISCOCABLESPECTRUMMIB.CcsSpectrumRequestTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsSpectrumDataTable", types.YChild{"CcsSpectrumDataTable", &cISCOCABLESPECTRUMMIB.CcsSpectrumDataTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsSNRRequestTable", types.YChild{"CcsSNRRequestTable", &cISCOCABLESPECTRUMMIB.CcsSNRRequestTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsUpInSpecGroupTable", types.YChild{"CcsUpInSpecGroupTable", &cISCOCABLESPECTRUMMIB.CcsUpInSpecGroupTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsUpInFiberNodeTable", types.YChild{"CcsUpInFiberNodeTable", &cISCOCABLESPECTRUMMIB.CcsUpInFiberNodeTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsUpSpecMgmtTable", types.YChild{"CcsUpSpecMgmtTable", &cISCOCABLESPECTRUMMIB.CcsUpSpecMgmtTable})
    cISCOCABLESPECTRUMMIB.EntityData.Children.Append("ccsSpecGroupFreqTable", types.YChild{"CcsSpecGroupFreqTable", &cISCOCABLESPECTRUMMIB.CcsSpecGroupFreqTable})
    cISCOCABLESPECTRUMMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCABLESPECTRUMMIB.EntityData.YListKeys = []string {}

    return &(cISCOCABLESPECTRUMMIB.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsFlapObjects
type CISCOCABLESPECTRUMMIB_CcsFlapObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of flapped modem entries (ccsCmFlapEntry) per Cable
    // downstream interface that  can be reported in the ccsCmFlapTable.  If the
    // Cable downstream interface has more flapped modems than the
    // ccsFlapListMaxSize, some of the flapped modems  will not be shown in the
    // ccsCmFlapTable.  In this case, the users may want to increase the
    // ccsFlapMaxSize. The type is interface{} with range: 1..65536. Units are
    // modems.
    CcsFlapListMaxSize interface{}

    // The total number of flapped modem entries (ccsCmFlapEntry) that reported in
    // the ccsCmFlapTable.  The maximum value  will be ccsFlapListMaxSize * <no of
    // downstreams>. The type is interface{} with range: 0..65536. Units are
    // modems.
    CcsFlapListCurrentSize interface{}

    // The flap entry aging threshold.  Periodically, the aging process scans
    // through the flap list and removes the cable modems that have not flapped
    // for that many minutes. The type is interface{} with range: 1..86400. Units
    // are minutes.
    CcsFlapAging interface{}

    // The insertion-time is an empirically derived, worst-case number of seconds
    // which the cable modem requires to complete registration.  The time taken by
    // cable modems to complete their registration is measured by the cable
    // operators and this information helps to determine the insertion time.  If
    // the cable modem has not completed the registration stage within this
    // insertion-time setting, the cable modem will be inserted into the
    // flap-list. The type is interface{} with range: 60..86400. Units are
    // seconds.
    CcsFlapInsertionTime interface{}

    // The power adjust threshold.  When the power of the modem is adjusted beyond
    // this threshold, the modem will be inserted into the flap-list. The type is
    // interface{} with range: 1..10. Units are db.
    CcsFlapPowerAdjustThreshold interface{}

    // Per modem miss threshold which triggers polling flap detector. When the
    // number of times a cable modem does not acknowledge a  MAC-layer keepalive
    // message from a cable modem card exceeds the  miss threshold, the cable
    // modem is placed in the flap list. The type is interface{} with range:
    // 1..12.
    CcsFlapMissThreshold interface{}

    // Setting this object to true(1) causes ccsFlapInsertionFailNum,
    // ccsFlapHitNum, ccsFlapMissNum, ccsFlapCrcErrorNum, 
    // ccsFlapPowerAdjustmentNum and ccsFlapTotalNum objects of each entry in
    // ccsFlapTable to be started from zero. Reading this  object always returns
    // false(2). The type is bool.
    CcsFlapResetAll interface{}

    // Setting this object to true(1) removes all cable modems from flap-list and
    // all the entries in the ccsFlapTable are destroyed. If a modem keeps
    // flapping, the modem will be added again into the flap list and a new entry
    // in the ccsFlapTable will be created.  The newly created entry for that
    // modem will have new value of ccsFlapCreateTime and all the statistical
    // objects will be started from zero. Reading this object always returns
    // false(2). The type is bool.
    CcsFlapClearAll interface{}

    // The last time that all the entries in the ccsFlapTable are destroyed. There
    // are several ways to destroy all the entries in the ccsFlapTable. Setting
    // the object ccsFlapClearAll to true is one way, and the other way is through
    // Command Line Interface. This timestamp can be used to know when all the
    // entries in the ccsFlapTable are destroyed. The special value of all '00'Hs
    // indicates that the entries in the ccsFlapTable have never been destroyed.
    // The type is string.
    CcsFlapLastClearTime interface{}
}

func (ccsFlapObjects *CISCOCABLESPECTRUMMIB_CcsFlapObjects) GetEntityData() *types.CommonEntityData {
    ccsFlapObjects.EntityData.YFilter = ccsFlapObjects.YFilter
    ccsFlapObjects.EntityData.YangName = "ccsFlapObjects"
    ccsFlapObjects.EntityData.BundleName = "cisco_ios_xe"
    ccsFlapObjects.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsFlapObjects.EntityData.SegmentPath = "ccsFlapObjects"
    ccsFlapObjects.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsFlapObjects.EntityData.SegmentPath
    ccsFlapObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsFlapObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsFlapObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsFlapObjects.EntityData.Children = types.NewOrderedMap()
    ccsFlapObjects.EntityData.Leafs = types.NewOrderedMap()
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapListMaxSize", types.YLeaf{"CcsFlapListMaxSize", ccsFlapObjects.CcsFlapListMaxSize})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapListCurrentSize", types.YLeaf{"CcsFlapListCurrentSize", ccsFlapObjects.CcsFlapListCurrentSize})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapAging", types.YLeaf{"CcsFlapAging", ccsFlapObjects.CcsFlapAging})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapInsertionTime", types.YLeaf{"CcsFlapInsertionTime", ccsFlapObjects.CcsFlapInsertionTime})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapPowerAdjustThreshold", types.YLeaf{"CcsFlapPowerAdjustThreshold", ccsFlapObjects.CcsFlapPowerAdjustThreshold})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapMissThreshold", types.YLeaf{"CcsFlapMissThreshold", ccsFlapObjects.CcsFlapMissThreshold})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapResetAll", types.YLeaf{"CcsFlapResetAll", ccsFlapObjects.CcsFlapResetAll})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapClearAll", types.YLeaf{"CcsFlapClearAll", ccsFlapObjects.CcsFlapClearAll})
    ccsFlapObjects.EntityData.Leafs.Append("ccsFlapLastClearTime", types.YLeaf{"CcsFlapLastClearTime", ccsFlapObjects.CcsFlapLastClearTime})

    ccsFlapObjects.EntityData.YListKeys = []string {}

    return &(ccsFlapObjects.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsFlapTable
// This table keeps the records of modem state changes.
// It can be used to identify the problematic cable modems.
// An entry can be deleted from the table but can not be
// added to the table.
type CISCOCABLESPECTRUMMIB_CcsFlapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of attributes for an entry in the ccsFlapTable. An entry in this table
    // exists for each cable modem that triggered one of our flap detectors. The
    // type is slice of CISCOCABLESPECTRUMMIB_CcsFlapTable_CcsFlapEntry.
    CcsFlapEntry []*CISCOCABLESPECTRUMMIB_CcsFlapTable_CcsFlapEntry
}

func (ccsFlapTable *CISCOCABLESPECTRUMMIB_CcsFlapTable) GetEntityData() *types.CommonEntityData {
    ccsFlapTable.EntityData.YFilter = ccsFlapTable.YFilter
    ccsFlapTable.EntityData.YangName = "ccsFlapTable"
    ccsFlapTable.EntityData.BundleName = "cisco_ios_xe"
    ccsFlapTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsFlapTable.EntityData.SegmentPath = "ccsFlapTable"
    ccsFlapTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsFlapTable.EntityData.SegmentPath
    ccsFlapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsFlapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsFlapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsFlapTable.EntityData.Children = types.NewOrderedMap()
    ccsFlapTable.EntityData.Children.Append("ccsFlapEntry", types.YChild{"CcsFlapEntry", nil})
    for i := range ccsFlapTable.CcsFlapEntry {
        ccsFlapTable.EntityData.Children.Append(types.GetSegmentPath(ccsFlapTable.CcsFlapEntry[i]), types.YChild{"CcsFlapEntry", ccsFlapTable.CcsFlapEntry[i]})
    }
    ccsFlapTable.EntityData.Leafs = types.NewOrderedMap()

    ccsFlapTable.EntityData.YListKeys = []string {}

    return &(ccsFlapTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsFlapTable_CcsFlapEntry
// List of attributes for an entry in the ccsFlapTable.
// An entry in this table exists for each cable modem that
// triggered one of our flap detectors.
type CISCOCABLESPECTRUMMIB_CcsFlapTable_CcsFlapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MAC address of the Cable Modem's Cable interface
    // which identifies a flap-list entry for a flapping  Cable Modem. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CcsFlapMacAddr interface{}

    // The ifIndex of the Cable upstream interface whose ifType is
    // docsCableUpstream(129).  The CMTS detects a flapping Cable Modem from its
    // Cable upstream interface. The type is interface{} with range:
    // 1..2147483647.
    CcsFlapUpstreamIfIndex interface{}

    // The ifIndex of the Cable downstream interface whose ifType is
    // docsCableDownstream(128). The type is interface{} with range:
    // 1..2147483647.
    CcsFlapDownstreamIfIndex interface{}

    // The number of times a Cable Modem registered more frequently than expected.
    // Excessive registration is defined as the presence of a time span between
    // two successive registration cycles which is less than a threshold span
    // (ccsFlapInsertionTime).  A Cable Modem may fail the ranging or registration
    // process due to not being able to get an IP address. When the Cable Modem
    // can not finish registration within the insertion time, it retries the
    // process and sends the Initial Maintenance packet again. CMTS will receive
    // the Initial Maintenance packet from the Cable Modem sooner than expected
    // and the Cable Modem is considered a flapping modem.  This count may
    // indicate:     Intermittent downstream sync loss, or     DHCP or modem
    // registration problems.  The Flap Count (ccsFlapTotal) will be incremented
    // when this counter is incremented.  Discontinuites in the value of this
    // counter can occur if this entry is removed from the table and then
    // re-added, and are indicated by a change in the value of ccsFlapCreateTime.
    // The type is interface{} with range: 0..4294967295.
    CcsFlapInsertionFails interface{}

    // The number of times the CMTS receives the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance transmit opportunity at a
    // typical rate of once every 10 seconds and waits for a Ranging request from
    // the Cable Modem.If the CMTS receives a Ranging request then the Hit count
    // will be increased by 1  If the FlapTotal count is high,both Hits and Misses
    // counts are high, and other counters are relatively low then the flapping is
    // probably caused by the modem going up and down. The Hits and Misses counts
    // are keep-alive polling statistics. The Hits count should be much greater
    // than the Misses count  Discontinuites in the value of this counter can
    // occur if this entry is removed from the table and then re-added, and are
    // indicated by a change in the value of ccsFlapCreateTime. The type is
    // interface{} with range: 0..4294967295.
    CcsFlapHits interface{}

    // The number of times the CMTS misses the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance packet every 10 seconds and
    // waits for a Ranging request from the Cable Modem. If the CMTS misses a
    // Ranging request within 25 msec then the Misses count will be incremented. 
    // If ccsFlapTotal is high, Hits and Misses are high but
    // ccsFlapPowerAdjustments and ccsFlapInsertionFails are low then the flapping
    // is probably caused by the modem going up and down.  Miss counts can
    // indicate:     Intermittent upstream,     Laser clipping, or     Noise
    // bursts.  Laser clipping can happen if the signal power is too high when the
    // upstream electrical signal is converted to an optical signal.  When it
    // happens the more input produces  less output, until finally there is no
    // more increase in  output.  This phenomena is called laser clipping. 
    // Discontinuites in the value of this counter can occur if this entry is
    // removed from the table and then re-added, and are indicated by a change in
    // the value of ccsFlapCreateTime. The type is interface{} with range:
    // 0..4294967295.
    CcsFlapMisses interface{}

    // The number of times the CMTS upstream receiver flagged a packet with a CRC
    // error.  If ccsFlapCrcErrors is high, it indicates the cable upstream may
    // have high noise level.  The modem may not be flapping yet but it may be a
    // potential problem.  This count can indicate:     Intermittent upstream,    
    // Laser clipping, or     Noise bursts.  Laser clipping can happen if the
    // signal power is too high when the upstream electrical signal is converted
    // to an optical signal.  When it happens the more input produces  less
    // output, until finally there is no more increase in  output.  This phenomena
    // is called laser clipping. Discontinuites in the value of this counter can
    // occur if this entry is removed from the table and then re-added, and are
    // indicated by a change in the value of ccsFlapCreateTime. The type is
    // interface{} with range: 0..4294967295.
    CcsFlapCrcErrors interface{}

    // The number of times the Cable Modem upstream transmit power is adjusted
    // during station maintenance.  When the adjustment is greater than the power
    // adjustment threshold the counter will be incremented. The power adjustment
    // threshold is chosen in an implementation-dependent manner.  The Flap Count
    // (ccsFlapTotal) will be incremented when this counter is incremented.  If
    // ccsFlapTotal is high, ccsFlapPowerAdjustments is high but the Hits and
    // Misses are low and ccsFlapInsertionFails are low then the flapping is
    // probably caused by an improper transmit power level setting at the modem
    // end.  This count can indicate:     Amplifier degradation,     Poor
    // connections, or     Wind, moisture, or temperature sensitivity. 
    // Discontinuites in the value of this counter can occur if this entry is
    // removed from the table and then re-added, and are indicated by a change in
    // the value of ccsFlapCreateTime. The type is interface{} with range:
    // 0..4294967295.
    CcsFlapPowerAdjustments interface{}

    // Whenever the Cable Modem passes flap detection, then the flap counter is
    // increased.  There are 3 flap detectors defined: (1) When
    // ccsFlapInsertionFails is increased the Flap count     will be increased.
    // (2) When the CMTS receives a Miss followed by a Hit     then the Flap count
    // will be increased. (3) When ccsFlapPowerAdjustments is increased the Flap  
    // count will be increased.  Discontinuites in the value of this counter can
    // occur if this entry is removed from the table and then re-added, and are
    // indicated by a change in the value of ccsFlapCreateTime. The type is
    // interface{} with range: 0..4294967295.
    CcsFlapTotal interface{}

    // The flap time is set whenever the Cable Modem triggers a flap detector. The
    // type is string.
    CcsFlapLastFlapTime interface{}

    // The time that this entry was added to the table. If an entry is removed and
    // then later re-added, there may be a discontinuity in the counters
    // associated with this entry. This timestamp can be used to detect those 
    // discontinuities. The type is string.
    CcsFlapCreateTime interface{}

    // Controls and reflects the status of rows in this table.  When a cable modem
    // triggers a flap detector, if an entry does not already exist for this cable
    // modem, and ccsFlapListCurrentSize is less than ccsFlapListMaxSize, then an
    // entry will be created in this table. The new  instance of this object will
    // be set to active(1).  All  flapping modems have the status of active(1). 
    // Active entries are removed from the table after they have not triggered any
    // additional flap detectors for the period of time defined in ccsFlapAging.
    // Alternatively, setting this instance to destroy(6) will remove the entry
    // immediately.  createAndGo(4) and createAndWait(5) are not supported. The
    // type is RowStatus.
    CcsFlapRowStatus interface{}

    // The number of times a Cable Modem registered more frequently than expected.
    // Excessive registration is defined as the presence of a time span between
    // two successive registration cycles which is less than a threshold span
    // (ccsFlapInsertionTime).  A Cable Modem may fail the ranging or registration
    // process due to not being able to get an IP address. When the Cable Modem
    // can not finish registration within the insertion time, it retries the
    // process and sends the Initial Maintenance packet again. CMTS will receive
    // the Initial Maintenance packet from the Cable Modem sooner than expected
    // and the Cable Modem is considered a flapping modem.  This object may
    // indicate:     Intermittent downstream sync loss, or     DHCP or modem
    // registration problems.  The Flap number (ccsFlapTotalNum) will be
    // incremented when this object is incremented.  This object is going to
    // replace the object ccsFlapInsertionFails and the value of this object can
    // be reset to zero if this entry is removed from the table and then re-added,
    // or if a user resets all the statistical objects for this entry. The value
    // of the object ccsFlapLastResetTime indicates the last reset time. The type
    // is interface{} with range: 0..4294967295.
    CcsFlapInsertionFailNum interface{}

    // The number of times the CMTS receives the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance transmit opportunity at a
    // typical rate of once every 10 seconds and waits for a Ranging request from
    // the Cable Modem. If the CMTS receives a Ranging request then the Hit number
    // will be increased by 1  If the FlapTotal object is high, both Hit and Miss
    // objects are high, and other statistical objects are relatively low then the
    // flapping is probably caused by the modem going up and down. The Hit and
    // Miss objects keep-alive polling statistics. The Hit object should be much
    // greater than the Misses count.  This object is going to replace the object
    // ccsFlapHits and the value of this object can be reset to zero if this entry
    // is removed from the table and then re-added, or if an user resets all the
    // statistical objects for this entry. The value of the object
    // ccsFlapLastResetTime indicates the last reset time. The type is interface{}
    // with range: 0..4294967295.
    CcsFlapHitNum interface{}

    // The number of times the CMTS misses the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance packet every 10 seconds and
    // waits for a Ranging request from the Cable Modem. If the CMTS misses a
    // Ranging request within 25 msec then the Miss Object will be incremented. 
    // If ccsFlapTotalNum is high, Hit and Miss are high but
    // ccsFlapPowerAdjustmentNum and ccsFlapInsertionFailNum are low then the
    // flapping is probably caused by the modem going up and down.  Miss object
    // can indicate:     Intermittent upstream,     Laser clipping, or     Noise
    // bursts.  Laser clipping can happen if the signal power is too high when the
    // upstream electrical signal is converted to an optical signal.  When it
    // happens the more input produces less output, until finally there is no more
    // increase in output. This phenomena is called laser clipping.  This object
    // is going to replace the object ccsFlapMisses and the value of this object
    // can be reset to zero if this entry is removed from the table and then
    // re-added, or if an user resets all the statistical objects for this entry.
    // The value of the object ccsFlapLastResetTime indicates the last reset time.
    // The type is interface{} with range: 0..4294967295.
    CcsFlapMissNum interface{}

    // The number of times the CMTS upstream receiver flagged a packet with a CRC
    // error.  If ccsFlapCrcErrorNum is high, it indicates the cable upstream may
    // have high noise level. The modem may not be flapping yet but it may be a
    // potential problem.  This object can indicate:     Intermittent upstream,   
    // Laser clipping, or     Noise bursts.  Laser clipping can happen if the
    // signal power is too high when the upstream electrical signal is converted
    // to an optical signal.  When it happens the more input produces  less
    // output, until finally there is no more increase in  output.  This phenomena
    // is called laser clipping.  This object is going to replace the object
    // ccsFlapCrcErrors and the value of this object can be reset to zero if this
    // entry is removed from the table and then re-added, or if a user resets all
    // the statistical objects for this entry. The value of the object
    // ccsFlapLastResetTime indicates the last reset time. The type is interface{}
    // with range: 0..4294967295.
    CcsFlapCrcErrorNum interface{}

    // The number of times the Cable Modem upstream transmit power is adjusted
    // during station maintenance.  When the adjustment  is greater than the power
    // adjustment threshold the number  will be incremented. The power adjustment
    // threshold is chosen  in an implementation-dependent manner  The Flap number
    // (ccsFlapTotalNum) will be incremented when this object is incremented.  If
    // ccsFlapTotalNum is high, ccsFlapPowerAdjustmentNum is high but the Hit and
    // Miss are low and ccsFlapInsertionFailNum are low then the flapping is
    // probably caused by an improper transmit power level setting at the modem
    // end.  This object can indicate:     Amplifier degradation,     Poor
    // connections, or     Wind, moisture, or temperature sensitivity.  This
    // object is going to replace the object ccsFlapPowerAdjustments and the value
    // of this object can be reset to zero if this entry is removed from the table
    // and then re-added, or if a user resets all the statistical objects for this
    // entry. The value of the object ccsFlapLastResetTime indicates the last
    // reset time. The type is interface{} with range: 0..4294967295.
    CcsFlapPowerAdjustmentNum interface{}

    // Whenever the Cable Modem passes flap detection, then the flap number is
    // increased.  There are 3 flap detectors defined: (1) When
    // ccsFlapInsertionFailNum is increased the Flap number     will be increased.
    // (2) When the CMTS receives a Miss followed by a Hit     then the Flap
    // number will be increased. (3) When ccsFlapPowerAdjustmentNum is increased
    // the Flap     number will be increased.  This object is going to replace the
    // object ccsFlapTotal and the value of this object can be reset to zero if
    // this entry is removed from the table and then re-added, or if an user
    // resets all the statistical objects for this entry. The value of the object
    // ccsFlapLastResetTime indicates the last reset time. The type is interface{}
    // with range: 0..4294967295.
    CcsFlapTotalNum interface{}

    // Setting this object to true(1) will set the following objects of this entry
    // to 0: ccsFlapInsertionFailsNum, ccsFlapHitsNum, ccsFlapMissesNum,
    // ccsFlapCrcErrorsNum, ccsFlapPowerAdjustmentsNum and ccsFlapTotalNum.
    // Setting this object to true does not destroy the entry, so the
    // ccsFlapCreateTime will be unchanged. Reading this object always returns
    // false(2). The type is bool.
    CcsFlapResetNow interface{}

    // The last time that all the statistical objects of this entry are started
    // from zero. There are several ways to restart the the statistical objects
    // from zero. Setting the object ccsFlapResetNow or ccsFlapResetAll to true
    // via SNMP is one way and and the other way is via command Line Interface.
    // This timestamp can be used to know the last time the statistical objects
    // are started from zero. The special value of all '00'Hs indicates that these
    // statistical objects of this entry in the ccsFlapTable have never been
    // reset. The type is string.
    CcsFlapLastResetTime interface{}
}

func (ccsFlapEntry *CISCOCABLESPECTRUMMIB_CcsFlapTable_CcsFlapEntry) GetEntityData() *types.CommonEntityData {
    ccsFlapEntry.EntityData.YFilter = ccsFlapEntry.YFilter
    ccsFlapEntry.EntityData.YangName = "ccsFlapEntry"
    ccsFlapEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsFlapEntry.EntityData.ParentYangName = "ccsFlapTable"
    ccsFlapEntry.EntityData.SegmentPath = "ccsFlapEntry" + types.AddKeyToken(ccsFlapEntry.CcsFlapMacAddr, "ccsFlapMacAddr")
    ccsFlapEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsFlapTable/" + ccsFlapEntry.EntityData.SegmentPath
    ccsFlapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsFlapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsFlapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsFlapEntry.EntityData.Children = types.NewOrderedMap()
    ccsFlapEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapMacAddr", types.YLeaf{"CcsFlapMacAddr", ccsFlapEntry.CcsFlapMacAddr})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapUpstreamIfIndex", types.YLeaf{"CcsFlapUpstreamIfIndex", ccsFlapEntry.CcsFlapUpstreamIfIndex})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapDownstreamIfIndex", types.YLeaf{"CcsFlapDownstreamIfIndex", ccsFlapEntry.CcsFlapDownstreamIfIndex})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapInsertionFails", types.YLeaf{"CcsFlapInsertionFails", ccsFlapEntry.CcsFlapInsertionFails})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapHits", types.YLeaf{"CcsFlapHits", ccsFlapEntry.CcsFlapHits})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapMisses", types.YLeaf{"CcsFlapMisses", ccsFlapEntry.CcsFlapMisses})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapCrcErrors", types.YLeaf{"CcsFlapCrcErrors", ccsFlapEntry.CcsFlapCrcErrors})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapPowerAdjustments", types.YLeaf{"CcsFlapPowerAdjustments", ccsFlapEntry.CcsFlapPowerAdjustments})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapTotal", types.YLeaf{"CcsFlapTotal", ccsFlapEntry.CcsFlapTotal})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapLastFlapTime", types.YLeaf{"CcsFlapLastFlapTime", ccsFlapEntry.CcsFlapLastFlapTime})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapCreateTime", types.YLeaf{"CcsFlapCreateTime", ccsFlapEntry.CcsFlapCreateTime})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapRowStatus", types.YLeaf{"CcsFlapRowStatus", ccsFlapEntry.CcsFlapRowStatus})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapInsertionFailNum", types.YLeaf{"CcsFlapInsertionFailNum", ccsFlapEntry.CcsFlapInsertionFailNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapHitNum", types.YLeaf{"CcsFlapHitNum", ccsFlapEntry.CcsFlapHitNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapMissNum", types.YLeaf{"CcsFlapMissNum", ccsFlapEntry.CcsFlapMissNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapCrcErrorNum", types.YLeaf{"CcsFlapCrcErrorNum", ccsFlapEntry.CcsFlapCrcErrorNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapPowerAdjustmentNum", types.YLeaf{"CcsFlapPowerAdjustmentNum", ccsFlapEntry.CcsFlapPowerAdjustmentNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapTotalNum", types.YLeaf{"CcsFlapTotalNum", ccsFlapEntry.CcsFlapTotalNum})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapResetNow", types.YLeaf{"CcsFlapResetNow", ccsFlapEntry.CcsFlapResetNow})
    ccsFlapEntry.EntityData.Leafs.Append("ccsFlapLastResetTime", types.YLeaf{"CcsFlapLastResetTime", ccsFlapEntry.CcsFlapLastResetTime})

    ccsFlapEntry.EntityData.YListKeys = []string {"CcsFlapMacAddr"}

    return &(ccsFlapEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsCmFlapTable
// This table keeps the records of modem state changes,
// so it can be used to identify the problematic cable 
// modems.
// An entry per modem is added to the table automatically 
// by the system when it detects any state changes to
// the modem.  Therefore, it can be deleted but 
// can not be added by the management system.
type CISCOCABLESPECTRUMMIB_CcsCmFlapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of attributes for an entry in the ccsCmFlapTable. An entry in this
    // table exists for each cable modem that triggered one of our flap detectors.
    // The type is slice of CISCOCABLESPECTRUMMIB_CcsCmFlapTable_CcsCmFlapEntry.
    CcsCmFlapEntry []*CISCOCABLESPECTRUMMIB_CcsCmFlapTable_CcsCmFlapEntry
}

func (ccsCmFlapTable *CISCOCABLESPECTRUMMIB_CcsCmFlapTable) GetEntityData() *types.CommonEntityData {
    ccsCmFlapTable.EntityData.YFilter = ccsCmFlapTable.YFilter
    ccsCmFlapTable.EntityData.YangName = "ccsCmFlapTable"
    ccsCmFlapTable.EntityData.BundleName = "cisco_ios_xe"
    ccsCmFlapTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsCmFlapTable.EntityData.SegmentPath = "ccsCmFlapTable"
    ccsCmFlapTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsCmFlapTable.EntityData.SegmentPath
    ccsCmFlapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsCmFlapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsCmFlapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsCmFlapTable.EntityData.Children = types.NewOrderedMap()
    ccsCmFlapTable.EntityData.Children.Append("ccsCmFlapEntry", types.YChild{"CcsCmFlapEntry", nil})
    for i := range ccsCmFlapTable.CcsCmFlapEntry {
        ccsCmFlapTable.EntityData.Children.Append(types.GetSegmentPath(ccsCmFlapTable.CcsCmFlapEntry[i]), types.YChild{"CcsCmFlapEntry", ccsCmFlapTable.CcsCmFlapEntry[i]})
    }
    ccsCmFlapTable.EntityData.Leafs = types.NewOrderedMap()

    ccsCmFlapTable.EntityData.YListKeys = []string {}

    return &(ccsCmFlapTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsCmFlapTable_CcsCmFlapEntry
// List of attributes for an entry in the ccsCmFlapTable.
// An entry in this table exists for each cable modem that
// triggered one of our flap detectors.
type CISCOCABLESPECTRUMMIB_CcsCmFlapTable_CcsCmFlapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ifIndex of the Cable downstream interface
    // whose ifType is docsCableDownstream(128). The type is interface{} with
    // range: 1..2147483647.
    CcsCmFlapDownstreamIfIndex interface{}

    // This attribute is a key. The ifIndex of the Cable upstream interface whose
    // ifType is docsCableUpstream(129).  The CMTS detects a flapping Cable Modem
    // from its Cable upstream interface. The type is interface{} with range:
    // 1..2147483647.
    CcsCmFlapUpstreamIfIndex interface{}

    // This attribute is a key. MAC address of the Cable Modem's Cable interface
    // which identifies a flapping Cable Modem. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CcsCmFlapMacAddr interface{}

    // The flap time is set whenever the Cable Modem triggers a flap detector. The
    // type is string.
    CcsCmFlapLastFlapTime interface{}

    // The time that this entry was added to the table. If an entry is removed and
    // then later re-added, there may be a discontinuity in the counters
    // associated with this entry. This timestamp can be used to detect those 
    // discontinuities. The type is string.
    CcsCmFlapCreateTime interface{}

    // The number of times a Cable Modem registered more frequently than expected.
    // Excessive registration is defined as the presence of a time span between
    // two successive registration cycles which is less than a threshold span
    // (ccsFlapInsertionTime).  A Cable Modem may fail the ranging or registration
    // process due to not being able to get an IP address. When the Cable Modem
    // can not finish registration within the insertion time, it retries the
    // process and sends the Initial Maintenance packet again. CMTS will receive
    // the Initial Maintenance packet from the Cable Modem sooner than expected
    // and the Cable Modem is considered a flapping modem.  This object may
    // indicate:     Intermittent downstream sync loss, or     DHCP or modem
    // registration problems.  The Flap number (ccsCmFlapTotalNum) will be
    // incremented when this object is incremented.  The value of this object can
    // be reset to zero if this entry  is removed from the table and then
    // re-added, or if the  ccsCmFlapResetNow object was set to true(1). The value
    // of the  object ccsCmFlapLastResetTime indicates the last reset time. The
    // type is interface{} with range: 0..4294967295.
    CcsCmFlapInsertionFailNum interface{}

    // The number of times the CMTS receives the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance transmit opportunity at a
    // typical rate of once every 10 seconds and waits for a Ranging request from
    // the Cable Modem. If the CMTS receives a Ranging request then the Hit number
    // will be increased by 1  If the FlapTotal object is high, both Hit and Miss
    // objects are high, and other statistical objects are relatively low then the
    // flapping is probably caused by the modem going up and down. The Hit and
    // Miss objects keep-alive polling statistics. The Hit object should be much
    // greater than the Misses count.  The value of this object can be reset to
    // zero if this entry  is removed from the table and then re-added, or if the 
    // ccsCmFlapResetNow object was set to true(1). The value of the  object
    // ccsCmFlapLastResetTime indicates the last reset time. The type is
    // interface{} with range: 0..4294967295.
    CcsCmFlapHitNum interface{}

    // The number of times the CMTS misses the Ranging request from the Cable
    // Modem.  The CMTS issues a Station Maintenance packet every 10 seconds and
    // waits for a Ranging request from the Cable Modem. If the CMTS misses a
    // Ranging request within 25 msec then the Miss Object will be incremented. 
    // If ccsCmFlapTotalNum is high, Hit and Miss are high but
    // ccsCmFlapPowerAdjustmentNum and ccsCmFlapInsertionFailNum  are low then the
    // flapping is probably caused by the modem  going up and down.  Miss object
    // can indicate:     Intermittent upstream,     Laser clipping, or     Noise
    // bursts.  Laser clipping can happen if the signal power is too high when the
    // upstream electrical signal is converted to an optical signal.  When it
    // happens the more input produces less output, until finally there is no more
    // increase in output. This phenomena is called laser clipping.  The value of
    // this object can be reset to zero if this entry  is removed from the table
    // and then re-added, or if the  ccsCmFlapResetNow object was set to true(1).
    // The value of the  object ccsCmFlapLastResetTime indicates the last reset
    // time. The type is interface{} with range: 0..4294967295.
    CcsCmFlapMissNum interface{}

    // The number of times the CMTS upstream receiver flagged a packet with a CRC
    // error.  If ccsCmFlapCrcErrorNum is high, it indicates the cable upstream
    // may have high noise level. The modem may not be flapping yet but it may be
    // a potential problem.  This object can indicate:     Intermittent upstream, 
    // Laser clipping, or     Noise bursts.  Laser clipping can happen if the
    // signal power is too high when the upstream electrical signal is converted
    // to an optical signal.  When it happens the more input produces less output,
    // until finally there is no more increase in output. This phenomena is called
    // laser clipping.  The value of this object can be reset to zero if this
    // entry  is removed from the table and then re-added, or if the 
    // ccsCmFlapResetNow object was set to true(1). The value of the  object
    // ccsCmFlapLastResetTime indicates the last reset time. The type is
    // interface{} with range: 0..4294967295.
    CcsCmFlapCrcErrorNum interface{}

    // The number of times the Cable Modem upstream transmit power is adjusted
    // during station maintenance. When the adjustment is greater than the power
    // adjustment threshold the number will be incremented. The power adjustment
    // threshold is chosen in an implementation-dependent manner  The Flap number
    // (ccsCmFlapTotalNum) will be incremented when this object is incremented. 
    // If ccsCmFlapTotalNum is high, ccsCmFlapPowerAdjustmentNum is  high but the
    // Hit and Miss are low and  ccsCmFlapInsertionFailNum are low then the
    // flapping is  probably caused by an improper transmit power level  setting
    // at the modem end.  This object can indicate:     Amplifier degradation,    
    // Poor connections, or     Wind, moisture, or temperature sensitivity.  The
    // value of this object can be reset to zero if this entry  is removed from
    // the table and then re-added, or if the  ccsCmFlapResetNow object was set to
    // true(1). The value of the  object ccsCmFlapLastResetTime indicates the last
    // reset time. The type is interface{} with range: 0..4294967295.
    CcsCmFlapPowerAdjustmentNum interface{}

    // Whenever the Cable Modem passes flap detection, then the flap number is
    // increased.  There are 3 flap detectors defined: (1) When
    // ccsCmFlapInsertionFailNum is increased the Flap number     will be
    // increased. (2) When the CMTS receives a Miss followed by a Hit     then the
    // Flap number will be increased. (3) When ccsCmFlapPowerAdjustmentNum is
    // increased the Flap     number will be increased.  The value of this object
    // can be reset to zero if this entry  is removed from the table and then
    // re-added, or if the  ccsCmFlapResetNow object was set to true(1). The value
    // of the  object ccsCmFlapLastResetTime indicates the last reset time. The
    // type is interface{} with range: 0..4294967295.
    CcsCmFlapTotalNum interface{}

    // Setting this object to true(1) will set the value of certain objects in
    // this table to 0 and does not destroy the entry, so the ccsCmFlapCreateTime
    // will be  unchanged. Reading this object always returns false(2). The type
    // is bool.
    CcsCmFlapResetNow interface{}

    // The last time that all the statistical objects of this entry are started
    // from zero. There are several ways to restart the the statistical objects
    // from zero. Setting the object ccsCmFlapResetNow or ccsCmFlapResetAll to
    // true via SNMP is one way and and the other way is via command Line
    // Interface. This timestamp can be used to know the last time the statistical
    // objects are started from zero. The special value of all '00'Hs indicates
    // that these statistical objects of this entry in the ccsCmFlapTable have
    // never been reset. The type is string.
    CcsCmFlapLastResetTime interface{}

    // Controls and reflects the status of rows in this table.  When a cable modem
    // triggers a flap detector, if an entry does not already exist for this cable
    // modem,  an entry will be created in this table.  The new instance of this
    // object will be set to active(1).  All flapping modems have the status of
    // active(1).  Active entries are removed from the table after they have not
    // triggered any additional flap detectors for the period of time defined in
    // ccsFlapAging. Alternatively, setting this instance to destroy(6) will
    // remove the entry immediately.  createAndGo(4) and createAndWait(5) are not
    // supported. The type is RowStatus.
    CcsCmFlapRowStatus interface{}
}

func (ccsCmFlapEntry *CISCOCABLESPECTRUMMIB_CcsCmFlapTable_CcsCmFlapEntry) GetEntityData() *types.CommonEntityData {
    ccsCmFlapEntry.EntityData.YFilter = ccsCmFlapEntry.YFilter
    ccsCmFlapEntry.EntityData.YangName = "ccsCmFlapEntry"
    ccsCmFlapEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsCmFlapEntry.EntityData.ParentYangName = "ccsCmFlapTable"
    ccsCmFlapEntry.EntityData.SegmentPath = "ccsCmFlapEntry" + types.AddKeyToken(ccsCmFlapEntry.CcsCmFlapDownstreamIfIndex, "ccsCmFlapDownstreamIfIndex") + types.AddKeyToken(ccsCmFlapEntry.CcsCmFlapUpstreamIfIndex, "ccsCmFlapUpstreamIfIndex") + types.AddKeyToken(ccsCmFlapEntry.CcsCmFlapMacAddr, "ccsCmFlapMacAddr")
    ccsCmFlapEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsCmFlapTable/" + ccsCmFlapEntry.EntityData.SegmentPath
    ccsCmFlapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsCmFlapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsCmFlapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsCmFlapEntry.EntityData.Children = types.NewOrderedMap()
    ccsCmFlapEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapDownstreamIfIndex", types.YLeaf{"CcsCmFlapDownstreamIfIndex", ccsCmFlapEntry.CcsCmFlapDownstreamIfIndex})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapUpstreamIfIndex", types.YLeaf{"CcsCmFlapUpstreamIfIndex", ccsCmFlapEntry.CcsCmFlapUpstreamIfIndex})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapMacAddr", types.YLeaf{"CcsCmFlapMacAddr", ccsCmFlapEntry.CcsCmFlapMacAddr})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapLastFlapTime", types.YLeaf{"CcsCmFlapLastFlapTime", ccsCmFlapEntry.CcsCmFlapLastFlapTime})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapCreateTime", types.YLeaf{"CcsCmFlapCreateTime", ccsCmFlapEntry.CcsCmFlapCreateTime})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapInsertionFailNum", types.YLeaf{"CcsCmFlapInsertionFailNum", ccsCmFlapEntry.CcsCmFlapInsertionFailNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapHitNum", types.YLeaf{"CcsCmFlapHitNum", ccsCmFlapEntry.CcsCmFlapHitNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapMissNum", types.YLeaf{"CcsCmFlapMissNum", ccsCmFlapEntry.CcsCmFlapMissNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapCrcErrorNum", types.YLeaf{"CcsCmFlapCrcErrorNum", ccsCmFlapEntry.CcsCmFlapCrcErrorNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapPowerAdjustmentNum", types.YLeaf{"CcsCmFlapPowerAdjustmentNum", ccsCmFlapEntry.CcsCmFlapPowerAdjustmentNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapTotalNum", types.YLeaf{"CcsCmFlapTotalNum", ccsCmFlapEntry.CcsCmFlapTotalNum})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapResetNow", types.YLeaf{"CcsCmFlapResetNow", ccsCmFlapEntry.CcsCmFlapResetNow})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapLastResetTime", types.YLeaf{"CcsCmFlapLastResetTime", ccsCmFlapEntry.CcsCmFlapLastResetTime})
    ccsCmFlapEntry.EntityData.Leafs.Append("ccsCmFlapRowStatus", types.YLeaf{"CcsCmFlapRowStatus", ccsCmFlapEntry.CcsCmFlapRowStatus})

    ccsCmFlapEntry.EntityData.YListKeys = []string {"CcsCmFlapDownstreamIfIndex", "CcsCmFlapUpstreamIfIndex", "CcsCmFlapMacAddr"}

    return &(ccsCmFlapEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable
// This table contains the spectrum data requests.
// 
// There are two types of request: background noise and SNR.
// Refer to ccsSpectrumRequestIfIndex and ccsSpectrumRequestMacAddr
// DESCRIPTIONS on how the type of request is determined.
type CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a spectrum data request.  The management system uses
    // ccsSpectrumRequestStatus to control entry modification, creation, and
    // deletion.  Setting ccsSpectrumRequestEntry to 'destroy' causes entry and
    // its associated data (example: ccsSpectrumDataEntry) to be cleaned up
    // properly.  It is suggested the entry to be set to 'destroy' when the row is
    // no longer in use. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry.
    CcsSpectrumRequestEntry []*CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry
}

func (ccsSpectrumRequestTable *CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable) GetEntityData() *types.CommonEntityData {
    ccsSpectrumRequestTable.EntityData.YFilter = ccsSpectrumRequestTable.YFilter
    ccsSpectrumRequestTable.EntityData.YangName = "ccsSpectrumRequestTable"
    ccsSpectrumRequestTable.EntityData.BundleName = "cisco_ios_xe"
    ccsSpectrumRequestTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsSpectrumRequestTable.EntityData.SegmentPath = "ccsSpectrumRequestTable"
    ccsSpectrumRequestTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsSpectrumRequestTable.EntityData.SegmentPath
    ccsSpectrumRequestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpectrumRequestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpectrumRequestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpectrumRequestTable.EntityData.Children = types.NewOrderedMap()
    ccsSpectrumRequestTable.EntityData.Children.Append("ccsSpectrumRequestEntry", types.YChild{"CcsSpectrumRequestEntry", nil})
    for i := range ccsSpectrumRequestTable.CcsSpectrumRequestEntry {
        ccsSpectrumRequestTable.EntityData.Children.Append(types.GetSegmentPath(ccsSpectrumRequestTable.CcsSpectrumRequestEntry[i]), types.YChild{"CcsSpectrumRequestEntry", ccsSpectrumRequestTable.CcsSpectrumRequestEntry[i]})
    }
    ccsSpectrumRequestTable.EntityData.Leafs = types.NewOrderedMap()

    ccsSpectrumRequestTable.EntityData.YListKeys = []string {}

    return &(ccsSpectrumRequestTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry
// Information about a spectrum data request.  The management
// system uses ccsSpectrumRequestStatus to control entry
// modification, creation, and deletion.
// 
// Setting ccsSpectrumRequestEntry to 'destroy' causes entry
// and its associated data (example: ccsSpectrumDataEntry)
// to be cleaned up properly.  It is suggested the entry
// to be set to 'destroy' when the row is no longer in use.
type CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An arbitrary integer to uniquely identify the
    // entry. The type is interface{} with range: 1..100.
    CcsSpectrumRequestIndex interface{}

    // The ifIndex of a docsCableUpstream(129) interface.  The background noise
    // measurement is requested when ccsSpectrumRequestIfIndex is specified.  The
    // receiving power measurement is requested when ccsSpectrumRequestMacAddr is
    // specified; In this case, ccsSpectrumRequestIfIndex is the ifIndex of the
    // remote CM's upstream. The type is interface{} with range: 0..2147483647.
    CcsSpectrumRequestIfIndex interface{}

    // A MAC address that identifies a remote CM.  The default value of
    // 0000.0000.0000 indicates that the background noise will be measured for the
    // upstream.  In this case, ccsSpectrumRequestIfIndex must be specified. 
    // Other values indicate that the receiving power test is requested for the
    // ccsSpectrumRequestMacAddr with CM signals. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CcsSpectrumRequestMacAddr interface{}

    // Start of frequency range.  The ccsSpectrumRequestLowFreq is adjusted
    // slightly to accurately represent the actual starting point of the frequency
    // range.  The adjustment is done as follows:    aFactor = (center frequency -
    // ccsSpectrumRequestLowFreq)/12K   ccsSpectrumRequestLowFreq = center
    // frequency - (aFactor * 12K)  where 12K is the FFT's bin size. The type is
    // interface{} with range: 5000..85000. Units are KHz.
    CcsSpectrumRequestLowFreq interface{}

    // End of frequency range.  With the adjustment done to the
    // ccsSpectrumRequestLowFreq, ccsSpectrumRequestUpperFreq will also be
    // adjusted to the last frequency within the specified range divisible by the
    // bin size.  Refer to the ccsSpectrumRequestLowFreq DESCRIPTION for the
    // adjustment calculation. The type is interface{} with range: 5000..85000.
    // Units are KHz.
    CcsSpectrumRequestUpperFreq interface{}

    // A span between two frequencies.  ccsSpectrumRequestResolution dictates the
    // amount of receiving power data to be returned in ccsSpectrumDataTable. The
    // finer the resolution, the more data returned.  ccsSpectrumRequestResolution
    // is adjusted to a value which is divisible by FFT's 12KHz bin size. The type
    // is interface{} with range: 12..37000. Units are KHz.
    CcsSpectrumRequestResolution interface{}

    // The control that allows 'start' or 'abort' of the test.  Since there is
    // only 1 FFT engine running on the CMTS, 'start' changes
    // ccsSpectrumRequestOperState to 'pending' state if the FFT is busy;
    // Otherwise, it changes ccsSpectrumRequestOperState to 'running'.  'abort'
    // changes ccsSpectrumRequestOperState to 'aborted' state.  'abort' is only
    // allowed when ccsSpectrumRequestOperState is in 'pending' state.  Only
    // 'start' when request is to be started and 'abort' when request is to be
    // aborted can be set by the user. It is set to 'none' only on completion of
    // the request by the FFT engine.  Note: The SNMP SET is rejected if
    // ccsSpectrumRequestStatus is not 'active'. The type is CCSRequestOperation.
    CcsSpectrumRequestOperation interface{}

    // The operational state of the test.  ccsSpectrumRequestIfIndex,
    // ccsSpectrumRequestMacAddr, ccsSpectrumRequestUpperFreq,
    // ccsSpectrumRequestLowFreq and ccsSpectrumRequestResolution cannot be
    // changed when CCSRequestOperState is in the 'running' state.  For a detailed
    // description, see the CCSRequestOperState DESCRIPTION. The type is
    // CCSRequestOperState.
    CcsSpectrumRequestOperState interface{}

    // The value of sysUpTime when the spectrum measurement operation starts. The
    // type is interface{} with range: 0..4294967295.
    CcsSpectrumRequestStartTime interface{}

    // The value of sysUpTime when the spectrum measurement operation stops. The
    // type is interface{} with range: 0..4294967295.
    CcsSpectrumRequestStoppedTime interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules, see the ccsSpectrumRequestEntry DESCRIPTION. The type
    // is RowStatus.
    CcsSpectrumRequestStatus interface{}
}

func (ccsSpectrumRequestEntry *CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry) GetEntityData() *types.CommonEntityData {
    ccsSpectrumRequestEntry.EntityData.YFilter = ccsSpectrumRequestEntry.YFilter
    ccsSpectrumRequestEntry.EntityData.YangName = "ccsSpectrumRequestEntry"
    ccsSpectrumRequestEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsSpectrumRequestEntry.EntityData.ParentYangName = "ccsSpectrumRequestTable"
    ccsSpectrumRequestEntry.EntityData.SegmentPath = "ccsSpectrumRequestEntry" + types.AddKeyToken(ccsSpectrumRequestEntry.CcsSpectrumRequestIndex, "ccsSpectrumRequestIndex")
    ccsSpectrumRequestEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsSpectrumRequestTable/" + ccsSpectrumRequestEntry.EntityData.SegmentPath
    ccsSpectrumRequestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpectrumRequestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpectrumRequestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpectrumRequestEntry.EntityData.Children = types.NewOrderedMap()
    ccsSpectrumRequestEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestIndex", types.YLeaf{"CcsSpectrumRequestIndex", ccsSpectrumRequestEntry.CcsSpectrumRequestIndex})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestIfIndex", types.YLeaf{"CcsSpectrumRequestIfIndex", ccsSpectrumRequestEntry.CcsSpectrumRequestIfIndex})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestMacAddr", types.YLeaf{"CcsSpectrumRequestMacAddr", ccsSpectrumRequestEntry.CcsSpectrumRequestMacAddr})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestLowFreq", types.YLeaf{"CcsSpectrumRequestLowFreq", ccsSpectrumRequestEntry.CcsSpectrumRequestLowFreq})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestUpperFreq", types.YLeaf{"CcsSpectrumRequestUpperFreq", ccsSpectrumRequestEntry.CcsSpectrumRequestUpperFreq})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestResolution", types.YLeaf{"CcsSpectrumRequestResolution", ccsSpectrumRequestEntry.CcsSpectrumRequestResolution})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestOperation", types.YLeaf{"CcsSpectrumRequestOperation", ccsSpectrumRequestEntry.CcsSpectrumRequestOperation})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestOperState", types.YLeaf{"CcsSpectrumRequestOperState", ccsSpectrumRequestEntry.CcsSpectrumRequestOperState})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestStartTime", types.YLeaf{"CcsSpectrumRequestStartTime", ccsSpectrumRequestEntry.CcsSpectrumRequestStartTime})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestStoppedTime", types.YLeaf{"CcsSpectrumRequestStoppedTime", ccsSpectrumRequestEntry.CcsSpectrumRequestStoppedTime})
    ccsSpectrumRequestEntry.EntityData.Leafs.Append("ccsSpectrumRequestStatus", types.YLeaf{"CcsSpectrumRequestStatus", ccsSpectrumRequestEntry.CcsSpectrumRequestStatus})

    ccsSpectrumRequestEntry.EntityData.YListKeys = []string {"CcsSpectrumRequestIndex"}

    return &(ccsSpectrumRequestEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable
// This table contains the receiving power or background
// noise measurement based on the criteria that is set in
// the ccsSpectrumRequestEntry.
type CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the receiving power or background noise measured at a
    // particular frequency for the ccsSpectrumRequestEntry. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable_CcsSpectrumDataEntry.
    CcsSpectrumDataEntry []*CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable_CcsSpectrumDataEntry
}

func (ccsSpectrumDataTable *CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable) GetEntityData() *types.CommonEntityData {
    ccsSpectrumDataTable.EntityData.YFilter = ccsSpectrumDataTable.YFilter
    ccsSpectrumDataTable.EntityData.YangName = "ccsSpectrumDataTable"
    ccsSpectrumDataTable.EntityData.BundleName = "cisco_ios_xe"
    ccsSpectrumDataTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsSpectrumDataTable.EntityData.SegmentPath = "ccsSpectrumDataTable"
    ccsSpectrumDataTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsSpectrumDataTable.EntityData.SegmentPath
    ccsSpectrumDataTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpectrumDataTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpectrumDataTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpectrumDataTable.EntityData.Children = types.NewOrderedMap()
    ccsSpectrumDataTable.EntityData.Children.Append("ccsSpectrumDataEntry", types.YChild{"CcsSpectrumDataEntry", nil})
    for i := range ccsSpectrumDataTable.CcsSpectrumDataEntry {
        ccsSpectrumDataTable.EntityData.Children.Append(types.GetSegmentPath(ccsSpectrumDataTable.CcsSpectrumDataEntry[i]), types.YChild{"CcsSpectrumDataEntry", ccsSpectrumDataTable.CcsSpectrumDataEntry[i]})
    }
    ccsSpectrumDataTable.EntityData.Leafs = types.NewOrderedMap()

    ccsSpectrumDataTable.EntityData.YListKeys = []string {}

    return &(ccsSpectrumDataTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable_CcsSpectrumDataEntry
// Information about the receiving power or background noise
// measured at a particular frequency for the
// ccsSpectrumRequestEntry.
type CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable_CcsSpectrumDataEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..100. Refers to
    // cisco_cable_spectrum_mib.CISCOCABLESPECTRUMMIB_CcsSpectrumRequestTable_CcsSpectrumRequestEntry_CcsSpectrumRequestIndex
    CcsSpectrumRequestIndex interface{}

    // This attribute is a key. ccsSpectrumDataPower measurement frequency.  Due
    // to the adjustment calculation the starting frequency range for the actual
    // measured frequency if off comparing to the configured frequency.  Refer to
    // ccsSpectrumRequestLowFreq DESCRIPTIONS for the adjustment calculation. The
    // type is interface{} with range: 4000..85000. Units are KHz.
    CcsSpectrumDataFreq interface{}

    // The receiving power measured at the ccsSpectrumDataFreq. The type is
    // interface{} with range: -50..50. Units are dBmV.
    CcsSpectrumDataPower interface{}
}

func (ccsSpectrumDataEntry *CISCOCABLESPECTRUMMIB_CcsSpectrumDataTable_CcsSpectrumDataEntry) GetEntityData() *types.CommonEntityData {
    ccsSpectrumDataEntry.EntityData.YFilter = ccsSpectrumDataEntry.YFilter
    ccsSpectrumDataEntry.EntityData.YangName = "ccsSpectrumDataEntry"
    ccsSpectrumDataEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsSpectrumDataEntry.EntityData.ParentYangName = "ccsSpectrumDataTable"
    ccsSpectrumDataEntry.EntityData.SegmentPath = "ccsSpectrumDataEntry" + types.AddKeyToken(ccsSpectrumDataEntry.CcsSpectrumRequestIndex, "ccsSpectrumRequestIndex") + types.AddKeyToken(ccsSpectrumDataEntry.CcsSpectrumDataFreq, "ccsSpectrumDataFreq")
    ccsSpectrumDataEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsSpectrumDataTable/" + ccsSpectrumDataEntry.EntityData.SegmentPath
    ccsSpectrumDataEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpectrumDataEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpectrumDataEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpectrumDataEntry.EntityData.Children = types.NewOrderedMap()
    ccsSpectrumDataEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsSpectrumDataEntry.EntityData.Leafs.Append("ccsSpectrumRequestIndex", types.YLeaf{"CcsSpectrumRequestIndex", ccsSpectrumDataEntry.CcsSpectrumRequestIndex})
    ccsSpectrumDataEntry.EntityData.Leafs.Append("ccsSpectrumDataFreq", types.YLeaf{"CcsSpectrumDataFreq", ccsSpectrumDataEntry.CcsSpectrumDataFreq})
    ccsSpectrumDataEntry.EntityData.Leafs.Append("ccsSpectrumDataPower", types.YLeaf{"CcsSpectrumDataPower", ccsSpectrumDataEntry.CcsSpectrumDataPower})

    ccsSpectrumDataEntry.EntityData.YListKeys = []string {"CcsSpectrumRequestIndex", "CcsSpectrumDataFreq"}

    return &(ccsSpectrumDataEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSNRRequestTable
// A table of CNR requests.
type CISCOCABLESPECTRUMMIB_CcsSNRRequestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an CNR request.  The management system uses
    // ccsSNRRequestStatus to control entry modification, creation, and deletion.
    // The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsSNRRequestTable_CcsSNRRequestEntry.
    CcsSNRRequestEntry []*CISCOCABLESPECTRUMMIB_CcsSNRRequestTable_CcsSNRRequestEntry
}

func (ccsSNRRequestTable *CISCOCABLESPECTRUMMIB_CcsSNRRequestTable) GetEntityData() *types.CommonEntityData {
    ccsSNRRequestTable.EntityData.YFilter = ccsSNRRequestTable.YFilter
    ccsSNRRequestTable.EntityData.YangName = "ccsSNRRequestTable"
    ccsSNRRequestTable.EntityData.BundleName = "cisco_ios_xe"
    ccsSNRRequestTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsSNRRequestTable.EntityData.SegmentPath = "ccsSNRRequestTable"
    ccsSNRRequestTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsSNRRequestTable.EntityData.SegmentPath
    ccsSNRRequestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSNRRequestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSNRRequestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSNRRequestTable.EntityData.Children = types.NewOrderedMap()
    ccsSNRRequestTable.EntityData.Children.Append("ccsSNRRequestEntry", types.YChild{"CcsSNRRequestEntry", nil})
    for i := range ccsSNRRequestTable.CcsSNRRequestEntry {
        ccsSNRRequestTable.EntityData.Children.Append(types.GetSegmentPath(ccsSNRRequestTable.CcsSNRRequestEntry[i]), types.YChild{"CcsSNRRequestEntry", ccsSNRRequestTable.CcsSNRRequestEntry[i]})
    }
    ccsSNRRequestTable.EntityData.Leafs = types.NewOrderedMap()

    ccsSNRRequestTable.EntityData.YListKeys = []string {}

    return &(ccsSNRRequestTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSNRRequestTable_CcsSNRRequestEntry
// Information about an CNR request.  The management
// system uses ccsSNRRequestStatus to control entry
// modification, creation, and deletion.
type CISCOCABLESPECTRUMMIB_CcsSNRRequestTable_CcsSNRRequestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An arbitrary integer to uniquely identify this
    // entry. The type is interface{} with range: 1..100.
    CcsSNRRequestIndex interface{}

    // A MAC address that identifies the remote online CM that the CNR measurement
    // operation is being performed on. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CcsSNRRequestMacAddr interface{}

    // A snap shot of the CNR value that is measured over the in-use band
    // frequency.  The ccsSNRRequestSNR is set to 0 when ccsSNRRequestOperState is
    // in the 'running' state. The type is interface{} with range: -100..100.
    // Units are dB.
    CcsSNRRequestSNR interface{}

    // The control that allows start or abort of the test.  Since there is only 1
    // FFT engine running on the CMTS, 'start' changes ccsSNRRequestOperState to
    // 'pending' state if the FFT is busy; Otherwise, it changes
    // ccsSNRRequestOperState to 'running'.  'abort' changes
    // ccsSNRRequestOperState to 'aborted' state.  Only 'start' when request is to
    // be started and 'abort' when request is to be aborted can be set by the
    // user. It is set to 'none' only on completion of the request by the FFT
    // engine. The type is CCSRequestOperation.
    CcsSNRRequestOperation interface{}

    // The operational state of the test.  ccsSNRRequestMacAddr, cannot be changed
    // when the ccsSNRRequestOperState is in the 'running' state. The type is
    // CCSRequestOperState.
    CcsSNRRequestOperState interface{}

    // The value of sysUpTime when the CNR measurement operation starts. The type
    // is interface{} with range: 0..4294967295.
    CcsSNRRequestStartTime interface{}

    // The value of sysUpTime when the CNR measurement operation stops. The type
    // is interface{} with range: 0..4294967295.
    CcsSNRRequestStoppedTime interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the ccsSpectrumRequestEntry DESCRIPTION. The type is
    // RowStatus.
    CcsSNRRequestStatus interface{}
}

func (ccsSNRRequestEntry *CISCOCABLESPECTRUMMIB_CcsSNRRequestTable_CcsSNRRequestEntry) GetEntityData() *types.CommonEntityData {
    ccsSNRRequestEntry.EntityData.YFilter = ccsSNRRequestEntry.YFilter
    ccsSNRRequestEntry.EntityData.YangName = "ccsSNRRequestEntry"
    ccsSNRRequestEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsSNRRequestEntry.EntityData.ParentYangName = "ccsSNRRequestTable"
    ccsSNRRequestEntry.EntityData.SegmentPath = "ccsSNRRequestEntry" + types.AddKeyToken(ccsSNRRequestEntry.CcsSNRRequestIndex, "ccsSNRRequestIndex")
    ccsSNRRequestEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsSNRRequestTable/" + ccsSNRRequestEntry.EntityData.SegmentPath
    ccsSNRRequestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSNRRequestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSNRRequestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSNRRequestEntry.EntityData.Children = types.NewOrderedMap()
    ccsSNRRequestEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestIndex", types.YLeaf{"CcsSNRRequestIndex", ccsSNRRequestEntry.CcsSNRRequestIndex})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestMacAddr", types.YLeaf{"CcsSNRRequestMacAddr", ccsSNRRequestEntry.CcsSNRRequestMacAddr})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestSNR", types.YLeaf{"CcsSNRRequestSNR", ccsSNRRequestEntry.CcsSNRRequestSNR})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestOperation", types.YLeaf{"CcsSNRRequestOperation", ccsSNRRequestEntry.CcsSNRRequestOperation})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestOperState", types.YLeaf{"CcsSNRRequestOperState", ccsSNRRequestEntry.CcsSNRRequestOperState})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestStartTime", types.YLeaf{"CcsSNRRequestStartTime", ccsSNRRequestEntry.CcsSNRRequestStartTime})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestStoppedTime", types.YLeaf{"CcsSNRRequestStoppedTime", ccsSNRRequestEntry.CcsSNRRequestStoppedTime})
    ccsSNRRequestEntry.EntityData.Leafs.Append("ccsSNRRequestStatus", types.YLeaf{"CcsSNRRequestStatus", ccsSNRRequestEntry.CcsSNRRequestStatus})

    ccsSNRRequestEntry.EntityData.YListKeys = []string {"CcsSNRRequestIndex"}

    return &(ccsSNRRequestEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable
// This table contains the cable upstream interfaces assigned to
// a spectrum group. A spectrum group contains one or more fixed
// frequencies or frequency bands which can be assigned to cable
// upstream interfaces in the spectrum group.
type CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in ccsUpInSpecGroupTable table. Each entry represents a cable
    // upstream interface assigned to a spectrum group. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry.
    CcsUpInSpecGroupEntry []*CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry
}

func (ccsUpInSpecGroupTable *CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable) GetEntityData() *types.CommonEntityData {
    ccsUpInSpecGroupTable.EntityData.YFilter = ccsUpInSpecGroupTable.YFilter
    ccsUpInSpecGroupTable.EntityData.YangName = "ccsUpInSpecGroupTable"
    ccsUpInSpecGroupTable.EntityData.BundleName = "cisco_ios_xe"
    ccsUpInSpecGroupTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsUpInSpecGroupTable.EntityData.SegmentPath = "ccsUpInSpecGroupTable"
    ccsUpInSpecGroupTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsUpInSpecGroupTable.EntityData.SegmentPath
    ccsUpInSpecGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpInSpecGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpInSpecGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpInSpecGroupTable.EntityData.Children = types.NewOrderedMap()
    ccsUpInSpecGroupTable.EntityData.Children.Append("ccsUpInSpecGroupEntry", types.YChild{"CcsUpInSpecGroupEntry", nil})
    for i := range ccsUpInSpecGroupTable.CcsUpInSpecGroupEntry {
        ccsUpInSpecGroupTable.EntityData.Children.Append(types.GetSegmentPath(ccsUpInSpecGroupTable.CcsUpInSpecGroupEntry[i]), types.YChild{"CcsUpInSpecGroupEntry", ccsUpInSpecGroupTable.CcsUpInSpecGroupEntry[i]})
    }
    ccsUpInSpecGroupTable.EntityData.Leafs = types.NewOrderedMap()

    ccsUpInSpecGroupTable.EntityData.YListKeys = []string {}

    return &(ccsUpInSpecGroupTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry
// An entry in ccsUpInSpecGroupTable table. Each entry represents
// a cable upstream interface assigned to a spectrum group.
type CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The spectrum group number. The value of the object
    // is same as the value of ccsUpSpecMgmtSpecGroup object except value 0. The
    // type is interface{} with range: 1..4294967295.
    CcsSpecGroupNumber interface{}

    // This attribute is a key. The ifIndex of the Cable upstream interface
    // belonging to this Spectrum Group. The value of the corresponding instance
    // of ifType is 'docsCableUpstream(129)'. The type is interface{} with range:
    // 1..2147483647.
    CcsSpecGroupUpstreamIfIndex interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcsSpecGroupUpstreamStorage interface{}

    // The status for this conceptual row. This object is used for
    // creating/deleting entries in ccsUpInSpecGroupTable. The type is RowStatus.
    CcsSpecGroupUpstreamRowStatus interface{}
}

func (ccsUpInSpecGroupEntry *CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry) GetEntityData() *types.CommonEntityData {
    ccsUpInSpecGroupEntry.EntityData.YFilter = ccsUpInSpecGroupEntry.YFilter
    ccsUpInSpecGroupEntry.EntityData.YangName = "ccsUpInSpecGroupEntry"
    ccsUpInSpecGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsUpInSpecGroupEntry.EntityData.ParentYangName = "ccsUpInSpecGroupTable"
    ccsUpInSpecGroupEntry.EntityData.SegmentPath = "ccsUpInSpecGroupEntry" + types.AddKeyToken(ccsUpInSpecGroupEntry.CcsSpecGroupNumber, "ccsSpecGroupNumber") + types.AddKeyToken(ccsUpInSpecGroupEntry.CcsSpecGroupUpstreamIfIndex, "ccsSpecGroupUpstreamIfIndex")
    ccsUpInSpecGroupEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsUpInSpecGroupTable/" + ccsUpInSpecGroupEntry.EntityData.SegmentPath
    ccsUpInSpecGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpInSpecGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpInSpecGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpInSpecGroupEntry.EntityData.Children = types.NewOrderedMap()
    ccsUpInSpecGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsUpInSpecGroupEntry.EntityData.Leafs.Append("ccsSpecGroupNumber", types.YLeaf{"CcsSpecGroupNumber", ccsUpInSpecGroupEntry.CcsSpecGroupNumber})
    ccsUpInSpecGroupEntry.EntityData.Leafs.Append("ccsSpecGroupUpstreamIfIndex", types.YLeaf{"CcsSpecGroupUpstreamIfIndex", ccsUpInSpecGroupEntry.CcsSpecGroupUpstreamIfIndex})
    ccsUpInSpecGroupEntry.EntityData.Leafs.Append("ccsSpecGroupUpstreamStorage", types.YLeaf{"CcsSpecGroupUpstreamStorage", ccsUpInSpecGroupEntry.CcsSpecGroupUpstreamStorage})
    ccsUpInSpecGroupEntry.EntityData.Leafs.Append("ccsSpecGroupUpstreamRowStatus", types.YLeaf{"CcsSpecGroupUpstreamRowStatus", ccsUpInSpecGroupEntry.CcsSpecGroupUpstreamRowStatus})

    ccsUpInSpecGroupEntry.EntityData.YListKeys = []string {"CcsSpecGroupNumber", "CcsSpecGroupUpstreamIfIndex"}

    return &(ccsUpInSpecGroupEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable
// This table contains the cable upstream interfaces in a
// fiber-node.Each fiber-node uniquely represents an RF
// domain.Cable upstream interfaces in the same fiber-node
// are physically combined together into the same RF domain.
type CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in ccsUpInFiberNodeTable. Each entry represents a cable upstream
    // interface assigned to a fiber-node. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable_CcsUpInFiberNodeEntry.
    CcsUpInFiberNodeEntry []*CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable_CcsUpInFiberNodeEntry
}

func (ccsUpInFiberNodeTable *CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable) GetEntityData() *types.CommonEntityData {
    ccsUpInFiberNodeTable.EntityData.YFilter = ccsUpInFiberNodeTable.YFilter
    ccsUpInFiberNodeTable.EntityData.YangName = "ccsUpInFiberNodeTable"
    ccsUpInFiberNodeTable.EntityData.BundleName = "cisco_ios_xe"
    ccsUpInFiberNodeTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsUpInFiberNodeTable.EntityData.SegmentPath = "ccsUpInFiberNodeTable"
    ccsUpInFiberNodeTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsUpInFiberNodeTable.EntityData.SegmentPath
    ccsUpInFiberNodeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpInFiberNodeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpInFiberNodeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpInFiberNodeTable.EntityData.Children = types.NewOrderedMap()
    ccsUpInFiberNodeTable.EntityData.Children.Append("ccsUpInFiberNodeEntry", types.YChild{"CcsUpInFiberNodeEntry", nil})
    for i := range ccsUpInFiberNodeTable.CcsUpInFiberNodeEntry {
        ccsUpInFiberNodeTable.EntityData.Children.Append(types.GetSegmentPath(ccsUpInFiberNodeTable.CcsUpInFiberNodeEntry[i]), types.YChild{"CcsUpInFiberNodeEntry", ccsUpInFiberNodeTable.CcsUpInFiberNodeEntry[i]})
    }
    ccsUpInFiberNodeTable.EntityData.Leafs = types.NewOrderedMap()

    ccsUpInFiberNodeTable.EntityData.YListKeys = []string {}

    return &(ccsUpInFiberNodeTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable_CcsUpInFiberNodeEntry
// An entry in ccsUpInFiberNodeTable. Each entry represents a
// cable upstream interface assigned to a fiber-node.
type CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable_CcsUpInFiberNodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The fiber-node number. The value of the object is
    // same as the value of ccsUpSpecMgmtSharedSpectrum except value 0. The type
    // is interface{} with range: 1..4294967295.
    CcsFiberNodeNumber interface{}

    // This attribute is a key. The ifIndex of the Cable upstream interface
    // belonging to this Spectrum Group. The value of the corresponding instance
    // of ifType is 'docsCableUpstream(129)'. The type is interface{} with range:
    // 1..2147483647.
    CcsFiberNodeUpstreamIfIndex interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcsFiberNodeUpstreamStorage interface{}

    // The status for this conceptual row. This object is used for
    // creating/deleting entries in ccsUpInFiberNodeTable. The type is RowStatus.
    CcsFiberNodeUpstreamRowStatus interface{}
}

func (ccsUpInFiberNodeEntry *CISCOCABLESPECTRUMMIB_CcsUpInFiberNodeTable_CcsUpInFiberNodeEntry) GetEntityData() *types.CommonEntityData {
    ccsUpInFiberNodeEntry.EntityData.YFilter = ccsUpInFiberNodeEntry.YFilter
    ccsUpInFiberNodeEntry.EntityData.YangName = "ccsUpInFiberNodeEntry"
    ccsUpInFiberNodeEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsUpInFiberNodeEntry.EntityData.ParentYangName = "ccsUpInFiberNodeTable"
    ccsUpInFiberNodeEntry.EntityData.SegmentPath = "ccsUpInFiberNodeEntry" + types.AddKeyToken(ccsUpInFiberNodeEntry.CcsFiberNodeNumber, "ccsFiberNodeNumber") + types.AddKeyToken(ccsUpInFiberNodeEntry.CcsFiberNodeUpstreamIfIndex, "ccsFiberNodeUpstreamIfIndex")
    ccsUpInFiberNodeEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsUpInFiberNodeTable/" + ccsUpInFiberNodeEntry.EntityData.SegmentPath
    ccsUpInFiberNodeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpInFiberNodeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpInFiberNodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpInFiberNodeEntry.EntityData.Children = types.NewOrderedMap()
    ccsUpInFiberNodeEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsUpInFiberNodeEntry.EntityData.Leafs.Append("ccsFiberNodeNumber", types.YLeaf{"CcsFiberNodeNumber", ccsUpInFiberNodeEntry.CcsFiberNodeNumber})
    ccsUpInFiberNodeEntry.EntityData.Leafs.Append("ccsFiberNodeUpstreamIfIndex", types.YLeaf{"CcsFiberNodeUpstreamIfIndex", ccsUpInFiberNodeEntry.CcsFiberNodeUpstreamIfIndex})
    ccsUpInFiberNodeEntry.EntityData.Leafs.Append("ccsFiberNodeUpstreamStorage", types.YLeaf{"CcsFiberNodeUpstreamStorage", ccsUpInFiberNodeEntry.CcsFiberNodeUpstreamStorage})
    ccsUpInFiberNodeEntry.EntityData.Leafs.Append("ccsFiberNodeUpstreamRowStatus", types.YLeaf{"CcsFiberNodeUpstreamRowStatus", ccsUpInFiberNodeEntry.CcsFiberNodeUpstreamRowStatus})

    ccsUpInFiberNodeEntry.EntityData.YListKeys = []string {"CcsFiberNodeNumber", "CcsFiberNodeUpstreamIfIndex"}

    return &(ccsUpInFiberNodeEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable
// This table contains the attributes of the cable
// upstream interfaces, ifType of docsCableUpstream(129),
// to be used for improving performance and proactive
// hopping.
// 
// Proactive hopping is achieved by setting the SNR 
// polling period over the in-use band without CM
// signals.
type CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Upstream interface's spectrum management information. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry.
    CcsUpSpecMgmtEntry []*CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry
}

func (ccsUpSpecMgmtTable *CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable) GetEntityData() *types.CommonEntityData {
    ccsUpSpecMgmtTable.EntityData.YFilter = ccsUpSpecMgmtTable.YFilter
    ccsUpSpecMgmtTable.EntityData.YangName = "ccsUpSpecMgmtTable"
    ccsUpSpecMgmtTable.EntityData.BundleName = "cisco_ios_xe"
    ccsUpSpecMgmtTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsUpSpecMgmtTable.EntityData.SegmentPath = "ccsUpSpecMgmtTable"
    ccsUpSpecMgmtTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsUpSpecMgmtTable.EntityData.SegmentPath
    ccsUpSpecMgmtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpSpecMgmtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpSpecMgmtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpSpecMgmtTable.EntityData.Children = types.NewOrderedMap()
    ccsUpSpecMgmtTable.EntityData.Children.Append("ccsUpSpecMgmtEntry", types.YChild{"CcsUpSpecMgmtEntry", nil})
    for i := range ccsUpSpecMgmtTable.CcsUpSpecMgmtEntry {
        ccsUpSpecMgmtTable.EntityData.Children.Append(types.GetSegmentPath(ccsUpSpecMgmtTable.CcsUpSpecMgmtEntry[i]), types.YChild{"CcsUpSpecMgmtEntry", ccsUpSpecMgmtTable.CcsUpSpecMgmtEntry[i]})
    }
    ccsUpSpecMgmtTable.EntityData.Leafs = types.NewOrderedMap()

    ccsUpSpecMgmtTable.EntityData.YListKeys = []string {}

    return &(ccsUpSpecMgmtTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry
// Upstream interface's spectrum management information.
type CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // A preference priority for changing the frequency, modulation, or channel
    // width supporting the automatic switching of the modulation scheme when the
    // channel becomes noisy.  The default priority is frequency, modulation, and
    // channel width.  With the default preference, the frequency is changed if
    // there is a clean band available.  If there's no clean band available, the
    // modulation is changed.  And if the clean band is still not available, the
    // bandwidth is reduced until an acceptable band is found or a minimum
    // bandwidth of 200KHz. The type is CcsUpSpecMgmtHopPriority.
    CcsUpSpecMgmtHopPriority interface{}

    // The Signal to Noise (SNR) threshold.  This object is applicable for
    // modulation profile 1.  When the CMTS detects that the SNR goes lower than
    // ccsUpSpecMgmtSnrThres1, it switches to profile 2. Therefore,
    // ccsUpSpecMgmtSnrThres1 should be larger than ccsUpSpecMgmtSnrThres2.  A
    // value 0 indicates to bypass the threshold check. The type is interface{}
    // with range: 0..None | 5..35. Units are dB.
    CcsUpSpecMgmtSnrThres1 interface{}

    // The Signal to Noise (SNR) threshold.  This object is applicable for
    // modulation profile 2.  Refer to ccsUpSpecMgmtCriteria on how 
    // ccsUpSpecMgmtSnrThres2 can trigger a change  in frequency, modulation or
    // channel width.  A value 0 indicates to bypass the threshold check.  Note:
    // The SNMP SET is rejected if both  ccsUpSpecMgmtSnrThres1,
    // ccsUpSpecMgmtSnrThres2 are non-zero and ccsUpSpecMgmtSnrThres2 is higher
    // than ccsUpSpecMgmtSnrThres1. The type is interface{} with range: 0..None |
    // 5..35. Units are dB.
    CcsUpSpecMgmtSnrThres2 interface{}

    // The Forward Error Correction (FEC) correctable count threshold.  This
    // object is applicable for profile 1.  A value 0 indicates to bypass the
    // threshold check.  When CMTS detects that FEC correctable count goes higher
    // than ccsUpSpecMgmtFecCorrectThres1, it switch to Profile 2.  Therefore,
    // ccsUpSpecMgmtFecCorrectThres1 should be smaller than
    // ccsUpSpecMgmtFecCorrectThres2. The type is interface{} with range: 0..30.
    CcsUpSpecMgmtFecCorrectThres1 interface{}

    // The FEC correctable count threshold.  This object is applicable for profile
    // 2.  When CMTS detects that FEC correctable count goes higher than
    // ccsUpSpecMgmtFecCorrectThres2, modulation change can occur, depending on
    // the type of ccsUpSpecMgmtHopPriority.  Note: SNMP SET will be rejected if 
    // ccsUpSpecMgmtFecCorrectThres2 is lower than  ccsUpSpecMgmtFecCorrectThres1.
    // The type is interface{} with range: 1..20.
    CcsUpSpecMgmtFecCorrectThres2 interface{}

    // The FEC uncorrectable count threshold.  This object is applicable for
    // modulation profile 1.  A value 0 indicates to bypass the threshold check. 
    // When CMTS detects that FEC uncorrectable count goes higher than
    // ccsUpSpecMgmtFecUnCorrectThres1, it switches to Profile 2. Therefore,
    // ccsUpSpecMgmtFecUnCorrectThres1 should be smaller than
    // ccsUpSpecMgmtFecUnCorrectThres2. The type is interface{} with range: 0..30.
    CcsUpSpecMgmtFecUnCorrectThres1 interface{}

    // The FEC uncorrectable count threshold.  This object is applicable for
    // modulation profile 2.  A value 0 indicates to bypass the threshold check. 
    // Refer to ccsUpSpecMgmtCriteria on how  ccsUpSpecMgmtFecUnCorrectThres2 can
    // trigger a change  in frequency, modulation or channel width.  Note: SNMP
    // SET is rejected if ccsUpSpecMgmtFecUnCorrectThres2 is lower than
    // ccsUpSpecMgmtFecUnCorrectThres1. The type is interface{} with range: 0..30.
    CcsUpSpecMgmtFecUnCorrectThres2 interface{}

    // A period between SNR pollings.  The SNR is collected from the Fast Fourier
    // Transform (FFT) measurement over the in-use band when there is no CM
    // signals. When the CMTS detects that SNR doesn't meet ccsUpSpecMgmtSnrThres1
    // or ccsUpSpecMgmtSnrThres2, a possible hopping occurs, depending on the type
    // of ccsUpSpecMgmtHopPriority. The type is interface{} with range: 0..None |
    // 500..25000. Units are msec.
    CcsUpSpecMgmtSnrPollPeriod interface{}

    // A condition that triggers hopping.  The SNR condition occurs when SNR does
    // not meet the ccsUpSpecMgmtSnrThres1 or ccsUpSpecMgmtSnrThres2.  The
    // stationMaintainenceMiss condition occurs when the  percentage of offline
    // CMs is reached. The type is CcsUpSpecMgmtHopCondition.
    CcsUpSpecMgmtHopCondition interface{}

    // Center frequency before hopping occurs.  A value 0 indicates that the
    // interface has no frequency assigned. The type is interface{} with range:
    // 0..None | 5000..65000. Units are KHz.
    CcsUpSpecMgmtFromCenterFreq interface{}

    // Current center frequency.  A value 0 indicates that the interface has no
    // frequency assigned. The type is interface{} with range: 0..None |
    // 5000..65000. Units are KHz.
    CcsUpSpecMgmtToCenterFreq interface{}

    // Bandwidth before hopping occurs. The type is interface{} with range:
    // 200..None | 400..None | 800..None | 1600..None | 3200..None | 6400..None.
    // Units are KHz.
    CcsUpSpecMgmtFromBandWidth interface{}

    // Current bandwidth. The type is interface{} with range: 200..None |
    // 400..None | 800..None | 1600..None | 3200..None | 6400..None. Units are
    // KHz.
    CcsUpSpecMgmtToBandWidth interface{}

    // Modulation profile index before hopping occurs. It is the index identical
    // to the docsIfModIndex in the docsIfCmtsModulationTable.  For the detailed
    // descriptions, see the docsIfCmtsModulationTable and docsIfCmtsModIndex
    // DESCRIPTIONS. The type is interface{} with range: 1..2147483647.
    CcsUpSpecMgmtFromModProfile interface{}

    // The current modulation profile index. It is the index identical to the
    // docsIfModIndex in the docsIfCmtsModulationTable.  For the detailed
    // descriptions, see the docsIfCmtsModulationTable and docsIfCmtsModIndex
    // DESCRIPTIONS. The type is interface{} with range: 1..2147483647.
    CcsUpSpecMgmtToModProfile interface{}

    // Current SNR. A value -99 indicates the system detected no  modem on line,
    // and a value of -100 indicates the  system was unable to retrieve the SNR
    // value. The type is interface{} with range: -100..100. Units are dB.
    CcsUpSpecMgmtSNR interface{}

    // Upper bound frequency that the upstream supports. The type is interface{}
    // with range: 5000..85000. Units are KHz.
    CcsUpSpecMgmtUpperBoundFreq interface{}

    // The Carrier to Noise (CNR) threshold.  This object is applicable for
    // modulation profile 1.   When the CMTS detects that the CNR goes lower than
    // ccsUpSpecMgmtCnrThres1, it switches to profile 2. Therefore,
    // ccsUpSpecMgmtCnrThres1 should be larger  than ccsUpSpecMgmtCnrThres2.  A
    // value 0 indicates to bypass the threshold check. The type is interface{}
    // with range: 0..None | 5..35. Units are dB.
    CcsUpSpecMgmtCnrThres1 interface{}

    // The Carrier to Noise (CNR) threshold.  This object is applicable for
    // modulation profile 2.  Refer to ccsUpSpecMgmtCriteria on how 
    // ccsUpSpecMgmtCnrThres2 can trigger a change  in frequency, modulation or
    // channel width.  A value 0 indicates to bypass the threshold check.  Note:
    // The SNMP SET is rejected if both  ccsUpSpecMgmtCnrThres1,
    // ccsUpSpecMgmtCnrThres2 are non-zero and ccsUpSpecMgmtCnrThres2 is higher
    // than ccsUpSpecMgmtCnrThres1. The type is interface{} with range: 0..None |
    // 5..35. Units are dB.
    CcsUpSpecMgmtCnrThres2 interface{}

    // Current CNR. A value -99 indicates the system detected no  modem on line,
    // and a value of -100 indicates the  system was unable to retrieve the CNR
    // value. The type is interface{} with range: -100..100. Units are dB.
    CcsUpSpecMgmtCNR interface{}

    // The missed maintenance message threshold in percentage.  A value 0
    // indicates that the interface has no spectrum  group assigned. i.e.
    // ccsUpSpecMgmtSpecGroup equals 0. The type is interface{} with range:
    // 0..100.
    CcsUpSpecMgmtMissedMaintMsgThres interface{}

    // The minimum time between frequency hops in seconds.  A value 0 indicates
    // that the interface has no spectrum group assigned. i.e.
    // ccsUpSpecMgmtSpecGroup equals to 0. The type is interface{} with range:
    // 0..3600. Units are seconds.
    CcsUpSpecMgmtHopPeriod interface{}

    // Defines the criteria that trigger a change in frequency hopping, modulation
    // or channel width.    Depending on the type of linecards, the criteria  are
    // slightly different:  For the linecards that have no Hardware  Based
    // Spectrum Management capability like  the uBR-MC1xC, change in modulation or
    // frequency occurs in the following criteria:  1) Change from modulation
    // profile 1 to     modulation profile 2 would occur when CMTS     detects the
    // SNR that goes below the     threshold and either corr FEC or uncorr FEC    
    // is above the threshold.  In this case,     snrBelowThres and either
    // corrFecAboveThres or     uncorrFecAboveThres bits will get set.   2) Change
    // from modulation profile 2 to     modulation profile 1 would occur when CMTS
    // detects the SNR goes above the threshold+3    and both corr FEC and uncorr
    // FEC are     below the threshold.  In this case,     snrAboveThres and
    // corrFecBelowThres    and uncorrFecBelowThres bits will get set.   3) Change
    // in frequency or frequency hopping would    occur when CMTS detects no
    // active modem on the link.      In this case, noActiveModem bit will get   
    // set accordingly.   For the linecards that have the Hardware Based  Spectrum
    // Management capability, change in  frequency, modulation or channel width
    // occurs  in the following criteria:  1) when CMTS detects the SNR or CNR
    // goes below the     threshold and either corr FEC or uncorr FEC is     above
    // the threshold.  In this case,     either snrBelowThres or cnrBelowThres and
    // either corrFecAboveThres or uncorrFecAboveThres     bits will get set.   2)
    // when CMTS detects the SNR or CNR goes above the     threshold + 3 and both
    // corr FEC and uncorr FEC are     below the threshold.  In this case,    
    // either snrAboveThres or cnrAboveThres and both     corrFecBelowThres and
    // uncorrFecBelowThres bits     will get set.   3) when CMTS detects no active
    // modem on the link or     uncorr FEC is above
    // ccsUpSpecMgmtFecUnCorrectThres2.     In this case noActiveModem or    
    // uncorrFecAboveSecondThres bit will get set     accordingly.  Note: The
    // order of frequency, modulation or channel  width changes for the advanced
    // spectrum management linecards will be determined based on the selection of 
    // the ccsUpSpecMgmtHopPriority. The type is map[string]bool.
    CcsUpSpecMgmtCriteria interface{}

    // The spectrum group assigned to the upstream. The value of 0 for the object
    // indicates that the upstream has no spectrum group assigned. The type is
    // interface{} with range: 0..4294967295.
    CcsUpSpecMgmtSpecGroup interface{}

    // The fiber-node assigned to the upstreams. Upstreams having same fiber-node
    // number indicates that they physically combine together into same RF domain
    // and must have unique frequency assigned. The value of 0 for the object
    // indicates that the upstream is not physically combine with any others. The
    // type is interface{} with range: 0..4294967295.
    CcsUpSpecMgmtSharedSpectrum interface{}
}

func (ccsUpSpecMgmtEntry *CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry) GetEntityData() *types.CommonEntityData {
    ccsUpSpecMgmtEntry.EntityData.YFilter = ccsUpSpecMgmtEntry.YFilter
    ccsUpSpecMgmtEntry.EntityData.YangName = "ccsUpSpecMgmtEntry"
    ccsUpSpecMgmtEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsUpSpecMgmtEntry.EntityData.ParentYangName = "ccsUpSpecMgmtTable"
    ccsUpSpecMgmtEntry.EntityData.SegmentPath = "ccsUpSpecMgmtEntry" + types.AddKeyToken(ccsUpSpecMgmtEntry.IfIndex, "ifIndex")
    ccsUpSpecMgmtEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsUpSpecMgmtTable/" + ccsUpSpecMgmtEntry.EntityData.SegmentPath
    ccsUpSpecMgmtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsUpSpecMgmtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsUpSpecMgmtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsUpSpecMgmtEntry.EntityData.Children = types.NewOrderedMap()
    ccsUpSpecMgmtEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ccsUpSpecMgmtEntry.IfIndex})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtHopPriority", types.YLeaf{"CcsUpSpecMgmtHopPriority", ccsUpSpecMgmtEntry.CcsUpSpecMgmtHopPriority})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSnrThres1", types.YLeaf{"CcsUpSpecMgmtSnrThres1", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSnrThres1})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSnrThres2", types.YLeaf{"CcsUpSpecMgmtSnrThres2", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSnrThres2})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFecCorrectThres1", types.YLeaf{"CcsUpSpecMgmtFecCorrectThres1", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFecCorrectThres1})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFecCorrectThres2", types.YLeaf{"CcsUpSpecMgmtFecCorrectThres2", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFecCorrectThres2})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFecUnCorrectThres1", types.YLeaf{"CcsUpSpecMgmtFecUnCorrectThres1", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFecUnCorrectThres1})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFecUnCorrectThres2", types.YLeaf{"CcsUpSpecMgmtFecUnCorrectThres2", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFecUnCorrectThres2})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSnrPollPeriod", types.YLeaf{"CcsUpSpecMgmtSnrPollPeriod", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSnrPollPeriod})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtHopCondition", types.YLeaf{"CcsUpSpecMgmtHopCondition", ccsUpSpecMgmtEntry.CcsUpSpecMgmtHopCondition})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFromCenterFreq", types.YLeaf{"CcsUpSpecMgmtFromCenterFreq", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFromCenterFreq})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtToCenterFreq", types.YLeaf{"CcsUpSpecMgmtToCenterFreq", ccsUpSpecMgmtEntry.CcsUpSpecMgmtToCenterFreq})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFromBandWidth", types.YLeaf{"CcsUpSpecMgmtFromBandWidth", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFromBandWidth})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtToBandWidth", types.YLeaf{"CcsUpSpecMgmtToBandWidth", ccsUpSpecMgmtEntry.CcsUpSpecMgmtToBandWidth})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtFromModProfile", types.YLeaf{"CcsUpSpecMgmtFromModProfile", ccsUpSpecMgmtEntry.CcsUpSpecMgmtFromModProfile})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtToModProfile", types.YLeaf{"CcsUpSpecMgmtToModProfile", ccsUpSpecMgmtEntry.CcsUpSpecMgmtToModProfile})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSNR", types.YLeaf{"CcsUpSpecMgmtSNR", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSNR})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtUpperBoundFreq", types.YLeaf{"CcsUpSpecMgmtUpperBoundFreq", ccsUpSpecMgmtEntry.CcsUpSpecMgmtUpperBoundFreq})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtCnrThres1", types.YLeaf{"CcsUpSpecMgmtCnrThres1", ccsUpSpecMgmtEntry.CcsUpSpecMgmtCnrThres1})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtCnrThres2", types.YLeaf{"CcsUpSpecMgmtCnrThres2", ccsUpSpecMgmtEntry.CcsUpSpecMgmtCnrThres2})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtCNR", types.YLeaf{"CcsUpSpecMgmtCNR", ccsUpSpecMgmtEntry.CcsUpSpecMgmtCNR})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtMissedMaintMsgThres", types.YLeaf{"CcsUpSpecMgmtMissedMaintMsgThres", ccsUpSpecMgmtEntry.CcsUpSpecMgmtMissedMaintMsgThres})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtHopPeriod", types.YLeaf{"CcsUpSpecMgmtHopPeriod", ccsUpSpecMgmtEntry.CcsUpSpecMgmtHopPeriod})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtCriteria", types.YLeaf{"CcsUpSpecMgmtCriteria", ccsUpSpecMgmtEntry.CcsUpSpecMgmtCriteria})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSpecGroup", types.YLeaf{"CcsUpSpecMgmtSpecGroup", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSpecGroup})
    ccsUpSpecMgmtEntry.EntityData.Leafs.Append("ccsUpSpecMgmtSharedSpectrum", types.YLeaf{"CcsUpSpecMgmtSharedSpectrum", ccsUpSpecMgmtEntry.CcsUpSpecMgmtSharedSpectrum})

    ccsUpSpecMgmtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(ccsUpSpecMgmtEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition represents percentage of offline CMs is reached.
type CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition string

const (
    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition_snr CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition = "snr"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition_stationMaintainenceMiss CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition = "stationMaintainenceMiss"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition_others CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopCondition = "others"
)

// CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority represents bandwidth of 200KHz.
type CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority string

const (
    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_frqModChannel CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "frqModChannel"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_frqChannelMod CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "frqChannelMod"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_modFrqChannel CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "modFrqChannel"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_modChannelFrq CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "modChannelFrq"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_channelFrqMod CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "channelFrqMod"

    CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority_channelModFrq CISCOCABLESPECTRUMMIB_CcsUpSpecMgmtTable_CcsUpSpecMgmtEntry_CcsUpSpecMgmtHopPriority = "channelModFrq"
)

// CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable
// This table contains the frequency and band configuration
// of the spectrum group.
type CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in ccsSpecGroupFreqTable. Each entry represents a center frequency
    // or a frequency band configured for the spectrum group. The type is slice of
    // CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry.
    CcsSpecGroupFreqEntry []*CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry
}

func (ccsSpecGroupFreqTable *CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable) GetEntityData() *types.CommonEntityData {
    ccsSpecGroupFreqTable.EntityData.YFilter = ccsSpecGroupFreqTable.YFilter
    ccsSpecGroupFreqTable.EntityData.YangName = "ccsSpecGroupFreqTable"
    ccsSpecGroupFreqTable.EntityData.BundleName = "cisco_ios_xe"
    ccsSpecGroupFreqTable.EntityData.ParentYangName = "CISCO-CABLE-SPECTRUM-MIB"
    ccsSpecGroupFreqTable.EntityData.SegmentPath = "ccsSpecGroupFreqTable"
    ccsSpecGroupFreqTable.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/" + ccsSpecGroupFreqTable.EntityData.SegmentPath
    ccsSpecGroupFreqTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpecGroupFreqTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpecGroupFreqTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpecGroupFreqTable.EntityData.Children = types.NewOrderedMap()
    ccsSpecGroupFreqTable.EntityData.Children.Append("ccsSpecGroupFreqEntry", types.YChild{"CcsSpecGroupFreqEntry", nil})
    for i := range ccsSpecGroupFreqTable.CcsSpecGroupFreqEntry {
        ccsSpecGroupFreqTable.EntityData.Children.Append(types.GetSegmentPath(ccsSpecGroupFreqTable.CcsSpecGroupFreqEntry[i]), types.YChild{"CcsSpecGroupFreqEntry", ccsSpecGroupFreqTable.CcsSpecGroupFreqEntry[i]})
    }
    ccsSpecGroupFreqTable.EntityData.Leafs = types.NewOrderedMap()

    ccsSpecGroupFreqTable.EntityData.YListKeys = []string {}

    return &(ccsSpecGroupFreqTable.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry
// An entry in ccsSpecGroupFreqTable. Each entry represents a
// center frequency or a frequency band configured for the
// spectrum group.
type CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_cable_spectrum_mib.CISCOCABLESPECTRUMMIB_CcsUpInSpecGroupTable_CcsUpInSpecGroupEntry_CcsSpecGroupNumber
    CcsSpecGroupNumber interface{}

    // This attribute is a key. An arbitrary index to uniquely identify
    // frequencies or bands configured in spectrum group. The type is interface{}
    // with range: 1..4294967295.
    CcsSpecGroupFreqIndex interface{}

    // The type of the frequency configured in spectrum group. This is a mandatory
    // object and can be modified when the row is active. If the value of
    // centerFreq(1) is specified, the values of the corresponding instance of
    // ccsSpecGroupFreqLower and csSpecGroupFreqUpper object must be identical. If
    // the value of bandFreq(2) is specified, the values of the corresponding
    // instance of ccsSpecGroupFreqLower and csSpecGroupFreqUpper must be unique.
    // The type is CcsSpecGroupFreqType.
    CcsSpecGroupFreqType interface{}

    // The lower frequency configured in spectrum group. This is a mandatory
    // object and can be modified when the row is active. To configure a band
    // frequency in the spectrum group, the value of this object must be lower
    // than the value of the corresponding instance of ccsSpecGroupFreqUpper. To
    // configure a fixed center frequency in the spectrum group, the value of this
    // object must be equal to the value of the corresponding instance of
    // ccsSpecGroupFreqUpper. The type is interface{} with range: 0..1000000000.
    // Units are Hz.
    CcsSpecGroupFreqLower interface{}

    // The upper frequency configured in spectrum group. This is a mandatory
    // object and can be modified when the row is active. To configure a band
    // frequency in the spectrum group, the value of this object must be greater
    // than the value of the corresponding instance of ccsSpecGroupFreqLower. To
    // configure a fixed center frequency in the spectrum group, the value of this
    // object must be equal to the value of the corresponding instance of
    // ccsSpecGroupFreqLower. The type is interface{} with range: 0..1000000000.
    // Units are Hz.
    CcsSpecGroupFreqUpper interface{}

    // The storage type for this conceptual row. The type is StorageType.
    CcsSpecGroupStorage interface{}

    // The status of this conceptual row. This object is used for
    // creating/deleting entries in ccsSpecGroupFreqTable.  A conceptual row may
    // not be created via SNMP without explicitly setting the value of
    // ccsSpecGroupRowStatus to createAndGo(4).  A conceptual row can not have the
    // status of active(1) until proper values have been assigned to the mandatory
    // objects ccsSpecGroupFreqType, ccsSpecGroupFreqLower, and
    // ccsSpecGroupFreqUpper.  A conceptual row may be modified or deleted any
    // time. If the frequency represents by the row has been assigned to a cable
    // upstream interface, modifying or deleting such row would cause the cable
    // upstream interface frequency reassignment. The type is RowStatus.
    CcsSpecGroupRowStatus interface{}
}

func (ccsSpecGroupFreqEntry *CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry) GetEntityData() *types.CommonEntityData {
    ccsSpecGroupFreqEntry.EntityData.YFilter = ccsSpecGroupFreqEntry.YFilter
    ccsSpecGroupFreqEntry.EntityData.YangName = "ccsSpecGroupFreqEntry"
    ccsSpecGroupFreqEntry.EntityData.BundleName = "cisco_ios_xe"
    ccsSpecGroupFreqEntry.EntityData.ParentYangName = "ccsSpecGroupFreqTable"
    ccsSpecGroupFreqEntry.EntityData.SegmentPath = "ccsSpecGroupFreqEntry" + types.AddKeyToken(ccsSpecGroupFreqEntry.CcsSpecGroupNumber, "ccsSpecGroupNumber") + types.AddKeyToken(ccsSpecGroupFreqEntry.CcsSpecGroupFreqIndex, "ccsSpecGroupFreqIndex")
    ccsSpecGroupFreqEntry.EntityData.AbsolutePath = "CISCO-CABLE-SPECTRUM-MIB:CISCO-CABLE-SPECTRUM-MIB/ccsSpecGroupFreqTable/" + ccsSpecGroupFreqEntry.EntityData.SegmentPath
    ccsSpecGroupFreqEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccsSpecGroupFreqEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccsSpecGroupFreqEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccsSpecGroupFreqEntry.EntityData.Children = types.NewOrderedMap()
    ccsSpecGroupFreqEntry.EntityData.Leafs = types.NewOrderedMap()
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupNumber", types.YLeaf{"CcsSpecGroupNumber", ccsSpecGroupFreqEntry.CcsSpecGroupNumber})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupFreqIndex", types.YLeaf{"CcsSpecGroupFreqIndex", ccsSpecGroupFreqEntry.CcsSpecGroupFreqIndex})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupFreqType", types.YLeaf{"CcsSpecGroupFreqType", ccsSpecGroupFreqEntry.CcsSpecGroupFreqType})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupFreqLower", types.YLeaf{"CcsSpecGroupFreqLower", ccsSpecGroupFreqEntry.CcsSpecGroupFreqLower})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupFreqUpper", types.YLeaf{"CcsSpecGroupFreqUpper", ccsSpecGroupFreqEntry.CcsSpecGroupFreqUpper})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupStorage", types.YLeaf{"CcsSpecGroupStorage", ccsSpecGroupFreqEntry.CcsSpecGroupStorage})
    ccsSpecGroupFreqEntry.EntityData.Leafs.Append("ccsSpecGroupRowStatus", types.YLeaf{"CcsSpecGroupRowStatus", ccsSpecGroupFreqEntry.CcsSpecGroupRowStatus})

    ccsSpecGroupFreqEntry.EntityData.YListKeys = []string {"CcsSpecGroupNumber", "CcsSpecGroupFreqIndex"}

    return &(ccsSpecGroupFreqEntry.EntityData)
}

// CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType represents must be unique.
type CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType string

const (
    CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType_centerFreq CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType = "centerFreq"

    CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType_bandFreq CISCOCABLESPECTRUMMIB_CcsSpecGroupFreqTable_CcsSpecGroupFreqEntry_CcsSpecGroupFreqType = "bandFreq"
)

