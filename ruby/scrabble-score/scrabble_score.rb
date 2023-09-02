=begin
Write your code for the 'Scrabble Score' exercise in this file. Make the tests in
`scrabble_score_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/scrabble-score` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

class Scrabble

  def initialize(word)
    @word = word
  end

  def self.score(word)
    self.new(word).score
  end

  def points(letter)
    {
      1  => %w{ a e i o u l n r s t },
      2  => %w{ d g },
      3  => %w{ b c m p },
      4  => %w{ f h v w y },
      5  => %w{ k },
      8  => %w{ j x },
      10 => %w{ q z},
    }.each do |points,letters|
      return points if letters.include?(letter)
    end
  end

  def score
    @word.to_s.downcase.gsub(/[^a-z]/, '').chars.map{ |letter| points(letter) }.inject(&:+) || 0
  end
end
