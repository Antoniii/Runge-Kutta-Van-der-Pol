import numpy  as np
import matplotlib.pyplot as plt

data = np.loadtxt('data.txt')


x = data[:, 0]
y = data[:, 1]
plt.plot(x, y,'r--')
plt.title('Runge–Kutta’s Method for Van der Pol oscillator')
plt.xlabel('t')
plt.ylabel('x')
plt.grid()
plt.savefig("fig.png")
plt.show()