// BGP policy based accounting information
package cisco_bgp_policy_accounting_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_bgp_policy_accounting_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-BGP-POLICY-ACCOUNTING-MIB CISCO-BGP-POLICY-ACCOUNTING-MIB}", reflect.TypeOf(CISCOBGPPOLICYACCOUNTINGMIB{}))
    ydk.RegisterEntity("CISCO-BGP-POLICY-ACCOUNTING-MIB:CISCO-BGP-POLICY-ACCOUNTING-MIB", reflect.TypeOf(CISCOBGPPOLICYACCOUNTINGMIB{}))
}

// CISCOBGPPOLICYACCOUNTINGMIB
type CISCOBGPPOLICYACCOUNTINGMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The cbpAcctTable provides statistics about ingress and egress  traffic on
    // an interface. This data could be used for purposes  like billing.
    Cbpaccttable CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetFilter() yfilter.YFilter { return cISCOBGPPOLICYACCOUNTINGMIB.YFilter }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) SetFilter(yf yfilter.YFilter) { cISCOBGPPOLICYACCOUNTINGMIB.YFilter = yf }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetGoName(yname string) string {
    if yname == "cbpAcctTable" { return "Cbpaccttable" }
    return ""
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetSegmentPath() string {
    return "CISCO-BGP-POLICY-ACCOUNTING-MIB:CISCO-BGP-POLICY-ACCOUNTING-MIB"
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbpAcctTable" {
        return &cISCOBGPPOLICYACCOUNTINGMIB.Cbpaccttable
    }
    return nil
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cbpAcctTable"] = &cISCOBGPPOLICYACCOUNTINGMIB.Cbpaccttable
    return children
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetYangName() string { return "CISCO-BGP-POLICY-ACCOUNTING-MIB" }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) SetParent(parent types.Entity) { cISCOBGPPOLICYACCOUNTINGMIB.parent = parent }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetParent() types.Entity { return cISCOBGPPOLICYACCOUNTINGMIB.parent }

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetParentYangName() string { return "CISCO-BGP-POLICY-ACCOUNTING-MIB" }

// CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable
// The cbpAcctTable provides statistics about ingress and egress 
// traffic on an interface. This data could be used for purposes 
// like billing.
type CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each cbpAcctEntry provides statistics for traffic of interest on an ingress
    // and/or egress interfaces. The traffic of interest  may be used for purposes
    // like billing, and is referred to from  here on in the MIB by the term
    // 'traffic-type', which corresponds to cbpAcctTrafficIndex. Traffic-types are
    // configured by the user on a per interface basis.  The statistics include
    // ingress packet counts, ingress octet counts, egress packet counts and
    // egress octet counts. Entries  are created when traffic-type is configured
    // on an interface. Entries are deleted automatically when the user  removes
    // the corresponding traffic-type configuration from an interface. The type is
    // slice of CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry.
    Cbpacctentry []CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetFilter() yfilter.YFilter { return cbpaccttable.YFilter }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) SetFilter(yf yfilter.YFilter) { cbpaccttable.YFilter = yf }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetGoName(yname string) string {
    if yname == "cbpAcctEntry" { return "Cbpacctentry" }
    return ""
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetSegmentPath() string {
    return "cbpAcctTable"
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cbpAcctEntry" {
        for _, c := range cbpaccttable.Cbpacctentry {
            if cbpaccttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry{}
        cbpaccttable.Cbpacctentry = append(cbpaccttable.Cbpacctentry, child)
        return &cbpaccttable.Cbpacctentry[len(cbpaccttable.Cbpacctentry)-1]
    }
    return nil
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cbpaccttable.Cbpacctentry {
        children[cbpaccttable.Cbpacctentry[i].GetSegmentPath()] = &cbpaccttable.Cbpacctentry[i]
    }
    return children
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetBundleName() string { return "cisco_ios_xe" }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetYangName() string { return "cbpAcctTable" }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) SetParent(parent types.Entity) { cbpaccttable.parent = parent }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetParent() types.Entity { return cbpaccttable.parent }

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetParentYangName() string { return "CISCO-BGP-POLICY-ACCOUNTING-MIB" }

// CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry
// Each cbpAcctEntry provides statistics for traffic of interest
// on an ingress and/or egress interfaces. The traffic of interest 
// may be used for purposes like billing, and is referred to from 
// here on in the MIB by the term 'traffic-type', which corresponds
// to cbpAcctTrafficIndex. Traffic-types are configured by the user
// on a per interface basis.
// 
// The statistics include ingress packet counts, ingress octet
// counts, egress packet counts and egress octet counts. Entries 
// are created when traffic-type is configured on an interface.
// Entries are deleted automatically when the user 
// removes the corresponding traffic-type configuration from an
// interface.
type CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An integer value greater than 0, that uniquely
    // identifies a traffic-type. The traffic-type has no intrinsic meaning. It
    // just means the traffic coming into an interface can be differentiated into
    // different types. It is up to the user to give meaning to and configure the
    // various traffic-types on an  interface. The type is interface{} with range:
    // 1..2147483647.
    Cbpaccttrafficindex interface{}

    // The total number of packets received for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    Cbpacctinpacketcount interface{}

    // The total number of octets received for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    Cbpacctinoctetcount interface{}

    // The total number of packets transmitted for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    Cbpacctoutpacketcount interface{}

    // The total number of octets transmitted for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    Cbpacctoutoctetcount interface{}
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetFilter() yfilter.YFilter { return cbpacctentry.YFilter }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) SetFilter(yf yfilter.YFilter) { cbpacctentry.YFilter = yf }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cbpAcctTrafficIndex" { return "Cbpaccttrafficindex" }
    if yname == "cbpAcctInPacketCount" { return "Cbpacctinpacketcount" }
    if yname == "cbpAcctInOctetCount" { return "Cbpacctinoctetcount" }
    if yname == "cbpAcctOutPacketCount" { return "Cbpacctoutpacketcount" }
    if yname == "cbpAcctOutOctetCount" { return "Cbpacctoutoctetcount" }
    return ""
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetSegmentPath() string {
    return "cbpAcctEntry" + "[ifIndex='" + fmt.Sprintf("%v", cbpacctentry.Ifindex) + "']" + "[cbpAcctTrafficIndex='" + fmt.Sprintf("%v", cbpacctentry.Cbpaccttrafficindex) + "']"
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cbpacctentry.Ifindex
    leafs["cbpAcctTrafficIndex"] = cbpacctentry.Cbpaccttrafficindex
    leafs["cbpAcctInPacketCount"] = cbpacctentry.Cbpacctinpacketcount
    leafs["cbpAcctInOctetCount"] = cbpacctentry.Cbpacctinoctetcount
    leafs["cbpAcctOutPacketCount"] = cbpacctentry.Cbpacctoutpacketcount
    leafs["cbpAcctOutOctetCount"] = cbpacctentry.Cbpacctoutoctetcount
    return leafs
}

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetBundleName() string { return "cisco_ios_xe" }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetYangName() string { return "cbpAcctEntry" }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) SetParent(parent types.Entity) { cbpacctentry.parent = parent }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetParent() types.Entity { return cbpacctentry.parent }

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetParentYangName() string { return "cbpAcctTable" }

