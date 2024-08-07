# Book Management System

## Proje Açıklaması

Book Management System, kullanıcıların kitapları yönetebileceği, kitaplara inceleme ve puan verebileceği, okudukları kitapları kaydedebileceği bir uygulamadır. Proje, kullanıcı yönetimi, güvenlik, ve CRUD işlemleri gibi temel işlevleri içerir.

## Gereksinimler

- Go 1.18 veya daha yeni bir sürüm
- SQLite (varsayılan veritabanı)

## Kurulum ve Çalıştırma

1. **Proje:**

   ```bash
   git clone git@github.com:mertburgu/book-management-system.git
   cd book-management-system

2. **Go Modüllerini Yükleme:**

   ```bash
   go mod download

3. **Veritabanını Başlatın:**

    Veritabanı otomatik olarak başlatılır ve book_management.db dosyası proje kök dizininde oluşturulur.


4. **Proje Yapılandırmasını Yapın:**

    config/database.go dosyasında veritabanı bağlantı ayarlarını kontrol edin.


5. **Projeyi Çalıştırma:**

   ```bash
   go run cmd/main.go

6. **Kullanım**

   
   Kullanıcı Kaydı: POST /register endpoint'ini kullanarak yeni kullanıcılar oluşturabilirsiniz.
    
   
   Kullanıcı Girişi: POST /login endpoint'ini kullanarak kullanıcılar giriş yapabilir ve JWT token alabilir.
    
   
   Kitap Ekleme: POST /books endpoint'ini kullanarak kitap ekleyebilirsiniz (JWT token gerekli).

   
   Kitap Listeleme: GET /books endpoint'ini kullanarak kitapları listeleyebilirsiniz (JWT token gerekli).
      
   
   Kitap Güncelleme: PUT /books/:id endpoint'ini kullanarak bir kitabı güncelleyebilirsiniz (JWT token gerekli).
   
   
   Kitap Silme: DELETE /books/:id endpoint'ini kullanarak bir kitabı silebilirsiniz (JWT token gerekli).

curl -X POST http://localhost:8082/login -H "Content-Type: application/json" -d '{"username":"testuser", "password":"password"}'

curl -X POST http://localhost:8082/books -H "Authorization: Bearer YOUR_TOKEN" -H "Content-Type: application/json" -d '{"title":"The Great Gatsby", "author":"F. Scott Fitzgerald"}'
