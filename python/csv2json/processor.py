import csv
import json


def processFile(in_file_path, out_file_path, processItem):
    with open(in_file_path, 'r') as in_file:
        reader = csv.DictReader(in_file)
        with open(out_file_path, 'w', newline='') as out_file:
            for in_data in reader:
                out_data = processItem(in_data)
                json.dump(out_data, out_file)
                out_file.write('\n')
