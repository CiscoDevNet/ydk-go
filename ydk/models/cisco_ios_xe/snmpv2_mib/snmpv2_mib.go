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

    
    Snmpset SNMPv2MIB_Snmpset

    // The (conceptual) table listing the capabilities of the local SNMP
    // application acting as a command responder with respect to various MIB
    // modules. SNMP entities having dynamically-configurable support of MIB
    // modules will have a dynamically-varying number of conceptual rows.
    Sysortable SNMPv2MIB_Sysortable
}

func (sNMPv2MIB *SNMPv2MIB) GetEntityData() *types.CommonEntityData {
    sNMPv2MIB.EntityData.YFilter = sNMPv2MIB.YFilter
    sNMPv2MIB.EntityData.YangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.BundleName = "cisco_ios_xe"
    sNMPv2MIB.EntityData.ParentYangName = "SNMPv2-MIB"
    sNMPv2MIB.EntityData.SegmentPath = "SNMPv2-MIB:SNMPv2-MIB"
    sNMPv2MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sNMPv2MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sNMPv2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sNMPv2MIB.EntityData.Children = make(map[string]types.YChild)
    sNMPv2MIB.EntityData.Children["system"] = types.YChild{"System", &sNMPv2MIB.System}
    sNMPv2MIB.EntityData.Children["snmp"] = types.YChild{"Snmp", &sNMPv2MIB.Snmp}
    sNMPv2MIB.EntityData.Children["snmpSet"] = types.YChild{"Snmpset", &sNMPv2MIB.Snmpset}
    sNMPv2MIB.EntityData.Children["sysORTable"] = types.YChild{"Sysortable", &sNMPv2MIB.Sysortable}
    sNMPv2MIB.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Sysdescr interface{}

    // The vendor's authoritative identification of the network management
    // subsystem contained in the entity. This value is allocated within the SMI
    // enterprises subtree (1.3.6.1.4.1) and provides an easy and unambiguous
    // means for determining `what kind of box' is being managed.  For example, if
    // vendor `Flintstones, Inc.' was assigned the subtree 1.3.6.1.4.1.424242, it
    // could assign the identifier 1.3.6.1.4.1.424242.1.1 to its `Fred Router'.
    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Sysobjectid interface{}

    // The time (in hundredths of a second) since the network management portion
    // of the system was last re-initialized. The type is interface{} with range:
    // 0..4294967295.
    Sysuptime interface{}

    // The textual identification of the contact person for this managed node,
    // together with information on how to contact this person.  If no contact
    // information is known, the value is the zero-length string. The type is
    // string with length: 0..255.
    Syscontact interface{}

    // An administratively-assigned name for this managed node.  By convention,
    // this is the node's fully-qualified domain name.  If the name is unknown,
    // the value is the zero-length string. The type is string with length:
    // 0..255.
    Sysname interface{}

    // The physical location of this node (e.g., 'telephone closet, 3rd floor'). 
    // If the location is unknown, the value is the zero-length string. The type
    // is string with length: 0..255.
    Syslocation interface{}

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
    Sysservices interface{}

    // The value of sysUpTime at the time of the most recent change in state or
    // value of any instance of sysORID. The type is interface{} with range:
    // 0..4294967295.
    Sysorlastchange interface{}
}

func (system *SNMPv2MIB_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xe"
    system.EntityData.ParentYangName = "SNMPv2-MIB"
    system.EntityData.SegmentPath = "system"
    system.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    system.EntityData.Children = make(map[string]types.YChild)
    system.EntityData.Leafs = make(map[string]types.YLeaf)
    system.EntityData.Leafs["sysDescr"] = types.YLeaf{"Sysdescr", system.Sysdescr}
    system.EntityData.Leafs["sysObjectID"] = types.YLeaf{"Sysobjectid", system.Sysobjectid}
    system.EntityData.Leafs["sysUpTime"] = types.YLeaf{"Sysuptime", system.Sysuptime}
    system.EntityData.Leafs["sysContact"] = types.YLeaf{"Syscontact", system.Syscontact}
    system.EntityData.Leafs["sysName"] = types.YLeaf{"Sysname", system.Sysname}
    system.EntityData.Leafs["sysLocation"] = types.YLeaf{"Syslocation", system.Syslocation}
    system.EntityData.Leafs["sysServices"] = types.YLeaf{"Sysservices", system.Sysservices}
    system.EntityData.Leafs["sysORLastChange"] = types.YLeaf{"Sysorlastchange", system.Sysorlastchange}
    return &(system.EntityData)
}

