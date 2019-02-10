// This MIB module defines the templates for IP SLA operations of
// ICMP echo, UDP echo and TCP connect.
// 
// The ICMP echo operation measures end-to-end response time between 
// a Cisco router and any IP enabled device by computing the time
// taken between sending an ICMP echo request message to the 
// destination and receiving an ICMP echo reply.
// 
// 
// The UDP echo operation measures end-to-end response time between 
// a Cisco router and any IP enabled device by computing the time
// taken between sending an UDP echo request message to the 
// destination and receiving an UDP echo reply.
// 
// The TCP connect operation measures end-to-end response time between 
// a Cisco router and any IP enabled device by computing the time
// taken to perform a TCP connect operation.
package cisco_ipsla_echo_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsla_echo_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSLA-ECHO-MIB CISCO-IPSLA-ECHO-MIB}", reflect.TypeOf(CISCOIPSLAECHOMIB{}))
    ydk.RegisterEntity("CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB", reflect.TypeOf(CISCOIPSLAECHOMIB{}))
}

// CISCOIPSLAECHOMIB
type CISCOIPSLAECHOMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table that contains ICMP echo template definitions.
    CipslaIcmpEchoTmplTable CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable

    // A table that contains UDP echo template specific definitions.
    CipslaUdpEchoTmplTable CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable

    // A table that contains TCP connect template specific definitions.
    CipslaTcpConnTmplTable CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSLAECHOMIB.EntityData.YFilter = cISCOIPSLAECHOMIB.YFilter
    cISCOIPSLAECHOMIB.EntityData.YangName = "CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSLAECHOMIB.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.SegmentPath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.AbsolutePath = cISCOIPSLAECHOMIB.EntityData.SegmentPath
    cISCOIPSLAECHOMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSLAECHOMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSLAECHOMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSLAECHOMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSLAECHOMIB.EntityData.Children.Append("cipslaIcmpEchoTmplTable", types.YChild{"CipslaIcmpEchoTmplTable", &cISCOIPSLAECHOMIB.CipslaIcmpEchoTmplTable})
    cISCOIPSLAECHOMIB.EntityData.Children.Append("cipslaUdpEchoTmplTable", types.YChild{"CipslaUdpEchoTmplTable", &cISCOIPSLAECHOMIB.CipslaUdpEchoTmplTable})
    cISCOIPSLAECHOMIB.EntityData.Children.Append("cipslaTcpConnTmplTable", types.YChild{"CipslaTcpConnTmplTable", &cISCOIPSLAECHOMIB.CipslaTcpConnTmplTable})
    cISCOIPSLAECHOMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSLAECHOMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSLAECHOMIB.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable
// A table that contains ICMP echo template definitions.
type CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA ICMP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry.
    CipslaIcmpEchoTmplEntry []*CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry
}

