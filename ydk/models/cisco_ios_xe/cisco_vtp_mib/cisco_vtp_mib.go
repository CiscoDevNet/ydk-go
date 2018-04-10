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
    EntityData types.CommonEntityData
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

func (cISCOVTPMIB *CISCOVTPMIB) GetEntityData() *types.CommonEntityData {
    cISCOVTPMIB.EntityData.YFilter = cISCOVTPMIB.YFilter
    cISCOVTPMIB.EntityData.YangName = "CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVTPMIB.EntityData.ParentYangName = "CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.SegmentPath = "CISCO-VTP-MIB:CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVTPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVTPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVTPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVTPMIB.EntityData.Children["vtpStatus"] = types.YChild{"Vtpstatus", &cISCOVTPMIB.Vtpstatus}
    cISCOVTPMIB.EntityData.Children["internalVlanInfo"] = types.YChild{"Internalvlaninfo", &cISCOVTPMIB.Internalvlaninfo}
    cISCOVTPMIB.EntityData.Children["vlanTrunkPorts"] = types.YChild{"Vlantrunkports", &cISCOVTPMIB.Vlantrunkports}
    cISCOVTPMIB.EntityData.Children["vlanStatistics"] = types.YChild{"Vlanstatistics", &cISCOVTPMIB.Vlanstatistics}
    cISCOVTPMIB.EntityData.Children["managementDomainTable"] = types.YChild{"Managementdomaintable", &cISCOVTPMIB.Managementdomaintable}
    cISCOVTPMIB.EntityData.Children["vtpVlanTable"] = types.YChild{"Vtpvlantable", &cISCOVTPMIB.Vtpvlantable}
    cISCOVTPMIB.EntityData.Children["vtpInternalVlanTable"] = types.YChild{"Vtpinternalvlantable", &cISCOVTPMIB.Vtpinternalvlantable}
    cISCOVTPMIB.EntityData.Children["vtpVlanEditTable"] = types.YChild{"Vtpvlanedittable", &cISCOVTPMIB.Vtpvlanedittable}
    cISCOVTPMIB.EntityData.Children["vtpVlanLocalShutdownTable"] = types.YChild{"Vtpvlanlocalshutdowntable", &cISCOVTPMIB.Vtpvlanlocalshutdowntable}
    cISCOVTPMIB.EntityData.Children["vlanTrunkPortTable"] = types.YChild{"Vlantrunkporttable", &cISCOVTPMIB.Vlantrunkporttable}
    cISCOVTPMIB.EntityData.Children["vtpDiscoverTable"] = types.YChild{"Vtpdiscovertable", &cISCOVTPMIB.Vtpdiscovertable}
    cISCOVTPMIB.EntityData.Children["vtpDiscoverResultTable"] = types.YChild{"Vtpdiscoverresulttable", &cISCOVTPMIB.Vtpdiscoverresulttable}
    cISCOVTPMIB.EntityData.Children["vtpDatabaseTable"] = types.YChild{"Vtpdatabasetable", &cISCOVTPMIB.Vtpdatabasetable}
    cISCOVTPMIB.EntityData.Children["vtpAuthenticationTable"] = types.YChild{"Vtpauthenticationtable", &cISCOVTPMIB.Vtpauthenticationtable}
    cISCOVTPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVTPMIB.EntityData)
}

