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

    
    Atmmibobjects ATMMIB_Atmmibobjects

    // This table contains ATM local interface configuration parameters, one entry
    // per ATM interface port.
    Atminterfaceconftable ATMMIB_Atminterfaceconftable

    // This table contains ATM interface DS3 PLCP parameters and state variables,
    // one entry per ATM interface port.
    Atminterfaceds3Plcptable ATMMIB_Atminterfaceds3Plcptable

    // This table contains ATM interface TC Sublayer parameters and state
    // variables, one entry per ATM interface port.
    Atminterfacetctable ATMMIB_Atminterfacetctable

    // This table contains information on ATM traffic descriptor type and the
    // associated parameters.
    Atmtrafficdescrparamtable ATMMIB_Atmtrafficdescrparamtable

    // The Virtual Path Link (VPL) table.  A bi-directional VPL is modeled as one
    // entry in this table. This table can be used for PVCs, SVCs and Soft PVCs.
    // Entries are not present in this table for the VPIs used by entries in the
    // atmVclTable.
    Atmvpltable ATMMIB_Atmvpltable

    // The Virtual Channel Link (VCL) table.  A bi-directional VCL is modeled as
    // one entry in this table. This table can be used for PVCs, SVCs and Soft
    // PVCs.
    Atmvcltable ATMMIB_Atmvcltable

    // The ATM VP Cross Connect table for PVCs. An entry in this table models two
    // cross-connected VPLs. Each VPL must have its atmConnKind set to pvc(1).
    Atmvpcrossconnecttable ATMMIB_Atmvpcrossconnecttable

    // The ATM VC Cross Connect table for PVCs. An entry in this table models two
    // cross-connected VCLs. Each VCL must have its atmConnKind set to pvc(1).
    Atmvccrossconnecttable ATMMIB_Atmvccrossconnecttable

    // This table contains AAL5 VCC performance parameters.
    Aal5Vcctable ATMMIB_Aal5Vcctable
}

func (aTMMIB *ATMMIB) GetEntityData() *types.CommonEntityData {
    aTMMIB.EntityData.YFilter = aTMMIB.YFilter
    aTMMIB.EntityData.YangName = "ATM-MIB"
    aTMMIB.EntityData.BundleName = "cisco_ios_xe"
    aTMMIB.EntityData.ParentYangName = "ATM-MIB"
    aTMMIB.EntityData.SegmentPath = "ATM-MIB:ATM-MIB"
    aTMMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aTMMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aTMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aTMMIB.EntityData.Children = make(map[string]types.YChild)
    aTMMIB.EntityData.Children["atmMIBObjects"] = types.YChild{"Atmmibobjects", &aTMMIB.Atmmibobjects}
    aTMMIB.EntityData.Children["atmInterfaceConfTable"] = types.YChild{"Atminterfaceconftable", &aTMMIB.Atminterfaceconftable}
    aTMMIB.EntityData.Children["atmInterfaceDs3PlcpTable"] = types.YChild{"Atminterfaceds3Plcptable", &aTMMIB.Atminterfaceds3Plcptable}
    aTMMIB.EntityData.Children["atmInterfaceTCTable"] = types.YChild{"Atminterfacetctable", &aTMMIB.Atminterfacetctable}
    aTMMIB.EntityData.Children["atmTrafficDescrParamTable"] = types.YChild{"Atmtrafficdescrparamtable", &aTMMIB.Atmtrafficdescrparamtable}
    aTMMIB.EntityData.Children["atmVplTable"] = types.YChild{"Atmvpltable", &aTMMIB.Atmvpltable}
    aTMMIB.EntityData.Children["atmVclTable"] = types.YChild{"Atmvcltable", &aTMMIB.Atmvcltable}
    aTMMIB.EntityData.Children["atmVpCrossConnectTable"] = types.YChild{"Atmvpcrossconnecttable", &aTMMIB.Atmvpcrossconnecttable}
    aTMMIB.EntityData.Children["atmVcCrossConnectTable"] = types.YChild{"Atmvccrossconnecttable", &aTMMIB.Atmvccrossconnecttable}
    aTMMIB.EntityData.Children["aal5VccTable"] = types.YChild{"Aal5Vcctable", &aTMMIB.Aal5Vcctable}
    aTMMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aTMMIB.EntityData)
}

// ATMMIB_Atmmibobjects
type ATMMIB_Atmmibobjects struct {
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
    Atmvpcrossconnectindexnext interface{}

    // This object contains an appropriate value to be used for
    // atmVcCrossConnectIndex when creating entries in the atmVcCrossConnectTable.
    // The value 0 indicates that no unassigned entries are available. To obtain
    // the atmVcCrossConnectIndex value for a new entry, the manager issues a
    // management protocol retrieval operation to obtain the current value of this
    // object.  After each retrieval, the agent should modify the value to the
    // next unassigned index. After a manager retrieves a value the agent will
    // determine through its local policy when this index value will be made
    // available for reuse. The type is interface{} with range: 0..2147483647.
    Atmvccrossconnectindexnext interface{}

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
    Atmtrafficdescrparamindexnext interface{}
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetEntityData() *types.CommonEntityData {
    atmmibobjects.EntityData.YFilter = atmmibobjects.YFilter
    atmmibobjects.EntityData.YangName = "atmMIBObjects"
    atmmibobjects.EntityData.BundleName = "cisco_ios_xe"
    atmmibobjects.EntityData.ParentYangName = "ATM-MIB"
    atmmibobjects.EntityData.SegmentPath = "atmMIBObjects"
    atmmibobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmmibobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmmibobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmmibobjects.EntityData.Children = make(map[string]types.YChild)
    atmmibobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    atmmibobjects.EntityData.Leafs["atmVpCrossConnectIndexNext"] = types.YLeaf{"Atmvpcrossconnectindexnext", atmmibobjects.Atmvpcrossconnectindexnext}
    atmmibobjects.EntityData.Leafs["atmVcCrossConnectIndexNext"] = types.YLeaf{"Atmvccrossconnectindexnext", atmmibobjects.Atmvccrossconnectindexnext}
    atmmibobjects.EntityData.Leafs["atmTrafficDescrParamIndexNext"] = types.YLeaf{"Atmtrafficdescrparamindexnext", atmmibobjects.Atmtrafficdescrparamindexnext}
    return &(atmmibobjects.EntityData)
}

// ATMMIB_Atminterfaceconftable
// This table contains ATM local interface
// configuration parameters, one entry per ATM
// interface port.
type ATMMIB_Atminterfaceconftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains ATM interface configuration parameters and state
    // variables and is indexed by ifIndex values of ATM interfaces. The type is
    // slice of ATMMIB_Atminterfaceconftable_Atminterfaceconfentry.
    Atminterfaceconfentry []ATMMIB_Atminterfaceconftable_Atminterfaceconfentry
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetEntityData() *types.CommonEntityData {
    atminterfaceconftable.EntityData.YFilter = atminterfaceconftable.YFilter
    atminterfaceconftable.EntityData.YangName = "atmInterfaceConfTable"
    atminterfaceconftable.EntityData.BundleName = "cisco_ios_xe"
    atminterfaceconftable.EntityData.ParentYangName = "ATM-MIB"
    atminterfaceconftable.EntityData.SegmentPath = "atmInterfaceConfTable"
    atminterfaceconftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfaceconftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfaceconftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfaceconftable.EntityData.Children = make(map[string]types.YChild)
    atminterfaceconftable.EntityData.Children["atmInterfaceConfEntry"] = types.YChild{"Atminterfaceconfentry", nil}
    for i := range atminterfaceconftable.Atminterfaceconfentry {
        atminterfaceconftable.EntityData.Children[types.GetSegmentPath(&atminterfaceconftable.Atminterfaceconfentry[i])] = types.YChild{"Atminterfaceconfentry", &atminterfaceconftable.Atminterfaceconfentry[i]}
    }
    atminterfaceconftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atminterfaceconftable.EntityData)
}