func (cipslaIcmpEchoTmplTable *CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable) GetEntityData() *types.CommonEntityData {
    cipslaIcmpEchoTmplTable.EntityData.YFilter = cipslaIcmpEchoTmplTable.YFilter
    cipslaIcmpEchoTmplTable.EntityData.YangName = "cipslaIcmpEchoTmplTable"
    cipslaIcmpEchoTmplTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaIcmpEchoTmplTable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslaIcmpEchoTmplTable.EntityData.SegmentPath = "cipslaIcmpEchoTmplTable"
    cipslaIcmpEchoTmplTable.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/" + cipslaIcmpEchoTmplTable.EntityData.SegmentPath
    cipslaIcmpEchoTmplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaIcmpEchoTmplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaIcmpEchoTmplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaIcmpEchoTmplTable.EntityData.Children = types.NewOrderedMap()
    cipslaIcmpEchoTmplTable.EntityData.Children.Append("cipslaIcmpEchoTmplEntry", types.YChild{"CipslaIcmpEchoTmplEntry", nil})
    for i := range cipslaIcmpEchoTmplTable.CipslaIcmpEchoTmplEntry {
        cipslaIcmpEchoTmplTable.EntityData.Children.Append(types.GetSegmentPath(cipslaIcmpEchoTmplTable.CipslaIcmpEchoTmplEntry[i]), types.YChild{"CipslaIcmpEchoTmplEntry", cipslaIcmpEchoTmplTable.CipslaIcmpEchoTmplEntry[i]})
    }
    cipslaIcmpEchoTmplTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaIcmpEchoTmplTable.EntityData.YListKeys = []string {}

    return &(cipslaIcmpEchoTmplTable.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry
// A row entry representing an IPSLA ICMP echo template.
type CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This field is used to specify the ICMP echo
    // template name. The type is string with length: 1..64.
    CipslaIcmpEchoTmplName interface{}

    // This field is used to provide description for the ICMP echo template. The
    // type is string with length: 0..128.
    CipslaIcmpEchoTmplDescription interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaIcmpEchoTmplSrcAddr object. The type is
    // InetAddressType.
    CipslaIcmpEchoTmplSrcAddrType interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    CipslaIcmpEchoTmplSrcAddr interface{}

    // Specifies the duration to wait for a IP SLA operation completion.   For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    CipslaIcmpEchoTmplTimeOut interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    CipslaIcmpEchoTmplVerifyData interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' IP SLA request/responses, this value represents the native
    // payload size.  REMEMBER:  The ARR Header overhead is not included          
    // in this value. The type is interface{} with range: 0..16384. Units are
    // octets.
    CipslaIcmpEchoTmplReqDataSize interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    CipslaIcmpEchoTmplTOS interface{}

    // This field is used to specify the VRF name with which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VRF routing table
    // for this operation. The type is string with length: 0..32.
    CipslaIcmpEchoTmplVrfName interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaIcmpEchoTmplHistFilter is satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    CipslaIcmpEchoTmplThreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    CipslaIcmpEchoTmplHistLives interface{}

    // The maximum number of history buckets to record. This value is set to the
    // number of operations  to keep per lifetime.  After
    // cipslaIcmpEchoTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaIcmpEchoTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    CipslaIcmpEchoTmplHistBuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaIcmpEchoTmplThreshold are                     recorded. failures(4)  
    // - the results of failed operations (only)                     are recorded.
    // The type is CipslaIcmpEchoTmplHistFilter.
    CipslaIcmpEchoTmplHistFilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    CipslaIcmpEchoTmplStatsHours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaIcmpEchoTmplStatsNumDistBuckets will be kept.  The last
    // cipslaIcmpEchoTmplStatsNumDistBucket will contain all entries from its
    // distribution interval start point to infinity. The type is interface{} with
    // range: 1..20.
    CipslaIcmpEchoTmplDistBuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  cipslaIcmpEchoTmplDistBuckets = 5 buckets
    // cipslaIcmpEchoTmplDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2 |
    // Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms | 30-39
    // ms | 40-Inf ms |  Odd Example:  cipslaIcmpEchoTmplDistBuckets = 1 buckets
    // cipslaIcmpEchoTmplDistInterval = 10 milliseconds  | Bucket 1  | |  0-Inf ms
    // |  Thus, this odd example shows that the value of
    // cipslaIcmpEchoTmplDistInterval does not apply when
    // cipslaIcmpEchoTmplDistBuckets is one. The type is interface{} with range:
    // 1..100. Units are milliseconds.
    CipslaIcmpEchoTmplDistInterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaIcmpEchoTmplStorageType interface{}

    // The status of the conceptual ICMP echo template control row. When the
    // status is active, all the read-create objects in that  row can be modified.
    // The type is RowStatus.
    CipslaIcmpEchoTmplRowStatus interface{}
}

func (cipslaIcmpEchoTmplEntry *CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry) GetEntityData() *types.CommonEntityData {
    cipslaIcmpEchoTmplEntry.EntityData.YFilter = cipslaIcmpEchoTmplEntry.YFilter
    cipslaIcmpEchoTmplEntry.EntityData.YangName = "cipslaIcmpEchoTmplEntry"
    cipslaIcmpEchoTmplEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaIcmpEchoTmplEntry.EntityData.ParentYangName = "cipslaIcmpEchoTmplTable"
    cipslaIcmpEchoTmplEntry.EntityData.SegmentPath = "cipslaIcmpEchoTmplEntry" + types.AddKeyToken(cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplName, "cipslaIcmpEchoTmplName")
    cipslaIcmpEchoTmplEntry.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/cipslaIcmpEchoTmplTable/" + cipslaIcmpEchoTmplEntry.EntityData.SegmentPath
    cipslaIcmpEchoTmplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaIcmpEchoTmplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaIcmpEchoTmplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaIcmpEchoTmplEntry.EntityData.Children = types.NewOrderedMap()
    cipslaIcmpEchoTmplEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplName", types.YLeaf{"CipslaIcmpEchoTmplName", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplName})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplDescription", types.YLeaf{"CipslaIcmpEchoTmplDescription", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplDescription})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplSrcAddrType", types.YLeaf{"CipslaIcmpEchoTmplSrcAddrType", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplSrcAddrType})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplSrcAddr", types.YLeaf{"CipslaIcmpEchoTmplSrcAddr", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplSrcAddr})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplTimeOut", types.YLeaf{"CipslaIcmpEchoTmplTimeOut", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplTimeOut})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplVerifyData", types.YLeaf{"CipslaIcmpEchoTmplVerifyData", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplVerifyData})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplReqDataSize", types.YLeaf{"CipslaIcmpEchoTmplReqDataSize", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplReqDataSize})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplTOS", types.YLeaf{"CipslaIcmpEchoTmplTOS", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplTOS})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplVrfName", types.YLeaf{"CipslaIcmpEchoTmplVrfName", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplVrfName})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplThreshold", types.YLeaf{"CipslaIcmpEchoTmplThreshold", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplThreshold})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplHistLives", types.YLeaf{"CipslaIcmpEchoTmplHistLives", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplHistLives})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplHistBuckets", types.YLeaf{"CipslaIcmpEchoTmplHistBuckets", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplHistBuckets})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplHistFilter", types.YLeaf{"CipslaIcmpEchoTmplHistFilter", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplHistFilter})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplStatsHours", types.YLeaf{"CipslaIcmpEchoTmplStatsHours", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplStatsHours})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplDistBuckets", types.YLeaf{"CipslaIcmpEchoTmplDistBuckets", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplDistBuckets})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplDistInterval", types.YLeaf{"CipslaIcmpEchoTmplDistInterval", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplDistInterval})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplStorageType", types.YLeaf{"CipslaIcmpEchoTmplStorageType", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplStorageType})
    cipslaIcmpEchoTmplEntry.EntityData.Leafs.Append("cipslaIcmpEchoTmplRowStatus", types.YLeaf{"CipslaIcmpEchoTmplRowStatus", cipslaIcmpEchoTmplEntry.CipslaIcmpEchoTmplRowStatus})

    cipslaIcmpEchoTmplEntry.EntityData.YListKeys = []string {"CipslaIcmpEchoTmplName"}

    return &(cipslaIcmpEchoTmplEntry.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter represents                    are recorded.
type CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter string

const (
    CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter_none CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter = "none"

    CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter_all CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter = "all"

    CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter_overThreshold CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter = "overThreshold"

    CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter_failures CISCOIPSLAECHOMIB_CipslaIcmpEchoTmplTable_CipslaIcmpEchoTmplEntry_CipslaIcmpEchoTmplHistFilter = "failures"
)

// CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable
// A table that contains UDP echo template specific definitions.
type CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry.
    CipslaUdpEchoTmplEntry []*CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry
}

