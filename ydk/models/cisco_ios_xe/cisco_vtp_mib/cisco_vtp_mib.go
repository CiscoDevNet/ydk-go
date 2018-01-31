// The MIB module for entities implementing the VTP
// protocol and Vlan management.
package cisco_vtp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_vtp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VTP-MIB CISCO-VTP-MIB}", reflect.TypeOf(CISCOVTPMIB{}))
    ydk.RegisterEntity("CISCO-VTP-MIB:CISCO-VTP-MIB", reflect.TypeOf(CISCOVTPMIB{}))
}

// VlanType represents to represent the bridged broadcast domain.
type VlanType string

const (
    VlanType_ethernet VlanType = "ethernet"

    VlanType_fddi VlanType = "fddi"

    VlanType_tokenRing VlanType = "tokenRing"

    VlanType_fddiNet VlanType = "fddiNet"

    VlanType_trNet VlanType = "trNet"

    VlanType_deprecated VlanType = "deprecated"
)

// CISCOVTPMIB
type CISCOVTPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Vtpstatus CISCOVTPMIB_Vtpstatus

    
    Internalvlaninfo CISCOVTPMIB_Internalvlaninfo

    
    Vlantrunkports CISCOVTPMIB_Vlantrunkports

    
    Vlanstatistics CISCOVTPMIB_Vlanstatistics

    // The table containing information on the management domains in which the
    // local system is participating.  Devices which support only one management
    // domain will support just one row in this table, and will not let it be
    // deleted nor let other rows be created.  Devices which support multiple
    // management domains will allow rows to be created and deleted, but will not
    // allow the last row to be deleted. If the device does  not support VTP, the
    // table is read-only.
    Managementdomaintable CISCOVTPMIB_Managementdomaintable

    // This table contains information on the VLANs which currently exist.
    Vtpvlantable CISCOVTPMIB_Vtpvlantable

    // A vtpInternalVlanTable entry contains information on an existing internal
    // VLAN. It is internally created by the device for a specific application
    // program  and hence owned by the application.   It cannot be modified or
    // deleted by (local  or network) management.
    Vtpinternalvlantable CISCOVTPMIB_Vtpinternalvlantable

    // The table which contains the information in the Edit Buffers, one Edit
    // Buffer per management domain.  The information for a particular management
    // domain is initialized, by a 'copy' operation, to be the current global VLAN
    // information for that management domain.  After initialization, editing can
    // be performed to add VLANs, delete VLANs, or modify their global parameters.
    // The information as modified through editing is local to this Edit Buffer. 
    // An apply operation using the vtpVlanEditOperation object is necessary to
    // instanciate the modified information as the new global VLAN information for
    // that management domain.  To use the Edit Buffer, a manager acts as follows:
    // 1. ensures the Edit Buffer for a management domain is empty, i.e., there
    // are no rows in this table for this management domain.  2. issues a SNMP set
    // operation which sets vtpVlanEditOperation to 'copy', and
    // vtpVlanEditBufferOwner to its own identifier (e.g., its own IP address). 
    // 3. if this set operation is successful, proceeds to edit the information in
    // the vtpVlanEditTable.  4. if and when the edited information is to be
    // instantiated, issues a SNMP set operation which sets vtpVlanEditOperation
    // to 'apply'.  5. issues retrieval requests to obtain the value of
    // vtpVlanApplyStatus, until the result of the apply is determined.  6.
    // releases the Edit Buffer by issuing a SNMP set operation which sets
    // vtpVlanEditOperation to 'release'.  Note that the information contained in
    // this table is not saved across agent reboots.
    Vtpvlanedittable CISCOVTPMIB_Vtpvlanedittable

    // Ths table contains the VLAN local shutdown information within management
    // domain.
    Vtpvlanlocalshutdowntable CISCOVTPMIB_Vtpvlanlocalshutdowntable

    // The table containing information on the local system's VLAN trunk ports.
    Vlantrunkporttable CISCOVTPMIB_Vlantrunkporttable

    // This table contains information related to the discovery of the VTP members
    // in the designated management domain. This table is not instantiated when 
    // managementDomainVersionInUse is version1(1), version2(3) or none(3).
    Vtpdiscovertable CISCOVTPMIB_Vtpdiscovertable

    // The table containing information of discovered VTP members in the
    // management domain in which the local system is participating. This table is
    // not instantiated when  managementDomainVersionInUse is version1(1),
    // version2(2) or none(3).
    Vtpdiscoverresulttable CISCOVTPMIB_Vtpdiscoverresulttable

    // This table contains information of the VTP databases. It is not
    // instantiated when managementDomainVersionInUse is version1(1),  version2(3)
    // or none(3).
    Vtpdatabasetable CISCOVTPMIB_Vtpdatabasetable

    // The table contains the authentication information of VTP in which the local
    // system participates.  The security mechanism of VTP relies on a secret key
    // that is used to alter the MD5 digest of the packets transmitted on the
    // wire.  The secret value is created from a password that may be saved in
    // plain text in the configuration or hidden from the configuration.  The
    // device creating or modifying the VTP configuration signs it using the MD5
    // digest generated from the secret key before advertising it. Other devices
    // in the domain receive this configuration use the same secret key to accept
    // it if correctly signed or drop it otherwise.  The user has the option to
    // hide the password from the configuration. Once the password is hidden, the
    // secret key generated from the password is shown in the  configuration
    // instead, and there is no other way to  show the password in plain text
    // again but clearing  it or resetting it.   In an un-trusted area, the
    // password on a device can  be configured without being unveiled. After that,
    // it has to be provided again by setting the same  value to
    // vtpDatabaseTakeOverPassword if the user  wants to take over the whole VTP
    // management domain of the database type.  When managementDomainVersionInUse
    // is version3(4), the  authentication mechanism is common to all VTP database
    // type.
    Vtpauthenticationtable CISCOVTPMIB_Vtpauthenticationtable
}

func (cISCOVTPMIB *CISCOVTPMIB) GetFilter() yfilter.YFilter { return cISCOVTPMIB.YFilter }

func (cISCOVTPMIB *CISCOVTPMIB) SetFilter(yf yfilter.YFilter) { cISCOVTPMIB.YFilter = yf }

func (cISCOVTPMIB *CISCOVTPMIB) GetGoName(yname string) string {
    if yname == "vtpStatus" { return "Vtpstatus" }
    if yname == "internalVlanInfo" { return "Internalvlaninfo" }
    if yname == "vlanTrunkPorts" { return "Vlantrunkports" }
    if yname == "vlanStatistics" { return "Vlanstatistics" }
    if yname == "managementDomainTable" { return "Managementdomaintable" }
    if yname == "vtpVlanTable" { return "Vtpvlantable" }
    if yname == "vtpInternalVlanTable" { return "Vtpinternalvlantable" }
    if yname == "vtpVlanEditTable" { return "Vtpvlanedittable" }
    if yname == "vtpVlanLocalShutdownTable" { return "Vtpvlanlocalshutdowntable" }
    if yname == "vlanTrunkPortTable" { return "Vlantrunkporttable" }
    if yname == "vtpDiscoverTable" { return "Vtpdiscovertable" }
    if yname == "vtpDiscoverResultTable" { return "Vtpdiscoverresulttable" }
    if yname == "vtpDatabaseTable" { return "Vtpdatabasetable" }
    if yname == "vtpAuthenticationTable" { return "Vtpauthenticationtable" }
    return ""
}

func (cISCOVTPMIB *CISCOVTPMIB) GetSegmentPath() string {
    return "CISCO-VTP-MIB:CISCO-VTP-MIB"
}

func (cISCOVTPMIB *CISCOVTPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpStatus" {
        return &cISCOVTPMIB.Vtpstatus
    }
    if childYangName == "internalVlanInfo" {
        return &cISCOVTPMIB.Internalvlaninfo
    }
    if childYangName == "vlanTrunkPorts" {
        return &cISCOVTPMIB.Vlantrunkports
    }
    if childYangName == "vlanStatistics" {
        return &cISCOVTPMIB.Vlanstatistics
    }
    if childYangName == "managementDomainTable" {
        return &cISCOVTPMIB.Managementdomaintable
    }
    if childYangName == "vtpVlanTable" {
        return &cISCOVTPMIB.Vtpvlantable
    }
    if childYangName == "vtpInternalVlanTable" {
        return &cISCOVTPMIB.Vtpinternalvlantable
    }
    if childYangName == "vtpVlanEditTable" {
        return &cISCOVTPMIB.Vtpvlanedittable
    }
    if childYangName == "vtpVlanLocalShutdownTable" {
        return &cISCOVTPMIB.Vtpvlanlocalshutdowntable
    }
    if childYangName == "vlanTrunkPortTable" {
        return &cISCOVTPMIB.Vlantrunkporttable
    }
    if childYangName == "vtpDiscoverTable" {
        return &cISCOVTPMIB.Vtpdiscovertable
    }
    if childYangName == "vtpDiscoverResultTable" {
        return &cISCOVTPMIB.Vtpdiscoverresulttable
    }
    if childYangName == "vtpDatabaseTable" {
        return &cISCOVTPMIB.Vtpdatabasetable
    }
    if childYangName == "vtpAuthenticationTable" {
        return &cISCOVTPMIB.Vtpauthenticationtable
    }
    return nil
}

func (cISCOVTPMIB *CISCOVTPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vtpStatus"] = &cISCOVTPMIB.Vtpstatus
    children["internalVlanInfo"] = &cISCOVTPMIB.Internalvlaninfo
    children["vlanTrunkPorts"] = &cISCOVTPMIB.Vlantrunkports
    children["vlanStatistics"] = &cISCOVTPMIB.Vlanstatistics
    children["managementDomainTable"] = &cISCOVTPMIB.Managementdomaintable
    children["vtpVlanTable"] = &cISCOVTPMIB.Vtpvlantable
    children["vtpInternalVlanTable"] = &cISCOVTPMIB.Vtpinternalvlantable
    children["vtpVlanEditTable"] = &cISCOVTPMIB.Vtpvlanedittable
    children["vtpVlanLocalShutdownTable"] = &cISCOVTPMIB.Vtpvlanlocalshutdowntable
    children["vlanTrunkPortTable"] = &cISCOVTPMIB.Vlantrunkporttable
    children["vtpDiscoverTable"] = &cISCOVTPMIB.Vtpdiscovertable
    children["vtpDiscoverResultTable"] = &cISCOVTPMIB.Vtpdiscoverresulttable
    children["vtpDatabaseTable"] = &cISCOVTPMIB.Vtpdatabasetable
    children["vtpAuthenticationTable"] = &cISCOVTPMIB.Vtpauthenticationtable
    return children
}

func (cISCOVTPMIB *CISCOVTPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVTPMIB *CISCOVTPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVTPMIB *CISCOVTPMIB) GetYangName() string { return "CISCO-VTP-MIB" }

func (cISCOVTPMIB *CISCOVTPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVTPMIB *CISCOVTPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVTPMIB *CISCOVTPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVTPMIB *CISCOVTPMIB) SetParent(parent types.Entity) { cISCOVTPMIB.parent = parent }

func (cISCOVTPMIB *CISCOVTPMIB) GetParent() types.Entity { return cISCOVTPMIB.parent }

func (cISCOVTPMIB *CISCOVTPMIB) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpstatus
type CISCOVTPMIB_Vtpstatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The version of VTP in use on the local system.  A device will report its
    // version capability and not any particular version in use on the device. If
    // the device does not support vtp, the version is none(3). The type is
    // Vtpversion.
    Vtpversion interface{}

    // An estimate of the maximum number of VLANs about which the local system can
    // recover complete VTP information after a reboot.  If the number of defined
    // VLANs is greater than this value, then the system can not act as a VTP
    // Server. For a device which has no means to calculate the estimated number,
    // this value is -1. The type is interface{} with range: -1..1023.
    Vtpmaxvlanstorage interface{}

    // An indication of whether the notifications/traps defined by the
    // vtpConfigNotificationsGroup, vtpConfigNotificationsGroup2, and
    // vtpConfigNotificationsGroup8 are enabled. The type is bool.
    Vtpnotificationsenabled interface{}

    // An indication of whether the notification should be generated when a VLAN
    // is created.   If the value of this object is 'true' then the vtpVlanCreated
    // notification will be generated.  If the value of this object is 'false'
    // then the vtpVlanCreated notification will not be generated. The type is
    // bool.
    Vtpvlancreatednotifenabled interface{}

    // An indication of whether the notification should be generated when a VLAN
    // is deleted.    If the value of this object is 'true' then the
    // vtpVlanDeleted notification will be generated.  If the value of this object
    // is 'false' then the vtpVlanDeleted notification will not be generated. The
    // type is bool.
    Vtpvlandeletednotifenabled interface{}
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetFilter() yfilter.YFilter { return vtpstatus.YFilter }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) SetFilter(yf yfilter.YFilter) { vtpstatus.YFilter = yf }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetGoName(yname string) string {
    if yname == "vtpVersion" { return "Vtpversion" }
    if yname == "vtpMaxVlanStorage" { return "Vtpmaxvlanstorage" }
    if yname == "vtpNotificationsEnabled" { return "Vtpnotificationsenabled" }
    if yname == "vtpVlanCreatedNotifEnabled" { return "Vtpvlancreatednotifenabled" }
    if yname == "vtpVlanDeletedNotifEnabled" { return "Vtpvlandeletednotifenabled" }
    return ""
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetSegmentPath() string {
    return "vtpStatus"
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vtpVersion"] = vtpstatus.Vtpversion
    leafs["vtpMaxVlanStorage"] = vtpstatus.Vtpmaxvlanstorage
    leafs["vtpNotificationsEnabled"] = vtpstatus.Vtpnotificationsenabled
    leafs["vtpVlanCreatedNotifEnabled"] = vtpstatus.Vtpvlancreatednotifenabled
    leafs["vtpVlanDeletedNotifEnabled"] = vtpstatus.Vtpvlandeletednotifenabled
    return leafs
}

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetBundleName() string { return "cisco_ios_xe" }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetYangName() string { return "vtpStatus" }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) SetParent(parent types.Entity) { vtpstatus.parent = parent }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetParent() types.Entity { return vtpstatus.parent }

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpstatus_Vtpversion represents vtp, the version is none(3).
type CISCOVTPMIB_Vtpstatus_Vtpversion string

const (
    CISCOVTPMIB_Vtpstatus_Vtpversion_one CISCOVTPMIB_Vtpstatus_Vtpversion = "one"

    CISCOVTPMIB_Vtpstatus_Vtpversion_two CISCOVTPMIB_Vtpstatus_Vtpversion = "two"

    CISCOVTPMIB_Vtpstatus_Vtpversion_none CISCOVTPMIB_Vtpstatus_Vtpversion = "none"

    CISCOVTPMIB_Vtpstatus_Vtpversion_three CISCOVTPMIB_Vtpstatus_Vtpversion = "three"
)

