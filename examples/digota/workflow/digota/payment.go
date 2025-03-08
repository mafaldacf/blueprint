package digota

type Charge struct {
	Id               string
	Statement        string
	ChargeAmount     uint64
	RefundAmount     uint64
	Refunds          []*Refund
	Currency         int32
	Email            string
	Paid             bool
	Refunded         bool
	ProviderId       int32
	ProviderChargeId string
	Created          int64
	Updated          int64
}

type Refund struct {
	RefundAmount     uint64
	ProviderRefundId string
	Reason           int32
	Created          int64
}

/* type RefundReason int32

const (
	RefundReason_GeneralError        RefundReason = 0
	RefundReason_Fraud               RefundReason = 1
	RefundReason_Duplicate           RefundReason = 2
	RefundReason_RequestedByCustomer RefundReason = 3
) */

/* var RefundReason_name = map[int32]string{
	0: "GeneralError",
	1: "Fraud",
	2: "Duplicate",
	3: "RequestedByCustomer",
}
var RefundReason_value = map[string]int32{
	"GeneralError":        0,
	"Fraud":               1,
	"Duplicate":           2,
	"RequestedByCustomer": 3,
} */

/* type PaymentProviderId int32

const (
	PaymentProviderId_PROVIDER_Reserved PaymentProviderId = 0
	PaymentProviderId_Stripe            PaymentProviderId = 1
	PaymentProviderId_Paypal            PaymentProviderId = 2
	PaymentProviderId_Braintree         PaymentProviderId = 3
) */

/* var PaymentProviderId_name = map[int32]string{
	0: "PROVIDER_Reserved",
	1: "Stripe",
	2: "Paypal",
	3: "Braintree",
}
var PaymentProviderId_value = map[string]int32{
	"PROVIDER_Reserved": 0,
	"Stripe":            1,
	"Paypal":            2,
	"Braintree":         3,
} */

