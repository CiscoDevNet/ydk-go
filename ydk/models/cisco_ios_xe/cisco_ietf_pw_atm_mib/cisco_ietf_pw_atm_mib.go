// This MIB contains managed object definitions for Pseudo Wire
// emulation of ATM over Packet Switched Networks(PSN).
// 
// This MIB reports to the PW-MIB. The PW-MIB contains
// structures and MIB associations generic to Pseudo-Wire
// Virtual Circuit (VC) emulation. VC-specific MIBs (such as
// this) contain config and stats for specific VC types.
package cisco_ietf_pw_atm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_atm_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-ATM-MIB CISCO-IETF-PW-ATM-MIB}", reflect.TypeOf(CISCOIETFPWATMMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB", reflect.TypeOf(CISCOIETFPWATMMIB{}))
}

// CISCOIETFPWATMMIB
type CISCOIETFPWATMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table specifies the information for an ATM interface, VC, VP to be
    // carried over PSN.
    CpwVcAtmTable CISCOIETFPWATMMIB_CpwVcAtmTable
}

func (cISCOIETFPWATMMIB *CISCOIETFPWATMMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWATMMIB.EntityData.YFilter = cISCOIETFPWATMMIB.YFilter
    cISCOIETFPWATMMIB.EntityData.YangName = "CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWATMMIB.EntityData.ParentYangName = "CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.SegmentPath = "CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB"
    cISCOIETFPWATMMIB.EntityData.AbsolutePath = cISCOIETFPWATMMIB.EntityData.SegmentPath
    cISCOIETFPWATMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWATMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWATMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWATMMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFPWATMMIB.EntityData.Children.Append("cpwVcAtmTable", types.YChild{"CpwVcAtmTable", &cISCOIETFPWATMMIB.CpwVcAtmTable})
    cISCOIETFPWATMMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFPWATMMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFPWATMMIB.EntityData)
}

// CISCOIETFPWATMMIB_CpwVcAtmTable
// This table specifies the information for an ATM interface, VC,
// VP to be carried over PSN.
type CISCOIETFPWATMMIB_CpwVcAtmTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents an ATM interface, VC, VP that needs to be
    // adapted and carried over PSN. This table is indexed by CpwVcIndex in
    // CISCO-IETF-PW-MIB. The type is slice of
    // CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry.
    CpwVcAtmEntry []*CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry
}

func (cpwVcAtmTable *CISCOIETFPWATMMIB_CpwVcAtmTable) GetEntityData() *types.CommonEntityData {
    cpwVcAtmTable.EntityData.YFilter = cpwVcAtmTable.YFilter
    cpwVcAtmTable.EntityData.YangName = "cpwVcAtmTable"
    cpwVcAtmTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcAtmTable.EntityData.ParentYangName = "CISCO-IETF-PW-ATM-MIB"
    cpwVcAtmTable.EntityData.SegmentPath = "cpwVcAtmTable"
    cpwVcAtmTable.EntityData.AbsolutePath = "CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB/" + cpwVcAtmTable.EntityData.SegmentPath
    cpwVcAtmTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcAtmTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcAtmTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcAtmTable.EntityData.Children = types.NewOrderedMap()
    cpwVcAtmTable.EntityData.Children.Append("cpwVcAtmEntry", types.YChild{"CpwVcAtmEntry", nil})
    for i := range cpwVcAtmTable.CpwVcAtmEntry {
        cpwVcAtmTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcAtmTable.CpwVcAtmEntry[i]), types.YChild{"CpwVcAtmEntry", cpwVcAtmTable.CpwVcAtmEntry[i]})
    }
    cpwVcAtmTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcAtmTable.EntityData.YListKeys = []string {}

    return &(cpwVcAtmTable.EntityData)
}

// CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry
// A row in this table represents an ATM interface, VC, VP
// that needs to be adapted and carried over PSN. This table
// is indexed by CpwVcIndex in CISCO-IETF-PW-MIB.
type CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // The ATM Interface that receives cells from the ATM network. The type is
    // interface{} with range: 1..2147483647.
    CpwAtmIf interface{}

    // VPI value of this ATM VC. The type is interface{} with range: 0..4095.
    CpwAtmVpi interface{}

    // VCI value of this ATM VC. The type is interface{} with range: 0..65535.
    CpwAtmVci interface{}

    // This Object indicates whether the CLP bits are considered when determining
    // the value placed in the Quality of Service fields (e.g. EXP fields of the
    // MPLS Label Stack) of the encapsulating protocol. The type is bool.
    CpwAtmClpQosMapping interface{}

    // This Object is used to create, modify or delete a row in this table. The
    // type is RowStatus.
    CpwAtmRowStatus interface{}

    // This Object indicates whether OAM Cells are transported on this VC. The
    // type is bool.
    CpwAtmOamCellSupported interface{}

    // This Object represents the scaling factor (% value) to be applied to ATM
    // QoS rates when calculating QoS rates for the PSN domain . For example, in
    // the cell transport mode the bandwidth needed in the PSN domain will be
    // higher (since PSN Transport header, PW header, and optional control word
    // have to transmitted with every cell), whereas in the AAL5 mode the
    // bandwidth needed in PSN domain will be less since cell headers will be
    // removed after reassembly. The type is interface{} with range:
    // -2147483648..2147483647.
    CpwAtmQosScalingFactor interface{}

    // This object is used to identify if the VC is configured to do Cell Packing.
    // The type is bool.
    CpwAtmCellPacking interface{}

    // This object indicates the maximum number of cells that get packed in one
    // packet. The type is interface{} with range: -2147483648..2147483647.
    CpwAtmMncp interface{}

    // This Object represents the maximum number of cell that can be packed in one
    // packet for peer interface. The type is interface{} with range:
    // -2147483648..2147483647.
    CpwAtmPeerMncp interface{}

    // This object indicates if the packet going on the pseudowire is mpls or
    // l2tpv3 encapsulated. The type is CpwAtmEncap.
    CpwAtmEncap interface{}

    // This Object represents which MCPT timeout value. The type is interface{}
    // with range: -2147483648..2147483647.
    CpwAtmMcptTimeout interface{}

    // This object can be used to obtain the information on the number of cells
    // that were received and sent to the PSN. The type is interface{} with range:
    // 0..4294967295.
    CpwAtmCellsReceived interface{}

    // This object can be used to obtain the information on the number of cells
    // that were received from the PSN and sent over the ATM network. The type is
    // interface{} with range: 0..4294967295.
    CpwAtmCellsSent interface{}

    // This Object indicates the number of cells that were rejected by this VC
    // because of policing. The type is interface{} with range: 0..4294967295.
    CpwAtmCellsRejected interface{}

    // This Object indicates the number of cells that were Tagged. The type is
    // interface{} with range: 0..4294967295.
    CpwAtmCellsTagged interface{}

    // High Capacity counter for the number of cells that were received by this
    // VC. The type is interface{} with range: 0..18446744073709551615.
    CpwAtmHCCellsReceived interface{}

    // High Capacity counter for the number of cells that were rejected by this VC
    // because of policing. The type is interface{} with range:
    // 0..18446744073709551615.
    CpwAtmHCCellsRejected interface{}

    // High Capacity counter for the number of cells that were tagged. The type is
    // interface{} with range: 0..18446744073709551615.
    CpwAtmHCCellsTagged interface{}

    // It indicates the Average number of cells that were received in one packet.
    // The type is interface{} with range: 0..4294967295.
    CpwAtmAvgCellsPacked interface{}

    // This object can be used to obtain the information on the number of packets
    // that were received and sent to the PSN. The type is interface{} with range:
    // 0..4294967295.
    CpwAtmPktsReceived interface{}

    // This object indicates the number of packets that were sent to the atm
    // network. The type is interface{} with range: 0..4294967295.
    CpwAtmPktsSent interface{}

    // This object indicates the number of packets that were rejected because of
    // Policing. The type is interface{} with range: 0..4294967295.
    CpwAtmPktsRejected interface{}
}

