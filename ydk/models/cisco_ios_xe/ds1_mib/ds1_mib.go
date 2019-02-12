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
    Dsx1ConfigTable DS1MIB_Dsx1ConfigTable

    // The DS1 current table contains various statistics being collected for the
    // current 15 minute interval.
    Dsx1CurrentTable DS1MIB_Dsx1CurrentTable

    // The DS1 Interval Table contains various statistics collected by each DS1
    // Interface over the previous 24 hours of operation.  The past 24 hours are
    // broken into 96 completed 15 minute intervals.  Each row in this table
    // represents one such interval (identified by dsx1IntervalNumber) for one
    // specific instance (identified by dsx1IntervalIndex).
    Dsx1IntervalTable DS1MIB_Dsx1IntervalTable

    // The DS1 Total Table contains the cumulative sum of the various statistics
    // for the 24 hour period preceding the current interval.
    Dsx1TotalTable DS1MIB_Dsx1TotalTable

    // The DS1 Far End Current table contains various statistics being collected
    // for the current 15 minute interval.  The statistics are collected from the
    // far end messages on the Facilities Data Link.  The definitions are the same
    // as described for the near-end information.
    Dsx1FarEndCurrentTable DS1MIB_Dsx1FarEndCurrentTable

    // The DS1 Far End Interval Table contains various statistics collected by
    // each DS1 interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals. Each row in this
    // table represents one such interval (identified by dsx1FarEndIntervalNumber)
    // for one specific instance (identified by dsx1FarEndIntervalIndex).
    Dsx1FarEndIntervalTable DS1MIB_Dsx1FarEndIntervalTable

    // The DS1 Far End Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx1FarEndTotalTable DS1MIB_Dsx1FarEndTotalTable

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
    Dsx1FracTable DS1MIB_Dsx1FracTable

    // The DS1 Channel Mapping table.  This table maps a DS1 channel number on a
    // particular DS3 into an ifIndex.  In the presence of DS2s, this table can be
    // used to map a DS2 channel number on a DS3 into an ifIndex, or used to map a
    // DS1 channel number on a DS2 onto an ifIndex.
    Dsx1ChanMappingTable DS1MIB_Dsx1ChanMappingTable
}

func (dS1MIB *DS1MIB) GetEntityData() *types.CommonEntityData {
    dS1MIB.EntityData.YFilter = dS1MIB.YFilter
    dS1MIB.EntityData.YangName = "DS1-MIB"
    dS1MIB.EntityData.BundleName = "cisco_ios_xe"
    dS1MIB.EntityData.ParentYangName = "DS1-MIB"
    dS1MIB.EntityData.SegmentPath = "DS1-MIB:DS1-MIB"
    dS1MIB.EntityData.AbsolutePath = dS1MIB.EntityData.SegmentPath
    dS1MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dS1MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dS1MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dS1MIB.EntityData.Children = types.NewOrderedMap()
    dS1MIB.EntityData.Children.Append("dsx1ConfigTable", types.YChild{"Dsx1ConfigTable", &dS1MIB.Dsx1ConfigTable})
    dS1MIB.EntityData.Children.Append("dsx1CurrentTable", types.YChild{"Dsx1CurrentTable", &dS1MIB.Dsx1CurrentTable})
    dS1MIB.EntityData.Children.Append("dsx1IntervalTable", types.YChild{"Dsx1IntervalTable", &dS1MIB.Dsx1IntervalTable})
    dS1MIB.EntityData.Children.Append("dsx1TotalTable", types.YChild{"Dsx1TotalTable", &dS1MIB.Dsx1TotalTable})
    dS1MIB.EntityData.Children.Append("dsx1FarEndCurrentTable", types.YChild{"Dsx1FarEndCurrentTable", &dS1MIB.Dsx1FarEndCurrentTable})
    dS1MIB.EntityData.Children.Append("dsx1FarEndIntervalTable", types.YChild{"Dsx1FarEndIntervalTable", &dS1MIB.Dsx1FarEndIntervalTable})
    dS1MIB.EntityData.Children.Append("dsx1FarEndTotalTable", types.YChild{"Dsx1FarEndTotalTable", &dS1MIB.Dsx1FarEndTotalTable})
    dS1MIB.EntityData.Children.Append("dsx1FracTable", types.YChild{"Dsx1FracTable", &dS1MIB.Dsx1FracTable})
    dS1MIB.EntityData.Children.Append("dsx1ChanMappingTable", types.YChild{"Dsx1ChanMappingTable", &dS1MIB.Dsx1ChanMappingTable})
    dS1MIB.EntityData.Leafs = types.NewOrderedMap()

    dS1MIB.EntityData.YListKeys = []string {}

    return &(dS1MIB.EntityData)
}

// DS1MIB_Dsx1ConfigTable
// The DS1 Configuration table.
type DS1MIB_Dsx1ConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Configuration table. The type is slice of
    // DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry.
    Dsx1ConfigEntry []*DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry
}

