// The MIB is created to provide ATM QoS information in
// these areas:
// 
// 1. Traffic shaping on a per-VC basis
// 2. Traffic shaping on a per-VP basis 
// 3. Per-VC queuing/buffering
// 
// Although the initial requirements of the MIB are
// driven to support the GSR TAZ line card,
// CISCO-ATM-QOS-MIB is designed as a generic MIB to
// support ATM interfaces across all platforms.
// 
// Here are the tables defined in this MIB:
// ciscoAtmQosVccTable 
//      - to provide information on traffic shaping on 
//        a per-VC basis.
//                       
// ciscoAtmQosVpcTable 
//      - to provide information on traffic shaping on
//        a per-VP basis.
//     
// ciscoAtmQosVcQueuingTable
// ciscoAtmQosVcQueuingClassTable
//      - to provide information on per-VC
//        queuing/buffering.
package cisco_atm_qos_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_atm_qos_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ATM-QOS-MIB CISCO-ATM-QOS-MIB}", reflect.TypeOf(CISCOATMQOSMIB{}))
    ydk.RegisterEntity("CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB", reflect.TypeOf(CISCOATMQOSMIB{}))
}

// VcParamConfigLocation represents 5 - VC-class configured on main-interface.
type VcParamConfigLocation string

const (
    VcParamConfigLocation_configDefault VcParamConfigLocation = "configDefault"

    VcParamConfigLocation_configVcDirect VcParamConfigLocation = "configVcDirect"

    VcParamConfigLocation_configVcClass VcParamConfigLocation = "configVcClass"

    VcParamConfigLocation_configVcClassSubInterface VcParamConfigLocation = "configVcClassSubInterface"

    VcParamConfigLocation_configVcClassInterface VcParamConfigLocation = "configVcClassInterface"
)

// VpState represents 2 - Active
type VpState string

const (
    VpState_vpStateInactive VpState = "vpStateInactive"

    VpState_vpStateActive VpState = "vpStateActive"
)

// CISCOATMQOSMIB
type CISCOATMQOSMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is defined to provide QoS information for each active ATM VC
    // existing on the interface.
    CaqVccParamsTable CISCOATMQOSMIB_CaqVccParamsTable

    // This table is defined to provide QoS information for each active ATM VP
    // existing on the interface.
    CaqVpcParamsTable CISCOATMQOSMIB_CaqVpcParamsTable

    // This table provides queuing related information for a VC existing on an ATM
    // interface.
    CaqQueuingParamsTable CISCOATMQOSMIB_CaqQueuingParamsTable

    // This table provides queuing information for all  queuing classes
    // associating with a VC.
    CaqQueuingParamsClassTable CISCOATMQOSMIB_CaqQueuingParamsClassTable
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetEntityData() *types.CommonEntityData {
    cISCOATMQOSMIB.EntityData.YFilter = cISCOATMQOSMIB.YFilter
    cISCOATMQOSMIB.EntityData.YangName = "CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOATMQOSMIB.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.SegmentPath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.AbsolutePath = cISCOATMQOSMIB.EntityData.SegmentPath
    cISCOATMQOSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOATMQOSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOATMQOSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOATMQOSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOATMQOSMIB.EntityData.Children.Append("caqVccParamsTable", types.YChild{"CaqVccParamsTable", &cISCOATMQOSMIB.CaqVccParamsTable})
    cISCOATMQOSMIB.EntityData.Children.Append("caqVpcParamsTable", types.YChild{"CaqVpcParamsTable", &cISCOATMQOSMIB.CaqVpcParamsTable})
    cISCOATMQOSMIB.EntityData.Children.Append("caqQueuingParamsTable", types.YChild{"CaqQueuingParamsTable", &cISCOATMQOSMIB.CaqQueuingParamsTable})
    cISCOATMQOSMIB.EntityData.Children.Append("caqQueuingParamsClassTable", types.YChild{"CaqQueuingParamsClassTable", &cISCOATMQOSMIB.CaqQueuingParamsClassTable})
    cISCOATMQOSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOATMQOSMIB.EntityData.YListKeys = []string {}

    return &(cISCOATMQOSMIB.EntityData)
}

