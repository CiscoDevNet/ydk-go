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
    Dsx3Configtable DS3MIB_Dsx3Configtable

    // The DS3/E3 current table contains various statistics being collected for
    // the current 15 minute interval.
    Dsx3Currenttable DS3MIB_Dsx3Currenttable

    // The DS3/E3 Interval Table contains various statistics collected by each
    // DS3/E3 Interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals.  Each row in this
    // table represents one such interval (identified by dsx3IntervalNumber) and
    // for one specific interface (identifed by dsx3IntervalIndex).
    Dsx3Intervaltable DS3MIB_Dsx3Intervaltable

    // The DS3/E3 Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx3Totaltable DS3MIB_Dsx3Totaltable

    // The DS3 Far End Configuration Table contains configuration information
    // reported in the C-bits from the remote end.
    Dsx3Farendconfigtable DS3MIB_Dsx3Farendconfigtable

    // The DS3 Far End Current table contains various statistics being collected
    // for the current 15 minute interval.  The statistics are collected from the
    // far end block error code within the C- bits.
    Dsx3Farendcurrenttable DS3MIB_Dsx3Farendcurrenttable

    // The DS3 Far End Interval Table contains various statistics collected by
    // each DS3 interface over the previous 24 hours of operation.  The past 24
    // hours are broken into 96 completed 15 minute intervals.
    Dsx3Farendintervaltable DS3MIB_Dsx3Farendintervaltable

    // The DS3 Far End Total Table contains the cumulative sum of the various
    // statistics for the 24 hour period preceding the current interval.
    Dsx3Farendtotaltable DS3MIB_Dsx3Farendtotaltable

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
    Dsx3Fractable DS3MIB_Dsx3Fractable
}

func (dS3MIB *DS3MIB) GetEntityData() *types.CommonEntityData {
    dS3MIB.EntityData.YFilter = dS3MIB.YFilter
    dS3MIB.EntityData.YangName = "DS3-MIB"
    dS3MIB.EntityData.BundleName = "cisco_ios_xe"
    dS3MIB.EntityData.ParentYangName = "DS3-MIB"
    dS3MIB.EntityData.SegmentPath = "DS3-MIB:DS3-MIB"
    dS3MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dS3MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dS3MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dS3MIB.EntityData.Children = make(map[string]types.YChild)
    dS3MIB.EntityData.Children["dsx3ConfigTable"] = types.YChild{"Dsx3Configtable", &dS3MIB.Dsx3Configtable}
    dS3MIB.EntityData.Children["dsx3CurrentTable"] = types.YChild{"Dsx3Currenttable", &dS3MIB.Dsx3Currenttable}
    dS3MIB.EntityData.Children["dsx3IntervalTable"] = types.YChild{"Dsx3Intervaltable", &dS3MIB.Dsx3Intervaltable}
    dS3MIB.EntityData.Children["dsx3TotalTable"] = types.YChild{"Dsx3Totaltable", &dS3MIB.Dsx3Totaltable}
    dS3MIB.EntityData.Children["dsx3FarEndConfigTable"] = types.YChild{"Dsx3Farendconfigtable", &dS3MIB.Dsx3Farendconfigtable}
    dS3MIB.EntityData.Children["dsx3FarEndCurrentTable"] = types.YChild{"Dsx3Farendcurrenttable", &dS3MIB.Dsx3Farendcurrenttable}
    dS3MIB.EntityData.Children["dsx3FarEndIntervalTable"] = types.YChild{"Dsx3Farendintervaltable", &dS3MIB.Dsx3Farendintervaltable}
    dS3MIB.EntityData.Children["dsx3FarEndTotalTable"] = types.YChild{"Dsx3Farendtotaltable", &dS3MIB.Dsx3Farendtotaltable}
    dS3MIB.EntityData.Children["dsx3FracTable"] = types.YChild{"Dsx3Fractable", &dS3MIB.Dsx3Fractable}
    dS3MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dS3MIB.EntityData)
}

// DS3MIB_Dsx3Configtable
// The DS3/E3 Configuration table.
type DS3MIB_Dsx3Configtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Configuration table. The type is slice of
    // DS3MIB_Dsx3Configtable_Dsx3Configentry.
    Dsx3Configentry []DS3MIB_Dsx3Configtable_Dsx3Configentry
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetEntityData() *types.CommonEntityData {
    dsx3Configtable.EntityData.YFilter = dsx3Configtable.YFilter
    dsx3Configtable.EntityData.YangName = "dsx3ConfigTable"
    dsx3Configtable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Configtable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Configtable.EntityData.SegmentPath = "dsx3ConfigTable"
    dsx3Configtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Configtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Configtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Configtable.EntityData.Children = make(map[string]types.YChild)
    dsx3Configtable.EntityData.Children["dsx3ConfigEntry"] = types.YChild{"Dsx3Configentry", nil}
    for i := range dsx3Configtable.Dsx3Configentry {
        dsx3Configtable.EntityData.Children[types.GetSegmentPath(&dsx3Configtable.Dsx3Configentry[i])] = types.YChild{"Dsx3Configentry", &dsx3Configtable.Dsx3Configentry[i]}
    }
    dsx3Configtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Configtable.EntityData)
}

