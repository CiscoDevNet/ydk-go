// This is the MIB Module for ATM and AAL5-related
// objects for managing ATM interfaces, ATM virtual
// links, ATM cross-connects, AAL5 entities, and
// and AAL5 connections.
package atm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ATM-MIB ATM-MIB}", reflect.TypeOf(ATMMIB{}))
    ydk.RegisterEntity("ATM-MIB:ATM-MIB", reflect.TypeOf(ATMMIB{}))
}

// ATMMIB
type ATMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    AtmMIBObjects ATMMIB_AtmMIBObjects

    // This table contains ATM local interface configuration parameters, one entry
    // per ATM interface port.
    AtmInterfaceConfTable ATMMIB_AtmInterfaceConfTable

    // This table contains ATM interface DS3 PLCP parameters and state variables,
    // one entry per ATM interface port.
    AtmInterfaceDs3PlcpTable ATMMIB_AtmInterfaceDs3PlcpTable

    // This table contains ATM interface TC Sublayer parameters and state
    // variables, one entry per ATM interface port.
    AtmInterfaceTCTable ATMMIB_AtmInterfaceTCTable

    // This table contains information on ATM traffic descriptor type and the
    // associated parameters.
    AtmTrafficDescrParamTable ATMMIB_AtmTrafficDescrParamTable

    // The Virtual Path Link (VPL) table.  A bi-directional VPL is modeled as one
    // entry in this table. This table can be used for PVCs, SVCs and Soft PVCs.
    // Entries are not present in this table for the VPIs used by entries in the
    // atmVclTable.
    AtmVplTable ATMMIB_AtmVplTable

    // The Virtual Channel Link (VCL) table.  A bi-directional VCL is modeled as
    // one entry in this table. This table can be used for PVCs, SVCs and Soft
    // PVCs.
    AtmVclTable ATMMIB_AtmVclTable

    // The ATM VP Cross Connect table for PVCs. An entry in this table models two
    // cross-connected VPLs. Each VPL must have its atmConnKind set to pvc(1).
    AtmVpCrossConnectTable ATMMIB_AtmVpCrossConnectTable

    // The ATM VC Cross Connect table for PVCs. An entry in this table models two
    // cross-connected VCLs. Each VCL must have its atmConnKind set to pvc(1).
    AtmVcCrossConnectTable ATMMIB_AtmVcCrossConnectTable

    // This table contains AAL5 VCC performance parameters.
    Aal5VccTable ATMMIB_Aal5VccTable
}

func (aTMMIB *ATMMIB) GetEntityData() *types.CommonEntityData {
    aTMMIB.EntityData.YFilter = aTMMIB.YFilter
    aTMMIB.EntityData.YangName = "ATM-MIB"
    aTMMIB.EntityData.BundleName = "cisco_ios_xe"
    aTMMIB.EntityData.ParentYangName = "ATM-MIB"
    aTMMIB.EntityData.SegmentPath = "ATM-MIB:ATM-MIB"
    aTMMIB.EntityData.AbsolutePath = aTMMIB.EntityData.SegmentPath
    aTMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aTMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aTMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aTMMIB.EntityData.Children = types.NewOrderedMap()
    aTMMIB.EntityData.Children.Append("atmMIBObjects", types.YChild{"AtmMIBObjects", &aTMMIB.AtmMIBObjects})
    aTMMIB.EntityData.Children.Append("atmInterfaceConfTable", types.YChild{"AtmInterfaceConfTable", &aTMMIB.AtmInterfaceConfTable})
    aTMMIB.EntityData.Children.Append("atmInterfaceDs3PlcpTable", types.YChild{"AtmInterfaceDs3PlcpTable", &aTMMIB.AtmInterfaceDs3PlcpTable})
    aTMMIB.EntityData.Children.Append("atmInterfaceTCTable", types.YChild{"AtmInterfaceTCTable", &aTMMIB.AtmInterfaceTCTable})
    aTMMIB.EntityData.Children.Append("atmTrafficDescrParamTable", types.YChild{"AtmTrafficDescrParamTable", &aTMMIB.AtmTrafficDescrParamTable})
    aTMMIB.EntityData.Children.Append("atmVplTable", types.YChild{"AtmVplTable", &aTMMIB.AtmVplTable})
    aTMMIB.EntityData.Children.Append("atmVclTable", types.YChild{"AtmVclTable", &aTMMIB.AtmVclTable})
    aTMMIB.EntityData.Children.Append("atmVpCrossConnectTable", types.YChild{"AtmVpCrossConnectTable", &aTMMIB.AtmVpCrossConnectTable})
    aTMMIB.EntityData.Children.Append("atmVcCrossConnectTable", types.YChild{"AtmVcCrossConnectTable", &aTMMIB.AtmVcCrossConnectTable})
    aTMMIB.EntityData.Children.Append("aal5VccTable", types.YChild{"Aal5VccTable", &aTMMIB.Aal5VccTable})
    aTMMIB.EntityData.Leafs = types.NewOrderedMap()

    aTMMIB.EntityData.YListKeys = []string {}

    return &(aTMMIB.EntityData)
}

// ATMMIB_AtmMIBObjects
type ATMMIB_AtmMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an appropriate value to be used for
    // atmVpCrossConnectIndex when creating entries in the atmVpCrossConnectTable.
    // The value 0 indicates that no unassigned entries are available. To obtain
    // the atmVpCrossConnectIndex value for a new entry, the manager issues a
    // management protocol retrieval operation to obtain the current value of this
    // object.  After each retrieval, the agent should modify the value to the
    // next unassigned index. After a manager retrieves a value the agent will
    // determine through its local policy when this index value will be made
    // available for reuse. The type is interface{} with range: 0..2147483647.
    AtmVpCrossConnectIndexNext interface{}

    // This object contains an appropriate value to be used for
    // atmVcCrossConnectIndex when creating entries in the atmVcCrossConnectTable.
    // The value 0 indicates that no unassigned entries are available. To obtain
    // the atmVcCrossConnectIndex value for a new entry, the manager issues a
    // management protocol retrieval operation to obtain the current value of this
    // object.  After each retrieval, the agent should modify the value to the
    // next unassigned index. After a manager retrieves a value the agent will
    // determine through its local policy when this index value will be made
    // available for reuse. The type is interface{} with range: 0..2147483647.
    AtmVcCrossConnectIndexNext interface{}

    // This object contains an appropriate value to be used for
    // atmTrafficDescrParamIndex when creating entries in the
    // atmTrafficDescrParamTable. The value 0 indicates that no unassigned entries
    // are available. To obtain the atmTrafficDescrParamIndex value for a new
    // entry, the manager issues a management protocol retrieval operation to
    // obtain the current value of this object.  After each retrieval, the agent
    // should modify the value to the next unassigned index. After a manager
    // retrieves a value the agent will determine through its local policy when
    // this index value will be made available for reuse. The type is interface{}
    // with range: 0..2147483647.
    AtmTrafficDescrParamIndexNext interface{}
}

func (atmMIBObjects *ATMMIB_AtmMIBObjects) GetEntityData() *types.CommonEntityData {
    atmMIBObjects.EntityData.YFilter = atmMIBObjects.YFilter
    atmMIBObjects.EntityData.YangName = "atmMIBObjects"
    atmMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    atmMIBObjects.EntityData.ParentYangName = "ATM-MIB"
    atmMIBObjects.EntityData.SegmentPath = "atmMIBObjects"
    atmMIBObjects.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmMIBObjects.EntityData.SegmentPath
    atmMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmMIBObjects.EntityData.Children = types.NewOrderedMap()
    atmMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    atmMIBObjects.EntityData.Leafs.Append("atmVpCrossConnectIndexNext", types.YLeaf{"AtmVpCrossConnectIndexNext", atmMIBObjects.AtmVpCrossConnectIndexNext})
    atmMIBObjects.EntityData.Leafs.Append("atmVcCrossConnectIndexNext", types.YLeaf{"AtmVcCrossConnectIndexNext", atmMIBObjects.AtmVcCrossConnectIndexNext})
    atmMIBObjects.EntityData.Leafs.Append("atmTrafficDescrParamIndexNext", types.YLeaf{"AtmTrafficDescrParamIndexNext", atmMIBObjects.AtmTrafficDescrParamIndexNext})

    atmMIBObjects.EntityData.YListKeys = []string {}

    return &(atmMIBObjects.EntityData)
}

// ATMMIB_AtmInterfaceConfTable
// This table contains ATM local interface
// configuration parameters, one entry per ATM
// interface port.
type ATMMIB_AtmInterfaceConfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains ATM interface configuration parameters and state
    // variables and is indexed by ifIndex values of ATM interfaces. The type is
    // slice of ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry.
    AtmInterfaceConfEntry []*ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry
}

func (atmInterfaceConfTable *ATMMIB_AtmInterfaceConfTable) GetEntityData() *types.CommonEntityData {
    atmInterfaceConfTable.EntityData.YFilter = atmInterfaceConfTable.YFilter
    atmInterfaceConfTable.EntityData.YangName = "atmInterfaceConfTable"
    atmInterfaceConfTable.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceConfTable.EntityData.ParentYangName = "ATM-MIB"
    atmInterfaceConfTable.EntityData.SegmentPath = "atmInterfaceConfTable"
    atmInterfaceConfTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmInterfaceConfTable.EntityData.SegmentPath
    atmInterfaceConfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceConfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceConfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceConfTable.EntityData.Children = types.NewOrderedMap()
    atmInterfaceConfTable.EntityData.Children.Append("atmInterfaceConfEntry", types.YChild{"AtmInterfaceConfEntry", nil})
    for i := range atmInterfaceConfTable.AtmInterfaceConfEntry {
        atmInterfaceConfTable.EntityData.Children.Append(types.GetSegmentPath(atmInterfaceConfTable.AtmInterfaceConfEntry[i]), types.YChild{"AtmInterfaceConfEntry", atmInterfaceConfTable.AtmInterfaceConfEntry[i]})
    }
    atmInterfaceConfTable.EntityData.Leafs = types.NewOrderedMap()

    atmInterfaceConfTable.EntityData.YListKeys = []string {}

    return &(atmInterfaceConfTable.EntityData)
}

// ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry
// This list contains ATM interface configuration
// parameters and state variables and is indexed
// by ifIndex values of ATM interfaces.
type ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The maximum number of VPCs (PVPCs and SVPCs) supported at this ATM
    // interface. At the ATM UNI, the maximum number of VPCs (PVPCs and SVPCs)
    // ranges from 0 to 256 only. The type is interface{} with range: 0..4096.
    AtmInterfaceMaxVpcs interface{}

    // The maximum number of VCCs (PVCCs and SVCCs) supported at this ATM
    // interface. The type is interface{} with range: 0..65536.
    AtmInterfaceMaxVccs interface{}

    // The number of VPCs (PVPC, Soft PVPC and SVPC) currently in use at this ATM
    // interface.  It includes the number of PVPCs and Soft PVPCs that are
    // configured at the interface, plus the number of SVPCs that are currently 
    // established at the interface.  At the ATM UNI, the configured number of
    // VPCs (PVPCs and SVPCs) can range from 0 to 256 only. The type is
    // interface{} with range: 0..4096.
    AtmInterfaceConfVpcs interface{}

    // The number of VCCs (PVCC, Soft PVCC and SVCC) currently in use at this ATM
    // interface.  It includes the number of PVCCs and Soft PVCCs that are
    // configured at the interface, plus the number of SVCCs that are currently 
    // established at the interface. The type is interface{} with range: 0..65536.
    AtmInterfaceConfVccs interface{}

    // The  maximum number of active VPI bits configured for use at the ATM
    // interface. At the ATM UNI, the maximum number of active VPI bits configured
    // for use ranges from 0 to 8 only. The type is interface{} with range: 0..12.
    AtmInterfaceMaxActiveVpiBits interface{}

    // The maximum number of active VCI bits configured for use at this ATM
    // interface. The type is interface{} with range: 0..16.
    AtmInterfaceMaxActiveVciBits interface{}

    // The VPI value of the VCC supporting the ILMI at this ATM interface.  If the
    // values of atmInterfaceIlmiVpi and atmInterfaceIlmiVci are both equal to
    // zero then the ILMI is not supported at this ATM interface. The type is
    // interface{} with range: 0..4095.
    AtmInterfaceIlmiVpi interface{}

    // The VCI value of the VCC supporting the ILMI at this ATM interface.  If the
    // values of atmInterfaceIlmiVpi and atmInterfaceIlmiVci are both equal to
    // zero then the ILMI is not supported at this ATM interface. The type is
    // interface{} with range: 0..65535.
    AtmInterfaceIlmiVci interface{}

    // The type of primary ATM address configured for use at this ATM interface.
    // The type is AtmInterfaceAddressType.
    AtmInterfaceAddressType interface{}

    // The primary address assigned for administrative purposes, for example, an
    // address associated with the service provider side of a public network UNI
    // (thus, the value of this address corresponds with the value of
    // ifPhysAddress at the host side). If this interface has no assigned
    // administrative address, or when the address used for administrative
    // purposes is the same as that used for ifPhysAddress, then this is an octet
    // string of zero length. The type is string.
    AtmInterfaceAdminAddress interface{}

    // The IP address of the neighbor system connected to the  far end of this
    // interface, to which a Network Management Station can send SNMP messages, as
    // IP datagrams sent to UDP port 161, in order to access network management
    // information concerning the operation of that system.  Note that the value
    // of this object may be obtained in different ways, e.g., by manual
    // configuration, or through ILMI interaction with the neighbor system. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AtmInterfaceMyNeighborIpAddress interface{}

    // The textual name of the interface on the neighbor system on the far end of
    // this interface, and to which this interface connects.  If the neighbor
    // system is manageable through SNMP and supports the object ifName, the value
    // of this object must be identical with that of ifName for the ifEntry of the
    // lowest level physical interface for this port.  If this interface does not
    // have a textual name, the value of this object is a zero length string. 
    // Note that the value of this object may be obtained in different ways, e.g.,
    // by manual configuration, or through ILMI interaction with the neighbor
    // system. The type is string.
    AtmInterfaceMyNeighborIfName interface{}

    // The maximum number of VPI Bits that may currently be used at this ATM
    // interface. The value is the minimum of atmInterfaceMaxActiveVpiBits, and
    // the atmInterfaceMaxActiveVpiBits of the interface's UNI/NNI peer.  If the
    // interface does not negotiate with its peer to determine the number of VPI
    // Bits that can be used on the interface, then the value of this object must
    // equal atmInterfaceMaxActiveVpiBits. The type is interface{} with range:
    // 0..12.
    AtmInterfaceCurrentMaxVpiBits interface{}

    // The maximum number of VCI Bits that may currently be used at this ATM
    // interface. The value is the minimum of atmInterfaceMaxActiveVciBits, and
    // the atmInterfaceMaxActiveVciBits of the interface's UNI/NNI peer.  If the
    // interface does not negotiate with its peer to determine the number of VCI
    // Bits that can be used on the interface, then the value of this object must
    // equal atmInterfaceMaxActiveVciBits. The type is interface{} with range:
    // 0..16.
    AtmInterfaceCurrentMaxVciBits interface{}

    // The identifier assigned by a service provider to the network side of a
    // public network UNI. If this interface has no assigned service provider
    // address, or for other interfaces this is an octet string of zero length.
    // The type is string.
    AtmInterfaceSubscrAddress interface{}

    // The number of times the operational status of a PVCL on this interface has
    // gone down. The type is interface{} with range: 0..4294967295.
    AtmIntfPvcFailures interface{}

    // The current number of VCLs on this interface for which there is an active
    // row in the atmVclTable having an atmVclConnKind value of `pvc' and an
    // atmVclOperStatus with a value other than `up'. The type is interface{} with
    // range: 0..4294967295.
    AtmIntfCurrentlyFailingPVcls interface{}

    // Allows the generation of traps in response to PVCL failures on this
    // interface. The type is bool.
    AtmIntfPvcFailuresTrapEnable interface{}

    // The minimum interval between the sending of cIntfPvcFailuresTrap
    // notifications for this interface. The type is interface{} with range:
    // 1..3600. Units are seconds.
    AtmIntfPvcNotificationInterval interface{}

    // The interval for storing the failed time in
    // atmPreviouslyFailedPVclTimeStamp. The type is interface{} with range:
    // 0..3600. Units are seconds.
    AtmPreviouslyFailedPVclInterval interface{}

    // The current number PVCLs on this interface which  changed state to 'up'
    // since the last  atmIntPvcUp2Trap was sent. The type is interface{} with
    // range: 0..4294967295.
    CatmIntfCurrentlyDownToUpPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the OAM
    // loopback failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    CatmIntfOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which the OAM loop back
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    CatmIntfCurrentOAMFailingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // Segment CC OAM failed condition  but the status of each PVCL remain in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    CatmIntfSegCCOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which the Segment CC OAM
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    CatmIntfCurSegCCOAMFailingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // End-to-End CC OAM failed condition  but the status of each PVCL remain in
    // the 'up' state. The type is interface{} with range: 0..4294967295.
    CatmIntfEndCCOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which the End-to-End CC
    // OAM has failed but the status of each PVCL  remain in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    CatmIntfCurEndCCOAMFailingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the AIS
    // RDI OAM failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    CatmIntfAISRDIOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which the AIS RDI OAM has
    // failed but the status of each PVCL remain  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurAISRDIOAMFailingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in any
    // type of OAM failed condition but  the status of each PVCL remain in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    CatmIntfAnyOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which  any of OAM has
    // failed but the status of each PVCL remain  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurAnyOAMFailingPVcls interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    CatmIntfTypeOfOAMFailure interface{}

    // The total number of PVCLs in this interface which  are currently in the OAM
    // loopback recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    CatmIntfOAMRcovedPVcls interface{}

    // The current number of PVCLs on this interface for which the OAM loop back
    // has recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurrentOAMRcovingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // Segment CC OAM recovered condition  and the status of each PVCL is in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    CatmIntfSegCCOAMRcovedPVcls interface{}

    // The current number of PVCLs on this interface for which the Segment CC OAM
    // has recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurSegCCOAMRcovingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // End-to-End CC OAM recovered condition  and the status of each PVCL is in
    // the 'up' state. The type is interface{} with range: 0..4294967295.
    CatmIntfEndCCOAMRcovedPVcls interface{}

    // The current number of PVCLs on this interface for which the End-to-End CC
    // OAM has recovered and the status of each PVCL  is in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    CatmIntfCurEndCCOAMRcovingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the AIS
    // RDI OAM recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    CatmIntfAISRDIOAMRcovedPVcls interface{}

    // The current number of PVCLs on this interface for which the AIS RDI OAM has
    // recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurAISRDIOAMRcovingPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in any
    // type of OAM recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    CatmIntfAnyOAMRcovedPVcls interface{}

    // The current number of PVCLs on this interface for which  any of OAM has
    // recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    CatmIntfCurAnyOAMRcovingPVcls interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    CatmIntfTypeOfOAMRecover interface{}

    // The current number PVCLs on this interface which  changed state to 'up'
    // since the last  atmIntPvcUpTrap was sent. The type is interface{} with
    // range: 0..4294967295.
    AtmIntfCurrentlyDownToUpPVcls interface{}

    // The total number of PVCLs in this interface which  are currently in the oam
    // loopback failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    AtmIntfOAMFailedPVcls interface{}

    // The current number of PVCLs on this interface for which the oam loop back
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    AtmIntfCurrentlyOAMFailingPVcls interface{}
}