// CISCOATMQOSMIB_CaqVccParamsTable
// This table is defined to provide QoS information for
// each active ATM VC existing on the interface.
type CISCOATMQOSMIB_CaqVccParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVccEntry.
    // The type is slice of CISCOATMQOSMIB_CaqVccParamsTable_CaqVccParamsEntry.
    CaqVccParamsEntry []*CISCOATMQOSMIB_CaqVccParamsTable_CaqVccParamsEntry
}

func (caqVccParamsTable *CISCOATMQOSMIB_CaqVccParamsTable) GetEntityData() *types.CommonEntityData {
    caqVccParamsTable.EntityData.YFilter = caqVccParamsTable.YFilter
    caqVccParamsTable.EntityData.YangName = "caqVccParamsTable"
    caqVccParamsTable.EntityData.BundleName = "cisco_ios_xe"
    caqVccParamsTable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqVccParamsTable.EntityData.SegmentPath = "caqVccParamsTable"
    caqVccParamsTable.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/" + caqVccParamsTable.EntityData.SegmentPath
    caqVccParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqVccParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqVccParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqVccParamsTable.EntityData.Children = types.NewOrderedMap()
    caqVccParamsTable.EntityData.Children.Append("caqVccParamsEntry", types.YChild{"CaqVccParamsEntry", nil})
    for i := range caqVccParamsTable.CaqVccParamsEntry {
        caqVccParamsTable.EntityData.Children.Append(types.GetSegmentPath(caqVccParamsTable.CaqVccParamsEntry[i]), types.YChild{"CaqVccParamsEntry", caqVccParamsTable.CaqVccParamsEntry[i]})
    }
    caqVccParamsTable.EntityData.Leafs = types.NewOrderedMap()

    caqVccParamsTable.EntityData.YListKeys = []string {}

    return &(caqVccParamsTable.EntityData)
}

// CISCOATMQOSMIB_CaqVccParamsTable_CaqVccParamsEntry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVccEntry.
type CISCOATMQOSMIB_CaqVccParamsTable_CaqVccParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // The service category of this virtual circuit connection. The type is
    // AtmServiceCategory.
    CaqVccParamsType interface{}

    // Input Peak Cell Rate (PCR) in kbps with  Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    CaqVccParamsPcrIn0 interface{}

    // Number of OAM F5 end to end loopback cells sent through the VCC. The type
    // is interface{} with range: 0..4294967295.
    CaqVccParamsPcrIn01 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    CaqVccParamsPcrOut0 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 1
    // (clp01). The type is interface{} with range: 0..4294967295.
    CaqVccParamsPcrOut01 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 0 (clp0). The type is interface{} with
    // range: 0..4294967295.
    CaqVccParamsScrIn0 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 1 (clp01). The type is interface{} with
    // range: 0..4294967295.
    CaqVccParamsScrIn01 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 0 (clp0). The type is interface{}
    // with range: 0..4294967295.
    CaqVccParamsScrOut0 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 1 (clp01). The type is interface{}
    // with range: 0..4294967295.
    CaqVccParamsScrOut01 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsBcsIn0 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsBcsIn01 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and  Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsBcsOut0 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsBcsOut01 interface{}

    // The source of configuration for peak cell rate. The type is
    // VcParamConfigLocation.
    CaqVccParamsInheritLevel interface{}

    // Input Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    CaqVccParamsMcrIn interface{}

    // Output Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    CaqVccParamsMcrOut interface{}

    // Inverse of rate decrease factor. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsInvRdf interface{}

    // Inverse of rate increase factor. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsInvRif interface{}

    // The source of configuration for rate factor. The type is
    // VcParamConfigLocation.
    CaqVccParamsRfl interface{}

    // Cell delay variation. The type is interface{} with range: 0..4294967295.
    CaqVccParamsCdv interface{}

    // Cell delay variation tolerance. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsCdvt interface{}

    // Initial cell rate. The type is interface{} with range: 0..4294967295.
    CaqVccParamsIcr interface{}

    // Transient buffer exposure. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsTbe interface{}

    // Fixed round-trip time. The type is interface{} with range: 0..4294967295.
    CaqVccParamsFrtt interface{}

    // Maximum number of tx cells for each forward rm cell. The type is
    // interface{} with range: 0..4294967295.
    CaqVccParamsNrm interface{}

    // Maximum time between forward rm cells. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsInvTrm interface{}

    // Inverse of cutoff decrease factor. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsInvCdf interface{}

    // Allowed cell rate decrease time factor. The type is interface{} with range:
    // 0..4294967295.
    CaqVccParamsAdtf interface{}
}

