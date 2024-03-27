/*
 * AUSF API
 *
 * AUSF UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.509 V17.9.0; 5G System; 3GPP TS Authentication Server services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.509
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the information related to the resource generated to handle the UE authentication. It contains at least the UE id, Serving Network, the Authentication Method and related EAP information or related 5G-AKA information.
type UeAuthenticationCtx struct {
	AuthType      AusfUeAuthenticationAuthType `json:"authType" yaml:"authType" bson:"authType,omitempty"`
	Var5gAuthData interface{}                  `json:"5gAuthData" yaml:"5gAuthData" bson:"5gAuthData,omitempty"`
	// A map(list of key-value pairs) where member serves as key
	Links              map[string][]Link `json:"_links" yaml:"_links" bson:"_links,omitempty"`
	ServingNetworkName string            `json:"servingNetworkName,omitempty" yaml:"servingNetworkName" bson:"servingNetworkName,omitempty"`
}