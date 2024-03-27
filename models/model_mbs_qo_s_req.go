/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.537 V17.3.0; 5G System; Multicast/Broadcast Policy Control Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.537/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represent MBS QoS requirements.
type MbsQoSReq struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.
	Var5qi int32 `json:"5qi" yaml:"5qi" bson:"5qi,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GuarBitRate string `json:"guarBitRate,omitempty" yaml:"guarBitRate" bson:"guarBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MaxBitRate string `json:"maxBitRate,omitempty" yaml:"maxBitRate" bson:"maxBitRate,omitempty"`
	// Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	AverWindow int32 `json:"averWindow,omitempty" yaml:"averWindow" bson:"averWindow,omitempty"`
	ReqMbsArp  *Arp  `json:"reqMbsArp,omitempty" yaml:"reqMbsArp" bson:"reqMbsArp,omitempty"`
}