// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra-locale package configuration.
// 
// This module contains definitions
// for the following management objects:
//   locale: Define the geographical locale
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_infra_locale_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_infra_locale_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-infra-locale-cfg locale}", reflect.TypeOf(Locale{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-infra-locale-cfg:locale", reflect.TypeOf(Locale{}))
}

// LocaleCountry represents Locale country
type LocaleCountry string

const (
    // Andorra
    LocaleCountry_ad LocaleCountry = "ad"

    // United Arab Emirates
    LocaleCountry_ae LocaleCountry = "ae"

    // Afghanistan
    LocaleCountry_af LocaleCountry = "af"

    // Antigua and Barbuda
    LocaleCountry_ag LocaleCountry = "ag"

    // Anguilla
    LocaleCountry_ai LocaleCountry = "ai"

    // Albania
    LocaleCountry_al LocaleCountry = "al"

    // Armenia
    LocaleCountry_am LocaleCountry = "am"

    // Netherlands Antilles
    LocaleCountry_an LocaleCountry = "an"

    // Angola
    LocaleCountry_ao LocaleCountry = "ao"

    // Antarctica
    LocaleCountry_aq LocaleCountry = "aq"

    // Argentina
    LocaleCountry_ar LocaleCountry = "ar"

    // American Samoa
    LocaleCountry_as LocaleCountry = "as"

    // Austria
    LocaleCountry_at LocaleCountry = "at"

    // Australia
    LocaleCountry_au LocaleCountry = "au"

    // Aruba
    LocaleCountry_aw LocaleCountry = "aw"

    // Azerbaijan
    LocaleCountry_az LocaleCountry = "az"

    // Bosnia and Herzegovina
    LocaleCountry_ba LocaleCountry = "ba"

    // Barbados
    LocaleCountry_bb LocaleCountry = "bb"

    // Bangladesh
    LocaleCountry_bd LocaleCountry = "bd"

    // Belgium
    LocaleCountry_be LocaleCountry = "be"

    // Burkina Faso
    LocaleCountry_bf LocaleCountry = "bf"

    // Bulgaria
    LocaleCountry_bg LocaleCountry = "bg"

    // Bahrain
    LocaleCountry_bh LocaleCountry = "bh"

    // Burundi
    LocaleCountry_bi LocaleCountry = "bi"

    // Benin
    LocaleCountry_bj LocaleCountry = "bj"

    // Bermuda
    LocaleCountry_bm LocaleCountry = "bm"

    // Brunei Darussalam
    LocaleCountry_bn LocaleCountry = "bn"

    // Bolivia
    LocaleCountry_bo LocaleCountry = "bo"

    // Brazil
    LocaleCountry_br LocaleCountry = "br"

    // Bahamas
    LocaleCountry_bs LocaleCountry = "bs"

    // Bhutan
    LocaleCountry_bt LocaleCountry = "bt"

    // Bouvet Island
    LocaleCountry_bv LocaleCountry = "bv"

    // Botswana
    LocaleCountry_bw LocaleCountry = "bw"

    // Belarus
    LocaleCountry_by LocaleCountry = "by"

    // Belize
    LocaleCountry_bz LocaleCountry = "bz"

    // Canada
    LocaleCountry_ca LocaleCountry = "ca"

    // Cocos (Keeling) Islands
    LocaleCountry_cc LocaleCountry = "cc"

    // Congo, The Democratic Republic of the (Zaire)
    LocaleCountry_cd LocaleCountry = "cd"

    // Central African Republic
    LocaleCountry_cf LocaleCountry = "cf"

    // Congo
    LocaleCountry_cg LocaleCountry = "cg"

    // Switzerland
    LocaleCountry_ch LocaleCountry = "ch"

    // Cote D'Ivoire
    LocaleCountry_ci LocaleCountry = "ci"

    // Cook Islands
    LocaleCountry_ck LocaleCountry = "ck"

    // Chile
    LocaleCountry_cl LocaleCountry = "cl"

    // Cameroon
    LocaleCountry_cm LocaleCountry = "cm"

    // China
    LocaleCountry_cn LocaleCountry = "cn"

    // Colombia
    LocaleCountry_co LocaleCountry = "co"

    // Costa Rica
    LocaleCountry_cr LocaleCountry = "cr"

    // Cuba
    LocaleCountry_cu LocaleCountry = "cu"

    // Cape Verde
    LocaleCountry_cv LocaleCountry = "cv"

    // Christmas Island
    LocaleCountry_cx LocaleCountry = "cx"

    // Cyprus
    LocaleCountry_cy LocaleCountry = "cy"

    // Czech Republic
    LocaleCountry_cz LocaleCountry = "cz"

    // Germany
    LocaleCountry_de LocaleCountry = "de"

    // Djibouti
    LocaleCountry_dj LocaleCountry = "dj"

    // Denmark
    LocaleCountry_dk LocaleCountry = "dk"

    // Dominica
    LocaleCountry_dm LocaleCountry = "dm"

    // Dominican Republic
    LocaleCountry_do LocaleCountry = "do"

    // Algeria
    LocaleCountry_dz LocaleCountry = "dz"

    // Ecuador
    LocaleCountry_ec LocaleCountry = "ec"

    // Estonia
    LocaleCountry_ee LocaleCountry = "ee"

    // Egypt
    LocaleCountry_eg LocaleCountry = "eg"

    // Western Sahara
    LocaleCountry_eh LocaleCountry = "eh"

    // Eritrea
    LocaleCountry_er LocaleCountry = "er"

    // Spain
    LocaleCountry_es LocaleCountry = "es"

    // Ethiopia
    LocaleCountry_et LocaleCountry = "et"

    // Finland
    LocaleCountry_fi LocaleCountry = "fi"

    // Fiji
    LocaleCountry_fj LocaleCountry = "fj"

    // Falkland Islands (Malvinas)
    LocaleCountry_fk LocaleCountry = "fk"

    // Micronesia, Federated States of
    LocaleCountry_fm LocaleCountry = "fm"

    // Faroe Islands
    LocaleCountry_fo LocaleCountry = "fo"

    // France
    LocaleCountry_fr LocaleCountry = "fr"

    // Gabon
    LocaleCountry_ga LocaleCountry = "ga"

    // United Kingdom
    LocaleCountry_gb LocaleCountry = "gb"

    // Grenada
    LocaleCountry_gd LocaleCountry = "gd"

    // Georgia
    LocaleCountry_ge LocaleCountry = "ge"

    // French Guiana
    LocaleCountry_gf LocaleCountry = "gf"

    // Ghana
    LocaleCountry_gh LocaleCountry = "gh"

    // Gibraltar
    LocaleCountry_gi LocaleCountry = "gi"

    // Greenland
    LocaleCountry_gl LocaleCountry = "gl"

    // Gambia
    LocaleCountry_gm LocaleCountry = "gm"

    // Guinea
    LocaleCountry_gn LocaleCountry = "gn"

    // Guadeloupe
    LocaleCountry_gp LocaleCountry = "gp"

    // Equatorial Guinea
    LocaleCountry_gq LocaleCountry = "gq"

    // Greece
    LocaleCountry_gr LocaleCountry = "gr"

    // South Georgia and the South Sandwich Islands
    LocaleCountry_gs LocaleCountry = "gs"

    // Guatemala
    LocaleCountry_gt LocaleCountry = "gt"

    // Guam
    LocaleCountry_gu LocaleCountry = "gu"

    // Guinea Bissau
    LocaleCountry_gw LocaleCountry = "gw"

    // Guyana
    LocaleCountry_gy LocaleCountry = "gy"

    // Hong Kong
    LocaleCountry_hk LocaleCountry = "hk"

    // Heard Island and McDonald Islands
    LocaleCountry_hm LocaleCountry = "hm"

    // Honduras
    LocaleCountry_hn LocaleCountry = "hn"

    // Croatia
    LocaleCountry_hr LocaleCountry = "hr"

    // Haiti
    LocaleCountry_ht LocaleCountry = "ht"

    // Hungary
    LocaleCountry_hu LocaleCountry = "hu"

    // Indonesia
    LocaleCountry_id LocaleCountry = "id"

    // Ireland
    LocaleCountry_ie LocaleCountry = "ie"

    // Israel
    LocaleCountry_il LocaleCountry = "il"

    // India
    LocaleCountry_in LocaleCountry = "in"

    // British Indian Ocean Territory
    LocaleCountry_io LocaleCountry = "io"

    // Iraq
    LocaleCountry_iq LocaleCountry = "iq"

    // Iran, Islamic Republic of
    LocaleCountry_ir LocaleCountry = "ir"

    // Iceland
    LocaleCountry_is LocaleCountry = "is"

    // Italy
    LocaleCountry_it LocaleCountry = "it"

    // Jamaica
    LocaleCountry_jm LocaleCountry = "jm"

    // Jordan
    LocaleCountry_jo LocaleCountry = "jo"

    // Japan
    LocaleCountry_jp LocaleCountry = "jp"

    // Kenya
    LocaleCountry_ke LocaleCountry = "ke"

    // Kyrgyzstan
    LocaleCountry_kg LocaleCountry = "kg"

    // Cambodia
    LocaleCountry_kh LocaleCountry = "kh"

    // Kiribati
    LocaleCountry_ki LocaleCountry = "ki"

    // Comoros
    LocaleCountry_km LocaleCountry = "km"

    // Saint Kitts and Nevis
    LocaleCountry_kn LocaleCountry = "kn"

    // Korea, Democratic People's Republic of
    LocaleCountry_kp LocaleCountry = "kp"

    // Korea, Republic of
    LocaleCountry_kr LocaleCountry = "kr"

    // Kuwait
    LocaleCountry_kw LocaleCountry = "kw"

    // Cayman Islands
    LocaleCountry_ky LocaleCountry = "ky"

    // Kazakstan
    LocaleCountry_kz LocaleCountry = "kz"

    // Lao People's Democratic Republic
    LocaleCountry_la LocaleCountry = "la"

    // Lebanon
    LocaleCountry_lb LocaleCountry = "lb"

    // Saint Lucia
    LocaleCountry_lc LocaleCountry = "lc"

    // Liechtenstein
    LocaleCountry_li LocaleCountry = "li"

    // Sri Lanka
    LocaleCountry_lk LocaleCountry = "lk"

    // Liberia
    LocaleCountry_lr LocaleCountry = "lr"

    // Lesotho
    LocaleCountry_ls LocaleCountry = "ls"

    // Lithuania
    LocaleCountry_lt LocaleCountry = "lt"

    // Luxembourg
    LocaleCountry_lu LocaleCountry = "lu"

    // Latvia
    LocaleCountry_lv LocaleCountry = "lv"

    // Libyan Arab Jamahiriya
    LocaleCountry_ly LocaleCountry = "ly"

    // Morocco
    LocaleCountry_ma LocaleCountry = "ma"

    // Monaco
    LocaleCountry_mc LocaleCountry = "mc"

    // Moldova, Republic of
    LocaleCountry_md LocaleCountry = "md"

    // Madagascar
    LocaleCountry_mg LocaleCountry = "mg"

    // Marshall Islands
    LocaleCountry_mh LocaleCountry = "mh"

    // Macedonia, The Former Yugoslav Republic of
    LocaleCountry_mk LocaleCountry = "mk"

    // Mali
    LocaleCountry_ml LocaleCountry = "ml"

    // Myanmar
    LocaleCountry_mm LocaleCountry = "mm"

    // Mongolia
    LocaleCountry_mn LocaleCountry = "mn"

    // Macau
    LocaleCountry_mo LocaleCountry = "mo"

    // Northern Mariana Islands
    LocaleCountry_mp LocaleCountry = "mp"

    // Martinique
    LocaleCountry_mq LocaleCountry = "mq"

    // Mauritania
    LocaleCountry_mr LocaleCountry = "mr"

    // Montserrat
    LocaleCountry_ms LocaleCountry = "ms"

    // Malta
    LocaleCountry_mt LocaleCountry = "mt"

    // Mauritius
    LocaleCountry_mu LocaleCountry = "mu"

    // Maldives
    LocaleCountry_mv LocaleCountry = "mv"

    // Malawi
    LocaleCountry_mw LocaleCountry = "mw"

    // Mexico
    LocaleCountry_mx LocaleCountry = "mx"

    // Malaysia
    LocaleCountry_my LocaleCountry = "my"

    // Mozambique
    LocaleCountry_mz LocaleCountry = "mz"

    // Namibia
    LocaleCountry_na LocaleCountry = "na"

    // New Caledonia
    LocaleCountry_nc LocaleCountry = "nc"

    // Niger
    LocaleCountry_ne LocaleCountry = "ne"

    // Norfolk Island
    LocaleCountry_nf LocaleCountry = "nf"

    // Nigeria
    LocaleCountry_ng LocaleCountry = "ng"

    // Nicaragua
    LocaleCountry_ni LocaleCountry = "ni"

    // Netherlands
    LocaleCountry_nl LocaleCountry = "nl"

    // Norway
    LocaleCountry_no LocaleCountry = "no"

    // Nepal
    LocaleCountry_np LocaleCountry = "np"

    // Nauru
    LocaleCountry_nr LocaleCountry = "nr"

    // Niue
    LocaleCountry_nu LocaleCountry = "nu"

    // New Zealand
    LocaleCountry_nz LocaleCountry = "nz"

    // Oman
    LocaleCountry_om LocaleCountry = "om"

    // Panama
    LocaleCountry_pa LocaleCountry = "pa"

    // Peru
    LocaleCountry_pe LocaleCountry = "pe"

    // French Polynesia
    LocaleCountry_pf LocaleCountry = "pf"

    // Papua New Guinea
    LocaleCountry_pg LocaleCountry = "pg"

    // Philippines
    LocaleCountry_ph LocaleCountry = "ph"

    // Pakistan
    LocaleCountry_pk LocaleCountry = "pk"

    // Poland
    LocaleCountry_pl LocaleCountry = "pl"

    // Saint Pierre and Miquelon
    LocaleCountry_pm LocaleCountry = "pm"

    // Pitcairn
    LocaleCountry_pn LocaleCountry = "pn"

    // Puerto Rico
    LocaleCountry_pr LocaleCountry = "pr"

    // Portugal
    LocaleCountry_pt LocaleCountry = "pt"

    // Palau
    LocaleCountry_pw LocaleCountry = "pw"

    // Paraguay
    LocaleCountry_py LocaleCountry = "py"

    // Qatar
    LocaleCountry_qa LocaleCountry = "qa"

    // Reunion
    LocaleCountry_re LocaleCountry = "re"

    // Romania
    LocaleCountry_ro LocaleCountry = "ro"

    // Russian Federation
    LocaleCountry_ru LocaleCountry = "ru"

    // Rwanda
    LocaleCountry_rw LocaleCountry = "rw"

    // Saudi Arabia
    LocaleCountry_sa LocaleCountry = "sa"

    // Solomon Islands
    LocaleCountry_sb LocaleCountry = "sb"

    // Seychelles
    LocaleCountry_sc LocaleCountry = "sc"

    // Sudan
    LocaleCountry_sd LocaleCountry = "sd"

    // Sweden
    LocaleCountry_se LocaleCountry = "se"

    // Singapore
    LocaleCountry_sg LocaleCountry = "sg"

    // Saint Helena
    LocaleCountry_sh LocaleCountry = "sh"

    // Slovenia
    LocaleCountry_si LocaleCountry = "si"

    // Svalbard and Jan Mayen
    LocaleCountry_sj LocaleCountry = "sj"

    // Slovakia
    LocaleCountry_sk LocaleCountry = "sk"

    // Sierra Leone
    LocaleCountry_sl LocaleCountry = "sl"

    // San Marino
    LocaleCountry_sm LocaleCountry = "sm"

    // Senegal
    LocaleCountry_sn LocaleCountry = "sn"

    // Somalia
    LocaleCountry_so LocaleCountry = "so"

    // Suriname
    LocaleCountry_sr LocaleCountry = "sr"

    // Sao Tome and Principe
    LocaleCountry_st LocaleCountry = "st"

    // El Salvador
    LocaleCountry_sv LocaleCountry = "sv"

    // Syrian Arab Republic
    LocaleCountry_sy LocaleCountry = "sy"

    // Swaziland
    LocaleCountry_sz LocaleCountry = "sz"

    // Turks and Caicos Islands
    LocaleCountry_tc LocaleCountry = "tc"

    // Chad
    LocaleCountry_td LocaleCountry = "td"

    // French Southern Territories
    LocaleCountry_tf LocaleCountry = "tf"

    // Togo
    LocaleCountry_tg LocaleCountry = "tg"

    // Thailand
    LocaleCountry_th LocaleCountry = "th"

    // Tajikistan
    LocaleCountry_tj LocaleCountry = "tj"

    // Tokelau
    LocaleCountry_tk LocaleCountry = "tk"

    // Turkmenistan
    LocaleCountry_tm LocaleCountry = "tm"

    // Tunisia
    LocaleCountry_tn LocaleCountry = "tn"

    // Tonga
    LocaleCountry_to LocaleCountry = "to"

    // East Timor
    LocaleCountry_tp LocaleCountry = "tp"

    // Turkey
    LocaleCountry_tr LocaleCountry = "tr"

    // Trinidad and Tobago
    LocaleCountry_tt LocaleCountry = "tt"

    // Tuvalu
    LocaleCountry_tv LocaleCountry = "tv"

    // Taiwan, Province of China
    LocaleCountry_tw LocaleCountry = "tw"

    // Tanzania, United Republic of
    LocaleCountry_tz LocaleCountry = "tz"

    // Ukraine
    LocaleCountry_ua LocaleCountry = "ua"

    // Uganda
    LocaleCountry_ug LocaleCountry = "ug"

    // United States Minor Outlying Islands
    LocaleCountry_um LocaleCountry = "um"

    // United States
    LocaleCountry_us LocaleCountry = "us"

    // Uruguay
    LocaleCountry_uy LocaleCountry = "uy"

    // Uzbekistan
    LocaleCountry_uz LocaleCountry = "uz"

    // Holy See (Vatican City State)
    LocaleCountry_va LocaleCountry = "va"

    // Saint Vincent and The Grenadines
    LocaleCountry_vc LocaleCountry = "vc"

    // Venezuela
    LocaleCountry_ve LocaleCountry = "ve"

    // Virgin Islands, British
    LocaleCountry_vg LocaleCountry = "vg"

    // Virgin Islands, U.S.
    LocaleCountry_vi LocaleCountry = "vi"

    // Vietnam
    LocaleCountry_vn LocaleCountry = "vn"

    // Vanuatu
    LocaleCountry_vu LocaleCountry = "vu"

    // Wallis and Futuna
    LocaleCountry_wf LocaleCountry = "wf"

    // Samoa
    LocaleCountry_ws LocaleCountry = "ws"

    // Yemen
    LocaleCountry_ye LocaleCountry = "ye"

    // Mayotte
    LocaleCountry_yt LocaleCountry = "yt"

    // Yugoslavia
    LocaleCountry_yu LocaleCountry = "yu"

    // South Africa
    LocaleCountry_za LocaleCountry = "za"

    // Zambia
    LocaleCountry_zm LocaleCountry = "zm"

    // Zimbabwe
    LocaleCountry_zw LocaleCountry = "zw"
)