/* type Currency int32

const (
	Currency_CUR_RESERVED Currency = 0
	Currency_AFN          Currency = 1
	Currency_ALL          Currency = 2
	Currency_AMD          Currency = 3
	Currency_ANG          Currency = 4
	Currency_ARS          Currency = 5
	Currency_AUD          Currency = 6
	Currency_AWG          Currency = 7
	Currency_AZN          Currency = 8
	Currency_BAM          Currency = 9
	Currency_BBD          Currency = 10
	Currency_BGN          Currency = 11
	Currency_BHD          Currency = 12
	Currency_BMD          Currency = 13
	Currency_BND          Currency = 14
	Currency_BOB          Currency = 15
	Currency_BRL          Currency = 16
	Currency_BSD          Currency = 17
	Currency_BWP          Currency = 18
	Currency_BYN          Currency = 19
	Currency_BYR          Currency = 20
	Currency_BZD          Currency = 21
	Currency_CAD          Currency = 22
	Currency_CLP          Currency = 23
	Currency_CNY          Currency = 24
	Currency_COP          Currency = 25
	Currency_CRC          Currency = 26
	Currency_CUP          Currency = 27
	Currency_CZK          Currency = 28
	Currency_DKK          Currency = 29
	Currency_DOP          Currency = 30
	Currency_DZD          Currency = 31
	Currency_EEK          Currency = 32
	Currency_EGP          Currency = 33
	Currency_EUR          Currency = 34
	Currency_FJD          Currency = 35
	Currency_FKP          Currency = 36
	Currency_GBP          Currency = 37
	Currency_GGP          Currency = 38
	Currency_GHC          Currency = 39
	Currency_GIP          Currency = 40
	Currency_GTQ          Currency = 41
	Currency_GYD          Currency = 42
	Currency_HKD          Currency = 43
	Currency_HNL          Currency = 44
	Currency_HRK          Currency = 45
	Currency_HUF          Currency = 46
	Currency_IDR          Currency = 47
	Currency_ILS          Currency = 48
	Currency_IMP          Currency = 49
	Currency_INR          Currency = 50
	Currency_IQD          Currency = 51
	Currency_IRR          Currency = 52
	Currency_ISK          Currency = 53
	Currency_JEP          Currency = 54
	Currency_JMD          Currency = 55
	Currency_JOD          Currency = 56
	Currency_JPY          Currency = 57
	Currency_KES          Currency = 58
	Currency_KGS          Currency = 59
	Currency_KHR          Currency = 60
	Currency_KPW          Currency = 61
	Currency_KRW          Currency = 62
	Currency_KWD          Currency = 63
	Currency_KYD          Currency = 64
	Currency_KZT          Currency = 65
	Currency_LAK          Currency = 66
	Currency_LBP          Currency = 67
	Currency_LKR          Currency = 68
	Currency_LRD          Currency = 69
	Currency_LTL          Currency = 70
	Currency_LVL          Currency = 71
	Currency_LYD          Currency = 72
	Currency_MAD          Currency = 73
	Currency_MKD          Currency = 74
	Currency_MNT          Currency = 75
	Currency_MUR          Currency = 76
	Currency_MXN          Currency = 77
	Currency_MWK          Currency = 78
	Currency_MYR          Currency = 79
	Currency_MZN          Currency = 80
	Currency_NAD          Currency = 81
	Currency_NGN          Currency = 82
	Currency_NIO          Currency = 83
	Currency_NOK          Currency = 84
	Currency_NPR          Currency = 85
	Currency_NZD          Currency = 86
	Currency_OMR          Currency = 87
	Currency_PAB          Currency = 88
	Currency_PEN          Currency = 89
	Currency_PHP          Currency = 90
	Currency_PKR          Currency = 91
	Currency_PLN          Currency = 92
	Currency_PYG          Currency = 93
	Currency_QAR          Currency = 94
	Currency_RON          Currency = 95
	Currency_RSD          Currency = 96
	Currency_RUB          Currency = 97
	Currency_RUR          Currency = 98
	Currency_SAR          Currency = 99
	Currency_SBD          Currency = 100
	Currency_SCR          Currency = 101
	Currency_SEK          Currency = 102
	Currency_SGD          Currency = 103
	Currency_SHP          Currency = 104
	Currency_SOS          Currency = 105
	Currency_SRD          Currency = 106
	Currency_SVC          Currency = 107
	Currency_SYP          Currency = 108
	Currency_THB          Currency = 109
	Currency_TND          Currency = 110
	Currency_TRL          Currency = 111
	Currency_TRY          Currency = 112
	Currency_TTD          Currency = 113
	Currency_TWD          Currency = 114
	Currency_TZS          Currency = 115
	Currency_UAH          Currency = 116
	Currency_UGX          Currency = 117
	Currency_AED          Currency = 118
	Currency_UYU          Currency = 119
	Currency_UZS          Currency = 120
	Currency_VEF          Currency = 121
	Currency_VND          Currency = 122
	Currency_XCD          Currency = 123
	Currency_YER          Currency = 124
	Currency_ZAR          Currency = 125
	Currency_ZMW          Currency = 126
	Currency_ZWD          Currency = 127
	Currency_USD          Currency = 128
) */

