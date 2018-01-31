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
    parent types.Entity
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

func (sNMPv2MIB *SNMPv2MIB) GetFilter() yfilter.YFilter { return sNMPv2MIB.YFilter }

func (sNMPv2MIB *SNMPv2MIB) SetFilter(yf yfilter.YFilter) { sNMPv2MIB.YFilter = yf }

func (sNMPv2MIB *SNMPv2MIB) GetGoName(yname string) string {
    if yname == "system" { return "System" }
    if yname == "snmp" { return "Snmp" }
    if yname == "snmpSet" { return "Snmpset" }
    if yname == "sysORTable" { return "Sysortable" }
    return ""
}

func (sNMPv2MIB *SNMPv2MIB) GetSegmentPath() string {
    return "SNMPv2-MIB:SNMPv2-MIB"
}

func (sNMPv2MIB *SNMPv2MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "system" {
        return &sNMPv2MIB.System
    }
    if childYangName == "snmp" {
        return &sNMPv2MIB.Snmp
    }
    if childYangName == "snmpSet" {
        return &sNMPv2MIB.Snmpset
    }
    if childYangName == "sysORTable" {
        return &sNMPv2MIB.Sysortable
    }
    return nil
}

func (sNMPv2MIB *SNMPv2MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["system"] = &sNMPv2MIB.System
    children["snmp"] = &sNMPv2MIB.Snmp
    children["snmpSet"] = &sNMPv2MIB.Snmpset
    children["sysORTable"] = &sNMPv2MIB.Sysortable
    return children
}

func (sNMPv2MIB *SNMPv2MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sNMPv2MIB *SNMPv2MIB) GetBundleName() string { return "cisco_ios_xe" }

func (sNMPv2MIB *SNMPv2MIB) GetYangName() string { return "SNMPv2-MIB" }

func (sNMPv2MIB *SNMPv2MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sNMPv2MIB *SNMPv2MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sNMPv2MIB *SNMPv2MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sNMPv2MIB *SNMPv2MIB) SetParent(parent types.Entity) { sNMPv2MIB.parent = parent }

func (sNMPv2MIB *SNMPv2MIB) GetParent() types.Entity { return sNMPv2MIB.parent }

func (sNMPv2MIB *SNMPv2MIB) GetParentYangName() string { return "SNMPv2-MIB" }

// SNMPv2MIB_System
type SNMPv2MIB_System struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (system *SNMPv2MIB_System) GetFilter() yfilter.YFilter { return system.YFilter }

func (system *SNMPv2MIB_System) SetFilter(yf yfilter.YFilter) { system.YFilter = yf }

func (system *SNMPv2MIB_System) GetGoName(yname string) string {
    if yname == "sysDescr" { return "Sysdescr" }
    if yname == "sysObjectID" { return "Sysobjectid" }
    if yname == "sysUpTime" { return "Sysuptime" }
    if yname == "sysContact" { return "Syscontact" }
    if yname == "sysName" { return "Sysname" }
    if yname == "sysLocation" { return "Syslocation" }
    if yname == "sysServices" { return "Sysservices" }
    if yname == "sysORLastChange" { return "Sysorlastchange" }
    return ""
}

func (system *SNMPv2MIB_System) GetSegmentPath() string {
    return "system"
}

func (system *SNMPv2MIB_System) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (system *SNMPv2MIB_System) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (system *SNMPv2MIB_System) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sysDescr"] = system.Sysdescr
    leafs["sysObjectID"] = system.Sysobjectid
    leafs["sysUpTime"] = system.Sysuptime
    leafs["sysContact"] = system.Syscontact
    leafs["sysName"] = system.Sysname
    leafs["sysLocation"] = system.Syslocation
    leafs["sysServices"] = system.Sysservices
    leafs["sysORLastChange"] = system.Sysorlastchange
    return leafs
}

func (system *SNMPv2MIB_System) GetBundleName() string { return "cisco_ios_xe" }

func (system *SNMPv2MIB_System) GetYangName() string { return "system" }

func (system *SNMPv2MIB_System) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (system *SNMPv2MIB_System) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (system *SNMPv2MIB_System) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (system *SNMPv2MIB_System) SetParent(parent types.Entity) { system.parent = parent }

func (system *SNMPv2MIB_System) GetParent() types.Entity { return system.parent }

