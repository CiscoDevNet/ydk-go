// The MIB module for SNMP entities.
// 
// Copyright (C) The Internet Society (2002). This
// version of this MIB module is part of RFC 3418;
// see the RFC itself for full legal notices.
package snmpv2_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmpv2_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMPv2-MIB SNMPv2-MIB}", reflect.TypeOf(SNMPv2MIB{}))
    ydk.RegisterEntity("SNMPv2-MIB:SNMPv2-MIB", reflect.TypeOf(SNMPv2MIB{}))
}

// SNMPv2MIB
type SNMPv2MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    System SNMPv2MIB_System

    
    Snmp SNMPv2MIB_Snmp

    
    SnmpSet SNMPv2MIB_SnmpSet

    // The (conceptual) table listing the capabilities of the local SNMP
    // application acting as a command responder with respect to various MIB
    // modules. SNMP entities having dynamically-configurable support of MIB
    // modules will have a dynamically-varying number of conceptual rows.
    SysORTable SNMPv2MIB_SysORTable
}

func (sNMPv2MIB *SNMPv2MIB) GetEntityData() *types.CommonEntityData {
    sNMPv2MIB.EntityData.YFilter = sNMPv2MIB.YFilter
    sNMPv2MIB.EntityData.YangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPv2MIB.EntityData.ParentYangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.SegmentPath = "SNMPv2-MIB:SNMPv2-MIB"
    sNMPv2MIB.EntityData.AbsolutePath = sNMPv2MIB.EntityData.SegmentPath
    sNMPv2MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPv2MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPv2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPv2MIB.EntityData.Children = types.NewOrderedMap()
    sNMPv2MIB.EntityData.Children.Append("system", types.YChild{"System", &sNMPv2MIB.System})
    sNMPv2MIB.EntityData.Children.Append("snmp", types.YChild{"Snmp", &sNMPv2MIB.Snmp})
    sNMPv2MIB.EntityData.Children.Append("snmpSet", types.YChild{"SnmpSet", &sNMPv2MIB.SnmpSet})
    sNMPv2MIB.EntityData.Children.Append("sysORTable", types.YChild{"SysORTable", &sNMPv2MIB.SysORTable})
    sNMPv2MIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPv2MIB.EntityData.YListKeys = []string {}

    return &(sNMPv2MIB.EntityData)
}

// SNMPv2MIB_System
type SNMPv2MIB_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A textual description of the entity.  This value should include the full
    // name and version identification of the system's hardware type, software
    // operating-system, and networking software. The type is string with length:
    // 0..255.
    SysDescr interface{}

    // The vendor's authoritative identification of the network management
    // subsystem contained in the entity. This value is allocated within the SMI
    // enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous
    // means for determining `what kind of box' is being managed.  For example, if
    // vendor `Flintstones, Inc.' was assigned the subtree 1.3.6.1.4.1.424242, it
    // could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.
    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SysObjectID interface{}

    // The time (in hundredths of a second) since the network management portion
    // of the system was last re-initialized. The type is interface{} with range:
    // 0..4294967295.
    SysUpTime interface{}

    // The textual identification of the contact person for this managed node,
    // together with information on how to contact this person.  If no contact
    // information is known, the value is the zero-length string. The type is
    // string with length: 0..255.
    SysContact interface{}

    // An administratively-assigned name for this managed node.  By convention,
    // this is the node's fully-qualified domain name.  If the name is unknown,
    // the value is the zero-length string. The type is string with length:
    // 0..255.
    SysName interface{}

    // The physical location of this node (e.g., 'telephone closet, 3rd floor'). 
    // If the location is unknown, the value is the zero-length string. The type
    // is string with length: 0..255.
    SysLocation interface{}

    // A value which indicates the set of services that this entity may
    // potentially offer.  The value is a sum. This sum initially takes the value
    // zero. Then, for each layer, L, in the range 1 through 7, that this node
    // performs transactions for, 2 raised to (L - 1) is added to the sum.  For
    // example, a node which performs only routing functions would have a value of
    // 4 (2^(3-1)). In contrast, a node which is a host offering application
    // services would have a value of 72 (2^(4-1) + 2^(7-1)). Note that in the
    // context of the Internet suite of protocols, values should be calculated
    // accordingly:       layer      functionality        1        physical (e.g.,
    // repeaters)        2        datalink/subnetwork (e.g., bridges)        3    
    // internet (e.g., supports the IP)        4        end-to-end  (e.g.,
    // supports the TCP)        7        applications (e.g., supports the SMTP) 
    // For systems including OSI protocols, layers 5 and 6 may also be counted.
    // The type is interface{} with range: 0..127.
    SysServices interface{}

    // The value of sysUpTime at the time of the most recent change in state or
    // value of any instance of sysORID. The type is interface{} with range:
    // 0..4294967295.
    SysORLastChange interface{}
}

