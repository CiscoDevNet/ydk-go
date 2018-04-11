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
    EntityData types.CommonEntityData
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

func (dS1MIB *DS1MIB) GetEntityData() *types.CommonEntityData {
    dS1MIB.EntityData.YFilter = dS1MIB.YFilter
    dS1MIB.EntityData.YangName = "DS1-MIB"
    dS1MIB.EntityData.BundleName = "cisco_ios_xe"
    dS1MIB.EntityData.ParentYangName = "DS1-MIB"
    dS1MIB.EntityData.SegmentPath = "DS1-MIB:DS1-MIB"
    dS1MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dS1MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dS1MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dS1MIB.EntityData.Children = make(map[string]types.YChild)
    dS1MIB.EntityData.Children["dsx1ConfigTable"] = types.YChild{"Dsx1Configtable", &dS1MIB.Dsx1Configtable}
    dS1MIB.EntityData.Children["dsx1CurrentTable"] = types.YChild{"Dsx1Currenttable", &dS1MIB.Dsx1Currenttable}
    dS1MIB.EntityData.Children["dsx1IntervalTable"] = types.YChild{"Dsx1Intervaltable", &dS1MIB.Dsx1Intervaltable}
    dS1MIB.EntityData.Children["dsx1TotalTable"] = types.YChild{"Dsx1Totaltable", &dS1MIB.Dsx1Totaltable}
    dS1MIB.EntityData.Children["dsx1FarEndCurrentTable"] = types.YChild{"Dsx1Farendcurrenttable", &dS1MIB.Dsx1Farendcurrenttable}
    dS1MIB.EntityData.Children["dsx1FarEndIntervalTable"] = types.YChild{"Dsx1Farendintervaltable", &dS1MIB.Dsx1Farendintervaltable}
    dS1MIB.EntityData.Children["dsx1FarEndTotalTable"] = types.YChild{"Dsx1Farendtotaltable", &dS1MIB.Dsx1Farendtotaltable}
    dS1MIB.EntityData.Children["dsx1FracTable"] = types.YChild{"Dsx1Fractable", &dS1MIB.Dsx1Fractable}
    dS1MIB.EntityData.Children["dsx1ChanMappingTable"] = types.YChild{"Dsx1Chanmappingtable", &dS1MIB.Dsx1Chanmappingtable}
    dS1MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dS1MIB.EntityData)
}

// DS1MIB_Dsx1Configtable
// The DS1 Configuration table.
type DS1MIB_Dsx1Configtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Configuration table. The type is slice of
    // DS1MIB_Dsx1Configtable_Dsx1Configentry.
    Dsx1Configentry []DS1MIB_Dsx1Configtable_Dsx1Configentry
}

func (dsx1Configtable *DS1MIB_Dsx1Configtable) GetEntityData() *types.CommonEntityData {
    dsx1Configtable.EntityData.YFilter = dsx1Configtable.YFilter
    dsx1Configtable.EntityData.YangName = "dsx1ConfigTable"
    dsx1Configtable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Configtable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Configtable.EntityData.SegmentPath = "dsx1ConfigTable"
    dsx1Configtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Configtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Configtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Configtable.EntityData.Children = make(map[string]types.YChild)
    dsx1Configtable.EntityData.Children["dsx1ConfigEntry"] = types.YChild{"Dsx1Configentry", nil}
    for i := range dsx1Configtable.Dsx1Configentry {
        dsx1Configtable.EntityData.Children[types.GetSegmentPath(&dsx1Configtable.Dsx1Configentry[i])] = types.YChild{"Dsx1Configentry", &dsx1Configtable.Dsx1Configentry[i]}
    }
    dsx1Configtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Configtable.EntityData)
}

