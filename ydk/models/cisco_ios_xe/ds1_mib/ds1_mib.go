// The MIB module to describe DS1, E1, DS2, and
// E2 interfaces objects.
package ds1_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ds1_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DS1-MIB DS1-MIB}", reflect.TypeOf(DS1MIB{}))
    ydk.RegisterEntity("DS1-MIB:DS1-MIB", reflect.TypeOf(DS1MIB{}))
}

// DS1MIB
type DS1MIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The DS1 Configuration table.
    Dsx1Configtable DS1MIB_Dsx1Configtable

    // The DS1 current table contains various statistics being collected for the
    // current 15 minute interval.
    Dsx1Currenttable DS1MIB_Dsx1Currenttable

    // The DS1 Interval Table contains various statistics collected by each DS1
    // Interface over the previous 24 hours of operation.  The past 24 hours are
    // broken into 96 completed 15 minute intervals.  Each row in this table
    // represents one such interval (identified by dsx1IntervalNumber) for one
    // specific instance (identified by dsx1IntervalIndex).
    Dsx1Intervaltable DS1MIB_Dsx1Intervaltable

    // The DS1 Total Table contains the cumulative sum of the various statistics
    // for the 24 hour period preceding the current interval.
    Dsx1Totaltable DS1MIB_Dsx1Totaltable

    // The DS1 Far End Current table contains various statistics being collected
    // for the current 15 minute interval.  The statistics are collected from the
    // far end messages on the Facilities Data Link.  The definitions are the same
    // as described for the near-end information.
    Dsx1Farendcurrenttable DS1MIB_Dsx1Farendcurrenttable

    // The DS1 Far End Interval Table contains various statistics collected by
    // each DS1 interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals. Each row in this
    // table represents one such interval (identified by dsx1FarEndIntervalNumber)
    // for one specific instance (identified by dsx1FarEndIntervalIndex).
    Dsx1Farendintervaltable DS1MIB_Dsx1Farendintervaltable

    // The DS1 Far End Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx1Farendtotaltable DS1MIB_Dsx1Farendtotaltable

    // This table is deprecated in favour of using ifStackTable.  The table was
    // mandatory for systems dividing a DS1 into channels containing different
    // data streams that are of local interest.  Systems which are indifferent to
    // data content, such as CSUs, need not implement it.  The DS1 fractional
    // table identifies which DS1 channels associated with a CSU are being used to
    // support a logical interface, i.e., an entry in the interfaces table from
    // the Internet-standard MIB.  For example, consider an application managing a
    // North American ISDN Primary Rate link whose division is a 384 kbit/s H1 _B_
    // Channel for Video, a second H1 for data to a primary routing peer, and 12
    // 64 kbit/s H0 _B_ Channels. Consider that some subset of the H0 channels are
    // used for voice and the remainder are available for dynamic data calls.  We
    // count a total of 14 interfaces multiplexed onto the DS1 interface. Six DS1
    // channels (for the sake of the example, channels 1..6) are used for Video,
    // six more (7..11 and 13) are used for data, and the remaining 12 are are in
    // channels 12 and 14..24.  Let us further imagine that ifIndex 2 is of type
    // DS1 and refers to the DS1 interface, and that the interfaces layered onto
    // it are numbered 3..16.  We might describe the allocation of channels, in
    // the dsx1FracTable, as follows: dsx1FracIfIndex.2. 1 = 3 
    // dsx1FracIfIndex.2.13 = 4 dsx1FracIfIndex.2. 2 = 3  dsx1FracIfIndex.2.14 = 6
    // dsx1FracIfIndex.2. 3 = 3  dsx1FracIfIndex.2.15 = 7 dsx1FracIfIndex.2. 4 = 3
    // dsx1FracIfIndex.2.16 = 8 dsx1FracIfIndex.2. 5 = 3  dsx1FracIfIndex.2.17 = 9
    // dsx1FracIfIndex.2. 6 = 3  dsx1FracIfIndex.2.18 = 10 dsx1FracIfIndex.2. 7 =
    // 4  dsx1FracIfIndex.2.19 = 11 dsx1FracIfIndex.2. 8 = 4  dsx1FracIfIndex.2.20
    // = 12 dsx1FracIfIndex.2. 9 = 4  dsx1FracIfIndex.2.21 = 13
    // dsx1FracIfIndex.2.10 = 4  dsx1FracIfIndex.2.22 = 14 dsx1FracIfIndex.2.11 =
    // 4  dsx1FracIfIndex.2.23 = 15 dsx1FracIfIndex.2.12 = 5  dsx1FracIfIndex.2.24
    // = 16  For North American (DS1) interfaces, there are 24 legal channels,
    // numbered 1 through 24.  For G.704 interfaces, there are 31 legal channels,
    // numbered 1 through 31.  The channels (1..31) correspond directly to the
    // equivalently numbered time-slots.
    Dsx1Fractable DS1MIB_Dsx1Fractable

    // The DS1 Channel Mapping table.  This table maps a DS1 channel number on a
    // particular DS3 into an ifIndex.  In the presence of DS2s, this table can be
    // used to map a DS2 channel number on a DS3 into an ifIndex, or used to map a
    // DS1 channel number on a DS2 onto an ifIndex.
    Dsx1Chanmappingtable DS1MIB_Dsx1Chanmappingtable
}

func (dS1MIB *DS1MIB) GetFilter() yfilter.YFilter { return dS1MIB.YFilter }

func (dS1MIB *DS1MIB) SetFilter(yf yfilter.YFilter) { dS1MIB.YFilter = yf }

func (dS1MIB *DS1MIB) GetGoName(yname string) string {
    if yname == "dsx1ConfigTable" { return "Dsx1Configtable" }
    if yname == "dsx1CurrentTable" { return "Dsx1Currenttable" }
    if yname == "dsx1IntervalTable" { return "Dsx1Intervaltable" }
    if yname == "dsx1TotalTable" { return "Dsx1Totaltable" }
    if yname == "dsx1FarEndCurrentTable" { return "Dsx1Farendcurrenttable" }
    if yname == "dsx1FarEndIntervalTable" { return "Dsx1Farendintervaltable" }
    if yname == "dsx1FarEndTotalTable" { return "Dsx1Farendtotaltable" }
    if yname == "dsx1FracTable" { return "Dsx1Fractable" }
    if yname == "dsx1ChanMappingTable" { return "Dsx1Chanmappingtable" }
    return ""
}

func (dS1MIB *DS1MIB) GetSegmentPath() string {
    return "DS1-MIB:DS1-MIB"
}

func (dS1MIB *DS1MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1ConfigTable" {
        return &dS1MIB.Dsx1Configtable
    }
    if childYangName == "dsx1CurrentTable" {
        return &dS1MIB.Dsx1Currenttable
    }
    if childYangName == "dsx1IntervalTable" {
        return &dS1MIB.Dsx1Intervaltable
    }
    if childYangName == "dsx1TotalTable" {
        return &dS1MIB.Dsx1Totaltable
    }
    if childYangName == "dsx1FarEndCurrentTable" {
        return &dS1MIB.Dsx1Farendcurrenttable
    }
    if childYangName == "dsx1FarEndIntervalTable" {
        return &dS1MIB.Dsx1Farendintervaltable
    }
    if childYangName == "dsx1FarEndTotalTable" {
        return &dS1MIB.Dsx1Farendtotaltable
    }
    if childYangName == "dsx1FracTable" {
        return &dS1MIB.Dsx1Fractable
    }
    if childYangName == "dsx1ChanMappingTable" {
        return &dS1MIB.Dsx1Chanmappingtable
    }
    return nil
}

func (dS1MIB *DS1MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dsx1ConfigTable"] = &dS1MIB.Dsx1Configtable
    children["dsx1CurrentTable"] = &dS1MIB.Dsx1Currenttable
    children["dsx1IntervalTable"] = &dS1MIB.Dsx1Intervaltable
    children["dsx1TotalTable"] = &dS1MIB.Dsx1Totaltable
    children["dsx1FarEndCurrentTable"] = &dS1MIB.Dsx1Farendcurrenttable
    children["dsx1FarEndIntervalTable"] = &dS1MIB.Dsx1Farendintervaltable
    children["dsx1FarEndTotalTable"] = &dS1MIB.Dsx1Farendtotaltable
    children["dsx1FracTable"] = &dS1MIB.Dsx1Fractable
    children["dsx1ChanMappingTable"] = &dS1MIB.Dsx1Chanmappingtable
    return children
}

func (dS1MIB *DS1MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dS1MIB *DS1MIB) GetBundleName() string { return "cisco_ios_xe" }

func (dS1MIB *DS1MIB) GetYangName() string { return "DS1-MIB" }

func (dS1MIB *DS1MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dS1MIB *DS1MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dS1MIB *DS1MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dS1MIB *DS1MIB) SetParent(parent types.Entity) { dS1MIB.parent = parent }

func (dS1MIB *DS1MIB) GetParent() types.Entity { return dS1MIB.parent }