func (system *SNMPv2MIB_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xe"
    system.EntityData.ParentYangName = "SNMPv2-MIB"
    system.EntityData.SegmentPath = "system"
    system.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + system.EntityData.SegmentPath
    system.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Leafs = types.NewOrderedMap()
    system.EntityData.Leafs.Append("sysDescr", types.YLeaf{"SysDescr", system.SysDescr})
    system.EntityData.Leafs.Append("sysObjectID", types.YLeaf{"SysObjectID", system.SysObjectID})
    system.EntityData.Leafs.Append("sysUpTime", types.YLeaf{"SysUpTime", system.SysUpTime})
    system.EntityData.Leafs.Append("sysContact", types.YLeaf{"SysContact", system.SysContact})
    system.EntityData.Leafs.Append("sysName", types.YLeaf{"SysName", system.SysName})
    system.EntityData.Leafs.Append("sysLocation", types.YLeaf{"SysLocation", system.SysLocation})
    system.EntityData.Leafs.Append("sysServices", types.YLeaf{"SysServices", system.SysServices})
    system.EntityData.Leafs.Append("sysORLastChange", types.YLeaf{"SysORLastChange", system.SysORLastChange})

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// SNMPv2MIB_Snmp
type SNMPv2MIB_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of messages delivered to the SNMP entity from the
    // transport service. The type is interface{} with range: 0..4294967295.
    SnmpInPkts interface{}

    // The total number of SNMP Messages which were passed from the SNMP protocol
    // entity to the transport service. The type is interface{} with range:
    // 0..4294967295.
    SnmpOutPkts interface{}

    // The total number of SNMP messages which were delivered to the SNMP entity
    // and were for an unsupported SNMP version. The type is interface{} with
    // range: 0..4294967295.
    SnmpInBadVersions interface{}

    // The total number of community-based SNMP messages (for example,  SNMPv1)
    // delivered to the SNMP entity which used an SNMP community name not known to
    // said entity. Also, implementations which authenticate community-based SNMP
    // messages using check(s) in addition to matching the community name (for
    // example, by also checking whether the message originated from a transport
    // address allowed to use a specified community name) MAY include in this
    // value the number of messages which failed the additional check(s).  It is
    // strongly RECOMMENDED that the documentation for any security model which is
    // used to authenticate community-based SNMP messages specify the precise
    // conditions that contribute to this value. The type is interface{} with
    // range: 0..4294967295.
    SnmpInBadCommunityNames interface{}

    // The total number of community-based SNMP messages (for example, SNMPv1)
    // delivered to the SNMP entity which represented an SNMP operation that was
    // not allowed for the SNMP community named in the message.  The precise
    // conditions under which this counter is incremented (if at all) depend on
    // how the SNMP entity implements its access control mechanism and how its
    // applications interact with that access control mechanism.  It is strongly
    // RECOMMENDED that the documentation for any access control mechanism which
    // is used to control access to and visibility of MIB instrumentation specify
    // the precise conditions that contribute to this value. The type is
    // interface{} with range: 0..4294967295.
    SnmpInBadCommunityUses interface{}

    // The total number of ASN.1 or BER errors encountered by the SNMP entity when
    // decoding received SNMP messages. The type is interface{} with range:
    // 0..4294967295.
    SnmpInASNParseErrs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `tooBig'. The
    // type is interface{} with range: 0..4294967295.
    SnmpInTooBigs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `noSuchName'.
    // The type is interface{} with range: 0..4294967295.
    SnmpInNoSuchNames interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `badValue'.
    // The type is interface{} with range: 0..4294967295.
    SnmpInBadValues interface{}

    // The total number valid SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `readOnly'. 
    // It should be noted that it is a protocol error to generate an SNMP PDU
    // which contains the value `readOnly' in the error-status field, as such this
    // object is provided as a means of detecting incorrect implementations of the
    // SNMP. The type is interface{} with range: 0..4294967295.
    SnmpInReadOnlys interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `genErr'. The
    // type is interface{} with range: 0..4294967295.
    SnmpInGenErrs interface{}

    // The total number of MIB objects which have been retrieved successfully by
    // the SNMP protocol entity as the result of receiving valid SNMP Get-Request
    // and Get-Next PDUs. The type is interface{} with range: 0..4294967295.
    SnmpInTotalReqVars interface{}

    // The total number of MIB objects which have been altered successfully by the
    // SNMP protocol entity as the result of receiving valid SNMP Set-Request
    // PDUs. The type is interface{} with range: 0..4294967295.
    SnmpInTotalSetVars interface{}

    // The total number of SNMP Get-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetRequests interface{}

    // The total number of SNMP Get-Next PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetNexts interface{}

    // The total number of SNMP Set-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInSetRequests interface{}

    // The total number of SNMP Get-Response PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInGetResponses interface{}

    // The total number of SNMP Trap PDUs which have been accepted and processed
    // by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    SnmpInTraps interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `tooBig.'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutTooBigs interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status was `noSuchName'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutNoSuchNames interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `badValue'.
    // The type is interface{} with range: 0..4294967295.
    SnmpOutBadValues interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `genErr'. The
    // type is interface{} with range: 0..4294967295.
    SnmpOutGenErrs interface{}

    // The total number of SNMP Get-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetRequests interface{}

    // The total number of SNMP Get-Next PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetNexts interface{}

    // The total number of SNMP Set-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutSetRequests interface{}

    // The total number of SNMP Get-Response PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutGetResponses interface{}

    // The total number of SNMP Trap PDUs which have been generated by the SNMP
    // protocol entity. The type is interface{} with range: 0..4294967295.
    SnmpOutTraps interface{}

    // Indicates whether the SNMP entity is permitted to generate
    // authenticationFailure traps.  The value of this object overrides any
    // configuration information; as such, it provides a means whereby all
    // authenticationFailure traps may be disabled.  Note that it is strongly
    // recommended that this object be stored in non-volatile memory so that it
    // remains constant across re-initializations of the network management
    // system. The type is SnmpEnableAuthenTraps.
    SnmpEnableAuthenTraps interface{}

    // The total number of Confirmed Class PDUs (such as GetRequest-PDUs,
    // GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and
    // InformRequest-PDUs) delivered to the SNMP entity which were silently
    // dropped because the size of a reply containing an alternate Response Class
    // PDU (such as a Response-PDU) with an empty variable-bindings field was
    // greater than either a local constraint or the maximum message size
    // associated with the originator of the request. The type is interface{} with
    // range: 0..4294967295.
    SnmpSilentDrops interface{}

    // The total number of Confirmed Class PDUs (such as GetRequest-PDUs,
    // GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and
    // InformRequest-PDUs) delivered to the SNMP entity which were silently
    // dropped because the transmission of the (possibly translated) message to a
    // proxy target failed in a manner (other than a time-out) such that no
    // Response Class PDU (such as a Response-PDU) could be returned. The type is
    // interface{} with range: 0..4294967295.
    SnmpProxyDrops interface{}
}

