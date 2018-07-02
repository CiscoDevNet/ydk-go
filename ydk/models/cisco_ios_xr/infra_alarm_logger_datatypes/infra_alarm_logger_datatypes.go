// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_alarm_logger_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_alarm_logger_datatypes"))
}

// AlarmLoggerSeverityLevel represents Alarm logger severity level
type AlarmLoggerSeverityLevel string

const (
    // Emergency
    AlarmLoggerSeverityLevel_emergency AlarmLoggerSeverityLevel = "emergency"

    // Alert
    AlarmLoggerSeverityLevel_alert AlarmLoggerSeverityLevel = "alert"

    // Critical
    AlarmLoggerSeverityLevel_critical AlarmLoggerSeverityLevel = "critical"

    // Error
    AlarmLoggerSeverityLevel_error_ AlarmLoggerSeverityLevel = "error"

    // Warning
    AlarmLoggerSeverityLevel_warning AlarmLoggerSeverityLevel = "warning"

    // Notice
    AlarmLoggerSeverityLevel_notice AlarmLoggerSeverityLevel = "notice"

    // Informational
    AlarmLoggerSeverityLevel_informational AlarmLoggerSeverityLevel = "informational"
)