func (dS1MIB *DS1MIB) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Configtable
// The DS1 Configuration table.
type DS1MIB_Dsx1Configtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Configuration table. The type is slice of
    // DS1MIB_Dsx1Configtable_Dsx1Configentry.
    Dsx1Configentry []DS1MIB_Dsx1Configtable_Dsx1Configentry
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetFilter() yfilter.YFilter { return dsx1Configtable.YFilter }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) SetFilter(yf yfilter.YFilter) { dsx1Configtable.YFilter = yf }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetGoName(yname string) string {
    if yname == "dsx1ConfigEntry" { return "Dsx1Configentry" }
    return ""
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetSegmentPath() string {
    return "dsx1ConfigTable"
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1ConfigEntry" {
        for _, c := range dsx1Configtable.Dsx1Configentry {
            if dsx1Configtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Configtable_Dsx1Configentry{}
        dsx1Configtable.Dsx1Configentry = append(dsx1Configtable.Dsx1Configentry, child)
        return &dsx1Configtable.Dsx1Configentry[len(dsx1Configtable.Dsx1Configentry)-1]
    }
    return nil
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Configtable.Dsx1Configentry {
        children[dsx1Configtable.Dsx1Configentry[i].GetSegmentPath()] = &dsx1Configtable.Dsx1Configentry[i]
    }
    return children
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetYangName() string { return "dsx1ConfigTable" }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) SetParent(parent types.Entity) { dsx1Configtable.parent = parent }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetParent() types.Entity { return dsx1Configtable.parent }

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Configtable_Dsx1Configentry
// An entry in the DS1 Configuration table.
type DS1MIB_Dsx1Configtable_Dsx1Configentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object should be made equal to ifIndex.  The
    // next paragraph describes its previous usage. Making the object equal to
    // ifIndex allows proper use of ifStackTable and ds0/ds0bundle mibs. 
    // Previously, this object is the identifier of a DS1 Interface on a managed
    // device.  If there is an ifEntry that is directly associated with this and
    // only this DS1 interface, it should have the same value as ifIndex. 
    // Otherwise, number the dsx1LineIndices with an unique identifier following
    // the rules of choosing a number that is greater than ifNumber and numbering
    // the inside interfaces (e.g., equipment side) with even numbers and outside
    // interfaces (e.g, network side) with odd numbers. The type is interface{}
    // with range: 1..2147483647.
    Dsx1Lineindex interface{}

    // This value for this object is equal to the value of ifIndex from the
    // Interfaces table of MIB II (RFC 1213). The type is interface{} with range:
    // 1..2147483647.
    Dsx1Ifindex interface{}

    // The number of seconds that have elapsed since the beginning of the near end
    // current error- measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx1Timeelapsed interface{}

    // The number of previous near end intervals for which data was collected. 
    // The value will be 96 unless the interface was brought online within the
    // last 24 hours, in which case the value will be the number of complete 15
    // minute near end intervals since the interface has been online.  In the case
    // where the agent is a proxy, it is possible that some intervals are
    // unavailable.  In this case, this interval is the maximum interval number
    // for which data is available. The type is interface{} with range: 0..96.
    Dsx1Validintervals interface{}

    // This variable indicates  the  variety  of  DS1 Line  implementing  this 
    // circuit.  The type of circuit affects the number of bits  per  second that 
    // the circuit can reasonably carry, as well as the interpretation of the 
    // usage  and  error statistics.  The values, in sequence, describe:  TITLE:  
    // SPECIFICATION: dsx1ESF         Extended SuperFrame DS1 (T1.107) dsx1D4     
    // AT&T D4 format DS1 (T1.107) dsx1E1          ITU-T Recommendation G.704     
    // (Table 4a) dsx1E1-CRC      ITU-T Recommendation G.704                 
    // (Table 4b) dsxE1-MF        G.704 (Table 4a) with TS16                 
    // multiframing enabled dsx1E1-CRC-MF   G.704 (Table 4b) with TS16            
    // multiframing enabled dsx1Unframed    DS1 with No Framing dsx1E1Unframed  E1
    // with No Framing (G.703) dsx1DS2M12      DS2 frame format (T1.107) dsx1E2   
    // E2 frame format (G.704)  For clarification, the capacity for each E1 type
    // is as listed below: dsx1E1Unframed - E1, no framing = 32 x 64k = 2048k
    // dsx1E1 or dsx1E1CRC - E1, with framing,    no signalling = 31 x 64k = 1984k
    // dsx1E1MF or dsx1E1CRCMF - E1, with framing,    signalling = 30 x 64k =
    // 1920k  For further information See ITU-T Recomm G.704. The type is
    // Dsx1Linetype.
    Dsx1Linetype interface{}

    // This variable describes the variety of Zero Code Suppression used on this
    // interface, which in turn affects a number of its characteristics.  dsx1JBZS
    // refers the Jammed Bit Zero Suppression, in which the AT&T specification of
    // at least one pulse every 8 bit periods is literally implemented by forcing
    // a pulse in bit 8 of each channel. Thus, only seven bits per channel, or
    // 1.344 Mbps, is available for data.  dsx1B8ZS refers to the use of a
    // specified pattern of normal bits and bipolar violations which are used to
    // replace a sequence of eight zero bits.  ANSI Clear Channels may use
    // dsx1ZBTSI, or Zero Byte Time Slot Interchange.  E1 links, with or without
    // CRC, use dsx1HDB3 or dsx1AMI.  dsx1AMI refers to a mode wherein no zero
    // code suppression is present and the line encoding does not solve the
    // problem directly.  In this application, the higher layer must provide data
    // which meets or exceeds the pulse density requirements, such as inverting
    // HDLC data.  dsx1B6ZS refers to the user of a specifed pattern of normal
    // bits and bipolar violations which are used to replace a sequence of six
    // zero bits.  Used for DS2. The type is Dsx1Linecoding.
    Dsx1Linecoding interface{}

    // This variable indicates what type of code is being sent across the DS1
    // interface by the device. Setting this variable causes the interface to send
    // the code requested.  The values mean: dsx1SendNoCode sending looped or
    // normal data  dsx1SendLineCode sending a request for a line loopback 
    // dsx1SendPayloadCode sending a request for a payload loopback 
    // dsx1SendResetCode sending a loopback termination request  dsx1SendQRS
    // sending a Quasi-Random Signal  (QRS)  test pattern  dsx1Send511Pattern
    // sending a 511 bit fixed test pattern  dsx1Send3in24Pattern sending a fixed
    // test pattern of 3 bits set in 24  dsx1SendOtherTestPattern sending a test
    // pattern  other  than  those described by this object. The type is
    // Dsx1Sendcode.
    Dsx1Sendcode interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. The type is string with
    // length: 0..255.
    Dsx1Circuitidentifier interface{}

    // This variable represents the desired loopback configuration of the DS1
    // interface.  Agents supporting read/write access should return
    // inconsistentValue in response to a requested loopback state that the
    // interface does not support.  The values mean:  dsx1NoLoop  Not in the
    // loopback state.  A device that is not capable of performing a loopback on
    // the interface shall always return this as its value.  dsx1PayloadLoop  The
    // received signal at this interface is looped through the device.  Typically
    // the received signal is looped back for retransmission after it has passed
    // through the device's framing function.  dsx1LineLoop  The received signal
    // at this interface does not go through the device (minimum penetration) but
    // is looped back out.  dsx1OtherLoop  Loopbacks that are not defined here. 
    // dsx1InwardLoop  The transmitted signal at this interface is looped back and
    // received by the same interface. What is transmitted onto the line is
    // product dependent.  dsx1DualLoop  Both dsx1LineLoop and dsx1InwardLoop will
    // be active simultaneously. The type is Dsx1Loopbackconfig.
    Dsx1Loopbackconfig interface{}

    // This variable indicates the Line Status of the interface.  It contains
    // loopback, failure, received 'alarm' and transmitted 'alarms information. 
    // The dsx1LineStatus is a bit map represented as a sum, therefore, it can
    // represent multiple failures (alarms) and a LoopbackState simultaneously. 
    // dsx1NoAlarm must be set if and only if no other flag is set.  If the
    // dsx1loopbackState bit is set, the loopback in effect can be determined from
    // the dsx1loopbackConfig object. The various bit positions are: 1    
    // dsx1NoAlarm           No alarm present 2     dsx1RcvFarEndLOF      Far end
    // LOF (a.k.a., Yellow Alarm) 4     dsx1XmtFarEndLOF      Near end sending LOF
    // Indication 8     dsx1RcvAIS            Far end sending AIS 16    
    // dsx1XmtAIS            Near end sending AIS 32     dsx1LossOfFrame      
    // Near end LOF (a.k.a., Red Alarm) 64     dsx1LossOfSignal      Near end Loss
    // Of Signal 128     dsx1LoopbackState     Near end is looped 256    
    // dsx1T16AIS            E1 TS16 AIS 512     dsx1RcvFarEndLOMF     Far End
    // Sending TS16 LOMF 1024     dsx1XmtFarEndLOMF     Near End Sending TS16 LOMF
    // 2048     dsx1RcvTestCode       Near End detects a test code 4096    
    // dsx1OtherFailure      any line status not defined here 8192    
    // dsx1UnavailSigState   Near End in Unavailable Signal                  State
    // 16384     dsx1NetEquipOOS       Carrier Equipment Out of Service 32768    
    // dsx1RcvPayloadAIS     DS2 Payload AIS 65536     dsx1Ds2PerfThreshold  DS2
    // Performance Threshold                  Exceeded. The type is interface{}
    // with range: 1..131071.
    Dsx1Linestatus interface{}

    // 'none' indicates that no bits are reserved for signaling on this channel. 
    // 'robbedBit' indicates that DS1 Robbed Bit  Sig- naling is in use. 
    // 'bitOriented' indicates that E1 Channel  Asso- ciated Signaling is in use. 
    // 'messageOriented' indicates that Common  Chan- nel Signaling is in use
    // either on channel 16 of an E1 link or channel 24 of a DS1. The type is
    // Dsx1Signalmode.
    Dsx1Signalmode interface{}

    // The source of Transmit Clock. 'loopTiming' indicates that the recovered re-
    // ceive clock is used as the transmit clock.  'localTiming' indicates that a
    // local clock source is used or when an external clock is attached to the box
    // containing the interface.  'throughTiming' indicates that recovered re-
    // ceive clock from another interface is used as the transmit clock. The type
    // is Dsx1Transmitclocksource.
    Dsx1Transmitclocksource interface{}

    // This bitmap describes the use of  the  facili- ties data link, and is the
    // sum of the capabili- ties.  Set any bits that are appropriate:  other(1),
    // dsx1AnsiT1403(2), dsx1Att54016(4), dsx1FdlNone(8)   'other' indicates that
    // a protocol  other  than one following is used.   'dsx1AnsiT1403' refers to
    // the  FDL  exchange recommended by ANSI.   'dsx1Att54016' refers to ESF FDL
    // exchanges.   'dsx1FdlNone' indicates that the device  does not use the FDL.
    // The type is interface{} with range: 1..15.
    Dsx1Fdl interface{}

    // The number of intervals in the range from 0 to dsx1ValidIntervals for which
    // no data is available.  This object will typically be zero except in cases
    // where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx1Invalidintervals interface{}

    // The length of the ds1 line in meters. This objects provides information for
    // line build out circuitry.  This object is only useful if the interface has
    // configurable line build out circuitry. The type is interface{} with range:
    // 0..64000. Units are meters.
    Dsx1Linelength interface{}

    // The value of MIB II's sysUpTime object at the time this DS1 entered its
    // current line status state.  If the current state was entered prior to the
    // last re-initialization of the proxy-agent, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    Dsx1Linestatuslastchange interface{}

    // Indicates whether dsx1LineStatusChange traps should be generated for this
    // interface. The type is Dsx1Linestatuschangetrapenable.
    Dsx1Linestatuschangetrapenable interface{}

    // This variable represents the current state of the loopback on the DS1
    // interface.  It contains information about loopbacks established by a
    // manager and remotely from the far end.  The dsx1LoopbackStatus is a bit map
    // represented as a sum, therefore is can represent multiple loopbacks
    // simultaneously.  The various bit positions are:  1  dsx1NoLoopback  2 
    // dsx1NearEndPayloadLoopback  4  dsx1NearEndLineLoopback  8 
    // dsx1NearEndOtherLoopback 16  dsx1NearEndInwardLoopback 32 
    // dsx1FarEndPayloadLoopback 64  dsx1FarEndLineLoopback. The type is
    // interface{} with range: 1..127.
    Dsx1Loopbackstatus interface{}

    // This variable represents the channel number of the DS1/E1 on its parent
    // Ds2/E2 or DS3/E3.  A value of 0 indicated this DS1/E1 does not have a
    // parent DS3/E3. The type is interface{} with range: 0..28.
    Dsx1Ds1Channelnumber interface{}

    // Indicates whether this ds1/e1 is channelized or unchannelized.  The value
    // of enabledDs0 indicates that this is a DS1 channelized into DS0s.  The
    // value of enabledDs1 indicated that this is a DS2 channelized into DS1s. 
    // Setting this value will cause the creation or deletion of entries in the
    // ifTable for the DS0s that are within the DS1. The type is
    // Dsx1Channelization.
    Dsx1Channelization interface{}
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetFilter() yfilter.YFilter { return dsx1Configentry.YFilter }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) SetFilter(yf yfilter.YFilter) { dsx1Configentry.YFilter = yf }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetGoName(yname string) string {
    if yname == "dsx1LineIndex" { return "Dsx1Lineindex" }
    if yname == "dsx1IfIndex" { return "Dsx1Ifindex" }
    if yname == "dsx1TimeElapsed" { return "Dsx1Timeelapsed" }
    if yname == "dsx1ValidIntervals" { return "Dsx1Validintervals" }
    if yname == "dsx1LineType" { return "Dsx1Linetype" }
    if yname == "dsx1LineCoding" { return "Dsx1Linecoding" }
    if yname == "dsx1SendCode" { return "Dsx1Sendcode" }
    if yname == "dsx1CircuitIdentifier" { return "Dsx1Circuitidentifier" }
    if yname == "dsx1LoopbackConfig" { return "Dsx1Loopbackconfig" }
    if yname == "dsx1LineStatus" { return "Dsx1Linestatus" }
    if yname == "dsx1SignalMode" { return "Dsx1Signalmode" }
    if yname == "dsx1TransmitClockSource" { return "Dsx1Transmitclocksource" }
    if yname == "dsx1Fdl" { return "Dsx1Fdl" }
    if yname == "dsx1InvalidIntervals" { return "Dsx1Invalidintervals" }
    if yname == "dsx1LineLength" { return "Dsx1Linelength" }
    if yname == "dsx1LineStatusLastChange" { return "Dsx1Linestatuslastchange" }
    if yname == "dsx1LineStatusChangeTrapEnable" { return "Dsx1Linestatuschangetrapenable" }
    if yname == "dsx1LoopbackStatus" { return "Dsx1Loopbackstatus" }
    if yname == "dsx1Ds1ChannelNumber" { return "Dsx1Ds1Channelnumber" }
    if yname == "dsx1Channelization" { return "Dsx1Channelization" }
    return ""
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetSegmentPath() string {
    return "dsx1ConfigEntry" + "[dsx1LineIndex='" + fmt.Sprintf("%v", dsx1Configentry.Dsx1Lineindex) + "']"
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1LineIndex"] = dsx1Configentry.Dsx1Lineindex
    leafs["dsx1IfIndex"] = dsx1Configentry.Dsx1Ifindex
    leafs["dsx1TimeElapsed"] = dsx1Configentry.Dsx1Timeelapsed
    leafs["dsx1ValidIntervals"] = dsx1Configentry.Dsx1Validintervals
    leafs["dsx1LineType"] = dsx1Configentry.Dsx1Linetype
    leafs["dsx1LineCoding"] = dsx1Configentry.Dsx1Linecoding
    leafs["dsx1SendCode"] = dsx1Configentry.Dsx1Sendcode
    leafs["dsx1CircuitIdentifier"] = dsx1Configentry.Dsx1Circuitidentifier
    leafs["dsx1LoopbackConfig"] = dsx1Configentry.Dsx1Loopbackconfig
    leafs["dsx1LineStatus"] = dsx1Configentry.Dsx1Linestatus
    leafs["dsx1SignalMode"] = dsx1Configentry.Dsx1Signalmode
    leafs["dsx1TransmitClockSource"] = dsx1Configentry.Dsx1Transmitclocksource
    leafs["dsx1Fdl"] = dsx1Configentry.Dsx1Fdl
    leafs["dsx1InvalidIntervals"] = dsx1Configentry.Dsx1Invalidintervals
    leafs["dsx1LineLength"] = dsx1Configentry.Dsx1Linelength
    leafs["dsx1LineStatusLastChange"] = dsx1Configentry.Dsx1Linestatuslastchange
    leafs["dsx1LineStatusChangeTrapEnable"] = dsx1Configentry.Dsx1Linestatuschangetrapenable
    leafs["dsx1LoopbackStatus"] = dsx1Configentry.Dsx1Loopbackstatus
    leafs["dsx1Ds1ChannelNumber"] = dsx1Configentry.Dsx1Ds1Channelnumber
    leafs["dsx1Channelization"] = dsx1Configentry.Dsx1Channelization
    return leafs
}

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetYangName() string { return "dsx1ConfigEntry" }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) SetParent(parent types.Entity) { dsx1Configentry.parent = parent }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetParent() types.Entity { return dsx1Configentry.parent }

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetParentYangName() string { return "dsx1ConfigTable" }

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization represents ifTable for the DS0s that are within the DS1.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization_disabled DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization = "disabled"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization_enabledDs0 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization = "enabledDs0"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization_enabledDs1 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Channelization = "enabledDs1"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding represents for DS2.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1JBZS DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1JBZS"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1B8ZS DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1B8ZS"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1HDB3 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1HDB3"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1ZBTSI DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1ZBTSI"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1AMI DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1AMI"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_other DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "other"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding_dsx1B6ZS DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linecoding = "dsx1B6ZS"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable represents should be generated for this interface.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable_enabled DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable = "enabled"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable_disabled DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linestatuschangetrapenable = "disabled"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype represents For further information See ITU-T Recomm G.704
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_other DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "other"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1ESF DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1ESF"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1D4 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1D4"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1E1 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1E1"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1E1CRC DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1E1CRC"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1E1MF DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1E1MF"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1E1CRCMF DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1E1CRCMF"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1Unframed DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1Unframed"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1E1Unframed DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1E1Unframed"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx1DS2M12 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx1DS2M12"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype_dsx2E2 DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Linetype = "dsx2E2"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig represents active simultaneously.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1NoLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1NoLoop"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1PayloadLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1PayloadLoop"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1LineLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1LineLoop"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1OtherLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1OtherLoop"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1InwardLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1InwardLoop"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig_dsx1DualLoop DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Loopbackconfig = "dsx1DualLoop"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode represents described by this object
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendNoCode DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendNoCode"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendLineCode DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendLineCode"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendPayloadCode DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendPayloadCode"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendResetCode DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendResetCode"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendQRS DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendQRS"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1Send511Pattern DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1Send511Pattern"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1Send3in24Pattern DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1Send3in24Pattern"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode_dsx1SendOtherTestPattern DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Sendcode = "dsx1SendOtherTestPattern"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode represents an E1 link or channel 24 of a DS1.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode_none DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode = "none"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode_robbedBit DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode = "robbedBit"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode_bitOriented DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode = "bitOriented"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode_messageOriented DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode = "messageOriented"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode_other DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Signalmode = "other"
)

// DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource represents the transmit clock.
type DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource string

const (
    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource_loopTiming DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource = "loopTiming"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource_localTiming DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource = "localTiming"

    DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource_throughTiming DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Transmitclocksource = "throughTiming"
)

// DS1MIB_Dsx1Currenttable
// The DS1 current table contains various statistics
// being collected for the current 15 minute
// interval.
type DS1MIB_Dsx1Currenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Current table. The type is slice of
    // DS1MIB_Dsx1Currenttable_Dsx1Currententry.
    Dsx1Currententry []DS1MIB_Dsx1Currenttable_Dsx1Currententry
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetFilter() yfilter.YFilter { return dsx1Currenttable.YFilter }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) SetFilter(yf yfilter.YFilter) { dsx1Currenttable.YFilter = yf }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetGoName(yname string) string {
    if yname == "dsx1CurrentEntry" { return "Dsx1Currententry" }
    return ""
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetSegmentPath() string {
    return "dsx1CurrentTable"
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1CurrentEntry" {
        for _, c := range dsx1Currenttable.Dsx1Currententry {
            if dsx1Currenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Currenttable_Dsx1Currententry{}
        dsx1Currenttable.Dsx1Currententry = append(dsx1Currenttable.Dsx1Currententry, child)
        return &dsx1Currenttable.Dsx1Currententry[len(dsx1Currenttable.Dsx1Currententry)-1]
    }
    return nil
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Currenttable.Dsx1Currententry {
        children[dsx1Currenttable.Dsx1Currententry[i].GetSegmentPath()] = &dsx1Currenttable.Dsx1Currententry[i]
    }
    return children
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetYangName() string { return "dsx1CurrentTable" }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) SetParent(parent types.Entity) { dsx1Currenttable.parent = parent }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetParent() types.Entity { return dsx1Currenttable.parent }

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Currenttable_Dsx1Currententry
// An entry in the DS1 Current table.
type DS1MIB_Dsx1Currenttable_Dsx1Currententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies  the DS1
    // interface to which this entry is applicable. The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1Currentindex interface{}

    // The number of Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentess interface{}

    // The number of Severely Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentsess interface{}

    // The number of Severely Errored Framing Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1Currentsefss interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentuass interface{}

    // The number of Controlled Slip Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentcsss interface{}

    // The number of Path Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentpcvs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentless interface{}

    // The number of Bursty Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentbess interface{}

    // The number of Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Currentdms interface{}

    // The number of Line Code Violations (LCVs). The type is interface{} with
    // range: 0..4294967295.
    Dsx1Currentlcvs interface{}
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetFilter() yfilter.YFilter { return dsx1Currententry.YFilter }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) SetFilter(yf yfilter.YFilter) { dsx1Currententry.YFilter = yf }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetGoName(yname string) string {
    if yname == "dsx1CurrentIndex" { return "Dsx1Currentindex" }
    if yname == "dsx1CurrentESs" { return "Dsx1Currentess" }
    if yname == "dsx1CurrentSESs" { return "Dsx1Currentsess" }
    if yname == "dsx1CurrentSEFSs" { return "Dsx1Currentsefss" }
    if yname == "dsx1CurrentUASs" { return "Dsx1Currentuass" }
    if yname == "dsx1CurrentCSSs" { return "Dsx1Currentcsss" }
    if yname == "dsx1CurrentPCVs" { return "Dsx1Currentpcvs" }
    if yname == "dsx1CurrentLESs" { return "Dsx1Currentless" }
    if yname == "dsx1CurrentBESs" { return "Dsx1Currentbess" }
    if yname == "dsx1CurrentDMs" { return "Dsx1Currentdms" }
    if yname == "dsx1CurrentLCVs" { return "Dsx1Currentlcvs" }
    return ""
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetSegmentPath() string {
    return "dsx1CurrentEntry" + "[dsx1CurrentIndex='" + fmt.Sprintf("%v", dsx1Currententry.Dsx1Currentindex) + "']"
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1CurrentIndex"] = dsx1Currententry.Dsx1Currentindex
    leafs["dsx1CurrentESs"] = dsx1Currententry.Dsx1Currentess
    leafs["dsx1CurrentSESs"] = dsx1Currententry.Dsx1Currentsess
    leafs["dsx1CurrentSEFSs"] = dsx1Currententry.Dsx1Currentsefss
    leafs["dsx1CurrentUASs"] = dsx1Currententry.Dsx1Currentuass
    leafs["dsx1CurrentCSSs"] = dsx1Currententry.Dsx1Currentcsss
    leafs["dsx1CurrentPCVs"] = dsx1Currententry.Dsx1Currentpcvs
    leafs["dsx1CurrentLESs"] = dsx1Currententry.Dsx1Currentless
    leafs["dsx1CurrentBESs"] = dsx1Currententry.Dsx1Currentbess
    leafs["dsx1CurrentDMs"] = dsx1Currententry.Dsx1Currentdms
    leafs["dsx1CurrentLCVs"] = dsx1Currententry.Dsx1Currentlcvs
    return leafs
}

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetYangName() string { return "dsx1CurrentEntry" }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) SetParent(parent types.Entity) { dsx1Currententry.parent = parent }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetParent() types.Entity { return dsx1Currententry.parent }

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetParentYangName() string { return "dsx1CurrentTable" }

