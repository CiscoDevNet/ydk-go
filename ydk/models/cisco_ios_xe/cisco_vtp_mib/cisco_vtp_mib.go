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

    
    VtpStatus CISCOVTPMIB_VtpStatus

    
    InternalVlanInfo CISCOVTPMIB_InternalVlanInfo

    
    VlanTrunkPorts CISCOVTPMIB_VlanTrunkPorts

    
    VlanStatistics CISCOVTPMIB_VlanStatistics

    // The table containing information on the management domains in which the
    // local system is participating.  Devices which support only one management
    // domain will support just one row in this table, and will not let it be
    // deleted nor let other rows be created.  Devices which support multiple
    // management domains will allow rows to be created and deleted, but will not
    // allow the last row to be deleted. If the device does  not support VTP, the
    // table is read-only.
    ManagementDomainTable CISCOVTPMIB_ManagementDomainTable

    // This table contains information on the VLANs which currently exist.
    VtpVlanTable CISCOVTPMIB_VtpVlanTable

    // A vtpInternalVlanTable entry contains information on an existing internal
    // VLAN. It is internally created by the device for a specific application
    // program  and hence owned by the application.   It cannot be modified or
    // deleted by (local  or network) management.
    VtpInternalVlanTable CISCOVTPMIB_VtpInternalVlanTable

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
    VtpVlanEditTable CISCOVTPMIB_VtpVlanEditTable

    // Ths table contains the VLAN local shutdown information within management
    // domain.
    VtpVlanLocalShutdownTable CISCOVTPMIB_VtpVlanLocalShutdownTable

    // The table containing information on the local system's VLAN trunk ports.
    VlanTrunkPortTable CISCOVTPMIB_VlanTrunkPortTable

    // This table contains information related to the discovery of the VTP members
    // in the designated management domain. This table is not instantiated when 
    // managementDomainVersionInUse is version1(1), version2(3) or none(3).
    VtpDiscoverTable CISCOVTPMIB_VtpDiscoverTable

    // The table containing information of discovered VTP members in the
    // management domain in which the local system is participating. This table is
    // not instantiated when  managementDomainVersionInUse is version1(1),
    // version2(2) or none(3).
    VtpDiscoverResultTable CISCOVTPMIB_VtpDiscoverResultTable

    // This table contains information of the VTP databases. It is not
    // instantiated when managementDomainVersionInUse is version1(1),  version2(3)
    // or none(3).
    VtpDatabaseTable CISCOVTPMIB_VtpDatabaseTable

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
    VtpAuthenticationTable CISCOVTPMIB_VtpAuthenticationTable
}

func (cISCOVTPMIB *CISCOVTPMIB) GetEntityData() *types.CommonEntityData {
    cISCOVTPMIB.EntityData.YFilter = cISCOVTPMIB.YFilter
    cISCOVTPMIB.EntityData.YangName = "CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVTPMIB.EntityData.ParentYangName = "CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.SegmentPath = "CISCO-VTP-MIB:CISCO-VTP-MIB"
    cISCOVTPMIB.EntityData.AbsolutePath = cISCOVTPMIB.EntityData.SegmentPath
    cISCOVTPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVTPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVTPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVTPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVTPMIB.EntityData.Children.Append("vtpStatus", types.YChild{"VtpStatus", &cISCOVTPMIB.VtpStatus})
    cISCOVTPMIB.EntityData.Children.Append("internalVlanInfo", types.YChild{"InternalVlanInfo", &cISCOVTPMIB.InternalVlanInfo})
    cISCOVTPMIB.EntityData.Children.Append("vlanTrunkPorts", types.YChild{"VlanTrunkPorts", &cISCOVTPMIB.VlanTrunkPorts})
    cISCOVTPMIB.EntityData.Children.Append("vlanStatistics", types.YChild{"VlanStatistics", &cISCOVTPMIB.VlanStatistics})
    cISCOVTPMIB.EntityData.Children.Append("managementDomainTable", types.YChild{"ManagementDomainTable", &cISCOVTPMIB.ManagementDomainTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpVlanTable", types.YChild{"VtpVlanTable", &cISCOVTPMIB.VtpVlanTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpInternalVlanTable", types.YChild{"VtpInternalVlanTable", &cISCOVTPMIB.VtpInternalVlanTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpVlanEditTable", types.YChild{"VtpVlanEditTable", &cISCOVTPMIB.VtpVlanEditTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpVlanLocalShutdownTable", types.YChild{"VtpVlanLocalShutdownTable", &cISCOVTPMIB.VtpVlanLocalShutdownTable})
    cISCOVTPMIB.EntityData.Children.Append("vlanTrunkPortTable", types.YChild{"VlanTrunkPortTable", &cISCOVTPMIB.VlanTrunkPortTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpDiscoverTable", types.YChild{"VtpDiscoverTable", &cISCOVTPMIB.VtpDiscoverTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpDiscoverResultTable", types.YChild{"VtpDiscoverResultTable", &cISCOVTPMIB.VtpDiscoverResultTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpDatabaseTable", types.YChild{"VtpDatabaseTable", &cISCOVTPMIB.VtpDatabaseTable})
    cISCOVTPMIB.EntityData.Children.Append("vtpAuthenticationTable", types.YChild{"VtpAuthenticationTable", &cISCOVTPMIB.VtpAuthenticationTable})
    cISCOVTPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVTPMIB.EntityData.YListKeys = []string {}

    return &(cISCOVTPMIB.EntityData)
}

// CISCOVTPMIB_VtpStatus
type CISCOVTPMIB_VtpStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The version of VTP in use on the local system.  A device will report its
    // version capability and not any particular version in use on the device. If
    // the device does not support vtp, the version is none(3). The type is
    // VtpVersion.
    VtpVersion interface{}

    // An estimate of the maximum number of VLANs about which the local system can
    // recover complete VTP information after a reboot.  If the number of defined
    // VLANs is greater than this value, then the system can not act as a VTP
    // Server. For a device which has no means to calculate the estimated number,
    // this value is -1. The type is interface{} with range: -1..1023.
    VtpMaxVlanStorage interface{}

    // An indication of whether the notifications/traps defined by the
    // vtpConfigNotificationsGroup, vtpConfigNotificationsGroup2, and
    // vtpConfigNotificationsGroup8 are enabled. The type is bool.
    VtpNotificationsEnabled interface{}

    // An indication of whether the notification should be generated when a VLAN
    // is created.   If the value of this object is 'true' then the vtpVlanCreated
    // notification will be generated.  If the value of this object is 'false'
    // then the vtpVlanCreated notification will not be generated. The type is
    // bool.
    VtpVlanCreatedNotifEnabled interface{}

    // An indication of whether the notification should be generated when a VLAN
    // is deleted.    If the value of this object is 'true' then the
    // vtpVlanDeleted notification will be generated.  If the value of this object
    // is 'false' then the vtpVlanDeleted notification will not be generated. The
    // type is bool.
    VtpVlanDeletedNotifEnabled interface{}
}

func (vtpStatus *CISCOVTPMIB_VtpStatus) GetEntityData() *types.CommonEntityData {
    vtpStatus.EntityData.YFilter = vtpStatus.YFilter
    vtpStatus.EntityData.YangName = "vtpStatus"
    vtpStatus.EntityData.BundleName = "cisco_ios_xe"
    vtpStatus.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpStatus.EntityData.SegmentPath = "vtpStatus"
    vtpStatus.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpStatus.EntityData.SegmentPath
    vtpStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpStatus.EntityData.Children = types.NewOrderedMap()
    vtpStatus.EntityData.Leafs = types.NewOrderedMap()
    vtpStatus.EntityData.Leafs.Append("vtpVersion", types.YLeaf{"VtpVersion", vtpStatus.VtpVersion})
    vtpStatus.EntityData.Leafs.Append("vtpMaxVlanStorage", types.YLeaf{"VtpMaxVlanStorage", vtpStatus.VtpMaxVlanStorage})
    vtpStatus.EntityData.Leafs.Append("vtpNotificationsEnabled", types.YLeaf{"VtpNotificationsEnabled", vtpStatus.VtpNotificationsEnabled})
    vtpStatus.EntityData.Leafs.Append("vtpVlanCreatedNotifEnabled", types.YLeaf{"VtpVlanCreatedNotifEnabled", vtpStatus.VtpVlanCreatedNotifEnabled})
    vtpStatus.EntityData.Leafs.Append("vtpVlanDeletedNotifEnabled", types.YLeaf{"VtpVlanDeletedNotifEnabled", vtpStatus.VtpVlanDeletedNotifEnabled})

    vtpStatus.EntityData.YListKeys = []string {}

    return &(vtpStatus.EntityData)
}

// CISCOVTPMIB_VtpStatus_VtpVersion represents vtp, the version is none(3).
type CISCOVTPMIB_VtpStatus_VtpVersion string

const (
    CISCOVTPMIB_VtpStatus_VtpVersion_one CISCOVTPMIB_VtpStatus_VtpVersion = "one"

    CISCOVTPMIB_VtpStatus_VtpVersion_two CISCOVTPMIB_VtpStatus_VtpVersion = "two"

    CISCOVTPMIB_VtpStatus_VtpVersion_none CISCOVTPMIB_VtpStatus_VtpVersion = "none"

    CISCOVTPMIB_VtpStatus_VtpVersion_three CISCOVTPMIB_VtpStatus_VtpVersion = "three"
)

// CISCOVTPMIB_InternalVlanInfo
type CISCOVTPMIB_InternalVlanInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The internal VLAN allocation policy.  'ascending'  - internal VLANs are
    // allocated                starting from a lowwer VLAN ID and                
    // upwards. 'descending' - internal VLANs are allocated               
    // starting from a higher VLAN ID and                downwards. The type is
    // VtpInternalVlanAllocPolicy.
    VtpInternalVlanAllocPolicy interface{}
}

func (internalVlanInfo *CISCOVTPMIB_InternalVlanInfo) GetEntityData() *types.CommonEntityData {
    internalVlanInfo.EntityData.YFilter = internalVlanInfo.YFilter
    internalVlanInfo.EntityData.YangName = "internalVlanInfo"
    internalVlanInfo.EntityData.BundleName = "cisco_ios_xe"
    internalVlanInfo.EntityData.ParentYangName = "CISCO-VTP-MIB"
    internalVlanInfo.EntityData.SegmentPath = "internalVlanInfo"
    internalVlanInfo.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + internalVlanInfo.EntityData.SegmentPath
    internalVlanInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    internalVlanInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    internalVlanInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    internalVlanInfo.EntityData.Children = types.NewOrderedMap()
    internalVlanInfo.EntityData.Leafs = types.NewOrderedMap()
    internalVlanInfo.EntityData.Leafs.Append("vtpInternalVlanAllocPolicy", types.YLeaf{"VtpInternalVlanAllocPolicy", internalVlanInfo.VtpInternalVlanAllocPolicy})

    internalVlanInfo.EntityData.YListKeys = []string {}

    return &(internalVlanInfo.EntityData)
}

// CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy represents                downwards.
type CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy string

const (
    CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy_ascending CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy = "ascending"

    CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy_descending CISCOVTPMIB_InternalVlanInfo_VtpInternalVlanAllocPolicy = "descending"
)

// CISCOVTPMIB_VlanTrunkPorts
type CISCOVTPMIB_VlanTrunkPorts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An advisory lock used to allow several cooperating SNMPv2 managers to
    // coordinate their use of the SNMPv2 set operation acting upon any instance
    // of vlanTrunkPortVlansEnabled. The type is interface{} with range:
    // 0..2147483647.
    VlanTrunkPortSetSerialNo interface{}

    // An indication of whether the tagging on all VLANs including native VLAN for
    // all 802.1q trunks is enabled.  If this object has a value of true(1) then
    // all VLANs including native VLAN are tagged.  If the value is false(2) then
    // all VLANs excluding native VLAN are tagged.  This object has been
    // deprecated and is replaced by the object 'cltcDot1qAllTaggedEnabled' in the
    // CISCO-L2-TUNNEL-CONFIG-MIB. The type is bool.
    VlanTrunkPortsDot1qTag interface{}
}

func (vlanTrunkPorts *CISCOVTPMIB_VlanTrunkPorts) GetEntityData() *types.CommonEntityData {
    vlanTrunkPorts.EntityData.YFilter = vlanTrunkPorts.YFilter
    vlanTrunkPorts.EntityData.YangName = "vlanTrunkPorts"
    vlanTrunkPorts.EntityData.BundleName = "cisco_ios_xe"
    vlanTrunkPorts.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlanTrunkPorts.EntityData.SegmentPath = "vlanTrunkPorts"
    vlanTrunkPorts.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vlanTrunkPorts.EntityData.SegmentPath
    vlanTrunkPorts.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanTrunkPorts.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanTrunkPorts.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanTrunkPorts.EntityData.Children = types.NewOrderedMap()
    vlanTrunkPorts.EntityData.Leafs = types.NewOrderedMap()
    vlanTrunkPorts.EntityData.Leafs.Append("vlanTrunkPortSetSerialNo", types.YLeaf{"VlanTrunkPortSetSerialNo", vlanTrunkPorts.VlanTrunkPortSetSerialNo})
    vlanTrunkPorts.EntityData.Leafs.Append("vlanTrunkPortsDot1qTag", types.YLeaf{"VlanTrunkPortsDot1qTag", vlanTrunkPorts.VlanTrunkPortsDot1qTag})

    vlanTrunkPorts.EntityData.YListKeys = []string {}

    return &(vlanTrunkPorts.EntityData)
}

// CISCOVTPMIB_VlanStatistics
type CISCOVTPMIB_VlanStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the number of the existing manageable VLANs with VLAN
    // indices from 1 to 1024 in the system. The type is interface{} with range:
    // 0..4294967295.
    VlanStatsVlans interface{}

    // This object indicates the number of the existing manageable VLANs with VLAN
    // indices greater than 1024 in the system. The type is interface{} with
    // range: 0..4294967295.
    VlanStatsExtendedVlans interface{}

    // This object indicates the number of the internal VLANs existing in the
    // system. The type is interface{} with range: 0..4294967295.
    VlanStatsInternalVlans interface{}

    // This object indicates the number of the free or unused VLANs in the system.
    // The type is interface{} with range: 0..4294967295.
    VlanStatsFreeVlans interface{}
}

func (vlanStatistics *CISCOVTPMIB_VlanStatistics) GetEntityData() *types.CommonEntityData {
    vlanStatistics.EntityData.YFilter = vlanStatistics.YFilter
    vlanStatistics.EntityData.YangName = "vlanStatistics"
    vlanStatistics.EntityData.BundleName = "cisco_ios_xe"
    vlanStatistics.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlanStatistics.EntityData.SegmentPath = "vlanStatistics"
    vlanStatistics.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vlanStatistics.EntityData.SegmentPath
    vlanStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanStatistics.EntityData.Children = types.NewOrderedMap()
    vlanStatistics.EntityData.Leafs = types.NewOrderedMap()
    vlanStatistics.EntityData.Leafs.Append("vlanStatsVlans", types.YLeaf{"VlanStatsVlans", vlanStatistics.VlanStatsVlans})
    vlanStatistics.EntityData.Leafs.Append("vlanStatsExtendedVlans", types.YLeaf{"VlanStatsExtendedVlans", vlanStatistics.VlanStatsExtendedVlans})
    vlanStatistics.EntityData.Leafs.Append("vlanStatsInternalVlans", types.YLeaf{"VlanStatsInternalVlans", vlanStatistics.VlanStatsInternalVlans})
    vlanStatistics.EntityData.Leafs.Append("vlanStatsFreeVlans", types.YLeaf{"VlanStatsFreeVlans", vlanStatistics.VlanStatsFreeVlans})

    vlanStatistics.EntityData.YListKeys = []string {}

    return &(vlanStatistics.EntityData)
}

// CISCOVTPMIB_ManagementDomainTable
// The table containing information on the management domains
// in which the local system is participating.  Devices which
// support only one management domain will support just one row
// in this table, and will not let it be deleted nor let other
// rows be created.  Devices which support multiple management
// domains will allow rows to be created and deleted, but will
// not allow the last row to be deleted. If the device does 
// not support VTP, the table is read-only.
type CISCOVTPMIB_ManagementDomainTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of one management domain. The type is slice of
    // CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry.
    ManagementDomainEntry []*CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry
}

