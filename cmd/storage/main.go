package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Função para baixar arquivos do MinIO
func downloadFile(endpoint, accessKey, secretKey, bucketName, objectName, destination string, useSSL bool) error {
	// Configura o cliente MinIO
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("falha ao criar cliente MinIO: %v", err)
	}

	// Abre um arquivo local para escrever o conteúdo
	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("falha ao criar arquivo local: %v", err)
	}
	defer file.Close()

	// Baixa o objeto
	err = client.FGetObject(context.Background(), bucketName, objectName, destination, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("falha ao baixar objeto: %v", err)
	}

	fmt.Printf("Arquivo %s baixado em %s\n", objectName, destination)
	return nil
}

// Função para enviar arquivos ao MinIO
func uploadFile(endpoint, accessKey, secretKey, bucketName, objectName, filePath string, useSSL bool) error {
	// Configura o cliente MinIO
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return fmt.Errorf("falha ao criar cliente MinIO: %v", err)
	}

	// Faz o upload do arquivo
	_, err = client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("falha ao fazer upload: %v", err)
	}

	fmt.Printf("Arquivo %s enviado para o bucket %s como %s\n", filePath, bucketName, objectName)
	return nil
}

func main() {
	// Configurações do MinIO
	endpoint := "br-se1.magaluobjects.com"
	accessKey := "a07868a3-a9db-454c-8064-3c977546d8de"
	secretKey := "38cd2fc9-102a-4b45-abc8-55052e714a8f"
	useSSL := true

	// Exemplo de download
	bucketName := "aulao-cloud"
	objectName := "panda-vermelho.png"
	destination := "./tmp/panda-vermelho.png"

	if err := downloadFile(endpoint, accessKey, secretKey, bucketName, objectName, destination, useSSL); err != nil {
		log.Fatalf("Erro ao baixar arquivo: %v\n", err)
	}

	// Exemplo de upload
	// filePath := "panda-vermelho.png"
	// uploadObjectName := "panda-vermelho.png"

	// if err := uploadFile(endpoint, accessKey, secretKey, bucketName, uploadObjectName, filePath, useSSL); err != nil {
	// 	log.Fatalf("Erro ao enviar arquivo: %v\n", err)
	// }
}







/*
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func downloadFile(bucket_name, object_name, destination string) error{
	ctx:= context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil{
		return fmt.Errorf("falha em criar conexao com o bucket da cloud: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucket_name)
	object := bucket.Object((object_name))

	reader, err := object.NewReader(ctx)
	if err != nil{
		return fmt.Errorf("falha em fazer um Reader com a cloud: %v", err)
	}
	defer reader.Close()

	file, err := os.Create(destination)
	if err != nil{
		return fmt.Errorf("falha em criar diretorio para baixar o arquivo: %v", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil{
		return fmt.Errorf("falha ao copiar dados para arquivo: %v", err)
	}

	fmt.Printf("Arquivo %s baixado em %s\n", object_name, destination)
	return nil
}


func sendFile(w io.Writer, bucket_name, name_file, object string) error{
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil{
		return fmt.Errorf("falha em criar conexao com o bucket da cloud: %w", err)
	}
	defer client.Close()

	file, err := os.Open(name_file)
	if err != nil {
		return fmt.Errorf("erro ao carregar arquivo: %w", err)
	}
	defer file.Close()

	o := client.Bucket(bucket_name).Object(object)

	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return fmt.Errorf("erro em io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("erro em Writer.Close: %w", err)
	}

	fmt.Fprintf(w, "Blob %v uploaded.\n", object)
	return nil
}


func main(){
	// bucketName := "video2gif"
	// objectName := "tmp/input.mp4"
	// destination := "./tmp/input.mp4"

	// if err := downloadFile(bucketName, objectName, destination); err != nil {
	// 	log.Fatalf("Erro ao baixar arquivo (a função deu pau): %v\n", err)
	// }

	bucketName := "video2gif"           // Nome do bucket no Google Cloud Storage
    objectName := "tmp/panda-vermelho.png"       // Caminho do objeto no bucket (coloca o arquivo na "pasta" tmp)
	file := "panda-vermelho.png"

    // Chama a função uploadFile e passa os parâmetros necessários
    if err := sendFile(os.Stdout, bucketName, file, objectName); err != nil {
        log.Fatalf("Falha no upload do arquivo: %v", err)
    }

}






// func main(){

// 	ctx := context.Background()

// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Read the object1 from bucket.
// 	rc, err := client.Bucket("video2gif").Object("tmp/input.mp4").NewReader(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer rc.Close()
// 	_, err = io.ReadAll(rc)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	v := rc.Attrs.Size
// 	fmt.Println(v)	

// }
*/
