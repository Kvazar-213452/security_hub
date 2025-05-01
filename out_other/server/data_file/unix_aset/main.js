function send_data() {
    const password = $('.input').val();
    console.log(password)
  
    if (!password) {
        message_window("Please enter your password!");
        return;
    }
  
    $.ajax({
      url: '/search',
      method: 'POST',
      contentType: 'application/json',
      data: JSON.stringify({ searchPassword: password }),
      success: function (data) {
        const link = $('<a>')
          .attr('href', `data:application/octet-stream;base64,${data.file}`)
          .attr('download', data.name)
          .appendTo('body');
  
        link[0].click();
        link.remove();
  
        message_window("File uploaded successfully!");
      },
      error: function (xhr) {
        if (xhr.status === 404) {
            message_window("File not found or password is incorrect.");
        } else {
            message_window("An error occurred. Try again");
        }
      },
    });
  }

function message_window(content) {
    const $block = $('<div class="animatedBlock hide"></div>').text(content);
    $('body').append($block);

    setTimeout(() => {
        $block.removeClass('hide').addClass('show');
    }, 0);

    setTimeout(() => {
        $block.removeClass('show').addClass('hide');

        setTimeout(() => {
            $block.remove();
        }, 1000);
    }, 3000);
}
  