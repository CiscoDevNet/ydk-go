// This MIB module defines MIB objects which provide
// mechanisms to remotely configure the parameters used
// by an SNMP entity for the generation of SNMP messages.
package snmp_target_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_target_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMP-TARGET-MIB SNMP-TARGET-MIB}", reflect.TypeOf(SNMPTARGETMIB{}))
    ydk.RegisterEntity("SNMP-TARGET-MIB:SNMP-TARGET-MIB", reflect.TypeOf(SNMPTARGETMIB{}))
}

// SNMPTARGETMIB
type SNMPTARGETMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Snmptargetobjects SNMPTARGETMIB_Snmptargetobjects

    // A table of transport addresses to be used in the generation of SNMP
    // messages.
    Snmptargetaddrtable SNMPTARGETMIB_Snmptargetaddrtable

    // A table of SNMP target information to be used in the generation of SNMP
    // messages.
    Snmptargetparamstable SNMPTARGETMIB_Snmptargetparamstable
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetFilter() yfilter.YFilter { return sNMPTARGETMIB.YFilter }

func (sNMPTARGETMIB *SNMPTARGETMIB) SetFilter(yf yfilter.YFilter) { sNMPTARGETMIB.YFilter = yf }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetGoName(yname string) string {
    if yname == "snmpTargetObjects" { return "Snmptargetobjects" }
    if yname == "snmpTargetAddrTable" { return "Snmptargetaddrtable" }
    if yname == "snmpTargetParamsTable" { return "Snmptargetparamstable" }
    return ""
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetSegmentPath() string {
    return "SNMP-TARGET-MIB:SNMP-TARGET-MIB"
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpTargetObjects" {
        return &sNMPTARGETMIB.Snmptargetobjects
    }
    if childYangName == "snmpTargetAddrTable" {
        return &sNMPTARGETMIB.Snmptargetaddrtable
    }
    if childYangName == "snmpTargetParamsTable" {
        return &sNMPTARGETMIB.Snmptargetparamstable
    }
    return nil
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snmpTargetObjects"] = &sNMPTARGETMIB.Snmptargetobjects
    children["snmpTargetAddrTable"] = &sNMPTARGETMIB.Snmptargetaddrtable
    children["snmpTargetParamsTable"] = &sNMPTARGETMIB.Snmptargetparamstable
    return children
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetBundleName() string { return "cisco_ios_xe" }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetYangName() string { return "SNMP-TARGET-MIB" }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sNMPTARGETMIB *SNMPTARGETMIB) SetParent(parent types.Entity) { sNMPTARGETMIB.parent = parent }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetParent() types.Entity { return sNMPTARGETMIB.parent }

func (sNMPTARGETMIB *SNMPTARGETMIB) GetParentYangName() string { return "SNMP-TARGET-MIB" }

// SNMPTARGETMIB_Snmptargetobjects
type SNMPTARGETMIB_Snmptargetobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object is used to facilitate modification of table entries in the
    // SNMP-TARGET-MIB module by multiple      managers.  In particular, it is
    // useful when modifying the value of the snmpTargetAddrTagList object.  The
    // procedure for modifying the snmpTargetAddrTagList object is as follows:    
    // 1.  Retrieve the value of snmpTargetSpinLock and         of
    // snmpTargetAddrTagList.      2.  Generate a new value for
    // snmpTargetAddrTagList.      3.  Set the value of snmpTargetSpinLock to the 
    // retrieved value, and the value of         snmpTargetAddrTagList to the new
    // value.  If         the set fails for the snmpTargetSpinLock         object,
    // go back to step 1. The type is interface{} with range: 0..2147483647.
    Snmptargetspinlock interface{}

    // The total number of packets received by the SNMP engine which were dropped
    // because the context contained in the message was unavailable. The type is
    // interface{} with range: 0..4294967295.
    Snmpunavailablecontexts interface{}

    // The total number of packets received by the SNMP engine which were dropped
    // because the context contained in the message was unknown. The type is
    // interface{} with range: 0..4294967295.
    Snmpunknowncontexts interface{}
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetFilter() yfilter.YFilter { return snmptargetobjects.YFilter }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) SetFilter(yf yfilter.YFilter) { snmptargetobjects.YFilter = yf }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetGoName(yname string) string {
    if yname == "snmpTargetSpinLock" { return "Snmptargetspinlock" }
    if yname == "snmpUnavailableContexts" { return "Snmpunavailablecontexts" }
    if yname == "snmpUnknownContexts" { return "Snmpunknowncontexts" }
    return ""
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetSegmentPath() string {
    return "snmpTargetObjects"
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpTargetSpinLock"] = snmptargetobjects.Snmptargetspinlock
    leafs["snmpUnavailableContexts"] = snmptargetobjects.Snmpunavailablecontexts
    leafs["snmpUnknownContexts"] = snmptargetobjects.Snmpunknowncontexts
    return leafs
}

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetBundleName() string { return "cisco_ios_xe" }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetYangName() string { return "snmpTargetObjects" }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) SetParent(parent types.Entity) { snmptargetobjects.parent = parent }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetParent() types.Entity { return snmptargetobjects.parent }

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetParentYangName() string { return "SNMP-TARGET-MIB" }

