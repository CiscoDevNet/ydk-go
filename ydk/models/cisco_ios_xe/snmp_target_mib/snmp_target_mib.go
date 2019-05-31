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

    
    SnmpTargetObjects SNMPTARGETMIB_SnmpTargetObjects

    // A table of transport addresses to be used in the generation of SNMP
    // messages.
    SnmpTargetAddrTable SNMPTARGETMIB_SnmpTargetAddrTable

    // A table of SNMP target information to be used in the generation of SNMP
    // messages.
    SnmpTargetParamsTable SNMPTARGETMIB_SnmpTargetParamsTable
}

func (sNMPTARGETMIB *SNMPTARGETMIB) GetEntityData() *types.CommonEntityData {
    sNMPTARGETMIB.EntityData.YFilter = sNMPTARGETMIB.YFilter
    sNMPTARGETMIB.EntityData.YangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPTARGETMIB.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.SegmentPath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB"
    sNMPTARGETMIB.EntityData.AbsolutePath = sNMPTARGETMIB.EntityData.SegmentPath
    sNMPTARGETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPTARGETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPTARGETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPTARGETMIB.EntityData.Children = types.NewOrderedMap()
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetObjects", types.YChild{"SnmpTargetObjects", &sNMPTARGETMIB.SnmpTargetObjects})
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetAddrTable", types.YChild{"SnmpTargetAddrTable", &sNMPTARGETMIB.SnmpTargetAddrTable})
    sNMPTARGETMIB.EntityData.Children.Append("snmpTargetParamsTable", types.YChild{"SnmpTargetParamsTable", &sNMPTARGETMIB.SnmpTargetParamsTable})
    sNMPTARGETMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPTARGETMIB.EntityData.YListKeys = []string {}

    return &(sNMPTARGETMIB.EntityData)
}

// SNMPTARGETMIB_SnmpTargetObjects
type SNMPTARGETMIB_SnmpTargetObjects struct {
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
    SnmpTargetSpinLock interface{}

    // The total number of packets received by the SNMP engine which were dropped
    // because the context contained in the message was unavailable. The type is
    // interface{} with range: 0..4294967295.
    SnmpUnavailableContexts interface{}

    // The total number of packets received by the SNMP engine which were dropped
    // because the context contained in the message was unknown. The type is
    // interface{} with range: 0..4294967295.
    SnmpUnknownContexts interface{}
}

func (snmpTargetObjects *SNMPTARGETMIB_SnmpTargetObjects) GetEntityData() *types.CommonEntityData {
    snmpTargetObjects.EntityData.YFilter = snmpTargetObjects.YFilter
    snmpTargetObjects.EntityData.YangName = "snmpTargetObjects"
    snmpTargetObjects.EntityData.BundleName = "cisco_ios_xe"
    snmpTargetObjects.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetObjects.EntityData.SegmentPath = "snmpTargetObjects"
    snmpTargetObjects.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetObjects.EntityData.SegmentPath
    snmpTargetObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTargetObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTargetObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTargetObjects.EntityData.Children = types.NewOrderedMap()
    snmpTargetObjects.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetObjects.EntityData.Leafs.Append("snmpTargetSpinLock", types.YLeaf{"SnmpTargetSpinLock", snmpTargetObjects.SnmpTargetSpinLock})
    snmpTargetObjects.EntityData.Leafs.Append("snmpUnavailableContexts", types.YLeaf{"SnmpUnavailableContexts", snmpTargetObjects.SnmpUnavailableContexts})
    snmpTargetObjects.EntityData.Leafs.Append("snmpUnknownContexts", types.YLeaf{"SnmpUnknownContexts", snmpTargetObjects.SnmpUnknownContexts})

    snmpTargetObjects.EntityData.YListKeys = []string {}

    return &(snmpTargetObjects.EntityData)
}

