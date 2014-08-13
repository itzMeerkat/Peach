#!/usr/bin/env python
# use this version of the code on PCs only

import socket
import time

TCP_SERVER_IP = '127.0.0.1'
TCP_SERVER_PORT = 9301
TCP_MESSAGE = 'tcp comm'

SEND_MSG_REP = 0

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

s.bind((TCP_SERVER_IP, TCP_SERVER_PORT))

s.listen(1)

conn, addr = s.accept()

while 1:
    conn.send(TCP_MESSAGE)
    print addr
    SEND_MSG_REP = SEND_MSG_REP + 1
    if SEND_MSG_REP == 10:
		raw_input("Press any key to continue...")
		break

s.close()
