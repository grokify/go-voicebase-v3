package main

import (
	"fmt"

	vb "github.com/grokify/go-voicebase-v3"
	"github.com/grokify/gotilla/fmt/fmtutil"
)

func CustomVocabRequestForString(name, s string) vb.VbVocabulary {
	return vb.VbVocabulary{
		VocabularyName: name,
		Terms:          []vb.VbVocabularyTerm{TermForString(s)},
	}
}

func TermForString(s string) vb.VbVocabularyTerm {
	return vb.VbVocabularyTerm{Term: s}
}

func main() {
	name := "E16505626570"
	str := "Quentyn Blackwood"

	vocabReq := CustomVocabRequestForString(name, str)
	fmtutil.PrintJSON(vocabReq)

	cfg := vb.VbConfiguration{
		Vocabularies: []vb.VbVocabularyConfiguration{
			{
				VocabularyName: name,
			},
		},
	}

	fmt.Println("DONE")
}