func (managementDomainTable *CISCOVTPMIB_ManagementDomainTable) GetEntityData() *types.CommonEntityData {
    managementDomainTable.EntityData.YFilter = managementDomainTable.YFilter
    managementDomainTable.EntityData.YangName = "managementDomainTable"
    managementDomainTable.EntityData.BundleName = "cisco_ios_xe"
    managementDomainTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    managementDomainTable.EntityData.SegmentPath = "managementDomainTable"
    managementDomainTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + managementDomainTable.EntityData.SegmentPath
    managementDomainTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    managementDomainTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    managementDomainTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    managementDomainTable.EntityData.Children = types.NewOrderedMap()
    managementDomainTable.EntityData.Children.Append("managementDomainEntry", types.YChild{"ManagementDomainEntry", nil})
    for i := range managementDomainTable.ManagementDomainEntry {
        managementDomainTable.EntityData.Children.Append(types.GetSegmentPath(managementDomainTable.ManagementDomainEntry[i]), types.YChild{"ManagementDomainEntry", managementDomainTable.ManagementDomainEntry[i]})
    }
    managementDomainTable.EntityData.Leafs = types.NewOrderedMap()

    managementDomainTable.EntityData.YListKeys = []string {}

    return &(managementDomainTable.EntityData)
}

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry
// Information about the status of one management domain.
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An arbitrary value to uniquely identify the
    // management domain on the local system. The type is interface{} with range:
    // 1..255.
    ManagementDomainIndex interface{}

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
    ManagementDomainName interface{}

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
    // ManagementDomainLocalMode.
    ManagementDomainLocalMode interface{}

    // The current Configuration Revision Number as known by the local device for
    // this management domain when  managementDomainVersionInUse is version1(1) or
    // version2(2).  If managementDomainVersionInUse is version3(4), this  object
    // has the same value with vtpDatabaseRevisionNumber  of VLAN database type. 
    // This value is updated (if necessary) whenever a VTP advertisement is
    // received or generated. When in the 'no management-domain' state, this value
    // is 0. The type is interface{} with range: 0..4294967295.
    ManagementDomainConfigRevNumber interface{}

    // The IP-address (or one of them) of the VTP Server which last updated the
    // Configuration Revision Number, as indicated in the most recently received
    // VTP advertisement for this management domain, when
    // managementDomainVersionInUse is version1(1) or version2(2).   If
    // managementDomainVersionInUse is version3(4), this object has the value of
    // 0.0.0.0.  Before an advertisement has been received, this value is 0.0.0.0.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ManagementDomainLastUpdater interface{}

    // The time at which the Configuration Revision Number was (last) increased to
    // its current value, as indicated in the most recently received VTP
    // advertisement for this management domain when managementDomainVersionInUse
    // is not version3(4) or in the most recently received VTP VLAN database 
    // advertisement for this management domain when  managementDomainVersionInUse
    // is version3(4).  The value 0x0000010100000000 indicates that the device
    // which last increased the Configuration Revision Number had no idea of the
    // date/time, or that no advertisement has been received. The type is string.
    ManagementDomainLastChange interface{}

    // The status of this conceptual row. The type is RowStatus.
    ManagementDomainRowStatus interface{}

    // The IP address of a TFTP Server in/from which VTP VLAN information for this
    // management domain is to be stored/retrieved.  If the information is being
    // locally stored in NVRAM, this object should take the value 0.0.0.0. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ManagementDomainTftpServer interface{}

    // The complete pathname of the file at the TFTP Server identified by the
    // value of managementDomainTftpServer in/from which VTP VLAN information for
    // this management domain is to be stored/retrieved.  If the value of
    // corresponding instance of managementDomainTftpServer is 0.0.0.0, the value
    // of this object is ignored. The type is string.
    ManagementDomainTftpPathname interface{}

    // An indication of whether VTP pruning is enabled or disabled in this
    // managament domain.   This object can only be modified, either when the 
    // corresponding instance value of managementDomainVersionInUse  is 'version1'
    // or 'version2' and the corresponding instance  value of
    // managementDomainLocalMode is 'server', or when the  corresponding instance
    // value of managementDomainVersionInUse  is 'version3' and the corresponding
    // instance value of  managementDomainLocalMode is 'server' or 'client'. The
    // type is ManagementDomainPruningState.
    ManagementDomainPruningState interface{}

    // The current version of the VTP that is in use by the designated management
    // domain.   This object can be set to none(3) only when  vtpVersion is
    // none(3). The type is ManagementDomainVersionInUse.
    ManagementDomainVersionInUse interface{}

    // Indicates whether VTP pruning is operationally enabled or disabled in this
    // managament domain. The type is ManagementDomainPruningStateOper.
    ManagementDomainPruningStateOper interface{}

    // The object specifies the interface to be used as the preferred source
    // interface for the VTP IP updater address.  A zero length value indicates
    // that a source interface is not specified. The type is string.
    ManagementDomainAdminSrcIf interface{}

    // The object specifies whether to use only the IP address of
    // managementDomainAdminSrcIf as the VTP IP updater address.   'true'
    // indicates to only use the IP address of         managementDomainAdminSrcIf
    // as the VTP IP         updater address.   'false' indicates to use the IP
    // address of          managementDomainAdminSrcIf as the VTP IP         
    // updater address if managementDomainAdminSrcIf          is configured with
    // an IP address.  Otherwise, the          first available IP address of the
    // system will         be used. The type is bool.
    ManagementDomainSourceOnlyMode interface{}

    // The object indicates the interface used as the preferred source interface
    // for the VTP IP updater address.  A zero length string indicates that a
    // source interface is not available. The type is string.
    ManagementDomainOperSrcIf interface{}

    // The object specifies the file name where VTP configuration is stored in the
    // format of <filename> or <devices>:[<filename>].  <device> can be (but not
    // limited to): flash, bootflash, slot0, slot1, disk0. The type is string.
    ManagementDomainConfigFile interface{}

    // The object indicates the type of the Internet address of the preferred
    // source interface for the VTP IP updater.  The value of this object is
    // 'unknown' if managementDomainVersionInUse is 'version3' or
    // managementDomainLocalMode is not 'server'. The type is InetAddressType.
    ManagementDomainLocalUpdaterType interface{}

    // The object indicates the Internet address of the preferred source interface
    // for the VTP IP updater. The type is string with length: 0..255.
    ManagementDomainLocalUpdater interface{}

    // The object indicates a value that uniquely identifies this device within a
    // VTP Domain.  The value of this object is zero length string if
    // managementDomainVersionInUse is not 'version3'. The type is string.
    ManagementDomainDeviceID interface{}

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
    // 'none' - no operation is performed. The type is VtpVlanEditOperation.
    VtpVlanEditOperation interface{}

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
    // VtpVlanApplyStatus.
    VtpVlanApplyStatus interface{}

    // The management station which is currently using the Edit Buffer for this
    // management domain.  When the Edit Buffer for a management domain is not
    // currently in use, the value of this object is the zero-length string.  Note
    // that it is also the zero-length string if a manager fails to set this
    // object when invoking a copy operation. The type is string with length:
    // 0..127.
    VtpVlanEditBufferOwner interface{}

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
    VtpVlanEditConfigRevNumber interface{}

    // The VLAN-id of the modified VLAN in the Edit Buffer. If the object has the
    // value of zero, any VLAN can  be edited. If the value of the object is not
    // zero, only this VLAN can be edited.  The object's value is reset to zero
    // after a successful 'apply' operation or a 'release' operation.   This
    // object is only supported for devices which allow  only one VLAN editing for
    // each 'apply' operation. For devices which allow multiple VLAN editing for
    // each 'apply' operation, this object is not supported. The type is
    // interface{} with range: 0..4095.
    VtpVlanEditModifiedVlan interface{}

    // The total number of VTP Summary Adverts received for this management
    // domain. The type is interface{} with range: 0..4294967295.
    VtpInSummaryAdverts interface{}

    // The total number of VTP Subset Adverts received for this management domain.
    // The type is interface{} with range: 0..4294967295.
    VtpInSubsetAdverts interface{}

    // The total number of VTP Advert Requests received for this management
    // domain. The type is interface{} with range: 0..4294967295.
    VtpInAdvertRequests interface{}

    // The total number of VTP Summary Adverts sent for this management domain.
    // The type is interface{} with range: 0..4294967295.
    VtpOutSummaryAdverts interface{}

    // The total number of VTP Subset Adverts sent for this management domain. The
    // type is interface{} with range: 0..4294967295.
    VtpOutSubsetAdverts interface{}

    // The total number of VTP Advert Requests sent for this management domain.
    // The type is interface{} with range: 0..4294967295.
    VtpOutAdvertRequests interface{}

    // The number of occurrences of configuration revision number errors for this
    // management domain.  A configuration revision number error occurs when a
    // device receives a VTP advertisement for which:  - the advertisement's
    // Configuration Revision Number is the   same as the current locally-held
    // value, and  - the advertisement's digest value is different from the  
    // current locally-held value. The type is interface{} with range:
    // 0..4294967295.
    VtpConfigRevNumberErrors interface{}

    // The number of occurrences of configuration digest errors for this
    // management domain.  A configuration digest error occurs when a device
    // receives a VTP advertisement for which:  - the advertisement's
    // Configuration Revision Number is   greater than the current locally-held
    // value, and  -  the advertisement's digest value computed by the  receiving
    // device does not match the checksum in the  summary advertisement that was
    // received earlier. This  can happen, for example, if there is a mismatch in
    // VTP  passwords between the VTP devices. The type is interface{} with range:
    // 0..4294967295.
    VtpConfigDigestErrors interface{}
}

