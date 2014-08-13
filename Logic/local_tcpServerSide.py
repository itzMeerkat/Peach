#!/usr/bin/env python

import socket
import time

TCP_SERVER_IP = '127.0.0.1'
TCP_SERVER_PORT = 9301

PRINT_ADDR_REP = 0

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

s.bind((TCP_SERVER_IP, TCP_SERVER_PORT))

s.listen(1)

conn, addr = s.accept()

while 1:
    conn.send("tcp comm")
    print addr
    PRINT_ADDR_REP = PRINT_ADDR_REP + 1
    if PRINT_ADDR_REP == 10:
        time.sleep(35)
        break

s.close()
