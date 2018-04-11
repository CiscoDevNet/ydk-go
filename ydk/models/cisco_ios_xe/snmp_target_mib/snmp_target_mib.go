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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Snmptargetobjects SNMPTARGETMIB_Snmptargetobjects

    // A table of transport addresses to be used in the generation of SNMP
    // messages.
    Snmptargetaddrtable SNMPTARGETMIB_Snmptargetaddrtable

    // A table of SNMP target information to be used in the generation of SNMP
    // messages.
    Snmptargetparamstable SNMPTARGETMIB_Snmptargetparamstable
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetEntityData() *types.CommonEntityData {
    sNMPTARGETMIB.EntityData.YFilter = sNMPTARGETMIB.YFilter
    sNMPTARGETMIB.EntityData.YangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPTARGETMIB.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.SegmentPath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPTARGETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPTARGETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPTARGETMIB.EntityData.Children = make(map[string]types.YChild)
    sNMPTARGETMIB.EntityData.Children["snmpTargetObjects"] = types.YChild{"Snmptargetobjects", &sNMPTARGETMIB.Snmptargetobjects}
    sNMPTARGETMIB.EntityData.Children["snmpTargetAddrTable"] = types.YChild{"Snmptargetaddrtable", &sNMPTARGETMIB.Snmptargetaddrtable}
    sNMPTARGETMIB.EntityData.Children["snmpTargetParamsTable"] = types.YChild{"Snmptargetparamstable", &sNMPTARGETMIB.Snmptargetparamstable}
    sNMPTARGETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sNMPTARGETMIB.EntityData)
}

// SNMPTARGETMIB_Snmptargetobjects
type SNMPTARGETMIB_Snmptargetobjects struct {
    EntityData types.CommonEntityData
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

func (snmptargetobjects *SNMPTARGETMIB_Snmptargetobjects) GetEntityData() *types.CommonEntityData {
    snmptargetobjects.EntityData.YFilter = snmptargetobjects.YFilter
    snmptargetobjects.EntityData.YangName = "snmpTargetObjects"
    snmptargetobjects.EntityData.BundleName = "cisco_ios_xe"
    snmptargetobjects.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmptargetobjects.EntityData.SegmentPath = "snmpTargetObjects"
    snmptargetobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmptargetobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmptargetobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmptargetobjects.EntityData.Children = make(map[string]types.YChild)
    snmptargetobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    snmptargetobjects.EntityData.Leafs["snmpTargetSpinLock"] = types.YLeaf{"Snmptargetspinlock", snmptargetobjects.Snmptargetspinlock}
    snmptargetobjects.EntityData.Leafs["snmpUnavailableContexts"] = types.YLeaf{"Snmpunavailablecontexts", snmptargetobjects.Snmpunavailablecontexts}
    snmptargetobjects.EntityData.Leafs["snmpUnknownContexts"] = types.YLeaf{"Snmpunknowncontexts", snmptargetobjects.Snmpunknowncontexts}
    return &(snmptargetobjects.EntityData)
}

// SNMPTARGETMIB_Snmptargetaddrtable
// A table of transport addresses to be used in the generation
// of SNMP messages.
type SNMPTARGETMIB_Snmptargetaddrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A transport address to be used in the generation of SNMP operations. 
    // Entries in the snmpTargetAddrTable are created and deleted using the
    // snmpTargetAddrRowStatus object. The type is slice of
    // SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry.
    Snmptargetaddrentry []SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry
}

func (snmptargetaddrtable *SNMPTARGETMIB_Snmptargetaddrtable) GetEntityData() *types.CommonEntityData {
    snmptargetaddrtable.EntityData.YFilter = snmptargetaddrtable.YFilter
    snmptargetaddrtable.EntityData.YangName = "snmpTargetAddrTable"
    snmptargetaddrtable.EntityData.BundleName = "cisco_ios_xe"
    snmptargetaddrtable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmptargetaddrtable.EntityData.SegmentPath = "snmpTargetAddrTable"
    snmptargetaddrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmptargetaddrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmptargetaddrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmptargetaddrtable.EntityData.Children = make(map[string]types.YChild)
    snmptargetaddrtable.EntityData.Children["snmpTargetAddrEntry"] = types.YChild{"Snmptargetaddrentry", nil}
    for i := range snmptargetaddrtable.Snmptargetaddrentry {
        snmptargetaddrtable.EntityData.Children[types.GetSegmentPath(&snmptargetaddrtable.Snmptargetaddrentry[i])] = types.YChild{"Snmptargetaddrentry", &snmptargetaddrtable.Snmptargetaddrentry[i]}
    }
    snmptargetaddrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snmptargetaddrtable.EntityData)
}

// SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry
// A transport address to be used in the generation
// of SNMP operations.
// 
// Entries in the snmpTargetAddrTable are created and
// deleted using the snmpTargetAddrRowStatus object.
type SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpTargetAddrEntry. The type is string with length:
    // 1..32.
    Snmptargetaddrname interface{}

    // This object indicates the transport type of the address contained in the
    // snmpTargetAddrTAddress object. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (snmptargetaddrentry *SNMPTARGETMIB_Snmptargetaddrtable_Snmptargetaddrentry) GetEntityData() *types.CommonEntityData {
    snmptargetaddrentry.EntityData.YFilter = snmptargetaddrentry.YFilter
    snmptargetaddrentry.EntityData.YangName = "snmpTargetAddrEntry"
    snmptargetaddrentry.EntityData.BundleName = "cisco_ios_xe"
    snmptargetaddrentry.EntityData.ParentYangName = "snmpTargetAddrTable"
    snmptargetaddrentry.EntityData.SegmentPath = "snmpTargetAddrEntry" + "[snmpTargetAddrName='" + fmt.Sprintf("%v", snmptargetaddrentry.Snmptargetaddrname) + "']"
    snmptargetaddrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmptargetaddrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmptargetaddrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmptargetaddrentry.EntityData.Children = make(map[string]types.YChild)
    snmptargetaddrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrName"] = types.YLeaf{"Snmptargetaddrname", snmptargetaddrentry.Snmptargetaddrname}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrTDomain"] = types.YLeaf{"Snmptargetaddrtdomain", snmptargetaddrentry.Snmptargetaddrtdomain}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrTAddress"] = types.YLeaf{"Snmptargetaddrtaddress", snmptargetaddrentry.Snmptargetaddrtaddress}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrTimeout"] = types.YLeaf{"Snmptargetaddrtimeout", snmptargetaddrentry.Snmptargetaddrtimeout}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrRetryCount"] = types.YLeaf{"Snmptargetaddrretrycount", snmptargetaddrentry.Snmptargetaddrretrycount}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrTagList"] = types.YLeaf{"Snmptargetaddrtaglist", snmptargetaddrentry.Snmptargetaddrtaglist}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrParams"] = types.YLeaf{"Snmptargetaddrparams", snmptargetaddrentry.Snmptargetaddrparams}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrStorageType"] = types.YLeaf{"Snmptargetaddrstoragetype", snmptargetaddrentry.Snmptargetaddrstoragetype}
    snmptargetaddrentry.EntityData.Leafs["snmpTargetAddrRowStatus"] = types.YLeaf{"Snmptargetaddrrowstatus", snmptargetaddrentry.Snmptargetaddrrowstatus}
    return &(snmptargetaddrentry.EntityData)
}

// SNMPTARGETMIB_Snmptargetparamstable
// A table of SNMP target information to be used
// in the generation of SNMP messages.
type SNMPTARGETMIB_Snmptargetparamstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of SNMP target information.  Entries in the snmpTargetParamsTable are
    // created and deleted using the snmpTargetParamsRowStatus object. The type is
    // slice of SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry.
    Snmptargetparamsentry []SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry
}

func (snmptargetparamstable *SNMPTARGETMIB_Snmptargetparamstable) GetEntityData() *types.CommonEntityData {
    snmptargetparamstable.EntityData.YFilter = snmptargetparamstable.YFilter
    snmptargetparamstable.EntityData.YangName = "snmpTargetParamsTable"
    snmptargetparamstable.EntityData.BundleName = "cisco_ios_xe"
    snmptargetparamstable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmptargetparamstable.EntityData.SegmentPath = "snmpTargetParamsTable"
    snmptargetparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmptargetparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmptargetparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmptargetparamstable.EntityData.Children = make(map[string]types.YChild)
    snmptargetparamstable.EntityData.Children["snmpTargetParamsEntry"] = types.YChild{"Snmptargetparamsentry", nil}
    for i := range snmptargetparamstable.Snmptargetparamsentry {
        snmptargetparamstable.EntityData.Children[types.GetSegmentPath(&snmptargetparamstable.Snmptargetparamsentry[i])] = types.YChild{"Snmptargetparamsentry", &snmptargetparamstable.Snmptargetparamsentry[i]}
    }
    snmptargetparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snmptargetparamstable.EntityData)
}

// SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry
// A set of SNMP target information.
// 
// Entries in the snmpTargetParamsTable are created and
// deleted using the snmpTargetParamsRowStatus object.
type SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry struct {
    EntityData types.CommonEntityData
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

func (snmptargetparamsentry *SNMPTARGETMIB_Snmptargetparamstable_Snmptargetparamsentry) GetEntityData() *types.CommonEntityData {
    snmptargetparamsentry.EntityData.YFilter = snmptargetparamsentry.YFilter
    snmptargetparamsentry.EntityData.YangName = "snmpTargetParamsEntry"
    snmptargetparamsentry.EntityData.BundleName = "cisco_ios_xe"
    snmptargetparamsentry.EntityData.ParentYangName = "snmpTargetParamsTable"
    snmptargetparamsentry.EntityData.SegmentPath = "snmpTargetParamsEntry" + "[snmpTargetParamsName='" + fmt.Sprintf("%v", snmptargetparamsentry.Snmptargetparamsname) + "']"
    snmptargetparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmptargetparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmptargetparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmptargetparamsentry.EntityData.Children = make(map[string]types.YChild)
    snmptargetparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsName"] = types.YLeaf{"Snmptargetparamsname", snmptargetparamsentry.Snmptargetparamsname}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsMPModel"] = types.YLeaf{"Snmptargetparamsmpmodel", snmptargetparamsentry.Snmptargetparamsmpmodel}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsSecurityModel"] = types.YLeaf{"Snmptargetparamssecuritymodel", snmptargetparamsentry.Snmptargetparamssecuritymodel}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsSecurityName"] = types.YLeaf{"Snmptargetparamssecurityname", snmptargetparamsentry.Snmptargetparamssecurityname}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsSecurityLevel"] = types.YLeaf{"Snmptargetparamssecuritylevel", snmptargetparamsentry.Snmptargetparamssecuritylevel}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsStorageType"] = types.YLeaf{"Snmptargetparamsstoragetype", snmptargetparamsentry.Snmptargetparamsstoragetype}
    snmptargetparamsentry.EntityData.Leafs["snmpTargetParamsRowStatus"] = types.YLeaf{"Snmptargetparamsrowstatus", snmptargetparamsentry.Snmptargetparamsrowstatus}
    return &(snmptargetparamsentry.EntityData)
}

