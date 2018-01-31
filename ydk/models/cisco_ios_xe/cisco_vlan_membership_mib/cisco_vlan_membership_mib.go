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
    parent types.Entity
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

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetFilter() yfilter.YFilter { return cISCOVLANMEMBERSHIPMIB.YFilter }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) SetFilter(yf yfilter.YFilter) { cISCOVLANMEMBERSHIPMIB.YFilter = yf }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetGoName(yname string) string {
    if yname == "vmVmps" { return "Vmvmps" }
    if yname == "vmMembership" { return "Vmmembership" }
    if yname == "vmStatistics" { return "Vmstatistics" }
    if yname == "vmStatus" { return "Vmstatus" }
    if yname == "vmVmpsTable" { return "Vmvmpstable" }
    if yname == "vmMembershipSummaryTable" { return "Vmmembershipsummarytable" }
    if yname == "vmMembershipTable" { return "Vmmembershiptable" }
    if yname == "vmMembershipSummaryExtTable" { return "Vmmembershipsummaryexttable" }
    if yname == "vmVoiceVlanTable" { return "Vmvoicevlantable" }
    return ""
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetSegmentPath() string {
    return "CISCO-VLAN-MEMBERSHIP-MIB:CISCO-VLAN-MEMBERSHIP-MIB"
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmVmps" {
        return &cISCOVLANMEMBERSHIPMIB.Vmvmps
    }
    if childYangName == "vmMembership" {
        return &cISCOVLANMEMBERSHIPMIB.Vmmembership
    }
    if childYangName == "vmStatistics" {
        return &cISCOVLANMEMBERSHIPMIB.Vmstatistics
    }
    if childYangName == "vmStatus" {
        return &cISCOVLANMEMBERSHIPMIB.Vmstatus
    }
    if childYangName == "vmVmpsTable" {
        return &cISCOVLANMEMBERSHIPMIB.Vmvmpstable
    }
    if childYangName == "vmMembershipSummaryTable" {
        return &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummarytable
    }
    if childYangName == "vmMembershipTable" {
        return &cISCOVLANMEMBERSHIPMIB.Vmmembershiptable
    }
    if childYangName == "vmMembershipSummaryExtTable" {
        return &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummaryexttable
    }
    if childYangName == "vmVoiceVlanTable" {
        return &cISCOVLANMEMBERSHIPMIB.Vmvoicevlantable
    }
    return nil
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vmVmps"] = &cISCOVLANMEMBERSHIPMIB.Vmvmps
    children["vmMembership"] = &cISCOVLANMEMBERSHIPMIB.Vmmembership
    children["vmStatistics"] = &cISCOVLANMEMBERSHIPMIB.Vmstatistics
    children["vmStatus"] = &cISCOVLANMEMBERSHIPMIB.Vmstatus
    children["vmVmpsTable"] = &cISCOVLANMEMBERSHIPMIB.Vmvmpstable
    children["vmMembershipSummaryTable"] = &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummarytable
    children["vmMembershipTable"] = &cISCOVLANMEMBERSHIPMIB.Vmmembershiptable
    children["vmMembershipSummaryExtTable"] = &cISCOVLANMEMBERSHIPMIB.Vmmembershipsummaryexttable
    children["vmVoiceVlanTable"] = &cISCOVLANMEMBERSHIPMIB.Vmvoicevlantable
    return children
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) SetParent(parent types.Entity) { cISCOVLANMEMBERSHIPMIB.parent = parent }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetParent() types.Entity { return cISCOVLANMEMBERSHIPMIB.parent }

