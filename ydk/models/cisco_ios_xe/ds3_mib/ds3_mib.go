// The is the MIB module that describes
// DS3 and E3 interfaces objects.
package ds3_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ds3_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DS3-MIB DS3-MIB}", reflect.TypeOf(DS3MIB{}))
    ydk.RegisterEntity("DS3-MIB:DS3-MIB", reflect.TypeOf(DS3MIB{}))
}

// DS3MIB
type DS3MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The DS3/E3 Configuration table.
    Dsx3ConfigTable DS3MIB_Dsx3ConfigTable

    // The DS3/E3 current table contains various statistics being collected for
    // the current 15 minute interval.
    Dsx3CurrentTable DS3MIB_Dsx3CurrentTable

    // The DS3/E3 Interval Table contains various statistics collected by each
    // DS3/E3 Interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals.  Each row in this
    // table represents one such interval (identified by dsx3IntervalNumber) and
    // for one specific interface (identifed by dsx3IntervalIndex).
    Dsx3IntervalTable DS3MIB_Dsx3IntervalTable

    // The DS3/E3 Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx3TotalTable DS3MIB_Dsx3TotalTable

    // The DS3 Far End Configuration Table contains configuration information
    // reported in the C-bits from the remote end.
    Dsx3FarEndConfigTable DS3MIB_Dsx3FarEndConfigTable

    // The DS3 Far End Current table contains various statistics being collected
    // for the current 15 minute interval.  The statistics are collected from the
    // far end block error code within the C- bits.
    Dsx3FarEndCurrentTable DS3MIB_Dsx3FarEndCurrentTable

    // The DS3 Far End Interval Table contains various statistics collected by
    // each DS3 interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals.
    Dsx3FarEndIntervalTable DS3MIB_Dsx3FarEndIntervalTable

    // The DS3 Far End Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx3FarEndTotalTable DS3MIB_Dsx3FarEndTotalTable

    // This table is deprecated in favour of using ifStackTable.  Implementation
    // of this table was optional.  It was designed for those systems dividing a
    // DS3/E3 into channels containing different data streams that are of local
    // interest.  The DS3/E3 fractional table identifies which DS3/E3 channels
    // associated with a CSU are being used to support a logical interface, i.e.,
    // an entry in the interfaces table from the Internet- standard MIB.  For
    // example, consider a DS3 device with 4 high speed links carrying router
    // traffic, a feed for voice, a feed for video, and a synchronous channel for
    // a non-routed protocol.  We might describe the allocation of channels, in
    // the dsx3FracTable, as follows: dsx3FracIfIndex.2. 1 = 3 
    // dsx3FracIfIndex.2.15 = 4 dsx3FracIfIndex.2. 2 = 3  dsx3FracIfIndex.2.16 = 6
    // dsx3FracIfIndex.2. 3 = 3  dsx3FracIfIndex.2.17 = 6 dsx3FracIfIndex.2. 4 = 3
    // dsx3FracIfIndex.2.18 = 6 dsx3FracIfIndex.2. 5 = 3  dsx3FracIfIndex.2.19 = 6
    // dsx3FracIfIndex.2. 6 = 3  dsx3FracIfIndex.2.20 = 6 dsx3FracIfIndex.2. 7 = 4
    // dsx3FracIfIndex.2.21 = 6 dsx3FracIfIndex.2. 8 = 4  dsx3FracIfIndex.2.22 = 6
    // dsx3FracIfIndex.2. 9 = 4  dsx3FracIfIndex.2.23 = 6 dsx3FracIfIndex.2.10 = 4
    // dsx3FracIfIndex.2.24 = 6 dsx3FracIfIndex.2.11 = 4  dsx3FracIfIndex.2.25 = 6
    // dsx3FracIfIndex.2.12 = 5  dsx3FracIfIndex.2.26 = 6 dsx3FracIfIndex.2.13 = 5
    // dsx3FracIfIndex.2.27 = 6 dsx3FracIfIndex.2.14 = 5  dsx3FracIfIndex.2.28 = 6
    // For dsx3M23, dsx3 SYNTRAN, dsx3CbitParity, and dsx3ClearChannel  there are
    // 28 legal channels, numbered 1 throug h 28.  For e3Framed there are 16 legal
    // channels, numbered 1 through 16.  The channels (1..16) correspond directly
    // to the equivalently numbered time-slots.
    Dsx3FracTable DS3MIB_Dsx3FracTable
}

func (dS3MIB *DS3MIB) GetEntityData() *types.CommonEntityData {
    dS3MIB.EntityData.YFilter = dS3MIB.YFilter
    dS3MIB.EntityData.YangName = "DS3-MIB"
    dS3MIB.EntityData.BundleName = "cisco_ios_xe"
    dS3MIB.EntityData.ParentYangName = "DS3-MIB"
    dS3MIB.EntityData.SegmentPath = "DS3-MIB:DS3-MIB"
    dS3MIB.EntityData.AbsolutePath = dS3MIB.EntityData.SegmentPath
    dS3MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dS3MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dS3MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dS3MIB.EntityData.Children = types.NewOrderedMap()
    dS3MIB.EntityData.Children.Append("dsx3ConfigTable", types.YChild{"Dsx3ConfigTable", &dS3MIB.Dsx3ConfigTable})
    dS3MIB.EntityData.Children.Append("dsx3CurrentTable", types.YChild{"Dsx3CurrentTable", &dS3MIB.Dsx3CurrentTable})
    dS3MIB.EntityData.Children.Append("dsx3IntervalTable", types.YChild{"Dsx3IntervalTable", &dS3MIB.Dsx3IntervalTable})
    dS3MIB.EntityData.Children.Append("dsx3TotalTable", types.YChild{"Dsx3TotalTable", &dS3MIB.Dsx3TotalTable})
    dS3MIB.EntityData.Children.Append("dsx3FarEndConfigTable", types.YChild{"Dsx3FarEndConfigTable", &dS3MIB.Dsx3FarEndConfigTable})
    dS3MIB.EntityData.Children.Append("dsx3FarEndCurrentTable", types.YChild{"Dsx3FarEndCurrentTable", &dS3MIB.Dsx3FarEndCurrentTable})
    dS3MIB.EntityData.Children.Append("dsx3FarEndIntervalTable", types.YChild{"Dsx3FarEndIntervalTable", &dS3MIB.Dsx3FarEndIntervalTable})
    dS3MIB.EntityData.Children.Append("dsx3FarEndTotalTable", types.YChild{"Dsx3FarEndTotalTable", &dS3MIB.Dsx3FarEndTotalTable})
    dS3MIB.EntityData.Children.Append("dsx3FracTable", types.YChild{"Dsx3FracTable", &dS3MIB.Dsx3FracTable})
    dS3MIB.EntityData.Leafs = types.NewOrderedMap()

    dS3MIB.EntityData.YListKeys = []string {}

    return &(dS3MIB.EntityData)
}

// DS3MIB_Dsx3ConfigTable
// The DS3/E3 Configuration table.
type DS3MIB_Dsx3ConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Configuration table. The type is slice of
    // DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry.
    Dsx3ConfigEntry []*DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry
}

