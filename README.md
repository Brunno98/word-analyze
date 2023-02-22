Descrição:
Aplicação qua analisa o texto informado e retorna algumas infomações como, por exemplo, uma lista de palavras e suas quantidades

Requisitos funcionais:
[x] usuario deve poder enviar um texto;
[x] usuario recebe resposta as palavras contidadas no texto e sua frequencia;
[x] na resposta deve vir tambem a informação das palavras mais e menos usadas, caso tenha mais de 1, usar a primeira de acordo com a ordem alfabetica
[ ] informar quais palavras são palindromo


Requisitos não funcionais:
- usar a linguagem Go;

Regras de negocio:
[ ] caso o texto enviado esteja vazio ou tenha somente espaços, deve retornar um BAD_REQUEST;
[x] deve-se tratar letras maiusculas iguais as minusculas;
[x] ignorar o caracter '.';
[ ] devem ser ignorados caracters nao alfanumericos (':', '/', etc) no inicio e fim da palavra. ex: 'lorem:' => 'lorem', '/foo' => 'foo';
[ ] deve ser considerado caracters não alfanumerocos no meio da palavra. ex: 'lo-rem', 'foo/bar';