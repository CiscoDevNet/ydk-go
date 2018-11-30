// This module defines YANG extensions that are used to translate
// SMIv2 concepts into YANG.
// 
// Copyright (c) 2012 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6643; see
// the RFC itself for full legal notices.
package yang_smiv2

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package yang_smiv2"))
}

type ObjectIdentity struct {
}

func (id ObjectIdentity) String() string {
	return "ietf-yang-smiv2:object-identity"
}

