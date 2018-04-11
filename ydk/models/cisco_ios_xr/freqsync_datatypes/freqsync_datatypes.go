// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package freqsync_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package freqsync_datatypes"))
}

// FsyncQlOption represents Fsync ql option
type FsyncQlOption string

const (
    // ITU-T Option 1
    FsyncQlOption_option_1 FsyncQlOption = "option-1"

    // ITU-T Option 2, Generation 1
    FsyncQlOption_option_2__COMMA___generation_1 FsyncQlOption = "option-2,-generation-1"

    // ITU-T Option 2, Generation 2
    FsyncQlOption_option_2__COMMA___generation_2 FsyncQlOption = "option-2,-generation-2"
)

// FsyncClock represents Fsync clock
type FsyncClock string

const (
    // Synchronous clock
    FsyncClock_sync FsyncClock = "sync"

    // Internal clock
    FsyncClock_internal FsyncClock = "internal"
)

// FsyncQlValue represents Fsync ql value
type FsyncQlValue string

const (
    // This signal should not be used for
    // synchronization
    FsyncQlValue_dnu FsyncQlValue = "dnu"

    // ITU-T Option 1: Primary reference clock
    FsyncQlValue_o1_prc FsyncQlValue = "o1-prc"

    // ITU-T Option 1: Type I or V slave clock
    FsyncQlValue_o1_ssu_a FsyncQlValue = "o1-ssu-a"

    // ITU-T Option 1: Type IV slave clock
    FsyncQlValue_o1_ssu_b FsyncQlValue = "o1-ssu-b"

    // ITU-T Option 1: SONET equipment clock
    FsyncQlValue_o1_sec FsyncQlValue = "o1-sec"

    // ITU-T Option 2, Generation 1: Primary reference
    // source
    FsyncQlValue_o2_g1_prs FsyncQlValue = "o2-g1-prs"

    // ITU-T Option 2, Generation 1: Synchronized -
    // traceability unknown
    FsyncQlValue_o2_g1_stu FsyncQlValue = "o2-g1-stu"

    // ITU-T Option 2, Generation 1: Stratum 2
    FsyncQlValue_o2_g1_st2 FsyncQlValue = "o2-g1-st2"

    // ITU-T Option 2, Generation 1: Stratum 3
    FsyncQlValue_o2_g1_st3 FsyncQlValue = "o2-g1-st3"

    // ITU-T Option 2, Generation 1: SONET clock self
    // timed
    FsyncQlValue_o2_g1_smc FsyncQlValue = "o2-g1-smc"

    // ITU-T Option 2, Generation 1: Stratum 4 freerun
    FsyncQlValue_o2_g1_st4 FsyncQlValue = "o2-g1-st4"

    // ITU-T Option 2, Generation 2: Primary reference
    // source
    FsyncQlValue_o2_g2_prs FsyncQlValue = "o2-g2-prs"

    // ITU-T Option 2, Generation 2: Synchronized -
    // traceability unknown
    FsyncQlValue_o2_g2_stu FsyncQlValue = "o2-g2-stu"

    // ITU-T Option 2, Generation 2: Stratum 2
    FsyncQlValue_o2_g2_st2 FsyncQlValue = "o2-g2-st2"

    // ITU-T Option 2, Generation 2: Stratum 3
    FsyncQlValue_o2_g2_st3 FsyncQlValue = "o2-g2-st3"

    // ITU-T Option 2, Generation 2: Transit node
    // clock
    FsyncQlValue_o2_g2_tnc FsyncQlValue = "o2-g2-tnc"

    // ITU-T Option 2, Generation 2: Stratum 3E
    FsyncQlValue_o2_g2_st3e FsyncQlValue = "o2-g2-st3e"

    // ITU-T Option 2, Generation 2: SONET clock self
    // timed
    FsyncQlValue_o2_g2_smc FsyncQlValue = "o2-g2-smc"

    // ITU-T Option 2, Generation 2: Stratum 4 freerun
    FsyncQlValue_o2_g2_st4 FsyncQlValue = "o2-g2-st4"
)

