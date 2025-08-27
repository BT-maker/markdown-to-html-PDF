# 🚀 Markdown to HTML/PDF Converter

Go diliyle geliştirilmiş, Markdown dosyalarını HTML ve PDF'e dönüştüren modern bir CLI ve web aracı.

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)

## ✨ Özellikler

- **🎨 Modern Tasarım**: Bootstrap 5 ile responsive ve şık HTML çıktısı
- **🌈 Syntax Highlighting**: Prism.js ile kod blokları için renklendirme
- **🌙 Tema Desteği**: Light ve Dark tema seçenekleri
- **💻 CLI Interface**: Kolay kullanım için komut satırı arayüzü
- **🌐 Web Interface**: Modern web arayüzü ile kolay kullanım
- **📄 PDF Support**: Markdown'dan PDF oluşturma
- **📁 Drag & Drop**: Web arayüzünde dosya sürükle-bırak
- **📝 GitHub Flavored Markdown**: GFM desteği ile gelişmiş markdown özellikleri
- **🔄 Otomatik İsimlendirme**: Çıktı dosyası belirtilmezse otomatik isimlendirme
- **⚡ Gerçek Zamanlı Önizleme**: Web arayüzünde anlık HTML önizleme
- **🎯 Akıllı İndirme**: Seçili formata göre HTML/PDF indirme

## 📋 Gereksinimler

- **Go 1.21** veya üzeri
- **Git**
- **wkhtmltopdf** (PDF dönüştürme için)

### wkhtmltopdf Kurulumu

**Windows:**
```bash
# Chocolatey ile
choco install wkhtmltopdf

# Manuel kurulum
# https://wkhtmltopdf.org/downloads.html adresinden indirin
```

**macOS:**
```bash
brew install wkhtmltopdf
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install wkhtmltopdf
```

**Linux (CentOS/RHEL):**
```bash
sudo yum install wkhtmltopdf
```

## 🛠️ Kurulum

1. **Projeyi klonlayın:**
```bash
git clone https://github.com/BT-maker/markdown-to-html.git
cd markdown-to-html
```

2. **Bağımlılıkları yükleyin:**
```bash
go mod tidy
```

3. **Uygulamayı derleyin:**
```bash
# CLI uygulaması
go build -o markdown-to-html cmd/main.go

# Web uygulaması
go build -o markdown-web cmd/web/main.go
```

4. **wkhtmltopdf kurulumunu kontrol edin:**
```bash
wkhtmltopdf --version
```

## 📖 Kullanım

### 🖥️ CLI Kullanımı

```bash
# Temel kullanım
./markdown-to-html input.md output.html

# PDF dönüştürme
./markdown-to-html input.md output.pdf --format pdf

# Otomatik isimlendirme
./markdown-to-html input.md

# Önizleme (terminalde HTML'i gösterir)
./markdown-to-html input.md --preview

# Dark tema ile dönüştürme
./markdown-to-html input.md --theme dark

# Dark tema ile PDF
./markdown-to-html input.md --theme dark --format pdf
```

### 🌐 Web Arayüzü

```bash
# Web arayüzünü başlat
./markdown-web

# Veya geliştirme modunda
go run cmd/web/main.go

# Windows batch dosyası ile
web.bat
```

**Tarayıcınızda:** `http://localhost:8080` adresine gidin.

### 🎯 Web Arayüzü Özellikleri

- **📁 Drag & Drop**: Markdown dosyalarını sürükleyip bırakın
- **👁️ Önizleme**: HTML ve PDF önizleme
- **⬇️ İndirme**: Tek tıkla HTML/PDF indirme
- **🌙 Tema**: Dark/Light tema değiştirme
- **📝 Editör**: Yerleşik markdown editörü

### 📋 Komut Satırı Seçenekleri

```bash
markdown-to-html [input.md] [output.html] [flags]

Flags:
  -p, --preview    Show preview in terminal
  -t, --theme      Theme: light or dark (default "light")
  -f, --format     Output format: html or pdf (default "html")
  -h, --help       Help for markdown-to-html
```

### 💡 Örnekler

```bash
# Örnek markdown dosyasını dönüştürme
./markdown-to-html examples/sample.md output/sample.html

# Dark tema ile önizleme
./markdown-to-html examples/sample.md --theme dark --preview

# PDF dönüştürme
./markdown-to-html examples/sample.md output/sample.pdf --format pdf

# Web arayüzü ile çalışma
./markdown-web
# Tarayıcıda http://localhost:8080 açın
```

## 🏗️ Proje Yapısı

```
markdown-to-html/
├── 📁 cmd/
│   ├── main.go              # 🖥️ Ana CLI giriş noktası
│   └── web/
│       └── main.go          # 🌐 Web arayüzü sunucusu
├── 📁 internal/
│   ├── converter/
│   │   ├── converter.go     # 🔄 Markdown → HTML dönüştürücü
│   │   └── pdf.go          # 📄 PDF dönüştürücü
│   └── utils/
│       └── file.go          # 📂 Dosya işlemleri yardımcıları
├── 📁 web/
│   ├── templates/
│   │   └── index.html       # 🎨 Web arayüzü template'i
│   └── static/
│       └── style.css        # 🎨 CSS stilleri
├── 📁 examples/
│   └── sample.md            # 📝 Örnek markdown dosyası
├── 📁 output/               # 📤 Çıktı dosyaları
├── 📄 go.mod               # 📦 Go modül dosyası
├── 📄 go.sum               # 🔒 Bağımlılık kontrolü
├── 📄 build.bat            # 🔨 Build script (Windows)
├── 📄 run.bat              # ▶️ Run script (Windows)
├── 📄 web.bat              # 🌐 Web arayüzü script (Windows)
├── 📄 LICENSE              # 📜 Lisans dosyası
└── 📄 README.md            # 📖 Bu dosya
```

