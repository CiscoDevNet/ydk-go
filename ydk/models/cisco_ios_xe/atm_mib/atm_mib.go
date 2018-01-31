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
    parent types.Entity
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

func (aTMMIB *ATMMIB) GetFilter() yfilter.YFilter { return aTMMIB.YFilter }

func (aTMMIB *ATMMIB) SetFilter(yf yfilter.YFilter) { aTMMIB.YFilter = yf }

func (aTMMIB *ATMMIB) GetGoName(yname string) string {
    if yname == "atmMIBObjects" { return "Atmmibobjects" }
    if yname == "atmInterfaceConfTable" { return "Atminterfaceconftable" }
    if yname == "atmInterfaceDs3PlcpTable" { return "Atminterfaceds3Plcptable" }
    if yname == "atmInterfaceTCTable" { return "Atminterfacetctable" }
    if yname == "atmTrafficDescrParamTable" { return "Atmtrafficdescrparamtable" }
    if yname == "atmVplTable" { return "Atmvpltable" }
    if yname == "atmVclTable" { return "Atmvcltable" }
    if yname == "atmVpCrossConnectTable" { return "Atmvpcrossconnecttable" }
    if yname == "atmVcCrossConnectTable" { return "Atmvccrossconnecttable" }
    if yname == "aal5VccTable" { return "Aal5Vcctable" }
    return ""
}

func (aTMMIB *ATMMIB) GetSegmentPath() string {
    return "ATM-MIB:ATM-MIB"
}

func (aTMMIB *ATMMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmMIBObjects" {
        return &aTMMIB.Atmmibobjects
    }
    if childYangName == "atmInterfaceConfTable" {
        return &aTMMIB.Atminterfaceconftable
    }
    if childYangName == "atmInterfaceDs3PlcpTable" {
        return &aTMMIB.Atminterfaceds3Plcptable
    }
    if childYangName == "atmInterfaceTCTable" {
        return &aTMMIB.Atminterfacetctable
    }
    if childYangName == "atmTrafficDescrParamTable" {
        return &aTMMIB.Atmtrafficdescrparamtable
    }
    if childYangName == "atmVplTable" {
        return &aTMMIB.Atmvpltable
    }
    if childYangName == "atmVclTable" {
        return &aTMMIB.Atmvcltable
    }
    if childYangName == "atmVpCrossConnectTable" {
        return &aTMMIB.Atmvpcrossconnecttable
    }
    if childYangName == "atmVcCrossConnectTable" {
        return &aTMMIB.Atmvccrossconnecttable
    }
    if childYangName == "aal5VccTable" {
        return &aTMMIB.Aal5Vcctable
    }
    return nil
}

func (aTMMIB *ATMMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["atmMIBObjects"] = &aTMMIB.Atmmibobjects
    children["atmInterfaceConfTable"] = &aTMMIB.Atminterfaceconftable
    children["atmInterfaceDs3PlcpTable"] = &aTMMIB.Atminterfaceds3Plcptable
    children["atmInterfaceTCTable"] = &aTMMIB.Atminterfacetctable
    children["atmTrafficDescrParamTable"] = &aTMMIB.Atmtrafficdescrparamtable
    children["atmVplTable"] = &aTMMIB.Atmvpltable
    children["atmVclTable"] = &aTMMIB.Atmvcltable
    children["atmVpCrossConnectTable"] = &aTMMIB.Atmvpcrossconnecttable
    children["atmVcCrossConnectTable"] = &aTMMIB.Atmvccrossconnecttable
    children["aal5VccTable"] = &aTMMIB.Aal5Vcctable
    return children
}

func (aTMMIB *ATMMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aTMMIB *ATMMIB) GetBundleName() string { return "cisco_ios_xe" }

func (aTMMIB *ATMMIB) GetYangName() string { return "ATM-MIB" }

func (aTMMIB *ATMMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aTMMIB *ATMMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aTMMIB *ATMMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aTMMIB *ATMMIB) SetParent(parent types.Entity) { aTMMIB.parent = parent }

func (aTMMIB *ATMMIB) GetParent() types.Entity { return aTMMIB.parent }