func (dsx3ConfigTable *DS3MIB_Dsx3ConfigTable) GetEntityData() *types.CommonEntityData {
    dsx3ConfigTable.EntityData.YFilter = dsx3ConfigTable.YFilter
    dsx3ConfigTable.EntityData.YangName = "dsx3ConfigTable"
    dsx3ConfigTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3ConfigTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3ConfigTable.EntityData.SegmentPath = "dsx3ConfigTable"
    dsx3ConfigTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3ConfigTable.EntityData.SegmentPath
    dsx3ConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3ConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3ConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3ConfigTable.EntityData.Children = types.NewOrderedMap()
    dsx3ConfigTable.EntityData.Children.Append("dsx3ConfigEntry", types.YChild{"Dsx3ConfigEntry", nil})
    for i := range dsx3ConfigTable.Dsx3ConfigEntry {
        dsx3ConfigTable.EntityData.Children.Append(types.GetSegmentPath(dsx3ConfigTable.Dsx3ConfigEntry[i]), types.YChild{"Dsx3ConfigEntry", dsx3ConfigTable.Dsx3ConfigEntry[i]})
    }
    dsx3ConfigTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3ConfigTable.EntityData.YListKeys = []string {}

    return &(dsx3ConfigTable.EntityData)
}

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry
// An entry in the DS3/E3 Configuration table.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object should be made equal to ifIndex.  The
    // next paragraph describes its previous usage. Making the object equal to
    // ifIndex allows propoer use of ifStackTable.  Previously, this object was
    // the identifier of a DS3/E3 Interface on a managed device.  If there is an
    // ifEntry that is directly associated with this and only this DS3/E3
    // interface, it should have the same value as ifIndex.  Otherwise, number the
    // dsx3LineIndices with an unique identifier following the rules of choosing a
    // number that is greater than ifNumber and numbering the inside interfaces
    // (e.g., equipment side) with even numbers and outside interfaces (e.g,
    // network side) with odd numbers. The type is interface{} with range:
    // 1..2147483647.
    Dsx3LineIndex interface{}

    // This value for this object is equal to the value of ifIndex from the
    // Interfaces table of MIB II (RFC 1213). The type is interface{} with range:
    // 1..2147483647.
    Dsx3IfIndex interface{}

    // The number of seconds that have elapsed since the beginning of the near end
    // current error- measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx3TimeElapsed interface{}

    // The number of previous near end intervals for which data was collected. 
    // The value will be 96 unless the interface was brought online within the
    // last 24 hours, in which case the value will be the number of complete 15
    // minute near end intervals since the interface has been online.  In the case
    // where the agent is a proxy, it is possible that some intervals are
    // unavailable.  In this case, this interval is the maximum interval number
    // for which data is available. The type is interface{} with range: 0..96.
    Dsx3ValidIntervals interface{}

    // This variable indicates the variety of DS3 C-bit or E3 application
    // implementing this interface. The type of interface affects the
    // interpretation of the usage and error statistics.  The rate of DS3 is
    // 44.736 Mbps and E3 is 34.368 Mbps.  The dsx3ClearChannel value means that
    // the C-bits are not used except for sending/receiving AIS. The values, in
    // sequence, describe:  TITLE:            SPECIFICATION: dsx3M23           
    // ANSI T1.107-1988 [9] dsx3SYNTRAN        ANSI T1.107-1988 [9] dsx3CbitParity
    // ANSI T1.107a-1990 [9a] dsx3ClearChannel   ANSI T1.102-1987 [8] e3Framed    
    // CCITT G.751 [12] e3Plcp             ETSI T/NA(91)18 [13]. The type is
    // Dsx3LineType.
    Dsx3LineType interface{}

    // This variable describes the variety of Zero Code Suppression used on this
    // interface, which in turn affects a number of its characteristics.  dsx3B3ZS
    // and e3HDB3 refer to the use of specified patterns of normal bits and
    // bipolar violations which are used to replace sequences of zero bits of a
    // specified length. The type is Dsx3LineCoding.
    Dsx3LineCoding interface{}

    // This variable indicates what type of code is being sent across the DS3/E3
    // interface by the device.  (These are optional for E3 interfaces.) Setting
    // this variable causes the interface to begin sending the code requested. The
    // values mean:     dsx3SendNoCode        sending looped or normal data    
    // dsx3SendLineCode        sending a request for a line loopback    
    // dsx3SendPayloadCode        sending a request for a payload loopback       
    // (i.e., all DS1/E1s in a DS3/E3 frame)     dsx3SendResetCode        sending
    // a loopback deactivation request     dsx3SendDS1LoopCode        requesting
    // to loopback a particular DS1/E1        within a DS3/E3 frame.  The DS1/E1
    // is        indicated in dsx3Ds1ForRemoteLoop.     dsx3SendTestPattern       
    // sending a test pattern. The type is Dsx3SendCode.
    Dsx3SendCode interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. The type is string with
    // length: 0..255.
    Dsx3CircuitIdentifier interface{}

    // This variable represents the desired loopback configuration of the DS3/E3
    // interface.  The values mean:  dsx3NoLoop   Not in the loopback state.  A
    // device that is   not capable of performing a loopback on   the interface
    // shall always return this as   its value.  dsx3PayloadLoop   The received
    // signal at this interface is looped   through the device.  Typically the
    // received signal   is looped back for retransmission after it has   passed
    // through the device's framing function.  dsx3LineLoop   The received signal
    // at this interface does not   go through the device (minimum penetration)
    // but   is looped back out.  dsx3OtherLoop   Loopbacks that are not defined
    // here.  dsx3InwardLoop   The sent signal at this interface is looped back  
    // through the device.  dsx3DualLoop   Both dsx1LineLoop and dsx1InwardLoop
    // will be   active simultaneously. The type is Dsx3LoopbackConfig.
    Dsx3LoopbackConfig interface{}

    // This variable indicates the Line Status of the interface.  It contains
    // loopback state information and failure state information.  The
    // dsx3LineStatus is a bit map represented as a sum, therefore, it can
    // represent multiple failures and a loopback (see dsx3LoopbackConfig object
    // for the type of loopback) simultaneously.  The dsx3NoAlarm must be set if
    // and only if no other flag is set.  If the dsx3loopbackState bit is set, the
    // loopback in effect can be determined from the dsx3loopbackConfig object.
    // The various bit positions are: 1     dsx3NoAlarm         No alarm present 2
    // dsx3RcvRAIFailure   Receiving Yellow/Remote                  Alarm
    // Indication 4     dsx3XmitRAIAlarm    Transmitting Yellow/Remote            
    // Alarm Indication 8     dsx3RcvAIS          Receiving AIS failure state 16  
    // dsx3XmitAIS         Transmitting AIS 32     dsx3LOF             Receiving
    // LOF failure state 64     dsx3LOS             Receiving LOS failure state
    // 128     dsx3LoopbackState   Looping the received signal 256    
    // dsx3RcvTestCode     Receiving a Test Pattern 512     dsx3OtherFailure   
    // any line status not defined                  here 1024    
    // dsx3UnavailSigState Near End in Unavailable Signal                  State
    // 2048     dsx3NetEquipOOS     Carrier Equipment Out of Service. The type is
    // interface{} with range: 1..4095.
    Dsx3LineStatus interface{}

    // The source of Transmit Clock.  loopTiming indicates that the recovered
    // receive clock is used as the transmit clock.  localTiming indicates that a
    // local clock source is used or that an external clock is attached to the box
    // containing the interface.  throughTiming indicates that transmit clock is
    // derived from the recovered receive clock of another DS3 interface. The type
    // is Dsx3TransmitClockSource.
    Dsx3TransmitClockSource interface{}

    // The number of intervals in the range from 0 to dsx3ValidIntervals for which
    // no data is available.  This object will typically be zero except in cases
    // where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx3InvalidIntervals interface{}

    // The length of the ds3 line in meters.  This object provides information for
    // line build out circuitry if it exists and can use this object to adjust the
    // line build out. The type is interface{} with range: 0..64000. Units are
    // meters.
    Dsx3LineLength interface{}

    // The value of MIB II's sysUpTime object at the time this DS3/E3 entered its
    // current line status state.  If the current state was entered prior to the
    // last re-initialization of the proxy-agent, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    Dsx3LineStatusLastChange interface{}

    // Indicates whether dsx3LineStatusChange traps should be generated for this
    // interface. The type is Dsx3LineStatusChangeTrapEnable.
    Dsx3LineStatusChangeTrapEnable interface{}

    // This variable represents the current state of the loopback on the DS3
    // interface.  It contains information about loopbacks established by a
    // manager and remotely from the far end.  The dsx3LoopbackStatus is a bit map
    // represented as a sum, therefore is can represent multiple loopbacks
    // simultaneously.  The various bit positions are:  1  dsx3NoLoopback  2 
    // dsx3NearEndPayloadLoopback  4  dsx3NearEndLineLoopback  8 
    // dsx3NearEndOtherLoopback 16  dsx3NearEndInwardLoopback 32 
    // dsx3FarEndPayloadLoopback 64  dsx3FarEndLineLoopback. The type is
    // interface{} with range: 1..127.
    Dsx3LoopbackStatus interface{}

    // Indicates whether this ds3/e3 is channelized or unchannelized.  The value
    // of enabledDs1 indicates that this is a DS3 channelized into DS1s.  The
    // value of enabledDs3 indicated that this is a DS3 channelized into DS2s. 
    // Setting this object will cause the creation or deletion of DS2 or DS1
    // entries in the ifTable.  . The type is Dsx3Channelization.
    Dsx3Channelization interface{}

    // Indicates which ds1/e1 on this ds3/e3 will be indicated in the remote ds1
    // loopback request.  A value of 0 means no DS1 will be looped.  A value of 29
    // means all ds1s/e1s will be looped. The type is interface{} with range:
    // 0..29.
    Dsx3Ds1ForRemoteLoop interface{}
}

