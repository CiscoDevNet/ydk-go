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
    CpwVcEnetTable CISCOIETFPWENETMIB_CpwVcEnetTable

    // This table may be used for MPLS PSNs if there is a need   to hold multiple
    // VC, each with different COS, for the same  user service (port + PW VLAN).
    // Such a need may arise if the  MPLS network is capable of L-LSP or E-LSP
    // without multiple  COS capabilities.  Each row is indexed by the cpwVcIndex 
    // and indicate the PRI bits on the packet recieved from the   user port (or
    // VPLS virtual port) that are  classified to this VC. Note that the EXP bit
    // value of the VC  is configured in the CISCO-IETF-PW-MPLS-MIB.
    CpwVcEnetMplsPriMappingTable CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable

    // This table contains statistical counters specific for   Ethernet PW.
    CpwVcEnetStatsTable CISCOIETFPWENETMIB_CpwVcEnetStatsTable
}

func (cISCOIETFPWENETMIB *CISCOIETFPWENETMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWENETMIB.EntityData.YFilter = cISCOIETFPWENETMIB.YFilter
    cISCOIETFPWENETMIB.EntityData.YangName = "CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWENETMIB.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.SegmentPath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB"
    cISCOIETFPWENETMIB.EntityData.AbsolutePath = cISCOIETFPWENETMIB.EntityData.SegmentPath
    cISCOIETFPWENETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWENETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWENETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWENETMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFPWENETMIB.EntityData.Children.Append("cpwVcEnetTable", types.YChild{"CpwVcEnetTable", &cISCOIETFPWENETMIB.CpwVcEnetTable})
    cISCOIETFPWENETMIB.EntityData.Children.Append("cpwVcEnetMplsPriMappingTable", types.YChild{"CpwVcEnetMplsPriMappingTable", &cISCOIETFPWENETMIB.CpwVcEnetMplsPriMappingTable})
    cISCOIETFPWENETMIB.EntityData.Children.Append("cpwVcEnetStatsTable", types.YChild{"CpwVcEnetStatsTable", &cISCOIETFPWENETMIB.CpwVcEnetStatsTable})
    cISCOIETFPWENETMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFPWENETMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFPWENETMIB.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetTable
// This table contains the index to the Ethernet tables  
// associated with this ETH VC, the VLAN configuration and  
// VLAN mode.
type CISCOIETFPWENETMIB_CpwVcEnetTable struct {
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
    // CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry.
    CpwVcEnetEntry []*CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry
}

func (cpwVcEnetTable *CISCOIETFPWENETMIB_CpwVcEnetTable) GetEntityData() *types.CommonEntityData {
    cpwVcEnetTable.EntityData.YFilter = cpwVcEnetTable.YFilter
    cpwVcEnetTable.EntityData.YangName = "cpwVcEnetTable"
    cpwVcEnetTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetTable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwVcEnetTable.EntityData.SegmentPath = "cpwVcEnetTable"
    cpwVcEnetTable.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/" + cpwVcEnetTable.EntityData.SegmentPath
    cpwVcEnetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetTable.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetTable.EntityData.Children.Append("cpwVcEnetEntry", types.YChild{"CpwVcEnetEntry", nil})
    for i := range cpwVcEnetTable.CpwVcEnetEntry {
        cpwVcEnetTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcEnetTable.CpwVcEnetEntry[i]), types.YChild{"CpwVcEnetEntry", cpwVcEnetTable.CpwVcEnetEntry[i]})
    }
    cpwVcEnetTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcEnetTable.EntityData.YListKeys = []string {}

    return &(cpwVcEnetTable.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry
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
type CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. This Object defines the VLAN on the VC. The value
    // of 4097  is used if the object is not applicable, for example when  mapping
    // all packets from an Ethernet port to this VC.  The value of 4096 is used to
    // indicate untagged frames (at   least from the PW point of view), for
    // example if   cpwVcEnetVlanMode is equal 'removeVLAN' or when  
    // cpwVcEnetVlanMode equal 'noChange' and cpwVcEnetPortVlan  is equal 4096.
    // The type is interface{} with range: 0..4097.
    CpwVcEnetPwVlan interface{}

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
    // implementation    is not required. . The type is CpwVcEnetVlanMode.
    CpwVcEnetVlanMode interface{}

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
    CpwVcEnetPortVlan interface{}

    // It is sometimes convenient to model the VC PW as a   virtual interface in
    // the ifTable. In these cases this   object hold the value of the ifIndex in
    // the ifTable   representing this VC PW. A value of zero indicate no such  
    // association or association is not yet known. The type is interface{} with
    // range: 0..2147483647.
    CpwVcEnetVcIfIndex interface{}

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
    CpwVcEnetPortIfIndex interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    CpwVcEnetRowStatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    CpwVcEnetStorageType interface{}
}

