# Código adaptado de https://github.com/tony-josi/pvd_steganography

import random
import string
from pathlib import Path
from .pvd_lib import pvd_lib

def embed_message_pvd(image_path, message, output_path):
    print("called")
    pvd_obj = pvd_lib()
    temp_file = Path.cwd() / 'temp_file'
    if temp_file.exists():
        temp_file.unlink()
    with open('temp_file', 'w') as file:
        file.write(message)
    temp_file_str = 'temp_file'
    pvd_obj.embed_data(image_path, temp_file_str, output_path)


# target_byte_count = 18288
# embed_str = ''.join(random.choice(string.ascii_letters) for x in range(target_byte_count))
# embed_message_pvd('/media/external/code/python/pvd_steganography/10001.png', embed_str, '/media/external/code/python/pvd_steganography/stego.png')