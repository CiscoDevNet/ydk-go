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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table that contains UDP jitter template specific definitions.
    Cipslaudpjittertmpltable CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable

    // A table that contains ICMP jitter template specific definitions.
    Cipslaicmpjittertmpltable CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable
}

func (cISCOIPSLAJITTERMIB *CISCOIPSLAJITTERMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSLAJITTERMIB.EntityData.YFilter = cISCOIPSLAJITTERMIB.YFilter
    cISCOIPSLAJITTERMIB.EntityData.YangName = "CISCO-IPSLA-JITTER-MIB"
    cISCOIPSLAJITTERMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSLAJITTERMIB.EntityData.ParentYangName = "CISCO-IPSLA-JITTER-MIB"
    cISCOIPSLAJITTERMIB.EntityData.SegmentPath = "CISCO-IPSLA-JITTER-MIB:CISCO-IPSLA-JITTER-MIB"
    cISCOIPSLAJITTERMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSLAJITTERMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSLAJITTERMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSLAJITTERMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSLAJITTERMIB.EntityData.Children["cipslaUdpJitterTmplTable"] = types.YChild{"Cipslaudpjittertmpltable", &cISCOIPSLAJITTERMIB.Cipslaudpjittertmpltable}
    cISCOIPSLAJITTERMIB.EntityData.Children["cipslaIcmpJitterTmplTable"] = types.YChild{"Cipslaicmpjittertmpltable", &cISCOIPSLAJITTERMIB.Cipslaicmpjittertmpltable}
    cISCOIPSLAJITTERMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSLAJITTERMIB.EntityData)
}

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable
// A table that contains UDP jitter template specific definitions.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP jitter template. The type is slice of
    // CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry.
    Cipslaudpjittertmplentry []CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry
}

func (cipslaudpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable) GetEntityData() *types.CommonEntityData {
    cipslaudpjittertmpltable.EntityData.YFilter = cipslaudpjittertmpltable.YFilter
    cipslaudpjittertmpltable.EntityData.YangName = "cipslaUdpJitterTmplTable"
    cipslaudpjittertmpltable.EntityData.BundleName = "cisco_ios_xe"
    cipslaudpjittertmpltable.EntityData.ParentYangName = "CISCO-IPSLA-JITTER-MIB"
    cipslaudpjittertmpltable.EntityData.SegmentPath = "cipslaUdpJitterTmplTable"
    cipslaudpjittertmpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaudpjittertmpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaudpjittertmpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaudpjittertmpltable.EntityData.Children = make(map[string]types.YChild)
    cipslaudpjittertmpltable.EntityData.Children["cipslaUdpJitterTmplEntry"] = types.YChild{"Cipslaudpjittertmplentry", nil}
    for i := range cipslaudpjittertmpltable.Cipslaudpjittertmplentry {
        cipslaudpjittertmpltable.EntityData.Children[types.GetSegmentPath(&cipslaudpjittertmpltable.Cipslaudpjittertmplentry[i])] = types.YChild{"Cipslaudpjittertmplentry", &cipslaudpjittertmpltable.Cipslaudpjittertmplentry[i]}
    }
    cipslaudpjittertmpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaudpjittertmpltable.EntityData)
}

// CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry
// A row entry representing an IPSLA UDP jitter template.
type CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry struct {
    EntityData types.CommonEntityData
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

func (cipslaudpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaudpjittertmpltable_Cipslaudpjittertmplentry) GetEntityData() *types.CommonEntityData {
    cipslaudpjittertmplentry.EntityData.YFilter = cipslaudpjittertmplentry.YFilter
    cipslaudpjittertmplentry.EntityData.YangName = "cipslaUdpJitterTmplEntry"
    cipslaudpjittertmplentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaudpjittertmplentry.EntityData.ParentYangName = "cipslaUdpJitterTmplTable"
    cipslaudpjittertmplentry.EntityData.SegmentPath = "cipslaUdpJitterTmplEntry" + "[cipslaUdpJitterTmplName='" + fmt.Sprintf("%v", cipslaudpjittertmplentry.Cipslaudpjittertmplname) + "']"
    cipslaudpjittertmplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaudpjittertmplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaudpjittertmplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaudpjittertmplentry.EntityData.Children = make(map[string]types.YChild)
    cipslaudpjittertmplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplName"] = types.YLeaf{"Cipslaudpjittertmplname", cipslaudpjittertmplentry.Cipslaudpjittertmplname}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplDescription"] = types.YLeaf{"Cipslaudpjittertmpldescription", cipslaudpjittertmplentry.Cipslaudpjittertmpldescription}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplControlEnable"] = types.YLeaf{"Cipslaudpjittertmplcontrolenable", cipslaudpjittertmplentry.Cipslaudpjittertmplcontrolenable}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplTimeOut"] = types.YLeaf{"Cipslaudpjittertmpltimeout", cipslaudpjittertmplentry.Cipslaudpjittertmpltimeout}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplVerifyData"] = types.YLeaf{"Cipslaudpjittertmplverifydata", cipslaudpjittertmplentry.Cipslaudpjittertmplverifydata}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplCodecType"] = types.YLeaf{"Cipslaudpjittertmplcodectype", cipslaudpjittertmplentry.Cipslaudpjittertmplcodectype}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplCodecInterval"] = types.YLeaf{"Cipslaudpjittertmplcodecinterval", cipslaudpjittertmplentry.Cipslaudpjittertmplcodecinterval}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplCodecPayload"] = types.YLeaf{"Cipslaudpjittertmplcodecpayload", cipslaudpjittertmplentry.Cipslaudpjittertmplcodecpayload}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplCodecNumPkts"] = types.YLeaf{"Cipslaudpjittertmplcodecnumpkts", cipslaudpjittertmplentry.Cipslaudpjittertmplcodecnumpkts}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplInterval"] = types.YLeaf{"Cipslaudpjittertmplinterval", cipslaudpjittertmplentry.Cipslaudpjittertmplinterval}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplNumPkts"] = types.YLeaf{"Cipslaudpjittertmplnumpkts", cipslaudpjittertmplentry.Cipslaudpjittertmplnumpkts}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplSrcAddrType"] = types.YLeaf{"Cipslaudpjittertmplsrcaddrtype", cipslaudpjittertmplentry.Cipslaudpjittertmplsrcaddrtype}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplSrcAddr"] = types.YLeaf{"Cipslaudpjittertmplsrcaddr", cipslaudpjittertmplentry.Cipslaudpjittertmplsrcaddr}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplSrcPort"] = types.YLeaf{"Cipslaudpjittertmplsrcport", cipslaudpjittertmplentry.Cipslaudpjittertmplsrcport}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplPrecision"] = types.YLeaf{"Cipslaudpjittertmplprecision", cipslaudpjittertmplentry.Cipslaudpjittertmplprecision}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplReqDataSize"] = types.YLeaf{"Cipslaudpjittertmplreqdatasize", cipslaudpjittertmplentry.Cipslaudpjittertmplreqdatasize}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplPktPriority"] = types.YLeaf{"Cipslaudpjittertmplpktpriority", cipslaudpjittertmplentry.Cipslaudpjittertmplpktpriority}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplTOS"] = types.YLeaf{"Cipslaudpjittertmpltos", cipslaudpjittertmplentry.Cipslaudpjittertmpltos}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplVrfName"] = types.YLeaf{"Cipslaudpjittertmplvrfname", cipslaudpjittertmplentry.Cipslaudpjittertmplvrfname}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplThreshold"] = types.YLeaf{"Cipslaudpjittertmplthreshold", cipslaudpjittertmplentry.Cipslaudpjittertmplthreshold}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplNTPTolAbs"] = types.YLeaf{"Cipslaudpjittertmplntptolabs", cipslaudpjittertmplentry.Cipslaudpjittertmplntptolabs}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplNTPTolPct"] = types.YLeaf{"Cipslaudpjittertmplntptolpct", cipslaudpjittertmplentry.Cipslaudpjittertmplntptolpct}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplNTPTolType"] = types.YLeaf{"Cipslaudpjittertmplntptoltype", cipslaudpjittertmplentry.Cipslaudpjittertmplntptoltype}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplIcpifFactor"] = types.YLeaf{"Cipslaudpjittertmplicpiffactor", cipslaudpjittertmplentry.Cipslaudpjittertmplicpiffactor}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplStatsHours"] = types.YLeaf{"Cipslaudpjittertmplstatshours", cipslaudpjittertmplentry.Cipslaudpjittertmplstatshours}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplDistBuckets"] = types.YLeaf{"Cipslaudpjittertmpldistbuckets", cipslaudpjittertmplentry.Cipslaudpjittertmpldistbuckets}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplDistInterval"] = types.YLeaf{"Cipslaudpjittertmpldistinterval", cipslaudpjittertmplentry.Cipslaudpjittertmpldistinterval}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplStorageType"] = types.YLeaf{"Cipslaudpjittertmplstoragetype", cipslaudpjittertmplentry.Cipslaudpjittertmplstoragetype}
    cipslaudpjittertmplentry.EntityData.Leafs["cipslaUdpJitterTmplRowStatus"] = types.YLeaf{"Cipslaudpjittertmplrowstatus", cipslaudpjittertmplentry.Cipslaudpjittertmplrowstatus}
    return &(cipslaudpjittertmplentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IP SLA ICMP Jitter template. The type is slice
    // of CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry.
    Cipslaicmpjittertmplentry []CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry
}

