package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/fajarslvn/go_rest_api/entity"
	"google.golang.org/api/iterator"
)

type repo struct {}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId		string = "pragmatic-reviews-ee857"
	collectionName 	string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to created firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":		post.ID,
		"Title":	post.Title,
		"Text": 	post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (repo) FindAll() ([]entity.Post, error ) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to created firestore client: %v", err)
		return nil, err
	}

	defer client.Close()
	
	var posts []entity.Post
	itr := client.Collection(collectionName).Documents(ctx)
	
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterated the list of posts: %v", err)
			return nil, err
		}
		
		post := entity.Post{
			ID: 	doc.Data()["ID"].(int64),
			Title: 	doc.Data()["Title"].(string),
			Text: 	doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
