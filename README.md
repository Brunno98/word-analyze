Descrição:
Aplicação qua analisa o texto informado e retorna algumas infomações como, por exemplo, uma lista de palavras e suas quantidades

Requisitos funcionais:
[x] usuario deve poder enviar um texto;
[ ] usuario recebe como resposta a lista de palavras contidadas no texto e sua frequencia;
    [ ] a resposta não esta sendo uma lista, e sim um array
[ ] na resposta deve vir tambem a informação das palavras mais e menos usadas, caso tenha mais de 1, usar a primeira de acorda com a ordem alfabetica
[ ] informar quais palavras são palindromo


Requisitos não funcionais:
- usar a linguagem Go;

Regras de negocio:
[ ] caso o texto enviado esteja vazio ou tenha somente espaços, deve retornar um BAD_REQUEST;
[x] deve-se tratar letras maiusculas iguais as minusculas;
[x] ignorar o caracter '.';