// SNMPTARGETMIB_Snmptargetaddrtable
// A table of transport addresses to be used in the generation
// of SNMP messages.
type SNMPTARGETMIB_Snmptargetaddrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A transport address to be used in the generation of SNMP operations. 
    // Entries in the snmpTargetAddrTable are created and deleted using the
    // snmpTargetAddrRowStatus object. The type is slice of
    // SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry.
    Snmptargetaddrentry []SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetFilter() yfilter.YFilter { return snmptargetaddrtable.YFilter }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) SetFilter(yf yfilter.YFilter) { snmptargetaddrtable.YFilter = yf }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetGoName(yname string) string {
    if yname == "snmpTargetAddrEntry" { return "Snmptargetaddrentry" }
    return ""
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetSegmentPath() string {
    return "snmpTargetAddrTable"
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpTargetAddrEntry" {
        for _, c := range snmptargetaddrtable.Snmptargetaddrentry {
            if snmptargetaddrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry{}
        snmptargetaddrtable.Snmptargetaddrentry = append(snmptargetaddrtable.Snmptargetaddrentry, child)
        return &snmptargetaddrtable.Snmptargetaddrentry[len(snmptargetaddrtable.Snmptargetaddrentry)-1]
    }
    return nil
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range snmptargetaddrtable.Snmptargetaddrentry {
        children[snmptargetaddrtable.Snmptargetaddrentry[i].GetSegmentPath()] = &snmptargetaddrtable.Snmptargetaddrentry[i]
    }
    return children
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetBundleName() string { return "cisco_ios_xe" }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetYangName() string { return "snmpTargetAddrTable" }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) SetParent(parent types.Entity) { snmptargetaddrtable.parent = parent }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetParent() types.Entity { return snmptargetaddrtable.parent }

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetParentYangName() string { return "SNMP-TARGET-MIB" }

// SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry
// A transport address to be used in the generation
// of SNMP operations.
// 
// Entries in the snmpTargetAddrTable are created and
// deleted using the snmpTargetAddrRowStatus object.
type SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpTargetAddrEntry. The type is string with length:
    // 1..32.
    Snmptargetaddrname interface{}

    // This object indicates the transport type of the address contained in the
    // snmpTargetAddrTAddress object. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Snmptargetaddrtdomain interface{}

    // This object contains a transport address.  The format of this address
    // depends on the value of the snmpTargetAddrTDomain object. The type is
    // string with length: 1..255.
    Snmptargetaddrtaddress interface{}

    // This object should reflect the expected maximum round trip time for
    // communicating with the transport address defined by this row.  When a
    // message is sent to this address, and a response (if one is expected) is not
    // received within this time period, an implementation may assume that the
    // response will not be delivered.  Note that the time interval that an
    // application waits for a response may actually be derived from the value of
    // this object.  The method for deriving the actual time interval is
    // implementation dependent.  One such method      is to derive the expected
    // round trip time based on a particular retransmission algorithm and on the
    // number of timeouts which have occurred.  The type of message may also be
    // considered when deriving expected round trip times for retransmissions. 
    // For example, if a message is being sent with a securityLevel that indicates
    // both authentication and privacy, the derived value may be increased to
    // compensate for extra processing time spent during authentication and
    // encryption processing. The type is interface{} with range: 0..2147483647.
    Snmptargetaddrtimeout interface{}

    // This object specifies a default number of retries to be attempted when a
    // response is not received for a generated message.  An application may
    // provide its own retry count, in which case the value of this object is
    // ignored. The type is interface{} with range: 0..255.
    Snmptargetaddrretrycount interface{}

    // This object contains a list of tag values which are used to select target
    // addresses for a particular operation. The type is string.
    Snmptargetaddrtaglist interface{}

    // The value of this object identifies an entry in the snmpTargetParamsTable. 
    // The identified entry contains SNMP parameters to be used when generating
    // messages to be sent to this transport address. The type is string with
    // length: 1..32.
    Snmptargetaddrparams interface{}

    // The storage type for this conceptual row. The type is StorageType.
    Snmptargetaddrstoragetype interface{}

    // The status of this conceptual row.  To create a row in this table, a
    // manager must set this object to either createAndGo(4) or createAndWait(5). 
    // Until instances of all corresponding columns are appropriately configured,
    // the value of the corresponding instance of the snmpTargetAddrRowStatus
    // column is 'notReady'.  In particular, a newly created row cannot be made
    // active until the corresponding instances of snmpTargetAddrTDomain,
    // snmpTargetAddrTAddress, and snmpTargetAddrParams have all been set.  The
    // following objects may not be modified while the value of this object is
    // active(1):     - snmpTargetAddrTDomain     - snmpTargetAddrTAddress An
    // attempt to set these objects while the value of snmpTargetAddrRowStatus is
    // active(1) will result in an inconsistentValue error. The type is RowStatus.
    Snmptargetaddrrowstatus interface{}
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetFilter() yfilter.YFilter { return snmptargetaddrentry.YFilter }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) SetFilter(yf yfilter.YFilter) { snmptargetaddrentry.YFilter = yf }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetGoName(yname string) string {
    if yname == "snmpTargetAddrName" { return "Snmptargetaddrname" }
    if yname == "snmpTargetAddrTDomain" { return "Snmptargetaddrtdomain" }
    if yname == "snmpTargetAddrTAddress" { return "Snmptargetaddrtaddress" }
    if yname == "snmpTargetAddrTimeout" { return "Snmptargetaddrtimeout" }
    if yname == "snmpTargetAddrRetryCount" { return "Snmptargetaddrretrycount" }
    if yname == "snmpTargetAddrTagList" { return "Snmptargetaddrtaglist" }
    if yname == "snmpTargetAddrParams" { return "Snmptargetaddrparams" }
    if yname == "snmpTargetAddrStorageType" { return "Snmptargetaddrstoragetype" }
    if yname == "snmpTargetAddrRowStatus" { return "Snmptargetaddrrowstatus" }
    return ""
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetSegmentPath() string {
    return "snmpTargetAddrEntry" + "[snmpTargetAddrName='" + fmt.Sprintf("%v", snmptargetaddrentry.Snmptargetaddrname) + "']"
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpTargetAddrName"] = snmptargetaddrentry.Snmptargetaddrname
    leafs["snmpTargetAddrTDomain"] = snmptargetaddrentry.Snmptargetaddrtdomain
    leafs["snmpTargetAddrTAddress"] = snmptargetaddrentry.Snmptargetaddrtaddress
    leafs["snmpTargetAddrTimeout"] = snmptargetaddrentry.Snmptargetaddrtimeout
    leafs["snmpTargetAddrRetryCount"] = snmptargetaddrentry.Snmptargetaddrretrycount
    leafs["snmpTargetAddrTagList"] = snmptargetaddrentry.Snmptargetaddrtaglist
    leafs["snmpTargetAddrParams"] = snmptargetaddrentry.Snmptargetaddrparams
    leafs["snmpTargetAddrStorageType"] = snmptargetaddrentry.Snmptargetaddrstoragetype
    leafs["snmpTargetAddrRowStatus"] = snmptargetaddrentry.Snmptargetaddrrowstatus
    return leafs
}

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetBundleName() string { return "cisco_ios_xe" }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetYangName() string { return "snmpTargetAddrEntry" }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) SetParent(parent types.Entity) { snmptargetaddrentry.parent = parent }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetParent() types.Entity { return snmptargetaddrentry.parent }

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetParentYangName() string { return "snmpTargetAddrTable" }