// DS1MIB_Dsx1Intervaltable
// The DS1 Interval Table contains various
// statistics collected by each DS1 Interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.  Each row in this table represents one
// such interval (identified by dsx1IntervalNumber)
// for one specific instance (identified by
// dsx1IntervalIndex).
type DS1MIB_Dsx1Intervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Interval table. The type is slice of
    // DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry.
    Dsx1Intervalentry []DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetFilter() yfilter.YFilter { return dsx1Intervaltable.YFilter }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) SetFilter(yf yfilter.YFilter) { dsx1Intervaltable.YFilter = yf }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetGoName(yname string) string {
    if yname == "dsx1IntervalEntry" { return "Dsx1Intervalentry" }
    return ""
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetSegmentPath() string {
    return "dsx1IntervalTable"
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1IntervalEntry" {
        for _, c := range dsx1Intervaltable.Dsx1Intervalentry {
            if dsx1Intervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry{}
        dsx1Intervaltable.Dsx1Intervalentry = append(dsx1Intervaltable.Dsx1Intervalentry, child)
        return &dsx1Intervaltable.Dsx1Intervalentry[len(dsx1Intervaltable.Dsx1Intervalentry)-1]
    }
    return nil
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Intervaltable.Dsx1Intervalentry {
        children[dsx1Intervaltable.Dsx1Intervalentry[i].GetSegmentPath()] = &dsx1Intervaltable.Dsx1Intervalentry[i]
    }
    return children
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetYangName() string { return "dsx1IntervalTable" }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) SetParent(parent types.Entity) { dsx1Intervaltable.parent = parent }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetParent() types.Entity { return dsx1Intervaltable.parent }

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry
// An entry in the DS1 Interval table.
type DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1Intervalindex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx1Intervalnumber interface{}

    // The number of Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervaless interface{}

    // The number of Severely Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervalsess interface{}

    // The number of Severely Errored Framing Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1Intervalsefss interface{}

    // The number of Unavailable Seconds.  This object may decrease if the
    // occurance of unavailable seconds occurs across an inteval boundary. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Intervaluass interface{}

    // The number of Controlled Slip Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervalcsss interface{}

    // The number of Path Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervalpcvs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervalless interface{}

    // The number of Bursty Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervalbess interface{}

    // The number of Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervaldms interface{}

    // The number of Line Code Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Intervallcvs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx1Intervalvaliddata interface{}
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetFilter() yfilter.YFilter { return dsx1Intervalentry.YFilter }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) SetFilter(yf yfilter.YFilter) { dsx1Intervalentry.YFilter = yf }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetGoName(yname string) string {
    if yname == "dsx1IntervalIndex" { return "Dsx1Intervalindex" }
    if yname == "dsx1IntervalNumber" { return "Dsx1Intervalnumber" }
    if yname == "dsx1IntervalESs" { return "Dsx1Intervaless" }
    if yname == "dsx1IntervalSESs" { return "Dsx1Intervalsess" }
    if yname == "dsx1IntervalSEFSs" { return "Dsx1Intervalsefss" }
    if yname == "dsx1IntervalUASs" { return "Dsx1Intervaluass" }
    if yname == "dsx1IntervalCSSs" { return "Dsx1Intervalcsss" }
    if yname == "dsx1IntervalPCVs" { return "Dsx1Intervalpcvs" }
    if yname == "dsx1IntervalLESs" { return "Dsx1Intervalless" }
    if yname == "dsx1IntervalBESs" { return "Dsx1Intervalbess" }
    if yname == "dsx1IntervalDMs" { return "Dsx1Intervaldms" }
    if yname == "dsx1IntervalLCVs" { return "Dsx1Intervallcvs" }
    if yname == "dsx1IntervalValidData" { return "Dsx1Intervalvaliddata" }
    return ""
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetSegmentPath() string {
    return "dsx1IntervalEntry" + "[dsx1IntervalIndex='" + fmt.Sprintf("%v", dsx1Intervalentry.Dsx1Intervalindex) + "']" + "[dsx1IntervalNumber='" + fmt.Sprintf("%v", dsx1Intervalentry.Dsx1Intervalnumber) + "']"
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1IntervalIndex"] = dsx1Intervalentry.Dsx1Intervalindex
    leafs["dsx1IntervalNumber"] = dsx1Intervalentry.Dsx1Intervalnumber
    leafs["dsx1IntervalESs"] = dsx1Intervalentry.Dsx1Intervaless
    leafs["dsx1IntervalSESs"] = dsx1Intervalentry.Dsx1Intervalsess
    leafs["dsx1IntervalSEFSs"] = dsx1Intervalentry.Dsx1Intervalsefss
    leafs["dsx1IntervalUASs"] = dsx1Intervalentry.Dsx1Intervaluass
    leafs["dsx1IntervalCSSs"] = dsx1Intervalentry.Dsx1Intervalcsss
    leafs["dsx1IntervalPCVs"] = dsx1Intervalentry.Dsx1Intervalpcvs
    leafs["dsx1IntervalLESs"] = dsx1Intervalentry.Dsx1Intervalless
    leafs["dsx1IntervalBESs"] = dsx1Intervalentry.Dsx1Intervalbess
    leafs["dsx1IntervalDMs"] = dsx1Intervalentry.Dsx1Intervaldms
    leafs["dsx1IntervalLCVs"] = dsx1Intervalentry.Dsx1Intervallcvs
    leafs["dsx1IntervalValidData"] = dsx1Intervalentry.Dsx1Intervalvaliddata
    return leafs
}

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetYangName() string { return "dsx1IntervalEntry" }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) SetParent(parent types.Entity) { dsx1Intervalentry.parent = parent }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetParent() types.Entity { return dsx1Intervalentry.parent }

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetParentYangName() string { return "dsx1IntervalTable" }

