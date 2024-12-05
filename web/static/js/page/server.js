$('#submit-btn').click(function(e) {
  e.preventDefault();

  let file = $('#file')[0].files[0];
  let password = $('#password').val();

  if (file && password) {
    let formData = new FormData();
    formData.append('file', file);
    formData.append('password', password);

    $.ajax({
      url: '/server_upload',
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

$('#searchForm').submit(function (event) {
  event.preventDefault();

  let searchPassword = $('#searchPassword').val();

  let formData = new FormData();
  formData.append('searchPassword', searchPassword);

  $.ajax({
    url: '/server_search',
    type: 'POST',
    data: formData,
    processData: false,
    contentType: false,
    success: function (response, status, xhr) {
      let filename = xhr.getResponseHeader('Content-Disposition');
      
      if (filename && filename.indexOf('attachment') !== -1) {
        let blob = new Blob([response], { type: 'application/octet-stream' });
        let link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = filename.split('filename=')[1];
        link.click();
      }
    },
    error: function (xhr, status, error) {
      console.log('Error:', error);
    }
  });
});