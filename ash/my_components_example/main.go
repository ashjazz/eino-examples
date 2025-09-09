package main

import (
	"context"
	"example.com/m/v2/my_loader"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	//embedder, err := my_embedding.NewEmbeddingModel(ctx)
	//if err != nil {
	//	log.Fatalf("embedder create failed: %v", err)
	//}
	docs, err := my_loader.LoadFile(ctx)
	if err != nil {
		log.Fatalf("doc load failed: %v", err)
	}
	if len(docs) == 0 {
		log.Println("docs load failed: no docs content")
		return
	}
	fmt.Println("=============== doc loads result ================")
	for i, doc := range docs {
		fmt.Printf("doc seq_%d: %s \n", i, doc.Content)
	}
}