func (dsx3ConfigEntry *DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry) GetEntityData() *types.CommonEntityData {
    dsx3ConfigEntry.EntityData.YFilter = dsx3ConfigEntry.YFilter
    dsx3ConfigEntry.EntityData.YangName = "dsx3ConfigEntry"
    dsx3ConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3ConfigEntry.EntityData.ParentYangName = "dsx3ConfigTable"
    dsx3ConfigEntry.EntityData.SegmentPath = "dsx3ConfigEntry" + types.AddKeyToken(dsx3ConfigEntry.Dsx3LineIndex, "dsx3LineIndex")
    dsx3ConfigEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3ConfigTable/" + dsx3ConfigEntry.EntityData.SegmentPath
    dsx3ConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3ConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3ConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3ConfigEntry.EntityData.Children = types.NewOrderedMap()
    dsx3ConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineIndex", types.YLeaf{"Dsx3LineIndex", dsx3ConfigEntry.Dsx3LineIndex})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3IfIndex", types.YLeaf{"Dsx3IfIndex", dsx3ConfigEntry.Dsx3IfIndex})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3TimeElapsed", types.YLeaf{"Dsx3TimeElapsed", dsx3ConfigEntry.Dsx3TimeElapsed})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3ValidIntervals", types.YLeaf{"Dsx3ValidIntervals", dsx3ConfigEntry.Dsx3ValidIntervals})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineType", types.YLeaf{"Dsx3LineType", dsx3ConfigEntry.Dsx3LineType})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineCoding", types.YLeaf{"Dsx3LineCoding", dsx3ConfigEntry.Dsx3LineCoding})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3SendCode", types.YLeaf{"Dsx3SendCode", dsx3ConfigEntry.Dsx3SendCode})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3CircuitIdentifier", types.YLeaf{"Dsx3CircuitIdentifier", dsx3ConfigEntry.Dsx3CircuitIdentifier})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LoopbackConfig", types.YLeaf{"Dsx3LoopbackConfig", dsx3ConfigEntry.Dsx3LoopbackConfig})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineStatus", types.YLeaf{"Dsx3LineStatus", dsx3ConfigEntry.Dsx3LineStatus})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3TransmitClockSource", types.YLeaf{"Dsx3TransmitClockSource", dsx3ConfigEntry.Dsx3TransmitClockSource})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3InvalidIntervals", types.YLeaf{"Dsx3InvalidIntervals", dsx3ConfigEntry.Dsx3InvalidIntervals})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineLength", types.YLeaf{"Dsx3LineLength", dsx3ConfigEntry.Dsx3LineLength})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineStatusLastChange", types.YLeaf{"Dsx3LineStatusLastChange", dsx3ConfigEntry.Dsx3LineStatusLastChange})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LineStatusChangeTrapEnable", types.YLeaf{"Dsx3LineStatusChangeTrapEnable", dsx3ConfigEntry.Dsx3LineStatusChangeTrapEnable})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3LoopbackStatus", types.YLeaf{"Dsx3LoopbackStatus", dsx3ConfigEntry.Dsx3LoopbackStatus})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3Channelization", types.YLeaf{"Dsx3Channelization", dsx3ConfigEntry.Dsx3Channelization})
    dsx3ConfigEntry.EntityData.Leafs.Append("dsx3Ds1ForRemoteLoop", types.YLeaf{"Dsx3Ds1ForRemoteLoop", dsx3ConfigEntry.Dsx3Ds1ForRemoteLoop})

    dsx3ConfigEntry.EntityData.YListKeys = []string {"Dsx3LineIndex"}

    return &(dsx3ConfigEntry.EntityData)
}

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization represents entries in the ifTable.  
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization_disabled DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization = "disabled"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization_enabledDs1 DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization = "enabledDs1"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization_enabledDs2 DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3Channelization = "enabledDs2"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding represents of a specified length.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding_dsx3Other DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding = "dsx3Other"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding_dsx3B3ZS DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding = "dsx3B3ZS"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding_e3HDB3 DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineCoding = "e3HDB3"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable represents should be generated for this interface.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable_enabled DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable = "enabled"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable_disabled DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineStatusChangeTrapEnable = "disabled"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType represents e3Plcp             ETSI T/NA(91)18 [13].
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_dsx3other DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "dsx3other"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_dsx3M23 DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "dsx3M23"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_dsx3SYNTRAN DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "dsx3SYNTRAN"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_dsx3CbitParity DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "dsx3CbitParity"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_dsx3ClearChannel DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "dsx3ClearChannel"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_e3other DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "e3other"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_e3Framed DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "e3Framed"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType_e3Plcp DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LineType = "e3Plcp"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig represents   active simultaneously.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3NoLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3NoLoop"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3PayloadLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3PayloadLoop"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3LineLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3LineLoop"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3OtherLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3OtherLoop"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3InwardLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3InwardLoop"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig_dsx3DualLoop DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3LoopbackConfig = "dsx3DualLoop"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode represents        sending a test pattern.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendNoCode DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendNoCode"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendLineCode DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendLineCode"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendPayloadCode DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendPayloadCode"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendResetCode DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendResetCode"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendDS1LoopCode DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendDS1LoopCode"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode_dsx3SendTestPattern DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3SendCode = "dsx3SendTestPattern"
)

// DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource represents interface.
type DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource string

const (
    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource_loopTiming DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource = "loopTiming"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource_localTiming DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource = "localTiming"

    DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource_throughTiming DS3MIB_Dsx3ConfigTable_Dsx3ConfigEntry_Dsx3TransmitClockSource = "throughTiming"
)

// DS3MIB_Dsx3CurrentTable
// The DS3/E3 current table contains various
// statistics being collected for the current 15
// minute interval.
type DS3MIB_Dsx3CurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Current table. The type is slice of
    // DS3MIB_Dsx3CurrentTable_Dsx3CurrentEntry.
    Dsx3CurrentEntry []*DS3MIB_Dsx3CurrentTable_Dsx3CurrentEntry
}

func (dsx3CurrentTable *DS3MIB_Dsx3CurrentTable) GetEntityData() *types.CommonEntityData {
    dsx3CurrentTable.EntityData.YFilter = dsx3CurrentTable.YFilter
    dsx3CurrentTable.EntityData.YangName = "dsx3CurrentTable"
    dsx3CurrentTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3CurrentTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3CurrentTable.EntityData.SegmentPath = "dsx3CurrentTable"
    dsx3CurrentTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3CurrentTable.EntityData.SegmentPath
    dsx3CurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3CurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3CurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3CurrentTable.EntityData.Children = types.NewOrderedMap()
    dsx3CurrentTable.EntityData.Children.Append("dsx3CurrentEntry", types.YChild{"Dsx3CurrentEntry", nil})
    for i := range dsx3CurrentTable.Dsx3CurrentEntry {
        dsx3CurrentTable.EntityData.Children.Append(types.GetSegmentPath(dsx3CurrentTable.Dsx3CurrentEntry[i]), types.YChild{"Dsx3CurrentEntry", dsx3CurrentTable.Dsx3CurrentEntry[i]})
    }
    dsx3CurrentTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3CurrentTable.EntityData.YListKeys = []string {}

    return &(dsx3CurrentTable.EntityData)
}

// DS3MIB_Dsx3CurrentTable_Dsx3CurrentEntry
// An entry in the DS3/E3 Current table.
type DS3MIB_Dsx3CurrentTable_Dsx3CurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3CurrentIndex interface{}

    // The counter associated with the number of P-bit Errored Seconds. The type
    // is interface{} with range: 0..4294967295.
    Dsx3CurrentPESs interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3CurrentPSESs interface{}

    // The counter associated with the number of Severely Errored Framing Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3CurrentSEFSs interface{}

    // The counter associated with the number of Unavailable Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx3CurrentUASs interface{}

    // The counter associated with the number of Line Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3CurrentLCVs interface{}

    // The counter associated with the number of P-bit Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3CurrentPCVs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3CurrentLESs interface{}

    // The number of C-bit Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx3CurrentCCVs interface{}

    // The number of C-bit Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3CurrentCESs interface{}

    // The number of C-bit Severely Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx3CurrentCSESs interface{}
}

