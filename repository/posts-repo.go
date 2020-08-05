package repository

import (
	"context"
	"go-firebase/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

//NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	//projectId      string = "go-firebase"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firestore-private-key.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firesbase App: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to adding a new post...: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./firestore-private-key.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}