// SNMPv2MIB_Snmp
type SNMPv2MIB_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of messages delivered to the SNMP entity from the
    // transport service. The type is interface{} with range: 0..4294967295.
    Snmpinpkts interface{}

    // The total number of SNMP Messages which were passed from the SNMP protocol
    // entity to the transport service. The type is interface{} with range:
    // 0..4294967295.
    Snmpoutpkts interface{}

    // The total number of SNMP messages which were delivered to the SNMP entity
    // and were for an unsupported SNMP version. The type is interface{} with
    // range: 0..4294967295.
    Snmpinbadversions interface{}

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
    Snmpinbadcommunitynames interface{}

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
    Snmpinbadcommunityuses interface{}

    // The total number of ASN.1 or BER errors encountered by the SNMP entity when
    // decoding received SNMP messages. The type is interface{} with range:
    // 0..4294967295.
    Snmpinasnparseerrs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `tooBig'. The
    // type is interface{} with range: 0..4294967295.
    Snmpintoobigs interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `noSuchName'.
    // The type is interface{} with range: 0..4294967295.
    Snmpinnosuchnames interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `badValue'.
    // The type is interface{} with range: 0..4294967295.
    Snmpinbadvalues interface{}

    // The total number valid SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `readOnly'. 
    // It should be noted that it is a protocol error to generate an SNMP PDU
    // which contains the value `readOnly' in the error-status field, as such this
    // object is provided as a means of detecting incorrect implementations of the
    // SNMP. The type is interface{} with range: 0..4294967295.
    Snmpinreadonlys interface{}

    // The total number of SNMP PDUs which were delivered to the SNMP protocol
    // entity and for which the value of the error-status field was `genErr'. The
    // type is interface{} with range: 0..4294967295.
    Snmpingenerrs interface{}

    // The total number of MIB objects which have been retrieved successfully by
    // the SNMP protocol entity as the result of receiving valid SNMP Get-Request
    // and Get-Next PDUs. The type is interface{} with range: 0..4294967295.
    Snmpintotalreqvars interface{}

    // The total number of MIB objects which have been altered successfully by the
    // SNMP protocol entity as the result of receiving valid SNMP Set-Request
    // PDUs. The type is interface{} with range: 0..4294967295.
    Snmpintotalsetvars interface{}

    // The total number of SNMP Get-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetrequests interface{}

    // The total number of SNMP Get-Next PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetnexts interface{}

    // The total number of SNMP Set-Request PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpinsetrequests interface{}

    // The total number of SNMP Get-Response PDUs which have been accepted and
    // processed by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpingetresponses interface{}

    // The total number of SNMP Trap PDUs which have been accepted and processed
    // by the SNMP protocol entity. The type is interface{} with range:
    // 0..4294967295.
    Snmpintraps interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `tooBig.'. The
    // type is interface{} with range: 0..4294967295.
    Snmpouttoobigs interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status was `noSuchName'. The
    // type is interface{} with range: 0..4294967295.
    Snmpoutnosuchnames interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `badValue'.
    // The type is interface{} with range: 0..4294967295.
    Snmpoutbadvalues interface{}

    // The total number of SNMP PDUs which were generated by the SNMP protocol
    // entity and for which the value of the error-status field was `genErr'. The
    // type is interface{} with range: 0..4294967295.
    Snmpoutgenerrs interface{}

    // The total number of SNMP Get-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetrequests interface{}

    // The total number of SNMP Get-Next PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetnexts interface{}

    // The total number of SNMP Set-Request PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutsetrequests interface{}

    // The total number of SNMP Get-Response PDUs which have been generated by the
    // SNMP protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpoutgetresponses interface{}

    // The total number of SNMP Trap PDUs which have been generated by the SNMP
    // protocol entity. The type is interface{} with range: 0..4294967295.
    Snmpouttraps interface{}

    // Indicates whether the SNMP entity is permitted to generate
    // authenticationFailure traps.  The value of this object overrides any
    // configuration information; as such, it provides a means whereby all
    // authenticationFailure traps may be disabled.  Note that it is strongly
    // recommended that this object be stored in non-volatile memory so that it
    // remains constant across re-initializations of the network management
    // system. The type is Snmpenableauthentraps.
    Snmpenableauthentraps interface{}

    // The total number of Confirmed Class PDUs (such as GetRequest-PDUs,
    // GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and
    // InformRequest-PDUs) delivered to the SNMP entity which were silently
    // dropped because the size of a reply containing an alternate Response Class
    // PDU (such as a Response-PDU) with an empty variable-bindings field was
    // greater than either a local constraint or the maximum message size
    // associated with the originator of the request. The type is interface{} with
    // range: 0..4294967295.
    Snmpsilentdrops interface{}

    // The total number of Confirmed Class PDUs (such as GetRequest-PDUs,
    // GetNextRequest-PDUs, GetBulkRequest-PDUs, SetRequest-PDUs, and
    // InformRequest-PDUs) delivered to the SNMP entity which were silently
    // dropped because the transmission of the (possibly translated) message to a
    // proxy target failed in a manner (other than a time-out) such that no
    // Response Class PDU (such as a Response-PDU) could be returned. The type is
    // interface{} with range: 0..4294967295.
    Snmpproxydrops interface{}
}

