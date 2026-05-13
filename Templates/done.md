<%*
const file = tp.file;
const currentDate = tp.date.now("YYYY-MM-DD");
await app.fileManager.processFrontMatter(file.find_tfile(file.path(true)), (fm) => {
  fm.status = "done";
  fm.due = currentDate;
});

// Запуск publish в фоне (не блокирует Obsidian)
const { exec } = require('child_process');
exec('/home/rudra/bin/publish-quartz.sh', (err, stdout, stderr) => {
  if (err) {
    new Notice(`Publish failed: ${err.message}`, 5000);
  } else {
    new Notice('✓ Quartz published', 3000);
  }
});
%>
