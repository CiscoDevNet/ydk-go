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
    CipslaUdpJitterTmplTable CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable

    // A table that contains ICMP jitter template specific definitions.
    CipslaIcmpJitterTmplTable CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable
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

    cISCOIPSLAJITTERMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSLAJITTERMIB.EntityData.Children.Append("cipslaUdpJitterTmplTable", types.YChild{"CipslaUdpJitterTmplTable", &cISCOIPSLAJITTERMIB.CipslaUdpJitterTmplTable})
    cISCOIPSLAJITTERMIB.EntityData.Children.Append("cipslaIcmpJitterTmplTable", types.YChild{"CipslaIcmpJitterTmplTable", &cISCOIPSLAJITTERMIB.CipslaIcmpJitterTmplTable})
    cISCOIPSLAJITTERMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSLAJITTERMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSLAJITTERMIB.EntityData)
}

// CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable
// A table that contains UDP jitter template specific definitions.
type CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP jitter template. The type is slice of
    // CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry.
    CipslaUdpJitterTmplEntry []*CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry
}

func (cipslaUdpJitterTmplTable *CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable) GetEntityData() *types.CommonEntityData {
    cipslaUdpJitterTmplTable.EntityData.YFilter = cipslaUdpJitterTmplTable.YFilter
    cipslaUdpJitterTmplTable.EntityData.YangName = "cipslaUdpJitterTmplTable"
    cipslaUdpJitterTmplTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaUdpJitterTmplTable.EntityData.ParentYangName = "CISCO-IPSLA-JITTER-MIB"
    cipslaUdpJitterTmplTable.EntityData.SegmentPath = "cipslaUdpJitterTmplTable"
    cipslaUdpJitterTmplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaUdpJitterTmplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaUdpJitterTmplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaUdpJitterTmplTable.EntityData.Children = types.NewOrderedMap()
    cipslaUdpJitterTmplTable.EntityData.Children.Append("cipslaUdpJitterTmplEntry", types.YChild{"CipslaUdpJitterTmplEntry", nil})
    for i := range cipslaUdpJitterTmplTable.CipslaUdpJitterTmplEntry {
        cipslaUdpJitterTmplTable.EntityData.Children.Append(types.GetSegmentPath(cipslaUdpJitterTmplTable.CipslaUdpJitterTmplEntry[i]), types.YChild{"CipslaUdpJitterTmplEntry", cipslaUdpJitterTmplTable.CipslaUdpJitterTmplEntry[i]})
    }
    cipslaUdpJitterTmplTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaUdpJitterTmplTable.EntityData.YListKeys = []string {}

    return &(cipslaUdpJitterTmplTable.EntityData)
}

// CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry
// A row entry representing an IPSLA UDP jitter template.
type CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the UDP Jitter template
    // name. The type is string with length: 1..64.
    CipslaUdpJitterTmplName interface{}

    // A string which provides description of UDP Jitter template. The type is
    // string with length: 0..128.
    CipslaUdpJitterTmplDescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    CipslaUdpJitterTmplControlEnable interface{}

    // Specifies the duration to wait for a IP SLA operation completion.   For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    CipslaUdpJitterTmplTimeOut interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    CipslaUdpJitterTmplVerifyData interface{}

    // Specifies the codec type to be used with UDP jitter operation.  If
    // codec-type is configured the following parameters cannot be  configured.
    // cipslaUdpJitterReqDataSize cipslaUdpJitterInterval cipslaUdpJitterNumPkts.
    // The type is IpSlaCodecType.
    CipslaUdpJitterTmplCodecType interface{}

    // This field represents the inter-packet delay between packets and is in
    // milliseconds. This object is applicable only to UDP jitter operation  which
    // uses codec type. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    CipslaUdpJitterTmplCodecInterval interface{}

    // This object represents the number of octets that needs to be placed into
    // the Data portion of the message. This value is used only for UDP jitter
    // operation  which uses codec type. The type is interface{} with range:
    // 0..16384. Units are octets.
    CipslaUdpJitterTmplCodecPayload interface{}

    // This value represents the number of packets that need to be transmitted.
    // This value is used only for UDP jitter  operation which uses codec type.
    // The type is interface{} with range: 1..60000. Units are packets.
    CipslaUdpJitterTmplCodecNumPkts interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    CipslaUdpJitterTmplInterval interface{}

    // This value represents the number of packets that need to be transmitted.
    // The type is interface{} with range: 1..60000. Units are packets.
    CipslaUdpJitterTmplNumPkts interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaUdpJitterTmplSrcAddr object. The type is
    // InetAddressType.
    CipslaUdpJitterTmplSrcAddrType interface{}

    // This field specifies the IP address of the source. The type is string with
    // length: 0..255.
    CipslaUdpJitterTmplSrcAddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    CipslaUdpJitterTmplSrcPort interface{}

    // This object specifies the accuracy of jitter statistics in
    // rttMonJitterStatsTable that needs to be calculated. milliseconds(1) - The
    // accuracy of stats will be of milliseconds. microseconds(2) - The accuracy
    // of stats will be in microseconds. The type is CipslaUdpJitterTmplPrecision.
    CipslaUdpJitterTmplPrecision interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' IP SLA request/responses, this value represents the native
    // payload size.  REMEMBER:  The ARR Header overhead is not included          
    // in this value. The type is interface{} with range: 16..65024. Units are
    // octets.
    CipslaUdpJitterTmplReqDataSize interface{}

    // This object specifies the priority that will be assigned to operation
    // packet.  normal(1) - The packet is of normal priority. high(2)   - The
    // packet is of high priority. The type is CipslaUdpJitterTmplPktPriority.
    CipslaUdpJitterTmplPktPriority interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    CipslaUdpJitterTmplTOS interface{}

    // This field is used to specify the VRF name in which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VPN routing table
    // for this operation. The type is string with length: 0..32.
    CipslaUdpJitterTmplVrfName interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit, then one threshold crossing occurrence
    // will be counted. The type is interface{} with range: 0..2147483647. Units
    // are milliseconds.
    CipslaUdpJitterTmplThreshold interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered tolerable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of NTP offsets on source and responder. The value specified is
    // microseconds. This value can be set only for UDP jitter operation  with
    // precision of microsecond. The type is interface{} with range: 0..100000.
    // Units are microseconds.
    CipslaUdpJitterTmplNTPTolAbs interface{}

    // This object specifies the total clock synchronization error on source and
    // responder that is considered tolerable for  oneway measurement when NTP is
    // used as clock synchronization  mechanism.  The total clock synchronization
    // error is sum of  NTP offsets on source and responder. The value is
    // expressed  as the percentage of actual oneway latency that is measured. 
    // This value can be set only for UDP jitter operation with precision  of
    // microsecond. The type is interface{} with range: 0..100.
    CipslaUdpJitterTmplNTPTolPct interface{}

    // This object specifies whether the value specified for oneway NTP sync
    // tolerance is absolute value or percent value.  percent(1)  - The value for
    // oneway NTP sync tolerance is                absolute value. absolute(2) -
    // The value for oneway NTP sync tolerance is                percent value.
    // The type is CipslaUdpJitterTmplNTPTolType.
    CipslaUdpJitterTmplNTPTolType interface{}

    // The advantage factor is dependant on the type of access and how the service
    // is to be used. Conventional Wire-line     0 Mobility within Building    5
    // Mobility within geographic area  10 Access to hard-to-reach location   20 
    // It is used when calculating the ICPIF value. The type is interface{} with
    // range: 0..20.
    CipslaUdpJitterTmplIcpifFactor interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    CipslaUdpJitterTmplStatsHours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaUdpJitterTmplDistBuckets will be kept.  The last bucket will contain
    // all entries from its  distribution interval start point to infinity. The
    // type is interface{} with range: 1..20.
    CipslaUdpJitterTmplDistBuckets interface{}

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
    CipslaUdpJitterTmplDistInterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaUdpJitterTmplStorageType interface{}

    // The status of the conceptual UDP Jitter template control row. When the
    // status is active, all the read-create objects in that  row can be modified.
    // The type is RowStatus.
    CipslaUdpJitterTmplRowStatus interface{}
}

