import board
import digitalio
import adafruit_character_lcd.character_lcd as character_lcd

import time

lcd_rs = digitalio.DigitalInOut(board.D17)
lcd_en = digitalio.DigitalInOut(board.D27)
lcd_d4 = digitalio.DigitalInOut(board.D22)
lcd_d5 = digitalio.DigitalInOut(board.D23)
lcd_d6 = digitalio.DigitalInOut(board.D24)
lcd_d7 = digitalio.DigitalInOut(board.D25)
# lcd_backlight = digitalio.DigitalInOut(board.D13)

lcd_columns = 16
lcd_rows = 2


# lcd = character_lcd.Character_LCD_Mono(lcd_rs, lcd_en, lcd_d4, lcd_d5, lcd_d6, lcd_d7, lcd_columns, lcd_rows, lcd_backlight)
lcd = character_lcd.Character_LCD_Mono(lcd_rs, lcd_en, lcd_d4, lcd_d5, lcd_d6, lcd_d7, lcd_columns, lcd_rows)
lcd.cursor = True

def type_message(message, delay=0.1):
    for x in range(0, len(message)):
        print(message[:x+1])
        lcd.message = message[:x+1]
        time.sleep(delay)

if __name__ == "__main__":
    message = "Hello there!\nHow can I help?"

    type_message(message)
    time.sleep(2)

    for i in range(16):
        lcd.move_left()
        time.sleep(0.05)

    lcd.clear()
    lcd.message = message
    time.sleep(2)

    for i in range(160):
        lcd.move_right()
        time.sleep(0.05)