// DS1MIB_Dsx1Totaltable
// The DS1 Total Table contains the cumulative sum
// of the various statistics for the 24 hour period
// preceding the current interval.
type DS1MIB_Dsx1Totaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Total table. The type is slice of
    // DS1MIB_Dsx1Totaltable_Dsx1Totalentry.
    Dsx1Totalentry []DS1MIB_Dsx1Totaltable_Dsx1Totalentry
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetFilter() yfilter.YFilter { return dsx1Totaltable.YFilter }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) SetFilter(yf yfilter.YFilter) { dsx1Totaltable.YFilter = yf }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetGoName(yname string) string {
    if yname == "dsx1TotalEntry" { return "Dsx1Totalentry" }
    return ""
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetSegmentPath() string {
    return "dsx1TotalTable"
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1TotalEntry" {
        for _, c := range dsx1Totaltable.Dsx1Totalentry {
            if dsx1Totaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Totaltable_Dsx1Totalentry{}
        dsx1Totaltable.Dsx1Totalentry = append(dsx1Totaltable.Dsx1Totalentry, child)
        return &dsx1Totaltable.Dsx1Totalentry[len(dsx1Totaltable.Dsx1Totalentry)-1]
    }
    return nil
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Totaltable.Dsx1Totalentry {
        children[dsx1Totaltable.Dsx1Totalentry[i].GetSegmentPath()] = &dsx1Totaltable.Dsx1Totalentry[i]
    }
    return children
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetYangName() string { return "dsx1TotalTable" }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) SetParent(parent types.Entity) { dsx1Totaltable.parent = parent }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetParent() types.Entity { return dsx1Totaltable.parent }

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Totaltable_Dsx1Totalentry
// An entry in the DS1 Total table.
type DS1MIB_Dsx1Totaltable_Dsx1Totalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1Totalindex interface{}

    // The sum of Errored Seconds encountered by a DS1 interface in the previous
    // 24 hour interval. Invalid 15 minute intervals count as 0. The type is
    // interface{} with range: 0..4294967295.
    Dsx1Totaless interface{}

    // The number of Severely Errored Seconds encountered by a DS1 interface in
    // the previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Totalsess interface{}

    // The number of Severely Errored Framing Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1Totalsefss interface{}

    // The number of Unavailable Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1Totaluass interface{}

    // The number of Controlled Slip Seconds encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Totalcsss interface{}

    // The number of Path Coding Violations encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Totalpcvs interface{}

    // The number of Line Errored Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1Totalless interface{}

    // The number of Bursty Errored Seconds (BESs) encountered by a DS1 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1Totalbess interface{}

    // The number of Degraded Minutes (DMs) encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Totaldms interface{}

    // The number of Line Code Violations (LCVs) encountered by a DS1 interface in
    // the current 15 minute interval.  Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1Totallcvs interface{}
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetFilter() yfilter.YFilter { return dsx1Totalentry.YFilter }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) SetFilter(yf yfilter.YFilter) { dsx1Totalentry.YFilter = yf }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetGoName(yname string) string {
    if yname == "dsx1TotalIndex" { return "Dsx1Totalindex" }
    if yname == "dsx1TotalESs" { return "Dsx1Totaless" }
    if yname == "dsx1TotalSESs" { return "Dsx1Totalsess" }
    if yname == "dsx1TotalSEFSs" { return "Dsx1Totalsefss" }
    if yname == "dsx1TotalUASs" { return "Dsx1Totaluass" }
    if yname == "dsx1TotalCSSs" { return "Dsx1Totalcsss" }
    if yname == "dsx1TotalPCVs" { return "Dsx1Totalpcvs" }
    if yname == "dsx1TotalLESs" { return "Dsx1Totalless" }
    if yname == "dsx1TotalBESs" { return "Dsx1Totalbess" }
    if yname == "dsx1TotalDMs" { return "Dsx1Totaldms" }
    if yname == "dsx1TotalLCVs" { return "Dsx1Totallcvs" }
    return ""
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetSegmentPath() string {
    return "dsx1TotalEntry" + "[dsx1TotalIndex='" + fmt.Sprintf("%v", dsx1Totalentry.Dsx1Totalindex) + "']"
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1TotalIndex"] = dsx1Totalentry.Dsx1Totalindex
    leafs["dsx1TotalESs"] = dsx1Totalentry.Dsx1Totaless
    leafs["dsx1TotalSESs"] = dsx1Totalentry.Dsx1Totalsess
    leafs["dsx1TotalSEFSs"] = dsx1Totalentry.Dsx1Totalsefss
    leafs["dsx1TotalUASs"] = dsx1Totalentry.Dsx1Totaluass
    leafs["dsx1TotalCSSs"] = dsx1Totalentry.Dsx1Totalcsss
    leafs["dsx1TotalPCVs"] = dsx1Totalentry.Dsx1Totalpcvs
    leafs["dsx1TotalLESs"] = dsx1Totalentry.Dsx1Totalless
    leafs["dsx1TotalBESs"] = dsx1Totalentry.Dsx1Totalbess
    leafs["dsx1TotalDMs"] = dsx1Totalentry.Dsx1Totaldms
    leafs["dsx1TotalLCVs"] = dsx1Totalentry.Dsx1Totallcvs
    return leafs
}

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetYangName() string { return "dsx1TotalEntry" }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) SetParent(parent types.Entity) { dsx1Totalentry.parent = parent }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetParent() types.Entity { return dsx1Totalentry.parent }

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetParentYangName() string { return "dsx1TotalTable" }

