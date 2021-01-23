# go-aws-translate
Sample code for AWS Translate by golang.

# Required

You need AWS Account and IAM User as attached following policy.

- TranslateFullAccess

# Install

## Environment variables

Set following environment variables.

- AWS_REGION
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

## Download

```
git clone git@github.com:yuukimiyo/go-aws-translate-example.git
```

# Usage

```
Usage of:
  -text string
        source text (default "これは翻訳のテストです")
  -slc string
        source language code [en|ja|fr]... (default "ja")
  -tlc string
        target language code [en|ja|fr]... (default "en")
```

# Example

## Translate to English from Japanese Text.

```sh
go run main.go --text="翻訳のテストです"

> Testing the translation
```

## Translate to Japanese from English Text.

```sh
go run main.go --text="This is test of AWS Translate" --slc="en" --tlc="ja"

> これは AWS 翻訳のテストです
```