func (caqVccParamsEntry *CISCOATMQOSMIB_CaqVccParamsTable_CaqVccParamsEntry) GetEntityData() *types.CommonEntityData {
    caqVccParamsEntry.EntityData.YFilter = caqVccParamsEntry.YFilter
    caqVccParamsEntry.EntityData.YangName = "caqVccParamsEntry"
    caqVccParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    caqVccParamsEntry.EntityData.ParentYangName = "caqVccParamsTable"
    caqVccParamsEntry.EntityData.SegmentPath = "caqVccParamsEntry" + types.AddKeyToken(caqVccParamsEntry.IfIndex, "ifIndex") + types.AddKeyToken(caqVccParamsEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(caqVccParamsEntry.AtmVclVci, "atmVclVci")
    caqVccParamsEntry.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/caqVccParamsTable/" + caqVccParamsEntry.EntityData.SegmentPath
    caqVccParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqVccParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqVccParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqVccParamsEntry.EntityData.Children = types.NewOrderedMap()
    caqVccParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    caqVccParamsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", caqVccParamsEntry.IfIndex})
    caqVccParamsEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", caqVccParamsEntry.AtmVclVpi})
    caqVccParamsEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", caqVccParamsEntry.AtmVclVci})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsType", types.YLeaf{"CaqVccParamsType", caqVccParamsEntry.CaqVccParamsType})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsPcrIn0", types.YLeaf{"CaqVccParamsPcrIn0", caqVccParamsEntry.CaqVccParamsPcrIn0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsPcrIn01", types.YLeaf{"CaqVccParamsPcrIn01", caqVccParamsEntry.CaqVccParamsPcrIn01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsPcrOut0", types.YLeaf{"CaqVccParamsPcrOut0", caqVccParamsEntry.CaqVccParamsPcrOut0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsPcrOut01", types.YLeaf{"CaqVccParamsPcrOut01", caqVccParamsEntry.CaqVccParamsPcrOut01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsScrIn0", types.YLeaf{"CaqVccParamsScrIn0", caqVccParamsEntry.CaqVccParamsScrIn0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsScrIn01", types.YLeaf{"CaqVccParamsScrIn01", caqVccParamsEntry.CaqVccParamsScrIn01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsScrOut0", types.YLeaf{"CaqVccParamsScrOut0", caqVccParamsEntry.CaqVccParamsScrOut0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsScrOut01", types.YLeaf{"CaqVccParamsScrOut01", caqVccParamsEntry.CaqVccParamsScrOut01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsBcsIn0", types.YLeaf{"CaqVccParamsBcsIn0", caqVccParamsEntry.CaqVccParamsBcsIn0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsBcsIn01", types.YLeaf{"CaqVccParamsBcsIn01", caqVccParamsEntry.CaqVccParamsBcsIn01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsBcsOut0", types.YLeaf{"CaqVccParamsBcsOut0", caqVccParamsEntry.CaqVccParamsBcsOut0})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsBcsOut01", types.YLeaf{"CaqVccParamsBcsOut01", caqVccParamsEntry.CaqVccParamsBcsOut01})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsInheritLevel", types.YLeaf{"CaqVccParamsInheritLevel", caqVccParamsEntry.CaqVccParamsInheritLevel})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsMcrIn", types.YLeaf{"CaqVccParamsMcrIn", caqVccParamsEntry.CaqVccParamsMcrIn})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsMcrOut", types.YLeaf{"CaqVccParamsMcrOut", caqVccParamsEntry.CaqVccParamsMcrOut})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsInvRdf", types.YLeaf{"CaqVccParamsInvRdf", caqVccParamsEntry.CaqVccParamsInvRdf})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsInvRif", types.YLeaf{"CaqVccParamsInvRif", caqVccParamsEntry.CaqVccParamsInvRif})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsRfl", types.YLeaf{"CaqVccParamsRfl", caqVccParamsEntry.CaqVccParamsRfl})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsCdv", types.YLeaf{"CaqVccParamsCdv", caqVccParamsEntry.CaqVccParamsCdv})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsCdvt", types.YLeaf{"CaqVccParamsCdvt", caqVccParamsEntry.CaqVccParamsCdvt})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsIcr", types.YLeaf{"CaqVccParamsIcr", caqVccParamsEntry.CaqVccParamsIcr})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsTbe", types.YLeaf{"CaqVccParamsTbe", caqVccParamsEntry.CaqVccParamsTbe})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsFrtt", types.YLeaf{"CaqVccParamsFrtt", caqVccParamsEntry.CaqVccParamsFrtt})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsNrm", types.YLeaf{"CaqVccParamsNrm", caqVccParamsEntry.CaqVccParamsNrm})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsInvTrm", types.YLeaf{"CaqVccParamsInvTrm", caqVccParamsEntry.CaqVccParamsInvTrm})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsInvCdf", types.YLeaf{"CaqVccParamsInvCdf", caqVccParamsEntry.CaqVccParamsInvCdf})
    caqVccParamsEntry.EntityData.Leafs.Append("caqVccParamsAdtf", types.YLeaf{"CaqVccParamsAdtf", caqVccParamsEntry.CaqVccParamsAdtf})

    caqVccParamsEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(caqVccParamsEntry.EntityData)
}