// CISCOVTPMIB_Vtpstatus
type CISCOVTPMIB_Vtpstatus struct {
    EntityData types.CommonEntityData
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

func (vtpstatus *CISCOVTPMIB_Vtpstatus) GetEntityData() *types.CommonEntityData {
    vtpstatus.EntityData.YFilter = vtpstatus.YFilter
    vtpstatus.EntityData.YangName = "vtpStatus"
    vtpstatus.EntityData.BundleName = "cisco_ios_xe"
    vtpstatus.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpstatus.EntityData.SegmentPath = "vtpStatus"
    vtpstatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpstatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpstatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpstatus.EntityData.Children = make(map[string]types.YChild)
    vtpstatus.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpstatus.EntityData.Leafs["vtpVersion"] = types.YLeaf{"Vtpversion", vtpstatus.Vtpversion}
    vtpstatus.EntityData.Leafs["vtpMaxVlanStorage"] = types.YLeaf{"Vtpmaxvlanstorage", vtpstatus.Vtpmaxvlanstorage}
    vtpstatus.EntityData.Leafs["vtpNotificationsEnabled"] = types.YLeaf{"Vtpnotificationsenabled", vtpstatus.Vtpnotificationsenabled}
    vtpstatus.EntityData.Leafs["vtpVlanCreatedNotifEnabled"] = types.YLeaf{"Vtpvlancreatednotifenabled", vtpstatus.Vtpvlancreatednotifenabled}
    vtpstatus.EntityData.Leafs["vtpVlanDeletedNotifEnabled"] = types.YLeaf{"Vtpvlandeletednotifenabled", vtpstatus.Vtpvlandeletednotifenabled}
    return &(vtpstatus.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The internal VLAN allocation policy.  'ascending'  - internal VLANs are
    // allocated                starting from a lowwer VLAN ID and                
    // upwards. 'descending' - internal VLANs are allocated               
    // starting from a higher VLAN ID and                downwards. The type is
    // Vtpinternalvlanallocpolicy.
    Vtpinternalvlanallocpolicy interface{}
}

func (internalvlaninfo *CISCOVTPMIB_Internalvlaninfo) GetEntityData() *types.CommonEntityData {
    internalvlaninfo.EntityData.YFilter = internalvlaninfo.YFilter
    internalvlaninfo.EntityData.YangName = "internalVlanInfo"
    internalvlaninfo.EntityData.BundleName = "cisco_ios_xe"
    internalvlaninfo.EntityData.ParentYangName = "CISCO-VTP-MIB"
    internalvlaninfo.EntityData.SegmentPath = "internalVlanInfo"
    internalvlaninfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    internalvlaninfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    internalvlaninfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    internalvlaninfo.EntityData.Children = make(map[string]types.YChild)
    internalvlaninfo.EntityData.Leafs = make(map[string]types.YLeaf)
    internalvlaninfo.EntityData.Leafs["vtpInternalVlanAllocPolicy"] = types.YLeaf{"Vtpinternalvlanallocpolicy", internalvlaninfo.Vtpinternalvlanallocpolicy}
    return &(internalvlaninfo.EntityData)
}

// CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy represents                downwards.
type CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy string

const (
    CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy_ascending CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy = "ascending"

    CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy_descending CISCOVTPMIB_Internalvlaninfo_Vtpinternalvlanallocpolicy = "descending"
)

// CISCOVTPMIB_Vlantrunkports
type CISCOVTPMIB_Vlantrunkports struct {
    EntityData types.CommonEntityData
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

func (vlantrunkports *CISCOVTPMIB_Vlantrunkports) GetEntityData() *types.CommonEntityData {
    vlantrunkports.EntityData.YFilter = vlantrunkports.YFilter
    vlantrunkports.EntityData.YangName = "vlanTrunkPorts"
    vlantrunkports.EntityData.BundleName = "cisco_ios_xe"
    vlantrunkports.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlantrunkports.EntityData.SegmentPath = "vlanTrunkPorts"
    vlantrunkports.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlantrunkports.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlantrunkports.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlantrunkports.EntityData.Children = make(map[string]types.YChild)
    vlantrunkports.EntityData.Leafs = make(map[string]types.YLeaf)
    vlantrunkports.EntityData.Leafs["vlanTrunkPortSetSerialNo"] = types.YLeaf{"Vlantrunkportsetserialno", vlantrunkports.Vlantrunkportsetserialno}
    vlantrunkports.EntityData.Leafs["vlanTrunkPortsDot1qTag"] = types.YLeaf{"Vlantrunkportsdot1Qtag", vlantrunkports.Vlantrunkportsdot1Qtag}
    return &(vlantrunkports.EntityData)
}

// CISCOVTPMIB_Vlanstatistics
type CISCOVTPMIB_Vlanstatistics struct {
    EntityData types.CommonEntityData
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

func (vlanstatistics *CISCOVTPMIB_Vlanstatistics) GetEntityData() *types.CommonEntityData {
    vlanstatistics.EntityData.YFilter = vlanstatistics.YFilter
    vlanstatistics.EntityData.YangName = "vlanStatistics"
    vlanstatistics.EntityData.BundleName = "cisco_ios_xe"
    vlanstatistics.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlanstatistics.EntityData.SegmentPath = "vlanStatistics"
    vlanstatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanstatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanstatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanstatistics.EntityData.Children = make(map[string]types.YChild)
    vlanstatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanstatistics.EntityData.Leafs["vlanStatsVlans"] = types.YLeaf{"Vlanstatsvlans", vlanstatistics.Vlanstatsvlans}
    vlanstatistics.EntityData.Leafs["vlanStatsExtendedVlans"] = types.YLeaf{"Vlanstatsextendedvlans", vlanstatistics.Vlanstatsextendedvlans}
    vlanstatistics.EntityData.Leafs["vlanStatsInternalVlans"] = types.YLeaf{"Vlanstatsinternalvlans", vlanstatistics.Vlanstatsinternalvlans}
    vlanstatistics.EntityData.Leafs["vlanStatsFreeVlans"] = types.YLeaf{"Vlanstatsfreevlans", vlanstatistics.Vlanstatsfreevlans}
    return &(vlanstatistics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of one management domain. The type is slice of
    // CISCOVTPMIB_Managementdomaintable_Managementdomainentry.
    Managementdomainentry []CISCOVTPMIB_Managementdomaintable_Managementdomainentry
}

func (managementdomaintable *CISCOVTPMIB_Managementdomaintable) GetEntityData() *types.CommonEntityData {
    managementdomaintable.EntityData.YFilter = managementdomaintable.YFilter
    managementdomaintable.EntityData.YangName = "managementDomainTable"
    managementdomaintable.EntityData.BundleName = "cisco_ios_xe"
    managementdomaintable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    managementdomaintable.EntityData.SegmentPath = "managementDomainTable"
    managementdomaintable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    managementdomaintable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    managementdomaintable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    managementdomaintable.EntityData.Children = make(map[string]types.YChild)
    managementdomaintable.EntityData.Children["managementDomainEntry"] = types.YChild{"Managementdomainentry", nil}
    for i := range managementdomaintable.Managementdomainentry {
        managementdomaintable.EntityData.Children[types.GetSegmentPath(&managementdomaintable.Managementdomainentry[i])] = types.YChild{"Managementdomainentry", &managementdomaintable.Managementdomainentry[i]}
    }
    managementdomaintable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(managementdomaintable.EntityData)
}

// CISCOVTPMIB_Managementdomaintable_Managementdomainentry
// Information about the status of one management domain.
type CISCOVTPMIB_Managementdomaintable_Managementdomainentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (managementdomainentry *CISCOVTPMIB_Managementdomaintable_Managementdomainentry) GetEntityData() *types.CommonEntityData {
    managementdomainentry.EntityData.YFilter = managementdomainentry.YFilter
    managementdomainentry.EntityData.YangName = "managementDomainEntry"
    managementdomainentry.EntityData.BundleName = "cisco_ios_xe"
    managementdomainentry.EntityData.ParentYangName = "managementDomainTable"
    managementdomainentry.EntityData.SegmentPath = "managementDomainEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", managementdomainentry.Managementdomainindex) + "']"
    managementdomainentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    managementdomainentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    managementdomainentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    managementdomainentry.EntityData.Children = make(map[string]types.YChild)
    managementdomainentry.EntityData.Leafs = make(map[string]types.YLeaf)
    managementdomainentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", managementdomainentry.Managementdomainindex}
    managementdomainentry.EntityData.Leafs["managementDomainName"] = types.YLeaf{"Managementdomainname", managementdomainentry.Managementdomainname}
    managementdomainentry.EntityData.Leafs["managementDomainLocalMode"] = types.YLeaf{"Managementdomainlocalmode", managementdomainentry.Managementdomainlocalmode}
    managementdomainentry.EntityData.Leafs["managementDomainConfigRevNumber"] = types.YLeaf{"Managementdomainconfigrevnumber", managementdomainentry.Managementdomainconfigrevnumber}
    managementdomainentry.EntityData.Leafs["managementDomainLastUpdater"] = types.YLeaf{"Managementdomainlastupdater", managementdomainentry.Managementdomainlastupdater}
    managementdomainentry.EntityData.Leafs["managementDomainLastChange"] = types.YLeaf{"Managementdomainlastchange", managementdomainentry.Managementdomainlastchange}
    managementdomainentry.EntityData.Leafs["managementDomainRowStatus"] = types.YLeaf{"Managementdomainrowstatus", managementdomainentry.Managementdomainrowstatus}
    managementdomainentry.EntityData.Leafs["managementDomainTftpServer"] = types.YLeaf{"Managementdomaintftpserver", managementdomainentry.Managementdomaintftpserver}
    managementdomainentry.EntityData.Leafs["managementDomainTftpPathname"] = types.YLeaf{"Managementdomaintftppathname", managementdomainentry.Managementdomaintftppathname}
    managementdomainentry.EntityData.Leafs["managementDomainPruningState"] = types.YLeaf{"Managementdomainpruningstate", managementdomainentry.Managementdomainpruningstate}
    managementdomainentry.EntityData.Leafs["managementDomainVersionInUse"] = types.YLeaf{"Managementdomainversioninuse", managementdomainentry.Managementdomainversioninuse}
    managementdomainentry.EntityData.Leafs["managementDomainPruningStateOper"] = types.YLeaf{"Managementdomainpruningstateoper", managementdomainentry.Managementdomainpruningstateoper}
    managementdomainentry.EntityData.Leafs["managementDomainAdminSrcIf"] = types.YLeaf{"Managementdomainadminsrcif", managementdomainentry.Managementdomainadminsrcif}
    managementdomainentry.EntityData.Leafs["managementDomainSourceOnlyMode"] = types.YLeaf{"Managementdomainsourceonlymode", managementdomainentry.Managementdomainsourceonlymode}
    managementdomainentry.EntityData.Leafs["managementDomainOperSrcIf"] = types.YLeaf{"Managementdomainopersrcif", managementdomainentry.Managementdomainopersrcif}
    managementdomainentry.EntityData.Leafs["managementDomainConfigFile"] = types.YLeaf{"Managementdomainconfigfile", managementdomainentry.Managementdomainconfigfile}
    managementdomainentry.EntityData.Leafs["managementDomainLocalUpdaterType"] = types.YLeaf{"Managementdomainlocalupdatertype", managementdomainentry.Managementdomainlocalupdatertype}
    managementdomainentry.EntityData.Leafs["managementDomainLocalUpdater"] = types.YLeaf{"Managementdomainlocalupdater", managementdomainentry.Managementdomainlocalupdater}
    managementdomainentry.EntityData.Leafs["managementDomainDeviceID"] = types.YLeaf{"Managementdomaindeviceid", managementdomainentry.Managementdomaindeviceid}
    managementdomainentry.EntityData.Leafs["vtpVlanEditOperation"] = types.YLeaf{"Vtpvlaneditoperation", managementdomainentry.Vtpvlaneditoperation}
    managementdomainentry.EntityData.Leafs["vtpVlanApplyStatus"] = types.YLeaf{"Vtpvlanapplystatus", managementdomainentry.Vtpvlanapplystatus}
    managementdomainentry.EntityData.Leafs["vtpVlanEditBufferOwner"] = types.YLeaf{"Vtpvlaneditbufferowner", managementdomainentry.Vtpvlaneditbufferowner}
    managementdomainentry.EntityData.Leafs["vtpVlanEditConfigRevNumber"] = types.YLeaf{"Vtpvlaneditconfigrevnumber", managementdomainentry.Vtpvlaneditconfigrevnumber}
    managementdomainentry.EntityData.Leafs["vtpVlanEditModifiedVlan"] = types.YLeaf{"Vtpvlaneditmodifiedvlan", managementdomainentry.Vtpvlaneditmodifiedvlan}
    managementdomainentry.EntityData.Leafs["vtpInSummaryAdverts"] = types.YLeaf{"Vtpinsummaryadverts", managementdomainentry.Vtpinsummaryadverts}
    managementdomainentry.EntityData.Leafs["vtpInSubsetAdverts"] = types.YLeaf{"Vtpinsubsetadverts", managementdomainentry.Vtpinsubsetadverts}
    managementdomainentry.EntityData.Leafs["vtpInAdvertRequests"] = types.YLeaf{"Vtpinadvertrequests", managementdomainentry.Vtpinadvertrequests}
    managementdomainentry.EntityData.Leafs["vtpOutSummaryAdverts"] = types.YLeaf{"Vtpoutsummaryadverts", managementdomainentry.Vtpoutsummaryadverts}
    managementdomainentry.EntityData.Leafs["vtpOutSubsetAdverts"] = types.YLeaf{"Vtpoutsubsetadverts", managementdomainentry.Vtpoutsubsetadverts}
    managementdomainentry.EntityData.Leafs["vtpOutAdvertRequests"] = types.YLeaf{"Vtpoutadvertrequests", managementdomainentry.Vtpoutadvertrequests}
    managementdomainentry.EntityData.Leafs["vtpConfigRevNumberErrors"] = types.YLeaf{"Vtpconfigrevnumbererrors", managementdomainentry.Vtpconfigrevnumbererrors}
    managementdomainentry.EntityData.Leafs["vtpConfigDigestErrors"] = types.YLeaf{"Vtpconfigdigesterrors", managementdomainentry.Vtpconfigdigesterrors}
    return &(managementdomainentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one current VLAN.  The managementDomainIndex value in the
    // INDEX clause indicates which management domain the VLAN is in. The type is
    // slice of CISCOVTPMIB_Vtpvlantable_Vtpvlanentry.
    Vtpvlanentry []CISCOVTPMIB_Vtpvlantable_Vtpvlanentry
}

func (vtpvlantable *CISCOVTPMIB_Vtpvlantable) GetEntityData() *types.CommonEntityData {
    vtpvlantable.EntityData.YFilter = vtpvlantable.YFilter
    vtpvlantable.EntityData.YangName = "vtpVlanTable"
    vtpvlantable.EntityData.BundleName = "cisco_ios_xe"
    vtpvlantable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpvlantable.EntityData.SegmentPath = "vtpVlanTable"
    vtpvlantable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlantable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlantable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlantable.EntityData.Children = make(map[string]types.YChild)
    vtpvlantable.EntityData.Children["vtpVlanEntry"] = types.YChild{"Vtpvlanentry", nil}
    for i := range vtpvlantable.Vtpvlanentry {
        vtpvlantable.EntityData.Children[types.GetSegmentPath(&vtpvlantable.Vtpvlanentry[i])] = types.YChild{"Vtpvlanentry", &vtpvlantable.Vtpvlanentry[i]}
    }
    vtpvlantable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpvlantable.EntityData)
}

// CISCOVTPMIB_Vtpvlantable_Vtpvlanentry
// Information about one current VLAN.  The
// managementDomainIndex value in the INDEX clause indicates
// which management domain the VLAN is in.
type CISCOVTPMIB_Vtpvlantable_Vtpvlanentry struct {
    EntityData types.CommonEntityData
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

func (vtpvlanentry *CISCOVTPMIB_Vtpvlantable_Vtpvlanentry) GetEntityData() *types.CommonEntityData {
    vtpvlanentry.EntityData.YFilter = vtpvlanentry.YFilter
    vtpvlanentry.EntityData.YangName = "vtpVlanEntry"
    vtpvlanentry.EntityData.BundleName = "cisco_ios_xe"
    vtpvlanentry.EntityData.ParentYangName = "vtpVlanTable"
    vtpvlanentry.EntityData.SegmentPath = "vtpVlanEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlanentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpvlanentry.Vtpvlanindex) + "']"
    vtpvlanentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlanentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlanentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlanentry.EntityData.Children = make(map[string]types.YChild)
    vtpvlanentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpvlanentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpvlanentry.Managementdomainindex}
    vtpvlanentry.EntityData.Leafs["vtpVlanIndex"] = types.YLeaf{"Vtpvlanindex", vtpvlanentry.Vtpvlanindex}
    vtpvlanentry.EntityData.Leafs["vtpVlanState"] = types.YLeaf{"Vtpvlanstate", vtpvlanentry.Vtpvlanstate}
    vtpvlanentry.EntityData.Leafs["vtpVlanType"] = types.YLeaf{"Vtpvlantype", vtpvlanentry.Vtpvlantype}
    vtpvlanentry.EntityData.Leafs["vtpVlanName"] = types.YLeaf{"Vtpvlanname", vtpvlanentry.Vtpvlanname}
    vtpvlanentry.EntityData.Leafs["vtpVlanMtu"] = types.YLeaf{"Vtpvlanmtu", vtpvlanentry.Vtpvlanmtu}
    vtpvlanentry.EntityData.Leafs["vtpVlanDot10Said"] = types.YLeaf{"Vtpvlandot10Said", vtpvlanentry.Vtpvlandot10Said}
    vtpvlanentry.EntityData.Leafs["vtpVlanRingNumber"] = types.YLeaf{"Vtpvlanringnumber", vtpvlanentry.Vtpvlanringnumber}
    vtpvlanentry.EntityData.Leafs["vtpVlanBridgeNumber"] = types.YLeaf{"Vtpvlanbridgenumber", vtpvlanentry.Vtpvlanbridgenumber}
    vtpvlanentry.EntityData.Leafs["vtpVlanStpType"] = types.YLeaf{"Vtpvlanstptype", vtpvlanentry.Vtpvlanstptype}
    vtpvlanentry.EntityData.Leafs["vtpVlanParentVlan"] = types.YLeaf{"Vtpvlanparentvlan", vtpvlanentry.Vtpvlanparentvlan}
    vtpvlanentry.EntityData.Leafs["vtpVlanTranslationalVlan1"] = types.YLeaf{"Vtpvlantranslationalvlan1", vtpvlanentry.Vtpvlantranslationalvlan1}
    vtpvlanentry.EntityData.Leafs["vtpVlanTranslationalVlan2"] = types.YLeaf{"Vtpvlantranslationalvlan2", vtpvlanentry.Vtpvlantranslationalvlan2}
    vtpvlanentry.EntityData.Leafs["vtpVlanBridgeType"] = types.YLeaf{"Vtpvlanbridgetype", vtpvlanentry.Vtpvlanbridgetype}
    vtpvlanentry.EntityData.Leafs["vtpVlanAreHopCount"] = types.YLeaf{"Vtpvlanarehopcount", vtpvlanentry.Vtpvlanarehopcount}
    vtpvlanentry.EntityData.Leafs["vtpVlanSteHopCount"] = types.YLeaf{"Vtpvlanstehopcount", vtpvlanentry.Vtpvlanstehopcount}
    vtpvlanentry.EntityData.Leafs["vtpVlanIsCRFBackup"] = types.YLeaf{"Vtpvlaniscrfbackup", vtpvlanentry.Vtpvlaniscrfbackup}
    vtpvlanentry.EntityData.Leafs["vtpVlanTypeExt"] = types.YLeaf{"Vtpvlantypeext", vtpvlanentry.Vtpvlantypeext}
    vtpvlanentry.EntityData.Leafs["vtpVlanIfIndex"] = types.YLeaf{"Vtpvlanifindex", vtpvlanentry.Vtpvlanifindex}
    vtpvlanentry.EntityData.Leafs["stpxVlanMISTPInstMapInstIndex"] = types.YLeaf{"Stpxvlanmistpinstmapinstindex", vtpvlanentry.Stpxvlanmistpinstmapinstindex}
    return &(vtpvlanentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one current internal VLAN. The type is slice of
    // CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry.
    Vtpinternalvlanentry []CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry
}

func (vtpinternalvlantable *CISCOVTPMIB_Vtpinternalvlantable) GetEntityData() *types.CommonEntityData {
    vtpinternalvlantable.EntityData.YFilter = vtpinternalvlantable.YFilter
    vtpinternalvlantable.EntityData.YangName = "vtpInternalVlanTable"
    vtpinternalvlantable.EntityData.BundleName = "cisco_ios_xe"
    vtpinternalvlantable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpinternalvlantable.EntityData.SegmentPath = "vtpInternalVlanTable"
    vtpinternalvlantable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpinternalvlantable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpinternalvlantable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpinternalvlantable.EntityData.Children = make(map[string]types.YChild)
    vtpinternalvlantable.EntityData.Children["vtpInternalVlanEntry"] = types.YChild{"Vtpinternalvlanentry", nil}
    for i := range vtpinternalvlantable.Vtpinternalvlanentry {
        vtpinternalvlantable.EntityData.Children[types.GetSegmentPath(&vtpinternalvlantable.Vtpinternalvlanentry[i])] = types.YChild{"Vtpinternalvlanentry", &vtpinternalvlantable.Vtpinternalvlanentry[i]}
    }
    vtpinternalvlantable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpinternalvlantable.EntityData)
}

// CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry
// Information about one current internal
// VLAN.
type CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry struct {
    EntityData types.CommonEntityData
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

func (vtpinternalvlanentry *CISCOVTPMIB_Vtpinternalvlantable_Vtpinternalvlanentry) GetEntityData() *types.CommonEntityData {
    vtpinternalvlanentry.EntityData.YFilter = vtpinternalvlanentry.YFilter
    vtpinternalvlanentry.EntityData.YangName = "vtpInternalVlanEntry"
    vtpinternalvlanentry.EntityData.BundleName = "cisco_ios_xe"
    vtpinternalvlanentry.EntityData.ParentYangName = "vtpInternalVlanTable"
    vtpinternalvlanentry.EntityData.SegmentPath = "vtpInternalVlanEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpinternalvlanentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpinternalvlanentry.Vtpvlanindex) + "']"
    vtpinternalvlanentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpinternalvlanentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpinternalvlanentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpinternalvlanentry.EntityData.Children = make(map[string]types.YChild)
    vtpinternalvlanentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpinternalvlanentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpinternalvlanentry.Managementdomainindex}
    vtpinternalvlanentry.EntityData.Leafs["vtpVlanIndex"] = types.YLeaf{"Vtpvlanindex", vtpinternalvlanentry.Vtpvlanindex}
    vtpinternalvlanentry.EntityData.Leafs["vtpInternalVlanOwner"] = types.YLeaf{"Vtpinternalvlanowner", vtpinternalvlanentry.Vtpinternalvlanowner}
    return &(vtpinternalvlanentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one VLAN in the Edit Buffer for a particular management
    // domain. The type is slice of CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry.
    Vtpvlaneditentry []CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry
}

func (vtpvlanedittable *CISCOVTPMIB_Vtpvlanedittable) GetEntityData() *types.CommonEntityData {
    vtpvlanedittable.EntityData.YFilter = vtpvlanedittable.YFilter
    vtpvlanedittable.EntityData.YangName = "vtpVlanEditTable"
    vtpvlanedittable.EntityData.BundleName = "cisco_ios_xe"
    vtpvlanedittable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpvlanedittable.EntityData.SegmentPath = "vtpVlanEditTable"
    vtpvlanedittable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlanedittable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlanedittable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlanedittable.EntityData.Children = make(map[string]types.YChild)
    vtpvlanedittable.EntityData.Children["vtpVlanEditEntry"] = types.YChild{"Vtpvlaneditentry", nil}
    for i := range vtpvlanedittable.Vtpvlaneditentry {
        vtpvlanedittable.EntityData.Children[types.GetSegmentPath(&vtpvlanedittable.Vtpvlaneditentry[i])] = types.YChild{"Vtpvlaneditentry", &vtpvlanedittable.Vtpvlaneditentry[i]}
    }
    vtpvlanedittable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpvlanedittable.EntityData)
}

// CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry
// Information about one VLAN in the Edit Buffer for a
// particular management domain.
type CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry struct {
    EntityData types.CommonEntityData
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

func (vtpvlaneditentry *CISCOVTPMIB_Vtpvlanedittable_Vtpvlaneditentry) GetEntityData() *types.CommonEntityData {
    vtpvlaneditentry.EntityData.YFilter = vtpvlaneditentry.YFilter
    vtpvlaneditentry.EntityData.YangName = "vtpVlanEditEntry"
    vtpvlaneditentry.EntityData.BundleName = "cisco_ios_xe"
    vtpvlaneditentry.EntityData.ParentYangName = "vtpVlanEditTable"
    vtpvlaneditentry.EntityData.SegmentPath = "vtpVlanEditEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlaneditentry.Managementdomainindex) + "']" + "[vtpVlanEditIndex='" + fmt.Sprintf("%v", vtpvlaneditentry.Vtpvlaneditindex) + "']"
    vtpvlaneditentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlaneditentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlaneditentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlaneditentry.EntityData.Children = make(map[string]types.YChild)
    vtpvlaneditentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpvlaneditentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpvlaneditentry.Managementdomainindex}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditIndex"] = types.YLeaf{"Vtpvlaneditindex", vtpvlaneditentry.Vtpvlaneditindex}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditState"] = types.YLeaf{"Vtpvlaneditstate", vtpvlaneditentry.Vtpvlaneditstate}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditType"] = types.YLeaf{"Vtpvlanedittype", vtpvlaneditentry.Vtpvlanedittype}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditName"] = types.YLeaf{"Vtpvlaneditname", vtpvlaneditentry.Vtpvlaneditname}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditMtu"] = types.YLeaf{"Vtpvlaneditmtu", vtpvlaneditentry.Vtpvlaneditmtu}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditDot10Said"] = types.YLeaf{"Vtpvlaneditdot10Said", vtpvlaneditentry.Vtpvlaneditdot10Said}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditRingNumber"] = types.YLeaf{"Vtpvlaneditringnumber", vtpvlaneditentry.Vtpvlaneditringnumber}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditBridgeNumber"] = types.YLeaf{"Vtpvlaneditbridgenumber", vtpvlaneditentry.Vtpvlaneditbridgenumber}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditStpType"] = types.YLeaf{"Vtpvlaneditstptype", vtpvlaneditentry.Vtpvlaneditstptype}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditParentVlan"] = types.YLeaf{"Vtpvlaneditparentvlan", vtpvlaneditentry.Vtpvlaneditparentvlan}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditRowStatus"] = types.YLeaf{"Vtpvlaneditrowstatus", vtpvlaneditentry.Vtpvlaneditrowstatus}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditTranslationalVlan1"] = types.YLeaf{"Vtpvlanedittranslationalvlan1", vtpvlaneditentry.Vtpvlanedittranslationalvlan1}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditTranslationalVlan2"] = types.YLeaf{"Vtpvlanedittranslationalvlan2", vtpvlaneditentry.Vtpvlanedittranslationalvlan2}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditBridgeType"] = types.YLeaf{"Vtpvlaneditbridgetype", vtpvlaneditentry.Vtpvlaneditbridgetype}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditAreHopCount"] = types.YLeaf{"Vtpvlaneditarehopcount", vtpvlaneditentry.Vtpvlaneditarehopcount}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditSteHopCount"] = types.YLeaf{"Vtpvlaneditstehopcount", vtpvlaneditentry.Vtpvlaneditstehopcount}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditIsCRFBackup"] = types.YLeaf{"Vtpvlaneditiscrfbackup", vtpvlaneditentry.Vtpvlaneditiscrfbackup}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditTypeExt"] = types.YLeaf{"Vtpvlanedittypeext", vtpvlaneditentry.Vtpvlanedittypeext}
    vtpvlaneditentry.EntityData.Leafs["vtpVlanEditTypeExt2"] = types.YLeaf{"Vtpvlanedittypeext2", vtpvlaneditentry.Vtpvlanedittypeext2}
    vtpvlaneditentry.EntityData.Leafs["stpxVlanMISTPInstMapEditInstIndex"] = types.YLeaf{"Stpxvlanmistpinstmapeditinstindex", vtpvlaneditentry.Stpxvlanmistpinstmapeditinstindex}
    return &(vtpvlaneditentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing VLAN local shutdown information for a particular VLAN
    // in the management domain.  An entry is created if a VLAN which supports
    // local shutdown has been created.  An entry is deleted if a VLAN which
    // supports local shutdown has been removed. The type is slice of
    // CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry.
    Vtpvlanlocalshutdownentry []CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry
}

