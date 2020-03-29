package main

// https://www.cia.gov/library/publications/the-world-factbook/rankorder/2004rank.html
var pppGDPperCapita = map[string]float64{
	"LIECHTENSTEIN":                    139100,
	"QATAR":                            124500,
	"MONACO":                           115700,
	"MACAU":                            111600,
	"LUXEMBOURG":                       106300,
	"BERMUDA":                          99400,
	"SINGAPORE":                        93900,
	"ISLE OF MAN":                      84600,
	"BRUNEI":                           78200,
	"IRELAND":                          75500,
	"NORWAY":                           71800,
	"FALKLAND ISLANDS":                 70800,
	"UNITED ARAB EMIRATES":             67700,
	"SINT MAARTEN":                     66800,
	"KUWAIT":                           66200,
	"GIBRALTAR":                        61700,
	"HONG KONG":                        61400,
	"SWITZERLAND":                      61400,
	"UNITED STATES OF AMERICA":         59500,
	"SAN MARINO":                       58600,
	"JERSEY":                           56600,
	"SAUDI ARABIA":                     54800,
	"NETHERLANDS":                      53600,
	"GUERNSEY":                         52500,
	"ICELAND":                          51800,
	"SWEDEN":                           51500,
	"GERMANY":                          50400,
	"TAIWAN":                           50300,
	"AUSTRALIA":                        50300,
	"AUSTRIA":                          49900,
	"DENMARK":                          49900,
	"ANDORRA":                          49900,
	"BAHRAIN":                          48500,
	"CANADA":                           48300,
	"BELGIUM":                          46600,
	"SAINT PIERRE AND MIQUELON":        46200,
	"OMAN":                             45200,
	"FINLAND":                          44300,
	"UNITED KINGDOM":                   44100,
	"CAYMAN ISLANDS":                   43800,
	"FRANCE":                           43800,
	"JAPAN":                            42800,
	"MALTA":                            42000,
	"GREENLAND":                        41800,
	"EUROPEAN UNION":                   40900,
	"FAROE ISLANDS":                    40000,
	"SOUTH KOREA":                      39400,
	"NEW ZEALAND":                      38900,
	"SPAIN":                            38300,
	"ITALY":                            38100,
	"PUERTO RICO":                      37300,
	"VIRGIN ISLANDS":                   37000,
	"CYPRUS":                           37000,
	"ISRAEL":                           36300,
	"EQUATORIAL GUINEA":                36000,
	"GUAM":                             35600,
	"CZECHIA":                          35500,
	"SLOVENIA":                         34400,
	"BRITISH VIRGIN ISLANDS":           34200,
	"MONTSERRAT":                       34000,
	"SLOVAKIA":                         33000,
	"LITHUANIA":                        32300,
	"ESTONIA":                          31800,
	"TRINIDAD AND TOBAGO":              31400,
	"BAHAMAS":                          31200,
	"NEW CALEDONIA":                    31100,
	"PORTUGAL":                         30400,
	"HUNGARY":                          29500,
	"POLAND":                           29500,
	"TURKS AND CAICOS ISLANDS":         29100,
	"MALAYSIA":                         29000,
	"SEYCHELLES":                       28900,
	"RUSSIA":                           27800,
	"GREECE":                           27700,
	"LATVIA":                           27600,
	"TURKEY":                           26900,
	"SAINT KITTS AND NEVIS":            26800,
	"KAZAKHSTAN":                       26300,
	"ANTIGUA AND BARBUDA":              26300,
	"PANAMA":                           25400,
	"ARUBA":                            25300,
	"NORTHERN MARIANA ISLANDS":         24500,
	"CHILE":                            24500,
	"ROMANIA":                          24500,
	"CROATIA":                          24400,
	"URUGUAY":                          22400,
	"BULGARIA":                         21700,
	"MAURITIUS":                        21600,
	"ARGENTINA":                        20900,
	"IRAN":                             20200,
	"MEXICO":                           19900,
	"LEBANON":                          19400,
	"SAINT MARTIN":                     19300,
	"GABON":                            19200,
	"MALDIVES":                         19100,
	"BELARUS":                          18900,
	"BARBADOS":                         18700,
	"TURKMENISTAN":                     18100,
	"THAILAND":                         17900,
	"BOTSWANA":                         17800,
	"MONTENEGRO":                       17700,
	"AZERBAIJAN":                       17500,
	"FRENCH POLYNESIA":                 17000,
	"IRAQ":                             17000,
	"COSTA RICA":                       16900,
	"DOMINICAN REPUBLIC":               16900,
	"COOK ISLANDS":                     16700,
	"CHINA":                            16700,
	"PALAU":                            16200,
	"BRAZIL":                           15600,
	"ALGERIA":                          15200,
	"SERBIA":                           15000,
	"CURACAO":                          15000,
	"MACEDONIA":                        14900,
	"GRENADA":                          14900,
	"SURINAME":                         14600,
	"COLOMBIA":                         14500,
	"SAINT LUCIA":                      14400,
	"SOUTH AFRICA":                     13500,
	"PERU":                             13300,
	"MONGOLIA":                         13000,
	"SRI LANKA":                        12800,
	"EGYPT":                            12700,
	"BOSNIA AND HERZEGOVINA":           12700,
	"ALBANIA":                          12500,
	"JORDAN":                           12500,
	"INDONESIA":                        12400,
	"CUBA":                             12300,
	"NAURU":                            12200,
	"ANGUILLA":                         12200,
	"VENEZUELA":                        12100,
	"TUNISIA":                          11800,
	"ECUADOR":                          11500,
	"SAINT VINCENT AND THE GRENADINES": 11500,
	"NAMIBIA":                          11300,
	"AMERICAN SAMOA":                   11200,
	"DOMINICA":                         11100,
	"GEORGIA":                          10700,
	"KOSOVO":                           10500,
	"LIBYA":                            10000,
	"ESWATINI":                         9900,
	"PARAGUAY":                         9800,
	"FIJI":                             9800,
	"ARMENIA":                          9500,
	"JAMAICA":                          9200,
	"EL SALVADOR":                      8900,
	"BHUTAN":                           8700,
	"UKRAINE":                          8700,
	"MOROCCO":                          8600,
	"BELIZE":                           8300,
	"PHILIPPINES":                      8300,
	"GUYANA":                           8200,
	"GUATEMALA":                        8100,
	"SAINT HELENA, ASCENSION, AND TRISTAN DA CUNHA": 7800,
	"BOLIVIA":                           7500,
	"LAOS":                              7400,
	"INDIA":                             7200,
	"CABO VERDE":                        6900,
	"VIETNAM":                           6900,
	"UZBEKISTAN":                        6900,
	"ANGOLA":                            6800,
	"CONGO":                             6600,
	"BURMA":                             6200,
	"NIGERIA":                           5900,
	"NICARAGUA":                         5800,
	"NIUE":                              5800,
	"SAMOA":                             5700,
	"MOLDOVA":                           5700,
	"TONGA":                             5600,
	"HONDURAS":                          5600,
	"TIMOR-LESTE":                       5400,
	"PAKISTAN":                          5400,
	"GHANA":                             4700,
	"SUDAN":                             4600,
	"MAURITANIA":                        4400,
	"WEST BANK":                         4300,
	"BANGLADESH":                        4200,
	"ZAMBIA":                            4000,
	"CAMBODIA":                          4000,
	"COTE D'IVOIRE":                     3900,
	"WALLIS AND FUTUNA":                 3800,
	"TUVALU":                            3800,
	"CAMEROON":                          3700,
	"PAPUA NEW GUINEA":                  3700,
	"KYRGYZSTAN":                        3700,
	"LESOTHO":                           3600,
	"DJIBOUTI":                          3600,
	"KENYA":                             3500,
	"MARSHALL ISLANDS":                  3400,
	"MICRONESIA":                        3400,
	"TAJIKISTAN":                        3200,
	"TANZANIA":                          3200,
	"SAO TOME AND PRINCIPE":             3200,
	"SYRIA":                             2900,
	"VANUATU":                           2700,
	"SENEGAL":                           2700,
	"NEPAL":                             2700,
	"WESTERN SAHARA":                    2500,
	"UGANDA":                            2400,
	"BENIN":                             2300,
	"ZIMBABWE":                          2300,
	"CHAD":                              2300,
	"SOLOMON ISLANDS":                   2200,
	"MALI":                              2200,
	"ETHIOPIA":                          2200,
	"RWANDA":                            2100,
	"AFGHANISTAN":                       2000,
	"KIRIBATI":                          2000,
	"GUINEA":                            2000,
	"BURKINA FASO":                      1900,
	"GUINEA-BISSAU":                     1800,
	"HAITI":                             1800,
	"GAMBIA, THE":                       1700,
	"TOGO":                              1700,
	"NORTH KOREA":                       1700,
	"COMOROS":                           1600,
	"SIERRA LEONE":                      1600,
	"MADAGASCAR":                        1600,
	"ERITREA":                           1600,
	"SOUTH SUDAN":                       1500,
	"LIBERIA":                           1400,
	"YEMEN":                             1300,
	"MALAWI":                            1200,
	"NIGER":                             1200,
	"MOZAMBIQUE":                        1200,
	"TOKELAU":                           1000,
	"CONGO, DEMOCRATIC REPUBLIC OF THE": 800,
	"CENTRAL AFRICAN REPUBLIC":          700,
	"BURUNDI":                           700,
}