// CISCOATMQOSMIB_CaqVpcParamsTable
// This table is defined to provide QoS information for
// each active ATM VP existing on the interface.
type CISCOATMQOSMIB_CaqVpcParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVpcEntry.
    // The type is slice of CISCOATMQOSMIB_CaqVpcParamsTable_CaqVpcParamsEntry.
    CaqVpcParamsEntry []*CISCOATMQOSMIB_CaqVpcParamsTable_CaqVpcParamsEntry
}

func (caqVpcParamsTable *CISCOATMQOSMIB_CaqVpcParamsTable) GetEntityData() *types.CommonEntityData {
    caqVpcParamsTable.EntityData.YFilter = caqVpcParamsTable.YFilter
    caqVpcParamsTable.EntityData.YangName = "caqVpcParamsTable"
    caqVpcParamsTable.EntityData.BundleName = "cisco_ios_xe"
    caqVpcParamsTable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqVpcParamsTable.EntityData.SegmentPath = "caqVpcParamsTable"
    caqVpcParamsTable.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/" + caqVpcParamsTable.EntityData.SegmentPath
    caqVpcParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqVpcParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqVpcParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqVpcParamsTable.EntityData.Children = types.NewOrderedMap()
    caqVpcParamsTable.EntityData.Children.Append("caqVpcParamsEntry", types.YChild{"CaqVpcParamsEntry", nil})
    for i := range caqVpcParamsTable.CaqVpcParamsEntry {
        caqVpcParamsTable.EntityData.Children.Append(types.GetSegmentPath(caqVpcParamsTable.CaqVpcParamsEntry[i]), types.YChild{"CaqVpcParamsEntry", caqVpcParamsTable.CaqVpcParamsEntry[i]})
    }
    caqVpcParamsTable.EntityData.Leafs = types.NewOrderedMap()

    caqVpcParamsTable.EntityData.YListKeys = []string {}

    return &(caqVpcParamsTable.EntityData)
}