func (dsx3CurrentEntry *DS3MIB_Dsx3CurrentTable_Dsx3CurrentEntry) GetEntityData() *types.CommonEntityData {
    dsx3CurrentEntry.EntityData.YFilter = dsx3CurrentEntry.YFilter
    dsx3CurrentEntry.EntityData.YangName = "dsx3CurrentEntry"
    dsx3CurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3CurrentEntry.EntityData.ParentYangName = "dsx3CurrentTable"
    dsx3CurrentEntry.EntityData.SegmentPath = "dsx3CurrentEntry" + types.AddKeyToken(dsx3CurrentEntry.Dsx3CurrentIndex, "dsx3CurrentIndex")
    dsx3CurrentEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3CurrentTable/" + dsx3CurrentEntry.EntityData.SegmentPath
    dsx3CurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3CurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3CurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3CurrentEntry.EntityData.Children = types.NewOrderedMap()
    dsx3CurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentIndex", types.YLeaf{"Dsx3CurrentIndex", dsx3CurrentEntry.Dsx3CurrentIndex})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentPESs", types.YLeaf{"Dsx3CurrentPESs", dsx3CurrentEntry.Dsx3CurrentPESs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentPSESs", types.YLeaf{"Dsx3CurrentPSESs", dsx3CurrentEntry.Dsx3CurrentPSESs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentSEFSs", types.YLeaf{"Dsx3CurrentSEFSs", dsx3CurrentEntry.Dsx3CurrentSEFSs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentUASs", types.YLeaf{"Dsx3CurrentUASs", dsx3CurrentEntry.Dsx3CurrentUASs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentLCVs", types.YLeaf{"Dsx3CurrentLCVs", dsx3CurrentEntry.Dsx3CurrentLCVs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentPCVs", types.YLeaf{"Dsx3CurrentPCVs", dsx3CurrentEntry.Dsx3CurrentPCVs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentLESs", types.YLeaf{"Dsx3CurrentLESs", dsx3CurrentEntry.Dsx3CurrentLESs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentCCVs", types.YLeaf{"Dsx3CurrentCCVs", dsx3CurrentEntry.Dsx3CurrentCCVs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentCESs", types.YLeaf{"Dsx3CurrentCESs", dsx3CurrentEntry.Dsx3CurrentCESs})
    dsx3CurrentEntry.EntityData.Leafs.Append("dsx3CurrentCSESs", types.YLeaf{"Dsx3CurrentCSESs", dsx3CurrentEntry.Dsx3CurrentCSESs})

    dsx3CurrentEntry.EntityData.YListKeys = []string {"Dsx3CurrentIndex"}

    return &(dsx3CurrentEntry.EntityData)
}

// DS3MIB_Dsx3IntervalTable
// The DS3/E3 Interval Table contains various
// statistics collected by each DS3/E3 Interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.  Each row in this table represents one
// such interval (identified by dsx3IntervalNumber)
// and for one specific interface (identifed by
// dsx3IntervalIndex).
type DS3MIB_Dsx3IntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Interval table. The type is slice of
    // DS3MIB_Dsx3IntervalTable_Dsx3IntervalEntry.
    Dsx3IntervalEntry []*DS3MIB_Dsx3IntervalTable_Dsx3IntervalEntry
}

func (dsx3IntervalTable *DS3MIB_Dsx3IntervalTable) GetEntityData() *types.CommonEntityData {
    dsx3IntervalTable.EntityData.YFilter = dsx3IntervalTable.YFilter
    dsx3IntervalTable.EntityData.YangName = "dsx3IntervalTable"
    dsx3IntervalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3IntervalTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3IntervalTable.EntityData.SegmentPath = "dsx3IntervalTable"
    dsx3IntervalTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3IntervalTable.EntityData.SegmentPath
    dsx3IntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3IntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3IntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3IntervalTable.EntityData.Children = types.NewOrderedMap()
    dsx3IntervalTable.EntityData.Children.Append("dsx3IntervalEntry", types.YChild{"Dsx3IntervalEntry", nil})
    for i := range dsx3IntervalTable.Dsx3IntervalEntry {
        dsx3IntervalTable.EntityData.Children.Append(types.GetSegmentPath(dsx3IntervalTable.Dsx3IntervalEntry[i]), types.YChild{"Dsx3IntervalEntry", dsx3IntervalTable.Dsx3IntervalEntry[i]})
    }
    dsx3IntervalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3IntervalTable.EntityData.YListKeys = []string {}

    return &(dsx3IntervalTable.EntityData)
}

// DS3MIB_Dsx3IntervalTable_Dsx3IntervalEntry
// An entry in the DS3/E3 Interval table.
type DS3MIB_Dsx3IntervalTable_Dsx3IntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3IntervalIndex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx3IntervalNumber interface{}

    // The counter associated with the number of P-bit Errored Seconds. The type
    // is interface{} with range: 0..4294967295.
    Dsx3IntervalPESs interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3IntervalPSESs interface{}

    // The counter associated with the number of Severely Errored Framing Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3IntervalSEFSs interface{}

    // The counter associated with the number of Unavailable Seconds.  This object
    // may decrease if the occurance of unavailable seconds occurs across an
    // inteval boundary. The type is interface{} with range: 0..4294967295.
    Dsx3IntervalUASs interface{}

    // The counter associated with the number of Line Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3IntervalLCVs interface{}

    // The counter associated with the number of P-bit Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3IntervalPCVs interface{}

    // The number of Line Errored  Seconds  (BPVs  or illegal  zero  sequences).
    // The type is interface{} with range: 0..4294967295.
    Dsx3IntervalLESs interface{}

    // The number of C-bit Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx3IntervalCCVs interface{}

    // The number of C-bit Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3IntervalCESs interface{}

    // The number of C-bit Severely Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx3IntervalCSESs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx3IntervalValidData interface{}
}

func (dsx3IntervalEntry *DS3MIB_Dsx3IntervalTable_Dsx3IntervalEntry) GetEntityData() *types.CommonEntityData {
    dsx3IntervalEntry.EntityData.YFilter = dsx3IntervalEntry.YFilter
    dsx3IntervalEntry.EntityData.YangName = "dsx3IntervalEntry"
    dsx3IntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3IntervalEntry.EntityData.ParentYangName = "dsx3IntervalTable"
    dsx3IntervalEntry.EntityData.SegmentPath = "dsx3IntervalEntry" + types.AddKeyToken(dsx3IntervalEntry.Dsx3IntervalIndex, "dsx3IntervalIndex") + types.AddKeyToken(dsx3IntervalEntry.Dsx3IntervalNumber, "dsx3IntervalNumber")
    dsx3IntervalEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3IntervalTable/" + dsx3IntervalEntry.EntityData.SegmentPath
    dsx3IntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3IntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3IntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3IntervalEntry.EntityData.Children = types.NewOrderedMap()
    dsx3IntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalIndex", types.YLeaf{"Dsx3IntervalIndex", dsx3IntervalEntry.Dsx3IntervalIndex})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalNumber", types.YLeaf{"Dsx3IntervalNumber", dsx3IntervalEntry.Dsx3IntervalNumber})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalPESs", types.YLeaf{"Dsx3IntervalPESs", dsx3IntervalEntry.Dsx3IntervalPESs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalPSESs", types.YLeaf{"Dsx3IntervalPSESs", dsx3IntervalEntry.Dsx3IntervalPSESs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalSEFSs", types.YLeaf{"Dsx3IntervalSEFSs", dsx3IntervalEntry.Dsx3IntervalSEFSs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalUASs", types.YLeaf{"Dsx3IntervalUASs", dsx3IntervalEntry.Dsx3IntervalUASs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalLCVs", types.YLeaf{"Dsx3IntervalLCVs", dsx3IntervalEntry.Dsx3IntervalLCVs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalPCVs", types.YLeaf{"Dsx3IntervalPCVs", dsx3IntervalEntry.Dsx3IntervalPCVs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalLESs", types.YLeaf{"Dsx3IntervalLESs", dsx3IntervalEntry.Dsx3IntervalLESs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalCCVs", types.YLeaf{"Dsx3IntervalCCVs", dsx3IntervalEntry.Dsx3IntervalCCVs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalCESs", types.YLeaf{"Dsx3IntervalCESs", dsx3IntervalEntry.Dsx3IntervalCESs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalCSESs", types.YLeaf{"Dsx3IntervalCSESs", dsx3IntervalEntry.Dsx3IntervalCSESs})
    dsx3IntervalEntry.EntityData.Leafs.Append("dsx3IntervalValidData", types.YLeaf{"Dsx3IntervalValidData", dsx3IntervalEntry.Dsx3IntervalValidData})

    dsx3IntervalEntry.EntityData.YListKeys = []string {"Dsx3IntervalIndex", "Dsx3IntervalNumber"}

    return &(dsx3IntervalEntry.EntityData)
}

// DS3MIB_Dsx3TotalTable
// The DS3/E3 Total Table contains the cumulative
// sum of the various statistics for the 24 hour
// period preceding the current interval.
type DS3MIB_Dsx3TotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Total table. The type is slice of
    // DS3MIB_Dsx3TotalTable_Dsx3TotalEntry.
    Dsx3TotalEntry []*DS3MIB_Dsx3TotalTable_Dsx3TotalEntry
}

func (dsx3TotalTable *DS3MIB_Dsx3TotalTable) GetEntityData() *types.CommonEntityData {
    dsx3TotalTable.EntityData.YFilter = dsx3TotalTable.YFilter
    dsx3TotalTable.EntityData.YangName = "dsx3TotalTable"
    dsx3TotalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3TotalTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3TotalTable.EntityData.SegmentPath = "dsx3TotalTable"
    dsx3TotalTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3TotalTable.EntityData.SegmentPath
    dsx3TotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3TotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3TotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3TotalTable.EntityData.Children = types.NewOrderedMap()
    dsx3TotalTable.EntityData.Children.Append("dsx3TotalEntry", types.YChild{"Dsx3TotalEntry", nil})
    for i := range dsx3TotalTable.Dsx3TotalEntry {
        dsx3TotalTable.EntityData.Children.Append(types.GetSegmentPath(dsx3TotalTable.Dsx3TotalEntry[i]), types.YChild{"Dsx3TotalEntry", dsx3TotalTable.Dsx3TotalEntry[i]})
    }
    dsx3TotalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3TotalTable.EntityData.YListKeys = []string {}

    return &(dsx3TotalTable.EntityData)
}

// DS3MIB_Dsx3TotalTable_Dsx3TotalEntry
// An entry in the DS3/E3 Total table.
type DS3MIB_Dsx3TotalTable_Dsx3TotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3TotalIndex interface{}

    // The counter associated with the number of P-bit Errored Seconds,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalPESs interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalPSESs interface{}

    // The counter associated with the number of Severely Errored Framing Seconds,
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalSEFSs interface{}

    // The counter associated with the number of Unavailable Seconds, encountered
    // by a DS3 interface in the previous 24 hour interval.  Invalid 15 minute
    // intervals count as 0. The type is interface{} with range: 0..4294967295.
    Dsx3TotalUASs interface{}

    // The counter associated with the number of Line Coding Violations
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalLCVs interface{}

    // The counter associated with the number of P-bit Coding Violations,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalPCVs interface{}

    // The number of Line Errored  Seconds  (BPVs  or illegal  zero  sequences)
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3TotalLESs interface{}

    // The number of C-bit Coding Violations encountered by a DS3 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx3TotalCCVs interface{}

    // The number of C-bit Errored Seconds encountered by a DS3 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx3TotalCESs interface{}

    // The number of C-bit Severely Errored Seconds encountered by a DS3 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx3TotalCSESs interface{}
}

func (dsx3TotalEntry *DS3MIB_Dsx3TotalTable_Dsx3TotalEntry) GetEntityData() *types.CommonEntityData {
    dsx3TotalEntry.EntityData.YFilter = dsx3TotalEntry.YFilter
    dsx3TotalEntry.EntityData.YangName = "dsx3TotalEntry"
    dsx3TotalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3TotalEntry.EntityData.ParentYangName = "dsx3TotalTable"
    dsx3TotalEntry.EntityData.SegmentPath = "dsx3TotalEntry" + types.AddKeyToken(dsx3TotalEntry.Dsx3TotalIndex, "dsx3TotalIndex")
    dsx3TotalEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3TotalTable/" + dsx3TotalEntry.EntityData.SegmentPath
    dsx3TotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3TotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3TotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3TotalEntry.EntityData.Children = types.NewOrderedMap()
    dsx3TotalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalIndex", types.YLeaf{"Dsx3TotalIndex", dsx3TotalEntry.Dsx3TotalIndex})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalPESs", types.YLeaf{"Dsx3TotalPESs", dsx3TotalEntry.Dsx3TotalPESs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalPSESs", types.YLeaf{"Dsx3TotalPSESs", dsx3TotalEntry.Dsx3TotalPSESs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalSEFSs", types.YLeaf{"Dsx3TotalSEFSs", dsx3TotalEntry.Dsx3TotalSEFSs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalUASs", types.YLeaf{"Dsx3TotalUASs", dsx3TotalEntry.Dsx3TotalUASs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalLCVs", types.YLeaf{"Dsx3TotalLCVs", dsx3TotalEntry.Dsx3TotalLCVs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalPCVs", types.YLeaf{"Dsx3TotalPCVs", dsx3TotalEntry.Dsx3TotalPCVs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalLESs", types.YLeaf{"Dsx3TotalLESs", dsx3TotalEntry.Dsx3TotalLESs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalCCVs", types.YLeaf{"Dsx3TotalCCVs", dsx3TotalEntry.Dsx3TotalCCVs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalCESs", types.YLeaf{"Dsx3TotalCESs", dsx3TotalEntry.Dsx3TotalCESs})
    dsx3TotalEntry.EntityData.Leafs.Append("dsx3TotalCSESs", types.YLeaf{"Dsx3TotalCSESs", dsx3TotalEntry.Dsx3TotalCSESs})

    dsx3TotalEntry.EntityData.YListKeys = []string {"Dsx3TotalIndex"}

    return &(dsx3TotalEntry.EntityData)
}

// DS3MIB_Dsx3FarEndConfigTable
// The DS3 Far End Configuration Table contains
// configuration information reported in the C-bits
// from the remote end.
type DS3MIB_Dsx3FarEndConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Configuration table. The type is slice of
    // DS3MIB_Dsx3FarEndConfigTable_Dsx3FarEndConfigEntry.
    Dsx3FarEndConfigEntry []*DS3MIB_Dsx3FarEndConfigTable_Dsx3FarEndConfigEntry
}

func (dsx3FarEndConfigTable *DS3MIB_Dsx3FarEndConfigTable) GetEntityData() *types.CommonEntityData {
    dsx3FarEndConfigTable.EntityData.YFilter = dsx3FarEndConfigTable.YFilter
    dsx3FarEndConfigTable.EntityData.YangName = "dsx3FarEndConfigTable"
    dsx3FarEndConfigTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndConfigTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3FarEndConfigTable.EntityData.SegmentPath = "dsx3FarEndConfigTable"
    dsx3FarEndConfigTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3FarEndConfigTable.EntityData.SegmentPath
    dsx3FarEndConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndConfigTable.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndConfigTable.EntityData.Children.Append("dsx3FarEndConfigEntry", types.YChild{"Dsx3FarEndConfigEntry", nil})
    for i := range dsx3FarEndConfigTable.Dsx3FarEndConfigEntry {
        dsx3FarEndConfigTable.EntityData.Children.Append(types.GetSegmentPath(dsx3FarEndConfigTable.Dsx3FarEndConfigEntry[i]), types.YChild{"Dsx3FarEndConfigEntry", dsx3FarEndConfigTable.Dsx3FarEndConfigEntry[i]})
    }
    dsx3FarEndConfigTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3FarEndConfigTable.EntityData.YListKeys = []string {}

    return &(dsx3FarEndConfigTable.EntityData)
}

// DS3MIB_Dsx3FarEndConfigTable_Dsx3FarEndConfigEntry
// An entry in the DS3 Far End Configuration table.
type DS3MIB_Dsx3FarEndConfigTable_Dsx3FarEndConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value an dsx3LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx3FarEndLineIndex interface{}

    // This is the Far End Equipment Identification code that describes the
    // specific piece of equipment. It is sent within the Path Identification
    // Message. The type is string with length: 0..10.
    Dsx3FarEndEquipCode interface{}

    // This is the Far End Location Identification code that describes the
    // specific location of the equipment.  It is sent within the Path
    // Identification Message. The type is string with length: 0..11.
    Dsx3FarEndLocationIDCode interface{}

    // This is the Far End Frame Identification code that identifies where the
    // equipment is located within a building at a given location.  It is sent
    // within the Path Identification Message. The type is string with length:
    // 0..10.
    Dsx3FarEndFrameIDCode interface{}

    // This is the Far End code that identifies the equipment location within a
    // bay.  It is sent within the Path Identification Message. The type is string
    // with length: 0..6.
    Dsx3FarEndUnitCode interface{}

    // This code identifies a specific Far End DS3 path. It is sent within the
    // Path Identification Message. The type is string with length: 0..38.
    Dsx3FarEndFacilityIDCode interface{}
}

func (dsx3FarEndConfigEntry *DS3MIB_Dsx3FarEndConfigTable_Dsx3FarEndConfigEntry) GetEntityData() *types.CommonEntityData {
    dsx3FarEndConfigEntry.EntityData.YFilter = dsx3FarEndConfigEntry.YFilter
    dsx3FarEndConfigEntry.EntityData.YangName = "dsx3FarEndConfigEntry"
    dsx3FarEndConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndConfigEntry.EntityData.ParentYangName = "dsx3FarEndConfigTable"
    dsx3FarEndConfigEntry.EntityData.SegmentPath = "dsx3FarEndConfigEntry" + types.AddKeyToken(dsx3FarEndConfigEntry.Dsx3FarEndLineIndex, "dsx3FarEndLineIndex")
    dsx3FarEndConfigEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3FarEndConfigTable/" + dsx3FarEndConfigEntry.EntityData.SegmentPath
    dsx3FarEndConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndConfigEntry.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndLineIndex", types.YLeaf{"Dsx3FarEndLineIndex", dsx3FarEndConfigEntry.Dsx3FarEndLineIndex})
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndEquipCode", types.YLeaf{"Dsx3FarEndEquipCode", dsx3FarEndConfigEntry.Dsx3FarEndEquipCode})
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndLocationIDCode", types.YLeaf{"Dsx3FarEndLocationIDCode", dsx3FarEndConfigEntry.Dsx3FarEndLocationIDCode})
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndFrameIDCode", types.YLeaf{"Dsx3FarEndFrameIDCode", dsx3FarEndConfigEntry.Dsx3FarEndFrameIDCode})
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndUnitCode", types.YLeaf{"Dsx3FarEndUnitCode", dsx3FarEndConfigEntry.Dsx3FarEndUnitCode})
    dsx3FarEndConfigEntry.EntityData.Leafs.Append("dsx3FarEndFacilityIDCode", types.YLeaf{"Dsx3FarEndFacilityIDCode", dsx3FarEndConfigEntry.Dsx3FarEndFacilityIDCode})

    dsx3FarEndConfigEntry.EntityData.YListKeys = []string {"Dsx3FarEndLineIndex"}

    return &(dsx3FarEndConfigEntry.EntityData)
}

// DS3MIB_Dsx3FarEndCurrentTable
// The DS3 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end block error code within the C-
// bits.
type DS3MIB_Dsx3FarEndCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Current table. The type is slice of
    // DS3MIB_Dsx3FarEndCurrentTable_Dsx3FarEndCurrentEntry.
    Dsx3FarEndCurrentEntry []*DS3MIB_Dsx3FarEndCurrentTable_Dsx3FarEndCurrentEntry
}

func (dsx3FarEndCurrentTable *DS3MIB_Dsx3FarEndCurrentTable) GetEntityData() *types.CommonEntityData {
    dsx3FarEndCurrentTable.EntityData.YFilter = dsx3FarEndCurrentTable.YFilter
    dsx3FarEndCurrentTable.EntityData.YangName = "dsx3FarEndCurrentTable"
    dsx3FarEndCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndCurrentTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3FarEndCurrentTable.EntityData.SegmentPath = "dsx3FarEndCurrentTable"
    dsx3FarEndCurrentTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3FarEndCurrentTable.EntityData.SegmentPath
    dsx3FarEndCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndCurrentTable.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndCurrentTable.EntityData.Children.Append("dsx3FarEndCurrentEntry", types.YChild{"Dsx3FarEndCurrentEntry", nil})
    for i := range dsx3FarEndCurrentTable.Dsx3FarEndCurrentEntry {
        dsx3FarEndCurrentTable.EntityData.Children.Append(types.GetSegmentPath(dsx3FarEndCurrentTable.Dsx3FarEndCurrentEntry[i]), types.YChild{"Dsx3FarEndCurrentEntry", dsx3FarEndCurrentTable.Dsx3FarEndCurrentEntry[i]})
    }
    dsx3FarEndCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3FarEndCurrentTable.EntityData.YListKeys = []string {}

    return &(dsx3FarEndCurrentTable.EntityData)
}

// DS3MIB_Dsx3FarEndCurrentTable_Dsx3FarEndCurrentEntry
// An entry in the DS3 Far End Current table.
type DS3MIB_Dsx3FarEndCurrentTable_Dsx3FarEndCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3FarEndCurrentIndex interface{}

    // The number of seconds that have elapsed since the beginning of the far end
    // current error-measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx3FarEndTimeElapsed interface{}

    // The number of previous far end intervals for which data was collected.  The
    // value will be 96 unless the interface was brought online within the last 24
    // hours, in which case the value will be the number of complete 15 minute far
    // end intervals since the interface has been online. The type is interface{}
    // with range: 0..96.
    Dsx3FarEndValidIntervals interface{}

    // The counter associated with the number of Far Far End C-bit Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3FarEndCurrentCESs interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3FarEndCurrentCSESs interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count. The type is interface{} with
    // range: 0..4294967295.
    Dsx3FarEndCurrentCCVs interface{}

    // The counter associated with the number of Far End unavailable seconds. The
    // type is interface{} with range: 0..4294967295.
    Dsx3FarEndCurrentUASs interface{}

    // The number of intervals in the range from 0 to dsx3FarEndValidIntervals for
    // which no data is available.  This object will typically be zero except in
    // cases where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx3FarEndInvalidIntervals interface{}
}

