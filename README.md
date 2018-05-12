## About

Package ode-nm is to solve ODEs with different numerical methods and
compare their accuracy watching methods plots.

The task is like this. We have a system of equations:

    1. y' = f(x, y)
    2. y(0) = y0
    3. x <- [0, 1]

Solving the system manually we get the precise solution, calculate
the partial f(x, y) derivatives and hardcode all these functions.

Then all methods implemented are applied to the system and plots
are built. Watching plots we can visually compare accuracy of all methods
since there is also the plot of the precise solution.
