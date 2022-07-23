import os
from pathlib import Path
from PIL import Image

ORIGINAL_PATH = Path(__file__).parent / 'original'
OUTPUT_PATH = Path(__file__).parent / 'original_jpg'

Path.mkdir(OUTPUT_PATH, exist_ok=True)

for file in ORIGINAL_PATH.iterdir():
    basename = file.stem
    ext = file.suffix
    if ext != '.bmp':
        continue
    print(f"converting file {file} ({basename}) to jpg")
    img = Image.open(file)
    output_path = OUTPUT_PATH / f'{basename}.jpg'
    img.save(output_path)
