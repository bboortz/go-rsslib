package rsslib

import (
	//"github.com/davecgh/go-spew/spew"
	//	"github.com/bboortz/go-utils"
	"github.com/stretchr/testify/assert"
	//"go-rsslib"
	"testing"
)

func TestCreateRssItem1(t *testing.T) {
	assert := assert.New(t)

	result := RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000001", Channel: "TestChannel", Title: "testtitle1", Link: "http://localhost"}

	assert.NotNil(result)
	assert.NotEmpty(result.Uuid)
	assert.NotEmpty(result.Channel)
	assert.NotEmpty(result.Title)
	assert.NotEmpty(result.Link)
	assert.Empty(result.Description)
	assert.Empty(result.PublishDate)
	assert.Empty(result.UpdateDate)
	assert.Equal(0, result.UpdateCount)
}

func TestCreateRssItems(t *testing.T) {
	assert := assert.New(t)

	testItemArr := [...]RssItem{
		RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000001", Channel: "TestChannel", Title: "testtitle1", Link: "http://localhost"},
		RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000002", Channel: "TestChannel", Title: "testtitle2", Link: "http://localhost"},
		RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000003", Channel: "TestChannel", Title: "testtitle3", Link: "http://localhost", PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC"},
		RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000004", Channel: "TestChannel", Title: "testtitle4", Link: "http://localhost", PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC", UpdateDate: "2017-01-03 00:10:35.180321993 +0000 UTC"},
		RssItem{Uuid: "68e9e42d-a0ba-5a4c-0001-000000000005", Channel: "TestChannel", Title: "testtitle4", Link: "http://localhost", PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC", UpdateDate: "2017-01-03 00:10:35.180321993 +0000 UTC", UpdateCount: 8},
	}

	for _, v := range testItemArr {
		assert.NotNil(v)
	}
}

func TestCompareAllFieldsTrue(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(true, result)
}

func TestCompareAllFieldsFalseWithId(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          22,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithUuid(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000002",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithChannel(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel2",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithTitle(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle12",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithLink(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost/",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithDescription(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description2",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithThumbnail(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb2.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithPublishDate(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321994 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithUpdateDate(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321994 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestCompareAllFieldsFalseWithUpdateCount(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 6,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(false, result)
}

func TestDiffTrue(t *testing.T) {
	assert := assert.New(t)

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	result := item1.CompareAllFields(item2)

	assert.Equal(true, result)
}

func TestDiffFalseWithId(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          22,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithUuid(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000002",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithChannel(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel2",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithTitle(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle12",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithLink(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost/",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithDescription(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description2",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithThumbnail(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb2.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithPublishDate(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321994 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithUpdateDate(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321994 +0000 UTC",
		UpdateCount: 5,
	}

	item1.Diff(item2)
}

func TestDiffFalseWithUpdateCount(t *testing.T) {

	item1 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 5,
	}
	item2 := RssItem{
		Id:          21,
		Uuid:        "68e9e42d-a0ba-5a4c-0002-000000000001",
		Channel:     "TestChannel",
		Title:       "testtitle1",
		Link:        "http://localhost",
		Description: "Test Description",
		Thumbnail:   "http://localhost/thumb.png",
		PublishDate: "2017-01-03 00:06:35.180321993 +0000 UTC",
		UpdateDate:  "2017-01-03 00:10:35.180321993 +0000 UTC",
		UpdateCount: 6,
	}

	item1.Diff(item2)
}
