/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicate the supported V2X Capability by the PCF.
type V2xCapability struct {
	LteV2x bool `json:"lteV2x,omitempty" yaml:"lteV2x" bson:"lteV2x,omitempty"`
	NrV2x  bool `json:"nrV2x,omitempty" yaml:"nrV2x" bson:"nrV2x,omitempty"`
}