func (cpwVcEnetEntry *CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry) GetEntityData() *types.CommonEntityData {
    cpwVcEnetEntry.EntityData.YFilter = cpwVcEnetEntry.YFilter
    cpwVcEnetEntry.EntityData.YangName = "cpwVcEnetEntry"
    cpwVcEnetEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetEntry.EntityData.ParentYangName = "cpwVcEnetTable"
    cpwVcEnetEntry.EntityData.SegmentPath = "cpwVcEnetEntry" + types.AddKeyToken(cpwVcEnetEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwVcEnetEntry.CpwVcEnetPwVlan, "cpwVcEnetPwVlan")
    cpwVcEnetEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/cpwVcEnetTable/" + cpwVcEnetEntry.EntityData.SegmentPath
    cpwVcEnetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcEnetEntry.CpwVcIndex})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetPwVlan", types.YLeaf{"CpwVcEnetPwVlan", cpwVcEnetEntry.CpwVcEnetPwVlan})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetVlanMode", types.YLeaf{"CpwVcEnetVlanMode", cpwVcEnetEntry.CpwVcEnetVlanMode})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetPortVlan", types.YLeaf{"CpwVcEnetPortVlan", cpwVcEnetEntry.CpwVcEnetPortVlan})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetVcIfIndex", types.YLeaf{"CpwVcEnetVcIfIndex", cpwVcEnetEntry.CpwVcEnetVcIfIndex})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetPortIfIndex", types.YLeaf{"CpwVcEnetPortIfIndex", cpwVcEnetEntry.CpwVcEnetPortIfIndex})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetRowStatus", types.YLeaf{"CpwVcEnetRowStatus", cpwVcEnetEntry.CpwVcEnetRowStatus})
    cpwVcEnetEntry.EntityData.Leafs.Append("cpwVcEnetStorageType", types.YLeaf{"CpwVcEnetStorageType", cpwVcEnetEntry.CpwVcEnetStorageType})

    cpwVcEnetEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwVcEnetPwVlan"}

    return &(cpwVcEnetEntry.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode represents   is not required. 
type CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode string

const (
    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_other CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "other"

    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_portBased CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "portBased"

    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_noChange CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "noChange"

    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_changeVlan CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "changeVlan"

    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_addVlan CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "addVlan"

    CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode_removeVlan CISCOIETFPWENETMIB_CpwVcEnetTable_CpwVcEnetEntry_CpwVcEnetVlanMode = "removeVlan"
)

// CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable
// This table may be used for MPLS PSNs if there is a need  
// to hold multiple VC, each with different COS, for the same 
// user service (port + PW VLAN). Such a need may arise if the 
// MPLS network is capable of L-LSP or E-LSP without multiple 
// COS capabilities.  Each row is indexed by the cpwVcIndex  
// and indicate the PRI bits on the packet recieved from the  
// user port (or VPLS virtual port) that are 
// classified to this VC. Note that the EXP bit value of the VC 
// is configured in the CISCO-IETF-PW-MPLS-MIB.
type CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry is created if special classification based on   the PRI bits is
    // required for this VC. The type is slice of
    // CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable_CpwVcEnetMplsPriMappingTableEntry.
    CpwVcEnetMplsPriMappingTableEntry []*CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable_CpwVcEnetMplsPriMappingTableEntry
}