func (dsx3FarEndCurrentEntry *DS3MIB_Dsx3FarEndCurrentTable_Dsx3FarEndCurrentEntry) GetEntityData() *types.CommonEntityData {
    dsx3FarEndCurrentEntry.EntityData.YFilter = dsx3FarEndCurrentEntry.YFilter
    dsx3FarEndCurrentEntry.EntityData.YangName = "dsx3FarEndCurrentEntry"
    dsx3FarEndCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndCurrentEntry.EntityData.ParentYangName = "dsx3FarEndCurrentTable"
    dsx3FarEndCurrentEntry.EntityData.SegmentPath = "dsx3FarEndCurrentEntry" + types.AddKeyToken(dsx3FarEndCurrentEntry.Dsx3FarEndCurrentIndex, "dsx3FarEndCurrentIndex")
    dsx3FarEndCurrentEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3FarEndCurrentTable/" + dsx3FarEndCurrentEntry.EntityData.SegmentPath
    dsx3FarEndCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndCurrentEntry.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndCurrentIndex", types.YLeaf{"Dsx3FarEndCurrentIndex", dsx3FarEndCurrentEntry.Dsx3FarEndCurrentIndex})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndTimeElapsed", types.YLeaf{"Dsx3FarEndTimeElapsed", dsx3FarEndCurrentEntry.Dsx3FarEndTimeElapsed})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndValidIntervals", types.YLeaf{"Dsx3FarEndValidIntervals", dsx3FarEndCurrentEntry.Dsx3FarEndValidIntervals})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndCurrentCESs", types.YLeaf{"Dsx3FarEndCurrentCESs", dsx3FarEndCurrentEntry.Dsx3FarEndCurrentCESs})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndCurrentCSESs", types.YLeaf{"Dsx3FarEndCurrentCSESs", dsx3FarEndCurrentEntry.Dsx3FarEndCurrentCSESs})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndCurrentCCVs", types.YLeaf{"Dsx3FarEndCurrentCCVs", dsx3FarEndCurrentEntry.Dsx3FarEndCurrentCCVs})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndCurrentUASs", types.YLeaf{"Dsx3FarEndCurrentUASs", dsx3FarEndCurrentEntry.Dsx3FarEndCurrentUASs})
    dsx3FarEndCurrentEntry.EntityData.Leafs.Append("dsx3FarEndInvalidIntervals", types.YLeaf{"Dsx3FarEndInvalidIntervals", dsx3FarEndCurrentEntry.Dsx3FarEndInvalidIntervals})

    dsx3FarEndCurrentEntry.EntityData.YListKeys = []string {"Dsx3FarEndCurrentIndex"}

    return &(dsx3FarEndCurrentEntry.EntityData)
}