func (managementDomainEntry *CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry) GetEntityData() *types.CommonEntityData {
    managementDomainEntry.EntityData.YFilter = managementDomainEntry.YFilter
    managementDomainEntry.EntityData.YangName = "managementDomainEntry"
    managementDomainEntry.EntityData.BundleName = "cisco_ios_xe"
    managementDomainEntry.EntityData.ParentYangName = "managementDomainTable"
    managementDomainEntry.EntityData.SegmentPath = "managementDomainEntry" + types.AddKeyToken(managementDomainEntry.ManagementDomainIndex, "managementDomainIndex")
    managementDomainEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/managementDomainTable/" + managementDomainEntry.EntityData.SegmentPath
    managementDomainEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    managementDomainEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    managementDomainEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    managementDomainEntry.EntityData.Children = types.NewOrderedMap()
    managementDomainEntry.EntityData.Leafs = types.NewOrderedMap()
    managementDomainEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", managementDomainEntry.ManagementDomainIndex})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainName", types.YLeaf{"ManagementDomainName", managementDomainEntry.ManagementDomainName})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainLocalMode", types.YLeaf{"ManagementDomainLocalMode", managementDomainEntry.ManagementDomainLocalMode})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainConfigRevNumber", types.YLeaf{"ManagementDomainConfigRevNumber", managementDomainEntry.ManagementDomainConfigRevNumber})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainLastUpdater", types.YLeaf{"ManagementDomainLastUpdater", managementDomainEntry.ManagementDomainLastUpdater})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainLastChange", types.YLeaf{"ManagementDomainLastChange", managementDomainEntry.ManagementDomainLastChange})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainRowStatus", types.YLeaf{"ManagementDomainRowStatus", managementDomainEntry.ManagementDomainRowStatus})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainTftpServer", types.YLeaf{"ManagementDomainTftpServer", managementDomainEntry.ManagementDomainTftpServer})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainTftpPathname", types.YLeaf{"ManagementDomainTftpPathname", managementDomainEntry.ManagementDomainTftpPathname})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainPruningState", types.YLeaf{"ManagementDomainPruningState", managementDomainEntry.ManagementDomainPruningState})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainVersionInUse", types.YLeaf{"ManagementDomainVersionInUse", managementDomainEntry.ManagementDomainVersionInUse})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainPruningStateOper", types.YLeaf{"ManagementDomainPruningStateOper", managementDomainEntry.ManagementDomainPruningStateOper})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainAdminSrcIf", types.YLeaf{"ManagementDomainAdminSrcIf", managementDomainEntry.ManagementDomainAdminSrcIf})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainSourceOnlyMode", types.YLeaf{"ManagementDomainSourceOnlyMode", managementDomainEntry.ManagementDomainSourceOnlyMode})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainOperSrcIf", types.YLeaf{"ManagementDomainOperSrcIf", managementDomainEntry.ManagementDomainOperSrcIf})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainConfigFile", types.YLeaf{"ManagementDomainConfigFile", managementDomainEntry.ManagementDomainConfigFile})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainLocalUpdaterType", types.YLeaf{"ManagementDomainLocalUpdaterType", managementDomainEntry.ManagementDomainLocalUpdaterType})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainLocalUpdater", types.YLeaf{"ManagementDomainLocalUpdater", managementDomainEntry.ManagementDomainLocalUpdater})
    managementDomainEntry.EntityData.Leafs.Append("managementDomainDeviceID", types.YLeaf{"ManagementDomainDeviceID", managementDomainEntry.ManagementDomainDeviceID})
    managementDomainEntry.EntityData.Leafs.Append("vtpVlanEditOperation", types.YLeaf{"VtpVlanEditOperation", managementDomainEntry.VtpVlanEditOperation})
    managementDomainEntry.EntityData.Leafs.Append("vtpVlanApplyStatus", types.YLeaf{"VtpVlanApplyStatus", managementDomainEntry.VtpVlanApplyStatus})
    managementDomainEntry.EntityData.Leafs.Append("vtpVlanEditBufferOwner", types.YLeaf{"VtpVlanEditBufferOwner", managementDomainEntry.VtpVlanEditBufferOwner})
    managementDomainEntry.EntityData.Leafs.Append("vtpVlanEditConfigRevNumber", types.YLeaf{"VtpVlanEditConfigRevNumber", managementDomainEntry.VtpVlanEditConfigRevNumber})
    managementDomainEntry.EntityData.Leafs.Append("vtpVlanEditModifiedVlan", types.YLeaf{"VtpVlanEditModifiedVlan", managementDomainEntry.VtpVlanEditModifiedVlan})
    managementDomainEntry.EntityData.Leafs.Append("vtpInSummaryAdverts", types.YLeaf{"VtpInSummaryAdverts", managementDomainEntry.VtpInSummaryAdverts})
    managementDomainEntry.EntityData.Leafs.Append("vtpInSubsetAdverts", types.YLeaf{"VtpInSubsetAdverts", managementDomainEntry.VtpInSubsetAdverts})
    managementDomainEntry.EntityData.Leafs.Append("vtpInAdvertRequests", types.YLeaf{"VtpInAdvertRequests", managementDomainEntry.VtpInAdvertRequests})
    managementDomainEntry.EntityData.Leafs.Append("vtpOutSummaryAdverts", types.YLeaf{"VtpOutSummaryAdverts", managementDomainEntry.VtpOutSummaryAdverts})
    managementDomainEntry.EntityData.Leafs.Append("vtpOutSubsetAdverts", types.YLeaf{"VtpOutSubsetAdverts", managementDomainEntry.VtpOutSubsetAdverts})
    managementDomainEntry.EntityData.Leafs.Append("vtpOutAdvertRequests", types.YLeaf{"VtpOutAdvertRequests", managementDomainEntry.VtpOutAdvertRequests})
    managementDomainEntry.EntityData.Leafs.Append("vtpConfigRevNumberErrors", types.YLeaf{"VtpConfigRevNumberErrors", managementDomainEntry.VtpConfigRevNumberErrors})
    managementDomainEntry.EntityData.Leafs.Append("vtpConfigDigestErrors", types.YLeaf{"VtpConfigDigestErrors", managementDomainEntry.VtpConfigDigestErrors})

    managementDomainEntry.EntityData.YListKeys = []string {"ManagementDomainIndex"}

    return &(managementDomainEntry.EntityData)
}

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode represents   generate, listen to or forward any VTP messages.
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode_client CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode = "client"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode_server CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode = "server"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode_transparent CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode = "transparent"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode_off CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainLocalMode = "off"
)

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState represents managementDomainLocalMode is 'server' or 'client'.
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState_enabled CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState = "enabled"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState_disabled CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningState = "disabled"
)

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper represents disabled in this managament domain.
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper_enabled CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper = "enabled"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper_disabled CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainPruningStateOper = "disabled"
)

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse represents vtpVersion is none(3).
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse_version1 CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse = "version1"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse_version2 CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse = "version2"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse_none CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse = "none"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse_version3 CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainVersionInUse = "version3"
)

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus represents           is version3(4).
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_inProgress CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "inProgress"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_succeeded CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "succeeded"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_configNumberError CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "configNumberError"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_inconsistentEdit CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "inconsistentEdit"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_tooBig CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "tooBig"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_localNVStoreFail CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "localNVStoreFail"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_remoteNVStoreFail CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "remoteNVStoreFail"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_editBufferEmpty CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "editBufferEmpty"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_someOtherError CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "someOtherError"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus_notPrimaryServer CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanApplyStatus = "notPrimaryServer"
)

// CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation represents  'none' - no operation is performed.
type CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation string

const (
    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation_none CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation = "none"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation_copy_ CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation = "copy"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation_apply CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation = "apply"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation_release CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation = "release"

    CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation_restartTimer CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_VtpVlanEditOperation = "restartTimer"
)

// CISCOVTPMIB_VtpVlanTable
// This table contains information on the VLANs which
// currently exist.
type CISCOVTPMIB_VtpVlanTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one current VLAN.  The managementDomainIndex value in the
    // INDEX clause indicates which management domain the VLAN is in. The type is
    // slice of CISCOVTPMIB_VtpVlanTable_VtpVlanEntry.
    VtpVlanEntry []*CISCOVTPMIB_VtpVlanTable_VtpVlanEntry
}

func (vtpVlanTable *CISCOVTPMIB_VtpVlanTable) GetEntityData() *types.CommonEntityData {
    vtpVlanTable.EntityData.YFilter = vtpVlanTable.YFilter
    vtpVlanTable.EntityData.YangName = "vtpVlanTable"
    vtpVlanTable.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpVlanTable.EntityData.SegmentPath = "vtpVlanTable"
    vtpVlanTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpVlanTable.EntityData.SegmentPath
    vtpVlanTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanTable.EntityData.Children = types.NewOrderedMap()
    vtpVlanTable.EntityData.Children.Append("vtpVlanEntry", types.YChild{"VtpVlanEntry", nil})
    for i := range vtpVlanTable.VtpVlanEntry {
        vtpVlanTable.EntityData.Children.Append(types.GetSegmentPath(vtpVlanTable.VtpVlanEntry[i]), types.YChild{"VtpVlanEntry", vtpVlanTable.VtpVlanEntry[i]})
    }
    vtpVlanTable.EntityData.Leafs = types.NewOrderedMap()

    vtpVlanTable.EntityData.YListKeys = []string {}

    return &(vtpVlanTable.EntityData)
}

// CISCOVTPMIB_VtpVlanTable_VtpVlanEntry
// Information about one current VLAN.  The
// managementDomainIndex value in the INDEX clause indicates
// which management domain the VLAN is in.
type CISCOVTPMIB_VtpVlanTable_VtpVlanEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. The VLAN-id of this VLAN on ISL or 802.1q trunks.
    // The type is interface{} with range: 0..4095.
    VtpVlanIndex interface{}

    // The state of this VLAN.  The state 'mtuTooBigForDevice' indicates that this
    // device cannot participate in this VLAN because the VLAN's MTU is larger
    // than the device can support.  The state 'mtuTooBigForTrunk' indicates that
    // while this VLAN's MTU is supported by this device, it is too large for one
    // or more of the device's trunk ports. The type is VtpVlanState.
    VtpVlanState interface{}

    // The type of this VLAN. The type is VlanType.
    VtpVlanType interface{}

    // The name of this VLAN.  This name is used as the ELAN-name for an ATM
    // LAN-Emulation segment of this VLAN. The type is string with length: 1..32.
    VtpVlanName interface{}

    // The MTU size on this VLAN, defined as the size of largest MAC-layer
    // (information field portion of the) data frame which can be transmitted on
    // the VLAN. The type is interface{} with range: 1500..18190.
    VtpVlanMtu interface{}

    // The value of the 802.10 SAID field for this VLAN. The type is string with
    // length: 4.
    VtpVlanDot10Said interface{}

    // The ring number of this VLAN.  This object is only instantiated when the
    // value of the corresponding instance of vtpVlanType has a value of 'fddi' or
    // 'tokenRing' and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 0..4095.
    VtpVlanRingNumber interface{}

    // The bridge number of the VTP-capable switches for this VLAN.  This object
    // is only instantiated for VLANs that are involved with emulating token ring
    // segments. The type is interface{} with range: 0..15.
    VtpVlanBridgeNumber interface{}

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
    // corresponds to this object is set to 'auto'. The type is VtpVlanStpType.
    VtpVlanStpType interface{}

    // The parent VLAN for this VLAN.  This object is only instantiated when the
    // value of the corresponding instance of vtpVlanType has a value of 'fddi' or
    // 'tokenRing' and Source Routing is in use on this VLAN.  The parent VLAN
    // must have  a vtpVlanType value of fddiNet(4) or trNet(5),  respectively.
    // The type is interface{} with range: 0..4095.
    VtpVlanParentVlan interface{}

    // A VLAN to which this VLAN is being translational-bridged. If this value and
    // the corresponding instance of vtpVlanTranslationalVlan2 are both zero, then
    // this VLAN is not being translational-bridged. The type is interface{} with
    // range: 0..4095.
    VtpVlanTranslationalVlan1 interface{}

    // Another VLAN, i.e., other than that indicated by vtpVlanTranslationalVlan1,
    // to which this VLAN is being translational-bridged.  If this value and the
    // corresponding instance of vtpVlanTranslationalVlan1 are both zero, then
    // this VLAN is not being translational-bridged. The type is interface{} with
    // range: 0..4095.
    VtpVlanTranslationalVlan2 interface{}

    // The type of the Source Route bridging mode in use on this VLAN.  This
    // object is only instantiated when the value of  the corresponding instance
    // of vtpVlanType has a value of  fddi(2) or tokenRing(3) and Source Routing
    // is in use on this VLAN. The type is VtpVlanBridgeType.
    VtpVlanBridgeType interface{}

    // The maximum number of bridge hops allowed in All Routes Explorer frames on
    // this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    VtpVlanAreHopCount interface{}

    // The maximum number of bridge hops allowed in Spanning Tree Explorer frames
    // on this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    VtpVlanSteHopCount interface{}

    // True if this VLAN is of type trCrf and also is acting as a backup trCrf for
    // the ISL distributed BRF. The type is bool.
    VtpVlanIsCRFBackup interface{}

    // The additional type information of this VLAN. The type is map[string]bool.
    VtpVlanTypeExt interface{}

    // The value of the ifIndex corresponding to this VLAN ID. If the VLAN ID does
    // not have its corresponding interface,  this object has the value of zero.
    // The type is interface{} with range: 0..2147483647.
    VtpVlanIfIndex interface{}

    // The MISTP instance, to which the corresponding vlan is mapped. If this
    // value of this mib object is 0,  the corresponding vlan  is not configured
    // to be mapped to any MISTP instance and all the ports under this VLAN remain
    // in blocking state. The type is interface{} with range: 0..256.
    StpxVlanMISTPInstMapInstIndex interface{}
}