// CISCOATMQOSMIB_CaqVpcParamsTable_CaqVpcParamsEntry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVpcEntry.
type CISCOATMQOSMIB_CaqVpcParamsTable_CaqVpcParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVplTable_AtmVplEntry_AtmVplVpi
    AtmVplVpi interface{}

    // VP state of the current permanent virtual path. The type is VpState.
    CaqVpcParamsVpState interface{}

    // Maximum rate in kbps at which the associated permanent virtual path can
    // transmit data. The type is interface{} with range: 0..4294967295.
    CaqVpcParamsPeakRate interface{}

    // Maximum rate in kbps at which CES VCs can transmit data with the associated
    // permanent virtual path. The type is interface{} with range: 0..4294967295.
    CaqVpcParamsCesRate interface{}

    // Number of data VCs currently associated with the permanent virtual path.
    // The type is interface{} with range: 0..65535.
    CaqVpcParamsDataVcCount interface{}

    // Number of CES VCs currently associated with the permanent virtual path. The
    // type is interface{} with range: 0..65535.
    CaqVpcParamsCesVcCount interface{}

    // Vcd for F4 OAM segment processing. The type is interface{} with range:
    // 0..65535.
    CaqVpcParamsVcdF4Seg interface{}

    // Vcd for F4 OAM end to end processing. The type is interface{} with range:
    // 0..65535.
    CaqVpcParamsVcdF4Ete interface{}

    // Sustained cell rate associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    CaqVpcParamsScr interface{}

    // Maximum burst size associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    CaqVpcParamsMbs interface{}

    // Bandwidth in Kbps currently currently available on this PVP. The type is
    // interface{} with range: 0..4294967295.
    CaqVpcParamsAvailBw interface{}
}

func (caqVpcParamsEntry *CISCOATMQOSMIB_CaqVpcParamsTable_CaqVpcParamsEntry) GetEntityData() *types.CommonEntityData {
    caqVpcParamsEntry.EntityData.YFilter = caqVpcParamsEntry.YFilter
    caqVpcParamsEntry.EntityData.YangName = "caqVpcParamsEntry"
    caqVpcParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    caqVpcParamsEntry.EntityData.ParentYangName = "caqVpcParamsTable"
    caqVpcParamsEntry.EntityData.SegmentPath = "caqVpcParamsEntry" + types.AddKeyToken(caqVpcParamsEntry.IfIndex, "ifIndex") + types.AddKeyToken(caqVpcParamsEntry.AtmVplVpi, "atmVplVpi")
    caqVpcParamsEntry.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/caqVpcParamsTable/" + caqVpcParamsEntry.EntityData.SegmentPath
    caqVpcParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqVpcParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqVpcParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqVpcParamsEntry.EntityData.Children = types.NewOrderedMap()
    caqVpcParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    caqVpcParamsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", caqVpcParamsEntry.IfIndex})
    caqVpcParamsEntry.EntityData.Leafs.Append("atmVplVpi", types.YLeaf{"AtmVplVpi", caqVpcParamsEntry.AtmVplVpi})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsVpState", types.YLeaf{"CaqVpcParamsVpState", caqVpcParamsEntry.CaqVpcParamsVpState})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsPeakRate", types.YLeaf{"CaqVpcParamsPeakRate", caqVpcParamsEntry.CaqVpcParamsPeakRate})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsCesRate", types.YLeaf{"CaqVpcParamsCesRate", caqVpcParamsEntry.CaqVpcParamsCesRate})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsDataVcCount", types.YLeaf{"CaqVpcParamsDataVcCount", caqVpcParamsEntry.CaqVpcParamsDataVcCount})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsCesVcCount", types.YLeaf{"CaqVpcParamsCesVcCount", caqVpcParamsEntry.CaqVpcParamsCesVcCount})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsVcdF4Seg", types.YLeaf{"CaqVpcParamsVcdF4Seg", caqVpcParamsEntry.CaqVpcParamsVcdF4Seg})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsVcdF4Ete", types.YLeaf{"CaqVpcParamsVcdF4Ete", caqVpcParamsEntry.CaqVpcParamsVcdF4Ete})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsScr", types.YLeaf{"CaqVpcParamsScr", caqVpcParamsEntry.CaqVpcParamsScr})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsMbs", types.YLeaf{"CaqVpcParamsMbs", caqVpcParamsEntry.CaqVpcParamsMbs})
    caqVpcParamsEntry.EntityData.Leafs.Append("caqVpcParamsAvailBw", types.YLeaf{"CaqVpcParamsAvailBw", caqVpcParamsEntry.CaqVpcParamsAvailBw})

    caqVpcParamsEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVplVpi"}

    return &(caqVpcParamsEntry.EntityData)
}

