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
	"github.com/yuin/goldmark/extension"
)

type Note struct {
	Title    string
	Content  template.HTML
	Filename string
	Previous *Note
	Next     *Note
}

func main() {
	postsDir := "posts"
	outputDir := "public"
	templateDir := "templates"

	// Ler e ordenar os arquivos Markdown
	files, err := ioutil.ReadDir(postsDir)
	if err != nil {
		log.Fatalf("Erro ao ler o diretório '%s': %v", postsDir, err)
	}

	var mdFiles []fs.FileInfo
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			mdFiles = append(mdFiles, file)
		}
	}

	sort.Slice(mdFiles, func(i, j int) bool {
		return mdFiles[i].Name() < mdFiles[j].Name()
	})

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
		md := goldmark.New(
			goldmark.WithExtensions(
				extension.Table,
			),
		)

		if err := md.Convert(content, &buf); err != nil {
			log.Printf("Erro ao converter Markdown para HTML: %v", err)
			continue
		}

		// Extrair título
		lines := strings.Split(string(content), "\n")
		title := "Sem Título"
		for _, line := range lines {
			if strings.HasPrefix(line, "# ") {
				title = strings.TrimPrefix(line, "# ")
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

	// Associar links de navegação
	for i := range notes {
		if i > 0 {
			notes[i].Previous = notes[i-1]
		}
		if i < len(notes)-1 {
			notes[i].Next = notes[i+1]
		}
	}

	// Criar diretório de saída
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Erro ao criar o diretório '%s': %v", outputDir, err)
	}

	// ---------- GERAR NOTAS INDIVIDUAIS ----------
	tmplNote, err := template.ParseFiles(
		filepath.Join(templateDir, "base.html"),
		filepath.Join(templateDir, "note.html"),
	)
	if err != nil {
		log.Fatalf("Erro ao carregar templates de notas: %v", err)
	}

	for _, note := range notes {
		outputPath := filepath.Join(outputDir, note.Filename)
		outFile, err := os.Create(outputPath)
		if err != nil {
			log.Printf("Erro ao criar %s: %v", outputPath, err)
			continue
		}

		err = tmplNote.ExecuteTemplate(outFile, "base", note)
		if err != nil {
			log.Printf("Erro ao gerar %s: %v", outputPath, err)
		} else {
			fmt.Printf("Página gerada: %s\n", outputPath)
		}

		outFile.Close()
	}

	// ---------- GERAR INDEX ----------
	tmplIndex, err := template.ParseFiles(
		filepath.Join(templateDir, "base.html"),
		filepath.Join(templateDir, "index.html"),
	)
	if err != nil {
		log.Fatalf("Erro ao carregar templates do índice: %v", err)
	}

	indexData := struct {
		Title string
		Notes []*Note
	}{
		Title: "Início",
		Notes: notes,
	}

	indexPath := filepath.Join(outputDir, "index.html")
	indexFile, err := os.Create(indexPath)
	if err != nil {
		log.Fatalf("Erro ao criar o index.html: %v", err)
	}
	defer indexFile.Close()

	err = tmplIndex.ExecuteTemplate(indexFile, "base", indexData)
	if err != nil {
		log.Fatalf("Erro ao gerar o index.html: %v", err)
	} else {
		fmt.Printf("Página de índice gerada: %s\n", indexPath)
	}
}