func (dsx1ConfigTable *DS1MIB_Dsx1ConfigTable) GetEntityData() *types.CommonEntityData {
    dsx1ConfigTable.EntityData.YFilter = dsx1ConfigTable.YFilter
    dsx1ConfigTable.EntityData.YangName = "dsx1ConfigTable"
    dsx1ConfigTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1ConfigTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1ConfigTable.EntityData.SegmentPath = "dsx1ConfigTable"
    dsx1ConfigTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1ConfigTable.EntityData.SegmentPath
    dsx1ConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1ConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1ConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1ConfigTable.EntityData.Children = types.NewOrderedMap()
    dsx1ConfigTable.EntityData.Children.Append("dsx1ConfigEntry", types.YChild{"Dsx1ConfigEntry", nil})
    for i := range dsx1ConfigTable.Dsx1ConfigEntry {
        dsx1ConfigTable.EntityData.Children.Append(types.GetSegmentPath(dsx1ConfigTable.Dsx1ConfigEntry[i]), types.YChild{"Dsx1ConfigEntry", dsx1ConfigTable.Dsx1ConfigEntry[i]})
    }
    dsx1ConfigTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1ConfigTable.EntityData.YListKeys = []string {}

    return &(dsx1ConfigTable.EntityData)
}

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry
// An entry in the DS1 Configuration table.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    Dsx1LineIndex interface{}

    // This value for this object is equal to the value of ifIndex from the
    // Interfaces table of MIB II (RFC 1213). The type is interface{} with range:
    // 1..2147483647.
    Dsx1IfIndex interface{}

    // The number of seconds that have elapsed since the beginning of the near end
    // current error- measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx1TimeElapsed interface{}

    // The number of previous near end intervals for which data was collected. 
    // The value will be 96 unless the interface was brought online within the
    // last 24 hours, in which case the value will be the number of complete 15
    // minute near end intervals since the interface has been online.  In the case
    // where the agent is a proxy, it is possible that some intervals are
    // unavailable.  In this case, this interval is the maximum interval number
    // for which data is available. The type is interface{} with range: 0..96.
    Dsx1ValidIntervals interface{}

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
    // Dsx1LineType.
    Dsx1LineType interface{}

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
    // zero bits.  Used for DS2. The type is Dsx1LineCoding.
    Dsx1LineCoding interface{}

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
    // Dsx1SendCode.
    Dsx1SendCode interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. The type is string with
    // length: 0..255.
    Dsx1CircuitIdentifier interface{}

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
    // be active simultaneously. The type is Dsx1LoopbackConfig.
    Dsx1LoopbackConfig interface{}

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
    Dsx1LineStatus interface{}

    // 'none' indicates that no bits are reserved for signaling on this channel. 
    // 'robbedBit' indicates that DS1 Robbed Bit  Sig- naling is in use. 
    // 'bitOriented' indicates that E1 Channel  Asso- ciated Signaling is in use. 
    // 'messageOriented' indicates that Common  Chan- nel Signaling is in use
    // either on channel 16 of an E1 link or channel 24 of a DS1. The type is
    // Dsx1SignalMode.
    Dsx1SignalMode interface{}

    // The source of Transmit Clock. 'loopTiming' indicates that the recovered re-
    // ceive clock is used as the transmit clock.  'localTiming' indicates that a
    // local clock source is used or when an external clock is attached to the box
    // containing the interface.  'throughTiming' indicates that recovered re-
    // ceive clock from another interface is used as the transmit clock. The type
    // is Dsx1TransmitClockSource.
    Dsx1TransmitClockSource interface{}

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
    Dsx1InvalidIntervals interface{}

    // The length of the ds1 line in meters. This objects provides information for
    // line build out circuitry.  This object is only useful if the interface has
    // configurable line build out circuitry. The type is interface{} with range:
    // 0..64000. Units are meters.
    Dsx1LineLength interface{}

    // The value of MIB II's sysUpTime object at the time this DS1 entered its
    // current line status state.  If the current state was entered prior to the
    // last re-initialization of the proxy-agent, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    Dsx1LineStatusLastChange interface{}

    // Indicates whether dsx1LineStatusChange traps should be generated for this
    // interface. The type is Dsx1LineStatusChangeTrapEnable.
    Dsx1LineStatusChangeTrapEnable interface{}

    // This variable represents the current state of the loopback on the DS1
    // interface.  It contains information about loopbacks established by a
    // manager and remotely from the far end.  The dsx1LoopbackStatus is a bit map
    // represented as a sum, therefore is can represent multiple loopbacks
    // simultaneously.  The various bit positions are:  1  dsx1NoLoopback  2 
    // dsx1NearEndPayloadLoopback  4  dsx1NearEndLineLoopback  8 
    // dsx1NearEndOtherLoopback 16  dsx1NearEndInwardLoopback 32 
    // dsx1FarEndPayloadLoopback 64  dsx1FarEndLineLoopback. The type is
    // interface{} with range: 1..127.
    Dsx1LoopbackStatus interface{}

    // This variable represents the channel number of the DS1/E1 on its parent
    // Ds2/E2 or DS3/E3.  A value of 0 indicated this DS1/E1 does not have a
    // parent DS3/E3. The type is interface{} with range: 0..28.
    Dsx1Ds1ChannelNumber interface{}

    // Indicates whether this ds1/e1 is channelized or unchannelized.  The value
    // of enabledDs0 indicates that this is a DS1 channelized into DS0s.  The
    // value of enabledDs1 indicated that this is a DS2 channelized into DS1s. 
    // Setting this value will cause the creation or deletion of entries in the
    // ifTable for the DS0s that are within the DS1. The type is
    // Dsx1Channelization.
    Dsx1Channelization interface{}
}

func (dsx1ConfigEntry *DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry) GetEntityData() *types.CommonEntityData {
    dsx1ConfigEntry.EntityData.YFilter = dsx1ConfigEntry.YFilter
    dsx1ConfigEntry.EntityData.YangName = "dsx1ConfigEntry"
    dsx1ConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1ConfigEntry.EntityData.ParentYangName = "dsx1ConfigTable"
    dsx1ConfigEntry.EntityData.SegmentPath = "dsx1ConfigEntry" + types.AddKeyToken(dsx1ConfigEntry.Dsx1LineIndex, "dsx1LineIndex")
    dsx1ConfigEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1ConfigTable/" + dsx1ConfigEntry.EntityData.SegmentPath
    dsx1ConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1ConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1ConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1ConfigEntry.EntityData.Children = types.NewOrderedMap()
    dsx1ConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineIndex", types.YLeaf{"Dsx1LineIndex", dsx1ConfigEntry.Dsx1LineIndex})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1IfIndex", types.YLeaf{"Dsx1IfIndex", dsx1ConfigEntry.Dsx1IfIndex})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1TimeElapsed", types.YLeaf{"Dsx1TimeElapsed", dsx1ConfigEntry.Dsx1TimeElapsed})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1ValidIntervals", types.YLeaf{"Dsx1ValidIntervals", dsx1ConfigEntry.Dsx1ValidIntervals})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineType", types.YLeaf{"Dsx1LineType", dsx1ConfigEntry.Dsx1LineType})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineCoding", types.YLeaf{"Dsx1LineCoding", dsx1ConfigEntry.Dsx1LineCoding})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1SendCode", types.YLeaf{"Dsx1SendCode", dsx1ConfigEntry.Dsx1SendCode})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1CircuitIdentifier", types.YLeaf{"Dsx1CircuitIdentifier", dsx1ConfigEntry.Dsx1CircuitIdentifier})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LoopbackConfig", types.YLeaf{"Dsx1LoopbackConfig", dsx1ConfigEntry.Dsx1LoopbackConfig})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineStatus", types.YLeaf{"Dsx1LineStatus", dsx1ConfigEntry.Dsx1LineStatus})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1SignalMode", types.YLeaf{"Dsx1SignalMode", dsx1ConfigEntry.Dsx1SignalMode})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1TransmitClockSource", types.YLeaf{"Dsx1TransmitClockSource", dsx1ConfigEntry.Dsx1TransmitClockSource})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1Fdl", types.YLeaf{"Dsx1Fdl", dsx1ConfigEntry.Dsx1Fdl})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1InvalidIntervals", types.YLeaf{"Dsx1InvalidIntervals", dsx1ConfigEntry.Dsx1InvalidIntervals})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineLength", types.YLeaf{"Dsx1LineLength", dsx1ConfigEntry.Dsx1LineLength})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineStatusLastChange", types.YLeaf{"Dsx1LineStatusLastChange", dsx1ConfigEntry.Dsx1LineStatusLastChange})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LineStatusChangeTrapEnable", types.YLeaf{"Dsx1LineStatusChangeTrapEnable", dsx1ConfigEntry.Dsx1LineStatusChangeTrapEnable})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1LoopbackStatus", types.YLeaf{"Dsx1LoopbackStatus", dsx1ConfigEntry.Dsx1LoopbackStatus})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1Ds1ChannelNumber", types.YLeaf{"Dsx1Ds1ChannelNumber", dsx1ConfigEntry.Dsx1Ds1ChannelNumber})
    dsx1ConfigEntry.EntityData.Leafs.Append("dsx1Channelization", types.YLeaf{"Dsx1Channelization", dsx1ConfigEntry.Dsx1Channelization})

    dsx1ConfigEntry.EntityData.YListKeys = []string {"Dsx1LineIndex"}

    return &(dsx1ConfigEntry.EntityData)
}

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization represents ifTable for the DS0s that are within the DS1.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization_disabled DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization = "disabled"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization_enabledDs0 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization = "enabledDs0"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization_enabledDs1 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Channelization = "enabledDs1"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding represents for DS2.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1JBZS DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1JBZS"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1B8ZS DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1B8ZS"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1HDB3 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1HDB3"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1ZBTSI DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1ZBTSI"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1AMI DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1AMI"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_other DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "other"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding_dsx1B6ZS DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineCoding = "dsx1B6ZS"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable represents should be generated for this interface.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable_enabled DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable = "enabled"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable_disabled DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineStatusChangeTrapEnable = "disabled"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType represents For further information See ITU-T Recomm G.704
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_other DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "other"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1ESF DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1ESF"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1D4 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1D4"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1E1 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1E1"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1E1CRC DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1E1CRC"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1E1MF DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1E1MF"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1E1CRCMF DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1E1CRCMF"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1Unframed DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1Unframed"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1E1Unframed DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1E1Unframed"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx1DS2M12 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx1DS2M12"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType_dsx2E2 DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LineType = "dsx2E2"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig represents active simultaneously.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1NoLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1NoLoop"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1PayloadLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1PayloadLoop"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1LineLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1LineLoop"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1OtherLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1OtherLoop"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1InwardLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1InwardLoop"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig_dsx1DualLoop DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1LoopbackConfig = "dsx1DualLoop"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode represents described by this object
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendNoCode DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendNoCode"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendLineCode DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendLineCode"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendPayloadCode DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendPayloadCode"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendResetCode DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendResetCode"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendQRS DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendQRS"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1Send511Pattern DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1Send511Pattern"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1Send3in24Pattern DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1Send3in24Pattern"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode_dsx1SendOtherTestPattern DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SendCode = "dsx1SendOtherTestPattern"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode represents an E1 link or channel 24 of a DS1.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode_none DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode = "none"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode_robbedBit DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode = "robbedBit"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode_bitOriented DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode = "bitOriented"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode_messageOriented DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode = "messageOriented"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode_other DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1SignalMode = "other"
)

// DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource represents the transmit clock.
type DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource string

const (
    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource_loopTiming DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource = "loopTiming"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource_localTiming DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource = "localTiming"

    DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource_throughTiming DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1TransmitClockSource = "throughTiming"
)

// DS1MIB_Dsx1CurrentTable
// The DS1 current table contains various statistics
// being collected for the current 15 minute
// interval.
type DS1MIB_Dsx1CurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Current table. The type is slice of
    // DS1MIB_Dsx1CurrentTable_Dsx1CurrentEntry.
    Dsx1CurrentEntry []*DS1MIB_Dsx1CurrentTable_Dsx1CurrentEntry
}

func (dsx1CurrentTable *DS1MIB_Dsx1CurrentTable) GetEntityData() *types.CommonEntityData {
    dsx1CurrentTable.EntityData.YFilter = dsx1CurrentTable.YFilter
    dsx1CurrentTable.EntityData.YangName = "dsx1CurrentTable"
    dsx1CurrentTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1CurrentTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1CurrentTable.EntityData.SegmentPath = "dsx1CurrentTable"
    dsx1CurrentTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1CurrentTable.EntityData.SegmentPath
    dsx1CurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1CurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1CurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1CurrentTable.EntityData.Children = types.NewOrderedMap()
    dsx1CurrentTable.EntityData.Children.Append("dsx1CurrentEntry", types.YChild{"Dsx1CurrentEntry", nil})
    for i := range dsx1CurrentTable.Dsx1CurrentEntry {
        dsx1CurrentTable.EntityData.Children.Append(types.GetSegmentPath(dsx1CurrentTable.Dsx1CurrentEntry[i]), types.YChild{"Dsx1CurrentEntry", dsx1CurrentTable.Dsx1CurrentEntry[i]})
    }
    dsx1CurrentTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1CurrentTable.EntityData.YListKeys = []string {}

    return &(dsx1CurrentTable.EntityData)
}

// DS1MIB_Dsx1CurrentTable_Dsx1CurrentEntry
// An entry in the DS1 Current table.
type DS1MIB_Dsx1CurrentTable_Dsx1CurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies  the DS1
    // interface to which this entry is applicable. The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1CurrentIndex interface{}

    // The number of Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentESs interface{}

    // The number of Severely Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentSESs interface{}

    // The number of Severely Errored Framing Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1CurrentSEFSs interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentUASs interface{}

    // The number of Controlled Slip Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentCSSs interface{}

    // The number of Path Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentPCVs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentLESs interface{}

    // The number of Bursty Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentBESs interface{}

    // The number of Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1CurrentDMs interface{}

    // The number of Line Code Violations (LCVs). The type is interface{} with
    // range: 0..4294967295.
    Dsx1CurrentLCVs interface{}
}

