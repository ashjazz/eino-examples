package my_loader

import (
	"context"
	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/schema"
)

func LoadFile(ctx context.Context) ([]*schema.Document, error) {
	loader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		UseNameAsID: true,
	})

	if err != nil {
		return nil, err
	}

	filePath := "./test_file/test.docx"
	docs, err := loader.Load(ctx, document.Source{
		URI: filePath,
	})
	if err != nil {
		return nil, err
	}
	return docs, nil
}