func (cipslaUdpJitterTmplEntry *CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry) GetEntityData() *types.CommonEntityData {
    cipslaUdpJitterTmplEntry.EntityData.YFilter = cipslaUdpJitterTmplEntry.YFilter
    cipslaUdpJitterTmplEntry.EntityData.YangName = "cipslaUdpJitterTmplEntry"
    cipslaUdpJitterTmplEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaUdpJitterTmplEntry.EntityData.ParentYangName = "cipslaUdpJitterTmplTable"
    cipslaUdpJitterTmplEntry.EntityData.SegmentPath = "cipslaUdpJitterTmplEntry" + types.AddKeyToken(cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplName, "cipslaUdpJitterTmplName")
    cipslaUdpJitterTmplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaUdpJitterTmplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaUdpJitterTmplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaUdpJitterTmplEntry.EntityData.Children = types.NewOrderedMap()
    cipslaUdpJitterTmplEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplName", types.YLeaf{"CipslaUdpJitterTmplName", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplName})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplDescription", types.YLeaf{"CipslaUdpJitterTmplDescription", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplDescription})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplControlEnable", types.YLeaf{"CipslaUdpJitterTmplControlEnable", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplControlEnable})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplTimeOut", types.YLeaf{"CipslaUdpJitterTmplTimeOut", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplTimeOut})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplVerifyData", types.YLeaf{"CipslaUdpJitterTmplVerifyData", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplVerifyData})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplCodecType", types.YLeaf{"CipslaUdpJitterTmplCodecType", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplCodecType})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplCodecInterval", types.YLeaf{"CipslaUdpJitterTmplCodecInterval", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplCodecInterval})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplCodecPayload", types.YLeaf{"CipslaUdpJitterTmplCodecPayload", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplCodecPayload})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplCodecNumPkts", types.YLeaf{"CipslaUdpJitterTmplCodecNumPkts", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplCodecNumPkts})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplInterval", types.YLeaf{"CipslaUdpJitterTmplInterval", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplInterval})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplNumPkts", types.YLeaf{"CipslaUdpJitterTmplNumPkts", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplNumPkts})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplSrcAddrType", types.YLeaf{"CipslaUdpJitterTmplSrcAddrType", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplSrcAddrType})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplSrcAddr", types.YLeaf{"CipslaUdpJitterTmplSrcAddr", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplSrcAddr})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplSrcPort", types.YLeaf{"CipslaUdpJitterTmplSrcPort", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplSrcPort})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplPrecision", types.YLeaf{"CipslaUdpJitterTmplPrecision", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplPrecision})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplReqDataSize", types.YLeaf{"CipslaUdpJitterTmplReqDataSize", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplReqDataSize})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplPktPriority", types.YLeaf{"CipslaUdpJitterTmplPktPriority", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplPktPriority})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplTOS", types.YLeaf{"CipslaUdpJitterTmplTOS", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplTOS})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplVrfName", types.YLeaf{"CipslaUdpJitterTmplVrfName", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplVrfName})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplThreshold", types.YLeaf{"CipslaUdpJitterTmplThreshold", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplThreshold})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplNTPTolAbs", types.YLeaf{"CipslaUdpJitterTmplNTPTolAbs", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplNTPTolAbs})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplNTPTolPct", types.YLeaf{"CipslaUdpJitterTmplNTPTolPct", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplNTPTolPct})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplNTPTolType", types.YLeaf{"CipslaUdpJitterTmplNTPTolType", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplNTPTolType})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplIcpifFactor", types.YLeaf{"CipslaUdpJitterTmplIcpifFactor", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplIcpifFactor})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplStatsHours", types.YLeaf{"CipslaUdpJitterTmplStatsHours", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplStatsHours})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplDistBuckets", types.YLeaf{"CipslaUdpJitterTmplDistBuckets", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplDistBuckets})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplDistInterval", types.YLeaf{"CipslaUdpJitterTmplDistInterval", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplDistInterval})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplStorageType", types.YLeaf{"CipslaUdpJitterTmplStorageType", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplStorageType})
    cipslaUdpJitterTmplEntry.EntityData.Leafs.Append("cipslaUdpJitterTmplRowStatus", types.YLeaf{"CipslaUdpJitterTmplRowStatus", cipslaUdpJitterTmplEntry.CipslaUdpJitterTmplRowStatus})

    cipslaUdpJitterTmplEntry.EntityData.YListKeys = []string {"CipslaUdpJitterTmplName"}

    return &(cipslaUdpJitterTmplEntry.EntityData)
}

// CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType represents               percent value.
type CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType string

const (
    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType_percent CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType = "percent"

    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType_absolute CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplNTPTolType = "absolute"
)

// CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority represents high(2)   - The packet is of high priority.
type CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority string

const (
    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority_normal CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority = "normal"

    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority_high CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPktPriority = "high"
)

// CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision represents microseconds(2) - The accuracy of stats will be in microseconds.
type CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision string

const (
    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision_milliseconds CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision = "milliseconds"

    CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision_microseconds CISCOIPSLAJITTERMIB_CipslaUdpJitterTmplTable_CipslaUdpJitterTmplEntry_CipslaUdpJitterTmplPrecision = "microseconds"
)

// CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable
// A table that contains ICMP jitter template specific definitions.
type CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IP SLA ICMP Jitter template. The type is slice
    // of CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable_CipslaIcmpJitterTmplEntry.
    CipslaIcmpJitterTmplEntry []*CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable_CipslaIcmpJitterTmplEntry
}

func (cipslaIcmpJitterTmplTable *CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable) GetEntityData() *types.CommonEntityData {
    cipslaIcmpJitterTmplTable.EntityData.YFilter = cipslaIcmpJitterTmplTable.YFilter
    cipslaIcmpJitterTmplTable.EntityData.YangName = "cipslaIcmpJitterTmplTable"
    cipslaIcmpJitterTmplTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaIcmpJitterTmplTable.EntityData.ParentYangName = "CISCO-IPSLA-JITTER-MIB"
    cipslaIcmpJitterTmplTable.EntityData.SegmentPath = "cipslaIcmpJitterTmplTable"
    cipslaIcmpJitterTmplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaIcmpJitterTmplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaIcmpJitterTmplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaIcmpJitterTmplTable.EntityData.Children = types.NewOrderedMap()
    cipslaIcmpJitterTmplTable.EntityData.Children.Append("cipslaIcmpJitterTmplEntry", types.YChild{"CipslaIcmpJitterTmplEntry", nil})
    for i := range cipslaIcmpJitterTmplTable.CipslaIcmpJitterTmplEntry {
        cipslaIcmpJitterTmplTable.EntityData.Children.Append(types.GetSegmentPath(cipslaIcmpJitterTmplTable.CipslaIcmpJitterTmplEntry[i]), types.YChild{"CipslaIcmpJitterTmplEntry", cipslaIcmpJitterTmplTable.CipslaIcmpJitterTmplEntry[i]})
    }
    cipslaIcmpJitterTmplTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaIcmpJitterTmplTable.EntityData.YListKeys = []string {}

    return &(cipslaIcmpJitterTmplTable.EntityData)
}

// CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable_CipslaIcmpJitterTmplEntry
// A row entry representing an IP SLA ICMP Jitter template.
type CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable_CipslaIcmpJitterTmplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the ICMP jitter template
    // name. The type is string with length: 1..64.
    CipslaIcmpJitterTmplName interface{}

    // A string which provides description of ICMP Jitter template. The type is
    // string with length: 0..128.
    CipslaIcmpJitterTmplDescription interface{}

    // Specifies the duration to wait for a IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    CipslaIcmpJitterTmplTimeOut interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    CipslaIcmpJitterTmplVerifyData interface{}

    // This value represents the number of packets that need to be transmitted.
    // The type is interface{} with range: 1..60000. Units are packets.
    CipslaIcmpJitterTmplNumPkts interface{}

    // This value represents the inter-packet delay between packets and is in
    // milliseconds. The type is interface{} with range: 4..60000. Units are
    // milliseconds.
    CipslaIcmpJitterTmplInterval interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaIcmpJitterTmplSrcAddr object. The type is
    // InetAddressType.
    CipslaIcmpJitterTmplSrcAddrType interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    CipslaIcmpJitterTmplSrcAddr interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    CipslaIcmpJitterTmplTOS interface{}

    // This field is used to specify the VRF name in which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VPN routing Table
    // for this operation. The type is string with length: 0..32.
    CipslaIcmpJitterTmplVrfName interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit, then  one threshold crossing occurrence
    // will be counted. The type is interface{} with range: 0..2147483647. Units
    // are milliseconds.
    CipslaIcmpJitterTmplThreshold interface{}

    // The maximum number of hourss for which statistics are maintained.
    // Specifically this is the number of hourly groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    CipslaIcmpJitterTmplStatsHours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaIcmpJitterTmplDistBuckets will be kept.  The last bucket will contain
    // all entries from its  distribution interval start point to infinity. The
    // type is interface{} with range: 1..20.
    CipslaIcmpJitterTmplDistBuckets interface{}

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
    CipslaIcmpJitterTmplDistInterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaIcmpJitterTmplStorageType interface{}

    // The status of the conceptual ICMP jitter template control row. When the
    // status is active, all the read-create objects in  that row can be modified.
    // The type is RowStatus.
    CipslaIcmpJitterTmplRowStatus interface{}
}

