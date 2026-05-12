<%*
const file = tp.file;
const currentDate = tp.date.now("YYYY-MM-DD");
await app.fileManager.processFrontMatter(file.find_tfile(file.path(true)), (fm) => {
  fm.status = "done";
  fm.due = currentDate;
});
%>
