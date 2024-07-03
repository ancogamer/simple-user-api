# simple-user-api

Algumas considerações.
Durante o codigo, você verá comentários em inglês, bom, nada a comentar sobre...

Bom o mock do testify pode não ser o mais rapido ou com menor número de dependências, mas é simples, então usei.

Sobre o google CMP, eu gosto então usei, ele da um bom print no diff;

Sobre os erros, bom, fiz desta maneira pois me da uma liberdade, mas vale lembrar que é possível encapsular o erro de cada "camada" e usar as interfaces da parte de errors.

Sobre alguns lugares estarem sem teste, bom, como é só uma api de exemplo, prefiri não colocar em todos os lugares.

# Makes do projeto
Make run -> executa

Make cover -> roda todos os testes e te da uma cobertura total

Make deps -> sobe as dependências
