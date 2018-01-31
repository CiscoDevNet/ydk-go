// The MIB module for managing Cisco extensions to
// the 802.1D Spanning Tree Protocol (STP).
package cisco_stp_extensions_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_stp_extensions_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-STP-EXTENSIONS-MIB CISCO-STP-EXTENSIONS-MIB}", reflect.TypeOf(CISCOSTPEXTENSIONSMIB{}))
    ydk.RegisterEntity("CISCO-STP-EXTENSIONS-MIB:CISCO-STP-EXTENSIONS-MIB", reflect.TypeOf(CISCOSTPEXTENSIONSMIB{}))
}

// CISCOSTPEXTENSIONSMIB
type CISCOSTPEXTENSIONSMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Stpxuplinkfastobjects CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects

    
    Stpxbackbonefastobjects CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects

    
    Stpxspanningtreeobjects CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects

    
    Stpxmistpobjects CISCOSTPEXTENSIONSMIB_Stpxmistpobjects

    
    Stpxloopguardobjects CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects

    
    Stpxfaststartobjects CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects

    
    Stpxbpduskewingobjects CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects

    
    Stpxmstobjects CISCOSTPEXTENSIONSMIB_Stpxmstobjects

    
    Stpxrstpobjects CISCOSTPEXTENSIONSMIB_Stpxrstpobjects

    
    Stpxsmstobjects CISCOSTPEXTENSIONSMIB_Stpxsmstobjects

    // A list of Virtual LAN entries containing information for Spanning Tree
    // PVST+ protocol.  An entry will exist for each VLAN existing on  the device.
    Stpxpvstvlantable CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable

    // A table containing a list of the ports for which a particular VLAN's
    // Spanning Tree has been found to have an inconsistency.  Two types of
    // inconsistency are discovered: 1) an inconsistency where two different port
    // types have been plugged together; and 2) an inconsistency where different
    // switches have different PVIDs for the same link.
    Stpxinconsistencytable CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable

    // A table containing a list of the bridge ports for which Spanning Tree
    // RootGuard capability can be configured.
    Stpxrootguardconfigtable CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable

    // A table containing a list of the bridge ports for which a particular
    // Spanning Tree instance has been found  to have an root-inconsistency. The
    // agent creates a new  entry in this table whenever it detects a new 
    // root-inconsistency, and deletes entries  when/soon after the inconsistency
    // is no longer present.
    Stpxrootinconsistencytable CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable

    // This table contains one entry for each instance of MISTP and  it contains
    // stpxMISTPInstanceNumber entries, numbered from 1 to
    // stpxMISTPInstanceNumber.  This table is only instantiated when the value of
    // stpxSpanningTreeType is mistp(2) or mistpPvstPlus(3).
    Stpxmistpinstancetable CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable

    // A table containing a list of the bridge ports for which Spanning Tree
    // LoopGuard capability can be configured.
    Stpxloopguardconfigtable CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable

    // A table containing a list of the bridge ports for which a particular
    // Spanning Tree instance has been found to have a loop-inconsistency. The
    // agent creates a new entry in this table whenever it detects a new
    // loop-inconsistency, and deletes entries when/soon after the inconsistency
    // is no longer present.
    Stpxloopinconsistencytable CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable

    // A table containing a list of the bridge ports for which Spanning Tree Port
    // Fast Start can be configured.
    Stpxfaststartporttable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable

    // A table containing a list of the bridge ports  for a particular Spanning
    // Tree Instance.
    Stpxfaststartopermodetable CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable

    // A table containing a list of the bridge ports for  which a particular
    // Spanning Tree instance has been  detected to have BPDU skewing occurred
    // since the  object value of stpxBpduSkewingDetectionEnable was last changed
    // to true(1).  The agent creates a new entry in this table whenever a port in
    // a particular Spanning Tree instance is  detected to be BPDU skewed since
    // the object value of  stpxBpduSkewingDetectionEnable object is changed to 
    // true(1). The agent deletes all the entries in this  table when the object
    // value of  stpxBpduSkewingDetectionEnable is changed to false(2) or the
    // object value of stpxSpanningTreeType is  changed.
    Stpxbpduskewingtable CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable

    // This table contains MST instance information with one entry for an MST
    // instance within the range of  0 to the object value of
    // stpxMSTMaxInstanceNumber.   This table is deprecated and replaced by 
    // stpxSMSTInstanceTable.
    Stpxmstinstancetable CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable

    // This table contains MST instance information in the  Edit Buffer with one
    // entry for each MST instance numbered from 0 to stpxMSTMaxInstanceNumber.  
    // This table is only instantiated when the  stpxMSTRegionEditBufferStatus has
    // the value of acquiredBySnmp(2).  This table is deprecated and replaced by 
    // stpxSMSTInstanceEditTable.
    Stpxmstinstanceedittable CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable

    // A table containing port information for the MST  Protocol on all the bridge
    // ports existing on the  system.
    Stpxmstporttable CISCOSTPEXTENSIONSMIB_Stpxmstporttable

    // A table containing a list of the bridge ports for a  particular MST
    // instance.  This table is only instantiated  when the stpxSpanningTreeType
    // is mst(4).   This table is deprecated and replaced with 
    // stpxRSTPPortRoleTable.
    Stpxmstportroletable CISCOSTPEXTENSIONSMIB_Stpxmstportroletable

    // A table containing port information for the RSTP  Protocol on all the
    // bridge ports existing in the  system.
    Stpxrstpporttable CISCOSTPEXTENSIONSMIB_Stpxrstpporttable

    // A table containing a list of the bridge ports for a  particular Spanning
    // Tree instance.  This table is  only instantiated when the
    // stpxSpanningTreeType is mst(4)  or rapidPvstPlus(5).
    Stpxrstpportroletable CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable

    // A table containing a list of the bridge ports  for a particular Spanning
    // Tree Instance. This table is only instantiated when the object value of
    // stpxSpanningTreeType is rapidPvstPlus(5).
    Stpxrpvstporttable CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable

    // This table contains MST instance information for IEEE MST.
    Stpxsmstinstancetable CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable

    // This table contains MST instance information in the  Edit Buffer.   This
    // table is only instantiated when the object value of 
    // stpxMSTRegionEditBufferStatus has the value of acquiredBySnmp(2).
    Stpxsmstinstanceedittable CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable

    // A table containing port information for the MST  Protocol on all the bridge
    // ports existing on the  system.  This table is only instantiated when the
    // object  value of stpxSpanningTreeType is mst(4).
    Stpxsmstporttable CISCOSTPEXTENSIONSMIB_Stpxsmstporttable
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetFilter() yfilter.YFilter { return cISCOSTPEXTENSIONSMIB.YFilter }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) SetFilter(yf yfilter.YFilter) { cISCOSTPEXTENSIONSMIB.YFilter = yf }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetGoName(yname string) string {
    if yname == "stpxUplinkFastObjects" { return "Stpxuplinkfastobjects" }
    if yname == "stpxBackboneFastObjects" { return "Stpxbackbonefastobjects" }
    if yname == "stpxSpanningTreeObjects" { return "Stpxspanningtreeobjects" }
    if yname == "stpxMISTPObjects" { return "Stpxmistpobjects" }
    if yname == "stpxLoopGuardObjects" { return "Stpxloopguardobjects" }
    if yname == "stpxFastStartObjects" { return "Stpxfaststartobjects" }
    if yname == "stpxBpduSkewingObjects" { return "Stpxbpduskewingobjects" }
    if yname == "stpxMSTObjects" { return "Stpxmstobjects" }
    if yname == "stpxRSTPObjects" { return "Stpxrstpobjects" }
    if yname == "stpxSMSTObjects" { return "Stpxsmstobjects" }
    if yname == "stpxPVSTVlanTable" { return "Stpxpvstvlantable" }
    if yname == "stpxInconsistencyTable" { return "Stpxinconsistencytable" }
    if yname == "stpxRootGuardConfigTable" { return "Stpxrootguardconfigtable" }
    if yname == "stpxRootInconsistencyTable" { return "Stpxrootinconsistencytable" }
    if yname == "stpxMISTPInstanceTable" { return "Stpxmistpinstancetable" }
    if yname == "stpxLoopGuardConfigTable" { return "Stpxloopguardconfigtable" }
    if yname == "stpxLoopInconsistencyTable" { return "Stpxloopinconsistencytable" }
    if yname == "stpxFastStartPortTable" { return "Stpxfaststartporttable" }
    if yname == "stpxFastStartOperModeTable" { return "Stpxfaststartopermodetable" }
    if yname == "stpxBpduSkewingTable" { return "Stpxbpduskewingtable" }
    if yname == "stpxMSTInstanceTable" { return "Stpxmstinstancetable" }
    if yname == "stpxMSTInstanceEditTable" { return "Stpxmstinstanceedittable" }
    if yname == "stpxMSTPortTable" { return "Stpxmstporttable" }
    if yname == "stpxMSTPortRoleTable" { return "Stpxmstportroletable" }
    if yname == "stpxRSTPPortTable" { return "Stpxrstpporttable" }
    if yname == "stpxRSTPPortRoleTable" { return "Stpxrstpportroletable" }
    if yname == "stpxRPVSTPortTable" { return "Stpxrpvstporttable" }
    if yname == "stpxSMSTInstanceTable" { return "Stpxsmstinstancetable" }
    if yname == "stpxSMSTInstanceEditTable" { return "Stpxsmstinstanceedittable" }
    if yname == "stpxSMSTPortTable" { return "Stpxsmstporttable" }
    return ""
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetSegmentPath() string {
    return "CISCO-STP-EXTENSIONS-MIB:CISCO-STP-EXTENSIONS-MIB"
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxUplinkFastObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxuplinkfastobjects
    }
    if childYangName == "stpxBackboneFastObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxbackbonefastobjects
    }
    if childYangName == "stpxSpanningTreeObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxspanningtreeobjects
    }
    if childYangName == "stpxMISTPObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmistpobjects
    }
    if childYangName == "stpxLoopGuardObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxloopguardobjects
    }
    if childYangName == "stpxFastStartObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxfaststartobjects
    }
    if childYangName == "stpxBpduSkewingObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingobjects
    }
    if childYangName == "stpxMSTObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmstobjects
    }
    if childYangName == "stpxRSTPObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrstpobjects
    }
    if childYangName == "stpxSMSTObjects" {
        return &cISCOSTPEXTENSIONSMIB.Stpxsmstobjects
    }
    if childYangName == "stpxPVSTVlanTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxpvstvlantable
    }
    if childYangName == "stpxInconsistencyTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxinconsistencytable
    }
    if childYangName == "stpxRootGuardConfigTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrootguardconfigtable
    }
    if childYangName == "stpxRootInconsistencyTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrootinconsistencytable
    }
    if childYangName == "stpxMISTPInstanceTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmistpinstancetable
    }
    if childYangName == "stpxLoopGuardConfigTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxloopguardconfigtable
    }
    if childYangName == "stpxLoopInconsistencyTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxloopinconsistencytable
    }
    if childYangName == "stpxFastStartPortTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxfaststartporttable
    }
    if childYangName == "stpxFastStartOperModeTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxfaststartopermodetable
    }
    if childYangName == "stpxBpduSkewingTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingtable
    }
    if childYangName == "stpxMSTInstanceTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmstinstancetable
    }
    if childYangName == "stpxMSTInstanceEditTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmstinstanceedittable
    }
    if childYangName == "stpxMSTPortTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmstporttable
    }
    if childYangName == "stpxMSTPortRoleTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxmstportroletable
    }
    if childYangName == "stpxRSTPPortTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrstpporttable
    }
    if childYangName == "stpxRSTPPortRoleTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrstpportroletable
    }
    if childYangName == "stpxRPVSTPortTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxrpvstporttable
    }
    if childYangName == "stpxSMSTInstanceTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxsmstinstancetable
    }
    if childYangName == "stpxSMSTInstanceEditTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxsmstinstanceedittable
    }
    if childYangName == "stpxSMSTPortTable" {
        return &cISCOSTPEXTENSIONSMIB.Stpxsmstporttable
    }
    return nil
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stpxUplinkFastObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxuplinkfastobjects
    children["stpxBackboneFastObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxbackbonefastobjects
    children["stpxSpanningTreeObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxspanningtreeobjects
    children["stpxMISTPObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxmistpobjects
    children["stpxLoopGuardObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxloopguardobjects
    children["stpxFastStartObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxfaststartobjects
    children["stpxBpduSkewingObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingobjects
    children["stpxMSTObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxmstobjects
    children["stpxRSTPObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxrstpobjects
    children["stpxSMSTObjects"] = &cISCOSTPEXTENSIONSMIB.Stpxsmstobjects
    children["stpxPVSTVlanTable"] = &cISCOSTPEXTENSIONSMIB.Stpxpvstvlantable
    children["stpxInconsistencyTable"] = &cISCOSTPEXTENSIONSMIB.Stpxinconsistencytable
    children["stpxRootGuardConfigTable"] = &cISCOSTPEXTENSIONSMIB.Stpxrootguardconfigtable
    children["stpxRootInconsistencyTable"] = &cISCOSTPEXTENSIONSMIB.Stpxrootinconsistencytable
    children["stpxMISTPInstanceTable"] = &cISCOSTPEXTENSIONSMIB.Stpxmistpinstancetable
    children["stpxLoopGuardConfigTable"] = &cISCOSTPEXTENSIONSMIB.Stpxloopguardconfigtable
    children["stpxLoopInconsistencyTable"] = &cISCOSTPEXTENSIONSMIB.Stpxloopinconsistencytable
    children["stpxFastStartPortTable"] = &cISCOSTPEXTENSIONSMIB.Stpxfaststartporttable
    children["stpxFastStartOperModeTable"] = &cISCOSTPEXTENSIONSMIB.Stpxfaststartopermodetable
    children["stpxBpduSkewingTable"] = &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingtable
    children["stpxMSTInstanceTable"] = &cISCOSTPEXTENSIONSMIB.Stpxmstinstancetable
    children["stpxMSTInstanceEditTable"] = &cISCOSTPEXTENSIONSMIB.Stpxmstinstanceedittable
    children["stpxMSTPortTable"] = &cISCOSTPEXTENSIONSMIB.Stpxmstporttable
    children["stpxMSTPortRoleTable"] = &cISCOSTPEXTENSIONSMIB.Stpxmstportroletable
    children["stpxRSTPPortTable"] = &cISCOSTPEXTENSIONSMIB.Stpxrstpporttable
    children["stpxRSTPPortRoleTable"] = &cISCOSTPEXTENSIONSMIB.Stpxrstpportroletable
    children["stpxRPVSTPortTable"] = &cISCOSTPEXTENSIONSMIB.Stpxrpvstporttable
    children["stpxSMSTInstanceTable"] = &cISCOSTPEXTENSIONSMIB.Stpxsmstinstancetable
    children["stpxSMSTInstanceEditTable"] = &cISCOSTPEXTENSIONSMIB.Stpxsmstinstanceedittable
    children["stpxSMSTPortTable"] = &cISCOSTPEXTENSIONSMIB.Stpxsmstporttable
    return children
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) SetParent(parent types.Entity) { cISCOSTPEXTENSIONSMIB.parent = parent }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetParent() types.Entity { return cISCOSTPEXTENSIONSMIB.parent }

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects
type CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An indication of whether the UplinkFast capability is administratively
    // enabled on the device.  If the platform does not support configuration of
    // this object when the object value of stpxSpanningTreeType is  mst(4), then
    // this object is not instantiated. The type is bool.
    Stpxuplinkfastenabled interface{}

    // The cumulative number of UplinkFast transitions (from the STP 'Blocking'
    // state directly to the STP 'Forwarding' state).  All transitions are
    // included in this counter, irrespective of the instance of the Spanning Tree
    // Protocol on which they occur.  If the platform supports the
    // stpxUplinkFastOperEnabled  object, then this object is not instantiated
    // when the  object value of stpxUplinkFastOperEnabled is false(2). If the
    // platform does not support the  stpxUplinkFastOperEnabled object, then this
    // object is  not instantiated when the object value of  stpxSpanningTreeType
    // is mst(4). The type is interface{} with range: 0..4294967295. Units are
    // transitions.
    Stpxuplinkfasttransitions interface{}

    // The maximum number of station-learning frames that this device will
    // generate in each 100 milli-second period after a UplinkFast transition.  By
    // configuring this object, the network administrator can limit the rate at
    // which station-learning frames are generated.    If the platform does not
    // support configuration of this object when the object value of
    // stpxSpanningTreeType is mst(4), then this object is not instantiated. The
    // type is interface{} with range: 0..32000. Units are frames.
    Stpxuplinkstationlearninggenrate interface{}

    // The cumulative number of station-learning frames generated due to
    // UplinkFast transitions.  All generated station-learning frames are included
    // in this counter, irrespective of the instance of the Spanning Tree Protocol
    // on which the UplinkFast transition occurred.  If the platform supports the
    // stpxUplinkFastOperEnabled  object, then this object is not instantiated
    // when the  object value of stpxUplinkFastOperEnabled is false(2). If the
    // platform does not support the  stpxUplinkFastOperEnabled object, then this
    // object is  not instantiated when the object value of  stpxSpanningTreeType
    // is mst(4). The type is interface{} with range: 0..4294967295. Units are
    // frames.
    Stpxuplinkstationlearningframes interface{}

    // An indication of whether the UplinkFast capability is  operationally
    // enabled on the device. The type is bool.
    Stpxuplinkfastoperenabled interface{}
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetFilter() yfilter.YFilter { return stpxuplinkfastobjects.YFilter }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) SetFilter(yf yfilter.YFilter) { stpxuplinkfastobjects.YFilter = yf }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetGoName(yname string) string {
    if yname == "stpxUplinkFastEnabled" { return "Stpxuplinkfastenabled" }
    if yname == "stpxUplinkFastTransitions" { return "Stpxuplinkfasttransitions" }
    if yname == "stpxUplinkStationLearningGenRate" { return "Stpxuplinkstationlearninggenrate" }
    if yname == "stpxUplinkStationLearningFrames" { return "Stpxuplinkstationlearningframes" }
    if yname == "stpxUplinkFastOperEnabled" { return "Stpxuplinkfastoperenabled" }
    return ""
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetSegmentPath() string {
    return "stpxUplinkFastObjects"
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxUplinkFastEnabled"] = stpxuplinkfastobjects.Stpxuplinkfastenabled
    leafs["stpxUplinkFastTransitions"] = stpxuplinkfastobjects.Stpxuplinkfasttransitions
    leafs["stpxUplinkStationLearningGenRate"] = stpxuplinkfastobjects.Stpxuplinkstationlearninggenrate
    leafs["stpxUplinkStationLearningFrames"] = stpxuplinkfastobjects.Stpxuplinkstationlearningframes
    leafs["stpxUplinkFastOperEnabled"] = stpxuplinkfastobjects.Stpxuplinkfastoperenabled
    return leafs
}

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetYangName() string { return "stpxUplinkFastObjects" }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) SetParent(parent types.Entity) { stpxuplinkfastobjects.parent = parent }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetParent() types.Entity { return stpxuplinkfastobjects.parent }

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects
type CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An indication of whether the BackboneFast capability is administratively
    // enabled on the device. The type is bool.
    Stpxbackbonefastenabled interface{}

    // The number of inferior BPDUs received by the switch  since the
    // stpxBackboneFastOperEnabled has become true(1). If the value of 
    // stpxBackboneFastOperEnabled is false(2), then this  mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    Stpxbackbonefastininferiorbpdus interface{}

    // The number of Root Link Query request PDUs received by the switch since the
    // stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    Stpxbackbonefastinrlqrequestpdus interface{}

    // The number of Root Link Query response PDUs received by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    Stpxbackbonefastinrlqresponsepdus interface{}

    // The number of Root Link Query request PDUs transmitted by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    Stpxbackbonefastoutrlqrequestpdus interface{}

    // The number of Root Link Query response PDUs transmitted by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    Stpxbackbonefastoutrlqresponsepdus interface{}

    // An indication of whether the BackboneFast capability is operationally
    // enabled on the device. The type is bool.
    Stpxbackbonefastoperenabled interface{}
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetFilter() yfilter.YFilter { return stpxbackbonefastobjects.YFilter }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) SetFilter(yf yfilter.YFilter) { stpxbackbonefastobjects.YFilter = yf }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetGoName(yname string) string {
    if yname == "stpxBackboneFastEnabled" { return "Stpxbackbonefastenabled" }
    if yname == "stpxBackboneFastInInferiorBPDUs" { return "Stpxbackbonefastininferiorbpdus" }
    if yname == "stpxBackboneFastInRLQRequestPDUs" { return "Stpxbackbonefastinrlqrequestpdus" }
    if yname == "stpxBackboneFastInRLQResponsePDUs" { return "Stpxbackbonefastinrlqresponsepdus" }
    if yname == "stpxBackboneFastOutRLQRequestPDUs" { return "Stpxbackbonefastoutrlqrequestpdus" }
    if yname == "stpxBackboneFastOutRLQResponsePDUs" { return "Stpxbackbonefastoutrlqresponsepdus" }
    if yname == "stpxBackboneFastOperEnabled" { return "Stpxbackbonefastoperenabled" }
    return ""
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetSegmentPath() string {
    return "stpxBackboneFastObjects"
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxBackboneFastEnabled"] = stpxbackbonefastobjects.Stpxbackbonefastenabled
    leafs["stpxBackboneFastInInferiorBPDUs"] = stpxbackbonefastobjects.Stpxbackbonefastininferiorbpdus
    leafs["stpxBackboneFastInRLQRequestPDUs"] = stpxbackbonefastobjects.Stpxbackbonefastinrlqrequestpdus
    leafs["stpxBackboneFastInRLQResponsePDUs"] = stpxbackbonefastobjects.Stpxbackbonefastinrlqresponsepdus
    leafs["stpxBackboneFastOutRLQRequestPDUs"] = stpxbackbonefastobjects.Stpxbackbonefastoutrlqrequestpdus
    leafs["stpxBackboneFastOutRLQResponsePDUs"] = stpxbackbonefastobjects.Stpxbackbonefastoutrlqresponsepdus
    leafs["stpxBackboneFastOperEnabled"] = stpxbackbonefastobjects.Stpxbackbonefastoperenabled
    return leafs
}

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetYangName() string { return "stpxBackboneFastObjects" }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) SetParent(parent types.Entity) { stpxbackbonefastobjects.parent = parent }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetParent() types.Entity { return stpxbackbonefastobjects.parent }

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects
type CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The actual mode of spanning tree protocol runs on the  device. It can be
    // one of the following:  pvstPlus -- PVST+ (Per VLAN Spanning Tree+
    // Protocol).  mistp -- MISTP (Multi Instance Spanning Tree Protocol). 
    // mistpPvstPlus --  MISTP with the tunneling scheme                     
    // enabled for PVST+.  mst -- IEEE 802.1s Multiple Spanning Tree (MST)       
    // with IEEE 802.1w Rapid Spanning Tree Protocol        (RSTP).  rapidPvstPlus
    // -- IEEE 802.1w Rapid Spanning Tree          Protocol (RSTP) for all vlans
    // in PVST+.  When the value of this MIB object gets changed, the  network
    // connectivity would be affected and the  connectivity to this device would
    // also be lost  temporarily. The type is Stpxspanningtreetype.
    Stpxspanningtreetype interface{}

    // Indicates the administrative  spanning tree path cost mode  configured on
    // device. The type is Stpxspanningtreepathcostmode.
    Stpxspanningtreepathcostmode interface{}

    // Indicates whether Extended System ID feature  is administratively enabled
    // on the device or not. The type is bool.
    Stpxextendedsysidadminenabled interface{}

    // Indicates whether Extended System ID feature  is operationaly enabled on
    // the device or not.  If the value of this object is true(1), then the
    // accepted values for dot1dStpPriority in BRIDGE-MIB should be multiples of
    // 4096 plus bridge instance ID, such as VlanIndex. Changing this object value
    // might cause the values of dot1dBaseBridgeAddress and dot1dStpPriority in
    // BRIDGE-MIB to be changed also. The type is bool.
    Stpxextendedsysidoperenabled interface{}

    // Indicates whether a specified notification is enabled or not. If a bit
    // corresponding to a notification is set to 1, then  the specified
    // notification can be generated.  newRoot -- the newRoot notification as
    // defined in BRIDGE-MIB.  topologyChange -- the topologyChange notification
    // as                   defined in BRIDGE-MIB.  inconsistency -- the
    // stpxInconsistencyUpdate notification.  rootInconsistency -- the
    // stpxRootInconsistencyUpdate                       notification. 
    // loopInconsistency -- the stpxLoopInconsistencyUpdate                      
    // notification. The type is map[string]bool.
    Stpxnotificationenable interface{}

    // Indicate the operational spanning tree path cost mode on device. This mode
    // applies to all instances of the Spanning Tree protocol running on the
    // device.   When the value of this MIB object gets changed, the path cost of
    // all ports will be reassigned to the default path cost values based on the
    // new spanning tree path cost mode and the ports' speed.  When the value of
    // this MIB object is long(2), the stpxLongStpPortPathCost MIB object must be
    // used in order to retrieve/configure the spanning tree port path cost as a
    // 32 bits value. The set operation on dot1dStpPortPathCost in BRIDGE-MIB will
    // be rejected. While retrieving the value of dot1dStpPortPathCost, the
    // maximum value of 65535 will be returned if the value of
    // stpxLongStpPortPathCost for the same instance exceeds 65535.  When the
    // value of this MIB object is short(1), the dot1dStpPortPathCost in
    // BRIDGE-MIB must be used. The type is Stpxspanningtreepathcostopermode.
    Stpxspanningtreepathcostopermode interface{}
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetFilter() yfilter.YFilter { return stpxspanningtreeobjects.YFilter }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) SetFilter(yf yfilter.YFilter) { stpxspanningtreeobjects.YFilter = yf }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetGoName(yname string) string {
    if yname == "stpxSpanningTreeType" { return "Stpxspanningtreetype" }
    if yname == "stpxSpanningTreePathCostMode" { return "Stpxspanningtreepathcostmode" }
    if yname == "stpxExtendedSysIDAdminEnabled" { return "Stpxextendedsysidadminenabled" }
    if yname == "stpxExtendedSysIDOperEnabled" { return "Stpxextendedsysidoperenabled" }
    if yname == "stpxNotificationEnable" { return "Stpxnotificationenable" }
    if yname == "stpxSpanningTreePathCostOperMode" { return "Stpxspanningtreepathcostopermode" }
    return ""
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetSegmentPath() string {
    return "stpxSpanningTreeObjects"
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxSpanningTreeType"] = stpxspanningtreeobjects.Stpxspanningtreetype
    leafs["stpxSpanningTreePathCostMode"] = stpxspanningtreeobjects.Stpxspanningtreepathcostmode
    leafs["stpxExtendedSysIDAdminEnabled"] = stpxspanningtreeobjects.Stpxextendedsysidadminenabled
    leafs["stpxExtendedSysIDOperEnabled"] = stpxspanningtreeobjects.Stpxextendedsysidoperenabled
    leafs["stpxNotificationEnable"] = stpxspanningtreeobjects.Stpxnotificationenable
    leafs["stpxSpanningTreePathCostOperMode"] = stpxspanningtreeobjects.Stpxspanningtreepathcostopermode
    return leafs
}

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetYangName() string { return "stpxSpanningTreeObjects" }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) SetParent(parent types.Entity) { stpxspanningtreeobjects.parent = parent }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetParent() types.Entity { return stpxspanningtreeobjects.parent }

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode represents configured on device.
type CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode_short CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode = "short"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode_long CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostmode = "long"
)

// CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode represents the dot1dStpPortPathCost in BRIDGE-MIB must be used.
type CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode_short CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode = "short"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode_long CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreepathcostopermode = "long"
)

// CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype represents temporarily.
type CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype string

const (
    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype_pvstPlus CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype = "pvstPlus"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype_mistp CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype = "mistp"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype_mistpPvstPlus CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype = "mistpPvstPlus"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype_mst CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype = "mst"

    CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype_rapidPvstPlus CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects_Stpxspanningtreetype = "rapidPvstPlus"
)

// CISCOSTPEXTENSIONSMIB_Stpxmistpobjects
type CISCOSTPEXTENSIONSMIB_Stpxmistpobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of MISTP instances, that are supported by the device  when the
    // value of stpxSpanningTreeType is either mistp(2) or mistpPvstPlus(3). The
    // type is interface{} with range: 1..256.
    Stpxmistpinstancenumber interface{}
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetFilter() yfilter.YFilter { return stpxmistpobjects.YFilter }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) SetFilter(yf yfilter.YFilter) { stpxmistpobjects.YFilter = yf }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetGoName(yname string) string {
    if yname == "stpxMISTPInstanceNumber" { return "Stpxmistpinstancenumber" }
    return ""
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetSegmentPath() string {
    return "stpxMISTPObjects"
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMISTPInstanceNumber"] = stpxmistpobjects.Stpxmistpinstancenumber
    return leafs
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetYangName() string { return "stpxMISTPObjects" }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) SetParent(parent types.Entity) { stpxmistpobjects.parent = parent }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetParent() types.Entity { return stpxmistpobjects.parent }

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects
type CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the global default config mode of LoopGuard  feature on the
    // device. The type is Stpxloopguardglobaldefaultmode.
    Stpxloopguardglobaldefaultmode interface{}
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetFilter() yfilter.YFilter { return stpxloopguardobjects.YFilter }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) SetFilter(yf yfilter.YFilter) { stpxloopguardobjects.YFilter = yf }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetGoName(yname string) string {
    if yname == "stpxLoopGuardGlobalDefaultMode" { return "Stpxloopguardglobaldefaultmode" }
    return ""
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetSegmentPath() string {
    return "stpxLoopGuardObjects"
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxLoopGuardGlobalDefaultMode"] = stpxloopguardobjects.Stpxloopguardglobaldefaultmode
    return leafs
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetYangName() string { return "stpxLoopGuardObjects" }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) SetParent(parent types.Entity) { stpxloopguardobjects.parent = parent }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetParent() types.Entity { return stpxloopguardobjects.parent }

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode represents feature on the device.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode_enable CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode_disable CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode = "disable"
)

// CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects
type CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates the global default mode of the Bpdu Guard feature on the device. 
    // On platforms that does not support per port  Bpdu Guard configuration as
    // indicated by the object stpxFastStartPortBpduGuardMode, if  the value of
    // this object is set to true(1),  and the Fast Start Feature is operationally
    // enabled on a port, then that port will be  immediately disabled when the
    // system receives a BPDU from that port. The type is bool.
    Stpxfaststartbpduguardenable interface{}

    // Indicates the global default mode of the Bpdu  Filter feature on the
    // device.  On platforms that does not support per port  Bpdu Filter
    // configuration as indicated by the object stpxFastStartPortBpduFilterMode,
    // if  the value of this object is set to true(1),  and the Fast Start Feature
    // is operationally  enabled on a port, then no BPDUs will be  transmitted on
    // this port. The type is bool.
    Stpxfaststartbpdufilterenable interface{}

    // Indicates the global default mode of the Fast  Start feature on the device.
    // The type is Stpxfaststartglobaldefaultmode.
    Stpxfaststartglobaldefaultmode interface{}
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetFilter() yfilter.YFilter { return stpxfaststartobjects.YFilter }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) SetFilter(yf yfilter.YFilter) { stpxfaststartobjects.YFilter = yf }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetGoName(yname string) string {
    if yname == "stpxFastStartBpduGuardEnable" { return "Stpxfaststartbpduguardenable" }
    if yname == "stpxFastStartBpduFilterEnable" { return "Stpxfaststartbpdufilterenable" }
    if yname == "stpxFastStartGlobalDefaultMode" { return "Stpxfaststartglobaldefaultmode" }
    return ""
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetSegmentPath() string {
    return "stpxFastStartObjects"
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxFastStartBpduGuardEnable"] = stpxfaststartobjects.Stpxfaststartbpduguardenable
    leafs["stpxFastStartBpduFilterEnable"] = stpxfaststartobjects.Stpxfaststartbpdufilterenable
    leafs["stpxFastStartGlobalDefaultMode"] = stpxfaststartobjects.Stpxfaststartglobaldefaultmode
    return leafs
}

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetYangName() string { return "stpxFastStartObjects" }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) SetParent(parent types.Entity) { stpxfaststartobjects.parent = parent }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetParent() types.Entity { return stpxfaststartobjects.parent }

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode represents Start feature on the device.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode_enable CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode_disable CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode = "disable"
)

// CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects
type CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates whether BPDU skewing detection feature is enabled or not on the
    // system. If this object has the value of true(1), then the system will
    // detect whether BPDUs received by any port on any Spanning  Tree instance
    // are processed at an interval longer than the object value of
    // dot1dStpHelloTime in the BIRDGE-MIB of the Spanning Tree instance. The type
    // is bool.
    Stpxbpduskewingdetectionenable interface{}
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetFilter() yfilter.YFilter { return stpxbpduskewingobjects.YFilter }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) SetFilter(yf yfilter.YFilter) { stpxbpduskewingobjects.YFilter = yf }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetGoName(yname string) string {
    if yname == "stpxBpduSkewingDetectionEnable" { return "Stpxbpduskewingdetectionenable" }
    return ""
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetSegmentPath() string {
    return "stpxBpduSkewingObjects"
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxBpduSkewingDetectionEnable"] = stpxbpduskewingobjects.Stpxbpduskewingdetectionenable
    return leafs
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetYangName() string { return "stpxBpduSkewingObjects" }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) SetParent(parent types.Entity) { stpxbpduskewingobjects.parent = parent }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetParent() types.Entity { return stpxbpduskewingobjects.parent }

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstobjects
type CISCOSTPEXTENSIONSMIB_Stpxmstobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The maximum MST (Multiple Spanning Tree) instance id,  that can be
    // supported by the device for Cisco proprietary implementation of the MST
    // Protocol.  This object is deprecated and replaced by 
    // stpxSMSTMaxInstanceID. The type is interface{} with range: 1..256.
    Stpxmstmaxinstancenumber interface{}

    // The operational MST region name. The type is string with length: 0..32.
    Stpxmstregionname interface{}

    // The operational MST region version.  This object is deprecated and replaced
    // by  stpxSMSTRegionRevision. The type is interface{} with range: 0..65535.
    Stpxmstregionrevision interface{}

    // Indicates the current ownership status of the unique  Region Config Edit
    // Buffer.   released -- the Edit Buffer can be acquired by any of            
    // the SNMP management stations.   acquiredBySnmp -- the Edit Buffer is
    // acquired by             any of the SNMP management stations.  
    // acquiredByNonSnmp -- the Edit Buffer is acquired by the             
    // non-SNMP users managing the device. The type is
    // Stpxmstregioneditbufferstatus.
    Stpxmstregioneditbufferstatus interface{}

    // Indicates the operation that is performed on the Region  Config Edit
    // Buffer.  other --   none of the following operations.    acquire -- acquire
    // the Edit Buffer. This operation can             only be performed when the
    // object             stpxMSTRegionEditBufferStatus has the value of          
    // released(1). After the successful operation of             this action, the
    // stpxMSTRegionEditBufferStatus            will be changed to
    // acquiredBySnmp(2).               releaseWithForce -- release the Edit
    // Buffer acquired by            non-SNMP users with force and discard the
    // changes            in the Edit Buffer. This operation can only be          
    // performed when the object             stpxMSTRegionEditBufferStatus has the
    // value of             acquiredByNonSnmp(2).  commit --  commit the changes
    // in the Edit Buffer            and release the Edit Buffer. The successful  
    // operation of this action will make the changes            in the Edit
    // Buffer effective on the device.            This operation can only be
    // performed when the             object stpxMSTRegionEditBufferStatus has the
    // value of acquiredBySnmp(3).   rollBack -- discard the changes in the Edit
    // Buffer            and release the Edit Buffer. This operation can          
    // only be performed when the object             stpxMSTRegionEditBufferStatus
    // has the value             of acquiredBySnmp(3).  This object always returns
    // other(1) when it is read. The type is Stpxmstregioneditbufferoperation.
    Stpxmstregioneditbufferoperation interface{}

    // The MST region name in the Edit Buffer.   This object is only instantiated
    // when the  stpxMSTRegionEditBufferStatus has the value of 
    // acquiredBySnmp(2). The type is string with length: 0..32.
    Stpxmstregioneditname interface{}

    // The MST region version in the Edit Buffer. This object is only instantiated
    // when the stpxMSTRegionEditBufferStatus  has the value of acquiredBySnmp(2).
    // This object is deprecated and replaced by stpxSMSTRegionEditRevision. The
    // type is interface{} with range: 1..65535.
    Stpxmstregioneditrevision interface{}

    // The maximum number of hops for the MST region.   This object will take on
    // value of 40 if the object value of stpxSMSTMaxHopCount is greater than 40. 
    // This object is deprecated and replaced by stpxSMSTMaxHopCount. The type is
    // interface{} with range: 1..40.
    Stpxmstmaxhopcount interface{}
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetFilter() yfilter.YFilter { return stpxmstobjects.YFilter }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) SetFilter(yf yfilter.YFilter) { stpxmstobjects.YFilter = yf }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetGoName(yname string) string {
    if yname == "stpxMSTMaxInstanceNumber" { return "Stpxmstmaxinstancenumber" }
    if yname == "stpxMSTRegionName" { return "Stpxmstregionname" }
    if yname == "stpxMSTRegionRevision" { return "Stpxmstregionrevision" }
    if yname == "stpxMSTRegionEditBufferStatus" { return "Stpxmstregioneditbufferstatus" }
    if yname == "stpxMSTRegionEditBufferOperation" { return "Stpxmstregioneditbufferoperation" }
    if yname == "stpxMSTRegionEditName" { return "Stpxmstregioneditname" }
    if yname == "stpxMSTRegionEditRevision" { return "Stpxmstregioneditrevision" }
    if yname == "stpxMSTMaxHopCount" { return "Stpxmstmaxhopcount" }
    return ""
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetSegmentPath() string {
    return "stpxMSTObjects"
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMSTMaxInstanceNumber"] = stpxmstobjects.Stpxmstmaxinstancenumber
    leafs["stpxMSTRegionName"] = stpxmstobjects.Stpxmstregionname
    leafs["stpxMSTRegionRevision"] = stpxmstobjects.Stpxmstregionrevision
    leafs["stpxMSTRegionEditBufferStatus"] = stpxmstobjects.Stpxmstregioneditbufferstatus
    leafs["stpxMSTRegionEditBufferOperation"] = stpxmstobjects.Stpxmstregioneditbufferoperation
    leafs["stpxMSTRegionEditName"] = stpxmstobjects.Stpxmstregioneditname
    leafs["stpxMSTRegionEditRevision"] = stpxmstobjects.Stpxmstregioneditrevision
    leafs["stpxMSTMaxHopCount"] = stpxmstobjects.Stpxmstmaxhopcount
    return leafs
}

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetYangName() string { return "stpxMSTObjects" }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) SetParent(parent types.Entity) { stpxmstobjects.parent = parent }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetParent() types.Entity { return stpxmstobjects.parent }

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation represents This object always returns other(1) when it is read.
type CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation string