## 🔧 Teknolojiler

| Teknoloji | Açıklama |
|-----------|----------|
| **Go 1.21+** | Ana programlama dili |
| **Goldmark** | Markdown parsing kütüphanesi |
| **Cobra** | CLI framework |
| **Bootstrap 5** | CSS framework |
| **Prism.js** | Syntax highlighting |
| **wkhtmltopdf** | PDF dönüştürme |
| **Font Awesome** | İkon kütüphanesi |

## 📝 Desteklenen Markdown Özellikleri

| Özellik | Durum | Açıklama |
|---------|-------|----------|
| **Başlıklar (H1-H6)** | ✅ | Tüm başlık seviyeleri |
| **Paragraflar** | ✅ | Temel metin formatlaması |
| **Kalın/İtalik** | ✅ | **bold**, *italic* |
| **Listeler** | ✅ | Sıralı ve sırasız listeler |
| **Bağlantılar** | ✅ | [link](url) formatı |
| **Resimler** | ✅ | ![alt](url) formatı |
| **Kod Blokları** | ✅ | Syntax highlighting ile |
| **Satır İçi Kod** | ✅ | `code` formatı |
| **Tablolar** | ✅ | Markdown tablo formatı |
| **Alıntılar** | ✅ | > quote formatı |
| **Yatay Çizgiler** | ✅ | --- formatı |
| **Görev Listeleri** | ✅ | - [x] formatı |
| **GitHub Flavored Markdown** | ✅ | GFM uzantıları |
| **Otomatik Başlık ID'leri** | ✅ | Başlık linkleri |

## 🎨 Tema Özellikleri

### 🌞 Light Theme (Varsayılan)
- **Temiz ve modern tasarım**
- **Beyaz arka plan** (#ffffff)
- **Koyu metin rengi** (#212529)
- **Bootstrap 5 bileşenleri**
- **Syntax highlighting** (Prism.js)

### 🌙 Dark Theme
- **Koyu arka plan** (#212529)
- **Açık metin rengi** (#f8f9fa)
- **Koyu kartlar** (#343a40)
- **Göz yorgunluğunu azaltan renkler**
- **Dark mode syntax highlighting**

## ✅ Tamamlanan Özellikler

- [x] **CLI Interface** - Komut satırı arayüzü
- [x] **Web Interface** - Modern web arayüzü
- [x] **PDF Support** - Markdown'dan PDF oluşturma
- [x] **Drag & Drop** - Dosya sürükle-bırak
- [x] **Real-time Preview** - Gerçek zamanlı önizleme
- [x] **Dark/Light Theme** - Tema desteği
- [x] **Syntax Highlighting** - Kod renklendirme
- [x] **Responsive Design** - Mobil uyumlu tasarım
- [x] **Auto-naming** - Otomatik dosya isimlendirme
- [x] **Error Handling** - Hata yönetimi

## 🚧 Gelecek Özellikler

- [ ] **Custom CSS Themes** - Özel CSS tema desteği
- [ ] **Watch Mode** - Dosya değişikliklerini izleme
- [ ] **Math Support** - Matematiksel formül desteği (KaTeX)
- [ ] **Multi-language** - Çoklu dil desteği
- [ ] **API Endpoints** - REST API
- [ ] **Docker Support** - Docker container desteği
- [ ] **Batch Processing** - Toplu dosya işleme
- [ ] **Export Formats** - Diğer formatlar (DOCX, RTF)

## 🤝 Katkıda Bulunma

Bu projeye katkıda bulunmak istiyorsanız:

1. **Fork yapın** - Projeyi kendi hesabınıza kopyalayın
2. **Feature branch oluşturun** - `git checkout -b feature/amazing-feature`
3. **Değişikliklerinizi commit edin** - `git commit -m 'Add amazing feature'`
4. **Branch'inizi push edin** - `git push origin feature/amazing-feature`
5. **Pull Request oluşturun** - Değişikliklerinizi gönderin

### 🐛 Bug Report

Bir hata bulduysanız, lütfen [Issue](https://github.com/BT-maker/markdown-to-html&PDF/issues) açın.

### 💡 Feature Request

Yeni özellik önerileriniz için [Issue](https://github.com/BT-maker/markdown-to-html&PDF/issues) açın.



## 🙏 Teşekkürler

| Kütüphane | Açıklama |
|-----------|----------|
| [Goldmark](https://github.com/yuin/goldmark) | Markdown parsing kütüphanesi |
| [Cobra](https://github.com/spf13/cobra) | CLI framework |
| [Bootstrap](https://getbootstrap.com/) | CSS framework |
| [Prism.js](https://prismjs.com/) | Syntax highlighting |
| [wkhtmltopdf](https://wkhtmltopdf.org/) | PDF dönüştürme |
| [Font Awesome](https://fontawesome.com/) | İkon kütüphanesi |

## 📞 İletişim

- **GitHub Issues**: [Issue Açın](https://github.com/yourusername/markdown-to-html/issues)
- **Email**: bahattok5@gmail.com
- **Website**: https://bt-maker.github.io/Portfolio/

## ⭐ Yıldız Verin

Bu projeyi beğendiyseniz, ⭐ yıldız vermeyi unutmayın!

---

**Not**: Bu proje eğitim amaçlı geliştirilmiştir. 🚀
# markdown-to-html-PDF
