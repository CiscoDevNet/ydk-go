// This module contains a collection of YANG definitions for configuring
// diffserv specification implementations. Copyright (c) 2014 IETF
// Trust and the persons identified as authors of the code. All rights
// reserved. Redistribution and use in source and binary forms, with
// or without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents (http://trustee.ietf.org/license-info).
// This version of this YANG module is part of RFC XXXX; see the
// RFC itself for full legal notices.
package diffserv_action

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_action"))
}

type Marking struct {
}

func (id Marking) String() string {
	return "ietf-diffserv-action:marking"
}

type AlwaysDrop struct {
}

func (id AlwaysDrop) String() string {
	return "ietf-diffserv-action:always-drop"
}

type TailDrop struct {
}

func (id TailDrop) String() string {
	return "ietf-diffserv-action:tail-drop"
}

type DropType struct {
}

func (id DropType) String() string {
	return "ietf-diffserv-action:drop-type"
}

type MeterActionDrop struct {
}

func (id MeterActionDrop) String() string {
	return "ietf-diffserv-action:meter-action-drop"
}

type MinRate struct {
}

func (id MinRate) String() string {
	return "ietf-diffserv-action:min-rate"
}

type Meter struct {
}

func (id Meter) String() string {
	return "ietf-diffserv-action:meter"
}

type Priority struct {
}

func (id Priority) String() string {
	return "ietf-diffserv-action:priority"
}

type MaxRate struct {
}

func (id MaxRate) String() string {
	return "ietf-diffserv-action:max-rate"
}

type MeterActionType struct {
}

func (id MeterActionType) String() string {
	return "ietf-diffserv-action:meter-action-type"
}

type MeterActionSet struct {
}

func (id MeterActionSet) String() string {
	return "ietf-diffserv-action:meter-action-set"
}

type RandomDetect struct {
}

func (id RandomDetect) String() string {
	return "ietf-diffserv-action:random-detect"
}

type AlgorithmicDrop struct {
}

func (id AlgorithmicDrop) String() string {
	return "ietf-diffserv-action:algorithmic-drop"
}