func (vtpVlanEntry *CISCOVTPMIB_VtpVlanTable_VtpVlanEntry) GetEntityData() *types.CommonEntityData {
    vtpVlanEntry.EntityData.YFilter = vtpVlanEntry.YFilter
    vtpVlanEntry.EntityData.YangName = "vtpVlanEntry"
    vtpVlanEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanEntry.EntityData.ParentYangName = "vtpVlanTable"
    vtpVlanEntry.EntityData.SegmentPath = "vtpVlanEntry" + types.AddKeyToken(vtpVlanEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpVlanEntry.VtpVlanIndex, "vtpVlanIndex")
    vtpVlanEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpVlanTable/" + vtpVlanEntry.EntityData.SegmentPath
    vtpVlanEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanEntry.EntityData.Children = types.NewOrderedMap()
    vtpVlanEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpVlanEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpVlanEntry.ManagementDomainIndex})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanIndex", types.YLeaf{"VtpVlanIndex", vtpVlanEntry.VtpVlanIndex})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanState", types.YLeaf{"VtpVlanState", vtpVlanEntry.VtpVlanState})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanType", types.YLeaf{"VtpVlanType", vtpVlanEntry.VtpVlanType})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanName", types.YLeaf{"VtpVlanName", vtpVlanEntry.VtpVlanName})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanMtu", types.YLeaf{"VtpVlanMtu", vtpVlanEntry.VtpVlanMtu})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanDot10Said", types.YLeaf{"VtpVlanDot10Said", vtpVlanEntry.VtpVlanDot10Said})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanRingNumber", types.YLeaf{"VtpVlanRingNumber", vtpVlanEntry.VtpVlanRingNumber})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanBridgeNumber", types.YLeaf{"VtpVlanBridgeNumber", vtpVlanEntry.VtpVlanBridgeNumber})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanStpType", types.YLeaf{"VtpVlanStpType", vtpVlanEntry.VtpVlanStpType})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanParentVlan", types.YLeaf{"VtpVlanParentVlan", vtpVlanEntry.VtpVlanParentVlan})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanTranslationalVlan1", types.YLeaf{"VtpVlanTranslationalVlan1", vtpVlanEntry.VtpVlanTranslationalVlan1})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanTranslationalVlan2", types.YLeaf{"VtpVlanTranslationalVlan2", vtpVlanEntry.VtpVlanTranslationalVlan2})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanBridgeType", types.YLeaf{"VtpVlanBridgeType", vtpVlanEntry.VtpVlanBridgeType})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanAreHopCount", types.YLeaf{"VtpVlanAreHopCount", vtpVlanEntry.VtpVlanAreHopCount})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanSteHopCount", types.YLeaf{"VtpVlanSteHopCount", vtpVlanEntry.VtpVlanSteHopCount})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanIsCRFBackup", types.YLeaf{"VtpVlanIsCRFBackup", vtpVlanEntry.VtpVlanIsCRFBackup})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanTypeExt", types.YLeaf{"VtpVlanTypeExt", vtpVlanEntry.VtpVlanTypeExt})
    vtpVlanEntry.EntityData.Leafs.Append("vtpVlanIfIndex", types.YLeaf{"VtpVlanIfIndex", vtpVlanEntry.VtpVlanIfIndex})
    vtpVlanEntry.EntityData.Leafs.Append("stpxVlanMISTPInstMapInstIndex", types.YLeaf{"StpxVlanMISTPInstMapInstIndex", vtpVlanEntry.StpxVlanMISTPInstMapInstIndex})

    vtpVlanEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpVlanIndex"}

    return &(vtpVlanEntry.EntityData)
}

// CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType represents this VLAN.
type CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType string

const (
    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType_none CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType = "none"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType_srt CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType = "srt"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType_srb CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanBridgeType = "srb"
)

// CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState represents one or more of the device's trunk ports.
type CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState string

const (
    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState_operational CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState = "operational"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState_suspended CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState = "suspended"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState_mtuTooBigForDevice CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState = "mtuTooBigForDevice"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState_mtuTooBigForTrunk CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanState = "mtuTooBigForTrunk"
)

// CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType represents to 'auto'.
type CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType string

const (
    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType_ieee CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType = "ieee"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType_ibm CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType = "ibm"

    CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType_hybrid CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanStpType = "hybrid"
)

// CISCOVTPMIB_VtpInternalVlanTable
// A vtpInternalVlanTable entry contains
// information on an existing internal
// VLAN. It is internally created by the
// device for a specific application program 
// and hence owned by the application.  
// It cannot be modified or deleted by (local 
// or network) management.
type CISCOVTPMIB_VtpInternalVlanTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one current internal VLAN. The type is slice of
    // CISCOVTPMIB_VtpInternalVlanTable_VtpInternalVlanEntry.
    VtpInternalVlanEntry []*CISCOVTPMIB_VtpInternalVlanTable_VtpInternalVlanEntry
}

func (vtpInternalVlanTable *CISCOVTPMIB_VtpInternalVlanTable) GetEntityData() *types.CommonEntityData {
    vtpInternalVlanTable.EntityData.YFilter = vtpInternalVlanTable.YFilter
    vtpInternalVlanTable.EntityData.YangName = "vtpInternalVlanTable"
    vtpInternalVlanTable.EntityData.BundleName = "cisco_ios_xe"
    vtpInternalVlanTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpInternalVlanTable.EntityData.SegmentPath = "vtpInternalVlanTable"
    vtpInternalVlanTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpInternalVlanTable.EntityData.SegmentPath
    vtpInternalVlanTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpInternalVlanTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpInternalVlanTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpInternalVlanTable.EntityData.Children = types.NewOrderedMap()
    vtpInternalVlanTable.EntityData.Children.Append("vtpInternalVlanEntry", types.YChild{"VtpInternalVlanEntry", nil})
    for i := range vtpInternalVlanTable.VtpInternalVlanEntry {
        vtpInternalVlanTable.EntityData.Children.Append(types.GetSegmentPath(vtpInternalVlanTable.VtpInternalVlanEntry[i]), types.YChild{"VtpInternalVlanEntry", vtpInternalVlanTable.VtpInternalVlanEntry[i]})
    }
    vtpInternalVlanTable.EntityData.Leafs = types.NewOrderedMap()

    vtpInternalVlanTable.EntityData.YListKeys = []string {}

    return &(vtpInternalVlanTable.EntityData)
}

// CISCOVTPMIB_VtpInternalVlanTable_VtpInternalVlanEntry
// Information about one current internal
// VLAN.
type CISCOVTPMIB_VtpInternalVlanTable_VtpInternalVlanEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanIndex
    VtpVlanIndex interface{}

    // The program name of the internal VLAN's owner application. This internal
    // VLAN is allocated by the device specifically for this application and no
    // one else could create, modify or delete this  VLAN. The type is string.
    VtpInternalVlanOwner interface{}
}

func (vtpInternalVlanEntry *CISCOVTPMIB_VtpInternalVlanTable_VtpInternalVlanEntry) GetEntityData() *types.CommonEntityData {
    vtpInternalVlanEntry.EntityData.YFilter = vtpInternalVlanEntry.YFilter
    vtpInternalVlanEntry.EntityData.YangName = "vtpInternalVlanEntry"
    vtpInternalVlanEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpInternalVlanEntry.EntityData.ParentYangName = "vtpInternalVlanTable"
    vtpInternalVlanEntry.EntityData.SegmentPath = "vtpInternalVlanEntry" + types.AddKeyToken(vtpInternalVlanEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpInternalVlanEntry.VtpVlanIndex, "vtpVlanIndex")
    vtpInternalVlanEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpInternalVlanTable/" + vtpInternalVlanEntry.EntityData.SegmentPath
    vtpInternalVlanEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpInternalVlanEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpInternalVlanEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpInternalVlanEntry.EntityData.Children = types.NewOrderedMap()
    vtpInternalVlanEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpInternalVlanEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpInternalVlanEntry.ManagementDomainIndex})
    vtpInternalVlanEntry.EntityData.Leafs.Append("vtpVlanIndex", types.YLeaf{"VtpVlanIndex", vtpInternalVlanEntry.VtpVlanIndex})
    vtpInternalVlanEntry.EntityData.Leafs.Append("vtpInternalVlanOwner", types.YLeaf{"VtpInternalVlanOwner", vtpInternalVlanEntry.VtpInternalVlanOwner})

    vtpInternalVlanEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpVlanIndex"}

    return &(vtpInternalVlanEntry.EntityData)
}

// CISCOVTPMIB_VtpVlanEditTable
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
type CISCOVTPMIB_VtpVlanEditTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one VLAN in the Edit Buffer for a particular management
    // domain. The type is slice of CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry.
    VtpVlanEditEntry []*CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry
}

func (vtpVlanEditTable *CISCOVTPMIB_VtpVlanEditTable) GetEntityData() *types.CommonEntityData {
    vtpVlanEditTable.EntityData.YFilter = vtpVlanEditTable.YFilter
    vtpVlanEditTable.EntityData.YangName = "vtpVlanEditTable"
    vtpVlanEditTable.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanEditTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpVlanEditTable.EntityData.SegmentPath = "vtpVlanEditTable"
    vtpVlanEditTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpVlanEditTable.EntityData.SegmentPath
    vtpVlanEditTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanEditTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanEditTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanEditTable.EntityData.Children = types.NewOrderedMap()
    vtpVlanEditTable.EntityData.Children.Append("vtpVlanEditEntry", types.YChild{"VtpVlanEditEntry", nil})
    for i := range vtpVlanEditTable.VtpVlanEditEntry {
        vtpVlanEditTable.EntityData.Children.Append(types.GetSegmentPath(vtpVlanEditTable.VtpVlanEditEntry[i]), types.YChild{"VtpVlanEditEntry", vtpVlanEditTable.VtpVlanEditEntry[i]})
    }
    vtpVlanEditTable.EntityData.Leafs = types.NewOrderedMap()

    vtpVlanEditTable.EntityData.YListKeys = []string {}

    return &(vtpVlanEditTable.EntityData)
}

// CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry
// Information about one VLAN in the Edit Buffer for a
// particular management domain.
type CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. The VLAN-id which this VLAN would have on ISL or
    // 802.1q trunks. The type is interface{} with range: 0..4095.
    VtpVlanEditIndex interface{}

    // The state which this VLAN would have. The type is VtpVlanEditState.
    VtpVlanEditState interface{}

    // The type which this VLAN would have. An implementation may restrict access
    // to this object. The type is VlanType.
    VtpVlanEditType interface{}

    // The name which this VLAN would have.  This name would be used as the
    // ELAN-name for an ATM LAN-Emulation segment of this VLAN.  An implementation
    // may restrict access to this object. The type is string with length: 1..32.
    VtpVlanEditName interface{}

    // The MTU size which this VLAN would have, defined as the size of largest
    // MAC-layer (information field portion of the) data frame which can be
    // transmitted on the VLAN.  An implementation may restrict access to this
    // object. The type is interface{} with range: 1500..18190.
    VtpVlanEditMtu interface{}

    // The value of the 802.10 SAID field which would be used for this VLAN.  An
    // implementation may restrict access to this object. The type is string with
    // length: 4.
    VtpVlanEditDot10Said interface{}

    // The ring number which would be used for this VLAN.  This object is only
    // instantiated when the value of the corresponding instance of
    // vtpVlanEditType has a value of 'fddi' or 'tokenRing' and Source Routing is
    // in use on  this VLAN. The type is interface{} with range: 0..4095.
    VtpVlanEditRingNumber interface{}

    // The bridge number of the VTP-capable switches which would be used for this
    // VLAN.  This object is only instantiated when the value of the corresponding
    // instance of vtpVlanEditType has a value of fddiNet(4) or trNet(5). The type
    // is interface{} with range: 0..15.
    VtpVlanEditBridgeNumber interface{}

    // The type of the Spanning Tree Protocol which would be running on this VLAN.
    // This object is only instantiated when the value of the corresponding
    // instance of vtpVlanEditType has a value of fddiNet(4) or trNet(5).  If
    // 'ieee' is selected, the STP that runs will be IEEE.  If 'ibm' is selected,
    // the STP that runs will be IBM.  If 'auto' is selected, the STP that runs
    // will be dependant on the values of vtpVlanEditBridgeType for all children
    // tokenRing/fddi type VLANs.  This will result in a 'hybrid' STP (see
    // vtpVlanStpType). The type is VtpVlanEditStpType.
    VtpVlanEditStpType interface{}

    // The VLAN index of the VLAN which would be the parent for this VLAN.  This
    // object is only instantiated when the value of the corresponding instance of
    // vtpVlanEditType has a value of 'fddi' or 'tokenRing' and Source Routing is
    // in use on this VLAN.  The parent VLAN must have a vtpVlanEditType  value of
    // fddiNet(4) or trNet(5), respectively. The type is interface{} with range:
    // 0..4095.
    VtpVlanEditParentVlan interface{}

    // The status of this row.  Any and all columnar objects in an existing row
    // can be modified irrespective of the status of the row.  A row is not
    // qualified for activation until instances of at least its vtpVlanEditType,
    // vtpVlanEditName and vtpVlanEditDot10Said columns have appropriate values. 
    // The management station should endeavor to make all rows consistent in the
    // table before 'apply'ing the buffer.  An inconsistent entry in the table
    // will cause the entire buffer to be rejected with the vtpVlanApplyStatus
    // object set to the appropriate error value. The type is RowStatus.
    VtpVlanEditRowStatus interface{}

    // A VLAN to which this VLAN would be translational-bridged. If this value and
    // the corresponding instance of vtpVlanTranslationalVlan2 are both zero, then
    // this VLAN would not be translational-bridged.  An implementation may
    // restrict access to this object. The type is interface{} with range:
    // 0..4095.
    VtpVlanEditTranslationalVlan1 interface{}

    // Another VLAN, i.e., other than that indicated by
    // vtpVlanEditTranslationalVlan1, to which this VLAN would be
    // translational-bridged.  If this value and the corresponding instance of
    // vtpVlanTranslationalVlan1 are both zero, then this VLAN would not be
    // translational-bridged.  An implementation may restrict access to this
    // object. The type is interface{} with range: 0..4095.
    VtpVlanEditTranslationalVlan2 interface{}

    // The type of Source Route bridging mode which would be in use on this VLAN. 
    // This object is only instantiated when  the value of  the corresponding
    // instance of vtpVlanEditType has a value of fddi(2) or tokenRing(3) and
    // Source Routing  is in use on this VLAN. The type is VtpVlanEditBridgeType.
    VtpVlanEditBridgeType interface{}

    // The maximum number of bridge hops allowed in All Routes Explorer frames on
    // this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    VtpVlanEditAreHopCount interface{}

    // The maximum number of bridge hops allowed in Spanning Tree Explorer frames
    // on this VLAN.  This object is only instantiated when the value of the
    // corresponding instance of vtpVlanType has a value of fddi(2) or
    // tokenRing(3) and Source Routing is in use on this VLAN. The type is
    // interface{} with range: 1..13.
    VtpVlanEditSteHopCount interface{}

    // True if this VLAN is of type trCrf and also is acting as a backup trCrf for
    // the ISL distributed BRF.  This object is only instantiated when the value
    // of the corresponding instance of vtpVlanEditType has a value of
    // tokenRing(3). The type is bool.
    VtpVlanEditIsCRFBackup interface{}

    // The additional type information of this VLAN. vtpVlanEditTypeExt object is
    // superseded by vtpVlanEditTypeExt2. The type is map[string]bool.
    VtpVlanEditTypeExt interface{}

    // The additional type information of this VLAN. The VlanTypeExt TC specifies
    // which bits may be written by a management application. The agent should
    // provide a default value. The type is map[string]bool.
    VtpVlanEditTypeExt2 interface{}

    // The MISTP instance, to which the corresponding vlan would be  mapped. The
    // value of this mib object is from 0 to the value of stpxMISTPInstanceNumber.
    // If setting the value of this object to 0, the corresponding vlan will not
    // be mapped to a MISTP  instance and all the ports under this VLAN will be
    // moved into the blocking state. The type is interface{} with range: 0..256.
    StpxVlanMISTPInstMapEditInstIndex interface{}
}

