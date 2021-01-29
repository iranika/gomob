package gomob

import "strings"

const AffiliateId = "iranica"
const maniax_work = "https://www.dlsite.com/maniax/work/=/product_id/"
const maniax_dlaf = "https://www.dlsite.com/maniax/dlaf/=/t/s/link/work/aid/" + AffiliateId + "/id/"
const home_work = "https://www.dlsite.com/home/work/=/product_id/"
const home_dlaf = "https://www.dlsite.com/home/dlaf/=/t/s/link/work/aid/" + AffiliateId + "/id/"

const (
	ManiaxWork int = iota
	ManiaxAffiliate
	HomeWork
	HomeAffiliate
	Unknown
)

func WhatisLinkPattern(urlstr string) int {
	if strings.Contains(urlstr, "/maniax/work/") {
		return ManiaxWork
	} else if strings.Contains(urlstr, "/maniax/dlaf/") {
		return ManiaxAffiliate
	} else if strings.Contains(urlstr, "/home/work/") {
		return HomeWork
	} else if strings.Contains(urlstr, "/home/dlaf/") {
		return HomeAffiliate
	} else {
		return Unknown
	}
}

func ReplaceWork(urlstr string) string {
	switch WhatisLinkPattern(urlstr) {
	case ManiaxWork:
		return urlstr
	case HomeWork:
		return urlstr
	case ManiaxAffiliate:
		code := GetProductCode(urlstr)
		return maniax_work + code + ".html"
	case HomeAffiliate:
		code := GetProductCode(urlstr)
		return home_work + code + ".html"
	default:
		return "undefined"
	}
}

func ReplaceAffiliate(urlstr string) string {
	switch WhatisLinkPattern(urlstr) {
	case ManiaxAffiliate:
		return urlstr
	case HomeAffiliate:
		return urlstr
	case ManiaxWork:
		code := GetProductCode(urlstr)
		return maniax_dlaf + code + ".html"
	case HomeWork:
		code := GetProductCode(urlstr)
		return home_dlaf + code + ".html"
	default:
		return "undefined"
	}
}
