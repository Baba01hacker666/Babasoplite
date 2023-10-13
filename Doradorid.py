print("welcome to doradorid")
import os 
import sys
import requests
ip = sys.argv[1]
r = requests.get('http://' + (ip), verify=False)
if r.status_code == 200:
  print("[+]no account authorized nedded")
if r.status_code == 401:
  print("[+]account authrozation needed")
if r.status_code == 404:
  print("soory web page not found")
if r.status_code == 403:
  print("webpage not allowed to be open")
if r.status_code == 502:
  print("bad getway")