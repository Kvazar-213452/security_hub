const express = require('express');
const cors = require('cors');

const app = express();
const port = 3000;

app.use(cors());


app.get('/data', (req, res) => {
    res.json({
        message: 'Version 6',
        desc: 'wefdvwfevdverf',
    });
});

app.listen(port, () => {
    console.log(`Сервер запущено на http://localhost:${port}`);
});
