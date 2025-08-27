package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"markdown-to-html/internal/converter"
	"markdown-to-html/internal/utils"

	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	preview    bool
	theme      string
	format     string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "markdown-converter [input.md] [output]",
		Short: "Convert Markdown files to HTML or PDF with Bootstrap styling",
		Long: `A command line tool to convert Markdown files to HTML or PDF.
		
Examples:
  markdown-converter input.md output.html
  markdown-converter input.md output.pdf --format pdf
  markdown-converter input.md --format pdf                    # Outputs to input.pdf
  markdown-converter input.md --preview                       # Shows HTML preview in terminal
  markdown-converter input.md --theme dark --format pdf       # Uses dark theme for PDF`,
		Args: cobra.MaximumNArgs(2),
		Run:  run,
	}

	rootCmd.Flags().BoolVarP(&preview, "preview", "p", false, "Show preview in terminal")
	rootCmd.Flags().StringVarP(&theme, "theme", "t", "light", "Theme: light or dark")
	rootCmd.Flags().StringVarP(&format, "format", "f", "html", "Output format: html or pdf")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	// Parse arguments
	if len(args) == 0 {
		fmt.Println("Error: Input file is required")
		fmt.Println("Usage: markdown-to-html input.md [output.html]")
		os.Exit(1)
	}

	inputFile = args[0]
	
	if len(args) == 2 {
		outputFile = args[1]
	} else {
		// Auto-generate output filename
		ext := filepath.Ext(inputFile)
		base := strings.TrimSuffix(inputFile, ext)
		if format == "pdf" {
			outputFile = base + ".pdf"
		} else {
			outputFile = base + ".html"
		}
	}

	// Validate input file
	if !utils.FileExists(inputFile) {
		fmt.Printf("Error: Input file '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	// Read markdown file
	content, err := utils.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Check if wkhtmltopdf is installed for PDF conversion
	if format == "pdf" && !converter.IsWkhtmltopdfInstalled() {
		fmt.Println("Error: wkhtmltopdf is not installed")
		fmt.Println(converter.GetWkhtmltopdfInstallInstructions())
		os.Exit(1)
	}

	// Convert based on format
	switch format {
	case "pdf":
		// Convert markdown to PDF
		err = converter.ConvertMarkdownFileToPDF(inputFile, outputFile, theme)
		if err != nil {
			fmt.Printf("Error converting to PDF: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully converted '%s' to '%s'\n", inputFile, outputFile)
		
	case "html":
		// Convert markdown to HTML
		html, err := converter.ConvertToHTML(content, theme)
		if err != nil {
			fmt.Printf("Error converting markdown: %v\n", err)
			os.Exit(1)
		}

		// Show preview if requested
		if preview {
			fmt.Println("=== HTML Preview ===")
			fmt.Println(html)
			fmt.Println("=== End Preview ===")
			return
		}

		// Write output file
		err = utils.WriteFile(outputFile, html)
		if err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully converted '%s' to '%s'\n", inputFile, outputFile)
		
	default:
		fmt.Printf("Error: Unsupported format '%s'. Supported formats: html, pdf\n", format)
		os.Exit(1)
	}
}