func (snmp *SNMPv2MIB_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xe"
    snmp.EntityData.ParentYangName = "SNMPv2-MIB"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmp.EntityData.Children = make(map[string]types.YChild)
    snmp.EntityData.Leafs = make(map[string]types.YLeaf)
    snmp.EntityData.Leafs["snmpInPkts"] = types.YLeaf{"Snmpinpkts", snmp.Snmpinpkts}
    snmp.EntityData.Leafs["snmpOutPkts"] = types.YLeaf{"Snmpoutpkts", snmp.Snmpoutpkts}
    snmp.EntityData.Leafs["snmpInBadVersions"] = types.YLeaf{"Snmpinbadversions", snmp.Snmpinbadversions}
    snmp.EntityData.Leafs["snmpInBadCommunityNames"] = types.YLeaf{"Snmpinbadcommunitynames", snmp.Snmpinbadcommunitynames}
    snmp.EntityData.Leafs["snmpInBadCommunityUses"] = types.YLeaf{"Snmpinbadcommunityuses", snmp.Snmpinbadcommunityuses}
    snmp.EntityData.Leafs["snmpInASNParseErrs"] = types.YLeaf{"Snmpinasnparseerrs", snmp.Snmpinasnparseerrs}
    snmp.EntityData.Leafs["snmpInTooBigs"] = types.YLeaf{"Snmpintoobigs", snmp.Snmpintoobigs}
    snmp.EntityData.Leafs["snmpInNoSuchNames"] = types.YLeaf{"Snmpinnosuchnames", snmp.Snmpinnosuchnames}
    snmp.EntityData.Leafs["snmpInBadValues"] = types.YLeaf{"Snmpinbadvalues", snmp.Snmpinbadvalues}
    snmp.EntityData.Leafs["snmpInReadOnlys"] = types.YLeaf{"Snmpinreadonlys", snmp.Snmpinreadonlys}
    snmp.EntityData.Leafs["snmpInGenErrs"] = types.YLeaf{"Snmpingenerrs", snmp.Snmpingenerrs}
    snmp.EntityData.Leafs["snmpInTotalReqVars"] = types.YLeaf{"Snmpintotalreqvars", snmp.Snmpintotalreqvars}
    snmp.EntityData.Leafs["snmpInTotalSetVars"] = types.YLeaf{"Snmpintotalsetvars", snmp.Snmpintotalsetvars}
    snmp.EntityData.Leafs["snmpInGetRequests"] = types.YLeaf{"Snmpingetrequests", snmp.Snmpingetrequests}
    snmp.EntityData.Leafs["snmpInGetNexts"] = types.YLeaf{"Snmpingetnexts", snmp.Snmpingetnexts}
    snmp.EntityData.Leafs["snmpInSetRequests"] = types.YLeaf{"Snmpinsetrequests", snmp.Snmpinsetrequests}
    snmp.EntityData.Leafs["snmpInGetResponses"] = types.YLeaf{"Snmpingetresponses", snmp.Snmpingetresponses}
    snmp.EntityData.Leafs["snmpInTraps"] = types.YLeaf{"Snmpintraps", snmp.Snmpintraps}
    snmp.EntityData.Leafs["snmpOutTooBigs"] = types.YLeaf{"Snmpouttoobigs", snmp.Snmpouttoobigs}
    snmp.EntityData.Leafs["snmpOutNoSuchNames"] = types.YLeaf{"Snmpoutnosuchnames", snmp.Snmpoutnosuchnames}
    snmp.EntityData.Leafs["snmpOutBadValues"] = types.YLeaf{"Snmpoutbadvalues", snmp.Snmpoutbadvalues}
    snmp.EntityData.Leafs["snmpOutGenErrs"] = types.YLeaf{"Snmpoutgenerrs", snmp.Snmpoutgenerrs}
    snmp.EntityData.Leafs["snmpOutGetRequests"] = types.YLeaf{"Snmpoutgetrequests", snmp.Snmpoutgetrequests}
    snmp.EntityData.Leafs["snmpOutGetNexts"] = types.YLeaf{"Snmpoutgetnexts", snmp.Snmpoutgetnexts}
    snmp.EntityData.Leafs["snmpOutSetRequests"] = types.YLeaf{"Snmpoutsetrequests", snmp.Snmpoutsetrequests}
    snmp.EntityData.Leafs["snmpOutGetResponses"] = types.YLeaf{"Snmpoutgetresponses", snmp.Snmpoutgetresponses}
    snmp.EntityData.Leafs["snmpOutTraps"] = types.YLeaf{"Snmpouttraps", snmp.Snmpouttraps}
    snmp.EntityData.Leafs["snmpEnableAuthenTraps"] = types.YLeaf{"Snmpenableauthentraps", snmp.Snmpenableauthentraps}
    snmp.EntityData.Leafs["snmpSilentDrops"] = types.YLeaf{"Snmpsilentdrops", snmp.Snmpsilentdrops}
    snmp.EntityData.Leafs["snmpProxyDrops"] = types.YLeaf{"Snmpproxydrops", snmp.Snmpproxydrops}
    return &(snmp.EntityData)
}