func (cISCOVLANMEMBERSHIPMIB *CISCOVLANMEMBERSHIPMIB) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmvmps
type CISCOVLANMEMBERSHIPMIB_Vmvmps struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Vmvmpscurrent interface{}
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetFilter() yfilter.YFilter { return vmvmps.YFilter }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) SetFilter(yf yfilter.YFilter) { vmvmps.YFilter = yf }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetGoName(yname string) string {
    if yname == "vmVmpsVQPVersion" { return "Vmvmpsvqpversion" }
    if yname == "vmVmpsRetries" { return "Vmvmpsretries" }
    if yname == "vmVmpsReconfirmInterval" { return "Vmvmpsreconfirminterval" }
    if yname == "vmVmpsReconfirm" { return "Vmvmpsreconfirm" }
    if yname == "vmVmpsReconfirmResult" { return "Vmvmpsreconfirmresult" }
    if yname == "vmVmpsCurrent" { return "Vmvmpscurrent" }
    return ""
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetSegmentPath() string {
    return "vmVmps"
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmVmpsVQPVersion"] = vmvmps.Vmvmpsvqpversion
    leafs["vmVmpsRetries"] = vmvmps.Vmvmpsretries
    leafs["vmVmpsReconfirmInterval"] = vmvmps.Vmvmpsreconfirminterval
    leafs["vmVmpsReconfirm"] = vmvmps.Vmvmpsreconfirm
    leafs["vmVmpsReconfirmResult"] = vmvmps.Vmvmpsreconfirmresult
    leafs["vmVmpsCurrent"] = vmvmps.Vmvmpscurrent
    return leafs
}

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetBundleName() string { return "cisco_ios_xe" }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetYangName() string { return "vmVmps" }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) SetParent(parent types.Entity) { vmvmps.parent = parent }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetParent() types.Entity { return vmvmps.parent }

func (vmvmps *CISCOVLANMEMBERSHIPMIB_Vmvmps) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

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
    parent types.Entity
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

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetFilter() yfilter.YFilter { return vmmembership.YFilter }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) SetFilter(yf yfilter.YFilter) { vmmembership.YFilter = yf }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetGoName(yname string) string {
    if yname == "vmVlanCreationMode" { return "Vmvlancreationmode" }
    return ""
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetSegmentPath() string {
    return "vmMembership"
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmVlanCreationMode"] = vmmembership.Vmvlancreationmode
    return leafs
}

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetYangName() string { return "vmMembership" }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) SetParent(parent types.Entity) { vmmembership.parent = parent }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetParent() types.Entity { return vmmembership.parent }

func (vmmembership *CISCOVLANMEMBERSHIPMIB_Vmmembership) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode represents                to a port.
type CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode string

const (
    CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode_automatic CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode = "automatic"

    CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode_manual CISCOVLANMEMBERSHIPMIB_Vmmembership_Vmvlancreationmode = "manual"
)

// CISCOVLANMEMBERSHIPMIB_Vmstatistics
type CISCOVLANMEMBERSHIPMIB_Vmstatistics struct {
    parent types.Entity
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

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetFilter() yfilter.YFilter { return vmstatistics.YFilter }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) SetFilter(yf yfilter.YFilter) { vmstatistics.YFilter = yf }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetGoName(yname string) string {
    if yname == "vmVQPQueries" { return "Vmvqpqueries" }
    if yname == "vmVQPResponses" { return "Vmvqpresponses" }
    if yname == "vmVmpsChanges" { return "Vmvmpschanges" }
    if yname == "vmVQPShutdown" { return "Vmvqpshutdown" }
    if yname == "vmVQPDenied" { return "Vmvqpdenied" }
    if yname == "vmVQPWrongDomain" { return "Vmvqpwrongdomain" }
    if yname == "vmVQPWrongVersion" { return "Vmvqpwrongversion" }
    if yname == "vmInsufficientResources" { return "Vminsufficientresources" }
    return ""
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetSegmentPath() string {
    return "vmStatistics"
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmVQPQueries"] = vmstatistics.Vmvqpqueries
    leafs["vmVQPResponses"] = vmstatistics.Vmvqpresponses
    leafs["vmVmpsChanges"] = vmstatistics.Vmvmpschanges
    leafs["vmVQPShutdown"] = vmstatistics.Vmvqpshutdown
    leafs["vmVQPDenied"] = vmstatistics.Vmvqpdenied
    leafs["vmVQPWrongDomain"] = vmstatistics.Vmvqpwrongdomain
    leafs["vmVQPWrongVersion"] = vmstatistics.Vmvqpwrongversion
    leafs["vmInsufficientResources"] = vmstatistics.Vminsufficientresources
    return leafs
}

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetBundleName() string { return "cisco_ios_xe" }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetYangName() string { return "vmStatistics" }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) SetParent(parent types.Entity) { vmstatistics.parent = parent }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetParent() types.Entity { return vmstatistics.parent }

