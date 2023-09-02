=begin
Write your code for the 'Bowling' exercise in this file. Make the tests in
`bowling_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/bowling` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end


class Game
  class BowlingError < ArgumentError; end

  class Frame
    def initialize
      @rolls = []
      @bonus = []
    end

    def roll(pins)
      raise(BowlingError, "Invalid roll, score too high") if score + pins > 10

      @rolls << pins
    end

    def score
      @rolls.sum + @bonus.sum
    end

    def bonus(pins)
      unless bonus_complete?
        raise(BowlingError, "Invalid bonus pin count") if @bonus.fetch(0, 0) != 10 and @bonus.fetch(0, 0) + pins > 10
        @bonus << pins
      end
    end

    def strike?
      @rolls[0] == 10
    end

    def spare?
      score == 10 and not strike?
    end

    def complete?
      strike? or @rolls.length == 2
    end

    def bonus_complete?
      (strike? and @bonus.length == 2) or (spare? and @bonus.length == 1) or (not strike? and not spare?)
    end
  end

  def initialize
    @frames = Array.new(10) { Frame.new }
    @current_frame = 0
    @rolls = []
  end

  def roll(pins)
    raise(BowlingError, "Invalid pin count") if pins > 10 or pins < 0
    raise(BowlingError, "Game has ended") if remaining_frames == 0 and frame.complete? and frame.bonus_complete?

    @rolls[-1]&.bonus(pins) if @rolls[-1]&.complete?
    @rolls[-2]&.bonus(pins) if @rolls[-2]&.complete?

    unless frame.complete?
      frame&.roll(pins)
    end

    @rolls << frame

    if frame&.complete? and remaining_frames > 0
      @current_frame += 1
    end
  end

  def remaining_frames
    @frames.length - 1 - @current_frame
  end

  def frame
    @frames[@current_frame]
  end

  def score
    raise(BowlingError, "Game still in progress") if remaining_frames > 0 or not @frames.all?{ |f| f.complete? and f.bonus_complete? }

    @frames.map(&:score).sum
  end
end
