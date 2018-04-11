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

    
    Cpwctdmobjects CISCOIETFPWTDMMIB_Cpwctdmobjects

    // This table contains basic information including ifIndex, and pointers to
    // entries in the relevant TDM config tables for this TDM PW.
    Cpwctdmtable CISCOIETFPWTDMMIB_Cpwctdmtable

    // This table contains a set of parameters that may be referenced by one or
    // more TDM PWs in cpwCTDMTable.
    Cpwctdmcfgtable CISCOIETFPWTDMMIB_Cpwctdmcfgtable

    // This table provides TDM PW performance information. This includes current
    // 15 minute interval counts.   The table includes counters that work together
    // to integrate errors and the lack of errors on the TDM PW. An error is
    // caused by a missing packet. Missing packet can be a result of, packet loss
    // in the network, (uncorrectable) packet out of sequence, packet length
    // error, jitter buffer overflow, and jitter buffer underflow. The result is
    // declaring whether or not the TDM PW is in Loss of Packet (LOPS) state.
    Cpwctdmperfcurrenttable CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable

    // This table provides performance information per TDM PW similar to the
    // cpwCTDMPerfCurrentTable above. However, these counts represent historical
    // 15 minute intervals. Typically, this table will have a maximum of 96
    // entries for a 24 hour period, but is not limited to this.
    Cpwctdmperfintervaltable CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable

    // This table provides performance information per TDM PW similar to the
    // cpwCTDMPerfIntervalTable above. However, these counters represent
    // historical 1 day intervals up to one full month. The table consists of real
    // time data, as such it is not persistence across re-boot.
    Cpwctdmperf1Dayintervaltable CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable
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

    cISCOIETFPWTDMMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMObjects"] = types.YChild{"Cpwctdmobjects", &cISCOIETFPWTDMMIB.Cpwctdmobjects}
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMTable"] = types.YChild{"Cpwctdmtable", &cISCOIETFPWTDMMIB.Cpwctdmtable}
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMCfgTable"] = types.YChild{"Cpwctdmcfgtable", &cISCOIETFPWTDMMIB.Cpwctdmcfgtable}
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMPerfCurrentTable"] = types.YChild{"Cpwctdmperfcurrenttable", &cISCOIETFPWTDMMIB.Cpwctdmperfcurrenttable}
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMPerfIntervalTable"] = types.YChild{"Cpwctdmperfintervaltable", &cISCOIETFPWTDMMIB.Cpwctdmperfintervaltable}
    cISCOIETFPWTDMMIB.EntityData.Children["cpwCTDMPerf1DayIntervalTable"] = types.YChild{"Cpwctdmperf1Dayintervaltable", &cISCOIETFPWTDMMIB.Cpwctdmperf1Dayintervaltable}
    cISCOIETFPWTDMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFPWTDMMIB.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmobjects
type CISCOIETFPWTDMMIB_Cpwctdmobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains the value to be used for cpwCTDMCfgIndex when creating
    // entries in the cpwCTDMCfgTable. The value 0 indicates that no unassigned
    // entries are available.  To obtain the value of cpwCTDMCfgIndexNext for a
    // new entry in the cpwCTDMCfgTable, the manager issues a management protocol
    // retrieval operation. The agent will determine through its local policy when
    // this index value will be made available for reuse. The type is interface{}
    // with range: 0..4294967295.
    Cpwctdmcfgindexnext interface{}
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetEntityData() *types.CommonEntityData {
    cpwctdmobjects.EntityData.YFilter = cpwctdmobjects.YFilter
    cpwctdmobjects.EntityData.YangName = "cpwCTDMObjects"
    cpwctdmobjects.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmobjects.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmobjects.EntityData.SegmentPath = "cpwCTDMObjects"
    cpwctdmobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmobjects.EntityData.Children = make(map[string]types.YChild)
    cpwctdmobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmobjects.EntityData.Leafs["cpwCTDMCfgIndexNext"] = types.YLeaf{"Cpwctdmcfgindexnext", cpwctdmobjects.Cpwctdmcfgindexnext}
    return &(cpwctdmobjects.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmtable
// This table contains basic information including ifIndex,
// and pointers to entries in the relevant TDM config
// tables for this TDM PW.
type CISCOIETFPWTDMMIB_Cpwctdmtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is indexed by the same index that was created for the associated
    // entry in the VC Table (in the CISCO-IETF-PW-MIB).    - The CpwVcIndex.  An
    // entry is created in this table by the agent for every entry in the
    // cpwVcTable with a cpwVcType equal to one of the following: e1Satop(12),
    // t1Satop(13), e3Satop(14), t3Satop(15), basicCesPsn(16), basicTdmIp(17), 
    // tdmCasCesPsn(18), tdmCasTdmIp(19). The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry.
    Cpwctdmentry []CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetEntityData() *types.CommonEntityData {
    cpwctdmtable.EntityData.YFilter = cpwctdmtable.YFilter
    cpwctdmtable.EntityData.YangName = "cpwCTDMTable"
    cpwctdmtable.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmtable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmtable.EntityData.SegmentPath = "cpwCTDMTable"
    cpwctdmtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmtable.EntityData.Children = make(map[string]types.YChild)
    cpwctdmtable.EntityData.Children["cpwCTDMEntry"] = types.YChild{"Cpwctdmentry", nil}
    for i := range cpwctdmtable.Cpwctdmentry {
        cpwctdmtable.EntityData.Children[types.GetSegmentPath(&cpwctdmtable.Cpwctdmentry[i])] = types.YChild{"Cpwctdmentry", &cpwctdmtable.Cpwctdmentry[i]}
    }
    cpwctdmtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwctdmtable.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry
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
type CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // The parameter represents the bit-rate of the TDM service in multiples of
    // the 'basic' 64 Kbit/s rate. It complements the definition of cpwVcType used
    // in CISCO-IETF-PW-MIB. For structure-agnostic the following should be used:
    // a) Satop E1 - 32 b) Satop T1 emulation:    i)   MUST be set to 24 in the
    // basic emulation mode    ii)  MUST be set to 25 for the 'Octet-aligned T1'  
    // emulation mode c) Satop E3 - 535 d) Satop T3 - 699 For all kinds of
    // structure-aware emulation, this parameter MUST be set to N where N is the
    // number of DS0 channels in the corresponding attachment circuit. The type is
    // interface{} with range: -2147483648..2147483647.
    Cpwctdmrate interface{}

    // This is a unique index within the ifTable. It represents the interface
    // index of the full link or the interface index for the bundle holding the
    // group of time slots to be transmitted via this PW connection.  A value of
    // zero indicates an interface index that has yet to be determined. Once set,
    // if the TDM ifIndex is (for some reason) later removed, the agent SHOULD
    // delete the associated PW rows (e.g., this cpwCTDMTable entry). If the agent
    // does not delete the rows,  the agent MUST set this object to zero. The type
    // is interface{} with range: 0..2147483647.
    Cpwctdmifindex interface{}

    // Index to the generic parameters in the TDM configuration table that appears
    // in this MIB module. It is likely that multiple TDM PWs of the same
    // characteristic will share a single TDM Cfg entry. The type is interface{}
    // with range: 0..4294967295.
    Cpwcgentdmcfgindex interface{}

    // Index to the relevant TDM configuration table entry that appears in one of
    // the related MIB modules such as TDMoIP or CESoPSN. It is likely that
    // multiple TDM PWs of the same characteristic will share a single
    // configuration entry of the relevant type. The value 0 implies no entry in
    // other related MIB. The type is interface{} with range: 0..4294967295.
    Cpwcreltdmcfgindex interface{}

    // Any of the bits are set if the local configuration is not compatible with
    // the peer configuration as available from the various parameters options. 
    // -tdmTypeIncompatible bit is set if the local configuration is not carrying
    // the same TDM type as the peer configuration.  -peerRtpIncompatible bit is
    // set if the local configuration is configured to send RTP packets for this
    // PW, and the remote is not capable of accepting RTP packets. 
    // -peerPayloadSizeIncompatible bit is set if the local configuration is not
    // carrying the same Payload Size as the peer configuration. The type is
    // map[string]bool.
    Cpwctdmconfigerror interface{}

    // The number of seconds, including partial seconds, that have elapsed since
    // the beginning of the current measurement period. If, for some reason, such
    // as an adjustment in the system's time-of-day clock, the current interval
    // exceeds the maximum value, the agent will return the maximum value. The
    // type is interface{} with range: 1..900. Units are seconds.
    Cpwctdmtimeelapsed interface{}

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
    Cpwctdmvalidintervals interface{}

    // The number of previous days for which data was collected. An agent with TDM
    // capability must be capable of supporting at least n intervals. The minimum
    // value of n is 1, The default of n is 1 and the maximum value of n is 30.
    // The type is interface{} with range: 0..30. Units are days.
    Cpwctdmvaliddayintervals interface{}

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
    Cpwctdmcurrentindications interface{}

    // The state of TDM indicators when the TDM PW last declared an error second
    // (either as ES, SES or a second with errors inside a UAS) condition. At this
    // time, only LOPS can create a failure. Since indicators other than LOPS are
    // useful, all are latched here. For bit definitions, see
    // cpwCTDMCurrentIndications above.  Note: the algorithm used to latch these
    // indications when entering a defect state is implementation specific. The
    // type is map[string]bool.
    Cpwctdmlatchedindications interface{}

    // The value of sysUpTime at the most recent occasion at which the TDM PW
    // entered the ES or SES state. The type is interface{} with range:
    // 0..4294967295.
    Cpwctdmlastestimestamp interface{}
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetEntityData() *types.CommonEntityData {
    cpwctdmentry.EntityData.YFilter = cpwctdmentry.YFilter
    cpwctdmentry.EntityData.YangName = "cpwCTDMEntry"
    cpwctdmentry.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmentry.EntityData.ParentYangName = "cpwCTDMTable"
    cpwctdmentry.EntityData.SegmentPath = "cpwCTDMEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmentry.Cpwvcindex) + "']"
    cpwctdmentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmentry.EntityData.Children = make(map[string]types.YChild)
    cpwctdmentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwctdmentry.Cpwvcindex}
    cpwctdmentry.EntityData.Leafs["cpwCTDMRate"] = types.YLeaf{"Cpwctdmrate", cpwctdmentry.Cpwctdmrate}
    cpwctdmentry.EntityData.Leafs["cpwCTDMIfIndex"] = types.YLeaf{"Cpwctdmifindex", cpwctdmentry.Cpwctdmifindex}
    cpwctdmentry.EntityData.Leafs["cpwCGenTDMCfgIndex"] = types.YLeaf{"Cpwcgentdmcfgindex", cpwctdmentry.Cpwcgentdmcfgindex}
    cpwctdmentry.EntityData.Leafs["cpwCRelTDMCfgIndex"] = types.YLeaf{"Cpwcreltdmcfgindex", cpwctdmentry.Cpwcreltdmcfgindex}
    cpwctdmentry.EntityData.Leafs["cpwCTDMConfigError"] = types.YLeaf{"Cpwctdmconfigerror", cpwctdmentry.Cpwctdmconfigerror}
    cpwctdmentry.EntityData.Leafs["cpwCTDMTimeElapsed"] = types.YLeaf{"Cpwctdmtimeelapsed", cpwctdmentry.Cpwctdmtimeelapsed}
    cpwctdmentry.EntityData.Leafs["cpwCTDMValidIntervals"] = types.YLeaf{"Cpwctdmvalidintervals", cpwctdmentry.Cpwctdmvalidintervals}
    cpwctdmentry.EntityData.Leafs["cpwCTDMValidDayIntervals"] = types.YLeaf{"Cpwctdmvaliddayintervals", cpwctdmentry.Cpwctdmvaliddayintervals}
    cpwctdmentry.EntityData.Leafs["cpwCTDMCurrentIndications"] = types.YLeaf{"Cpwctdmcurrentindications", cpwctdmentry.Cpwctdmcurrentindications}
    cpwctdmentry.EntityData.Leafs["cpwCTDMLatchedIndications"] = types.YLeaf{"Cpwctdmlatchedindications", cpwctdmentry.Cpwctdmlatchedindications}
    cpwctdmentry.EntityData.Leafs["cpwCTDMLastEsTimeStamp"] = types.YLeaf{"Cpwctdmlastestimestamp", cpwctdmentry.Cpwctdmlastestimestamp}
    return &(cpwctdmentry.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable
// This table contains a set of parameters that may be
// referenced by one or more TDM PWs in cpwCTDMTable.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // These parameters define the characteristics of a TDM PW. They are grouped
    // here to ease NMS burden. Once an entry is created here it may be re-used by
    // many PWs. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry.
    Cpwctdmcfgentry []CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetEntityData() *types.CommonEntityData {
    cpwctdmcfgtable.EntityData.YFilter = cpwctdmcfgtable.YFilter
    cpwctdmcfgtable.EntityData.YangName = "cpwCTDMCfgTable"
    cpwctdmcfgtable.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmcfgtable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmcfgtable.EntityData.SegmentPath = "cpwCTDMCfgTable"
    cpwctdmcfgtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmcfgtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmcfgtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmcfgtable.EntityData.Children = make(map[string]types.YChild)
    cpwctdmcfgtable.EntityData.Children["cpwCTDMCfgEntry"] = types.YChild{"Cpwctdmcfgentry", nil}
    for i := range cpwctdmcfgtable.Cpwctdmcfgentry {
        cpwctdmcfgtable.EntityData.Children[types.GetSegmentPath(&cpwctdmcfgtable.Cpwctdmcfgentry[i])] = types.YChild{"Cpwctdmcfgentry", &cpwctdmcfgtable.Cpwctdmcfgentry[i]}
    }
    cpwctdmcfgtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwctdmcfgtable.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry
// These parameters define the characteristics of a
// TDM PW. They are grouped here to ease NMS burden.
// Once an entry is created here it may be re-used
// by many PWs.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index to an entry in this table. The value is a
    // copy of the assigned cpwCTDMCfgIndexNext. The type is interface{} with
    // range: 0..4294967295.
    Cpwctdmcfgindex interface{}

    // This object indicates the various configuration errors, illegal settings
    // within the cpwCTDMCfg table. The errors can be due to several reasons, like
    // Payload size mismatch or Jitter Buffer depth value mistmatch.   payloadSize
    // - This bit is set if there is Payload size               mismatch between
    // the local and peer               configurations.  jtrBfrDepth - This bit is
    // set if there is Jitter Buffer               depth value mistmatch. other   
    // - This bit is set if the error is not due to               payloadSize and
    // jtrBfrDepth mismatch. The type is map[string]bool.
    Cpwctdmcfgconferr interface{}

    // The value of this object indicates the PayLoad Size (in bytes) to be
    // defined during the PW setUp. Upon TX, implementation must be capable of
    // carrying that amount of bytes. Upon RX, when the LEN field is set to 0, the
    // payload of packet  MUST assume this size, and if the actual packet size is
    // inconsistent with this length, the packet MUST be considered to be
    // malformed. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    Cpwctdmcfgpayloadsize interface{}

    // If set to true, CE bound packets are queued in the jitter buffer, out of
    // order packets are re-ordered. The maximum sequence number differential
    // (i.e., the range in which re-sequencing can occur) is dependant on the
    // depth of the jitter buffer. See cpwCTDMCfgJtrBfrDepth. The type is bool.
    Cpwctdmcfgpktreorder interface{}

    // If the value of this object is set to false, then a RTP header is not
    // pre-pended to the TDM packet. The type is bool.
    Cpwctdmcfgrtphdrused interface{}

    // The size of this buffer SHOULD be locally configured to allow accommodation
    // to the PSN-specific packet delay variation.  If configured to a value not
    // supported by the implementation, the agent MUST return an error code
    // 'jtrBfrDepth' in 'cpwCTDMConfigError '. Jitter buffers are a limited
    // resource to be managed. The actual size should be at least twice as big as
    // the value of cpwCTDMCfgJtrBfrDepth. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    Cpwctdmcfgjtrbfrdepth interface{}

    // This object indicates whether the Payload suppression is eanbled or
    // disabled. Payload MAY be omitted in order to conserve bandwidth.  enable  -
    // Payload suppression is allowed. disable - No Payload suppresion under any
    // condition. The type is Cpwctdmcfgpayloadsuppression.
    Cpwctdmcfgpayloadsuppression interface{}

    // The number of consecutive packets with sequential sequence numbers that are
    // required to exit the Loss of Packets (LOPS) state. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Cpwctdmcfgconsecpktsinsynch interface{}

    // The number of consecutive missing packets that are required to enter the
    // LOPS state. The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cpwctdmcfgconsecmisspktsoutsynch interface{}

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
    Cpwctdmcfgsetup2Synchtimeout interface{}

    // This parameter determines the value to be played when CE bound packets have
    // over/underflow the jitter buffer, or are missing for any reason. This AIS
    // (Alarm Indication Signal) pattern is sent (played) on the TDM line.  ais   
    // - AIS (Alarm Indication Signal)                          pattern is sent
    // (played) on                          the TDM line.  implementationSpecific
    // - Implementation specific pattern is                          sent on the
    // TDM line. The type is Cpwctdmcfgpktreplacepolicy.
    Cpwctdmcfgpktreplacepolicy interface{}

    // The length of time over which the average packet loss rate should be
    // computed to detect Excessive packet loss rate. The type is interface{} with
    // range: -2147483648..2147483647. Units are millisecond.
    Cpwctdmcfgavepktlosstimewindow interface{}

    // Excessive packet loss rate is detected by computing the average packetloss
    // rate over a  cpwCTDMCfgAvePktLossTimeWindow amount of time and comparing it
    // with this threshold value. The type is interface{} with range:
    // 0..4294967295.
    Cpwctdmcfgexcessivepktlossthreshold interface{}

    // Alarms are only reported when the defect state persists for the length of
    // time specified by this object. The object's unit is millisec. The type is
    // interface{} with range: 0..4294967295. Units are milliseconds.
    Cpwctdmcfgalarmthreshold interface{}

    // Alarm MUST be cleared after the corresponding defect is undetected for the
    // amount of time specified by this object. The object's unit is millisec. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds.
    Cpwctdmcfgclearalarmthreshold interface{}

    // Number of missing packets detected (consecutive or not) within a 1 second
    // window to cause a Severely Error Second (SES) to be counted. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmcfgmissingpktstoses interface{}

    // Timestamp generation MAY be used in one of the following modes: 1. Absolute
    // mode: the PSN-bound IWF sets timestamps  using the clock recovered from the
    // incoming TDM attachment  circuit. As a consequence, the timestamps are
    // closely  correlated with the sequence numbers. All TDM   implementations
    // that support usage of the RTP header MUST  support this mode. 2.
    // Differential mode: Both IWFs have access to a common  high-quality timing
    // source, and this source is used for  timestamp generation. Support of this
    // mode is OPTIONAL. The type is Cpwctdmcfgtimestampmode.
    Cpwctdmcfgtimestampmode interface{}

    // This variable indicates the storage type for this row. The following are
    // the read-write objects in permanent(4) rows, that an agent must allow to be
    // writable: cpwCTDMCfgPayloadSize, cpwCTDMCfgPktReorder,
    // cpwCTDMCfgRtpHdrUsed, cpwCTDMCfgJtrBfrDepth, cpwCTDMCfgPayloadSuppression,
    // cpwCTDMCfgConfErr. The type is StorageType.
    Cpwctdmcfgstoragetype interface{}

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
    Cpwctdmcfgrowstatus interface{}
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetEntityData() *types.CommonEntityData {
    cpwctdmcfgentry.EntityData.YFilter = cpwctdmcfgentry.YFilter
    cpwctdmcfgentry.EntityData.YangName = "cpwCTDMCfgEntry"
    cpwctdmcfgentry.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmcfgentry.EntityData.ParentYangName = "cpwCTDMCfgTable"
    cpwctdmcfgentry.EntityData.SegmentPath = "cpwCTDMCfgEntry" + "[cpwCTDMCfgIndex='" + fmt.Sprintf("%v", cpwctdmcfgentry.Cpwctdmcfgindex) + "']"
    cpwctdmcfgentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmcfgentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmcfgentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmcfgentry.EntityData.Children = make(map[string]types.YChild)
    cpwctdmcfgentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgIndex"] = types.YLeaf{"Cpwctdmcfgindex", cpwctdmcfgentry.Cpwctdmcfgindex}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgConfErr"] = types.YLeaf{"Cpwctdmcfgconferr", cpwctdmcfgentry.Cpwctdmcfgconferr}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgPayloadSize"] = types.YLeaf{"Cpwctdmcfgpayloadsize", cpwctdmcfgentry.Cpwctdmcfgpayloadsize}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgPktReorder"] = types.YLeaf{"Cpwctdmcfgpktreorder", cpwctdmcfgentry.Cpwctdmcfgpktreorder}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgRtpHdrUsed"] = types.YLeaf{"Cpwctdmcfgrtphdrused", cpwctdmcfgentry.Cpwctdmcfgrtphdrused}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgJtrBfrDepth"] = types.YLeaf{"Cpwctdmcfgjtrbfrdepth", cpwctdmcfgentry.Cpwctdmcfgjtrbfrdepth}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgPayloadSuppression"] = types.YLeaf{"Cpwctdmcfgpayloadsuppression", cpwctdmcfgentry.Cpwctdmcfgpayloadsuppression}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgConsecPktsInSynch"] = types.YLeaf{"Cpwctdmcfgconsecpktsinsynch", cpwctdmcfgentry.Cpwctdmcfgconsecpktsinsynch}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgConsecMissPktsOutSynch"] = types.YLeaf{"Cpwctdmcfgconsecmisspktsoutsynch", cpwctdmcfgentry.Cpwctdmcfgconsecmisspktsoutsynch}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgSetUp2SynchTimeOut"] = types.YLeaf{"Cpwctdmcfgsetup2Synchtimeout", cpwctdmcfgentry.Cpwctdmcfgsetup2Synchtimeout}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgPktReplacePolicy"] = types.YLeaf{"Cpwctdmcfgpktreplacepolicy", cpwctdmcfgentry.Cpwctdmcfgpktreplacepolicy}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgAvePktLossTimeWindow"] = types.YLeaf{"Cpwctdmcfgavepktlosstimewindow", cpwctdmcfgentry.Cpwctdmcfgavepktlosstimewindow}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgExcessivePktLossThreshold"] = types.YLeaf{"Cpwctdmcfgexcessivepktlossthreshold", cpwctdmcfgentry.Cpwctdmcfgexcessivepktlossthreshold}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgAlarmThreshold"] = types.YLeaf{"Cpwctdmcfgalarmthreshold", cpwctdmcfgentry.Cpwctdmcfgalarmthreshold}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgClearAlarmThreshold"] = types.YLeaf{"Cpwctdmcfgclearalarmthreshold", cpwctdmcfgentry.Cpwctdmcfgclearalarmthreshold}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgMissingPktsToSes"] = types.YLeaf{"Cpwctdmcfgmissingpktstoses", cpwctdmcfgentry.Cpwctdmcfgmissingpktstoses}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgTimestampMode"] = types.YLeaf{"Cpwctdmcfgtimestampmode", cpwctdmcfgentry.Cpwctdmcfgtimestampmode}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgStorageType"] = types.YLeaf{"Cpwctdmcfgstoragetype", cpwctdmcfgentry.Cpwctdmcfgstoragetype}
    cpwctdmcfgentry.EntityData.Leafs["cpwCTDMCfgRowStatus"] = types.YLeaf{"Cpwctdmcfgrowstatus", cpwctdmcfgentry.Cpwctdmcfgrowstatus}
    return &(cpwctdmcfgentry.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression represents disable - No Payload suppresion under any condition.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression string

const (
    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression_enable CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression = "enable"

    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression_disable CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpayloadsuppression = "disable"
)

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy represents                          sent on the TDM line.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy string

const (
    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy_ais CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy = "ais"

    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy_implementationSpecific CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgpktreplacepolicy = "implementationSpecific"
)

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode represents  timestamp generation. Support of this mode is OPTIONAL.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode string

const (
    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode_notApplicable CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode = "notApplicable"

    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode_absolute CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode = "absolute"

    CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode_differential CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry_Cpwctdmcfgtimestampmode = "differential"
)

// CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable
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
type CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every cpwCTDMTable
    // entry. After 15 minutes, the contents of this table entry are copied to a
    // new entry in the cpwCTDMPerfInterval table and the counts in this entry are
    // reset to zero. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry.
    Cpwctdmperfcurrententry []CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetEntityData() *types.CommonEntityData {
    cpwctdmperfcurrenttable.EntityData.YFilter = cpwctdmperfcurrenttable.YFilter
    cpwctdmperfcurrenttable.EntityData.YangName = "cpwCTDMPerfCurrentTable"
    cpwctdmperfcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperfcurrenttable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmperfcurrenttable.EntityData.SegmentPath = "cpwCTDMPerfCurrentTable"
    cpwctdmperfcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperfcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperfcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperfcurrenttable.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperfcurrenttable.EntityData.Children["cpwCTDMPerfCurrentEntry"] = types.YChild{"Cpwctdmperfcurrententry", nil}
    for i := range cpwctdmperfcurrenttable.Cpwctdmperfcurrententry {
        cpwctdmperfcurrenttable.EntityData.Children[types.GetSegmentPath(&cpwctdmperfcurrenttable.Cpwctdmperfcurrententry[i])] = types.YChild{"Cpwctdmperfcurrententry", &cpwctdmperfcurrenttable.Cpwctdmperfcurrententry[i]}
    }
    cpwctdmperfcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwctdmperfcurrenttable.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry
// An entry in this table is created by the agent for every
// cpwCTDMTable entry. After 15 minutes, the contents of this
// table entry are copied to a new entry in the
// cpwCTDMPerfInterval table and the counts in this entry
// are reset to zero.
type CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cpwctdmperfcurrentmissingpkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cpwctdmperfcurrentpktsreorder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    Cpwctdmperfcurrentjtrbfrunderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperfcurrentmisorderdropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperfcurrentmalformedpkt interface{}

    // The counter associated with the number of Error Seconds encountered. Any
    // malformed packet, sequence error and similar are considered as error
    // second. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Cpwctdmperfcurrentess interface{}

    // The counter associated with the number of Severely Error Seconds
    // encountered. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Cpwctdmperfcurrentsess interface{}

    // The counter associated with the number of Unavailable Seconds encountered.
    // Any consequtive five seconds of SES are counted as one UAS. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperfcurrentuass interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. A failure event that begins in one period and ends in another
    // period is counted only in the period in which it begins. The type is
    // interface{} with range: 0..4294967295.
    Cpwctdmperfcurrentfc interface{}
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetEntityData() *types.CommonEntityData {
    cpwctdmperfcurrententry.EntityData.YFilter = cpwctdmperfcurrententry.YFilter
    cpwctdmperfcurrententry.EntityData.YangName = "cpwCTDMPerfCurrentEntry"
    cpwctdmperfcurrententry.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperfcurrententry.EntityData.ParentYangName = "cpwCTDMPerfCurrentTable"
    cpwctdmperfcurrententry.EntityData.SegmentPath = "cpwCTDMPerfCurrentEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperfcurrententry.Cpwvcindex) + "']"
    cpwctdmperfcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperfcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperfcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperfcurrententry.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperfcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmperfcurrententry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwctdmperfcurrententry.Cpwvcindex}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentMissingPkts"] = types.YLeaf{"Cpwctdmperfcurrentmissingpkts", cpwctdmperfcurrententry.Cpwctdmperfcurrentmissingpkts}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentPktsReOrder"] = types.YLeaf{"Cpwctdmperfcurrentpktsreorder", cpwctdmperfcurrententry.Cpwctdmperfcurrentpktsreorder}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentJtrBfrUnderruns"] = types.YLeaf{"Cpwctdmperfcurrentjtrbfrunderruns", cpwctdmperfcurrententry.Cpwctdmperfcurrentjtrbfrunderruns}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentMisOrderDropped"] = types.YLeaf{"Cpwctdmperfcurrentmisorderdropped", cpwctdmperfcurrententry.Cpwctdmperfcurrentmisorderdropped}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentMalformedPkt"] = types.YLeaf{"Cpwctdmperfcurrentmalformedpkt", cpwctdmperfcurrententry.Cpwctdmperfcurrentmalformedpkt}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentESs"] = types.YLeaf{"Cpwctdmperfcurrentess", cpwctdmperfcurrententry.Cpwctdmperfcurrentess}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentSESs"] = types.YLeaf{"Cpwctdmperfcurrentsess", cpwctdmperfcurrententry.Cpwctdmperfcurrentsess}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentUASs"] = types.YLeaf{"Cpwctdmperfcurrentuass", cpwctdmperfcurrententry.Cpwctdmperfcurrentuass}
    cpwctdmperfcurrententry.EntityData.Leafs["cpwCTDMPerfCurrentFC"] = types.YLeaf{"Cpwctdmperfcurrentfc", cpwctdmperfcurrententry.Cpwctdmperfcurrentfc}
    return &(cpwctdmperfcurrententry.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfCurrentTable above. However,
// these counts represent historical 15 minute intervals.
// Typically, this table will have a maximum of 96 entries
// for a 24 hour period, but is not limited to this.
type CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every
    // cpwCTDMPerfCurrentEntry that is 15 minutes old. The contents of the Current
    // entry are copied to the new entry here. The Current entry, then resets its
    // counts to zero for the next current 15 minute interval. The type is slice
    // of CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry.
    Cpwctdmperfintervalentry []CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetEntityData() *types.CommonEntityData {
    cpwctdmperfintervaltable.EntityData.YFilter = cpwctdmperfintervaltable.YFilter
    cpwctdmperfintervaltable.EntityData.YangName = "cpwCTDMPerfIntervalTable"
    cpwctdmperfintervaltable.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperfintervaltable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmperfintervaltable.EntityData.SegmentPath = "cpwCTDMPerfIntervalTable"
    cpwctdmperfintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperfintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperfintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperfintervaltable.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperfintervaltable.EntityData.Children["cpwCTDMPerfIntervalEntry"] = types.YChild{"Cpwctdmperfintervalentry", nil}
    for i := range cpwctdmperfintervaltable.Cpwctdmperfintervalentry {
        cpwctdmperfintervaltable.EntityData.Children[types.GetSegmentPath(&cpwctdmperfintervaltable.Cpwctdmperfintervalentry[i])] = types.YChild{"Cpwctdmperfintervalentry", &cpwctdmperfintervaltable.Cpwctdmperfintervalentry[i]}
    }
    cpwctdmperfintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwctdmperfintervaltable.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry
// An entry in this table is created by the agent for
// every cpwCTDMPerfCurrentEntry that is 15 minutes old.
// The contents of the Current entry are copied to the new
// entry here. The Current entry, then resets its counts
// to zero for the next current 15 minute interval.
type CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. This object indicates a number (normally between 1
    // and 96 to cover a 24 hour period) which identifies the interval for which
    // the set of statistics is available. The interval identified by 1 is the
    // most recently completed 15 minute interval, and the interval identified by
    // N is the interval immediately preceding the one identified by N-1. The
    // minimum range of N is 1 through 4.The default range is 1 through 32. The
    // maximum value of N is 1 through 96. The type is interface{} with range:
    // 0..4294967295.
    Cpwctdmperfintervalnumber interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Cpwctdmperfintervalvaliddata interface{}

    // The duration of a particular interval in seconds. Adjustments in the
    // system's time-of-day clock, may cause the interval to be greater or less
    // than, the normal value. Therefore this actual interval value is provided.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperfintervalduration interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cpwctdmperfintervalmissingpkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cpwctdmperfintervalpktsreorder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    Cpwctdmperfintervaljtrbfrunderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperfintervalmisorderdropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperfintervalmalformedpkt interface{}

    // The counter associated with the number of Error Seconds encountered. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperfintervaless interface{}

    // The counter associated with the number of Severely Error Seconds
    // encountered. The type is interface{} with range: 0..4294967295.
    Cpwctdmperfintervalsess interface{}

    // The counter associated with the number of Unavailable Seconds encountered.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperfintervaluass interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. A failure event that begins in one period and ends in another
    // period is counted only in the period in which it begins. The type is
    // interface{} with range: 0..4294967295.
    Cpwctdmperfintervalfc interface{}
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetEntityData() *types.CommonEntityData {
    cpwctdmperfintervalentry.EntityData.YFilter = cpwctdmperfintervalentry.YFilter
    cpwctdmperfintervalentry.EntityData.YangName = "cpwCTDMPerfIntervalEntry"
    cpwctdmperfintervalentry.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperfintervalentry.EntityData.ParentYangName = "cpwCTDMPerfIntervalTable"
    cpwctdmperfintervalentry.EntityData.SegmentPath = "cpwCTDMPerfIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperfintervalentry.Cpwvcindex) + "']" + "[cpwCTDMPerfIntervalNumber='" + fmt.Sprintf("%v", cpwctdmperfintervalentry.Cpwctdmperfintervalnumber) + "']"
    cpwctdmperfintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperfintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperfintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperfintervalentry.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperfintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmperfintervalentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwctdmperfintervalentry.Cpwvcindex}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalNumber"] = types.YLeaf{"Cpwctdmperfintervalnumber", cpwctdmperfintervalentry.Cpwctdmperfintervalnumber}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalValidData"] = types.YLeaf{"Cpwctdmperfintervalvaliddata", cpwctdmperfintervalentry.Cpwctdmperfintervalvaliddata}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalDuration"] = types.YLeaf{"Cpwctdmperfintervalduration", cpwctdmperfintervalentry.Cpwctdmperfintervalduration}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalMissingPkts"] = types.YLeaf{"Cpwctdmperfintervalmissingpkts", cpwctdmperfintervalentry.Cpwctdmperfintervalmissingpkts}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalPktsReOrder"] = types.YLeaf{"Cpwctdmperfintervalpktsreorder", cpwctdmperfintervalentry.Cpwctdmperfintervalpktsreorder}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalJtrBfrUnderruns"] = types.YLeaf{"Cpwctdmperfintervaljtrbfrunderruns", cpwctdmperfintervalentry.Cpwctdmperfintervaljtrbfrunderruns}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalMisOrderDropped"] = types.YLeaf{"Cpwctdmperfintervalmisorderdropped", cpwctdmperfintervalentry.Cpwctdmperfintervalmisorderdropped}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalMalformedPkt"] = types.YLeaf{"Cpwctdmperfintervalmalformedpkt", cpwctdmperfintervalentry.Cpwctdmperfintervalmalformedpkt}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalESs"] = types.YLeaf{"Cpwctdmperfintervaless", cpwctdmperfintervalentry.Cpwctdmperfintervaless}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalSESs"] = types.YLeaf{"Cpwctdmperfintervalsess", cpwctdmperfintervalentry.Cpwctdmperfintervalsess}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalUASs"] = types.YLeaf{"Cpwctdmperfintervaluass", cpwctdmperfintervalentry.Cpwctdmperfintervaluass}
    cpwctdmperfintervalentry.EntityData.Leafs["cpwCTDMPerfIntervalFC"] = types.YLeaf{"Cpwctdmperfintervalfc", cpwctdmperfintervalentry.Cpwctdmperfintervalfc}
    return &(cpwctdmperfintervalentry.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfIntervalTable above. However,
// these counters represent historical 1 day intervals up to
// one full month. The table consists of real time data, as
// such it is not persistence across re-boot.
type CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry is created in this table by the agent for every entry in the
    // cpwCTDMTable table. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry.
    Cpwctdmperf1Dayintervalentry []CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetEntityData() *types.CommonEntityData {
    cpwctdmperf1Dayintervaltable.EntityData.YFilter = cpwctdmperf1Dayintervaltable.YFilter
    cpwctdmperf1Dayintervaltable.EntityData.YangName = "cpwCTDMPerf1DayIntervalTable"
    cpwctdmperf1Dayintervaltable.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperf1Dayintervaltable.EntityData.ParentYangName = "CISCO-IETF-PW-TDM-MIB"
    cpwctdmperf1Dayintervaltable.EntityData.SegmentPath = "cpwCTDMPerf1DayIntervalTable"
    cpwctdmperf1Dayintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperf1Dayintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperf1Dayintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperf1Dayintervaltable.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperf1Dayintervaltable.EntityData.Children["cpwCTDMPerf1DayIntervalEntry"] = types.YChild{"Cpwctdmperf1Dayintervalentry", nil}
    for i := range cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry {
        cpwctdmperf1Dayintervaltable.EntityData.Children[types.GetSegmentPath(&cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry[i])] = types.YChild{"Cpwctdmperf1Dayintervalentry", &cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry[i]}
    }
    cpwctdmperf1Dayintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwctdmperf1Dayintervaltable.EntityData)
}

// CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry
// An entry is created in this table by the agent
// for every entry in the cpwCTDMTable table.
type CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. The number of interval, where 1 indicates current
    // day measured period and 2 and above indicate previous days respectively.
    // The type is interface{} with range: 0..4294967295.
    Cpwctdmperf1Dayintervalnumber interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Cpwctdmperf1Dayintervalvaliddata interface{}

    // The duration of a particular interval in seconds, Adjustments in the
    // system's time-of-day clock, may cause the interval to be greater or less
    // than, the normal value. Therefore this actual interval value is provided.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperf1Dayintervalduration interface{}

    // Number of missing packets (as detected via control word sequence number
    // gaps). The type is interface{} with range: 0..4294967295. Units are
    // packets.
    Cpwctdmperf1Dayintervalmissingpkts interface{}

    // Number of packets detected out of sequence (via control word sequence
    // number), but successfully re-ordered. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Cpwctdmperf1Dayintervalpktsreorder interface{}

    // Number of times a packet needed to be played out and the jitter buffer was
    // empty. The type is interface{} with range: 0..4294967295.
    Cpwctdmperf1Dayintervaljtrbfrunderruns interface{}

    // Number of packets detected out of order(via control word sequence numbers),
    // and could not be re-ordered, or could not fit in the jitter buffer. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperf1Dayintervalmisorderdropped interface{}

    // Number of packets detected with unexpected size, or bad headers' stack. The
    // type is interface{} with range: 0..4294967295. Units are packets.
    Cpwctdmperf1Dayintervalmalformedpkt interface{}

    // The counter associated with the number of Error Seconds encountered. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperf1Dayintervaless interface{}

    // The counter associated with the number of Severely Error Seconds. The type
    // is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperf1Dayintervalsess interface{}

    // The counter associated with the number of UnAvailable Seconds. When first
    // entering the UAS state, the number of SES To UAS is added to this object,
    // then as each additional UAS occurs, this object increments by one. The type
    // is interface{} with range: 0..4294967295. Units are seconds.
    Cpwctdmperf1Dayintervaluass interface{}

    // This object represents the number of TDM failure events. A failure event
    // begins when the LOPS failure is declared, and ends when the failure is
    // cleared. The type is interface{} with range: 0..4294967295.
    Cpwctdmperf1Dayintervalfc interface{}
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetEntityData() *types.CommonEntityData {
    cpwctdmperf1Dayintervalentry.EntityData.YFilter = cpwctdmperf1Dayintervalentry.YFilter
    cpwctdmperf1Dayintervalentry.EntityData.YangName = "cpwCTDMPerf1DayIntervalEntry"
    cpwctdmperf1Dayintervalentry.EntityData.BundleName = "cisco_ios_xe"
    cpwctdmperf1Dayintervalentry.EntityData.ParentYangName = "cpwCTDMPerf1DayIntervalTable"
    cpwctdmperf1Dayintervalentry.EntityData.SegmentPath = "cpwCTDMPerf1DayIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperf1Dayintervalentry.Cpwvcindex) + "']" + "[cpwCTDMPerf1DayIntervalNumber='" + fmt.Sprintf("%v", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalnumber) + "']"
    cpwctdmperf1Dayintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwctdmperf1Dayintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwctdmperf1Dayintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwctdmperf1Dayintervalentry.EntityData.Children = make(map[string]types.YChild)
    cpwctdmperf1Dayintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwctdmperf1Dayintervalentry.Cpwvcindex}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalNumber"] = types.YLeaf{"Cpwctdmperf1Dayintervalnumber", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalnumber}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalValidData"] = types.YLeaf{"Cpwctdmperf1Dayintervalvaliddata", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalvaliddata}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalDuration"] = types.YLeaf{"Cpwctdmperf1Dayintervalduration", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalduration}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalMissingPkts"] = types.YLeaf{"Cpwctdmperf1Dayintervalmissingpkts", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmissingpkts}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalPktsReOrder"] = types.YLeaf{"Cpwctdmperf1Dayintervalpktsreorder", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalpktsreorder}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalJtrBfrUnderruns"] = types.YLeaf{"Cpwctdmperf1Dayintervaljtrbfrunderruns", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaljtrbfrunderruns}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalMisOrderDropped"] = types.YLeaf{"Cpwctdmperf1Dayintervalmisorderdropped", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmisorderdropped}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalMalformedPkt"] = types.YLeaf{"Cpwctdmperf1Dayintervalmalformedpkt", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmalformedpkt}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalESs"] = types.YLeaf{"Cpwctdmperf1Dayintervaless", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaless}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalSESs"] = types.YLeaf{"Cpwctdmperf1Dayintervalsess", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalsess}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalUASs"] = types.YLeaf{"Cpwctdmperf1Dayintervaluass", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaluass}
    cpwctdmperf1Dayintervalentry.EntityData.Leafs["cpwCTDMPerf1DayIntervalFC"] = types.YLeaf{"Cpwctdmperf1Dayintervalfc", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalfc}
    return &(cpwctdmperf1Dayintervalentry.EntityData)
}