// SNMPv2MIB_Snmp_Snmpenableauthentraps represents management system.
type SNMPv2MIB_Snmp_Snmpenableauthentraps string

const (
    SNMPv2MIB_Snmp_Snmpenableauthentraps_enabled SNMPv2MIB_Snmp_Snmpenableauthentraps = "enabled"

    SNMPv2MIB_Snmp_Snmpenableauthentraps_disabled SNMPv2MIB_Snmp_Snmpenableauthentraps = "disabled"
)

// SNMPv2MIB_Snmpset
type SNMPv2MIB_Snmpset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An advisory lock used to allow several cooperating command generator
    // applications to coordinate their use of the SNMP set operation.  This
    // object is used for coarse-grain coordination. To achieve fine-grain
    // coordination, one or more similar objects might be defined within each MIB
    // group, as appropriate. The type is interface{} with range: 0..2147483647.
    Snmpsetserialno interface{}
}

func (snmpset *SNMPv2MIB_Snmpset) GetEntityData() *types.CommonEntityData {
    snmpset.EntityData.YFilter = snmpset.YFilter
    snmpset.EntityData.YangName = "snmpSet"
    snmpset.EntityData.BundleName = "cisco_ios_xe"
    snmpset.EntityData.ParentYangName = "SNMPv2-MIB"
    snmpset.EntityData.SegmentPath = "snmpSet"
    snmpset.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpset.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpset.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpset.EntityData.Children = make(map[string]types.YChild)
    snmpset.EntityData.Leafs = make(map[string]types.YLeaf)
    snmpset.EntityData.Leafs["snmpSetSerialNo"] = types.YLeaf{"Snmpsetserialno", snmpset.Snmpsetserialno}
    return &(snmpset.EntityData)
}

