package my_transformer

import (
	"context"
	"errors"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/recursive"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

func IdGeneratorFunc(ctx context.Context, originId string, splitIndex int) string {
	fileUUID := uuid.New().String()
	if originId == "" {
		originId = fileUUID
	}
	return strings.Join([]string{originId, "_", strconv.Itoa(splitIndex)}, "")
}

func SplitChildSegment(ctx context.Context, docs []*schema.Document, parentIdentifier string, childIdentifier string) ([]*schema.Document, error) {
	splitter, err := recursive.NewSplitter(ctx, &recursive.Config{
		ChunkSize:   1,
		OverlapSize: 0,
		Separators:  []string{parentIdentifier},
		LenFunc:     nil,
		KeepType:    recursive.KeepTypeNone,
		IDGenerator: IdGeneratorFunc,
	})

	if err != nil {
		return nil, err
	}
	parentDocs, err := splitter.Transform(ctx, docs)
	if err != nil {
		return nil, err
	}

	childSplitter, err := recursive.NewSplitter(ctx, &recursive.Config{
		ChunkSize:   1,
		OverlapSize: 0,
		Separators:  []string{childIdentifier},
		LenFunc:     nil,
		KeepType:    recursive.KeepTypeNone,
		IDGenerator: IdGeneratorFunc,
	})
	for _, parentDoc := range parentDocs {
		docs, err := childSplitter.Transform(ctx, []*schema.Document{parentDoc})
		if err != nil {
			return nil, err
		}
		if len(docs) == 0 {
			return nil, errors.New("文档子切片解析错误，请检查文档")
		}
		if parentDoc.MetaData == nil {
			parentDoc.MetaData = make(map[string]any)
		}
		parentDoc.MetaData["children"] = docs
	}
	return parentDocs, nil
}