func (aTMMIB *ATMMIB) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atmmibobjects
type ATMMIB_Atmmibobjects struct {
    parent types.Entity
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

func (atmmibobjects *ATMMIB_Atmmibobjects) GetFilter() yfilter.YFilter { return atmmibobjects.YFilter }

func (atmmibobjects *ATMMIB_Atmmibobjects) SetFilter(yf yfilter.YFilter) { atmmibobjects.YFilter = yf }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetGoName(yname string) string {
    if yname == "atmVpCrossConnectIndexNext" { return "Atmvpcrossconnectindexnext" }
    if yname == "atmVcCrossConnectIndexNext" { return "Atmvccrossconnectindexnext" }
    if yname == "atmTrafficDescrParamIndexNext" { return "Atmtrafficdescrparamindexnext" }
    return ""
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetSegmentPath() string {
    return "atmMIBObjects"
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atmVpCrossConnectIndexNext"] = atmmibobjects.Atmvpcrossconnectindexnext
    leafs["atmVcCrossConnectIndexNext"] = atmmibobjects.Atmvccrossconnectindexnext
    leafs["atmTrafficDescrParamIndexNext"] = atmmibobjects.Atmtrafficdescrparamindexnext
    return leafs
}

func (atmmibobjects *ATMMIB_Atmmibobjects) GetBundleName() string { return "cisco_ios_xe" }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetYangName() string { return "atmMIBObjects" }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmmibobjects *ATMMIB_Atmmibobjects) SetParent(parent types.Entity) { atmmibobjects.parent = parent }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetParent() types.Entity { return atmmibobjects.parent }

func (atmmibobjects *ATMMIB_Atmmibobjects) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atminterfaceconftable
// This table contains ATM local interface
// configuration parameters, one entry per ATM
// interface port.
type ATMMIB_Atminterfaceconftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains ATM interface configuration parameters and state
    // variables and is indexed by ifIndex values of ATM interfaces. The type is
    // slice of ATMMIB_Atminterfaceconftable_Atminterfaceconfentry.
    Atminterfaceconfentry []ATMMIB_Atminterfaceconftable_Atminterfaceconfentry
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetFilter() yfilter.YFilter { return atminterfaceconftable.YFilter }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) SetFilter(yf yfilter.YFilter) { atminterfaceconftable.YFilter = yf }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetGoName(yname string) string {
    if yname == "atmInterfaceConfEntry" { return "Atminterfaceconfentry" }
    return ""
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetSegmentPath() string {
    return "atmInterfaceConfTable"
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmInterfaceConfEntry" {
        for _, c := range atminterfaceconftable.Atminterfaceconfentry {
            if atminterfaceconftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atminterfaceconftable_Atminterfaceconfentry{}
        atminterfaceconftable.Atminterfaceconfentry = append(atminterfaceconftable.Atminterfaceconfentry, child)
        return &atminterfaceconftable.Atminterfaceconfentry[len(atminterfaceconftable.Atminterfaceconfentry)-1]
    }
    return nil
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atminterfaceconftable.Atminterfaceconfentry {
        children[atminterfaceconftable.Atminterfaceconfentry[i].GetSegmentPath()] = &atminterfaceconftable.Atminterfaceconfentry[i]
    }
    return children
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetYangName() string { return "atmInterfaceConfTable" }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) SetParent(parent types.Entity) { atminterfaceconftable.parent = parent }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetParent() types.Entity { return atminterfaceconftable.parent }

func (atminterfaceconftable *ATMMIB_Atminterfaceconftable) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atminterfaceconftable_Atminterfaceconfentry
// This list contains ATM interface configuration
// parameters and state variables and is indexed
// by ifIndex values of ATM interfaces.
type ATMMIB_Atminterfaceconftable_Atminterfaceconfentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetFilter() yfilter.YFilter { return atminterfaceconfentry.YFilter }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) SetFilter(yf yfilter.YFilter) { atminterfaceconfentry.YFilter = yf }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmInterfaceMaxVpcs" { return "Atminterfacemaxvpcs" }
    if yname == "atmInterfaceMaxVccs" { return "Atminterfacemaxvccs" }
    if yname == "atmInterfaceConfVpcs" { return "Atminterfaceconfvpcs" }
    if yname == "atmInterfaceConfVccs" { return "Atminterfaceconfvccs" }
    if yname == "atmInterfaceMaxActiveVpiBits" { return "Atminterfacemaxactivevpibits" }
    if yname == "atmInterfaceMaxActiveVciBits" { return "Atminterfacemaxactivevcibits" }
    if yname == "atmInterfaceIlmiVpi" { return "Atminterfaceilmivpi" }
    if yname == "atmInterfaceIlmiVci" { return "Atminterfaceilmivci" }
    if yname == "atmInterfaceAddressType" { return "Atminterfaceaddresstype" }
    if yname == "atmInterfaceAdminAddress" { return "Atminterfaceadminaddress" }
    if yname == "atmInterfaceMyNeighborIpAddress" { return "Atminterfacemyneighboripaddress" }
    if yname == "atmInterfaceMyNeighborIfName" { return "Atminterfacemyneighborifname" }
    if yname == "atmInterfaceCurrentMaxVpiBits" { return "Atminterfacecurrentmaxvpibits" }
    if yname == "atmInterfaceCurrentMaxVciBits" { return "Atminterfacecurrentmaxvcibits" }
    if yname == "atmInterfaceSubscrAddress" { return "Atminterfacesubscraddress" }
    if yname == "atmIntfPvcFailures" { return "Atmintfpvcfailures" }
    if yname == "atmIntfCurrentlyFailingPVcls" { return "Atmintfcurrentlyfailingpvcls" }
    if yname == "atmIntfPvcFailuresTrapEnable" { return "Atmintfpvcfailurestrapenable" }
    if yname == "atmIntfPvcNotificationInterval" { return "Atmintfpvcnotificationinterval" }
    if yname == "atmPreviouslyFailedPVclInterval" { return "Atmpreviouslyfailedpvclinterval" }
    if yname == "catmIntfCurrentlyDownToUpPVcls" { return "Catmintfcurrentlydowntouppvcls" }
    if yname == "catmIntfOAMFailedPVcls" { return "Catmintfoamfailedpvcls" }
    if yname == "catmIntfCurrentOAMFailingPVcls" { return "Catmintfcurrentoamfailingpvcls" }
    if yname == "catmIntfSegCCOAMFailedPVcls" { return "Catmintfsegccoamfailedpvcls" }
    if yname == "catmIntfCurSegCCOAMFailingPVcls" { return "Catmintfcursegccoamfailingpvcls" }
    if yname == "catmIntfEndCCOAMFailedPVcls" { return "Catmintfendccoamfailedpvcls" }
    if yname == "catmIntfCurEndCCOAMFailingPVcls" { return "Catmintfcurendccoamfailingpvcls" }
    if yname == "catmIntfAISRDIOAMFailedPVcls" { return "Catmintfaisrdioamfailedpvcls" }
    if yname == "catmIntfCurAISRDIOAMFailingPVcls" { return "Catmintfcuraisrdioamfailingpvcls" }
    if yname == "catmIntfAnyOAMFailedPVcls" { return "Catmintfanyoamfailedpvcls" }
    if yname == "catmIntfCurAnyOAMFailingPVcls" { return "Catmintfcuranyoamfailingpvcls" }
    if yname == "catmIntfTypeOfOAMFailure" { return "Catmintftypeofoamfailure" }
    if yname == "catmIntfOAMRcovedPVcls" { return "Catmintfoamrcovedpvcls" }
    if yname == "catmIntfCurrentOAMRcovingPVcls" { return "Catmintfcurrentoamrcovingpvcls" }
    if yname == "catmIntfSegCCOAMRcovedPVcls" { return "Catmintfsegccoamrcovedpvcls" }
    if yname == "catmIntfCurSegCCOAMRcovingPVcls" { return "Catmintfcursegccoamrcovingpvcls" }
    if yname == "catmIntfEndCCOAMRcovedPVcls" { return "Catmintfendccoamrcovedpvcls" }
    if yname == "catmIntfCurEndCCOAMRcovingPVcls" { return "Catmintfcurendccoamrcovingpvcls" }
    if yname == "catmIntfAISRDIOAMRcovedPVcls" { return "Catmintfaisrdioamrcovedpvcls" }
    if yname == "catmIntfCurAISRDIOAMRcovingPVcls" { return "Catmintfcuraisrdioamrcovingpvcls" }
    if yname == "catmIntfAnyOAMRcovedPVcls" { return "Catmintfanyoamrcovedpvcls" }
    if yname == "catmIntfCurAnyOAMRcovingPVcls" { return "Catmintfcuranyoamrcovingpvcls" }
    if yname == "catmIntfTypeOfOAMRecover" { return "Catmintftypeofoamrecover" }
    if yname == "atmIntfCurrentlyDownToUpPVcls" { return "Atmintfcurrentlydowntouppvcls" }
    if yname == "atmIntfOAMFailedPVcls" { return "Atmintfoamfailedpvcls" }
    if yname == "atmIntfCurrentlyOAMFailingPVcls" { return "Atmintfcurrentlyoamfailingpvcls" }
    return ""
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetSegmentPath() string {
    return "atmInterfaceConfEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfaceconfentry.Ifindex) + "']"
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atminterfaceconfentry.Ifindex
    leafs["atmInterfaceMaxVpcs"] = atminterfaceconfentry.Atminterfacemaxvpcs
    leafs["atmInterfaceMaxVccs"] = atminterfaceconfentry.Atminterfacemaxvccs
    leafs["atmInterfaceConfVpcs"] = atminterfaceconfentry.Atminterfaceconfvpcs
    leafs["atmInterfaceConfVccs"] = atminterfaceconfentry.Atminterfaceconfvccs
    leafs["atmInterfaceMaxActiveVpiBits"] = atminterfaceconfentry.Atminterfacemaxactivevpibits
    leafs["atmInterfaceMaxActiveVciBits"] = atminterfaceconfentry.Atminterfacemaxactivevcibits
    leafs["atmInterfaceIlmiVpi"] = atminterfaceconfentry.Atminterfaceilmivpi
    leafs["atmInterfaceIlmiVci"] = atminterfaceconfentry.Atminterfaceilmivci
    leafs["atmInterfaceAddressType"] = atminterfaceconfentry.Atminterfaceaddresstype
    leafs["atmInterfaceAdminAddress"] = atminterfaceconfentry.Atminterfaceadminaddress
    leafs["atmInterfaceMyNeighborIpAddress"] = atminterfaceconfentry.Atminterfacemyneighboripaddress
    leafs["atmInterfaceMyNeighborIfName"] = atminterfaceconfentry.Atminterfacemyneighborifname
    leafs["atmInterfaceCurrentMaxVpiBits"] = atminterfaceconfentry.Atminterfacecurrentmaxvpibits
    leafs["atmInterfaceCurrentMaxVciBits"] = atminterfaceconfentry.Atminterfacecurrentmaxvcibits
    leafs["atmInterfaceSubscrAddress"] = atminterfaceconfentry.Atminterfacesubscraddress
    leafs["atmIntfPvcFailures"] = atminterfaceconfentry.Atmintfpvcfailures
    leafs["atmIntfCurrentlyFailingPVcls"] = atminterfaceconfentry.Atmintfcurrentlyfailingpvcls
    leafs["atmIntfPvcFailuresTrapEnable"] = atminterfaceconfentry.Atmintfpvcfailurestrapenable
    leafs["atmIntfPvcNotificationInterval"] = atminterfaceconfentry.Atmintfpvcnotificationinterval
    leafs["atmPreviouslyFailedPVclInterval"] = atminterfaceconfentry.Atmpreviouslyfailedpvclinterval
    leafs["catmIntfCurrentlyDownToUpPVcls"] = atminterfaceconfentry.Catmintfcurrentlydowntouppvcls
    leafs["catmIntfOAMFailedPVcls"] = atminterfaceconfentry.Catmintfoamfailedpvcls
    leafs["catmIntfCurrentOAMFailingPVcls"] = atminterfaceconfentry.Catmintfcurrentoamfailingpvcls
    leafs["catmIntfSegCCOAMFailedPVcls"] = atminterfaceconfentry.Catmintfsegccoamfailedpvcls
    leafs["catmIntfCurSegCCOAMFailingPVcls"] = atminterfaceconfentry.Catmintfcursegccoamfailingpvcls
    leafs["catmIntfEndCCOAMFailedPVcls"] = atminterfaceconfentry.Catmintfendccoamfailedpvcls
    leafs["catmIntfCurEndCCOAMFailingPVcls"] = atminterfaceconfentry.Catmintfcurendccoamfailingpvcls
    leafs["catmIntfAISRDIOAMFailedPVcls"] = atminterfaceconfentry.Catmintfaisrdioamfailedpvcls
    leafs["catmIntfCurAISRDIOAMFailingPVcls"] = atminterfaceconfentry.Catmintfcuraisrdioamfailingpvcls
    leafs["catmIntfAnyOAMFailedPVcls"] = atminterfaceconfentry.Catmintfanyoamfailedpvcls
    leafs["catmIntfCurAnyOAMFailingPVcls"] = atminterfaceconfentry.Catmintfcuranyoamfailingpvcls
    leafs["catmIntfTypeOfOAMFailure"] = atminterfaceconfentry.Catmintftypeofoamfailure
    leafs["catmIntfOAMRcovedPVcls"] = atminterfaceconfentry.Catmintfoamrcovedpvcls
    leafs["catmIntfCurrentOAMRcovingPVcls"] = atminterfaceconfentry.Catmintfcurrentoamrcovingpvcls
    leafs["catmIntfSegCCOAMRcovedPVcls"] = atminterfaceconfentry.Catmintfsegccoamrcovedpvcls
    leafs["catmIntfCurSegCCOAMRcovingPVcls"] = atminterfaceconfentry.Catmintfcursegccoamrcovingpvcls
    leafs["catmIntfEndCCOAMRcovedPVcls"] = atminterfaceconfentry.Catmintfendccoamrcovedpvcls
    leafs["catmIntfCurEndCCOAMRcovingPVcls"] = atminterfaceconfentry.Catmintfcurendccoamrcovingpvcls
    leafs["catmIntfAISRDIOAMRcovedPVcls"] = atminterfaceconfentry.Catmintfaisrdioamrcovedpvcls
    leafs["catmIntfCurAISRDIOAMRcovingPVcls"] = atminterfaceconfentry.Catmintfcuraisrdioamrcovingpvcls
    leafs["catmIntfAnyOAMRcovedPVcls"] = atminterfaceconfentry.Catmintfanyoamrcovedpvcls
    leafs["catmIntfCurAnyOAMRcovingPVcls"] = atminterfaceconfentry.Catmintfcuranyoamrcovingpvcls
    leafs["catmIntfTypeOfOAMRecover"] = atminterfaceconfentry.Catmintftypeofoamrecover
    leafs["atmIntfCurrentlyDownToUpPVcls"] = atminterfaceconfentry.Atmintfcurrentlydowntouppvcls
    leafs["atmIntfOAMFailedPVcls"] = atminterfaceconfentry.Atmintfoamfailedpvcls
    leafs["atmIntfCurrentlyOAMFailingPVcls"] = atminterfaceconfentry.Atmintfcurrentlyoamfailingpvcls
    return leafs
}

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetYangName() string { return "atmInterfaceConfEntry" }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) SetParent(parent types.Entity) { atminterfaceconfentry.parent = parent }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetParent() types.Entity { return atminterfaceconfentry.parent }