func (atmInterfaceConfEntry *ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry) GetEntityData() *types.CommonEntityData {
    atmInterfaceConfEntry.EntityData.YFilter = atmInterfaceConfEntry.YFilter
    atmInterfaceConfEntry.EntityData.YangName = "atmInterfaceConfEntry"
    atmInterfaceConfEntry.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceConfEntry.EntityData.ParentYangName = "atmInterfaceConfTable"
    atmInterfaceConfEntry.EntityData.SegmentPath = "atmInterfaceConfEntry" + types.AddKeyToken(atmInterfaceConfEntry.IfIndex, "ifIndex")
    atmInterfaceConfEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmInterfaceConfTable/" + atmInterfaceConfEntry.EntityData.SegmentPath
    atmInterfaceConfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceConfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceConfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceConfEntry.EntityData.Children = types.NewOrderedMap()
    atmInterfaceConfEntry.EntityData.Leafs = types.NewOrderedMap()
    atmInterfaceConfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmInterfaceConfEntry.IfIndex})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMaxVpcs", types.YLeaf{"AtmInterfaceMaxVpcs", atmInterfaceConfEntry.AtmInterfaceMaxVpcs})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMaxVccs", types.YLeaf{"AtmInterfaceMaxVccs", atmInterfaceConfEntry.AtmInterfaceMaxVccs})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceConfVpcs", types.YLeaf{"AtmInterfaceConfVpcs", atmInterfaceConfEntry.AtmInterfaceConfVpcs})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceConfVccs", types.YLeaf{"AtmInterfaceConfVccs", atmInterfaceConfEntry.AtmInterfaceConfVccs})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMaxActiveVpiBits", types.YLeaf{"AtmInterfaceMaxActiveVpiBits", atmInterfaceConfEntry.AtmInterfaceMaxActiveVpiBits})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMaxActiveVciBits", types.YLeaf{"AtmInterfaceMaxActiveVciBits", atmInterfaceConfEntry.AtmInterfaceMaxActiveVciBits})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceIlmiVpi", types.YLeaf{"AtmInterfaceIlmiVpi", atmInterfaceConfEntry.AtmInterfaceIlmiVpi})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceIlmiVci", types.YLeaf{"AtmInterfaceIlmiVci", atmInterfaceConfEntry.AtmInterfaceIlmiVci})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceAddressType", types.YLeaf{"AtmInterfaceAddressType", atmInterfaceConfEntry.AtmInterfaceAddressType})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceAdminAddress", types.YLeaf{"AtmInterfaceAdminAddress", atmInterfaceConfEntry.AtmInterfaceAdminAddress})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMyNeighborIpAddress", types.YLeaf{"AtmInterfaceMyNeighborIpAddress", atmInterfaceConfEntry.AtmInterfaceMyNeighborIpAddress})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceMyNeighborIfName", types.YLeaf{"AtmInterfaceMyNeighborIfName", atmInterfaceConfEntry.AtmInterfaceMyNeighborIfName})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceCurrentMaxVpiBits", types.YLeaf{"AtmInterfaceCurrentMaxVpiBits", atmInterfaceConfEntry.AtmInterfaceCurrentMaxVpiBits})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceCurrentMaxVciBits", types.YLeaf{"AtmInterfaceCurrentMaxVciBits", atmInterfaceConfEntry.AtmInterfaceCurrentMaxVciBits})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmInterfaceSubscrAddress", types.YLeaf{"AtmInterfaceSubscrAddress", atmInterfaceConfEntry.AtmInterfaceSubscrAddress})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfPvcFailures", types.YLeaf{"AtmIntfPvcFailures", atmInterfaceConfEntry.AtmIntfPvcFailures})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfCurrentlyFailingPVcls", types.YLeaf{"AtmIntfCurrentlyFailingPVcls", atmInterfaceConfEntry.AtmIntfCurrentlyFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfPvcFailuresTrapEnable", types.YLeaf{"AtmIntfPvcFailuresTrapEnable", atmInterfaceConfEntry.AtmIntfPvcFailuresTrapEnable})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfPvcNotificationInterval", types.YLeaf{"AtmIntfPvcNotificationInterval", atmInterfaceConfEntry.AtmIntfPvcNotificationInterval})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmPreviouslyFailedPVclInterval", types.YLeaf{"AtmPreviouslyFailedPVclInterval", atmInterfaceConfEntry.AtmPreviouslyFailedPVclInterval})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurrentlyDownToUpPVcls", types.YLeaf{"CatmIntfCurrentlyDownToUpPVcls", atmInterfaceConfEntry.CatmIntfCurrentlyDownToUpPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfOAMFailedPVcls", types.YLeaf{"CatmIntfOAMFailedPVcls", atmInterfaceConfEntry.CatmIntfOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurrentOAMFailingPVcls", types.YLeaf{"CatmIntfCurrentOAMFailingPVcls", atmInterfaceConfEntry.CatmIntfCurrentOAMFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfSegCCOAMFailedPVcls", types.YLeaf{"CatmIntfSegCCOAMFailedPVcls", atmInterfaceConfEntry.CatmIntfSegCCOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurSegCCOAMFailingPVcls", types.YLeaf{"CatmIntfCurSegCCOAMFailingPVcls", atmInterfaceConfEntry.CatmIntfCurSegCCOAMFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfEndCCOAMFailedPVcls", types.YLeaf{"CatmIntfEndCCOAMFailedPVcls", atmInterfaceConfEntry.CatmIntfEndCCOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurEndCCOAMFailingPVcls", types.YLeaf{"CatmIntfCurEndCCOAMFailingPVcls", atmInterfaceConfEntry.CatmIntfCurEndCCOAMFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfAISRDIOAMFailedPVcls", types.YLeaf{"CatmIntfAISRDIOAMFailedPVcls", atmInterfaceConfEntry.CatmIntfAISRDIOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurAISRDIOAMFailingPVcls", types.YLeaf{"CatmIntfCurAISRDIOAMFailingPVcls", atmInterfaceConfEntry.CatmIntfCurAISRDIOAMFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfAnyOAMFailedPVcls", types.YLeaf{"CatmIntfAnyOAMFailedPVcls", atmInterfaceConfEntry.CatmIntfAnyOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurAnyOAMFailingPVcls", types.YLeaf{"CatmIntfCurAnyOAMFailingPVcls", atmInterfaceConfEntry.CatmIntfCurAnyOAMFailingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfTypeOfOAMFailure", types.YLeaf{"CatmIntfTypeOfOAMFailure", atmInterfaceConfEntry.CatmIntfTypeOfOAMFailure})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfOAMRcovedPVcls", types.YLeaf{"CatmIntfOAMRcovedPVcls", atmInterfaceConfEntry.CatmIntfOAMRcovedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurrentOAMRcovingPVcls", types.YLeaf{"CatmIntfCurrentOAMRcovingPVcls", atmInterfaceConfEntry.CatmIntfCurrentOAMRcovingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfSegCCOAMRcovedPVcls", types.YLeaf{"CatmIntfSegCCOAMRcovedPVcls", atmInterfaceConfEntry.CatmIntfSegCCOAMRcovedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurSegCCOAMRcovingPVcls", types.YLeaf{"CatmIntfCurSegCCOAMRcovingPVcls", atmInterfaceConfEntry.CatmIntfCurSegCCOAMRcovingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfEndCCOAMRcovedPVcls", types.YLeaf{"CatmIntfEndCCOAMRcovedPVcls", atmInterfaceConfEntry.CatmIntfEndCCOAMRcovedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurEndCCOAMRcovingPVcls", types.YLeaf{"CatmIntfCurEndCCOAMRcovingPVcls", atmInterfaceConfEntry.CatmIntfCurEndCCOAMRcovingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfAISRDIOAMRcovedPVcls", types.YLeaf{"CatmIntfAISRDIOAMRcovedPVcls", atmInterfaceConfEntry.CatmIntfAISRDIOAMRcovedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurAISRDIOAMRcovingPVcls", types.YLeaf{"CatmIntfCurAISRDIOAMRcovingPVcls", atmInterfaceConfEntry.CatmIntfCurAISRDIOAMRcovingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfAnyOAMRcovedPVcls", types.YLeaf{"CatmIntfAnyOAMRcovedPVcls", atmInterfaceConfEntry.CatmIntfAnyOAMRcovedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfCurAnyOAMRcovingPVcls", types.YLeaf{"CatmIntfCurAnyOAMRcovingPVcls", atmInterfaceConfEntry.CatmIntfCurAnyOAMRcovingPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("catmIntfTypeOfOAMRecover", types.YLeaf{"CatmIntfTypeOfOAMRecover", atmInterfaceConfEntry.CatmIntfTypeOfOAMRecover})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfCurrentlyDownToUpPVcls", types.YLeaf{"AtmIntfCurrentlyDownToUpPVcls", atmInterfaceConfEntry.AtmIntfCurrentlyDownToUpPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfOAMFailedPVcls", types.YLeaf{"AtmIntfOAMFailedPVcls", atmInterfaceConfEntry.AtmIntfOAMFailedPVcls})
    atmInterfaceConfEntry.EntityData.Leafs.Append("atmIntfCurrentlyOAMFailingPVcls", types.YLeaf{"AtmIntfCurrentlyOAMFailingPVcls", atmInterfaceConfEntry.AtmIntfCurrentlyOAMFailingPVcls})

    atmInterfaceConfEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(atmInterfaceConfEntry.EntityData)
}

// ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType represents for use at this ATM interface.
type ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType string

const (
    ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType_private ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType = "private"

    ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType_nsapE164 ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType = "nsapE164"

    ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType_nativeE164 ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType = "nativeE164"

    ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType_other ATMMIB_AtmInterfaceConfTable_AtmInterfaceConfEntry_AtmInterfaceAddressType = "other"
)

// ATMMIB_AtmInterfaceDs3PlcpTable
// This table contains ATM interface DS3 PLCP
// parameters and state variables, one entry per
// ATM interface port.
type ATMMIB_AtmInterfaceDs3PlcpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains DS3 PLCP parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry.
    AtmInterfaceDs3PlcpEntry []*ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry
}

func (atmInterfaceDs3PlcpTable *ATMMIB_AtmInterfaceDs3PlcpTable) GetEntityData() *types.CommonEntityData {
    atmInterfaceDs3PlcpTable.EntityData.YFilter = atmInterfaceDs3PlcpTable.YFilter
    atmInterfaceDs3PlcpTable.EntityData.YangName = "atmInterfaceDs3PlcpTable"
    atmInterfaceDs3PlcpTable.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceDs3PlcpTable.EntityData.ParentYangName = "ATM-MIB"
    atmInterfaceDs3PlcpTable.EntityData.SegmentPath = "atmInterfaceDs3PlcpTable"
    atmInterfaceDs3PlcpTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmInterfaceDs3PlcpTable.EntityData.SegmentPath
    atmInterfaceDs3PlcpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceDs3PlcpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceDs3PlcpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceDs3PlcpTable.EntityData.Children = types.NewOrderedMap()
    atmInterfaceDs3PlcpTable.EntityData.Children.Append("atmInterfaceDs3PlcpEntry", types.YChild{"AtmInterfaceDs3PlcpEntry", nil})
    for i := range atmInterfaceDs3PlcpTable.AtmInterfaceDs3PlcpEntry {
        atmInterfaceDs3PlcpTable.EntityData.Children.Append(types.GetSegmentPath(atmInterfaceDs3PlcpTable.AtmInterfaceDs3PlcpEntry[i]), types.YChild{"AtmInterfaceDs3PlcpEntry", atmInterfaceDs3PlcpTable.AtmInterfaceDs3PlcpEntry[i]})
    }
    atmInterfaceDs3PlcpTable.EntityData.Leafs = types.NewOrderedMap()

    atmInterfaceDs3PlcpTable.EntityData.YListKeys = []string {}

    return &(atmInterfaceDs3PlcpTable.EntityData)
}

// ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry
// This list contains DS3 PLCP parameters and
// state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of DS3 PLCP Severely Errored Framing Seconds (SEFS). Each SEFS
    // represents a one-second interval which contains one or more SEF events. The
    // type is interface{} with range: 0..4294967295.
    AtmInterfaceDs3PlcpSEFSs interface{}

    // This variable indicates if there is an alarm present for the DS3 PLCP.  The
    // value receivedFarEndAlarm means that the DS3 PLCP has received an incoming
    // Yellow Signal, the value incomingLOF means that the DS3 PLCP has declared a
    // loss of frame (LOF) failure condition, and the value noAlarm means that
    // there are no alarms present. Transition from the failure to the no alarm
    // state occurs when no defects (e.g., LOF) are received for more than 10
    // seconds. The type is AtmInterfaceDs3PlcpAlarmState.
    AtmInterfaceDs3PlcpAlarmState interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by the PLCP. The type is interface{} with range: 0..4294967295.
    AtmInterfaceDs3PlcpUASs interface{}
}

func (atmInterfaceDs3PlcpEntry *ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry) GetEntityData() *types.CommonEntityData {
    atmInterfaceDs3PlcpEntry.EntityData.YFilter = atmInterfaceDs3PlcpEntry.YFilter
    atmInterfaceDs3PlcpEntry.EntityData.YangName = "atmInterfaceDs3PlcpEntry"
    atmInterfaceDs3PlcpEntry.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceDs3PlcpEntry.EntityData.ParentYangName = "atmInterfaceDs3PlcpTable"
    atmInterfaceDs3PlcpEntry.EntityData.SegmentPath = "atmInterfaceDs3PlcpEntry" + types.AddKeyToken(atmInterfaceDs3PlcpEntry.IfIndex, "ifIndex")
    atmInterfaceDs3PlcpEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmInterfaceDs3PlcpTable/" + atmInterfaceDs3PlcpEntry.EntityData.SegmentPath
    atmInterfaceDs3PlcpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceDs3PlcpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceDs3PlcpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceDs3PlcpEntry.EntityData.Children = types.NewOrderedMap()
    atmInterfaceDs3PlcpEntry.EntityData.Leafs = types.NewOrderedMap()
    atmInterfaceDs3PlcpEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmInterfaceDs3PlcpEntry.IfIndex})
    atmInterfaceDs3PlcpEntry.EntityData.Leafs.Append("atmInterfaceDs3PlcpSEFSs", types.YLeaf{"AtmInterfaceDs3PlcpSEFSs", atmInterfaceDs3PlcpEntry.AtmInterfaceDs3PlcpSEFSs})
    atmInterfaceDs3PlcpEntry.EntityData.Leafs.Append("atmInterfaceDs3PlcpAlarmState", types.YLeaf{"AtmInterfaceDs3PlcpAlarmState", atmInterfaceDs3PlcpEntry.AtmInterfaceDs3PlcpAlarmState})
    atmInterfaceDs3PlcpEntry.EntityData.Leafs.Append("atmInterfaceDs3PlcpUASs", types.YLeaf{"AtmInterfaceDs3PlcpUASs", atmInterfaceDs3PlcpEntry.AtmInterfaceDs3PlcpUASs})

    atmInterfaceDs3PlcpEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(atmInterfaceDs3PlcpEntry.EntityData)
}

// ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState represents for more than 10 seconds.
type ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState string

const (
    ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState_noAlarm ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState = "noAlarm"

    ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState_receivedFarEndAlarm ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState = "receivedFarEndAlarm"

    ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState_incomingLOF ATMMIB_AtmInterfaceDs3PlcpTable_AtmInterfaceDs3PlcpEntry_AtmInterfaceDs3PlcpAlarmState = "incomingLOF"
)

// ATMMIB_AtmInterfaceTCTable
// This table contains ATM interface TC
// Sublayer parameters and state variables,
// one entry per ATM interface port.
type ATMMIB_AtmInterfaceTCTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains TC Sublayer parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry.
    AtmInterfaceTCEntry []*ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry
}

func (atmInterfaceTCTable *ATMMIB_AtmInterfaceTCTable) GetEntityData() *types.CommonEntityData {
    atmInterfaceTCTable.EntityData.YFilter = atmInterfaceTCTable.YFilter
    atmInterfaceTCTable.EntityData.YangName = "atmInterfaceTCTable"
    atmInterfaceTCTable.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceTCTable.EntityData.ParentYangName = "ATM-MIB"
    atmInterfaceTCTable.EntityData.SegmentPath = "atmInterfaceTCTable"
    atmInterfaceTCTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmInterfaceTCTable.EntityData.SegmentPath
    atmInterfaceTCTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceTCTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceTCTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceTCTable.EntityData.Children = types.NewOrderedMap()
    atmInterfaceTCTable.EntityData.Children.Append("atmInterfaceTCEntry", types.YChild{"AtmInterfaceTCEntry", nil})
    for i := range atmInterfaceTCTable.AtmInterfaceTCEntry {
        atmInterfaceTCTable.EntityData.Children.Append(types.GetSegmentPath(atmInterfaceTCTable.AtmInterfaceTCEntry[i]), types.YChild{"AtmInterfaceTCEntry", atmInterfaceTCTable.AtmInterfaceTCEntry[i]})
    }
    atmInterfaceTCTable.EntityData.Leafs = types.NewOrderedMap()

    atmInterfaceTCTable.EntityData.YListKeys = []string {}

    return &(atmInterfaceTCTable.EntityData)
}

// ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry
// This list contains TC Sublayer parameters
// and state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of times the Out of Cell Delineation (OCD) events occur.  If
    // seven consecutive ATM cells have Header Error Control (HEC) violations, an
    // OCD event occurs. A high number of OCD events may indicate a problem with
    // the TC Sublayer. The type is interface{} with range: 0..4294967295.
    AtmInterfaceOCDEvents interface{}

    // This variable indicates if there is an alarm present for the TC Sublayer. 
    // The value lcdFailure(2) indicates that the TC Sublayer is currently in the
    // Loss of Cell Delineation (LCD) defect maintenance state.  The value
    // noAlarm(1) indicates that the TC Sublayer is currently not in the LCD
    // defect maintenance state. The type is AtmInterfaceTCAlarmState.
    AtmInterfaceTCAlarmState interface{}
}

func (atmInterfaceTCEntry *ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry) GetEntityData() *types.CommonEntityData {
    atmInterfaceTCEntry.EntityData.YFilter = atmInterfaceTCEntry.YFilter
    atmInterfaceTCEntry.EntityData.YangName = "atmInterfaceTCEntry"
    atmInterfaceTCEntry.EntityData.BundleName = "cisco_ios_xe"
    atmInterfaceTCEntry.EntityData.ParentYangName = "atmInterfaceTCTable"
    atmInterfaceTCEntry.EntityData.SegmentPath = "atmInterfaceTCEntry" + types.AddKeyToken(atmInterfaceTCEntry.IfIndex, "ifIndex")
    atmInterfaceTCEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmInterfaceTCTable/" + atmInterfaceTCEntry.EntityData.SegmentPath
    atmInterfaceTCEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmInterfaceTCEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmInterfaceTCEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmInterfaceTCEntry.EntityData.Children = types.NewOrderedMap()
    atmInterfaceTCEntry.EntityData.Leafs = types.NewOrderedMap()
    atmInterfaceTCEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmInterfaceTCEntry.IfIndex})
    atmInterfaceTCEntry.EntityData.Leafs.Append("atmInterfaceOCDEvents", types.YLeaf{"AtmInterfaceOCDEvents", atmInterfaceTCEntry.AtmInterfaceOCDEvents})
    atmInterfaceTCEntry.EntityData.Leafs.Append("atmInterfaceTCAlarmState", types.YLeaf{"AtmInterfaceTCAlarmState", atmInterfaceTCEntry.AtmInterfaceTCAlarmState})

    atmInterfaceTCEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(atmInterfaceTCEntry.EntityData)
}

// ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState represents maintenance state.
type ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState string

const (
    ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState_noAlarm ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState = "noAlarm"

    ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState_lcdFailure ATMMIB_AtmInterfaceTCTable_AtmInterfaceTCEntry_AtmInterfaceTCAlarmState = "lcdFailure"
)

// ATMMIB_AtmTrafficDescrParamTable
// This table contains information on ATM traffic
// descriptor type and the associated parameters.
type ATMMIB_AtmTrafficDescrParamTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains ATM traffic descriptor type and the associated
    // parameters. The type is slice of
    // ATMMIB_AtmTrafficDescrParamTable_AtmTrafficDescrParamEntry.
    AtmTrafficDescrParamEntry []*ATMMIB_AtmTrafficDescrParamTable_AtmTrafficDescrParamEntry
}

func (atmTrafficDescrParamTable *ATMMIB_AtmTrafficDescrParamTable) GetEntityData() *types.CommonEntityData {
    atmTrafficDescrParamTable.EntityData.YFilter = atmTrafficDescrParamTable.YFilter
    atmTrafficDescrParamTable.EntityData.YangName = "atmTrafficDescrParamTable"
    atmTrafficDescrParamTable.EntityData.BundleName = "cisco_ios_xe"
    atmTrafficDescrParamTable.EntityData.ParentYangName = "ATM-MIB"
    atmTrafficDescrParamTable.EntityData.SegmentPath = "atmTrafficDescrParamTable"
    atmTrafficDescrParamTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmTrafficDescrParamTable.EntityData.SegmentPath
    atmTrafficDescrParamTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmTrafficDescrParamTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmTrafficDescrParamTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmTrafficDescrParamTable.EntityData.Children = types.NewOrderedMap()
    atmTrafficDescrParamTable.EntityData.Children.Append("atmTrafficDescrParamEntry", types.YChild{"AtmTrafficDescrParamEntry", nil})
    for i := range atmTrafficDescrParamTable.AtmTrafficDescrParamEntry {
        atmTrafficDescrParamTable.EntityData.Children.Append(types.GetSegmentPath(atmTrafficDescrParamTable.AtmTrafficDescrParamEntry[i]), types.YChild{"AtmTrafficDescrParamEntry", atmTrafficDescrParamTable.AtmTrafficDescrParamEntry[i]})
    }
    atmTrafficDescrParamTable.EntityData.Leafs = types.NewOrderedMap()

    atmTrafficDescrParamTable.EntityData.YListKeys = []string {}

    return &(atmTrafficDescrParamTable.EntityData)
}

// ATMMIB_AtmTrafficDescrParamTable_AtmTrafficDescrParamEntry
// This list contains ATM traffic descriptor
// type and the associated parameters.
type ATMMIB_AtmTrafficDescrParamTable_AtmTrafficDescrParamEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object is used by the virtual link table
    // (i.e., VPL or VCL table) to identify the row of this table. When creating a
    // new row in the table the value of this index may be obtained by retrieving
    // the value of atmTrafficDescrParamIndexNext. The type is interface{} with
    // range: 1..2147483647.
    AtmTrafficDescrParamIndex interface{}

    // The value of this object identifies the type of ATM traffic descriptor. The
    // type may indicate no traffic descriptor or traffic descriptor with one or
    // more parameters. These parameters are specified as a parameter vector, in
    // the corresponding instances of the objects:     atmTrafficDescrParam1    
    // atmTrafficDescrParam2     atmTrafficDescrParam3     atmTrafficDescrParam4  
    // atmTrafficDescrParam5. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    AtmTrafficDescrType interface{}

    // The first parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    AtmTrafficDescrParam1 interface{}

    // The second parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    AtmTrafficDescrParam2 interface{}

    // The third parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    AtmTrafficDescrParam3 interface{}

    // The fourth parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    AtmTrafficDescrParam4 interface{}

    // The fifth parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    AtmTrafficDescrParam5 interface{}

    // The value of this object identifies the QoS Class. Four Service classes
    // have been specified in the ATM Forum UNI Specification: Service Class A:
    // Constant bit rate video and                  Circuit emulation Service
    // Class B: Variable bit rate video/audio Service Class C: Connection-oriented
    // data Service Class D: Connectionless data Four QoS classes numbered 1, 2,
    // 3, and 4 have been specified with the aim to support service classes A, B,
    // C, and D respectively. An unspecified QoS Class numbered `0' is used for
    // best effort traffic. The type is interface{} with range: 0..255.
    AtmTrafficQoSClass interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The type is RowStatus.
    AtmTrafficDescrRowStatus interface{}

    // The ATM service category. The type is AtmServiceCategory.
    AtmServiceCategory interface{}

    // If set to 'true', this object indicates that the network is requested to
    // treat data for this connection, in the given direction, as frames (e.g.
    // AAL5 CPCS_PDU's) rather than as individual cells.  While the precise
    // implementation is network-specific, this treatment may for example involve
    // discarding entire frames during congestion, rather than a few cells from
    // many frames. The type is bool.
    AtmTrafficFrameDiscard interface{}
}

func (atmTrafficDescrParamEntry *ATMMIB_AtmTrafficDescrParamTable_AtmTrafficDescrParamEntry) GetEntityData() *types.CommonEntityData {
    atmTrafficDescrParamEntry.EntityData.YFilter = atmTrafficDescrParamEntry.YFilter
    atmTrafficDescrParamEntry.EntityData.YangName = "atmTrafficDescrParamEntry"
    atmTrafficDescrParamEntry.EntityData.BundleName = "cisco_ios_xe"
    atmTrafficDescrParamEntry.EntityData.ParentYangName = "atmTrafficDescrParamTable"
    atmTrafficDescrParamEntry.EntityData.SegmentPath = "atmTrafficDescrParamEntry" + types.AddKeyToken(atmTrafficDescrParamEntry.AtmTrafficDescrParamIndex, "atmTrafficDescrParamIndex")
    atmTrafficDescrParamEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmTrafficDescrParamTable/" + atmTrafficDescrParamEntry.EntityData.SegmentPath
    atmTrafficDescrParamEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmTrafficDescrParamEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmTrafficDescrParamEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmTrafficDescrParamEntry.EntityData.Children = types.NewOrderedMap()
    atmTrafficDescrParamEntry.EntityData.Leafs = types.NewOrderedMap()
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParamIndex", types.YLeaf{"AtmTrafficDescrParamIndex", atmTrafficDescrParamEntry.AtmTrafficDescrParamIndex})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrType", types.YLeaf{"AtmTrafficDescrType", atmTrafficDescrParamEntry.AtmTrafficDescrType})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParam1", types.YLeaf{"AtmTrafficDescrParam1", atmTrafficDescrParamEntry.AtmTrafficDescrParam1})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParam2", types.YLeaf{"AtmTrafficDescrParam2", atmTrafficDescrParamEntry.AtmTrafficDescrParam2})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParam3", types.YLeaf{"AtmTrafficDescrParam3", atmTrafficDescrParamEntry.AtmTrafficDescrParam3})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParam4", types.YLeaf{"AtmTrafficDescrParam4", atmTrafficDescrParamEntry.AtmTrafficDescrParam4})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrParam5", types.YLeaf{"AtmTrafficDescrParam5", atmTrafficDescrParamEntry.AtmTrafficDescrParam5})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficQoSClass", types.YLeaf{"AtmTrafficQoSClass", atmTrafficDescrParamEntry.AtmTrafficQoSClass})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficDescrRowStatus", types.YLeaf{"AtmTrafficDescrRowStatus", atmTrafficDescrParamEntry.AtmTrafficDescrRowStatus})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmServiceCategory", types.YLeaf{"AtmServiceCategory", atmTrafficDescrParamEntry.AtmServiceCategory})
    atmTrafficDescrParamEntry.EntityData.Leafs.Append("atmTrafficFrameDiscard", types.YLeaf{"AtmTrafficFrameDiscard", atmTrafficDescrParamEntry.AtmTrafficFrameDiscard})

    atmTrafficDescrParamEntry.EntityData.YListKeys = []string {"AtmTrafficDescrParamIndex"}

    return &(atmTrafficDescrParamEntry.EntityData)
}

