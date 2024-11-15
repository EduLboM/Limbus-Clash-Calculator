using Distributions, Random

function round_simulation(BPAlly::Int, BPEnemy::Int, CNAlly::Int, CNEnemy::Int, CPAlly::Int, CPEnemy::Int, CHAlly::Float64, CHEnemy::Float64)
    PWAlly = BPAlly + (CPAlly * rand(Binomial(CNAlly, CHAlly)))
    PWEnemy = BPEnemy + (CPEnemy * rand(Binomial(CNEnemy, CHEnemy)))

    if PWAlly > PWEnemy
        return :Ally
    elseif PWEnemy > PWAlly
        return :Enemy
    else
        return :Tie
    end
end

function monte_carlo_final_simulation(SMNumbers::Int, BPAlly::Int, BPEnemy::Int, CNAlly::Int, CNEnemy::Int, CPAlly::Int, CPEnemy::Int, CHAlly::Float64, CHEnemy::Float64)
    WNally = 0
    WNEnemy = 0

    for _ in 1:SMNumbers
        CNSMAlly = CNAlly
        CNSMEnemy = CNEnemy

        while CNSMAlly > 0 && CNSMEnemy > 0
            result = round_simulation(BPAlly, BPEnemy, CNSMAlly, CNSMEnemy, CPAlly, CPEnemy, CHAlly, CHEnemy)

            if result == :Ally
                CNSMEnemy -= 1
            elseif result == :Enemy
                CNSMAlly -= 1
            end
        end

        if CNSMEnemy == 0
            WNally += 1
        elseif CNSMAlly == 0
            WNEnemy += 1
        end
    end

    CHWNAlly = WNally / SMNumbers
    CHWNEnemy = WNEnemy / SMNumbers

    return CHWNAlly, CHWNEnemy
end


println("LIMBUS COMPANY CLASHING CALCULATOR ⛓️🚂")

println("\nSinner's Attack 😇")
println("What's your base power?")
BPAlly = parse(Int64, readline())
println("How many coins do you have?")
CNAlly = parse(Int64, readline())
println("What's your coin power?")
CPAlly = parse(Int64, readline())
println("What's your sanity?")
SNAlly = parse(Int64, readline())
CHAlly = 0.5 + (SNAlly / 100)

println("\nEnemy's Attack 👿")
println("What's the enemy base power?")
BPEnemy = parse(Int64, readline())
println("How many coins does the enemy have?")
CNEnemy = parse(Int64, readline())
println("What's the enemy's coin power?")
CPEnemy = parse(Int64, readline())
println("What's the enemy's sanity?")
SNEnemy = parse(Int64, readline())
CHEnemy = 0.5 + (SNEnemy / 100)

println("How many simulations? (recommended: 10k+)")
SMNumbers = parse(Int64, readline())

CHWNAlly, CHWNEnemy = monte_carlo_final_simulation(SMNumbers, BPAlly, BPEnemy, CNAlly, CNEnemy, CPAlly, CPEnemy, CHAlly, CHEnemy)
CHWNAlly = round(CHWNAlly, sigdigits=5) * 100
CHWNEnemy = round(CHWNEnemy, sigdigits=5) * 100

println("\nFinal Win Probability:")
println("Ally Win Probability: $(CHWNAlly)%")
println("Enemy Win Probability: $(CHWNEnemy)%")