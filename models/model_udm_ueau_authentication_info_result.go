/*
 * Nudm_UEAU
 *
 * UDM UE Authentication Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UdmUeauAuthenticationInfoResult struct {
	AuthType UdmUeauAuthType `json:"authType" yaml:"authType" bson:"authType,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures    string                `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
	AuthenticationVector *AuthenticationVector `json:"authenticationVector,omitempty" yaml:"authenticationVector" bson:"authenticationVector,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi      string                 `json:"supi,omitempty" yaml:"supi" bson:"supi,omitempty"`
	AkmaInd   bool                   `json:"akmaInd,omitempty" yaml:"akmaInd" bson:"akmaInd,omitempty"`
	AuthAaa   bool                   `json:"authAaa,omitempty" yaml:"authAaa" bson:"authAaa,omitempty"`
	RoutingId string                 `json:"routingId,omitempty" yaml:"routingId" bson:"routingId,omitempty"`
	PvsInfo   []ServerAddressingInfo `json:"pvsInfo,omitempty" yaml:"pvsInfo" bson:"pvsInfo,omitempty"`
}