// DS1MIB_Dsx1Configtable_Dsx1Configentry
// An entry in the DS1 Configuration table.
type DS1MIB_Dsx1Configtable_Dsx1Configentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Configentry *DS1MIB_Dsx1Configtable_Dsx1Configentry) GetEntityData() *types.CommonEntityData {
    dsx1Configentry.EntityData.YFilter = dsx1Configentry.YFilter
    dsx1Configentry.EntityData.YangName = "dsx1ConfigEntry"
    dsx1Configentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Configentry.EntityData.ParentYangName = "dsx1ConfigTable"
    dsx1Configentry.EntityData.SegmentPath = "dsx1ConfigEntry" + "[dsx1LineIndex='" + fmt.Sprintf("%v", dsx1Configentry.Dsx1Lineindex) + "']"
    dsx1Configentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Configentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Configentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Configentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Configentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Configentry.EntityData.Leafs["dsx1LineIndex"] = types.YLeaf{"Dsx1Lineindex", dsx1Configentry.Dsx1Lineindex}
    dsx1Configentry.EntityData.Leafs["dsx1IfIndex"] = types.YLeaf{"Dsx1Ifindex", dsx1Configentry.Dsx1Ifindex}
    dsx1Configentry.EntityData.Leafs["dsx1TimeElapsed"] = types.YLeaf{"Dsx1Timeelapsed", dsx1Configentry.Dsx1Timeelapsed}
    dsx1Configentry.EntityData.Leafs["dsx1ValidIntervals"] = types.YLeaf{"Dsx1Validintervals", dsx1Configentry.Dsx1Validintervals}
    dsx1Configentry.EntityData.Leafs["dsx1LineType"] = types.YLeaf{"Dsx1Linetype", dsx1Configentry.Dsx1Linetype}
    dsx1Configentry.EntityData.Leafs["dsx1LineCoding"] = types.YLeaf{"Dsx1Linecoding", dsx1Configentry.Dsx1Linecoding}
    dsx1Configentry.EntityData.Leafs["dsx1SendCode"] = types.YLeaf{"Dsx1Sendcode", dsx1Configentry.Dsx1Sendcode}
    dsx1Configentry.EntityData.Leafs["dsx1CircuitIdentifier"] = types.YLeaf{"Dsx1Circuitidentifier", dsx1Configentry.Dsx1Circuitidentifier}
    dsx1Configentry.EntityData.Leafs["dsx1LoopbackConfig"] = types.YLeaf{"Dsx1Loopbackconfig", dsx1Configentry.Dsx1Loopbackconfig}
    dsx1Configentry.EntityData.Leafs["dsx1LineStatus"] = types.YLeaf{"Dsx1Linestatus", dsx1Configentry.Dsx1Linestatus}
    dsx1Configentry.EntityData.Leafs["dsx1SignalMode"] = types.YLeaf{"Dsx1Signalmode", dsx1Configentry.Dsx1Signalmode}
    dsx1Configentry.EntityData.Leafs["dsx1TransmitClockSource"] = types.YLeaf{"Dsx1Transmitclocksource", dsx1Configentry.Dsx1Transmitclocksource}
    dsx1Configentry.EntityData.Leafs["dsx1Fdl"] = types.YLeaf{"Dsx1Fdl", dsx1Configentry.Dsx1Fdl}
    dsx1Configentry.EntityData.Leafs["dsx1InvalidIntervals"] = types.YLeaf{"Dsx1Invalidintervals", dsx1Configentry.Dsx1Invalidintervals}
    dsx1Configentry.EntityData.Leafs["dsx1LineLength"] = types.YLeaf{"Dsx1Linelength", dsx1Configentry.Dsx1Linelength}
    dsx1Configentry.EntityData.Leafs["dsx1LineStatusLastChange"] = types.YLeaf{"Dsx1Linestatuslastchange", dsx1Configentry.Dsx1Linestatuslastchange}
    dsx1Configentry.EntityData.Leafs["dsx1LineStatusChangeTrapEnable"] = types.YLeaf{"Dsx1Linestatuschangetrapenable", dsx1Configentry.Dsx1Linestatuschangetrapenable}
    dsx1Configentry.EntityData.Leafs["dsx1LoopbackStatus"] = types.YLeaf{"Dsx1Loopbackstatus", dsx1Configentry.Dsx1Loopbackstatus}
    dsx1Configentry.EntityData.Leafs["dsx1Ds1ChannelNumber"] = types.YLeaf{"Dsx1Ds1Channelnumber", dsx1Configentry.Dsx1Ds1Channelnumber}
    dsx1Configentry.EntityData.Leafs["dsx1Channelization"] = types.YLeaf{"Dsx1Channelization", dsx1Configentry.Dsx1Channelization}
    return &(dsx1Configentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Current table. The type is slice of
    // DS1MIB_Dsx1Currenttable_Dsx1Currententry.
    Dsx1Currententry []DS1MIB_Dsx1Currenttable_Dsx1Currententry
}

func (dsx1Currenttable *DS1MIB_Dsx1Currenttable) GetEntityData() *types.CommonEntityData {
    dsx1Currenttable.EntityData.YFilter = dsx1Currenttable.YFilter
    dsx1Currenttable.EntityData.YangName = "dsx1CurrentTable"
    dsx1Currenttable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Currenttable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Currenttable.EntityData.SegmentPath = "dsx1CurrentTable"
    dsx1Currenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Currenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Currenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Currenttable.EntityData.Children = make(map[string]types.YChild)
    dsx1Currenttable.EntityData.Children["dsx1CurrentEntry"] = types.YChild{"Dsx1Currententry", nil}
    for i := range dsx1Currenttable.Dsx1Currententry {
        dsx1Currenttable.EntityData.Children[types.GetSegmentPath(&dsx1Currenttable.Dsx1Currententry[i])] = types.YChild{"Dsx1Currententry", &dsx1Currenttable.Dsx1Currententry[i]}
    }
    dsx1Currenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Currenttable.EntityData)
}

