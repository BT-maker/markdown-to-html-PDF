@echo off
echo 🌐 Markdown to HTML/PDF Web Arayüzü Başlatılıyor...
echo.
echo 📝 Özellikler:
echo    - Drag & Drop dosya yükleme
echo    - Gerçek zamanlı önizleme
echo    - HTML ve PDF dönüştürme
echo    - Dark/Light tema desteği
echo    - Syntax highlighting
echo.
echo 🚀 Web arayüzü: http://localhost:8080
echo.
echo Durdurmak için Ctrl+C tuşlayın
echo.

go run cmd/web/main.go

pause
