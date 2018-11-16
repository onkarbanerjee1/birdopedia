package models

import (
	"fmt"
)

// Bird denotes a bird
type Bird struct {
	GenericName,
	commonName,
	scientificName,
	pictureURL string
	habitat    []string
	endangered bool
	postedBy   string
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

// BirdBuilder returns a new BirdBuilder
func NewBirdBuilder(genericName string) *BirdBuilder {
	return &BirdBuilder{genericName: genericName}
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
		GenericName:    b.genericName,
		commonName:     b.commonName,
		scientificName: b.scientificName,
		habitat:        b.habitat,
		endangered:     b.endangered,
		postedBy:       b.postedBy,
	}
}

func (bird *Bird) String() string {
	return fmt.Sprintf(
		"Generic Name : %s,	 Common Name : %s, Scientific Name : %s, Habitat : %s, Endangered %t, Posted By %s\n",
		bird.GenericName, bird.commonName, bird.scientificName, bird.habitat, bird.endangered, bird.postedBy,
	)
}