// DS1MIB_Dsx1Farendcurrenttable
// The DS1 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end messages on the Facilities Data
// Link.  The definitions are the same as described
// for the near-end information.
type DS1MIB_Dsx1Farendcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Current table. The type is slice of
    // DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry.
    Dsx1Farendcurrententry []DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetFilter() yfilter.YFilter { return dsx1Farendcurrenttable.YFilter }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) SetFilter(yf yfilter.YFilter) { dsx1Farendcurrenttable.YFilter = yf }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetGoName(yname string) string {
    if yname == "dsx1FarEndCurrentEntry" { return "Dsx1Farendcurrententry" }
    return ""
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetSegmentPath() string {
    return "dsx1FarEndCurrentTable"
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1FarEndCurrentEntry" {
        for _, c := range dsx1Farendcurrenttable.Dsx1Farendcurrententry {
            if dsx1Farendcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry{}
        dsx1Farendcurrenttable.Dsx1Farendcurrententry = append(dsx1Farendcurrenttable.Dsx1Farendcurrententry, child)
        return &dsx1Farendcurrenttable.Dsx1Farendcurrententry[len(dsx1Farendcurrenttable.Dsx1Farendcurrententry)-1]
    }
    return nil
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Farendcurrenttable.Dsx1Farendcurrententry {
        children[dsx1Farendcurrenttable.Dsx1Farendcurrententry[i].GetSegmentPath()] = &dsx1Farendcurrenttable.Dsx1Farendcurrententry[i]
    }
    return children
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetYangName() string { return "dsx1FarEndCurrentTable" }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) SetParent(parent types.Entity) { dsx1Farendcurrenttable.parent = parent }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetParent() types.Entity { return dsx1Farendcurrenttable.parent }

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry
// An entry in the DS1 Far End Current table.
type DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1Farendcurrentindex interface{}

    // The number of seconds that have elapsed since the beginning of the far end
    // current error-measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx1Farendtimeelapsed interface{}

    // The number of previous far end intervals for which data was collected.  The
    // value will be 96 unless the interface was brought online within the last 24
    // hours, in which case the value will be the number of complete 15 minute far
    // end intervals since the interface has been online. The type is interface{}
    // with range: 0..96.
    Dsx1Farendvalidintervals interface{}

    // The number of Far End Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendcurrentess interface{}

    // The number of Far End Severely Errored Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1Farendcurrentsess interface{}

    // The number of Far End Severely Errored Framing Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx1Farendcurrentsefss interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendcurrentuass interface{}

    // The number of Far End Controlled Slip Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendcurrentcsss interface{}

    // The number of Far End Line Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendcurrentless interface{}

    // The number of Far End Path Coding Violations. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendcurrentpcvs interface{}

    // The number of Far End Bursty Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendcurrentbess interface{}

    // The number of Far End Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendcurrentdms interface{}

    // The number of intervals in the range from 0 to dsx1FarEndValidIntervals for
    // which no data is available.  This object will typically be zero except in
    // cases where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx1Farendinvalidintervals interface{}
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetFilter() yfilter.YFilter { return dsx1Farendcurrententry.YFilter }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) SetFilter(yf yfilter.YFilter) { dsx1Farendcurrententry.YFilter = yf }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetGoName(yname string) string {
    if yname == "dsx1FarEndCurrentIndex" { return "Dsx1Farendcurrentindex" }
    if yname == "dsx1FarEndTimeElapsed" { return "Dsx1Farendtimeelapsed" }
    if yname == "dsx1FarEndValidIntervals" { return "Dsx1Farendvalidintervals" }
    if yname == "dsx1FarEndCurrentESs" { return "Dsx1Farendcurrentess" }
    if yname == "dsx1FarEndCurrentSESs" { return "Dsx1Farendcurrentsess" }
    if yname == "dsx1FarEndCurrentSEFSs" { return "Dsx1Farendcurrentsefss" }
    if yname == "dsx1FarEndCurrentUASs" { return "Dsx1Farendcurrentuass" }
    if yname == "dsx1FarEndCurrentCSSs" { return "Dsx1Farendcurrentcsss" }
    if yname == "dsx1FarEndCurrentLESs" { return "Dsx1Farendcurrentless" }
    if yname == "dsx1FarEndCurrentPCVs" { return "Dsx1Farendcurrentpcvs" }
    if yname == "dsx1FarEndCurrentBESs" { return "Dsx1Farendcurrentbess" }
    if yname == "dsx1FarEndCurrentDMs" { return "Dsx1Farendcurrentdms" }
    if yname == "dsx1FarEndInvalidIntervals" { return "Dsx1Farendinvalidintervals" }
    return ""
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetSegmentPath() string {
    return "dsx1FarEndCurrentEntry" + "[dsx1FarEndCurrentIndex='" + fmt.Sprintf("%v", dsx1Farendcurrententry.Dsx1Farendcurrentindex) + "']"
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1FarEndCurrentIndex"] = dsx1Farendcurrententry.Dsx1Farendcurrentindex
    leafs["dsx1FarEndTimeElapsed"] = dsx1Farendcurrententry.Dsx1Farendtimeelapsed
    leafs["dsx1FarEndValidIntervals"] = dsx1Farendcurrententry.Dsx1Farendvalidintervals
    leafs["dsx1FarEndCurrentESs"] = dsx1Farendcurrententry.Dsx1Farendcurrentess
    leafs["dsx1FarEndCurrentSESs"] = dsx1Farendcurrententry.Dsx1Farendcurrentsess
    leafs["dsx1FarEndCurrentSEFSs"] = dsx1Farendcurrententry.Dsx1Farendcurrentsefss
    leafs["dsx1FarEndCurrentUASs"] = dsx1Farendcurrententry.Dsx1Farendcurrentuass
    leafs["dsx1FarEndCurrentCSSs"] = dsx1Farendcurrententry.Dsx1Farendcurrentcsss
    leafs["dsx1FarEndCurrentLESs"] = dsx1Farendcurrententry.Dsx1Farendcurrentless
    leafs["dsx1FarEndCurrentPCVs"] = dsx1Farendcurrententry.Dsx1Farendcurrentpcvs
    leafs["dsx1FarEndCurrentBESs"] = dsx1Farendcurrententry.Dsx1Farendcurrentbess
    leafs["dsx1FarEndCurrentDMs"] = dsx1Farendcurrententry.Dsx1Farendcurrentdms
    leafs["dsx1FarEndInvalidIntervals"] = dsx1Farendcurrententry.Dsx1Farendinvalidintervals
    return leafs
}

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetYangName() string { return "dsx1FarEndCurrentEntry" }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) SetParent(parent types.Entity) { dsx1Farendcurrententry.parent = parent }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetParent() types.Entity { return dsx1Farendcurrententry.parent }

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetParentYangName() string { return "dsx1FarEndCurrentTable" }

// DS1MIB_Dsx1Farendintervaltable
// The DS1 Far End Interval Table contains various
// statistics collected by each DS1 interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals. Each row in this table represents one
// such interval (identified by
// dsx1FarEndIntervalNumber) for one specific
// instance (identified by dsx1FarEndIntervalIndex).
type DS1MIB_Dsx1Farendintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Interval table. The type is slice of
    // DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry.
    Dsx1Farendintervalentry []DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetFilter() yfilter.YFilter { return dsx1Farendintervaltable.YFilter }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) SetFilter(yf yfilter.YFilter) { dsx1Farendintervaltable.YFilter = yf }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetGoName(yname string) string {
    if yname == "dsx1FarEndIntervalEntry" { return "Dsx1Farendintervalentry" }
    return ""
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetSegmentPath() string {
    return "dsx1FarEndIntervalTable"
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1FarEndIntervalEntry" {
        for _, c := range dsx1Farendintervaltable.Dsx1Farendintervalentry {
            if dsx1Farendintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry{}
        dsx1Farendintervaltable.Dsx1Farendintervalentry = append(dsx1Farendintervaltable.Dsx1Farendintervalentry, child)
        return &dsx1Farendintervaltable.Dsx1Farendintervalentry[len(dsx1Farendintervaltable.Dsx1Farendintervalentry)-1]
    }
    return nil
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Farendintervaltable.Dsx1Farendintervalentry {
        children[dsx1Farendintervaltable.Dsx1Farendintervalentry[i].GetSegmentPath()] = &dsx1Farendintervaltable.Dsx1Farendintervalentry[i]
    }
    return children
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetYangName() string { return "dsx1FarEndIntervalTable" }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) SetParent(parent types.Entity) { dsx1Farendintervaltable.parent = parent }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetParent() types.Entity { return dsx1Farendintervaltable.parent }

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry
// An entry in the DS1 Far End Interval table.
type DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1Farendintervalindex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx1Farendintervalnumber interface{}

    // The number of Far End Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendintervaless interface{}

    // The number of Far End Severely Errored Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1Farendintervalsess interface{}

    // The number of Far End Severely Errored Framing Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx1Farendintervalsefss interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendintervaluass interface{}

    // The number of Far End Controlled Slip Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendintervalcsss interface{}

    // The number of Far End Line Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendintervalless interface{}

    // The number of Far End Path Coding Violations. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendintervalpcvs interface{}

    // The number of Far End Bursty Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1Farendintervalbess interface{}

    // The number of Far End Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1Farendintervaldms interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx1Farendintervalvaliddata interface{}
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetFilter() yfilter.YFilter { return dsx1Farendintervalentry.YFilter }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) SetFilter(yf yfilter.YFilter) { dsx1Farendintervalentry.YFilter = yf }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetGoName(yname string) string {
    if yname == "dsx1FarEndIntervalIndex" { return "Dsx1Farendintervalindex" }
    if yname == "dsx1FarEndIntervalNumber" { return "Dsx1Farendintervalnumber" }
    if yname == "dsx1FarEndIntervalESs" { return "Dsx1Farendintervaless" }
    if yname == "dsx1FarEndIntervalSESs" { return "Dsx1Farendintervalsess" }
    if yname == "dsx1FarEndIntervalSEFSs" { return "Dsx1Farendintervalsefss" }
    if yname == "dsx1FarEndIntervalUASs" { return "Dsx1Farendintervaluass" }
    if yname == "dsx1FarEndIntervalCSSs" { return "Dsx1Farendintervalcsss" }
    if yname == "dsx1FarEndIntervalLESs" { return "Dsx1Farendintervalless" }
    if yname == "dsx1FarEndIntervalPCVs" { return "Dsx1Farendintervalpcvs" }
    if yname == "dsx1FarEndIntervalBESs" { return "Dsx1Farendintervalbess" }
    if yname == "dsx1FarEndIntervalDMs" { return "Dsx1Farendintervaldms" }
    if yname == "dsx1FarEndIntervalValidData" { return "Dsx1Farendintervalvaliddata" }
    return ""
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetSegmentPath() string {
    return "dsx1FarEndIntervalEntry" + "[dsx1FarEndIntervalIndex='" + fmt.Sprintf("%v", dsx1Farendintervalentry.Dsx1Farendintervalindex) + "']" + "[dsx1FarEndIntervalNumber='" + fmt.Sprintf("%v", dsx1Farendintervalentry.Dsx1Farendintervalnumber) + "']"
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1FarEndIntervalIndex"] = dsx1Farendintervalentry.Dsx1Farendintervalindex
    leafs["dsx1FarEndIntervalNumber"] = dsx1Farendintervalentry.Dsx1Farendintervalnumber
    leafs["dsx1FarEndIntervalESs"] = dsx1Farendintervalentry.Dsx1Farendintervaless
    leafs["dsx1FarEndIntervalSESs"] = dsx1Farendintervalentry.Dsx1Farendintervalsess
    leafs["dsx1FarEndIntervalSEFSs"] = dsx1Farendintervalentry.Dsx1Farendintervalsefss
    leafs["dsx1FarEndIntervalUASs"] = dsx1Farendintervalentry.Dsx1Farendintervaluass
    leafs["dsx1FarEndIntervalCSSs"] = dsx1Farendintervalentry.Dsx1Farendintervalcsss
    leafs["dsx1FarEndIntervalLESs"] = dsx1Farendintervalentry.Dsx1Farendintervalless
    leafs["dsx1FarEndIntervalPCVs"] = dsx1Farendintervalentry.Dsx1Farendintervalpcvs
    leafs["dsx1FarEndIntervalBESs"] = dsx1Farendintervalentry.Dsx1Farendintervalbess
    leafs["dsx1FarEndIntervalDMs"] = dsx1Farendintervalentry.Dsx1Farendintervaldms
    leafs["dsx1FarEndIntervalValidData"] = dsx1Farendintervalentry.Dsx1Farendintervalvaliddata
    return leafs
}

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetYangName() string { return "dsx1FarEndIntervalEntry" }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) SetParent(parent types.Entity) { dsx1Farendintervalentry.parent = parent }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetParent() types.Entity { return dsx1Farendintervalentry.parent }

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetParentYangName() string { return "dsx1FarEndIntervalTable" }

