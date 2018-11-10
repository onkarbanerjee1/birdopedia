package bird

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

// Builder helps to build a bird
type Builder struct {
	genericName,
	commonName,
	scientificName,
	pictureURL string
	habitat    []string
	endangered bool
	postedBy   string
}

// NewBuilder returns a new Builder
func NewBuilder(genericName string) *Builder {
	return &Builder{genericName: genericName}
}

// CommonName sets the common name
func (b *Builder) CommonName(commonName string) *Builder {
	b.commonName = commonName
	return b
}

// ScientificName sets the ScientificName
func (b *Builder) ScientificName(scientificName string) *Builder {
	b.scientificName = scientificName
	return b
}

// Habitat sets the Habitat
func (b *Builder) Habitat(habitat []string) *Builder {
	b.habitat = habitat
	return b
}

// Endangered sets the endangered
func (b *Builder) Endangered(endangered bool) *Builder {
	b.endangered = endangered
	return b
}

// PostedBy sets the postedBy
func (b *Builder) PostedBy(postedBy string) *Builder {
	b.postedBy = postedBy
	return b
}

// Build returns a bird
func (b *Builder) Build() *Bird {
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