func (atminterfaceconfentry *ATMMIB_Atminterfaceconftable_Atminterfaceconfentry) GetParentYangName() string { return "atmInterfaceConfTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains DS3 PLCP parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry.
    Atminterfaceds3Plcpentry []ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetFilter() yfilter.YFilter { return atminterfaceds3Plcptable.YFilter }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) SetFilter(yf yfilter.YFilter) { atminterfaceds3Plcptable.YFilter = yf }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetGoName(yname string) string {
    if yname == "atmInterfaceDs3PlcpEntry" { return "Atminterfaceds3Plcpentry" }
    return ""
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetSegmentPath() string {
    return "atmInterfaceDs3PlcpTable"
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmInterfaceDs3PlcpEntry" {
        for _, c := range atminterfaceds3Plcptable.Atminterfaceds3Plcpentry {
            if atminterfaceds3Plcptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry{}
        atminterfaceds3Plcptable.Atminterfaceds3Plcpentry = append(atminterfaceds3Plcptable.Atminterfaceds3Plcpentry, child)
        return &atminterfaceds3Plcptable.Atminterfaceds3Plcpentry[len(atminterfaceds3Plcptable.Atminterfaceds3Plcpentry)-1]
    }
    return nil
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atminterfaceds3Plcptable.Atminterfaceds3Plcpentry {
        children[atminterfaceds3Plcptable.Atminterfaceds3Plcpentry[i].GetSegmentPath()] = &atminterfaceds3Plcptable.Atminterfaceds3Plcpentry[i]
    }
    return children
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetYangName() string { return "atmInterfaceDs3PlcpTable" }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) SetParent(parent types.Entity) { atminterfaceds3Plcptable.parent = parent }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetParent() types.Entity { return atminterfaceds3Plcptable.parent }

func (atminterfaceds3Plcptable *ATMMIB_Atminterfaceds3Plcptable) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry
// This list contains DS3 PLCP parameters and
// state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry struct {
    parent types.Entity
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

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetFilter() yfilter.YFilter { return atminterfaceds3Plcpentry.YFilter }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) SetFilter(yf yfilter.YFilter) { atminterfaceds3Plcpentry.YFilter = yf }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmInterfaceDs3PlcpSEFSs" { return "Atminterfaceds3Plcpsefss" }
    if yname == "atmInterfaceDs3PlcpAlarmState" { return "Atminterfaceds3Plcpalarmstate" }
    if yname == "atmInterfaceDs3PlcpUASs" { return "Atminterfaceds3Plcpuass" }
    return ""
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetSegmentPath() string {
    return "atmInterfaceDs3PlcpEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfaceds3Plcpentry.Ifindex) + "']"
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atminterfaceds3Plcpentry.Ifindex
    leafs["atmInterfaceDs3PlcpSEFSs"] = atminterfaceds3Plcpentry.Atminterfaceds3Plcpsefss
    leafs["atmInterfaceDs3PlcpAlarmState"] = atminterfaceds3Plcpentry.Atminterfaceds3Plcpalarmstate
    leafs["atmInterfaceDs3PlcpUASs"] = atminterfaceds3Plcpentry.Atminterfaceds3Plcpuass
    return leafs
}

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetYangName() string { return "atmInterfaceDs3PlcpEntry" }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) SetParent(parent types.Entity) { atminterfaceds3Plcpentry.parent = parent }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetParent() types.Entity { return atminterfaceds3Plcpentry.parent }

func (atminterfaceds3Plcpentry *ATMMIB_Atminterfaceds3Plcptable_Atminterfaceds3Plcpentry) GetParentYangName() string { return "atmInterfaceDs3PlcpTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains TC Sublayer parameters and state variables at the ATM
    // interface and is indexed by the ifIndex value of the ATM interface. The
    // type is slice of ATMMIB_Atminterfacetctable_Atminterfacetcentry.
    Atminterfacetcentry []ATMMIB_Atminterfacetctable_Atminterfacetcentry
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetFilter() yfilter.YFilter { return atminterfacetctable.YFilter }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) SetFilter(yf yfilter.YFilter) { atminterfacetctable.YFilter = yf }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetGoName(yname string) string {
    if yname == "atmInterfaceTCEntry" { return "Atminterfacetcentry" }
    return ""
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetSegmentPath() string {
    return "atmInterfaceTCTable"
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmInterfaceTCEntry" {
        for _, c := range atminterfacetctable.Atminterfacetcentry {
            if atminterfacetctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atminterfacetctable_Atminterfacetcentry{}
        atminterfacetctable.Atminterfacetcentry = append(atminterfacetctable.Atminterfacetcentry, child)
        return &atminterfacetctable.Atminterfacetcentry[len(atminterfacetctable.Atminterfacetcentry)-1]
    }
    return nil
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atminterfacetctable.Atminterfacetcentry {
        children[atminterfacetctable.Atminterfacetcentry[i].GetSegmentPath()] = &atminterfacetctable.Atminterfacetcentry[i]
    }
    return children
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetYangName() string { return "atmInterfaceTCTable" }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) SetParent(parent types.Entity) { atminterfacetctable.parent = parent }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetParent() types.Entity { return atminterfacetctable.parent }

func (atminterfacetctable *ATMMIB_Atminterfacetctable) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atminterfacetctable_Atminterfacetcentry
// This list contains TC Sublayer parameters
// and state variables at the ATM interface and is
// indexed by the ifIndex value of the ATM interface.
type ATMMIB_Atminterfacetctable_Atminterfacetcentry struct {
    parent types.Entity
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

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetFilter() yfilter.YFilter { return atminterfacetcentry.YFilter }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) SetFilter(yf yfilter.YFilter) { atminterfacetcentry.YFilter = yf }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmInterfaceOCDEvents" { return "Atminterfaceocdevents" }
    if yname == "atmInterfaceTCAlarmState" { return "Atminterfacetcalarmstate" }
    return ""
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetSegmentPath() string {
    return "atmInterfaceTCEntry" + "[ifIndex='" + fmt.Sprintf("%v", atminterfacetcentry.Ifindex) + "']"
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atminterfacetcentry.Ifindex
    leafs["atmInterfaceOCDEvents"] = atminterfacetcentry.Atminterfaceocdevents
    leafs["atmInterfaceTCAlarmState"] = atminterfacetcentry.Atminterfacetcalarmstate
    return leafs
}

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetBundleName() string { return "cisco_ios_xe" }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetYangName() string { return "atmInterfaceTCEntry" }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) SetParent(parent types.Entity) { atminterfacetcentry.parent = parent }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetParent() types.Entity { return atminterfacetcentry.parent }

