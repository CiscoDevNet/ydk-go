// The SNMP Management Architecture MIB
// 
// Copyright (C) The Internet Society (2002). This
// version of this MIB module is part of RFC 3411;
// see the RFC itself for full legal notices.
package snmp_framework_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_framework_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:SNMP-FRAMEWORK-MIB SNMP-FRAMEWORK-MIB}", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
    ydk.RegisterEntity("SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB", reflect.TypeOf(SNMPFRAMEWORKMIB{}))
}

type Snmpprivprotocols struct {
}

func (id Snmpprivprotocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpPrivProtocols"
}

type Snmpauthprotocols struct {
}

func (id Snmpauthprotocols) String() string {
	return "SNMP-FRAMEWORK-MIB:snmpAuthProtocols"
}

// SnmpSecurityLevel represents authNoPriv is less than authPriv.
type SnmpSecurityLevel string

const (
    SnmpSecurityLevel_noAuthNoPriv SnmpSecurityLevel = "noAuthNoPriv"

    SnmpSecurityLevel_authNoPriv SnmpSecurityLevel = "authNoPriv"

    SnmpSecurityLevel_authPriv SnmpSecurityLevel = "authPriv"
)

// SNMPFRAMEWORKMIB
type SNMPFRAMEWORKMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Snmpengine SNMPFRAMEWORKMIB_Snmpengine
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetFilter() yfilter.YFilter { return sNMPFRAMEWORKMIB.YFilter }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) SetFilter(yf yfilter.YFilter) { sNMPFRAMEWORKMIB.YFilter = yf }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetGoName(yname string) string {
    if yname == "snmpEngine" { return "Snmpengine" }
    return ""
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetSegmentPath() string {
    return "SNMP-FRAMEWORK-MIB:SNMP-FRAMEWORK-MIB"
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmpEngine" {
        return &sNMPFRAMEWORKMIB.Snmpengine
    }
    return nil
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snmpEngine"] = &sNMPFRAMEWORKMIB.Snmpengine
    return children
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetBundleName() string { return "cisco_ios_xe" }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetYangName() string { return "SNMP-FRAMEWORK-MIB" }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) SetParent(parent types.Entity) { sNMPFRAMEWORKMIB.parent = parent }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetParent() types.Entity { return sNMPFRAMEWORKMIB.parent }

func (sNMPFRAMEWORKMIB *SNMPFRAMEWORKMIB) GetParentYangName() string { return "SNMP-FRAMEWORK-MIB" }

// SNMPFRAMEWORKMIB_Snmpengine
type SNMPFRAMEWORKMIB_Snmpengine struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An SNMP engine's administratively-unique identifier.  This information
    // SHOULD be stored in non-volatile storage so that it remains constant across
    // re-initializations of the SNMP engine. The type is string with length:
    // 5..32.
    Snmpengineid interface{}

    // The number of times that the SNMP engine has (re-)initialized itself since
    // snmpEngineID was last configured. The type is interface{} with range:
    // 1..2147483647.
    Snmpengineboots interface{}

    // The number of seconds since the value of the snmpEngineBoots object last
    // changed. When incrementing this object's value would cause it to exceed its
    // maximum, snmpEngineBoots is incremented as if a re-initialization had
    // occurred, and this object's value consequently reverts to zero. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    Snmpenginetime interface{}

    // The maximum length in octets of an SNMP message which this SNMP engine can
    // send or receive and process, determined as the minimum of the maximum
    // message size values supported among all of the transports available to and
    // supported by the engine. The type is interface{} with range:
    // 484..2147483647.
    Snmpenginemaxmessagesize interface{}
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetFilter() yfilter.YFilter { return snmpengine.YFilter }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) SetFilter(yf yfilter.YFilter) { snmpengine.YFilter = yf }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetGoName(yname string) string {
    if yname == "snmpEngineID" { return "Snmpengineid" }
    if yname == "snmpEngineBoots" { return "Snmpengineboots" }
    if yname == "snmpEngineTime" { return "Snmpenginetime" }
    if yname == "snmpEngineMaxMessageSize" { return "Snmpenginemaxmessagesize" }
    return ""
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetSegmentPath() string {
    return "snmpEngine"
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["snmpEngineID"] = snmpengine.Snmpengineid
    leafs["snmpEngineBoots"] = snmpengine.Snmpengineboots
    leafs["snmpEngineTime"] = snmpengine.Snmpenginetime
    leafs["snmpEngineMaxMessageSize"] = snmpengine.Snmpenginemaxmessagesize
    return leafs
}

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetBundleName() string { return "cisco_ios_xe" }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetYangName() string { return "snmpEngine" }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) SetParent(parent types.Entity) { snmpengine.parent = parent }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetParent() types.Entity { return snmpengine.parent }

func (snmpengine *SNMPFRAMEWORKMIB_Snmpengine) GetParentYangName() string { return "SNMP-FRAMEWORK-MIB" }

