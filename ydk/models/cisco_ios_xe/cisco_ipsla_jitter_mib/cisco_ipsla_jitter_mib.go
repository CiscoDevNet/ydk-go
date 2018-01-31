// This MIB module defines templates for IP SLA operations of UDP
// Jitter and ICMP Jitter. 
// 
// The UDP Jitter operation is designed to measure the delay 
// variance and packet loss in IP networks by generating synthetic 
// UDP traffic. 
// 
// The ICMP Jitter operation provides capability to measure metrics 
// such as RTT (Round Trip Time), jitter, packet loss, one-way 
// latency by sending ICMP Timestamp stream to the destination 
// devices.
package cisco_ipsla_jitter_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsla_jitter_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSLA-JITTER-MIB CISCO-IPSLA-JITTER-MIB}", reflect.TypeOf(CISCOIPSLAJITTERMIB{}))
    ydk.RegisterEntity("CISCO-IPSLA-JITTER-MIB:CISCO-IPSLA-JITTER-MIB", reflect.TypeOf(CISCOIPSLAJITTERMIB{}))
}

// CISCOIPSLAJITTERMIB
type CISCOIPSLAJITTERMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table that contains UDP jitter template specific definitions.
    Cipslaudpjittertmpltable CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable

    // A table that contains ICMP jitter template specific definitions.
    Cipslaicmpjittertmpltable CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetFilter() yfilter.YFilter { return cISCOIPSLAJITTERMIB.YFilter }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSLAJITTERMIB.YFilter = yf }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetGoName(yname string) string {
    if yname == "cipslaUdpJitterTmplTable" { return "Cipslaudpjittertmpltable" }
    if yname == "cipslaIcmpJitterTmplTable" { return "Cipslaicmpjittertmpltable" }
    return ""
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetSegmentPath() string {
    return "CISCO-IPSLA-JITTER-MIB:CISCO-IPSLA-JITTER-MIB"
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaUdpJitterTmplTable" {
        return &cISCOIPSLAJITTERMIB.Cipslaudpjittertmpltable
    }
    if childYangName == "cipslaIcmpJitterTmplTable" {
        return &cISCOIPSLAJITTERMIB.Cipslaicmpjittertmpltable
    }
    return nil
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipslaUdpJitterTmplTable"] = &cISCOIPSLAJITTERMIB.Cipslaudpjittertmpltable
    children["cipslaIcmpJitterTmplTable"] = &cISCOIPSLAJITTERMIB.Cipslaicmpjittertmpltable
    return children
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetYangName() string { return "CISCO-IPSLA-JITTER-MIB" }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) SetParent(parent types.Entity) { cISCOIPSLAJITTERMIB.parent = parent }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetParent() types.Entity { return cISCOIPSLAJITTERMIB.parent }

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetParentYangName() string { return "CISCO-IPSLA-JITTER-MIB" }

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable
// A table that contains UDP jitter template specific definitions.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP jitter template. The type is slice of
    // CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry.
    Cipslaudpjittertmplentry []CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetFilter() yfilter.YFilter { return cipslaudpjittertmpltable.YFilter }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) SetFilter(yf yfilter.YFilter) { cipslaudpjittertmpltable.YFilter = yf }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetGoName(yname string) string {
    if yname == "cipslaUdpJitterTmplEntry" { return "Cipslaudpjittertmplentry" }
    return ""
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetSegmentPath() string {
    return "cipslaUdpJitterTmplTable"
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaUdpJitterTmplEntry" {
        for _, c := range cipslaudpjittertmpltable.Cipslaudpjittertmplentry {
            if cipslaudpjittertmpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry{}
        cipslaudpjittertmpltable.Cipslaudpjittertmplentry = append(cipslaudpjittertmpltable.Cipslaudpjittertmplentry, child)
        return &cipslaudpjittertmpltable.Cipslaudpjittertmplentry[len(cipslaudpjittertmpltable.Cipslaudpjittertmplentry)-1]
    }
    return nil
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaudpjittertmpltable.Cipslaudpjittertmplentry {
        children[cipslaudpjittertmpltable.Cipslaudpjittertmplentry[i].GetSegmentPath()] = &cipslaudpjittertmpltable.Cipslaudpjittertmplentry[i]
    }
    return children
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetYangName() string { return "cipslaUdpJitterTmplTable" }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) SetParent(parent types.Entity) { cipslaudpjittertmpltable.parent = parent }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetParent() types.Entity { return cipslaudpjittertmpltable.parent }

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetParentYangName() string { return "CISCO-IPSLA-JITTER-MIB" }

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry
// A row entry representing an IPSLA UDP jitter template.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the UDP Jitter template
    // name. The type is string with length: 1..64.
    Cipslaudpjittertmplname interface{}

    // A string which provides description of UDP Jitter template. The type is
    // string with length: 0..128.
    Cipslaudpjittertmpldescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    Cipslaudpjittertmplcontrolenable interface{}

    // Specifies the duration to wait for a IP SLA operation completion.   For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    Cipslaudpjittertmpltimeout interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    Cipslaudpjittertmplverifydata interface{}

    // Specifies the codec type to be used with UDP jitter operation.  If
    // codec-type is configured the following parameters cannot be  configured.
    // cipslaUdpJitterReqDataSize cipslaUdpJitterInterval cipslaUdpJitterNumPkts.
    // The type is IpSlaCodecType.
    Cipslaudpjittertmplcodectype interface{}

    // This field represents the inter-packet delay between packets and is in
    // milliseconds. This object is applicable only to UDP jitter operation  which
    // uses codec type. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    Cipslaudpjittertmplcodecinterval interface{}

    // This object represents the number of octets that needs to be placed into
    // the Data portion of the message. This value is used only for UDP jitter
    // operation  which uses codec type. The type is interface{} with range:
    // 0..16384. Units are octets.
    Cipslaudpjittertmplcodecpayload interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is used only for UDP jitter  operation which uses codec type.
    // The type is interface{} with range: 1..60000. Units are packets.
    Cipslaudpjittertmplcodecnumpkts interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    Cipslaudpjittertmplinterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // The type is interface{} with range: 1..60000. Units are packets.
    Cipslaudpjittertmplnumpkts interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaUdpJitterTmplSrcAddr object. The type is
    // InetAddressType.
    Cipslaudpjittertmplsrcaddrtype interface{}

    // This field specifies the IP address of the source. The type is string with
    // length: 0..255.
    Cipslaudpjittertmplsrcaddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    Cipslaudpjittertmplsrcport interface{}

    // This object specifies the accuracy of jitter statistics in
    // rttMonJitterStatsTable that needs to be calculated. milliseconds(1) - The
    // accuracy of stats will be of milliseconds. microseconds(2) - The accuracy
    // of stats will be in microseconds. The type is Cipslaudpjittertmplprecision.
    Cipslaudpjittertmplprecision interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' IP SLA request/responses, this value represents the native
    // payload size.  REMEMBER:  The ARR Header overhead is not included          
    // in this value. The type is interface{} with range: 16..65024. Units are
    // octets.
    Cipslaudpjittertmplreqdatasize interface{}

    // This object specifies the priority that will be assigned to operation
    // packet.  normal(1) - The packet is of normal priority. high(2)   - The
    // packet is of high priority. The type is Cipslaudpjittertmplpktpriority.
    Cipslaudpjittertmplpktpriority interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    Cipslaudpjittertmpltos interface{}

    // This field is used to specify the VRF name in which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VPN routing table
    // for this operation. The type is string with length: 0..32.
    Cipslaudpjittertmplvrfname interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit, then one threshold crossing occurrence
    // will be counted. The type is interface{} with range: 0..2147483647. Units
    // are milliseconds.
    Cipslaudpjittertmplthreshold interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered tolerable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of NTP offsets on source and responder. The value specified is
    // microseconds. This value can be set only for UDP jitter operation  with
    // precision of microsecond. The type is interface{} with range: 0..100000.
    // Units are microseconds.
    Cipslaudpjittertmplntptolabs interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered tolerable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of  NTP offsets on source and responder. The value is
    // expressed  as the percentage of actual oneway latency that is measured. 
    // This value can be set only for UDP jitter operation with precision  of
    // microsecond. The type is interface{} with range: 0..100.
    Cipslaudpjittertmplntptolpct interface{}

    // This object specifies whether the value specified for oneway NTP sync
    // tolerance is absolute value or percent value.  percent(1)  - The value for
    // oneway NTP sync tolerance is                absolute value. absolute(2) -
    // The value for oneway NTP sync tolerance is                percent value.
    // The type is Cipslaudpjittertmplntptoltype.
    Cipslaudpjittertmplntptoltype interface{}

    // The advantage factor is dependant on the type of access and how the service
    // is to be used. Conventional Wire-line     0 Mobility within Building    5
    // Mobility within geographic area  10 Access to hard-to-reach location   20 
    // It is used when calculating the ICPIF value. The type is interface{} with
    // range: 0..20.
    Cipslaudpjittertmplicpiffactor interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    Cipslaudpjittertmplstatshours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaUdpJitterTmplDistBuckets will be kept.  The last bucket will contain
    // all entries from its  distribution interval start point to infinity. The
    // type is interface{} with range: 1..20.
    Cipslaudpjittertmpldistbuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  cipslaUdpJitterTmplDistBuckets = 5 buckets
    // cipslaUdpJitterTmplDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2 |
    // Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms | 30-39
    // ms | 40-Inf ms |  Odd Example:  cipslaUdpJitterTmplDistBuckets = 1 buckets
    // cipslaUdpJitterTmplDistInterval = 10 milliseconds  | Bucket 1  | |  0-Inf
    // ms |  Thus, this odd example shows that the value of
    // cipslaUdpJitterTmplDistInterval does not apply when
    // cipslaUdpJitterTmplDistBuckets is one. The type is interface{} with range:
    // 1..100. Units are milliseconds.
    Cipslaudpjittertmpldistinterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslaudpjittertmplstoragetype interface{}

    // The status of the conceptual UDP Jitter template control row. When the
    // status is active, all the read-create objects in that  row can be modified.
    // The type is RowStatus.
    Cipslaudpjittertmplrowstatus interface{}
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetFilter() yfilter.YFilter { return cipslaudpjittertmplentry.YFilter }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) SetFilter(yf yfilter.YFilter) { cipslaudpjittertmplentry.YFilter = yf }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetGoName(yname string) string {
    if yname == "cipslaUdpJitterTmplName" { return "Cipslaudpjittertmplname" }
    if yname == "cipslaUdpJitterTmplDescription" { return "Cipslaudpjittertmpldescription" }
    if yname == "cipslaUdpJitterTmplControlEnable" { return "Cipslaudpjittertmplcontrolenable" }
    if yname == "cipslaUdpJitterTmplTimeOut" { return "Cipslaudpjittertmpltimeout" }
    if yname == "cipslaUdpJitterTmplVerifyData" { return "Cipslaudpjittertmplverifydata" }
    if yname == "cipslaUdpJitterTmplCodecType" { return "Cipslaudpjittertmplcodectype" }
    if yname == "cipslaUdpJitterTmplCodecInterval" { return "Cipslaudpjittertmplcodecinterval" }
    if yname == "cipslaUdpJitterTmplCodecPayload" { return "Cipslaudpjittertmplcodecpayload" }
    if yname == "cipslaUdpJitterTmplCodecNumPkts" { return "Cipslaudpjittertmplcodecnumpkts" }
    if yname == "cipslaUdpJitterTmplInterval" { return "Cipslaudpjittertmplinterval" }
    if yname == "cipslaUdpJitterTmplNumPkts" { return "Cipslaudpjittertmplnumpkts" }
    if yname == "cipslaUdpJitterTmplSrcAddrType" { return "Cipslaudpjittertmplsrcaddrtype" }
    if yname == "cipslaUdpJitterTmplSrcAddr" { return "Cipslaudpjittertmplsrcaddr" }
    if yname == "cipslaUdpJitterTmplSrcPort" { return "Cipslaudpjittertmplsrcport" }
    if yname == "cipslaUdpJitterTmplPrecision" { return "Cipslaudpjittertmplprecision" }
    if yname == "cipslaUdpJitterTmplReqDataSize" { return "Cipslaudpjittertmplreqdatasize" }
    if yname == "cipslaUdpJitterTmplPktPriority" { return "Cipslaudpjittertmplpktpriority" }
    if yname == "cipslaUdpJitterTmplTOS" { return "Cipslaudpjittertmpltos" }
    if yname == "cipslaUdpJitterTmplVrfName" { return "Cipslaudpjittertmplvrfname" }
    if yname == "cipslaUdpJitterTmplThreshold" { return "Cipslaudpjittertmplthreshold" }
    if yname == "cipslaUdpJitterTmplNTPTolAbs" { return "Cipslaudpjittertmplntptolabs" }
    if yname == "cipslaUdpJitterTmplNTPTolPct" { return "Cipslaudpjittertmplntptolpct" }
    if yname == "cipslaUdpJitterTmplNTPTolType" { return "Cipslaudpjittertmplntptoltype" }
    if yname == "cipslaUdpJitterTmplIcpifFactor" { return "Cipslaudpjittertmplicpiffactor" }
    if yname == "cipslaUdpJitterTmplStatsHours" { return "Cipslaudpjittertmplstatshours" }
    if yname == "cipslaUdpJitterTmplDistBuckets" { return "Cipslaudpjittertmpldistbuckets" }
    if yname == "cipslaUdpJitterTmplDistInterval" { return "Cipslaudpjittertmpldistinterval" }
    if yname == "cipslaUdpJitterTmplStorageType" { return "Cipslaudpjittertmplstoragetype" }
    if yname == "cipslaUdpJitterTmplRowStatus" { return "Cipslaudpjittertmplrowstatus" }
    return ""
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetSegmentPath() string {
    return "cipslaUdpJitterTmplEntry" + "[cipslaUdpJitterTmplName='" + fmt.Sprintf("%v", cipslaudpjittertmplentry.Cipslaudpjittertmplname) + "']"
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaUdpJitterTmplName"] = cipslaudpjittertmplentry.Cipslaudpjittertmplname
    leafs["cipslaUdpJitterTmplDescription"] = cipslaudpjittertmplentry.Cipslaudpjittertmpldescription
    leafs["cipslaUdpJitterTmplControlEnable"] = cipslaudpjittertmplentry.Cipslaudpjittertmplcontrolenable
    leafs["cipslaUdpJitterTmplTimeOut"] = cipslaudpjittertmplentry.Cipslaudpjittertmpltimeout
    leafs["cipslaUdpJitterTmplVerifyData"] = cipslaudpjittertmplentry.Cipslaudpjittertmplverifydata
    leafs["cipslaUdpJitterTmplCodecType"] = cipslaudpjittertmplentry.Cipslaudpjittertmplcodectype
    leafs["cipslaUdpJitterTmplCodecInterval"] = cipslaudpjittertmplentry.Cipslaudpjittertmplcodecinterval
    leafs["cipslaUdpJitterTmplCodecPayload"] = cipslaudpjittertmplentry.Cipslaudpjittertmplcodecpayload
    leafs["cipslaUdpJitterTmplCodecNumPkts"] = cipslaudpjittertmplentry.Cipslaudpjittertmplcodecnumpkts
    leafs["cipslaUdpJitterTmplInterval"] = cipslaudpjittertmplentry.Cipslaudpjittertmplinterval
    leafs["cipslaUdpJitterTmplNumPkts"] = cipslaudpjittertmplentry.Cipslaudpjittertmplnumpkts
    leafs["cipslaUdpJitterTmplSrcAddrType"] = cipslaudpjittertmplentry.Cipslaudpjittertmplsrcaddrtype
    leafs["cipslaUdpJitterTmplSrcAddr"] = cipslaudpjittertmplentry.Cipslaudpjittertmplsrcaddr
    leafs["cipslaUdpJitterTmplSrcPort"] = cipslaudpjittertmplentry.Cipslaudpjittertmplsrcport
    leafs["cipslaUdpJitterTmplPrecision"] = cipslaudpjittertmplentry.Cipslaudpjittertmplprecision
    leafs["cipslaUdpJitterTmplReqDataSize"] = cipslaudpjittertmplentry.Cipslaudpjittertmplreqdatasize
    leafs["cipslaUdpJitterTmplPktPriority"] = cipslaudpjittertmplentry.Cipslaudpjittertmplpktpriority
    leafs["cipslaUdpJitterTmplTOS"] = cipslaudpjittertmplentry.Cipslaudpjittertmpltos
    leafs["cipslaUdpJitterTmplVrfName"] = cipslaudpjittertmplentry.Cipslaudpjittertmplvrfname
    leafs["cipslaUdpJitterTmplThreshold"] = cipslaudpjittertmplentry.Cipslaudpjittertmplthreshold
    leafs["cipslaUdpJitterTmplNTPTolAbs"] = cipslaudpjittertmplentry.Cipslaudpjittertmplntptolabs
    leafs["cipslaUdpJitterTmplNTPTolPct"] = cipslaudpjittertmplentry.Cipslaudpjittertmplntptolpct
    leafs["cipslaUdpJitterTmplNTPTolType"] = cipslaudpjittertmplentry.Cipslaudpjittertmplntptoltype
    leafs["cipslaUdpJitterTmplIcpifFactor"] = cipslaudpjittertmplentry.Cipslaudpjittertmplicpiffactor
    leafs["cipslaUdpJitterTmplStatsHours"] = cipslaudpjittertmplentry.Cipslaudpjittertmplstatshours
    leafs["cipslaUdpJitterTmplDistBuckets"] = cipslaudpjittertmplentry.Cipslaudpjittertmpldistbuckets
    leafs["cipslaUdpJitterTmplDistInterval"] = cipslaudpjittertmplentry.Cipslaudpjittertmpldistinterval
    leafs["cipslaUdpJitterTmplStorageType"] = cipslaudpjittertmplentry.Cipslaudpjittertmplstoragetype
    leafs["cipslaUdpJitterTmplRowStatus"] = cipslaudpjittertmplentry.Cipslaudpjittertmplrowstatus
    return leafs
}

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetYangName() string { return "cipslaUdpJitterTmplEntry" }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) SetParent(parent types.Entity) { cipslaudpjittertmplentry.parent = parent }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetParent() types.Entity { return cipslaudpjittertmplentry.parent }

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetParentYangName() string { return "cipslaUdpJitterTmplTable" }

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype represents               percent value.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype string