func (atminterfacetcentry *ATMMIB_Atminterfacetctable_Atminterfacetcentry) GetParentYangName() string { return "atmInterfaceTCTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains ATM traffic descriptor type and the associated
    // parameters. The type is slice of
    // ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry.
    Atmtrafficdescrparamentry []ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetFilter() yfilter.YFilter { return atmtrafficdescrparamtable.YFilter }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) SetFilter(yf yfilter.YFilter) { atmtrafficdescrparamtable.YFilter = yf }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetGoName(yname string) string {
    if yname == "atmTrafficDescrParamEntry" { return "Atmtrafficdescrparamentry" }
    return ""
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetSegmentPath() string {
    return "atmTrafficDescrParamTable"
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmTrafficDescrParamEntry" {
        for _, c := range atmtrafficdescrparamtable.Atmtrafficdescrparamentry {
            if atmtrafficdescrparamtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry{}
        atmtrafficdescrparamtable.Atmtrafficdescrparamentry = append(atmtrafficdescrparamtable.Atmtrafficdescrparamentry, child)
        return &atmtrafficdescrparamtable.Atmtrafficdescrparamentry[len(atmtrafficdescrparamtable.Atmtrafficdescrparamentry)-1]
    }
    return nil
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmtrafficdescrparamtable.Atmtrafficdescrparamentry {
        children[atmtrafficdescrparamtable.Atmtrafficdescrparamentry[i].GetSegmentPath()] = &atmtrafficdescrparamtable.Atmtrafficdescrparamentry[i]
    }
    return children
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetBundleName() string { return "cisco_ios_xe" }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetYangName() string { return "atmTrafficDescrParamTable" }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) SetParent(parent types.Entity) { atmtrafficdescrparamtable.parent = parent }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetParent() types.Entity { return atmtrafficdescrparamtable.parent }

func (atmtrafficdescrparamtable *ATMMIB_Atmtrafficdescrparamtable) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry
// This list contains ATM traffic descriptor
// type and the associated parameters.
type ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetFilter() yfilter.YFilter { return atmtrafficdescrparamentry.YFilter }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) SetFilter(yf yfilter.YFilter) { atmtrafficdescrparamentry.YFilter = yf }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetGoName(yname string) string {
    if yname == "atmTrafficDescrParamIndex" { return "Atmtrafficdescrparamindex" }
    if yname == "atmTrafficDescrType" { return "Atmtrafficdescrtype" }
    if yname == "atmTrafficDescrParam1" { return "Atmtrafficdescrparam1" }
    if yname == "atmTrafficDescrParam2" { return "Atmtrafficdescrparam2" }
    if yname == "atmTrafficDescrParam3" { return "Atmtrafficdescrparam3" }
    if yname == "atmTrafficDescrParam4" { return "Atmtrafficdescrparam4" }
    if yname == "atmTrafficDescrParam5" { return "Atmtrafficdescrparam5" }
    if yname == "atmTrafficQoSClass" { return "Atmtrafficqosclass" }
    if yname == "atmTrafficDescrRowStatus" { return "Atmtrafficdescrrowstatus" }
    if yname == "atmServiceCategory" { return "Atmservicecategory" }
    if yname == "atmTrafficFrameDiscard" { return "Atmtrafficframediscard" }
    return ""
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetSegmentPath() string {
    return "atmTrafficDescrParamEntry" + "[atmTrafficDescrParamIndex='" + fmt.Sprintf("%v", atmtrafficdescrparamentry.Atmtrafficdescrparamindex) + "']"
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atmTrafficDescrParamIndex"] = atmtrafficdescrparamentry.Atmtrafficdescrparamindex
    leafs["atmTrafficDescrType"] = atmtrafficdescrparamentry.Atmtrafficdescrtype
    leafs["atmTrafficDescrParam1"] = atmtrafficdescrparamentry.Atmtrafficdescrparam1
    leafs["atmTrafficDescrParam2"] = atmtrafficdescrparamentry.Atmtrafficdescrparam2
    leafs["atmTrafficDescrParam3"] = atmtrafficdescrparamentry.Atmtrafficdescrparam3
    leafs["atmTrafficDescrParam4"] = atmtrafficdescrparamentry.Atmtrafficdescrparam4
    leafs["atmTrafficDescrParam5"] = atmtrafficdescrparamentry.Atmtrafficdescrparam5
    leafs["atmTrafficQoSClass"] = atmtrafficdescrparamentry.Atmtrafficqosclass
    leafs["atmTrafficDescrRowStatus"] = atmtrafficdescrparamentry.Atmtrafficdescrrowstatus
    leafs["atmServiceCategory"] = atmtrafficdescrparamentry.Atmservicecategory
    leafs["atmTrafficFrameDiscard"] = atmtrafficdescrparamentry.Atmtrafficframediscard
    return leafs
}

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetYangName() string { return "atmTrafficDescrParamEntry" }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) SetParent(parent types.Entity) { atmtrafficdescrparamentry.parent = parent }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetParent() types.Entity { return atmtrafficdescrparamentry.parent }

func (atmtrafficdescrparamentry *ATMMIB_Atmtrafficdescrparamtable_Atmtrafficdescrparamentry) GetParentYangName() string { return "atmTrafficDescrParamTable" }