// ATMMIB_AtmVplTable
// The Virtual Path Link (VPL) table.  A
// bi-directional VPL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
// Entries are not present in this table for
// the VPIs used by entries in the atmVclTable.
type ATMMIB_AtmVplTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the VPL table.  This entry is used to model a bi-directional
    // VPL. To create a VPL at an ATM interface, either of the following
    // procedures are used:  Negotiated VPL establishment  (1) The management
    // application creates   a VPL entry in the atmVplTable   by setting
    // atmVplRowStatus to createAndWait(5).   This may fail for the following
    // reasons:   - The selected VPI value is unavailable,   - The selected VPI
    // value is in use.   Otherwise, the agent creates a row and   reserves the
    // VPI value on that port.  (2) The manager selects an existing row(s) in the 
    // atmTrafficDescrParamTable,   thereby, selecting a set of self-consistent  
    // ATM traffic parameters and the service category   for receive and transmit
    // directions of the VPL.  (2a) If no suitable row(s) in the  
    // atmTrafficDescrParamTable exists,   the manager must create a new row(s)  
    // in that table.  (2b) The manager characterizes the VPL's traffic  
    // parameters through setting the   atmVplReceiveTrafficDescrIndex and the  
    // atmVplTransmitTrafficDescrIndex values   in the VPL table, which point to
    // the rows   containing desired ATM traffic parameter values   in the
    // atmTrafficDescrParamTable.  The agent   will check the availability of
    // resources and   may refuse the request.   If the transmit and receive
    // service categories   are inconsistent, the agent should refuse the  
    // request.  (3) The manager activates the VPL by setting the   the
    // atmVplRowStatus to active(1).   If this set is successful, the agent has  
    // reserved the resources to satisfy the requested   traffic parameter values
    // and the service category   for that VPL.  (4) If the VPL terminates a VPC
    // in the ATM host   or switch, the manager turns on the   atmVplAdminStatus
    // to up(1) to turn the VPL   traffic flow on.  Otherwise, the  
    // atmVpCrossConnectTable  must be used   to cross-connect the VPL to another
    // VPL(s)   in an ATM switch or network.  One-Shot VPL Establishment  A VPL
    // may also be established in one step by a set-request with all necessary VPL
    // parameter values and atmVplRowStatus set to createAndGo(4).  In contrast to
    // the negotiated VPL establishment which allows for detailed error checking
    // (i.e., set errors are explicitly linked to particular resource acquisition
    // failures), the one-shot VPL establishment performs the setup on one
    // operation but does not have the advantage of step-wise error checking.  VPL
    // Retirement  A VPL is released by setting atmVplRowStatus to destroy(6), and
    // the agent may release all associated resources. The type is slice of
    // ATMMIB_AtmVplTable_AtmVplEntry.
    AtmVplEntry []*ATMMIB_AtmVplTable_AtmVplEntry
}

func (atmVplTable *ATMMIB_AtmVplTable) GetEntityData() *types.CommonEntityData {
    atmVplTable.EntityData.YFilter = atmVplTable.YFilter
    atmVplTable.EntityData.YangName = "atmVplTable"
    atmVplTable.EntityData.BundleName = "cisco_ios_xe"
    atmVplTable.EntityData.ParentYangName = "ATM-MIB"
    atmVplTable.EntityData.SegmentPath = "atmVplTable"
    atmVplTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmVplTable.EntityData.SegmentPath
    atmVplTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVplTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVplTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVplTable.EntityData.Children = types.NewOrderedMap()
    atmVplTable.EntityData.Children.Append("atmVplEntry", types.YChild{"AtmVplEntry", nil})
    for i := range atmVplTable.AtmVplEntry {
        atmVplTable.EntityData.Children.Append(types.GetSegmentPath(atmVplTable.AtmVplEntry[i]), types.YChild{"AtmVplEntry", atmVplTable.AtmVplEntry[i]})
    }
    atmVplTable.EntityData.Leafs = types.NewOrderedMap()

    atmVplTable.EntityData.YListKeys = []string {}

    return &(atmVplTable.EntityData)
}

// ATMMIB_AtmVplTable_AtmVplEntry
// An entry in the VPL table.  This entry is
// used to model a bi-directional VPL.
// To create a VPL at an ATM interface,
// either of the following procedures are used:
// 
// Negotiated VPL establishment
// 
// (1) The management application creates
//   a VPL entry in the atmVplTable
//   by setting atmVplRowStatus to createAndWait(5).
//   This may fail for the following reasons:
//   - The selected VPI value is unavailable,
//   - The selected VPI value is in use.
//   Otherwise, the agent creates a row and
//   reserves the VPI value on that port.
// 
// (2) The manager selects an existing row(s) in the
//   atmTrafficDescrParamTable,
//   thereby, selecting a set of self-consistent
//   ATM traffic parameters and the service category
//   for receive and transmit directions of the VPL.
// 
// (2a) If no suitable row(s) in the
//   atmTrafficDescrParamTable exists,
//   the manager must create a new row(s)
//   in that table.
// 
// (2b) The manager characterizes the VPL's traffic
//   parameters through setting the
//   atmVplReceiveTrafficDescrIndex and the
//   atmVplTransmitTrafficDescrIndex values
//   in the VPL table, which point to the rows
//   containing desired ATM traffic parameter values
//   in the atmTrafficDescrParamTable.  The agent
//   will check the availability of resources and
//   may refuse the request.
//   If the transmit and receive service categories
//   are inconsistent, the agent should refuse the
//   request.
// 
// (3) The manager activates the VPL by setting the
//   the atmVplRowStatus to active(1).
//   If this set is successful, the agent has
//   reserved the resources to satisfy the requested
//   traffic parameter values and the service category
//   for that VPL.
// 
// (4) If the VPL terminates a VPC in the ATM host
//   or switch, the manager turns on the
//   atmVplAdminStatus to up(1) to turn the VPL
//   traffic flow on.  Otherwise, the
//   atmVpCrossConnectTable  must be used
//   to cross-connect the VPL to another VPL(s)
//   in an ATM switch or network.
// 
// One-Shot VPL Establishment
// 
// A VPL may also be established in one step by a
// set-request with all necessary VPL parameter
// values and atmVplRowStatus set to createAndGo(4).
// 
// In contrast to the negotiated VPL establishment
// which allows for detailed error checking
// (i.e., set errors are explicitly linked to
// particular resource acquisition failures),
// the one-shot VPL establishment
// performs the setup on one operation but
// does not have the advantage of step-wise
// error checking.
// 
// VPL Retirement
// 
// A VPL is released by setting atmVplRowStatus to
// destroy(6), and the agent may release all
// associated resources.
type ATMMIB_AtmVplTable_AtmVplEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The VPI value of the VPL. The type is interface{}
    // with range: 0..4095.
    AtmVplVpi interface{}

    // This object is instanciated only for a VPL which terminates a VPC (i.e.,
    // one which is NOT cross-connected to other VPLs). Its value specifies the
    // desired administrative state of the VPL. The type is AtmVorXAdminStatus.
    AtmVplAdminStatus interface{}

    // The current operational status of the VPL. The type is AtmVorXOperStatus.
    AtmVplOperStatus interface{}

    // The value of sysUpTime at the time this VPL entered its current operational
    // state. The type is interface{} with range: 0..4294967295.
    AtmVplLastChange interface{}

    // The value of this object identifies the row in the
    // atmTrafficDescrParamTable which applies to the receive direction of the
    // VPL. The type is interface{} with range: 0..2147483647.
    AtmVplReceiveTrafficDescrIndex interface{}

    // The value of this object identifies the row in the
    // atmTrafficDescrParamTable which applies to the transmit direction of the
    // VPL. The type is interface{} with range: 0..2147483647.
    AtmVplTransmitTrafficDescrIndex interface{}

    // This object is instantiated only for a VPL which is cross-connected to
    // other VPLs that belong to the same VPC.  All such associated VPLs have the
    // same value of this object, and all their cross-connections are identified
    // either by entries that are indexed by the same value of
    // atmVpCrossConnectIndex in the atmVpCrossConnectTable of this MIB module or
    // by the same value of the cross-connect index in the cross-connect table for
    // SVCs and Soft PVCs (defined in a separate MIB module). At no time should
    // entries in these respective cross-connect tables exist simultaneously with
    // the same cross-connect index value. The value of this object is initialized
    // by the agent after the associated entries in the atmVpCrossConnectTable
    // have been created. The type is interface{} with range: 0..2147483647.
    AtmVplCrossConnectIdentifier interface{}

    // This object is used to create, delete or modify a row in this table. To
    // create a new VCL, this object is initially set to 'createAndWait' or
    // 'createAndGo'.  This object should not be set to 'active' unless the
    // following columnar objects have been set to their desired value in this
    // row: atmVplReceiveTrafficDescrIndex and atmVplTransmitTrafficDescrIndex.
    // The DESCRIPTION of atmVplEntry provides further guidance to row treatment
    // in this table. The type is RowStatus.
    AtmVplRowStatus interface{}

    // The connection topology type. The type is AtmConnCastType.
    AtmVplCastType interface{}

    // The use of call control. The type is AtmConnKind.
    AtmVplConnKind interface{}
}

func (atmVplEntry *ATMMIB_AtmVplTable_AtmVplEntry) GetEntityData() *types.CommonEntityData {
    atmVplEntry.EntityData.YFilter = atmVplEntry.YFilter
    atmVplEntry.EntityData.YangName = "atmVplEntry"
    atmVplEntry.EntityData.BundleName = "cisco_ios_xe"
    atmVplEntry.EntityData.ParentYangName = "atmVplTable"
    atmVplEntry.EntityData.SegmentPath = "atmVplEntry" + types.AddKeyToken(atmVplEntry.IfIndex, "ifIndex") + types.AddKeyToken(atmVplEntry.AtmVplVpi, "atmVplVpi")
    atmVplEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmVplTable/" + atmVplEntry.EntityData.SegmentPath
    atmVplEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVplEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVplEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVplEntry.EntityData.Children = types.NewOrderedMap()
    atmVplEntry.EntityData.Leafs = types.NewOrderedMap()
    atmVplEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmVplEntry.IfIndex})
    atmVplEntry.EntityData.Leafs.Append("atmVplVpi", types.YLeaf{"AtmVplVpi", atmVplEntry.AtmVplVpi})
    atmVplEntry.EntityData.Leafs.Append("atmVplAdminStatus", types.YLeaf{"AtmVplAdminStatus", atmVplEntry.AtmVplAdminStatus})
    atmVplEntry.EntityData.Leafs.Append("atmVplOperStatus", types.YLeaf{"AtmVplOperStatus", atmVplEntry.AtmVplOperStatus})
    atmVplEntry.EntityData.Leafs.Append("atmVplLastChange", types.YLeaf{"AtmVplLastChange", atmVplEntry.AtmVplLastChange})
    atmVplEntry.EntityData.Leafs.Append("atmVplReceiveTrafficDescrIndex", types.YLeaf{"AtmVplReceiveTrafficDescrIndex", atmVplEntry.AtmVplReceiveTrafficDescrIndex})
    atmVplEntry.EntityData.Leafs.Append("atmVplTransmitTrafficDescrIndex", types.YLeaf{"AtmVplTransmitTrafficDescrIndex", atmVplEntry.AtmVplTransmitTrafficDescrIndex})
    atmVplEntry.EntityData.Leafs.Append("atmVplCrossConnectIdentifier", types.YLeaf{"AtmVplCrossConnectIdentifier", atmVplEntry.AtmVplCrossConnectIdentifier})
    atmVplEntry.EntityData.Leafs.Append("atmVplRowStatus", types.YLeaf{"AtmVplRowStatus", atmVplEntry.AtmVplRowStatus})
    atmVplEntry.EntityData.Leafs.Append("atmVplCastType", types.YLeaf{"AtmVplCastType", atmVplEntry.AtmVplCastType})
    atmVplEntry.EntityData.Leafs.Append("atmVplConnKind", types.YLeaf{"AtmVplConnKind", atmVplEntry.AtmVplConnKind})

    atmVplEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVplVpi"}

    return &(atmVplEntry.EntityData)
}

// ATMMIB_AtmVclTable
// The Virtual Channel Link (VCL) table.  A
// bi-directional VCL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
type ATMMIB_AtmVclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the VCL table. This entry is used to model a bi-directional
    // VCL. To create a VCL at an ATM interface, either of the following
    // procedures are used:  Negotiated VCL establishment  (1) The management
    // application creates   a VCL entry in the atmVclTable   by setting
    // atmVclRowStatus to createAndWait(5).   This may fail for the following
    // reasons:   - The selected VPI/VCI values are unavailable,   - The selected
    // VPI/VCI values are in use.   Otherwise, the agent creates a row and  
    // reserves the VPI/VCI values on that port.  (2) The manager selects an
    // existing row(s) in the   atmTrafficDescrParamTable,   thereby, selecting a
    // set of self-consistent   ATM traffic parameters and the service category  
    // for receive and transmit directions of the VCL.   (2a) If no suitable
    // row(s) in the   atmTrafficDescrParamTable exists,   the manager must create
    // a new row(s)   in that table.  (2b) The manager characterizes the VCL's
    // traffic   parameters through setting the   atmVclReceiveTrafficDescrIndex
    // and the   atmVclTransmitTrafficDescrIndex values   in the VCL table, which
    // point to the rows   containing desired ATM traffic parameter values   in
    // the atmTrafficDescrParamTable.  The agent   will check the availability of
    // resources and   may refuse the request.   If the transmit and receive
    // service categories   are inconsistent, the agent should refuse the  
    // request.  (3) The manager activates the VCL by setting the   the
    // atmVclRowStatus to active(1) (for   requirements on this activation see the
    // description of atmVclRowStatus).   If this set is successful, the agent has
    // reserved the resources to satisfy the requested   traffic parameter values
    // and the service category   for that VCL. (4) If the VCL terminates a VCC in
    // the ATM host   or switch, the manager turns on the   atmVclAdminStatus to
    // up(1) to turn the VCL   traffic flow on.  Otherwise, the  
    // atmVcCrossConnectTable  must be used   to cross-connect the VCL to another
    // VCL(s)   in an ATM switch or network.  One-Shot VCL Establishment  A VCL
    // may also be established in one step by a set-request with all necessary VCL
    // parameter values and atmVclRowStatus set to createAndGo(4).  In contrast to
    // the negotiated VCL establishment which allows for detailed error checking
    // (i.e., set errors are explicitly linked to particular resource acquisition
    // failures), the one-shot VCL establishment performs the setup on one
    // operation but does not have the advantage of step-wise error checking.  VCL
    // Retirement  A VCL is released by setting atmVclRowStatus to destroy(6), and
    // the agent may release all associated resources. The type is slice of
    // ATMMIB_AtmVclTable_AtmVclEntry.
    AtmVclEntry []*ATMMIB_AtmVclTable_AtmVclEntry
}

func (atmVclTable *ATMMIB_AtmVclTable) GetEntityData() *types.CommonEntityData {
    atmVclTable.EntityData.YFilter = atmVclTable.YFilter
    atmVclTable.EntityData.YangName = "atmVclTable"
    atmVclTable.EntityData.BundleName = "cisco_ios_xe"
    atmVclTable.EntityData.ParentYangName = "ATM-MIB"
    atmVclTable.EntityData.SegmentPath = "atmVclTable"
    atmVclTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmVclTable.EntityData.SegmentPath
    atmVclTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVclTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVclTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVclTable.EntityData.Children = types.NewOrderedMap()
    atmVclTable.EntityData.Children.Append("atmVclEntry", types.YChild{"AtmVclEntry", nil})
    for i := range atmVclTable.AtmVclEntry {
        atmVclTable.EntityData.Children.Append(types.GetSegmentPath(atmVclTable.AtmVclEntry[i]), types.YChild{"AtmVclEntry", atmVclTable.AtmVclEntry[i]})
    }
    atmVclTable.EntityData.Leafs = types.NewOrderedMap()

    atmVclTable.EntityData.YListKeys = []string {}

    return &(atmVclTable.EntityData)
}