func (cpwVcEnetMplsPriMappingTable *CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable) GetEntityData() *types.CommonEntityData {
    cpwVcEnetMplsPriMappingTable.EntityData.YFilter = cpwVcEnetMplsPriMappingTable.YFilter
    cpwVcEnetMplsPriMappingTable.EntityData.YangName = "cpwVcEnetMplsPriMappingTable"
    cpwVcEnetMplsPriMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetMplsPriMappingTable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwVcEnetMplsPriMappingTable.EntityData.SegmentPath = "cpwVcEnetMplsPriMappingTable"
    cpwVcEnetMplsPriMappingTable.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/" + cpwVcEnetMplsPriMappingTable.EntityData.SegmentPath
    cpwVcEnetMplsPriMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetMplsPriMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetMplsPriMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetMplsPriMappingTable.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetMplsPriMappingTable.EntityData.Children.Append("cpwVcEnetMplsPriMappingTableEntry", types.YChild{"CpwVcEnetMplsPriMappingTableEntry", nil})
    for i := range cpwVcEnetMplsPriMappingTable.CpwVcEnetMplsPriMappingTableEntry {
        cpwVcEnetMplsPriMappingTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcEnetMplsPriMappingTable.CpwVcEnetMplsPriMappingTableEntry[i]), types.YChild{"CpwVcEnetMplsPriMappingTableEntry", cpwVcEnetMplsPriMappingTable.CpwVcEnetMplsPriMappingTableEntry[i]})
    }
    cpwVcEnetMplsPriMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcEnetMplsPriMappingTable.EntityData.YListKeys = []string {}

    return &(cpwVcEnetMplsPriMappingTable.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable_CpwVcEnetMplsPriMappingTableEntry
// Each entry is created if special classification based on  
// the PRI bits is required for this VC.
type CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable_CpwVcEnetMplsPriMappingTableEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

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
    CpwVcEnetMplsPriMapping interface{}

    // Enable creating, deleting and modifying this row. The type is RowStatus.
    CpwVcEnetMplsPriMappingRowStatus interface{}

    // Indicates the storage type of this row. The type is StorageType.
    CpwVcEnetMplsPriMappingStorageType interface{}
}

func (cpwVcEnetMplsPriMappingTableEntry *CISCOIETFPWENETMIB_CpwVcEnetMplsPriMappingTable_CpwVcEnetMplsPriMappingTableEntry) GetEntityData() *types.CommonEntityData {
    cpwVcEnetMplsPriMappingTableEntry.EntityData.YFilter = cpwVcEnetMplsPriMappingTableEntry.YFilter
    cpwVcEnetMplsPriMappingTableEntry.EntityData.YangName = "cpwVcEnetMplsPriMappingTableEntry"
    cpwVcEnetMplsPriMappingTableEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetMplsPriMappingTableEntry.EntityData.ParentYangName = "cpwVcEnetMplsPriMappingTable"
    cpwVcEnetMplsPriMappingTableEntry.EntityData.SegmentPath = "cpwVcEnetMplsPriMappingTableEntry" + types.AddKeyToken(cpwVcEnetMplsPriMappingTableEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcEnetMplsPriMappingTableEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/cpwVcEnetMplsPriMappingTable/" + cpwVcEnetMplsPriMappingTableEntry.EntityData.SegmentPath
    cpwVcEnetMplsPriMappingTableEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetMplsPriMappingTableEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetMplsPriMappingTableEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetMplsPriMappingTableEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetMplsPriMappingTableEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcEnetMplsPriMappingTableEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcEnetMplsPriMappingTableEntry.CpwVcIndex})
    cpwVcEnetMplsPriMappingTableEntry.EntityData.Leafs.Append("cpwVcEnetMplsPriMapping", types.YLeaf{"CpwVcEnetMplsPriMapping", cpwVcEnetMplsPriMappingTableEntry.CpwVcEnetMplsPriMapping})
    cpwVcEnetMplsPriMappingTableEntry.EntityData.Leafs.Append("cpwVcEnetMplsPriMappingRowStatus", types.YLeaf{"CpwVcEnetMplsPriMappingRowStatus", cpwVcEnetMplsPriMappingTableEntry.CpwVcEnetMplsPriMappingRowStatus})
    cpwVcEnetMplsPriMappingTableEntry.EntityData.Leafs.Append("cpwVcEnetMplsPriMappingStorageType", types.YLeaf{"CpwVcEnetMplsPriMappingStorageType", cpwVcEnetMplsPriMappingTableEntry.CpwVcEnetMplsPriMappingStorageType})

    cpwVcEnetMplsPriMappingTableEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcEnetMplsPriMappingTableEntry.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetStatsTable
// This table contains statistical counters specific for  
// Ethernet PW.
type CISCOIETFPWENETMIB_CpwVcEnetStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents the statistics gathered for the   VC carrying the
    // Ethernet packets since this VC was   first created in the cpwVcEnetTable.
    // The type is slice of
    // CISCOIETFPWENETMIB_CpwVcEnetStatsTable_CpwVcEnetStatsEntry.
    CpwVcEnetStatsEntry []*CISCOIETFPWENETMIB_CpwVcEnetStatsTable_CpwVcEnetStatsEntry
}

