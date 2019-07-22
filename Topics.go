package ral

import (
	"strconv"
	"fmt"
	"encoding/json"

	"github.com/eidolon/wordwrap"
)

// Topics are posted to a certain [Continuity/Year]
// Example: [Anime/2019/1] is the first Topic posted on
// the Anime continuity this year
type TopicList []Topic
type Topic struct {
	Id int
	Topic int
	Created string
	Continuity string
	Content string
	Replies int
	Year int }

// Generate and execute an API request which returns the full list of
// Topics on the given Site and given Continuity for the given Year
func (s Site) Topics(continuity string, year int) (ret TopicList, err error) {
	params := map[string]string{
		"continuity": continuity,
		"year": strconv.Itoa(year)}
	areq := s.APIRequest(View, params)

	body, err := areq.Go()
	if err != nil { return }

	err = json.Unmarshal(body, &ret)
	return
}
// Serialize to console
func (tl TopicList) Print(f Format) {
	switch(f) {
	case FormatSimple:
		for _, t := range tl {
			fmt.Printf("[%s/%d/%d] (%s) (%d replies)\n",
				t.Continuity,
				t.Year,
				t.Topic,
				t.Created,
				t.Replies)

			wrapper := wordwrap.Wrapper(76, false)
			fmt.Printf("%s\n", wordwrap.Indent(wrapper(t.Content), "    ", true))
		}
	} }