// CISCOVTPMIB_Internalvlaninfo
type CISCOVTPMIB_Internalvlaninfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The internal VLAN allocation policy.  'ascending'  - internal VLANs are
    // allocated                starting from a lowwer VLAN ID and                
    // upwards. 'descending' - internal VLANs are allocated               
    // starting from a higher VLAN ID and                downwards. The type is
    // Vtpinternalvlanallocpolicy.
    Vtpinternalvlanallocpolicy interface{}
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetFilter() yfilter.YFilter { return internalvlaninfo.YFilter }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) SetFilter(yf yfilter.YFilter) { internalvlaninfo.YFilter = yf }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetGoName(yname string) string {
    if yname == "vtpInternalVlanAllocPolicy" { return "Vtpinternalvlanallocpolicy" }
    return ""
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetSegmentPath() string {
    return "internalVlanInfo"
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vtpInternalVlanAllocPolicy"] = internalvlaninfo.Vtpinternalvlanallocpolicy
    return leafs
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetBundleName() string { return "cisco_ios_xe" }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetYangName() string { return "internalVlanInfo" }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) SetParent(parent types.Entity) { internalvlaninfo.parent = parent }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetParent() types.Entity { return internalvlaninfo.parent }

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy represents                downwards.
type CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy string

const (
    CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy_ascending CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy = "ascending"

    CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy_descending CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy = "descending"
)

// CISCOVTPMIB_Vlantrunkports
type CISCOVTPMIB_Vlantrunkports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An advisory lock used to allow several cooperating SNMPv2 managers to
    // coordinate their use of the SNMPv2 set operation acting upon any instance
    // of vlanTrunkPortVlansEnabled. The type is interface{} with range:
    // 0..2147483647.
    Vlantrunkportsetserialno interface{}

    // An indication of whether the tagging on all VLANs including native VLAN for
    // all 802.1q trunks is enabled.  If this object has a value of true(1) then
    // all VLANs including native VLAN are tagged.  If the value is false(2) then
    // all VLANs excluding native VLAN are tagged.  This object has been
    // deprecated and is replaced by the object 'cltcDot1qAllTaggedEnabled' in the
    // CISCO-L2-TUNNEL-CONFIG-MIB. The type is bool.
    Vlantrunkportsdot1Qtag interface{}
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetFilter() yfilter.YFilter { return vlantrunkports.YFilter }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) SetFilter(yf yfilter.YFilter) { vlantrunkports.YFilter = yf }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetGoName(yname string) string {
    if yname == "vlanTrunkPortSetSerialNo" { return "Vlantrunkportsetserialno" }
    if yname == "vlanTrunkPortsDot1qTag" { return "Vlantrunkportsdot1Qtag" }
    return ""
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetSegmentPath() string {
    return "vlanTrunkPorts"
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlanTrunkPortSetSerialNo"] = vlantrunkports.Vlantrunkportsetserialno
    leafs["vlanTrunkPortsDot1qTag"] = vlantrunkports.Vlantrunkportsdot1Qtag
    return leafs
}

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetBundleName() string { return "cisco_ios_xe" }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetYangName() string { return "vlanTrunkPorts" }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) SetParent(parent types.Entity) { vlantrunkports.parent = parent }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetParent() types.Entity { return vlantrunkports.parent }

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vlanstatistics
type CISCOVTPMIB_Vlanstatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object indicates the number of the existing manageable VLANs with VLAN
    // indices from 1 to 1024 in the system. The type is interface{} with range:
    // 0..4294967295.
    Vlanstatsvlans interface{}

    // This object indicates the number of the existing manageable VLANs with VLAN
    // indices greater than 1024 in the system. The type is interface{} with
    // range: 0..4294967295.
    Vlanstatsextendedvlans interface{}

    // This object indicates the number of the internal VLANs existing in the
    // system. The type is interface{} with range: 0..4294967295.
    Vlanstatsinternalvlans interface{}

    // This object indicates the number of the free or unused VLANs in the system.
    // The type is interface{} with range: 0..4294967295.
    Vlanstatsfreevlans interface{}
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetFilter() yfilter.YFilter { return vlanstatistics.YFilter }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) SetFilter(yf yfilter.YFilter) { vlanstatistics.YFilter = yf }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetGoName(yname string) string {
    if yname == "vlanStatsVlans" { return "Vlanstatsvlans" }
    if yname == "vlanStatsExtendedVlans" { return "Vlanstatsextendedvlans" }
    if yname == "vlanStatsInternalVlans" { return "Vlanstatsinternalvlans" }
    if yname == "vlanStatsFreeVlans" { return "Vlanstatsfreevlans" }
    return ""
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetSegmentPath() string {
    return "vlanStatistics"
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlanStatsVlans"] = vlanstatistics.Vlanstatsvlans
    leafs["vlanStatsExtendedVlans"] = vlanstatistics.Vlanstatsextendedvlans
    leafs["vlanStatsInternalVlans"] = vlanstatistics.Vlanstatsinternalvlans
    leafs["vlanStatsFreeVlans"] = vlanstatistics.Vlanstatsfreevlans
    return leafs
}

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetBundleName() string { return "cisco_ios_xe" }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetYangName() string { return "vlanStatistics" }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) SetParent(parent types.Entity) { vlanstatistics.parent = parent }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetParent() types.Entity { return vlanstatistics.parent }

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Managementdomaintable
// The table containing information on the management domains
// in which the local system is participating.  Devices which
// support only one management domain will support just one row
// in this table, and will not let it be deleted nor let other
// rows be created.  Devices which support multiple management
// domains will allow rows to be created and deleted, but will
// not allow the last row to be deleted. If the device does 
// not support VTP, the table is read-only.
type CISCOVTPMIB_Managementdomaintable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about the status of one management domain. The type is slice of
    // CISCOVTPMIB_Managementdomaintable_Managementdomainentry.
    Managementdomainentry []CISCOVTPMIB_Managementdomaintable_Managementdomainentry
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetFilter() yfilter.YFilter { return managementdomaintable.YFilter }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) SetFilter(yf yfilter.YFilter) { managementdomaintable.YFilter = yf }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetGoName(yname string) string {
    if yname == "managementDomainEntry" { return "Managementdomainentry" }
    return ""
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetSegmentPath() string {
    return "managementDomainTable"
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "managementDomainEntry" {
        for _, c := range managementdomaintable.Managementdomainentry {
            if managementdomaintable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Managementdomaintable_Managementdomainentry{}
        managementdomaintable.Managementdomainentry = append(managementdomaintable.Managementdomainentry, child)
        return &managementdomaintable.Managementdomainentry[len(managementdomaintable.Managementdomainentry)-1]
    }
    return nil
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range managementdomaintable.Managementdomainentry {
        children[managementdomaintable.Managementdomainentry[i].GetSegmentPath()] = &managementdomaintable.Managementdomainentry[i]
    }
    return children
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetBundleName() string { return "cisco_ios_xe" }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetYangName() string { return "managementDomainTable" }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) SetParent(parent types.Entity) { managementdomaintable.parent = parent }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetParent() types.Entity { return managementdomaintable.parent }

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry
// Information about the status of one management domain.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary value to uniquely identify the
    // management domain on the local system. The type is interface{} with range:
    // 1..255.
    Managementdomainindex interface{}

    // The management name of a domain in which the local system is participating.
    // The zero-length name corresponds to the 'no management-domain' state which
    // is the initial value at installation-time if not configured otherwise. 
    // Note that the zero-length name does not correspond to an operational
    // management domain, and a device does not send VTP advertisements while in
    // the 'no management-domain' state.  A device leaves the 'no
    // management-domain' state when it obtains a management-domain name, either
    // through configuration or through inheriting the management-domain name from
    // a received VTP advertisement.  When the value of an existing instance of
    // this object is modified by network management, the local system should re-
    // initialize its VLAN information (for the given management domain) as if it
    // had just been configured with a management domain name at installation
    // time. The type is string with length: 0..32.
    Managementdomainname interface{}

    // The local VTP mode in this management domain when
    // managementDomainVersionInUse is version1(1) or version2(2).  If
    // managementDomainVersionInUse is version3(4), this  object has the same
    // value with vtpDatabaseLocalMode  of VLAN database type.  - 'client'
    // indicates that the local system is acting   as a VTP client.  - 'server'
    // indicates that the local system is acting   as a VTP server.  -
    // 'transparent' indicates that the local system does   not generate or listen
    // to VTP messages, but forwards   messages. This mode can also be set by the
    // device   itself when the amount of VLAN information is too   large for it
    // to hold in DRAM.  - 'off' indicates that the local system does not  
    // generate, listen to or forward any VTP messages. The type is
    // Managementdomainlocalmode.
    Managementdomainlocalmode interface{}

    // The current Configuration Revision Number as known by the local device for
    // this management domain when  managementDomainVersionInUse is version1(1) or
    // version2(2).  If managementDomainVersionInUse is version3(4), this  object
    // has the same value with vtpDatabaseRevisionNumber  of VLAN database type. 
    // This value is updated (if necessary) whenever a VTP advertisement is
    // received or generated. When in the 'no management-domain' state, this value
    // is 0. The type is interface{} with range: 0..4294967295.
    Managementdomainconfigrevnumber interface{}

    // The IP-address (or one of them) of the VTP Server which last updated the
    // Configuration Revision Number, as indicated in the most recently received
    // VTP advertisement for this management domain, when
    // managementDomainVersionInUse is version1(1) or version2(2).   If
    // managementDomainVersionInUse is version3(4), this object has the value of
    // 0.0.0.0.  Before an advertisement has been received, this value is 0.0.0.0.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Managementdomainlastupdater interface{}

    // The time at which the Configuration Revision Number was (last) increased to
    // its current value, as indicated in the most recently received VTP
    // advertisement for this management domain when managementDomainVersionInUse
    // is not version3(4) or in the most recently received VTP VLAN database 
    // advertisement for this management domain when  managementDomainVersionInUse
    // is version3(4).  The value 0x0000010100000000 indicates that the device
    // which last increased the Configuration Revision Number had no idea of the
    // date/time, or that no advertisement has been received. The type is string.
    Managementdomainlastchange interface{}

    // The status of this conceptual row. The type is RowStatus.
    Managementdomainrowstatus interface{}

    // The IP address of a TFTP Server in/from which VTP VLAN information for this
    // management domain is to be stored/retrieved.  If the information is being
    // locally stored in NVRAM, this object should take the value 0.0.0.0. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Managementdomaintftpserver interface{}

    // The complete pathname of the file at the TFTP Server identified by the
    // value of managementDomainTftpServer in/from which VTP VLAN information for
    // this management domain is to be stored/retrieved.  If the value of
    // corresponding instance of managementDomainTftpServer is 0.0.0.0, the value
    // of this object is ignored. The type is string.
    Managementdomaintftppathname interface{}

    // An indication of whether VTP pruning is enabled or disabled in this
    // managament domain.   This object can only be modified, either when the 
    // corresponding instance value of managementDomainVersionInUse  is 'version1'
    // or 'version2' and the corresponding instance  value of
    // managementDomainLocalMode is 'server', or when the  corresponding instance
    // value of managementDomainVersionInUse  is 'version3' and the corresponding
    // instance value of  managementDomainLocalMode is 'server' or 'client'. The
    // type is Managementdomainpruningstate.
    Managementdomainpruningstate interface{}

    // The current version of the VTP that is in use by the designated management
    // domain.   This object can be set to none(3) only when  vtpVersion is
    // none(3). The type is Managementdomainversioninuse.
    Managementdomainversioninuse interface{}

    // Indicates whether VTP pruning is operationally enabled or disabled in this
    // managament domain. The type is Managementdomainpruningstateoper.
    Managementdomainpruningstateoper interface{}

    // The object specifies the interface to be used as the preferred source
    // interface for the VTP IP updater address.  A zero length value indicates
    // that a source interface is not specified. The type is string.
    Managementdomainadminsrcif interface{}

    // The object specifies whether to use only the IP address of
    // managementDomainAdminSrcIf as the VTP IP updater address.   'true'
    // indicates to only use the IP address of         managementDomainAdminSrcIf
    // as the VTP IP         updater address.   'false' indicates to use the IP
    // address of          managementDomainAdminSrcIf as the VTP IP         
    // updater address if managementDomainAdminSrcIf          is configured with
    // an IP address.  Otherwise, the          first available IP address of the
    // system will         be used. The type is bool.
    Managementdomainsourceonlymode interface{}

    // The object indicates the interface used as the preferred source interface
    // for the VTP IP updater address.  A zero length string indicates that a
    // source interface is not available. The type is string.
    Managementdomainopersrcif interface{}

    // The object specifies the file name where VTP configuration is stored in the
    // format of <filename> or <devices>:[<filename>].  <device> can be (but not
    // limited to): flash, bootflash, slot0, slot1, disk0. The type is string.
    Managementdomainconfigfile interface{}

    // The object indicates the type of the Internet address of the preferred
    // source interface for the VTP IP updater.  The value of this object is
    // 'unknown' if managementDomainVersionInUse is 'version3' or
    // managementDomainLocalMode is not 'server'. The type is InetAddressType.
    Managementdomainlocalupdatertype interface{}

    // The object indicates the Internet address of the preferred source interface
    // for the VTP IP updater. The type is string with length: 0..255.
    Managementdomainlocalupdater interface{}

    // The object indicates a value that uniquely identifies this device within a
    // VTP Domain.  The value of this object is zero length string if
    // managementDomainVersionInUse is not 'version3'. The type is string.
    Managementdomaindeviceid interface{}

    // This object always has the value 'none' when read.  When written, each
    // value causes the appropriate action:   'copy' - causes the creation of rows
    // in the vtpVlanEditTable exactly corresponding to the current global VLAN
    // information for this management domain.  If the Edit Buffer (for this
    // management domain) is not currently empty, a copy operation fails.  A
    // successful copy operation starts the deadman-timer.   'apply' - first
    // performs a consistent check on the the modified information contained in
    // the Edit Buffer, and if consistent, then tries to instanciate the modified
    // information as the new global VLAN information.  Note that an empty Edit
    // Buffer (for the management domain) would always result in an inconsistency
    // since the default VLANs are required to be present.   'release' - flushes
    // the Edit Buffer (for this management domain), clears the Owner information,
    // and aborts the deadman-timer.  A release is generated automatically if the
    // deadman-timer ever expires.   'restartTimer' - restarts the deadman-timer. 
    // 'none' - no operation is performed. The type is Vtpvlaneditoperation.
    Vtpvlaneditoperation interface{}

    // The current status of an 'apply' operation to instanciate the Edit Buffer
    // as the new global VLAN information (for this management domain).  If no
    // apply is currently active, the status represented is that of the most
    // recently completed apply.  The possible values are:     inProgress -
    // 'apply' operation in progress;     succeeded - the 'apply' was successful
    // (this value is           also used when no apply has been invoked since the
    // last time the local system restarted);     configNumberError - the apply
    // failed because the value of           vtpVlanEditConfigRevNumber was less
    // or equal to           the value of current value of           
    // managementDomainConfigRevNumber;     inconsistentEdit - the apply failed
    // because the modified           information was not self-consistent;    
    // tooBig - the apply failed because the modified           information was
    // too large to fit in this VTP           Server's non-volatile storage
    // location;     localNVStoreFail - the apply failed in trying to store       
    // the new information in a local non-volatile           storage location;    
    // remoteNVStoreFail - the apply failed in trying to store           the new
    // information in a remote non-volatile           storage location;    
    // editBufferEmpty - the apply failed because the Edit           Buffer was
    // empty (for this management domain).     someOtherError - the apply failed
    // for some other reason           (e.g., insufficient memory).    
    // notPrimaryServer - the apply failed because the local            device is
    // not a VTP primary server for VLAN            database type when
    // managementDomainVersionInUse           is version3(4). The type is
    // Vtpvlanapplystatus.
    Vtpvlanapplystatus interface{}

    // The management station which is currently using the Edit Buffer for this
    // management domain.  When the Edit Buffer for a management domain is not
    // currently in use, the value of this object is the zero-length string.  Note
    // that it is also the zero-length string if a manager fails to set this
    // object when invoking a copy operation. The type is string with length:
    // 0..127.
    Vtpvlaneditbufferowner interface{}

    // The Configuration Revision Number to be used for the next apply operation. 
    // This value is initialized (by the agent) on a copy operation to be one
    // greater than the value of managementDomainConfigRevNumber. On an apply, if
    // the  number is less or equal to the value of 
    // managementDomainConfigRevNumber, then the apply fails. The value can be
    // modified (increased) by network management before an apply to ensure that
    // an apply does not fail for  this reason.  This object is used to allow
    // management control over whether a configuration revision received via a VTP
    // advertisement after a copy operation but before the succeeding apply
    // operation is lost by being overwritten by the (local) edit operation.  By
    // default, the apply operation will fail in this situation.  By increasing
    // this object's value after the copy but before the apply, management can
    // control whether the apply is to succeed (with the update via VTP
    // advertisement being lost). The type is interface{} with range:
    // 0..4294967295.
    Vtpvlaneditconfigrevnumber interface{}

    // The VLAN-id of the modified VLAN in the Edit Buffer. If the object has the
    // value of zero, any VLAN can  be edited. If the value of the object is not
    // zero, only this VLAN can be edited.  The object's value is reset to zero
    // after a successful 'apply' operation or a 'release' operation.   This
    // object is only supported for devices which allow  only one VLAN editing for
    // each 'apply' operation. For devices which allow multiple VLAN editing for
    // each 'apply' operation, this object is not supported. The type is
    // interface{} with range: 0..4095.
    Vtpvlaneditmodifiedvlan interface{}

    // The total number of VTP Summary Adverts received for this management
    // domain. The type is interface{} with range: 0..4294967295.
    Vtpinsummaryadverts interface{}

    // The total number of VTP Subset Adverts received for this management domain.
    // The type is interface{} with range: 0..4294967295.
    Vtpinsubsetadverts interface{}

    // The total number of VTP Advert Requests received for this management
    // domain. The type is interface{} with range: 0..4294967295.
    Vtpinadvertrequests interface{}

    // The total number of VTP Summary Adverts sent for this management domain.
    // The type is interface{} with range: 0..4294967295.
    Vtpoutsummaryadverts interface{}

    // The total number of VTP Subset Adverts sent for this management domain. The
    // type is interface{} with range: 0..4294967295.
    Vtpoutsubsetadverts interface{}

    // The total number of VTP Advert Requests sent for this management domain.
    // The type is interface{} with range: 0..4294967295.
    Vtpoutadvertrequests interface{}

    // The number of occurrences of configuration revision number errors for this
    // management domain.  A configuration revision number error occurs when a
    // device receives a VTP advertisement for which:  - the advertisement's
    // Configuration Revision Number is the   same as the current locally-held
    // value, and  - the advertisement's digest value is different from the  
    // current locally-held value. The type is interface{} with range:
    // 0..4294967295.
    Vtpconfigrevnumbererrors interface{}

    // The number of occurrences of configuration digest errors for this
    // management domain.  A configuration digest error occurs when a device
    // receives a VTP advertisement for which:  - the advertisement's
    // Configuration Revision Number is   greater than the current locally-held
    // value, and  -  the advertisement's digest value computed by the  receiving
    // device does not match the checksum in the  summary advertisement that was
    // received earlier. This  can happen, for example, if there is a mismatch in
    // VTP  passwords between the VTP devices. The type is interface{} with range:
    // 0..4294967295.
    Vtpconfigdigesterrors interface{}
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetFilter() yfilter.YFilter { return managementdomainentry.YFilter }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) SetFilter(yf yfilter.YFilter) { managementdomainentry.YFilter = yf }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "managementDomainName" { return "Managementdomainname" }
    if yname == "managementDomainLocalMode" { return "Managementdomainlocalmode" }
    if yname == "managementDomainConfigRevNumber" { return "Managementdomainconfigrevnumber" }
    if yname == "managementDomainLastUpdater" { return "Managementdomainlastupdater" }
    if yname == "managementDomainLastChange" { return "Managementdomainlastchange" }
    if yname == "managementDomainRowStatus" { return "Managementdomainrowstatus" }
    if yname == "managementDomainTftpServer" { return "Managementdomaintftpserver" }
    if yname == "managementDomainTftpPathname" { return "Managementdomaintftppathname" }
    if yname == "managementDomainPruningState" { return "Managementdomainpruningstate" }
    if yname == "managementDomainVersionInUse" { return "Managementdomainversioninuse" }
    if yname == "managementDomainPruningStateOper" { return "Managementdomainpruningstateoper" }
    if yname == "managementDomainAdminSrcIf" { return "Managementdomainadminsrcif" }
    if yname == "managementDomainSourceOnlyMode" { return "Managementdomainsourceonlymode" }
    if yname == "managementDomainOperSrcIf" { return "Managementdomainopersrcif" }
    if yname == "managementDomainConfigFile" { return "Managementdomainconfigfile" }
    if yname == "managementDomainLocalUpdaterType" { return "Managementdomainlocalupdatertype" }
    if yname == "managementDomainLocalUpdater" { return "Managementdomainlocalupdater" }
    if yname == "managementDomainDeviceID" { return "Managementdomaindeviceid" }
    if yname == "vtpVlanEditOperation" { return "Vtpvlaneditoperation" }
    if yname == "vtpVlanApplyStatus" { return "Vtpvlanapplystatus" }
    if yname == "vtpVlanEditBufferOwner" { return "Vtpvlaneditbufferowner" }
    if yname == "vtpVlanEditConfigRevNumber" { return "Vtpvlaneditconfigrevnumber" }
    if yname == "vtpVlanEditModifiedVlan" { return "Vtpvlaneditmodifiedvlan" }
    if yname == "vtpInSummaryAdverts" { return "Vtpinsummaryadverts" }
    if yname == "vtpInSubsetAdverts" { return "Vtpinsubsetadverts" }
    if yname == "vtpInAdvertRequests" { return "Vtpinadvertrequests" }
    if yname == "vtpOutSummaryAdverts" { return "Vtpoutsummaryadverts" }
    if yname == "vtpOutSubsetAdverts" { return "Vtpoutsubsetadverts" }
    if yname == "vtpOutAdvertRequests" { return "Vtpoutadvertrequests" }
    if yname == "vtpConfigRevNumberErrors" { return "Vtpconfigrevnumbererrors" }
    if yname == "vtpConfigDigestErrors" { return "Vtpconfigdigesterrors" }
    return ""
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetSegmentPath() string {
    return "managementDomainEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", managementdomainentry.Managementdomainindex) + "']"
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = managementdomainentry.Managementdomainindex
    leafs["managementDomainName"] = managementdomainentry.Managementdomainname
    leafs["managementDomainLocalMode"] = managementdomainentry.Managementdomainlocalmode
    leafs["managementDomainConfigRevNumber"] = managementdomainentry.Managementdomainconfigrevnumber
    leafs["managementDomainLastUpdater"] = managementdomainentry.Managementdomainlastupdater
    leafs["managementDomainLastChange"] = managementdomainentry.Managementdomainlastchange
    leafs["managementDomainRowStatus"] = managementdomainentry.Managementdomainrowstatus
    leafs["managementDomainTftpServer"] = managementdomainentry.Managementdomaintftpserver
    leafs["managementDomainTftpPathname"] = managementdomainentry.Managementdomaintftppathname
    leafs["managementDomainPruningState"] = managementdomainentry.Managementdomainpruningstate
    leafs["managementDomainVersionInUse"] = managementdomainentry.Managementdomainversioninuse
    leafs["managementDomainPruningStateOper"] = managementdomainentry.Managementdomainpruningstateoper
    leafs["managementDomainAdminSrcIf"] = managementdomainentry.Managementdomainadminsrcif
    leafs["managementDomainSourceOnlyMode"] = managementdomainentry.Managementdomainsourceonlymode
    leafs["managementDomainOperSrcIf"] = managementdomainentry.Managementdomainopersrcif
    leafs["managementDomainConfigFile"] = managementdomainentry.Managementdomainconfigfile
    leafs["managementDomainLocalUpdaterType"] = managementdomainentry.Managementdomainlocalupdatertype
    leafs["managementDomainLocalUpdater"] = managementdomainentry.Managementdomainlocalupdater
    leafs["managementDomainDeviceID"] = managementdomainentry.Managementdomaindeviceid
    leafs["vtpVlanEditOperation"] = managementdomainentry.Vtpvlaneditoperation
    leafs["vtpVlanApplyStatus"] = managementdomainentry.Vtpvlanapplystatus
    leafs["vtpVlanEditBufferOwner"] = managementdomainentry.Vtpvlaneditbufferowner
    leafs["vtpVlanEditConfigRevNumber"] = managementdomainentry.Vtpvlaneditconfigrevnumber
    leafs["vtpVlanEditModifiedVlan"] = managementdomainentry.Vtpvlaneditmodifiedvlan
    leafs["vtpInSummaryAdverts"] = managementdomainentry.Vtpinsummaryadverts
    leafs["vtpInSubsetAdverts"] = managementdomainentry.Vtpinsubsetadverts
    leafs["vtpInAdvertRequests"] = managementdomainentry.Vtpinadvertrequests
    leafs["vtpOutSummaryAdverts"] = managementdomainentry.Vtpoutsummaryadverts
    leafs["vtpOutSubsetAdverts"] = managementdomainentry.Vtpoutsubsetadverts
    leafs["vtpOutAdvertRequests"] = managementdomainentry.Vtpoutadvertrequests
    leafs["vtpConfigRevNumberErrors"] = managementdomainentry.Vtpconfigrevnumbererrors
    leafs["vtpConfigDigestErrors"] = managementdomainentry.Vtpconfigdigesterrors
    return leafs
}

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetBundleName() string { return "cisco_ios_xe" }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetYangName() string { return "managementDomainEntry" }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) SetParent(parent types.Entity) { managementdomainentry.parent = parent }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetParent() types.Entity { return managementdomainentry.parent }

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetParentYangName() string { return "managementDomainTable" }

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode represents   generate, listen to or forward any VTP messages.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode_client CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode = "client"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode_server CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode = "server"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode_transparent CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode = "transparent"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode_off CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainlocalmode = "off"
)

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate represents managementDomainLocalMode is 'server' or 'client'.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate_enabled CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate = "enabled"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate_disabled CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstate = "disabled"
)

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper represents disabled in this managament domain.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper_enabled CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper = "enabled"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper_disabled CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainpruningstateoper = "disabled"
)

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse represents vtpVersion is none(3).
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse_version1 CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse = "version1"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse_version2 CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse = "version2"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse_none CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse = "none"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse_version3 CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainversioninuse = "version3"
)

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus represents           is version3(4).
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_inProgress CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "inProgress"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_succeeded CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "succeeded"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_configNumberError CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "configNumberError"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_inconsistentEdit CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "inconsistentEdit"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_tooBig CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "tooBig"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_localNVStoreFail CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "localNVStoreFail"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_remoteNVStoreFail CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "remoteNVStoreFail"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_editBufferEmpty CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "editBufferEmpty"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_someOtherError CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "someOtherError"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus_notPrimaryServer CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlanapplystatus = "notPrimaryServer"
)

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation represents  'none' - no operation is performed.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation string