func (vtpvlanlocalshutdowntable *CISCOVTPMIB_Vtpvlanlocalshutdowntable) GetEntityData() *types.CommonEntityData {
    vtpvlanlocalshutdowntable.EntityData.YFilter = vtpvlanlocalshutdowntable.YFilter
    vtpvlanlocalshutdowntable.EntityData.YangName = "vtpVlanLocalShutdownTable"
    vtpvlanlocalshutdowntable.EntityData.BundleName = "cisco_ios_xe"
    vtpvlanlocalshutdowntable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpvlanlocalshutdowntable.EntityData.SegmentPath = "vtpVlanLocalShutdownTable"
    vtpvlanlocalshutdowntable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlanlocalshutdowntable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlanlocalshutdowntable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlanlocalshutdowntable.EntityData.Children = make(map[string]types.YChild)
    vtpvlanlocalshutdowntable.EntityData.Children["vtpVlanLocalShutdownEntry"] = types.YChild{"Vtpvlanlocalshutdownentry", nil}
    for i := range vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry {
        vtpvlanlocalshutdowntable.EntityData.Children[types.GetSegmentPath(&vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry[i])] = types.YChild{"Vtpvlanlocalshutdownentry", &vtpvlanlocalshutdowntable.Vtpvlanlocalshutdownentry[i]}
    }
    vtpvlanlocalshutdowntable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpvlanlocalshutdowntable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (vtpvlanlocalshutdownentry *CISCOVTPMIB_Vtpvlanlocalshutdowntable_Vtpvlanlocalshutdownentry) GetEntityData() *types.CommonEntityData {
    vtpvlanlocalshutdownentry.EntityData.YFilter = vtpvlanlocalshutdownentry.YFilter
    vtpvlanlocalshutdownentry.EntityData.YangName = "vtpVlanLocalShutdownEntry"
    vtpvlanlocalshutdownentry.EntityData.BundleName = "cisco_ios_xe"
    vtpvlanlocalshutdownentry.EntityData.ParentYangName = "vtpVlanLocalShutdownTable"
    vtpvlanlocalshutdownentry.EntityData.SegmentPath = "vtpVlanLocalShutdownEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpvlanlocalshutdownentry.Managementdomainindex) + "']" + "[vtpVlanIndex='" + fmt.Sprintf("%v", vtpvlanlocalshutdownentry.Vtpvlanindex) + "']"
    vtpvlanlocalshutdownentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpvlanlocalshutdownentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpvlanlocalshutdownentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpvlanlocalshutdownentry.EntityData.Children = make(map[string]types.YChild)
    vtpvlanlocalshutdownentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpvlanlocalshutdownentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpvlanlocalshutdownentry.Managementdomainindex}
    vtpvlanlocalshutdownentry.EntityData.Leafs["vtpVlanIndex"] = types.YLeaf{"Vtpvlanindex", vtpvlanlocalshutdownentry.Vtpvlanindex}
    vtpvlanlocalshutdownentry.EntityData.Leafs["vtpVlanLocalShutdown"] = types.YLeaf{"Vtpvlanlocalshutdown", vtpvlanlocalshutdownentry.Vtpvlanlocalshutdown}
    return &(vtpvlanlocalshutdownentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one trunk port. The type is slice of
    // CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry.
    Vlantrunkportentry []CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry
}

func (vlantrunkporttable *CISCOVTPMIB_Vlantrunkporttable) GetEntityData() *types.CommonEntityData {
    vlantrunkporttable.EntityData.YFilter = vlantrunkporttable.YFilter
    vlantrunkporttable.EntityData.YangName = "vlanTrunkPortTable"
    vlantrunkporttable.EntityData.BundleName = "cisco_ios_xe"
    vlantrunkporttable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlantrunkporttable.EntityData.SegmentPath = "vlanTrunkPortTable"
    vlantrunkporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlantrunkporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlantrunkporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlantrunkporttable.EntityData.Children = make(map[string]types.YChild)
    vlantrunkporttable.EntityData.Children["vlanTrunkPortEntry"] = types.YChild{"Vlantrunkportentry", nil}
    for i := range vlantrunkporttable.Vlantrunkportentry {
        vlantrunkporttable.EntityData.Children[types.GetSegmentPath(&vlantrunkporttable.Vlantrunkportentry[i])] = types.YChild{"Vlantrunkportentry", &vlantrunkporttable.Vlantrunkportentry[i]}
    }
    vlantrunkporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlantrunkporttable.EntityData)
}

// CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry
// Information about one trunk port.
type CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry struct {
    EntityData types.CommonEntityData
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

func (vlantrunkportentry *CISCOVTPMIB_Vlantrunkporttable_Vlantrunkportentry) GetEntityData() *types.CommonEntityData {
    vlantrunkportentry.EntityData.YFilter = vlantrunkportentry.YFilter
    vlantrunkportentry.EntityData.YangName = "vlanTrunkPortEntry"
    vlantrunkportentry.EntityData.BundleName = "cisco_ios_xe"
    vlantrunkportentry.EntityData.ParentYangName = "vlanTrunkPortTable"
    vlantrunkportentry.EntityData.SegmentPath = "vlanTrunkPortEntry" + "[vlanTrunkPortIfIndex='" + fmt.Sprintf("%v", vlantrunkportentry.Vlantrunkportifindex) + "']"
    vlantrunkportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlantrunkportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlantrunkportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlantrunkportentry.EntityData.Children = make(map[string]types.YChild)
    vlantrunkportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortIfIndex"] = types.YLeaf{"Vlantrunkportifindex", vlantrunkportentry.Vlantrunkportifindex}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortManagementDomain"] = types.YLeaf{"Vlantrunkportmanagementdomain", vlantrunkportentry.Vlantrunkportmanagementdomain}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortEncapsulationType"] = types.YLeaf{"Vlantrunkportencapsulationtype", vlantrunkportentry.Vlantrunkportencapsulationtype}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansEnabled"] = types.YLeaf{"Vlantrunkportvlansenabled", vlantrunkportentry.Vlantrunkportvlansenabled}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortNativeVlan"] = types.YLeaf{"Vlantrunkportnativevlan", vlantrunkportentry.Vlantrunkportnativevlan}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortRowStatus"] = types.YLeaf{"Vlantrunkportrowstatus", vlantrunkportentry.Vlantrunkportrowstatus}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortInJoins"] = types.YLeaf{"Vlantrunkportinjoins", vlantrunkportentry.Vlantrunkportinjoins}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortOutJoins"] = types.YLeaf{"Vlantrunkportoutjoins", vlantrunkportentry.Vlantrunkportoutjoins}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortOldAdverts"] = types.YLeaf{"Vlantrunkportoldadverts", vlantrunkportentry.Vlantrunkportoldadverts}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansPruningEligible"] = types.YLeaf{"Vlantrunkportvlanspruningeligible", vlantrunkportentry.Vlantrunkportvlanspruningeligible}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansXmitJoined"] = types.YLeaf{"Vlantrunkportvlansxmitjoined", vlantrunkportentry.Vlantrunkportvlansxmitjoined}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansRcvJoined"] = types.YLeaf{"Vlantrunkportvlansrcvjoined", vlantrunkportentry.Vlantrunkportvlansrcvjoined}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortDynamicState"] = types.YLeaf{"Vlantrunkportdynamicstate", vlantrunkportentry.Vlantrunkportdynamicstate}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortDynamicStatus"] = types.YLeaf{"Vlantrunkportdynamicstatus", vlantrunkportentry.Vlantrunkportdynamicstatus}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVtpEnabled"] = types.YLeaf{"Vlantrunkportvtpenabled", vlantrunkportentry.Vlantrunkportvtpenabled}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortEncapsulationOperType"] = types.YLeaf{"Vlantrunkportencapsulationopertype", vlantrunkportentry.Vlantrunkportencapsulationopertype}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansEnabled2k"] = types.YLeaf{"Vlantrunkportvlansenabled2K", vlantrunkportentry.Vlantrunkportvlansenabled2K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansEnabled3k"] = types.YLeaf{"Vlantrunkportvlansenabled3K", vlantrunkportentry.Vlantrunkportvlansenabled3K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansEnabled4k"] = types.YLeaf{"Vlantrunkportvlansenabled4K", vlantrunkportentry.Vlantrunkportvlansenabled4K}
    vlantrunkportentry.EntityData.Leafs["vtpVlansPruningEligible2k"] = types.YLeaf{"Vtpvlanspruningeligible2K", vlantrunkportentry.Vtpvlanspruningeligible2K}
    vlantrunkportentry.EntityData.Leafs["vtpVlansPruningEligible3k"] = types.YLeaf{"Vtpvlanspruningeligible3K", vlantrunkportentry.Vtpvlanspruningeligible3K}
    vlantrunkportentry.EntityData.Leafs["vtpVlansPruningEligible4k"] = types.YLeaf{"Vtpvlanspruningeligible4K", vlantrunkportentry.Vtpvlanspruningeligible4K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansXmitJoined2k"] = types.YLeaf{"Vlantrunkportvlansxmitjoined2K", vlantrunkportentry.Vlantrunkportvlansxmitjoined2K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansXmitJoined3k"] = types.YLeaf{"Vlantrunkportvlansxmitjoined3K", vlantrunkportentry.Vlantrunkportvlansxmitjoined3K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansXmitJoined4k"] = types.YLeaf{"Vlantrunkportvlansxmitjoined4K", vlantrunkportentry.Vlantrunkportvlansxmitjoined4K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansRcvJoined2k"] = types.YLeaf{"Vlantrunkportvlansrcvjoined2K", vlantrunkportentry.Vlantrunkportvlansrcvjoined2K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansRcvJoined3k"] = types.YLeaf{"Vlantrunkportvlansrcvjoined3K", vlantrunkportentry.Vlantrunkportvlansrcvjoined3K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansRcvJoined4k"] = types.YLeaf{"Vlantrunkportvlansrcvjoined4K", vlantrunkportentry.Vlantrunkportvlansrcvjoined4K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortDot1qTunnel"] = types.YLeaf{"Vlantrunkportdot1Qtunnel", vlantrunkportentry.Vlantrunkportdot1Qtunnel}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansActiveFirst2k"] = types.YLeaf{"Vlantrunkportvlansactivefirst2K", vlantrunkportentry.Vlantrunkportvlansactivefirst2K}
    vlantrunkportentry.EntityData.Leafs["vlanTrunkPortVlansActiveSecond2k"] = types.YLeaf{"Vlantrunkportvlansactivesecond2K", vlantrunkportentry.Vlantrunkportvlansactivesecond2K}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredVlansMap"] = types.YLeaf{"Stpxpreferredvlansmap", vlantrunkportentry.Stpxpreferredvlansmap}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredVlansMap2k"] = types.YLeaf{"Stpxpreferredvlansmap2K", vlantrunkportentry.Stpxpreferredvlansmap2K}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredVlansMap3k"] = types.YLeaf{"Stpxpreferredvlansmap3K", vlantrunkportentry.Stpxpreferredvlansmap3K}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredVlansMap4k"] = types.YLeaf{"Stpxpreferredvlansmap4K", vlantrunkportentry.Stpxpreferredvlansmap4K}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredMISTPInstancesMap"] = types.YLeaf{"Stpxpreferredmistpinstancesmap", vlantrunkportentry.Stpxpreferredmistpinstancesmap}
    vlantrunkportentry.EntityData.Leafs["stpxPreferredMSTInstancesMap"] = types.YLeaf{"Stpxpreferredmstinstancesmap", vlantrunkportentry.Stpxpreferredmstinstancesmap}
    return &(vlantrunkportentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information related to the discovery of the VTP members in one management
    // domain. The type is slice of CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry.
    Vtpdiscoverentry []CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry
}