func (cipslaIcmpJitterTmplEntry *CISCOIPSLAJITTERMIB_CipslaIcmpJitterTmplTable_CipslaIcmpJitterTmplEntry) GetEntityData() *types.CommonEntityData {
    cipslaIcmpJitterTmplEntry.EntityData.YFilter = cipslaIcmpJitterTmplEntry.YFilter
    cipslaIcmpJitterTmplEntry.EntityData.YangName = "cipslaIcmpJitterTmplEntry"
    cipslaIcmpJitterTmplEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaIcmpJitterTmplEntry.EntityData.ParentYangName = "cipslaIcmpJitterTmplTable"
    cipslaIcmpJitterTmplEntry.EntityData.SegmentPath = "cipslaIcmpJitterTmplEntry" + types.AddKeyToken(cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplName, "cipslaIcmpJitterTmplName")
    cipslaIcmpJitterTmplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaIcmpJitterTmplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaIcmpJitterTmplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaIcmpJitterTmplEntry.EntityData.Children = types.NewOrderedMap()
    cipslaIcmpJitterTmplEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplName", types.YLeaf{"CipslaIcmpJitterTmplName", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplName})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplDescription", types.YLeaf{"CipslaIcmpJitterTmplDescription", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplDescription})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplTimeOut", types.YLeaf{"CipslaIcmpJitterTmplTimeOut", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplTimeOut})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplVerifyData", types.YLeaf{"CipslaIcmpJitterTmplVerifyData", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplVerifyData})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplNumPkts", types.YLeaf{"CipslaIcmpJitterTmplNumPkts", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplNumPkts})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplInterval", types.YLeaf{"CipslaIcmpJitterTmplInterval", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplInterval})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplSrcAddrType", types.YLeaf{"CipslaIcmpJitterTmplSrcAddrType", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplSrcAddrType})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplSrcAddr", types.YLeaf{"CipslaIcmpJitterTmplSrcAddr", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplSrcAddr})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplTOS", types.YLeaf{"CipslaIcmpJitterTmplTOS", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplTOS})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplVrfName", types.YLeaf{"CipslaIcmpJitterTmplVrfName", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplVrfName})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplThreshold", types.YLeaf{"CipslaIcmpJitterTmplThreshold", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplThreshold})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplStatsHours", types.YLeaf{"CipslaIcmpJitterTmplStatsHours", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplStatsHours})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplDistBuckets", types.YLeaf{"CipslaIcmpJitterTmplDistBuckets", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplDistBuckets})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplDistInterval", types.YLeaf{"CipslaIcmpJitterTmplDistInterval", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplDistInterval})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplStorageType", types.YLeaf{"CipslaIcmpJitterTmplStorageType", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplStorageType})
    cipslaIcmpJitterTmplEntry.EntityData.Leafs.Append("cipslaIcmpJitterTmplRowStatus", types.YLeaf{"CipslaIcmpJitterTmplRowStatus", cipslaIcmpJitterTmplEntry.CipslaIcmpJitterTmplRowStatus})

    cipslaIcmpJitterTmplEntry.EntityData.YListKeys = []string {"CipslaIcmpJitterTmplName"}

    return &(cipslaIcmpJitterTmplEntry.EntityData)
}