func (vtpVlanEditEntry *CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry) GetEntityData() *types.CommonEntityData {
    vtpVlanEditEntry.EntityData.YFilter = vtpVlanEditEntry.YFilter
    vtpVlanEditEntry.EntityData.YangName = "vtpVlanEditEntry"
    vtpVlanEditEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanEditEntry.EntityData.ParentYangName = "vtpVlanEditTable"
    vtpVlanEditEntry.EntityData.SegmentPath = "vtpVlanEditEntry" + types.AddKeyToken(vtpVlanEditEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpVlanEditEntry.VtpVlanEditIndex, "vtpVlanEditIndex")
    vtpVlanEditEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpVlanEditTable/" + vtpVlanEditEntry.EntityData.SegmentPath
    vtpVlanEditEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanEditEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanEditEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanEditEntry.EntityData.Children = types.NewOrderedMap()
    vtpVlanEditEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpVlanEditEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpVlanEditEntry.ManagementDomainIndex})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditIndex", types.YLeaf{"VtpVlanEditIndex", vtpVlanEditEntry.VtpVlanEditIndex})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditState", types.YLeaf{"VtpVlanEditState", vtpVlanEditEntry.VtpVlanEditState})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditType", types.YLeaf{"VtpVlanEditType", vtpVlanEditEntry.VtpVlanEditType})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditName", types.YLeaf{"VtpVlanEditName", vtpVlanEditEntry.VtpVlanEditName})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditMtu", types.YLeaf{"VtpVlanEditMtu", vtpVlanEditEntry.VtpVlanEditMtu})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditDot10Said", types.YLeaf{"VtpVlanEditDot10Said", vtpVlanEditEntry.VtpVlanEditDot10Said})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditRingNumber", types.YLeaf{"VtpVlanEditRingNumber", vtpVlanEditEntry.VtpVlanEditRingNumber})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditBridgeNumber", types.YLeaf{"VtpVlanEditBridgeNumber", vtpVlanEditEntry.VtpVlanEditBridgeNumber})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditStpType", types.YLeaf{"VtpVlanEditStpType", vtpVlanEditEntry.VtpVlanEditStpType})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditParentVlan", types.YLeaf{"VtpVlanEditParentVlan", vtpVlanEditEntry.VtpVlanEditParentVlan})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditRowStatus", types.YLeaf{"VtpVlanEditRowStatus", vtpVlanEditEntry.VtpVlanEditRowStatus})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditTranslationalVlan1", types.YLeaf{"VtpVlanEditTranslationalVlan1", vtpVlanEditEntry.VtpVlanEditTranslationalVlan1})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditTranslationalVlan2", types.YLeaf{"VtpVlanEditTranslationalVlan2", vtpVlanEditEntry.VtpVlanEditTranslationalVlan2})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditBridgeType", types.YLeaf{"VtpVlanEditBridgeType", vtpVlanEditEntry.VtpVlanEditBridgeType})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditAreHopCount", types.YLeaf{"VtpVlanEditAreHopCount", vtpVlanEditEntry.VtpVlanEditAreHopCount})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditSteHopCount", types.YLeaf{"VtpVlanEditSteHopCount", vtpVlanEditEntry.VtpVlanEditSteHopCount})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditIsCRFBackup", types.YLeaf{"VtpVlanEditIsCRFBackup", vtpVlanEditEntry.VtpVlanEditIsCRFBackup})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditTypeExt", types.YLeaf{"VtpVlanEditTypeExt", vtpVlanEditEntry.VtpVlanEditTypeExt})
    vtpVlanEditEntry.EntityData.Leafs.Append("vtpVlanEditTypeExt2", types.YLeaf{"VtpVlanEditTypeExt2", vtpVlanEditEntry.VtpVlanEditTypeExt2})
    vtpVlanEditEntry.EntityData.Leafs.Append("stpxVlanMISTPInstMapEditInstIndex", types.YLeaf{"StpxVlanMISTPInstMapEditInstIndex", vtpVlanEditEntry.StpxVlanMISTPInstMapEditInstIndex})

    vtpVlanEditEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpVlanEditIndex"}

    return &(vtpVlanEditEntry.EntityData)
}

// CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType represents is in use on this VLAN.
type CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType string

const (
    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType_srt CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType = "srt"

    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType_srb CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditBridgeType = "srb"
)

// CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState represents The state which this VLAN would have.
type CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState string

const (
    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState_operational CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState = "operational"

    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState_suspended CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditState = "suspended"
)

// CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType represents a 'hybrid' STP (see vtpVlanStpType).
type CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType string

const (
    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType_ieee CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType = "ieee"

    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType_ibm CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType = "ibm"

    CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType_auto CISCOVTPMIB_VtpVlanEditTable_VtpVlanEditEntry_VtpVlanEditStpType = "auto"
)

// CISCOVTPMIB_VtpVlanLocalShutdownTable
// Ths table contains the VLAN local shutdown
// information within management domain.
type CISCOVTPMIB_VtpVlanLocalShutdownTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing VLAN local shutdown information for a particular VLAN
    // in the management domain.  An entry is created if a VLAN which supports
    // local shutdown has been created.  An entry is deleted if a VLAN which
    // supports local shutdown has been removed. The type is slice of
    // CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry.
    VtpVlanLocalShutdownEntry []*CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry
}

func (vtpVlanLocalShutdownTable *CISCOVTPMIB_VtpVlanLocalShutdownTable) GetEntityData() *types.CommonEntityData {
    vtpVlanLocalShutdownTable.EntityData.YFilter = vtpVlanLocalShutdownTable.YFilter
    vtpVlanLocalShutdownTable.EntityData.YangName = "vtpVlanLocalShutdownTable"
    vtpVlanLocalShutdownTable.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanLocalShutdownTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpVlanLocalShutdownTable.EntityData.SegmentPath = "vtpVlanLocalShutdownTable"
    vtpVlanLocalShutdownTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpVlanLocalShutdownTable.EntityData.SegmentPath
    vtpVlanLocalShutdownTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanLocalShutdownTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanLocalShutdownTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanLocalShutdownTable.EntityData.Children = types.NewOrderedMap()
    vtpVlanLocalShutdownTable.EntityData.Children.Append("vtpVlanLocalShutdownEntry", types.YChild{"VtpVlanLocalShutdownEntry", nil})
    for i := range vtpVlanLocalShutdownTable.VtpVlanLocalShutdownEntry {
        vtpVlanLocalShutdownTable.EntityData.Children.Append(types.GetSegmentPath(vtpVlanLocalShutdownTable.VtpVlanLocalShutdownEntry[i]), types.YChild{"VtpVlanLocalShutdownEntry", vtpVlanLocalShutdownTable.VtpVlanLocalShutdownEntry[i]})
    }
    vtpVlanLocalShutdownTable.EntityData.Leafs = types.NewOrderedMap()

    vtpVlanLocalShutdownTable.EntityData.YListKeys = []string {}

    return &(vtpVlanLocalShutdownTable.EntityData)
}

// CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry
// An entry containing VLAN local shutdown information for a
// particular VLAN in the management domain.
// 
// An entry is created if a VLAN which supports local shutdown
// has been created.
// 
// An entry is deleted if a VLAN which supports local shutdown
// has been removed.
type CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_VtpVlanTable_VtpVlanEntry_VtpVlanIndex
    VtpVlanIndex interface{}

    // The object specifies the VLAN local shutdown state. The type is
    // VtpVlanLocalShutdown.
    VtpVlanLocalShutdown interface{}
}

func (vtpVlanLocalShutdownEntry *CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry) GetEntityData() *types.CommonEntityData {
    vtpVlanLocalShutdownEntry.EntityData.YFilter = vtpVlanLocalShutdownEntry.YFilter
    vtpVlanLocalShutdownEntry.EntityData.YangName = "vtpVlanLocalShutdownEntry"
    vtpVlanLocalShutdownEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpVlanLocalShutdownEntry.EntityData.ParentYangName = "vtpVlanLocalShutdownTable"
    vtpVlanLocalShutdownEntry.EntityData.SegmentPath = "vtpVlanLocalShutdownEntry" + types.AddKeyToken(vtpVlanLocalShutdownEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpVlanLocalShutdownEntry.VtpVlanIndex, "vtpVlanIndex")
    vtpVlanLocalShutdownEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpVlanLocalShutdownTable/" + vtpVlanLocalShutdownEntry.EntityData.SegmentPath
    vtpVlanLocalShutdownEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpVlanLocalShutdownEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpVlanLocalShutdownEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpVlanLocalShutdownEntry.EntityData.Children = types.NewOrderedMap()
    vtpVlanLocalShutdownEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpVlanLocalShutdownEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpVlanLocalShutdownEntry.ManagementDomainIndex})
    vtpVlanLocalShutdownEntry.EntityData.Leafs.Append("vtpVlanIndex", types.YLeaf{"VtpVlanIndex", vtpVlanLocalShutdownEntry.VtpVlanIndex})
    vtpVlanLocalShutdownEntry.EntityData.Leafs.Append("vtpVlanLocalShutdown", types.YLeaf{"VtpVlanLocalShutdown", vtpVlanLocalShutdownEntry.VtpVlanLocalShutdown})

    vtpVlanLocalShutdownEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpVlanIndex"}

    return &(vtpVlanLocalShutdownEntry.EntityData)
}

// CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown represents The object specifies the VLAN local shutdown state.
type CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown string

const (
    CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown_up CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown = "up"

    CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown_down CISCOVTPMIB_VtpVlanLocalShutdownTable_VtpVlanLocalShutdownEntry_VtpVlanLocalShutdown = "down"
)

// CISCOVTPMIB_VlanTrunkPortTable
// The table containing information on the local system's VLAN
// trunk ports.
type CISCOVTPMIB_VlanTrunkPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about one trunk port. The type is slice of
    // CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry.
    VlanTrunkPortEntry []*CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry
}