const (
    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation_none CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation = "none"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation_copy CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation = "copy"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation_apply CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation = "apply"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation_release CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation = "release"

    CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation_restartTimer CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Vtpvlaneditoperation = "restartTimer"
)

// CISCOVTPMIB_Vtpvlantable
// This table contains information on the VLANs which
// currently exist.
type CISCOVTPMIB_Vtpvlantable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about one current VLAN.  The managementDomainIndex value in the
    // INDEX clause indicates which management domain the VLAN is in. The type is
    // slice of CISCOVTPMIB_Vtpvlantable_Vtpvlanentry.
    Vtpvlanentry []CISCOVTPMIB_Vtpvlantable_Vtpvlanentry
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetFilter() yfilter.YFilter { return vtpvlantable.YFilter }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) SetFilter(yf yfilter.YFilter) { vtpvlantable.YFilter = yf }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetGoName(yname string) string {
    if yname == "vtpVlanEntry" { return "Vtpvlanentry" }
    return ""
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetSegmentPath() string {
    return "vtpVlanTable"
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpVlanEntry" {
        for _, c := range vtpvlantable.Vtpvlanentry {
            if vtpvlantable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpvlantable_Vtpvlanentry{}
        vtpvlantable.Vtpvlanentry = append(vtpvlantable.Vtpvlanentry, child)
        return &vtpvlantable.Vtpvlanentry[len(vtpvlantable.Vtpvlanentry)-1]
    }
    return nil
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpvlantable.Vtpvlanentry {
        children[vtpvlantable.Vtpvlanentry[i].GetSegmentPath()] = &vtpvlantable.Vtpvlanentry[i]
    }
    return children
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetYangName() string { return "vtpVlanTable" }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) SetParent(parent types.Entity) { vtpvlantable.parent = parent }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetParent() types.Entity { return vtpvlantable.parent }

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpvlantable_Vtpvlanentry
// Information about one current VLAN.  The
// managementDomainIndex value in the INDEX clause indicates
// which management domain the VLAN is in.
type CISCOVTPMIB_Vtpvlantable_Vtpvlanentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. The VLAN-id of this VLAN on ISL or 802.1q trunks.
    // The type is interface{} with range: 0..4095.
    Vtpvlanindex interface{}

    // The state of this VLAN.  The state 'mtuTooBigForDevice' indicates that this
    // device cannot participate in this VLAN because the VLAN's MTU is larger
    // than the device can support.  The state 'mtuTooBigForTrunk' indicates that
    // while this VLAN's MTU is supported by this device, it is too large for one
    // or more of the device's trunk ports. The type is Vtpvlanstate.
    Vtpvlanstate interface{}

    // The type of this VLAN. The type is VlanType.
    Vtpvlantype interface{}

    // The name of this VLAN.  This name is used as the ELAN-name for an ATM
    // LAN-Emulation segment of this VLAN. The type is string with length: 1..32.
    Vtpvlanname interface{}

    // The MTU size on this VLAN, defined as the size of largest MAC-layer
    // (information field portion of the) data frame which can be transmitted on
    // the VLAN. The type is interface{} with range: 1500..18190.
    Vtpvlanmtu interface{}

    // The value of the 802.10 SAID field for this VLAN. The type is string with
    // length: 4.
    Vtpvlandot10Said interface{}

    // The ring number of this VLAN.  This object is only instantiated when the
    // value of the corresponding instance of vtpVlanType has a value of 'fddi' or
    // 'tokenRing' and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 0..4095.
    Vtpvlanringnumber interface{}

    // The bridge number of the VTP-capable switches for this VLAN.  This object
    // is only instantiated for VLANs that are involved with emulating token ring
    // segments. The type is interface{} with range: 0..15.
    Vtpvlanbridgenumber interface{}

    // The type of the Spanning Tree Protocol (STP) running on this VLAN.  This
    // object is only instanciated when the value of the corresponding instance of
    // vtpVlanType has a value of 'fddiNet' or 'trNet'.  The value returned by
    // this object depends upon the value of the corresponding instance of
    // vtpVlanEditStpType.  - 'ieee' indicates IEEE STP is running exclusively.  -
    // 'ibm' indicates IBM STP is running exclusively.  - 'hybrid' indicates a STP
    // that allows a combination of   IEEE and IBM is running.  The 'hybrid' STP
    // type results from tokenRing/fddi VLANs that are children of this
    // trNet/fddiNet parent VLAN being configured in a combination of SRT and SRB
    // vtpVlanBridgeTypes while the instance of vtpVlanEditStpType that
    // corresponds to this object is set to 'auto'. The type is Vtpvlanstptype.
    Vtpvlanstptype interface{}

    // The parent VLAN for this VLAN.  This object is only instantiated when the
    // value of the corresponding instance of vtpVlanType has a value of 'fddi' or
    // 'tokenRing' and Source Routing is in use on this VLAN.  The parent VLAN
    // must have  a vtpVlanType value of fddiNet(4) or trNet(5),  respectively.
    // The type is interface{} with range: 0..4095.
    Vtpvlanparentvlan interface{}

    // A VLAN to which this VLAN is being translational-bridged. If this value and
    // the corresponding instance of vtpVlanTranslationalVlan2 are both zero, then
    // this VLAN is not being translational-bridged. The type is interface{} with
    // range: 0..4095.
    Vtpvlantranslationalvlan1 interface{}

    // Another VLAN, i.e., other than that indicated by vtpVlanTranslationalVlan1,
    // to which this VLAN is being translational-bridged.  If this value and the
    // corresponding instance of vtpVlanTranslationalVlan1 are both zero, then
    // this VLAN is not being translational-bridged. The type is interface{} with
    // range: 0..4095.
    Vtpvlantranslationalvlan2 interface{}

    // The type of the Source Route bridging mode in use on this VLAN.  This
    // object is only instantiated when the value of  the corresponding instance
    // of vtpVlanType has a value of  fddi(2) or tokenRing(3) and Source Routing
    // is in use on this VLAN. The type is Vtpvlanbridgetype.
    Vtpvlanbridgetype interface{}

    // The maximum number of bridge hops allowed in All Routes Explorer frames on
    // this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    Vtpvlanarehopcount interface{}

    // The maximum number of bridge hops allowed in Spanning Tree Explorer frames
    // on this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    Vtpvlanstehopcount interface{}

    // True if this VLAN is of type trCrf and also is acting as a backup trCrf for
    // the ISL distributed BRF. The type is bool.
    Vtpvlaniscrfbackup interface{}

    // The additional type information of this VLAN. The type is map[string]bool.
    Vtpvlantypeext interface{}

    // The value of the ifIndex corresponding to this VLAN ID. If the VLAN ID does
    // not have its corresponding interface,  this object has the value of zero.
    // The type is interface{} with range: 0..2147483647.
    Vtpvlanifindex interface{}

    // The MISTP instance, to which the corresponding vlan is mapped. If this
    // value of this mib object is 0,  the corresponding vlan  is not configured
    // to be mapped to any MISTP instance and all the ports under this VLAN remain
    // in blocking state. The type is interface{} with range: 0..256.
    Stpxvlanmistpinstmapinstindex interface{}
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetFilter() yfilter.YFilter { return vtpvlanentry.YFilter }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) SetFilter(yf yfilter.YFilter) { vtpvlanentry.YFilter = yf }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpVlanIndex" { return "Vtpvlanindex" }
    if yname == "vtpVlanState" { return "Vtpvlanstate" }
    if yname == "vtpVlanType" { return "Vtpvlantype" }
    if yname == "vtpVlanName" { return "Vtpvlanname" }
    if yname == "vtpVlanMtu" { return "Vtpvlanmtu" }
    if yname == "vtpVlanDot10Said" { return "Vtpvlandot10Said" }
    if yname == "vtpVlanRingNumber" { return "Vtpvlanringnumber" }
    if yname == "vtpVlanBridgeNumber" { return "Vtpvlanbridgenumber" }
    if yname == "vtpVlanStpType" { return "Vtpvlanstptype" }
    if yname == "vtpVlanParentVlan" { return "Vtpvlanparentvlan" }
    if yname == "vtpVlanTranslationalVlan1" { return "Vtpvlantranslationalvlan1" }
    if yname == "vtpVlanTranslationalVlan2" { return "Vtpvlantranslationalvlan2" }
    if yname == "vtpVlanBridgeType" { return "Vtpvlanbridgetype" }
    if yname == "vtpVlanAreHopCount" { return "Vtpvlanarehopcount" }
    if yname == "vtpVlanSteHopCount" { return "Vtpvlanstehopcount" }
    if yname == "vtpVlanIsCRFBackup" { return "Vtpvlaniscrfbackup" }
    if yname == "vtpVlanTypeExt" { return "Vtpvlantypeext" }
    if yname == "vtpVlanIfIndex" { return "Vtpvlanifindex" }
    if yname == "stpxVlanMISTPInstMapInstIndex" { return "Stpxvlanmistpinstmapinstindex" }
    return ""
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetSegmentPath() string {
    return "vtpVlanEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlanentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpvlanentry.Vtpvlanindex) + "']"
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpvlanentry.Managementdomainindex
    leafs["vtpVlanIndex"] = vtpvlanentry.Vtpvlanindex
    leafs["vtpVlanState"] = vtpvlanentry.Vtpvlanstate
    leafs["vtpVlanType"] = vtpvlanentry.Vtpvlantype
    leafs["vtpVlanName"] = vtpvlanentry.Vtpvlanname
    leafs["vtpVlanMtu"] = vtpvlanentry.Vtpvlanmtu
    leafs["vtpVlanDot10Said"] = vtpvlanentry.Vtpvlandot10Said
    leafs["vtpVlanRingNumber"] = vtpvlanentry.Vtpvlanringnumber
    leafs["vtpVlanBridgeNumber"] = vtpvlanentry.Vtpvlanbridgenumber
    leafs["vtpVlanStpType"] = vtpvlanentry.Vtpvlanstptype
    leafs["vtpVlanParentVlan"] = vtpvlanentry.Vtpvlanparentvlan
    leafs["vtpVlanTranslationalVlan1"] = vtpvlanentry.Vtpvlantranslationalvlan1
    leafs["vtpVlanTranslationalVlan2"] = vtpvlanentry.Vtpvlantranslationalvlan2
    leafs["vtpVlanBridgeType"] = vtpvlanentry.Vtpvlanbridgetype
    leafs["vtpVlanAreHopCount"] = vtpvlanentry.Vtpvlanarehopcount
    leafs["vtpVlanSteHopCount"] = vtpvlanentry.Vtpvlanstehopcount
    leafs["vtpVlanIsCRFBackup"] = vtpvlanentry.Vtpvlaniscrfbackup
    leafs["vtpVlanTypeExt"] = vtpvlanentry.Vtpvlantypeext
    leafs["vtpVlanIfIndex"] = vtpvlanentry.Vtpvlanifindex
    leafs["stpxVlanMISTPInstMapInstIndex"] = vtpvlanentry.Stpxvlanmistpinstmapinstindex
    return leafs
}

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetYangName() string { return "vtpVlanEntry" }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) SetParent(parent types.Entity) { vtpvlanentry.parent = parent }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetParent() types.Entity { return vtpvlanentry.parent }

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetParentYangName() string { return "vtpVlanTable" }

// CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype represents this VLAN.
type CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype string

const (
    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype_none CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype = "none"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype_srt CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype = "srt"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype_srb CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanbridgetype = "srb"
)

// CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate represents one or more of the device's trunk ports.
type CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate string

const (
    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate_operational CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate = "operational"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate_suspended CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate = "suspended"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate_mtuTooBigForDevice CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate = "mtuTooBigForDevice"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate_mtuTooBigForTrunk CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstate = "mtuTooBigForTrunk"
)

// CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype represents to 'auto'.
type CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype string

const (
    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype_ieee CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype = "ieee"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype_ibm CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype = "ibm"

    CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype_hybrid CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanstptype = "hybrid"
)

// CISCOVTPMIB_Vtpinternalvlantable
// A vtpInternalVlanTable entry contains
// information on an existing internal
// VLAN. It is internally created by the
// device for a specific application program 
// and hence owned by the application.  
// It cannot be modified or deleted by (local 
// or network) management.
type CISCOVTPMIB_Vtpinternalvlantable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about one current internal VLAN. The type is slice of
    // CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry.
    Vtpinternalvlanentry []CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetFilter() yfilter.YFilter { return vtpinternalvlantable.YFilter }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) SetFilter(yf yfilter.YFilter) { vtpinternalvlantable.YFilter = yf }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetGoName(yname string) string {
    if yname == "vtpInternalVlanEntry" { return "Vtpinternalvlanentry" }
    return ""
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetSegmentPath() string {
    return "vtpInternalVlanTable"
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpInternalVlanEntry" {
        for _, c := range vtpinternalvlantable.Vtpinternalvlanentry {
            if vtpinternalvlantable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry{}
        vtpinternalvlantable.Vtpinternalvlanentry = append(vtpinternalvlantable.Vtpinternalvlanentry, child)
        return &vtpinternalvlantable.Vtpinternalvlanentry[len(vtpinternalvlantable.Vtpinternalvlanentry)-1]
    }
    return nil
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpinternalvlantable.Vtpinternalvlanentry {
        children[vtpinternalvlantable.Vtpinternalvlanentry[i].GetSegmentPath()] = &vtpinternalvlantable.Vtpinternalvlanentry[i]
    }
    return children
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetYangName() string { return "vtpInternalVlanTable" }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) SetParent(parent types.Entity) { vtpinternalvlantable.parent = parent }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetParent() types.Entity { return vtpinternalvlantable.parent }

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry
// Information about one current internal
// VLAN.
type CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanindex
    Vtpvlanindex interface{}

    // The program name of the internal VLAN's owner application. This internal
    // VLAN is allocated by the device specifically for this application and no
    // one else could create, modify or delete this  VLAN. The type is string.
    Vtpinternalvlanowner interface{}
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetFilter() yfilter.YFilter { return vtpinternalvlanentry.YFilter }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) SetFilter(yf yfilter.YFilter) { vtpinternalvlanentry.YFilter = yf }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpVlanIndex" { return "Vtpvlanindex" }
    if yname == "vtpInternalVlanOwner" { return "Vtpinternalvlanowner" }
    return ""
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetSegmentPath() string {
    return "vtpInternalVlanEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpinternalvlanentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpinternalvlanentry.Vtpvlanindex) + "']"
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpinternalvlanentry.Managementdomainindex
    leafs["vtpVlanIndex"] = vtpinternalvlanentry.Vtpvlanindex
    leafs["vtpInternalVlanOwner"] = vtpinternalvlanentry.Vtpinternalvlanowner
    return leafs
}

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetYangName() string { return "vtpInternalVlanEntry" }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) SetParent(parent types.Entity) { vtpinternalvlanentry.parent = parent }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetParent() types.Entity { return vtpinternalvlanentry.parent }

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetParentYangName() string { return "vtpInternalVlanTable" }

// CISCOVTPMIB_Vtpvlanedittable
// The table which contains the information in the Edit
// Buffers, one Edit Buffer per management domain.  The
// information for a particular management domain is
// initialized, by a 'copy' operation, to be the current global
// VLAN information for that management domain.  After
// initialization, editing can be performed to add VLANs,
// delete VLANs, or modify their global parameters.  The
// information as modified through editing is local to this
// Edit Buffer.  An apply operation using the
// vtpVlanEditOperation object is necessary to instanciate the
// modified information as the new global VLAN information for
// that management domain.
// 
// To use the Edit Buffer, a manager acts as follows:
// 
// 1. ensures the Edit Buffer for a management domain is empty,
// i.e., there are no rows in this table for this management
// domain.
// 
// 2. issues a SNMP set operation which sets
// vtpVlanEditOperation to 'copy', and vtpVlanEditBufferOwner
// to its own identifier (e.g., its own IP address).
// 
// 3. if this set operation is successful, proceeds to edit the
// information in the vtpVlanEditTable.
// 
// 4. if and when the edited information is to be instantiated,
// issues a SNMP set operation which sets vtpVlanEditOperation
// to 'apply'.
// 
// 5. issues retrieval requests to obtain the value of
// vtpVlanApplyStatus, until the result of the apply is
// determined.
// 
// 6. releases the Edit Buffer by issuing a SNMP set operation
// which sets vtpVlanEditOperation to 'release'.
// 
// Note that the information contained in this table is not
// saved across agent reboots.
type CISCOVTPMIB_Vtpvlanedittable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about one VLAN in the Edit Buffer for a particular management
    // domain. The type is slice of CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry.
    Vtpvlaneditentry []CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetFilter() yfilter.YFilter { return vtpvlanedittable.YFilter }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) SetFilter(yf yfilter.YFilter) { vtpvlanedittable.YFilter = yf }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetGoName(yname string) string {
    if yname == "vtpVlanEditEntry" { return "Vtpvlaneditentry" }
    return ""
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetSegmentPath() string {
    return "vtpVlanEditTable"
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpVlanEditEntry" {
        for _, c := range vtpvlanedittable.Vtpvlaneditentry {
            if vtpvlanedittable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry{}
        vtpvlanedittable.Vtpvlaneditentry = append(vtpvlanedittable.Vtpvlaneditentry, child)
        return &vtpvlanedittable.Vtpvlaneditentry[len(vtpvlanedittable.Vtpvlaneditentry)-1]
    }
    return nil
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpvlanedittable.Vtpvlaneditentry {
        children[vtpvlanedittable.Vtpvlaneditentry[i].GetSegmentPath()] = &vtpvlanedittable.Vtpvlaneditentry[i]
    }
    return children
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetYangName() string { return "vtpVlanEditTable" }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) SetParent(parent types.Entity) { vtpvlanedittable.parent = parent }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetParent() types.Entity { return vtpvlanedittable.parent }

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry
// Information about one VLAN in the Edit Buffer for a
// particular management domain.
type CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. The VLAN-id which this VLAN would have on ISL or
    // 802.1q trunks. The type is interface{} with range: 0..4095.
    Vtpvlaneditindex interface{}

    // The state which this VLAN would have. The type is Vtpvlaneditstate.
    Vtpvlaneditstate interface{}

    // The type which this VLAN would have. An implementation may restrict access
    // to this object. The type is VlanType.
    Vtpvlanedittype interface{}

    // The name which this VLAN would have.  This name would be used as the
    // ELAN-name for an ATM LAN-Emulation segment of this VLAN.  An implementation
    // may restrict access to this object. The type is string with length: 1..32.
    Vtpvlaneditname interface{}

    // The MTU size which this VLAN would have, defined as the size of largest
    // MAC-layer (information field portion of the) data frame which can be
    // transmitted on the VLAN.  An implementation may restrict access to this
    // object. The type is interface{} with range: 1500..18190.
    Vtpvlaneditmtu interface{}

    // The value of the 802.10 SAID field which would be used for this VLAN.  An
    // implementation may restrict access to this object. The type is string with
    // length: 4.
    Vtpvlaneditdot10Said interface{}

    // The ring number which would be used for this VLAN.  This object is only
    // instantiated when the value of the corresponding instance of
    // vtpVlanEditType has a value of 'fddi' or 'tokenRing' and Source Routing is
    // in use on  this VLAN. The type is interface{} with range: 0..4095.
    Vtpvlaneditringnumber interface{}

    // The bridge number of the VTP-capable switches which would be used for this
    // VLAN.  This object is only instantiated when the value of the corresponding
    // instance of vtpVlanEditType has a value of fddiNet(4) or trNet(5). The type
    // is interface{} with range: 0..15.
    Vtpvlaneditbridgenumber interface{}

    // The type of the Spanning Tree Protocol which would be running on this VLAN.
    // This object is only instantiated when the value of the corresponding
    // instance of vtpVlanEditType has a value of fddiNet(4) or trNet(5).  If
    // 'ieee' is selected, the STP that runs will be IEEE.  If 'ibm' is selected,
    // the STP that runs will be IBM.  If 'auto' is selected, the STP that runs
    // will be dependant on the values of vtpVlanEditBridgeType for all children
    // tokenRing/fddi type VLANs.  This will result in a 'hybrid' STP (see
    // vtpVlanStpType). The type is Vtpvlaneditstptype.
    Vtpvlaneditstptype interface{}

    // The VLAN index of the VLAN which would be the parent for this VLAN.  This
    // object is only instantiated when the value of the corresponding instance of
    // vtpVlanEditType has a value of 'fddi' or 'tokenRing' and Source Routing is
    // in use on this VLAN.  The parent VLAN must have a vtpVlanEditType  value of
    // fddiNet(4) or trNet(5), respectively. The type is interface{} with range:
    // 0..4095.
    Vtpvlaneditparentvlan interface{}

    // The status of this row.  Any and all columnar objects in an existing row
    // can be modified irrespective of the status of the row.  A row is not
    // qualified for activation until instances of at least its vtpVlanEditType,
    // vtpVlanEditName and vtpVlanEditDot10Said columns have appropriate values. 
    // The management station should endeavor to make all rows consistent in the
    // table before 'apply'ing the buffer.  An inconsistent entry in the table
    // will cause the entire buffer to be rejected with the vtpVlanApplyStatus
    // object set to the appropriate error value. The type is RowStatus.
    Vtpvlaneditrowstatus interface{}

    // A VLAN to which this VLAN would be translational-bridged. If this value and
    // the corresponding instance of vtpVlanTranslationalVlan2 are both zero, then
    // this VLAN would not be translational-bridged.  An implementation may
    // restrict access to this object. The type is interface{} with range:
    // 0..4095.
    Vtpvlanedittranslationalvlan1 interface{}

    // Another VLAN, i.e., other than that indicated by
    // vtpVlanEditTranslationalVlan1, to which this VLAN would be
    // translational-bridged.  If this value and the corresponding instance of
    // vtpVlanTranslationalVlan1 are both zero, then this VLAN would not be
    // translational-bridged.  An implementation may restrict access to this
    // object. The type is interface{} with range: 0..4095.
    Vtpvlanedittranslationalvlan2 interface{}

    // The type of Source Route bridging mode which would be in use on this VLAN. 
    // This object is only instantiated when  the value of  the corresponding
    // instance of vtpVlanEditType has a value of fddi(2) or tokenRing(3) and
    // Source Routing  is in use on this VLAN. The type is Vtpvlaneditbridgetype.
    Vtpvlaneditbridgetype interface{}

    // The maximum number of bridge hops allowed in All Routes Explorer frames on
    // this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    Vtpvlaneditarehopcount interface{}

    // The maximum number of bridge hops allowed in Spanning Tree Explorer frames
    // on this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    Vtpvlaneditstehopcount interface{}

    // True if this VLAN is of type trCrf and also is acting as a backup trCrf for
    // the ISL distributed BRF.  This object is only instantiated when the value
    // of the corresponding instance of vtpVlanEditType has a value of
    // tokenRing(3). The type is bool.
    Vtpvlaneditiscrfbackup interface{}

    // The additional type information of this VLAN. vtpVlanEditTypeExt object is
    // superseded by vtpVlanEditTypeExt2. The type is map[string]bool.
    Vtpvlanedittypeext interface{}

    // The additional type information of this VLAN. The VlanTypeExt TC specifies
    // which bits may be written by a management application. The agent should
    // provide a default value. The type is map[string]bool.
    Vtpvlanedittypeext2 interface{}

    // The MISTP instance, to which the corresponding vlan would be  mapped. The
    // value of this mib object is from 0 to the value of stpxMISTPInstanceNumber.
    // If setting the value of this object to 0, the corresponding vlan will not
    // be mapped to a MISTP  instance and all the ports under this VLAN will be
    // moved into the blocking state. The type is interface{} with range: 0..256.
    Stpxvlanmistpinstmapeditinstindex interface{}
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetFilter() yfilter.YFilter { return vtpvlaneditentry.YFilter }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) SetFilter(yf yfilter.YFilter) { vtpvlaneditentry.YFilter = yf }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpVlanEditIndex" { return "Vtpvlaneditindex" }
    if yname == "vtpVlanEditState" { return "Vtpvlaneditstate" }
    if yname == "vtpVlanEditType" { return "Vtpvlanedittype" }
    if yname == "vtpVlanEditName" { return "Vtpvlaneditname" }
    if yname == "vtpVlanEditMtu" { return "Vtpvlaneditmtu" }
    if yname == "vtpVlanEditDot10Said" { return "Vtpvlaneditdot10Said" }
    if yname == "vtpVlanEditRingNumber" { return "Vtpvlaneditringnumber" }
    if yname == "vtpVlanEditBridgeNumber" { return "Vtpvlaneditbridgenumber" }
    if yname == "vtpVlanEditStpType" { return "Vtpvlaneditstptype" }
    if yname == "vtpVlanEditParentVlan" { return "Vtpvlaneditparentvlan" }
    if yname == "vtpVlanEditRowStatus" { return "Vtpvlaneditrowstatus" }
    if yname == "vtpVlanEditTranslationalVlan1" { return "Vtpvlanedittranslationalvlan1" }
    if yname == "vtpVlanEditTranslationalVlan2" { return "Vtpvlanedittranslationalvlan2" }
    if yname == "vtpVlanEditBridgeType" { return "Vtpvlaneditbridgetype" }
    if yname == "vtpVlanEditAreHopCount" { return "Vtpvlaneditarehopcount" }
    if yname == "vtpVlanEditSteHopCount" { return "Vtpvlaneditstehopcount" }
    if yname == "vtpVlanEditIsCRFBackup" { return "Vtpvlaneditiscrfbackup" }
    if yname == "vtpVlanEditTypeExt" { return "Vtpvlanedittypeext" }
    if yname == "vtpVlanEditTypeExt2" { return "Vtpvlanedittypeext2" }
    if yname == "stpxVlanMISTPInstMapEditInstIndex" { return "Stpxvlanmistpinstmapeditinstindex" }
    return ""
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetSegmentPath() string {
    return "vtpVlanEditEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlaneditentry.Managementdomainindex) + "']" + "[vtpVlanEditIndex='" + fmt.Sprintf("%v", vtpvlaneditentry.Vtpvlaneditindex) + "']"
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpvlaneditentry.Managementdomainindex
    leafs["vtpVlanEditIndex"] = vtpvlaneditentry.Vtpvlaneditindex
    leafs["vtpVlanEditState"] = vtpvlaneditentry.Vtpvlaneditstate
    leafs["vtpVlanEditType"] = vtpvlaneditentry.Vtpvlanedittype
    leafs["vtpVlanEditName"] = vtpvlaneditentry.Vtpvlaneditname
    leafs["vtpVlanEditMtu"] = vtpvlaneditentry.Vtpvlaneditmtu
    leafs["vtpVlanEditDot10Said"] = vtpvlaneditentry.Vtpvlaneditdot10Said
    leafs["vtpVlanEditRingNumber"] = vtpvlaneditentry.Vtpvlaneditringnumber
    leafs["vtpVlanEditBridgeNumber"] = vtpvlaneditentry.Vtpvlaneditbridgenumber
    leafs["vtpVlanEditStpType"] = vtpvlaneditentry.Vtpvlaneditstptype
    leafs["vtpVlanEditParentVlan"] = vtpvlaneditentry.Vtpvlaneditparentvlan
    leafs["vtpVlanEditRowStatus"] = vtpvlaneditentry.Vtpvlaneditrowstatus
    leafs["vtpVlanEditTranslationalVlan1"] = vtpvlaneditentry.Vtpvlanedittranslationalvlan1
    leafs["vtpVlanEditTranslationalVlan2"] = vtpvlaneditentry.Vtpvlanedittranslationalvlan2
    leafs["vtpVlanEditBridgeType"] = vtpvlaneditentry.Vtpvlaneditbridgetype
    leafs["vtpVlanEditAreHopCount"] = vtpvlaneditentry.Vtpvlaneditarehopcount
    leafs["vtpVlanEditSteHopCount"] = vtpvlaneditentry.Vtpvlaneditstehopcount
    leafs["vtpVlanEditIsCRFBackup"] = vtpvlaneditentry.Vtpvlaneditiscrfbackup
    leafs["vtpVlanEditTypeExt"] = vtpvlaneditentry.Vtpvlanedittypeext
    leafs["vtpVlanEditTypeExt2"] = vtpvlaneditentry.Vtpvlanedittypeext2
    leafs["stpxVlanMISTPInstMapEditInstIndex"] = vtpvlaneditentry.Stpxvlanmistpinstmapeditinstindex
    return leafs
}

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetYangName() string { return "vtpVlanEditEntry" }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) SetParent(parent types.Entity) { vtpvlaneditentry.parent = parent }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetParent() types.Entity { return vtpvlaneditentry.parent }

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetParentYangName() string { return "vtpVlanEditTable" }

// CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype represents is in use on this VLAN.
type CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype string

const (
    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype_srt CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype = "srt"

    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype_srb CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditbridgetype = "srb"
)

// CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate represents The state which this VLAN would have.
type CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate string

const (
    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate_operational CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate = "operational"

    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate_suspended CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstate = "suspended"
)

// CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype represents a 'hybrid' STP (see vtpVlanStpType).
type CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype string

const (
    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype_ieee CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype = "ieee"

    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype_ibm CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype = "ibm"

    CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype_auto CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry_Vtpvlaneditstptype = "auto"
)

// CISCOVTPMIB_Vtpvlanlocalshutdowntable
// Ths table contains the VLAN local shutdown
// information within management domain.
type CISCOVTPMIB_Vtpvlanlocalshutdowntable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing VLAN local shutdown information for a particular VLAN
    // in the management domain.  An entry is created if a VLAN which supports
    // local shutdown has been created.  An entry is deleted if a VLAN which
    // supports local shutdown has been removed. The type is slice of
    // CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry.
    Vtpvlanlocalshutdownentry []CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetFilter() yfilter.YFilter { return vtpvlanlocalshutdowntable.YFilter }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) SetFilter(yf yfilter.YFilter) { vtpvlanlocalshutdowntable.YFilter = yf }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetGoName(yname string) string {
    if yname == "vtpVlanLocalShutdownEntry" { return "Vtpvlanlocalshutdownentry" }
    return ""
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetSegmentPath() string {
    return "vtpVlanLocalShutdownTable"
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpVlanLocalShutdownEntry" {
        for _, c := range vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry {
            if vtpvlanlocalshutdowntable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry{}
        vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry = append(vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry, child)
        return &vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry[len(vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry)-1]
    }
    return nil
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry {
        children[vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry[i].GetSegmentPath()] = &vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry[i]
    }
    return children
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetYangName() string { return "vtpVlanLocalShutdownTable" }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) SetParent(parent types.Entity) { vtpvlanlocalshutdowntable.parent = parent }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetParent() types.Entity { return vtpvlanlocalshutdowntable.parent }

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry
// An entry containing VLAN local shutdown information for a
// particular VLAN in the management domain.
// 
// An entry is created if a VLAN which supports local shutdown
// has been created.
// 
// An entry is deleted if a VLAN which supports local shutdown
// has been removed.
type CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Vtpvlantable_Vtpvlanentry_Vtpvlanindex
    Vtpvlanindex interface{}

    // The object specifies the VLAN local shutdown state. The type is
    // Vtpvlanlocalshutdown.
    Vtpvlanlocalshutdown interface{}
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetFilter() yfilter.YFilter { return vtpvlanlocalshutdownentry.YFilter }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) SetFilter(yf yfilter.YFilter) { vtpvlanlocalshutdownentry.YFilter = yf }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpVlanIndex" { return "Vtpvlanindex" }
    if yname == "vtpVlanLocalShutdown" { return "Vtpvlanlocalshutdown" }
    return ""
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetSegmentPath() string {
    return "vtpVlanLocalShutdownEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlanlocalshutdownentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpvlanlocalshutdownentry.Vtpvlanindex) + "']"
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpvlanlocalshutdownentry.Managementdomainindex
    leafs["vtpVlanIndex"] = vtpvlanlocalshutdownentry.Vtpvlanindex
    leafs["vtpVlanLocalShutdown"] = vtpvlanlocalshutdownentry.Vtpvlanlocalshutdown
    return leafs
}

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetYangName() string { return "vtpVlanLocalShutdownEntry" }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) SetParent(parent types.Entity) { vtpvlanlocalshutdownentry.parent = parent }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetParent() types.Entity { return vtpvlanlocalshutdownentry.parent }

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetParentYangName() string { return "vtpVlanLocalShutdownTable" }

// CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown represents The object specifies the VLAN local shutdown state.
type CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown string

const (
    CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown_up CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown = "up"

    CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown_down CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry_Vtpvlanlocalshutdown = "down"
)

