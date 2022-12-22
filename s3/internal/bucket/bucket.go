package bucket

import (
	"context"
	"os"
	"s3/internal/pkg/cloud"

	"github.com/sirupsen/logrus"
)

func Bucket(client cloud.BucketClient) {
	ctx := context.Background()

	if err := create(ctx, client); err != nil {
		logrus.Printf("creating bucket error : %v", err)
	} else {
		logrus.Println("bucket created")
	}

	url, err := uploadObject(ctx, client)
	if err != nil {
		logrus.Printf("uploading file error : %v", err)
	} else {
		logrus.Printf("file uploaded: %s", url)
	}

	if err = downloadObject(ctx, client); err != nil {
		logrus.Printf("downloading file error: %v", err)
	} else {
		logrus.Println("file downloaded")
	}

	if err = deleteObject(ctx, client); err != nil {
		logrus.Printf("deleted file error: %v", err)
	} else {
		logrus.Println("file deleted")
	}

	list, err := listObjects(ctx, client)
	if err != nil {
		logrus.Println("getting list error: %v", list)
	} else {
		logrus.Printf("objects list: %v", list)
	}
}

func create(ctx context.Context, client cloud.BucketClient) error {
	if err := client.Create(ctx, "aws-test"); err != nil {
		return err
	}

	return nil
}

func uploadObject(ctx context.Context, client cloud.BucketClient) (string, error) {
	file, err := os.Open("./assets/id.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	url, err := client.UploadObject(ctx, "aws-test", "id.txt", file)
	if err != nil {
		return "", err
	}

	return url, nil
}

func downloadObject(ctx context.Context, client cloud.BucketClient) error {
	file, err := os.Create("./tmp/id.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := client.DownloadObject(ctx, "aws-test", "id.txt", file); err != nil {
		return err
	}

	return nil
}

func deleteObject(ctx context.Context, client cloud.BucketClient) error {
	if err := client.DeleteObject(ctx, "aws-test", "id.txt"); err != nil {
		return err
	}

	return nil
}

func listObjects(ctx context.Context, client cloud.BucketClient) ([]*cloud.Object, error) {
	objects, err := client.ListObjects(ctx, "aws-test")
	if err != nil {
		return nil, err
	}

	return objects, nil
}