func (cipslaUdpEchoTmplTable *CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable) GetEntityData() *types.CommonEntityData {
    cipslaUdpEchoTmplTable.EntityData.YFilter = cipslaUdpEchoTmplTable.YFilter
    cipslaUdpEchoTmplTable.EntityData.YangName = "cipslaUdpEchoTmplTable"
    cipslaUdpEchoTmplTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaUdpEchoTmplTable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslaUdpEchoTmplTable.EntityData.SegmentPath = "cipslaUdpEchoTmplTable"
    cipslaUdpEchoTmplTable.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/" + cipslaUdpEchoTmplTable.EntityData.SegmentPath
    cipslaUdpEchoTmplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaUdpEchoTmplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaUdpEchoTmplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaUdpEchoTmplTable.EntityData.Children = types.NewOrderedMap()
    cipslaUdpEchoTmplTable.EntityData.Children.Append("cipslaUdpEchoTmplEntry", types.YChild{"CipslaUdpEchoTmplEntry", nil})
    for i := range cipslaUdpEchoTmplTable.CipslaUdpEchoTmplEntry {
        cipslaUdpEchoTmplTable.EntityData.Children.Append(types.GetSegmentPath(cipslaUdpEchoTmplTable.CipslaUdpEchoTmplEntry[i]), types.YChild{"CipslaUdpEchoTmplEntry", cipslaUdpEchoTmplTable.CipslaUdpEchoTmplEntry[i]})
    }
    cipslaUdpEchoTmplTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaUdpEchoTmplTable.EntityData.YListKeys = []string {}

    return &(cipslaUdpEchoTmplTable.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry
// A row entry representing an IPSLA UDP echo template.
type CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A string which specifies the UDP echo template
    // name. The type is string with length: 1..64.
    CipslaUdpEchoTmplName interface{}

    // A string which provides description to the UDP echo template. The type is
    // string with length: 0..128.
    CipslaUdpEchoTmplDescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    CipslaUdpEchoTmplControlEnable interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaUdpEchoTmplSrcAddr object. The type is
    // InetAddressType.
    CipslaUdpEchoTmplSrcAddrType interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    CipslaUdpEchoTmplSrcAddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    CipslaUdpEchoTmplSrcPort interface{}

    // Specifies the duration to wait for an IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    CipslaUdpEchoTmplTimeOut interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    CipslaUdpEchoTmplVerifyData interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' RTT request/responses, this value represents the native payload
    // size.  REMEMBER:  The ARR Header overhead is not included            in
    // this value. The type is interface{} with range: 4..1500. Units are octets.
    CipslaUdpEchoTmplReqDataSize interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    CipslaUdpEchoTmplTOS interface{}

    // This field is used to specify the VRF name with which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VRF routing Table
    // for this operation. The type is string with length: 0..32.
    CipslaUdpEchoTmplVrfName interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaUdpEchoTmplHistFilter is  satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    CipslaUdpEchoTmplThreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    CipslaUdpEchoTmplHistLives interface{}

    // The maximum number of history buckets to record. This value should be set
    // to the number of operations  to keep per lifetime.  After
    // cipslaUdpEchoTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaUdpEchoTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    CipslaUdpEchoTmplHistBuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaUdpEchoTmplThreshold are                     recorded. failures(4)   
    // - the results of failed operations (only)                     are recorded.
    // The type is CipslaUdpEchoTmplHistFilter.
    CipslaUdpEchoTmplHistFilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    CipslaUdpEchoTmplStatsHours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaUdpEchoTmplStatsNumDistBuckets will be kept.  The last
    // cipslaUdpEchoTmplStatsNumDistBuckets will contain all entries from its
    // distribution interval start point to infinity. The type is interface{} with
    // range: 1..20.
    CipslaUdpEchoTmplDistBuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  cipslaUdpEchoTmplDistBuckets = 5 buckets
    // cipslaUdpEchoTmplDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2 |
    // Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms | 30-39
    // ms | 40-Inf ms |  Odd Example:  cipslaUdpEchoTmplDistBuckets = 1 buckets
    // cipslaUdpEchoTmplDistInterval = 10 milliseconds  | Bucket 1  | |  0-Inf ms
    // |  Thus, this odd example shows that the value of
    // cipslaUdpEchoTmplDistInterval does not apply when
    // cipslaUdpEchoTmplDistBuckets is one. The type is interface{} with range:
    // 1..100. Units are milliseconds.
    CipslaUdpEchoTmplDistInterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaUdpEchoTmplStorageType interface{}

    // The status of the conceptual UDP echo template control row. When the status
    // is active, all the read-create objects in  that row can be modified. The
    // type is RowStatus.
    CipslaUdpEchoTmplRowStatus interface{}
}

func (cipslaUdpEchoTmplEntry *CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry) GetEntityData() *types.CommonEntityData {
    cipslaUdpEchoTmplEntry.EntityData.YFilter = cipslaUdpEchoTmplEntry.YFilter
    cipslaUdpEchoTmplEntry.EntityData.YangName = "cipslaUdpEchoTmplEntry"
    cipslaUdpEchoTmplEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaUdpEchoTmplEntry.EntityData.ParentYangName = "cipslaUdpEchoTmplTable"
    cipslaUdpEchoTmplEntry.EntityData.SegmentPath = "cipslaUdpEchoTmplEntry" + types.AddKeyToken(cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplName, "cipslaUdpEchoTmplName")
    cipslaUdpEchoTmplEntry.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/cipslaUdpEchoTmplTable/" + cipslaUdpEchoTmplEntry.EntityData.SegmentPath
    cipslaUdpEchoTmplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaUdpEchoTmplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaUdpEchoTmplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaUdpEchoTmplEntry.EntityData.Children = types.NewOrderedMap()
    cipslaUdpEchoTmplEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplName", types.YLeaf{"CipslaUdpEchoTmplName", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplName})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplDescription", types.YLeaf{"CipslaUdpEchoTmplDescription", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplDescription})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplControlEnable", types.YLeaf{"CipslaUdpEchoTmplControlEnable", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplControlEnable})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplSrcAddrType", types.YLeaf{"CipslaUdpEchoTmplSrcAddrType", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplSrcAddrType})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplSrcAddr", types.YLeaf{"CipslaUdpEchoTmplSrcAddr", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplSrcAddr})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplSrcPort", types.YLeaf{"CipslaUdpEchoTmplSrcPort", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplSrcPort})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplTimeOut", types.YLeaf{"CipslaUdpEchoTmplTimeOut", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplTimeOut})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplVerifyData", types.YLeaf{"CipslaUdpEchoTmplVerifyData", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplVerifyData})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplReqDataSize", types.YLeaf{"CipslaUdpEchoTmplReqDataSize", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplReqDataSize})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplTOS", types.YLeaf{"CipslaUdpEchoTmplTOS", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplTOS})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplVrfName", types.YLeaf{"CipslaUdpEchoTmplVrfName", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplVrfName})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplThreshold", types.YLeaf{"CipslaUdpEchoTmplThreshold", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplThreshold})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplHistLives", types.YLeaf{"CipslaUdpEchoTmplHistLives", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplHistLives})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplHistBuckets", types.YLeaf{"CipslaUdpEchoTmplHistBuckets", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplHistBuckets})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplHistFilter", types.YLeaf{"CipslaUdpEchoTmplHistFilter", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplHistFilter})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplStatsHours", types.YLeaf{"CipslaUdpEchoTmplStatsHours", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplStatsHours})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplDistBuckets", types.YLeaf{"CipslaUdpEchoTmplDistBuckets", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplDistBuckets})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplDistInterval", types.YLeaf{"CipslaUdpEchoTmplDistInterval", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplDistInterval})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplStorageType", types.YLeaf{"CipslaUdpEchoTmplStorageType", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplStorageType})
    cipslaUdpEchoTmplEntry.EntityData.Leafs.Append("cipslaUdpEchoTmplRowStatus", types.YLeaf{"CipslaUdpEchoTmplRowStatus", cipslaUdpEchoTmplEntry.CipslaUdpEchoTmplRowStatus})

    cipslaUdpEchoTmplEntry.EntityData.YListKeys = []string {"CipslaUdpEchoTmplName"}

    return &(cipslaUdpEchoTmplEntry.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter represents                    are recorded.
type CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter string

const (
    CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter_none CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter = "none"

    CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter_all CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter = "all"

    CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter_overThreshold CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter = "overThreshold"

    CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter_failures CISCOIPSLAECHOMIB_CipslaUdpEchoTmplTable_CipslaUdpEchoTmplEntry_CipslaUdpEchoTmplHistFilter = "failures"
)

// CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable
// A table that contains TCP connect template specific definitions.
type CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA TCP connect template. The type is slice
    // of CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry.
    CipslaTcpConnTmplEntry []*CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry
}

