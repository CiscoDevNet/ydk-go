package sysadmin_fpd_infra_cli_fpd_service

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_fpd_infra_cli_fpd_service"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd-service location}", reflect.TypeOf(Location{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd-service:location", reflect.TypeOf(Location{}))
}

// Location
type Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Loc interface{}

    // The type is slice of Location_Fpd2.
    Fpd2 []Location_Fpd2
}

func (location *Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd-service"
    location.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-fpd-infra-cli-fpd-service:location" + "[loc='" + fmt.Sprintf("%v", location.Loc) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fpd2"] = types.YChild{"Fpd2", nil}
    for i := range location.Fpd2 {
        location.EntityData.Children[types.GetSegmentPath(&location.Fpd2[i])] = types.YChild{"Fpd2", &location.Fpd2[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["loc"] = types.YLeaf{"Loc", location.Loc}
    return &(location.EntityData)
}

// Location_Fpd2
type Location_Fpd2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}
}

func (fpd2 *Location_Fpd2) GetEntityData() *types.CommonEntityData {
    fpd2.EntityData.YFilter = fpd2.YFilter
    fpd2.EntityData.YangName = "fpd2"
    fpd2.EntityData.BundleName = "cisco_ios_xr"
    fpd2.EntityData.ParentYangName = "location"
    fpd2.EntityData.SegmentPath = "fpd2" + "[name='" + fmt.Sprintf("%v", fpd2.Name) + "']"
    fpd2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd2.EntityData.Children = make(map[string]types.YChild)
    fpd2.EntityData.Leafs = make(map[string]types.YLeaf)
    fpd2.EntityData.Leafs["name"] = types.YLeaf{"Name", fpd2.Name}
    return &(fpd2.EntityData)
}