func (dsx1CurrentEntry *DS1MIB_Dsx1CurrentTable_Dsx1CurrentEntry) GetEntityData() *types.CommonEntityData {
    dsx1CurrentEntry.EntityData.YFilter = dsx1CurrentEntry.YFilter
    dsx1CurrentEntry.EntityData.YangName = "dsx1CurrentEntry"
    dsx1CurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1CurrentEntry.EntityData.ParentYangName = "dsx1CurrentTable"
    dsx1CurrentEntry.EntityData.SegmentPath = "dsx1CurrentEntry" + types.AddKeyToken(dsx1CurrentEntry.Dsx1CurrentIndex, "dsx1CurrentIndex")
    dsx1CurrentEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1CurrentTable/" + dsx1CurrentEntry.EntityData.SegmentPath
    dsx1CurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1CurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1CurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1CurrentEntry.EntityData.Children = types.NewOrderedMap()
    dsx1CurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentIndex", types.YLeaf{"Dsx1CurrentIndex", dsx1CurrentEntry.Dsx1CurrentIndex})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentESs", types.YLeaf{"Dsx1CurrentESs", dsx1CurrentEntry.Dsx1CurrentESs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentSESs", types.YLeaf{"Dsx1CurrentSESs", dsx1CurrentEntry.Dsx1CurrentSESs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentSEFSs", types.YLeaf{"Dsx1CurrentSEFSs", dsx1CurrentEntry.Dsx1CurrentSEFSs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentUASs", types.YLeaf{"Dsx1CurrentUASs", dsx1CurrentEntry.Dsx1CurrentUASs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentCSSs", types.YLeaf{"Dsx1CurrentCSSs", dsx1CurrentEntry.Dsx1CurrentCSSs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentPCVs", types.YLeaf{"Dsx1CurrentPCVs", dsx1CurrentEntry.Dsx1CurrentPCVs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentLESs", types.YLeaf{"Dsx1CurrentLESs", dsx1CurrentEntry.Dsx1CurrentLESs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentBESs", types.YLeaf{"Dsx1CurrentBESs", dsx1CurrentEntry.Dsx1CurrentBESs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentDMs", types.YLeaf{"Dsx1CurrentDMs", dsx1CurrentEntry.Dsx1CurrentDMs})
    dsx1CurrentEntry.EntityData.Leafs.Append("dsx1CurrentLCVs", types.YLeaf{"Dsx1CurrentLCVs", dsx1CurrentEntry.Dsx1CurrentLCVs})

    dsx1CurrentEntry.EntityData.YListKeys = []string {"Dsx1CurrentIndex"}

    return &(dsx1CurrentEntry.EntityData)
}

// DS1MIB_Dsx1IntervalTable
// The DS1 Interval Table contains various
// statistics collected by each DS1 Interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.  Each row in this table represents one
// such interval (identified by dsx1IntervalNumber)
// for one specific instance (identified by
// dsx1IntervalIndex).
type DS1MIB_Dsx1IntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Interval table. The type is slice of
    // DS1MIB_Dsx1IntervalTable_Dsx1IntervalEntry.
    Dsx1IntervalEntry []*DS1MIB_Dsx1IntervalTable_Dsx1IntervalEntry
}

func (dsx1IntervalTable *DS1MIB_Dsx1IntervalTable) GetEntityData() *types.CommonEntityData {
    dsx1IntervalTable.EntityData.YFilter = dsx1IntervalTable.YFilter
    dsx1IntervalTable.EntityData.YangName = "dsx1IntervalTable"
    dsx1IntervalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1IntervalTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1IntervalTable.EntityData.SegmentPath = "dsx1IntervalTable"
    dsx1IntervalTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1IntervalTable.EntityData.SegmentPath
    dsx1IntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1IntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1IntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1IntervalTable.EntityData.Children = types.NewOrderedMap()
    dsx1IntervalTable.EntityData.Children.Append("dsx1IntervalEntry", types.YChild{"Dsx1IntervalEntry", nil})
    for i := range dsx1IntervalTable.Dsx1IntervalEntry {
        dsx1IntervalTable.EntityData.Children.Append(types.GetSegmentPath(dsx1IntervalTable.Dsx1IntervalEntry[i]), types.YChild{"Dsx1IntervalEntry", dsx1IntervalTable.Dsx1IntervalEntry[i]})
    }
    dsx1IntervalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1IntervalTable.EntityData.YListKeys = []string {}

    return &(dsx1IntervalTable.EntityData)
}

// DS1MIB_Dsx1IntervalTable_Dsx1IntervalEntry
// An entry in the DS1 Interval table.
type DS1MIB_Dsx1IntervalTable_Dsx1IntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1IntervalIndex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx1IntervalNumber interface{}

    // The number of Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalESs interface{}

    // The number of Severely Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalSESs interface{}

    // The number of Severely Errored Framing Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1IntervalSEFSs interface{}

    // The number of Unavailable Seconds.  This object may decrease if the
    // occurance of unavailable seconds occurs across an inteval boundary. The
    // type is interface{} with range: 0..4294967295.
    Dsx1IntervalUASs interface{}

    // The number of Controlled Slip Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalCSSs interface{}

    // The number of Path Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalPCVs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalLESs interface{}

    // The number of Bursty Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalBESs interface{}

    // The number of Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalDMs interface{}

    // The number of Line Code Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx1IntervalLCVs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx1IntervalValidData interface{}
}

func (dsx1IntervalEntry *DS1MIB_Dsx1IntervalTable_Dsx1IntervalEntry) GetEntityData() *types.CommonEntityData {
    dsx1IntervalEntry.EntityData.YFilter = dsx1IntervalEntry.YFilter
    dsx1IntervalEntry.EntityData.YangName = "dsx1IntervalEntry"
    dsx1IntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1IntervalEntry.EntityData.ParentYangName = "dsx1IntervalTable"
    dsx1IntervalEntry.EntityData.SegmentPath = "dsx1IntervalEntry" + types.AddKeyToken(dsx1IntervalEntry.Dsx1IntervalIndex, "dsx1IntervalIndex") + types.AddKeyToken(dsx1IntervalEntry.Dsx1IntervalNumber, "dsx1IntervalNumber")
    dsx1IntervalEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1IntervalTable/" + dsx1IntervalEntry.EntityData.SegmentPath
    dsx1IntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1IntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1IntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1IntervalEntry.EntityData.Children = types.NewOrderedMap()
    dsx1IntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalIndex", types.YLeaf{"Dsx1IntervalIndex", dsx1IntervalEntry.Dsx1IntervalIndex})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalNumber", types.YLeaf{"Dsx1IntervalNumber", dsx1IntervalEntry.Dsx1IntervalNumber})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalESs", types.YLeaf{"Dsx1IntervalESs", dsx1IntervalEntry.Dsx1IntervalESs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalSESs", types.YLeaf{"Dsx1IntervalSESs", dsx1IntervalEntry.Dsx1IntervalSESs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalSEFSs", types.YLeaf{"Dsx1IntervalSEFSs", dsx1IntervalEntry.Dsx1IntervalSEFSs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalUASs", types.YLeaf{"Dsx1IntervalUASs", dsx1IntervalEntry.Dsx1IntervalUASs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalCSSs", types.YLeaf{"Dsx1IntervalCSSs", dsx1IntervalEntry.Dsx1IntervalCSSs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalPCVs", types.YLeaf{"Dsx1IntervalPCVs", dsx1IntervalEntry.Dsx1IntervalPCVs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalLESs", types.YLeaf{"Dsx1IntervalLESs", dsx1IntervalEntry.Dsx1IntervalLESs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalBESs", types.YLeaf{"Dsx1IntervalBESs", dsx1IntervalEntry.Dsx1IntervalBESs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalDMs", types.YLeaf{"Dsx1IntervalDMs", dsx1IntervalEntry.Dsx1IntervalDMs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalLCVs", types.YLeaf{"Dsx1IntervalLCVs", dsx1IntervalEntry.Dsx1IntervalLCVs})
    dsx1IntervalEntry.EntityData.Leafs.Append("dsx1IntervalValidData", types.YLeaf{"Dsx1IntervalValidData", dsx1IntervalEntry.Dsx1IntervalValidData})

    dsx1IntervalEntry.EntityData.YListKeys = []string {"Dsx1IntervalIndex", "Dsx1IntervalNumber"}

    return &(dsx1IntervalEntry.EntityData)
}

// DS1MIB_Dsx1TotalTable
// The DS1 Total Table contains the cumulative sum
// of the various statistics for the 24 hour period
// preceding the current interval.
type DS1MIB_Dsx1TotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Total table. The type is slice of
    // DS1MIB_Dsx1TotalTable_Dsx1TotalEntry.
    Dsx1TotalEntry []*DS1MIB_Dsx1TotalTable_Dsx1TotalEntry
}

func (dsx1TotalTable *DS1MIB_Dsx1TotalTable) GetEntityData() *types.CommonEntityData {
    dsx1TotalTable.EntityData.YFilter = dsx1TotalTable.YFilter
    dsx1TotalTable.EntityData.YangName = "dsx1TotalTable"
    dsx1TotalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1TotalTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1TotalTable.EntityData.SegmentPath = "dsx1TotalTable"
    dsx1TotalTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1TotalTable.EntityData.SegmentPath
    dsx1TotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1TotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1TotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1TotalTable.EntityData.Children = types.NewOrderedMap()
    dsx1TotalTable.EntityData.Children.Append("dsx1TotalEntry", types.YChild{"Dsx1TotalEntry", nil})
    for i := range dsx1TotalTable.Dsx1TotalEntry {
        dsx1TotalTable.EntityData.Children.Append(types.GetSegmentPath(dsx1TotalTable.Dsx1TotalEntry[i]), types.YChild{"Dsx1TotalEntry", dsx1TotalTable.Dsx1TotalEntry[i]})
    }
    dsx1TotalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1TotalTable.EntityData.YListKeys = []string {}

    return &(dsx1TotalTable.EntityData)
}

// DS1MIB_Dsx1TotalTable_Dsx1TotalEntry
// An entry in the DS1 Total table.
type DS1MIB_Dsx1TotalTable_Dsx1TotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value as a dsx1LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx1TotalIndex interface{}

    // The sum of Errored Seconds encountered by a DS1 interface in the previous
    // 24 hour interval. Invalid 15 minute intervals count as 0. The type is
    // interface{} with range: 0..4294967295.
    Dsx1TotalESs interface{}

    // The number of Severely Errored Seconds encountered by a DS1 interface in
    // the previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1TotalSESs interface{}

    // The number of Severely Errored Framing Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1TotalSEFSs interface{}

    // The number of Unavailable Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1TotalUASs interface{}

    // The number of Controlled Slip Seconds encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1TotalCSSs interface{}

    // The number of Path Coding Violations encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1TotalPCVs interface{}

    // The number of Line Errored Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1TotalLESs interface{}

    // The number of Bursty Errored Seconds (BESs) encountered by a DS1 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1TotalBESs interface{}

    // The number of Degraded Minutes (DMs) encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1TotalDMs interface{}

    // The number of Line Code Violations (LCVs) encountered by a DS1 interface in
    // the current 15 minute interval.  Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1TotalLCVs interface{}
}

func (dsx1TotalEntry *DS1MIB_Dsx1TotalTable_Dsx1TotalEntry) GetEntityData() *types.CommonEntityData {
    dsx1TotalEntry.EntityData.YFilter = dsx1TotalEntry.YFilter
    dsx1TotalEntry.EntityData.YangName = "dsx1TotalEntry"
    dsx1TotalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1TotalEntry.EntityData.ParentYangName = "dsx1TotalTable"
    dsx1TotalEntry.EntityData.SegmentPath = "dsx1TotalEntry" + types.AddKeyToken(dsx1TotalEntry.Dsx1TotalIndex, "dsx1TotalIndex")
    dsx1TotalEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1TotalTable/" + dsx1TotalEntry.EntityData.SegmentPath
    dsx1TotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1TotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1TotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1TotalEntry.EntityData.Children = types.NewOrderedMap()
    dsx1TotalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalIndex", types.YLeaf{"Dsx1TotalIndex", dsx1TotalEntry.Dsx1TotalIndex})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalESs", types.YLeaf{"Dsx1TotalESs", dsx1TotalEntry.Dsx1TotalESs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalSESs", types.YLeaf{"Dsx1TotalSESs", dsx1TotalEntry.Dsx1TotalSESs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalSEFSs", types.YLeaf{"Dsx1TotalSEFSs", dsx1TotalEntry.Dsx1TotalSEFSs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalUASs", types.YLeaf{"Dsx1TotalUASs", dsx1TotalEntry.Dsx1TotalUASs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalCSSs", types.YLeaf{"Dsx1TotalCSSs", dsx1TotalEntry.Dsx1TotalCSSs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalPCVs", types.YLeaf{"Dsx1TotalPCVs", dsx1TotalEntry.Dsx1TotalPCVs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalLESs", types.YLeaf{"Dsx1TotalLESs", dsx1TotalEntry.Dsx1TotalLESs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalBESs", types.YLeaf{"Dsx1TotalBESs", dsx1TotalEntry.Dsx1TotalBESs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalDMs", types.YLeaf{"Dsx1TotalDMs", dsx1TotalEntry.Dsx1TotalDMs})
    dsx1TotalEntry.EntityData.Leafs.Append("dsx1TotalLCVs", types.YLeaf{"Dsx1TotalLCVs", dsx1TotalEntry.Dsx1TotalLCVs})

    dsx1TotalEntry.EntityData.YListKeys = []string {"Dsx1TotalIndex"}

    return &(dsx1TotalEntry.EntityData)
}

// DS1MIB_Dsx1FarEndCurrentTable
// The DS1 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end messages on the Facilities Data
// Link.  The definitions are the same as described
// for the near-end information.
type DS1MIB_Dsx1FarEndCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Current table. The type is slice of
    // DS1MIB_Dsx1FarEndCurrentTable_Dsx1FarEndCurrentEntry.
    Dsx1FarEndCurrentEntry []*DS1MIB_Dsx1FarEndCurrentTable_Dsx1FarEndCurrentEntry
}

func (dsx1FarEndCurrentTable *DS1MIB_Dsx1FarEndCurrentTable) GetEntityData() *types.CommonEntityData {
    dsx1FarEndCurrentTable.EntityData.YFilter = dsx1FarEndCurrentTable.YFilter
    dsx1FarEndCurrentTable.EntityData.YangName = "dsx1FarEndCurrentTable"
    dsx1FarEndCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndCurrentTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1FarEndCurrentTable.EntityData.SegmentPath = "dsx1FarEndCurrentTable"
    dsx1FarEndCurrentTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1FarEndCurrentTable.EntityData.SegmentPath
    dsx1FarEndCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndCurrentTable.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndCurrentTable.EntityData.Children.Append("dsx1FarEndCurrentEntry", types.YChild{"Dsx1FarEndCurrentEntry", nil})
    for i := range dsx1FarEndCurrentTable.Dsx1FarEndCurrentEntry {
        dsx1FarEndCurrentTable.EntityData.Children.Append(types.GetSegmentPath(dsx1FarEndCurrentTable.Dsx1FarEndCurrentEntry[i]), types.YChild{"Dsx1FarEndCurrentEntry", dsx1FarEndCurrentTable.Dsx1FarEndCurrentEntry[i]})
    }
    dsx1FarEndCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1FarEndCurrentTable.EntityData.YListKeys = []string {}

    return &(dsx1FarEndCurrentTable.EntityData)
}

// DS1MIB_Dsx1FarEndCurrentTable_Dsx1FarEndCurrentEntry
// An entry in the DS1 Far End Current table.
type DS1MIB_Dsx1FarEndCurrentTable_Dsx1FarEndCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1FarEndCurrentIndex interface{}

    // The number of seconds that have elapsed since the beginning of the far end
    // current error-measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx1FarEndTimeElapsed interface{}

    // The number of previous far end intervals for which data was collected.  The
    // value will be 96 unless the interface was brought online within the last 24
    // hours, in which case the value will be the number of complete 15 minute far
    // end intervals since the interface has been online. The type is interface{}
    // with range: 0..96.
    Dsx1FarEndValidIntervals interface{}

    // The number of Far End Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndCurrentESs interface{}

    // The number of Far End Severely Errored Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1FarEndCurrentSESs interface{}

    // The number of Far End Severely Errored Framing Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx1FarEndCurrentSEFSs interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndCurrentUASs interface{}

    // The number of Far End Controlled Slip Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndCurrentCSSs interface{}

    // The number of Far End Line Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndCurrentLESs interface{}

    // The number of Far End Path Coding Violations. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndCurrentPCVs interface{}

    // The number of Far End Bursty Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndCurrentBESs interface{}

    // The number of Far End Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndCurrentDMs interface{}

    // The number of intervals in the range from 0 to dsx1FarEndValidIntervals for
    // which no data is available.  This object will typically be zero except in
    // cases where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx1FarEndInvalidIntervals interface{}
}

func (dsx1FarEndCurrentEntry *DS1MIB_Dsx1FarEndCurrentTable_Dsx1FarEndCurrentEntry) GetEntityData() *types.CommonEntityData {
    dsx1FarEndCurrentEntry.EntityData.YFilter = dsx1FarEndCurrentEntry.YFilter
    dsx1FarEndCurrentEntry.EntityData.YangName = "dsx1FarEndCurrentEntry"
    dsx1FarEndCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndCurrentEntry.EntityData.ParentYangName = "dsx1FarEndCurrentTable"
    dsx1FarEndCurrentEntry.EntityData.SegmentPath = "dsx1FarEndCurrentEntry" + types.AddKeyToken(dsx1FarEndCurrentEntry.Dsx1FarEndCurrentIndex, "dsx1FarEndCurrentIndex")
    dsx1FarEndCurrentEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1FarEndCurrentTable/" + dsx1FarEndCurrentEntry.EntityData.SegmentPath
    dsx1FarEndCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndCurrentEntry.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentIndex", types.YLeaf{"Dsx1FarEndCurrentIndex", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentIndex})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndTimeElapsed", types.YLeaf{"Dsx1FarEndTimeElapsed", dsx1FarEndCurrentEntry.Dsx1FarEndTimeElapsed})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndValidIntervals", types.YLeaf{"Dsx1FarEndValidIntervals", dsx1FarEndCurrentEntry.Dsx1FarEndValidIntervals})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentESs", types.YLeaf{"Dsx1FarEndCurrentESs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentESs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentSESs", types.YLeaf{"Dsx1FarEndCurrentSESs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentSESs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentSEFSs", types.YLeaf{"Dsx1FarEndCurrentSEFSs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentSEFSs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentUASs", types.YLeaf{"Dsx1FarEndCurrentUASs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentUASs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentCSSs", types.YLeaf{"Dsx1FarEndCurrentCSSs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentCSSs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentLESs", types.YLeaf{"Dsx1FarEndCurrentLESs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentLESs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentPCVs", types.YLeaf{"Dsx1FarEndCurrentPCVs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentPCVs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentBESs", types.YLeaf{"Dsx1FarEndCurrentBESs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentBESs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndCurrentDMs", types.YLeaf{"Dsx1FarEndCurrentDMs", dsx1FarEndCurrentEntry.Dsx1FarEndCurrentDMs})
    dsx1FarEndCurrentEntry.EntityData.Leafs.Append("dsx1FarEndInvalidIntervals", types.YLeaf{"Dsx1FarEndInvalidIntervals", dsx1FarEndCurrentEntry.Dsx1FarEndInvalidIntervals})

    dsx1FarEndCurrentEntry.EntityData.YListKeys = []string {"Dsx1FarEndCurrentIndex"}

    return &(dsx1FarEndCurrentEntry.EntityData)
}

// DS1MIB_Dsx1FarEndIntervalTable
// The DS1 Far End Interval Table contains various
// statistics collected by each DS1 interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals. Each row in this table represents one
// such interval (identified by
// dsx1FarEndIntervalNumber) for one specific
// instance (identified by dsx1FarEndIntervalIndex).
type DS1MIB_Dsx1FarEndIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Interval table. The type is slice of
    // DS1MIB_Dsx1FarEndIntervalTable_Dsx1FarEndIntervalEntry.
    Dsx1FarEndIntervalEntry []*DS1MIB_Dsx1FarEndIntervalTable_Dsx1FarEndIntervalEntry
}

func (dsx1FarEndIntervalTable *DS1MIB_Dsx1FarEndIntervalTable) GetEntityData() *types.CommonEntityData {
    dsx1FarEndIntervalTable.EntityData.YFilter = dsx1FarEndIntervalTable.YFilter
    dsx1FarEndIntervalTable.EntityData.YangName = "dsx1FarEndIntervalTable"
    dsx1FarEndIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndIntervalTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1FarEndIntervalTable.EntityData.SegmentPath = "dsx1FarEndIntervalTable"
    dsx1FarEndIntervalTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1FarEndIntervalTable.EntityData.SegmentPath
    dsx1FarEndIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndIntervalTable.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndIntervalTable.EntityData.Children.Append("dsx1FarEndIntervalEntry", types.YChild{"Dsx1FarEndIntervalEntry", nil})
    for i := range dsx1FarEndIntervalTable.Dsx1FarEndIntervalEntry {
        dsx1FarEndIntervalTable.EntityData.Children.Append(types.GetSegmentPath(dsx1FarEndIntervalTable.Dsx1FarEndIntervalEntry[i]), types.YChild{"Dsx1FarEndIntervalEntry", dsx1FarEndIntervalTable.Dsx1FarEndIntervalEntry[i]})
    }
    dsx1FarEndIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1FarEndIntervalTable.EntityData.YListKeys = []string {}

    return &(dsx1FarEndIntervalTable.EntityData)
}

// DS1MIB_Dsx1FarEndIntervalTable_Dsx1FarEndIntervalEntry
// An entry in the DS1 Far End Interval table.
type DS1MIB_Dsx1FarEndIntervalTable_Dsx1FarEndIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1FarEndIntervalIndex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx1FarEndIntervalNumber interface{}

    // The number of Far End Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndIntervalESs interface{}

    // The number of Far End Severely Errored Seconds. The type is interface{}
    // with range: 0..4294967295.
    Dsx1FarEndIntervalSESs interface{}

    // The number of Far End Severely Errored Framing Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx1FarEndIntervalSEFSs interface{}

    // The number of Unavailable Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndIntervalUASs interface{}

    // The number of Far End Controlled Slip Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndIntervalCSSs interface{}

    // The number of Far End Line Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndIntervalLESs interface{}

    // The number of Far End Path Coding Violations. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndIntervalPCVs interface{}

    // The number of Far End Bursty Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx1FarEndIntervalBESs interface{}

    // The number of Far End Degraded Minutes. The type is interface{} with range:
    // 0..4294967295.
    Dsx1FarEndIntervalDMs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx1FarEndIntervalValidData interface{}
}

func (dsx1FarEndIntervalEntry *DS1MIB_Dsx1FarEndIntervalTable_Dsx1FarEndIntervalEntry) GetEntityData() *types.CommonEntityData {
    dsx1FarEndIntervalEntry.EntityData.YFilter = dsx1FarEndIntervalEntry.YFilter
    dsx1FarEndIntervalEntry.EntityData.YangName = "dsx1FarEndIntervalEntry"
    dsx1FarEndIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndIntervalEntry.EntityData.ParentYangName = "dsx1FarEndIntervalTable"
    dsx1FarEndIntervalEntry.EntityData.SegmentPath = "dsx1FarEndIntervalEntry" + types.AddKeyToken(dsx1FarEndIntervalEntry.Dsx1FarEndIntervalIndex, "dsx1FarEndIntervalIndex") + types.AddKeyToken(dsx1FarEndIntervalEntry.Dsx1FarEndIntervalNumber, "dsx1FarEndIntervalNumber")
    dsx1FarEndIntervalEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1FarEndIntervalTable/" + dsx1FarEndIntervalEntry.EntityData.SegmentPath
    dsx1FarEndIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndIntervalEntry.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalIndex", types.YLeaf{"Dsx1FarEndIntervalIndex", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalIndex})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalNumber", types.YLeaf{"Dsx1FarEndIntervalNumber", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalNumber})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalESs", types.YLeaf{"Dsx1FarEndIntervalESs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalESs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalSESs", types.YLeaf{"Dsx1FarEndIntervalSESs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalSESs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalSEFSs", types.YLeaf{"Dsx1FarEndIntervalSEFSs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalSEFSs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalUASs", types.YLeaf{"Dsx1FarEndIntervalUASs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalUASs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalCSSs", types.YLeaf{"Dsx1FarEndIntervalCSSs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalCSSs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalLESs", types.YLeaf{"Dsx1FarEndIntervalLESs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalLESs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalPCVs", types.YLeaf{"Dsx1FarEndIntervalPCVs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalPCVs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalBESs", types.YLeaf{"Dsx1FarEndIntervalBESs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalBESs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalDMs", types.YLeaf{"Dsx1FarEndIntervalDMs", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalDMs})
    dsx1FarEndIntervalEntry.EntityData.Leafs.Append("dsx1FarEndIntervalValidData", types.YLeaf{"Dsx1FarEndIntervalValidData", dsx1FarEndIntervalEntry.Dsx1FarEndIntervalValidData})

    dsx1FarEndIntervalEntry.EntityData.YListKeys = []string {"Dsx1FarEndIntervalIndex", "Dsx1FarEndIntervalNumber"}

    return &(dsx1FarEndIntervalEntry.EntityData)
}

// DS1MIB_Dsx1FarEndTotalTable
// The DS1 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS1MIB_Dsx1FarEndTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Far End Total table. The type is slice of
    // DS1MIB_Dsx1FarEndTotalTable_Dsx1FarEndTotalEntry.
    Dsx1FarEndTotalEntry []*DS1MIB_Dsx1FarEndTotalTable_Dsx1FarEndTotalEntry
}

func (dsx1FarEndTotalTable *DS1MIB_Dsx1FarEndTotalTable) GetEntityData() *types.CommonEntityData {
    dsx1FarEndTotalTable.EntityData.YFilter = dsx1FarEndTotalTable.YFilter
    dsx1FarEndTotalTable.EntityData.YangName = "dsx1FarEndTotalTable"
    dsx1FarEndTotalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndTotalTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1FarEndTotalTable.EntityData.SegmentPath = "dsx1FarEndTotalTable"
    dsx1FarEndTotalTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1FarEndTotalTable.EntityData.SegmentPath
    dsx1FarEndTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndTotalTable.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndTotalTable.EntityData.Children.Append("dsx1FarEndTotalEntry", types.YChild{"Dsx1FarEndTotalEntry", nil})
    for i := range dsx1FarEndTotalTable.Dsx1FarEndTotalEntry {
        dsx1FarEndTotalTable.EntityData.Children.Append(types.GetSegmentPath(dsx1FarEndTotalTable.Dsx1FarEndTotalEntry[i]), types.YChild{"Dsx1FarEndTotalEntry", dsx1FarEndTotalTable.Dsx1FarEndTotalEntry[i]})
    }
    dsx1FarEndTotalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1FarEndTotalTable.EntityData.YListKeys = []string {}

    return &(dsx1FarEndTotalTable.EntityData)
}

// DS1MIB_Dsx1FarEndTotalTable_Dsx1FarEndTotalEntry
// An entry in the DS1 Far End Total table.
type DS1MIB_Dsx1FarEndTotalTable_Dsx1FarEndTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS1
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx1LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx1FarEndTotalIndex interface{}

    // The number of Far End Errored Seconds encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalESs interface{}

    // The number of Far End Severely Errored Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalSESs interface{}

    // The number of Far End Severely Errored Framing Seconds encountered by a DS1
    // interface in the previous 24 hour interval. Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalSEFSs interface{}

    // The number of Unavailable Seconds encountered by a DS1 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalUASs interface{}

    // The number of Far End Controlled Slip Seconds encountered by a DS1
    // interface in the previous 24 hour interval.  Invalid 15 minute intervals
    // count as 0. The type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalCSSs interface{}

    // The number of Far End Line Errored Seconds encountered by a DS1 interface
    // in the previous 24 hour interval.  Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalLESs interface{}

    // The number of Far End Path Coding Violations reported via the far end block
    // error count encountered by a DS1 interface in the previous 24 hour
    // interval.  Invalid 15 minute intervals count as 0. The type is interface{}
    // with range: 0..4294967295.
    Dsx1FarEndTotalPCVs interface{}

    // The number of Bursty Errored Seconds (BESs) encountered by a DS1 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalBESs interface{}

    // The number of Degraded Minutes (DMs) encountered by a DS1 interface in the
    // previous 24 hour interval.  Invalid 15 minute intervals count as 0. The
    // type is interface{} with range: 0..4294967295.
    Dsx1FarEndTotalDMs interface{}
}

func (dsx1FarEndTotalEntry *DS1MIB_Dsx1FarEndTotalTable_Dsx1FarEndTotalEntry) GetEntityData() *types.CommonEntityData {
    dsx1FarEndTotalEntry.EntityData.YFilter = dsx1FarEndTotalEntry.YFilter
    dsx1FarEndTotalEntry.EntityData.YangName = "dsx1FarEndTotalEntry"
    dsx1FarEndTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1FarEndTotalEntry.EntityData.ParentYangName = "dsx1FarEndTotalTable"
    dsx1FarEndTotalEntry.EntityData.SegmentPath = "dsx1FarEndTotalEntry" + types.AddKeyToken(dsx1FarEndTotalEntry.Dsx1FarEndTotalIndex, "dsx1FarEndTotalIndex")
    dsx1FarEndTotalEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1FarEndTotalTable/" + dsx1FarEndTotalEntry.EntityData.SegmentPath
    dsx1FarEndTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FarEndTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FarEndTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FarEndTotalEntry.EntityData.Children = types.NewOrderedMap()
    dsx1FarEndTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalIndex", types.YLeaf{"Dsx1FarEndTotalIndex", dsx1FarEndTotalEntry.Dsx1FarEndTotalIndex})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalESs", types.YLeaf{"Dsx1FarEndTotalESs", dsx1FarEndTotalEntry.Dsx1FarEndTotalESs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalSESs", types.YLeaf{"Dsx1FarEndTotalSESs", dsx1FarEndTotalEntry.Dsx1FarEndTotalSESs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalSEFSs", types.YLeaf{"Dsx1FarEndTotalSEFSs", dsx1FarEndTotalEntry.Dsx1FarEndTotalSEFSs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalUASs", types.YLeaf{"Dsx1FarEndTotalUASs", dsx1FarEndTotalEntry.Dsx1FarEndTotalUASs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalCSSs", types.YLeaf{"Dsx1FarEndTotalCSSs", dsx1FarEndTotalEntry.Dsx1FarEndTotalCSSs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalLESs", types.YLeaf{"Dsx1FarEndTotalLESs", dsx1FarEndTotalEntry.Dsx1FarEndTotalLESs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalPCVs", types.YLeaf{"Dsx1FarEndTotalPCVs", dsx1FarEndTotalEntry.Dsx1FarEndTotalPCVs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalBESs", types.YLeaf{"Dsx1FarEndTotalBESs", dsx1FarEndTotalEntry.Dsx1FarEndTotalBESs})
    dsx1FarEndTotalEntry.EntityData.Leafs.Append("dsx1FarEndTotalDMs", types.YLeaf{"Dsx1FarEndTotalDMs", dsx1FarEndTotalEntry.Dsx1FarEndTotalDMs})

    dsx1FarEndTotalEntry.EntityData.YListKeys = []string {"Dsx1FarEndTotalIndex"}

    return &(dsx1FarEndTotalEntry.EntityData)
}

// DS1MIB_Dsx1FracTable
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
type DS1MIB_Dsx1FracTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS1 Fractional table. The type is slice of
    // DS1MIB_Dsx1FracTable_Dsx1FracEntry.
    Dsx1FracEntry []*DS1MIB_Dsx1FracTable_Dsx1FracEntry
}

func (dsx1FracTable *DS1MIB_Dsx1FracTable) GetEntityData() *types.CommonEntityData {
    dsx1FracTable.EntityData.YFilter = dsx1FracTable.YFilter
    dsx1FracTable.EntityData.YangName = "dsx1FracTable"
    dsx1FracTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1FracTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1FracTable.EntityData.SegmentPath = "dsx1FracTable"
    dsx1FracTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1FracTable.EntityData.SegmentPath
    dsx1FracTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FracTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FracTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FracTable.EntityData.Children = types.NewOrderedMap()
    dsx1FracTable.EntityData.Children.Append("dsx1FracEntry", types.YChild{"Dsx1FracEntry", nil})
    for i := range dsx1FracTable.Dsx1FracEntry {
        dsx1FracTable.EntityData.Children.Append(types.GetSegmentPath(dsx1FracTable.Dsx1FracEntry[i]), types.YChild{"Dsx1FracEntry", dsx1FracTable.Dsx1FracEntry[i]})
    }
    dsx1FracTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1FracTable.EntityData.YListKeys = []string {}

    return &(dsx1FracTable.EntityData)
}

// DS1MIB_Dsx1FracTable_Dsx1FracEntry
// An entry in the DS1 Fractional table.
type DS1MIB_Dsx1FracTable_Dsx1FracEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies  the DS1
    // interface  to which this entry is applicable The interface identified by a 
    // particular value  of  this  index is the same interface as identified by
    // the same value  an  dsx1LineIndex object instance. The type is interface{}
    // with range: 1..2147483647.
    Dsx1FracIndex interface{}

    // This attribute is a key. The channel number for this entry. The type is
    // interface{} with range: 1..31.
    Dsx1FracNumber interface{}

    // An index value that uniquely identifies an interface.  The interface
    // identified by a particular value of this index is the same  interface as 
    // identified by the same value an ifIndex object instance. If no interface is
    // currently using a channel, the value should be zero.  If a single interface
    // occupies more  than  one  time slot,  that ifIndex value will be found in
    // multiple time slots. The type is interface{} with range: 1..2147483647.
    Dsx1FracIfIndex interface{}
}

func (dsx1FracEntry *DS1MIB_Dsx1FracTable_Dsx1FracEntry) GetEntityData() *types.CommonEntityData {
    dsx1FracEntry.EntityData.YFilter = dsx1FracEntry.YFilter
    dsx1FracEntry.EntityData.YangName = "dsx1FracEntry"
    dsx1FracEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1FracEntry.EntityData.ParentYangName = "dsx1FracTable"
    dsx1FracEntry.EntityData.SegmentPath = "dsx1FracEntry" + types.AddKeyToken(dsx1FracEntry.Dsx1FracIndex, "dsx1FracIndex") + types.AddKeyToken(dsx1FracEntry.Dsx1FracNumber, "dsx1FracNumber")
    dsx1FracEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1FracTable/" + dsx1FracEntry.EntityData.SegmentPath
    dsx1FracEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1FracEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1FracEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1FracEntry.EntityData.Children = types.NewOrderedMap()
    dsx1FracEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1FracEntry.EntityData.Leafs.Append("dsx1FracIndex", types.YLeaf{"Dsx1FracIndex", dsx1FracEntry.Dsx1FracIndex})
    dsx1FracEntry.EntityData.Leafs.Append("dsx1FracNumber", types.YLeaf{"Dsx1FracNumber", dsx1FracEntry.Dsx1FracNumber})
    dsx1FracEntry.EntityData.Leafs.Append("dsx1FracIfIndex", types.YLeaf{"Dsx1FracIfIndex", dsx1FracEntry.Dsx1FracIfIndex})

    dsx1FracEntry.EntityData.YListKeys = []string {"Dsx1FracIndex", "Dsx1FracNumber"}

    return &(dsx1FracEntry.EntityData)
}

// DS1MIB_Dsx1ChanMappingTable
// The DS1 Channel Mapping table.  This table maps a
// DS1 channel number on a particular DS3 into an
// ifIndex.  In the presence of DS2s, this table can
// be used to map a DS2 channel number on a DS3 into
// an ifIndex, or used to map a DS1 channel number on
// a DS2 onto an ifIndex.
type DS1MIB_Dsx1ChanMappingTable struct {
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
    // type is slice of DS1MIB_Dsx1ChanMappingTable_Dsx1ChanMappingEntry.
    Dsx1ChanMappingEntry []*DS1MIB_Dsx1ChanMappingTable_Dsx1ChanMappingEntry
}

func (dsx1ChanMappingTable *DS1MIB_Dsx1ChanMappingTable) GetEntityData() *types.CommonEntityData {
    dsx1ChanMappingTable.EntityData.YFilter = dsx1ChanMappingTable.YFilter
    dsx1ChanMappingTable.EntityData.YangName = "dsx1ChanMappingTable"
    dsx1ChanMappingTable.EntityData.BundleName = "cisco_ios_xe"
    dsx1ChanMappingTable.EntityData.ParentYangName = "DS1-MIB"
    dsx1ChanMappingTable.EntityData.SegmentPath = "dsx1ChanMappingTable"
    dsx1ChanMappingTable.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/" + dsx1ChanMappingTable.EntityData.SegmentPath
    dsx1ChanMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1ChanMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1ChanMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1ChanMappingTable.EntityData.Children = types.NewOrderedMap()
    dsx1ChanMappingTable.EntityData.Children.Append("dsx1ChanMappingEntry", types.YChild{"Dsx1ChanMappingEntry", nil})
    for i := range dsx1ChanMappingTable.Dsx1ChanMappingEntry {
        dsx1ChanMappingTable.EntityData.Children.Append(types.GetSegmentPath(dsx1ChanMappingTable.Dsx1ChanMappingEntry[i]), types.YChild{"Dsx1ChanMappingEntry", dsx1ChanMappingTable.Dsx1ChanMappingEntry[i]})
    }
    dsx1ChanMappingTable.EntityData.Leafs = types.NewOrderedMap()

    dsx1ChanMappingTable.EntityData.YListKeys = []string {}

    return &(dsx1ChanMappingTable.EntityData)
}

// DS1MIB_Dsx1ChanMappingTable_Dsx1ChanMappingEntry
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
type DS1MIB_Dsx1ChanMappingTable_Dsx1ChanMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..28. Refers to
    // ds1_mib.DS1MIB_Dsx1ConfigTable_Dsx1ConfigEntry_Dsx1Ds1ChannelNumber
    Dsx1Ds1ChannelNumber interface{}

    // This object indicates the ifIndex value assigned by the agent for the
    // individual ds1 ifEntry that corresponds to the given DS1 channel number
    // (specified by the INDEX element dsx1Ds1ChannelNumber) of the given
    // channelized interface (specified by INDEX element ifIndex). The type is
    // interface{} with range: 1..2147483647.
    Dsx1ChanMappedIfIndex interface{}
}

func (dsx1ChanMappingEntry *DS1MIB_Dsx1ChanMappingTable_Dsx1ChanMappingEntry) GetEntityData() *types.CommonEntityData {
    dsx1ChanMappingEntry.EntityData.YFilter = dsx1ChanMappingEntry.YFilter
    dsx1ChanMappingEntry.EntityData.YangName = "dsx1ChanMappingEntry"
    dsx1ChanMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx1ChanMappingEntry.EntityData.ParentYangName = "dsx1ChanMappingTable"
    dsx1ChanMappingEntry.EntityData.SegmentPath = "dsx1ChanMappingEntry" + types.AddKeyToken(dsx1ChanMappingEntry.IfIndex, "ifIndex") + types.AddKeyToken(dsx1ChanMappingEntry.Dsx1Ds1ChannelNumber, "dsx1Ds1ChannelNumber")
    dsx1ChanMappingEntry.EntityData.AbsolutePath = "DS1-MIB:DS1-MIB/dsx1ChanMappingTable/" + dsx1ChanMappingEntry.EntityData.SegmentPath
    dsx1ChanMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx1ChanMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx1ChanMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx1ChanMappingEntry.EntityData.Children = types.NewOrderedMap()
    dsx1ChanMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx1ChanMappingEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", dsx1ChanMappingEntry.IfIndex})
    dsx1ChanMappingEntry.EntityData.Leafs.Append("dsx1Ds1ChannelNumber", types.YLeaf{"Dsx1Ds1ChannelNumber", dsx1ChanMappingEntry.Dsx1Ds1ChannelNumber})
    dsx1ChanMappingEntry.EntityData.Leafs.Append("dsx1ChanMappedIfIndex", types.YLeaf{"Dsx1ChanMappedIfIndex", dsx1ChanMappingEntry.Dsx1ChanMappedIfIndex})

    dsx1ChanMappingEntry.EntityData.YListKeys = []string {"IfIndex", "Dsx1Ds1ChannelNumber"}

    return &(dsx1ChanMappingEntry.EntityData)
}

