package user

import (
	"testing"

	gomob "github.com/iranika/gomob"
	"github.com/stretchr/testify/assert"
)

func TestDLSite(t *testing.T) {

	maniaxWorkLinks := [...]string{
		"https://www.dlsite.com/maniax/work/=/product_id/RJ303351.html",
		"https://www.dlsite.com/maniax/work/=/product_id/RJ220032.html",
	}
	for _, link := range maniaxWorkLinks {
		result := gomob.WhatisLinkPattern(link)
		assert.Equal(t, result, gomob.ManiaxWork)
		t.Logf("link: %s , result: %d", link, result)
	}

	maniaxAffiliatelLinks := [...]string{
		"https://www.dlsite.com/maniax/dlaf/=/t/s/link/work/aid/iranica/id/RJ303351.html",
		"https://www.dlsite.com/maniax/dlaf/=/t/s/link/work/aid/iranica/id/RJ314617.html",
	}
	for _, link := range maniaxAffiliatelLinks {
		result := gomob.WhatisLinkPattern(link)
		assert.Equal(t, result, gomob.ManiaxAffiliate)
		t.Logf("link: %s , result: %d", link, result)
	}

	link := "https://www.dlsite.com/maniax/work/=/product_id/RJ312136.html"
	aflink := gomob.ReplaceAffiliate(link)
	assert.Equal(t, aflink, "https://www.dlsite.com/maniax/dlaf/=/t/s/link/work/aid/iranica/id/RJ312136.html")
	t.Logf("link: %s , result: %s", link, aflink)

	link = "https://www.dlsite.com/maniax/dlaf/=/t/s/link/work/aid/iranica/id/RJ312136.html"
	aflink = gomob.ReplaceWork(link)
	assert.Equal(t, aflink, "https://www.dlsite.com/maniax/work/=/product_id/RJ312136.html")
	t.Logf("link: %s , result: %s", link, aflink)

}
