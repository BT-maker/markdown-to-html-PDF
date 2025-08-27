# Markdown to HTML Converter

Bu proje, Markdown dosyalarını HTML'e dönüştüren bir Go uygulamasıdır.

## Özellikler

- **Bootstrap Styling**: Modern ve responsive tasarım
- **Syntax Highlighting**: Kod blokları için renklendirme
- **Dark/Light Theme**: İki farklı tema seçeneği
- **CLI Interface**: Kolay kullanım için komut satırı arayüzü

## Kullanım

```bash
# Temel kullanım
go run cmd/main.go input.md output.html

# Otomatik isimlendirme
go run cmd/main.go input.md

# Önizleme
go run cmd/main.go input.md --preview

# Dark tema
go run cmd/main.go input.md --theme dark
```

## Kod Örneği

```go
package main

import (
    "fmt"
    "markdown-to-html/internal/converter"
)

func main() {
    markdown := "# Merhaba Dünya\nBu bir **markdown** örneğidir."
    html, err := converter.ConvertToHTML(markdown, "light")
    if err != nil {
        panic(err)
    }
    fmt.Println(html)
}
```

## Tablo Örneği

| Özellik | Açıklama | Durum |
|---------|----------|-------|
| Markdown Parsing | Goldmark kütüphanesi | ✅ |
| Bootstrap CSS | Responsive tasarım | ✅ |
| Syntax Highlighting | Prism.js | ✅ |
| Dark Theme | Koyu tema desteği | ✅ |

## Liste Örnekleri

### Sıralı Liste
1. İlk adım
2. İkinci adım
3. Üçüncü adım

### Sırasız Liste
- Elma
- Armut
- Muz
  - Sarı muz
  - Yeşil muz

## Alıntı Örneği

> Bu bir alıntı örneğidir. Markdown'da `>` işareti ile alıntı yapabilirsiniz.

## Bağlantı ve Resim

[GitHub](https://github.com) bağlantısı örneği.

<!-- Resim örneği (internet bağlantısı gerektirdiği için yorum satırına alındı) -->
<!-- ![Örnek Resim](https://via.placeholder.com/300x200) -->

## Kod Satır İçi

Bu satırda `kod` örneği bulunmaktadır.

## Yatay Çizgi

---

Bu çizginin altında kalan içerik.

## Matematiksel İfadeler

Euler'ın formülü: $e^{i\pi} + 1 = 0$

## Görev Listesi

- [x] Markdown parsing
- [x] HTML çıktısı
- [x] Bootstrap styling
- [ ] Çoklu dosya desteği
- [ ] Web arayüzü

## Sonuç

Bu proje, Markdown dosyalarını güzel HTML sayfalarına dönüştürmek için geliştirilmiştir. Go dilinin gücü ve modern web teknolojileri birleştirilerek kullanıcı dostu bir araç oluşturulmuştur.
