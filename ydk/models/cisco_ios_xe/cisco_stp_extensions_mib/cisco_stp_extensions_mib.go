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

    
    StpxUplinkFastObjects CISCOSTPEXTENSIONSMIB_StpxUplinkFastObjects

    
    StpxBackboneFastObjects CISCOSTPEXTENSIONSMIB_StpxBackboneFastObjects

    
    StpxSpanningTreeObjects CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects

    
    StpxMISTPObjects CISCOSTPEXTENSIONSMIB_StpxMISTPObjects

    
    StpxLoopGuardObjects CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects

    
    StpxFastStartObjects CISCOSTPEXTENSIONSMIB_StpxFastStartObjects

    
    StpxBpduSkewingObjects CISCOSTPEXTENSIONSMIB_StpxBpduSkewingObjects

    
    StpxMSTObjects CISCOSTPEXTENSIONSMIB_StpxMSTObjects

    
    StpxRSTPObjects CISCOSTPEXTENSIONSMIB_StpxRSTPObjects

    
    StpxSMSTObjects CISCOSTPEXTENSIONSMIB_StpxSMSTObjects

    // A list of Virtual LAN entries containing information for Spanning Tree
    // PVST+ protocol.  An entry will exist for each VLAN existing on  the device.
    StpxPVSTVlanTable CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable

    // A table containing a list of the ports for which a particular VLAN's
    // Spanning Tree has been found to have an inconsistency.  Two types of
    // inconsistency are discovered: 1) an inconsistency where two different port
    // types have been plugged together; and 2) an inconsistency where different
    // switches have different PVIDs for the same link.
    StpxInconsistencyTable CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable

    // A table containing a list of the bridge ports for which Spanning Tree
    // RootGuard capability can be configured.
    StpxRootGuardConfigTable CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable

    // A table containing a list of the bridge ports for which a particular
    // Spanning Tree instance has been found  to have an root-inconsistency. The
    // agent creates a new  entry in this table whenever it detects a new 
    // root-inconsistency, and deletes entries  when/soon after the inconsistency
    // is no longer present.
    StpxRootInconsistencyTable CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable

    // This table contains one entry for each instance of MISTP and  it contains
    // stpxMISTPInstanceNumber entries, numbered from 1 to
    // stpxMISTPInstanceNumber.  This table is only instantiated when the value of
    // stpxSpanningTreeType is mistp(2) or mistpPvstPlus(3).
    StpxMISTPInstanceTable CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable

    // A table containing a list of the bridge ports for which Spanning Tree
    // LoopGuard capability can be configured.
    StpxLoopGuardConfigTable CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable

    // A table containing a list of the bridge ports for which a particular
    // Spanning Tree instance has been found to have a loop-inconsistency. The
    // agent creates a new entry in this table whenever it detects a new
    // loop-inconsistency, and deletes entries when/soon after the inconsistency
    // is no longer present.
    StpxLoopInconsistencyTable CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable

    // A table containing a list of the bridge ports for which Spanning Tree Port
    // Fast Start can be configured.
    StpxFastStartPortTable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable

    // A table containing a list of the bridge ports  for a particular Spanning
    // Tree Instance.
    StpxFastStartOperModeTable CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable

    // A table containing a list of the bridge ports for  which a particular
    // Spanning Tree instance has been  detected to have BPDU skewing occurred
    // since the  object value of stpxBpduSkewingDetectionEnable was last changed
    // to true(1).  The agent creates a new entry in this table whenever a port in
    // a particular Spanning Tree instance is  detected to be BPDU skewed since
    // the object value of  stpxBpduSkewingDetectionEnable object is changed to 
    // true(1). The agent deletes all the entries in this  table when the object
    // value of  stpxBpduSkewingDetectionEnable is changed to false(2) or the
    // object value of stpxSpanningTreeType is  changed.
    StpxBpduSkewingTable CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable

    // This table contains MST instance information with one entry for an MST
    // instance within the range of  0 to the object value of
    // stpxMSTMaxInstanceNumber.   This table is deprecated and replaced by 
    // stpxSMSTInstanceTable.
    StpxMSTInstanceTable CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable

    // This table contains MST instance information in the  Edit Buffer with one
    // entry for each MST instance numbered from 0 to stpxMSTMaxInstanceNumber.  
    // This table is only instantiated when the  stpxMSTRegionEditBufferStatus has
    // the value of acquiredBySnmp(2).  This table is deprecated and replaced by 
    // stpxSMSTInstanceEditTable.
    StpxMSTInstanceEditTable CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable

    // A table containing port information for the MST  Protocol on all the bridge
    // ports existing on the  system.
    StpxMSTPortTable CISCOSTPEXTENSIONSMIB_StpxMSTPortTable

    // A table containing a list of the bridge ports for a  particular MST
    // instance.  This table is only instantiated  when the stpxSpanningTreeType
    // is mst(4).   This table is deprecated and replaced with 
    // stpxRSTPPortRoleTable.
    StpxMSTPortRoleTable CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable

    // A table containing port information for the RSTP  Protocol on all the
    // bridge ports existing in the  system.
    StpxRSTPPortTable CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable

    // A table containing a list of the bridge ports for a  particular Spanning
    // Tree instance.  This table is  only instantiated when the
    // stpxSpanningTreeType is mst(4)  or rapidPvstPlus(5).
    StpxRSTPPortRoleTable CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable

    // A table containing a list of the bridge ports  for a particular Spanning
    // Tree Instance. This table is only instantiated when the object value of
    // stpxSpanningTreeType is rapidPvstPlus(5).
    StpxRPVSTPortTable CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable

    // This table contains MST instance information for IEEE MST.
    StpxSMSTInstanceTable CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable

    // This table contains MST instance information in the  Edit Buffer.   This
    // table is only instantiated when the object value of 
    // stpxMSTRegionEditBufferStatus has the value of acquiredBySnmp(2).
    StpxSMSTInstanceEditTable CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable

    // A table containing port information for the MST  Protocol on all the bridge
    // ports existing on the  system.  This table is only instantiated when the
    // object  value of stpxSpanningTreeType is mst(4).
    StpxSMSTPortTable CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable
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

    cISCOSTPEXTENSIONSMIB.EntityData.Children = types.NewOrderedMap()
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxUplinkFastObjects", types.YChild{"StpxUplinkFastObjects", &cISCOSTPEXTENSIONSMIB.StpxUplinkFastObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxBackboneFastObjects", types.YChild{"StpxBackboneFastObjects", &cISCOSTPEXTENSIONSMIB.StpxBackboneFastObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxSpanningTreeObjects", types.YChild{"StpxSpanningTreeObjects", &cISCOSTPEXTENSIONSMIB.StpxSpanningTreeObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMISTPObjects", types.YChild{"StpxMISTPObjects", &cISCOSTPEXTENSIONSMIB.StpxMISTPObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxLoopGuardObjects", types.YChild{"StpxLoopGuardObjects", &cISCOSTPEXTENSIONSMIB.StpxLoopGuardObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxFastStartObjects", types.YChild{"StpxFastStartObjects", &cISCOSTPEXTENSIONSMIB.StpxFastStartObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxBpduSkewingObjects", types.YChild{"StpxBpduSkewingObjects", &cISCOSTPEXTENSIONSMIB.StpxBpduSkewingObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMSTObjects", types.YChild{"StpxMSTObjects", &cISCOSTPEXTENSIONSMIB.StpxMSTObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRSTPObjects", types.YChild{"StpxRSTPObjects", &cISCOSTPEXTENSIONSMIB.StpxRSTPObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxSMSTObjects", types.YChild{"StpxSMSTObjects", &cISCOSTPEXTENSIONSMIB.StpxSMSTObjects})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxPVSTVlanTable", types.YChild{"StpxPVSTVlanTable", &cISCOSTPEXTENSIONSMIB.StpxPVSTVlanTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxInconsistencyTable", types.YChild{"StpxInconsistencyTable", &cISCOSTPEXTENSIONSMIB.StpxInconsistencyTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRootGuardConfigTable", types.YChild{"StpxRootGuardConfigTable", &cISCOSTPEXTENSIONSMIB.StpxRootGuardConfigTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRootInconsistencyTable", types.YChild{"StpxRootInconsistencyTable", &cISCOSTPEXTENSIONSMIB.StpxRootInconsistencyTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMISTPInstanceTable", types.YChild{"StpxMISTPInstanceTable", &cISCOSTPEXTENSIONSMIB.StpxMISTPInstanceTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxLoopGuardConfigTable", types.YChild{"StpxLoopGuardConfigTable", &cISCOSTPEXTENSIONSMIB.StpxLoopGuardConfigTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxLoopInconsistencyTable", types.YChild{"StpxLoopInconsistencyTable", &cISCOSTPEXTENSIONSMIB.StpxLoopInconsistencyTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxFastStartPortTable", types.YChild{"StpxFastStartPortTable", &cISCOSTPEXTENSIONSMIB.StpxFastStartPortTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxFastStartOperModeTable", types.YChild{"StpxFastStartOperModeTable", &cISCOSTPEXTENSIONSMIB.StpxFastStartOperModeTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxBpduSkewingTable", types.YChild{"StpxBpduSkewingTable", &cISCOSTPEXTENSIONSMIB.StpxBpduSkewingTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMSTInstanceTable", types.YChild{"StpxMSTInstanceTable", &cISCOSTPEXTENSIONSMIB.StpxMSTInstanceTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMSTInstanceEditTable", types.YChild{"StpxMSTInstanceEditTable", &cISCOSTPEXTENSIONSMIB.StpxMSTInstanceEditTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMSTPortTable", types.YChild{"StpxMSTPortTable", &cISCOSTPEXTENSIONSMIB.StpxMSTPortTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxMSTPortRoleTable", types.YChild{"StpxMSTPortRoleTable", &cISCOSTPEXTENSIONSMIB.StpxMSTPortRoleTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRSTPPortTable", types.YChild{"StpxRSTPPortTable", &cISCOSTPEXTENSIONSMIB.StpxRSTPPortTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRSTPPortRoleTable", types.YChild{"StpxRSTPPortRoleTable", &cISCOSTPEXTENSIONSMIB.StpxRSTPPortRoleTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxRPVSTPortTable", types.YChild{"StpxRPVSTPortTable", &cISCOSTPEXTENSIONSMIB.StpxRPVSTPortTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxSMSTInstanceTable", types.YChild{"StpxSMSTInstanceTable", &cISCOSTPEXTENSIONSMIB.StpxSMSTInstanceTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxSMSTInstanceEditTable", types.YChild{"StpxSMSTInstanceEditTable", &cISCOSTPEXTENSIONSMIB.StpxSMSTInstanceEditTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Children.Append("stpxSMSTPortTable", types.YChild{"StpxSMSTPortTable", &cISCOSTPEXTENSIONSMIB.StpxSMSTPortTable})
    cISCOSTPEXTENSIONSMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOSTPEXTENSIONSMIB.EntityData.YListKeys = []string {}

    return &(cISCOSTPEXTENSIONSMIB.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxUplinkFastObjects
type CISCOSTPEXTENSIONSMIB_StpxUplinkFastObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the UplinkFast capability is administratively
    // enabled on the device.  If the platform does not support configuration of
    // this object when the object value of stpxSpanningTreeType is  mst(4), then
    // this object is not instantiated. The type is bool.
    StpxUplinkFastEnabled interface{}

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
    StpxUplinkFastTransitions interface{}

    // The maximum number of station-learning frames that this device will
    // generate in each 100 milli-second period after a UplinkFast transition.  By
    // configuring this object, the network administrator can limit the rate at
    // which station-learning frames are generated.    If the platform does not
    // support configuration of this object when the object value of
    // stpxSpanningTreeType is mst(4), then this object is not instantiated. The
    // type is interface{} with range: 0..32000. Units are frames.
    StpxUplinkStationLearningGenRate interface{}

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
    StpxUplinkStationLearningFrames interface{}

    // An indication of whether the UplinkFast capability is  operationally
    // enabled on the device. The type is bool.
    StpxUplinkFastOperEnabled interface{}
}

func (stpxUplinkFastObjects *CISCOSTPEXTENSIONSMIB_StpxUplinkFastObjects) GetEntityData() *types.CommonEntityData {
    stpxUplinkFastObjects.EntityData.YFilter = stpxUplinkFastObjects.YFilter
    stpxUplinkFastObjects.EntityData.YangName = "stpxUplinkFastObjects"
    stpxUplinkFastObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxUplinkFastObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxUplinkFastObjects.EntityData.SegmentPath = "stpxUplinkFastObjects"
    stpxUplinkFastObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxUplinkFastObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxUplinkFastObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxUplinkFastObjects.EntityData.Children = types.NewOrderedMap()
    stpxUplinkFastObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxUplinkFastObjects.EntityData.Leafs.Append("stpxUplinkFastEnabled", types.YLeaf{"StpxUplinkFastEnabled", stpxUplinkFastObjects.StpxUplinkFastEnabled})
    stpxUplinkFastObjects.EntityData.Leafs.Append("stpxUplinkFastTransitions", types.YLeaf{"StpxUplinkFastTransitions", stpxUplinkFastObjects.StpxUplinkFastTransitions})
    stpxUplinkFastObjects.EntityData.Leafs.Append("stpxUplinkStationLearningGenRate", types.YLeaf{"StpxUplinkStationLearningGenRate", stpxUplinkFastObjects.StpxUplinkStationLearningGenRate})
    stpxUplinkFastObjects.EntityData.Leafs.Append("stpxUplinkStationLearningFrames", types.YLeaf{"StpxUplinkStationLearningFrames", stpxUplinkFastObjects.StpxUplinkStationLearningFrames})
    stpxUplinkFastObjects.EntityData.Leafs.Append("stpxUplinkFastOperEnabled", types.YLeaf{"StpxUplinkFastOperEnabled", stpxUplinkFastObjects.StpxUplinkFastOperEnabled})

    stpxUplinkFastObjects.EntityData.YListKeys = []string {}

    return &(stpxUplinkFastObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxBackboneFastObjects
type CISCOSTPEXTENSIONSMIB_StpxBackboneFastObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the BackboneFast capability is administratively
    // enabled on the device. The type is bool.
    StpxBackboneFastEnabled interface{}

    // The number of inferior BPDUs received by the switch  since the
    // stpxBackboneFastOperEnabled has become true(1). If the value of 
    // stpxBackboneFastOperEnabled is false(2), then this  mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    StpxBackboneFastInInferiorBPDUs interface{}

    // The number of Root Link Query request PDUs received by the switch since the
    // stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    StpxBackboneFastInRLQRequestPDUs interface{}

    // The number of Root Link Query response PDUs received by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    StpxBackboneFastInRLQResponsePDUs interface{}

    // The number of Root Link Query request PDUs transmitted by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    StpxBackboneFastOutRLQRequestPDUs interface{}

    // The number of Root Link Query response PDUs transmitted by the switch since
    // the stpxBackboneFastOperEnabled has become true(1). If the value of
    // stpxBackboneFastOperEnabled is false(2), then this mib object will have a
    // value of 0. The type is interface{} with range: 0..4294967295.
    StpxBackboneFastOutRLQResponsePDUs interface{}

    // An indication of whether the BackboneFast capability is operationally
    // enabled on the device. The type is bool.
    StpxBackboneFastOperEnabled interface{}
}

func (stpxBackboneFastObjects *CISCOSTPEXTENSIONSMIB_StpxBackboneFastObjects) GetEntityData() *types.CommonEntityData {
    stpxBackboneFastObjects.EntityData.YFilter = stpxBackboneFastObjects.YFilter
    stpxBackboneFastObjects.EntityData.YangName = "stpxBackboneFastObjects"
    stpxBackboneFastObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxBackboneFastObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxBackboneFastObjects.EntityData.SegmentPath = "stpxBackboneFastObjects"
    stpxBackboneFastObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxBackboneFastObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxBackboneFastObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxBackboneFastObjects.EntityData.Children = types.NewOrderedMap()
    stpxBackboneFastObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastEnabled", types.YLeaf{"StpxBackboneFastEnabled", stpxBackboneFastObjects.StpxBackboneFastEnabled})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastInInferiorBPDUs", types.YLeaf{"StpxBackboneFastInInferiorBPDUs", stpxBackboneFastObjects.StpxBackboneFastInInferiorBPDUs})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastInRLQRequestPDUs", types.YLeaf{"StpxBackboneFastInRLQRequestPDUs", stpxBackboneFastObjects.StpxBackboneFastInRLQRequestPDUs})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastInRLQResponsePDUs", types.YLeaf{"StpxBackboneFastInRLQResponsePDUs", stpxBackboneFastObjects.StpxBackboneFastInRLQResponsePDUs})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastOutRLQRequestPDUs", types.YLeaf{"StpxBackboneFastOutRLQRequestPDUs", stpxBackboneFastObjects.StpxBackboneFastOutRLQRequestPDUs})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastOutRLQResponsePDUs", types.YLeaf{"StpxBackboneFastOutRLQResponsePDUs", stpxBackboneFastObjects.StpxBackboneFastOutRLQResponsePDUs})
    stpxBackboneFastObjects.EntityData.Leafs.Append("stpxBackboneFastOperEnabled", types.YLeaf{"StpxBackboneFastOperEnabled", stpxBackboneFastObjects.StpxBackboneFastOperEnabled})

    stpxBackboneFastObjects.EntityData.YListKeys = []string {}

    return &(stpxBackboneFastObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects
type CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects struct {
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
    // also be lost  temporarily. The type is StpxSpanningTreeType.
    StpxSpanningTreeType interface{}

    // Indicates the administrative  spanning tree path cost mode  configured on
    // device. The type is StpxSpanningTreePathCostMode.
    StpxSpanningTreePathCostMode interface{}

    // Indicates whether Extended System ID feature  is administratively enabled
    // on the device or not. The type is bool.
    StpxExtendedSysIDAdminEnabled interface{}

    // Indicates whether Extended System ID feature  is operationaly enabled on
    // the device or not.  If the value of this object is true(1), then the
    // accepted values for dot1dStpPriority in BRIDGE-MIB should be multiples of
    // 4096 plus bridge instance ID, such as VlanIndex. Changing this object value
    // might cause the values of dot1dBaseBridgeAddress and dot1dStpPriority in
    // BRIDGE-MIB to be changed also. The type is bool.
    StpxExtendedSysIDOperEnabled interface{}

    // Indicates whether a specified notification is enabled or not. If a bit
    // corresponding to a notification is set to 1, then  the specified
    // notification can be generated.  newRoot -- the newRoot notification as
    // defined in BRIDGE-MIB.  topologyChange -- the topologyChange notification
    // as                   defined in BRIDGE-MIB.  inconsistency -- the
    // stpxInconsistencyUpdate notification.  rootInconsistency -- the
    // stpxRootInconsistencyUpdate                       notification. 
    // loopInconsistency -- the stpxLoopInconsistencyUpdate                      
    // notification. The type is map[string]bool.
    StpxNotificationEnable interface{}

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
    // BRIDGE-MIB must be used. The type is StpxSpanningTreePathCostOperMode.
    StpxSpanningTreePathCostOperMode interface{}
}

func (stpxSpanningTreeObjects *CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects) GetEntityData() *types.CommonEntityData {
    stpxSpanningTreeObjects.EntityData.YFilter = stpxSpanningTreeObjects.YFilter
    stpxSpanningTreeObjects.EntityData.YangName = "stpxSpanningTreeObjects"
    stpxSpanningTreeObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxSpanningTreeObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxSpanningTreeObjects.EntityData.SegmentPath = "stpxSpanningTreeObjects"
    stpxSpanningTreeObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSpanningTreeObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSpanningTreeObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSpanningTreeObjects.EntityData.Children = types.NewOrderedMap()
    stpxSpanningTreeObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxSpanningTreeType", types.YLeaf{"StpxSpanningTreeType", stpxSpanningTreeObjects.StpxSpanningTreeType})
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxSpanningTreePathCostMode", types.YLeaf{"StpxSpanningTreePathCostMode", stpxSpanningTreeObjects.StpxSpanningTreePathCostMode})
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxExtendedSysIDAdminEnabled", types.YLeaf{"StpxExtendedSysIDAdminEnabled", stpxSpanningTreeObjects.StpxExtendedSysIDAdminEnabled})
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxExtendedSysIDOperEnabled", types.YLeaf{"StpxExtendedSysIDOperEnabled", stpxSpanningTreeObjects.StpxExtendedSysIDOperEnabled})
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxNotificationEnable", types.YLeaf{"StpxNotificationEnable", stpxSpanningTreeObjects.StpxNotificationEnable})
    stpxSpanningTreeObjects.EntityData.Leafs.Append("stpxSpanningTreePathCostOperMode", types.YLeaf{"StpxSpanningTreePathCostOperMode", stpxSpanningTreeObjects.StpxSpanningTreePathCostOperMode})

    stpxSpanningTreeObjects.EntityData.YListKeys = []string {}

    return &(stpxSpanningTreeObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode represents configured on device.
type CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode_short CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode = "short"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode_long CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostMode = "long"
)

// CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode represents the dot1dStpPortPathCost in BRIDGE-MIB must be used.
type CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode_short CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode = "short"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode_long CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreePathCostOperMode = "long"
)

// CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType represents temporarily.
type CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType string

const (
    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType_pvstPlus CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType = "pvstPlus"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType_mistp CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType = "mistp"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType_mistpPvstPlus CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType = "mistpPvstPlus"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType_mst CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType = "mst"

    CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType_rapidPvstPlus CISCOSTPEXTENSIONSMIB_StpxSpanningTreeObjects_StpxSpanningTreeType = "rapidPvstPlus"
)

// CISCOSTPEXTENSIONSMIB_StpxMISTPObjects
type CISCOSTPEXTENSIONSMIB_StpxMISTPObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of MISTP instances, that are supported by the device  when the
    // value of stpxSpanningTreeType is either mistp(2) or mistpPvstPlus(3). The
    // type is interface{} with range: 1..256.
    StpxMISTPInstanceNumber interface{}
}

func (stpxMISTPObjects *CISCOSTPEXTENSIONSMIB_StpxMISTPObjects) GetEntityData() *types.CommonEntityData {
    stpxMISTPObjects.EntityData.YFilter = stpxMISTPObjects.YFilter
    stpxMISTPObjects.EntityData.YangName = "stpxMISTPObjects"
    stpxMISTPObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxMISTPObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMISTPObjects.EntityData.SegmentPath = "stpxMISTPObjects"
    stpxMISTPObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMISTPObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMISTPObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMISTPObjects.EntityData.Children = types.NewOrderedMap()
    stpxMISTPObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxMISTPObjects.EntityData.Leafs.Append("stpxMISTPInstanceNumber", types.YLeaf{"StpxMISTPInstanceNumber", stpxMISTPObjects.StpxMISTPInstanceNumber})

    stpxMISTPObjects.EntityData.YListKeys = []string {}

    return &(stpxMISTPObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects
type CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the global default config mode of LoopGuard  feature on the
    // device. The type is StpxLoopGuardGlobalDefaultMode.
    StpxLoopGuardGlobalDefaultMode interface{}
}

func (stpxLoopGuardObjects *CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects) GetEntityData() *types.CommonEntityData {
    stpxLoopGuardObjects.EntityData.YFilter = stpxLoopGuardObjects.YFilter
    stpxLoopGuardObjects.EntityData.YangName = "stpxLoopGuardObjects"
    stpxLoopGuardObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxLoopGuardObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxLoopGuardObjects.EntityData.SegmentPath = "stpxLoopGuardObjects"
    stpxLoopGuardObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxLoopGuardObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxLoopGuardObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxLoopGuardObjects.EntityData.Children = types.NewOrderedMap()
    stpxLoopGuardObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxLoopGuardObjects.EntityData.Leafs.Append("stpxLoopGuardGlobalDefaultMode", types.YLeaf{"StpxLoopGuardGlobalDefaultMode", stpxLoopGuardObjects.StpxLoopGuardGlobalDefaultMode})

    stpxLoopGuardObjects.EntityData.YListKeys = []string {}

    return &(stpxLoopGuardObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode represents feature on the device.
type CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode_enable CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode_disable CISCOSTPEXTENSIONSMIB_StpxLoopGuardObjects_StpxLoopGuardGlobalDefaultMode = "disable"
)

// CISCOSTPEXTENSIONSMIB_StpxFastStartObjects
type CISCOSTPEXTENSIONSMIB_StpxFastStartObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the global default mode of the Bpdu Guard feature on the device. 
    // On platforms that does not support per port  Bpdu Guard configuration as
    // indicated by the object stpxFastStartPortBpduGuardMode, if  the value of
    // this object is set to true(1),  and the Fast Start Feature is operationally
    // enabled on a port, then that port will be  immediately disabled when the
    // system receives a BPDU from that port. The type is bool.
    StpxFastStartBpduGuardEnable interface{}

    // Indicates the global default mode of the Bpdu  Filter feature on the
    // device.  On platforms that does not support per port  Bpdu Filter
    // configuration as indicated by the object stpxFastStartPortBpduFilterMode,
    // if  the value of this object is set to true(1),  and the Fast Start Feature
    // is operationally  enabled on a port, then no BPDUs will be  transmitted on
    // this port. The type is bool.
    StpxFastStartBpduFilterEnable interface{}

    // Indicates the global default mode of the Fast  Start feature on the device.
    // The type is StpxFastStartGlobalDefaultMode.
    StpxFastStartGlobalDefaultMode interface{}
}

func (stpxFastStartObjects *CISCOSTPEXTENSIONSMIB_StpxFastStartObjects) GetEntityData() *types.CommonEntityData {
    stpxFastStartObjects.EntityData.YFilter = stpxFastStartObjects.YFilter
    stpxFastStartObjects.EntityData.YangName = "stpxFastStartObjects"
    stpxFastStartObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxFastStartObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxFastStartObjects.EntityData.SegmentPath = "stpxFastStartObjects"
    stpxFastStartObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxFastStartObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxFastStartObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxFastStartObjects.EntityData.Children = types.NewOrderedMap()
    stpxFastStartObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxFastStartObjects.EntityData.Leafs.Append("stpxFastStartBpduGuardEnable", types.YLeaf{"StpxFastStartBpduGuardEnable", stpxFastStartObjects.StpxFastStartBpduGuardEnable})
    stpxFastStartObjects.EntityData.Leafs.Append("stpxFastStartBpduFilterEnable", types.YLeaf{"StpxFastStartBpduFilterEnable", stpxFastStartObjects.StpxFastStartBpduFilterEnable})
    stpxFastStartObjects.EntityData.Leafs.Append("stpxFastStartGlobalDefaultMode", types.YLeaf{"StpxFastStartGlobalDefaultMode", stpxFastStartObjects.StpxFastStartGlobalDefaultMode})

    stpxFastStartObjects.EntityData.YListKeys = []string {}

    return &(stpxFastStartObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode represents Start feature on the device.
type CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode_enable CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode_disable CISCOSTPEXTENSIONSMIB_StpxFastStartObjects_StpxFastStartGlobalDefaultMode = "disable"
)

// CISCOSTPEXTENSIONSMIB_StpxBpduSkewingObjects
type CISCOSTPEXTENSIONSMIB_StpxBpduSkewingObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates whether BPDU skewing detection feature is enabled or not on the
    // system. If this object has the value of true(1), then the system will
    // detect whether BPDUs received by any port on any Spanning  Tree instance
    // are processed at an interval longer than the object value of
    // dot1dStpHelloTime in the BIRDGE-MIB of the Spanning Tree instance. The type
    // is bool.
    StpxBpduSkewingDetectionEnable interface{}
}

func (stpxBpduSkewingObjects *CISCOSTPEXTENSIONSMIB_StpxBpduSkewingObjects) GetEntityData() *types.CommonEntityData {
    stpxBpduSkewingObjects.EntityData.YFilter = stpxBpduSkewingObjects.YFilter
    stpxBpduSkewingObjects.EntityData.YangName = "stpxBpduSkewingObjects"
    stpxBpduSkewingObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxBpduSkewingObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxBpduSkewingObjects.EntityData.SegmentPath = "stpxBpduSkewingObjects"
    stpxBpduSkewingObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxBpduSkewingObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxBpduSkewingObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxBpduSkewingObjects.EntityData.Children = types.NewOrderedMap()
    stpxBpduSkewingObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxBpduSkewingObjects.EntityData.Leafs.Append("stpxBpduSkewingDetectionEnable", types.YLeaf{"StpxBpduSkewingDetectionEnable", stpxBpduSkewingObjects.StpxBpduSkewingDetectionEnable})

    stpxBpduSkewingObjects.EntityData.YListKeys = []string {}

    return &(stpxBpduSkewingObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTObjects
type CISCOSTPEXTENSIONSMIB_StpxMSTObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum MST (Multiple Spanning Tree) instance id,  that can be
    // supported by the device for Cisco proprietary implementation of the MST
    // Protocol.  This object is deprecated and replaced by 
    // stpxSMSTMaxInstanceID. The type is interface{} with range: 1..256.
    StpxMSTMaxInstanceNumber interface{}

    // The operational MST region name. The type is string with length: 0..32.
    StpxMSTRegionName interface{}

    // The operational MST region version.  This object is deprecated and replaced
    // by  stpxSMSTRegionRevision. The type is interface{} with range: 0..65535.
    StpxMSTRegionRevision interface{}

    // Indicates the current ownership status of the unique  Region Config Edit
    // Buffer.   released -- the Edit Buffer can be acquired by any of            
    // the SNMP management stations.   acquiredBySnmp -- the Edit Buffer is
    // acquired by             any of the SNMP management stations.  
    // acquiredByNonSnmp -- the Edit Buffer is acquired by the             
    // non-SNMP users managing the device. The type is
    // StpxMSTRegionEditBufferStatus.
    StpxMSTRegionEditBufferStatus interface{}

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
    // other(1) when it is read. The type is StpxMSTRegionEditBufferOperation.
    StpxMSTRegionEditBufferOperation interface{}

    // The MST region name in the Edit Buffer.   This object is only instantiated
    // when the  stpxMSTRegionEditBufferStatus has the value of 
    // acquiredBySnmp(2). The type is string with length: 0..32.
    StpxMSTRegionEditName interface{}

    // The MST region version in the Edit Buffer. This object is only instantiated
    // when the stpxMSTRegionEditBufferStatus  has the value of acquiredBySnmp(2).
    // This object is deprecated and replaced by stpxSMSTRegionEditRevision. The
    // type is interface{} with range: 1..65535.
    StpxMSTRegionEditRevision interface{}

    // The maximum number of hops for the MST region.   This object will take on
    // value of 40 if the object value of stpxSMSTMaxHopCount is greater than 40. 
    // This object is deprecated and replaced by stpxSMSTMaxHopCount. The type is
    // interface{} with range: 1..40.
    StpxMSTMaxHopCount interface{}
}

func (stpxMSTObjects *CISCOSTPEXTENSIONSMIB_StpxMSTObjects) GetEntityData() *types.CommonEntityData {
    stpxMSTObjects.EntityData.YFilter = stpxMSTObjects.YFilter
    stpxMSTObjects.EntityData.YangName = "stpxMSTObjects"
    stpxMSTObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMSTObjects.EntityData.SegmentPath = "stpxMSTObjects"
    stpxMSTObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTObjects.EntityData.Children = types.NewOrderedMap()
    stpxMSTObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTMaxInstanceNumber", types.YLeaf{"StpxMSTMaxInstanceNumber", stpxMSTObjects.StpxMSTMaxInstanceNumber})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionName", types.YLeaf{"StpxMSTRegionName", stpxMSTObjects.StpxMSTRegionName})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionRevision", types.YLeaf{"StpxMSTRegionRevision", stpxMSTObjects.StpxMSTRegionRevision})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionEditBufferStatus", types.YLeaf{"StpxMSTRegionEditBufferStatus", stpxMSTObjects.StpxMSTRegionEditBufferStatus})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionEditBufferOperation", types.YLeaf{"StpxMSTRegionEditBufferOperation", stpxMSTObjects.StpxMSTRegionEditBufferOperation})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionEditName", types.YLeaf{"StpxMSTRegionEditName", stpxMSTObjects.StpxMSTRegionEditName})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTRegionEditRevision", types.YLeaf{"StpxMSTRegionEditRevision", stpxMSTObjects.StpxMSTRegionEditRevision})
    stpxMSTObjects.EntityData.Leafs.Append("stpxMSTMaxHopCount", types.YLeaf{"StpxMSTMaxHopCount", stpxMSTObjects.StpxMSTMaxHopCount})

    stpxMSTObjects.EntityData.YListKeys = []string {}

    return &(stpxMSTObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation represents This object always returns other(1) when it is read.
type CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation string

const (
    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation_other CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation = "other"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation_acquire CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation = "acquire"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation_releaseWithForce CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation = "releaseWithForce"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation_commit CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation = "commit"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation_rollBack CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferOperation = "rollBack"
)

// CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus represents             non-SNMP users managing the device.
type CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus string

const (
    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus_released CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus = "released"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus_acquiredBySnmp CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus = "acquiredBySnmp"

    CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus_acquiredByNonSnmp CISCOSTPEXTENSIONSMIB_StpxMSTObjects_StpxMSTRegionEditBufferStatus = "acquiredByNonSnmp"
)

// CISCOSTPEXTENSIONSMIB_StpxRSTPObjects
type CISCOSTPEXTENSIONSMIB_StpxRSTPObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Transmit Hold Count. The type is interface{} with range: 0..4294967295.
    StpxRSTPTransmitHoldCount interface{}
}

func (stpxRSTPObjects *CISCOSTPEXTENSIONSMIB_StpxRSTPObjects) GetEntityData() *types.CommonEntityData {
    stpxRSTPObjects.EntityData.YFilter = stpxRSTPObjects.YFilter
    stpxRSTPObjects.EntityData.YangName = "stpxRSTPObjects"
    stpxRSTPObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxRSTPObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRSTPObjects.EntityData.SegmentPath = "stpxRSTPObjects"
    stpxRSTPObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRSTPObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRSTPObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRSTPObjects.EntityData.Children = types.NewOrderedMap()
    stpxRSTPObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxRSTPObjects.EntityData.Leafs.Append("stpxRSTPTransmitHoldCount", types.YLeaf{"StpxRSTPTransmitHoldCount", stpxRSTPObjects.StpxRSTPTransmitHoldCount})

    stpxRSTPObjects.EntityData.YListKeys = []string {}

    return &(stpxRSTPObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTObjects
type CISCOSTPEXTENSIONSMIB_StpxSMSTObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of MST instances that can be  supported by the device
    // for IEEE MST. The type is interface{} with range: 0..4294967295.
    StpxSMSTMaxInstances interface{}

    // The maximum MST instance ID that can be supported  by the device for IEEE
    // MST. The type is interface{} with range: 0..4294967295.
    StpxSMSTMaxInstanceID interface{}

    // The operational region version for IEEE MST. The type is interface{} with
    // range: 0..4294967295.
    StpxSMSTRegionRevision interface{}

    // The MST region version in the Edit Buffer for IEEE  MST.  This object is
    // only instantiated when the  stpxMSTRegionEditBufferStatus has the value of 
    // acquiredBySnmp(2). The type is interface{} with range: 0..4294967295.
    StpxSMSTRegionEditRevision interface{}

    // The maximum number of hops for the IEEE MST region. The type is interface{}
    // with range: 0..4294967295.
    StpxSMSTMaxHopCount interface{}

    // The IEEE MST region configuration digest. The type is string.
    StpxSMSTConfigDigest interface{}

    // The pre-standard MST region configuration digest. The type is string.
    StpxSMSTConfigPreStandardDigest interface{}
}

func (stpxSMSTObjects *CISCOSTPEXTENSIONSMIB_StpxSMSTObjects) GetEntityData() *types.CommonEntityData {
    stpxSMSTObjects.EntityData.YFilter = stpxSMSTObjects.YFilter
    stpxSMSTObjects.EntityData.YangName = "stpxSMSTObjects"
    stpxSMSTObjects.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTObjects.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxSMSTObjects.EntityData.SegmentPath = "stpxSMSTObjects"
    stpxSMSTObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTObjects.EntityData.Children = types.NewOrderedMap()
    stpxSMSTObjects.EntityData.Leafs = types.NewOrderedMap()
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTMaxInstances", types.YLeaf{"StpxSMSTMaxInstances", stpxSMSTObjects.StpxSMSTMaxInstances})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTMaxInstanceID", types.YLeaf{"StpxSMSTMaxInstanceID", stpxSMSTObjects.StpxSMSTMaxInstanceID})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTRegionRevision", types.YLeaf{"StpxSMSTRegionRevision", stpxSMSTObjects.StpxSMSTRegionRevision})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTRegionEditRevision", types.YLeaf{"StpxSMSTRegionEditRevision", stpxSMSTObjects.StpxSMSTRegionEditRevision})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTMaxHopCount", types.YLeaf{"StpxSMSTMaxHopCount", stpxSMSTObjects.StpxSMSTMaxHopCount})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTConfigDigest", types.YLeaf{"StpxSMSTConfigDigest", stpxSMSTObjects.StpxSMSTConfigDigest})
    stpxSMSTObjects.EntityData.Leafs.Append("stpxSMSTConfigPreStandardDigest", types.YLeaf{"StpxSMSTConfigPreStandardDigest", stpxSMSTObjects.StpxSMSTConfigPreStandardDigest})

    stpxSMSTObjects.EntityData.YListKeys = []string {}

    return &(stpxSMSTObjects.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable
// A list of Virtual LAN entries containing
// information for Spanning Tree PVST+ protocol. 
// An entry will exist for each VLAN existing on 
// the device.
type CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing Spanning Tree PVST+ Protocol  information for a
    // particular Virtual LAN. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry.
    StpxPVSTVlanEntry []*CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry
}

func (stpxPVSTVlanTable *CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable) GetEntityData() *types.CommonEntityData {
    stpxPVSTVlanTable.EntityData.YFilter = stpxPVSTVlanTable.YFilter
    stpxPVSTVlanTable.EntityData.YangName = "stpxPVSTVlanTable"
    stpxPVSTVlanTable.EntityData.BundleName = "cisco_ios_xe"
    stpxPVSTVlanTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxPVSTVlanTable.EntityData.SegmentPath = "stpxPVSTVlanTable"
    stpxPVSTVlanTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxPVSTVlanTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxPVSTVlanTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxPVSTVlanTable.EntityData.Children = types.NewOrderedMap()
    stpxPVSTVlanTable.EntityData.Children.Append("stpxPVSTVlanEntry", types.YChild{"StpxPVSTVlanEntry", nil})
    for i := range stpxPVSTVlanTable.StpxPVSTVlanEntry {
        stpxPVSTVlanTable.EntityData.Children.Append(types.GetSegmentPath(stpxPVSTVlanTable.StpxPVSTVlanEntry[i]), types.YChild{"StpxPVSTVlanEntry", stpxPVSTVlanTable.StpxPVSTVlanEntry[i]})
    }
    stpxPVSTVlanTable.EntityData.Leafs = types.NewOrderedMap()

    stpxPVSTVlanTable.EntityData.YListKeys = []string {}

    return &(stpxPVSTVlanTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry
// An entry containing Spanning Tree PVST+ Protocol 
// information for a particular Virtual LAN.
type CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index value that uniquely identifies the
    // Virtual LAN associated with this information. The type is interface{} with
    // range: 0..4095.
    StpxPVSTVlanIndex interface{}

    // Indicates whether Spanning Tree PVST+   Protocol is enabled for this
    // Virtual LAN. If  Spanning Tree PVST+ Protocol is not supported  on this
    // VLAN, then notApplicable(3) will be  returned while retrieving the object
    // value for  this VLAN.  If the device only supports a single global Spanning
    // Tree PVST+ Protocol enable/disable  for all the existing VLANs, then the
    // object  value assigned to this VLAN will be applied to the object values of
    // all the instances in this table which do not have the value of
    // notApplicable(3).  If the value of stpxSpanningTreeType is neither 
    // pvstPlus(1) nor rapidPvstPlus(5), then the value  of stpxPVSTVlanEnable for
    // this VLAN can not be  changed. The type is StpxPVSTVlanEnable.
    StpxPVSTVlanEnable interface{}
}

func (stpxPVSTVlanEntry *CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry) GetEntityData() *types.CommonEntityData {
    stpxPVSTVlanEntry.EntityData.YFilter = stpxPVSTVlanEntry.YFilter
    stpxPVSTVlanEntry.EntityData.YangName = "stpxPVSTVlanEntry"
    stpxPVSTVlanEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxPVSTVlanEntry.EntityData.ParentYangName = "stpxPVSTVlanTable"
    stpxPVSTVlanEntry.EntityData.SegmentPath = "stpxPVSTVlanEntry" + types.AddKeyToken(stpxPVSTVlanEntry.StpxPVSTVlanIndex, "stpxPVSTVlanIndex")
    stpxPVSTVlanEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxPVSTVlanEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxPVSTVlanEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxPVSTVlanEntry.EntityData.Children = types.NewOrderedMap()
    stpxPVSTVlanEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxPVSTVlanEntry.EntityData.Leafs.Append("stpxPVSTVlanIndex", types.YLeaf{"StpxPVSTVlanIndex", stpxPVSTVlanEntry.StpxPVSTVlanIndex})
    stpxPVSTVlanEntry.EntityData.Leafs.Append("stpxPVSTVlanEnable", types.YLeaf{"StpxPVSTVlanEnable", stpxPVSTVlanEntry.StpxPVSTVlanEnable})

    stpxPVSTVlanEntry.EntityData.YListKeys = []string {"StpxPVSTVlanIndex"}

    return &(stpxPVSTVlanEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable represents changed.
type CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable string

const (
    CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable_enabled CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable = "enabled"

    CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable_disabled CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable = "disabled"

    CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable_notApplicable CISCOSTPEXTENSIONSMIB_StpxPVSTVlanTable_StpxPVSTVlanEntry_StpxPVSTVlanEnable = "notApplicable"
)

// CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable
// A table containing a list of the ports for which
// a particular VLAN's Spanning Tree has been found to
// have an inconsistency.  Two types of inconsistency
// are discovered: 1) an inconsistency where two different
// port types have been plugged together; and 2) an
// inconsistency where different switches have different
// PVIDs for the same link.
type CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VLAN on a particular port for which a Spanning Tree inconsistency is
    // currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable_StpxInconsistencyEntry.
    StpxInconsistencyEntry []*CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable_StpxInconsistencyEntry
}

func (stpxInconsistencyTable *CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable) GetEntityData() *types.CommonEntityData {
    stpxInconsistencyTable.EntityData.YFilter = stpxInconsistencyTable.YFilter
    stpxInconsistencyTable.EntityData.YangName = "stpxInconsistencyTable"
    stpxInconsistencyTable.EntityData.BundleName = "cisco_ios_xe"
    stpxInconsistencyTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxInconsistencyTable.EntityData.SegmentPath = "stpxInconsistencyTable"
    stpxInconsistencyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxInconsistencyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxInconsistencyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxInconsistencyTable.EntityData.Children = types.NewOrderedMap()
    stpxInconsistencyTable.EntityData.Children.Append("stpxInconsistencyEntry", types.YChild{"StpxInconsistencyEntry", nil})
    for i := range stpxInconsistencyTable.StpxInconsistencyEntry {
        stpxInconsistencyTable.EntityData.Children.Append(types.GetSegmentPath(stpxInconsistencyTable.StpxInconsistencyEntry[i]), types.YChild{"StpxInconsistencyEntry", stpxInconsistencyTable.StpxInconsistencyEntry[i]})
    }
    stpxInconsistencyTable.EntityData.Leafs = types.NewOrderedMap()

    stpxInconsistencyTable.EntityData.YListKeys = []string {}

    return &(stpxInconsistencyTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable_StpxInconsistencyEntry
// A VLAN on a particular port for which a Spanning Tree
// inconsistency is currently in effect.
type CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable_StpxInconsistencyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    StpxVlanIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxPortIndex interface{}

    // The types of inconsistency which have been discovered on this port for this
    // VLAN's Spanning Tree.  When this object exists, the value of the
    // corresponding instance of the Bridge MIB's dot1dStpPortState object will be
    // 'broken(6)'. The type is map[string]bool.
    StpxInconsistentState interface{}
}

func (stpxInconsistencyEntry *CISCOSTPEXTENSIONSMIB_StpxInconsistencyTable_StpxInconsistencyEntry) GetEntityData() *types.CommonEntityData {
    stpxInconsistencyEntry.EntityData.YFilter = stpxInconsistencyEntry.YFilter
    stpxInconsistencyEntry.EntityData.YangName = "stpxInconsistencyEntry"
    stpxInconsistencyEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxInconsistencyEntry.EntityData.ParentYangName = "stpxInconsistencyTable"
    stpxInconsistencyEntry.EntityData.SegmentPath = "stpxInconsistencyEntry" + types.AddKeyToken(stpxInconsistencyEntry.StpxVlanIndex, "stpxVlanIndex") + types.AddKeyToken(stpxInconsistencyEntry.StpxPortIndex, "stpxPortIndex")
    stpxInconsistencyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxInconsistencyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxInconsistencyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxInconsistencyEntry.EntityData.Children = types.NewOrderedMap()
    stpxInconsistencyEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxInconsistencyEntry.EntityData.Leafs.Append("stpxVlanIndex", types.YLeaf{"StpxVlanIndex", stpxInconsistencyEntry.StpxVlanIndex})
    stpxInconsistencyEntry.EntityData.Leafs.Append("stpxPortIndex", types.YLeaf{"StpxPortIndex", stpxInconsistencyEntry.StpxPortIndex})
    stpxInconsistencyEntry.EntityData.Leafs.Append("stpxInconsistentState", types.YLeaf{"StpxInconsistentState", stpxInconsistencyEntry.StpxInconsistentState})

    stpxInconsistencyEntry.EntityData.YListKeys = []string {"StpxVlanIndex", "StpxPortIndex"}

    return &(stpxInconsistencyEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable
// A table containing a list of the bridge ports for which
// Spanning Tree RootGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree RootGuard capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable_StpxRootGuardConfigEntry.
    StpxRootGuardConfigEntry []*CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable_StpxRootGuardConfigEntry
}

func (stpxRootGuardConfigTable *CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable) GetEntityData() *types.CommonEntityData {
    stpxRootGuardConfigTable.EntityData.YFilter = stpxRootGuardConfigTable.YFilter
    stpxRootGuardConfigTable.EntityData.YangName = "stpxRootGuardConfigTable"
    stpxRootGuardConfigTable.EntityData.BundleName = "cisco_ios_xe"
    stpxRootGuardConfigTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRootGuardConfigTable.EntityData.SegmentPath = "stpxRootGuardConfigTable"
    stpxRootGuardConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRootGuardConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRootGuardConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRootGuardConfigTable.EntityData.Children = types.NewOrderedMap()
    stpxRootGuardConfigTable.EntityData.Children.Append("stpxRootGuardConfigEntry", types.YChild{"StpxRootGuardConfigEntry", nil})
    for i := range stpxRootGuardConfigTable.StpxRootGuardConfigEntry {
        stpxRootGuardConfigTable.EntityData.Children.Append(types.GetSegmentPath(stpxRootGuardConfigTable.StpxRootGuardConfigEntry[i]), types.YChild{"StpxRootGuardConfigEntry", stpxRootGuardConfigTable.StpxRootGuardConfigEntry[i]})
    }
    stpxRootGuardConfigTable.EntityData.Leafs = types.NewOrderedMap()

    stpxRootGuardConfigTable.EntityData.YListKeys = []string {}

    return &(stpxRootGuardConfigTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable_StpxRootGuardConfigEntry
// A bridge port for which Spanning Tree RootGuard
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable_StpxRootGuardConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxRootGuardConfigPortIndex interface{}

    // An indication of whether the RootGuard capability is  enabled on this port
    // or not. This configuration will be applied to all Spanning Tree instances
    // in which this port  exists. The type is bool.
    StpxRootGuardConfigEnabled interface{}
}

func (stpxRootGuardConfigEntry *CISCOSTPEXTENSIONSMIB_StpxRootGuardConfigTable_StpxRootGuardConfigEntry) GetEntityData() *types.CommonEntityData {
    stpxRootGuardConfigEntry.EntityData.YFilter = stpxRootGuardConfigEntry.YFilter
    stpxRootGuardConfigEntry.EntityData.YangName = "stpxRootGuardConfigEntry"
    stpxRootGuardConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxRootGuardConfigEntry.EntityData.ParentYangName = "stpxRootGuardConfigTable"
    stpxRootGuardConfigEntry.EntityData.SegmentPath = "stpxRootGuardConfigEntry" + types.AddKeyToken(stpxRootGuardConfigEntry.StpxRootGuardConfigPortIndex, "stpxRootGuardConfigPortIndex")
    stpxRootGuardConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRootGuardConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRootGuardConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRootGuardConfigEntry.EntityData.Children = types.NewOrderedMap()
    stpxRootGuardConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxRootGuardConfigEntry.EntityData.Leafs.Append("stpxRootGuardConfigPortIndex", types.YLeaf{"StpxRootGuardConfigPortIndex", stpxRootGuardConfigEntry.StpxRootGuardConfigPortIndex})
    stpxRootGuardConfigEntry.EntityData.Leafs.Append("stpxRootGuardConfigEnabled", types.YLeaf{"StpxRootGuardConfigEnabled", stpxRootGuardConfigEntry.StpxRootGuardConfigEnabled})

    stpxRootGuardConfigEntry.EntityData.YListKeys = []string {"StpxRootGuardConfigPortIndex"}

    return &(stpxRootGuardConfigEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable
// A table containing a list of the bridge ports for which
// a particular Spanning Tree instance has been found 
// to have an root-inconsistency. The agent creates a new 
// entry in this table whenever it detects a new 
// root-inconsistency, and deletes entries 
// when/soon after the inconsistency is no longer present.
type CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for  which a Spanning Tree
    // root-inconsistency is currently  in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable_StpxRootInconsistencyEntry.
    StpxRootInconsistencyEntry []*CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable_StpxRootInconsistencyEntry
}

func (stpxRootInconsistencyTable *CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable) GetEntityData() *types.CommonEntityData {
    stpxRootInconsistencyTable.EntityData.YFilter = stpxRootInconsistencyTable.YFilter
    stpxRootInconsistencyTable.EntityData.YangName = "stpxRootInconsistencyTable"
    stpxRootInconsistencyTable.EntityData.BundleName = "cisco_ios_xe"
    stpxRootInconsistencyTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRootInconsistencyTable.EntityData.SegmentPath = "stpxRootInconsistencyTable"
    stpxRootInconsistencyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRootInconsistencyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRootInconsistencyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRootInconsistencyTable.EntityData.Children = types.NewOrderedMap()
    stpxRootInconsistencyTable.EntityData.Children.Append("stpxRootInconsistencyEntry", types.YChild{"StpxRootInconsistencyEntry", nil})
    for i := range stpxRootInconsistencyTable.StpxRootInconsistencyEntry {
        stpxRootInconsistencyTable.EntityData.Children.Append(types.GetSegmentPath(stpxRootInconsistencyTable.StpxRootInconsistencyEntry[i]), types.YChild{"StpxRootInconsistencyEntry", stpxRootInconsistencyTable.StpxRootInconsistencyEntry[i]})
    }
    stpxRootInconsistencyTable.EntityData.Leafs = types.NewOrderedMap()

    stpxRootInconsistencyTable.EntityData.YListKeys = []string {}

    return &(stpxRootInconsistencyTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable_StpxRootInconsistencyEntry
// A Spanning Tree instance on a particular port for 
// which a Spanning Tree root-inconsistency is currently 
// in effect.
type CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable_StpxRootInconsistencyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1) or
    // rapidPvstPlus(5). The type is interface{} with range: 0..65535.
    StpxRootInconsistencyIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxRootInconsistencyPortIndex interface{}

    // Indicates whether the port on a particular Spanning  Tree instance is
    // currently in root-inconsistent  state or not. The type is bool.
    StpxRootInconsistencyState interface{}
}

func (stpxRootInconsistencyEntry *CISCOSTPEXTENSIONSMIB_StpxRootInconsistencyTable_StpxRootInconsistencyEntry) GetEntityData() *types.CommonEntityData {
    stpxRootInconsistencyEntry.EntityData.YFilter = stpxRootInconsistencyEntry.YFilter
    stpxRootInconsistencyEntry.EntityData.YangName = "stpxRootInconsistencyEntry"
    stpxRootInconsistencyEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxRootInconsistencyEntry.EntityData.ParentYangName = "stpxRootInconsistencyTable"
    stpxRootInconsistencyEntry.EntityData.SegmentPath = "stpxRootInconsistencyEntry" + types.AddKeyToken(stpxRootInconsistencyEntry.StpxRootInconsistencyIndex, "stpxRootInconsistencyIndex") + types.AddKeyToken(stpxRootInconsistencyEntry.StpxRootInconsistencyPortIndex, "stpxRootInconsistencyPortIndex")
    stpxRootInconsistencyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRootInconsistencyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRootInconsistencyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRootInconsistencyEntry.EntityData.Children = types.NewOrderedMap()
    stpxRootInconsistencyEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxRootInconsistencyEntry.EntityData.Leafs.Append("stpxRootInconsistencyIndex", types.YLeaf{"StpxRootInconsistencyIndex", stpxRootInconsistencyEntry.StpxRootInconsistencyIndex})
    stpxRootInconsistencyEntry.EntityData.Leafs.Append("stpxRootInconsistencyPortIndex", types.YLeaf{"StpxRootInconsistencyPortIndex", stpxRootInconsistencyEntry.StpxRootInconsistencyPortIndex})
    stpxRootInconsistencyEntry.EntityData.Leafs.Append("stpxRootInconsistencyState", types.YLeaf{"StpxRootInconsistencyState", stpxRootInconsistencyEntry.StpxRootInconsistencyState})

    stpxRootInconsistencyEntry.EntityData.YListKeys = []string {"StpxRootInconsistencyIndex", "StpxRootInconsistencyPortIndex"}

    return &(stpxRootInconsistencyEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable
// This table contains one entry for each instance of MISTP and 
// it contains stpxMISTPInstanceNumber entries, numbered from 1
// to stpxMISTPInstanceNumber.
// 
// This table is only instantiated when the value of 
// stpxSpanningTreeType is mistp(2) or mistpPvstPlus(3).
type CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the status of the MISTP  instance. The type is
    // slice of
    // CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable_StpxMISTPInstanceEntry.
    StpxMISTPInstanceEntry []*CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable_StpxMISTPInstanceEntry
}

func (stpxMISTPInstanceTable *CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable) GetEntityData() *types.CommonEntityData {
    stpxMISTPInstanceTable.EntityData.YFilter = stpxMISTPInstanceTable.YFilter
    stpxMISTPInstanceTable.EntityData.YangName = "stpxMISTPInstanceTable"
    stpxMISTPInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    stpxMISTPInstanceTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMISTPInstanceTable.EntityData.SegmentPath = "stpxMISTPInstanceTable"
    stpxMISTPInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMISTPInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMISTPInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMISTPInstanceTable.EntityData.Children = types.NewOrderedMap()
    stpxMISTPInstanceTable.EntityData.Children.Append("stpxMISTPInstanceEntry", types.YChild{"StpxMISTPInstanceEntry", nil})
    for i := range stpxMISTPInstanceTable.StpxMISTPInstanceEntry {
        stpxMISTPInstanceTable.EntityData.Children.Append(types.GetSegmentPath(stpxMISTPInstanceTable.StpxMISTPInstanceEntry[i]), types.YChild{"StpxMISTPInstanceEntry", stpxMISTPInstanceTable.StpxMISTPInstanceEntry[i]})
    }
    stpxMISTPInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    stpxMISTPInstanceTable.EntityData.YListKeys = []string {}

    return &(stpxMISTPInstanceTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable_StpxMISTPInstanceEntry
// A conceptual row containing the status of the MISTP 
// instance.
type CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable_StpxMISTPInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary integer within the range from 1 to
    // the value of stpxMISTPInstanceNumber that uniquely identifies an instance.
    // The type is interface{} with range: 1..256.
    StpxMISTPInstanceIndex interface{}

    // This object indicates whether the MISTP protocol is currently enabled on
    // the MISTP instance.  If this object is set to    'true'    - the MISTP
    // protocol will run on this instance.                   'false'   - the MISTP
    // protocol will stop running on this                 instance. The type is
    // bool.
    StpxMISTPInstanceEnable interface{}

    // A string of octets containing one bit per VLAN. The first octet corresponds
    // to VLANs with VlanIndex values of 0 through 7; the second octet to VLANs 8
    // through 15; etc.  The most significant bit of each octet corresponds to the
    // lowest value VlanIndex in that octet.  For each VLAN, if it is mapped to
    // this MISTP instance, then the bit corresponding to that VLAN is set to '1'.
    // The type is string with length: 0..128.
    StpxMISTPInstanceVlansMapped interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 1024 through 1031; the second octet to VLANs 1032
    // through 1039; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    StpxMISTPInstanceVlansMapped2k interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    StpxMISTPInstanceVlansMapped3k interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 3072 through 3079; the second octet to VLANs 3080
    // through 3087; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MISTP instance, then the bit corresponding to that VLAN is set to
    // '1'.  This object is only instantiated on devices with  support for
    // VlanIndex up to 4095. The type is string with length: 0..128.
    StpxMISTPInstanceVlansMapped4k interface{}
}

func (stpxMISTPInstanceEntry *CISCOSTPEXTENSIONSMIB_StpxMISTPInstanceTable_StpxMISTPInstanceEntry) GetEntityData() *types.CommonEntityData {
    stpxMISTPInstanceEntry.EntityData.YFilter = stpxMISTPInstanceEntry.YFilter
    stpxMISTPInstanceEntry.EntityData.YangName = "stpxMISTPInstanceEntry"
    stpxMISTPInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxMISTPInstanceEntry.EntityData.ParentYangName = "stpxMISTPInstanceTable"
    stpxMISTPInstanceEntry.EntityData.SegmentPath = "stpxMISTPInstanceEntry" + types.AddKeyToken(stpxMISTPInstanceEntry.StpxMISTPInstanceIndex, "stpxMISTPInstanceIndex")
    stpxMISTPInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMISTPInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMISTPInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMISTPInstanceEntry.EntityData.Children = types.NewOrderedMap()
    stpxMISTPInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceIndex", types.YLeaf{"StpxMISTPInstanceIndex", stpxMISTPInstanceEntry.StpxMISTPInstanceIndex})
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceEnable", types.YLeaf{"StpxMISTPInstanceEnable", stpxMISTPInstanceEntry.StpxMISTPInstanceEnable})
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceVlansMapped", types.YLeaf{"StpxMISTPInstanceVlansMapped", stpxMISTPInstanceEntry.StpxMISTPInstanceVlansMapped})
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceVlansMapped2k", types.YLeaf{"StpxMISTPInstanceVlansMapped2k", stpxMISTPInstanceEntry.StpxMISTPInstanceVlansMapped2k})
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceVlansMapped3k", types.YLeaf{"StpxMISTPInstanceVlansMapped3k", stpxMISTPInstanceEntry.StpxMISTPInstanceVlansMapped3k})
    stpxMISTPInstanceEntry.EntityData.Leafs.Append("stpxMISTPInstanceVlansMapped4k", types.YLeaf{"StpxMISTPInstanceVlansMapped4k", stpxMISTPInstanceEntry.StpxMISTPInstanceVlansMapped4k})

    stpxMISTPInstanceEntry.EntityData.YListKeys = []string {"StpxMISTPInstanceIndex"}

    return &(stpxMISTPInstanceEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable
// A table containing a list of the bridge ports for which
// Spanning Tree LoopGuard capability can be configured.
type CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree LoopGuard  capability can be
    // configured. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry.
    StpxLoopGuardConfigEntry []*CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry
}

func (stpxLoopGuardConfigTable *CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable) GetEntityData() *types.CommonEntityData {
    stpxLoopGuardConfigTable.EntityData.YFilter = stpxLoopGuardConfigTable.YFilter
    stpxLoopGuardConfigTable.EntityData.YangName = "stpxLoopGuardConfigTable"
    stpxLoopGuardConfigTable.EntityData.BundleName = "cisco_ios_xe"
    stpxLoopGuardConfigTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxLoopGuardConfigTable.EntityData.SegmentPath = "stpxLoopGuardConfigTable"
    stpxLoopGuardConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxLoopGuardConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxLoopGuardConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxLoopGuardConfigTable.EntityData.Children = types.NewOrderedMap()
    stpxLoopGuardConfigTable.EntityData.Children.Append("stpxLoopGuardConfigEntry", types.YChild{"StpxLoopGuardConfigEntry", nil})
    for i := range stpxLoopGuardConfigTable.StpxLoopGuardConfigEntry {
        stpxLoopGuardConfigTable.EntityData.Children.Append(types.GetSegmentPath(stpxLoopGuardConfigTable.StpxLoopGuardConfigEntry[i]), types.YChild{"StpxLoopGuardConfigEntry", stpxLoopGuardConfigTable.StpxLoopGuardConfigEntry[i]})
    }
    stpxLoopGuardConfigTable.EntityData.Leafs = types.NewOrderedMap()

    stpxLoopGuardConfigTable.EntityData.YListKeys = []string {}

    return &(stpxLoopGuardConfigTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry
// A bridge port for which Spanning Tree LoopGuard 
// capability can be configured.
type CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxLoopGuardConfigPortIndex interface{}

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
    StpxLoopGuardConfigEnabled interface{}

    // Indicates the mode of Loop Guard Feature on this  port. This configuration
    // will be applied to all  the Spanning Tree instances in which this port 
    // exists.  enable -- the Loop Guard feature is enabled on this           
    // port.   disable -- the Loop Guard feature is disabled on this           
    // port.    default -- whether the Loop Guard feature is enabled            or
    // not on this port depends on the object             value of
    // stpxLoopGuardGlobalDefaultMode. The type is StpxLoopGuardConfigMode.
    StpxLoopGuardConfigMode interface{}
}

func (stpxLoopGuardConfigEntry *CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry) GetEntityData() *types.CommonEntityData {
    stpxLoopGuardConfigEntry.EntityData.YFilter = stpxLoopGuardConfigEntry.YFilter
    stpxLoopGuardConfigEntry.EntityData.YangName = "stpxLoopGuardConfigEntry"
    stpxLoopGuardConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxLoopGuardConfigEntry.EntityData.ParentYangName = "stpxLoopGuardConfigTable"
    stpxLoopGuardConfigEntry.EntityData.SegmentPath = "stpxLoopGuardConfigEntry" + types.AddKeyToken(stpxLoopGuardConfigEntry.StpxLoopGuardConfigPortIndex, "stpxLoopGuardConfigPortIndex")
    stpxLoopGuardConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxLoopGuardConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxLoopGuardConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxLoopGuardConfigEntry.EntityData.Children = types.NewOrderedMap()
    stpxLoopGuardConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxLoopGuardConfigEntry.EntityData.Leafs.Append("stpxLoopGuardConfigPortIndex", types.YLeaf{"StpxLoopGuardConfigPortIndex", stpxLoopGuardConfigEntry.StpxLoopGuardConfigPortIndex})
    stpxLoopGuardConfigEntry.EntityData.Leafs.Append("stpxLoopGuardConfigEnabled", types.YLeaf{"StpxLoopGuardConfigEnabled", stpxLoopGuardConfigEntry.StpxLoopGuardConfigEnabled})
    stpxLoopGuardConfigEntry.EntityData.Leafs.Append("stpxLoopGuardConfigMode", types.YLeaf{"StpxLoopGuardConfigMode", stpxLoopGuardConfigEntry.StpxLoopGuardConfigMode})

    stpxLoopGuardConfigEntry.EntityData.YListKeys = []string {"StpxLoopGuardConfigPortIndex"}

    return &(stpxLoopGuardConfigEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode represents            value of stpxLoopGuardGlobalDefaultMode.
type CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode_enable CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode_disable CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode = "disable"

    CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode_default_ CISCOSTPEXTENSIONSMIB_StpxLoopGuardConfigTable_StpxLoopGuardConfigEntry_StpxLoopGuardConfigMode = "default"
)

// CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable
// A table containing a list of the bridge ports for which
// a particular Spanning Tree instance has been found
// to have a loop-inconsistency. The agent creates a new
// entry in this table whenever it detects a new
// loop-inconsistency, and deletes entries
// when/soon after the inconsistency is no longer present.
type CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which a Spanning Tree
    // loop-inconsistency is currently in effect. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable_StpxLoopInconsistencyEntry.
    StpxLoopInconsistencyEntry []*CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable_StpxLoopInconsistencyEntry
}

func (stpxLoopInconsistencyTable *CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable) GetEntityData() *types.CommonEntityData {
    stpxLoopInconsistencyTable.EntityData.YFilter = stpxLoopInconsistencyTable.YFilter
    stpxLoopInconsistencyTable.EntityData.YangName = "stpxLoopInconsistencyTable"
    stpxLoopInconsistencyTable.EntityData.BundleName = "cisco_ios_xe"
    stpxLoopInconsistencyTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxLoopInconsistencyTable.EntityData.SegmentPath = "stpxLoopInconsistencyTable"
    stpxLoopInconsistencyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxLoopInconsistencyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxLoopInconsistencyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxLoopInconsistencyTable.EntityData.Children = types.NewOrderedMap()
    stpxLoopInconsistencyTable.EntityData.Children.Append("stpxLoopInconsistencyEntry", types.YChild{"StpxLoopInconsistencyEntry", nil})
    for i := range stpxLoopInconsistencyTable.StpxLoopInconsistencyEntry {
        stpxLoopInconsistencyTable.EntityData.Children.Append(types.GetSegmentPath(stpxLoopInconsistencyTable.StpxLoopInconsistencyEntry[i]), types.YChild{"StpxLoopInconsistencyEntry", stpxLoopInconsistencyTable.StpxLoopInconsistencyEntry[i]})
    }
    stpxLoopInconsistencyTable.EntityData.Leafs = types.NewOrderedMap()

    stpxLoopInconsistencyTable.EntityData.YListKeys = []string {}

    return &(stpxLoopInconsistencyTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable_StpxLoopInconsistencyEntry
// A Spanning Tree instance on a particular port for
// which a Spanning Tree loop-inconsistency is currently
// in effect.
type CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable_StpxLoopInconsistencyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1) or
    // rapidPvstPlus(5). The type is interface{} with range: 0..65535.
    StpxLoopInconsistencyIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxLoopInconsistencyPortIndex interface{}

    // Indicates whether the port on a particular Spanning  Tree instance is
    // currently in loop-inconsistent  state or not. The type is bool.
    StpxLoopInconsistencyState interface{}
}

func (stpxLoopInconsistencyEntry *CISCOSTPEXTENSIONSMIB_StpxLoopInconsistencyTable_StpxLoopInconsistencyEntry) GetEntityData() *types.CommonEntityData {
    stpxLoopInconsistencyEntry.EntityData.YFilter = stpxLoopInconsistencyEntry.YFilter
    stpxLoopInconsistencyEntry.EntityData.YangName = "stpxLoopInconsistencyEntry"
    stpxLoopInconsistencyEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxLoopInconsistencyEntry.EntityData.ParentYangName = "stpxLoopInconsistencyTable"
    stpxLoopInconsistencyEntry.EntityData.SegmentPath = "stpxLoopInconsistencyEntry" + types.AddKeyToken(stpxLoopInconsistencyEntry.StpxLoopInconsistencyIndex, "stpxLoopInconsistencyIndex") + types.AddKeyToken(stpxLoopInconsistencyEntry.StpxLoopInconsistencyPortIndex, "stpxLoopInconsistencyPortIndex")
    stpxLoopInconsistencyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxLoopInconsistencyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxLoopInconsistencyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxLoopInconsistencyEntry.EntityData.Children = types.NewOrderedMap()
    stpxLoopInconsistencyEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxLoopInconsistencyEntry.EntityData.Leafs.Append("stpxLoopInconsistencyIndex", types.YLeaf{"StpxLoopInconsistencyIndex", stpxLoopInconsistencyEntry.StpxLoopInconsistencyIndex})
    stpxLoopInconsistencyEntry.EntityData.Leafs.Append("stpxLoopInconsistencyPortIndex", types.YLeaf{"StpxLoopInconsistencyPortIndex", stpxLoopInconsistencyEntry.StpxLoopInconsistencyPortIndex})
    stpxLoopInconsistencyEntry.EntityData.Leafs.Append("stpxLoopInconsistencyState", types.YLeaf{"StpxLoopInconsistencyState", stpxLoopInconsistencyEntry.StpxLoopInconsistencyState})

    stpxLoopInconsistencyEntry.EntityData.YListKeys = []string {"StpxLoopInconsistencyIndex", "StpxLoopInconsistencyPortIndex"}

    return &(stpxLoopInconsistencyEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable
// A table containing a list of the bridge ports for
// which Spanning Tree Port Fast Start can be
// configured.
type CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A bridge port for which Spanning Tree Port Fast Start can be configured.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry.
    StpxFastStartPortEntry []*CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry
}

func (stpxFastStartPortTable *CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable) GetEntityData() *types.CommonEntityData {
    stpxFastStartPortTable.EntityData.YFilter = stpxFastStartPortTable.YFilter
    stpxFastStartPortTable.EntityData.YangName = "stpxFastStartPortTable"
    stpxFastStartPortTable.EntityData.BundleName = "cisco_ios_xe"
    stpxFastStartPortTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxFastStartPortTable.EntityData.SegmentPath = "stpxFastStartPortTable"
    stpxFastStartPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxFastStartPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxFastStartPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxFastStartPortTable.EntityData.Children = types.NewOrderedMap()
    stpxFastStartPortTable.EntityData.Children.Append("stpxFastStartPortEntry", types.YChild{"StpxFastStartPortEntry", nil})
    for i := range stpxFastStartPortTable.StpxFastStartPortEntry {
        stpxFastStartPortTable.EntityData.Children.Append(types.GetSegmentPath(stpxFastStartPortTable.StpxFastStartPortEntry[i]), types.YChild{"StpxFastStartPortEntry", stpxFastStartPortTable.StpxFastStartPortEntry[i]})
    }
    stpxFastStartPortTable.EntityData.Leafs = types.NewOrderedMap()

    stpxFastStartPortTable.EntityData.YListKeys = []string {}

    return &(stpxFastStartPortTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry
// A bridge port for which Spanning Tree Port Fast
// Start can be configured.
type CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxFastStartPortIndex interface{}

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
    StpxFastStartPortEnable interface{}

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
    // StpxFastStartPortMode.
    StpxFastStartPortMode interface{}

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
    // BPDU from this port. The type is StpxFastStartPortBpduGuardMode.
    StpxFastStartPortBpduGuardMode interface{}

    // Indicates the mode of Bpdu Filter Feature on the port. The system will not
    // transmit BPDUs on a port  with Bpdu Filter feature enabled.  enable -- the
    // Bpdu Filter feature is enabled on this            port.   disable -- the
    // Bpdu Filter feature is disabled on this            port.  default --
    // whether the Bpdu Filter feature is enabled            or not on this port
    // depends on the object            value of stpxFastStartBpduFilterEnable. If
    // the value of stpxFastStartBpduFilterEnable            is true(1) and Fast
    // Start feature is also            enabled operationally on this port, then  
    // no BPDUs will be transmitted on this port. The type is
    // StpxFastStartPortBpduFilterMode.
    StpxFastStartPortBpduFilterMode interface{}
}

func (stpxFastStartPortEntry *CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry) GetEntityData() *types.CommonEntityData {
    stpxFastStartPortEntry.EntityData.YFilter = stpxFastStartPortEntry.YFilter
    stpxFastStartPortEntry.EntityData.YangName = "stpxFastStartPortEntry"
    stpxFastStartPortEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxFastStartPortEntry.EntityData.ParentYangName = "stpxFastStartPortTable"
    stpxFastStartPortEntry.EntityData.SegmentPath = "stpxFastStartPortEntry" + types.AddKeyToken(stpxFastStartPortEntry.StpxFastStartPortIndex, "stpxFastStartPortIndex")
    stpxFastStartPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxFastStartPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxFastStartPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxFastStartPortEntry.EntityData.Children = types.NewOrderedMap()
    stpxFastStartPortEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxFastStartPortEntry.EntityData.Leafs.Append("stpxFastStartPortIndex", types.YLeaf{"StpxFastStartPortIndex", stpxFastStartPortEntry.StpxFastStartPortIndex})
    stpxFastStartPortEntry.EntityData.Leafs.Append("stpxFastStartPortEnable", types.YLeaf{"StpxFastStartPortEnable", stpxFastStartPortEntry.StpxFastStartPortEnable})
    stpxFastStartPortEntry.EntityData.Leafs.Append("stpxFastStartPortMode", types.YLeaf{"StpxFastStartPortMode", stpxFastStartPortEntry.StpxFastStartPortMode})
    stpxFastStartPortEntry.EntityData.Leafs.Append("stpxFastStartPortBpduGuardMode", types.YLeaf{"StpxFastStartPortBpduGuardMode", stpxFastStartPortEntry.StpxFastStartPortBpduGuardMode})
    stpxFastStartPortEntry.EntityData.Leafs.Append("stpxFastStartPortBpduFilterMode", types.YLeaf{"StpxFastStartPortBpduFilterMode", stpxFastStartPortEntry.StpxFastStartPortBpduFilterMode})

    stpxFastStartPortEntry.EntityData.YListKeys = []string {"StpxFastStartPortIndex"}

    return &(stpxFastStartPortEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode represents            no BPDUs will be transmitted on this port.
type CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode_enable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode_disable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode = "disable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode_default_ CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduFilterMode = "default"
)

// CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode represents            the system receives a BPDU from this port.
type CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode_enable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode_disable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode = "disable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode_default_ CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortBpduGuardMode = "default"
)

// CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode represents            this port.
type CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode_enable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode = "enable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode_disable CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode = "disable"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode_enableForTrunk CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode = "enableForTrunk"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode_default_ CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode = "default"

    CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode_network CISCOSTPEXTENSIONSMIB_StpxFastStartPortTable_StpxFastStartPortEntry_StpxFastStartPortMode = "network"
)

// CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable
// A table containing a list of the bridge ports 
// for a particular Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port fast start oper mode  information on a bridge port for a
    // particular  Spanning Tree Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry.
    StpxFastStartOperModeEntry []*CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry
}

func (stpxFastStartOperModeTable *CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable) GetEntityData() *types.CommonEntityData {
    stpxFastStartOperModeTable.EntityData.YFilter = stpxFastStartOperModeTable.YFilter
    stpxFastStartOperModeTable.EntityData.YangName = "stpxFastStartOperModeTable"
    stpxFastStartOperModeTable.EntityData.BundleName = "cisco_ios_xe"
    stpxFastStartOperModeTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxFastStartOperModeTable.EntityData.SegmentPath = "stpxFastStartOperModeTable"
    stpxFastStartOperModeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxFastStartOperModeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxFastStartOperModeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxFastStartOperModeTable.EntityData.Children = types.NewOrderedMap()
    stpxFastStartOperModeTable.EntityData.Children.Append("stpxFastStartOperModeEntry", types.YChild{"StpxFastStartOperModeEntry", nil})
    for i := range stpxFastStartOperModeTable.StpxFastStartOperModeEntry {
        stpxFastStartOperModeTable.EntityData.Children.Append(types.GetSegmentPath(stpxFastStartOperModeTable.StpxFastStartOperModeEntry[i]), types.YChild{"StpxFastStartOperModeEntry", stpxFastStartOperModeTable.StpxFastStartOperModeEntry[i]})
    }
    stpxFastStartOperModeTable.EntityData.Leafs = types.NewOrderedMap()

    stpxFastStartOperModeTable.EntityData.YListKeys = []string {}

    return &(stpxFastStartOperModeTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry
// An entry with port fast start oper mode 
// information on a bridge port for a particular 
// Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType is pvstPlus(1). The
    // type is interface{} with range: 0..65535.
    StpxFastStartOperModeInstIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxFastStartOperModePortIndex interface{}

    // Indicates the fast start operational status of the  port on a particular
    // Spanning Tree Instance. The type is StpxFastStartOperMode.
    StpxFastStartOperMode interface{}
}

func (stpxFastStartOperModeEntry *CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry) GetEntityData() *types.CommonEntityData {
    stpxFastStartOperModeEntry.EntityData.YFilter = stpxFastStartOperModeEntry.YFilter
    stpxFastStartOperModeEntry.EntityData.YangName = "stpxFastStartOperModeEntry"
    stpxFastStartOperModeEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxFastStartOperModeEntry.EntityData.ParentYangName = "stpxFastStartOperModeTable"
    stpxFastStartOperModeEntry.EntityData.SegmentPath = "stpxFastStartOperModeEntry" + types.AddKeyToken(stpxFastStartOperModeEntry.StpxFastStartOperModeInstIndex, "stpxFastStartOperModeInstIndex") + types.AddKeyToken(stpxFastStartOperModeEntry.StpxFastStartOperModePortIndex, "stpxFastStartOperModePortIndex")
    stpxFastStartOperModeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxFastStartOperModeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxFastStartOperModeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxFastStartOperModeEntry.EntityData.Children = types.NewOrderedMap()
    stpxFastStartOperModeEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxFastStartOperModeEntry.EntityData.Leafs.Append("stpxFastStartOperModeInstIndex", types.YLeaf{"StpxFastStartOperModeInstIndex", stpxFastStartOperModeEntry.StpxFastStartOperModeInstIndex})
    stpxFastStartOperModeEntry.EntityData.Leafs.Append("stpxFastStartOperModePortIndex", types.YLeaf{"StpxFastStartOperModePortIndex", stpxFastStartOperModeEntry.StpxFastStartOperModePortIndex})
    stpxFastStartOperModeEntry.EntityData.Leafs.Append("stpxFastStartOperMode", types.YLeaf{"StpxFastStartOperMode", stpxFastStartOperModeEntry.StpxFastStartOperMode})

    stpxFastStartOperModeEntry.EntityData.YListKeys = []string {"StpxFastStartOperModeInstIndex", "StpxFastStartOperModePortIndex"}

    return &(stpxFastStartOperModeEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode represents port on a particular Spanning Tree Instance.
type CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode_enabled CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode = "enabled"

    CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode_disabled CISCOSTPEXTENSIONSMIB_StpxFastStartOperModeTable_StpxFastStartOperModeEntry_StpxFastStartOperMode = "disabled"
)

// CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable
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
type CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Spanning Tree instance on a particular port for which BPDU skewing has
    // been detected. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable_StpxBpduSkewingEntry.
    StpxBpduSkewingEntry []*CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable_StpxBpduSkewingEntry
}

func (stpxBpduSkewingTable *CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable) GetEntityData() *types.CommonEntityData {
    stpxBpduSkewingTable.EntityData.YFilter = stpxBpduSkewingTable.YFilter
    stpxBpduSkewingTable.EntityData.YangName = "stpxBpduSkewingTable"
    stpxBpduSkewingTable.EntityData.BundleName = "cisco_ios_xe"
    stpxBpduSkewingTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxBpduSkewingTable.EntityData.SegmentPath = "stpxBpduSkewingTable"
    stpxBpduSkewingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxBpduSkewingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxBpduSkewingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxBpduSkewingTable.EntityData.Children = types.NewOrderedMap()
    stpxBpduSkewingTable.EntityData.Children.Append("stpxBpduSkewingEntry", types.YChild{"StpxBpduSkewingEntry", nil})
    for i := range stpxBpduSkewingTable.StpxBpduSkewingEntry {
        stpxBpduSkewingTable.EntityData.Children.Append(types.GetSegmentPath(stpxBpduSkewingTable.StpxBpduSkewingEntry[i]), types.YChild{"StpxBpduSkewingEntry", stpxBpduSkewingTable.StpxBpduSkewingEntry[i]})
    }
    stpxBpduSkewingTable.EntityData.Leafs = types.NewOrderedMap()

    stpxBpduSkewingTable.EntityData.YListKeys = []string {}

    return &(stpxBpduSkewingTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable_StpxBpduSkewingEntry
// A Spanning Tree instance on a particular port for
// which BPDU skewing has been detected.
type CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable_StpxBpduSkewingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, such as the VLAN id
    // of the VLAN if the object value of stpxSpanningTreeType  is pvstPlus(1).
    // The type is interface{} with range: 0..65535.
    StpxBpduSkewingInstanceIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxBpduSkewingPortIndex interface{}

    // Indicates the skew duration in milliseconds of the last BPDU skewing
    // detected. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    StpxBpduSkewingLastSkewDuration interface{}

    // Indicates the skew duration in milliseconds of the worst BPDU skewing
    // detected. The type is interface{} with range: 0..4294967295. Units are
    // milliseconds.
    StpxBpduSkewingWorstSkewDuration interface{}

    // Indicates the value of sysUpTime when the worst BPDU skewing was detected.
    // The type is interface{} with range: 0..4294967295.
    StpxBpduSkewingWorstSkewTime interface{}
}

func (stpxBpduSkewingEntry *CISCOSTPEXTENSIONSMIB_StpxBpduSkewingTable_StpxBpduSkewingEntry) GetEntityData() *types.CommonEntityData {
    stpxBpduSkewingEntry.EntityData.YFilter = stpxBpduSkewingEntry.YFilter
    stpxBpduSkewingEntry.EntityData.YangName = "stpxBpduSkewingEntry"
    stpxBpduSkewingEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxBpduSkewingEntry.EntityData.ParentYangName = "stpxBpduSkewingTable"
    stpxBpduSkewingEntry.EntityData.SegmentPath = "stpxBpduSkewingEntry" + types.AddKeyToken(stpxBpduSkewingEntry.StpxBpduSkewingInstanceIndex, "stpxBpduSkewingInstanceIndex") + types.AddKeyToken(stpxBpduSkewingEntry.StpxBpduSkewingPortIndex, "stpxBpduSkewingPortIndex")
    stpxBpduSkewingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxBpduSkewingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxBpduSkewingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxBpduSkewingEntry.EntityData.Children = types.NewOrderedMap()
    stpxBpduSkewingEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxBpduSkewingEntry.EntityData.Leafs.Append("stpxBpduSkewingInstanceIndex", types.YLeaf{"StpxBpduSkewingInstanceIndex", stpxBpduSkewingEntry.StpxBpduSkewingInstanceIndex})
    stpxBpduSkewingEntry.EntityData.Leafs.Append("stpxBpduSkewingPortIndex", types.YLeaf{"StpxBpduSkewingPortIndex", stpxBpduSkewingEntry.StpxBpduSkewingPortIndex})
    stpxBpduSkewingEntry.EntityData.Leafs.Append("stpxBpduSkewingLastSkewDuration", types.YLeaf{"StpxBpduSkewingLastSkewDuration", stpxBpduSkewingEntry.StpxBpduSkewingLastSkewDuration})
    stpxBpduSkewingEntry.EntityData.Leafs.Append("stpxBpduSkewingWorstSkewDuration", types.YLeaf{"StpxBpduSkewingWorstSkewDuration", stpxBpduSkewingEntry.StpxBpduSkewingWorstSkewDuration})
    stpxBpduSkewingEntry.EntityData.Leafs.Append("stpxBpduSkewingWorstSkewTime", types.YLeaf{"StpxBpduSkewingWorstSkewTime", stpxBpduSkewingEntry.StpxBpduSkewingWorstSkewTime})

    stpxBpduSkewingEntry.EntityData.YListKeys = []string {"StpxBpduSkewingInstanceIndex", "StpxBpduSkewingPortIndex"}

    return &(stpxBpduSkewingEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable
// This table contains MST instance information with
// one entry for an MST instance within the range of 
// 0 to the object value of stpxMSTMaxInstanceNumber. 
// 
// This table is deprecated and replaced by 
// stpxSMSTInstanceTable.
type CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information. The type is
    // slice of CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable_StpxMSTInstanceEntry.
    StpxMSTInstanceEntry []*CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable_StpxMSTInstanceEntry
}

func (stpxMSTInstanceTable *CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable) GetEntityData() *types.CommonEntityData {
    stpxMSTInstanceTable.EntityData.YFilter = stpxMSTInstanceTable.YFilter
    stpxMSTInstanceTable.EntityData.YangName = "stpxMSTInstanceTable"
    stpxMSTInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTInstanceTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMSTInstanceTable.EntityData.SegmentPath = "stpxMSTInstanceTable"
    stpxMSTInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTInstanceTable.EntityData.Children = types.NewOrderedMap()
    stpxMSTInstanceTable.EntityData.Children.Append("stpxMSTInstanceEntry", types.YChild{"StpxMSTInstanceEntry", nil})
    for i := range stpxMSTInstanceTable.StpxMSTInstanceEntry {
        stpxMSTInstanceTable.EntityData.Children.Append(types.GetSegmentPath(stpxMSTInstanceTable.StpxMSTInstanceEntry[i]), types.YChild{"StpxMSTInstanceEntry", stpxMSTInstanceTable.StpxMSTInstanceEntry[i]})
    }
    stpxMSTInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    stpxMSTInstanceTable.EntityData.YListKeys = []string {}

    return &(stpxMSTInstanceTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable_StpxMSTInstanceEntry
// A conceptual row containing the MST instance 
// information.
type CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable_StpxMSTInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer that uniquely identifies an MST
    // instance  within the range of 0 to the object value of
    // stpxMSTMaxInstanceNumber.  This object is deprecated and replaced by 
    // stpxSMSTInstanceIndex. The type is interface{} with range: 0..256.
    StpxMSTInstanceIndex interface{}

    // A string of octets containing one bit per VLAN. The first octet corresponds
    // to VLANs with VlanIndex values of 0 through 7; the second octet to VLANs 8
    // through 15; etc.  The most significant bit of each octet corresponds to the
    // lowest value VlanIndex in that octet.  For each VLAN, if it is mapped to
    // this MST instance,  then the bit corresponding to that VLAN is set to '1'. 
    // This object is deprecated and replaced by  stpxSMSTInstanceVlansMapped1k2k.
    // The type is string with length: 0..128.
    StpxMSTInstanceVlansMapped interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 1024 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 1024 through 1031; the second octet to VLANs 1032
    // through 1039; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by 
    // stpxSMSTInstanceVlansMapped1k2k. The type is string with length: 0..128.
    StpxMSTInstanceVlansMapped2k interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 3071. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by 
    // stpxSMSTInstanceVlansMapped3k4k. The type is string with length: 0..128.
    StpxMSTInstanceVlansMapped3k interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 3072 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 3072 through 3079; the second octet to VLANs 3080
    // through 3087; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance,  then the bit corresponding to that VLAN is set to
    // '1'.  This object is deprecated and replaced by
    // stpxSMSTInstanceVlansMapped3k4k. The type is string with length: 0..128.
    StpxMSTInstanceVlansMapped4k interface{}

    // The remaining hop count for this MST instance.  This object will take on
    // value of 40 if the object value of stpxSMSTInstanceRemainingHopCount is
    // greater than 40.  This object is only instantiated when the object value of
    // stpxSpanningTreeType is mst(4).  This object is deprecated and replaced by 
    // stpxSMSTInstanceRemainingHopCount. The type is interface{} with range:
    // 0..40.
    StpxMSTInstanceRemainingHopCount interface{}
}

func (stpxMSTInstanceEntry *CISCOSTPEXTENSIONSMIB_StpxMSTInstanceTable_StpxMSTInstanceEntry) GetEntityData() *types.CommonEntityData {
    stpxMSTInstanceEntry.EntityData.YFilter = stpxMSTInstanceEntry.YFilter
    stpxMSTInstanceEntry.EntityData.YangName = "stpxMSTInstanceEntry"
    stpxMSTInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTInstanceEntry.EntityData.ParentYangName = "stpxMSTInstanceTable"
    stpxMSTInstanceEntry.EntityData.SegmentPath = "stpxMSTInstanceEntry" + types.AddKeyToken(stpxMSTInstanceEntry.StpxMSTInstanceIndex, "stpxMSTInstanceIndex")
    stpxMSTInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTInstanceEntry.EntityData.Children = types.NewOrderedMap()
    stpxMSTInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceIndex", types.YLeaf{"StpxMSTInstanceIndex", stpxMSTInstanceEntry.StpxMSTInstanceIndex})
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceVlansMapped", types.YLeaf{"StpxMSTInstanceVlansMapped", stpxMSTInstanceEntry.StpxMSTInstanceVlansMapped})
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceVlansMapped2k", types.YLeaf{"StpxMSTInstanceVlansMapped2k", stpxMSTInstanceEntry.StpxMSTInstanceVlansMapped2k})
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceVlansMapped3k", types.YLeaf{"StpxMSTInstanceVlansMapped3k", stpxMSTInstanceEntry.StpxMSTInstanceVlansMapped3k})
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceVlansMapped4k", types.YLeaf{"StpxMSTInstanceVlansMapped4k", stpxMSTInstanceEntry.StpxMSTInstanceVlansMapped4k})
    stpxMSTInstanceEntry.EntityData.Leafs.Append("stpxMSTInstanceRemainingHopCount", types.YLeaf{"StpxMSTInstanceRemainingHopCount", stpxMSTInstanceEntry.StpxMSTInstanceRemainingHopCount})

    stpxMSTInstanceEntry.EntityData.YListKeys = []string {"StpxMSTInstanceIndex"}

    return &(stpxMSTInstanceEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable
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
type CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer.
    // The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable_StpxMSTInstanceEditEntry.
    StpxMSTInstanceEditEntry []*CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable_StpxMSTInstanceEditEntry
}

func (stpxMSTInstanceEditTable *CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable) GetEntityData() *types.CommonEntityData {
    stpxMSTInstanceEditTable.EntityData.YFilter = stpxMSTInstanceEditTable.YFilter
    stpxMSTInstanceEditTable.EntityData.YangName = "stpxMSTInstanceEditTable"
    stpxMSTInstanceEditTable.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTInstanceEditTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMSTInstanceEditTable.EntityData.SegmentPath = "stpxMSTInstanceEditTable"
    stpxMSTInstanceEditTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTInstanceEditTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTInstanceEditTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTInstanceEditTable.EntityData.Children = types.NewOrderedMap()
    stpxMSTInstanceEditTable.EntityData.Children.Append("stpxMSTInstanceEditEntry", types.YChild{"StpxMSTInstanceEditEntry", nil})
    for i := range stpxMSTInstanceEditTable.StpxMSTInstanceEditEntry {
        stpxMSTInstanceEditTable.EntityData.Children.Append(types.GetSegmentPath(stpxMSTInstanceEditTable.StpxMSTInstanceEditEntry[i]), types.YChild{"StpxMSTInstanceEditEntry", stpxMSTInstanceEditTable.StpxMSTInstanceEditEntry[i]})
    }
    stpxMSTInstanceEditTable.EntityData.Leafs = types.NewOrderedMap()

    stpxMSTInstanceEditTable.EntityData.YListKeys = []string {}

    return &(stpxMSTInstanceEditTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable_StpxMSTInstanceEditEntry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
type CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable_StpxMSTInstanceEditEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer that uniquely identifies an MST
    // instance  from 0 to the object value of stpxMSTMaxInstanceNumber.  The
    // instances of this table entry with  stpxMSTInstanceEditIndex of zero can
    // not be  modified.  This object is deprecated and replaced by 
    // stpxSMSTInstanceEditIndex. The type is interface{} with range: 0..256.
    StpxMSTInstanceEditIndex interface{}

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
    StpxMSTInstanceEditVlansMap interface{}

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
    StpxMSTInstanceEditVlansMap2k interface{}

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
    StpxMSTInstanceEditVlansMap3k interface{}

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
    StpxMSTInstanceEditVlansMap4k interface{}
}

func (stpxMSTInstanceEditEntry *CISCOSTPEXTENSIONSMIB_StpxMSTInstanceEditTable_StpxMSTInstanceEditEntry) GetEntityData() *types.CommonEntityData {
    stpxMSTInstanceEditEntry.EntityData.YFilter = stpxMSTInstanceEditEntry.YFilter
    stpxMSTInstanceEditEntry.EntityData.YangName = "stpxMSTInstanceEditEntry"
    stpxMSTInstanceEditEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTInstanceEditEntry.EntityData.ParentYangName = "stpxMSTInstanceEditTable"
    stpxMSTInstanceEditEntry.EntityData.SegmentPath = "stpxMSTInstanceEditEntry" + types.AddKeyToken(stpxMSTInstanceEditEntry.StpxMSTInstanceEditIndex, "stpxMSTInstanceEditIndex")
    stpxMSTInstanceEditEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTInstanceEditEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTInstanceEditEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTInstanceEditEntry.EntityData.Children = types.NewOrderedMap()
    stpxMSTInstanceEditEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxMSTInstanceEditEntry.EntityData.Leafs.Append("stpxMSTInstanceEditIndex", types.YLeaf{"StpxMSTInstanceEditIndex", stpxMSTInstanceEditEntry.StpxMSTInstanceEditIndex})
    stpxMSTInstanceEditEntry.EntityData.Leafs.Append("stpxMSTInstanceEditVlansMap", types.YLeaf{"StpxMSTInstanceEditVlansMap", stpxMSTInstanceEditEntry.StpxMSTInstanceEditVlansMap})
    stpxMSTInstanceEditEntry.EntityData.Leafs.Append("stpxMSTInstanceEditVlansMap2k", types.YLeaf{"StpxMSTInstanceEditVlansMap2k", stpxMSTInstanceEditEntry.StpxMSTInstanceEditVlansMap2k})
    stpxMSTInstanceEditEntry.EntityData.Leafs.Append("stpxMSTInstanceEditVlansMap3k", types.YLeaf{"StpxMSTInstanceEditVlansMap3k", stpxMSTInstanceEditEntry.StpxMSTInstanceEditVlansMap3k})
    stpxMSTInstanceEditEntry.EntityData.Leafs.Append("stpxMSTInstanceEditVlansMap4k", types.YLeaf{"StpxMSTInstanceEditVlansMap4k", stpxMSTInstanceEditEntry.StpxMSTInstanceEditVlansMap4k})

    stpxMSTInstanceEditEntry.EntityData.YListKeys = []string {"StpxMSTInstanceEditIndex"}

    return &(stpxMSTInstanceEditEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTPortTable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the MST Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry.
    StpxMSTPortEntry []*CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry
}

func (stpxMSTPortTable *CISCOSTPEXTENSIONSMIB_StpxMSTPortTable) GetEntityData() *types.CommonEntityData {
    stpxMSTPortTable.EntityData.YFilter = stpxMSTPortTable.YFilter
    stpxMSTPortTable.EntityData.YangName = "stpxMSTPortTable"
    stpxMSTPortTable.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTPortTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMSTPortTable.EntityData.SegmentPath = "stpxMSTPortTable"
    stpxMSTPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTPortTable.EntityData.Children = types.NewOrderedMap()
    stpxMSTPortTable.EntityData.Children.Append("stpxMSTPortEntry", types.YChild{"StpxMSTPortEntry", nil})
    for i := range stpxMSTPortTable.StpxMSTPortEntry {
        stpxMSTPortTable.EntityData.Children.Append(types.GetSegmentPath(stpxMSTPortTable.StpxMSTPortEntry[i]), types.YChild{"StpxMSTPortEntry", stpxMSTPortTable.StpxMSTPortEntry[i]})
    }
    stpxMSTPortTable.EntityData.Leafs = types.NewOrderedMap()

    stpxMSTPortTable.EntityData.YListKeys = []string {}

    return &(stpxMSTPortTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry
// An entry with port information for the MST Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxMSTPortIndex interface{}

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
    // stpxRSTPPortAdminLinkType. The type is StpxMSTPortAdminLinkType.
    StpxMSTPortAdminLinkType interface{}

    // Indicates the operational link type of a bridge port for the MST protocol. 
    // pointToPoint -- the port is operationally connected to         a
    // point-to-point link.  shared -- the port is operationally connected to     
    // a shared medium.  other -- none of the above.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4). 
    // stpxMSTPortOperLinkType  is deprecated and replaced with
    // stpxRSTPPortOperLinkType. The type is StpxMSTPortOperLinkType.
    StpxMSTPortOperLinkType interface{}

    // The protocol migration control on this port. When the  object value of 
    // stpxSpanningTreeType is mst(4) or  rapidPvstPlus(5), setting true(1) to
    // this object forces  the device to try using version 2 BPDUs on this port. 
    // When the object value of stpxSpanningTreeType is neither  mst(4) nor
    // rapidPvstPlus(5), setting true(1) to this  object has no effect. Setting
    // false(2) to this object has  no effect. This object always returns false(2)
    // when read. stpxMSTPortProtocolMigration is deprecated and  replaced with
    // stpxRSTPPortProtocolMigration. The type is bool.
    StpxMSTPortProtocolMigration interface{}

    // Indicates the operational status of the port for the  MST protocol.   edge
    // -- this port is an edge port for the MST region.  boundary -- this port is
    // a boundary port for the          MST region.  pvst --  this port is
    // connected to a PVST/PVST+ bridge.     stp -- this port is connected to a
    // Single Spanning         Tree bridge.   This object is only instantiated
    // when the object value of stpxSpanningTreeType is mst(4).  This object is
    // deprecated and replaced by  stpxSMSTPortStatus. The type is
    // map[string]bool.
    StpxMSTPortStatus interface{}
}

func (stpxMSTPortEntry *CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry) GetEntityData() *types.CommonEntityData {
    stpxMSTPortEntry.EntityData.YFilter = stpxMSTPortEntry.YFilter
    stpxMSTPortEntry.EntityData.YangName = "stpxMSTPortEntry"
    stpxMSTPortEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTPortEntry.EntityData.ParentYangName = "stpxMSTPortTable"
    stpxMSTPortEntry.EntityData.SegmentPath = "stpxMSTPortEntry" + types.AddKeyToken(stpxMSTPortEntry.StpxMSTPortIndex, "stpxMSTPortIndex")
    stpxMSTPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTPortEntry.EntityData.Children = types.NewOrderedMap()
    stpxMSTPortEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxMSTPortEntry.EntityData.Leafs.Append("stpxMSTPortIndex", types.YLeaf{"StpxMSTPortIndex", stpxMSTPortEntry.StpxMSTPortIndex})
    stpxMSTPortEntry.EntityData.Leafs.Append("stpxMSTPortAdminLinkType", types.YLeaf{"StpxMSTPortAdminLinkType", stpxMSTPortEntry.StpxMSTPortAdminLinkType})
    stpxMSTPortEntry.EntityData.Leafs.Append("stpxMSTPortOperLinkType", types.YLeaf{"StpxMSTPortOperLinkType", stpxMSTPortEntry.StpxMSTPortOperLinkType})
    stpxMSTPortEntry.EntityData.Leafs.Append("stpxMSTPortProtocolMigration", types.YLeaf{"StpxMSTPortProtocolMigration", stpxMSTPortEntry.StpxMSTPortProtocolMigration})
    stpxMSTPortEntry.EntityData.Leafs.Append("stpxMSTPortStatus", types.YLeaf{"StpxMSTPortStatus", stpxMSTPortEntry.StpxMSTPortStatus})

    stpxMSTPortEntry.EntityData.YListKeys = []string {"StpxMSTPortIndex"}

    return &(stpxMSTPortEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType represents with stpxRSTPPortAdminLinkType.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType string

const (
    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType_pointToPoint CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType_shared CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType = "shared"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType_auto CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortAdminLinkType = "auto"
)

// CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType represents is deprecated and replaced with stpxRSTPPortOperLinkType.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType string

const (
    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType_pointToPoint CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType_shared CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType = "shared"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType_other CISCOSTPEXTENSIONSMIB_StpxMSTPortTable_StpxMSTPortEntry_StpxMSTPortOperLinkType = "other"
)

// CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable
// A table containing a list of the bridge ports for a 
// particular MST instance.  This table is only instantiated 
// when the stpxSpanningTreeType is mst(4). 
// 
// This table is deprecated and replaced with 
// stpxRSTPPortRoleTable.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the port role information for the MST protocol on a
    // port for a particular MST instance existing on the system. The type is
    // slice of CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry.
    StpxMSTPortRoleEntry []*CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry
}

func (stpxMSTPortRoleTable *CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable) GetEntityData() *types.CommonEntityData {
    stpxMSTPortRoleTable.EntityData.YFilter = stpxMSTPortRoleTable.YFilter
    stpxMSTPortRoleTable.EntityData.YangName = "stpxMSTPortRoleTable"
    stpxMSTPortRoleTable.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTPortRoleTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxMSTPortRoleTable.EntityData.SegmentPath = "stpxMSTPortRoleTable"
    stpxMSTPortRoleTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTPortRoleTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTPortRoleTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTPortRoleTable.EntityData.Children = types.NewOrderedMap()
    stpxMSTPortRoleTable.EntityData.Children.Append("stpxMSTPortRoleEntry", types.YChild{"StpxMSTPortRoleEntry", nil})
    for i := range stpxMSTPortRoleTable.StpxMSTPortRoleEntry {
        stpxMSTPortRoleTable.EntityData.Children.Append(types.GetSegmentPath(stpxMSTPortRoleTable.StpxMSTPortRoleEntry[i]), types.YChild{"StpxMSTPortRoleEntry", stpxMSTPortRoleTable.StpxMSTPortRoleEntry[i]})
    }
    stpxMSTPortRoleTable.EntityData.Leafs = types.NewOrderedMap()

    stpxMSTPortRoleTable.EntityData.YListKeys = []string {}

    return &(stpxMSTPortRoleTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry
// An entry containing the port role information for the MST
// protocol on a port for a particular MST instance existing
// on the system.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance id within the range of 0 to
    // stpxMSTMaxInstanceNumber. The type is interface{} with range: 0..256.
    StpxMSTPortRoleInstanceIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxMSTPortRolePortIndex interface{}

    // Indicates the port role on a particular MST instance for the MST protocol. 
    // disabled --  this port has no role on this MST instance.   root -- this
    // port has the role of root port on this MST             instance.  
    // designated -- this port has the role of designated              port on
    // this MST instance.  alternate -- this port has the role of alternate port  
    // on this MST instance.  backUp -- this port has the role of backup port on
    // this               MST instance.  boundary -- this port has the role of
    // boundary port on              this MST instance.  master -- this port has
    // the role of master port on           this MST instance. The type is
    // StpxMSTPortRoleValue.
    StpxMSTPortRoleValue interface{}
}

func (stpxMSTPortRoleEntry *CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry) GetEntityData() *types.CommonEntityData {
    stpxMSTPortRoleEntry.EntityData.YFilter = stpxMSTPortRoleEntry.YFilter
    stpxMSTPortRoleEntry.EntityData.YangName = "stpxMSTPortRoleEntry"
    stpxMSTPortRoleEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxMSTPortRoleEntry.EntityData.ParentYangName = "stpxMSTPortRoleTable"
    stpxMSTPortRoleEntry.EntityData.SegmentPath = "stpxMSTPortRoleEntry" + types.AddKeyToken(stpxMSTPortRoleEntry.StpxMSTPortRoleInstanceIndex, "stpxMSTPortRoleInstanceIndex") + types.AddKeyToken(stpxMSTPortRoleEntry.StpxMSTPortRolePortIndex, "stpxMSTPortRolePortIndex")
    stpxMSTPortRoleEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxMSTPortRoleEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxMSTPortRoleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxMSTPortRoleEntry.EntityData.Children = types.NewOrderedMap()
    stpxMSTPortRoleEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxMSTPortRoleEntry.EntityData.Leafs.Append("stpxMSTPortRoleInstanceIndex", types.YLeaf{"StpxMSTPortRoleInstanceIndex", stpxMSTPortRoleEntry.StpxMSTPortRoleInstanceIndex})
    stpxMSTPortRoleEntry.EntityData.Leafs.Append("stpxMSTPortRolePortIndex", types.YLeaf{"StpxMSTPortRolePortIndex", stpxMSTPortRoleEntry.StpxMSTPortRolePortIndex})
    stpxMSTPortRoleEntry.EntityData.Leafs.Append("stpxMSTPortRoleValue", types.YLeaf{"StpxMSTPortRoleValue", stpxMSTPortRoleEntry.StpxMSTPortRoleValue})

    stpxMSTPortRoleEntry.EntityData.YListKeys = []string {"StpxMSTPortRoleInstanceIndex", "StpxMSTPortRolePortIndex"}

    return &(stpxMSTPortRoleEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue represents           this MST instance.
type CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue string

const (
    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_disabled CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "disabled"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_root CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "root"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_designated CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "designated"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_alternate CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "alternate"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_backUp CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "backUp"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_boundary CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "boundary"

    CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue_master CISCOSTPEXTENSIONSMIB_StpxMSTPortRoleTable_StpxMSTPortRoleEntry_StpxMSTPortRoleValue = "master"
)

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable
// A table containing port information for the RSTP 
// Protocol on all the bridge ports existing in the 
// system.
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the RSTP Protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry.
    StpxRSTPPortEntry []*CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry
}

func (stpxRSTPPortTable *CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable) GetEntityData() *types.CommonEntityData {
    stpxRSTPPortTable.EntityData.YFilter = stpxRSTPPortTable.YFilter
    stpxRSTPPortTable.EntityData.YangName = "stpxRSTPPortTable"
    stpxRSTPPortTable.EntityData.BundleName = "cisco_ios_xe"
    stpxRSTPPortTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRSTPPortTable.EntityData.SegmentPath = "stpxRSTPPortTable"
    stpxRSTPPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRSTPPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRSTPPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRSTPPortTable.EntityData.Children = types.NewOrderedMap()
    stpxRSTPPortTable.EntityData.Children.Append("stpxRSTPPortEntry", types.YChild{"StpxRSTPPortEntry", nil})
    for i := range stpxRSTPPortTable.StpxRSTPPortEntry {
        stpxRSTPPortTable.EntityData.Children.Append(types.GetSegmentPath(stpxRSTPPortTable.StpxRSTPPortEntry[i]), types.YChild{"StpxRSTPPortEntry", stpxRSTPPortTable.StpxRSTPPortEntry[i]})
    }
    stpxRSTPPortTable.EntityData.Leafs = types.NewOrderedMap()

    stpxRSTPPortTable.EntityData.YListKeys = []string {}

    return &(stpxRSTPPortTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry
// An entry with port information for the RSTP Protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxRSTPPortIndex interface{}

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
    // The type is StpxRSTPPortAdminLinkType.
    StpxRSTPPortAdminLinkType interface{}

    // Indicates the operational link type of a bridge port for the RSTP protocol.
    // pointToPoint -- the port is operationally connected to         a
    // point-to-point link.  shared -- the port is operationally connected to     
    // a shared medium.  other -- none of the above.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4) or
    // rapidPvstPlus(5). The type is StpxRSTPPortOperLinkType.
    StpxRSTPPortOperLinkType interface{}

    // The protocol migration control on this port. When the  object value of 
    // stpxSpanningTreeType is mst(4) or  rapidPvstPlus(5), setting true(1) to
    // this object forces  the device to try using version 2 BPDUs on this port. 
    // When the object value of stpxSpanningTreeType is neither  mst(4) nor
    // rapidPvstPlus(5), setting true(1) to  this object has no effect. Setting
    // false(2) to this  object has no effect. This object always returns 
    // false(2) when read. The type is bool.
    StpxRSTPPortProtocolMigration interface{}
}

func (stpxRSTPPortEntry *CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry) GetEntityData() *types.CommonEntityData {
    stpxRSTPPortEntry.EntityData.YFilter = stpxRSTPPortEntry.YFilter
    stpxRSTPPortEntry.EntityData.YangName = "stpxRSTPPortEntry"
    stpxRSTPPortEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxRSTPPortEntry.EntityData.ParentYangName = "stpxRSTPPortTable"
    stpxRSTPPortEntry.EntityData.SegmentPath = "stpxRSTPPortEntry" + types.AddKeyToken(stpxRSTPPortEntry.StpxRSTPPortIndex, "stpxRSTPPortIndex")
    stpxRSTPPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRSTPPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRSTPPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRSTPPortEntry.EntityData.Children = types.NewOrderedMap()
    stpxRSTPPortEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxRSTPPortEntry.EntityData.Leafs.Append("stpxRSTPPortIndex", types.YLeaf{"StpxRSTPPortIndex", stpxRSTPPortEntry.StpxRSTPPortIndex})
    stpxRSTPPortEntry.EntityData.Leafs.Append("stpxRSTPPortAdminLinkType", types.YLeaf{"StpxRSTPPortAdminLinkType", stpxRSTPPortEntry.StpxRSTPPortAdminLinkType})
    stpxRSTPPortEntry.EntityData.Leafs.Append("stpxRSTPPortOperLinkType", types.YLeaf{"StpxRSTPPortOperLinkType", stpxRSTPPortEntry.StpxRSTPPortOperLinkType})
    stpxRSTPPortEntry.EntityData.Leafs.Append("stpxRSTPPortProtocolMigration", types.YLeaf{"StpxRSTPPortProtocolMigration", stpxRSTPPortEntry.StpxRSTPPortProtocolMigration})

    stpxRSTPPortEntry.EntityData.YListKeys = []string {"StpxRSTPPortIndex"}

    return &(stpxRSTPPortEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType represents stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType string

const (
    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType_pointToPoint CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType_shared CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType = "shared"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType_auto CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortAdminLinkType = "auto"
)

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType represents stpxSpanningTreeType is mst(4) or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType string

const (
    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType_pointToPoint CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType = "pointToPoint"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType_shared CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType = "shared"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType_other CISCOSTPEXTENSIONSMIB_StpxRSTPPortTable_StpxRSTPPortEntry_StpxRSTPPortOperLinkType = "other"
)

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable
// A table containing a list of the bridge ports for a 
// particular Spanning Tree instance.  This table is 
// only instantiated when the stpxSpanningTreeType is mst(4) 
// or rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the port role information for the RSTP protocol on a
    // port for a particular Spanning Tree instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry.
    StpxRSTPPortRoleEntry []*CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry
}

func (stpxRSTPPortRoleTable *CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable) GetEntityData() *types.CommonEntityData {
    stpxRSTPPortRoleTable.EntityData.YFilter = stpxRSTPPortRoleTable.YFilter
    stpxRSTPPortRoleTable.EntityData.YangName = "stpxRSTPPortRoleTable"
    stpxRSTPPortRoleTable.EntityData.BundleName = "cisco_ios_xe"
    stpxRSTPPortRoleTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRSTPPortRoleTable.EntityData.SegmentPath = "stpxRSTPPortRoleTable"
    stpxRSTPPortRoleTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRSTPPortRoleTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRSTPPortRoleTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRSTPPortRoleTable.EntityData.Children = types.NewOrderedMap()
    stpxRSTPPortRoleTable.EntityData.Children.Append("stpxRSTPPortRoleEntry", types.YChild{"StpxRSTPPortRoleEntry", nil})
    for i := range stpxRSTPPortRoleTable.StpxRSTPPortRoleEntry {
        stpxRSTPPortRoleTable.EntityData.Children.Append(types.GetSegmentPath(stpxRSTPPortRoleTable.StpxRSTPPortRoleEntry[i]), types.YChild{"StpxRSTPPortRoleEntry", stpxRSTPPortRoleTable.StpxRSTPPortRoleEntry[i]})
    }
    stpxRSTPPortRoleTable.EntityData.Leafs = types.NewOrderedMap()

    stpxRSTPPortRoleTable.EntityData.YListKeys = []string {}

    return &(stpxRSTPPortRoleTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry
// An entry containing the port role information for the RSTP
// protocol on a port for a particular Spanning Tree instance.
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Spanning Tree instance id, it can either be a 
    // VLAN number if the stpxSpanningTreeType is rapidPvstPlus(5)  or an MST
    // instance id if the stpxSpanningTreeType is mst(4). The type is interface{}
    // with range: 0..4095.
    StpxRSTPPortRoleInstanceIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxRSTPPortRolePortIndex interface{}

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
    // value of stpxSpanningTreeType is mst(4). The type is StpxRSTPPortRoleValue.
    StpxRSTPPortRoleValue interface{}
}

func (stpxRSTPPortRoleEntry *CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry) GetEntityData() *types.CommonEntityData {
    stpxRSTPPortRoleEntry.EntityData.YFilter = stpxRSTPPortRoleEntry.YFilter
    stpxRSTPPortRoleEntry.EntityData.YangName = "stpxRSTPPortRoleEntry"
    stpxRSTPPortRoleEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxRSTPPortRoleEntry.EntityData.ParentYangName = "stpxRSTPPortRoleTable"
    stpxRSTPPortRoleEntry.EntityData.SegmentPath = "stpxRSTPPortRoleEntry" + types.AddKeyToken(stpxRSTPPortRoleEntry.StpxRSTPPortRoleInstanceIndex, "stpxRSTPPortRoleInstanceIndex") + types.AddKeyToken(stpxRSTPPortRoleEntry.StpxRSTPPortRolePortIndex, "stpxRSTPPortRolePortIndex")
    stpxRSTPPortRoleEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRSTPPortRoleEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRSTPPortRoleEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRSTPPortRoleEntry.EntityData.Children = types.NewOrderedMap()
    stpxRSTPPortRoleEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxRSTPPortRoleEntry.EntityData.Leafs.Append("stpxRSTPPortRoleInstanceIndex", types.YLeaf{"StpxRSTPPortRoleInstanceIndex", stpxRSTPPortRoleEntry.StpxRSTPPortRoleInstanceIndex})
    stpxRSTPPortRoleEntry.EntityData.Leafs.Append("stpxRSTPPortRolePortIndex", types.YLeaf{"StpxRSTPPortRolePortIndex", stpxRSTPPortRoleEntry.StpxRSTPPortRolePortIndex})
    stpxRSTPPortRoleEntry.EntityData.Leafs.Append("stpxRSTPPortRoleValue", types.YLeaf{"StpxRSTPPortRoleValue", stpxRSTPPortRoleEntry.StpxRSTPPortRoleValue})

    stpxRSTPPortRoleEntry.EntityData.YListKeys = []string {"StpxRSTPPortRoleInstanceIndex", "StpxRSTPPortRolePortIndex"}

    return &(stpxRSTPPortRoleEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue represents only when the object value of stpxSpanningTreeType is mst(4).
type CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue string

const (
    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_disabled CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "disabled"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_root CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "root"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_designated CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "designated"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_alternate CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "alternate"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_backUp CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "backUp"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_boundary CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "boundary"

    CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue_master CISCOSTPEXTENSIONSMIB_StpxRSTPPortRoleTable_StpxRSTPPortRoleEntry_StpxRSTPPortRoleValue = "master"
)

// CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable
// A table containing a list of the bridge ports 
// for a particular Spanning Tree Instance.
// This table is only instantiated when the object value
// of stpxSpanningTreeType is rapidPvstPlus(5).
type CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port status information on a  bridge port for a particular
    // Spanning Tree  Instance. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable_StpxRPVSTPortEntry.
    StpxRPVSTPortEntry []*CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable_StpxRPVSTPortEntry
}

func (stpxRPVSTPortTable *CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable) GetEntityData() *types.CommonEntityData {
    stpxRPVSTPortTable.EntityData.YFilter = stpxRPVSTPortTable.YFilter
    stpxRPVSTPortTable.EntityData.YangName = "stpxRPVSTPortTable"
    stpxRPVSTPortTable.EntityData.BundleName = "cisco_ios_xe"
    stpxRPVSTPortTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxRPVSTPortTable.EntityData.SegmentPath = "stpxRPVSTPortTable"
    stpxRPVSTPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRPVSTPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRPVSTPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRPVSTPortTable.EntityData.Children = types.NewOrderedMap()
    stpxRPVSTPortTable.EntityData.Children.Append("stpxRPVSTPortEntry", types.YChild{"StpxRPVSTPortEntry", nil})
    for i := range stpxRPVSTPortTable.StpxRPVSTPortEntry {
        stpxRPVSTPortTable.EntityData.Children.Append(types.GetSegmentPath(stpxRPVSTPortTable.StpxRPVSTPortEntry[i]), types.YChild{"StpxRPVSTPortEntry", stpxRPVSTPortTable.StpxRPVSTPortEntry[i]})
    }
    stpxRPVSTPortTable.EntityData.Leafs = types.NewOrderedMap()

    stpxRPVSTPortTable.EntityData.YListKeys = []string {}

    return &(stpxRPVSTPortTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable_StpxRPVSTPortEntry
// An entry with port status information on a 
// bridge port for a particular Spanning Tree 
// Instance.
type CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable_StpxRPVSTPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VLAN id of the VLAN. The type is interface{}
    // with range: 0..4095.
    StpxRPVSTPortVlanIndex interface{}

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxRPVSTPortIndex interface{}

    // Indicates the operational status of the port for the  Rapid PVST+ protocol.
    // edge -- this port is an edge port for the RST region.  unused1 -- unused
    // bit 1.  unused2 -- unused bit 2.  stp -- this port is connected to a Single
    // Spanning        Tree/PVST+ bridge.  dispute -- this port, as a designated
    // port, received an        inferior BPDU with a designated role and the      
    // learning bit being set. The type is map[string]bool.
    StpxRPVSTPortStatus interface{}
}

func (stpxRPVSTPortEntry *CISCOSTPEXTENSIONSMIB_StpxRPVSTPortTable_StpxRPVSTPortEntry) GetEntityData() *types.CommonEntityData {
    stpxRPVSTPortEntry.EntityData.YFilter = stpxRPVSTPortEntry.YFilter
    stpxRPVSTPortEntry.EntityData.YangName = "stpxRPVSTPortEntry"
    stpxRPVSTPortEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxRPVSTPortEntry.EntityData.ParentYangName = "stpxRPVSTPortTable"
    stpxRPVSTPortEntry.EntityData.SegmentPath = "stpxRPVSTPortEntry" + types.AddKeyToken(stpxRPVSTPortEntry.StpxRPVSTPortVlanIndex, "stpxRPVSTPortVlanIndex") + types.AddKeyToken(stpxRPVSTPortEntry.StpxRPVSTPortIndex, "stpxRPVSTPortIndex")
    stpxRPVSTPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxRPVSTPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxRPVSTPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxRPVSTPortEntry.EntityData.Children = types.NewOrderedMap()
    stpxRPVSTPortEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxRPVSTPortEntry.EntityData.Leafs.Append("stpxRPVSTPortVlanIndex", types.YLeaf{"StpxRPVSTPortVlanIndex", stpxRPVSTPortEntry.StpxRPVSTPortVlanIndex})
    stpxRPVSTPortEntry.EntityData.Leafs.Append("stpxRPVSTPortIndex", types.YLeaf{"StpxRPVSTPortIndex", stpxRPVSTPortEntry.StpxRPVSTPortIndex})
    stpxRPVSTPortEntry.EntityData.Leafs.Append("stpxRPVSTPortStatus", types.YLeaf{"StpxRPVSTPortStatus", stpxRPVSTPortEntry.StpxRPVSTPortStatus})

    stpxRPVSTPortEntry.EntityData.YListKeys = []string {"StpxRPVSTPortVlanIndex", "StpxRPVSTPortIndex"}

    return &(stpxRPVSTPortEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable
// This table contains MST instance information
// for IEEE MST.
type CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing the MST instance  information for IEEE MST. The
    // type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable_StpxSMSTInstanceEntry.
    StpxSMSTInstanceEntry []*CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable_StpxSMSTInstanceEntry
}

func (stpxSMSTInstanceTable *CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable) GetEntityData() *types.CommonEntityData {
    stpxSMSTInstanceTable.EntityData.YFilter = stpxSMSTInstanceTable.YFilter
    stpxSMSTInstanceTable.EntityData.YangName = "stpxSMSTInstanceTable"
    stpxSMSTInstanceTable.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTInstanceTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxSMSTInstanceTable.EntityData.SegmentPath = "stpxSMSTInstanceTable"
    stpxSMSTInstanceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTInstanceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTInstanceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTInstanceTable.EntityData.Children = types.NewOrderedMap()
    stpxSMSTInstanceTable.EntityData.Children.Append("stpxSMSTInstanceEntry", types.YChild{"StpxSMSTInstanceEntry", nil})
    for i := range stpxSMSTInstanceTable.StpxSMSTInstanceEntry {
        stpxSMSTInstanceTable.EntityData.Children.Append(types.GetSegmentPath(stpxSMSTInstanceTable.StpxSMSTInstanceEntry[i]), types.YChild{"StpxSMSTInstanceEntry", stpxSMSTInstanceTable.StpxSMSTInstanceEntry[i]})
    }
    stpxSMSTInstanceTable.EntityData.Leafs = types.NewOrderedMap()

    stpxSMSTInstanceTable.EntityData.YListKeys = []string {}

    return &(stpxSMSTInstanceTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable_StpxSMSTInstanceEntry
// A conceptual row containing the MST instance 
// information for IEEE MST.
type CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable_StpxSMSTInstanceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance ID, the value of which is in the
    // range  from 0 to stpxSMSTMaxInstanceID. The type is interface{} with range:
    // 0..4294967295.
    StpxSMSTInstanceIndex interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 0 through 2047. The first octet corresponds to VLANs with
    // VlanIndex values of 0 through 7; the second octet to VLANs 8 through 15;
    // etc.  The most significant bit of each octet corresponds to the lowest
    // value VlanIndex in that octet.  For each VLAN, if it is mapped to this MST
    // instance, then the bit corresponding to that VLAN is set to '1'. If the
    // length of this string is less than 256 octets, any 'missing' octets are
    // assumed to contain the value  of zero. The type is string with length:
    // 0..256.
    StpxSMSTInstanceVlansMapped1k2k interface{}

    // A string of octets containing one bit per VLAN for VLANS with VlanIndex
    // values of 2048 through 4095. The first octet corresponds to VLANs with
    // VlanIndex values of 2048 through 2055; the second octet to VLANs 2056
    // through 2063; etc.  The most significant bit of each octet corresponds to
    // the lowest value VlanIndex in that octet.  For each VLAN, if it is mapped
    // to this MST instance, then the bit corresponding to that VLAN is set to
    // '1'. If the length of this string is less than 256 octets, any 'missing'
    // octets are assumed to contain the value  of zero. The type is string with
    // length: 0..256.
    StpxSMSTInstanceVlansMapped3k4k interface{}

    // The remaining hop count for this MST instance. If this object value is not
    // applicable on an MST instance, then the value retrieved for this object for
    // that MST instance will be -1.   This object is only instantiated when the
    // object value of stpxSpanningTreeType is mst(4). The type is interface{}
    // with range: -1..2147483647.
    StpxSMSTInstanceRemainingHopCount interface{}

    // Indicates the Bridge Identifier (refer to BridgeId  defined in BRIDGE-MIB)
    // of CIST (Common and Internal  Spanning Tree) Regional Root for the MST
    // region.  This object is only instantiated when the object value of
    // stpxSpanningTreeType is mst(4) and stpxSMSTInstanceIndex is 0. The type is
    // string with length: 8.
    StpxSMSTInstanceCISTRegionalRoot interface{}

    // Indicates the CIST Internal Root Path Cost, i.e., the path cost to the CIST
    // Regional Root as specified by the corresponding
    // stpxSMSTInstanceCISTRegionalRoot for the  MST region.  This object is only
    // instantiated when the object value of stpxSpanningTreeType is mst(4) and
    // stpxSMSTInstanceIndex is 0. The type is interface{} with range:
    // 0..4294967295.
    StpxSMSTInstanceCISTIntRootCost interface{}
}

func (stpxSMSTInstanceEntry *CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceTable_StpxSMSTInstanceEntry) GetEntityData() *types.CommonEntityData {
    stpxSMSTInstanceEntry.EntityData.YFilter = stpxSMSTInstanceEntry.YFilter
    stpxSMSTInstanceEntry.EntityData.YangName = "stpxSMSTInstanceEntry"
    stpxSMSTInstanceEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTInstanceEntry.EntityData.ParentYangName = "stpxSMSTInstanceTable"
    stpxSMSTInstanceEntry.EntityData.SegmentPath = "stpxSMSTInstanceEntry" + types.AddKeyToken(stpxSMSTInstanceEntry.StpxSMSTInstanceIndex, "stpxSMSTInstanceIndex")
    stpxSMSTInstanceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTInstanceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTInstanceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTInstanceEntry.EntityData.Children = types.NewOrderedMap()
    stpxSMSTInstanceEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceIndex", types.YLeaf{"StpxSMSTInstanceIndex", stpxSMSTInstanceEntry.StpxSMSTInstanceIndex})
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceVlansMapped1k2k", types.YLeaf{"StpxSMSTInstanceVlansMapped1k2k", stpxSMSTInstanceEntry.StpxSMSTInstanceVlansMapped1k2k})
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceVlansMapped3k4k", types.YLeaf{"StpxSMSTInstanceVlansMapped3k4k", stpxSMSTInstanceEntry.StpxSMSTInstanceVlansMapped3k4k})
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceRemainingHopCount", types.YLeaf{"StpxSMSTInstanceRemainingHopCount", stpxSMSTInstanceEntry.StpxSMSTInstanceRemainingHopCount})
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceCISTRegionalRoot", types.YLeaf{"StpxSMSTInstanceCISTRegionalRoot", stpxSMSTInstanceEntry.StpxSMSTInstanceCISTRegionalRoot})
    stpxSMSTInstanceEntry.EntityData.Leafs.Append("stpxSMSTInstanceCISTIntRootCost", types.YLeaf{"StpxSMSTInstanceCISTIntRootCost", stpxSMSTInstanceEntry.StpxSMSTInstanceCISTIntRootCost})

    stpxSMSTInstanceEntry.EntityData.YListKeys = []string {"StpxSMSTInstanceIndex"}

    return &(stpxSMSTInstanceEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable
// This table contains MST instance information in the 
// Edit Buffer. 
// 
// This table is only instantiated when the object value
// of  stpxMSTRegionEditBufferStatus has the value of
// acquiredBySnmp(2).
type CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row containing MST instance information  in the Edit Buffer. 
    // The total number of entries in this table has to be  less than or equal to
    // the object value of stpxSMSTMaxInstances. The type is slice of
    // CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable_StpxSMSTInstanceEditEntry.
    StpxSMSTInstanceEditEntry []*CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable_StpxSMSTInstanceEditEntry
}

func (stpxSMSTInstanceEditTable *CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable) GetEntityData() *types.CommonEntityData {
    stpxSMSTInstanceEditTable.EntityData.YFilter = stpxSMSTInstanceEditTable.YFilter
    stpxSMSTInstanceEditTable.EntityData.YangName = "stpxSMSTInstanceEditTable"
    stpxSMSTInstanceEditTable.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTInstanceEditTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxSMSTInstanceEditTable.EntityData.SegmentPath = "stpxSMSTInstanceEditTable"
    stpxSMSTInstanceEditTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTInstanceEditTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTInstanceEditTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTInstanceEditTable.EntityData.Children = types.NewOrderedMap()
    stpxSMSTInstanceEditTable.EntityData.Children.Append("stpxSMSTInstanceEditEntry", types.YChild{"StpxSMSTInstanceEditEntry", nil})
    for i := range stpxSMSTInstanceEditTable.StpxSMSTInstanceEditEntry {
        stpxSMSTInstanceEditTable.EntityData.Children.Append(types.GetSegmentPath(stpxSMSTInstanceEditTable.StpxSMSTInstanceEditEntry[i]), types.YChild{"StpxSMSTInstanceEditEntry", stpxSMSTInstanceEditTable.StpxSMSTInstanceEditEntry[i]})
    }
    stpxSMSTInstanceEditTable.EntityData.Leafs = types.NewOrderedMap()

    stpxSMSTInstanceEditTable.EntityData.YListKeys = []string {}

    return &(stpxSMSTInstanceEditTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable_StpxSMSTInstanceEditEntry
// A conceptual row containing MST instance information 
// in the Edit Buffer.
// 
// The total number of entries in this table has to be 
// less than or equal to the object value of stpxSMSTMaxInstances.
type CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable_StpxSMSTInstanceEditEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MST instance ID, the value of which is in the
    // range from 0 to stpxSMSTMaxInstanceID.   The instances of this table entry
    // with  stpxSMSTInstanceEditIndex of zero is automatically created by the
    // device and can not modified. The type is interface{} with range:
    // 0..4294967295.
    StpxSMSTInstanceEditIndex interface{}

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
    StpxSMSTInstanceEditVlansMap1k2k interface{}

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
    StpxSMSTInstanceEditVlansMap3k4k interface{}

    // This object controls the creation and deletion of a row  in
    // stpxSMSTInstanceEditTable.  When creating an entry in this table,
    // 'createAndGo' method is used and the value of this object is set to
    // 'active'. Deactivation of an 'active' entry is not allowed.  When  deleting
    // an entry in this table, 'destroy' method is used.  Once a row becomes
    // active, value in any other column  within such a row may be modified. When
    // a row is active,  setting the instance of stpxSMSTInstanceEditVlansMap1k2k
    // stpxSMSTInstanceEditVlansMap3k4k for the same MST instance both to the
    // value of zero length can not be allowed. The type is RowStatus.
    StpxSMSTInstanceEditRowStatus interface{}
}

func (stpxSMSTInstanceEditEntry *CISCOSTPEXTENSIONSMIB_StpxSMSTInstanceEditTable_StpxSMSTInstanceEditEntry) GetEntityData() *types.CommonEntityData {
    stpxSMSTInstanceEditEntry.EntityData.YFilter = stpxSMSTInstanceEditEntry.YFilter
    stpxSMSTInstanceEditEntry.EntityData.YangName = "stpxSMSTInstanceEditEntry"
    stpxSMSTInstanceEditEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTInstanceEditEntry.EntityData.ParentYangName = "stpxSMSTInstanceEditTable"
    stpxSMSTInstanceEditEntry.EntityData.SegmentPath = "stpxSMSTInstanceEditEntry" + types.AddKeyToken(stpxSMSTInstanceEditEntry.StpxSMSTInstanceEditIndex, "stpxSMSTInstanceEditIndex")
    stpxSMSTInstanceEditEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTInstanceEditEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTInstanceEditEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTInstanceEditEntry.EntityData.Children = types.NewOrderedMap()
    stpxSMSTInstanceEditEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxSMSTInstanceEditEntry.EntityData.Leafs.Append("stpxSMSTInstanceEditIndex", types.YLeaf{"StpxSMSTInstanceEditIndex", stpxSMSTInstanceEditEntry.StpxSMSTInstanceEditIndex})
    stpxSMSTInstanceEditEntry.EntityData.Leafs.Append("stpxSMSTInstanceEditVlansMap1k2k", types.YLeaf{"StpxSMSTInstanceEditVlansMap1k2k", stpxSMSTInstanceEditEntry.StpxSMSTInstanceEditVlansMap1k2k})
    stpxSMSTInstanceEditEntry.EntityData.Leafs.Append("stpxSMSTInstanceEditVlansMap3k4k", types.YLeaf{"StpxSMSTInstanceEditVlansMap3k4k", stpxSMSTInstanceEditEntry.StpxSMSTInstanceEditVlansMap3k4k})
    stpxSMSTInstanceEditEntry.EntityData.Leafs.Append("stpxSMSTInstanceEditRowStatus", types.YLeaf{"StpxSMSTInstanceEditRowStatus", stpxSMSTInstanceEditEntry.StpxSMSTInstanceEditRowStatus})

    stpxSMSTInstanceEditEntry.EntityData.YListKeys = []string {"StpxSMSTInstanceEditIndex"}

    return &(stpxSMSTInstanceEditEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable
// A table containing port information for the MST 
// Protocol on all the bridge ports existing on the 
// system.
// 
// This table is only instantiated when the object 
// value of stpxSpanningTreeType is mst(4)
type CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry with port information for the MST protocol on a bridge port. The
    // type is slice of CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry.
    StpxSMSTPortEntry []*CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry
}

func (stpxSMSTPortTable *CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable) GetEntityData() *types.CommonEntityData {
    stpxSMSTPortTable.EntityData.YFilter = stpxSMSTPortTable.YFilter
    stpxSMSTPortTable.EntityData.YangName = "stpxSMSTPortTable"
    stpxSMSTPortTable.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTPortTable.EntityData.ParentYangName = "CISCO-STP-EXTENSIONS-MIB"
    stpxSMSTPortTable.EntityData.SegmentPath = "stpxSMSTPortTable"
    stpxSMSTPortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTPortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTPortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTPortTable.EntityData.Children = types.NewOrderedMap()
    stpxSMSTPortTable.EntityData.Children.Append("stpxSMSTPortEntry", types.YChild{"StpxSMSTPortEntry", nil})
    for i := range stpxSMSTPortTable.StpxSMSTPortEntry {
        stpxSMSTPortTable.EntityData.Children.Append(types.GetSegmentPath(stpxSMSTPortTable.StpxSMSTPortEntry[i]), types.YChild{"StpxSMSTPortEntry", stpxSMSTPortTable.StpxSMSTPortEntry[i]})
    }
    stpxSMSTPortTable.EntityData.Leafs = types.NewOrderedMap()

    stpxSMSTPortTable.EntityData.YListKeys = []string {}

    return &(stpxSMSTPortTable.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry
// An entry with port information for the MST protocol
// on a bridge port.
type CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of dot1dBasePort (i.e. dot1dBridge.1.4)
    // for the bridge port. The type is interface{} with range: 1..65535.
    StpxSMSTPortIndex interface{}

    // Indicates the operational status of the port for the  MST protocol.   edge
    // -- this port is an edge port for the MST region.  boundary -- this port is
    // a boundary port for the          MST region.  pvst --  this port is
    // connected to a PVST/PVST+ bridge.     stp -- this port is connected to a
    // Single Spanning         Tree bridge.  dispute -- this port, as a designated
    // port, received an         inferior BPDU with a designated role and the     
    // learning bit being set.  rstp -- this port is connected to a RSTP bridge or
    // an          MST bridge in a different MST region. The type is
    // map[string]bool.
    StpxSMSTPortStatus interface{}

    // The adminitratively configured hello time in hundredth  of seconds on a
    // port for IEEE MST. The granularity  of this timer is 1 second. An agent may
    // return a badValue  error if a set is attempted to a value which is not a 
    // whole number of seconds. This object value of zero means the hello time is
    // not specifically configured on  this port and object value of
    // stpxSMSTPortConfigedHelloTime retrieved for this port will take on the
    // value of  dot1dStpBridgeHelloTime defined in BRIDGE-MIB. The type is
    // interface{} with range: 0..4294967295. Units are hundredth of seconds.
    StpxSMSTPortAdminHelloTime interface{}

    // Indicates the effective configuration of the hello time on  a port. The
    // type is interface{} with range: 0..4294967295. Units are hundredth of
    // seconds.
    StpxSMSTPortConfigedHelloTime interface{}

    // The operational hello time in hundredth of seconds on a  port for IEEE MST.
    // If this object value is not applicable on a port, then the value retrieved
    // on that port will be -1. The type is interface{} with range:
    // -1..2147483647. Units are hundredth of seconds.
    StpxSMSTPortOperHelloTime interface{}

    // The desired MST mode of this port.  preStandard -- this port is
    // administratively configured to     transmit pre-standard, i.e. pre IEEE
    // MST, BPDUs.  auto -- the BPDU transmission mode of this port is based     
    // on automatic detection of neighbor ports. The type is
    // StpxSMSTPortAdminMSTMode.
    StpxSMSTPortAdminMSTMode interface{}

    // Indicates the current operational MST mode of this port.  unknown -- the
    // operational mode is currently unknown.  preStandard -- this port is
    // currently operating in      pre-standard MSTP BPDU transmission mode. 
    // standard -- this port is currently operating in IEEE MST      BPDU
    // transmission mode. The type is StpxSMSTPortOperMSTMode.
    StpxSMSTPortOperMSTMode interface{}
}

func (stpxSMSTPortEntry *CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry) GetEntityData() *types.CommonEntityData {
    stpxSMSTPortEntry.EntityData.YFilter = stpxSMSTPortEntry.YFilter
    stpxSMSTPortEntry.EntityData.YangName = "stpxSMSTPortEntry"
    stpxSMSTPortEntry.EntityData.BundleName = "cisco_ios_xe"
    stpxSMSTPortEntry.EntityData.ParentYangName = "stpxSMSTPortTable"
    stpxSMSTPortEntry.EntityData.SegmentPath = "stpxSMSTPortEntry" + types.AddKeyToken(stpxSMSTPortEntry.StpxSMSTPortIndex, "stpxSMSTPortIndex")
    stpxSMSTPortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stpxSMSTPortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stpxSMSTPortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stpxSMSTPortEntry.EntityData.Children = types.NewOrderedMap()
    stpxSMSTPortEntry.EntityData.Leafs = types.NewOrderedMap()
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortIndex", types.YLeaf{"StpxSMSTPortIndex", stpxSMSTPortEntry.StpxSMSTPortIndex})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortStatus", types.YLeaf{"StpxSMSTPortStatus", stpxSMSTPortEntry.StpxSMSTPortStatus})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortAdminHelloTime", types.YLeaf{"StpxSMSTPortAdminHelloTime", stpxSMSTPortEntry.StpxSMSTPortAdminHelloTime})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortConfigedHelloTime", types.YLeaf{"StpxSMSTPortConfigedHelloTime", stpxSMSTPortEntry.StpxSMSTPortConfigedHelloTime})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortOperHelloTime", types.YLeaf{"StpxSMSTPortOperHelloTime", stpxSMSTPortEntry.StpxSMSTPortOperHelloTime})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortAdminMSTMode", types.YLeaf{"StpxSMSTPortAdminMSTMode", stpxSMSTPortEntry.StpxSMSTPortAdminMSTMode})
    stpxSMSTPortEntry.EntityData.Leafs.Append("stpxSMSTPortOperMSTMode", types.YLeaf{"StpxSMSTPortOperMSTMode", stpxSMSTPortEntry.StpxSMSTPortOperMSTMode})

    stpxSMSTPortEntry.EntityData.YListKeys = []string {"StpxSMSTPortIndex"}

    return &(stpxSMSTPortEntry.EntityData)
}

// CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode represents     on automatic detection of neighbor ports.
type CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode_preStandard CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode = "preStandard"

    CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode_auto CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortAdminMSTMode = "auto"
)

// CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode represents     BPDU transmission mode.
type CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode string

const (
    CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode_unknown CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode = "unknown"

    CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode_preStandard CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode = "preStandard"

    CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode_standard CISCOSTPEXTENSIONSMIB_StpxSMSTPortTable_StpxSMSTPortEntry_StpxSMSTPortOperMSTMode = "standard"
)