// CISCOVTPMIB_Vlantrunkporttable
// The table containing information on the local system's VLAN
// trunk ports.
type CISCOVTPMIB_Vlantrunkporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about one trunk port. The type is slice of
    // CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry.
    Vlantrunkportentry []CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetFilter() yfilter.YFilter { return vlantrunkporttable.YFilter }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) SetFilter(yf yfilter.YFilter) { vlantrunkporttable.YFilter = yf }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetGoName(yname string) string {
    if yname == "vlanTrunkPortEntry" { return "Vlantrunkportentry" }
    return ""
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetSegmentPath() string {
    return "vlanTrunkPortTable"
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlanTrunkPortEntry" {
        for _, c := range vlantrunkporttable.Vlantrunkportentry {
            if vlantrunkporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry{}
        vlantrunkporttable.Vlantrunkportentry = append(vlantrunkporttable.Vlantrunkportentry, child)
        return &vlantrunkporttable.Vlantrunkportentry[len(vlantrunkporttable.Vlantrunkportentry)-1]
    }
    return nil
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vlantrunkporttable.Vlantrunkportentry {
        children[vlantrunkporttable.Vlantrunkportentry[i].GetSegmentPath()] = &vlantrunkporttable.Vlantrunkportentry[i]
    }
    return children
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetBundleName() string { return "cisco_ios_xe" }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetYangName() string { return "vlanTrunkPortTable" }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) SetParent(parent types.Entity) { vlantrunkporttable.parent = parent }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetParent() types.Entity { return vlantrunkporttable.parent }

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry
// Information about one trunk port.
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of ifIndex for the interface
    // corresponding to this trunk port. The type is interface{} with range:
    // 1..2147483647.
    Vlantrunkportifindex interface{}

    // The value of managementDomainIndex for the management domain on this trunk
    // port.  Devices which support only one management domain will support this
    // object read-only. The type is interface{} with range: 1..255.
    Vlantrunkportmanagementdomain interface{}

    // The type of VLAN encapsulation desired to be used on this trunk port. It is
    // either a particular type, or 'negotiate' meaning whatever type results from
    // the negotiation. negotiate(5) is not allowed if the port does not support
    // negotiation or if its vlanTrunkPortDynamicState is set to on(1) or
    // onNoNegotiate(5). Whether writing to this object in order to modify the
    // encapsulation is supported is both device and interface specific. The type
    // is Vlantrunkportencapsulationtype.
    Vlantrunkportencapsulationtype interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then the local system is enabled for sending and receiving frames on that
    // VLAN; if the bit is set to '0', then the system is disabled from sending
    // and receiving frames on that VLAN.  To avoid conflicts between overlapping
    // partial updates by multiple managers, i.e., updates which modify only a
    // portion of an instance of this object (e.g., enable/disable a single VLAN
    // on the trunk port), any SNMP Set operation accessing an instance of this
    // object should also write the value of vlanTrunkPortSetSerialNo. The type is
    // string with length: 128.
    Vlantrunkportvlansenabled interface{}

    // The VlanIndex of the VLAN which is represented by native frames on this
    // trunk port.  For trunk ports not supporting the sending and receiving of
    // native frames, this value should be set to zero. The type is interface{}
    // with range: 0..4095.
    Vlantrunkportnativevlan interface{}

    // The status of this row.  In some circumstances, the creation of a row in
    // this table is needed to enable the appropriate trunking/tagging protocol on
    // the port, to enable the use of VTP on the port, and to assign the port to
    // the appropriate management domain.  In other circumstances, rows in this
    // table will be created as a by-product of other operations. The type is
    // RowStatus.
    Vlantrunkportrowstatus interface{}

    // The number of VTP Join messages received on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    Vlantrunkportinjoins interface{}

    // The number of VTP Join messages sent on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    Vlantrunkportoutjoins interface{}

    // The number of VTP Advertisement messages which indicated the sender does
    // not support VLAN-pruning received on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    Vlantrunkportoldadverts interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then the local system is permitted to prune that VLAN on this trunk port;
    // if the bit is set to '0', then the system must not prune that VLAN on this
    // trunk port.  To avoid conflicts between overlapping partial updates by
    // multiple managers, i.e., updates which modify only a portion of an instance
    // of this object (e.g., enable/disable a single VLAN on the trunk port), any
    // SNMP Set operation accessing an instance of this object should also write
    // the value of vlanTrunkPortSetSerialNo. The type is string with length: 128.
    Vlantrunkportvlanspruningeligible interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then this VLAN is presently being forwarded on this trunk port, i.e., it is
    // not pruned; if the bit is set to '0', then this VLAN is presently not being
    // forwarded on this trunk port, either because it is pruned or for some other
    // reason. The type is string with length: 128.
    Vlantrunkportvlansxmitjoined interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then the local switch is currently sending joins for this VLAN on this
    // trunk port, i.e., it is asking to receive frames for this VLAN; if the bit
    // is set to '0', then the local switch is not currently sending joins for
    // this VLAN on this trunk port. The type is string with length: 128.
    Vlantrunkportvlansrcvjoined interface{}

    // For devices that allows dynamic determination of whether a link between two
    // switches should be a trunk or not, this object allows the operator to
    // mandate the behavior of that dynamic mechanism.  on(1) dictates that the
    // interface will always be a trunk. This is the value for static entries
    // (those that show no dynamic behavior). If the negotiation is supported on
    // this port, negotiation will take place with the far end to attempt to bring
    // the far end into trunking state.  off(2) allows an operator to specify that
    // the specified interface is never to be trunk, regardless of any dynamic
    // mechanisms to the contrary.  This value is useful for overriding the
    // default behavior of some switches. If the negotiation is supported on this
    // port, negotiation will take place with the far end to attempt on the link
    // to bring the far end into non-trunking state.  desirable(3) is used to
    // indicate that it is desirable for the interface to become a trunk.  The
    // device will initiate any negotiation necessary to become a trunk but will
    // not become a trunk unless it receives confirmation from the far end on the
    // link.  auto(4) is used to indicate that the interface is capable and
    // willing to become a trunk but will not initiate trunking negotiations.  The
    // far end on the link are required to either start negotiations or start
    // sending encapsulated packets, on which event the specified interface will
    // become a trunk.  onNoNegotiate(5) is used to indicate that the interface is
    // permanently set to be a trunk, and no negotiation takes place with the far
    // end on the link to ensure consistent operation. This is similar to on(1)
    // except no negotiation takes place with the far end.  If the port does not
    // support negotiation or its vlanTrunkPortEncapsulationType is set to
    // negotiate(5), onNoNegotiate(5) is not allowed.  Devices that do no support
    // dynamic determination (for just a particular interface, encapsulation or
    // for the whole device) need only support the 'on', and 'off' values. The
    // type is Vlantrunkportdynamicstate.
    Vlantrunkportdynamicstate interface{}

    // Indicates whether the specified interface is either acting as a trunk or
    // not. This is a result of the vlanTrunkPortDynamicState and the ifOperStatus
    // of the trunk port itself. The type is Vlantrunkportdynamicstatus.
    Vlantrunkportdynamicstatus interface{}

    // Some trunk interface modules allow VTP to be enabled/disabled seperately
    // from that of the central device.  In such a case this object provides
    // management a way to remotely enable VTP on that module.  If a module does
    // not support a seperate VTP enabled state then this object shall always
    // return 'true' and will accept no other value during a SET operation. The
    // type is bool.
    Vlantrunkportvtpenabled interface{}

    // The type of VLAN encapsulation in use on this trunk port. For intefaces
    // with vlanTrunkPortDynamicStatus of notTrunking(2) the
    // vlanTrunkPortEncapsulationOperType shall be notApplicable(6). The type is
    // Vlantrunkportencapsulationopertype.
    Vlantrunkportencapsulationopertype interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 1024 through
    // 1031; the second octet to VLANs 1032 through 1039; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet. If the bit corresponding to a VLAN is set to '1', then the
    // local system is enabled for sending and receiving frames on that VLAN; if
    // the bit is set to '0', then the system is disabled from sending and
    // receiving frames on that VLAN. The default value is zero length string.  To
    // avoid conflicts between overlapping partial updates by multiple managers,
    // i.e., updates which modify only a portion of an instance of this object
    // (e.g., enable/disable a single VLAN on the trunk port), any SNMP Set
    // operation accessing an instance of this object should also write the value
    // of vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vlantrunkportvlansenabled2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 2048 through
    // 2055; the second octet to VLANs 2056 through 2063; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet. If the bit corresponding to a VLAN is set to '1', then the
    // local system is enabled for sending and receiving frames on that VLAN; if
    // the bit is set to '0', then the system is disabled from sending and
    // receiving frames on that VLAN. The default value is zero length string.  To
    // avoid conflicts between overlapping partial updates by multiple managers,
    // i.e., updates which modify only a portion of an instance of this object
    // (e.g., enable/disable a single VLAN on the trunk port), any SNMP Set
    // operation accessing an instance of this object should also write the value
    // of vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vlantrunkportvlansenabled3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 3072 through
    // 3079; the second octet to VLANs 3080 through 3087; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet. If the bit corresponding to a VLAN is set to '1', then the
    // local system is enabled for sending and receiving frames on that VLAN; if
    // the bit is set to '0', then the system is disabled from sending and
    // receiving frames on that VLAN. The default value is zero length string.  To
    // avoid conflicts between overlapping partial updates by multiple managers,
    // i.e., updates which modify only a portion of an instance of this object
    // (e.g., enable/disable a single VLAN on the trunk port), any SNMP Set
    // operation accessing an instance of this object should also write the value
    // of vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vlantrunkportvlansenabled4K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 1024 through
    // 1031; the second octet to VLANs 1032 through 1039; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local system is permitted to prune that VLAN on this trunk port; if the bit
    // is set to '0', then the system must not prune that VLAN on this trunk port.
    // The default value is zero length string.  To avoid conflicts between
    // overlapping partial updates by multiple managers, i.e., updates which
    // modify only a portion of an instance of this object (e.g., enable/disable a
    // single VLAN on the trunk port), any SNMP Set operation accessing an
    // instance of this object should also write the value of
    // vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vtpvlanspruningeligible2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 2048 through
    // 2055; the second octet to VLANs 2056 through 2063; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local system is permitted to prune that VLAN on this trunk port; if the bit
    // is set to '0', then the system must not prune that VLAN on this trunk port.
    // The default value is zero length string.  To avoid conflicts between
    // overlapping partial updates by multiple managers, i.e., updates which
    // modify only a portion of an instance of this object (e.g., enable/disable a
    // single VLAN on the trunk port), any SNMP Set operation accessing an
    // instance of this object should also write the value of
    // vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vtpvlanspruningeligible3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 3072 through
    // 3079; the second octet to VLANs 3080 through 3087; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local system is permitted to prune that VLAN on this trunk port; if the bit
    // is set to '0', then the system must not prune that VLAN on this trunk port.
    // The default value is zero length string.  To avoid conflicts between
    // overlapping partial updates by multiple managers, i.e., updates which
    // modify only a portion of an instance of this object (e.g., enable/disable a
    // single VLAN on the trunk port), any SNMP Set operation accessing an
    // instance of this object should also write the value of
    // vlanTrunkPortSetSerialNo. The type is string with length: 0..128.
    Vtpvlanspruningeligible4K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 1024 through
    // 1031; the second octet to VLANs 1032 through 1039; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then this
    // VLAN is presently being forwarded on this trunk port, i.e., it is not
    // pruned; if the bit is set to '0', then this VLAN is presently not being
    // forwarded on this trunk port, either because it is pruned or for some other
    // reason. The type is string with length: 0..128.
    Vlantrunkportvlansxmitjoined2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 2048 through
    // 2055; the second octet to VLANs 2056 through 2063; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then this
    // VLAN is presently being forwarded on this trunk port, i.e., it is not
    // pruned; if the bit is set to '0', then this VLAN is presently not being
    // forwarded on this trunk port, either because it is pruned or for some other
    // reason. The type is string with length: 0..128.
    Vlantrunkportvlansxmitjoined3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 3072 through
    // 3079; the second octet to VLANs 3080 through 3087; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then this
    // VLAN is presently being forwarded on this trunk port, i.e., it is not
    // pruned; if the bit is set to '0', then this VLAN is presently not being
    // forwarded on this trunk port, either because it is pruned or for some other
    // reason. The type is string with length: 0..128.
    Vlantrunkportvlansxmitjoined4K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 1024 through
    // 1031; the second octet to VLANs 1032 through 1039; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local switch is currently sending joins for this VLAN on this trunk port,
    // i.e., it is asking to receive frames for this VLAN; if the bit is set to
    // '0', then the local switch is not currently sending joins for this VLAN on
    // this trunk port. The type is string with length: 0..128.
    Vlantrunkportvlansrcvjoined2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 2048 through
    // 2055; the second octet to VLANs 2056 through 2063; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local switch is currently sending joins for this VLAN on this trunk port,
    // i.e., it is asking to receive frames for this VLAN; if the bit is set to
    // '0', then the local switch is not currently sending joins for this VLAN on
    // this trunk port. The type is string with length: 0..128.
    Vlantrunkportvlansrcvjoined3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095 in the management domain on this trunk port. 
    // The first octet corresponds to VLANs with VlanIndex values of 3072 through
    // 3079; the second octet to VLANs 3080 through 3087; etc.  The most
    // significant bit of each octet corresponds to the lowest value VlanIndex in
    // that octet.  If the bit corresponding to a VLAN is set to '1', then the
    // local switch is currently sending joins for this VLAN on this trunk port,
    // i.e., it is asking to receive frames for this VLAN; if the bit is set to
    // '0', then the local switch is not currently sending joins for this VLAN on
    // this trunk port. The type is string with length: 0..128.
    Vlantrunkportvlansrcvjoined4K interface{}

    // Indicates dot1qtunnel mode of the port.  If the portDot1qTunnel  is set to
    // 'trunk' mode, the port's vlanTrunkPortDynamicState will be changed to
    // 'onNoNegotiate' and the vlanTrunkPortEncapsulationType will be set to
    // 'dot1Q'. These values cannot be changed unless dot1q tunnel is disabled on
    // this port.  If the portDot1qTunnel mode is set to 'access' mode, the port's
    // vlanTrunkPortDynamicState will be set to 'off'.And the value of
    // vlanTrunkPortDynamicState cannot be changed unless dot1q tunnel is disabled
    // on this port. 1Q packets received on this access port will remain.  Setting
    // the port to dot1q tunnel 'disabled' mode causes the dot1q tunnel feature to
    // be disabled on this port.  This object can't be set to 'trunk' or 'access'
    // mode, when vlanTrunkPortsDot1qTag  object is set to 'false'.  This object
    // has been deprecated and is replaced by the object 'cltcDot1qTunnelMode' in
    // the CISCO-L2-TUNNEL-CONFIG-MIB. The type is Vlantrunkportdot1Qtunnel.
    Vlantrunkportdot1Qtunnel interface{}

    // A string of octets containing one bit per VLAN with VlanIndex values of 0
    // through 2047.  If the bit corresponding to a VLAN is set to 1, it indicates
    // that vlan is allowed and active in management domain.  If the bit
    // corresponding to a VLAN is set to 0, it indicates that vlan is not allowed
    // or not active in management domain. The type is string with length: 0..256.
    Vlantrunkportvlansactivefirst2K interface{}

    // A string of octets containing one bit per VLAN with VlanIndex values of
    // 2048 through 4095.  If the bit corresponding to a VLAN is set to 1, it
    // indicates that vlan is allowed and active in management domain.  If the bit
    // corresponding to a VLAN is set to 0, it indicates that vlan is not allowed
    // or not active in management domain. The type is string with length: 0..256.
    Vlantrunkportvlansactivesecond2K interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  For each VLAN, if it is preferred on this trunk
    // port, then the bit corresponding to that VLAN is set to '1'. The default
    // value is 128 bytes of zeros.  To avoid conflicts between overlapping
    // partial updates by multiple managers, i.e., updates which modify only a
    // portion of an instance of this object (e.g., enable/disable a single VLAN
    // on the trunk port), any SNMP Set operation accessing an instance of this
    // object should also write the value of vlanTrunkPortSetSerialNo. The type is
    // string with length: 128.
    Stpxpreferredvlansmap interface{}

    // A string of octets containing one bit per VLAN for VLANS  with VlanIndex
    // values of 1024 through 2047 in the management domain on this trunk port. 
    // The first octet corresponds to  VLANs with VlanIndex values of 1024 through
    // 1031;  the second octet to VLANs 1032 through 1039; etc.  The most
    // significant bit of each octet corresponds to the  lowest value VlanIndex in
    // that octet.   For each VLAN, if it is preferred on this trunk port, then 
    // the bit corresponding to that VLAN is set to '1'.  The default value is 128
    // bytes of zeros.   To avoid conflicts between overlapping partial updates by
    // multiple managers, i.e., updates which modify only a portion of an instance
    // of this object (e.g., enable/disable a single VLAN on the trunk port), any
    // SNMP Set operation accessing an instance of this object should also write
    // the value of  vlanTrunkPortSetSerialNo. The type is string with length:
    // 0..128.
    Stpxpreferredvlansmap2K interface{}

    // A string of octets containing one bit per VLAN for VLANS  with VlanIndex
    // values of 2048 through 3071 in the management domain on this trunk port. 
    // The first octet corresponds to  VLANs with VlanIndex values of 2048 through
    // 2055;  the second octet to VLANs 2056 through 2063; etc.  The most
    // significant bit of each octet corresponds to the  lowest value VlanIndex in
    // that octet.   For each VLAN, if it is preferred on this trunk port, then 
    // the bit corresponding to that VLAN is set to '1'.  The default value is 128
    // bytes of zeros.   To avoid conflicts between overlapping partial updates by
    // multiple managers, i.e., updates which modify only a portion of an instance
    // of this object (e.g., enable/disable a single VLAN on the trunk port), any
    // SNMP Set operation accessing an instance of this object should also write
    // the value of  vlanTrunkPortSetSerialNo. The type is string with length:
    // 0..128.
    Stpxpreferredvlansmap3K interface{}

    // A string of octets containing one bit per VLAN for VLANS  with VlanIndex
    // values of 3072 through 4095 in the management domain on this trunk port. 
    // The first octet corresponds to  VLANs with VlanIndex values of 3072 through
    // 3079;  the second octet to VLANs 3080 through 3087; etc.  The most
    // significant bit of each octet corresponds to the  lowest value VlanIndex in
    // that octet.   For each VLAN, if it is preferred on this trunk port, then 
    // the bit corresponding to that VLAN is set to '1'.  The default value is 128
    // bytes of zeros.   To avoid conflicts between overlapping partial updates by
    // multiple managers, i.e., updates which modify only a portion of an instance
    // of this object (e.g., enable/disable a single VLAN on the trunk port), any
    // SNMP Set operation accessing an instance of this object should also write
    // the value of  vlanTrunkPortSetSerialNo. The type is string with length:
    // 0..128.
    Stpxpreferredvlansmap4K interface{}

    // A string of octets containing one bit per MISTP instances  in the
    // management domain on this trunk port. The first octet corresponds to MISTP
    // instances with InstIndex values of 1  through 8; the second octet to MISTP
    // instances 9 through 16; etc. The most significant bit of each octet
    // corresponds to  the lowest value instanceIndex in that octet. The number of
    // bits for this mib object will be determined by the value of 
    // stpxMISTPInstanceNumber.  For each instance, if it is preferred on this
    // trunk port, then the bit corresponding to that instance is set to '1'.   To
    // avoid conflicts between overlapping partial updates by multiple managers,
    // i.e., updates which modify only a portion of an instance of this object
    // (e.g., enable/disable a single instance on the trunk port), any SNMP Set
    // operation  accessing an instance of this object should also write the 
    // value of vlanTrunkPortSetSerialNo. The type is string with length: 1..32.
    Stpxpreferredmistpinstancesmap interface{}

    // A string of octets containing one bit per MST instances  on this trunk
    // port.  The first octet corresponds to MST  instances of 0 through 7; the
    // second octet to MST instances  8 through 15; etc. The most significant bit
    // of each octet  corresponds to the lowest MST instance value in that octet. 
    // The number of bits for this mib object will be determined  by the value of
    // stpxMSTMaxInstanceNumber.  For each instance, if it is preferred on this
    // trunk port,  then the bit corresponding to that instance is set to '1'. The
    // type is string with length: 1..32.
    Stpxpreferredmstinstancesmap interface{}
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetFilter() yfilter.YFilter { return vlantrunkportentry.YFilter }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) SetFilter(yf yfilter.YFilter) { vlantrunkportentry.YFilter = yf }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetGoName(yname string) string {
    if yname == "vlanTrunkPortIfIndex" { return "Vlantrunkportifindex" }
    if yname == "vlanTrunkPortManagementDomain" { return "Vlantrunkportmanagementdomain" }
    if yname == "vlanTrunkPortEncapsulationType" { return "Vlantrunkportencapsulationtype" }
    if yname == "vlanTrunkPortVlansEnabled" { return "Vlantrunkportvlansenabled" }
    if yname == "vlanTrunkPortNativeVlan" { return "Vlantrunkportnativevlan" }
    if yname == "vlanTrunkPortRowStatus" { return "Vlantrunkportrowstatus" }
    if yname == "vlanTrunkPortInJoins" { return "Vlantrunkportinjoins" }
    if yname == "vlanTrunkPortOutJoins" { return "Vlantrunkportoutjoins" }
    if yname == "vlanTrunkPortOldAdverts" { return "Vlantrunkportoldadverts" }
    if yname == "vlanTrunkPortVlansPruningEligible" { return "Vlantrunkportvlanspruningeligible" }
    if yname == "vlanTrunkPortVlansXmitJoined" { return "Vlantrunkportvlansxmitjoined" }
    if yname == "vlanTrunkPortVlansRcvJoined" { return "Vlantrunkportvlansrcvjoined" }
    if yname == "vlanTrunkPortDynamicState" { return "Vlantrunkportdynamicstate" }
    if yname == "vlanTrunkPortDynamicStatus" { return "Vlantrunkportdynamicstatus" }
    if yname == "vlanTrunkPortVtpEnabled" { return "Vlantrunkportvtpenabled" }
    if yname == "vlanTrunkPortEncapsulationOperType" { return "Vlantrunkportencapsulationopertype" }
    if yname == "vlanTrunkPortVlansEnabled2k" { return "Vlantrunkportvlansenabled2K" }
    if yname == "vlanTrunkPortVlansEnabled3k" { return "Vlantrunkportvlansenabled3K" }
    if yname == "vlanTrunkPortVlansEnabled4k" { return "Vlantrunkportvlansenabled4K" }
    if yname == "vtpVlansPruningEligible2k" { return "Vtpvlanspruningeligible2K" }
    if yname == "vtpVlansPruningEligible3k" { return "Vtpvlanspruningeligible3K" }
    if yname == "vtpVlansPruningEligible4k" { return "Vtpvlanspruningeligible4K" }
    if yname == "vlanTrunkPortVlansXmitJoined2k" { return "Vlantrunkportvlansxmitjoined2K" }
    if yname == "vlanTrunkPortVlansXmitJoined3k" { return "Vlantrunkportvlansxmitjoined3K" }
    if yname == "vlanTrunkPortVlansXmitJoined4k" { return "Vlantrunkportvlansxmitjoined4K" }
    if yname == "vlanTrunkPortVlansRcvJoined2k" { return "Vlantrunkportvlansrcvjoined2K" }
    if yname == "vlanTrunkPortVlansRcvJoined3k" { return "Vlantrunkportvlansrcvjoined3K" }
    if yname == "vlanTrunkPortVlansRcvJoined4k" { return "Vlantrunkportvlansrcvjoined4K" }
    if yname == "vlanTrunkPortDot1qTunnel" { return "Vlantrunkportdot1Qtunnel" }
    if yname == "vlanTrunkPortVlansActiveFirst2k" { return "Vlantrunkportvlansactivefirst2K" }
    if yname == "vlanTrunkPortVlansActiveSecond2k" { return "Vlantrunkportvlansactivesecond2K" }
    if yname == "stpxPreferredVlansMap" { return "Stpxpreferredvlansmap" }
    if yname == "stpxPreferredVlansMap2k" { return "Stpxpreferredvlansmap2K" }
    if yname == "stpxPreferredVlansMap3k" { return "Stpxpreferredvlansmap3K" }
    if yname == "stpxPreferredVlansMap4k" { return "Stpxpreferredvlansmap4K" }
    if yname == "stpxPreferredMISTPInstancesMap" { return "Stpxpreferredmistpinstancesmap" }
    if yname == "stpxPreferredMSTInstancesMap" { return "Stpxpreferredmstinstancesmap" }
    return ""
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetSegmentPath() string {
    return "vlanTrunkPortEntry" + "[vlanTrunkPortIfIndex='" + fmt.Sprintf("%v", vlantrunkportentry.Vlantrunkportifindex) + "']"
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlanTrunkPortIfIndex"] = vlantrunkportentry.Vlantrunkportifindex
    leafs["vlanTrunkPortManagementDomain"] = vlantrunkportentry.Vlantrunkportmanagementdomain
    leafs["vlanTrunkPortEncapsulationType"] = vlantrunkportentry.Vlantrunkportencapsulationtype
    leafs["vlanTrunkPortVlansEnabled"] = vlantrunkportentry.Vlantrunkportvlansenabled
    leafs["vlanTrunkPortNativeVlan"] = vlantrunkportentry.Vlantrunkportnativevlan
    leafs["vlanTrunkPortRowStatus"] = vlantrunkportentry.Vlantrunkportrowstatus
    leafs["vlanTrunkPortInJoins"] = vlantrunkportentry.Vlantrunkportinjoins
    leafs["vlanTrunkPortOutJoins"] = vlantrunkportentry.Vlantrunkportoutjoins
    leafs["vlanTrunkPortOldAdverts"] = vlantrunkportentry.Vlantrunkportoldadverts
    leafs["vlanTrunkPortVlansPruningEligible"] = vlantrunkportentry.Vlantrunkportvlanspruningeligible
    leafs["vlanTrunkPortVlansXmitJoined"] = vlantrunkportentry.Vlantrunkportvlansxmitjoined
    leafs["vlanTrunkPortVlansRcvJoined"] = vlantrunkportentry.Vlantrunkportvlansrcvjoined
    leafs["vlanTrunkPortDynamicState"] = vlantrunkportentry.Vlantrunkportdynamicstate
    leafs["vlanTrunkPortDynamicStatus"] = vlantrunkportentry.Vlantrunkportdynamicstatus
    leafs["vlanTrunkPortVtpEnabled"] = vlantrunkportentry.Vlantrunkportvtpenabled
    leafs["vlanTrunkPortEncapsulationOperType"] = vlantrunkportentry.Vlantrunkportencapsulationopertype
    leafs["vlanTrunkPortVlansEnabled2k"] = vlantrunkportentry.Vlantrunkportvlansenabled2K
    leafs["vlanTrunkPortVlansEnabled3k"] = vlantrunkportentry.Vlantrunkportvlansenabled3K
    leafs["vlanTrunkPortVlansEnabled4k"] = vlantrunkportentry.Vlantrunkportvlansenabled4K
    leafs["vtpVlansPruningEligible2k"] = vlantrunkportentry.Vtpvlanspruningeligible2K
    leafs["vtpVlansPruningEligible3k"] = vlantrunkportentry.Vtpvlanspruningeligible3K
    leafs["vtpVlansPruningEligible4k"] = vlantrunkportentry.Vtpvlanspruningeligible4K
    leafs["vlanTrunkPortVlansXmitJoined2k"] = vlantrunkportentry.Vlantrunkportvlansxmitjoined2K
    leafs["vlanTrunkPortVlansXmitJoined3k"] = vlantrunkportentry.Vlantrunkportvlansxmitjoined3K
    leafs["vlanTrunkPortVlansXmitJoined4k"] = vlantrunkportentry.Vlantrunkportvlansxmitjoined4K
    leafs["vlanTrunkPortVlansRcvJoined2k"] = vlantrunkportentry.Vlantrunkportvlansrcvjoined2K
    leafs["vlanTrunkPortVlansRcvJoined3k"] = vlantrunkportentry.Vlantrunkportvlansrcvjoined3K
    leafs["vlanTrunkPortVlansRcvJoined4k"] = vlantrunkportentry.Vlantrunkportvlansrcvjoined4K
    leafs["vlanTrunkPortDot1qTunnel"] = vlantrunkportentry.Vlantrunkportdot1Qtunnel
    leafs["vlanTrunkPortVlansActiveFirst2k"] = vlantrunkportentry.Vlantrunkportvlansactivefirst2K
    leafs["vlanTrunkPortVlansActiveSecond2k"] = vlantrunkportentry.Vlantrunkportvlansactivesecond2K
    leafs["stpxPreferredVlansMap"] = vlantrunkportentry.Stpxpreferredvlansmap
    leafs["stpxPreferredVlansMap2k"] = vlantrunkportentry.Stpxpreferredvlansmap2K
    leafs["stpxPreferredVlansMap3k"] = vlantrunkportentry.Stpxpreferredvlansmap3K
    leafs["stpxPreferredVlansMap4k"] = vlantrunkportentry.Stpxpreferredvlansmap4K
    leafs["stpxPreferredMISTPInstancesMap"] = vlantrunkportentry.Stpxpreferredmistpinstancesmap
    leafs["stpxPreferredMSTInstancesMap"] = vlantrunkportentry.Stpxpreferredmstinstancesmap
    return leafs
}

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetBundleName() string { return "cisco_ios_xe" }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetYangName() string { return "vlanTrunkPortEntry" }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) SetParent(parent types.Entity) { vlantrunkportentry.parent = parent }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetParent() types.Entity { return vlantrunkportentry.parent }

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetParentYangName() string { return "vlanTrunkPortTable" }

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel represents CISCO-L2-TUNNEL-CONFIG-MIB
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel string

