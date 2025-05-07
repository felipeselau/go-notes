package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
)

type Note struct {
	Title    string
	Content  template.HTML
	Filename string
	Previous *Note
	Next     *Note
}

func main() {
	// Diretórios
	postsDir := "posts"
	outputDir := "public"
	templateFile := "templates/base.html"

	// Ler arquivos Markdown
	files, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatalf("Erro ao ler o diretório '%s': %v", postsDir, err)
	}

	// Filtrar e ordenar arquivos .md
	var mdFiles []fs.FileInfo
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			mdFiles = append(mdFiles, file)
		}
	}
	sort.Slice(mdFiles, func(i, j int) bool {
		return mdFiles[i].Name() < mdFiles[j].Name()
	})

	// Processar notas
	var notes []*Note
	for _, file := range mdFiles {
		inputPath := filepath.Join(postsDir, file.Name())
		content, err := ioutil.ReadFile(inputPath)
		if err != nil {
			log.Printf("Erro ao ler o arquivo %s: %v", inputPath, err)
			continue
		}

		// Converter Markdown para HTML
		var buf bytes.Buffer
		if err := goldmark.Convert(content, &buf); err != nil {
			log.Printf("Erro ao converter Markdown para HTML: %v", err)
			continue
		}

		// Extrair título (primeira linha que começa com '# ')
		lines := strings.Split(string(content), "\n")
		title := "Sem Título"
		for _, line := range lines {
			if strings.HasPrefix(line, "# ") {
				title = strings.TrimSpace(strings.TrimPrefix(line, "# "))
				break
			}
		}

		note := &Note{
			Title:    title,
			Content:  template.HTML(buf.String()),
			Filename: strings.TrimSuffix(file.Name(), ".md") + ".html",
		}
		notes = append(notes, note)
	}

	// Associar notas anteriores e próximas
	for i := range notes {
		if i > 0 {
			notes[i].Previous = notes[i-1]
		}
		if i < len(notes)-1 {
			notes[i].Next = notes[i+1]
		}
	}

	// Carregar template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Erro ao carregar o template: %v", err)
	}

	// Criar diretório de saída
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Erro ao criar o diretório '%s': %v", outputDir, err)
	}

	// Gerar arquivos HTML
	for _, note := range notes {
		outputPath := filepath.Join(outputDir, note.Filename)
		outFile, err := os.Create(outputPath)
		if err != nil {
			log.Printf("Erro ao criar o arquivo %s: %v", outputPath, err)
			continue
		}

		if err := tmpl.Execute(outFile, note); err != nil {
			log.Printf("Erro ao executar o template para %s: %v", outputPath, err)
			outFile.Close()
			continue
		}

		outFile.Close()
		fmt.Printf("Página gerada: %s\n", outputPath)
	}
}
