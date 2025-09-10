package main

import (
	"context"
	"example.com/m/v2/my_embedding"
	"example.com/m/v2/my_loader"
	"example.com/m/v2/my_transformer"
	"github.com/cloudwego/eino/schema"
	"log"
)

func main() {
	ctx := context.Background()
	//embedder, err := my_embedding.NewEmbeddingModel(ctx)
	//if err != nil {
	//	log.Fatalf("embedder create failed: %v", err)
	//}

	// 读取文档
	docs, err := my_loader.LoadFile(ctx)
	if err != nil {
		log.Fatalf("doc load failed: %v", err)
	}
	if len(docs) == 0 {
		log.Fatalln("docs load failed: no docs content")
	}

	// 拆分为父子结构
	parentDocsWithChildren, err := my_transformer.SplitChildSegment(ctx, docs, "-----", "\n")

	// 子切片向量化
	embedder, err := my_embedding.NewEmbeddingModel(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, parentDoc := range parentDocsWithChildren {
		for _, childDoc := range parentDoc.MetaData["children"].([]*schema.Document) {
			embeddingRes, err := embedder.EmbedStrings(ctx, []string{childDoc.Content})
			if err != nil {
				log.Fatalln(err)
			}
			childDoc.MetaData["vector"] = embeddingRes[0]
		}
	}
	return
}
