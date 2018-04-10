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
    Caqvccparamstable CISCOATMQOSMIB_Caqvccparamstable

    // This table is defined to provide QoS information for each active ATM VP
    // existing on the interface.
    Caqvpcparamstable CISCOATMQOSMIB_Caqvpcparamstable

    // This table provides queuing related information for a VC existing on an ATM
    // interface.
    Caqqueuingparamstable CISCOATMQOSMIB_Caqqueuingparamstable

    // This table provides queuing information for all  queuing classes
    // associating with a VC.
    Caqqueuingparamsclasstable CISCOATMQOSMIB_Caqqueuingparamsclasstable
}

func (cISCOATMQOSMIB *CISCOATMQOSMIB) GetEntityData() *types.CommonEntityData {
    cISCOATMQOSMIB.EntityData.YFilter = cISCOATMQOSMIB.YFilter
    cISCOATMQOSMIB.EntityData.YangName = "CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOATMQOSMIB.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.SegmentPath = "CISCO-ATM-QOS-MIB:CISCO-ATM-QOS-MIB"
    cISCOATMQOSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOATMQOSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOATMQOSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOATMQOSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOATMQOSMIB.EntityData.Children["caqVccParamsTable"] = types.YChild{"Caqvccparamstable", &cISCOATMQOSMIB.Caqvccparamstable}
    cISCOATMQOSMIB.EntityData.Children["caqVpcParamsTable"] = types.YChild{"Caqvpcparamstable", &cISCOATMQOSMIB.Caqvpcparamstable}
    cISCOATMQOSMIB.EntityData.Children["caqQueuingParamsTable"] = types.YChild{"Caqqueuingparamstable", &cISCOATMQOSMIB.Caqqueuingparamstable}
    cISCOATMQOSMIB.EntityData.Children["caqQueuingParamsClassTable"] = types.YChild{"Caqqueuingparamsclasstable", &cISCOATMQOSMIB.Caqqueuingparamsclasstable}
    cISCOATMQOSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOATMQOSMIB.EntityData)
}

// CISCOATMQOSMIB_Caqvccparamstable
// This table is defined to provide QoS information for
// each active ATM VC existing on the interface.
type CISCOATMQOSMIB_Caqvccparamstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVccEntry.
    // The type is slice of CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry.
    Caqvccparamsentry []CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry
}

func (caqvccparamstable *CISCOATMQOSMIB_Caqvccparamstable) GetEntityData() *types.CommonEntityData {
    caqvccparamstable.EntityData.YFilter = caqvccparamstable.YFilter
    caqvccparamstable.EntityData.YangName = "caqVccParamsTable"
    caqvccparamstable.EntityData.BundleName = "cisco_ios_xe"
    caqvccparamstable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqvccparamstable.EntityData.SegmentPath = "caqVccParamsTable"
    caqvccparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqvccparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqvccparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqvccparamstable.EntityData.Children = make(map[string]types.YChild)
    caqvccparamstable.EntityData.Children["caqVccParamsEntry"] = types.YChild{"Caqvccparamsentry", nil}
    for i := range caqvccparamstable.Caqvccparamsentry {
        caqvccparamstable.EntityData.Children[types.GetSegmentPath(&caqvccparamstable.Caqvccparamsentry[i])] = types.YChild{"Caqvccparamsentry", &caqvccparamstable.Caqvccparamsentry[i]}
    }
    caqvccparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(caqvccparamstable.EntityData)
}

// CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVccEntry.
type CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // The service category of this virtual circuit connection. The type is
    // AtmServiceCategory.
    Caqvccparamstype interface{}

    // Input Peak Cell Rate (PCR) in kbps with  Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrin0 interface{}

    // Number of OAM F5 end to end loopback cells sent through the VCC. The type
    // is interface{} with range: 0..4294967295.
    Caqvccparamspcrin01 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 0
    // (clp0). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrout0 interface{}

    // Output Peak Cell Rate (PCR) in kbps with Cell Loss Priority bit set to 1
    // (clp01). The type is interface{} with range: 0..4294967295.
    Caqvccparamspcrout01 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 0 (clp0). The type is interface{} with
    // range: 0..4294967295.
    Caqvccparamsscrin0 interface{}

    // Input Sustained Cell Rate (SCR) in kbps for connection with VBR type of QoS
    // and Cell Loss Priority bit set to 1 (clp01). The type is interface{} with
    // range: 0..4294967295.
    Caqvccparamsscrin01 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 0 (clp0). The type is interface{}
    // with range: 0..4294967295.
    Caqvccparamsscrout0 interface{}

    // Output Sustained Cell Rate (SCR) in kbps for connection with VBR type of
    // QoS and Cell Loss Priority bit set to 1 (clp01). The type is interface{}
    // with range: 0..4294967295.
    Caqvccparamsscrout01 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsin0 interface{}

    // Input Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsin01 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and  Cell
    // Loss Priority bit set to 0 (clp0). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsout0 interface{}

    // Output Burst Cell Size (BCS) for connection with VBR type of QoS and Cell
    // Loss Priority bit set to 1 (clp01). The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsbcsout01 interface{}

    // The source of configuration for peak cell rate. The type is
    // VcParamConfigLocation.
    Caqvccparamsinheritlevel interface{}

    // Input Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    Caqvccparamsmcrin interface{}

    // Output Minimum Cell Rate (MCR) in kbps for connection with VBR-nrt or ABR
    // type of QoS. The type is interface{} with range: 0..4294967295.
    Caqvccparamsmcrout interface{}

    // Inverse of rate decrease factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvrdf interface{}

    // Inverse of rate increase factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvrif interface{}

    // The source of configuration for rate factor. The type is
    // VcParamConfigLocation.
    Caqvccparamsrfl interface{}

    // Cell delay variation. The type is interface{} with range: 0..4294967295.
    Caqvccparamscdv interface{}

    // Cell delay variation tolerance. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamscdvt interface{}

    // Initial cell rate. The type is interface{} with range: 0..4294967295.
    Caqvccparamsicr interface{}

    // Transient buffer exposure. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamstbe interface{}

    // Fixed round-trip time. The type is interface{} with range: 0..4294967295.
    Caqvccparamsfrtt interface{}

    // Maximum number of tx cells for each forward rm cell. The type is
    // interface{} with range: 0..4294967295.
    Caqvccparamsnrm interface{}

    // Maximum time between forward rm cells. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvtrm interface{}

    // Inverse of cutoff decrease factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsinvcdf interface{}

    // Allowed cell rate decrease time factor. The type is interface{} with range:
    // 0..4294967295.
    Caqvccparamsadtf interface{}
}

