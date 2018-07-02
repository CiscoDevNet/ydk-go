// This MIB contains managed object definitions for
// encapsulating TDM (T1,E1, T3, E3, NxDS0) as
// pseudo-wires over packet-switching networks (PSN).
// 
// This MIB supplements the CISCO-IETF-PW-MIB.
// The CISCO-IETF-PW-MIB contains structures and MIB
// associations generic to Pseudo-Wire (PW) emulation.
// PW-specific MIBs (such as this) contain config and stats for
// specific PW types.
package cisco_ietf_pw_tdm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_tdm_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-TDM-MIB CISCO-IETF-PW-TDM-MIB}", reflect.TypeOf(CISCOIETFPWTDMMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-TDM-MIB:CISCO-IETF-PW-TDM-MIB", reflect.TypeOf(CISCOIETFPWTDMMIB{}))
}

// CISCOIETFPWTDMMIB
type CISCOIETFPWTDMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CpwCTDMObjects CISCOIETFPWTDMMIB_CpwCTDMObjects

    // This table contains basic information including ifIndex, and pointers to
    // entries in the relevant TDM config tables for this TDM PW.
    CpwCTDMTable CISCOIETFPWTDMMIB_CpwCTDMTable

    // This table contains a set of parameters that may be referenced by one or
    // more TDM PWs in cpwCTDMTable.
    CpwCTDMCfgTable CISCOIETFPWTDMMIB_CpwCTDMCfgTable

    // This table provides TDM PW performance information. This includes current
    // 15 minute interval counts.   The table includes counters that work together
    // to integrate errors and the lack of errors on the TDM PW. An error is
    // caused by a missing packet. Missing packet can be a result of, packet loss
    // in the network, (uncorrectable) packet out of sequence, packet length
    // error, jitter buffer overflow, and jitter buffer underflow. The result is
    // declaring whether or not the TDM PW is in Loss of Packet (LOPS) state.
    CpwCTDMPerfCurrentTable CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable

    // This table provides performance information per TDM PW similar to the
    // cpwCTDMPerfCurrentTable above. However, these counts represent historical
    // 15 minute intervals. Typically, this table will have a maximum of 96
    // entries for a 24 hour period, but is not limited to this.
    CpwCTDMPerfIntervalTable CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable

    // This table provides performance information per TDM PW similar to the
    // cpwCTDMPerfIntervalTable above. However, these counters represent
    // historical 1 day intervals up to one full month. The table consists of real
    // time data, as such it is not persistence across re-boot.
    CpwCTDMPerf1DayIntervalTable CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWTDMMIB.EntityData.YFilter = cISCOIETFPWTDMMIB.YFilter
    cISCOIETFPWTDMMIB.EntityData.YangName = "CISCO-IETF-PW-TDM-MIB"
    cISCOIETFPWTDMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWTDMMIB.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cISCOIETFPWTDMMIB.EntityData.SegmentPath = "CISCO-IETF-PW-TDM-MIB:CISCO-IETF-PW-TDM-MIB"
    cISCOIETFPWTDMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWTDMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWTDMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWTDMMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMObjects", types.YChild{"CpwCTDMObjects", &cISCOIETFPWTDMMIB.CpwCTDMObjects})
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMTable", types.YChild{"CpwCTDMTable", &cISCOIETFPWTDMMIB.CpwCTDMTable})
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMCfgTable", types.YChild{"CpwCTDMCfgTable", &cISCOIETFPWTDMMIB.CpwCTDMCfgTable})
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMPerfCurrentTable", types.YChild{"CpwCTDMPerfCurrentTable", &cISCOIETFPWTDMMIB.CpwCTDMPerfCurrentTable})
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMPerfIntervalTable", types.YChild{"CpwCTDMPerfIntervalTable", &cISCOIETFPWTDMMIB.CpwCTDMPerfIntervalTable})
    cISCOIETFPWTDMMIB.EntityData.Children.Append("cpwCTDMPerf1DayIntervalTable", types.YChild{"CpwCTDMPerf1DayIntervalTable", &cISCOIETFPWTDMMIB.CpwCTDMPerf1DayIntervalTable})
    cISCOIETFPWTDMMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFPWTDMMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFPWTDMMIB.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMObjects
type CISCOIETFPWTDMMIB_CpwCTDMObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the value to be used for cpwCTDMCfgIndex when creating
    // entries in the cpwCTDMCfgTable. The value 0 indicates that no unassigned
    // entries are available.  To obtain the value of cpwCTDMCfgIndexNext for a
    // new entry in the cpwCTDMCfgTable, the manager issues a management protocol
    // retrieval operation. The agent will determine through its local policy when
    // this index value will be made available for reuse. The type is interface{}
    // with range: 0..4294967295.
    CpwCTDMCfgIndexNext interface{}
}