func (vmstatistics *CISCOVLANMEMBERSHIPMIB_Vmstatistics) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmstatus
type CISCOVLANMEMBERSHIPMIB_Vmstatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An indication of whether the notifications/traps defined in this MIB are
    // enabled. The type is bool.
    Vmnotificationsenabled interface{}
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetFilter() yfilter.YFilter { return vmstatus.YFilter }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) SetFilter(yf yfilter.YFilter) { vmstatus.YFilter = yf }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetGoName(yname string) string {
    if yname == "vmNotificationsEnabled" { return "Vmnotificationsenabled" }
    return ""
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetSegmentPath() string {
    return "vmStatus"
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmNotificationsEnabled"] = vmstatus.Vmnotificationsenabled
    return leafs
}

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetBundleName() string { return "cisco_ios_xe" }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetYangName() string { return "vmStatus" }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) SetParent(parent types.Entity) { vmstatus.parent = parent }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetParent() types.Entity { return vmstatus.parent }

func (vmstatus *CISCOVLANMEMBERSHIPMIB_Vmstatus) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVmpsTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry.
    Vmvmpsentry []CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetFilter() yfilter.YFilter { return vmvmpstable.YFilter }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) SetFilter(yf yfilter.YFilter) { vmvmpstable.YFilter = yf }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetGoName(yname string) string {
    if yname == "vmVmpsEntry" { return "Vmvmpsentry" }
    return ""
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetSegmentPath() string {
    return "vmVmpsTable"
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmVmpsEntry" {
        for _, c := range vmvmpstable.Vmvmpsentry {
            if vmvmpstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry{}
        vmvmpstable.Vmvmpsentry = append(vmvmpstable.Vmvmpsentry, child)
        return &vmvmpstable.Vmvmpsentry[len(vmvmpstable.Vmvmpsentry)-1]
    }
    return nil
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vmvmpstable.Vmvmpsentry {
        children[vmvmpstable.Vmvmpsentry[i].GetSegmentPath()] = &vmvmpstable.Vmvmpsentry[i]
    }
    return children
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetBundleName() string { return "cisco_ios_xe" }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetYangName() string { return "vmVmpsTable" }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) SetParent(parent types.Entity) { vmvmpstable.parent = parent }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetParent() types.Entity { return vmvmpstable.parent }

func (vmvmpstable *CISCOVLANMEMBERSHIPMIB_Vmvmpstable) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry
// An entry (conceptual row) in the vmVmpsTable.
type CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Ip Address of the VMPS. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetFilter() yfilter.YFilter { return vmvmpsentry.YFilter }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) SetFilter(yf yfilter.YFilter) { vmvmpsentry.YFilter = yf }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetGoName(yname string) string {
    if yname == "vmVmpsIpAddress" { return "Vmvmpsipaddress" }
    if yname == "vmVmpsPrimary" { return "Vmvmpsprimary" }
    if yname == "vmVmpsRowStatus" { return "Vmvmpsrowstatus" }
    return ""
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetSegmentPath() string {
    return "vmVmpsEntry" + "[vmVmpsIpAddress='" + fmt.Sprintf("%v", vmvmpsentry.Vmvmpsipaddress) + "']"
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmVmpsIpAddress"] = vmvmpsentry.Vmvmpsipaddress
    leafs["vmVmpsPrimary"] = vmvmpsentry.Vmvmpsprimary
    leafs["vmVmpsRowStatus"] = vmvmpsentry.Vmvmpsrowstatus
    return leafs
}

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetBundleName() string { return "cisco_ios_xe" }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetYangName() string { return "vmVmpsEntry" }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) SetParent(parent types.Entity) { vmvmpsentry.parent = parent }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetParent() types.Entity { return vmvmpsentry.parent }

