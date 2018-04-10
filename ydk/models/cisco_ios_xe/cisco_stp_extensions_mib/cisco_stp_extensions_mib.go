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
    EntityData types.CommonEntityData
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

func (cISCOSTPEXTENSIONSMIB *CISCOSTPEXTENSIONSMIB) GetEntityData() *types.CommonEntityData {
    cISCOSTPEXTENSIONSMIB.EntityData.YFilter = cISCOSTPEXTENSIONSMIB.YFilter
    cISCOSTPEXTENSIONSMIB.EntityData.YangName = "CISCO-STP-EXTENSIONS-MIB"
    cISCOSTPEXTENSIONSMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOSTPEXTENSIONSMIB.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    cISCOSTPEXTENSIONSMIB.EntityData.SegmentPath = "CISCO-STP-EXTENSIONS-MIB:CISCO-STP-EXTENSIONS-MIB"
    cISCOSTPEXTENSIONSMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOSTPEXTENSIONSMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOSTPEXTENSIONSMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOSTPEXTENSIONSMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxUplinkFastObjects"] = types.YChild{"Stpxuplinkfastobjects", &cISCOSTPEXTENSIONSMIB.Stpxuplinkfastobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxBackboneFastObjects"] = types.YChild{"Stpxbackbonefastobjects", &cISCOSTPEXTENSIONSMIB.Stpxbackbonefastobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxSpanningTreeObjects"] = types.YChild{"Stpxspanningtreeobjects", &cISCOSTPEXTENSIONSMIB.Stpxspanningtreeobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMISTPObjects"] = types.YChild{"Stpxmistpobjects", &cISCOSTPEXTENSIONSMIB.Stpxmistpobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxLoopGuardObjects"] = types.YChild{"Stpxloopguardobjects", &cISCOSTPEXTENSIONSMIB.Stpxloopguardobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxFastStartObjects"] = types.YChild{"Stpxfaststartobjects", &cISCOSTPEXTENSIONSMIB.Stpxfaststartobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxBpduSkewingObjects"] = types.YChild{"Stpxbpduskewingobjects", &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMSTObjects"] = types.YChild{"Stpxmstobjects", &cISCOSTPEXTENSIONSMIB.Stpxmstobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRSTPObjects"] = types.YChild{"Stpxrstpobjects", &cISCOSTPEXTENSIONSMIB.Stpxrstpobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxSMSTObjects"] = types.YChild{"Stpxsmstobjects", &cISCOSTPEXTENSIONSMIB.Stpxsmstobjects}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxPVSTVlanTable"] = types.YChild{"Stpxpvstvlantable", &cISCOSTPEXTENSIONSMIB.Stpxpvstvlantable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxInconsistencyTable"] = types.YChild{"Stpxinconsistencytable", &cISCOSTPEXTENSIONSMIB.Stpxinconsistencytable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRootGuardConfigTable"] = types.YChild{"Stpxrootguardconfigtable", &cISCOSTPEXTENSIONSMIB.Stpxrootguardconfigtable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRootInconsistencyTable"] = types.YChild{"Stpxrootinconsistencytable", &cISCOSTPEXTENSIONSMIB.Stpxrootinconsistencytable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMISTPInstanceTable"] = types.YChild{"Stpxmistpinstancetable", &cISCOSTPEXTENSIONSMIB.Stpxmistpinstancetable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxLoopGuardConfigTable"] = types.YChild{"Stpxloopguardconfigtable", &cISCOSTPEXTENSIONSMIB.Stpxloopguardconfigtable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxLoopInconsistencyTable"] = types.YChild{"Stpxloopinconsistencytable", &cISCOSTPEXTENSIONSMIB.Stpxloopinconsistencytable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxFastStartPortTable"] = types.YChild{"Stpxfaststartporttable", &cISCOSTPEXTENSIONSMIB.Stpxfaststartporttable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxFastStartOperModeTable"] = types.YChild{"Stpxfaststartopermodetable", &cISCOSTPEXTENSIONSMIB.Stpxfaststartopermodetable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxBpduSkewingTable"] = types.YChild{"Stpxbpduskewingtable", &cISCOSTPEXTENSIONSMIB.Stpxbpduskewingtable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMSTInstanceTable"] = types.YChild{"Stpxmstinstancetable", &cISCOSTPEXTENSIONSMIB.Stpxmstinstancetable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMSTInstanceEditTable"] = types.YChild{"Stpxmstinstanceedittable", &cISCOSTPEXTENSIONSMIB.Stpxmstinstanceedittable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMSTPortTable"] = types.YChild{"Stpxmstporttable", &cISCOSTPEXTENSIONSMIB.Stpxmstporttable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxMSTPortRoleTable"] = types.YChild{"Stpxmstportroletable", &cISCOSTPEXTENSIONSMIB.Stpxmstportroletable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRSTPPortTable"] = types.YChild{"Stpxrstpporttable", &cISCOSTPEXTENSIONSMIB.Stpxrstpporttable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRSTPPortRoleTable"] = types.YChild{"Stpxrstpportroletable", &cISCOSTPEXTENSIONSMIB.Stpxrstpportroletable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxRPVSTPortTable"] = types.YChild{"Stpxrpvstporttable", &cISCOSTPEXTENSIONSMIB.Stpxrpvstporttable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxSMSTInstanceTable"] = types.YChild{"Stpxsmstinstancetable", &cISCOSTPEXTENSIONSMIB.Stpxsmstinstancetable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxSMSTInstanceEditTable"] = types.YChild{"Stpxsmstinstanceedittable", &cISCOSTPEXTENSIONSMIB.Stpxsmstinstanceedittable}
    cISCOSTPEXTENSIONSMIB.EntityData.Children["stpxSMSTPortTable"] = types.YChild{"Stpxsmstporttable", &cISCOSTPEXTENSIONSMIB.Stpxsmstporttable}
    cISCOSTPEXTENSIONSMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOSTPEXTENSIONSMIB.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects
type CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects struct {
    EntityData types.CommonEntityData
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

func (stpxuplinkfastobjects *CISCOSTPEXTENSIONSMIB_Stpxuplinkfastobjects) GetEntityData() *types.CommonEntityData {
    stpxuplinkfastobjects.EntityData.YFilter = stpxuplinkfastobjects.YFilter
    stpxuplinkfastobjects.EntityData.YangName = "stpxUplinkFastObjects"
    stpxuplinkfastobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxuplinkfastobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxuplinkfastobjects.EntityData.SegmentPath = "stpxUplinkFastObjects"
    stpxuplinkfastobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxuplinkfastobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxuplinkfastobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxuplinkfastobjects.EntityData.Children = make(map[string]types.YChild)
    stpxuplinkfastobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxuplinkfastobjects.EntityData.Leafs["stpxUplinkFastEnabled"] = types.YLeaf{"Stpxuplinkfastenabled", stpxuplinkfastobjects.Stpxuplinkfastenabled}
    stpxuplinkfastobjects.EntityData.Leafs["stpxUplinkFastTransitions"] = types.YLeaf{"Stpxuplinkfasttransitions", stpxuplinkfastobjects.Stpxuplinkfasttransitions}
    stpxuplinkfastobjects.EntityData.Leafs["stpxUplinkStationLearningGenRate"] = types.YLeaf{"Stpxuplinkstationlearninggenrate", stpxuplinkfastobjects.Stpxuplinkstationlearninggenrate}
    stpxuplinkfastobjects.EntityData.Leafs["stpxUplinkStationLearningFrames"] = types.YLeaf{"Stpxuplinkstationlearningframes", stpxuplinkfastobjects.Stpxuplinkstationlearningframes}
    stpxuplinkfastobjects.EntityData.Leafs["stpxUplinkFastOperEnabled"] = types.YLeaf{"Stpxuplinkfastoperenabled", stpxuplinkfastobjects.Stpxuplinkfastoperenabled}
    return &(stpxuplinkfastobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects
type CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects struct {
    EntityData types.CommonEntityData
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

func (stpxbackbonefastobjects *CISCOSTPEXTENSIONSMIB_Stpxbackbonefastobjects) GetEntityData() *types.CommonEntityData {
    stpxbackbonefastobjects.EntityData.YFilter = stpxbackbonefastobjects.YFilter
    stpxbackbonefastobjects.EntityData.YangName = "stpxBackboneFastObjects"
    stpxbackbonefastobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxbackbonefastobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxbackbonefastobjects.EntityData.SegmentPath = "stpxBackboneFastObjects"
    stpxbackbonefastobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxbackbonefastobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxbackbonefastobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxbackbonefastobjects.EntityData.Children = make(map[string]types.YChild)
    stpxbackbonefastobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastEnabled"] = types.YLeaf{"Stpxbackbonefastenabled", stpxbackbonefastobjects.Stpxbackbonefastenabled}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastInInferiorBPDUs"] = types.YLeaf{"Stpxbackbonefastininferiorbpdus", stpxbackbonefastobjects.Stpxbackbonefastininferiorbpdus}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastInRLQRequestPDUs"] = types.YLeaf{"Stpxbackbonefastinrlqrequestpdus", stpxbackbonefastobjects.Stpxbackbonefastinrlqrequestpdus}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastInRLQResponsePDUs"] = types.YLeaf{"Stpxbackbonefastinrlqresponsepdus", stpxbackbonefastobjects.Stpxbackbonefastinrlqresponsepdus}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastOutRLQRequestPDUs"] = types.YLeaf{"Stpxbackbonefastoutrlqrequestpdus", stpxbackbonefastobjects.Stpxbackbonefastoutrlqrequestpdus}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastOutRLQResponsePDUs"] = types.YLeaf{"Stpxbackbonefastoutrlqresponsepdus", stpxbackbonefastobjects.Stpxbackbonefastoutrlqresponsepdus}
    stpxbackbonefastobjects.EntityData.Leafs["stpxBackboneFastOperEnabled"] = types.YLeaf{"Stpxbackbonefastoperenabled", stpxbackbonefastobjects.Stpxbackbonefastoperenabled}
    return &(stpxbackbonefastobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects
type CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects struct {
    EntityData types.CommonEntityData
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

func (stpxspanningtreeobjects *CISCOSTPEXTENSIONSMIB_Stpxspanningtreeobjects) GetEntityData() *types.CommonEntityData {
    stpxspanningtreeobjects.EntityData.YFilter = stpxspanningtreeobjects.YFilter
    stpxspanningtreeobjects.EntityData.YangName = "stpxSpanningTreeObjects"
    stpxspanningtreeobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxspanningtreeobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxspanningtreeobjects.EntityData.SegmentPath = "stpxSpanningTreeObjects"
    stpxspanningtreeobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxspanningtreeobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxspanningtreeobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxspanningtreeobjects.EntityData.Children = make(map[string]types.YChild)
    stpxspanningtreeobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxspanningtreeobjects.EntityData.Leafs["stpxSpanningTreeType"] = types.YLeaf{"Stpxspanningtreetype", stpxspanningtreeobjects.Stpxspanningtreetype}
    stpxspanningtreeobjects.EntityData.Leafs["stpxSpanningTreePathCostMode"] = types.YLeaf{"Stpxspanningtreepathcostmode", stpxspanningtreeobjects.Stpxspanningtreepathcostmode}
    stpxspanningtreeobjects.EntityData.Leafs["stpxExtendedSysIDAdminEnabled"] = types.YLeaf{"Stpxextendedsysidadminenabled", stpxspanningtreeobjects.Stpxextendedsysidadminenabled}
    stpxspanningtreeobjects.EntityData.Leafs["stpxExtendedSysIDOperEnabled"] = types.YLeaf{"Stpxextendedsysidoperenabled", stpxspanningtreeobjects.Stpxextendedsysidoperenabled}
    stpxspanningtreeobjects.EntityData.Leafs["stpxNotificationEnable"] = types.YLeaf{"Stpxnotificationenable", stpxspanningtreeobjects.Stpxnotificationenable}
    stpxspanningtreeobjects.EntityData.Leafs["stpxSpanningTreePathCostOperMode"] = types.YLeaf{"Stpxspanningtreepathcostopermode", stpxspanningtreeobjects.Stpxspanningtreepathcostopermode}
    return &(stpxspanningtreeobjects.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of MISTP instances, that are supported by the device  when the
    // value of stpxSpanningTreeType is either mistp(2) or mistpPvstPlus(3). The
    // type is interface{} with range: 1..256.
    Stpxmistpinstancenumber interface{}
}

func (stpxmistpobjects *CISCOSTPEXTENSIONSMIB_Stpxmistpobjects) GetEntityData() *types.CommonEntityData {
    stpxmistpobjects.EntityData.YFilter = stpxmistpobjects.YFilter
    stpxmistpobjects.EntityData.YangName = "stpxMISTPObjects"
    stpxmistpobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxmistpobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmistpobjects.EntityData.SegmentPath = "stpxMISTPObjects"
    stpxmistpobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmistpobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmistpobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmistpobjects.EntityData.Children = make(map[string]types.YChild)
    stpxmistpobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmistpobjects.EntityData.Leafs["stpxMISTPInstanceNumber"] = types.YLeaf{"Stpxmistpinstancenumber", stpxmistpobjects.Stpxmistpinstancenumber}
    return &(stpxmistpobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects
type CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the global default config mode of LoopGuard  feature on the
    // device. The type is Stpxloopguardglobaldefaultmode.
    Stpxloopguardglobaldefaultmode interface{}
}

func (stpxloopguardobjects *CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects) GetEntityData() *types.CommonEntityData {
    stpxloopguardobjects.EntityData.YFilter = stpxloopguardobjects.YFilter
    stpxloopguardobjects.EntityData.YangName = "stpxLoopGuardObjects"
    stpxloopguardobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxloopguardobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxloopguardobjects.EntityData.SegmentPath = "stpxLoopGuardObjects"
    stpxloopguardobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxloopguardobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxloopguardobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxloopguardobjects.EntityData.Children = make(map[string]types.YChild)
    stpxloopguardobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxloopguardobjects.EntityData.Leafs["stpxLoopGuardGlobalDefaultMode"] = types.YLeaf{"Stpxloopguardglobaldefaultmode", stpxloopguardobjects.Stpxloopguardglobaldefaultmode}
    return &(stpxloopguardobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode represents feature on the device.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode_enable CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode_disable CISCOSTPEXTENSIONSMIB_Stpxloopguardobjects_Stpxloopguardglobaldefaultmode = "disable"
)

// CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects
type CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects struct {
    EntityData types.CommonEntityData
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

func (stpxfaststartobjects *CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects) GetEntityData() *types.CommonEntityData {
    stpxfaststartobjects.EntityData.YFilter = stpxfaststartobjects.YFilter
    stpxfaststartobjects.EntityData.YangName = "stpxFastStartObjects"
    stpxfaststartobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxfaststartobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxfaststartobjects.EntityData.SegmentPath = "stpxFastStartObjects"
    stpxfaststartobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxfaststartobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxfaststartobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxfaststartobjects.EntityData.Children = make(map[string]types.YChild)
    stpxfaststartobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxfaststartobjects.EntityData.Leafs["stpxFastStartBpduGuardEnable"] = types.YLeaf{"Stpxfaststartbpduguardenable", stpxfaststartobjects.Stpxfaststartbpduguardenable}
    stpxfaststartobjects.EntityData.Leafs["stpxFastStartBpduFilterEnable"] = types.YLeaf{"Stpxfaststartbpdufilterenable", stpxfaststartobjects.Stpxfaststartbpdufilterenable}
    stpxfaststartobjects.EntityData.Leafs["stpxFastStartGlobalDefaultMode"] = types.YLeaf{"Stpxfaststartglobaldefaultmode", stpxfaststartobjects.Stpxfaststartglobaldefaultmode}
    return &(stpxfaststartobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode represents Start feature on the device.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode string

const (
    CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode_enable CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode = "enable"

    CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode_disable CISCOSTPEXTENSIONSMIB_Stpxfaststartobjects_Stpxfaststartglobaldefaultmode = "disable"
)

// CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects
type CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether BPDU skewing detection feature is enabled or not on the
    // system. If this object has the value of true(1), then the system will
    // detect whether BPDUs received by any port on any Spanning  Tree instance
    // are processed at an interval longer than the object value of
    // dot1dStpHelloTime in the BIRDGE-MIB of the Spanning Tree instance. The type
    // is bool.
    Stpxbpduskewingdetectionenable interface{}
}

func (stpxbpduskewingobjects *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingobjects) GetEntityData() *types.CommonEntityData {
    stpxbpduskewingobjects.EntityData.YFilter = stpxbpduskewingobjects.YFilter
    stpxbpduskewingobjects.EntityData.YangName = "stpxBpduSkewingObjects"
    stpxbpduskewingobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxbpduskewingobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxbpduskewingobjects.EntityData.SegmentPath = "stpxBpduSkewingObjects"
    stpxbpduskewingobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxbpduskewingobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxbpduskewingobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxbpduskewingobjects.EntityData.Children = make(map[string]types.YChild)
    stpxbpduskewingobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxbpduskewingobjects.EntityData.Leafs["stpxBpduSkewingDetectionEnable"] = types.YLeaf{"Stpxbpduskewingdetectionenable", stpxbpduskewingobjects.Stpxbpduskewingdetectionenable}
    return &(stpxbpduskewingobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstobjects
type CISCOSTPEXTENSIONSMIB_Stpxmstobjects struct {
    EntityData types.CommonEntityData
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

func (stpxmstobjects *CISCOSTPEXTENSIONSMIB_Stpxmstobjects) GetEntityData() *types.CommonEntityData {
    stpxmstobjects.EntityData.YFilter = stpxmstobjects.YFilter
    stpxmstobjects.EntityData.YangName = "stpxMSTObjects"
    stpxmstobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxmstobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmstobjects.EntityData.SegmentPath = "stpxMSTObjects"
    stpxmstobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstobjects.EntityData.Children = make(map[string]types.YChild)
    stpxmstobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmstobjects.EntityData.Leafs["stpxMSTMaxInstanceNumber"] = types.YLeaf{"Stpxmstmaxinstancenumber", stpxmstobjects.Stpxmstmaxinstancenumber}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionName"] = types.YLeaf{"Stpxmstregionname", stpxmstobjects.Stpxmstregionname}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionRevision"] = types.YLeaf{"Stpxmstregionrevision", stpxmstobjects.Stpxmstregionrevision}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionEditBufferStatus"] = types.YLeaf{"Stpxmstregioneditbufferstatus", stpxmstobjects.Stpxmstregioneditbufferstatus}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionEditBufferOperation"] = types.YLeaf{"Stpxmstregioneditbufferoperation", stpxmstobjects.Stpxmstregioneditbufferoperation}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionEditName"] = types.YLeaf{"Stpxmstregioneditname", stpxmstobjects.Stpxmstregioneditname}
    stpxmstobjects.EntityData.Leafs["stpxMSTRegionEditRevision"] = types.YLeaf{"Stpxmstregioneditrevision", stpxmstobjects.Stpxmstregioneditrevision}
    stpxmstobjects.EntityData.Leafs["stpxMSTMaxHopCount"] = types.YLeaf{"Stpxmstmaxhopcount", stpxmstobjects.Stpxmstmaxhopcount}
    return &(stpxmstobjects.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Transmit Hold Count. The type is interface{} with range: 0..4294967295.
    Stpxrstptransmitholdcount interface{}
}

func (stpxrstpobjects *CISCOSTPEXTENSIONSMIB_Stpxrstpobjects) GetEntityData() *types.CommonEntityData {
    stpxrstpobjects.EntityData.YFilter = stpxrstpobjects.YFilter
    stpxrstpobjects.EntityData.YangName = "stpxRSTPObjects"
    stpxrstpobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxrstpobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrstpobjects.EntityData.SegmentPath = "stpxRSTPObjects"
    stpxrstpobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrstpobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrstpobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrstpobjects.EntityData.Children = make(map[string]types.YChild)
    stpxrstpobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrstpobjects.EntityData.Leafs["stpxRSTPTransmitHoldCount"] = types.YLeaf{"Stpxrstptransmitholdcount", stpxrstpobjects.Stpxrstptransmitholdcount}
    return &(stpxrstpobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstobjects
type CISCOSTPEXTENSIONSMIB_Stpxsmstobjects struct {
    EntityData types.CommonEntityData
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

func (stpxsmstobjects *CISCOSTPEXTENSIONSMIB_Stpxsmstobjects) GetEntityData() *types.CommonEntityData {
    stpxsmstobjects.EntityData.YFilter = stpxsmstobjects.YFilter
    stpxsmstobjects.EntityData.YangName = "stpxSMSTObjects"
    stpxsmstobjects.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstobjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxsmstobjects.EntityData.SegmentPath = "stpxSMSTObjects"
    stpxsmstobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstobjects.EntityData.Children = make(map[string]types.YChild)
    stpxsmstobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxsmstobjects.EntityData.Leafs["stpxSMSTMaxInstances"] = types.YLeaf{"Stpxsmstmaxinstances", stpxsmstobjects.Stpxsmstmaxinstances}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTMaxInstanceID"] = types.YLeaf{"Stpxsmstmaxinstanceid", stpxsmstobjects.Stpxsmstmaxinstanceid}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTRegionRevision"] = types.YLeaf{"Stpxsmstregionrevision", stpxsmstobjects.Stpxsmstregionrevision}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTRegionEditRevision"] = types.YLeaf{"Stpxsmstregioneditrevision", stpxsmstobjects.Stpxsmstregioneditrevision}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTMaxHopCount"] = types.YLeaf{"Stpxsmstmaxhopcount", stpxsmstobjects.Stpxsmstmaxhopcount}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTConfigDigest"] = types.YLeaf{"Stpxsmstconfigdigest", stpxsmstobjects.Stpxsmstconfigdigest}
    stpxsmstobjects.EntityData.Leafs["stpxSMSTConfigPreStandardDigest"] = types.YLeaf{"Stpxsmstconfigprestandarddigest", stpxsmstobjects.Stpxsmstconfigprestandarddigest}
    return &(stpxsmstobjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable
// A list of Virtual LAN entries containing
// information for Spanning Tree PVST+ protocol. 
// An entry will exist for each VLAN existing on 
// the device.
type CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing Spanning Tree PVST+ Protocol  information for a
    // particular Virtual LAN. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry.
    Stpxpvstvlanentry []CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry
}

func (stpxpvstvlantable *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable) GetEntityData() *types.CommonEntityData {
    stpxpvstvlantable.EntityData.YFilter = stpxpvstvlantable.YFilter
    stpxpvstvlantable.EntityData.YangName = "stpxPVSTVlanTable"
    stpxpvstvlantable.EntityData.BundleName = "cisco_ios_xe"
    stpxpvstvlantable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxpvstvlantable.EntityData.SegmentPath = "stpxPVSTVlanTable"
    stpxpvstvlantable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxpvstvlantable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxpvstvlantable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxpvstvlantable.EntityData.Children = make(map[string]types.YChild)
    stpxpvstvlantable.EntityData.Children["stpxPVSTVlanEntry"] = types.YChild{"Stpxpvstvlanentry", nil}
    for i := range stpxpvstvlantable.Stpxpvstvlanentry {
        stpxpvstvlantable.EntityData.Children[types.GetSegmentPath(&stpxpvstvlantable.Stpxpvstvlanentry[i])] = types.YChild{"Stpxpvstvlanentry", &stpxpvstvlantable.Stpxpvstvlanentry[i]}
    }
    stpxpvstvlantable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxpvstvlantable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry
// An entry containing Spanning Tree PVST+ Protocol 
// information for a particular Virtual LAN.
type CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry struct {
    EntityData types.CommonEntityData
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

func (stpxpvstvlanentry *CISCOSTPEXTENSIONSMIB_Stpxpvstvlantable_Stpxpvstvlanentry) GetEntityData() *types.CommonEntityData {
    stpxpvstvlanentry.EntityData.YFilter = stpxpvstvlanentry.YFilter
    stpxpvstvlanentry.EntityData.YangName = "stpxPVSTVlanEntry"
    stpxpvstvlanentry.EntityData.BundleName = "cisco_ios_xe"
    stpxpvstvlanentry.EntityData.ParentYangName = "stpxPVSTVlanTable"
    stpxpvstvlanentry.EntityData.SegmentPath = "stpxPVSTVlanEntry" + "[stpxPVSTVlanIndex='" + fmt.Sprintf("%v", stpxpvstvlanentry.Stpxpvstvlanindex) + "']"
    stpxpvstvlanentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxpvstvlanentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxpvstvlanentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxpvstvlanentry.EntityData.Children = make(map[string]types.YChild)
    stpxpvstvlanentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxpvstvlanentry.EntityData.Leafs["stpxPVSTVlanIndex"] = types.YLeaf{"Stpxpvstvlanindex", stpxpvstvlanentry.Stpxpvstvlanindex}
    stpxpvstvlanentry.EntityData.Leafs["stpxPVSTVlanEnable"] = types.YLeaf{"Stpxpvstvlanenable", stpxpvstvlanentry.Stpxpvstvlanenable}
    return &(stpxpvstvlanentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VLAN on a particular port for which a Spanning Tree inconsistency is
    // currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry.
    Stpxinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry
}

func (stpxinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable) GetEntityData() *types.CommonEntityData {
    stpxinconsistencytable.EntityData.YFilter = stpxinconsistencytable.YFilter
    stpxinconsistencytable.EntityData.YangName = "stpxInconsistencyTable"
    stpxinconsistencytable.EntityData.BundleName = "cisco_ios_xe"
    stpxinconsistencytable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxinconsistencytable.EntityData.SegmentPath = "stpxInconsistencyTable"
    stpxinconsistencytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxinconsistencytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxinconsistencytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxinconsistencytable.EntityData.Children = make(map[string]types.YChild)
    stpxinconsistencytable.EntityData.Children["stpxInconsistencyEntry"] = types.YChild{"Stpxinconsistencyentry", nil}
    for i := range stpxinconsistencytable.Stpxinconsistencyentry {
        stpxinconsistencytable.EntityData.Children[types.GetSegmentPath(&stpxinconsistencytable.Stpxinconsistencyentry[i])] = types.YChild{"Stpxinconsistencyentry", &stpxinconsistencytable.Stpxinconsistencyentry[i]}
    }
    stpxinconsistencytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxinconsistencytable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry
// A VLAN on a particular port for which a Spanning Tree
// inconsistency is currently in effect.
type CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry struct {
    EntityData types.CommonEntityData
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

func (stpxinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxinconsistencytable_Stpxinconsistencyentry) GetEntityData() *types.CommonEntityData {
    stpxinconsistencyentry.EntityData.YFilter = stpxinconsistencyentry.YFilter
    stpxinconsistencyentry.EntityData.YangName = "stpxInconsistencyEntry"
    stpxinconsistencyentry.EntityData.BundleName = "cisco_ios_xe"
    stpxinconsistencyentry.EntityData.ParentYangName = "stpxInconsistencyTable"
    stpxinconsistencyentry.EntityData.SegmentPath = "stpxInconsistencyEntry" + "[stpxVlanIndex='" + fmt.Sprintf("%v", stpxinconsistencyentry.Stpxvlanindex) + "']" + "[stpxPortIndex='" + fmt.Sprintf("%v", stpxinconsistencyentry.Stpxportindex) + "']"
    stpxinconsistencyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxinconsistencyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxinconsistencyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxinconsistencyentry.EntityData.Children = make(map[string]types.YChild)
    stpxinconsistencyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxinconsistencyentry.EntityData.Leafs["stpxVlanIndex"] = types.YLeaf{"Stpxvlanindex", stpxinconsistencyentry.Stpxvlanindex}
    stpxinconsistencyentry.EntityData.Leafs["stpxPortIndex"] = types.YLeaf{"Stpxportindex", stpxinconsistencyentry.Stpxportindex}
    stpxinconsistencyentry.EntityData.Leafs["stpxInconsistentState"] = types.YLeaf{"Stpxinconsistentstate", stpxinconsistencyentry.Stpxinconsistentstate}
    return &(stpxinconsistencyentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable
// A table containing a list of the bridge ports for which
// Spanning Tree RootGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree RootGuard capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry.
    Stpxrootguardconfigentry []CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry
}

func (stpxrootguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable) GetEntityData() *types.CommonEntityData {
    stpxrootguardconfigtable.EntityData.YFilter = stpxrootguardconfigtable.YFilter
    stpxrootguardconfigtable.EntityData.YangName = "stpxRootGuardConfigTable"
    stpxrootguardconfigtable.EntityData.BundleName = "cisco_ios_xe"
    stpxrootguardconfigtable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrootguardconfigtable.EntityData.SegmentPath = "stpxRootGuardConfigTable"
    stpxrootguardconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrootguardconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrootguardconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrootguardconfigtable.EntityData.Children = make(map[string]types.YChild)
    stpxrootguardconfigtable.EntityData.Children["stpxRootGuardConfigEntry"] = types.YChild{"Stpxrootguardconfigentry", nil}
    for i := range stpxrootguardconfigtable.Stpxrootguardconfigentry {
        stpxrootguardconfigtable.EntityData.Children[types.GetSegmentPath(&stpxrootguardconfigtable.Stpxrootguardconfigentry[i])] = types.YChild{"Stpxrootguardconfigentry", &stpxrootguardconfigtable.Stpxrootguardconfigentry[i]}
    }
    stpxrootguardconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxrootguardconfigtable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry
// A bridge port for which Spanning Tree RootGuard
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    Stpxrootguardconfigportindex interface{}

    // An indication of whether the RootGuard capability is  enabled on this port
    // or not. This configuration will be applied to all Spanning Tree instances
    // in which this port  exists. The type is bool.
    Stpxrootguardconfigenabled interface{}
}

func (stpxrootguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxrootguardconfigtable_Stpxrootguardconfigentry) GetEntityData() *types.CommonEntityData {
    stpxrootguardconfigentry.EntityData.YFilter = stpxrootguardconfigentry.YFilter
    stpxrootguardconfigentry.EntityData.YangName = "stpxRootGuardConfigEntry"
    stpxrootguardconfigentry.EntityData.BundleName = "cisco_ios_xe"
    stpxrootguardconfigentry.EntityData.ParentYangName = "stpxRootGuardConfigTable"
    stpxrootguardconfigentry.EntityData.SegmentPath = "stpxRootGuardConfigEntry" + "[stpxRootGuardConfigPortIndex='" + fmt.Sprintf("%v", stpxrootguardconfigentry.Stpxrootguardconfigportindex) + "']"
    stpxrootguardconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrootguardconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrootguardconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrootguardconfigentry.EntityData.Children = make(map[string]types.YChild)
    stpxrootguardconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrootguardconfigentry.EntityData.Leafs["stpxRootGuardConfigPortIndex"] = types.YLeaf{"Stpxrootguardconfigportindex", stpxrootguardconfigentry.Stpxrootguardconfigportindex}
    stpxrootguardconfigentry.EntityData.Leafs["stpxRootGuardConfigEnabled"] = types.YLeaf{"Stpxrootguardconfigenabled", stpxrootguardconfigentry.Stpxrootguardconfigenabled}
    return &(stpxrootguardconfigentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable
// A table containing a list of the bridge ports for which
// a particular Spanning Tree instance has been found 
// to have an root-inconsistency. The agent creates a new 
// entry in this table whenever it detects a new 
// root-inconsistency, and deletes entries 
// when/soon after the inconsistency is no longer present.
type CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for  which a Spanning Tree
    // root-inconsistency is currently  in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry.
    Stpxrootinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry
}

func (stpxrootinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable) GetEntityData() *types.CommonEntityData {
    stpxrootinconsistencytable.EntityData.YFilter = stpxrootinconsistencytable.YFilter
    stpxrootinconsistencytable.EntityData.YangName = "stpxRootInconsistencyTable"
    stpxrootinconsistencytable.EntityData.BundleName = "cisco_ios_xe"
    stpxrootinconsistencytable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrootinconsistencytable.EntityData.SegmentPath = "stpxRootInconsistencyTable"
    stpxrootinconsistencytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrootinconsistencytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrootinconsistencytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrootinconsistencytable.EntityData.Children = make(map[string]types.YChild)
    stpxrootinconsistencytable.EntityData.Children["stpxRootInconsistencyEntry"] = types.YChild{"Stpxrootinconsistencyentry", nil}
    for i := range stpxrootinconsistencytable.Stpxrootinconsistencyentry {
        stpxrootinconsistencytable.EntityData.Children[types.GetSegmentPath(&stpxrootinconsistencytable.Stpxrootinconsistencyentry[i])] = types.YChild{"Stpxrootinconsistencyentry", &stpxrootinconsistencytable.Stpxrootinconsistencyentry[i]}
    }
    stpxrootinconsistencytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxrootinconsistencytable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry
// A Spanning Tree instance on a particular port for 
// which a Spanning Tree root-inconsistency is currently 
// in effect.
type CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry struct {
    EntityData types.CommonEntityData
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

func (stpxrootinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxrootinconsistencytable_Stpxrootinconsistencyentry) GetEntityData() *types.CommonEntityData {
    stpxrootinconsistencyentry.EntityData.YFilter = stpxrootinconsistencyentry.YFilter
    stpxrootinconsistencyentry.EntityData.YangName = "stpxRootInconsistencyEntry"
    stpxrootinconsistencyentry.EntityData.BundleName = "cisco_ios_xe"
    stpxrootinconsistencyentry.EntityData.ParentYangName = "stpxRootInconsistencyTable"
    stpxrootinconsistencyentry.EntityData.SegmentPath = "stpxRootInconsistencyEntry" + "[stpxRootInconsistencyIndex='" + fmt.Sprintf("%v", stpxrootinconsistencyentry.Stpxrootinconsistencyindex) + "']" + "[stpxRootInconsistencyPortIndex='" + fmt.Sprintf("%v", stpxrootinconsistencyentry.Stpxrootinconsistencyportindex) + "']"
    stpxrootinconsistencyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrootinconsistencyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrootinconsistencyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrootinconsistencyentry.EntityData.Children = make(map[string]types.YChild)
    stpxrootinconsistencyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrootinconsistencyentry.EntityData.Leafs["stpxRootInconsistencyIndex"] = types.YLeaf{"Stpxrootinconsistencyindex", stpxrootinconsistencyentry.Stpxrootinconsistencyindex}
    stpxrootinconsistencyentry.EntityData.Leafs["stpxRootInconsistencyPortIndex"] = types.YLeaf{"Stpxrootinconsistencyportindex", stpxrootinconsistencyentry.Stpxrootinconsistencyportindex}
    stpxrootinconsistencyentry.EntityData.Leafs["stpxRootInconsistencyState"] = types.YLeaf{"Stpxrootinconsistencystate", stpxrootinconsistencyentry.Stpxrootinconsistencystate}
    return &(stpxrootinconsistencyentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable
// This table contains one entry for each instance of MISTP and 
// it contains stpxMISTPInstanceNumber entries, numbered from 1
// to stpxMISTPInstanceNumber.
// 
// This table is only instantiated when the value of 
// stpxSpanningTreeType is mistp(2) or mistpPvstPlus(3).
type CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the status of the MISTP  instance. The type is
    // slice of
    // CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry.
    Stpxmistpinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry
}

func (stpxmistpinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable) GetEntityData() *types.CommonEntityData {
    stpxmistpinstancetable.EntityData.YFilter = stpxmistpinstancetable.YFilter
    stpxmistpinstancetable.EntityData.YangName = "stpxMISTPInstanceTable"
    stpxmistpinstancetable.EntityData.BundleName = "cisco_ios_xe"
    stpxmistpinstancetable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmistpinstancetable.EntityData.SegmentPath = "stpxMISTPInstanceTable"
    stpxmistpinstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmistpinstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmistpinstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmistpinstancetable.EntityData.Children = make(map[string]types.YChild)
    stpxmistpinstancetable.EntityData.Children["stpxMISTPInstanceEntry"] = types.YChild{"Stpxmistpinstanceentry", nil}
    for i := range stpxmistpinstancetable.Stpxmistpinstanceentry {
        stpxmistpinstancetable.EntityData.Children[types.GetSegmentPath(&stpxmistpinstancetable.Stpxmistpinstanceentry[i])] = types.YChild{"Stpxmistpinstanceentry", &stpxmistpinstancetable.Stpxmistpinstanceentry[i]}
    }
    stpxmistpinstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxmistpinstancetable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry
// A conceptual row containing the status of the MISTP 
// instance.
type CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry struct {
    EntityData types.CommonEntityData
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

func (stpxmistpinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmistpinstancetable_Stpxmistpinstanceentry) GetEntityData() *types.CommonEntityData {
    stpxmistpinstanceentry.EntityData.YFilter = stpxmistpinstanceentry.YFilter
    stpxmistpinstanceentry.EntityData.YangName = "stpxMISTPInstanceEntry"
    stpxmistpinstanceentry.EntityData.BundleName = "cisco_ios_xe"
    stpxmistpinstanceentry.EntityData.ParentYangName = "stpxMISTPInstanceTable"
    stpxmistpinstanceentry.EntityData.SegmentPath = "stpxMISTPInstanceEntry" + "[stpxMISTPInstanceIndex='" + fmt.Sprintf("%v", stpxmistpinstanceentry.Stpxmistpinstanceindex) + "']"
    stpxmistpinstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmistpinstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmistpinstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmistpinstanceentry.EntityData.Children = make(map[string]types.YChild)
    stpxmistpinstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceIndex"] = types.YLeaf{"Stpxmistpinstanceindex", stpxmistpinstanceentry.Stpxmistpinstanceindex}
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceEnable"] = types.YLeaf{"Stpxmistpinstanceenable", stpxmistpinstanceentry.Stpxmistpinstanceenable}
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceVlansMapped"] = types.YLeaf{"Stpxmistpinstancevlansmapped", stpxmistpinstanceentry.Stpxmistpinstancevlansmapped}
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceVlansMapped2k"] = types.YLeaf{"Stpxmistpinstancevlansmapped2K", stpxmistpinstanceentry.Stpxmistpinstancevlansmapped2K}
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceVlansMapped3k"] = types.YLeaf{"Stpxmistpinstancevlansmapped3K", stpxmistpinstanceentry.Stpxmistpinstancevlansmapped3K}
    stpxmistpinstanceentry.EntityData.Leafs["stpxMISTPInstanceVlansMapped4k"] = types.YLeaf{"Stpxmistpinstancevlansmapped4K", stpxmistpinstanceentry.Stpxmistpinstancevlansmapped4K}
    return &(stpxmistpinstanceentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable
// A table containing a list of the bridge ports for which
// Spanning Tree LoopGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree LoopGuard  capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry.
    Stpxloopguardconfigentry []CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry
}

func (stpxloopguardconfigtable *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable) GetEntityData() *types.CommonEntityData {
    stpxloopguardconfigtable.EntityData.YFilter = stpxloopguardconfigtable.YFilter
    stpxloopguardconfigtable.EntityData.YangName = "stpxLoopGuardConfigTable"
    stpxloopguardconfigtable.EntityData.BundleName = "cisco_ios_xe"
    stpxloopguardconfigtable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxloopguardconfigtable.EntityData.SegmentPath = "stpxLoopGuardConfigTable"
    stpxloopguardconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxloopguardconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxloopguardconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxloopguardconfigtable.EntityData.Children = make(map[string]types.YChild)
    stpxloopguardconfigtable.EntityData.Children["stpxLoopGuardConfigEntry"] = types.YChild{"Stpxloopguardconfigentry", nil}
    for i := range stpxloopguardconfigtable.Stpxloopguardconfigentry {
        stpxloopguardconfigtable.EntityData.Children[types.GetSegmentPath(&stpxloopguardconfigtable.Stpxloopguardconfigentry[i])] = types.YChild{"Stpxloopguardconfigentry", &stpxloopguardconfigtable.Stpxloopguardconfigentry[i]}
    }
    stpxloopguardconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxloopguardconfigtable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry
// A bridge port for which Spanning Tree LoopGuard 
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry struct {
    EntityData types.CommonEntityData
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

func (stpxloopguardconfigentry *CISCOSTPEXTENSIONSMIB_Stpxloopguardconfigtable_Stpxloopguardconfigentry) GetEntityData() *types.CommonEntityData {
    stpxloopguardconfigentry.EntityData.YFilter = stpxloopguardconfigentry.YFilter
    stpxloopguardconfigentry.EntityData.YangName = "stpxLoopGuardConfigEntry"
    stpxloopguardconfigentry.EntityData.BundleName = "cisco_ios_xe"
    stpxloopguardconfigentry.EntityData.ParentYangName = "stpxLoopGuardConfigTable"
    stpxloopguardconfigentry.EntityData.SegmentPath = "stpxLoopGuardConfigEntry" + "[stpxLoopGuardConfigPortIndex='" + fmt.Sprintf("%v", stpxloopguardconfigentry.Stpxloopguardconfigportindex) + "']"
    stpxloopguardconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxloopguardconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxloopguardconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxloopguardconfigentry.EntityData.Children = make(map[string]types.YChild)
    stpxloopguardconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxloopguardconfigentry.EntityData.Leafs["stpxLoopGuardConfigPortIndex"] = types.YLeaf{"Stpxloopguardconfigportindex", stpxloopguardconfigentry.Stpxloopguardconfigportindex}
    stpxloopguardconfigentry.EntityData.Leafs["stpxLoopGuardConfigEnabled"] = types.YLeaf{"Stpxloopguardconfigenabled", stpxloopguardconfigentry.Stpxloopguardconfigenabled}
    stpxloopguardconfigentry.EntityData.Leafs["stpxLoopGuardConfigMode"] = types.YLeaf{"Stpxloopguardconfigmode", stpxloopguardconfigentry.Stpxloopguardconfigmode}
    return &(stpxloopguardconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which a Spanning Tree
    // loop-inconsistency is currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry.
    Stpxloopinconsistencyentry []CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry
}

func (stpxloopinconsistencytable *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable) GetEntityData() *types.CommonEntityData {
    stpxloopinconsistencytable.EntityData.YFilter = stpxloopinconsistencytable.YFilter
    stpxloopinconsistencytable.EntityData.YangName = "stpxLoopInconsistencyTable"
    stpxloopinconsistencytable.EntityData.BundleName = "cisco_ios_xe"
    stpxloopinconsistencytable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxloopinconsistencytable.EntityData.SegmentPath = "stpxLoopInconsistencyTable"
    stpxloopinconsistencytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxloopinconsistencytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxloopinconsistencytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxloopinconsistencytable.EntityData.Children = make(map[string]types.YChild)
    stpxloopinconsistencytable.EntityData.Children["stpxLoopInconsistencyEntry"] = types.YChild{"Stpxloopinconsistencyentry", nil}
    for i := range stpxloopinconsistencytable.Stpxloopinconsistencyentry {
        stpxloopinconsistencytable.EntityData.Children[types.GetSegmentPath(&stpxloopinconsistencytable.Stpxloopinconsistencyentry[i])] = types.YChild{"Stpxloopinconsistencyentry", &stpxloopinconsistencytable.Stpxloopinconsistencyentry[i]}
    }
    stpxloopinconsistencytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxloopinconsistencytable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry
// A Spanning Tree instance on a particular port for
// which a Spanning Tree loop-inconsistency is currently
// in effect.
type CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry struct {
    EntityData types.CommonEntityData
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

func (stpxloopinconsistencyentry *CISCOSTPEXTENSIONSMIB_Stpxloopinconsistencytable_Stpxloopinconsistencyentry) GetEntityData() *types.CommonEntityData {
    stpxloopinconsistencyentry.EntityData.YFilter = stpxloopinconsistencyentry.YFilter
    stpxloopinconsistencyentry.EntityData.YangName = "stpxLoopInconsistencyEntry"
    stpxloopinconsistencyentry.EntityData.BundleName = "cisco_ios_xe"
    stpxloopinconsistencyentry.EntityData.ParentYangName = "stpxLoopInconsistencyTable"
    stpxloopinconsistencyentry.EntityData.SegmentPath = "stpxLoopInconsistencyEntry" + "[stpxLoopInconsistencyIndex='" + fmt.Sprintf("%v", stpxloopinconsistencyentry.Stpxloopinconsistencyindex) + "']" + "[stpxLoopInconsistencyPortIndex='" + fmt.Sprintf("%v", stpxloopinconsistencyentry.Stpxloopinconsistencyportindex) + "']"
    stpxloopinconsistencyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxloopinconsistencyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxloopinconsistencyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxloopinconsistencyentry.EntityData.Children = make(map[string]types.YChild)
    stpxloopinconsistencyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxloopinconsistencyentry.EntityData.Leafs["stpxLoopInconsistencyIndex"] = types.YLeaf{"Stpxloopinconsistencyindex", stpxloopinconsistencyentry.Stpxloopinconsistencyindex}
    stpxloopinconsistencyentry.EntityData.Leafs["stpxLoopInconsistencyPortIndex"] = types.YLeaf{"Stpxloopinconsistencyportindex", stpxloopinconsistencyentry.Stpxloopinconsistencyportindex}
    stpxloopinconsistencyentry.EntityData.Leafs["stpxLoopInconsistencyState"] = types.YLeaf{"Stpxloopinconsistencystate", stpxloopinconsistencyentry.Stpxloopinconsistencystate}
    return &(stpxloopinconsistencyentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable
// A table containing a list of the bridge ports for
// which Spanning Tree Port Fast Start can be
// configured.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree Port Fast Start can be configured.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry.
    Stpxfaststartportentry []CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry
}

func (stpxfaststartporttable *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable) GetEntityData() *types.CommonEntityData {
    stpxfaststartporttable.EntityData.YFilter = stpxfaststartporttable.YFilter
    stpxfaststartporttable.EntityData.YangName = "stpxFastStartPortTable"
    stpxfaststartporttable.EntityData.BundleName = "cisco_ios_xe"
    stpxfaststartporttable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxfaststartporttable.EntityData.SegmentPath = "stpxFastStartPortTable"
    stpxfaststartporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxfaststartporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxfaststartporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxfaststartporttable.EntityData.Children = make(map[string]types.YChild)
    stpxfaststartporttable.EntityData.Children["stpxFastStartPortEntry"] = types.YChild{"Stpxfaststartportentry", nil}
    for i := range stpxfaststartporttable.Stpxfaststartportentry {
        stpxfaststartporttable.EntityData.Children[types.GetSegmentPath(&stpxfaststartporttable.Stpxfaststartportentry[i])] = types.YChild{"Stpxfaststartportentry", &stpxfaststartporttable.Stpxfaststartportentry[i]}
    }
    stpxfaststartporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxfaststartporttable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry
// A bridge port for which Spanning Tree Port Fast
// Start can be configured.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry struct {
    EntityData types.CommonEntityData
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

func (stpxfaststartportentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartporttable_Stpxfaststartportentry) GetEntityData() *types.CommonEntityData {
    stpxfaststartportentry.EntityData.YFilter = stpxfaststartportentry.YFilter
    stpxfaststartportentry.EntityData.YangName = "stpxFastStartPortEntry"
    stpxfaststartportentry.EntityData.BundleName = "cisco_ios_xe"
    stpxfaststartportentry.EntityData.ParentYangName = "stpxFastStartPortTable"
    stpxfaststartportentry.EntityData.SegmentPath = "stpxFastStartPortEntry" + "[stpxFastStartPortIndex='" + fmt.Sprintf("%v", stpxfaststartportentry.Stpxfaststartportindex) + "']"
    stpxfaststartportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxfaststartportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxfaststartportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxfaststartportentry.EntityData.Children = make(map[string]types.YChild)
    stpxfaststartportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxfaststartportentry.EntityData.Leafs["stpxFastStartPortIndex"] = types.YLeaf{"Stpxfaststartportindex", stpxfaststartportentry.Stpxfaststartportindex}
    stpxfaststartportentry.EntityData.Leafs["stpxFastStartPortEnable"] = types.YLeaf{"Stpxfaststartportenable", stpxfaststartportentry.Stpxfaststartportenable}
    stpxfaststartportentry.EntityData.Leafs["stpxFastStartPortMode"] = types.YLeaf{"Stpxfaststartportmode", stpxfaststartportentry.Stpxfaststartportmode}
    stpxfaststartportentry.EntityData.Leafs["stpxFastStartPortBpduGuardMode"] = types.YLeaf{"Stpxfaststartportbpduguardmode", stpxfaststartportentry.Stpxfaststartportbpduguardmode}
    stpxfaststartportentry.EntityData.Leafs["stpxFastStartPortBpduFilterMode"] = types.YLeaf{"Stpxfaststartportbpdufiltermode", stpxfaststartportentry.Stpxfaststartportbpdufiltermode}
    return &(stpxfaststartportentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port fast start oper mode  information on a bridge port for a
    // particular  Spanning Tree Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry.
    Stpxfaststartopermodeentry []CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry
}

func (stpxfaststartopermodetable *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable) GetEntityData() *types.CommonEntityData {
    stpxfaststartopermodetable.EntityData.YFilter = stpxfaststartopermodetable.YFilter
    stpxfaststartopermodetable.EntityData.YangName = "stpxFastStartOperModeTable"
    stpxfaststartopermodetable.EntityData.BundleName = "cisco_ios_xe"
    stpxfaststartopermodetable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxfaststartopermodetable.EntityData.SegmentPath = "stpxFastStartOperModeTable"
    stpxfaststartopermodetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxfaststartopermodetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxfaststartopermodetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxfaststartopermodetable.EntityData.Children = make(map[string]types.YChild)
    stpxfaststartopermodetable.EntityData.Children["stpxFastStartOperModeEntry"] = types.YChild{"Stpxfaststartopermodeentry", nil}
    for i := range stpxfaststartopermodetable.Stpxfaststartopermodeentry {
        stpxfaststartopermodetable.EntityData.Children[types.GetSegmentPath(&stpxfaststartopermodetable.Stpxfaststartopermodeentry[i])] = types.YChild{"Stpxfaststartopermodeentry", &stpxfaststartopermodetable.Stpxfaststartopermodeentry[i]}
    }
    stpxfaststartopermodetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxfaststartopermodetable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry
// An entry with port fast start oper mode 
// information on a bridge port for a particular 
// Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry struct {
    EntityData types.CommonEntityData
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

func (stpxfaststartopermodeentry *CISCOSTPEXTENSIONSMIB_Stpxfaststartopermodetable_Stpxfaststartopermodeentry) GetEntityData() *types.CommonEntityData {
    stpxfaststartopermodeentry.EntityData.YFilter = stpxfaststartopermodeentry.YFilter
    stpxfaststartopermodeentry.EntityData.YangName = "stpxFastStartOperModeEntry"
    stpxfaststartopermodeentry.EntityData.BundleName = "cisco_ios_xe"
    stpxfaststartopermodeentry.EntityData.ParentYangName = "stpxFastStartOperModeTable"
    stpxfaststartopermodeentry.EntityData.SegmentPath = "stpxFastStartOperModeEntry" + "[stpxFastStartOperModeInstIndex='" + fmt.Sprintf("%v", stpxfaststartopermodeentry.Stpxfaststartopermodeinstindex) + "']" + "[stpxFastStartOperModePortIndex='" + fmt.Sprintf("%v", stpxfaststartopermodeentry.Stpxfaststartopermodeportindex) + "']"
    stpxfaststartopermodeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxfaststartopermodeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxfaststartopermodeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxfaststartopermodeentry.EntityData.Children = make(map[string]types.YChild)
    stpxfaststartopermodeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxfaststartopermodeentry.EntityData.Leafs["stpxFastStartOperModeInstIndex"] = types.YLeaf{"Stpxfaststartopermodeinstindex", stpxfaststartopermodeentry.Stpxfaststartopermodeinstindex}
    stpxfaststartopermodeentry.EntityData.Leafs["stpxFastStartOperModePortIndex"] = types.YLeaf{"Stpxfaststartopermodeportindex", stpxfaststartopermodeentry.Stpxfaststartopermodeportindex}
    stpxfaststartopermodeentry.EntityData.Leafs["stpxFastStartOperMode"] = types.YLeaf{"Stpxfaststartopermode", stpxfaststartopermodeentry.Stpxfaststartopermode}
    return &(stpxfaststartopermodeentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which BPDU skewing has
    // been detected. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry.
    Stpxbpduskewingentry []CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry
}

func (stpxbpduskewingtable *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable) GetEntityData() *types.CommonEntityData {
    stpxbpduskewingtable.EntityData.YFilter = stpxbpduskewingtable.YFilter
    stpxbpduskewingtable.EntityData.YangName = "stpxBpduSkewingTable"
    stpxbpduskewingtable.EntityData.BundleName = "cisco_ios_xe"
    stpxbpduskewingtable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxbpduskewingtable.EntityData.SegmentPath = "stpxBpduSkewingTable"
    stpxbpduskewingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxbpduskewingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxbpduskewingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxbpduskewingtable.EntityData.Children = make(map[string]types.YChild)
    stpxbpduskewingtable.EntityData.Children["stpxBpduSkewingEntry"] = types.YChild{"Stpxbpduskewingentry", nil}
    for i := range stpxbpduskewingtable.Stpxbpduskewingentry {
        stpxbpduskewingtable.EntityData.Children[types.GetSegmentPath(&stpxbpduskewingtable.Stpxbpduskewingentry[i])] = types.YChild{"Stpxbpduskewingentry", &stpxbpduskewingtable.Stpxbpduskewingentry[i]}
    }
    stpxbpduskewingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxbpduskewingtable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry
// A Spanning Tree instance on a particular port for
// which BPDU skewing has been detected.
type CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry struct {
    EntityData types.CommonEntityData
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

func (stpxbpduskewingentry *CISCOSTPEXTENSIONSMIB_Stpxbpduskewingtable_Stpxbpduskewingentry) GetEntityData() *types.CommonEntityData {
    stpxbpduskewingentry.EntityData.YFilter = stpxbpduskewingentry.YFilter
    stpxbpduskewingentry.EntityData.YangName = "stpxBpduSkewingEntry"
    stpxbpduskewingentry.EntityData.BundleName = "cisco_ios_xe"
    stpxbpduskewingentry.EntityData.ParentYangName = "stpxBpduSkewingTable"
    stpxbpduskewingentry.EntityData.SegmentPath = "stpxBpduSkewingEntry" + "[stpxBpduSkewingInstanceIndex='" + fmt.Sprintf("%v", stpxbpduskewingentry.Stpxbpduskewinginstanceindex) + "']" + "[stpxBpduSkewingPortIndex='" + fmt.Sprintf("%v", stpxbpduskewingentry.Stpxbpduskewingportindex) + "']"
    stpxbpduskewingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxbpduskewingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxbpduskewingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxbpduskewingentry.EntityData.Children = make(map[string]types.YChild)
    stpxbpduskewingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxbpduskewingentry.EntityData.Leafs["stpxBpduSkewingInstanceIndex"] = types.YLeaf{"Stpxbpduskewinginstanceindex", stpxbpduskewingentry.Stpxbpduskewinginstanceindex}
    stpxbpduskewingentry.EntityData.Leafs["stpxBpduSkewingPortIndex"] = types.YLeaf{"Stpxbpduskewingportindex", stpxbpduskewingentry.Stpxbpduskewingportindex}
    stpxbpduskewingentry.EntityData.Leafs["stpxBpduSkewingLastSkewDuration"] = types.YLeaf{"Stpxbpduskewinglastskewduration", stpxbpduskewingentry.Stpxbpduskewinglastskewduration}
    stpxbpduskewingentry.EntityData.Leafs["stpxBpduSkewingWorstSkewDuration"] = types.YLeaf{"Stpxbpduskewingworstskewduration", stpxbpduskewingentry.Stpxbpduskewingworstskewduration}
    stpxbpduskewingentry.EntityData.Leafs["stpxBpduSkewingWorstSkewTime"] = types.YLeaf{"Stpxbpduskewingworstskewtime", stpxbpduskewingentry.Stpxbpduskewingworstskewtime}
    return &(stpxbpduskewingentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable
// This table contains MST instance information with
// one entry for an MST instance within the range of 
// 0 to the object value of stpxMSTMaxInstanceNumber. 
// 
// This table is deprecated and replaced by 
// stpxSMSTInstanceTable.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information. The type is
    // slice of CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry.
    Stpxmstinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry
}

func (stpxmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable) GetEntityData() *types.CommonEntityData {
    stpxmstinstancetable.EntityData.YFilter = stpxmstinstancetable.YFilter
    stpxmstinstancetable.EntityData.YangName = "stpxMSTInstanceTable"
    stpxmstinstancetable.EntityData.BundleName = "cisco_ios_xe"
    stpxmstinstancetable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmstinstancetable.EntityData.SegmentPath = "stpxMSTInstanceTable"
    stpxmstinstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstinstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstinstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstinstancetable.EntityData.Children = make(map[string]types.YChild)
    stpxmstinstancetable.EntityData.Children["stpxMSTInstanceEntry"] = types.YChild{"Stpxmstinstanceentry", nil}
    for i := range stpxmstinstancetable.Stpxmstinstanceentry {
        stpxmstinstancetable.EntityData.Children[types.GetSegmentPath(&stpxmstinstancetable.Stpxmstinstanceentry[i])] = types.YChild{"Stpxmstinstanceentry", &stpxmstinstancetable.Stpxmstinstanceentry[i]}
    }
    stpxmstinstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxmstinstancetable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry
// A conceptual row containing the MST instance 
// information.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry struct {
    EntityData types.CommonEntityData
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

func (stpxmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstancetable_Stpxmstinstanceentry) GetEntityData() *types.CommonEntityData {
    stpxmstinstanceentry.EntityData.YFilter = stpxmstinstanceentry.YFilter
    stpxmstinstanceentry.EntityData.YangName = "stpxMSTInstanceEntry"
    stpxmstinstanceentry.EntityData.BundleName = "cisco_ios_xe"
    stpxmstinstanceentry.EntityData.ParentYangName = "stpxMSTInstanceTable"
    stpxmstinstanceentry.EntityData.SegmentPath = "stpxMSTInstanceEntry" + "[stpxMSTInstanceIndex='" + fmt.Sprintf("%v", stpxmstinstanceentry.Stpxmstinstanceindex) + "']"
    stpxmstinstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstinstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstinstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstinstanceentry.EntityData.Children = make(map[string]types.YChild)
    stpxmstinstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceIndex"] = types.YLeaf{"Stpxmstinstanceindex", stpxmstinstanceentry.Stpxmstinstanceindex}
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceVlansMapped"] = types.YLeaf{"Stpxmstinstancevlansmapped", stpxmstinstanceentry.Stpxmstinstancevlansmapped}
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceVlansMapped2k"] = types.YLeaf{"Stpxmstinstancevlansmapped2K", stpxmstinstanceentry.Stpxmstinstancevlansmapped2K}
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceVlansMapped3k"] = types.YLeaf{"Stpxmstinstancevlansmapped3K", stpxmstinstanceentry.Stpxmstinstancevlansmapped3K}
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceVlansMapped4k"] = types.YLeaf{"Stpxmstinstancevlansmapped4K", stpxmstinstanceentry.Stpxmstinstancevlansmapped4K}
    stpxmstinstanceentry.EntityData.Leafs["stpxMSTInstanceRemainingHopCount"] = types.YLeaf{"Stpxmstinstanceremaininghopcount", stpxmstinstanceentry.Stpxmstinstanceremaininghopcount}
    return &(stpxmstinstanceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry.
    Stpxmstinstanceeditentry []CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry
}

func (stpxmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable) GetEntityData() *types.CommonEntityData {
    stpxmstinstanceedittable.EntityData.YFilter = stpxmstinstanceedittable.YFilter
    stpxmstinstanceedittable.EntityData.YangName = "stpxMSTInstanceEditTable"
    stpxmstinstanceedittable.EntityData.BundleName = "cisco_ios_xe"
    stpxmstinstanceedittable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmstinstanceedittable.EntityData.SegmentPath = "stpxMSTInstanceEditTable"
    stpxmstinstanceedittable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstinstanceedittable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstinstanceedittable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstinstanceedittable.EntityData.Children = make(map[string]types.YChild)
    stpxmstinstanceedittable.EntityData.Children["stpxMSTInstanceEditEntry"] = types.YChild{"Stpxmstinstanceeditentry", nil}
    for i := range stpxmstinstanceedittable.Stpxmstinstanceeditentry {
        stpxmstinstanceedittable.EntityData.Children[types.GetSegmentPath(&stpxmstinstanceedittable.Stpxmstinstanceeditentry[i])] = types.YChild{"Stpxmstinstanceeditentry", &stpxmstinstanceedittable.Stpxmstinstanceeditentry[i]}
    }
    stpxmstinstanceedittable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxmstinstanceedittable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
type CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry struct {
    EntityData types.CommonEntityData
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

func (stpxmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxmstinstanceedittable_Stpxmstinstanceeditentry) GetEntityData() *types.CommonEntityData {
    stpxmstinstanceeditentry.EntityData.YFilter = stpxmstinstanceeditentry.YFilter
    stpxmstinstanceeditentry.EntityData.YangName = "stpxMSTInstanceEditEntry"
    stpxmstinstanceeditentry.EntityData.BundleName = "cisco_ios_xe"
    stpxmstinstanceeditentry.EntityData.ParentYangName = "stpxMSTInstanceEditTable"
    stpxmstinstanceeditentry.EntityData.SegmentPath = "stpxMSTInstanceEditEntry" + "[stpxMSTInstanceEditIndex='" + fmt.Sprintf("%v", stpxmstinstanceeditentry.Stpxmstinstanceeditindex) + "']"
    stpxmstinstanceeditentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstinstanceeditentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstinstanceeditentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstinstanceeditentry.EntityData.Children = make(map[string]types.YChild)
    stpxmstinstanceeditentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmstinstanceeditentry.EntityData.Leafs["stpxMSTInstanceEditIndex"] = types.YLeaf{"Stpxmstinstanceeditindex", stpxmstinstanceeditentry.Stpxmstinstanceeditindex}
    stpxmstinstanceeditentry.EntityData.Leafs["stpxMSTInstanceEditVlansMap"] = types.YLeaf{"Stpxmstinstanceeditvlansmap", stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap}
    stpxmstinstanceeditentry.EntityData.Leafs["stpxMSTInstanceEditVlansMap2k"] = types.YLeaf{"Stpxmstinstanceeditvlansmap2K", stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap2K}
    stpxmstinstanceeditentry.EntityData.Leafs["stpxMSTInstanceEditVlansMap3k"] = types.YLeaf{"Stpxmstinstanceeditvlansmap3K", stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap3K}
    stpxmstinstanceeditentry.EntityData.Leafs["stpxMSTInstanceEditVlansMap4k"] = types.YLeaf{"Stpxmstinstanceeditvlansmap4K", stpxmstinstanceeditentry.Stpxmstinstanceeditvlansmap4K}
    return &(stpxmstinstanceeditentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the MST Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry.
    Stpxmstportentry []CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry
}

func (stpxmstporttable *CISCOSTPEXTENSIONSMIB_Stpxmstporttable) GetEntityData() *types.CommonEntityData {
    stpxmstporttable.EntityData.YFilter = stpxmstporttable.YFilter
    stpxmstporttable.EntityData.YangName = "stpxMSTPortTable"
    stpxmstporttable.EntityData.BundleName = "cisco_ios_xe"
    stpxmstporttable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmstporttable.EntityData.SegmentPath = "stpxMSTPortTable"
    stpxmstporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstporttable.EntityData.Children = make(map[string]types.YChild)
    stpxmstporttable.EntityData.Children["stpxMSTPortEntry"] = types.YChild{"Stpxmstportentry", nil}
    for i := range stpxmstporttable.Stpxmstportentry {
        stpxmstporttable.EntityData.Children[types.GetSegmentPath(&stpxmstporttable.Stpxmstportentry[i])] = types.YChild{"Stpxmstportentry", &stpxmstporttable.Stpxmstportentry[i]}
    }
    stpxmstporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxmstporttable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry
// An entry with port information for the MST Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry struct {
    EntityData types.CommonEntityData
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

func (stpxmstportentry *CISCOSTPEXTENSIONSMIB_Stpxmstporttable_Stpxmstportentry) GetEntityData() *types.CommonEntityData {
    stpxmstportentry.EntityData.YFilter = stpxmstportentry.YFilter
    stpxmstportentry.EntityData.YangName = "stpxMSTPortEntry"
    stpxmstportentry.EntityData.BundleName = "cisco_ios_xe"
    stpxmstportentry.EntityData.ParentYangName = "stpxMSTPortTable"
    stpxmstportentry.EntityData.SegmentPath = "stpxMSTPortEntry" + "[stpxMSTPortIndex='" + fmt.Sprintf("%v", stpxmstportentry.Stpxmstportindex) + "']"
    stpxmstportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstportentry.EntityData.Children = make(map[string]types.YChild)
    stpxmstportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmstportentry.EntityData.Leafs["stpxMSTPortIndex"] = types.YLeaf{"Stpxmstportindex", stpxmstportentry.Stpxmstportindex}
    stpxmstportentry.EntityData.Leafs["stpxMSTPortAdminLinkType"] = types.YLeaf{"Stpxmstportadminlinktype", stpxmstportentry.Stpxmstportadminlinktype}
    stpxmstportentry.EntityData.Leafs["stpxMSTPortOperLinkType"] = types.YLeaf{"Stpxmstportoperlinktype", stpxmstportentry.Stpxmstportoperlinktype}
    stpxmstportentry.EntityData.Leafs["stpxMSTPortProtocolMigration"] = types.YLeaf{"Stpxmstportprotocolmigration", stpxmstportentry.Stpxmstportprotocolmigration}
    stpxmstportentry.EntityData.Leafs["stpxMSTPortStatus"] = types.YLeaf{"Stpxmstportstatus", stpxmstportentry.Stpxmstportstatus}
    return &(stpxmstportentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the port role information for the MST protocol on a
    // port for a particular MST instance existing on the system. The type is
    // slice of CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry.
    Stpxmstportroleentry []CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry
}

func (stpxmstportroletable *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable) GetEntityData() *types.CommonEntityData {
    stpxmstportroletable.EntityData.YFilter = stpxmstportroletable.YFilter
    stpxmstportroletable.EntityData.YangName = "stpxMSTPortRoleTable"
    stpxmstportroletable.EntityData.BundleName = "cisco_ios_xe"
    stpxmstportroletable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxmstportroletable.EntityData.SegmentPath = "stpxMSTPortRoleTable"
    stpxmstportroletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstportroletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstportroletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstportroletable.EntityData.Children = make(map[string]types.YChild)
    stpxmstportroletable.EntityData.Children["stpxMSTPortRoleEntry"] = types.YChild{"Stpxmstportroleentry", nil}
    for i := range stpxmstportroletable.Stpxmstportroleentry {
        stpxmstportroletable.EntityData.Children[types.GetSegmentPath(&stpxmstportroletable.Stpxmstportroleentry[i])] = types.YChild{"Stpxmstportroleentry", &stpxmstportroletable.Stpxmstportroleentry[i]}
    }
    stpxmstportroletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxmstportroletable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry
// An entry containing the port role information for the MST
// protocol on a port for a particular MST instance existing
// on the system.
type CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry struct {
    EntityData types.CommonEntityData
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

func (stpxmstportroleentry *CISCOSTPEXTENSIONSMIB_Stpxmstportroletable_Stpxmstportroleentry) GetEntityData() *types.CommonEntityData {
    stpxmstportroleentry.EntityData.YFilter = stpxmstportroleentry.YFilter
    stpxmstportroleentry.EntityData.YangName = "stpxMSTPortRoleEntry"
    stpxmstportroleentry.EntityData.BundleName = "cisco_ios_xe"
    stpxmstportroleentry.EntityData.ParentYangName = "stpxMSTPortRoleTable"
    stpxmstportroleentry.EntityData.SegmentPath = "stpxMSTPortRoleEntry" + "[stpxMSTPortRoleInstanceIndex='" + fmt.Sprintf("%v", stpxmstportroleentry.Stpxmstportroleinstanceindex) + "']" + "[stpxMSTPortRolePortIndex='" + fmt.Sprintf("%v", stpxmstportroleentry.Stpxmstportroleportindex) + "']"
    stpxmstportroleentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxmstportroleentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxmstportroleentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxmstportroleentry.EntityData.Children = make(map[string]types.YChild)
    stpxmstportroleentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxmstportroleentry.EntityData.Leafs["stpxMSTPortRoleInstanceIndex"] = types.YLeaf{"Stpxmstportroleinstanceindex", stpxmstportroleentry.Stpxmstportroleinstanceindex}
    stpxmstportroleentry.EntityData.Leafs["stpxMSTPortRolePortIndex"] = types.YLeaf{"Stpxmstportroleportindex", stpxmstportroleentry.Stpxmstportroleportindex}
    stpxmstportroleentry.EntityData.Leafs["stpxMSTPortRoleValue"] = types.YLeaf{"Stpxmstportrolevalue", stpxmstportroleentry.Stpxmstportrolevalue}
    return &(stpxmstportroleentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the RSTP Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry.
    Stpxrstpportentry []CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry
}

func (stpxrstpporttable *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable) GetEntityData() *types.CommonEntityData {
    stpxrstpporttable.EntityData.YFilter = stpxrstpporttable.YFilter
    stpxrstpporttable.EntityData.YangName = "stpxRSTPPortTable"
    stpxrstpporttable.EntityData.BundleName = "cisco_ios_xe"
    stpxrstpporttable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrstpporttable.EntityData.SegmentPath = "stpxRSTPPortTable"
    stpxrstpporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrstpporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrstpporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrstpporttable.EntityData.Children = make(map[string]types.YChild)
    stpxrstpporttable.EntityData.Children["stpxRSTPPortEntry"] = types.YChild{"Stpxrstpportentry", nil}
    for i := range stpxrstpporttable.Stpxrstpportentry {
        stpxrstpporttable.EntityData.Children[types.GetSegmentPath(&stpxrstpporttable.Stpxrstpportentry[i])] = types.YChild{"Stpxrstpportentry", &stpxrstpporttable.Stpxrstpportentry[i]}
    }
    stpxrstpporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxrstpporttable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry
// An entry with port information for the RSTP Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry struct {
    EntityData types.CommonEntityData
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

func (stpxrstpportentry *CISCOSTPEXTENSIONSMIB_Stpxrstpporttable_Stpxrstpportentry) GetEntityData() *types.CommonEntityData {
    stpxrstpportentry.EntityData.YFilter = stpxrstpportentry.YFilter
    stpxrstpportentry.EntityData.YangName = "stpxRSTPPortEntry"
    stpxrstpportentry.EntityData.BundleName = "cisco_ios_xe"
    stpxrstpportentry.EntityData.ParentYangName = "stpxRSTPPortTable"
    stpxrstpportentry.EntityData.SegmentPath = "stpxRSTPPortEntry" + "[stpxRSTPPortIndex='" + fmt.Sprintf("%v", stpxrstpportentry.Stpxrstpportindex) + "']"
    stpxrstpportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrstpportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrstpportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrstpportentry.EntityData.Children = make(map[string]types.YChild)
    stpxrstpportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrstpportentry.EntityData.Leafs["stpxRSTPPortIndex"] = types.YLeaf{"Stpxrstpportindex", stpxrstpportentry.Stpxrstpportindex}
    stpxrstpportentry.EntityData.Leafs["stpxRSTPPortAdminLinkType"] = types.YLeaf{"Stpxrstpportadminlinktype", stpxrstpportentry.Stpxrstpportadminlinktype}
    stpxrstpportentry.EntityData.Leafs["stpxRSTPPortOperLinkType"] = types.YLeaf{"Stpxrstpportoperlinktype", stpxrstpportentry.Stpxrstpportoperlinktype}
    stpxrstpportentry.EntityData.Leafs["stpxRSTPPortProtocolMigration"] = types.YLeaf{"Stpxrstpportprotocolmigration", stpxrstpportentry.Stpxrstpportprotocolmigration}
    return &(stpxrstpportentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the port role information for the RSTP protocol on a
    // port for a particular Spanning Tree instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry.
    Stpxrstpportroleentry []CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry
}

func (stpxrstpportroletable *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable) GetEntityData() *types.CommonEntityData {
    stpxrstpportroletable.EntityData.YFilter = stpxrstpportroletable.YFilter
    stpxrstpportroletable.EntityData.YangName = "stpxRSTPPortRoleTable"
    stpxrstpportroletable.EntityData.BundleName = "cisco_ios_xe"
    stpxrstpportroletable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrstpportroletable.EntityData.SegmentPath = "stpxRSTPPortRoleTable"
    stpxrstpportroletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrstpportroletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrstpportroletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrstpportroletable.EntityData.Children = make(map[string]types.YChild)
    stpxrstpportroletable.EntityData.Children["stpxRSTPPortRoleEntry"] = types.YChild{"Stpxrstpportroleentry", nil}
    for i := range stpxrstpportroletable.Stpxrstpportroleentry {
        stpxrstpportroletable.EntityData.Children[types.GetSegmentPath(&stpxrstpportroletable.Stpxrstpportroleentry[i])] = types.YChild{"Stpxrstpportroleentry", &stpxrstpportroletable.Stpxrstpportroleentry[i]}
    }
    stpxrstpportroletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxrstpportroletable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry
// An entry containing the port role information for the RSTP
// protocol on a port for a particular Spanning Tree instance.
type CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry struct {
    EntityData types.CommonEntityData
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

func (stpxrstpportroleentry *CISCOSTPEXTENSIONSMIB_Stpxrstpportroletable_Stpxrstpportroleentry) GetEntityData() *types.CommonEntityData {
    stpxrstpportroleentry.EntityData.YFilter = stpxrstpportroleentry.YFilter
    stpxrstpportroleentry.EntityData.YangName = "stpxRSTPPortRoleEntry"
    stpxrstpportroleentry.EntityData.BundleName = "cisco_ios_xe"
    stpxrstpportroleentry.EntityData.ParentYangName = "stpxRSTPPortRoleTable"
    stpxrstpportroleentry.EntityData.SegmentPath = "stpxRSTPPortRoleEntry" + "[stpxRSTPPortRoleInstanceIndex='" + fmt.Sprintf("%v", stpxrstpportroleentry.Stpxrstpportroleinstanceindex) + "']" + "[stpxRSTPPortRolePortIndex='" + fmt.Sprintf("%v", stpxrstpportroleentry.Stpxrstpportroleportindex) + "']"
    stpxrstpportroleentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrstpportroleentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrstpportroleentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrstpportroleentry.EntityData.Children = make(map[string]types.YChild)
    stpxrstpportroleentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrstpportroleentry.EntityData.Leafs["stpxRSTPPortRoleInstanceIndex"] = types.YLeaf{"Stpxrstpportroleinstanceindex", stpxrstpportroleentry.Stpxrstpportroleinstanceindex}
    stpxrstpportroleentry.EntityData.Leafs["stpxRSTPPortRolePortIndex"] = types.YLeaf{"Stpxrstpportroleportindex", stpxrstpportroleentry.Stpxrstpportroleportindex}
    stpxrstpportroleentry.EntityData.Leafs["stpxRSTPPortRoleValue"] = types.YLeaf{"Stpxrstpportrolevalue", stpxrstpportroleentry.Stpxrstpportrolevalue}
    return &(stpxrstpportroleentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port status information on a  bridge port for a particular
    // Spanning Tree  Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry.
    Stpxrpvstportentry []CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry
}

func (stpxrpvstporttable *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable) GetEntityData() *types.CommonEntityData {
    stpxrpvstporttable.EntityData.YFilter = stpxrpvstporttable.YFilter
    stpxrpvstporttable.EntityData.YangName = "stpxRPVSTPortTable"
    stpxrpvstporttable.EntityData.BundleName = "cisco_ios_xe"
    stpxrpvstporttable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxrpvstporttable.EntityData.SegmentPath = "stpxRPVSTPortTable"
    stpxrpvstporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrpvstporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrpvstporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrpvstporttable.EntityData.Children = make(map[string]types.YChild)
    stpxrpvstporttable.EntityData.Children["stpxRPVSTPortEntry"] = types.YChild{"Stpxrpvstportentry", nil}
    for i := range stpxrpvstporttable.Stpxrpvstportentry {
        stpxrpvstporttable.EntityData.Children[types.GetSegmentPath(&stpxrpvstporttable.Stpxrpvstportentry[i])] = types.YChild{"Stpxrpvstportentry", &stpxrpvstporttable.Stpxrpvstportentry[i]}
    }
    stpxrpvstporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxrpvstporttable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry
// An entry with port status information on a 
// bridge port for a particular Spanning Tree 
// Instance.
type CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry struct {
    EntityData types.CommonEntityData
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

func (stpxrpvstportentry *CISCOSTPEXTENSIONSMIB_Stpxrpvstporttable_Stpxrpvstportentry) GetEntityData() *types.CommonEntityData {
    stpxrpvstportentry.EntityData.YFilter = stpxrpvstportentry.YFilter
    stpxrpvstportentry.EntityData.YangName = "stpxRPVSTPortEntry"
    stpxrpvstportentry.EntityData.BundleName = "cisco_ios_xe"
    stpxrpvstportentry.EntityData.ParentYangName = "stpxRPVSTPortTable"
    stpxrpvstportentry.EntityData.SegmentPath = "stpxRPVSTPortEntry" + "[stpxRPVSTPortVlanIndex='" + fmt.Sprintf("%v", stpxrpvstportentry.Stpxrpvstportvlanindex) + "']" + "[stpxRPVSTPortIndex='" + fmt.Sprintf("%v", stpxrpvstportentry.Stpxrpvstportindex) + "']"
    stpxrpvstportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxrpvstportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxrpvstportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxrpvstportentry.EntityData.Children = make(map[string]types.YChild)
    stpxrpvstportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxrpvstportentry.EntityData.Leafs["stpxRPVSTPortVlanIndex"] = types.YLeaf{"Stpxrpvstportvlanindex", stpxrpvstportentry.Stpxrpvstportvlanindex}
    stpxrpvstportentry.EntityData.Leafs["stpxRPVSTPortIndex"] = types.YLeaf{"Stpxrpvstportindex", stpxrpvstportentry.Stpxrpvstportindex}
    stpxrpvstportentry.EntityData.Leafs["stpxRPVSTPortStatus"] = types.YLeaf{"Stpxrpvstportstatus", stpxrpvstportentry.Stpxrpvstportstatus}
    return &(stpxrpvstportentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable
// This table contains MST instance information
// for IEEE MST.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information for IEEE MST. The
    // type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry.
    Stpxsmstinstanceentry []CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry
}

func (stpxsmstinstancetable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable) GetEntityData() *types.CommonEntityData {
    stpxsmstinstancetable.EntityData.YFilter = stpxsmstinstancetable.YFilter
    stpxsmstinstancetable.EntityData.YangName = "stpxSMSTInstanceTable"
    stpxsmstinstancetable.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstinstancetable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxsmstinstancetable.EntityData.SegmentPath = "stpxSMSTInstanceTable"
    stpxsmstinstancetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstinstancetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstinstancetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstinstancetable.EntityData.Children = make(map[string]types.YChild)
    stpxsmstinstancetable.EntityData.Children["stpxSMSTInstanceEntry"] = types.YChild{"Stpxsmstinstanceentry", nil}
    for i := range stpxsmstinstancetable.Stpxsmstinstanceentry {
        stpxsmstinstancetable.EntityData.Children[types.GetSegmentPath(&stpxsmstinstancetable.Stpxsmstinstanceentry[i])] = types.YChild{"Stpxsmstinstanceentry", &stpxsmstinstancetable.Stpxsmstinstanceentry[i]}
    }
    stpxsmstinstancetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxsmstinstancetable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry
// A conceptual row containing the MST instance 
// information for IEEE MST.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry struct {
    EntityData types.CommonEntityData
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

func (stpxsmstinstanceentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstancetable_Stpxsmstinstanceentry) GetEntityData() *types.CommonEntityData {
    stpxsmstinstanceentry.EntityData.YFilter = stpxsmstinstanceentry.YFilter
    stpxsmstinstanceentry.EntityData.YangName = "stpxSMSTInstanceEntry"
    stpxsmstinstanceentry.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstinstanceentry.EntityData.ParentYangName = "stpxSMSTInstanceTable"
    stpxsmstinstanceentry.EntityData.SegmentPath = "stpxSMSTInstanceEntry" + "[stpxSMSTInstanceIndex='" + fmt.Sprintf("%v", stpxsmstinstanceentry.Stpxsmstinstanceindex) + "']"
    stpxsmstinstanceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstinstanceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstinstanceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstinstanceentry.EntityData.Children = make(map[string]types.YChild)
    stpxsmstinstanceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceIndex"] = types.YLeaf{"Stpxsmstinstanceindex", stpxsmstinstanceentry.Stpxsmstinstanceindex}
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceVlansMapped1k2k"] = types.YLeaf{"Stpxsmstinstancevlansmapped1K2K", stpxsmstinstanceentry.Stpxsmstinstancevlansmapped1K2K}
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceVlansMapped3k4k"] = types.YLeaf{"Stpxsmstinstancevlansmapped3K4K", stpxsmstinstanceentry.Stpxsmstinstancevlansmapped3K4K}
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceRemainingHopCount"] = types.YLeaf{"Stpxsmstinstanceremaininghopcount", stpxsmstinstanceentry.Stpxsmstinstanceremaininghopcount}
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceCISTRegionalRoot"] = types.YLeaf{"Stpxsmstinstancecistregionalroot", stpxsmstinstanceentry.Stpxsmstinstancecistregionalroot}
    stpxsmstinstanceentry.EntityData.Leafs["stpxSMSTInstanceCISTIntRootCost"] = types.YLeaf{"Stpxsmstinstancecistintrootcost", stpxsmstinstanceentry.Stpxsmstinstancecistintrootcost}
    return &(stpxsmstinstanceentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable
// This table contains MST instance information in the 
// Edit Buffer. 
// 
// This table is only instantiated when the object value
// of  stpxMSTRegionEditBufferStatus has the value of
// acquiredBySnmp(2).
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer. 
    // The total number of entries in this table has to be  less than or equal to
    // the object value of stpxSMSTMaxInstances. The type is slice of
    // CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry.
    Stpxsmstinstanceeditentry []CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry
}

func (stpxsmstinstanceedittable *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable) GetEntityData() *types.CommonEntityData {
    stpxsmstinstanceedittable.EntityData.YFilter = stpxsmstinstanceedittable.YFilter
    stpxsmstinstanceedittable.EntityData.YangName = "stpxSMSTInstanceEditTable"
    stpxsmstinstanceedittable.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstinstanceedittable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxsmstinstanceedittable.EntityData.SegmentPath = "stpxSMSTInstanceEditTable"
    stpxsmstinstanceedittable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstinstanceedittable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstinstanceedittable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstinstanceedittable.EntityData.Children = make(map[string]types.YChild)
    stpxsmstinstanceedittable.EntityData.Children["stpxSMSTInstanceEditEntry"] = types.YChild{"Stpxsmstinstanceeditentry", nil}
    for i := range stpxsmstinstanceedittable.Stpxsmstinstanceeditentry {
        stpxsmstinstanceedittable.EntityData.Children[types.GetSegmentPath(&stpxsmstinstanceedittable.Stpxsmstinstanceeditentry[i])] = types.YChild{"Stpxsmstinstanceeditentry", &stpxsmstinstanceedittable.Stpxsmstinstanceeditentry[i]}
    }
    stpxsmstinstanceedittable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxsmstinstanceedittable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
// 
// The total number of entries in this table has to be 
// less than or equal to the object value of stpxSMSTMaxInstances.
type CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry struct {
    EntityData types.CommonEntityData
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

func (stpxsmstinstanceeditentry *CISCOSTPEXTENSIONSMIB_Stpxsmstinstanceedittable_Stpxsmstinstanceeditentry) GetEntityData() *types.CommonEntityData {
    stpxsmstinstanceeditentry.EntityData.YFilter = stpxsmstinstanceeditentry.YFilter
    stpxsmstinstanceeditentry.EntityData.YangName = "stpxSMSTInstanceEditEntry"
    stpxsmstinstanceeditentry.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstinstanceeditentry.EntityData.ParentYangName = "stpxSMSTInstanceEditTable"
    stpxsmstinstanceeditentry.EntityData.SegmentPath = "stpxSMSTInstanceEditEntry" + "[stpxSMSTInstanceEditIndex='" + fmt.Sprintf("%v", stpxsmstinstanceeditentry.Stpxsmstinstanceeditindex) + "']"
    stpxsmstinstanceeditentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstinstanceeditentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstinstanceeditentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstinstanceeditentry.EntityData.Children = make(map[string]types.YChild)
    stpxsmstinstanceeditentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxsmstinstanceeditentry.EntityData.Leafs["stpxSMSTInstanceEditIndex"] = types.YLeaf{"Stpxsmstinstanceeditindex", stpxsmstinstanceeditentry.Stpxsmstinstanceeditindex}
    stpxsmstinstanceeditentry.EntityData.Leafs["stpxSMSTInstanceEditVlansMap1k2k"] = types.YLeaf{"Stpxsmstinstanceeditvlansmap1K2K", stpxsmstinstanceeditentry.Stpxsmstinstanceeditvlansmap1K2K}
    stpxsmstinstanceeditentry.EntityData.Leafs["stpxSMSTInstanceEditVlansMap3k4k"] = types.YLeaf{"Stpxsmstinstanceeditvlansmap3K4K", stpxsmstinstanceeditentry.Stpxsmstinstanceeditvlansmap3K4K}
    stpxsmstinstanceeditentry.EntityData.Leafs["stpxSMSTInstanceEditRowStatus"] = types.YLeaf{"Stpxsmstinstanceeditrowstatus", stpxsmstinstanceeditentry.Stpxsmstinstanceeditrowstatus}
    return &(stpxsmstinstanceeditentry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
// 
// This table is only instantiated when the object 
// value of stpxSpanningTreeType is mst(4)
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the MST protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry.
    Stpxsmstportentry []CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry
}

func (stpxsmstporttable *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable) GetEntityData() *types.CommonEntityData {
    stpxsmstporttable.EntityData.YFilter = stpxsmstporttable.YFilter
    stpxsmstporttable.EntityData.YangName = "stpxSMSTPortTable"
    stpxsmstporttable.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstporttable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxsmstporttable.EntityData.SegmentPath = "stpxSMSTPortTable"
    stpxsmstporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstporttable.EntityData.Children = make(map[string]types.YChild)
    stpxsmstporttable.EntityData.Children["stpxSMSTPortEntry"] = types.YChild{"Stpxsmstportentry", nil}
    for i := range stpxsmstporttable.Stpxsmstportentry {
        stpxsmstporttable.EntityData.Children[types.GetSegmentPath(&stpxsmstporttable.Stpxsmstportentry[i])] = types.YChild{"Stpxsmstportentry", &stpxsmstporttable.Stpxsmstportentry[i]}
    }
    stpxsmstporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stpxsmstporttable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry
// An entry with port information for the MST protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry struct {
    EntityData types.CommonEntityData
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

func (stpxsmstportentry *CISCOSTPEXTENSIONSMIB_Stpxsmstporttable_Stpxsmstportentry) GetEntityData() *types.CommonEntityData {
    stpxsmstportentry.EntityData.YFilter = stpxsmstportentry.YFilter
    stpxsmstportentry.EntityData.YangName = "stpxSMSTPortEntry"
    stpxsmstportentry.EntityData.BundleName = "cisco_ios_xe"
    stpxsmstportentry.EntityData.ParentYangName = "stpxSMSTPortTable"
    stpxsmstportentry.EntityData.SegmentPath = "stpxSMSTPortEntry" + "[stpxSMSTPortIndex='" + fmt.Sprintf("%v", stpxsmstportentry.Stpxsmstportindex) + "']"
    stpxsmstportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxsmstportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxsmstportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxsmstportentry.EntityData.Children = make(map[string]types.YChild)
    stpxsmstportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortIndex"] = types.YLeaf{"Stpxsmstportindex", stpxsmstportentry.Stpxsmstportindex}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortStatus"] = types.YLeaf{"Stpxsmstportstatus", stpxsmstportentry.Stpxsmstportstatus}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortAdminHelloTime"] = types.YLeaf{"Stpxsmstportadminhellotime", stpxsmstportentry.Stpxsmstportadminhellotime}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortConfigedHelloTime"] = types.YLeaf{"Stpxsmstportconfigedhellotime", stpxsmstportentry.Stpxsmstportconfigedhellotime}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortOperHelloTime"] = types.YLeaf{"Stpxsmstportoperhellotime", stpxsmstportentry.Stpxsmstportoperhellotime}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortAdminMSTMode"] = types.YLeaf{"Stpxsmstportadminmstmode", stpxsmstportentry.Stpxsmstportadminmstmode}
    stpxsmstportentry.EntityData.Leafs["stpxSMSTPortOperMSTMode"] = types.YLeaf{"Stpxsmstportopermstmode", stpxsmstportentry.Stpxsmstportopermstmode}
    return &(stpxsmstportentry.EntityData)
}

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

