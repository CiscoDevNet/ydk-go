// The MIB module for the management of the VLAN
// Membership within the frame  work of Cisco
// VLAN Architecture, v 2.0 by Keith McCloghrie. The MIB
// provides information on VLAN Membership Policy Servers
// used by a device and VLAN membership assignments of
// non-trunk bridge ports of the device.
package cisco_vlan_membership_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_vlan_membership_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VLAN-MEMBERSHIP-MIB CISCO-VLAN-MEMBERSHIP-MIB}", reflect.TypeOf(CISCOVLANMEMBERSHIPMIB{}))
    ydk.RegisterEntity("CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB", reflect.TypeOf(CISCOVLANMEMBERSHIPMIB{}))
}

// CISCOVLANMEMBERSHIPMIB
type CISCOVLANMEMBERSHIPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    VmVmps CISCOVLANMEMBERSHIPMIB_VmVmps

    
    VmMembership CISCOVLANMEMBERSHIPMIB_VmMembership

    
    VmStatistics CISCOVLANMEMBERSHIPMIB_VmStatistics

    
    VmStatus CISCOVLANMEMBERSHIPMIB_VmStatus

    // A table of VMPS to use. The device will use the the primary VMPS by
    // default. If the device is unable to reach the primary server after
    // vmVmpsRetries retries, it uses the first secondary server in the table
    // until it runs out of secondary servers, in which case it will return to
    // using the primary server. Entries in this table may be created and deleted
    // via this MIB or the management console on a device.
    VmVmpsTable CISCOVLANMEMBERSHIPMIB_VmVmpsTable

    // A summary of VLAN membership of non-trunk bridge ports. This is a
    // convenience table for retrieving VLAN membership information.  A row is
    // created for a VLAN if: a) the VLAN exists, or b) a port is assigned to a
    // non-existent VLAN.  VLAN membership can only be modified via the
    // vmMembershipTable.
    VmMembershipSummaryTable CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable

    // A table for configuring VLAN port membership. There is one row for each
    // bridge port that is assigned to a static or dynamic access port. Trunk
    // ports are not  represented in this table.  An entry may be created and
    // deleted when ports are created or deleted via SNMP or the management
    // console on a  device.
    VmMembershipTable CISCOVLANMEMBERSHIPMIB_VmMembershipTable

    // A summary of VLAN membership of non-trunk bridge ports. This table is used
    // for  retrieving VLAN membership information for the device which supports
    // dot1dBasePort  with value greater than 2048.  A row is created for a VLAN
    // and a particular bridge port range, where at least one port  in the range
    // is assigned to this VLAN.  VLAN membership can only be modified via the
    // vmMembershipTable.
    VmMembershipSummaryExtTable CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable

    // A table for configuring the Voice VLAN-ID for the ports. An entry will
    // exist for each interface which supports Voice Vlan feature.
    VmVoiceVlanTable CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetEntityData() *types.CommonEntityData {
    cISCOVLANMEMBERSHIPMIB.EntityData.YFilter = cISCOVLANMEMBERSHIPMIB.YFilter
    cISCOVLANMEMBERSHIPMIB.EntityData.YangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVLANMEMBERSHIPMIB.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.SegmentPath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.AbsolutePath = cISCOVLANMEMBERSHIPMIB.EntityData.SegmentPath
    cISCOVLANMEMBERSHIPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVLANMEMBERSHIPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVLANMEMBERSHIPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVLANMEMBERSHIPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmVmps", types.YChild{"VmVmps", &cISCOVLANMEMBERSHIPMIB.VmVmps})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmMembership", types.YChild{"VmMembership", &cISCOVLANMEMBERSHIPMIB.VmMembership})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmStatistics", types.YChild{"VmStatistics", &cISCOVLANMEMBERSHIPMIB.VmStatistics})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmStatus", types.YChild{"VmStatus", &cISCOVLANMEMBERSHIPMIB.VmStatus})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmVmpsTable", types.YChild{"VmVmpsTable", &cISCOVLANMEMBERSHIPMIB.VmVmpsTable})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmMembershipSummaryTable", types.YChild{"VmMembershipSummaryTable", &cISCOVLANMEMBERSHIPMIB.VmMembershipSummaryTable})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmMembershipTable", types.YChild{"VmMembershipTable", &cISCOVLANMEMBERSHIPMIB.VmMembershipTable})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmMembershipSummaryExtTable", types.YChild{"VmMembershipSummaryExtTable", &cISCOVLANMEMBERSHIPMIB.VmMembershipSummaryExtTable})
    cISCOVLANMEMBERSHIPMIB.EntityData.Children.Append("vmVoiceVlanTable", types.YChild{"VmVoiceVlanTable", &cISCOVLANMEMBERSHIPMIB.VmVoiceVlanTable})
    cISCOVLANMEMBERSHIPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOVLANMEMBERSHIPMIB.EntityData.YListKeys = []string {}

    return &(cISCOVLANMEMBERSHIPMIB.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVmps
type CISCOVLANMEMBERSHIPMIB_VmVmps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VLAN Query Protocol (VQP) version supported on the device. VQP is the
    // protocol used to query VLAN Membership Policy Server (VMPS) for VLAN
    // membership assignments of dynamic VLAN ports. A VMPS provides VLAN
    // membership policy assignments based on the content of the packets received
    // on a port. The type is interface{} with range: -2147483648..2147483647.
    VmVmpsVQPVersion interface{}

    // The number of retries for VQP requests to a VMPS before using the next
    // available VMPS. The type is interface{} with range: 1..10.
    VmVmpsRetries interface{}

    // The switch will reconfirm membership of addresses on each port with VMPS
    // periodically. This object specifies the interval to perform reconfirmation.
    // If the value is set to 0, the switch does not reconfirm membership with
    // VMPS. The type is interface{} with range: 0..120. Units are Minutes.
    VmVmpsReconfirmInterval interface{}

    // Setting this object to execute(2) causes the switch to reconfirm membership
    // of every dynamic port. Reading this object always return ready(1). The type
    // is VmVmpsReconfirm.
    VmVmpsReconfirm interface{}

    // This object returns the result of the last request that sets
    // vmVmpsReconfirm to execute(2). The semantics of the possible results are as
    // follows:       other(1)           - none of following      inProgress(2)   
    // - reconfirm in progress      success(3)         - reconfirm completed
    // successfully      noResponse(4)      - reconfirm failed because no         
    // VMPS responded      noVmps(5)          - No VMPS configured     
    // noDynamicPort(6)   - No dynamic ports configured      noHostConnected(7) -
    // No hosts on dynamic ports. The type is VmVmpsReconfirmResult.
    VmVmpsReconfirmResult interface{}

    // This is the IpAddress of the current VMPS used. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    VmVmpsCurrent interface{}
}

