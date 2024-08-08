# Book Management System

## Proje Açıklaması

Book Management System, kullanıcıların kitaplara inceleme ve puan ekleyebildiği, kendi kütüphanelerini düzenleyip public/private olarak listeleyebildiği bir uygulamadır. 
Editör ve admin rolleri ile de kitapların yönetimi yapılmaktadır.
Proje, kullanıcı yönetimi, güvenlik, ve CRUD işlemleri gibi temel işlevleri içerir.

## Gereksinimler

- Go 1.22
- PostgreSQL 14.12
- Docker version 27.1.1

## Kurulum ve Çalıştırma


1. **Proje:**

   ```bash
   git clone git@github.com:mertburgu/book-management-system.git
   cd book-management-system


2. **Go Modüllerini Yükleme:**

   ```bash
   go mod download


3. **Veritabanını Başlatın:**

    Veritabanı docker ile otomatik olarak başlatılır 

    database:book_management

    user:user
 
    password:password
 
    host:localhost

    port:5432


4. **Projeyi Çalıştırma:**

   ```bash
   docker-compose build
   docker-compose up -d


5. **Kullanım**

   Kullanıcı Kaydı: POST /register endpoint'ini kullanarak yeni kullanıcılar oluşturabilirsiniz.
   ```bash
   POST
   http://localhost:8082/register
   
   Request Headers
   Content-Type application/vnd.api+json
   Body raw json
   {
       "username": "testuser",
       "password": "testpassword",
       "email": "testuser@example.com",
       "name": "Test",
       "surname": "User",
       "phone": "1234567890"
   }
   ```
   
   Kullanıcı Girişi: POST /login endpoint'ini kullanarak kullanıcılar giriş yapabilir ve JWT token alabilir.
   
   ```bash
      POST
      http://localhost:8082/login
   
      Request Headers
      Content-Type application/vnd.api+json
      Body raw json
      {
      "username": "testuser",
      "password": "testpassword",
      }
      
      RESPONSE
      {
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxOTc4MDUsImlkIjoiNTJkZGUwM2QtMDk1My00ZTFiLWI2OTMtOWIxMjhkZjBhMWIyIiwicm9sZSI6InVzZXIifQ.kl9aYsKZw1h33w7GpTNpD2Y1Q9OelylWrcd2aGegOXE"
      }
   ```
   
[comment]: <> (   Kitap Ekleme: POST /books endpoint'ini kullanarak kitap ekleyebilirsiniz &#40;JWT token gerekli&#41;.)

   
[comment]: <> (   Kitap Listeleme: GET /books endpoint'ini kullanarak kitapları listeleyebilirsiniz &#40;JWT token gerekli&#41;.)
      
   
[comment]: <> (   Kitap Güncelleme: PUT /books/:id endpoint'ini kullanarak bir kitabı güncelleyebilirsiniz &#40;JWT token gerekli&#41;.)
   
   
[comment]: <> (   Kitap Silme: DELETE /books/:id endpoint'ini kullanarak bir kitabı silebilirsiniz &#40;JWT token gerekli&#41;.)