func (vlanTrunkPortTable *CISCOVTPMIB_VlanTrunkPortTable) GetEntityData() *types.CommonEntityData {
    vlanTrunkPortTable.EntityData.YFilter = vlanTrunkPortTable.YFilter
    vlanTrunkPortTable.EntityData.YangName = "vlanTrunkPortTable"
    vlanTrunkPortTable.EntityData.BundleName = "cisco_ios_xe"
    vlanTrunkPortTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vlanTrunkPortTable.EntityData.SegmentPath = "vlanTrunkPortTable"
    vlanTrunkPortTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vlanTrunkPortTable.EntityData.SegmentPath
    vlanTrunkPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanTrunkPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanTrunkPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanTrunkPortTable.EntityData.Children = types.NewOrderedMap()
    vlanTrunkPortTable.EntityData.Children.Append("vlanTrunkPortEntry", types.YChild{"VlanTrunkPortEntry", nil})
    for i := range vlanTrunkPortTable.VlanTrunkPortEntry {
        vlanTrunkPortTable.EntityData.Children.Append(types.GetSegmentPath(vlanTrunkPortTable.VlanTrunkPortEntry[i]), types.YChild{"VlanTrunkPortEntry", vlanTrunkPortTable.VlanTrunkPortEntry[i]})
    }
    vlanTrunkPortTable.EntityData.Leafs = types.NewOrderedMap()

    vlanTrunkPortTable.EntityData.YListKeys = []string {}

    return &(vlanTrunkPortTable.EntityData)
}

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry
// Information about one trunk port.
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The value of ifIndex for the interface
    // corresponding to this trunk port. The type is interface{} with range:
    // 1..2147483647.
    VlanTrunkPortIfIndex interface{}

    // The value of managementDomainIndex for the management domain on this trunk
    // port.  Devices which support only one management domain will support this
    // object read-only. The type is interface{} with range: 1..255.
    VlanTrunkPortManagementDomain interface{}

    // The type of VLAN encapsulation desired to be used on this trunk port. It is
    // either a particular type, or 'negotiate' meaning whatever type results from
    // the negotiation. negotiate(5) is not allowed if the port does not support
    // negotiation or if its vlanTrunkPortDynamicState is set to on(1) or
    // onNoNegotiate(5). Whether writing to this object in order to modify the
    // encapsulation is supported is both device and interface specific. The type
    // is VlanTrunkPortEncapsulationType.
    VlanTrunkPortEncapsulationType interface{}

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
    VlanTrunkPortVlansEnabled interface{}

    // The VlanIndex of the VLAN which is represented by native frames on this
    // trunk port.  For trunk ports not supporting the sending and receiving of
    // native frames, this value should be set to zero. The type is interface{}
    // with range: 0..4095.
    VlanTrunkPortNativeVlan interface{}

    // The status of this row.  In some circumstances, the creation of a row in
    // this table is needed to enable the appropriate trunking/tagging protocol on
    // the port, to enable the use of VTP on the port, and to assign the port to
    // the appropriate management domain.  In other circumstances, rows in this
    // table will be created as a by-product of other operations. The type is
    // RowStatus.
    VlanTrunkPortRowStatus interface{}

    // The number of VTP Join messages received on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    VlanTrunkPortInJoins interface{}

    // The number of VTP Join messages sent on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    VlanTrunkPortOutJoins interface{}

    // The number of VTP Advertisement messages which indicated the sender does
    // not support VLAN-pruning received on this trunk port. The type is
    // interface{} with range: 0..4294967295.
    VlanTrunkPortOldAdverts interface{}

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
    VlanTrunkPortVlansPruningEligible interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then this VLAN is presently being forwarded on this trunk port, i.e., it is
    // not pruned; if the bit is set to '0', then this VLAN is presently not being
    // forwarded on this trunk port, either because it is pruned or for some other
    // reason. The type is string with length: 128.
    VlanTrunkPortVlansXmitJoined interface{}

    // A string of octets containing one bit per VLAN in the management domain on
    // this trunk port.  The first octet corresponds to VLANs with VlanIndex
    // values of 0 through 7; the second octet to VLANs 8 through 15; etc.  The
    // most significant bit of each octet corresponds to the lowest value
    // VlanIndex in that octet.  If the bit corresponding to a VLAN is set to '1',
    // then the local switch is currently sending joins for this VLAN on this
    // trunk port, i.e., it is asking to receive frames for this VLAN; if the bit
    // is set to '0', then the local switch is not currently sending joins for
    // this VLAN on this trunk port. The type is string with length: 128.
    VlanTrunkPortVlansRcvJoined interface{}

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
    // type is VlanTrunkPortDynamicState.
    VlanTrunkPortDynamicState interface{}

    // Indicates whether the specified interface is either acting as a trunk or
    // not. This is a result of the vlanTrunkPortDynamicState and the ifOperStatus
    // of the trunk port itself. The type is VlanTrunkPortDynamicStatus.
    VlanTrunkPortDynamicStatus interface{}

    // Some trunk interface modules allow VTP to be enabled/disabled seperately
    // from that of the central device.  In such a case this object provides
    // management a way to remotely enable VTP on that module.  If a module does
    // not support a seperate VTP enabled state then this object shall always
    // return 'true' and will accept no other value during a SET operation. The
    // type is bool.
    VlanTrunkPortVtpEnabled interface{}

    // The type of VLAN encapsulation in use on this trunk port. For intefaces
    // with vlanTrunkPortDynamicStatus of notTrunking(2) the
    // vlanTrunkPortEncapsulationOperType shall be notApplicable(6). The type is
    // VlanTrunkPortEncapsulationOperType.
    VlanTrunkPortEncapsulationOperType interface{}

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
    VlanTrunkPortVlansEnabled2k interface{}

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
    VlanTrunkPortVlansEnabled3k interface{}

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
    VlanTrunkPortVlansEnabled4k interface{}

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
    VtpVlansPruningEligible2k interface{}

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
    VtpVlansPruningEligible3k interface{}

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
    VtpVlansPruningEligible4k interface{}

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
    VlanTrunkPortVlansXmitJoined2k interface{}

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
    VlanTrunkPortVlansXmitJoined3k interface{}

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
    VlanTrunkPortVlansXmitJoined4k interface{}

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
    VlanTrunkPortVlansRcvJoined2k interface{}

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
    VlanTrunkPortVlansRcvJoined3k interface{}

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
    VlanTrunkPortVlansRcvJoined4k interface{}

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
    // the CISCO-L2-TUNNEL-CONFIG-MIB. The type is VlanTrunkPortDot1qTunnel.
    VlanTrunkPortDot1qTunnel interface{}

    // A string of octets containing one bit per VLAN with VlanIndex values of 0
    // through 2047.  If the bit corresponding to a VLAN is set to 1, it indicates
    // that vlan is allowed and active in management domain.  If the bit
    // corresponding to a VLAN is set to 0, it indicates that vlan is not allowed
    // or not active in management domain. The type is string with length: 0..256.
    VlanTrunkPortVlansActiveFirst2k interface{}

    // A string of octets containing one bit per VLAN with VlanIndex values of
    // 2048 through 4095.  If the bit corresponding to a VLAN is set to 1, it
    // indicates that vlan is allowed and active in management domain.  If the bit
    // corresponding to a VLAN is set to 0, it indicates that vlan is not allowed
    // or not active in management domain. The type is string with length: 0..256.
    VlanTrunkPortVlansActiveSecond2k interface{}

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
    StpxPreferredVlansMap interface{}

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
    StpxPreferredVlansMap2k interface{}

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
    StpxPreferredVlansMap3k interface{}

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
    StpxPreferredVlansMap4k interface{}

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
    StpxPreferredMISTPInstancesMap interface{}

    // A string of octets containing one bit per MST instances  on this trunk
    // port.  The first octet corresponds to MST  instances of 0 through 7; the
    // second octet to MST instances  8 through 15; etc. The most significant bit
    // of each octet  corresponds to the lowest MST instance value in that octet. 
    // The number of bits for this mib object will be determined  by the value of
    // stpxMSTMaxInstanceNumber.  For each instance, if it is preferred on this
    // trunk port,  then the bit corresponding to that instance is set to '1'. The
    // type is string with length: 1..32.
    StpxPreferredMSTInstancesMap interface{}
}

