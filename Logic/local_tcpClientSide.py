#!/usr/bin/env python

import socket
import time

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

TCP_SOURCE_IP = '127.0.0.1'
TCP_SOURCE_PORT = 7723

TCP_DESTINATION_IP = '127.0.0.1'
TCP_DESTINATION_PORT = 9301

TCP_BUFFER_SIZE = 16384

MSG_RECV_REP = 0

s.bind((TCP_SOURCE_IP, TCP_SOURCE_PORT))
s.connect((TCP_DESTINATION_IP, TCP_DESTINATION_PORT))

TCP_RESP = s.recv(TCP_BUFFER_SIZE)

while 1:
    if not TCP_RESP:
        print "No response from TCP server! Exiting now"
        break
    else:
        print "Data from TCP server:", TCP_RESP
        MSG_RECV_REP = MSG_RECV_REP + 1
        if MSG_RECV_REP == 10:
            time.sleep(30)
            break

s.close()