const (
    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel_trunk CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel = "trunk"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel_access CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel = "access"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel_disabled CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdot1Qtunnel = "disabled"
)

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate represents device) need only support the 'on', and 'off' values.
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate string

const (
    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate_on CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate = "on"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate_off CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate = "off"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate_desirable CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate = "desirable"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate_auto CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate = "auto"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate_onNoNegotiate CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstate = "onNoNegotiate"
)

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus represents trunk port itself.
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus string

const (
    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus_trunking CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus = "trunking"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus_notTrunking CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportdynamicstatus = "notTrunking"
)

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype represents be notApplicable(6).
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype string

const (
    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_isl CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "isl"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_dot10 CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "dot10"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_lane CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "lane"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_dot1Q CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "dot1Q"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_negotiating CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "negotiating"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype_notApplicable CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationopertype = "notApplicable"
)

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype represents device and interface specific.
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype string

const (
    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype_isl CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype = "isl"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype_dot10 CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype = "dot10"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype_lane CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype = "lane"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype_dot1Q CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype = "dot1Q"

    CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype_negotiate CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry_Vlantrunkportencapsulationtype = "negotiate"
)

// CISCOVTPMIB_Vtpdiscovertable
// This table contains information related to the discovery
// of the VTP members in the designated management
// domain. This table is not instantiated when 
// managementDomainVersionInUse is version1(1), version2(3)
// or none(3).
type CISCOVTPMIB_Vtpdiscovertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information related to the discovery of the VTP members in one management
    // domain. The type is slice of CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry.
    Vtpdiscoverentry []CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetFilter() yfilter.YFilter { return vtpdiscovertable.YFilter }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) SetFilter(yf yfilter.YFilter) { vtpdiscovertable.YFilter = yf }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetGoName(yname string) string {
    if yname == "vtpDiscoverEntry" { return "Vtpdiscoverentry" }
    return ""
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetSegmentPath() string {
    return "vtpDiscoverTable"
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpDiscoverEntry" {
        for _, c := range vtpdiscovertable.Vtpdiscoverentry {
            if vtpdiscovertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry{}
        vtpdiscovertable.Vtpdiscoverentry = append(vtpdiscovertable.Vtpdiscoverentry, child)
        return &vtpdiscovertable.Vtpdiscoverentry[len(vtpdiscovertable.Vtpdiscoverentry)-1]
    }
    return nil
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpdiscovertable.Vtpdiscoverentry {
        children[vtpdiscovertable.Vtpdiscoverentry[i].GetSegmentPath()] = &vtpdiscovertable.Vtpdiscoverentry[i]
    }
    return children
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetYangName() string { return "vtpDiscoverTable" }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) SetParent(parent types.Entity) { vtpdiscovertable.parent = parent }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetParent() types.Entity { return vtpdiscovertable.parent }

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry
// Information related to the discovery of the
// VTP members in one management domain.
type CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // When this object is set to discover(1), all the entries in
    // vtpDiscoverResultTable for the corresponding management domain will be
    // removed  and the local device will begin to discover all VTP members in the
    // management domain. Upon the successful completion of discovery, the
    // discovered result will be stored in the vtpDiscoverResultTable.  If
    // vtpDiscoverStatus is inProgress(1), setting  vtpDiscoverAction to
    // discover(1) will fail.   When this object is set to purgeResult(3),  all
    // the entries of vtpDiscoverResultTable for  the corresponding management
    // domain will be removed from vtpDiscoverResultTable.  When this object is
    // set to noOperation(2), no action will be taken. When read, this object
    // always returns noOperation(2). The type is Vtpdiscoveraction.
    Vtpdiscoveraction interface{}

    // The current status of VTP discovery.  inProgress - a discovery is in
    // progress;  succeeded - the discovery was completed successfully            
    // (this value is also used when              no discover has been invoked
    // since the             last time the local system restarted); 
    // resourceUnavailable - the discovery failed because             the required
    // allocation of a resource is             presently unavailable. 
    // someOtherError - 'the discovery failed due to a             reason no
    // listed. The type is Vtpdiscoverstatus.
    Vtpdiscoverstatus interface{}

    // The value of sysUpTime at which the last discovery was completed.  A value
    // of zero indicates that no discovery has been invoked since last time the
    // local system restarted. The type is interface{} with range: 0..4294967295.
    Vtplastdiscovertime interface{}
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetFilter() yfilter.YFilter { return vtpdiscoverentry.YFilter }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) SetFilter(yf yfilter.YFilter) { vtpdiscoverentry.YFilter = yf }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpDiscoverAction" { return "Vtpdiscoveraction" }
    if yname == "vtpDiscoverStatus" { return "Vtpdiscoverstatus" }
    if yname == "vtpLastDiscoverTime" { return "Vtplastdiscovertime" }
    return ""
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetSegmentPath() string {
    return "vtpDiscoverEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdiscoverentry.Managementdomainindex) + "']"
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpdiscoverentry.Managementdomainindex
    leafs["vtpDiscoverAction"] = vtpdiscoverentry.Vtpdiscoveraction
    leafs["vtpDiscoverStatus"] = vtpdiscoverentry.Vtpdiscoverstatus
    leafs["vtpLastDiscoverTime"] = vtpdiscoverentry.Vtplastdiscovertime
    return leafs
}

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetYangName() string { return "vtpDiscoverEntry" }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) SetParent(parent types.Entity) { vtpdiscoverentry.parent = parent }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetParent() types.Entity { return vtpdiscoverentry.parent }

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetParentYangName() string { return "vtpDiscoverTable" }

// CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction represents always returns noOperation(2).
type CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction string

const (
    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction_discover CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction = "discover"

    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction_noOperation CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction = "noOperation"

    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction_purgeResult CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoveraction = "purgeResult"
)

// CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus represents             reason no listed.
type CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus string

const (
    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus_inProgress CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus = "inProgress"

    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus_succeeded CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus = "succeeded"

    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus_resourceUnavailable CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus = "resourceUnavailable"

    CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus_someOtherError CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry_Vtpdiscoverstatus = "someOtherError"
)

// CISCOVTPMIB_Vtpdiscoverresulttable
// The table containing information of discovered VTP members
// in the management domain in which the local system is
// participating. This table is not instantiated when 
// managementDomainVersionInUse is version1(1), version2(2) or
// none(3).
type CISCOVTPMIB_Vtpdiscoverresulttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row is created for each VTP member which is found through
    // successful discovery. The type is slice of
    // CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry.
    Vtpdiscoverresultentry []CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetFilter() yfilter.YFilter { return vtpdiscoverresulttable.YFilter }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) SetFilter(yf yfilter.YFilter) { vtpdiscoverresulttable.YFilter = yf }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetGoName(yname string) string {
    if yname == "vtpDiscoverResultEntry" { return "Vtpdiscoverresultentry" }
    return ""
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetSegmentPath() string {
    return "vtpDiscoverResultTable"
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpDiscoverResultEntry" {
        for _, c := range vtpdiscoverresulttable.Vtpdiscoverresultentry {
            if vtpdiscoverresulttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry{}
        vtpdiscoverresulttable.Vtpdiscoverresultentry = append(vtpdiscoverresulttable.Vtpdiscoverresultentry, child)
        return &vtpdiscoverresulttable.Vtpdiscoverresultentry[len(vtpdiscoverresulttable.Vtpdiscoverresultentry)-1]
    }
    return nil
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpdiscoverresulttable.Vtpdiscoverresultentry {
        children[vtpdiscoverresulttable.Vtpdiscoverresultentry[i].GetSegmentPath()] = &vtpdiscoverresulttable.Vtpdiscoverresultentry[i]
    }
    return children
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetYangName() string { return "vtpDiscoverResultTable" }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) SetParent(parent types.Entity) { vtpdiscoverresulttable.parent = parent }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetParent() types.Entity { return vtpdiscoverresulttable.parent }

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry
// A conceptual row is created for each VTP member which
// is found through successful discovery.
type CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. A value assigned by the system which identifies a
    // VTP member and the associated database in the  management domain. The type
    // is interface{} with range: 0..4294967295.
    Vtpdiscoverresultindex interface{}

    // The database name associated with the discovered VTP member. The type is
    // string with length: 0..64.
    Vtpdiscoverresultdatabasename interface{}

    // Indicates whether this VTP member contains conflicting information. 
    // true(1) indicates that this member has conflicting  information of the
    // database type in the management domain.  false(2) indicates that there is
    // no conflicting information of the database type in the management domain.
    // The type is bool.
    Vtpdiscoverresultconflicting interface{}

    // The unique identifier of the device for this VTP member. The type is string
    // with length: 0..64.
    Vtpdiscoverresultdeviceid interface{}

    // The unique identifier of the primary server for this VTP member and the
    // associated database type.  There are two different VTP servers, the primary
    // server and the secondary server.  When a local device is configured as a
    // server for a certain database type, it becomes secondary server by default.
    // Primary server is an operational role under which a server can initiate or
    // change the VTP configuration of the database type.  If this VTP member
    // itself is the primary server, the    value of this object is the same as
    // the value of  vtpDiscoverResultDeviceId of the instance. The type is string
    // with length: 0..64.
    Vtpdiscoverresultprimaryserver interface{}

    // The current configuration revision number as known by the VTP member. When
    // the database type is unknown for the VTP member, this value is 0. The type
    // is interface{} with range: 0..4294967295.
    Vtpdiscoverresultrevnumber interface{}

    // sysName of the VTP member. The type is string with length: 0..64.
    Vtpdiscoverresultsystemname interface{}
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetFilter() yfilter.YFilter { return vtpdiscoverresultentry.YFilter }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) SetFilter(yf yfilter.YFilter) { vtpdiscoverresultentry.YFilter = yf }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpDiscoverResultIndex" { return "Vtpdiscoverresultindex" }
    if yname == "vtpDiscoverResultDatabaseName" { return "Vtpdiscoverresultdatabasename" }
    if yname == "vtpDiscoverResultConflicting" { return "Vtpdiscoverresultconflicting" }
    if yname == "vtpDiscoverResultDeviceId" { return "Vtpdiscoverresultdeviceid" }
    if yname == "vtpDiscoverResultPrimaryServer" { return "Vtpdiscoverresultprimaryserver" }
    if yname == "vtpDiscoverResultRevNumber" { return "Vtpdiscoverresultrevnumber" }
    if yname == "vtpDiscoverResultSystemName" { return "Vtpdiscoverresultsystemname" }
    return ""
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetSegmentPath() string {
    return "vtpDiscoverResultEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdiscoverresultentry.Managementdomainindex) + "']" + "[vtpDiscoverResultIndex='" + fmt.Sprintf("%v", vtpdiscoverresultentry.Vtpdiscoverresultindex) + "']"
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpdiscoverresultentry.Managementdomainindex
    leafs["vtpDiscoverResultIndex"] = vtpdiscoverresultentry.Vtpdiscoverresultindex
    leafs["vtpDiscoverResultDatabaseName"] = vtpdiscoverresultentry.Vtpdiscoverresultdatabasename
    leafs["vtpDiscoverResultConflicting"] = vtpdiscoverresultentry.Vtpdiscoverresultconflicting
    leafs["vtpDiscoverResultDeviceId"] = vtpdiscoverresultentry.Vtpdiscoverresultdeviceid
    leafs["vtpDiscoverResultPrimaryServer"] = vtpdiscoverresultentry.Vtpdiscoverresultprimaryserver
    leafs["vtpDiscoverResultRevNumber"] = vtpdiscoverresultentry.Vtpdiscoverresultrevnumber
    leafs["vtpDiscoverResultSystemName"] = vtpdiscoverresultentry.Vtpdiscoverresultsystemname
    return leafs
}

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetYangName() string { return "vtpDiscoverResultEntry" }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) SetParent(parent types.Entity) { vtpdiscoverresultentry.parent = parent }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetParent() types.Entity { return vtpdiscoverresultentry.parent }

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetParentYangName() string { return "vtpDiscoverResultTable" }