func (cpwCTDMObjects *CISCOIETFPWTDMMIB_CpwCTDMObjects) GetEntityData() *types.CommonEntityData {
    cpwCTDMObjects.EntityData.YFilter = cpwCTDMObjects.YFilter
    cpwCTDMObjects.EntityData.YangName = "cpwCTDMObjects"
    cpwCTDMObjects.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMObjects.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMObjects.EntityData.SegmentPath = "cpwCTDMObjects"
    cpwCTDMObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMObjects.EntityData.Children = types.NewOrderedMap()
    cpwCTDMObjects.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMObjects.EntityData.Leafs.Append("cpwCTDMCfgIndexNext", types.YLeaf{"CpwCTDMCfgIndexNext", cpwCTDMObjects.CpwCTDMCfgIndexNext})

    cpwCTDMObjects.EntityData.YListKeys = []string {}

    return &(cpwCTDMObjects.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMTable
// This table contains basic information including ifIndex,
// and pointers to entries in the relevant TDM config
// tables for this TDM PW.
type CISCOIETFPWTDMMIB_CpwCTDMTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is indexed by the same index that was created for the associated
    // entry in the VC Table (in the CISCO-IETF-PW-MIB).    - The CpwVcIndex.  An
    // entry is created in this table by the agent for every entry in the
    // cpwVcTable with a cpwVcType equal to one of the following: e1Satop(12),
    // t1Satop(13), e3Satop(14), t3Satop(15), basicCesPsn(16), basicTdmIp(17), 
    // tdmCasCesPsn(18), tdmCasTdmIp(19). The type is slice of
    // CISCOIETFPWTDMMIB_CpwCTDMTable_CpwCTDMEntry.
    CpwCTDMEntry []*CISCOIETFPWTDMMIB_CpwCTDMTable_CpwCTDMEntry
}

func (cpwCTDMTable *CISCOIETFPWTDMMIB_CpwCTDMTable) GetEntityData() *types.CommonEntityData {
    cpwCTDMTable.EntityData.YFilter = cpwCTDMTable.YFilter
    cpwCTDMTable.EntityData.YangName = "cpwCTDMTable"
    cpwCTDMTable.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMTable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMTable.EntityData.SegmentPath = "cpwCTDMTable"
    cpwCTDMTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMTable.EntityData.Children = types.NewOrderedMap()
    cpwCTDMTable.EntityData.Children.Append("cpwCTDMEntry", types.YChild{"CpwCTDMEntry", nil})
    for i := range cpwCTDMTable.CpwCTDMEntry {
        cpwCTDMTable.EntityData.Children.Append(types.GetSegmentPath(cpwCTDMTable.CpwCTDMEntry[i]), types.YChild{"CpwCTDMEntry", cpwCTDMTable.CpwCTDMEntry[i]})
    }
    cpwCTDMTable.EntityData.Leafs = types.NewOrderedMap()

    cpwCTDMTable.EntityData.YListKeys = []string {}

    return &(cpwCTDMTable.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMTable_CpwCTDMEntry
// This table is indexed by the same index that was
// created for the associated entry in the VC Table
// (in the CISCO-IETF-PW-MIB).
// 
//   - The CpwVcIndex.
// 
// An entry is created in this table by the agent for every
// entry in the cpwVcTable with a cpwVcType equal to one of the
// following:
// e1Satop(12), t1Satop(13), e3Satop(14), t3Satop(15),
// basicCesPsn(16), basicTdmIp(17),  tdmCasCesPsn(18),
// tdmCasTdmIp(19).
type CISCOIETFPWTDMMIB_CpwCTDMTable_CpwCTDMEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // The parameter represents the bit-rate of the TDM service in multiples of
    // the 'basic' 64 Kbit/s rate. It complements the definition of cpwVcType used
    // in CISCO-IETF-PW-MIB. For structure-agnostic the following should be used:
    // a) Satop E1 - 32 b) Satop T1 emulation:    i)   MUST be set to 24 in the
    // basic emulation mode    ii)  MUST be set to 25 for the 'Octet-aligned T1'  
    // emulation mode c) Satop E3 - 535 d) Satop T3 - 699 For all kinds of
    // structure-aware emulation, this parameter MUST be set to N where N is the
    // number of DS0 channels in the corresponding attachment circuit. The type is
    // interface{} with range: -2147483648..2147483647.
    CpwCTDMRate interface{}

    // This is a unique index within the ifTable. It represents the interface
    // index of the full link or the interface index for the bundle holding the
    // group of time slots to be transmitted via this PW connection.  A value of
    // zero indicates an interface index that has yet to be determined. Once set,
    // if the TDM ifIndex is (for some reason) later removed, the agent SHOULD
    // delete the associated PW rows (e.g., this cpwCTDMTable entry). If the agent
    // does not delete the rows,  the agent MUST set this object to zero. The type
    // is interface{} with range: 0..2147483647.
    CpwCTDMIfIndex interface{}

    // Index to the generic parameters in the TDM configuration table that appears
    // in this MIB module. It is likely that multiple TDM PWs of the same
    // characteristic will share a single TDM Cfg entry. The type is interface{}
    // with range: 0..4294967295.
    CpwCGenTDMCfgIndex interface{}

    // Index to the relevant TDM configuration table entry that appears in one of
    // the related MIB modules such as TDMoIP or CESoPSN. It is likely that
    // multiple TDM PWs of the same characteristic will share a single
    // configuration entry of the relevant type. The value 0 implies no entry in
    // other related MIB. The type is interface{} with range: 0..4294967295.
    CpwCRelTDMCfgIndex interface{}

    // Any of the bits are set if the local configuration is not compatible with
    // the peer configuration as available from the various parameters options. 
    // -tdmTypeIncompatible bit is set if the local configuration is not carrying
    // the same TDM type as the peer configuration.  -peerRtpIncompatible bit is
    // set if the local configuration is configured to send RTP packets for this
    // PW, and the remote is not capable of accepting RTP packets. 
    // -peerPayloadSizeIncompatible bit is set if the local configuration is not
    // carrying the same Payload Size as the peer configuration. The type is
    // map[string]bool.
    CpwCTDMConfigError interface{}

    // The number of seconds, including partial seconds, that have elapsed since
    // the beginning of the current measurement period. If, for some reason, such
    // as an adjustment in the system's time-of-day clock, the current interval
    // exceeds the maximum value, the agent will return the maximum value. The
    // type is interface{} with range: 1..900. Units are seconds.
    CpwCTDMTimeElapsed interface{}

    // The number of previous 15-minute intervals for which data was collected. An
    // agent with TDM capability must be capable of supporting at least n
    // intervals. The minimum value of n is 4, The default of n is 32 and the
    // maximum value of n is 96. The value will be <n> unless the measurement was
    // (re-) started within the last (<n>*15) minutes, in which case the value
    // will be the number of complete 15 minute intervals for which the agent has
    // at least some data. In certain cases(e.g., in the case where the agent is a
    // proxy) it is possible that some intervals are unavailable. In this case,
    // this interval is the maximum interval number for which data is available.
    // The type is interface{} with range: 0..96. Units are minutes.
    CpwCTDMValidIntervals interface{}

    // The number of previous days for which data was collected. An agent with TDM
    // capability must be capable of supporting at least n intervals. The minimum
    // value of n is 1, The default of n is 1 and the maximum value of n is 30.
    // The type is interface{} with range: 0..30. Units are days.
    CpwCTDMValidDayIntervals interface{}

    // The following defects should be detected and reported upon request:  -Stray
    // packets MAY be detected by the PSN and multiplexing layers. Stray packets
    // MUST be discarded by the CE-bound IWF and their detection MUST NOT affect
    // mechanisms for detection of packet loss.  -Malformed packets are detected
    // by mismatch between the expected packet size (taking the value of the L bit
    // into account) and the actual packet size inferred from the PSN and
    // multiplexing layers. Malformed in-order packets MUST be discarded by the
    // CE-bound IWF and replacement data generated as for lost packets. 
    // -Excessive packet loss rate is detected by computing the average packet
    // loss rate over the value of cpwCTDMAvePktLossTimeWindow and comparing it
    // with a preconfigured threshold [SATOP].  -Buffer overrun is detected in the
    // normal operation state when the CE bound IWF's jitter buffer cannot
    // accommodate newly arrived packets.  -Remote packet loss is indicated by
    // reception of packets  with their R bit set.  -Packet misorder is detected
    // by looking at the Sequence number provided by the control word.  -TDM
    // Fault, if L bit in the control word is set, it indicates that TDM data
    // carried in the payload is invalid due an attachment circuit fault.  When
    // the L bit is set the payload MAY be omitted in order to conserve bandwidth.
    // Note: the algorithm used to capture these indications is implementation
    // specific. The type is map[string]bool.
    CpwCTDMCurrentIndications interface{}

    // The state of TDM indicators when the TDM PW last declared an error second
    // (either as ES, SES or a second with errors inside a UAS) condition. At this
    // time, only LOPS can create a failure. Since indicators other than LOPS are
    // useful, all are latched here. For bit definitions, see
    // cpwCTDMCurrentIndications above.  Note: the algorithm used to latch these
    // indications when entering a defect state is implementation specific. The
    // type is map[string]bool.
    CpwCTDMLatchedIndications interface{}

    // The value of sysUpTime at the most recent occasion at which the TDM PW
    // entered the ES or SES state. The type is interface{} with range:
    // 0..4294967295.
    CpwCTDMLastEsTimeStamp interface{}
}

func (cpwCTDMEntry *CISCOIETFPWTDMMIB_CpwCTDMTable_CpwCTDMEntry) GetEntityData() *types.CommonEntityData {
    cpwCTDMEntry.EntityData.YFilter = cpwCTDMEntry.YFilter
    cpwCTDMEntry.EntityData.YangName = "cpwCTDMEntry"
    cpwCTDMEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMEntry.EntityData.ParentYangName = "cpwCTDMTable"
    cpwCTDMEntry.EntityData.SegmentPath = "cpwCTDMEntry" + types.AddKeyToken(cpwCTDMEntry.CpwVcIndex, "cpwVcIndex")
    cpwCTDMEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMEntry.EntityData.Children = types.NewOrderedMap()
    cpwCTDMEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwCTDMEntry.CpwVcIndex})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMRate", types.YLeaf{"CpwCTDMRate", cpwCTDMEntry.CpwCTDMRate})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMIfIndex", types.YLeaf{"CpwCTDMIfIndex", cpwCTDMEntry.CpwCTDMIfIndex})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCGenTDMCfgIndex", types.YLeaf{"CpwCGenTDMCfgIndex", cpwCTDMEntry.CpwCGenTDMCfgIndex})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCRelTDMCfgIndex", types.YLeaf{"CpwCRelTDMCfgIndex", cpwCTDMEntry.CpwCRelTDMCfgIndex})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMConfigError", types.YLeaf{"CpwCTDMConfigError", cpwCTDMEntry.CpwCTDMConfigError})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMTimeElapsed", types.YLeaf{"CpwCTDMTimeElapsed", cpwCTDMEntry.CpwCTDMTimeElapsed})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMValidIntervals", types.YLeaf{"CpwCTDMValidIntervals", cpwCTDMEntry.CpwCTDMValidIntervals})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMValidDayIntervals", types.YLeaf{"CpwCTDMValidDayIntervals", cpwCTDMEntry.CpwCTDMValidDayIntervals})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMCurrentIndications", types.YLeaf{"CpwCTDMCurrentIndications", cpwCTDMEntry.CpwCTDMCurrentIndications})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMLatchedIndications", types.YLeaf{"CpwCTDMLatchedIndications", cpwCTDMEntry.CpwCTDMLatchedIndications})
    cpwCTDMEntry.EntityData.Leafs.Append("cpwCTDMLastEsTimeStamp", types.YLeaf{"CpwCTDMLastEsTimeStamp", cpwCTDMEntry.CpwCTDMLastEsTimeStamp})

    cpwCTDMEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwCTDMEntry.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMCfgTable
// This table contains a set of parameters that may be
// referenced by one or more TDM PWs in cpwCTDMTable.
type CISCOIETFPWTDMMIB_CpwCTDMCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // These parameters define the characteristics of a TDM PW. They are grouped
    // here to ease NMS burden. Once an entry is created here it may be re-used by
    // many PWs. The type is slice of
    // CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry.
    CpwCTDMCfgEntry []*CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry
}

func (cpwCTDMCfgTable *CISCOIETFPWTDMMIB_CpwCTDMCfgTable) GetEntityData() *types.CommonEntityData {
    cpwCTDMCfgTable.EntityData.YFilter = cpwCTDMCfgTable.YFilter
    cpwCTDMCfgTable.EntityData.YangName = "cpwCTDMCfgTable"
    cpwCTDMCfgTable.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMCfgTable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMCfgTable.EntityData.SegmentPath = "cpwCTDMCfgTable"
    cpwCTDMCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMCfgTable.EntityData.Children = types.NewOrderedMap()
    cpwCTDMCfgTable.EntityData.Children.Append("cpwCTDMCfgEntry", types.YChild{"CpwCTDMCfgEntry", nil})
    for i := range cpwCTDMCfgTable.CpwCTDMCfgEntry {
        cpwCTDMCfgTable.EntityData.Children.Append(types.GetSegmentPath(cpwCTDMCfgTable.CpwCTDMCfgEntry[i]), types.YChild{"CpwCTDMCfgEntry", cpwCTDMCfgTable.CpwCTDMCfgEntry[i]})
    }
    cpwCTDMCfgTable.EntityData.Leafs = types.NewOrderedMap()

    cpwCTDMCfgTable.EntityData.YListKeys = []string {}

    return &(cpwCTDMCfgTable.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry
// These parameters define the characteristics of a
// TDM PW. They are grouped here to ease NMS burden.
// Once an entry is created here it may be re-used
// by many PWs.
type CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index to an entry in this table. The value is a
    // copy of the assigned cpwCTDMCfgIndexNext. The type is interface{} with
    // range: 0..4294967295.
    CpwCTDMCfgIndex interface{}

    // This object indicates the various configuration errors, illegal settings
    // within the cpwCTDMCfg table. The errors can be due to several reasons, like
    // Payload size mismatch or Jitter Buffer depth value mistmatch.   payloadSize
    // - This bit is set if there is Payload size               mismatch between
    // the local and peer               configurations.  jtrBfrDepth - This bit is
    // set if there is Jitter Buffer               depth value mistmatch. other   
    // - This bit is set if the error is not due to               payloadSize and
    // jtrBfrDepth mismatch. The type is map[string]bool.
    CpwCTDMCfgConfErr interface{}

    // The value of this object indicates the PayLoad Size (in bytes) to be
    // defined during the PW setUp. Upon TX, implementation must be capable of
    // carrying that amount of bytes. Upon RX, when the LEN field is set to 0, the
    // payload of packet  MUST assume this size, and if the actual packet size is
    // inconsistent with this length, the packet MUST be considered to be
    // malformed. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    CpwCTDMCfgPayloadSize interface{}

    // If set to true, CE bound packets are queued in the jitter buffer, out of
    // order packets are re-ordered. The maximum sequence number differential
    // (i.e., the range in which re-sequencing can occur) is dependant on the
    // depth of the jitter buffer. See cpwCTDMCfgJtrBfrDepth. The type is bool.
    CpwCTDMCfgPktReorder interface{}

    // If the value of this object is set to false, then a RTP header is not
    // pre-pended to the TDM packet. The type is bool.
    CpwCTDMCfgRtpHdrUsed interface{}

    // The size of this buffer SHOULD be locally configured to allow accommodation
    // to the PSN-specific packet delay variation.  If configured to a value not
    // supported by the implementation, the agent MUST return an error code
    // 'jtrBfrDepth' in 'cpwCTDMConfigError '. Jitter buffers are a limited
    // resource to be managed. The actual size should be at least twice as big as
    // the value of cpwCTDMCfgJtrBfrDepth. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    CpwCTDMCfgJtrBfrDepth interface{}

    // This object indicates whether the Payload suppression is eanbled or
    // disabled. Payload MAY be omitted in order to conserve bandwidth.  enable  -
    // Payload suppression is allowed. disable - No Payload suppresion under any
    // condition. The type is CpwCTDMCfgPayloadSuppression.
    CpwCTDMCfgPayloadSuppression interface{}

    // The number of consecutive packets with sequential sequence numbers that are
    // required to exit the Loss of Packets (LOPS) state. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CpwCTDMCfgConsecPktsInSynch interface{}

    // The number of consecutive missing packets that are required to enter the
    // LOPS state. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CpwCTDMCfgConsecMissPktsOutSynch interface{}

    // The amount of time the host should wait before declaring the pseudo wire in
    // down state,  if the number of consecutive TDM packets that have been
    // received after changing the administrative status to up and after
    // finalization of signaling (if supported) between the two PEs is smaller
    // than cpwCTDMCfgConsecPktsInSynch. Once the the PW has OperStatus of 'up'
    // this parameter is no longer valid. This parameter is defined to ensure that
    // the host does not prematurely inform failure of the PW. In particular PW
    // 'down' notifications should not be sent before expiration of this timer.
    // This parameter is valid only after adminisrative changes of the status of
    // the PW. If the PW fails due to network impairments a 'down' notification
    // should be sent. The type is interface{} with range: 0..4294967295. Units
    // are millisecond.
    CpwCTDMCfgSetUp2SynchTimeOut interface{}

    // This parameter determines the value to be played when CE bound packets have
    // over/underflow the jitter buffer, or are missing for any reason. This AIS
    // (Alarm Indication Signal) pattern is sent (played) on the TDM line.  ais   
    // - AIS (Alarm Indication Signal)                          pattern is sent
    // (played) on                          the TDM line.  implementationSpecific
    // - Implementation specific pattern is                          sent on the
    // TDM line. The type is CpwCTDMCfgPktReplacePolicy.
    CpwCTDMCfgPktReplacePolicy interface{}

    // The length of time over which the average packet loss rate should be
    // computed to detect Excessive packet loss rate. The type is interface{} with
    // range: -2147483648..2147483647. Units are millisecond.
    CpwCTDMCfgAvePktLossTimeWindow interface{}

    // Excessive packet loss rate is detected by computing the average packetloss
    // rate over a  cpwCTDMCfgAvePktLossTimeWindow amount of time and comparing it
    // with this threshold value. The type is interface{} with range:
    // 0..4294967295.
    CpwCTDMCfgExcessivePktLossThreshold interface{}

    // Alarms are only reported when the defect state persists for the length of
    // time specified by this object. The object's unit is millisec. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    CpwCTDMCfgAlarmThreshold interface{}

    // Alarm MUST be cleared after the corresponding defect is undetected for the
    // amount of time specified by this object. The object's unit is millisec. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    CpwCTDMCfgClearAlarmThreshold interface{}

    // Number of missing packets detected (consecutive or not) within a 1 second
    // window to cause a Severely Error Second (SES) to be counted. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMCfgMissingPktsToSes interface{}

    // Timestamp generation MAY be used in one of the following modes: 1. Absolute
    // mode: the PSN-bound IWF sets timestamps  using the clock recovered from the
    // incoming TDM attachment  circuit. As a consequence, the timestamps are
    // closely  correlated with the sequence numbers. All TDM   implementations
    // that support usage of the RTP header MUST  support this mode. 2.
    // Differential mode: Both IWFs have access to a common  high-quality timing
    // source, and this source is used for  timestamp generation. Support of this
    // mode is OPTIONAL. The type is CpwCTDMCfgTimestampMode.
    CpwCTDMCfgTimestampMode interface{}

    // This variable indicates the storage type for this row. The following are
    // the read-write objects in permanent(4) rows, that an agent must allow to be
    // writable: cpwCTDMCfgPayloadSize, cpwCTDMCfgPktReorder,
    // cpwCTDMCfgRtpHdrUsed, cpwCTDMCfgJtrBfrDepth, cpwCTDMCfgPayloadSuppression,
    // cpwCTDMCfgConfErr. The type is StorageType.
    CpwCTDMCfgStorageType interface{}

    // The status of this conceptual row.  To create a row in this table, a
    // manager must set this object to either createAndGo(4) or createAndWait(5). 
    // All of the columnar objects have to be set to valid values before the row
    // can be activated. Default value will be automatically provisioned if for
    // those objects not specified during row creation.  No objects in cascading
    // tables have to be populated with related data before the row can be
    // activated.  The following objects cannot be modified if the RowStatus is
    // active: cpwCTDMCfgPayloadSize, cpwCTDMCfgRtpHdrUsed, cpwCTDMCfgJtrBfrDepth,
    // and cpwCTDMCfgPayloadSuppression.  If the RowStatus is active, the
    // following parameters can be modified:  cpwCTDMCfgConfErr,
    // cpwCTDMCfgPktReorder,  cpwCTDMCfgConsecPktsInSynch,
    // cpwCTDMCfgConsecMissPktsOutSynch, cpwCTDMCfgSetUp2SynchTimeOut,
    // cpwCTDMCfgPktReplacePolicy, cpwCTDMCfgAvePktLossTimeWindow, 
    // cpwCTDMCfgExcessivePktLossThreshold, cpwCTDMCfgAlarmThreshold,
    // cpwCTDMCfgClearAlarmThreshold, cpwCTDMCfgMissingPktsToSes,
    // cpwCTDMCfgTimestampMode, cpwCTDMCfgStorageType.  A row may be deleted by
    // setting the RowStatus to 'destroy'. The type is RowStatus.
    CpwCTDMCfgRowStatus interface{}
}

func (cpwCTDMCfgEntry *CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry) GetEntityData() *types.CommonEntityData {
    cpwCTDMCfgEntry.EntityData.YFilter = cpwCTDMCfgEntry.YFilter
    cpwCTDMCfgEntry.EntityData.YangName = "cpwCTDMCfgEntry"
    cpwCTDMCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMCfgEntry.EntityData.ParentYangName = "cpwCTDMCfgTable"
    cpwCTDMCfgEntry.EntityData.SegmentPath = "cpwCTDMCfgEntry" + types.AddKeyToken(cpwCTDMCfgEntry.CpwCTDMCfgIndex, "cpwCTDMCfgIndex")
    cpwCTDMCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMCfgEntry.EntityData.Children = types.NewOrderedMap()
    cpwCTDMCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgIndex", types.YLeaf{"CpwCTDMCfgIndex", cpwCTDMCfgEntry.CpwCTDMCfgIndex})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgConfErr", types.YLeaf{"CpwCTDMCfgConfErr", cpwCTDMCfgEntry.CpwCTDMCfgConfErr})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgPayloadSize", types.YLeaf{"CpwCTDMCfgPayloadSize", cpwCTDMCfgEntry.CpwCTDMCfgPayloadSize})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgPktReorder", types.YLeaf{"CpwCTDMCfgPktReorder", cpwCTDMCfgEntry.CpwCTDMCfgPktReorder})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgRtpHdrUsed", types.YLeaf{"CpwCTDMCfgRtpHdrUsed", cpwCTDMCfgEntry.CpwCTDMCfgRtpHdrUsed})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgJtrBfrDepth", types.YLeaf{"CpwCTDMCfgJtrBfrDepth", cpwCTDMCfgEntry.CpwCTDMCfgJtrBfrDepth})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgPayloadSuppression", types.YLeaf{"CpwCTDMCfgPayloadSuppression", cpwCTDMCfgEntry.CpwCTDMCfgPayloadSuppression})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgConsecPktsInSynch", types.YLeaf{"CpwCTDMCfgConsecPktsInSynch", cpwCTDMCfgEntry.CpwCTDMCfgConsecPktsInSynch})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgConsecMissPktsOutSynch", types.YLeaf{"CpwCTDMCfgConsecMissPktsOutSynch", cpwCTDMCfgEntry.CpwCTDMCfgConsecMissPktsOutSynch})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgSetUp2SynchTimeOut", types.YLeaf{"CpwCTDMCfgSetUp2SynchTimeOut", cpwCTDMCfgEntry.CpwCTDMCfgSetUp2SynchTimeOut})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgPktReplacePolicy", types.YLeaf{"CpwCTDMCfgPktReplacePolicy", cpwCTDMCfgEntry.CpwCTDMCfgPktReplacePolicy})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgAvePktLossTimeWindow", types.YLeaf{"CpwCTDMCfgAvePktLossTimeWindow", cpwCTDMCfgEntry.CpwCTDMCfgAvePktLossTimeWindow})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgExcessivePktLossThreshold", types.YLeaf{"CpwCTDMCfgExcessivePktLossThreshold", cpwCTDMCfgEntry.CpwCTDMCfgExcessivePktLossThreshold})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgAlarmThreshold", types.YLeaf{"CpwCTDMCfgAlarmThreshold", cpwCTDMCfgEntry.CpwCTDMCfgAlarmThreshold})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgClearAlarmThreshold", types.YLeaf{"CpwCTDMCfgClearAlarmThreshold", cpwCTDMCfgEntry.CpwCTDMCfgClearAlarmThreshold})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgMissingPktsToSes", types.YLeaf{"CpwCTDMCfgMissingPktsToSes", cpwCTDMCfgEntry.CpwCTDMCfgMissingPktsToSes})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgTimestampMode", types.YLeaf{"CpwCTDMCfgTimestampMode", cpwCTDMCfgEntry.CpwCTDMCfgTimestampMode})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgStorageType", types.YLeaf{"CpwCTDMCfgStorageType", cpwCTDMCfgEntry.CpwCTDMCfgStorageType})
    cpwCTDMCfgEntry.EntityData.Leafs.Append("cpwCTDMCfgRowStatus", types.YLeaf{"CpwCTDMCfgRowStatus", cpwCTDMCfgEntry.CpwCTDMCfgRowStatus})

    cpwCTDMCfgEntry.EntityData.YListKeys = []string {"CpwCTDMCfgIndex"}

    return &(cpwCTDMCfgEntry.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression represents disable - No Payload suppresion under any condition.
type CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression string

const (
    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression_enable CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression = "enable"

    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression_disable CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPayloadSuppression = "disable"
)

// CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy represents                          sent on the TDM line.
type CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy string

const (
    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy_ais CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy = "ais"

    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy_implementationSpecific CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgPktReplacePolicy = "implementationSpecific"
)

// CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode represents  timestamp generation. Support of this mode is OPTIONAL.
type CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode string

const (
    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode_notApplicable CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode = "notApplicable"

    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode_absolute CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode = "absolute"

    CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode_differential CISCOIETFPWTDMMIB_CpwCTDMCfgTable_CpwCTDMCfgEntry_CpwCTDMCfgTimestampMode = "differential"
)

// CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable
// This table provides TDM PW performance information.
// This includes current 15 minute interval counts. 
// 
// The table includes counters that work together to
// integrate errors and the lack of errors on the TDM PW.
// An error is caused by a missing packet. Missing packet
// can be a result of, packet loss in the network,
// (uncorrectable) packet out of sequence, packet length error,
// jitter buffer overflow, and jitter buffer underflow.
// The result is declaring whether or not the TDM PW is in
// Loss of Packet (LOPS) state.
type CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every cpwCTDMTable
    // entry. After 15 minutes, the contents of this table entry are copied to a
    // new entry in the cpwCTDMPerfInterval table and the counts in this entry are
    // reset to zero. The type is slice of
    // CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable_CpwCTDMPerfCurrentEntry.
    CpwCTDMPerfCurrentEntry []*CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable_CpwCTDMPerfCurrentEntry
}

