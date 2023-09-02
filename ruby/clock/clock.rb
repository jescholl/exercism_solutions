=begin
Write your code for the 'Clock' exercise in this file. Make the tests in
`clock_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/clock` directory.
=end

module Minitest
  class Test
    def skip; end
  end
end

class Clock
  attr_reader :hour, :minute

  def initialize(hour: 0, minute: 0)
    @minute = minute % 60
    @hour = (hour + minute/60) % 24
  end

  def to_s
    "%02d:%02d" % [@hour, @minute]
  end

  def -(clock)
    Clock.new(hour: hour - clock.hour, minute: minute - clock.minute)
  end

  def +(clock)
    Clock.new(hour: hour + clock.hour, minute: minute + clock.minute)
  end

  def ==(clock)
    to_s == clock.to_s
  end
end
