// This MIB describes a model for managing Ethernet  
// point-to-point pseudo wire services over a Packet  
// Switched Network (PSN).
package cisco_ietf_pw_enet_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_enet_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-ENET-MIB CISCO-IETF-PW-ENET-MIB}", reflect.TypeOf(CISCOIETFPWENETMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB", reflect.TypeOf(CISCOIETFPWENETMIB{}))
}

// CISCOIETFPWENETMIB
type CISCOIETFPWENETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table contains the index to the Ethernet tables   associated with this
    // ETH VC, the VLAN configuration and   VLAN mode.
    Cpwvcenettable CISCOIETFPWENETMIB_Cpwvcenettable

    // This table may be used for MPLS PSNs if there is a need   to hold multiple
    // VC, each with different COS, for the same  user service (port + PW VLAN).
    // Such a need may arise if the  MPLS network is capable of L-LSP or E-LSP
    // without multiple  COS capabilities.  Each row is indexed by the cpwVcIndex 
    // and indicate the PRI bits on the packet recieved from the   user port (or
    // VPLS virtual port) that are  classified to this VC. Note that the EXP bit
    // value of the VC  is configured in the CISCO-IETF-PW-MPLS-MIB.
    Cpwvcenetmplsprimappingtable CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable

    // This table contains statistical counters specific for   Ethernet PW.
    Cpwvcenetstatstable CISCOIETFPWENETMIB_Cpwvcenetstatstable
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWENETMIB.EntityData.YFilter = cISCOIETFPWENETMIB.YFilter
    cISCOIETFPWENETMIB.EntityData.YangName = "CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWENETMIB.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.SegmentPath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWENETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWENETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWENETMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFPWENETMIB.EntityData.Children["cpwVcEnetTable"] = types.YChild{"Cpwvcenettable", &cISCOIETFPWENETMIB.Cpwvcenettable}
    cISCOIETFPWENETMIB.EntityData.Children["cpwVcEnetMplsPriMappingTable"] = types.YChild{"Cpwvcenetmplsprimappingtable", &cISCOIETFPWENETMIB.Cpwvcenetmplsprimappingtable}
    cISCOIETFPWENETMIB.EntityData.Children["cpwVcEnetStatsTable"] = types.YChild{"Cpwvcenetstatstable", &cISCOIETFPWENETMIB.Cpwvcenetstatstable}
    cISCOIETFPWENETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFPWENETMIB.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenettable
// This table contains the index to the Ethernet tables  
// associated with this ETH VC, the VLAN configuration and  
// VLAN mode.
type CISCOIETFPWENETMIB_Cpwvcenettable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table is indexed by the same index that was created   for the
    // associated entry in the PW VC Table in the  CISCO-IETF-PW-MIB.  The
    // CpwVcIndex and the cpwVcEnetPwVlan  are used as indexes to allow multiple
    // VLANs to exist on  the same PW.   An entry is created in this table by the
    // agent for every   entry in the cpwVc table with a VcType of 'ethernetVLAN',
    // 'ethernet' or 'ethernetVPLS'. Additional rows may be   created by the
    // operator or the agent if multiple entries  are required for the same VC.  
    // This table provides Ethernet port mapping and VLAN   configuration for each
    // Ethernet VC. The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry.
    Cpwvcenetentry []CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry
}