// CISCOVTPMIB_Vtpdatabasetable
// This table contains information of the VTP
// databases. It is not instantiated when
// managementDomainVersionInUse is version1(1), 
// version2(3) or none(3).
type CISCOVTPMIB_Vtpdatabasetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about the status of the VTP database in the domain.  Each VTP
    // database type known to the local device type has an entry in this table. An
    // entry is also created for unknown database which is notified through VTP
    // advertisements from other VTP servers. The type is slice of
    // CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry.
    Vtpdatabaseentry []CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetFilter() yfilter.YFilter { return vtpdatabasetable.YFilter }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) SetFilter(yf yfilter.YFilter) { vtpdatabasetable.YFilter = yf }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetGoName(yname string) string {
    if yname == "vtpDatabaseEntry" { return "Vtpdatabaseentry" }
    return ""
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetSegmentPath() string {
    return "vtpDatabaseTable"
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpDatabaseEntry" {
        for _, c := range vtpdatabasetable.Vtpdatabaseentry {
            if vtpdatabasetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry{}
        vtpdatabasetable.Vtpdatabaseentry = append(vtpdatabasetable.Vtpdatabaseentry, child)
        return &vtpdatabasetable.Vtpdatabaseentry[len(vtpdatabasetable.Vtpdatabaseentry)-1]
    }
    return nil
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpdatabasetable.Vtpdatabaseentry {
        children[vtpdatabasetable.Vtpdatabaseentry[i].GetSegmentPath()] = &vtpdatabasetable.Vtpdatabaseentry[i]
    }
    return children
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetYangName() string { return "vtpDatabaseTable" }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) SetParent(parent types.Entity) { vtpdatabasetable.parent = parent }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetParent() types.Entity { return vtpdatabasetable.parent }

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry
// Information about the status of the VTP database
// in the domain.  Each VTP database type known to the
// local device type has an entry in this table.
// An entry is also created for unknown database which is
// notified through VTP advertisements from other VTP
// servers.
type CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // This attribute is a key. A value assigned by the system which uniquely
    // identifies a VTP database in the local system. The type is interface{} with
    // range: 0..4294967295.
    Vtpdatabaseindex interface{}

    // The name of the database. The type is string with length: 0..64.
    Vtpdatabasename interface{}

    // The local VTP mode for a particular database type in this administrative
    // domain.  - 'client' indicates that the local system is acting   as a VTP
    // client of the database type.  - 'server' indicates that the local system is
    // acting   as a VTP server of the database type.  - 'transparent' indicates
    // that the local system does   not generate or listen to VTP messages of this
    // database type, but forwards   messages. This mode can also be set by the
    // device   itself when the size of database is too large for it   to hold in
    // DRAM.  - 'off' indicates that the local system does not   generate, listen
    // to or forward any VTP messages   of this database type.  The default mode
    // is 'client' for the database type  known to the local device and
    // 'transparent' for the unknown database type. The type is
    // Vtpdatabaselocalmode.
    Vtpdatabaselocalmode interface{}

    // The current configuration revision number as known by the local device for
    // this VTP 3 database type in the management domain.  This value is updated
    // (if necessary) whenever a  VTP advertisement for the database type is
    // received  or generated. When the database type is unknown to the  local
    // device or no VTP advertisement for the database type is received or
    // generated, its value is 0. The type is interface{} with range:
    // 0..4294967295.
    Vtpdatabaserevnumber interface{}

    // There are two kinds of VTP version 3 servers for a certain database type -
    // the primary server and the secondary server. When a local device is
    // configured as a server for a certain database type, it becomes secondary
    // server by default. Primary server is an operational role under which a
    // server can initiate or change the VTP configuration of the database type. 
    // A true(1) value indicates that the local device is the  primary server of
    // the database type in the management domain. A false(2) value indicates that
    // the local device is not the primary server, or the database type is unknown
    // to the local device. The type is bool.
    Vtpdatabaseprimaryserver interface{}

    // The unique identifier of the primary server in the management domain for
    // the database type.   If no primary server is discovered for the database
    // type, the object has a value of zero length string. The type is string with
    // length: 0..64.
    Vtpdatabaseprimaryserverid interface{}

    // There are two kinds of VTP version 3 servers for a certain database type -
    // the primary server and the secondary server. When a local device is
    // configured as a server for a certain database type, it becomes secondary
    // server by default. Primary server is an operational role under which a
    // server can initiate or change the VTP configuration of the database type. 
    // Setting this object to a true(1) value will advertise the configuration of
    // this database type to the whole domain.  In order to successfully setting
    // this object to true(1), the value of vtpDatabaseLocalMode must be
    // server(2). Besides that, when the VTP password is hidden from the
    // configuration file, the password (vtpDatabaseTakeOverPassword) which 
    // matches  the secret key (vtpAuthSecretKey) must be provided in the same
    // data packet.   When read, the object always returns false(2). The type is
    // bool.
    Vtpdatabasetakeoverprimary interface{}

    // When read, this object always returns the value of a zero-length octet
    // string.  In the case that the VTP password is hidden from the 
    // configuration and the local device intends to take over the whole domain,
    // this object must be  set to the matching password with the secret key 
    // (vtpAuthSecretKey) in the same data packet as which  the
    // vtpDatabaseTakeOverPrimary is in. In all the  other situations, setting a
    // valid value to this object  has no impact on the system. The type is string
    // with length: 0..64.
    Vtpdatabasetakeoverpassword interface{}
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetFilter() yfilter.YFilter { return vtpdatabaseentry.YFilter }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) SetFilter(yf yfilter.YFilter) { vtpdatabaseentry.YFilter = yf }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpDatabaseIndex" { return "Vtpdatabaseindex" }
    if yname == "vtpDatabaseName" { return "Vtpdatabasename" }
    if yname == "vtpDatabaseLocalMode" { return "Vtpdatabaselocalmode" }
    if yname == "vtpDatabaseRevNumber" { return "Vtpdatabaserevnumber" }
    if yname == "vtpDatabasePrimaryServer" { return "Vtpdatabaseprimaryserver" }
    if yname == "vtpDatabasePrimaryServerId" { return "Vtpdatabaseprimaryserverid" }
    if yname == "vtpDatabaseTakeOverPrimary" { return "Vtpdatabasetakeoverprimary" }
    if yname == "vtpDatabaseTakeOverPassword" { return "Vtpdatabasetakeoverpassword" }
    return ""
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetSegmentPath() string {
    return "vtpDatabaseEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdatabaseentry.Managementdomainindex) + "']" + "[vtpDatabaseIndex='" + fmt.Sprintf("%v", vtpdatabaseentry.Vtpdatabaseindex) + "']"
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpdatabaseentry.Managementdomainindex
    leafs["vtpDatabaseIndex"] = vtpdatabaseentry.Vtpdatabaseindex
    leafs["vtpDatabaseName"] = vtpdatabaseentry.Vtpdatabasename
    leafs["vtpDatabaseLocalMode"] = vtpdatabaseentry.Vtpdatabaselocalmode
    leafs["vtpDatabaseRevNumber"] = vtpdatabaseentry.Vtpdatabaserevnumber
    leafs["vtpDatabasePrimaryServer"] = vtpdatabaseentry.Vtpdatabaseprimaryserver
    leafs["vtpDatabasePrimaryServerId"] = vtpdatabaseentry.Vtpdatabaseprimaryserverid
    leafs["vtpDatabaseTakeOverPrimary"] = vtpdatabaseentry.Vtpdatabasetakeoverprimary
    leafs["vtpDatabaseTakeOverPassword"] = vtpdatabaseentry.Vtpdatabasetakeoverpassword
    return leafs
}

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetYangName() string { return "vtpDatabaseEntry" }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) SetParent(parent types.Entity) { vtpdatabaseentry.parent = parent }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetParent() types.Entity { return vtpdatabaseentry.parent }

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetParentYangName() string { return "vtpDatabaseTable" }

// CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode represents unknown database type.
type CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode string

const (
    CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode_client CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode = "client"

    CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode_server CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode = "server"

    CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode_transparent CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode = "transparent"

    CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode_off CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry_Vtpdatabaselocalmode = "off"
)

// CISCOVTPMIB_Vtpauthenticationtable
// The table contains the authentication information of VTP
// in which the local system participates.
// 
// The security mechanism of VTP relies on a secret key
// that is used to alter the MD5 digest of the packets
// transmitted on the wire.  The secret value is
// created from a password that may be saved in plain text
// in the configuration or hidden from the configuration.
// 
// The device creating or modifying the VTP configuration
// signs it using the MD5 digest generated from the secret
// key before advertising it. Other devices in the domain
// receive this configuration use the same secret key
// to accept it if correctly signed or drop it otherwise.
// 
// The user has the option to hide the password from the
// configuration. Once the password is hidden, the secret
// key generated from the password is shown in the 
// configuration instead, and there is no other way to 
// show the password in plain text again but clearing 
// it or resetting it. 
// 
// In an un-trusted area, the password on a device can 
// be configured without being unveiled. After that,
// it has to be provided again by setting the same 
// value to vtpDatabaseTakeOverPassword if the user 
// wants to take over the whole VTP management domain
// of the database type.
// 
// When managementDomainVersionInUse is version3(4), the 
// authentication mechanism is common to all VTP
// database type.
type CISCOVTPMIB_Vtpauthenticationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about the status of the VTP authentication information in one
    // domain. The type is slice of
    // CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry.
    Vtpauthentry []CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetFilter() yfilter.YFilter { return vtpauthenticationtable.YFilter }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) SetFilter(yf yfilter.YFilter) { vtpauthenticationtable.YFilter = yf }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetGoName(yname string) string {
    if yname == "vtpAuthEntry" { return "Vtpauthentry" }
    return ""
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetSegmentPath() string {
    return "vtpAuthenticationTable"
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vtpAuthEntry" {
        for _, c := range vtpauthenticationtable.Vtpauthentry {
            if vtpauthenticationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry{}
        vtpauthenticationtable.Vtpauthentry = append(vtpauthenticationtable.Vtpauthentry, child)
        return &vtpauthenticationtable.Vtpauthentry[len(vtpauthenticationtable.Vtpauthentry)-1]
    }
    return nil
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vtpauthenticationtable.Vtpauthentry {
        children[vtpauthenticationtable.Vtpauthentry[i].GetSegmentPath()] = &vtpauthenticationtable.Vtpauthentry[i]
    }
    return children
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetBundleName() string { return "cisco_ios_xe" }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetYangName() string { return "vtpAuthenticationTable" }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) SetParent(parent types.Entity) { vtpauthenticationtable.parent = parent }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetParent() types.Entity { return vtpauthenticationtable.parent }

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetParentYangName() string { return "CISCO-VTP-MIB" }

// CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry
// Information about the status of the VTP
// authentication information in one domain.
type CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_Managementdomaintable_Managementdomainentry_Managementdomainindex
    Managementdomainindex interface{}

    // By default, this object has a value of a zero-length character string and
    // is considered to be not configured.  The device uses the password to
    // generate the  secret key. It can be stored in the configuration in  plain
    // text or hidden from the configuration. If a VTP  server intends to modify
    // the database's configuration in the domain but the password was hidden from
    // the configuration, the same password (vtpDatabaseTakeOverPassword) as the
    // hidden one has to be provided.  When this object is set alone,
    // vtpAuthPasswordType is set to plaintext(1) automatically by the system.
    // Setting this object to a zero length character string resets the password
    // to its default value and the password is considered as not configured. 
    // This object is not allowed to be set at the same time when 
    // vtpAuthSecretKey is set.  When the vtpAuthPasswordType is hidden(2), this
    // object  will return a zero-length character string when read. The type is
    // string with length: 0..64.
    Vtpauthpassword interface{}

    // By default this object has the value as plaintext(1) and the VTP password
    // is stored in the configuration file in plain text.  Setting this object to
    // hidden(2) will hide the password from the configuration.  Once this object
    // is set to hidden(2), it cannot be set to plaintext(1) alone. However, it
    // may be set to plaintext(1) at the same time the password is set. The type
    // is Vtpauthpasswordtype.
    Vtpauthpasswordtype interface{}

    // The device creating or modifying the VTP configuration signs it using the
    // MD5 digest generated from the secret key before advertising it. Other
    // devices in the domain receiving this configuration use the same secret key
    // to accept it if it was correctly signed or drop it  otherwise.  By default,
    // the object has the value as a zero-length string and this value is read
    // only. It is set  to this value automatically when the password 
    // (vtpAuthPassword) is set to a zero-length octet string.  The secret key can
    // be either generated using the password or configured by the user. Once the
    // secret key is configured by the user, it is stored as a hexadecimal string
    // in the device's configuration and the password is considered to be the
    // secret key's matching password and hidden from the configuration
    // automatically.  This object is not allowed to be set at the same time when
    // vtpAuthPassword is set.  The secret key is overwritten by a newly generated
    // secret key when the password is re-configured. The type is string with
    // length: 0 | 16.
    Vtpauthsecretkey interface{}
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetFilter() yfilter.YFilter { return vtpauthentry.YFilter }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) SetFilter(yf yfilter.YFilter) { vtpauthentry.YFilter = yf }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetGoName(yname string) string {
    if yname == "managementDomainIndex" { return "Managementdomainindex" }
    if yname == "vtpAuthPassword" { return "Vtpauthpassword" }
    if yname == "vtpAuthPasswordType" { return "Vtpauthpasswordtype" }
    if yname == "vtpAuthSecretKey" { return "Vtpauthsecretkey" }
    return ""
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetSegmentPath() string {
    return "vtpAuthEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpauthentry.Managementdomainindex) + "']"
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["managementDomainIndex"] = vtpauthentry.Managementdomainindex
    leafs["vtpAuthPassword"] = vtpauthentry.Vtpauthpassword
    leafs["vtpAuthPasswordType"] = vtpauthentry.Vtpauthpasswordtype
    leafs["vtpAuthSecretKey"] = vtpauthentry.Vtpauthsecretkey
    return leafs
}

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetBundleName() string { return "cisco_ios_xe" }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetYangName() string { return "vtpAuthEntry" }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) SetParent(parent types.Entity) { vtpauthentry.parent = parent }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetParent() types.Entity { return vtpauthentry.parent }

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetParentYangName() string { return "vtpAuthenticationTable" }

// CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype represents password is set.
type CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype string

const (
    CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype_plaintext CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype = "plaintext"

    CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype_hidden CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype = "hidden"
)

