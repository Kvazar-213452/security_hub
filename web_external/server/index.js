const express = require('express');
const bodyParser = require('body-parser');
const multer = require('multer');
const fs = require('fs-extra');
const path = require('path');
const cors = require('cors');

const app = express();
const PORT = 3000;

function get_time() {
    const now = new Date();

    now.setHours(now.getHours() + 1);
  
    const formattedTime = now.toLocaleTimeString('en-GB', {
      hour12: false,
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
    });
  
    return formattedTime;
};

const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    const uploadPath = path.join(__dirname, 'static/file');
    fs.ensureDirSync(uploadPath);
    cb(null, uploadPath);
  },
  filename: (req, file, cb) => {
    const uniqueSuffix = Date.now() + '-' + Math.round(Math.random() * 1E9);
    const extension = path.extname(file.originalname);
    cb(null, `${uniqueSuffix}${extension}`);
  },
});

const upload = multer({ storage });

app.use(bodyParser.urlencoded({ extended: true }));
app.use(bodyParser.json());
app.use(express.static('static'));

const readDB = () => {
  const dbPath = path.join(__dirname, 'db.json');
  if (!fs.existsSync(dbPath)) return [];
  return JSON.parse(fs.readFileSync(dbPath, 'utf8'));
};

const writeDB = (data) => {
  const dbPath = path.join(__dirname, 'db.json');
  fs.writeFileSync(dbPath, JSON.stringify(data, null, 2), 'utf8');
};

const deleteFileAfterDelay = (filePath, password) => {
  setTimeout(() => {
    if (fs.existsSync(filePath)) {
      fs.unlinkSync(filePath);
      console.log(`Файл ${filePath} видалено.`);
    }

    const db = readDB();
    const updatedDB = db.filter((entry) => entry.password !== password);
    writeDB(updatedDB);
    console.log(`Запис із паролем "${password}" видалено з db.json.`);
  }, 43200000);
};

app.get('/', (req, res) => {
    res.send(`unix server`);
});

app.post('/upload', upload.single('file'), (req, res) => {
  const password = req.body.password;
  const originalName = req.file.originalname;
  const newFileName = `renamed-${Date.now()}${path.extname(originalName)}`;
  const newFilePath = path.join(__dirname, 'static/file', newFileName);
  const currentTime = get_time();

  fs.rename(req.file.path, newFilePath, (err) => {
    if (err) {
      console.error('Помилка перейменування:', err);
      return res.status(500).send('Виникла помилка при обробці файлу.');
    }

    const db = readDB();
    db.push({ password, fileName: newFileName, time: currentTime });
    writeDB(db);

    deleteFileAfterDelay(newFilePath, password);

    res.send(`good`);
  });
});

app.post('/search', (req, res) => {
  const searchPassword = req.body.searchPassword;

  const db = readDB();
  const foundEntry = db.find((entry) => entry.password === searchPassword);

  if (foundEntry) {
    const filePath = path.join(__dirname, 'static/file', foundEntry.fileName);
    if (fs.existsSync(filePath)) {
      return res.download(filePath);
    } else {
      return res.status(404).send('Файл не знайдено на сервері.');
    }
  }

  res.status(404).send('Пароль не знайдено.');
});

app.listen(PORT, () => {
  app.use(cors());
  console.log(`Сервер запущено на http://localhost:${PORT}`);
});