// DS1MIB_Dsx1Farendtotaltable
// The DS1 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS1MIB_Dsx1Farendtotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Total table. The type is slice of
    // DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry.
    Dsx1Farendtotalentry []DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetFilter() yfilter.YFilter { return dsx1Farendtotaltable.YFilter }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) SetFilter(yf yfilter.YFilter) { dsx1Farendtotaltable.YFilter = yf }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetGoName(yname string) string {
    if yname == "dsx1FarEndTotalEntry" { return "Dsx1Farendtotalentry" }
    return ""
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetSegmentPath() string {
    return "dsx1FarEndTotalTable"
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1FarEndTotalEntry" {
        for _, c := range dsx1Farendtotaltable.Dsx1Farendtotalentry {
            if dsx1Farendtotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry{}
        dsx1Farendtotaltable.Dsx1Farendtotalentry = append(dsx1Farendtotaltable.Dsx1Farendtotalentry, child)
        return &dsx1Farendtotaltable.Dsx1Farendtotalentry[len(dsx1Farendtotaltable.Dsx1Farendtotalentry)-1]
    }
    return nil
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Farendtotaltable.Dsx1Farendtotalentry {
        children[dsx1Farendtotaltable.Dsx1Farendtotalentry[i].GetSegmentPath()] = &dsx1Farendtotaltable.Dsx1Farendtotalentry[i]
    }
    return children
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetYangName() string { return "dsx1FarEndTotalTable" }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) SetParent(parent types.Entity) { dsx1Farendtotaltable.parent = parent }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetParent() types.Entity { return dsx1Farendtotaltable.parent }

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry
// An entry in the DS1 Far End Total table.
type DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1Farendtotalindex interface{}

    // The number of Far End Errored Seconds encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Farendtotaless interface{}

    // The number of Far End Severely Errored Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1Farendtotalsess interface{}

    // The number of Far End Severely Errored Framing Seconds encountered by a DS1
    // interface in the previous 24 hour interval. Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1Farendtotalsefss interface{}

    // The number of Unavailable Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1Farendtotaluass interface{}

    // The number of Far End Controlled Slip Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1Farendtotalcsss interface{}

    // The number of Far End Line Errored Seconds encountered by a DS1 interface
    // in the previous 24 hour interval.  Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1Farendtotalless interface{}

    // The number of Far End Path Coding Violations reported via the far end block
    // error count encountered by a DS1 interface in the previous 24 hour
    // interval.  Invalid 15 minute intervals count as 0. The type is interface{}
    // with range: 0..4294967295.
    Dsx1Farendtotalpcvs interface{}

    // The number of Bursty Errored Seconds (BESs) encountered by a DS1 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1Farendtotalbess interface{}

    // The number of Degraded Minutes (DMs) encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1Farendtotaldms interface{}
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetFilter() yfilter.YFilter { return dsx1Farendtotalentry.YFilter }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) SetFilter(yf yfilter.YFilter) { dsx1Farendtotalentry.YFilter = yf }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetGoName(yname string) string {
    if yname == "dsx1FarEndTotalIndex" { return "Dsx1Farendtotalindex" }
    if yname == "dsx1FarEndTotalESs" { return "Dsx1Farendtotaless" }
    if yname == "dsx1FarEndTotalSESs" { return "Dsx1Farendtotalsess" }
    if yname == "dsx1FarEndTotalSEFSs" { return "Dsx1Farendtotalsefss" }
    if yname == "dsx1FarEndTotalUASs" { return "Dsx1Farendtotaluass" }
    if yname == "dsx1FarEndTotalCSSs" { return "Dsx1Farendtotalcsss" }
    if yname == "dsx1FarEndTotalLESs" { return "Dsx1Farendtotalless" }
    if yname == "dsx1FarEndTotalPCVs" { return "Dsx1Farendtotalpcvs" }
    if yname == "dsx1FarEndTotalBESs" { return "Dsx1Farendtotalbess" }
    if yname == "dsx1FarEndTotalDMs" { return "Dsx1Farendtotaldms" }
    return ""
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetSegmentPath() string {
    return "dsx1FarEndTotalEntry" + "[dsx1FarEndTotalIndex='" + fmt.Sprintf("%v", dsx1Farendtotalentry.Dsx1Farendtotalindex) + "']"
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1FarEndTotalIndex"] = dsx1Farendtotalentry.Dsx1Farendtotalindex
    leafs["dsx1FarEndTotalESs"] = dsx1Farendtotalentry.Dsx1Farendtotaless
    leafs["dsx1FarEndTotalSESs"] = dsx1Farendtotalentry.Dsx1Farendtotalsess
    leafs["dsx1FarEndTotalSEFSs"] = dsx1Farendtotalentry.Dsx1Farendtotalsefss
    leafs["dsx1FarEndTotalUASs"] = dsx1Farendtotalentry.Dsx1Farendtotaluass
    leafs["dsx1FarEndTotalCSSs"] = dsx1Farendtotalentry.Dsx1Farendtotalcsss
    leafs["dsx1FarEndTotalLESs"] = dsx1Farendtotalentry.Dsx1Farendtotalless
    leafs["dsx1FarEndTotalPCVs"] = dsx1Farendtotalentry.Dsx1Farendtotalpcvs
    leafs["dsx1FarEndTotalBESs"] = dsx1Farendtotalentry.Dsx1Farendtotalbess
    leafs["dsx1FarEndTotalDMs"] = dsx1Farendtotalentry.Dsx1Farendtotaldms
    return leafs
}

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetYangName() string { return "dsx1FarEndTotalEntry" }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) SetParent(parent types.Entity) { dsx1Farendtotalentry.parent = parent }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetParent() types.Entity { return dsx1Farendtotalentry.parent }

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetParentYangName() string { return "dsx1FarEndTotalTable" }