func (vmvmpsentry *CISCOVLANMEMBERSHIPMIB_Vmvmpstable_Vmvmpsentry) GetParentYangName() string { return "vmVmpsTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry.
    Vmmembershipsummaryentry []CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetFilter() yfilter.YFilter { return vmmembershipsummarytable.YFilter }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) SetFilter(yf yfilter.YFilter) { vmmembershipsummarytable.YFilter = yf }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetGoName(yname string) string {
    if yname == "vmMembershipSummaryEntry" { return "Vmmembershipsummaryentry" }
    return ""
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetSegmentPath() string {
    return "vmMembershipSummaryTable"
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmMembershipSummaryEntry" {
        for _, c := range vmmembershipsummarytable.Vmmembershipsummaryentry {
            if vmmembershipsummarytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry{}
        vmmembershipsummarytable.Vmmembershipsummaryentry = append(vmmembershipsummarytable.Vmmembershipsummaryentry, child)
        return &vmmembershipsummarytable.Vmmembershipsummaryentry[len(vmmembershipsummarytable.Vmmembershipsummaryentry)-1]
    }
    return nil
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vmmembershipsummarytable.Vmmembershipsummaryentry {
        children[vmmembershipsummarytable.Vmmembershipsummaryentry[i].GetSegmentPath()] = &vmmembershipsummarytable.Vmmembershipsummaryentry[i]
    }
    return children
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetYangName() string { return "vmMembershipSummaryTable" }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) SetParent(parent types.Entity) { vmmembershipsummarytable.parent = parent }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetParent() types.Entity { return vmmembershipsummarytable.parent }

func (vmmembershipsummarytable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry
// An entry (conceptual row) in the
// vmMembershipSummaryTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry struct {
    parent types.Entity
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

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetFilter() yfilter.YFilter { return vmmembershipsummaryentry.YFilter }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) SetFilter(yf yfilter.YFilter) { vmmembershipsummaryentry.YFilter = yf }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetGoName(yname string) string {
    if yname == "vmMembershipSummaryVlanIndex" { return "Vmmembershipsummaryvlanindex" }
    if yname == "vmMembershipSummaryMemberPorts" { return "Vmmembershipsummarymemberports" }
    if yname == "vmMembershipSummaryMember2kPorts" { return "Vmmembershipsummarymember2Kports" }
    return ""
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetSegmentPath() string {
    return "vmMembershipSummaryEntry" + "[vmMembershipSummaryVlanIndex='" + fmt.Sprintf("%v", vmmembershipsummaryentry.Vmmembershipsummaryvlanindex) + "']"
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmMembershipSummaryVlanIndex"] = vmmembershipsummaryentry.Vmmembershipsummaryvlanindex
    leafs["vmMembershipSummaryMemberPorts"] = vmmembershipsummaryentry.Vmmembershipsummarymemberports
    leafs["vmMembershipSummaryMember2kPorts"] = vmmembershipsummaryentry.Vmmembershipsummarymember2Kports
    return leafs
}

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetYangName() string { return "vmMembershipSummaryEntry" }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) SetParent(parent types.Entity) { vmmembershipsummaryentry.parent = parent }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetParent() types.Entity { return vmmembershipsummaryentry.parent }

func (vmmembershipsummaryentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummarytable_Vmmembershipsummaryentry) GetParentYangName() string { return "vmMembershipSummaryTable" }

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable
// A table for configuring VLAN port membership.
// There is one row for each bridge port that is
// assigned to a static or dynamic access port. Trunk
// ports are not  represented in this table.  An entry
// may be created and deleted when ports are created or
// deleted via SNMP or the management console on a 
// device.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipTable. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry.
    Vmmembershipentry []CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetFilter() yfilter.YFilter { return vmmembershiptable.YFilter }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) SetFilter(yf yfilter.YFilter) { vmmembershiptable.YFilter = yf }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetGoName(yname string) string {
    if yname == "vmMembershipEntry" { return "Vmmembershipentry" }
    return ""
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetSegmentPath() string {
    return "vmMembershipTable"
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmMembershipEntry" {
        for _, c := range vmmembershiptable.Vmmembershipentry {
            if vmmembershiptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry{}
        vmmembershiptable.Vmmembershipentry = append(vmmembershiptable.Vmmembershipentry, child)
        return &vmmembershiptable.Vmmembershipentry[len(vmmembershiptable.Vmmembershipentry)-1]
    }
    return nil
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vmmembershiptable.Vmmembershipentry {
        children[vmmembershiptable.Vmmembershipentry[i].GetSegmentPath()] = &vmmembershiptable.Vmmembershipentry[i]
    }
    return children
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetYangName() string { return "vmMembershipTable" }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) SetParent(parent types.Entity) { vmmembershiptable.parent = parent }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetParent() types.Entity { return vmmembershiptable.parent }

func (vmmembershiptable *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry
// An entry (conceptual row) in the vmMembershipTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry struct {
    parent types.Entity
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

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetFilter() yfilter.YFilter { return vmmembershipentry.YFilter }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) SetFilter(yf yfilter.YFilter) { vmmembershipentry.YFilter = yf }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "vmVlanType" { return "Vmvlantype" }
    if yname == "vmVlan" { return "Vmvlan" }
    if yname == "vmPortStatus" { return "Vmportstatus" }
    if yname == "vmVlans" { return "Vmvlans" }
    if yname == "vmVlans2k" { return "Vmvlans2K" }
    if yname == "vmVlans3k" { return "Vmvlans3K" }
    if yname == "vmVlans4k" { return "Vmvlans4K" }
    return ""
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetSegmentPath() string {
    return "vmMembershipEntry" + "[ifIndex='" + fmt.Sprintf("%v", vmmembershipentry.Ifindex) + "']"
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = vmmembershipentry.Ifindex
    leafs["vmVlanType"] = vmmembershipentry.Vmvlantype
    leafs["vmVlan"] = vmmembershipentry.Vmvlan
    leafs["vmPortStatus"] = vmmembershipentry.Vmportstatus
    leafs["vmVlans"] = vmmembershipentry.Vmvlans
    leafs["vmVlans2k"] = vmmembershipentry.Vmvlans2K
    leafs["vmVlans3k"] = vmmembershipentry.Vmvlans3K
    leafs["vmVlans4k"] = vmmembershipentry.Vmvlans4K
    return leafs
}

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetYangName() string { return "vmMembershipEntry" }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) SetParent(parent types.Entity) { vmmembershipentry.parent = parent }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetParent() types.Entity { return vmmembershipentry.parent }

