Aplicativo de Exemplo em Golang com Arquitetura Hexagonal
=========================================================

Este é um exemplo de aplicativo desenvolvido em Golang utilizando a arquitetura hexagonal, também conhecida como arquitetura ports and adapters. Neste exemplo, a camada externa é implementada utilizando REST e a base de dados é gerenciada pelo MongoDB.

Arquitetura Hexagonal
---------------------

A arquitetura hexagonal é um padrão de arquitetura de software que visa desacoplar as preocupações do núcleo do negócio das implementações externas, como interfaces de usuário, serviços externos, etc. Isso é alcançado através da definição de "portas" de entrada e saída que isolam o núcleo do negócio das implementações específicas. Esta arquitetura promove a testabilidade, a manutenção e a escalabilidade do software.

Componentes
-----------

### 1\. Core (Núcleo)

O núcleo do aplicativo contém a lógica de negócio. Ele não conhece os detalhes de implementação de entrada e saída. Em vez disso, define interfaces (portas) que devem ser implementadas pelas camadas externas.

### 2\. Adaptadores

Os adaptadores são responsáveis por implementar as interfaces definidas no núcleo. Neste exemplo, temos adaptadores para lidar com as operações REST e MongoDB.

Pré-requisitos
--------------

Certifique-se de ter Golang e MongoDB instalados em seu sistema antes de prosseguir.

Configuração
------------

1.  Clone este repositório:


`git clone https://github.com/erkylima/hexagonal-golang-example.git`

1.  Instale as dependências:


`go mod tidy`

1.  Configure as variáveis de ambiente necessárias:


`MONGO_URL=sua-uri-do-mongodb`
`MONGO_DB=beneficiary`
`MONGO_TIMEOUT=3`


1.  Inicie o MongoDB.
    
2.  Execute o aplicativo:
    

`go run main.go`

Uso
---

Este exemplo fornece endpoints REST básicos para manipulação de entidades. Você pode testar os endpoints utilizando ferramentas como cURL ou Postman.

*   GET `/{name}`: Retorna uma entidade específica com o ID fornecido.
*   POST `/`: Cria uma nova entidade. Envie os dados da entidade no corpo da solicitação.

Certifique-se de substituir `sua-uri-do-mongodb` pela URI real do seu banco de dados MongoDB.

Contribuição
------------

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas ou enviar solicitações de pull.

Licença
-------

Este projeto está licenciado sob a MIT License.
