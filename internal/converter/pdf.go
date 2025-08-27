package converter

import (
	"fmt"
	"os"
	"path/filepath"

	"markdown-to-html/internal/utils"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// ConvertToPDF converts markdown content to PDF
func ConvertToPDF(markdown string, outputPath string, theme string) error {
	// First convert markdown to HTML
	html, err := ConvertToHTML(markdown, theme)
	if err != nil {
		return fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	// Create temporary HTML file
	tempHTML := filepath.Join(os.TempDir(), "temp_markdown.html")
	err = os.WriteFile(tempHTML, []byte(html), 0644)
	if err != nil {
		return fmt.Errorf("failed to create temporary HTML file: %w", err)
	}
	defer os.Remove(tempHTML)

	// Set wkhtmltopdf path if not in PATH
	wkhtmltopdfPath := getWkhtmltopdfPath()
	if wkhtmltopdfPath != "" {
		wkhtmltopdf.SetPath(wkhtmltopdfPath)
	}

	// Create PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("failed to create PDF generator: %w", err)
	}

	// Set PDF options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginTop.Set(20)
	pdfg.MarginBottom.Set(20)
	pdfg.MarginLeft.Set(20)
	pdfg.MarginRight.Set(20)

	// Add page
	page := wkhtmltopdf.NewPage(tempHTML)
	page.EnableLocalFileAccess.Set(true)
	page.LoadErrorHandling.Set("ignore")
	page.LoadMediaErrorHandling.Set("ignore")
	
	// Set page options for better rendering
	page.Zoom.Set(1.0)
	page.JavascriptDelay.Set(1000) // Wait for JS to load
	
	pdfg.AddPage(page)

	// Generate PDF
	err = pdfg.Create()
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	// Ensure output directory exists
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write PDF to file
	err = pdfg.WriteFile(outputPath)
	if err != nil {
		return fmt.Errorf("failed to write PDF file: %w", err)
	}

	return nil
}

// ConvertMarkdownFileToPDF converts a markdown file to PDF
func ConvertMarkdownFileToPDF(inputFile string, outputFile string, theme string) error {
	// Read markdown file
	content, err := utils.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read markdown file: %w", err)
	}

	// Convert to PDF
	return ConvertToPDF(content, outputFile, theme)
}

// IsWkhtmltopdfInstalled checks if wkhtmltopdf is installed
func IsWkhtmltopdfInstalled() bool {
	// Check common installation paths
	paths := []string{
		"/usr/local/bin/wkhtmltopdf",
		"C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",
		"C:\\Program Files (x86)\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",
		"wkhtmltopdf", // Check if it's in PATH
	}
	
	for _, path := range paths {
		if path == "wkhtmltopdf" {
			// Check if it's in PATH
			if _, err := os.Stat("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe"); err == nil {
				return true
			}
			if _, err := os.Stat("C:\\Program Files (x86)\\wkhtmltopdf\\bin\\wkhtmltopdf.exe"); err == nil {
				return true
			}
		} else {
			if _, err := os.Stat(path); err == nil {
				return true
			}
		}
	}
	
	return false
}

// getWkhtmltopdfPath returns the path to wkhtmltopdf executable
func getWkhtmltopdfPath() string {
	// Check common installation paths
	paths := []string{
		"C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",
		"C:\\Program Files (x86)\\wkhtmltopdf\\bin\\wkhtmltopdf.exe",
		"/usr/local/bin/wkhtmltopdf",
		"/usr/bin/wkhtmltopdf",
	}
	
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	
	return ""
}

// GetWkhtmltopdfInstallInstructions returns installation instructions
func GetWkhtmltopdfInstallInstructions() string {
	return `wkhtmltopdf is required for PDF generation.

Installation:

Windows:
1. Download from: https://wkhtmltopdf.org/downloads.html
2. Install to default location: C:\Program Files\wkhtmltopdf\
3. Add to PATH: C:\Program Files\wkhtmltopdf\bin\

macOS:
brew install wkhtmltopdf

Linux (Ubuntu/Debian):
sudo apt-get install wkhtmltopdf

Linux (CentOS/RHEL):
sudo yum install wkhtmltopdf

After installation, restart your terminal and try again.`
}