// ATMMIB_Atminterfaceconftable_Atminterfaceconfentry
// This list contains ATM interface configuration
// parameters and state variables and is indexed
// by ifIndex values of ATM interfaces.
type ATMMIB_Atminterfaceconftable_Atminterfaceconfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The maximum number of VPCs (PVPCs and SVPCs) supported at this ATM
    // interface. At the ATM UNI, the maximum number of VPCs (PVPCs and SVPCs)
    // ranges from 0 to 256 only. The type is interface{} with range: 0..4096.
    Atminterfacemaxvpcs interface{}

    // The maximum number of VCCs (PVCCs and SVCCs) supported at this ATM
    // interface. The type is interface{} with range: 0..65536.
    Atminterfacemaxvccs interface{}

    // The number of VPCs (PVPC, Soft PVPC and SVPC) currently in use at this ATM
    // interface.  It includes the number of PVPCs and Soft PVPCs that are
    // configured at the interface, plus the number of SVPCs that are currently 
    // established at the interface.  At the ATM UNI, the configured number of
    // VPCs (PVPCs and SVPCs) can range from 0 to 256 only. The type is
    // interface{} with range: 0..4096.
    Atminterfaceconfvpcs interface{}

    // The number of VCCs (PVCC, Soft PVCC and SVCC) currently in use at this ATM
    // interface.  It includes the number of PVCCs and Soft PVCCs that are
    // configured at the interface, plus the number of SVCCs that are currently 
    // established at the interface. The type is interface{} with range: 0..65536.
    Atminterfaceconfvccs interface{}

    // The  maximum number of active VPI bits configured for use at the ATM
    // interface. At the ATM UNI, the maximum number of active VPI bits configured
    // for use ranges from 0 to 8 only. The type is interface{} with range: 0..12.
    Atminterfacemaxactivevpibits interface{}

    // The maximum number of active VCI bits configured for use at this ATM
    // interface. The type is interface{} with range: 0..16.
    Atminterfacemaxactivevcibits interface{}

    // The VPI value of the VCC supporting the ILMI at this ATM interface.  If the
    // values of atmInterfaceIlmiVpi and atmInterfaceIlmiVci are both equal to
    // zero then the ILMI is not supported at this ATM interface. The type is
    // interface{} with range: 0..4095.
    Atminterfaceilmivpi interface{}

    // The VCI value of the VCC supporting the ILMI at this ATM interface.  If the
    // values of atmInterfaceIlmiVpi and atmInterfaceIlmiVci are both equal to
    // zero then the ILMI is not supported at this ATM interface. The type is
    // interface{} with range: 0..65535.
    Atminterfaceilmivci interface{}

    // The type of primary ATM address configured for use at this ATM interface.
    // The type is Atminterfaceaddresstype.
    Atminterfaceaddresstype interface{}

    // The primary address assigned for administrative purposes, for example, an
    // address associated with the service provider side of a public network UNI
    // (thus, the value of this address corresponds with the value of
    // ifPhysAddress at the host side). If this interface has no assigned
    // administrative address, or when the address used for administrative
    // purposes is the same as that used for ifPhysAddress, then this is an octet
    // string of zero length. The type is string.
    Atminterfaceadminaddress interface{}

    // The IP address of the neighbor system connected to the  far end of this
    // interface, to which a Network Management Station can send SNMP messages, as
    // IP datagrams sent to UDP port 161, in order to access network management
    // information concerning the operation of that system.  Note that the value
    // of this object may be obtained in different ways, e.g., by manual
    // configuration, or through ILMI interaction with the neighbor system. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Atminterfacemyneighboripaddress interface{}

    // The textual name of the interface on the neighbor system on the far end of
    // this interface, and to which this interface connects.  If the neighbor
    // system is manageable through SNMP and supports the object ifName, the value
    // of this object must be identical with that of ifName for the ifEntry of the
    // lowest level physical interface for this port.  If this interface does not
    // have a textual name, the value of this object is a zero length string. 
    // Note that the value of this object may be obtained in different ways, e.g.,
    // by manual configuration, or through ILMI interaction with the neighbor
    // system. The type is string.
    Atminterfacemyneighborifname interface{}

    // The maximum number of VPI Bits that may currently be used at this ATM
    // interface. The value is the minimum of atmInterfaceMaxActiveVpiBits, and
    // the atmInterfaceMaxActiveVpiBits of the interface's UNI/NNI peer.  If the
    // interface does not negotiate with its peer to determine the number of VPI
    // Bits that can be used on the interface, then the value of this object must
    // equal atmInterfaceMaxActiveVpiBits. The type is interface{} with range:
    // 0..12.
    Atminterfacecurrentmaxvpibits interface{}

    // The maximum number of VCI Bits that may currently be used at this ATM
    // interface. The value is the minimum of atmInterfaceMaxActiveVciBits, and
    // the atmInterfaceMaxActiveVciBits of the interface's UNI/NNI peer.  If the
    // interface does not negotiate with its peer to determine the number of VCI
    // Bits that can be used on the interface, then the value of this object must
    // equal atmInterfaceMaxActiveVciBits. The type is interface{} with range:
    // 0..16.
    Atminterfacecurrentmaxvcibits interface{}

    // The identifier assigned by a service provider to the network side of a
    // public network UNI. If this interface has no assigned service provider
    // address, or for other interfaces this is an octet string of zero length.
    // The type is string.
    Atminterfacesubscraddress interface{}

    // The number of times the operational status of a PVCL on this interface has
    // gone down. The type is interface{} with range: 0..4294967295.
    Atmintfpvcfailures interface{}

    // The current number of VCLs on this interface for which there is an active
    // row in the atmVclTable having an atmVclConnKind value of `pvc' and an
    // atmVclOperStatus with a value other than `up'. The type is interface{} with
    // range: 0..4294967295.
    Atmintfcurrentlyfailingpvcls interface{}

    // Allows the generation of traps in response to PVCL failures on this
    // interface. The type is bool.
    Atmintfpvcfailurestrapenable interface{}

    // The minimum interval between the sending of cIntfPvcFailuresTrap
    // notifications for this interface. The type is interface{} with range:
    // 1..3600. Units are seconds.
    Atmintfpvcnotificationinterval interface{}

    // The interval for storing the failed time in
    // atmPreviouslyFailedPVclTimeStamp. The type is interface{} with range:
    // 0..3600. Units are seconds.
    Atmpreviouslyfailedpvclinterval interface{}

    // The current number PVCLs on this interface which  changed state to 'up'
    // since the last  atmIntPvcUp2Trap was sent. The type is interface{} with
    // range: 0..4294967295.
    Catmintfcurrentlydowntouppvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the OAM
    // loopback failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Catmintfoamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which the OAM loop back
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Catmintfcurrentoamfailingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // Segment CC OAM failed condition  but the status of each PVCL remain in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    Catmintfsegccoamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which the Segment CC OAM
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Catmintfcursegccoamfailingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // End-to-End CC OAM failed condition  but the status of each PVCL remain in
    // the 'up' state. The type is interface{} with range: 0..4294967295.
    Catmintfendccoamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which the End-to-End CC
    // OAM has failed but the status of each PVCL  remain in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Catmintfcurendccoamfailingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the AIS
    // RDI OAM failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Catmintfaisrdioamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which the AIS RDI OAM has
    // failed but the status of each PVCL remain  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcuraisrdioamfailingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in any
    // type of OAM failed condition but  the status of each PVCL remain in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    Catmintfanyoamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which  any of OAM has
    // failed but the status of each PVCL remain  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcuranyoamfailingpvcls interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    Catmintftypeofoamfailure interface{}

    // The total number of PVCLs in this interface which  are currently in the OAM
    // loopback recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Catmintfoamrcovedpvcls interface{}

    // The current number of PVCLs on this interface for which the OAM loop back
    // has recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcurrentoamrcovingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // Segment CC OAM recovered condition  and the status of each PVCL is in the
    // 'up' state. The type is interface{} with range: 0..4294967295.
    Catmintfsegccoamrcovedpvcls interface{}

    // The current number of PVCLs on this interface for which the Segment CC OAM
    // has recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcursegccoamrcovingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the
    // End-to-End CC OAM recovered condition  and the status of each PVCL is in
    // the 'up' state. The type is interface{} with range: 0..4294967295.
    Catmintfendccoamrcovedpvcls interface{}

    // The current number of PVCLs on this interface for which the End-to-End CC
    // OAM has recovered and the status of each PVCL  is in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Catmintfcurendccoamrcovingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the AIS
    // RDI OAM recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Catmintfaisrdioamrcovedpvcls interface{}

    // The current number of PVCLs on this interface for which the AIS RDI OAM has
    // recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcuraisrdioamrcovingpvcls interface{}

    // The total number of PVCLs in this interface which  are currently in any
    // type of OAM recovered condition and  the status of each PVCL is in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Catmintfanyoamrcovedpvcls interface{}

    // The current number of PVCLs on this interface for which  any of OAM has
    // recovered and the status of each PVCL is  in the 'up' state in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Catmintfcuranyoamrcovingpvcls interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    Catmintftypeofoamrecover interface{}

    // The current number PVCLs on this interface which  changed state to 'up'
    // since the last  atmIntPvcUpTrap was sent. The type is interface{} with
    // range: 0..4294967295.
    Atmintfcurrentlydowntouppvcls interface{}

    // The total number of PVCLs in this interface which  are currently in the oam
    // loopback failed condition but  the status of each PVCL remain in the 'up'
    // state. The type is interface{} with range: 0..4294967295.
    Atmintfoamfailedpvcls interface{}

    // The current number of PVCLs on this interface for which the oam loop back
    // has failed but the status of each PVCL remain  in the 'up' state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmintfcurrentlyoamfailingpvcls interface{}
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetEntityData() *types.CommonEntityData {
    atminterfaceconfentry.EntityData.YFilter = atminterfaceconfentry.YFilter
    atminterfaceconfentry.EntityData.YangName = "atmInterfaceConfEntry"
    atminterfaceconfentry.EntityData.BundleName = "cisco_ios_xe"
    atminterfaceconfentry.EntityData.ParentYangName = "atmInterfaceConfTable"
    atminterfaceconfentry.EntityData.SegmentPath = "atmInterfaceConfEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfaceconfentry.Ifindex) + "']"
    atminterfaceconfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfaceconfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfaceconfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfaceconfentry.EntityData.Children = make(map[string]types.YChild)
    atminterfaceconfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atminterfaceconfentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atminterfaceconfentry.Ifindex}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMaxVpcs"] = types.YLeaf{"Atminterfacemaxvpcs", atminterfaceconfentry.Atminterfacemaxvpcs}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMaxVccs"] = types.YLeaf{"Atminterfacemaxvccs", atminterfaceconfentry.Atminterfacemaxvccs}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceConfVpcs"] = types.YLeaf{"Atminterfaceconfvpcs", atminterfaceconfentry.Atminterfaceconfvpcs}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceConfVccs"] = types.YLeaf{"Atminterfaceconfvccs", atminterfaceconfentry.Atminterfaceconfvccs}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMaxActiveVpiBits"] = types.YLeaf{"Atminterfacemaxactivevpibits", atminterfaceconfentry.Atminterfacemaxactivevpibits}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMaxActiveVciBits"] = types.YLeaf{"Atminterfacemaxactivevcibits", atminterfaceconfentry.Atminterfacemaxactivevcibits}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceIlmiVpi"] = types.YLeaf{"Atminterfaceilmivpi", atminterfaceconfentry.Atminterfaceilmivpi}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceIlmiVci"] = types.YLeaf{"Atminterfaceilmivci", atminterfaceconfentry.Atminterfaceilmivci}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceAddressType"] = types.YLeaf{"Atminterfaceaddresstype", atminterfaceconfentry.Atminterfaceaddresstype}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceAdminAddress"] = types.YLeaf{"Atminterfaceadminaddress", atminterfaceconfentry.Atminterfaceadminaddress}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMyNeighborIpAddress"] = types.YLeaf{"Atminterfacemyneighboripaddress", atminterfaceconfentry.Atminterfacemyneighboripaddress}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceMyNeighborIfName"] = types.YLeaf{"Atminterfacemyneighborifname", atminterfaceconfentry.Atminterfacemyneighborifname}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceCurrentMaxVpiBits"] = types.YLeaf{"Atminterfacecurrentmaxvpibits", atminterfaceconfentry.Atminterfacecurrentmaxvpibits}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceCurrentMaxVciBits"] = types.YLeaf{"Atminterfacecurrentmaxvcibits", atminterfaceconfentry.Atminterfacecurrentmaxvcibits}
    atminterfaceconfentry.EntityData.Leafs["atmInterfaceSubscrAddress"] = types.YLeaf{"Atminterfacesubscraddress", atminterfaceconfentry.Atminterfacesubscraddress}
    atminterfaceconfentry.EntityData.Leafs["atmIntfPvcFailures"] = types.YLeaf{"Atmintfpvcfailures", atminterfaceconfentry.Atmintfpvcfailures}
    atminterfaceconfentry.EntityData.Leafs["atmIntfCurrentlyFailingPVcls"] = types.YLeaf{"Atmintfcurrentlyfailingpvcls", atminterfaceconfentry.Atmintfcurrentlyfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["atmIntfPvcFailuresTrapEnable"] = types.YLeaf{"Atmintfpvcfailurestrapenable", atminterfaceconfentry.Atmintfpvcfailurestrapenable}
    atminterfaceconfentry.EntityData.Leafs["atmIntfPvcNotificationInterval"] = types.YLeaf{"Atmintfpvcnotificationinterval", atminterfaceconfentry.Atmintfpvcnotificationinterval}
    atminterfaceconfentry.EntityData.Leafs["atmPreviouslyFailedPVclInterval"] = types.YLeaf{"Atmpreviouslyfailedpvclinterval", atminterfaceconfentry.Atmpreviouslyfailedpvclinterval}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurrentlyDownToUpPVcls"] = types.YLeaf{"Catmintfcurrentlydowntouppvcls", atminterfaceconfentry.Catmintfcurrentlydowntouppvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfOAMFailedPVcls"] = types.YLeaf{"Catmintfoamfailedpvcls", atminterfaceconfentry.Catmintfoamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurrentOAMFailingPVcls"] = types.YLeaf{"Catmintfcurrentoamfailingpvcls", atminterfaceconfentry.Catmintfcurrentoamfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfSegCCOAMFailedPVcls"] = types.YLeaf{"Catmintfsegccoamfailedpvcls", atminterfaceconfentry.Catmintfsegccoamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurSegCCOAMFailingPVcls"] = types.YLeaf{"Catmintfcursegccoamfailingpvcls", atminterfaceconfentry.Catmintfcursegccoamfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfEndCCOAMFailedPVcls"] = types.YLeaf{"Catmintfendccoamfailedpvcls", atminterfaceconfentry.Catmintfendccoamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurEndCCOAMFailingPVcls"] = types.YLeaf{"Catmintfcurendccoamfailingpvcls", atminterfaceconfentry.Catmintfcurendccoamfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfAISRDIOAMFailedPVcls"] = types.YLeaf{"Catmintfaisrdioamfailedpvcls", atminterfaceconfentry.Catmintfaisrdioamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurAISRDIOAMFailingPVcls"] = types.YLeaf{"Catmintfcuraisrdioamfailingpvcls", atminterfaceconfentry.Catmintfcuraisrdioamfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfAnyOAMFailedPVcls"] = types.YLeaf{"Catmintfanyoamfailedpvcls", atminterfaceconfentry.Catmintfanyoamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurAnyOAMFailingPVcls"] = types.YLeaf{"Catmintfcuranyoamfailingpvcls", atminterfaceconfentry.Catmintfcuranyoamfailingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfTypeOfOAMFailure"] = types.YLeaf{"Catmintftypeofoamfailure", atminterfaceconfentry.Catmintftypeofoamfailure}
    atminterfaceconfentry.EntityData.Leafs["catmIntfOAMRcovedPVcls"] = types.YLeaf{"Catmintfoamrcovedpvcls", atminterfaceconfentry.Catmintfoamrcovedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurrentOAMRcovingPVcls"] = types.YLeaf{"Catmintfcurrentoamrcovingpvcls", atminterfaceconfentry.Catmintfcurrentoamrcovingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfSegCCOAMRcovedPVcls"] = types.YLeaf{"Catmintfsegccoamrcovedpvcls", atminterfaceconfentry.Catmintfsegccoamrcovedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurSegCCOAMRcovingPVcls"] = types.YLeaf{"Catmintfcursegccoamrcovingpvcls", atminterfaceconfentry.Catmintfcursegccoamrcovingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfEndCCOAMRcovedPVcls"] = types.YLeaf{"Catmintfendccoamrcovedpvcls", atminterfaceconfentry.Catmintfendccoamrcovedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurEndCCOAMRcovingPVcls"] = types.YLeaf{"Catmintfcurendccoamrcovingpvcls", atminterfaceconfentry.Catmintfcurendccoamrcovingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfAISRDIOAMRcovedPVcls"] = types.YLeaf{"Catmintfaisrdioamrcovedpvcls", atminterfaceconfentry.Catmintfaisrdioamrcovedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurAISRDIOAMRcovingPVcls"] = types.YLeaf{"Catmintfcuraisrdioamrcovingpvcls", atminterfaceconfentry.Catmintfcuraisrdioamrcovingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfAnyOAMRcovedPVcls"] = types.YLeaf{"Catmintfanyoamrcovedpvcls", atminterfaceconfentry.Catmintfanyoamrcovedpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfCurAnyOAMRcovingPVcls"] = types.YLeaf{"Catmintfcuranyoamrcovingpvcls", atminterfaceconfentry.Catmintfcuranyoamrcovingpvcls}
    atminterfaceconfentry.EntityData.Leafs["catmIntfTypeOfOAMRecover"] = types.YLeaf{"Catmintftypeofoamrecover", atminterfaceconfentry.Catmintftypeofoamrecover}
    atminterfaceconfentry.EntityData.Leafs["atmIntfCurrentlyDownToUpPVcls"] = types.YLeaf{"Atmintfcurrentlydowntouppvcls", atminterfaceconfentry.Atmintfcurrentlydowntouppvcls}
    atminterfaceconfentry.EntityData.Leafs["atmIntfOAMFailedPVcls"] = types.YLeaf{"Atmintfoamfailedpvcls", atminterfaceconfentry.Atmintfoamfailedpvcls}
    atminterfaceconfentry.EntityData.Leafs["atmIntfCurrentlyOAMFailingPVcls"] = types.YLeaf{"Atmintfcurrentlyoamfailingpvcls", atminterfaceconfentry.Atmintfcurrentlyoamfailingpvcls}
    return &(atminterfaceconfentry.EntityData)
}

// ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype represents for use at this ATM interface.
type ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype string

const (
    ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype_private ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype = "private"

    ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype_nsapE164 ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype = "nsapE164"

    ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype_nativeE164 ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype = "nativeE164"

    ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype_other ATMMIB_Atminterfaceconftable_Atminterfaceconfentry_Atminterfaceaddresstype = "other"
)

// ATMMIB_Atminterfaceds3Plcptable
// This table contains ATM interface DS3 PLCP
// parameters and state variables, one entry per
// ATM interface port.
type ATMMIB_Atminterfaceds3Plcptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains DS3 PLCP parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry.
    Atminterfaceds3Plcpentry []ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetEntityData() *types.CommonEntityData {
    atminterfaceds3Plcptable.EntityData.YFilter = atminterfaceds3Plcptable.YFilter
    atminterfaceds3Plcptable.EntityData.YangName = "atmInterfaceDs3PlcpTable"
    atminterfaceds3Plcptable.EntityData.BundleName = "cisco_ios_xe"
    atminterfaceds3Plcptable.EntityData.ParentYangName = "ATM-MIB"
    atminterfaceds3Plcptable.EntityData.SegmentPath = "atmInterfaceDs3PlcpTable"
    atminterfaceds3Plcptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfaceds3Plcptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfaceds3Plcptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfaceds3Plcptable.EntityData.Children = make(map[string]types.YChild)
    atminterfaceds3Plcptable.EntityData.Children["atmInterfaceDs3PlcpEntry"] = types.YChild{"Atminterfaceds3Plcpentry", nil}
    for i := range atminterfaceds3Plcptable.Atminterfaceds3Plcpentry {
        atminterfaceds3Plcptable.EntityData.Children[types.GetSegmentPath(&atminterfaceds3Plcptable.Atminterfaceds3Plcpentry[i])] = types.YChild{"Atminterfaceds3Plcpentry", &atminterfaceds3Plcptable.Atminterfaceds3Plcpentry[i]}
    }
    atminterfaceds3Plcptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atminterfaceds3Plcptable.EntityData)
}

// ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry
// This list contains DS3 PLCP parameters and
// state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of DS3 PLCP Severely Errored Framing Seconds (SEFS). Each SEFS
    // represents a one-second interval which contains one or more SEF events. The
    // type is interface{} with range: 0..4294967295.
    Atminterfaceds3Plcpsefss interface{}

    // This variable indicates if there is an alarm present for the DS3 PLCP.  The
    // value receivedFarEndAlarm means that the DS3 PLCP has received an incoming
    // Yellow Signal, the value incomingLOF means that the DS3 PLCP has declared a
    // loss of frame (LOF) failure condition, and the value noAlarm means that
    // there are no alarms present. Transition from the failure to the no alarm
    // state occurs when no defects (e.g., LOF) are received for more than 10
    // seconds. The type is Atminterfaceds3Plcpalarmstate.
    Atminterfaceds3Plcpalarmstate interface{}

    // The counter associated with the number of Unavailable Seconds encountered
    // by the PLCP. The type is interface{} with range: 0..4294967295.
    Atminterfaceds3Plcpuass interface{}
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetEntityData() *types.CommonEntityData {
    atminterfaceds3Plcpentry.EntityData.YFilter = atminterfaceds3Plcpentry.YFilter
    atminterfaceds3Plcpentry.EntityData.YangName = "atmInterfaceDs3PlcpEntry"
    atminterfaceds3Plcpentry.EntityData.BundleName = "cisco_ios_xe"
    atminterfaceds3Plcpentry.EntityData.ParentYangName = "atmInterfaceDs3PlcpTable"
    atminterfaceds3Plcpentry.EntityData.SegmentPath = "atmInterfaceDs3PlcpEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfaceds3Plcpentry.Ifindex) + "']"
    atminterfaceds3Plcpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfaceds3Plcpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfaceds3Plcpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfaceds3Plcpentry.EntityData.Children = make(map[string]types.YChild)
    atminterfaceds3Plcpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atminterfaceds3Plcpentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atminterfaceds3Plcpentry.Ifindex}
    atminterfaceds3Plcpentry.EntityData.Leafs["atmInterfaceDs3PlcpSEFSs"] = types.YLeaf{"Atminterfaceds3Plcpsefss", atminterfaceds3Plcpentry.Atminterfaceds3Plcpsefss}
    atminterfaceds3Plcpentry.EntityData.Leafs["atmInterfaceDs3PlcpAlarmState"] = types.YLeaf{"Atminterfaceds3Plcpalarmstate", atminterfaceds3Plcpentry.Atminterfaceds3Plcpalarmstate}
    atminterfaceds3Plcpentry.EntityData.Leafs["atmInterfaceDs3PlcpUASs"] = types.YLeaf{"Atminterfaceds3Plcpuass", atminterfaceds3Plcpentry.Atminterfaceds3Plcpuass}
    return &(atminterfaceds3Plcpentry.EntityData)
}

// ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate represents for more than 10 seconds.
type ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate string

const (
    ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate_noAlarm ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate = "noAlarm"

    ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate_receivedFarEndAlarm ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate = "receivedFarEndAlarm"

    ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate_incomingLOF ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry_Atminterfaceds3Plcpalarmstate = "incomingLOF"
)

// ATMMIB_Atminterfacetctable
// This table contains ATM interface TC
// Sublayer parameters and state variables,
// one entry per ATM interface port.
type ATMMIB_Atminterfacetctable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains TC Sublayer parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_Atminterfacetctable_Atminterfacetcentry.
    Atminterfacetcentry []ATMMIB_Atminterfacetctable_Atminterfacetcentry
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetEntityData() *types.CommonEntityData {
    atminterfacetctable.EntityData.YFilter = atminterfacetctable.YFilter
    atminterfacetctable.EntityData.YangName = "atmInterfaceTCTable"
    atminterfacetctable.EntityData.BundleName = "cisco_ios_xe"
    atminterfacetctable.EntityData.ParentYangName = "ATM-MIB"
    atminterfacetctable.EntityData.SegmentPath = "atmInterfaceTCTable"
    atminterfacetctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfacetctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfacetctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfacetctable.EntityData.Children = make(map[string]types.YChild)
    atminterfacetctable.EntityData.Children["atmInterfaceTCEntry"] = types.YChild{"Atminterfacetcentry", nil}
    for i := range atminterfacetctable.Atminterfacetcentry {
        atminterfacetctable.EntityData.Children[types.GetSegmentPath(&atminterfacetctable.Atminterfacetcentry[i])] = types.YChild{"Atminterfacetcentry", &atminterfacetctable.Atminterfacetcentry[i]}
    }
    atminterfacetctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atminterfacetctable.EntityData)
}

// ATMMIB_Atminterfacetctable_Atminterfacetcentry
// This list contains TC Sublayer parameters
// and state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_Atminterfacetctable_Atminterfacetcentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of times the Out of Cell Delineation (OCD) events occur.  If
    // seven consecutive ATM cells have Header Error Control (HEC) violations, an
    // OCD event occurs. A high number of OCD events may indicate a problem with
    // the TC Sublayer. The type is interface{} with range: 0..4294967295.
    Atminterfaceocdevents interface{}

    // This variable indicates if there is an alarm present for the TC Sublayer. 
    // The value lcdFailure(2) indicates that the TC Sublayer is currently in the
    // Loss of Cell Delineation (LCD) defect maintenance state.  The value
    // noAlarm(1) indicates that the TC Sublayer is currently not in the LCD
    // defect maintenance state. The type is Atminterfacetcalarmstate.
    Atminterfacetcalarmstate interface{}
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetEntityData() *types.CommonEntityData {
    atminterfacetcentry.EntityData.YFilter = atminterfacetcentry.YFilter
    atminterfacetcentry.EntityData.YangName = "atmInterfaceTCEntry"
    atminterfacetcentry.EntityData.BundleName = "cisco_ios_xe"
    atminterfacetcentry.EntityData.ParentYangName = "atmInterfaceTCTable"
    atminterfacetcentry.EntityData.SegmentPath = "atmInterfaceTCEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfacetcentry.Ifindex) + "']"
    atminterfacetcentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atminterfacetcentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atminterfacetcentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atminterfacetcentry.EntityData.Children = make(map[string]types.YChild)
    atminterfacetcentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atminterfacetcentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atminterfacetcentry.Ifindex}
    atminterfacetcentry.EntityData.Leafs["atmInterfaceOCDEvents"] = types.YLeaf{"Atminterfaceocdevents", atminterfacetcentry.Atminterfaceocdevents}
    atminterfacetcentry.EntityData.Leafs["atmInterfaceTCAlarmState"] = types.YLeaf{"Atminterfacetcalarmstate", atminterfacetcentry.Atminterfacetcalarmstate}
    return &(atminterfacetcentry.EntityData)
}

// ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate represents maintenance state.
type ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate string

const (
    ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate_noAlarm ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate = "noAlarm"

    ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate_lcdFailure ATMMIB_Atminterfacetctable_Atminterfacetcentry_Atminterfacetcalarmstate = "lcdFailure"
)

// ATMMIB_Atmtrafficdescrparamtable
// This table contains information on ATM traffic
// descriptor type and the associated parameters.
type ATMMIB_Atmtrafficdescrparamtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains ATM traffic descriptor type and the associated
    // parameters. The type is slice of
    // ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry.
    Atmtrafficdescrparamentry []ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetEntityData() *types.CommonEntityData {
    atmtrafficdescrparamtable.EntityData.YFilter = atmtrafficdescrparamtable.YFilter
    atmtrafficdescrparamtable.EntityData.YangName = "atmTrafficDescrParamTable"
    atmtrafficdescrparamtable.EntityData.BundleName = "cisco_ios_xe"
    atmtrafficdescrparamtable.EntityData.ParentYangName = "ATM-MIB"
    atmtrafficdescrparamtable.EntityData.SegmentPath = "atmTrafficDescrParamTable"
    atmtrafficdescrparamtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmtrafficdescrparamtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmtrafficdescrparamtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmtrafficdescrparamtable.EntityData.Children = make(map[string]types.YChild)
    atmtrafficdescrparamtable.EntityData.Children["atmTrafficDescrParamEntry"] = types.YChild{"Atmtrafficdescrparamentry", nil}
    for i := range atmtrafficdescrparamtable.Atmtrafficdescrparamentry {
        atmtrafficdescrparamtable.EntityData.Children[types.GetSegmentPath(&atmtrafficdescrparamtable.Atmtrafficdescrparamentry[i])] = types.YChild{"Atmtrafficdescrparamentry", &atmtrafficdescrparamtable.Atmtrafficdescrparamentry[i]}
    }
    atmtrafficdescrparamtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmtrafficdescrparamtable.EntityData)
}

// ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry
// This list contains ATM traffic descriptor
// type and the associated parameters.
type ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object is used by the virtual link table
    // (i.e., VPL or VCL table) to identify the row of this table. When creating a
    // new row in the table the value of this index may be obtained by retrieving
    // the value of atmTrafficDescrParamIndexNext. The type is interface{} with
    // range: 1..2147483647.
    Atmtrafficdescrparamindex interface{}

    // The value of this object identifies the type of ATM traffic descriptor. The
    // type may indicate no traffic descriptor or traffic descriptor with one or
    // more parameters. These parameters are specified as a parameter vector, in
    // the corresponding instances of the objects:     atmTrafficDescrParam1    
    // atmTrafficDescrParam2     atmTrafficDescrParam3     atmTrafficDescrParam4  
    // atmTrafficDescrParam5. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Atmtrafficdescrtype interface{}

    // The first parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    Atmtrafficdescrparam1 interface{}

    // The second parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    Atmtrafficdescrparam2 interface{}

    // The third parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    Atmtrafficdescrparam3 interface{}

    // The fourth parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    Atmtrafficdescrparam4 interface{}

    // The fifth parameter of the ATM traffic descriptor used according to the
    // value of atmTrafficDescrType. The type is interface{} with range:
    // -2147483648..2147483647.
    Atmtrafficdescrparam5 interface{}

    // The value of this object identifies the QoS Class. Four Service classes
    // have been specified in the ATM Forum UNI Specification: Service Class A:
    // Constant bit rate video and                  Circuit emulation Service
    // Class B: Variable bit rate video/audio Service Class C: Connection-oriented
    // data Service Class D: Connectionless data Four QoS classes numbered 1, 2,
    // 3, and 4 have been specified with the aim to support service classes A, B,
    // C, and D respectively. An unspecified QoS Class numbered `0' is used for
    // best effort traffic. The type is interface{} with range: 0..255.
    Atmtrafficqosclass interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The type is RowStatus.
    Atmtrafficdescrrowstatus interface{}

    // The ATM service category. The type is AtmServiceCategory.
    Atmservicecategory interface{}

    // If set to 'true', this object indicates that the network is requested to
    // treat data for this connection, in the given direction, as frames (e.g.
    // AAL5 CPCS_PDU's) rather than as individual cells.  While the precise
    // implementation is network-specific, this treatment may for example involve
    // discarding entire frames during congestion, rather than a few cells from
    // many frames. The type is bool.
    Atmtrafficframediscard interface{}
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetEntityData() *types.CommonEntityData {
    atmtrafficdescrparamentry.EntityData.YFilter = atmtrafficdescrparamentry.YFilter
    atmtrafficdescrparamentry.EntityData.YangName = "atmTrafficDescrParamEntry"
    atmtrafficdescrparamentry.EntityData.BundleName = "cisco_ios_xe"
    atmtrafficdescrparamentry.EntityData.ParentYangName = "atmTrafficDescrParamTable"
    atmtrafficdescrparamentry.EntityData.SegmentPath = "atmTrafficDescrParamEntry" + "[atmTrafficDescrParamIndex='" + fmt.Sprintf("%v", atmtrafficdescrparamentry.Atmtrafficdescrparamindex) + "']"
    atmtrafficdescrparamentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmtrafficdescrparamentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmtrafficdescrparamentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmtrafficdescrparamentry.EntityData.Children = make(map[string]types.YChild)
    atmtrafficdescrparamentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParamIndex"] = types.YLeaf{"Atmtrafficdescrparamindex", atmtrafficdescrparamentry.Atmtrafficdescrparamindex}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrType"] = types.YLeaf{"Atmtrafficdescrtype", atmtrafficdescrparamentry.Atmtrafficdescrtype}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParam1"] = types.YLeaf{"Atmtrafficdescrparam1", atmtrafficdescrparamentry.Atmtrafficdescrparam1}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParam2"] = types.YLeaf{"Atmtrafficdescrparam2", atmtrafficdescrparamentry.Atmtrafficdescrparam2}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParam3"] = types.YLeaf{"Atmtrafficdescrparam3", atmtrafficdescrparamentry.Atmtrafficdescrparam3}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParam4"] = types.YLeaf{"Atmtrafficdescrparam4", atmtrafficdescrparamentry.Atmtrafficdescrparam4}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrParam5"] = types.YLeaf{"Atmtrafficdescrparam5", atmtrafficdescrparamentry.Atmtrafficdescrparam5}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficQoSClass"] = types.YLeaf{"Atmtrafficqosclass", atmtrafficdescrparamentry.Atmtrafficqosclass}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficDescrRowStatus"] = types.YLeaf{"Atmtrafficdescrrowstatus", atmtrafficdescrparamentry.Atmtrafficdescrrowstatus}
    atmtrafficdescrparamentry.EntityData.Leafs["atmServiceCategory"] = types.YLeaf{"Atmservicecategory", atmtrafficdescrparamentry.Atmservicecategory}
    atmtrafficdescrparamentry.EntityData.Leafs["atmTrafficFrameDiscard"] = types.YLeaf{"Atmtrafficframediscard", atmtrafficdescrparamentry.Atmtrafficframediscard}
    return &(atmtrafficdescrparamentry.EntityData)
}

// ATMMIB_Atmvpltable
// The Virtual Path Link (VPL) table.  A
// bi-directional VPL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
// Entries are not present in this table for
// the VPIs used by entries in the atmVclTable.
type ATMMIB_Atmvpltable struct {
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
    // ATMMIB_Atmvpltable_Atmvplentry.
    Atmvplentry []ATMMIB_Atmvpltable_Atmvplentry
}

func (atmvpltable *ATMMIB_Atmvpltable) GetEntityData() *types.CommonEntityData {
    atmvpltable.EntityData.YFilter = atmvpltable.YFilter
    atmvpltable.EntityData.YangName = "atmVplTable"
    atmvpltable.EntityData.BundleName = "cisco_ios_xe"
    atmvpltable.EntityData.ParentYangName = "ATM-MIB"
    atmvpltable.EntityData.SegmentPath = "atmVplTable"
    atmvpltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvpltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvpltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvpltable.EntityData.Children = make(map[string]types.YChild)
    atmvpltable.EntityData.Children["atmVplEntry"] = types.YChild{"Atmvplentry", nil}
    for i := range atmvpltable.Atmvplentry {
        atmvpltable.EntityData.Children[types.GetSegmentPath(&atmvpltable.Atmvplentry[i])] = types.YChild{"Atmvplentry", &atmvpltable.Atmvplentry[i]}
    }
    atmvpltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmvpltable.EntityData)
}

// ATMMIB_Atmvpltable_Atmvplentry
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
type ATMMIB_Atmvpltable_Atmvplentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The VPI value of the VPL. The type is interface{}
    // with range: 0..4095.
    Atmvplvpi interface{}

    // This object is instanciated only for a VPL which terminates a VPC (i.e.,
    // one which is NOT cross-connected to other VPLs). Its value specifies the
    // desired administrative state of the VPL. The type is AtmVorXAdminStatus.
    Atmvpladminstatus interface{}

    // The current operational status of the VPL. The type is AtmVorXOperStatus.
    Atmvploperstatus interface{}

    // The value of sysUpTime at the time this VPL entered its current operational
    // state. The type is interface{} with range: 0..4294967295.
    Atmvpllastchange interface{}

    // The value of this object identifies the row in the
    // atmTrafficDescrParamTable which applies to the receive direction of the
    // VPL. The type is interface{} with range: 0..2147483647.
    Atmvplreceivetrafficdescrindex interface{}

    // The value of this object identifies the row in the
    // atmTrafficDescrParamTable which applies to the transmit direction of the
    // VPL. The type is interface{} with range: 0..2147483647.
    Atmvpltransmittrafficdescrindex interface{}

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
    Atmvplcrossconnectidentifier interface{}

    // This object is used to create, delete or modify a row in this table. To
    // create a new VCL, this object is initially set to 'createAndWait' or
    // 'createAndGo'.  This object should not be set to 'active' unless the
    // following columnar objects have been set to their desired value in this
    // row: atmVplReceiveTrafficDescrIndex and atmVplTransmitTrafficDescrIndex.
    // The DESCRIPTION of atmVplEntry provides further guidance to row treatment
    // in this table. The type is RowStatus.
    Atmvplrowstatus interface{}

    // The connection topology type. The type is AtmConnCastType.
    Atmvplcasttype interface{}

    // The use of call control. The type is AtmConnKind.
    Atmvplconnkind interface{}
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetEntityData() *types.CommonEntityData {
    atmvplentry.EntityData.YFilter = atmvplentry.YFilter
    atmvplentry.EntityData.YangName = "atmVplEntry"
    atmvplentry.EntityData.BundleName = "cisco_ios_xe"
    atmvplentry.EntityData.ParentYangName = "atmVplTable"
    atmvplentry.EntityData.SegmentPath = "atmVplEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmvplentry.Ifindex) + "']" + "[atmVplVpi='" + fmt.Sprintf("%v", atmvplentry.Atmvplvpi) + "']"
    atmvplentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvplentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvplentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvplentry.EntityData.Children = make(map[string]types.YChild)
    atmvplentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmvplentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atmvplentry.Ifindex}
    atmvplentry.EntityData.Leafs["atmVplVpi"] = types.YLeaf{"Atmvplvpi", atmvplentry.Atmvplvpi}
    atmvplentry.EntityData.Leafs["atmVplAdminStatus"] = types.YLeaf{"Atmvpladminstatus", atmvplentry.Atmvpladminstatus}
    atmvplentry.EntityData.Leafs["atmVplOperStatus"] = types.YLeaf{"Atmvploperstatus", atmvplentry.Atmvploperstatus}
    atmvplentry.EntityData.Leafs["atmVplLastChange"] = types.YLeaf{"Atmvpllastchange", atmvplentry.Atmvpllastchange}
    atmvplentry.EntityData.Leafs["atmVplReceiveTrafficDescrIndex"] = types.YLeaf{"Atmvplreceivetrafficdescrindex", atmvplentry.Atmvplreceivetrafficdescrindex}
    atmvplentry.EntityData.Leafs["atmVplTransmitTrafficDescrIndex"] = types.YLeaf{"Atmvpltransmittrafficdescrindex", atmvplentry.Atmvpltransmittrafficdescrindex}
    atmvplentry.EntityData.Leafs["atmVplCrossConnectIdentifier"] = types.YLeaf{"Atmvplcrossconnectidentifier", atmvplentry.Atmvplcrossconnectidentifier}
    atmvplentry.EntityData.Leafs["atmVplRowStatus"] = types.YLeaf{"Atmvplrowstatus", atmvplentry.Atmvplrowstatus}
    atmvplentry.EntityData.Leafs["atmVplCastType"] = types.YLeaf{"Atmvplcasttype", atmvplentry.Atmvplcasttype}
    atmvplentry.EntityData.Leafs["atmVplConnKind"] = types.YLeaf{"Atmvplconnkind", atmvplentry.Atmvplconnkind}
    return &(atmvplentry.EntityData)
}

