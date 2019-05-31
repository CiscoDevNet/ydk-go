// This module contains a collection of YANG type definitions for 
// SYSLOG.
package syslog_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package syslog_types"))
}

type SyslogFacility struct {
}

func (id SyslogFacility) String() string {
	return "ietf-syslog-types:syslog-facility"
}

type Kern struct {
}

func (id Kern) String() string {
	return "ietf-syslog-types:kern"
}

type User struct {
}

func (id User) String() string {
	return "ietf-syslog-types:user"
}

type Mail struct {
}

func (id Mail) String() string {
	return "ietf-syslog-types:mail"
}

type Daemon struct {
}

func (id Daemon) String() string {
	return "ietf-syslog-types:daemon"
}

type Auth struct {
}

func (id Auth) String() string {
	return "ietf-syslog-types:auth"
}

type Syslog struct {
}

func (id Syslog) String() string {
	return "ietf-syslog-types:syslog"
}

type Lpr struct {
}

func (id Lpr) String() string {
	return "ietf-syslog-types:lpr"
}

type News struct {
}

func (id News) String() string {
	return "ietf-syslog-types:news"
}

type Uucp struct {
}

func (id Uucp) String() string {
	return "ietf-syslog-types:uucp"
}

type Cron struct {
}

func (id Cron) String() string {
	return "ietf-syslog-types:cron"
}

type Authpriv struct {
}

func (id Authpriv) String() string {
	return "ietf-syslog-types:authpriv"
}

type Ftp struct {
}

func (id Ftp) String() string {
	return "ietf-syslog-types:ftp"
}

type Ntp struct {
}

func (id Ntp) String() string {
	return "ietf-syslog-types:ntp"
}

type Audit struct {
}

func (id Audit) String() string {
	return "ietf-syslog-types:audit"
}

type Console struct {
}

func (id Console) String() string {
	return "ietf-syslog-types:console"
}

type Cron2 struct {
}

func (id Cron2) String() string {
	return "ietf-syslog-types:cron2"
}

type Local0 struct {
}

func (id Local0) String() string {
	return "ietf-syslog-types:local0"
}

type Local1 struct {
}

func (id Local1) String() string {
	return "ietf-syslog-types:local1"
}

type Local2 struct {
}

func (id Local2) String() string {
	return "ietf-syslog-types:local2"
}

type Local3 struct {
}

func (id Local3) String() string {
	return "ietf-syslog-types:local3"
}

type Local4 struct {
}

func (id Local4) String() string {
	return "ietf-syslog-types:local4"
}

type Local5 struct {
}

func (id Local5) String() string {
	return "ietf-syslog-types:local5"
}

type Local6 struct {
}

func (id Local6) String() string {
	return "ietf-syslog-types:local6"
}

type Local7 struct {
}

func (id Local7) String() string {
	return "ietf-syslog-types:local7"
}

// Severity represents The definitions for Syslog message severity as per RFC 5424.
type Severity string

const (
    // Emergency Level Msg
    Severity_emergency Severity = "emergency"

    // Alert Level Msg
    Severity_alert Severity = "alert"

    // Critical Level Msg
    Severity_critical Severity = "critical"

    // Error Level Msg
    Severity_error_ Severity = "error"

    // Warning Level Msg
    Severity_warning Severity = "warning"

    // Notification Level Msg
    Severity_notice Severity = "notice"

    // Informational Level Msg
    Severity_info Severity = "info"

    // Debugging Level Msg
    Severity_debug Severity = "debug"
)

