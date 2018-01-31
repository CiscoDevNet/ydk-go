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
    parent types.Entity
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

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetFilter() yfilter.YFilter { return cISCOIETFPWTDMMIB.YFilter }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFPWTDMMIB.YFilter = yf }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetGoName(yname string) string {
    if yname == "cpwCTDMObjects" { return "Cpwctdmobjects" }
    if yname == "cpwCTDMTable" { return "Cpwctdmtable" }
    if yname == "cpwCTDMCfgTable" { return "Cpwctdmcfgtable" }
    if yname == "cpwCTDMPerfCurrentTable" { return "Cpwctdmperfcurrenttable" }
    if yname == "cpwCTDMPerfIntervalTable" { return "Cpwctdmperfintervaltable" }
    if yname == "cpwCTDMPerf1DayIntervalTable" { return "Cpwctdmperf1Dayintervaltable" }
    return ""
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetSegmentPath() string {
    return "CISCO-IETF-PW-TDM-MIB:CISCO-IETF-PW-TDM-MIB"
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMObjects" {
        return &cISCOIETFPWTDMMIB.Cpwctdmobjects
    }
    if childYangName == "cpwCTDMTable" {
        return &cISCOIETFPWTDMMIB.Cpwctdmtable
    }
    if childYangName == "cpwCTDMCfgTable" {
        return &cISCOIETFPWTDMMIB.Cpwctdmcfgtable
    }
    if childYangName == "cpwCTDMPerfCurrentTable" {
        return &cISCOIETFPWTDMMIB.Cpwctdmperfcurrenttable
    }
    if childYangName == "cpwCTDMPerfIntervalTable" {
        return &cISCOIETFPWTDMMIB.Cpwctdmperfintervaltable
    }
    if childYangName == "cpwCTDMPerf1DayIntervalTable" {
        return &cISCOIETFPWTDMMIB.Cpwctdmperf1Dayintervaltable
    }
    return nil
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpwCTDMObjects"] = &cISCOIETFPWTDMMIB.Cpwctdmobjects
    children["cpwCTDMTable"] = &cISCOIETFPWTDMMIB.Cpwctdmtable
    children["cpwCTDMCfgTable"] = &cISCOIETFPWTDMMIB.Cpwctdmcfgtable
    children["cpwCTDMPerfCurrentTable"] = &cISCOIETFPWTDMMIB.Cpwctdmperfcurrenttable
    children["cpwCTDMPerfIntervalTable"] = &cISCOIETFPWTDMMIB.Cpwctdmperfintervaltable
    children["cpwCTDMPerf1DayIntervalTable"] = &cISCOIETFPWTDMMIB.Cpwctdmperf1Dayintervaltable
    return children
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) SetParent(parent types.Entity) { cISCOIETFPWTDMMIB.parent = parent }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetParent() types.Entity { return cISCOIETFPWTDMMIB.parent }

func (cISCOIETFPWTDMMIB *CISCOIETFPWTDMMIB) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmobjects
type CISCOIETFPWTDMMIB_Cpwctdmobjects struct {
    parent types.Entity
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

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetFilter() yfilter.YFilter { return cpwctdmobjects.YFilter }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) SetFilter(yf yfilter.YFilter) { cpwctdmobjects.YFilter = yf }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetGoName(yname string) string {
    if yname == "cpwCTDMCfgIndexNext" { return "Cpwctdmcfgindexnext" }
    return ""
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetSegmentPath() string {
    return "cpwCTDMObjects"
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwCTDMCfgIndexNext"] = cpwctdmobjects.Cpwctdmcfgindexnext
    return leafs
}

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetYangName() string { return "cpwCTDMObjects" }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) SetParent(parent types.Entity) { cpwctdmobjects.parent = parent }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetParent() types.Entity { return cpwctdmobjects.parent }