func (vlanTrunkPortEntry *CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry) GetEntityData() *types.CommonEntityData {
    vlanTrunkPortEntry.EntityData.YFilter = vlanTrunkPortEntry.YFilter
    vlanTrunkPortEntry.EntityData.YangName = "vlanTrunkPortEntry"
    vlanTrunkPortEntry.EntityData.BundleName = "cisco_ios_xe"
    vlanTrunkPortEntry.EntityData.ParentYangName = "vlanTrunkPortTable"
    vlanTrunkPortEntry.EntityData.SegmentPath = "vlanTrunkPortEntry" + types.AddKeyToken(vlanTrunkPortEntry.VlanTrunkPortIfIndex, "vlanTrunkPortIfIndex")
    vlanTrunkPortEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vlanTrunkPortTable/" + vlanTrunkPortEntry.EntityData.SegmentPath
    vlanTrunkPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vlanTrunkPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vlanTrunkPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vlanTrunkPortEntry.EntityData.Children = types.NewOrderedMap()
    vlanTrunkPortEntry.EntityData.Leafs = types.NewOrderedMap()
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortIfIndex", types.YLeaf{"VlanTrunkPortIfIndex", vlanTrunkPortEntry.VlanTrunkPortIfIndex})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortManagementDomain", types.YLeaf{"VlanTrunkPortManagementDomain", vlanTrunkPortEntry.VlanTrunkPortManagementDomain})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortEncapsulationType", types.YLeaf{"VlanTrunkPortEncapsulationType", vlanTrunkPortEntry.VlanTrunkPortEncapsulationType})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansEnabled", types.YLeaf{"VlanTrunkPortVlansEnabled", vlanTrunkPortEntry.VlanTrunkPortVlansEnabled})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortNativeVlan", types.YLeaf{"VlanTrunkPortNativeVlan", vlanTrunkPortEntry.VlanTrunkPortNativeVlan})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortRowStatus", types.YLeaf{"VlanTrunkPortRowStatus", vlanTrunkPortEntry.VlanTrunkPortRowStatus})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortInJoins", types.YLeaf{"VlanTrunkPortInJoins", vlanTrunkPortEntry.VlanTrunkPortInJoins})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortOutJoins", types.YLeaf{"VlanTrunkPortOutJoins", vlanTrunkPortEntry.VlanTrunkPortOutJoins})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortOldAdverts", types.YLeaf{"VlanTrunkPortOldAdverts", vlanTrunkPortEntry.VlanTrunkPortOldAdverts})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansPruningEligible", types.YLeaf{"VlanTrunkPortVlansPruningEligible", vlanTrunkPortEntry.VlanTrunkPortVlansPruningEligible})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansXmitJoined", types.YLeaf{"VlanTrunkPortVlansXmitJoined", vlanTrunkPortEntry.VlanTrunkPortVlansXmitJoined})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansRcvJoined", types.YLeaf{"VlanTrunkPortVlansRcvJoined", vlanTrunkPortEntry.VlanTrunkPortVlansRcvJoined})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortDynamicState", types.YLeaf{"VlanTrunkPortDynamicState", vlanTrunkPortEntry.VlanTrunkPortDynamicState})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortDynamicStatus", types.YLeaf{"VlanTrunkPortDynamicStatus", vlanTrunkPortEntry.VlanTrunkPortDynamicStatus})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVtpEnabled", types.YLeaf{"VlanTrunkPortVtpEnabled", vlanTrunkPortEntry.VlanTrunkPortVtpEnabled})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortEncapsulationOperType", types.YLeaf{"VlanTrunkPortEncapsulationOperType", vlanTrunkPortEntry.VlanTrunkPortEncapsulationOperType})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansEnabled2k", types.YLeaf{"VlanTrunkPortVlansEnabled2k", vlanTrunkPortEntry.VlanTrunkPortVlansEnabled2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansEnabled3k", types.YLeaf{"VlanTrunkPortVlansEnabled3k", vlanTrunkPortEntry.VlanTrunkPortVlansEnabled3k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansEnabled4k", types.YLeaf{"VlanTrunkPortVlansEnabled4k", vlanTrunkPortEntry.VlanTrunkPortVlansEnabled4k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vtpVlansPruningEligible2k", types.YLeaf{"VtpVlansPruningEligible2k", vlanTrunkPortEntry.VtpVlansPruningEligible2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vtpVlansPruningEligible3k", types.YLeaf{"VtpVlansPruningEligible3k", vlanTrunkPortEntry.VtpVlansPruningEligible3k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vtpVlansPruningEligible4k", types.YLeaf{"VtpVlansPruningEligible4k", vlanTrunkPortEntry.VtpVlansPruningEligible4k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansXmitJoined2k", types.YLeaf{"VlanTrunkPortVlansXmitJoined2k", vlanTrunkPortEntry.VlanTrunkPortVlansXmitJoined2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansXmitJoined3k", types.YLeaf{"VlanTrunkPortVlansXmitJoined3k", vlanTrunkPortEntry.VlanTrunkPortVlansXmitJoined3k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansXmitJoined4k", types.YLeaf{"VlanTrunkPortVlansXmitJoined4k", vlanTrunkPortEntry.VlanTrunkPortVlansXmitJoined4k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansRcvJoined2k", types.YLeaf{"VlanTrunkPortVlansRcvJoined2k", vlanTrunkPortEntry.VlanTrunkPortVlansRcvJoined2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansRcvJoined3k", types.YLeaf{"VlanTrunkPortVlansRcvJoined3k", vlanTrunkPortEntry.VlanTrunkPortVlansRcvJoined3k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansRcvJoined4k", types.YLeaf{"VlanTrunkPortVlansRcvJoined4k", vlanTrunkPortEntry.VlanTrunkPortVlansRcvJoined4k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortDot1qTunnel", types.YLeaf{"VlanTrunkPortDot1qTunnel", vlanTrunkPortEntry.VlanTrunkPortDot1qTunnel})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansActiveFirst2k", types.YLeaf{"VlanTrunkPortVlansActiveFirst2k", vlanTrunkPortEntry.VlanTrunkPortVlansActiveFirst2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("vlanTrunkPortVlansActiveSecond2k", types.YLeaf{"VlanTrunkPortVlansActiveSecond2k", vlanTrunkPortEntry.VlanTrunkPortVlansActiveSecond2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredVlansMap", types.YLeaf{"StpxPreferredVlansMap", vlanTrunkPortEntry.StpxPreferredVlansMap})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredVlansMap2k", types.YLeaf{"StpxPreferredVlansMap2k", vlanTrunkPortEntry.StpxPreferredVlansMap2k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredVlansMap3k", types.YLeaf{"StpxPreferredVlansMap3k", vlanTrunkPortEntry.StpxPreferredVlansMap3k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredVlansMap4k", types.YLeaf{"StpxPreferredVlansMap4k", vlanTrunkPortEntry.StpxPreferredVlansMap4k})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredMISTPInstancesMap", types.YLeaf{"StpxPreferredMISTPInstancesMap", vlanTrunkPortEntry.StpxPreferredMISTPInstancesMap})
    vlanTrunkPortEntry.EntityData.Leafs.Append("stpxPreferredMSTInstancesMap", types.YLeaf{"StpxPreferredMSTInstancesMap", vlanTrunkPortEntry.StpxPreferredMSTInstancesMap})

    vlanTrunkPortEntry.EntityData.YListKeys = []string {"VlanTrunkPortIfIndex"}

    return &(vlanTrunkPortEntry.EntityData)
}

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel represents CISCO-L2-TUNNEL-CONFIG-MIB
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel string

const (
    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel_trunk CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel = "trunk"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel_access CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel = "access"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel_disabled CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDot1qTunnel = "disabled"
)

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState represents device) need only support the 'on', and 'off' values.
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState string

const (
    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState_on CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState = "on"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState_off CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState = "off"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState_desirable CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState = "desirable"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState_auto CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState = "auto"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState_onNoNegotiate CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicState = "onNoNegotiate"
)

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus represents trunk port itself.
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus string

const (
    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus_trunking CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus = "trunking"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus_notTrunking CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortDynamicStatus = "notTrunking"
)

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType represents be notApplicable(6).
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType string

const (
    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_isl CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "isl"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_dot10 CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "dot10"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_lane CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "lane"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_dot1Q CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "dot1Q"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_negotiating CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "negotiating"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType_notApplicable CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationOperType = "notApplicable"
)

// CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType represents device and interface specific.
type CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType string

const (
    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType_isl CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType = "isl"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType_dot10 CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType = "dot10"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType_lane CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType = "lane"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType_dot1Q CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType = "dot1Q"

    CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType_negotiate CISCOVTPMIB_VlanTrunkPortTable_VlanTrunkPortEntry_VlanTrunkPortEncapsulationType = "negotiate"
)

// CISCOVTPMIB_VtpDiscoverTable
// This table contains information related to the discovery
// of the VTP members in the designated management
// domain. This table is not instantiated when 
// managementDomainVersionInUse is version1(1), version2(3)
// or none(3).
type CISCOVTPMIB_VtpDiscoverTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information related to the discovery of the VTP members in one management
    // domain. The type is slice of CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry.
    VtpDiscoverEntry []*CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry
}

func (vtpDiscoverTable *CISCOVTPMIB_VtpDiscoverTable) GetEntityData() *types.CommonEntityData {
    vtpDiscoverTable.EntityData.YFilter = vtpDiscoverTable.YFilter
    vtpDiscoverTable.EntityData.YangName = "vtpDiscoverTable"
    vtpDiscoverTable.EntityData.BundleName = "cisco_ios_xe"
    vtpDiscoverTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpDiscoverTable.EntityData.SegmentPath = "vtpDiscoverTable"
    vtpDiscoverTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpDiscoverTable.EntityData.SegmentPath
    vtpDiscoverTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDiscoverTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDiscoverTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDiscoverTable.EntityData.Children = types.NewOrderedMap()
    vtpDiscoverTable.EntityData.Children.Append("vtpDiscoverEntry", types.YChild{"VtpDiscoverEntry", nil})
    for i := range vtpDiscoverTable.VtpDiscoverEntry {
        vtpDiscoverTable.EntityData.Children.Append(types.GetSegmentPath(vtpDiscoverTable.VtpDiscoverEntry[i]), types.YChild{"VtpDiscoverEntry", vtpDiscoverTable.VtpDiscoverEntry[i]})
    }
    vtpDiscoverTable.EntityData.Leafs = types.NewOrderedMap()

    vtpDiscoverTable.EntityData.YListKeys = []string {}

    return &(vtpDiscoverTable.EntityData)
}

// CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry
// Information related to the discovery of the
// VTP members in one management domain.
type CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

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
    // always returns noOperation(2). The type is VtpDiscoverAction.
    VtpDiscoverAction interface{}

    // The current status of VTP discovery.  inProgress - a discovery is in
    // progress;  succeeded - the discovery was completed successfully            
    // (this value is also used when              no discover has been invoked
    // since the             last time the local system restarted); 
    // resourceUnavailable - the discovery failed because             the required
    // allocation of a resource is             presently unavailable. 
    // someOtherError - 'the discovery failed due to a             reason no
    // listed. The type is VtpDiscoverStatus.
    VtpDiscoverStatus interface{}

    // The value of sysUpTime at which the last discovery was completed.  A value
    // of zero indicates that no discovery has been invoked since last time the
    // local system restarted. The type is interface{} with range: 0..4294967295.
    VtpLastDiscoverTime interface{}
}

func (vtpDiscoverEntry *CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry) GetEntityData() *types.CommonEntityData {
    vtpDiscoverEntry.EntityData.YFilter = vtpDiscoverEntry.YFilter
    vtpDiscoverEntry.EntityData.YangName = "vtpDiscoverEntry"
    vtpDiscoverEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpDiscoverEntry.EntityData.ParentYangName = "vtpDiscoverTable"
    vtpDiscoverEntry.EntityData.SegmentPath = "vtpDiscoverEntry" + types.AddKeyToken(vtpDiscoverEntry.ManagementDomainIndex, "managementDomainIndex")
    vtpDiscoverEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpDiscoverTable/" + vtpDiscoverEntry.EntityData.SegmentPath
    vtpDiscoverEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDiscoverEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDiscoverEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDiscoverEntry.EntityData.Children = types.NewOrderedMap()
    vtpDiscoverEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpDiscoverEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpDiscoverEntry.ManagementDomainIndex})
    vtpDiscoverEntry.EntityData.Leafs.Append("vtpDiscoverAction", types.YLeaf{"VtpDiscoverAction", vtpDiscoverEntry.VtpDiscoverAction})
    vtpDiscoverEntry.EntityData.Leafs.Append("vtpDiscoverStatus", types.YLeaf{"VtpDiscoverStatus", vtpDiscoverEntry.VtpDiscoverStatus})
    vtpDiscoverEntry.EntityData.Leafs.Append("vtpLastDiscoverTime", types.YLeaf{"VtpLastDiscoverTime", vtpDiscoverEntry.VtpLastDiscoverTime})

    vtpDiscoverEntry.EntityData.YListKeys = []string {"ManagementDomainIndex"}

    return &(vtpDiscoverEntry.EntityData)
}

// CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction represents always returns noOperation(2).
type CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction string

const (
    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction_discover CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction = "discover"

    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction_noOperation CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction = "noOperation"

    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction_purgeResult CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverAction = "purgeResult"
)

// CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus represents             reason no listed.
type CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus string

const (
    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus_inProgress CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus = "inProgress"

    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus_succeeded CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus = "succeeded"

    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus_resourceUnavailable CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus = "resourceUnavailable"

    CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus_someOtherError CISCOVTPMIB_VtpDiscoverTable_VtpDiscoverEntry_VtpDiscoverStatus = "someOtherError"
)

// CISCOVTPMIB_VtpDiscoverResultTable
// The table containing information of discovered VTP members
// in the management domain in which the local system is
// participating. This table is not instantiated when 
// managementDomainVersionInUse is version1(1), version2(2) or
// none(3).
type CISCOVTPMIB_VtpDiscoverResultTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row is created for each VTP member which is found through
    // successful discovery. The type is slice of
    // CISCOVTPMIB_VtpDiscoverResultTable_VtpDiscoverResultEntry.
    VtpDiscoverResultEntry []*CISCOVTPMIB_VtpDiscoverResultTable_VtpDiscoverResultEntry
}

func (vtpDiscoverResultTable *CISCOVTPMIB_VtpDiscoverResultTable) GetEntityData() *types.CommonEntityData {
    vtpDiscoverResultTable.EntityData.YFilter = vtpDiscoverResultTable.YFilter
    vtpDiscoverResultTable.EntityData.YangName = "vtpDiscoverResultTable"
    vtpDiscoverResultTable.EntityData.BundleName = "cisco_ios_xe"
    vtpDiscoverResultTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpDiscoverResultTable.EntityData.SegmentPath = "vtpDiscoverResultTable"
    vtpDiscoverResultTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpDiscoverResultTable.EntityData.SegmentPath
    vtpDiscoverResultTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDiscoverResultTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDiscoverResultTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDiscoverResultTable.EntityData.Children = types.NewOrderedMap()
    vtpDiscoverResultTable.EntityData.Children.Append("vtpDiscoverResultEntry", types.YChild{"VtpDiscoverResultEntry", nil})
    for i := range vtpDiscoverResultTable.VtpDiscoverResultEntry {
        vtpDiscoverResultTable.EntityData.Children.Append(types.GetSegmentPath(vtpDiscoverResultTable.VtpDiscoverResultEntry[i]), types.YChild{"VtpDiscoverResultEntry", vtpDiscoverResultTable.VtpDiscoverResultEntry[i]})
    }
    vtpDiscoverResultTable.EntityData.Leafs = types.NewOrderedMap()

    vtpDiscoverResultTable.EntityData.YListKeys = []string {}

    return &(vtpDiscoverResultTable.EntityData)
}

// CISCOVTPMIB_VtpDiscoverResultTable_VtpDiscoverResultEntry
// A conceptual row is created for each VTP member which
// is found through successful discovery.
type CISCOVTPMIB_VtpDiscoverResultTable_VtpDiscoverResultEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. A value assigned by the system which identifies a
    // VTP member and the associated database in the  management domain. The type
    // is interface{} with range: 0..4294967295.
    VtpDiscoverResultIndex interface{}

    // The database name associated with the discovered VTP member. The type is
    // string with length: 0..64.
    VtpDiscoverResultDatabaseName interface{}

    // Indicates whether this VTP member contains conflicting information. 
    // true(1) indicates that this member has conflicting  information of the
    // database type in the management domain.  false(2) indicates that there is
    // no conflicting information of the database type in the management domain.
    // The type is bool.
    VtpDiscoverResultConflicting interface{}

    // The unique identifier of the device for this VTP member. The type is string
    // with length: 0..64.
    VtpDiscoverResultDeviceId interface{}

    // The unique identifier of the primary server for this VTP member and the
    // associated database type.  There are two different VTP servers, the primary
    // server and the secondary server.  When a local device is configured as a
    // server for a certain database type, it becomes secondary server by default.
    // Primary server is an operational role under which a server can initiate or
    // change the VTP configuration of the database type.  If this VTP member
    // itself is the primary server, the    value of this object is the same as
    // the value of  vtpDiscoverResultDeviceId of the instance. The type is string
    // with length: 0..64.
    VtpDiscoverResultPrimaryServer interface{}

    // The current configuration revision number as known by the VTP member. When
    // the database type is unknown for the VTP member, this value is 0. The type
    // is interface{} with range: 0..4294967295.
    VtpDiscoverResultRevNumber interface{}

    // sysName of the VTP member. The type is string with length: 0..64.
    VtpDiscoverResultSystemName interface{}
}

func (vtpDiscoverResultEntry *CISCOVTPMIB_VtpDiscoverResultTable_VtpDiscoverResultEntry) GetEntityData() *types.CommonEntityData {
    vtpDiscoverResultEntry.EntityData.YFilter = vtpDiscoverResultEntry.YFilter
    vtpDiscoverResultEntry.EntityData.YangName = "vtpDiscoverResultEntry"
    vtpDiscoverResultEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpDiscoverResultEntry.EntityData.ParentYangName = "vtpDiscoverResultTable"
    vtpDiscoverResultEntry.EntityData.SegmentPath = "vtpDiscoverResultEntry" + types.AddKeyToken(vtpDiscoverResultEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpDiscoverResultEntry.VtpDiscoverResultIndex, "vtpDiscoverResultIndex")
    vtpDiscoverResultEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpDiscoverResultTable/" + vtpDiscoverResultEntry.EntityData.SegmentPath
    vtpDiscoverResultEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDiscoverResultEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDiscoverResultEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDiscoverResultEntry.EntityData.Children = types.NewOrderedMap()
    vtpDiscoverResultEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpDiscoverResultEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpDiscoverResultEntry.ManagementDomainIndex})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultIndex", types.YLeaf{"VtpDiscoverResultIndex", vtpDiscoverResultEntry.VtpDiscoverResultIndex})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultDatabaseName", types.YLeaf{"VtpDiscoverResultDatabaseName", vtpDiscoverResultEntry.VtpDiscoverResultDatabaseName})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultConflicting", types.YLeaf{"VtpDiscoverResultConflicting", vtpDiscoverResultEntry.VtpDiscoverResultConflicting})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultDeviceId", types.YLeaf{"VtpDiscoverResultDeviceId", vtpDiscoverResultEntry.VtpDiscoverResultDeviceId})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultPrimaryServer", types.YLeaf{"VtpDiscoverResultPrimaryServer", vtpDiscoverResultEntry.VtpDiscoverResultPrimaryServer})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultRevNumber", types.YLeaf{"VtpDiscoverResultRevNumber", vtpDiscoverResultEntry.VtpDiscoverResultRevNumber})
    vtpDiscoverResultEntry.EntityData.Leafs.Append("vtpDiscoverResultSystemName", types.YLeaf{"VtpDiscoverResultSystemName", vtpDiscoverResultEntry.VtpDiscoverResultSystemName})

    vtpDiscoverResultEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpDiscoverResultIndex"}

    return &(vtpDiscoverResultEntry.EntityData)
}