func (snmp *SNMPv2MIB_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xe"
    snmp.EntityData.ParentYangName = "SNMPv2-MIB"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("snmpInPkts", types.YLeaf{"SnmpInPkts", snmp.SnmpInPkts})
    snmp.EntityData.Leafs.Append("snmpOutPkts", types.YLeaf{"SnmpOutPkts", snmp.SnmpOutPkts})
    snmp.EntityData.Leafs.Append("snmpInBadVersions", types.YLeaf{"SnmpInBadVersions", snmp.SnmpInBadVersions})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityNames", types.YLeaf{"SnmpInBadCommunityNames", snmp.SnmpInBadCommunityNames})
    snmp.EntityData.Leafs.Append("snmpInBadCommunityUses", types.YLeaf{"SnmpInBadCommunityUses", snmp.SnmpInBadCommunityUses})
    snmp.EntityData.Leafs.Append("snmpInASNParseErrs", types.YLeaf{"SnmpInASNParseErrs", snmp.SnmpInASNParseErrs})
    snmp.EntityData.Leafs.Append("snmpInTooBigs", types.YLeaf{"SnmpInTooBigs", snmp.SnmpInTooBigs})
    snmp.EntityData.Leafs.Append("snmpInNoSuchNames", types.YLeaf{"SnmpInNoSuchNames", snmp.SnmpInNoSuchNames})
    snmp.EntityData.Leafs.Append("snmpInBadValues", types.YLeaf{"SnmpInBadValues", snmp.SnmpInBadValues})
    snmp.EntityData.Leafs.Append("snmpInReadOnlys", types.YLeaf{"SnmpInReadOnlys", snmp.SnmpInReadOnlys})
    snmp.EntityData.Leafs.Append("snmpInGenErrs", types.YLeaf{"SnmpInGenErrs", snmp.SnmpInGenErrs})
    snmp.EntityData.Leafs.Append("snmpInTotalReqVars", types.YLeaf{"SnmpInTotalReqVars", snmp.SnmpInTotalReqVars})
    snmp.EntityData.Leafs.Append("snmpInTotalSetVars", types.YLeaf{"SnmpInTotalSetVars", snmp.SnmpInTotalSetVars})
    snmp.EntityData.Leafs.Append("snmpInGetRequests", types.YLeaf{"SnmpInGetRequests", snmp.SnmpInGetRequests})
    snmp.EntityData.Leafs.Append("snmpInGetNexts", types.YLeaf{"SnmpInGetNexts", snmp.SnmpInGetNexts})
    snmp.EntityData.Leafs.Append("snmpInSetRequests", types.YLeaf{"SnmpInSetRequests", snmp.SnmpInSetRequests})
    snmp.EntityData.Leafs.Append("snmpInGetResponses", types.YLeaf{"SnmpInGetResponses", snmp.SnmpInGetResponses})
    snmp.EntityData.Leafs.Append("snmpInTraps", types.YLeaf{"SnmpInTraps", snmp.SnmpInTraps})
    snmp.EntityData.Leafs.Append("snmpOutTooBigs", types.YLeaf{"SnmpOutTooBigs", snmp.SnmpOutTooBigs})
    snmp.EntityData.Leafs.Append("snmpOutNoSuchNames", types.YLeaf{"SnmpOutNoSuchNames", snmp.SnmpOutNoSuchNames})
    snmp.EntityData.Leafs.Append("snmpOutBadValues", types.YLeaf{"SnmpOutBadValues", snmp.SnmpOutBadValues})
    snmp.EntityData.Leafs.Append("snmpOutGenErrs", types.YLeaf{"SnmpOutGenErrs", snmp.SnmpOutGenErrs})
    snmp.EntityData.Leafs.Append("snmpOutGetRequests", types.YLeaf{"SnmpOutGetRequests", snmp.SnmpOutGetRequests})
    snmp.EntityData.Leafs.Append("snmpOutGetNexts", types.YLeaf{"SnmpOutGetNexts", snmp.SnmpOutGetNexts})
    snmp.EntityData.Leafs.Append("snmpOutSetRequests", types.YLeaf{"SnmpOutSetRequests", snmp.SnmpOutSetRequests})
    snmp.EntityData.Leafs.Append("snmpOutGetResponses", types.YLeaf{"SnmpOutGetResponses", snmp.SnmpOutGetResponses})
    snmp.EntityData.Leafs.Append("snmpOutTraps", types.YLeaf{"SnmpOutTraps", snmp.SnmpOutTraps})
    snmp.EntityData.Leafs.Append("snmpEnableAuthenTraps", types.YLeaf{"SnmpEnableAuthenTraps", snmp.SnmpEnableAuthenTraps})
    snmp.EntityData.Leafs.Append("snmpSilentDrops", types.YLeaf{"SnmpSilentDrops", snmp.SnmpSilentDrops})
    snmp.EntityData.Leafs.Append("snmpProxyDrops", types.YLeaf{"SnmpProxyDrops", snmp.SnmpProxyDrops})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// SNMPv2MIB_Snmp_SnmpEnableAuthenTraps represents management system.
type SNMPv2MIB_Snmp_SnmpEnableAuthenTraps string

const (
    SNMPv2MIB_Snmp_SnmpEnableAuthenTraps_enabled SNMPv2MIB_Snmp_SnmpEnableAuthenTraps = "enabled"

    SNMPv2MIB_Snmp_SnmpEnableAuthenTraps_disabled SNMPv2MIB_Snmp_SnmpEnableAuthenTraps = "disabled"
)

// SNMPv2MIB_SnmpSet
type SNMPv2MIB_SnmpSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An advisory lock used to allow several cooperating command generator
    // applications to coordinate their use of the SNMP set operation.  This
    // object is used for coarse-grain coordination. To achieve fine-grain
    // coordination, one or more similar objects might be defined within each MIB
    // group, as appropriate. The type is interface{} with range: 0..2147483647.
    SnmpSetSerialNo interface{}
}