func (vtpdiscovertable *CISCOVTPMIB_Vtpdiscovertable) GetEntityData() *types.CommonEntityData {
    vtpdiscovertable.EntityData.YFilter = vtpdiscovertable.YFilter
    vtpdiscovertable.EntityData.YangName = "vtpDiscoverTable"
    vtpdiscovertable.EntityData.BundleName = "cisco_ios_xe"
    vtpdiscovertable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpdiscovertable.EntityData.SegmentPath = "vtpDiscoverTable"
    vtpdiscovertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdiscovertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdiscovertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdiscovertable.EntityData.Children = make(map[string]types.YChild)
    vtpdiscovertable.EntityData.Children["vtpDiscoverEntry"] = types.YChild{"Vtpdiscoverentry", nil}
    for i := range vtpdiscovertable.Vtpdiscoverentry {
        vtpdiscovertable.EntityData.Children[types.GetSegmentPath(&vtpdiscovertable.Vtpdiscoverentry[i])] = types.YChild{"Vtpdiscoverentry", &vtpdiscovertable.Vtpdiscoverentry[i]}
    }
    vtpdiscovertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpdiscovertable.EntityData)
}

// CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry
// Information related to the discovery of the
// VTP members in one management domain.
type CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry struct {
    EntityData types.CommonEntityData
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

func (vtpdiscoverentry *CISCOVTPMIB_Vtpdiscovertable_Vtpdiscoverentry) GetEntityData() *types.CommonEntityData {
    vtpdiscoverentry.EntityData.YFilter = vtpdiscoverentry.YFilter
    vtpdiscoverentry.EntityData.YangName = "vtpDiscoverEntry"
    vtpdiscoverentry.EntityData.BundleName = "cisco_ios_xe"
    vtpdiscoverentry.EntityData.ParentYangName = "vtpDiscoverTable"
    vtpdiscoverentry.EntityData.SegmentPath = "vtpDiscoverEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdiscoverentry.Managementdomainindex) + "']"
    vtpdiscoverentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdiscoverentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdiscoverentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdiscoverentry.EntityData.Children = make(map[string]types.YChild)
    vtpdiscoverentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpdiscoverentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpdiscoverentry.Managementdomainindex}
    vtpdiscoverentry.EntityData.Leafs["vtpDiscoverAction"] = types.YLeaf{"Vtpdiscoveraction", vtpdiscoverentry.Vtpdiscoveraction}
    vtpdiscoverentry.EntityData.Leafs["vtpDiscoverStatus"] = types.YLeaf{"Vtpdiscoverstatus", vtpdiscoverentry.Vtpdiscoverstatus}
    vtpdiscoverentry.EntityData.Leafs["vtpLastDiscoverTime"] = types.YLeaf{"Vtplastdiscovertime", vtpdiscoverentry.Vtplastdiscovertime}
    return &(vtpdiscoverentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row is created for each VTP member which is found through
    // successful discovery. The type is slice of
    // CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry.
    Vtpdiscoverresultentry []CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry
}

func (vtpdiscoverresulttable *CISCOVTPMIB_Vtpdiscoverresulttable) GetEntityData() *types.CommonEntityData {
    vtpdiscoverresulttable.EntityData.YFilter = vtpdiscoverresulttable.YFilter
    vtpdiscoverresulttable.EntityData.YangName = "vtpDiscoverResultTable"
    vtpdiscoverresulttable.EntityData.BundleName = "cisco_ios_xe"
    vtpdiscoverresulttable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpdiscoverresulttable.EntityData.SegmentPath = "vtpDiscoverResultTable"
    vtpdiscoverresulttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdiscoverresulttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdiscoverresulttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdiscoverresulttable.EntityData.Children = make(map[string]types.YChild)
    vtpdiscoverresulttable.EntityData.Children["vtpDiscoverResultEntry"] = types.YChild{"Vtpdiscoverresultentry", nil}
    for i := range vtpdiscoverresulttable.Vtpdiscoverresultentry {
        vtpdiscoverresulttable.EntityData.Children[types.GetSegmentPath(&vtpdiscoverresulttable.Vtpdiscoverresultentry[i])] = types.YChild{"Vtpdiscoverresultentry", &vtpdiscoverresulttable.Vtpdiscoverresultentry[i]}
    }
    vtpdiscoverresulttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpdiscoverresulttable.EntityData)
}

// CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry
// A conceptual row is created for each VTP member which
// is found through successful discovery.
type CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry struct {
    EntityData types.CommonEntityData
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

func (vtpdiscoverresultentry *CISCOVTPMIB_Vtpdiscoverresulttable_Vtpdiscoverresultentry) GetEntityData() *types.CommonEntityData {
    vtpdiscoverresultentry.EntityData.YFilter = vtpdiscoverresultentry.YFilter
    vtpdiscoverresultentry.EntityData.YangName = "vtpDiscoverResultEntry"
    vtpdiscoverresultentry.EntityData.BundleName = "cisco_ios_xe"
    vtpdiscoverresultentry.EntityData.ParentYangName = "vtpDiscoverResultTable"
    vtpdiscoverresultentry.EntityData.SegmentPath = "vtpDiscoverResultEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdiscoverresultentry.Managementdomainindex) + "']" + "[vtpDiscoverResultIndex='" + fmt.Sprintf("%v", vtpdiscoverresultentry.Vtpdiscoverresultindex) + "']"
    vtpdiscoverresultentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdiscoverresultentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdiscoverresultentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdiscoverresultentry.EntityData.Children = make(map[string]types.YChild)
    vtpdiscoverresultentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpdiscoverresultentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpdiscoverresultentry.Managementdomainindex}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultIndex"] = types.YLeaf{"Vtpdiscoverresultindex", vtpdiscoverresultentry.Vtpdiscoverresultindex}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultDatabaseName"] = types.YLeaf{"Vtpdiscoverresultdatabasename", vtpdiscoverresultentry.Vtpdiscoverresultdatabasename}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultConflicting"] = types.YLeaf{"Vtpdiscoverresultconflicting", vtpdiscoverresultentry.Vtpdiscoverresultconflicting}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultDeviceId"] = types.YLeaf{"Vtpdiscoverresultdeviceid", vtpdiscoverresultentry.Vtpdiscoverresultdeviceid}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultPrimaryServer"] = types.YLeaf{"Vtpdiscoverresultprimaryserver", vtpdiscoverresultentry.Vtpdiscoverresultprimaryserver}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultRevNumber"] = types.YLeaf{"Vtpdiscoverresultrevnumber", vtpdiscoverresultentry.Vtpdiscoverresultrevnumber}
    vtpdiscoverresultentry.EntityData.Leafs["vtpDiscoverResultSystemName"] = types.YLeaf{"Vtpdiscoverresultsystemname", vtpdiscoverresultentry.Vtpdiscoverresultsystemname}
    return &(vtpdiscoverresultentry.EntityData)
}