// DS1MIB_Dsx1Currenttable_Dsx1Currententry
// An entry in the DS1 Current table.
type DS1MIB_Dsx1Currenttable_Dsx1Currententry struct {
    EntityData types.CommonEntityData
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

func (dsx1Currententry *DS1MIB_Dsx1Currenttable_Dsx1Currententry) GetEntityData() *types.CommonEntityData {
    dsx1Currententry.EntityData.YFilter = dsx1Currententry.YFilter
    dsx1Currententry.EntityData.YangName = "dsx1CurrentEntry"
    dsx1Currententry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Currententry.EntityData.ParentYangName = "dsx1CurrentTable"
    dsx1Currententry.EntityData.SegmentPath = "dsx1CurrentEntry" + "[dsx1CurrentIndex='" + fmt.Sprintf("%v", dsx1Currententry.Dsx1Currentindex) + "']"
    dsx1Currententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Currententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Currententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Currententry.EntityData.Children = make(map[string]types.YChild)
    dsx1Currententry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Currententry.EntityData.Leafs["dsx1CurrentIndex"] = types.YLeaf{"Dsx1Currentindex", dsx1Currententry.Dsx1Currentindex}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentESs"] = types.YLeaf{"Dsx1Currentess", dsx1Currententry.Dsx1Currentess}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentSESs"] = types.YLeaf{"Dsx1Currentsess", dsx1Currententry.Dsx1Currentsess}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentSEFSs"] = types.YLeaf{"Dsx1Currentsefss", dsx1Currententry.Dsx1Currentsefss}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentUASs"] = types.YLeaf{"Dsx1Currentuass", dsx1Currententry.Dsx1Currentuass}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentCSSs"] = types.YLeaf{"Dsx1Currentcsss", dsx1Currententry.Dsx1Currentcsss}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentPCVs"] = types.YLeaf{"Dsx1Currentpcvs", dsx1Currententry.Dsx1Currentpcvs}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentLESs"] = types.YLeaf{"Dsx1Currentless", dsx1Currententry.Dsx1Currentless}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentBESs"] = types.YLeaf{"Dsx1Currentbess", dsx1Currententry.Dsx1Currentbess}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentDMs"] = types.YLeaf{"Dsx1Currentdms", dsx1Currententry.Dsx1Currentdms}
    dsx1Currententry.EntityData.Leafs["dsx1CurrentLCVs"] = types.YLeaf{"Dsx1Currentlcvs", dsx1Currententry.Dsx1Currentlcvs}
    return &(dsx1Currententry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Interval table. The type is slice of
    // DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry.
    Dsx1Intervalentry []DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry
}

func (dsx1Intervaltable *DS1MIB_Dsx1Intervaltable) GetEntityData() *types.CommonEntityData {
    dsx1Intervaltable.EntityData.YFilter = dsx1Intervaltable.YFilter
    dsx1Intervaltable.EntityData.YangName = "dsx1IntervalTable"
    dsx1Intervaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Intervaltable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Intervaltable.EntityData.SegmentPath = "dsx1IntervalTable"
    dsx1Intervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Intervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Intervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Intervaltable.EntityData.Children = make(map[string]types.YChild)
    dsx1Intervaltable.EntityData.Children["dsx1IntervalEntry"] = types.YChild{"Dsx1Intervalentry", nil}
    for i := range dsx1Intervaltable.Dsx1Intervalentry {
        dsx1Intervaltable.EntityData.Children[types.GetSegmentPath(&dsx1Intervaltable.Dsx1Intervalentry[i])] = types.YChild{"Dsx1Intervalentry", &dsx1Intervaltable.Dsx1Intervalentry[i]}
    }
    dsx1Intervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Intervaltable.EntityData)
}

// DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry
// An entry in the DS1 Interval table.
type DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Intervalentry *DS1MIB_Dsx1Intervaltable_Dsx1Intervalentry) GetEntityData() *types.CommonEntityData {
    dsx1Intervalentry.EntityData.YFilter = dsx1Intervalentry.YFilter
    dsx1Intervalentry.EntityData.YangName = "dsx1IntervalEntry"
    dsx1Intervalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Intervalentry.EntityData.ParentYangName = "dsx1IntervalTable"
    dsx1Intervalentry.EntityData.SegmentPath = "dsx1IntervalEntry" + "[dsx1IntervalIndex='" + fmt.Sprintf("%v", dsx1Intervalentry.Dsx1Intervalindex) + "']" + "[dsx1IntervalNumber='" + fmt.Sprintf("%v", dsx1Intervalentry.Dsx1Intervalnumber) + "']"
    dsx1Intervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Intervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Intervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Intervalentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Intervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalIndex"] = types.YLeaf{"Dsx1Intervalindex", dsx1Intervalentry.Dsx1Intervalindex}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalNumber"] = types.YLeaf{"Dsx1Intervalnumber", dsx1Intervalentry.Dsx1Intervalnumber}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalESs"] = types.YLeaf{"Dsx1Intervaless", dsx1Intervalentry.Dsx1Intervaless}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalSESs"] = types.YLeaf{"Dsx1Intervalsess", dsx1Intervalentry.Dsx1Intervalsess}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalSEFSs"] = types.YLeaf{"Dsx1Intervalsefss", dsx1Intervalentry.Dsx1Intervalsefss}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalUASs"] = types.YLeaf{"Dsx1Intervaluass", dsx1Intervalentry.Dsx1Intervaluass}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalCSSs"] = types.YLeaf{"Dsx1Intervalcsss", dsx1Intervalentry.Dsx1Intervalcsss}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalPCVs"] = types.YLeaf{"Dsx1Intervalpcvs", dsx1Intervalentry.Dsx1Intervalpcvs}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalLESs"] = types.YLeaf{"Dsx1Intervalless", dsx1Intervalentry.Dsx1Intervalless}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalBESs"] = types.YLeaf{"Dsx1Intervalbess", dsx1Intervalentry.Dsx1Intervalbess}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalDMs"] = types.YLeaf{"Dsx1Intervaldms", dsx1Intervalentry.Dsx1Intervaldms}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalLCVs"] = types.YLeaf{"Dsx1Intervallcvs", dsx1Intervalentry.Dsx1Intervallcvs}
    dsx1Intervalentry.EntityData.Leafs["dsx1IntervalValidData"] = types.YLeaf{"Dsx1Intervalvaliddata", dsx1Intervalentry.Dsx1Intervalvaliddata}
    return &(dsx1Intervalentry.EntityData)
}

// DS1MIB_Dsx1Totaltable
// The DS1 Total Table contains the cumulative sum
// of the various statistics for the 24 hour period
// preceding the current interval.
type DS1MIB_Dsx1Totaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Total table. The type is slice of
    // DS1MIB_Dsx1Totaltable_Dsx1Totalentry.
    Dsx1Totalentry []DS1MIB_Dsx1Totaltable_Dsx1Totalentry
}

func (dsx1Totaltable *DS1MIB_Dsx1Totaltable) GetEntityData() *types.CommonEntityData {
    dsx1Totaltable.EntityData.YFilter = dsx1Totaltable.YFilter
    dsx1Totaltable.EntityData.YangName = "dsx1TotalTable"
    dsx1Totaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Totaltable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Totaltable.EntityData.SegmentPath = "dsx1TotalTable"
    dsx1Totaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Totaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Totaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Totaltable.EntityData.Children = make(map[string]types.YChild)
    dsx1Totaltable.EntityData.Children["dsx1TotalEntry"] = types.YChild{"Dsx1Totalentry", nil}
    for i := range dsx1Totaltable.Dsx1Totalentry {
        dsx1Totaltable.EntityData.Children[types.GetSegmentPath(&dsx1Totaltable.Dsx1Totalentry[i])] = types.YChild{"Dsx1Totalentry", &dsx1Totaltable.Dsx1Totalentry[i]}
    }
    dsx1Totaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Totaltable.EntityData)
}

// DS1MIB_Dsx1Totaltable_Dsx1Totalentry
// An entry in the DS1 Total table.
type DS1MIB_Dsx1Totaltable_Dsx1Totalentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Totalentry *DS1MIB_Dsx1Totaltable_Dsx1Totalentry) GetEntityData() *types.CommonEntityData {
    dsx1Totalentry.EntityData.YFilter = dsx1Totalentry.YFilter
    dsx1Totalentry.EntityData.YangName = "dsx1TotalEntry"
    dsx1Totalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Totalentry.EntityData.ParentYangName = "dsx1TotalTable"
    dsx1Totalentry.EntityData.SegmentPath = "dsx1TotalEntry" + "[dsx1TotalIndex='" + fmt.Sprintf("%v", dsx1Totalentry.Dsx1Totalindex) + "']"
    dsx1Totalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Totalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Totalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Totalentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Totalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Totalentry.EntityData.Leafs["dsx1TotalIndex"] = types.YLeaf{"Dsx1Totalindex", dsx1Totalentry.Dsx1Totalindex}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalESs"] = types.YLeaf{"Dsx1Totaless", dsx1Totalentry.Dsx1Totaless}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalSESs"] = types.YLeaf{"Dsx1Totalsess", dsx1Totalentry.Dsx1Totalsess}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalSEFSs"] = types.YLeaf{"Dsx1Totalsefss", dsx1Totalentry.Dsx1Totalsefss}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalUASs"] = types.YLeaf{"Dsx1Totaluass", dsx1Totalentry.Dsx1Totaluass}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalCSSs"] = types.YLeaf{"Dsx1Totalcsss", dsx1Totalentry.Dsx1Totalcsss}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalPCVs"] = types.YLeaf{"Dsx1Totalpcvs", dsx1Totalentry.Dsx1Totalpcvs}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalLESs"] = types.YLeaf{"Dsx1Totalless", dsx1Totalentry.Dsx1Totalless}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalBESs"] = types.YLeaf{"Dsx1Totalbess", dsx1Totalentry.Dsx1Totalbess}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalDMs"] = types.YLeaf{"Dsx1Totaldms", dsx1Totalentry.Dsx1Totaldms}
    dsx1Totalentry.EntityData.Leafs["dsx1TotalLCVs"] = types.YLeaf{"Dsx1Totallcvs", dsx1Totalentry.Dsx1Totallcvs}
    return &(dsx1Totalentry.EntityData)
}