// DS1MIB_Dsx1Fractable
// This table is deprecated in favour of using
// ifStackTable.
// 
// The table was mandatory for systems dividing a DS1
// into channels containing different data streams
// that are of local interest.  Systems which are
// indifferent to data content, such as CSUs, need
// not implement it.
// 
// The DS1 fractional table identifies which DS1
// channels associated with a CSU are being used to
// support a logical interface, i.e., an entry in the
// interfaces table from the Internet-standard MIB.
// 
// For example, consider an application managing a
// North American ISDN Primary Rate link whose
// division is a 384 kbit/s H1 _B_ Channel for Video,
// a second H1 for data to a primary routing peer,
// and 12 64 kbit/s H0 _B_ Channels. Consider that
// some subset of the H0 channels are used for voice
// and the remainder are available for dynamic data
// calls.
// 
// We count a total of 14 interfaces multiplexed onto
// the DS1 interface. Six DS1 channels (for the sake
// of the example, channels 1..6) are used for Video,
// six more (7..11 and 13) are used for data, and the
// remaining 12 are are in channels 12 and 14..24.
// 
// Let us further imagine that ifIndex 2 is of type
// DS1 and refers to the DS1 interface, and that the
// interfaces layered onto it are numbered 3..16.
// 
// We might describe the allocation of channels, in
// the dsx1FracTable, as follows:
// dsx1FracIfIndex.2. 1 = 3  dsx1FracIfIndex.2.13 = 4
// dsx1FracIfIndex.2. 2 = 3  dsx1FracIfIndex.2.14 = 6
// dsx1FracIfIndex.2. 3 = 3  dsx1FracIfIndex.2.15 = 7
// dsx1FracIfIndex.2. 4 = 3  dsx1FracIfIndex.2.16 = 8
// dsx1FracIfIndex.2. 5 = 3  dsx1FracIfIndex.2.17 = 9
// dsx1FracIfIndex.2. 6 = 3  dsx1FracIfIndex.2.18 = 10
// dsx1FracIfIndex.2. 7 = 4  dsx1FracIfIndex.2.19 = 11
// dsx1FracIfIndex.2. 8 = 4  dsx1FracIfIndex.2.20 = 12
// dsx1FracIfIndex.2. 9 = 4  dsx1FracIfIndex.2.21 = 13
// dsx1FracIfIndex.2.10 = 4  dsx1FracIfIndex.2.22 = 14
// dsx1FracIfIndex.2.11 = 4  dsx1FracIfIndex.2.23 = 15
// dsx1FracIfIndex.2.12 = 5  dsx1FracIfIndex.2.24 = 16
// 
// For North American (DS1) interfaces, there are 24
// legal channels, numbered 1 through 24.
// 
// For G.704 interfaces, there are 31 legal channels,
// numbered 1 through 31.  The channels (1..31)
// correspond directly to the equivalently numbered
// time-slots.
type DS1MIB_Dsx1Fractable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Fractional table. The type is slice of
    // DS1MIB_Dsx1Fractable_Dsx1Fracentry.
    Dsx1Fracentry []DS1MIB_Dsx1Fractable_Dsx1Fracentry
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetFilter() yfilter.YFilter { return dsx1Fractable.YFilter }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) SetFilter(yf yfilter.YFilter) { dsx1Fractable.YFilter = yf }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetGoName(yname string) string {
    if yname == "dsx1FracEntry" { return "Dsx1Fracentry" }
    return ""
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetSegmentPath() string {
    return "dsx1FracTable"
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1FracEntry" {
        for _, c := range dsx1Fractable.Dsx1Fracentry {
            if dsx1Fractable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Fractable_Dsx1Fracentry{}
        dsx1Fractable.Dsx1Fracentry = append(dsx1Fractable.Dsx1Fracentry, child)
        return &dsx1Fractable.Dsx1Fracentry[len(dsx1Fractable.Dsx1Fracentry)-1]
    }
    return nil
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Fractable.Dsx1Fracentry {
        children[dsx1Fractable.Dsx1Fracentry[i].GetSegmentPath()] = &dsx1Fractable.Dsx1Fracentry[i]
    }
    return children
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetYangName() string { return "dsx1FracTable" }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) SetParent(parent types.Entity) { dsx1Fractable.parent = parent }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetParent() types.Entity { return dsx1Fractable.parent }

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Fractable_Dsx1Fracentry
// An entry in the DS1 Fractional table.
type DS1MIB_Dsx1Fractable_Dsx1Fracentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies  the DS1
    // interface  to which this entry is applicable The interface identified by a 
    // particular value  of  this  index is the same interface as identified by
    // the same value  an  dsx1LineIndex object instance. The type is interface{}
    // with range: 1..2147483647.
    Dsx1Fracindex interface{}

    // This attribute is a key. The channel number for this entry. The type is
    // interface{} with range: 1..31.
    Dsx1Fracnumber interface{}

    // An index value that uniquely identifies an interface.  The interface
    // identified by a particular value of this index is the same  interface as 
    // identified by the same value an ifIndex object instance. If no interface is
    // currently using a channel, the value should be zero.  If a single interface
    // occupies more  than  one  time slot,  that ifIndex value will be found in
    // multiple time slots. The type is interface{} with range: 1..2147483647.
    Dsx1Fracifindex interface{}
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetFilter() yfilter.YFilter { return dsx1Fracentry.YFilter }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) SetFilter(yf yfilter.YFilter) { dsx1Fracentry.YFilter = yf }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetGoName(yname string) string {
    if yname == "dsx1FracIndex" { return "Dsx1Fracindex" }
    if yname == "dsx1FracNumber" { return "Dsx1Fracnumber" }
    if yname == "dsx1FracIfIndex" { return "Dsx1Fracifindex" }
    return ""
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetSegmentPath() string {
    return "dsx1FracEntry" + "[dsx1FracIndex='" + fmt.Sprintf("%v", dsx1Fracentry.Dsx1Fracindex) + "']" + "[dsx1FracNumber='" + fmt.Sprintf("%v", dsx1Fracentry.Dsx1Fracnumber) + "']"
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx1FracIndex"] = dsx1Fracentry.Dsx1Fracindex
    leafs["dsx1FracNumber"] = dsx1Fracentry.Dsx1Fracnumber
    leafs["dsx1FracIfIndex"] = dsx1Fracentry.Dsx1Fracifindex
    return leafs
}

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetYangName() string { return "dsx1FracEntry" }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) SetParent(parent types.Entity) { dsx1Fracentry.parent = parent }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetParent() types.Entity { return dsx1Fracentry.parent }

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetParentYangName() string { return "dsx1FracTable" }

// DS1MIB_Dsx1Chanmappingtable
// The DS1 Channel Mapping table.  This table maps a
// DS1 channel number on a particular DS3 into an
// ifIndex.  In the presence of DS2s, this table can
// be used to map a DS2 channel number on a DS3 into
// an ifIndex, or used to map a DS1 channel number on
// a DS2 onto an ifIndex.
type DS1MIB_Dsx1Chanmappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS1 Channel Mapping table.  There is an entry in this table
    // corresponding to each ds1 ifEntry within any interface that is channelized
    // to the individual ds1 ifEntry level.  This table is intended to facilitate
    // mapping from channelized interface / channel number to DS1 ifEntry.  (e.g.
    // mapping (DS3 ifIndex, DS1 Channel Number) -> ifIndex)  While this table
    // provides information that can also be found in the ifStackTable and
    // dsx1ConfigTable, it provides this same information with a single table
    // lookup, rather than by walking the ifStackTable to find the various
    // constituent ds1 ifTable entries, and testing various dsx1ConfigTable
    // entries to check for the entry with the applicable DS1 channel number. The
    // type is slice of DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry.
    Dsx1Chanmappingentry []DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetFilter() yfilter.YFilter { return dsx1Chanmappingtable.YFilter }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) SetFilter(yf yfilter.YFilter) { dsx1Chanmappingtable.YFilter = yf }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetGoName(yname string) string {
    if yname == "dsx1ChanMappingEntry" { return "Dsx1Chanmappingentry" }
    return ""
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetSegmentPath() string {
    return "dsx1ChanMappingTable"
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx1ChanMappingEntry" {
        for _, c := range dsx1Chanmappingtable.Dsx1Chanmappingentry {
            if dsx1Chanmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry{}
        dsx1Chanmappingtable.Dsx1Chanmappingentry = append(dsx1Chanmappingtable.Dsx1Chanmappingentry, child)
        return &dsx1Chanmappingtable.Dsx1Chanmappingentry[len(dsx1Chanmappingtable.Dsx1Chanmappingentry)-1]
    }
    return nil
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx1Chanmappingtable.Dsx1Chanmappingentry {
        children[dsx1Chanmappingtable.Dsx1Chanmappingentry[i].GetSegmentPath()] = &dsx1Chanmappingtable.Dsx1Chanmappingentry[i]
    }
    return children
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetYangName() string { return "dsx1ChanMappingTable" }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) SetParent(parent types.Entity) { dsx1Chanmappingtable.parent = parent }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetParent() types.Entity { return dsx1Chanmappingtable.parent }

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetParentYangName() string { return "DS1-MIB" }

// DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry
// An entry in the DS1 Channel Mapping table.  There
// is an entry in this table corresponding to each
// ds1 ifEntry within any interface that is
// channelized to the individual ds1 ifEntry level.
// 
// This table is intended to facilitate mapping from
// channelized interface / channel number to DS1
// ifEntry.  (e.g. mapping (DS3 ifIndex, DS1 Channel
// Number) -> ifIndex)
// 
// While this table provides information that can
// also be found in the ifStackTable and
// dsx1ConfigTable, it provides this same information
// with a single table lookup, rather than by walking
// the ifStackTable to find the various constituent
// ds1 ifTable entries, and testing various
// dsx1ConfigTable entries to check for the entry
// with the applicable DS1 channel number.
type DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..28. Refers to
    // ds1_mib.DS1MIB_Dsx1Configtable_Dsx1Configentry_Dsx1Ds1Channelnumber
    Dsx1Ds1Channelnumber interface{}

    // This object indicates the ifIndex value assigned by the agent for the
    // individual ds1 ifEntry that corresponds to the given DS1 channel number
    // (specified by the INDEX element dsx1Ds1ChannelNumber) of the given
    // channelized interface (specified by INDEX element ifIndex). The type is
    // interface{} with range: 1..2147483647.
    Dsx1Chanmappedifindex interface{}
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetFilter() yfilter.YFilter { return dsx1Chanmappingentry.YFilter }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) SetFilter(yf yfilter.YFilter) { dsx1Chanmappingentry.YFilter = yf }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "dsx1Ds1ChannelNumber" { return "Dsx1Ds1Channelnumber" }
    if yname == "dsx1ChanMappedIfIndex" { return "Dsx1Chanmappedifindex" }
    return ""
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetSegmentPath() string {
    return "dsx1ChanMappingEntry" + "[ifIndex='" + fmt.Sprintf("%v", dsx1Chanmappingentry.Ifindex) + "']" + "[dsx1Ds1ChannelNumber='" + fmt.Sprintf("%v", dsx1Chanmappingentry.Dsx1Ds1Channelnumber) + "']"
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = dsx1Chanmappingentry.Ifindex
    leafs["dsx1Ds1ChannelNumber"] = dsx1Chanmappingentry.Dsx1Ds1Channelnumber
    leafs["dsx1ChanMappedIfIndex"] = dsx1Chanmappingentry.Dsx1Chanmappedifindex
    return leafs
}

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetYangName() string { return "dsx1ChanMappingEntry" }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) SetParent(parent types.Entity) { dsx1Chanmappingentry.parent = parent }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetParent() types.Entity { return dsx1Chanmappingentry.parent }

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetParentYangName() string { return "dsx1ChanMappingTable" }

