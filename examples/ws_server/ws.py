import websocket
import json
import ctypes
import threading
import time

class Request:
    def __init__(self):
        self.user = ""
        self.handler = ""
        self.payload = ""

def on_message(ws, message):
    print("Received message:", message)

def on_error(ws, error):
    print("Error:", error)

def on_close(ws):
    print("Connection closed")

def on_open(ws):
    request = Request()
    request.user = "12D3KooWGQ4ncdUVMSaVrWrCU1fyM8ZdcVvuWa7MdwqkUu4SSDo4"
    request.handler = "MyHandler"
    request.payload = "Hello, world!"

    # Convert the Request object to JSON
    request_json = json.dumps(request.__dict__)
    # Send a test message

    while True:
        time.sleep(0.1)
        ws.send(request_json)


def app():
    library = ctypes.cdll.LoadLibrary('./library.so')
    hello_world = library.start
    hello_world()

if __name__ == "__main__":

    x = threading.Thread(target=app, args=())
    x.start()

    time.sleep(3)

    websocket.enableTrace(True)

    # WebSocket server URL
    ws_url = "ws://localhost:8080/ws"

    # Create WebSocket connection
    ws = websocket.WebSocketApp(ws_url,
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    ws.on_open = on_open

    # Start WebSocket connection
    ws.run_forever()
