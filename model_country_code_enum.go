/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// CountryCodeEnum the model 'CountryCodeEnum'
type CountryCodeEnum string

// List of CountryCodeEnum
const (
	COUNTRYCODEENUM_AF CountryCodeEnum = "AF"
	COUNTRYCODEENUM_AX CountryCodeEnum = "AX"
	COUNTRYCODEENUM_AL CountryCodeEnum = "AL"
	COUNTRYCODEENUM_DZ CountryCodeEnum = "DZ"
	COUNTRYCODEENUM_AS CountryCodeEnum = "AS"
	COUNTRYCODEENUM_AD CountryCodeEnum = "AD"
	COUNTRYCODEENUM_AO CountryCodeEnum = "AO"
	COUNTRYCODEENUM_AI CountryCodeEnum = "AI"
	COUNTRYCODEENUM_AQ CountryCodeEnum = "AQ"
	COUNTRYCODEENUM_AG CountryCodeEnum = "AG"
	COUNTRYCODEENUM_AR CountryCodeEnum = "AR"
	COUNTRYCODEENUM_AM CountryCodeEnum = "AM"
	COUNTRYCODEENUM_AW CountryCodeEnum = "AW"
	COUNTRYCODEENUM_AU CountryCodeEnum = "AU"
	COUNTRYCODEENUM_AT CountryCodeEnum = "AT"
	COUNTRYCODEENUM_AZ CountryCodeEnum = "AZ"
	COUNTRYCODEENUM_BS CountryCodeEnum = "BS"
	COUNTRYCODEENUM_BH CountryCodeEnum = "BH"
	COUNTRYCODEENUM_BD CountryCodeEnum = "BD"
	COUNTRYCODEENUM_BB CountryCodeEnum = "BB"
	COUNTRYCODEENUM_BY CountryCodeEnum = "BY"
	COUNTRYCODEENUM_BE CountryCodeEnum = "BE"
	COUNTRYCODEENUM_BZ CountryCodeEnum = "BZ"
	COUNTRYCODEENUM_BJ CountryCodeEnum = "BJ"
	COUNTRYCODEENUM_BM CountryCodeEnum = "BM"
	COUNTRYCODEENUM_BT CountryCodeEnum = "BT"
	COUNTRYCODEENUM_BO CountryCodeEnum = "BO"
	COUNTRYCODEENUM_BQ CountryCodeEnum = "BQ"
	COUNTRYCODEENUM_BA CountryCodeEnum = "BA"
	COUNTRYCODEENUM_BW CountryCodeEnum = "BW"
	COUNTRYCODEENUM_BV CountryCodeEnum = "BV"
	COUNTRYCODEENUM_BR CountryCodeEnum = "BR"
	COUNTRYCODEENUM_IO CountryCodeEnum = "IO"
	COUNTRYCODEENUM_BN CountryCodeEnum = "BN"
	COUNTRYCODEENUM_BG CountryCodeEnum = "BG"
	COUNTRYCODEENUM_BF CountryCodeEnum = "BF"
	COUNTRYCODEENUM_BI CountryCodeEnum = "BI"
	COUNTRYCODEENUM_CV CountryCodeEnum = "CV"
	COUNTRYCODEENUM_KH CountryCodeEnum = "KH"
	COUNTRYCODEENUM_CM CountryCodeEnum = "CM"
	COUNTRYCODEENUM_CA CountryCodeEnum = "CA"
	COUNTRYCODEENUM_KY CountryCodeEnum = "KY"
	COUNTRYCODEENUM_CF CountryCodeEnum = "CF"
	COUNTRYCODEENUM_TD CountryCodeEnum = "TD"
	COUNTRYCODEENUM_CL CountryCodeEnum = "CL"
	COUNTRYCODEENUM_CN CountryCodeEnum = "CN"
	COUNTRYCODEENUM_CX CountryCodeEnum = "CX"
	COUNTRYCODEENUM_CC CountryCodeEnum = "CC"
	COUNTRYCODEENUM_CO CountryCodeEnum = "CO"
	COUNTRYCODEENUM_KM CountryCodeEnum = "KM"
	COUNTRYCODEENUM_CG CountryCodeEnum = "CG"
	COUNTRYCODEENUM_CD CountryCodeEnum = "CD"
	COUNTRYCODEENUM_CK CountryCodeEnum = "CK"
	COUNTRYCODEENUM_CR CountryCodeEnum = "CR"
	COUNTRYCODEENUM_CI CountryCodeEnum = "CI"
	COUNTRYCODEENUM_HR CountryCodeEnum = "HR"
	COUNTRYCODEENUM_CU CountryCodeEnum = "CU"
	COUNTRYCODEENUM_CW CountryCodeEnum = "CW"
	COUNTRYCODEENUM_CY CountryCodeEnum = "CY"
	COUNTRYCODEENUM_CZ CountryCodeEnum = "CZ"
	COUNTRYCODEENUM_DK CountryCodeEnum = "DK"
	COUNTRYCODEENUM_DJ CountryCodeEnum = "DJ"
	COUNTRYCODEENUM_DM CountryCodeEnum = "DM"
	COUNTRYCODEENUM_DO CountryCodeEnum = "DO"
	COUNTRYCODEENUM_EC CountryCodeEnum = "EC"
	COUNTRYCODEENUM_EG CountryCodeEnum = "EG"
	COUNTRYCODEENUM_SV CountryCodeEnum = "SV"
	COUNTRYCODEENUM_GQ CountryCodeEnum = "GQ"
	COUNTRYCODEENUM_ER CountryCodeEnum = "ER"
	COUNTRYCODEENUM_EE CountryCodeEnum = "EE"
	COUNTRYCODEENUM_SZ CountryCodeEnum = "SZ"
	COUNTRYCODEENUM_ET CountryCodeEnum = "ET"
	COUNTRYCODEENUM_FK CountryCodeEnum = "FK"
	COUNTRYCODEENUM_FO CountryCodeEnum = "FO"
	COUNTRYCODEENUM_FJ CountryCodeEnum = "FJ"
	COUNTRYCODEENUM_FI CountryCodeEnum = "FI"
	COUNTRYCODEENUM_FR CountryCodeEnum = "FR"
	COUNTRYCODEENUM_GF CountryCodeEnum = "GF"
	COUNTRYCODEENUM_PF CountryCodeEnum = "PF"
	COUNTRYCODEENUM_TF CountryCodeEnum = "TF"
	COUNTRYCODEENUM_GA CountryCodeEnum = "GA"
	COUNTRYCODEENUM_GM CountryCodeEnum = "GM"
	COUNTRYCODEENUM_GE CountryCodeEnum = "GE"
	COUNTRYCODEENUM_DE CountryCodeEnum = "DE"
	COUNTRYCODEENUM_GH CountryCodeEnum = "GH"
	COUNTRYCODEENUM_GI CountryCodeEnum = "GI"
	COUNTRYCODEENUM_GR CountryCodeEnum = "GR"
	COUNTRYCODEENUM_GL CountryCodeEnum = "GL"
	COUNTRYCODEENUM_GD CountryCodeEnum = "GD"
	COUNTRYCODEENUM_GP CountryCodeEnum = "GP"
	COUNTRYCODEENUM_GU CountryCodeEnum = "GU"
	COUNTRYCODEENUM_GT CountryCodeEnum = "GT"
	COUNTRYCODEENUM_GG CountryCodeEnum = "GG"
	COUNTRYCODEENUM_GN CountryCodeEnum = "GN"
	COUNTRYCODEENUM_GW CountryCodeEnum = "GW"
	COUNTRYCODEENUM_GY CountryCodeEnum = "GY"
	COUNTRYCODEENUM_HT CountryCodeEnum = "HT"
	COUNTRYCODEENUM_HM CountryCodeEnum = "HM"
	COUNTRYCODEENUM_VA CountryCodeEnum = "VA"
	COUNTRYCODEENUM_HN CountryCodeEnum = "HN"
	COUNTRYCODEENUM_HK CountryCodeEnum = "HK"
	COUNTRYCODEENUM_HU CountryCodeEnum = "HU"
	COUNTRYCODEENUM_IS CountryCodeEnum = "IS"
	COUNTRYCODEENUM_IN CountryCodeEnum = "IN"
	COUNTRYCODEENUM_ID CountryCodeEnum = "ID"
	COUNTRYCODEENUM_IR CountryCodeEnum = "IR"
	COUNTRYCODEENUM_IQ CountryCodeEnum = "IQ"
	COUNTRYCODEENUM_IE CountryCodeEnum = "IE"
	COUNTRYCODEENUM_IM CountryCodeEnum = "IM"
	COUNTRYCODEENUM_IL CountryCodeEnum = "IL"
	COUNTRYCODEENUM_IT CountryCodeEnum = "IT"
	COUNTRYCODEENUM_JM CountryCodeEnum = "JM"
	COUNTRYCODEENUM_JP CountryCodeEnum = "JP"
	COUNTRYCODEENUM_JE CountryCodeEnum = "JE"
	COUNTRYCODEENUM_JO CountryCodeEnum = "JO"
	COUNTRYCODEENUM_KZ CountryCodeEnum = "KZ"
	COUNTRYCODEENUM_KE CountryCodeEnum = "KE"
	COUNTRYCODEENUM_KI CountryCodeEnum = "KI"
	COUNTRYCODEENUM_KW CountryCodeEnum = "KW"
	COUNTRYCODEENUM_KG CountryCodeEnum = "KG"
	COUNTRYCODEENUM_LA CountryCodeEnum = "LA"
	COUNTRYCODEENUM_LV CountryCodeEnum = "LV"
	COUNTRYCODEENUM_LB CountryCodeEnum = "LB"
	COUNTRYCODEENUM_LS CountryCodeEnum = "LS"
	COUNTRYCODEENUM_LR CountryCodeEnum = "LR"
	COUNTRYCODEENUM_LY CountryCodeEnum = "LY"
	COUNTRYCODEENUM_LI CountryCodeEnum = "LI"
	COUNTRYCODEENUM_LT CountryCodeEnum = "LT"
	COUNTRYCODEENUM_LU CountryCodeEnum = "LU"
	COUNTRYCODEENUM_MO CountryCodeEnum = "MO"
	COUNTRYCODEENUM_MG CountryCodeEnum = "MG"
	COUNTRYCODEENUM_MW CountryCodeEnum = "MW"
	COUNTRYCODEENUM_MY CountryCodeEnum = "MY"
	COUNTRYCODEENUM_MV CountryCodeEnum = "MV"
	COUNTRYCODEENUM_ML CountryCodeEnum = "ML"
	COUNTRYCODEENUM_MT CountryCodeEnum = "MT"
	COUNTRYCODEENUM_MH CountryCodeEnum = "MH"
	COUNTRYCODEENUM_MQ CountryCodeEnum = "MQ"
	COUNTRYCODEENUM_MR CountryCodeEnum = "MR"
	COUNTRYCODEENUM_MU CountryCodeEnum = "MU"
	COUNTRYCODEENUM_YT CountryCodeEnum = "YT"
	COUNTRYCODEENUM_MX CountryCodeEnum = "MX"
	COUNTRYCODEENUM_FM CountryCodeEnum = "FM"
	COUNTRYCODEENUM_MD CountryCodeEnum = "MD"
	COUNTRYCODEENUM_MC CountryCodeEnum = "MC"
	COUNTRYCODEENUM_MN CountryCodeEnum = "MN"
	COUNTRYCODEENUM_ME CountryCodeEnum = "ME"
	COUNTRYCODEENUM_MS CountryCodeEnum = "MS"
	COUNTRYCODEENUM_MA CountryCodeEnum = "MA"
	COUNTRYCODEENUM_MZ CountryCodeEnum = "MZ"
	COUNTRYCODEENUM_MM CountryCodeEnum = "MM"
	COUNTRYCODEENUM_NA CountryCodeEnum = "NA"
	COUNTRYCODEENUM_NR CountryCodeEnum = "NR"
	COUNTRYCODEENUM_NP CountryCodeEnum = "NP"
	COUNTRYCODEENUM_NL CountryCodeEnum = "NL"
	COUNTRYCODEENUM_NC CountryCodeEnum = "NC"
	COUNTRYCODEENUM_NZ CountryCodeEnum = "NZ"
	COUNTRYCODEENUM_NI CountryCodeEnum = "NI"
	COUNTRYCODEENUM_NE CountryCodeEnum = "NE"
	COUNTRYCODEENUM_NG CountryCodeEnum = "NG"
	COUNTRYCODEENUM_NU CountryCodeEnum = "NU"
	COUNTRYCODEENUM_NF CountryCodeEnum = "NF"
	COUNTRYCODEENUM_KP CountryCodeEnum = "KP"
	COUNTRYCODEENUM_MK CountryCodeEnum = "MK"
	COUNTRYCODEENUM_MP CountryCodeEnum = "MP"
	COUNTRYCODEENUM_NO CountryCodeEnum = "NO"
	COUNTRYCODEENUM_OM CountryCodeEnum = "OM"
	COUNTRYCODEENUM_PK CountryCodeEnum = "PK"
	COUNTRYCODEENUM_PW CountryCodeEnum = "PW"
	COUNTRYCODEENUM_PS CountryCodeEnum = "PS"
	COUNTRYCODEENUM_PA CountryCodeEnum = "PA"
	COUNTRYCODEENUM_PG CountryCodeEnum = "PG"
	COUNTRYCODEENUM_PY CountryCodeEnum = "PY"
	COUNTRYCODEENUM_PE CountryCodeEnum = "PE"
	COUNTRYCODEENUM_PH CountryCodeEnum = "PH"
	COUNTRYCODEENUM_PN CountryCodeEnum = "PN"
	COUNTRYCODEENUM_PL CountryCodeEnum = "PL"
	COUNTRYCODEENUM_PT CountryCodeEnum = "PT"
	COUNTRYCODEENUM_PR CountryCodeEnum = "PR"
	COUNTRYCODEENUM_QA CountryCodeEnum = "QA"
	COUNTRYCODEENUM_RE CountryCodeEnum = "RE"
	COUNTRYCODEENUM_RO CountryCodeEnum = "RO"
	COUNTRYCODEENUM_RU CountryCodeEnum = "RU"
	COUNTRYCODEENUM_RW CountryCodeEnum = "RW"
	COUNTRYCODEENUM_BL CountryCodeEnum = "BL"
	COUNTRYCODEENUM_SH CountryCodeEnum = "SH"
	COUNTRYCODEENUM_KN CountryCodeEnum = "KN"
	COUNTRYCODEENUM_LC CountryCodeEnum = "LC"
	COUNTRYCODEENUM_MF CountryCodeEnum = "MF"
	COUNTRYCODEENUM_PM CountryCodeEnum = "PM"
	COUNTRYCODEENUM_VC CountryCodeEnum = "VC"
	COUNTRYCODEENUM_WS CountryCodeEnum = "WS"
	COUNTRYCODEENUM_SM CountryCodeEnum = "SM"
	COUNTRYCODEENUM_ST CountryCodeEnum = "ST"
	COUNTRYCODEENUM_SA CountryCodeEnum = "SA"
	COUNTRYCODEENUM_SN CountryCodeEnum = "SN"
	COUNTRYCODEENUM_RS CountryCodeEnum = "RS"
	COUNTRYCODEENUM_SC CountryCodeEnum = "SC"
	COUNTRYCODEENUM_SL CountryCodeEnum = "SL"
	COUNTRYCODEENUM_SG CountryCodeEnum = "SG"
	COUNTRYCODEENUM_SX CountryCodeEnum = "SX"
	COUNTRYCODEENUM_SK CountryCodeEnum = "SK"
	COUNTRYCODEENUM_SI CountryCodeEnum = "SI"
	COUNTRYCODEENUM_SB CountryCodeEnum = "SB"
	COUNTRYCODEENUM_SO CountryCodeEnum = "SO"
	COUNTRYCODEENUM_ZA CountryCodeEnum = "ZA"
	COUNTRYCODEENUM_GS CountryCodeEnum = "GS"
	COUNTRYCODEENUM_KR CountryCodeEnum = "KR"
	COUNTRYCODEENUM_SS CountryCodeEnum = "SS"
	COUNTRYCODEENUM_ES CountryCodeEnum = "ES"
	COUNTRYCODEENUM_LK CountryCodeEnum = "LK"
	COUNTRYCODEENUM_SD CountryCodeEnum = "SD"
	COUNTRYCODEENUM_SR CountryCodeEnum = "SR"
	COUNTRYCODEENUM_SJ CountryCodeEnum = "SJ"
	COUNTRYCODEENUM_SE CountryCodeEnum = "SE"
	COUNTRYCODEENUM_CH CountryCodeEnum = "CH"
	COUNTRYCODEENUM_SY CountryCodeEnum = "SY"
	COUNTRYCODEENUM_TW CountryCodeEnum = "TW"
	COUNTRYCODEENUM_TJ CountryCodeEnum = "TJ"
	COUNTRYCODEENUM_TZ CountryCodeEnum = "TZ"
	COUNTRYCODEENUM_TH CountryCodeEnum = "TH"
	COUNTRYCODEENUM_TL CountryCodeEnum = "TL"
	COUNTRYCODEENUM_TG CountryCodeEnum = "TG"
	COUNTRYCODEENUM_TK CountryCodeEnum = "TK"
	COUNTRYCODEENUM_TO CountryCodeEnum = "TO"
	COUNTRYCODEENUM_TT CountryCodeEnum = "TT"
	COUNTRYCODEENUM_TN CountryCodeEnum = "TN"
	COUNTRYCODEENUM_TR CountryCodeEnum = "TR"
	COUNTRYCODEENUM_TM CountryCodeEnum = "TM"
	COUNTRYCODEENUM_TC CountryCodeEnum = "TC"
	COUNTRYCODEENUM_TV CountryCodeEnum = "TV"
	COUNTRYCODEENUM_UG CountryCodeEnum = "UG"
	COUNTRYCODEENUM_UA CountryCodeEnum = "UA"
	COUNTRYCODEENUM_AE CountryCodeEnum = "AE"
	COUNTRYCODEENUM_GB CountryCodeEnum = "GB"
	COUNTRYCODEENUM_UM CountryCodeEnum = "UM"
	COUNTRYCODEENUM_US CountryCodeEnum = "US"
	COUNTRYCODEENUM_UY CountryCodeEnum = "UY"
	COUNTRYCODEENUM_UZ CountryCodeEnum = "UZ"
	COUNTRYCODEENUM_VU CountryCodeEnum = "VU"
	COUNTRYCODEENUM_VE CountryCodeEnum = "VE"
	COUNTRYCODEENUM_VN CountryCodeEnum = "VN"
	COUNTRYCODEENUM_VG CountryCodeEnum = "VG"
	COUNTRYCODEENUM_VI CountryCodeEnum = "VI"
	COUNTRYCODEENUM_WF CountryCodeEnum = "WF"
	COUNTRYCODEENUM_EH CountryCodeEnum = "EH"
	COUNTRYCODEENUM_YE CountryCodeEnum = "YE"
	COUNTRYCODEENUM_ZM CountryCodeEnum = "ZM"
	COUNTRYCODEENUM_ZW CountryCodeEnum = "ZW"
)

