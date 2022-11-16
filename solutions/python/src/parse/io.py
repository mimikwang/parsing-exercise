import os
import json
from typing import Tuple

from parse import output


def get_files(directory: str) -> Tuple[str, dict]:
    """Return the file name and data of json files in directory"""
    for file in os.listdir(directory):
        if not file.endswith(".json"):
            continue
        with open(os.path.join(directory, file)) as f:
            data = json.load(f)
            yield file, data


def write_output(directory: str, file_name: str, out: output.Output):
    """Write output"""
    with open(os.path.join(directory, file_name), "w") as f:
        f.write(out.json())
