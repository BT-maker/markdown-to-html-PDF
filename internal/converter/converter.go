package converter

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"markdown-to-html/internal/utils"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// ConvertToHTML converts markdown content to HTML with Bootstrap styling
func ConvertToHTML(markdown string, theme string) (string, error) {
	// Create markdown parser with extensions
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM), // GitHub Flavored Markdown
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", fmt.Errorf("failed to convert markdown: %w", err)
	}

	htmlContent := buf.String()

	// Wrap in HTML template with Bootstrap
	return wrapInHTMLTemplate(htmlContent, theme), nil
}

// wrapInHTMLTemplate wraps the HTML content in a complete HTML document with Bootstrap
func wrapInHTMLTemplate(content string, theme string) string {
	var cssTheme string
	var bodyClass string

	switch theme {
	case "dark":
		cssTheme = `
		body {
			background-color: #212529;
			color: #ffffff;
		}
		.card {
			background-color: #343a40;
			border-color: #495057;
		}
		.navbar {
			background-color: #343a40 !important;
		}
		.navbar-brand {
			color: #ffffff !important;
		}
		.nav-link {
			color: #ffffff !important;
		}
		.nav-link:hover {
			color: #ffffff !important;
		}
		.navbar-text {
			color: #ffffff !important;
		}
		.markdown-content {
			color: #ffffff;
		}
		.markdown-content h1,
		.markdown-content h2,
		.markdown-content h3,
		.markdown-content h4,
		.markdown-content h5,
		.markdown-content h6 {
			color: #ffffff;
		}
		.markdown-content p {
			color: #ffffff;
		}
		.markdown-content li {
			color: #ffffff;
		}
		.markdown-content blockquote {
			color: #e9ecef;
		}
		.markdown-content code {
			background-color: #2d3748 !important;
			color: #e2e8f0 !important;
		}
		.markdown-content pre {
			background-color: #2d3748 !important;
			color: #e2e8f0 !important;
		}
		.markdown-content pre code {
			background-color: transparent !important;
			color: #e2e8f0 !important;
		}
		/* Override Prism.js styles */
		pre[class*="language-"] {
			background-color: #2d3748 !important;
			color: #e2e8f0 !important;
		}
		code[class*="language-"] {
			background-color: transparent !important;
			color: #e2e8f0 !important;
		}
		.token.comment,
		.token.prolog,
		.token.doctype,
		.token.cdata {
			color: #718096 !important;
		}
		.token.punctuation {
			color: #e2e8f0 !important;
		}
		.token.property,
		.token.tag,
		.token.boolean,
		.token.number,
		.token.constant,
		.token.symbol,
		.token.deleted {
			color: #f687b3 !important;
		}
		.token.selector,
		.token.attr-name,
		.token.string,
		.token.char,
		.token.builtin,
		.token.inserted {
			color: #68d391 !important;
		}
		.token.operator,
		.token.entity,
		.token.url,
		.language-css .token.string,
		.style .token.string {
			color: #f6ad55 !important;
		}
		.token.atrule,
		.token.attr-value,
		.token.keyword {
			color: #63b3ed !important;
		}
		.token.function,
		.token.class-name {
			color: #b794f4 !important;
		}
		.token.regex,
		.token.important,
		.token.variable {
			color: #fc8181 !important;
		}
		.markdown-content table th,
		.markdown-content table td {
			color: #ffffff;
		}
		.markdown-content a {
			color: #86b7fe;
		}
		.markdown-content a:hover {
			color: #b6d4fe;
		}`
		bodyClass = "bg-dark text-light"
	default:
		cssTheme = ""
		bodyClass = ""
	}

	template := fmt.Sprintf(`<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Markdown to HTML</title>
    
    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    
    <!-- Prism.js for syntax highlighting -->
    <link href="https://cdn.jsdelivr.net/npm/prismjs@1.29.0/themes/prism.min.css" rel="stylesheet">
    
    <!-- Custom styles -->
    <style>
        %s
        
        .markdown-content {
            line-height: 1.6;
        }
        
        .markdown-content h1,
        .markdown-content h2,
        .markdown-content h3,
        .markdown-content h4,
        .markdown-content h5,
        .markdown-content h6 {
            margin-top: 2rem;
            margin-bottom: 1rem;
            font-weight: 600;
        }
        
        .markdown-content h1 {
            border-bottom: 2px solid #dee2e6;
            padding-bottom: 0.5rem;
        }
        
        .markdown-content h2 {
            border-bottom: 1px solid #dee2e6;
            padding-bottom: 0.3rem;
        }
        
        .markdown-content p {
            margin-bottom: 1rem;
        }
        
        .markdown-content blockquote {
            border-left: 4px solid #007bff;
            padding-left: 1rem;
            margin-left: 0;
            color: #6c757d;
        }
        
        .markdown-content code {
            background-color: #f8f9fa;
            padding: 0.2rem 0.4rem;
            border-radius: 0.25rem;
            font-size: 0.875em;
        }
        
        .markdown-content pre {
            background-color: #f8f9fa;
            padding: 1rem;
            border-radius: 0.5rem;
            overflow-x: auto;
        }
        
        .markdown-content pre code {
            background-color: transparent;
            padding: 0;
        }
        
        .markdown-content table {
            width: 100%%;
            margin-bottom: 1rem;
        }
        
        .markdown-content table th,
        .markdown-content table td {
            padding: 0.75rem;
            border-top: 1px solid #dee2e6;
        }
        
        .markdown-content table thead th {
            vertical-align: bottom;
            border-bottom: 2px solid #dee2e6;
        }
        
        .markdown-content img {
            max-width: 100%%;
            height: auto;
            border-radius: 0.5rem;
            margin: 1rem 0;
        }
        
        .markdown-content ul,
        .markdown-content ol {
            margin-bottom: 1rem;
            padding-left: 2rem;
        }
        
        .markdown-content li {
            margin-bottom: 0.5rem;
        }
        
        .markdown-content hr {
            margin: 2rem 0;
            border: 0;
            border-top: 1px solid #dee2e6;
        }
    </style>
</head>
<body class="%s">
    <!-- Navigation -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light mb-4">
        <div class="container">
            <a class="navbar-brand" href="#">
                <i class="bi bi-markdown"></i>
                Markdown to HTML Converter
            </a>
            <div class="navbar-nav ms-auto">
                <span class="navbar-text">
                    Generated with Go
                </span>
            </div>
        </div>
    </nav>

    <!-- Main Content -->
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-lg-8">
                <div class="card">
                    <div class="card-body">
                        <div class="markdown-content">
                            %s
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Bootstrap JS -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>
    
    <!-- Prism.js for syntax highlighting -->
    <script src="https://cdn.jsdelivr.net/npm/prismjs@1.29.0/components/prism-core.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/prismjs@1.29.0/plugins/autoloader/prism-autoloader.min.js"></script>
    
    <!-- Initialize syntax highlighting -->
    <script>
        Prism.highlightAll();
    </script>
</body>
</html>`, cssTheme, bodyClass, content)

	return template
}

// ConvertMultipleFiles converts multiple markdown files to HTML
func ConvertMultipleFiles(inputFiles []string, outputDir string, theme string) error {
	for _, inputFile := range inputFiles {
		// Read markdown file
		content, err := utils.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", inputFile, err)
		}

		// Convert to HTML
		html, err := ConvertToHTML(content, theme)
		if err != nil {
			return fmt.Errorf("failed to convert %s: %w", inputFile, err)
		}

		// Generate output filename
		base := filepath.Base(inputFile)
		ext := filepath.Ext(base)
		outputFile := filepath.Join(outputDir, strings.TrimSuffix(base, ext)+".html")

		// Write output file
		if err := utils.WriteFile(outputFile, html); err != nil {
			return fmt.Errorf("failed to write %s: %w", outputFile, err)
		}

		fmt.Printf("Converted %s to %s\n", inputFile, outputFile)
	}

	return nil
}
