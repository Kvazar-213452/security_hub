const fs = require('fs-extra');
const path = require('path');

// server/data_file/main_com/func.js

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
}

const readDB = () => {
  const dbPath = path.join(__dirname, '../db.json');
  if (!fs.existsSync(dbPath)) return [];
  return JSON.parse(fs.readFileSync(dbPath, 'utf8'));
};

const writeDB = (data) => {
  const dbPath = path.join(__dirname, '../db.json');
  fs.writeFileSync(dbPath, JSON.stringify(data, null, 2), 'utf8');
};

const deleteFileAfterDelay = (filePath, password) => {
  setTimeout(() => {
    if (fs.existsSync(filePath)) {
      fs.unlinkSync(filePath);
    }

    const db = readDB();
    const updatedDB = db.filter((entry) => entry.password !== password);
    writeDB(updatedDB);
  }, 43200000 * 2); // 12 hours in milliseconds
};

module.exports = {
  get_time,
  readDB,
  writeDB,
  deleteFileAfterDelay
};
