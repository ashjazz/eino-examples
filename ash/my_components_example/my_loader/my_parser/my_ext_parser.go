package my_parser

import (
	"context"
	"github.com/cloudwego/eino-ext/components/document/parser/docx"
	"github.com/cloudwego/eino-ext/components/document/parser/html"
	"github.com/cloudwego/eino-ext/components/document/parser/pdf"
	"github.com/cloudwego/eino/components/document/parser"
)

func NewExtParser(ctx context.Context) (*parser.ExtParser, error) {
	textParser := parser.TextParser{}

	htmlSelector := "body"
	htmlParser, err := html.NewParser(ctx, &html.Config{
		Selector: &htmlSelector,
	})
	docxParser, err := docx.NewDocxParser(ctx, &docx.Config{
		ToSections:      true,
		IncludeHeaders:  false,
		IncludeComments: false,
		IncludeFooters:  false,
		IncludeTables:   false,
	})
	pdfParser, err := pdf.NewPDFParser(ctx, &pdf.Config{})
	if err != nil {
		return nil, err
	}

	// 创建扩展解析器
	extParser, err := parser.NewExtParser(ctx, &parser.ExtParserConfig{
		// 注册扩展名解析器worker
		Parsers: map[string]parser.Parser{
			".docx": docxParser,
			".pdf":  pdfParser,
			".html": htmlParser,
		},
		FallbackParser: textParser,
	})
	if err != nil {
		return nil, err
	}
	return extParser, nil
}
