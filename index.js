// Just for looping and creating README.md
const fs = require('fs');
const path = require('path');

async function readDirectory(filePath) {
  const data = await fs.readdirSync(filePath);
  const folders = data.filter(item => fs.lstatSync(item).isDirectory());

  folders.map(folder => {
    const subFilePath = path.join(filePath, `/${folder}` + `/README.md`);
    fs.writeFile(subFilePath, '', err => {
      if (err) console.log(err);
      else {
        console.log(`File ${subFilePath} created`);
      }
    });
  });
}

readDirectory(path.resolve(__dirname));