// DS3MIB_Dsx3Configtable_Dsx3Configentry
// An entry in the DS3/E3 Configuration table.
type DS3MIB_Dsx3Configtable_Dsx3Configentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    Dsx3Lineindex interface{}

    // This value for this object is equal to the value of ifIndex from the
    // Interfaces table of MIB II (RFC 1213). The type is interface{} with range:
    // 1..2147483647.
    Dsx3Ifindex interface{}

    // The number of seconds that have elapsed since the beginning of the near end
    // current error- measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx3Timeelapsed interface{}

    // The number of previous near end intervals for which data was collected. 
    // The value will be 96 unless the interface was brought online within the
    // last 24 hours, in which case the value will be the number of complete 15
    // minute near end intervals since the interface has been online.  In the case
    // where the agent is a proxy, it is possible that some intervals are
    // unavailable.  In this case, this interval is the maximum interval number
    // for which data is available. The type is interface{} with range: 0..96.
    Dsx3Validintervals interface{}

    // This variable indicates the variety of DS3 C-bit or E3 application
    // implementing this interface. The type of interface affects the
    // interpretation of the usage and error statistics.  The rate of DS3 is
    // 44.736 Mbps and E3 is 34.368 Mbps.  The dsx3ClearChannel value means that
    // the C-bits are not used except for sending/receiving AIS. The values, in
    // sequence, describe:  TITLE:            SPECIFICATION: dsx3M23           
    // ANSI T1.107-1988 [9] dsx3SYNTRAN        ANSI T1.107-1988 [9] dsx3CbitParity
    // ANSI T1.107a-1990 [9a] dsx3ClearChannel   ANSI T1.102-1987 [8] e3Framed    
    // CCITT G.751 [12] e3Plcp             ETSI T/NA(91)18 [13]. The type is
    // Dsx3Linetype.
    Dsx3Linetype interface{}

    // This variable describes the variety of Zero Code Suppression used on this
    // interface, which in turn affects a number of its characteristics.  dsx3B3ZS
    // and e3HDB3 refer to the use of specified patterns of normal bits and
    // bipolar violations which are used to replace sequences of zero bits of a
    // specified length. The type is Dsx3Linecoding.
    Dsx3Linecoding interface{}

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
    // sending a test pattern. The type is Dsx3Sendcode.
    Dsx3Sendcode interface{}

    // This variable contains the transmission vendor's circuit identifier, for
    // the purpose of facilitating troubleshooting. The type is string with
    // length: 0..255.
    Dsx3Circuitidentifier interface{}

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
    // will be   active simultaneously. The type is Dsx3Loopbackconfig.
    Dsx3Loopbackconfig interface{}

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
    Dsx3Linestatus interface{}

    // The source of Transmit Clock.  loopTiming indicates that the recovered
    // receive clock is used as the transmit clock.  localTiming indicates that a
    // local clock source is used or that an external clock is attached to the box
    // containing the interface.  throughTiming indicates that transmit clock is
    // derived from the recovered receive clock of another DS3 interface. The type
    // is Dsx3Transmitclocksource.
    Dsx3Transmitclocksource interface{}

    // The number of intervals in the range from 0 to dsx3ValidIntervals for which
    // no data is available.  This object will typically be zero except in cases
    // where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx3Invalidintervals interface{}

    // The length of the ds3 line in meters.  This object provides information for
    // line build out circuitry if it exists and can use this object to adjust the
    // line build out. The type is interface{} with range: 0..64000. Units are
    // meters.
    Dsx3Linelength interface{}

    // The value of MIB II's sysUpTime object at the time this DS3/E3 entered its
    // current line status state.  If the current state was entered prior to the
    // last re-initialization of the proxy-agent, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    Dsx3Linestatuslastchange interface{}

    // Indicates whether dsx3LineStatusChange traps should be generated for this
    // interface. The type is Dsx3Linestatuschangetrapenable.
    Dsx3Linestatuschangetrapenable interface{}

    // This variable represents the current state of the loopback on the DS3
    // interface.  It contains information about loopbacks established by a
    // manager and remotely from the far end.  The dsx3LoopbackStatus is a bit map
    // represented as a sum, therefore is can represent multiple loopbacks
    // simultaneously.  The various bit positions are:  1  dsx3NoLoopback  2 
    // dsx3NearEndPayloadLoopback  4  dsx3NearEndLineLoopback  8 
    // dsx3NearEndOtherLoopback 16  dsx3NearEndInwardLoopback 32 
    // dsx3FarEndPayloadLoopback 64  dsx3FarEndLineLoopback. The type is
    // interface{} with range: 1..127.
    Dsx3Loopbackstatus interface{}

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
    Dsx3Ds1Forremoteloop interface{}
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetEntityData() *types.CommonEntityData {
    dsx3Configentry.EntityData.YFilter = dsx3Configentry.YFilter
    dsx3Configentry.EntityData.YangName = "dsx3ConfigEntry"
    dsx3Configentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Configentry.EntityData.ParentYangName = "dsx3ConfigTable"
    dsx3Configentry.EntityData.SegmentPath = "dsx3ConfigEntry" + "[dsx3LineIndex='" + fmt.Sprintf("%v", dsx3Configentry.Dsx3Lineindex) + "']"
    dsx3Configentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Configentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Configentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Configentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Configentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Configentry.EntityData.Leafs["dsx3LineIndex"] = types.YLeaf{"Dsx3Lineindex", dsx3Configentry.Dsx3Lineindex}
    dsx3Configentry.EntityData.Leafs["dsx3IfIndex"] = types.YLeaf{"Dsx3Ifindex", dsx3Configentry.Dsx3Ifindex}
    dsx3Configentry.EntityData.Leafs["dsx3TimeElapsed"] = types.YLeaf{"Dsx3Timeelapsed", dsx3Configentry.Dsx3Timeelapsed}
    dsx3Configentry.EntityData.Leafs["dsx3ValidIntervals"] = types.YLeaf{"Dsx3Validintervals", dsx3Configentry.Dsx3Validintervals}
    dsx3Configentry.EntityData.Leafs["dsx3LineType"] = types.YLeaf{"Dsx3Linetype", dsx3Configentry.Dsx3Linetype}
    dsx3Configentry.EntityData.Leafs["dsx3LineCoding"] = types.YLeaf{"Dsx3Linecoding", dsx3Configentry.Dsx3Linecoding}
    dsx3Configentry.EntityData.Leafs["dsx3SendCode"] = types.YLeaf{"Dsx3Sendcode", dsx3Configentry.Dsx3Sendcode}
    dsx3Configentry.EntityData.Leafs["dsx3CircuitIdentifier"] = types.YLeaf{"Dsx3Circuitidentifier", dsx3Configentry.Dsx3Circuitidentifier}
    dsx3Configentry.EntityData.Leafs["dsx3LoopbackConfig"] = types.YLeaf{"Dsx3Loopbackconfig", dsx3Configentry.Dsx3Loopbackconfig}
    dsx3Configentry.EntityData.Leafs["dsx3LineStatus"] = types.YLeaf{"Dsx3Linestatus", dsx3Configentry.Dsx3Linestatus}
    dsx3Configentry.EntityData.Leafs["dsx3TransmitClockSource"] = types.YLeaf{"Dsx3Transmitclocksource", dsx3Configentry.Dsx3Transmitclocksource}
    dsx3Configentry.EntityData.Leafs["dsx3InvalidIntervals"] = types.YLeaf{"Dsx3Invalidintervals", dsx3Configentry.Dsx3Invalidintervals}
    dsx3Configentry.EntityData.Leafs["dsx3LineLength"] = types.YLeaf{"Dsx3Linelength", dsx3Configentry.Dsx3Linelength}
    dsx3Configentry.EntityData.Leafs["dsx3LineStatusLastChange"] = types.YLeaf{"Dsx3Linestatuslastchange", dsx3Configentry.Dsx3Linestatuslastchange}
    dsx3Configentry.EntityData.Leafs["dsx3LineStatusChangeTrapEnable"] = types.YLeaf{"Dsx3Linestatuschangetrapenable", dsx3Configentry.Dsx3Linestatuschangetrapenable}
    dsx3Configentry.EntityData.Leafs["dsx3LoopbackStatus"] = types.YLeaf{"Dsx3Loopbackstatus", dsx3Configentry.Dsx3Loopbackstatus}
    dsx3Configentry.EntityData.Leafs["dsx3Channelization"] = types.YLeaf{"Dsx3Channelization", dsx3Configentry.Dsx3Channelization}
    dsx3Configentry.EntityData.Leafs["dsx3Ds1ForRemoteLoop"] = types.YLeaf{"Dsx3Ds1Forremoteloop", dsx3Configentry.Dsx3Ds1Forremoteloop}
    return &(dsx3Configentry.EntityData)
}

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization represents entries in the ifTable.  
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization_disabled DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization = "disabled"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization_enabledDs1 DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization = "enabledDs1"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization_enabledDs2 DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Channelization = "enabledDs2"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding represents of a specified length.
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding_dsx3Other DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding = "dsx3Other"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding_dsx3B3ZS DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding = "dsx3B3ZS"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding_e3HDB3 DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linecoding = "e3HDB3"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable represents should be generated for this interface.
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable_enabled DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable = "enabled"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable_disabled DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linestatuschangetrapenable = "disabled"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype represents e3Plcp             ETSI T/NA(91)18 [13].
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_dsx3other DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "dsx3other"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_dsx3M23 DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "dsx3M23"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_dsx3SYNTRAN DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "dsx3SYNTRAN"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_dsx3CbitParity DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "dsx3CbitParity"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_dsx3ClearChannel DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "dsx3ClearChannel"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_e3other DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "e3other"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_e3Framed DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "e3Framed"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype_e3Plcp DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Linetype = "e3Plcp"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig represents   active simultaneously.
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3NoLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3NoLoop"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3PayloadLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3PayloadLoop"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3LineLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3LineLoop"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3OtherLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3OtherLoop"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3InwardLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3InwardLoop"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig_dsx3DualLoop DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Loopbackconfig = "dsx3DualLoop"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode represents        sending a test pattern.
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendNoCode DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendNoCode"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendLineCode DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendLineCode"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendPayloadCode DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendPayloadCode"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendResetCode DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendResetCode"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendDS1LoopCode DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendDS1LoopCode"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode_dsx3SendTestPattern DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Sendcode = "dsx3SendTestPattern"
)

// DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource represents interface.
type DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource string

const (
    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource_loopTiming DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource = "loopTiming"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource_localTiming DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource = "localTiming"

    DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource_throughTiming DS3MIB_Dsx3Configtable_Dsx3Configentry_Dsx3Transmitclocksource = "throughTiming"
)

// DS3MIB_Dsx3Currenttable
// The DS3/E3 current table contains various
// statistics being collected for the current 15
// minute interval.
type DS3MIB_Dsx3Currenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Current table. The type is slice of
    // DS3MIB_Dsx3Currenttable_Dsx3Currententry.
    Dsx3Currententry []DS3MIB_Dsx3Currenttable_Dsx3Currententry
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetEntityData() *types.CommonEntityData {
    dsx3Currenttable.EntityData.YFilter = dsx3Currenttable.YFilter
    dsx3Currenttable.EntityData.YangName = "dsx3CurrentTable"
    dsx3Currenttable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Currenttable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Currenttable.EntityData.SegmentPath = "dsx3CurrentTable"
    dsx3Currenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Currenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Currenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Currenttable.EntityData.Children = make(map[string]types.YChild)
    dsx3Currenttable.EntityData.Children["dsx3CurrentEntry"] = types.YChild{"Dsx3Currententry", nil}
    for i := range dsx3Currenttable.Dsx3Currententry {
        dsx3Currenttable.EntityData.Children[types.GetSegmentPath(&dsx3Currenttable.Dsx3Currententry[i])] = types.YChild{"Dsx3Currententry", &dsx3Currenttable.Dsx3Currententry[i]}
    }
    dsx3Currenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Currenttable.EntityData)
}

// DS3MIB_Dsx3Currenttable_Dsx3Currententry
// An entry in the DS3/E3 Current table.
type DS3MIB_Dsx3Currenttable_Dsx3Currententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3Currentindex interface{}

    // The counter associated with the number of P-bit Errored Seconds. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Currentpess interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Currentpsess interface{}

    // The counter associated with the number of Severely Errored Framing Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Currentsefss interface{}

    // The counter associated with the number of Unavailable Seconds. The type is
    // interface{} with range: 0..4294967295.
    Dsx3Currentuass interface{}

    // The counter associated with the number of Line Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Currentlcvs interface{}

    // The counter associated with the number of P-bit Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Currentpcvs interface{}

    // The number of Line Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Currentless interface{}

    // The number of C-bit Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Currentccvs interface{}

    // The number of C-bit Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Currentcess interface{}

    // The number of C-bit Severely Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx3Currentcsess interface{}
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetEntityData() *types.CommonEntityData {
    dsx3Currententry.EntityData.YFilter = dsx3Currententry.YFilter
    dsx3Currententry.EntityData.YangName = "dsx3CurrentEntry"
    dsx3Currententry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Currententry.EntityData.ParentYangName = "dsx3CurrentTable"
    dsx3Currententry.EntityData.SegmentPath = "dsx3CurrentEntry" + "[dsx3CurrentIndex='" + fmt.Sprintf("%v", dsx3Currententry.Dsx3Currentindex) + "']"
    dsx3Currententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Currententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Currententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Currententry.EntityData.Children = make(map[string]types.YChild)
    dsx3Currententry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Currententry.EntityData.Leafs["dsx3CurrentIndex"] = types.YLeaf{"Dsx3Currentindex", dsx3Currententry.Dsx3Currentindex}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentPESs"] = types.YLeaf{"Dsx3Currentpess", dsx3Currententry.Dsx3Currentpess}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentPSESs"] = types.YLeaf{"Dsx3Currentpsess", dsx3Currententry.Dsx3Currentpsess}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentSEFSs"] = types.YLeaf{"Dsx3Currentsefss", dsx3Currententry.Dsx3Currentsefss}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentUASs"] = types.YLeaf{"Dsx3Currentuass", dsx3Currententry.Dsx3Currentuass}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentLCVs"] = types.YLeaf{"Dsx3Currentlcvs", dsx3Currententry.Dsx3Currentlcvs}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentPCVs"] = types.YLeaf{"Dsx3Currentpcvs", dsx3Currententry.Dsx3Currentpcvs}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentLESs"] = types.YLeaf{"Dsx3Currentless", dsx3Currententry.Dsx3Currentless}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentCCVs"] = types.YLeaf{"Dsx3Currentccvs", dsx3Currententry.Dsx3Currentccvs}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentCESs"] = types.YLeaf{"Dsx3Currentcess", dsx3Currententry.Dsx3Currentcess}
    dsx3Currententry.EntityData.Leafs["dsx3CurrentCSESs"] = types.YLeaf{"Dsx3Currentcsess", dsx3Currententry.Dsx3Currentcsess}
    return &(dsx3Currententry.EntityData)
}