// SNMPTARGETMIB_SnmpTargetAddrTable
// A table of transport addresses to be used in the generation
// of SNMP messages.
type SNMPTARGETMIB_SnmpTargetAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A transport address to be used in the generation of SNMP operations. 
    // Entries in the snmpTargetAddrTable are created and deleted using the
    // snmpTargetAddrRowStatus object. The type is slice of
    // SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry.
    SnmpTargetAddrEntry []*SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry
}

func (snmpTargetAddrTable *SNMPTARGETMIB_SnmpTargetAddrTable) GetEntityData() *types.CommonEntityData {
    snmpTargetAddrTable.EntityData.YFilter = snmpTargetAddrTable.YFilter
    snmpTargetAddrTable.EntityData.YangName = "snmpTargetAddrTable"
    snmpTargetAddrTable.EntityData.BundleName = "cisco_ios_xe"
    snmpTargetAddrTable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetAddrTable.EntityData.SegmentPath = "snmpTargetAddrTable"
    snmpTargetAddrTable.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetAddrTable.EntityData.SegmentPath
    snmpTargetAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTargetAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTargetAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTargetAddrTable.EntityData.Children = types.NewOrderedMap()
    snmpTargetAddrTable.EntityData.Children.Append("snmpTargetAddrEntry", types.YChild{"SnmpTargetAddrEntry", nil})
    for i := range snmpTargetAddrTable.SnmpTargetAddrEntry {
        snmpTargetAddrTable.EntityData.Children.Append(types.GetSegmentPath(snmpTargetAddrTable.SnmpTargetAddrEntry[i]), types.YChild{"SnmpTargetAddrEntry", snmpTargetAddrTable.SnmpTargetAddrEntry[i]})
    }
    snmpTargetAddrTable.EntityData.Leafs = types.NewOrderedMap()

    snmpTargetAddrTable.EntityData.YListKeys = []string {}

    return &(snmpTargetAddrTable.EntityData)
}

// SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry
// A transport address to be used in the generation
// of SNMP operations.
// 
// Entries in the snmpTargetAddrTable are created and
// deleted using the snmpTargetAddrRowStatus object.
type SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpTargetAddrEntry. The type is string with length:
    // 1..32.
    SnmpTargetAddrName interface{}

    // This object indicates the transport type of the address contained in the
    // snmpTargetAddrTAddress object. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SnmpTargetAddrTDomain interface{}

    // This object contains a transport address.  The format of this address
    // depends on the value of the snmpTargetAddrTDomain object. The type is
    // string with length: 1..255.
    SnmpTargetAddrTAddress interface{}

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
    SnmpTargetAddrTimeout interface{}

    // This object specifies a default number of retries to be attempted when a
    // response is not received for a generated message.  An application may
    // provide its own retry count, in which case the value of this object is
    // ignored. The type is interface{} with range: 0..255.
    SnmpTargetAddrRetryCount interface{}

    // This object contains a list of tag values which are used to select target
    // addresses for a particular operation. The type is string.
    SnmpTargetAddrTagList interface{}

    // The value of this object identifies an entry in the snmpTargetParamsTable. 
    // The identified entry contains SNMP parameters to be used when generating
    // messages to be sent to this transport address. The type is string with
    // length: 1..32.
    SnmpTargetAddrParams interface{}

    // The storage type for this conceptual row. The type is StorageType.
    SnmpTargetAddrStorageType interface{}

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
    SnmpTargetAddrRowStatus interface{}
}

