package npmdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall_error_bad_request(t *testing.T) {

	_, err := call("/point/asdf/express")

	assert.Error(t, err)

	e := err.(*Error)
	assert.Equal(t, e.Error(), "invalid date")
	assert.Equal(t, "Bad Request", e.Status)
	assert.Equal(t, 400, e.StatusCode)
}

func TestGetPointCounts_no_period_no_package(t *testing.T) {

	out, err := GetPointCounts("", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Empty(t, out.Package)
}

func TestGetPointCounts_no_period(t *testing.T) {

	out, err := GetPointCounts("", "express")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Equal(t, "express", out.Package)
}

func TestGetPointCounts_no_package(t *testing.T) {

	out, err := GetPointCounts("last-week", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Empty(t, out.Package)
}

func TestGetPointCounts(t *testing.T) {

	out, err := GetPointCounts("last-month", "express")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Equal(t, "express", out.Package)
}

func TestGetRangeCounts_no_period_no_package(t *testing.T) {

	out, err := GetRangeCounts("", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Empty(t, out.Package)
}

func TestGetRangeCounts_no_period(t *testing.T) {

	out, err := GetRangeCounts("", "express")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Equal(t, "express", out.Package)
}

func TestGetRangeCounts_no_package(t *testing.T) {

	out, err := GetRangeCounts("last-week", "")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Empty(t, out.Package)
}

func TestGetRangeCounts(t *testing.T) {

	out, err := GetRangeCounts("last-week", "express")

	assert.NoError(t, err)
	assert.NotEmpty(t, out.Downloads)
	assert.NotEmpty(t, out.Start)
	assert.NotEmpty(t, out.End)
	assert.Equal(t, "express", out.Package)
}
