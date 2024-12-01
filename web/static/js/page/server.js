$('#submit-btn').click(function(e) {
  e.preventDefault();

  var file = $('#file')[0].files[0];
  var password = $('#password').val();

  if (file && password) {
    var formData = new FormData();
    formData.append('file', file);
    formData.append('password', password);

    $.ajax({
      url: 'http://localhost:3000/upload',
      type: 'POST',
      data: formData,
      processData: false,
      contentType: false,
      success: function(response) {
        if (lang_global === "uk") {
          message_window('Завантажено');
        } else if (lang_global === "en") {
            message_window('Downloaded');
        }
      },
      error: function(xhr, status, error) {
        if (lang_global === "uk") {
          message_window('Завантажено');
        } else if (lang_global === "en") {
            message_window('Downloaded');
        }
      }
    });
  } else {
    message_window()
    if (lang_global === "uk") {
      message_window('Будь ласка, виберіть файл та введіть пароль');
    } else if (lang_global === "en") {
        message_window('Please select a file and enter a password');
    }
  }
});
