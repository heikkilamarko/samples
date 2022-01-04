import argparse
from processor import process_file
from person import process_person

parser = argparse.ArgumentParser()
parser.add_argument("input", help="input file")
parser.add_argument("-o", "--output", help="output file")
args = parser.parse_args()

if args.output is None:
    args.output = args.input + ".json"

process_file(args.input, args.output, process_person)
