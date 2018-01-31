// This module contains a collection of IOS-XR derived YANG data 
// types.
//     
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package types"))
}

// EncryptionType represents The type of encryption used on a password string.
type EncryptionType string

const (
    // The password string is clear text.
    EncryptionType_none EncryptionType = "none"

    // The password is encrypted to an MD5 digest.
    EncryptionType_md5 EncryptionType = "md5"

    // The password is encrypted using Cisco type 7 
    // password encryption.
    EncryptionType_proprietary EncryptionType = "proprietary"

    // The password is encrypted using Cisco type 6 
    // password encryption.
    EncryptionType_type6 EncryptionType = "type6"
)

