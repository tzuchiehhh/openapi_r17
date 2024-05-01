/*
 * Npcf_UEPolicyControl
 *
 * UE Policy Control Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 29.525 V16.9.0; 5G System; UE Policy Control Service.
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.525/
 *
 * API version: 1.1.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UePolicyAssociationReleaseCause string

// List of UePolicyAssociationReleaseCause
const (
	UePolicyAssociationReleaseCause_UNSPECIFIED      UePolicyAssociationReleaseCause = "UNSPECIFIED"
	UePolicyAssociationReleaseCause_UE_SUBSCRIPTION  UePolicyAssociationReleaseCause = "UE_SUBSCRIPTION"
	UePolicyAssociationReleaseCause_INSUFFICIENT_RES UePolicyAssociationReleaseCause = "INSUFFICIENT_RES"
)