// CISCOATMQOSMIB_CaqQueuingParamsTable
// This table provides queuing related information
// for a VC existing on an ATM interface.
type CISCOATMQOSMIB_CaqQueuingParamsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is defined as an entry in caqQueuingParamsTable. The type is slice of
    // CISCOATMQOSMIB_CaqQueuingParamsTable_CaqQueuingParamsEntry.
    CaqQueuingParamsEntry []*CISCOATMQOSMIB_CaqQueuingParamsTable_CaqQueuingParamsEntry
}

func (caqQueuingParamsTable *CISCOATMQOSMIB_CaqQueuingParamsTable) GetEntityData() *types.CommonEntityData {
    caqQueuingParamsTable.EntityData.YFilter = caqQueuingParamsTable.YFilter
    caqQueuingParamsTable.EntityData.YangName = "caqQueuingParamsTable"
    caqQueuingParamsTable.EntityData.BundleName = "cisco_ios_xe"
    caqQueuingParamsTable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqQueuingParamsTable.EntityData.SegmentPath = "caqQueuingParamsTable"
    caqQueuingParamsTable.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/" + caqQueuingParamsTable.EntityData.SegmentPath
    caqQueuingParamsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqQueuingParamsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqQueuingParamsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqQueuingParamsTable.EntityData.Children = types.NewOrderedMap()
    caqQueuingParamsTable.EntityData.Children.Append("caqQueuingParamsEntry", types.YChild{"CaqQueuingParamsEntry", nil})
    for i := range caqQueuingParamsTable.CaqQueuingParamsEntry {
        caqQueuingParamsTable.EntityData.Children.Append(types.GetSegmentPath(caqQueuingParamsTable.CaqQueuingParamsEntry[i]), types.YChild{"CaqQueuingParamsEntry", caqQueuingParamsTable.CaqQueuingParamsEntry[i]})
    }
    caqQueuingParamsTable.EntityData.Leafs = types.NewOrderedMap()

    caqQueuingParamsTable.EntityData.YListKeys = []string {}

    return &(caqQueuingParamsTable.EntityData)
}

// CISCOATMQOSMIB_CaqQueuingParamsTable_CaqQueuingParamsEntry
// This is defined as an entry in caqQueuingParamsTable.
type CISCOATMQOSMIB_CaqQueuingParamsTable_CaqQueuingParamsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // Mean Queue Depth associated with the vc.  This value is calculated based on
    // the actual queue depth on the interface and the exponential weighting
    // constant. The type is interface{} with range: 0..4294967295.
    CaqQueuingParamsMeanQDepth interface{}
}