func (cpwCTDMPerfCurrentTable *CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerfCurrentTable.EntityData.YFilter = cpwCTDMPerfCurrentTable.YFilter
    cpwCTDMPerfCurrentTable.EntityData.YangName = "cpwCTDMPerfCurrentTable"
    cpwCTDMPerfCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerfCurrentTable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMPerfCurrentTable.EntityData.SegmentPath = "cpwCTDMPerfCurrentTable"
    cpwCTDMPerfCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerfCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerfCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerfCurrentTable.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerfCurrentTable.EntityData.Children.Append("cpwCTDMPerfCurrentEntry", types.YChild{"CpwCTDMPerfCurrentEntry", nil})
    for i := range cpwCTDMPerfCurrentTable.CpwCTDMPerfCurrentEntry {
        cpwCTDMPerfCurrentTable.EntityData.Children.Append(types.GetSegmentPath(cpwCTDMPerfCurrentTable.CpwCTDMPerfCurrentEntry[i]), types.YChild{"CpwCTDMPerfCurrentEntry", cpwCTDMPerfCurrentTable.CpwCTDMPerfCurrentEntry[i]})
    }
    cpwCTDMPerfCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    cpwCTDMPerfCurrentTable.EntityData.YListKeys = []string {}

    return &(cpwCTDMPerfCurrentTable.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable_CpwCTDMPerfCurrentEntry
// An entry in this table is created by the agent for every
// cpwCTDMTable entry. After 15 minutes, the contents of this
// table entry are copied to a new entry in the
// cpwCTDMPerfInterval table and the counts in this entry
// are reset to zero.
type CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable_CpwCTDMPerfCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CpwCTDMPerfCurrentMissingPkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CpwCTDMPerfCurrentPktsReOrder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    CpwCTDMPerfCurrentJtrBfrUnderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerfCurrentMisOrderDropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerfCurrentMalformedPkt interface{}

    // The counter associated with the number of Error Seconds encountered. Any
    // malformed packet, sequence error and similar are considered as error
    // second. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CpwCTDMPerfCurrentESs interface{}

    // The counter associated with the number of Severely Error Seconds
    // encountered. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CpwCTDMPerfCurrentSESs interface{}

    // The counter associated with the number of Unavailable Seconds encountered.
    // Any consequtive five seconds of SES are counted as one UAS. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerfCurrentUASs interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. A failure event that begins in one period and ends in another
    // period is counted only in the period in which it begins. The type is
    // interface{} with range: 0..4294967295.
    CpwCTDMPerfCurrentFC interface{}
}

func (cpwCTDMPerfCurrentEntry *CISCOIETFPWTDMMIB_CpwCTDMPerfCurrentTable_CpwCTDMPerfCurrentEntry) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerfCurrentEntry.EntityData.YFilter = cpwCTDMPerfCurrentEntry.YFilter
    cpwCTDMPerfCurrentEntry.EntityData.YangName = "cpwCTDMPerfCurrentEntry"
    cpwCTDMPerfCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerfCurrentEntry.EntityData.ParentYangName = "cpwCTDMPerfCurrentTable"
    cpwCTDMPerfCurrentEntry.EntityData.SegmentPath = "cpwCTDMPerfCurrentEntry" + types.AddKeyToken(cpwCTDMPerfCurrentEntry.CpwVcIndex, "cpwVcIndex")
    cpwCTDMPerfCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerfCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerfCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerfCurrentEntry.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerfCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwCTDMPerfCurrentEntry.CpwVcIndex})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentMissingPkts", types.YLeaf{"CpwCTDMPerfCurrentMissingPkts", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentMissingPkts})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentPktsReOrder", types.YLeaf{"CpwCTDMPerfCurrentPktsReOrder", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentPktsReOrder})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentJtrBfrUnderruns", types.YLeaf{"CpwCTDMPerfCurrentJtrBfrUnderruns", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentJtrBfrUnderruns})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentMisOrderDropped", types.YLeaf{"CpwCTDMPerfCurrentMisOrderDropped", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentMisOrderDropped})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentMalformedPkt", types.YLeaf{"CpwCTDMPerfCurrentMalformedPkt", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentMalformedPkt})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentESs", types.YLeaf{"CpwCTDMPerfCurrentESs", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentESs})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentSESs", types.YLeaf{"CpwCTDMPerfCurrentSESs", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentSESs})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentUASs", types.YLeaf{"CpwCTDMPerfCurrentUASs", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentUASs})
    cpwCTDMPerfCurrentEntry.EntityData.Leafs.Append("cpwCTDMPerfCurrentFC", types.YLeaf{"CpwCTDMPerfCurrentFC", cpwCTDMPerfCurrentEntry.CpwCTDMPerfCurrentFC})

    cpwCTDMPerfCurrentEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwCTDMPerfCurrentEntry.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfCurrentTable above. However,