func (vmmembershipentry *CISCOVLANMEMBERSHIPMIB_Vmmembershiptable_Vmmembershipentry) GetParentYangName() string { return "vmMembershipTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmMembershipSummaryExtTable. The type is
    // slice of
    // CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry.
    Vmmembershipsummaryextentry []CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetFilter() yfilter.YFilter { return vmmembershipsummaryexttable.YFilter }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) SetFilter(yf yfilter.YFilter) { vmmembershipsummaryexttable.YFilter = yf }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetGoName(yname string) string {
    if yname == "vmMembershipSummaryExtEntry" { return "Vmmembershipsummaryextentry" }
    return ""
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetSegmentPath() string {
    return "vmMembershipSummaryExtTable"
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmMembershipSummaryExtEntry" {
        for _, c := range vmmembershipsummaryexttable.Vmmembershipsummaryextentry {
            if vmmembershipsummaryexttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry{}
        vmmembershipsummaryexttable.Vmmembershipsummaryextentry = append(vmmembershipsummaryexttable.Vmmembershipsummaryextentry, child)
        return &vmmembershipsummaryexttable.Vmmembershipsummaryextentry[len(vmmembershipsummaryexttable.Vmmembershipsummaryextentry)-1]
    }
    return nil
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vmmembershipsummaryexttable.Vmmembershipsummaryextentry {
        children[vmmembershipsummaryexttable.Vmmembershipsummaryextentry[i].GetSegmentPath()] = &vmmembershipsummaryexttable.Vmmembershipsummaryextentry[i]
    }
    return children
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetYangName() string { return "vmMembershipSummaryExtTable" }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) SetParent(parent types.Entity) { vmmembershipsummaryexttable.parent = parent }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetParent() types.Entity { return vmmembershipsummaryexttable.parent }

func (vmmembershipsummaryexttable *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry
// An entry (conceptual row) in the
// vmMembershipSummaryExtTable.
type CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry struct {
    parent types.Entity
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

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetFilter() yfilter.YFilter { return vmmembershipsummaryextentry.YFilter }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) SetFilter(yf yfilter.YFilter) { vmmembershipsummaryextentry.YFilter = yf }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetGoName(yname string) string {
    if yname == "vmMembershipSummaryVlanIndex" { return "Vmmembershipsummaryvlanindex" }
    if yname == "vmMembershipPortRangeIndex" { return "Vmmembershipportrangeindex" }
    if yname == "vmMembershipSummaryExtPorts" { return "Vmmembershipsummaryextports" }
    return ""
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetSegmentPath() string {
    return "vmMembershipSummaryExtEntry" + "[vmMembershipSummaryVlanIndex='" + fmt.Sprintf("%v", vmmembershipsummaryextentry.Vmmembershipsummaryvlanindex) + "']" + "[vmMembershipPortRangeIndex='" + fmt.Sprintf("%v", vmmembershipsummaryextentry.Vmmembershipportrangeindex) + "']"
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vmMembershipSummaryVlanIndex"] = vmmembershipsummaryextentry.Vmmembershipsummaryvlanindex
    leafs["vmMembershipPortRangeIndex"] = vmmembershipsummaryextentry.Vmmembershipportrangeindex
    leafs["vmMembershipSummaryExtPorts"] = vmmembershipsummaryextentry.Vmmembershipsummaryextports
    return leafs
}

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetBundleName() string { return "cisco_ios_xe" }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetYangName() string { return "vmMembershipSummaryExtEntry" }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) SetParent(parent types.Entity) { vmmembershipsummaryextentry.parent = parent }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetParent() types.Entity { return vmmembershipsummaryextentry.parent }

func (vmmembershipsummaryextentry *CISCOVLANMEMBERSHIPMIB_Vmmembershipsummaryexttable_Vmmembershipsummaryextentry) GetParentYangName() string { return "vmMembershipSummaryExtTable" }

// CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable
// A table for configuring the Voice VLAN-ID
// for the ports. An entry will exist for each
// interface which supports Voice Vlan feature.
type CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the vmVoiceVlanTable. Only interfaces which
    // support Voice Vlan feature are shown. The type is slice of
    // CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry.
    Vmvoicevlanentry []CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetFilter() yfilter.YFilter { return vmvoicevlantable.YFilter }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) SetFilter(yf yfilter.YFilter) { vmvoicevlantable.YFilter = yf }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetGoName(yname string) string {
    if yname == "vmVoiceVlanEntry" { return "Vmvoicevlanentry" }
    return ""
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetSegmentPath() string {
    return "vmVoiceVlanTable"
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vmVoiceVlanEntry" {
        for _, c := range vmvoicevlantable.Vmvoicevlanentry {
            if vmvoicevlantable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry{}
        vmvoicevlantable.Vmvoicevlanentry = append(vmvoicevlantable.Vmvoicevlanentry, child)
        return &vmvoicevlantable.Vmvoicevlanentry[len(vmvoicevlantable.Vmvoicevlanentry)-1]
    }
    return nil
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vmvoicevlantable.Vmvoicevlanentry {
        children[vmvoicevlantable.Vmvoicevlanentry[i].GetSegmentPath()] = &vmvoicevlantable.Vmvoicevlanentry[i]
    }
    return children
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetBundleName() string { return "cisco_ios_xe" }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetYangName() string { return "vmVoiceVlanTable" }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) SetParent(parent types.Entity) { vmvoicevlantable.parent = parent }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetParent() types.Entity { return vmvoicevlantable.parent }

func (vmvoicevlantable *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable) GetParentYangName() string { return "CISCO-VLAN-MEMBERSHIP-MIB" }

// CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry
// An entry (conceptual row) in the vmVoiceVlanTable.
// Only interfaces which support Voice Vlan feature
// are shown.
type CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry struct {
    parent types.Entity
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

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetFilter() yfilter.YFilter { return vmvoicevlanentry.YFilter }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) SetFilter(yf yfilter.YFilter) { vmvoicevlanentry.YFilter = yf }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "vmVoiceVlanId" { return "Vmvoicevlanid" }
    if yname == "vmVoiceVlanCdpVerifyEnable" { return "Vmvoicevlancdpverifyenable" }
    return ""
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetSegmentPath() string {
    return "vmVoiceVlanEntry" + "[ifIndex='" + fmt.Sprintf("%v", vmvoicevlanentry.Ifindex) + "']"
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = vmvoicevlanentry.Ifindex
    leafs["vmVoiceVlanId"] = vmvoicevlanentry.Vmvoicevlanid
    leafs["vmVoiceVlanCdpVerifyEnable"] = vmvoicevlanentry.Vmvoicevlancdpverifyenable
    return leafs
}

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetBundleName() string { return "cisco_ios_xe" }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetYangName() string { return "vmVoiceVlanEntry" }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) SetParent(parent types.Entity) { vmvoicevlanentry.parent = parent }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetParent() types.Entity { return vmvoicevlanentry.parent }

func (vmvoicevlanentry *CISCOVLANMEMBERSHIPMIB_Vmvoicevlantable_Vmvoicevlanentry) GetParentYangName() string { return "vmVoiceVlanTable" }

