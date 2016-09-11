package nl_BQ

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type nl_BQ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositivePrefix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'nl_BQ' locale
func New() locales.Translator {
	return &nl_BQ{
		locale:                 "nl_BQ",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6, 2},
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "$", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mrt.", "apr.", "mei", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		daysNarrow:             []string{"Z", "M", "D", "W", "D", "V", "Z"},
		daysShort:              []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		daysWide:               []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		periodsNarrow:          []string{"a.m.", "p.m."},
		periodsWide:            []string{"a.m.", "p.m."},
		erasAbbreviated:        []string{"v.Chr.", "n.Chr."},
		erasNarrow:             []string{"v.C.", "n.C."},
		erasWide:               []string{"voor Christus", "na Christus"},
		timezones:              map[string]string{"PST": "Pacific-standaardtijd", "ECT": "Ecuadoraanse tijd", "AST": "Atlantic-standaardtijd", "HKST": "Hongkongse zomertijd", "HAST": "Hawaii-Aleoetische standaardtijd", "ADT": "Atlantic-zomertijd", "HNT": "Newfoundland-standaardtijd", "ART": "Argentijnse standaardtijd", "UYST": "Uruguayaanse zomertijd", "SRT": "Surinaamse tijd", "CLT": "Chileense standaardtijd", "CLST": "Chileense zomertijd", "EAT": "Oost-Afrikaanse tijd", "MEZ": "Midden-Europese standaardtijd", "GMT": "Greenwich Mean Time", "TMST": "Turkmeense zomertijd", "IST": "Indiase tijd", "JST": "Japanse standaardtijd", "JDT": "Japanse zomertijd", "EST": "Eastern-standaardtijd", "ACST": "Midden-Australische standaardtijd", "CHAST": "Chatham-standaardtijd", "BT": "Bhutaanse tijd", "CDT": "Central-zomertijd", "COST": "Colombiaanse zomertijd", "CST": "Central-standaardtijd", "AKDT": "Alaska-zomertijd", "MESZ": "Midden-Europese zomertijd", "ACWDT": "Midden-Australische westelijke zomertijd", "NZST": "Nieuw-Zeelandse standaardtijd", "OEZ": "Oost-Europese standaardtijd", "PDT": "Pacific-zomertijd", "AEST": "Oost-Australische standaardtijd", "AEDT": "Oost-Australische zomertijd", "LHDT": "Lord Howe-eilandse zomertijd", "CAT": "Centraal-Afrikaanse tijd", "WESZ": "West-Europese zomertijd", "GYT": "Guyaanse tijd", "AKST": "Alaska-standaardtijd", "TMT": "Turkmeense standaardtijd", "OESZ": "Oost-Europese zomertijd", "GFT": "Frans-Guyaanse tijd", "WEZ": "West-Europese standaardtijd", "MYT": "Maleisische tijd", "SAST": "Zuid-Afrikaanse tijd", "WARST": "West-Argentijnse zomertijd", "WIB": "West-Indonesische tijd", "HADT": "Hawaii-Aleoetische zomertijd", "LHST": "Lord Howe-eilandse standaardtijd", "∅∅∅": "Azoren-zomertijd", "ACDT": "Midden-Australische zomertijd", "EDT": "Eastern-zomertijd", "HKT": "Hongkongse standaardtijd", "ChST": "Chamorro-tijd", "BOT": "Boliviaanse tijd", "MDT": "Mountain-zomertijd", "SGT": "Singaporese standaardtijd", "ACWST": "Midden-Australische westelijke standaardtijd", "ARST": "Argentijnse zomertijd", "AWST": "West-Australische standaardtijd", "WIT": "Oost-Indonesische tijd", "MST": "Mountain-standaardtijd", "VET": "Venezolaanse tijd", "CHADT": "Chatham-zomertijd", "HAT": "Newfoundland-zomertijd", "WITA": "Centraal-Indonesische tijd", "WAT": "West-Afrikaanse standaardtijd", "WAST": "West-Afrikaanse zomertijd", "COT": "Colombiaanse standaardtijd", "WART": "West-Argentijnse standaardtijd", "AWDT": "West-Australische zomertijd", "NZDT": "Nieuw-Zeelandse zomertijd", "UYT": "Uruguayaanse standaardtijd"},
	}
}

// Locale returns the current translators string locale
func (nl *nl_BQ) Locale() string {
	return nl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nl_BQ'
func (nl *nl_BQ) PluralsCardinal() []locales.PluralRule {
	return nl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nl_BQ'
func (nl *nl_BQ) PluralsOrdinal() []locales.PluralRule {
	return nl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nl_BQ'
func (nl *nl_BQ) PluralsRange() []locales.PluralRule {
	return nl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nl_BQ'
func (nl *nl_BQ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nl_BQ'
func (nl *nl_BQ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nl_BQ'
func (nl *nl_BQ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := nl.CardinalPluralRule(num1, v1)
	end := nl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nl *nl_BQ) MonthAbbreviated(month time.Month) string {
	return nl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nl *nl_BQ) MonthsAbbreviated() []string {
	return nl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nl *nl_BQ) MonthNarrow(month time.Month) string {
	return nl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nl *nl_BQ) MonthsNarrow() []string {
	return nl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nl *nl_BQ) MonthWide(month time.Month) string {
	return nl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nl *nl_BQ) MonthsWide() []string {
	return nl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nl *nl_BQ) WeekdayAbbreviated(weekday time.Weekday) string {
	return nl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nl *nl_BQ) WeekdaysAbbreviated() []string {
	return nl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nl *nl_BQ) WeekdayNarrow(weekday time.Weekday) string {
	return nl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nl *nl_BQ) WeekdaysNarrow() []string {
	return nl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nl *nl_BQ) WeekdayShort(weekday time.Weekday) string {
	return nl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nl *nl_BQ) WeekdaysShort() []string {
	return nl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nl *nl_BQ) WeekdayWide(weekday time.Weekday) string {
	return nl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nl *nl_BQ) WeekdaysWide() []string {
	return nl.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nl_BQ' and handles both Whole and Real numbers based on 'v'
func (nl *nl_BQ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 0 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nl_BQ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nl *nl_BQ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(nl.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nl_BQ'
func (nl *nl_BQ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 2 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(nl.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, nl.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nl_BQ'
// in accounting notation.
func (nl *nl_BQ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 0*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(nl.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nl.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(nl.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, nl.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nl.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nl_BQ'
func (nl *nl_BQ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
