import numpy as np
from PIL import Image
from luma.led_matrix.device import max7219
from luma.core.interface.serial import spi, noop
import time
import cv2

# create LED matrix device 
serial = spi(port=0, device=0, gpio=noop())
device = max7219(serial, height=8, width=8, rotate=0)
uint8='uint8'

# Clear display
device.clear()

# Set LED brightness (int 0-255)
device.contrast(255)


def read_img(img_path):
    img = cv2.imread(img_path, cv2.IMREAD_COLOR)
    return cv2.resize(img, (8, 8))


def img_to_bw_array(img_path):
    originalImage = read_img(img_path)
    grayImage = cv2.cvtColor(originalImage, cv2.COLOR_BGR2GRAY)
    (thresh, blackAndWhiteImage) = cv2.threshold(grayImage, 127, 255, cv2.THRESH_BINARY)
    return blackAndWhiteImage

def matrix_display(matrix, device=device, invert=False):
    """Display a 2D numpy array on the LED matrix.

    This is a wrapper function that converts a 2D numpy
    matrix into a 1-bit B/W image and then displays it
    on the LED matrix.
    
    Parameters
    ----------
    device : luma.led_matrix.device.max7219
        A Luma LED matrix device object.
    matrix: uint8 numpy.ndarray
        A 2D binary numpy array.
        
    """
    # Convert "1" values in binary matrix to 255
    matrix_g = np.where(matrix==1, 255, matrix)

    if invert:
        matrix_g = np.invert(matrix_g)

    # Creating a 1-bit B/W image directly from a binary
    # numpy array does not work for some reason (maybe
    # a bug in Pillow).
    # Create greyscale PIL image from a 2D numpy array
    image_g = Image.fromarray(matrix_g, mode='L')
    
    # Convert greyscale image to 1-bit B/W
    image_bw = image_g.convert('1')
    
    # Display image on LED matrix
    device.display(image_bw)
    
def string_to_array(string):
    if len(string) != 64:
        raise ValueError("String is not 64 characters long")
    else:
        array = []
        for i in range(0, 64, 8):
            row = [y for y in string[i:i+8]]
            array.append(row)
    return np.array(array, dtype='uint8')

def image_display(image_name):
    # This is mainly for the web server to use
    matrix_display(img_to_bw_array(f"images/{image_name}.png"))

def string_display(string):
    matrix = string_to_array(string)
    matrix_display(matrix)

if __name__ == '__main__':
    matrix_display(img_to_bw_array('images/heart.png'), invert=True)
    time.sleep(1)

    for _ in range(0,3):
        matrix_display(img_to_bw_array('images/exclaim.png'), invert=True)
        time.sleep(.3)
        matrix_display(img_to_bw_array('images/exclaim.png'))
        time.sleep(.3)
        
    matrix_display(img_to_bw_array('images/key.png'), invert=True)
    time.sleep(1)
    matrix_display(img_to_bw_array('images/key.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/skull.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/skull3.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/cat.png'), invert=True)
    time.sleep(1)
    matrix_display(img_to_bw_array('images/happy.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/sad.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/sword.png'))
    time.sleep(1)
    matrix_display(img_to_bw_array('images/bow.png'))
    time.sleep(1)
    # matrix_display(device, img_to_bw_array('images/skull3.png'))
    # time.sleep(1)