const (
    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation_other CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation = "other"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation_acquire CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation = "acquire"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation_releaseWithForce CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation = "releaseWithForce"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation_commit CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation = "commit"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation_rollBack CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferoperation = "rollBack"
)

// CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus represents             non-SNMP users managing the device.
type CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus string

const (
    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus_released CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus = "released"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus_acquiredBySnmp CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus = "acquiredBySnmp"

    CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus_acquiredByNonSnmp CISCOSTPEXTENSIONSMIB_Stpxmstobjects_Stpxmstregioneditbufferstatus = "acquiredByNonSnmp"
)

// CISCOSTPEXTENSIONSMIB_Stpxrstpobjects
type CISCOSTPEXTENSIONSMIB_Stpxrstpobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Transmit Hold Count. The type is interface{} with range: 0..4294967295.
    Stpxrstptransmitholdcount interface{}
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetFilter() yfilter.YFilter { return stpxrstpobjects.YFilter }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) SetFilter(yf yfilter.YFilter) { stpxrstpobjects.YFilter = yf }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetGoName(yname string) string {
    if yname == "stpxRSTPTransmitHoldCount" { return "Stpxrstptransmitholdcount" }
    return ""
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetSegmentPath() string {
    return "stpxRSTPObjects"
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRSTPTransmitHoldCount"] = stpxrstpobjects.Stpxrstptransmitholdcount
    return leafs
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetYangName() string { return "stpxRSTPObjects" }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) SetParent(parent types.Entity) { stpxrstpobjects.parent = parent }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetParent() types.Entity { return stpxrstpobjects.parent }

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstobjects
type CISCOSTPEXTENSIONSMIB_Stpxsmstobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The maximum number of MST instances that can be  supported by the device
    // for IEEE MST. The type is interface{} with range: 0..4294967295.
    Stpxsmstmaxinstances interface{}

    // The maximum MST instance ID that can be supported  by the device for IEEE
    // MST. The type is interface{} with range: 0..4294967295.
    Stpxsmstmaxinstanceid interface{}

    // The operational region version for IEEE MST. The type is interface{} with
    // range: 0..4294967295.
    Stpxsmstregionrevision interface{}

    // The MST region version in the Edit Buffer for IEEE  MST.  This object is
    // only instantiated when the  stpxMSTRegionEditBufferStatus has the value of 
    // acquiredBySnmp(2). The type is interface{} with range: 0..4294967295.
    Stpxsmstregioneditrevision interface{}

    // The maximum number of hops for the IEEE MST region. The type is interface{}
    // with range: 0..4294967295.
    Stpxsmstmaxhopcount interface{}

    // The IEEE MST region configuration digest. The type is string.
    Stpxsmstconfigdigest interface{}

    // The pre-standard MST region configuration digest. The type is string.
    Stpxsmstconfigprestandarddigest interface{}
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetFilter() yfilter.YFilter { return stpxsmstobjects.YFilter }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) SetFilter(yf yfilter.YFilter) { stpxsmstobjects.YFilter = yf }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetGoName(yname string) string {
    if yname == "stpxSMSTMaxInstances" { return "Stpxsmstmaxinstances" }
    if yname == "stpxSMSTMaxInstanceID" { return "Stpxsmstmaxinstanceid" }
    if yname == "stpxSMSTRegionRevision" { return "Stpxsmstregionrevision" }
    if yname == "stpxSMSTRegionEditRevision" { return "Stpxsmstregioneditrevision" }
    if yname == "stpxSMSTMaxHopCount" { return "Stpxsmstmaxhopcount" }
    if yname == "stpxSMSTConfigDigest" { return "Stpxsmstconfigdigest" }
    if yname == "stpxSMSTConfigPreStandardDigest" { return "Stpxsmstconfigprestandarddigest" }
    return ""
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetSegmentPath() string {
    return "stpxSMSTObjects"
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxSMSTMaxInstances"] = stpxsmstobjects.Stpxsmstmaxinstances
    leafs["stpxSMSTMaxInstanceID"] = stpxsmstobjects.Stpxsmstmaxinstanceid
    leafs["stpxSMSTRegionRevision"] = stpxsmstobjects.Stpxsmstregionrevision
    leafs["stpxSMSTRegionEditRevision"] = stpxsmstobjects.Stpxsmstregioneditrevision
    leafs["stpxSMSTMaxHopCount"] = stpxsmstobjects.Stpxsmstmaxhopcount
    leafs["stpxSMSTConfigDigest"] = stpxsmstobjects.Stpxsmstconfigdigest
    leafs["stpxSMSTConfigPreStandardDigest"] = stpxsmstobjects.Stpxsmstconfigprestandarddigest
    return leafs
}

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetYangName() string { return "stpxSMSTObjects" }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) SetParent(parent types.Entity) { stpxsmstobjects.parent = parent }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetParent() types.Entity { return stpxsmstobjects.parent }

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable
// A list of Virtual LAN entries containing
// information for Spanning Tree PVST+ protocol. 
// An entry will exist for each VLAN existing on 
// the device.
type CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing Spanning Tree PVST+ Protocol  information for a
    // particular Virtual LAN. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry.
    Stpxpvstvlanentry []CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetFilter() yfilter.YFilter { return stpxpvstvlantable.YFilter }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) SetFilter(yf yfilter.YFilter) { stpxpvstvlantable.YFilter = yf }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetGoName(yname string) string {
    if yname == "stpxPVSTVlanEntry" { return "Stpxpvstvlanentry" }
    return ""
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetSegmentPath() string {
    return "stpxPVSTVlanTable"
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxPVSTVlanEntry" {
        for _, c := range stpxpvstvlantable.Stpxpvstvlanentry {
            if stpxpvstvlantable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry{}
        stpxpvstvlantable.Stpxpvstvlanentry = append(stpxpvstvlantable.Stpxpvstvlanentry, child)
        return &stpxpvstvlantable.Stpxpvstvlanentry[len(stpxpvstvlantable.Stpxpvstvlanentry)-1]
    }
    return nil
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxpvstvlantable.Stpxpvstvlanentry {
        children[stpxpvstvlantable.Stpxpvstvlanentry[i].GetSegmentPath()] = &stpxpvstvlantable.Stpxpvstvlanentry[i]
    }
    return children
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetYangName() string { return "stpxPVSTVlanTable" }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) SetParent(parent types.Entity) { stpxpvstvlantable.parent = parent }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetParent() types.Entity { return stpxpvstvlantable.parent }

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry
// An entry containing Spanning Tree PVST+ Protocol 
// information for a particular Virtual LAN.
type CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An index value that uniquely identifies the
    // Virtual LAN associated with this information. The type is interface{} with
    // range: 0..4095.
    Stpxpvstvlanindex interface{}

    // Indicates whether Spanning Tree PVST+   Protocol is enabled for this
    // Virtual LAN. If  Spanning Tree PVST+ Protocol is not supported  on this
    // VLAN, then notApplicable(3) will be  returned while retrieving the object
    // value for  this VLAN.  If the device only supports a single global Spanning
    // Tree PVST+ Protocol enable/disable  for all the existing VLANs, then the
    // object  value assigned to this VLAN will be applied to the object values of
    // all the instances in this table which do not have the value of
    // notApplicable(3).  If the value of stpxSpanningTreeType is neither 
    // pvstPlus(1) nor rapidPvstPlus(5), then the value  of stpxPVSTVlanEnable for
    // this VLAN can not be  changed. The type is Stpxpvstvlanenable.
    Stpxpvstvlanenable interface{}
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetFilter() yfilter.YFilter { return stpxpvstvlanentry.YFilter }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) SetFilter(yf yfilter.YFilter) { stpxpvstvlanentry.YFilter = yf }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetGoName(yname string) string {
    if yname == "stpxPVSTVlanIndex" { return "Stpxpvstvlanindex" }
    if yname == "stpxPVSTVlanEnable" { return "Stpxpvstvlanenable" }
    return ""
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetSegmentPath() string {
    return "stpxPVSTVlanEntry" + "[stpxPVSTVlanIndex='" + fmt.Sprintf("%v", stpxpvstvlanentry.Stpxpvstvlanindex) + "']"
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxPVSTVlanIndex"] = stpxpvstvlanentry.Stpxpvstvlanindex
    leafs["stpxPVSTVlanEnable"] = stpxpvstvlanentry.Stpxpvstvlanenable
    return leafs
}

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetYangName() string { return "stpxPVSTVlanEntry" }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) SetParent(parent types.Entity) { stpxpvstvlanentry.parent = parent }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetParent() types.Entity { return stpxpvstvlanentry.parent }

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetParentYangName() string { return "stpxPVSTVlanTable" }

// CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable represents changed.
type CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable string

const (
    CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable_enabled CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable = "enabled"

    CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable_disabled CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable = "disabled"

    CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable_notApplicable CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry_Stpxpvstvlanenable = "notApplicable"
)

// CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable
// A table containing a list of the ports for which
// a particular VLAN's Spanning Tree has been found to
// have an inconsistency.  Two types of inconsistency
// are discovered: 1) an inconsistency where two different
// port types have been plugged together; and 2) an
// inconsistency where different switches have different
// PVIDs for the same link.
type CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VLAN on a particular port for which a Spanning Tree inconsistency is
    // currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry.
    Stpxinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetFilter() yfilter.YFilter { return stpxinconsistencytable.YFilter }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) SetFilter(yf yfilter.YFilter) { stpxinconsistencytable.YFilter = yf }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetGoName(yname string) string {
    if yname == "stpxInconsistencyEntry" { return "Stpxinconsistencyentry" }
    return ""
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetSegmentPath() string {
    return "stpxInconsistencyTable"
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxInconsistencyEntry" {
        for _, c := range stpxinconsistencytable.Stpxinconsistencyentry {
            if stpxinconsistencytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry{}
        stpxinconsistencytable.Stpxinconsistencyentry = append(stpxinconsistencytable.Stpxinconsistencyentry, child)
        return &stpxinconsistencytable.Stpxinconsistencyentry[len(stpxinconsistencytable.Stpxinconsistencyentry)-1]
    }
    return nil
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxinconsistencytable.Stpxinconsistencyentry {
        children[stpxinconsistencytable.Stpxinconsistencyentry[i].GetSegmentPath()] = &stpxinconsistencytable.Stpxinconsistencyentry[i]
    }
    return children
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetYangName() string { return "stpxInconsistencyTable" }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) SetParent(parent types.Entity) { stpxinconsistencytable.parent = parent }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetParent() types.Entity { return stpxinconsistencytable.parent }

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry
// A VLAN on a particular port for which a Spanning Tree
// inconsistency is currently in effect.
type CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    Stpxvlanindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxportindex interface{}

    // The types of inconsistency which have been discovered on this port for this
    // VLAN's Spanning Tree.  When this object exists, the value of the
    // corresponding instance of the Bridge MIB's dot1dStpPortState object will be
    // 'broken(6)'. The type is map[string]bool.
    Stpxinconsistentstate interface{}
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetFilter() yfilter.YFilter { return stpxinconsistencyentry.YFilter }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) SetFilter(yf yfilter.YFilter) { stpxinconsistencyentry.YFilter = yf }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetGoName(yname string) string {
    if yname == "stpxVlanIndex" { return "Stpxvlanindex" }
    if yname == "stpxPortIndex" { return "Stpxportindex" }
    if yname == "stpxInconsistentState" { return "Stpxinconsistentstate" }
    return ""
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetSegmentPath() string {
    return "stpxInconsistencyEntry" + "[stpxVlanIndex='" + fmt.Sprintf("%v", stpxinconsistencyentry.Stpxvlanindex) + "']" + "[stpxPortIndex='" + fmt.Sprintf("%v", stpxinconsistencyentry.Stpxportindex) + "']"
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxVlanIndex"] = stpxinconsistencyentry.Stpxvlanindex
    leafs["stpxPortIndex"] = stpxinconsistencyentry.Stpxportindex
    leafs["stpxInconsistentState"] = stpxinconsistencyentry.Stpxinconsistentstate
    return leafs
}

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetYangName() string { return "stpxInconsistencyEntry" }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) SetParent(parent types.Entity) { stpxinconsistencyentry.parent = parent }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetParent() types.Entity { return stpxinconsistencyentry.parent }

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetParentYangName() string { return "stpxInconsistencyTable" }

// CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable
// A table containing a list of the bridge ports for which
// Spanning Tree RootGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree RootGuard capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry.
    Stpxrootguardconfigentry []CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetFilter() yfilter.YFilter { return stpxrootguardconfigtable.YFilter }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) SetFilter(yf yfilter.YFilter) { stpxrootguardconfigtable.YFilter = yf }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetGoName(yname string) string {
    if yname == "stpxRootGuardConfigEntry" { return "Stpxrootguardconfigentry" }
    return ""
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetSegmentPath() string {
    return "stpxRootGuardConfigTable"
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxRootGuardConfigEntry" {
        for _, c := range stpxrootguardconfigtable.Stpxrootguardconfigentry {
            if stpxrootguardconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry{}
        stpxrootguardconfigtable.Stpxrootguardconfigentry = append(stpxrootguardconfigtable.Stpxrootguardconfigentry, child)
        return &stpxrootguardconfigtable.Stpxrootguardconfigentry[len(stpxrootguardconfigtable.Stpxrootguardconfigentry)-1]
    }
    return nil
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxrootguardconfigtable.Stpxrootguardconfigentry {
        children[stpxrootguardconfigtable.Stpxrootguardconfigentry[i].GetSegmentPath()] = &stpxrootguardconfigtable.Stpxrootguardconfigentry[i]
    }
    return children
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetYangName() string { return "stpxRootGuardConfigTable" }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) SetParent(parent types.Entity) { stpxrootguardconfigtable.parent = parent }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetParent() types.Entity { return stpxrootguardconfigtable.parent }

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry
// A bridge port for which Spanning Tree RootGuard
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrootguardconfigportindex interface{}

    // An indication of whether the RootGuard capability is  enabled on this port
    // or not. This configuration will be applied to all Spanning Tree instances
    // in which this port  exists. The type is bool.
    Stpxrootguardconfigenabled interface{}
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetFilter() yfilter.YFilter { return stpxrootguardconfigentry.YFilter }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) SetFilter(yf yfilter.YFilter) { stpxrootguardconfigentry.YFilter = yf }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetGoName(yname string) string {
    if yname == "stpxRootGuardConfigPortIndex" { return "Stpxrootguardconfigportindex" }
    if yname == "stpxRootGuardConfigEnabled" { return "Stpxrootguardconfigenabled" }
    return ""
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetSegmentPath() string {
    return "stpxRootGuardConfigEntry" + "[stpxRootGuardConfigPortIndex='" + fmt.Sprintf("%v", stpxrootguardconfigentry.Stpxrootguardconfigportindex) + "']"
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRootGuardConfigPortIndex"] = stpxrootguardconfigentry.Stpxrootguardconfigportindex
    leafs["stpxRootGuardConfigEnabled"] = stpxrootguardconfigentry.Stpxrootguardconfigenabled
    return leafs
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetYangName() string { return "stpxRootGuardConfigEntry" }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) SetParent(parent types.Entity) { stpxrootguardconfigentry.parent = parent }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetParent() types.Entity { return stpxrootguardconfigentry.parent }

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetParentYangName() string { return "stpxRootGuardConfigTable" }

// CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable
// A table containing a list of the bridge ports for which
// a particular Spanning Tree instance has been found 
// to have an root-inconsistency. The agent creates a new 
// entry in this table whenever it detects a new 
// root-inconsistency, and deletes entries 
// when/soon after the inconsistency is no longer present.
type CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for  which a Spanning Tree
    // root-inconsistency is currently  in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry.
    Stpxrootinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetFilter() yfilter.YFilter { return stpxrootinconsistencytable.YFilter }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) SetFilter(yf yfilter.YFilter) { stpxrootinconsistencytable.YFilter = yf }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetGoName(yname string) string {
    if yname == "stpxRootInconsistencyEntry" { return "Stpxrootinconsistencyentry" }
    return ""
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetSegmentPath() string {
    return "stpxRootInconsistencyTable"
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxRootInconsistencyEntry" {
        for _, c := range stpxrootinconsistencytable.Stpxrootinconsistencyentry {
            if stpxrootinconsistencytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry{}
        stpxrootinconsistencytable.Stpxrootinconsistencyentry = append(stpxrootinconsistencytable.Stpxrootinconsistencyentry, child)
        return &stpxrootinconsistencytable.Stpxrootinconsistencyentry[len(stpxrootinconsistencytable.Stpxrootinconsistencyentry)-1]
    }
    return nil
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxrootinconsistencytable.Stpxrootinconsistencyentry {
        children[stpxrootinconsistencytable.Stpxrootinconsistencyentry[i].GetSegmentPath()] = &stpxrootinconsistencytable.Stpxrootinconsistencyentry[i]
    }
    return children
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetYangName() string { return "stpxRootInconsistencyTable" }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) SetParent(parent types.Entity) { stpxrootinconsistencytable.parent = parent }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetParent() types.Entity { return stpxrootinconsistencytable.parent }

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry
// A Spanning Tree instance on a particular port for 
// which a Spanning Tree root-inconsistency is currently 
// in effect.
type CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1) or
    // rapidPvstPlus(5). The type is interface{} with range: 0..65535.
    Stpxrootinconsistencyindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrootinconsistencyportindex interface{}

    // Indicates whether the port on a particular Spanning  Tree instance is
    // currently in root-inconsistent  state or not. The type is bool.
    Stpxrootinconsistencystate interface{}
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetFilter() yfilter.YFilter { return stpxrootinconsistencyentry.YFilter }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) SetFilter(yf yfilter.YFilter) { stpxrootinconsistencyentry.YFilter = yf }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetGoName(yname string) string {
    if yname == "stpxRootInconsistencyIndex" { return "Stpxrootinconsistencyindex" }
    if yname == "stpxRootInconsistencyPortIndex" { return "Stpxrootinconsistencyportindex" }
    if yname == "stpxRootInconsistencyState" { return "Stpxrootinconsistencystate" }
    return ""
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetSegmentPath() string {
    return "stpxRootInconsistencyEntry" + "[stpxRootInconsistencyIndex='" + fmt.Sprintf("%v", stpxrootinconsistencyentry.Stpxrootinconsistencyindex) + "']" + "[stpxRootInconsistencyPortIndex='" + fmt.Sprintf("%v", stpxrootinconsistencyentry.Stpxrootinconsistencyportindex) + "']"
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRootInconsistencyIndex"] = stpxrootinconsistencyentry.Stpxrootinconsistencyindex
    leafs["stpxRootInconsistencyPortIndex"] = stpxrootinconsistencyentry.Stpxrootinconsistencyportindex
    leafs["stpxRootInconsistencyState"] = stpxrootinconsistencyentry.Stpxrootinconsistencystate
    return leafs
}

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetYangName() string { return "stpxRootInconsistencyEntry" }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) SetParent(parent types.Entity) { stpxrootinconsistencyentry.parent = parent }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetParent() types.Entity { return stpxrootinconsistencyentry.parent }

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetParentYangName() string { return "stpxRootInconsistencyTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable
// This table contains one entry for each instance of MISTP and 
// it contains stpxMISTPInstanceNumber entries, numbered from 1
// to stpxMISTPInstanceNumber.
// 
// This table is only instantiated when the value of 
// stpxSpanningTreeType is mistp(2) or mistpPvstPlus(3).
type CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row containing the status of the MISTP  instance. The type is
    // slice of
    // CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry.
    Stpxmistpinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetFilter() yfilter.YFilter { return stpxmistpinstancetable.YFilter }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) SetFilter(yf yfilter.YFilter) { stpxmistpinstancetable.YFilter = yf }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetGoName(yname string) string {
    if yname == "stpxMISTPInstanceEntry" { return "Stpxmistpinstanceentry" }
    return ""
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetSegmentPath() string {
    return "stpxMISTPInstanceTable"
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxMISTPInstanceEntry" {
        for _, c := range stpxmistpinstancetable.Stpxmistpinstanceentry {
            if stpxmistpinstancetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry{}
        stpxmistpinstancetable.Stpxmistpinstanceentry = append(stpxmistpinstancetable.Stpxmistpinstanceentry, child)
        return &stpxmistpinstancetable.Stpxmistpinstanceentry[len(stpxmistpinstancetable.Stpxmistpinstanceentry)-1]
    }
    return nil
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxmistpinstancetable.Stpxmistpinstanceentry {
        children[stpxmistpinstancetable.Stpxmistpinstanceentry[i].GetSegmentPath()] = &stpxmistpinstancetable.Stpxmistpinstanceentry[i]
    }
    return children
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetYangName() string { return "stpxMISTPInstanceTable" }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) SetParent(parent types.Entity) { stpxmistpinstancetable.parent = parent }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetParent() types.Entity { return stpxmistpinstancetable.parent }

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry
// A conceptual row containing the status of the MISTP 
// instance.
type CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer within the range from 1 to
    // the value of stpxMISTPInstanceNumber that uniquely identifies an instance.
    // The type is interface{} with range: 1..256.
    Stpxmistpinstanceindex interface{}

    // This object indicates whether the MISTP protocol is currently enabled on
    // the MISTP instance.  If this object is set to    'true'    - the MISTP
    // protocol will run on this instance.                   'false'   - the MISTP
    // protocol will stop running on this                 instance. The type is
    // bool.
    Stpxmistpinstanceenable interface{}

    // A string of octets containing one bit per VLAN. The first octet corresponds
    // to VLANs with VlanIndex values of 0 through 7; the second octet to VLANs 8
    // through 15; etc.  The most significant bit of each octet corresponds to the
    // lowest value VlanIndex in that octet.  For each VLAN, if it is mapped to
    // this MISTP instance, then the bit corresponding to that VLAN is set to '1'.
    // The type is string with length: 0..128.
    Stpxmistpinstancevlansmapped interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 1024 through 1031; the second octet to VLANs 1032
    // through 1039; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    Stpxmistpinstancevlansmapped2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    Stpxmistpinstancevlansmapped3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 3072 through 3079; the second octet to VLANs 3080
    // through 3087; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    Stpxmistpinstancevlansmapped4K interface{}
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetFilter() yfilter.YFilter { return stpxmistpinstanceentry.YFilter }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) SetFilter(yf yfilter.YFilter) { stpxmistpinstanceentry.YFilter = yf }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetGoName(yname string) string {
    if yname == "stpxMISTPInstanceIndex" { return "Stpxmistpinstanceindex" }
    if yname == "stpxMISTPInstanceEnable" { return "Stpxmistpinstanceenable" }
    if yname == "stpxMISTPInstanceVlansMapped" { return "Stpxmistpinstancevlansmapped" }
    if yname == "stpxMISTPInstanceVlansMapped2k" { return "Stpxmistpinstancevlansmapped2K" }
    if yname == "stpxMISTPInstanceVlansMapped3k" { return "Stpxmistpinstancevlansmapped3K" }
    if yname == "stpxMISTPInstanceVlansMapped4k" { return "Stpxmistpinstancevlansmapped4K" }
    return ""
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetSegmentPath() string {
    return "stpxMISTPInstanceEntry" + "[stpxMISTPInstanceIndex='" + fmt.Sprintf("%v", stpxmistpinstanceentry.Stpxmistpinstanceindex) + "']"
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMISTPInstanceIndex"] = stpxmistpinstanceentry.Stpxmistpinstanceindex
    leafs["stpxMISTPInstanceEnable"] = stpxmistpinstanceentry.Stpxmistpinstanceenable
    leafs["stpxMISTPInstanceVlansMapped"] = stpxmistpinstanceentry.Stpxmistpinstancevlansmapped
    leafs["stpxMISTPInstanceVlansMapped2k"] = stpxmistpinstanceentry.Stpxmistpinstancevlansmapped2K
    leafs["stpxMISTPInstanceVlansMapped3k"] = stpxmistpinstanceentry.Stpxmistpinstancevlansmapped3K
    leafs["stpxMISTPInstanceVlansMapped4k"] = stpxmistpinstanceentry.Stpxmistpinstancevlansmapped4K
    return leafs
}

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetYangName() string { return "stpxMISTPInstanceEntry" }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) SetParent(parent types.Entity) { stpxmistpinstanceentry.parent = parent }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetParent() types.Entity { return stpxmistpinstanceentry.parent }

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetParentYangName() string { return "stpxMISTPInstanceTable" }

// CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable
// A table containing a list of the bridge ports for which
// Spanning Tree LoopGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree LoopGuard  capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry.
    Stpxloopguardconfigentry []CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetFilter() yfilter.YFilter { return stpxloopguardconfigtable.YFilter }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) SetFilter(yf yfilter.YFilter) { stpxloopguardconfigtable.YFilter = yf }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetGoName(yname string) string {
    if yname == "stpxLoopGuardConfigEntry" { return "Stpxloopguardconfigentry" }
    return ""
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetSegmentPath() string {
    return "stpxLoopGuardConfigTable"
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxLoopGuardConfigEntry" {
        for _, c := range stpxloopguardconfigtable.Stpxloopguardconfigentry {
            if stpxloopguardconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry{}
        stpxloopguardconfigtable.Stpxloopguardconfigentry = append(stpxloopguardconfigtable.Stpxloopguardconfigentry, child)
        return &stpxloopguardconfigtable.Stpxloopguardconfigentry[len(stpxloopguardconfigtable.Stpxloopguardconfigentry)-1]
    }
    return nil
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxloopguardconfigtable.Stpxloopguardconfigentry {
        children[stpxloopguardconfigtable.Stpxloopguardconfigentry[i].GetSegmentPath()] = &stpxloopguardconfigtable.Stpxloopguardconfigentry[i]
    }
    return children
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetYangName() string { return "stpxLoopGuardConfigTable" }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) SetParent(parent types.Entity) { stpxloopguardconfigtable.parent = parent }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetParent() types.Entity { return stpxloopguardconfigtable.parent }

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry
// A bridge port for which Spanning Tree LoopGuard 
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxloopguardconfigportindex interface{}

    // An indication of whether the LoopGuard capability is  enabled on this port
    // or not. This configuration will be applied to all the Spanning Tree
    // instances in which this  port exists.  In order to support additional Loop
    // Guard config mode (default) as defined in stpxLoopGuardConfigMode other 
    // than enable (true(1)) or disable (false(2)) as defined  in this object,
    // stpxLoopGuardConfigMode object needs to  be used.  When the
    // stpxLoopGuardConfigMode object has the value of enable(1), the value of
    // stpxLoopGuardConfigEnabled for  the same instance will be true(1). When the
    // stpxLoopGuardConfigMode object has the value of disable(2),  the value of
    // stpxLoopGuardConfigEnabled for the same  instance will be false(2). When
    // the stpxLoopGuardConfigMode  object has the value of default(3), the value
    // of  stpxLoopGuardConfigEnabled for the same instance will  depend on the
    // object value of  stpxLoopGuardGlobalDefaultMode. The type is bool.
    Stpxloopguardconfigenabled interface{}

    // Indicates the mode of Loop Guard Feature on this  port. This configuration
    // will be applied to all  the Spanning Tree instances in which this port 
    // exists.  enable -- the Loop Guard feature is enabled on this           
    // port.   disable -- the Loop Guard feature is disabled on this           
    // port.    default -- whether the Loop Guard feature is enabled            or
    // not on this port depends on the object             value of
    // stpxLoopGuardGlobalDefaultMode. The type is Stpxloopguardconfigmode.
    Stpxloopguardconfigmode interface{}
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetFilter() yfilter.YFilter { return stpxloopguardconfigentry.YFilter }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) SetFilter(yf yfilter.YFilter) { stpxloopguardconfigentry.YFilter = yf }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetGoName(yname string) string {
    if yname == "stpxLoopGuardConfigPortIndex" { return "Stpxloopguardconfigportindex" }
    if yname == "stpxLoopGuardConfigEnabled" { return "Stpxloopguardconfigenabled" }
    if yname == "stpxLoopGuardConfigMode" { return "Stpxloopguardconfigmode" }
    return ""
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetSegmentPath() string {
    return "stpxLoopGuardConfigEntry" + "[stpxLoopGuardConfigPortIndex='" + fmt.Sprintf("%v", stpxloopguardconfigentry.Stpxloopguardconfigportindex) + "']"
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxLoopGuardConfigPortIndex"] = stpxloopguardconfigentry.Stpxloopguardconfigportindex
    leafs["stpxLoopGuardConfigEnabled"] = stpxloopguardconfigentry.Stpxloopguardconfigenabled
    leafs["stpxLoopGuardConfigMode"] = stpxloopguardconfigentry.Stpxloopguardconfigmode
    return leafs
}

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetYangName() string { return "stpxLoopGuardConfigEntry" }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) SetParent(parent types.Entity) { stpxloopguardconfigentry.parent = parent }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetParent() types.Entity { return stpxloopguardconfigentry.parent }

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetParentYangName() string { return "stpxLoopGuardConfigTable" }

// CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode represents            value of stpxLoopGuardGlobalDefaultMode.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode_enable CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode_disable CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode = "disable"

    CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode_default_ CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry_Stpxloopguardconfigmode = "default"
)

// CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable
// A table containing a list of the bridge ports for which
// a particular Spanning Tree instance has been found
// to have a loop-inconsistency. The agent creates a new
// entry in this table whenever it detects a new
// loop-inconsistency, and deletes entries
// when/soon after the inconsistency is no longer present.
type CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which a Spanning Tree
    // loop-inconsistency is currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry.
    Stpxloopinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetFilter() yfilter.YFilter { return stpxloopinconsistencytable.YFilter }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) SetFilter(yf yfilter.YFilter) { stpxloopinconsistencytable.YFilter = yf }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetGoName(yname string) string {
    if yname == "stpxLoopInconsistencyEntry" { return "Stpxloopinconsistencyentry" }
    return ""
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetSegmentPath() string {
    return "stpxLoopInconsistencyTable"
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxLoopInconsistencyEntry" {
        for _, c := range stpxloopinconsistencytable.Stpxloopinconsistencyentry {
            if stpxloopinconsistencytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry{}
        stpxloopinconsistencytable.Stpxloopinconsistencyentry = append(stpxloopinconsistencytable.Stpxloopinconsistencyentry, child)
        return &stpxloopinconsistencytable.Stpxloopinconsistencyentry[len(stpxloopinconsistencytable.Stpxloopinconsistencyentry)-1]
    }
    return nil
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxloopinconsistencytable.Stpxloopinconsistencyentry {
        children[stpxloopinconsistencytable.Stpxloopinconsistencyentry[i].GetSegmentPath()] = &stpxloopinconsistencytable.Stpxloopinconsistencyentry[i]
    }
    return children
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetYangName() string { return "stpxLoopInconsistencyTable" }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) SetParent(parent types.Entity) { stpxloopinconsistencytable.parent = parent }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetParent() types.Entity { return stpxloopinconsistencytable.parent }

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry
// A Spanning Tree instance on a particular port for
// which a Spanning Tree loop-inconsistency is currently
// in effect.
type CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1) or
    // rapidPvstPlus(5). The type is interface{} with range: 0..65535.
    Stpxloopinconsistencyindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxloopinconsistencyportindex interface{}

    // Indicates whether the port on a particular Spanning  Tree instance is
    // currently in loop-inconsistent  state or not. The type is bool.
    Stpxloopinconsistencystate interface{}
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetFilter() yfilter.YFilter { return stpxloopinconsistencyentry.YFilter }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) SetFilter(yf yfilter.YFilter) { stpxloopinconsistencyentry.YFilter = yf }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetGoName(yname string) string {
    if yname == "stpxLoopInconsistencyIndex" { return "Stpxloopinconsistencyindex" }
    if yname == "stpxLoopInconsistencyPortIndex" { return "Stpxloopinconsistencyportindex" }
    if yname == "stpxLoopInconsistencyState" { return "Stpxloopinconsistencystate" }
    return ""
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetSegmentPath() string {
    return "stpxLoopInconsistencyEntry" + "[stpxLoopInconsistencyIndex='" + fmt.Sprintf("%v", stpxloopinconsistencyentry.Stpxloopinconsistencyindex) + "']" + "[stpxLoopInconsistencyPortIndex='" + fmt.Sprintf("%v", stpxloopinconsistencyentry.Stpxloopinconsistencyportindex) + "']"
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxLoopInconsistencyIndex"] = stpxloopinconsistencyentry.Stpxloopinconsistencyindex
    leafs["stpxLoopInconsistencyPortIndex"] = stpxloopinconsistencyentry.Stpxloopinconsistencyportindex
    leafs["stpxLoopInconsistencyState"] = stpxloopinconsistencyentry.Stpxloopinconsistencystate
    return leafs
}

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetYangName() string { return "stpxLoopInconsistencyEntry" }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) SetParent(parent types.Entity) { stpxloopinconsistencyentry.parent = parent }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetParent() types.Entity { return stpxloopinconsistencyentry.parent }

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetParentYangName() string { return "stpxLoopInconsistencyTable" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable
// A table containing a list of the bridge ports for
// which Spanning Tree Port Fast Start can be
// configured.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree Port Fast Start can be configured.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry.
    Stpxfaststartportentry []CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetFilter() yfilter.YFilter { return stpxfaststartporttable.YFilter }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) SetFilter(yf yfilter.YFilter) { stpxfaststartporttable.YFilter = yf }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetGoName(yname string) string {
    if yname == "stpxFastStartPortEntry" { return "Stpxfaststartportentry" }
    return ""
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetSegmentPath() string {
    return "stpxFastStartPortTable"
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxFastStartPortEntry" {
        for _, c := range stpxfaststartporttable.Stpxfaststartportentry {
            if stpxfaststartporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry{}
        stpxfaststartporttable.Stpxfaststartportentry = append(stpxfaststartporttable.Stpxfaststartportentry, child)
        return &stpxfaststartporttable.Stpxfaststartportentry[len(stpxfaststartporttable.Stpxfaststartportentry)-1]
    }
    return nil
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxfaststartporttable.Stpxfaststartportentry {
        children[stpxfaststartporttable.Stpxfaststartportentry[i].GetSegmentPath()] = &stpxfaststartporttable.Stpxfaststartportentry[i]
    }
    return children
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetYangName() string { return "stpxFastStartPortTable" }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) SetParent(parent types.Entity) { stpxfaststartporttable.parent = parent }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetParent() types.Entity { return stpxfaststartporttable.parent }

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry
// A bridge port for which Spanning Tree Port Fast
// Start can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxfaststartportindex interface{}

    // Indicates whether the port is operating in spantree fast start mode.  A
    // port with fast start enabled is immediately put in spanning tree forwarding
    // state when that port is detected by the Spanning Tree, rather  than
    // starting in blocking state which is the normal  operation.  In order to
    // support additional Fast Start enable mode (enableForTrunk and default) as
    // defined in stpxFastStartPortMode other than enable (true(1)) or disable
    // (false(2)) as defined in this object, stpxFastStartPortMode object needs to
    // be used.  When the stpxFastStartPortMode has the value of enable(1) or
    // enableForTrunk(3), the value of stpxFastStartPortEnable for the same
    // instance will be true(1). When the stpxFastStartPortMode has the value of
    // disable(2), the value of  stpxFastStartPortEnable for the same instance
    // will be  false(2). When the stpxFastStartPortMode has the value  of
    // default(4), the value of stpxFastStartPortEnable for  the same instance
    // depends on the object value of  stpxFastStartGlobalDefaultMode. The type is
    // bool.
    Stpxfaststartportenable interface{}

    // Indicates the mode of Fast Start Feature on the  port. A port with fast
    // start enabled is immediately  put in spanning tree forwarding state when
    // the port is detected by the Spanning Tree, rather than  starting in
    // blocking state which is the normal  operation.  enable -- the fast start
    // feature is enabled on this            port but will only take effect when
    // the            object value of its            vlanTrunkPortDynamicStatus as
    // specified            in CISCO-VTP-MIB is notTrunking(2).  disable -- the
    // fast start feature is disabled on this            port.    enableForTrunk
    // -- the fast start feature is enabled            on this port and will take
    // effect            regardless of the object value of            its
    // vlanTrunkPortDynamicStatus.  default -- whether the fast start feature is
    // enabled            or not on this port depends on the object            
    // value of stpxFastStartGlobalDefaultMode.  network -- the fast start network
    // mode is enabled on             this port. The type is
    // Stpxfaststartportmode.
    Stpxfaststartportmode interface{}

    // Indicates the mode of Bpdu Guard Feature on the port. A port with Bpdu
    // Guard enabled is  immediately disabled when the system  receives a BPDU
    // from that port.   enable -- the Bpdu Guard feature is enabled on this      
    // port.   disable -- the Bpdu Guard feature is disabled on this          
    // port.  default -- whether the Bpdu Guard feature is enabled            or
    // not on this port depends on the object            value of
    // stpxFastStartBpduGuardEnable. If             the value of
    // stpxFastStartBpduGuardEnable            is true(1) and Fast Start feature
    // is also             enabled operationally on this port, then           
    // this port is immediately disabled when             the system receives a
    // BPDU from this port. The type is Stpxfaststartportbpduguardmode.
    Stpxfaststartportbpduguardmode interface{}

    // Indicates the mode of Bpdu Filter Feature on the port. The system will not
    // transmit BPDUs on a port  with Bpdu Filter feature enabled.  enable -- the
    // Bpdu Filter feature is enabled on this            port.   disable -- the
    // Bpdu Filter feature is disabled on this            port.  default --
    // whether the Bpdu Filter feature is enabled            or not on this port
    // depends on the object            value of stpxFastStartBpduFilterEnable. If
    // the value of stpxFastStartBpduFilterEnable            is true(1) and Fast
    // Start feature is also            enabled operationally on this port, then  
    // no BPDUs will be transmitted on this port. The type is
    // Stpxfaststartportbpdufiltermode.
    Stpxfaststartportbpdufiltermode interface{}
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetFilter() yfilter.YFilter { return stpxfaststartportentry.YFilter }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) SetFilter(yf yfilter.YFilter) { stpxfaststartportentry.YFilter = yf }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetGoName(yname string) string {
    if yname == "stpxFastStartPortIndex" { return "Stpxfaststartportindex" }
    if yname == "stpxFastStartPortEnable" { return "Stpxfaststartportenable" }
    if yname == "stpxFastStartPortMode" { return "Stpxfaststartportmode" }
    if yname == "stpxFastStartPortBpduGuardMode" { return "Stpxfaststartportbpduguardmode" }
    if yname == "stpxFastStartPortBpduFilterMode" { return "Stpxfaststartportbpdufiltermode" }
    return ""
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetSegmentPath() string {
    return "stpxFastStartPortEntry" + "[stpxFastStartPortIndex='" + fmt.Sprintf("%v", stpxfaststartportentry.Stpxfaststartportindex) + "']"
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxFastStartPortIndex"] = stpxfaststartportentry.Stpxfaststartportindex
    leafs["stpxFastStartPortEnable"] = stpxfaststartportentry.Stpxfaststartportenable
    leafs["stpxFastStartPortMode"] = stpxfaststartportentry.Stpxfaststartportmode
    leafs["stpxFastStartPortBpduGuardMode"] = stpxfaststartportentry.Stpxfaststartportbpduguardmode
    leafs["stpxFastStartPortBpduFilterMode"] = stpxfaststartportentry.Stpxfaststartportbpdufiltermode
    return leafs
}

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetYangName() string { return "stpxFastStartPortEntry" }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) SetParent(parent types.Entity) { stpxfaststartportentry.parent = parent }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetParent() types.Entity { return stpxfaststartportentry.parent }

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetParentYangName() string { return "stpxFastStartPortTable" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode represents            no BPDUs will be transmitted on this port.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode_enable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode_disable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode = "disable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode_default_ CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpdufiltermode = "default"
)

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode represents            the system receives a BPDU from this port.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode_enable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode_disable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode = "disable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode_default_ CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportbpduguardmode = "default"
)

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode represents            this port.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode_enable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode_disable CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode = "disable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode_enableForTrunk CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode = "enableForTrunk"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode_default_ CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode = "default"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode_network CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry_Stpxfaststartportmode = "network"
)

// CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable
// A table containing a list of the bridge ports 
// for a particular Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with port fast start oper mode  information on a bridge port for a
    // particular  Spanning Tree Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry.
    Stpxfaststartopermodeentry []CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetFilter() yfilter.YFilter { return stpxfaststartopermodetable.YFilter }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) SetFilter(yf yfilter.YFilter) { stpxfaststartopermodetable.YFilter = yf }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetGoName(yname string) string {
    if yname == "stpxFastStartOperModeEntry" { return "Stpxfaststartopermodeentry" }
    return ""
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetSegmentPath() string {
    return "stpxFastStartOperModeTable"
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxFastStartOperModeEntry" {
        for _, c := range stpxfaststartopermodetable.Stpxfaststartopermodeentry {
            if stpxfaststartopermodetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry{}
        stpxfaststartopermodetable.Stpxfaststartopermodeentry = append(stpxfaststartopermodetable.Stpxfaststartopermodeentry, child)
        return &stpxfaststartopermodetable.Stpxfaststartopermodeentry[len(stpxfaststartopermodetable.Stpxfaststartopermodeentry)-1]
    }
    return nil
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxfaststartopermodetable.Stpxfaststartopermodeentry {
        children[stpxfaststartopermodetable.Stpxfaststartopermodeentry[i].GetSegmentPath()] = &stpxfaststartopermodetable.Stpxfaststartopermodeentry[i]
    }
    return children
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetYangName() string { return "stpxFastStartOperModeTable" }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) SetParent(parent types.Entity) { stpxfaststartopermodetable.parent = parent }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetParent() types.Entity { return stpxfaststartopermodetable.parent }

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry
// An entry with port fast start oper mode 
// information on a bridge port for a particular 
// Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1). The
    // type is interface{} with range: 0..65535.
    Stpxfaststartopermodeinstindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxfaststartopermodeportindex interface{}

    // Indicates the fast start operational status of the  port on a particular
    // Spanning Tree Instance. The type is Stpxfaststartopermode.
    Stpxfaststartopermode interface{}
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetFilter() yfilter.YFilter { return stpxfaststartopermodeentry.YFilter }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) SetFilter(yf yfilter.YFilter) { stpxfaststartopermodeentry.YFilter = yf }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetGoName(yname string) string {
    if yname == "stpxFastStartOperModeInstIndex" { return "Stpxfaststartopermodeinstindex" }
    if yname == "stpxFastStartOperModePortIndex" { return "Stpxfaststartopermodeportindex" }
    if yname == "stpxFastStartOperMode" { return "Stpxfaststartopermode" }
    return ""
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetSegmentPath() string {
    return "stpxFastStartOperModeEntry" + "[stpxFastStartOperModeInstIndex='" + fmt.Sprintf("%v", stpxfaststartopermodeentry.Stpxfaststartopermodeinstindex) + "']" + "[stpxFastStartOperModePortIndex='" + fmt.Sprintf("%v", stpxfaststartopermodeentry.Stpxfaststartopermodeportindex) + "']"
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxFastStartOperModeInstIndex"] = stpxfaststartopermodeentry.Stpxfaststartopermodeinstindex
    leafs["stpxFastStartOperModePortIndex"] = stpxfaststartopermodeentry.Stpxfaststartopermodeportindex
    leafs["stpxFastStartOperMode"] = stpxfaststartopermodeentry.Stpxfaststartopermode
    return leafs
}

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetYangName() string { return "stpxFastStartOperModeEntry" }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) SetParent(parent types.Entity) { stpxfaststartopermodeentry.parent = parent }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetParent() types.Entity { return stpxfaststartopermodeentry.parent }

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetParentYangName() string { return "stpxFastStartOperModeTable" }

// CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode represents port on a particular Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode_enabled CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode = "enabled"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode_disabled CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry_Stpxfaststartopermode = "disabled"
)

// CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable
// A table containing a list of the bridge ports for 
// which a particular Spanning Tree instance has been 
// detected to have BPDU skewing occurred since the 
// object value of stpxBpduSkewingDetectionEnable was
// last changed to true(1).
// 
// The agent creates a new entry in this table whenever
// a port in a particular Spanning Tree instance is 
// detected to be BPDU skewed since the object value of 
// stpxBpduSkewingDetectionEnable object is changed to 
// true(1). The agent deletes all the entries in this 
// table when the object value of 
// stpxBpduSkewingDetectionEnable is changed to false(2)
// or the object value of stpxSpanningTreeType is 
// changed.
type CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which BPDU skewing has
    // been detected. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry.
    Stpxbpduskewingentry []CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetFilter() yfilter.YFilter { return stpxbpduskewingtable.YFilter }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) SetFilter(yf yfilter.YFilter) { stpxbpduskewingtable.YFilter = yf }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetGoName(yname string) string {
    if yname == "stpxBpduSkewingEntry" { return "Stpxbpduskewingentry" }
    return ""
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetSegmentPath() string {
    return "stpxBpduSkewingTable"
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxBpduSkewingEntry" {
        for _, c := range stpxbpduskewingtable.Stpxbpduskewingentry {
            if stpxbpduskewingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry{}
        stpxbpduskewingtable.Stpxbpduskewingentry = append(stpxbpduskewingtable.Stpxbpduskewingentry, child)
        return &stpxbpduskewingtable.Stpxbpduskewingentry[len(stpxbpduskewingtable.Stpxbpduskewingentry)-1]
    }
    return nil
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxbpduskewingtable.Stpxbpduskewingentry {
        children[stpxbpduskewingtable.Stpxbpduskewingentry[i].GetSegmentPath()] = &stpxbpduskewingtable.Stpxbpduskewingentry[i]
    }
    return children
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetYangName() string { return "stpxBpduSkewingTable" }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) SetParent(parent types.Entity) { stpxbpduskewingtable.parent = parent }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetParent() types.Entity { return stpxbpduskewingtable.parent }

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry
// A Spanning Tree instance on a particular port for
// which BPDU skewing has been detected.
type CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType  is pvstPlus(1).
    // The type is interface{} with range: 0..65535.
    Stpxbpduskewinginstanceindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxbpduskewingportindex interface{}

    // Indicates the skew duration in milliseconds of the last BPDU skewing
    // detected. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Stpxbpduskewinglastskewduration interface{}

    // Indicates the skew duration in milliseconds of the worst BPDU skewing
    // detected. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    Stpxbpduskewingworstskewduration interface{}

    // Indicates the value of sysUpTime when the worst BPDU skewing was detected.
    // The type is interface{} with range: 0..4294967295.
    Stpxbpduskewingworstskewtime interface{}
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetFilter() yfilter.YFilter { return stpxbpduskewingentry.YFilter }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) SetFilter(yf yfilter.YFilter) { stpxbpduskewingentry.YFilter = yf }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetGoName(yname string) string {
    if yname == "stpxBpduSkewingInstanceIndex" { return "Stpxbpduskewinginstanceindex" }
    if yname == "stpxBpduSkewingPortIndex" { return "Stpxbpduskewingportindex" }
    if yname == "stpxBpduSkewingLastSkewDuration" { return "Stpxbpduskewinglastskewduration" }
    if yname == "stpxBpduSkewingWorstSkewDuration" { return "Stpxbpduskewingworstskewduration" }
    if yname == "stpxBpduSkewingWorstSkewTime" { return "Stpxbpduskewingworstskewtime" }
    return ""
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetSegmentPath() string {
    return "stpxBpduSkewingEntry" + "[stpxBpduSkewingInstanceIndex='" + fmt.Sprintf("%v", stpxbpduskewingentry.Stpxbpduskewinginstanceindex) + "']" + "[stpxBpduSkewingPortIndex='" + fmt.Sprintf("%v", stpxbpduskewingentry.Stpxbpduskewingportindex) + "']"
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxBpduSkewingInstanceIndex"] = stpxbpduskewingentry.Stpxbpduskewinginstanceindex
    leafs["stpxBpduSkewingPortIndex"] = stpxbpduskewingentry.Stpxbpduskewingportindex
    leafs["stpxBpduSkewingLastSkewDuration"] = stpxbpduskewingentry.Stpxbpduskewinglastskewduration
    leafs["stpxBpduSkewingWorstSkewDuration"] = stpxbpduskewingentry.Stpxbpduskewingworstskewduration
    leafs["stpxBpduSkewingWorstSkewTime"] = stpxbpduskewingentry.Stpxbpduskewingworstskewtime
    return leafs
}

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetYangName() string { return "stpxBpduSkewingEntry" }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) SetParent(parent types.Entity) { stpxbpduskewingentry.parent = parent }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetParent() types.Entity { return stpxbpduskewingentry.parent }

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetParentYangName() string { return "stpxBpduSkewingTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable
// This table contains MST instance information with
// one entry for an MST instance within the range of 
// 0 to the object value of stpxMSTMaxInstanceNumber. 
// 
// This table is deprecated and replaced by 
// stpxSMSTInstanceTable.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information. The type is
    // slice of CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry.
    Stpxmstinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetFilter() yfilter.YFilter { return stpxmstinstancetable.YFilter }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) SetFilter(yf yfilter.YFilter) { stpxmstinstancetable.YFilter = yf }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetGoName(yname string) string {
    if yname == "stpxMSTInstanceEntry" { return "Stpxmstinstanceentry" }
    return ""
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetSegmentPath() string {
    return "stpxMSTInstanceTable"
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxMSTInstanceEntry" {
        for _, c := range stpxmstinstancetable.Stpxmstinstanceentry {
            if stpxmstinstancetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry{}
        stpxmstinstancetable.Stpxmstinstanceentry = append(stpxmstinstancetable.Stpxmstinstanceentry, child)
        return &stpxmstinstancetable.Stpxmstinstanceentry[len(stpxmstinstancetable.Stpxmstinstanceentry)-1]
    }
    return nil
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxmstinstancetable.Stpxmstinstanceentry {
        children[stpxmstinstancetable.Stpxmstinstanceentry[i].GetSegmentPath()] = &stpxmstinstancetable.Stpxmstinstanceentry[i]
    }
    return children
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetYangName() string { return "stpxMSTInstanceTable" }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) SetParent(parent types.Entity) { stpxmstinstancetable.parent = parent }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetParent() types.Entity { return stpxmstinstancetable.parent }

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry
// A conceptual row containing the MST instance 
// information.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An integer that uniquely identifies an MST
    // instance  within the range of 0 to the object value of
    // stpxMSTMaxInstanceNumber.  This object is deprecated and replaced by 
    // stpxSMSTInstanceIndex. The type is interface{} with range: 0..256.
    Stpxmstinstanceindex interface{}

    // A string of octets containing one bit per VLAN. The first octet corresponds
    // to VLANs with VlanIndex values of 0 through 7; the second octet to VLANs 8
    // through 15; etc.  The most significant bit of each octet corresponds to the
    // lowest value VlanIndex in that octet.  For each VLAN, if it is mapped to
    // this MST instance,  then the bit corresponding to that VLAN is set to '1'. 
    // This object is deprecated and replaced by  stpxSMSTInstanceVlansMapped1k2k.
    // The type is string with length: 0..128.
    Stpxmstinstancevlansmapped interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 1024 through 1031; the second octet to VLANs 1032
    // through 1039; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by 
    // stpxSMSTInstanceVlansMapped1k2k. The type is string with length: 0..128.
    Stpxmstinstancevlansmapped2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by 
    // stpxSMSTInstanceVlansMapped3k4k. The type is string with length: 0..128.
    Stpxmstinstancevlansmapped3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 3072 through 3079; the second octet to VLANs 3080
    // through 3087; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by
    // stpxSMSTInstanceVlansMapped3k4k. The type is string with length: 0..128.
    Stpxmstinstancevlansmapped4K interface{}

    // The remaining hop count for this MST instance.  This object will take on
    // value of 40 if the object value of stpxSMSTInstanceRemainingHopCount is
    // greater than 40.  This object is only instantiated when the object value of
    // stpxSpanningTreeType is mst(4).  This object is deprecated and replaced by 
    // stpxSMSTInstanceRemainingHopCount. The type is interface{} with range:
    // 0..40.
    Stpxmstinstanceremaininghopcount interface{}
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetFilter() yfilter.YFilter { return stpxmstinstanceentry.YFilter }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) SetFilter(yf yfilter.YFilter) { stpxmstinstanceentry.YFilter = yf }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetGoName(yname string) string {
    if yname == "stpxMSTInstanceIndex" { return "Stpxmstinstanceindex" }
    if yname == "stpxMSTInstanceVlansMapped" { return "Stpxmstinstancevlansmapped" }
    if yname == "stpxMSTInstanceVlansMapped2k" { return "Stpxmstinstancevlansmapped2K" }
    if yname == "stpxMSTInstanceVlansMapped3k" { return "Stpxmstinstancevlansmapped3K" }
    if yname == "stpxMSTInstanceVlansMapped4k" { return "Stpxmstinstancevlansmapped4K" }
    if yname == "stpxMSTInstanceRemainingHopCount" { return "Stpxmstinstanceremaininghopcount" }
    return ""
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetSegmentPath() string {
    return "stpxMSTInstanceEntry" + "[stpxMSTInstanceIndex='" + fmt.Sprintf("%v", stpxmstinstanceentry.Stpxmstinstanceindex) + "']"
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMSTInstanceIndex"] = stpxmstinstanceentry.Stpxmstinstanceindex
    leafs["stpxMSTInstanceVlansMapped"] = stpxmstinstanceentry.Stpxmstinstancevlansmapped
    leafs["stpxMSTInstanceVlansMapped2k"] = stpxmstinstanceentry.Stpxmstinstancevlansmapped2K
    leafs["stpxMSTInstanceVlansMapped3k"] = stpxmstinstanceentry.Stpxmstinstancevlansmapped3K
    leafs["stpxMSTInstanceVlansMapped4k"] = stpxmstinstanceentry.Stpxmstinstancevlansmapped4K
    leafs["stpxMSTInstanceRemainingHopCount"] = stpxmstinstanceentry.Stpxmstinstanceremaininghopcount
    return leafs
}

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetYangName() string { return "stpxMSTInstanceEntry" }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) SetParent(parent types.Entity) { stpxmstinstanceentry.parent = parent }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetParent() types.Entity { return stpxmstinstanceentry.parent }

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetParentYangName() string { return "stpxMSTInstanceTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable
// This table contains MST instance information in the 
// Edit Buffer with one entry for each MST
// instance numbered from 0 to stpxMSTMaxInstanceNumber. 
// 
// This table is only instantiated when the 
// stpxMSTRegionEditBufferStatus has the value of
// acquiredBySnmp(2).
// 
// This table is deprecated and replaced by 
// stpxSMSTInstanceEditTable.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry.
    Stpxmstinstanceeditentry []CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetFilter() yfilter.YFilter { return stpxmstinstanceedittable.YFilter }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) SetFilter(yf yfilter.YFilter) { stpxmstinstanceedittable.YFilter = yf }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetGoName(yname string) string {
    if yname == "stpxMSTInstanceEditEntry" { return "Stpxmstinstanceeditentry" }
    return ""
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetSegmentPath() string {
    return "stpxMSTInstanceEditTable"
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxMSTInstanceEditEntry" {
        for _, c := range stpxmstinstanceedittable.Stpxmstinstanceeditentry {
            if stpxmstinstanceedittable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry{}
        stpxmstinstanceedittable.Stpxmstinstanceeditentry = append(stpxmstinstanceedittable.Stpxmstinstanceeditentry, child)
        return &stpxmstinstanceedittable.Stpxmstinstanceeditentry[len(stpxmstinstanceedittable.Stpxmstinstanceeditentry)-1]
    }
    return nil
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxmstinstanceedittable.Stpxmstinstanceeditentry {
        children[stpxmstinstanceedittable.Stpxmstinstanceeditentry[i].GetSegmentPath()] = &stpxmstinstanceedittable.Stpxmstinstanceeditentry[i]
    }
    return children
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetYangName() string { return "stpxMSTInstanceEditTable" }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) SetParent(parent types.Entity) { stpxmstinstanceedittable.parent = parent }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetParent() types.Entity { return stpxmstinstanceedittable.parent }

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An integer that uniquely identifies an MST
    // instance  from 0 to the object value of stpxMSTMaxInstanceNumber.  The
    // instances of this table entry with  stpxMSTInstanceEditIndex of zero can
    // not be  modified.  This object is deprecated and replaced by 
    // stpxSMSTInstanceEditIndex. The type is interface{} with range: 0..256.
    Stpxmstinstanceeditindex interface{}

    // A string of octets containing one bit per VLAN. The first octet corresponds
    // to VLANs with VlanIndex values of 0 through 7; the second octet to VLANs 8
    // through 15; etc.  The most significant bit of each octet corresponds to the
    // lowest value VlanIndex in that octet.  For each VLAN, if it is mapped to
    // this MST instance,  then the bit corresponding to that VLAN is set to  '1'.
    // Each VLAN can only be mapped to one unique MST  instance in the range from
    // 1 to stpxMSTMaxInstanceNumber. If the bit corresponding to a VLAN is
    // changed from '1'  to '0', then that VLAN will be automatically mapped to 
    // MST instance 0 by the device.  This object is deprecated and replaced by 
    // stpxSMSTInstanceEditVlansMap1k2k. The type is string with length: 0..128.
    Stpxmstinstanceeditvlansmap interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 1024 through 1031; the second octet to VLANs 1032
    // through 1039; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. Each VLAN can only be mapped to one unique MST instance in the range
    // from 1 to stpxMSTMaxInstanceNumber. If the bit corresponding to a VLAN is
    // changed from '1'  to '0', then that VLAN will be automatically mapped to 
    // MST instance 0 by the device.  This object is deprecated and replaced by 
    // stpxSMSTInstanceEditVlansMap1k2k. The type is string with length: 0..128.
    Stpxmstinstanceeditvlansmap2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. Each VLAN can only be mapped to one unique MST instance in the range
    // from 1 to stpxMSTMaxInstanceNumber. If the bit corresponding to a VLAN is
    // changed from '1'  to '0', then that VLAN will be automatically mapped to 
    // MST instance 0 by the device.  This object is deprecated and replaced by 
    // stpxSMSTInstanceEditVlansMap3k4k. The type is string with length: 0..128.
    Stpxmstinstanceeditvlansmap3K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 3072 through 3079; the second octet to VLANs 3080
    // through 3087; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. Each VLAN can only be mapped to one unique MST instance in the range
    // from 1 to stpxMSTMaxInstanceNumber. If the bit corresponding to a VLAN is
    // changed from '1'  to '0', then that VLAN will be automatically mapped to 
    // MST instance 0 by the device.  This object is deprecated and replaced by
    // stpxSMSTInstanceEditVlansMap3k4k. The type is string with length: 0..128.
    Stpxmstinstanceeditvlansmap4K interface{}
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetFilter() yfilter.YFilter { return stpxmstinstanceeditentry.YFilter }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) SetFilter(yf yfilter.YFilter) { stpxmstinstanceeditentry.YFilter = yf }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetGoName(yname string) string {
    if yname == "stpxMSTInstanceEditIndex" { return "Stpxmstinstanceeditindex" }
    if yname == "stpxMSTInstanceEditVlansMap" { return "Stpxmstinstanceeditvlansmap" }
    if yname == "stpxMSTInstanceEditVlansMap2k" { return "Stpxmstinstanceeditvlansmap2K" }
    if yname == "stpxMSTInstanceEditVlansMap3k" { return "Stpxmstinstanceeditvlansmap3K" }
    if yname == "stpxMSTInstanceEditVlansMap4k" { return "Stpxmstinstanceeditvlansmap4K" }
    return ""
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetSegmentPath() string {
    return "stpxMSTInstanceEditEntry" + "[stpxMSTInstanceEditIndex='" + fmt.Sprintf("%v", stpxmstinstanceeditentry.Stpxmstinstanceeditindex) + "']"
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMSTInstanceEditIndex"] = stpxmstinstanceeditentry.Stpxmstinstanceeditindex
    leafs["stpxMSTInstanceEditVlansMap"] = stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap
    leafs["stpxMSTInstanceEditVlansMap2k"] = stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap2K
    leafs["stpxMSTInstanceEditVlansMap3k"] = stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap3K
    leafs["stpxMSTInstanceEditVlansMap4k"] = stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap4K
    return leafs
}

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetYangName() string { return "stpxMSTInstanceEditEntry" }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) SetParent(parent types.Entity) { stpxmstinstanceeditentry.parent = parent }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetParent() types.Entity { return stpxmstinstanceeditentry.parent }

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetParentYangName() string { return "stpxMSTInstanceEditTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with port information for the MST Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry.
    Stpxmstportentry []CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetFilter() yfilter.YFilter { return stpxmstporttable.YFilter }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) SetFilter(yf yfilter.YFilter) { stpxmstporttable.YFilter = yf }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetGoName(yname string) string {
    if yname == "stpxMSTPortEntry" { return "Stpxmstportentry" }
    return ""
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetSegmentPath() string {
    return "stpxMSTPortTable"
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxMSTPortEntry" {
        for _, c := range stpxmstporttable.Stpxmstportentry {
            if stpxmstporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry{}
        stpxmstporttable.Stpxmstportentry = append(stpxmstporttable.Stpxmstportentry, child)
        return &stpxmstporttable.Stpxmstportentry[len(stpxmstporttable.Stpxmstportentry)-1]
    }
    return nil
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxmstporttable.Stpxmstportentry {
        children[stpxmstporttable.Stpxmstportentry[i].GetSegmentPath()] = &stpxmstporttable.Stpxmstportentry[i]
    }
    return children
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetYangName() string { return "stpxMSTPortTable" }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) SetParent(parent types.Entity) { stpxmstporttable.parent = parent }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetParent() types.Entity { return stpxmstporttable.parent }

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry
// An entry with port information for the MST Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxmstportindex interface{}

    // Indicates the administrative link type configuration of  a bridge port for
    // the MST protocol.   pointToPoint -- the port is administratively configured
    // to         be connected to a point-to-point link.  shared -- the port is
    // administratively configured to be         connected to a shared medium.  
    // auto -- the administrative configuration of the port's          link type
    // depends on link duplex of the port.         If the port link is
    // full-duplex, the administrative          link type configuration on this
    // port will be taken          as pointTopoint(1). If the port link is
    // half-duplex,          the administrative link type configuration on this   
    // port will be taken as shared(2).  This configuration of this object only
    // takes effect when the stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
    // stpxMSTPortAdminLinkType is deprecated and replaced  with
    // stpxRSTPPortAdminLinkType. The type is Stpxmstportadminlinktype.
    Stpxmstportadminlinktype interface{}

    // Indicates the operational link type of a bridge port for the MST protocol. 
    // pointToPoint -- the port is operationally connected to         a
    // point-to-point link.  shared -- the port is operationally connected to     
    // a shared medium.  other -- none of the above.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4). 
    // stpxMSTPortOperLinkType  is deprecated and replaced with
    // stpxRSTPPortOperLinkType. The type is Stpxmstportoperlinktype.
    Stpxmstportoperlinktype interface{}

    // The protocol migration control on this port. When the  object value of 
    // stpxSpanningTreeType is mst(4) or  rapidPvstPlus(5), setting true(1) to
    // this object forces  the device to try using version 2 BPDUs on this port. 
    // When the object value of stpxSpanningTreeType is neither  mst(4) nor
    // rapidPvstPlus(5), setting true(1) to this  object has no effect. Setting
    // false(2) to this object has  no effect. This object always returns false(2)
    // when read. stpxMSTPortProtocolMigration is deprecated and  replaced with
    // stpxRSTPPortProtocolMigration. The type is bool.
    Stpxmstportprotocolmigration interface{}

    // Indicates the operational status of the port for the  MST protocol.   edge
    // -- this port is an edge port for the MST region.  boundary -- this port is
    // a boundary port for the          MST region.  pvst --  this port is
    // connected to a PVST/PVST+ bridge.     stp -- this port is connected to a
    // Single Spanning         Tree bridge.   This object is only instantiated
    // when the object value of stpxSpanningTreeType is mst(4).  This object is
    // deprecated and replaced by  stpxSMSTPortStatus. The type is
    // map[string]bool.
    Stpxmstportstatus interface{}
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetFilter() yfilter.YFilter { return stpxmstportentry.YFilter }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) SetFilter(yf yfilter.YFilter) { stpxmstportentry.YFilter = yf }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetGoName(yname string) string {
    if yname == "stpxMSTPortIndex" { return "Stpxmstportindex" }
    if yname == "stpxMSTPortAdminLinkType" { return "Stpxmstportadminlinktype" }
    if yname == "stpxMSTPortOperLinkType" { return "Stpxmstportoperlinktype" }
    if yname == "stpxMSTPortProtocolMigration" { return "Stpxmstportprotocolmigration" }
    if yname == "stpxMSTPortStatus" { return "Stpxmstportstatus" }
    return ""
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetSegmentPath() string {
    return "stpxMSTPortEntry" + "[stpxMSTPortIndex='" + fmt.Sprintf("%v", stpxmstportentry.Stpxmstportindex) + "']"
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMSTPortIndex"] = stpxmstportentry.Stpxmstportindex
    leafs["stpxMSTPortAdminLinkType"] = stpxmstportentry.Stpxmstportadminlinktype
    leafs["stpxMSTPortOperLinkType"] = stpxmstportentry.Stpxmstportoperlinktype
    leafs["stpxMSTPortProtocolMigration"] = stpxmstportentry.Stpxmstportprotocolmigration
    leafs["stpxMSTPortStatus"] = stpxmstportentry.Stpxmstportstatus
    return leafs
}

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetYangName() string { return "stpxMSTPortEntry" }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) SetParent(parent types.Entity) { stpxmstportentry.parent = parent }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetParent() types.Entity { return stpxmstportentry.parent }

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetParentYangName() string { return "stpxMSTPortTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype represents with stpxRSTPPortAdminLinkType.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype string

const (
    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype_pointToPoint CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype_shared CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype = "shared"

    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype_auto CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportadminlinktype = "auto"
)

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype represents is deprecated and replaced with stpxRSTPPortOperLinkType.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype string

const (
    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype_pointToPoint CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype_shared CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype = "shared"

    CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype_other CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry_Stpxmstportoperlinktype = "other"
)

// CISCOSTPEXTENSIONSMIB_Stpxmstportroletable
// A table containing a list of the bridge ports for a 
// particular MST instance.  This table is only instantiated 
// when the stpxSpanningTreeType is mst(4). 
// 
// This table is deprecated and replaced with 
// stpxRSTPPortRoleTable.
type CISCOSTPEXTENSIONSMIB_Stpxmstportroletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the port role information for the MST protocol on a
    // port for a particular MST instance existing on the system. The type is
    // slice of CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry.
    Stpxmstportroleentry []CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetFilter() yfilter.YFilter { return stpxmstportroletable.YFilter }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) SetFilter(yf yfilter.YFilter) { stpxmstportroletable.YFilter = yf }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetGoName(yname string) string {
    if yname == "stpxMSTPortRoleEntry" { return "Stpxmstportroleentry" }
    return ""
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetSegmentPath() string {
    return "stpxMSTPortRoleTable"
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxMSTPortRoleEntry" {
        for _, c := range stpxmstportroletable.Stpxmstportroleentry {
            if stpxmstportroletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry{}
        stpxmstportroletable.Stpxmstportroleentry = append(stpxmstportroletable.Stpxmstportroleentry, child)
        return &stpxmstportroletable.Stpxmstportroleentry[len(stpxmstportroletable.Stpxmstportroleentry)-1]
    }
    return nil
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxmstportroletable.Stpxmstportroleentry {
        children[stpxmstportroletable.Stpxmstportroleentry[i].GetSegmentPath()] = &stpxmstportroletable.Stpxmstportroleentry[i]
    }
    return children
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetYangName() string { return "stpxMSTPortRoleTable" }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) SetParent(parent types.Entity) { stpxmstportroletable.parent = parent }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetParent() types.Entity { return stpxmstportroletable.parent }

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry
// An entry containing the port role information for the MST
// protocol on a port for a particular MST instance existing
// on the system.
type CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance id within the range of 0 to
    // stpxMSTMaxInstanceNumber. The type is interface{} with range: 0..256.
    Stpxmstportroleinstanceindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxmstportroleportindex interface{}

    // Indicates the port role on a particular MST instance for the MST protocol. 
    // disabled --  this port has no role on this MST instance.   root -- this
    // port has the role of root port on this MST             instance.  
    // designated -- this port has the role of designated              port on
    // this MST instance.  alternate -- this port has the role of alternate port  
    // on this MST instance.  backUp -- this port has the role of backup port on
    // this               MST instance.  boundary -- this port has the role of
    // boundary port on              this MST instance.  master -- this port has
    // the role of master port on           this MST instance. The type is
    // Stpxmstportrolevalue.
    Stpxmstportrolevalue interface{}
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetFilter() yfilter.YFilter { return stpxmstportroleentry.YFilter }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) SetFilter(yf yfilter.YFilter) { stpxmstportroleentry.YFilter = yf }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetGoName(yname string) string {
    if yname == "stpxMSTPortRoleInstanceIndex" { return "Stpxmstportroleinstanceindex" }
    if yname == "stpxMSTPortRolePortIndex" { return "Stpxmstportroleportindex" }
    if yname == "stpxMSTPortRoleValue" { return "Stpxmstportrolevalue" }
    return ""
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetSegmentPath() string {
    return "stpxMSTPortRoleEntry" + "[stpxMSTPortRoleInstanceIndex='" + fmt.Sprintf("%v", stpxmstportroleentry.Stpxmstportroleinstanceindex) + "']" + "[stpxMSTPortRolePortIndex='" + fmt.Sprintf("%v", stpxmstportroleentry.Stpxmstportroleportindex) + "']"
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxMSTPortRoleInstanceIndex"] = stpxmstportroleentry.Stpxmstportroleinstanceindex
    leafs["stpxMSTPortRolePortIndex"] = stpxmstportroleentry.Stpxmstportroleportindex
    leafs["stpxMSTPortRoleValue"] = stpxmstportroleentry.Stpxmstportrolevalue
    return leafs
}

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetYangName() string { return "stpxMSTPortRoleEntry" }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) SetParent(parent types.Entity) { stpxmstportroleentry.parent = parent }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetParent() types.Entity { return stpxmstportroleentry.parent }

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetParentYangName() string { return "stpxMSTPortRoleTable" }

// CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue represents           this MST instance.
type CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue string

const (
    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_disabled CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "disabled"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_root CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "root"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_designated CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "designated"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_alternate CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "alternate"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_backUp CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "backUp"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_boundary CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "boundary"

    CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue_master CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry_Stpxmstportrolevalue = "master"
)

// CISCOSTPEXTENSIONSMIB_Stpxrstpporttable
// A table containing port information for the RSTP 
// Protocol on all the bridge ports existing in the 
// system.
type CISCOSTPEXTENSIONSMIB_Stpxrstpporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with port information for the RSTP Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry.
    Stpxrstpportentry []CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetFilter() yfilter.YFilter { return stpxrstpporttable.YFilter }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) SetFilter(yf yfilter.YFilter) { stpxrstpporttable.YFilter = yf }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetGoName(yname string) string {
    if yname == "stpxRSTPPortEntry" { return "Stpxrstpportentry" }
    return ""
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetSegmentPath() string {
    return "stpxRSTPPortTable"
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxRSTPPortEntry" {
        for _, c := range stpxrstpporttable.Stpxrstpportentry {
            if stpxrstpporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry{}
        stpxrstpporttable.Stpxrstpportentry = append(stpxrstpporttable.Stpxrstpportentry, child)
        return &stpxrstpporttable.Stpxrstpportentry[len(stpxrstpporttable.Stpxrstpportentry)-1]
    }
    return nil
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxrstpporttable.Stpxrstpportentry {
        children[stpxrstpporttable.Stpxrstpportentry[i].GetSegmentPath()] = &stpxrstpporttable.Stpxrstpportentry[i]
    }
    return children
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetYangName() string { return "stpxRSTPPortTable" }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) SetParent(parent types.Entity) { stpxrstpporttable.parent = parent }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetParent() types.Entity { return stpxrstpporttable.parent }

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry
// An entry with port information for the RSTP Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrstpportindex interface{}

    // Indicates the administrative link type configuration of  a bridge port for
    // the RSTP protocol.   pointToPoint -- the port is administratively
    // configured to         be connected to a point-to-point link.  shared -- the
    // port is administratively configured to be         connected to a shared
    // medium.   auto -- the administrative configuration of the port's         
    // link type depends on link duplex of the port.         If the port link is
    // full-duplex, the administrative          link type configuration on this
    // port will be taken          as pointTopoint(1). If the port link is
    // half-duplex,          the administrative link type configuration on this   
    // port will be taken as shared(2).  This configuration of this object only
    // takes effect when the stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
    // The type is Stpxrstpportadminlinktype.
    Stpxrstpportadminlinktype interface{}

    // Indicates the operational link type of a bridge port for the RSTP protocol.
    // pointToPoint -- the port is operationally connected to         a
    // point-to-point link.  shared -- the port is operationally connected to     
    // a shared medium.  other -- none of the above.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4) or
    // rapidPvstPlus(5). The type is Stpxrstpportoperlinktype.
    Stpxrstpportoperlinktype interface{}

    // The protocol migration control on this port. When the  object value of 
    // stpxSpanningTreeType is mst(4) or  rapidPvstPlus(5), setting true(1) to
    // this object forces  the device to try using version 2 BPDUs on this port. 
    // When the object value of stpxSpanningTreeType is neither  mst(4) nor
    // rapidPvstPlus(5), setting true(1) to  this object has no effect. Setting
    // false(2) to this  object has no effect. This object always returns 
    // false(2) when read. The type is bool.
    Stpxrstpportprotocolmigration interface{}
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetFilter() yfilter.YFilter { return stpxrstpportentry.YFilter }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) SetFilter(yf yfilter.YFilter) { stpxrstpportentry.YFilter = yf }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetGoName(yname string) string {
    if yname == "stpxRSTPPortIndex" { return "Stpxrstpportindex" }
    if yname == "stpxRSTPPortAdminLinkType" { return "Stpxrstpportadminlinktype" }
    if yname == "stpxRSTPPortOperLinkType" { return "Stpxrstpportoperlinktype" }
    if yname == "stpxRSTPPortProtocolMigration" { return "Stpxrstpportprotocolmigration" }
    return ""
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetSegmentPath() string {
    return "stpxRSTPPortEntry" + "[stpxRSTPPortIndex='" + fmt.Sprintf("%v", stpxrstpportentry.Stpxrstpportindex) + "']"
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRSTPPortIndex"] = stpxrstpportentry.Stpxrstpportindex
    leafs["stpxRSTPPortAdminLinkType"] = stpxrstpportentry.Stpxrstpportadminlinktype
    leafs["stpxRSTPPortOperLinkType"] = stpxrstpportentry.Stpxrstpportoperlinktype
    leafs["stpxRSTPPortProtocolMigration"] = stpxrstpportentry.Stpxrstpportprotocolmigration
    return leafs
}

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetYangName() string { return "stpxRSTPPortEntry" }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) SetParent(parent types.Entity) { stpxrstpportentry.parent = parent }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetParent() types.Entity { return stpxrstpportentry.parent }

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetParentYangName() string { return "stpxRSTPPortTable" }

// CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype represents stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype string

const (
    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype_pointToPoint CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype_shared CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype = "shared"

    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype_auto CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportadminlinktype = "auto"
)

// CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype represents stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype string

const (
    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype_pointToPoint CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype_shared CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype = "shared"

    CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype_other CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry_Stpxrstpportoperlinktype = "other"
)

// CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable
// A table containing a list of the bridge ports for a 
// particular Spanning Tree instance.  This table is 
// only instantiated when the stpxSpanningTreeType is mst(4) 
// or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the port role information for the RSTP protocol on a
    // port for a particular Spanning Tree instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry.
    Stpxrstpportroleentry []CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetFilter() yfilter.YFilter { return stpxrstpportroletable.YFilter }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) SetFilter(yf yfilter.YFilter) { stpxrstpportroletable.YFilter = yf }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetGoName(yname string) string {
    if yname == "stpxRSTPPortRoleEntry" { return "Stpxrstpportroleentry" }
    return ""
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetSegmentPath() string {
    return "stpxRSTPPortRoleTable"
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxRSTPPortRoleEntry" {
        for _, c := range stpxrstpportroletable.Stpxrstpportroleentry {
            if stpxrstpportroletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry{}
        stpxrstpportroletable.Stpxrstpportroleentry = append(stpxrstpportroletable.Stpxrstpportroleentry, child)
        return &stpxrstpportroletable.Stpxrstpportroleentry[len(stpxrstpportroletable.Stpxrstpportroleentry)-1]
    }
    return nil
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxrstpportroletable.Stpxrstpportroleentry {
        children[stpxrstpportroletable.Stpxrstpportroleentry[i].GetSegmentPath()] = &stpxrstpportroletable.Stpxrstpportroleentry[i]
    }
    return children
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetYangName() string { return "stpxRSTPPortRoleTable" }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) SetParent(parent types.Entity) { stpxrstpportroletable.parent = parent }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetParent() types.Entity { return stpxrstpportroletable.parent }

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry
// An entry containing the port role information for the RSTP
// protocol on a port for a particular Spanning Tree instance.
type CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, it can either be a 
    // VLAN number if the stpxSpanningTreeType is rapidPvstPlus(5)  or an MST
    // instance id if the stpxSpanningTreeType is mst(4). The type is interface{}
    // with range: 0..4095.
    Stpxrstpportroleinstanceindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrstpportroleportindex interface{}

    // Indicates the port role on a particular Spanning Tree  instance for the
    // RSTP protocol.   disabled --  this port has no role in this Spanning       
    // Tree instance.   root -- this port has the role of root port in this       
    // Spanning Tree instance.   designated -- this port has the role of
    // designated              port in this Spanning Tree instance.  alternate --
    // this port has the role of alternate port             in this Spanning Tree
    // instance.  backUp -- this port has the role of backup port in this         
    // Spanning Tree instance.  boundary -- this port has the role of boundary
    // port in              this Spanning Tree instance.  master -- this port has
    // the role of master port in             this Spanning Tree instance.  This
    // object could have a value of 'boundary' or 'master' only when the object
    // value of stpxSpanningTreeType is mst(4). The type is Stpxrstpportrolevalue.
    Stpxrstpportrolevalue interface{}
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetFilter() yfilter.YFilter { return stpxrstpportroleentry.YFilter }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) SetFilter(yf yfilter.YFilter) { stpxrstpportroleentry.YFilter = yf }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetGoName(yname string) string {
    if yname == "stpxRSTPPortRoleInstanceIndex" { return "Stpxrstpportroleinstanceindex" }
    if yname == "stpxRSTPPortRolePortIndex" { return "Stpxrstpportroleportindex" }
    if yname == "stpxRSTPPortRoleValue" { return "Stpxrstpportrolevalue" }
    return ""
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetSegmentPath() string {
    return "stpxRSTPPortRoleEntry" + "[stpxRSTPPortRoleInstanceIndex='" + fmt.Sprintf("%v", stpxrstpportroleentry.Stpxrstpportroleinstanceindex) + "']" + "[stpxRSTPPortRolePortIndex='" + fmt.Sprintf("%v", stpxrstpportroleentry.Stpxrstpportroleportindex) + "']"
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRSTPPortRoleInstanceIndex"] = stpxrstpportroleentry.Stpxrstpportroleinstanceindex
    leafs["stpxRSTPPortRolePortIndex"] = stpxrstpportroleentry.Stpxrstpportroleportindex
    leafs["stpxRSTPPortRoleValue"] = stpxrstpportroleentry.Stpxrstpportrolevalue
    return leafs
}

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetYangName() string { return "stpxRSTPPortRoleEntry" }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) SetParent(parent types.Entity) { stpxrstpportroleentry.parent = parent }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetParent() types.Entity { return stpxrstpportroleentry.parent }

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetParentYangName() string { return "stpxRSTPPortRoleTable" }

// CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue represents only when the object value of stpxSpanningTreeType is mst(4).
type CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue string

const (
    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_disabled CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "disabled"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_root CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "root"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_designated CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "designated"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_alternate CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "alternate"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_backUp CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "backUp"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_boundary CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "boundary"

    CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue_master CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry_Stpxrstpportrolevalue = "master"
)

// CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable
// A table containing a list of the bridge ports 
// for a particular Spanning Tree Instance.
// This table is only instantiated when the object value
// of stpxSpanningTreeType is rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with port status information on a  bridge port for a particular
    // Spanning Tree  Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry.
    Stpxrpvstportentry []CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetFilter() yfilter.YFilter { return stpxrpvstporttable.YFilter }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) SetFilter(yf yfilter.YFilter) { stpxrpvstporttable.YFilter = yf }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetGoName(yname string) string {
    if yname == "stpxRPVSTPortEntry" { return "Stpxrpvstportentry" }
    return ""
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetSegmentPath() string {
    return "stpxRPVSTPortTable"
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxRPVSTPortEntry" {
        for _, c := range stpxrpvstporttable.Stpxrpvstportentry {
            if stpxrpvstporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry{}
        stpxrpvstporttable.Stpxrpvstportentry = append(stpxrpvstporttable.Stpxrpvstportentry, child)
        return &stpxrpvstporttable.Stpxrpvstportentry[len(stpxrpvstporttable.Stpxrpvstportentry)-1]
    }
    return nil
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxrpvstporttable.Stpxrpvstportentry {
        children[stpxrpvstporttable.Stpxrpvstportentry[i].GetSegmentPath()] = &stpxrpvstporttable.Stpxrpvstportentry[i]
    }
    return children
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetYangName() string { return "stpxRPVSTPortTable" }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) SetParent(parent types.Entity) { stpxrpvstporttable.parent = parent }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetParent() types.Entity { return stpxrpvstporttable.parent }

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry
// An entry with port status information on a 
// bridge port for a particular Spanning Tree 
// Instance.
type CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    Stpxrpvstportvlanindex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrpvstportindex interface{}

    // Indicates the operational status of the port for the  Rapid PVST+ protocol.
    // edge -- this port is an edge port for the RST region.  unused1 -- unused
    // bit 1.  unused2 -- unused bit 2.  stp -- this port is connected to a Single
    // Spanning        Tree/PVST+ bridge.  dispute -- this port, as a designated
    // port, received an        inferior BPDU with a designated role and the      
    // learning bit being set. The type is map[string]bool.
    Stpxrpvstportstatus interface{}
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetFilter() yfilter.YFilter { return stpxrpvstportentry.YFilter }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) SetFilter(yf yfilter.YFilter) { stpxrpvstportentry.YFilter = yf }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetGoName(yname string) string {
    if yname == "stpxRPVSTPortVlanIndex" { return "Stpxrpvstportvlanindex" }
    if yname == "stpxRPVSTPortIndex" { return "Stpxrpvstportindex" }
    if yname == "stpxRPVSTPortStatus" { return "Stpxrpvstportstatus" }
    return ""
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetSegmentPath() string {
    return "stpxRPVSTPortEntry" + "[stpxRPVSTPortVlanIndex='" + fmt.Sprintf("%v", stpxrpvstportentry.Stpxrpvstportvlanindex) + "']" + "[stpxRPVSTPortIndex='" + fmt.Sprintf("%v", stpxrpvstportentry.Stpxrpvstportindex) + "']"
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxRPVSTPortVlanIndex"] = stpxrpvstportentry.Stpxrpvstportvlanindex
    leafs["stpxRPVSTPortIndex"] = stpxrpvstportentry.Stpxrpvstportindex
    leafs["stpxRPVSTPortStatus"] = stpxrpvstportentry.Stpxrpvstportstatus
    return leafs
}

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetYangName() string { return "stpxRPVSTPortEntry" }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) SetParent(parent types.Entity) { stpxrpvstportentry.parent = parent }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetParent() types.Entity { return stpxrpvstportentry.parent }

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetParentYangName() string { return "stpxRPVSTPortTable" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable
// This table contains MST instance information
// for IEEE MST.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information for IEEE MST. The
    // type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry.
    Stpxsmstinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetFilter() yfilter.YFilter { return stpxsmstinstancetable.YFilter }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) SetFilter(yf yfilter.YFilter) { stpxsmstinstancetable.YFilter = yf }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetGoName(yname string) string {
    if yname == "stpxSMSTInstanceEntry" { return "Stpxsmstinstanceentry" }
    return ""
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetSegmentPath() string {
    return "stpxSMSTInstanceTable"
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxSMSTInstanceEntry" {
        for _, c := range stpxsmstinstancetable.Stpxsmstinstanceentry {
            if stpxsmstinstancetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry{}
        stpxsmstinstancetable.Stpxsmstinstanceentry = append(stpxsmstinstancetable.Stpxsmstinstanceentry, child)
        return &stpxsmstinstancetable.Stpxsmstinstanceentry[len(stpxsmstinstancetable.Stpxsmstinstanceentry)-1]
    }
    return nil
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxsmstinstancetable.Stpxsmstinstanceentry {
        children[stpxsmstinstancetable.Stpxsmstinstanceentry[i].GetSegmentPath()] = &stpxsmstinstancetable.Stpxsmstinstanceentry[i]
    }
    return children
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetYangName() string { return "stpxSMSTInstanceTable" }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) SetParent(parent types.Entity) { stpxsmstinstancetable.parent = parent }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetParent() types.Entity { return stpxsmstinstancetable.parent }

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry
// A conceptual row containing the MST instance 
// information for IEEE MST.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance ID, the value of which is in the
    // range  from 0 to stpxSMSTMaxInstanceID. The type is interface{} with range:
    // 0..4294967295.
    Stpxsmstinstanceindex interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 0 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 0 through 7; the second octet to VLANs 8 through 15;
    // etc.  The most significant bit of each octet corresponds to the lowest
    // value VlanIndex in that octet.  For each VLAN, if it is mapped to this MST
    // instance, then the bit corresponding to that VLAN is set to '1'. If the
    // length of this string is less than 256 octets, any 'missing' octets are
    // assumed to contain the value  of zero. The type is string with length:
    // 0..256.
    Stpxsmstinstancevlansmapped1K2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. If the length of this string is less than 256 octets, any 'missing'
    // octets are assumed to contain the value  of zero. The type is string with
    // length: 0..256.
    Stpxsmstinstancevlansmapped3K4K interface{}

    // The remaining hop count for this MST instance. If this object value is not
    // applicable on an MST instance, then the value retrieved for this object for
    // that MST instance will be -1.   This object is only instantiated when the
    // object value of stpxSpanningTreeType is mst(4). The type is interface{}
    // with range: -1..2147483647.
    Stpxsmstinstanceremaininghopcount interface{}

    // Indicates the Bridge Identifier (refer to BridgeId  defined in BRIDGE-MIB)
    // of CIST (Common and Internal  Spanning Tree) Regional Root for the MST
    // region.  This object is only instantiated when the object value of
    // stpxSpanningTreeType is mst(4) and stpxSMSTInstanceIndex is 0. The type is
    // string with length: 8.
    Stpxsmstinstancecistregionalroot interface{}

    // Indicates the CIST Internal Root Path Cost, i.e., the path cost to the CIST
    // Regional Root as specified by the corresponding
    // stpxSMSTInstanceCISTRegionalRoot for the  MST region.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4) and
    // stpxSMSTInstanceIndex is 0. The type is interface{} with range:
    // 0..4294967295.
    Stpxsmstinstancecistintrootcost interface{}
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetFilter() yfilter.YFilter { return stpxsmstinstanceentry.YFilter }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) SetFilter(yf yfilter.YFilter) { stpxsmstinstanceentry.YFilter = yf }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetGoName(yname string) string {
    if yname == "stpxSMSTInstanceIndex" { return "Stpxsmstinstanceindex" }
    if yname == "stpxSMSTInstanceVlansMapped1k2k" { return "Stpxsmstinstancevlansmapped1K2K" }
    if yname == "stpxSMSTInstanceVlansMapped3k4k" { return "Stpxsmstinstancevlansmapped3K4K" }
    if yname == "stpxSMSTInstanceRemainingHopCount" { return "Stpxsmstinstanceremaininghopcount" }
    if yname == "stpxSMSTInstanceCISTRegionalRoot" { return "Stpxsmstinstancecistregionalroot" }
    if yname == "stpxSMSTInstanceCISTIntRootCost" { return "Stpxsmstinstancecistintrootcost" }
    return ""
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetSegmentPath() string {
    return "stpxSMSTInstanceEntry" + "[stpxSMSTInstanceIndex='" + fmt.Sprintf("%v", stpxsmstinstanceentry.Stpxsmstinstanceindex) + "']"
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxSMSTInstanceIndex"] = stpxsmstinstanceentry.Stpxsmstinstanceindex
    leafs["stpxSMSTInstanceVlansMapped1k2k"] = stpxsmstinstanceentry.Stpxsmstinstancevlansmapped1K2K
    leafs["stpxSMSTInstanceVlansMapped3k4k"] = stpxsmstinstanceentry.Stpxsmstinstancevlansmapped3K4K
    leafs["stpxSMSTInstanceRemainingHopCount"] = stpxsmstinstanceentry.Stpxsmstinstanceremaininghopcount
    leafs["stpxSMSTInstanceCISTRegionalRoot"] = stpxsmstinstanceentry.Stpxsmstinstancecistregionalroot
    leafs["stpxSMSTInstanceCISTIntRootCost"] = stpxsmstinstanceentry.Stpxsmstinstancecistintrootcost
    return leafs
}

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetYangName() string { return "stpxSMSTInstanceEntry" }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) SetParent(parent types.Entity) { stpxsmstinstanceentry.parent = parent }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetParent() types.Entity { return stpxsmstinstanceentry.parent }

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetParentYangName() string { return "stpxSMSTInstanceTable" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable
// This table contains MST instance information in the 
// Edit Buffer. 
// 
// This table is only instantiated when the object value
// of  stpxMSTRegionEditBufferStatus has the value of
// acquiredBySnmp(2).
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer. 
    // The total number of entries in this table has to be  less than or equal to
    // the object value of stpxSMSTMaxInstances. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry.
    Stpxsmstinstanceeditentry []CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetFilter() yfilter.YFilter { return stpxsmstinstanceedittable.YFilter }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) SetFilter(yf yfilter.YFilter) { stpxsmstinstanceedittable.YFilter = yf }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetGoName(yname string) string {
    if yname == "stpxSMSTInstanceEditEntry" { return "Stpxsmstinstanceeditentry" }
    return ""
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetSegmentPath() string {
    return "stpxSMSTInstanceEditTable"
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxSMSTInstanceEditEntry" {
        for _, c := range stpxsmstinstanceedittable.Stpxsmstinstanceeditentry {
            if stpxsmstinstanceedittable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry{}
        stpxsmstinstanceedittable.Stpxsmstinstanceeditentry = append(stpxsmstinstanceedittable.Stpxsmstinstanceeditentry, child)
        return &stpxsmstinstanceedittable.Stpxsmstinstanceeditentry[len(stpxsmstinstanceedittable.Stpxsmstinstanceeditentry)-1]
    }
    return nil
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxsmstinstanceedittable.Stpxsmstinstanceeditentry {
        children[stpxsmstinstanceedittable.Stpxsmstinstanceeditentry[i].GetSegmentPath()] = &stpxsmstinstanceedittable.Stpxsmstinstanceeditentry[i]
    }
    return children
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetYangName() string { return "stpxSMSTInstanceEditTable" }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) SetParent(parent types.Entity) { stpxsmstinstanceedittable.parent = parent }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetParent() types.Entity { return stpxsmstinstanceedittable.parent }

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
// 
// The total number of entries in this table has to be 
// less than or equal to the object value of stpxSMSTMaxInstances.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance ID, the value of which is in the
    // range from 0 to stpxSMSTMaxInstanceID.   The instances of this table entry
    // with  stpxSMSTInstanceEditIndex of zero is automatically created by the
    // device and can not modified. The type is interface{} with range:
    // 0..4294967295.
    Stpxsmstinstanceeditindex interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 0 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 0 through 7; the second octet to VLANs 8 through 15;
    // etc.  The most significant bit of each octet corresponds to the lowest
    // value VlanIndex in that octet.  For each VLAN, if it is mapped to this MST
    // instance,  then the bit corresponding to that VLAN is set to  '1'. Each
    // VLAN can only be mapped to one unique MST  instance with the range from 0
    // to stpxSMSTMaxInstanceNumber. If the bit corresponding to a VLAN is changed
    // from '1'  to '0', then that VLAN will be automatically mapped to  SMST
    // instance 0 by the device. If the bit corresponding  to a VLAN is changed
    // from '0' to '1', then that VLAN will  be automatically removed from the MST
    // instance this VLAN was  previously mapped to. If the length of this string
    // is  less than 256 octets, any 'missing' octets are assumed to  contain the
    // value of zero. The type is string with length: 0..256.
    Stpxsmstinstanceeditvlansmap1K2K interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. Each VLAN can only be mapped to one unique MST instance with the range
    // from 0 to stpxSMSTMaxInstanceNumber. If the bit corresponding to a VLAN is
    // changed from '1' to '0', then that VLAN will be automatically mapped to
    // SMST instance 0 by the device. If the bit corresponding to a VLAN is
    // changed from '0' to '1', then that VLAN will be automatically removed from
    // the MST instance this VLAN was previously mapped to. If the length of this
    // string is  less than 256 octets, any 'missing' octets are assumed to 
    // contain the value of zero. The type is string with length: 0..256.
    Stpxsmstinstanceeditvlansmap3K4K interface{}

    // This object controls the creation and deletion of a row  in
    // stpxSMSTInstanceEditTable.  When creating an entry in this table,
    // 'createAndGo' method is used and the value of this object is set to
    // 'active'. Deactivation of an 'active' entry is not allowed.  When  deleting
    // an entry in this table, 'destroy' method is used.  Once a row becomes
    // active, value in any other column  within such a row may be modified. When
    // a row is active,  setting the instance of stpxSMSTInstanceEditVlansMap1k2k
    // stpxSMSTInstanceEditVlansMap3k4k for the same MST instance both to the
    // value of zero length can not be allowed. The type is RowStatus.
    Stpxsmstinstanceeditrowstatus interface{}
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetFilter() yfilter.YFilter { return stpxsmstinstanceeditentry.YFilter }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) SetFilter(yf yfilter.YFilter) { stpxsmstinstanceeditentry.YFilter = yf }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetGoName(yname string) string {
    if yname == "stpxSMSTInstanceEditIndex" { return "Stpxsmstinstanceeditindex" }
    if yname == "stpxSMSTInstanceEditVlansMap1k2k" { return "Stpxsmstinstanceeditvlansmap1K2K" }
    if yname == "stpxSMSTInstanceEditVlansMap3k4k" { return "Stpxsmstinstanceeditvlansmap3K4K" }
    if yname == "stpxSMSTInstanceEditRowStatus" { return "Stpxsmstinstanceeditrowstatus" }
    return ""
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetSegmentPath() string {
    return "stpxSMSTInstanceEditEntry" + "[stpxSMSTInstanceEditIndex='" + fmt.Sprintf("%v", stpxsmstinstanceeditentry.Stpxsmstinstanceeditindex) + "']"
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxSMSTInstanceEditIndex"] = stpxsmstinstanceeditentry.Stpxsmstinstanceeditindex
    leafs["stpxSMSTInstanceEditVlansMap1k2k"] = stpxsmstinstanceeditentry.Stpxsmstinstanceeditvlansmap1K2K
    leafs["stpxSMSTInstanceEditVlansMap3k4k"] = stpxsmstinstanceeditentry.Stpxsmstinstanceeditvlansmap3K4K
    leafs["stpxSMSTInstanceEditRowStatus"] = stpxsmstinstanceeditentry.Stpxsmstinstanceeditrowstatus
    return leafs
}

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetYangName() string { return "stpxSMSTInstanceEditEntry" }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) SetParent(parent types.Entity) { stpxsmstinstanceeditentry.parent = parent }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetParent() types.Entity { return stpxsmstinstanceeditentry.parent }

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetParentYangName() string { return "stpxSMSTInstanceEditTable" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
// 
// This table is only instantiated when the object 
// value of stpxSpanningTreeType is mst(4)
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry with port information for the MST protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry.
    Stpxsmstportentry []CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetFilter() yfilter.YFilter { return stpxsmstporttable.YFilter }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) SetFilter(yf yfilter.YFilter) { stpxsmstporttable.YFilter = yf }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetGoName(yname string) string {
    if yname == "stpxSMSTPortEntry" { return "Stpxsmstportentry" }
    return ""
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetSegmentPath() string {
    return "stpxSMSTPortTable"
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stpxSMSTPortEntry" {
        for _, c := range stpxsmstporttable.Stpxsmstportentry {
            if stpxsmstporttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry{}
        stpxsmstporttable.Stpxsmstportentry = append(stpxsmstporttable.Stpxsmstportentry, child)
        return &stpxsmstporttable.Stpxsmstportentry[len(stpxsmstporttable.Stpxsmstportentry)-1]
    }
    return nil
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stpxsmstporttable.Stpxsmstportentry {
        children[stpxsmstporttable.Stpxsmstportentry[i].GetSegmentPath()] = &stpxsmstporttable.Stpxsmstportentry[i]
    }
    return children
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetYangName() string { return "stpxSMSTPortTable" }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) SetParent(parent types.Entity) { stpxsmstporttable.parent = parent }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetParent() types.Entity { return stpxsmstporttable.parent }

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetParentYangName() string { return "CISCO-STP-EXTENSIONS-MIB" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry
// An entry with port information for the MST protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxsmstportindex interface{}

    // Indicates the operational status of the port for the  MST protocol.   edge
    // -- this port is an edge port for the MST region.  boundary -- this port is
    // a boundary port for the          MST region.  pvst --  this port is
    // connected to a PVST/PVST+ bridge.     stp -- this port is connected to a
    // Single Spanning         Tree bridge.  dispute -- this port, as a designated
    // port, received an         inferior BPDU with a designated role and the     
    // learning bit being set.  rstp -- this port is connected to a RSTP bridge or
    // an          MST bridge in a different MST region. The type is
    // map[string]bool.
    Stpxsmstportstatus interface{}

    // The adminitratively configured hello time in hundredth  of seconds on a
    // port for IEEE MST. The granularity  of this timer is 1 second. An agent may
    // return a badValue  error if a set is attempted to a value which is not a 
    // whole number of seconds. This object value of zero means the hello time is
    // not specifically configured on  this port and object value of
    // stpxSMSTPortConfigedHelloTime retrieved for this port will take on the
    // value of  dot1dStpBridgeHelloTime defined in BRIDGE-MIB. The type is
    // interface{} with range: 0..4294967295. Units are hundredth of seconds.
    Stpxsmstportadminhellotime interface{}

    // Indicates the effective configuration of the hello time on  a port. The
    // type is interface{} with range: 0..4294967295. Units are hundredth of
    // seconds.
    Stpxsmstportconfigedhellotime interface{}

    // The operational hello time in hundredth of seconds on a  port for IEEE MST.
    // If this object value is not applicable on a port, then the value retrieved
    // on that port will be -1. The type is interface{} with range:
    // -1..2147483647. Units are hundredth of seconds.
    Stpxsmstportoperhellotime interface{}

    // The desired MST mode of this port.  preStandard -- this port is
    // administratively configured to     transmit pre-standard, i.e. pre IEEE
    // MST, BPDUs.  auto -- the BPDU transmission mode of this port is based     
    // on automatic detection of neighbor ports. The type is
    // Stpxsmstportadminmstmode.
    Stpxsmstportadminmstmode interface{}

    // Indicates the current operational MST mode of this port.  unknown -- the
    // operational mode is currently unknown.  preStandard -- this port is
    // currently operating in      pre-standard MSTP BPDU transmission mode. 
    // standard -- this port is currently operating in IEEE MST      BPDU
    // transmission mode. The type is Stpxsmstportopermstmode.
    Stpxsmstportopermstmode interface{}
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetFilter() yfilter.YFilter { return stpxsmstportentry.YFilter }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) SetFilter(yf yfilter.YFilter) { stpxsmstportentry.YFilter = yf }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetGoName(yname string) string {
    if yname == "stpxSMSTPortIndex" { return "Stpxsmstportindex" }
    if yname == "stpxSMSTPortStatus" { return "Stpxsmstportstatus" }
    if yname == "stpxSMSTPortAdminHelloTime" { return "Stpxsmstportadminhellotime" }
    if yname == "stpxSMSTPortConfigedHelloTime" { return "Stpxsmstportconfigedhellotime" }
    if yname == "stpxSMSTPortOperHelloTime" { return "Stpxsmstportoperhellotime" }
    if yname == "stpxSMSTPortAdminMSTMode" { return "Stpxsmstportadminmstmode" }
    if yname == "stpxSMSTPortOperMSTMode" { return "Stpxsmstportopermstmode" }
    return ""
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetSegmentPath() string {
    return "stpxSMSTPortEntry" + "[stpxSMSTPortIndex='" + fmt.Sprintf("%v", stpxsmstportentry.Stpxsmstportindex) + "']"
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stpxSMSTPortIndex"] = stpxsmstportentry.Stpxsmstportindex
    leafs["stpxSMSTPortStatus"] = stpxsmstportentry.Stpxsmstportstatus
    leafs["stpxSMSTPortAdminHelloTime"] = stpxsmstportentry.Stpxsmstportadminhellotime
    leafs["stpxSMSTPortConfigedHelloTime"] = stpxsmstportentry.Stpxsmstportconfigedhellotime
    leafs["stpxSMSTPortOperHelloTime"] = stpxsmstportentry.Stpxsmstportoperhellotime
    leafs["stpxSMSTPortAdminMSTMode"] = stpxsmstportentry.Stpxsmstportadminmstmode
    leafs["stpxSMSTPortOperMSTMode"] = stpxsmstportentry.Stpxsmstportopermstmode
    return leafs
}

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetBundleName() string { return "cisco_ios_xe" }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetYangName() string { return "stpxSMSTPortEntry" }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) SetParent(parent types.Entity) { stpxsmstportentry.parent = parent }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetParent() types.Entity { return stpxsmstportentry.parent }

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetParentYangName() string { return "stpxSMSTPortTable" }

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode represents     on automatic detection of neighbor ports.
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode_preStandard CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode = "preStandard"

    CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode_auto CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportadminmstmode = "auto"
)

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode represents     BPDU transmission mode.
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode_unknown CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode = "unknown"

    CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode_preStandard CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode = "preStandard"

    CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode_standard CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry_Stpxsmstportopermstmode = "standard"
)

