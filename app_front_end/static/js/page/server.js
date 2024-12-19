// <!DOCTYPE html>
// <html lang="uk">
// <head>
//   <meta charset="UTF-8">
//   <meta name="viewport" content="width=device-width, initial-scale=1.0">
//   <title>File Upload and Search</title>
//   <style>
//     .padoge {
//       padding: 20px;
//       background: #f0f0f0;
//     }
//     .file-input-container {
//       margin-bottom: 20px;
//     }
//     .ooo33ff3cv2122 {
//       padding: 10px;
//       margin: 5px 0;
//       width: 100%;
//     }
//     .dqw21fvvvv {
//       background-color: #4CAF50;
//       color: white;
//       padding: 10px 20px;
//       border: none;
//       cursor: pointer;
//     }
//   </style>
// </head>
// <body>

//   <!-- Форма для завантаження файлу -->
//   <div class="padoge">
//     <h2>Завантажити файл</h2>
//     <div class="file-input-container">
//       <input type="file" name="file" id="file" required>
//       <label for="file" class="file-input-button">Виберіть файл</label>
//     </div>
//     <label for="password">Пароль для доступу до файлу:</label>
//     <input class="ooo33ff3cv2122" type="text" name="password" id="password" required>
//     <br><br>
//     <button class="dqw21fvvvv" type="button" id="submit-btn">Завантажити</button>
//   </div>

//   <br>

//   <!-- Форма для пошуку файлу -->
//   <div class="padoge">
//     <h2>Пошук файлу за паролем</h2>
//     <form id="searchForm">
//       <label for="searchPassword">Пароль:</label>
//       <input class="ooo33ff3cv2122" type="text" name="searchPassword" id="searchPassword" required>
//       <br><br>
//       <button class="dqw21fvvvv" type="submit">Знайти файл</button>
//     </form>
//     <div id="searchResult"></div>
//   </div>

//   <script>
//     // Завантаження файлу
//     document.getElementById('submit-btn').addEventListener('click', function() {
//       const fileInput = document.getElementById('file');
//       const password = document.getElementById('password').value;
//       const file = fileInput.files[0];

//       if (!file || !password) {
//         alert('Будь ласка, виберіть файл і введіть пароль.');
//         return;
//       }

//       const formData = new FormData();
//       formData.append('file', file);
//       formData.append('password', password);

//       fetch('http://localhost:3000/upload', {
//         method: 'POST',
//         body: formData
//       })
//       .then(response => response.text())
//       .then(data => {
//         alert('Файл успішно завантажено');
//       })
//       .catch(error => {
//         console.error('Error:', error);
//         alert('Сталася помилка при завантаженні файлу');
//       });
//     });


    
// document.getElementById('searchForm').addEventListener('submit', function(event) {
//   event.preventDefault();
//   const searchPassword = document.getElementById('searchPassword').value;

//   fetch('http://localhost:3000/search', {
//     method: 'POST',
//     headers: {
//       'Content-Type': 'application/json'
//     },
//     body: JSON.stringify({ searchPassword })
//   })
//   .then(response => {
//     // Перевіряємо тип відповіді
//     const contentType = response.headers.get('Content-Type');
    
//     if (contentType && contentType.includes('application/json')) {
//       // Якщо сервер повертає JSON, обробляємо його
//       return response.json();
//     } else {
//       // Якщо це бінарний файл (наприклад, PDF), обробляємо його як blob
//       return response.blob();
//     }
//   })
//   .then(data => {
//     const contentType = data instanceof Blob ? 'blob' : 'json';

//     if (contentType === 'json') {
//       // Якщо сервер повернув JSON, отримуємо шлях до файлу і ім'я
//       const filePath = data.file;
//       const filename = data.name;

//       // Завантажуємо файл
//       fetch(filePath)
//         .then(fileResponse => {
//           if (!fileResponse.ok) {
//             throw new Error('Не вдалося завантажити файл');
//           }
//           return fileResponse.blob();  // Отримуємо файл як blob
//         })
//         .then(fileBlob => {
//           const fileURL = URL.createObjectURL(fileBlob);
//           const link = document.createElement('a');
//           link.href = fileURL;
//           link.download = filename;  // Встановлюємо ім'я файлу для завантаження
//           link.click();
//         })
//         .catch(error => {
//           document.getElementById('searchResult').innerText = error.message;
//           console.error('Error:', error);
//         });
//     } else if (contentType === 'blob') {
//       // Якщо відповідь - бінарний файл (наприклад, PDF), створюємо URL для скачування
//       const fileURL = URL.createObjectURL(data);
//       const link = document.createElement('a');
//       link.href = fileURL;
//       link.download = 'file';  // Можна задати конкретне ім'я файлу
//       link.click();
//     }
//   })
//   .catch(error => {
//     document.getElementById('searchResult').innerText = error.message;
//     console.error('Error:', error);
//   });
// });


//   </script>

// </body>
// </html>





$('#submit-btn').click(function() {
  const fileInput = $('#file')[0];
  const password = $('#password').val();
  const file = fileInput.files[0];

  if (!file || !password) {
    alert('Будь ласка, виберіть файл і введіть пароль.');
    return;
  }

  const formData = new FormData();
  formData.append('file', file);
  formData.append('password', password);

  $.ajax({
    url: 'http://localhost:3000/upload',
    type: 'POST',
    data: formData,
    contentType: false,
    processData: false,
    success: function(response) {
      alert('Файл успішно завантажено');
    },
    error: function(error) {
      console.error('Error:', error);
      alert('Сталася помилка при завантаженні файлу');
    }
  });
});


$('#searchForm').submit(function(event) {
  event.preventDefault();
  const searchPassword = $('#searchPassword').val();

  $.ajax({
    url: 'http://localhost:3000/search',
    type: 'POST',
    contentType: 'application/json',
    data: JSON.stringify({ searchPassword }),
    success: function(response) {
      const contentType = response.headers['Content-Type'];

      if (contentType && contentType.includes('application/json')) {
        const filePath = response.file;
        const filename = response.name;

        $.ajax({
          url: filePath,
          type: 'GET',
          success: function(fileResponse) {
            const fileURL = URL.createObjectURL(fileResponse);
            const link = document.createElement('a');
            link.href = fileURL;
            link.download = filename;
            link.click();
          },
          error: function(error) {
            $('#searchResult').text('Не вдалося завантажити файл');
            console.error('Error:', error);
          }
        });
      } else if (contentType && contentType.includes('blob')) {
        const fileURL = URL.createObjectURL(response);
        const link = document.createElement('a');
        link.href = fileURL;
        link.download = 'file';  // Можна задати конкретне ім'я файлу
        link.click();
      }
    },
    error: function(error) {
      $('#searchResult').text(error.message);
      console.error('Error:', error);
    }
  });
});
