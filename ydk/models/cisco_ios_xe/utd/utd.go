// Cisco XE Native Unified Threat Defense (UTD) Yang model.
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package utd

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package utd"))
}

// UtdCategoryType
type UtdCategoryType string

const (
    UtdCategoryType_abortion UtdCategoryType = "abortion"

    UtdCategoryType_abused_drugs UtdCategoryType = "abused-drugs"

    UtdCategoryType_adult_and_pornography UtdCategoryType = "adult-and-pornography"

    UtdCategoryType_alcohol_and_tobacco UtdCategoryType = "alcohol-and-tobacco"

    UtdCategoryType_auctions UtdCategoryType = "auctions"

    UtdCategoryType_bot_nets UtdCategoryType = "bot-nets"

    UtdCategoryType_business_and_economy UtdCategoryType = "business-and-economy"

    UtdCategoryType_cdns UtdCategoryType = "cdns"

    UtdCategoryType_cheating UtdCategoryType = "cheating"

    UtdCategoryType_computer_and_internet_info UtdCategoryType = "computer-and-internet-info"

    UtdCategoryType_computer_and_internet_security UtdCategoryType = "computer-and-internet-security"

    UtdCategoryType_confirmed_spam_sources UtdCategoryType = "confirmed-spam-sources"

    UtdCategoryType_cult_and_occult UtdCategoryType = "cult-and-occult"

    UtdCategoryType_dating UtdCategoryType = "dating"

    UtdCategoryType_dead_sites UtdCategoryType = "dead-sites"

    UtdCategoryType_dynamic_content UtdCategoryType = "dynamic-content"

    UtdCategoryType_educational_institutions UtdCategoryType = "educational-institutions"

    UtdCategoryType_entertainment_and_arts UtdCategoryType = "entertainment-and-arts"

    UtdCategoryType_fashion_and_beauty UtdCategoryType = "fashion-and-beauty"

    UtdCategoryType_financial_services UtdCategoryType = "financial-services"

    UtdCategoryType_gambling UtdCategoryType = "gambling"

    UtdCategoryType_games UtdCategoryType = "games"

    UtdCategoryType_government UtdCategoryType = "government"

    UtdCategoryType_gross UtdCategoryType = "gross"

    UtdCategoryType_hacking UtdCategoryType = "hacking"

    UtdCategoryType_hate_and_racism UtdCategoryType = "hate-and-racism"

    UtdCategoryType_health_and_medicine UtdCategoryType = "health-and-medicine"

    UtdCategoryType_home UtdCategoryType = "home"

    UtdCategoryType_hunting_and_fishing UtdCategoryType = "hunting-and-fishing"

    UtdCategoryType_illegal UtdCategoryType = "illegal"

    UtdCategoryType_image_and_video_search UtdCategoryType = "image-and-video-search"

    UtdCategoryType_individual_stock_advice_and_tools UtdCategoryType = "individual-stock-advice-and-tools"

    UtdCategoryType_internet_communications UtdCategoryType = "internet-communications"

    UtdCategoryType_internet_portals UtdCategoryType = "internet-portals"

    UtdCategoryType_job_search UtdCategoryType = "job-search"

    UtdCategoryType_keyloggers_and_monitoring UtdCategoryType = "keyloggers-and-monitoring"

    UtdCategoryType_kids UtdCategoryType = "kids"

    UtdCategoryType_legal UtdCategoryType = "legal"

    UtdCategoryType_local_information UtdCategoryType = "local-information"

    UtdCategoryType_malware_sites UtdCategoryType = "malware-sites"

    UtdCategoryType_marijuana UtdCategoryType = "marijuana"

    UtdCategoryType_military UtdCategoryType = "military"

    UtdCategoryType_motor_vehicles UtdCategoryType = "motor-vehicles"

    UtdCategoryType_music UtdCategoryType = "music"

    UtdCategoryType_news_and_media UtdCategoryType = "news-and-media"

    UtdCategoryType_nudity UtdCategoryType = "nudity"

    UtdCategoryType_online_greeting_cards UtdCategoryType = "online-greeting-cards"

    UtdCategoryType_online_personal_storage UtdCategoryType = "online-personal-storage"

    UtdCategoryType_open_http_proxies UtdCategoryType = "open-http-proxies"

    UtdCategoryType_p2p UtdCategoryType = "p2p"

    UtdCategoryType_parked_sites UtdCategoryType = "parked-sites"

    UtdCategoryType_pay_to_surf UtdCategoryType = "pay-to-surf"

    UtdCategoryType_personal_sites_and_blogs UtdCategoryType = "personal-sites-and-blogs"

    UtdCategoryType_philosophy_and_political_advocacy UtdCategoryType = "philosophy-and-political-advocacy"

    UtdCategoryType_phishing_and_other_frauds UtdCategoryType = "phishing-and-other-frauds"

    UtdCategoryType_private_ip_addresses UtdCategoryType = "private-ip-addresses"

    UtdCategoryType_proxy_avoid_and_anonymizers UtdCategoryType = "proxy-avoid-and-anonymizers"

    UtdCategoryType_questionable UtdCategoryType = "questionable"

    UtdCategoryType_real_estate UtdCategoryType = "real-estate"

    UtdCategoryType_recreation_and_hobbies UtdCategoryType = "recreation-and-hobbies"

    UtdCategoryType_reference_and_research UtdCategoryType = "reference-and-research"

    UtdCategoryType_religion UtdCategoryType = "religion"

    UtdCategoryType_search_engines UtdCategoryType = "search-engines"

    UtdCategoryType_sex_education UtdCategoryType = "sex-education"

    UtdCategoryType_shareware_and_freeware UtdCategoryType = "shareware-and-freeware"

    UtdCategoryType_shopping UtdCategoryType = "shopping"

    UtdCategoryType_social_network UtdCategoryType = "social-network"

    UtdCategoryType_society UtdCategoryType = "society"

    UtdCategoryType_spam_urls UtdCategoryType = "spam-urls"

    UtdCategoryType_sports UtdCategoryType = "sports"

    UtdCategoryType_spyware_and_adware UtdCategoryType = "spyware-and-adware"

    UtdCategoryType_streaming_media UtdCategoryType = "streaming-media"

    UtdCategoryType_swimsuits_and_intimate_apparel UtdCategoryType = "swimsuits-and-intimate-apparel"

    UtdCategoryType_training_and_tools UtdCategoryType = "training-and-tools"

    UtdCategoryType_translation UtdCategoryType = "translation"

    UtdCategoryType_travel UtdCategoryType = "travel"

    UtdCategoryType_uncategorized UtdCategoryType = "uncategorized"

    UtdCategoryType_unconfirmed_spam_sources UtdCategoryType = "unconfirmed-spam-sources"

    UtdCategoryType_violence UtdCategoryType = "violence"

    UtdCategoryType_weapons UtdCategoryType = "weapons"

    UtdCategoryType_web_advertisements UtdCategoryType = "web-advertisements"

    UtdCategoryType_web_based_email UtdCategoryType = "web-based-email"

    UtdCategoryType_web_hosting UtdCategoryType = "web-hosting"
)