func (caqvccparamsentry *CISCOATMQOSMIB_Caqvccparamstable_Caqvccparamsentry) GetEntityData() *types.CommonEntityData {
    caqvccparamsentry.EntityData.YFilter = caqvccparamsentry.YFilter
    caqvccparamsentry.EntityData.YangName = "caqVccParamsEntry"
    caqvccparamsentry.EntityData.BundleName = "cisco_ios_xe"
    caqvccparamsentry.EntityData.ParentYangName = "caqVccParamsTable"
    caqvccparamsentry.EntityData.SegmentPath = "caqVccParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqvccparamsentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqvccparamsentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqvccparamsentry.Atmvclvci) + "']"
    caqvccparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqvccparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqvccparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqvccparamsentry.EntityData.Children = make(map[string]types.YChild)
    caqvccparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    caqvccparamsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", caqvccparamsentry.Ifindex}
    caqvccparamsentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", caqvccparamsentry.Atmvclvpi}
    caqvccparamsentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", caqvccparamsentry.Atmvclvci}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsType"] = types.YLeaf{"Caqvccparamstype", caqvccparamsentry.Caqvccparamstype}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsPcrIn0"] = types.YLeaf{"Caqvccparamspcrin0", caqvccparamsentry.Caqvccparamspcrin0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsPcrIn01"] = types.YLeaf{"Caqvccparamspcrin01", caqvccparamsentry.Caqvccparamspcrin01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsPcrOut0"] = types.YLeaf{"Caqvccparamspcrout0", caqvccparamsentry.Caqvccparamspcrout0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsPcrOut01"] = types.YLeaf{"Caqvccparamspcrout01", caqvccparamsentry.Caqvccparamspcrout01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsScrIn0"] = types.YLeaf{"Caqvccparamsscrin0", caqvccparamsentry.Caqvccparamsscrin0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsScrIn01"] = types.YLeaf{"Caqvccparamsscrin01", caqvccparamsentry.Caqvccparamsscrin01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsScrOut0"] = types.YLeaf{"Caqvccparamsscrout0", caqvccparamsentry.Caqvccparamsscrout0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsScrOut01"] = types.YLeaf{"Caqvccparamsscrout01", caqvccparamsentry.Caqvccparamsscrout01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsBcsIn0"] = types.YLeaf{"Caqvccparamsbcsin0", caqvccparamsentry.Caqvccparamsbcsin0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsBcsIn01"] = types.YLeaf{"Caqvccparamsbcsin01", caqvccparamsentry.Caqvccparamsbcsin01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsBcsOut0"] = types.YLeaf{"Caqvccparamsbcsout0", caqvccparamsentry.Caqvccparamsbcsout0}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsBcsOut01"] = types.YLeaf{"Caqvccparamsbcsout01", caqvccparamsentry.Caqvccparamsbcsout01}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsInheritLevel"] = types.YLeaf{"Caqvccparamsinheritlevel", caqvccparamsentry.Caqvccparamsinheritlevel}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsMcrIn"] = types.YLeaf{"Caqvccparamsmcrin", caqvccparamsentry.Caqvccparamsmcrin}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsMcrOut"] = types.YLeaf{"Caqvccparamsmcrout", caqvccparamsentry.Caqvccparamsmcrout}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsInvRdf"] = types.YLeaf{"Caqvccparamsinvrdf", caqvccparamsentry.Caqvccparamsinvrdf}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsInvRif"] = types.YLeaf{"Caqvccparamsinvrif", caqvccparamsentry.Caqvccparamsinvrif}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsRfl"] = types.YLeaf{"Caqvccparamsrfl", caqvccparamsentry.Caqvccparamsrfl}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsCdv"] = types.YLeaf{"Caqvccparamscdv", caqvccparamsentry.Caqvccparamscdv}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsCdvt"] = types.YLeaf{"Caqvccparamscdvt", caqvccparamsentry.Caqvccparamscdvt}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsIcr"] = types.YLeaf{"Caqvccparamsicr", caqvccparamsentry.Caqvccparamsicr}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsTbe"] = types.YLeaf{"Caqvccparamstbe", caqvccparamsentry.Caqvccparamstbe}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsFrtt"] = types.YLeaf{"Caqvccparamsfrtt", caqvccparamsentry.Caqvccparamsfrtt}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsNrm"] = types.YLeaf{"Caqvccparamsnrm", caqvccparamsentry.Caqvccparamsnrm}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsInvTrm"] = types.YLeaf{"Caqvccparamsinvtrm", caqvccparamsentry.Caqvccparamsinvtrm}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsInvCdf"] = types.YLeaf{"Caqvccparamsinvcdf", caqvccparamsentry.Caqvccparamsinvcdf}
    caqvccparamsentry.EntityData.Leafs["caqVccParamsAdtf"] = types.YLeaf{"Caqvccparamsadtf", caqvccparamsentry.Caqvccparamsadtf}
    return &(caqvccparamsentry.EntityData)
}

// CISCOATMQOSMIB_Caqvpcparamstable
// This table is defined to provide QoS information for
// each active ATM VP existing on the interface.
type CISCOATMQOSMIB_Caqvpcparamstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the ATM QoS parameters provided by ciscoAtmQosVpcEntry.
    // The type is slice of CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry.
    Caqvpcparamsentry []CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry
}

func (caqvpcparamstable *CISCOATMQOSMIB_Caqvpcparamstable) GetEntityData() *types.CommonEntityData {
    caqvpcparamstable.EntityData.YFilter = caqvpcparamstable.YFilter
    caqvpcparamstable.EntityData.YangName = "caqVpcParamsTable"
    caqvpcparamstable.EntityData.BundleName = "cisco_ios_xe"
    caqvpcparamstable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqvpcparamstable.EntityData.SegmentPath = "caqVpcParamsTable"
    caqvpcparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqvpcparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqvpcparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqvpcparamstable.EntityData.Children = make(map[string]types.YChild)
    caqvpcparamstable.EntityData.Children["caqVpcParamsEntry"] = types.YChild{"Caqvpcparamsentry", nil}
    for i := range caqvpcparamstable.Caqvpcparamsentry {
        caqvpcparamstable.EntityData.Children[types.GetSegmentPath(&caqvpcparamstable.Caqvpcparamsentry[i])] = types.YChild{"Caqvpcparamsentry", &caqvpcparamstable.Caqvpcparamsentry[i]}
    }
    caqvpcparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(caqvpcparamstable.EntityData)
}

// CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry
// This list contains the ATM QoS parameters provided by
// ciscoAtmQosVpcEntry.
type CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvpltable_Atmvplentry_Atmvplvpi
    Atmvplvpi interface{}

    // VP state of the current permanent virtual path. The type is VpState.
    Caqvpcparamsvpstate interface{}

    // Maximum rate in kbps at which the associated permanent virtual path can
    // transmit data. The type is interface{} with range: 0..4294967295.
    Caqvpcparamspeakrate interface{}

    // Maximum rate in kbps at which CES VCs can transmit data with the associated
    // permanent virtual path. The type is interface{} with range: 0..4294967295.
    Caqvpcparamscesrate interface{}

    // Number of data VCs currently associated with the permanent virtual path.
    // The type is interface{} with range: 0..65535.
    Caqvpcparamsdatavccount interface{}

    // Number of CES VCs currently associated with the permanent virtual path. The
    // type is interface{} with range: 0..65535.
    Caqvpcparamscesvccount interface{}

    // Vcd for F4 OAM segment processing. The type is interface{} with range:
    // 0..65535.
    Caqvpcparamsvcdf4Seg interface{}

    // Vcd for F4 OAM end to end processing. The type is interface{} with range:
    // 0..65535.
    Caqvpcparamsvcdf4Ete interface{}

    // Sustained cell rate associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    Caqvpcparamsscr interface{}

    // Maximum burst size associated with the PVP. The type is interface{} with
    // range: 0..4294967295.
    Caqvpcparamsmbs interface{}

    // Bandwidth in Kbps currently currently available on this PVP. The type is
    // interface{} with range: 0..4294967295.
    Caqvpcparamsavailbw interface{}
}

func (caqvpcparamsentry *CISCOATMQOSMIB_Caqvpcparamstable_Caqvpcparamsentry) GetEntityData() *types.CommonEntityData {
    caqvpcparamsentry.EntityData.YFilter = caqvpcparamsentry.YFilter
    caqvpcparamsentry.EntityData.YangName = "caqVpcParamsEntry"
    caqvpcparamsentry.EntityData.BundleName = "cisco_ios_xe"
    caqvpcparamsentry.EntityData.ParentYangName = "caqVpcParamsTable"
    caqvpcparamsentry.EntityData.SegmentPath = "caqVpcParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqvpcparamsentry.Ifindex) + "']" + "[atmVplVpi='" + fmt.Sprintf("%v", caqvpcparamsentry.Atmvplvpi) + "']"
    caqvpcparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqvpcparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqvpcparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqvpcparamsentry.EntityData.Children = make(map[string]types.YChild)
    caqvpcparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    caqvpcparamsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", caqvpcparamsentry.Ifindex}
    caqvpcparamsentry.EntityData.Leafs["atmVplVpi"] = types.YLeaf{"Atmvplvpi", caqvpcparamsentry.Atmvplvpi}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsVpState"] = types.YLeaf{"Caqvpcparamsvpstate", caqvpcparamsentry.Caqvpcparamsvpstate}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsPeakRate"] = types.YLeaf{"Caqvpcparamspeakrate", caqvpcparamsentry.Caqvpcparamspeakrate}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsCesRate"] = types.YLeaf{"Caqvpcparamscesrate", caqvpcparamsentry.Caqvpcparamscesrate}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsDataVcCount"] = types.YLeaf{"Caqvpcparamsdatavccount", caqvpcparamsentry.Caqvpcparamsdatavccount}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsCesVcCount"] = types.YLeaf{"Caqvpcparamscesvccount", caqvpcparamsentry.Caqvpcparamscesvccount}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsVcdF4Seg"] = types.YLeaf{"Caqvpcparamsvcdf4Seg", caqvpcparamsentry.Caqvpcparamsvcdf4Seg}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsVcdF4Ete"] = types.YLeaf{"Caqvpcparamsvcdf4Ete", caqvpcparamsentry.Caqvpcparamsvcdf4Ete}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsScr"] = types.YLeaf{"Caqvpcparamsscr", caqvpcparamsentry.Caqvpcparamsscr}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsMbs"] = types.YLeaf{"Caqvpcparamsmbs", caqvpcparamsentry.Caqvpcparamsmbs}
    caqvpcparamsentry.EntityData.Leafs["caqVpcParamsAvailBw"] = types.YLeaf{"Caqvpcparamsavailbw", caqvpcparamsentry.Caqvpcparamsavailbw}
    return &(caqvpcparamsentry.EntityData)
}