// CISCOVTPMIB_VtpDatabaseTable
// This table contains information of the VTP
// databases. It is not instantiated when
// managementDomainVersionInUse is version1(1), 
// version2(3) or none(3).
type CISCOVTPMIB_VtpDatabaseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of the VTP database in the domain.  Each VTP
    // database type known to the local device type has an entry in this table. An
    // entry is also created for unknown database which is notified through VTP
    // advertisements from other VTP servers. The type is slice of
    // CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry.
    VtpDatabaseEntry []*CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry
}

func (vtpDatabaseTable *CISCOVTPMIB_VtpDatabaseTable) GetEntityData() *types.CommonEntityData {
    vtpDatabaseTable.EntityData.YFilter = vtpDatabaseTable.YFilter
    vtpDatabaseTable.EntityData.YangName = "vtpDatabaseTable"
    vtpDatabaseTable.EntityData.BundleName = "cisco_ios_xe"
    vtpDatabaseTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpDatabaseTable.EntityData.SegmentPath = "vtpDatabaseTable"
    vtpDatabaseTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpDatabaseTable.EntityData.SegmentPath
    vtpDatabaseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDatabaseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDatabaseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDatabaseTable.EntityData.Children = types.NewOrderedMap()
    vtpDatabaseTable.EntityData.Children.Append("vtpDatabaseEntry", types.YChild{"VtpDatabaseEntry", nil})
    for i := range vtpDatabaseTable.VtpDatabaseEntry {
        vtpDatabaseTable.EntityData.Children.Append(types.GetSegmentPath(vtpDatabaseTable.VtpDatabaseEntry[i]), types.YChild{"VtpDatabaseEntry", vtpDatabaseTable.VtpDatabaseEntry[i]})
    }
    vtpDatabaseTable.EntityData.Leafs = types.NewOrderedMap()

    vtpDatabaseTable.EntityData.YListKeys = []string {}

    return &(vtpDatabaseTable.EntityData)
}

// CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry
// Information about the status of the VTP database
// in the domain.  Each VTP database type known to the
// local device type has an entry in this table.
// An entry is also created for unknown database which is
// notified through VTP advertisements from other VTP
// servers.
type CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

    // This attribute is a key. A value assigned by the system which uniquely
    // identifies a VTP database in the local system. The type is interface{} with
    // range: 0..4294967295.
    VtpDatabaseIndex interface{}

    // The name of the database. The type is string with length: 0..64.
    VtpDatabaseName interface{}

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
    // VtpDatabaseLocalMode.
    VtpDatabaseLocalMode interface{}

    // The current configuration revision number as known by the local device for
    // this VTP 3 database type in the management domain.  This value is updated
    // (if necessary) whenever a  VTP advertisement for the database type is
    // received  or generated. When the database type is unknown to the  local
    // device or no VTP advertisement for the database type is received or
    // generated, its value is 0. The type is interface{} with range:
    // 0..4294967295.
    VtpDatabaseRevNumber interface{}

    // There are two kinds of VTP version 3 servers for a certain database type -
    // the primary server and the secondary server. When a local device is
    // configured as a server for a certain database type, it becomes secondary
    // server by default. Primary server is an operational role under which a
    // server can initiate or change the VTP configuration of the database type. 
    // A true(1) value indicates that the local device is the  primary server of
    // the database type in the management domain. A false(2) value indicates that
    // the local device is not the primary server, or the database type is unknown
    // to the local device. The type is bool.
    VtpDatabasePrimaryServer interface{}

    // The unique identifier of the primary server in the management domain for
    // the database type.   If no primary server is discovered for the database
    // type, the object has a value of zero length string. The type is string with
    // length: 0..64.
    VtpDatabasePrimaryServerId interface{}

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
    VtpDatabaseTakeOverPrimary interface{}

    // When read, this object always returns the value of a zero-length octet
    // string.  In the case that the VTP password is hidden from the 
    // configuration and the local device intends to take over the whole domain,
    // this object must be  set to the matching password with the secret key 
    // (vtpAuthSecretKey) in the same data packet as which  the
    // vtpDatabaseTakeOverPrimary is in. In all the  other situations, setting a
    // valid value to this object  has no impact on the system. The type is string
    // with length: 0..64.
    VtpDatabaseTakeOverPassword interface{}
}

func (vtpDatabaseEntry *CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry) GetEntityData() *types.CommonEntityData {
    vtpDatabaseEntry.EntityData.YFilter = vtpDatabaseEntry.YFilter
    vtpDatabaseEntry.EntityData.YangName = "vtpDatabaseEntry"
    vtpDatabaseEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpDatabaseEntry.EntityData.ParentYangName = "vtpDatabaseTable"
    vtpDatabaseEntry.EntityData.SegmentPath = "vtpDatabaseEntry" + types.AddKeyToken(vtpDatabaseEntry.ManagementDomainIndex, "managementDomainIndex") + types.AddKeyToken(vtpDatabaseEntry.VtpDatabaseIndex, "vtpDatabaseIndex")
    vtpDatabaseEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpDatabaseTable/" + vtpDatabaseEntry.EntityData.SegmentPath
    vtpDatabaseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpDatabaseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpDatabaseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpDatabaseEntry.EntityData.Children = types.NewOrderedMap()
    vtpDatabaseEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpDatabaseEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpDatabaseEntry.ManagementDomainIndex})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseIndex", types.YLeaf{"VtpDatabaseIndex", vtpDatabaseEntry.VtpDatabaseIndex})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseName", types.YLeaf{"VtpDatabaseName", vtpDatabaseEntry.VtpDatabaseName})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseLocalMode", types.YLeaf{"VtpDatabaseLocalMode", vtpDatabaseEntry.VtpDatabaseLocalMode})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseRevNumber", types.YLeaf{"VtpDatabaseRevNumber", vtpDatabaseEntry.VtpDatabaseRevNumber})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabasePrimaryServer", types.YLeaf{"VtpDatabasePrimaryServer", vtpDatabaseEntry.VtpDatabasePrimaryServer})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabasePrimaryServerId", types.YLeaf{"VtpDatabasePrimaryServerId", vtpDatabaseEntry.VtpDatabasePrimaryServerId})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseTakeOverPrimary", types.YLeaf{"VtpDatabaseTakeOverPrimary", vtpDatabaseEntry.VtpDatabaseTakeOverPrimary})
    vtpDatabaseEntry.EntityData.Leafs.Append("vtpDatabaseTakeOverPassword", types.YLeaf{"VtpDatabaseTakeOverPassword", vtpDatabaseEntry.VtpDatabaseTakeOverPassword})

    vtpDatabaseEntry.EntityData.YListKeys = []string {"ManagementDomainIndex", "VtpDatabaseIndex"}

    return &(vtpDatabaseEntry.EntityData)
}

// CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode represents unknown database type.
type CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode string

const (
    CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode_client CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode = "client"

    CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode_server CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode = "server"

    CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode_transparent CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode = "transparent"

    CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode_off CISCOVTPMIB_VtpDatabaseTable_VtpDatabaseEntry_VtpDatabaseLocalMode = "off"
)

// CISCOVTPMIB_VtpAuthenticationTable
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
type CISCOVTPMIB_VtpAuthenticationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the status of the VTP authentication information in one
    // domain. The type is slice of
    // CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry.
    VtpAuthEntry []*CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry
}

func (vtpAuthenticationTable *CISCOVTPMIB_VtpAuthenticationTable) GetEntityData() *types.CommonEntityData {
    vtpAuthenticationTable.EntityData.YFilter = vtpAuthenticationTable.YFilter
    vtpAuthenticationTable.EntityData.YangName = "vtpAuthenticationTable"
    vtpAuthenticationTable.EntityData.BundleName = "cisco_ios_xe"
    vtpAuthenticationTable.EntityData.ParentYangName = "CISCO-VTP-MIB"
    vtpAuthenticationTable.EntityData.SegmentPath = "vtpAuthenticationTable"
    vtpAuthenticationTable.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/" + vtpAuthenticationTable.EntityData.SegmentPath
    vtpAuthenticationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpAuthenticationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpAuthenticationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpAuthenticationTable.EntityData.Children = types.NewOrderedMap()
    vtpAuthenticationTable.EntityData.Children.Append("vtpAuthEntry", types.YChild{"VtpAuthEntry", nil})
    for i := range vtpAuthenticationTable.VtpAuthEntry {
        vtpAuthenticationTable.EntityData.Children.Append(types.GetSegmentPath(vtpAuthenticationTable.VtpAuthEntry[i]), types.YChild{"VtpAuthEntry", vtpAuthenticationTable.VtpAuthEntry[i]})
    }
    vtpAuthenticationTable.EntityData.Leafs = types.NewOrderedMap()

    vtpAuthenticationTable.EntityData.YListKeys = []string {}

    return &(vtpAuthenticationTable.EntityData)
}

// CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry
// Information about the status of the VTP
// authentication information in one domain.
type CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..255. Refers to
    // cisco_vtp_mib.CISCOVTPMIB_ManagementDomainTable_ManagementDomainEntry_ManagementDomainIndex
    ManagementDomainIndex interface{}

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
    VtpAuthPassword interface{}

    // By default this object has the value as plaintext(1) and the VTP password
    // is stored in the configuration file in plain text.  Setting this object to
    // hidden(2) will hide the password from the configuration.  Once this object
    // is set to hidden(2), it cannot be set to plaintext(1) alone. However, it
    // may be set to plaintext(1) at the same time the password is set. The type
    // is VtpAuthPasswordType.
    VtpAuthPasswordType interface{}

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
    VtpAuthSecretKey interface{}
}

func (vtpAuthEntry *CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry) GetEntityData() *types.CommonEntityData {
    vtpAuthEntry.EntityData.YFilter = vtpAuthEntry.YFilter
    vtpAuthEntry.EntityData.YangName = "vtpAuthEntry"
    vtpAuthEntry.EntityData.BundleName = "cisco_ios_xe"
    vtpAuthEntry.EntityData.ParentYangName = "vtpAuthenticationTable"
    vtpAuthEntry.EntityData.SegmentPath = "vtpAuthEntry" + types.AddKeyToken(vtpAuthEntry.ManagementDomainIndex, "managementDomainIndex")
    vtpAuthEntry.EntityData.AbsolutePath = "CISCO-VTP-MIB:CISCO-VTP-MIB/vtpAuthenticationTable/" + vtpAuthEntry.EntityData.SegmentPath
    vtpAuthEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vtpAuthEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vtpAuthEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vtpAuthEntry.EntityData.Children = types.NewOrderedMap()
    vtpAuthEntry.EntityData.Leafs = types.NewOrderedMap()
    vtpAuthEntry.EntityData.Leafs.Append("managementDomainIndex", types.YLeaf{"ManagementDomainIndex", vtpAuthEntry.ManagementDomainIndex})
    vtpAuthEntry.EntityData.Leafs.Append("vtpAuthPassword", types.YLeaf{"VtpAuthPassword", vtpAuthEntry.VtpAuthPassword})
    vtpAuthEntry.EntityData.Leafs.Append("vtpAuthPasswordType", types.YLeaf{"VtpAuthPasswordType", vtpAuthEntry.VtpAuthPasswordType})
    vtpAuthEntry.EntityData.Leafs.Append("vtpAuthSecretKey", types.YLeaf{"VtpAuthSecretKey", vtpAuthEntry.VtpAuthSecretKey})

    vtpAuthEntry.EntityData.YListKeys = []string {"ManagementDomainIndex"}

    return &(vtpAuthEntry.EntityData)
}

// CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType represents password is set.
type CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType string

const (
    CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType_plaintext CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType = "plaintext"

    CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType_hidden CISCOVTPMIB_VtpAuthenticationTable_VtpAuthEntry_VtpAuthPasswordType = "hidden"
)

