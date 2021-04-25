import time
import pathlib

from matplotlib.animation import FuncAnimation
import matplotlib.pyplot as plt
import matplotlib.image as im


frame_counter = 0

start_time = time.time()
current_time = time.time()


def grab_frame():
    global frame_counter

    file = pathlib.Path('/home/krishna/Desktop/network_programming_project/frame_save/img' + str(frame_counter) + '.jpg')
    missing_frame_counter = 0
    while (not file.exists() and missing_frame_counter <= 100):

        # check for the frame every TIMEOUT seconds
        TIMEOUT = 0.2
        time.sleep(TIMEOUT)
        missing_frame_counter += 1
        print("missing frame count = ",missing_frame_counter)

    image = im.imread('/home/krishna/Desktop/network_programming_project/frame_save/img' + str(frame_counter) + '.jpg')
    frame_counter += 1

    current_time = time.time()
    print("fps: ", frame_counter / (current_time - start_time))

    # delete the frame from the disk.
    file.unlink()
    return image

# create axes
ax1 = plt.subplot(111)

# create axes
im1 = ax1.imshow(grab_frame())

def update(i):
    im1.set_data(grab_frame())

# Animate the view every 10 miliseconds
# Keep the interval the same as delay in transmission else the frame directory size grows, could lead to a crash/hang XD

ani = FuncAnimation(plt.gcf(), update, interval=10)
plt.show()