var ccodes = map[string]string{
	"ad": "andorra",
	"ae": "united arab emirates",
	"af": "afghanistan",
	"ag": "antigua and barbuda",
	"ai": "anguilla",
	"al": "albania",
	"am": "armenia",
	"ao": "angola",
	"aq": "antarctica",
	"ar": "argentina",
	"as": "american samoa",
	"at": "austria",
	"au": "australia",
	"aw": "aruba",
	"ax": "åland islands",
	"az": "azerbaijan",
	"ba": "bosnia and herzegovina",
	"bb": "barbados",
	"bd": "bangladesh",
	"be": "belgium",
	"bf": "burkina faso",
	"bg": "bulgaria",
	"bh": "bahrain",
	"bi": "burundi",
	"bj": "benin",
	"bl": "saint barthélemy",
	"bm": "bermuda",
	"bn": "brunei darussalam",
	"bo": "bolivia",
	"bq": "bonaire, sint eustatius and saba",
	"br": "brazil",
	"bs": "bahamas",
	"bt": "bhutan",
	"bv": "bouvet island",
	"bw": "botswana",
	"by": "belarus",
	"bz": "belize",
	"ca": "canada",
	"cc": "cocos (keeling) islands",
	"cd": "congo, democratic republic of the",
	"cf": "central african republic",
	"cg": "congo",
	"ch": "switzerland",
	"ci": "cote d'ivoire",
	"ck": "cook islands",
	"cl": "chile",
	"cm": "cameroon",
	"cn": "china",
	"co": "colombia",
	"cr": "costa rica",
	"cu": "cuba",
	"cv": "cabo verde",
	"cw": "curaçao",
	"cx": "christmas island",
	"cy": "cyprus",
	"cz": "czechia",
	"de": "germany",
	"dj": "djibouti",
	"dk": "denmark",
	"dm": "dominica",
	"do": "dominican republic",
	"dz": "algeria",
	"ec": "ecuador",
	"ee": "estonia",
	"eg": "egypt",
	"eh": "western sahara",
	"er": "eritrea",
	"es": "spain",
	"et": "ethiopia",
	"fi": "finland",
	"fj": "fiji",
	"fk": "falkland islands",
	"fm": "micronesia",
	"fo": "faroe islands",
	"fr": "france",
	"ga": "gabon",
	"gb": "united kingdom",
	"gd": "grenada",
	"ge": "georgia",
	"gf": "french guiana",
	"gg": "guernsey",
	"gh": "ghana",
	"gi": "gibraltar",
	"gl": "greenland",
	"gm": "gambia",
	"gn": "guinea",
	"gp": "guadeloupe",
	"gq": "equatorial guinea",
	"gr": "greece",
	"gs": "south georgia and the south sandwich islands",
	"gt": "guatemala",
	"gu": "guam",
	"gw": "guinea-bissau",
	"gy": "guyana",
	"hk": "hong kong",
	"hm": "heard island and mcdonald islands",
	"hn": "honduras",
	"hr": "croatia",
	"ht": "haiti",
	"hu": "hungary",
	"id": "indonesia",
	"ie": "ireland",
	"il": "israel",
	"im": "isle of man",
	"in": "india",
	"io": "british indian ocean territory",
	"iq": "iraq",
	"ir": "iran",
	"is": "iceland",
	"it": "italy",
	"je": "jersey",
	"jm": "jamaica",
	"jo": "jordan",
	"jp": "japan",
	"ke": "kenya",
	"kg": "kyrgyzstan",
	"kh": "cambodia",
	"ki": "kiribati",
	"km": "comoros",
	"kn": "saint kitts and nevis",
	"kp": "north korea",
	"kr": "south korea",
	"kw": "kuwait",
	"ky": "cayman islands",
	"kz": "kazakhstan",
	"la": "lao",
	"lb": "lebanon",
	"lc": "saint lucia",
	"li": "liechtenstein",
	"lk": "sri lanka",
	"lr": "liberia",
	"ls": "lesotho",
	"lt": "lithuania",
	"lu": "luxembourg",
	"lv": "latvia",
	"ly": "libya",
	"ma": "morocco",
	"mc": "monaco",
	"md": "moldova",
	"me": "montenegro",
	"mf": "saint martin (french part)",
	"mg": "madagascar",
	"mh": "marshall islands",
	"mk": "north macedonia",
	"ml": "mali",
	"mm": "myanmar",
	"mn": "mongolia",
	"mo": "macao",
	"mp": "northern mariana islands",
	"mq": "martinique",
	"mr": "mauritania",
	"ms": "montserrat",
	"mt": "malta",
	"mu": "mauritius",
	"mv": "maldives",
	"mw": "malawi",
	"mx": "mexico",
	"my": "malaysia",
	"mz": "mozambique",
	"na": "namibia",
	"nc": "new caledonia",
	"ne": "niger",
	"nf": "norfolk island",
	"ng": "nigeria",
	"ni": "nicaragua",
	"nl": "netherlands",
	"no": "norway",
	"np": "nepal",
	"nr": "nauru",
	"nu": "niue",
	"nz": "new zealand",
	"om": "oman",
	"pa": "panama",
	"pe": "peru",
	"pf": "french polynesia",
	"pg": "papua new guinea",
	"ph": "philippines",
	"pk": "pakistan",
	"pl": "poland",
	"pm": "saint pierre and miquelon",
	"pn": "pitcairn",
	"pr": "puerto rico",
	"ps": "palestine",
	"pt": "portugal",
	"pw": "palau",
	"py": "paraguay",
	"qa": "qatar",
	"re": "réunion",
	"ro": "romania",
	"rs": "serbia",
	"ru": "russia",
	"rw": "rwanda",
	"sa": "saudi arabia",
	"sb": "solomon islands",
	"sc": "seychelles",
	"sd": "sudan",
	"se": "sweden",
	"sg": "singapore",
	"sh": "saint helena, ascension and tristan da cunha",
	"si": "slovenia",
	"sj": "svalbard and jan mayen",
	"sk": "slovakia",
	"sl": "sierra leone",
	"sm": "san marino",
	"sn": "senegal",
	"so": "somalia",
	"sr": "suriname",
	"ss": "south sudan",
	"st": "sao tome and principe",
	"sv": "el salvador",
	"sx": "sint maarten (dutch part)",
	"sy": "syria",
	"sz": "eswatini",
	"tc": "turks and caicos islands",
	"td": "chad",
	"tf": "french southern territories",
	"tg": "togo",
	"th": "thailand",
	"tj": "tajikistan",
	"tk": "tokelau",
	"tl": "timor-leste",
	"tm": "turkmenistan",
	"tn": "tunisia",
	"to": "tonga",
	"tr": "turkey",
	"tt": "trinidad and tobago",
	"tv": "tuvalu",
	"tw": "taiwan",
	"tz": "tanzania",
	"ua": "ukraine",
	"ug": "uganda",
	"um": "united states minor outlying islands",
	"us": "united states of america",
	"uy": "uruguay",
	"uz": "uzbekistan",
	"va": "holy see",
	"vc": "saint vincent and the grenadines",
	"ve": "venezuela (bolivarian republic of)",
	"vg": "virgin islands (british)",
	"vi": "virgin islands (u.s.)",
	"vn": "vietnam",
	"vu": "vanuatu",
	"wf": "wallis and futuna",
	"ws": "samoa",
	"ye": "yemen",
	"yt": "mayotte",
	"za": "south africa",
	"zm": "zambia",
	"zw": "zimbabwe",
}