// DS3MIB_Dsx3Intervaltable
// The DS3/E3 Interval Table contains various
// statistics collected by each DS3/E3 Interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.  Each row in this table represents one
// such interval (identified by dsx3IntervalNumber)
// and for one specific interface (identifed by
// dsx3IntervalIndex).
type DS3MIB_Dsx3Intervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Interval table. The type is slice of
    // DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry.
    Dsx3Intervalentry []DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetEntityData() *types.CommonEntityData {
    dsx3Intervaltable.EntityData.YFilter = dsx3Intervaltable.YFilter
    dsx3Intervaltable.EntityData.YangName = "dsx3IntervalTable"
    dsx3Intervaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Intervaltable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Intervaltable.EntityData.SegmentPath = "dsx3IntervalTable"
    dsx3Intervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Intervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Intervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Intervaltable.EntityData.Children = make(map[string]types.YChild)
    dsx3Intervaltable.EntityData.Children["dsx3IntervalEntry"] = types.YChild{"Dsx3Intervalentry", nil}
    for i := range dsx3Intervaltable.Dsx3Intervalentry {
        dsx3Intervaltable.EntityData.Children[types.GetSegmentPath(&dsx3Intervaltable.Dsx3Intervalentry[i])] = types.YChild{"Dsx3Intervalentry", &dsx3Intervaltable.Dsx3Intervalentry[i]}
    }
    dsx3Intervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Intervaltable.EntityData)
}

// DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry
// An entry in the DS3/E3 Interval table.
type DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3Intervalindex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx3Intervalnumber interface{}

    // The counter associated with the number of P-bit Errored Seconds. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Intervalpess interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Intervalpsess interface{}

    // The counter associated with the number of Severely Errored Framing Seconds.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Intervalsefss interface{}

    // The counter associated with the number of Unavailable Seconds.  This object
    // may decrease if the occurance of unavailable seconds occurs across an
    // inteval boundary. The type is interface{} with range: 0..4294967295.
    Dsx3Intervaluass interface{}

    // The counter associated with the number of Line Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Intervallcvs interface{}

    // The counter associated with the number of P-bit Coding Violations. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Intervalpcvs interface{}

    // The number of Line Errored  Seconds  (BPVs  or illegal  zero  sequences).
    // The type is interface{} with range: 0..4294967295.
    Dsx3Intervalless interface{}

    // The number of C-bit Coding Violations. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Intervalccvs interface{}

    // The number of C-bit Errored Seconds. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Intervalcess interface{}

    // The number of C-bit Severely Errored Seconds. The type is interface{} with
    // range: 0..4294967295.
    Dsx3Intervalcsess interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx3Intervalvaliddata interface{}
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetEntityData() *types.CommonEntityData {
    dsx3Intervalentry.EntityData.YFilter = dsx3Intervalentry.YFilter
    dsx3Intervalentry.EntityData.YangName = "dsx3IntervalEntry"
    dsx3Intervalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Intervalentry.EntityData.ParentYangName = "dsx3IntervalTable"
    dsx3Intervalentry.EntityData.SegmentPath = "dsx3IntervalEntry" + "[dsx3IntervalIndex='" + fmt.Sprintf("%v", dsx3Intervalentry.Dsx3Intervalindex) + "']" + "[dsx3IntervalNumber='" + fmt.Sprintf("%v", dsx3Intervalentry.Dsx3Intervalnumber) + "']"
    dsx3Intervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Intervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Intervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Intervalentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Intervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalIndex"] = types.YLeaf{"Dsx3Intervalindex", dsx3Intervalentry.Dsx3Intervalindex}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalNumber"] = types.YLeaf{"Dsx3Intervalnumber", dsx3Intervalentry.Dsx3Intervalnumber}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalPESs"] = types.YLeaf{"Dsx3Intervalpess", dsx3Intervalentry.Dsx3Intervalpess}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalPSESs"] = types.YLeaf{"Dsx3Intervalpsess", dsx3Intervalentry.Dsx3Intervalpsess}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalSEFSs"] = types.YLeaf{"Dsx3Intervalsefss", dsx3Intervalentry.Dsx3Intervalsefss}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalUASs"] = types.YLeaf{"Dsx3Intervaluass", dsx3Intervalentry.Dsx3Intervaluass}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalLCVs"] = types.YLeaf{"Dsx3Intervallcvs", dsx3Intervalentry.Dsx3Intervallcvs}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalPCVs"] = types.YLeaf{"Dsx3Intervalpcvs", dsx3Intervalentry.Dsx3Intervalpcvs}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalLESs"] = types.YLeaf{"Dsx3Intervalless", dsx3Intervalentry.Dsx3Intervalless}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalCCVs"] = types.YLeaf{"Dsx3Intervalccvs", dsx3Intervalentry.Dsx3Intervalccvs}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalCESs"] = types.YLeaf{"Dsx3Intervalcess", dsx3Intervalentry.Dsx3Intervalcess}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalCSESs"] = types.YLeaf{"Dsx3Intervalcsess", dsx3Intervalentry.Dsx3Intervalcsess}
    dsx3Intervalentry.EntityData.Leafs["dsx3IntervalValidData"] = types.YLeaf{"Dsx3Intervalvaliddata", dsx3Intervalentry.Dsx3Intervalvaliddata}
    return &(dsx3Intervalentry.EntityData)
}

// DS3MIB_Dsx3Totaltable
// The DS3/E3 Total Table contains the cumulative
// sum of the various statistics for the 24 hour
// period preceding the current interval.
type DS3MIB_Dsx3Totaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Total table. The type is slice of
    // DS3MIB_Dsx3Totaltable_Dsx3Totalentry.
    Dsx3Totalentry []DS3MIB_Dsx3Totaltable_Dsx3Totalentry
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetEntityData() *types.CommonEntityData {
    dsx3Totaltable.EntityData.YFilter = dsx3Totaltable.YFilter
    dsx3Totaltable.EntityData.YangName = "dsx3TotalTable"
    dsx3Totaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Totaltable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Totaltable.EntityData.SegmentPath = "dsx3TotalTable"
    dsx3Totaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Totaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Totaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Totaltable.EntityData.Children = make(map[string]types.YChild)
    dsx3Totaltable.EntityData.Children["dsx3TotalEntry"] = types.YChild{"Dsx3Totalentry", nil}
    for i := range dsx3Totaltable.Dsx3Totalentry {
        dsx3Totaltable.EntityData.Children[types.GetSegmentPath(&dsx3Totaltable.Dsx3Totalentry[i])] = types.YChild{"Dsx3Totalentry", &dsx3Totaltable.Dsx3Totalentry[i]}
    }
    dsx3Totaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Totaltable.EntityData)
}

