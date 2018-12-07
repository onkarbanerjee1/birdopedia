package birds

import (
	"fmt"
)

// Bird denotes a bird
type Bird struct {
	CommonName,
	ScientificName,
	PictureURL string
	Habitat    []string
	Endangered bool
	PostedBy   string
}

// BirdBuilder helps to build a bird
type BirdBuilder struct {
	genericName,
	commonName,
	scientificName,
	pictureURL string
	habitat    []string
	endangered bool
	postedBy   string
}

// NewBirdBuilder returns a new BirdBuilder
func NewBirdBuilder(commonName string) *BirdBuilder {
	return &BirdBuilder{commonName: commonName}
}

// CommonName sets the common name
func (b *BirdBuilder) CommonName(commonName string) *BirdBuilder {
	b.commonName = commonName
	return b
}

// ScientificName sets the ScientificName
func (b *BirdBuilder) ScientificName(scientificName string) *BirdBuilder {
	b.scientificName = scientificName
	return b
}

// PictureURL sets the picture URL
func (b *BirdBuilder) PictureURL(picURL string) *BirdBuilder {
	b.pictureURL = picURL
	return b
}

// Habitat sets the Habitat
func (b *BirdBuilder) Habitat(habitat []string) *BirdBuilder {
	b.habitat = habitat
	return b
}

// Endangered sets the endangered
func (b *BirdBuilder) Endangered(endangered bool) *BirdBuilder {
	b.endangered = endangered
	return b
}

// PostedBy sets the postedBy
func (b *BirdBuilder) PostedBy(postedBy string) *BirdBuilder {
	b.postedBy = postedBy
	return b
}

// Build returns a bird
func (b *BirdBuilder) Build() *Bird {
	return &Bird{
		CommonName:     b.commonName,
		ScientificName: b.scientificName,
		Habitat:        b.habitat,
		Endangered:     b.endangered,
		PostedBy:       b.postedBy,
	}
}

func (bird *Bird) String() string {
	return fmt.Sprintf(
		"Common Name : %s, Scientific Name : %s, Habitat : %s, Endangered %t, Posted By %s\n",
		bird.CommonName, bird.ScientificName, bird.Habitat, bird.Endangered, bird.PostedBy,
	)
}