// ATMMIB_AtmVclTable_AtmVclEntry
// An entry in the VCL table. This entry is
// used to model a bi-directional VCL.
// To create a VCL at an ATM interface,
// either of the following procedures are used:
// 
// Negotiated VCL establishment
// 
// (1) The management application creates
//   a VCL entry in the atmVclTable
//   by setting atmVclRowStatus to createAndWait(5).
//   This may fail for the following reasons:
//   - The selected VPI/VCI values are unavailable,
//   - The selected VPI/VCI values are in use.
//   Otherwise, the agent creates a row and
//   reserves the VPI/VCI values on that port.
// 
// (2) The manager selects an existing row(s) in the
//   atmTrafficDescrParamTable,
//   thereby, selecting a set of self-consistent
//   ATM traffic parameters and the service category
//   for receive and transmit directions of the VCL.
// 
// 
// (2a) If no suitable row(s) in the
//   atmTrafficDescrParamTable exists,
//   the manager must create a new row(s)
//   in that table.
// 
// (2b) The manager characterizes the VCL's traffic
//   parameters through setting the
//   atmVclReceiveTrafficDescrIndex and the
//   atmVclTransmitTrafficDescrIndex values
//   in the VCL table, which point to the rows
//   containing desired ATM traffic parameter values
//   in the atmTrafficDescrParamTable.  The agent
//   will check the availability of resources and
//   may refuse the request.
//   If the transmit and receive service categories
//   are inconsistent, the agent should refuse the
//   request.
// 
// (3) The manager activates the VCL by setting the
//   the atmVclRowStatus to active(1) (for
//   requirements on this activation see the
//   description of atmVclRowStatus).
//   If this set is successful, the agent has
//   reserved the resources to satisfy the requested
//   traffic parameter values and the service category
//   for that VCL.
// (4) If the VCL terminates a VCC in the ATM host
//   or switch, the manager turns on the
//   atmVclAdminStatus to up(1) to turn the VCL
//   traffic flow on.  Otherwise, the
//   atmVcCrossConnectTable  must be used
//   to cross-connect the VCL to another VCL(s)
//   in an ATM switch or network.
// 
// One-Shot VCL Establishment
// 
// A VCL may also be established in one step by a
// set-request with all necessary VCL parameter
// values and atmVclRowStatus set to createAndGo(4).
// 
// In contrast to the negotiated VCL establishment
// which allows for detailed error checking
// (i.e., set errors are explicitly linked to
// particular resource acquisition failures),
// the one-shot VCL establishment
// performs the setup on one operation but
// does not have the advantage of step-wise
// error checking.
// 
// VCL Retirement
// 
// A VCL is released by setting atmVclRowStatus to
// destroy(6), and the agent may release all
// associated resources.
type ATMMIB_AtmVclTable_AtmVclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The VPI value of the VCL. The type is interface{}
    // with range: 0..4095.
    AtmVclVpi interface{}

    // This attribute is a key. The VCI value of the VCL. The type is interface{}
    // with range: 0..65535.
    AtmVclVci interface{}

    // This object is instanciated only for a VCL which terminates a VCC (i.e.,
    // one which is NOT cross-connected to other VCLs). Its value specifies the
    // desired administrative state of the VCL. The type is AtmVorXAdminStatus.
    AtmVclAdminStatus interface{}

    // The current operational status of the VCL. The type is AtmVorXOperStatus.
    AtmVclOperStatus interface{}

    // The value of sysUpTime at the time this VCL entered its current operational
    // state. The type is interface{} with range: 0..4294967295.
    AtmVclLastChange interface{}

    // The value of this object identifies the row in the ATM Traffic Descriptor
    // Table which applies to the receive direction of this VCL. The type is
    // interface{} with range: 0..2147483647.
    AtmVclReceiveTrafficDescrIndex interface{}

    // The value of this object identifies the row of the ATM Traffic Descriptor
    // Table which applies to the transmit direction of this VCL. The type is
    // interface{} with range: 0..2147483647.
    AtmVclTransmitTrafficDescrIndex interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL is in use. The type of AAL used on this VCC. The
    // AAL type includes AAL1, AAL2, AAL3/4, and AAL5. The other(4) may be
    // user-defined AAL type.  The unknown type indicates that the AAL type cannot
    // be determined. The type is AtmVccAalType.
    AtmVccAalType interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The maximum AAL5 CPCS SDU size in
    // octets that is supported on the transmit direction of this VCC. The type is
    // interface{} with range: 1..65535.
    AtmVccAal5CpcsTransmitSduSize interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The maximum AAL5 CPCS SDU size in
    // octets that is supported on the receive direction of this VCC. The type is
    // interface{} with range: 1..65535.
    AtmVccAal5CpcsReceiveSduSize interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The type of data encapsulation used
    // over the AAL5 SSCS layer. The definitions reference RFC 1483 Multiprotocol
    // Encapsulation over ATM AAL5 and to the ATM Forum LAN Emulation
    // specification. The type is AtmVccAal5EncapsType.
    AtmVccAal5EncapsType interface{}

    // This object is instantiated only for a VCL which is cross-connected to
    // other VCLs that belong to the same VCC.  All such associated VCLs have the
    // same value of this object, and all their cross-connections are identified
    // either by entries that are indexed by the same value of
    // atmVcCrossConnectIndex in the atmVcCrossConnectTable of this MIB module or
    // by the same value of the cross-connect index in the cross-connect table for
    // SVCs and Soft PVCs (defined in a separate MIB module).  At no time should
    // entries in these respective cross-connect tables exist simultaneously with
    // the same cross-connect index value. The value of this object is initialized
    // by the agent after the associated entries in the atmVcCrossConnectTable
    // have been created. The type is interface{} with range: 0..2147483647.
    AtmVclCrossConnectIdentifier interface{}

    // This object is used to create, delete or modify a row in this table.  To
    // create a new VCL, this object is initially set to 'createAndWait' or
    // 'createAndGo'. This object should not be set to 'active' unless the
    // following columnar objects have been set to their desired value in this
    // row: atmVclReceiveTrafficDescrIndex, atmVclTransmitTrafficDescrIndex. In
    // addition, if the local VCL end-point is also the VCC end-point:
    // atmVccAalType. In addition, for AAL5 connections only:
    // atmVccAal5CpcsTransmitSduSize, atmVccAal5CpcsReceiveSduSize, and
    // atmVccAal5EncapsType. (The existence of these objects imply the AAL
    // connection type.). The DESCRIPTION of atmVclEntry provides further guidance
    // to row treatment in this table. The type is RowStatus.
    AtmVclRowStatus interface{}

    // The connection topology type. The type is AtmConnCastType.
    AtmVclCastType interface{}

    // The use of call control. The type is AtmConnKind.
    AtmVclConnKind interface{}

    // Specifies OAM loopback frequency. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    CatmxVclOamLoopbackFreq interface{}

    // Specifies OAM retry polling frequency. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    CatmxVclOamRetryFreq interface{}

    // Specifies OAM retry count before declaring a VC  is up. The type is
    // interface{} with range: 0..4294967295.
    CatmxVclOamUpRetryCount interface{}

    // Specifies OAM retry count before declaring a VC  is down. The type is
    // interface{} with range: 0..4294967295.
    CatmxVclOamDownRetryCount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Activation retry count. The
    // type is interface{} with range: 0..4294967295.
    CatmxVclOamEndCCActCount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Deactivation retry count.
    // The type is interface{} with range: 0..4294967295.
    CatmxVclOamEndCCDeActCount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Activation/Deactivation
    // retry frequency. The type is interface{} with range: 0..4294967295. Units
    // are seconds.
    CatmxVclOamEndCCRetryFreq interface{}

    // Specifies OAM Segment Continuity check (CC)  Activation retry count. The
    // type is interface{} with range: 0..4294967295.
    CatmxVclOamSegCCActCount interface{}

    // Specifies OAM Segment Continuity check (CC)  Deactivation retry count. The
    // type is interface{} with range: 0..4294967295.
    CatmxVclOamSegCCDeActCount interface{}

    // Specifies OAM Segment Continuity check (CC)  Activation/Deactivation retry
    // frequency. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    CatmxVclOamSegCCRetryFreq interface{}

    // Specifies OAM Enable/Disable on the VC. true(1) indicates that OAM is
    // enabled on the VC. false(2) indicates that OAM is disabled on the VC. The
    // type is bool.
    CatmxVclOamManage interface{}

    // Indicates OAM loopback status of the VC. disabled(1)  --   No OAMs on this
    // VC. sent(2)      --   OAM sent, waiting for echo. received(3)  --   OAM
    // received from target. failed(4)    --   Last OAM did not return. The type
    // is CatmxVclOamLoopBkStatus.
    CatmxVclOamLoopBkStatus interface{}

    // Indicates the state of VC OAM. downRetry(1)   --  Loopback failed. Retry
    // sending                     loopbacks with retry frequency.                
    // VC is up. verified(2)    --  Loopback is successful. notVerified(3) --  Not
    // verified by loopback,                     AIS/RDI conditions are cleared.
    // upRetry(4)     --  Retry successive loopbacks.                     VC is
    // down. aisRDI(5)      --  Received AIS/RDI. Loopback are                    
    // not sent in this state. aisOut(6)      --  Sending AIS. Loopback and reply
    // are                     not sent in this state. notManaged(7)  --  VC is
    // not managed by OAM. The type is CatmxVclOamVcState.
    CatmxVclOamVcState interface{}

    // Indicates OAM End-to-end Continuity check (CC)  status. The type is
    // OamCCStatus.
    CatmxVclOamEndCCStatus interface{}

    // Indicates OAM Segment Continuity check (CC) status. The type is
    // OamCCStatus.
    CatmxVclOamSegCCStatus interface{}

    // Indicates OAM End-to-end Continuity check (CC)  VC state. The type is
    // OamCCVcState.
    CatmxVclOamEndCCVcState interface{}

    // Indicates OAM Segment Continuity check (CC) VC  state. The type is
    // OamCCVcState.
    CatmxVclOamSegCCVcState interface{}

    // Indicates the number of OAM cells received on  this VC. The type is
    // interface{} with range: 0..4294967295. Units are cells.
    CatmxVclOamCellsReceived interface{}

    // Indicates the number of OAM cells sent on  this VC. The type is interface{}
    // with range: 0..4294967295. Units are cells.
    CatmxVclOamCellsSent interface{}

    // Indicates the number of OAM cells dropped on  this VC. The type is
    // interface{} with range: 0..4294967295. Units are cells.
    CatmxVclOamCellsDropped interface{}

    // Indicates the number of received OAM  F5 Alarm Indication Signal (AIS)
    // cells from the VC. The type is interface{} with range: 0..4294967295. Units
    // are cells.
    CatmxVclOamInF5ais interface{}

    // Indicates the number of transmitted OAM  F5 Alarm Indication Signal (AIS)
    // cells to the VC. The type is interface{} with range: 0..4294967295. Units
    // are cells.
    CatmxVclOamOutF5ais interface{}

    // Indicates the number of received OAM  F5 Remote Detection Indication (RDI)
    // cells from  the VC. The type is interface{} with range: 0..4294967295.
    // Units are cells.
    CatmxVclOamInF5rdi interface{}

    // Indicates the number of transmitted OAM  F5 Remote Detection Indication
    // (RDI) cells to the VC. The type is interface{} with range: 0..4294967295.
    // Units are cells.
    CatmxVclOamOutF5rdi interface{}
}

func (atmVclEntry *ATMMIB_AtmVclTable_AtmVclEntry) GetEntityData() *types.CommonEntityData {
    atmVclEntry.EntityData.YFilter = atmVclEntry.YFilter
    atmVclEntry.EntityData.YangName = "atmVclEntry"
    atmVclEntry.EntityData.BundleName = "cisco_ios_xe"
    atmVclEntry.EntityData.ParentYangName = "atmVclTable"
    atmVclEntry.EntityData.SegmentPath = "atmVclEntry" + types.AddKeyToken(atmVclEntry.IfIndex, "ifIndex") + types.AddKeyToken(atmVclEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(atmVclEntry.AtmVclVci, "atmVclVci")
    atmVclEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmVclTable/" + atmVclEntry.EntityData.SegmentPath
    atmVclEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVclEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVclEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVclEntry.EntityData.Children = types.NewOrderedMap()
    atmVclEntry.EntityData.Leafs = types.NewOrderedMap()
    atmVclEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmVclEntry.IfIndex})
    atmVclEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", atmVclEntry.AtmVclVpi})
    atmVclEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", atmVclEntry.AtmVclVci})
    atmVclEntry.EntityData.Leafs.Append("atmVclAdminStatus", types.YLeaf{"AtmVclAdminStatus", atmVclEntry.AtmVclAdminStatus})
    atmVclEntry.EntityData.Leafs.Append("atmVclOperStatus", types.YLeaf{"AtmVclOperStatus", atmVclEntry.AtmVclOperStatus})
    atmVclEntry.EntityData.Leafs.Append("atmVclLastChange", types.YLeaf{"AtmVclLastChange", atmVclEntry.AtmVclLastChange})
    atmVclEntry.EntityData.Leafs.Append("atmVclReceiveTrafficDescrIndex", types.YLeaf{"AtmVclReceiveTrafficDescrIndex", atmVclEntry.AtmVclReceiveTrafficDescrIndex})
    atmVclEntry.EntityData.Leafs.Append("atmVclTransmitTrafficDescrIndex", types.YLeaf{"AtmVclTransmitTrafficDescrIndex", atmVclEntry.AtmVclTransmitTrafficDescrIndex})
    atmVclEntry.EntityData.Leafs.Append("atmVccAalType", types.YLeaf{"AtmVccAalType", atmVclEntry.AtmVccAalType})
    atmVclEntry.EntityData.Leafs.Append("atmVccAal5CpcsTransmitSduSize", types.YLeaf{"AtmVccAal5CpcsTransmitSduSize", atmVclEntry.AtmVccAal5CpcsTransmitSduSize})
    atmVclEntry.EntityData.Leafs.Append("atmVccAal5CpcsReceiveSduSize", types.YLeaf{"AtmVccAal5CpcsReceiveSduSize", atmVclEntry.AtmVccAal5CpcsReceiveSduSize})
    atmVclEntry.EntityData.Leafs.Append("atmVccAal5EncapsType", types.YLeaf{"AtmVccAal5EncapsType", atmVclEntry.AtmVccAal5EncapsType})
    atmVclEntry.EntityData.Leafs.Append("atmVclCrossConnectIdentifier", types.YLeaf{"AtmVclCrossConnectIdentifier", atmVclEntry.AtmVclCrossConnectIdentifier})
    atmVclEntry.EntityData.Leafs.Append("atmVclRowStatus", types.YLeaf{"AtmVclRowStatus", atmVclEntry.AtmVclRowStatus})
    atmVclEntry.EntityData.Leafs.Append("atmVclCastType", types.YLeaf{"AtmVclCastType", atmVclEntry.AtmVclCastType})
    atmVclEntry.EntityData.Leafs.Append("atmVclConnKind", types.YLeaf{"AtmVclConnKind", atmVclEntry.AtmVclConnKind})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamLoopbackFreq", types.YLeaf{"CatmxVclOamLoopbackFreq", atmVclEntry.CatmxVclOamLoopbackFreq})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamRetryFreq", types.YLeaf{"CatmxVclOamRetryFreq", atmVclEntry.CatmxVclOamRetryFreq})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamUpRetryCount", types.YLeaf{"CatmxVclOamUpRetryCount", atmVclEntry.CatmxVclOamUpRetryCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamDownRetryCount", types.YLeaf{"CatmxVclOamDownRetryCount", atmVclEntry.CatmxVclOamDownRetryCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamEndCCActCount", types.YLeaf{"CatmxVclOamEndCCActCount", atmVclEntry.CatmxVclOamEndCCActCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamEndCCDeActCount", types.YLeaf{"CatmxVclOamEndCCDeActCount", atmVclEntry.CatmxVclOamEndCCDeActCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamEndCCRetryFreq", types.YLeaf{"CatmxVclOamEndCCRetryFreq", atmVclEntry.CatmxVclOamEndCCRetryFreq})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamSegCCActCount", types.YLeaf{"CatmxVclOamSegCCActCount", atmVclEntry.CatmxVclOamSegCCActCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamSegCCDeActCount", types.YLeaf{"CatmxVclOamSegCCDeActCount", atmVclEntry.CatmxVclOamSegCCDeActCount})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamSegCCRetryFreq", types.YLeaf{"CatmxVclOamSegCCRetryFreq", atmVclEntry.CatmxVclOamSegCCRetryFreq})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamManage", types.YLeaf{"CatmxVclOamManage", atmVclEntry.CatmxVclOamManage})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamLoopBkStatus", types.YLeaf{"CatmxVclOamLoopBkStatus", atmVclEntry.CatmxVclOamLoopBkStatus})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamVcState", types.YLeaf{"CatmxVclOamVcState", atmVclEntry.CatmxVclOamVcState})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamEndCCStatus", types.YLeaf{"CatmxVclOamEndCCStatus", atmVclEntry.CatmxVclOamEndCCStatus})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamSegCCStatus", types.YLeaf{"CatmxVclOamSegCCStatus", atmVclEntry.CatmxVclOamSegCCStatus})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamEndCCVcState", types.YLeaf{"CatmxVclOamEndCCVcState", atmVclEntry.CatmxVclOamEndCCVcState})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamSegCCVcState", types.YLeaf{"CatmxVclOamSegCCVcState", atmVclEntry.CatmxVclOamSegCCVcState})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamCellsReceived", types.YLeaf{"CatmxVclOamCellsReceived", atmVclEntry.CatmxVclOamCellsReceived})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamCellsSent", types.YLeaf{"CatmxVclOamCellsSent", atmVclEntry.CatmxVclOamCellsSent})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamCellsDropped", types.YLeaf{"CatmxVclOamCellsDropped", atmVclEntry.CatmxVclOamCellsDropped})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamInF5ais", types.YLeaf{"CatmxVclOamInF5ais", atmVclEntry.CatmxVclOamInF5ais})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamOutF5ais", types.YLeaf{"CatmxVclOamOutF5ais", atmVclEntry.CatmxVclOamOutF5ais})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamInF5rdi", types.YLeaf{"CatmxVclOamInF5rdi", atmVclEntry.CatmxVclOamInF5rdi})
    atmVclEntry.EntityData.Leafs.Append("catmxVclOamOutF5rdi", types.YLeaf{"CatmxVclOamOutF5rdi", atmVclEntry.CatmxVclOamOutF5rdi})

    atmVclEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(atmVclEntry.EntityData)
}

// ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType represents LAN Emulation specification.
type ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType string

const (
    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexRoutedProtocol ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexRoutedProtocol"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexBridgedProtocol8023 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexBridgedProtocol8023"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexBridgedProtocol8025 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexBridgedProtocol8025"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexBridgedProtocol8026 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexBridgedProtocol8026"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexLANemulation8023 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexLANemulation8023"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_vcMultiplexLANemulation8025 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "vcMultiplexLANemulation8025"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_llcEncapsulation ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "llcEncapsulation"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_multiprotocolFrameRelaySscs ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "multiprotocolFrameRelaySscs"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_other ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "other"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType_unknown ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAal5EncapsType = "unknown"
)

// ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType represents the AAL type cannot be determined.
type ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType string

const (
    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_aal1 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "aal1"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_aal34 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "aal34"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_aal5 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "aal5"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_other ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "other"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_unknown ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "unknown"

    ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType_aal2 ATMMIB_AtmVclTable_AtmVclEntry_AtmVccAalType = "aal2"
)

// ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus represents failed(4)    --   Last OAM did not return.
type ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus string

const (
    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus_disabled ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus = "disabled"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus_sent ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus = "sent"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus_received ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus = "received"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus_failed ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamLoopBkStatus = "failed"
)

// ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState represents notManaged(7)  --  VC is not managed by OAM.
type ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState string

const (
    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_downRetry ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "downRetry"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_verified ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "verified"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_notVerified ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "notVerified"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_upRetry ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "upRetry"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_aisRDI ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "aisRDI"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_aisOut ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "aisOut"

    ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState_notManaged ATMMIB_AtmVclTable_AtmVclEntry_CatmxVclOamVcState = "notManaged"
)

// ATMMIB_AtmVpCrossConnectTable
// The ATM VP Cross Connect table for PVCs.
// An entry in this table models two
// cross-connected VPLs.
// Each VPL must have its atmConnKind set
// to pvc(1).
type ATMMIB_AtmVpCrossConnectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ATM VP Cross Connect table. This entry is used to model a
    // bi-directional ATM VP cross-connect which cross-connects two VPLs. 
    // Step-wise Procedures to set up a VP Cross-connect  Once the entries in the
    // atmVplTable are created, the following procedures are used to cross-connect
    // the VPLs together.  (1) The manager obtains a unique   
    // atmVpCrossConnectIndex by reading the    atmVpCrossConnectIndexNext object.
    // (2) Next, the manager creates a set of one    or more rows in the ATM VP
    // Cross Connect    Table, one for each cross-connection between    two VPLs. 
    // Each row is indexed by the ATM    interface port numbers and VPI values of
    // the    two ends of that cross-connection.    This set of rows specifies the
    // topology of the    VPC cross-connect and is identified by a single    value
    // of atmVpCrossConnectIndex.  Negotiated VP Cross-Connect Establishment  (2a)
    // The manager creates a row in this table by    setting
    // atmVpCrossConnectRowStatus to    createAndWait(5).  The agent checks the   
    // requested topology and the mutual sanity of    the ATM traffic parameters
    // and    service categories, i.e., the row creation    fails if:    - the
    // requested topology is incompatible with      associated values of
    // atmVplCastType,    - the requested topology is not supported      by the
    // agent,    - the traffic/service category parameter values      associated
    // with the requested row are      incompatible with those of already existing
    // rows for this VP cross-connect.    [For example, for setting up    a
    // point-to-point VP cross-connect, the    ATM traffic parameters in the
    // receive direction    of a VPL at the low end of the cross-connect    must
    // equal to the traffic parameters in the    transmit direction of the other
    // VPL at the    high end of the cross-connect,    otherwise, the row creation
    // fails.]    The agent also checks for internal errors    in building the
    // cross-connect.     The atmVpCrossConnectIndex values in the   
    // corresponding atmVplTable rows are filled    in by the agent at this point.
    // (2b) The manager promotes the row in the    atmVpCrossConnectTable by
    // setting    atmVpCrossConnectRowStatus to active(1).  If    this set is
    // successful, the agent has reserved    the resources specified by the ATM
    // traffic    parameter and Service category values    for each direction of
    // the VP cross-connect    in an ATM switch or network.  (3) The manager sets
    // the    atmVpCrossConnectAdminStatus to up(1) in all    rows of this VP
    // cross-connect to turn the    traffic flow on.   One-Shot VP Cross-Connect
    // Establishment  A VP cross-connect may also be established in one step by a
    // set-request with all necessary parameter values and
    // atmVpCrossConnectRowStatus set to createAndGo(4).  In contrast to the
    // negotiated VP cross-connect establishment which allows for detailed error
    // checking (i.e., set errors are explicitly linked to particular resource
    // acquisition failures), the one-shot VP cross-connect establishment performs
    // the setup on one operation but does not have the advantage of step-wise
    // error checking.  VP Cross-Connect Retirement  A VP cross-connect identified
    // by a particular value of atmVpCrossConnectIndex is released by:  (1)
    // Setting atmVpCrossConnectRowStatus of all    rows identified by this value
    // of    atmVpCrossConnectIndex to destroy(6).    The agent may release all   
    // associated resources, and the    atmVpCrossConnectIndex values in the   
    // corresponding atmVplTable row are removed.    Note that a situation when
    // only a subset of    the associated rows are deleted corresponds    to a VP
    // topology change.  (2) After deletion of the appropriate   
    // atmVpCrossConnectEntries, the manager may    set atmVplRowStatus to
    // destroy(6) the    associated VPLs.  The agent releases    the resources and
    // removes the associated    rows in the atmVplTable.  VP Cross-connect
    // Reconfiguration  At the discretion of the agent, a VP cross-connect may be
    // reconfigured by adding and/or deleting leafs to/from the VP topology as per
    // the VP cross-connect establishment/retirement procedures. Reconfiguration
    // of traffic/service category parameter values requires release of the VP
    // cross-connect before those parameter values may by changed for individual
    // VPLs. The type is slice of
    // ATMMIB_AtmVpCrossConnectTable_AtmVpCrossConnectEntry.
    AtmVpCrossConnectEntry []*ATMMIB_AtmVpCrossConnectTable_AtmVpCrossConnectEntry
}

func (atmVpCrossConnectTable *ATMMIB_AtmVpCrossConnectTable) GetEntityData() *types.CommonEntityData {
    atmVpCrossConnectTable.EntityData.YFilter = atmVpCrossConnectTable.YFilter
    atmVpCrossConnectTable.EntityData.YangName = "atmVpCrossConnectTable"
    atmVpCrossConnectTable.EntityData.BundleName = "cisco_ios_xe"
    atmVpCrossConnectTable.EntityData.ParentYangName = "ATM-MIB"
    atmVpCrossConnectTable.EntityData.SegmentPath = "atmVpCrossConnectTable"
    atmVpCrossConnectTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmVpCrossConnectTable.EntityData.SegmentPath
    atmVpCrossConnectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVpCrossConnectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVpCrossConnectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVpCrossConnectTable.EntityData.Children = types.NewOrderedMap()
    atmVpCrossConnectTable.EntityData.Children.Append("atmVpCrossConnectEntry", types.YChild{"AtmVpCrossConnectEntry", nil})
    for i := range atmVpCrossConnectTable.AtmVpCrossConnectEntry {
        atmVpCrossConnectTable.EntityData.Children.Append(types.GetSegmentPath(atmVpCrossConnectTable.AtmVpCrossConnectEntry[i]), types.YChild{"AtmVpCrossConnectEntry", atmVpCrossConnectTable.AtmVpCrossConnectEntry[i]})
    }
    atmVpCrossConnectTable.EntityData.Leafs = types.NewOrderedMap()

    atmVpCrossConnectTable.EntityData.YListKeys = []string {}

    return &(atmVpCrossConnectTable.EntityData)
}

// ATMMIB_AtmVpCrossConnectTable_AtmVpCrossConnectEntry
// An entry in the ATM VP Cross Connect table.
// This entry is used to model a bi-directional
// ATM VP cross-connect which cross-connects
// two VPLs.
// 
// Step-wise Procedures to set up a VP Cross-connect
// 
// Once the entries in the atmVplTable are created,
// the following procedures are used
// to cross-connect the VPLs together.
// 
// (1) The manager obtains a unique
//    atmVpCrossConnectIndex by reading the
//    atmVpCrossConnectIndexNext object.
// 
// (2) Next, the manager creates a set of one
//    or more rows in the ATM VP Cross Connect
//    Table, one for each cross-connection between
//    two VPLs.  Each row is indexed by the ATM
//    interface port numbers and VPI values of the
//    two ends of that cross-connection.
//    This set of rows specifies the topology of the
//    VPC cross-connect and is identified by a single
//    value of atmVpCrossConnectIndex.
// 
// Negotiated VP Cross-Connect Establishment
// 
// (2a) The manager creates a row in this table by
//    setting atmVpCrossConnectRowStatus to
//    createAndWait(5).  The agent checks the
//    requested topology and the mutual sanity of
//    the ATM traffic parameters and
//    service categories, i.e., the row creation
//    fails if:
//    - the requested topology is incompatible with
//      associated values of atmVplCastType,
//    - the requested topology is not supported
//      by the agent,
//    - the traffic/service category parameter values
//      associated with the requested row are
//      incompatible with those of already existing
//      rows for this VP cross-connect.
//    [For example, for setting up
//    a point-to-point VP cross-connect, the
//    ATM traffic parameters in the receive direction
//    of a VPL at the low end of the cross-connect
//    must equal to the traffic parameters in the
//    transmit direction of the other VPL at the
//    high end of the cross-connect,
//    otherwise, the row creation fails.]
//    The agent also checks for internal errors
//    in building the cross-connect.
// 
//    The atmVpCrossConnectIndex values in the
//    corresponding atmVplTable rows are filled
//    in by the agent at this point.
// 
// (2b) The manager promotes the row in the
//    atmVpCrossConnectTable by setting
//    atmVpCrossConnectRowStatus to active(1).  If
//    this set is successful, the agent has reserved
//    the resources specified by the ATM traffic
//    parameter and Service category values
//    for each direction of the VP cross-connect
//    in an ATM switch or network.
// 
// (3) The manager sets the
//    atmVpCrossConnectAdminStatus to up(1) in all
//    rows of this VP cross-connect to turn the
//    traffic flow on.
// 
// 
// One-Shot VP Cross-Connect Establishment
// 
// A VP cross-connect may also be established in
// one step by a set-request with all necessary
// parameter values and atmVpCrossConnectRowStatus
// set to createAndGo(4).
// 
// In contrast to the negotiated VP cross-connect
// establishment which allows for detailed error
// checking (i.e., set errors are explicitly linked
// to particular resource acquisition failures),
// the one-shot VP cross-connect establishment
// performs the setup on one operation but does not
// have the advantage of step-wise error checking.
// 
// VP Cross-Connect Retirement
// 
// A VP cross-connect identified by a particular
// value of atmVpCrossConnectIndex is released by:
// 
// (1) Setting atmVpCrossConnectRowStatus of all
//    rows identified by this value of
//    atmVpCrossConnectIndex to destroy(6).
//    The agent may release all
//    associated resources, and the
//    atmVpCrossConnectIndex values in the
//    corresponding atmVplTable row are removed.
//    Note that a situation when only a subset of
//    the associated rows are deleted corresponds
//    to a VP topology change.
// 
// (2) After deletion of the appropriate
//    atmVpCrossConnectEntries, the manager may
//    set atmVplRowStatus to destroy(6) the
//    associated VPLs.  The agent releases
//    the resources and removes the associated
//    rows in the atmVplTable.
// 
// VP Cross-connect Reconfiguration
// 
// At the discretion of the agent, a VP
// cross-connect may be reconfigured by
// adding and/or deleting leafs to/from
// the VP topology as per the VP cross-connect
// establishment/retirement procedures.
// Reconfiguration of traffic/service category parameter
// values requires release of the VP cross-connect
// before those parameter values may by changed
// for individual VPLs.
type ATMMIB_AtmVpCrossConnectTable_AtmVpCrossConnectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique value to identify this VP cross-connect.
    // For each VPL associated with this cross-connect, the agent reports this
    // cross-connect index value in the atmVplCrossConnectIdentifier attribute of
    // the corresponding atmVplTable entries. The type is interface{} with range:
    // 1..2147483647.
    AtmVpCrossConnectIndex interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VP
    // cross-connect. The term low implies that this ATM interface has the
    // numerically lower ifIndex value than the other ATM interface identified in
    // the same atmVpCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    AtmVpCrossConnectLowIfIndex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VP cross-connect that is identified by atmVpCrossConnectLowIfIndex. The
    // type is interface{} with range: 0..4095.
    AtmVpCrossConnectLowVpi interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VP
    // cross-connect. The term high implies that this ATM interface has the
    // numerically higher ifIndex value than the  other ATM interface identified
    // in the same atmVpCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    AtmVpCrossConnectHighIfIndex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VP cross-connect that is identified by atmVpCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..4095.
    AtmVpCrossConnectHighVpi interface{}

    // The desired administrative status of this bi-directional VP cross-connect.
    // The type is AtmVorXAdminStatus.
    AtmVpCrossConnectAdminStatus interface{}

    // The operational status of the VP cross-connect in one direction; (i.e.,
    // from the low to high direction). The type is AtmVorXOperStatus.
    AtmVpCrossConnectL2HOperStatus interface{}

    // The operational status of the VP cross-connect in one direction; (i.e.,
    // from the high to low direction). The type is AtmVorXOperStatus.
    AtmVpCrossConnectH2LOperStatus interface{}

    // The value of sysUpTime at the time this VP cross-connect entered its
    // current operational state in the low to high direction. The type is
    // interface{} with range: 0..4294967295.
    AtmVpCrossConnectL2HLastChange interface{}

    // The value of sysUpTime at the time this VP cross-connect entered its
    // current operational in the high to low direction. The type is interface{}
    // with range: 0..4294967295.
    AtmVpCrossConnectH2LLastChange interface{}

    // The status of this entry in the atmVpCrossConnectTable.  This object is
    // used to create a cross-connect for cross-connecting VPLs which are created
    // using the atmVplTable or to change or delete an existing cross-connect.
    // This object must be initially set to `createAndWait' or 'createAndGo'. To
    // turn on a VP cross-connect, the atmVpCrossConnectAdminStatus is set to
    // `up'. The type is RowStatus.
    AtmVpCrossConnectRowStatus interface{}
}