// DS1MIB_Dsx1Farendcurrenttable
// The DS1 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end messages on the Facilities Data
// Link.  The definitions are the same as described
// for the near-end information.
type DS1MIB_Dsx1Farendcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Current table. The type is slice of
    // DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry.
    Dsx1Farendcurrententry []DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry
}

func (dsx1Farendcurrenttable *DS1MIB_Dsx1Farendcurrenttable) GetEntityData() *types.CommonEntityData {
    dsx1Farendcurrenttable.EntityData.YFilter = dsx1Farendcurrenttable.YFilter
    dsx1Farendcurrenttable.EntityData.YangName = "dsx1FarEndCurrentTable"
    dsx1Farendcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendcurrenttable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Farendcurrenttable.EntityData.SegmentPath = "dsx1FarEndCurrentTable"
    dsx1Farendcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendcurrenttable.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendcurrenttable.EntityData.Children["dsx1FarEndCurrentEntry"] = types.YChild{"Dsx1Farendcurrententry", nil}
    for i := range dsx1Farendcurrenttable.Dsx1Farendcurrententry {
        dsx1Farendcurrenttable.EntityData.Children[types.GetSegmentPath(&dsx1Farendcurrenttable.Dsx1Farendcurrententry[i])] = types.YChild{"Dsx1Farendcurrententry", &dsx1Farendcurrenttable.Dsx1Farendcurrententry[i]}
    }
    dsx1Farendcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Farendcurrenttable.EntityData)
}

// DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry
// An entry in the DS1 Far End Current table.
type DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry struct {
    EntityData types.CommonEntityData
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

func (dsx1Farendcurrententry *DS1MIB_Dsx1Farendcurrenttable_Dsx1Farendcurrententry) GetEntityData() *types.CommonEntityData {
    dsx1Farendcurrententry.EntityData.YFilter = dsx1Farendcurrententry.YFilter
    dsx1Farendcurrententry.EntityData.YangName = "dsx1FarEndCurrentEntry"
    dsx1Farendcurrententry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendcurrententry.EntityData.ParentYangName = "dsx1FarEndCurrentTable"
    dsx1Farendcurrententry.EntityData.SegmentPath = "dsx1FarEndCurrentEntry" + "[dsx1FarEndCurrentIndex='" + fmt.Sprintf("%v", dsx1Farendcurrententry.Dsx1Farendcurrentindex) + "']"
    dsx1Farendcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendcurrententry.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentIndex"] = types.YLeaf{"Dsx1Farendcurrentindex", dsx1Farendcurrententry.Dsx1Farendcurrentindex}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndTimeElapsed"] = types.YLeaf{"Dsx1Farendtimeelapsed", dsx1Farendcurrententry.Dsx1Farendtimeelapsed}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndValidIntervals"] = types.YLeaf{"Dsx1Farendvalidintervals", dsx1Farendcurrententry.Dsx1Farendvalidintervals}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentESs"] = types.YLeaf{"Dsx1Farendcurrentess", dsx1Farendcurrententry.Dsx1Farendcurrentess}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentSESs"] = types.YLeaf{"Dsx1Farendcurrentsess", dsx1Farendcurrententry.Dsx1Farendcurrentsess}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentSEFSs"] = types.YLeaf{"Dsx1Farendcurrentsefss", dsx1Farendcurrententry.Dsx1Farendcurrentsefss}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentUASs"] = types.YLeaf{"Dsx1Farendcurrentuass", dsx1Farendcurrententry.Dsx1Farendcurrentuass}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentCSSs"] = types.YLeaf{"Dsx1Farendcurrentcsss", dsx1Farendcurrententry.Dsx1Farendcurrentcsss}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentLESs"] = types.YLeaf{"Dsx1Farendcurrentless", dsx1Farendcurrententry.Dsx1Farendcurrentless}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentPCVs"] = types.YLeaf{"Dsx1Farendcurrentpcvs", dsx1Farendcurrententry.Dsx1Farendcurrentpcvs}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentBESs"] = types.YLeaf{"Dsx1Farendcurrentbess", dsx1Farendcurrententry.Dsx1Farendcurrentbess}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndCurrentDMs"] = types.YLeaf{"Dsx1Farendcurrentdms", dsx1Farendcurrententry.Dsx1Farendcurrentdms}
    dsx1Farendcurrententry.EntityData.Leafs["dsx1FarEndInvalidIntervals"] = types.YLeaf{"Dsx1Farendinvalidintervals", dsx1Farendcurrententry.Dsx1Farendinvalidintervals}
    return &(dsx1Farendcurrententry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Interval table. The type is slice of
    // DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry.
    Dsx1Farendintervalentry []DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry
}