// these counts represent historical 15 minute intervals.
// Typically, this table will have a maximum of 96 entries
// for a 24 hour period, but is not limited to this.
type CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every
    // cpwCTDMPerfCurrentEntry that is 15 minutes old. The contents of the Current
    // entry are copied to the new entry here. The Current entry, then resets its
    // counts to zero for the next current 15 minute interval. The type is slice
    // of CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable_CpwCTDMPerfIntervalEntry.
    CpwCTDMPerfIntervalEntry []*CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable_CpwCTDMPerfIntervalEntry
}

func (cpwCTDMPerfIntervalTable *CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerfIntervalTable.EntityData.YFilter = cpwCTDMPerfIntervalTable.YFilter
    cpwCTDMPerfIntervalTable.EntityData.YangName = "cpwCTDMPerfIntervalTable"
    cpwCTDMPerfIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerfIntervalTable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMPerfIntervalTable.EntityData.SegmentPath = "cpwCTDMPerfIntervalTable"
    cpwCTDMPerfIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerfIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerfIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerfIntervalTable.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerfIntervalTable.EntityData.Children.Append("cpwCTDMPerfIntervalEntry", types.YChild{"CpwCTDMPerfIntervalEntry", nil})
    for i := range cpwCTDMPerfIntervalTable.CpwCTDMPerfIntervalEntry {
        cpwCTDMPerfIntervalTable.EntityData.Children.Append(types.GetSegmentPath(cpwCTDMPerfIntervalTable.CpwCTDMPerfIntervalEntry[i]), types.YChild{"CpwCTDMPerfIntervalEntry", cpwCTDMPerfIntervalTable.CpwCTDMPerfIntervalEntry[i]})
    }
    cpwCTDMPerfIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    cpwCTDMPerfIntervalTable.EntityData.YListKeys = []string {}

    return &(cpwCTDMPerfIntervalTable.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable_CpwCTDMPerfIntervalEntry
// An entry in this table is created by the agent for
// every cpwCTDMPerfCurrentEntry that is 15 minutes old.
// The contents of the Current entry are copied to the new
// entry here. The Current entry, then resets its counts
// to zero for the next current 15 minute interval.
type CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable_CpwCTDMPerfIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. This object indicates a number (normally between 1
    // and 96 to cover a 24 hour period) which identifies the interval for which
    // the set of statistics is available. The interval identified by 1 is the
    // most recently completed 15 minute interval, and the interval identified by
    // N is the interval immediately preceding the one identified by N-1. The
    // minimum range of N is 1 through 4.The default range is 1 through 32. The
    // maximum value of N is 1 through 96. The type is interface{} with range:
    // 0..4294967295.
    CpwCTDMPerfIntervalNumber interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    CpwCTDMPerfIntervalValidData interface{}

    // The duration of a particular interval in seconds. Adjustments in the
    // system's time-of-day clock, may cause the interval to be greater or less
    // than, the normal value. Therefore this actual interval value is provided.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerfIntervalDuration interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CpwCTDMPerfIntervalMissingPkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CpwCTDMPerfIntervalPktsReOrder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    CpwCTDMPerfIntervalJtrBfrUnderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerfIntervalMisOrderDropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerfIntervalMalformedPkt interface{}

    // The counter associated with the number of Error Seconds encountered. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerfIntervalESs interface{}

    // The counter associated with the number of Severely Error Seconds
    // encountered. The type is interface{} with range: 0..4294967295.
    CpwCTDMPerfIntervalSESs interface{}

    // The counter associated with the number of Unavailable Seconds encountered.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerfIntervalUASs interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. A failure event that begins in one period and ends in another
    // period is counted only in the period in which it begins. The type is
    // interface{} with range: 0..4294967295.
    CpwCTDMPerfIntervalFC interface{}
}

func (cpwCTDMPerfIntervalEntry *CISCOIETFPWTDMMIB_CpwCTDMPerfIntervalTable_CpwCTDMPerfIntervalEntry) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerfIntervalEntry.EntityData.YFilter = cpwCTDMPerfIntervalEntry.YFilter
    cpwCTDMPerfIntervalEntry.EntityData.YangName = "cpwCTDMPerfIntervalEntry"
    cpwCTDMPerfIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerfIntervalEntry.EntityData.ParentYangName = "cpwCTDMPerfIntervalTable"
    cpwCTDMPerfIntervalEntry.EntityData.SegmentPath = "cpwCTDMPerfIntervalEntry" + types.AddKeyToken(cpwCTDMPerfIntervalEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalNumber, "cpwCTDMPerfIntervalNumber")
    cpwCTDMPerfIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerfIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerfIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerfIntervalEntry.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerfIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwCTDMPerfIntervalEntry.CpwVcIndex})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalNumber", types.YLeaf{"CpwCTDMPerfIntervalNumber", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalNumber})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalValidData", types.YLeaf{"CpwCTDMPerfIntervalValidData", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalValidData})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalDuration", types.YLeaf{"CpwCTDMPerfIntervalDuration", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalDuration})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalMissingPkts", types.YLeaf{"CpwCTDMPerfIntervalMissingPkts", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalMissingPkts})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalPktsReOrder", types.YLeaf{"CpwCTDMPerfIntervalPktsReOrder", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalPktsReOrder})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalJtrBfrUnderruns", types.YLeaf{"CpwCTDMPerfIntervalJtrBfrUnderruns", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalJtrBfrUnderruns})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalMisOrderDropped", types.YLeaf{"CpwCTDMPerfIntervalMisOrderDropped", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalMisOrderDropped})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalMalformedPkt", types.YLeaf{"CpwCTDMPerfIntervalMalformedPkt", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalMalformedPkt})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalESs", types.YLeaf{"CpwCTDMPerfIntervalESs", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalESs})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalSESs", types.YLeaf{"CpwCTDMPerfIntervalSESs", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalSESs})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalUASs", types.YLeaf{"CpwCTDMPerfIntervalUASs", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalUASs})
    cpwCTDMPerfIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerfIntervalFC", types.YLeaf{"CpwCTDMPerfIntervalFC", cpwCTDMPerfIntervalEntry.CpwCTDMPerfIntervalFC})

    cpwCTDMPerfIntervalEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwCTDMPerfIntervalNumber"}

    return &(cpwCTDMPerfIntervalEntry.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfIntervalTable above. However,
// these counters represent historical 1 day intervals up to
// one full month. The table consists of real time data, as
// such it is not persistence across re-boot.
type CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in this table by the agent for every entry in the
    // cpwCTDMTable table. The type is slice of
    // CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable_CpwCTDMPerf1DayIntervalEntry.
    CpwCTDMPerf1DayIntervalEntry []*CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable_CpwCTDMPerf1DayIntervalEntry
}

