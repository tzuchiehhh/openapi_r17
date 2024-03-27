/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Indicates if access is allowed to a given serving network, e.g. a PLMN (MCC, MNC) or an  SNPN (MCC, MNC, NID).
type RoamingRestrictions struct {
	AccessAllowed bool `json:"accessAllowed,omitempty" yaml:"accessAllowed" bson:"accessAllowed,omitempty"`
}