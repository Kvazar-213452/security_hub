$('#submit-btn').click(function () {
  const fileInput = $('#file')[0];
  const password = $('#password').val();
  const file = fileInput.files[0];

  if (!file || !password) {
    if (lang_global === "uk") {
      message_window('Заповніть всі поля'); 
    } else if (lang_global === "en") {
        message_window('Fill in all fields'); 
    }
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
      if (lang_global === "uk") {
        message_window('Файл успішно завантажено'); 
      } else if (lang_global === "en") {
          message_window('File uploaded successfully'); 
      }
    }
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
        if (lang_global === "uk") {
          message_window('Файл не знайдено'); 
        } else if (lang_global === "en") {
            message_window('File not found'); 
        }
      }
    }
  });
});