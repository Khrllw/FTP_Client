package client

import (
	"context"
	"encoding/base64"
	"log"
	"testing"

	"ftp_service/internal/domain/models"
)

const (
	testHost       = "192.168.30.92"
	testPort       = "21"
	testFileName   = "example.txt"
	rawContent     = "Hello, FTP!"
	testTargetPath = "upload/"
)

var testFileContent = base64.StdEncoding.EncodeToString([]byte(rawContent))

func TestSendAndGetFileWithAuth(t *testing.T) {
	api := NewClient("http://localhost:8080")
	ctx := context.Background()
	username := "tester"
	password := "password"
	fullFilename := testTargetPath + testFileName

	// Отправка файла с авторизацией
	log.Println("Отправка файла с авторизацией...")
	sendReq := models.SendFileRequest{
		Host:          testHost,
		Port:          testPort,
		Username:      username,
		Password:      password,
		Filename:      testFileName,
		ContentBase64: testFileContent,
		TargetPath:    testTargetPath,
	}
	sendResp, _, err := api.SendFileWithAuth(ctx, &sendReq)
	if err != nil {
		t.Fatalf("Ошибка при отправке файла с авторизацией: %v", err)
	}
	if sendResp.Status != "success" {
		t.Fatalf("Ошибка в ответе: %+v", sendResp)
	}
	log.Printf("Файл успешно отправлен с авторизацией: %s\n", sendResp.Message)

	// Получение файла с авторизацией
	log.Println("Получение файла с авторизацией...")
	getReq := models.GetFileWithAuthRequest{
		Host:     testHost,
		Port:     testPort,
		Username: username,
		Password: password,
		FilePath: fullFilename,
	}
	getResp, _, err := api.GetFileWithAuth(ctx, &getReq)
	if err != nil {
		t.Fatalf("Ошибка при получении файла с авторизацией: %v", err)
	}
	if getResp.Data.Filename != testFileName || getResp.Data.ContentBase64 != testFileContent {
		t.Fatalf("Содержимое файла не совпадает: %+v", getResp)
	}
	log.Printf("Файл получен с авторизацией: %s\n", getResp.Data.Filename)
}

func TestSendAndGetFileAnonymous(t *testing.T) {
	api := NewClient("http://localhost:8080")
	ctx := context.Background()
	anonPath := "anon/"
	fullFilename := anonPath + testFileName

	// Отправка файла анонимно
	log.Println("Отправка файла анонимно...")
	sendReq := models.SendFileRequest{
		Host:          testHost,
		Port:          testPort,
		Filename:      testFileName,
		ContentBase64: testFileContent,
		TargetPath:    anonPath,
	}
	sendResp, _, err := api.SendFileAnonymous(ctx, &sendReq)
	if err != nil {
		t.Fatalf("Ошибка при анонимной отправке файла: %v", err)
	}
	if sendResp.Status != "success" {
		t.Fatalf("Ошибка в ответе: %+v", sendResp)
	}
	log.Printf("Файл успешно отправлен анонимно: %s\n", sendResp.Message)

	// Получение файла анонимно
	log.Println("Получение файла анонимно...")
	getReq := models.GetFileAnonymousRequest{
		Host:     testHost,
		Port:     testPort,
		FilePath: fullFilename,
	}
	getResp, _, err := api.GetFileAnonymous(ctx, &getReq)
	if err != nil {
		t.Fatalf("Ошибка при анонимном получении файла: %v", err)
	}
	if getResp.Data.Filename != testFileName || getResp.Data.ContentBase64 != testFileContent {
		t.Fatalf("Содержимое анонимно полученного файла не совпадает: %+v", getResp)
	}
	log.Printf("Файл успешно получен анонимно: %s\n", getResp.Data.Filename)
}