func (cpwvcenettable *CISCOIETFPWENETMIB_Cpwvcenettable) GetEntityData() *types.CommonEntityData {
    cpwvcenettable.EntityData.YFilter = cpwvcenettable.YFilter
    cpwvcenettable.EntityData.YangName = "cpwVcEnetTable"
    cpwvcenettable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenettable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwvcenettable.EntityData.SegmentPath = "cpwVcEnetTable"
    cpwvcenettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenettable.EntityData.Children = make(map[string]types.YChild)
    cpwvcenettable.EntityData.Children["cpwVcEnetEntry"] = types.YChild{"Cpwvcenetentry", nil}
    for i := range cpwvcenettable.Cpwvcenetentry {
        cpwvcenettable.EntityData.Children[types.GetSegmentPath(&cpwvcenettable.Cpwvcenetentry[i])] = types.YChild{"Cpwvcenetentry", &cpwvcenettable.Cpwvcenetentry[i]}
    }
    cpwvcenettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcenettable.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry
// This table is indexed by the same index that was created  
// for the associated entry in the PW VC Table in the 
// CISCO-IETF-PW-MIB.  The CpwVcIndex and the cpwVcEnetPwVlan 
// are used as indexes to allow multiple VLANs to exist on 
// the same PW. 
// 
// An entry is created in this table by the agent for every  
// entry in the cpwVc table with a VcType of 'ethernetVLAN', 
// 'ethernet' or 'ethernetVPLS'. Additional rows may be  
// created by the operator or the agent if multiple entries 
// are required for the same VC. 
// 
// This table provides Ethernet port mapping and VLAN  
// configuration for each Ethernet VC.
type CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. This Object defines the VLAN on the VC. The value
    // of 4097  is used if the object is not applicable, for example when  mapping
    // all packets from an Ethernet port to this VC.  The value of 4096 is used to
    // indicate untagged frames (at   least from the PW point of view), for
    // example if   cpwVcEnetVlanMode is equal 'removeVLAN' or when  
    // cpwVcEnetVlanMode equal 'noChange' and cpwVcEnetPortVlan  is equal 4096.
    // The type is interface{} with range: 0..4097.
    Cpwvcenetpwvlan interface{}

    // Indicate the mode of VLAN handling between the port   associated to the VC
    // and the VC encapsulation itself.   - 'other' indicate operation that is not
    // defined by    this MIB.   - 'portBased' indicates that the forwarder will
    // forward    packets between the port and the PW independent of their   
    // structure.   - 'noChange' indicates that the VC contains the original    
    // user VLAN, as specified in cpwVcEnetPortVlan.   - 'changeVlan' indicates
    // that the VLAN field on the VC     may be different than the VLAN field on
    // the user's     port.   - 'removeVlan' indicates that the encapsulation on
    // the     VC does not include the original VLAN field. Note     that PRI bits
    // transparency is lost in this case.   - 'addVlan' indicate that a VLAN field
    // will be added    on the PSN bound direction. cpwVcEnetPwVlan indicate   
    // the value that will be added.    - 'removeVlan', 'addVlan' and 'changeVlan'
    // implementation    is not required. . The type is Cpwvcenetvlanmode.
    Cpwvcenetvlanmode interface{}

    // This object define the VLAN value on the physical port (or   VPLS virtual
    // port) if a change is required to the VLAN value  between the VC and the
    // physical/virtual port.   The value of this object can be ignored if the
    // whole traffic   from the port is forwarded to one VC independent of the  
    // tagging on the port, but it is RECOMENDED that the value in  this case will
    // be '4097' indicating not relevant.   It MUST be equal to cpwVcEnetPwVlan if
    // 'noChange' mode  is used.   The value 4096 indicate that no VLAN (i.e.
    // untagged   frames) on the port are associated to this VC. This   allows the
    // same behaviors as assigning 'Default VLAN'   to un-tagged frames. . The
    // type is interface{} with range: 0..4097.
    Cpwvcenetportvlan interface{}

    // It is sometimes convenient to model the VC PW as a   virtual interface in
    // the ifTable. In these cases this   object hold the value of the ifIndex in
    // the ifTable   representing this VC PW. A value of zero indicate no such  
    // association or association is not yet known. The type is interface{} with
    // range: 0..2147483647.
    Cpwvcenetvcifindex interface{}

    // This object is used to specify the ifIndex of the ETHERNET  port associated
    // with this VC for point-to-point Ethernet   service, or the ifIndex of the
    // virtual interface of the VPLS   instance associated with the PW if the
    // service is VPLS. Two   rows in this table can point to the same ifIndex
    // only if:   1) It is required to support multiple COS on a MPLS PSN      for
    // the same service (i.e.: a combination of ports and      VLANs) by the use
    // of multiple VC, each with a different     COS.   2) There is no overlap of
    // VLAN values specified in      cpwVcEnetPortVlan that are associated with
    // this port.   A value of zero indicate that association to an ifIndex is 
    // not yet known. The type is interface{} with range: 0..2147483647.
    Cpwvcenetportifindex interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    Cpwvcenetrowstatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    Cpwvcenetstoragetype interface{}
}