func (snmpSet *SNMPv2MIB_SnmpSet) GetEntityData() *types.CommonEntityData {
    snmpSet.EntityData.YFilter = snmpSet.YFilter
    snmpSet.EntityData.YangName = "snmpSet"
    snmpSet.EntityData.BundleName = "cisco_ios_xe"
    snmpSet.EntityData.ParentYangName = "SNMPv2-MIB"
    snmpSet.EntityData.SegmentPath = "snmpSet"
    snmpSet.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + snmpSet.EntityData.SegmentPath
    snmpSet.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpSet.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpSet.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpSet.EntityData.Children = types.NewOrderedMap()
    snmpSet.EntityData.Leafs = types.NewOrderedMap()
    snmpSet.EntityData.Leafs.Append("snmpSetSerialNo", types.YLeaf{"SnmpSetSerialNo", snmpSet.SnmpSetSerialNo})

    snmpSet.EntityData.YListKeys = []string {}

    return &(snmpSet.EntityData)
}

// SNMPv2MIB_SysORTable
// The (conceptual) table listing the capabilities of
// the local SNMP application acting as a command
// responder with respect to various MIB modules.
// SNMP entities having dynamically-configurable support
// of MIB modules will have a dynamically-varying number
// of conceptual rows.
type SNMPv2MIB_SysORTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the sysORTable. The type is slice of
    // SNMPv2MIB_SysORTable_SysOREntry.
    SysOREntry []*SNMPv2MIB_SysORTable_SysOREntry
}