const (
    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype_percent CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype = "percent"

    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype_absolute CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplntptoltype = "absolute"
)

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority represents high(2)   - The packet is of high priority.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority string

const (
    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority_normal CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority = "normal"

    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority_high CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplpktpriority = "high"
)

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision represents microseconds(2) - The accuracy of stats will be in microseconds.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision string

const (
    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision_milliseconds CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision = "milliseconds"

    CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision_microseconds CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry_Cipslaudpjittertmplprecision = "microseconds"
)

// CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable
// A table that contains ICMP jitter template specific definitions.
type CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row entry representing an IP SLA ICMP Jitter template. The type is slice
    // of CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry.
    Cipslaicmpjittertmplentry []CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetFilter() yfilter.YFilter { return cipslaicmpjittertmpltable.YFilter }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) SetFilter(yf yfilter.YFilter) { cipslaicmpjittertmpltable.YFilter = yf }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetGoName(yname string) string {
    if yname == "cipslaIcmpJitterTmplEntry" { return "Cipslaicmpjittertmplentry" }
    return ""
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetSegmentPath() string {
    return "cipslaIcmpJitterTmplTable"
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaIcmpJitterTmplEntry" {
        for _, c := range cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry {
            if cipslaicmpjittertmpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry{}
        cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry = append(cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry, child)
        return &cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry[len(cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry)-1]
    }
    return nil
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry {
        children[cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry[i].GetSegmentPath()] = &cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry[i]
    }
    return children
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetYangName() string { return "cipslaIcmpJitterTmplTable" }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) SetParent(parent types.Entity) { cipslaicmpjittertmpltable.parent = parent }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetParent() types.Entity { return cipslaicmpjittertmpltable.parent }

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetParentYangName() string { return "CISCO-IPSLA-JITTER-MIB" }

// CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry
// A row entry representing an IP SLA ICMP Jitter template.
type CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the ICMP jitter template
    // name. The type is string with length: 1..64.
    Cipslaicmpjittertmplname interface{}

    // A string which provides description of ICMP Jitter template. The type is
    // string with length: 0..128.
    Cipslaicmpjittertmpldescription interface{}

    // Specifies the duration to wait for a IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    Cipslaicmpjittertmpltimeout interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    Cipslaicmpjittertmplverifydata interface{}

    // This value represents the number of packets that need to be transmitted.
    // The type is interface{} with range: 1..60000. Units are packets.
    Cipslaicmpjittertmplnumpkts interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    Cipslaicmpjittertmplinterval interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaIcmpJitterTmplSrcAddr object. The type is
    // InetAddressType.
    Cipslaicmpjittertmplsrcaddrtype interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    Cipslaicmpjittertmplsrcaddr interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    Cipslaicmpjittertmpltos interface{}

    // This field is used to specify the VRF name in which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VPN routing Table
    // for this operation. The type is string with length: 0..32.
    Cipslaicmpjittertmplvrfname interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit, then  one threshold crossing occurrence
    // will be counted. The type is interface{} with range: 0..2147483647. Units
    // are milliseconds.
    Cipslaicmpjittertmplthreshold interface{}

    // The maximum number of hourss for which statistics are maintained.
    // Specifically this is the number of hourly groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    Cipslaicmpjittertmplstatshours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaIcmpJitterTmplDistBuckets will be kept.  The last bucket will contain
    // all entries from its  distribution interval start point to infinity. The
    // type is interface{} with range: 1..20.
    Cipslaicmpjittertmpldistbuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  cipslaIcmpJitterTmplDistBuckets = 5 buckets
    // cipslaIcmpJitterTmplDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2 |
    // Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms | 30-39
    // ms | 40-Inf ms |  Odd Example:  cipslaIcmpJitterTmplDistBuckets = 1 buckets
    // cipslaIcmpJitterTmplDistInterval = 10 milliseconds  | Bucket 1  | |  0-Inf
    // ms |  Thus, this odd example shows that the value of
    // cipslaIcmpJitterTmplDistInterval does not apply when
    // cipslaIcmpJitterTmplDistBuckets is one. The type is interface{} with range:
    // 1..100. Units are milliseconds.
    Cipslaicmpjittertmpldistinterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslaicmpjittertmplstoragetype interface{}

    // The status of the conceptual ICMP jitter template control row. When the
    // status is active, all the read-create objects in  that row can be modified.
    // The type is RowStatus.
    Cipslaicmpjittertmplrowstatus interface{}
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetFilter() yfilter.YFilter { return cipslaicmpjittertmplentry.YFilter }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) SetFilter(yf yfilter.YFilter) { cipslaicmpjittertmplentry.YFilter = yf }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetGoName(yname string) string {
    if yname == "cipslaIcmpJitterTmplName" { return "Cipslaicmpjittertmplname" }
    if yname == "cipslaIcmpJitterTmplDescription" { return "Cipslaicmpjittertmpldescription" }
    if yname == "cipslaIcmpJitterTmplTimeOut" { return "Cipslaicmpjittertmpltimeout" }
    if yname == "cipslaIcmpJitterTmplVerifyData" { return "Cipslaicmpjittertmplverifydata" }
    if yname == "cipslaIcmpJitterTmplNumPkts" { return "Cipslaicmpjittertmplnumpkts" }
    if yname == "cipslaIcmpJitterTmplInterval" { return "Cipslaicmpjittertmplinterval" }
    if yname == "cipslaIcmpJitterTmplSrcAddrType" { return "Cipslaicmpjittertmplsrcaddrtype" }
    if yname == "cipslaIcmpJitterTmplSrcAddr" { return "Cipslaicmpjittertmplsrcaddr" }
    if yname == "cipslaIcmpJitterTmplTOS" { return "Cipslaicmpjittertmpltos" }
    if yname == "cipslaIcmpJitterTmplVrfName" { return "Cipslaicmpjittertmplvrfname" }
    if yname == "cipslaIcmpJitterTmplThreshold" { return "Cipslaicmpjittertmplthreshold" }
    if yname == "cipslaIcmpJitterTmplStatsHours" { return "Cipslaicmpjittertmplstatshours" }
    if yname == "cipslaIcmpJitterTmplDistBuckets" { return "Cipslaicmpjittertmpldistbuckets" }
    if yname == "cipslaIcmpJitterTmplDistInterval" { return "Cipslaicmpjittertmpldistinterval" }
    if yname == "cipslaIcmpJitterTmplStorageType" { return "Cipslaicmpjittertmplstoragetype" }
    if yname == "cipslaIcmpJitterTmplRowStatus" { return "Cipslaicmpjittertmplrowstatus" }
    return ""
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetSegmentPath() string {
    return "cipslaIcmpJitterTmplEntry" + "[cipslaIcmpJitterTmplName='" + fmt.Sprintf("%v", cipslaicmpjittertmplentry.Cipslaicmpjittertmplname) + "']"
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaIcmpJitterTmplName"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplname
    leafs["cipslaIcmpJitterTmplDescription"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmpldescription
    leafs["cipslaIcmpJitterTmplTimeOut"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmpltimeout
    leafs["cipslaIcmpJitterTmplVerifyData"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplverifydata
    leafs["cipslaIcmpJitterTmplNumPkts"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplnumpkts
    leafs["cipslaIcmpJitterTmplInterval"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplinterval
    leafs["cipslaIcmpJitterTmplSrcAddrType"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplsrcaddrtype
    leafs["cipslaIcmpJitterTmplSrcAddr"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplsrcaddr
    leafs["cipslaIcmpJitterTmplTOS"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmpltos
    leafs["cipslaIcmpJitterTmplVrfName"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplvrfname
    leafs["cipslaIcmpJitterTmplThreshold"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplthreshold
    leafs["cipslaIcmpJitterTmplStatsHours"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplstatshours
    leafs["cipslaIcmpJitterTmplDistBuckets"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmpldistbuckets
    leafs["cipslaIcmpJitterTmplDistInterval"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmpldistinterval
    leafs["cipslaIcmpJitterTmplStorageType"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplstoragetype
    leafs["cipslaIcmpJitterTmplRowStatus"] = cipslaicmpjittertmplentry.Cipslaicmpjittertmplrowstatus
    return leafs
}

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetYangName() string { return "cipslaIcmpJitterTmplEntry" }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) SetParent(parent types.Entity) { cipslaicmpjittertmplentry.parent = parent }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetParent() types.Entity { return cipslaicmpjittertmplentry.parent }

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetParentYangName() string { return "cipslaIcmpJitterTmplTable" }