func (cpwCTDMPerf1DayIntervalTable *CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerf1DayIntervalTable.EntityData.YFilter = cpwCTDMPerf1DayIntervalTable.YFilter
    cpwCTDMPerf1DayIntervalTable.EntityData.YangName = "cpwCTDMPerf1DayIntervalTable"
    cpwCTDMPerf1DayIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerf1DayIntervalTable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwCTDMPerf1DayIntervalTable.EntityData.SegmentPath = "cpwCTDMPerf1DayIntervalTable"
    cpwCTDMPerf1DayIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerf1DayIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerf1DayIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerf1DayIntervalTable.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerf1DayIntervalTable.EntityData.Children.Append("cpwCTDMPerf1DayIntervalEntry", types.YChild{"CpwCTDMPerf1DayIntervalEntry", nil})
    for i := range cpwCTDMPerf1DayIntervalTable.CpwCTDMPerf1DayIntervalEntry {
        cpwCTDMPerf1DayIntervalTable.EntityData.Children.Append(types.GetSegmentPath(cpwCTDMPerf1DayIntervalTable.CpwCTDMPerf1DayIntervalEntry[i]), types.YChild{"CpwCTDMPerf1DayIntervalEntry", cpwCTDMPerf1DayIntervalTable.CpwCTDMPerf1DayIntervalEntry[i]})
    }
    cpwCTDMPerf1DayIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    cpwCTDMPerf1DayIntervalTable.EntityData.YListKeys = []string {}

    return &(cpwCTDMPerf1DayIntervalTable.EntityData)
}

// CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable_CpwCTDMPerf1DayIntervalEntry
// An entry is created in this table by the agent
// for every entry in the cpwCTDMTable table.
type CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable_CpwCTDMPerf1DayIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. The number of interval, where 1 indicates current
    // day measured period and 2 and above indicate previous days respectively.
    // The type is interface{} with range: 0..4294967295.
    CpwCTDMPerf1DayIntervalNumber interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    CpwCTDMPerf1DayIntervalValidData interface{}

    // The duration of a particular interval in seconds, Adjustments in the
    // system's time-of-day clock, may cause the interval to be greater or less
    // than, the normal value. Therefore this actual interval value is provided.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerf1DayIntervalDuration interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    CpwCTDMPerf1DayIntervalMissingPkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CpwCTDMPerf1DayIntervalPktsReOrder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    CpwCTDMPerf1DayIntervalJtrBfrUnderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerf1DayIntervalMisOrderDropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    CpwCTDMPerf1DayIntervalMalformedPkt interface{}

    // The counter associated with the number of Error Seconds encountered. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerf1DayIntervalESs interface{}

    // The counter associated with the number of Severely Error Seconds. The type
    // is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerf1DayIntervalSESs interface{}

    // The counter associated with the number of UnAvailable Seconds. When first
    // entering the UAS state, the number of SES To UAS is added to this object,
    // then as each additional UAS occurs, this object increments by one. The type
    // is interface{} with range: 0..4294967295. Units are seconds.
    CpwCTDMPerf1DayIntervalUASs interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. The type is interface{} with range: 0..4294967295.
    CpwCTDMPerf1DayIntervalFC interface{}
}

