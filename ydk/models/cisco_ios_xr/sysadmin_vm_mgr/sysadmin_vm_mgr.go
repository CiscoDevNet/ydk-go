// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2012-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_vm_mgr

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_vm_mgr"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-vm-mgr VM}", reflect.TypeOf(VM{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-vm-mgr:VM", reflect.TypeOf(VM{}))
}

// VM
// VM Info
type VM struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of VM_AllLocations.
    AllLocations []*VM_AllLocations
}

func (vM *VM) GetEntityData() *types.CommonEntityData {
    vM.EntityData.YFilter = vM.YFilter
    vM.EntityData.YangName = "VM"
    vM.EntityData.BundleName = "cisco_ios_xr"
    vM.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-vm-mgr"
    vM.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-vm-mgr:VM"
    vM.EntityData.AbsolutePath = vM.EntityData.SegmentPath
    vM.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vM.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vM.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vM.EntityData.Children = types.NewOrderedMap()
    vM.EntityData.Children.Append("all-locations", types.YChild{"AllLocations", nil})
    for i := range vM.AllLocations {
        vM.EntityData.Children.Append(types.GetSegmentPath(vM.AllLocations[i]), types.YChild{"AllLocations", vM.AllLocations[i]})
    }
    vM.EntityData.Leafs = types.NewOrderedMap()

    vM.EntityData.YListKeys = []string {}

    return &(vM.EntityData)
}

// VM_AllLocations
type VM_AllLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Location interface{}

    // The type is slice of VM_AllLocations_AllUiids.
    AllUiids []*VM_AllLocations_AllUiids
}

func (allLocations *VM_AllLocations) GetEntityData() *types.CommonEntityData {
    allLocations.EntityData.YFilter = allLocations.YFilter
    allLocations.EntityData.YangName = "all-locations"
    allLocations.EntityData.BundleName = "cisco_ios_xr"
    allLocations.EntityData.ParentYangName = "VM"
    allLocations.EntityData.SegmentPath = "all-locations" + types.AddKeyToken(allLocations.Location, "location")
    allLocations.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm-mgr:VM/" + allLocations.EntityData.SegmentPath
    allLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLocations.EntityData.Children = types.NewOrderedMap()
    allLocations.EntityData.Children.Append("all-uiids", types.YChild{"AllUiids", nil})
    for i := range allLocations.AllUiids {
        allLocations.EntityData.Children.Append(types.GetSegmentPath(allLocations.AllUiids[i]), types.YChild{"AllUiids", allLocations.AllUiids[i]})
    }
    allLocations.EntityData.Leafs = types.NewOrderedMap()
    allLocations.EntityData.Leafs.Append("location", types.YLeaf{"Location", allLocations.Location})

    allLocations.EntityData.YListKeys = []string {"Location"}

    return &(allLocations.EntityData)
}

// VM_AllLocations_AllUiids
type VM_AllLocations_AllUiids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique Immutable ID. The type is string.
    Uiid interface{}

    // ID of the VM. The type is string.
    Id interface{}

    // Status of the VM. The type is string.
    Status interface{}

    // CE IP address. The type is string.
    Ipaddr interface{}

    // Heartbeat interval sec. The type is string.
    HbIntervalS interface{}

    // Heartbeat interval nsec. The type is string.
    HbIntervalNs interface{}

    // Last heartbeat sent. The type is string.
    LastHbSent interface{}

    // Last heartbeat received. The type is string.
    LastHbRec interface{}

    // ISSU role. The type is string.
    Role interface{}
}

func (allUiids *VM_AllLocations_AllUiids) GetEntityData() *types.CommonEntityData {
    allUiids.EntityData.YFilter = allUiids.YFilter
    allUiids.EntityData.YangName = "all-uiids"
    allUiids.EntityData.BundleName = "cisco_ios_xr"
    allUiids.EntityData.ParentYangName = "all-locations"
    allUiids.EntityData.SegmentPath = "all-uiids" + types.AddKeyToken(allUiids.Uiid, "uiid")
    allUiids.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-vm-mgr:VM/all-locations/" + allUiids.EntityData.SegmentPath
    allUiids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allUiids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allUiids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allUiids.EntityData.Children = types.NewOrderedMap()
    allUiids.EntityData.Leafs = types.NewOrderedMap()
    allUiids.EntityData.Leafs.Append("uiid", types.YLeaf{"Uiid", allUiids.Uiid})
    allUiids.EntityData.Leafs.Append("id", types.YLeaf{"Id", allUiids.Id})
    allUiids.EntityData.Leafs.Append("status", types.YLeaf{"Status", allUiids.Status})
    allUiids.EntityData.Leafs.Append("ipaddr", types.YLeaf{"Ipaddr", allUiids.Ipaddr})
    allUiids.EntityData.Leafs.Append("hb_interval_s", types.YLeaf{"HbIntervalS", allUiids.HbIntervalS})
    allUiids.EntityData.Leafs.Append("hb_interval_ns", types.YLeaf{"HbIntervalNs", allUiids.HbIntervalNs})
    allUiids.EntityData.Leafs.Append("last_hb_sent", types.YLeaf{"LastHbSent", allUiids.LastHbSent})
    allUiids.EntityData.Leafs.Append("last_hb_rec", types.YLeaf{"LastHbRec", allUiids.LastHbRec})
    allUiids.EntityData.Leafs.Append("role", types.YLeaf{"Role", allUiids.Role})

    allUiids.EntityData.YListKeys = []string {"Uiid"}

    return &(allUiids.EntityData)
}

