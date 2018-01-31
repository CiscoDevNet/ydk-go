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
    parent types.Entity
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

func (dS3MIB *DS3MIB) GetFilter() yfilter.YFilter { return dS3MIB.YFilter }

func (dS3MIB *DS3MIB) SetFilter(yf yfilter.YFilter) { dS3MIB.YFilter = yf }

func (dS3MIB *DS3MIB) GetGoName(yname string) string {
    if yname == "dsx3ConfigTable" { return "Dsx3Configtable" }
    if yname == "dsx3CurrentTable" { return "Dsx3Currenttable" }
    if yname == "dsx3IntervalTable" { return "Dsx3Intervaltable" }
    if yname == "dsx3TotalTable" { return "Dsx3Totaltable" }
    if yname == "dsx3FarEndConfigTable" { return "Dsx3Farendconfigtable" }
    if yname == "dsx3FarEndCurrentTable" { return "Dsx3Farendcurrenttable" }
    if yname == "dsx3FarEndIntervalTable" { return "Dsx3Farendintervaltable" }
    if yname == "dsx3FarEndTotalTable" { return "Dsx3Farendtotaltable" }
    if yname == "dsx3FracTable" { return "Dsx3Fractable" }
    return ""
}

func (dS3MIB *DS3MIB) GetSegmentPath() string {
    return "DS3-MIB:DS3-MIB"
}

func (dS3MIB *DS3MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3ConfigTable" {
        return &dS3MIB.Dsx3Configtable
    }
    if childYangName == "dsx3CurrentTable" {
        return &dS3MIB.Dsx3Currenttable
    }
    if childYangName == "dsx3IntervalTable" {
        return &dS3MIB.Dsx3Intervaltable
    }
    if childYangName == "dsx3TotalTable" {
        return &dS3MIB.Dsx3Totaltable
    }
    if childYangName == "dsx3FarEndConfigTable" {
        return &dS3MIB.Dsx3Farendconfigtable
    }
    if childYangName == "dsx3FarEndCurrentTable" {
        return &dS3MIB.Dsx3Farendcurrenttable
    }
    if childYangName == "dsx3FarEndIntervalTable" {
        return &dS3MIB.Dsx3Farendintervaltable
    }
    if childYangName == "dsx3FarEndTotalTable" {
        return &dS3MIB.Dsx3Farendtotaltable
    }
    if childYangName == "dsx3FracTable" {
        return &dS3MIB.Dsx3Fractable
    }
    return nil
}

func (dS3MIB *DS3MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dsx3ConfigTable"] = &dS3MIB.Dsx3Configtable
    children["dsx3CurrentTable"] = &dS3MIB.Dsx3Currenttable
    children["dsx3IntervalTable"] = &dS3MIB.Dsx3Intervaltable
    children["dsx3TotalTable"] = &dS3MIB.Dsx3Totaltable
    children["dsx3FarEndConfigTable"] = &dS3MIB.Dsx3Farendconfigtable
    children["dsx3FarEndCurrentTable"] = &dS3MIB.Dsx3Farendcurrenttable
    children["dsx3FarEndIntervalTable"] = &dS3MIB.Dsx3Farendintervaltable
    children["dsx3FarEndTotalTable"] = &dS3MIB.Dsx3Farendtotaltable
    children["dsx3FracTable"] = &dS3MIB.Dsx3Fractable
    return children
}

func (dS3MIB *DS3MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dS3MIB *DS3MIB) GetBundleName() string { return "cisco_ios_xe" }

func (dS3MIB *DS3MIB) GetYangName() string { return "DS3-MIB" }

func (dS3MIB *DS3MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dS3MIB *DS3MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dS3MIB *DS3MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dS3MIB *DS3MIB) SetParent(parent types.Entity) { dS3MIB.parent = parent }

func (dS3MIB *DS3MIB) GetParent() types.Entity { return dS3MIB.parent }