func (cipslaicmpjittertmpltable *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable) GetEntityData() *types.CommonEntityData {
    cipslaicmpjittertmpltable.EntityData.YFilter = cipslaicmpjittertmpltable.YFilter
    cipslaicmpjittertmpltable.EntityData.YangName = "cipslaIcmpJitterTmplTable"
    cipslaicmpjittertmpltable.EntityData.BundleName = "cisco_ios_xe"
    cipslaicmpjittertmpltable.EntityData.ParentYangName = "CISCO-IPSLA-JITTER-MIB"
    cipslaicmpjittertmpltable.EntityData.SegmentPath = "cipslaIcmpJitterTmplTable"
    cipslaicmpjittertmpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaicmpjittertmpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaicmpjittertmpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaicmpjittertmpltable.EntityData.Children = make(map[string]types.YChild)
    cipslaicmpjittertmpltable.EntityData.Children["cipslaIcmpJitterTmplEntry"] = types.YChild{"Cipslaicmpjittertmplentry", nil}
    for i := range cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry {
        cipslaicmpjittertmpltable.EntityData.Children[types.GetSegmentPath(&cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry[i])] = types.YChild{"Cipslaicmpjittertmplentry", &cipslaicmpjittertmpltable.Cipslaicmpjittertmplentry[i]}
    }
    cipslaicmpjittertmpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaicmpjittertmpltable.EntityData)
}

// CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry
// A row entry representing an IP SLA ICMP Jitter template.
type CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry struct {
    EntityData types.CommonEntityData
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

func (cipslaicmpjittertmplentry *CISCOIPSLAJITTERMIB_Cipslaicmpjittertmpltable_Cipslaicmpjittertmplentry) GetEntityData() *types.CommonEntityData {
    cipslaicmpjittertmplentry.EntityData.YFilter = cipslaicmpjittertmplentry.YFilter
    cipslaicmpjittertmplentry.EntityData.YangName = "cipslaIcmpJitterTmplEntry"
    cipslaicmpjittertmplentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaicmpjittertmplentry.EntityData.ParentYangName = "cipslaIcmpJitterTmplTable"
    cipslaicmpjittertmplentry.EntityData.SegmentPath = "cipslaIcmpJitterTmplEntry" + "[cipslaIcmpJitterTmplName='" + fmt.Sprintf("%v", cipslaicmpjittertmplentry.Cipslaicmpjittertmplname) + "']"
    cipslaicmpjittertmplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaicmpjittertmplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaicmpjittertmplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaicmpjittertmplentry.EntityData.Children = make(map[string]types.YChild)
    cipslaicmpjittertmplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplName"] = types.YLeaf{"Cipslaicmpjittertmplname", cipslaicmpjittertmplentry.Cipslaicmpjittertmplname}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplDescription"] = types.YLeaf{"Cipslaicmpjittertmpldescription", cipslaicmpjittertmplentry.Cipslaicmpjittertmpldescription}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplTimeOut"] = types.YLeaf{"Cipslaicmpjittertmpltimeout", cipslaicmpjittertmplentry.Cipslaicmpjittertmpltimeout}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplVerifyData"] = types.YLeaf{"Cipslaicmpjittertmplverifydata", cipslaicmpjittertmplentry.Cipslaicmpjittertmplverifydata}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplNumPkts"] = types.YLeaf{"Cipslaicmpjittertmplnumpkts", cipslaicmpjittertmplentry.Cipslaicmpjittertmplnumpkts}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplInterval"] = types.YLeaf{"Cipslaicmpjittertmplinterval", cipslaicmpjittertmplentry.Cipslaicmpjittertmplinterval}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplSrcAddrType"] = types.YLeaf{"Cipslaicmpjittertmplsrcaddrtype", cipslaicmpjittertmplentry.Cipslaicmpjittertmplsrcaddrtype}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplSrcAddr"] = types.YLeaf{"Cipslaicmpjittertmplsrcaddr", cipslaicmpjittertmplentry.Cipslaicmpjittertmplsrcaddr}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplTOS"] = types.YLeaf{"Cipslaicmpjittertmpltos", cipslaicmpjittertmplentry.Cipslaicmpjittertmpltos}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplVrfName"] = types.YLeaf{"Cipslaicmpjittertmplvrfname", cipslaicmpjittertmplentry.Cipslaicmpjittertmplvrfname}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplThreshold"] = types.YLeaf{"Cipslaicmpjittertmplthreshold", cipslaicmpjittertmplentry.Cipslaicmpjittertmplthreshold}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplStatsHours"] = types.YLeaf{"Cipslaicmpjittertmplstatshours", cipslaicmpjittertmplentry.Cipslaicmpjittertmplstatshours}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplDistBuckets"] = types.YLeaf{"Cipslaicmpjittertmpldistbuckets", cipslaicmpjittertmplentry.Cipslaicmpjittertmpldistbuckets}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplDistInterval"] = types.YLeaf{"Cipslaicmpjittertmpldistinterval", cipslaicmpjittertmplentry.Cipslaicmpjittertmpldistinterval}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplStorageType"] = types.YLeaf{"Cipslaicmpjittertmplstoragetype", cipslaicmpjittertmplentry.Cipslaicmpjittertmplstoragetype}
    cipslaicmpjittertmplentry.EntityData.Leafs["cipslaIcmpJitterTmplRowStatus"] = types.YLeaf{"Cipslaicmpjittertmplrowstatus", cipslaicmpjittertmplentry.Cipslaicmpjittertmplrowstatus}
    return &(cipslaicmpjittertmplentry.EntityData)
}