// DS3MIB_Dsx3FarEndIntervalTable
// The DS3 Far End Interval Table contains various
// statistics collected by each DS3 interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.
type DS3MIB_Dsx3FarEndIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Interval table. The type is slice of
    // DS3MIB_Dsx3FarEndIntervalTable_Dsx3FarEndIntervalEntry.
    Dsx3FarEndIntervalEntry []*DS3MIB_Dsx3FarEndIntervalTable_Dsx3FarEndIntervalEntry
}

func (dsx3FarEndIntervalTable *DS3MIB_Dsx3FarEndIntervalTable) GetEntityData() *types.CommonEntityData {
    dsx3FarEndIntervalTable.EntityData.YFilter = dsx3FarEndIntervalTable.YFilter
    dsx3FarEndIntervalTable.EntityData.YangName = "dsx3FarEndIntervalTable"
    dsx3FarEndIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndIntervalTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3FarEndIntervalTable.EntityData.SegmentPath = "dsx3FarEndIntervalTable"
    dsx3FarEndIntervalTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3FarEndIntervalTable.EntityData.SegmentPath
    dsx3FarEndIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndIntervalTable.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndIntervalTable.EntityData.Children.Append("dsx3FarEndIntervalEntry", types.YChild{"Dsx3FarEndIntervalEntry", nil})
    for i := range dsx3FarEndIntervalTable.Dsx3FarEndIntervalEntry {
        dsx3FarEndIntervalTable.EntityData.Children.Append(types.GetSegmentPath(dsx3FarEndIntervalTable.Dsx3FarEndIntervalEntry[i]), types.YChild{"Dsx3FarEndIntervalEntry", dsx3FarEndIntervalTable.Dsx3FarEndIntervalEntry[i]})
    }
    dsx3FarEndIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3FarEndIntervalTable.EntityData.YListKeys = []string {}

    return &(dsx3FarEndIntervalTable.EntityData)
}