var Currency_name = map[int32]string{
	0:   "CUR_RESERVED",
	1:   "AFN",
	2:   "ALL",
	3:   "AMD",
	4:   "ANG",
	5:   "ARS",
	6:   "AUD",
	7:   "AWG",
	8:   "AZN",
	9:   "BAM",
	10:  "BBD",
	11:  "BGN",
	12:  "BHD",
	13:  "BMD",
	14:  "BND",
	15:  "BOB",
	16:  "BRL",
	17:  "BSD",
	18:  "BWP",
	19:  "BYN",
	20:  "BYR",
	21:  "BZD",
	22:  "CAD",
	23:  "CLP",
	24:  "CNY",
	25:  "COP",
	26:  "CRC",
	27:  "CUP",
	28:  "CZK",
	29:  "DKK",
	30:  "DOP",
	31:  "DZD",
	32:  "EEK",
	33:  "EGP",
	34:  "EUR",
	35:  "FJD",
	36:  "FKP",
	37:  "GBP",
	38:  "GGP",
	39:  "GHC",
	40:  "GIP",
	41:  "GTQ",
	42:  "GYD",
	43:  "HKD",
	44:  "HNL",
	45:  "HRK",
	46:  "HUF",
	47:  "IDR",
	48:  "ILS",
	49:  "IMP",
	50:  "INR",
	51:  "IQD",
	52:  "IRR",
	53:  "ISK",
	54:  "JEP",
	55:  "JMD",
	56:  "JOD",
	57:  "JPY",
	58:  "KES",
	59:  "KGS",
	60:  "KHR",
	61:  "KPW",
	62:  "KRW",
	63:  "KWD",
	64:  "KYD",
	65:  "KZT",
	66:  "LAK",
	67:  "LBP",
	68:  "LKR",
	69:  "LRD",
	70:  "LTL",
	71:  "LVL",
	72:  "LYD",
	73:  "MAD",
	74:  "MKD",
	75:  "MNT",
	76:  "MUR",
	77:  "MXN",
	78:  "MWK",
	79:  "MYR",
	80:  "MZN",
	81:  "NAD",
	82:  "NGN",
	83:  "NIO",
	84:  "NOK",
	85:  "NPR",
	86:  "NZD",
	87:  "OMR",
	88:  "PAB",
	89:  "PEN",
	90:  "PHP",
	91:  "PKR",
	92:  "PLN",
	93:  "PYG",
	94:  "QAR",
	95:  "RON",
	96:  "RSD",
	97:  "RUB",
	98:  "RUR",
	99:  "SAR",
	100: "SBD",
	101: "SCR",
	102: "SEK",
	103: "SGD",
	104: "SHP",
	105: "SOS",
	106: "SRD",
	107: "SVC",
	108: "SYP",
	109: "THB",
	110: "TND",
	111: "TRL",
	112: "TRY",
	113: "TTD",
	114: "TWD",
	115: "TZS",
	116: "UAH",
	117: "UGX",
	118: "AED",
	119: "UYU",
	120: "UZS",
	121: "VEF",
	122: "VND",
	123: "XCD",
	124: "YER",
	125: "ZAR",
	126: "ZMW",
	127: "ZWD",
	128: "USD",
}
var Currency_value = map[string]int32{
	"CUR_RESERVED": 0,
	"AFN":          1,
	"ALL":          2,
	"AMD":          3,
	"ANG":          4,
	"ARS":          5,
	"AUD":          6,
	"AWG":          7,
	"AZN":          8,
	"BAM":          9,
	"BBD":          10,
	"BGN":          11,
	"BHD":          12,
	"BMD":          13,
	"BND":          14,
	"BOB":          15,
	"BRL":          16,
	"BSD":          17,
	"BWP":          18,
	"BYN":          19,
	"BYR":          20,
	"BZD":          21,
	"CAD":          22,
	"CLP":          23,
	"CNY":          24,
	"COP":          25,
	"CRC":          26,
	"CUP":          27,
	"CZK":          28,
	"DKK":          29,
	"DOP":          30,
	"DZD":          31,
	"EEK":          32,
	"EGP":          33,
	"EUR":          34,
	"FJD":          35,
	"FKP":          36,
	"GBP":          37,
	"GGP":          38,
	"GHC":          39,
	"GIP":          40,
	"GTQ":          41,
	"GYD":          42,
	"HKD":          43,
	"HNL":          44,
	"HRK":          45,
	"HUF":          46,
	"IDR":          47,
	"ILS":          48,
	"IMP":          49,
	"INR":          50,
	"IQD":          51,
	"IRR":          52,
	"ISK":          53,
	"JEP":          54,
	"JMD":          55,
	"JOD":          56,
	"JPY":          57,
	"KES":          58,
	"KGS":          59,
	"KHR":          60,
	"KPW":          61,
	"KRW":          62,
	"KWD":          63,
	"KYD":          64,
	"KZT":          65,
	"LAK":          66,
	"LBP":          67,
	"LKR":          68,
	"LRD":          69,
	"LTL":          70,
	"LVL":          71,
	"LYD":          72,
	"MAD":          73,
	"MKD":          74,
	"MNT":          75,
	"MUR":          76,
	"MXN":          77,
	"MWK":          78,
	"MYR":          79,
	"MZN":          80,
	"NAD":          81,
	"NGN":          82,
	"NIO":          83,
	"NOK":          84,
	"NPR":          85,
	"NZD":          86,
	"OMR":          87,
	"PAB":          88,
	"PEN":          89,
	"PHP":          90,
	"PKR":          91,
	"PLN":          92,
	"PYG":          93,
	"QAR":          94,
	"RON":          95,
	"RSD":          96,
	"RUB":          97,
	"RUR":          98,
	"SAR":          99,
	"SBD":          100,
	"SCR":          101,
	"SEK":          102,
	"SGD":          103,
	"SHP":          104,
	"SOS":          105,
	"SRD":          106,
	"SVC":          107,
	"SYP":          108,
	"THB":          109,
	"TND":          110,
	"TRL":          111,
	"TRY":          112,
	"TTD":          113,
	"TWD":          114,
	"TZS":          115,
	"UAH":          116,
	"UGX":          117,
	"AED":          118,
	"UYU":          119,
	"UZS":          120,
	"VEF":          121,
	"VND":          122,
	"XCD":          123,
	"YER":          124,
	"ZAR":          125,
	"ZMW":          126,
	"ZWD":          127,
	"USD":          128,
}

