using Distributions, Random, Gtk

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
    return CHWNAlly
end

glade = GtkBuilder(filename="GUI.glade")
Window = glade["Window"]
BPAllyInput = glade["BPAllyInput"]
CNAllyInput = glade["CNAllyInput"]
CPAllyInput = glade["CPAllyInput"]
SNAllyInput = glade["SNAllyInput"]
BPEnemyInput = glade["BPEnemyInput"]
CNEnemyInput = glade["CNEnemyInput"]
CPEnemyInput = glade["CPEnemyInput"]
SNEnemyInput = glade["SNEnemyInput"]
SMNumberInput = glade["SMNumberInput"]
Simulation = glade["Simulation"]
CHWNAllyOutput = glade["CHWNAllyOutput"]
AllyImage = glade["AllyImage"]
EnemyImage = glade["EnemyImage"]

function MainChanceCalculator()
    BPAlly = parse(Int64, get_gtk_property(BPAllyInput, :text, String))
    CNAlly = parse(Int64, get_gtk_property(CNAllyInput, :text, String))
    CPAlly = parse(Int64, get_gtk_property(CPAllyInput, :text, String))
    SNAlly = parse(Int64, get_gtk_property(SNAllyInput, :text, String))
    CHAlly = 0.5 + (SNAlly / 100)
    BPEnemy = parse(Int64, get_gtk_property(BPEnemyInput, :text, String))
    CNEnemy = parse(Int64, get_gtk_property(CNEnemyInput, :text, String))
    CPEnemy = parse(Int64, get_gtk_property(CPEnemyInput, :text, String))
    SNEnemy = parse(Int64, get_gtk_property(SNEnemyInput, :text, String))
    CHEnemy = 0.5 + (SNEnemy / 100)
    SMNumbers = parse(Int64, get_gtk_property(SMNumberInput, :text, String))

    CHWNAlly = monte_carlo_final_simulation(SMNumbers, BPAlly, BPEnemy, CNAlly, CNEnemy, CPAlly, CPEnemy, CHAlly, CHEnemy)
    CHWNAlly = round(CHWNAlly, sigdigits=5) * 100
    return CHWNAlly
end

id = signal_connect(Simulation, "button-press-event") do widget, event
    CHWNAlly = MainChanceCalculator()
    GAccessor.text(CHWNAllyOutput, "ALLY WIN CHANCE: $(CHWNAlly)%")
end

showall(Window)
