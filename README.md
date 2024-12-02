* Video -> GIF

* API p/ receber os parametros. start_time, end_time, fps ✅
* recebeu video usando http ✅
* salvou video localmente ✅
* Passar isso p/ func de convert video to gif ✅

* Salva isso no storage
* Retornar o link do gif
* Deployar no Cloud Run


POST http://localhost:8080/covert?start_time=0&end_time=10&fps=10

curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@input.mp4;type=video/mp4" "http://localhost:8080/convert?start_time=0&end_time=10&fps=10"

criamos uma VM na magaluCloud, instalamos docker nela.

desabilitamos a firewall

Escreveu um docker compose 

Adicionou DNS de luiz-teste.patos.dev (pra nao precisar escerver o IP)

fizemos:
- git clone na VM
- docker compose up --build (fez com que a porta do container se conectasse com a porta da máquina)
- por meio de outra maquina, fizemos o comando curl acima com o nome da VM: curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@input.mp4;type=video/mp4" "http://teste-luiz.patos.dev:8080/convert?start_time=0&end_time=10&fps=10"
- ja esta rodando