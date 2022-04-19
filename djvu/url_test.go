package djvu

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UrlTestSuite struct {
	suite.Suite
	SimpleUrl *Url // `http://www.lizardtech.com/file%201.djvu`
}

func TestUrlSuite(t *testing.T) {
	suite.Run(t, new(UrlTestSuite))
}

func (s *UrlTestSuite) SetupTest() {
	var err error
	s.SimpleUrl, err = NewUrl(`http://www.lizardtech.com/file%201.djvu`)
	s.NoError(err)
	s.False(s.SimpleUrl.IsEmpty())
}

func (s *UrlTestSuite) TestPathname() {
	s.Equal(`/file%201.djvu`, s.SimpleUrl.Pathname())
}

func (s *UrlTestSuite) TestUrlName() {
	s.Equal(`file%201.djvu`, s.SimpleUrl.Name())
}