// ATMMIB_Atmvcltable
// The Virtual Channel Link (VCL) table.  A
// bi-directional VCL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
type ATMMIB_Atmvcltable struct {
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
    // ATMMIB_Atmvcltable_Atmvclentry.
    Atmvclentry []ATMMIB_Atmvcltable_Atmvclentry
}

func (atmvcltable *ATMMIB_Atmvcltable) GetEntityData() *types.CommonEntityData {
    atmvcltable.EntityData.YFilter = atmvcltable.YFilter
    atmvcltable.EntityData.YangName = "atmVclTable"
    atmvcltable.EntityData.BundleName = "cisco_ios_xe"
    atmvcltable.EntityData.ParentYangName = "ATM-MIB"
    atmvcltable.EntityData.SegmentPath = "atmVclTable"
    atmvcltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvcltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvcltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvcltable.EntityData.Children = make(map[string]types.YChild)
    atmvcltable.EntityData.Children["atmVclEntry"] = types.YChild{"Atmvclentry", nil}
    for i := range atmvcltable.Atmvclentry {
        atmvcltable.EntityData.Children[types.GetSegmentPath(&atmvcltable.Atmvclentry[i])] = types.YChild{"Atmvclentry", &atmvcltable.Atmvclentry[i]}
    }
    atmvcltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmvcltable.EntityData)
}

// ATMMIB_Atmvcltable_Atmvclentry
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
type ATMMIB_Atmvcltable_Atmvclentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The VPI value of the VCL. The type is interface{}
    // with range: 0..4095.
    Atmvclvpi interface{}

    // This attribute is a key. The VCI value of the VCL. The type is interface{}
    // with range: 0..65535.
    Atmvclvci interface{}

    // This object is instanciated only for a VCL which terminates a VCC (i.e.,
    // one which is NOT cross-connected to other VCLs). Its value specifies the
    // desired administrative state of the VCL. The type is AtmVorXAdminStatus.
    Atmvcladminstatus interface{}

    // The current operational status of the VCL. The type is AtmVorXOperStatus.
    Atmvcloperstatus interface{}

    // The value of sysUpTime at the time this VCL entered its current operational
    // state. The type is interface{} with range: 0..4294967295.
    Atmvcllastchange interface{}

    // The value of this object identifies the row in the ATM Traffic Descriptor
    // Table which applies to the receive direction of this VCL. The type is
    // interface{} with range: 0..2147483647.
    Atmvclreceivetrafficdescrindex interface{}

    // The value of this object identifies the row of the ATM Traffic Descriptor
    // Table which applies to the transmit direction of this VCL. The type is
    // interface{} with range: 0..2147483647.
    Atmvcltransmittrafficdescrindex interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL is in use. The type of AAL used on this VCC. The
    // AAL type includes AAL1, AAL2, AAL3/4, and AAL5. The other(4) may be
    // user-defined AAL type.  The unknown type indicates that the AAL type cannot
    // be determined. The type is Atmvccaaltype.
    Atmvccaaltype interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The maximum AAL5 CPCS SDU size in
    // octets that is supported on the transmit direction of this VCC. The type is
    // interface{} with range: 1..65535.
    Atmvccaal5Cpcstransmitsdusize interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The maximum AAL5 CPCS SDU size in
    // octets that is supported on the receive direction of this VCC. The type is
    // interface{} with range: 1..65535.
    Atmvccaal5Cpcsreceivesdusize interface{}

    // An instance of this object only exists when the local VCL end-point is also
    // the VCC end-point, and AAL5 is in use. The type of data encapsulation used
    // over the AAL5 SSCS layer. The definitions reference RFC 1483 Multiprotocol
    // Encapsulation over ATM AAL5 and to the ATM Forum LAN Emulation
    // specification. The type is Atmvccaal5Encapstype.
    Atmvccaal5Encapstype interface{}

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
    Atmvclcrossconnectidentifier interface{}

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
    Atmvclrowstatus interface{}

    // The connection topology type. The type is AtmConnCastType.
    Atmvclcasttype interface{}

    // The use of call control. The type is AtmConnKind.
    Atmvclconnkind interface{}

    // Specifies OAM loopback frequency. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    Catmxvcloamloopbackfreq interface{}

    // Specifies OAM retry polling frequency. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    Catmxvcloamretryfreq interface{}

    // Specifies OAM retry count before declaring a VC  is up. The type is
    // interface{} with range: 0..4294967295.
    Catmxvcloamupretrycount interface{}

    // Specifies OAM retry count before declaring a VC  is down. The type is
    // interface{} with range: 0..4294967295.
    Catmxvcloamdownretrycount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Activation retry count. The
    // type is interface{} with range: 0..4294967295.
    Catmxvcloamendccactcount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Deactivation retry count.
    // The type is interface{} with range: 0..4294967295.
    Catmxvcloamendccdeactcount interface{}

    // Specifies OAM End-to-end Continuity check (CC)  Activation/Deactivation
    // retry frequency. The type is interface{} with range: 0..4294967295. Units
    // are seconds.
    Catmxvcloamendccretryfreq interface{}

    // Specifies OAM Segment Continuity check (CC)  Activation retry count. The
    // type is interface{} with range: 0..4294967295.
    Catmxvcloamsegccactcount interface{}

    // Specifies OAM Segment Continuity check (CC)  Deactivation retry count. The
    // type is interface{} with range: 0..4294967295.
    Catmxvcloamsegccdeactcount interface{}

    // Specifies OAM Segment Continuity check (CC)  Activation/Deactivation retry
    // frequency. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Catmxvcloamsegccretryfreq interface{}

    // Specifies OAM Enable/Disable on the VC. true(1) indicates that OAM is
    // enabled on the VC. false(2) indicates that OAM is disabled on the VC. The
    // type is bool.
    Catmxvcloammanage interface{}

    // Indicates OAM loopback status of the VC. disabled(1)  --   No OAMs on this
    // VC. sent(2)      --   OAM sent, waiting for echo. received(3)  --   OAM
    // received from target. failed(4)    --   Last OAM did not return. The type
    // is Catmxvcloamloopbkstatus.
    Catmxvcloamloopbkstatus interface{}

    // Indicates the state of VC OAM. downRetry(1)   --  Loopback failed. Retry
    // sending                     loopbacks with retry frequency.                
    // VC is up. verified(2)    --  Loopback is successful. notVerified(3) --  Not
    // verified by loopback,                     AIS/RDI conditions are cleared.
    // upRetry(4)     --  Retry successive loopbacks.                     VC is
    // down. aisRDI(5)      --  Received AIS/RDI. Loopback are                    
    // not sent in this state. aisOut(6)      --  Sending AIS. Loopback and reply
    // are                     not sent in this state. notManaged(7)  --  VC is
    // not managed by OAM. The type is Catmxvcloamvcstate.
    Catmxvcloamvcstate interface{}

    // Indicates OAM End-to-end Continuity check (CC)  status. The type is
    // OamCCStatus.
    Catmxvcloamendccstatus interface{}

    // Indicates OAM Segment Continuity check (CC) status. The type is
    // OamCCStatus.
    Catmxvcloamsegccstatus interface{}

    // Indicates OAM End-to-end Continuity check (CC)  VC state. The type is
    // OamCCVcState.
    Catmxvcloamendccvcstate interface{}

    // Indicates OAM Segment Continuity check (CC) VC  state. The type is
    // OamCCVcState.
    Catmxvcloamsegccvcstate interface{}

    // Indicates the number of OAM cells received on  this VC. The type is
    // interface{} with range: 0..4294967295. Units are cells.
    Catmxvcloamcellsreceived interface{}

    // Indicates the number of OAM cells sent on  this VC. The type is interface{}
    // with range: 0..4294967295. Units are cells.
    Catmxvcloamcellssent interface{}

    // Indicates the number of OAM cells dropped on  this VC. The type is
    // interface{} with range: 0..4294967295. Units are cells.
    Catmxvcloamcellsdropped interface{}

    // Indicates the number of received OAM  F5 Alarm Indication Signal (AIS)
    // cells from the VC. The type is interface{} with range: 0..4294967295. Units
    // are cells.
    Catmxvcloaminf5Ais interface{}

    // Indicates the number of transmitted OAM  F5 Alarm Indication Signal (AIS)
    // cells to the VC. The type is interface{} with range: 0..4294967295. Units
    // are cells.
    Catmxvcloamoutf5Ais interface{}

    // Indicates the number of received OAM  F5 Remote Detection Indication (RDI)
    // cells from  the VC. The type is interface{} with range: 0..4294967295.
    // Units are cells.
    Catmxvcloaminf5Rdi interface{}

    // Indicates the number of transmitted OAM  F5 Remote Detection Indication
    // (RDI) cells to the VC. The type is interface{} with range: 0..4294967295.
    // Units are cells.
    Catmxvcloamoutf5Rdi interface{}
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetEntityData() *types.CommonEntityData {
    atmvclentry.EntityData.YFilter = atmvclentry.YFilter
    atmvclentry.EntityData.YangName = "atmVclEntry"
    atmvclentry.EntityData.BundleName = "cisco_ios_xe"
    atmvclentry.EntityData.ParentYangName = "atmVclTable"
    atmvclentry.EntityData.SegmentPath = "atmVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmvclentry.Atmvclvci) + "']"
    atmvclentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvclentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvclentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvclentry.EntityData.Children = make(map[string]types.YChild)
    atmvclentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmvclentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atmvclentry.Ifindex}
    atmvclentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", atmvclentry.Atmvclvpi}
    atmvclentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", atmvclentry.Atmvclvci}
    atmvclentry.EntityData.Leafs["atmVclAdminStatus"] = types.YLeaf{"Atmvcladminstatus", atmvclentry.Atmvcladminstatus}
    atmvclentry.EntityData.Leafs["atmVclOperStatus"] = types.YLeaf{"Atmvcloperstatus", atmvclentry.Atmvcloperstatus}
    atmvclentry.EntityData.Leafs["atmVclLastChange"] = types.YLeaf{"Atmvcllastchange", atmvclentry.Atmvcllastchange}
    atmvclentry.EntityData.Leafs["atmVclReceiveTrafficDescrIndex"] = types.YLeaf{"Atmvclreceivetrafficdescrindex", atmvclentry.Atmvclreceivetrafficdescrindex}
    atmvclentry.EntityData.Leafs["atmVclTransmitTrafficDescrIndex"] = types.YLeaf{"Atmvcltransmittrafficdescrindex", atmvclentry.Atmvcltransmittrafficdescrindex}
    atmvclentry.EntityData.Leafs["atmVccAalType"] = types.YLeaf{"Atmvccaaltype", atmvclentry.Atmvccaaltype}
    atmvclentry.EntityData.Leafs["atmVccAal5CpcsTransmitSduSize"] = types.YLeaf{"Atmvccaal5Cpcstransmitsdusize", atmvclentry.Atmvccaal5Cpcstransmitsdusize}
    atmvclentry.EntityData.Leafs["atmVccAal5CpcsReceiveSduSize"] = types.YLeaf{"Atmvccaal5Cpcsreceivesdusize", atmvclentry.Atmvccaal5Cpcsreceivesdusize}
    atmvclentry.EntityData.Leafs["atmVccAal5EncapsType"] = types.YLeaf{"Atmvccaal5Encapstype", atmvclentry.Atmvccaal5Encapstype}
    atmvclentry.EntityData.Leafs["atmVclCrossConnectIdentifier"] = types.YLeaf{"Atmvclcrossconnectidentifier", atmvclentry.Atmvclcrossconnectidentifier}
    atmvclentry.EntityData.Leafs["atmVclRowStatus"] = types.YLeaf{"Atmvclrowstatus", atmvclentry.Atmvclrowstatus}
    atmvclentry.EntityData.Leafs["atmVclCastType"] = types.YLeaf{"Atmvclcasttype", atmvclentry.Atmvclcasttype}
    atmvclentry.EntityData.Leafs["atmVclConnKind"] = types.YLeaf{"Atmvclconnkind", atmvclentry.Atmvclconnkind}
    atmvclentry.EntityData.Leafs["catmxVclOamLoopbackFreq"] = types.YLeaf{"Catmxvcloamloopbackfreq", atmvclentry.Catmxvcloamloopbackfreq}
    atmvclentry.EntityData.Leafs["catmxVclOamRetryFreq"] = types.YLeaf{"Catmxvcloamretryfreq", atmvclentry.Catmxvcloamretryfreq}
    atmvclentry.EntityData.Leafs["catmxVclOamUpRetryCount"] = types.YLeaf{"Catmxvcloamupretrycount", atmvclentry.Catmxvcloamupretrycount}
    atmvclentry.EntityData.Leafs["catmxVclOamDownRetryCount"] = types.YLeaf{"Catmxvcloamdownretrycount", atmvclentry.Catmxvcloamdownretrycount}
    atmvclentry.EntityData.Leafs["catmxVclOamEndCCActCount"] = types.YLeaf{"Catmxvcloamendccactcount", atmvclentry.Catmxvcloamendccactcount}
    atmvclentry.EntityData.Leafs["catmxVclOamEndCCDeActCount"] = types.YLeaf{"Catmxvcloamendccdeactcount", atmvclentry.Catmxvcloamendccdeactcount}
    atmvclentry.EntityData.Leafs["catmxVclOamEndCCRetryFreq"] = types.YLeaf{"Catmxvcloamendccretryfreq", atmvclentry.Catmxvcloamendccretryfreq}
    atmvclentry.EntityData.Leafs["catmxVclOamSegCCActCount"] = types.YLeaf{"Catmxvcloamsegccactcount", atmvclentry.Catmxvcloamsegccactcount}
    atmvclentry.EntityData.Leafs["catmxVclOamSegCCDeActCount"] = types.YLeaf{"Catmxvcloamsegccdeactcount", atmvclentry.Catmxvcloamsegccdeactcount}
    atmvclentry.EntityData.Leafs["catmxVclOamSegCCRetryFreq"] = types.YLeaf{"Catmxvcloamsegccretryfreq", atmvclentry.Catmxvcloamsegccretryfreq}
    atmvclentry.EntityData.Leafs["catmxVclOamManage"] = types.YLeaf{"Catmxvcloammanage", atmvclentry.Catmxvcloammanage}
    atmvclentry.EntityData.Leafs["catmxVclOamLoopBkStatus"] = types.YLeaf{"Catmxvcloamloopbkstatus", atmvclentry.Catmxvcloamloopbkstatus}
    atmvclentry.EntityData.Leafs["catmxVclOamVcState"] = types.YLeaf{"Catmxvcloamvcstate", atmvclentry.Catmxvcloamvcstate}
    atmvclentry.EntityData.Leafs["catmxVclOamEndCCStatus"] = types.YLeaf{"Catmxvcloamendccstatus", atmvclentry.Catmxvcloamendccstatus}
    atmvclentry.EntityData.Leafs["catmxVclOamSegCCStatus"] = types.YLeaf{"Catmxvcloamsegccstatus", atmvclentry.Catmxvcloamsegccstatus}
    atmvclentry.EntityData.Leafs["catmxVclOamEndCCVcState"] = types.YLeaf{"Catmxvcloamendccvcstate", atmvclentry.Catmxvcloamendccvcstate}
    atmvclentry.EntityData.Leafs["catmxVclOamSegCCVcState"] = types.YLeaf{"Catmxvcloamsegccvcstate", atmvclentry.Catmxvcloamsegccvcstate}
    atmvclentry.EntityData.Leafs["catmxVclOamCellsReceived"] = types.YLeaf{"Catmxvcloamcellsreceived", atmvclentry.Catmxvcloamcellsreceived}
    atmvclentry.EntityData.Leafs["catmxVclOamCellsSent"] = types.YLeaf{"Catmxvcloamcellssent", atmvclentry.Catmxvcloamcellssent}
    atmvclentry.EntityData.Leafs["catmxVclOamCellsDropped"] = types.YLeaf{"Catmxvcloamcellsdropped", atmvclentry.Catmxvcloamcellsdropped}
    atmvclentry.EntityData.Leafs["catmxVclOamInF5ais"] = types.YLeaf{"Catmxvcloaminf5Ais", atmvclentry.Catmxvcloaminf5Ais}
    atmvclentry.EntityData.Leafs["catmxVclOamOutF5ais"] = types.YLeaf{"Catmxvcloamoutf5Ais", atmvclentry.Catmxvcloamoutf5Ais}
    atmvclentry.EntityData.Leafs["catmxVclOamInF5rdi"] = types.YLeaf{"Catmxvcloaminf5Rdi", atmvclentry.Catmxvcloaminf5Rdi}
    atmvclentry.EntityData.Leafs["catmxVclOamOutF5rdi"] = types.YLeaf{"Catmxvcloamoutf5Rdi", atmvclentry.Catmxvcloamoutf5Rdi}
    return &(atmvclentry.EntityData)
}

// ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype represents LAN Emulation specification.
type ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype string

const (
    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexRoutedProtocol ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexRoutedProtocol"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexBridgedProtocol8023 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexBridgedProtocol8023"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexBridgedProtocol8025 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexBridgedProtocol8025"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexBridgedProtocol8026 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexBridgedProtocol8026"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexLANemulation8023 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexLANemulation8023"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_vcMultiplexLANemulation8025 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "vcMultiplexLANemulation8025"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_llcEncapsulation ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "llcEncapsulation"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_multiprotocolFrameRelaySscs ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "multiprotocolFrameRelaySscs"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_other ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "other"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype_unknown ATMMIB_Atmvcltable_Atmvclentry_Atmvccaal5Encapstype = "unknown"
)

// ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype represents the AAL type cannot be determined.
type ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype string

const (
    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_aal1 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "aal1"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_aal34 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "aal34"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_aal5 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "aal5"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_other ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "other"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_unknown ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "unknown"

    ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype_aal2 ATMMIB_Atmvcltable_Atmvclentry_Atmvccaaltype = "aal2"
)

// ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus represents failed(4)    --   Last OAM did not return.
type ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus string

const (
    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus_disabled ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus = "disabled"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus_sent ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus = "sent"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus_received ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus = "received"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus_failed ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamloopbkstatus = "failed"
)

// ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate represents notManaged(7)  --  VC is not managed by OAM.
type ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate string

const (
    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_downRetry ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "downRetry"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_verified ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "verified"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_notVerified ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "notVerified"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_upRetry ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "upRetry"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_aisRDI ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "aisRDI"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_aisOut ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "aisOut"

    ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate_notManaged ATMMIB_Atmvcltable_Atmvclentry_Catmxvcloamvcstate = "notManaged"
)

// ATMMIB_Atmvpcrossconnecttable
// The ATM VP Cross Connect table for PVCs.
// An entry in this table models two
// cross-connected VPLs.
// Each VPL must have its atmConnKind set
// to pvc(1).
type ATMMIB_Atmvpcrossconnecttable struct {
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
    // ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry.
    Atmvpcrossconnectentry []ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetEntityData() *types.CommonEntityData {
    atmvpcrossconnecttable.EntityData.YFilter = atmvpcrossconnecttable.YFilter
    atmvpcrossconnecttable.EntityData.YangName = "atmVpCrossConnectTable"
    atmvpcrossconnecttable.EntityData.BundleName = "cisco_ios_xe"
    atmvpcrossconnecttable.EntityData.ParentYangName = "ATM-MIB"
    atmvpcrossconnecttable.EntityData.SegmentPath = "atmVpCrossConnectTable"
    atmvpcrossconnecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvpcrossconnecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvpcrossconnecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvpcrossconnecttable.EntityData.Children = make(map[string]types.YChild)
    atmvpcrossconnecttable.EntityData.Children["atmVpCrossConnectEntry"] = types.YChild{"Atmvpcrossconnectentry", nil}
    for i := range atmvpcrossconnecttable.Atmvpcrossconnectentry {
        atmvpcrossconnecttable.EntityData.Children[types.GetSegmentPath(&atmvpcrossconnecttable.Atmvpcrossconnectentry[i])] = types.YChild{"Atmvpcrossconnectentry", &atmvpcrossconnecttable.Atmvpcrossconnectentry[i]}
    }
    atmvpcrossconnecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmvpcrossconnecttable.EntityData)
}

// ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry
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
type ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique value to identify this VP cross-connect.
    // For each VPL associated with this cross-connect, the agent reports this
    // cross-connect index value in the atmVplCrossConnectIdentifier attribute of
    // the corresponding atmVplTable entries. The type is interface{} with range:
    // 1..2147483647.
    Atmvpcrossconnectindex interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VP
    // cross-connect. The term low implies that this ATM interface has the
    // numerically lower ifIndex value than the other ATM interface identified in
    // the same atmVpCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    Atmvpcrossconnectlowifindex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VP cross-connect that is identified by atmVpCrossConnectLowIfIndex. The
    // type is interface{} with range: 0..4095.
    Atmvpcrossconnectlowvpi interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VP
    // cross-connect. The term high implies that this ATM interface has the
    // numerically higher ifIndex value than the  other ATM interface identified
    // in the same atmVpCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    Atmvpcrossconnecthighifindex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VP cross-connect that is identified by atmVpCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..4095.
    Atmvpcrossconnecthighvpi interface{}

    // The desired administrative status of this bi-directional VP cross-connect.
    // The type is AtmVorXAdminStatus.
    Atmvpcrossconnectadminstatus interface{}

    // The operational status of the VP cross-connect in one direction; (i.e.,
    // from the low to high direction). The type is AtmVorXOperStatus.
    Atmvpcrossconnectl2Hoperstatus interface{}

    // The operational status of the VP cross-connect in one direction; (i.e.,
    // from the high to low direction). The type is AtmVorXOperStatus.
    Atmvpcrossconnecth2Loperstatus interface{}

    // The value of sysUpTime at the time this VP cross-connect entered its
    // current operational state in the low to high direction. The type is
    // interface{} with range: 0..4294967295.
    Atmvpcrossconnectl2Hlastchange interface{}

    // The value of sysUpTime at the time this VP cross-connect entered its
    // current operational in the high to low direction. The type is interface{}
    // with range: 0..4294967295.
    Atmvpcrossconnecth2Llastchange interface{}

    // The status of this entry in the atmVpCrossConnectTable.  This object is
    // used to create a cross-connect for cross-connecting VPLs which are created
    // using the atmVplTable or to change or delete an existing cross-connect.
    // This object must be initially set to `createAndWait' or 'createAndGo'. To
    // turn on a VP cross-connect, the atmVpCrossConnectAdminStatus is set to
    // `up'. The type is RowStatus.
    Atmvpcrossconnectrowstatus interface{}
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetEntityData() *types.CommonEntityData {
    atmvpcrossconnectentry.EntityData.YFilter = atmvpcrossconnectentry.YFilter
    atmvpcrossconnectentry.EntityData.YangName = "atmVpCrossConnectEntry"
    atmvpcrossconnectentry.EntityData.BundleName = "cisco_ios_xe"
    atmvpcrossconnectentry.EntityData.ParentYangName = "atmVpCrossConnectTable"
    atmvpcrossconnectentry.EntityData.SegmentPath = "atmVpCrossConnectEntry" + "[atmVpCrossConnectIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectindex) + "']" + "[atmVpCrossConnectLowIfIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectlowifindex) + "']" + "[atmVpCrossConnectLowVpi='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectlowvpi) + "']" + "[atmVpCrossConnectHighIfIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnecthighifindex) + "']" + "[atmVpCrossConnectHighVpi='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnecthighvpi) + "']"
    atmvpcrossconnectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvpcrossconnectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvpcrossconnectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvpcrossconnectentry.EntityData.Children = make(map[string]types.YChild)
    atmvpcrossconnectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectIndex"] = types.YLeaf{"Atmvpcrossconnectindex", atmvpcrossconnectentry.Atmvpcrossconnectindex}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectLowIfIndex"] = types.YLeaf{"Atmvpcrossconnectlowifindex", atmvpcrossconnectentry.Atmvpcrossconnectlowifindex}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectLowVpi"] = types.YLeaf{"Atmvpcrossconnectlowvpi", atmvpcrossconnectentry.Atmvpcrossconnectlowvpi}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectHighIfIndex"] = types.YLeaf{"Atmvpcrossconnecthighifindex", atmvpcrossconnectentry.Atmvpcrossconnecthighifindex}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectHighVpi"] = types.YLeaf{"Atmvpcrossconnecthighvpi", atmvpcrossconnectentry.Atmvpcrossconnecthighvpi}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectAdminStatus"] = types.YLeaf{"Atmvpcrossconnectadminstatus", atmvpcrossconnectentry.Atmvpcrossconnectadminstatus}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectL2HOperStatus"] = types.YLeaf{"Atmvpcrossconnectl2Hoperstatus", atmvpcrossconnectentry.Atmvpcrossconnectl2Hoperstatus}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectH2LOperStatus"] = types.YLeaf{"Atmvpcrossconnecth2Loperstatus", atmvpcrossconnectentry.Atmvpcrossconnecth2Loperstatus}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectL2HLastChange"] = types.YLeaf{"Atmvpcrossconnectl2Hlastchange", atmvpcrossconnectentry.Atmvpcrossconnectl2Hlastchange}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectH2LLastChange"] = types.YLeaf{"Atmvpcrossconnecth2Llastchange", atmvpcrossconnectentry.Atmvpcrossconnecth2Llastchange}
    atmvpcrossconnectentry.EntityData.Leafs["atmVpCrossConnectRowStatus"] = types.YLeaf{"Atmvpcrossconnectrowstatus", atmvpcrossconnectentry.Atmvpcrossconnectrowstatus}
    return &(atmvpcrossconnectentry.EntityData)
}

// ATMMIB_Atmvccrossconnecttable
// The ATM VC Cross Connect table for PVCs.
// An entry in this table models two
// cross-connected VCLs.
// Each VCL must have its atmConnKind set
// to pvc(1).
type ATMMIB_Atmvccrossconnecttable struct {
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
    // ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry.
    Atmvccrossconnectentry []ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetEntityData() *types.CommonEntityData {
    atmvccrossconnecttable.EntityData.YFilter = atmvccrossconnecttable.YFilter
    atmvccrossconnecttable.EntityData.YangName = "atmVcCrossConnectTable"
    atmvccrossconnecttable.EntityData.BundleName = "cisco_ios_xe"
    atmvccrossconnecttable.EntityData.ParentYangName = "ATM-MIB"
    atmvccrossconnecttable.EntityData.SegmentPath = "atmVcCrossConnectTable"
    atmvccrossconnecttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvccrossconnecttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvccrossconnecttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvccrossconnecttable.EntityData.Children = make(map[string]types.YChild)
    atmvccrossconnecttable.EntityData.Children["atmVcCrossConnectEntry"] = types.YChild{"Atmvccrossconnectentry", nil}
    for i := range atmvccrossconnecttable.Atmvccrossconnectentry {
        atmvccrossconnecttable.EntityData.Children[types.GetSegmentPath(&atmvccrossconnecttable.Atmvccrossconnectentry[i])] = types.YChild{"Atmvccrossconnectentry", &atmvccrossconnecttable.Atmvccrossconnectentry[i]}
    }
    atmvccrossconnecttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmvccrossconnecttable.EntityData)
}

// ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry
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
type ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique value to identify this VC cross-connect.
    // For each VCL associated with this cross-connect, the agent reports this
    // cross-connect index value in the atmVclCrossConnectIdentifier attribute of
    // the corresponding atmVclTable entries. The type is interface{} with range:
    // 1..2147483647.
    Atmvccrossconnectindex interface{}

    // This attribute is a key. The ifIndex value of the ATM interface for this VC
    // cross-connect. The term low implies that this ATM interface has the
    // numerically lower ifIndex value than the other ATM interface identified in
    // the same atmVcCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    Atmvccrossconnectlowifindex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectLowIfIndex. The
    // type is interface{} with range: 0..4095.
    Atmvccrossconnectlowvpi interface{}

    // This attribute is a key. The VCI value at the ATM interface associated with
    // this VC cross-connect that is identified by atmVcCrossConnectLowIfIndex.
    // The type is interface{} with range: 0..65535.
    Atmvccrossconnectlowvci interface{}

    // This attribute is a key. The ifIndex value for the ATM interface for this
    // VC cross-connect. The term high implies that this ATM interface has the
    // numerically higher ifIndex value than the other ATM interface identified in
    // the same atmVcCrossConnectEntry. The type is interface{} with range:
    // 1..2147483647.
    Atmvccrossconnecthighifindex interface{}

    // This attribute is a key. The VPI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..4095.
    Atmvccrossconnecthighvpi interface{}

    // This attribute is a key. The VCI value at the ATM interface associated with
    // the VC cross-connect that is identified by atmVcCrossConnectHighIfIndex.
    // The type is interface{} with range: 0..65535.
    Atmvccrossconnecthighvci interface{}

    // The desired administrative status of this bi-directional VC cross-connect.
    // The type is AtmVorXAdminStatus.
    Atmvccrossconnectadminstatus interface{}

    // The current operational status of the VC cross-connect in one direction;
    // (i.e., from the low to high direction). The type is AtmVorXOperStatus.
    Atmvccrossconnectl2Hoperstatus interface{}

    // The current operational status of the VC cross-connect in one direction;
    // (i.e., from the high to low direction). The type is AtmVorXOperStatus.
    Atmvccrossconnecth2Loperstatus interface{}

    // The value of sysUpTime at the time this VC cross-connect entered its
    // current operational state in low to high direction. The type is interface{}
    // with range: 0..4294967295.
    Atmvccrossconnectl2Hlastchange interface{}

    // The value of sysUpTime at the time this VC cross-connect entered its
    // current operational state in high to low direction. The type is interface{}
    // with range: 0..4294967295.
    Atmvccrossconnecth2Llastchange interface{}

    // The status of this entry in the atmVcCrossConnectTable.  This object is
    // used to create a new cross-connect for cross-connecting VCLs which are
    // created using the atmVclTable or to change or delete existing
    // cross-connect. This object must be initially set to `createAndWait' or
    // 'createAndGo'. To turn on a VC cross-connect, the
    // atmVcCrossConnectAdminStatus is set to `up'. The type is RowStatus.
    Atmvccrossconnectrowstatus interface{}
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetEntityData() *types.CommonEntityData {
    atmvccrossconnectentry.EntityData.YFilter = atmvccrossconnectentry.YFilter
    atmvccrossconnectentry.EntityData.YangName = "atmVcCrossConnectEntry"
    atmvccrossconnectentry.EntityData.BundleName = "cisco_ios_xe"
    atmvccrossconnectentry.EntityData.ParentYangName = "atmVcCrossConnectTable"
    atmvccrossconnectentry.EntityData.SegmentPath = "atmVcCrossConnectEntry" + "[atmVcCrossConnectIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectindex) + "']" + "[atmVcCrossConnectLowIfIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowifindex) + "']" + "[atmVcCrossConnectLowVpi='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowvpi) + "']" + "[atmVcCrossConnectLowVci='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowvci) + "']" + "[atmVcCrossConnectHighIfIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighifindex) + "']" + "[atmVcCrossConnectHighVpi='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighvpi) + "']" + "[atmVcCrossConnectHighVci='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighvci) + "']"
    atmvccrossconnectentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmvccrossconnectentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmvccrossconnectentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmvccrossconnectentry.EntityData.Children = make(map[string]types.YChild)
    atmvccrossconnectentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectIndex"] = types.YLeaf{"Atmvccrossconnectindex", atmvccrossconnectentry.Atmvccrossconnectindex}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectLowIfIndex"] = types.YLeaf{"Atmvccrossconnectlowifindex", atmvccrossconnectentry.Atmvccrossconnectlowifindex}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectLowVpi"] = types.YLeaf{"Atmvccrossconnectlowvpi", atmvccrossconnectentry.Atmvccrossconnectlowvpi}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectLowVci"] = types.YLeaf{"Atmvccrossconnectlowvci", atmvccrossconnectentry.Atmvccrossconnectlowvci}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectHighIfIndex"] = types.YLeaf{"Atmvccrossconnecthighifindex", atmvccrossconnectentry.Atmvccrossconnecthighifindex}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectHighVpi"] = types.YLeaf{"Atmvccrossconnecthighvpi", atmvccrossconnectentry.Atmvccrossconnecthighvpi}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectHighVci"] = types.YLeaf{"Atmvccrossconnecthighvci", atmvccrossconnectentry.Atmvccrossconnecthighvci}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectAdminStatus"] = types.YLeaf{"Atmvccrossconnectadminstatus", atmvccrossconnectentry.Atmvccrossconnectadminstatus}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectL2HOperStatus"] = types.YLeaf{"Atmvccrossconnectl2Hoperstatus", atmvccrossconnectentry.Atmvccrossconnectl2Hoperstatus}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectH2LOperStatus"] = types.YLeaf{"Atmvccrossconnecth2Loperstatus", atmvccrossconnectentry.Atmvccrossconnecth2Loperstatus}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectL2HLastChange"] = types.YLeaf{"Atmvccrossconnectl2Hlastchange", atmvccrossconnectentry.Atmvccrossconnectl2Hlastchange}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectH2LLastChange"] = types.YLeaf{"Atmvccrossconnecth2Llastchange", atmvccrossconnectentry.Atmvccrossconnecth2Llastchange}
    atmvccrossconnectentry.EntityData.Leafs["atmVcCrossConnectRowStatus"] = types.YLeaf{"Atmvccrossconnectrowstatus", atmvccrossconnectentry.Atmvccrossconnectrowstatus}
    return &(atmvccrossconnectentry.EntityData)
}