func (cpwVcEnetStatsTable *CISCOIETFPWENETMIB_CpwVcEnetStatsTable) GetEntityData() *types.CommonEntityData {
    cpwVcEnetStatsTable.EntityData.YFilter = cpwVcEnetStatsTable.YFilter
    cpwVcEnetStatsTable.EntityData.YangName = "cpwVcEnetStatsTable"
    cpwVcEnetStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetStatsTable.EntityData.ParentYangName = "CISCO-IETF-PW-ENET-MIB"
    cpwVcEnetStatsTable.EntityData.SegmentPath = "cpwVcEnetStatsTable"
    cpwVcEnetStatsTable.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/" + cpwVcEnetStatsTable.EntityData.SegmentPath
    cpwVcEnetStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetStatsTable.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetStatsTable.EntityData.Children.Append("cpwVcEnetStatsEntry", types.YChild{"CpwVcEnetStatsEntry", nil})
    for i := range cpwVcEnetStatsTable.CpwVcEnetStatsEntry {
        cpwVcEnetStatsTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcEnetStatsTable.CpwVcEnetStatsEntry[i]), types.YChild{"CpwVcEnetStatsEntry", cpwVcEnetStatsTable.CpwVcEnetStatsEntry[i]})
    }
    cpwVcEnetStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcEnetStatsTable.EntityData.YListKeys = []string {}

    return &(cpwVcEnetStatsTable.EntityData)
}

// CISCOIETFPWENETMIB_CpwVcEnetStatsTable_CpwVcEnetStatsEntry
// Each entry represents the statistics gathered for the  
// VC carrying the Ethernet packets since this VC was  
// first created in the cpwVcEnetTable.
type CISCOIETFPWENETMIB_CpwVcEnetStatsTable_CpwVcEnetStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // The number of packets received (from the PSN) on this VC with   an illegal
    // VLAN field, missing VLAN field that was expected, or   A VLAN field when it
    // was not expected. This counter is not   relevant if the VC type is
    // 'ethernet' (i.e. raw mode), and   should be set to 0 by the agent to
    // indicate this. The type is interface{} with range: 0..18446744073709551615.
    CpwVcEnetStatsIllegalVlan interface{}

    // The number of packets that were received with an illegal   Ethernet packet
    // length on this VC. An illegal length is defined  as being greater than the
    // value in the advertised maximum MTU   supported, or shorter than the
    // allowed Ethernet packet size. The type is interface{} with range:
    // 0..18446744073709551615.
    CpwVcEnetStatsIllegalLength interface{}
}

func (cpwVcEnetStatsEntry *CISCOIETFPWENETMIB_CpwVcEnetStatsTable_CpwVcEnetStatsEntry) GetEntityData() *types.CommonEntityData {
    cpwVcEnetStatsEntry.EntityData.YFilter = cpwVcEnetStatsEntry.YFilter
    cpwVcEnetStatsEntry.EntityData.YangName = "cpwVcEnetStatsEntry"
    cpwVcEnetStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEnetStatsEntry.EntityData.ParentYangName = "cpwVcEnetStatsTable"
    cpwVcEnetStatsEntry.EntityData.SegmentPath = "cpwVcEnetStatsEntry" + types.AddKeyToken(cpwVcEnetStatsEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcEnetStatsEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-ENET-MIB:CISCO-IETF-PW-ENET-MIB/cpwVcEnetStatsTable/" + cpwVcEnetStatsEntry.EntityData.SegmentPath
    cpwVcEnetStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEnetStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEnetStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEnetStatsEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcEnetStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcEnetStatsEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcEnetStatsEntry.CpwVcIndex})
    cpwVcEnetStatsEntry.EntityData.Leafs.Append("cpwVcEnetStatsIllegalVlan", types.YLeaf{"CpwVcEnetStatsIllegalVlan", cpwVcEnetStatsEntry.CpwVcEnetStatsIllegalVlan})
    cpwVcEnetStatsEntry.EntityData.Leafs.Append("cpwVcEnetStatsIllegalLength", types.YLeaf{"CpwVcEnetStatsIllegalLength", cpwVcEnetStatsEntry.CpwVcEnetStatsIllegalLength})

    cpwVcEnetStatsEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcEnetStatsEntry.EntityData)
}

