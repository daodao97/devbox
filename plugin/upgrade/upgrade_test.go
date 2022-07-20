package upgrade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	Init("https://github.com/daodao97/EasyClash/releases/download/v1.0/EasyClash-amd-v%d.%d.zip", "0.0")
}

func TestGetCheckVersion(t *testing.T) {
	v := getNeedCheckVersion()
	fmt.Println("new version", v)
}

func TestCheckNewVersion(t *testing.T) {
	have, err := CheckNewVersion()
	assert.Equal(t, nil, err)
	fmt.Println("get new version", have)
}

func TestDonwoad(t *testing.T) {
	process := Download(".", "https://github.com/daodao97/EasyClash/releases/download/v1.0/EasyClash-amd-v1.0.zip")
	for p := range process {
		fmt.Printf("transferred %v / %v bytes error %v \n", p.Complate, p.Size, p.Error)
	}
}