func (vmVmps *CISCOVLANMEMBERSHIPMIB_VmVmps) GetEntityData() *types.CommonEntityData {
    vmVmps.EntityData.YFilter = vmVmps.YFilter
    vmVmps.EntityData.YangName = "vmVmps"
    vmVmps.EntityData.BundleName = "cisco_ios_xe"
    vmVmps.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmVmps.EntityData.SegmentPath = "vmVmps"
    vmVmps.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmVmps.EntityData.SegmentPath
    vmVmps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmVmps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmVmps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmVmps.EntityData.Children = types.NewOrderedMap()
    vmVmps.EntityData.Leafs = types.NewOrderedMap()
    vmVmps.EntityData.Leafs.Append("vmVmpsVQPVersion", types.YLeaf{"VmVmpsVQPVersion", vmVmps.VmVmpsVQPVersion})
    vmVmps.EntityData.Leafs.Append("vmVmpsRetries", types.YLeaf{"VmVmpsRetries", vmVmps.VmVmpsRetries})
    vmVmps.EntityData.Leafs.Append("vmVmpsReconfirmInterval", types.YLeaf{"VmVmpsReconfirmInterval", vmVmps.VmVmpsReconfirmInterval})
    vmVmps.EntityData.Leafs.Append("vmVmpsReconfirm", types.YLeaf{"VmVmpsReconfirm", vmVmps.VmVmpsReconfirm})
    vmVmps.EntityData.Leafs.Append("vmVmpsReconfirmResult", types.YLeaf{"VmVmpsReconfirmResult", vmVmps.VmVmpsReconfirmResult})
    vmVmps.EntityData.Leafs.Append("vmVmpsCurrent", types.YLeaf{"VmVmpsCurrent", vmVmps.VmVmpsCurrent})

    vmVmps.EntityData.YListKeys = []string {}

    return &(vmVmps.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm represents Reading this object always return ready(1).
type CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm string

const (
    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm_ready CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm = "ready"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm_execute CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirm = "execute"
)

// CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult represents      noHostConnected(7) - No hosts on dynamic ports
type CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult string

const (
    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_other CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "other"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_inProgress CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "inProgress"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_success CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "success"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_noResponse CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "noResponse"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_noVmps CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "noVmps"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_noDynamicPort CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "noDynamicPort"

    CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult_noHostConnected CISCOVLANMEMBERSHIPMIB_VmVmps_VmVmpsReconfirmResult = "noHostConnected"
)

// CISCOVLANMEMBERSHIPMIB_VmMembership
type CISCOVLANMEMBERSHIPMIB_VmMembership struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object is used to determine whether or not a non-existing VLAN will be
    // created automatically by the system after assigned to a port. 
    // automatic(1):  a non-existing VLAN will be created               
    // automatically by the system after                assigned to a port. 
    // manual(2):     a non-existing VLAN will not be created               
    // automatically by the system and need to be                manually created
    // by the users after assigned                to a port. The type is
    // VmVlanCreationMode.
    VmVlanCreationMode interface{}
}

func (vmMembership *CISCOVLANMEMBERSHIPMIB_VmMembership) GetEntityData() *types.CommonEntityData {
    vmMembership.EntityData.YFilter = vmMembership.YFilter
    vmMembership.EntityData.YangName = "vmMembership"
    vmMembership.EntityData.BundleName = "cisco_ios_xe"
    vmMembership.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmMembership.EntityData.SegmentPath = "vmMembership"
    vmMembership.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmMembership.EntityData.SegmentPath
    vmMembership.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembership.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembership.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembership.EntityData.Children = types.NewOrderedMap()
    vmMembership.EntityData.Leafs = types.NewOrderedMap()
    vmMembership.EntityData.Leafs.Append("vmVlanCreationMode", types.YLeaf{"VmVlanCreationMode", vmMembership.VmVlanCreationMode})

    vmMembership.EntityData.YListKeys = []string {}

    return &(vmMembership.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode represents                to a port.
type CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode string

const (
    CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode_automatic CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode = "automatic"

    CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode_manual CISCOVLANMEMBERSHIPMIB_VmMembership_VmVlanCreationMode = "manual"
)

// CISCOVLANMEMBERSHIPMIB_VmStatistics
type CISCOVLANMEMBERSHIPMIB_VmStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of VQP requests sent by this device to all VMPS since last
    // system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    VmVQPQueries interface{}

    // The number of VQP responses received by this device from all VMPS since
    // last system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    VmVQPResponses interface{}

    // The number of times, since last system re-initialization, the current VMPS
    // was changed. The current VMPS is changed whenever the VMPS fails to 
    // response after vmVmpsRetries of a VQP request. The type is interface{} with
    // range: 0..4294967295.
    VmVmpsChanges interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  'shutdown'. A 'shutdown' response is a result of  the membership
    // policy configured at a VMPS by the administrator. The type is interface{}
    // with range: 0..4294967295.
    VmVQPShutdown interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  'denied'. A 'denied' response is a result of  the membership
    // policy configured at a VMPS by the administrator. The type is interface{}
    // with range: 0..4294967295.
    VmVQPDenied interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates wrong  management domain. A wrong management domain  response
    // indicates that the VMPS used serves a  management domain that is different
    // from the device's management domain. The type is interface{} with range:
    // 0..4294967295.
    VmVQPWrongDomain interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates wrong  VQP version. A wrong VQP version response  indicates that
    // the VMPS used supports a VQP  version that is different from the device's 
    // VQP version. The type is interface{} with range: 0..4294967295.
    VmVQPWrongVersion interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  insufficient resources. An insufficient resources  response
    // indicates that the VMPS used does not  have the required resources to
    // verify the membership assignment requested. The type is interface{} with
    // range: 0..4294967295.
    VmInsufficientResources interface{}
}

func (vmStatistics *CISCOVLANMEMBERSHIPMIB_VmStatistics) GetEntityData() *types.CommonEntityData {
    vmStatistics.EntityData.YFilter = vmStatistics.YFilter
    vmStatistics.EntityData.YangName = "vmStatistics"
    vmStatistics.EntityData.BundleName = "cisco_ios_xe"
    vmStatistics.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmStatistics.EntityData.SegmentPath = "vmStatistics"
    vmStatistics.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmStatistics.EntityData.SegmentPath
    vmStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmStatistics.EntityData.Children = types.NewOrderedMap()
    vmStatistics.EntityData.Leafs = types.NewOrderedMap()
    vmStatistics.EntityData.Leafs.Append("vmVQPQueries", types.YLeaf{"VmVQPQueries", vmStatistics.VmVQPQueries})
    vmStatistics.EntityData.Leafs.Append("vmVQPResponses", types.YLeaf{"VmVQPResponses", vmStatistics.VmVQPResponses})
    vmStatistics.EntityData.Leafs.Append("vmVmpsChanges", types.YLeaf{"VmVmpsChanges", vmStatistics.VmVmpsChanges})
    vmStatistics.EntityData.Leafs.Append("vmVQPShutdown", types.YLeaf{"VmVQPShutdown", vmStatistics.VmVQPShutdown})
    vmStatistics.EntityData.Leafs.Append("vmVQPDenied", types.YLeaf{"VmVQPDenied", vmStatistics.VmVQPDenied})
    vmStatistics.EntityData.Leafs.Append("vmVQPWrongDomain", types.YLeaf{"VmVQPWrongDomain", vmStatistics.VmVQPWrongDomain})
    vmStatistics.EntityData.Leafs.Append("vmVQPWrongVersion", types.YLeaf{"VmVQPWrongVersion", vmStatistics.VmVQPWrongVersion})
    vmStatistics.EntityData.Leafs.Append("vmInsufficientResources", types.YLeaf{"VmInsufficientResources", vmStatistics.VmInsufficientResources})

    vmStatistics.EntityData.YListKeys = []string {}

    return &(vmStatistics.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmStatus
type CISCOVLANMEMBERSHIPMIB_VmStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the notifications/traps defined in this MIB are
    // enabled. The type is bool.
    VmNotificationsEnabled interface{}
}

func (vmStatus *CISCOVLANMEMBERSHIPMIB_VmStatus) GetEntityData() *types.CommonEntityData {
    vmStatus.EntityData.YFilter = vmStatus.YFilter
    vmStatus.EntityData.YangName = "vmStatus"
    vmStatus.EntityData.BundleName = "cisco_ios_xe"
    vmStatus.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmStatus.EntityData.SegmentPath = "vmStatus"
    vmStatus.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmStatus.EntityData.SegmentPath
    vmStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmStatus.EntityData.Children = types.NewOrderedMap()
    vmStatus.EntityData.Leafs = types.NewOrderedMap()
    vmStatus.EntityData.Leafs.Append("vmNotificationsEnabled", types.YLeaf{"VmNotificationsEnabled", vmStatus.VmNotificationsEnabled})

    vmStatus.EntityData.YListKeys = []string {}

    return &(vmStatus.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVmpsTable
// A table of VMPS to use. The device will use
// the the primary VMPS by default. If the
// device is unable to reach the primary server
// after vmVmpsRetries retries, it uses the first
// secondary server in the table until it runs out
// of secondary servers, in which case it will return
// to using the primary server. Entries in this table
// may be created and deleted via this MIB or
// the management console on a device.
type CISCOVLANMEMBERSHIPMIB_VmVmpsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVmpsTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_VmVmpsTable_VmVmpsEntry.
    VmVmpsEntry []*CISCOVLANMEMBERSHIPMIB_VmVmpsTable_VmVmpsEntry
}

func (vmVmpsTable *CISCOVLANMEMBERSHIPMIB_VmVmpsTable) GetEntityData() *types.CommonEntityData {
    vmVmpsTable.EntityData.YFilter = vmVmpsTable.YFilter
    vmVmpsTable.EntityData.YangName = "vmVmpsTable"
    vmVmpsTable.EntityData.BundleName = "cisco_ios_xe"
    vmVmpsTable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmVmpsTable.EntityData.SegmentPath = "vmVmpsTable"
    vmVmpsTable.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmVmpsTable.EntityData.SegmentPath
    vmVmpsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmVmpsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmVmpsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmVmpsTable.EntityData.Children = types.NewOrderedMap()
    vmVmpsTable.EntityData.Children.Append("vmVmpsEntry", types.YChild{"VmVmpsEntry", nil})
    for i := range vmVmpsTable.VmVmpsEntry {
        vmVmpsTable.EntityData.Children.Append(types.GetSegmentPath(vmVmpsTable.VmVmpsEntry[i]), types.YChild{"VmVmpsEntry", vmVmpsTable.VmVmpsEntry[i]})
    }
    vmVmpsTable.EntityData.Leafs = types.NewOrderedMap()

    vmVmpsTable.EntityData.YListKeys = []string {}

    return &(vmVmpsTable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVmpsTable_VmVmpsEntry
// An entry (conceptual row) in the vmVmpsTable.
type CISCOVLANMEMBERSHIPMIB_VmVmpsTable_VmVmpsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Ip Address of the VMPS. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    VmVmpsIpAddress interface{}

    // The status of the VMPS. Setting this value to true will make this VMPS the
    // primary server and make the switch use this as the current server. Setting
    // this entry to true causes other rows to transition to false. Attempting to
    // write a value of false after creation will result in a return of bad value.
    // Deleting an entry whose value is true will result in the first entry in the
    // table being set to true. The type is bool.
    VmVmpsPrimary interface{}

    // The status of this conceptual row. The type is RowStatus.
    VmVmpsRowStatus interface{}
}

func (vmVmpsEntry *CISCOVLANMEMBERSHIPMIB_VmVmpsTable_VmVmpsEntry) GetEntityData() *types.CommonEntityData {
    vmVmpsEntry.EntityData.YFilter = vmVmpsEntry.YFilter
    vmVmpsEntry.EntityData.YangName = "vmVmpsEntry"
    vmVmpsEntry.EntityData.BundleName = "cisco_ios_xe"
    vmVmpsEntry.EntityData.ParentYangName = "vmVmpsTable"
    vmVmpsEntry.EntityData.SegmentPath = "vmVmpsEntry" + types.AddKeyToken(vmVmpsEntry.VmVmpsIpAddress, "vmVmpsIpAddress")
    vmVmpsEntry.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/vmVmpsTable/" + vmVmpsEntry.EntityData.SegmentPath
    vmVmpsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmVmpsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmVmpsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmVmpsEntry.EntityData.Children = types.NewOrderedMap()
    vmVmpsEntry.EntityData.Leafs = types.NewOrderedMap()
    vmVmpsEntry.EntityData.Leafs.Append("vmVmpsIpAddress", types.YLeaf{"VmVmpsIpAddress", vmVmpsEntry.VmVmpsIpAddress})
    vmVmpsEntry.EntityData.Leafs.Append("vmVmpsPrimary", types.YLeaf{"VmVmpsPrimary", vmVmpsEntry.VmVmpsPrimary})
    vmVmpsEntry.EntityData.Leafs.Append("vmVmpsRowStatus", types.YLeaf{"VmVmpsRowStatus", vmVmpsEntry.VmVmpsRowStatus})

    vmVmpsEntry.EntityData.YListKeys = []string {"VmVmpsIpAddress"}

    return &(vmVmpsEntry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable
// A summary of VLAN membership of non-trunk
// bridge ports. This is a convenience table
// for retrieving VLAN membership information.
// 
// A row is created for a VLAN if:
// a) the VLAN exists, or
// b) a port is assigned to a non-existent VLAN.
// 
// VLAN membership can only be modified via the
// vmMembershipTable.
type CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry.
    VmMembershipSummaryEntry []*CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry
}

func (vmMembershipSummaryTable *CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable) GetEntityData() *types.CommonEntityData {
    vmMembershipSummaryTable.EntityData.YFilter = vmMembershipSummaryTable.YFilter
    vmMembershipSummaryTable.EntityData.YangName = "vmMembershipSummaryTable"
    vmMembershipSummaryTable.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipSummaryTable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmMembershipSummaryTable.EntityData.SegmentPath = "vmMembershipSummaryTable"
    vmMembershipSummaryTable.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmMembershipSummaryTable.EntityData.SegmentPath
    vmMembershipSummaryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipSummaryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipSummaryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipSummaryTable.EntityData.Children = types.NewOrderedMap()
    vmMembershipSummaryTable.EntityData.Children.Append("vmMembershipSummaryEntry", types.YChild{"VmMembershipSummaryEntry", nil})
    for i := range vmMembershipSummaryTable.VmMembershipSummaryEntry {
        vmMembershipSummaryTable.EntityData.Children.Append(types.GetSegmentPath(vmMembershipSummaryTable.VmMembershipSummaryEntry[i]), types.YChild{"VmMembershipSummaryEntry", vmMembershipSummaryTable.VmMembershipSummaryEntry[i]})
    }
    vmMembershipSummaryTable.EntityData.Leafs = types.NewOrderedMap()

    vmMembershipSummaryTable.EntityData.YListKeys = []string {}

    return &(vmMembershipSummaryTable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry
// An entry (conceptual row) in the
// vmMembershipSummaryTable.
type CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    VmMembershipSummaryVlanIndex interface{}

    // The set of the device's member ports that belong to the VLAN.  Each octet
    // within the value of this object specifies a set of eight ports, with the
    // first octet specifying ports 1 through 8, the second octet specifying ports
    // 9 through 16, etc.   Within each octet, the most significant bit represents
    // the lowest numbered port, and the least significant bit represents the
    // highest numbered port.  Thus, each port of the VLAN is represented by a
    // single bit within the value of this object.  If that bit has a value of '1'
    // then that port is included in the set of ports; the port is not included if
    // its bit has a value of '0'.  A port number is the value of dot1dBasePort
    // for the port in the BRIDGE-MIB (RFC 1493). The type is string with length:
    // 0..128.
    VmMembershipSummaryMemberPorts interface{}

    // The set of the device's member ports that belong to the VLAN. It has the
    // VLAN membership information of up to 2048 ports with the port number from 1
    // to  2048.  Each octet within the value of this object specifies a set of
    // eight ports, with the first octet specifying  ports 1 through 8, the second
    // octet specifying ports 9 through 16, etc.   Within each octet, the most
    // significant bit represents the lowest numbered port, and the least
    // significant bit represents the highest numbered port.  Thus, each port of
    // the VLAN is represented by a single bit within the value of this object. 
    // If that bit has a value of '1' then that port is included in the set of
    // ports; the port is not included if its bit has a value of '0'.  A port
    // number is the value of dot1dBasePort for the port in the BRIDGE-MIB (RFC
    // 1493). The type is string with length: 0..256.
    VmMembershipSummaryMember2kPorts interface{}
}

func (vmMembershipSummaryEntry *CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry) GetEntityData() *types.CommonEntityData {
    vmMembershipSummaryEntry.EntityData.YFilter = vmMembershipSummaryEntry.YFilter
    vmMembershipSummaryEntry.EntityData.YangName = "vmMembershipSummaryEntry"
    vmMembershipSummaryEntry.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipSummaryEntry.EntityData.ParentYangName = "vmMembershipSummaryTable"
    vmMembershipSummaryEntry.EntityData.SegmentPath = "vmMembershipSummaryEntry" + types.AddKeyToken(vmMembershipSummaryEntry.VmMembershipSummaryVlanIndex, "vmMembershipSummaryVlanIndex")
    vmMembershipSummaryEntry.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/vmMembershipSummaryTable/" + vmMembershipSummaryEntry.EntityData.SegmentPath
    vmMembershipSummaryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipSummaryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipSummaryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipSummaryEntry.EntityData.Children = types.NewOrderedMap()
    vmMembershipSummaryEntry.EntityData.Leafs = types.NewOrderedMap()
    vmMembershipSummaryEntry.EntityData.Leafs.Append("vmMembershipSummaryVlanIndex", types.YLeaf{"VmMembershipSummaryVlanIndex", vmMembershipSummaryEntry.VmMembershipSummaryVlanIndex})
    vmMembershipSummaryEntry.EntityData.Leafs.Append("vmMembershipSummaryMemberPorts", types.YLeaf{"VmMembershipSummaryMemberPorts", vmMembershipSummaryEntry.VmMembershipSummaryMemberPorts})
    vmMembershipSummaryEntry.EntityData.Leafs.Append("vmMembershipSummaryMember2kPorts", types.YLeaf{"VmMembershipSummaryMember2kPorts", vmMembershipSummaryEntry.VmMembershipSummaryMember2kPorts})

    vmMembershipSummaryEntry.EntityData.YListKeys = []string {"VmMembershipSummaryVlanIndex"}

    return &(vmMembershipSummaryEntry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipTable
// A table for configuring VLAN port membership.
// There is one row for each bridge port that is
// assigned to a static or dynamic access port. Trunk
// ports are not  represented in this table.  An entry
// may be created and deleted when ports are created or
// deleted via SNMP or the management console on a 
// device.
type CISCOVLANMEMBERSHIPMIB_VmMembershipTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry.
    VmMembershipEntry []*CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry
}

func (vmMembershipTable *CISCOVLANMEMBERSHIPMIB_VmMembershipTable) GetEntityData() *types.CommonEntityData {
    vmMembershipTable.EntityData.YFilter = vmMembershipTable.YFilter
    vmMembershipTable.EntityData.YangName = "vmMembershipTable"
    vmMembershipTable.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipTable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmMembershipTable.EntityData.SegmentPath = "vmMembershipTable"
    vmMembershipTable.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmMembershipTable.EntityData.SegmentPath
    vmMembershipTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipTable.EntityData.Children = types.NewOrderedMap()
    vmMembershipTable.EntityData.Children.Append("vmMembershipEntry", types.YChild{"VmMembershipEntry", nil})
    for i := range vmMembershipTable.VmMembershipEntry {
        vmMembershipTable.EntityData.Children.Append(types.GetSegmentPath(vmMembershipTable.VmMembershipEntry[i]), types.YChild{"VmMembershipEntry", vmMembershipTable.VmMembershipEntry[i]})
    }
    vmMembershipTable.EntityData.Leafs = types.NewOrderedMap()

    vmMembershipTable.EntityData.YListKeys = []string {}

    return &(vmMembershipTable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry
// An entry (conceptual row) in the vmMembershipTable.
type CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The type of VLAN membership assigned to this port. A port with static vlan
    // membership is assigned to a single VLAN directly. A port with dynamic
    // membership is assigned a single VLAN based on content of packets received
    // on the port and via VQP queries to VMPS. A port with multiVlan membership
    // may be assigned to one or more VLANs directly.  A static or dynamic port
    // membership is specified by the value of vmVlan. A multiVlan port membership
    // is specified by the value of vmVlans. The type is VmVlanType.
    VmVlanType interface{}

    // The VLAN id of the VLAN the port is assigned to when vmVlanType is set to
    // static or dynamic. This object is not instantiated if not applicable.  The
    // value may be 0 if the port is not assigned to a VLAN.  If vmVlanType is
    // static, the port is always assigned to a VLAN and the object may not be set
    // to 0.  If vmVlanType is dynamic the object's value is 0 if the port is
    // currently not assigned to a VLAN. In addition, the object may be set to 0
    // only. The type is interface{} with range: 0..4095.
    VmVlan interface{}

    // An indication of the current VLAN status of the port. A status of
    // inactive(1) indicates that a dynamic port does not yet have a VLAN
    // assigned, or a port is  assigned to a VLAN that is currently not active. A 
    // status of active(2) indicates that the currently  assigned VLAN is active.
    // A status of shutdown(3)  indicates that the port has been disabled as a
    // result of VQP shutdown response. The type is VmPortStatus.
    VmPortStatus interface{}

    // The VLAN(s) the port is assigned to when the port's vmVlanType is set to
    // multiVlan. This object is not instantiated if not applicable.  The port is
    // always assigned to one or more VLANs and the object may not be set so that
    // there are no vlans assigned.  Each octet within the value of this object
    // specifies a set of eight VLANs, with the first octet specifying VLAN id 1
    // through 8, the second octet specifying VLAN ids 9 through 16, etc.   Within
    // each octet, the most significant bit represents the lowest numbered VLAN
    // id, and the least significant bit represents the highest numbered VLAN id. 
    // Thus, each VLAN of the port is represented by a single bit within the value
    // of this object.  If that bit has a value of '1' then that VLAN is included
    // in the set of VLANs; the VLAN is not included if its bit has a value of
    // '0'. The type is string with length: 0..128.
    VmVlans interface{}

    // The VLAN(s) the port is assigned to when the port's vmVlanType is set to
    // multiVlan. This object is not instantiated if not applicable.  The port is
    // always assigned to one or more VLANs and the object may not be set so that
    // there are no vlans assigned.  Each octet within the value of this object
    // specifies a set of eight VLANs, with the first octet specifying VLAN id
    // 1024 through 1031, the second octet specifying  VLAN ids 1032 through 1039,
    // etc.  Within each octet,  the most significant bit represents the lowest 
    // numbered VLAN id, and the least significant bit  represents the highest
    // numbered VLAN id.  Thus, each  VLAN of the port is represented by a single
    // bit within the value of this object.  If that bit has a value of '1' then
    // that VLAN is included in the set of VLANs; the VLAN is not included if its
    // bit has a value of '0'. The type is string with length: 0..128.
    VmVlans2k interface{}

    // The VLAN(s) the port is assigned to when the port's vmVlanType is set to
    // multiVlan. This object is not instantiated if not applicable.  The port is
    // always assigned to one or more VLANs and the object may not be set so that
    // there are no vlans assigned.  Each octet within the value of this object
    // specifies a set of eight VLANs, with the first octet specifying VLAN id
    // 2048 through 2055, the second octet specifying  VLAN ids 2056 through 2063,
    // etc.   Within each octet,  the most significant bit represents the lowest 
    // numbered VLAN id, and the least significant bit  represents the highest
    // numbered VLAN id.  Thus, each VLAN of the port is represented by a single
    // bit within the value of this object.  If that bit has a value of '1' then
    // that VLAN is included in the set of VLANs; the VLAN is not included if its
    // bit has a value of '0'. The type is string with length: 0..128.
    VmVlans3k interface{}

    // The VLAN(s) the port is assigned to when the port's vmVlanType is set to
    // multiVlan. This object is not instantiated if not applicable.  The port is
    // always assigned to one or more VLANs and the object may not be set so that
    // there are no vlans assigned.  Each octet within the value of this object
    // specifies a set of eight VLANs, with the first octet specifying VLAN id
    // 3072 through 3079, the second octet specifying  VLAN ids 3040 through 3047,
    // etc.   Within each octet,  the most significant bit represents the lowest 
    // numbered VLAN id, and the least significant bit  represents the highest
    // numbered VLAN id.  Thus, each VLAN of the port is represented by a single
    // bit within the value of this object.  If that bit has a value of '1' then
    // that VLAN is included in the set of VLANs; the VLAN is not included if its
    // bit has a value of '0'. The type is string with length: 0..128.
    VmVlans4k interface{}
}

func (vmMembershipEntry *CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry) GetEntityData() *types.CommonEntityData {
    vmMembershipEntry.EntityData.YFilter = vmMembershipEntry.YFilter
    vmMembershipEntry.EntityData.YangName = "vmMembershipEntry"
    vmMembershipEntry.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipEntry.EntityData.ParentYangName = "vmMembershipTable"
    vmMembershipEntry.EntityData.SegmentPath = "vmMembershipEntry" + types.AddKeyToken(vmMembershipEntry.IfIndex, "ifIndex")
    vmMembershipEntry.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/vmMembershipTable/" + vmMembershipEntry.EntityData.SegmentPath
    vmMembershipEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipEntry.EntityData.Children = types.NewOrderedMap()
    vmMembershipEntry.EntityData.Leafs = types.NewOrderedMap()
    vmMembershipEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", vmMembershipEntry.IfIndex})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlanType", types.YLeaf{"VmVlanType", vmMembershipEntry.VmVlanType})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlan", types.YLeaf{"VmVlan", vmMembershipEntry.VmVlan})
    vmMembershipEntry.EntityData.Leafs.Append("vmPortStatus", types.YLeaf{"VmPortStatus", vmMembershipEntry.VmPortStatus})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlans", types.YLeaf{"VmVlans", vmMembershipEntry.VmVlans})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlans2k", types.YLeaf{"VmVlans2k", vmMembershipEntry.VmVlans2k})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlans3k", types.YLeaf{"VmVlans3k", vmMembershipEntry.VmVlans3k})
    vmMembershipEntry.EntityData.Leafs.Append("vmVlans4k", types.YLeaf{"VmVlans4k", vmMembershipEntry.VmVlans4k})

    vmMembershipEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(vmMembershipEntry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus represents of VQP shutdown response.
type CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus string

const (
    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus_inactive CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus = "inactive"

    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus_active CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus = "active"

    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus_shutdown CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmPortStatus = "shutdown"
)

// CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType represents specified by the value of vmVlans.
type CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType string

const (
    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType_static CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType = "static"

    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType_dynamic CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType = "dynamic"

    CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType_multiVlan CISCOVLANMEMBERSHIPMIB_VmMembershipTable_VmMembershipEntry_VmVlanType = "multiVlan"
)

// CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable
// A summary of VLAN membership of non-trunk
// bridge ports. This table is used for 
// retrieving VLAN membership information
// for the device which supports dot1dBasePort 
// with value greater than 2048.
// 
// A row is created for a VLAN and a particular
// bridge port range, where at least one port 
// in the range is assigned to this VLAN.
// 
// VLAN membership can only be modified via the
// vmMembershipTable.
type CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryExtTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable_VmMembershipSummaryExtEntry.
    VmMembershipSummaryExtEntry []*CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable_VmMembershipSummaryExtEntry
}

func (vmMembershipSummaryExtTable *CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable) GetEntityData() *types.CommonEntityData {
    vmMembershipSummaryExtTable.EntityData.YFilter = vmMembershipSummaryExtTable.YFilter
    vmMembershipSummaryExtTable.EntityData.YangName = "vmMembershipSummaryExtTable"
    vmMembershipSummaryExtTable.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipSummaryExtTable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmMembershipSummaryExtTable.EntityData.SegmentPath = "vmMembershipSummaryExtTable"
    vmMembershipSummaryExtTable.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmMembershipSummaryExtTable.EntityData.SegmentPath
    vmMembershipSummaryExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipSummaryExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipSummaryExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipSummaryExtTable.EntityData.Children = types.NewOrderedMap()
    vmMembershipSummaryExtTable.EntityData.Children.Append("vmMembershipSummaryExtEntry", types.YChild{"VmMembershipSummaryExtEntry", nil})
    for i := range vmMembershipSummaryExtTable.VmMembershipSummaryExtEntry {
        vmMembershipSummaryExtTable.EntityData.Children.Append(types.GetSegmentPath(vmMembershipSummaryExtTable.VmMembershipSummaryExtEntry[i]), types.YChild{"VmMembershipSummaryExtEntry", vmMembershipSummaryExtTable.VmMembershipSummaryExtEntry[i]})
    }
    vmMembershipSummaryExtTable.EntityData.Leafs = types.NewOrderedMap()

    vmMembershipSummaryExtTable.EntityData.YListKeys = []string {}

    return &(vmMembershipSummaryExtTable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable_VmMembershipSummaryExtEntry
// An entry (conceptual row) in the
// vmMembershipSummaryExtTable.
type CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable_VmMembershipSummaryExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vlan_membership_mib.CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryTable_VmMembershipSummaryEntry_VmMembershipSummaryVlanIndex
    VmMembershipSummaryVlanIndex interface{}

    // This attribute is a key. The bridge port range index of this row. The type
    // is CiscoPortListRange.
    VmMembershipPortRangeIndex interface{}

    // The set of the device's member ports that belong to the VLAN. It has the
    // VLAN membership information of up to 2k ports with the port number starting
    // from the information indicated in vmMembershipPortRangeIndex object of the
    // same row. For example, if the value of vmMembershipPortRangeIndex is
    // 'twoKto4K', the port number indicated in this object starting from 2049 and
    // ending to 4096.   A port number is the value of dot1dBasePort for the port
    // in the BRIDGE-MIB (RFC 1493). The type is string with length: 0..256.
    VmMembershipSummaryExtPorts interface{}
}

func (vmMembershipSummaryExtEntry *CISCOVLANMEMBERSHIPMIB_VmMembershipSummaryExtTable_VmMembershipSummaryExtEntry) GetEntityData() *types.CommonEntityData {
    vmMembershipSummaryExtEntry.EntityData.YFilter = vmMembershipSummaryExtEntry.YFilter
    vmMembershipSummaryExtEntry.EntityData.YangName = "vmMembershipSummaryExtEntry"
    vmMembershipSummaryExtEntry.EntityData.BundleName = "cisco_ios_xe"
    vmMembershipSummaryExtEntry.EntityData.ParentYangName = "vmMembershipSummaryExtTable"
    vmMembershipSummaryExtEntry.EntityData.SegmentPath = "vmMembershipSummaryExtEntry" + types.AddKeyToken(vmMembershipSummaryExtEntry.VmMembershipSummaryVlanIndex, "vmMembershipSummaryVlanIndex") + types.AddKeyToken(vmMembershipSummaryExtEntry.VmMembershipPortRangeIndex, "vmMembershipPortRangeIndex")
    vmMembershipSummaryExtEntry.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/vmMembershipSummaryExtTable/" + vmMembershipSummaryExtEntry.EntityData.SegmentPath
    vmMembershipSummaryExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmMembershipSummaryExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmMembershipSummaryExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmMembershipSummaryExtEntry.EntityData.Children = types.NewOrderedMap()
    vmMembershipSummaryExtEntry.EntityData.Leafs = types.NewOrderedMap()
    vmMembershipSummaryExtEntry.EntityData.Leafs.Append("vmMembershipSummaryVlanIndex", types.YLeaf{"VmMembershipSummaryVlanIndex", vmMembershipSummaryExtEntry.VmMembershipSummaryVlanIndex})
    vmMembershipSummaryExtEntry.EntityData.Leafs.Append("vmMembershipPortRangeIndex", types.YLeaf{"VmMembershipPortRangeIndex", vmMembershipSummaryExtEntry.VmMembershipPortRangeIndex})
    vmMembershipSummaryExtEntry.EntityData.Leafs.Append("vmMembershipSummaryExtPorts", types.YLeaf{"VmMembershipSummaryExtPorts", vmMembershipSummaryExtEntry.VmMembershipSummaryExtPorts})

    vmMembershipSummaryExtEntry.EntityData.YListKeys = []string {"VmMembershipSummaryVlanIndex", "VmMembershipPortRangeIndex"}

    return &(vmMembershipSummaryExtEntry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable
// A table for configuring the Voice VLAN-ID
// for the ports. An entry will exist for each
// interface which supports Voice Vlan feature.
type CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVoiceVlanTable. Only interfaces which
    // support Voice Vlan feature are shown. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable_VmVoiceVlanEntry.
    VmVoiceVlanEntry []*CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable_VmVoiceVlanEntry
}

func (vmVoiceVlanTable *CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable) GetEntityData() *types.CommonEntityData {
    vmVoiceVlanTable.EntityData.YFilter = vmVoiceVlanTable.YFilter
    vmVoiceVlanTable.EntityData.YangName = "vmVoiceVlanTable"
    vmVoiceVlanTable.EntityData.BundleName = "cisco_ios_xe"
    vmVoiceVlanTable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmVoiceVlanTable.EntityData.SegmentPath = "vmVoiceVlanTable"
    vmVoiceVlanTable.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/" + vmVoiceVlanTable.EntityData.SegmentPath
    vmVoiceVlanTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmVoiceVlanTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmVoiceVlanTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmVoiceVlanTable.EntityData.Children = types.NewOrderedMap()
    vmVoiceVlanTable.EntityData.Children.Append("vmVoiceVlanEntry", types.YChild{"VmVoiceVlanEntry", nil})
    for i := range vmVoiceVlanTable.VmVoiceVlanEntry {
        vmVoiceVlanTable.EntityData.Children.Append(types.GetSegmentPath(vmVoiceVlanTable.VmVoiceVlanEntry[i]), types.YChild{"VmVoiceVlanEntry", vmVoiceVlanTable.VmVoiceVlanEntry[i]})
    }
    vmVoiceVlanTable.EntityData.Leafs = types.NewOrderedMap()

    vmVoiceVlanTable.EntityData.YListKeys = []string {}

    return &(vmVoiceVlanTable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable_VmVoiceVlanEntry
// An entry (conceptual row) in the vmVoiceVlanTable.
// Only interfaces which support Voice Vlan feature
// are shown.
type CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable_VmVoiceVlanEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The Voice Vlan ID (VVID) to which this port belongs to.  0    -    The CDP
    // packets transmitting            through this port would contain          
    // Appliance VLAN-ID TLV with value            of 0. VoIP and related packets 
    // are expected to be sent and            received with VLAN-id=0 and an      
    // 802.1p priority.   1..4094 - The CDP packets transmitting           through
    // this port would contain           Appliance VLAN-ID TLV with N.          
    // VoIP and related packets are           expected to be sent and received    
    // with VLAN-id=N and an 802.1p           priority.  4095  -   The CDP packets
    // transmitting           through this port would contain           Appliance
    // VLAN-ID TLV with value           of 4095. VoIP and related packets         
    // are expected to be sent and            received untagged without an        
    // 802.1p priority.  4096  -   The CDP packets transmitting            through
    // this port would not            include Appliance VLAN-ID TLV;           
    // or, if the VVID is not supported            on the port, this MIB object
    // will           not be configurable and will            return 4096. The
    // type is interface{} with range: 0..4096.
    VmVoiceVlanId interface{}

    // Enable or Disable the feature of CDP message verification of voice VLANs. 
    // true   - The voice VLAN vmVoiceVlan is enabled           only after CDP
    // messages are received           from the IP phone.  false -  The voice VLAN
    // vmVoiceVlan is enabled          as soon as the IP phone interface is       
    // up. There is no verification needed           from CDP messages from the IP
    // phone. The type is bool.
    VmVoiceVlanCdpVerifyEnable interface{}
}

func (vmVoiceVlanEntry *CISCOVLANMEMBERSHIPMIB_VmVoiceVlanTable_VmVoiceVlanEntry) GetEntityData() *types.CommonEntityData {
    vmVoiceVlanEntry.EntityData.YFilter = vmVoiceVlanEntry.YFilter
    vmVoiceVlanEntry.EntityData.YangName = "vmVoiceVlanEntry"
    vmVoiceVlanEntry.EntityData.BundleName = "cisco_ios_xe"
    vmVoiceVlanEntry.EntityData.ParentYangName = "vmVoiceVlanTable"
    vmVoiceVlanEntry.EntityData.SegmentPath = "vmVoiceVlanEntry" + types.AddKeyToken(vmVoiceVlanEntry.IfIndex, "ifIndex")
    vmVoiceVlanEntry.EntityData.AbsolutePath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB/vmVoiceVlanTable/" + vmVoiceVlanEntry.EntityData.SegmentPath
    vmVoiceVlanEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmVoiceVlanEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmVoiceVlanEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmVoiceVlanEntry.EntityData.Children = types.NewOrderedMap()
    vmVoiceVlanEntry.EntityData.Leafs = types.NewOrderedMap()
    vmVoiceVlanEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", vmVoiceVlanEntry.IfIndex})
    vmVoiceVlanEntry.EntityData.Leafs.Append("vmVoiceVlanId", types.YLeaf{"VmVoiceVlanId", vmVoiceVlanEntry.VmVoiceVlanId})
    vmVoiceVlanEntry.EntityData.Leafs.Append("vmVoiceVlanCdpVerifyEnable", types.YLeaf{"VmVoiceVlanCdpVerifyEnable", vmVoiceVlanEntry.VmVoiceVlanCdpVerifyEnable})

    vmVoiceVlanEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(vmVoiceVlanEntry.EntityData)
}

