package litxapfwew

import (
	"bytes"
	"log"
	"slices"
	"strings"
	"sync"

	fwewlib "github.com/fwew/fwew-lib/v5"
	"github.com/gissleh/litxap"
	"github.com/gissleh/litxap/litxaputil"
)

type fwewDict struct{}

func (d *fwewDict) LookupEntries(word string) ([]litxap.Entry, error) {
	res, err := fwewlib.TranslateFromNaviHash(word, true, false, true)
	if err != nil {
		return nil, err
	}

	entries := make([]litxap.Entry, 0, len(res))

	for _, matches := range res {
		for _, match := range matches {
			if match.ID == "" {
				continue
			}

			syllables := strings.Split(strings.ReplaceAll(strings.ToLower(match.Syllables), " ", "-"), "-")

			for _, ipa := range strings.Split(match.IPA, "or") {
				ipa = strings.Trim(ipa, " []")
				ipaSyllables := strings.Split(strings.ReplaceAll(ipa, " ", "."), ".")
				if len(ipaSyllables) != len(syllables) {
					continue
				}

				stressIndex := 0
				for i, syllable := range ipaSyllables {
					if strings.HasPrefix(syllable, "ˈ") {
						stressIndex = i
						break
					}
				}

				suffixes := append([]string{}, match.Affixes.Suffix...)

				slices.Reverse(suffixes)

				entry := litxap.Entry{
					Word:        match.Navi,
					Translation: match.EN,
					Syllables:   syllables,
					Stress:      stressIndex,
					InfixPos:    litxaputil.InfixPositionsFromBrackets(strings.ReplaceAll(match.InfixLocations, " ", ""), syllables),
					Prefixes:    match.Affixes.Prefix,
					Infixes:     match.Affixes.Infix,
					Suffixes:    suffixes,
				}

				entries = append(entries, entry)
			}
		}
	}

	return entries, nil
}

func Adpositions() ([]string, error) {
	list, err := fwewlib.List([]string{"pos", "has", "adp."}, 0)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(list))
	for _, match := range list {
		res = append(res, strings.TrimSuffix(match.Navi, "+"))
	}

	return res, nil
}

func FindMultipartWords() map[string]string {
	// Calculate the multiword words needed at startup
	// Make sure we have words that must be multiword words
	doubles := map[string]string{}
	multis := fwewlib.GetMultiwordWords()
	fullWord := bytes.NewBuffer(make([]byte, 0, 16))
	IPAstring := []string{}
	for key, val := range multis {
		for _, stringArray := range val {
			fullWord.Reset()
			fullWord.WriteString(key + " ")
			for i, multiword := range stringArray {
				fullWord.WriteString(multiword)
				if i+1 != len(stringArray) {
					fullWord.WriteString(" ")
				}
			}
			result1, _ := fwewlib.TranslateFromNaviHash(key, true, false, true)
			result2, _ := fwewlib.TranslateFromNaviHash(fullWord.String(), true, false, true)
			if len(result2[0]) == 1 {
				log.Println(fullWord.String(), "-- not found")
				continue
			}

			IPAstring = strings.Split(result2[0][1].IPA, " ")

			if len(result1[0]) < 2 {
				doubles[key] = IPAstring[0]
			}

			for i, multiword := range stringArray {
				res3, _ := fwewlib.TranslateFromNaviHash(multiword, true, false, true)
				if len(res3[0]) < 2 {
					doubles[multiword] = IPAstring[i+1]
				}
			}
		}
	}

	return doubles
}

// MultiWordPartDictionary makes a dictionary using litxap.CustomWords for any parts of a multipart word
// that don't already exist as separate words, like "tsaheyl".
func MultiWordPartDictionary() litxap.Dictionary {
	names := make([]string, 0, 16)
	for _, ipa := range FindMultipartWords() {
		wordOptions, stressOptions := litxaputil.RomanizeIPA(ipa)
		for i := range wordOptions {
			words := wordOptions[i]
			stress := stressOptions[i]

			if len(words) == 1 && len(stress) == 1 {
				syllables := words[0]
				stressIndex := -1
				for _, index := range stress {
					if index != -1 {
						stressIndex = index
						break
					}
				}

				if stressIndex == -1 {
					syllables[0] = "-" + syllables[0]
				} else {
					for i := range syllables {
						if i == stressIndex {
							syllables[i] = "*" + syllables[i]
						}
					}
				}

				names = append(names, strings.Join(syllables, "-"))
			}
		}
	}

	res := litxap.CustomWords(names, "Part of multi-part word")
	return res
}

var startEverythingOnce sync.Once

// Global returns a wrapper around fwew that implements the litxap.Dictionary interface.
// It does use the global `fwew` dictionary and will initialize it on the first call.
func Global() litxap.Dictionary {
	startEverythingOnce.Do(func() {
		fwewlib.StartEverything()
	})

	return &fwewDict{}
}
