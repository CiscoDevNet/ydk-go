// This module contains a collection of YANG definitions
// for Cisco IOS-XR li package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lawful-intercept: Lawful intercept configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package li_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package li_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-li-cfg lawful-intercept}", reflect.TypeOf(LawfulIntercept{}))
    ydk.RegisterEntity("Cisco-IOS-XR-li-cfg:lawful-intercept", reflect.TypeOf(LawfulIntercept{}))
}

// LawfulIntercept
// Lawful intercept configuration
type LawfulIntercept struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable lawful intercept feature. The type is interface{}.
    Disable interface{}
}

func (lawfulIntercept *LawfulIntercept) GetFilter() yfilter.YFilter { return lawfulIntercept.YFilter }

func (lawfulIntercept *LawfulIntercept) SetFilter(yf yfilter.YFilter) { lawfulIntercept.YFilter = yf }

func (lawfulIntercept *LawfulIntercept) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (lawfulIntercept *LawfulIntercept) GetSegmentPath() string {
    return "Cisco-IOS-XR-li-cfg:lawful-intercept"
}

func (lawfulIntercept *LawfulIntercept) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lawfulIntercept *LawfulIntercept) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lawfulIntercept *LawfulIntercept) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = lawfulIntercept.Disable
    return leafs
}

func (lawfulIntercept *LawfulIntercept) GetBundleName() string { return "cisco_ios_xr" }

func (lawfulIntercept *LawfulIntercept) GetYangName() string { return "lawful-intercept" }

func (lawfulIntercept *LawfulIntercept) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lawfulIntercept *LawfulIntercept) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lawfulIntercept *LawfulIntercept) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lawfulIntercept *LawfulIntercept) SetParent(parent types.Entity) { lawfulIntercept.parent = parent }

func (lawfulIntercept *LawfulIntercept) GetParent() types.Entity { return lawfulIntercept.parent }

func (lawfulIntercept *LawfulIntercept) GetParentYangName() string { return "Cisco-IOS-XR-li-cfg" }