// DS3MIB_Dsx3FarEndIntervalTable_Dsx3FarEndIntervalEntry
// An entry in the DS3 Far End Interval table.
type DS3MIB_Dsx3FarEndIntervalTable_Dsx3FarEndIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3FarEndIntervalIndex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx3FarEndIntervalNumber interface{}

    // The counter associated with the number of Far End C-bit Errored Seconds
    // encountered by a DS3 interface in one of the previous 96, individual 15
    // minute, intervals. In the case where the agent is a proxy and data is not
    // available, return noSuchInstance. The type is interface{} with range:
    // 0..4294967295.
    Dsx3FarEndIntervalCESs interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3FarEndIntervalCSESs interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count. The type is interface{} with
    // range: 0..4294967295.
    Dsx3FarEndIntervalCCVs interface{}

    // The counter associated with the number of Far End unavailable seconds. The
    // type is interface{} with range: 0..4294967295.
    Dsx3FarEndIntervalUASs interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx3FarEndIntervalValidData interface{}
}

func (dsx3FarEndIntervalEntry *DS3MIB_Dsx3FarEndIntervalTable_Dsx3FarEndIntervalEntry) GetEntityData() *types.CommonEntityData {
    dsx3FarEndIntervalEntry.EntityData.YFilter = dsx3FarEndIntervalEntry.YFilter
    dsx3FarEndIntervalEntry.EntityData.YangName = "dsx3FarEndIntervalEntry"
    dsx3FarEndIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndIntervalEntry.EntityData.ParentYangName = "dsx3FarEndIntervalTable"
    dsx3FarEndIntervalEntry.EntityData.SegmentPath = "dsx3FarEndIntervalEntry" + types.AddKeyToken(dsx3FarEndIntervalEntry.Dsx3FarEndIntervalIndex, "dsx3FarEndIntervalIndex") + types.AddKeyToken(dsx3FarEndIntervalEntry.Dsx3FarEndIntervalNumber, "dsx3FarEndIntervalNumber")
    dsx3FarEndIntervalEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3FarEndIntervalTable/" + dsx3FarEndIntervalEntry.EntityData.SegmentPath
    dsx3FarEndIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndIntervalEntry.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalIndex", types.YLeaf{"Dsx3FarEndIntervalIndex", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalIndex})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalNumber", types.YLeaf{"Dsx3FarEndIntervalNumber", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalNumber})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalCESs", types.YLeaf{"Dsx3FarEndIntervalCESs", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalCESs})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalCSESs", types.YLeaf{"Dsx3FarEndIntervalCSESs", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalCSESs})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalCCVs", types.YLeaf{"Dsx3FarEndIntervalCCVs", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalCCVs})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalUASs", types.YLeaf{"Dsx3FarEndIntervalUASs", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalUASs})
    dsx3FarEndIntervalEntry.EntityData.Leafs.Append("dsx3FarEndIntervalValidData", types.YLeaf{"Dsx3FarEndIntervalValidData", dsx3FarEndIntervalEntry.Dsx3FarEndIntervalValidData})

    dsx3FarEndIntervalEntry.EntityData.YListKeys = []string {"Dsx3FarEndIntervalIndex", "Dsx3FarEndIntervalNumber"}

    return &(dsx3FarEndIntervalEntry.EntityData)
}

// DS3MIB_Dsx3FarEndTotalTable
// The DS3 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS3MIB_Dsx3FarEndTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Total table. The type is slice of
    // DS3MIB_Dsx3FarEndTotalTable_Dsx3FarEndTotalEntry.
    Dsx3FarEndTotalEntry []*DS3MIB_Dsx3FarEndTotalTable_Dsx3FarEndTotalEntry
}

func (dsx3FarEndTotalTable *DS3MIB_Dsx3FarEndTotalTable) GetEntityData() *types.CommonEntityData {
    dsx3FarEndTotalTable.EntityData.YFilter = dsx3FarEndTotalTable.YFilter
    dsx3FarEndTotalTable.EntityData.YangName = "dsx3FarEndTotalTable"
    dsx3FarEndTotalTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndTotalTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3FarEndTotalTable.EntityData.SegmentPath = "dsx3FarEndTotalTable"
    dsx3FarEndTotalTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3FarEndTotalTable.EntityData.SegmentPath
    dsx3FarEndTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndTotalTable.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndTotalTable.EntityData.Children.Append("dsx3FarEndTotalEntry", types.YChild{"Dsx3FarEndTotalEntry", nil})
    for i := range dsx3FarEndTotalTable.Dsx3FarEndTotalEntry {
        dsx3FarEndTotalTable.EntityData.Children.Append(types.GetSegmentPath(dsx3FarEndTotalTable.Dsx3FarEndTotalEntry[i]), types.YChild{"Dsx3FarEndTotalEntry", dsx3FarEndTotalTable.Dsx3FarEndTotalEntry[i]})
    }
    dsx3FarEndTotalTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3FarEndTotalTable.EntityData.YListKeys = []string {}

    return &(dsx3FarEndTotalTable.EntityData)
}

// DS3MIB_Dsx3FarEndTotalTable_Dsx3FarEndTotalEntry
// An entry in the DS3 Far End Total table.
type DS3MIB_Dsx3FarEndTotalTable_Dsx3FarEndTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3FarEndTotalIndex interface{}

    // The counter associated with the number of Far End C-bit Errored Seconds
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3FarEndTotalCESs interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds encountered by a DS3 interface in the previous 24 hour interval.
    // Invalid 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3FarEndTotalCSESs interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count encountered by a DS3 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx3FarEndTotalCCVs interface{}

    // The counter associated with the number of Far End unavailable seconds
    // encountered by a DS3 interface in the previous 24 hour interval.  Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3FarEndTotalUASs interface{}
}

