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

    
    Vmvmps CISCOVLANMEMBERSHIPMIB_Vmvmps

    
    Vmmembership CISCOVLANMEMBERSHIPMIB_Vmmembership

    
    Vmstatistics CISCOVLANMEMBERSHIPMIB_Vmstatistics

    
    Vmstatus CISCOVLANMEMBERSHIPMIB_Vmstatus

    // A table of VMPS to use. The device will use the the primary VMPS by
    // default. If the device is unable to reach the primary server after
    // vmVmpsRetries retries, it uses the first secondary server in the table
    // until it runs out of secondary servers, in which case it will return to
    // using the primary server. Entries in this table may be created and deleted
    // via this MIB or the management console on a device.
    Vmvmpstable CISCOVLANMEMBERSHIPMIB_Vmvmpstable

    // A summary of VLAN membership of non-trunk bridge ports. This is a
    // convenience table for retrieving VLAN membership information.  A row is
    // created for a VLAN if: a) the VLAN exists, or b) a port is assigned to a
    // non-existent VLAN.  VLAN membership can only be modified via the
    // vmMembershipTable.
    Vmmembershipsummarytable CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable

    // A table for configuring VLAN port membership. There is one row for each
    // bridge port that is assigned to a static or dynamic access port. Trunk
    // ports are not  represented in this table.  An entry may be created and
    // deleted when ports are created or deleted via SNMP or the management
    // console on a  device.
    Vmmembershiptable CISCOVLANMEMBERSHIPMIB_Vmmembershiptable

    // A summary of VLAN membership of non-trunk bridge ports. This table is used
    // for  retrieving VLAN membership information for the device which supports
    // dot1dBasePort  with value greater than 2048.  A row is created for a VLAN
    // and a particular bridge port range, where at least one port  in the range
    // is assigned to this VLAN.  VLAN membership can only be modified via the
    // vmMembershipTable.
    Vmmembershipsummaryexttable CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable

    // A table for configuring the Voice VLAN-ID for the ports. An entry will
    // exist for each interface which supports Voice Vlan feature.
    Vmvoicevlantable CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetEntityData() *types.CommonEntityData {
    cISCOVLANMEMBERSHIPMIB.EntityData.YFilter = cISCOVLANMEMBERSHIPMIB.YFilter
    cISCOVLANMEMBERSHIPMIB.EntityData.YangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOVLANMEMBERSHIPMIB.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.SegmentPath = "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB"
    cISCOVLANMEMBERSHIPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOVLANMEMBERSHIPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOVLANMEMBERSHIPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOVLANMEMBERSHIPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmVmps"] = types.YChild{"Vmvmps", &cISCOVLANMEMBERSHIPMIB.Vmvmps}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmMembership"] = types.YChild{"Vmmembership", &cISCOVLANMEMBERSHIPMIB.Vmmembership}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmStatistics"] = types.YChild{"Vmstatistics", &cISCOVLANMEMBERSHIPMIB.Vmstatistics}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmStatus"] = types.YChild{"Vmstatus", &cISCOVLANMEMBERSHIPMIB.Vmstatus}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmVmpsTable"] = types.YChild{"Vmvmpstable", &cISCOVLANMEMBERSHIPMIB.Vmvmpstable}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmMembershipSummaryTable"] = types.YChild{"Vmmembershipsummarytable", &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummarytable}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmMembershipTable"] = types.YChild{"Vmmembershiptable", &cISCOVLANMEMBERSHIPMIB.Vmmembershiptable}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmMembershipSummaryExtTable"] = types.YChild{"Vmmembershipsummaryexttable", &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummaryexttable}
    cISCOVLANMEMBERSHIPMIB.EntityData.Children["vmVoiceVlanTable"] = types.YChild{"Vmvoicevlantable", &cISCOVLANMEMBERSHIPMIB.Vmvoicevlantable}
    cISCOVLANMEMBERSHIPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOVLANMEMBERSHIPMIB.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvmps
type CISCOVLANMEMBERSHIPMIB_Vmvmps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VLAN Query Protocol (VQP) version supported on the device. VQP is the
    // protocol used to query VLAN Membership Policy Server (VMPS) for VLAN
    // membership assignments of dynamic VLAN ports. A VMPS provides VLAN
    // membership policy assignments based on the content of the packets received
    // on a port. The type is interface{} with range: -2147483648..2147483647.
    Vmvmpsvqpversion interface{}

    // The number of retries for VQP requests to a VMPS before using the next
    // available VMPS. The type is interface{} with range: 1..10.
    Vmvmpsretries interface{}

    // The switch will reconfirm membership of addresses on each port with VMPS
    // periodically. This object specifies the interval to perform reconfirmation.
    // If the value is set to 0, the switch does not reconfirm membership with
    // VMPS. The type is interface{} with range: 0..120. Units are Minutes.
    Vmvmpsreconfirminterval interface{}

    // Setting this object to execute(2) causes the switch to reconfirm membership
    // of every dynamic port. Reading this object always return ready(1). The type
    // is Vmvmpsreconfirm.
    Vmvmpsreconfirm interface{}

    // This object returns the result of the last request that sets
    // vmVmpsReconfirm to execute(2). The semantics of the possible results are as
    // follows:       other(1)           - none of following      inProgress(2)   
    // - reconfirm in progress      success(3)         - reconfirm completed
    // successfully      noResponse(4)      - reconfirm failed because no         
    // VMPS responded      noVmps(5)          - No VMPS configured     
    // noDynamicPort(6)   - No dynamic ports configured      noHostConnected(7) -
    // No hosts on dynamic ports. The type is Vmvmpsreconfirmresult.
    Vmvmpsreconfirmresult interface{}

    // This is the IpAddress of the current VMPS used. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Vmvmpscurrent interface{}
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetEntityData() *types.CommonEntityData {
    vmvmps.EntityData.YFilter = vmvmps.YFilter
    vmvmps.EntityData.YangName = "vmVmps"
    vmvmps.EntityData.BundleName = "cisco_ios_xe"
    vmvmps.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmvmps.EntityData.SegmentPath = "vmVmps"
    vmvmps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmvmps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmvmps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmvmps.EntityData.Children = make(map[string]types.YChild)
    vmvmps.EntityData.Leafs = make(map[string]types.YLeaf)
    vmvmps.EntityData.Leafs["vmVmpsVQPVersion"] = types.YLeaf{"Vmvmpsvqpversion", vmvmps.Vmvmpsvqpversion}
    vmvmps.EntityData.Leafs["vmVmpsRetries"] = types.YLeaf{"Vmvmpsretries", vmvmps.Vmvmpsretries}
    vmvmps.EntityData.Leafs["vmVmpsReconfirmInterval"] = types.YLeaf{"Vmvmpsreconfirminterval", vmvmps.Vmvmpsreconfirminterval}
    vmvmps.EntityData.Leafs["vmVmpsReconfirm"] = types.YLeaf{"Vmvmpsreconfirm", vmvmps.Vmvmpsreconfirm}
    vmvmps.EntityData.Leafs["vmVmpsReconfirmResult"] = types.YLeaf{"Vmvmpsreconfirmresult", vmvmps.Vmvmpsreconfirmresult}
    vmvmps.EntityData.Leafs["vmVmpsCurrent"] = types.YLeaf{"Vmvmpscurrent", vmvmps.Vmvmpscurrent}
    return &(vmvmps.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm represents Reading this object always return ready(1).
type CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm string

const (
    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm_ready CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm = "ready"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm_execute CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirm = "execute"
)

// CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult represents      noHostConnected(7) - No hosts on dynamic ports
type CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult string

const (
    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_other CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "other"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_inProgress CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "inProgress"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_success CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "success"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_noResponse CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "noResponse"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_noVmps CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "noVmps"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_noDynamicPort CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "noDynamicPort"

    CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult_noHostConnected CISCOVLANMEMBERSHIPMIB_Vmvmps_Vmvmpsreconfirmresult = "noHostConnected"
)

// CISCOVLANMEMBERSHIPMIB_Vmmembership
type CISCOVLANMEMBERSHIPMIB_Vmmembership struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object is used to determine whether or not a non-existing VLAN will be
    // created automatically by the system after assigned to a port. 
    // automatic(1):  a non-existing VLAN will be created               
    // automatically by the system after                assigned to a port. 
    // manual(2):     a non-existing VLAN will not be created               
    // automatically by the system and need to be                manually created
    // by the users after assigned                to a port. The type is
    // Vmvlancreationmode.
    Vmvlancreationmode interface{}
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetEntityData() *types.CommonEntityData {
    vmmembership.EntityData.YFilter = vmmembership.YFilter
    vmmembership.EntityData.YangName = "vmMembership"
    vmmembership.EntityData.BundleName = "cisco_ios_xe"
    vmmembership.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmmembership.EntityData.SegmentPath = "vmMembership"
    vmmembership.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembership.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembership.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembership.EntityData.Children = make(map[string]types.YChild)
    vmmembership.EntityData.Leafs = make(map[string]types.YLeaf)
    vmmembership.EntityData.Leafs["vmVlanCreationMode"] = types.YLeaf{"Vmvlancreationmode", vmmembership.Vmvlancreationmode}
    return &(vmmembership.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode represents                to a port.
type CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode string

const (
    CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode_automatic CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode = "automatic"

    CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode_manual CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode = "manual"
)

// CISCOVLANMEMBERSHIPMIB_Vmstatistics
type CISCOVLANMEMBERSHIPMIB_Vmstatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of VQP requests sent by this device to all VMPS since last
    // system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    Vmvqpqueries interface{}

    // The number of VQP responses received by this device from all VMPS since
    // last system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    Vmvqpresponses interface{}

    // The number of times, since last system re-initialization, the current VMPS
    // was changed. The current VMPS is changed whenever the VMPS fails to 
    // response after vmVmpsRetries of a VQP request. The type is interface{} with
    // range: 0..4294967295.
    Vmvmpschanges interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  'shutdown'. A 'shutdown' response is a result of  the membership
    // policy configured at a VMPS by the administrator. The type is interface{}
    // with range: 0..4294967295.
    Vmvqpshutdown interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  'denied'. A 'denied' response is a result of  the membership
    // policy configured at a VMPS by the administrator. The type is interface{}
    // with range: 0..4294967295.
    Vmvqpdenied interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates wrong  management domain. A wrong management domain  response
    // indicates that the VMPS used serves a  management domain that is different
    // from the device's management domain. The type is interface{} with range:
    // 0..4294967295.
    Vmvqpwrongdomain interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates wrong  VQP version. A wrong VQP version response  indicates that
    // the VMPS used supports a VQP  version that is different from the device's 
    // VQP version. The type is interface{} with range: 0..4294967295.
    Vmvqpwrongversion interface{}

    // The number of times, since last system re-initialization, a VQP response
    // indicates  insufficient resources. An insufficient resources  response
    // indicates that the VMPS used does not  have the required resources to
    // verify the membership assignment requested. The type is interface{} with
    // range: 0..4294967295.
    Vminsufficientresources interface{}
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetEntityData() *types.CommonEntityData {
    vmstatistics.EntityData.YFilter = vmstatistics.YFilter
    vmstatistics.EntityData.YangName = "vmStatistics"
    vmstatistics.EntityData.BundleName = "cisco_ios_xe"
    vmstatistics.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmstatistics.EntityData.SegmentPath = "vmStatistics"
    vmstatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmstatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmstatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmstatistics.EntityData.Children = make(map[string]types.YChild)
    vmstatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    vmstatistics.EntityData.Leafs["vmVQPQueries"] = types.YLeaf{"Vmvqpqueries", vmstatistics.Vmvqpqueries}
    vmstatistics.EntityData.Leafs["vmVQPResponses"] = types.YLeaf{"Vmvqpresponses", vmstatistics.Vmvqpresponses}
    vmstatistics.EntityData.Leafs["vmVmpsChanges"] = types.YLeaf{"Vmvmpschanges", vmstatistics.Vmvmpschanges}
    vmstatistics.EntityData.Leafs["vmVQPShutdown"] = types.YLeaf{"Vmvqpshutdown", vmstatistics.Vmvqpshutdown}
    vmstatistics.EntityData.Leafs["vmVQPDenied"] = types.YLeaf{"Vmvqpdenied", vmstatistics.Vmvqpdenied}
    vmstatistics.EntityData.Leafs["vmVQPWrongDomain"] = types.YLeaf{"Vmvqpwrongdomain", vmstatistics.Vmvqpwrongdomain}
    vmstatistics.EntityData.Leafs["vmVQPWrongVersion"] = types.YLeaf{"Vmvqpwrongversion", vmstatistics.Vmvqpwrongversion}
    vmstatistics.EntityData.Leafs["vmInsufficientResources"] = types.YLeaf{"Vminsufficientresources", vmstatistics.Vminsufficientresources}
    return &(vmstatistics.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmstatus
type CISCOVLANMEMBERSHIPMIB_Vmstatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the notifications/traps defined in this MIB are
    // enabled. The type is bool.
    Vmnotificationsenabled interface{}
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetEntityData() *types.CommonEntityData {
    vmstatus.EntityData.YFilter = vmstatus.YFilter
    vmstatus.EntityData.YangName = "vmStatus"
    vmstatus.EntityData.BundleName = "cisco_ios_xe"
    vmstatus.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmstatus.EntityData.SegmentPath = "vmStatus"
    vmstatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmstatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmstatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmstatus.EntityData.Children = make(map[string]types.YChild)
    vmstatus.EntityData.Leafs = make(map[string]types.YLeaf)
    vmstatus.EntityData.Leafs["vmNotificationsEnabled"] = types.YLeaf{"Vmnotificationsenabled", vmstatus.Vmnotificationsenabled}
    return &(vmstatus.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvmpstable
// A table of VMPS to use. The device will use
// the the primary VMPS by default. If the
// device is unable to reach the primary server
// after vmVmpsRetries retries, it uses the first
// secondary server in the table until it runs out
// of secondary servers, in which case it will return
// to using the primary server. Entries in this table
// may be created and deleted via this MIB or
// the management console on a device.
type CISCOVLANMEMBERSHIPMIB_Vmvmpstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVmpsTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry.
    Vmvmpsentry []CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetEntityData() *types.CommonEntityData {
    vmvmpstable.EntityData.YFilter = vmvmpstable.YFilter
    vmvmpstable.EntityData.YangName = "vmVmpsTable"
    vmvmpstable.EntityData.BundleName = "cisco_ios_xe"
    vmvmpstable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmvmpstable.EntityData.SegmentPath = "vmVmpsTable"
    vmvmpstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmvmpstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmvmpstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmvmpstable.EntityData.Children = make(map[string]types.YChild)
    vmvmpstable.EntityData.Children["vmVmpsEntry"] = types.YChild{"Vmvmpsentry", nil}
    for i := range vmvmpstable.Vmvmpsentry {
        vmvmpstable.EntityData.Children[types.GetSegmentPath(&vmvmpstable.Vmvmpsentry[i])] = types.YChild{"Vmvmpsentry", &vmvmpstable.Vmvmpsentry[i]}
    }
    vmvmpstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vmvmpstable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry
// An entry (conceptual row) in the vmVmpsTable.
type CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Ip Address of the VMPS. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Vmvmpsipaddress interface{}

    // The status of the VMPS. Setting this value to true will make this VMPS the
    // primary server and make the switch use this as the current server. Setting
    // this entry to true causes other rows to transition to false. Attempting to
    // write a value of false after creation will result in a return of bad value.
    // Deleting an entry whose value is true will result in the first entry in the
    // table being set to true. The type is bool.
    Vmvmpsprimary interface{}

    // The status of this conceptual row. The type is RowStatus.
    Vmvmpsrowstatus interface{}
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetEntityData() *types.CommonEntityData {
    vmvmpsentry.EntityData.YFilter = vmvmpsentry.YFilter
    vmvmpsentry.EntityData.YangName = "vmVmpsEntry"
    vmvmpsentry.EntityData.BundleName = "cisco_ios_xe"
    vmvmpsentry.EntityData.ParentYangName = "vmVmpsTable"
    vmvmpsentry.EntityData.SegmentPath = "vmVmpsEntry" + "[vmVmpsIpAddress='" + fmt.Sprintf("%v", vmvmpsentry.Vmvmpsipaddress) + "']"
    vmvmpsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmvmpsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmvmpsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmvmpsentry.EntityData.Children = make(map[string]types.YChild)
    vmvmpsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vmvmpsentry.EntityData.Leafs["vmVmpsIpAddress"] = types.YLeaf{"Vmvmpsipaddress", vmvmpsentry.Vmvmpsipaddress}
    vmvmpsentry.EntityData.Leafs["vmVmpsPrimary"] = types.YLeaf{"Vmvmpsprimary", vmvmpsentry.Vmvmpsprimary}
    vmvmpsentry.EntityData.Leafs["vmVmpsRowStatus"] = types.YLeaf{"Vmvmpsrowstatus", vmvmpsentry.Vmvmpsrowstatus}
    return &(vmvmpsentry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable
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
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry.
    Vmmembershipsummaryentry []CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetEntityData() *types.CommonEntityData {
    vmmembershipsummarytable.EntityData.YFilter = vmmembershipsummarytable.YFilter
    vmmembershipsummarytable.EntityData.YangName = "vmMembershipSummaryTable"
    vmmembershipsummarytable.EntityData.BundleName = "cisco_ios_xe"
    vmmembershipsummarytable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmmembershipsummarytable.EntityData.SegmentPath = "vmMembershipSummaryTable"
    vmmembershipsummarytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershipsummarytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershipsummarytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershipsummarytable.EntityData.Children = make(map[string]types.YChild)
    vmmembershipsummarytable.EntityData.Children["vmMembershipSummaryEntry"] = types.YChild{"Vmmembershipsummaryentry", nil}
    for i := range vmmembershipsummarytable.Vmmembershipsummaryentry {
        vmmembershipsummarytable.EntityData.Children[types.GetSegmentPath(&vmmembershipsummarytable.Vmmembershipsummaryentry[i])] = types.YChild{"Vmmembershipsummaryentry", &vmmembershipsummarytable.Vmmembershipsummaryentry[i]}
    }
    vmmembershipsummarytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vmmembershipsummarytable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry
// An entry (conceptual row) in the
// vmMembershipSummaryTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    Vmmembershipsummaryvlanindex interface{}

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
    Vmmembershipsummarymemberports interface{}

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
    Vmmembershipsummarymember2Kports interface{}
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetEntityData() *types.CommonEntityData {
    vmmembershipsummaryentry.EntityData.YFilter = vmmembershipsummaryentry.YFilter
    vmmembershipsummaryentry.EntityData.YangName = "vmMembershipSummaryEntry"
    vmmembershipsummaryentry.EntityData.BundleName = "cisco_ios_xe"
    vmmembershipsummaryentry.EntityData.ParentYangName = "vmMembershipSummaryTable"
    vmmembershipsummaryentry.EntityData.SegmentPath = "vmMembershipSummaryEntry" + "[vmMembershipSummaryVlanIndex='" + fmt.Sprintf("%v", vmmembershipsummaryentry.Vmmembershipsummaryvlanindex) + "']"
    vmmembershipsummaryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershipsummaryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershipsummaryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershipsummaryentry.EntityData.Children = make(map[string]types.YChild)
    vmmembershipsummaryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vmmembershipsummaryentry.EntityData.Leafs["vmMembershipSummaryVlanIndex"] = types.YLeaf{"Vmmembershipsummaryvlanindex", vmmembershipsummaryentry.Vmmembershipsummaryvlanindex}
    vmmembershipsummaryentry.EntityData.Leafs["vmMembershipSummaryMemberPorts"] = types.YLeaf{"Vmmembershipsummarymemberports", vmmembershipsummaryentry.Vmmembershipsummarymemberports}
    vmmembershipsummaryentry.EntityData.Leafs["vmMembershipSummaryMember2kPorts"] = types.YLeaf{"Vmmembershipsummarymember2Kports", vmmembershipsummaryentry.Vmmembershipsummarymember2Kports}
    return &(vmmembershipsummaryentry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable
// A table for configuring VLAN port membership.
// There is one row for each bridge port that is
// assigned to a static or dynamic access port. Trunk
// ports are not  represented in this table.  An entry
// may be created and deleted when ports are created or
// deleted via SNMP or the management console on a 
// device.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry.
    Vmmembershipentry []CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetEntityData() *types.CommonEntityData {
    vmmembershiptable.EntityData.YFilter = vmmembershiptable.YFilter
    vmmembershiptable.EntityData.YangName = "vmMembershipTable"
    vmmembershiptable.EntityData.BundleName = "cisco_ios_xe"
    vmmembershiptable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmmembershiptable.EntityData.SegmentPath = "vmMembershipTable"
    vmmembershiptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershiptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershiptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershiptable.EntityData.Children = make(map[string]types.YChild)
    vmmembershiptable.EntityData.Children["vmMembershipEntry"] = types.YChild{"Vmmembershipentry", nil}
    for i := range vmmembershiptable.Vmmembershipentry {
        vmmembershiptable.EntityData.Children[types.GetSegmentPath(&vmmembershiptable.Vmmembershipentry[i])] = types.YChild{"Vmmembershipentry", &vmmembershiptable.Vmmembershipentry[i]}
    }
    vmmembershiptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vmmembershiptable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry
// An entry (conceptual row) in the vmMembershipTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The type of VLAN membership assigned to this port. A port with static vlan
    // membership is assigned to a single VLAN directly. A port with dynamic
    // membership is assigned a single VLAN based on content of packets received
    // on the port and via VQP queries to VMPS. A port with multiVlan membership
    // may be assigned to one or more VLANs directly.  A static or dynamic port
    // membership is specified by the value of vmVlan. A multiVlan port membership
    // is specified by the value of vmVlans. The type is Vmvlantype.
    Vmvlantype interface{}

    // The VLAN id of the VLAN the port is assigned to when vmVlanType is set to
    // static or dynamic. This object is not instantiated if not applicable.  The
    // value may be 0 if the port is not assigned to a VLAN.  If vmVlanType is
    // static, the port is always assigned to a VLAN and the object may not be set
    // to 0.  If vmVlanType is dynamic the object's value is 0 if the port is
    // currently not assigned to a VLAN. In addition, the object may be set to 0
    // only. The type is interface{} with range: 0..4095.
    Vmvlan interface{}

    // An indication of the current VLAN status of the port. A status of
    // inactive(1) indicates that a dynamic port does not yet have a VLAN
    // assigned, or a port is  assigned to a VLAN that is currently not active. A 
    // status of active(2) indicates that the currently  assigned VLAN is active.
    // A status of shutdown(3)  indicates that the port has been disabled as a
    // result of VQP shutdown response. The type is Vmportstatus.
    Vmportstatus interface{}

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
    Vmvlans interface{}

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
    Vmvlans2K interface{}

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
    Vmvlans3K interface{}

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
    Vmvlans4K interface{}
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetEntityData() *types.CommonEntityData {
    vmmembershipentry.EntityData.YFilter = vmmembershipentry.YFilter
    vmmembershipentry.EntityData.YangName = "vmMembershipEntry"
    vmmembershipentry.EntityData.BundleName = "cisco_ios_xe"
    vmmembershipentry.EntityData.ParentYangName = "vmMembershipTable"
    vmmembershipentry.EntityData.SegmentPath = "vmMembershipEntry" + "[ifIndex='" + fmt.Sprintf("%v", vmmembershipentry.Ifindex) + "']"
    vmmembershipentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershipentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershipentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershipentry.EntityData.Children = make(map[string]types.YChild)
    vmmembershipentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vmmembershipentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", vmmembershipentry.Ifindex}
    vmmembershipentry.EntityData.Leafs["vmVlanType"] = types.YLeaf{"Vmvlantype", vmmembershipentry.Vmvlantype}
    vmmembershipentry.EntityData.Leafs["vmVlan"] = types.YLeaf{"Vmvlan", vmmembershipentry.Vmvlan}
    vmmembershipentry.EntityData.Leafs["vmPortStatus"] = types.YLeaf{"Vmportstatus", vmmembershipentry.Vmportstatus}
    vmmembershipentry.EntityData.Leafs["vmVlans"] = types.YLeaf{"Vmvlans", vmmembershipentry.Vmvlans}
    vmmembershipentry.EntityData.Leafs["vmVlans2k"] = types.YLeaf{"Vmvlans2K", vmmembershipentry.Vmvlans2K}
    vmmembershipentry.EntityData.Leafs["vmVlans3k"] = types.YLeaf{"Vmvlans3K", vmmembershipentry.Vmvlans3K}
    vmmembershipentry.EntityData.Leafs["vmVlans4k"] = types.YLeaf{"Vmvlans4K", vmmembershipentry.Vmvlans4K}
    return &(vmmembershipentry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus represents of VQP shutdown response.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus string

const (
    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus_inactive CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus = "inactive"

    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus_active CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus = "active"

    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus_shutdown CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmportstatus = "shutdown"
)

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype represents specified by the value of vmVlans.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype string

const (
    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype_static CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype = "static"

    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype_dynamic CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype = "dynamic"

    CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype_multiVlan CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry_Vmvlantype = "multiVlan"
)

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable
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
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryExtTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry.
    Vmmembershipsummaryextentry []CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetEntityData() *types.CommonEntityData {
    vmmembershipsummaryexttable.EntityData.YFilter = vmmembershipsummaryexttable.YFilter
    vmmembershipsummaryexttable.EntityData.YangName = "vmMembershipSummaryExtTable"
    vmmembershipsummaryexttable.EntityData.BundleName = "cisco_ios_xe"
    vmmembershipsummaryexttable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmmembershipsummaryexttable.EntityData.SegmentPath = "vmMembershipSummaryExtTable"
    vmmembershipsummaryexttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershipsummaryexttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershipsummaryexttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershipsummaryexttable.EntityData.Children = make(map[string]types.YChild)
    vmmembershipsummaryexttable.EntityData.Children["vmMembershipSummaryExtEntry"] = types.YChild{"Vmmembershipsummaryextentry", nil}
    for i := range vmmembershipsummaryexttable.Vmmembershipsummaryextentry {
        vmmembershipsummaryexttable.EntityData.Children[types.GetSegmentPath(&vmmembershipsummaryexttable.Vmmembershipsummaryextentry[i])] = types.YChild{"Vmmembershipsummaryextentry", &vmmembershipsummaryexttable.Vmmembershipsummaryextentry[i]}
    }
    vmmembershipsummaryexttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vmmembershipsummaryexttable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry
// An entry (conceptual row) in the
// vmMembershipSummaryExtTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // cisco_vlan_membership_mib.CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry_Vmmembershipsummaryvlanindex
    Vmmembershipsummaryvlanindex interface{}

    // This attribute is a key. The bridge port range index of this row. The type
    // is CiscoPortListRange.
    Vmmembershipportrangeindex interface{}

    // The set of the device's member ports that belong to the VLAN. It has the
    // VLAN membership information of up to 2k ports with the port number starting
    // from the information indicated in vmMembershipPortRangeIndex object of the
    // same row. For example, if the value of vmMembershipPortRangeIndex is
    // 'twoKto4K', the port number indicated in this object starting from 2049 and
    // ending to 4096.   A port number is the value of dot1dBasePort for the port
    // in the BRIDGE-MIB (RFC 1493). The type is string with length: 0..256.
    Vmmembershipsummaryextports interface{}
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetEntityData() *types.CommonEntityData {
    vmmembershipsummaryextentry.EntityData.YFilter = vmmembershipsummaryextentry.YFilter
    vmmembershipsummaryextentry.EntityData.YangName = "vmMembershipSummaryExtEntry"
    vmmembershipsummaryextentry.EntityData.BundleName = "cisco_ios_xe"
    vmmembershipsummaryextentry.EntityData.ParentYangName = "vmMembershipSummaryExtTable"
    vmmembershipsummaryextentry.EntityData.SegmentPath = "vmMembershipSummaryExtEntry" + "[vmMembershipSummaryVlanIndex='" + fmt.Sprintf("%v", vmmembershipsummaryextentry.Vmmembershipsummaryvlanindex) + "']" + "[vmMembershipPortRangeIndex='" + fmt.Sprintf("%v", vmmembershipsummaryextentry.Vmmembershipportrangeindex) + "']"
    vmmembershipsummaryextentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmmembershipsummaryextentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmmembershipsummaryextentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmmembershipsummaryextentry.EntityData.Children = make(map[string]types.YChild)
    vmmembershipsummaryextentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vmmembershipsummaryextentry.EntityData.Leafs["vmMembershipSummaryVlanIndex"] = types.YLeaf{"Vmmembershipsummaryvlanindex", vmmembershipsummaryextentry.Vmmembershipsummaryvlanindex}
    vmmembershipsummaryextentry.EntityData.Leafs["vmMembershipPortRangeIndex"] = types.YLeaf{"Vmmembershipportrangeindex", vmmembershipsummaryextentry.Vmmembershipportrangeindex}
    vmmembershipsummaryextentry.EntityData.Leafs["vmMembershipSummaryExtPorts"] = types.YLeaf{"Vmmembershipsummaryextports", vmmembershipsummaryextentry.Vmmembershipsummaryextports}
    return &(vmmembershipsummaryextentry.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable
// A table for configuring the Voice VLAN-ID
// for the ports. An entry will exist for each
// interface which supports Voice Vlan feature.
type CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVoiceVlanTable. Only interfaces which
    // support Voice Vlan feature are shown. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry.
    Vmvoicevlanentry []CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetEntityData() *types.CommonEntityData {
    vmvoicevlantable.EntityData.YFilter = vmvoicevlantable.YFilter
    vmvoicevlantable.EntityData.YangName = "vmVoiceVlanTable"
    vmvoicevlantable.EntityData.BundleName = "cisco_ios_xe"
    vmvoicevlantable.EntityData.ParentYangName = "CISCO-VLAN-MEMBERSHIP-MIB"
    vmvoicevlantable.EntityData.SegmentPath = "vmVoiceVlanTable"
    vmvoicevlantable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmvoicevlantable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmvoicevlantable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmvoicevlantable.EntityData.Children = make(map[string]types.YChild)
    vmvoicevlantable.EntityData.Children["vmVoiceVlanEntry"] = types.YChild{"Vmvoicevlanentry", nil}
    for i := range vmvoicevlantable.Vmvoicevlanentry {
        vmvoicevlantable.EntityData.Children[types.GetSegmentPath(&vmvoicevlantable.Vmvoicevlanentry[i])] = types.YChild{"Vmvoicevlanentry", &vmvoicevlantable.Vmvoicevlanentry[i]}
    }
    vmvoicevlantable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vmvoicevlantable.EntityData)
}

// CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry
// An entry (conceptual row) in the vmVoiceVlanTable.
// Only interfaces which support Voice Vlan feature
// are shown.
type CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

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
    Vmvoicevlanid interface{}

    // Enable or Disable the feature of CDP message verification of voice VLANs. 
    // true   - The voice VLAN vmVoiceVlan is enabled           only after CDP
    // messages are received           from the IP phone.  false -  The voice VLAN
    // vmVoiceVlan is enabled          as soon as the IP phone interface is       
    // up. There is no verification needed           from CDP messages from the IP
    // phone. The type is bool.
    Vmvoicevlancdpverifyenable interface{}
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetEntityData() *types.CommonEntityData {
    vmvoicevlanentry.EntityData.YFilter = vmvoicevlanentry.YFilter
    vmvoicevlanentry.EntityData.YangName = "vmVoiceVlanEntry"
    vmvoicevlanentry.EntityData.BundleName = "cisco_ios_xe"
    vmvoicevlanentry.EntityData.ParentYangName = "vmVoiceVlanTable"
    vmvoicevlanentry.EntityData.SegmentPath = "vmVoiceVlanEntry" + "[ifIndex='" + fmt.Sprintf("%v", vmvoicevlanentry.Ifindex) + "']"
    vmvoicevlanentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vmvoicevlanentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vmvoicevlanentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vmvoicevlanentry.EntityData.Children = make(map[string]types.YChild)
    vmvoicevlanentry.EntityData.Leafs = make(map[string]types.YLeaf)
    vmvoicevlanentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", vmvoicevlanentry.Ifindex}
    vmvoicevlanentry.EntityData.Leafs["vmVoiceVlanId"] = types.YLeaf{"Vmvoicevlanid", vmvoicevlanentry.Vmvoicevlanid}
    vmvoicevlanentry.EntityData.Leafs["vmVoiceVlanCdpVerifyEnable"] = types.YLeaf{"Vmvoicevlancdpverifyenable", vmvoicevlanentry.Vmvoicevlancdpverifyenable}
    return &(vmvoicevlanentry.EntityData)
}

