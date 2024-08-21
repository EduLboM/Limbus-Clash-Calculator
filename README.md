⏰🚌 Limbus Company Clash Simulation Program 🔥⛓️


-> Overview


This program simulates the clash mechanics in the game Limbus Company using absorbing Markov chains (thanks MusicOnline for the math) to determine the probability of victory for two players based on various in-game factors such as base attack power, buffs, debuffs, coin flips, and sanity levels.


-> Features


Sanity Calculation: Adjusts the probability of a successful coin flip based on the player's sanity.

Buffs and Debuffs: Accounts for power increases/decreases and special effects like "Declared Duel" and "Tremor-Chains".


-> Benefits of Using Go


Performance: Go provides high performance and efficient memory usage, which is beneficial for running complex simulations.

Concurrency: Go’s powerful concurrency model allows for efficient handling of multiple simulations or processes.

Ease of Use: Go’s syntax and comprehensive standard library make it easier to implement mathematical models and handle user input.

Static Typing: Helps catch errors at compile-time, ensuring the robustness of the simulation program.


-> Challenges! this is my first real programming project, there were several challenges:


Implementing Markov Chains: Translating the mathematical model of absorbing Markov chains into efficient Go code.

Handling Buffs/Debuffs: Accurately modeling the various buffs and debuffs and their impact on the game simulation.

Sanity and Coin Flips: Implementing the dynamic probability changes based on sanity levels and ensuring the simulation reflects these changes correctly.

User Input Management: Creating a user-friendly interface for inputting various game parameters and validating those inputs.


-> Installation and Usage


Prerequisites

Go 1.16 or later

Installation

Clone the Repository:

git clone https://github.com/yourusername/limbus-clash-simulation.git

cd limbus-clash-simulation

Install Dependencies:

go get -u gonum.org/v1/gonum

Run the Program:

go run main.go


-> Conclusion


This project has been an excellent learning experience in both programming and game theory. Implementing a Markov chain model in Go for game simulations has provided valuable insights into the power and efficiency of Go for complex mathematical and statistical computations.
For further development, additional game mechanics and more sophisticated models can be incorporated to enhance the simulation's accuracy and depth.

Feel free to reach out with any questions or suggestions for improving the program. Thank you for using the Limbus Company Clash Simulation Program!

🏆