func (dsx3FarEndTotalEntry *DS3MIB_Dsx3FarEndTotalTable_Dsx3FarEndTotalEntry) GetEntityData() *types.CommonEntityData {
    dsx3FarEndTotalEntry.EntityData.YFilter = dsx3FarEndTotalEntry.YFilter
    dsx3FarEndTotalEntry.EntityData.YangName = "dsx3FarEndTotalEntry"
    dsx3FarEndTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3FarEndTotalEntry.EntityData.ParentYangName = "dsx3FarEndTotalTable"
    dsx3FarEndTotalEntry.EntityData.SegmentPath = "dsx3FarEndTotalEntry" + types.AddKeyToken(dsx3FarEndTotalEntry.Dsx3FarEndTotalIndex, "dsx3FarEndTotalIndex")
    dsx3FarEndTotalEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3FarEndTotalTable/" + dsx3FarEndTotalEntry.EntityData.SegmentPath
    dsx3FarEndTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FarEndTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FarEndTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FarEndTotalEntry.EntityData.Children = types.NewOrderedMap()
    dsx3FarEndTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3FarEndTotalEntry.EntityData.Leafs.Append("dsx3FarEndTotalIndex", types.YLeaf{"Dsx3FarEndTotalIndex", dsx3FarEndTotalEntry.Dsx3FarEndTotalIndex})
    dsx3FarEndTotalEntry.EntityData.Leafs.Append("dsx3FarEndTotalCESs", types.YLeaf{"Dsx3FarEndTotalCESs", dsx3FarEndTotalEntry.Dsx3FarEndTotalCESs})
    dsx3FarEndTotalEntry.EntityData.Leafs.Append("dsx3FarEndTotalCSESs", types.YLeaf{"Dsx3FarEndTotalCSESs", dsx3FarEndTotalEntry.Dsx3FarEndTotalCSESs})
    dsx3FarEndTotalEntry.EntityData.Leafs.Append("dsx3FarEndTotalCCVs", types.YLeaf{"Dsx3FarEndTotalCCVs", dsx3FarEndTotalEntry.Dsx3FarEndTotalCCVs})
    dsx3FarEndTotalEntry.EntityData.Leafs.Append("dsx3FarEndTotalUASs", types.YLeaf{"Dsx3FarEndTotalUASs", dsx3FarEndTotalEntry.Dsx3FarEndTotalUASs})

    dsx3FarEndTotalEntry.EntityData.YListKeys = []string {"Dsx3FarEndTotalIndex"}

    return &(dsx3FarEndTotalEntry.EntityData)
}

// DS3MIB_Dsx3FracTable
// This table is deprecated in favour of using
// ifStackTable.
// 
// Implementation of this table was optional.  It was
// designed for those systems dividing a DS3/E3 into
// channels containing different data streams that
// are of local interest.
// 
// The DS3/E3 fractional table identifies which
// DS3/E3 channels associated with a CSU are being
// used to support a logical interface, i.e., an
// entry in the interfaces table from the Internet-
// standard MIB.
// 
// For example, consider a DS3 device with 4 high
// speed links carrying router traffic, a feed for
// voice, a feed for video, and a synchronous channel
// for a non-routed protocol.  We might describe the
// allocation of channels, in the dsx3FracTable, as
// follows:
// dsx3FracIfIndex.2. 1 = 3  dsx3FracIfIndex.2.15 = 4
// dsx3FracIfIndex.2. 2 = 3  dsx3FracIfIndex.2.16 = 6
// dsx3FracIfIndex.2. 3 = 3  dsx3FracIfIndex.2.17 = 6
// dsx3FracIfIndex.2. 4 = 3  dsx3FracIfIndex.2.18 = 6
// dsx3FracIfIndex.2. 5 = 3  dsx3FracIfIndex.2.19 = 6
// dsx3FracIfIndex.2. 6 = 3  dsx3FracIfIndex.2.20 = 6
// dsx3FracIfIndex.2. 7 = 4  dsx3FracIfIndex.2.21 = 6
// dsx3FracIfIndex.2. 8 = 4  dsx3FracIfIndex.2.22 = 6
// dsx3FracIfIndex.2. 9 = 4  dsx3FracIfIndex.2.23 = 6
// dsx3FracIfIndex.2.10 = 4  dsx3FracIfIndex.2.24 = 6
// dsx3FracIfIndex.2.11 = 4  dsx3FracIfIndex.2.25 = 6
// dsx3FracIfIndex.2.12 = 5  dsx3FracIfIndex.2.26 = 6
// dsx3FracIfIndex.2.13 = 5  dsx3FracIfIndex.2.27 = 6
// dsx3FracIfIndex.2.14 = 5  dsx3FracIfIndex.2.28 = 6
// For dsx3M23, dsx3 SYNTRAN, dsx3CbitParity, and
// dsx3ClearChannel  there are 28 legal channels,
// numbered 1 throug h 28.
// 
// For e3Framed there are 16 legal channels, numbered
// 1 through 16.  The channels (1..16) correspond
// directly to the equivalently numbered time-slots.
type DS3MIB_Dsx3FracTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Fractional table. The type is slice of
    // DS3MIB_Dsx3FracTable_Dsx3FracEntry.
    Dsx3FracEntry []*DS3MIB_Dsx3FracTable_Dsx3FracEntry
}

func (dsx3FracTable *DS3MIB_Dsx3FracTable) GetEntityData() *types.CommonEntityData {
    dsx3FracTable.EntityData.YFilter = dsx3FracTable.YFilter
    dsx3FracTable.EntityData.YangName = "dsx3FracTable"
    dsx3FracTable.EntityData.BundleName = "cisco_ios_xe"
    dsx3FracTable.EntityData.ParentYangName = "DS3-MIB"
    dsx3FracTable.EntityData.SegmentPath = "dsx3FracTable"
    dsx3FracTable.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/" + dsx3FracTable.EntityData.SegmentPath
    dsx3FracTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FracTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FracTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FracTable.EntityData.Children = types.NewOrderedMap()
    dsx3FracTable.EntityData.Children.Append("dsx3FracEntry", types.YChild{"Dsx3FracEntry", nil})
    for i := range dsx3FracTable.Dsx3FracEntry {
        dsx3FracTable.EntityData.Children.Append(types.GetSegmentPath(dsx3FracTable.Dsx3FracEntry[i]), types.YChild{"Dsx3FracEntry", dsx3FracTable.Dsx3FracEntry[i]})
    }
    dsx3FracTable.EntityData.Leafs = types.NewOrderedMap()

    dsx3FracTable.EntityData.YListKeys = []string {}

    return &(dsx3FracTable.EntityData)
}

// DS3MIB_Dsx3FracTable_Dsx3FracEntry
// An entry in the DS3 Fractional table.
type DS3MIB_Dsx3FracTable_Dsx3FracEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index value which uniquely identifies  the DS3
    // interface  to which this entry is applicable The interface identified by a 
    // particular value of  this  index is the same interface as identified by the
    // same value  an  dsx3LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx3FracIndex interface{}

    // This attribute is a key. The channel number for this entry. The type is
    // interface{} with range: 1..31.
    Dsx3FracNumber interface{}

    // An index value that uniquely identifies an interface.  The interface
    // identified by a particular value of this index is the same interface as 
    // identified by the same value an ifIndex object instance. If no interface is
    // currently using a channel, the value should be zero.  If a single interface
    // occupies more  than one  time slot,  that ifIndex value will be found in
    // multiple time slots. The type is interface{} with range: 1..2147483647.
    Dsx3FracIfIndex interface{}
}

func (dsx3FracEntry *DS3MIB_Dsx3FracTable_Dsx3FracEntry) GetEntityData() *types.CommonEntityData {
    dsx3FracEntry.EntityData.YFilter = dsx3FracEntry.YFilter
    dsx3FracEntry.EntityData.YangName = "dsx3FracEntry"
    dsx3FracEntry.EntityData.BundleName = "cisco_ios_xe"
    dsx3FracEntry.EntityData.ParentYangName = "dsx3FracTable"
    dsx3FracEntry.EntityData.SegmentPath = "dsx3FracEntry" + types.AddKeyToken(dsx3FracEntry.Dsx3FracIndex, "dsx3FracIndex") + types.AddKeyToken(dsx3FracEntry.Dsx3FracNumber, "dsx3FracNumber")
    dsx3FracEntry.EntityData.AbsolutePath = "DS3-MIB:DS3-MIB/dsx3FracTable/" + dsx3FracEntry.EntityData.SegmentPath
    dsx3FracEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3FracEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3FracEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3FracEntry.EntityData.Children = types.NewOrderedMap()
    dsx3FracEntry.EntityData.Leafs = types.NewOrderedMap()
    dsx3FracEntry.EntityData.Leafs.Append("dsx3FracIndex", types.YLeaf{"Dsx3FracIndex", dsx3FracEntry.Dsx3FracIndex})
    dsx3FracEntry.EntityData.Leafs.Append("dsx3FracNumber", types.YLeaf{"Dsx3FracNumber", dsx3FracEntry.Dsx3FracNumber})
    dsx3FracEntry.EntityData.Leafs.Append("dsx3FracIfIndex", types.YLeaf{"Dsx3FracIfIndex", dsx3FracEntry.Dsx3FracIfIndex})

    dsx3FracEntry.EntityData.YListKeys = []string {"Dsx3FracIndex", "Dsx3FracNumber"}

    return &(dsx3FracEntry.EntityData)
}

