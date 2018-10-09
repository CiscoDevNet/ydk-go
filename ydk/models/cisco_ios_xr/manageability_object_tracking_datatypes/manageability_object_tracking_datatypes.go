// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package manageability_object_tracking_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_object_tracking_datatypes"))
}

// ObjectTrackingBooleanSign represents Object tracking boolean sign
type ObjectTrackingBooleanSign string

const (
    // Object without not
    ObjectTrackingBooleanSign_without_not ObjectTrackingBooleanSign = "without-not"

    // Object with not
    ObjectTrackingBooleanSign_with_not ObjectTrackingBooleanSign = "with-not"
)