// All allowed values of CountryCodeEnum enum
var AllowedCountryCodeEnumEnumValues = []CountryCodeEnum{
	"AF",
	"AX",
	"AL",
	"DZ",
	"AS",
	"AD",
	"AO",
	"AI",
	"AQ",
	"AG",
	"AR",
	"AM",
	"AW",
	"AU",
	"AT",
	"AZ",
	"BS",
	"BH",
	"BD",
	"BB",
	"BY",
	"BE",
	"BZ",
	"BJ",
	"BM",
	"BT",
	"BO",
	"BQ",
	"BA",
	"BW",
	"BV",
	"BR",
	"IO",
	"BN",
	"BG",
	"BF",
	"BI",
	"CV",
	"KH",
	"CM",
	"CA",
	"KY",
	"CF",
	"TD",
	"CL",
	"CN",
	"CX",
	"CC",
	"CO",
	"KM",
	"CG",
	"CD",
	"CK",
	"CR",
	"CI",
	"HR",
	"CU",
	"CW",
	"CY",
	"CZ",
	"DK",
	"DJ",
	"DM",
	"DO",
	"EC",
	"EG",
	"SV",
	"GQ",
	"ER",
	"EE",
	"SZ",
	"ET",
	"FK",
	"FO",
	"FJ",
	"FI",
	"FR",
	"GF",
	"PF",
	"TF",
	"GA",
	"GM",
	"GE",
	"DE",
	"GH",
	"GI",
	"GR",
	"GL",
	"GD",
	"GP",
	"GU",
	"GT",
	"GG",
	"GN",
	"GW",
	"GY",
	"HT",
	"HM",
	"VA",
	"HN",
	"HK",
	"HU",
	"IS",
	"IN",
	"ID",
	"IR",
	"IQ",
	"IE",
	"IM",
	"IL",
	"IT",
	"JM",
	"JP",
	"JE",
	"JO",
	"KZ",
	"KE",
	"KI",
	"KW",
	"KG",
	"LA",
	"LV",
	"LB",
	"LS",
	"LR",
	"LY",
	"LI",
	"LT",
	"LU",
	"MO",
	"MG",
	"MW",
	"MY",
	"MV",
	"ML",
	"MT",
	"MH",
	"MQ",
	"MR",
	"MU",
	"YT",
	"MX",
	"FM",
	"MD",
	"MC",
	"MN",
	"ME",
	"MS",
	"MA",
	"MZ",
	"MM",
	"NA",
	"NR",
	"NP",
	"NL",
	"NC",
	"NZ",
	"NI",
	"NE",
	"NG",
	"NU",
	"NF",
	"KP",
	"MK",
	"MP",
	"NO",
	"OM",
	"PK",
	"PW",
	"PS",
	"PA",
	"PG",
	"PY",
	"PE",
	"PH",
	"PN",
	"PL",
	"PT",
	"PR",
	"QA",
	"RE",
	"RO",
	"RU",
	"RW",
	"BL",
	"SH",
	"KN",
	"LC",
	"MF",
	"PM",
	"VC",
	"WS",
	"SM",
	"ST",
	"SA",
	"SN",
	"RS",
	"SC",
	"SL",
	"SG",
	"SX",
	"SK",
	"SI",
	"SB",
	"SO",
	"ZA",
	"GS",
	"KR",
	"SS",
	"ES",
	"LK",
	"SD",
	"SR",
	"SJ",
	"SE",
	"CH",
	"SY",
	"TW",
	"TJ",
	"TZ",
	"TH",
	"TL",
	"TG",
	"TK",
	"TO",
	"TT",
	"TN",
	"TR",
	"TM",
	"TC",
	"TV",
	"UG",
	"UA",
	"AE",
	"GB",
	"UM",
	"US",
	"UY",
	"UZ",
	"VU",
	"VE",
	"VN",
	"VG",
	"VI",
	"WF",
	"EH",
	"YE",
	"ZM",
	"ZW",
}

func (v *CountryCodeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountryCodeEnum(value)
	for _, existing := range AllowedCountryCodeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountryCodeEnum", value)
}

// NewCountryCodeEnumFromValue returns a pointer to a valid CountryCodeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountryCodeEnumFromValue(v string) (*CountryCodeEnum, error) {
	ev := CountryCodeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountryCodeEnum: valid values are %v", v, AllowedCountryCodeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountryCodeEnum) IsValid() bool {
	for _, existing := range AllowedCountryCodeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CountryCodeEnum value
func (v CountryCodeEnum) Ptr() *CountryCodeEnum {
	return &v
}

type NullableCountryCodeEnum struct {
	value *CountryCodeEnum
	isSet bool
}

func (v NullableCountryCodeEnum) Get() *CountryCodeEnum {
	return v.value
}

func (v *NullableCountryCodeEnum) Set(val *CountryCodeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryCodeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryCodeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryCodeEnum(val *CountryCodeEnum) *NullableCountryCodeEnum {
	return &NullableCountryCodeEnum{value: val, isSet: true}
}

func (v NullableCountryCodeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryCodeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