// DS3MIB_Dsx3Totaltable_Dsx3Totalentry
// An entry in the DS3/E3 Total table.
type DS3MIB_Dsx3Totaltable_Dsx3Totalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the
    // DS3/E3 interface to which this entry is applicable.  The interface
    // identified by a particular value of this index is the same interface as
    // identified by the same value an dsx3LineIndex object instance. The type is
    // interface{} with range: 1..2147483647.
    Dsx3Totalindex interface{}

    // The counter associated with the number of P-bit Errored Seconds,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totalpess interface{}

    // The counter associated with the number of P-bit Severely Errored Seconds,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totalpsess interface{}

    // The counter associated with the number of Severely Errored Framing Seconds,
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totalsefss interface{}

    // The counter associated with the number of Unavailable Seconds, encountered
    // by a DS3 interface in the previous 24 hour interval.  Invalid 15 minute
    // intervals count as 0. The type is interface{} with range: 0..4294967295.
    Dsx3Totaluass interface{}

    // The counter associated with the number of Line Coding Violations
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totallcvs interface{}

    // The counter associated with the number of P-bit Coding Violations,
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totalpcvs interface{}

    // The number of Line Errored  Seconds  (BPVs  or illegal  zero  sequences)
    // encountered by a DS3/E3 interface in the previous 24 hour interval. Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Totalless interface{}

    // The number of C-bit Coding Violations encountered by a DS3 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Totalccvs interface{}

    // The number of C-bit Errored Seconds encountered by a DS3 interface in the
    // previous 24 hour interval. Invalid 15 minute intervals count as 0. The type
    // is interface{} with range: 0..4294967295.
    Dsx3Totalcess interface{}

    // The number of C-bit Severely Errored Seconds encountered by a DS3 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Totalcsess interface{}
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetEntityData() *types.CommonEntityData {
    dsx3Totalentry.EntityData.YFilter = dsx3Totalentry.YFilter
    dsx3Totalentry.EntityData.YangName = "dsx3TotalEntry"
    dsx3Totalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Totalentry.EntityData.ParentYangName = "dsx3TotalTable"
    dsx3Totalentry.EntityData.SegmentPath = "dsx3TotalEntry" + "[dsx3TotalIndex='" + fmt.Sprintf("%v", dsx3Totalentry.Dsx3Totalindex) + "']"
    dsx3Totalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Totalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Totalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Totalentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Totalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Totalentry.EntityData.Leafs["dsx3TotalIndex"] = types.YLeaf{"Dsx3Totalindex", dsx3Totalentry.Dsx3Totalindex}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalPESs"] = types.YLeaf{"Dsx3Totalpess", dsx3Totalentry.Dsx3Totalpess}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalPSESs"] = types.YLeaf{"Dsx3Totalpsess", dsx3Totalentry.Dsx3Totalpsess}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalSEFSs"] = types.YLeaf{"Dsx3Totalsefss", dsx3Totalentry.Dsx3Totalsefss}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalUASs"] = types.YLeaf{"Dsx3Totaluass", dsx3Totalentry.Dsx3Totaluass}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalLCVs"] = types.YLeaf{"Dsx3Totallcvs", dsx3Totalentry.Dsx3Totallcvs}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalPCVs"] = types.YLeaf{"Dsx3Totalpcvs", dsx3Totalentry.Dsx3Totalpcvs}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalLESs"] = types.YLeaf{"Dsx3Totalless", dsx3Totalentry.Dsx3Totalless}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalCCVs"] = types.YLeaf{"Dsx3Totalccvs", dsx3Totalentry.Dsx3Totalccvs}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalCESs"] = types.YLeaf{"Dsx3Totalcess", dsx3Totalentry.Dsx3Totalcess}
    dsx3Totalentry.EntityData.Leafs["dsx3TotalCSESs"] = types.YLeaf{"Dsx3Totalcsess", dsx3Totalentry.Dsx3Totalcsess}
    return &(dsx3Totalentry.EntityData)
}

// DS3MIB_Dsx3Farendconfigtable
// The DS3 Far End Configuration Table contains
// configuration information reported in the C-bits
// from the remote end.
type DS3MIB_Dsx3Farendconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Configuration table. The type is slice of
    // DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry.
    Dsx3Farendconfigentry []DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetEntityData() *types.CommonEntityData {
    dsx3Farendconfigtable.EntityData.YFilter = dsx3Farendconfigtable.YFilter
    dsx3Farendconfigtable.EntityData.YangName = "dsx3FarEndConfigTable"
    dsx3Farendconfigtable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendconfigtable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Farendconfigtable.EntityData.SegmentPath = "dsx3FarEndConfigTable"
    dsx3Farendconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendconfigtable.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendconfigtable.EntityData.Children["dsx3FarEndConfigEntry"] = types.YChild{"Dsx3Farendconfigentry", nil}
    for i := range dsx3Farendconfigtable.Dsx3Farendconfigentry {
        dsx3Farendconfigtable.EntityData.Children[types.GetSegmentPath(&dsx3Farendconfigtable.Dsx3Farendconfigentry[i])] = types.YChild{"Dsx3Farendconfigentry", &dsx3Farendconfigtable.Dsx3Farendconfigentry[i]}
    }
    dsx3Farendconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Farendconfigtable.EntityData)
}

// DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry
// An entry in the DS3 Far End Configuration table.
type DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is the same interface as identified by the
    // same value an dsx3LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx3Farendlineindex interface{}

    // This is the Far End Equipment Identification code that describes the
    // specific piece of equipment. It is sent within the Path Identification
    // Message. The type is string with length: 0..10.
    Dsx3Farendequipcode interface{}

    // This is the Far End Location Identification code that describes the
    // specific location of the equipment.  It is sent within the Path
    // Identification Message. The type is string with length: 0..11.
    Dsx3Farendlocationidcode interface{}

    // This is the Far End Frame Identification code that identifies where the
    // equipment is located within a building at a given location.  It is sent
    // within the Path Identification Message. The type is string with length:
    // 0..10.
    Dsx3Farendframeidcode interface{}

    // This is the Far End code that identifies the equipment location within a
    // bay.  It is sent within the Path Identification Message. The type is string
    // with length: 0..6.
    Dsx3Farendunitcode interface{}

    // This code identifies a specific Far End DS3 path. It is sent within the
    // Path Identification Message. The type is string with length: 0..38.
    Dsx3Farendfacilityidcode interface{}
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetEntityData() *types.CommonEntityData {
    dsx3Farendconfigentry.EntityData.YFilter = dsx3Farendconfigentry.YFilter
    dsx3Farendconfigentry.EntityData.YangName = "dsx3FarEndConfigEntry"
    dsx3Farendconfigentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendconfigentry.EntityData.ParentYangName = "dsx3FarEndConfigTable"
    dsx3Farendconfigentry.EntityData.SegmentPath = "dsx3FarEndConfigEntry" + "[dsx3FarEndLineIndex='" + fmt.Sprintf("%v", dsx3Farendconfigentry.Dsx3Farendlineindex) + "']"
    dsx3Farendconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendconfigentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndLineIndex"] = types.YLeaf{"Dsx3Farendlineindex", dsx3Farendconfigentry.Dsx3Farendlineindex}
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndEquipCode"] = types.YLeaf{"Dsx3Farendequipcode", dsx3Farendconfigentry.Dsx3Farendequipcode}
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndLocationIDCode"] = types.YLeaf{"Dsx3Farendlocationidcode", dsx3Farendconfigentry.Dsx3Farendlocationidcode}
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndFrameIDCode"] = types.YLeaf{"Dsx3Farendframeidcode", dsx3Farendconfigentry.Dsx3Farendframeidcode}
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndUnitCode"] = types.YLeaf{"Dsx3Farendunitcode", dsx3Farendconfigentry.Dsx3Farendunitcode}
    dsx3Farendconfigentry.EntityData.Leafs["dsx3FarEndFacilityIDCode"] = types.YLeaf{"Dsx3Farendfacilityidcode", dsx3Farendconfigentry.Dsx3Farendfacilityidcode}
    return &(dsx3Farendconfigentry.EntityData)
}

// DS3MIB_Dsx3Farendcurrenttable
// The DS3 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end block error code within the C-
// bits.
type DS3MIB_Dsx3Farendcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Current table. The type is slice of
    // DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry.
    Dsx3Farendcurrententry []DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetEntityData() *types.CommonEntityData {
    dsx3Farendcurrenttable.EntityData.YFilter = dsx3Farendcurrenttable.YFilter
    dsx3Farendcurrenttable.EntityData.YangName = "dsx3FarEndCurrentTable"
    dsx3Farendcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendcurrenttable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Farendcurrenttable.EntityData.SegmentPath = "dsx3FarEndCurrentTable"
    dsx3Farendcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendcurrenttable.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendcurrenttable.EntityData.Children["dsx3FarEndCurrentEntry"] = types.YChild{"Dsx3Farendcurrententry", nil}
    for i := range dsx3Farendcurrenttable.Dsx3Farendcurrententry {
        dsx3Farendcurrenttable.EntityData.Children[types.GetSegmentPath(&dsx3Farendcurrenttable.Dsx3Farendcurrententry[i])] = types.YChild{"Dsx3Farendcurrententry", &dsx3Farendcurrenttable.Dsx3Farendcurrententry[i]}
    }
    dsx3Farendcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Farendcurrenttable.EntityData)
}

// DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry
// An entry in the DS3 Far End Current table.
type DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3Farendcurrentindex interface{}

    // The number of seconds that have elapsed since the beginning of the far end
    // current error-measurement period.  If, for some reason, such as an
    // adjustment in the system's time-of-day clock, the current interval exceeds
    // the maximum value, the agent will return the maximum value. The type is
    // interface{} with range: 0..899.
    Dsx3Farendtimeelapsed interface{}

    // The number of previous far end intervals for which data was collected.  The
    // value will be 96 unless the interface was brought online within the last 24
    // hours, in which case the value will be the number of complete 15 minute far
    // end intervals since the interface has been online. The type is interface{}
    // with range: 0..96.
    Dsx3Farendvalidintervals interface{}

    // The counter associated with the number of Far Far End C-bit Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3Farendcurrentcess interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3Farendcurrentcsess interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count. The type is interface{} with
    // range: 0..4294967295.
    Dsx3Farendcurrentccvs interface{}

    // The counter associated with the number of Far End unavailable seconds. The
    // type is interface{} with range: 0..4294967295.
    Dsx3Farendcurrentuass interface{}

    // The number of intervals in the range from 0 to dsx3FarEndValidIntervals for
    // which no data is available.  This object will typically be zero except in
    // cases where the data for some intervals are not available (e.g., in proxy
    // situations). The type is interface{} with range: 0..96.
    Dsx3Farendinvalidintervals interface{}
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetEntityData() *types.CommonEntityData {
    dsx3Farendcurrententry.EntityData.YFilter = dsx3Farendcurrententry.YFilter
    dsx3Farendcurrententry.EntityData.YangName = "dsx3FarEndCurrentEntry"
    dsx3Farendcurrententry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendcurrententry.EntityData.ParentYangName = "dsx3FarEndCurrentTable"
    dsx3Farendcurrententry.EntityData.SegmentPath = "dsx3FarEndCurrentEntry" + "[dsx3FarEndCurrentIndex='" + fmt.Sprintf("%v", dsx3Farendcurrententry.Dsx3Farendcurrentindex) + "']"
    dsx3Farendcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendcurrententry.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndCurrentIndex"] = types.YLeaf{"Dsx3Farendcurrentindex", dsx3Farendcurrententry.Dsx3Farendcurrentindex}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndTimeElapsed"] = types.YLeaf{"Dsx3Farendtimeelapsed", dsx3Farendcurrententry.Dsx3Farendtimeelapsed}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndValidIntervals"] = types.YLeaf{"Dsx3Farendvalidintervals", dsx3Farendcurrententry.Dsx3Farendvalidintervals}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndCurrentCESs"] = types.YLeaf{"Dsx3Farendcurrentcess", dsx3Farendcurrententry.Dsx3Farendcurrentcess}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndCurrentCSESs"] = types.YLeaf{"Dsx3Farendcurrentcsess", dsx3Farendcurrententry.Dsx3Farendcurrentcsess}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndCurrentCCVs"] = types.YLeaf{"Dsx3Farendcurrentccvs", dsx3Farendcurrententry.Dsx3Farendcurrentccvs}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndCurrentUASs"] = types.YLeaf{"Dsx3Farendcurrentuass", dsx3Farendcurrententry.Dsx3Farendcurrentuass}
    dsx3Farendcurrententry.EntityData.Leafs["dsx3FarEndInvalidIntervals"] = types.YLeaf{"Dsx3Farendinvalidintervals", dsx3Farendcurrententry.Dsx3Farendinvalidintervals}
    return &(dsx3Farendcurrententry.EntityData)
}