// ATMMIB_Atmvpltable
// The Virtual Path Link (VPL) table.  A
// bi-directional VPL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
// Entries are not present in this table for
// the VPIs used by entries in the atmVclTable.
type ATMMIB_Atmvpltable struct {
    parent types.Entity
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

func (atmvpltable *ATMMIB_Atmvpltable) GetFilter() yfilter.YFilter { return atmvpltable.YFilter }

func (atmvpltable *ATMMIB_Atmvpltable) SetFilter(yf yfilter.YFilter) { atmvpltable.YFilter = yf }

func (atmvpltable *ATMMIB_Atmvpltable) GetGoName(yname string) string {
    if yname == "atmVplEntry" { return "Atmvplentry" }
    return ""
}

func (atmvpltable *ATMMIB_Atmvpltable) GetSegmentPath() string {
    return "atmVplTable"
}

func (atmvpltable *ATMMIB_Atmvpltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmVplEntry" {
        for _, c := range atmvpltable.Atmvplentry {
            if atmvpltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atmvpltable_Atmvplentry{}
        atmvpltable.Atmvplentry = append(atmvpltable.Atmvplentry, child)
        return &atmvpltable.Atmvplentry[len(atmvpltable.Atmvplentry)-1]
    }
    return nil
}

func (atmvpltable *ATMMIB_Atmvpltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmvpltable.Atmvplentry {
        children[atmvpltable.Atmvplentry[i].GetSegmentPath()] = &atmvpltable.Atmvplentry[i]
    }
    return children
}

func (atmvpltable *ATMMIB_Atmvpltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmvpltable *ATMMIB_Atmvpltable) GetBundleName() string { return "cisco_ios_xe" }

func (atmvpltable *ATMMIB_Atmvpltable) GetYangName() string { return "atmVplTable" }

func (atmvpltable *ATMMIB_Atmvpltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvpltable *ATMMIB_Atmvpltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvpltable *ATMMIB_Atmvpltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvpltable *ATMMIB_Atmvpltable) SetParent(parent types.Entity) { atmvpltable.parent = parent }

func (atmvpltable *ATMMIB_Atmvpltable) GetParent() types.Entity { return atmvpltable.parent }

func (atmvpltable *ATMMIB_Atmvpltable) GetParentYangName() string { return "ATM-MIB" }

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
    parent types.Entity
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

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetFilter() yfilter.YFilter { return atmvplentry.YFilter }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) SetFilter(yf yfilter.YFilter) { atmvplentry.YFilter = yf }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVplVpi" { return "Atmvplvpi" }
    if yname == "atmVplAdminStatus" { return "Atmvpladminstatus" }
    if yname == "atmVplOperStatus" { return "Atmvploperstatus" }
    if yname == "atmVplLastChange" { return "Atmvpllastchange" }
    if yname == "atmVplReceiveTrafficDescrIndex" { return "Atmvplreceivetrafficdescrindex" }
    if yname == "atmVplTransmitTrafficDescrIndex" { return "Atmvpltransmittrafficdescrindex" }
    if yname == "atmVplCrossConnectIdentifier" { return "Atmvplcrossconnectidentifier" }
    if yname == "atmVplRowStatus" { return "Atmvplrowstatus" }
    if yname == "atmVplCastType" { return "Atmvplcasttype" }
    if yname == "atmVplConnKind" { return "Atmvplconnkind" }
    return ""
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetSegmentPath() string {
    return "atmVplEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmvplentry.Ifindex) + "']" + "[atmVplVpi='" + fmt.Sprintf("%v", atmvplentry.Atmvplvpi) + "']"
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atmvplentry.Ifindex
    leafs["atmVplVpi"] = atmvplentry.Atmvplvpi
    leafs["atmVplAdminStatus"] = atmvplentry.Atmvpladminstatus
    leafs["atmVplOperStatus"] = atmvplentry.Atmvploperstatus
    leafs["atmVplLastChange"] = atmvplentry.Atmvpllastchange
    leafs["atmVplReceiveTrafficDescrIndex"] = atmvplentry.Atmvplreceivetrafficdescrindex
    leafs["atmVplTransmitTrafficDescrIndex"] = atmvplentry.Atmvpltransmittrafficdescrindex
    leafs["atmVplCrossConnectIdentifier"] = atmvplentry.Atmvplcrossconnectidentifier
    leafs["atmVplRowStatus"] = atmvplentry.Atmvplrowstatus
    leafs["atmVplCastType"] = atmvplentry.Atmvplcasttype
    leafs["atmVplConnKind"] = atmvplentry.Atmvplconnkind
    return leafs
}

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetYangName() string { return "atmVplEntry" }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) SetParent(parent types.Entity) { atmvplentry.parent = parent }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetParent() types.Entity { return atmvplentry.parent }

func (atmvplentry *ATMMIB_Atmvpltable_Atmvplentry) GetParentYangName() string { return "atmVplTable" }