func (caqQueuingParamsEntry *CISCOATMQOSMIB_CaqQueuingParamsTable_CaqQueuingParamsEntry) GetEntityData() *types.CommonEntityData {
    caqQueuingParamsEntry.EntityData.YFilter = caqQueuingParamsEntry.YFilter
    caqQueuingParamsEntry.EntityData.YangName = "caqQueuingParamsEntry"
    caqQueuingParamsEntry.EntityData.BundleName = "cisco_ios_xe"
    caqQueuingParamsEntry.EntityData.ParentYangName = "caqQueuingParamsTable"
    caqQueuingParamsEntry.EntityData.SegmentPath = "caqQueuingParamsEntry" + types.AddKeyToken(caqQueuingParamsEntry.IfIndex, "ifIndex") + types.AddKeyToken(caqQueuingParamsEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(caqQueuingParamsEntry.AtmVclVci, "atmVclVci")
    caqQueuingParamsEntry.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/caqQueuingParamsTable/" + caqQueuingParamsEntry.EntityData.SegmentPath
    caqQueuingParamsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqQueuingParamsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqQueuingParamsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqQueuingParamsEntry.EntityData.Children = types.NewOrderedMap()
    caqQueuingParamsEntry.EntityData.Leafs = types.NewOrderedMap()
    caqQueuingParamsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", caqQueuingParamsEntry.IfIndex})
    caqQueuingParamsEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", caqQueuingParamsEntry.AtmVclVpi})
    caqQueuingParamsEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", caqQueuingParamsEntry.AtmVclVci})
    caqQueuingParamsEntry.EntityData.Leafs.Append("caqQueuingParamsMeanQDepth", types.YLeaf{"CaqQueuingParamsMeanQDepth", caqQueuingParamsEntry.CaqQueuingParamsMeanQDepth})

    caqQueuingParamsEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(caqQueuingParamsEntry.EntityData)
}

// CISCOATMQOSMIB_CaqQueuingParamsClassTable
// This table provides queuing information for all 
// queuing classes associating with a VC.
type CISCOATMQOSMIB_CaqQueuingParamsClassTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is defined as an entry in ciscoAtmQosVcQueuingClassTable to provide
    // queuing information of a specific class. The type is slice of
    // CISCOATMQOSMIB_CaqQueuingParamsClassTable_CaqQueuingParamsClassEntry.
    CaqQueuingParamsClassEntry []*CISCOATMQOSMIB_CaqQueuingParamsClassTable_CaqQueuingParamsClassEntry
}

func (caqQueuingParamsClassTable *CISCOATMQOSMIB_CaqQueuingParamsClassTable) GetEntityData() *types.CommonEntityData {
    caqQueuingParamsClassTable.EntityData.YFilter = caqQueuingParamsClassTable.YFilter
    caqQueuingParamsClassTable.EntityData.YangName = "caqQueuingParamsClassTable"
    caqQueuingParamsClassTable.EntityData.BundleName = "cisco_ios_xe"
    caqQueuingParamsClassTable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqQueuingParamsClassTable.EntityData.SegmentPath = "caqQueuingParamsClassTable"
    caqQueuingParamsClassTable.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/" + caqQueuingParamsClassTable.EntityData.SegmentPath
    caqQueuingParamsClassTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqQueuingParamsClassTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqQueuingParamsClassTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqQueuingParamsClassTable.EntityData.Children = types.NewOrderedMap()
    caqQueuingParamsClassTable.EntityData.Children.Append("caqQueuingParamsClassEntry", types.YChild{"CaqQueuingParamsClassEntry", nil})
    for i := range caqQueuingParamsClassTable.CaqQueuingParamsClassEntry {
        caqQueuingParamsClassTable.EntityData.Children.Append(types.GetSegmentPath(caqQueuingParamsClassTable.CaqQueuingParamsClassEntry[i]), types.YChild{"CaqQueuingParamsClassEntry", caqQueuingParamsClassTable.CaqQueuingParamsClassEntry[i]})
    }
    caqQueuingParamsClassTable.EntityData.Leafs = types.NewOrderedMap()

    caqQueuingParamsClassTable.EntityData.YListKeys = []string {}

    return &(caqQueuingParamsClassTable.EntityData)
}

