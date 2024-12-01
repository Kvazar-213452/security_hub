$(document).ready(function() {
    // Завантаження файлу
    $('#uploadBtn').click(function() {
      var file = $('#file')[0].files[0]; // Отримуємо файл
      var password = $('#password').val(); // Отримуємо пароль
  
      if (!file || !password) {
        alert('Будь ласка, виберіть файл і введіть пароль.');
        return;
      }
  
      var formData = new FormData();
      formData.append('file', file);
      formData.append('password', password);
  
      $.ajax({
        url: 'http://localhost:3000/upload',
        type: 'POST',
        data: formData,
        contentType: false, // Вказуємо, що ми не хочемо відправляти заголовки contentType
        processData: false, // Не перетворюємо дані
        success: function(response) {
          $('#result').html('<p>Файл успішно завантажено!</p>');
        },
        error: function(xhr, status, error) {
          $('#result').html('<p>Помилка при завантаженні файлу.</p>');
        }
      });
    });
  
    // Пошук файлу за паролем
    $('#searchBtn').click(function() {
      var searchPassword = $('#searchPassword').val(); // Отримуємо пароль для пошуку
  
      if (!searchPassword) {
        alert('Будь ласка, введіть пароль для пошуку.');
        return;
      }
  
      $.ajax({
        url: 'http://localhost:3000/search',
        type: 'POST',
        data: { searchPassword: searchPassword },
        success: function(response) {
          if (response) {
            $('#result').html('<p>Файл знайдено! Завантажуємо...</p>');
          } else {
            $('#result').html('<p>Файл не знайдено.</p>');
          }
        },
        error: function(xhr, status, error) {
          $('#result').html('<p>Помилка при пошуку файлу.</p>');
        }
      });
    });
  });
  