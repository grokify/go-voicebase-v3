package clientutil

import (
	"strings"

	"github.com/grokify/go-voicebase-v3/client"
)

// BuildVocabulariesForStrings takes a list of strings and converts them
// to a list of `VbVocabulary` structs for use in uploading media.
func BuildVocabulariesForStrings(vocabNames ...string) []voicebase.VbVocabulary {
	vocabs := []voicebase.VbVocabulary{}
	for _, vocabName := range vocabNames {
		vocabName = strings.TrimSpace(vocabName)
		if len(vocabName) > 0 {
			vocabs = append(vocabs, voicebase.VbVocabulary{
				VocabularyName: vocabName,
			})
		}
	}
	return vocabs
}

type Vocabulary struct {
	VocabularyName string
	Term           string
	SoundsLike     []string
}

func (v *Vocabulary) ToVbVocabulary() voicebase.VbVocabulary {
	return voicebase.VbVocabulary{
		VocabularyName: v.VocabularyName,
		VocabularyType: voicebase.TERMS,
		Terms: []voicebase.VbVocabularyTerm{
			{
				Term:       v.Term,
				SoundsLike: v.SoundsLike,
			},
		},
	}
}