func (dsx1Farendintervaltable *DS1MIB_Dsx1Farendintervaltable) GetEntityData() *types.CommonEntityData {
    dsx1Farendintervaltable.EntityData.YFilter = dsx1Farendintervaltable.YFilter
    dsx1Farendintervaltable.EntityData.YangName = "dsx1FarEndIntervalTable"
    dsx1Farendintervaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendintervaltable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Farendintervaltable.EntityData.SegmentPath = "dsx1FarEndIntervalTable"
    dsx1Farendintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendintervaltable.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendintervaltable.EntityData.Children["dsx1FarEndIntervalEntry"] = types.YChild{"Dsx1Farendintervalentry", nil}
    for i := range dsx1Farendintervaltable.Dsx1Farendintervalentry {
        dsx1Farendintervaltable.EntityData.Children[types.GetSegmentPath(&dsx1Farendintervaltable.Dsx1Farendintervalentry[i])] = types.YChild{"Dsx1Farendintervalentry", &dsx1Farendintervaltable.Dsx1Farendintervalentry[i]}
    }
    dsx1Farendintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Farendintervaltable.EntityData)
}

// DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry
// An entry in the DS1 Far End Interval table.
type DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Farendintervalentry *DS1MIB_Dsx1Farendintervaltable_Dsx1Farendintervalentry) GetEntityData() *types.CommonEntityData {
    dsx1Farendintervalentry.EntityData.YFilter = dsx1Farendintervalentry.YFilter
    dsx1Farendintervalentry.EntityData.YangName = "dsx1FarEndIntervalEntry"
    dsx1Farendintervalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendintervalentry.EntityData.ParentYangName = "dsx1FarEndIntervalTable"
    dsx1Farendintervalentry.EntityData.SegmentPath = "dsx1FarEndIntervalEntry" + "[dsx1FarEndIntervalIndex='" + fmt.Sprintf("%v", dsx1Farendintervalentry.Dsx1Farendintervalindex) + "']" + "[dsx1FarEndIntervalNumber='" + fmt.Sprintf("%v", dsx1Farendintervalentry.Dsx1Farendintervalnumber) + "']"
    dsx1Farendintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendintervalentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalIndex"] = types.YLeaf{"Dsx1Farendintervalindex", dsx1Farendintervalentry.Dsx1Farendintervalindex}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalNumber"] = types.YLeaf{"Dsx1Farendintervalnumber", dsx1Farendintervalentry.Dsx1Farendintervalnumber}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalESs"] = types.YLeaf{"Dsx1Farendintervaless", dsx1Farendintervalentry.Dsx1Farendintervaless}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalSESs"] = types.YLeaf{"Dsx1Farendintervalsess", dsx1Farendintervalentry.Dsx1Farendintervalsess}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalSEFSs"] = types.YLeaf{"Dsx1Farendintervalsefss", dsx1Farendintervalentry.Dsx1Farendintervalsefss}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalUASs"] = types.YLeaf{"Dsx1Farendintervaluass", dsx1Farendintervalentry.Dsx1Farendintervaluass}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalCSSs"] = types.YLeaf{"Dsx1Farendintervalcsss", dsx1Farendintervalentry.Dsx1Farendintervalcsss}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalLESs"] = types.YLeaf{"Dsx1Farendintervalless", dsx1Farendintervalentry.Dsx1Farendintervalless}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalPCVs"] = types.YLeaf{"Dsx1Farendintervalpcvs", dsx1Farendintervalentry.Dsx1Farendintervalpcvs}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalBESs"] = types.YLeaf{"Dsx1Farendintervalbess", dsx1Farendintervalentry.Dsx1Farendintervalbess}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalDMs"] = types.YLeaf{"Dsx1Farendintervaldms", dsx1Farendintervalentry.Dsx1Farendintervaldms}
    dsx1Farendintervalentry.EntityData.Leafs["dsx1FarEndIntervalValidData"] = types.YLeaf{"Dsx1Farendintervalvaliddata", dsx1Farendintervalentry.Dsx1Farendintervalvaliddata}
    return &(dsx1Farendintervalentry.EntityData)
}

// DS1MIB_Dsx1Farendtotaltable
// The DS1 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS1MIB_Dsx1Farendtotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Total table. The type is slice of
    // DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry.
    Dsx1Farendtotalentry []DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry
}

