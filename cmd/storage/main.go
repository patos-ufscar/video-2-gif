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

