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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The cbpAcctTable provides statistics about ingress and egress  traffic on
    // an interface. This data could be used for purposes  like billing.
    Cbpaccttable CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable
}

func (cISCOBGPPOLICYACCOUNTINGMIB *CISCOBGPPOLICYACCOUNTINGMIB) GetEntityData() *types.CommonEntityData {
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.YFilter = cISCOBGPPOLICYACCOUNTINGMIB.YFilter
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.YangName = "CISCO-BGP-POLICY-ACCOUNTING-MIB"
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.ParentYangName = "CISCO-BGP-POLICY-ACCOUNTING-MIB"
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.SegmentPath = "CISCO-BGP-POLICY-ACCOUNTING-MIB:CISCO-BGP-POLICY-ACCOUNTING-MIB"
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Children["cbpAcctTable"] = types.YChild{"Cbpaccttable", &cISCOBGPPOLICYACCOUNTINGMIB.Cbpaccttable}
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOBGPPOLICYACCOUNTINGMIB.EntityData)
}

// CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable
// The cbpAcctTable provides statistics about ingress and egress 
// traffic on an interface. This data could be used for purposes 
// like billing.
type CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable struct {
    EntityData types.CommonEntityData
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

func (cbpaccttable *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable) GetEntityData() *types.CommonEntityData {
    cbpaccttable.EntityData.YFilter = cbpaccttable.YFilter
    cbpaccttable.EntityData.YangName = "cbpAcctTable"
    cbpaccttable.EntityData.BundleName = "cisco_ios_xe"
    cbpaccttable.EntityData.ParentYangName = "CISCO-BGP-POLICY-ACCOUNTING-MIB"
    cbpaccttable.EntityData.SegmentPath = "cbpAcctTable"
    cbpaccttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbpaccttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbpaccttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbpaccttable.EntityData.Children = make(map[string]types.YChild)
    cbpaccttable.EntityData.Children["cbpAcctEntry"] = types.YChild{"Cbpacctentry", nil}
    for i := range cbpaccttable.Cbpacctentry {
        cbpaccttable.EntityData.Children[types.GetSegmentPath(&cbpaccttable.Cbpacctentry[i])] = types.YChild{"Cbpacctentry", &cbpaccttable.Cbpacctentry[i]}
    }
    cbpaccttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cbpaccttable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cbpacctentry *CISCOBGPPOLICYACCOUNTINGMIB_Cbpaccttable_Cbpacctentry) GetEntityData() *types.CommonEntityData {
    cbpacctentry.EntityData.YFilter = cbpacctentry.YFilter
    cbpacctentry.EntityData.YangName = "cbpAcctEntry"
    cbpacctentry.EntityData.BundleName = "cisco_ios_xe"
    cbpacctentry.EntityData.ParentYangName = "cbpAcctTable"
    cbpacctentry.EntityData.SegmentPath = "cbpAcctEntry" + "[ifIndex='" + fmt.Sprintf("%v", cbpacctentry.Ifindex) + "']" + "[cbpAcctTrafficIndex='" + fmt.Sprintf("%v", cbpacctentry.Cbpaccttrafficindex) + "']"
    cbpacctentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbpacctentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbpacctentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbpacctentry.EntityData.Children = make(map[string]types.YChild)
    cbpacctentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cbpacctentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cbpacctentry.Ifindex}
    cbpacctentry.EntityData.Leafs["cbpAcctTrafficIndex"] = types.YLeaf{"Cbpaccttrafficindex", cbpacctentry.Cbpaccttrafficindex}
    cbpacctentry.EntityData.Leafs["cbpAcctInPacketCount"] = types.YLeaf{"Cbpacctinpacketcount", cbpacctentry.Cbpacctinpacketcount}
    cbpacctentry.EntityData.Leafs["cbpAcctInOctetCount"] = types.YLeaf{"Cbpacctinoctetcount", cbpacctentry.Cbpacctinoctetcount}
    cbpacctentry.EntityData.Leafs["cbpAcctOutPacketCount"] = types.YLeaf{"Cbpacctoutpacketcount", cbpacctentry.Cbpacctoutpacketcount}
    cbpacctentry.EntityData.Leafs["cbpAcctOutOctetCount"] = types.YLeaf{"Cbpacctoutoctetcount", cbpacctentry.Cbpacctoutoctetcount}
    return &(cbpacctentry.EntityData)
}

