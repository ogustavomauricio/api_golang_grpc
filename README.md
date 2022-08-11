# API-GOLANG-WITH-GRPC

## 😃 Seja bem-vindo(a)
Projeto realizado como parte do desafio técnico da klever

### 💻 Sobre o projeto
The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an upvote service endpoints.

### ⚙️ Features
Technical requirements:

- [x] Keep the code in Github

API:

- [x] The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
- [x] The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct
- [x]  The API should contain unit test of methods it uses

Extra:

- [ ] Deliver the whole solution running in some free cloud service

## 🌀 Desenvolvimento
Para iniciar o desenvolvimento, é necessário clonar o projeto do GitHub para sua máquina.

```git@github.com:ogustavomauricio/api_golang_grpc.git```

##  🛠️ Instalação do Insomnia
- **O que é?** 
  Uma ferramenta cliente de API REST, como o Postman, mas tem alguns recursos adicionais, como suporte a GraphQL, gRPC, entre outros.

##  🚀 Pré-requisitos
 Antes de começar, rode os comandos abaixo em sequência.

Para **instalar as bibliotecas**, necessárias da aplicação, rode na raiz do projeto o comando: 
`make mod`

Para **buildar** o servidor, execute;
`make build`

Para **inicializar o servidor**, execute:
`make server`

#### ✨ É esperada a seguinte mensagem:

![step-one](/assets/Server.png)

Pronto! Agora você já pode utilizar a apliacação

## Testes
Para checar se a aplicação está funcionando devidamente, rode na raiz do projeto o comando: ```make test```

##  🛠️ Como utilizar o insomnia?


Abra o insomnia e clique no botão Create e logo em seguida no botão Request Collection

----
![step-one](/assets/image1.jpeg)





Em seguida, você deve definir um nome para o novo workspace.  No meu caso, irei colocar o nome da minha aplicação, que é Challenge-Klever. Depois clique em **Create**.
Obs: Caso queira você pode definir outro nome.



----
![step-one](/assets/image2.jpeg)


Para criar uma requisição podemos usar o atalho  _Ctrl + N_  ou clicar no botão **+** e depois ir na opção  _New Request_, como na imagem a seguir:


 Logo em seguida, clique no **+** e selecione gRPC Request
![step-one](/assets/image3.jpeg)

Clique no botão  **Add Proto File**  e selecione o arquivo  **service.proto**, localizado em  `./proto`  na pasta onde o projeto foi clonado.
![step-one](/assets/image4.jpeg)


Insira o valor _localhost:50051_, fazendo isso o client do Insomnia conseguirá se conectar com o nosso server. Por fim, selecione um dos métodos da aplicação gRPC e estará pronto para uso.
![step-one](/assets/image5.jpeg)


## Primeira requisição
- **/CurrencyCoinService/createCoin**

A primeira requisição que irei criar será responsável por cadastrar uma nova moeda. Envia uma requisição com nome e preço. O retorno deve ser o nome, preço e a quantidade de votos.
**Obs: Os votos não DEVEM ser enviados. Eles recebem o valor automático de 0.**
<br>
<br>
<br>
✨ Resultado

Exemplo de requisição:
```
{
	"name": "BITCOIN",
	"price": 23.724,8
}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 23.724,8,
    "vote": 0
}
```


## Segunda requisição

- **/CurrencyCoinService/ListCoins**

A segunda requisição que irei criar será responsável pelo retorno de todas as moedas. Envia uma requisição vazia e será retornado todas as moedas.
<br>
<br>
<br>
✨ Resultado
 
Exemplo de requisição:
```
{}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 23724,8,
    "vote": 0
}
```

Exemplo de response:
```
{
	"name": "Ethereum",
	"price": 1825,48,
    "vote": 865
}
```

Exemplo de response:
```
{
	"name": "Tether",
	"price": 10002,
    "vote": 5856
}
```

## Terceira requisição
- **/CurrencyCoinService/UpvoteCoin**

Responsável por aumentar o número de votos em 1. O corpo da requisição deve ser o nome da moeda e o retorno será o nome, preço e voto atualizado.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**
<br>
<br>
<br>
✨ Resultado
Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
    "vote": "1"
}
```


## Quarta requisição
- **/CurrencyCoinService/DownvoteCoin**
Responsável por decrementar o número de votos em 1. O corpo da requisição deve ser o nome da moeda e o retorno será o nome, preço e voto atualizado.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**
<br>
<br>
<br>
✨ Resultado
Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
    "vote": 0
}
```

## Quinta requisição
- **/CurrencyCoinService/Delete**

Responsável por deletar uma moeda. O corpo da requisição deve ser o nome da moeda e o retorno será uma mensagem confirmando a exclusão.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**
<br>
<br>
<br>
✨ Resultado
Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"message": "BITCOIN was deleted successful!
}
```



###   ✅ Autor
---

<a href="https://www.linkedin.com/in/o-gustavo-mauricio/">
 <img style="border-radius: 50%;" src="./assets/FotoGUstavo.png" width="100px;" alt=""/>
 <br />
 <sub><b>Gustavo Mauricio</b></sub></a> <a href="https://www.linkedin.com/in/o-gustavo-mauricio/" title="Linkedin">🚀</a>


Feito com ❤️ por Gustavo Mauricio 👋🏽 Entre em contato!

[![Linkedin Badge](https://img.shields.io/badge/-Gustavo-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/o-gustavo-mauricio/)](https://www.linkedin.com/in/o-gustavo-mauricio/) 
[![Gmail Badge](https://img.shields.io/badge/-gmauricio207@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:gmauricio207@gmail.com)](mailto:gmauricio207@gmail.com)
## Sujestões e FeedBack são sempre bem-vindos :)
