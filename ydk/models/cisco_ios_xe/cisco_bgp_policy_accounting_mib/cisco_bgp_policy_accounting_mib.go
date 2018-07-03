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
    CbpAcctTable CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable
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

    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Children = types.NewOrderedMap()
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Children.Append("cbpAcctTable", types.YChild{"CbpAcctTable", &cISCOBGPPOLICYACCOUNTINGMIB.CbpAcctTable})
    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOBGPPOLICYACCOUNTINGMIB.EntityData.YListKeys = []string {}

    return &(cISCOBGPPOLICYACCOUNTINGMIB.EntityData)
}

// CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable
// The cbpAcctTable provides statistics about ingress and egress 
// traffic on an interface. This data could be used for purposes 
// like billing.
type CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable struct {
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
    // slice of CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable_CbpAcctEntry.
    CbpAcctEntry []*CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable_CbpAcctEntry
}

func (cbpAcctTable *CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable) GetEntityData() *types.CommonEntityData {
    cbpAcctTable.EntityData.YFilter = cbpAcctTable.YFilter
    cbpAcctTable.EntityData.YangName = "cbpAcctTable"
    cbpAcctTable.EntityData.BundleName = "cisco_ios_xe"
    cbpAcctTable.EntityData.ParentYangName = "CISCO-BGP-POLICY-ACCOUNTING-MIB"
    cbpAcctTable.EntityData.SegmentPath = "cbpAcctTable"
    cbpAcctTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbpAcctTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbpAcctTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbpAcctTable.EntityData.Children = types.NewOrderedMap()
    cbpAcctTable.EntityData.Children.Append("cbpAcctEntry", types.YChild{"CbpAcctEntry", nil})
    for i := range cbpAcctTable.CbpAcctEntry {
        cbpAcctTable.EntityData.Children.Append(types.GetSegmentPath(cbpAcctTable.CbpAcctEntry[i]), types.YChild{"CbpAcctEntry", cbpAcctTable.CbpAcctEntry[i]})
    }
    cbpAcctTable.EntityData.Leafs = types.NewOrderedMap()

    cbpAcctTable.EntityData.YListKeys = []string {}

    return &(cbpAcctTable.EntityData)
}

// CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable_CbpAcctEntry
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
type CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable_CbpAcctEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An integer value greater than 0, that uniquely
    // identifies a traffic-type. The traffic-type has no intrinsic meaning. It
    // just means the traffic coming into an interface can be differentiated into
    // different types. It is up to the user to give meaning to and configure the
    // various traffic-types on an  interface. The type is interface{} with range:
    // 1..2147483647.
    CbpAcctTrafficIndex interface{}

    // The total number of packets received for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    CbpAcctInPacketCount interface{}

    // The total number of octets received for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    CbpAcctInOctetCount interface{}

    // The total number of packets transmitted for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    CbpAcctOutPacketCount interface{}

    // The total number of octets transmitted for a particular traffic-type on an
    // interface. The type is interface{} with range: 0..18446744073709551615.
    CbpAcctOutOctetCount interface{}
}

func (cbpAcctEntry *CISCOBGPPOLICYACCOUNTINGMIB_CbpAcctTable_CbpAcctEntry) GetEntityData() *types.CommonEntityData {
    cbpAcctEntry.EntityData.YFilter = cbpAcctEntry.YFilter
    cbpAcctEntry.EntityData.YangName = "cbpAcctEntry"
    cbpAcctEntry.EntityData.BundleName = "cisco_ios_xe"
    cbpAcctEntry.EntityData.ParentYangName = "cbpAcctTable"
    cbpAcctEntry.EntityData.SegmentPath = "cbpAcctEntry" + types.AddKeyToken(cbpAcctEntry.IfIndex, "ifIndex") + types.AddKeyToken(cbpAcctEntry.CbpAcctTrafficIndex, "cbpAcctTrafficIndex")
    cbpAcctEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cbpAcctEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cbpAcctEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cbpAcctEntry.EntityData.Children = types.NewOrderedMap()
    cbpAcctEntry.EntityData.Leafs = types.NewOrderedMap()
    cbpAcctEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cbpAcctEntry.IfIndex})
    cbpAcctEntry.EntityData.Leafs.Append("cbpAcctTrafficIndex", types.YLeaf{"CbpAcctTrafficIndex", cbpAcctEntry.CbpAcctTrafficIndex})
    cbpAcctEntry.EntityData.Leafs.Append("cbpAcctInPacketCount", types.YLeaf{"CbpAcctInPacketCount", cbpAcctEntry.CbpAcctInPacketCount})
    cbpAcctEntry.EntityData.Leafs.Append("cbpAcctInOctetCount", types.YLeaf{"CbpAcctInOctetCount", cbpAcctEntry.CbpAcctInOctetCount})
    cbpAcctEntry.EntityData.Leafs.Append("cbpAcctOutPacketCount", types.YLeaf{"CbpAcctOutPacketCount", cbpAcctEntry.CbpAcctOutPacketCount})
    cbpAcctEntry.EntityData.Leafs.Append("cbpAcctOutOctetCount", types.YLeaf{"CbpAcctOutOctetCount", cbpAcctEntry.CbpAcctOutOctetCount})

    cbpAcctEntry.EntityData.YListKeys = []string {"IfIndex", "CbpAcctTrafficIndex"}

    return &(cbpAcctEntry.EntityData)
}