// ATMMIB_Atmvcltable
// The Virtual Channel Link (VCL) table.  A
// bi-directional VCL is modeled as one entry
// in this table. This table can be used for
// PVCs, SVCs and Soft PVCs.
type ATMMIB_Atmvcltable struct {
    parent types.Entity
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

func (atmvcltable *ATMMIB_Atmvcltable) GetFilter() yfilter.YFilter { return atmvcltable.YFilter }

func (atmvcltable *ATMMIB_Atmvcltable) SetFilter(yf yfilter.YFilter) { atmvcltable.YFilter = yf }

func (atmvcltable *ATMMIB_Atmvcltable) GetGoName(yname string) string {
    if yname == "atmVclEntry" { return "Atmvclentry" }
    return ""
}

func (atmvcltable *ATMMIB_Atmvcltable) GetSegmentPath() string {
    return "atmVclTable"
}

func (atmvcltable *ATMMIB_Atmvcltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmVclEntry" {
        for _, c := range atmvcltable.Atmvclentry {
            if atmvcltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atmvcltable_Atmvclentry{}
        atmvcltable.Atmvclentry = append(atmvcltable.Atmvclentry, child)
        return &atmvcltable.Atmvclentry[len(atmvcltable.Atmvclentry)-1]
    }
    return nil
}

func (atmvcltable *ATMMIB_Atmvcltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmvcltable.Atmvclentry {
        children[atmvcltable.Atmvclentry[i].GetSegmentPath()] = &atmvcltable.Atmvclentry[i]
    }
    return children
}

func (atmvcltable *ATMMIB_Atmvcltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmvcltable *ATMMIB_Atmvcltable) GetBundleName() string { return "cisco_ios_xe" }

func (atmvcltable *ATMMIB_Atmvcltable) GetYangName() string { return "atmVclTable" }

func (atmvcltable *ATMMIB_Atmvcltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvcltable *ATMMIB_Atmvcltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvcltable *ATMMIB_Atmvcltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvcltable *ATMMIB_Atmvcltable) SetParent(parent types.Entity) { atmvcltable.parent = parent }

func (atmvcltable *ATMMIB_Atmvcltable) GetParent() types.Entity { return atmvcltable.parent }

func (atmvcltable *ATMMIB_Atmvcltable) GetParentYangName() string { return "ATM-MIB" }

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
    parent types.Entity
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

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetFilter() yfilter.YFilter { return atmvclentry.YFilter }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) SetFilter(yf yfilter.YFilter) { atmvclentry.YFilter = yf }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "atmVclAdminStatus" { return "Atmvcladminstatus" }
    if yname == "atmVclOperStatus" { return "Atmvcloperstatus" }
    if yname == "atmVclLastChange" { return "Atmvcllastchange" }
    if yname == "atmVclReceiveTrafficDescrIndex" { return "Atmvclreceivetrafficdescrindex" }
    if yname == "atmVclTransmitTrafficDescrIndex" { return "Atmvcltransmittrafficdescrindex" }
    if yname == "atmVccAalType" { return "Atmvccaaltype" }
    if yname == "atmVccAal5CpcsTransmitSduSize" { return "Atmvccaal5Cpcstransmitsdusize" }
    if yname == "atmVccAal5CpcsReceiveSduSize" { return "Atmvccaal5Cpcsreceivesdusize" }
    if yname == "atmVccAal5EncapsType" { return "Atmvccaal5Encapstype" }
    if yname == "atmVclCrossConnectIdentifier" { return "Atmvclcrossconnectidentifier" }
    if yname == "atmVclRowStatus" { return "Atmvclrowstatus" }
    if yname == "atmVclCastType" { return "Atmvclcasttype" }
    if yname == "atmVclConnKind" { return "Atmvclconnkind" }
    if yname == "catmxVclOamLoopbackFreq" { return "Catmxvcloamloopbackfreq" }
    if yname == "catmxVclOamRetryFreq" { return "Catmxvcloamretryfreq" }
    if yname == "catmxVclOamUpRetryCount" { return "Catmxvcloamupretrycount" }
    if yname == "catmxVclOamDownRetryCount" { return "Catmxvcloamdownretrycount" }
    if yname == "catmxVclOamEndCCActCount" { return "Catmxvcloamendccactcount" }
    if yname == "catmxVclOamEndCCDeActCount" { return "Catmxvcloamendccdeactcount" }
    if yname == "catmxVclOamEndCCRetryFreq" { return "Catmxvcloamendccretryfreq" }
    if yname == "catmxVclOamSegCCActCount" { return "Catmxvcloamsegccactcount" }
    if yname == "catmxVclOamSegCCDeActCount" { return "Catmxvcloamsegccdeactcount" }
    if yname == "catmxVclOamSegCCRetryFreq" { return "Catmxvcloamsegccretryfreq" }
    if yname == "catmxVclOamManage" { return "Catmxvcloammanage" }
    if yname == "catmxVclOamLoopBkStatus" { return "Catmxvcloamloopbkstatus" }
    if yname == "catmxVclOamVcState" { return "Catmxvcloamvcstate" }
    if yname == "catmxVclOamEndCCStatus" { return "Catmxvcloamendccstatus" }
    if yname == "catmxVclOamSegCCStatus" { return "Catmxvcloamsegccstatus" }
    if yname == "catmxVclOamEndCCVcState" { return "Catmxvcloamendccvcstate" }
    if yname == "catmxVclOamSegCCVcState" { return "Catmxvcloamsegccvcstate" }
    if yname == "catmxVclOamCellsReceived" { return "Catmxvcloamcellsreceived" }
    if yname == "catmxVclOamCellsSent" { return "Catmxvcloamcellssent" }
    if yname == "catmxVclOamCellsDropped" { return "Catmxvcloamcellsdropped" }
    if yname == "catmxVclOamInF5ais" { return "Catmxvcloaminf5Ais" }
    if yname == "catmxVclOamOutF5ais" { return "Catmxvcloamoutf5Ais" }
    if yname == "catmxVclOamInF5rdi" { return "Catmxvcloaminf5Rdi" }
    if yname == "catmxVclOamOutF5rdi" { return "Catmxvcloamoutf5Rdi" }
    return ""
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetSegmentPath() string {
    return "atmVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmvclentry.Atmvclvci) + "']"
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atmvclentry.Ifindex
    leafs["atmVclVpi"] = atmvclentry.Atmvclvpi
    leafs["atmVclVci"] = atmvclentry.Atmvclvci
    leafs["atmVclAdminStatus"] = atmvclentry.Atmvcladminstatus
    leafs["atmVclOperStatus"] = atmvclentry.Atmvcloperstatus
    leafs["atmVclLastChange"] = atmvclentry.Atmvcllastchange
    leafs["atmVclReceiveTrafficDescrIndex"] = atmvclentry.Atmvclreceivetrafficdescrindex
    leafs["atmVclTransmitTrafficDescrIndex"] = atmvclentry.Atmvcltransmittrafficdescrindex
    leafs["atmVccAalType"] = atmvclentry.Atmvccaaltype
    leafs["atmVccAal5CpcsTransmitSduSize"] = atmvclentry.Atmvccaal5Cpcstransmitsdusize
    leafs["atmVccAal5CpcsReceiveSduSize"] = atmvclentry.Atmvccaal5Cpcsreceivesdusize
    leafs["atmVccAal5EncapsType"] = atmvclentry.Atmvccaal5Encapstype
    leafs["atmVclCrossConnectIdentifier"] = atmvclentry.Atmvclcrossconnectidentifier
    leafs["atmVclRowStatus"] = atmvclentry.Atmvclrowstatus
    leafs["atmVclCastType"] = atmvclentry.Atmvclcasttype
    leafs["atmVclConnKind"] = atmvclentry.Atmvclconnkind
    leafs["catmxVclOamLoopbackFreq"] = atmvclentry.Catmxvcloamloopbackfreq
    leafs["catmxVclOamRetryFreq"] = atmvclentry.Catmxvcloamretryfreq
    leafs["catmxVclOamUpRetryCount"] = atmvclentry.Catmxvcloamupretrycount
    leafs["catmxVclOamDownRetryCount"] = atmvclentry.Catmxvcloamdownretrycount
    leafs["catmxVclOamEndCCActCount"] = atmvclentry.Catmxvcloamendccactcount
    leafs["catmxVclOamEndCCDeActCount"] = atmvclentry.Catmxvcloamendccdeactcount
    leafs["catmxVclOamEndCCRetryFreq"] = atmvclentry.Catmxvcloamendccretryfreq
    leafs["catmxVclOamSegCCActCount"] = atmvclentry.Catmxvcloamsegccactcount
    leafs["catmxVclOamSegCCDeActCount"] = atmvclentry.Catmxvcloamsegccdeactcount
    leafs["catmxVclOamSegCCRetryFreq"] = atmvclentry.Catmxvcloamsegccretryfreq
    leafs["catmxVclOamManage"] = atmvclentry.Catmxvcloammanage
    leafs["catmxVclOamLoopBkStatus"] = atmvclentry.Catmxvcloamloopbkstatus
    leafs["catmxVclOamVcState"] = atmvclentry.Catmxvcloamvcstate
    leafs["catmxVclOamEndCCStatus"] = atmvclentry.Catmxvcloamendccstatus
    leafs["catmxVclOamSegCCStatus"] = atmvclentry.Catmxvcloamsegccstatus
    leafs["catmxVclOamEndCCVcState"] = atmvclentry.Catmxvcloamendccvcstate
    leafs["catmxVclOamSegCCVcState"] = atmvclentry.Catmxvcloamsegccvcstate
    leafs["catmxVclOamCellsReceived"] = atmvclentry.Catmxvcloamcellsreceived
    leafs["catmxVclOamCellsSent"] = atmvclentry.Catmxvcloamcellssent
    leafs["catmxVclOamCellsDropped"] = atmvclentry.Catmxvcloamcellsdropped
    leafs["catmxVclOamInF5ais"] = atmvclentry.Catmxvcloaminf5Ais
    leafs["catmxVclOamOutF5ais"] = atmvclentry.Catmxvcloamoutf5Ais
    leafs["catmxVclOamInF5rdi"] = atmvclentry.Catmxvcloaminf5Rdi
    leafs["catmxVclOamOutF5rdi"] = atmvclentry.Catmxvcloamoutf5Rdi
    return leafs
}

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetYangName() string { return "atmVclEntry" }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) SetParent(parent types.Entity) { atmvclentry.parent = parent }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetParent() types.Entity { return atmvclentry.parent }

func (atmvclentry *ATMMIB_Atmvcltable_Atmvclentry) GetParentYangName() string { return "atmVclTable" }

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
    parent types.Entity
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

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetFilter() yfilter.YFilter { return atmvpcrossconnecttable.YFilter }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) SetFilter(yf yfilter.YFilter) { atmvpcrossconnecttable.YFilter = yf }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetGoName(yname string) string {
    if yname == "atmVpCrossConnectEntry" { return "Atmvpcrossconnectentry" }
    return ""
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetSegmentPath() string {
    return "atmVpCrossConnectTable"
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmVpCrossConnectEntry" {
        for _, c := range atmvpcrossconnecttable.Atmvpcrossconnectentry {
            if atmvpcrossconnecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry{}
        atmvpcrossconnecttable.Atmvpcrossconnectentry = append(atmvpcrossconnecttable.Atmvpcrossconnectentry, child)
        return &atmvpcrossconnecttable.Atmvpcrossconnectentry[len(atmvpcrossconnecttable.Atmvpcrossconnectentry)-1]
    }
    return nil
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmvpcrossconnecttable.Atmvpcrossconnectentry {
        children[atmvpcrossconnecttable.Atmvpcrossconnectentry[i].GetSegmentPath()] = &atmvpcrossconnecttable.Atmvpcrossconnectentry[i]
    }
    return children
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetBundleName() string { return "cisco_ios_xe" }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetYangName() string { return "atmVpCrossConnectTable" }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) SetParent(parent types.Entity) { atmvpcrossconnecttable.parent = parent }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetParent() types.Entity { return atmvpcrossconnecttable.parent }

func (atmvpcrossconnecttable *ATMMIB_Atmvpcrossconnecttable) GetParentYangName() string { return "ATM-MIB" }

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
    parent types.Entity
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

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetFilter() yfilter.YFilter { return atmvpcrossconnectentry.YFilter }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) SetFilter(yf yfilter.YFilter) { atmvpcrossconnectentry.YFilter = yf }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetGoName(yname string) string {
    if yname == "atmVpCrossConnectIndex" { return "Atmvpcrossconnectindex" }
    if yname == "atmVpCrossConnectLowIfIndex" { return "Atmvpcrossconnectlowifindex" }
    if yname == "atmVpCrossConnectLowVpi" { return "Atmvpcrossconnectlowvpi" }
    if yname == "atmVpCrossConnectHighIfIndex" { return "Atmvpcrossconnecthighifindex" }
    if yname == "atmVpCrossConnectHighVpi" { return "Atmvpcrossconnecthighvpi" }
    if yname == "atmVpCrossConnectAdminStatus" { return "Atmvpcrossconnectadminstatus" }
    if yname == "atmVpCrossConnectL2HOperStatus" { return "Atmvpcrossconnectl2Hoperstatus" }
    if yname == "atmVpCrossConnectH2LOperStatus" { return "Atmvpcrossconnecth2Loperstatus" }
    if yname == "atmVpCrossConnectL2HLastChange" { return "Atmvpcrossconnectl2Hlastchange" }
    if yname == "atmVpCrossConnectH2LLastChange" { return "Atmvpcrossconnecth2Llastchange" }
    if yname == "atmVpCrossConnectRowStatus" { return "Atmvpcrossconnectrowstatus" }
    return ""
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetSegmentPath() string {
    return "atmVpCrossConnectEntry" + "[atmVpCrossConnectIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectindex) + "']" + "[atmVpCrossConnectLowIfIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectlowifindex) + "']" + "[atmVpCrossConnectLowVpi='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnectlowvpi) + "']" + "[atmVpCrossConnectHighIfIndex='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnecthighifindex) + "']" + "[atmVpCrossConnectHighVpi='" + fmt.Sprintf("%v", atmvpcrossconnectentry.Atmvpcrossconnecthighvpi) + "']"
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atmVpCrossConnectIndex"] = atmvpcrossconnectentry.Atmvpcrossconnectindex
    leafs["atmVpCrossConnectLowIfIndex"] = atmvpcrossconnectentry.Atmvpcrossconnectlowifindex
    leafs["atmVpCrossConnectLowVpi"] = atmvpcrossconnectentry.Atmvpcrossconnectlowvpi
    leafs["atmVpCrossConnectHighIfIndex"] = atmvpcrossconnectentry.Atmvpcrossconnecthighifindex
    leafs["atmVpCrossConnectHighVpi"] = atmvpcrossconnectentry.Atmvpcrossconnecthighvpi
    leafs["atmVpCrossConnectAdminStatus"] = atmvpcrossconnectentry.Atmvpcrossconnectadminstatus
    leafs["atmVpCrossConnectL2HOperStatus"] = atmvpcrossconnectentry.Atmvpcrossconnectl2Hoperstatus
    leafs["atmVpCrossConnectH2LOperStatus"] = atmvpcrossconnectentry.Atmvpcrossconnecth2Loperstatus
    leafs["atmVpCrossConnectL2HLastChange"] = atmvpcrossconnectentry.Atmvpcrossconnectl2Hlastchange
    leafs["atmVpCrossConnectH2LLastChange"] = atmvpcrossconnectentry.Atmvpcrossconnecth2Llastchange
    leafs["atmVpCrossConnectRowStatus"] = atmvpcrossconnectentry.Atmvpcrossconnectrowstatus
    return leafs
}

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetYangName() string { return "atmVpCrossConnectEntry" }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) SetParent(parent types.Entity) { atmvpcrossconnectentry.parent = parent }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetParent() types.Entity { return atmvpcrossconnectentry.parent }