// DS3MIB_Dsx3Farendintervaltable
// The DS3 Far End Interval Table contains various
// statistics collected by each DS3 interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.
type DS3MIB_Dsx3Farendintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Interval table. The type is slice of
    // DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry.
    Dsx3Farendintervalentry []DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetEntityData() *types.CommonEntityData {
    dsx3Farendintervaltable.EntityData.YFilter = dsx3Farendintervaltable.YFilter
    dsx3Farendintervaltable.EntityData.YangName = "dsx3FarEndIntervalTable"
    dsx3Farendintervaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendintervaltable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Farendintervaltable.EntityData.SegmentPath = "dsx3FarEndIntervalTable"
    dsx3Farendintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendintervaltable.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendintervaltable.EntityData.Children["dsx3FarEndIntervalEntry"] = types.YChild{"Dsx3Farendintervalentry", nil}
    for i := range dsx3Farendintervaltable.Dsx3Farendintervalentry {
        dsx3Farendintervaltable.EntityData.Children[types.GetSegmentPath(&dsx3Farendintervaltable.Dsx3Farendintervalentry[i])] = types.YChild{"Dsx3Farendintervalentry", &dsx3Farendintervaltable.Dsx3Farendintervalentry[i]}
    }
    dsx3Farendintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Farendintervaltable.EntityData)
}

// DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry
// An entry in the DS3 Far End Interval table.
type DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3Farendintervalindex interface{}

    // This attribute is a key. A number between 1 and 96, where 1 is the most
    // recently completed 15 minute interval and 96 is the 15 minutes interval
    // completed 23 hours and 45 minutes prior to interval 1. The type is
    // interface{} with range: 1..96.
    Dsx3Farendintervalnumber interface{}

    // The counter associated with the number of Far End C-bit Errored Seconds
    // encountered by a DS3 interface in one of the previous 96, individual 15
    // minute, intervals. In the case where the agent is a proxy and data is not
    // available, return noSuchInstance. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Farendintervalcess interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds. The type is interface{} with range: 0..4294967295.
    Dsx3Farendintervalcsess interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count. The type is interface{} with
    // range: 0..4294967295.
    Dsx3Farendintervalccvs interface{}

    // The counter associated with the number of Far End unavailable seconds. The
    // type is interface{} with range: 0..4294967295.
    Dsx3Farendintervaluass interface{}

    // This variable indicates if the data for this interval is valid. The type is
    // bool.
    Dsx3Farendintervalvaliddata interface{}
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetEntityData() *types.CommonEntityData {
    dsx3Farendintervalentry.EntityData.YFilter = dsx3Farendintervalentry.YFilter
    dsx3Farendintervalentry.EntityData.YangName = "dsx3FarEndIntervalEntry"
    dsx3Farendintervalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendintervalentry.EntityData.ParentYangName = "dsx3FarEndIntervalTable"
    dsx3Farendintervalentry.EntityData.SegmentPath = "dsx3FarEndIntervalEntry" + "[dsx3FarEndIntervalIndex='" + fmt.Sprintf("%v", dsx3Farendintervalentry.Dsx3Farendintervalindex) + "']" + "[dsx3FarEndIntervalNumber='" + fmt.Sprintf("%v", dsx3Farendintervalentry.Dsx3Farendintervalnumber) + "']"
    dsx3Farendintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendintervalentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalIndex"] = types.YLeaf{"Dsx3Farendintervalindex", dsx3Farendintervalentry.Dsx3Farendintervalindex}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalNumber"] = types.YLeaf{"Dsx3Farendintervalnumber", dsx3Farendintervalentry.Dsx3Farendintervalnumber}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalCESs"] = types.YLeaf{"Dsx3Farendintervalcess", dsx3Farendintervalentry.Dsx3Farendintervalcess}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalCSESs"] = types.YLeaf{"Dsx3Farendintervalcsess", dsx3Farendintervalentry.Dsx3Farendintervalcsess}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalCCVs"] = types.YLeaf{"Dsx3Farendintervalccvs", dsx3Farendintervalentry.Dsx3Farendintervalccvs}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalUASs"] = types.YLeaf{"Dsx3Farendintervaluass", dsx3Farendintervalentry.Dsx3Farendintervaluass}
    dsx3Farendintervalentry.EntityData.Leafs["dsx3FarEndIntervalValidData"] = types.YLeaf{"Dsx3Farendintervalvaliddata", dsx3Farendintervalentry.Dsx3Farendintervalvaliddata}
    return &(dsx3Farendintervalentry.EntityData)
}

// DS3MIB_Dsx3Farendtotaltable
// The DS3 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS3MIB_Dsx3Farendtotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Total table. The type is slice of
    // DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry.
    Dsx3Farendtotalentry []DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetEntityData() *types.CommonEntityData {
    dsx3Farendtotaltable.EntityData.YFilter = dsx3Farendtotaltable.YFilter
    dsx3Farendtotaltable.EntityData.YangName = "dsx3FarEndTotalTable"
    dsx3Farendtotaltable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendtotaltable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Farendtotaltable.EntityData.SegmentPath = "dsx3FarEndTotalTable"
    dsx3Farendtotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendtotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendtotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendtotaltable.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendtotaltable.EntityData.Children["dsx3FarEndTotalEntry"] = types.YChild{"Dsx3Farendtotalentry", nil}
    for i := range dsx3Farendtotaltable.Dsx3Farendtotalentry {
        dsx3Farendtotaltable.EntityData.Children[types.GetSegmentPath(&dsx3Farendtotaltable.Dsx3Farendtotalentry[i])] = types.YChild{"Dsx3Farendtotalentry", &dsx3Farendtotaltable.Dsx3Farendtotalentry[i]}
    }
    dsx3Farendtotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Farendtotaltable.EntityData)
}

// DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry
// An entry in the DS3 Far End Total table.
type DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies the DS3
    // interface to which this entry is applicable.  The interface identified by a
    // particular value of this index is identical to the interface identified by
    // the same value of dsx3LineIndex. The type is interface{} with range:
    // 1..2147483647.
    Dsx3Farendtotalindex interface{}

    // The counter associated with the number of Far End C-bit Errored Seconds
    // encountered by a DS3 interface in the previous 24 hour interval. Invalid 15
    // minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Farendtotalcess interface{}

    // The counter associated with the number of Far End C-bit Severely Errored
    // Seconds encountered by a DS3 interface in the previous 24 hour interval.
    // Invalid 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Farendtotalcsess interface{}

    // The counter associated with the number of Far End C-bit Coding Violations
    // reported via the far end block error count encountered by a DS3 interface
    // in the previous 24 hour interval. Invalid 15 minute intervals count as 0.
    // The type is interface{} with range: 0..4294967295.
    Dsx3Farendtotalccvs interface{}

    // The counter associated with the number of Far End unavailable seconds
    // encountered by a DS3 interface in the previous 24 hour interval.  Invalid
    // 15 minute intervals count as 0. The type is interface{} with range:
    // 0..4294967295.
    Dsx3Farendtotaluass interface{}
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetEntityData() *types.CommonEntityData {
    dsx3Farendtotalentry.EntityData.YFilter = dsx3Farendtotalentry.YFilter
    dsx3Farendtotalentry.EntityData.YangName = "dsx3FarEndTotalEntry"
    dsx3Farendtotalentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Farendtotalentry.EntityData.ParentYangName = "dsx3FarEndTotalTable"
    dsx3Farendtotalentry.EntityData.SegmentPath = "dsx3FarEndTotalEntry" + "[dsx3FarEndTotalIndex='" + fmt.Sprintf("%v", dsx3Farendtotalentry.Dsx3Farendtotalindex) + "']"
    dsx3Farendtotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Farendtotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Farendtotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Farendtotalentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Farendtotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Farendtotalentry.EntityData.Leafs["dsx3FarEndTotalIndex"] = types.YLeaf{"Dsx3Farendtotalindex", dsx3Farendtotalentry.Dsx3Farendtotalindex}
    dsx3Farendtotalentry.EntityData.Leafs["dsx3FarEndTotalCESs"] = types.YLeaf{"Dsx3Farendtotalcess", dsx3Farendtotalentry.Dsx3Farendtotalcess}
    dsx3Farendtotalentry.EntityData.Leafs["dsx3FarEndTotalCSESs"] = types.YLeaf{"Dsx3Farendtotalcsess", dsx3Farendtotalentry.Dsx3Farendtotalcsess}
    dsx3Farendtotalentry.EntityData.Leafs["dsx3FarEndTotalCCVs"] = types.YLeaf{"Dsx3Farendtotalccvs", dsx3Farendtotalentry.Dsx3Farendtotalccvs}
    dsx3Farendtotalentry.EntityData.Leafs["dsx3FarEndTotalUASs"] = types.YLeaf{"Dsx3Farendtotaluass", dsx3Farendtotalentry.Dsx3Farendtotaluass}
    return &(dsx3Farendtotalentry.EntityData)
}

// DS3MIB_Dsx3Fractable
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
type DS3MIB_Dsx3Fractable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DS3 Fractional table. The type is slice of
    // DS3MIB_Dsx3Fractable_Dsx3Fracentry.
    Dsx3Fracentry []DS3MIB_Dsx3Fractable_Dsx3Fracentry
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetEntityData() *types.CommonEntityData {
    dsx3Fractable.EntityData.YFilter = dsx3Fractable.YFilter
    dsx3Fractable.EntityData.YangName = "dsx3FracTable"
    dsx3Fractable.EntityData.BundleName = "cisco_ios_xe"
    dsx3Fractable.EntityData.ParentYangName = "DS3-MIB"
    dsx3Fractable.EntityData.SegmentPath = "dsx3FracTable"
    dsx3Fractable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Fractable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Fractable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Fractable.EntityData.Children = make(map[string]types.YChild)
    dsx3Fractable.EntityData.Children["dsx3FracEntry"] = types.YChild{"Dsx3Fracentry", nil}
    for i := range dsx3Fractable.Dsx3Fracentry {
        dsx3Fractable.EntityData.Children[types.GetSegmentPath(&dsx3Fractable.Dsx3Fracentry[i])] = types.YChild{"Dsx3Fracentry", &dsx3Fractable.Dsx3Fracentry[i]}
    }
    dsx3Fractable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dsx3Fractable.EntityData)
}

// DS3MIB_Dsx3Fractable_Dsx3Fracentry
// An entry in the DS3 Fractional table.
type DS3MIB_Dsx3Fractable_Dsx3Fracentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index value which uniquely identifies  the DS3
    // interface  to which this entry is applicable The interface identified by a 
    // particular value of  this  index is the same interface as identified by the
    // same value  an  dsx3LineIndex object instance. The type is interface{} with
    // range: 1..2147483647.
    Dsx3Fracindex interface{}

    // This attribute is a key. The channel number for this entry. The type is
    // interface{} with range: 1..31.
    Dsx3Fracnumber interface{}

    // An index value that uniquely identifies an interface.  The interface
    // identified by a particular value of this index is the same interface as 
    // identified by the same value an ifIndex object instance. If no interface is
    // currently using a channel, the value should be zero.  If a single interface
    // occupies more  than one  time slot,  that ifIndex value will be found in
    // multiple time slots. The type is interface{} with range: 1..2147483647.
    Dsx3Fracifindex interface{}
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetEntityData() *types.CommonEntityData {
    dsx3Fracentry.EntityData.YFilter = dsx3Fracentry.YFilter
    dsx3Fracentry.EntityData.YangName = "dsx3FracEntry"
    dsx3Fracentry.EntityData.BundleName = "cisco_ios_xe"
    dsx3Fracentry.EntityData.ParentYangName = "dsx3FracTable"
    dsx3Fracentry.EntityData.SegmentPath = "dsx3FracEntry" + "[dsx3FracIndex='" + fmt.Sprintf("%v", dsx3Fracentry.Dsx3Fracindex) + "']" + "[dsx3FracNumber='" + fmt.Sprintf("%v", dsx3Fracentry.Dsx3Fracnumber) + "']"
    dsx3Fracentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsx3Fracentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsx3Fracentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsx3Fracentry.EntityData.Children = make(map[string]types.YChild)
    dsx3Fracentry.EntityData.Leafs = make(map[string]types.YLeaf)
    dsx3Fracentry.EntityData.Leafs["dsx3FracIndex"] = types.YLeaf{"Dsx3Fracindex", dsx3Fracentry.Dsx3Fracindex}
    dsx3Fracentry.EntityData.Leafs["dsx3FracNumber"] = types.YLeaf{"Dsx3Fracnumber", dsx3Fracentry.Dsx3Fracnumber}
    dsx3Fracentry.EntityData.Leafs["dsx3FracIfIndex"] = types.YLeaf{"Dsx3Fracifindex", dsx3Fracentry.Dsx3Fracifindex}
    return &(dsx3Fracentry.EntityData)
}