func (snmpTargetAddrEntry *SNMPTARGETMIB_SnmpTargetAddrTable_SnmpTargetAddrEntry) GetEntityData() *types.CommonEntityData {
    snmpTargetAddrEntry.EntityData.YFilter = snmpTargetAddrEntry.YFilter
    snmpTargetAddrEntry.EntityData.YangName = "snmpTargetAddrEntry"
    snmpTargetAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    snmpTargetAddrEntry.EntityData.ParentYangName = "snmpTargetAddrTable"
    snmpTargetAddrEntry.EntityData.SegmentPath = "snmpTargetAddrEntry" + types.AddKeyToken(snmpTargetAddrEntry.SnmpTargetAddrName, "snmpTargetAddrName")
    snmpTargetAddrEntry.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/snmpTargetAddrTable/" + snmpTargetAddrEntry.EntityData.SegmentPath
    snmpTargetAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTargetAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTargetAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTargetAddrEntry.EntityData.Children = types.NewOrderedMap()
    snmpTargetAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrName", types.YLeaf{"SnmpTargetAddrName", snmpTargetAddrEntry.SnmpTargetAddrName})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTDomain", types.YLeaf{"SnmpTargetAddrTDomain", snmpTargetAddrEntry.SnmpTargetAddrTDomain})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTAddress", types.YLeaf{"SnmpTargetAddrTAddress", snmpTargetAddrEntry.SnmpTargetAddrTAddress})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTimeout", types.YLeaf{"SnmpTargetAddrTimeout", snmpTargetAddrEntry.SnmpTargetAddrTimeout})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrRetryCount", types.YLeaf{"SnmpTargetAddrRetryCount", snmpTargetAddrEntry.SnmpTargetAddrRetryCount})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrTagList", types.YLeaf{"SnmpTargetAddrTagList", snmpTargetAddrEntry.SnmpTargetAddrTagList})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrParams", types.YLeaf{"SnmpTargetAddrParams", snmpTargetAddrEntry.SnmpTargetAddrParams})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrStorageType", types.YLeaf{"SnmpTargetAddrStorageType", snmpTargetAddrEntry.SnmpTargetAddrStorageType})
    snmpTargetAddrEntry.EntityData.Leafs.Append("snmpTargetAddrRowStatus", types.YLeaf{"SnmpTargetAddrRowStatus", snmpTargetAddrEntry.SnmpTargetAddrRowStatus})

    snmpTargetAddrEntry.EntityData.YListKeys = []string {"SnmpTargetAddrName"}

    return &(snmpTargetAddrEntry.EntityData)
}

// SNMPTARGETMIB_SnmpTargetParamsTable
// A table of SNMP target information to be used
// in the generation of SNMP messages.
type SNMPTARGETMIB_SnmpTargetParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of SNMP target information.  Entries in the snmpTargetParamsTable are
    // created and deleted using the snmpTargetParamsRowStatus object. The type is
    // slice of SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry.
    SnmpTargetParamsEntry []*SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry
}

func (snmpTargetParamsTable *SNMPTARGETMIB_SnmpTargetParamsTable) GetEntityData() *types.CommonEntityData {
    snmpTargetParamsTable.EntityData.YFilter = snmpTargetParamsTable.YFilter
    snmpTargetParamsTable.EntityData.YangName = "snmpTargetParamsTable"
    snmpTargetParamsTable.EntityData.BundleName = "cisco_ios_xe"
    snmpTargetParamsTable.EntityData.ParentYangName = "SNMP-TARGET-MIB"
    snmpTargetParamsTable.EntityData.SegmentPath = "snmpTargetParamsTable"
    snmpTargetParamsTable.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/" + snmpTargetParamsTable.EntityData.SegmentPath
    snmpTargetParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTargetParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTargetParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTargetParamsTable.EntityData.Children = types.NewOrderedMap()
    snmpTargetParamsTable.EntityData.Children.Append("snmpTargetParamsEntry", types.YChild{"SnmpTargetParamsEntry", nil})
    for i := range snmpTargetParamsTable.SnmpTargetParamsEntry {
        snmpTargetParamsTable.EntityData.Children.Append(types.GetSegmentPath(snmpTargetParamsTable.SnmpTargetParamsEntry[i]), types.YChild{"SnmpTargetParamsEntry", snmpTargetParamsTable.SnmpTargetParamsEntry[i]})
    }
    snmpTargetParamsTable.EntityData.Leafs = types.NewOrderedMap()

    snmpTargetParamsTable.EntityData.YListKeys = []string {}

    return &(snmpTargetParamsTable.EntityData)
}

// SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry
// A set of SNMP target information.
// 
// Entries in the snmpTargetParamsTable are created and
// deleted using the snmpTargetParamsRowStatus object.
type SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The locally arbitrary, but unique identifier
    // associated with this snmpTargetParamsEntry. The type is string with length:
    // 1..32.
    SnmpTargetParamsName interface{}

    // The Message Processing Model to be used when generating SNMP messages using
    // this entry. The type is interface{} with range: 0..2147483647.
    SnmpTargetParamsMPModel interface{}

    // The Security Model to be used when generating SNMP messages using this
    // entry.  An implementation may choose to return an inconsistentValue error
    // if an attempt is made to set this variable to a value for a security model
    // which the implementation does      not support. The type is interface{}
    // with range: 1..2147483647.
    SnmpTargetParamsSecurityModel interface{}

    // The securityName which identifies the Principal on whose behalf SNMP
    // messages will be generated using this entry. The type is string.
    SnmpTargetParamsSecurityName interface{}

    // The Level of Security to be used when generating SNMP messages using this
    // entry. The type is SnmpSecurityLevel.
    SnmpTargetParamsSecurityLevel interface{}

    // The storage type for this conceptual row. The type is StorageType.
    SnmpTargetParamsStorageType interface{}

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
    SnmpTargetParamsRowStatus interface{}
}

func (snmpTargetParamsEntry *SNMPTARGETMIB_SnmpTargetParamsTable_SnmpTargetParamsEntry) GetEntityData() *types.CommonEntityData {
    snmpTargetParamsEntry.EntityData.YFilter = snmpTargetParamsEntry.YFilter
    snmpTargetParamsEntry.EntityData.YangName = "snmpTargetParamsEntry"
    snmpTargetParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    snmpTargetParamsEntry.EntityData.ParentYangName = "snmpTargetParamsTable"
    snmpTargetParamsEntry.EntityData.SegmentPath = "snmpTargetParamsEntry" + types.AddKeyToken(snmpTargetParamsEntry.SnmpTargetParamsName, "snmpTargetParamsName")
    snmpTargetParamsEntry.EntityData.AbsolutePath = "SNMP-TARGET-MIB:SNMP-TARGET-MIB/snmpTargetParamsTable/" + snmpTargetParamsEntry.EntityData.SegmentPath
    snmpTargetParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTargetParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTargetParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTargetParamsEntry.EntityData.Children = types.NewOrderedMap()
    snmpTargetParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsName", types.YLeaf{"SnmpTargetParamsName", snmpTargetParamsEntry.SnmpTargetParamsName})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsMPModel", types.YLeaf{"SnmpTargetParamsMPModel", snmpTargetParamsEntry.SnmpTargetParamsMPModel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityModel", types.YLeaf{"SnmpTargetParamsSecurityModel", snmpTargetParamsEntry.SnmpTargetParamsSecurityModel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityName", types.YLeaf{"SnmpTargetParamsSecurityName", snmpTargetParamsEntry.SnmpTargetParamsSecurityName})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsSecurityLevel", types.YLeaf{"SnmpTargetParamsSecurityLevel", snmpTargetParamsEntry.SnmpTargetParamsSecurityLevel})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsStorageType", types.YLeaf{"SnmpTargetParamsStorageType", snmpTargetParamsEntry.SnmpTargetParamsStorageType})
    snmpTargetParamsEntry.EntityData.Leafs.Append("snmpTargetParamsRowStatus", types.YLeaf{"SnmpTargetParamsRowStatus", snmpTargetParamsEntry.SnmpTargetParamsRowStatus})

    snmpTargetParamsEntry.EntityData.YListKeys = []string {"SnmpTargetParamsName"}

    return &(snmpTargetParamsEntry.EntityData)
}