// CISCOVTPMIB_Vtpdatabasetable
// This table contains information of the VTP
// databases. It is not instantiated when
// managementDomainVersionInUse is version1(1), 
// version2(3) or none(3).
type CISCOVTPMIB_Vtpdatabasetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of the VTP database in the domain.  Each VTP
    // database type known to the local device type has an entry in this table. An
    // entry is also created for unknown database which is notified through VTP
    // advertisements from other VTP servers. The type is slice of
    // CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry.
    Vtpdatabaseentry []CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry
}

func (vtpdatabasetable *CISCOVTPMIB_Vtpdatabasetable) GetEntityData() *types.CommonEntityData {
    vtpdatabasetable.EntityData.YFilter = vtpdatabasetable.YFilter
    vtpdatabasetable.EntityData.YangName = "vtpDatabaseTable"
    vtpdatabasetable.EntityData.BundleName = "cisco_ios_xe"
    vtpdatabasetable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpdatabasetable.EntityData.SegmentPath = "vtpDatabaseTable"
    vtpdatabasetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdatabasetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdatabasetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdatabasetable.EntityData.Children = make(map[string]types.YChild)
    vtpdatabasetable.EntityData.Children["vtpDatabaseEntry"] = types.YChild{"Vtpdatabaseentry", nil}
    for i := range vtpdatabasetable.Vtpdatabaseentry {
        vtpdatabasetable.EntityData.Children[types.GetSegmentPath(&vtpdatabasetable.Vtpdatabaseentry[i])] = types.YChild{"Vtpdatabaseentry", &vtpdatabasetable.Vtpdatabaseentry[i]}
    }
    vtpdatabasetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpdatabasetable.EntityData)
}

// CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry
// Information about the status of the VTP database
// in the domain.  Each VTP database type known to the
// local device type has an entry in this table.
// An entry is also created for unknown database which is
// notified through VTP advertisements from other VTP
// servers.
type CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry struct {
    EntityData types.CommonEntityData
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

func (vtpdatabaseentry *CISCOVTPMIB_Vtpdatabasetable_Vtpdatabaseentry) GetEntityData() *types.CommonEntityData {
    vtpdatabaseentry.EntityData.YFilter = vtpdatabaseentry.YFilter
    vtpdatabaseentry.EntityData.YangName = "vtpDatabaseEntry"
    vtpdatabaseentry.EntityData.BundleName = "cisco_ios_xe"
    vtpdatabaseentry.EntityData.ParentYangName = "vtpDatabaseTable"
    vtpdatabaseentry.EntityData.SegmentPath = "vtpDatabaseEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpdatabaseentry.Managementdomainindex) + "']" + "[vtpDatabaseIndex='" + fmt.Sprintf("%v", vtpdatabaseentry.Vtpdatabaseindex) + "']"
    vtpdatabaseentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpdatabaseentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpdatabaseentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpdatabaseentry.EntityData.Children = make(map[string]types.YChild)
    vtpdatabaseentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpdatabaseentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpdatabaseentry.Managementdomainindex}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseIndex"] = types.YLeaf{"Vtpdatabaseindex", vtpdatabaseentry.Vtpdatabaseindex}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseName"] = types.YLeaf{"Vtpdatabasename", vtpdatabaseentry.Vtpdatabasename}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseLocalMode"] = types.YLeaf{"Vtpdatabaselocalmode", vtpdatabaseentry.Vtpdatabaselocalmode}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseRevNumber"] = types.YLeaf{"Vtpdatabaserevnumber", vtpdatabaseentry.Vtpdatabaserevnumber}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabasePrimaryServer"] = types.YLeaf{"Vtpdatabaseprimaryserver", vtpdatabaseentry.Vtpdatabaseprimaryserver}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabasePrimaryServerId"] = types.YLeaf{"Vtpdatabaseprimaryserverid", vtpdatabaseentry.Vtpdatabaseprimaryserverid}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseTakeOverPrimary"] = types.YLeaf{"Vtpdatabasetakeoverprimary", vtpdatabaseentry.Vtpdatabasetakeoverprimary}
    vtpdatabaseentry.EntityData.Leafs["vtpDatabaseTakeOverPassword"] = types.YLeaf{"Vtpdatabasetakeoverpassword", vtpdatabaseentry.Vtpdatabasetakeoverpassword}
    return &(vtpdatabaseentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of the VTP authentication information in one
    // domain. The type is slice of
    // CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry.
    Vtpauthentry []CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry
}

func (vtpauthenticationtable *CISCOVTPMIB_Vtpauthenticationtable) GetEntityData() *types.CommonEntityData {
    vtpauthenticationtable.EntityData.YFilter = vtpauthenticationtable.YFilter
    vtpauthenticationtable.EntityData.YangName = "vtpAuthenticationTable"
    vtpauthenticationtable.EntityData.BundleName = "cisco_ios_xe"
    vtpauthenticationtable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpauthenticationtable.EntityData.SegmentPath = "vtpAuthenticationTable"
    vtpauthenticationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpauthenticationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpauthenticationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpauthenticationtable.EntityData.Children = make(map[string]types.YChild)
    vtpauthenticationtable.EntityData.Children["vtpAuthEntry"] = types.YChild{"Vtpauthentry", nil}
    for i := range vtpauthenticationtable.Vtpauthentry {
        vtpauthenticationtable.EntityData.Children[types.GetSegmentPath(&vtpauthenticationtable.Vtpauthentry[i])] = types.YChild{"Vtpauthentry", &vtpauthenticationtable.Vtpauthentry[i]}
    }
    vtpauthenticationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vtpauthenticationtable.EntityData)
}

// CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry
// Information about the status of the VTP
// authentication information in one domain.
type CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry struct {
    EntityData types.CommonEntityData
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

func (vtpauthentry *CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry) GetEntityData() *types.CommonEntityData {
    vtpauthentry.EntityData.YFilter = vtpauthentry.YFilter
    vtpauthentry.EntityData.YangName = "vtpAuthEntry"
    vtpauthentry.EntityData.BundleName = "cisco_ios_xe"
    vtpauthentry.EntityData.ParentYangName = "vtpAuthenticationTable"
    vtpauthentry.EntityData.SegmentPath = "vtpAuthEntry" + "[managementDomainIndex='" + fmt.Sprintf("%v", vtpauthentry.Managementdomainindex) + "']"
    vtpauthentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpauthentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpauthentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpauthentry.EntityData.Children = make(map[string]types.YChild)
    vtpauthentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vtpauthentry.EntityData.Leafs["managementDomainIndex"] = types.YLeaf{"Managementdomainindex", vtpauthentry.Managementdomainindex}
    vtpauthentry.EntityData.Leafs["vtpAuthPassword"] = types.YLeaf{"Vtpauthpassword", vtpauthentry.Vtpauthpassword}
    vtpauthentry.EntityData.Leafs["vtpAuthPasswordType"] = types.YLeaf{"Vtpauthpasswordtype", vtpauthentry.Vtpauthpasswordtype}
    vtpauthentry.EntityData.Leafs["vtpAuthSecretKey"] = types.YLeaf{"Vtpauthsecretkey", vtpauthentry.Vtpauthsecretkey}
    return &(vtpauthentry.EntityData)
}

// CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype represents password is set.
type CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype string

const (
    CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype_plaintext CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype = "plaintext"

    CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype_hidden CISCOVTPMIB_Vtpauthenticationtable_Vtpauthentry_Vtpauthpasswordtype = "hidden"
)