func (atmvpcrossconnectentry *ATMMIB_Atmvpcrossconnecttable_Atmvpcrossconnectentry) GetParentYangName() string { return "atmVpCrossConnectTable" }

// ATMMIB_Atmvccrossconnecttable
// The ATM VC Cross Connect table for PVCs.
// An entry in this table models two
// cross-connected VCLs.
// Each VCL must have its atmConnKind set
// to pvc(1).
type ATMMIB_Atmvccrossconnecttable struct {
    parent types.Entity
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

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetFilter() yfilter.YFilter { return atmvccrossconnecttable.YFilter }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) SetFilter(yf yfilter.YFilter) { atmvccrossconnecttable.YFilter = yf }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetGoName(yname string) string {
    if yname == "atmVcCrossConnectEntry" { return "Atmvccrossconnectentry" }
    return ""
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetSegmentPath() string {
    return "atmVcCrossConnectTable"
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmVcCrossConnectEntry" {
        for _, c := range atmvccrossconnecttable.Atmvccrossconnectentry {
            if atmvccrossconnecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry{}
        atmvccrossconnecttable.Atmvccrossconnectentry = append(atmvccrossconnecttable.Atmvccrossconnectentry, child)
        return &atmvccrossconnecttable.Atmvccrossconnectentry[len(atmvccrossconnecttable.Atmvccrossconnectentry)-1]
    }
    return nil
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmvccrossconnecttable.Atmvccrossconnectentry {
        children[atmvccrossconnecttable.Atmvccrossconnectentry[i].GetSegmentPath()] = &atmvccrossconnecttable.Atmvccrossconnectentry[i]
    }
    return children
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetBundleName() string { return "cisco_ios_xe" }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetYangName() string { return "atmVcCrossConnectTable" }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) SetParent(parent types.Entity) { atmvccrossconnecttable.parent = parent }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetParent() types.Entity { return atmvccrossconnecttable.parent }

func (atmvccrossconnecttable *ATMMIB_Atmvccrossconnecttable) GetParentYangName() string { return "ATM-MIB" }

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
    parent types.Entity
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

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetFilter() yfilter.YFilter { return atmvccrossconnectentry.YFilter }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) SetFilter(yf yfilter.YFilter) { atmvccrossconnectentry.YFilter = yf }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetGoName(yname string) string {
    if yname == "atmVcCrossConnectIndex" { return "Atmvccrossconnectindex" }
    if yname == "atmVcCrossConnectLowIfIndex" { return "Atmvccrossconnectlowifindex" }
    if yname == "atmVcCrossConnectLowVpi" { return "Atmvccrossconnectlowvpi" }
    if yname == "atmVcCrossConnectLowVci" { return "Atmvccrossconnectlowvci" }
    if yname == "atmVcCrossConnectHighIfIndex" { return "Atmvccrossconnecthighifindex" }
    if yname == "atmVcCrossConnectHighVpi" { return "Atmvccrossconnecthighvpi" }
    if yname == "atmVcCrossConnectHighVci" { return "Atmvccrossconnecthighvci" }
    if yname == "atmVcCrossConnectAdminStatus" { return "Atmvccrossconnectadminstatus" }
    if yname == "atmVcCrossConnectL2HOperStatus" { return "Atmvccrossconnectl2Hoperstatus" }
    if yname == "atmVcCrossConnectH2LOperStatus" { return "Atmvccrossconnecth2Loperstatus" }
    if yname == "atmVcCrossConnectL2HLastChange" { return "Atmvccrossconnectl2Hlastchange" }
    if yname == "atmVcCrossConnectH2LLastChange" { return "Atmvccrossconnecth2Llastchange" }
    if yname == "atmVcCrossConnectRowStatus" { return "Atmvccrossconnectrowstatus" }
    return ""
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetSegmentPath() string {
    return "atmVcCrossConnectEntry" + "[atmVcCrossConnectIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectindex) + "']" + "[atmVcCrossConnectLowIfIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowifindex) + "']" + "[atmVcCrossConnectLowVpi='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowvpi) + "']" + "[atmVcCrossConnectLowVci='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnectlowvci) + "']" + "[atmVcCrossConnectHighIfIndex='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighifindex) + "']" + "[atmVcCrossConnectHighVpi='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighvpi) + "']" + "[atmVcCrossConnectHighVci='" + fmt.Sprintf("%v", atmvccrossconnectentry.Atmvccrossconnecthighvci) + "']"
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["atmVcCrossConnectIndex"] = atmvccrossconnectentry.Atmvccrossconnectindex
    leafs["atmVcCrossConnectLowIfIndex"] = atmvccrossconnectentry.Atmvccrossconnectlowifindex
    leafs["atmVcCrossConnectLowVpi"] = atmvccrossconnectentry.Atmvccrossconnectlowvpi
    leafs["atmVcCrossConnectLowVci"] = atmvccrossconnectentry.Atmvccrossconnectlowvci
    leafs["atmVcCrossConnectHighIfIndex"] = atmvccrossconnectentry.Atmvccrossconnecthighifindex
    leafs["atmVcCrossConnectHighVpi"] = atmvccrossconnectentry.Atmvccrossconnecthighvpi
    leafs["atmVcCrossConnectHighVci"] = atmvccrossconnectentry.Atmvccrossconnecthighvci
    leafs["atmVcCrossConnectAdminStatus"] = atmvccrossconnectentry.Atmvccrossconnectadminstatus
    leafs["atmVcCrossConnectL2HOperStatus"] = atmvccrossconnectentry.Atmvccrossconnectl2Hoperstatus
    leafs["atmVcCrossConnectH2LOperStatus"] = atmvccrossconnectentry.Atmvccrossconnecth2Loperstatus
    leafs["atmVcCrossConnectL2HLastChange"] = atmvccrossconnectentry.Atmvccrossconnectl2Hlastchange
    leafs["atmVcCrossConnectH2LLastChange"] = atmvccrossconnectentry.Atmvccrossconnecth2Llastchange
    leafs["atmVcCrossConnectRowStatus"] = atmvccrossconnectentry.Atmvccrossconnectrowstatus
    return leafs
}

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetYangName() string { return "atmVcCrossConnectEntry" }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) SetParent(parent types.Entity) { atmvccrossconnectentry.parent = parent }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetParent() types.Entity { return atmvccrossconnectentry.parent }