type Card struct {
	Number      string
	ExpireMonth string
	ExpireYear  string
	FirstName   string
	LastName    string
	CVC         string
	Type        int32
}

/* type CardType int32

const (
	CardType_CARD_Reserved   CardType = 0
	CardType_Mastercard      CardType = 1
	CardType_Visa            CardType = 2
	CardType_AmericanExpress CardType = 3
	CardType_JCB             CardType = 4
	CardType_Discover        CardType = 5
	CardType_DinersClub      CardType = 6
) */

/* var CardType_name = map[int32]string{
	0: "CARD_Reserved",
	1: "Mastercard",
	2: "Visa",
	3: "AmericanExpress",
	4: "JCB",
	5: "Discover",
	6: "DinersClub",
}
var CardType_value = map[string]int32{
	"CARD_Reserved":   0,
	"Mastercard":      1,
	"Visa":            2,
	"AmericanExpress": 3,
	"JCB":             4,
	"Discover":        5,
	"DinersClub":      6,
} */

type ListRequest struct {
	Page  int64
	Limit int64
	Sort  int32
}

/* type ListRequest_Sort int32

const (
	ListRequest_Natural     ListRequest_Sort = 0
	ListRequest_CreatedDesc ListRequest_Sort = 1
	ListRequest_CreatedAsc  ListRequest_Sort = 2
	ListRequest_UpdatedDesc ListRequest_Sort = 3
	ListRequest_UpdatedAsc  ListRequest_Sort = 4
) */

/* var ListRequest_Sort_name = map[int32]string{
	0: "Natural",
	1: "CreatedDesc",
	2: "CreatedAsc",
	3: "UpdatedDesc",
	4: "UpdatedAsc",
}
var ListRequest_Sort_value = map[string]int32{
	"Natural":     0,
	"CreatedDesc": 1,
	"CreatedAsc":  2,
	"UpdatedDesc": 3,
	"UpdatedAsc":  4,
} */

type ChargeList struct {
	Charges []*Charge
	Total   int32
}
