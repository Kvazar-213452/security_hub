const express = require('express');
const bodyParser = require('body-parser');
const multer = require('multer');
const fs = require('fs-extra');
const path = require('path');
const cors = require('cors');

// server/data_file/index.js

const { get_time, readDB, writeDB, deleteFileAfterDelay } = require('./main_com/func');

const app = express();
const PORT = 3000;

app.use(cors());

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));

app.use('/unix_aset', express.static(path.join(__dirname, 'unix_aset')));

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

app.get('/', (req, res) => {
  res.send(`unix server`);
});

app.get('/main', (req, res) => {
  res.render('index');
});

app.post('/get_how_many', (req, res) => {
  const dbPath = path.join(__dirname, './db.json');
  if (!fs.existsSync(dbPath)) {
      return res.status(404).send('Database not found');
  }

  const dbData = JSON.parse(fs.readFileSync(dbPath, 'utf8'));
  res.send({ count: dbData.length });
});

app.post('/server_unix', (req, res) => {
    res.send(`1`);
});

app.post('/upload', upload.single('file'), (req, res) => {
  const password = req.body.password;
  const originalName = req.file.originalname;
  const newFileName = `renamed-${Date.now()}${path.extname(originalName)}`;
  const newFilePath = path.join(__dirname, 'static/file', newFileName);
  const currentTime = get_time();

  fs.rename(req.file.path, newFilePath, (err) => {
    if (err) {
      return res.status(500).send('0.');
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
      const fileContent = fs.readFileSync(filePath);
      return res.send({
        file: fileContent.toString('base64'),
        name: foundEntry.fileName
      });
    } else {
      return res.status(404).send('1');
    }
  }

  res.status(404).send('0');
});

app.listen(PORT, () => {
  console.log(`start http://localhost:${PORT}`);
});
