package pb

import (
	"fmt"
	"strings"

	"github.com/xuender/kit/types"
)

func (p *Hit) Ack() string {
	return fmt.Sprintf("%d %d", p.GetCol(), p.GetLen())
}

func NewHits(text string) []*Hit {
	end := strings.Index(text, ":")
	if end < 0 {
		end = len(text)
	}

	start := strings.Index(text, ";")
	if start > end || start < 0 {
		start = 0
	}

	cols := strings.Split(text[start+1:end], ",")
	hits := make([]*Hit, len(cols))

	for idx, col := range cols {
		strs := strings.Split(col, " ")

		col, err := types.ParseInteger[uint32](strs[0])
		if err != nil {
			continue
		}

		length, err := types.ParseInteger[uint32](strs[1])
		if err != nil {
			continue
		}

		hits[idx] = &Hit{
			Col: col,
			Len: length,
		}
	}

	return hits
}
