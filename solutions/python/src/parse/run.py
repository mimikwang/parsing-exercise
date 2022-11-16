import argparse
import logging

from parse import input, io

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class Args:
    def __init__(self):
        parser = argparse.ArgumentParser(description="parse-exercise python implementation")
        parser.add_argument("input_dir")
        parser.add_argument("output_dir")
        args = parser.parse_args()

        self.input_dir = args.input_dir
        self.output_dir = args.output_dir


def read_input(data: str) -> input.Parse:
    formats = [
        input.FormatA,
        input.FormatB
    ]
    for format in formats:
        try:
            p = format.parse_obj(data)
            return p
        except Exception as e:
            continue
    raise ValueError("unknown format")


def process_input(p: input.Parse, file_name: str, output_dir: str):
    out = p.parse()
    logger.info(f"Writing {file_name}")
    io.write_output(output_dir, file_name, out)


def run():
    args = Args()
    for file_name, data in io.get_files(args.input_dir):
        logger.info(f"Processing {file_name}")
        try:
            p = read_input(data)
            process_input(p, file_name, args.output_dir)
        except Exception as e:
            logger.info(f"Cannot process {file_name}: {e}")
