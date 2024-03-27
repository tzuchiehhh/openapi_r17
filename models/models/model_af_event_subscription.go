/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Describes the event information delivered in the subscription.
type AfEventSubscription struct {
	Event       PcfPolicyAuthorizationAfEvent `json:"event" yaml:"event" bson:"event,omitempty"`
	NotifMethod AfNotifMethod                 `json:"notifMethod,omitempty" yaml:"notifMethod" bson:"notifMethod,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	RepPeriod int32 `json:"repPeriod,omitempty" yaml:"repPeriod" bson:"repPeriod,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	WaitTime int32 `json:"waitTime,omitempty" yaml:"waitTime" bson:"waitTime,omitempty"`
}