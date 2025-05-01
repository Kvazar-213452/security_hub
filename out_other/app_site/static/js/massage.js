$('#sendButton').click(function() {
    const inputData = $('#myInput').val();
    const textData = $('#myTextarea').val();
    $.ajax({
        url: '/server/save_data',
        type: 'POST',
        data: { title: inputData, text: textData },
        success: function(response) {
            fetchData();
        },
        error: function() {
            alert('Error saving data.');
        }
    });
});
function fetchData() {
    $.ajax({
        url: '/db',
        type: 'GET',
        dataType: 'json',
        success: function(data) {
            $('#weg4gwevds').html(null);
            
            for (let i = 0; i < data.length; i++) {
                let text = `
                    <div class="cwd23103100c">
                        <p class="e46fetgr">${data[i]['title']}</p>
                        <p class="dqw32wert54eg">${data[i]['text']}</p>
                    </div>
                `;
                $('#weg4gwevds').append(text);
            }
        },
        error: function() {
            console.error('Error fetching data from db.json');
        }
    });
}