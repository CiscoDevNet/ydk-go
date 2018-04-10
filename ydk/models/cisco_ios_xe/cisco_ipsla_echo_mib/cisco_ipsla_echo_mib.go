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
    Cipslaicmpechotmpltable CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable

    // A table that contains UDP echo template specific definitions.
    Cipslaudpechotmpltable CISCOIPSLAECHOMIB_Cipslaudpechotmpltable

    // A table that contains TCP connect template specific definitions.
    Cipslatcpconntmpltable CISCOIPSLAECHOMIB_Cipslatcpconntmpltable
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSLAECHOMIB.EntityData.YFilter = cISCOIPSLAECHOMIB.YFilter
    cISCOIPSLAECHOMIB.EntityData.YangName = "CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSLAECHOMIB.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.SegmentPath = "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB"
    cISCOIPSLAECHOMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSLAECHOMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSLAECHOMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSLAECHOMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSLAECHOMIB.EntityData.Children["cipslaIcmpEchoTmplTable"] = types.YChild{"Cipslaicmpechotmpltable", &cISCOIPSLAECHOMIB.Cipslaicmpechotmpltable}
    cISCOIPSLAECHOMIB.EntityData.Children["cipslaUdpEchoTmplTable"] = types.YChild{"Cipslaudpechotmpltable", &cISCOIPSLAECHOMIB.Cipslaudpechotmpltable}
    cISCOIPSLAECHOMIB.EntityData.Children["cipslaTcpConnTmplTable"] = types.YChild{"Cipslatcpconntmpltable", &cISCOIPSLAECHOMIB.Cipslatcpconntmpltable}
    cISCOIPSLAECHOMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSLAECHOMIB.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable
// A table that contains ICMP echo template definitions.
type CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA ICMP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry.
    Cipslaicmpechotmplentry []CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetEntityData() *types.CommonEntityData {
    cipslaicmpechotmpltable.EntityData.YFilter = cipslaicmpechotmpltable.YFilter
    cipslaicmpechotmpltable.EntityData.YangName = "cipslaIcmpEchoTmplTable"
    cipslaicmpechotmpltable.EntityData.BundleName = "cisco_ios_xe"
    cipslaicmpechotmpltable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslaicmpechotmpltable.EntityData.SegmentPath = "cipslaIcmpEchoTmplTable"
    cipslaicmpechotmpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaicmpechotmpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaicmpechotmpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaicmpechotmpltable.EntityData.Children = make(map[string]types.YChild)
    cipslaicmpechotmpltable.EntityData.Children["cipslaIcmpEchoTmplEntry"] = types.YChild{"Cipslaicmpechotmplentry", nil}
    for i := range cipslaicmpechotmpltable.Cipslaicmpechotmplentry {
        cipslaicmpechotmpltable.EntityData.Children[types.GetSegmentPath(&cipslaicmpechotmpltable.Cipslaicmpechotmplentry[i])] = types.YChild{"Cipslaicmpechotmplentry", &cipslaicmpechotmpltable.Cipslaicmpechotmplentry[i]}
    }
    cipslaicmpechotmpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaicmpechotmpltable.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry
// A row entry representing an IPSLA ICMP echo template.
type CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This field is used to specify the ICMP echo
    // template name. The type is string with length: 1..64.
    Cipslaicmpechotmplname interface{}

    // This field is used to provide description for the ICMP echo template. The
    // type is string with length: 0..128.
    Cipslaicmpechotmpldescription interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaIcmpEchoTmplSrcAddr object. The type is
    // InetAddressType.
    Cipslaicmpechotmplsrcaddrtype interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    Cipslaicmpechotmplsrcaddr interface{}

    // Specifies the duration to wait for a IP SLA operation completion.   For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    Cipslaicmpechotmpltimeout interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    Cipslaicmpechotmplverifydata interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' IP SLA request/responses, this value represents the native
    // payload size.  REMEMBER:  The ARR Header overhead is not included          
    // in this value. The type is interface{} with range: 0..16384. Units are
    // octets.
    Cipslaicmpechotmplreqdatasize interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    Cipslaicmpechotmpltos interface{}

    // This field is used to specify the VRF name with which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VRF routing table
    // for this operation. The type is string with length: 0..32.
    Cipslaicmpechotmplvrfname interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaIcmpEchoTmplHistFilter is satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    Cipslaicmpechotmplthreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    Cipslaicmpechotmplhistlives interface{}

    // The maximum number of history buckets to record. This value is set to the
    // number of operations  to keep per lifetime.  After
    // cipslaIcmpEchoTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaIcmpEchoTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    Cipslaicmpechotmplhistbuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaIcmpEchoTmplThreshold are                     recorded. failures(4)  
    // - the results of failed operations (only)                     are recorded.
    // The type is Cipslaicmpechotmplhistfilter.
    Cipslaicmpechotmplhistfilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    Cipslaicmpechotmplstatshours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaIcmpEchoTmplStatsNumDistBuckets will be kept.  The last
    // cipslaIcmpEchoTmplStatsNumDistBucket will contain all entries from its
    // distribution interval start point to infinity. The type is interface{} with
    // range: 1..20.
    Cipslaicmpechotmpldistbuckets interface{}

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
    Cipslaicmpechotmpldistinterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslaicmpechotmplstoragetype interface{}

    // The status of the conceptual ICMP echo template control row. When the
    // status is active, all the read-create objects in that  row can be modified.
    // The type is RowStatus.
    Cipslaicmpechotmplrowstatus interface{}
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetEntityData() *types.CommonEntityData {
    cipslaicmpechotmplentry.EntityData.YFilter = cipslaicmpechotmplentry.YFilter
    cipslaicmpechotmplentry.EntityData.YangName = "cipslaIcmpEchoTmplEntry"
    cipslaicmpechotmplentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaicmpechotmplentry.EntityData.ParentYangName = "cipslaIcmpEchoTmplTable"
    cipslaicmpechotmplentry.EntityData.SegmentPath = "cipslaIcmpEchoTmplEntry" + "[cipslaIcmpEchoTmplName='" + fmt.Sprintf("%v", cipslaicmpechotmplentry.Cipslaicmpechotmplname) + "']"
    cipslaicmpechotmplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaicmpechotmplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaicmpechotmplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaicmpechotmplentry.EntityData.Children = make(map[string]types.YChild)
    cipslaicmpechotmplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplName"] = types.YLeaf{"Cipslaicmpechotmplname", cipslaicmpechotmplentry.Cipslaicmpechotmplname}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplDescription"] = types.YLeaf{"Cipslaicmpechotmpldescription", cipslaicmpechotmplentry.Cipslaicmpechotmpldescription}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplSrcAddrType"] = types.YLeaf{"Cipslaicmpechotmplsrcaddrtype", cipslaicmpechotmplentry.Cipslaicmpechotmplsrcaddrtype}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplSrcAddr"] = types.YLeaf{"Cipslaicmpechotmplsrcaddr", cipslaicmpechotmplentry.Cipslaicmpechotmplsrcaddr}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplTimeOut"] = types.YLeaf{"Cipslaicmpechotmpltimeout", cipslaicmpechotmplentry.Cipslaicmpechotmpltimeout}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplVerifyData"] = types.YLeaf{"Cipslaicmpechotmplverifydata", cipslaicmpechotmplentry.Cipslaicmpechotmplverifydata}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplReqDataSize"] = types.YLeaf{"Cipslaicmpechotmplreqdatasize", cipslaicmpechotmplentry.Cipslaicmpechotmplreqdatasize}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplTOS"] = types.YLeaf{"Cipslaicmpechotmpltos", cipslaicmpechotmplentry.Cipslaicmpechotmpltos}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplVrfName"] = types.YLeaf{"Cipslaicmpechotmplvrfname", cipslaicmpechotmplentry.Cipslaicmpechotmplvrfname}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplThreshold"] = types.YLeaf{"Cipslaicmpechotmplthreshold", cipslaicmpechotmplentry.Cipslaicmpechotmplthreshold}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplHistLives"] = types.YLeaf{"Cipslaicmpechotmplhistlives", cipslaicmpechotmplentry.Cipslaicmpechotmplhistlives}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplHistBuckets"] = types.YLeaf{"Cipslaicmpechotmplhistbuckets", cipslaicmpechotmplentry.Cipslaicmpechotmplhistbuckets}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplHistFilter"] = types.YLeaf{"Cipslaicmpechotmplhistfilter", cipslaicmpechotmplentry.Cipslaicmpechotmplhistfilter}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplStatsHours"] = types.YLeaf{"Cipslaicmpechotmplstatshours", cipslaicmpechotmplentry.Cipslaicmpechotmplstatshours}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplDistBuckets"] = types.YLeaf{"Cipslaicmpechotmpldistbuckets", cipslaicmpechotmplentry.Cipslaicmpechotmpldistbuckets}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplDistInterval"] = types.YLeaf{"Cipslaicmpechotmpldistinterval", cipslaicmpechotmplentry.Cipslaicmpechotmpldistinterval}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplStorageType"] = types.YLeaf{"Cipslaicmpechotmplstoragetype", cipslaicmpechotmplentry.Cipslaicmpechotmplstoragetype}
    cipslaicmpechotmplentry.EntityData.Leafs["cipslaIcmpEchoTmplRowStatus"] = types.YLeaf{"Cipslaicmpechotmplrowstatus", cipslaicmpechotmplentry.Cipslaicmpechotmplrowstatus}
    return &(cipslaicmpechotmplentry.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter represents                    are recorded.
type CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter string

const (
    CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter_none CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter = "none"

    CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter_all CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter = "all"

    CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter_overThreshold CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter = "overThreshold"

    CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter_failures CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry_Cipslaicmpechotmplhistfilter = "failures"
)

// CISCOIPSLAECHOMIB_Cipslaudpechotmpltable
// A table that contains UDP echo template specific definitions.
type CISCOIPSLAECHOMIB_Cipslaudpechotmpltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry.
    Cipslaudpechotmplentry []CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetEntityData() *types.CommonEntityData {
    cipslaudpechotmpltable.EntityData.YFilter = cipslaudpechotmpltable.YFilter
    cipslaudpechotmpltable.EntityData.YangName = "cipslaUdpEchoTmplTable"
    cipslaudpechotmpltable.EntityData.BundleName = "cisco_ios_xe"
    cipslaudpechotmpltable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslaudpechotmpltable.EntityData.SegmentPath = "cipslaUdpEchoTmplTable"
    cipslaudpechotmpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaudpechotmpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaudpechotmpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaudpechotmpltable.EntityData.Children = make(map[string]types.YChild)
    cipslaudpechotmpltable.EntityData.Children["cipslaUdpEchoTmplEntry"] = types.YChild{"Cipslaudpechotmplentry", nil}
    for i := range cipslaudpechotmpltable.Cipslaudpechotmplentry {
        cipslaudpechotmpltable.EntityData.Children[types.GetSegmentPath(&cipslaudpechotmpltable.Cipslaudpechotmplentry[i])] = types.YChild{"Cipslaudpechotmplentry", &cipslaudpechotmpltable.Cipslaudpechotmplentry[i]}
    }
    cipslaudpechotmpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaudpechotmpltable.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry
// A row entry representing an IPSLA UDP echo template.
type CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the UDP echo template
    // name. The type is string with length: 1..64.
    Cipslaudpechotmplname interface{}

    // A string which provides description to the UDP echo template. The type is
    // string with length: 0..128.
    Cipslaudpechotmpldescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    Cipslaudpechotmplcontrolenable interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaUdpEchoTmplSrcAddr object. The type is
    // InetAddressType.
    Cipslaudpechotmplsrcaddrtype interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    Cipslaudpechotmplsrcaddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    Cipslaudpechotmplsrcport interface{}

    // Specifies the duration to wait for an IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    Cipslaudpechotmpltimeout interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    Cipslaudpechotmplverifydata interface{}

    // This object represents the number of octets to be placed into the ARR Data
    // portion of the request message, when using SNA protocols.  For non-ARR
    // protocols' RTT request/responses, this value represents the native payload
    // size.  REMEMBER:  The ARR Header overhead is not included            in
    // this value. The type is interface{} with range: 4..1500. Units are octets.
    Cipslaudpechotmplreqdatasize interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    Cipslaudpechotmpltos interface{}

    // This field is used to specify the VRF name with which the IP SLA operation
    // will be used. For regular IP SLA operation this field should not be
    // configured. The agent will use this field to identify the VRF routing Table
    // for this operation. The type is string with length: 0..32.
    Cipslaudpechotmplvrfname interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaUdpEchoTmplHistFilter is  satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    Cipslaudpechotmplthreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    Cipslaudpechotmplhistlives interface{}

    // The maximum number of history buckets to record. This value should be set
    // to the number of operations  to keep per lifetime.  After
    // cipslaUdpEchoTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaUdpEchoTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    Cipslaudpechotmplhistbuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaUdpEchoTmplThreshold are                     recorded. failures(4)   
    // - the results of failed operations (only)                     are recorded.
    // The type is Cipslaudpechotmplhistfilter.
    Cipslaudpechotmplhistfilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    Cipslaudpechotmplstatshours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaUdpEchoTmplStatsNumDistBuckets will be kept.  The last
    // cipslaUdpEchoTmplStatsNumDistBuckets will contain all entries from its
    // distribution interval start point to infinity. The type is interface{} with
    // range: 1..20.
    Cipslaudpechotmpldistbuckets interface{}

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
    Cipslaudpechotmpldistinterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslaudpechotmplstoragetype interface{}

    // The status of the conceptual UDP echo template control row. When the status
    // is active, all the read-create objects in  that row can be modified. The
    // type is RowStatus.
    Cipslaudpechotmplrowstatus interface{}
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetEntityData() *types.CommonEntityData {
    cipslaudpechotmplentry.EntityData.YFilter = cipslaudpechotmplentry.YFilter
    cipslaudpechotmplentry.EntityData.YangName = "cipslaUdpEchoTmplEntry"
    cipslaudpechotmplentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaudpechotmplentry.EntityData.ParentYangName = "cipslaUdpEchoTmplTable"
    cipslaudpechotmplentry.EntityData.SegmentPath = "cipslaUdpEchoTmplEntry" + "[cipslaUdpEchoTmplName='" + fmt.Sprintf("%v", cipslaudpechotmplentry.Cipslaudpechotmplname) + "']"
    cipslaudpechotmplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaudpechotmplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaudpechotmplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaudpechotmplentry.EntityData.Children = make(map[string]types.YChild)
    cipslaudpechotmplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplName"] = types.YLeaf{"Cipslaudpechotmplname", cipslaudpechotmplentry.Cipslaudpechotmplname}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplDescription"] = types.YLeaf{"Cipslaudpechotmpldescription", cipslaudpechotmplentry.Cipslaudpechotmpldescription}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplControlEnable"] = types.YLeaf{"Cipslaudpechotmplcontrolenable", cipslaudpechotmplentry.Cipslaudpechotmplcontrolenable}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplSrcAddrType"] = types.YLeaf{"Cipslaudpechotmplsrcaddrtype", cipslaudpechotmplentry.Cipslaudpechotmplsrcaddrtype}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplSrcAddr"] = types.YLeaf{"Cipslaudpechotmplsrcaddr", cipslaudpechotmplentry.Cipslaudpechotmplsrcaddr}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplSrcPort"] = types.YLeaf{"Cipslaudpechotmplsrcport", cipslaudpechotmplentry.Cipslaudpechotmplsrcport}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplTimeOut"] = types.YLeaf{"Cipslaudpechotmpltimeout", cipslaudpechotmplentry.Cipslaudpechotmpltimeout}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplVerifyData"] = types.YLeaf{"Cipslaudpechotmplverifydata", cipslaudpechotmplentry.Cipslaudpechotmplverifydata}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplReqDataSize"] = types.YLeaf{"Cipslaudpechotmplreqdatasize", cipslaudpechotmplentry.Cipslaudpechotmplreqdatasize}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplTOS"] = types.YLeaf{"Cipslaudpechotmpltos", cipslaudpechotmplentry.Cipslaudpechotmpltos}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplVrfName"] = types.YLeaf{"Cipslaudpechotmplvrfname", cipslaudpechotmplentry.Cipslaudpechotmplvrfname}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplThreshold"] = types.YLeaf{"Cipslaudpechotmplthreshold", cipslaudpechotmplentry.Cipslaudpechotmplthreshold}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplHistLives"] = types.YLeaf{"Cipslaudpechotmplhistlives", cipslaudpechotmplentry.Cipslaudpechotmplhistlives}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplHistBuckets"] = types.YLeaf{"Cipslaudpechotmplhistbuckets", cipslaudpechotmplentry.Cipslaudpechotmplhistbuckets}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplHistFilter"] = types.YLeaf{"Cipslaudpechotmplhistfilter", cipslaudpechotmplentry.Cipslaudpechotmplhistfilter}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplStatsHours"] = types.YLeaf{"Cipslaudpechotmplstatshours", cipslaudpechotmplentry.Cipslaudpechotmplstatshours}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplDistBuckets"] = types.YLeaf{"Cipslaudpechotmpldistbuckets", cipslaudpechotmplentry.Cipslaudpechotmpldistbuckets}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplDistInterval"] = types.YLeaf{"Cipslaudpechotmpldistinterval", cipslaudpechotmplentry.Cipslaudpechotmpldistinterval}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplStorageType"] = types.YLeaf{"Cipslaudpechotmplstoragetype", cipslaudpechotmplentry.Cipslaudpechotmplstoragetype}
    cipslaudpechotmplentry.EntityData.Leafs["cipslaUdpEchoTmplRowStatus"] = types.YLeaf{"Cipslaudpechotmplrowstatus", cipslaudpechotmplentry.Cipslaudpechotmplrowstatus}
    return &(cipslaudpechotmplentry.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter represents                    are recorded.
type CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter string

const (
    CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter_none CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter = "none"

    CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter_all CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter = "all"

    CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter_overThreshold CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter = "overThreshold"

    CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter_failures CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry_Cipslaudpechotmplhistfilter = "failures"
)

// CISCOIPSLAECHOMIB_Cipslatcpconntmpltable
// A table that contains TCP connect template specific definitions.
type CISCOIPSLAECHOMIB_Cipslatcpconntmpltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA TCP connect template. The type is slice
    // of CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry.
    Cipslatcpconntmplentry []CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetEntityData() *types.CommonEntityData {
    cipslatcpconntmpltable.EntityData.YFilter = cipslatcpconntmpltable.YFilter
    cipslatcpconntmpltable.EntityData.YangName = "cipslaTcpConnTmplTable"
    cipslatcpconntmpltable.EntityData.BundleName = "cisco_ios_xe"
    cipslatcpconntmpltable.EntityData.ParentYangName = "CISCO-IPSLA-ECHO-MIB"
    cipslatcpconntmpltable.EntityData.SegmentPath = "cipslaTcpConnTmplTable"
    cipslatcpconntmpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslatcpconntmpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslatcpconntmpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslatcpconntmpltable.EntityData.Children = make(map[string]types.YChild)
    cipslatcpconntmpltable.EntityData.Children["cipslaTcpConnTmplEntry"] = types.YChild{"Cipslatcpconntmplentry", nil}
    for i := range cipslatcpconntmpltable.Cipslatcpconntmplentry {
        cipslatcpconntmpltable.EntityData.Children[types.GetSegmentPath(&cipslatcpconntmpltable.Cipslatcpconntmplentry[i])] = types.YChild{"Cipslatcpconntmplentry", &cipslatcpconntmpltable.Cipslatcpconntmplentry[i]}
    }
    cipslatcpconntmpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslatcpconntmpltable.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry
// A row entry representing an IPSLA TCP connect template.
type CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string which specifies the TCP connect template
    // name. The type is string with length: 1..64.
    Cipslatcpconntmplname interface{}

    // A string which provides description for the TCP connect template. The type
    // is string with length: 0..128.
    Cipslatcpconntmpldescription interface{}

    // If this object is enabled, then the IP SLA application will send control
    // messages to a responder, residing on the target router to respond to the
    // data request packets being sent by the source router. The type is bool.
    Cipslatcpconntmplcontrolenable interface{}

    // An enumerated value which specifies the IP address type of the source. It
    // must be used along with the cipslaTcpConnTmplSrcAddr object. The type is
    // InetAddressType.
    Cipslatcpconntmplsrcaddrtype interface{}

    // A string which specifies the IP address of the source. The type is string
    // with length: 0..255.
    Cipslatcpconntmplsrcaddr interface{}

    // This object represents the source's port number. If this object is not
    // specified, the application will get a port allocated by the system. The
    // type is interface{} with range: 0..65535.
    Cipslatcpconntmplsrcport interface{}

    // Specifies the duration to wait for an IP SLA operation completion.  For
    // connection oriented protocols, this may cause the connection to be closed
    // by the operation.  Once closed, it will be assumed that the connection
    // reestablishment will be performed.  To prevent unwanted closure of
    // connections, be sure to set this value to a realistic connection timeout.
    // The type is interface{} with range: 0..604800000. Units are milliseconds.
    Cipslatcpconntmpltimeout interface{}

    // When set to true, the resulting data in each IP SLA operation is compared
    // with the expected data.  This includes checking header information (if
    // possible) and exact packet size. The type is bool.
    Cipslatcpconntmplverifydata interface{}

    // This object represents the type of service octet in an IP header. The type
    // is interface{} with range: 0..255.
    Cipslatcpconntmpltos interface{}

    // This object defines an administrative threshold limit. If the IP SLA
    // operation time exceeds this limit and if the condition specified in
    // cipslaTcpConnTmplHistFilter is  satisfied, one threshold crossing
    // occurrence will be counted. The type is interface{} with range:
    // 0..2147483647. Units are milliseconds.
    Cipslatcpconntmplthreshold interface{}

    // The maximum number of history lives to record.  A life is defined by the
    // countdown (or transition) to zero  by the cipslaAutoGroupScheduleLife
    // object.  A new life is created when the same conceptual control row is
    // restarted via the transition of the  cipslaAutoGroupScheduleLife object and
    // its subsequent  countdown.  The value of zero will shut off all data
    // collection. The type is interface{} with range: 0..2.
    Cipslatcpconntmplhistlives interface{}

    // The maximum number of history buckets to record. This value should be set
    // to the number of operations  to keep per lifetime.  After
    // cipslaTcpConnTmplHistBuckets are filled, the  oldest entries are deleted
    // and the most recent cipslaTcpConnTmplHistBuckets buckets are retained. The
    // type is interface{} with range: 1..60.
    Cipslatcpconntmplhistbuckets interface{}

    // Defines a filter for adding RTT results to the history buffer:  none(1)    
    // - no history is recorded all(2)           - the results of all completion
    // times                     and failed completions are recorded
    // overThreshold(3) - the results of completion times                    over
    // cipslaTcpConnTmplThreshold are                     recorded. failures(4)   
    // - the results of failed operations (only)                     are recorded.
    // The type is Cipslatcpconntmplhistfilter.
    Cipslatcpconntmplhistfilter interface{}

    // The maximum number of hours for which statistics are maintained.
    // Specifically this is the number of hourly  groups to keep before rolling
    // over.  The value of one is not advisable because the hourly group will
    // close and immediately be deleted before the network management station will
    // have the opportunity to retrieve the statistics.  The value of zero will
    // shut off data collection. The type is interface{} with range: 0..25. Units
    // are hours.
    Cipslatcpconntmplstatshours interface{}

    // The maximum number of statistical distribution buckets to accumulate. 
    // Since this index does not rollover, only the first
    // cipslaTcpConnTmplDistBuckets will be kept.  The last
    // cipslaTcpConnTmplDistBuckets will contain all entries from its distribution
    // interval start point to infinity. The type is interface{} with range:
    // 1..20.
    Cipslatcpconntmpldistbuckets interface{}

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
    Cipslatcpconntmpldistinterval interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslatcpconntmplstoragetype interface{}

    // The status of the conceptual tcp connect control row. When the status is
    // active, all the read-create objects  in that row can be modified. The type
    // is RowStatus.
    Cipslatcpconntmplrowstatus interface{}
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetEntityData() *types.CommonEntityData {
    cipslatcpconntmplentry.EntityData.YFilter = cipslatcpconntmplentry.YFilter
    cipslatcpconntmplentry.EntityData.YangName = "cipslaTcpConnTmplEntry"
    cipslatcpconntmplentry.EntityData.BundleName = "cisco_ios_xe"
    cipslatcpconntmplentry.EntityData.ParentYangName = "cipslaTcpConnTmplTable"
    cipslatcpconntmplentry.EntityData.SegmentPath = "cipslaTcpConnTmplEntry" + "[cipslaTcpConnTmplName='" + fmt.Sprintf("%v", cipslatcpconntmplentry.Cipslatcpconntmplname) + "']"
    cipslatcpconntmplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslatcpconntmplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslatcpconntmplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslatcpconntmplentry.EntityData.Children = make(map[string]types.YChild)
    cipslatcpconntmplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplName"] = types.YLeaf{"Cipslatcpconntmplname", cipslatcpconntmplentry.Cipslatcpconntmplname}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplDescription"] = types.YLeaf{"Cipslatcpconntmpldescription", cipslatcpconntmplentry.Cipslatcpconntmpldescription}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplControlEnable"] = types.YLeaf{"Cipslatcpconntmplcontrolenable", cipslatcpconntmplentry.Cipslatcpconntmplcontrolenable}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplSrcAddrType"] = types.YLeaf{"Cipslatcpconntmplsrcaddrtype", cipslatcpconntmplentry.Cipslatcpconntmplsrcaddrtype}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplSrcAddr"] = types.YLeaf{"Cipslatcpconntmplsrcaddr", cipslatcpconntmplentry.Cipslatcpconntmplsrcaddr}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplSrcPort"] = types.YLeaf{"Cipslatcpconntmplsrcport", cipslatcpconntmplentry.Cipslatcpconntmplsrcport}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplTimeOut"] = types.YLeaf{"Cipslatcpconntmpltimeout", cipslatcpconntmplentry.Cipslatcpconntmpltimeout}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplVerifyData"] = types.YLeaf{"Cipslatcpconntmplverifydata", cipslatcpconntmplentry.Cipslatcpconntmplverifydata}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplTOS"] = types.YLeaf{"Cipslatcpconntmpltos", cipslatcpconntmplentry.Cipslatcpconntmpltos}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplThreshold"] = types.YLeaf{"Cipslatcpconntmplthreshold", cipslatcpconntmplentry.Cipslatcpconntmplthreshold}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplHistLives"] = types.YLeaf{"Cipslatcpconntmplhistlives", cipslatcpconntmplentry.Cipslatcpconntmplhistlives}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplHistBuckets"] = types.YLeaf{"Cipslatcpconntmplhistbuckets", cipslatcpconntmplentry.Cipslatcpconntmplhistbuckets}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplHistFilter"] = types.YLeaf{"Cipslatcpconntmplhistfilter", cipslatcpconntmplentry.Cipslatcpconntmplhistfilter}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplStatsHours"] = types.YLeaf{"Cipslatcpconntmplstatshours", cipslatcpconntmplentry.Cipslatcpconntmplstatshours}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplDistBuckets"] = types.YLeaf{"Cipslatcpconntmpldistbuckets", cipslatcpconntmplentry.Cipslatcpconntmpldistbuckets}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplDistInterval"] = types.YLeaf{"Cipslatcpconntmpldistinterval", cipslatcpconntmplentry.Cipslatcpconntmpldistinterval}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplStorageType"] = types.YLeaf{"Cipslatcpconntmplstoragetype", cipslatcpconntmplentry.Cipslatcpconntmplstoragetype}
    cipslatcpconntmplentry.EntityData.Leafs["cipslaTcpConnTmplRowStatus"] = types.YLeaf{"Cipslatcpconntmplrowstatus", cipslatcpconntmplentry.Cipslatcpconntmplrowstatus}
    return &(cipslatcpconntmplentry.EntityData)
}

// CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter represents                    are recorded.
type CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter string

const (
    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_none CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "none"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_all CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "all"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_overThreshold CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "overThreshold"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_failures CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "failures"
)