// SNMPv2MIB_Sysortable
// The (conceptual) table listing the capabilities of
// the local SNMP application acting as a command
// responder with respect to various MIB modules.
// SNMP entities having dynamically-configurable support
// of MIB modules will have a dynamically-varying number
// of conceptual rows.
type SNMPv2MIB_Sysortable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the sysORTable. The type is slice of
    // SNMPv2MIB_Sysortable_Sysorentry.
    Sysorentry []SNMPv2MIB_Sysortable_Sysorentry
}

func (sysortable *SNMPv2MIB_Sysortable) GetEntityData() *types.CommonEntityData {
    sysortable.EntityData.YFilter = sysortable.YFilter
    sysortable.EntityData.YangName = "sysORTable"
    sysortable.EntityData.BundleName = "cisco_ios_xe"
    sysortable.EntityData.ParentYangName = "SNMPv2-MIB"
    sysortable.EntityData.SegmentPath = "sysORTable"
    sysortable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sysortable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sysortable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sysortable.EntityData.Children = make(map[string]types.YChild)
    sysortable.EntityData.Children["sysOREntry"] = types.YChild{"Sysorentry", nil}
    for i := range sysortable.Sysorentry {
        sysortable.EntityData.Children[types.GetSegmentPath(&sysortable.Sysorentry[i])] = types.YChild{"Sysorentry", &sysortable.Sysorentry[i]}
    }
    sysortable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sysortable.EntityData)
}

// SNMPv2MIB_Sysortable_Sysorentry
// An entry (conceptual row) in the sysORTable.
type SNMPv2MIB_Sysortable_Sysorentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The auxiliary variable used for identifying
    // instances of the columnar objects in the sysORTable. The type is
    // interface{} with range: 1..2147483647.
    Sysorindex interface{}

    // An authoritative identification of a capabilities statement with respect to
    // various MIB modules supported by the local SNMP application acting as a
    // command responder. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Sysorid interface{}

    // A textual description of the capabilities identified by the corresponding
    // instance of sysORID. The type is string.
    Sysordescr interface{}

    // The value of sysUpTime at the time this conceptual row was last
    // instantiated. The type is interface{} with range: 0..4294967295.
    Sysoruptime interface{}
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetEntityData() *types.CommonEntityData {
    sysorentry.EntityData.YFilter = sysorentry.YFilter
    sysorentry.EntityData.YangName = "sysOREntry"
    sysorentry.EntityData.BundleName = "cisco_ios_xe"
    sysorentry.EntityData.ParentYangName = "sysORTable"
    sysorentry.EntityData.SegmentPath = "sysOREntry" + "[sysORIndex='" + fmt.Sprintf("%v", sysorentry.Sysorindex) + "']"
    sysorentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sysorentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sysorentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sysorentry.EntityData.Children = make(map[string]types.YChild)
    sysorentry.EntityData.Leafs = make(map[string]types.YLeaf)
    sysorentry.EntityData.Leafs["sysORIndex"] = types.YLeaf{"Sysorindex", sysorentry.Sysorindex}
    sysorentry.EntityData.Leafs["sysORID"] = types.YLeaf{"Sysorid", sysorentry.Sysorid}
    sysorentry.EntityData.Leafs["sysORDescr"] = types.YLeaf{"Sysordescr", sysorentry.Sysordescr}
    sysorentry.EntityData.Leafs["sysORUpTime"] = types.YLeaf{"Sysoruptime", sysorentry.Sysoruptime}
    return &(sysorentry.EntityData)
}

