// This module defines an extension to the NETCONF protocol
// that allows the NETCONF client to control how default
// values are handled by the server in particular NETCONF
// operations.
// 
// Copyright (c) 2011 IETF Trust and the persons identified as
// the document authors.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6243; see
// the RFC itself for full legal notices.
package netconf_with_defaults

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package netconf_with_defaults"))
}

// WithDefaultsMode represents Possible modes to report default data.
type WithDefaultsMode string

const (
    // All default data is reported.
    WithDefaultsMode_report_all WithDefaultsMode = "report-all"

    // All default data is reported.
    // Any nodes considered to be default data
    // will contain a 'default' XML attribute,
    // set to 'true' or '1'.
    WithDefaultsMode_report_all_tagged WithDefaultsMode = "report-all-tagged"

    // Values are not reported if they contain the default.
    WithDefaultsMode_trim WithDefaultsMode = "trim"

    // Report values that contain the definition of
    // explicitly set data.
    WithDefaultsMode_explicit WithDefaultsMode = "explicit"
)

