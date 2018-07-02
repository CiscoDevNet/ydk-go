// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package atm_common_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_common_datatypes"))
}

// AtmVpShaping represents Atm vp shaping
type AtmVpShaping string

const (
    // Constant Bit Rate
    AtmVpShaping_cbr AtmVpShaping = "cbr"

    // Variable Bit Rate-non real time
    AtmVpShaping_vbr_nrt AtmVpShaping = "vbr-nrt"

    // Variable Bit Rate-real time
    AtmVpShaping_vbr_rt AtmVpShaping = "vbr-rt"

    // Unspecified Bit Rate
    AtmVpShaping_ubr AtmVpShaping = "ubr"
)

// AtmPvcShaping represents Atm pvc shaping
type AtmPvcShaping string

const (
    // Constant Bit Rate
    AtmPvcShaping_cbr AtmPvcShaping = "cbr"

    // Variable Bit Rate-non real time
    AtmPvcShaping_vbr_nrt AtmPvcShaping = "vbr-nrt"

    // Variable Bit Rate-real time
    AtmPvcShaping_vbr_rt AtmPvcShaping = "vbr-rt"

    // Unspecified Bit Rate
    AtmPvcShaping_ubr AtmPvcShaping = "ubr"
)

// AtmPvcEncapsulation represents Atm pvc encapsulation
type AtmPvcEncapsulation string

const (
    // SNAP
    AtmPvcEncapsulation_snap AtmPvcEncapsulation = "snap"

    // VC MUX
    AtmPvcEncapsulation_vc_mux AtmPvcEncapsulation = "vc-mux"

    // NLPID
    AtmPvcEncapsulation_nlpid AtmPvcEncapsulation = "nlpid"

    // AAL0
    AtmPvcEncapsulation_aal0 AtmPvcEncapsulation = "aal0"

    // AAL5
    AtmPvcEncapsulation_aal5 AtmPvcEncapsulation = "aal5"
)

// AtmPvcData represents Atm pvc data
type AtmPvcData string

const (
    // Data
    AtmPvcData_data AtmPvcData = "data"

    // ILMI
    AtmPvcData_ilmi AtmPvcData = "ilmi"

    // Layer2
    AtmPvcData_layer2 AtmPvcData = "layer2"
)