func (cpwvcenetentry *CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry) GetEntityData() *types.CommonEntityData {
    cpwvcenetentry.EntityData.YFilter = cpwvcenetentry.YFilter
    cpwvcenetentry.EntityData.YangName = "cpwVcEnetEntry"
    cpwvcenetentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenetentry.EntityData.ParentYangName = "cpwVcEnetTable"
    cpwvcenetentry.EntityData.SegmentPath = "cpwVcEnetEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetentry.Cpwvcindex) + "']" + "[cpwVcEnetPwVlan='" + fmt.Sprintf("%v", cpwvcenetentry.Cpwvcenetpwvlan) + "']"
    cpwvcenetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenetentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcenetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcenetentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcenetentry.Cpwvcindex}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetPwVlan"] = types.YLeaf{"Cpwvcenetpwvlan", cpwvcenetentry.Cpwvcenetpwvlan}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetVlanMode"] = types.YLeaf{"Cpwvcenetvlanmode", cpwvcenetentry.Cpwvcenetvlanmode}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetPortVlan"] = types.YLeaf{"Cpwvcenetportvlan", cpwvcenetentry.Cpwvcenetportvlan}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetVcIfIndex"] = types.YLeaf{"Cpwvcenetvcifindex", cpwvcenetentry.Cpwvcenetvcifindex}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetPortIfIndex"] = types.YLeaf{"Cpwvcenetportifindex", cpwvcenetentry.Cpwvcenetportifindex}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetRowStatus"] = types.YLeaf{"Cpwvcenetrowstatus", cpwvcenetentry.Cpwvcenetrowstatus}
    cpwvcenetentry.EntityData.Leafs["cpwVcEnetStorageType"] = types.YLeaf{"Cpwvcenetstoragetype", cpwvcenetentry.Cpwvcenetstoragetype}
    return &(cpwvcenetentry.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode represents   is not required. 
type CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode string

const (
    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_other CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "other"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_portBased CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "portBased"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_noChange CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "noChange"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_changeVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "changeVlan"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_addVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "addVlan"

    CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode_removeVlan CISCOIETFPWENETMIB_Cpwvcenettable_Cpwvcenetentry_Cpwvcenetvlanmode = "removeVlan"
)

// CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable
// This table may be used for MPLS PSNs if there is a need  
// to hold multiple VC, each with different COS, for the same 
// user service (port + PW VLAN). Such a need may arise if the 
// MPLS network is capable of L-LSP or E-LSP without multiple 
// COS capabilities.  Each row is indexed by the cpwVcIndex  
// and indicate the PRI bits on the packet recieved from the  
// user port (or VPLS virtual port) that are 
// classified to this VC. Note that the EXP bit value of the VC 
// is configured in the CISCO-IETF-PW-MPLS-MIB.
type CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry is created if special classification based on   the PRI bits is
    // required for this VC. The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry.
    Cpwvcenetmplsprimappingtableentry []CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry
}

