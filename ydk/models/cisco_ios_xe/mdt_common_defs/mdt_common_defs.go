// This module contains a collection of common YANG 
// definitions for streaming telemetry.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package mdt_common_defs

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mdt_common_defs"))
}

// MdtSubFilterType represents Types of subscription filters.
type MdtSubFilterType string

const (
    // Indicates that no filter has been specified.
    MdtSubFilterType_sub_filter_type_none MdtSubFilterType = "sub-filter-type-none"

    // Xpath defining the data items of interest.
    // A limited set of the Xpath 1.0 expressions is
    // supported.
    MdtSubFilterType_sub_filter_type_xpath MdtSubFilterType = "sub-filter-type-xpath"

    // TDL-URI defining the data items of interest.
    MdtSubFilterType_sub_filter_type_tdl_uri MdtSubFilterType = "sub-filter-type-tdl-uri"

    // tdl-transform define transform scheme 
    // to be subscribed.
    MdtSubFilterType_sub_filter_type_tdl_transform MdtSubFilterType = "sub-filter-type-tdl-transform"
)

// MdtSubUpdateTrigger represents Types of subscription update triggers.
type MdtSubUpdateTrigger string

const (
    // Indicates trigger has not been specified.
    MdtSubUpdateTrigger_sub_upd_trig_none MdtSubUpdateTrigger = "sub-upd-trig-none"

    // Subscription is triggered on a periodic basis.
    MdtSubUpdateTrigger_sub_upd_trig_periodic MdtSubUpdateTrigger = "sub-upd-trig-periodic"

    // Subscription is triggered when a value changes.
    MdtSubUpdateTrigger_sub_upd_trig_on_change MdtSubUpdateTrigger = "sub-upd-trig-on-change"
)

