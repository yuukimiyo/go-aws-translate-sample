package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

func main() {
	sourceText := flag.String("text", "これは翻訳のテストです", "source text")
	sourceLC := flag.String("slc", "ja", "source language code [en|ja|fr]...")
	targetLC := flag.String("tlc", "en", "target language code [en|ja|fr]...")
	flag.Parse()

	sess := session.Must(session.NewSession())
	trs := translate.New(sess)

	result, err := trs.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(*sourceLC),
		TargetLanguageCode: aws.String(*targetLC),
		Text:               aws.String(*sourceText),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf(*result.TranslatedText)
}