func (atmVpCrossConnectEntry *ATMMIB_AtmVpCrossConnectTable_AtmVpCrossConnectEntry) GetEntityData() *types.CommonEntityData {
    atmVpCrossConnectEntry.EntityData.YFilter = atmVpCrossConnectEntry.YFilter
    atmVpCrossConnectEntry.EntityData.YangName = "atmVpCrossConnectEntry"
    atmVpCrossConnectEntry.EntityData.BundleName = "cisco_ios_xe"
    atmVpCrossConnectEntry.EntityData.ParentYangName = "atmVpCrossConnectTable"
    atmVpCrossConnectEntry.EntityData.SegmentPath = "atmVpCrossConnectEntry" + types.AddKeyToken(atmVpCrossConnectEntry.AtmVpCrossConnectIndex, "atmVpCrossConnectIndex") + types.AddKeyToken(atmVpCrossConnectEntry.AtmVpCrossConnectLowIfIndex, "atmVpCrossConnectLowIfIndex") + types.AddKeyToken(atmVpCrossConnectEntry.AtmVpCrossConnectLowVpi, "atmVpCrossConnectLowVpi") + types.AddKeyToken(atmVpCrossConnectEntry.AtmVpCrossConnectHighIfIndex, "atmVpCrossConnectHighIfIndex") + types.AddKeyToken(atmVpCrossConnectEntry.AtmVpCrossConnectHighVpi, "atmVpCrossConnectHighVpi")
    atmVpCrossConnectEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmVpCrossConnectTable/" + atmVpCrossConnectEntry.EntityData.SegmentPath
    atmVpCrossConnectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVpCrossConnectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVpCrossConnectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVpCrossConnectEntry.EntityData.Children = types.NewOrderedMap()
    atmVpCrossConnectEntry.EntityData.Leafs = types.NewOrderedMap()
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectIndex", types.YLeaf{"AtmVpCrossConnectIndex", atmVpCrossConnectEntry.AtmVpCrossConnectIndex})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectLowIfIndex", types.YLeaf{"AtmVpCrossConnectLowIfIndex", atmVpCrossConnectEntry.AtmVpCrossConnectLowIfIndex})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectLowVpi", types.YLeaf{"AtmVpCrossConnectLowVpi", atmVpCrossConnectEntry.AtmVpCrossConnectLowVpi})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectHighIfIndex", types.YLeaf{"AtmVpCrossConnectHighIfIndex", atmVpCrossConnectEntry.AtmVpCrossConnectHighIfIndex})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectHighVpi", types.YLeaf{"AtmVpCrossConnectHighVpi", atmVpCrossConnectEntry.AtmVpCrossConnectHighVpi})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectAdminStatus", types.YLeaf{"AtmVpCrossConnectAdminStatus", atmVpCrossConnectEntry.AtmVpCrossConnectAdminStatus})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectL2HOperStatus", types.YLeaf{"AtmVpCrossConnectL2HOperStatus", atmVpCrossConnectEntry.AtmVpCrossConnectL2HOperStatus})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectH2LOperStatus", types.YLeaf{"AtmVpCrossConnectH2LOperStatus", atmVpCrossConnectEntry.AtmVpCrossConnectH2LOperStatus})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectL2HLastChange", types.YLeaf{"AtmVpCrossConnectL2HLastChange", atmVpCrossConnectEntry.AtmVpCrossConnectL2HLastChange})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectH2LLastChange", types.YLeaf{"AtmVpCrossConnectH2LLastChange", atmVpCrossConnectEntry.AtmVpCrossConnectH2LLastChange})
    atmVpCrossConnectEntry.EntityData.Leafs.Append("atmVpCrossConnectRowStatus", types.YLeaf{"AtmVpCrossConnectRowStatus", atmVpCrossConnectEntry.AtmVpCrossConnectRowStatus})

    atmVpCrossConnectEntry.EntityData.YListKeys = []string {"AtmVpCrossConnectIndex", "AtmVpCrossConnectLowIfIndex", "AtmVpCrossConnectLowVpi", "AtmVpCrossConnectHighIfIndex", "AtmVpCrossConnectHighVpi"}

    return &(atmVpCrossConnectEntry.EntityData)
}

// ATMMIB_AtmVcCrossConnectTable
// The ATM VC Cross Connect table for PVCs.
// An entry in this table models two
// cross-connected VCLs.
// Each VCL must have its atmConnKind set
// to pvc(1).
type ATMMIB_AtmVcCrossConnectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the ATM VC Cross Connect table. This entry is used to model a
    // bi-directional ATM VC cross-connect cross-connecting two end points. 
    // Step-wise Procedures to set up a VC Cross-connect  Once the entries in the
    // atmVclTable are created, the following procedures are used to cross-connect
    // the VCLs together to form a VCC segment.  (1) The manager obtains a unique 
    // atmVcCrossConnectIndex by reading the    atmVcCrossConnectIndexNext object.
    // (2) Next, the manager creates a set of one    or more rows in the ATM VC
    // Cross Connect    Table, one for each cross-connection between    two VCLs. 
    // Each row is indexed by the ATM    interface port numbers and VPI/VCI values
    // of    the two ends of that cross-connection.    This set of rows specifies
    // the topology of the    VCC cross-connect and is identified by a single   
    // value of atmVcCrossConnectIndex.  Negotiated VC Cross-Connect Establishment
    // (2a) The manager creates a row in this table by    setting
    // atmVcCrossConnectRowStatus to    createAndWait(5).  The agent checks the   
    // requested topology and the mutual sanity of    the ATM traffic parameters
    // and    service categories, i.e., the row creation    fails if:    - the
    // requested topology is incompatible with      associated values of
    // atmVclCastType,    - the requested topology is not supported      by the
    // agent,    - the traffic/service category parameter values      associated
    // with the requested row are      incompatible with those of already existing
    // rows for this VC cross-connect.    [For example, for setting up    a
    // point-to-point VC cross-connect, the    ATM traffic parameters in the
    // receive direction    of a VCL at the low end of the cross-connect    must
    // equal to the traffic parameters in the    transmit direction of the other
    // VCL at the    high end of the cross-connect,    otherwise, the row creation
    // fails.]    The agent also checks for internal errors    in building the
    // cross-connect.     The atmVcCrossConnectIndex values in the   
    // corresponding atmVclTable rows are filled    in by the agent at this point.
    // (2b) The manager promotes the row in the    atmVcCrossConnectTable by
    // setting    atmVcCrossConnectRowStatus to active(1).  If    this set is
    // successful, the agent has reserved    the resources specified by the ATM
    // traffic    parameter and Service category values    for each direction of
    // the VC cross-connect    in an ATM switch or network.  (3) The manager sets
    // the    atmVcCrossConnectAdminStatus to up(1)    in all rows of this VC
    // cross-connect to    turn the traffic flow on.   One-Shot VC Cross-Connect
    // Establishment  A VC cross-connect may also be established in one step by a
    // set-request with all necessary parameter values and
    // atmVcCrossConnectRowStatus set to createAndGo(4).  In contrast to the
    // negotiated VC cross-connect establishment which allows for detailed error
    // checking i.e., set errors are explicitly linked to particular resource
    // acquisition failures), the one-shot VC cross-connect establishment performs
    // the setup on one operation but does not have the advantage of step-wise
    // error checking.  VC Cross-Connect Retirement  A VC cross-connect identified
    // by a particular value of atmVcCrossConnectIndex is released by:  (1)
    // Setting atmVcCrossConnectRowStatus of all rows    identified by this value
    // of    atmVcCrossConnectIndex to destroy(6).    The agent may release all   
    // associated resources, and the    atmVcCrossConnectIndex values in the   
    // corresponding atmVclTable row are removed.    Note that a situation when
    // only a subset of    the associated rows are deleted corresponds    to a VC
    // topology change.  (2) After deletion of the appropriate   
    // atmVcCrossConnectEntries, the manager may    set atmVclRowStatus to
    // destroy(6) the    associated VCLs.  The agent releases    the resources and
    // removes the associated    rows in the atmVclTable.  VC Cross-Connect
    // Reconfiguration  At the discretion of the agent, a VC cross-connect may be
    // reconfigured by adding and/or deleting leafs to/from the VC topology as per
    // the VC cross-connect establishment/retirement procedures. Reconfiguration
    // of traffic/service category parameter values requires release of the VC
    // cross-connect before those parameter values may by changed for individual
    // VCLs. The type is slice of
    // ATMMIB_AtmVcCrossConnectTable_AtmVcCrossConnectEntry.
    AtmVcCrossConnectEntry []*ATMMIB_AtmVcCrossConnectTable_AtmVcCrossConnectEntry
}

func (atmVcCrossConnectTable *ATMMIB_AtmVcCrossConnectTable) GetEntityData() *types.CommonEntityData {
    atmVcCrossConnectTable.EntityData.YFilter = atmVcCrossConnectTable.YFilter
    atmVcCrossConnectTable.EntityData.YangName = "atmVcCrossConnectTable"
    atmVcCrossConnectTable.EntityData.BundleName = "cisco_ios_xe"
    atmVcCrossConnectTable.EntityData.ParentYangName = "ATM-MIB"
    atmVcCrossConnectTable.EntityData.SegmentPath = "atmVcCrossConnectTable"
    atmVcCrossConnectTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + atmVcCrossConnectTable.EntityData.SegmentPath
    atmVcCrossConnectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVcCrossConnectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVcCrossConnectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVcCrossConnectTable.EntityData.Children = types.NewOrderedMap()
    atmVcCrossConnectTable.EntityData.Children.Append("atmVcCrossConnectEntry", types.YChild{"AtmVcCrossConnectEntry", nil})
    for i := range atmVcCrossConnectTable.AtmVcCrossConnectEntry {
        atmVcCrossConnectTable.EntityData.Children.Append(types.GetSegmentPath(atmVcCrossConnectTable.AtmVcCrossConnectEntry[i]), types.YChild{"AtmVcCrossConnectEntry", atmVcCrossConnectTable.AtmVcCrossConnectEntry[i]})
    }
    atmVcCrossConnectTable.EntityData.Leafs = types.NewOrderedMap()

    atmVcCrossConnectTable.EntityData.YListKeys = []string {}

    return &(atmVcCrossConnectTable.EntityData)
}

// ATMMIB_AtmVcCrossConnectTable_AtmVcCrossConnectEntry
// An entry in the ATM VC Cross Connect table.
// This entry is used to model a bi-directional ATM
// VC cross-connect cross-connecting two end points.
// 
// Step-wise Procedures to set up a VC Cross-connect
// 
// Once the entries in the atmVclTable are created,
// the following procedures are used
// to cross-connect the VCLs together to
// form a VCC segment.
// 
// (1) The manager obtains a unique
//    atmVcCrossConnectIndex by reading the
//    atmVcCrossConnectIndexNext object.
// 
// (2) Next, the manager creates a set of one
//    or more rows in the ATM VC Cross Connect
//    Table, one for each cross-connection between
//    two VCLs.  Each row is indexed by the ATM
//    interface port numbers and VPI/VCI values of
//    the two ends of that cross-connection.
//    This set of rows specifies the topology of the
//    VCC cross-connect and is identified by a single
//    value of atmVcCrossConnectIndex.
// 
// Negotiated VC Cross-Connect Establishment
// 
// (2a) The manager creates a row in this table by
//    setting atmVcCrossConnectRowStatus to
//    createAndWait(5).  The agent checks the
//    requested topology and the mutual sanity of
//    the ATM traffic parameters and
//    service categories, i.e., the row creation
//    fails if:
//    - the requested topology is incompatible with
//      associated values of atmVclCastType,
//    - the requested topology is not supported
//      by the agent,
//    - the traffic/service category parameter values
//      associated with the requested row are
//      incompatible with those of already existing
//      rows for this VC cross-connect.
//    [For example, for setting up
//    a point-to-point VC cross-connect, the
//    ATM traffic parameters in the receive direction
//    of a VCL at the low end of the cross-connect
//    must equal to the traffic parameters in the
//    transmit direction of the other VCL at the
//    high end of the cross-connect,
//    otherwise, the row creation fails.]
//    The agent also checks for internal errors
//    in building the cross-connect.
// 
//    The atmVcCrossConnectIndex values in the
//    corresponding atmVclTable rows are filled
//    in by the agent at this point.
// 
// (2b) The manager promotes the row in the
//    atmVcCrossConnectTable by setting
//    atmVcCrossConnectRowStatus to active(1).  If
//    this set is successful, the agent has reserved
//    the resources specified by the ATM traffic
//    parameter and Service category values
//    for each direction of the VC cross-connect
//    in an ATM switch or network.
// 
// (3) The manager sets the
//    atmVcCrossConnectAdminStatus to up(1)
//    in all rows of this VC cross-connect to
//    turn the traffic flow on.
// 
// 
// One-Shot VC Cross-Connect Establishment
// 
// A VC cross-connect may also be established in
// one step by a set-request with all necessary
// parameter values and atmVcCrossConnectRowStatus
// set to createAndGo(4).
// 
// In contrast to the negotiated VC cross-connect
// establishment which allows for detailed error
// checking i.e., set errors are explicitly linked to
// particular resource acquisition failures), the
// one-shot VC cross-connect establishment
// performs the setup on one operation but does
// not have the advantage of step-wise error
// checking.
// 
// VC Cross-Connect Retirement
// 
// A VC cross-connect identified by a particular
// value of atmVcCrossConnectIndex is released by:
// 
// (1) Setting atmVcCrossConnectRowStatus of all rows
//    identified by this value of
//    atmVcCrossConnectIndex to destroy(6).
//    The agent may release all
//    associated resources, and the
//    atmVcCrossConnectIndex values in the
//    corresponding atmVclTable row are removed.
//    Note that a situation when only a subset of
//    the associated rows are deleted corresponds
//    to a VC topology change.
// 
// (2) After deletion of the appropriate
//    atmVcCrossConnectEntries, the manager may
//    set atmVclRowStatus to destroy(6) the
//    associated VCLs.  The agent releases
//    the resources and removes the associated
//    rows in the atmVclTable.
// 
// VC Cross-Connect Reconfiguration
// 
// At the discretion of the agent, a VC
// cross-connect may be reconfigured by
// adding and/or deleting leafs to/from
// the VC topology as per the VC cross-connect
// establishment/retirement procedures.
// Reconfiguration of traffic/service category parameter
// values requires release of the VC cross-connect
// before those parameter values may by changed
// for individual VCLs.
type ATMMIB_AtmVcCrossConnectTable_AtmVcCrossConnectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique value to identify this VC cross-connect.
    // For each VCL associated with this cross-connect, the agent reports this
    // cross-connect index value in the atmVclCrossConnectIdentifier attribute of
    // the corresponding atmVclTable entries. The type is interface{} with range:
    // 1..2147483647.
    AtmVcCrossConnectIndex interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VC
    // cross-connect. The term low implies that this ATM interface has the
    // numerically lower ifIndex value than the other ATM interface identified in
    // the same atmVcCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    AtmVcCrossConnectLowIfIndex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectLowIfIndex. The
    // type is interface{} with range: 0..4095.
    AtmVcCrossConnectLowVpi interface{}

    // This attribute is a key. The VCI value at the ATM interface associated with
    // this VC cross-connect that is identified by atmVcCrossConnectLowIfIndex.
    // The type is interface{} with range: 0..65535.
    AtmVcCrossConnectLowVci interface{}

    // This attribute is a key. The ifIndex value for the ATM interface for this
    // VC cross-connect. The term high implies that this ATM interface has the
    // numerically higher ifIndex value than the other ATM interface identified in
    // the same atmVcCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    AtmVcCrossConnectHighIfIndex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..4095.
    AtmVcCrossConnectHighVpi interface{}

    // This attribute is a key. The VCI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..65535.
    AtmVcCrossConnectHighVci interface{}

    // The desired administrative status of this bi-directional VC cross-connect.
    // The type is AtmVorXAdminStatus.
    AtmVcCrossConnectAdminStatus interface{}

    // The current operational status of the VC cross-connect in one direction;
    // (i.e., from the low to high direction). The type is AtmVorXOperStatus.
    AtmVcCrossConnectL2HOperStatus interface{}

    // The current operational status of the VC cross-connect in one direction;
    // (i.e., from the high to low direction). The type is AtmVorXOperStatus.
    AtmVcCrossConnectH2LOperStatus interface{}

    // The value of sysUpTime at the time this VC cross-connect entered its
    // current operational state in low to high direction. The type is interface{}
    // with range: 0..4294967295.
    AtmVcCrossConnectL2HLastChange interface{}

    // The value of sysUpTime at the time this VC cross-connect entered its
    // current operational state in high to low direction. The type is interface{}
    // with range: 0..4294967295.
    AtmVcCrossConnectH2LLastChange interface{}

    // The status of this entry in the atmVcCrossConnectTable.  This object is
    // used to create a new cross-connect for cross-connecting VCLs which are
    // created using the atmVclTable or to change or delete existing
    // cross-connect. This object must be initially set to `createAndWait' or
    // 'createAndGo'. To turn on a VC cross-connect, the
    // atmVcCrossConnectAdminStatus is set to `up'. The type is RowStatus.
    AtmVcCrossConnectRowStatus interface{}
}

