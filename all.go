package currency

var (
	byNumeric = map[string]*Currency{}
	byISO     = map[string]*Currency{}

	// Testing currency
	// For tesing purposes this currency should have 2 decimal digits
	// just like other major currencies
	xts = Currency{
		Numeric:  "963",
		ISO:      "XTS",
		Decimals: 2,
		Name:     "Testing currency",
	}
)

func add(c Currency) {
	ref := &c

	byNumeric[c.Numeric] = ref
	byISO[c.ISO] = ref
}

func init() {
	Add(
		Currency{Numeric: "008", ISO: "ALL", Decimals: 2, Name: "Lek"},
		Currency{Numeric: "012", ISO: "DZD", Decimals: 2, Name: "Algerian Dinar"},
		Currency{Numeric: "032", ISO: "ARS", Decimals: 2, Name: "Argentine Peso"},
		Currency{Numeric: "036", ISO: "AUD", Decimals: 2, Name: "Australian Dollar"},
		Currency{Numeric: "044", ISO: "BSD", Decimals: 2, Name: "Bahamian Dollar"},
		Currency{Numeric: "048", ISO: "BHD", Decimals: 3, Name: "Bahraini Dinar"},
		Currency{Numeric: "050", ISO: "BDT", Decimals: 2, Name: "Taka"},
		Currency{Numeric: "051", ISO: "AMD", Decimals: 2, Name: "Armenian Dram"},
		Currency{Numeric: "052", ISO: "BBD", Decimals: 2, Name: "Barbados Dollar"},
		Currency{Numeric: "060", ISO: "BMD", Decimals: 2, Name: "Bermudian Dollar"},
		Currency{Numeric: "064", ISO: "BTN", Decimals: 2, Name: "Ngultrum"},
		Currency{Numeric: "068", ISO: "BOB", Decimals: 2, Name: "Boliviano"},
		Currency{Numeric: "072", ISO: "BWP", Decimals: 2, Name: "Pula"},
		Currency{Numeric: "084", ISO: "BZD", Decimals: 2, Name: "Belize Dollar"},
		Currency{Numeric: "090", ISO: "SBD", Decimals: 2, Name: "Solomon Islands Dollar"},
		Currency{Numeric: "096", ISO: "BND", Decimals: 2, Name: "Brunei Dollar"},
		Currency{Numeric: "104", ISO: "MMK", Decimals: 2, Name: "Kyat"},
		Currency{Numeric: "108", ISO: "BIF", Decimals: 0, Name: "Burundi Franc"},
		Currency{Numeric: "116", ISO: "KHR", Decimals: 2, Name: "Riel"},
		Currency{Numeric: "124", ISO: "CAD", Decimals: 2, Name: "Canadian Dollar"},
		Currency{Numeric: "132", ISO: "CVE", Decimals: 2, Name: "Cabo Verde Escudo"},
		Currency{Numeric: "136", ISO: "KYD", Decimals: 2, Name: "Cayman Islands Dollar"},
		Currency{Numeric: "144", ISO: "LKR", Decimals: 2, Name: "Sri Lanka Rupee"},
		Currency{Numeric: "152", ISO: "CLP", Decimals: 0, Name: "Chilean Peso"},
		Currency{Numeric: "156", ISO: "CNY", Decimals: 2, Name: "Yuan Renminbi"},
		Currency{Numeric: "170", ISO: "COP", Decimals: 2, Name: "Colombian Peso"},
		Currency{Numeric: "174", ISO: "KMF", Decimals: 0, Name: "Comorian Franc "},
		Currency{Numeric: "188", ISO: "CRC", Decimals: 2, Name: "Costa Rican Colon"},
		Currency{Numeric: "191", ISO: "HRK", Decimals: 2, Name: "Kuna"},
		Currency{Numeric: "192", ISO: "CUP", Decimals: 2, Name: "Cuban Peso"},
		Currency{Numeric: "203", ISO: "CZK", Decimals: 2, Name: "Czech Koruna"},
		Currency{Numeric: "208", ISO: "DKK", Decimals: 2, Name: "Danish Krone"},
		Currency{Numeric: "214", ISO: "DOP", Decimals: 2, Name: "Dominican Peso"},
		Currency{Numeric: "222", ISO: "SVC", Decimals: 2, Name: "El Salvador Colon"},
		Currency{Numeric: "230", ISO: "ETB", Decimals: 2, Name: "Ethiopian Birr"},
		Currency{Numeric: "232", ISO: "ERN", Decimals: 2, Name: "Nakfa"},
		Currency{Numeric: "238", ISO: "FKP", Decimals: 2, Name: "Falkland Islands Pound"},
		Currency{Numeric: "242", ISO: "FJD", Decimals: 2, Name: "Fiji Dollar"},
		Currency{Numeric: "262", ISO: "DJF", Decimals: 0, Name: "Djibouti Franc"},
		Currency{Numeric: "270", ISO: "GMD", Decimals: 2, Name: "Dalasi"},
		Currency{Numeric: "292", ISO: "GIP", Decimals: 2, Name: "Gibraltar Pound"},
		Currency{Numeric: "320", ISO: "GTQ", Decimals: 2, Name: "Quetzal"},
		Currency{Numeric: "324", ISO: "GNF", Decimals: 0, Name: "Guinean Franc"},
		Currency{Numeric: "328", ISO: "GYD", Decimals: 2, Name: "Guyana Dollar"},
		Currency{Numeric: "332", ISO: "HTG", Decimals: 2, Name: "Gourde"},
		Currency{Numeric: "340", ISO: "HNL", Decimals: 2, Name: "Lempira"},
		Currency{Numeric: "344", ISO: "HKD", Decimals: 2, Name: "Hong Kong Dollar"},
		Currency{Numeric: "348", ISO: "HUF", Decimals: 2, Name: "Forint"},
		Currency{Numeric: "352", ISO: "ISK", Decimals: 0, Name: "Iceland Krona"},
		Currency{Numeric: "356", ISO: "INR", Decimals: 2, Name: "Indian Rupee"},
		Currency{Numeric: "360", ISO: "IDR", Decimals: 2, Name: "Rupiah"},
		Currency{Numeric: "364", ISO: "IRR", Decimals: 2, Name: "Iranian Rial"},
		Currency{Numeric: "368", ISO: "IQD", Decimals: 3, Name: "Iraqi Dinar"},
		Currency{Numeric: "376", ISO: "ILS", Decimals: 2, Name: "New Israeli Sheqel"},
		Currency{Numeric: "388", ISO: "JMD", Decimals: 2, Name: "Jamaican Dollar"},
		Currency{Numeric: "392", ISO: "JPY", Decimals: 0, Name: "Yen"},
		Currency{Numeric: "398", ISO: "KZT", Decimals: 2, Name: "Tenge"},
		Currency{Numeric: "400", ISO: "JOD", Decimals: 3, Name: "Jordanian Dinar"},
		Currency{Numeric: "404", ISO: "KES", Decimals: 2, Name: "Kenyan Shilling"},
		Currency{Numeric: "408", ISO: "KPW", Decimals: 2, Name: "North Korean Won"},
		Currency{Numeric: "410", ISO: "KRW", Decimals: 0, Name: "Won"},
		Currency{Numeric: "414", ISO: "KWD", Decimals: 3, Name: "Kuwaiti Dinar"},
		Currency{Numeric: "417", ISO: "KGS", Decimals: 2, Name: "Som"},
		Currency{Numeric: "418", ISO: "LAK", Decimals: 2, Name: "Lao Kip"},
		Currency{Numeric: "422", ISO: "LBP", Decimals: 2, Name: "Lebanese Pound"},
		Currency{Numeric: "426", ISO: "LSL", Decimals: 2, Name: "Loti"},
		Currency{Numeric: "430", ISO: "LRD", Decimals: 2, Name: "Liberian Dollar"},
		Currency{Numeric: "434", ISO: "LYD", Decimals: 3, Name: "Libyan Dinar"},
		Currency{Numeric: "446", ISO: "MOP", Decimals: 2, Name: "Pataca"},
		Currency{Numeric: "454", ISO: "MWK", Decimals: 2, Name: "Malawi Kwacha"},
		Currency{Numeric: "458", ISO: "MYR", Decimals: 2, Name: "Malaysian Ringgit"},
		Currency{Numeric: "462", ISO: "MVR", Decimals: 2, Name: "Rufiyaa"},
		Currency{Numeric: "480", ISO: "MUR", Decimals: 2, Name: "Mauritius Rupee"},
		Currency{Numeric: "484", ISO: "MXN", Decimals: 2, Name: "Mexican Peso"},
		Currency{Numeric: "496", ISO: "MNT", Decimals: 2, Name: "Tugrik"},
		Currency{Numeric: "498", ISO: "MDL", Decimals: 2, Name: "Moldovan Leu"},
		Currency{Numeric: "504", ISO: "MAD", Decimals: 2, Name: "Moroccan Dirham"},
		Currency{Numeric: "512", ISO: "OMR", Decimals: 3, Name: "Rial Omani"},
		Currency{Numeric: "516", ISO: "NAD", Decimals: 2, Name: "Namibia Dollar"},
		Currency{Numeric: "524", ISO: "NPR", Decimals: 2, Name: "Nepalese Rupee"},
		Currency{Numeric: "532", ISO: "ANG", Decimals: 2, Name: "Netherlands Antillean Guilder"},
		Currency{Numeric: "533", ISO: "AWG", Decimals: 2, Name: "Aruban Florin"},
		Currency{Numeric: "548", ISO: "VUV", Decimals: 0, Name: "Vatu"},
		Currency{Numeric: "554", ISO: "NZD", Decimals: 2, Name: "New Zealand Dollar"},
		Currency{Numeric: "558", ISO: "NIO", Decimals: 2, Name: "Cordoba Oro"},
		Currency{Numeric: "566", ISO: "NGN", Decimals: 2, Name: "Naira"},
		Currency{Numeric: "578", ISO: "NOK", Decimals: 2, Name: "Norwegian Krone"},
		Currency{Numeric: "586", ISO: "PKR", Decimals: 2, Name: "Pakistan Rupee"},
		Currency{Numeric: "590", ISO: "PAB", Decimals: 2, Name: "Balboa"},
		Currency{Numeric: "598", ISO: "PGK", Decimals: 2, Name: "Kina"},
		Currency{Numeric: "600", ISO: "PYG", Decimals: 0, Name: "Guarani"},
		Currency{Numeric: "604", ISO: "PEN", Decimals: 2, Name: "Sol"},
		Currency{Numeric: "608", ISO: "PHP", Decimals: 2, Name: "Philippine Peso"},
		Currency{Numeric: "634", ISO: "QAR", Decimals: 2, Name: "Qatari Rial"},
		Currency{Numeric: "643", ISO: "RUB", Decimals: 2, Name: "Russian Ruble"},
		Currency{Numeric: "646", ISO: "RWF", Decimals: 0, Name: "Rwanda Franc"},
		Currency{Numeric: "654", ISO: "SHP", Decimals: 2, Name: "Saint Helena Pound"},
		Currency{Numeric: "682", ISO: "SAR", Decimals: 2, Name: "Saudi Riyal"},
		Currency{Numeric: "690", ISO: "SCR", Decimals: 2, Name: "Seychelles Rupee"},
		Currency{Numeric: "694", ISO: "SLL", Decimals: 2, Name: "Leone"},
		Currency{Numeric: "702", ISO: "SGD", Decimals: 2, Name: "Singapore Dollar"},
		Currency{Numeric: "704", ISO: "VND", Decimals: 0, Name: "Dong"},
		Currency{Numeric: "706", ISO: "SOS", Decimals: 2, Name: "Somali Shilling"},
		Currency{Numeric: "710", ISO: "ZAR", Decimals: 2, Name: "Rand"},
		Currency{Numeric: "728", ISO: "SSP", Decimals: 2, Name: "South Sudanese Pound"},
		Currency{Numeric: "748", ISO: "SZL", Decimals: 2, Name: "Lilangeni"},
		Currency{Numeric: "752", ISO: "SEK", Decimals: 2, Name: "Swedish Krona"},
		Currency{Numeric: "756", ISO: "CHF", Decimals: 2, Name: "Swiss Franc"},
		Currency{Numeric: "760", ISO: "SYP", Decimals: 2, Name: "Syrian Pound"},
		Currency{Numeric: "764", ISO: "THB", Decimals: 2, Name: "Baht"},
		Currency{Numeric: "776", ISO: "TOP", Decimals: 2, Name: "Pa???anga"},
		Currency{Numeric: "780", ISO: "TTD", Decimals: 2, Name: "Trinidad and Tobago Dollar"},
		Currency{Numeric: "784", ISO: "AED", Decimals: 2, Name: "UAE Dirham"},
		Currency{Numeric: "788", ISO: "TND", Decimals: 3, Name: "Tunisian Dinar"},
		Currency{Numeric: "800", ISO: "UGX", Decimals: 0, Name: "Uganda Shilling"},
		Currency{Numeric: "807", ISO: "MKD", Decimals: 2, Name: "Denar"},
		Currency{Numeric: "818", ISO: "EGP", Decimals: 2, Name: "Egyptian Pound"},
		Currency{Numeric: "826", ISO: "GBP", Decimals: 2, Name: "Pound Sterling"},
		Currency{Numeric: "834", ISO: "TZS", Decimals: 2, Name: "Tanzanian Shilling"},
		Currency{Numeric: "840", ISO: "USD", Decimals: 2, Name: "US Dollar"},
		Currency{Numeric: "858", ISO: "UYU", Decimals: 2, Name: "Peso Uruguayo"},
		Currency{Numeric: "860", ISO: "UZS", Decimals: 2, Name: "Uzbekistan Sum"},
		Currency{Numeric: "882", ISO: "WST", Decimals: 2, Name: "Tala"},
		Currency{Numeric: "886", ISO: "YER", Decimals: 2, Name: "Yemeni Rial"},
		Currency{Numeric: "901", ISO: "TWD", Decimals: 2, Name: "New Taiwan Dollar"},
		Currency{Numeric: "927", ISO: "UYW", Decimals: 4, Name: "Unidad Previsional"},
		Currency{Numeric: "928", ISO: "VES", Decimals: 2, Name: "Bol??var Soberano"},
		Currency{Numeric: "929", ISO: "MRU", Decimals: 2, Name: "Ouguiya"},
		Currency{Numeric: "930", ISO: "STN", Decimals: 2, Name: "Dobra"},
		Currency{Numeric: "931", ISO: "CUC", Decimals: 2, Name: "Peso Convertible"},
		Currency{Numeric: "932", ISO: "ZWL", Decimals: 2, Name: "Zimbabwe Dollar"},
		Currency{Numeric: "933", ISO: "BYN", Decimals: 2, Name: "Belarusian Ruble"},
		Currency{Numeric: "934", ISO: "TMT", Decimals: 2, Name: "Turkmenistan New Manat"},
		Currency{Numeric: "936", ISO: "GHS", Decimals: 2, Name: "Ghana Cedi"},
		Currency{Numeric: "938", ISO: "SDG", Decimals: 2, Name: "Sudanese Pound"},
		Currency{Numeric: "940", ISO: "UYI", Decimals: 0, Name: "Uruguay Peso en Unidades Indexadas (UI)"},
		Currency{Numeric: "941", ISO: "RSD", Decimals: 2, Name: "Serbian Dinar"},
		Currency{Numeric: "943", ISO: "MZN", Decimals: 2, Name: "Mozambique Metical"},
		Currency{Numeric: "944", ISO: "AZN", Decimals: 2, Name: "Azerbaijan Manat"},
		Currency{Numeric: "946", ISO: "RON", Decimals: 2, Name: "Romanian Leu"},
		Currency{Numeric: "947", ISO: "CHE", Decimals: 2, Name: "WIR Euro"},
		Currency{Numeric: "948", ISO: "CHW", Decimals: 2, Name: "WIR Franc"},
		Currency{Numeric: "949", ISO: "TRY", Decimals: 2, Name: "Turkish Lira"},
		Currency{Numeric: "950", ISO: "XAF", Decimals: 0, Name: "CFA Franc BEAC"},
		Currency{Numeric: "951", ISO: "XCD", Decimals: 2, Name: "East Caribbean Dollar"},
		Currency{Numeric: "952", ISO: "XOF", Decimals: 0, Name: "CFA Franc BCEAO"},
		Currency{Numeric: "953", ISO: "XPF", Decimals: 0, Name: "CFP Franc"},
		Currency{Numeric: "967", ISO: "ZMW", Decimals: 2, Name: "Zambian Kwacha"},
		Currency{Numeric: "968", ISO: "SRD", Decimals: 2, Name: "Surinam Dollar"},
		Currency{Numeric: "969", ISO: "MGA", Decimals: 2, Name: "Malagasy Ariary"},
		Currency{Numeric: "970", ISO: "COU", Decimals: 2, Name: "Unidad de Valor Real"},
		Currency{Numeric: "971", ISO: "AFN", Decimals: 2, Name: "Afghani"},
		Currency{Numeric: "972", ISO: "TJS", Decimals: 2, Name: "Somoni"},
		Currency{Numeric: "973", ISO: "AOA", Decimals: 2, Name: "Kwanza"},
		Currency{Numeric: "975", ISO: "BGN", Decimals: 2, Name: "Bulgarian Lev"},
		Currency{Numeric: "976", ISO: "CDF", Decimals: 2, Name: "Congolese Franc"},
		Currency{Numeric: "977", ISO: "BAM", Decimals: 2, Name: "Convertible Mark"},
		Currency{Numeric: "978", ISO: "EUR", Decimals: 2, Name: "Euro"},
		Currency{Numeric: "979", ISO: "MXV", Decimals: 2, Name: "Mexican Unidad de Inversion (UDI)"},
		Currency{Numeric: "980", ISO: "UAH", Decimals: 2, Name: "Hryvnia"},
		Currency{Numeric: "981", ISO: "GEL", Decimals: 2, Name: "Lari"},
		Currency{Numeric: "984", ISO: "BOV", Decimals: 2, Name: "Mvdol"},
		Currency{Numeric: "985", ISO: "PLN", Decimals: 2, Name: "Zloty"},
		Currency{Numeric: "986", ISO: "BRL", Decimals: 2, Name: "Brazilian Real"},
		Currency{Numeric: "990", ISO: "CLF", Decimals: 4, Name: "Unidad de Fomento"},

		// Registering testing currency
		xts,
	)
}