func (cpwVcAtmEntry *CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry) GetEntityData() *types.CommonEntityData {
    cpwVcAtmEntry.EntityData.YFilter = cpwVcAtmEntry.YFilter
    cpwVcAtmEntry.EntityData.YangName = "cpwVcAtmEntry"
    cpwVcAtmEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcAtmEntry.EntityData.ParentYangName = "cpwVcAtmTable"
    cpwVcAtmEntry.EntityData.SegmentPath = "cpwVcAtmEntry" + types.AddKeyToken(cpwVcAtmEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcAtmEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-ATM-MIB:CISCO-IETF-PW-ATM-MIB/cpwVcAtmTable/" + cpwVcAtmEntry.EntityData.SegmentPath
    cpwVcAtmEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcAtmEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcAtmEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcAtmEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcAtmEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcAtmEntry.CpwVcIndex})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmIf", types.YLeaf{"CpwAtmIf", cpwVcAtmEntry.CpwAtmIf})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmVpi", types.YLeaf{"CpwAtmVpi", cpwVcAtmEntry.CpwAtmVpi})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmVci", types.YLeaf{"CpwAtmVci", cpwVcAtmEntry.CpwAtmVci})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmClpQosMapping", types.YLeaf{"CpwAtmClpQosMapping", cpwVcAtmEntry.CpwAtmClpQosMapping})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmRowStatus", types.YLeaf{"CpwAtmRowStatus", cpwVcAtmEntry.CpwAtmRowStatus})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmOamCellSupported", types.YLeaf{"CpwAtmOamCellSupported", cpwVcAtmEntry.CpwAtmOamCellSupported})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmQosScalingFactor", types.YLeaf{"CpwAtmQosScalingFactor", cpwVcAtmEntry.CpwAtmQosScalingFactor})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmCellPacking", types.YLeaf{"CpwAtmCellPacking", cpwVcAtmEntry.CpwAtmCellPacking})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmMncp", types.YLeaf{"CpwAtmMncp", cpwVcAtmEntry.CpwAtmMncp})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmPeerMncp", types.YLeaf{"CpwAtmPeerMncp", cpwVcAtmEntry.CpwAtmPeerMncp})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmEncap", types.YLeaf{"CpwAtmEncap", cpwVcAtmEntry.CpwAtmEncap})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmMcptTimeout", types.YLeaf{"CpwAtmMcptTimeout", cpwVcAtmEntry.CpwAtmMcptTimeout})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmCellsReceived", types.YLeaf{"CpwAtmCellsReceived", cpwVcAtmEntry.CpwAtmCellsReceived})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmCellsSent", types.YLeaf{"CpwAtmCellsSent", cpwVcAtmEntry.CpwAtmCellsSent})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmCellsRejected", types.YLeaf{"CpwAtmCellsRejected", cpwVcAtmEntry.CpwAtmCellsRejected})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmCellsTagged", types.YLeaf{"CpwAtmCellsTagged", cpwVcAtmEntry.CpwAtmCellsTagged})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmHCCellsReceived", types.YLeaf{"CpwAtmHCCellsReceived", cpwVcAtmEntry.CpwAtmHCCellsReceived})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmHCCellsRejected", types.YLeaf{"CpwAtmHCCellsRejected", cpwVcAtmEntry.CpwAtmHCCellsRejected})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmHCCellsTagged", types.YLeaf{"CpwAtmHCCellsTagged", cpwVcAtmEntry.CpwAtmHCCellsTagged})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmAvgCellsPacked", types.YLeaf{"CpwAtmAvgCellsPacked", cpwVcAtmEntry.CpwAtmAvgCellsPacked})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmPktsReceived", types.YLeaf{"CpwAtmPktsReceived", cpwVcAtmEntry.CpwAtmPktsReceived})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmPktsSent", types.YLeaf{"CpwAtmPktsSent", cpwVcAtmEntry.CpwAtmPktsSent})
    cpwVcAtmEntry.EntityData.Leafs.Append("cpwAtmPktsRejected", types.YLeaf{"CpwAtmPktsRejected", cpwVcAtmEntry.CpwAtmPktsRejected})

    cpwVcAtmEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcAtmEntry.EntityData)
}

// CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap represents is mpls or l2tpv3 encapsulated.
type CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap string

const (
    CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap_mpls CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap = "mpls"

    CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap_l2tpv3 CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap = "l2tpv3"

    CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap_unknown CISCOIETFPWATMMIB_CpwVcAtmTable_CpwVcAtmEntry_CpwAtmEncap = "unknown"
)