func (cpwctdmobjects *CISCOIETFPWTDMMIB_Cpwctdmobjects) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmtable
// This table contains basic information including ifIndex,
// and pointers to entries in the relevant TDM config
// tables for this TDM PW.
type CISCOIETFPWTDMMIB_Cpwctdmtable struct {
    parent types.Entity
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

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetFilter() yfilter.YFilter { return cpwctdmtable.YFilter }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) SetFilter(yf yfilter.YFilter) { cpwctdmtable.YFilter = yf }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetGoName(yname string) string {
    if yname == "cpwCTDMEntry" { return "Cpwctdmentry" }
    return ""
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetSegmentPath() string {
    return "cpwCTDMTable"
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMEntry" {
        for _, c := range cpwctdmtable.Cpwctdmentry {
            if cpwctdmtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry{}
        cpwctdmtable.Cpwctdmentry = append(cpwctdmtable.Cpwctdmentry, child)
        return &cpwctdmtable.Cpwctdmentry[len(cpwctdmtable.Cpwctdmentry)-1]
    }
    return nil
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwctdmtable.Cpwctdmentry {
        children[cpwctdmtable.Cpwctdmentry[i].GetSegmentPath()] = &cpwctdmtable.Cpwctdmentry[i]
    }
    return children
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetYangName() string { return "cpwCTDMTable" }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) SetParent(parent types.Entity) { cpwctdmtable.parent = parent }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetParent() types.Entity { return cpwctdmtable.parent }

func (cpwctdmtable *CISCOIETFPWTDMMIB_Cpwctdmtable) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

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
    parent types.Entity
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

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetFilter() yfilter.YFilter { return cpwctdmentry.YFilter }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) SetFilter(yf yfilter.YFilter) { cpwctdmentry.YFilter = yf }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwCTDMRate" { return "Cpwctdmrate" }
    if yname == "cpwCTDMIfIndex" { return "Cpwctdmifindex" }
    if yname == "cpwCGenTDMCfgIndex" { return "Cpwcgentdmcfgindex" }
    if yname == "cpwCRelTDMCfgIndex" { return "Cpwcreltdmcfgindex" }
    if yname == "cpwCTDMConfigError" { return "Cpwctdmconfigerror" }
    if yname == "cpwCTDMTimeElapsed" { return "Cpwctdmtimeelapsed" }
    if yname == "cpwCTDMValidIntervals" { return "Cpwctdmvalidintervals" }
    if yname == "cpwCTDMValidDayIntervals" { return "Cpwctdmvaliddayintervals" }
    if yname == "cpwCTDMCurrentIndications" { return "Cpwctdmcurrentindications" }
    if yname == "cpwCTDMLatchedIndications" { return "Cpwctdmlatchedindications" }
    if yname == "cpwCTDMLastEsTimeStamp" { return "Cpwctdmlastestimestamp" }
    return ""
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetSegmentPath() string {
    return "cpwCTDMEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmentry.Cpwvcindex) + "']"
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwctdmentry.Cpwvcindex
    leafs["cpwCTDMRate"] = cpwctdmentry.Cpwctdmrate
    leafs["cpwCTDMIfIndex"] = cpwctdmentry.Cpwctdmifindex
    leafs["cpwCGenTDMCfgIndex"] = cpwctdmentry.Cpwcgentdmcfgindex
    leafs["cpwCRelTDMCfgIndex"] = cpwctdmentry.Cpwcreltdmcfgindex
    leafs["cpwCTDMConfigError"] = cpwctdmentry.Cpwctdmconfigerror
    leafs["cpwCTDMTimeElapsed"] = cpwctdmentry.Cpwctdmtimeelapsed
    leafs["cpwCTDMValidIntervals"] = cpwctdmentry.Cpwctdmvalidintervals
    leafs["cpwCTDMValidDayIntervals"] = cpwctdmentry.Cpwctdmvaliddayintervals
    leafs["cpwCTDMCurrentIndications"] = cpwctdmentry.Cpwctdmcurrentindications
    leafs["cpwCTDMLatchedIndications"] = cpwctdmentry.Cpwctdmlatchedindications
    leafs["cpwCTDMLastEsTimeStamp"] = cpwctdmentry.Cpwctdmlastestimestamp
    return leafs
}

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetYangName() string { return "cpwCTDMEntry" }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) SetParent(parent types.Entity) { cpwctdmentry.parent = parent }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetParent() types.Entity { return cpwctdmentry.parent }

