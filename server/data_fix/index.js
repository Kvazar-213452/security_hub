const express = require("express");
const path = require("path");
const fs = require("fs");
const cors = require("cors");

const app = express();
const PORT = 23113;

const staticDir = path.join(__dirname, "public");

app.use(cors());

app.use(express.static(staticDir));

app.get("/font", (req, res) => {
    const fontPath = path.join(staticDir, "Minecraft_1.1.ttf");
    
    fs.readFile(fontPath, (err, data) => {
        if (err) {
            res.status(500).send("Помилка при завантаженні шрифта");
        } else {
            res.setHeader("Content-Type", "font/ttf");
            res.send(data);
        }
    });
});

app.listen(PORT, () => {
    console.log(`Сервер запущено на http://localhost:${PORT}`);
});
