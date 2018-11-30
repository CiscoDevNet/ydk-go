// This MIB module contains the management objects for the
// management of fiber nodes in the Cable plant.
package clab_topo_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package clab_topo_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CLAB-TOPO-MIB CLAB-TOPO-MIB}", reflect.TypeOf(CLABTOPOMIB{}))
    ydk.RegisterEntity("CLAB-TOPO-MIB:CLAB-TOPO-MIB", reflect.TypeOf(CLABTOPOMIB{}))
}

// CLABTOPOMIB
type CLABTOPOMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object defines the cable HFC plant Fiber Nodes known at a CMTS. This
    // object supports the creation and deletion of multiple instances.
    ClabTopoFiberNodeCfgTable CLABTOPOMIB_ClabTopoFiberNodeCfgTable

    // This object defines the RF topology by defining the connectivity of a
    // CMTS's downstream and upstream channels to the fiber nodes. Each instance
    // of this object describes connectivity of one downstream or upstream channel
    // with a single fiber node. This object supports the creation and deletion of
    // multiple instances.
    ClabTopoChFnCfgTable CLABTOPOMIB_ClabTopoChFnCfgTable
}

func (cLABTOPOMIB *CLABTOPOMIB) GetEntityData() *types.CommonEntityData {
    cLABTOPOMIB.EntityData.YFilter = cLABTOPOMIB.YFilter
    cLABTOPOMIB.EntityData.YangName = "CLAB-TOPO-MIB"
    cLABTOPOMIB.EntityData.BundleName = "cisco_ios_xe"
    cLABTOPOMIB.EntityData.ParentYangName = "CLAB-TOPO-MIB"
    cLABTOPOMIB.EntityData.SegmentPath = "CLAB-TOPO-MIB:CLAB-TOPO-MIB"
    cLABTOPOMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cLABTOPOMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cLABTOPOMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cLABTOPOMIB.EntityData.Children = types.NewOrderedMap()
    cLABTOPOMIB.EntityData.Children.Append("clabTopoFiberNodeCfgTable", types.YChild{"ClabTopoFiberNodeCfgTable", &cLABTOPOMIB.ClabTopoFiberNodeCfgTable})
    cLABTOPOMIB.EntityData.Children.Append("clabTopoChFnCfgTable", types.YChild{"ClabTopoChFnCfgTable", &cLABTOPOMIB.ClabTopoChFnCfgTable})
    cLABTOPOMIB.EntityData.Leafs = types.NewOrderedMap()

    cLABTOPOMIB.EntityData.YListKeys = []string {}

    return &(cLABTOPOMIB.EntityData)
}

// CLABTOPOMIB_ClabTopoFiberNodeCfgTable
// This object defines the cable HFC plant Fiber Nodes
// known at a CMTS.
// This object supports the creation and deletion of multiple
// instances.
type CLABTOPOMIB_ClabTopoFiberNodeCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of clabTopoFiberNodeCfg. The CMTS persists all instances
    // of FiberNodeCfg across reinitializations. The type is slice of
    // CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry.
    ClabTopoFiberNodeCfgEntry []*CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry
}

func (clabTopoFiberNodeCfgTable *CLABTOPOMIB_ClabTopoFiberNodeCfgTable) GetEntityData() *types.CommonEntityData {
    clabTopoFiberNodeCfgTable.EntityData.YFilter = clabTopoFiberNodeCfgTable.YFilter
    clabTopoFiberNodeCfgTable.EntityData.YangName = "clabTopoFiberNodeCfgTable"
    clabTopoFiberNodeCfgTable.EntityData.BundleName = "cisco_ios_xe"
    clabTopoFiberNodeCfgTable.EntityData.ParentYangName = "CLAB-TOPO-MIB"
    clabTopoFiberNodeCfgTable.EntityData.SegmentPath = "clabTopoFiberNodeCfgTable"
    clabTopoFiberNodeCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clabTopoFiberNodeCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clabTopoFiberNodeCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clabTopoFiberNodeCfgTable.EntityData.Children = types.NewOrderedMap()
    clabTopoFiberNodeCfgTable.EntityData.Children.Append("clabTopoFiberNodeCfgEntry", types.YChild{"ClabTopoFiberNodeCfgEntry", nil})
    for i := range clabTopoFiberNodeCfgTable.ClabTopoFiberNodeCfgEntry {
        clabTopoFiberNodeCfgTable.EntityData.Children.Append(types.GetSegmentPath(clabTopoFiberNodeCfgTable.ClabTopoFiberNodeCfgEntry[i]), types.YChild{"ClabTopoFiberNodeCfgEntry", clabTopoFiberNodeCfgTable.ClabTopoFiberNodeCfgEntry[i]})
    }
    clabTopoFiberNodeCfgTable.EntityData.Leafs = types.NewOrderedMap()

    clabTopoFiberNodeCfgTable.EntityData.YListKeys = []string {}

    return &(clabTopoFiberNodeCfgTable.EntityData)
}

// CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry
// The conceptual row of clabTopoFiberNodeCfg.
// The CMTS persists all instances of FiberNodeCfg
// across reinitializations.
type CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This key represents a human-readable name for a
    // fiber node. The type is string with length: 1..16.
    ClabTopoFiberNodeCfgNodeName interface{}

    // Administratively configured human-readable description of the fiber node.
    // The type is string.
    ClabTopoFiberNodeCfgNodeDescr interface{}

    // The status of this instance. The type is RowStatus.
    ClabTopoFiberNodeCfgRowStatus interface{}
}

func (clabTopoFiberNodeCfgEntry *CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry) GetEntityData() *types.CommonEntityData {
    clabTopoFiberNodeCfgEntry.EntityData.YFilter = clabTopoFiberNodeCfgEntry.YFilter
    clabTopoFiberNodeCfgEntry.EntityData.YangName = "clabTopoFiberNodeCfgEntry"
    clabTopoFiberNodeCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    clabTopoFiberNodeCfgEntry.EntityData.ParentYangName = "clabTopoFiberNodeCfgTable"
    clabTopoFiberNodeCfgEntry.EntityData.SegmentPath = "clabTopoFiberNodeCfgEntry" + types.AddKeyToken(clabTopoFiberNodeCfgEntry.ClabTopoFiberNodeCfgNodeName, "clabTopoFiberNodeCfgNodeName")
    clabTopoFiberNodeCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clabTopoFiberNodeCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clabTopoFiberNodeCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clabTopoFiberNodeCfgEntry.EntityData.Children = types.NewOrderedMap()
    clabTopoFiberNodeCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    clabTopoFiberNodeCfgEntry.EntityData.Leafs.Append("clabTopoFiberNodeCfgNodeName", types.YLeaf{"ClabTopoFiberNodeCfgNodeName", clabTopoFiberNodeCfgEntry.ClabTopoFiberNodeCfgNodeName})
    clabTopoFiberNodeCfgEntry.EntityData.Leafs.Append("clabTopoFiberNodeCfgNodeDescr", types.YLeaf{"ClabTopoFiberNodeCfgNodeDescr", clabTopoFiberNodeCfgEntry.ClabTopoFiberNodeCfgNodeDescr})
    clabTopoFiberNodeCfgEntry.EntityData.Leafs.Append("clabTopoFiberNodeCfgRowStatus", types.YLeaf{"ClabTopoFiberNodeCfgRowStatus", clabTopoFiberNodeCfgEntry.ClabTopoFiberNodeCfgRowStatus})

    clabTopoFiberNodeCfgEntry.EntityData.YListKeys = []string {"ClabTopoFiberNodeCfgNodeName"}

    return &(clabTopoFiberNodeCfgEntry.EntityData)
}

// CLABTOPOMIB_ClabTopoChFnCfgTable
// This object defines the RF topology by defining the
// connectivity of a CMTS's downstream and upstream channels
// to the fiber nodes. Each instance of this object
// describes connectivity of one downstream or upstream
// channel with a single fiber node.
// This object supports the creation and deletion of multiple
// instances.
type CLABTOPOMIB_ClabTopoChFnCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of clabTopoChFnCfg. The CMTS persists all instances of
    // ChFnCfg across reinitializations. The type is slice of
    // CLABTOPOMIB_ClabTopoChFnCfgTable_ClabTopoChFnCfgEntry.
    ClabTopoChFnCfgEntry []*CLABTOPOMIB_ClabTopoChFnCfgTable_ClabTopoChFnCfgEntry
}