func (cpwCTDMPerf1DayIntervalEntry *CISCOIETFPWTDMMIB_CpwCTDMPerf1DayIntervalTable_CpwCTDMPerf1DayIntervalEntry) GetEntityData() *types.CommonEntityData {
    cpwCTDMPerf1DayIntervalEntry.EntityData.YFilter = cpwCTDMPerf1DayIntervalEntry.YFilter
    cpwCTDMPerf1DayIntervalEntry.EntityData.YangName = "cpwCTDMPerf1DayIntervalEntry"
    cpwCTDMPerf1DayIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwCTDMPerf1DayIntervalEntry.EntityData.ParentYangName = "cpwCTDMPerf1DayIntervalTable"
    cpwCTDMPerf1DayIntervalEntry.EntityData.SegmentPath = "cpwCTDMPerf1DayIntervalEntry" + types.AddKeyToken(cpwCTDMPerf1DayIntervalEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalNumber, "cpwCTDMPerf1DayIntervalNumber")
    cpwCTDMPerf1DayIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwCTDMPerf1DayIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwCTDMPerf1DayIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwCTDMPerf1DayIntervalEntry.EntityData.Children = types.NewOrderedMap()
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwCTDMPerf1DayIntervalEntry.CpwVcIndex})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalNumber", types.YLeaf{"CpwCTDMPerf1DayIntervalNumber", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalNumber})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalValidData", types.YLeaf{"CpwCTDMPerf1DayIntervalValidData", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalValidData})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalDuration", types.YLeaf{"CpwCTDMPerf1DayIntervalDuration", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalDuration})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalMissingPkts", types.YLeaf{"CpwCTDMPerf1DayIntervalMissingPkts", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalMissingPkts})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalPktsReOrder", types.YLeaf{"CpwCTDMPerf1DayIntervalPktsReOrder", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalPktsReOrder})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalJtrBfrUnderruns", types.YLeaf{"CpwCTDMPerf1DayIntervalJtrBfrUnderruns", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalJtrBfrUnderruns})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalMisOrderDropped", types.YLeaf{"CpwCTDMPerf1DayIntervalMisOrderDropped", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalMisOrderDropped})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalMalformedPkt", types.YLeaf{"CpwCTDMPerf1DayIntervalMalformedPkt", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalMalformedPkt})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalESs", types.YLeaf{"CpwCTDMPerf1DayIntervalESs", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalESs})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalSESs", types.YLeaf{"CpwCTDMPerf1DayIntervalSESs", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalSESs})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalUASs", types.YLeaf{"CpwCTDMPerf1DayIntervalUASs", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalUASs})
    cpwCTDMPerf1DayIntervalEntry.EntityData.Leafs.Append("cpwCTDMPerf1DayIntervalFC", types.YLeaf{"CpwCTDMPerf1DayIntervalFC", cpwCTDMPerf1DayIntervalEntry.CpwCTDMPerf1DayIntervalFC})

    cpwCTDMPerf1DayIntervalEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwCTDMPerf1DayIntervalNumber"}

    return &(cpwCTDMPerf1DayIntervalEntry.EntityData)
}

