package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"markdown-to-html/internal/converter"
	"markdown-to-html/internal/utils"
)

type ConversionRequest struct {
	Markdown string `json:"markdown"`
	Theme    string `json:"theme"`
	Format   string `json:"format"`
}

type ConversionResponse struct {
	Success bool   `json:"success"`
	HTML    string `json:"html,omitempty"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	
	// Routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/convert", handleConvert)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/download", handleDownload)
	
	port := ":8080"
	fmt.Printf("🚀 Web arayüzü başlatılıyor... http://localhost%s\n", port)
	fmt.Println("📝 Markdown dosyalarınızı web arayüzünden dönüştürebilirsiniz!")
	
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := map[string]interface{}{
		"Title": "Markdown to HTML/PDF Converter",
	}
	
	tmpl.Execute(w, data)
}

func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Validate input
	if strings.TrimSpace(req.Markdown) == "" {
		json.NewEncoder(w).Encode(ConversionResponse{
			Success: false,
			Error:   "Markdown içeriği boş olamaz",
		})
		return
	}
	
	// Set default theme
	if req.Theme == "" {
		req.Theme = "light"
	}
	
	// Set default format
	if req.Format == "" {
		req.Format = "html"
	}
	
	var response ConversionResponse
	
	if req.Format == "html" {
		// Convert to HTML
		html, err := converter.ConvertToHTML(req.Markdown, req.Theme)
		if err != nil {
			response = ConversionResponse{
				Success: false,
				Error:   "HTML dönüştürme hatası: " + err.Error(),
			}
		} else {
			response = ConversionResponse{
				Success: true,
				HTML:    html,
			}
		}
	} else if req.Format == "pdf" {
		// For PDF, we'll return a success message and handle download separately
		response = ConversionResponse{
			Success: true,
			HTML:    "PDF hazırlanıyor...",
		}
	} else {
		response = ConversionResponse{
			Success: false,
			Error:   "Desteklenmeyen format: " + req.Format,
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "File upload error: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	// Check file extension
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".md") {
		http.Error(w, "Sadece .md dosyaları kabul edilir", http.StatusBadRequest)
		return
	}
	
	// Read file content
	content, err := utils.ReadFileFromReader(file)
	if err != nil {
		http.Error(w, "File read error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Return file content as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"content": content,
		"filename": header.Filename,
	})
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	if req.Format == "pdf" {
		// Generate PDF
		if !converter.IsWkhtmltopdfInstalled() {
			http.Error(w, "wkhtmltopdf yüklü değil", http.StatusInternalServerError)
			return
		}
		
		// Create temporary file
		tempFile := filepath.Join("output", "temp_download.pdf")
		
		err := converter.ConvertToPDF(req.Markdown, tempFile, req.Theme)
		if err != nil {
			http.Error(w, "PDF generation error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		
		// Serve file for download
		w.Header().Set("Content-Disposition", "attachment; filename=converted.pdf")
		w.Header().Set("Content-Type", "application/pdf")
		http.ServeFile(w, r, tempFile)
		
		// Clean up temp file
		// Note: In production, you might want to use a proper temp file system
	} else {
		http.Error(w, "Unsupported format for download", http.StatusBadRequest)
	}
}