func (cpwvcenetmplsprimappingtable *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable) GetEntityData() *types.CommonEntityData {
    cpwvcenetmplsprimappingtable.EntityData.YFilter = cpwvcenetmplsprimappingtable.YFilter
    cpwvcenetmplsprimappingtable.EntityData.YangName = "cpwVcEnetMplsPriMappingTable"
    cpwvcenetmplsprimappingtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenetmplsprimappingtable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwvcenetmplsprimappingtable.EntityData.SegmentPath = "cpwVcEnetMplsPriMappingTable"
    cpwvcenetmplsprimappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenetmplsprimappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenetmplsprimappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenetmplsprimappingtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcenetmplsprimappingtable.EntityData.Children["cpwVcEnetMplsPriMappingTableEntry"] = types.YChild{"Cpwvcenetmplsprimappingtableentry", nil}
    for i := range cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry {
        cpwvcenetmplsprimappingtable.EntityData.Children[types.GetSegmentPath(&cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry[i])] = types.YChild{"Cpwvcenetmplsprimappingtableentry", &cpwvcenetmplsprimappingtable.Cpwvcenetmplsprimappingtableentry[i]}
    }
    cpwvcenetmplsprimappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcenetmplsprimappingtable.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry
// Each entry is created if special classification based on  
// the PRI bits is required for this VC.
type CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This object defines the groups of user PRI mapped into  this VC. Each bit
    // set indicates that this user priority   is assigned to this VC.   The value
    // 'untagged' is used to indicate that untagged   frames are also associated
    // to this VC.   This object allow the use of different PSN COS based on  
    // user marking of PRI bits in MPLS PSN with L-LSP or   E-LSP without multiple
    // COS support. In all other cases,   the default value MUST be used.   It is
    // REQUIRED that there is no overlap on this object   between rows serving the
    // same service (port+ PW VLAN).   In case of missing BIT configuration
    // between rows to   the same service, incoming packets with PRI marking not  
    // configured should be handled by the VC with the lowest   COS. . The type is
    // map[string]bool.
    Cpwvcenetmplsprimapping interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    Cpwvcenetmplsprimappingrowstatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    Cpwvcenetmplsprimappingstoragetype interface{}
}

func (cpwvcenetmplsprimappingtableentry *CISCOIETFPWENETMIB_Cpwvcenetmplsprimappingtable_Cpwvcenetmplsprimappingtableentry) GetEntityData() *types.CommonEntityData {
    cpwvcenetmplsprimappingtableentry.EntityData.YFilter = cpwvcenetmplsprimappingtableentry.YFilter
    cpwvcenetmplsprimappingtableentry.EntityData.YangName = "cpwVcEnetMplsPriMappingTableEntry"
    cpwvcenetmplsprimappingtableentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenetmplsprimappingtableentry.EntityData.ParentYangName = "cpwVcEnetMplsPriMappingTable"
    cpwvcenetmplsprimappingtableentry.EntityData.SegmentPath = "cpwVcEnetMplsPriMappingTableEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetmplsprimappingtableentry.Cpwvcindex) + "']"
    cpwvcenetmplsprimappingtableentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenetmplsprimappingtableentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenetmplsprimappingtableentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenetmplsprimappingtableentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcenetmplsprimappingtableentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcenetmplsprimappingtableentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcenetmplsprimappingtableentry.Cpwvcindex}
    cpwvcenetmplsprimappingtableentry.EntityData.Leafs["cpwVcEnetMplsPriMapping"] = types.YLeaf{"Cpwvcenetmplsprimapping", cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimapping}
    cpwvcenetmplsprimappingtableentry.EntityData.Leafs["cpwVcEnetMplsPriMappingRowStatus"] = types.YLeaf{"Cpwvcenetmplsprimappingrowstatus", cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimappingrowstatus}
    cpwvcenetmplsprimappingtableentry.EntityData.Leafs["cpwVcEnetMplsPriMappingStorageType"] = types.YLeaf{"Cpwvcenetmplsprimappingstoragetype", cpwvcenetmplsprimappingtableentry.Cpwvcenetmplsprimappingstoragetype}
    return &(cpwvcenetmplsprimappingtableentry.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenetstatstable
// This table contains statistical counters specific for  
// Ethernet PW.
type CISCOIETFPWENETMIB_Cpwvcenetstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents the statistics gathered for the   VC carrying the
    // Ethernet packets since this VC was   first created in the cpwVcEnetTable.
    // The type is slice of
    // CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry.
    Cpwvcenetstatsentry []CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry
}