func (cipslaTcpConnTmplTable *CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable) GetEntityData() *types.CommonEntityData {
    cipslaTcpConnTmplTable.EntityData.YFilter = cipslaTcpConnTmplTable.YFilter
    cipslaTcpConnTmplTable.EntityData.YangName = "cipslaTcpConnTmplTable"
    cipslaTcpConnTmplTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaTcpConnTmplTable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslaTcpConnTmplTable.EntityData.SegmentPath = "cipslaTcpConnTmplTable"
    cipslaTcpConnTmplTable.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/" + cipslaTcpConnTmplTable.EntityData.SegmentPath
    cipslaTcpConnTmplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaTcpConnTmplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaTcpConnTmplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaTcpConnTmplTable.EntityData.Children = types.NewOrderedMap()
    cipslaTcpConnTmplTable.EntityData.Children.Append("cipslaTcpConnTmplEntry", types.YChild{"CipslaTcpConnTmplEntry", nil})
    for i := range cipslaTcpConnTmplTable.CipslaTcpConnTmplEntry {
        cipslaTcpConnTmplTable.EntityData.Children.Append(types.GetSegmentPath(cipslaTcpConnTmplTable.CipslaTcpConnTmplEntry[i]), types.YChild{"CipslaTcpConnTmplEntry", cipslaTcpConnTmplTable.CipslaTcpConnTmplEntry[i]})
    }
    cipslaTcpConnTmplTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaTcpConnTmplTable.EntityData.YListKeys = []string {}

    return &(cipslaTcpConnTmplTable.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry
// A row entry representing an IPSLA TCP connect template.
type CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A string which specifies the TCP connect template
    // name. The type is string with length: 1..64.
    CipslaTcpConnTmplName interface{}

    // A string which provides description for the TCP connect template. The type
    // is string with length: 0..128.
    CipslaTcpConnTmplDescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    CipslaTcpConnTmplControlEnable interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaTcpConnTmplSrcAddr object. The type is
    // InetAddressType.
    CipslaTcpConnTmplSrcAddrType interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    CipslaTcpConnTmplSrcAddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    CipslaTcpConnTmplSrcPort interface{}

    // Specifies the duration to wait for an IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    CipslaTcpConnTmplTimeOut interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    CipslaTcpConnTmplVerifyData interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    CipslaTcpConnTmplTOS interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaTcpConnTmplHistFilter is  satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    CipslaTcpConnTmplThreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    CipslaTcpConnTmplHistLives interface{}

    // The maximum number of history buckets to record. This value should be set
    // to the number of operations  to keep per lifetime.  After
    // cipslaTcpConnTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaTcpConnTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    CipslaTcpConnTmplHistBuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaTcpConnTmplThreshold are                     recorded. failures(4)   
    // - the results of failed operations (only)                     are recorded.
    // The type is CipslaTcpConnTmplHistFilter.
    CipslaTcpConnTmplHistFilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    CipslaTcpConnTmplStatsHours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaTcpConnTmplDistBuckets will be kept.  The last
    // cipslaTcpConnTmplDistBuckets will contain all entries from its distribution
    // interval start point to infinity. The type is interface{} with range:
    // 1..20.
    CipslaTcpConnTmplDistBuckets interface{}

    // The statistical distribution buckets interval.  Distribution Bucket
    // Example:  cipslaTcpConnTmplDistBuckets = 5 buckets
    // cipslaTcpConnTmplDistInterval = 10 milliseconds  | Bucket 1 | Bucket 2 |
    // Bucket 3 | Bucket 4 | Bucket 5  | |  0-9 ms  | 10-19 ms | 20-29 ms | 30-39
    // ms | 40-Inf ms |  Odd Example:  cipslaTcpConnTmplDistBuckets = 1 buckets
    // cipslaTcpConnTmplDistInterval = 10 milliseconds  | Bucket 1  | |  0-Inf ms
    // |  Thus, this odd example shows that the value of
    // cipslaTcpConnTmplDistInterval does not apply when
    // cipslaTcpConnTmplDistBuckets is one. The type is interface{} with range:
    // 1..100. Units are milliseconds.
    CipslaTcpConnTmplDistInterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaTcpConnTmplStorageType interface{}

    // The status of the conceptual tcp connect control row. When the status is
    // active, all the read-create objects  in that row can be modified. The type
    // is RowStatus.
    CipslaTcpConnTmplRowStatus interface{}
}

func (cipslaTcpConnTmplEntry *CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry) GetEntityData() *types.CommonEntityData {
    cipslaTcpConnTmplEntry.EntityData.YFilter = cipslaTcpConnTmplEntry.YFilter
    cipslaTcpConnTmplEntry.EntityData.YangName = "cipslaTcpConnTmplEntry"
    cipslaTcpConnTmplEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaTcpConnTmplEntry.EntityData.ParentYangName = "cipslaTcpConnTmplTable"
    cipslaTcpConnTmplEntry.EntityData.SegmentPath = "cipslaTcpConnTmplEntry" + types.AddKeyToken(cipslaTcpConnTmplEntry.CipslaTcpConnTmplName, "cipslaTcpConnTmplName")
    cipslaTcpConnTmplEntry.EntityData.AbsolutePath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB/cipslaTcpConnTmplTable/" + cipslaTcpConnTmplEntry.EntityData.SegmentPath
    cipslaTcpConnTmplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaTcpConnTmplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaTcpConnTmplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaTcpConnTmplEntry.EntityData.Children = types.NewOrderedMap()
    cipslaTcpConnTmplEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplName", types.YLeaf{"CipslaTcpConnTmplName", cipslaTcpConnTmplEntry.CipslaTcpConnTmplName})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplDescription", types.YLeaf{"CipslaTcpConnTmplDescription", cipslaTcpConnTmplEntry.CipslaTcpConnTmplDescription})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplControlEnable", types.YLeaf{"CipslaTcpConnTmplControlEnable", cipslaTcpConnTmplEntry.CipslaTcpConnTmplControlEnable})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplSrcAddrType", types.YLeaf{"CipslaTcpConnTmplSrcAddrType", cipslaTcpConnTmplEntry.CipslaTcpConnTmplSrcAddrType})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplSrcAddr", types.YLeaf{"CipslaTcpConnTmplSrcAddr", cipslaTcpConnTmplEntry.CipslaTcpConnTmplSrcAddr})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplSrcPort", types.YLeaf{"CipslaTcpConnTmplSrcPort", cipslaTcpConnTmplEntry.CipslaTcpConnTmplSrcPort})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplTimeOut", types.YLeaf{"CipslaTcpConnTmplTimeOut", cipslaTcpConnTmplEntry.CipslaTcpConnTmplTimeOut})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplVerifyData", types.YLeaf{"CipslaTcpConnTmplVerifyData", cipslaTcpConnTmplEntry.CipslaTcpConnTmplVerifyData})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplTOS", types.YLeaf{"CipslaTcpConnTmplTOS", cipslaTcpConnTmplEntry.CipslaTcpConnTmplTOS})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplThreshold", types.YLeaf{"CipslaTcpConnTmplThreshold", cipslaTcpConnTmplEntry.CipslaTcpConnTmplThreshold})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplHistLives", types.YLeaf{"CipslaTcpConnTmplHistLives", cipslaTcpConnTmplEntry.CipslaTcpConnTmplHistLives})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplHistBuckets", types.YLeaf{"CipslaTcpConnTmplHistBuckets", cipslaTcpConnTmplEntry.CipslaTcpConnTmplHistBuckets})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplHistFilter", types.YLeaf{"CipslaTcpConnTmplHistFilter", cipslaTcpConnTmplEntry.CipslaTcpConnTmplHistFilter})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplStatsHours", types.YLeaf{"CipslaTcpConnTmplStatsHours", cipslaTcpConnTmplEntry.CipslaTcpConnTmplStatsHours})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplDistBuckets", types.YLeaf{"CipslaTcpConnTmplDistBuckets", cipslaTcpConnTmplEntry.CipslaTcpConnTmplDistBuckets})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplDistInterval", types.YLeaf{"CipslaTcpConnTmplDistInterval", cipslaTcpConnTmplEntry.CipslaTcpConnTmplDistInterval})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplStorageType", types.YLeaf{"CipslaTcpConnTmplStorageType", cipslaTcpConnTmplEntry.CipslaTcpConnTmplStorageType})
    cipslaTcpConnTmplEntry.EntityData.Leafs.Append("cipslaTcpConnTmplRowStatus", types.YLeaf{"CipslaTcpConnTmplRowStatus", cipslaTcpConnTmplEntry.CipslaTcpConnTmplRowStatus})

    cipslaTcpConnTmplEntry.EntityData.YListKeys = []string {"CipslaTcpConnTmplName"}

    return &(cipslaTcpConnTmplEntry.EntityData)
}

// CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter represents                    are recorded.
type CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter string

const (
    CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter_none CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter = "none"

    CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter_all CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter = "all"

    CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter_overThreshold CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter = "overThreshold"

    CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter_failures CISCOIPSLAECHOMIB_CipslaTcpConnTmplTable_CipslaTcpConnTmplEntry_CipslaTcpConnTmplHistFilter = "failures"
)