func (system *SNMPv2MIB_System) GetParentYangName() string { return "SNMPv2-MIB" }

// SNMPv2MIB_Snmp
type SNMPv2MIB_Snmp struct {
    parent types.Entity
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

func (snmp *SNMPv2MIB_Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *SNMPv2MIB_Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *SNMPv2MIB_Snmp) GetGoName(yname string) string {
    if yname == "snmpInPkts" { return "Snmpinpkts" }
    if yname == "snmpOutPkts" { return "Snmpoutpkts" }
    if yname == "snmpInBadVersions" { return "Snmpinbadversions" }
    if yname == "snmpInBadCommunityNames" { return "Snmpinbadcommunitynames" }
    if yname == "snmpInBadCommunityUses" { return "Snmpinbadcommunityuses" }
    if yname == "snmpInASNParseErrs" { return "Snmpinasnparseerrs" }
    if yname == "snmpInTooBigs" { return "Snmpintoobigs" }
    if yname == "snmpInNoSuchNames" { return "Snmpinnosuchnames" }
    if yname == "snmpInBadValues" { return "Snmpinbadvalues" }
    if yname == "snmpInReadOnlys" { return "Snmpinreadonlys" }
    if yname == "snmpInGenErrs" { return "Snmpingenerrs" }
    if yname == "snmpInTotalReqVars" { return "Snmpintotalreqvars" }
    if yname == "snmpInTotalSetVars" { return "Snmpintotalsetvars" }
    if yname == "snmpInGetRequests" { return "Snmpingetrequests" }
    if yname == "snmpInGetNexts" { return "Snmpingetnexts" }
    if yname == "snmpInSetRequests" { return "Snmpinsetrequests" }
    if yname == "snmpInGetResponses" { return "Snmpingetresponses" }
    if yname == "snmpInTraps" { return "Snmpintraps" }
    if yname == "snmpOutTooBigs" { return "Snmpouttoobigs" }
    if yname == "snmpOutNoSuchNames" { return "Snmpoutnosuchnames" }
    if yname == "snmpOutBadValues" { return "Snmpoutbadvalues" }
    if yname == "snmpOutGenErrs" { return "Snmpoutgenerrs" }
    if yname == "snmpOutGetRequests" { return "Snmpoutgetrequests" }
    if yname == "snmpOutGetNexts" { return "Snmpoutgetnexts" }
    if yname == "snmpOutSetRequests" { return "Snmpoutsetrequests" }
    if yname == "snmpOutGetResponses" { return "Snmpoutgetresponses" }
    if yname == "snmpOutTraps" { return "Snmpouttraps" }
    if yname == "snmpEnableAuthenTraps" { return "Snmpenableauthentraps" }
    if yname == "snmpSilentDrops" { return "Snmpsilentdrops" }
    if yname == "snmpProxyDrops" { return "Snmpproxydrops" }
    return ""
}

func (snmp *SNMPv2MIB_Snmp) GetSegmentPath() string {
    return "snmp"
}

func (snmp *SNMPv2MIB_Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmp *SNMPv2MIB_Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmp *SNMPv2MIB_Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpInPkts"] = snmp.Snmpinpkts
    leafs["snmpOutPkts"] = snmp.Snmpoutpkts
    leafs["snmpInBadVersions"] = snmp.Snmpinbadversions
    leafs["snmpInBadCommunityNames"] = snmp.Snmpinbadcommunitynames
    leafs["snmpInBadCommunityUses"] = snmp.Snmpinbadcommunityuses
    leafs["snmpInASNParseErrs"] = snmp.Snmpinasnparseerrs
    leafs["snmpInTooBigs"] = snmp.Snmpintoobigs
    leafs["snmpInNoSuchNames"] = snmp.Snmpinnosuchnames
    leafs["snmpInBadValues"] = snmp.Snmpinbadvalues
    leafs["snmpInReadOnlys"] = snmp.Snmpinreadonlys
    leafs["snmpInGenErrs"] = snmp.Snmpingenerrs
    leafs["snmpInTotalReqVars"] = snmp.Snmpintotalreqvars
    leafs["snmpInTotalSetVars"] = snmp.Snmpintotalsetvars
    leafs["snmpInGetRequests"] = snmp.Snmpingetrequests
    leafs["snmpInGetNexts"] = snmp.Snmpingetnexts
    leafs["snmpInSetRequests"] = snmp.Snmpinsetrequests
    leafs["snmpInGetResponses"] = snmp.Snmpingetresponses
    leafs["snmpInTraps"] = snmp.Snmpintraps
    leafs["snmpOutTooBigs"] = snmp.Snmpouttoobigs
    leafs["snmpOutNoSuchNames"] = snmp.Snmpoutnosuchnames
    leafs["snmpOutBadValues"] = snmp.Snmpoutbadvalues
    leafs["snmpOutGenErrs"] = snmp.Snmpoutgenerrs
    leafs["snmpOutGetRequests"] = snmp.Snmpoutgetrequests
    leafs["snmpOutGetNexts"] = snmp.Snmpoutgetnexts
    leafs["snmpOutSetRequests"] = snmp.Snmpoutsetrequests
    leafs["snmpOutGetResponses"] = snmp.Snmpoutgetresponses
    leafs["snmpOutTraps"] = snmp.Snmpouttraps
    leafs["snmpEnableAuthenTraps"] = snmp.Snmpenableauthentraps
    leafs["snmpSilentDrops"] = snmp.Snmpsilentdrops
    leafs["snmpProxyDrops"] = snmp.Snmpproxydrops
    return leafs
}

func (snmp *SNMPv2MIB_Snmp) GetBundleName() string { return "cisco_ios_xe" }

func (snmp *SNMPv2MIB_Snmp) GetYangName() string { return "snmp" }

func (snmp *SNMPv2MIB_Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmp *SNMPv2MIB_Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmp *SNMPv2MIB_Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmp *SNMPv2MIB_Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *SNMPv2MIB_Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *SNMPv2MIB_Snmp) GetParentYangName() string { return "SNMPv2-MIB" }

// SNMPv2MIB_Snmp_Snmpenableauthentraps represents management system.
type SNMPv2MIB_Snmp_Snmpenableauthentraps string

const (
    SNMPv2MIB_Snmp_Snmpenableauthentraps_enabled SNMPv2MIB_Snmp_Snmpenableauthentraps = "enabled"

    SNMPv2MIB_Snmp_Snmpenableauthentraps_disabled SNMPv2MIB_Snmp_Snmpenableauthentraps = "disabled"
)

// SNMPv2MIB_Snmpset
type SNMPv2MIB_Snmpset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An advisory lock used to allow several cooperating command generator
    // applications to coordinate their use of the SNMP set operation.  This
    // object is used for coarse-grain coordination. To achieve fine-grain
    // coordination, one or more similar objects might be defined within each MIB
    // group, as appropriate. The type is interface{} with range: 0..2147483647.
    Snmpsetserialno interface{}
}

