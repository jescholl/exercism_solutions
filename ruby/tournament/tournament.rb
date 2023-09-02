=begin
Write your code for the 'Tournament' exercise in this file. Make the tests in
`tournament_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/tournament` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

module Tournament
  class Team
    attr_accessor :wins, :losses, :draws

    def initialize
      @wins = @losses = @draws = 0
    end

    def win
      @wins += 1
    end

    def loss
      @losses += 1
    end

    def draw
      @draws += 1
    end

    def matches
      wins + losses + draws
    end

    def points
      wins * 3 + draws
    end
  end

  @winloss_flip = {
    "win" => "loss",
    "loss" => "win",
    "draw" => "draw"
  }

  def self.tally(input)
    teams = Hash.new { |hash,key| hash[key] = Team.new}
    input.each_line do |line|
      team1, team2, winloss = line.chomp.split(";")
      if team1 and team2 and winloss
        teams[team1].__send__(winloss.to_sym)
        teams[team2].__send__(@winloss_flip[winloss].to_sym)
      end
    end

    results = [%w{Team MP W D L P}]

    teams.sort.sort{ |(_,team1),(_,team2)| (team1.points <=> team2.points) * -1 }.each do |name, team|
      results << [name, team.matches, team.wins, team.draws, team.losses, team.points]
    end

    results.map do |line|
      sprintf("%-30s | %2s | %2s | %2s | %2s | %2s\n", *line)
    end.join
  end
end
