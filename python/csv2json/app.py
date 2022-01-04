import argparse
from processor import processFile
from person import processPerson

parser = argparse.ArgumentParser()
parser.add_argument("input", help="input file")
parser.add_argument("-o", "--output", help="output file")
args = parser.parse_args()

if args.output is None:
    args.output = args.input + ".json"

processFile(args.input, args.output, processPerson)