// LocaleLanguage represents Locale language
type LocaleLanguage string

const (
    // Afar
    LocaleLanguage_aa LocaleLanguage = "aa"

    // Abkhazian
    LocaleLanguage_ab LocaleLanguage = "ab"

    // Afrikaans
    LocaleLanguage_af LocaleLanguage = "af"

    // Amharic
    LocaleLanguage_am LocaleLanguage = "am"

    // Arabic
    LocaleLanguage_ar LocaleLanguage = "ar"

    // Assamese
    LocaleLanguage_as LocaleLanguage = "as"

    // Aymara
    LocaleLanguage_ay LocaleLanguage = "ay"

    // Azerbaijani
    LocaleLanguage_az LocaleLanguage = "az"

    // Bashkir
    LocaleLanguage_ba LocaleLanguage = "ba"

    // Byelorussian
    LocaleLanguage_be LocaleLanguage = "be"

    // Bulgarian
    LocaleLanguage_bg LocaleLanguage = "bg"

    // Bihari
    LocaleLanguage_bh LocaleLanguage = "bh"

    // Bislama
    LocaleLanguage_bi LocaleLanguage = "bi"

    // Bengali
    LocaleLanguage_bn LocaleLanguage = "bn"

    // Tibetan
    LocaleLanguage_bo LocaleLanguage = "bo"

    // Breton
    LocaleLanguage_br LocaleLanguage = "br"

    // Catalan
    LocaleLanguage_ca LocaleLanguage = "ca"

    // Corsican
    LocaleLanguage_co LocaleLanguage = "co"

    // Czech
    LocaleLanguage_cs LocaleLanguage = "cs"

    // Welsh
    LocaleLanguage_cy LocaleLanguage = "cy"

    // Danish
    LocaleLanguage_da LocaleLanguage = "da"

    // German
    LocaleLanguage_de LocaleLanguage = "de"

    // Bhutani
    LocaleLanguage_dz LocaleLanguage = "dz"

    // Greek
    LocaleLanguage_el LocaleLanguage = "el"

    // English
    LocaleLanguage_en LocaleLanguage = "en"

    // Esperanto
    LocaleLanguage_eo LocaleLanguage = "eo"

    // Spanish
    LocaleLanguage_es LocaleLanguage = "es"

    // Estonian
    LocaleLanguage_et LocaleLanguage = "et"

    // Basque
    LocaleLanguage_eu LocaleLanguage = "eu"

    // Persian
    LocaleLanguage_fa LocaleLanguage = "fa"

    // Finnish
    LocaleLanguage_fi LocaleLanguage = "fi"

    // Fiji
    LocaleLanguage_fj LocaleLanguage = "fj"

    // Faroese
    LocaleLanguage_fo LocaleLanguage = "fo"

    // French
    LocaleLanguage_fr LocaleLanguage = "fr"

    // Frisian
    LocaleLanguage_fy LocaleLanguage = "fy"

    // Irish
    LocaleLanguage_ga LocaleLanguage = "ga"

    // Scots Gaelic
    LocaleLanguage_gd LocaleLanguage = "gd"

    // Galician
    LocaleLanguage_gl LocaleLanguage = "gl"

    // Guarani
    LocaleLanguage_gn LocaleLanguage = "gn"

    // Gujarati
    LocaleLanguage_gu LocaleLanguage = "gu"

    // Hausa
    LocaleLanguage_ha LocaleLanguage = "ha"

    // Hebrew
    LocaleLanguage_he LocaleLanguage = "he"

    // Hindi
    LocaleLanguage_hi LocaleLanguage = "hi"

    // Croatian
    LocaleLanguage_hr LocaleLanguage = "hr"

    // Hungarian
    LocaleLanguage_hu LocaleLanguage = "hu"

    // Armenian
    LocaleLanguage_hy LocaleLanguage = "hy"

    // Interlingua
    LocaleLanguage_ia LocaleLanguage = "ia"

    // Indonesian
    LocaleLanguage_id LocaleLanguage = "id"

    // Interlingue
    LocaleLanguage_ie LocaleLanguage = "ie"

    // Inupiak
    LocaleLanguage_ik LocaleLanguage = "ik"

    // Icelandic
    LocaleLanguage_is LocaleLanguage = "is"

    // Italian
    LocaleLanguage_it LocaleLanguage = "it"

    // Inuktitut
    LocaleLanguage_iu LocaleLanguage = "iu"

    // Japanese
    LocaleLanguage_ja LocaleLanguage = "ja"

    // Javanese
    LocaleLanguage_jw LocaleLanguage = "jw"

    // Georgian
    LocaleLanguage_ka LocaleLanguage = "ka"

    // Kazakh
    LocaleLanguage_kk LocaleLanguage = "kk"

    // Greenlandic
    LocaleLanguage_kl LocaleLanguage = "kl"

    // Cambodian
    LocaleLanguage_km LocaleLanguage = "km"

    // Kannada
    LocaleLanguage_kn LocaleLanguage = "kn"

    // Korean
    LocaleLanguage_ko LocaleLanguage = "ko"

    // Kashmiri
    LocaleLanguage_ks LocaleLanguage = "ks"

    // Kurdish
    LocaleLanguage_ku LocaleLanguage = "ku"

    // Kirghiz
    LocaleLanguage_ky LocaleLanguage = "ky"

    // Latin
    LocaleLanguage_la LocaleLanguage = "la"

    // Lingala
    LocaleLanguage_ln LocaleLanguage = "ln"

    // Laothian
    LocaleLanguage_lo LocaleLanguage = "lo"

    // Lithuanian
    LocaleLanguage_lt LocaleLanguage = "lt"

    // Latvian, Lettish
    LocaleLanguage_lv LocaleLanguage = "lv"

    // Malagasy
    LocaleLanguage_mg LocaleLanguage = "mg"

    // Maori
    LocaleLanguage_mi LocaleLanguage = "mi"

    // Macedonian
    LocaleLanguage_mk LocaleLanguage = "mk"

    // Malayalam
    LocaleLanguage_ml LocaleLanguage = "ml"

    // Mongolian
    LocaleLanguage_mn LocaleLanguage = "mn"

    // Moldavian
    LocaleLanguage_mo LocaleLanguage = "mo"

    // Marathi
    LocaleLanguage_mr LocaleLanguage = "mr"

    // Malay
    LocaleLanguage_ms LocaleLanguage = "ms"

    // Maltese
    LocaleLanguage_mt LocaleLanguage = "mt"

    // Burmese
    LocaleLanguage_my LocaleLanguage = "my"

    // Nauru
    LocaleLanguage_na LocaleLanguage = "na"

    // Nepali
    LocaleLanguage_ne LocaleLanguage = "ne"

    // Dutch
    LocaleLanguage_nl LocaleLanguage = "nl"

    // Norwegian
    LocaleLanguage_no LocaleLanguage = "no"

    // Occitan
    LocaleLanguage_oc LocaleLanguage = "oc"

    // (Afan) Oromo
    LocaleLanguage_om LocaleLanguage = "om"

    // Oriya
    LocaleLanguage_or LocaleLanguage = "or"

    // Punjabi
    LocaleLanguage_pa LocaleLanguage = "pa"

    // Polish
    LocaleLanguage_pl LocaleLanguage = "pl"

    // Pashto, Pushto
    LocaleLanguage_ps LocaleLanguage = "ps"

    // Portuguese
    LocaleLanguage_pt LocaleLanguage = "pt"

    // Quechua
    LocaleLanguage_qu LocaleLanguage = "qu"

    // Rhaeto Romance
    LocaleLanguage_rm LocaleLanguage = "rm"

    // Kirundi
    LocaleLanguage_rn LocaleLanguage = "rn"

    // Romanian
    LocaleLanguage_ro LocaleLanguage = "ro"

    // Russian
    LocaleLanguage_ru LocaleLanguage = "ru"

    // Kinyarwanda
    LocaleLanguage_rw LocaleLanguage = "rw"

    // Sanskrit
    LocaleLanguage_sa LocaleLanguage = "sa"

    // Sindhi
    LocaleLanguage_sd LocaleLanguage = "sd"

    // Sangho
    LocaleLanguage_sg LocaleLanguage = "sg"

    // Serbo Croatian
    LocaleLanguage_sh LocaleLanguage = "sh"

    // Sinhalese
    LocaleLanguage_si LocaleLanguage = "si"

    // Slovak
    LocaleLanguage_sk LocaleLanguage = "sk"

    // Slovenian
    LocaleLanguage_sl LocaleLanguage = "sl"

    // Samoan
    LocaleLanguage_sm LocaleLanguage = "sm"

    // Shona
    LocaleLanguage_sn LocaleLanguage = "sn"

    // Somali
    LocaleLanguage_so LocaleLanguage = "so"

    // Albanian
    LocaleLanguage_sq LocaleLanguage = "sq"

    // Serbian
    LocaleLanguage_sr LocaleLanguage = "sr"

    // Siswati
    LocaleLanguage_ss LocaleLanguage = "ss"

    // Sesotho
    LocaleLanguage_st LocaleLanguage = "st"

    // Sundanese
    LocaleLanguage_su LocaleLanguage = "su"

    // Swedish
    LocaleLanguage_sv LocaleLanguage = "sv"

    // Swahili
    LocaleLanguage_sw LocaleLanguage = "sw"

    // Tamil
    LocaleLanguage_ta LocaleLanguage = "ta"

    // Telugu
    LocaleLanguage_te LocaleLanguage = "te"

    // Tajik
    LocaleLanguage_tg LocaleLanguage = "tg"

    // Thai
    LocaleLanguage_th LocaleLanguage = "th"

    // Tigrinya
    LocaleLanguage_ti LocaleLanguage = "ti"

    // Turkmen
    LocaleLanguage_tk LocaleLanguage = "tk"

    // Tagalog
    LocaleLanguage_tl LocaleLanguage = "tl"

    // Setswana
    LocaleLanguage_tn LocaleLanguage = "tn"

    // Tonga
    LocaleLanguage_to LocaleLanguage = "to"

    // Turkish
    LocaleLanguage_tr LocaleLanguage = "tr"

    // Tsonga
    LocaleLanguage_ts LocaleLanguage = "ts"

    // Tatar
    LocaleLanguage_tt LocaleLanguage = "tt"

    // Twi
    LocaleLanguage_tw LocaleLanguage = "tw"

    // Uighur
    LocaleLanguage_ug LocaleLanguage = "ug"

    // Ukrainian
    LocaleLanguage_uk LocaleLanguage = "uk"

    // Urdu
    LocaleLanguage_ur LocaleLanguage = "ur"

    // Uzbek
    LocaleLanguage_uz LocaleLanguage = "uz"

    // Vietnamese
    LocaleLanguage_vi LocaleLanguage = "vi"

    // Volapuk
    LocaleLanguage_vo LocaleLanguage = "vo"

    // Wolof
    LocaleLanguage_wo LocaleLanguage = "wo"

    // Xhosa
    LocaleLanguage_xh LocaleLanguage = "xh"

    // Yiddish
    LocaleLanguage_yi LocaleLanguage = "yi"

    // Yoruba
    LocaleLanguage_yo LocaleLanguage = "yo"

    // Zhuang
    LocaleLanguage_za LocaleLanguage = "za"

    // Chinese
    LocaleLanguage_zh LocaleLanguage = "zh"

    // Zulu
    LocaleLanguage_zu LocaleLanguage = "zu"
)

// Locale
// Define the geographical locale
type Locale struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of country locale. The type is LocaleCountry.
    Country interface{}

    // Name of language locale. The type is LocaleLanguage.
    Language interface{}
}

func (locale *Locale) GetEntityData() *types.CommonEntityData {
    locale.EntityData.YFilter = locale.YFilter
    locale.EntityData.YangName = "locale"
    locale.EntityData.BundleName = "cisco_ios_xr"
    locale.EntityData.ParentYangName = "Cisco-IOS-XR-infra-infra-locale-cfg"
    locale.EntityData.SegmentPath = "Cisco-IOS-XR-infra-infra-locale-cfg:locale"
    locale.EntityData.AbsolutePath = locale.EntityData.SegmentPath
    locale.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locale.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locale.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locale.EntityData.Children = types.NewOrderedMap()
    locale.EntityData.Leafs = types.NewOrderedMap()
    locale.EntityData.Leafs.Append("country", types.YLeaf{"Country", locale.Country})
    locale.EntityData.Leafs.Append("language", types.YLeaf{"Language", locale.Language})

    locale.EntityData.YListKeys = []string {}

    return &(locale.EntityData)
}