func (cpwvcenetstatstable *CISCOIETFPWENETMIB_Cpwvcenetstatstable) GetEntityData() *types.CommonEntityData {
    cpwvcenetstatstable.EntityData.YFilter = cpwvcenetstatstable.YFilter
    cpwvcenetstatstable.EntityData.YangName = "cpwVcEnetStatsTable"
    cpwvcenetstatstable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenetstatstable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwvcenetstatstable.EntityData.SegmentPath = "cpwVcEnetStatsTable"
    cpwvcenetstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenetstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenetstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenetstatstable.EntityData.Children = make(map[string]types.YChild)
    cpwvcenetstatstable.EntityData.Children["cpwVcEnetStatsEntry"] = types.YChild{"Cpwvcenetstatsentry", nil}
    for i := range cpwvcenetstatstable.Cpwvcenetstatsentry {
        cpwvcenetstatstable.EntityData.Children[types.GetSegmentPath(&cpwvcenetstatstable.Cpwvcenetstatsentry[i])] = types.YChild{"Cpwvcenetstatsentry", &cpwvcenetstatstable.Cpwvcenetstatsentry[i]}
    }
    cpwvcenetstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcenetstatstable.EntityData)
}

// CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry
// Each entry represents the statistics gathered for the  
// VC carrying the Ethernet packets since this VC was  
// first created in the cpwVcEnetTable.
type CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // The number of packets received (from the PSN) on this VC with   an illegal
    // VLAN field, missing VLAN field that was expected, or   A VLAN field when it
    // was not expected. This counter is not   relevant if the VC type is
    // 'ethernet' (i.e. raw mode), and   should be set to 0 by the agent to
    // indicate this. The type is interface{} with range: 0..18446744073709551615.
    Cpwvcenetstatsillegalvlan interface{}

    // The number of packets that were received with an illegal   Ethernet packet
    // length on this VC. An illegal length is defined  as being greater than the
    // value in the advertised maximum MTU   supported, or shorter than the
    // allowed Ethernet packet size. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwvcenetstatsillegallength interface{}
}

func (cpwvcenetstatsentry *CISCOIETFPWENETMIB_Cpwvcenetstatstable_Cpwvcenetstatsentry) GetEntityData() *types.CommonEntityData {
    cpwvcenetstatsentry.EntityData.YFilter = cpwvcenetstatsentry.YFilter
    cpwvcenetstatsentry.EntityData.YangName = "cpwVcEnetStatsEntry"
    cpwvcenetstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcenetstatsentry.EntityData.ParentYangName = "cpwVcEnetStatsTable"
    cpwvcenetstatsentry.EntityData.SegmentPath = "cpwVcEnetStatsEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcenetstatsentry.Cpwvcindex) + "']"
    cpwvcenetstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcenetstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcenetstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcenetstatsentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcenetstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcenetstatsentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcenetstatsentry.Cpwvcindex}
    cpwvcenetstatsentry.EntityData.Leafs["cpwVcEnetStatsIllegalVlan"] = types.YLeaf{"Cpwvcenetstatsillegalvlan", cpwvcenetstatsentry.Cpwvcenetstatsillegalvlan}
    cpwvcenetstatsentry.EntityData.Leafs["cpwVcEnetStatsIllegalLength"] = types.YLeaf{"Cpwvcenetstatsillegallength", cpwvcenetstatsentry.Cpwvcenetstatsillegallength}
    return &(cpwvcenetstatsentry.EntityData)
}

