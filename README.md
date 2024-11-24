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
