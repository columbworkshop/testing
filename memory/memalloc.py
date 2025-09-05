import time 
import logging


num_gb = 2
logging.info(f"Allocate number of gb = {num_gb}")
#
try:
    x = bytearray(num_gb * 1024 * 1024 * 1024)
except Exception as e:
   logging.error(f"An unexpected error occurred: {e}")
   exit(1)

i = 0
while True:
    logging.info(f"Heartbeat {i}")
    i += 1
    time.sleep(30)
