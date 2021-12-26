import fs from "fs";
import { parse } from "csv-parse";

export async function processFile(inFile, outFile, processItem) {
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