// CISCOATMQOSMIB_CaqQueuingParamsClassTable_CaqQueuingParamsClassEntry
// This is defined as an entry in ciscoAtmQosVcQueuingClassTable
// to provide queuing information of a specific class.
type CISCOATMQOSMIB_CaqQueuingParamsClassTable_CaqQueuingParamsClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // This attribute is a key. A class index, which associates with an IP
    // precedence (0 to 8), is defined to reference individual
    // caqQueuingParamsClassEntry. The type is interface{} with range: 0..8.
    CaqQueuingParamsClassIndex interface{}

    // Number of packets dropped when Mean Queue Length is between Minimum
    // Threshold and Maximum Threshold range. The type is interface{} with range:
    // 0..4294967295.
    CaqQueuingParamsClassRandDrp interface{}

    // Number of packets dropped because the Mean Queue Depth exceeds the Maximum
    // Threshold value. The type is interface{} with range: 0..4294967295.
    CaqQueuingParamsClassTailDrp interface{}

    // Minimum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    CaqQueuingParamsClassMinThre interface{}

    // Maximum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    CaqQueuingParamsClassMaxThre interface{}

    // Mark probability denominator.  This is the value used in the calculation of
    // a packet being dropped when the average queue size is between the minimum
    // threshold and the maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    CaqQueuingParamsClassMrkProb interface{}
}

func (caqQueuingParamsClassEntry *CISCOATMQOSMIB_CaqQueuingParamsClassTable_CaqQueuingParamsClassEntry) GetEntityData() *types.CommonEntityData {
    caqQueuingParamsClassEntry.EntityData.YFilter = caqQueuingParamsClassEntry.YFilter
    caqQueuingParamsClassEntry.EntityData.YangName = "caqQueuingParamsClassEntry"
    caqQueuingParamsClassEntry.EntityData.BundleName = "cisco_ios_xe"
    caqQueuingParamsClassEntry.EntityData.ParentYangName = "caqQueuingParamsClassTable"
    caqQueuingParamsClassEntry.EntityData.SegmentPath = "caqQueuingParamsClassEntry" + types.AddKeyToken(caqQueuingParamsClassEntry.IfIndex, "ifIndex") + types.AddKeyToken(caqQueuingParamsClassEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(caqQueuingParamsClassEntry.AtmVclVci, "atmVclVci") + types.AddKeyToken(caqQueuingParamsClassEntry.CaqQueuingParamsClassIndex, "caqQueuingParamsClassIndex")
    caqQueuingParamsClassEntry.EntityData.AbsolutePath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB/caqQueuingParamsClassTable/" + caqQueuingParamsClassEntry.EntityData.SegmentPath
    caqQueuingParamsClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqQueuingParamsClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqQueuingParamsClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqQueuingParamsClassEntry.EntityData.Children = types.NewOrderedMap()
    caqQueuingParamsClassEntry.EntityData.Leafs = types.NewOrderedMap()
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", caqQueuingParamsClassEntry.IfIndex})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", caqQueuingParamsClassEntry.AtmVclVpi})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", caqQueuingParamsClassEntry.AtmVclVci})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassIndex", types.YLeaf{"CaqQueuingParamsClassIndex", caqQueuingParamsClassEntry.CaqQueuingParamsClassIndex})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassRandDrp", types.YLeaf{"CaqQueuingParamsClassRandDrp", caqQueuingParamsClassEntry.CaqQueuingParamsClassRandDrp})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassTailDrp", types.YLeaf{"CaqQueuingParamsClassTailDrp", caqQueuingParamsClassEntry.CaqQueuingParamsClassTailDrp})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassMinThre", types.YLeaf{"CaqQueuingParamsClassMinThre", caqQueuingParamsClassEntry.CaqQueuingParamsClassMinThre})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassMaxThre", types.YLeaf{"CaqQueuingParamsClassMaxThre", caqQueuingParamsClassEntry.CaqQueuingParamsClassMaxThre})
    caqQueuingParamsClassEntry.EntityData.Leafs.Append("caqQueuingParamsClassMrkProb", types.YLeaf{"CaqQueuingParamsClassMrkProb", caqQueuingParamsClassEntry.CaqQueuingParamsClassMrkProb})

    caqQueuingParamsClassEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci", "CaqQueuingParamsClassIndex"}

    return &(caqQueuingParamsClassEntry.EntityData)
}

