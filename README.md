# 📝 Go-Notes

Um gerador de site estático simples que transforma arquivos Markdown em páginas HTML estilizadas com Bootstrap 5 e a paleta de cores Catppuccin Mocha.

## 🚀 Demonstração

Acesse a versão hospedada em: [https://notes.luizfelipedev.com](https://notes.luizfelipedev.com)

## 📂 Estrutura do Projeto

```
├── public/             # Arquivos HTML gerados
├── templates/          # Templates HTML (base, nota, índice)
├── css/                # Estilos personalizados (custom.css)
├── notes/              # Arquivos Markdown (.md)
├── main.go             # Script principal em Go
├── deploy.yml          # Workflow do GitHub Actions
└── README.md           # Este arquivo
```

## ⚙️ Tecnologias Utilizadas

* [Go](https://golang.org/) para processamento de templates e geração de páginas
* [Bootstrap 5](https://getbootstrap.com/) para responsividade e estilo
* [Catppuccin Mocha](https://github.com/catppuccin/catppuccin) para paleta de cores
* [GitHub Pages](https://pages.github.com/) para hospedagem gratuita
* [GitHub Actions](https://github.com/features/actions) para CI/CD([warcontent.com][1])

## 🛠️ Como Usar

1. Clone o repositório:

   ```bash
   git clone https://github.com/luizfelipedev/notes.git
   cd notes
   ```

2. Adicione suas notas em formato Markdown na pasta `notes/`.

3. Execute o script principal para gerar as páginas HTML:

   ```bash
   go run main.go
   ```

4. Os arquivos HTML serão gerados na pasta `public/`.

## 🌐 Hospedagem no GitHub Pages

O projeto está configurado para ser hospedado na branch `gh-pages`. O arquivo `deploy.yml` automatiza o processo de build e deploy utilizando o GitHub Actions.

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
