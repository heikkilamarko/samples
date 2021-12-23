import fs from "fs";
import { parse } from "csv-parse";

async function processFile(inFile, outFile) {
  const inStream = fs.createReadStream(inFile).pipe(
    parse({
      columns: true,
      relax_quotes: true,
    })
  );

  const outStream = fs.createWriteStream(outFile);

  for await (const inItem of inStream) {
    const outItem = processItem(inItem);
    outStream.write(JSON.stringify(outItem) + "\n");
  }
}

function processItem(item) {
  // TODO: Add transformation logic here.
  return item;
}

await processFile("in.csv", "out.json");