func (dS3MIB *DS3MIB) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Configtable
// The DS3/E3 Configuration table.
type DS3MIB_Dsx3Configtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Configuration table. The type is slice of
    // DS3MIB_Dsx3Configtable_Dsx3Configentry.
    Dsx3Configentry []DS3MIB_Dsx3Configtable_Dsx3Configentry
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetFilter() yfilter.YFilter { return dsx3Configtable.YFilter }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) SetFilter(yf yfilter.YFilter) { dsx3Configtable.YFilter = yf }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetGoName(yname string) string {
    if yname == "dsx3ConfigEntry" { return "Dsx3Configentry" }
    return ""
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetSegmentPath() string {
    return "dsx3ConfigTable"
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3ConfigEntry" {
        for _, c := range dsx3Configtable.Dsx3Configentry {
            if dsx3Configtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Configtable_Dsx3Configentry{}
        dsx3Configtable.Dsx3Configentry = append(dsx3Configtable.Dsx3Configentry, child)
        return &dsx3Configtable.Dsx3Configentry[len(dsx3Configtable.Dsx3Configentry)-1]
    }
    return nil
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Configtable.Dsx3Configentry {
        children[dsx3Configtable.Dsx3Configentry[i].GetSegmentPath()] = &dsx3Configtable.Dsx3Configentry[i]
    }
    return children
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetYangName() string { return "dsx3ConfigTable" }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) SetParent(parent types.Entity) { dsx3Configtable.parent = parent }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetParent() types.Entity { return dsx3Configtable.parent }

func (dsx3Configtable *DS3MIB_Dsx3Configtable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Configtable_Dsx3Configentry
// An entry in the DS3/E3 Configuration table.
type DS3MIB_Dsx3Configtable_Dsx3Configentry struct {
    parent types.Entity
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

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetFilter() yfilter.YFilter { return dsx3Configentry.YFilter }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) SetFilter(yf yfilter.YFilter) { dsx3Configentry.YFilter = yf }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetGoName(yname string) string {
    if yname == "dsx3LineIndex" { return "Dsx3Lineindex" }
    if yname == "dsx3IfIndex" { return "Dsx3Ifindex" }
    if yname == "dsx3TimeElapsed" { return "Dsx3Timeelapsed" }
    if yname == "dsx3ValidIntervals" { return "Dsx3Validintervals" }
    if yname == "dsx3LineType" { return "Dsx3Linetype" }
    if yname == "dsx3LineCoding" { return "Dsx3Linecoding" }
    if yname == "dsx3SendCode" { return "Dsx3Sendcode" }
    if yname == "dsx3CircuitIdentifier" { return "Dsx3Circuitidentifier" }
    if yname == "dsx3LoopbackConfig" { return "Dsx3Loopbackconfig" }
    if yname == "dsx3LineStatus" { return "Dsx3Linestatus" }
    if yname == "dsx3TransmitClockSource" { return "Dsx3Transmitclocksource" }
    if yname == "dsx3InvalidIntervals" { return "Dsx3Invalidintervals" }
    if yname == "dsx3LineLength" { return "Dsx3Linelength" }
    if yname == "dsx3LineStatusLastChange" { return "Dsx3Linestatuslastchange" }
    if yname == "dsx3LineStatusChangeTrapEnable" { return "Dsx3Linestatuschangetrapenable" }
    if yname == "dsx3LoopbackStatus" { return "Dsx3Loopbackstatus" }
    if yname == "dsx3Channelization" { return "Dsx3Channelization" }
    if yname == "dsx3Ds1ForRemoteLoop" { return "Dsx3Ds1Forremoteloop" }
    return ""
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetSegmentPath() string {
    return "dsx3ConfigEntry" + "[dsx3LineIndex='" + fmt.Sprintf("%v", dsx3Configentry.Dsx3Lineindex) + "']"
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3LineIndex"] = dsx3Configentry.Dsx3Lineindex
    leafs["dsx3IfIndex"] = dsx3Configentry.Dsx3Ifindex
    leafs["dsx3TimeElapsed"] = dsx3Configentry.Dsx3Timeelapsed
    leafs["dsx3ValidIntervals"] = dsx3Configentry.Dsx3Validintervals
    leafs["dsx3LineType"] = dsx3Configentry.Dsx3Linetype
    leafs["dsx3LineCoding"] = dsx3Configentry.Dsx3Linecoding
    leafs["dsx3SendCode"] = dsx3Configentry.Dsx3Sendcode
    leafs["dsx3CircuitIdentifier"] = dsx3Configentry.Dsx3Circuitidentifier
    leafs["dsx3LoopbackConfig"] = dsx3Configentry.Dsx3Loopbackconfig
    leafs["dsx3LineStatus"] = dsx3Configentry.Dsx3Linestatus
    leafs["dsx3TransmitClockSource"] = dsx3Configentry.Dsx3Transmitclocksource
    leafs["dsx3InvalidIntervals"] = dsx3Configentry.Dsx3Invalidintervals
    leafs["dsx3LineLength"] = dsx3Configentry.Dsx3Linelength
    leafs["dsx3LineStatusLastChange"] = dsx3Configentry.Dsx3Linestatuslastchange
    leafs["dsx3LineStatusChangeTrapEnable"] = dsx3Configentry.Dsx3Linestatuschangetrapenable
    leafs["dsx3LoopbackStatus"] = dsx3Configentry.Dsx3Loopbackstatus
    leafs["dsx3Channelization"] = dsx3Configentry.Dsx3Channelization
    leafs["dsx3Ds1ForRemoteLoop"] = dsx3Configentry.Dsx3Ds1Forremoteloop
    return leafs
}

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetYangName() string { return "dsx3ConfigEntry" }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) SetParent(parent types.Entity) { dsx3Configentry.parent = parent }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetParent() types.Entity { return dsx3Configentry.parent }

func (dsx3Configentry *DS3MIB_Dsx3Configtable_Dsx3Configentry) GetParentYangName() string { return "dsx3ConfigTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Current table. The type is slice of
    // DS3MIB_Dsx3Currenttable_Dsx3Currententry.
    Dsx3Currententry []DS3MIB_Dsx3Currenttable_Dsx3Currententry
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetFilter() yfilter.YFilter { return dsx3Currenttable.YFilter }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) SetFilter(yf yfilter.YFilter) { dsx3Currenttable.YFilter = yf }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetGoName(yname string) string {
    if yname == "dsx3CurrentEntry" { return "Dsx3Currententry" }
    return ""
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetSegmentPath() string {
    return "dsx3CurrentTable"
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3CurrentEntry" {
        for _, c := range dsx3Currenttable.Dsx3Currententry {
            if dsx3Currenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Currenttable_Dsx3Currententry{}
        dsx3Currenttable.Dsx3Currententry = append(dsx3Currenttable.Dsx3Currententry, child)
        return &dsx3Currenttable.Dsx3Currententry[len(dsx3Currenttable.Dsx3Currententry)-1]
    }
    return nil
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Currenttable.Dsx3Currententry {
        children[dsx3Currenttable.Dsx3Currententry[i].GetSegmentPath()] = &dsx3Currenttable.Dsx3Currententry[i]
    }
    return children
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetYangName() string { return "dsx3CurrentTable" }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) SetParent(parent types.Entity) { dsx3Currenttable.parent = parent }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetParent() types.Entity { return dsx3Currenttable.parent }

func (dsx3Currenttable *DS3MIB_Dsx3Currenttable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Currenttable_Dsx3Currententry
// An entry in the DS3/E3 Current table.
type DS3MIB_Dsx3Currenttable_Dsx3Currententry struct {
    parent types.Entity
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

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetFilter() yfilter.YFilter { return dsx3Currententry.YFilter }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) SetFilter(yf yfilter.YFilter) { dsx3Currententry.YFilter = yf }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetGoName(yname string) string {
    if yname == "dsx3CurrentIndex" { return "Dsx3Currentindex" }
    if yname == "dsx3CurrentPESs" { return "Dsx3Currentpess" }
    if yname == "dsx3CurrentPSESs" { return "Dsx3Currentpsess" }
    if yname == "dsx3CurrentSEFSs" { return "Dsx3Currentsefss" }
    if yname == "dsx3CurrentUASs" { return "Dsx3Currentuass" }
    if yname == "dsx3CurrentLCVs" { return "Dsx3Currentlcvs" }
    if yname == "dsx3CurrentPCVs" { return "Dsx3Currentpcvs" }
    if yname == "dsx3CurrentLESs" { return "Dsx3Currentless" }
    if yname == "dsx3CurrentCCVs" { return "Dsx3Currentccvs" }
    if yname == "dsx3CurrentCESs" { return "Dsx3Currentcess" }
    if yname == "dsx3CurrentCSESs" { return "Dsx3Currentcsess" }
    return ""
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetSegmentPath() string {
    return "dsx3CurrentEntry" + "[dsx3CurrentIndex='" + fmt.Sprintf("%v", dsx3Currententry.Dsx3Currentindex) + "']"
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3CurrentIndex"] = dsx3Currententry.Dsx3Currentindex
    leafs["dsx3CurrentPESs"] = dsx3Currententry.Dsx3Currentpess
    leafs["dsx3CurrentPSESs"] = dsx3Currententry.Dsx3Currentpsess
    leafs["dsx3CurrentSEFSs"] = dsx3Currententry.Dsx3Currentsefss
    leafs["dsx3CurrentUASs"] = dsx3Currententry.Dsx3Currentuass
    leafs["dsx3CurrentLCVs"] = dsx3Currententry.Dsx3Currentlcvs
    leafs["dsx3CurrentPCVs"] = dsx3Currententry.Dsx3Currentpcvs
    leafs["dsx3CurrentLESs"] = dsx3Currententry.Dsx3Currentless
    leafs["dsx3CurrentCCVs"] = dsx3Currententry.Dsx3Currentccvs
    leafs["dsx3CurrentCESs"] = dsx3Currententry.Dsx3Currentcess
    leafs["dsx3CurrentCSESs"] = dsx3Currententry.Dsx3Currentcsess
    return leafs
}

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetYangName() string { return "dsx3CurrentEntry" }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) SetParent(parent types.Entity) { dsx3Currententry.parent = parent }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetParent() types.Entity { return dsx3Currententry.parent }

func (dsx3Currententry *DS3MIB_Dsx3Currenttable_Dsx3Currententry) GetParentYangName() string { return "dsx3CurrentTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Interval table. The type is slice of
    // DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry.
    Dsx3Intervalentry []DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetFilter() yfilter.YFilter { return dsx3Intervaltable.YFilter }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) SetFilter(yf yfilter.YFilter) { dsx3Intervaltable.YFilter = yf }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetGoName(yname string) string {
    if yname == "dsx3IntervalEntry" { return "Dsx3Intervalentry" }
    return ""
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetSegmentPath() string {
    return "dsx3IntervalTable"
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3IntervalEntry" {
        for _, c := range dsx3Intervaltable.Dsx3Intervalentry {
            if dsx3Intervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry{}
        dsx3Intervaltable.Dsx3Intervalentry = append(dsx3Intervaltable.Dsx3Intervalentry, child)
        return &dsx3Intervaltable.Dsx3Intervalentry[len(dsx3Intervaltable.Dsx3Intervalentry)-1]
    }
    return nil
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Intervaltable.Dsx3Intervalentry {
        children[dsx3Intervaltable.Dsx3Intervalentry[i].GetSegmentPath()] = &dsx3Intervaltable.Dsx3Intervalentry[i]
    }
    return children
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetYangName() string { return "dsx3IntervalTable" }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) SetParent(parent types.Entity) { dsx3Intervaltable.parent = parent }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetParent() types.Entity { return dsx3Intervaltable.parent }

func (dsx3Intervaltable *DS3MIB_Dsx3Intervaltable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry
// An entry in the DS3/E3 Interval table.
type DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry struct {
    parent types.Entity
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

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetFilter() yfilter.YFilter { return dsx3Intervalentry.YFilter }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) SetFilter(yf yfilter.YFilter) { dsx3Intervalentry.YFilter = yf }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetGoName(yname string) string {
    if yname == "dsx3IntervalIndex" { return "Dsx3Intervalindex" }
    if yname == "dsx3IntervalNumber" { return "Dsx3Intervalnumber" }
    if yname == "dsx3IntervalPESs" { return "Dsx3Intervalpess" }
    if yname == "dsx3IntervalPSESs" { return "Dsx3Intervalpsess" }
    if yname == "dsx3IntervalSEFSs" { return "Dsx3Intervalsefss" }
    if yname == "dsx3IntervalUASs" { return "Dsx3Intervaluass" }
    if yname == "dsx3IntervalLCVs" { return "Dsx3Intervallcvs" }
    if yname == "dsx3IntervalPCVs" { return "Dsx3Intervalpcvs" }
    if yname == "dsx3IntervalLESs" { return "Dsx3Intervalless" }
    if yname == "dsx3IntervalCCVs" { return "Dsx3Intervalccvs" }
    if yname == "dsx3IntervalCESs" { return "Dsx3Intervalcess" }
    if yname == "dsx3IntervalCSESs" { return "Dsx3Intervalcsess" }
    if yname == "dsx3IntervalValidData" { return "Dsx3Intervalvaliddata" }
    return ""
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetSegmentPath() string {
    return "dsx3IntervalEntry" + "[dsx3IntervalIndex='" + fmt.Sprintf("%v", dsx3Intervalentry.Dsx3Intervalindex) + "']" + "[dsx3IntervalNumber='" + fmt.Sprintf("%v", dsx3Intervalentry.Dsx3Intervalnumber) + "']"
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3IntervalIndex"] = dsx3Intervalentry.Dsx3Intervalindex
    leafs["dsx3IntervalNumber"] = dsx3Intervalentry.Dsx3Intervalnumber
    leafs["dsx3IntervalPESs"] = dsx3Intervalentry.Dsx3Intervalpess
    leafs["dsx3IntervalPSESs"] = dsx3Intervalentry.Dsx3Intervalpsess
    leafs["dsx3IntervalSEFSs"] = dsx3Intervalentry.Dsx3Intervalsefss
    leafs["dsx3IntervalUASs"] = dsx3Intervalentry.Dsx3Intervaluass
    leafs["dsx3IntervalLCVs"] = dsx3Intervalentry.Dsx3Intervallcvs
    leafs["dsx3IntervalPCVs"] = dsx3Intervalentry.Dsx3Intervalpcvs
    leafs["dsx3IntervalLESs"] = dsx3Intervalentry.Dsx3Intervalless
    leafs["dsx3IntervalCCVs"] = dsx3Intervalentry.Dsx3Intervalccvs
    leafs["dsx3IntervalCESs"] = dsx3Intervalentry.Dsx3Intervalcess
    leafs["dsx3IntervalCSESs"] = dsx3Intervalentry.Dsx3Intervalcsess
    leafs["dsx3IntervalValidData"] = dsx3Intervalentry.Dsx3Intervalvaliddata
    return leafs
}

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetYangName() string { return "dsx3IntervalEntry" }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) SetParent(parent types.Entity) { dsx3Intervalentry.parent = parent }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetParent() types.Entity { return dsx3Intervalentry.parent }

func (dsx3Intervalentry *DS3MIB_Dsx3Intervaltable_Dsx3Intervalentry) GetParentYangName() string { return "dsx3IntervalTable" }

// DS3MIB_Dsx3Totaltable
// The DS3/E3 Total Table contains the cumulative
// sum of the various statistics for the 24 hour
// period preceding the current interval.
type DS3MIB_Dsx3Totaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3/E3 Total table. The type is slice of
    // DS3MIB_Dsx3Totaltable_Dsx3Totalentry.
    Dsx3Totalentry []DS3MIB_Dsx3Totaltable_Dsx3Totalentry
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetFilter() yfilter.YFilter { return dsx3Totaltable.YFilter }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) SetFilter(yf yfilter.YFilter) { dsx3Totaltable.YFilter = yf }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetGoName(yname string) string {
    if yname == "dsx3TotalEntry" { return "Dsx3Totalentry" }
    return ""
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetSegmentPath() string {
    return "dsx3TotalTable"
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3TotalEntry" {
        for _, c := range dsx3Totaltable.Dsx3Totalentry {
            if dsx3Totaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Totaltable_Dsx3Totalentry{}
        dsx3Totaltable.Dsx3Totalentry = append(dsx3Totaltable.Dsx3Totalentry, child)
        return &dsx3Totaltable.Dsx3Totalentry[len(dsx3Totaltable.Dsx3Totalentry)-1]
    }
    return nil
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Totaltable.Dsx3Totalentry {
        children[dsx3Totaltable.Dsx3Totalentry[i].GetSegmentPath()] = &dsx3Totaltable.Dsx3Totalentry[i]
    }
    return children
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetYangName() string { return "dsx3TotalTable" }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) SetParent(parent types.Entity) { dsx3Totaltable.parent = parent }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetParent() types.Entity { return dsx3Totaltable.parent }

func (dsx3Totaltable *DS3MIB_Dsx3Totaltable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Totaltable_Dsx3Totalentry
// An entry in the DS3/E3 Total table.
type DS3MIB_Dsx3Totaltable_Dsx3Totalentry struct {
    parent types.Entity
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

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetFilter() yfilter.YFilter { return dsx3Totalentry.YFilter }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) SetFilter(yf yfilter.YFilter) { dsx3Totalentry.YFilter = yf }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetGoName(yname string) string {
    if yname == "dsx3TotalIndex" { return "Dsx3Totalindex" }
    if yname == "dsx3TotalPESs" { return "Dsx3Totalpess" }
    if yname == "dsx3TotalPSESs" { return "Dsx3Totalpsess" }
    if yname == "dsx3TotalSEFSs" { return "Dsx3Totalsefss" }
    if yname == "dsx3TotalUASs" { return "Dsx3Totaluass" }
    if yname == "dsx3TotalLCVs" { return "Dsx3Totallcvs" }
    if yname == "dsx3TotalPCVs" { return "Dsx3Totalpcvs" }
    if yname == "dsx3TotalLESs" { return "Dsx3Totalless" }
    if yname == "dsx3TotalCCVs" { return "Dsx3Totalccvs" }
    if yname == "dsx3TotalCESs" { return "Dsx3Totalcess" }
    if yname == "dsx3TotalCSESs" { return "Dsx3Totalcsess" }
    return ""
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetSegmentPath() string {
    return "dsx3TotalEntry" + "[dsx3TotalIndex='" + fmt.Sprintf("%v", dsx3Totalentry.Dsx3Totalindex) + "']"
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3TotalIndex"] = dsx3Totalentry.Dsx3Totalindex
    leafs["dsx3TotalPESs"] = dsx3Totalentry.Dsx3Totalpess
    leafs["dsx3TotalPSESs"] = dsx3Totalentry.Dsx3Totalpsess
    leafs["dsx3TotalSEFSs"] = dsx3Totalentry.Dsx3Totalsefss
    leafs["dsx3TotalUASs"] = dsx3Totalentry.Dsx3Totaluass
    leafs["dsx3TotalLCVs"] = dsx3Totalentry.Dsx3Totallcvs
    leafs["dsx3TotalPCVs"] = dsx3Totalentry.Dsx3Totalpcvs
    leafs["dsx3TotalLESs"] = dsx3Totalentry.Dsx3Totalless
    leafs["dsx3TotalCCVs"] = dsx3Totalentry.Dsx3Totalccvs
    leafs["dsx3TotalCESs"] = dsx3Totalentry.Dsx3Totalcess
    leafs["dsx3TotalCSESs"] = dsx3Totalentry.Dsx3Totalcsess
    return leafs
}

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetYangName() string { return "dsx3TotalEntry" }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) SetParent(parent types.Entity) { dsx3Totalentry.parent = parent }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetParent() types.Entity { return dsx3Totalentry.parent }

func (dsx3Totalentry *DS3MIB_Dsx3Totaltable_Dsx3Totalentry) GetParentYangName() string { return "dsx3TotalTable" }

// DS3MIB_Dsx3Farendconfigtable
// The DS3 Far End Configuration Table contains
// configuration information reported in the C-bits
// from the remote end.
type DS3MIB_Dsx3Farendconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Configuration table. The type is slice of
    // DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry.
    Dsx3Farendconfigentry []DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetFilter() yfilter.YFilter { return dsx3Farendconfigtable.YFilter }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) SetFilter(yf yfilter.YFilter) { dsx3Farendconfigtable.YFilter = yf }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetGoName(yname string) string {
    if yname == "dsx3FarEndConfigEntry" { return "Dsx3Farendconfigentry" }
    return ""
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetSegmentPath() string {
    return "dsx3FarEndConfigTable"
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3FarEndConfigEntry" {
        for _, c := range dsx3Farendconfigtable.Dsx3Farendconfigentry {
            if dsx3Farendconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry{}
        dsx3Farendconfigtable.Dsx3Farendconfigentry = append(dsx3Farendconfigtable.Dsx3Farendconfigentry, child)
        return &dsx3Farendconfigtable.Dsx3Farendconfigentry[len(dsx3Farendconfigtable.Dsx3Farendconfigentry)-1]
    }
    return nil
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Farendconfigtable.Dsx3Farendconfigentry {
        children[dsx3Farendconfigtable.Dsx3Farendconfigentry[i].GetSegmentPath()] = &dsx3Farendconfigtable.Dsx3Farendconfigentry[i]
    }
    return children
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetYangName() string { return "dsx3FarEndConfigTable" }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) SetParent(parent types.Entity) { dsx3Farendconfigtable.parent = parent }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetParent() types.Entity { return dsx3Farendconfigtable.parent }

func (dsx3Farendconfigtable *DS3MIB_Dsx3Farendconfigtable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry
// An entry in the DS3 Far End Configuration table.
type DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry struct {
    parent types.Entity
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

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetFilter() yfilter.YFilter { return dsx3Farendconfigentry.YFilter }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) SetFilter(yf yfilter.YFilter) { dsx3Farendconfigentry.YFilter = yf }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetGoName(yname string) string {
    if yname == "dsx3FarEndLineIndex" { return "Dsx3Farendlineindex" }
    if yname == "dsx3FarEndEquipCode" { return "Dsx3Farendequipcode" }
    if yname == "dsx3FarEndLocationIDCode" { return "Dsx3Farendlocationidcode" }
    if yname == "dsx3FarEndFrameIDCode" { return "Dsx3Farendframeidcode" }
    if yname == "dsx3FarEndUnitCode" { return "Dsx3Farendunitcode" }
    if yname == "dsx3FarEndFacilityIDCode" { return "Dsx3Farendfacilityidcode" }
    return ""
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetSegmentPath() string {
    return "dsx3FarEndConfigEntry" + "[dsx3FarEndLineIndex='" + fmt.Sprintf("%v", dsx3Farendconfigentry.Dsx3Farendlineindex) + "']"
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3FarEndLineIndex"] = dsx3Farendconfigentry.Dsx3Farendlineindex
    leafs["dsx3FarEndEquipCode"] = dsx3Farendconfigentry.Dsx3Farendequipcode
    leafs["dsx3FarEndLocationIDCode"] = dsx3Farendconfigentry.Dsx3Farendlocationidcode
    leafs["dsx3FarEndFrameIDCode"] = dsx3Farendconfigentry.Dsx3Farendframeidcode
    leafs["dsx3FarEndUnitCode"] = dsx3Farendconfigentry.Dsx3Farendunitcode
    leafs["dsx3FarEndFacilityIDCode"] = dsx3Farendconfigentry.Dsx3Farendfacilityidcode
    return leafs
}

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetYangName() string { return "dsx3FarEndConfigEntry" }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) SetParent(parent types.Entity) { dsx3Farendconfigentry.parent = parent }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetParent() types.Entity { return dsx3Farendconfigentry.parent }

func (dsx3Farendconfigentry *DS3MIB_Dsx3Farendconfigtable_Dsx3Farendconfigentry) GetParentYangName() string { return "dsx3FarEndConfigTable" }

// DS3MIB_Dsx3Farendcurrenttable
// The DS3 Far End Current table contains various
// statistics being collected for the current 15
// minute interval.  The statistics are collected
// from the far end block error code within the C-
// bits.
type DS3MIB_Dsx3Farendcurrenttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Current table. The type is slice of
    // DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry.
    Dsx3Farendcurrententry []DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetFilter() yfilter.YFilter { return dsx3Farendcurrenttable.YFilter }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) SetFilter(yf yfilter.YFilter) { dsx3Farendcurrenttable.YFilter = yf }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetGoName(yname string) string {
    if yname == "dsx3FarEndCurrentEntry" { return "Dsx3Farendcurrententry" }
    return ""
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetSegmentPath() string {
    return "dsx3FarEndCurrentTable"
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3FarEndCurrentEntry" {
        for _, c := range dsx3Farendcurrenttable.Dsx3Farendcurrententry {
            if dsx3Farendcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry{}
        dsx3Farendcurrenttable.Dsx3Farendcurrententry = append(dsx3Farendcurrenttable.Dsx3Farendcurrententry, child)
        return &dsx3Farendcurrenttable.Dsx3Farendcurrententry[len(dsx3Farendcurrenttable.Dsx3Farendcurrententry)-1]
    }
    return nil
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Farendcurrenttable.Dsx3Farendcurrententry {
        children[dsx3Farendcurrenttable.Dsx3Farendcurrententry[i].GetSegmentPath()] = &dsx3Farendcurrenttable.Dsx3Farendcurrententry[i]
    }
    return children
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetYangName() string { return "dsx3FarEndCurrentTable" }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) SetParent(parent types.Entity) { dsx3Farendcurrenttable.parent = parent }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetParent() types.Entity { return dsx3Farendcurrenttable.parent }

func (dsx3Farendcurrenttable *DS3MIB_Dsx3Farendcurrenttable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry
// An entry in the DS3 Far End Current table.
type DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry struct {
    parent types.Entity
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

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetFilter() yfilter.YFilter { return dsx3Farendcurrententry.YFilter }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) SetFilter(yf yfilter.YFilter) { dsx3Farendcurrententry.YFilter = yf }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetGoName(yname string) string {
    if yname == "dsx3FarEndCurrentIndex" { return "Dsx3Farendcurrentindex" }
    if yname == "dsx3FarEndTimeElapsed" { return "Dsx3Farendtimeelapsed" }
    if yname == "dsx3FarEndValidIntervals" { return "Dsx3Farendvalidintervals" }
    if yname == "dsx3FarEndCurrentCESs" { return "Dsx3Farendcurrentcess" }
    if yname == "dsx3FarEndCurrentCSESs" { return "Dsx3Farendcurrentcsess" }
    if yname == "dsx3FarEndCurrentCCVs" { return "Dsx3Farendcurrentccvs" }
    if yname == "dsx3FarEndCurrentUASs" { return "Dsx3Farendcurrentuass" }
    if yname == "dsx3FarEndInvalidIntervals" { return "Dsx3Farendinvalidintervals" }
    return ""
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetSegmentPath() string {
    return "dsx3FarEndCurrentEntry" + "[dsx3FarEndCurrentIndex='" + fmt.Sprintf("%v", dsx3Farendcurrententry.Dsx3Farendcurrentindex) + "']"
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3FarEndCurrentIndex"] = dsx3Farendcurrententry.Dsx3Farendcurrentindex
    leafs["dsx3FarEndTimeElapsed"] = dsx3Farendcurrententry.Dsx3Farendtimeelapsed
    leafs["dsx3FarEndValidIntervals"] = dsx3Farendcurrententry.Dsx3Farendvalidintervals
    leafs["dsx3FarEndCurrentCESs"] = dsx3Farendcurrententry.Dsx3Farendcurrentcess
    leafs["dsx3FarEndCurrentCSESs"] = dsx3Farendcurrententry.Dsx3Farendcurrentcsess
    leafs["dsx3FarEndCurrentCCVs"] = dsx3Farendcurrententry.Dsx3Farendcurrentccvs
    leafs["dsx3FarEndCurrentUASs"] = dsx3Farendcurrententry.Dsx3Farendcurrentuass
    leafs["dsx3FarEndInvalidIntervals"] = dsx3Farendcurrententry.Dsx3Farendinvalidintervals
    return leafs
}

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetYangName() string { return "dsx3FarEndCurrentEntry" }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) SetParent(parent types.Entity) { dsx3Farendcurrententry.parent = parent }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetParent() types.Entity { return dsx3Farendcurrententry.parent }

func (dsx3Farendcurrententry *DS3MIB_Dsx3Farendcurrenttable_Dsx3Farendcurrententry) GetParentYangName() string { return "dsx3FarEndCurrentTable" }

// DS3MIB_Dsx3Farendintervaltable
// The DS3 Far End Interval Table contains various
// statistics collected by each DS3 interface over
// the previous 24 hours of operation.  The past 24
// hours are broken into 96 completed 15 minute
// intervals.
type DS3MIB_Dsx3Farendintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Interval table. The type is slice of
    // DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry.
    Dsx3Farendintervalentry []DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetFilter() yfilter.YFilter { return dsx3Farendintervaltable.YFilter }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) SetFilter(yf yfilter.YFilter) { dsx3Farendintervaltable.YFilter = yf }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetGoName(yname string) string {
    if yname == "dsx3FarEndIntervalEntry" { return "Dsx3Farendintervalentry" }
    return ""
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetSegmentPath() string {
    return "dsx3FarEndIntervalTable"
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3FarEndIntervalEntry" {
        for _, c := range dsx3Farendintervaltable.Dsx3Farendintervalentry {
            if dsx3Farendintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry{}
        dsx3Farendintervaltable.Dsx3Farendintervalentry = append(dsx3Farendintervaltable.Dsx3Farendintervalentry, child)
        return &dsx3Farendintervaltable.Dsx3Farendintervalentry[len(dsx3Farendintervaltable.Dsx3Farendintervalentry)-1]
    }
    return nil
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Farendintervaltable.Dsx3Farendintervalentry {
        children[dsx3Farendintervaltable.Dsx3Farendintervalentry[i].GetSegmentPath()] = &dsx3Farendintervaltable.Dsx3Farendintervalentry[i]
    }
    return children
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetYangName() string { return "dsx3FarEndIntervalTable" }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) SetParent(parent types.Entity) { dsx3Farendintervaltable.parent = parent }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetParent() types.Entity { return dsx3Farendintervaltable.parent }

func (dsx3Farendintervaltable *DS3MIB_Dsx3Farendintervaltable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry
// An entry in the DS3 Far End Interval table.
type DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry struct {
    parent types.Entity
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

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetFilter() yfilter.YFilter { return dsx3Farendintervalentry.YFilter }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) SetFilter(yf yfilter.YFilter) { dsx3Farendintervalentry.YFilter = yf }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetGoName(yname string) string {
    if yname == "dsx3FarEndIntervalIndex" { return "Dsx3Farendintervalindex" }
    if yname == "dsx3FarEndIntervalNumber" { return "Dsx3Farendintervalnumber" }
    if yname == "dsx3FarEndIntervalCESs" { return "Dsx3Farendintervalcess" }
    if yname == "dsx3FarEndIntervalCSESs" { return "Dsx3Farendintervalcsess" }
    if yname == "dsx3FarEndIntervalCCVs" { return "Dsx3Farendintervalccvs" }
    if yname == "dsx3FarEndIntervalUASs" { return "Dsx3Farendintervaluass" }
    if yname == "dsx3FarEndIntervalValidData" { return "Dsx3Farendintervalvaliddata" }
    return ""
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetSegmentPath() string {
    return "dsx3FarEndIntervalEntry" + "[dsx3FarEndIntervalIndex='" + fmt.Sprintf("%v", dsx3Farendintervalentry.Dsx3Farendintervalindex) + "']" + "[dsx3FarEndIntervalNumber='" + fmt.Sprintf("%v", dsx3Farendintervalentry.Dsx3Farendintervalnumber) + "']"
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3FarEndIntervalIndex"] = dsx3Farendintervalentry.Dsx3Farendintervalindex
    leafs["dsx3FarEndIntervalNumber"] = dsx3Farendintervalentry.Dsx3Farendintervalnumber
    leafs["dsx3FarEndIntervalCESs"] = dsx3Farendintervalentry.Dsx3Farendintervalcess
    leafs["dsx3FarEndIntervalCSESs"] = dsx3Farendintervalentry.Dsx3Farendintervalcsess
    leafs["dsx3FarEndIntervalCCVs"] = dsx3Farendintervalentry.Dsx3Farendintervalccvs
    leafs["dsx3FarEndIntervalUASs"] = dsx3Farendintervalentry.Dsx3Farendintervaluass
    leafs["dsx3FarEndIntervalValidData"] = dsx3Farendintervalentry.Dsx3Farendintervalvaliddata
    return leafs
}

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetYangName() string { return "dsx3FarEndIntervalEntry" }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) SetParent(parent types.Entity) { dsx3Farendintervalentry.parent = parent }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetParent() types.Entity { return dsx3Farendintervalentry.parent }

func (dsx3Farendintervalentry *DS3MIB_Dsx3Farendintervaltable_Dsx3Farendintervalentry) GetParentYangName() string { return "dsx3FarEndIntervalTable" }

// DS3MIB_Dsx3Farendtotaltable
// The DS3 Far End Total Table contains the
// cumulative sum of the various statistics for the
// 24 hour period preceding the current interval.
type DS3MIB_Dsx3Farendtotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3 Far End Total table. The type is slice of
    // DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry.
    Dsx3Farendtotalentry []DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetFilter() yfilter.YFilter { return dsx3Farendtotaltable.YFilter }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) SetFilter(yf yfilter.YFilter) { dsx3Farendtotaltable.YFilter = yf }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetGoName(yname string) string {
    if yname == "dsx3FarEndTotalEntry" { return "Dsx3Farendtotalentry" }
    return ""
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetSegmentPath() string {
    return "dsx3FarEndTotalTable"
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3FarEndTotalEntry" {
        for _, c := range dsx3Farendtotaltable.Dsx3Farendtotalentry {
            if dsx3Farendtotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry{}
        dsx3Farendtotaltable.Dsx3Farendtotalentry = append(dsx3Farendtotaltable.Dsx3Farendtotalentry, child)
        return &dsx3Farendtotaltable.Dsx3Farendtotalentry[len(dsx3Farendtotaltable.Dsx3Farendtotalentry)-1]
    }
    return nil
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Farendtotaltable.Dsx3Farendtotalentry {
        children[dsx3Farendtotaltable.Dsx3Farendtotalentry[i].GetSegmentPath()] = &dsx3Farendtotaltable.Dsx3Farendtotalentry[i]
    }
    return children
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetYangName() string { return "dsx3FarEndTotalTable" }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) SetParent(parent types.Entity) { dsx3Farendtotaltable.parent = parent }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetParent() types.Entity { return dsx3Farendtotaltable.parent }

func (dsx3Farendtotaltable *DS3MIB_Dsx3Farendtotaltable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry
// An entry in the DS3 Far End Total table.
type DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry struct {
    parent types.Entity
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

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetFilter() yfilter.YFilter { return dsx3Farendtotalentry.YFilter }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) SetFilter(yf yfilter.YFilter) { dsx3Farendtotalentry.YFilter = yf }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetGoName(yname string) string {
    if yname == "dsx3FarEndTotalIndex" { return "Dsx3Farendtotalindex" }
    if yname == "dsx3FarEndTotalCESs" { return "Dsx3Farendtotalcess" }
    if yname == "dsx3FarEndTotalCSESs" { return "Dsx3Farendtotalcsess" }
    if yname == "dsx3FarEndTotalCCVs" { return "Dsx3Farendtotalccvs" }
    if yname == "dsx3FarEndTotalUASs" { return "Dsx3Farendtotaluass" }
    return ""
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetSegmentPath() string {
    return "dsx3FarEndTotalEntry" + "[dsx3FarEndTotalIndex='" + fmt.Sprintf("%v", dsx3Farendtotalentry.Dsx3Farendtotalindex) + "']"
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3FarEndTotalIndex"] = dsx3Farendtotalentry.Dsx3Farendtotalindex
    leafs["dsx3FarEndTotalCESs"] = dsx3Farendtotalentry.Dsx3Farendtotalcess
    leafs["dsx3FarEndTotalCSESs"] = dsx3Farendtotalentry.Dsx3Farendtotalcsess
    leafs["dsx3FarEndTotalCCVs"] = dsx3Farendtotalentry.Dsx3Farendtotalccvs
    leafs["dsx3FarEndTotalUASs"] = dsx3Farendtotalentry.Dsx3Farendtotaluass
    return leafs
}

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetYangName() string { return "dsx3FarEndTotalEntry" }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) SetParent(parent types.Entity) { dsx3Farendtotalentry.parent = parent }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetParent() types.Entity { return dsx3Farendtotalentry.parent }

func (dsx3Farendtotalentry *DS3MIB_Dsx3Farendtotaltable_Dsx3Farendtotalentry) GetParentYangName() string { return "dsx3FarEndTotalTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DS3 Fractional table. The type is slice of
    // DS3MIB_Dsx3Fractable_Dsx3Fracentry.
    Dsx3Fracentry []DS3MIB_Dsx3Fractable_Dsx3Fracentry
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetFilter() yfilter.YFilter { return dsx3Fractable.YFilter }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) SetFilter(yf yfilter.YFilter) { dsx3Fractable.YFilter = yf }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetGoName(yname string) string {
    if yname == "dsx3FracEntry" { return "Dsx3Fracentry" }
    return ""
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetSegmentPath() string {
    return "dsx3FracTable"
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dsx3FracEntry" {
        for _, c := range dsx3Fractable.Dsx3Fracentry {
            if dsx3Fractable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DS3MIB_Dsx3Fractable_Dsx3Fracentry{}
        dsx3Fractable.Dsx3Fracentry = append(dsx3Fractable.Dsx3Fracentry, child)
        return &dsx3Fractable.Dsx3Fracentry[len(dsx3Fractable.Dsx3Fracentry)-1]
    }
    return nil
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dsx3Fractable.Dsx3Fracentry {
        children[dsx3Fractable.Dsx3Fracentry[i].GetSegmentPath()] = &dsx3Fractable.Dsx3Fracentry[i]
    }
    return children
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetYangName() string { return "dsx3FracTable" }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) SetParent(parent types.Entity) { dsx3Fractable.parent = parent }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetParent() types.Entity { return dsx3Fractable.parent }

func (dsx3Fractable *DS3MIB_Dsx3Fractable) GetParentYangName() string { return "DS3-MIB" }

// DS3MIB_Dsx3Fractable_Dsx3Fracentry
// An entry in the DS3 Fractional table.
type DS3MIB_Dsx3Fractable_Dsx3Fracentry struct {
    parent types.Entity
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

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetFilter() yfilter.YFilter { return dsx3Fracentry.YFilter }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) SetFilter(yf yfilter.YFilter) { dsx3Fracentry.YFilter = yf }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetGoName(yname string) string {
    if yname == "dsx3FracIndex" { return "Dsx3Fracindex" }
    if yname == "dsx3FracNumber" { return "Dsx3Fracnumber" }
    if yname == "dsx3FracIfIndex" { return "Dsx3Fracifindex" }
    return ""
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetSegmentPath() string {
    return "dsx3FracEntry" + "[dsx3FracIndex='" + fmt.Sprintf("%v", dsx3Fracentry.Dsx3Fracindex) + "']" + "[dsx3FracNumber='" + fmt.Sprintf("%v", dsx3Fracentry.Dsx3Fracnumber) + "']"
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dsx3FracIndex"] = dsx3Fracentry.Dsx3Fracindex
    leafs["dsx3FracNumber"] = dsx3Fracentry.Dsx3Fracnumber
    leafs["dsx3FracIfIndex"] = dsx3Fracentry.Dsx3Fracifindex
    return leafs
}

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetBundleName() string { return "cisco_ios_xe" }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetYangName() string { return "dsx3FracEntry" }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) SetParent(parent types.Entity) { dsx3Fracentry.parent = parent }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetParent() types.Entity { return dsx3Fracentry.parent }

func (dsx3Fracentry *DS3MIB_Dsx3Fractable_Dsx3Fracentry) GetParentYangName() string { return "dsx3FracTable" }

