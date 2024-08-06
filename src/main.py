import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import numpy as np

# Function to read coordinates from file
def read_coordinates(file_path):
    with open(file_path, 'r') as file:
        lines = file.readlines()
    
    coordinates = []
    for line in lines:
        x, y, z = map(float, line.strip().split())
        coordinates.append((x, y, z))
    return coordinates

# File path
file_path = 'dest.txt'

# Read coordinates from file
coords = read_coordinates(file_path)

# Separate the coordinates into x, y, z lists
x_coords = [coord[0] for coord in coords]
y_coords = [coord[1] for coord in coords]
z_coords = [coord[2] for coord in coords]

# Create a 3D plot
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# Plot the points
ax.scatter(x_coords, y_coords, z_coords)

# Set labels
ax.set_xlabel('X Coordinate')
ax.set_ylabel('Y Coordinate')
ax.set_zlabel('Z Coordinate')

# Show the plot
plt.show()

