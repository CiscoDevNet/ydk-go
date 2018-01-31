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
    parent types.Entity
    YFilter yfilter.YFilter

    // A table that contains ICMP echo template definitions.
    Cipslaicmpechotmpltable CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable

    // A table that contains UDP echo template specific definitions.
    Cipslaudpechotmpltable CISCOIPSLAECHOMIB_Cipslaudpechotmpltable

    // A table that contains TCP connect template specific definitions.
    Cipslatcpconntmpltable CISCOIPSLAECHOMIB_Cipslatcpconntmpltable
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetFilter() yfilter.YFilter { return cISCOIPSLAECHOMIB.YFilter }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSLAECHOMIB.YFilter = yf }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetGoName(yname string) string {
    if yname == "cipslaIcmpEchoTmplTable" { return "Cipslaicmpechotmpltable" }
    if yname == "cipslaUdpEchoTmplTable" { return "Cipslaudpechotmpltable" }
    if yname == "cipslaTcpConnTmplTable" { return "Cipslatcpconntmpltable" }
    return ""
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetSegmentPath() string {
    return "CISCO-IPSLA-ECHO-MIB:CISCO-IPSLA-ECHO-MIB"
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaIcmpEchoTmplTable" {
        return &cISCOIPSLAECHOMIB.Cipslaicmpechotmpltable
    }
    if childYangName == "cipslaUdpEchoTmplTable" {
        return &cISCOIPSLAECHOMIB.Cipslaudpechotmpltable
    }
    if childYangName == "cipslaTcpConnTmplTable" {
        return &cISCOIPSLAECHOMIB.Cipslatcpconntmpltable
    }
    return nil
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipslaIcmpEchoTmplTable"] = &cISCOIPSLAECHOMIB.Cipslaicmpechotmpltable
    children["cipslaUdpEchoTmplTable"] = &cISCOIPSLAECHOMIB.Cipslaudpechotmpltable
    children["cipslaTcpConnTmplTable"] = &cISCOIPSLAECHOMIB.Cipslatcpconntmpltable
    return children
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetYangName() string { return "CISCO-IPSLA-ECHO-MIB" }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) SetParent(parent types.Entity) { cISCOIPSLAECHOMIB.parent = parent }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetParent() types.Entity { return cISCOIPSLAECHOMIB.parent }

func (cISCOIPSLAECHOMIB *CISCOIPSLAECHOMIB) GetParentYangName() string { return "CISCO-IPSLA-ECHO-MIB" }

// CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable
// A table that contains ICMP echo template definitions.
type CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA ICMP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry.
    Cipslaicmpechotmplentry []CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetFilter() yfilter.YFilter { return cipslaicmpechotmpltable.YFilter }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) SetFilter(yf yfilter.YFilter) { cipslaicmpechotmpltable.YFilter = yf }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetGoName(yname string) string {
    if yname == "cipslaIcmpEchoTmplEntry" { return "Cipslaicmpechotmplentry" }
    return ""
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetSegmentPath() string {
    return "cipslaIcmpEchoTmplTable"
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaIcmpEchoTmplEntry" {
        for _, c := range cipslaicmpechotmpltable.Cipslaicmpechotmplentry {
            if cipslaicmpechotmpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry{}
        cipslaicmpechotmpltable.Cipslaicmpechotmplentry = append(cipslaicmpechotmpltable.Cipslaicmpechotmplentry, child)
        return &cipslaicmpechotmpltable.Cipslaicmpechotmplentry[len(cipslaicmpechotmpltable.Cipslaicmpechotmplentry)-1]
    }
    return nil
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaicmpechotmpltable.Cipslaicmpechotmplentry {
        children[cipslaicmpechotmpltable.Cipslaicmpechotmplentry[i].GetSegmentPath()] = &cipslaicmpechotmpltable.Cipslaicmpechotmplentry[i]
    }
    return children
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetYangName() string { return "cipslaIcmpEchoTmplTable" }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) SetParent(parent types.Entity) { cipslaicmpechotmpltable.parent = parent }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetParent() types.Entity { return cipslaicmpechotmpltable.parent }

func (cipslaicmpechotmpltable *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable) GetParentYangName() string { return "CISCO-IPSLA-ECHO-MIB" }

// CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry
// A row entry representing an IPSLA ICMP echo template.
type CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry struct {
    parent types.Entity
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

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetFilter() yfilter.YFilter { return cipslaicmpechotmplentry.YFilter }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) SetFilter(yf yfilter.YFilter) { cipslaicmpechotmplentry.YFilter = yf }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetGoName(yname string) string {
    if yname == "cipslaIcmpEchoTmplName" { return "Cipslaicmpechotmplname" }
    if yname == "cipslaIcmpEchoTmplDescription" { return "Cipslaicmpechotmpldescription" }
    if yname == "cipslaIcmpEchoTmplSrcAddrType" { return "Cipslaicmpechotmplsrcaddrtype" }
    if yname == "cipslaIcmpEchoTmplSrcAddr" { return "Cipslaicmpechotmplsrcaddr" }
    if yname == "cipslaIcmpEchoTmplTimeOut" { return "Cipslaicmpechotmpltimeout" }
    if yname == "cipslaIcmpEchoTmplVerifyData" { return "Cipslaicmpechotmplverifydata" }
    if yname == "cipslaIcmpEchoTmplReqDataSize" { return "Cipslaicmpechotmplreqdatasize" }
    if yname == "cipslaIcmpEchoTmplTOS" { return "Cipslaicmpechotmpltos" }
    if yname == "cipslaIcmpEchoTmplVrfName" { return "Cipslaicmpechotmplvrfname" }
    if yname == "cipslaIcmpEchoTmplThreshold" { return "Cipslaicmpechotmplthreshold" }
    if yname == "cipslaIcmpEchoTmplHistLives" { return "Cipslaicmpechotmplhistlives" }
    if yname == "cipslaIcmpEchoTmplHistBuckets" { return "Cipslaicmpechotmplhistbuckets" }
    if yname == "cipslaIcmpEchoTmplHistFilter" { return "Cipslaicmpechotmplhistfilter" }
    if yname == "cipslaIcmpEchoTmplStatsHours" { return "Cipslaicmpechotmplstatshours" }
    if yname == "cipslaIcmpEchoTmplDistBuckets" { return "Cipslaicmpechotmpldistbuckets" }
    if yname == "cipslaIcmpEchoTmplDistInterval" { return "Cipslaicmpechotmpldistinterval" }
    if yname == "cipslaIcmpEchoTmplStorageType" { return "Cipslaicmpechotmplstoragetype" }
    if yname == "cipslaIcmpEchoTmplRowStatus" { return "Cipslaicmpechotmplrowstatus" }
    return ""
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetSegmentPath() string {
    return "cipslaIcmpEchoTmplEntry" + "[cipslaIcmpEchoTmplName='" + fmt.Sprintf("%v", cipslaicmpechotmplentry.Cipslaicmpechotmplname) + "']"
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaIcmpEchoTmplName"] = cipslaicmpechotmplentry.Cipslaicmpechotmplname
    leafs["cipslaIcmpEchoTmplDescription"] = cipslaicmpechotmplentry.Cipslaicmpechotmpldescription
    leafs["cipslaIcmpEchoTmplSrcAddrType"] = cipslaicmpechotmplentry.Cipslaicmpechotmplsrcaddrtype
    leafs["cipslaIcmpEchoTmplSrcAddr"] = cipslaicmpechotmplentry.Cipslaicmpechotmplsrcaddr
    leafs["cipslaIcmpEchoTmplTimeOut"] = cipslaicmpechotmplentry.Cipslaicmpechotmpltimeout
    leafs["cipslaIcmpEchoTmplVerifyData"] = cipslaicmpechotmplentry.Cipslaicmpechotmplverifydata
    leafs["cipslaIcmpEchoTmplReqDataSize"] = cipslaicmpechotmplentry.Cipslaicmpechotmplreqdatasize
    leafs["cipslaIcmpEchoTmplTOS"] = cipslaicmpechotmplentry.Cipslaicmpechotmpltos
    leafs["cipslaIcmpEchoTmplVrfName"] = cipslaicmpechotmplentry.Cipslaicmpechotmplvrfname
    leafs["cipslaIcmpEchoTmplThreshold"] = cipslaicmpechotmplentry.Cipslaicmpechotmplthreshold
    leafs["cipslaIcmpEchoTmplHistLives"] = cipslaicmpechotmplentry.Cipslaicmpechotmplhistlives
    leafs["cipslaIcmpEchoTmplHistBuckets"] = cipslaicmpechotmplentry.Cipslaicmpechotmplhistbuckets
    leafs["cipslaIcmpEchoTmplHistFilter"] = cipslaicmpechotmplentry.Cipslaicmpechotmplhistfilter
    leafs["cipslaIcmpEchoTmplStatsHours"] = cipslaicmpechotmplentry.Cipslaicmpechotmplstatshours
    leafs["cipslaIcmpEchoTmplDistBuckets"] = cipslaicmpechotmplentry.Cipslaicmpechotmpldistbuckets
    leafs["cipslaIcmpEchoTmplDistInterval"] = cipslaicmpechotmplentry.Cipslaicmpechotmpldistinterval
    leafs["cipslaIcmpEchoTmplStorageType"] = cipslaicmpechotmplentry.Cipslaicmpechotmplstoragetype
    leafs["cipslaIcmpEchoTmplRowStatus"] = cipslaicmpechotmplentry.Cipslaicmpechotmplrowstatus
    return leafs
}

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetYangName() string { return "cipslaIcmpEchoTmplEntry" }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) SetParent(parent types.Entity) { cipslaicmpechotmplentry.parent = parent }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetParent() types.Entity { return cipslaicmpechotmplentry.parent }

func (cipslaicmpechotmplentry *CISCOIPSLAECHOMIB_Cipslaicmpechotmpltable_Cipslaicmpechotmplentry) GetParentYangName() string { return "cipslaIcmpEchoTmplTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA UDP echo template. The type is slice of
    // CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry.
    Cipslaudpechotmplentry []CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetFilter() yfilter.YFilter { return cipslaudpechotmpltable.YFilter }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) SetFilter(yf yfilter.YFilter) { cipslaudpechotmpltable.YFilter = yf }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetGoName(yname string) string {
    if yname == "cipslaUdpEchoTmplEntry" { return "Cipslaudpechotmplentry" }
    return ""
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetSegmentPath() string {
    return "cipslaUdpEchoTmplTable"
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaUdpEchoTmplEntry" {
        for _, c := range cipslaudpechotmpltable.Cipslaudpechotmplentry {
            if cipslaudpechotmpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry{}
        cipslaudpechotmpltable.Cipslaudpechotmplentry = append(cipslaudpechotmpltable.Cipslaudpechotmplentry, child)
        return &cipslaudpechotmpltable.Cipslaudpechotmplentry[len(cipslaudpechotmpltable.Cipslaudpechotmplentry)-1]
    }
    return nil
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaudpechotmpltable.Cipslaudpechotmplentry {
        children[cipslaudpechotmpltable.Cipslaudpechotmplentry[i].GetSegmentPath()] = &cipslaudpechotmpltable.Cipslaudpechotmplentry[i]
    }
    return children
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetYangName() string { return "cipslaUdpEchoTmplTable" }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) SetParent(parent types.Entity) { cipslaudpechotmpltable.parent = parent }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetParent() types.Entity { return cipslaudpechotmpltable.parent }

func (cipslaudpechotmpltable *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable) GetParentYangName() string { return "CISCO-IPSLA-ECHO-MIB" }

// CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry
// A row entry representing an IPSLA UDP echo template.
type CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry struct {
    parent types.Entity
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

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetFilter() yfilter.YFilter { return cipslaudpechotmplentry.YFilter }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) SetFilter(yf yfilter.YFilter) { cipslaudpechotmplentry.YFilter = yf }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetGoName(yname string) string {
    if yname == "cipslaUdpEchoTmplName" { return "Cipslaudpechotmplname" }
    if yname == "cipslaUdpEchoTmplDescription" { return "Cipslaudpechotmpldescription" }
    if yname == "cipslaUdpEchoTmplControlEnable" { return "Cipslaudpechotmplcontrolenable" }
    if yname == "cipslaUdpEchoTmplSrcAddrType" { return "Cipslaudpechotmplsrcaddrtype" }
    if yname == "cipslaUdpEchoTmplSrcAddr" { return "Cipslaudpechotmplsrcaddr" }
    if yname == "cipslaUdpEchoTmplSrcPort" { return "Cipslaudpechotmplsrcport" }
    if yname == "cipslaUdpEchoTmplTimeOut" { return "Cipslaudpechotmpltimeout" }
    if yname == "cipslaUdpEchoTmplVerifyData" { return "Cipslaudpechotmplverifydata" }
    if yname == "cipslaUdpEchoTmplReqDataSize" { return "Cipslaudpechotmplreqdatasize" }
    if yname == "cipslaUdpEchoTmplTOS" { return "Cipslaudpechotmpltos" }
    if yname == "cipslaUdpEchoTmplVrfName" { return "Cipslaudpechotmplvrfname" }
    if yname == "cipslaUdpEchoTmplThreshold" { return "Cipslaudpechotmplthreshold" }
    if yname == "cipslaUdpEchoTmplHistLives" { return "Cipslaudpechotmplhistlives" }
    if yname == "cipslaUdpEchoTmplHistBuckets" { return "Cipslaudpechotmplhistbuckets" }
    if yname == "cipslaUdpEchoTmplHistFilter" { return "Cipslaudpechotmplhistfilter" }
    if yname == "cipslaUdpEchoTmplStatsHours" { return "Cipslaudpechotmplstatshours" }
    if yname == "cipslaUdpEchoTmplDistBuckets" { return "Cipslaudpechotmpldistbuckets" }
    if yname == "cipslaUdpEchoTmplDistInterval" { return "Cipslaudpechotmpldistinterval" }
    if yname == "cipslaUdpEchoTmplStorageType" { return "Cipslaudpechotmplstoragetype" }
    if yname == "cipslaUdpEchoTmplRowStatus" { return "Cipslaudpechotmplrowstatus" }
    return ""
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetSegmentPath() string {
    return "cipslaUdpEchoTmplEntry" + "[cipslaUdpEchoTmplName='" + fmt.Sprintf("%v", cipslaudpechotmplentry.Cipslaudpechotmplname) + "']"
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaUdpEchoTmplName"] = cipslaudpechotmplentry.Cipslaudpechotmplname
    leafs["cipslaUdpEchoTmplDescription"] = cipslaudpechotmplentry.Cipslaudpechotmpldescription
    leafs["cipslaUdpEchoTmplControlEnable"] = cipslaudpechotmplentry.Cipslaudpechotmplcontrolenable
    leafs["cipslaUdpEchoTmplSrcAddrType"] = cipslaudpechotmplentry.Cipslaudpechotmplsrcaddrtype
    leafs["cipslaUdpEchoTmplSrcAddr"] = cipslaudpechotmplentry.Cipslaudpechotmplsrcaddr
    leafs["cipslaUdpEchoTmplSrcPort"] = cipslaudpechotmplentry.Cipslaudpechotmplsrcport
    leafs["cipslaUdpEchoTmplTimeOut"] = cipslaudpechotmplentry.Cipslaudpechotmpltimeout
    leafs["cipslaUdpEchoTmplVerifyData"] = cipslaudpechotmplentry.Cipslaudpechotmplverifydata
    leafs["cipslaUdpEchoTmplReqDataSize"] = cipslaudpechotmplentry.Cipslaudpechotmplreqdatasize
    leafs["cipslaUdpEchoTmplTOS"] = cipslaudpechotmplentry.Cipslaudpechotmpltos
    leafs["cipslaUdpEchoTmplVrfName"] = cipslaudpechotmplentry.Cipslaudpechotmplvrfname
    leafs["cipslaUdpEchoTmplThreshold"] = cipslaudpechotmplentry.Cipslaudpechotmplthreshold
    leafs["cipslaUdpEchoTmplHistLives"] = cipslaudpechotmplentry.Cipslaudpechotmplhistlives
    leafs["cipslaUdpEchoTmplHistBuckets"] = cipslaudpechotmplentry.Cipslaudpechotmplhistbuckets
    leafs["cipslaUdpEchoTmplHistFilter"] = cipslaudpechotmplentry.Cipslaudpechotmplhistfilter
    leafs["cipslaUdpEchoTmplStatsHours"] = cipslaudpechotmplentry.Cipslaudpechotmplstatshours
    leafs["cipslaUdpEchoTmplDistBuckets"] = cipslaudpechotmplentry.Cipslaudpechotmpldistbuckets
    leafs["cipslaUdpEchoTmplDistInterval"] = cipslaudpechotmplentry.Cipslaudpechotmpldistinterval
    leafs["cipslaUdpEchoTmplStorageType"] = cipslaudpechotmplentry.Cipslaudpechotmplstoragetype
    leafs["cipslaUdpEchoTmplRowStatus"] = cipslaudpechotmplentry.Cipslaudpechotmplrowstatus
    return leafs
}

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetYangName() string { return "cipslaUdpEchoTmplEntry" }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) SetParent(parent types.Entity) { cipslaudpechotmplentry.parent = parent }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetParent() types.Entity { return cipslaudpechotmplentry.parent }

func (cipslaudpechotmplentry *CISCOIPSLAECHOMIB_Cipslaudpechotmpltable_Cipslaudpechotmplentry) GetParentYangName() string { return "cipslaUdpEchoTmplTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A row entry representing an IPSLA TCP connect template. The type is slice
    // of CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry.
    Cipslatcpconntmplentry []CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetFilter() yfilter.YFilter { return cipslatcpconntmpltable.YFilter }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) SetFilter(yf yfilter.YFilter) { cipslatcpconntmpltable.YFilter = yf }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetGoName(yname string) string {
    if yname == "cipslaTcpConnTmplEntry" { return "Cipslatcpconntmplentry" }
    return ""
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetSegmentPath() string {
    return "cipslaTcpConnTmplTable"
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaTcpConnTmplEntry" {
        for _, c := range cipslatcpconntmpltable.Cipslatcpconntmplentry {
            if cipslatcpconntmpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry{}
        cipslatcpconntmpltable.Cipslatcpconntmplentry = append(cipslatcpconntmpltable.Cipslatcpconntmplentry, child)
        return &cipslatcpconntmpltable.Cipslatcpconntmplentry[len(cipslatcpconntmpltable.Cipslatcpconntmplentry)-1]
    }
    return nil
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslatcpconntmpltable.Cipslatcpconntmplentry {
        children[cipslatcpconntmpltable.Cipslatcpconntmplentry[i].GetSegmentPath()] = &cipslatcpconntmpltable.Cipslatcpconntmplentry[i]
    }
    return children
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetYangName() string { return "cipslaTcpConnTmplTable" }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) SetParent(parent types.Entity) { cipslatcpconntmpltable.parent = parent }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetParent() types.Entity { return cipslatcpconntmpltable.parent }

func (cipslatcpconntmpltable *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable) GetParentYangName() string { return "CISCO-IPSLA-ECHO-MIB" }

// CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry
// A row entry representing an IPSLA TCP connect template.
type CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry struct {
    parent types.Entity
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

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetFilter() yfilter.YFilter { return cipslatcpconntmplentry.YFilter }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) SetFilter(yf yfilter.YFilter) { cipslatcpconntmplentry.YFilter = yf }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetGoName(yname string) string {
    if yname == "cipslaTcpConnTmplName" { return "Cipslatcpconntmplname" }
    if yname == "cipslaTcpConnTmplDescription" { return "Cipslatcpconntmpldescription" }
    if yname == "cipslaTcpConnTmplControlEnable" { return "Cipslatcpconntmplcontrolenable" }
    if yname == "cipslaTcpConnTmplSrcAddrType" { return "Cipslatcpconntmplsrcaddrtype" }
    if yname == "cipslaTcpConnTmplSrcAddr" { return "Cipslatcpconntmplsrcaddr" }
    if yname == "cipslaTcpConnTmplSrcPort" { return "Cipslatcpconntmplsrcport" }
    if yname == "cipslaTcpConnTmplTimeOut" { return "Cipslatcpconntmpltimeout" }
    if yname == "cipslaTcpConnTmplVerifyData" { return "Cipslatcpconntmplverifydata" }
    if yname == "cipslaTcpConnTmplTOS" { return "Cipslatcpconntmpltos" }
    if yname == "cipslaTcpConnTmplThreshold" { return "Cipslatcpconntmplthreshold" }
    if yname == "cipslaTcpConnTmplHistLives" { return "Cipslatcpconntmplhistlives" }
    if yname == "cipslaTcpConnTmplHistBuckets" { return "Cipslatcpconntmplhistbuckets" }
    if yname == "cipslaTcpConnTmplHistFilter" { return "Cipslatcpconntmplhistfilter" }
    if yname == "cipslaTcpConnTmplStatsHours" { return "Cipslatcpconntmplstatshours" }
    if yname == "cipslaTcpConnTmplDistBuckets" { return "Cipslatcpconntmpldistbuckets" }
    if yname == "cipslaTcpConnTmplDistInterval" { return "Cipslatcpconntmpldistinterval" }
    if yname == "cipslaTcpConnTmplStorageType" { return "Cipslatcpconntmplstoragetype" }
    if yname == "cipslaTcpConnTmplRowStatus" { return "Cipslatcpconntmplrowstatus" }
    return ""
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetSegmentPath() string {
    return "cipslaTcpConnTmplEntry" + "[cipslaTcpConnTmplName='" + fmt.Sprintf("%v", cipslatcpconntmplentry.Cipslatcpconntmplname) + "']"
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaTcpConnTmplName"] = cipslatcpconntmplentry.Cipslatcpconntmplname
    leafs["cipslaTcpConnTmplDescription"] = cipslatcpconntmplentry.Cipslatcpconntmpldescription
    leafs["cipslaTcpConnTmplControlEnable"] = cipslatcpconntmplentry.Cipslatcpconntmplcontrolenable
    leafs["cipslaTcpConnTmplSrcAddrType"] = cipslatcpconntmplentry.Cipslatcpconntmplsrcaddrtype
    leafs["cipslaTcpConnTmplSrcAddr"] = cipslatcpconntmplentry.Cipslatcpconntmplsrcaddr
    leafs["cipslaTcpConnTmplSrcPort"] = cipslatcpconntmplentry.Cipslatcpconntmplsrcport
    leafs["cipslaTcpConnTmplTimeOut"] = cipslatcpconntmplentry.Cipslatcpconntmpltimeout
    leafs["cipslaTcpConnTmplVerifyData"] = cipslatcpconntmplentry.Cipslatcpconntmplverifydata
    leafs["cipslaTcpConnTmplTOS"] = cipslatcpconntmplentry.Cipslatcpconntmpltos
    leafs["cipslaTcpConnTmplThreshold"] = cipslatcpconntmplentry.Cipslatcpconntmplthreshold
    leafs["cipslaTcpConnTmplHistLives"] = cipslatcpconntmplentry.Cipslatcpconntmplhistlives
    leafs["cipslaTcpConnTmplHistBuckets"] = cipslatcpconntmplentry.Cipslatcpconntmplhistbuckets
    leafs["cipslaTcpConnTmplHistFilter"] = cipslatcpconntmplentry.Cipslatcpconntmplhistfilter
    leafs["cipslaTcpConnTmplStatsHours"] = cipslatcpconntmplentry.Cipslatcpconntmplstatshours
    leafs["cipslaTcpConnTmplDistBuckets"] = cipslatcpconntmplentry.Cipslatcpconntmpldistbuckets
    leafs["cipslaTcpConnTmplDistInterval"] = cipslatcpconntmplentry.Cipslatcpconntmpldistinterval
    leafs["cipslaTcpConnTmplStorageType"] = cipslatcpconntmplentry.Cipslatcpconntmplstoragetype
    leafs["cipslaTcpConnTmplRowStatus"] = cipslatcpconntmplentry.Cipslatcpconntmplrowstatus
    return leafs
}

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetYangName() string { return "cipslaTcpConnTmplEntry" }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) SetParent(parent types.Entity) { cipslatcpconntmplentry.parent = parent }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetParent() types.Entity { return cipslatcpconntmplentry.parent }

func (cipslatcpconntmplentry *CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry) GetParentYangName() string { return "cipslaTcpConnTmplTable" }

// CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter represents                    are recorded.
type CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter string

const (
    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_none CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "none"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_all CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "all"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_overThreshold CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "overThreshold"

    CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter_failures CISCOIPSLAECHOMIB_Cipslatcpconntmpltable_Cipslatcpconntmplentry_Cipslatcpconntmplhistfilter = "failures"
)