func (dsx1Farendtotaltable *DS1MIB_Dsx1Farendtotaltable) GetEntityData() *types.CommonEntityData {
    dsx1Farendtotaltable.EntityData.YFilter = dsx1Farendtotaltable.YFilter
    dsx1Farendtotaltable.EntityData.YangName = "dsx1FarEndTotalTable"
    dsx1Farendtotaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendtotaltable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Farendtotaltable.EntityData.SegmentPath = "dsx1FarEndTotalTable"
    dsx1Farendtotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendtotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendtotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendtotaltable.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendtotaltable.EntityData.Children["dsx1FarEndTotalEntry"] = types.YChild{"Dsx1Farendtotalentry", nil}
    for i := range dsx1Farendtotaltable.Dsx1Farendtotalentry {
        dsx1Farendtotaltable.EntityData.Children[types.GetSegmentPath(&dsx1Farendtotaltable.Dsx1Farendtotalentry[i])] = types.YChild{"Dsx1Farendtotalentry", &dsx1Farendtotaltable.Dsx1Farendtotalentry[i]}
    }
    dsx1Farendtotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Farendtotaltable.EntityData)
}

// DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry
// An entry in the DS1 Far End Total table.
type DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Farendtotalentry *DS1MIB_Dsx1Farendtotaltable_Dsx1Farendtotalentry) GetEntityData() *types.CommonEntityData {
    dsx1Farendtotalentry.EntityData.YFilter = dsx1Farendtotalentry.YFilter
    dsx1Farendtotalentry.EntityData.YangName = "dsx1FarEndTotalEntry"
    dsx1Farendtotalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Farendtotalentry.EntityData.ParentYangName = "dsx1FarEndTotalTable"
    dsx1Farendtotalentry.EntityData.SegmentPath = "dsx1FarEndTotalEntry" + "[dsx1FarEndTotalIndex='" + fmt.Sprintf("%v", dsx1Farendtotalentry.Dsx1Farendtotalindex) + "']"
    dsx1Farendtotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Farendtotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Farendtotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Farendtotalentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Farendtotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalIndex"] = types.YLeaf{"Dsx1Farendtotalindex", dsx1Farendtotalentry.Dsx1Farendtotalindex}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalESs"] = types.YLeaf{"Dsx1Farendtotaless", dsx1Farendtotalentry.Dsx1Farendtotaless}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalSESs"] = types.YLeaf{"Dsx1Farendtotalsess", dsx1Farendtotalentry.Dsx1Farendtotalsess}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalSEFSs"] = types.YLeaf{"Dsx1Farendtotalsefss", dsx1Farendtotalentry.Dsx1Farendtotalsefss}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalUASs"] = types.YLeaf{"Dsx1Farendtotaluass", dsx1Farendtotalentry.Dsx1Farendtotaluass}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalCSSs"] = types.YLeaf{"Dsx1Farendtotalcsss", dsx1Farendtotalentry.Dsx1Farendtotalcsss}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalLESs"] = types.YLeaf{"Dsx1Farendtotalless", dsx1Farendtotalentry.Dsx1Farendtotalless}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalPCVs"] = types.YLeaf{"Dsx1Farendtotalpcvs", dsx1Farendtotalentry.Dsx1Farendtotalpcvs}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalBESs"] = types.YLeaf{"Dsx1Farendtotalbess", dsx1Farendtotalentry.Dsx1Farendtotalbess}
    dsx1Farendtotalentry.EntityData.Leafs["dsx1FarEndTotalDMs"] = types.YLeaf{"Dsx1Farendtotaldms", dsx1Farendtotalentry.Dsx1Farendtotaldms}
    return &(dsx1Farendtotalentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Fractional table. The type is slice of
    // DS1MIB_Dsx1Fractable_Dsx1Fracentry.
    Dsx1Fracentry []DS1MIB_Dsx1Fractable_Dsx1Fracentry
}

func (dsx1Fractable *DS1MIB_Dsx1Fractable) GetEntityData() *types.CommonEntityData {
    dsx1Fractable.EntityData.YFilter = dsx1Fractable.YFilter
    dsx1Fractable.EntityData.YangName = "dsx1FracTable"
    dsx1Fractable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Fractable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Fractable.EntityData.SegmentPath = "dsx1FracTable"
    dsx1Fractable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Fractable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Fractable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Fractable.EntityData.Children = make(map[string]types.YChild)
    dsx1Fractable.EntityData.Children["dsx1FracEntry"] = types.YChild{"Dsx1Fracentry", nil}
    for i := range dsx1Fractable.Dsx1Fracentry {
        dsx1Fractable.EntityData.Children[types.GetSegmentPath(&dsx1Fractable.Dsx1Fracentry[i])] = types.YChild{"Dsx1Fracentry", &dsx1Fractable.Dsx1Fracentry[i]}
    }
    dsx1Fractable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Fractable.EntityData)
}