func (cpwctdmentry *CISCOIETFPWTDMMIB_Cpwctdmtable_Cpwctdmentry) GetParentYangName() string { return "cpwCTDMTable" }

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable
// This table contains a set of parameters that may be
// referenced by one or more TDM PWs in cpwCTDMTable.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // These parameters define the characteristics of a TDM PW. They are grouped
    // here to ease NMS burden. Once an entry is created here it may be re-used by
    // many PWs. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry.
    Cpwctdmcfgentry []CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetFilter() yfilter.YFilter { return cpwctdmcfgtable.YFilter }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) SetFilter(yf yfilter.YFilter) { cpwctdmcfgtable.YFilter = yf }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetGoName(yname string) string {
    if yname == "cpwCTDMCfgEntry" { return "Cpwctdmcfgentry" }
    return ""
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetSegmentPath() string {
    return "cpwCTDMCfgTable"
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMCfgEntry" {
        for _, c := range cpwctdmcfgtable.Cpwctdmcfgentry {
            if cpwctdmcfgtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry{}
        cpwctdmcfgtable.Cpwctdmcfgentry = append(cpwctdmcfgtable.Cpwctdmcfgentry, child)
        return &cpwctdmcfgtable.Cpwctdmcfgentry[len(cpwctdmcfgtable.Cpwctdmcfgentry)-1]
    }
    return nil
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwctdmcfgtable.Cpwctdmcfgentry {
        children[cpwctdmcfgtable.Cpwctdmcfgentry[i].GetSegmentPath()] = &cpwctdmcfgtable.Cpwctdmcfgentry[i]
    }
    return children
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetYangName() string { return "cpwCTDMCfgTable" }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) SetParent(parent types.Entity) { cpwctdmcfgtable.parent = parent }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetParent() types.Entity { return cpwctdmcfgtable.parent }

func (cpwctdmcfgtable *CISCOIETFPWTDMMIB_Cpwctdmcfgtable) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry
// These parameters define the characteristics of a
// TDM PW. They are grouped here to ease NMS burden.
// Once an entry is created here it may be re-used
// by many PWs.
type CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry struct {
    parent types.Entity
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

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetFilter() yfilter.YFilter { return cpwctdmcfgentry.YFilter }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) SetFilter(yf yfilter.YFilter) { cpwctdmcfgentry.YFilter = yf }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetGoName(yname string) string {
    if yname == "cpwCTDMCfgIndex" { return "Cpwctdmcfgindex" }
    if yname == "cpwCTDMCfgConfErr" { return "Cpwctdmcfgconferr" }
    if yname == "cpwCTDMCfgPayloadSize" { return "Cpwctdmcfgpayloadsize" }
    if yname == "cpwCTDMCfgPktReorder" { return "Cpwctdmcfgpktreorder" }
    if yname == "cpwCTDMCfgRtpHdrUsed" { return "Cpwctdmcfgrtphdrused" }
    if yname == "cpwCTDMCfgJtrBfrDepth" { return "Cpwctdmcfgjtrbfrdepth" }
    if yname == "cpwCTDMCfgPayloadSuppression" { return "Cpwctdmcfgpayloadsuppression" }
    if yname == "cpwCTDMCfgConsecPktsInSynch" { return "Cpwctdmcfgconsecpktsinsynch" }
    if yname == "cpwCTDMCfgConsecMissPktsOutSynch" { return "Cpwctdmcfgconsecmisspktsoutsynch" }
    if yname == "cpwCTDMCfgSetUp2SynchTimeOut" { return "Cpwctdmcfgsetup2Synchtimeout" }
    if yname == "cpwCTDMCfgPktReplacePolicy" { return "Cpwctdmcfgpktreplacepolicy" }
    if yname == "cpwCTDMCfgAvePktLossTimeWindow" { return "Cpwctdmcfgavepktlosstimewindow" }
    if yname == "cpwCTDMCfgExcessivePktLossThreshold" { return "Cpwctdmcfgexcessivepktlossthreshold" }
    if yname == "cpwCTDMCfgAlarmThreshold" { return "Cpwctdmcfgalarmthreshold" }
    if yname == "cpwCTDMCfgClearAlarmThreshold" { return "Cpwctdmcfgclearalarmthreshold" }
    if yname == "cpwCTDMCfgMissingPktsToSes" { return "Cpwctdmcfgmissingpktstoses" }
    if yname == "cpwCTDMCfgTimestampMode" { return "Cpwctdmcfgtimestampmode" }
    if yname == "cpwCTDMCfgStorageType" { return "Cpwctdmcfgstoragetype" }
    if yname == "cpwCTDMCfgRowStatus" { return "Cpwctdmcfgrowstatus" }
    return ""
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetSegmentPath() string {
    return "cpwCTDMCfgEntry" + "[cpwCTDMCfgIndex='" + fmt.Sprintf("%v", cpwctdmcfgentry.Cpwctdmcfgindex) + "']"
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwCTDMCfgIndex"] = cpwctdmcfgentry.Cpwctdmcfgindex
    leafs["cpwCTDMCfgConfErr"] = cpwctdmcfgentry.Cpwctdmcfgconferr
    leafs["cpwCTDMCfgPayloadSize"] = cpwctdmcfgentry.Cpwctdmcfgpayloadsize
    leafs["cpwCTDMCfgPktReorder"] = cpwctdmcfgentry.Cpwctdmcfgpktreorder
    leafs["cpwCTDMCfgRtpHdrUsed"] = cpwctdmcfgentry.Cpwctdmcfgrtphdrused
    leafs["cpwCTDMCfgJtrBfrDepth"] = cpwctdmcfgentry.Cpwctdmcfgjtrbfrdepth
    leafs["cpwCTDMCfgPayloadSuppression"] = cpwctdmcfgentry.Cpwctdmcfgpayloadsuppression
    leafs["cpwCTDMCfgConsecPktsInSynch"] = cpwctdmcfgentry.Cpwctdmcfgconsecpktsinsynch
    leafs["cpwCTDMCfgConsecMissPktsOutSynch"] = cpwctdmcfgentry.Cpwctdmcfgconsecmisspktsoutsynch
    leafs["cpwCTDMCfgSetUp2SynchTimeOut"] = cpwctdmcfgentry.Cpwctdmcfgsetup2Synchtimeout
    leafs["cpwCTDMCfgPktReplacePolicy"] = cpwctdmcfgentry.Cpwctdmcfgpktreplacepolicy
    leafs["cpwCTDMCfgAvePktLossTimeWindow"] = cpwctdmcfgentry.Cpwctdmcfgavepktlosstimewindow
    leafs["cpwCTDMCfgExcessivePktLossThreshold"] = cpwctdmcfgentry.Cpwctdmcfgexcessivepktlossthreshold
    leafs["cpwCTDMCfgAlarmThreshold"] = cpwctdmcfgentry.Cpwctdmcfgalarmthreshold
    leafs["cpwCTDMCfgClearAlarmThreshold"] = cpwctdmcfgentry.Cpwctdmcfgclearalarmthreshold
    leafs["cpwCTDMCfgMissingPktsToSes"] = cpwctdmcfgentry.Cpwctdmcfgmissingpktstoses
    leafs["cpwCTDMCfgTimestampMode"] = cpwctdmcfgentry.Cpwctdmcfgtimestampmode
    leafs["cpwCTDMCfgStorageType"] = cpwctdmcfgentry.Cpwctdmcfgstoragetype
    leafs["cpwCTDMCfgRowStatus"] = cpwctdmcfgentry.Cpwctdmcfgrowstatus
    return leafs
}

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetYangName() string { return "cpwCTDMCfgEntry" }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) SetParent(parent types.Entity) { cpwctdmcfgentry.parent = parent }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetParent() types.Entity { return cpwctdmcfgentry.parent }

func (cpwctdmcfgentry *CISCOIETFPWTDMMIB_Cpwctdmcfgtable_Cpwctdmcfgentry) GetParentYangName() string { return "cpwCTDMCfgTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every cpwCTDMTable
    // entry. After 15 minutes, the contents of this table entry are copied to a
    // new entry in the cpwCTDMPerfInterval table and the counts in this entry are
    // reset to zero. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry.
    Cpwctdmperfcurrententry []CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetFilter() yfilter.YFilter { return cpwctdmperfcurrenttable.YFilter }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) SetFilter(yf yfilter.YFilter) { cpwctdmperfcurrenttable.YFilter = yf }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetGoName(yname string) string {
    if yname == "cpwCTDMPerfCurrentEntry" { return "Cpwctdmperfcurrententry" }
    return ""
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetSegmentPath() string {
    return "cpwCTDMPerfCurrentTable"
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMPerfCurrentEntry" {
        for _, c := range cpwctdmperfcurrenttable.Cpwctdmperfcurrententry {
            if cpwctdmperfcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry{}
        cpwctdmperfcurrenttable.Cpwctdmperfcurrententry = append(cpwctdmperfcurrenttable.Cpwctdmperfcurrententry, child)
        return &cpwctdmperfcurrenttable.Cpwctdmperfcurrententry[len(cpwctdmperfcurrenttable.Cpwctdmperfcurrententry)-1]
    }
    return nil
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwctdmperfcurrenttable.Cpwctdmperfcurrententry {
        children[cpwctdmperfcurrenttable.Cpwctdmperfcurrententry[i].GetSegmentPath()] = &cpwctdmperfcurrenttable.Cpwctdmperfcurrententry[i]
    }
    return children
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetYangName() string { return "cpwCTDMPerfCurrentTable" }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) SetParent(parent types.Entity) { cpwctdmperfcurrenttable.parent = parent }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetParent() types.Entity { return cpwctdmperfcurrenttable.parent }

func (cpwctdmperfcurrenttable *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry
// An entry in this table is created by the agent for every
// cpwCTDMTable entry. After 15 minutes, the contents of this
// table entry are copied to a new entry in the
// cpwCTDMPerfInterval table and the counts in this entry
// are reset to zero.
type CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry struct {
    parent types.Entity
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

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetFilter() yfilter.YFilter { return cpwctdmperfcurrententry.YFilter }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) SetFilter(yf yfilter.YFilter) { cpwctdmperfcurrententry.YFilter = yf }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwCTDMPerfCurrentMissingPkts" { return "Cpwctdmperfcurrentmissingpkts" }
    if yname == "cpwCTDMPerfCurrentPktsReOrder" { return "Cpwctdmperfcurrentpktsreorder" }
    if yname == "cpwCTDMPerfCurrentJtrBfrUnderruns" { return "Cpwctdmperfcurrentjtrbfrunderruns" }
    if yname == "cpwCTDMPerfCurrentMisOrderDropped" { return "Cpwctdmperfcurrentmisorderdropped" }
    if yname == "cpwCTDMPerfCurrentMalformedPkt" { return "Cpwctdmperfcurrentmalformedpkt" }
    if yname == "cpwCTDMPerfCurrentESs" { return "Cpwctdmperfcurrentess" }
    if yname == "cpwCTDMPerfCurrentSESs" { return "Cpwctdmperfcurrentsess" }
    if yname == "cpwCTDMPerfCurrentUASs" { return "Cpwctdmperfcurrentuass" }
    if yname == "cpwCTDMPerfCurrentFC" { return "Cpwctdmperfcurrentfc" }
    return ""
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetSegmentPath() string {
    return "cpwCTDMPerfCurrentEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperfcurrententry.Cpwvcindex) + "']"
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwctdmperfcurrententry.Cpwvcindex
    leafs["cpwCTDMPerfCurrentMissingPkts"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentmissingpkts
    leafs["cpwCTDMPerfCurrentPktsReOrder"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentpktsreorder
    leafs["cpwCTDMPerfCurrentJtrBfrUnderruns"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentjtrbfrunderruns
    leafs["cpwCTDMPerfCurrentMisOrderDropped"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentmisorderdropped
    leafs["cpwCTDMPerfCurrentMalformedPkt"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentmalformedpkt
    leafs["cpwCTDMPerfCurrentESs"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentess
    leafs["cpwCTDMPerfCurrentSESs"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentsess
    leafs["cpwCTDMPerfCurrentUASs"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentuass
    leafs["cpwCTDMPerfCurrentFC"] = cpwctdmperfcurrententry.Cpwctdmperfcurrentfc
    return leafs
}

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetYangName() string { return "cpwCTDMPerfCurrentEntry" }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) SetParent(parent types.Entity) { cpwctdmperfcurrententry.parent = parent }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetParent() types.Entity { return cpwctdmperfcurrententry.parent }

func (cpwctdmperfcurrententry *CISCOIETFPWTDMMIB_Cpwctdmperfcurrenttable_Cpwctdmperfcurrententry) GetParentYangName() string { return "cpwCTDMPerfCurrentTable" }

// CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfCurrentTable above. However,
// these counts represent historical 15 minute intervals.
// Typically, this table will have a maximum of 96 entries
// for a 24 hour period, but is not limited to this.
type CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every
    // cpwCTDMPerfCurrentEntry that is 15 minutes old. The contents of the Current
    // entry are copied to the new entry here. The Current entry, then resets its
    // counts to zero for the next current 15 minute interval. The type is slice
    // of CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry.
    Cpwctdmperfintervalentry []CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetFilter() yfilter.YFilter { return cpwctdmperfintervaltable.YFilter }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) SetFilter(yf yfilter.YFilter) { cpwctdmperfintervaltable.YFilter = yf }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetGoName(yname string) string {
    if yname == "cpwCTDMPerfIntervalEntry" { return "Cpwctdmperfintervalentry" }
    return ""
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetSegmentPath() string {
    return "cpwCTDMPerfIntervalTable"
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMPerfIntervalEntry" {
        for _, c := range cpwctdmperfintervaltable.Cpwctdmperfintervalentry {
            if cpwctdmperfintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry{}
        cpwctdmperfintervaltable.Cpwctdmperfintervalentry = append(cpwctdmperfintervaltable.Cpwctdmperfintervalentry, child)
        return &cpwctdmperfintervaltable.Cpwctdmperfintervalentry[len(cpwctdmperfintervaltable.Cpwctdmperfintervalentry)-1]
    }
    return nil
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwctdmperfintervaltable.Cpwctdmperfintervalentry {
        children[cpwctdmperfintervaltable.Cpwctdmperfintervalentry[i].GetSegmentPath()] = &cpwctdmperfintervaltable.Cpwctdmperfintervalentry[i]
    }
    return children
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetYangName() string { return "cpwCTDMPerfIntervalTable" }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) SetParent(parent types.Entity) { cpwctdmperfintervaltable.parent = parent }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetParent() types.Entity { return cpwctdmperfintervaltable.parent }

func (cpwctdmperfintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry
// An entry in this table is created by the agent for
// every cpwCTDMPerfCurrentEntry that is 15 minutes old.
// The contents of the Current entry are copied to the new
// entry here. The Current entry, then resets its counts
// to zero for the next current 15 minute interval.
type CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry struct {
    parent types.Entity
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

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetFilter() yfilter.YFilter { return cpwctdmperfintervalentry.YFilter }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) SetFilter(yf yfilter.YFilter) { cpwctdmperfintervalentry.YFilter = yf }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwCTDMPerfIntervalNumber" { return "Cpwctdmperfintervalnumber" }
    if yname == "cpwCTDMPerfIntervalValidData" { return "Cpwctdmperfintervalvaliddata" }
    if yname == "cpwCTDMPerfIntervalDuration" { return "Cpwctdmperfintervalduration" }
    if yname == "cpwCTDMPerfIntervalMissingPkts" { return "Cpwctdmperfintervalmissingpkts" }
    if yname == "cpwCTDMPerfIntervalPktsReOrder" { return "Cpwctdmperfintervalpktsreorder" }
    if yname == "cpwCTDMPerfIntervalJtrBfrUnderruns" { return "Cpwctdmperfintervaljtrbfrunderruns" }
    if yname == "cpwCTDMPerfIntervalMisOrderDropped" { return "Cpwctdmperfintervalmisorderdropped" }
    if yname == "cpwCTDMPerfIntervalMalformedPkt" { return "Cpwctdmperfintervalmalformedpkt" }
    if yname == "cpwCTDMPerfIntervalESs" { return "Cpwctdmperfintervaless" }
    if yname == "cpwCTDMPerfIntervalSESs" { return "Cpwctdmperfintervalsess" }
    if yname == "cpwCTDMPerfIntervalUASs" { return "Cpwctdmperfintervaluass" }
    if yname == "cpwCTDMPerfIntervalFC" { return "Cpwctdmperfintervalfc" }
    return ""
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetSegmentPath() string {
    return "cpwCTDMPerfIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperfintervalentry.Cpwvcindex) + "']" + "[cpwCTDMPerfIntervalNumber='" + fmt.Sprintf("%v", cpwctdmperfintervalentry.Cpwctdmperfintervalnumber) + "']"
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwctdmperfintervalentry.Cpwvcindex
    leafs["cpwCTDMPerfIntervalNumber"] = cpwctdmperfintervalentry.Cpwctdmperfintervalnumber
    leafs["cpwCTDMPerfIntervalValidData"] = cpwctdmperfintervalentry.Cpwctdmperfintervalvaliddata
    leafs["cpwCTDMPerfIntervalDuration"] = cpwctdmperfintervalentry.Cpwctdmperfintervalduration
    leafs["cpwCTDMPerfIntervalMissingPkts"] = cpwctdmperfintervalentry.Cpwctdmperfintervalmissingpkts
    leafs["cpwCTDMPerfIntervalPktsReOrder"] = cpwctdmperfintervalentry.Cpwctdmperfintervalpktsreorder
    leafs["cpwCTDMPerfIntervalJtrBfrUnderruns"] = cpwctdmperfintervalentry.Cpwctdmperfintervaljtrbfrunderruns
    leafs["cpwCTDMPerfIntervalMisOrderDropped"] = cpwctdmperfintervalentry.Cpwctdmperfintervalmisorderdropped
    leafs["cpwCTDMPerfIntervalMalformedPkt"] = cpwctdmperfintervalentry.Cpwctdmperfintervalmalformedpkt
    leafs["cpwCTDMPerfIntervalESs"] = cpwctdmperfintervalentry.Cpwctdmperfintervaless
    leafs["cpwCTDMPerfIntervalSESs"] = cpwctdmperfintervalentry.Cpwctdmperfintervalsess
    leafs["cpwCTDMPerfIntervalUASs"] = cpwctdmperfintervalentry.Cpwctdmperfintervaluass
    leafs["cpwCTDMPerfIntervalFC"] = cpwctdmperfintervalentry.Cpwctdmperfintervalfc
    return leafs
}

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetYangName() string { return "cpwCTDMPerfIntervalEntry" }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) SetParent(parent types.Entity) { cpwctdmperfintervalentry.parent = parent }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetParent() types.Entity { return cpwctdmperfintervalentry.parent }

func (cpwctdmperfintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperfintervaltable_Cpwctdmperfintervalentry) GetParentYangName() string { return "cpwCTDMPerfIntervalTable" }

// CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable
// This table provides performance information per TDM PW
// similar to the cpwCTDMPerfIntervalTable above. However,
// these counters represent historical 1 day intervals up to
// one full month. The table consists of real time data, as
// such it is not persistence across re-boot.
type CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry is created in this table by the agent for every entry in the
    // cpwCTDMTable table. The type is slice of
    // CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry.
    Cpwctdmperf1Dayintervalentry []CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetFilter() yfilter.YFilter { return cpwctdmperf1Dayintervaltable.YFilter }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) SetFilter(yf yfilter.YFilter) { cpwctdmperf1Dayintervaltable.YFilter = yf }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetGoName(yname string) string {
    if yname == "cpwCTDMPerf1DayIntervalEntry" { return "Cpwctdmperf1Dayintervalentry" }
    return ""
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetSegmentPath() string {
    return "cpwCTDMPerf1DayIntervalTable"
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwCTDMPerf1DayIntervalEntry" {
        for _, c := range cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry {
            if cpwctdmperf1Dayintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry{}
        cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry = append(cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry, child)
        return &cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry[len(cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry)-1]
    }
    return nil
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry {
        children[cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry[i].GetSegmentPath()] = &cpwctdmperf1Dayintervaltable.Cpwctdmperf1Dayintervalentry[i]
    }
    return children
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetYangName() string { return "cpwCTDMPerf1DayIntervalTable" }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) SetParent(parent types.Entity) { cpwctdmperf1Dayintervaltable.parent = parent }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetParent() types.Entity { return cpwctdmperf1Dayintervaltable.parent }

func (cpwctdmperf1Dayintervaltable *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable) GetParentYangName() string { return "CISCO-IETF-PW-TDM-MIB" }

// CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry
// An entry is created in this table by the agent
// for every entry in the cpwCTDMTable table.
type CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry struct {
    parent types.Entity
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

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetFilter() yfilter.YFilter { return cpwctdmperf1Dayintervalentry.YFilter }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) SetFilter(yf yfilter.YFilter) { cpwctdmperf1Dayintervalentry.YFilter = yf }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwCTDMPerf1DayIntervalNumber" { return "Cpwctdmperf1Dayintervalnumber" }
    if yname == "cpwCTDMPerf1DayIntervalValidData" { return "Cpwctdmperf1Dayintervalvaliddata" }
    if yname == "cpwCTDMPerf1DayIntervalDuration" { return "Cpwctdmperf1Dayintervalduration" }
    if yname == "cpwCTDMPerf1DayIntervalMissingPkts" { return "Cpwctdmperf1Dayintervalmissingpkts" }
    if yname == "cpwCTDMPerf1DayIntervalPktsReOrder" { return "Cpwctdmperf1Dayintervalpktsreorder" }
    if yname == "cpwCTDMPerf1DayIntervalJtrBfrUnderruns" { return "Cpwctdmperf1Dayintervaljtrbfrunderruns" }
    if yname == "cpwCTDMPerf1DayIntervalMisOrderDropped" { return "Cpwctdmperf1Dayintervalmisorderdropped" }
    if yname == "cpwCTDMPerf1DayIntervalMalformedPkt" { return "Cpwctdmperf1Dayintervalmalformedpkt" }
    if yname == "cpwCTDMPerf1DayIntervalESs" { return "Cpwctdmperf1Dayintervaless" }
    if yname == "cpwCTDMPerf1DayIntervalSESs" { return "Cpwctdmperf1Dayintervalsess" }
    if yname == "cpwCTDMPerf1DayIntervalUASs" { return "Cpwctdmperf1Dayintervaluass" }
    if yname == "cpwCTDMPerf1DayIntervalFC" { return "Cpwctdmperf1Dayintervalfc" }
    return ""
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetSegmentPath() string {
    return "cpwCTDMPerf1DayIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwctdmperf1Dayintervalentry.Cpwvcindex) + "']" + "[cpwCTDMPerf1DayIntervalNumber='" + fmt.Sprintf("%v", cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalnumber) + "']"
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwctdmperf1Dayintervalentry.Cpwvcindex
    leafs["cpwCTDMPerf1DayIntervalNumber"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalnumber
    leafs["cpwCTDMPerf1DayIntervalValidData"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalvaliddata
    leafs["cpwCTDMPerf1DayIntervalDuration"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalduration
    leafs["cpwCTDMPerf1DayIntervalMissingPkts"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmissingpkts
    leafs["cpwCTDMPerf1DayIntervalPktsReOrder"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalpktsreorder
    leafs["cpwCTDMPerf1DayIntervalJtrBfrUnderruns"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaljtrbfrunderruns
    leafs["cpwCTDMPerf1DayIntervalMisOrderDropped"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmisorderdropped
    leafs["cpwCTDMPerf1DayIntervalMalformedPkt"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalmalformedpkt
    leafs["cpwCTDMPerf1DayIntervalESs"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaless
    leafs["cpwCTDMPerf1DayIntervalSESs"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalsess
    leafs["cpwCTDMPerf1DayIntervalUASs"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervaluass
    leafs["cpwCTDMPerf1DayIntervalFC"] = cpwctdmperf1Dayintervalentry.Cpwctdmperf1Dayintervalfc
    return leafs
}

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetYangName() string { return "cpwCTDMPerf1DayIntervalEntry" }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) SetParent(parent types.Entity) { cpwctdmperf1Dayintervalentry.parent = parent }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetParent() types.Entity { return cpwctdmperf1Dayintervalentry.parent }

func (cpwctdmperf1Dayintervalentry *CISCOIETFPWTDMMIB_Cpwctdmperf1Dayintervaltable_Cpwctdmperf1Dayintervalentry) GetParentYangName() string { return "cpwCTDMPerf1DayIntervalTable" }

