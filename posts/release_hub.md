# “Landing pages elegantes para GitHub Releases”.

# **🎯** Proposta de Valor

Permitir que qualquer repositório GitHub tenha uma **landing page de downloads**, onde os releases são apresentados de forma organizada, filtrada e estilizada — configurável via YAML no próprio repo.

# **⚙️**  **Como funciona (alto nível)**

### **1. O app é instalado em um repositório GitHub**
- Pode ser via GitHub App ou apenas via workflow + action.
- Ele lê o arquivo .releasehub.yml na raiz.
### **2.** **A LP é gerada automaticamente**
- Está hospedada:
    - Via GitHub Pages (ideal)
    - Ou via teu servidor / Vercel
- O app lê a aba **Releases** do GitHub:
    - Tags
    - Assets (zip, tar.gz, exe, apk, dmg…)
    - Changelog
    - Data e versão
### **3.** **O YAML controla tudo**

Exemplo:

```yaml
title: "Downloads oficiais"
description: "Versões estáveis e nightly builds"
theme: "dark"

tabs:
  - id: stable
    label: "Stable"
    filter:
      prerelease: false

  - id: nightly
    label: "Nightly"
    filter:
      prerelease: true

  - id: custom
    label: "Custom"
    include:
      - "v1.2.0"
      - "2024-09-beta-*"
    extensions:
      - ".zip"
      - ".exe"

global:
  hide_assets:
    - "*.sig"
  hide_releases:
    - "v0.*"
```

---

# **🧩**  **Funcionalidades chave** **(com detalhamento técnico)**


## **1. Filtragem avançada de releases**

- por regex
- por extensão
- por pré-lançamento / rascunho
- por nome do asset
- ocultar versões antigas
- mostrar só releases LTS
- ordenar por data/semver

## **2. Até 3 tabs**

Cada tab funciona como uma LP independente, mas **herda todas as configs do global**:

- Tab 1: Stable    
- Tab 2: Beta
- Tab 3: Custom (qualquer regra do YAML)

Tabs aparecem como “abas” na UI
  

## **3. LP central que agrega vários repositórios**
  
### Repo central →  .releasehub.yml :

```
aggregate:
  - owner/project-a
  - owner/project-b
  - owner/project-c
```

Então gera uma LP assim:

```
/ → Mostra todos os projetos agregados
/project-a → LP do projeto A
/project-b → LP do projeto B
```

Perfeito para empresas com vários projetos open-source.

---

# **🏗** **Arquitetura da Solução**

- GitHub REST API (ou GraphQL)
    Endpoints usados:
    - GET /repos/{owner}/{repo}/releases
    - GET /repos/{owner}/{repo}/releases/assets
    - GET /repos/{owner}/{repo} (metadados)
### **✔ Output**
- Build de páginas estáticas (SSG):
    - Astro.js (melhor)
    - Next.js static export
    - Eleventy / Jekyll
- Hospedado automaticamente no:
    - GitHub Pages
    - OU Vercel (melhor experiência)

### **✔ GitHub Action**
Toda vez que:
- lançar release
- modificar .releasehub.yml
    → Regenera a LP.

---

# **🔧**  **Instalação no repositório**

O usuário adiciona:

```
.github/workflows/releasehub.yml
.releasehub.yml
```

E pronto.

---

# **🎨** **UI / UX da LP**
  
A LP pode ter:
- Nome do projeto
- Descrição
- Tag mais recente destacada
- Download por plataforma
- Changelog expandível    
- Badges (stable / prerelease)
- Botões grandes e intuitivos para cada asset
- QR code opcional para downloads mobile
- Analytics opcional


Exemplo de seção por release:

```
Version v1.3.2 — 12 Dec 2025
---------------------------------
Windows 64-bit (.exe)
Linux (.AppImage)
Mac (.dmg)
Source Code (.zip)
Changelog ▶
```

---

# **🧩**  Funcionalidades extras


## 1. Short URLs
yourapp.com/p/project/release/latest

## **2.**  **Auto-detectar plataforma**
Quando abrir no Windows → já mostra .exe
Quando abrir no Android → já mostra .apk

## **3.** **Analytics básicos**
- Downloads por dia
- Sistemas operacionais
- Version ranking

## **4.**  **Proteção opcional**
- Releases privados? → exige token/personal access.

---

# **🛠**  **Roadmap para implementar**
### **MVP**
1. Ler releases via GitHub API
2. Interpretar .releasehub.yml
3. Gerar LP estática
4. Suportar tabs
5. Suportar filtros básicos

### **v1**

6. Agregador de vários repositórios
7. Tema dark/light
8. Deploy automático (GitHub Pages)
### **v2**
9. Analytics
10. Detecção de plataforma
11. Short URLs
12. Customizações profundas via YAML
13. Templates customizáveis por HTML/MDX

---
