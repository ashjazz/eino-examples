package my_loader

import (
	"context"
	"example.com/m/v2/my_loader/my_parser"
	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/schema"
	"strings"
)

func LoadFile(ctx context.Context) ([]*schema.Document, error) {
	extParser, err := my_parser.NewExtParser(ctx)
	if err != nil {
		return nil, err
	}
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		UseNameAsID: false,
		Parser:      extParser,
	})

	if err != nil {
		return nil, err
	}

	filePath := "./test_file/test.docx"
	docs, err := loader.Load(ctx, document.Source{
		URI: filePath,
	})
	for _, doc := range docs {
		doc.Content = strings.Replace(doc.Content, "=== MAIN CONTENT ===", "", -1)
	}
	if err != nil {
		return nil, err
	}
	return docs, nil
}