func (atmvccrossconnectentry *ATMMIB_Atmvccrossconnecttable_Atmvccrossconnectentry) GetParentYangName() string { return "atmVcCrossConnectTable" }

// ATMMIB_Aal5Vcctable
// This table contains AAL5 VCC performance
// parameters.
type ATMMIB_Aal5Vcctable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains the AAL5 VCC performance parameters and is indexed by
    // ifIndex values of AAL5 interfaces and the associated VPI/VCI values. The
    // type is slice of ATMMIB_Aal5Vcctable_Aal5Vccentry.
    Aal5Vccentry []ATMMIB_Aal5Vcctable_Aal5Vccentry
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetFilter() yfilter.YFilter { return aal5Vcctable.YFilter }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) SetFilter(yf yfilter.YFilter) { aal5Vcctable.YFilter = yf }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetGoName(yname string) string {
    if yname == "aal5VccEntry" { return "Aal5Vccentry" }
    return ""
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetSegmentPath() string {
    return "aal5VccTable"
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aal5VccEntry" {
        for _, c := range aal5Vcctable.Aal5Vccentry {
            if aal5Vcctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ATMMIB_Aal5Vcctable_Aal5Vccentry{}
        aal5Vcctable.Aal5Vccentry = append(aal5Vcctable.Aal5Vccentry, child)
        return &aal5Vcctable.Aal5Vccentry[len(aal5Vcctable.Aal5Vccentry)-1]
    }
    return nil
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aal5Vcctable.Aal5Vccentry {
        children[aal5Vcctable.Aal5Vccentry[i].GetSegmentPath()] = &aal5Vcctable.Aal5Vccentry[i]
    }
    return children
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetBundleName() string { return "cisco_ios_xe" }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetYangName() string { return "aal5VccTable" }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) SetParent(parent types.Entity) { aal5Vcctable.parent = parent }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetParent() types.Entity { return aal5Vcctable.parent }

func (aal5Vcctable *ATMMIB_Aal5Vcctable) GetParentYangName() string { return "ATM-MIB" }

// ATMMIB_Aal5Vcctable_Aal5Vccentry
// This list contains the AAL5 VCC
// performance parameters and is indexed
// by ifIndex values of AAL5 interfaces
// and the associated VPI/VCI values.
type ATMMIB_Aal5Vcctable_Aal5Vccentry struct {
    parent types.Entity
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

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetFilter() yfilter.YFilter { return aal5Vccentry.YFilter }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) SetFilter(yf yfilter.YFilter) { aal5Vccentry.YFilter = yf }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "aal5VccVpi" { return "Aal5Vccvpi" }
    if yname == "aal5VccVci" { return "Aal5Vccvci" }
    if yname == "aal5VccCrcErrors" { return "Aal5Vcccrcerrors" }
    if yname == "aal5VccSarTimeOuts" { return "Aal5Vccsartimeouts" }
    if yname == "aal5VccOverSizedSDUs" { return "Aal5Vccoversizedsdus" }
    if yname == "cAal5VccInPkts" { return "Caal5Vccinpkts" }
    if yname == "cAal5VccOutPkts" { return "Caal5Vccoutpkts" }
    if yname == "cAal5VccInOctets" { return "Caal5Vccinoctets" }
    if yname == "cAal5VccOutOctets" { return "Caal5Vccoutoctets" }
    if yname == "cAal5VccInDroppedPkts" { return "Caal5Vccindroppedpkts" }
    if yname == "cAal5VccOutDroppedPkts" { return "Caal5Vccoutdroppedpkts" }
    if yname == "cAal5VccInDroppedOctets" { return "Caal5Vccindroppedoctets" }
    if yname == "cAal5VccOutDroppedOctets" { return "Caal5Vccoutdroppedoctets" }
    if yname == "cAal5VccHCInPkts" { return "Caal5Vcchcinpkts" }
    if yname == "cAal5VccHCOutPkts" { return "Caal5Vcchcoutpkts" }
    if yname == "cAal5VccHCInOctets" { return "Caal5Vcchcinoctets" }
    if yname == "cAal5VccHCOutOctets" { return "Caal5Vcchcoutoctets" }
    if yname == "cAal5VccExtCompEnabled" { return "Caal5Vccextcompenabled" }
    if yname == "cAal5VccExtVoice" { return "Caal5Vccextvoice" }
    if yname == "cAal5VccExtInF5OamCells" { return "Caal5Vccextinf5Oamcells" }
    if yname == "cAal5VccExtOutF5OamCells" { return "Caal5Vccextoutf5Oamcells" }
    return ""
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetSegmentPath() string {
    return "aal5VccEntry" + "[ifIndex='" + fmt.Sprintf("%v", aal5Vccentry.Ifindex) + "']" + "[aal5VccVpi='" + fmt.Sprintf("%v", aal5Vccentry.Aal5Vccvpi) + "']" + "[aal5VccVci='" + fmt.Sprintf("%v", aal5Vccentry.Aal5Vccvci) + "']"
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = aal5Vccentry.Ifindex
    leafs["aal5VccVpi"] = aal5Vccentry.Aal5Vccvpi
    leafs["aal5VccVci"] = aal5Vccentry.Aal5Vccvci
    leafs["aal5VccCrcErrors"] = aal5Vccentry.Aal5Vcccrcerrors
    leafs["aal5VccSarTimeOuts"] = aal5Vccentry.Aal5Vccsartimeouts
    leafs["aal5VccOverSizedSDUs"] = aal5Vccentry.Aal5Vccoversizedsdus
    leafs["cAal5VccInPkts"] = aal5Vccentry.Caal5Vccinpkts
    leafs["cAal5VccOutPkts"] = aal5Vccentry.Caal5Vccoutpkts
    leafs["cAal5VccInOctets"] = aal5Vccentry.Caal5Vccinoctets
    leafs["cAal5VccOutOctets"] = aal5Vccentry.Caal5Vccoutoctets
    leafs["cAal5VccInDroppedPkts"] = aal5Vccentry.Caal5Vccindroppedpkts
    leafs["cAal5VccOutDroppedPkts"] = aal5Vccentry.Caal5Vccoutdroppedpkts
    leafs["cAal5VccInDroppedOctets"] = aal5Vccentry.Caal5Vccindroppedoctets
    leafs["cAal5VccOutDroppedOctets"] = aal5Vccentry.Caal5Vccoutdroppedoctets
    leafs["cAal5VccHCInPkts"] = aal5Vccentry.Caal5Vcchcinpkts
    leafs["cAal5VccHCOutPkts"] = aal5Vccentry.Caal5Vcchcoutpkts
    leafs["cAal5VccHCInOctets"] = aal5Vccentry.Caal5Vcchcinoctets
    leafs["cAal5VccHCOutOctets"] = aal5Vccentry.Caal5Vcchcoutoctets
    leafs["cAal5VccExtCompEnabled"] = aal5Vccentry.Caal5Vccextcompenabled
    leafs["cAal5VccExtVoice"] = aal5Vccentry.Caal5Vccextvoice
    leafs["cAal5VccExtInF5OamCells"] = aal5Vccentry.Caal5Vccextinf5Oamcells
    leafs["cAal5VccExtOutF5OamCells"] = aal5Vccentry.Caal5Vccextoutf5Oamcells
    return leafs
}

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetBundleName() string { return "cisco_ios_xe" }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetYangName() string { return "aal5VccEntry" }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) SetParent(parent types.Entity) { aal5Vccentry.parent = parent }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetParent() types.Entity { return aal5Vccentry.parent }

func (aal5Vccentry *ATMMIB_Aal5Vcctable_Aal5Vccentry) GetParentYangName() string { return "aal5VccTable" }