// ATMMIB_Aal5Vcctable
// This table contains AAL5 VCC performance
// parameters.
type ATMMIB_Aal5Vcctable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains the AAL5 VCC performance parameters and is indexed by
    // ifIndex values of AAL5 interfaces and the associated VPI/VCI values. The
    // type is slice of ATMMIB_Aal5Vcctable_Aal5Vccentry.
    Aal5Vccentry []ATMMIB_Aal5Vcctable_Aal5Vccentry
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetEntityData() *types.CommonEntityData {
    aal5Vcctable.EntityData.YFilter = aal5Vcctable.YFilter
    aal5Vcctable.EntityData.YangName = "aal5VccTable"
    aal5Vcctable.EntityData.BundleName = "cisco_ios_xe"
    aal5Vcctable.EntityData.ParentYangName = "ATM-MIB"
    aal5Vcctable.EntityData.SegmentPath = "aal5VccTable"
    aal5Vcctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aal5Vcctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aal5Vcctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aal5Vcctable.EntityData.Children = make(map[string]types.YChild)
    aal5Vcctable.EntityData.Children["aal5VccEntry"] = types.YChild{"Aal5Vccentry", nil}
    for i := range aal5Vcctable.Aal5Vccentry {
        aal5Vcctable.EntityData.Children[types.GetSegmentPath(&aal5Vcctable.Aal5Vccentry[i])] = types.YChild{"Aal5Vccentry", &aal5Vcctable.Aal5Vccentry[i]}
    }
    aal5Vcctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aal5Vcctable.EntityData)
}

// ATMMIB_Aal5Vcctable_Aal5Vccentry
// This list contains the AAL5 VCC
// performance parameters and is indexed
// by ifIndex values of AAL5 interfaces
// and the associated VPI/VCI values.
type ATMMIB_Aal5Vcctable_Aal5Vccentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The VPI value of the AAL5 VCC at the interface
    // identified by the ifIndex. The type is interface{} with range: 0..4095.
    Aal5Vccvpi interface{}

    // This attribute is a key. The VCI value of the AAL5 VCC at the interface
    // identified by the ifIndex. The type is interface{} with range: 0..65535.
    Aal5Vccvci interface{}

    // The number of AAL5 CPCS PDUs received with CRC-32 errors on this AAL5 VCC
    // at the interface associated with an AAL5 entity. The type is interface{}
    // with range: 0..4294967295.
    Aal5Vcccrcerrors interface{}

    // The number of partially re-assembled AAL5 CPCS PDUs which were discarded on
    // this AAL5 VCC at the interface associated with an AAL5 entity because they
    // were not fully re-assembled within the required time period.  If the
    // re-assembly timer is not supported, then this object contains a zero value.
    // The type is interface{} with range: 0..4294967295.
    Aal5Vccsartimeouts interface{}

    // The number of AAL5 CPCS PDUs discarded on this AAL5 VCC at the interface
    // associated with an AAL5 entity because the AAL5 SDUs were too large. The
    // type is interface{} with range: 0..4294967295.
    Aal5Vccoversizedsdus interface{}

    // The number of AAL5 CPCS PDUs received on this AAL5 VCC at the interface
    // associated with an AAL5 entity. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Caal5Vccinpkts interface{}

    // The number of AAL5 CPCS PDUs transmitted on this AAL5 VCC at the interface
    // associated with an AAL5 entity. The type is interface{} with range:
    // 0..4294967295. Units are packets.
    Caal5Vccoutpkts interface{}

    // The number of AAL5 CPCS PDU octets received on this AAL5 VCC at the
    // interface associated with an AAL5 entity. The type is interface{} with
    // range: 0..4294967295. Units are octets.
    Caal5Vccinoctets interface{}

    // The number of AAL5 CPCS PDU octets transmitted on this AAL5  VCC at the
    // interface associated with an AAL5 entity. The type is interface{} with
    // range: 0..4294967295. Units are octets.
    Caal5Vccoutoctets interface{}

    // The number of AAL5 CPCS PDUs dropped at the  receive side of this AAL5 VCC
    // at the interface  associated with an AAL5 entity. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Caal5Vccindroppedpkts interface{}

    // The number of AAL5 CPCS PDUs dropped at the transmit side  of this AAL5 VCC
    // at the interface associated with an  AAL5 entity. The type is interface{}
    // with range: 0..4294967295. Units are packets.
    Caal5Vccoutdroppedpkts interface{}

    // The number of AAL5 CPCS PDU Octets dropped at the  receive side of this
    // AAL5 VCC at the interface  associated with an AAL5 entity. The type is
    // interface{} with range: 0..4294967295. Units are octets.
    Caal5Vccindroppedoctets interface{}

    // The number of AAL5 CPCS PDU Octets dropped at the  transmit side of this
    // AAL5 VCC at the interface  associated with an AAL5 entity. The type is
    // interface{} with range: 0..4294967295. Units are octets.
    Caal5Vccoutdroppedoctets interface{}

    // This is 64bit (High Capacity) version of cAal5VccInPkts  counters. The type
    // is interface{} with range: 0..18446744073709551615.
    Caal5Vcchcinpkts interface{}

    // This is 64bit (High Capacity) version of cAal5VccOutPkts  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    Caal5Vcchcoutpkts interface{}

    // This is 64bit (High Capacity) version of cAal5VccInOctets  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    Caal5Vcchcinoctets interface{}

    // This is 64bit (High Capacity) version of cAal5VccOutOctets  counters. The
    // type is interface{} with range: 0..18446744073709551615.
    Caal5Vcchcoutoctets interface{}

    // Boolean, if compression enabled for VCC. The type is bool.
    Caal5Vccextcompenabled interface{}

    // Boolean, TRUE if VCC is used to carry voice. The type is bool.
    Caal5Vccextvoice interface{}

    // Number of OAM F5 end to end loopback cells  received through the VCC. The
    // type is interface{} with range: 0..4294967295.
    Caal5Vccextinf5Oamcells interface{}

    // Number of OAM F5 end to end loopback cells sent  through the VCC. The type
    // is interface{} with range: 0..4294967295.
    Caal5Vccextoutf5Oamcells interface{}
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetEntityData() *types.CommonEntityData {
    aal5Vccentry.EntityData.YFilter = aal5Vccentry.YFilter
    aal5Vccentry.EntityData.YangName = "aal5VccEntry"
    aal5Vccentry.EntityData.BundleName = "cisco_ios_xe"
    aal5Vccentry.EntityData.ParentYangName = "aal5VccTable"
    aal5Vccentry.EntityData.SegmentPath = "aal5VccEntry" + "[ifIndex='" + fmt.Sprintf("%v", aal5Vccentry.Ifindex) + "']" + "[aal5VccVpi='" + fmt.Sprintf("%v", aal5Vccentry.Aal5Vccvpi) + "']" + "[aal5VccVci='" + fmt.Sprintf("%v", aal5Vccentry.Aal5Vccvci) + "']"
    aal5Vccentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aal5Vccentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aal5Vccentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aal5Vccentry.EntityData.Children = make(map[string]types.YChild)
    aal5Vccentry.EntityData.Leafs = make(map[string]types.YLeaf)
    aal5Vccentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", aal5Vccentry.Ifindex}
    aal5Vccentry.EntityData.Leafs["aal5VccVpi"] = types.YLeaf{"Aal5Vccvpi", aal5Vccentry.Aal5Vccvpi}
    aal5Vccentry.EntityData.Leafs["aal5VccVci"] = types.YLeaf{"Aal5Vccvci", aal5Vccentry.Aal5Vccvci}
    aal5Vccentry.EntityData.Leafs["aal5VccCrcErrors"] = types.YLeaf{"Aal5Vcccrcerrors", aal5Vccentry.Aal5Vcccrcerrors}
    aal5Vccentry.EntityData.Leafs["aal5VccSarTimeOuts"] = types.YLeaf{"Aal5Vccsartimeouts", aal5Vccentry.Aal5Vccsartimeouts}
    aal5Vccentry.EntityData.Leafs["aal5VccOverSizedSDUs"] = types.YLeaf{"Aal5Vccoversizedsdus", aal5Vccentry.Aal5Vccoversizedsdus}
    aal5Vccentry.EntityData.Leafs["cAal5VccInPkts"] = types.YLeaf{"Caal5Vccinpkts", aal5Vccentry.Caal5Vccinpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccOutPkts"] = types.YLeaf{"Caal5Vccoutpkts", aal5Vccentry.Caal5Vccoutpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccInOctets"] = types.YLeaf{"Caal5Vccinoctets", aal5Vccentry.Caal5Vccinoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccOutOctets"] = types.YLeaf{"Caal5Vccoutoctets", aal5Vccentry.Caal5Vccoutoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccInDroppedPkts"] = types.YLeaf{"Caal5Vccindroppedpkts", aal5Vccentry.Caal5Vccindroppedpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccOutDroppedPkts"] = types.YLeaf{"Caal5Vccoutdroppedpkts", aal5Vccentry.Caal5Vccoutdroppedpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccInDroppedOctets"] = types.YLeaf{"Caal5Vccindroppedoctets", aal5Vccentry.Caal5Vccindroppedoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccOutDroppedOctets"] = types.YLeaf{"Caal5Vccoutdroppedoctets", aal5Vccentry.Caal5Vccoutdroppedoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccHCInPkts"] = types.YLeaf{"Caal5Vcchcinpkts", aal5Vccentry.Caal5Vcchcinpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccHCOutPkts"] = types.YLeaf{"Caal5Vcchcoutpkts", aal5Vccentry.Caal5Vcchcoutpkts}
    aal5Vccentry.EntityData.Leafs["cAal5VccHCInOctets"] = types.YLeaf{"Caal5Vcchcinoctets", aal5Vccentry.Caal5Vcchcinoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccHCOutOctets"] = types.YLeaf{"Caal5Vcchcoutoctets", aal5Vccentry.Caal5Vcchcoutoctets}
    aal5Vccentry.EntityData.Leafs["cAal5VccExtCompEnabled"] = types.YLeaf{"Caal5Vccextcompenabled", aal5Vccentry.Caal5Vccextcompenabled}
    aal5Vccentry.EntityData.Leafs["cAal5VccExtVoice"] = types.YLeaf{"Caal5Vccextvoice", aal5Vccentry.Caal5Vccextvoice}
    aal5Vccentry.EntityData.Leafs["cAal5VccExtInF5OamCells"] = types.YLeaf{"Caal5Vccextinf5Oamcells", aal5Vccentry.Caal5Vccextinf5Oamcells}
    aal5Vccentry.EntityData.Leafs["cAal5VccExtOutF5OamCells"] = types.YLeaf{"Caal5Vccextoutf5Oamcells", aal5Vccentry.Caal5Vccextoutf5Oamcells}
    return &(aal5Vccentry.EntityData)
}