// CISCOATMQOSMIB_Caqqueuingparamstable
// This table provides queuing related information
// for a VC existing on an ATM interface.
type CISCOATMQOSMIB_Caqqueuingparamstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is defined as an entry in caqQueuingParamsTable. The type is slice of
    // CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry.
    Caqqueuingparamsentry []CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry
}

func (caqqueuingparamstable *CISCOATMQOSMIB_Caqqueuingparamstable) GetEntityData() *types.CommonEntityData {
    caqqueuingparamstable.EntityData.YFilter = caqqueuingparamstable.YFilter
    caqqueuingparamstable.EntityData.YangName = "caqQueuingParamsTable"
    caqqueuingparamstable.EntityData.BundleName = "cisco_ios_xe"
    caqqueuingparamstable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqqueuingparamstable.EntityData.SegmentPath = "caqQueuingParamsTable"
    caqqueuingparamstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqqueuingparamstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqqueuingparamstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqqueuingparamstable.EntityData.Children = make(map[string]types.YChild)
    caqqueuingparamstable.EntityData.Children["caqQueuingParamsEntry"] = types.YChild{"Caqqueuingparamsentry", nil}
    for i := range caqqueuingparamstable.Caqqueuingparamsentry {
        caqqueuingparamstable.EntityData.Children[types.GetSegmentPath(&caqqueuingparamstable.Caqqueuingparamsentry[i])] = types.YChild{"Caqqueuingparamsentry", &caqqueuingparamstable.Caqqueuingparamsentry[i]}
    }
    caqqueuingparamstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(caqqueuingparamstable.EntityData)
}

// CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry
// This is defined as an entry in caqQueuingParamsTable.
type CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // Mean Queue Depth associated with the vc.  This value is calculated based on
    // the actual queue depth on the interface and the exponential weighting
    // constant. The type is interface{} with range: 0..4294967295.
    Caqqueuingparamsmeanqdepth interface{}
}

func (caqqueuingparamsentry *CISCOATMQOSMIB_Caqqueuingparamstable_Caqqueuingparamsentry) GetEntityData() *types.CommonEntityData {
    caqqueuingparamsentry.EntityData.YFilter = caqqueuingparamsentry.YFilter
    caqqueuingparamsentry.EntityData.YangName = "caqQueuingParamsEntry"
    caqqueuingparamsentry.EntityData.BundleName = "cisco_ios_xe"
    caqqueuingparamsentry.EntityData.ParentYangName = "caqQueuingParamsTable"
    caqqueuingparamsentry.EntityData.SegmentPath = "caqQueuingParamsEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqqueuingparamsentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqqueuingparamsentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqqueuingparamsentry.Atmvclvci) + "']"
    caqqueuingparamsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqqueuingparamsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqqueuingparamsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqqueuingparamsentry.EntityData.Children = make(map[string]types.YChild)
    caqqueuingparamsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    caqqueuingparamsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", caqqueuingparamsentry.Ifindex}
    caqqueuingparamsentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", caqqueuingparamsentry.Atmvclvpi}
    caqqueuingparamsentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", caqqueuingparamsentry.Atmvclvci}
    caqqueuingparamsentry.EntityData.Leafs["caqQueuingParamsMeanQDepth"] = types.YLeaf{"Caqqueuingparamsmeanqdepth", caqqueuingparamsentry.Caqqueuingparamsmeanqdepth}
    return &(caqqueuingparamsentry.EntityData)
}

// CISCOATMQOSMIB_Caqqueuingparamsclasstable
// This table provides queuing information for all 
// queuing classes associating with a VC.
type CISCOATMQOSMIB_Caqqueuingparamsclasstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is defined as an entry in ciscoAtmQosVcQueuingClassTable to provide
    // queuing information of a specific class. The type is slice of
    // CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry.
    Caqqueuingparamsclassentry []CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry
}

func (caqqueuingparamsclasstable *CISCOATMQOSMIB_Caqqueuingparamsclasstable) GetEntityData() *types.CommonEntityData {
    caqqueuingparamsclasstable.EntityData.YFilter = caqqueuingparamsclasstable.YFilter
    caqqueuingparamsclasstable.EntityData.YangName = "caqQueuingParamsClassTable"
    caqqueuingparamsclasstable.EntityData.BundleName = "cisco_ios_xe"
    caqqueuingparamsclasstable.EntityData.ParentYangName = "CISCO-ATM-QOS-MIB"
    caqqueuingparamsclasstable.EntityData.SegmentPath = "caqQueuingParamsClassTable"
    caqqueuingparamsclasstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqqueuingparamsclasstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqqueuingparamsclasstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqqueuingparamsclasstable.EntityData.Children = make(map[string]types.YChild)
    caqqueuingparamsclasstable.EntityData.Children["caqQueuingParamsClassEntry"] = types.YChild{"Caqqueuingparamsclassentry", nil}
    for i := range caqqueuingparamsclasstable.Caqqueuingparamsclassentry {
        caqqueuingparamsclasstable.EntityData.Children[types.GetSegmentPath(&caqqueuingparamsclasstable.Caqqueuingparamsclassentry[i])] = types.YChild{"Caqqueuingparamsclassentry", &caqqueuingparamsclasstable.Caqqueuingparamsclassentry[i]}
    }
    caqqueuingparamsclasstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(caqqueuingparamsclasstable.EntityData)
}

// CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry
// This is defined as an entry in ciscoAtmQosVcQueuingClassTable
// to provide queuing information of a specific class.
type CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // This attribute is a key. A class index, which associates with an IP
    // precedence (0 to 8), is defined to reference individual
    // caqQueuingParamsClassEntry. The type is interface{} with range: 0..8.
    Caqqueuingparamsclassindex interface{}

    // Number of packets dropped when Mean Queue Length is between Minimum
    // Threshold and Maximum Threshold range. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassranddrp interface{}

    // Number of packets dropped because the Mean Queue Depth exceeds the Maximum
    // Threshold value. The type is interface{} with range: 0..4294967295.
    Caqqueuingparamsclasstaildrp interface{}

    // Minimum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassminthre interface{}

    // Maximum Threshold value in kbps. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassmaxthre interface{}

    // Mark probability denominator.  This is the value used in the calculation of
    // a packet being dropped when the average queue size is between the minimum
    // threshold and the maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    Caqqueuingparamsclassmrkprob interface{}
}

func (caqqueuingparamsclassentry *CISCOATMQOSMIB_Caqqueuingparamsclasstable_Caqqueuingparamsclassentry) GetEntityData() *types.CommonEntityData {
    caqqueuingparamsclassentry.EntityData.YFilter = caqqueuingparamsclassentry.YFilter
    caqqueuingparamsclassentry.EntityData.YangName = "caqQueuingParamsClassEntry"
    caqqueuingparamsclassentry.EntityData.BundleName = "cisco_ios_xe"
    caqqueuingparamsclassentry.EntityData.ParentYangName = "caqQueuingParamsClassTable"
    caqqueuingparamsclassentry.EntityData.SegmentPath = "caqQueuingParamsClassEntry" + "[ifIndex='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Atmvclvci) + "']" + "[caqQueuingParamsClassIndex='" + fmt.Sprintf("%v", caqqueuingparamsclassentry.Caqqueuingparamsclassindex) + "']"
    caqqueuingparamsclassentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    caqqueuingparamsclassentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    caqqueuingparamsclassentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    caqqueuingparamsclassentry.EntityData.Children = make(map[string]types.YChild)
    caqqueuingparamsclassentry.EntityData.Leafs = make(map[string]types.YLeaf)
    caqqueuingparamsclassentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", caqqueuingparamsclassentry.Ifindex}
    caqqueuingparamsclassentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", caqqueuingparamsclassentry.Atmvclvpi}
    caqqueuingparamsclassentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", caqqueuingparamsclassentry.Atmvclvci}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassIndex"] = types.YLeaf{"Caqqueuingparamsclassindex", caqqueuingparamsclassentry.Caqqueuingparamsclassindex}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassRandDrp"] = types.YLeaf{"Caqqueuingparamsclassranddrp", caqqueuingparamsclassentry.Caqqueuingparamsclassranddrp}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassTailDrp"] = types.YLeaf{"Caqqueuingparamsclasstaildrp", caqqueuingparamsclassentry.Caqqueuingparamsclasstaildrp}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassMinThre"] = types.YLeaf{"Caqqueuingparamsclassminthre", caqqueuingparamsclassentry.Caqqueuingparamsclassminthre}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassMaxThre"] = types.YLeaf{"Caqqueuingparamsclassmaxthre", caqqueuingparamsclassentry.Caqqueuingparamsclassmaxthre}
    caqqueuingparamsclassentry.EntityData.Leafs["caqQueuingParamsClassMrkProb"] = types.YLeaf{"Caqqueuingparamsclassmrkprob", caqqueuingparamsclassentry.Caqqueuingparamsclassmrkprob}
    return &(caqqueuingparamsclassentry.EntityData)
}

