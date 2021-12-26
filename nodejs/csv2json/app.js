import { Command } from "commander";
import { processFile } from "./processor.js";
import { processPerson } from "./person.js";

const { input, output } = parseFlags();

await processFile(input, output, processPerson);

function parseFlags() {
  const program = new Command();

  program.requiredOption("-i, --input <file>", "input file");
  program.option("-o, --output <file>", "output file");

  program.parse();

  const opts = program.opts();

  if (!opts.output) {
    opts.output = opts.input + ".json";
  }

  return opts;
}
