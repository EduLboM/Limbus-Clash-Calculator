Welcome to the Limbus Company Clashing Calculator! ⛓️🚂

This program calculates the probability of victory in the game Limbus Company, now with a brand-new Graphical User Interface (GUI)! 🖥️🎉 Follow the steps below to set up and start exploring your chances of victory in style.

Program Description 📝

The Limbus Company Clashing Calculator simulates battles between allies and enemies by considering each character's base attack power, number of coins, and sanity levels. Battles are resolved in rounds, where characters with lower attack power lose coins until one side runs out. The coin flips (heads or tails) depend on sanity, which ranges from -45% to +45%.

Key Features 🔍

Graphical User Interface (GUI): The new interface makes it easier and more engaging to input your data and view results!
Power Calculation: Calculates attack power based on base power and coin count.
Monte Carlo Simulations: Estimates win probabilities through numerous simulated battles.
Clear Results: Displays final win probabilities for both allies and enemies.

How to Get Started 🚀

Step 1: Download Julia and Required Packages
Download and install Julia from the official website: JuliaLang.org
Install the following packages by running these commands in the Julia REPL:

    using Pkg
    Pkg.add("Gtk")           # GUI for the program
    Pkg.add("Distributions")  # For handling probability distributions
    Pkg.add("Random")         # For random number generation

Step 2: Run the Program

Launch the Program: Open the code in your Julia environment and run it to start the GUI.
Enter Ally and Enemy Data in the GUI:
Ally Data:
Base Power (BP): Your character's initial power.
Number of Coins (CN): How many coins you have.
Coin Power (CP): Value each coin adds to your power.
Sanity (SN): Affects the probability of coins landing heads.
Enemy Data:
Enemy Base Power (BPEnemy)
Enemy Coins (CNEnemy)
Coin Power for Enemy (CPEnemy)
Enemy Sanity (SNEnemy)
Set Number of Simulations: Choose a high number (10,000+ recommended) for accuracy.
View Results: After running, see your and your enemy’s win probabilities. 🏆

Conclusion 🎉
You're all set to dive into the Limbus Company Clashing Calculator and experience the battles with our enhanced GUI! Have fun exploring your chances of victory, and feel free to reach out with questions or suggestions! 🤗💬