func (clabTopoChFnCfgTable *CLABTOPOMIB_ClabTopoChFnCfgTable) GetEntityData() *types.CommonEntityData {
    clabTopoChFnCfgTable.EntityData.YFilter = clabTopoChFnCfgTable.YFilter
    clabTopoChFnCfgTable.EntityData.YangName = "clabTopoChFnCfgTable"
    clabTopoChFnCfgTable.EntityData.BundleName = "cisco_ios_xe"
    clabTopoChFnCfgTable.EntityData.ParentYangName = "CLAB-TOPO-MIB"
    clabTopoChFnCfgTable.EntityData.SegmentPath = "clabTopoChFnCfgTable"
    clabTopoChFnCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clabTopoChFnCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clabTopoChFnCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clabTopoChFnCfgTable.EntityData.Children = types.NewOrderedMap()
    clabTopoChFnCfgTable.EntityData.Children.Append("clabTopoChFnCfgEntry", types.YChild{"ClabTopoChFnCfgEntry", nil})
    for i := range clabTopoChFnCfgTable.ClabTopoChFnCfgEntry {
        clabTopoChFnCfgTable.EntityData.Children.Append(types.GetSegmentPath(clabTopoChFnCfgTable.ClabTopoChFnCfgEntry[i]), types.YChild{"ClabTopoChFnCfgEntry", clabTopoChFnCfgTable.ClabTopoChFnCfgEntry[i]})
    }
    clabTopoChFnCfgTable.EntityData.Leafs = types.NewOrderedMap()

    clabTopoChFnCfgTable.EntityData.YListKeys = []string {}

    return &(clabTopoChFnCfgTable.EntityData)
}

// CLABTOPOMIB_ClabTopoChFnCfgTable_ClabTopoChFnCfgEntry
// The conceptual row of clabTopoChFnCfg.
// The CMTS persists all instances of ChFnCfg
// across reinitializations.
type CLABTOPOMIB_ClabTopoChFnCfgTable_ClabTopoChFnCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..16. Refers to
    // clab_topo_mib.CLABTOPOMIB_ClabTopoFiberNodeCfgTable_ClabTopoFiberNodeCfgEntry_ClabTopoFiberNodeCfgNodeName
    ClabTopoFiberNodeCfgNodeName interface{}

    // This attribute is a key. This key represents the interface index of an
    // upstream or downstream channel associated with this fiber node. In the
    // upstream direction, only ifIndices docsCableUpstream channels are
    // reflected. The type is interface{} with range: 1..2147483647.
    ClabTopoChFnCfgChIfIndex interface{}

    // The status of this instance. The type is RowStatus.
    ClabTopoChFnCfgRowStatus interface{}
}

func (clabTopoChFnCfgEntry *CLABTOPOMIB_ClabTopoChFnCfgTable_ClabTopoChFnCfgEntry) GetEntityData() *types.CommonEntityData {
    clabTopoChFnCfgEntry.EntityData.YFilter = clabTopoChFnCfgEntry.YFilter
    clabTopoChFnCfgEntry.EntityData.YangName = "clabTopoChFnCfgEntry"
    clabTopoChFnCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    clabTopoChFnCfgEntry.EntityData.ParentYangName = "clabTopoChFnCfgTable"
    clabTopoChFnCfgEntry.EntityData.SegmentPath = "clabTopoChFnCfgEntry" + types.AddKeyToken(clabTopoChFnCfgEntry.ClabTopoFiberNodeCfgNodeName, "clabTopoFiberNodeCfgNodeName") + types.AddKeyToken(clabTopoChFnCfgEntry.ClabTopoChFnCfgChIfIndex, "clabTopoChFnCfgChIfIndex")
    clabTopoChFnCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clabTopoChFnCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clabTopoChFnCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clabTopoChFnCfgEntry.EntityData.Children = types.NewOrderedMap()
    clabTopoChFnCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    clabTopoChFnCfgEntry.EntityData.Leafs.Append("clabTopoFiberNodeCfgNodeName", types.YLeaf{"ClabTopoFiberNodeCfgNodeName", clabTopoChFnCfgEntry.ClabTopoFiberNodeCfgNodeName})
    clabTopoChFnCfgEntry.EntityData.Leafs.Append("clabTopoChFnCfgChIfIndex", types.YLeaf{"ClabTopoChFnCfgChIfIndex", clabTopoChFnCfgEntry.ClabTopoChFnCfgChIfIndex})
    clabTopoChFnCfgEntry.EntityData.Leafs.Append("clabTopoChFnCfgRowStatus", types.YLeaf{"ClabTopoChFnCfgRowStatus", clabTopoChFnCfgEntry.ClabTopoChFnCfgRowStatus})

    clabTopoChFnCfgEntry.EntityData.YListKeys = []string {"ClabTopoFiberNodeCfgNodeName", "ClabTopoChFnCfgChIfIndex"}

    return &(clabTopoChFnCfgEntry.EntityData)
}