// SNMPTARGETMIB_Snmptargetparamstable
// A table of SNMP target information to be used
// in the generation of SNMP messages.
type SNMPTARGETMIB_Snmptargetparamstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A set of SNMP target information.  Entries in the snmpTargetParamsTable are
    // created and deleted using the snmpTargetParamsRowStatus object. The type is
    // slice of SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry.
    Snmptargetparamsentry []SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetFilter() yfilter.YFilter { return snmptargetparamstable.YFilter }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) SetFilter(yf yfilter.YFilter) { snmptargetparamstable.YFilter = yf }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetGoName(yname string) string {
    if yname == "snmpTargetParamsEntry" { return "Snmptargetparamsentry" }
    return ""
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetSegmentPath() string {
    return "snmpTargetParamsTable"
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpTargetParamsEntry" {
        for _, c := range snmptargetparamstable.Snmptargetparamsentry {
            if snmptargetparamstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry{}
        snmptargetparamstable.Snmptargetparamsentry = append(snmptargetparamstable.Snmptargetparamsentry, child)
        return &snmptargetparamstable.Snmptargetparamsentry[len(snmptargetparamstable.Snmptargetparamsentry)-1]
    }
    return nil
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range snmptargetparamstable.Snmptargetparamsentry {
        children[snmptargetparamstable.Snmptargetparamsentry[i].GetSegmentPath()] = &snmptargetparamstable.Snmptargetparamsentry[i]
    }
    return children
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetBundleName() string { return "cisco_ios_xe" }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetYangName() string { return "snmpTargetParamsTable" }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) SetParent(parent types.Entity) { snmptargetparamstable.parent = parent }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetParent() types.Entity { return snmptargetparamstable.parent }

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetParentYangName() string { return "SNMP-TARGET-MIB" }

// SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry
// A set of SNMP target information.
// 
// Entries in the snmpTargetParamsTable are created and
// deleted using the snmpTargetParamsRowStatus object.
type SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpTargetParamsEntry. The type is string with length:
    // 1..32.
    Snmptargetparamsname interface{}

    // The Message Processing Model to be used when generating SNMP messages using
    // this entry. The type is interface{} with range: 0..2147483647.
    Snmptargetparamsmpmodel interface{}

    // The Security Model to be used when generating SNMP messages using this
    // entry.  An implementation may choose to return an inconsistentValue error
    // if an attempt is made to set this variable to a value for a security model
    // which the implementation does      not support. The type is interface{}
    // with range: 1..2147483647.
    Snmptargetparamssecuritymodel interface{}

    // The securityName which identifies the Principal on whose behalf SNMP
    // messages will be generated using this entry. The type is string.
    Snmptargetparamssecurityname interface{}

    // The Level of Security to be used when generating SNMP messages using this
    // entry. The type is SnmpSecurityLevel.
    Snmptargetparamssecuritylevel interface{}

    // The storage type for this conceptual row. The type is StorageType.
    Snmptargetparamsstoragetype interface{}

    // The status of this conceptual row.  To create a row in this table, a
    // manager must set this object to either createAndGo(4) or createAndWait(5). 
    // Until instances of all corresponding columns are appropriately configured,
    // the value of the corresponding instance of the snmpTargetParamsRowStatus
    // column is 'notReady'.  In particular, a newly created row cannot be made   
    // active until the corresponding snmpTargetParamsMPModel,
    // snmpTargetParamsSecurityModel, snmpTargetParamsSecurityName, and
    // snmpTargetParamsSecurityLevel have all been set. The following objects may
    // not be modified while the value of this object is active(1):     -
    // snmpTargetParamsMPModel     - snmpTargetParamsSecurityModel     -
    // snmpTargetParamsSecurityName     - snmpTargetParamsSecurityLevel An attempt
    // to set these objects while the value of snmpTargetParamsRowStatus is
    // active(1) will result in an inconsistentValue error. The type is RowStatus.
    Snmptargetparamsrowstatus interface{}
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetFilter() yfilter.YFilter { return snmptargetparamsentry.YFilter }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) SetFilter(yf yfilter.YFilter) { snmptargetparamsentry.YFilter = yf }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetGoName(yname string) string {
    if yname == "snmpTargetParamsName" { return "Snmptargetparamsname" }
    if yname == "snmpTargetParamsMPModel" { return "Snmptargetparamsmpmodel" }
    if yname == "snmpTargetParamsSecurityModel" { return "Snmptargetparamssecuritymodel" }
    if yname == "snmpTargetParamsSecurityName" { return "Snmptargetparamssecurityname" }
    if yname == "snmpTargetParamsSecurityLevel" { return "Snmptargetparamssecuritylevel" }
    if yname == "snmpTargetParamsStorageType" { return "Snmptargetparamsstoragetype" }
    if yname == "snmpTargetParamsRowStatus" { return "Snmptargetparamsrowstatus" }
    return ""
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetSegmentPath() string {
    return "snmpTargetParamsEntry" + "[snmpTargetParamsName='" + fmt.Sprintf("%v", snmptargetparamsentry.Snmptargetparamsname) + "']"
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpTargetParamsName"] = snmptargetparamsentry.Snmptargetparamsname
    leafs["snmpTargetParamsMPModel"] = snmptargetparamsentry.Snmptargetparamsmpmodel
    leafs["snmpTargetParamsSecurityModel"] = snmptargetparamsentry.Snmptargetparamssecuritymodel
    leafs["snmpTargetParamsSecurityName"] = snmptargetparamsentry.Snmptargetparamssecurityname
    leafs["snmpTargetParamsSecurityLevel"] = snmptargetparamsentry.Snmptargetparamssecuritylevel
    leafs["snmpTargetParamsStorageType"] = snmptargetparamsentry.Snmptargetparamsstoragetype
    leafs["snmpTargetParamsRowStatus"] = snmptargetparamsentry.Snmptargetparamsrowstatus
    return leafs
}

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetBundleName() string { return "cisco_ios_xe" }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetYangName() string { return "snmpTargetParamsEntry" }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) SetParent(parent types.Entity) { snmptargetparamsentry.parent = parent }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetParent() types.Entity { return snmptargetparamsentry.parent }

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetParentYangName() string { return "snmpTargetParamsTable" }

