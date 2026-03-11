import time

def tick():
    print("Tick...")

def tock():
    print("Tock...")

if __name__ == "__main__":
    print("Loop app starting...")
    for i in range(10):
        time.sleep(1)
        tick()
        time.sleep(1)
        tock()
