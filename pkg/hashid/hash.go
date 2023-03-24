package hashid

import (
	"encoding/json"
	"log"

	"github.com/speps/go-hashids/v2"
)

var (
	hashid     *hashids.HashID
	emptySlice []uint
)

func init() {
	var err error
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 6
	hashid, err = hashids.NewWithData(hd)
	if err != nil {
		log.Fatal(err)
	}
}

type HashID struct {
	slice   []uint
	encoded string
}

func New(id uint) (*HashID, error) {
	h := new(HashID)
	err := h.SetID(id)
	return h, err
}

func FromString(encoded string) *HashID {
	h := new(HashID)
	h.slice = toUint(hashid.Decode(encoded))
	h.encoded = encoded
	return h
}

func (h *HashID) ID() uint {
	return uint(h.slice[0])
}

func (h *HashID) Tags() []uint {
	if len(h.slice) > 1 {
		return h.slice[1:]
	}

	return emptySlice
}

func (h *HashID) String() string {
	return h.encoded
}

func (h *HashID) SetTags(tags ...uint) error {
	h.slice = append(h.slice[0:1], tags...)
	return h.encode()

}
func (h *HashID) SetID(id uint) error {
	if len(h.slice) == 0 {
		h.slice = []uint{0}
	}

	h.slice[0] = id

	return h.encode()
}

func (h *HashID) encode() error {
	var err error

	h.encoded, err = hashid.Encode(toInt(h.slice))

	return err
}

func (h *HashID) MarshalJSON() ([]byte, error) {
	if len(h.slice) == 0 {
		return []byte(""), nil
	}

	return json.Marshal(h.encoded)
}

func (h *HashID) UnmarshalJSON(b []byte) error {
	h = FromString(string(b))

	return nil
}

func toUint(sliceInt []int) []uint {
	var sliceUint []uint

	for _, v := range sliceInt {
		sliceUint = append(sliceUint, uint(v))
	}

	return sliceUint
}

func toInt(sliceUint []uint) []int {
	var sliceInt []int

	for _, v := range sliceUint {
		sliceInt = append(sliceInt, int(v))
	}

	return sliceInt
}