func (snmpset *SNMPv2MIB_Snmpset) GetFilter() yfilter.YFilter { return snmpset.YFilter }

func (snmpset *SNMPv2MIB_Snmpset) SetFilter(yf yfilter.YFilter) { snmpset.YFilter = yf }

func (snmpset *SNMPv2MIB_Snmpset) GetGoName(yname string) string {
    if yname == "snmpSetSerialNo" { return "Snmpsetserialno" }
    return ""
}

func (snmpset *SNMPv2MIB_Snmpset) GetSegmentPath() string {
    return "snmpSet"
}

func (snmpset *SNMPv2MIB_Snmpset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmpset *SNMPv2MIB_Snmpset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmpset *SNMPv2MIB_Snmpset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpSetSerialNo"] = snmpset.Snmpsetserialno
    return leafs
}

func (snmpset *SNMPv2MIB_Snmpset) GetBundleName() string { return "cisco_ios_xe" }

func (snmpset *SNMPv2MIB_Snmpset) GetYangName() string { return "snmpSet" }

func (snmpset *SNMPv2MIB_Snmpset) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpset *SNMPv2MIB_Snmpset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpset *SNMPv2MIB_Snmpset) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpset *SNMPv2MIB_Snmpset) SetParent(parent types.Entity) { snmpset.parent = parent }

func (snmpset *SNMPv2MIB_Snmpset) GetParent() types.Entity { return snmpset.parent }

func (snmpset *SNMPv2MIB_Snmpset) GetParentYangName() string { return "SNMPv2-MIB" }

