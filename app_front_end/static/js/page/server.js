$('#submit-btn').click(function () {
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
    url: '/post_file_server',
    type: 'POST',
    data: formData,
    processData: false,
    contentType: false,
    success: function () {
      alert('Файл успішно завантажено');
    },
    error: function (xhr, status, error) {
      console.error('Error:', error);
      alert('Сталася помилка при завантаженні файлу');
    },
  });
});

$('#search-btn').click(function () {
  const searchPassword = $('#searchPassword').val();

  $.ajax({
    url: '/search_server',
    type: 'POST',
    contentType: 'application/json',
    data: JSON.stringify({ searchPassword }),
    success: function (data) {
      console.log(data);
      if (data.file && data.name) {
        const binaryData = atob(data.file);
        const byteArray = new Uint8Array(binaryData.length);
        for (let i = 0; i < binaryData.length; i++) {
          byteArray[i] = binaryData.charCodeAt(i);
        }
        const blob = new Blob([byteArray]);

        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = data.name;
        link.click();
      } else {
        $('#searchResult').text('Файл не знайдено.');
      }
    },
    error: function (xhr, status, error) {
      console.error('Error:', error);
      $('#searchResult').text('Сталася помилка при пошуку файлу.');
    },
  });
});