func (atmVcCrossConnectEntry *ATMMIB_AtmVcCrossConnectTable_AtmVcCrossConnectEntry) GetEntityData() *types.CommonEntityData {
    atmVcCrossConnectEntry.EntityData.YFilter = atmVcCrossConnectEntry.YFilter
    atmVcCrossConnectEntry.EntityData.YangName = "atmVcCrossConnectEntry"
    atmVcCrossConnectEntry.EntityData.BundleName = "cisco_ios_xe"
    atmVcCrossConnectEntry.EntityData.ParentYangName = "atmVcCrossConnectTable"
    atmVcCrossConnectEntry.EntityData.SegmentPath = "atmVcCrossConnectEntry" + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectIndex, "atmVcCrossConnectIndex") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectLowIfIndex, "atmVcCrossConnectLowIfIndex") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectLowVpi, "atmVcCrossConnectLowVpi") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectLowVci, "atmVcCrossConnectLowVci") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectHighIfIndex, "atmVcCrossConnectHighIfIndex") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectHighVpi, "atmVcCrossConnectHighVpi") + types.AddKeyToken(atmVcCrossConnectEntry.AtmVcCrossConnectHighVci, "atmVcCrossConnectHighVci")
    atmVcCrossConnectEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/atmVcCrossConnectTable/" + atmVcCrossConnectEntry.EntityData.SegmentPath
    atmVcCrossConnectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmVcCrossConnectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmVcCrossConnectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmVcCrossConnectEntry.EntityData.Children = types.NewOrderedMap()
    atmVcCrossConnectEntry.EntityData.Leafs = types.NewOrderedMap()
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectIndex", types.YLeaf{"AtmVcCrossConnectIndex", atmVcCrossConnectEntry.AtmVcCrossConnectIndex})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectLowIfIndex", types.YLeaf{"AtmVcCrossConnectLowIfIndex", atmVcCrossConnectEntry.AtmVcCrossConnectLowIfIndex})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectLowVpi", types.YLeaf{"AtmVcCrossConnectLowVpi", atmVcCrossConnectEntry.AtmVcCrossConnectLowVpi})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectLowVci", types.YLeaf{"AtmVcCrossConnectLowVci", atmVcCrossConnectEntry.AtmVcCrossConnectLowVci})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectHighIfIndex", types.YLeaf{"AtmVcCrossConnectHighIfIndex", atmVcCrossConnectEntry.AtmVcCrossConnectHighIfIndex})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectHighVpi", types.YLeaf{"AtmVcCrossConnectHighVpi", atmVcCrossConnectEntry.AtmVcCrossConnectHighVpi})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectHighVci", types.YLeaf{"AtmVcCrossConnectHighVci", atmVcCrossConnectEntry.AtmVcCrossConnectHighVci})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectAdminStatus", types.YLeaf{"AtmVcCrossConnectAdminStatus", atmVcCrossConnectEntry.AtmVcCrossConnectAdminStatus})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectL2HOperStatus", types.YLeaf{"AtmVcCrossConnectL2HOperStatus", atmVcCrossConnectEntry.AtmVcCrossConnectL2HOperStatus})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectH2LOperStatus", types.YLeaf{"AtmVcCrossConnectH2LOperStatus", atmVcCrossConnectEntry.AtmVcCrossConnectH2LOperStatus})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectL2HLastChange", types.YLeaf{"AtmVcCrossConnectL2HLastChange", atmVcCrossConnectEntry.AtmVcCrossConnectL2HLastChange})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectH2LLastChange", types.YLeaf{"AtmVcCrossConnectH2LLastChange", atmVcCrossConnectEntry.AtmVcCrossConnectH2LLastChange})
    atmVcCrossConnectEntry.EntityData.Leafs.Append("atmVcCrossConnectRowStatus", types.YLeaf{"AtmVcCrossConnectRowStatus", atmVcCrossConnectEntry.AtmVcCrossConnectRowStatus})

    atmVcCrossConnectEntry.EntityData.YListKeys = []string {"AtmVcCrossConnectIndex", "AtmVcCrossConnectLowIfIndex", "AtmVcCrossConnectLowVpi", "AtmVcCrossConnectLowVci", "AtmVcCrossConnectHighIfIndex", "AtmVcCrossConnectHighVpi", "AtmVcCrossConnectHighVci"}

    return &(atmVcCrossConnectEntry.EntityData)
}

// ATMMIB_Aal5VccTable
// This table contains AAL5 VCC performance
// parameters.
type ATMMIB_Aal5VccTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the AAL5 VCC performance parameters and is indexed by
    // ifIndex values of AAL5 interfaces and the associated VPI/VCI values. The
    // type is slice of ATMMIB_Aal5VccTable_Aal5VccEntry.
    Aal5VccEntry []*ATMMIB_Aal5VccTable_Aal5VccEntry
}

func (aal5VccTable *ATMMIB_Aal5VccTable) GetEntityData() *types.CommonEntityData {
    aal5VccTable.EntityData.YFilter = aal5VccTable.YFilter
    aal5VccTable.EntityData.YangName = "aal5VccTable"
    aal5VccTable.EntityData.BundleName = "cisco_ios_xe"
    aal5VccTable.EntityData.ParentYangName = "ATM-MIB"
    aal5VccTable.EntityData.SegmentPath = "aal5VccTable"
    aal5VccTable.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/" + aal5VccTable.EntityData.SegmentPath
    aal5VccTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aal5VccTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aal5VccTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aal5VccTable.EntityData.Children = types.NewOrderedMap()
    aal5VccTable.EntityData.Children.Append("aal5VccEntry", types.YChild{"Aal5VccEntry", nil})
    for i := range aal5VccTable.Aal5VccEntry {
        aal5VccTable.EntityData.Children.Append(types.GetSegmentPath(aal5VccTable.Aal5VccEntry[i]), types.YChild{"Aal5VccEntry", aal5VccTable.Aal5VccEntry[i]})
    }
    aal5VccTable.EntityData.Leafs = types.NewOrderedMap()

    aal5VccTable.EntityData.YListKeys = []string {}

    return &(aal5VccTable.EntityData)
}

// ATMMIB_Aal5VccTable_Aal5VccEntry
// This list contains the AAL5 VCC
// performance parameters and is indexed
// by ifIndex values of AAL5 interfaces
// and the associated VPI/VCI values.
type ATMMIB_Aal5VccTable_Aal5VccEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The VPI value of the AAL5 VCC at the interface
    // identified by the ifIndex. The type is interface{} with range: 0..4095.
    Aal5VccVpi interface{}

    // This attribute is a key. The VCI value of the AAL5 VCC at the interface
    // identified by the ifIndex. The type is interface{} with range: 0..65535.
    Aal5VccVci interface{}

    // The number of AAL5 CPCS PDUs received with CRC-32 errors on this AAL5 VCC
    // at the interface associated with an AAL5 entity. The type is interface{}
    // with range: 0..4294967295.
    Aal5VccCrcErrors interface{}

    // The number of partially re-assembled AAL5 CPCS PDUs which were discarded on
    // this AAL5 VCC at the interface associated with an AAL5 entity because they
    // were not fully re-assembled within the required time period.  If the
    // re-assembly timer is not supported, then this object contains a zero value.
    // The type is interface{} with range: 0..4294967295.
    Aal5VccSarTimeOuts interface{}

    // The number of AAL5 CPCS PDUs discarded on this AAL5 VCC at the interface
    // associated with an AAL5 entity because the AAL5 SDUs were too large. The
    // type is interface{} with range: 0..4294967295.
    Aal5VccOverSizedSDUs interface{}

    // The number of AAL5 CPCS PDUs received on this AAL5 VCC at the interface
    // associated with an AAL5 entity. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CAal5VccInPkts interface{}

    // The number of AAL5 CPCS PDUs transmitted on this AAL5 VCC at the interface
    // associated with an AAL5 entity. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    CAal5VccOutPkts interface{}

    // The number of AAL5 CPCS PDU octets received on this AAL5 VCC at the
    // interface associated with an AAL5 entity. The type is interface{} with
    // range: 0..4294967295. Units are octets.
    CAal5VccInOctets interface{}

    // The number of AAL5 CPCS PDU octets transmitted on this AAL5  VCC at the
    // interface associated with an AAL5 entity. The type is interface{} with
    // range: 0..4294967295. Units are octets.
    CAal5VccOutOctets interface{}

    // The number of AAL5 CPCS PDUs dropped at the  receive side of this AAL5 VCC
    // at the interface  associated with an AAL5 entity. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CAal5VccInDroppedPkts interface{}

    // The number of AAL5 CPCS PDUs dropped at the transmit side  of this AAL5 VCC
    // at the interface associated with an  AAL5 entity. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    CAal5VccOutDroppedPkts interface{}

    // The number of AAL5 CPCS PDU Octets dropped at the  receive side of this
    // AAL5 VCC at the interface  associated with an AAL5 entity. The type is
    // interface{} with range: 0..4294967295. Units are octets.
    CAal5VccInDroppedOctets interface{}

    // The number of AAL5 CPCS PDU Octets dropped at the  transmit side of this
    // AAL5 VCC at the interface  associated with an AAL5 entity. The type is
    // interface{} with range: 0..4294967295. Units are octets.
    CAal5VccOutDroppedOctets interface{}

    // This is 64bit (High Capacity) version of cAal5VccInPkts  counters. The type
    // is interface{} with range: 0..18446744073709551615.
    CAal5VccHCInPkts interface{}

    // This is 64bit (High Capacity) version of cAal5VccOutPkts  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    CAal5VccHCOutPkts interface{}

    // This is 64bit (High Capacity) version of cAal5VccInOctets  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    CAal5VccHCInOctets interface{}

    // This is 64bit (High Capacity) version of cAal5VccOutOctets  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    CAal5VccHCOutOctets interface{}

    // Boolean, if compression enabled for VCC. The type is bool.
    CAal5VccExtCompEnabled interface{}

    // Boolean, TRUE if VCC is used to carry voice. The type is bool.
    CAal5VccExtVoice interface{}

    // Number of OAM F5 end to end loopback cells  received through the VCC. The
    // type is interface{} with range: 0..4294967295.
    CAal5VccExtInF5OamCells interface{}

    // Number of OAM F5 end to end loopback cells sent  through the VCC. The type
    // is interface{} with range: 0..4294967295.
    CAal5VccExtOutF5OamCells interface{}
}

func (aal5VccEntry *ATMMIB_Aal5VccTable_Aal5VccEntry) GetEntityData() *types.CommonEntityData {
    aal5VccEntry.EntityData.YFilter = aal5VccEntry.YFilter
    aal5VccEntry.EntityData.YangName = "aal5VccEntry"
    aal5VccEntry.EntityData.BundleName = "cisco_ios_xe"
    aal5VccEntry.EntityData.ParentYangName = "aal5VccTable"
    aal5VccEntry.EntityData.SegmentPath = "aal5VccEntry" + types.AddKeyToken(aal5VccEntry.IfIndex, "ifIndex") + types.AddKeyToken(aal5VccEntry.Aal5VccVpi, "aal5VccVpi") + types.AddKeyToken(aal5VccEntry.Aal5VccVci, "aal5VccVci")
    aal5VccEntry.EntityData.AbsolutePath = "ATM-MIB:ATM-MIB/aal5VccTable/" + aal5VccEntry.EntityData.SegmentPath
    aal5VccEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aal5VccEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aal5VccEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aal5VccEntry.EntityData.Children = types.NewOrderedMap()
    aal5VccEntry.EntityData.Leafs = types.NewOrderedMap()
    aal5VccEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", aal5VccEntry.IfIndex})
    aal5VccEntry.EntityData.Leafs.Append("aal5VccVpi", types.YLeaf{"Aal5VccVpi", aal5VccEntry.Aal5VccVpi})
    aal5VccEntry.EntityData.Leafs.Append("aal5VccVci", types.YLeaf{"Aal5VccVci", aal5VccEntry.Aal5VccVci})
    aal5VccEntry.EntityData.Leafs.Append("aal5VccCrcErrors", types.YLeaf{"Aal5VccCrcErrors", aal5VccEntry.Aal5VccCrcErrors})
    aal5VccEntry.EntityData.Leafs.Append("aal5VccSarTimeOuts", types.YLeaf{"Aal5VccSarTimeOuts", aal5VccEntry.Aal5VccSarTimeOuts})
    aal5VccEntry.EntityData.Leafs.Append("aal5VccOverSizedSDUs", types.YLeaf{"Aal5VccOverSizedSDUs", aal5VccEntry.Aal5VccOverSizedSDUs})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccInPkts", types.YLeaf{"CAal5VccInPkts", aal5VccEntry.CAal5VccInPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccOutPkts", types.YLeaf{"CAal5VccOutPkts", aal5VccEntry.CAal5VccOutPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccInOctets", types.YLeaf{"CAal5VccInOctets", aal5VccEntry.CAal5VccInOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccOutOctets", types.YLeaf{"CAal5VccOutOctets", aal5VccEntry.CAal5VccOutOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccInDroppedPkts", types.YLeaf{"CAal5VccInDroppedPkts", aal5VccEntry.CAal5VccInDroppedPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccOutDroppedPkts", types.YLeaf{"CAal5VccOutDroppedPkts", aal5VccEntry.CAal5VccOutDroppedPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccInDroppedOctets", types.YLeaf{"CAal5VccInDroppedOctets", aal5VccEntry.CAal5VccInDroppedOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccOutDroppedOctets", types.YLeaf{"CAal5VccOutDroppedOctets", aal5VccEntry.CAal5VccOutDroppedOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccHCInPkts", types.YLeaf{"CAal5VccHCInPkts", aal5VccEntry.CAal5VccHCInPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccHCOutPkts", types.YLeaf{"CAal5VccHCOutPkts", aal5VccEntry.CAal5VccHCOutPkts})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccHCInOctets", types.YLeaf{"CAal5VccHCInOctets", aal5VccEntry.CAal5VccHCInOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccHCOutOctets", types.YLeaf{"CAal5VccHCOutOctets", aal5VccEntry.CAal5VccHCOutOctets})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccExtCompEnabled", types.YLeaf{"CAal5VccExtCompEnabled", aal5VccEntry.CAal5VccExtCompEnabled})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccExtVoice", types.YLeaf{"CAal5VccExtVoice", aal5VccEntry.CAal5VccExtVoice})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccExtInF5OamCells", types.YLeaf{"CAal5VccExtInF5OamCells", aal5VccEntry.CAal5VccExtInF5OamCells})
    aal5VccEntry.EntityData.Leafs.Append("cAal5VccExtOutF5OamCells", types.YLeaf{"CAal5VccExtOutF5OamCells", aal5VccEntry.CAal5VccExtOutF5OamCells})

    aal5VccEntry.EntityData.YListKeys = []string {"IfIndex", "Aal5VccVpi", "Aal5VccVci"}

    return &(aal5VccEntry.EntityData)
}