// SNMPv2MIB_Sysortable
// The (conceptual) table listing the capabilities of
// the local SNMP application acting as a command
// responder with respect to various MIB modules.
// SNMP entities having dynamically-configurable support
// of MIB modules will have a dynamically-varying number
// of conceptual rows.
type SNMPv2MIB_Sysortable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the sysORTable. The type is slice of
    // SNMPv2MIB_Sysortable_Sysorentry.
    Sysorentry []SNMPv2MIB_Sysortable_Sysorentry
}

func (sysortable *SNMPv2MIB_Sysortable) GetFilter() yfilter.YFilter { return sysortable.YFilter }

func (sysortable *SNMPv2MIB_Sysortable) SetFilter(yf yfilter.YFilter) { sysortable.YFilter = yf }

func (sysortable *SNMPv2MIB_Sysortable) GetGoName(yname string) string {
    if yname == "sysOREntry" { return "Sysorentry" }
    return ""
}

func (sysortable *SNMPv2MIB_Sysortable) GetSegmentPath() string {
    return "sysORTable"
}

func (sysortable *SNMPv2MIB_Sysortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sysOREntry" {
        for _, c := range sysortable.Sysorentry {
            if sysortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SNMPv2MIB_Sysortable_Sysorentry{}
        sysortable.Sysorentry = append(sysortable.Sysorentry, child)
        return &sysortable.Sysorentry[len(sysortable.Sysorentry)-1]
    }
    return nil
}

func (sysortable *SNMPv2MIB_Sysortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sysortable.Sysorentry {
        children[sysortable.Sysorentry[i].GetSegmentPath()] = &sysortable.Sysorentry[i]
    }
    return children
}

func (sysortable *SNMPv2MIB_Sysortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sysortable *SNMPv2MIB_Sysortable) GetBundleName() string { return "cisco_ios_xe" }

func (sysortable *SNMPv2MIB_Sysortable) GetYangName() string { return "sysORTable" }

func (sysortable *SNMPv2MIB_Sysortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sysortable *SNMPv2MIB_Sysortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sysortable *SNMPv2MIB_Sysortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sysortable *SNMPv2MIB_Sysortable) SetParent(parent types.Entity) { sysortable.parent = parent }

func (sysortable *SNMPv2MIB_Sysortable) GetParent() types.Entity { return sysortable.parent }

func (sysortable *SNMPv2MIB_Sysortable) GetParentYangName() string { return "SNMPv2-MIB" }

// SNMPv2MIB_Sysortable_Sysorentry
// An entry (conceptual row) in the sysORTable.
type SNMPv2MIB_Sysortable_Sysorentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The auxiliary variable used for identifying
    // instances of the columnar objects in the sysORTable. The type is
    // interface{} with range: 1..2147483647.
    Sysorindex interface{}

    // An authoritative identification of a capabilities statement with respect to
    // various MIB modules supported by the local SNMP application acting as a
    // command responder. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Sysorid interface{}

    // A textual description of the capabilities identified by the corresponding
    // instance of sysORID. The type is string.
    Sysordescr interface{}

    // The value of sysUpTime at the time this conceptual row was last
    // instantiated. The type is interface{} with range: 0..4294967295.
    Sysoruptime interface{}
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetFilter() yfilter.YFilter { return sysorentry.YFilter }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) SetFilter(yf yfilter.YFilter) { sysorentry.YFilter = yf }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetGoName(yname string) string {
    if yname == "sysORIndex" { return "Sysorindex" }
    if yname == "sysORID" { return "Sysorid" }
    if yname == "sysORDescr" { return "Sysordescr" }
    if yname == "sysORUpTime" { return "Sysoruptime" }
    return ""
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetSegmentPath() string {
    return "sysOREntry" + "[sysORIndex='" + fmt.Sprintf("%v", sysorentry.Sysorindex) + "']"
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sysORIndex"] = sysorentry.Sysorindex
    leafs["sysORID"] = sysorentry.Sysorid
    leafs["sysORDescr"] = sysorentry.Sysordescr
    leafs["sysORUpTime"] = sysorentry.Sysoruptime
    return leafs
}

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetBundleName() string { return "cisco_ios_xe" }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetYangName() string { return "sysOREntry" }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) SetParent(parent types.Entity) { sysorentry.parent = parent }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetParent() types.Entity { return sysorentry.parent }

func (sysorentry *SNMPv2MIB_Sysortable_Sysorentry) GetParentYangName() string { return "sysORTable" }