// DS1MIB_Dsx1Fractable_Dsx1Fracentry
// An entry in the DS1 Fractional table.
type DS1MIB_Dsx1Fractable_Dsx1Fracentry struct {
    EntityData types.CommonEntityData
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

func (dsx1Fracentry *DS1MIB_Dsx1Fractable_Dsx1Fracentry) GetEntityData() *types.CommonEntityData {
    dsx1Fracentry.EntityData.YFilter = dsx1Fracentry.YFilter
    dsx1Fracentry.EntityData.YangName = "dsx1FracEntry"
    dsx1Fracentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Fracentry.EntityData.ParentYangName = "dsx1FracTable"
    dsx1Fracentry.EntityData.SegmentPath = "dsx1FracEntry" + "[dsx1FracIndex='" + fmt.Sprintf("%v", dsx1Fracentry.Dsx1Fracindex) + "']" + "[dsx1FracNumber='" + fmt.Sprintf("%v", dsx1Fracentry.Dsx1Fracnumber) + "']"
    dsx1Fracentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Fracentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Fracentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Fracentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Fracentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Fracentry.EntityData.Leafs["dsx1FracIndex"] = types.YLeaf{"Dsx1Fracindex", dsx1Fracentry.Dsx1Fracindex}
    dsx1Fracentry.EntityData.Leafs["dsx1FracNumber"] = types.YLeaf{"Dsx1Fracnumber", dsx1Fracentry.Dsx1Fracnumber}
    dsx1Fracentry.EntityData.Leafs["dsx1FracIfIndex"] = types.YLeaf{"Dsx1Fracifindex", dsx1Fracentry.Dsx1Fracifindex}
    return &(dsx1Fracentry.EntityData)
}

// DS1MIB_Dsx1Chanmappingtable
// The DS1 Channel Mapping table.  This table maps a
// DS1 channel number on a particular DS3 into an
// ifIndex.  In the presence of DS2s, this table can
// be used to map a DS2 channel number on a DS3 into
// an ifIndex, or used to map a DS1 channel number on
// a DS2 onto an ifIndex.
type DS1MIB_Dsx1Chanmappingtable struct {
    EntityData types.CommonEntityData
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

func (dsx1Chanmappingtable *DS1MIB_Dsx1Chanmappingtable) GetEntityData() *types.CommonEntityData {
    dsx1Chanmappingtable.EntityData.YFilter = dsx1Chanmappingtable.YFilter
    dsx1Chanmappingtable.EntityData.YangName = "dsx1ChanMappingTable"
    dsx1Chanmappingtable.EntityData.BundleName = "cisco_ios_xe"
    dsx1Chanmappingtable.EntityData.ParentYangName = "DS1-MIB"
    dsx1Chanmappingtable.EntityData.SegmentPath = "dsx1ChanMappingTable"
    dsx1Chanmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Chanmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Chanmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Chanmappingtable.EntityData.Children = make(map[string]types.YChild)
    dsx1Chanmappingtable.EntityData.Children["dsx1ChanMappingEntry"] = types.YChild{"Dsx1Chanmappingentry", nil}
    for i := range dsx1Chanmappingtable.Dsx1Chanmappingentry {
        dsx1Chanmappingtable.EntityData.Children[types.GetSegmentPath(&dsx1Chanmappingtable.Dsx1Chanmappingentry[i])] = types.YChild{"Dsx1Chanmappingentry", &dsx1Chanmappingtable.Dsx1Chanmappingentry[i]}
    }
    dsx1Chanmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx1Chanmappingtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (dsx1Chanmappingentry *DS1MIB_Dsx1Chanmappingtable_Dsx1Chanmappingentry) GetEntityData() *types.CommonEntityData {
    dsx1Chanmappingentry.EntityData.YFilter = dsx1Chanmappingentry.YFilter
    dsx1Chanmappingentry.EntityData.YangName = "dsx1ChanMappingEntry"
    dsx1Chanmappingentry.EntityData.BundleName = "cisco_ios_xe"
    dsx1Chanmappingentry.EntityData.ParentYangName = "dsx1ChanMappingTable"
    dsx1Chanmappingentry.EntityData.SegmentPath = "dsx1ChanMappingEntry" + "[ifIndex='" + fmt.Sprintf("%v", dsx1Chanmappingentry.Ifindex) + "']" + "[dsx1Ds1ChannelNumber='" + fmt.Sprintf("%v", dsx1Chanmappingentry.Dsx1Ds1Channelnumber) + "']"
    dsx1Chanmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1Chanmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1Chanmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1Chanmappingentry.EntityData.Children = make(map[string]types.YChild)
    dsx1Chanmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx1Chanmappingentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", dsx1Chanmappingentry.Ifindex}
    dsx1Chanmappingentry.EntityData.Leafs["dsx1Ds1ChannelNumber"] = types.YLeaf{"Dsx1Ds1Channelnumber", dsx1Chanmappingentry.Dsx1Ds1Channelnumber}
    dsx1Chanmappingentry.EntityData.Leafs["dsx1ChanMappedIfIndex"] = types.YLeaf{"Dsx1Chanmappedifindex", dsx1Chanmappingentry.Dsx1Chanmappedifindex}
    return &(dsx1Chanmappingentry.EntityData)
}