func (sysORTable *SNMPv2MIB_SysORTable) GetEntityData() *types.CommonEntityData {
    sysORTable.EntityData.YFilter = sysORTable.YFilter
    sysORTable.EntityData.YangName = "sysORTable"
    sysORTable.EntityData.BundleName = "cisco_ios_xe"
    sysORTable.EntityData.ParentYangName = "SNMPv2-MIB"
    sysORTable.EntityData.SegmentPath = "sysORTable"
    sysORTable.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/" + sysORTable.EntityData.SegmentPath
    sysORTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sysORTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sysORTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sysORTable.EntityData.Children = types.NewOrderedMap()
    sysORTable.EntityData.Children.Append("sysOREntry", types.YChild{"SysOREntry", nil})
    for i := range sysORTable.SysOREntry {
        sysORTable.EntityData.Children.Append(types.GetSegmentPath(sysORTable.SysOREntry[i]), types.YChild{"SysOREntry", sysORTable.SysOREntry[i]})
    }
    sysORTable.EntityData.Leafs = types.NewOrderedMap()

    sysORTable.EntityData.YListKeys = []string {}

    return &(sysORTable.EntityData)
}

// SNMPv2MIB_SysORTable_SysOREntry
// An entry (conceptual row) in the sysORTable.
type SNMPv2MIB_SysORTable_SysOREntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The auxiliary variable used for identifying
    // instances of the columnar objects in the sysORTable. The type is
    // interface{} with range: 1..2147483647.
    SysORIndex interface{}

    // An authoritative identification of a capabilities statement with respect to
    // various MIB modules supported by the local SNMP application acting as a
    // command responder. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    SysORID interface{}

    // A textual description of the capabilities identified by the corresponding
    // instance of sysORID. The type is string.
    SysORDescr interface{}

    // The value of sysUpTime at the time this conceptual row was last
    // instantiated. The type is interface{} with range: 0..4294967295.
    SysORUpTime interface{}
}

func (sysOREntry *SNMPv2MIB_SysORTable_SysOREntry) GetEntityData() *types.CommonEntityData {
    sysOREntry.EntityData.YFilter = sysOREntry.YFilter
    sysOREntry.EntityData.YangName = "sysOREntry"
    sysOREntry.EntityData.BundleName = "cisco_ios_xe"
    sysOREntry.EntityData.ParentYangName = "sysORTable"
    sysOREntry.EntityData.SegmentPath = "sysOREntry" + types.AddKeyToken(sysOREntry.SysORIndex, "sysORIndex")
    sysOREntry.EntityData.AbsolutePath = "SNMPv2-MIB:SNMPv2-MIB/sysORTable/" + sysOREntry.EntityData.SegmentPath
    sysOREntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sysOREntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sysOREntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sysOREntry.EntityData.Children = types.NewOrderedMap()
    sysOREntry.EntityData.Leafs = types.NewOrderedMap()
    sysOREntry.EntityData.Leafs.Append("sysORIndex", types.YLeaf{"SysORIndex", sysOREntry.SysORIndex})
    sysOREntry.EntityData.Leafs.Append("sysORID", types.YLeaf{"SysORID", sysOREntry.SysORID})
    sysOREntry.EntityData.Leafs.Append("sysORDescr", types.YLeaf{"SysORDescr", sysOREntry.SysORDescr})
    sysOREntry.EntityData.Leafs.Append("sysORUpTime", types.YLeaf{"SysORUpTime", sysOREntry.SysORUpTime})

    sysOREntry.EntityData.YListKeys = []string {"SysORIndex"}

    return &(sysOREntry.EntityData)
}

