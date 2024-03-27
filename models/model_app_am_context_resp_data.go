/*
 * Npcf_AMPolicyAuthorization Service API
 *
 * PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.534 V17.3.0; 5G System; Access and Mobility Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.534/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// It represents a response to a modification or creation request of an Individual Application AM resource. It may contain the notification of the already met events.
type AppAmContextRespData struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi" yaml:"supi" bson:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi string `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	TermNotifUri string             `json:"termNotifUri" yaml:"termNotifUri" bson:"termNotifUri,omitempty"`
	EvSubsc      *AmEventsSubscData `json:"evSubsc,omitempty" yaml:"evSubsc" bson:"evSubsc,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat,omitempty" yaml:"suppFeat" bson:"suppFeat,omitempty"`
	// indicating a time in seconds.
	Expiry int32 `json:"expiry,omitempty" yaml:"expiry" bson:"expiry,omitempty"`
	// Indicates whether high throughput is desired for the indicated UE traffic.
	HighThruInd bool `json:"highThruInd,omitempty" yaml:"highThruInd" bson:"highThruInd,omitempty"`
	// Identifies a list of Tracking Areas per serving network where service is allowed.
	CovReq         []ServiceAreaCoverageInfo                  `json:"covReq,omitempty" yaml:"covReq" bson:"covReq,omitempty"`
	AsTimeDisParam *PcfAmPolicyControlAsTimeDistributionParam `json:"asTimeDisParam,omitempty" yaml:"asTimeDisParam" bson:"asTimeDisParam,omitempty"`
	// Contains the AM Policy Events Subscription resource identifier related to the event notification.
	AppAmContextId string                `json:"appAmContextId,omitempty" yaml:"appAmContextId" bson:"appAmContextId,omitempty"`
	RepEvents      []AmEventNotification `json:"repEvents" yaml:"repEvents" bson:"repEvents,omitempty"`
}