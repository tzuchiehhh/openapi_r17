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

import (
	"time"
)

// UE MM Transaction Report Item per Slice
type MmTransactionSliceReportItem struct {
	Snssai *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	Timestamp    *time.Time `json:"timestamp" yaml:"timestamp" bson:"timestamp,omitempty"`
	Transactions int32      `json:"transactions" yaml:"transactions" bson:"transactions,omitempty"`
}